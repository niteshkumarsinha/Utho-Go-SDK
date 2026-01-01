# Utho Go SDK & CLI

[![Go Reference](https://pkg.go.dev/badge/github.com/niteshkumarsinha/utho-sdk-go.svg)](https://pkg.go.dev/github.com/niteshkumarsinha/utho-sdk-go)

The official project for interacting with the Utho Cloud API using Go. This repository contains both the **Utho Go SDK** for programmatic access and the **Utho CLI** for terminal-based management.

## 🚀 Key Features

- **Full API Coverage**: Support for all 22 Utho service categories.
- **Unified Client**: Easy initialization with lazy-loaded services.
- **Powerful CLI**: Manage resources directly from your terminal.
- **Flexible Configuration**: Securely store credentials in a config file or use environment variables.

---

## 🛠 Installation

### Go SDK
Add the SDK to your project:
```bash
go get github.com/niteshkumarsinha/utho-sdk-go
```

### Utho CLI
Build the CLI tool from source:
```bash
go build -o utho ./cmd/utho/main.go
# Optional: Move to your path
mv utho /usr/local/bin/
```

---

## 🔑 Configuration

The SDK and CLI both support two methods of authentication.

### 1. Configuration File (Recommended for CLI)
Run the `configure` command (CLI only) or create the file manually at `~/.utho/config.json`:
```bash
./utho configure
```
```json
{
  "apikey": "your-api-key-here"
}
```

### 2. Environment Variables
Set the `UTHO_APIKEY` environment variable in your `~/.bash_profile` or `~/.zshrc`:
```bash
export UTHO_APIKEY="your-api-key-here"
```

---

## 📖 Usage Examples

### Using the Go SDK
```go
package main

import (
    "fmt"
    "github.com/niteshkumarsinha/utho-sdk-go"
)

func main() {
    client, _ := utho.NewClient("") // Automatically looks for UTHO_APIKEY or config file
    servers, _ := client.CloudServer.List()
    for _, s := range servers {
        fmt.Printf("Server: %s (%s)\n", s.Hostname, s.IP)
    }
}
```

### Using the Utho CLI
```bash
# List all cloud servers
utho cloudserver list

# Get account information
utho account info

# List VPCs
utho vpc list
```

---

## 📊 Service Coverage & CLI Mapping

| Category | SDK Client Field | CLI Command |
| :--- | :--- | :--- |
| **Compute** | `client.CloudServer` | `utho cloudserver` |
| | `client.Autoscaling` | `utho autoscaling` |
| | `client.Backups` | `utho backups` |
| | `client.Snapshots` | `utho snapshots` |
| | `client.ISO` | `utho iso` |
| **Networking** | `client.VPC` | `utho vpc` |
| | `client.LoadBalancer` | `utho loadbalancer` |
| | `client.Networking` | `utho networking` |
| | `client.VPN` | `utho vpn` |
| | `client.WAF` | `utho waf` |
| **Storage** | `client.Storage` | `utho storage` |
| | `client.ObjectStorage` | `utho objectstorage` |
| **Managed Services**| `client.Database` | `utho database` |
| | `client.Kubernetes` | `utho kubernetes` |
| **App Services** | `client.SQS` | `utho sqs` |
| | `client.Registry` | `utho registry` |
| | `client.Stacks` | `utho stacks` |
| | `client.Transfer` | `utho transfer` |
| **Admin & Security**| `client.Account` | `utho account` |
| | `client.Security` | `utho security` |
| | `client.SSL` | `utho ssl` |
| | `client.Monitoring` | `utho monitoring` |

---

## 🧪 Testing
```bash
go test ./...
```

## 📄 License
This project is licensed under the MIT License.
