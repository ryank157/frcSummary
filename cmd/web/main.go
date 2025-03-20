package main

import (
	"fmt"
	"log"
	"net/http"

	"frcSummary/internal/config"
	"frcSummary/internal/handler"
	"frcSummary/internal/llm"
	"frcSummary/internal/service"
	"frcSummary/internal/utils"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize Logger
	logger := utils.NewLogger(cfg.LogLevel)

	// Initialize OpenRouter Client
	openRouterClient := llm.NewOpenRouterClient(cfg.OpenRouterAPIKey)

	// Initialize Services
	analysisService := service.NewAnalysisService(openRouterClient)

	// Initialize Handlers
	analysisHandler := handler.NewAnalysisHandler(analysisService)
	homeHandler := handler.NewHomeHandler()

	// Set up Routes
	http.HandleFunc("/analyze", analysisHandler.Analyze)
	http.HandleFunc("/", homeHandler.Home)

	// Serve Static Files
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/web/", http.StripPrefix("/web/", fs))

	addr := fmt.Sprintf(":%d", cfg.Port)
	logger.Infof("Server listening on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}
