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

// QuerySemanticMemoryService contains methods and other services that help with
// interacting with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQuerySemanticMemoryService] method instead.
type QuerySemanticMemoryService struct {
	Options []option.RequestOption
}

// NewQuerySemanticMemoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewQuerySemanticMemoryService(opts ...option.RequestOption) (r QuerySemanticMemoryService) {
	r = QuerySemanticMemoryService{}
	r.Options = opts
	return
}

// Searches across semantic memory documents using natural language queries. Uses
// vector embeddings and semantic similarity to find relevant knowledge documents
// regardless of exact keyword matches.
func (r *QuerySemanticMemoryService) Search(ctx context.Context, body QuerySemanticMemorySearchParams, opts ...option.RequestOption) (res *QuerySemanticMemorySearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/search_semantic_memory"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type QuerySemanticMemorySearchResponse struct {
	// Search results with matching documents
	DocumentSearchResponse QuerySemanticMemorySearchResponseDocumentSearchResponse `json:"documentSearchResponse,nullable"`
	// Error message if the search failed
	Error string `json:"error,nullable"`
	// Indicates whether the search was performed successfully
	Success bool `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DocumentSearchResponse respjson.Field
		Error                  respjson.Field
		Success                respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuerySemanticMemorySearchResponse) RawJSON() string { return r.JSON.raw }
func (r *QuerySemanticMemorySearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search results with matching documents
type QuerySemanticMemorySearchResponseDocumentSearchResponse struct {
	// List of matching documents ordered by relevance
	Results []QuerySemanticMemorySearchResponseDocumentSearchResponseResult `json:"results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuerySemanticMemorySearchResponseDocumentSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *QuerySemanticMemorySearchResponseDocumentSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QuerySemanticMemorySearchResponseDocumentSearchResponseResult struct {
	// Unique signature for this search result chunk
	ChunkSignature string `json:"chunkSignature,nullable"`
	// Embedding vector information (if available)
	Embed string `json:"embed,nullable"`
	// Payload signature for the original document
	PayloadSignature string `json:"payloadSignature,nullable"`
	// Relevance score for this search result
	Score float64 `json:"score,nullable"`
	// Source reference for the matched content
	Source string `json:"source,nullable"`
	// Matched text content from the document
	Text string `json:"text,nullable"`
	// Type of the matched content
	Type string `json:"type,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChunkSignature   respjson.Field
		Embed            respjson.Field
		PayloadSignature respjson.Field
		Score            respjson.Field
		Source           respjson.Field
		Text             respjson.Field
		Type             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuerySemanticMemorySearchResponseDocumentSearchResponseResult) RawJSON() string {
	return r.JSON.raw
}
func (r *QuerySemanticMemorySearchResponseDocumentSearchResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QuerySemanticMemorySearchParams struct {
	// Natural language search query to find relevant documents
	Needle string `json:"needle,required"`
	// Smart memory locator for targeting the correct smart memory instance
	SmartMemoryLocation QuerySemanticMemorySearchParamsSmartMemoryLocation `json:"smartMemoryLocation,omitzero,required"`
	paramObj
}

func (r QuerySemanticMemorySearchParams) MarshalJSON() (data []byte, err error) {
	type shadow QuerySemanticMemorySearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuerySemanticMemorySearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SmartMemory is required.
type QuerySemanticMemorySearchParamsSmartMemoryLocation struct {
	// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
	// **REQUIRED** FALSE
	SmartMemory QuerySemanticMemorySearchParamsSmartMemoryLocationSmartMemory `json:"smartMemory,omitzero,required"`
	paramObj
}

func (r QuerySemanticMemorySearchParamsSmartMemoryLocation) MarshalJSON() (data []byte, err error) {
	type shadow QuerySemanticMemorySearchParamsSmartMemoryLocation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuerySemanticMemorySearchParamsSmartMemoryLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
// **REQUIRED** FALSE
//
// The property Name is required.
type QuerySemanticMemorySearchParamsSmartMemoryLocationSmartMemory struct {
	// The name of the smart memory **EXAMPLE** "my-smartmemory" **REQUIRED** TRUE
	Name string `json:"name,required"`
	// Optional Application **EXAMPLE** "my-app" **REQUIRED** FALSE
	ApplicationName param.Opt[string] `json:"applicationName,omitzero"`
	// Optional version of the smart memory **EXAMPLE** "01jtryx2f2f61ryk06vd8mr91p"
	// **REQUIRED** FALSE
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r QuerySemanticMemorySearchParamsSmartMemoryLocationSmartMemory) MarshalJSON() (data []byte, err error) {
	type shadow QuerySemanticMemorySearchParamsSmartMemoryLocationSmartMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuerySemanticMemorySearchParamsSmartMemoryLocationSmartMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
