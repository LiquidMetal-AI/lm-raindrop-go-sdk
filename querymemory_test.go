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

func TestQueryMemorySearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Query.Memory.Search(context.TODO(), raindrop.QueryMemorySearchParams{
		BodySessionID1: "01jxanr45haeswhay4n0q8340y",
		BodySmartMemoryLocation1: raindrop.QueryMemorySearchParamsSmartMemoryLocationUnion{
			OfModuleID: &raindrop.QueryMemorySearchParamsSmartMemoryLocationModuleID{
				ModuleID: "moduleId",
			},
		},
		Terms:            "user interface preferences",
		BodyEndTime1:     raindrop.Time(time.Now()),
		BodyEndTime2:     raindrop.Time(time.Now()),
		BodyNMostRecent1: raindrop.Int(10),
		BodyNMostRecent2: raindrop.Int(10),
		BodySessionID2:   raindrop.String("01jxanr45haeswhay4n0q8340y"),
		BodySmartMemoryLocation2: raindrop.QueryMemorySearchParamsSmartMemoryLocationUnion{
			OfModuleID: &raindrop.QueryMemorySearchParamsSmartMemoryLocationModuleID{
				ModuleID: "moduleId",
			},
		},
		BodyStartTime1: raindrop.Time(time.Now()),
		BodyStartTime2: raindrop.Time(time.Now()),
		Timeline:       raindrop.String("user-conversation-2024"),
	})
	if err != nil {
		var apierr *raindrop.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
