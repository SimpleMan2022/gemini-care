package usecase

import (
	"encoding/json"
	"fmt"
	"gemini-care/dto"
	"gemini-care/external/gemini"
	"strings"
)

type SymptomUsecase interface {
	SymptomChecker(req *dto.SymptomsRequest) (*dto.SymptomResponse, error)
}

type symptomUsecase struct {
	gemini.GeminiClient
}

func NewSymptomUsecase(geminiClient gemini.GeminiClient) *symptomUsecase {
	return &symptomUsecase{GeminiClient: geminiClient}
}

func (s *symptomUsecase) SymptomChecker(req *dto.SymptomsRequest) (*dto.SymptomResponse, error) {
	question := "saya merasa "
	for _, quest := range req.Symptoms {
		question += quest + ", "
	}
	res, err := s.GeminiClient.AnswerQuestion(question)
	if err != nil {
		return nil, err
	}
	cleanedRes := strings.TrimSpace(res)
	cleanedRes = strings.TrimPrefix(cleanedRes, "```")
	cleanedRes = strings.TrimPrefix(cleanedRes, "json")
	cleanedRes = strings.TrimSuffix(cleanedRes, "```")

	var symptomResponse *dto.SymptomResponse
	err = json.Unmarshal([]byte(cleanedRes), &symptomResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return symptomResponse, nil
}
