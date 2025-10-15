// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package raindrop

import (
	"context"
	"net/http"
	"slices"

	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/apijson"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/requestconfig"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/option"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/param"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/respjson"
)

// EndSessionService contains methods and other services that help with interacting
// with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEndSessionService] method instead.
type EndSessionService struct {
	Options []option.RequestOption
}

// NewEndSessionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEndSessionService(opts ...option.RequestOption) (r EndSessionService) {
	r = EndSessionService{}
	r.Options = opts
	return
}

// Ends a working memory session, optionally flushing working memory to long-term
// storage. When flush is enabled, important memories are processed and stored for
// future retrieval.
func (r *EndSessionService) New(ctx context.Context, body EndSessionNewParams, opts ...option.RequestOption) (res *EndSessionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/end_session"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EndSessionNewResponse struct {
	// Indicates whether the session was ended successfully
	Success bool `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EndSessionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *EndSessionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EndSessionNewParams struct {
	// Unique session identifier to end
	SessionID string `json:"session_id,required"`
	// Smart memory locator for targeting the correct smart memory instance
	SmartMemoryLocation EndSessionNewParamsSmartMemoryLocationUnion `json:"smart_memory_location,omitzero,required"`
	// Whether to flush working memory to long-term storage
	Flush param.Opt[bool] `json:"flush,omitzero"`
	// Optional custom system prompt for memory summarization during flush
	SystemPrompt   param.Opt[string] `json:"system_prompt,omitzero"`
	OrganizationID param.Opt[string] `json:"organization_id,omitzero"`
	UserID         param.Opt[string] `json:"user_id,omitzero"`
	paramObj
}

func (r EndSessionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EndSessionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EndSessionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EndSessionNewParamsSmartMemoryLocationUnion struct {
	OfModuleID    *EndSessionNewParamsSmartMemoryLocationModuleID    `json:",omitzero,inline"`
	OfSmartMemory *EndSessionNewParamsSmartMemoryLocationSmartMemory `json:",omitzero,inline"`
	paramUnion
}

func (u EndSessionNewParamsSmartMemoryLocationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfModuleID, u.OfSmartMemory)
}
func (u *EndSessionNewParamsSmartMemoryLocationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EndSessionNewParamsSmartMemoryLocationUnion) asAny() any {
	if !param.IsOmitted(u.OfModuleID) {
		return u.OfModuleID
	} else if !param.IsOmitted(u.OfSmartMemory) {
		return u.OfSmartMemory
	}
	return nil
}

// The property ModuleID is required.
type EndSessionNewParamsSmartMemoryLocationModuleID struct {
	// **REQUIRED** FALSE
	ModuleID string `json:"module_id,required"`
	paramObj
}

func (r EndSessionNewParamsSmartMemoryLocationModuleID) MarshalJSON() (data []byte, err error) {
	type shadow EndSessionNewParamsSmartMemoryLocationModuleID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EndSessionNewParamsSmartMemoryLocationModuleID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SmartMemory is required.
type EndSessionNewParamsSmartMemoryLocationSmartMemory struct {
	// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
	// **REQUIRED** FALSE
	SmartMemory EndSessionNewParamsSmartMemoryLocationSmartMemorySmartMemory `json:"smart_memory,omitzero,required"`
	paramObj
}

func (r EndSessionNewParamsSmartMemoryLocationSmartMemory) MarshalJSON() (data []byte, err error) {
	type shadow EndSessionNewParamsSmartMemoryLocationSmartMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EndSessionNewParamsSmartMemoryLocationSmartMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
// **REQUIRED** FALSE
//
// The property Name is required.
type EndSessionNewParamsSmartMemoryLocationSmartMemorySmartMemory struct {
	// The name of the smart memory **EXAMPLE** "my-smartmemory" **REQUIRED** TRUE
	Name string `json:"name,required"`
	// Optional Application **EXAMPLE** "my-app" **REQUIRED** FALSE
	ApplicationName param.Opt[string] `json:"application_name,omitzero"`
	// Optional version of the smart memory **EXAMPLE** "01jtryx2f2f61ryk06vd8mr91p"
	// **REQUIRED** FALSE
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r EndSessionNewParamsSmartMemoryLocationSmartMemorySmartMemory) MarshalJSON() (data []byte, err error) {
	type shadow EndSessionNewParamsSmartMemoryLocationSmartMemorySmartMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EndSessionNewParamsSmartMemoryLocationSmartMemorySmartMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
