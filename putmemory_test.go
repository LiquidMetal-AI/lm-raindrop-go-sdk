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
		Content:   "User prefers dark theme for the interface",
		SessionID: "01jxanr45haeswhay4n0q8340y",
		SmartMemoryLocation: raindrop.PutMemoryNewParamsSmartMemoryLocationUnion{
			OfSmartMemory: &raindrop.PutMemoryNewParamsSmartMemoryLocationSmartMemory{
				SmartMemory: raindrop.PutMemoryNewParamsSmartMemoryLocationSmartMemorySmartMemory{
					Name:            "memory-name",
					ApplicationName: raindrop.String("demo"),
					Version:         raindrop.String("1234"),
				},
			},
		},
		Agent:          raindrop.String("assistant-v1"),
		Key:            raindrop.String("user-preference-theme"),
		OrganizationID: raindrop.String("organization_id"),
		Timeline:       raindrop.String("user-conversation-2024"),
		UserID:         raindrop.String("user_id"),
	})
	if err != nil {
		var apierr *raindrop.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
