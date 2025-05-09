// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package raindrop

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/apijson"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/requestconfig"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/option"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/respjson"
)

// StorageObjectService contains methods and other services that help with
// interacting with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStorageObjectService] method instead.
type StorageObjectService struct {
	Options []option.RequestOption
}

// NewStorageObjectService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStorageObjectService(opts ...option.RequestOption) (r StorageObjectService) {
	r = StorageObjectService{}
	r.Options = opts
	return
}

// Delete a file from a SmartBucket or regular bucket. The bucket parameter (ID) is
// used to identify the bucket to delete from. The key is the path to the object in
// the bucket.
func (r *StorageObjectService) Delete(ctx context.Context, key string, body StorageObjectDeleteParams, opts ...option.RequestOption) (res *StorageObjectDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.Bucket == "" {
		err = errors.New("missing required bucket parameter")
		return
	}
	if key == "" {
		err = errors.New("missing required key parameter")
		return
	}
	path := fmt.Sprintf("v1/object/%s/%s", body.Bucket, key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type StorageObjectDeleteResponse struct {
	Success bool `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageObjectDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageObjectDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageObjectDeleteParams struct {
	Bucket string `path:"bucket,required" json:"-"`
	paramObj
}
