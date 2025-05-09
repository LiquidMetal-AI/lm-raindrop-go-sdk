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
func (r *SummarizePageService) New(ctx context.Context, body SummarizePageNewParams, opts ...option.RequestOption) (res *SummarizePageNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/summarize_page"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SummarizePageNewResponse struct {
	// AI-generated summary including key themes and topics, content type distribution,
	// important findings, and document relationships
	Summary string `json:"summary,required"`
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
	// Client-provided search session identifier from the original search
	RequestID string `json:"request_id,required"`
	// Target page number (1-based)
	Page param.Opt[int64] `json:"page,omitzero"`
	// Results per page. Affects how many documents are included in the summary
	PageSize param.Opt[int64] `json:"page_size,omitzero"`
	paramObj
}

func (r SummarizePageNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SummarizePageNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SummarizePageNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
