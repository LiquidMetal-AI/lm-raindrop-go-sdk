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
	t.Skip("Prism tests are disabled")
	response, err := client.Query.DocumentQuery(context.TODO(), raindrop.QueryDocumentQueryParams{
		BucketLocation: raindrop.BucketLocatorParam{Bucket: raindrop.LiquidmetalV1alpha1BucketNameParam{ApplicationName: "my-app", Name: "my-bucket", Version: "01jtryx2f2f61ryk06vd8mr91p"}},
		Input:          "What are the key points in this document?",
		ObjectID:       "document.pdf",
		RequestID:      "<YOUR-REQUEST-ID>",
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response.Answer)
}
