package helper

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type ValidationHelper interface {
	ValidateRequest(req any) any
}

type validationHelper struct{}

func NewValidationHelper() *validationHelper {
	return &validationHelper{}
}

type validation struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Errors  []errorCustom `json:"errors"`
}

type errorCustom struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

func GenerateValidationResponse(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return "REQUIRED"
	case "min":
		return fmt.Sprintf("MIN%sCHARACTER", fieldError.Param())
	case "max":
		return fmt.Sprintf("MAX%sCHARACTER", fieldError.Param())
	}
	return fieldError.Error()
}

func (v *validationHelper) ValidateRequest(req any) any {
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			validationsAPI := make([]errorCustom, len(ve))
			for i, fieldError := range ve {
				validationsAPI[i] = errorCustom{
					Field: fieldError.Field(),
					Error: GenerateValidationResponse(fieldError),
				}
			}
			return validation{
				Status:  "Failed",
				Message: "Please check your input",
				Errors:  validationsAPI,
			}
		}
	}
	return nil
}
