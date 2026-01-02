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

// CreateParams represents the parameters for creating a WAF.
type CreateParams struct {
	Name   string `json:"name"`
	DCSlug string `json:"dcslug"`
}

// Create creates a new WAF instance.
func (s *WafService) Create(params CreateParams) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	err := s.client.Request(http.MethodPost, "/waf/create", params, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}

// Delete destroys a WAF instance.
func (s *WafService) Delete(id string) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	url := fmt.Sprintf("/waf/%s/delete", id)
	err := s.client.Request(http.MethodDelete, url, nil, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}

// AttachParams represents parameters for attaching WAF to a resource.
type AttachParams struct {
	ResourceID   string `json:"resource_id"`
	ResourceType string `json:"resource_type"`
}

// Attach attaches a WAF to a resource (e.g. Load Balancer).
func (s *WafService) Attach(wafID string, params AttachParams) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	url := fmt.Sprintf("/waf/%s/attach", wafID)
	err := s.client.Request(http.MethodPost, url, params, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}

// Detach detaches a WAF from a resource.
func (s *WafService) Detach(wafID string) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	url := fmt.Sprintf("/waf/%s/detach", wafID)
	err := s.client.Request(http.MethodPost, url, nil, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}
