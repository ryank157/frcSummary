package handler

import (
	"frcSummary/web/templates" // Import the generated template package
	"log"
	"net/http"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) Home(w http.ResponseWriter, r *http.Request) {
	// Render the Home template within the Base template
	homeContent := templates.Home()
	err := templates.Base("Home", homeContent).Render(r.Context(), w) // Use Base template
	if err != nil {
		log.Printf("Error rendering template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
