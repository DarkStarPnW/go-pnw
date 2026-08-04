// Package pnw provides a client for the Politics and War GraphQL API (v3).
// API endpoint: https://api.politicsandwar.com/graphql
// Get your API key at: https://politicsandwar.com/account/
package pnw

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const defaultEndpoint = "https://api.politicsandwar.com/graphql"

// Client is the Politics and War API client.
type Client struct {
	apiKey     string
	botKey     string
	endpoint   string
	httpClient *http.Client
}

// Option configures the Client.
type Option func(*Client)

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) { c.httpClient = hc }
}

// WithBotKey sets the verified bot key required by mutations such as
// BankWithdraw and BankDeposit. Queries do not need it.
func WithBotKey(key string) Option {
	return func(c *Client) { c.botKey = key }
}

// WithEndpoint overrides the default API endpoint.
func WithEndpoint(url string) Option {
	return func(c *Client) { c.endpoint = url }
}

// NewClient creates a new API client using the given API key.
func NewClient(apiKey string, opts ...Option) *Client {
	c := &Client{
		apiKey:   apiKey,
		endpoint: defaultEndpoint,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
	for _, o := range opts {
		o(c)
	}
	return c
}

type graphQLRequest struct {
	Query     string         `json:"query"`
	Variables map[string]any `json:"variables,omitempty"`
}

type graphQLResponse[T any] struct {
	Data   T              `json:"data"`
	Errors []graphQLError `json:"errors"`
}

type graphQLError struct {
	Message    string `json:"message"`
	Extensions struct {
		Category string `json:"category"`
	} `json:"extensions"`
}

func (e graphQLError) Error() string { return e.Message }

// do executes a raw GraphQL query and decodes the response into out.
func (c *Client) do(ctx context.Context, query string, variables map[string]any, out any) error {
	body, err := json.Marshal(graphQLRequest{Query: query, Variables: variables})
	if err != nil {
		return fmt.Errorf("pnw: marshal request: %w", err)
	}

	url := c.endpoint + "?api_key=" + c.apiKey
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("pnw: build request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	// Mutations are authenticated by the bot key pair rather than the query string.
	if c.botKey != "" {
		req.Header.Set("X-Bot-Key", c.botKey)
		req.Header.Set("X-Api-Key", c.apiKey)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("pnw: http: %w", err)
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("pnw: read body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("pnw: http %d: %s", resp.StatusCode, raw)
	}

	var wrapper struct {
		Data   json.RawMessage `json:"data"`
		Errors []graphQLError  `json:"errors"`
	}
	if err := json.Unmarshal(raw, &wrapper); err != nil {
		return fmt.Errorf("pnw: decode response: %w", err)
	}
	if len(wrapper.Errors) > 0 {
		return APIError(wrapper.Errors)
	}
	if out != nil && wrapper.Data != nil {
		if err := json.Unmarshal(wrapper.Data, out); err != nil {
			return fmt.Errorf("pnw: decode data: %w", err)
		}
	}
	return nil
}
