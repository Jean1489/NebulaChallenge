package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"NebulaChallenge/analyzer"
	"NebulaChallenge/formatter"
)

func main() {
	// Definir flags
	hostPtr := flag.String("host", "", "Hostname to analyze (required)")
	publishPtr := flag.Bool("publish", false, "Publish results on SSL Labs boards")
	jsonPtr := flag.Bool("json", false, "Output results as JSON")
	helpPtr := flag.Bool("help", false, "Show help")

	flag.Parse()

	if *helpPtr {
		printHelp()
		os.Exit(0)
	}

	if *hostPtr == "" {
		fmt.Println("Error: --host flag is required")
		fmt.Println("Use --help for more information")
		os.Exit(1)
	}

	// Setup context para manejar Ctrl+C
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Capturar Ctrl+C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\n\nAnalysis cancelled by user")
		cancel()
		os.Exit(0)
	}()

	// Crear analizador y ejecutar
	a := analyzer.NewAnalyzer()

	result, err := a.Run(*hostPtr, *publishPtr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Mostrar resultados
	if *jsonPtr {
		jsonOutput, err := formatter.ExportJSON(result)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error exporting JSON: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(jsonOutput)
	} else {
		formatter.PrintReport(result)
	}
}

func printHelp() {
	fmt.Println("Nebula Challenge - SSL Labs Security Scanner")
	fmt.Println("\nUsage:")
	fmt.Println("  nebula-challenge --host=<hostname> [options]")
	fmt.Println("\nOptions:")
	fmt.Println("  --host string      Hostname to analyze (required)")
	fmt.Println("  --publish          Publish results on SSL Labs public boards")
	fmt.Println("  --json             Output results as JSON")
	fmt.Println("  --help             Show this help message")
	fmt.Println("\nExamples:")
	fmt.Println("  go run . --host=google.com")
	fmt.Println("  go run . --host=facebook.com --json")
	fmt.Println("  go run . --host=github.com --publish")
}
