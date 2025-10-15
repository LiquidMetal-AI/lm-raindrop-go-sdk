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

// GetSemanticMemoryService contains methods and other services that help with
// interacting with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGetSemanticMemoryService] method instead.
type GetSemanticMemoryService struct {
	Options []option.RequestOption
}

// NewGetSemanticMemoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGetSemanticMemoryService(opts ...option.RequestOption) (r GetSemanticMemoryService) {
	r = GetSemanticMemoryService{}
	r.Options = opts
	return
}

// Retrieves a specific semantic memory document by its object ID. Returns the
// complete document with all its stored properties and metadata.
func (r *GetSemanticMemoryService) New(ctx context.Context, body GetSemanticMemoryNewParams, opts ...option.RequestOption) (res *GetSemanticMemoryNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/get_semantic_memory"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type GetSemanticMemoryNewResponse struct {
	// JSON-encoded document content if found
	Document string `json:"document,nullable"`
	// Error message if the operation failed
	Error string `json:"error,nullable"`
	// Indicates whether the document was retrieved successfully
	Success bool `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Document    respjson.Field
		Error       respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetSemanticMemoryNewResponse) RawJSON() string { return r.JSON.raw }
func (r *GetSemanticMemoryNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetSemanticMemoryNewParams struct {
	// Unique object identifier of the document to retrieve
	ObjectID string `json:"object_id,required"`
	// Smart memory locator for targeting the correct smart memory instance
	SmartMemoryLocation GetSemanticMemoryNewParamsSmartMemoryLocationUnion `json:"smart_memory_location,omitzero,required"`
	OrganizationID      param.Opt[string]                                  `json:"organization_id,omitzero"`
	UserID              param.Opt[string]                                  `json:"user_id,omitzero"`
	paramObj
}

func (r GetSemanticMemoryNewParams) MarshalJSON() (data []byte, err error) {
	type shadow GetSemanticMemoryNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GetSemanticMemoryNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type GetSemanticMemoryNewParamsSmartMemoryLocationUnion struct {
	OfModuleID    *GetSemanticMemoryNewParamsSmartMemoryLocationModuleID    `json:",omitzero,inline"`
	OfSmartMemory *GetSemanticMemoryNewParamsSmartMemoryLocationSmartMemory `json:",omitzero,inline"`
	paramUnion
}

func (u GetSemanticMemoryNewParamsSmartMemoryLocationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfModuleID, u.OfSmartMemory)
}
func (u *GetSemanticMemoryNewParamsSmartMemoryLocationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *GetSemanticMemoryNewParamsSmartMemoryLocationUnion) asAny() any {
	if !param.IsOmitted(u.OfModuleID) {
		return u.OfModuleID
	} else if !param.IsOmitted(u.OfSmartMemory) {
		return u.OfSmartMemory
	}
	return nil
}

// The property ModuleID is required.
type GetSemanticMemoryNewParamsSmartMemoryLocationModuleID struct {
	// **REQUIRED** FALSE
	ModuleID string `json:"module_id,required"`
	paramObj
}

func (r GetSemanticMemoryNewParamsSmartMemoryLocationModuleID) MarshalJSON() (data []byte, err error) {
	type shadow GetSemanticMemoryNewParamsSmartMemoryLocationModuleID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GetSemanticMemoryNewParamsSmartMemoryLocationModuleID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SmartMemory is required.
type GetSemanticMemoryNewParamsSmartMemoryLocationSmartMemory struct {
	// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
	// **REQUIRED** FALSE
	SmartMemory GetSemanticMemoryNewParamsSmartMemoryLocationSmartMemorySmartMemory `json:"smart_memory,omitzero,required"`
	paramObj
}

func (r GetSemanticMemoryNewParamsSmartMemoryLocationSmartMemory) MarshalJSON() (data []byte, err error) {
	type shadow GetSemanticMemoryNewParamsSmartMemoryLocationSmartMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GetSemanticMemoryNewParamsSmartMemoryLocationSmartMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
// **REQUIRED** FALSE
//
// The property Name is required.
type GetSemanticMemoryNewParamsSmartMemoryLocationSmartMemorySmartMemory struct {
	// The name of the smart memory **EXAMPLE** "my-smartmemory" **REQUIRED** TRUE
	Name string `json:"name,required"`
	// Optional Application **EXAMPLE** "my-app" **REQUIRED** FALSE
	ApplicationName param.Opt[string] `json:"application_name,omitzero"`
	// Optional version of the smart memory **EXAMPLE** "01jtryx2f2f61ryk06vd8mr91p"
	// **REQUIRED** FALSE
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r GetSemanticMemoryNewParamsSmartMemoryLocationSmartMemorySmartMemory) MarshalJSON() (data []byte, err error) {
	type shadow GetSemanticMemoryNewParamsSmartMemoryLocationSmartMemorySmartMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GetSemanticMemoryNewParamsSmartMemoryLocationSmartMemorySmartMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
