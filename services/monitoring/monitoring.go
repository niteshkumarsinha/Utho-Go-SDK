package monitoring

import (
	"fmt"
	"net/http"

	"github.com/niteshkumarsinha/utho-sdk-go/internal/client"
)

// MonitoringService handles communication with the monitoring and alert related methods of the Utho API.
type MonitoringService struct {
	client *client.Client
}

// NewService creates a new MonitoringService.
func NewService(client *client.Client) *MonitoringService {
	return &MonitoringService{
		client: client,
	}
}

// AlertPolicy represents a monitoring alert policy.
type AlertPolicy struct {
	ID       string `json:"id"`
	Label    string `json:"label"`
	Active   bool   `json:"active"`
	Resource string `json:"resource_type"`
}

// ListAlertPoliciesResponse represents the response for listing alert policies.
type ListAlertPoliciesResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    []AlertPolicy `json:"data"`
}

// ListAlertPolicies returns a list of all monitoring alert policies.
func (s *MonitoringService) ListAlertPolicies() ([]AlertPolicy, error) {
	var resp ListAlertPoliciesResponse
	err := s.client.Request(http.MethodGet, "/monitoring/alerts", nil, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return resp.Data, nil
}

// CreateAlertPolicyParams represents the parameters for creating an alert policy.
type CreateAlertPolicyParams struct {
	Label        string   `json:"label"`
	ResourceType string   `json:"resource_type"`
	Contacts     []string `json:"contacts"`
	Thresholds   struct {
		CPU       int `json:"cpu,omitempty"`
		RAM       int `json:"ram,omitempty"`
		Disk      int `json:"disk,omitempty"`
		Bandwidth int `json:"bandwidth,omitempty"`
	} `json:"thresholds"`
}

// CreateAlertPolicy creates a new monitoring alert policy.
func (s *MonitoringService) CreateAlertPolicy(params CreateAlertPolicyParams) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	err := s.client.Request(http.MethodPost, "/monitoring/alerts/create", params, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}

// DeleteAlertPolicy deletes a monitoring alert policy.
func (s *MonitoringService) DeleteAlertPolicy(id string) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	url := fmt.Sprintf("/monitoring/alerts/%s/delete", id)
	err := s.client.Request(http.MethodDelete, url, nil, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}
