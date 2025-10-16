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
	ObjectID string `json:"objectId,required"`
	// Smart memory locator for targeting the correct smart memory instance
	SmartMemoryLocation GetSemanticMemoryNewParamsSmartMemoryLocation `json:"smartMemoryLocation,omitzero,required"`
	paramObj
}

func (r GetSemanticMemoryNewParams) MarshalJSON() (data []byte, err error) {
	type shadow GetSemanticMemoryNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GetSemanticMemoryNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SmartMemory is required.
type GetSemanticMemoryNewParamsSmartMemoryLocation struct {
	// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
	// **REQUIRED** FALSE
	SmartMemory shared.LiquidmetalV1alpha1SmartMemoryNameParam `json:"smartMemory,omitzero,required"`
	paramObj
}

func (r GetSemanticMemoryNewParamsSmartMemoryLocation) MarshalJSON() (data []byte, err error) {
	type shadow GetSemanticMemoryNewParamsSmartMemoryLocation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GetSemanticMemoryNewParamsSmartMemoryLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
