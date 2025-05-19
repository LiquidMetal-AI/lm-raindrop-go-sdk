// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package raindrop

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/apijson"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/requestconfig"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/option"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/param"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/respjson"
)

// ObjectService contains methods and other services that help with interacting
// with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectService] method instead.
type ObjectService struct {
	Options []option.RequestOption
}

// NewObjectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewObjectService(opts ...option.RequestOption) (r ObjectService) {
	r = ObjectService{}
	r.Options = opts
	return
}

// Download a file from a SmartBucket or regular bucket. The bucket parameter (ID)
// is used to identify the bucket to download from. The key is the path to the
// object in the bucket.
func (r *ObjectService) Get(ctx context.Context, key string, query ObjectGetParams, opts ...option.RequestOption) (res *ObjectGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.ModuleID == "" {
		err = errors.New("missing required module_id parameter")
		return
	}
	if key == "" {
		err = errors.New("missing required key parameter")
		return
	}
	path := fmt.Sprintf("v1/object/%s/%s", query.ModuleID, key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all objects in a SmartBucket or regular bucket. The bucket parameter (ID)
// is used to identify the bucket to list objects from.
func (r *ObjectService) List(ctx context.Context, moduleID string, opts ...option.RequestOption) (res *ObjectListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if moduleID == "" {
		err = errors.New("missing required module_id parameter")
		return
	}
	path := fmt.Sprintf("v1/object/%s", moduleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a file from a SmartBucket or regular bucket. The bucket parameter (ID) is
// used to identify the bucket to delete from. The key is the path to the object in
// the bucket.
func (r *ObjectService) Delete(ctx context.Context, key string, body ObjectDeleteParams, opts ...option.RequestOption) (res *ObjectDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.ModuleID == "" {
		err = errors.New("missing required module_id parameter")
		return
	}
	if key == "" {
		err = errors.New("missing required key parameter")
		return
	}
	path := fmt.Sprintf("v1/object/%s/%s", body.ModuleID, key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Upload a file to a SmartBucket or regular bucket. The bucket parameter (ID) is
// used to identify the bucket to upload to. The key is the path to the object in
// the bucket.
func (r *ObjectService) Upload(ctx context.Context, key string, params ObjectUploadParams, opts ...option.RequestOption) (res *ObjectUploadResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.PathModuleID == "" {
		err = errors.New("missing required module_id parameter")
		return
	}
	if key == "" {
		err = errors.New("missing required key parameter")
		return
	}
	path := fmt.Sprintf("v1/object/%s/%s", params.PathModuleID, key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type ObjectGetResponse struct {
	// No specific comments in original for these fields directly, but they were part
	// of the original GetObjectResponse.
	Content     string `json:"content" format:"byte"`
	ContentType string `json:"content_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ContentType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectListResponse struct {
	// List of objects in the bucket with their metadata
	Objects []ObjectListResponseObject `json:"objects"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Objects     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectListResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectInfo represents metadata about a single object
type ObjectListResponseObject struct {
	// MIME type of the object
	ContentType string `json:"content_type"`
	// Object key/path in the bucket
	Key string `json:"key"`
	// Last modification timestamp
	LastModified time.Time `json:"last_modified" format:"date-time"`
	// Size of the object in bytes
	Size ObjectListResponseObjectSizeUnion `json:"size" format:"int64"`
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
func (r ObjectListResponseObject) RawJSON() string { return r.JSON.raw }
func (r *ObjectListResponseObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectListResponseObjectSizeUnion contains all possible properties and values
// from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type ObjectListResponseObjectSizeUnion struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u ObjectListResponseObjectSizeUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectListResponseObjectSizeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ObjectListResponseObjectSizeUnion) RawJSON() string { return u.JSON.raw }

func (r *ObjectListResponseObjectSizeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectDeleteResponse = any

type ObjectUploadResponse struct {
	// Information about the bucket where the object was uploaded
	Bucket ObjectUploadResponseBucket `json:"bucket"`
	// Key/path of the uploaded object
	Key string `json:"key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bucket      respjson.Field
		Key         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectUploadResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectUploadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about the bucket where the object was uploaded
type ObjectUploadResponseBucket struct {
	// **EXAMPLE** "my-app"
	ApplicationName string `json:"application_name"`
	// **EXAMPLE** "01jtryx2f2f61ryk06vd8mr91p"
	ApplicationVersionID string `json:"application_version_id"`
	// **EXAMPLE** "my-smartbucket"
	BucketName string `json:"bucket_name"`
	// **EXAMPLE** "01jtryx2f2f61ryk06vd8mr91p"
	ModuleID string `json:"module_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApplicationName      respjson.Field
		ApplicationVersionID respjson.Field
		BucketName           respjson.Field
		ModuleID             respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectUploadResponseBucket) RawJSON() string { return r.JSON.raw }
func (r *ObjectUploadResponseBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectGetParams struct {
	// Module ID identifying the bucket
	ModuleID string `path:"module_id,required" json:"-"`
	paramObj
}

type ObjectDeleteParams struct {
	// Module ID identifying the bucket
	ModuleID string `path:"module_id,required" json:"-"`
	paramObj
}

type ObjectUploadParams struct {
	// Module ID identifying the bucket
	PathModuleID string `path:"module_id,required" json:"-"`
	// Binary content of the object
	Content string `json:"content,required" format:"byte"`
	// MIME type of the object
	ContentType param.Opt[string] `json:"content_type,omitzero"`
	// Module ID identifying the bucket
	BodyModuleID param.Opt[string] `json:"module_id,omitzero"`
	paramObj
}

func (r ObjectUploadParams) MarshalJSON() (data []byte, err error) {
	type shadow ObjectUploadParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectUploadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
