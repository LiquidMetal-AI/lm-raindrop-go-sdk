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

// DeleteSemanticMemoryService contains methods and other services that help with
// interacting with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeleteSemanticMemoryService] method instead.
type DeleteSemanticMemoryService struct {
	Options []option.RequestOption
}

// NewDeleteSemanticMemoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDeleteSemanticMemoryService(opts ...option.RequestOption) (r DeleteSemanticMemoryService) {
	r = DeleteSemanticMemoryService{}
	r.Options = opts
	return
}

// Removes a specific semantic memory document by its object ID. This operation
// permanently deletes the document and is irreversible.
func (r *DeleteSemanticMemoryService) Delete(ctx context.Context, body DeleteSemanticMemoryDeleteParams, opts ...option.RequestOption) (res *DeleteSemanticMemoryDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/delete_semantic_memory"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type DeleteSemanticMemoryDeleteResponse struct {
	// Error message if the operation failed
	Error string `json:"error,nullable"`
	// Indicates whether the document was deleted successfully
	Success bool `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeleteSemanticMemoryDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *DeleteSemanticMemoryDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeleteSemanticMemoryDeleteParams struct {
	// Unique object identifier of the document to delete
	ObjectID string `json:"objectId,required"`
	// Smart memory locator for targeting the correct smart memory instance
	SmartMemoryLocation DeleteSemanticMemoryDeleteParamsSmartMemoryLocation `json:"smartMemoryLocation,omitzero,required"`
	paramObj
}

func (r DeleteSemanticMemoryDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow DeleteSemanticMemoryDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeleteSemanticMemoryDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SmartMemory is required.
type DeleteSemanticMemoryDeleteParamsSmartMemoryLocation struct {
	// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
	// **REQUIRED** TRUE
	SmartMemory DeleteSemanticMemoryDeleteParamsSmartMemoryLocationSmartMemory `json:"smartMemory,omitzero,required"`
	paramObj
}

func (r DeleteSemanticMemoryDeleteParamsSmartMemoryLocation) MarshalJSON() (data []byte, err error) {
	type shadow DeleteSemanticMemoryDeleteParamsSmartMemoryLocation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeleteSemanticMemoryDeleteParamsSmartMemoryLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
// **REQUIRED** TRUE
//
// The property Name is required.
type DeleteSemanticMemoryDeleteParamsSmartMemoryLocationSmartMemory struct {
	// The name of the smart memory **EXAMPLE** "my-smartmemory" **REQUIRED** TRUE
	Name string `json:"name,required"`
	// Optional Application **EXAMPLE** "my-app" **REQUIRED** FALSE
	ApplicationName param.Opt[string] `json:"applicationName,omitzero"`
	// Optional version of the smart memory **EXAMPLE** "01jtryx2f2f61ryk06vd8mr91p"
	// **REQUIRED** FALSE
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r DeleteSemanticMemoryDeleteParamsSmartMemoryLocationSmartMemory) MarshalJSON() (data []byte, err error) {
	type shadow DeleteSemanticMemoryDeleteParamsSmartMemoryLocationSmartMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeleteSemanticMemoryDeleteParamsSmartMemoryLocationSmartMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
