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

// PutObjectService contains methods and other services that help with interacting
// with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPutObjectService] method instead.
type PutObjectService struct {
	Options []option.RequestOption
}

// NewPutObjectService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPutObjectService(opts ...option.RequestOption) (r PutObjectService) {
	r = PutObjectService{}
	r.Options = opts
	return
}

// Upload a file to a SmartBucket or regular bucket. The bucket parameter (ID) is
// used to identify the bucket to upload to. The key is the path to the object in
// the bucket.
func (r *PutObjectService) Upload(ctx context.Context, body PutObjectUploadParams, opts ...option.RequestOption) (res *PutObjectUploadResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/put_object"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type PutObjectUploadResponse struct {
	// Information about the bucket where the object was uploaded
	Bucket PutObjectUploadResponseBucket `json:"bucket"`
	// Key/path of the uploaded object
	Key string `json:"key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bucket      respjson.Field
		Key         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PutObjectUploadResponse) RawJSON() string { return r.JSON.raw }
func (r *PutObjectUploadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about the bucket where the object was uploaded
type PutObjectUploadResponseBucket struct {
	// **EXAMPLE** "my-app"
	ApplicationName string `json:"application_name"`
	// **EXAMPLE** "01jtryx2f2f61ryk06vd8mr91p"
	ApplicationVersionID string `json:"application_version_id"`
	// **EXAMPLE** "my-smartbucket"
	BucketName string `json:"bucket_name"`
	// **EXAMPLE** "01jtryx2f2f61ryk06vd8mr91p"
	ModuleID string `json:"module_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApplicationName      respjson.Field
		ApplicationVersionID respjson.Field
		BucketName           respjson.Field
		ModuleID             respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PutObjectUploadResponseBucket) RawJSON() string { return r.JSON.raw }
func (r *PutObjectUploadResponseBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PutObjectUploadParams struct {
	// Binary content of the object
	Content string `json:"content,required" format:"byte"`
	// MIME type of the object
	ContentType param.Opt[string] `json:"content_type,omitzero"`
	// Object key/path in the bucket
	Key param.Opt[string] `json:"key,omitzero"`
	// Module ID identifying the bucket
	ModuleID param.Opt[string] `json:"module_id,omitzero"`
	paramObj
}

func (r PutObjectUploadParams) MarshalJSON() (data []byte, err error) {
	type shadow PutObjectUploadParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PutObjectUploadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
