package utho_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/niteshkumarsinha/utho-sdk-go/services/account"
	"github.com/niteshkumarsinha/utho-sdk-go/services/autoscaling"
	"github.com/niteshkumarsinha/utho-sdk-go/services/backups"
	"github.com/niteshkumarsinha/utho-sdk-go/services/kubernetes"
	"github.com/niteshkumarsinha/utho-sdk-go/services/loadbalancer"
	"github.com/niteshkumarsinha/utho-sdk-go/services/networking"
	"github.com/niteshkumarsinha/utho-sdk-go/services/security"
	"github.com/niteshkumarsinha/utho-sdk-go/services/storage"
	"github.com/niteshkumarsinha/utho-sdk-go/services/waf"
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
func TestCloudServerGet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/cloud/123" {
			t.Errorf("Expected path /cloud/123, got: %s", r.URL.Path)
		}
		resp := struct {
			Status  string `json:"status"`
			Message string `json:"message"`
			Data    struct {
				ID       string `json:"id"`
				Hostname string `json:"hostname"`
			} `json:"data"`
		}{
			Status:  "success",
			Message: "Server retrieved",
			Data: struct {
				ID       string `json:"id"`
				Hostname string `json:"hostname"`
			}{
				ID:       "123",
				Hostname: "test-server",
			},
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client, _ := utho.NewClientWithConfig(utho.Config{BaseURL: server.URL})
	instance, err := client.CloudServer.Get("123")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if instance.ID != "123" || instance.Hostname != "test-server" {
		t.Errorf("Unexpected instance data: %+v", instance)
	}
}

func TestDatabaseGet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/databases/456" {
			t.Errorf("Expected path /databases/456, got: %s", r.URL.Path)
		}
		resp := struct {
			Status  string `json:"status"`
			Message string `json:"message"`
			Data    struct {
				ID           string `json:"id"`
				ClusterLabel string `json:"cluster_label"`
			} `json:"data"`
		}{
			Status:  "success",
			Message: "DB retrieved",
			Data: struct {
				ID           string `json:"id"`
				ClusterLabel string `json:"cluster_label"`
			}{
				ID:           "456",
				ClusterLabel: "test-db",
			},
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client, _ := utho.NewClientWithConfig(utho.Config{BaseURL: server.URL})
	db, err := client.Database.Get("456")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if db.ID != "456" || db.ClusterLabel != "test-db" {
		t.Errorf("Unexpected db data: %+v", db)
	}
}

func TestVPCDelete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE method, got: %s", r.Method)
		}
		if r.URL.Path != "/vpc/789" {
			t.Errorf("Expected path /vpc/789, got: %s", r.URL.Path)
		}
		resp := struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{
			Status:  "success",
			Message: "VPC deleted",
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client, _ := utho.NewClientWithConfig(utho.Config{BaseURL: server.URL})
	err := client.VPC.Delete("789")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestNetworkingDomain(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/dns/adddomain" && r.Method == http.MethodPost {
			json.NewEncoder(w).Encode(struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{Status: "success", Message: "Domain added"})
			return
		}
		if r.URL.Path == "/dns/test.com/delete" && r.Method == http.MethodDelete {
			json.NewEncoder(w).Encode(struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{Status: "success", Message: "Domain deleted"})
			return
		}
		t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
	}))
	defer server.Close()

	client, _ := utho.NewClientWithConfig(utho.Config{BaseURL: server.URL})

	// Test Create
	err := client.Networking.CreateDomain(networking.CreateDomainParams{Domain: "test.com"})
	if err != nil {
		t.Fatalf("CreateDomain failed: %v", err)
	}

	// Test Delete
	err = client.Networking.DeleteDomain("test.com")
	if err != nil {
		t.Fatalf("DeleteDomain failed: %v", err)
	}
}

func TestObjectStorageBucketDelete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/objectstorage/my-dc/bucket/my-bucket/delete/" && r.Method == http.MethodDelete {
			json.NewEncoder(w).Encode(struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{Status: "success", Message: "Bucket deleted"})
			return
		}
		t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
	}))
	defer server.Close()

	client, _ := utho.NewClientWithConfig(utho.Config{BaseURL: server.URL})
	err := client.ObjectStorage.DeleteBucket("my-dc", "my-bucket")
	if err != nil {
		t.Fatalf("DeleteBucket failed: %v", err)
	}
}

func TestLoadBalancerUpdate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/loadbalancer/lb-123/update" && r.Method == http.MethodPut {
			json.NewEncoder(w).Encode(struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{Status: "success", Message: "LB updated"})
			return
		}
		t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
	}))
	defer server.Close()

	client, _ := utho.NewClientWithConfig(utho.Config{BaseURL: server.URL})
	err := client.LoadBalancer.Update("lb-123", loadbalancer.UpdateParams{Name: "new-name"})
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}
}

func TestSnapshotDelete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/cloud/c-123/snapshot/s-456/delete" && r.Method == http.MethodDelete {
			json.NewEncoder(w).Encode(struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{Status: "success", Message: "Snapshot deleted"})
			return
		}
		t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
	}))
	defer server.Close()

	client, _ := utho.NewClientWithConfig(utho.Config{BaseURL: server.URL})
	err := client.Snapshots.Delete("c-123", "s-456")
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
}

func TestKubernetesLifecycle(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/kubernetes/k8s-123" && r.Method == http.MethodGet {
			json.NewEncoder(w).Encode(struct {
				Status  string             `json:"status"`
				Message string             `json:"message"`
				Data    kubernetes.Cluster `json:"data"`
			}{Status: "success", Message: "Cluster retrieved", Data: kubernetes.Cluster{ID: "k8s-123"}})
			return
		}
		if r.URL.Path == "/kubernetes/k8s-123/destroy" && r.Method == http.MethodDelete {
			json.NewEncoder(w).Encode(struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{Status: "success", Message: "Cluster deleted"})
			return
		}
		t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
	}))
	defer server.Close()

	client, _ := utho.NewClientWithConfig(utho.Config{BaseURL: server.URL})

	// Test Get
	cluster, err := client.Kubernetes.Get("k8s-123")
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if cluster.ID != "k8s-123" {
		t.Errorf("Expected ID k8s-123, got %s", cluster.ID)
	}

	// Test Delete
	err = client.Kubernetes.Delete("k8s-123")
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
}

func TestStorageLifecycle(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/ebs/vol-123/destroy" && r.Method == http.MethodDelete {
			json.NewEncoder(w).Encode(struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{Status: "success", Message: "Volume deleted"})
			return
		}
		if r.URL.Path == "/ebs/vol-123/attach" && r.Method == http.MethodPost {
			json.NewEncoder(w).Encode(struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{Status: "success", Message: "Volume attached"})
			return
		}
		if r.URL.Path == "/ebs/vol-123/detach" && r.Method == http.MethodPost {
			json.NewEncoder(w).Encode(struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{Status: "success", Message: "Volume detached"})
			return
		}
		t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
	}))
	defer server.Close()

	client, _ := utho.NewClientWithConfig(utho.Config{BaseURL: server.URL})

	// Test Attach
	err := client.Storage.Attach("vol-123", storage.AttachParams{ServerID: "srv-123"})
	if err != nil {
		t.Fatalf("Attach failed: %v", err)
	}

	// Test Detach
	err = client.Storage.Detach("vol-123")
	if err != nil {
		t.Fatalf("Detach failed: %v", err)
	}

	// Test Delete
	err = client.Storage.Delete("vol-123")
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
}

func TestSecurityLifecycle(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/key/key-123/delete" && r.Method == http.MethodDelete {
			json.NewEncoder(w).Encode(struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{Status: "success", Message: "Key deleted"})
			return
		}
		t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
	}))
	defer server.Close()

	client, _ := utho.NewClientWithConfig(utho.Config{BaseURL: server.URL})
	err := client.Security.DeleteSSHKey("key-123")
	if err != nil {
		t.Fatalf("DeleteSSHKey failed: %v", err)
	}
}

func TestAutoscalingLifecycle(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/autoscaling/asg-123/delete" && r.Method == http.MethodDelete {
			json.NewEncoder(w).Encode(struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{Status: "success", Message: "ASG deleted"})
			return
		}
		if r.URL.Path == "/autoscaling/create" && r.Method == http.MethodPost {
			json.NewEncoder(w).Encode(struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{Status: "success", Message: "ASG created"})
			return
		}
		t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
	}))
	defer server.Close()

	client, _ := utho.NewClientWithConfig(utho.Config{BaseURL: server.URL})

	// Test Create
	err := client.Autoscaling.Create(autoscaling.CreateParams{Name: "test-asg"})
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	// Test Delete
	err = client.Autoscaling.Delete("asg-123")
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
}

func TestMonitoringLifecycle(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/monitoring/alerts/pol-123/delete" && r.Method == http.MethodDelete {
			json.NewEncoder(w).Encode(struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{Status: "success", Message: "Policy deleted"})
			return
		}
		t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
	}))
	defer server.Close()

	client, _ := utho.NewClientWithConfig(utho.Config{BaseURL: server.URL})
	err := client.Monitoring.DeleteAlertPolicy("pol-123")
	if err != nil {
		t.Fatalf("DeleteAlertPolicy failed: %v", err)
	}
}

func TestRegistryLifecycle(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/registry/reg-123/delete" && r.Method == http.MethodDelete {
			json.NewEncoder(w).Encode(struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{Status: "success", Message: "Registry deleted"})
			return
		}
		t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
	}))
	defer server.Close()

	client, _ := utho.NewClientWithConfig(utho.Config{BaseURL: server.URL})
	err := client.Registry.Delete("reg-123")
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
}

func TestSQSLifecycle(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/sqs/sqs-123/delete" && r.Method == http.MethodDelete {
			json.NewEncoder(w).Encode(struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{Status: "success", Message: "SQS deleted"})
			return
		}
		t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
	}))
	defer server.Close()

	client, _ := utho.NewClientWithConfig(utho.Config{BaseURL: server.URL})
	err := client.SQS.Delete("sqs-123")
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
}

func TestSSLLifecycle(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/certificates/cert-123" && r.Method == http.MethodDelete {
			json.NewEncoder(w).Encode(struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{Status: "success", Message: "Certificate deleted"})
			return
		}
		t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
	}))
	defer server.Close()

	client, _ := utho.NewClientWithConfig(utho.Config{BaseURL: server.URL})
	err := client.SSL.Delete("cert-123")
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
}

func TestStacksLifecycle(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/stacks/stack-123" && r.Method == http.MethodDelete {
			json.NewEncoder(w).Encode(struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{Status: "success", Message: "Stack deleted"})
			return
		}
		t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
	}))
	defer server.Close()

	client, _ := utho.NewClientWithConfig(utho.Config{BaseURL: server.URL})
	err := client.Stacks.Delete("stack-123")
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
}

func TestVPNLifecycle(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/vpn/vpn-123/delete" && r.Method == http.MethodDelete {
			json.NewEncoder(w).Encode(struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{Status: "success", Message: "VPN deleted"})
			return
		}
		t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
	}))
	defer server.Close()

	client, _ := utho.NewClientWithConfig(utho.Config{BaseURL: server.URL})
	err := client.VPN.Delete("vpn-123")
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
}

func TestWAFLifecycle(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/waf/waf-123/delete" && r.Method == http.MethodDelete {
			json.NewEncoder(w).Encode(struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{Status: "success", Message: "WAF deleted"})
			return
		}
		if r.URL.Path == "/waf/waf-123/attach" && r.Method == http.MethodPost {
			json.NewEncoder(w).Encode(struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{Status: "success", Message: "WAF attached"})
			return
		}
		t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
	}))
	defer server.Close()

	client, _ := utho.NewClientWithConfig(utho.Config{BaseURL: server.URL})

	// Test Attach
	err := client.WAF.Attach("waf-123", waf.AttachParams{ResourceID: "lb-123", ResourceType: "loadbalancer"})
	if err != nil {
		t.Fatalf("Attach failed: %v", err)
	}

	// Test Delete
	err = client.WAF.Delete("waf-123")
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
}

func TestBackupsLifecycle(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/backups/bkp-123/delete" && r.Method == http.MethodDelete {
			json.NewEncoder(w).Encode(struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{Status: "success", Message: "Backup deleted"})
			return
		}
		if r.URL.Path == "/backups/bkp-123/restore" && r.Method == http.MethodPost {
			json.NewEncoder(w).Encode(struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{Status: "success", Message: "Backup restored"})
			return
		}
		t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
	}))
	defer server.Close()

	client, _ := utho.NewClientWithConfig(utho.Config{BaseURL: server.URL})

	// Test Restore
	err := client.Backups.Restore("bkp-123", backups.RestoreParams{CloudID: "cloud-123"})
	if err != nil {
		t.Fatalf("Restore failed: %v", err)
	}

	// Test Delete
	err = client.Backups.Delete("bkp-123")
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
}

func TestSecurityAPIKeys(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/generate" && r.Method == http.MethodPost {
			json.NewEncoder(w).Encode(struct {
				Status  string          `json:"status"`
				Message string          `json:"message"`
				Data    security.APIKey `json:"data"`
			}{Status: "success", Message: "Key generated", Data: security.APIKey{ID: "key-123"}})
			return
		}
		if r.URL.Path == "/api/key-123/delete" && r.Method == http.MethodDelete {
			json.NewEncoder(w).Encode(struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{Status: "success", Message: "Key deleted"})
			return
		}
		t.Errorf("Unexpected request: %s %s", r.Method, r.URL.Path)
	}))
	defer server.Close()

	client, _ := utho.NewClientWithConfig(utho.Config{BaseURL: server.URL})

	// Test Generate
	key, err := client.Security.GenerateAPIKey(security.GenerateAPIKeyParams{Label: "test-key"})
	if err != nil {
		t.Fatalf("GenerateAPIKey failed: %v", err)
	}
	if key.ID != "key-123" {
		t.Errorf("Expected Key ID key-123, got %s", key.ID)
	}

	// Test Delete
	err = client.Security.DeleteAPIKey("key-123")
	if err != nil {
		t.Fatalf("DeleteAPIKey failed: %v", err)
	}
}
