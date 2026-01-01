package kubernetes

import (
	"fmt"
	"net/http"

	"github.com/niteshkumarsinha/utho-sdk-go/internal/client"
)

// KubernetesService handles communication with the kubernetes related methods of the Utho API.
type KubernetesService struct {
	client *client.Client
}

// NewService creates a new KubernetesService.
func NewService(client *client.Client) *KubernetesService {
	return &KubernetesService{
		client: client,
	}
}

// Cluster represents a Utho Kubernetes cluster.
type Cluster struct {
	ID             string `json:"id"`
	ClusterLabel   string `json:"cluster_label"`
	ClusterVersion string `json:"cluster_version"`
	Status         string `json:"status"`
	CreatedAt      string `json:"created_at"`
	DC             string `json:"dc"`
}

// ListClustersResponse represents the response for listing clusters.
type ListClustersResponse struct {
	Status  string    `json:"status"`
	Message string    `json:"message"`
	Data    []Cluster `json:"data"`
}

// List returns a list of all Kubernetes clusters.
func (s *KubernetesService) List() ([]Cluster, error) {
	var resp ListClustersResponse
	err := s.client.Request(http.MethodGet, "/kubernetes/", nil, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return resp.Data, nil
}

// NodePoolConfig represents the configuration for a node pool.
type NodePoolConfig struct {
	Label string `json:"label"`
	Size  string `json:"size"`
	Count int    `json:"count"`
}

// CreateParams represents the parameters for creating a Kubernetes cluster.
type CreateParams struct {
	DCSlug         string           `json:"dcslug"`
	ClusterLabel   string           `json:"cluster_label"`
	ClusterVersion string           `json:"cluster_version"`
	NodePools      []NodePoolConfig `json:"nodepools"`
	VPC            string           `json:"vpc"`
}

// CreateResponse represents the response for creating a cluster.
type CreateResponse struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	ClusterID string `json:"clusterid"`
}

// Create creates a new Kubernetes cluster.
func (s *KubernetesService) Create(params CreateParams) (*CreateResponse, error) {
	var resp CreateResponse
	err := s.client.Request(http.MethodPost, "/kubernetes/deploy", params, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return &resp, nil
}
