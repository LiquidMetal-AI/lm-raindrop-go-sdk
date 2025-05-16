// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package raindrop

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/raindrop-go/internal/apijson"
	"github.com/stainless-sdks/raindrop-go/internal/requestconfig"
	"github.com/stainless-sdks/raindrop-go/option"
	"github.com/stainless-sdks/raindrop-go/packages/param"
	"github.com/stainless-sdks/raindrop-go/packages/respjson"
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
	Results []TextResult `json:"results"`
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

type SearchFindParams struct {
	// The buckets to search. If provided, the search will only return results from
	// these buckets
	BucketLocations []BucketLocatorUnionParam `json:"bucket_locations,omitzero,required"`
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
