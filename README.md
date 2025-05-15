# Prompt Engineering
> By example

## Intro - prompt engineering?

Prompt engineering is the art and science of formulating clear and effective instructions for artificial intelligence models to obtain the most relevant and useful responses possible. It's like learning to communicate effectively with a highly capable assistant who needs precise directions.

According to OpenAI, prompt engineering is "the process of designing and optimizing input prompts to effectively guide a language model's responses." While a prompt can be a simple text input, the way you structure and refine it dramatically impacts the quality of the AI's output.

## Personality and Ststem instructions
> It's important especially with local models

> - The "Persona" Technique: ask the AI to adopt a specific role to get tailored responses.
> - The "System Message" Technique: use system messages to set overall behavior and guidelines

- `00-chat-stream`
- `01-chat-stream`
- `02-chat-stream`

## Reasoning

> - The "Chain of Thought" Technique: encourage the AI to break down a complex problem into reasoning steps.
>  - "Explain step by step..."
> - The "Few-Shot Learning" Technique: provide a few examples of what you expect as a result. (not only for reasoning)
> - The "XML Tagging" Technique: use XML tags to structure prompts and help the LLM (not only for reasoning)

- `03-reasoning`
- `04-reasoning`

## Gordon and I

### Meta prompts

> - magic keywords to change the behavior of the LLM

```text
[For Kids] explain ECI
```

### Structuration

> - The "Structured Output" Technique: specify exactly how you want the information formatted.

slides -> structured
for a french person -> notes

