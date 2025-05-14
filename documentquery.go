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

// DocumentQueryService contains methods and other services that help with
// interacting with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDocumentQueryService] method instead.
type DocumentQueryService struct {
	Options []option.RequestOption
}

// NewDocumentQueryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDocumentQueryService(opts ...option.RequestOption) (r DocumentQueryService) {
	r = DocumentQueryService{}
	r.Options = opts
	return
}

// Enables natural conversational interactions with documents stored in
// SmartBuckets. This endpoint allows users to ask questions, request summaries,
// and explore document content through an intuitive conversational interface. The
// system understands context and can handle complex queries about document
// contents.
//
// The query system maintains conversation context throught the `request_id`,
// enabling follow-up questions and deep exploration of document content. It works
// across all supported file types and automatically handles multi-page documents,
// making complex file interaction as simple as having a conversation.
//
// The system will:
//
// - Maintain conversation history for context when using the same `request_id`
// - Process questions against file content
// - Generate contextual, relevant responses
//
// Document query is supported for all file types, including PDFs, images, and
// audio files.
func (r *DocumentQueryService) Ask(ctx context.Context, body DocumentQueryAskParams, opts ...option.RequestOption) (res *DocumentQueryAskResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/document_query"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type DocumentQueryAskResponse struct {
	// AI-generated response that may include direct document quotes, content
	// summaries, contextual explanations, references to specific sections, and related
	// content suggestions
	Answer string `json:"answer,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Answer      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentQueryAskResponse) RawJSON() string { return r.JSON.raw }
func (r *DocumentQueryAskResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentQueryAskParams struct {
	// The storage bucket ID containing the target document. Must be an accessible
	// Smart Bucket
	Bucket string `json:"bucket,required"`
	// User's input or question about the document. Can be natural language questions,
	// commands, or requests
	Input string `json:"input,required"`
	// Document identifier within the bucket. Typically matches the storage path or key
	ObjectID string `json:"object_id,required"`
	// Client-provided conversation session identifier. Required for maintaining
	// context in follow-up questions. We recommend using a UUID or ULID for this
	// value.
	RequestID string `json:"request_id,required"`
	paramObj
}

func (r DocumentQueryAskParams) MarshalJSON() (data []byte, err error) {
	type shadow DocumentQueryAskParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentQueryAskParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
