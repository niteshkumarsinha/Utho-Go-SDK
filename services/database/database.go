package database

import (
	"fmt"
	"net/http"

	"github.com/niteshkumarsinha/utho-sdk-go/internal/client"
)

// DatabaseService handles communication with the database related methods of the Utho API.
type DatabaseService struct {
	client *client.Client
}

// NewService creates a new DatabaseService.
func NewService(client *client.Client) *DatabaseService {
	return &DatabaseService{
		client: client,
	}
}

// DBCluster represents a Utho Managed Database cluster.
type DBCluster struct {
	ID             string `json:"id"`
	ClusterLabel   string `json:"cluster_label"`
	ClusterEngine  string `json:"cluster_engine"`
	ClusterVersion string `json:"cluster_version"`
	Status         string `json:"status"`
	DC             string `json:"dc"`
}

// ListDBResponse represents the response for listing database clusters.
type ListDBResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    []DBCluster `json:"data"`
}

// List returns a list of all database clusters.
func (s *DatabaseService) List() ([]DBCluster, error) {
	var resp ListDBResponse
	err := s.client.Request(http.MethodGet, "/databases", nil, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return resp.Data, nil
}

// CreateParams represents the parameters for creating a database cluster.
type CreateParams struct {
	DCSlug         string `json:"dcslug"`
	ClusterLabel   string `json:"cluster_label"`
	ClusterEngine  string `json:"cluster_engine"`
	ClusterVersion string `json:"cluster_version"`
	Size           string `json:"size"`
	ReplicaCount   int    `json:"replica_count"`
}

// CreateResponse represents the response for creating a database cluster.
type CreateResponse struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	ClusterID string `json:"clusterid"`
}

// Create creates a new database cluster.
func (s *DatabaseService) Create(params CreateParams) (*CreateResponse, error) {
	var resp CreateResponse
	err := s.client.Request(http.MethodPost, "/databases", params, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return &resp, nil
}

// Get retrieves details about a single database cluster by its ID.
func (s *DatabaseService) Get(id string) (*DBCluster, error) {
	var resp struct {
		Status  string    `json:"status"`
		Message string    `json:"message"`
		Data    DBCluster `json:"data"`
	}
	url := fmt.Sprintf("/databases/%s", id)
	err := s.client.Request(http.MethodGet, url, nil, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return &resp.Data, nil
}

// Delete destroys a database cluster by its ID.
func (s *DatabaseService) Delete(id string) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	url := fmt.Sprintf("/databases/%s", id)
	err := s.client.Request(http.MethodDelete, url, nil, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}
