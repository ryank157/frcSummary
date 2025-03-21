package service

import (
	"fmt"
	"frcSummary/internal/llm"
	"frcSummary/internal/model"
)

type AnalysisService struct {
	llmClient llm.LLMClient
}

func NewAnalysisService(llmClient llm.LLMClient) *AnalysisService {
	return &AnalysisService{llmClient: llmClient}
}

func (s *AnalysisService) PerformAnalysis(request model.AnalysisRequest) (model.AnalysisResponse, error) {
	// 1. Construct the prompt for the LLM
	prompt := fmt.Sprintf("Respond with an otter ascii")

	// 2. Call the LLM client
	llmResponse, err := s.llmClient.Generate(prompt)
	if err != nil {
		return model.AnalysisResponse{}, fmt.Errorf("failed to generate analysis: %w", err)
	}

	// 3.  Process the response (Optional: this can contain parsing, validation etc)
	response := model.AnalysisResponse{
		Result:      llmResponse,
		Explanation: "Analysis by LLM",
	}

	return response, nil
}
