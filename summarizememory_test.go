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

func TestSummarizeMemoryNewWithOptionalParams(t *testing.T) {
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
	_, err := client.SummarizeMemory.New(context.TODO(), raindrop.SummarizeMemoryNewParams{
		AgentMemoryLocation: raindrop.SummarizeMemoryNewParamsAgentMemoryLocationUnion{
			OfAgentMemory: &raindrop.SummarizeMemoryNewParamsAgentMemoryLocationAgentMemory{
				AgentMemory: raindrop.SummarizeMemoryNewParamsAgentMemoryLocationAgentMemoryAgentMemory{
					Name:            "memory-name",
					ApplicationName: raindrop.String("my-app"),
					Version:         raindrop.String("1234"),
				},
			},
		},
		MemoryIDs:    []string{"01jxanr45haeswhay4n0q8340y", "01jxanr45haeswhay4n0q8341z"},
		SessionID:    "01jxanr45haeswhay4n0q8340y",
		SystemPrompt: raindrop.String("Summarize the key decisions and action items"),
	})
	if err != nil {
		var apierr *raindrop.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
