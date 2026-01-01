package registry

import (
	"fmt"
	"net/http"

	"github.com/niteshkumarsinha/utho-sdk-go/internal/client"
)

// RegistryService handles communication with the container registry related methods of the Utho API.
type RegistryService struct {
	client *client.Client
}

// NewService creates a new RegistryService.
func NewService(client *client.Client) *RegistryService {
	return &RegistryService{
		client: client,
	}
}

// Registry represents a Utho Container Registry.
type Registry struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	URL    string `json:"url"`
	Status string `json:"status"`
}

// ListRegistriesResponse represents the response for listing registries.
type ListRegistriesResponse struct {
	Status  string     `json:"status"`
	Message string     `json:"message"`
	Data    []Registry `json:"data"`
}

// List returns a list of all container registries.
func (s *RegistryService) List() ([]Registry, error) {
	var resp ListRegistriesResponse
	err := s.client.Request(http.MethodGet, "/registry", nil, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return resp.Data, nil
}
