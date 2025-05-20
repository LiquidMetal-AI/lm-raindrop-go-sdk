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

func TestQueryChunkSearchWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Query.ChunkSearch(context.TODO(), raindrop.QueryChunkSearchParams{
		BucketLocations: []raindrop.BucketLocatorUnionParam{{
			OfBucket: &raindrop.BucketLocatorBucketParam{
				Bucket: raindrop.BucketLocatorBucketBucketParam{
					Name:            "my-bucket",
					ApplicationName: raindrop.String("my-app"),
					Version:         raindrop.String("01jtgtraw3b5qbahrhvrj3ygbb"),
				},
			},
		}},
		Input:          "Find documents about revenue in Q4 2023",
		RequestID:      "123e4567-e89b-12d3-a456-426614174000",
		OrganizationID: raindrop.String("organizationId"),
		UserID:         raindrop.String("userId"),
	})
	if err != nil {
		var apierr *raindrop.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestQueryDocumentQueryWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Query.DocumentQuery(context.TODO(), raindrop.QueryDocumentQueryParams{
		BucketLocation: raindrop.BucketLocatorUnionParam{
			OfBucket: &raindrop.BucketLocatorBucketParam{
				Bucket: raindrop.BucketLocatorBucketBucketParam{
					Name:            "my-bucket",
					ApplicationName: raindrop.String("my-app"),
					Version:         raindrop.String("01jtgtraw3b5qbahrhvrj3ygbb"),
				},
			},
		},
		Input:          "What are the key points in this document?",
		ObjectID:       "document.pdf",
		RequestID:      "123e4567-e89b-12d3-a456-426614174000",
		OrganizationID: raindrop.String("organizationId"),
		UserID:         raindrop.String("userId"),
	})
	if err != nil {
		var apierr *raindrop.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestQueryGetPaginatedSearchWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Query.GetPaginatedSearch(context.TODO(), raindrop.QueryGetPaginatedSearchParams{
		Page:           raindrop.Int(2),
		PageSize:       raindrop.Int(10),
		RequestID:      "123e4567-e89b-12d3-a456-426614174000",
		OrganizationID: raindrop.String("organizationId"),
		UserID:         raindrop.String("userId"),
	})
	if err != nil {
		var apierr *raindrop.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestQuerySearchWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Query.Search(context.TODO(), raindrop.QuerySearchParams{
		BucketLocations: []raindrop.BucketLocatorUnionParam{{
			OfBucket: &raindrop.BucketLocatorBucketParam{
				Bucket: raindrop.BucketLocatorBucketBucketParam{
					Name:            "my-bucket",
					ApplicationName: raindrop.String("my-app"),
					Version:         raindrop.String("01jtgtraw3b5qbahrhvrj3ygbb"),
				},
			},
		}},
		Input:          "Show me documents containing credit card numbers or social security numbers",
		RequestID:      "123e4567-e89b-12d3-a456-426614174000",
		OrganizationID: raindrop.String("organizationId"),
		UserID:         raindrop.String("userId"),
	})
	if err != nil {
		var apierr *raindrop.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestQuerySumarizePageWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Query.SumarizePage(context.TODO(), raindrop.QuerySumarizePageParams{
		Page:           1,
		PageSize:       10,
		RequestID:      "123e4567-e89b-12d3-a456-426614174000",
		OrganizationID: raindrop.String("organizationId"),
		UserID:         raindrop.String("userId"),
	})
	if err != nil {
		var apierr *raindrop.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
