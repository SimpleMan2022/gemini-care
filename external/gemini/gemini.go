package gemini

import (
	"context"
	"gemini-care/helper"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"log"
	"os"
)

type GeminiClient interface {
	AnswerQuestion(question string) (string, error)
}

type geminiClient struct {
	client *genai.Client
}

func NewGeminiClient() *geminiClient {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))

	if err != nil {
		log.Fatal("Error creating client: ", err)
	}
	return &geminiClient{client: client}
}

func (g *geminiClient) AnswerQuestion(question string) (string, error) {
	ctx := context.Background()

	model := g.client.GenerativeModel("gemini-1.5-flash")
	model.SystemInstruction = helper.GetSystemInstruction()

	cs := model.StartChat()

	message, err := cs.SendMessage(ctx, genai.Text(question))

	if err != nil {
		return "AI sedang bermasalah", err
	}

	response := helper.ToAIResponse(message)
	return response, nil
}
