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

// UpdateMetadataService contains methods and other services that help with
// interacting with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUpdateMetadataService] method instead.
type UpdateMetadataService struct {
	Options []option.RequestOption
}

// NewUpdateMetadataService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUpdateMetadataService(opts ...option.RequestOption) (r UpdateMetadataService) {
	r = UpdateMetadataService{}
	r.Options = opts
	return
}

// Updates database schema metadata manually. Allows for explicit metadata
// management when automatic detection is insufficient or needs correction.
//
// Use cases:
//
// - Manual schema corrections
// - Bulk metadata updates
// - Custom metadata annotations
func (r *UpdateMetadataService) Update(ctx context.Context, body UpdateMetadataUpdateParams, opts ...option.RequestOption) (res *UpdateMetadataUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/update_metadata"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type UpdateMetadataUpdateResponse struct {
	// Indicates whether the update was successful
	Success bool `json:"success"`
	// Number of tables updated
	TablesUpdated int64 `json:"tablesUpdated"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success       respjson.Field
		TablesUpdated respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UpdateMetadataUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *UpdateMetadataUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UpdateMetadataUpdateParams struct {
	// Smart SQL locator for targeting the correct smart SQL instance
	SmartSqlLocation UpdateMetadataUpdateParamsSmartSqlLocation `json:"smartSqlLocation,omitzero,required"`
	// Table metadata to update or create
	Tables []UpdateMetadataUpdateParamsTable `json:"tables,omitzero,required"`
	// Update mode: replace (overwrite), merge (preserve existing), or append (only new
	// entries)
	//
	// Any of "UPDATE_MODE_UNSPECIFIED", "UPDATE_MODE_REPLACE", "UPDATE_MODE_MERGE",
	// "UPDATE_MODE_APPEND".
	Mode UpdateMetadataUpdateParamsMode `json:"mode,omitzero"`
	paramObj
}

func (r UpdateMetadataUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow UpdateMetadataUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateMetadataUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SmartSql is required.
type UpdateMetadataUpdateParamsSmartSqlLocation struct {
	// Name-based smart SQL instance identifier (recommended)
	SmartSql UpdateMetadataUpdateParamsSmartSqlLocationSmartSql `json:"smartSql,omitzero,required"`
	paramObj
}

func (r UpdateMetadataUpdateParamsSmartSqlLocation) MarshalJSON() (data []byte, err error) {
	type shadow UpdateMetadataUpdateParamsSmartSqlLocation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateMetadataUpdateParamsSmartSqlLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Name-based smart SQL instance identifier (recommended)
//
// The property Name is required.
type UpdateMetadataUpdateParamsSmartSqlLocationSmartSql struct {
	// The name of the smart SQL instance
	Name string `json:"name,required"`
	// Optional application name that owns this smart SQL instance
	ApplicationName param.Opt[string] `json:"applicationName,omitzero"`
	// Optional version identifier for the smart SQL instance
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r UpdateMetadataUpdateParamsSmartSqlLocationSmartSql) MarshalJSON() (data []byte, err error) {
	type shadow UpdateMetadataUpdateParamsSmartSqlLocationSmartSql
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateMetadataUpdateParamsSmartSqlLocationSmartSql) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UpdateMetadataUpdateParamsTable struct {
	// When this table metadata was created
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// When this table metadata was last updated
	UpdatedAt param.Opt[time.Time] `json:"updatedAt,omitzero" format:"date-time"`
	// Name of the database table
	TableName param.Opt[string] `json:"tableName,omitzero"`
	// List of column information for the table
	Columns []UpdateMetadataUpdateParamsTableColumn `json:"columns,omitzero"`
	paramObj
}

func (r UpdateMetadataUpdateParamsTable) MarshalJSON() (data []byte, err error) {
	type shadow UpdateMetadataUpdateParamsTable
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateMetadataUpdateParamsTable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UpdateMetadataUpdateParamsTableColumn struct {
	// Sample data for AI context (nullable)
	SampleData param.Opt[string] `json:"sampleData,omitzero"`
	// Name of the database column
	ColumnName param.Opt[string] `json:"columnName,omitzero"`
	// Data type of the column
	DataType param.Opt[string] `json:"dataType,omitzero"`
	// Whether the column is a primary key
	IsPrimaryKey param.Opt[bool] `json:"isPrimaryKey,omitzero"`
	// Whether the column can contain null values
	Nullable param.Opt[bool] `json:"nullable,omitzero"`
	paramObj
}

func (r UpdateMetadataUpdateParamsTableColumn) MarshalJSON() (data []byte, err error) {
	type shadow UpdateMetadataUpdateParamsTableColumn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateMetadataUpdateParamsTableColumn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Update mode: replace (overwrite), merge (preserve existing), or append (only new
// entries)
type UpdateMetadataUpdateParamsMode string

const (
	UpdateMetadataUpdateParamsModeUpdateModeUnspecified UpdateMetadataUpdateParamsMode = "UPDATE_MODE_UNSPECIFIED"
	UpdateMetadataUpdateParamsModeUpdateModeReplace     UpdateMetadataUpdateParamsMode = "UPDATE_MODE_REPLACE"
	UpdateMetadataUpdateParamsModeUpdateModeMerge       UpdateMetadataUpdateParamsMode = "UPDATE_MODE_MERGE"
	UpdateMetadataUpdateParamsModeUpdateModeAppend      UpdateMetadataUpdateParamsMode = "UPDATE_MODE_APPEND"
)
