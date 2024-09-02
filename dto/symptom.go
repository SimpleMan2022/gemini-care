package dto

type SymptomsRequest struct {
	Symptoms []string `json:"symptoms"`
}

type SymptomResponse struct {
	Diagnosis  string `json:"diagnosis"`
	Confidence string `json:"confidence"`
	Tips       string `json:"tips"`
}
