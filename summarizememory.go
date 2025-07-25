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

// SummarizeMemoryService contains methods and other services that help with
// interacting with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSummarizeMemoryService] method instead.
type SummarizeMemoryService struct {
	Options []option.RequestOption
}

// NewSummarizeMemoryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSummarizeMemoryService(opts ...option.RequestOption) (r SummarizeMemoryService) {
	r = SummarizeMemoryService{}
	r.Options = opts
	return
}

// Generates intelligent summaries of a collection of memories using AI. Can
// optionally accept custom system prompts to guide the summarization style.
//
// The summarization system:
//
// - Identifies key themes and patterns
// - Extracts important events and decisions
// - Maintains temporal context
// - Supports custom summarization instructions
func (r *SummarizeMemoryService) New(ctx context.Context, body SummarizeMemoryNewParams, opts ...option.RequestOption) (res *SummarizeMemoryNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/summarize_memory"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SummarizeMemoryNewResponse struct {
	// List of memory IDs that were summarized
	SummarizedMemoryIDs []string `json:"summarizedMemoryIds"`
	// AI-generated summary of the memories
	Summary string `json:"summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SummarizedMemoryIDs respjson.Field
		Summary             respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SummarizeMemoryNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SummarizeMemoryNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SummarizeMemoryNewParams struct {
	// List of memory IDs to summarize
	MemoryIDs []string `json:"memoryIds,omitzero,required"`
	// Unique session identifier for the working memory instance
	SessionID string `json:"sessionId,required"`
	// Smart memory locator for targeting the correct smart memory instance
	SmartMemoryLocation SummarizeMemoryNewParamsSmartMemoryLocation `json:"smartMemoryLocation,omitzero,required"`
	// Optional custom system prompt for summarization
	SystemPrompt param.Opt[string] `json:"systemPrompt,omitzero"`
	paramObj
}

func (r SummarizeMemoryNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SummarizeMemoryNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SummarizeMemoryNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SmartMemory is required.
type SummarizeMemoryNewParamsSmartMemoryLocation struct {
	// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
	// **REQUIRED** TRUE
	SmartMemory SummarizeMemoryNewParamsSmartMemoryLocationSmartMemory `json:"smartMemory,omitzero,required"`
	paramObj
}

func (r SummarizeMemoryNewParamsSmartMemoryLocation) MarshalJSON() (data []byte, err error) {
	type shadow SummarizeMemoryNewParamsSmartMemoryLocation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SummarizeMemoryNewParamsSmartMemoryLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
// **REQUIRED** TRUE
//
// The property Name is required.
type SummarizeMemoryNewParamsSmartMemoryLocationSmartMemory struct {
	// The name of the smart memory **EXAMPLE** "my-smartmemory" **REQUIRED** TRUE
	Name string `json:"name,required"`
	// Optional Application **EXAMPLE** "my-app" **REQUIRED** FALSE
	ApplicationName param.Opt[string] `json:"applicationName,omitzero"`
	// Optional version of the smart memory **EXAMPLE** "01jtryx2f2f61ryk06vd8mr91p"
	// **REQUIRED** FALSE
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r SummarizeMemoryNewParamsSmartMemoryLocationSmartMemory) MarshalJSON() (data []byte, err error) {
	type shadow SummarizeMemoryNewParamsSmartMemoryLocationSmartMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SummarizeMemoryNewParamsSmartMemoryLocationSmartMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
