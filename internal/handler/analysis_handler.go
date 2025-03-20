package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"frcSummary/internal/model"
	"frcSummary/internal/service"
)

type AnalysisHandler struct {
	analysisService *service.AnalysisService
}

func NewAnalysisHandler(analysisService *service.AnalysisService) *AnalysisHandler {
	return &AnalysisHandler{analysisService: analysisService}
}

func (h *AnalysisHandler) Analyze(w http.ResponseWriter, r *http.Request) {
	// 1. Extract parameters from the URL query
	vars := r.URL.Query()
	var1 := vars.Get("var1")
	var2Str := vars.Get("var2")
	var3Str := vars.Get("var3")

	// 2. Validate parameters
	if var1 == "" || var2Str == "" || var3Str == "" {
		http.Error(w, "Missing parameters", http.StatusBadRequest)
		return
	}

	var2, err := strconv.Atoi(var2Str)
	if err != nil {
		http.Error(w, "Invalid var2", http.StatusBadRequest)
		return
	}

	var3, err := strconv.ParseBool(var3Str)
	if err != nil {
		http.Error(w, "Invalid var3", http.StatusBadRequest)
		return
	}

	// 3. Create AnalysisRequest
	request := model.AnalysisRequest{
		Var1: var1,
		Var2: var2,
		Var3: var3,
	}

	// 4. Call the AnalysisService
	response, err := h.analysisService.PerformAnalysis(request)
	if err != nil {
		log.Printf("Error performing analysis: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 5. Render the response to the analysis.html template
	tmpl, err := template.ParseFiles("web/templates/base.html") // Assuming base template includes content
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Request":      request,
		"Response":     response,
		"TemplateName": "analysis", // Pass template name to the base
	}

	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		fmt.Println("Failed to parsing HTML")
		return
	}

	// fmt.Fprintf(w, "Response: %s", response.Result)
}
