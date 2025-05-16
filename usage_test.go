// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package raindrop_test

import (
	"context"
	"os"
	"testing"

	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/testutil"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/option"
)

func TestUsage(t *testing.T) {
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
	response, err := client.DocumentQuery.Ask(context.TODO(), raindrop.DocumentQueryAskParams{
		BucketLocation: raindrop.BucketLocatorUnionParam{
			OfBucket: &raindrop.BucketLocatorBucketParam{
				Bucket: raindrop.BucketLocatorBucketBucketParam{},
			},
		},
		Input:     "What are the key points in this document?",
		ObjectID:  "document.pdf",
		RequestID: "123e4567-e89b-12d3-a456-426614174000",
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response.Answer)
}
