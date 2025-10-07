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

// RehydrateSessionService contains methods and other services that help with
// interacting with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRehydrateSessionService] method instead.
type RehydrateSessionService struct {
	Options []option.RequestOption
}

// NewRehydrateSessionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRehydrateSessionService(opts ...option.RequestOption) (r RehydrateSessionService) {
	r = RehydrateSessionService{}
	r.Options = opts
	return
}

// Rehydrates a previous session from episodic memory storage. Allows resuming work
// from where a previous session left off by restoring either all memories or just
// a summary of the previous session.
func (r *RehydrateSessionService) Rehydrate(ctx context.Context, body RehydrateSessionRehydrateParams, opts ...option.RequestOption) (res *RehydrateSessionRehydrateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/rehydrate_session"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type RehydrateSessionRehydrateResponse struct {
	// Operation status: 'initiated' for async processing, 'failed' for immediate
	// failure
	Operation string `json:"operation"`
	// Storage key for checking async operation status (optional)
	StatusKey string `json:"statusKey,nullable"`
	// Indicates whether the rehydration was successful
	Success bool `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operation   respjson.Field
		StatusKey   respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RehydrateSessionRehydrateResponse) RawJSON() string { return r.JSON.raw }
func (r *RehydrateSessionRehydrateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RehydrateSessionRehydrateParams struct {
	// Session identifier to restore from episodic memory
	SessionID string `json:"sessionId,required"`
	// Smart memory locator for targeting the correct smart memory instance
	SmartMemoryLocation RehydrateSessionRehydrateParamsSmartMemoryLocation `json:"smartMemoryLocation,omitzero,required"`
	// If true, only restore a summary. If false, restore all memories
	SummaryOnly param.Opt[bool] `json:"summaryOnly,omitzero"`
	paramObj
}

func (r RehydrateSessionRehydrateParams) MarshalJSON() (data []byte, err error) {
	type shadow RehydrateSessionRehydrateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RehydrateSessionRehydrateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SmartMemory is required.
type RehydrateSessionRehydrateParamsSmartMemoryLocation struct {
	// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
	// **REQUIRED** FALSE
	SmartMemory RehydrateSessionRehydrateParamsSmartMemoryLocationSmartMemory `json:"smartMemory,omitzero,required"`
	paramObj
}

func (r RehydrateSessionRehydrateParamsSmartMemoryLocation) MarshalJSON() (data []byte, err error) {
	type shadow RehydrateSessionRehydrateParamsSmartMemoryLocation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RehydrateSessionRehydrateParamsSmartMemoryLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
// **REQUIRED** FALSE
//
// The property Name is required.
type RehydrateSessionRehydrateParamsSmartMemoryLocationSmartMemory struct {
	// The name of the smart memory **EXAMPLE** "my-smartmemory" **REQUIRED** TRUE
	Name string `json:"name,required"`
	// Optional Application **EXAMPLE** "my-app" **REQUIRED** FALSE
	ApplicationName param.Opt[string] `json:"applicationName,omitzero"`
	// Optional version of the smart memory **EXAMPLE** "01jtryx2f2f61ryk06vd8mr91p"
	// **REQUIRED** FALSE
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r RehydrateSessionRehydrateParamsSmartMemoryLocationSmartMemory) MarshalJSON() (data []byte, err error) {
	type shadow RehydrateSessionRehydrateParamsSmartMemoryLocationSmartMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RehydrateSessionRehydrateParamsSmartMemoryLocationSmartMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
