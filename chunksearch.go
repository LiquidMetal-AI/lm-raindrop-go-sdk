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

type ChunkSearchFindResponse struct {
	// Semantically relevant results with metadata and relevance scoring
	Results []TextResult `json:"results,required"`
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
	BucketLocations []ChunkSearchFindParamsBucketLocationUnion `json:"bucket_locations,omitzero,required"`
	// Natural language query or question. Can include complex criteria and
	// relationships
	Input string `json:"input,required"`
	// Client-provided search session identifier. We recommend using a UUID or ULID for
	// this value.
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChunkSearchFindParamsBucketLocationUnion struct {
	OfChunkSearchFindsBucketLocationModuleID *ChunkSearchFindParamsBucketLocationModuleID `json:",omitzero,inline"`
	OfChunkSearchFindsBucketLocationBucket   *ChunkSearchFindParamsBucketLocationBucket   `json:",omitzero,inline"`
	paramUnion
}

func (u ChunkSearchFindParamsBucketLocationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[ChunkSearchFindParamsBucketLocationUnion](u.OfChunkSearchFindsBucketLocationModuleID, u.OfChunkSearchFindsBucketLocationBucket)
}
func (u *ChunkSearchFindParamsBucketLocationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChunkSearchFindParamsBucketLocationUnion) asAny() any {
	if !param.IsOmitted(u.OfChunkSearchFindsBucketLocationModuleID) {
		return u.OfChunkSearchFindsBucketLocationModuleID
	} else if !param.IsOmitted(u.OfChunkSearchFindsBucketLocationBucket) {
		return u.OfChunkSearchFindsBucketLocationBucket
	}
	return nil
}

// The property ModuleID is required.
type ChunkSearchFindParamsBucketLocationModuleID struct {
	// Version-agnostic identifier for a module
	ModuleID string `json:"module_id,required"`
	paramObj
}

func (r ChunkSearchFindParamsBucketLocationModuleID) MarshalJSON() (data []byte, err error) {
	type shadow ChunkSearchFindParamsBucketLocationModuleID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChunkSearchFindParamsBucketLocationModuleID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Bucket is required.
type ChunkSearchFindParamsBucketLocationBucket struct {
	Bucket ChunkSearchFindParamsBucketLocationBucketBucket `json:"bucket,omitzero,required"`
	paramObj
}

func (r ChunkSearchFindParamsBucketLocationBucket) MarshalJSON() (data []byte, err error) {
	type shadow ChunkSearchFindParamsBucketLocationBucket
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChunkSearchFindParamsBucketLocationBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ApplicationName, Name, Version are required.
type ChunkSearchFindParamsBucketLocationBucketBucket struct {
	// Name of the application
	ApplicationName string `json:"application_name,required"`
	// Name of the bucket
	Name string `json:"name,required"`
	// Version of the bucket
	Version string `json:"version,required"`
	paramObj
}

func (r ChunkSearchFindParamsBucketLocationBucketBucket) MarshalJSON() (data []byte, err error) {
	type shadow ChunkSearchFindParamsBucketLocationBucketBucket
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChunkSearchFindParamsBucketLocationBucketBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
