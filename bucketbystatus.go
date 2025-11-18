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

// BucketByStatusService contains methods and other services that help with
// interacting with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBucketByStatusService] method instead.
type BucketByStatusService struct {
	Options []option.RequestOption
}

// NewBucketByStatusService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBucketByStatusService(opts ...option.RequestOption) (r BucketByStatusService) {
	r = BucketByStatusService{}
	r.Options = opts
	return
}

// List objects filtered by indexing status. Efficiently queries document storage
// across all shards to find objects matching (or excluding) specified statuses.
// Useful for identifying objects that need attention (e.g., failed, processing) or
// tracking indexing progress.
func (r *BucketByStatusService) ListObjects(ctx context.Context, body BucketByStatusListObjectsParams, opts ...option.RequestOption) (res *BucketByStatusListObjectsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/list_objects_by_status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type BucketByStatusListObjectsResponse struct {
	// Documents matching the status filter with their full status information
	Documents []BucketByStatusListObjectsResponseDocument `json:"documents"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Documents   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BucketByStatusListObjectsResponse) RawJSON() string { return r.JSON.raw }
func (r *BucketByStatusListObjectsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BucketByStatusListObjectsResponseDocument struct {
	// Embedding stage information
	Embedding BucketByStatusListObjectsResponseDocumentEmbedding `json:"embedding,nullable"`
	// Any errors encountered during indexing
	Errors []string `json:"errors"`
	// Ingest stage information
	Ingest BucketByStatusListObjectsResponseDocumentIngest `json:"ingest,nullable"`
	// Document identifier (object key)
	ObjectID string `json:"objectId"`
	// PII detection stage information
	Pii BucketByStatusListObjectsResponseDocumentPii `json:"pii,nullable"`
	// Relationships stage information
	Relationships BucketByStatusListObjectsResponseDocumentRelationships `json:"relationships,nullable"`
	// Overall document status
	Status string `json:"status"`
	// Vector index stage information
	VectorIndex BucketByStatusListObjectsResponseDocumentVectorIndex `json:"vectorIndex,nullable"`
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
func (r BucketByStatusListObjectsResponseDocument) RawJSON() string { return r.JSON.raw }
func (r *BucketByStatusListObjectsResponseDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Embedding stage information
type BucketByStatusListObjectsResponseDocumentEmbedding struct {
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
func (r BucketByStatusListObjectsResponseDocumentEmbedding) RawJSON() string { return r.JSON.raw }
func (r *BucketByStatusListObjectsResponseDocumentEmbedding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ingest stage information
type BucketByStatusListObjectsResponseDocumentIngest struct {
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
func (r BucketByStatusListObjectsResponseDocumentIngest) RawJSON() string { return r.JSON.raw }
func (r *BucketByStatusListObjectsResponseDocumentIngest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PII detection stage information
type BucketByStatusListObjectsResponseDocumentPii struct {
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
func (r BucketByStatusListObjectsResponseDocumentPii) RawJSON() string { return r.JSON.raw }
func (r *BucketByStatusListObjectsResponseDocumentPii) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Relationships stage information
type BucketByStatusListObjectsResponseDocumentRelationships struct {
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
func (r BucketByStatusListObjectsResponseDocumentRelationships) RawJSON() string { return r.JSON.raw }
func (r *BucketByStatusListObjectsResponseDocumentRelationships) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Vector index stage information
type BucketByStatusListObjectsResponseDocumentVectorIndex struct {
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
func (r BucketByStatusListObjectsResponseDocumentVectorIndex) RawJSON() string { return r.JSON.raw }
func (r *BucketByStatusListObjectsResponseDocumentVectorIndex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BucketByStatusListObjectsParams struct {
	// The storage bucket to query
	BucketLocation BucketLocatorParam `json:"bucketLocation,omitzero,required"`
	// Status values to filter by (e.g., "completed", "failed", "processing",
	// "ingesting", "partial", "uploading", "not_found")
	Statuses []string `json:"statuses,omitzero,required"`
	// If true, returns objects NOT matching the specified statuses (inverts the
	// filter)
	Exclude param.Opt[bool] `json:"exclude,omitzero"`
	// Partition to query (defaults to "default")
	Partition param.Opt[string] `json:"partition,omitzero"`
	// Optional prefix to filter object keys (e.g., "documents/" to only search in
	// documents folder)
	Prefix param.Opt[string] `json:"prefix,omitzero"`
	paramObj
}

func (r BucketByStatusListObjectsParams) MarshalJSON() (data []byte, err error) {
	type shadow BucketByStatusListObjectsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BucketByStatusListObjectsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
