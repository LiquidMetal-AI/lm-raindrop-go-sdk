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

// ListObjectService contains methods and other services that help with interacting
// with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewListObjectService] method instead.
type ListObjectService struct {
	Options []option.RequestOption
}

// NewListObjectService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewListObjectService(opts ...option.RequestOption) (r ListObjectService) {
	r = ListObjectService{}
	r.Options = opts
	return
}

// List all objects in a SmartBucket or regular bucket. The bucket parameter (ID)
// is used to identify the bucket to list objects from.
func (r *ListObjectService) New(ctx context.Context, body ListObjectNewParams, opts ...option.RequestOption) (res *ListObjectNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/list_objects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ListObjectNewResponse struct {
	// List of objects in the bucket with their metadata
	Objects []ListObjectNewResponseObject `json:"objects"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Objects     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListObjectNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ListObjectNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectInfo represents metadata about a single object
type ListObjectNewResponseObject struct {
	// MIME type of the object
	ContentType string `json:"content_type"`
	// Object key/path in the bucket
	Key string `json:"key"`
	// Last modification timestamp
	LastModified time.Time `json:"last_modified" format:"date-time"`
	// Size of the object in bytes
	Size ListObjectNewResponseObjectSizeUnion `json:"size" format:"int64"`
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
func (r ListObjectNewResponseObject) RawJSON() string { return r.JSON.raw }
func (r *ListObjectNewResponseObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ListObjectNewResponseObjectSizeUnion contains all possible properties and values
// from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type ListObjectNewResponseObjectSizeUnion struct {
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

func (u ListObjectNewResponseObjectSizeUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListObjectNewResponseObjectSizeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ListObjectNewResponseObjectSizeUnion) RawJSON() string { return u.JSON.raw }

func (r *ListObjectNewResponseObjectSizeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListObjectNewParams struct {
	// Module ID identifying the bucket
	ModuleID param.Opt[string] `json:"module_id,omitzero"`
	paramObj
}

func (r ListObjectNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ListObjectNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListObjectNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
