package loadbalancer

import (
	"fmt"
	"net/http"

	"github.com/niteshkumarsinha/utho-sdk-go/internal/client"
)

// LoadBalancerService handles communication with the loadbalancer related methods of the Utho API.
type LoadBalancerService struct {
	client *client.Client
}

// NewService creates a new LoadBalancerService.
func NewService(client *client.Client) *LoadBalancerService {
	return &LoadBalancerService{
		client: client,
	}
}

// LoadBalancer represents a Utho Load Balancer.
type LoadBalancer struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Status string `json:"status"`
	IP     string `json:"ip"`
	DC     string `json:"dc"`
}

// ListLBResponse represents the response for listing load balancers.
type ListLBResponse struct {
	Status  string         `json:"status"`
	Message string         `json:"message"`
	Data    []LoadBalancer `json:"data"`
}

// List returns a list of all load balancers.
func (s *LoadBalancerService) List() ([]LoadBalancer, error) {
	var resp ListLBResponse
	err := s.client.Request(http.MethodGet, "/loadbalancer", nil, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return resp.Data, nil
}

// CreateParams represents the parameters for creating a load balancer.
type CreateParams struct {
	DCSlug string `json:"dcslug"`
	Name   string `json:"name"`
	Type   string `json:"type"`
}

// CreateResponse represents the response for creating a load balancer.
type CreateResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	LBID    string `json:"lbid"`
}

// Create creates a new load balancer.
func (s *LoadBalancerService) Create(params CreateParams) (*CreateResponse, error) {
	var resp CreateResponse
	err := s.client.Request(http.MethodPost, "/loadbalancer/add", params, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return &resp, nil
}
