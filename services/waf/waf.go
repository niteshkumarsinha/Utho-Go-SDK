package waf

import (
	"fmt"
	"net/http"

	"github.com/niteshkumarsinha/utho-sdk-go/internal/client"
)

// WafService handles communication with the WAF related methods of the Utho API.
type WafService struct {
	client *client.Client
}

// NewService creates a new WafService.
func NewService(client *client.Client) *WafService {
	return &WafService{
		client: client,
	}
}

// WafInstance represents a Utho WAF instance.
type WafInstance struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

// ListWafResponse represents the response for listing WAF instances.
type ListWafResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    []WafInstance `json:"data"`
}

// List returns a list of all WAF instances.
func (s *WafService) List() ([]WafInstance, error) {
	var resp ListWafResponse
	err := s.client.Request(http.MethodGet, "/waf", nil, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return resp.Data, nil
}
