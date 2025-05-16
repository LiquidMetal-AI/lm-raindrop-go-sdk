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

func TestObjectGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Object.Get(
		context.TODO(),
		"object_key",
		raindrop.ObjectGetParams{
			BucketName:     "bucket_name",
			Key:            raindrop.String("key"),
			ModuleID:       raindrop.String("module_id"),
			OrganizationID: raindrop.String("organization_id"),
			UserID:         raindrop.String("user_id"),
		},
	)
	if err != nil {
		var apierr *raindrop.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectListWithOptionalParams(t *testing.T) {
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
	_, err := client.Object.List(
		context.TODO(),
		"bucket_name",
		raindrop.ObjectListParams{
			ModuleID:       raindrop.String("module_id"),
			OrganizationID: raindrop.String("organization_id"),
			UserID:         raindrop.String("user_id"),
		},
	)
	if err != nil {
		var apierr *raindrop.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectUploadWithOptionalParams(t *testing.T) {
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
	_, err := client.Object.Upload(
		context.TODO(),
		"object_key",
		raindrop.ObjectUploadParams{
			PathBucketName: "bucket_name",
			BodyBucketName: raindrop.String("bucket_name"),
			Content:        raindrop.String("U3RhaW5sZXNzIHJvY2tz"),
			ContentType:    raindrop.String("content_type"),
			Key:            raindrop.String("key"),
			ModuleID:       raindrop.String("module_id"),
			ObjectKey:      raindrop.String("object_key"),
			OrganizationID: raindrop.String("organization_id"),
			UserID:         raindrop.String("user_id"),
		},
	)
	if err != nil {
		var apierr *raindrop.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
