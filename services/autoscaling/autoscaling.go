package autoscaling

import (
	"fmt"
	"net/http"

	"github.com/niteshkumarsinha/utho-sdk-go/internal/client"
)

// AutoscalingService handles communication with the autoscaling related methods of the Utho API.
type AutoscalingService struct {
	client *client.Client
}

// NewService creates a new AutoscalingService.
func NewService(client *client.Client) *AutoscalingService {
	return &AutoscalingService{
		client: client,
	}
}

// ASGroup represents a Utho Autoscaling Group.
type ASGroup struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	MinSize int    `json:"min_size"`
	MaxSize int    `json:"max_size"`
	Status  string `json:"status"`
	DC      string `json:"dc"`
}

// ListASGroupsResponse represents the response for listing autoscaling groups.
type ListASGroupsResponse struct {
	Status  string    `json:"status"`
	Message string    `json:"message"`
	Data    []ASGroup `json:"data"`
}

// List returns a list of all autoscaling groups.
func (s *AutoscalingService) List() ([]ASGroup, error) {
	var resp ListASGroupsResponse
	err := s.client.Request(http.MethodGet, "/autoscaling", nil, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return resp.Data, nil
}

// CreateParams represents the parameters for creating an autoscaling group.
type CreateParams struct {
	Name    string `json:"name"`
	MinSize int    `json:"min_size"`
	MaxSize int    `json:"max_size"`
	Image   string `json:"image"`
	Plan    string `json:"plan"`
	Script  string `json:"script"`
}

// Create creates a new autoscaling group.
func (s *AutoscalingService) Create(params CreateParams) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	err := s.client.Request(http.MethodPost, "/autoscaling/create", params, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}

// Delete destroys an autoscaling group.
func (s *AutoscalingService) Delete(id string) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	url := fmt.Sprintf("/autoscaling/%s/delete", id)
	err := s.client.Request(http.MethodDelete, url, nil, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}

// UpdateParams represents parameters for updating an autoscaling group.
type UpdateParams struct {
	MinSize int `json:"min_size,omitempty"`
	MaxSize int `json:"max_size,omitempty"`
}

// Update updates an autoscaling group configuration.
func (s *AutoscalingService) Update(id string, params UpdateParams) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	url := fmt.Sprintf("/autoscaling/%s/update", id)
	err := s.client.Request(http.MethodPut, url, params, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}
