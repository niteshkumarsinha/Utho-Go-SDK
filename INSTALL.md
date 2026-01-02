# Installing Utho CLI

You can install the Utho CLI on macOS and Linux by building it from source.

## Prerequisites

- [Go](https://go.dev/dl/) 1.18 or later

## Installation Steps

1.  **Clone the repository**:
    ```bash
    git clone https://github.com/niteshkumarsinha/utho-sdk-go.git
    cd utho-sdk-go
    ```

2.  **Build the CLI binary**:
    ```bash
    go build -o utho cmd/utho/main.go
    ```

3.  **Install the binary**:
    Move the binary to a directory in your system's PATH (e.g., `/usr/local/bin`):
    ```bash
    sudo mv utho /usr/local/bin/
    ```

4.  **Verify installation**:
    ```bash
    utho --version
    ```
    (Note: If version command is not yet implemented, try `utho help`)

## Configuration

Before using the CLI, configure your API Key:

```bash
utho configure
```
Enter your Utho API Key when prompted. This will save your credentials to `~/.utho/config.json`.

Alternatively, you can set the `UTHO_APIKEY` environment variable:
```bash
export UTHO_APIKEY=your_api_key_here
```

## Usage

Run `utho help` to see all available commands.

Example: List all cloud servers
```bash
utho cloudserver list
```
