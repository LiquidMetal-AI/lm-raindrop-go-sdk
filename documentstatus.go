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

// DocumentStatusService contains methods and other services that help with
// interacting with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDocumentStatusService] method instead.
type DocumentStatusService struct {
	Options []option.RequestOption
}

// NewDocumentStatusService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDocumentStatusService(opts ...option.RequestOption) (r DocumentStatusService) {
	r = DocumentStatusService{}
	r.Options = opts
	return
}

// Get the indexing status of a document by its object key. This endpoint returns
// the current indexing status of a document including progress through various
// processing stages.
func (r *DocumentStatusService) GetStatus(ctx context.Context, body DocumentStatusGetStatusParams, opts ...option.RequestOption) (res *DocumentStatusGetStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/document_status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type DocumentStatusGetStatusResponse struct {
	// Embedding stage information
	Embedding DocumentStatusGetStatusResponseEmbedding `json:"embedding,nullable"`
	// Any errors encountered during indexing
	Errors []string `json:"errors"`
	// Ingest stage information
	Ingest DocumentStatusGetStatusResponseIngest `json:"ingest,nullable"`
	// PII detection stage information
	Pii DocumentStatusGetStatusResponsePii `json:"pii,nullable"`
	// Relationships stage information
	Relationships DocumentStatusGetStatusResponseRelationships `json:"relationships,nullable"`
	// Overall document status
	Status string `json:"status"`
	// Vector index stage information
	VectorIndex DocumentStatusGetStatusResponseVectorIndex `json:"vectorIndex,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Embedding     respjson.Field
		Errors        respjson.Field
		Ingest        respjson.Field
		Pii           respjson.Field
		Relationships respjson.Field
		Status        respjson.Field
		VectorIndex   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentStatusGetStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *DocumentStatusGetStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Embedding stage information
type DocumentStatusGetStatusResponseEmbedding struct {
	// Number of items remaining to be processed
	ItemsRemaining int64 `json:"itemsRemaining"`
	// Total number of items expected to be processed
	TotalExpected int64 `json:"totalExpected"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemsRemaining respjson.Field
		TotalExpected  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentStatusGetStatusResponseEmbedding) RawJSON() string { return r.JSON.raw }
func (r *DocumentStatusGetStatusResponseEmbedding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ingest stage information
type DocumentStatusGetStatusResponseIngest struct {
	// Number of chunks queued for processing
	ChunksQueued int64 `json:"chunksQueued"`
	// Whether chunk creation is complete
	CreationComplete bool `json:"creationComplete"`
	// Total number of chunks created
	TotalChunksCreated int64 `json:"totalChunksCreated"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChunksQueued       respjson.Field
		CreationComplete   respjson.Field
		TotalChunksCreated respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentStatusGetStatusResponseIngest) RawJSON() string { return r.JSON.raw }
func (r *DocumentStatusGetStatusResponseIngest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PII detection stage information
type DocumentStatusGetStatusResponsePii struct {
	// Number of items remaining to be processed
	ItemsRemaining int64 `json:"itemsRemaining"`
	// Total number of items expected to be processed
	TotalExpected int64 `json:"totalExpected"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemsRemaining respjson.Field
		TotalExpected  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentStatusGetStatusResponsePii) RawJSON() string { return r.JSON.raw }
func (r *DocumentStatusGetStatusResponsePii) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Relationships stage information
type DocumentStatusGetStatusResponseRelationships struct {
	// Number of items remaining to be processed
	ItemsRemaining int64 `json:"itemsRemaining"`
	// Total number of items expected to be processed
	TotalExpected int64 `json:"totalExpected"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemsRemaining respjson.Field
		TotalExpected  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentStatusGetStatusResponseRelationships) RawJSON() string { return r.JSON.raw }
func (r *DocumentStatusGetStatusResponseRelationships) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Vector index stage information
type DocumentStatusGetStatusResponseVectorIndex struct {
	// Number of items remaining to be processed
	ItemsRemaining int64 `json:"itemsRemaining"`
	// Total number of items expected to be processed
	TotalExpected int64 `json:"totalExpected"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemsRemaining respjson.Field
		TotalExpected  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentStatusGetStatusResponseVectorIndex) RawJSON() string { return r.JSON.raw }
func (r *DocumentStatusGetStatusResponseVectorIndex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentStatusGetStatusParams struct {
	// The storage bucket containing the target document
	BucketLocation BucketLocatorParam `json:"bucketLocation,omitzero,required"`
	// Document identifier within the bucket (object key)
	ObjectID string `json:"objectId,required"`
	// Optional partition identifier for multi-tenant data isolation. Defaults to
	// 'default' if not specified
	Partition param.Opt[string] `json:"partition,omitzero"`
	paramObj
}

func (r DocumentStatusGetStatusParams) MarshalJSON() (data []byte, err error) {
	type shadow DocumentStatusGetStatusParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentStatusGetStatusParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
