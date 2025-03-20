package handler

import (
	"html/template"
	"log"
	"net/http"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/base.html", "web/templates/home.html") // Adjust path as needed
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Message":      "Hello, HTMX!",
		"TemplateName": "home", // Add this line to specify home tempalte
	}

	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		log.Printf("Error executing template: %v. Data is %v. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError) // Proper Handling
		return
	}
}
