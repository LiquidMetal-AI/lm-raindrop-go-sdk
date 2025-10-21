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

func TestUpdateMetadataUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.UpdateMetadata.Update(context.TODO(), raindrop.UpdateMetadataUpdateParams{
		SmartSqlLocation: raindrop.UpdateMetadataUpdateParamsSmartSqlLocation{SmartSql: raindrop.UpdateMetadataUpdateParamsSmartSqlLocationSmartSql{Name: "analytics-sql", ApplicationName: raindrop.String("data-analytics-app"), Version: raindrop.String("v1.2.0")}},
		Tables: []raindrop.UpdateMetadataUpdateParamsTable{{
			Columns: []raindrop.UpdateMetadataUpdateParamsTableColumn{{
				ColumnName:   raindrop.String("user_id"),
				DataType:     raindrop.String("INTEGER"),
				IsPrimaryKey: raindrop.Bool(true),
				Nullable:     raindrop.Bool(true),
				SampleData:   raindrop.String("123"),
			}},
			CreatedAt: raindrop.Time(time.Now()),
			TableName: raindrop.String("users"),
			UpdatedAt: raindrop.Time(time.Now()),
		}},
		Mode: raindrop.UpdateMetadataUpdateParamsModeUpdateModeMerge,
	})
	if err != nil {
		var apierr *raindrop.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
