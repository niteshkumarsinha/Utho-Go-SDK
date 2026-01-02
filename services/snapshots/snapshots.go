package snapshots

import (
	"fmt"
	"net/http"

	"github.com/niteshkumarsinha/utho-sdk-go/internal/client"
)

// SnapshotsService handles communication with the snapshot related methods of the Utho API.
type SnapshotsService struct {
	client *client.Client
}

// NewService creates a new SnapshotsService.
func NewService(client *client.Client) *SnapshotsService {
	return &SnapshotsService{
		client: client,
	}
}

// Snapshot represents a Utho cloud server snapshot.
type Snapshot struct {
	ID        string `json:"id"`
	CloudID   string `json:"cloudid"`
	Hostname  string `json:"hostname"`
	Name      string `json:"name"`
	Size      string `json:"size"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

// ListSnapshotsResponse represents the response for listing snapshots.
type ListSnapshotsResponse struct {
	Status  string     `json:"status"`
	Message string     `json:"message"`
	Data    []Snapshot `json:"data"`
}

// List returns a list of all snapshots for the account.
func (s *SnapshotsService) List() ([]Snapshot, error) {
	var resp ListSnapshotsResponse
	err := s.client.Request(http.MethodGet, "/snapshots", nil, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return resp.Data, nil
}

// CreateParams represents the parameters for creating a snapshot.
type CreateParams struct {
	CloudID string `json:"cloudid"`
	Name    string `json:"name"`
}

// CreateResponse represents the response for creating a snapshot.
type CreateResponse struct {
	Status     string `json:"status"`
	Message    string `json:"message"`
	SnapshotID string `json:"snapshotid"`
}

// Create creates a new snapshot for a cloud server.
func (s *SnapshotsService) Create(params CreateParams) (*CreateResponse, error) {
	var resp CreateResponse
	err := s.client.Request(http.MethodPost, "/snapshots/create", params, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return &resp, nil
}

// Delete removes a snapshot.
func (s *SnapshotsService) Delete(cloudID, snapshotID string) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	url := fmt.Sprintf("/cloud/%s/snapshot/%s/delete", cloudID, snapshotID)
	err := s.client.Request(http.MethodDelete, url, nil, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}
