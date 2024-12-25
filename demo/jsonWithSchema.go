package demo

import (
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func JsonWithSchema() {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-pro-latest")
	// Ask the model to respond with JSON.
	model.ResponseMIMEType = "application/json"
	prompt := `List the latest movies in 2024 using this JSON schema:
                   Movie = {'movieName': string}
	           Return: Array<Movie>`
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}

	printResponse(resp)
}
