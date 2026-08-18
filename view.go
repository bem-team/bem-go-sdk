// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bem

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/bem-team/bem-go-sdk/internal/apijson"
	"github.com/bem-team/bem-go-sdk/internal/apiquery"
	shimjson "github.com/bem-team/bem-go-sdk/internal/encoding/json"
	"github.com/bem-team/bem-go-sdk/internal/requestconfig"
	"github.com/bem-team/bem-go-sdk/option"
	"github.com/bem-team/bem-go-sdk/packages/pagination"
	"github.com/bem-team/bem-go-sdk/packages/param"
	"github.com/bem-team/bem-go-sdk/packages/respjson"
)

// Views are tabular projections over the `transformations` your functions produce
// — a saved query that turns raw extracted JSON into a filterable, paginatable,
// aggregatable table.
//
// ## Anatomy
//
// A view declares:
//
//   - One or more **functions** to read from (by `functionID` or `functionName`).
//   - A list of **columns**, each pinned to a `valueSchemaPath` (a JSON Pointer into
//     the function's output schema).
//   - Optional **filters** (string equality, numeric comparators, null-checks) and
//     **aggregations** (`count`, `count_distinct`, `sum`, `average`, `min`, `max`).
//
// Views are versioned: every update produces a new version, and the previous
// version remains immutable and addressable. Function types that produce
// transformations with an output schema — `extract`, `transform`, `analyze`,
// `join` — are all queryable through views; `extract` works uniformly across
// vision and OCR inputs.
//
// ## Reading data
//
//   - **`POST /v3/views/table-data`** — paginated rows of column values. Each row
//     reports the underlying event's `eventID` (the externally-stable KSUID used
//     everywhere else in V3) plus the projected column values.
//   - **`POST /v3/views/aggregation-data`** — group-by-able aggregate values across
//     the same query surface.
//
// Both endpoints take a `timeWindow` to bound the transformation set and require
// at least one `function` to read from.
//
// ViewService contains methods and other services that help with interacting with
// the bem API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewViewService] method instead.
type ViewService struct {
	options []option.RequestOption
}

// NewViewService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewViewService(opts ...option.RequestOption) (r ViewService) {
	r = ViewService{}
	r.options = opts
	return
}

// **Create a view.**
//
// A view is a tabular projection over the `transformations` produced by one or
// more functions. Each column declares a `valueSchemaPath` — a JSON Pointer path
// into the function's output schema — and the view can additionally carry filters
// and aggregations.
//
// Supported for every function type that produces correctable transformations and
// an output schema: `extract`, `transform`, `analyze`, `join`. Extract works on
// both vision (PDF/PNG/JPEG/HEIC/HEIF/WebP) and OCR-routed inputs — the resulting
// rows surface through views uniformly.
//
// The new view is created at `versionNum: 1`. Subsequent updates produce new
// versions; the version-1 configuration remains addressable.
func (r *ViewService) New(ctx context.Context, body ViewNewParams, opts ...option.RequestOption) (res *View, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/views"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// **Retrieve a view by ID.**
//
// Returns the view's current version. To inspect a historical version, fetch the
// list of versions on the View object and re-request with the desired version
// pinned (versions are immutable once created).
func (r *ViewService) Get(ctx context.Context, viewID string, opts ...option.RequestOption) (res *View, err error) {
	opts = slices.Concat(r.options, opts)
	if viewID == "" {
		err = errors.New("missing required view_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/views/%s", url.PathEscape(viewID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// **Update a view. Updates create a new version.**
//
// The previous version remains addressable and immutable. The new configuration is
// fully replacing — pass the complete view body, not a patch. The version number
// is auto-incremented.
func (r *ViewService) Update(ctx context.Context, viewID string, body ViewUpdateParams, opts ...option.RequestOption) (res *View, err error) {
	opts = slices.Concat(r.options, opts)
	if viewID == "" {
		err = errors.New("missing required view_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/views/%s", url.PathEscape(viewID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// **List views in the current environment, optionally filtered by the functions
// they read from.**
//
// Views are tabular projections over `transformations` rows: each view names one
// or more functions and a list of columns (JSON-pointer paths into
// `extractedJson`), and produces a uniform table that can be filtered, paginated,
// and aggregated.
//
// Filters AND together when combined. Pagination is cursor-based on `viewID`;
// default limit is 50, maximum 100.
func (r *ViewService) List(ctx context.Context, query ViewListParams, opts ...option.RequestOption) (res *pagination.ViewsPage[View], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v3/views"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// **List views in the current environment, optionally filtered by the functions
// they read from.**
//
// Views are tabular projections over `transformations` rows: each view names one
// or more functions and a list of columns (JSON-pointer paths into
// `extractedJson`), and produces a uniform table that can be filtered, paginated,
// and aggregated.
//
// Filters AND together when combined. Pagination is cursor-based on `viewID`;
// default limit is 50, maximum 100.
func (r *ViewService) ListAutoPaging(ctx context.Context, query ViewListParams, opts ...option.RequestOption) *pagination.ViewsPageAutoPager[View] {
	return pagination.NewViewsPageAutoPager(r.List(ctx, query, opts...))
}

// **Delete a view and every one of its versions.**
//
// Permanent. Any cached data-table or aggregation result clients have fetched
// remains valid, but subsequent calls to `POST /v3/views/table-data` or
// `POST /v3/views/aggregation-data` for this view will fail.
func (r *ViewService) Delete(ctx context.Context, viewID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if viewID == "" {
		err = errors.New("missing required view_id parameter")
		return err
	}
	path := fmt.Sprintf("v3/views/%s", url.PathEscape(viewID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// **Generate aggregation results for a view.**
//
// Executes each aggregation declared on the view against the `transformations`
// rows produced by the named functions inside the supplied `timeWindow`, applying
// the view's filters. Supported aggregation functions: `count`, `count_distinct`,
// `sum`, `average`, `min`, `max`. Grouped aggregations return up to 200 groups per
// aggregation; non-grouped aggregations return a single group with an empty
// `groupName`.
//
// As with table-data, the `functions` field is required.
func (r *ViewService) GenerateAggregationData(ctx context.Context, body ViewGenerateAggregationDataParams, opts ...option.RequestOption) (res *ViewGenerateAggregationDataResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/views/aggregation-data"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// **Generate paginated table data for a view.**
//
// Executes the view's query against `transformations` rows produced by the named
// functions inside the supplied `timeWindow`, applies the view's filters, and
// returns matching rows. Each row reports the event `eventID` (externally-stable
// KSUID) plus the projected column values.
//
// The `functions` field is required — at least one `functionID` or `functionName`
// must be supplied. `limit` defaults to 50 with a maximum of 200; `offset` is
// zero-based. The response's `totalCount` reflects the match count before
// pagination, so paging can be driven off it.
func (r *ViewService) GenerateTableData(ctx context.Context, body ViewGenerateTableDataParams, opts ...option.RequestOption) (res *ViewGenerateTableDataResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/views/table-data"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type FunctionIdentifier struct {
	// Unique identifier of function. Provide either id or name, not both.
	ID string `json:"id"`
	// Name of function. Must be UNIQUE on a per-environment basis. Provide either id
	// or name, not both.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionIdentifier) RawJSON() string { return r.JSON.raw }
func (r *FunctionIdentifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FunctionIdentifier to a FunctionIdentifierParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FunctionIdentifierParam.Overrides()
func (r FunctionIdentifier) ToParam() FunctionIdentifierParam {
	return param.Override[FunctionIdentifierParam](json.RawMessage(r.RawJSON()))
}

type FunctionIdentifierParam struct {
	// Unique identifier of function. Provide either id or name, not both.
	ID param.Opt[string] `json:"id,omitzero"`
	// Name of function. Must be UNIQUE on a per-environment basis. Provide either id
	// or name, not both.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r FunctionIdentifierParam) MarshalJSON() (data []byte, err error) {
	type shadow FunctionIdentifierParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionIdentifierParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Time window for filtering transformations in a view
//
// The properties End, Start are required.
type TimeWindowParam struct {
	// End of the time window in ISO 8601 (RFC 3339) format in UTC
	End time.Time `json:"end" api:"required" format:"date-time"`
	// Start of the time window in ISO 8601 (RFC 3339) format in UTC
	Start time.Time `json:"start" api:"required" format:"date-time"`
	paramObj
}

func (r TimeWindowParam) MarshalJSON() (data []byte, err error) {
	type shadow TimeWindowParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TimeWindowParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A view is a table visualization of transformations that allows customers to have
// insight into their transformations
type View struct {
	// List of aggregations defined for the view
	Aggregations []ViewAggregation `json:"aggregations" api:"required"`
	// List of columns in the view
	Columns []ViewColumn `json:"columns" api:"required"`
	// Current version number of the view
	CurrentVersionNum int64 `json:"currentVersionNum" api:"required"`
	// List of filters applied to the view
	Filters []ViewFilter `json:"filters" api:"required"`
	// List of functions that this view queries transformations from
	Functions []FunctionIdentifier `json:"functions" api:"required"`
	// Name of the view
	Name string `json:"name" api:"required"`
	// Unique identifier of the view
	ViewID string `json:"viewID" api:"required"`
	// Description of the view
	Description string `json:"description" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Aggregations      respjson.Field
		Columns           respjson.Field
		CurrentVersionNum respjson.Field
		Filters           respjson.Field
		Functions         respjson.Field
		Name              respjson.Field
		ViewID            respjson.Field
		Description       respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r View) RawJSON() string { return r.JSON.raw }
func (r *View) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An aggregation definition for a view
type ViewAggregation struct {
	// Aggregation function to apply to a view column
	//
	// Any of "count", "count_distinct", "sum", "average", "min", "max".
	Function ViewAggregationFunction `json:"function" api:"required"`
	// Name of the aggregation
	Name string `json:"name" api:"required"`
	// Name of the column to aggregate (required for count_distinct, sum, average, min,
	// max functions)
	AggregateColumnName string `json:"aggregateColumnName" api:"nullable"`
	// How to display the aggregation results
	//
	// Any of "table", "bar_chart", "pie_chart".
	DisplayType ViewAggregationDisplayType `json:"displayType"`
	// Name of the column to group by (optional, for grouped aggregations)
	GroupByColumnName string `json:"groupByColumnName" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Function            respjson.Field
		Name                respjson.Field
		AggregateColumnName respjson.Field
		DisplayType         respjson.Field
		GroupByColumnName   respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewAggregation) RawJSON() string { return r.JSON.raw }
func (r *ViewAggregation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ViewAggregation to a ViewAggregationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ViewAggregationParam.Overrides()
func (r ViewAggregation) ToParam() ViewAggregationParam {
	return param.Override[ViewAggregationParam](json.RawMessage(r.RawJSON()))
}

// Aggregation function to apply to a view column
type ViewAggregationFunction string

const (
	ViewAggregationFunctionCount         ViewAggregationFunction = "count"
	ViewAggregationFunctionCountDistinct ViewAggregationFunction = "count_distinct"
	ViewAggregationFunctionSum           ViewAggregationFunction = "sum"
	ViewAggregationFunctionAverage       ViewAggregationFunction = "average"
	ViewAggregationFunctionMin           ViewAggregationFunction = "min"
	ViewAggregationFunctionMax           ViewAggregationFunction = "max"
)

// How to display the aggregation results
type ViewAggregationDisplayType string

const (
	ViewAggregationDisplayTypeTable    ViewAggregationDisplayType = "table"
	ViewAggregationDisplayTypeBarChart ViewAggregationDisplayType = "bar_chart"
	ViewAggregationDisplayTypePieChart ViewAggregationDisplayType = "pie_chart"
)

// An aggregation definition for a view
//
// The properties Function, Name are required.
type ViewAggregationParam struct {
	// Aggregation function to apply to a view column
	//
	// Any of "count", "count_distinct", "sum", "average", "min", "max".
	Function ViewAggregationFunction `json:"function,omitzero" api:"required"`
	// Name of the aggregation
	Name string `json:"name" api:"required"`
	// Name of the column to aggregate (required for count_distinct, sum, average, min,
	// max functions)
	AggregateColumnName param.Opt[string] `json:"aggregateColumnName,omitzero"`
	// Name of the column to group by (optional, for grouped aggregations)
	GroupByColumnName param.Opt[string] `json:"groupByColumnName,omitzero"`
	// How to display the aggregation results
	//
	// Any of "table", "bar_chart", "pie_chart".
	DisplayType ViewAggregationDisplayType `json:"displayType,omitzero"`
	paramObj
}

func (r ViewAggregationParam) MarshalJSON() (data []byte, err error) {
	type shadow ViewAggregationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewAggregationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A column definition in a view
type ViewColumn struct {
	// Order in which this column should be displayed (0-based index)
	DisplayOrderIndex int64 `json:"displayOrderIndex" api:"required"`
	// Name of the column
	Name string `json:"name" api:"required"`
	// JSON path to the value in the transformation output schema (e.g.,
	// ["invoiceDetails", "invoiceNumber"])
	ValueSchemaPath []string `json:"valueSchemaPath" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayOrderIndex respjson.Field
		Name              respjson.Field
		ValueSchemaPath   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewColumn) RawJSON() string { return r.JSON.raw }
func (r *ViewColumn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ViewColumn to a ViewColumnParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ViewColumnParam.Overrides()
func (r ViewColumn) ToParam() ViewColumnParam {
	return param.Override[ViewColumnParam](json.RawMessage(r.RawJSON()))
}

// A column definition in a view
//
// The properties DisplayOrderIndex, Name, ValueSchemaPath are required.
type ViewColumnParam struct {
	// Order in which this column should be displayed (0-based index)
	DisplayOrderIndex int64 `json:"displayOrderIndex" api:"required"`
	// Name of the column
	Name string `json:"name" api:"required"`
	// JSON path to the value in the transformation output schema (e.g.,
	// ["invoiceDetails", "invoiceNumber"])
	ValueSchemaPath []string `json:"valueSchemaPath,omitzero" api:"required"`
	paramObj
}

func (r ViewColumnParam) MarshalJSON() (data []byte, err error) {
	type shadow ViewColumnParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewColumnParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request to create a new view or update an existing view
//
// The properties Aggregations, Columns, Filters, Functions, Name are required.
type ViewCreateParam struct {
	// List of aggregations defined for the view
	Aggregations []ViewAggregationParam `json:"aggregations,omitzero" api:"required"`
	// List of columns in the view
	Columns []ViewColumnParam `json:"columns,omitzero" api:"required"`
	// List of filters applied to the view
	Filters []ViewFilterParam `json:"filters,omitzero" api:"required"`
	// List of functions that this view queries transformations from
	Functions []FunctionIdentifierParam `json:"functions,omitzero" api:"required"`
	// Name of the view
	Name string `json:"name" api:"required"`
	// Description of the view
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r ViewCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow ViewCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewCreateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A filter to apply to a view column
type ViewFilter struct {
	// Name of the column to filter on
	ColumnName string `json:"columnName" api:"required"`
	// Type of filter to apply to a view column
	//
	// Any of "equals_string", "equals_number", "less_than_number",
	// "less_than_equal_number", "greater_than_number", "greater_than_equal_number",
	// "is_null", "is_not_null".
	FilterType ViewFilterFilterType `json:"filterType" api:"required"`
	// Numeric value for the filter (required for number filter types)
	Number float64 `json:"number" api:"nullable"`
	// String value for the filter (required for string filter types)
	String string `json:"string" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ColumnName  respjson.Field
		FilterType  respjson.Field
		Number      respjson.Field
		String      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewFilter) RawJSON() string { return r.JSON.raw }
func (r *ViewFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ViewFilter to a ViewFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ViewFilterParam.Overrides()
func (r ViewFilter) ToParam() ViewFilterParam {
	return param.Override[ViewFilterParam](json.RawMessage(r.RawJSON()))
}

// Type of filter to apply to a view column
type ViewFilterFilterType string

const (
	ViewFilterFilterTypeEqualsString           ViewFilterFilterType = "equals_string"
	ViewFilterFilterTypeEqualsNumber           ViewFilterFilterType = "equals_number"
	ViewFilterFilterTypeLessThanNumber         ViewFilterFilterType = "less_than_number"
	ViewFilterFilterTypeLessThanEqualNumber    ViewFilterFilterType = "less_than_equal_number"
	ViewFilterFilterTypeGreaterThanNumber      ViewFilterFilterType = "greater_than_number"
	ViewFilterFilterTypeGreaterThanEqualNumber ViewFilterFilterType = "greater_than_equal_number"
	ViewFilterFilterTypeIsNull                 ViewFilterFilterType = "is_null"
	ViewFilterFilterTypeIsNotNull              ViewFilterFilterType = "is_not_null"
)

// A filter to apply to a view column
//
// The properties ColumnName, FilterType are required.
type ViewFilterParam struct {
	// Name of the column to filter on
	ColumnName string `json:"columnName" api:"required"`
	// Type of filter to apply to a view column
	//
	// Any of "equals_string", "equals_number", "less_than_number",
	// "less_than_equal_number", "greater_than_number", "greater_than_equal_number",
	// "is_null", "is_not_null".
	FilterType ViewFilterFilterType `json:"filterType,omitzero" api:"required"`
	// Numeric value for the filter (required for number filter types)
	Number param.Opt[float64] `json:"number,omitzero"`
	// String value for the filter (required for string filter types)
	String param.Opt[string] `json:"string,omitzero"`
	paramObj
}

func (r ViewFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow ViewFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing aggregation data for a view
type ViewGenerateAggregationDataResponse struct {
	// Array of aggregation results
	Aggregations []ViewGenerateAggregationDataResponseAggregation `json:"aggregations" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Aggregations respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewGenerateAggregationDataResponse) RawJSON() string { return r.JSON.raw }
func (r *ViewGenerateAggregationDataResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Aggregation result for a single aggregation definition
type ViewGenerateAggregationDataResponseAggregation struct {
	// Array of group results (single group for non-grouped aggregations)
	Groups []ViewGenerateAggregationDataResponseAggregationGroup `json:"groups" api:"required"`
	// Name of the aggregation
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Groups      respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewGenerateAggregationDataResponseAggregation) RawJSON() string { return r.JSON.raw }
func (r *ViewGenerateAggregationDataResponseAggregation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single group result in an aggregation response
type ViewGenerateAggregationDataResponseAggregationGroup struct {
	// Name of the group (empty string for non-grouped aggregations)
	GroupName string `json:"groupName" api:"required"`
	// Aggregated value for this group
	Value float64 `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GroupName   respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewGenerateAggregationDataResponseAggregationGroup) RawJSON() string { return r.JSON.raw }
func (r *ViewGenerateAggregationDataResponseAggregationGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing paginated view table data
type ViewGenerateTableDataResponse struct {
	// Array of rows matching the view configuration
	Rows []ViewGenerateTableDataResponseRow `json:"rows" api:"required"`
	// Total number of rows matching the view (before pagination)
	TotalCount int64 `json:"totalCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Rows        respjson.Field
		TotalCount  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewGenerateTableDataResponse) RawJSON() string { return r.JSON.raw }
func (r *ViewGenerateTableDataResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single row in the view table data response
type ViewGenerateTableDataResponseRow struct {
	// Column entries for this row
	Columns []ViewGenerateTableDataResponseRowColumn `json:"columns" api:"required"`
	// Externally-stable KSUID of the event whose underlying transformation produced
	// this row.
	EventID string `json:"eventID" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Columns     respjson.Field
		EventID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewGenerateTableDataResponseRow) RawJSON() string { return r.JSON.raw }
func (r *ViewGenerateTableDataResponseRow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single column entry in a view table data row
type ViewGenerateTableDataResponseRowColumn struct {
	// Name of the column
	ColumnName string `json:"columnName" api:"required"`
	// Value of the column (can be any JSON type)
	Value ViewGenerateTableDataResponseRowColumnValueUnion `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ColumnName  respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewGenerateTableDataResponseRowColumn) RawJSON() string { return r.JSON.raw }
func (r *ViewGenerateTableDataResponseRowColumn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ViewGenerateTableDataResponseRowColumnValueUnion contains all possible
// properties and values from [string], [float64], [bool], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type ViewGenerateTableDataResponseRowColumnValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u ViewGenerateTableDataResponseRowColumnValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ViewGenerateTableDataResponseRowColumnValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ViewGenerateTableDataResponseRowColumnValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ViewGenerateTableDataResponseRowColumnValueUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ViewGenerateTableDataResponseRowColumnValueUnion) RawJSON() string { return u.JSON.raw }

func (r *ViewGenerateTableDataResponseRowColumnValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewNewParams struct {
	// Request to create a new view or update an existing view
	ViewCreate ViewCreateParam
	paramObj
}

func (r ViewNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ViewCreate)
}
func (r *ViewNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewUpdateParams struct {
	// Request to create a new view or update an existing view
	ViewCreate ViewCreateParam
	paramObj
}

func (r ViewUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ViewCreate)
}
func (r *ViewUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewListParams struct {
	// Cursor — a `viewID` defining your place in the list.
	EndingBefore param.Opt[string] `query:"endingBefore,omitzero" json:"-"`
	Limit        param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	// Cursor — a `viewID` defining your place in the list.
	StartingAfter param.Opt[string] `query:"startingAfter,omitzero" json:"-"`
	// Case-insensitive substring search over view names.
	ViewNameSubstring param.Opt[string] `query:"viewNameSubstring,omitzero" json:"-"`
	// Return only views that read from at least one of the named functions.
	FunctionIDs []string `query:"functionIDs,omitzero" json:"-"`
	// Return only views that read from at least one of the named functions.
	FunctionNames []string `query:"functionNames,omitzero" json:"-"`
	// Sort order over view IDs (default `asc`).
	//
	// Any of "asc", "desc".
	SortOrder ViewListParamsSortOrder `query:"sortOrder,omitzero" json:"-"`
	// Return only the specified view IDs.
	ViewIDs []string `query:"viewIDs,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ViewListParams]'s query parameters as `url.Values`.
func (r ViewListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order over view IDs (default `asc`).
type ViewListParamsSortOrder string

const (
	ViewListParamsSortOrderAsc  ViewListParamsSortOrder = "asc"
	ViewListParamsSortOrderDesc ViewListParamsSortOrder = "desc"
)

type ViewGenerateAggregationDataParams struct {
	// List of aggregations defined for the view
	Aggregations []ViewAggregationParam `json:"aggregations,omitzero" api:"required"`
	// List of columns in the view
	Columns []ViewColumnParam `json:"columns,omitzero" api:"required"`
	// List of filters applied to the view
	Filters []ViewFilterParam `json:"filters,omitzero" api:"required"`
	// List of functions that this view queries transformations from
	Functions []FunctionIdentifierParam `json:"functions,omitzero" api:"required"`
	// Name of the view
	Name string `json:"name" api:"required"`
	// Time window for filtering transformations in a view
	TimeWindow TimeWindowParam `json:"timeWindow,omitzero" api:"required"`
	// Description of the view
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r ViewGenerateAggregationDataParams) MarshalJSON() (data []byte, err error) {
	type shadow ViewGenerateAggregationDataParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewGenerateAggregationDataParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewGenerateTableDataParams struct {
	// List of aggregations defined for the view
	Aggregations []ViewAggregationParam `json:"aggregations,omitzero" api:"required"`
	// List of columns in the view
	Columns []ViewColumnParam `json:"columns,omitzero" api:"required"`
	// List of filters applied to the view
	Filters []ViewFilterParam `json:"filters,omitzero" api:"required"`
	// List of functions that this view queries transformations from
	Functions []FunctionIdentifierParam `json:"functions,omitzero" api:"required"`
	// Name of the view
	Name string `json:"name" api:"required"`
	// Time window for filtering transformations in a view
	TimeWindow TimeWindowParam `json:"timeWindow,omitzero" api:"required"`
	// Maximum number of rows to return (default: 50, max: 200)
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// Number of rows to skip for pagination
	Offset param.Opt[int64] `json:"offset,omitzero"`
	// Description of the view
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r ViewGenerateTableDataParams) MarshalJSON() (data []byte, err error) {
	type shadow ViewGenerateTableDataParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewGenerateTableDataParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
