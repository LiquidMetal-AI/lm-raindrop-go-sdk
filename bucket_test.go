// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package raindrop_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/testutil"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/option"
)

func TestBucketListWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := raindrop.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Bucket.List(context.TODO(), raindrop.BucketListParams{
		BucketLocation: raindrop.BucketLocatorParam{Bucket: raindrop.LiquidmetalV1alpha1BucketNameParam{ApplicationName: "my-app", Name: "my-bucket", Version: "01jtryx2f2f61ryk06vd8mr91p"}},
		Partition:      raindrop.String("tenant-123"),
		Prefix:         raindrop.String("documents/"),
	})
	if err != nil {
		var apierr *raindrop.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBucketDeleteWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := raindrop.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Bucket.Delete(context.TODO(), raindrop.BucketDeleteParams{
		BucketLocation: raindrop.BucketLocatorParam{Bucket: raindrop.LiquidmetalV1alpha1BucketNameParam{ApplicationName: "my-app", Name: "my-bucket", Version: "01jtryx2f2f61ryk06vd8mr91p"}},
		Key:            "my-key",
		Partition:      raindrop.String("tenant-123"),
	})
	if err != nil {
		var apierr *raindrop.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBucketGetWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := raindrop.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Bucket.Get(context.TODO(), raindrop.BucketGetParams{
		BucketLocation: raindrop.BucketLocatorParam{Bucket: raindrop.LiquidmetalV1alpha1BucketNameParam{ApplicationName: "my-app", Name: "my-bucket", Version: "01jtryx2f2f61ryk06vd8mr91p"}},
		Key:            "my-key",
		Partition:      raindrop.String("tenant-123"),
	})
	if err != nil {
		var apierr *raindrop.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBucketPutWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := raindrop.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Bucket.Put(context.TODO(), raindrop.BucketPutParams{
		BucketLocation: raindrop.BucketLocatorParam{Bucket: raindrop.LiquidmetalV1alpha1BucketNameParam{ApplicationName: "my-app", Name: "my-bucket", Version: "01jtryx2f2f61ryk06vd8mr91p"}},
		Content:        "U3RhaW5sZXNzIHJvY2tz",
		ContentType:    "application/pdf",
		Key:            "my-key",
		Partition:      raindrop.String("tenant-123"),
	})
	if err != nil {
		var apierr *raindrop.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
