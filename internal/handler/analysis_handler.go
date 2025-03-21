package handler

import (
	"log"
	"net/http"

	"frcSummary/internal/model"
	"frcSummary/internal/service"
	"frcSummary/web/templates"
)

type AnalysisHandler struct {
	analysisService *service.AnalysisService
}

func NewAnalysisHandler(analysisService *service.AnalysisService) *AnalysisHandler {
	return &AnalysisHandler{analysisService: analysisService}
}

func (h *AnalysisHandler) Analyze(w http.ResponseWriter, r *http.Request) {
	// 1. Extract parameters from the URL query
	// vars := r.URL.Query()
	// var1 := vars.Get("var1")
	// var2Str := vars.Get("var2")
	// var3Str := vars.Get("var3")

	// // 2. Validate parameters
	// if var1 == "" || var2Str == "" {
	// 	http.Error(w, "Missing parameters", http.StatusBadRequest)
	// 	return
	// }

	// var2, err := strconv.Atoi(var2Str)
	// if err != nil {
	// 	http.Error(w, "Invalid var2", http.StatusBadRequest)
	// 	return
	// }

	// var3, err := strconv.ParseBool(var3Str)
	// if err != nil {
	// 	http.Error(w, "Invalid var3", http.StatusBadRequest)
	// 	return
	// }

	// 3. Create AnalysisRequest
	request := model.AnalysisRequest{}

	resp, err := h.analysisService.PerformAnalysis(request)
	if err != nil {
		log.Printf("Error performing analysis: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 4. Render the analysis template only
	err = templates.Analysis(request, resp).Render(r.Context(), w)
	if err != nil {
		log.Printf("Error rendering template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
