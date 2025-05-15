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
	//userQuestion := `What is the best pizza in the world?`
	//userQuestion := `What are the ingredients of the hawaiian pizza?``
	userQuestion := `What is your name?`

	messages := []openai.ChatCompletionMessageParamUnion{
		openai.UserMessage(userQuestion),
	}

	//! Request parameters

	param := openai.ChatCompletionNewParams{
		Messages:    messages,
		Model:       model,
		Temperature: openai.Opt(0.5),
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
