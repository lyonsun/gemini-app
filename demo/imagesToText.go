package demo

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func ImagesToText() {
	ctx := context.Background()
	// Access your API key as an environment variable
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("models/gemini-1.5-pro")

	imagePath1 := "https://upload.wikimedia.org/wikipedia/commons/thumb/8/87/Palace_of_Westminster_from_the_dome_on_Methodist_Central_Hall.jpg/2560px-Palace_of_Westminster_from_the_dome_on_Methodist_Central_Hall.jpg"
	imageResp1, err := http.Get(imagePath1)
	if err != nil {
		panic(err)
	}
	defer imageResp1.Body.Close()
	imageBytes1, err := io.ReadAll(imageResp1.Body)
	if err != nil {
		panic(err)
	}

	imagePath2 := "https://images.unsplash.com/photo-1732221560326-f6c8cb063e8b"
	imageResp2, err := http.Get(imagePath2)
	if err != nil {
		panic(err)
	}
	defer imageResp2.Body.Close()
	imageBytes2, err := io.ReadAll(imageResp2.Body)
	if err != nil {
		panic(err)
	}

	// Create the request.
	req := []genai.Part{
		genai.ImageData("jpeg", imageBytes1),
		genai.ImageData("jpeg", imageBytes2),
		genai.Text("Generate a list of all the objects contained in both images."),
	}

	// Generate content.
	resp, err := model.GenerateContent(ctx, req...)
	if err != nil {
		panic(err)
	}

	// Handle the response of generated text.
	for _, c := range resp.Candidates {
		if c.Content != nil {
			fmt.Println(*c.Content)
		}
	}
}
