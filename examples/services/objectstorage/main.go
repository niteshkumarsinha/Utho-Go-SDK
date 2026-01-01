package main

import (
	"fmt"
	"os"

	"github.com/niteshkumarsinha/utho-sdk-go"
	"github.com/niteshkumarsinha/utho-sdk-go/services/objectstorage"
)

func main() {
	apiKey := os.Getenv("UTHO_API_KEY")
	client, _ := utho.NewClient(apiKey)

	dc := "inmumbaizone2"
	fmt.Printf("Creating a bucket in %s...\n", dc)

	params := objectstorage.CreateBucketParams{
		Name:    "my-static-assets",
		DCSlug:  dc,
		Size:    "100", // 100 GB
		Billing: "hourly",
	}

	resp, err := client.ObjectStorage.CreateBucket(params)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Bucket Creation Status: %s\n", resp.Status)
	}

	fmt.Println("\nListing buckets...")
	buckets, _ := client.ObjectStorage.ListBuckets(dc)
	for _, b := range buckets {
		fmt.Printf("- %s (%s GB) [%s]\n", b.Name, b.Size, b.Status)
	}
}
