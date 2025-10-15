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

func TestPutProcedureNewWithOptionalParams(t *testing.T) {
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
	_, err := client.PutProcedure.New(context.TODO(), raindrop.PutProcedureNewParams{
		Key: "TechnicalReportSystemPrompt",
		SmartMemoryLocation: raindrop.PutProcedureNewParamsSmartMemoryLocationUnion{
			OfSmartMemory: &raindrop.PutProcedureNewParamsSmartMemoryLocationSmartMemory{
				SmartMemory: raindrop.PutProcedureNewParamsSmartMemoryLocationSmartMemorySmartMemory{
					Name:            "memory-name",
					ApplicationName: raindrop.String("demo"),
					Version:         raindrop.String("1234"),
				},
			},
		},
		Value:              "You are a technical documentation assistant...",
		OrganizationID:     raindrop.String("organization_id"),
		ProceduralMemoryID: raindrop.String("demo-smartmemory"),
		UserID:             raindrop.String("user_id"),
	})
	if err != nil {
		var apierr *raindrop.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
