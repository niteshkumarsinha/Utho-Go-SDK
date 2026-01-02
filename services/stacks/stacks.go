package stacks

import (
	"fmt"
	"net/http"

	"github.com/niteshkumarsinha/utho-sdk-go/internal/client"
)

// StacksService handles communication with the stack related methods of the Utho API.
type StacksService struct {
	client *client.Client
}

// NewService creates a new StacksService.
func NewService(client *client.Client) *StacksService {
	return &StacksService{
		client: client,
	}
}

// Stack represents a Utho automation stack.
type Stack struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Images      string `json:"images"`
	Status      string `json:"status"`
	Script      string `json:"script"`
}

// ListStacksResponse represents the response for listing stacks.
type ListStacksResponse struct {
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Data    []Stack `json:"data"`
}

// List returns a list of all stacks.
func (s *StacksService) List() ([]Stack, error) {
	var resp ListStacksResponse
	err := s.client.Request(http.MethodGet, "/stacks", nil, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return resp.Data, nil
}

// CreateParams represents the parameters for creating a stack.
type CreateParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Images      string `json:"images"`
	Status      string `json:"status"`
	Script      string `json:"script"`
}

// Create creates a new stack.
func (s *StacksService) Create(params CreateParams) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	err := s.client.Request(http.MethodPost, "/stacks", params, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}

// Delete destroys a stack.
func (s *StacksService) Delete(id string) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	url := fmt.Sprintf("/stacks/%s", id)
	err := s.client.Request(http.MethodDelete, url, nil, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}

// UpdateParams represents parameters for updating a stack.
type UpdateParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Images      string `json:"images"`
	Status      string `json:"status"`
	Script      string `json:"script"`
}

// Update updates a stack configuration.
func (s *StacksService) Update(id string, params UpdateParams) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	url := fmt.Sprintf("/stacks/%s", id)
	err := s.client.Request(http.MethodPut, url, params, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}
