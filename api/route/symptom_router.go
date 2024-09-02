package route

import (
	"gemini-care/api/handler"
	"gemini-care/external/gemini"
	"gemini-care/usecase"
	"github.com/labstack/echo/v4"
)

func SymptomRouter(r *echo.Group) {
	geminiClient := gemini.NewGeminiClient()
	symptomUsecase := usecase.NewSymptomUsecase(geminiClient)
	symptomHandler := handler.NewSymptomHandler(symptomUsecase)
	r.POST("/symptom-checker", symptomHandler.SymptomChecker)
}
