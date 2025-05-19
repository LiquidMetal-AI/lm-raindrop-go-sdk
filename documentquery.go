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
func (r *DocumentQueryService) Ask(ctx context.Context, body DocumentQueryAskParams, opts ...option.RequestOption) (res *DocumentQueryAskResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/document_query"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func BucketLocatorParamOfBucket(bucket BucketLocatorBucketBucketParam) BucketLocatorUnionParam {
	var variant BucketLocatorBucketParam
	variant.Bucket = bucket
	return BucketLocatorUnionParam{OfBucket: &variant}
}

func BucketLocatorParamOfModuleID(moduleID string) BucketLocatorUnionParam {
	var variant BucketLocatorModuleIDParam
	variant.ModuleID = moduleID
	return BucketLocatorUnionParam{OfModuleID: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BucketLocatorUnionParam struct {
	OfBucket   *BucketLocatorBucketParam   `json:",omitzero,inline"`
	OfModuleID *BucketLocatorModuleIDParam `json:",omitzero,inline"`
	paramUnion
}

func (u BucketLocatorUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[BucketLocatorUnionParam](u.OfBucket, u.OfModuleID)
}
func (u *BucketLocatorUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BucketLocatorUnionParam) asAny() any {
	if !param.IsOmitted(u.OfBucket) {
		return u.OfBucket
	} else if !param.IsOmitted(u.OfModuleID) {
		return u.OfModuleID
	}
	return nil
}

// The property Bucket is required.
type BucketLocatorBucketParam struct {
	// **EXAMPLE** { name: 'my-smartbucket' } **REQUIRED** FALSE
	Bucket BucketLocatorBucketBucketParam `json:"bucket,omitzero,required"`
	paramObj
}

func (r BucketLocatorBucketParam) MarshalJSON() (data []byte, err error) {
	type shadow BucketLocatorBucketParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BucketLocatorBucketParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **EXAMPLE** { name: 'my-smartbucket' } **REQUIRED** FALSE
//
// The property Name is required.
type BucketLocatorBucketBucketParam struct {
	// The name of the bucket **EXAMPLE** "my-bucket" **REQUIRED** TRUE
	Name string `json:"name,required"`
	// Optional Application **EXAMPLE** "my-app" **REQUIRED** FALSE
	ApplicationName param.Opt[string] `json:"application_name,omitzero"`
	// Optional version of the bucket **EXAMPLE** "01jtryx2f2f61ryk06vd8mr91p"
	// **REQUIRED** FALSE
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r BucketLocatorBucketBucketParam) MarshalJSON() (data []byte, err error) {
	type shadow BucketLocatorBucketBucketParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BucketLocatorBucketBucketParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ModuleID is required.
type BucketLocatorModuleIDParam struct {
	// **EXAMPLE** "01jtryx2f2f61ryk06vd8mr91p" **REQUIRED** FALSE
	ModuleID string `json:"module_id,required"`
	paramObj
}

func (r BucketLocatorModuleIDParam) MarshalJSON() (data []byte, err error) {
	type shadow BucketLocatorModuleIDParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BucketLocatorModuleIDParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentQueryAskResponse struct {
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
func (r DocumentQueryAskResponse) RawJSON() string { return r.JSON.raw }
func (r *DocumentQueryAskResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentQueryAskParams struct {
	// The storage bucket containing the target document. Must be a valid, registered
	// Smart Bucket. Used to identify which bucket to query against
	BucketLocation BucketLocatorUnionParam `json:"bucket_location,omitzero,required"`
	// User's input or question about the document. Can be natural language questions,
	// commands, or requests. The system will process this against the document content
	Input string `json:"input,required"`
	// Document identifier within the bucket. Typically matches the storage path or
	// key. Used to identify which document to chat with
	ObjectID string `json:"object_id,required"`
	// Client-provided conversation session identifier. Required for maintaining
	// context in follow-up questions. We recommend using a UUID or ULID for this value
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
