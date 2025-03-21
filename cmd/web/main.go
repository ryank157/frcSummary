package main

import (
	"fmt"
	"log"
	"net/http"

	"frcSummary/internal/config"
	"frcSummary/internal/handler"
	"frcSummary/internal/llm"
	"frcSummary/internal/service"
	"frcSummary/internal/statbotics" // Import statbotics package

	"context"
	"frcSummary/internal/utils"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize Logger
	fmt.Printf("level %s", cfg.LogLevel)
	logger := utils.NewLogger(cfg.LogLevel)

	// Initialize OpenRouter Client
	openRouterClient := llm.NewOpenRouterClient(cfg.OpenRouterAPIKey)

	statboticsClient := statbotics.NewClient(cfg.StatboticsUrl, nil) // Replace with your Statbotics Base URL from config
	// Example usage of the Statbotics client (replace with your actual use case)
	ctx := context.Background()
	defaultResponse, err := statboticsClient.GetDefault(ctx)
	if err != nil {
		logger.Errorf("Failed to get default from Statbotics: %v", err)
	} else {
		logger.Infof("Statbotics Default Response: %+v", defaultResponse)
	}

	// Initialize Services
	analysisService := service.NewAnalysisService(openRouterClient)

	// Initialize Handlers
	analysisHandler := handler.NewAnalysisHandler(analysisService)
	homeHandler := handler.NewHomeHandler()
	matchHandler := handler.NewMatchHandler(statboticsClient) // Initialize MatchHandler

	// Set up Routes
	http.HandleFunc("/analyze", analysisHandler.Analyze)
	http.HandleFunc("/", homeHandler.Home)
	http.HandleFunc("/match", matchHandler.Match) // Add the /match route

	// Serve Static Files
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/web/", http.StripPrefix("/web/", fs))

	addr := fmt.Sprintf(":%d", cfg.Port)
	logger.Infof("Server listening on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}
