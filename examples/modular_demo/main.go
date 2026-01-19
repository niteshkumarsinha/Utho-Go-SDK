package main

import (
	"fmt"
	"log"
	"os"

	"github.com/niteshkumarsinha/utho-sdk-go/services/objectstorage"
	"github.com/subosito/gotenv"
)

func main() {
	// Load environment variables
	_ = gotenv.Load()

	apiKey := os.Getenv("UTHO_API_KEY")
	if apiKey == "" {
		log.Fatal("UTHO_API_KEY not found")
	}

	// Modular SDK: Initialize ONLY the service you need!
	// No need to import the entire github.com/niteshkumarsinha/utho-sdk-go package.
	svc, err := objectstorage.NewClient(apiKey)
	if err != nil {
		log.Fatalf("Failed to create service client: %v", err)
	}

	fmt.Println("=== Modular SDK Demo (Object Storage) ===")

	// 1. List Buckets
	fmt.Println("--- Listing Buckets (Modular) ---")
	buckets, err := svc.ListBuckets("innoida")
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		fmt.Printf("Found %d buckets\n", len(buckets))
	}

	fmt.Println("\n=== Demo Complete ===")
}
