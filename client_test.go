// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package raindrop_test

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/option"
)

type closureTransport struct {
	fn func(req *http.Request) (*http.Response, error)
}

func (t *closureTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.fn(req)
}

func TestUserAgentHeader(t *testing.T) {
	var userAgent string
	client := raindrop.NewClient(
		option.WithAPIKey("My API Key"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					userAgent = req.Header.Get("User-Agent")
					return &http.Response{
						StatusCode: http.StatusOK,
					}, nil
				},
			},
		}),
	)
	client.Query.DocumentQuery(context.Background(), raindrop.QueryDocumentQueryParams{
		BucketLocation: raindrop.BucketLocatorParam{Bucket: raindrop.BucketLocatorBucketParam{Name: "my-bucket", ApplicationName: raindrop.String("my-app"), Version: raindrop.String("01jtryx2f2f61ryk06vd8mr91p")}},
		Input:          "What are the key points in this document?",
		ObjectID:       "document.pdf",
		RequestID:      "<YOUR-REQUEST-ID>",
	})
	if userAgent != fmt.Sprintf("Raindrop/Go %s", internal.PackageVersion) {
		t.Errorf("Expected User-Agent to be correct, but got: %#v", userAgent)
	}
}

func TestRetryAfter(t *testing.T) {
	retryCountHeaders := make([]string, 0)
	client := raindrop.NewClient(
		option.WithAPIKey("My API Key"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					retryCountHeaders = append(retryCountHeaders, req.Header.Get("X-Stainless-Retry-Count"))
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header: http.Header{
							http.CanonicalHeaderKey("Retry-After"): []string{"0.1"},
						},
					}, nil
				},
			},
		}),
	)
	_, err := client.Query.DocumentQuery(context.Background(), raindrop.QueryDocumentQueryParams{
		BucketLocation: raindrop.BucketLocatorParam{Bucket: raindrop.BucketLocatorBucketParam{Name: "my-bucket", ApplicationName: raindrop.String("my-app"), Version: raindrop.String("01jtryx2f2f61ryk06vd8mr91p")}},
		Input:          "What are the key points in this document?",
		ObjectID:       "document.pdf",
		RequestID:      "<YOUR-REQUEST-ID>",
	})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}

	attempts := len(retryCountHeaders)
	if attempts != 3 {
		t.Errorf("Expected %d attempts, got %d", 3, attempts)
	}

	expectedRetryCountHeaders := []string{"0", "1", "2"}
	if !reflect.DeepEqual(retryCountHeaders, expectedRetryCountHeaders) {
		t.Errorf("Expected %v retry count headers, got %v", expectedRetryCountHeaders, retryCountHeaders)
	}
}

func TestDeleteRetryCountHeader(t *testing.T) {
	retryCountHeaders := make([]string, 0)
	client := raindrop.NewClient(
		option.WithAPIKey("My API Key"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					retryCountHeaders = append(retryCountHeaders, req.Header.Get("X-Stainless-Retry-Count"))
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header: http.Header{
							http.CanonicalHeaderKey("Retry-After"): []string{"0.1"},
						},
					}, nil
				},
			},
		}),
		option.WithHeaderDel("X-Stainless-Retry-Count"),
	)
	_, err := client.Query.DocumentQuery(context.Background(), raindrop.QueryDocumentQueryParams{
		BucketLocation: raindrop.BucketLocatorParam{Bucket: raindrop.BucketLocatorBucketParam{Name: "my-bucket", ApplicationName: raindrop.String("my-app"), Version: raindrop.String("01jtryx2f2f61ryk06vd8mr91p")}},
		Input:          "What are the key points in this document?",
		ObjectID:       "document.pdf",
		RequestID:      "<YOUR-REQUEST-ID>",
	})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}

	expectedRetryCountHeaders := []string{"", "", ""}
	if !reflect.DeepEqual(retryCountHeaders, expectedRetryCountHeaders) {
		t.Errorf("Expected %v retry count headers, got %v", expectedRetryCountHeaders, retryCountHeaders)
	}
}

func TestOverwriteRetryCountHeader(t *testing.T) {
	retryCountHeaders := make([]string, 0)
	client := raindrop.NewClient(
		option.WithAPIKey("My API Key"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					retryCountHeaders = append(retryCountHeaders, req.Header.Get("X-Stainless-Retry-Count"))
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header: http.Header{
							http.CanonicalHeaderKey("Retry-After"): []string{"0.1"},
						},
					}, nil
				},
			},
		}),
		option.WithHeader("X-Stainless-Retry-Count", "42"),
	)
	_, err := client.Query.DocumentQuery(context.Background(), raindrop.QueryDocumentQueryParams{
		BucketLocation: raindrop.BucketLocatorParam{Bucket: raindrop.BucketLocatorBucketParam{Name: "my-bucket", ApplicationName: raindrop.String("my-app"), Version: raindrop.String("01jtryx2f2f61ryk06vd8mr91p")}},
		Input:          "What are the key points in this document?",
		ObjectID:       "document.pdf",
		RequestID:      "<YOUR-REQUEST-ID>",
	})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}

	expectedRetryCountHeaders := []string{"42", "42", "42"}
	if !reflect.DeepEqual(retryCountHeaders, expectedRetryCountHeaders) {
		t.Errorf("Expected %v retry count headers, got %v", expectedRetryCountHeaders, retryCountHeaders)
	}
}

func TestRetryAfterMs(t *testing.T) {
	attempts := 0
	client := raindrop.NewClient(
		option.WithAPIKey("My API Key"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					attempts++
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header: http.Header{
							http.CanonicalHeaderKey("Retry-After-Ms"): []string{"100"},
						},
					}, nil
				},
			},
		}),
	)
	_, err := client.Query.DocumentQuery(context.Background(), raindrop.QueryDocumentQueryParams{
		BucketLocation: raindrop.BucketLocatorParam{Bucket: raindrop.BucketLocatorBucketParam{Name: "my-bucket", ApplicationName: raindrop.String("my-app"), Version: raindrop.String("01jtryx2f2f61ryk06vd8mr91p")}},
		Input:          "What are the key points in this document?",
		ObjectID:       "document.pdf",
		RequestID:      "<YOUR-REQUEST-ID>",
	})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}
	if want := 3; attempts != want {
		t.Errorf("Expected %d attempts, got %d", want, attempts)
	}
}

func TestContextCancel(t *testing.T) {
	client := raindrop.NewClient(
		option.WithAPIKey("My API Key"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					<-req.Context().Done()
					return nil, req.Context().Err()
				},
			},
		}),
	)
	cancelCtx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := client.Query.DocumentQuery(cancelCtx, raindrop.QueryDocumentQueryParams{
		BucketLocation: raindrop.BucketLocatorParam{Bucket: raindrop.BucketLocatorBucketParam{Name: "my-bucket", ApplicationName: raindrop.String("my-app"), Version: raindrop.String("01jtryx2f2f61ryk06vd8mr91p")}},
		Input:          "What are the key points in this document?",
		ObjectID:       "document.pdf",
		RequestID:      "<YOUR-REQUEST-ID>",
	})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}
}

func TestContextCancelDelay(t *testing.T) {
	client := raindrop.NewClient(
		option.WithAPIKey("My API Key"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					<-req.Context().Done()
					return nil, req.Context().Err()
				},
			},
		}),
	)
	cancelCtx, cancel := context.WithTimeout(context.Background(), 2*time.Millisecond)
	defer cancel()
	_, err := client.Query.DocumentQuery(cancelCtx, raindrop.QueryDocumentQueryParams{
		BucketLocation: raindrop.BucketLocatorParam{Bucket: raindrop.BucketLocatorBucketParam{Name: "my-bucket", ApplicationName: raindrop.String("my-app"), Version: raindrop.String("01jtryx2f2f61ryk06vd8mr91p")}},
		Input:          "What are the key points in this document?",
		ObjectID:       "document.pdf",
		RequestID:      "<YOUR-REQUEST-ID>",
	})
	if err == nil {
		t.Error("expected there to be a cancel error")
	}
}

func TestContextDeadline(t *testing.T) {
	testTimeout := time.After(3 * time.Second)
	testDone := make(chan struct{})

	deadline := time.Now().Add(100 * time.Millisecond)
	deadlineCtx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	go func() {
		client := raindrop.NewClient(
			option.WithAPIKey("My API Key"),
			option.WithHTTPClient(&http.Client{
				Transport: &closureTransport{
					fn: func(req *http.Request) (*http.Response, error) {
						<-req.Context().Done()
						return nil, req.Context().Err()
					},
				},
			}),
		)
		_, err := client.Query.DocumentQuery(deadlineCtx, raindrop.QueryDocumentQueryParams{
			BucketLocation: raindrop.BucketLocatorParam{Bucket: raindrop.BucketLocatorBucketParam{Name: "my-bucket", ApplicationName: raindrop.String("my-app"), Version: raindrop.String("01jtryx2f2f61ryk06vd8mr91p")}},
			Input:          "What are the key points in this document?",
			ObjectID:       "document.pdf",
			RequestID:      "<YOUR-REQUEST-ID>",
		})
		if err == nil {
			t.Error("expected there to be a deadline error")
		}
		close(testDone)
	}()

	select {
	case <-testTimeout:
		t.Fatal("client didn't finish in time")
	case <-testDone:
		if diff := time.Since(deadline); diff < -30*time.Millisecond || 30*time.Millisecond < diff {
			t.Fatalf("client did not return within 30ms of context deadline, got %s", diff)
		}
	}
}
