// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package raindrop

import (
	"github.com/stainless-sdks/raindrop-go/internal/apijson"
	"github.com/stainless-sdks/raindrop-go/option"
	"github.com/stainless-sdks/raindrop-go/packages/respjson"
)

// ObjectService contains methods and other services that help with interacting
// with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectService] method instead.
type ObjectService struct {
	Options []option.RequestOption
}

// NewObjectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewObjectService(opts ...option.RequestOption) (r ObjectService) {
	r = ObjectService{}
	r.Options = opts
	return
}

type BucketResponse struct {
	ApplicationName      string `json:"application_name"`
	ApplicationVersionID string `json:"application_version_id"`
	BucketName           string `json:"bucket_name"`
	ModuleID             string `json:"module_id"`
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
func (r BucketResponse) RawJSON() string { return r.JSON.raw }
func (r *BucketResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
