package handler

import (
	"encoding/json"
	"frcSummary/internal/llm"
	"frcSummary/internal/statbotics" // Import statbotics packagestrconv
	"net/http"

	"context"
	"frcSummary/web/templates"
	"log"
)

// MatchHandler handles requests to get match data from Statbotics.
type MatchHandler struct {
	StatboticsClient *statbotics.Client
	llmClient        llm.LLMClient
}

// NewMatchHandler creates a new MatchHandler.
func NewMatchHandler(statboticsClient *statbotics.Client, llmClient llm.LLMClient) *MatchHandler {
	return &MatchHandler{StatboticsClient: statboticsClient, llmClient: llmClient}
}

// Match retrieves match data from Statbotics and renders it.
func (h *MatchHandler) Match(w http.ResponseWriter, r *http.Request) {
	matchKey := r.URL.Query().Get("matchKey")
	if matchKey == "" {
		http.Error(w, "matchKey parameter is required", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	matchResponse, err := h.StatboticsClient.GetMatch(ctx, matchKey)
	log.Printf("Match Response: %+v", matchResponse)
	if err != nil {
		log.Printf("Failed to get match data from Statbotics: %v", err)
		http.Error(w, "Failed to get match data", http.StatusInternalServerError)
		return
	}

	matchResponseBytes, err := json.Marshal(matchResponse)
	if err != nil {
		// Handle the error appropriately.  Perhaps log it and return an error.
		// The simplest (but potentially dangerous) thing to do would be to panic:
		panic(err)
		// Or, return an error to the caller, e.g.:
		// return "", fmt.Errorf("failed to marshal matchResponse: %w", err)
	}
	strMatchRes := string(matchResponseBytes)
	llmRes, err := h.llmClient.Generate(strMatchRes)
	if err != nil {

	}

	// Render the match data in a user-friendly way.  You'll need to create a template for this.
	// or you can use a simple string-based output for now.
	// For now, just output the raw JSON. In a real application you would use
	// a html template.
	err = templates.Match_Analysis(llmRes).Render(ctx, w)
	if err != nil {
		log.Printf("Error rendering template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
