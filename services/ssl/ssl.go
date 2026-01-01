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
