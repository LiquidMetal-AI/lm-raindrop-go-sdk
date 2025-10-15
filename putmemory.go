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
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/shared"
)

// PutMemoryService contains methods and other services that help with interacting
// with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPutMemoryService] method instead.
type PutMemoryService struct {
	Options []option.RequestOption
}

// NewPutMemoryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPutMemoryService(opts ...option.RequestOption) (r PutMemoryService) {
	r = PutMemoryService{}
	r.Options = opts
	return
}

// Stores a new memory entry in the agent's working memory. Memories are organized
// by timeline and can include contextual information like the agent responsible
// and triggering events.
//
// The system will:
//
// - Store the memory with automatic timestamping
// - Generate embeddings for semantic search
// - Associate the memory with the specified timeline
// - Enable future retrieval and search operations
func (r *PutMemoryService) New(ctx context.Context, body PutMemoryNewParams, opts ...option.RequestOption) (res *PutMemoryNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/put_memory"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type PutMemoryNewResponse struct {
	// Unique identifier for the stored memory entry
	MemoryID string `json:"memory_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MemoryID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PutMemoryNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PutMemoryNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PutMemoryNewParams struct {
	// The actual memory content to store
	Content string `json:"content,required"`
	// Unique session identifier for the working memory instance
	SessionID string `json:"session_id,required"`
	// Smart memory locator for targeting the correct smart memory instance
	SmartMemoryLocation PutMemoryNewParamsSmartMemoryLocationUnion `json:"smart_memory_location,omitzero,required"`
	// Agent identifier responsible for this memory
	Agent param.Opt[string] `json:"agent,omitzero"`
	// Optional key for direct memory retrieval
	Key param.Opt[string] `json:"key,omitzero"`
	// Timeline identifier for organizing related memories
	Timeline       param.Opt[string] `json:"timeline,omitzero"`
	OrganizationID param.Opt[string] `json:"organization_id,omitzero"`
	UserID         param.Opt[string] `json:"user_id,omitzero"`
	paramObj
}

func (r PutMemoryNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PutMemoryNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PutMemoryNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PutMemoryNewParamsSmartMemoryLocationUnion struct {
	OfModuleID    *PutMemoryNewParamsSmartMemoryLocationModuleID    `json:",omitzero,inline"`
	OfSmartMemory *PutMemoryNewParamsSmartMemoryLocationSmartMemory `json:",omitzero,inline"`
	paramUnion
}

func (u PutMemoryNewParamsSmartMemoryLocationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfModuleID, u.OfSmartMemory)
}
func (u *PutMemoryNewParamsSmartMemoryLocationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PutMemoryNewParamsSmartMemoryLocationUnion) asAny() any {
	if !param.IsOmitted(u.OfModuleID) {
		return u.OfModuleID
	} else if !param.IsOmitted(u.OfSmartMemory) {
		return u.OfSmartMemory
	}
	return nil
}

// The property ModuleID is required.
type PutMemoryNewParamsSmartMemoryLocationModuleID struct {
	// **REQUIRED** FALSE
	ModuleID string `json:"module_id,required"`
	paramObj
}

func (r PutMemoryNewParamsSmartMemoryLocationModuleID) MarshalJSON() (data []byte, err error) {
	type shadow PutMemoryNewParamsSmartMemoryLocationModuleID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PutMemoryNewParamsSmartMemoryLocationModuleID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SmartMemory is required.
type PutMemoryNewParamsSmartMemoryLocationSmartMemory struct {
	// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
	// **REQUIRED** FALSE
	SmartMemory shared.LiquidmetalV1alpha1SmartMemoryNameParam `json:"smart_memory,omitzero,required"`
	paramObj
}

func (r PutMemoryNewParamsSmartMemoryLocationSmartMemory) MarshalJSON() (data []byte, err error) {
	type shadow PutMemoryNewParamsSmartMemoryLocationSmartMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PutMemoryNewParamsSmartMemoryLocationSmartMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
