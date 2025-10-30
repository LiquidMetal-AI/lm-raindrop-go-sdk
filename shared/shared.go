// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/apijson"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/param"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type LiquidmetalV1alpha1BucketResponse struct {
	// **EXAMPLE** "my-app"
	ApplicationName string `json:"applicationName"`
	// **EXAMPLE** "01jtryx2f2f61ryk06vd8mr91p"
	ApplicationVersionID string `json:"applicationVersionId"`
	// **EXAMPLE** "my-smartbucket"
	BucketName string `json:"bucketName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApplicationName      respjson.Field
		ApplicationVersionID respjson.Field
		BucketName           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LiquidmetalV1alpha1BucketResponse) RawJSON() string { return r.JSON.raw }
func (r *LiquidmetalV1alpha1BucketResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SmartMemoryName represents a smart memory name with an optional version
//
// The properties ApplicationName, Name, Version are required.
type LiquidmetalV1alpha1SmartMemoryNameParam struct {
	// Optional Application **EXAMPLE** "my-app" **REQUIRED** TRUE
	ApplicationName param.Opt[string] `json:"applicationName,omitzero,required"`
	// Optional version of the smart memory **EXAMPLE** "01jtryx2f2f61ryk06vd8mr91p"
	// **REQUIRED** TRUE
	Version param.Opt[string] `json:"version,omitzero,required"`
	// The name of the smart memory **EXAMPLE** "my-smartmemory" **REQUIRED** TRUE
	Name string `json:"name,required"`
	// Optional Application **EXAMPLE** "my-app" **REQUIRED** TRUE (Alias: accepts both
	// 'applicationName' and 'application_name')
	ApplicationName param.Opt[string] `json:"application_name,omitzero"`
	paramObj
}

func (r LiquidmetalV1alpha1SmartMemoryNameParam) MarshalJSON() (data []byte, err error) {
	type shadow LiquidmetalV1alpha1SmartMemoryNameParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LiquidmetalV1alpha1SmartMemoryNameParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
