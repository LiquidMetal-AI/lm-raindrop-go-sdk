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

func TestGetPiiDataGetWithOptionalParams(t *testing.T) {
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
	_, err := client.GetPiiData.Get(context.TODO(), raindrop.GetPiiDataGetParams{
		BodySmartSqlLocation1: raindrop.GetPiiDataGetParamsSmartSqlLocationUnion{
			OfSmartSql: &raindrop.GetPiiDataGetParamsSmartSqlLocationSmartSql{
				SmartSql: raindrop.GetPiiDataGetParamsSmartSqlLocationSmartSqlSmartSql{
					Name:            "analytics-sql",
					ApplicationName: raindrop.String("data-analytics-app"),
					ApplicationName: raindrop.String("data-analytics-app"),
					Version:         raindrop.String("v1.2.0"),
				},
			},
		},
		BodyTableName1: "users",
		BodyRecordID1:  raindrop.String("user_123"),
		BodyRecordID2:  raindrop.String("user_123"),
		BodySmartSqlLocation2: raindrop.GetPiiDataGetParamsSmartSqlLocationUnion{
			OfSmartSql: &raindrop.GetPiiDataGetParamsSmartSqlLocationSmartSql{
				SmartSql: raindrop.GetPiiDataGetParamsSmartSqlLocationSmartSqlSmartSql{
					Name:            "analytics-sql",
					ApplicationName: raindrop.String("data-analytics-app"),
					ApplicationName: raindrop.String("data-analytics-app"),
					Version:         raindrop.String("v1.2.0"),
				},
			},
		},
		BodyTableName2: raindrop.String("users"),
	})
	if err != nil {
		var apierr *raindrop.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
