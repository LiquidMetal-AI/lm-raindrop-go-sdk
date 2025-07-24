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

// PutProcedureService contains methods and other services that help with
// interacting with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPutProcedureService] method instead.
type PutProcedureService struct {
	Options []option.RequestOption
}

// NewPutProcedureService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPutProcedureService(opts ...option.RequestOption) (r PutProcedureService) {
	r = PutProcedureService{}
	r.Options = opts
	return
}

// Stores a new procedure in the agent's procedural memory. Procedures are reusable
// knowledge artifacts like system prompts, templates, workflows, or instructions
// that can be retrieved and applied across different sessions and contexts.
func (r *PutProcedureService) New(ctx context.Context, body PutProcedureNewParams, opts ...option.RequestOption) (res *PutProcedureNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/put_procedure"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type PutProcedureNewResponse struct {
	// Indicates whether the procedure was stored successfully
	Success bool `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PutProcedureNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PutProcedureNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PutProcedureNewParams struct {
	// Unique key to identify this procedure
	Key string `json:"key,required"`
	// Smart memory locator for targeting the correct smart memory instance
	SmartMemoryLocation PutProcedureNewParamsSmartMemoryLocation `json:"smartMemoryLocation,omitzero,required"`
	// The procedure content (prompt, template, instructions, etc.)
	Value string `json:"value,required"`
	// Optional procedural memory ID to use for actor isolation
	ProceduralMemoryID param.Opt[string] `json:"proceduralMemoryId,omitzero"`
	paramObj
}

func (r PutProcedureNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PutProcedureNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PutProcedureNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SmartMemory is required.
type PutProcedureNewParamsSmartMemoryLocation struct {
	// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
	// **REQUIRED** TRUE
	SmartMemory PutProcedureNewParamsSmartMemoryLocationSmartMemory `json:"smartMemory,omitzero,required"`
	paramObj
}

func (r PutProcedureNewParamsSmartMemoryLocation) MarshalJSON() (data []byte, err error) {
	type shadow PutProcedureNewParamsSmartMemoryLocation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PutProcedureNewParamsSmartMemoryLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
// **REQUIRED** TRUE
//
// The property Name is required.
type PutProcedureNewParamsSmartMemoryLocationSmartMemory struct {
	// The name of the smart memory **EXAMPLE** "my-smartmemory" **REQUIRED** TRUE
	Name string `json:"name,required"`
	// Optional Application **EXAMPLE** "my-app" **REQUIRED** FALSE
	ApplicationName param.Opt[string] `json:"applicationName,omitzero"`
	// Optional version of the smart memory **EXAMPLE** "01jtryx2f2f61ryk06vd8mr91p"
	// **REQUIRED** FALSE
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r PutProcedureNewParamsSmartMemoryLocationSmartMemory) MarshalJSON() (data []byte, err error) {
	type shadow PutProcedureNewParamsSmartMemoryLocationSmartMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PutProcedureNewParamsSmartMemoryLocationSmartMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
