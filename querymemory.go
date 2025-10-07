// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package raindrop

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/apijson"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/requestconfig"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/option"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/param"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/respjson"
)

// QueryMemoryService contains methods and other services that help with
// interacting with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQueryMemoryService] method instead.
type QueryMemoryService struct {
	Options []option.RequestOption
}

// NewQueryMemoryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewQueryMemoryService(opts ...option.RequestOption) (r QueryMemoryService) {
	r = QueryMemoryService{}
	r.Options = opts
	return
}

// Performs semantic search across stored memories using natural language queries.
// The system uses vector embeddings to find semantically similar content
// regardless of exact keyword matches.
//
// Search features:
//
// - Semantic similarity matching
// - Timeline-specific search
// - Temporal filtering
// - Relevance-based ranking
func (r *QueryMemoryService) Search(ctx context.Context, body QueryMemorySearchParams, opts ...option.RequestOption) (res *QueryMemorySearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/search_memory"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type QueryMemorySearchResponse struct {
	// List of matching memory entries ordered by relevance
	Memories []QueryMemorySearchResponseMemory `json:"memories"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Memories    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryMemorySearchResponse) RawJSON() string { return r.JSON.raw }
func (r *QueryMemorySearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryMemorySearchResponseMemory struct {
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
func (r QueryMemorySearchResponseMemory) RawJSON() string { return r.JSON.raw }
func (r *QueryMemorySearchResponseMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryMemorySearchParams struct {
	// Unique session identifier for the working memory instance
	SessionID string `json:"sessionId,required"`
	// Smart memory locator for targeting the correct smart memory instance
	SmartMemoryLocation QueryMemorySearchParamsSmartMemoryLocation `json:"smartMemoryLocation,omitzero,required"`
	// Natural language search query
	Terms string `json:"terms,required"`
	// End time for temporal filtering
	EndTime param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// Maximum number of most recent results to return
	NMostRecent param.Opt[int64] `json:"nMostRecent,omitzero"`
	// Start time for temporal filtering
	StartTime param.Opt[time.Time] `json:"startTime,omitzero" format:"date-time"`
	// Timeline to filter search results
	Timeline param.Opt[string] `json:"timeline,omitzero"`
	paramObj
}

func (r QueryMemorySearchParams) MarshalJSON() (data []byte, err error) {
	type shadow QueryMemorySearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QueryMemorySearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SmartMemory is required.
type QueryMemorySearchParamsSmartMemoryLocation struct {
	// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
	// **REQUIRED** FALSE
	SmartMemory QueryMemorySearchParamsSmartMemoryLocationSmartMemory `json:"smartMemory,omitzero,required"`
	paramObj
}

func (r QueryMemorySearchParamsSmartMemoryLocation) MarshalJSON() (data []byte, err error) {
	type shadow QueryMemorySearchParamsSmartMemoryLocation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QueryMemorySearchParamsSmartMemoryLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
// **REQUIRED** FALSE
//
// The property Name is required.
type QueryMemorySearchParamsSmartMemoryLocationSmartMemory struct {
	// The name of the smart memory **EXAMPLE** "my-smartmemory" **REQUIRED** TRUE
	Name string `json:"name,required"`
	// Optional Application **EXAMPLE** "my-app" **REQUIRED** FALSE
	ApplicationName param.Opt[string] `json:"applicationName,omitzero"`
	// Optional version of the smart memory **EXAMPLE** "01jtryx2f2f61ryk06vd8mr91p"
	// **REQUIRED** FALSE
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r QueryMemorySearchParamsSmartMemoryLocationSmartMemory) MarshalJSON() (data []byte, err error) {
	type shadow QueryMemorySearchParamsSmartMemoryLocationSmartMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QueryMemorySearchParamsSmartMemoryLocationSmartMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
