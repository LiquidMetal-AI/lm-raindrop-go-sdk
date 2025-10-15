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

// PutSemanticMemoryService contains methods and other services that help with
// interacting with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPutSemanticMemoryService] method instead.
type PutSemanticMemoryService struct {
	Options []option.RequestOption
}

// NewPutSemanticMemoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPutSemanticMemoryService(opts ...option.RequestOption) (r PutSemanticMemoryService) {
	r = PutSemanticMemoryService{}
	r.Options = opts
	return
}

// Stores a semantic memory document for long-term knowledge retrieval. Semantic
// memory is used for storing structured knowledge, facts, and information that can
// be searched and retrieved across different sessions.
func (r *PutSemanticMemoryService) New(ctx context.Context, body PutSemanticMemoryNewParams, opts ...option.RequestOption) (res *PutSemanticMemoryNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/put_semantic_memory"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type PutSemanticMemoryNewResponse struct {
	// Error message if the operation failed
	Error string `json:"error,nullable"`
	// Unique object identifier for the stored document
	ObjectID string `json:"object_id,nullable"`
	// Indicates whether the document was stored successfully
	Success bool `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		ObjectID    respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PutSemanticMemoryNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PutSemanticMemoryNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PutSemanticMemoryNewParams struct {
	// JSON-encoded document content to store in semantic memory
	Document string `json:"document,required"`
	// Smart memory locator for targeting the correct smart memory instance
	SmartMemoryLocation PutSemanticMemoryNewParamsSmartMemoryLocationUnion `json:"smart_memory_location,omitzero,required"`
	OrganizationID      param.Opt[string]                                  `json:"organization_id,omitzero"`
	UserID              param.Opt[string]                                  `json:"user_id,omitzero"`
	paramObj
}

func (r PutSemanticMemoryNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PutSemanticMemoryNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PutSemanticMemoryNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PutSemanticMemoryNewParamsSmartMemoryLocationUnion struct {
	OfModuleID    *PutSemanticMemoryNewParamsSmartMemoryLocationModuleID    `json:",omitzero,inline"`
	OfSmartMemory *PutSemanticMemoryNewParamsSmartMemoryLocationSmartMemory `json:",omitzero,inline"`
	paramUnion
}

func (u PutSemanticMemoryNewParamsSmartMemoryLocationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfModuleID, u.OfSmartMemory)
}
func (u *PutSemanticMemoryNewParamsSmartMemoryLocationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PutSemanticMemoryNewParamsSmartMemoryLocationUnion) asAny() any {
	if !param.IsOmitted(u.OfModuleID) {
		return u.OfModuleID
	} else if !param.IsOmitted(u.OfSmartMemory) {
		return u.OfSmartMemory
	}
	return nil
}

// The property ModuleID is required.
type PutSemanticMemoryNewParamsSmartMemoryLocationModuleID struct {
	// **REQUIRED** FALSE
	ModuleID string `json:"module_id,required"`
	paramObj
}

func (r PutSemanticMemoryNewParamsSmartMemoryLocationModuleID) MarshalJSON() (data []byte, err error) {
	type shadow PutSemanticMemoryNewParamsSmartMemoryLocationModuleID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PutSemanticMemoryNewParamsSmartMemoryLocationModuleID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SmartMemory is required.
type PutSemanticMemoryNewParamsSmartMemoryLocationSmartMemory struct {
	// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
	// **REQUIRED** FALSE
	SmartMemory shared.LiquidmetalV1alpha1SmartMemoryNameParam `json:"smart_memory,omitzero,required"`
	paramObj
}

func (r PutSemanticMemoryNewParamsSmartMemoryLocationSmartMemory) MarshalJSON() (data []byte, err error) {
	type shadow PutSemanticMemoryNewParamsSmartMemoryLocationSmartMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PutSemanticMemoryNewParamsSmartMemoryLocationSmartMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
