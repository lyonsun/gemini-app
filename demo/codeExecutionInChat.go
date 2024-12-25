package demo

import (
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func CodeExecutionInChat() {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-pro")
	// To enable code execution, set the `CodeExecution` tool.
	model.Tools = []*genai.Tool{
		{CodeExecution: &genai.CodeExecution{}},
	}

	cs := model.StartChat()
	res, err := cs.SendMessage(ctx, genai.Text(`
  What is the sum of the first 50 prime numbers?
  Generate and run code for the calculation, and make sure you get all 50.
`))
	if err != nil {
		log.Fatal(err)
	}

	// do something with `res`
	printResponse(res)
}
