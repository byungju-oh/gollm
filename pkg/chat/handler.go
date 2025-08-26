package chat

import (
	"fmt"
	"gollm/config"
	"gollm/pkg/ollama"
)

type Handler struct {
	client *ollama.Client
	config *config.Config
}

func NewHandler(client *ollama.Client, config *config.Config) *Handler {
	return &Handler{
		client: client,
		config: config,
	}
}

func (h *Handler) ProcessMessage(message string) (string, error) {
	response, err := h.client.Generate(h.config.DefaultModel, message)
	if err != nil {
		return "", fmt.Errorf("메시지 처리 중 오류: %v", err)
	}
	return response, nil
}

func (h *Handler) IsExitCommand(input string) bool {
	return input == "quit" || input == "exit"
}
