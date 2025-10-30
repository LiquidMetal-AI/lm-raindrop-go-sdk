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

func TestStartSessionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.StartSession.New(context.TODO(), raindrop.StartSessionNewParams{
		BodySmartMemoryLocation1: raindrop.StartSessionNewParamsSmartMemoryLocationUnion{
			OfModuleID: &raindrop.StartSessionNewParamsSmartMemoryLocationModuleID{
				ModuleID: "moduleId",
			},
		},
		BodySmartMemoryLocation2: raindrop.StartSessionNewParamsSmartMemoryLocationUnion{
			OfModuleID: &raindrop.StartSessionNewParamsSmartMemoryLocationModuleID{
				ModuleID: "moduleId",
			},
		},
	})
	if err != nil {
		var apierr *raindrop.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
