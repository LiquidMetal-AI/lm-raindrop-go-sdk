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

func TestEndSessionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.EndSession.New(context.TODO(), raindrop.EndSessionNewParams{
		SessionID:           "01jxanr45haeswhay4n0q8340y",
		SmartMemoryLocation: raindrop.EndSessionNewParamsSmartMemoryLocation{SmartMemory: raindrop.EndSessionNewParamsSmartMemoryLocationSmartMemory{Name: "memory-name", ApplicationName: raindrop.String("my-app"), Version: raindrop.String("1234")}},
		Flush:               raindrop.Bool(true),
		SystemPrompt:        raindrop.String("Summarize the key decisions and action items from this session"),
	})
	if err != nil {
		var apierr *raindrop.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
