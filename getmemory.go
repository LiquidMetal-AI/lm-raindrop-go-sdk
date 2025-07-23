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

// GetMemoryService contains methods and other services that help with interacting
// with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGetMemoryService] method instead.
type GetMemoryService struct {
	Options []option.RequestOption
}

// NewGetMemoryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGetMemoryService(opts ...option.RequestOption) (r GetMemoryService) {
	r = GetMemoryService{}
	r.Options = opts
	return
}

// Retrieves memories based on timeline, key, or temporal criteria. Supports
// filtering by specific timelines, time ranges, and limiting results to the most
// recent entries.
//
// Query capabilities:
//
// - Timeline-specific retrieval
// - Key-based lookup
// - Temporal range queries
// - Most recent N entries
func (r *GetMemoryService) Get(ctx context.Context, body GetMemoryGetParams, opts ...option.RequestOption) (res *GetMemoryGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/get_memory"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type GetMemoryGetResponse struct {
	// List of matching memory entries
	Memories []GetMemoryGetResponseMemory `json:"memories"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Memories    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetMemoryGetResponse) RawJSON() string { return r.JSON.raw }
func (r *GetMemoryGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetMemoryGetResponseMemory struct {
	// Unique identifier for this memory entry
	ID string `json:"id"`
	// Optional agent identifier
	Agent string `json:"agent,nullable"`
	// When this memory was created
	At time.Time `json:"at,nullable" format:"date-time"`
	// Agent that created this memory
	By string `json:"by"`
	// The actual memory content
	Content string `json:"content"`
	// What triggered this memory creation
	DueTo string `json:"dueTo"`
	// Optional key for direct retrieval
	Key string `json:"key,nullable"`
	// Session identifier where this memory was created
	SessionID string `json:"sessionId"`
	// Timeline this memory belongs to
	Timeline string `json:"timeline"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Agent       respjson.Field
		At          respjson.Field
		By          respjson.Field
		Content     respjson.Field
		DueTo       respjson.Field
		Key         respjson.Field
		SessionID   respjson.Field
		Timeline    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetMemoryGetResponseMemory) RawJSON() string { return r.JSON.raw }
func (r *GetMemoryGetResponseMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetMemoryGetParams struct {
	// Unique session identifier for the working memory instance
	SessionID string `json:"sessionId,required"`
	// Smart memory locator for targeting the correct smart memory instance
	SmartMemoryLocation GetMemoryGetParamsSmartMemoryLocation `json:"smartMemoryLocation,omitzero,required"`
	// End time for temporal filtering
	EndTime param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// Specific key to retrieve
	Key param.Opt[string] `json:"key,omitzero"`
	// Maximum number of most recent memories to return
	NMostRecent param.Opt[int64] `json:"nMostRecent,omitzero"`
	// Start time for temporal filtering
	StartTime param.Opt[time.Time] `json:"startTime,omitzero" format:"date-time"`
	// Timeline to filter memories
	Timeline param.Opt[string] `json:"timeline,omitzero"`
	paramObj
}

func (r GetMemoryGetParams) MarshalJSON() (data []byte, err error) {
	type shadow GetMemoryGetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GetMemoryGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SmartMemory is required.
type GetMemoryGetParamsSmartMemoryLocation struct {
	// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
	// **REQUIRED** FALSE
	SmartMemory GetMemoryGetParamsSmartMemoryLocationSmartMemory `json:"smartMemory,omitzero,required"`
	paramObj
}

func (r GetMemoryGetParamsSmartMemoryLocation) MarshalJSON() (data []byte, err error) {
	type shadow GetMemoryGetParamsSmartMemoryLocation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GetMemoryGetParamsSmartMemoryLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
// **REQUIRED** FALSE
//
// The property Name is required.
type GetMemoryGetParamsSmartMemoryLocationSmartMemory struct {
	// The name of the smart memory **EXAMPLE** "my-smartmemory" **REQUIRED** TRUE
	Name string `json:"name,required"`
	// Optional Application **EXAMPLE** "my-app" **REQUIRED** FALSE
	ApplicationName param.Opt[string] `json:"applicationName,omitzero"`
	// Optional version of the smart memory **EXAMPLE** "01jtryx2f2f61ryk06vd8mr91p"
	// **REQUIRED** FALSE
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r GetMemoryGetParamsSmartMemoryLocationSmartMemory) MarshalJSON() (data []byte, err error) {
	type shadow GetMemoryGetParamsSmartMemoryLocationSmartMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GetMemoryGetParamsSmartMemoryLocationSmartMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
