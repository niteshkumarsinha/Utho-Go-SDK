package objectstorage

import (
	"fmt"
	"net/http"

	"github.com/niteshkumarsinha/utho-sdk-go/internal/client"
)

// ObjectStorageService handles communication with the object storage related methods of the Utho API.
type ObjectStorageService struct {
	client *client.Client
}

// NewService creates a new ObjectStorageService.
func NewService(client *client.Client) *ObjectStorageService {
	return &ObjectStorageService{
		client: client,
	}
}

// Bucket represents a Utho Object Storage bucket.
type Bucket struct {
	Name      string `json:"name"`
	DC        string `json:"dc"`
	Size      string `json:"size"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

// ListBucketsResponse represents the response for listing buckets.
type ListBucketsResponse struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Data    []Bucket `json:"data"`
}

// ListBuckets returns a list of all buckets in a specific datacenter.
func (s *ObjectStorageService) ListBuckets(dcslug string) ([]Bucket, error) {
	var resp ListBucketsResponse
	err := s.client.Request(http.MethodGet, fmt.Sprintf("/objectstorage/%s/bucket", dcslug), nil, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return resp.Data, nil
}

// CreateBucketParams represents the parameters for creating a bucket.
type CreateBucketParams struct {
	Name    string `json:"name"`
	DCSlug  string `json:"dcslug"`
	Size    string `json:"size"`
	Billing string `json:"billing"`
}

// CreateBucketResponse represents the response for creating a bucket.
type CreateBucketResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// CreateBucket creates a new object storage bucket.
func (s *ObjectStorageService) CreateBucket(params CreateBucketParams) (*CreateBucketResponse, error) {
	var resp CreateBucketResponse
	err := s.client.Request(http.MethodPost, "/objectstorage/bucket/create", params, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("API error: %s", resp.Message)
	}
	return &resp, nil
}
