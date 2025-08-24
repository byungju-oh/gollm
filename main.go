package main

import (
	"gollm/cmd"
	"gollm/config"
	"gollm/pkg/chat"
	"gollm/pkg/ollama"
)

func main() {
	cfg := config.New()
	
	client := ollama.NewClient(cfg.OllamaURL)
	
	handler := chat.NewHandler(client, cfg)
	
	chatbot := cmd.NewChatbot(handler)
	
	chatbot.Start()
}