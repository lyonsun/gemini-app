package demo

import (
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func TextGeneration() {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text("What will the world be like in 2025"))
	if err != nil {
		log.Fatal(err)
	}

	printResponse(resp)
}
