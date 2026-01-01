# Utho Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/niteshkumarsinha/utho-sdk-go.svg)](https://pkg.go.dev/github.com/niteshkumarsinha/utho-sdk-go)

The official Utho SDK for the Go programming language. This SDK allows you to easily interact with the Utho REST API v2 to manage your cloud resources, including Cloud Servers, VPCs, Kubernetes clusters, Databases, and more.

## Features

- **Full Coverage**: Implements all 22 Utho service categories.
- **Modular Design**: Service-specific packages (e.g., `cloudserver`, `vpc`, `kubernetes`) for a clean and efficient development experience.
- **Easy Initialization**: Single client entry point with lazy-loaded services.
- **Developer Friendly**: Built-in Go structs for all request/response models.

## Installation

```bash
go get github.com/niteshkumarsinha/utho-sdk-go
```

## Quick Start

### Authentication

You'll need a Utho API key. You can find or create one in the [Utho Console](https://console.utho.com/).

```go
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/niteshkumarsinha/utho-sdk-go"
)

func main() {
    // 1. Initialize the client
    apiKey := os.Getenv("UTHO_API_KEY")
    client, err := utho.NewClient(apiKey)
    if err != nil {
        log.Fatalf("failed to create client: %v", err)
    }

    // 2. Use a service (e.g., CloudServer)
    servers, err := client.CloudServer.List()
    if err != nil {
        log.Fatalf("failed to list servers: %v", err)
    }

    for _, s := range servers {
        fmt.Printf("Server: %s (%s) - Status: %s\n", s.Hostname, s.IP, s.Status)
    }
}
```

## Service Coverage

The SDK supports the following services:

| Category | Client Field | Package |
| :--- | :--- | :--- |
| **Compute** | `client.CloudServer` | `services/cloudserver` |
| | `client.Autoscaling` | `services/autoscaling` |
| | `client.Backups` | `services/backups` |
| | `client.Snapshots` | `services/snapshots` |
| | `client.ISO` | `services/iso` |
| **Networking** | `client.VPC` | `services/vpc` |
| | `client.LoadBalancer` | `services/loadbalancer` |
| | `client.Networking` | `services/networking` |
| | `client.VPN` | `services/vpn` |
| | `client.WAF` | `services/waf` |
| **Storage** | `client.Storage` | `services/storage` |
| | `client.ObjectStorage` | `services/objectstorage` |
| **Managed Services** | `client.Database` | `services/database` |
| | `client.Kubernetes` | `services/kubernetes` |
| **App Services** | `client.SQS` | `services/sqs` |
| | `client.Registry` | `services/registry` |
| | `client.Stacks` | `services/stacks` |
| | `client.Transfer` | `services/transfer` |
| **Admin & Security** | `client.Account` | `services/account` |
| | `client.Security` | `services/security` |
| | `client.SSL` | `services/ssl` |
| | `client.Monitoring` | `services/monitoring` |

## Examples

Detailed examples for each service can be found in the [examples/services](https://github.com/niteshkumarsinha/utho-sdk-go/tree/main/examples/services) directory.

### Deploying a Cloud Server

```go
params := cloudserver.DeployParams{
    DCSlug:       "inmumbaizone2",
    PlanID:       "10045",
    BillingCycle: "hourly",
    Image:        "ubuntu-22.04-x86_64",
    Cloud: []cloudserver.InstanceConfig{
        {Hostname: "my-web-server"},
    },
}

resp, err := client.CloudServer.Deploy(params)
```

## Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

### Running Tests

```bash
go test ./...
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
