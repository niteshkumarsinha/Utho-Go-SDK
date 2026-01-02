# Utho Go SDK & CLI

[![Go Reference](https://pkg.go.dev/badge/github.com/niteshkumarsinha/utho-sdk-go.svg)](https://pkg.go.dev/github.com/niteshkumarsinha/utho-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/niteshkumarsinha/utho-sdk-go)](https://goreportcard.com/report/github.com/niteshkumarsinha/utho-sdk-go)

Official Go SDK and Command Line Interface for the [Utho Cloud Platform](https://utho.com). Manage your entire cloud infrastructure programmatically or via terminal with full coverage of all 22 Utho services.

## Features

- 🚀 **Complete API Coverage** - All 22 Utho services with full CRUD operations
- 🧩 **Modular Design** - Import only what you need
- 💻 **Powerful CLI** - Manage resources from your terminal
- 📝 **Well Documented** - Comprehensive examples and godoc comments
- ✅ **Type Safe** - Strongly typed Go structs for all API responses
- 🔧 **Easy Configuration** - Multiple authentication methods

## Table of Contents

- [Installation](#installation)
- [Quick Start](#quick-start)
- [Authentication](#authentication)
- [SDK Usage](#sdk-usage)
- [CLI Usage](#cli-usage)
- [Service Examples](#service-examples)
- [Documentation](#documentation)

## Installation

### Go SDK

```bash
go get github.com/niteshkumarsinha/utho-sdk-go
```

### CLI Tool

#### Build from Source

```bash
git clone https://github.com/niteshkumarsinha/utho-sdk-go.git
cd utho-sdk-go
go build -o utho cmd/utho/main.go
sudo mv utho /usr/local/bin/  # Optional: install globally
```

See [INSTALL.md](INSTALL.md) for detailed installation instructions.

## Quick Start

### SDK Example

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/niteshkumarsinha/utho-sdk-go"
)

func main() {
    // Create client
    client, err := utho.NewClient("your-api-key")
    if err != nil {
        log.Fatal(err)
    }
    
    // List cloud servers
    servers, err := client.CloudServer.List()
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Found %d servers\n", len(servers))
}
```

### CLI Example

```bash
# Configure API key
export UTHO_APIKEY="your-api-key"
# Or use: utho configure

# List cloud servers
utho cloudserver list

# Deploy a new server
utho cloudserver deploy --zone inmumbaizone2 --plan 10045 --image ubuntu-20.04-x64 --hostname my-server
```

## Authentication

### Environment Variable

```bash
export UTHO_APIKEY="your-api-key-here"
```

### Configuration File (CLI)

```bash
utho configure
# Enter your API key when prompted
# Saves to ~/.utho/config.json
```

### Programmatic

```go
// Method 1: Direct API key
client, err := utho.NewClient("your-api-key")

// Method 2: Custom configuration
config := utho.Config{
    APIKey:  "your-api-key",
    BaseURL: "https://api.utho.com/v2",
}
client, err := utho.NewClientWithConfig(config)
```

## SDK Usage

### Client Initialization

```go
import "github.com/niteshkumarsinha/utho-sdk-go"

// Simple initialization
client, err := utho.NewClient("your-api-key")

// With custom HTTP client
httpClient := &http.Client{Timeout: 60 * time.Second}
config := utho.Config{
    APIKey:     "your-api-key",
    HTTPClient: httpClient,
}
client, err := utho.NewClientWithConfig(config)
```

### Available Services

The SDK provides access to all 22 Utho services:

```go
client.Account        // Account information
client.Autoscaling    // Autoscaling groups
client.Backups        // Server backups
client.CloudServer    // Cloud servers (VMs)
client.Database       // Managed databases
client.ISO            // Custom ISO images
client.Kubernetes     // Kubernetes clusters
client.LoadBalancer   // Load balancers
client.Monitoring     // Alert policies
client.Networking     // DNS & Firewalls
client.ObjectStorage  // S3-compatible storage
client.Registry       // Container registries
client.Security       // SSH & API keys
client.Snapshots      // Server snapshots
client.SQS            // Message queues
client.SSL            // SSL certificates
client.Stacks         // Automation templates
client.Storage        // Block storage (EBS)
client.Transfer       // Resource transfers
client.VPC            // Virtual Private Cloud
client.VPN            // VPN instances
client.WAF            // Web Application Firewall
```

## CLI Usage

### Basic Commands

```bash
# Get help
utho --help
utho <service> --help

# Configure credentials
utho configure

# List resources
utho cloudserver list
utho kubernetes list
utho database list

# Get specific resource
utho cloudserver get <server-id>
utho kubernetes get <cluster-id>
```

### Common Patterns

```bash
# Create resources
utho cloudserver deploy --zone <zone> --plan <plan-id> --image <image> --hostname <name>
utho kubernetes create --name <name> --zone <zone> --nodes <count> --plan <plan-id>
utho database create --name <name> --engine mysql --version 8.0 --zone <zone>

# Delete resources
utho cloudserver delete <server-id>
utho kubernetes delete <cluster-id>
utho database delete <db-id>

# Manage resources
utho cloudserver poweroff <server-id>
utho cloudserver poweron <server-id>
utho storage attach <ebs-id> --cloudid <server-id>
```

## Service Examples

### Cloud Server

**SDK:**
```go
// List servers
servers, err := client.CloudServer.List()

// Get server details
server, err := client.CloudServer.Get("server-id")

// Deploy new server
params := cloudserver.DeployParams{
    DCSlug:         "inmumbaizone2",
    PlanID:         "10045",
    BillingCycle:   "hourly",
    Auth:           "password",
    EnablePublicIP: "1",
    Image:          "ubuntu-20.04-x64",
    Cloud: []cloudserver.InstanceConfig{
        {Hostname: "my-server"},
    },
}
resp, err := client.CloudServer.Deploy(params)

// Power operations
err = client.CloudServer.PowerOff("server-id")
err = client.CloudServer.PowerOn("server-id")
err = client.CloudServer.HardReboot("server-id")

// Delete server
err = client.CloudServer.Delete("server-id")
```

**CLI:**
```bash
utho cloudserver list
utho cloudserver get <server-id>
utho cloudserver deploy --zone inmumbaizone2 --plan 10045 --image ubuntu-20.04-x64 --hostname my-server
utho cloudserver poweroff <server-id>
utho cloudserver poweron <server-id>
utho cloudserver delete <server-id>
```

### Kubernetes

**SDK:**
```go
// List clusters
clusters, err := client.Kubernetes.List()

// Create cluster
params := kubernetes.CreateParams{
    ClusterLabel: "my-k8s-cluster",
    DCSlug:       "inmumbaizone2",
    NodePools: []kubernetes.NodePool{
        {
            Label: "worker-pool",
            Count: 3,
            Size:  "10045",
        },
    },
}
err = client.Kubernetes.Create(params)

// Delete cluster
err = client.Kubernetes.Delete("cluster-id")
```

**CLI:**
```bash
utho kubernetes list
utho kubernetes create --name my-cluster --zone inmumbaizone2 --nodes 3 --plan 10045
utho kubernetes delete <cluster-id>
```

### Database

**SDK:**
```go
// List databases
databases, err := client.Database.List()

// Create database
params := database.CreateParams{
    Name:     "my-db",
    Engine:   "mysql",
    Version:  "8.0",
    DCSlug:   "inmumbaizone2",
    PlanID:   "10045",
    Replicas: 2,
}
err = client.Database.Create(params)

// Delete database
err = client.Database.Delete("db-id")
```

**CLI:**
```bash
utho database list
utho database create --name my-db --engine mysql --version 8.0 --zone inmumbaizone2
utho database delete <db-id>
```

### Storage (EBS)

**SDK:**
```go
// List volumes
volumes, err := client.Storage.List()

// Create volume
params := storage.CreateParams{
    Name:     "my-volume",
    DCSlug:   "inmumbaizone2",
    Size:     "50",
    DiskType: "ssd",
}
resp, err := client.Storage.Create(params)

// Attach to server
attachParams := storage.AttachParams{
    ServerID: "server-id",
}
err = client.Storage.Attach("ebs-id", attachParams)

// Detach and delete
err = client.Storage.Detach("ebs-id")
err = client.Storage.Delete("ebs-id")
```

**CLI:**
```bash
utho storage list
utho storage create --name my-volume --zone inmumbaizone2 --size 50 --type ssd
utho storage attach <ebs-id> --cloudid <server-id>
utho storage detach <ebs-id>
utho storage delete <ebs-id>
```

### Load Balancer

**SDK:**
```go
// List load balancers
lbs, err := client.LoadBalancer.List()

// Create load balancer
params := loadbalancer.CreateParams{
    DCSlug: "inmumbaizone2",
    Name:   "my-lb",
    Type:   "http",
}
resp, err := client.LoadBalancer.Create(params)

// Update
updateParams := loadbalancer.UpdateParams{
    Name: "updated-name",
}
err = client.LoadBalancer.Update("lb-id", updateParams)

// Delete
err = client.LoadBalancer.Delete("lb-id")
```

**CLI:**
```bash
utho loadbalancer list
utho loadbalancer create --name my-lb --zone inmumbaizone2 --type http
utho loadbalancer update <lb-id> --name new-name
utho loadbalancer delete <lb-id>
```

### VPC

**SDK:**
```go
// List VPCs
vpcs, err := client.VPC.List()

// Create VPC
params := vpc.CreateParams{
    Name:   "my-vpc",
    Region: "inmumbaizone2",
    Range:  "10.0.0.0/16",
}
err = client.VPC.Create(params)

// Delete VPC
err = client.VPC.Delete("vpc-id")
```

**CLI:**
```bash
utho vpc list
utho vpc create --name my-vpc --region inmumbaizone2 --range 10.0.0.0/16
utho vpc delete <vpc-id>
```

### Object Storage

**SDK:**
```go
// List buckets
buckets, err := client.ObjectStorage.ListBuckets("inmumbaizone2")

// Create bucket
params := objectstorage.CreateBucketParams{
    Name:    "my-bucket",
    DCSlug:  "inmumbaizone2",
    Size:    "250GB",
    Billing: "monthly",
}
resp, err := client.ObjectStorage.CreateBucket(params)

// Manage access keys
keys, err := client.ObjectStorage.ListAccessKeys("inmumbaizone2")
key, err := client.ObjectStorage.CreateAccessKey("inmumbaizone2")

// Delete bucket
err = client.ObjectStorage.DeleteBucket("inmumbaizone2", "my-bucket")
```

**CLI:**
```bash
utho objectstorage list <dcslug>
utho objectstorage create-bucket <dcslug> --name my-bucket --size 250GB
utho objectstorage list-keys <dcslug>
utho objectstorage create-key <dcslug>
utho objectstorage delete-bucket <dcslug> <bucket-name>
```

### Networking (DNS & Firewalls)

**SDK:**
```go
// DNS Management
domains, err := client.Networking.ListDomains()

createParams := networking.CreateDomainParams{
    Domain: "example.com",
}
err = client.Networking.CreateDomain(createParams)
err = client.Networking.DeleteDomain("example.com")

// Firewall Management
firewalls, err := client.Networking.ListFirewalls()

fwParams := networking.CreateFirewallParams{
    Name: "my-firewall",
}
err = client.Networking.CreateFirewall(fwParams)
err = client.Networking.DeleteFirewall("firewall-id")
```

**CLI:**
```bash
# DNS
utho networking list-domains
utho networking create-domain example.com
utho networking delete-domain example.com

# Firewalls
utho networking list-firewalls
utho networking create-firewall my-firewall
utho networking delete-firewall <firewall-id>
```

### Security (SSH & API Keys)

**SDK:**
```go
// SSH Keys
sshKeys, err := client.Security.ListSSHKeys()

importParams := security.ImportSSHKeyParams{
    Name: "my-key",
    Key:  "ssh-rsa AAAAB3...",
}
err = client.Security.ImportSSHKey(importParams)
err = client.Security.DeleteSSHKey("key-id")

// API Keys
apiKeys, err := client.Security.ListAPIKeys()

genParams := security.GenerateAPIKeyParams{
    Label: "my-api-key",
}
key, err := client.Security.GenerateAPIKey(genParams)
err = client.Security.DeleteAPIKey("key-id")
```

**CLI:**
```bash
# SSH Keys
utho security list-sshkeys
utho security import-sshkey --name my-key --key "ssh-rsa..."
utho security delete-sshkey <key-id>

# API Keys
utho security list-apikeys
utho security generate-apikey --label my-key
utho security delete-apikey <key-id>
```

### Autoscaling

**SDK:**
```go
// List groups
groups, err := client.Autoscaling.List()

// Create group
params := autoscaling.CreateParams{
    Name:    "my-as-group",
    MinSize: 2,
    MaxSize: 10,
    Image:   "ubuntu-20.04-x64",
    Plan:    "10045",
    Script:  "#!/bin/bash\necho 'Setup'",
}
err = client.Autoscaling.Create(params)

// Update group
updateParams := autoscaling.UpdateParams{
    MinSize: 3,
    MaxSize: 15,
}
err = client.Autoscaling.Update("group-id", updateParams)

// Delete group
err = client.Autoscaling.Delete("group-id")
```

**CLI:**
```bash
utho autoscaling list
utho autoscaling create --name my-group --min 2 --max 10 --image ubuntu-20.04-x64 --plan 10045
utho autoscaling update <group-id> --min 3 --max 15
utho autoscaling delete <group-id>
```

### Monitoring

**SDK:**
```go
// List alert policies
policies, err := client.Monitoring.ListAlertPolicies()

// Create alert policy
params := monitoring.CreateAlertPolicyParams{
    Label:        "high-cpu-alert",
    ResourceType: "cloudserver",
    Contacts:     []string{"contact-id"},
}
params.Thresholds.CPU = 80
params.Thresholds.RAM = 90

err = client.Monitoring.CreateAlertPolicy(params)

// Delete policy
err = client.Monitoring.DeleteAlertPolicy("policy-id")
```

**CLI:**
```bash
utho monitoring alerts
utho monitoring create-alert --label high-cpu --cpu 80 --ram 90
utho monitoring delete-alert <policy-id>
```

### Additional Services

For complete examples of all 22 services, see the [examples](examples/) directory.

**Other Available Services:**
- **Backups** - Manage server backups and restores
- **Snapshots** - Create point-in-time snapshots
- **ISO** - Upload custom ISO images
- **SQS** - Message queue management
- **Registry** - Container registry management
- **SSL** - SSL certificate management
- **Stacks** - Infrastructure automation templates
- **VPN** - VPN instance management
- **WAF** - Web Application Firewall
- **Transfer** - Resource transfer between accounts
- **Account** - Account information and billing

## Documentation

- **[API Documentation](https://pkg.go.dev/github.com/niteshkumarsinha/utho-sdk-go)** - Complete Go package documentation
- **[HTML Docs](docs/index.html)** - Interactive service documentation
- **[Examples](examples/)** - Runnable examples for all services
- **[Installation Guide](INSTALL.md)** - Detailed installation instructions

## Error Handling

```go
servers, err := client.CloudServer.List()
if err != nil {
    log.Fatalf("Error listing servers: %v", err)
}

// API errors include status and message
resp, err := client.CloudServer.Deploy(params)
if err != nil {
    // Handle error - could be network, auth, or API error
    log.Printf("Deployment failed: %v", err)
    return
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

- **Documentation**: [https://utho.com/docs](https://utho.com/docs)
- **API Reference**: [https://utho.com/api-docs](https://utho.com/api-docs)
- **Issues**: [GitHub Issues](https://github.com/niteshkumarsinha/utho-sdk-go/issues)

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for version history and updates.
