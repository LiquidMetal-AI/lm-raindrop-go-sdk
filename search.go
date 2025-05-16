// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package raindrop

import (
	"context"
	"net/http"

	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/apijson"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/requestconfig"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/option"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/param"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/respjson"
)

// SearchService contains methods and other services that help with interacting
// with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSearchService] method instead.
type SearchService struct {
	Options []option.RequestOption
}

// NewSearchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSearchService(opts ...option.RequestOption) (r SearchService) {
	r = SearchService{}
	r.Options = opts
	return
}

// Primary search endpoint that provides advanced search capabilities across all
// document types stored in SmartBuckets.
//
// Supports recursive object search within objects, enabling nested content search
// like embedded images, text content, and personally identifiable information
// (PII).
//
// The system supports complex queries like:
//
//   - 'Show me documents containing credit card numbers or social security numbers'
//   - 'Find images of landscapes taken during sunset'
//   - 'Get documents mentioning revenue forecasts from Q4 2023'
//   - 'Find me all PDF documents that contain pictures of a cat'
//   - 'Find me all audio files that contain information about the weather in SF in
//     2024'
//
// Key capabilities:
//
// - Natural language query understanding
// - Content-based search across text, images, and audio
// - Automatic PII detection
// - Multi-modal search (text, images, audio)
func (r *SearchService) Find(ctx context.Context, body SearchFindParams, opts ...option.RequestOption) (res *SearchFindResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SearchFindResponse struct {
	// Pagination details for result navigation
	Pagination SearchFindResponsePagination `json:"pagination"`
	// Matched results with metadata
	Results []SearchFindResponseResult `json:"results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pagination  respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchFindResponse) RawJSON() string { return r.JSON.raw }
func (r *SearchFindResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination details for result navigation
type SearchFindResponsePagination struct {
	// Indicates more results available. Used for infinite scroll implementation
	HasMore bool `json:"has_more"`
	// Current page number (1-based)
	Page int64 `json:"page"`
	// Results per page. May be adjusted for performance
	PageSize int64 `json:"page_size"`
	// Total number of available results
	Total int64 `json:"total"`
	// Total available pages. Calculated as ceil(total/page_size)
	TotalPages int64 `json:"total_pages"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		Page        respjson.Field
		PageSize    respjson.Field
		Total       respjson.Field
		TotalPages  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchFindResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *SearchFindResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchFindResponseResult struct {
	// Unique identifier for this text segment. Used for deduplication and result
	// tracking
	ChunkSignature string `json:"chunk_signature,nullable"`
	// Vector representation for similarity matching. Used in semantic search
	// operations
	Embed string `json:"embed,nullable"`
	// Parent document identifier. Links related content chunks together
	PayloadSignature string `json:"payload_signature,nullable"`
	// Relevance score (0.0 to 1.0). Higher scores indicate better matches
	Score float64 `json:"score,nullable"`
	// Source document references. Contains bucket and object information
	Source SearchFindResponseResultSource `json:"source"`
	// The actual content of the result. May be a document excerpt or full content
	Text string `json:"text,nullable"`
	// Content MIME type. Helps with proper result rendering
	Type string `json:"type,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChunkSignature   respjson.Field
		Embed            respjson.Field
		PayloadSignature respjson.Field
		Score            respjson.Field
		Source           respjson.Field
		Text             respjson.Field
		Type             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchFindResponseResult) RawJSON() string { return r.JSON.raw }
func (r *SearchFindResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Source document references. Contains bucket and object information
type SearchFindResponseResultSource struct {
	// The bucket information containing this result
	Bucket SearchFindResponseResultSourceBucket `json:"bucket"`
	// The object key within the bucket
	Object string `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bucket      respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchFindResponseResultSource) RawJSON() string { return r.JSON.raw }
func (r *SearchFindResponseResultSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The bucket information containing this result
type SearchFindResponseResultSourceBucket struct {
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
func (r SearchFindResponseResultSourceBucket) RawJSON() string { return r.JSON.raw }
func (r *SearchFindResponseResultSourceBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchFindParams struct {
	// The buckets to search. If provided, the search will only return results from
	// these buckets
	BucketLocations []SearchFindParamsBucketLocationUnion `json:"bucket_locations,omitzero,required"`
	// Natural language search query that can include complex criteria. Supports
	// queries like finding documents with specific content types, PII, or semantic
	// meaning
	Input string `json:"input,required"`
	// Client-provided search session identifier. Required for pagination and result
	// tracking. We recommend using a UUID or ULID for this value
	RequestID string `json:"request_id,required"`
	paramObj
}

func (r SearchFindParams) MarshalJSON() (data []byte, err error) {
	type shadow SearchFindParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchFindParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SearchFindParamsBucketLocationUnion struct {
	OfBucket   *SearchFindParamsBucketLocationBucket   `json:",omitzero,inline"`
	OfModuleID *SearchFindParamsBucketLocationModuleID `json:",omitzero,inline"`
	paramUnion
}

func (u SearchFindParamsBucketLocationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[SearchFindParamsBucketLocationUnion](u.OfBucket, u.OfModuleID)
}
func (u *SearchFindParamsBucketLocationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SearchFindParamsBucketLocationUnion) asAny() any {
	if !param.IsOmitted(u.OfBucket) {
		return u.OfBucket
	} else if !param.IsOmitted(u.OfModuleID) {
		return u.OfModuleID
	}
	return nil
}

// The property Bucket is required.
type SearchFindParamsBucketLocationBucket struct {
	// BucketName represents a bucket name with an optional version
	Bucket SearchFindParamsBucketLocationBucketBucket `json:"bucket,omitzero,required"`
	paramObj
}

func (r SearchFindParamsBucketLocationBucket) MarshalJSON() (data []byte, err error) {
	type shadow SearchFindParamsBucketLocationBucket
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchFindParamsBucketLocationBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BucketName represents a bucket name with an optional version
type SearchFindParamsBucketLocationBucketBucket struct {
	// Optional Application
	ApplicationName param.Opt[string] `json:"application_name,omitzero"`
	// Optional version of the bucket
	Version param.Opt[string] `json:"version,omitzero"`
	// The name of the bucket
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r SearchFindParamsBucketLocationBucketBucket) MarshalJSON() (data []byte, err error) {
	type shadow SearchFindParamsBucketLocationBucketBucket
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchFindParamsBucketLocationBucketBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ModuleID is required.
type SearchFindParamsBucketLocationModuleID struct {
	ModuleID string `json:"module_id,required"`
	paramObj
}

func (r SearchFindParamsBucketLocationModuleID) MarshalJSON() (data []byte, err error) {
	type shadow SearchFindParamsBucketLocationModuleID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchFindParamsBucketLocationModuleID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
