package handler

import (
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

	resp, err := h.analysisService.PerformAnalysis(request)
	if err != nil {
		log.Printf("Error performing analysis: %v", err)
		//Added Model if the analysis does not connect properly.
		response := model.AnalysisResponse{
			Result:      "Default",
			Explanation: "error, unable to connect to LLM",
		}

		data := map[string]interface{}{
			"Request":      request,
			"Response":     response,
			"TemplateName": "analysis", // Pass   template name to the base
		}

		// 5. Render the response to the analysis.html template
		tmpl, err := template.ParseFiles("web/templates/base.html", "web/templates/analysis.html") // Assuming base template includes content
		if err != nil {
			log.Printf("Error parsing template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
			log.Printf("Error executing template: %v. data info is %+v", err, data)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError) // Proper Handling
			return
		}
		return //MAKE SURE to have a RETURN STATMENT here since it continues to load if it does not return
	}

	data := map[string]interface{}{
		"Request":      request,
		"Response":     resp,
		"TemplateName": "analysis", // Pass template name to the base
	}
	tmpl, err := template.ParseFiles("web/templates/base.html", "web/templates/analysis.html") // Assuming base template includes content

	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		log.Printf("Error executing template: %v.Data info is %+v", err, data) //Print statement to verify data output
		http.Error(w, "Internal Server Error", http.StatusInternalServerError) //Proper Error and error

		return
	}

	// fmt.Fprintf(w, "Response: %s", response.Result)
}
