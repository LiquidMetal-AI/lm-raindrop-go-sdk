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
	_, err := client.SummarizeMemory.New(context.TODO(), raindrop.SummarizeMemoryNewParams{
		BodyMemoryIDs1: []string{"01jxanr45haeswhay4n0q8340y", "01jxanr45haeswhay4n0q8341z"},
		BodySessionID1: "01jxanr45haeswhay4n0q8340y",
		BodySmartMemoryLocation1: raindrop.SummarizeMemoryNewParamsSmartMemoryLocationUnion{
			OfModuleID: &raindrop.SummarizeMemoryNewParamsSmartMemoryLocationModuleID{
				ModuleID: "moduleId",
			},
		},
		BodyMemoryIDs2: []string{"01jxanr45haeswhay4n0q8340y", "01jxanr45haeswhay4n0q8341z"},
		BodySessionID2: raindrop.String("01jxanr45haeswhay4n0q8340y"),
		BodySmartMemoryLocation2: raindrop.SummarizeMemoryNewParamsSmartMemoryLocationUnion{
			OfModuleID: &raindrop.SummarizeMemoryNewParamsSmartMemoryLocationModuleID{
				ModuleID: "moduleId",
			},
		},
		BodySystemPrompt1: raindrop.String("Summarize the key decisions and action items"),
		BodySystemPrompt2: raindrop.String("Summarize the key decisions and action items"),
	})
	if err != nil {
		var apierr *raindrop.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
