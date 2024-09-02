package handler

import (
	"fmt"
	"gemini-care/dto"
	errorHandler "gemini-care/error"
	"gemini-care/helper"
	"gemini-care/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type sysmtomHandler struct {
	usecase usecase.SymptomUsecase
}

func NewSymptomHandler(usecase usecase.SymptomUsecase) *sysmtomHandler {
	return &sysmtomHandler{usecase: usecase}
}

func (s *sysmtomHandler) SymptomChecker(ctx echo.Context) error {
	var request dto.SymptomsRequest

	if err := ctx.Bind(&request); err != nil {
		return errorHandler.HandleError(ctx, &errorHandler.BadRequestError{Message: err.Error()})
	}

	response, err := s.usecase.SymptomChecker(&request)
	if err != nil {
		return errorHandler.HandleError(ctx, err)
	}
	fmt.Println(response)
	return ctx.JSON(http.StatusOK, helper.Response(dto.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "Berhasil mendapatkan jawaban",
		Data:       response,
	}))
}
