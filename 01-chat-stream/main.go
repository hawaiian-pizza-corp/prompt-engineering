package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func main() {
	ctx := context.Background()

	llmURL := os.Getenv("MODEL_RUNNER_BASE_URL") + "/engines/llama.cpp/v1/"
	model := os.Getenv("MODEL_RUNNER_LLM_CHAT")

	client := openai.NewClient(
		option.WithBaseURL(llmURL),
		option.WithAPIKey(""),
	)

	//! Prompt
	systemInstructions := `
	You are a Hawaiian pizza expert. Your name is Bob.
	Provide accurate, enthusiastic information about Hawaiian pizza's history 
	(invented in Canada in 1962 by Sam Panopoulos), 
	ingredients (ham, pineapple, cheese on tomato sauce), preparation methods, and cultural impact.
	Use a friendly tone with occasional pizza puns. 
	Defend pineapple on pizza good-naturedly while respecting differing opinions. 
	If asked about other pizzas, briefly answer but return focus to Hawaiian pizza. 
	Emphasize the sweet-savory flavor combination that makes Hawaiian pizza special.
	USE ONLY THE INFORMATION PROVIDED IN THE KNOWLEDGE BASE.	
	`
	
	userQuestion := `What is the best pizza in the world?`
	//userQuestion := `What are the ingredients of the hawaiian pizza?`
	//userQuestion := `What is your name?`

	messages := []openai.ChatCompletionMessageParamUnion{
		openai.SystemMessage(systemInstructions),
		openai.UserMessage(userQuestion),
	}

	//! Request parameters

	param := openai.ChatCompletionNewParams{
		Messages:    messages,
		Model:       model,
		Temperature: openai.Opt(0.0),
	}

	StreamCompletion(ctx, &client, param)


}


func StreamCompletion(ctx context.Context, client *openai.Client, param openai.ChatCompletionNewParams) {
	fmt.Println()

	stream := client.Chat.Completions.NewStreaming(ctx, param)

	for stream.Next() {
		chunk := stream.Current()
		if len(chunk.Choices) > 0 && chunk.Choices[0].Delta.Content != "" {
			fmt.Print(chunk.Choices[0].Delta.Content)
		}
	}

	if err := stream.Err(); err != nil {
		log.Fatalln("😡:", err)
	}
	fmt.Println()
	fmt.Println()

}
