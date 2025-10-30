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

// DeleteProcedureService contains methods and other services that help with
// interacting with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeleteProcedureService] method instead.
type DeleteProcedureService struct {
	Options []option.RequestOption
}

// NewDeleteProcedureService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeleteProcedureService(opts ...option.RequestOption) (r DeleteProcedureService) {
	r = DeleteProcedureService{}
	r.Options = opts
	return
}

// Removes a specific procedure from procedural memory. This operation is permanent
// and affects all future sessions.
func (r *DeleteProcedureService) New(ctx context.Context, body DeleteProcedureNewParams, opts ...option.RequestOption) (res *DeleteProcedureNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/delete_procedure"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type DeleteProcedureNewResponse struct {
	// Indicates whether the procedure was deleted successfully
	Success bool `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeleteProcedureNewResponse) RawJSON() string { return r.JSON.raw }
func (r *DeleteProcedureNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeleteProcedureNewParams struct {
	// Unique key of the procedure to delete
	Key string `json:"key,required"`
	// Smart memory locator for targeting the correct smart memory instance
	BodySmartMemoryLocation1 DeleteProcedureNewParamsSmartMemoryLocation `json:"smartMemoryLocation,omitzero,required"`
	// Optional procedural memory ID to use for actor isolation (Alias: accepts both
	// 'proceduralMemoryId' and 'procedural_memory_id')
	BodyProceduralMemoryID1 param.Opt[string] `json:"procedural_memory_id,omitzero"`
	// Optional procedural memory ID to use for actor isolation
	BodyProceduralMemoryID2 param.Opt[string] `json:"proceduralMemoryId,omitzero"`
	// Smart memory locator for targeting the correct smart memory instance (Alias:
	// accepts both 'smartMemoryLocation' and 'smart_memory_location')
	BodySmartMemoryLocation2 DeleteProcedureNewParamsSmartMemoryLocation `json:"smart_memory_location,omitzero"`
	paramObj
}

func (r DeleteProcedureNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DeleteProcedureNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeleteProcedureNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SmartMemory is required.
type DeleteProcedureNewParamsSmartMemoryLocation struct {
	// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
	// **REQUIRED** FALSE
	SmartMemory shared.LiquidmetalV1alpha1SmartMemoryNameParam `json:"smartMemory,omitzero,required"`
	paramObj
}

func (r DeleteProcedureNewParamsSmartMemoryLocation) MarshalJSON() (data []byte, err error) {
	type shadow DeleteProcedureNewParamsSmartMemoryLocation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeleteProcedureNewParamsSmartMemoryLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
