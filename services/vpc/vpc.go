package vpc

import (
	"fmt"
	"net/http"

	"github.com/niteshkumarsinha/utho-sdk-go/internal/client"
)

// VPCService handles communication with the VPC related methods of the Utho API.
type VPCService struct {
	client *client.Client
}

// NewService creates a new VPCService.
func NewService(client *client.Client) *VPCService {
	return &VPCService{
		client: client,
	}
}

// VPC represents a Utho VPC.
type VPC struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IPv4Range string `json:"ipv4_range"`
	Region    string `json:"region"`
	IsDefault bool   `json:"is_default"`
	CreatedAt string `json:"created_at"`
}

// ListVPCsResponse represents the response for listing VPCs.
type ListVPCsResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    []VPC  `json:"data"`
}

// List returns a list of all VPCs.
func (s *VPCService) List() ([]VPC, error) {
	var resp ListVPCsResponse
	err := s.client.Request(http.MethodGet, "/vpc/", nil, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return resp.Data, nil
}

// CreateParams represents the parameters for creating a VPC.
type CreateParams struct {
	Name      string `json:"name"`
	Region    string `json:"region"`
	IPv4Range string `json:"ipv4_range"`
}

// CreateResponse represents the response for creating a VPC.
type CreateResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	VPCID   string `json:"vpcid"`
}

// Create provisions a new Virtual Private Cloud (VPC) with the specified configuration.
func (s *VPCService) Create(params CreateParams) (*CreateResponse, error) {
	var resp CreateResponse
	err := s.client.Request(http.MethodPost, "/vpc/", params, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return &resp, nil
}
