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

// DocumentStatusBulkService contains methods and other services that help with
// interacting with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDocumentStatusBulkService] method instead.
type DocumentStatusBulkService struct {
	Options []option.RequestOption
}

// NewDocumentStatusBulkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDocumentStatusBulkService(opts ...option.RequestOption) (r DocumentStatusBulkService) {
	r = DocumentStatusBulkService{}
	r.Options = opts
	return
}

// Get the indexing status for multiple documents in a single request. This is
// significantly more efficient than making individual GetDocumentStatus calls, as
// it searches shards once and returns status for all requested documents.
func (r *DocumentStatusBulkService) GetStatusBulk(ctx context.Context, body DocumentStatusBulkGetStatusBulkParams, opts ...option.RequestOption) (res *DocumentStatusBulkGetStatusBulkResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/document_status_bulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type DocumentStatusBulkGetStatusBulkResponse struct {
	// Status information for each requested document, keyed by object_id
	Documents []DocumentStatusBulkGetStatusBulkResponseDocument `json:"documents"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Documents   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentStatusBulkGetStatusBulkResponse) RawJSON() string { return r.JSON.raw }
func (r *DocumentStatusBulkGetStatusBulkResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentStatusBulkGetStatusBulkResponseDocument struct {
	// Embedding stage information
	Embedding DocumentStatusBulkGetStatusBulkResponseDocumentEmbedding `json:"embedding,nullable"`
	// Any errors encountered during indexing
	Errors []string `json:"errors"`
	// Ingest stage information
	Ingest DocumentStatusBulkGetStatusBulkResponseDocumentIngest `json:"ingest,nullable"`
	// Document identifier (object key)
	ObjectID string `json:"objectId"`
	// PII detection stage information
	Pii DocumentStatusBulkGetStatusBulkResponseDocumentPii `json:"pii,nullable"`
	// Relationships stage information
	Relationships DocumentStatusBulkGetStatusBulkResponseDocumentRelationships `json:"relationships,nullable"`
	// Overall document status
	Status string `json:"status"`
	// Vector index stage information
	VectorIndex DocumentStatusBulkGetStatusBulkResponseDocumentVectorIndex `json:"vectorIndex,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Embedding     respjson.Field
		Errors        respjson.Field
		Ingest        respjson.Field
		ObjectID      respjson.Field
		Pii           respjson.Field
		Relationships respjson.Field
		Status        respjson.Field
		VectorIndex   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentStatusBulkGetStatusBulkResponseDocument) RawJSON() string { return r.JSON.raw }
func (r *DocumentStatusBulkGetStatusBulkResponseDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Embedding stage information
type DocumentStatusBulkGetStatusBulkResponseDocumentEmbedding struct {
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
func (r DocumentStatusBulkGetStatusBulkResponseDocumentEmbedding) RawJSON() string { return r.JSON.raw }
func (r *DocumentStatusBulkGetStatusBulkResponseDocumentEmbedding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ingest stage information
type DocumentStatusBulkGetStatusBulkResponseDocumentIngest struct {
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
func (r DocumentStatusBulkGetStatusBulkResponseDocumentIngest) RawJSON() string { return r.JSON.raw }
func (r *DocumentStatusBulkGetStatusBulkResponseDocumentIngest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PII detection stage information
type DocumentStatusBulkGetStatusBulkResponseDocumentPii struct {
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
func (r DocumentStatusBulkGetStatusBulkResponseDocumentPii) RawJSON() string { return r.JSON.raw }
func (r *DocumentStatusBulkGetStatusBulkResponseDocumentPii) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Relationships stage information
type DocumentStatusBulkGetStatusBulkResponseDocumentRelationships struct {
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
func (r DocumentStatusBulkGetStatusBulkResponseDocumentRelationships) RawJSON() string {
	return r.JSON.raw
}
func (r *DocumentStatusBulkGetStatusBulkResponseDocumentRelationships) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Vector index stage information
type DocumentStatusBulkGetStatusBulkResponseDocumentVectorIndex struct {
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
func (r DocumentStatusBulkGetStatusBulkResponseDocumentVectorIndex) RawJSON() string {
	return r.JSON.raw
}
func (r *DocumentStatusBulkGetStatusBulkResponseDocumentVectorIndex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentStatusBulkGetStatusBulkParams struct {
	// The storage bucket containing the target documents
	BucketLocation BucketLocatorParam `json:"bucketLocation,omitzero,required"`
	// List of document identifiers (object keys) to get status for
	ObjectIDs []string `json:"objectIds,omitzero,required"`
	// Optional partition identifier for multi-tenant data isolation. Defaults to
	// 'default' if not specified
	Partition param.Opt[string] `json:"partition,omitzero"`
	paramObj
}

func (r DocumentStatusBulkGetStatusBulkParams) MarshalJSON() (data []byte, err error) {
	type shadow DocumentStatusBulkGetStatusBulkParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentStatusBulkGetStatusBulkParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
