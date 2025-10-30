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

// GetMetadataService contains methods and other services that help with
// interacting with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGetMetadataService] method instead.
type GetMetadataService struct {
	Options []option.RequestOption
}

// NewGetMetadataService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGetMetadataService(opts ...option.RequestOption) (r GetMetadataService) {
	r = GetMetadataService{}
	r.Options = opts
	return
}

// Retrieves database schema metadata for a smart SQL instance. Returns table
// structures, column information, and sample data that can be used for AI context
// or application development.
//
// Metadata includes:
//
// - Table names and structures
// - Column names and data types
// - Sample data for AI context
// - Schema versioning information
func (r *GetMetadataService) Get(ctx context.Context, body GetMetadataGetParams, opts ...option.RequestOption) (res *GetMetadataGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/get_metadata"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type GetMetadataGetResponse struct {
	// Timestamp when metadata was last updated
	LastUpdated time.Time `json:"lastUpdated,nullable" format:"date-time"`
	// List of table metadata entries
	Tables []GetMetadataGetResponseTable `json:"tables"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastUpdated respjson.Field
		Tables      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetMetadataGetResponse) RawJSON() string { return r.JSON.raw }
func (r *GetMetadataGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetMetadataGetResponseTable struct {
	// List of column information for the table
	Columns []GetMetadataGetResponseTableColumn `json:"columns"`
	// When this table metadata was created
	CreatedAt time.Time `json:"createdAt,nullable" format:"date-time"`
	// Name of the database table
	TableName string `json:"tableName"`
	// When this table metadata was last updated
	UpdatedAt time.Time `json:"updatedAt,nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Columns     respjson.Field
		CreatedAt   respjson.Field
		TableName   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetMetadataGetResponseTable) RawJSON() string { return r.JSON.raw }
func (r *GetMetadataGetResponseTable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetMetadataGetResponseTableColumn struct {
	// Name of the database column
	ColumnName string `json:"columnName"`
	// Data type of the column
	DataType string `json:"dataType"`
	// Whether the column is a primary key
	IsPrimaryKey bool `json:"isPrimaryKey"`
	// Whether the column can contain null values
	Nullable bool `json:"nullable"`
	// Sample data for AI context (nullable)
	SampleData string `json:"sampleData,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ColumnName   respjson.Field
		DataType     respjson.Field
		IsPrimaryKey respjson.Field
		Nullable     respjson.Field
		SampleData   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetMetadataGetResponseTableColumn) RawJSON() string { return r.JSON.raw }
func (r *GetMetadataGetResponseTableColumn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetMetadataGetParams struct {
	// Smart SQL locator for targeting the correct smart SQL instance
	BodySmartSqlLocation1 GetMetadataGetParamsSmartSqlLocationUnion `json:"smartSqlLocation,omitzero,required"`
	// Optional table name to filter metadata (Alias: accepts both 'tableName' and
	// 'table_name')
	BodyTableName1 param.Opt[string] `json:"table_name,omitzero"`
	// Optional table name to filter metadata
	BodyTableName2 param.Opt[string] `json:"tableName,omitzero"`
	// Smart SQL locator for targeting the correct smart SQL instance (Alias: accepts
	// both 'smartSqlLocation' and 'smart_sql_location')
	BodySmartSqlLocation2 GetMetadataGetParamsSmartSqlLocationUnion `json:"smart_sql_location,omitzero"`
	paramObj
}

func (r GetMetadataGetParams) MarshalJSON() (data []byte, err error) {
	type shadow GetMetadataGetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GetMetadataGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type GetMetadataGetParamsSmartSqlLocationUnion struct {
	OfModuleID *GetMetadataGetParamsSmartSqlLocationModuleID `json:",omitzero,inline"`
	OfSmartSql *GetMetadataGetParamsSmartSqlLocationSmartSql `json:",omitzero,inline"`
	paramUnion
}

func (u GetMetadataGetParamsSmartSqlLocationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfModuleID, u.OfSmartSql)
}
func (u *GetMetadataGetParamsSmartSqlLocationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *GetMetadataGetParamsSmartSqlLocationUnion) asAny() any {
	if !param.IsOmitted(u.OfModuleID) {
		return u.OfModuleID
	} else if !param.IsOmitted(u.OfSmartSql) {
		return u.OfSmartSql
	}
	return nil
}

// The property ModuleID is required.
type GetMetadataGetParamsSmartSqlLocationModuleID struct {
	// Direct module ID for smart SQL instance (fallback, prefer name-based resolution)
	ModuleID string `json:"moduleId,required"`
	paramObj
}

func (r GetMetadataGetParamsSmartSqlLocationModuleID) MarshalJSON() (data []byte, err error) {
	type shadow GetMetadataGetParamsSmartSqlLocationModuleID
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GetMetadataGetParamsSmartSqlLocationModuleID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SmartSql is required.
type GetMetadataGetParamsSmartSqlLocationSmartSql struct {
	// Name-based smart SQL instance identifier (recommended)
	SmartSql GetMetadataGetParamsSmartSqlLocationSmartSqlSmartSql `json:"smartSql,omitzero,required"`
	paramObj
}

func (r GetMetadataGetParamsSmartSqlLocationSmartSql) MarshalJSON() (data []byte, err error) {
	type shadow GetMetadataGetParamsSmartSqlLocationSmartSql
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GetMetadataGetParamsSmartSqlLocationSmartSql) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Name-based smart SQL instance identifier (recommended)
//
// The property Name is required.
type GetMetadataGetParamsSmartSqlLocationSmartSqlSmartSql struct {
	// The name of the smart SQL instance
	Name string `json:"name,required"`
	// Optional application name that owns this smart SQL instance (Alias: accepts both
	// 'applicationName' and 'application_name')
	ApplicationName param.Opt[string] `json:"application_name,omitzero"`
	// Optional application name that owns this smart SQL instance
	ApplicationName param.Opt[string] `json:"applicationName,omitzero"`
	// Optional version identifier for the smart SQL instance
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r GetMetadataGetParamsSmartSqlLocationSmartSqlSmartSql) MarshalJSON() (data []byte, err error) {
	type shadow GetMetadataGetParamsSmartSqlLocationSmartSqlSmartSql
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GetMetadataGetParamsSmartSqlLocationSmartSqlSmartSql) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
