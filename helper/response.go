package helper

import "gemini-care/dto"

type ResponseWithoutData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseWithData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func GenerateResponseWithoutData(status string, message string) ResponseWithoutData {
	return ResponseWithoutData{status, message}
}

func GenerateResponseWithData(status string, message string, data any) ResponseWithData {
	return ResponseWithData{status, message, data}
}

func Response(params dto.ResponseParam) any {
	var status string
	if params.StatusCode >= 200 && params.StatusCode < 300 {
		status = "success"
	} else {
		status = "failed"
	}

	if params.Data == nil {
		return GenerateResponseWithoutData(status, params.Message)
	} else {
		return GenerateResponseWithData(status, params.Message, params.Data)
	}
}
