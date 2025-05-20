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

func TestManualPagination(t *testing.T) {
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
	page, err := client.Query.GetPaginatedSearch(context.TODO(), raindrop.QueryGetPaginatedSearchParams{
		Page:      raindrop.Int(1),
		PageSize:  raindrop.Int(15),
		RequestID: "123e4567-e89b-12d3-a456-426614174000",
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, query := range page.Results {
		t.Logf("%+v\n", query.ChunkSignature)
	}
	// Prism mock isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, query := range page.Results {
			t.Logf("%+v\n", query.ChunkSignature)
		}
	}
}
