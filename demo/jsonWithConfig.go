package demo

import (
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func JsonWithConfig() {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-pro-latest")
	// Ask the model to respond with JSON.
	model.ResponseMIMEType = "application/json"
	// Specify the schema.
	model.ResponseSchema = &genai.Schema{
		Type:  genai.TypeArray,
		Items: &genai.Schema{Type: genai.TypeString},
	}
	resp, err := model.GenerateContent(ctx, genai.Text("List top 5 movies in 2024 using this JSON schema."))
	if err != nil {
		log.Fatal(err)
	}

	printResponse(resp)
	// for _, part := range resp.Candidates[0].Content.Parts {
	// 	if txt, ok := part.(genai.Text); ok {
	// 		var recipes []string
	// 		if err := json.Unmarshal([]byte(txt), &recipes); err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		fmt.Println(recipes)
	// 	}
	// }
}
