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

func TestPutSemanticMemoryNewWithOptionalParams(t *testing.T) {
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
	_, err := client.PutSemanticMemory.New(context.TODO(), raindrop.PutSemanticMemoryNewParams{
		Document: "document",
		SmartMemoryLocation: raindrop.PutSemanticMemoryNewParamsSmartMemoryLocationUnion{
			OfSmartMemory: &raindrop.PutSemanticMemoryNewParamsSmartMemoryLocationSmartMemory{
				SmartMemory: shared.LiquidmetalV1alpha1SmartMemoryNameParam{
					ApplicationName: raindrop.String("demo"),
					Name:            "memory-name",
					Version:         raindrop.String("1234"),
				},
			},
		},
		OrganizationID: raindrop.String("organization_id"),
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
