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

	"github.com/stainless-sdks/raindrop-go/internal/apijson"
	"github.com/stainless-sdks/raindrop-go/internal/apiquery"
	"github.com/stainless-sdks/raindrop-go/internal/requestconfig"
	"github.com/stainless-sdks/raindrop-go/option"
	"github.com/stainless-sdks/raindrop-go/packages/param"
	"github.com/stainless-sdks/raindrop-go/packages/respjson"
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

// **DESCRIPTION** Delete a file from a SmartBucket or regular bucket. The bucket
// parameter (ID) is used to identify the bucket to delete from. The key is the
// path to the object in the bucket.
func (r *ObjectService) Get(ctx context.Context, objectKey string, params ObjectGetParams, opts ...option.RequestOption) (res *ObjectGetResponse, err error) {
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

// **DESCRIPTION** List all objects in a SmartBucket or regular bucket. The bucket
// parameter (ID) is used to identify the bucket to list objects from.
func (r *ObjectService) List(ctx context.Context, bucketName string, query ObjectListParams, opts ...option.RequestOption) (res *ObjectListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("v1/object/%s", bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// PutObject uploads an object to a bucket **DESCRIPTION** Upload a file to a
// SmartBucket or regular bucket. The bucket parameter (ID) is used to identify the
// bucket to upload to. The key is the path to the object in the bucket.
func (r *ObjectService) Upload(ctx context.Context, objectKey string, params ObjectUploadParams, opts ...option.RequestOption) (res *ObjectUploadResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.PathBucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	if objectKey == "" {
		err = errors.New("missing required object_key parameter")
		return
	}
	path := fmt.Sprintf("v1/object/%s/%s", params.PathBucketName, objectKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type BucketResponse struct {
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
func (r BucketResponse) RawJSON() string { return r.JSON.raw }
func (r *BucketResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectGetResponse struct {
	// **DESCRIPTION** Information about the bucket where the object is stored
	// **REQUIRED** true **EXAMPLE** {"module_id": "01jt3vs2nyt2xwk2f54x2bkn84",
	// "bucket_name": "mr-bucket", "application_version_id":
	// "01jt3vs1qggsy39eeyq4k295q1", "application_name": "demo-smartbucket"}
	Bucket BucketResponse `json:"bucket"`
	// **DESCRIPTION** MIME type of the object **REQUIRED** true **EXAMPLE**
	// "application/pdf"
	ContentType string `json:"content_type"`
	// **DESCRIPTION** The object data as a base64 encoded string **REQUIRED** true
	// **EXAMPLE** "SGVsbG8gV29ybGQh"
	Data string `json:"data" format:"byte"`
	// **DESCRIPTION** Key/path of the object **REQUIRED** true **EXAMPLE**
	// "08036c5a50a93da84c5c45ba468c58159d75281e.pdf"
	Key string `json:"key"`
	// **DESCRIPTION** Last modification timestamp **REQUIRED** true **EXAMPLE**
	// "2025-05-05T18:36:43.029Z"
	LastModified time.Time `json:"last_modified" format:"date-time"`
	// **DESCRIPTION** Size of the object in bytes **REQUIRED** true **EXAMPLE** 401107
	Size ObjectGetResponseSizeUnion `json:"size" format:"int64"`
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
func (r ObjectGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectGetResponseSizeUnion contains all possible properties and values from
// [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type ObjectGetResponseSizeUnion struct {
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

func (u ObjectGetResponseSizeUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectGetResponseSizeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ObjectGetResponseSizeUnion) RawJSON() string { return u.JSON.raw }

func (r *ObjectGetResponseSizeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectListResponse struct {
	// **DESCRIPTION** List of objects in the bucket with their metadata **REQUIRED**
	// true **EXAMPLE** [{"key": "08036c5a50a93da84c5c45ba468c58159d75281e.pdf",
	// "size": "401107", "content_type": "application/pdf", "last_modified":
	// "2025-05-05T18:36:43.029Z"}, {"key":
	// "0a29925ccc5e6299e132a73325956a3abef6dd26.pdf", "size": "57173", "content_type":
	// "application/pdf", "last_modified": "2025-05-05T18:36:43.985Z"}, {"key":
	// "0e21835a42a6df2405496f62647058ff855743c1.pdf", "size": "1223197",
	// "content_type": "application/pdf", "last_modified": "2025-05-05T18:36:45.362Z"}]
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
	// **DESCRIPTION** MIME type of the object **REQUIRED** true **EXAMPLE**
	// "application/pdf"
	ContentType string `json:"content_type"`
	// **DESCRIPTION** Object key/path in the bucket **REQUIRED** true **EXAMPLE**
	// "08036c5a50a93da84c5c45ba468c58159d75281e.pdf"
	Key string `json:"key"`
	// **DESCRIPTION** Last modification timestamp **REQUIRED** true **EXAMPLE**
	// "2025-05-05T18:36:43.029Z"
	LastModified time.Time `json:"last_modified" format:"date-time"`
	// **DESCRIPTION** Size of the object in bytes **REQUIRED** true **EXAMPLE** 401107
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

type ObjectUploadResponse struct {
	// **DESCRIPTION** Information about the bucket where the object was uploaded
	// **REQUIRED** true **EXAMPLE** {"module_id": "01jt3vs2nyt2xwk2f54x2bkn84",
	// "bucket_name": "mr-bucket", "application_version_id":
	// "01jt3vs1qggsy39eeyq4k295q1", "application_name": "demo-smartbucket"}
	Bucket BucketResponse `json:"bucket"`
	// **DESCRIPTION** Key/path of the uploaded object **REQUIRED** true **EXAMPLE**
	// "test-object.txt"
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

type ObjectGetParams struct {
	// **DESCRIPTION** Name of the bucket **REQUIRED** true
	BucketName string `path:"bucket_name,required" json:"-"`
	// **DESCRIPTION** Object key/path to delete **REQUIRED** true **EXAMPLE** "my-key"
	Key param.Opt[string] `query:"key,omitzero" json:"-"`
	// **DESCRIPTION** Module ID identifying the bucket **REQUIRED** true **EXAMPLE**
	// "01jtgtrd37acrqf7k24dggg31s"
	ModuleID param.Opt[string] `query:"module_id,omitzero" json:"-"`
	// **DESCRIPTION** Organization ID for access control **REQUIRED** true
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" json:"-"`
	// **DESCRIPTION** User ID for access control **REQUIRED** true
	UserID param.Opt[string] `query:"user_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectGetParams]'s query parameters as `url.Values`.
func (r ObjectGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectListParams struct {
	// **DESCRIPTION** Module ID identifying the bucket **REQUIRED** true **EXAMPLE**
	// "01jtgtrd37acrqf7k24dggg31s"
	ModuleID param.Opt[string] `query:"module_id,omitzero" json:"-"`
	// **DESCRIPTION** Organization ID for access control **REQUIRED** true
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" json:"-"`
	// **DESCRIPTION** User ID for access control **REQUIRED** true
	UserID param.Opt[string] `query:"user_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectListParams]'s query parameters as `url.Values`.
func (r ObjectListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectUploadParams struct {
	// **DESCRIPTION** Name of the bucket **REQUIRED** true
	PathBucketName string `path:"bucket_name,required" json:"-"`
	// **DESCRIPTION** Name of the bucket **REQUIRED** true
	BodyBucketName param.Opt[string] `json:"bucket_name,omitzero"`
	// **DESCRIPTION** Binary content of the object **REQUIRED** true
	Content param.Opt[string] `json:"content,omitzero" format:"byte"`
	// **DESCRIPTION** MIME type of the object **REQUIRED** true **EXAMPLE**
	// "application/pdf"
	ContentType param.Opt[string] `json:"content_type,omitzero"`
	// **DESCRIPTION** Object key/path in the bucket **REQUIRED** true **EXAMPLE**
	// "my-key"
	Key param.Opt[string] `json:"key,omitzero"`
	// **DESCRIPTION** Module ID identifying the bucket **REQUIRED** true **EXAMPLE**
	// "01jtgtrd37acrqf7k24dggg31s"
	ModuleID param.Opt[string] `json:"module_id,omitzero"`
	// **DESCRIPTION** Key/path of the object in the bucket **REQUIRED** true
	ObjectKey param.Opt[string] `json:"object_key,omitzero"`
	// **DESCRIPTION** Organization ID for access control **REQUIRED** true
	OrganizationID param.Opt[string] `json:"organization_id,omitzero"`
	// **DESCRIPTION** User ID for access control **REQUIRED** true
	UserID param.Opt[string] `json:"user_id,omitzero"`
	paramObj
}

func (r ObjectUploadParams) MarshalJSON() (data []byte, err error) {
	type shadow ObjectUploadParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectUploadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
