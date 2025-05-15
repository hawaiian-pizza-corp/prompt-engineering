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
	You are an expert with iris species.
	# Instructions for Iris Species Classification

	As an LLM tasked with iris species classification, 
	you must follow these steps to analyze the four key measurements and determine the species 
	(Setosa, Versicolor, or Verginica).

	## Input Features
	You will be given four numerical measurements:
	1. Petal width (in cm)
	2. Petal length (in cm)
	3. Sepal width (in cm)
	4. Sepal length (in cm)
	`
	//? Context, examples, ...
	classificationProcess := `
	## Classification Process

	### Step 1: Primary Feature Analysis
	First, examine the petal measurements as they are the most discriminative features:
	- Setosa has distinctively small petals
	- Petal length < 2 cm
	- Petal width < 0.5 cm

	### Step 2: Secondary Feature Analysis
	If the specimen is not clearly Setosa, analyze the combination of features:

	For Versicolor:
	- Petal length typically between 3-5 cm
	- Petal width between 1.0-1.8 cm
	- Sepal length typically between 5-7 cm
	- Sepal width typically between 2-3.5 cm

	For Verginica:
	- Petal length typically > 4.5 cm
	- Petal width typically > 1.4 cm
	- Sepal length typically > 6 cm
	- Sepal width typically between 2.5-3.8 cm

	### Step 3: Decision Making
	1. If petal measurements match Setosa's distinctive small size → Classify as Setosa
	2. If measurements fall in the intermediate range → Classify as Versicolor
	3. If measurements show larger values, especially in petal length → Classify as Verginica

	### Step 4: Confidence Check
	- Consider the clarity of the distinction:
	- Are the measurements clearly in one category's range?
	- Are there any overlapping characteristics?
	- Express any uncertainty if measurements are in borderline ranges

	### Step 5: Explanation
	Provide reasoning for your classification by:
	1. Highlighting which measurements were most decisive
	2. Explaining why certain features led to your conclusion
	3. Noting any unusual or borderline measurements

	## Example Reasoning
	"Given a specimen with:
	- Petal width: 0.2 cm
	- Petal length: 1.4 cm
	- Sepal width: 3.5 cm
	- Sepal length: 5.1 cm

	Classification process:
	1. The very small petal measurements (width 0.2 cm, length 1.4 cm) are highly characteristic of Setosa
	2. These petal dimensions are well below the ranges for Versicolor and Verginica
	3. The sepal measurements support this classification, being in the typical range for Setosa
	4. Confidence is high due to the distinctive petal size

	Therefore, this specimen is classified as Setosa with high confidence."
	`

	irisDatabase := `
	## Iris Species Database
	<irisDataset>
		<metadata>
			<attributes>
				<attribute>Species_No</attribute>
				<attribute>Petal_width</attribute>
				<attribute>Petal_length</attribute>
				<attribute>Sepal_width</attribute>
				<attribute>Sepal_length</attribute>
				<attribute>Species_name</attribute>
			</attributes>
		</metadata>
		<records>
			<iris>
				<Species_No>1</Species_No>
				<Petal_width>0,2</Petal_width>
				<Petal_length>1,4</Petal_length>
				<Sepal_width>3,5</Sepal_width>
				<Sepal_length>5,1</Sepal_length>
				<Species_name>Setosa</Species_name>
			</iris>
			<iris>
				<Species_No>1</Species_No>
				<Petal_width>0,2</Petal_width>
				<Petal_length>1,4</Petal_length>
				<Sepal_width>3</Sepal_width>
				<Sepal_length>4,9</Sepal_length>
				<Species_name>Setosa</Species_name>
			</iris>
			<iris>
				<Species_No>1</Species_No>
				<Petal_width>0,2</Petal_width>
				<Petal_length>1,3</Petal_length>
				<Sepal_width>3,2</Sepal_width>
				<Sepal_length>4,7</Sepal_length>
				<Species_name>Setosa</Species_name>
			</iris>
			<iris>
				<Species_No>1</Species_No>
				<Petal_width>0,2</Petal_width>
				<Petal_length>1,5</Petal_length>
				<Sepal_width>3,1</Sepal_width>
				<Sepal_length>4,6</Sepal_length>
				<Species_name>Setosa</Species_name>
			</iris>
			<iris>
				<Species_No>1</Species_No>
				<Petal_width>0,2</Petal_width>
				<Petal_length>1,4</Petal_length>
				<Sepal_width>3,6</Sepal_width>
				<Sepal_length>5</Sepal_length>
				<Species_name>Setosa</Species_name>
			</iris>
			<iris>
				<Species_No>1</Species_No>
				<Petal_width>0,4</Petal_width>
				<Petal_length>1,7</Petal_length>
				<Sepal_width>3,9</Sepal_width>
				<Sepal_length>5,4</Sepal_length>
				<Species_name>Setosa</Species_name>
			</iris>
			<iris>
				<Species_No>1</Species_No>
				<Petal_width>0,3</Petal_width>
				<Petal_length>1,4</Petal_length>
				<Sepal_width>3,4</Sepal_width>
				<Sepal_length>4,6</Sepal_length>
				<Species_name>Setosa</Species_name>
			</iris>
			<iris>
				<Species_No>1</Species_No>
				<Petal_width>0,2</Petal_width>
				<Petal_length>1,5</Petal_length>
				<Sepal_width>3,4</Sepal_width>
				<Sepal_length>5</Sepal_length>
				<Species_name>Setosa</Species_name>
			</iris>
			<iris>
				<Species_No>1</Species_No>
				<Petal_width>0,2</Petal_width>
				<Petal_length>1,4</Petal_length>
				<Sepal_width>2,9</Sepal_width>
				<Sepal_length>4,4</Sepal_length>
				<Species_name>Setosa</Species_name>
			</iris>
			<iris>
				<Species_No>1</Species_No>
				<Petal_width>0,1</Petal_width>
				<Petal_length>1,5</Petal_length>
				<Sepal_width>3,1</Sepal_width>
				<Sepal_length>4,9</Sepal_length>
				<Species_name>Setosa</Species_name>
			</iris>
			<iris>
				<Species_No>2</Species_No>
				<Petal_width>1,4</Petal_width>
				<Petal_length>4,7</Petal_length>
				<Sepal_width>3,2</Sepal_width>
				<Sepal_length>7</Sepal_length>
				<Species_name>Versicolor</Species_name>
			</iris>
			<iris>
				<Species_No>2</Species_No>
				<Petal_width>1,5</Petal_width>
				<Petal_length>4,5</Petal_length>
				<Sepal_width>3,2</Sepal_width>
				<Sepal_length>6,4</Sepal_length>
				<Species_name>Versicolor</Species_name>
			</iris>
			<iris>
				<Species_No>3</Species_No>
				<Petal_width>2,5</Petal_width>
				<Petal_length>6</Petal_length>
				<Sepal_width>3,3</Sepal_width>
				<Sepal_length>6,3</Sepal_length>
				<Species_name>Verginica</Species_name>
			</iris>
			<iris>
				<Species_No>3</Species_No>
				<Petal_width>1,9</Petal_width>
				<Petal_length>5,1</Petal_length>
				<Sepal_width>2,7</Sepal_width>
				<Sepal_length>5,8</Sepal_length>
				<Species_name>Verginica</Species_name>
			</iris>
			<iris>
				<Species_No>3</Species_No>
				<Petal_width>2,1</Petal_width>
				<Petal_length>5,9</Petal_length>
				<Sepal_width>3</Sepal_width>
				<Sepal_length>7,1</Sepal_length>
				<Species_name>Verginica</Species_name>
			</iris>
		</records>
	</irisDataset>			
	`

/* 	userQuestion := `
	Using the below information, 
	Given a specimen with:
	- Petal width: 2,5 cm
	- Petal length: 6 cm
	- Sepal width: 3,3 cm
	- Sepal length: 6,3 cm
	What is the species of the iris?
	` */
	//* Iris Virginica
	
	
	userQuestion := `
	Using the below information, 
	Given a specimen with:
	- Petal width: 1,5 cm
	- Petal length: 4,5 cm
	- Sepal width: 3,2 cm
	- Sepal length: 6,4 cm
	What is the species of the iris?
	`
	//* Iris Versicolor

	messages := []openai.ChatCompletionMessageParamUnion{
		openai.SystemMessage(systemInstructions),
		openai.SystemMessage(classificationProcess),
		openai.SystemMessage(irisDatabase),
		openai.UserMessage("/think " + userQuestion),
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
