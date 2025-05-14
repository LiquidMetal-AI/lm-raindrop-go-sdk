// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package raindrop

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/apiform"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/apijson"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/requestconfig"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/option"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/respjson"
)

// StorageObjectService contains methods and other services that help with
// interacting with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStorageObjectService] method instead.
type StorageObjectService struct {
	Options []option.RequestOption
}

// NewStorageObjectService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStorageObjectService(opts ...option.RequestOption) (r StorageObjectService) {
	r = StorageObjectService{}
	r.Options = opts
	return
}

// List all objects in a SmartBucket or regular bucket. The bucket parameter (ID)
// is used to identify the bucket to list objects from.
func (r *StorageObjectService) List(ctx context.Context, bucket string, opts ...option.RequestOption) (res *StorageObjectListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if bucket == "" {
		err = errors.New("missing required bucket parameter")
		return
	}
	path := fmt.Sprintf("v1/object/%s", bucket)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a file from a SmartBucket or regular bucket. The bucket parameter (ID) is
// used to identify the bucket to delete from. The key is the path to the object in
// the bucket.
func (r *StorageObjectService) Delete(ctx context.Context, key string, body StorageObjectDeleteParams, opts ...option.RequestOption) (res *StorageObjectDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.Bucket == "" {
		err = errors.New("missing required bucket parameter")
		return
	}
	if key == "" {
		err = errors.New("missing required key parameter")
		return
	}
	path := fmt.Sprintf("v1/object/%s/%s", body.Bucket, key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Download a file from a SmartBucket or regular bucket. The bucket parameter (ID)
// is used to identify the bucket to download from. The key is the path to the
// object in the bucket.
func (r *StorageObjectService) Download(ctx context.Context, key string, query StorageObjectDownloadParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if query.Bucket == "" {
		err = errors.New("missing required bucket parameter")
		return
	}
	if key == "" {
		err = errors.New("missing required key parameter")
		return
	}
	path := fmt.Sprintf("v1/object/%s/%s", query.Bucket, key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Upload a file to a SmartBucket or regular bucket. The bucket parameter (ID) is
// used to identify the bucket to upload to. The key is the path to the object in
// the bucket.
func (r *StorageObjectService) Upload(ctx context.Context, key string, params StorageObjectUploadParams, opts ...option.RequestOption) (res *StorageObjectUploadResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.Bucket == "" {
		err = errors.New("missing required bucket parameter")
		return
	}
	if key == "" {
		err = errors.New("missing required key parameter")
		return
	}
	path := fmt.Sprintf("v1/object/%s/%s", params.Bucket, key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type StorageObjectListResponse struct {
	Objects []StorageObjectListResponseObject `json:"objects"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Objects     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageObjectListResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageObjectListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageObjectListResponseObject struct {
	// MIME type of the object
	ContentType string `json:"content_type"`
	// Object key/path in the bucket
	Key string `json:"key"`
	// Last modification timestamp
	LastModified time.Time `json:"last_modified" format:"date-time"`
	// Size of the object in bytes (as string due to potential BigInt values)
	Size string `json:"size"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentType  respjson.Field
		Key          respjson.Field
		LastModified respjson.Field
		Size         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageObjectListResponseObject) RawJSON() string { return r.JSON.raw }
func (r *StorageObjectListResponseObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageObjectDeleteResponse struct {
	Success bool `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageObjectDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageObjectDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageObjectUploadResponse struct {
	Bucket  string `json:"bucket,required"`
	Key     string `json:"key,required"`
	Success bool   `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bucket      respjson.Field
		Key         respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageObjectUploadResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageObjectUploadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageObjectDeleteParams struct {
	Bucket string `path:"bucket,required" json:"-"`
	paramObj
}

type StorageObjectDownloadParams struct {
	Bucket string `path:"bucket,required" json:"-"`
	paramObj
}

type StorageObjectUploadParams struct {
	Bucket string `path:"bucket,required" json:"-"`
	Body   io.Reader
	paramObj
}

func (r StorageObjectUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
