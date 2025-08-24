package ollama

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	baseURL string
	client  *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

func (c *Client) Generate(model, prompt string) (string, error) {
	reqData := Request{
		Model:  model,
		Prompt: prompt,
		Stream: false,
	}

	jsonData, err := json.Marshal(reqData)
	if err != nil {
		return "", fmt.Errorf("JSON 마샬링 오류: %v", err)
	}

	resp, err := c.client.Post(c.baseURL+"/api/generate", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("Ollama 서버 연결 오류: %v (Ollama가 실행 중인지 확인하세요)", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP 오류: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("응답 읽기 오류: %v", err)
	}

	var ollamaResp Response
	if err := json.Unmarshal(body, &ollamaResp); err != nil {
		return "", fmt.Errorf("JSON 언마샬링 오류: %v", err)
	}

	return strings.TrimSpace(ollamaResp.Response), nil
}