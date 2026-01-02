package utho

import (
	"net/http"
	"time"

	"github.com/niteshkumarsinha/utho-sdk-go/internal/client"
	"github.com/niteshkumarsinha/utho-sdk-go/services/account"
	"github.com/niteshkumarsinha/utho-sdk-go/services/autoscaling"
	"github.com/niteshkumarsinha/utho-sdk-go/services/backups"
	"github.com/niteshkumarsinha/utho-sdk-go/services/cloudserver"
	"github.com/niteshkumarsinha/utho-sdk-go/services/database"
	"github.com/niteshkumarsinha/utho-sdk-go/services/iso"
	"github.com/niteshkumarsinha/utho-sdk-go/services/kubernetes"
	"github.com/niteshkumarsinha/utho-sdk-go/services/loadbalancer"
	"github.com/niteshkumarsinha/utho-sdk-go/services/monitoring"
	"github.com/niteshkumarsinha/utho-sdk-go/services/networking"
	"github.com/niteshkumarsinha/utho-sdk-go/services/objectstorage"
	"github.com/niteshkumarsinha/utho-sdk-go/services/registry"
	"github.com/niteshkumarsinha/utho-sdk-go/services/security"
	"github.com/niteshkumarsinha/utho-sdk-go/services/snapshots"
	"github.com/niteshkumarsinha/utho-sdk-go/services/sqs"
	"github.com/niteshkumarsinha/utho-sdk-go/services/ssl"
	"github.com/niteshkumarsinha/utho-sdk-go/services/stacks"
	"github.com/niteshkumarsinha/utho-sdk-go/services/storage"
	"github.com/niteshkumarsinha/utho-sdk-go/services/transfer"
	"github.com/niteshkumarsinha/utho-sdk-go/services/vpc"
	"github.com/niteshkumarsinha/utho-sdk-go/services/vpn"
	"github.com/niteshkumarsinha/utho-sdk-go/services/waf"
)

const (
	// DefaultBaseURL is the default endpoint for the Utho API v2.
	DefaultBaseURL = "https://api.utho.com/v2"
	// DefaultTimeout is the default timeout for HTTP requests to the API.
	DefaultTimeout = 30 * time.Second
)

// Config holds the configuration for the Utho SDK client.
// It allows setting the BaseURL, APIKey, and a custom HTTPClient.
type Config struct {
	// BaseURL is the Utho API version endpoint (default: https://api.utho.com/v2).
	BaseURL string
	// APIKey is your Utho API key used for authentication.
	APIKey string
	// HTTPClient is an optional custom *http.Client.
	HTTPClient *http.Client
}

// Client is the main entry point for the Utho SDK.
// It provides access to all 22 service clients.
type Client struct {
	config     Config
	httpClient *client.Client

	Account       *account.AccountService
	Autoscaling   *autoscaling.AutoscalingService
	Backups       *backups.BackupsService
	CloudServer   *cloudserver.CloudServerService
	Database      *database.DatabaseService
	ISO           *iso.IsoService
	Kubernetes    *kubernetes.KubernetesService
	LoadBalancer  *loadbalancer.LoadBalancerService
	Monitoring    *monitoring.MonitoringService
	Networking    *networking.NetworkingService
	ObjectStorage *objectstorage.ObjectStorageService
	Registry      *registry.RegistryService
	Security      *security.SecurityService
	Snapshots     *snapshots.SnapshotsService
	SQS           *sqs.SqsService
	SSL           *ssl.SslService
	Stacks        *stacks.StacksService
	Storage       *storage.StorageService
	Transfer      *transfer.TransferService
	VPC           *vpc.VPCService
	VPN           *vpn.VpnService
	WAF           *waf.WafService
}

// NewClient creates a new Utho SDK client with the provided API key.
// It uses default configuration values for BaseURL and HTTPClient.
func NewClient(apiKey string) (*Client, error) {
	return NewClientWithConfig(Config{
		APIKey: apiKey,
	})
}

// NewClientWithConfig creates a new Utho SDK client using the provided custom configuration.
// It defaults missing fields to their respective DefaultBaseURL and DefaultTimeout.
func NewClientWithConfig(cfg Config) (*Client, error) {
	if cfg.BaseURL == "" {
		cfg.BaseURL = DefaultBaseURL
	}
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = &http.Client{
			Timeout: DefaultTimeout,
		}
	}

	c := &client.Client{
		BaseURL:    cfg.BaseURL,
		APIKey:     cfg.APIKey,
		HTTPClient: cfg.HTTPClient,
	}

	return &Client{
		config:        cfg,
		httpClient:    c,
		Account:       account.NewService(c),
		Autoscaling:   autoscaling.NewService(c),
		Backups:       backups.NewService(c),
		CloudServer:   cloudserver.NewService(c),
		Database:      database.NewService(c),
		ISO:           iso.NewService(c),
		Kubernetes:    kubernetes.NewService(c),
		LoadBalancer:  loadbalancer.NewService(c),
		Monitoring:    monitoring.NewService(c),
		Networking:    networking.NewService(c),
		ObjectStorage: objectstorage.NewService(c),
		Registry:      registry.NewService(c),
		Security:      security.NewService(c),
		Snapshots:     snapshots.NewService(c),
		SQS:           sqs.NewService(c),
		SSL:           ssl.NewService(c),
		Stacks:        stacks.NewService(c),
		Storage:       storage.NewService(c),
		Transfer:      transfer.NewService(c),
		VPC:           vpc.NewService(c),
		VPN:           vpn.NewService(c),
		WAF:           waf.NewService(c),
	}, nil
}
