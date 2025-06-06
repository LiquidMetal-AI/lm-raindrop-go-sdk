// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package raindrop

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/apijson"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/requestconfig"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/option"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/param"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/respjson"
)

// BucketService contains methods and other services that help with interacting
// with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBucketService] method instead.
type BucketService struct {
	Options []option.RequestOption
}

// NewBucketService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBucketService(opts ...option.RequestOption) (r BucketService) {
	r = BucketService{}
	r.Options = opts
	return
}

// List all objects in a SmartBucket or regular bucket. The bucket parameter (ID)
// is used to identify the bucket to list objects from.
func (r *BucketService) List(ctx context.Context, body BucketListParams, opts ...option.RequestOption) (res *BucketListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/list_objects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete a file from a SmartBucket or regular bucket. The bucket parameter (ID) is
// used to identify the bucket to delete from. The key is the path to the object in
// the bucket.
func (r *BucketService) Delete(ctx context.Context, body BucketDeleteParams, opts ...option.RequestOption) (res *BucketDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/delete_object"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Download a file from a SmartBucket or regular bucket. The bucket parameter (ID)
// is used to identify the bucket to download from. The key is the path to the
// object in the bucket.
func (r *BucketService) Get(ctx context.Context, body BucketGetParams, opts ...option.RequestOption) (res *BucketGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/get_object"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Upload a file to a SmartBucket or regular bucket. The bucket parameter (ID) is
// used to identify the bucket to upload to. The key is the path to the object in
// the bucket.
func (r *BucketService) Put(ctx context.Context, body BucketPutParams, opts ...option.RequestOption) (res *BucketPutResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/put_object"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type BucketListResponse struct {
	// List of objects in the bucket with their metadata.
	Objects []BucketListResponseObject `json:"objects,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Objects     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BucketListResponse) RawJSON() string { return r.JSON.raw }
func (r *BucketListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectInfo represents metadata about a single object
type BucketListResponseObject struct {
	// MIME type of the object
	ContentType string `json:"contentType,required"`
	// Object key/path in the bucket
	Key string `json:"key,required"`
	// Last modification timestamp
	LastModified time.Time `json:"lastModified,required" format:"date-time"`
	// Size of the object in bytes
	Size BucketListResponseObjectSizeUnion `json:"size,required" format:"int64"`
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
func (r BucketListResponseObject) RawJSON() string { return r.JSON.raw }
func (r *BucketListResponseObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BucketListResponseObjectSizeUnion contains all possible properties and values
// from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type BucketListResponseObjectSizeUnion struct {
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

func (u BucketListResponseObjectSizeUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BucketListResponseObjectSizeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BucketListResponseObjectSizeUnion) RawJSON() string { return u.JSON.raw }

func (r *BucketListResponseObjectSizeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BucketDeleteResponse = any

type BucketGetResponse struct {
	// No specific comments in original for these fields directly, but they were part
	// of the original GetObjectResponse.
	Content     string `json:"content" format:"byte"`
	ContentType string `json:"contentType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ContentType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BucketGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BucketGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BucketPutResponse struct {
	// Information about the bucket where the object was uploaded
	Bucket BucketPutResponseBucket `json:"bucket"`
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
func (r BucketPutResponse) RawJSON() string { return r.JSON.raw }
func (r *BucketPutResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about the bucket where the object was uploaded
type BucketPutResponseBucket struct {
	// **EXAMPLE** "my-app"
	ApplicationName string `json:"applicationName"`
	// **EXAMPLE** "01jtryx2f2f61ryk06vd8mr91p"
	ApplicationVersionID string `json:"applicationVersionId"`
	// **EXAMPLE** "my-smartbucket"
	BucketName string `json:"bucketName"`
	// **EXAMPLE** "01jtryx2f2f61ryk06vd8mr91p"
	ModuleID string `json:"moduleId"`
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
func (r BucketPutResponseBucket) RawJSON() string { return r.JSON.raw }
func (r *BucketPutResponseBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BucketListParams struct {
	// The buckets to search. If provided, the search will only return results from
	// these buckets
	BucketLocation BucketLocatorUnionParam `json:"bucketLocation,omitzero,required"`
	paramObj
}

func (r BucketListParams) MarshalJSON() (data []byte, err error) {
	type shadow BucketListParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BucketListParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BucketDeleteParams struct {
	// The buckets to search. If provided, the search will only return results from
	// these buckets
	BucketLocation BucketLocatorUnionParam `json:"bucketLocation,omitzero,required"`
	// Object key/path to delete
	Key string `json:"key,required"`
	paramObj
}

func (r BucketDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow BucketDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BucketDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BucketGetParams struct {
	// The buckets to search. If provided, the search will only return results from
	// these buckets
	BucketLocation BucketLocatorUnionParam `json:"bucketLocation,omitzero,required"`
	// Object key/path to download
	Key string `json:"key,required"`
	paramObj
}

func (r BucketGetParams) MarshalJSON() (data []byte, err error) {
	type shadow BucketGetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BucketGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BucketPutParams struct {
	// The buckets to search. If provided, the search will only return results from
	// these buckets
	BucketLocation BucketLocatorUnionParam `json:"bucketLocation,omitzero,required"`
	// Binary content of the object
	Content string `json:"content,required" format:"byte"`
	// MIME type of the object
	ContentType string `json:"contentType,required"`
	// Object key/path in the bucket
	Key string `json:"key,required"`
	paramObj
}

func (r BucketPutParams) MarshalJSON() (data []byte, err error) {
	type shadow BucketPutParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BucketPutParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
