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

// DeleteMemoryService contains methods and other services that help with
// interacting with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeleteMemoryService] method instead.
type DeleteMemoryService struct {
	Options []option.RequestOption
}

// NewDeleteMemoryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeleteMemoryService(opts ...option.RequestOption) (r DeleteMemoryService) {
	r = DeleteMemoryService{}
	r.Options = opts
	return
}

// Removes a specific memory entry from storage. This operation is permanent and
// cannot be undone.
func (r *DeleteMemoryService) New(ctx context.Context, body DeleteMemoryNewParams, opts ...option.RequestOption) (res *DeleteMemoryNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/delete_memory"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type DeleteMemoryNewResponse struct {
	// Indicates whether the deletion was successful
	Success bool `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeleteMemoryNewResponse) RawJSON() string { return r.JSON.raw }
func (r *DeleteMemoryNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeleteMemoryNewParams struct {
	// Unique identifier of the memory entry to delete
	MemoryID string `json:"memoryId,required"`
	// Unique session identifier for the working memory instance
	SessionID string `json:"sessionId,required"`
	// Smart memory locator for targeting the correct smart memory instance
	SmartMemoryLocation DeleteMemoryNewParamsSmartMemoryLocation `json:"smartMemoryLocation,omitzero,required"`
	paramObj
}

func (r DeleteMemoryNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DeleteMemoryNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeleteMemoryNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SmartMemory is required.
type DeleteMemoryNewParamsSmartMemoryLocation struct {
	// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
	// **REQUIRED** TRUE
	SmartMemory DeleteMemoryNewParamsSmartMemoryLocationSmartMemory `json:"smartMemory,omitzero,required"`
	paramObj
}

func (r DeleteMemoryNewParamsSmartMemoryLocation) MarshalJSON() (data []byte, err error) {
	type shadow DeleteMemoryNewParamsSmartMemoryLocation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeleteMemoryNewParamsSmartMemoryLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
// **REQUIRED** TRUE
//
// The property Name is required.
type DeleteMemoryNewParamsSmartMemoryLocationSmartMemory struct {
	// The name of the smart memory **EXAMPLE** "my-smartmemory" **REQUIRED** TRUE
	Name string `json:"name,required"`
	// Optional Application **EXAMPLE** "my-app" **REQUIRED** FALSE
	ApplicationName param.Opt[string] `json:"applicationName,omitzero"`
	// Optional version of the smart memory **EXAMPLE** "01jtryx2f2f61ryk06vd8mr91p"
	// **REQUIRED** FALSE
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r DeleteMemoryNewParamsSmartMemoryLocationSmartMemory) MarshalJSON() (data []byte, err error) {
	type shadow DeleteMemoryNewParamsSmartMemoryLocationSmartMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeleteMemoryNewParamsSmartMemoryLocationSmartMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
