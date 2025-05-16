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

// ChunkSearchService contains methods and other services that help with
// interacting with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChunkSearchService] method instead.
type ChunkSearchService struct {
	Options []option.RequestOption
}

// NewChunkSearchService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewChunkSearchService(opts ...option.RequestOption) (r ChunkSearchService) {
	r = ChunkSearchService{}
	r.Options = opts
	return
}

// Chunk Search provides search capabilities that serve as a complete drop-in
// replacement for traditional RAG pipelines. This system enables AI agents to
// leverage private data stored in SmartBuckets with zero additional configuration.
//
// Each input query is processed by our AI agent to determine the best way to
// search the data. The system will then return the most relevant results from the
// data ranked by relevance on the input query.
func (r *ChunkSearchService) Find(ctx context.Context, body ChunkSearchFindParams, opts ...option.RequestOption) (res *ChunkSearchFindResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/chunk_search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type TextResult struct {
	// Unique identifier for this text segment. Used for deduplication and result
	// tracking
	ChunkSignature string `json:"chunk_signature,nullable"`
	// Vector representation for similarity matching. Used in semantic search
	// operations
	Embed string `json:"embed,nullable"`
	// Parent document identifier. Links related content chunks together
	PayloadSignature string `json:"payload_signature,nullable"`
	// Relevance score (0.0 to 1.0). Higher scores indicate better matches
	Score float64 `json:"score,nullable"`
	// Source document references. Contains bucket and object information
	Source TextResultSource `json:"source"`
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
func (r TextResult) RawJSON() string { return r.JSON.raw }
func (r *TextResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Source document references. Contains bucket and object information
type TextResultSource struct {
	// The bucket information containing this result
	Bucket TextResultSourceBucket `json:"bucket"`
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
func (r TextResultSource) RawJSON() string { return r.JSON.raw }
func (r *TextResultSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The bucket information containing this result
type TextResultSourceBucket struct {
	ApplicationName      string `json:"application_name"`
	ApplicationVersionID string `json:"application_version_id"`
	BucketName           string `json:"bucket_name"`
	ModuleID             string `json:"module_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApplicationName      respjson.Field
		ApplicationVersionID respjson.Field
		BucketName           respjson.Field
		ModuleID             respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TextResultSourceBucket) RawJSON() string { return r.JSON.raw }
func (r *TextResultSourceBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChunkSearchFindResponse struct {
	// Ordered list of relevant text segments. Each result includes full context and
	// metadata
	Results []TextResult `json:"results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChunkSearchFindResponse) RawJSON() string { return r.JSON.raw }
func (r *ChunkSearchFindResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChunkSearchFindParams struct {
	// The buckets to search. If provided, the search will only return results from
	// these buckets
	BucketLocations []BucketLocatorUnionParam `json:"bucket_locations,omitzero,required"`
	// Natural language query or question. Can include complex criteria and
	// relationships. The system will optimize the search strategy based on this input
	Input string `json:"input,required"`
	// Client-provided search session identifier. Required for pagination and result
	// tracking. We recommend using a UUID or ULID for this value
	RequestID string `json:"request_id,required"`
	paramObj
}

func (r ChunkSearchFindParams) MarshalJSON() (data []byte, err error) {
	type shadow ChunkSearchFindParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChunkSearchFindParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
