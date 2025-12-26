package formatter

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"NebulaChallenge/models"
)

// PrintReport imprime el reporte de forma legible
func PrintReport(host *models.Host) {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Printf("SSL/TLS SECURITY ASSESSMENT REPORT\n")
	fmt.Println(strings.Repeat("=", 80))

	fmt.Printf("\nHost: %s\n", host.Host)
	fmt.Printf("Port: %d\n", host.Port)
	fmt.Printf("Protocol: %s\n", host.Protocol)
	fmt.Printf("Status: %s\n", host.Status)

	if host.TestTime > 0 {
		testTime := time.Unix(host.TestTime/1000, 0)
		fmt.Printf("Test Time: %s\n", testTime.Format("2006-01-02 15:04:05 MST"))
	}

	fmt.Printf("\nEngine Version: %s\n", host.EngineVersion)
	fmt.Printf("Criteria Version: %s\n", host.CriteriaVersion)

	// Endpoints
	fmt.Printf("\n%s\n", strings.Repeat("-", 80))
	fmt.Printf("ENDPOINTS (%d)\n", len(host.Endpoints))
	fmt.Printf("%s\n", strings.Repeat("-", 80))

	for i, ep := range host.Endpoints {
		printEndpoint(i+1, &ep)
	}
}

func printEndpoint(num int, ep *models.Endpoint) {
	fmt.Printf("\n[%d] IP Address: %s\n", num, ep.IPAddress)

	if ep.ServerName != "" {
		fmt.Printf("    Server Name: %s\n", ep.ServerName)
	}

	// Grade con color (simulado)
	gradeColor := getGradeDisplay(ep.Grade)
	fmt.Printf("    Grade: %s\n", gradeColor)

	if ep.GradeTrustIgnored != "" && ep.GradeTrustIgnored != ep.Grade {
		fmt.Printf("    Grade (Trust Ignored): %s\n", ep.GradeTrustIgnored)
	}

	fmt.Printf("    Status: %s\n", ep.StatusMessage)

	if ep.HasWarnings {
		fmt.Printf("Has Warnings\n")
	}

	if ep.IsExceptional {
		fmt.Printf("Exceptional Configuration\n")
	}

	// Detalles si están disponibles
	if ep.Details != nil {
		printEndpointDetails(ep.Details)
	}
}

func printEndpointDetails(details *models.EndpointDetails) {
	fmt.Println("\n    === DETAILED INFORMATION ===")

	// Protocolos
	if len(details.Protocols) > 0 {
		fmt.Printf("\n    Supported Protocols:\n")
		for _, proto := range details.Protocols {
			fmt.Printf("      - %s %s\n", proto.Name, proto.Version)
		}
	}

	// Certificado
	fmt.Printf("\n    Certificate:\n")
	fmt.Printf("      Subject: %s\n", details.Cert.Subject)
	fmt.Printf("      Issuer: %s\n", details.Cert.IssuerLabel)

	notBefore := time.Unix(details.Cert.NotBefore/1000, 0)
	notAfter := time.Unix(details.Cert.NotAfter/1000, 0)
	fmt.Printf("      Valid From: %s\n", notBefore.Format("2006-01-02"))
	fmt.Printf("      Valid Until: %s\n", notAfter.Format("2006-01-02"))

	// Key
	fmt.Printf("\n    Key:\n")
	fmt.Printf("      Algorithm: %s\n", details.Key.Alg)
	fmt.Printf("      Size: %d bits\n", details.Key.Size)
	fmt.Printf("      Strength: %d bits\n", details.Key.Strength)

	// Vulnerabilidades
	fmt.Printf("\n    Security Issues:\n")
	printVulnerabilities(details)

	// Cipher Suites (primeros 5)
	if len(details.Suites.List) > 0 {
		fmt.Printf("\n    Cipher Suites (showing first 5 of %d):\n", len(details.Suites.List))
		maxShow := 5
		if len(details.Suites.List) < maxShow {
			maxShow = len(details.Suites.List)
		}
		for i := 0; i < maxShow; i++ {
			suite := details.Suites.List[i]
			fmt.Printf("      - %s (%d bits)\n", suite.Name, suite.CipherStrength)
		}
	}
}

func printVulnerabilities(details *models.EndpointDetails) {
	hasVulnerabilities := false

	if details.VulnBeast {
		fmt.Printf("    BEAST: Vulnerable\n")
		hasVulnerabilities = true
	}

	if details.Heartbleed {
		fmt.Printf("    Heartbleed: Vulnerable\n")
		hasVulnerabilities = true
	}

	if details.Poodle {
		fmt.Printf("     POODLE (SSL): Vulnerable\n")
		hasVulnerabilities = true
	}

	if details.PoodleTls == 2 {
		fmt.Printf("     POODLE (TLS): Vulnerable\n")
		hasVulnerabilities = true
	}

	if details.Freak {
		fmt.Printf("      FREAK: Vulnerable\n")
		hasVulnerabilities = true
	}

	if details.Logjam {
		fmt.Printf("      Logjam: Vulnerable\n")
		hasVulnerabilities = true
	}

	if details.Rc4Only {
		fmt.Printf("      RC4 Only\n")
		hasVulnerabilities = true
	}

	if !hasVulnerabilities {
		fmt.Printf("      No major vulnerabilities detected\n")
	}

	// Forward Secrecy
	fmt.Printf("\n      Forward Secrecy: %s\n", getForwardSecrecyStatus(details.ForwardSecrecy))
}

func getGradeDisplay(grade string) string {
	// En una terminal real, podrías usar colores ANSI
	switch grade {
	case "A+", "A", "A-":
		return "🟢 " + grade
	case "B":
		return "🟡 " + grade
	case "C", "D":
		return "🟠 " + grade
	case "F", "T", "M":
		return "🔴 " + grade
	default:
		return grade
	}
}

func getForwardSecrecyStatus(fs int) string {
	if fs&4 != 0 {
		return "Yes (with all simulated clients)"
	} else if fs&2 != 0 {
		return "Yes (with modern clients)"
	} else if fs&1 != 0 {
		return "Yes (with some clients)"
	}
	return "No"
}

// ExportJSON exporta el resultado a JSON
func ExportJSON(host *models.Host) (string, error) {
	data, err := json.MarshalIndent(host, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error marshaling to JSON: %w", err)
	}
	return string(data), nil
}
