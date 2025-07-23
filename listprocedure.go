// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package raindrop

import (
	"context"
	"net/http"
	"time"

	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/apijson"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/requestconfig"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/option"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/param"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/respjson"
)

// ListProcedureService contains methods and other services that help with
// interacting with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewListProcedureService] method instead.
type ListProcedureService struct {
	Options []option.RequestOption
}

// NewListProcedureService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewListProcedureService(opts ...option.RequestOption) (r ListProcedureService) {
	r = ListProcedureService{}
	r.Options = opts
	return
}

// Lists all procedures stored in procedural memory. Returns metadata about each
// procedure including creation and modification times.
func (r *ListProcedureService) New(ctx context.Context, body ListProcedureNewParams, opts ...option.RequestOption) (res *ListProcedureNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/list_procedures"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ListProcedureNewResponse struct {
	// List of all stored procedures
	Procedures []ListProcedureNewResponseProcedure `json:"procedures"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Procedures  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListProcedureNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ListProcedureNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListProcedureNewResponseProcedure struct {
	// When this procedure was first created
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Unique key for this procedure
	Key string `json:"key"`
	// When this procedure was last updated
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
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
func (r ListProcedureNewResponseProcedure) RawJSON() string { return r.JSON.raw }
func (r *ListProcedureNewResponseProcedure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListProcedureNewParams struct {
	// Smart memory locator for targeting the correct smart memory instance
	SmartMemoryLocation ListProcedureNewParamsSmartMemoryLocation `json:"smartMemoryLocation,omitzero,required"`
	// Optional procedural memory ID to use for actor isolation
	ProceduralMemoryID param.Opt[string] `json:"proceduralMemoryId,omitzero"`
	paramObj
}

func (r ListProcedureNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ListProcedureNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListProcedureNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SmartMemory is required.
type ListProcedureNewParamsSmartMemoryLocation struct {
	// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
	// **REQUIRED** FALSE
	SmartMemory ListProcedureNewParamsSmartMemoryLocationSmartMemory `json:"smartMemory,omitzero,required"`
	paramObj
}

func (r ListProcedureNewParamsSmartMemoryLocation) MarshalJSON() (data []byte, err error) {
	type shadow ListProcedureNewParamsSmartMemoryLocation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListProcedureNewParamsSmartMemoryLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
// **REQUIRED** FALSE
//
// The property Name is required.
type ListProcedureNewParamsSmartMemoryLocationSmartMemory struct {
	// The name of the smart memory **EXAMPLE** "my-smartmemory" **REQUIRED** TRUE
	Name string `json:"name,required"`
	// Optional Application **EXAMPLE** "my-app" **REQUIRED** FALSE
	ApplicationName param.Opt[string] `json:"applicationName,omitzero"`
	// Optional version of the smart memory **EXAMPLE** "01jtryx2f2f61ryk06vd8mr91p"
	// **REQUIRED** FALSE
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r ListProcedureNewParamsSmartMemoryLocationSmartMemory) MarshalJSON() (data []byte, err error) {
	type shadow ListProcedureNewParamsSmartMemoryLocationSmartMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ListProcedureNewParamsSmartMemoryLocationSmartMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
