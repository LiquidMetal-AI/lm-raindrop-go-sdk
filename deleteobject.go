// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package raindrop

import (
	"context"
	"net/http"
	"net/url"

	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/apiquery"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/requestconfig"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/option"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/param"
)

// DeleteObjectService contains methods and other services that help with
// interacting with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeleteObjectService] method instead.
type DeleteObjectService struct {
	Options []option.RequestOption
}

// NewDeleteObjectService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeleteObjectService(opts ...option.RequestOption) (r DeleteObjectService) {
	r = DeleteObjectService{}
	r.Options = opts
	return
}

// Delete a file from a SmartBucket or regular bucket. The bucket parameter (ID) is
// used to identify the bucket to delete from. The key is the path to the object in
// the bucket.
func (r *DeleteObjectService) Delete(ctx context.Context, body DeleteObjectDeleteParams, opts ...option.RequestOption) (res *DeleteObjectDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/delete_object"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type DeleteObjectDeleteResponse = any

type DeleteObjectDeleteParams struct {
	// Object key/path to delete
	Key param.Opt[string] `query:"key,omitzero" json:"-"`
	// Module ID identifying the bucket
	ModuleID param.Opt[string] `query:"module_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DeleteObjectDeleteParams]'s query parameters as
// `url.Values`.
func (r DeleteObjectDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
