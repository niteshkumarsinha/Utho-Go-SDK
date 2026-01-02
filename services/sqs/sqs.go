package sqs

import (
	"fmt"
	"net/http"

	"github.com/niteshkumarsinha/utho-sdk-go/internal/client"
)

// SqsService handles communication with the SQS related methods of the Utho API.
type SqsService struct {
	client *client.Client
}

// NewService creates a new SqsService.
func NewService(client *client.Client) *SqsService {
	return &SqsService{
		client: client,
	}
}

// SQSInstance represents an SQS instance.
type SQSInstance struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

// ListSqsResponse represents the response for listing SQS instances.
type ListSqsResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    []SQSInstance `json:"data"`
}

// List returns a list of all SQS instances.
func (s *SqsService) List() ([]SQSInstance, error) {
	var resp ListSqsResponse
	err := s.client.Request(http.MethodGet, "/sqs", nil, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return resp.Data, nil
}

// CreateParams represents the parameters for creating an SQS instance.
type CreateParams struct {
	Name   string `json:"name"`
	DCSlug string `json:"dcslug"`
}

// Create creates a new SQS instance.
func (s *SqsService) Create(params CreateParams) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	err := s.client.Request(http.MethodPost, "/sqs/create", params, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}

// Delete destroys an SQS instance.
func (s *SqsService) Delete(id string) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	url := fmt.Sprintf("/sqs/%s/delete", id)
	err := s.client.Request(http.MethodDelete, url, nil, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}
