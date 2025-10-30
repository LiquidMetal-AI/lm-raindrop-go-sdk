// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package raindrop

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"
	"time"

	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/apijson"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/requestconfig"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/option"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/param"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/respjson"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/shared"
)

// QueryEpisodicMemoryService contains methods and other services that help with
// interacting with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQueryEpisodicMemoryService] method instead.
type QueryEpisodicMemoryService struct {
	Options []option.RequestOption
}

// NewQueryEpisodicMemoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewQueryEpisodicMemoryService(opts ...option.RequestOption) (r QueryEpisodicMemoryService) {
	r = QueryEpisodicMemoryService{}
	r.Options = opts
	return
}

// Searches across episodic memory documents stored in the SmartBucket. Allows
// finding relevant past sessions based on natural language queries. Returns
// summaries and metadata from stored episodic memory sessions.
func (r *QueryEpisodicMemoryService) Search(ctx context.Context, body QueryEpisodicMemorySearchParams, opts ...option.RequestOption) (res *QueryEpisodicMemorySearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/search_episodic_memory"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type QueryEpisodicMemorySearchResponse struct {
	// List of matching episodic memory entries ordered by relevance
	Entries []QueryEpisodicMemorySearchResponseEntry `json:"entries"`
	// Pagination information for the search results
	Pagination QueryEpisodicMemorySearchResponsePagination `json:"pagination,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entries     respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryEpisodicMemorySearchResponse) RawJSON() string { return r.JSON.raw }
func (r *QueryEpisodicMemorySearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryEpisodicMemorySearchResponseEntry struct {
	// Agent that created this episodic memory
	Agent string `json:"agent"`
	// When this episodic memory was created
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Duration of the session in milliseconds
	Duration QueryEpisodicMemorySearchResponseEntryDurationUnion `json:"duration" format:"int64"`
	// Number of individual memory entries in this session
	EntryCount int64 `json:"entryCount"`
	// Relevance score for this search result
	Score float64 `json:"score,nullable"`
	// Session identifier for this episodic memory
	SessionID string `json:"sessionId"`
	// AI-generated summary of the session
	Summary string `json:"summary"`
	// Number of different timelines in this session
	TimelineCount int64 `json:"timelineCount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Agent         respjson.Field
		CreatedAt     respjson.Field
		Duration      respjson.Field
		EntryCount    respjson.Field
		Score         respjson.Field
		SessionID     respjson.Field
		Summary       respjson.Field
		TimelineCount respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryEpisodicMemorySearchResponseEntry) RawJSON() string { return r.JSON.raw }
func (r *QueryEpisodicMemorySearchResponseEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// QueryEpisodicMemorySearchResponseEntryDurationUnion contains all possible
// properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type QueryEpisodicMemorySearchResponseEntryDurationUnion struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u QueryEpisodicMemorySearchResponseEntryDurationUnion) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QueryEpisodicMemorySearchResponseEntryDurationUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u QueryEpisodicMemorySearchResponseEntryDurationUnion) RawJSON() string { return u.JSON.raw }

func (r *QueryEpisodicMemorySearchResponseEntryDurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination information for the search results
type QueryEpisodicMemorySearchResponsePagination struct {
	// Whether there are more results available
	HasMore bool `json:"hasMore"`
	// Current page number
	Page int64 `json:"page"`
	// Number of results per page
	PageSize int64 `json:"pageSize"`
	// Total number of results available
	Total int64 `json:"total"`
	// Total number of pages available
	TotalPages int64 `json:"totalPages"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		Page        respjson.Field
		PageSize    respjson.Field
		Total       respjson.Field
		TotalPages  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryEpisodicMemorySearchResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *QueryEpisodicMemorySearchResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryEpisodicMemorySearchParams struct {
	// Smart memory locator for targeting the correct smart memory instance
	BodySmartMemoryLocation1 QueryEpisodicMemorySearchParamsSmartMemoryLocationUnion `json:"smartMemoryLocation,omitzero,required"`
	// Natural language search query to find relevant episodic memory sessions
	Terms string `json:"terms,required"`
	// End time for temporal filtering (Alias: accepts both 'endTime' and 'end_time')
	BodyEndTime1 param.Opt[time.Time] `json:"end_time,omitzero" format:"date-time"`
	// End time for temporal filtering
	BodyEndTime2 param.Opt[time.Time] `json:"endTime,omitzero" format:"date-time"`
	// Maximum number of most recent results to return (Alias: accepts both
	// 'nMostRecent' and 'n_most_recent')
	BodyNMostRecent1 param.Opt[int64] `json:"n_most_recent,omitzero"`
	// Maximum number of most recent results to return
	BodyNMostRecent2 param.Opt[int64] `json:"nMostRecent,omitzero"`
	// Start time for temporal filtering (Alias: accepts both 'startTime' and
	// 'start_time')
	BodyStartTime1 param.Opt[time.Time] `json:"start_time,omitzero" format:"date-time"`
	// Start time for temporal filtering
	BodyStartTime2 param.Opt[time.Time] `json:"startTime,omitzero" format:"date-time"`
	// Smart memory locator for targeting the correct smart memory instance (Alias:
	// accepts both 'smartMemoryLocation' and 'smart_memory_location')
	BodySmartMemoryLocation2 QueryEpisodicMemorySearchParamsSmartMemoryLocationUnion `json:"smart_memory_location,omitzero"`
	paramObj
}

func (r QueryEpisodicMemorySearchParams) MarshalJSON() (data []byte, err error) {
	type shadow QueryEpisodicMemorySearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QueryEpisodicMemorySearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type QueryEpisodicMemorySearchParamsSmartMemoryLocationUnion struct {
	OfModuleID    *QueryEpisodicMemorySearchParamsSmartMemoryLocationModuleID    `json:",omitzero,inline"`
	OfSmartMemory *QueryEpisodicMemorySearchParamsSmartMemoryLocationSmartMemory `json:",omitzero,inline"`
	paramUnion
}

func (u QueryEpisodicMemorySearchParamsSmartMemoryLocationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfModuleID, u.OfSmartMemory)
}
func (u *QueryEpisodicMemorySearchParamsSmartMemoryLocationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *QueryEpisodicMemorySearchParamsSmartMemoryLocationUnion) asAny() any {
	if !param.IsOmitted(u.OfModuleID) {
		return u.OfModuleID
	} else if !param.IsOmitted(u.OfSmartMemory) {
		return u.OfSmartMemory
	}
	return nil
}

// The property ModuleID is required.
type QueryEpisodicMemorySearchParamsSmartMemoryLocationModuleID struct {
	// **REQUIRED** FALSE
	ModuleID string `json:"moduleId,required"`
	paramObj
}

func (r QueryEpisodicMemorySearchParamsSmartMemoryLocationModuleID) MarshalJSON() (data []byte, err error) {
	type shadow QueryEpisodicMemorySearchParamsSmartMemoryLocationModuleID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QueryEpisodicMemorySearchParamsSmartMemoryLocationModuleID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SmartMemory is required.
type QueryEpisodicMemorySearchParamsSmartMemoryLocationSmartMemory struct {
	// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
	// **REQUIRED** FALSE
	SmartMemory shared.LiquidmetalV1alpha1SmartMemoryNameParam `json:"smartMemory,omitzero,required"`
	paramObj
}

func (r QueryEpisodicMemorySearchParamsSmartMemoryLocationSmartMemory) MarshalJSON() (data []byte, err error) {
	type shadow QueryEpisodicMemorySearchParamsSmartMemoryLocationSmartMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QueryEpisodicMemorySearchParamsSmartMemoryLocationSmartMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
