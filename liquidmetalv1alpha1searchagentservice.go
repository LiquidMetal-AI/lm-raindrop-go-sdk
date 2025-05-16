// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package raindrop

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/raindrop-go/internal/apijson"
	"github.com/stainless-sdks/raindrop-go/internal/requestconfig"
	"github.com/stainless-sdks/raindrop-go/option"
	"github.com/stainless-sdks/raindrop-go/packages/param"
	"github.com/stainless-sdks/raindrop-go/packages/respjson"
)

// LiquidmetalV1alpha1SearchAgentServiceService contains methods and other services
// that help with interacting with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLiquidmetalV1alpha1SearchAgentServiceService] method instead.
type LiquidmetalV1alpha1SearchAgentServiceService struct {
	Options []option.RequestOption
}

// NewLiquidmetalV1alpha1SearchAgentServiceService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewLiquidmetalV1alpha1SearchAgentServiceService(opts ...option.RequestOption) (r LiquidmetalV1alpha1SearchAgentServiceService) {
	r = LiquidmetalV1alpha1SearchAgentServiceService{}
	r.Options = opts
	return
}

// Retrieve additional pages from a previous search. This endpoint enables
// navigation through large result sets while maintaining search context and result
// relevance. Retrieving paginated results requires a valid request_id from a
// previously completed search.
func (r *LiquidmetalV1alpha1SearchAgentServiceService) GetPaginatedResults(ctx context.Context, params LiquidmetalV1alpha1SearchAgentServiceGetPaginatedResultsParams, opts ...option.RequestOption) (res *LiquidmetalV1alpha1SearchAgentServiceGetPaginatedResultsResponse, err error) {
	if !param.IsOmitted(params.ConnectProtocolVersion) {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if !param.IsOmitted(params.ConnectTimeoutMs) {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs.Value)))
	}
	opts = append(r.Options[:], opts...)
	path := "liquidmetal.v1alpha1.SearchAgentService/GetPaginatedResults"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type LiquidmetalV1alpha1SearchAgentServiceGetPaginatedResultsResponse struct {
	// **DESCRIPTION** Updated pagination information **EXAMPLE** {"total": 100,
	// "page": 2, "page_size": 10, "total_pages": 10, "has_more": true}
	Pagination PaginationInfo `json:"pagination"`
	// **DESCRIPTION** Page results with full metadata **EXAMPLE** [{"chunk_signature":
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
func (r LiquidmetalV1alpha1SearchAgentServiceGetPaginatedResultsResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *LiquidmetalV1alpha1SearchAgentServiceGetPaginatedResultsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LiquidmetalV1alpha1SearchAgentServiceGetPaginatedResultsParams struct {
	// **DESCRIPTION** Requested page number **EXAMPLE** 2 **REQUIRED** TRUE
	Page param.Opt[int64] `json:"page,omitzero"`
	// **DESCRIPTION** Results per page **EXAMPLE** 10 **REQUIRED** TRUE
	PageSize       param.Opt[int64]  `json:"page_size,omitzero"`
	OrganizationID param.Opt[string] `json:"organization_id,omitzero"`
	// **DESCRIPTION** Original search session identifier from the initial search
	// **EXAMPLE** "123e4567-e89b-12d3-a456-426614174000" **REQUIRED** TRUE
	RequestID param.Opt[string] `json:"request_id,omitzero"`
	UserID    param.Opt[string] `json:"user_id,omitzero"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Opt[float64] `header:"Connect-Timeout-Ms,omitzero" json:"-"`
	// Define the version of the Connect protocol
	//
	// This field can be elided, and will marshal its zero value as 1.
	ConnectProtocolVersion float64 `header:"Connect-Protocol-Version,omitzero,required" json:"-"`
	paramObj
}

func (r LiquidmetalV1alpha1SearchAgentServiceGetPaginatedResultsParams) MarshalJSON() (data []byte, err error) {
	type shadow LiquidmetalV1alpha1SearchAgentServiceGetPaginatedResultsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LiquidmetalV1alpha1SearchAgentServiceGetPaginatedResultsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
