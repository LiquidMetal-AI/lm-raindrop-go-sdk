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

func TestQueryProcedureSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Query.Procedures.Search(context.TODO(), raindrop.QueryProcedureSearchParams{
		BodySmartMemoryLocation1: raindrop.QueryProcedureSearchParamsSmartMemoryLocationUnion{
			OfModuleID: &raindrop.QueryProcedureSearchParamsSmartMemoryLocationModuleID{
				ModuleID: "moduleId",
			},
		},
		Terms:                   "system prompt",
		BodyNMostRecent1:        raindrop.Int(10),
		BodyNMostRecent2:        raindrop.Int(10),
		BodyProceduralMemoryID1: raindrop.String("demo-smartmemory"),
		BodyProceduralMemoryID2: raindrop.String("demo-smartmemory"),
		BodySearchKeys1:         raindrop.Bool(true),
		BodySearchValues1:       raindrop.Bool(true),
		BodySearchKeys2:         raindrop.Bool(true),
		BodySearchValues2:       raindrop.Bool(true),
		BodySmartMemoryLocation2: raindrop.QueryProcedureSearchParamsSmartMemoryLocationUnion{
			OfModuleID: &raindrop.QueryProcedureSearchParamsSmartMemoryLocationModuleID{
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
