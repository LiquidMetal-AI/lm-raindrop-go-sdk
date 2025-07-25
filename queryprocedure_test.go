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

func TestQueryProcedureSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Query.Procedures.Search(context.TODO(), raindrop.QueryProcedureSearchParams{
		SmartMemoryLocation: raindrop.QueryProcedureSearchParamsSmartMemoryLocation{SmartMemory: raindrop.QueryProcedureSearchParamsSmartMemoryLocationSmartMemory{Name: "memory-name", ApplicationName: raindrop.String("my-app"), Version: raindrop.String("1234")}},
		Terms:               "system prompt",
		NMostRecent:         raindrop.Int(10),
		ProceduralMemoryID:  raindrop.String("demo-smartmemory"),
		SearchKeys:          raindrop.Bool(true),
		SearchValues:        raindrop.Bool(true),
	})
	if err != nil {
		var apierr *raindrop.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
