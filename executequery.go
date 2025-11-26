// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package raindrop

import (
	"context"
	"net/http"
	"slices"

	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/apijson"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/internal/requestconfig"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/option"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/param"
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/packages/respjson"
)

// ExecuteQueryService contains methods and other services that help with
// interacting with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExecuteQueryService] method instead.
type ExecuteQueryService struct {
	Options []option.RequestOption
}

// NewExecuteQueryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewExecuteQueryService(opts ...option.RequestOption) (r ExecuteQueryService) {
	r = ExecuteQueryService{}
	r.Options = opts
	return
}

// Executes a SQL query or converts natural language to SQL and executes it.
// Supports both direct SQL execution and AI-powered natural language to SQL
// conversion. Automatically handles metadata updates and PII detection for data
// governance.
//
// Features:
//
// - Direct SQL query execution
// - Natural language to SQL conversion using AI
// - Automatic metadata tracking for schema changes
// - PII detection for security
// - Multiple output formats (JSON, CSV)
func (r *ExecuteQueryService) Execute(ctx context.Context, body ExecuteQueryExecuteParams, opts ...option.RequestOption) (res *ExecuteQueryExecuteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/execute_query"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ExecuteQueryExecuteResponse struct {
	// AI reasoning for natural language queries (if applicable)
	AIReasoning string `json:"aiReasoning,nullable"`
	// Format of the results field
	Format string `json:"format"`
	// The actual SQL query that was executed
	QueryExecuted string `json:"queryExecuted"`
	// Query results as a string in the requested format (JSON or CSV)
	Results string `json:"results"`
	// HTTP status code for the operation
	Status int64 `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AIReasoning   respjson.Field
		Format        respjson.Field
		QueryExecuted respjson.Field
		Results       respjson.Field
		Status        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecuteQueryExecuteResponse) RawJSON() string { return r.JSON.raw }
func (r *ExecuteQueryExecuteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExecuteQueryExecuteParams struct {
	// Smart SQL locator for targeting the correct smart SQL instance
	SmartSqlLocation ExecuteQueryExecuteParamsSmartSqlLocation `json:"smartSqlLocation,omitzero,required"`
	// Direct SQL query to execute (mutually exclusive with text_query)
	SqlQuery param.Opt[string] `json:"sqlQuery,omitzero"`
	// Natural language query to convert to SQL (mutually exclusive with sql_query)
	TextQuery param.Opt[string] `json:"textQuery,omitzero"`
	// Desired output format for query results
	//
	// Any of "OUTPUT_FORMAT_UNSPECIFIED", "OUTPUT_FORMAT_JSON", "OUTPUT_FORMAT_CSV".
	Format ExecuteQueryExecuteParamsFormat `json:"format,omitzero"`
	paramObj
}

func (r ExecuteQueryExecuteParams) MarshalJSON() (data []byte, err error) {
	type shadow ExecuteQueryExecuteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExecuteQueryExecuteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property SmartSql is required.
type ExecuteQueryExecuteParamsSmartSqlLocation struct {
	// Name-based smart SQL instance identifier (recommended)
	SmartSql ExecuteQueryExecuteParamsSmartSqlLocationSmartSql `json:"smartSql,omitzero,required"`
	paramObj
}

func (r ExecuteQueryExecuteParamsSmartSqlLocation) MarshalJSON() (data []byte, err error) {
	type shadow ExecuteQueryExecuteParamsSmartSqlLocation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExecuteQueryExecuteParamsSmartSqlLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Name-based smart SQL instance identifier (recommended)
//
// The property Name is required.
type ExecuteQueryExecuteParamsSmartSqlLocationSmartSql struct {
	// The name of the smart SQL instance
	Name string `json:"name,required"`
	// Optional application name that owns this smart SQL instance
	ApplicationName param.Opt[string] `json:"applicationName,omitzero"`
	// Optional version identifier for the smart SQL instance
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r ExecuteQueryExecuteParamsSmartSqlLocationSmartSql) MarshalJSON() (data []byte, err error) {
	type shadow ExecuteQueryExecuteParamsSmartSqlLocationSmartSql
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExecuteQueryExecuteParamsSmartSqlLocationSmartSql) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Desired output format for query results
type ExecuteQueryExecuteParamsFormat string

const (
	ExecuteQueryExecuteParamsFormatOutputFormatUnspecified ExecuteQueryExecuteParamsFormat = "OUTPUT_FORMAT_UNSPECIFIED"
	ExecuteQueryExecuteParamsFormatOutputFormatJson        ExecuteQueryExecuteParamsFormat = "OUTPUT_FORMAT_JSON"
	ExecuteQueryExecuteParamsFormatOutputFormatCsv         ExecuteQueryExecuteParamsFormat = "OUTPUT_FORMAT_CSV"
)
