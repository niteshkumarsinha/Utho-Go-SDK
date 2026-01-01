package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Client handles the low-level HTTP communication with the Utho API.
type Client struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

// Request sends an HTTP request to the Utho API.
func (c *Client) Request(method, path string, body interface{}, out interface{}) error {
	var bodyReader io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("error marshaling request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonData)
	}

	url := c.BaseURL + path
	if !strings.HasPrefix(path, "/") {
		url = c.BaseURL + "/" + path
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
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
