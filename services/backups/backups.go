package backups

import (
	"fmt"
	"net/http"

	"github.com/niteshkumarsinha/utho-sdk-go/internal/client"
)

// BackupsService handles communication with the backup related methods of the Utho API.
type BackupsService struct {
	client *client.Client
}

// NewService creates a new BackupsService.
func NewService(client *client.Client) *BackupsService {
	return &BackupsService{
		client: client,
	}
}

// Backup represents a Utho cloud server backup.
type Backup struct {
	ID        string `json:"id"`
	CloudID   string `json:"cloudid"`
	Hostname  string `json:"hostname"`
	Size      string `json:"size"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

// ListBackupsResponse represents the response for listing backups.
type ListBackupsResponse struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Data    []Backup `json:"data"`
}

// List returns a list of all backups for the account.
func (s *BackupsService) List() ([]Backup, error) {
	var resp ListBackupsResponse
	err := s.client.Request(http.MethodGet, "/backups", nil, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return resp.Data, nil
}

// Delete destroys a backup.
func (s *BackupsService) Delete(id string) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	url := fmt.Sprintf("/backups/%s/delete", id)
	err := s.client.Request(http.MethodDelete, url, nil, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}

// RestoreParams represents the parameters for restoring a backup.
type RestoreParams struct {
	CloudID string `json:"cloudid"`
}

// Restore restores a backup to a cloud server.
func (s *BackupsService) Restore(backupID string, params RestoreParams) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	url := fmt.Sprintf("/backups/%s/restore", backupID)
	err := s.client.Request(http.MethodPost, url, params, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}
