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
	"github.com/bem-team/bem-go-sdk/internal/requestconfig"
	"github.com/bem-team/bem-go-sdk/option"
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
func (r *ViewService) New(ctx context.Context, body ViewNewParams, opts ...option.RequestOption) (res *ViewNewResponse, err error) {
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
func (r *ViewService) Get(ctx context.Context, viewID string, opts ...option.RequestOption) (res *ViewGetResponse, err error) {
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
func (r *ViewService) Update(ctx context.Context, viewID string, body ViewUpdateParams, opts ...option.RequestOption) (res *ViewUpdateResponse, err error) {
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
func (r *ViewService) List(ctx context.Context, query ViewListParams, opts ...option.RequestOption) (res *ViewListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/views"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
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

// A view is a table visualization of transformations that allows customers to have
// insight into their transformations
type ViewNewResponse struct {
	// List of aggregations defined for the view
	Aggregations []ViewNewResponseAggregation `json:"aggregations" api:"required"`
	// List of columns in the view
	Columns []ViewNewResponseColumn `json:"columns" api:"required"`
	// Current version number of the view
	CurrentVersionNum int64 `json:"currentVersionNum" api:"required"`
	// List of filters applied to the view
	Filters []ViewNewResponseFilter `json:"filters" api:"required"`
	// List of functions that this view queries transformations from
	Functions []ViewNewResponseFunction `json:"functions" api:"required"`
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
func (r ViewNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ViewNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An aggregation definition for a view
type ViewNewResponseAggregation struct {
	// Aggregation function to apply to a view column
	//
	// Any of "count", "count_distinct", "sum", "average", "min", "max".
	Function string `json:"function" api:"required"`
	// Name of the aggregation
	Name string `json:"name" api:"required"`
	// Name of the column to aggregate (required for count_distinct, sum, average, min,
	// max functions)
	AggregateColumnName string `json:"aggregateColumnName" api:"nullable"`
	// How to display the aggregation results
	//
	// Any of "table", "bar_chart", "pie_chart".
	DisplayType string `json:"displayType"`
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
func (r ViewNewResponseAggregation) RawJSON() string { return r.JSON.raw }
func (r *ViewNewResponseAggregation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A column definition in a view
type ViewNewResponseColumn struct {
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
func (r ViewNewResponseColumn) RawJSON() string { return r.JSON.raw }
func (r *ViewNewResponseColumn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A filter to apply to a view column
type ViewNewResponseFilter struct {
	// Name of the column to filter on
	ColumnName string `json:"columnName" api:"required"`
	// Type of filter to apply to a view column
	//
	// Any of "equals_string", "equals_number", "less_than_number",
	// "less_than_equal_number", "greater_than_number", "greater_than_equal_number",
	// "is_null", "is_not_null".
	FilterType string `json:"filterType" api:"required"`
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
func (r ViewNewResponseFilter) RawJSON() string { return r.JSON.raw }
func (r *ViewNewResponseFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewNewResponseFunction struct {
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
func (r ViewNewResponseFunction) RawJSON() string { return r.JSON.raw }
func (r *ViewNewResponseFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A view is a table visualization of transformations that allows customers to have
// insight into their transformations
type ViewGetResponse struct {
	// List of aggregations defined for the view
	Aggregations []ViewGetResponseAggregation `json:"aggregations" api:"required"`
	// List of columns in the view
	Columns []ViewGetResponseColumn `json:"columns" api:"required"`
	// Current version number of the view
	CurrentVersionNum int64 `json:"currentVersionNum" api:"required"`
	// List of filters applied to the view
	Filters []ViewGetResponseFilter `json:"filters" api:"required"`
	// List of functions that this view queries transformations from
	Functions []ViewGetResponseFunction `json:"functions" api:"required"`
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
func (r ViewGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ViewGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An aggregation definition for a view
type ViewGetResponseAggregation struct {
	// Aggregation function to apply to a view column
	//
	// Any of "count", "count_distinct", "sum", "average", "min", "max".
	Function string `json:"function" api:"required"`
	// Name of the aggregation
	Name string `json:"name" api:"required"`
	// Name of the column to aggregate (required for count_distinct, sum, average, min,
	// max functions)
	AggregateColumnName string `json:"aggregateColumnName" api:"nullable"`
	// How to display the aggregation results
	//
	// Any of "table", "bar_chart", "pie_chart".
	DisplayType string `json:"displayType"`
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
func (r ViewGetResponseAggregation) RawJSON() string { return r.JSON.raw }
func (r *ViewGetResponseAggregation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A column definition in a view
type ViewGetResponseColumn struct {
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
func (r ViewGetResponseColumn) RawJSON() string { return r.JSON.raw }
func (r *ViewGetResponseColumn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A filter to apply to a view column
type ViewGetResponseFilter struct {
	// Name of the column to filter on
	ColumnName string `json:"columnName" api:"required"`
	// Type of filter to apply to a view column
	//
	// Any of "equals_string", "equals_number", "less_than_number",
	// "less_than_equal_number", "greater_than_number", "greater_than_equal_number",
	// "is_null", "is_not_null".
	FilterType string `json:"filterType" api:"required"`
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
func (r ViewGetResponseFilter) RawJSON() string { return r.JSON.raw }
func (r *ViewGetResponseFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewGetResponseFunction struct {
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
func (r ViewGetResponseFunction) RawJSON() string { return r.JSON.raw }
func (r *ViewGetResponseFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A view is a table visualization of transformations that allows customers to have
// insight into their transformations
type ViewUpdateResponse struct {
	// List of aggregations defined for the view
	Aggregations []ViewUpdateResponseAggregation `json:"aggregations" api:"required"`
	// List of columns in the view
	Columns []ViewUpdateResponseColumn `json:"columns" api:"required"`
	// Current version number of the view
	CurrentVersionNum int64 `json:"currentVersionNum" api:"required"`
	// List of filters applied to the view
	Filters []ViewUpdateResponseFilter `json:"filters" api:"required"`
	// List of functions that this view queries transformations from
	Functions []ViewUpdateResponseFunction `json:"functions" api:"required"`
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
func (r ViewUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ViewUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An aggregation definition for a view
type ViewUpdateResponseAggregation struct {
	// Aggregation function to apply to a view column
	//
	// Any of "count", "count_distinct", "sum", "average", "min", "max".
	Function string `json:"function" api:"required"`
	// Name of the aggregation
	Name string `json:"name" api:"required"`
	// Name of the column to aggregate (required for count_distinct, sum, average, min,
	// max functions)
	AggregateColumnName string `json:"aggregateColumnName" api:"nullable"`
	// How to display the aggregation results
	//
	// Any of "table", "bar_chart", "pie_chart".
	DisplayType string `json:"displayType"`
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
func (r ViewUpdateResponseAggregation) RawJSON() string { return r.JSON.raw }
func (r *ViewUpdateResponseAggregation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A column definition in a view
type ViewUpdateResponseColumn struct {
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
func (r ViewUpdateResponseColumn) RawJSON() string { return r.JSON.raw }
func (r *ViewUpdateResponseColumn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A filter to apply to a view column
type ViewUpdateResponseFilter struct {
	// Name of the column to filter on
	ColumnName string `json:"columnName" api:"required"`
	// Type of filter to apply to a view column
	//
	// Any of "equals_string", "equals_number", "less_than_number",
	// "less_than_equal_number", "greater_than_number", "greater_than_equal_number",
	// "is_null", "is_not_null".
	FilterType string `json:"filterType" api:"required"`
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
func (r ViewUpdateResponseFilter) RawJSON() string { return r.JSON.raw }
func (r *ViewUpdateResponseFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewUpdateResponseFunction struct {
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
func (r ViewUpdateResponseFunction) RawJSON() string { return r.JSON.raw }
func (r *ViewUpdateResponseFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing a list of views
type ViewListResponse struct {
	// Total number of views matching the query
	TotalCount int64 `json:"totalCount" api:"required"`
	// Array of views
	Views []ViewListResponseView `json:"views" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TotalCount  respjson.Field
		Views       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ViewListResponse) RawJSON() string { return r.JSON.raw }
func (r *ViewListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A view is a table visualization of transformations that allows customers to have
// insight into their transformations
type ViewListResponseView struct {
	// List of aggregations defined for the view
	Aggregations []ViewListResponseViewAggregation `json:"aggregations" api:"required"`
	// List of columns in the view
	Columns []ViewListResponseViewColumn `json:"columns" api:"required"`
	// Current version number of the view
	CurrentVersionNum int64 `json:"currentVersionNum" api:"required"`
	// List of filters applied to the view
	Filters []ViewListResponseViewFilter `json:"filters" api:"required"`
	// List of functions that this view queries transformations from
	Functions []ViewListResponseViewFunction `json:"functions" api:"required"`
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
func (r ViewListResponseView) RawJSON() string { return r.JSON.raw }
func (r *ViewListResponseView) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An aggregation definition for a view
type ViewListResponseViewAggregation struct {
	// Aggregation function to apply to a view column
	//
	// Any of "count", "count_distinct", "sum", "average", "min", "max".
	Function string `json:"function" api:"required"`
	// Name of the aggregation
	Name string `json:"name" api:"required"`
	// Name of the column to aggregate (required for count_distinct, sum, average, min,
	// max functions)
	AggregateColumnName string `json:"aggregateColumnName" api:"nullable"`
	// How to display the aggregation results
	//
	// Any of "table", "bar_chart", "pie_chart".
	DisplayType string `json:"displayType"`
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
func (r ViewListResponseViewAggregation) RawJSON() string { return r.JSON.raw }
func (r *ViewListResponseViewAggregation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A column definition in a view
type ViewListResponseViewColumn struct {
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
func (r ViewListResponseViewColumn) RawJSON() string { return r.JSON.raw }
func (r *ViewListResponseViewColumn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A filter to apply to a view column
type ViewListResponseViewFilter struct {
	// Name of the column to filter on
	ColumnName string `json:"columnName" api:"required"`
	// Type of filter to apply to a view column
	//
	// Any of "equals_string", "equals_number", "less_than_number",
	// "less_than_equal_number", "greater_than_number", "greater_than_equal_number",
	// "is_null", "is_not_null".
	FilterType string `json:"filterType" api:"required"`
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
func (r ViewListResponseViewFilter) RawJSON() string { return r.JSON.raw }
func (r *ViewListResponseViewFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewListResponseViewFunction struct {
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
func (r ViewListResponseViewFunction) RawJSON() string { return r.JSON.raw }
func (r *ViewListResponseViewFunction) UnmarshalJSON(data []byte) error {
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
	// List of aggregations defined for the view
	Aggregations []ViewNewParamsAggregation `json:"aggregations,omitzero" api:"required"`
	// List of columns in the view
	Columns []ViewNewParamsColumn `json:"columns,omitzero" api:"required"`
	// List of filters applied to the view
	Filters []ViewNewParamsFilter `json:"filters,omitzero" api:"required"`
	// List of functions that this view queries transformations from
	Functions []ViewNewParamsFunction `json:"functions,omitzero" api:"required"`
	// Name of the view
	Name string `json:"name" api:"required"`
	// Description of the view
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r ViewNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ViewNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An aggregation definition for a view
//
// The properties Function, Name are required.
type ViewNewParamsAggregation struct {
	// Aggregation function to apply to a view column
	//
	// Any of "count", "count_distinct", "sum", "average", "min", "max".
	Function string `json:"function,omitzero" api:"required"`
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
	DisplayType string `json:"displayType,omitzero"`
	paramObj
}

func (r ViewNewParamsAggregation) MarshalJSON() (data []byte, err error) {
	type shadow ViewNewParamsAggregation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewNewParamsAggregation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ViewNewParamsAggregation](
		"function", "count", "count_distinct", "sum", "average", "min", "max",
	)
	apijson.RegisterFieldValidator[ViewNewParamsAggregation](
		"displayType", "table", "bar_chart", "pie_chart",
	)
}

// A column definition in a view
//
// The properties DisplayOrderIndex, Name, ValueSchemaPath are required.
type ViewNewParamsColumn struct {
	// Order in which this column should be displayed (0-based index)
	DisplayOrderIndex int64 `json:"displayOrderIndex" api:"required"`
	// Name of the column
	Name string `json:"name" api:"required"`
	// JSON path to the value in the transformation output schema (e.g.,
	// ["invoiceDetails", "invoiceNumber"])
	ValueSchemaPath []string `json:"valueSchemaPath,omitzero" api:"required"`
	paramObj
}

func (r ViewNewParamsColumn) MarshalJSON() (data []byte, err error) {
	type shadow ViewNewParamsColumn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewNewParamsColumn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A filter to apply to a view column
//
// The properties ColumnName, FilterType are required.
type ViewNewParamsFilter struct {
	// Name of the column to filter on
	ColumnName string `json:"columnName" api:"required"`
	// Type of filter to apply to a view column
	//
	// Any of "equals_string", "equals_number", "less_than_number",
	// "less_than_equal_number", "greater_than_number", "greater_than_equal_number",
	// "is_null", "is_not_null".
	FilterType string `json:"filterType,omitzero" api:"required"`
	// Numeric value for the filter (required for number filter types)
	Number param.Opt[float64] `json:"number,omitzero"`
	// String value for the filter (required for string filter types)
	String param.Opt[string] `json:"string,omitzero"`
	paramObj
}

func (r ViewNewParamsFilter) MarshalJSON() (data []byte, err error) {
	type shadow ViewNewParamsFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewNewParamsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ViewNewParamsFilter](
		"filterType", "equals_string", "equals_number", "less_than_number", "less_than_equal_number", "greater_than_number", "greater_than_equal_number", "is_null", "is_not_null",
	)
}

type ViewNewParamsFunction struct {
	// Unique identifier of function. Provide either id or name, not both.
	ID param.Opt[string] `json:"id,omitzero"`
	// Name of function. Must be UNIQUE on a per-environment basis. Provide either id
	// or name, not both.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r ViewNewParamsFunction) MarshalJSON() (data []byte, err error) {
	type shadow ViewNewParamsFunction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewNewParamsFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewUpdateParams struct {
	// List of aggregations defined for the view
	Aggregations []ViewUpdateParamsAggregation `json:"aggregations,omitzero" api:"required"`
	// List of columns in the view
	Columns []ViewUpdateParamsColumn `json:"columns,omitzero" api:"required"`
	// List of filters applied to the view
	Filters []ViewUpdateParamsFilter `json:"filters,omitzero" api:"required"`
	// List of functions that this view queries transformations from
	Functions []ViewUpdateParamsFunction `json:"functions,omitzero" api:"required"`
	// Name of the view
	Name string `json:"name" api:"required"`
	// Description of the view
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r ViewUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ViewUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An aggregation definition for a view
//
// The properties Function, Name are required.
type ViewUpdateParamsAggregation struct {
	// Aggregation function to apply to a view column
	//
	// Any of "count", "count_distinct", "sum", "average", "min", "max".
	Function string `json:"function,omitzero" api:"required"`
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
	DisplayType string `json:"displayType,omitzero"`
	paramObj
}

func (r ViewUpdateParamsAggregation) MarshalJSON() (data []byte, err error) {
	type shadow ViewUpdateParamsAggregation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewUpdateParamsAggregation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ViewUpdateParamsAggregation](
		"function", "count", "count_distinct", "sum", "average", "min", "max",
	)
	apijson.RegisterFieldValidator[ViewUpdateParamsAggregation](
		"displayType", "table", "bar_chart", "pie_chart",
	)
}

// A column definition in a view
//
// The properties DisplayOrderIndex, Name, ValueSchemaPath are required.
type ViewUpdateParamsColumn struct {
	// Order in which this column should be displayed (0-based index)
	DisplayOrderIndex int64 `json:"displayOrderIndex" api:"required"`
	// Name of the column
	Name string `json:"name" api:"required"`
	// JSON path to the value in the transformation output schema (e.g.,
	// ["invoiceDetails", "invoiceNumber"])
	ValueSchemaPath []string `json:"valueSchemaPath,omitzero" api:"required"`
	paramObj
}

func (r ViewUpdateParamsColumn) MarshalJSON() (data []byte, err error) {
	type shadow ViewUpdateParamsColumn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewUpdateParamsColumn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A filter to apply to a view column
//
// The properties ColumnName, FilterType are required.
type ViewUpdateParamsFilter struct {
	// Name of the column to filter on
	ColumnName string `json:"columnName" api:"required"`
	// Type of filter to apply to a view column
	//
	// Any of "equals_string", "equals_number", "less_than_number",
	// "less_than_equal_number", "greater_than_number", "greater_than_equal_number",
	// "is_null", "is_not_null".
	FilterType string `json:"filterType,omitzero" api:"required"`
	// Numeric value for the filter (required for number filter types)
	Number param.Opt[float64] `json:"number,omitzero"`
	// String value for the filter (required for string filter types)
	String param.Opt[string] `json:"string,omitzero"`
	paramObj
}

func (r ViewUpdateParamsFilter) MarshalJSON() (data []byte, err error) {
	type shadow ViewUpdateParamsFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewUpdateParamsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ViewUpdateParamsFilter](
		"filterType", "equals_string", "equals_number", "less_than_number", "less_than_equal_number", "greater_than_number", "greater_than_equal_number", "is_null", "is_not_null",
	)
}

type ViewUpdateParamsFunction struct {
	// Unique identifier of function. Provide either id or name, not both.
	ID param.Opt[string] `json:"id,omitzero"`
	// Name of function. Must be UNIQUE on a per-environment basis. Provide either id
	// or name, not both.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r ViewUpdateParamsFunction) MarshalJSON() (data []byte, err error) {
	type shadow ViewUpdateParamsFunction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewUpdateParamsFunction) UnmarshalJSON(data []byte) error {
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
	Aggregations []ViewGenerateAggregationDataParamsAggregation `json:"aggregations,omitzero" api:"required"`
	// List of columns in the view
	Columns []ViewGenerateAggregationDataParamsColumn `json:"columns,omitzero" api:"required"`
	// List of filters applied to the view
	Filters []ViewGenerateAggregationDataParamsFilter `json:"filters,omitzero" api:"required"`
	// List of functions that this view queries transformations from
	Functions []ViewGenerateAggregationDataParamsFunction `json:"functions,omitzero" api:"required"`
	// Name of the view
	Name string `json:"name" api:"required"`
	// Time window for filtering transformations in a view
	TimeWindow ViewGenerateAggregationDataParamsTimeWindow `json:"timeWindow,omitzero" api:"required"`
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

// An aggregation definition for a view
//
// The properties Function, Name are required.
type ViewGenerateAggregationDataParamsAggregation struct {
	// Aggregation function to apply to a view column
	//
	// Any of "count", "count_distinct", "sum", "average", "min", "max".
	Function string `json:"function,omitzero" api:"required"`
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
	DisplayType string `json:"displayType,omitzero"`
	paramObj
}

func (r ViewGenerateAggregationDataParamsAggregation) MarshalJSON() (data []byte, err error) {
	type shadow ViewGenerateAggregationDataParamsAggregation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewGenerateAggregationDataParamsAggregation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ViewGenerateAggregationDataParamsAggregation](
		"function", "count", "count_distinct", "sum", "average", "min", "max",
	)
	apijson.RegisterFieldValidator[ViewGenerateAggregationDataParamsAggregation](
		"displayType", "table", "bar_chart", "pie_chart",
	)
}

// A column definition in a view
//
// The properties DisplayOrderIndex, Name, ValueSchemaPath are required.
type ViewGenerateAggregationDataParamsColumn struct {
	// Order in which this column should be displayed (0-based index)
	DisplayOrderIndex int64 `json:"displayOrderIndex" api:"required"`
	// Name of the column
	Name string `json:"name" api:"required"`
	// JSON path to the value in the transformation output schema (e.g.,
	// ["invoiceDetails", "invoiceNumber"])
	ValueSchemaPath []string `json:"valueSchemaPath,omitzero" api:"required"`
	paramObj
}

func (r ViewGenerateAggregationDataParamsColumn) MarshalJSON() (data []byte, err error) {
	type shadow ViewGenerateAggregationDataParamsColumn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewGenerateAggregationDataParamsColumn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A filter to apply to a view column
//
// The properties ColumnName, FilterType are required.
type ViewGenerateAggregationDataParamsFilter struct {
	// Name of the column to filter on
	ColumnName string `json:"columnName" api:"required"`
	// Type of filter to apply to a view column
	//
	// Any of "equals_string", "equals_number", "less_than_number",
	// "less_than_equal_number", "greater_than_number", "greater_than_equal_number",
	// "is_null", "is_not_null".
	FilterType string `json:"filterType,omitzero" api:"required"`
	// Numeric value for the filter (required for number filter types)
	Number param.Opt[float64] `json:"number,omitzero"`
	// String value for the filter (required for string filter types)
	String param.Opt[string] `json:"string,omitzero"`
	paramObj
}

func (r ViewGenerateAggregationDataParamsFilter) MarshalJSON() (data []byte, err error) {
	type shadow ViewGenerateAggregationDataParamsFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewGenerateAggregationDataParamsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ViewGenerateAggregationDataParamsFilter](
		"filterType", "equals_string", "equals_number", "less_than_number", "less_than_equal_number", "greater_than_number", "greater_than_equal_number", "is_null", "is_not_null",
	)
}

type ViewGenerateAggregationDataParamsFunction struct {
	// Unique identifier of function. Provide either id or name, not both.
	ID param.Opt[string] `json:"id,omitzero"`
	// Name of function. Must be UNIQUE on a per-environment basis. Provide either id
	// or name, not both.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r ViewGenerateAggregationDataParamsFunction) MarshalJSON() (data []byte, err error) {
	type shadow ViewGenerateAggregationDataParamsFunction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewGenerateAggregationDataParamsFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Time window for filtering transformations in a view
//
// The properties End, Start are required.
type ViewGenerateAggregationDataParamsTimeWindow struct {
	// End of the time window in ISO 8601 (RFC 3339) format in UTC
	End time.Time `json:"end" api:"required" format:"date-time"`
	// Start of the time window in ISO 8601 (RFC 3339) format in UTC
	Start time.Time `json:"start" api:"required" format:"date-time"`
	paramObj
}

func (r ViewGenerateAggregationDataParamsTimeWindow) MarshalJSON() (data []byte, err error) {
	type shadow ViewGenerateAggregationDataParamsTimeWindow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewGenerateAggregationDataParamsTimeWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ViewGenerateTableDataParams struct {
	// List of aggregations defined for the view
	Aggregations []ViewGenerateTableDataParamsAggregation `json:"aggregations,omitzero" api:"required"`
	// List of columns in the view
	Columns []ViewGenerateTableDataParamsColumn `json:"columns,omitzero" api:"required"`
	// List of filters applied to the view
	Filters []ViewGenerateTableDataParamsFilter `json:"filters,omitzero" api:"required"`
	// List of functions that this view queries transformations from
	Functions []ViewGenerateTableDataParamsFunction `json:"functions,omitzero" api:"required"`
	// Name of the view
	Name string `json:"name" api:"required"`
	// Time window for filtering transformations in a view
	TimeWindow ViewGenerateTableDataParamsTimeWindow `json:"timeWindow,omitzero" api:"required"`
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

// An aggregation definition for a view
//
// The properties Function, Name are required.
type ViewGenerateTableDataParamsAggregation struct {
	// Aggregation function to apply to a view column
	//
	// Any of "count", "count_distinct", "sum", "average", "min", "max".
	Function string `json:"function,omitzero" api:"required"`
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
	DisplayType string `json:"displayType,omitzero"`
	paramObj
}

func (r ViewGenerateTableDataParamsAggregation) MarshalJSON() (data []byte, err error) {
	type shadow ViewGenerateTableDataParamsAggregation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewGenerateTableDataParamsAggregation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ViewGenerateTableDataParamsAggregation](
		"function", "count", "count_distinct", "sum", "average", "min", "max",
	)
	apijson.RegisterFieldValidator[ViewGenerateTableDataParamsAggregation](
		"displayType", "table", "bar_chart", "pie_chart",
	)
}

// A column definition in a view
//
// The properties DisplayOrderIndex, Name, ValueSchemaPath are required.
type ViewGenerateTableDataParamsColumn struct {
	// Order in which this column should be displayed (0-based index)
	DisplayOrderIndex int64 `json:"displayOrderIndex" api:"required"`
	// Name of the column
	Name string `json:"name" api:"required"`
	// JSON path to the value in the transformation output schema (e.g.,
	// ["invoiceDetails", "invoiceNumber"])
	ValueSchemaPath []string `json:"valueSchemaPath,omitzero" api:"required"`
	paramObj
}

func (r ViewGenerateTableDataParamsColumn) MarshalJSON() (data []byte, err error) {
	type shadow ViewGenerateTableDataParamsColumn
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewGenerateTableDataParamsColumn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A filter to apply to a view column
//
// The properties ColumnName, FilterType are required.
type ViewGenerateTableDataParamsFilter struct {
	// Name of the column to filter on
	ColumnName string `json:"columnName" api:"required"`
	// Type of filter to apply to a view column
	//
	// Any of "equals_string", "equals_number", "less_than_number",
	// "less_than_equal_number", "greater_than_number", "greater_than_equal_number",
	// "is_null", "is_not_null".
	FilterType string `json:"filterType,omitzero" api:"required"`
	// Numeric value for the filter (required for number filter types)
	Number param.Opt[float64] `json:"number,omitzero"`
	// String value for the filter (required for string filter types)
	String param.Opt[string] `json:"string,omitzero"`
	paramObj
}

func (r ViewGenerateTableDataParamsFilter) MarshalJSON() (data []byte, err error) {
	type shadow ViewGenerateTableDataParamsFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewGenerateTableDataParamsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ViewGenerateTableDataParamsFilter](
		"filterType", "equals_string", "equals_number", "less_than_number", "less_than_equal_number", "greater_than_number", "greater_than_equal_number", "is_null", "is_not_null",
	)
}

type ViewGenerateTableDataParamsFunction struct {
	// Unique identifier of function. Provide either id or name, not both.
	ID param.Opt[string] `json:"id,omitzero"`
	// Name of function. Must be UNIQUE on a per-environment basis. Provide either id
	// or name, not both.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r ViewGenerateTableDataParamsFunction) MarshalJSON() (data []byte, err error) {
	type shadow ViewGenerateTableDataParamsFunction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewGenerateTableDataParamsFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Time window for filtering transformations in a view
//
// The properties End, Start are required.
type ViewGenerateTableDataParamsTimeWindow struct {
	// End of the time window in ISO 8601 (RFC 3339) format in UTC
	End time.Time `json:"end" api:"required" format:"date-time"`
	// Start of the time window in ISO 8601 (RFC 3339) format in UTC
	Start time.Time `json:"start" api:"required" format:"date-time"`
	paramObj
}

func (r ViewGenerateTableDataParamsTimeWindow) MarshalJSON() (data []byte, err error) {
	type shadow ViewGenerateTableDataParamsTimeWindow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ViewGenerateTableDataParamsTimeWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
