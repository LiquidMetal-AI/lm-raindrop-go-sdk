// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package raindrop

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/apijson"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/apiquery"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/requestconfig"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/option"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/param"
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
func (r *StorageObjectService) ListObjects(ctx context.Context, bucketName string, query StorageObjectListObjectsParams, opts ...option.RequestOption) (res *StorageObjectListObjectsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("v1/object/%s", bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Upload a file to a SmartBucket or regular bucket. The bucket parameter (ID) is
// used to identify the bucket to upload to. The key is the path to the object in
// the bucket.
func (r *StorageObjectService) PutObject(ctx context.Context, objectKey string, params StorageObjectPutObjectParams, opts ...option.RequestOption) (res *StorageObjectPutObjectResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.BucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	if objectKey == "" {
		err = errors.New("missing required object_key parameter")
		return
	}
	path := fmt.Sprintf("v1/object/%s/%s", params.BucketName, objectKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Delete a file from a SmartBucket or regular bucket. The bucket parameter (ID) is
// used to identify the bucket to delete from. The key is the path to the object in
// the bucket.
func (r *StorageObjectService) GetObject(ctx context.Context, objectKey string, params StorageObjectGetObjectParams, opts ...option.RequestOption) (res *StorageObjectGetObjectResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.BucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	if objectKey == "" {
		err = errors.New("missing required object_key parameter")
		return
	}
	path := fmt.Sprintf("v1/object/%s/%s", params.BucketName, objectKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type StorageObjectListObjectsResponse struct {
	// List of objects in the bucket with their metadata
	Objects []StorageObjectListObjectsResponseObject `json:"objects"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Objects     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageObjectListObjectsResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageObjectListObjectsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectInfo represents metadata about a single object
type StorageObjectListObjectsResponseObject struct {
	// MIME type of the object
	ContentType string `json:"content_type"`
	// Object key/path in the bucket
	Key string `json:"key"`
	// Last modification timestamp
	LastModified time.Time `json:"last_modified" format:"date-time"`
	// Size of the object in bytes
	Size StorageObjectListObjectsResponseObjectSizeUnion `json:"size" format:"int64"`
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
func (r StorageObjectListObjectsResponseObject) RawJSON() string { return r.JSON.raw }
func (r *StorageObjectListObjectsResponseObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// StorageObjectListObjectsResponseObjectSizeUnion contains all possible properties
// and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type StorageObjectListObjectsResponseObjectSizeUnion struct {
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

func (u StorageObjectListObjectsResponseObjectSizeUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u StorageObjectListObjectsResponseObjectSizeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u StorageObjectListObjectsResponseObjectSizeUnion) RawJSON() string { return u.JSON.raw }

func (r *StorageObjectListObjectsResponseObjectSizeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageObjectPutObjectResponse struct {
	// Information about the bucket where the object was uploaded
	Bucket StorageObjectPutObjectResponseBucket `json:"bucket"`
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
func (r StorageObjectPutObjectResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageObjectPutObjectResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about the bucket where the object was uploaded
type StorageObjectPutObjectResponseBucket struct {
	ApplicationName      string `json:"application_name"`
	ApplicationVersionID string `json:"application_version_id"`
	BucketName           string `json:"bucket_name"`
	ModuleID             string `json:"module_id"`
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
func (r StorageObjectPutObjectResponseBucket) RawJSON() string { return r.JSON.raw }
func (r *StorageObjectPutObjectResponseBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageObjectGetObjectResponse struct {
	// Information about the bucket where the object is stored
	Bucket StorageObjectGetObjectResponseBucket `json:"bucket"`
	// MIME type of the object
	ContentType string `json:"content_type"`
	// The object data as a base64 encoded string
	Data string `json:"data" format:"byte"`
	// Key/path of the object
	Key string `json:"key"`
	// Last modification timestamp
	LastModified time.Time `json:"last_modified" format:"date-time"`
	// Size of the object in bytes
	Size StorageObjectGetObjectResponseSizeUnion `json:"size" format:"int64"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bucket       respjson.Field
		ContentType  respjson.Field
		Data         respjson.Field
		Key          respjson.Field
		LastModified respjson.Field
		Size         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageObjectGetObjectResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageObjectGetObjectResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about the bucket where the object is stored
type StorageObjectGetObjectResponseBucket struct {
	ApplicationName      string `json:"application_name"`
	ApplicationVersionID string `json:"application_version_id"`
	BucketName           string `json:"bucket_name"`
	ModuleID             string `json:"module_id"`
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
func (r StorageObjectGetObjectResponseBucket) RawJSON() string { return r.JSON.raw }
func (r *StorageObjectGetObjectResponseBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// StorageObjectGetObjectResponseSizeUnion contains all possible properties and
// values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type StorageObjectGetObjectResponseSizeUnion struct {
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

func (u StorageObjectGetObjectResponseSizeUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u StorageObjectGetObjectResponseSizeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u StorageObjectGetObjectResponseSizeUnion) RawJSON() string { return u.JSON.raw }

func (r *StorageObjectGetObjectResponseSizeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageObjectListObjectsParams struct {
	// Module ID identifying the bucket
	ModuleID param.Opt[string] `query:"module_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StorageObjectListObjectsParams]'s query parameters as
// `url.Values`.
func (r StorageObjectListObjectsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StorageObjectPutObjectParams struct {
	BucketName string `path:"bucket_name,required" json:"-"`
	// Binary content of the object
	Content string `json:"content,required" format:"byte"`
	// MIME type of the object
	ContentType param.Opt[string] `json:"content_type,omitzero"`
	// Object key/path in the bucket
	Key param.Opt[string] `json:"key,omitzero"`
	// Module ID identifying the bucket
	ModuleID param.Opt[string] `json:"module_id,omitzero"`
	paramObj
}

func (r StorageObjectPutObjectParams) MarshalJSON() (data []byte, err error) {
	type shadow StorageObjectPutObjectParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StorageObjectPutObjectParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageObjectGetObjectParams struct {
	BucketName string `path:"bucket_name,required" json:"-"`
	// Object key/path to delete
	Key param.Opt[string] `query:"key,omitzero" json:"-"`
	// Module ID identifying the bucket
	ModuleID param.Opt[string] `query:"module_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StorageObjectGetObjectParams]'s query parameters as
// `url.Values`.
func (r StorageObjectGetObjectParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
