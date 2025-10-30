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

// GetProcedureService contains methods and other services that help with
// interacting with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGetProcedureService] method instead.
type GetProcedureService struct {
	Options []option.RequestOption
}

// NewGetProcedureService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGetProcedureService(opts ...option.RequestOption) (r GetProcedureService) {
	r = GetProcedureService{}
	r.Options = opts
	return
}

// Retrieves a specific procedure by key from procedural memory. Procedures are
// persistent knowledge artifacts that remain available across all sessions and can
// be shared between different agent instances.
func (r *GetProcedureService) New(ctx context.Context, body GetProcedureNewParams, opts ...option.RequestOption) (res *GetProcedureNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/get_procedure"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type GetProcedureNewResponse struct {
	// Indicates whether the procedure was found
	Found bool `json:"found"`
	// The procedure content, or empty if not found
	Value string `json:"value,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Found       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetProcedureNewResponse) RawJSON() string { return r.JSON.raw }
func (r *GetProcedureNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetProcedureNewParams struct {
	// Unique key of the procedure to retrieve
	Key string `json:"key,required"`
	// Smart memory locator for targeting the correct smart memory instance
	BodySmartMemoryLocation1 GetProcedureNewParamsSmartMemoryLocationUnion `json:"smartMemoryLocation,omitzero,required"`
	// Optional procedural memory ID to use for actor isolation (Alias: accepts both
	// 'proceduralMemoryId' and 'procedural_memory_id')
	BodyProceduralMemoryID1 param.Opt[string] `json:"procedural_memory_id,omitzero"`
	// Optional procedural memory ID to use for actor isolation
	BodyProceduralMemoryID2 param.Opt[string] `json:"proceduralMemoryId,omitzero"`
	// Smart memory locator for targeting the correct smart memory instance (Alias:
	// accepts both 'smartMemoryLocation' and 'smart_memory_location')
	BodySmartMemoryLocation2 GetProcedureNewParamsSmartMemoryLocationUnion `json:"smart_memory_location,omitzero"`
	paramObj
}

func (r GetProcedureNewParams) MarshalJSON() (data []byte, err error) {
	type shadow GetProcedureNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GetProcedureNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type GetProcedureNewParamsSmartMemoryLocationUnion struct {
	OfModuleID    *GetProcedureNewParamsSmartMemoryLocationModuleID    `json:",omitzero,inline"`
	OfSmartMemory *GetProcedureNewParamsSmartMemoryLocationSmartMemory `json:",omitzero,inline"`
	paramUnion
}

func (u GetProcedureNewParamsSmartMemoryLocationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfModuleID, u.OfSmartMemory)
}
func (u *GetProcedureNewParamsSmartMemoryLocationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *GetProcedureNewParamsSmartMemoryLocationUnion) asAny() any {
	if !param.IsOmitted(u.OfModuleID) {
		return u.OfModuleID
	} else if !param.IsOmitted(u.OfSmartMemory) {
		return u.OfSmartMemory
	}
	return nil
}

// The property ModuleID is required.
type GetProcedureNewParamsSmartMemoryLocationModuleID struct {
	// **REQUIRED** FALSE
	ModuleID string `json:"moduleId,required"`
	paramObj
}

func (r GetProcedureNewParamsSmartMemoryLocationModuleID) MarshalJSON() (data []byte, err error) {
	type shadow GetProcedureNewParamsSmartMemoryLocationModuleID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GetProcedureNewParamsSmartMemoryLocationModuleID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SmartMemory is required.
type GetProcedureNewParamsSmartMemoryLocationSmartMemory struct {
	// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
	// **REQUIRED** FALSE
	SmartMemory shared.LiquidmetalV1alpha1SmartMemoryNameParam `json:"smartMemory,omitzero,required"`
	paramObj
}

func (r GetProcedureNewParamsSmartMemoryLocationSmartMemory) MarshalJSON() (data []byte, err error) {
	type shadow GetProcedureNewParamsSmartMemoryLocationSmartMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GetProcedureNewParamsSmartMemoryLocationSmartMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
