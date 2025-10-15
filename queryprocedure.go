// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package raindrop

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/apijson"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/requestconfig"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/option"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/param"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/respjson"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/shared"
)

// QueryProcedureService contains methods and other services that help with
// interacting with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQueryProcedureService] method instead.
type QueryProcedureService struct {
	Options []option.RequestOption
}

// NewQueryProcedureService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewQueryProcedureService(opts ...option.RequestOption) (r QueryProcedureService) {
	r = QueryProcedureService{}
	r.Options = opts
	return
}

// Searches procedures using text matching across keys and values. Supports
// filtering by procedure keys, values, or both with fuzzy matching and relevance
// scoring.
//
// TODO: Future enhancement will include vector search for semantic similarity.
func (r *QueryProcedureService) Search(ctx context.Context, body QueryProcedureSearchParams, opts ...option.RequestOption) (res *QueryProcedureSearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/search_procedures"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type QueryProcedureSearchResponse struct {
	// List of matching procedures ordered by relevance
	Procedures []QueryProcedureSearchResponseProcedure `json:"procedures"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Procedures  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryProcedureSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *QueryProcedureSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryProcedureSearchResponseProcedure struct {
	// When this procedure was first created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Unique key for this procedure
	Key string `json:"key"`
	// When this procedure was last updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The procedure content
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Key         respjson.Field
		UpdatedAt   respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryProcedureSearchResponseProcedure) RawJSON() string { return r.JSON.raw }
func (r *QueryProcedureSearchResponseProcedure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryProcedureSearchParams struct {
	// Smart memory locator for targeting the correct smart memory instance
	SmartMemoryLocation QueryProcedureSearchParamsSmartMemoryLocationUnion `json:"smart_memory_location,omitzero,required"`
	// Search terms to match against procedure keys and values
	Terms string `json:"terms,required"`
	// Maximum number of results to return
	NMostRecent param.Opt[int64] `json:"n_most_recent,omitzero"`
	// Optional procedural memory ID to use for actor isolation
	ProceduralMemoryID param.Opt[string] `json:"procedural_memory_id,omitzero"`
	// Whether to search in procedure keys
	SearchKeys param.Opt[bool] `json:"search_keys,omitzero"`
	// Whether to search in procedure values
	SearchValues   param.Opt[bool]   `json:"search_values,omitzero"`
	OrganizationID param.Opt[string] `json:"organization_id,omitzero"`
	UserID         param.Opt[string] `json:"user_id,omitzero"`
	paramObj
}

func (r QueryProcedureSearchParams) MarshalJSON() (data []byte, err error) {
	type shadow QueryProcedureSearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QueryProcedureSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type QueryProcedureSearchParamsSmartMemoryLocationUnion struct {
	OfModuleID    *QueryProcedureSearchParamsSmartMemoryLocationModuleID    `json:",omitzero,inline"`
	OfSmartMemory *QueryProcedureSearchParamsSmartMemoryLocationSmartMemory `json:",omitzero,inline"`
	paramUnion
}

func (u QueryProcedureSearchParamsSmartMemoryLocationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfModuleID, u.OfSmartMemory)
}
func (u *QueryProcedureSearchParamsSmartMemoryLocationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *QueryProcedureSearchParamsSmartMemoryLocationUnion) asAny() any {
	if !param.IsOmitted(u.OfModuleID) {
		return u.OfModuleID
	} else if !param.IsOmitted(u.OfSmartMemory) {
		return u.OfSmartMemory
	}
	return nil
}

// The property ModuleID is required.
type QueryProcedureSearchParamsSmartMemoryLocationModuleID struct {
	// **REQUIRED** FALSE
	ModuleID string `json:"module_id,required"`
	paramObj
}

func (r QueryProcedureSearchParamsSmartMemoryLocationModuleID) MarshalJSON() (data []byte, err error) {
	type shadow QueryProcedureSearchParamsSmartMemoryLocationModuleID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QueryProcedureSearchParamsSmartMemoryLocationModuleID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SmartMemory is required.
type QueryProcedureSearchParamsSmartMemoryLocationSmartMemory struct {
	// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
	// **REQUIRED** FALSE
	SmartMemory shared.LiquidmetalV1alpha1SmartMemoryNameParam `json:"smart_memory,omitzero,required"`
	paramObj
}

func (r QueryProcedureSearchParamsSmartMemoryLocationSmartMemory) MarshalJSON() (data []byte, err error) {
	type shadow QueryProcedureSearchParamsSmartMemoryLocationSmartMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QueryProcedureSearchParamsSmartMemoryLocationSmartMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
