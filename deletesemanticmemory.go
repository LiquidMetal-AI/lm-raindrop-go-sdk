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
	ObjectID string `json:"object_id,required"`
	// Smart memory locator for targeting the correct smart memory instance
	SmartMemoryLocation DeleteSemanticMemoryDeleteParamsSmartMemoryLocationUnion `json:"smart_memory_location,omitzero,required"`
	OrganizationID      param.Opt[string]                                        `json:"organization_id,omitzero"`
	UserID              param.Opt[string]                                        `json:"user_id,omitzero"`
	paramObj
}

func (r DeleteSemanticMemoryDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow DeleteSemanticMemoryDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeleteSemanticMemoryDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DeleteSemanticMemoryDeleteParamsSmartMemoryLocationUnion struct {
	OfModuleID    *DeleteSemanticMemoryDeleteParamsSmartMemoryLocationModuleID    `json:",omitzero,inline"`
	OfSmartMemory *DeleteSemanticMemoryDeleteParamsSmartMemoryLocationSmartMemory `json:",omitzero,inline"`
	paramUnion
}

func (u DeleteSemanticMemoryDeleteParamsSmartMemoryLocationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfModuleID, u.OfSmartMemory)
}
func (u *DeleteSemanticMemoryDeleteParamsSmartMemoryLocationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *DeleteSemanticMemoryDeleteParamsSmartMemoryLocationUnion) asAny() any {
	if !param.IsOmitted(u.OfModuleID) {
		return u.OfModuleID
	} else if !param.IsOmitted(u.OfSmartMemory) {
		return u.OfSmartMemory
	}
	return nil
}

// The property ModuleID is required.
type DeleteSemanticMemoryDeleteParamsSmartMemoryLocationModuleID struct {
	// **REQUIRED** FALSE
	ModuleID string `json:"module_id,required"`
	paramObj
}

func (r DeleteSemanticMemoryDeleteParamsSmartMemoryLocationModuleID) MarshalJSON() (data []byte, err error) {
	type shadow DeleteSemanticMemoryDeleteParamsSmartMemoryLocationModuleID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeleteSemanticMemoryDeleteParamsSmartMemoryLocationModuleID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SmartMemory is required.
type DeleteSemanticMemoryDeleteParamsSmartMemoryLocationSmartMemory struct {
	// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
	// **REQUIRED** FALSE
	SmartMemory shared.LiquidmetalV1alpha1SmartMemoryNameParam `json:"smart_memory,omitzero,required"`
	paramObj
}

func (r DeleteSemanticMemoryDeleteParamsSmartMemoryLocationSmartMemory) MarshalJSON() (data []byte, err error) {
	type shadow DeleteSemanticMemoryDeleteParamsSmartMemoryLocationSmartMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeleteSemanticMemoryDeleteParamsSmartMemoryLocationSmartMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
