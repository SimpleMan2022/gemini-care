package errorHandler

import (
	"errors"
	"gemini-care/dto"
	"gemini-care/helper"
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandleError(ctx echo.Context, err error) error {
	var statusCode int
	var badRequestError *BadRequestError
	var notFoundError *NotFoundError
	var internalServerError *InternalServerError
	switch {
	case errors.As(err, &badRequestError):
		statusCode = http.StatusBadRequest
	case errors.As(err, &notFoundError):
		statusCode = http.StatusNotFound
	case errors.As(err, &internalServerError):
		statusCode = http.StatusInternalServerError
	default:
		statusCode = http.StatusInternalServerError
	}

	response := helper.Response(dto.ResponseParam{
		StatusCode: statusCode,
		Message:    err.Error(),
	})

	return ctx.JSON(statusCode, response)
}
