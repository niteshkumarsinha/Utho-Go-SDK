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

// CreateDomainParams represents the parameters for creating a domain.
type CreateDomainParams struct {
	Domain string `json:"domain"`
}

// CreateDomain adds a new domain to Utho DNS.
func (s *NetworkingService) CreateDomain(params CreateDomainParams) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	err := s.client.Request(http.MethodPost, "/dns/adddomain", params, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}

// DeleteDomain removes a domain from Utho DNS.
func (s *NetworkingService) DeleteDomain(domain string) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	url := fmt.Sprintf("/dns/%s/delete", domain)
	err := s.client.Request(http.MethodDelete, url, nil, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}

// CreateFirewallParams represents the parameters for creating a firewall.
type CreateFirewallParams struct {
	Name string `json:"name"`
	// Add other necessary fields based on API docs if needed, though 'name' is often minimal required.
	// Documentation says POST /firewall/create.
}

// CreateFirewall creates a new firewall.
func (s *NetworkingService) CreateFirewall(params CreateFirewallParams) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	err := s.client.Request(http.MethodPost, "/firewall/create", params, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}

// DeleteFirewall destroys a firewall by its ID.
func (s *NetworkingService) DeleteFirewall(id string) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	url := fmt.Sprintf("/firewall/%s/destroy", id)
	err := s.client.Request(http.MethodDelete, url, nil, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}
