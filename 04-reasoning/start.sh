#!/bin/bash
MODEL_RUNNER_BASE_URL=http://model-runner.docker.internal \
MODEL_RUNNER_LLM_CHAT=ai/deepseek-r1-distill-llama \
go run main.go

#MODEL_RUNNER_LLM_CHAT=ai/qwen3 \
