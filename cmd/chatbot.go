package cmd

import (
	"bufio"
	"fmt"
	"gollm/pkg/chat"
	"os"
	"strings"
)

type Chatbot struct {
	handler *chat.Handler
	scanner *bufio.Scanner
}

func NewChatbot(handler *chat.Handler) *Chatbot {
	return &Chatbot{
		handler: handler,
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (c *Chatbot) Start() {
	c.printWelcome()

	for {
		fmt.Print("사용자: ")
		if !c.scanner.Scan() {
			break
		}

		input := strings.TrimSpace(c.scanner.Text())

		if input == "" {
			continue
		}

		if c.handler.IsExitCommand(input) {
			c.printGoodbye()
			break
		}

		response, err := c.handler.ProcessMessage(input)
		if err != nil {
			fmt.Printf("오류 발생: %v\n", err)
			continue
		}

		fmt.Printf("봇: %s\n\n", response)
	}
}

func (c *Chatbot) printWelcome() {
	fmt.Println("=== Ollama 챗봇 ===")
	fmt.Println("종료하려면 'quit' 또는 'exit'를 입력하세요.")
	fmt.Println()
}

func (c *Chatbot) printGoodbye() {
	fmt.Println("챗봇을 종료합니다.")
}
