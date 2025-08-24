# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is "gollm", a Go-based Korean chatbot application that interfaces with Ollama for language model interactions. The application provides a command-line chat interface in Korean.

## Build and Run Commands

```bash
# Build the application
go build -o gollm .

# Run the application
./gollm

# Run directly with go run
go run main.go
```

## Development Commands

```bash
# Format code
go fmt ./...

# Run vet (static analysis)
go vet ./...

# Get dependencies
go mod tidy

# Test (when tests are added)
go test ./...
```

## Architecture

The application follows a clean architecture pattern:

- **main.go**: Application entry point, orchestrates dependency injection
- **cmd/chatbot.go**: CLI interface and user interaction loop
- **config/config.go**: Configuration management with environment variables
- **pkg/chat/handler.go**: Message processing business logic
- **pkg/ollama/**: Ollama API client implementation
  - **client.go**: HTTP client for Ollama API
  - **types.go**: Request/response data structures

## Configuration

The application uses environment variables with defaults:
- `OLLAMA_URL`: Ollama server URL (default: `http://localhost:11434`)
- `OLLAMA_MODEL`: Default model to use (default: `llama2`)

## Key Components

- **Chatbot**: CLI interface handler that manages user input/output
- **Handler**: Business logic layer that processes messages
- **Client**: HTTP client that communicates with Ollama API

The application is designed with Korean language UI and error messages.

## Commit Message Guidelines

When making commits to this repository, follow these conventions:

- **feat**: 새로운 기능 추가
- **fix**: 버그 수정
- **docs**: 문서 변경 (README, CLAUDE.md 등)
- **style**: 코드 포맷팅, 세미콜론 누락 등 (기능 변경 없음)
- **refactor**: 코드 리팩토링 (기능 변경 없음)
- **test**: 테스트 추가 또는 수정
- **chore**: 빌드/패키지 관리, 의존성 업데이트

**형식**: `타입: 한글 설명`

**예시**:
- `feat: 챗봇 메시지 히스토리 기능 추가`
- `fix: Ollama 연결 오류 처리 개선`
- `docs: 설치 가이드 업데이트`
- `refactor: 핸들러 코드 구조 개선`