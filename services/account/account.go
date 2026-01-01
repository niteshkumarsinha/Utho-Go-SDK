package account

import (
	"fmt"
	"net/http"

	"github.com/niteshkumarsinha/utho-sdk-go/internal/client"
)

// AccountService handles communication with the account related methods of the Utho API.
type AccountService struct {
	client *client.Client
}

// NewService creates a new AccountService.
func NewService(client *client.Client) *AccountService {
	return &AccountService{
		client: client,
	}
}

// AccountInfo represents Utho account information.
type AccountInfo struct {
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Status    string `json:"status"`
	Currency  string `json:"currency"`
	Balance   string `json:"balance"`
}

// GetInfoResponse represents the response for account info.
type GetInfoResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    AccountInfo `json:"data"`
}

// GetInfo returns the account information.
func (s *AccountService) GetInfo() (*AccountInfo, error) {
	var resp GetInfoResponse
	err := s.client.Request(http.MethodGet, "/account/", nil, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return &resp.Data, nil
}
