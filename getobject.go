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

// GetObjectService contains methods and other services that help with interacting
// with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGetObjectService] method instead.
type GetObjectService struct {
	Options []option.RequestOption
}

// NewGetObjectService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGetObjectService(opts ...option.RequestOption) (r GetObjectService) {
	r = GetObjectService{}
	r.Options = opts
	return
}

// Download a file from a SmartBucket or regular bucket. The bucket parameter (ID)
// is used to identify the bucket to download from. The key is the path to the
// object in the bucket.
func (r *GetObjectService) Download(ctx context.Context, body GetObjectDownloadParams, opts ...option.RequestOption) (res *GetObjectDownloadResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/get_object"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type GetObjectDownloadResponse struct {
	// No specific comments in original for these fields directly, but they were part
	// of the original GetObjectResponse.
	Content     string `json:"content" format:"byte"`
	ContentType string `json:"content_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ContentType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetObjectDownloadResponse) RawJSON() string { return r.JSON.raw }
func (r *GetObjectDownloadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetObjectDownloadParams struct {
	// Object key/path to download
	Key param.Opt[string] `json:"key,omitzero"`
	// Module ID identifying the bucket
	ModuleID param.Opt[string] `json:"module_id,omitzero"`
	paramObj
}

func (r GetObjectDownloadParams) MarshalJSON() (data []byte, err error) {
	type shadow GetObjectDownloadParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GetObjectDownloadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
