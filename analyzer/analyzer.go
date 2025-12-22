package analyzer

import (
	"fmt"
	"time"

	"NebulaChallenge/client"
	"NebulaChallenge/models"
	"NebulaChallenge/utils"
)

// Analyzer orquesta el análisis de SSL
type Analyzer struct {
	client *client.Client
}

// NewAnalyzer crea una nueva instancia del analizador
func NewAnalyzer() *Analyzer {
	return &Analyzer{
		client: client.NewClient(),
	}
}

// Run ejecuta el análisis completo de un host
func (a *Analyzer) Run(host string, publish bool) (*models.Host, error) {
	// 1. Validar el host
	if err := utils.ValidateHost(host); err != nil {
		return nil, fmt.Errorf("invalid host: %w", err)
	}

	// 2. Verificar disponibilidad del servicio
	info, err := a.client.GetInfo()
	if err != nil {
		return nil, fmt.Errorf("SSL Labs service unavailable: %w", err)
	}

	fmt.Printf("SSL Labs API v%s (Criteria: %s)\n", info.Version, info.CriteriaVersion)
	fmt.Printf("Max concurrent assessments: %d\n", info.MaxAssessments)
	fmt.Printf("Current assessments: %d\n\n", info.CurrentAssessments)

	// 3. Iniciar análisis
	fmt.Printf("Starting analysis for %s...\n", host)
	result, err := a.client.StartAnalysis(host, publish)
	if err != nil {
		return nil, fmt.Errorf("error starting analysis: %w", err)
	}

	// 4. Si ya está listo, retornar
	if result.Status == "READY" || result.Status == "ERROR" {
		return result, nil
	}

	// 5. Hacer polling hasta que termine
	fmt.Println("Analysis in progress...")
	return a.pollAnalysis(host)
}

// pollAnalysis hace polling periódico hasta que el análisis termine
func (a *Analyzer) pollAnalysis(host string) (*models.Host, error) {
	var pollInterval time.Duration
	inProgress := false

	for {
		// Polling variable: 5s hasta IN_PROGRESS, luego 10s
		if inProgress {
			pollInterval = 10 * time.Second
		} else {
			pollInterval = 5 * time.Second
		}

		time.Sleep(pollInterval)

		result, err := a.client.CheckAnalysis(host)
		if err != nil {
			return nil, fmt.Errorf("error checking analysis: %w", err)
		}

		fmt.Printf("Status: %s", result.Status)

		// Mostrar progreso de endpoints
		if len(result.Endpoints) > 0 {
			fmt.Print(" | Endpoints: ")
			for i, ep := range result.Endpoints {
				if i > 0 {
					fmt.Print(", ")
				}
				fmt.Printf("%s (%d%%)", ep.IPAddress, ep.Progress)
			}
		}
		fmt.Println()

		switch result.Status {
		case "READY":
			fmt.Println("\nAnalysis complete!")
			return result, nil

		case "ERROR":
			return nil, fmt.Errorf("analysis failed: %s", result.StatusMessage)

		case "IN_PROGRESS":
			inProgress = true

		case "DNS":
			fmt.Println("Resolving DNS...")
		}
	}
}
