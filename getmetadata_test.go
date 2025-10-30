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

func TestGetMetadataGetWithOptionalParams(t *testing.T) {
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
	_, err := client.GetMetadata.Get(context.TODO(), raindrop.GetMetadataGetParams{
		BodySmartSqlLocation1: raindrop.GetMetadataGetParamsSmartSqlLocationUnion{
			OfSmartSql: &raindrop.GetMetadataGetParamsSmartSqlLocationSmartSql{
				SmartSql: raindrop.GetMetadataGetParamsSmartSqlLocationSmartSqlSmartSql{
					Name:            "analytics-sql",
					ApplicationName: raindrop.String("data-analytics-app"),
					ApplicationName: raindrop.String("data-analytics-app"),
					Version:         raindrop.String("v1.2.0"),
				},
			},
		},
		BodySmartSqlLocation2: raindrop.GetMetadataGetParamsSmartSqlLocationUnion{
			OfSmartSql: &raindrop.GetMetadataGetParamsSmartSqlLocationSmartSql{
				SmartSql: raindrop.GetMetadataGetParamsSmartSqlLocationSmartSqlSmartSql{
					Name:            "analytics-sql",
					ApplicationName: raindrop.String("data-analytics-app"),
					ApplicationName: raindrop.String("data-analytics-app"),
					Version:         raindrop.String("v1.2.0"),
				},
			},
		},
		BodyTableName1: raindrop.String("users"),
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
