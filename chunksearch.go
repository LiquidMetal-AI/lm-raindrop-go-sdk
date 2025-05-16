// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package raindrop

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/raindrop-go/internal/apijson"
	"github.com/stainless-sdks/raindrop-go/internal/requestconfig"
	"github.com/stainless-sdks/raindrop-go/option"
	"github.com/stainless-sdks/raindrop-go/packages/param"
	"github.com/stainless-sdks/raindrop-go/packages/respjson"
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
func (r *ChunkSearchService) Execute(ctx context.Context, body ChunkSearchExecuteParams, opts ...option.RequestOption) (res *ChunkSearchExecuteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/chunk_search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type TextResult struct {
	// **DESCRIPTION** Unique identifier for this text segment. Used for deduplication
	// and result tracking **EXAMPLE**
	// "51abb575a5e438a2db5fa064611995dfd76aa14d9e4b2a44c29a6374203126a5"
	ChunkSignature string `json:"chunk_signature,nullable"`
	// **DESCRIPTION** Vector representation for similarity matching. Used in semantic
	// search operations **EXAMPLE** "base64_encoded_vector_data"
	Embed string `json:"embed,nullable"`
	// **DESCRIPTION** Parent document identifier. Links related content chunks
	// together **EXAMPLE**
	// "e2ec3b118e205ff5d627e0c866224a25ba52e6d3ab758a3ef3d49e80908d7444"
	PayloadSignature string `json:"payload_signature,nullable"`
	// **DESCRIPTION** Relevance score (0.0 to 1.0). Higher scores indicate better
	// matches **EXAMPLE** 0.95
	Score float64 `json:"score,nullable"`
	// **DESCRIPTION** Source document references. Contains bucket and object
	// information **EXAMPLE** {"bucket": {"moduleId": "01jt3vs2nyt2xwk2f54x2bkn84",
	// "bucketName": "mr-bucket"}, "object": "document.pdf"}
	Source TextResultSource `json:"source"`
	// **DESCRIPTION** The actual content of the result. May be a document excerpt or
	// full content **EXAMPLE** "This is a sample text chunk from the document"
	Text string `json:"text,nullable"`
	// **DESCRIPTION** Content MIME type. Helps with proper result rendering
	// **EXAMPLE** "application/pdf"
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

// **DESCRIPTION** Source document references. Contains bucket and object
// information **EXAMPLE** {"bucket": {"moduleId": "01jt3vs2nyt2xwk2f54x2bkn84",
// "bucketName": "mr-bucket"}, "object": "document.pdf"}
type TextResultSource struct {
	// **DESCRIPTION** The bucket information containing this result **EXAMPLE**
	// {"moduleId": "01jt3vs2nyt2xwk2f54x2bkn84", "bucketName": "mr-bucket"}
	Bucket BucketResponse `json:"bucket"`
	// **DESCRIPTION** The object key within the bucket **EXAMPLE** "document.pdf"
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

type ChunkSearchExecuteResponse struct {
	// **DESCRIPTION** Ordered list of relevant text segments. Each result includes
	// full context and metadata **EXAMPLE** [{"chunk_signature": "chunk_123abc",
	// "text": "Sample text", "score": 0.95}]
	Results []TextResult `json:"results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChunkSearchExecuteResponse) RawJSON() string { return r.JSON.raw }
func (r *ChunkSearchExecuteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChunkSearchExecuteParams struct {
	// **DESCRIPTION** Natural language query or question. Can include complex criteria
	// and relationships. The system will optimize the search strategy based on this
	// input **EXAMPLE** "Find documents about revenue in Q4 2023" **REQUIRED** TRUE
	Input          param.Opt[string] `json:"input,omitzero"`
	OrganizationID param.Opt[string] `json:"organization_id,omitzero"`
	// **DESCRIPTION** Client-provided search session identifier. Required for
	// pagination and result tracking. We recommend using a UUID or ULID for this value
	// **EXAMPLE** "123e4567-e89b-12d3-a456-426614174000" **REQUIRED** TRUE
	RequestID param.Opt[string] `json:"request_id,omitzero"`
	UserID    param.Opt[string] `json:"user_id,omitzero"`
	// **DESCRIPTION** The buckets to search. If provided, the search will only return
	// results from these buckets **EXAMPLE** [{"bucket": {"name": "my-bucket",
	// "version": "01jtgtraw3b5qbahrhvrj3ygbb", "application_name": "my-app"}}]
	// **REQUIRED** TRUE
	BucketLocations []BucketLocatorUnionParam `json:"bucket_locations,omitzero"`
	paramObj
}

func (r ChunkSearchExecuteParams) MarshalJSON() (data []byte, err error) {
	type shadow ChunkSearchExecuteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChunkSearchExecuteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
