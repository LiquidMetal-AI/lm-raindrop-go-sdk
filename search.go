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
func (r *SearchService) Run(ctx context.Context, body SearchRunParams, opts ...option.RequestOption) (res *SearchRunResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SearchRunResponse struct {
	// **DESCRIPTION** Pagination details for result navigation **EXAMPLE** {"total":
	// 100, "page": 1, "page_size": 10, "total_pages": 10, "has_more": true}
	Pagination SearchRunResponsePagination `json:"pagination"`
	// **DESCRIPTION** Matched results with metadata **EXAMPLE** [{"chunk_signature":
	// "chunk_123abc", "text": "Sample text", "score": 0.95}]
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
func (r SearchRunResponse) RawJSON() string { return r.JSON.raw }
func (r *SearchRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **DESCRIPTION** Pagination details for result navigation **EXAMPLE** {"total":
// 100, "page": 1, "page_size": 10, "total_pages": 10, "has_more": true}
type SearchRunResponsePagination struct {
	// **DESCRIPTION** Indicates more results available. Used for infinite scroll
	// implementation **EXAMPLE** true
	HasMore bool `json:"has_more"`
	// **DESCRIPTION** Current page number (1-based) **EXAMPLE** 1
	Page int64 `json:"page"`
	// **DESCRIPTION** Results per page. May be adjusted for performance **EXAMPLE** 15
	PageSize int64 `json:"page_size"`
	// **DESCRIPTION** Total number of available results **EXAMPLE** 1020
	Total int64 `json:"total"`
	// **DESCRIPTION** Total available pages. Calculated as ceil(total/page_size)
	// **EXAMPLE** 68
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
func (r SearchRunResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *SearchRunResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchRunParams struct {
	// **DESCRIPTION** Natural language search query that can include complex criteria.
	// Supports queries like finding documents with specific content types, PII, or
	// semantic meaning **EXAMPLE** "Show me documents containing credit card numbers
	// or social security numbers" **REQUIRED** TRUE
	Input          param.Opt[string] `json:"input,omitzero"`
	OrganizationID param.Opt[string] `json:"organization_id,omitzero"`
	// **DESCRIPTION** Client-provided search session identifier. Required for
	// pagination and result tracking. We recommend using a UUID or ULID for this value
	// **EXAMPLE** "123e4567-e89b-12d3-a456-426614174000" **REQUIRED** TRUE
	RequestID param.Opt[string] `json:"request_id,omitzero"`
	UserID    param.Opt[string] `json:"user_id,omitzero"`
	// **DESCRIPTION** The buckets to search. If provided, the search will only return
	// results from these buckets **EXAMPLE** [{"bucket": {"name": "my-bucket",
	// "version": "01jtgtraw3b5qbahrhvrj3ygbb", "application_name": "my-app"}}]
	// **REQUIRED** TRUE
	BucketLocations []BucketLocatorUnionParam `json:"bucket_locations,omitzero"`
	paramObj
}

func (r SearchRunParams) MarshalJSON() (data []byte, err error) {
	type shadow SearchRunParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchRunParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
