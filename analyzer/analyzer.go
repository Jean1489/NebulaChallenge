package analyzer

import (
	_ "context"
	"fmt"
	"strings"
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
	sanitizedHost := utils.SanitizeHost(host)
	if err := utils.ValidateHost(sanitizedHost); err != nil {
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
	fmt.Printf("Starting analysis for %s...\n", sanitizedHost)
	result, err := a.client.StartAnalysis(sanitizedHost, publish)
	if err != nil {
		return nil, fmt.Errorf("error starting analysis: %w", err)
	}

	// 4. Si ya está listo, retornar
	if client.IsAnalysisComplete(result.Status) {
		if result.Status == "ERROR" {
			return nil, fmt.Errorf("analysis failed: %s", result.StatusMessage)
		}
		return result, nil
	}

	// 5. Hacer polling hasta que termine
	fmt.Println("Analysis in progress...")
	return a.pollAnalysis(sanitizedHost)
}

// pollAnalysis hace polling periódico hasta que el análisis termine
func (a *Analyzer) pollAnalysis(host string) (*models.Host, error) {
	pollInterval := 5 * time.Second
	inProgress := false

	for {
		time.Sleep(pollInterval)

		result, err := a.client.CheckAnalysis(host)
		if err != nil {
			return nil, fmt.Errorf("error checking analysis: %w", err)
		}

		// Mostrar progreso
		a.printProgress(result)

		// Verificar si terminó
		if client.IsAnalysisComplete(result.Status) {
			fmt.Print("\r" + strings.Repeat(" ", 100) + "\r") // Limpiar línea

			if result.Status == "ERROR" {
				return nil, fmt.Errorf("analysis failed: %s", result.StatusMessage)
			}

			fmt.Println("\n Analysis complete!")
			return result, nil
		}

		// Ajustar intervalo cuando entra en progreso
		if result.Status == "IN_PROGRESS" && !inProgress {
			inProgress = true
			pollInterval = 10 * time.Second
		}
	}
}

// printProgress muestra el progreso actual
func (a *Analyzer) printProgress(result *models.Host) {
	fmt.Printf("\rStatus: %-15s", result.Status)

	if len(result.Endpoints) > 0 {
		fmt.Print(" | Endpoints: ")
		for i, ep := range result.Endpoints {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%s (%d%%)", ep.IPAddress, ep.Progress)
		}
	}
}
