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

func TestGetMemoryGetWithOptionalParams(t *testing.T) {
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
	_, err := client.GetMemory.Get(context.TODO(), raindrop.GetMemoryGetParams{
		AgentMemoryLocation: raindrop.GetMemoryGetParamsAgentMemoryLocationUnion{
			OfAgentMemory: &raindrop.GetMemoryGetParamsAgentMemoryLocationAgentMemory{
				AgentMemory: raindrop.GetMemoryGetParamsAgentMemoryLocationAgentMemoryAgentMemory{
					Name:            "my-agent-memory",
					ApplicationName: raindrop.String("my-app"),
					Version:         raindrop.String("01jtryx2f2f61ryk06vd8mr91p"),
				},
			},
		},
		SessionID:   "01jxanr45haeswhay4n0q8340y",
		EndTime:     raindrop.Time(time.Now()),
		Key:         raindrop.String("user-preference-theme"),
		NMostRecent: raindrop.Int(10),
		StartTime:   raindrop.Time(time.Now()),
		Timeline:    raindrop.String("user-conversation-2024"),
	})
	if err != nil {
		var apierr *raindrop.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
