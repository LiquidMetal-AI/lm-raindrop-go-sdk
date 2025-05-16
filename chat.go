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

// ChatService contains methods and other services that help with interacting with
// the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChatService] method instead.
type ChatService struct {
	Options []option.RequestOption
}

// NewChatService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewChatService(opts ...option.RequestOption) (r ChatService) {
	r = ChatService{}
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
func (r *ChatService) Interact(ctx context.Context, body ChatInteractParams, opts ...option.RequestOption) (res *ChatInteractResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/chat"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ChatInteractResponse struct {
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
func (r ChatInteractResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatInteractResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatInteractParams struct {
	// User's input or question about the document. Can be natural language questions,
	// commands, or requests. The system will process this against the document content
	Input param.Opt[string] `json:"input,omitzero"`
	// Document identifier within the bucket. Typically matches the storage path or
	// key. Used to identify which document to chat with
	ObjectID param.Opt[string] `json:"object_id,omitzero"`
	// Client-provided conversation session identifier. Required for maintaining
	// context in follow-up questions. We recommend using a UUID or ULID for this value
	RequestID param.Opt[string] `json:"request_id,omitzero"`
	// The storage bucket containing the target document. Must be a valid, registered
	// Smart Bucket. Used to identify which bucket to query against
	BucketLocation ChatInteractParamsBucketLocationUnion `json:"bucket_location,omitzero"`
	paramObj
}

func (r ChatInteractParams) MarshalJSON() (data []byte, err error) {
	type shadow ChatInteractParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatInteractParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ChatInteractParamsBucketLocationUnion struct {
	OfBucket   *ChatInteractParamsBucketLocationBucket   `json:",omitzero,inline"`
	OfModuleID *ChatInteractParamsBucketLocationModuleID `json:",omitzero,inline"`
	paramUnion
}

func (u ChatInteractParamsBucketLocationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion[ChatInteractParamsBucketLocationUnion](u.OfBucket, u.OfModuleID)
}
func (u *ChatInteractParamsBucketLocationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ChatInteractParamsBucketLocationUnion) asAny() any {
	if !param.IsOmitted(u.OfBucket) {
		return u.OfBucket
	} else if !param.IsOmitted(u.OfModuleID) {
		return u.OfModuleID
	}
	return nil
}

// The property Bucket is required.
type ChatInteractParamsBucketLocationBucket struct {
	// BucketName represents a bucket name with an optional version
	Bucket ChatInteractParamsBucketLocationBucketBucket `json:"bucket,omitzero,required"`
	paramObj
}

func (r ChatInteractParamsBucketLocationBucket) MarshalJSON() (data []byte, err error) {
	type shadow ChatInteractParamsBucketLocationBucket
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatInteractParamsBucketLocationBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BucketName represents a bucket name with an optional version
type ChatInteractParamsBucketLocationBucketBucket struct {
	// Optional Application
	ApplicationName param.Opt[string] `json:"application_name,omitzero"`
	// Optional version of the bucket
	Version param.Opt[string] `json:"version,omitzero"`
	// The name of the bucket
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r ChatInteractParamsBucketLocationBucketBucket) MarshalJSON() (data []byte, err error) {
	type shadow ChatInteractParamsBucketLocationBucketBucket
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatInteractParamsBucketLocationBucketBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ModuleID is required.
type ChatInteractParamsBucketLocationModuleID struct {
	ModuleID string `json:"module_id,required"`
	paramObj
}

func (r ChatInteractParamsBucketLocationModuleID) MarshalJSON() (data []byte, err error) {
	type shadow ChatInteractParamsBucketLocationModuleID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatInteractParamsBucketLocationModuleID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
