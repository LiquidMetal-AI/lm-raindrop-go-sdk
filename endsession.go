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
	opts = append(r.Options[:], opts...)
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
	// Agent memory locator for targeting the correct agent memory instance
	AgentMemoryLocation EndSessionNewParamsAgentMemoryLocationUnion `json:"agentMemoryLocation,omitzero,required"`
	// Unique session identifier to end
	SessionID string `json:"sessionId,required"`
	// Whether to flush working memory to long-term storage
	Flush param.Opt[bool] `json:"flush,omitzero"`
	// Optional custom system prompt for memory summarization during flush
	SystemPrompt param.Opt[string] `json:"systemPrompt,omitzero"`
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
type EndSessionNewParamsAgentMemoryLocationUnion struct {
	OfAgentMemory *EndSessionNewParamsAgentMemoryLocationAgentMemory `json:",omitzero,inline"`
	OfModuleID    *EndSessionNewParamsAgentMemoryLocationModuleID    `json:",omitzero,inline"`
	paramUnion
}

func (u EndSessionNewParamsAgentMemoryLocationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAgentMemory, u.OfModuleID)
}
func (u *EndSessionNewParamsAgentMemoryLocationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EndSessionNewParamsAgentMemoryLocationUnion) asAny() any {
	if !param.IsOmitted(u.OfAgentMemory) {
		return u.OfAgentMemory
	} else if !param.IsOmitted(u.OfModuleID) {
		return u.OfModuleID
	}
	return nil
}

// The property AgentMemory is required.
type EndSessionNewParamsAgentMemoryLocationAgentMemory struct {
	// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
	// **REQUIRED** FALSE
	AgentMemory EndSessionNewParamsAgentMemoryLocationAgentMemoryAgentMemory `json:"agentMemory,omitzero,required"`
	paramObj
}

func (r EndSessionNewParamsAgentMemoryLocationAgentMemory) MarshalJSON() (data []byte, err error) {
	type shadow EndSessionNewParamsAgentMemoryLocationAgentMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EndSessionNewParamsAgentMemoryLocationAgentMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
// **REQUIRED** FALSE
//
// The property Name is required.
type EndSessionNewParamsAgentMemoryLocationAgentMemoryAgentMemory struct {
	// The name of the agent memory **EXAMPLE** "my-agent-memory" **REQUIRED** TRUE
	Name string `json:"name,required"`
	// Optional Application **EXAMPLE** "my-app" **REQUIRED** FALSE
	ApplicationName param.Opt[string] `json:"applicationName,omitzero"`
	// Optional version of the agent memory **EXAMPLE** "01jtryx2f2f61ryk06vd8mr91p"
	// **REQUIRED** FALSE
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r EndSessionNewParamsAgentMemoryLocationAgentMemoryAgentMemory) MarshalJSON() (data []byte, err error) {
	type shadow EndSessionNewParamsAgentMemoryLocationAgentMemoryAgentMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EndSessionNewParamsAgentMemoryLocationAgentMemoryAgentMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ModuleID is required.
type EndSessionNewParamsAgentMemoryLocationModuleID struct {
	ModuleID string `json:"moduleId,required"`
	paramObj
}

func (r EndSessionNewParamsAgentMemoryLocationModuleID) MarshalJSON() (data []byte, err error) {
	type shadow EndSessionNewParamsAgentMemoryLocationModuleID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EndSessionNewParamsAgentMemoryLocationModuleID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
