// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package raindrop

import (
	"context"
	"net/http"

	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/apijson"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/requestconfig"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/option"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/pagination"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/param"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/respjson"
)

// QueryService contains methods and other services that help with interacting with
// the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQueryService] method instead.
type QueryService struct {
	Options        []option.RequestOption
	Memory         QueryMemoryService
	EpisodicMemory QueryEpisodicMemoryService
	Procedures     QueryProcedureService
	SemanticMemory QuerySemanticMemoryService
}

// NewQueryService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewQueryService(opts ...option.RequestOption) (r QueryService) {
	r = QueryService{}
	r.Options = opts
	r.Memory = NewQueryMemoryService(opts...)
	r.EpisodicMemory = NewQueryEpisodicMemoryService(opts...)
	r.Procedures = NewQueryProcedureService(opts...)
	r.SemanticMemory = NewQuerySemanticMemoryService(opts...)
	return
}

// Chunk Search provides search capabilities that serve as a complete drop-in
// replacement for traditional RAG pipelines. This system enables AI agents to
// leverage private data stored in SmartBuckets with zero additional configuration.
//
// Each input query is processed by our AI agent to determine the best way to
// search the data. The system will then return the most relevant results from the
// data ranked by relevance on the input query.
func (r *QueryService) ChunkSearch(ctx context.Context, body QueryChunkSearchParams, opts ...option.RequestOption) (res *QueryChunkSearchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/chunk_search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Enables natural conversational interactions with documents stored in
// SmartBuckets. This endpoint allows users to ask questions, request summaries,
// and explore document content through an intuitive conversational interface. The
// system understands context and can handle complex queries about document
// contents.
//
// The query system maintains conversation context throught the request_id,
// enabling follow-up questions and deep exploration of document content. It works
// across all supported file types and automatically handles multi-page documents,
// making complex file interaction as simple as having a conversation.
//
// The system will:
//
// - Maintain conversation history for context when using the same request_id
// - Process questions against file content
// - Generate contextual, relevant responses
//
// Document query is supported for all file types, including PDFs, images, and
// audio files.
func (r *QueryService) DocumentQuery(ctx context.Context, body QueryDocumentQueryParams, opts ...option.RequestOption) (res *QueryDocumentQueryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/document_query"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve additional pages from a previous search. This endpoint enables
// navigation through large result sets while maintaining search context and result
// relevance. Retrieving paginated results requires a valid request_id from a
// previously completed search.
func (r *QueryService) GetPaginatedSearch(ctx context.Context, body QueryGetPaginatedSearchParams, opts ...option.RequestOption) (res *pagination.PageNumber[QueryGetPaginatedSearchResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/search_get_page"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, body, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieve additional pages from a previous search. This endpoint enables
// navigation through large result sets while maintaining search context and result
// relevance. Retrieving paginated results requires a valid request_id from a
// previously completed search.
func (r *QueryService) GetPaginatedSearchAutoPaging(ctx context.Context, body QueryGetPaginatedSearchParams, opts ...option.RequestOption) *pagination.PageNumberAutoPager[QueryGetPaginatedSearchResponse] {
	return pagination.NewPageNumberAutoPager(r.GetPaginatedSearch(ctx, body, opts...))
}

// Primary search endpoint that provides advanced search capabilities across all
// document types stored in SmartBuckets.
//
// Supports recursive object search within objects, enabling nested content search
// like embedded images, text content, and personally identifiable information
// (PII).
//
// The system supports complex queries like:
//
//   - 'Show me documents containing credit card numbers or social security numbers'
//   - 'Find images of landscapes taken during sunset'
//   - 'Get documents mentioning revenue forecasts from Q4 2023'
//   - 'Find me all PDF documents that contain pictures of a cat'
//   - 'Find me all audio files that contain information about the weather in SF in
//     2024'
//
// Key capabilities:
//
// - Natural language query understanding
// - Content-based search across text, images, and audio
// - Automatic PII detection
// - Multi-modal search (text, images, audio)
func (r *QueryService) Search(ctx context.Context, body QuerySearchParams, opts ...option.RequestOption) (res *QuerySearchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Generates intelligent summaries of search result pages, helping users quickly
// understand large result sets without reading through every document. The system
// analyzes the content of all results on a given page and generates a detailed
// overview.
//
// The summary system:
//
// - Identifies key themes and topics
// - Extracts important findings
// - Highlights document relationships
// - Provides content type distribution
// - Summarizes metadata patterns
//
// This is particularly valuable when dealing with:
//
// - Large document collections
// - Mixed content types
// - Technical documentation
// - Research materials
func (r *QueryService) SumarizePage(ctx context.Context, body QuerySumarizePageParams, opts ...option.RequestOption) (res *QuerySumarizePageResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/summarize_page"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The property Bucket is required.
type BucketLocatorParam struct {
	// **EXAMPLE** { name: 'my-smartbucket' } **REQUIRED** FALSE
	Bucket BucketLocatorBucketParam `json:"bucket,omitzero,required"`
	paramObj
}

func (r BucketLocatorParam) MarshalJSON() (data []byte, err error) {
	type shadow BucketLocatorParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BucketLocatorParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **EXAMPLE** { name: 'my-smartbucket' } **REQUIRED** FALSE
//
// The property Name is required.
type BucketLocatorBucketParam struct {
	// The name of the bucket **EXAMPLE** "my-bucket" **REQUIRED** TRUE
	Name string `json:"name,required"`
	// Optional Application **EXAMPLE** "my-app" **REQUIRED** FALSE
	ApplicationName param.Opt[string] `json:"applicationName,omitzero"`
	// Optional version of the bucket **EXAMPLE** "01jtryx2f2f61ryk06vd8mr91p"
	// **REQUIRED** FALSE
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r BucketLocatorBucketParam) MarshalJSON() (data []byte, err error) {
	type shadow BucketLocatorBucketParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BucketLocatorBucketParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryChunkSearchResponse struct {
	// Ordered list of relevant text segments. Each result includes full context and
	// metadata
	Results []QueryChunkSearchResponseResult `json:"results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryChunkSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *QueryChunkSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryChunkSearchResponseResult struct {
	// Unique identifier for this text segment. Used for deduplication and result
	// tracking
	ChunkSignature string `json:"chunkSignature,nullable"`
	// Vector representation for similarity matching. Used in semantic search
	// operations
	Embed string `json:"embed,nullable"`
	// Parent document identifier. Links related content chunks together
	PayloadSignature string `json:"payloadSignature,nullable"`
	// Relevance score (0.0 to 1.0). Higher scores indicate better matches
	Score float64 `json:"score,nullable"`
	// Source document references. Contains bucket and object information
	Source QueryChunkSearchResponseResultSource `json:"source"`
	// The actual content of the result. May be a document excerpt or full content
	Text string `json:"text,nullable"`
	// Content MIME type. Helps with proper result rendering
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
func (r QueryChunkSearchResponseResult) RawJSON() string { return r.JSON.raw }
func (r *QueryChunkSearchResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Source document references. Contains bucket and object information
type QueryChunkSearchResponseResultSource struct {
	// The bucket information containing this result
	Bucket QueryChunkSearchResponseResultSourceBucket `json:"bucket"`
	// The object key within the bucket
	Object string `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bucket      respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryChunkSearchResponseResultSource) RawJSON() string { return r.JSON.raw }
func (r *QueryChunkSearchResponseResultSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The bucket information containing this result
type QueryChunkSearchResponseResultSourceBucket struct {
	// **EXAMPLE** "my-app"
	ApplicationName string `json:"applicationName"`
	// **EXAMPLE** "01jtryx2f2f61ryk06vd8mr91p"
	ApplicationVersionID string `json:"applicationVersionId"`
	// **EXAMPLE** "my-smartbucket"
	BucketName string `json:"bucketName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApplicationName      respjson.Field
		ApplicationVersionID respjson.Field
		BucketName           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryChunkSearchResponseResultSourceBucket) RawJSON() string { return r.JSON.raw }
func (r *QueryChunkSearchResponseResultSourceBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryDocumentQueryResponse struct {
	// AI-generated response that may include direct document quotes, content
	// summaries, contextual explanations, references to specific sections, and related
	// content suggestions
	Answer string `json:"answer"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Answer      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryDocumentQueryResponse) RawJSON() string { return r.JSON.raw }
func (r *QueryDocumentQueryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryGetPaginatedSearchResponse struct {
	// Unique identifier for this text segment. Used for deduplication and result
	// tracking
	ChunkSignature string `json:"chunkSignature,nullable"`
	// Vector representation for similarity matching. Used in semantic search
	// operations
	Embed string `json:"embed,nullable"`
	// Parent document identifier. Links related content chunks together
	PayloadSignature string `json:"payloadSignature,nullable"`
	// Relevance score (0.0 to 1.0). Higher scores indicate better matches
	Score float64 `json:"score,nullable"`
	// Source document references. Contains bucket and object information
	Source QueryGetPaginatedSearchResponseSource `json:"source"`
	// The actual content of the result. May be a document excerpt or full content
	Text string `json:"text,nullable"`
	// Content MIME type. Helps with proper result rendering
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
func (r QueryGetPaginatedSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *QueryGetPaginatedSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Source document references. Contains bucket and object information
type QueryGetPaginatedSearchResponseSource struct {
	// The bucket information containing this result
	Bucket QueryGetPaginatedSearchResponseSourceBucket `json:"bucket"`
	// The object key within the bucket
	Object string `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bucket      respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryGetPaginatedSearchResponseSource) RawJSON() string { return r.JSON.raw }
func (r *QueryGetPaginatedSearchResponseSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The bucket information containing this result
type QueryGetPaginatedSearchResponseSourceBucket struct {
	// **EXAMPLE** "my-app"
	ApplicationName string `json:"applicationName"`
	// **EXAMPLE** "01jtryx2f2f61ryk06vd8mr91p"
	ApplicationVersionID string `json:"applicationVersionId"`
	// **EXAMPLE** "my-smartbucket"
	BucketName string `json:"bucketName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApplicationName      respjson.Field
		ApplicationVersionID respjson.Field
		BucketName           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryGetPaginatedSearchResponseSourceBucket) RawJSON() string { return r.JSON.raw }
func (r *QueryGetPaginatedSearchResponseSourceBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QuerySearchResponse struct {
	// Pagination details for result navigation
	Pagination QuerySearchResponsePagination `json:"pagination"`
	// Matched results with metadata
	Results []QuerySearchResponseResult `json:"results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pagination  respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuerySearchResponse) RawJSON() string { return r.JSON.raw }
func (r *QuerySearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination details for result navigation
type QuerySearchResponsePagination struct {
	// Current page number (1-based)
	Page int64 `json:"page,required"`
	// Results per page. May be adjusted for performance
	PageSize int64 `json:"pageSize,required"`
	// Indicates more results available. Used for infinite scroll implementation
	HasMore bool `json:"hasMore"`
	// Total number of available results
	Total int64 `json:"total"`
	// Total available pages. Calculated as ceil(total/pageSize)
	TotalPages int64 `json:"totalPages"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Page        respjson.Field
		PageSize    respjson.Field
		HasMore     respjson.Field
		Total       respjson.Field
		TotalPages  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuerySearchResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *QuerySearchResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QuerySearchResponseResult struct {
	// Unique identifier for this text segment. Used for deduplication and result
	// tracking
	ChunkSignature string `json:"chunkSignature,nullable"`
	// Vector representation for similarity matching. Used in semantic search
	// operations
	Embed string `json:"embed,nullable"`
	// Parent document identifier. Links related content chunks together
	PayloadSignature string `json:"payloadSignature,nullable"`
	// Relevance score (0.0 to 1.0). Higher scores indicate better matches
	Score float64 `json:"score,nullable"`
	// Source document references. Contains bucket and object information
	Source QuerySearchResponseResultSource `json:"source"`
	// The actual content of the result. May be a document excerpt or full content
	Text string `json:"text,nullable"`
	// Content MIME type. Helps with proper result rendering
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
func (r QuerySearchResponseResult) RawJSON() string { return r.JSON.raw }
func (r *QuerySearchResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Source document references. Contains bucket and object information
type QuerySearchResponseResultSource struct {
	// The bucket information containing this result
	Bucket QuerySearchResponseResultSourceBucket `json:"bucket"`
	// The object key within the bucket
	Object string `json:"object"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bucket      respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuerySearchResponseResultSource) RawJSON() string { return r.JSON.raw }
func (r *QuerySearchResponseResultSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The bucket information containing this result
type QuerySearchResponseResultSourceBucket struct {
	// **EXAMPLE** "my-app"
	ApplicationName string `json:"applicationName"`
	// **EXAMPLE** "01jtryx2f2f61ryk06vd8mr91p"
	ApplicationVersionID string `json:"applicationVersionId"`
	// **EXAMPLE** "my-smartbucket"
	BucketName string `json:"bucketName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApplicationName      respjson.Field
		ApplicationVersionID respjson.Field
		BucketName           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuerySearchResponseResultSourceBucket) RawJSON() string { return r.JSON.raw }
func (r *QuerySearchResponseResultSourceBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QuerySumarizePageResponse struct {
	// AI-generated summary including key themes and topics, content type distribution,
	// important findings, and document relationships
	Summary string `json:"summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuerySumarizePageResponse) RawJSON() string { return r.JSON.raw }
func (r *QuerySumarizePageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryChunkSearchParams struct {
	// The buckets to search. If provided, the search will only return results from
	// these buckets
	BucketLocations []BucketLocatorParam `json:"bucketLocations,omitzero,required"`
	// Natural language query or question. Can include complex criteria and
	// relationships. The system will optimize the search strategy based on this input
	Input string `json:"input,required"`
	// Client-provided search session identifier. Required for pagination and result
	// tracking. We recommend using a UUID or ULID for this value
	RequestID string `json:"requestId,required"`
	paramObj
}

func (r QueryChunkSearchParams) MarshalJSON() (data []byte, err error) {
	type shadow QueryChunkSearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QueryChunkSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryDocumentQueryParams struct {
	// The storage bucket containing the target document. Must be a valid, registered
	// Smart Bucket. Used to identify which bucket to query against
	BucketLocation BucketLocatorParam `json:"bucketLocation,omitzero,required"`
	// User's input or question about the document. Can be natural language questions,
	// commands, or requests. The system will process this against the document content
	Input string `json:"input,required"`
	// Document identifier within the bucket. Typically matches the storage path or
	// key. Used to identify which document to chat with
	ObjectID string `json:"objectId,required"`
	// Client-provided conversation session identifier. Required for maintaining
	// context in follow-up questions. We recommend using a UUID or ULID for this value
	RequestID string `json:"requestId,required"`
	paramObj
}

func (r QueryDocumentQueryParams) MarshalJSON() (data []byte, err error) {
	type shadow QueryDocumentQueryParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QueryDocumentQueryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryGetPaginatedSearchParams struct {
	// Requested page number
	Page param.Opt[int64] `json:"page,omitzero,required"`
	// Results per page
	PageSize param.Opt[int64] `json:"pageSize,omitzero,required"`
	// Original search session identifier from the initial search
	RequestID string `json:"requestId,required"`
	paramObj
}

func (r QueryGetPaginatedSearchParams) MarshalJSON() (data []byte, err error) {
	type shadow QueryGetPaginatedSearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QueryGetPaginatedSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QuerySearchParams struct {
	// The buckets to search. If provided, the search will only return results from
	// these buckets
	BucketLocations []BucketLocatorParam `json:"bucketLocations,omitzero,required"`
	// Natural language search query that can include complex criteria. Supports
	// queries like finding documents with specific content types, PII, or semantic
	// meaning
	Input string `json:"input,required"`
	// Client-provided search session identifier. Required for pagination and result
	// tracking. We recommend using a UUID or ULID for this value
	RequestID string `json:"requestId,required"`
	paramObj
}

func (r QuerySearchParams) MarshalJSON() (data []byte, err error) {
	type shadow QuerySearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuerySearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QuerySumarizePageParams struct {
	// Target page number (1-based)
	Page int64 `json:"page,required"`
	// Results per page. Affects summary granularity
	PageSize int64 `json:"pageSize,required"`
	// Original search session identifier from the initial search
	RequestID string `json:"requestId,required"`
	paramObj
}

func (r QuerySumarizePageParams) MarshalJSON() (data []byte, err error) {
	type shadow QuerySumarizePageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuerySumarizePageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
