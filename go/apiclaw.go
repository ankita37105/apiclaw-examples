// Package apiclaw is a tiny, dependency-free Go client for APIClaw
// (https://apiclaw.biz), a flat-rate, OpenAI-compatible AI API gateway.
//
// It only wraps POST /v1/chat/completions so you can see exactly what goes
// over the wire. Any OpenAI-compatible Go SDK (for example openai-go or
// go-openai) also works: set its base URL to https://apiclaw.biz/v1.
//
// Maintained by the APIClaw team. APIClaw is an independent gateway and is
// not affiliated with Anthropic, OpenAI or Google.
package apiclaw

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

// DefaultBaseURL is the OpenAI-compatible base URL of APIClaw.
const DefaultBaseURL = "https://apiclaw.biz/v1"

// DefaultModel is the model ID the examples use. List the IDs your plan
// exposes with GET https://apiclaw.biz/v1/models.
const DefaultModel = "apiclaw/claude-sonnet-5.5"

// Message is one chat message.
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatRequest is the subset of the OpenAI chat completions request body
// used by this package.
type ChatRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	MaxTokens   int       `json:"max_tokens,omitempty"`
	Temperature *float64  `json:"temperature,omitempty"`
}

// Usage is the token usage block returned by the API.
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// ChatResponse is the subset of the OpenAI chat completions response used
// by this package.
type ChatResponse struct {
	ID      string `json:"id"`
	Model   string `json:"model"`
	Choices []struct {
		Index        int     `json:"index"`
		Message      Message `json:"message"`
		FinishReason string  `json:"finish_reason"`
	} `json:"choices"`
	Usage Usage `json:"usage"`
}

// Text returns the content of the first choice, or "" if there is none.
func (r *ChatResponse) Text() string {
	if r == nil || len(r.Choices) == 0 {
		return ""
	}
	return r.Choices[0].Message.Content
}

// APIError is returned for non-2xx responses. RetryAfter is set when the
// server sent a Retry-After header in seconds (for example on HTTP 429 when
// the daily request quota is used up).
type APIError struct {
	StatusCode int
	Body       string
	RetryAfter time.Duration
}

func (e *APIError) Error() string {
	return fmt.Sprintf("apiclaw: HTTP %d: %s", e.StatusCode, e.Body)
}

// Client calls an OpenAI-compatible chat completions endpoint.
type Client struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient returns a client for APIClaw. If apiKey is empty it reads the
// APICLAW_API_KEY environment variable.
func NewClient(apiKey string) *Client {
	if apiKey == "" {
		apiKey = os.Getenv("APICLAW_API_KEY")
	}
	return &Client{
		APIKey:     apiKey,
		BaseURL:    DefaultBaseURL,
		HTTPClient: &http.Client{Timeout: 120 * time.Second},
	}
}

// Chat sends one chat completions request.
func (c *Client) Chat(ctx context.Context, req ChatRequest) (*ChatResponse, error) {
	if req.Model == "" {
		req.Model = DefaultModel
	}
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	base := c.BaseURL
	if base == "" {
		base = DefaultBaseURL
	}
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, base+"/chat/completions", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+c.APIKey)

	hc := c.HTTPClient
	if hc == nil {
		hc = http.DefaultClient
	}
	resp, err := hc.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		apiErr := &APIError{StatusCode: resp.StatusCode, Body: string(raw)}
		if s, convErr := strconv.Atoi(resp.Header.Get("Retry-After")); convErr == nil {
			apiErr.RetryAfter = time.Duration(s) * time.Second
		}
		return nil, apiErr
	}
	var out ChatResponse
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, fmt.Errorf("apiclaw: decode response: %w", err)
	}
	return &out, nil
}

// Ask is a convenience wrapper: one user prompt in, the reply text out.
func (c *Client) Ask(ctx context.Context, model, prompt string) (string, error) {
	resp, err := c.Chat(ctx, ChatRequest{
		Model:    model,
		Messages: []Message{{Role: "user", Content: prompt}},
	})
	if err != nil {
		return "", err
	}
	return resp.Text(), nil
}
