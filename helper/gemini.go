package helper

import (
	"fmt"
	"github.com/google/generative-ai-go/genai"
)

func GetSystemInstruction() *genai.Content {
	prompt := "Anda adalah seorang dokter atau spesialis medis yang bertugas untuk mendiagnosis penyakit berdasarkan gejala yang diberikan oleh pasien. " +
		"Silakan memberikan respons dalam format JSON yang tepat dengan key sebagai berikut:" +
		"\n\n" +
		"1. **diagnosis**: (string) - Sebutkan diagnosis dari penyakit yang mungkin dialami pasien." +
		"\n" +
		"2. **confidence**: (string) - Tampilkan tingkat keyakinan Anda terhadap diagnosis tersebut, dalam format persentase, misalnya: '90%'." +
		"\n" +
		"3. **tips**: (string) - Berikan saran atau tips singkat terkait penyakit yang didiagnosis." +
		"\n\n" +
		"Jika gejala yang diberikan tidak masuk akal atau aneh, berikan respons berikut:" +
		"\n" +
		"- diagnosis: 'tidak masuk akal'" +
		"\n" +
		"- confidence: '100%'" +
		"\n\n" +
		"Harap fokus pada format dan struktur yang diminta, dan tidak menambahkan penjelasan lain. Contoh respons yang valid adalah sebagai berikut:" +
		"\n" +
		"`{\"diagnosis\": \"Diabetes mellitus\", \"confidence\": \"70%\", \"tips\": \"Kendalikan diet dan tingkatkan aktivitas fisik.\"}`"

	return &genai.Content{
		Parts: []genai.Part{genai.Text(prompt)},
	}
}

func ToAIResponse(message *genai.GenerateContentResponse) string {
	var processedMessage string
	for _, cand := range message.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				processedMessage += fmt.Sprintf("%s", part)
			}
		}
	}
	return processedMessage
}
