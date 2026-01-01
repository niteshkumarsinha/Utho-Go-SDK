package cloudserver

import (
	"fmt"
	"net/http"

	"github.com/niteshkumarsinha/utho-sdk-go/internal/client"
)

// CloudServerService handles communication with the cloud server related methods of the Utho API.
// This includes listing instances, deploying new servers, and managing power states.
type CloudServerService struct {
	client *client.Client
}

// NewService creates a new CloudServerService.
func NewService(client *client.Client) *CloudServerService {
	return &CloudServerService{
		client: client,
	}
}

// Instance represents a Utho cloud instance.
type Instance struct {
	ID           string `json:"id"`
	Hostname     string `json:"hostname"`
	IP           string `json:"ip"`
	Status       string `json:"status"`
	PlanID       string `json:"planid"`
	DC           string `json:"dc"`
	Image        string `json:"image"`
	CreatedAt    string `json:"created_at"`
	BillingCycle string `json:"billingcycle"`
}

// ListInstancesResponse represents the response for listing instances.
type ListInstancesResponse struct {
	Status  string     `json:"status"`
	Message string     `json:"message"`
	Data    []Instance `json:"data"`
}

// List returns a list of all active cloud servers in your Utho account.
// It returns a slice of Instance structs containing details like IP, Status, and DC.
func (s *CloudServerService) List() ([]Instance, error) {
	var resp ListInstancesResponse
	err := s.client.Request(http.MethodGet, "/cloud/", nil, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return resp.Data, nil
}

// DeployParams represents the parameters for deploying an instance.
type DeployParams struct {
	DCSlug         string           `json:"dcslug"`
	PlanID         string           `json:"planid"`
	BillingCycle   string           `json:"billingcycle"`
	Auth           string           `json:"auth"`
	EnablePublicIP string           `json:"enable_publicip"`
	Image          string           `json:"image"`
	Cloud          []InstanceConfig `json:"cloud"`
}

// InstanceConfig represents the configuration for a single instance in a deploy request.
type InstanceConfig struct {
	Hostname string `json:"hostname"`
}

// DeployResponse represents the response for a deploy request.
type DeployResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	OrderID string `json:"orderid"`
}

// Deploy starts the deployment of one or more cloud instances based on the provided params.
// It returns a DeployResponse containing the OrderID for tracking the deployment progress.
func (s *CloudServerService) Deploy(params DeployParams) (*DeployResponse, error) {
	var resp DeployResponse
	err := s.client.Request(http.MethodPost, "/cloud/deploy/", params, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return &resp, nil
}
