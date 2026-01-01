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
