#!/bin/bash
MODEL_RUNNER_BASE_URL=http://model-runner.docker.internal \
MODEL_RUNNER_LLM_CHAT=ai/qwen2.5:0.5B-F16 \
go run main.go
