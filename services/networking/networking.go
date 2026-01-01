package networking

import (
	"fmt"
	"net/http"

	"github.com/niteshkumarsinha/utho-sdk-go/internal/client"
)

// NetworkingService handles communication with the networking (DNS, Firewall) related methods of the Utho API.
type NetworkingService struct {
	client *client.Client
}

// NewService creates a new NetworkingService.
func NewService(client *client.Client) *NetworkingService {
	return &NetworkingService{
		client: client,
	}
}

// Domain represents a DNS domain.
type Domain struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ListDomainsResponse represents the response for listing domains.
type ListDomainsResponse struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Data    []Domain `json:"data"`
}

// ListDomains returns a list of all DNS domains.
func (s *NetworkingService) ListDomains() ([]Domain, error) {
	var resp ListDomainsResponse
	err := s.client.Request(http.MethodGet, "/dns", nil, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return resp.Data, nil
}

// Firewall represents a Utho Firewall.
type Firewall struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

// ListFirewallsResponse represents the response for listing firewalls.
type ListFirewallsResponse struct {
	Status  string     `json:"status"`
	Message string     `json:"message"`
	Data    []Firewall `json:"data"`
}

// ListFirewalls returns a list of all firewalls.
func (s *NetworkingService) ListFirewalls() ([]Firewall, error) {
	var resp ListFirewallsResponse
	err := s.client.Request(http.MethodGet, "/firewall", nil, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return resp.Data, nil
}
