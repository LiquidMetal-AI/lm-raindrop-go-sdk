// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package raindrop_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/testutil"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/option"
)

func TestQueryEpisodicMemorySearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Query.EpisodicMemory.Search(context.TODO(), raindrop.QueryEpisodicMemorySearchParams{
		SmartMemoryLocation: raindrop.QueryEpisodicMemorySearchParamsSmartMemoryLocation{SmartMemory: raindrop.QueryEpisodicMemorySearchParamsSmartMemoryLocationSmartMemory{Name: "memory-name", ApplicationName: raindrop.String("my-app"), Version: raindrop.String("1234")}},
		Terms:               "sessions about user interface preferences",
		EndTime:             raindrop.Time(time.Now()),
		NMostRecent:         raindrop.Int(10),
		StartTime:           raindrop.Time(time.Now()),
	})
	if err != nil {
		var apierr *raindrop.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
