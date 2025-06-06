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

// AnswerService contains methods and other services that help with interacting
// with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAnswerService] method instead.
type AnswerService struct {
	Options []option.RequestOption
}

// NewAnswerService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAnswerService(opts ...option.RequestOption) (r AnswerService) {
	r = AnswerService{}
	r.Options = opts
	return
}

// Answers a question based on the entire content of a bucket. This combines a
// chunk search and an LLM to answer the question.
func (r *AnswerService) New(ctx context.Context, body AnswerNewParams, opts ...option.RequestOption) (res *AnswerNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/answer"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AnswerNewResponse struct {
	// The answer to the question based on all your documents in your bucket.
	Answer string `json:"answer"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Answer      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AnswerNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AnswerNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AnswerNewParams struct {
	// The bucket to search. Must be a valid, registered Smart Bucket.
	BucketLocation BucketLocatorUnionParam `json:"bucketLocation,omitzero,required"`
	// The question to answer.
	Input string `json:"input,required"`
	paramObj
}

func (r AnswerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AnswerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AnswerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
