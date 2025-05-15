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
	//? Context, examples, ...
	knowledgeBase := `
	## Traditional Ingredients
	- Base: Traditional pizza dough
	- Sauce: Tomato-based pizza sauce
	- Cheese: Mozzarella cheese
	- Key toppings: Ham (or Canadian bacon) and pineapple
	- Optional additional toppings: Bacon, mushrooms, bell peppers, jalapeños

	## Regional Variations
	- Australia: "Hawaiian and bacon" adds extra bacon to the traditional recipe
	- Brazil: "Portuguesa com abacaxi" combines the traditional Portuguese pizza (with ham, onions, hard-boiled eggs, olives) with pineapple
	- Japan: Sometimes includes teriyaki chicken instead of ham
	- Germany: "Hawaii-Toast" is a related open-faced sandwich with ham, pineapple, and cheese
	- Sweden: "Flying Jacob" pizza includes banana, pineapple, curry powder, and chicken

	## Popular Questions and Answers

	**Q: Why would anyone put pineapple on pizza?**
	A: The combination of sweet pineapple with salty ham creates a contrasting flavor profile that many find appealing. 
	The acidity of pineapple also cuts through the richness of cheese and complements the savory elements. It's a classic sweet-and-savory pairing found in many cuisines worldwide.

	**Q: Is Hawaiian pizza really from Hawaii?**
	A: No, Hawaiian pizza was created in Canada in 1962 by Greek-born Sam Panopoulos. 
	The name comes from the brand of canned pineapple used, not its origin.

	**Q: What cheese is best for Hawaiian pizza?**
	A: Traditional mozzarella works well, but some prefer a blend of mozzarella and provolone for more flavor depth. 
	Some gourmet versions use smoked mozzarella to complement the pineapple and ham.

	**Q: Should pineapple be fresh or canned on pizza?**
	A: While both can work, canned pineapple is often preferred because it's pre-cooked, has consistent sweetness, 
	and the canning process neutralizes the bromelain enzyme that can affect cheese. Fresh pineapple should be pre-cooked or at least well-drained.

	**Q: What drink pairs best with Hawaiian pizza?**
	A: Light, crisp beers like lagers or wheat beers work well. For wine, try an off-dry Riesling or rosé. 
	Non-alcoholic options include lemonade or iced tea with a touch of sweetness.
	`

	//userQuestion := `What are the ingredients of the hawaiian pizza?`
	//userQuestion := `What cheese is best for Hawaiian pizza?`
	userQuestion := `What is the Japanese variation of the hawaiian pizza?`

	messages := []openai.ChatCompletionMessageParamUnion{
		openai.SystemMessage(systemInstructions),
		openai.SystemMessage(knowledgeBase),
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
