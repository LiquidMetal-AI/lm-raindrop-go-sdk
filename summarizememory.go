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
	opts = slices.Concat(r.Options, opts)
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
	BodyMemoryIDs1 []string `json:"memoryIds,omitzero,required"`
	// Unique session identifier for the working memory instance
	BodySessionID1 string `json:"sessionId,required"`
	// Smart memory locator for targeting the correct smart memory instance
	BodySmartMemoryLocation1 SummarizeMemoryNewParamsSmartMemoryLocation `json:"smartMemoryLocation,omitzero,required"`
	// Optional custom system prompt for summarization (Alias: accepts both
	// 'systemPrompt' and 'system_prompt')
	BodySystemPrompt1 param.Opt[string] `json:"system_prompt,omitzero"`
	// Optional custom system prompt for summarization
	BodySystemPrompt2 param.Opt[string] `json:"systemPrompt,omitzero"`
	// Unique session identifier for the working memory instance (Alias: accepts both
	// 'sessionId' and 'session_id')
	BodySessionID2 param.Opt[string] `json:"session_id,omitzero"`
	// List of memory IDs to summarize (Alias: accepts both 'memoryIds' and
	// 'memory_ids')
	BodyMemoryIDs2 []string `json:"memory_ids,omitzero"`
	// Smart memory locator for targeting the correct smart memory instance (Alias:
	// accepts both 'smartMemoryLocation' and 'smart_memory_location')
	BodySmartMemoryLocation2 SummarizeMemoryNewParamsSmartMemoryLocation `json:"smart_memory_location,omitzero"`
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
	// **REQUIRED** FALSE
	SmartMemory shared.LiquidmetalV1alpha1SmartMemoryNameParam `json:"smartMemory,omitzero,required"`
	paramObj
}

func (r SummarizeMemoryNewParamsSmartMemoryLocation) MarshalJSON() (data []byte, err error) {
	type shadow SummarizeMemoryNewParamsSmartMemoryLocation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SummarizeMemoryNewParamsSmartMemoryLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
