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

// RehydrationStatusService contains methods and other services that help with
// interacting with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRehydrationStatusService] method instead.
type RehydrationStatusService struct {
	Options []option.RequestOption
}

// NewRehydrationStatusService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRehydrationStatusService(opts ...option.RequestOption) (r RehydrationStatusService) {
	r = RehydrationStatusService{}
	r.Options = opts
	return
}

// Gets the current status of an async rehydration operation. Use this to poll for
// completion after calling RehydrateSession. Returns status information including
// progress and any errors.
func (r *RehydrationStatusService) New(ctx context.Context, body RehydrationStatusNewParams, opts ...option.RequestOption) (res *RehydrationStatusNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/rehydration_status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type RehydrationStatusNewResponse struct {
	// When the rehydration operation completed
	CompletedAt string `json:"completedAt,nullable"`
	// Number of entries restored (available when completed)
	EntriesRestored int64 `json:"entriesRestored,nullable"`
	// Error message if the operation failed
	Error string `json:"error,nullable"`
	// When the rehydration operation was initiated
	InitiatedAt string `json:"initiatedAt,nullable"`
	// The session being rehydrated
	SessionID string `json:"sessionId,nullable"`
	// Current status: 'none', 'initiated', 'processing', 'completed', or 'failed'
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt     respjson.Field
		EntriesRestored respjson.Field
		Error           respjson.Field
		InitiatedAt     respjson.Field
		SessionID       respjson.Field
		Status          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RehydrationStatusNewResponse) RawJSON() string { return r.JSON.raw }
func (r *RehydrationStatusNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RehydrationStatusNewParams struct {
	// Session identifier to check rehydration status for
	SessionID string `json:"sessionId,required"`
	// Smart memory locator for targeting the correct smart memory instance
	SmartMemoryLocation RehydrationStatusNewParamsSmartMemoryLocation `json:"smartMemoryLocation,omitzero,required"`
	paramObj
}

func (r RehydrationStatusNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RehydrationStatusNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RehydrationStatusNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SmartMemory is required.
type RehydrationStatusNewParamsSmartMemoryLocation struct {
	// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
	// **REQUIRED** FALSE
	SmartMemory shared.LiquidmetalV1alpha1SmartMemoryNameParam `json:"smartMemory,omitzero,required"`
	paramObj
}

func (r RehydrationStatusNewParamsSmartMemoryLocation) MarshalJSON() (data []byte, err error) {
	type shadow RehydrationStatusNewParamsSmartMemoryLocation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RehydrationStatusNewParamsSmartMemoryLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
