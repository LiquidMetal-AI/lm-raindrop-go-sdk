// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package raindrop_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/raindrop-go"
	"github.com/stainless-sdks/raindrop-go/internal/testutil"
	"github.com/stainless-sdks/raindrop-go/option"
)

func TestSummarizePageNewSummaryWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.SummarizePage.NewSummary(context.TODO(), raindrop.SummarizePageNewSummaryParams{
		OrganizationID: raindrop.String("organization_id"),
		Page:           raindrop.Int(0),
		PageSize:       raindrop.Int(0),
		RequestID:      raindrop.String("request_id"),
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
