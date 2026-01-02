package ssl

import (
	"fmt"
	"net/http"

	"github.com/niteshkumarsinha/utho-sdk-go/internal/client"
)

// SslService handles communication with the SSL certificate related methods of the Utho API.
type SslService struct {
	client *client.Client
}

// NewService creates a new SslService.
func NewService(client *client.Client) *SslService {
	return &SslService{
		client: client,
	}
}

// Certificate represents an SSL certificate.
type Certificate struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// ListSslResponse represents the response for listing SSL certificates.
type ListSslResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    []Certificate `json:"data"`
}

// List returns a list of all SSL certificates.
func (s *SslService) List() ([]Certificate, error) {
	var resp ListSslResponse
	err := s.client.Request(http.MethodGet, "/certificates", nil, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return resp.Data, nil
}

// CreateParams represents the parameters for creating (uploading) an SSL certificate.
type CreateParams struct {
	Name        string `json:"name"`
	Certificate string `json:"certificate"`
	PrivateKey  string `json:"private_key"`
	Chain       string `json:"chain"`
}

// Create uploads a new SSL certificate.
func (s *SslService) Create(params CreateParams) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	err := s.client.Request(http.MethodPost, "/certificates", params, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}

// Delete destroys an SSL certificate.
func (s *SslService) Delete(id string) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	url := fmt.Sprintf("/certificates/%s", id)
	err := s.client.Request(http.MethodDelete, url, nil, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}
