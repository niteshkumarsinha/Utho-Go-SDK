package transfer

import (
	"fmt"
	"net/http"

	"github.com/niteshkumarsinha/utho-sdk-go/internal/client"
)

// TransferService handles communication with the resource transfer related methods of the Utho API.
type TransferService struct {
	client *client.Client
}

// NewService creates a new TransferService.
func NewService(client *client.Client) *TransferService {
	return &TransferService{
		client: client,
	}
}

// ReceiveParams represents the parameters for receiving a resource.
type ReceiveParams struct {
	ID    string `json:"id"`
	Token string `json:"token"`
	Type  string `json:"type"`
}

// Receive processes a resource transfer request.
func (s *TransferService) Receive(params ReceiveParams) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	err := s.client.Request(http.MethodPost, "/transfer/process/", params, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}

// Initiate initiates a resource transfer.
func (s *TransferService) Initiate(resourceType, resourceID string) error {
	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	url := fmt.Sprintf("/transfer/%s/%s/", resourceType, resourceID)
	err := s.client.Request(http.MethodGet, url, nil, &resp)
	if err != nil {
		return err
	}
	if resp.Status != "success" {
		return fmt.Errorf("API error: %s", resp.Message)
	}
	return nil
}
