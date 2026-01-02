# Utho SDK Go - Examples

This directory contains comprehensive examples demonstrating how to use the Utho Go SDK for all 22 services.

## Quick Start

1. **Set your API key**:
   ```bash
   export UTHO_API_KEY="your-api-key-here"
   ```

2. **Run the quick demo** (lists resources from all services):
   ```bash
   cd examples
   go run main.go
   ```

3. **Run individual service examples**:
   ```bash
   cd examples/services/cloudserver
   go run main.go
   ```

## Available Examples

### Compute & Storage
- **[cloudserver](services/cloudserver/)** - Deploy, manage, and control cloud servers
- **[kubernetes](services/kubernetes/)** - Create and manage Kubernetes clusters
- **[storage](services/storage/)** - Manage elastic block storage (EBS) volumes
- **[snapshots](services/snapshots/)** - Create and manage server snapshots
- **[backups](services/backups/)** - Manage server backups and restores
- **[iso](services/iso/)** - Upload and manage custom ISO images

### Networking & Security
- **[networking](services/networking/)** - Manage DNS domains and firewalls
- **[vpc](services/vpc/)** - Create and manage Virtual Private Clouds
- **[loadbalancer](services/loadbalancer/)** - Configure load balancers
- **[security](services/security/)** - Manage SSH keys and API keys
- **[ssl](services/ssl/)** - Upload and manage SSL certificates
- **[vpn](services/vpn/)** - Deploy VPN instances
- **[waf](services/waf/)** - Configure Web Application Firewalls

### Platform Services
- **[database](services/database/)** - Deploy managed database clusters
- **[autoscaling](services/autoscaling/)** - Configure autoscaling groups
- **[monitoring](services/monitoring/)** - Set up alert policies
- **[objectstorage](services/objectstorage/)** - Manage object storage buckets
- **[registry](services/registry/)** - Manage container registries
- **[stacks](services/stacks/)** - Create automation stacks
- **[sqs](services/sqs/)** - Manage SQS instances

### Utilities
- **[account](services/account/)** - Get account information
- **[transfer](services/transfer/)** - Transfer resources between accounts

## Example Pattern

Each service example follows this structure:

```go
package main

import (
    "fmt"
    "log"
    "os"
    
    "github.com/niteshkumarsinha/utho-sdk-go"
)

func main() {
    // Get API key from environment
    apiKey := os.Getenv("UTHO_API_KEY")
    if apiKey == "" {
        log.Fatal("UTHO_API_KEY environment variable is required")
    }
    
    // Create client
    client, err := utho.NewClient(apiKey)
    if err != nil {
        log.Fatalf("Error creating client: %v", err)
    }
    
    // Demonstrate CRUD operations
    // ...
}
```

## Notes

- All examples require the `UTHO_API_KEY` environment variable
- Examples demonstrate CRUD operations where applicable
- Some operations (create, delete) are commented out to prevent accidental resource creation
- Uncomment and modify parameters as needed for your use case
