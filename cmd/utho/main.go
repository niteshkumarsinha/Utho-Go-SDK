package main

import (
	"fmt"
	"os"

	"github.com/niteshkumarsinha/utho-sdk-go/cmd/utho/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
