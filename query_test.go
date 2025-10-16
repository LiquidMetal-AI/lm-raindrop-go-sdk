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
	_, err := client.Query.ChunkSearch(context.TODO(), raindrop.QueryChunkSearchParams{
		BucketLocations: []raindrop.BucketLocatorParam{{Bucket: raindrop.LiquidmetalV1alpha1BucketNameParam{Name: "my-bucket", ApplicationName: raindrop.String("my-app"), Version: raindrop.String("01jtryx2f2f61ryk06vd8mr91p")}}},
		Input:           "Find documents about revenue in Q4 2023",
		RequestID:       "<YOUR-REQUEST-ID>",
		Partition:       raindrop.String("tenant-123"),
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
	_, err := client.Query.DocumentQuery(context.TODO(), raindrop.QueryDocumentQueryParams{
		BucketLocation: raindrop.BucketLocatorParam{Bucket: raindrop.LiquidmetalV1alpha1BucketNameParam{Name: "my-bucket", ApplicationName: raindrop.String("my-app"), Version: raindrop.String("01jtryx2f2f61ryk06vd8mr91p")}},
		Input:          "What are the key points in this document?",
		ObjectID:       "document.pdf",
		RequestID:      "<YOUR-REQUEST-ID>",
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

func TestQueryGetPaginatedSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Query.GetPaginatedSearch(context.TODO(), raindrop.QueryGetPaginatedSearchParams{
		Page:      raindrop.Int(1),
		PageSize:  raindrop.Int(10),
		RequestID: "<YOUR-REQUEST-ID>",
		Partition: raindrop.String("tenant-123"),
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
	_, err := client.Query.Search(context.TODO(), raindrop.QuerySearchParams{
		BucketLocations: []raindrop.BucketLocatorParam{{Bucket: raindrop.LiquidmetalV1alpha1BucketNameParam{Name: "my-bucket", ApplicationName: raindrop.String("my-app"), Version: raindrop.String("01jtryx2f2f61ryk06vd8mr91p")}}},
		Input:           "All my files",
		RequestID:       "<YOUR-REQUEST-ID>",
		Partition:       raindrop.String("tenant-123"),
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
	_, err := client.Query.SumarizePage(context.TODO(), raindrop.QuerySumarizePageParams{
		Page:      1,
		PageSize:  10,
		RequestID: "<YOUR-REQUEST-ID>",
		Partition: raindrop.String("tenant-123"),
	})
	if err != nil {
		var apierr *raindrop.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
