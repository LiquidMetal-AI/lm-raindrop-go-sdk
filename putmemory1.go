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
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/shared"
)

// PutMemoryService contains methods and other services that help with interacting
// with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPutMemoryService] method instead.
type PutMemoryService struct {
	Options []option.RequestOption
}

// NewPutMemoryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPutMemoryService(opts ...option.RequestOption) (r PutMemoryService) {
	r = PutMemoryService{}
	r.Options = opts
	return
}

// Stores multiple memory entries in a single batch operation. This is
// significantly more efficient than calling PutMemory multiple times, as it
// amortizes network overhead and batches internal storage operations.
//
// The system will:
//
// - Store all memories with automatic timestamping
// - Batch generate embeddings for semantic search
// - Associate memories with their specified timelines
// - Return memory IDs in the same order as the input entries
func (r *PutMemoryService) NewBatch(ctx context.Context, body PutMemoryNewBatchParams, opts ...option.RequestOption) (res *PutMemoryNewBatchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/put_memories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type PutMemoryNewBatchResponse struct {
	// Unique identifiers for each stored memory entry, in same order as input
	MemoryIDs []string `json:"memoryIds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MemoryIDs   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PutMemoryNewBatchResponse) RawJSON() string { return r.JSON.raw }
func (r *PutMemoryNewBatchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PutMemoryNewBatchParams struct {
	// Array of memory entries to store in batch
	Entries []PutMemoryNewBatchParamsEntry `json:"entries,omitzero,required"`
	// Unique session identifier for the working memory instance
	SessionID string `json:"sessionId,required"`
	// Smart memory locator for targeting the correct smart memory instance
	SmartMemoryLocation PutMemoryNewBatchParamsSmartMemoryLocation `json:"smartMemoryLocation,omitzero,required"`
	paramObj
}

func (r PutMemoryNewBatchParams) MarshalJSON() (data []byte, err error) {
	type shadow PutMemoryNewBatchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PutMemoryNewBatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Content is required.
type PutMemoryNewBatchParamsEntry struct {
	// The actual memory content to store
	Content string `json:"content,required"`
	// Agent identifier responsible for this memory
	Agent param.Opt[string] `json:"agent,omitzero"`
	// Optional key for direct memory retrieval
	Key param.Opt[string] `json:"key,omitzero"`
	// Timeline identifier for organizing related memories
	Timeline param.Opt[string] `json:"timeline,omitzero"`
	paramObj
}

func (r PutMemoryNewBatchParamsEntry) MarshalJSON() (data []byte, err error) {
	type shadow PutMemoryNewBatchParamsEntry
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PutMemoryNewBatchParamsEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SmartMemory is required.
type PutMemoryNewBatchParamsSmartMemoryLocation struct {
	// **EXAMPLE** {"name":"memory-name","application_name":"demo","version":"1234"}
	// **REQUIRED** FALSE
	SmartMemory shared.LiquidmetalV1alpha1SmartMemoryNameParam `json:"smartMemory,omitzero,required"`
	paramObj
}

func (r PutMemoryNewBatchParamsSmartMemoryLocation) MarshalJSON() (data []byte, err error) {
	type shadow PutMemoryNewBatchParamsSmartMemoryLocation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PutMemoryNewBatchParamsSmartMemoryLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
