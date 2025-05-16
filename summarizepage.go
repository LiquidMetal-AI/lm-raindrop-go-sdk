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

// SummarizePageService contains methods and other services that help with
// interacting with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSummarizePageService] method instead.
type SummarizePageService struct {
	Options []option.RequestOption
}

// NewSummarizePageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSummarizePageService(opts ...option.RequestOption) (r SummarizePageService) {
	r = SummarizePageService{}
	r.Options = opts
	return
}

// Generates intelligent summaries of search result pages, helping users quickly
// understand large result sets without reading through every document. The system
// analyzes the content of all results on a given page and generates a detailed
// overview.
//
// The summary system:
//
// - Identifies key themes and topics
// - Extracts important findings
// - Highlights document relationships
// - Provides content type distribution
// - Summarizes metadata patterns
//
// This is particularly valuable when dealing with:
//
// - Large document collections
// - Mixed content types
// - Technical documentation
// - Research materials
func (r *SummarizePageService) New(ctx context.Context, body SummarizePageNewParams, opts ...option.RequestOption) (res *SummarizePageNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/summarize_page"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SummarizePageNewResponse struct {
	// AI-generated summary including key themes and topics, content type distribution,
	// important findings, and document relationships
	Summary string `json:"summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SummarizePageNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SummarizePageNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SummarizePageNewParams struct {
	// Target page number (1-based)
	Page param.Opt[int64] `json:"page,omitzero"`
	// Results per page. Affects summary granularity
	PageSize param.Opt[int64] `json:"page_size,omitzero"`
	// Original search session identifier from the initial search
	RequestID param.Opt[string] `json:"request_id,omitzero"`
	paramObj
}

func (r SummarizePageNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SummarizePageNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SummarizePageNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
