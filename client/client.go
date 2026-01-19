// Package client provides the low-level HTTP client for communicating with the Utho API.
// It handles authentication, request marshaling, and response decoding.
package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	// DefaultBaseURL is the default endpoint for the Utho API v2.
	DefaultBaseURL = "https://api.utho.com/v2"
	// DefaultTimeout is the default timeout for HTTP requests to the API.
	DefaultTimeout = 30 * time.Second
)

// Config holds the configuration for the Utho SDK client.
type Config struct {
	// BaseURL is the Utho API version endpoint (default: https://api.utho.com/v2).
	BaseURL string
	// APIKey is your Utho API key used for authentication.
	APIKey string
	// HTTPClient is an optional custom *http.Client.
	HTTPClient *http.Client
}

// Client handles the low-level HTTP communication with the Utho API.
// It can be used directly or passed to individual service constructors.
type Client struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

// New creates a new Client with the provided API key and default configuration.
// It returns a pointer to the Client and an error if initialization fails.
func New(apiKey string) (*Client, error) {
	return NewWithConfig(Config{
		APIKey: apiKey,
	})
}

// NewWithConfig creates a new Client using the provided custom configuration.
// If BaseURL or HTTPClient are not provided, default values are used.
func NewWithConfig(cfg Config) (*Client, error) {
	if cfg.BaseURL == "" {
		cfg.BaseURL = DefaultBaseURL
	}
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = &http.Client{
			Timeout: DefaultTimeout,
		}
	}

	return &Client{
		BaseURL:    cfg.BaseURL,
		APIKey:     cfg.APIKey,
		HTTPClient: cfg.HTTPClient,
	}, nil
}

// Request sends an HTTP request to the Utho API.
func (c *Client) Request(method, path string, body interface{}, out interface{}) error {
	var bodyReader io.Reader
	if body != nil {
		if r, ok := body.(io.Reader); ok {
			bodyReader = r
		} else {
			jsonData, err := json.Marshal(body)
			if err != nil {
				return fmt.Errorf("error marshaling request body: %w", err)
			}
			bodyReader = bytes.NewReader(jsonData)
		}
	}

	url := c.BaseURL + path
	if !strings.HasPrefix(path, "/") {
		url = c.BaseURL + "/" + path
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	if _, ok := body.(io.Reader); !ok {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(respBody))
	}

	if out != nil {
		if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
			return fmt.Errorf("error decoding response: %w", err)
		}
	}

	return nil
}
