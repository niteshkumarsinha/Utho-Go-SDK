package utho_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/niteshkumarsinha/utho-sdk-go/services/account"
)

func TestAccountGetInfo(t *testing.T) {
	// Mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/account/" {
			t.Errorf("Expected to request '/account/', got: %s", r.URL.Path)
		}
		if r.Header.Get("Authorization") != "Bearer test-api-key" {
			t.Errorf("Expected Bearer Token auth, got: %s", r.Header.Get("Authorization"))
		}

		resp := account.GetInfoResponse{
			Status:  "success",
			Message: "Account info retrieved",
			Data: account.AccountInfo{
				Email:    "test@example.com",
				Balance:  "100.00",
				Currency: "USD",
			},
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	// Initialize client with mock server URL
	client, _ := utho.NewClientWithConfig(utho.Config{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
	})

	info, err := client.Account.GetInfo()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if info.Email != "test@example.com" {
		t.Errorf("Expected email test@example.com, got: %s", info.Email)
	}
	if info.Balance != "100.00" {
		t.Errorf("Expected balance 100.00, got: %s", info.Balance)
	}
}
