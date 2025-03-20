package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/base.html") // Adjust path as needed
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Message":      "Hello, HTMX!",
		"TemplateName": "home", // Add this line to specify home tempalte
	}

	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil{
			fmt.Println("Failed to parsing HTML")
			return
	}
}
