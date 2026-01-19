// Package utho provides a monolithic client wrapper for all 22 Utho Cloud services.
// While individual services can be used independently using their respective packages
// under services/, this package provides a convenient way to access all services
// through a single Client object.
package utho

import (
	"github.com/niteshkumarsinha/utho-sdk-go/client"
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

// Client is the main entry point for the Utho SDK.
// It provides access to all 22 service clients.
type Client struct {
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
func NewClient(apiKey string) (*Client, error) {
	return NewClientWithConfig(client.Config{
		APIKey: apiKey,
	})
}

// NewClientWithConfig creates a new Utho SDK client using the provided custom configuration.
func NewClientWithConfig(cfg client.Config) (*Client, error) {
	c, err := client.NewWithConfig(cfg)
	if err != nil {
		return nil, err
	}

	return &Client{
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
