package vpn

import (
	"fmt"
	"net/http"

	"github.com/niteshkumarsinha/utho-sdk-go/internal/client"
)

// VpnService handles communication with the VPN and IPsec related methods of the Utho API.
type VpnService struct {
	client *client.Client
}

// NewService creates a new VpnService.
func NewService(client *client.Client) *VpnService {
	return &VpnService{
		client: client,
	}
}

// VpnInstance represents a Utho VPN instance.
type VpnInstance struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	IP     string `json:"ip"`
}

// ListVpnResponse represents the response for listing VPN instances.
type ListVpnResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    []VpnInstance `json:"data"`
}

// List returns a list of all VPN instances.
func (s *VpnService) List() ([]VpnInstance, error) {
	var resp ListVpnResponse
	err := s.client.Request(http.MethodGet, "/vpn", nil, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return resp.Data, nil
}
