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
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/shared"
)

func TestGetMemoryGetWithOptionalParams(t *testing.T) {
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
	_, err := client.GetMemory.Get(context.TODO(), raindrop.GetMemoryGetParams{
		SessionID: "01jxanr45haeswhay4n0q8340y",
		SmartMemoryLocation: raindrop.GetMemoryGetParamsSmartMemoryLocationUnion{
			OfSmartMemory: &raindrop.GetMemoryGetParamsSmartMemoryLocationSmartMemory{
				SmartMemory: shared.LiquidmetalV1alpha1SmartMemoryNameParam{
					ApplicationName: raindrop.String("demo"),
					Name:            "memory-name",
					Version:         raindrop.String("1234"),
				},
			},
		},
		EndTime:        raindrop.Time(time.Now()),
		Key:            raindrop.String("user-preference-theme"),
		NMostRecent:    raindrop.Int(10),
		OrganizationID: raindrop.String("organization_id"),
		StartTime:      raindrop.Time(time.Now()),
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
