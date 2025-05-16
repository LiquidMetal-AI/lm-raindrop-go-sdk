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

func TestObjectGet(t *testing.T) {
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
			Key:            "key",
			ModuleID:       "module_id",
			OrganizationID: "organization_id",
			UserID:         "user_id",
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

func TestObjectList(t *testing.T) {
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
			ModuleID:       "module_id",
			OrganizationID: "organization_id",
			UserID:         "user_id",
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
			BucketName:          "bucket_name",
			QueryKey:            "key",
			QueryModuleID:       "module_id",
			QueryOrganizationID: "organization_id",
			QueryUserID:         "user_id",
			Content:             raindrop.String("U3RhaW5sZXNzIHJvY2tz"),
			ContentType:         raindrop.String("content_type"),
			BodyKey:             raindrop.String("key"),
			BodyModuleID:        raindrop.String("module_id"),
			BodyOrganizationID:  raindrop.String("organization_id"),
			BodyUserID:          raindrop.String("user_id"),
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
