package iso

import (
	"fmt"
	"net/http"

	"github.com/niteshkumarsinha/utho-sdk-go/internal/client"
)

// IsoService handles communication with the ISO related methods of the Utho API.
type IsoService struct {
	client *client.Client
}

// NewService creates a new IsoService.
func NewService(client *client.Client) *IsoService {
	return &IsoService{
		client: client,
	}
}

// ISO represents an ISO image.
type ISO struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Size   string `json:"size"`
	Status string `json:"status"`
}

// ListIsoResponse represents the response for listing ISOs.
type ListIsoResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    []ISO  `json:"data"`
}

// List returns a list of all custom ISOs.
func (s *IsoService) List() ([]ISO, error) {
	var resp ListIsoResponse
	err := s.client.Request(http.MethodGet, "/iso", nil, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return resp.Data, nil
}

// CreateParams represents the parameters for adding a custom ISO.
type CreateParams struct {
	DCSlug string `json:"dcslug"`
	Name   string `json:"name"`
	URL    string `json:"url"`
}

// Create adds a new custom ISO from a remote URL.
func (s *IsoService) Create(params CreateParams) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	err := s.client.Request(http.MethodPost, "/iso/add", params, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}

// Delete removes a custom ISO.
func (s *IsoService) Delete(id string) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	url := fmt.Sprintf("/iso/%s/delete", id)
	err := s.client.Request(http.MethodDelete, url, nil, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}
