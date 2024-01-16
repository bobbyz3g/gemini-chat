package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	key := os.Getenv("API_KEY")
	if key == "" {
		log.Fatal("missing env API_KEY")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	model := client.GenerativeModel("gemini-pro")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("You: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "" {
			continue
		}
		if text == "exit()" {
			break
		}
		fmt.Println("Gemini:")
		resp, err := model.GenerateContent(ctx, genai.Text(text))
		if err != nil {
			log.Fatal(err)
		}
		for _, candidate := range resp.Candidates {
			for _, part := range candidate.Content.Parts {
				fmt.Print(part)
			}
		}
		fmt.Println()
	}
}
