package main

import (
	"fmt"
	"log"
	"os"

	"github.com/niteshkumarsinha/utho-sdk-go"
)

func main() {
	apiKey := os.Getenv("UTHO_API_KEY")
	if apiKey == "" {
		log.Fatal("UTHO_API_KEY environment variable is required")
	}

	client, err := utho.NewClient(apiKey)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	fmt.Println("=== Object Storage Example ===\n")

	dcslug := "inmumbaizone2"

	// 1. List all buckets
	fmt.Printf("1. Listing buckets in %s...\n", dcslug)
	buckets, err := client.ObjectStorage.ListBuckets(dcslug)
	if err != nil {
		log.Fatalf("Error listing buckets: %v", err)
	}
	fmt.Printf("Found %d buckets\n", len(buckets))
	for _, bucket := range buckets {
		fmt.Printf("  - Name: %s, Size: %s, Status: %s\n", bucket.Name, bucket.Size, bucket.Status)
	}

	// 2. List access keys
	fmt.Printf("\n2. Listing access keys for %s...\n", dcslug)
	keys, err := client.ObjectStorage.ListAccessKeys(dcslug)
	if err != nil {
		log.Printf("Error listing access keys: %v", err)
	} else {
		fmt.Printf("Found %d access keys\n", len(keys))
		for _, key := range keys {
			fmt.Printf("  - Access Key: %s, Status: %s\n", key.AccessKey, key.Status)
		}
	}

	// 3. Create a bucket (COMMENTED OUT)
	/*
		fmt.Println("\n3. Creating a bucket...")
		createParams := objectstorage.CreateBucketParams{
			Name:    "my-bucket",
			DCSlug:  dcslug,
			Size:    "250GB",
			Billing: "monthly",
		}
		createResp, err := client.ObjectStorage.CreateBucket(createParams)
		if err != nil {
			log.Printf("Error creating bucket: %v", err)
		} else {
			fmt.Printf("Bucket created: %s\n", createResp.Message)
		}
	*/

	// 4. Create an access key (COMMENTED OUT)
	/*
		fmt.Println("\n4. Creating an access key...")
		accessKey, err := client.ObjectStorage.CreateAccessKey(dcslug)
		if err != nil {
			log.Printf("Error creating access key: %v", err)
		} else {
			fmt.Printf("Access Key: %s\n", accessKey.AccessKey)
			fmt.Printf("Secret Key: %s\n", accessKey.SecretKey)
		}
	*/

	// 5. Delete a bucket (COMMENTED OUT)
	/*
		if len(buckets) > 0 {
			bucketName := buckets[0].Name
			fmt.Printf("\n5. Deleting bucket %s...\n", bucketName)
			err = client.ObjectStorage.DeleteBucket(dcslug, bucketName)
			if err != nil {
				log.Printf("Error deleting bucket: %v", err)
			} else {
				fmt.Println("Bucket deleted successfully")
			}
		}
	*/

	fmt.Println("\n=== Example completed ===")
}
