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

	"github.com/stainless-sdks/bem-go/internal/apijson"
	"github.com/stainless-sdks/bem-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/bem-go/internal/encoding/json"
	"github.com/stainless-sdks/bem-go/internal/requestconfig"
	"github.com/stainless-sdks/bem-go/option"
	"github.com/stainless-sdks/bem-go/packages/pagination"
	"github.com/stainless-sdks/bem-go/packages/param"
	"github.com/stainless-sdks/bem-go/packages/respjson"
	"github.com/stainless-sdks/bem-go/shared/constant"
)

// Functions are the core building blocks of data transformation in Bem. Each
// function type serves a specific purpose:
//
//   - **Transform**: Extract structured JSON data from unstructured documents (PDFs,
//     emails, images)
//   - **Analyze**: Perform visual analysis on documents to extract layout-aware
//     information
//   - **Route**: Direct data to different processing paths based on conditions
//   - **Split**: Break multi-page documents into individual pages for parallel
//     processing
//   - **Join**: Combine outputs from multiple function calls into a single result
//   - **Payload Shaping**: Transform and restructure data using JMESPath expressions
//   - **Enrich**: Enhance data with semantic search against collections
//
// Use these endpoints to create, update, list, and manage your functions.
//
// FunctionService contains methods and other services that help with interacting
// with the bem API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFunctionService] method instead.
type FunctionService struct {
	options []option.RequestOption
	// Functions are the core building blocks of data transformation in Bem. Each
	// function type serves a specific purpose:
	//
	//   - **Transform**: Extract structured JSON data from unstructured documents (PDFs,
	//     emails, images)
	//   - **Analyze**: Perform visual analysis on documents to extract layout-aware
	//     information
	//   - **Route**: Direct data to different processing paths based on conditions
	//   - **Split**: Break multi-page documents into individual pages for parallel
	//     processing
	//   - **Join**: Combine outputs from multiple function calls into a single result
	//   - **Payload Shaping**: Transform and restructure data using JMESPath expressions
	//   - **Enrich**: Enhance data with semantic search against collections
	//
	// Use these endpoints to create, update, list, and manage your functions.
	Copy FunctionCopyService
	// Functions are the core building blocks of data transformation in Bem. Each
	// function type serves a specific purpose:
	//
	//   - **Transform**: Extract structured JSON data from unstructured documents (PDFs,
	//     emails, images)
	//   - **Analyze**: Perform visual analysis on documents to extract layout-aware
	//     information
	//   - **Route**: Direct data to different processing paths based on conditions
	//   - **Split**: Break multi-page documents into individual pages for parallel
	//     processing
	//   - **Join**: Combine outputs from multiple function calls into a single result
	//   - **Payload Shaping**: Transform and restructure data using JMESPath expressions
	//   - **Enrich**: Enhance data with semantic search against collections
	//
	// Use these endpoints to create, update, list, and manage your functions.
	Versions FunctionVersionService
}

// NewFunctionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFunctionService(opts ...option.RequestOption) (r FunctionService) {
	r = FunctionService{}
	r.options = opts
	r.Copy = NewFunctionCopyService(opts...)
	r.Versions = NewFunctionVersionService(opts...)
	return
}

// Create a Function
func (r *FunctionService) New(ctx context.Context, body FunctionNewParams, opts ...option.RequestOption) (res *FunctionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/functions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a Function
func (r *FunctionService) Get(ctx context.Context, functionName string, opts ...option.RequestOption) (res *FunctionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if functionName == "" {
		err = errors.New("missing required functionName parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/functions/%s", url.PathEscape(functionName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a Function
func (r *FunctionService) Update(ctx context.Context, pathFunctionName string, body FunctionUpdateParams, opts ...option.RequestOption) (res *FunctionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if pathFunctionName == "" {
		err = errors.New("missing required path_function_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/functions/%s", url.PathEscape(pathFunctionName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List Functions
func (r *FunctionService) List(ctx context.Context, query FunctionListParams, opts ...option.RequestOption) (res *pagination.FunctionsPage[FunctionUnion], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v3/functions"
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

// List Functions
func (r *FunctionService) ListAutoPaging(ctx context.Context, query FunctionListParams, opts ...option.RequestOption) *pagination.FunctionsPageAutoPager[FunctionUnion] {
	return pagination.NewFunctionsPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a Function
func (r *FunctionService) Delete(ctx context.Context, functionName string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if functionName == "" {
		err = errors.New("missing required functionName parameter")
		return err
	}
	path := fmt.Sprintf("v3/functions/%s", url.PathEscape(functionName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

func CreateFunctionParamOfTransform(functionName string) CreateFunctionUnionParam {
	var transform CreateFunctionTransformParam
	transform.FunctionName = functionName
	return CreateFunctionUnionParam{OfTransform: &transform}
}

func CreateFunctionParamOfAnalyze(functionName string) CreateFunctionUnionParam {
	var analyze CreateFunctionAnalyzeParam
	analyze.FunctionName = functionName
	return CreateFunctionUnionParam{OfAnalyze: &analyze}
}

func CreateFunctionParamOfRoute(functionName string) CreateFunctionUnionParam {
	var route CreateFunctionRouteParam
	route.FunctionName = functionName
	return CreateFunctionUnionParam{OfRoute: &route}
}

func CreateFunctionParamOfSplit(functionName string) CreateFunctionUnionParam {
	var split CreateFunctionSplitParam
	split.FunctionName = functionName
	return CreateFunctionUnionParam{OfSplit: &split}
}

func CreateFunctionParamOfJoin(functionName string) CreateFunctionUnionParam {
	var join CreateFunctionJoinParam
	join.FunctionName = functionName
	return CreateFunctionUnionParam{OfJoin: &join}
}

func CreateFunctionParamOfPayloadShaping(functionName string) CreateFunctionUnionParam {
	var payloadShaping CreateFunctionPayloadShapingParam
	payloadShaping.FunctionName = functionName
	return CreateFunctionUnionParam{OfPayloadShaping: &payloadShaping}
}

func CreateFunctionParamOfEnrich(functionName string) CreateFunctionUnionParam {
	var enrich CreateFunctionEnrichParam
	enrich.FunctionName = functionName
	return CreateFunctionUnionParam{OfEnrich: &enrich}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CreateFunctionUnionParam struct {
	OfTransform      *CreateFunctionTransformParam      `json:",omitzero,inline"`
	OfAnalyze        *CreateFunctionAnalyzeParam        `json:",omitzero,inline"`
	OfRoute          *CreateFunctionRouteParam          `json:",omitzero,inline"`
	OfSplit          *CreateFunctionSplitParam          `json:",omitzero,inline"`
	OfJoin           *CreateFunctionJoinParam           `json:",omitzero,inline"`
	OfPayloadShaping *CreateFunctionPayloadShapingParam `json:",omitzero,inline"`
	OfEnrich         *CreateFunctionEnrichParam         `json:",omitzero,inline"`
	paramUnion
}

func (u CreateFunctionUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfTransform,
		u.OfAnalyze,
		u.OfRoute,
		u.OfSplit,
		u.OfJoin,
		u.OfPayloadShaping,
		u.OfEnrich)
}
func (u *CreateFunctionUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[CreateFunctionUnionParam](
		"type",
		apijson.Discriminator[CreateFunctionTransformParam]("transform"),
		apijson.Discriminator[CreateFunctionAnalyzeParam]("analyze"),
		apijson.Discriminator[CreateFunctionRouteParam]("route"),
		apijson.Discriminator[CreateFunctionSplitParam]("split"),
		apijson.Discriminator[CreateFunctionJoinParam]("join"),
		apijson.Discriminator[CreateFunctionPayloadShapingParam]("payload_shaping"),
		apijson.Discriminator[CreateFunctionEnrichParam]("enrich"),
	)
}

// The properties FunctionName, Type are required.
type CreateFunctionTransformParam struct {
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string `json:"functionName" api:"required"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Name of output schema object.
	OutputSchemaName param.Opt[string] `json:"outputSchemaName,omitzero"`
	// Whether tabular chunking is enabled on the pipeline. This processes tables in
	// CSV/Excel in row batches, rather than all rows at once.
	TabularChunkingEnabled param.Opt[bool] `json:"tabularChunkingEnabled,omitzero"`
	// Desired output structure defined in standard JSON Schema convention.
	OutputSchema any `json:"outputSchema,omitzero"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags,omitzero"`
	// This field can be elided, and will marshal its zero value as "transform".
	Type constant.Transform `json:"type" default:"transform"`
	paramObj
}

func (r CreateFunctionTransformParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateFunctionTransformParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateFunctionTransformParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FunctionName, Type are required.
type CreateFunctionAnalyzeParam struct {
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string `json:"functionName" api:"required"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Name of output schema object.
	OutputSchemaName param.Opt[string] `json:"outputSchemaName,omitzero"`
	// Desired output structure defined in standard JSON Schema convention.
	OutputSchema any `json:"outputSchema,omitzero"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags,omitzero"`
	// This field can be elided, and will marshal its zero value as "analyze".
	Type constant.Analyze `json:"type" default:"analyze"`
	paramObj
}

func (r CreateFunctionAnalyzeParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateFunctionAnalyzeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateFunctionAnalyzeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FunctionName, Type are required.
type CreateFunctionRouteParam struct {
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string `json:"functionName" api:"required"`
	// Description of router. Can be used to provide additional context on router's
	// purpose and expected inputs.
	Description param.Opt[string] `json:"description,omitzero"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// List of routes.
	Routes []RouteListItemParam `json:"routes,omitzero"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags,omitzero"`
	// This field can be elided, and will marshal its zero value as "route".
	Type constant.Route `json:"type" default:"route"`
	paramObj
}

func (r CreateFunctionRouteParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateFunctionRouteParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateFunctionRouteParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FunctionName, Type are required.
type CreateFunctionSplitParam struct {
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string `json:"functionName" api:"required"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName             param.Opt[string]                               `json:"displayName,omitzero"`
	PrintPageSplitConfig    CreateFunctionSplitPrintPageSplitConfigParam    `json:"printPageSplitConfig,omitzero"`
	SemanticPageSplitConfig CreateFunctionSplitSemanticPageSplitConfigParam `json:"semanticPageSplitConfig,omitzero"`
	// Any of "print_page", "semantic_page".
	SplitType string `json:"splitType,omitzero"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags,omitzero"`
	// This field can be elided, and will marshal its zero value as "split".
	Type constant.Split `json:"type" default:"split"`
	paramObj
}

func (r CreateFunctionSplitParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateFunctionSplitParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateFunctionSplitParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CreateFunctionSplitParam](
		"splitType", "print_page", "semantic_page",
	)
}

type CreateFunctionSplitPrintPageSplitConfigParam struct {
	NextFunctionID   param.Opt[string] `json:"nextFunctionID,omitzero"`
	NextFunctionName param.Opt[string] `json:"nextFunctionName,omitzero"`
	paramObj
}

func (r CreateFunctionSplitPrintPageSplitConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateFunctionSplitPrintPageSplitConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateFunctionSplitPrintPageSplitConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreateFunctionSplitSemanticPageSplitConfigParam struct {
	ItemClasses []SplitFunctionSemanticPageItemClassParam `json:"itemClasses,omitzero"`
	paramObj
}

func (r CreateFunctionSplitSemanticPageSplitConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateFunctionSplitSemanticPageSplitConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateFunctionSplitSemanticPageSplitConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FunctionName, Type are required.
type CreateFunctionJoinParam struct {
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string `json:"functionName" api:"required"`
	// Description of join function.
	Description param.Opt[string] `json:"description,omitzero"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Name of output schema object.
	OutputSchemaName param.Opt[string] `json:"outputSchemaName,omitzero"`
	// The type of join to perform.
	//
	// Any of "standard".
	JoinType string `json:"joinType,omitzero"`
	// Desired output structure defined in standard JSON Schema convention.
	OutputSchema any `json:"outputSchema,omitzero"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags,omitzero"`
	// This field can be elided, and will marshal its zero value as "join".
	Type constant.Join `json:"type" default:"join"`
	paramObj
}

func (r CreateFunctionJoinParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateFunctionJoinParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateFunctionJoinParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CreateFunctionJoinParam](
		"joinType", "standard",
	)
}

// The properties FunctionName, Type are required.
type CreateFunctionPayloadShapingParam struct {
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string `json:"functionName" api:"required"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// JMESPath expression that defines how to transform and customize the input
	// payload structure. Payload shaping allows you to extract, reshape, and
	// reorganize data from complex input payloads into a simplified, standardized
	// output format. Use JMESPath syntax to select specific fields, perform
	// calculations, and create new data structures tailored to your needs.
	ShapingSchema param.Opt[string] `json:"shapingSchema,omitzero"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags,omitzero"`
	// This field can be elided, and will marshal its zero value as "payload_shaping".
	Type constant.PayloadShaping `json:"type" default:"payload_shaping"`
	paramObj
}

func (r CreateFunctionPayloadShapingParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateFunctionPayloadShapingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateFunctionPayloadShapingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FunctionName, Type are required.
type CreateFunctionEnrichParam struct {
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string `json:"functionName" api:"required"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Configuration for enrich function with semantic search steps.
	//
	// **How Enrich Functions Work:**
	//
	// Enrich functions use semantic search to augment JSON data with relevant
	// information from collections. They take JSON input (typically from a transform
	// function), extract specified fields, perform vector-based semantic search
	// against collections, and inject the results back into the data.
	//
	// **Input Requirements:**
	//
	// - Must receive JSON input (typically uploaded to S3 from a previous function)
	// - Can be chained after transform or other functions that produce JSON output
	//
	// **Example Use Cases:**
	//
	// - Match product descriptions to SKU codes from a product catalog
	// - Enrich customer data with account information
	// - Link order line items to inventory records
	//
	// **Configuration:**
	//
	// - Define one or more enrichment steps
	// - Each step extracts values, searches a collection, and injects results
	// - Steps are executed sequentially
	Config EnrichConfigParam `json:"config,omitzero"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags,omitzero"`
	// This field can be elided, and will marshal its zero value as "enrich".
	Type constant.Enrich `json:"type" default:"enrich"`
	paramObj
}

func (r CreateFunctionEnrichParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateFunctionEnrichParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateFunctionEnrichParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for enrich function with semantic search steps.
//
// **How Enrich Functions Work:**
//
// Enrich functions use semantic search to augment JSON data with relevant
// information from collections. They take JSON input (typically from a transform
// function), extract specified fields, perform vector-based semantic search
// against collections, and inject the results back into the data.
//
// **Input Requirements:**
//
// - Must receive JSON input (typically uploaded to S3 from a previous function)
// - Can be chained after transform or other functions that produce JSON output
//
// **Example Use Cases:**
//
// - Match product descriptions to SKU codes from a product catalog
// - Enrich customer data with account information
// - Link order line items to inventory records
//
// **Configuration:**
//
// - Define one or more enrichment steps
// - Each step extracts values, searches a collection, and injects results
// - Steps are executed sequentially
type EnrichConfig struct {
	// Array of enrichment steps to execute sequentially
	Steps []EnrichStep `json:"steps" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Steps       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnrichConfig) RawJSON() string { return r.JSON.raw }
func (r *EnrichConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EnrichConfig to a EnrichConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EnrichConfigParam.Overrides()
func (r EnrichConfig) ToParam() EnrichConfigParam {
	return param.Override[EnrichConfigParam](json.RawMessage(r.RawJSON()))
}

// Configuration for enrich function with semantic search steps.
//
// **How Enrich Functions Work:**
//
// Enrich functions use semantic search to augment JSON data with relevant
// information from collections. They take JSON input (typically from a transform
// function), extract specified fields, perform vector-based semantic search
// against collections, and inject the results back into the data.
//
// **Input Requirements:**
//
// - Must receive JSON input (typically uploaded to S3 from a previous function)
// - Can be chained after transform or other functions that produce JSON output
//
// **Example Use Cases:**
//
// - Match product descriptions to SKU codes from a product catalog
// - Enrich customer data with account information
// - Link order line items to inventory records
//
// **Configuration:**
//
// - Define one or more enrichment steps
// - Each step extracts values, searches a collection, and injects results
// - Steps are executed sequentially
//
// The property Steps is required.
type EnrichConfigParam struct {
	// Array of enrichment steps to execute sequentially
	Steps []EnrichStepParam `json:"steps,omitzero" api:"required"`
	paramObj
}

func (r EnrichConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow EnrichConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnrichConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single enrichment step configuration.
//
// **Process Flow:**
//
//  1. Extract values from `sourceField` using JMESPath
//  2. Perform search against the specified collection (semantic, exact, or hybrid
//     based on `searchMode`)
//  3. Return top K matches sorted by relevance (best match first)
//  4. Inject results into `targetField`
//
// **Search Modes:**
//
//   - `semantic` (default): Vector similarity search - best for natural language and
//     conceptual matching
//   - `exact`: Exact keyword matching - best for SKU numbers, IDs, routing numbers
//   - `hybrid`: Combined semantic + keyword search - best for tags and categories
//
// **Result Format:**
//
//   - Results are always returned as an array (list), regardless of `topK` value
//   - Array is sorted by relevance (best match first)
//   - Each result contains `data` (the collection item) and optionally
//     `cosineDistance`
//   - With `topK=1`: Returns array with single best match:
//     `[{data: {...}, cosineDistance: 0.15}]`
//   - With `topK>1`: Returns array with multiple matches sorted by relevance
type EnrichStep struct {
	// Name of the collection to search against. The collection must exist and contain
	// items. Supports hierarchical paths when used with `includeSubcollections`.
	CollectionName string `json:"collectionName" api:"required"`
	// JMESPath expression to extract source data for semantic search. Can extract
	// single values or arrays. All extracted values will be used for search.
	SourceField string `json:"sourceField" api:"required"`
	// Field path where enriched results should be placed. Use simple field names
	// (e.g., "enriched_products"). Results are always injected as an array (list),
	// regardless of topK value.
	TargetField string `json:"targetField" api:"required"`
	// Whether to include cosine distance scores in results. Cosine distance ranges
	// from 0.0 (perfect match) to 2.0 (completely dissimilar). Lower scores indicate
	// better semantic similarity.
	//
	// When enabled, each result includes a `cosineDistance` field.
	IncludeCosineDistance bool `json:"includeCosineDistance"`
	// When true, searches all collections under the hierarchical path. For example,
	// "customers" will match "customers", "customers.premium", etc.
	IncludeSubcollections bool `json:"includeSubcollections"`
	// Maximum cosine distance threshold for filtering results (default: 0.6). Results
	// with cosine distance above this threshold are excluded.
	//
	// **Only applies to `semantic` and `hybrid` search modes.** Exact search does not
	// use cosine distance and ignores this setting.
	//
	// Cosine distance ranges from 0.0 (identical) to 2.0 (opposite):
	//
	// - 0.0 - 0.3: Very similar (strict threshold, high-quality matches only)
	// - 0.3 - 0.6: Reasonably similar (moderate threshold)
	// - 0.6 - 1.0: Loosely related (lenient threshold)
	// - > 1.0: Rarely useful — allows nearly unrelated results
	//
	// For most semantic search use cases, good matches typically fall in the 0.2 - 0.5
	// range.
	ScoreThreshold float64 `json:"scoreThreshold"`
	// Search mode to use for enrichment (default: "semantic").
	//
	// **semantic**: Vector similarity search using dense embeddings. Best for finding
	// conceptually similar items.
	//
	// - Use for: Product descriptions, natural language content
	// - Example: "red sports car" matches "crimson convertible automobile"
	//
	// **exact**: Exact keyword matching using PostgreSQL text search. Best for exact
	// identifiers.
	//
	// - Use for: SKU numbers, routing numbers, account IDs, exact tags
	// - Example: "SKU-12345" only matches items containing that exact text
	//
	// **hybrid**: Combined search using 20% semantic + 80% sparse embeddings
	// (keyword-based).
	//
	// - Use for: Tags, categories, partial identifiers
	// - Example: Balances semantic meaning with exact keyword matching
	//
	// Any of "semantic", "exact", "hybrid".
	SearchMode EnrichStepSearchMode `json:"searchMode"`
	// Number of top matching results to return per query (default: 1). Results are
	// always returned as an array (list) and automatically sorted by cosine distance
	// (best match = lowest distance first).
	//
	// - 1: Returns array with single best match: `[{...}]`
	// - > 1: Returns array with multiple matches: `[{...}, {...}, ...]`
	TopK int64 `json:"topK"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CollectionName        respjson.Field
		SourceField           respjson.Field
		TargetField           respjson.Field
		IncludeCosineDistance respjson.Field
		IncludeSubcollections respjson.Field
		ScoreThreshold        respjson.Field
		SearchMode            respjson.Field
		TopK                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnrichStep) RawJSON() string { return r.JSON.raw }
func (r *EnrichStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EnrichStep to a EnrichStepParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EnrichStepParam.Overrides()
func (r EnrichStep) ToParam() EnrichStepParam {
	return param.Override[EnrichStepParam](json.RawMessage(r.RawJSON()))
}

// Search mode to use for enrichment (default: "semantic").
//
// **semantic**: Vector similarity search using dense embeddings. Best for finding
// conceptually similar items.
//
// - Use for: Product descriptions, natural language content
// - Example: "red sports car" matches "crimson convertible automobile"
//
// **exact**: Exact keyword matching using PostgreSQL text search. Best for exact
// identifiers.
//
// - Use for: SKU numbers, routing numbers, account IDs, exact tags
// - Example: "SKU-12345" only matches items containing that exact text
//
// **hybrid**: Combined search using 20% semantic + 80% sparse embeddings
// (keyword-based).
//
// - Use for: Tags, categories, partial identifiers
// - Example: Balances semantic meaning with exact keyword matching
type EnrichStepSearchMode string

const (
	EnrichStepSearchModeSemantic EnrichStepSearchMode = "semantic"
	EnrichStepSearchModeExact    EnrichStepSearchMode = "exact"
	EnrichStepSearchModeHybrid   EnrichStepSearchMode = "hybrid"
)

// Single enrichment step configuration.
//
// **Process Flow:**
//
//  1. Extract values from `sourceField` using JMESPath
//  2. Perform search against the specified collection (semantic, exact, or hybrid
//     based on `searchMode`)
//  3. Return top K matches sorted by relevance (best match first)
//  4. Inject results into `targetField`
//
// **Search Modes:**
//
//   - `semantic` (default): Vector similarity search - best for natural language and
//     conceptual matching
//   - `exact`: Exact keyword matching - best for SKU numbers, IDs, routing numbers
//   - `hybrid`: Combined semantic + keyword search - best for tags and categories
//
// **Result Format:**
//
//   - Results are always returned as an array (list), regardless of `topK` value
//   - Array is sorted by relevance (best match first)
//   - Each result contains `data` (the collection item) and optionally
//     `cosineDistance`
//   - With `topK=1`: Returns array with single best match:
//     `[{data: {...}, cosineDistance: 0.15}]`
//   - With `topK>1`: Returns array with multiple matches sorted by relevance
//
// The properties CollectionName, SourceField, TargetField are required.
type EnrichStepParam struct {
	// Name of the collection to search against. The collection must exist and contain
	// items. Supports hierarchical paths when used with `includeSubcollections`.
	CollectionName string `json:"collectionName" api:"required"`
	// JMESPath expression to extract source data for semantic search. Can extract
	// single values or arrays. All extracted values will be used for search.
	SourceField string `json:"sourceField" api:"required"`
	// Field path where enriched results should be placed. Use simple field names
	// (e.g., "enriched_products"). Results are always injected as an array (list),
	// regardless of topK value.
	TargetField string `json:"targetField" api:"required"`
	// Whether to include cosine distance scores in results. Cosine distance ranges
	// from 0.0 (perfect match) to 2.0 (completely dissimilar). Lower scores indicate
	// better semantic similarity.
	//
	// When enabled, each result includes a `cosineDistance` field.
	IncludeCosineDistance param.Opt[bool] `json:"includeCosineDistance,omitzero"`
	// When true, searches all collections under the hierarchical path. For example,
	// "customers" will match "customers", "customers.premium", etc.
	IncludeSubcollections param.Opt[bool] `json:"includeSubcollections,omitzero"`
	// Maximum cosine distance threshold for filtering results (default: 0.6). Results
	// with cosine distance above this threshold are excluded.
	//
	// **Only applies to `semantic` and `hybrid` search modes.** Exact search does not
	// use cosine distance and ignores this setting.
	//
	// Cosine distance ranges from 0.0 (identical) to 2.0 (opposite):
	//
	// - 0.0 - 0.3: Very similar (strict threshold, high-quality matches only)
	// - 0.3 - 0.6: Reasonably similar (moderate threshold)
	// - 0.6 - 1.0: Loosely related (lenient threshold)
	// - > 1.0: Rarely useful — allows nearly unrelated results
	//
	// For most semantic search use cases, good matches typically fall in the 0.2 - 0.5
	// range.
	ScoreThreshold param.Opt[float64] `json:"scoreThreshold,omitzero"`
	// Number of top matching results to return per query (default: 1). Results are
	// always returned as an array (list) and automatically sorted by cosine distance
	// (best match = lowest distance first).
	//
	// - 1: Returns array with single best match: `[{...}]`
	// - > 1: Returns array with multiple matches: `[{...}, {...}, ...]`
	TopK param.Opt[int64] `json:"topK,omitzero"`
	// Search mode to use for enrichment (default: "semantic").
	//
	// **semantic**: Vector similarity search using dense embeddings. Best for finding
	// conceptually similar items.
	//
	// - Use for: Product descriptions, natural language content
	// - Example: "red sports car" matches "crimson convertible automobile"
	//
	// **exact**: Exact keyword matching using PostgreSQL text search. Best for exact
	// identifiers.
	//
	// - Use for: SKU numbers, routing numbers, account IDs, exact tags
	// - Example: "SKU-12345" only matches items containing that exact text
	//
	// **hybrid**: Combined search using 20% semantic + 80% sparse embeddings
	// (keyword-based).
	//
	// - Use for: Tags, categories, partial identifiers
	// - Example: Balances semantic meaning with exact keyword matching
	//
	// Any of "semantic", "exact", "hybrid".
	SearchMode EnrichStepSearchMode `json:"searchMode,omitzero"`
	paramObj
}

func (r EnrichStepParam) MarshalJSON() (data []byte, err error) {
	type shadow EnrichStepParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnrichStepParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunctionUnion contains all possible properties and values from
// [FunctionTransform], [FunctionAnalyze], [FunctionRoute], [FunctionSplit],
// [FunctionJoin], [FunctionPayloadShaping], [FunctionEnrich].
//
// Use the [FunctionUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FunctionUnion struct {
	EmailAddress     string `json:"emailAddress"`
	FunctionID       string `json:"functionID"`
	FunctionName     string `json:"functionName"`
	OutputSchema     any    `json:"outputSchema"`
	OutputSchemaName string `json:"outputSchemaName"`
	// This field is from variant [FunctionTransform].
	TabularChunkingEnabled bool `json:"tabularChunkingEnabled"`
	// Any of "transform", "analyze", "route", "split", "join", "payload_shaping",
	// "enrich".
	Type       string `json:"type"`
	VersionNum int64  `json:"versionNum"`
	// This field is from variant [FunctionTransform].
	Audit           FunctionAudit       `json:"audit"`
	DisplayName     string              `json:"displayName"`
	Tags            []string            `json:"tags"`
	UsedInWorkflows []WorkflowUsageInfo `json:"usedInWorkflows"`
	Description     string              `json:"description"`
	// This field is from variant [FunctionRoute].
	Routes []RouteListItem `json:"routes"`
	// This field is from variant [FunctionSplit].
	SplitType string `json:"splitType"`
	// This field is from variant [FunctionSplit].
	PrintPageSplitConfig FunctionSplitPrintPageSplitConfig `json:"printPageSplitConfig"`
	// This field is from variant [FunctionSplit].
	SemanticPageSplitConfig FunctionSplitSemanticPageSplitConfig `json:"semanticPageSplitConfig"`
	// This field is from variant [FunctionJoin].
	JoinType string `json:"joinType"`
	// This field is from variant [FunctionPayloadShaping].
	ShapingSchema string `json:"shapingSchema"`
	// This field is from variant [FunctionEnrich].
	Config EnrichConfig `json:"config"`
	JSON   struct {
		EmailAddress            respjson.Field
		FunctionID              respjson.Field
		FunctionName            respjson.Field
		OutputSchema            respjson.Field
		OutputSchemaName        respjson.Field
		TabularChunkingEnabled  respjson.Field
		Type                    respjson.Field
		VersionNum              respjson.Field
		Audit                   respjson.Field
		DisplayName             respjson.Field
		Tags                    respjson.Field
		UsedInWorkflows         respjson.Field
		Description             respjson.Field
		Routes                  respjson.Field
		SplitType               respjson.Field
		PrintPageSplitConfig    respjson.Field
		SemanticPageSplitConfig respjson.Field
		JoinType                respjson.Field
		ShapingSchema           respjson.Field
		Config                  respjson.Field
		raw                     string
	} `json:"-"`
}

// anyFunction is implemented by each variant of [FunctionUnion] to add type safety
// for the return type of [FunctionUnion.AsAny]
type anyFunction interface {
	implFunctionUnion()
}

func (FunctionTransform) implFunctionUnion()      {}
func (FunctionAnalyze) implFunctionUnion()        {}
func (FunctionRoute) implFunctionUnion()          {}
func (FunctionSplit) implFunctionUnion()          {}
func (FunctionJoin) implFunctionUnion()           {}
func (FunctionPayloadShaping) implFunctionUnion() {}
func (FunctionEnrich) implFunctionUnion()         {}

// Use the following switch statement to find the correct variant
//
//	switch variant := FunctionUnion.AsAny().(type) {
//	case bem.FunctionTransform:
//	case bem.FunctionAnalyze:
//	case bem.FunctionRoute:
//	case bem.FunctionSplit:
//	case bem.FunctionJoin:
//	case bem.FunctionPayloadShaping:
//	case bem.FunctionEnrich:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u FunctionUnion) AsAny() anyFunction {
	switch u.Type {
	case "transform":
		return u.AsTransform()
	case "analyze":
		return u.AsAnalyze()
	case "route":
		return u.AsRoute()
	case "split":
		return u.AsSplit()
	case "join":
		return u.AsJoin()
	case "payload_shaping":
		return u.AsPayloadShaping()
	case "enrich":
		return u.AsEnrich()
	}
	return nil
}

func (u FunctionUnion) AsTransform() (v FunctionTransform) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionUnion) AsAnalyze() (v FunctionAnalyze) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionUnion) AsRoute() (v FunctionRoute) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionUnion) AsSplit() (v FunctionSplit) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionUnion) AsJoin() (v FunctionJoin) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionUnion) AsPayloadShaping() (v FunctionPayloadShaping) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionUnion) AsEnrich() (v FunctionEnrich) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunctionUnion) RawJSON() string { return u.JSON.raw }

func (r *FunctionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionTransform struct {
	// Email address automatically created by bem. You can forward emails with or
	// without attachments, to be transformed.
	EmailAddress string `json:"emailAddress" api:"required"`
	// Unique identifier of function.
	FunctionID string `json:"functionID" api:"required"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string `json:"functionName" api:"required"`
	// Desired output structure defined in standard JSON Schema convention.
	OutputSchema any `json:"outputSchema" api:"required"`
	// Name of output schema object.
	OutputSchemaName string `json:"outputSchemaName" api:"required"`
	// Whether tabular chunking is enabled on the pipeline. This processes tables in
	// CSV/Excel in row batches, rather than all rows at once.
	TabularChunkingEnabled bool               `json:"tabularChunkingEnabled" api:"required"`
	Type                   constant.Transform `json:"type" default:"transform"`
	// Version number of function.
	VersionNum int64 `json:"versionNum" api:"required"`
	// Audit trail information for the function.
	Audit FunctionAudit `json:"audit"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName string `json:"displayName"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags"`
	// List of workflows that use this function.
	UsedInWorkflows []WorkflowUsageInfo `json:"usedInWorkflows"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EmailAddress           respjson.Field
		FunctionID             respjson.Field
		FunctionName           respjson.Field
		OutputSchema           respjson.Field
		OutputSchemaName       respjson.Field
		TabularChunkingEnabled respjson.Field
		Type                   respjson.Field
		VersionNum             respjson.Field
		Audit                  respjson.Field
		DisplayName            respjson.Field
		Tags                   respjson.Field
		UsedInWorkflows        respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionTransform) RawJSON() string { return r.JSON.raw }
func (r *FunctionTransform) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionAnalyze struct {
	// Unique identifier of function.
	FunctionID string `json:"functionID" api:"required"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string `json:"functionName" api:"required"`
	// Desired output structure defined in standard JSON Schema convention.
	OutputSchema any `json:"outputSchema" api:"required"`
	// Name of output schema object.
	OutputSchemaName string           `json:"outputSchemaName" api:"required"`
	Type             constant.Analyze `json:"type" default:"analyze"`
	// Version number of function.
	VersionNum int64 `json:"versionNum" api:"required"`
	// Audit trail information for the function.
	Audit FunctionAudit `json:"audit"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName string `json:"displayName"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags"`
	// List of workflows that use this function.
	UsedInWorkflows []WorkflowUsageInfo `json:"usedInWorkflows"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FunctionID       respjson.Field
		FunctionName     respjson.Field
		OutputSchema     respjson.Field
		OutputSchemaName respjson.Field
		Type             respjson.Field
		VersionNum       respjson.Field
		Audit            respjson.Field
		DisplayName      respjson.Field
		Tags             respjson.Field
		UsedInWorkflows  respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionAnalyze) RawJSON() string { return r.JSON.raw }
func (r *FunctionAnalyze) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionRoute struct {
	// Description of router. Can be used to provide additional context on router's
	// purpose and expected inputs.
	Description string `json:"description" api:"required"`
	// Email address automatically created by bem. You can forward emails with or
	// without attachments, to be routed.
	EmailAddress string `json:"emailAddress" api:"required"`
	// Unique identifier of function.
	FunctionID string `json:"functionID" api:"required"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string `json:"functionName" api:"required"`
	// List of routes.
	Routes []RouteListItem `json:"routes" api:"required"`
	Type   constant.Route  `json:"type" default:"route"`
	// Version number of function.
	VersionNum int64 `json:"versionNum" api:"required"`
	// Audit trail information for the function.
	Audit FunctionAudit `json:"audit"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName string `json:"displayName"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags"`
	// List of workflows that use this function.
	UsedInWorkflows []WorkflowUsageInfo `json:"usedInWorkflows"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description     respjson.Field
		EmailAddress    respjson.Field
		FunctionID      respjson.Field
		FunctionName    respjson.Field
		Routes          respjson.Field
		Type            respjson.Field
		VersionNum      respjson.Field
		Audit           respjson.Field
		DisplayName     respjson.Field
		Tags            respjson.Field
		UsedInWorkflows respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionRoute) RawJSON() string { return r.JSON.raw }
func (r *FunctionRoute) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionSplit struct {
	// Unique identifier of function.
	FunctionID string `json:"functionID" api:"required"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string `json:"functionName" api:"required"`
	// The method used to split pages.
	//
	// Any of "print_page", "semantic_page".
	SplitType string         `json:"splitType" api:"required"`
	Type      constant.Split `json:"type" default:"split"`
	// Version number of function.
	VersionNum int64 `json:"versionNum" api:"required"`
	// Audit trail information for the function.
	Audit FunctionAudit `json:"audit"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName string `json:"displayName"`
	// Configuration for print page splitting.
	PrintPageSplitConfig FunctionSplitPrintPageSplitConfig `json:"printPageSplitConfig"`
	// Configuration for semantic page splitting.
	SemanticPageSplitConfig FunctionSplitSemanticPageSplitConfig `json:"semanticPageSplitConfig"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags"`
	// List of workflows that use this function.
	UsedInWorkflows []WorkflowUsageInfo `json:"usedInWorkflows"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FunctionID              respjson.Field
		FunctionName            respjson.Field
		SplitType               respjson.Field
		Type                    respjson.Field
		VersionNum              respjson.Field
		Audit                   respjson.Field
		DisplayName             respjson.Field
		PrintPageSplitConfig    respjson.Field
		SemanticPageSplitConfig respjson.Field
		Tags                    respjson.Field
		UsedInWorkflows         respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionSplit) RawJSON() string { return r.JSON.raw }
func (r *FunctionSplit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for print page splitting.
type FunctionSplitPrintPageSplitConfig struct {
	NextFunctionID string `json:"nextFunctionID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextFunctionID respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionSplitPrintPageSplitConfig) RawJSON() string { return r.JSON.raw }
func (r *FunctionSplitPrintPageSplitConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for semantic page splitting.
type FunctionSplitSemanticPageSplitConfig struct {
	ItemClasses []SplitFunctionSemanticPageItemClass `json:"itemClasses"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemClasses respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionSplitSemanticPageSplitConfig) RawJSON() string { return r.JSON.raw }
func (r *FunctionSplitSemanticPageSplitConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionJoin struct {
	// Description of join function.
	Description string `json:"description" api:"required"`
	// Unique identifier of function.
	FunctionID string `json:"functionID" api:"required"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string `json:"functionName" api:"required"`
	// The type of join to perform.
	//
	// Any of "standard".
	JoinType string `json:"joinType" api:"required"`
	// Desired output structure defined in standard JSON Schema convention.
	OutputSchema any `json:"outputSchema" api:"required"`
	// Name of output schema object.
	OutputSchemaName string        `json:"outputSchemaName" api:"required"`
	Type             constant.Join `json:"type" default:"join"`
	// Version number of function.
	VersionNum int64 `json:"versionNum" api:"required"`
	// Audit trail information for the function.
	Audit FunctionAudit `json:"audit"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName string `json:"displayName"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags"`
	// List of workflows that use this function.
	UsedInWorkflows []WorkflowUsageInfo `json:"usedInWorkflows"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description      respjson.Field
		FunctionID       respjson.Field
		FunctionName     respjson.Field
		JoinType         respjson.Field
		OutputSchema     respjson.Field
		OutputSchemaName respjson.Field
		Type             respjson.Field
		VersionNum       respjson.Field
		Audit            respjson.Field
		DisplayName      respjson.Field
		Tags             respjson.Field
		UsedInWorkflows  respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionJoin) RawJSON() string { return r.JSON.raw }
func (r *FunctionJoin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A function that transforms and customizes input payloads using JMESPath
// expressions. Payload shaping allows you to extract specific data, perform
// calculations, and reshape complex input structures into simplified, standardized
// output formats tailored to your downstream systems or business requirements.
type FunctionPayloadShaping struct {
	// Unique identifier of function.
	FunctionID string `json:"functionID" api:"required"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string `json:"functionName" api:"required"`
	// JMESPath expression that defines how to transform and customize the input
	// payload structure. Payload shaping allows you to extract, reshape, and
	// reorganize data from complex input payloads into a simplified, standardized
	// output format. Use JMESPath syntax to select specific fields, perform
	// calculations, and create new data structures tailored to your needs.
	ShapingSchema string                  `json:"shapingSchema" api:"required"`
	Type          constant.PayloadShaping `json:"type" default:"payload_shaping"`
	// Version number of function.
	VersionNum int64 `json:"versionNum" api:"required"`
	// Audit trail information for the function.
	Audit FunctionAudit `json:"audit"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName string `json:"displayName"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags"`
	// List of workflows that use this function.
	UsedInWorkflows []WorkflowUsageInfo `json:"usedInWorkflows"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FunctionID      respjson.Field
		FunctionName    respjson.Field
		ShapingSchema   respjson.Field
		Type            respjson.Field
		VersionNum      respjson.Field
		Audit           respjson.Field
		DisplayName     respjson.Field
		Tags            respjson.Field
		UsedInWorkflows respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionPayloadShaping) RawJSON() string { return r.JSON.raw }
func (r *FunctionPayloadShaping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionEnrich struct {
	// Configuration for enrich function with semantic search steps.
	//
	// **How Enrich Functions Work:**
	//
	// Enrich functions use semantic search to augment JSON data with relevant
	// information from collections. They take JSON input (typically from a transform
	// function), extract specified fields, perform vector-based semantic search
	// against collections, and inject the results back into the data.
	//
	// **Input Requirements:**
	//
	// - Must receive JSON input (typically uploaded to S3 from a previous function)
	// - Can be chained after transform or other functions that produce JSON output
	//
	// **Example Use Cases:**
	//
	// - Match product descriptions to SKU codes from a product catalog
	// - Enrich customer data with account information
	// - Link order line items to inventory records
	//
	// **Configuration:**
	//
	// - Define one or more enrichment steps
	// - Each step extracts values, searches a collection, and injects results
	// - Steps are executed sequentially
	Config EnrichConfig `json:"config" api:"required"`
	// Unique identifier of function.
	FunctionID string `json:"functionID" api:"required"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string          `json:"functionName" api:"required"`
	Type         constant.Enrich `json:"type" default:"enrich"`
	// Version number of function.
	VersionNum int64 `json:"versionNum" api:"required"`
	// Audit trail information for the function.
	Audit FunctionAudit `json:"audit"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName string `json:"displayName"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags"`
	// List of workflows that use this function.
	UsedInWorkflows []WorkflowUsageInfo `json:"usedInWorkflows"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Config          respjson.Field
		FunctionID      respjson.Field
		FunctionName    respjson.Field
		Type            respjson.Field
		VersionNum      respjson.Field
		Audit           respjson.Field
		DisplayName     respjson.Field
		Tags            respjson.Field
		UsedInWorkflows respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionEnrich) RawJSON() string { return r.JSON.raw }
func (r *FunctionEnrich) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionAudit struct {
	// Information about who created the function.
	FunctionCreatedBy UserActionSummary `json:"functionCreatedBy"`
	// Information about who last updated the function.
	FunctionLastUpdatedBy UserActionSummary `json:"functionLastUpdatedBy"`
	// Information about who created the current version.
	VersionCreatedBy UserActionSummary `json:"versionCreatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FunctionCreatedBy     respjson.Field
		FunctionLastUpdatedBy respjson.Field
		VersionCreatedBy      respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionAudit) RawJSON() string { return r.JSON.raw }
func (r *FunctionAudit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single-function response wrapper used by V3 function endpoints. V3 wraps
// individual function responses in a `{"function": ...}` envelope for consistency
// with other V3 resource endpoints.
type FunctionResponse struct {
	// A function that transforms and customizes input payloads using JMESPath
	// expressions. Payload shaping allows you to extract specific data, perform
	// calculations, and reshape complex input structures into simplified, standardized
	// output formats tailored to your downstream systems or business requirements.
	Function FunctionUnion `json:"function" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Function    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionResponse) RawJSON() string { return r.JSON.raw }
func (r *FunctionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the function.
type FunctionType string

const (
	FunctionTypeTransform      FunctionType = "transform"
	FunctionTypeRoute          FunctionType = "route"
	FunctionTypeSplit          FunctionType = "split"
	FunctionTypeJoin           FunctionType = "join"
	FunctionTypeAnalyze        FunctionType = "analyze"
	FunctionTypePayloadShaping FunctionType = "payload_shaping"
	FunctionTypeEnrich         FunctionType = "enrich"
)

type ListFunctionsResponse struct {
	Functions []FunctionUnion `json:"functions"`
	// The total number of results available.
	TotalCount int64 `json:"totalCount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Functions   respjson.Field
		TotalCount  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListFunctionsResponse) RawJSON() string { return r.JSON.raw }
func (r *ListFunctionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RouteListItem struct {
	Name            string              `json:"name" api:"required"`
	Description     string              `json:"description"`
	FunctionID      string              `json:"functionID"`
	FunctionName    string              `json:"functionName"`
	IsErrorFallback bool                `json:"isErrorFallback"`
	Origin          RouteListItemOrigin `json:"origin"`
	Regex           RouteListItemRegex  `json:"regex"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name            respjson.Field
		Description     respjson.Field
		FunctionID      respjson.Field
		FunctionName    respjson.Field
		IsErrorFallback respjson.Field
		Origin          respjson.Field
		Regex           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouteListItem) RawJSON() string { return r.JSON.raw }
func (r *RouteListItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RouteListItem to a RouteListItemParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RouteListItemParam.Overrides()
func (r RouteListItem) ToParam() RouteListItemParam {
	return param.Override[RouteListItemParam](json.RawMessage(r.RawJSON()))
}

type RouteListItemOrigin struct {
	Email RouteListItemOriginEmail `json:"email"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouteListItemOrigin) RawJSON() string { return r.JSON.raw }
func (r *RouteListItemOrigin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RouteListItemOriginEmail struct {
	Patterns []string `json:"patterns"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Patterns    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouteListItemOriginEmail) RawJSON() string { return r.JSON.raw }
func (r *RouteListItemOriginEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RouteListItemRegex struct {
	Patterns []string `json:"patterns"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Patterns    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouteListItemRegex) RawJSON() string { return r.JSON.raw }
func (r *RouteListItemRegex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type RouteListItemParam struct {
	Name            string                   `json:"name" api:"required"`
	Description     param.Opt[string]        `json:"description,omitzero"`
	FunctionID      param.Opt[string]        `json:"functionID,omitzero"`
	FunctionName    param.Opt[string]        `json:"functionName,omitzero"`
	IsErrorFallback param.Opt[bool]          `json:"isErrorFallback,omitzero"`
	Origin          RouteListItemOriginParam `json:"origin,omitzero"`
	Regex           RouteListItemRegexParam  `json:"regex,omitzero"`
	paramObj
}

func (r RouteListItemParam) MarshalJSON() (data []byte, err error) {
	type shadow RouteListItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RouteListItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RouteListItemOriginParam struct {
	Email RouteListItemOriginEmailParam `json:"email,omitzero"`
	paramObj
}

func (r RouteListItemOriginParam) MarshalJSON() (data []byte, err error) {
	type shadow RouteListItemOriginParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RouteListItemOriginParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RouteListItemOriginEmailParam struct {
	Patterns []string `json:"patterns,omitzero"`
	paramObj
}

func (r RouteListItemOriginEmailParam) MarshalJSON() (data []byte, err error) {
	type shadow RouteListItemOriginEmailParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RouteListItemOriginEmailParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RouteListItemRegexParam struct {
	Patterns []string `json:"patterns,omitzero"`
	paramObj
}

func (r RouteListItemRegexParam) MarshalJSON() (data []byte, err error) {
	type shadow RouteListItemRegexParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RouteListItemRegexParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SplitFunctionSemanticPageItemClass struct {
	Name        string `json:"name" api:"required"`
	Description string `json:"description"`
	// The unique ID of the function you want to use for this item class.
	NextFunctionID string `json:"nextFunctionID"`
	// The unique name of the function you want to use for this item class.
	NextFunctionName string `json:"nextFunctionName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name             respjson.Field
		Description      respjson.Field
		NextFunctionID   respjson.Field
		NextFunctionName respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SplitFunctionSemanticPageItemClass) RawJSON() string { return r.JSON.raw }
func (r *SplitFunctionSemanticPageItemClass) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SplitFunctionSemanticPageItemClass to a
// SplitFunctionSemanticPageItemClassParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SplitFunctionSemanticPageItemClassParam.Overrides()
func (r SplitFunctionSemanticPageItemClass) ToParam() SplitFunctionSemanticPageItemClassParam {
	return param.Override[SplitFunctionSemanticPageItemClassParam](json.RawMessage(r.RawJSON()))
}

// The property Name is required.
type SplitFunctionSemanticPageItemClassParam struct {
	Name        string            `json:"name" api:"required"`
	Description param.Opt[string] `json:"description,omitzero"`
	// The unique ID of the function you want to use for this item class.
	NextFunctionID param.Opt[string] `json:"nextFunctionID,omitzero"`
	// The unique name of the function you want to use for this item class.
	NextFunctionName param.Opt[string] `json:"nextFunctionName,omitzero"`
	paramObj
}

func (r SplitFunctionSemanticPageItemClassParam) MarshalJSON() (data []byte, err error) {
	type shadow SplitFunctionSemanticPageItemClassParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SplitFunctionSemanticPageItemClassParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type UpdateFunctionUnionParam struct {
	OfTransform      *UpdateFunctionTransformParam      `json:",omitzero,inline"`
	OfAnalyze        *UpdateFunctionAnalyzeParam        `json:",omitzero,inline"`
	OfRoute          *UpdateFunctionRouteParam          `json:",omitzero,inline"`
	OfSplit          *UpdateFunctionSplitParam          `json:",omitzero,inline"`
	OfJoin           *UpdateFunctionJoinParam           `json:",omitzero,inline"`
	OfPayloadShaping *UpdateFunctionPayloadShapingParam `json:",omitzero,inline"`
	OfEnrich         *UpdateFunctionEnrichParam         `json:",omitzero,inline"`
	paramUnion
}

func (u UpdateFunctionUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfTransform,
		u.OfAnalyze,
		u.OfRoute,
		u.OfSplit,
		u.OfJoin,
		u.OfPayloadShaping,
		u.OfEnrich)
}
func (u *UpdateFunctionUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[UpdateFunctionUnionParam](
		"type",
		apijson.Discriminator[UpdateFunctionTransformParam]("transform"),
		apijson.Discriminator[UpdateFunctionAnalyzeParam]("analyze"),
		apijson.Discriminator[UpdateFunctionRouteParam]("route"),
		apijson.Discriminator[UpdateFunctionSplitParam]("split"),
		apijson.Discriminator[UpdateFunctionJoinParam]("join"),
		apijson.Discriminator[UpdateFunctionPayloadShapingParam]("payload_shaping"),
		apijson.Discriminator[UpdateFunctionEnrichParam]("enrich"),
	)
}

// The property Type is required.
type UpdateFunctionTransformParam struct {
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName param.Opt[string] `json:"functionName,omitzero"`
	// Name of output schema object.
	OutputSchemaName param.Opt[string] `json:"outputSchemaName,omitzero"`
	// Whether tabular chunking is enabled on the pipeline. This processes tables in
	// CSV/Excel in row batches, rather than all rows at once.
	TabularChunkingEnabled param.Opt[bool] `json:"tabularChunkingEnabled,omitzero"`
	// Desired output structure defined in standard JSON Schema convention.
	OutputSchema any `json:"outputSchema,omitzero"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags,omitzero"`
	// This field can be elided, and will marshal its zero value as "transform".
	Type constant.Transform `json:"type" default:"transform"`
	paramObj
}

func (r UpdateFunctionTransformParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateFunctionTransformParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateFunctionTransformParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type UpdateFunctionAnalyzeParam struct {
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName param.Opt[string] `json:"functionName,omitzero"`
	// Name of output schema object.
	OutputSchemaName param.Opt[string] `json:"outputSchemaName,omitzero"`
	// Desired output structure defined in standard JSON Schema convention.
	OutputSchema any `json:"outputSchema,omitzero"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags,omitzero"`
	// This field can be elided, and will marshal its zero value as "analyze".
	Type constant.Analyze `json:"type" default:"analyze"`
	paramObj
}

func (r UpdateFunctionAnalyzeParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateFunctionAnalyzeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateFunctionAnalyzeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type UpdateFunctionRouteParam struct {
	// Description of router. Can be used to provide additional context on router's
	// purpose and expected inputs.
	Description param.Opt[string] `json:"description,omitzero"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName param.Opt[string] `json:"functionName,omitzero"`
	// List of routes.
	Routes []RouteListItemParam `json:"routes,omitzero"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags,omitzero"`
	// This field can be elided, and will marshal its zero value as "route".
	Type constant.Route `json:"type" default:"route"`
	paramObj
}

func (r UpdateFunctionRouteParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateFunctionRouteParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateFunctionRouteParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type UpdateFunctionSplitParam struct {
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName            param.Opt[string]                               `json:"functionName,omitzero"`
	PrintPageSplitConfig    UpdateFunctionSplitPrintPageSplitConfigParam    `json:"printPageSplitConfig,omitzero"`
	SemanticPageSplitConfig UpdateFunctionSplitSemanticPageSplitConfigParam `json:"semanticPageSplitConfig,omitzero"`
	// Any of "print_page", "semantic_page".
	SplitType string `json:"splitType,omitzero"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags,omitzero"`
	// This field can be elided, and will marshal its zero value as "split".
	Type constant.Split `json:"type" default:"split"`
	paramObj
}

func (r UpdateFunctionSplitParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateFunctionSplitParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateFunctionSplitParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UpdateFunctionSplitParam](
		"splitType", "print_page", "semantic_page",
	)
}

type UpdateFunctionSplitPrintPageSplitConfigParam struct {
	NextFunctionID   param.Opt[string] `json:"nextFunctionID,omitzero"`
	NextFunctionName param.Opt[string] `json:"nextFunctionName,omitzero"`
	paramObj
}

func (r UpdateFunctionSplitPrintPageSplitConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateFunctionSplitPrintPageSplitConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateFunctionSplitPrintPageSplitConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UpdateFunctionSplitSemanticPageSplitConfigParam struct {
	ItemClasses []SplitFunctionSemanticPageItemClassParam `json:"itemClasses,omitzero"`
	paramObj
}

func (r UpdateFunctionSplitSemanticPageSplitConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateFunctionSplitSemanticPageSplitConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateFunctionSplitSemanticPageSplitConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type UpdateFunctionJoinParam struct {
	// Description of join function.
	Description param.Opt[string] `json:"description,omitzero"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName param.Opt[string] `json:"functionName,omitzero"`
	// Name of output schema object.
	OutputSchemaName param.Opt[string] `json:"outputSchemaName,omitzero"`
	// The type of join to perform.
	//
	// Any of "standard".
	JoinType string `json:"joinType,omitzero"`
	// Desired output structure defined in standard JSON Schema convention.
	OutputSchema any `json:"outputSchema,omitzero"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags,omitzero"`
	// This field can be elided, and will marshal its zero value as "join".
	Type constant.Join `json:"type" default:"join"`
	paramObj
}

func (r UpdateFunctionJoinParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateFunctionJoinParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateFunctionJoinParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UpdateFunctionJoinParam](
		"joinType", "standard",
	)
}

// A function that transforms and customizes input payloads using JMESPath
// expressions. Payload shaping allows you to extract specific data, perform
// calculations, and reshape complex input structures into simplified, standardized
// output formats tailored to your downstream systems or business requirements.
//
// The property Type is required.
type UpdateFunctionPayloadShapingParam struct {
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName param.Opt[string] `json:"functionName,omitzero"`
	// JMESPath expression that defines how to transform and customize the input
	// payload structure. Payload shaping allows you to extract, reshape, and
	// reorganize data from complex input payloads into a simplified, standardized
	// output format. Use JMESPath syntax to select specific fields, perform
	// calculations, and create new data structures tailored to your needs.
	ShapingSchema param.Opt[string] `json:"shapingSchema,omitzero"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags,omitzero"`
	// This field can be elided, and will marshal its zero value as "payload_shaping".
	Type constant.PayloadShaping `json:"type" default:"payload_shaping"`
	paramObj
}

func (r UpdateFunctionPayloadShapingParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateFunctionPayloadShapingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateFunctionPayloadShapingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type UpdateFunctionEnrichParam struct {
	// Configuration for enrich function with semantic search steps.
	//
	// **How Enrich Functions Work:**
	//
	// Enrich functions use semantic search to augment JSON data with relevant
	// information from collections. They take JSON input (typically from a transform
	// function), extract specified fields, perform vector-based semantic search
	// against collections, and inject the results back into the data.
	//
	// **Input Requirements:**
	//
	// - Must receive JSON input (typically uploaded to S3 from a previous function)
	// - Can be chained after transform or other functions that produce JSON output
	//
	// **Example Use Cases:**
	//
	// - Match product descriptions to SKU codes from a product catalog
	// - Enrich customer data with account information
	// - Link order line items to inventory records
	//
	// **Configuration:**
	//
	// - Define one or more enrichment steps
	// - Each step extracts values, searches a collection, and injects results
	// - Steps are executed sequentially
	Config EnrichConfigParam `json:"config,omitzero"`
	// This field can be elided, and will marshal its zero value as "enrich".
	Type constant.Enrich `json:"type" default:"enrich"`
	paramObj
}

func (r UpdateFunctionEnrichParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateFunctionEnrichParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateFunctionEnrichParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserActionSummary struct {
	// The date and time the action was created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Unique identifier of the user action.
	UserActionID string `json:"userActionID" api:"required"`
	// API key name. Present for API key-initiated actions.
	APIKeyName string `json:"apiKeyName"`
	// Email address. Present for email-initiated actions.
	EmailAddress string `json:"emailAddress"`
	// User's email address. Present for user-initiated actions.
	UserEmail string `json:"userEmail"`
	// User's ID. Present for user-initiated actions.
	UserID string `json:"userID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt    respjson.Field
		UserActionID respjson.Field
		APIKeyName   respjson.Field
		EmailAddress respjson.Field
		UserEmail    respjson.Field
		UserID       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserActionSummary) RawJSON() string { return r.JSON.raw }
func (r *UserActionSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowUsageInfo struct {
	// Current version number of workflow, provided for reference - compare to
	// usedInWorkflowVersionNums to see whether the current version of the workflow
	// uses this function version.
	CurrentVersionNum int64 `json:"currentVersionNum" api:"required"`
	// Version numbers of workflows that this function version is used in.
	UsedInWorkflowVersionNums []int64 `json:"usedInWorkflowVersionNums" api:"required"`
	// Unique identifier of workflow.
	WorkflowID string `json:"workflowID" api:"required"`
	// Name of workflow.
	WorkflowName string `json:"workflowName" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentVersionNum         respjson.Field
		UsedInWorkflowVersionNums respjson.Field
		WorkflowID                respjson.Field
		WorkflowName              respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowUsageInfo) RawJSON() string { return r.JSON.raw }
func (r *WorkflowUsageInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionNewParams struct {
	CreateFunction CreateFunctionUnionParam
	paramObj
}

func (r FunctionNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateFunction)
}
func (r *FunctionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionUpdateParams struct {
	// A function that transforms and customizes input payloads using JMESPath
	// expressions. Payload shaping allows you to extract specific data, perform
	// calculations, and reshape complex input structures into simplified, standardized
	// output formats tailored to your downstream systems or business requirements.
	UpdateFunction UpdateFunctionUnionParam
	paramObj
}

func (r FunctionUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateFunction)
}
func (r *FunctionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionListParams struct {
	DisplayName   param.Opt[string] `query:"displayName,omitzero" json:"-"`
	EndingBefore  param.Opt[string] `query:"endingBefore,omitzero" json:"-"`
	Limit         param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	StartingAfter param.Opt[string] `query:"startingAfter,omitzero" json:"-"`
	FunctionIDs   []string          `query:"functionIDs,omitzero" json:"-"`
	FunctionNames []string          `query:"functionNames,omitzero" json:"-"`
	// Any of "asc", "desc".
	SortOrder     FunctionListParamsSortOrder `query:"sortOrder,omitzero" json:"-"`
	Tags          []string                    `query:"tags,omitzero" json:"-"`
	Types         []FunctionType              `query:"types,omitzero" json:"-"`
	WorkflowIDs   []string                    `query:"workflowIDs,omitzero" json:"-"`
	WorkflowNames []string                    `query:"workflowNames,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FunctionListParams]'s query parameters as `url.Values`.
func (r FunctionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FunctionListParamsSortOrder string

const (
	FunctionListParamsSortOrderAsc  FunctionListParamsSortOrder = "asc"
	FunctionListParamsSortOrderDesc FunctionListParamsSortOrder = "desc"
)
