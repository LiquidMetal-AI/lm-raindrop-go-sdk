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
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/shared"
)

func TestPutMemoryNewWithOptionalParams(t *testing.T) {
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
	_, err := client.PutMemory.New(context.TODO(), raindrop.PutMemoryNewParams{
		Content:                  "User prefers dark theme for the interface",
		BodySessionID1:           "01jxanr45haeswhay4n0q8340y",
		BodySmartMemoryLocation1: raindrop.PutMemoryNewParamsSmartMemoryLocation{SmartMemory: shared.LiquidmetalV1alpha1SmartMemoryNameParam{ApplicationName: raindrop.String("my-app"), Name: "memory-name", Version: raindrop.String("1234"), ApplicationName: raindrop.String("demo")}},
		Agent:                    raindrop.String("assistant-v1"),
		Key:                      raindrop.String("user-preference-theme"),
		BodySessionID2:           raindrop.String("01jxanr45haeswhay4n0q8340y"),
		BodySmartMemoryLocation2: raindrop.PutMemoryNewParamsSmartMemoryLocation{SmartMemory: shared.LiquidmetalV1alpha1SmartMemoryNameParam{ApplicationName: raindrop.String("my-app"), Name: "memory-name", Version: raindrop.String("1234"), ApplicationName: raindrop.String("demo")}},
		Timeline:                 raindrop.String("user-conversation-2024"),
	})
	if err != nil {
		var apierr *raindrop.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
