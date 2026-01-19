package main

import (
	"fmt"
	"log"
	"os"

	"github.com/niteshkumarsinha/utho-sdk-go/services/objectstorage"
	"github.com/subosito/gotenv"
)

func main() {
	// Load environment variables from .env
	_ = gotenv.Load() // Ignore error if .env doesn't exist, might be in system env

	apiKey := os.Getenv("UTHO_API_KEY")
	if apiKey == "" {
		log.Fatal("UTHO_API_KEY not found")
	}

	// Modular Initialization: Only initialize the service needed
	svc, err := objectstorage.NewClient(apiKey)
	if err != nil {
		log.Fatalf("Failed to create service client: %v", err)
	}

	dcslug := "innoida"

	fmt.Println("=== Access Key Verification ===")

	// 1. Create Access Key
	fmt.Printf("\n--- 1. Creating Access Key in DC: %s ---\n", dcslug)
	newKey, err := svc.CreateAccessKey(dcslug)
	if err != nil {
		log.Printf("Error creating access key: %v", err)
	} else {
		fmt.Printf("Created Access Key: %s, Secret: %s, Status: %s\n", newKey.AccessKey, newKey.SecretKey, newKey.Status)

		// 2. Modify Access Key Status
		fmt.Printf("\n--- 2. Modifying Access Key Status to 'inactive' ---\n")
		err = svc.UpdateAccessKeyStatus(dcslug, newKey.AccessKey, "inactive")
		if err != nil {
			log.Printf("Error updating access key status: %v", err)
		} else {
			fmt.Println("Access key status updated successfully.")
		}
	}

	// 3. List Access Keys
	fmt.Printf("\n--- 3. Listing Access Keys for DC: %s ---\n", dcslug)
	keys, err := svc.ListAccessKeys(dcslug)
	if err != nil {
		log.Printf("Error listing access keys: %v", err)
	} else {
		fmt.Printf("Found %d access keys\n", len(keys))
		for _, k := range keys {
			fmt.Printf("  - Key: %s, Status: %s\n", k.AccessKey, k.Status)
		}
	}

	fmt.Println("\n=== Object Storage Bucket Verification ===")
	bucketName := "test-bucket-antigravity-" + fmt.Sprintf("%d", os.Getpid())

	// 4. Create Bucket
	fmt.Printf("\n--- 4. Creating Bucket: %s ---\n", bucketName)
	createParams := objectstorage.CreateBucketParams{
		Name:    bucketName,
		DCSlug:  dcslug,
		Size:    "10",
		Billing: "hourly",
	}
	createResp, err := svc.CreateBucket(createParams)
	if err != nil {
		log.Printf("Failed to create bucket: %v", err)
	} else {
		fmt.Printf("Create Response: %+v\n", createResp)
	}

	// 5. Delete Bucket
	fmt.Printf("\n--- 5. Deleting Bucket: %s ---\n", bucketName)
	err = svc.DeleteBucket(dcslug, bucketName)
	if err != nil {
		log.Printf("Error deleting bucket: %v", err)
	} else {
		fmt.Println("Bucket deleted successfully.")
	}
}
