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
func (r *DocumentQueryService) New(ctx context.Context, body DocumentQueryNewParams, opts ...option.RequestOption) (res *DocumentQueryNewResponse, err error) {
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
	// BucketName represents a bucket name with an optional version
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

// BucketName represents a bucket name with an optional version
type BucketLocatorBucketBucketParam struct {
	// Optional Application
	ApplicationName param.Opt[string] `json:"application_name,omitzero"`
	// Optional version of the bucket
	Version param.Opt[string] `json:"version,omitzero"`
	// The name of the bucket
	Name param.Opt[string] `json:"name,omitzero"`
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

type DocumentQueryNewResponse struct {
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
func (r DocumentQueryNewResponse) RawJSON() string { return r.JSON.raw }
func (r *DocumentQueryNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentQueryNewParams struct {
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

func (r DocumentQueryNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DocumentQueryNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentQueryNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
