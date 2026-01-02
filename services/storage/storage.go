package storage

import (
	"fmt"
	"net/http"

	"github.com/niteshkumarsinha/utho-sdk-go/internal/client"
)

// StorageService handles communication with the elastic block storage related methods of the Utho API.
type StorageService struct {
	client *client.Client
}

// NewService creates a new StorageService.
func NewService(client *client.Client) *StorageService {
	return &StorageService{
		client: client,
	}
}

// EBS represents an Elastic Block Storage volume.
type EBS struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Size     string `json:"size"`
	Type     string `json:"type"`
	Status   string `json:"status"`
	DC       string `json:"dc"`
	Attached string `json:"attached_to"`
}

// ListEBSResponse represents the response for listing EBS volumes.
type ListEBSResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    []EBS  `json:"data"`
}

// List returns a list of all EBS volumes.
func (s *StorageService) List() ([]EBS, error) {
	var resp ListEBSResponse
	err := s.client.Request(http.MethodGet, "/ebs", nil, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return resp.Data, nil
}

// CreateParams represents the parameters for creating an EBS volume.
type CreateParams struct {
	Name     string `json:"name"`
	DCSlug   string `json:"dcslug"`
	Size     string `json:"disk"`
	DiskType string `json:"disk_type"`
}

// CreateResponse represents the response for creating an EBS volume.
type CreateResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	EBSID   string `json:"ebsid"`
}

// Create creates a new EBS volume.
func (s *StorageService) Create(params CreateParams) (*CreateResponse, error) {
	var resp CreateResponse
	err := s.client.Request(http.MethodPost, "/ebs", params, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return &resp, nil
}

// Delete destroys an EBS volume.
func (s *StorageService) Delete(id string) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	url := fmt.Sprintf("/ebs/%s/destroy", id)
	err := s.client.Request(http.MethodDelete, url, nil, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}

// AttachParams represents the parameters for attaching an EBS volume.
type AttachParams struct {
	ServerID string `json:"serverid"`
}

// Attach attaches an EBS volume to a cloud server.
func (s *StorageService) Attach(ebsID string, params AttachParams) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	url := fmt.Sprintf("/ebs/%s/attach", ebsID)
	err := s.client.Request(http.MethodPost, url, params, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}

// Detach detaches an EBS volume from its server.
func (s *StorageService) Detach(ebsID string) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	url := fmt.Sprintf("/ebs/%s/detach", ebsID)
	// Using POST for state change actions to be consistent with other action endpoints (e.g. Cloud Server power actions).
	err := s.client.Request(http.MethodPost, url, nil, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}
