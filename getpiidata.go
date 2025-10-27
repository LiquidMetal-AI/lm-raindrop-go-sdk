// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package raindrop

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/apijson"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/requestconfig"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/option"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/param"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/respjson"
)

// GetPiiDataService contains methods and other services that help with interacting
// with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGetPiiDataService] method instead.
type GetPiiDataService struct {
	Options []option.RequestOption
}

// NewGetPiiDataService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGetPiiDataService(opts ...option.RequestOption) (r GetPiiDataService) {
	r = GetPiiDataService{}
	r.Options = opts
	return
}

// Retrieves PII detection results for specific database records. Returns detailed
// information about detected personally identifiable information for compliance
// and auditing purposes.
//
// PII information includes:
//
// - Entity types detected
// - Confidence scores
// - Character positions
// - Detection timestamps
func (r *GetPiiDataService) Get(ctx context.Context, body GetPiiDataGetParams, opts ...option.RequestOption) (res *GetPiiDataGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/get_pii_data"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type GetPiiDataGetResponse struct {
	// List of PII detection results
	PiiDetections []GetPiiDataGetResponsePiiDetection `json:"piiDetections"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PiiDetections respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetPiiDataGetResponse) RawJSON() string { return r.JSON.raw }
func (r *GetPiiDataGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetPiiDataGetResponsePiiDetection struct {
	// When the PII detection was performed
	DetectedAt time.Time `json:"detectedAt" format:"date-time"`
	// Unique identifier for this PII detection record
	DetectionID string `json:"detectionId"`
	// List of detected PII entities
	Entities []GetPiiDataGetResponsePiiDetectionEntity `json:"entities"`
	// Record identifier within the table
	RecordID string `json:"recordId"`
	// Table name where PII was detected
	TableName string `json:"tableName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DetectedAt  respjson.Field
		DetectionID respjson.Field
		Entities    respjson.Field
		RecordID    respjson.Field
		TableName   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetPiiDataGetResponsePiiDetection) RawJSON() string { return r.JSON.raw }
func (r *GetPiiDataGetResponsePiiDetection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetPiiDataGetResponsePiiDetectionEntity struct {
	// Confidence score for this detection (0.0 to 1.0)
	ConfidenceScore float64 `json:"confidenceScore"`
	// The detected text/token
	DetectedText string `json:"detectedText"`
	// End character position in the original text
	EndPosition int64 `json:"endPosition"`
	// Type of PII entity detected
	EntityType string `json:"entityType"`
	// Start character position in the original text
	StartPosition int64 `json:"startPosition"`
	// Token index in the tokenized text
	TokenIndex int64 `json:"tokenIndex"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConfidenceScore respjson.Field
		DetectedText    respjson.Field
		EndPosition     respjson.Field
		EntityType      respjson.Field
		StartPosition   respjson.Field
		TokenIndex      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetPiiDataGetResponsePiiDetectionEntity) RawJSON() string { return r.JSON.raw }
func (r *GetPiiDataGetResponsePiiDetectionEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetPiiDataGetParams struct {
	// Smart SQL locator for targeting the correct smart SQL instance
	SmartSqlLocation GetPiiDataGetParamsSmartSqlLocation `json:"smartSqlLocation,omitzero,required"`
	// Table name to retrieve PII data from
	TableName string `json:"tableName,required"`
	// Optional record identifier to filter PII data
	RecordID param.Opt[string] `json:"recordId,omitzero"`
	paramObj
}

func (r GetPiiDataGetParams) MarshalJSON() (data []byte, err error) {
	type shadow GetPiiDataGetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GetPiiDataGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SmartSql is required.
type GetPiiDataGetParamsSmartSqlLocation struct {
	// Name-based smart SQL instance identifier (recommended)
	SmartSql GetPiiDataGetParamsSmartSqlLocationSmartSql `json:"smartSql,omitzero,required"`
	paramObj
}

func (r GetPiiDataGetParamsSmartSqlLocation) MarshalJSON() (data []byte, err error) {
	type shadow GetPiiDataGetParamsSmartSqlLocation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GetPiiDataGetParamsSmartSqlLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Name-based smart SQL instance identifier (recommended)
//
// The property Name is required.
type GetPiiDataGetParamsSmartSqlLocationSmartSql struct {
	// The name of the smart SQL instance
	Name string `json:"name,required"`
	// Optional application name that owns this smart SQL instance
	ApplicationName param.Opt[string] `json:"applicationName,omitzero"`
	// Optional version identifier for the smart SQL instance
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r GetPiiDataGetParamsSmartSqlLocationSmartSql) MarshalJSON() (data []byte, err error) {
	type shadow GetPiiDataGetParamsSmartSqlLocationSmartSql
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GetPiiDataGetParamsSmartSqlLocationSmartSql) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
