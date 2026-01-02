package security

import (
	"fmt"
	"net/http"

	"github.com/niteshkumarsinha/utho-sdk-go/internal/client"
)

// SecurityService handles communication with the security (SSH Keys) related methods of the Utho API.
type SecurityService struct {
	client *client.Client
}

// NewService creates a new SecurityService.
func NewService(client *client.Client) *SecurityService {
	return &SecurityService{
		client: client,
	}
}

// SSHKey represents a Utho SSH key.
type SSHKey struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	PublicKey string `json:"public_key"`
}

// ListSSHKeysResponse represents the response for listing SSH keys.
type ListSSHKeysResponse struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Data    []SSHKey `json:"data"`
}

// ListSSHKeys returns a list of all SSH keys.
func (s *SecurityService) ListSSHKeys() ([]SSHKey, error) {
	var resp ListSSHKeysResponse
	err := s.client.Request(http.MethodGet, "/key", nil, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return resp.Data, nil
}

// ImportSSHKeyParams represents the parameters for importing an SSH key.
type ImportSSHKeyParams struct {
	Name      string `json:"name"`
	PublicKey string `json:"public_key"`
}

// ImportSSHKeyResponse represents the response for importing an SSH key.
type ImportSSHKeyResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	KeyID   string `json:"keyid"`
}

// ImportSSHKey imports a new SSH key.
func (s *SecurityService) ImportSSHKey(params ImportSSHKeyParams) (*ImportSSHKeyResponse, error) {
	var resp ImportSSHKeyResponse
	err := s.client.Request(http.MethodPost, "/key/import", params, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return &resp, nil
}

// APIKey represents a Utho API key.
type APIKey struct {
	ID        string `json:"id"`
	Label     string `json:"label"`
	APIKey    string `json:"apikey"`
	CreatedAt string `json:"created_at"`
}

// ListAPIKeysResponse represents the response for listing API keys.
type ListAPIKeysResponse struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Data    []APIKey `json:"data"`
}

// ListAPIKeys returns a list of all API keys.
func (s *SecurityService) ListAPIKeys() ([]APIKey, error) {
	var resp ListAPIKeysResponse
	err := s.client.Request(http.MethodGet, "/api", nil, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return resp.Data, nil
}

// DeleteSSHKey deletes an SSH key.
func (s *SecurityService) DeleteSSHKey(id string) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	url := fmt.Sprintf("/key/%s/delete", id)
	err := s.client.Request(http.MethodDelete, url, nil, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}

// GenerateAPIKeyParams represents the parameters for generating an API key.
type GenerateAPIKeyParams struct {
	Label string `json:"label"`
}

// GenerateAPIKey generates a new API key.
func (s *SecurityService) GenerateAPIKey(params GenerateAPIKeyParams) (*APIKey, error) {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Data    APIKey `json:"data"`
	}
	err := s.client.Request(http.MethodPost, "/api/generate", params, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return &resp.Data, nil
}

// DeleteAPIKey deletes an API key.
func (s *SecurityService) DeleteAPIKey(id string) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	url := fmt.Sprintf("/api/%s/delete", id)
	err := s.client.Request(http.MethodDelete, url, nil, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}
