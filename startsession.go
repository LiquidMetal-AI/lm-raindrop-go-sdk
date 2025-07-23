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

// StartSessionService contains methods and other services that help with
// interacting with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStartSessionService] method instead.
type StartSessionService struct {
	Options []option.RequestOption
}

// NewStartSessionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStartSessionService(opts ...option.RequestOption) (r StartSessionService) {
	r = StartSessionService{}
	r.Options = opts
	return
}

// Starts a new working memory session for an agent. Each session provides isolated
// memory operations and automatic cleanup capabilities.
func (r *StartSessionService) New(ctx context.Context, body StartSessionNewParams, opts ...option.RequestOption) (res *StartSessionNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/start_session"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type StartSessionNewResponse struct {
	// Unique identifier for the new session
	SessionID string `json:"sessionId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SessionID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StartSessionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *StartSessionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StartSessionNewParams struct {
	// Smart memory locator for targeting the correct smart memory instance
	SmartMemoryLocation StartSessionNewParamsSmartMemoryLocationUnion `json:"smartMemoryLocation,omitzero,required"`
	paramObj
}

func (r StartSessionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow StartSessionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StartSessionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type StartSessionNewParamsSmartMemoryLocationUnion struct {
	OfModuleID    *StartSessionNewParamsSmartMemoryLocationModuleID    `json:",omitzero,inline"`
	OfSmartMemory *StartSessionNewParamsSmartMemoryLocationSmartMemory `json:",omitzero,inline"`
	paramUnion
}

func (u StartSessionNewParamsSmartMemoryLocationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfModuleID, u.OfSmartMemory)
}
func (u *StartSessionNewParamsSmartMemoryLocationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *StartSessionNewParamsSmartMemoryLocationUnion) asAny() any {
	if !param.IsOmitted(u.OfModuleID) {
		return u.OfModuleID
	} else if !param.IsOmitted(u.OfSmartMemory) {
		return u.OfSmartMemory
	}
	return nil
}

// The property ModuleID is required.
type StartSessionNewParamsSmartMemoryLocationModuleID struct {
	ModuleID string `json:"moduleId,required"`
	paramObj
}

func (r StartSessionNewParamsSmartMemoryLocationModuleID) MarshalJSON() (data []byte, err error) {
	type shadow StartSessionNewParamsSmartMemoryLocationModuleID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StartSessionNewParamsSmartMemoryLocationModuleID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SmartMemory is required.
type StartSessionNewParamsSmartMemoryLocationSmartMemory struct {
	// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
	// **REQUIRED** FALSE
	SmartMemory StartSessionNewParamsSmartMemoryLocationSmartMemorySmartMemory `json:"smartMemory,omitzero,required"`
	paramObj
}

func (r StartSessionNewParamsSmartMemoryLocationSmartMemory) MarshalJSON() (data []byte, err error) {
	type shadow StartSessionNewParamsSmartMemoryLocationSmartMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StartSessionNewParamsSmartMemoryLocationSmartMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
// **REQUIRED** FALSE
//
// The property Name is required.
type StartSessionNewParamsSmartMemoryLocationSmartMemorySmartMemory struct {
	// The name of the smart memory **EXAMPLE** "my-smartmemory" **REQUIRED** TRUE
	Name string `json:"name,required"`
	// Optional Application **EXAMPLE** "my-app" **REQUIRED** FALSE
	ApplicationName param.Opt[string] `json:"applicationName,omitzero"`
	// Optional version of the smart memory **EXAMPLE** "01jtryx2f2f61ryk06vd8mr91p"
	// **REQUIRED** FALSE
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r StartSessionNewParamsSmartMemoryLocationSmartMemorySmartMemory) MarshalJSON() (data []byte, err error) {
	type shadow StartSessionNewParamsSmartMemoryLocationSmartMemorySmartMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StartSessionNewParamsSmartMemoryLocationSmartMemorySmartMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
