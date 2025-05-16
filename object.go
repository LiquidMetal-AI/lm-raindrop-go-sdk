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

// List all objects in a SmartBucket or regular bucket. The bucket parameter (ID)
// is used to identify the bucket to list objects from.
func (r *ObjectService) ListObjects(ctx context.Context, bucketName string, query ObjectListObjectsParams, opts ...option.RequestOption) (res *ObjectListObjectsResponse, err error) {
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
func (r *ObjectService) PutObject(ctx context.Context, objectKey string, params ObjectPutObjectParams, opts ...option.RequestOption) (res *ObjectPutObjectResponse, err error) {
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
func (r *ObjectService) GetObject(ctx context.Context, objectKey string, params ObjectGetObjectParams, opts ...option.RequestOption) (res *ObjectGetObjectResponse, err error) {
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

type ObjectListObjectsResponse struct {
	// List of objects in the bucket with their metadata
	Objects []ObjectListObjectsResponseObject `json:"objects"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Objects     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectListObjectsResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectListObjectsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectInfo represents metadata about a single object
type ObjectListObjectsResponseObject struct {
	// MIME type of the object
	ContentType string `json:"content_type"`
	// Object key/path in the bucket
	Key string `json:"key"`
	// Last modification timestamp
	LastModified time.Time `json:"last_modified" format:"date-time"`
	// Size of the object in bytes
	Size ObjectListObjectsResponseObjectSizeUnion `json:"size" format:"int64"`
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
func (r ObjectListObjectsResponseObject) RawJSON() string { return r.JSON.raw }
func (r *ObjectListObjectsResponseObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectListObjectsResponseObjectSizeUnion contains all possible properties and
// values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type ObjectListObjectsResponseObjectSizeUnion struct {
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

func (u ObjectListObjectsResponseObjectSizeUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectListObjectsResponseObjectSizeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ObjectListObjectsResponseObjectSizeUnion) RawJSON() string { return u.JSON.raw }

func (r *ObjectListObjectsResponseObjectSizeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectPutObjectResponse struct {
	// Information about the bucket where the object was uploaded
	Bucket ObjectPutObjectResponseBucket `json:"bucket"`
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
func (r ObjectPutObjectResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectPutObjectResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about the bucket where the object was uploaded
type ObjectPutObjectResponseBucket struct {
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
func (r ObjectPutObjectResponseBucket) RawJSON() string { return r.JSON.raw }
func (r *ObjectPutObjectResponseBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectGetObjectResponse struct {
	// Information about the bucket where the object is stored
	Bucket ObjectGetObjectResponseBucket `json:"bucket"`
	// MIME type of the object
	ContentType string `json:"content_type"`
	// The object data as a base64 encoded string
	Data string `json:"data" format:"byte"`
	// Key/path of the object
	Key string `json:"key"`
	// Last modification timestamp
	LastModified time.Time `json:"last_modified" format:"date-time"`
	// Size of the object in bytes
	Size ObjectGetObjectResponseSizeUnion `json:"size" format:"int64"`
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
func (r ObjectGetObjectResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectGetObjectResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about the bucket where the object is stored
type ObjectGetObjectResponseBucket struct {
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
func (r ObjectGetObjectResponseBucket) RawJSON() string { return r.JSON.raw }
func (r *ObjectGetObjectResponseBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectGetObjectResponseSizeUnion contains all possible properties and values
// from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type ObjectGetObjectResponseSizeUnion struct {
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

func (u ObjectGetObjectResponseSizeUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectGetObjectResponseSizeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ObjectGetObjectResponseSizeUnion) RawJSON() string { return u.JSON.raw }

func (r *ObjectGetObjectResponseSizeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectListObjectsParams struct {
	// Module ID identifying the bucket
	ModuleID param.Opt[string] `query:"module_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectListObjectsParams]'s query parameters as
// `url.Values`.
func (r ObjectListObjectsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectPutObjectParams struct {
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

func (r ObjectPutObjectParams) MarshalJSON() (data []byte, err error) {
	type shadow ObjectPutObjectParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectPutObjectParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectGetObjectParams struct {
	BucketName string `path:"bucket_name,required" json:"-"`
	// Object key/path to delete
	Key param.Opt[string] `query:"key,omitzero" json:"-"`
	// Module ID identifying the bucket
	ModuleID param.Opt[string] `query:"module_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectGetObjectParams]'s query parameters as `url.Values`.
func (r ObjectGetObjectParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
