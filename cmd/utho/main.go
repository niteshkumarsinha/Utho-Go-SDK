package main

import (
	"fmt"
	"os"

	"github.com/niteshkumarsinha/utho-sdk-go/cmd/utho/commands"
)

func main() {
	if err := commands.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
