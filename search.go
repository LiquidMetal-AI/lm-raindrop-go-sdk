// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package raindrop

import (
	"context"
	"net/http"
	"net/url"

	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/apijson"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/apiquery"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/requestconfig"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/option"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/pagination"
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

// Retrieve additional pages from a previous search. This endpoint enables
// navigation through large result sets while maintaining search context and result
// relevance. Retrieving paginated results requires a valid `request_id` from a
// previously completed search.
func (r *SearchService) Get(ctx context.Context, query SearchGetParams, opts ...option.RequestOption) (res *pagination.SearchPage[TextResult], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/search"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieve additional pages from a previous search. This endpoint enables
// navigation through large result sets while maintaining search context and result
// relevance. Retrieving paginated results requires a valid `request_id` from a
// previously completed search.
func (r *SearchService) GetAutoPaging(ctx context.Context, query SearchGetParams, opts ...option.RequestOption) *pagination.SearchPageAutoPager[TextResult] {
	return pagination.NewSearchPageAutoPager(r.Get(ctx, query, opts...))
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
//   - 'Find me all audio files that contain infomration about the weather in SF in
//     2024'
//
// Key capabilities:
//
// - Natural language query understanding
// - Content-based search across text, images, and audio
// - Automatic PII detection
// - Multi-modal search (text, images, audio)
func (r *SearchService) Find(ctx context.Context, body SearchFindParams, opts ...option.RequestOption) (res *SearchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SearchResponse struct {
	Pagination SearchResponsePagination `json:"pagination,required"`
	// Matched results with metadata
	Results []TextResult `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pagination  respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchResponse) RawJSON() string { return r.JSON.raw }
func (r *SearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchResponsePagination struct {
	// Indicates more results available
	HasMore bool `json:"has_more,required"`
	// Current page number (1-based)
	Page int64 `json:"page,required"`
	// Results per page
	PageSize int64 `json:"page_size,required"`
	// Total number of available results
	Total int64 `json:"total,required"`
	// Total available pages
	TotalPages int64 `json:"total_pages,required"`
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
func (r SearchResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *SearchResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TextResult struct {
	// Unique identifier for this text segment
	ChunkSignature string `json:"chunk_signature,required"`
	// Parent document identifier
	PayloadSignature string `json:"payload_signature"`
	// Relevance score (0.0 to 1.0)
	Score float64 `json:"score"`
	// Source document information in JSON format
	Source string `json:"source"`
	// The actual content of the result
	Text string `json:"text"`
	// Content MIME type
	//
	// Any of "text/plain", "application/pdf", "image/jpeg", "image/png".
	Type TextResultType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChunkSignature   respjson.Field
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
func (r TextResult) RawJSON() string { return r.JSON.raw }
func (r *TextResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Content MIME type
type TextResultType string

const (
	TextResultTypeTextPlain      TextResultType = "text/plain"
	TextResultTypeApplicationPdf TextResultType = "application/pdf"
	TextResultTypeImageJpeg      TextResultType = "image/jpeg"
	TextResultTypeImagePng       TextResultType = "image/png"
)

type SearchGetParams struct {
	// Client-provided search session identifier from the initial search
	RequestID string `query:"request_id,required" json:"-"`
	// Requested page number
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Results per page
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SearchGetParams]'s query parameters as `url.Values`.
func (r SearchGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SearchFindParams struct {
	BucketLocations []SearchFindParamsBucketLocationUnion `json:"bucket_locations,omitzero,required"`
	// Natural language search query that can include complex criteria
	Input string `json:"input,required"`
	// Client-provided search session identifier. Required for pagination and result
	// tracking. We recommend using a UUID or ULID for this value.
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
	OfSearchFindsBucketLocationModuleID *SearchFindParamsBucketLocationModuleID `json:",omitzero,inline"`
	OfSearchFindsBucketLocationBucket   *SearchFindParamsBucketLocationBucket   `json:",omitzero,inline"`
	paramUnion
}

func (u SearchFindParamsBucketLocationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[SearchFindParamsBucketLocationUnion](u.OfSearchFindsBucketLocationModuleID, u.OfSearchFindsBucketLocationBucket)
}
func (u *SearchFindParamsBucketLocationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SearchFindParamsBucketLocationUnion) asAny() any {
	if !param.IsOmitted(u.OfSearchFindsBucketLocationModuleID) {
		return u.OfSearchFindsBucketLocationModuleID
	} else if !param.IsOmitted(u.OfSearchFindsBucketLocationBucket) {
		return u.OfSearchFindsBucketLocationBucket
	}
	return nil
}

// The property ModuleID is required.
type SearchFindParamsBucketLocationModuleID struct {
	// Version-agnostic identifier for a module
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

// The property Bucket is required.
type SearchFindParamsBucketLocationBucket struct {
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

// The properties ApplicationName, Name, Version are required.
type SearchFindParamsBucketLocationBucketBucket struct {
	// Name of the application
	ApplicationName string `json:"application_name,required"`
	// Name of the bucket
	Name string `json:"name,required"`
	// Version of the bucket
	Version string `json:"version,required"`
	paramObj
}

func (r SearchFindParamsBucketLocationBucketBucket) MarshalJSON() (data []byte, err error) {
	type shadow SearchFindParamsBucketLocationBucketBucket
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchFindParamsBucketLocationBucketBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
