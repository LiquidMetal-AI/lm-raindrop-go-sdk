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
func (r *SummarizePageService) NewSummary(ctx context.Context, body SummarizePageNewSummaryParams, opts ...option.RequestOption) (res *SummarizePageNewSummaryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/summarize_page"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SummarizePageNewSummaryResponse struct {
	// **DESCRIPTION** AI-generated summary including key themes and topics, content
	// type distribution, important findings, and document relationships **EXAMPLE**
	// "The search results contain information about..."
	Summary string `json:"summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SummarizePageNewSummaryResponse) RawJSON() string { return r.JSON.raw }
func (r *SummarizePageNewSummaryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SummarizePageNewSummaryParams struct {
	OrganizationID param.Opt[string] `json:"organization_id,omitzero"`
	// **DESCRIPTION** Target page number (1-based) **EXAMPLE** 1 **REQUIRED** TRUE
	Page param.Opt[int64] `json:"page,omitzero"`
	// **DESCRIPTION** Results per page. Affects summary granularity **EXAMPLE** 10
	// **REQUIRED** TRUE
	PageSize param.Opt[int64] `json:"page_size,omitzero"`
	// **DESCRIPTION** Original search session identifier from the initial search
	// **EXAMPLE** "123e4567-e89b-12d3-a456-426614174000" **REQUIRED** TRUE
	RequestID param.Opt[string] `json:"request_id,omitzero"`
	UserID    param.Opt[string] `json:"user_id,omitzero"`
	paramObj
}

func (r SummarizePageNewSummaryParams) MarshalJSON() (data []byte, err error) {
	type shadow SummarizePageNewSummaryParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SummarizePageNewSummaryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
