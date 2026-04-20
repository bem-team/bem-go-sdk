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
	"github.com/bem-team/bem-go-sdk/packages/pagination"
	"github.com/bem-team/bem-go-sdk/packages/param"
	"github.com/bem-team/bem-go-sdk/packages/respjson"
	"github.com/bem-team/bem-go-sdk/shared/constant"
)

// Functions are the core building blocks of data transformation in Bem. Each
// function type serves a specific purpose:
//
//   - **Extract**: Extract structured JSON data from unstructured documents (PDFs,
//     emails, images, spreadsheets), with optional layout-aware bounding-box
//     extraction
//   - **Route**: Direct data to different processing paths based on conditions
//   - **Split**: Break multi-page documents into individual pages for parallel
//     processing
//   - **Join**: Combine outputs from multiple function calls into a single result
//   - **Payload Shaping**: Transform and restructure data using JMESPath expressions
//   - **Enrich**: Enhance data with semantic search against collections
//   - **Send**: Deliver workflow outputs to downstream destinations
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
	//   - **Extract**: Extract structured JSON data from unstructured documents (PDFs,
	//     emails, images, spreadsheets), with optional layout-aware bounding-box
	//     extraction
	//   - **Route**: Direct data to different processing paths based on conditions
	//   - **Split**: Break multi-page documents into individual pages for parallel
	//     processing
	//   - **Join**: Combine outputs from multiple function calls into a single result
	//   - **Payload Shaping**: Transform and restructure data using JMESPath expressions
	//   - **Enrich**: Enhance data with semantic search against collections
	//   - **Send**: Deliver workflow outputs to downstream destinations
	//
	// Use these endpoints to create, update, list, and manage your functions.
	Copy FunctionCopyService
	// Functions are the core building blocks of data transformation in Bem. Each
	// function type serves a specific purpose:
	//
	//   - **Extract**: Extract structured JSON data from unstructured documents (PDFs,
	//     emails, images, spreadsheets), with optional layout-aware bounding-box
	//     extraction
	//   - **Route**: Direct data to different processing paths based on conditions
	//   - **Split**: Break multi-page documents into individual pages for parallel
	//     processing
	//   - **Join**: Combine outputs from multiple function calls into a single result
	//   - **Payload Shaping**: Transform and restructure data using JMESPath expressions
	//   - **Enrich**: Enhance data with semantic search against collections
	//   - **Send**: Deliver workflow outputs to downstream destinations
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
func (r *FunctionService) List(ctx context.Context, query FunctionListParams, opts ...option.RequestOption) (res *pagination.FunctionsPage[FunctionListResponseUnion], err error) {
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
func (r *FunctionService) ListAutoPaging(ctx context.Context, query FunctionListParams, opts ...option.RequestOption) *pagination.FunctionsPageAutoPager[FunctionListResponseUnion] {
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
	// V3 read-side union. Same shape as the shared `Function` union but with
	// `classify` in place of `route`. Legacy `transform` and `analyze` functions
	// remain readable via V3.
	Function FunctionResponseFunctionUnion `json:"function" api:"required"`
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

// FunctionResponseFunctionUnion contains all possible properties and values from
// [FunctionResponseFunctionTransform], [FunctionResponseFunctionExtract],
// [FunctionResponseFunctionAnalyze], [FunctionResponseFunctionClassify],
// [FunctionResponseFunctionSend], [FunctionResponseFunctionSplit],
// [FunctionResponseFunctionJoin], [FunctionResponseFunctionPayloadShaping],
// [FunctionResponseFunctionEnrich].
//
// Use the [FunctionResponseFunctionUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FunctionResponseFunctionUnion struct {
	EmailAddress           string `json:"emailAddress"`
	FunctionID             string `json:"functionID"`
	FunctionName           string `json:"functionName"`
	OutputSchema           any    `json:"outputSchema"`
	OutputSchemaName       string `json:"outputSchemaName"`
	TabularChunkingEnabled bool   `json:"tabularChunkingEnabled"`
	// Any of "transform", "extract", "analyze", "classify", "send", "split", "join",
	// "payload_shaping", "enrich".
	Type       string `json:"type"`
	VersionNum int64  `json:"versionNum"`
	// This field is from variant [FunctionResponseFunctionTransform].
	Audit           FunctionAudit       `json:"audit"`
	DisplayName     string              `json:"displayName"`
	Tags            []string            `json:"tags"`
	UsedInWorkflows []WorkflowUsageInfo `json:"usedInWorkflows"`
	// This field is from variant [FunctionResponseFunctionAnalyze].
	EnableBoundingBoxes bool `json:"enableBoundingBoxes"`
	// This field is from variant [FunctionResponseFunctionAnalyze].
	PreCount bool `json:"preCount"`
	// This field is from variant [FunctionResponseFunctionClassify].
	Classifications []FunctionResponseFunctionClassifyClassification `json:"classifications"`
	Description     string                                           `json:"description"`
	// This field is from variant [FunctionResponseFunctionSend].
	DestinationType string `json:"destinationType"`
	// This field is from variant [FunctionResponseFunctionSend].
	GoogleDriveFolderID string `json:"googleDriveFolderId"`
	// This field is from variant [FunctionResponseFunctionSend].
	S3Bucket string `json:"s3Bucket"`
	// This field is from variant [FunctionResponseFunctionSend].
	S3Prefix string `json:"s3Prefix"`
	// This field is from variant [FunctionResponseFunctionSend].
	WebhookSigningEnabled bool `json:"webhookSigningEnabled"`
	// This field is from variant [FunctionResponseFunctionSend].
	WebhookURL string `json:"webhookUrl"`
	// This field is from variant [FunctionResponseFunctionSplit].
	SplitType string `json:"splitType"`
	// This field is from variant [FunctionResponseFunctionSplit].
	PrintPageSplitConfig FunctionResponseFunctionSplitPrintPageSplitConfig `json:"printPageSplitConfig"`
	// This field is from variant [FunctionResponseFunctionSplit].
	SemanticPageSplitConfig FunctionResponseFunctionSplitSemanticPageSplitConfig `json:"semanticPageSplitConfig"`
	// This field is from variant [FunctionResponseFunctionJoin].
	JoinType string `json:"joinType"`
	// This field is from variant [FunctionResponseFunctionPayloadShaping].
	ShapingSchema string `json:"shapingSchema"`
	// This field is from variant [FunctionResponseFunctionEnrich].
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
		EnableBoundingBoxes     respjson.Field
		PreCount                respjson.Field
		Classifications         respjson.Field
		Description             respjson.Field
		DestinationType         respjson.Field
		GoogleDriveFolderID     respjson.Field
		S3Bucket                respjson.Field
		S3Prefix                respjson.Field
		WebhookSigningEnabled   respjson.Field
		WebhookURL              respjson.Field
		SplitType               respjson.Field
		PrintPageSplitConfig    respjson.Field
		SemanticPageSplitConfig respjson.Field
		JoinType                respjson.Field
		ShapingSchema           respjson.Field
		Config                  respjson.Field
		raw                     string
	} `json:"-"`
}

// anyFunctionResponseFunction is implemented by each variant of
// [FunctionResponseFunctionUnion] to add type safety for the return type of
// [FunctionResponseFunctionUnion.AsAny]
type anyFunctionResponseFunction interface {
	implFunctionResponseFunctionUnion()
}

func (FunctionResponseFunctionTransform) implFunctionResponseFunctionUnion()      {}
func (FunctionResponseFunctionExtract) implFunctionResponseFunctionUnion()        {}
func (FunctionResponseFunctionAnalyze) implFunctionResponseFunctionUnion()        {}
func (FunctionResponseFunctionClassify) implFunctionResponseFunctionUnion()       {}
func (FunctionResponseFunctionSend) implFunctionResponseFunctionUnion()           {}
func (FunctionResponseFunctionSplit) implFunctionResponseFunctionUnion()          {}
func (FunctionResponseFunctionJoin) implFunctionResponseFunctionUnion()           {}
func (FunctionResponseFunctionPayloadShaping) implFunctionResponseFunctionUnion() {}
func (FunctionResponseFunctionEnrich) implFunctionResponseFunctionUnion()         {}

// Use the following switch statement to find the correct variant
//
//	switch variant := FunctionResponseFunctionUnion.AsAny().(type) {
//	case bem.FunctionResponseFunctionTransform:
//	case bem.FunctionResponseFunctionExtract:
//	case bem.FunctionResponseFunctionAnalyze:
//	case bem.FunctionResponseFunctionClassify:
//	case bem.FunctionResponseFunctionSend:
//	case bem.FunctionResponseFunctionSplit:
//	case bem.FunctionResponseFunctionJoin:
//	case bem.FunctionResponseFunctionPayloadShaping:
//	case bem.FunctionResponseFunctionEnrich:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u FunctionResponseFunctionUnion) AsAny() anyFunctionResponseFunction {
	switch u.Type {
	case "transform":
		return u.AsTransform()
	case "extract":
		return u.AsExtract()
	case "analyze":
		return u.AsAnalyze()
	case "classify":
		return u.AsClassify()
	case "send":
		return u.AsSend()
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

func (u FunctionResponseFunctionUnion) AsTransform() (v FunctionResponseFunctionTransform) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionResponseFunctionUnion) AsExtract() (v FunctionResponseFunctionExtract) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionResponseFunctionUnion) AsAnalyze() (v FunctionResponseFunctionAnalyze) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionResponseFunctionUnion) AsClassify() (v FunctionResponseFunctionClassify) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionResponseFunctionUnion) AsSend() (v FunctionResponseFunctionSend) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionResponseFunctionUnion) AsSplit() (v FunctionResponseFunctionSplit) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionResponseFunctionUnion) AsJoin() (v FunctionResponseFunctionJoin) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionResponseFunctionUnion) AsPayloadShaping() (v FunctionResponseFunctionPayloadShaping) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionResponseFunctionUnion) AsEnrich() (v FunctionResponseFunctionEnrich) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunctionResponseFunctionUnion) RawJSON() string { return u.JSON.raw }

func (r *FunctionResponseFunctionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionResponseFunctionTransform struct {
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
func (r FunctionResponseFunctionTransform) RawJSON() string { return r.JSON.raw }
func (r *FunctionResponseFunctionTransform) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A function that extracts structured JSON from documents and images. Accepts a
// wide range of input types including PDFs, images, spreadsheets, emails, and
// more.
type FunctionResponseFunctionExtract struct {
	// Unique identifier of function.
	FunctionID string `json:"functionID" api:"required"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string `json:"functionName" api:"required"`
	// Desired output structure defined in standard JSON Schema convention.
	OutputSchema any `json:"outputSchema" api:"required"`
	// Name of output schema object.
	OutputSchemaName string `json:"outputSchemaName" api:"required"`
	// Whether tabular chunking is enabled. When true, tables in CSV/Excel files are
	// processed in row batches rather than all at once.
	TabularChunkingEnabled bool             `json:"tabularChunkingEnabled" api:"required"`
	Type                   constant.Extract `json:"type" default:"extract"`
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
func (r FunctionResponseFunctionExtract) RawJSON() string { return r.JSON.raw }
func (r *FunctionResponseFunctionExtract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionResponseFunctionAnalyze struct {
	// Whether bounding box extraction is enabled. Only applicable to analyze and
	// extract functions. When true, the function returns the document regions (page,
	// coordinates) from which each field was extracted.
	EnableBoundingBoxes bool `json:"enableBoundingBoxes" api:"required"`
	// Unique identifier of function.
	FunctionID string `json:"functionID" api:"required"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string `json:"functionName" api:"required"`
	// Desired output structure defined in standard JSON Schema convention.
	OutputSchema any `json:"outputSchema" api:"required"`
	// Name of output schema object.
	OutputSchemaName string `json:"outputSchemaName" api:"required"`
	// Reducing the risk of the model stopping early on long documents. Trade-off:
	// Increases total latency.
	PreCount bool             `json:"preCount" api:"required"`
	Type     constant.Analyze `json:"type" default:"analyze"`
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
		EnableBoundingBoxes respjson.Field
		FunctionID          respjson.Field
		FunctionName        respjson.Field
		OutputSchema        respjson.Field
		OutputSchemaName    respjson.Field
		PreCount            respjson.Field
		Type                respjson.Field
		VersionNum          respjson.Field
		Audit               respjson.Field
		DisplayName         respjson.Field
		Tags                respjson.Field
		UsedInWorkflows     respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionResponseFunctionAnalyze) RawJSON() string { return r.JSON.raw }
func (r *FunctionResponseFunctionAnalyze) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V3 read-side shape of a Classify (internally Route) function. Mirrors {
type FunctionResponseFunctionClassify struct {
	// V3 create/update variants of the shared function payloads.
	//
	// The V3 Functions API no longer accepts the legacy `transform` or `analyze`
	// function types when creating new functions or updating existing ones — both have
	// been unified under `extract`. Existing functions of those types remain readable
	// and callable via V3, so the V3 read-side unions still include `transform` and
	// `analyze` variants.
	//
	// The V3 API also renames the internal `route` function type to `classify` on the
	// wire, and the associated `routes` field to `classifications` (type
	// `ClassificationList`). Platform-internal storage and processing still use
	// `route` / `routes`; the rename is applied only at the V3 API boundary.V3-facing
	// name for the list of classifications a classify function can produce.
	Classifications []FunctionResponseFunctionClassifyClassification `json:"classifications" api:"required"`
	// Description of classifier. Can be used to provide additional context on
	// classifier's purpose and expected inputs.
	Description string `json:"description" api:"required"`
	// Email address automatically created by bem. You can forward emails with or
	// without attachments, to be classified.
	EmailAddress string `json:"emailAddress" api:"required"`
	// Unique identifier of function.
	FunctionID string `json:"functionID" api:"required"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string            `json:"functionName" api:"required"`
	Type         constant.Classify `json:"type" default:"classify"`
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
		Classifications respjson.Field
		Description     respjson.Field
		EmailAddress    respjson.Field
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
func (r FunctionResponseFunctionClassify) RawJSON() string { return r.JSON.raw }
func (r *FunctionResponseFunctionClassify) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionResponseFunctionClassifyClassification struct {
	Name            string                                               `json:"name" api:"required"`
	Description     string                                               `json:"description"`
	FunctionID      string                                               `json:"functionID"`
	FunctionName    string                                               `json:"functionName"`
	IsErrorFallback bool                                                 `json:"isErrorFallback"`
	Origin          FunctionResponseFunctionClassifyClassificationOrigin `json:"origin"`
	Regex           FunctionResponseFunctionClassifyClassificationRegex  `json:"regex"`
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
func (r FunctionResponseFunctionClassifyClassification) RawJSON() string { return r.JSON.raw }
func (r *FunctionResponseFunctionClassifyClassification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionResponseFunctionClassifyClassificationOrigin struct {
	Email FunctionResponseFunctionClassifyClassificationOriginEmail `json:"email"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionResponseFunctionClassifyClassificationOrigin) RawJSON() string { return r.JSON.raw }
func (r *FunctionResponseFunctionClassifyClassificationOrigin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionResponseFunctionClassifyClassificationOriginEmail struct {
	Patterns []string `json:"patterns"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Patterns    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionResponseFunctionClassifyClassificationOriginEmail) RawJSON() string {
	return r.JSON.raw
}
func (r *FunctionResponseFunctionClassifyClassificationOriginEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionResponseFunctionClassifyClassificationRegex struct {
	Patterns []string `json:"patterns"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Patterns    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionResponseFunctionClassifyClassificationRegex) RawJSON() string { return r.JSON.raw }
func (r *FunctionResponseFunctionClassifyClassificationRegex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A function that delivers workflow outputs to an external destination. Send
// functions receive the output of an upstream workflow node and forward it to a
// webhook, S3 bucket, or Google Drive folder.
type FunctionResponseFunctionSend struct {
	// Destination type for a Send function.
	//
	// Any of "webhook", "s3", "google_drive".
	DestinationType string `json:"destinationType" api:"required"`
	// Unique identifier of function.
	FunctionID string `json:"functionID" api:"required"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string        `json:"functionName" api:"required"`
	Type         constant.Send `json:"type" default:"send"`
	// Version number of function.
	VersionNum int64 `json:"versionNum" api:"required"`
	// Audit trail information for the function.
	Audit FunctionAudit `json:"audit"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName string `json:"displayName"`
	// Google Drive folder ID. Present when destinationType is google_drive. Managed
	// via Paragon OAuth.
	GoogleDriveFolderID string `json:"googleDriveFolderId"`
	// S3 bucket to upload the payload to. Present when destinationType is s3.
	S3Bucket string `json:"s3Bucket"`
	// S3 key prefix (folder path). Optional, present when destinationType is s3.
	S3Prefix string `json:"s3Prefix"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags"`
	// List of workflows that use this function.
	UsedInWorkflows []WorkflowUsageInfo `json:"usedInWorkflows"`
	// Whether webhook payloads are signed with an HMAC-SHA256 `bem-signature` header.
	WebhookSigningEnabled bool `json:"webhookSigningEnabled"`
	// Webhook URL to POST the payload to. Present when destinationType is webhook.
	WebhookURL string `json:"webhookUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DestinationType       respjson.Field
		FunctionID            respjson.Field
		FunctionName          respjson.Field
		Type                  respjson.Field
		VersionNum            respjson.Field
		Audit                 respjson.Field
		DisplayName           respjson.Field
		GoogleDriveFolderID   respjson.Field
		S3Bucket              respjson.Field
		S3Prefix              respjson.Field
		Tags                  respjson.Field
		UsedInWorkflows       respjson.Field
		WebhookSigningEnabled respjson.Field
		WebhookURL            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionResponseFunctionSend) RawJSON() string { return r.JSON.raw }
func (r *FunctionResponseFunctionSend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionResponseFunctionSplit struct {
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
	PrintPageSplitConfig FunctionResponseFunctionSplitPrintPageSplitConfig `json:"printPageSplitConfig"`
	// Configuration for semantic page splitting.
	SemanticPageSplitConfig FunctionResponseFunctionSplitSemanticPageSplitConfig `json:"semanticPageSplitConfig"`
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
func (r FunctionResponseFunctionSplit) RawJSON() string { return r.JSON.raw }
func (r *FunctionResponseFunctionSplit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for print page splitting.
type FunctionResponseFunctionSplitPrintPageSplitConfig struct {
	NextFunctionID string `json:"nextFunctionID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextFunctionID respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionResponseFunctionSplitPrintPageSplitConfig) RawJSON() string { return r.JSON.raw }
func (r *FunctionResponseFunctionSplitPrintPageSplitConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for semantic page splitting.
type FunctionResponseFunctionSplitSemanticPageSplitConfig struct {
	ItemClasses []SplitFunctionSemanticPageItemClass `json:"itemClasses"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemClasses respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionResponseFunctionSplitSemanticPageSplitConfig) RawJSON() string { return r.JSON.raw }
func (r *FunctionResponseFunctionSplitSemanticPageSplitConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionResponseFunctionJoin struct {
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
func (r FunctionResponseFunctionJoin) RawJSON() string { return r.JSON.raw }
func (r *FunctionResponseFunctionJoin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A function that transforms and customizes input payloads using JMESPath
// expressions. Payload shaping allows you to extract specific data, perform
// calculations, and reshape complex input structures into simplified, standardized
// output formats tailored to your downstream systems or business requirements.
type FunctionResponseFunctionPayloadShaping struct {
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
func (r FunctionResponseFunctionPayloadShaping) RawJSON() string { return r.JSON.raw }
func (r *FunctionResponseFunctionPayloadShaping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionResponseFunctionEnrich struct {
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
func (r FunctionResponseFunctionEnrich) RawJSON() string { return r.JSON.raw }
func (r *FunctionResponseFunctionEnrich) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the function.
type FunctionType string

const (
	FunctionTypeTransform      FunctionType = "transform"
	FunctionTypeExtract        FunctionType = "extract"
	FunctionTypeRoute          FunctionType = "route"
	FunctionTypeSend           FunctionType = "send"
	FunctionTypeSplit          FunctionType = "split"
	FunctionTypeJoin           FunctionType = "join"
	FunctionTypeAnalyze        FunctionType = "analyze"
	FunctionTypePayloadShaping FunctionType = "payload_shaping"
	FunctionTypeEnrich         FunctionType = "enrich"
)

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

// FunctionListResponseUnion contains all possible properties and values from
// [FunctionListResponseTransform], [FunctionListResponseExtract],
// [FunctionListResponseAnalyze], [FunctionListResponseClassify],
// [FunctionListResponseSend], [FunctionListResponseSplit],
// [FunctionListResponseJoin], [FunctionListResponsePayloadShaping],
// [FunctionListResponseEnrich].
//
// Use the [FunctionListResponseUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FunctionListResponseUnion struct {
	EmailAddress           string `json:"emailAddress"`
	FunctionID             string `json:"functionID"`
	FunctionName           string `json:"functionName"`
	OutputSchema           any    `json:"outputSchema"`
	OutputSchemaName       string `json:"outputSchemaName"`
	TabularChunkingEnabled bool   `json:"tabularChunkingEnabled"`
	// Any of "transform", "extract", "analyze", "classify", "send", "split", "join",
	// "payload_shaping", "enrich".
	Type       string `json:"type"`
	VersionNum int64  `json:"versionNum"`
	// This field is from variant [FunctionListResponseTransform].
	Audit           FunctionAudit       `json:"audit"`
	DisplayName     string              `json:"displayName"`
	Tags            []string            `json:"tags"`
	UsedInWorkflows []WorkflowUsageInfo `json:"usedInWorkflows"`
	// This field is from variant [FunctionListResponseAnalyze].
	EnableBoundingBoxes bool `json:"enableBoundingBoxes"`
	// This field is from variant [FunctionListResponseAnalyze].
	PreCount bool `json:"preCount"`
	// This field is from variant [FunctionListResponseClassify].
	Classifications []FunctionListResponseClassifyClassification `json:"classifications"`
	Description     string                                       `json:"description"`
	// This field is from variant [FunctionListResponseSend].
	DestinationType string `json:"destinationType"`
	// This field is from variant [FunctionListResponseSend].
	GoogleDriveFolderID string `json:"googleDriveFolderId"`
	// This field is from variant [FunctionListResponseSend].
	S3Bucket string `json:"s3Bucket"`
	// This field is from variant [FunctionListResponseSend].
	S3Prefix string `json:"s3Prefix"`
	// This field is from variant [FunctionListResponseSend].
	WebhookSigningEnabled bool `json:"webhookSigningEnabled"`
	// This field is from variant [FunctionListResponseSend].
	WebhookURL string `json:"webhookUrl"`
	// This field is from variant [FunctionListResponseSplit].
	SplitType string `json:"splitType"`
	// This field is from variant [FunctionListResponseSplit].
	PrintPageSplitConfig FunctionListResponseSplitPrintPageSplitConfig `json:"printPageSplitConfig"`
	// This field is from variant [FunctionListResponseSplit].
	SemanticPageSplitConfig FunctionListResponseSplitSemanticPageSplitConfig `json:"semanticPageSplitConfig"`
	// This field is from variant [FunctionListResponseJoin].
	JoinType string `json:"joinType"`
	// This field is from variant [FunctionListResponsePayloadShaping].
	ShapingSchema string `json:"shapingSchema"`
	// This field is from variant [FunctionListResponseEnrich].
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
		EnableBoundingBoxes     respjson.Field
		PreCount                respjson.Field
		Classifications         respjson.Field
		Description             respjson.Field
		DestinationType         respjson.Field
		GoogleDriveFolderID     respjson.Field
		S3Bucket                respjson.Field
		S3Prefix                respjson.Field
		WebhookSigningEnabled   respjson.Field
		WebhookURL              respjson.Field
		SplitType               respjson.Field
		PrintPageSplitConfig    respjson.Field
		SemanticPageSplitConfig respjson.Field
		JoinType                respjson.Field
		ShapingSchema           respjson.Field
		Config                  respjson.Field
		raw                     string
	} `json:"-"`
}

// anyFunctionListResponse is implemented by each variant of
// [FunctionListResponseUnion] to add type safety for the return type of
// [FunctionListResponseUnion.AsAny]
type anyFunctionListResponse interface {
	implFunctionListResponseUnion()
}

func (FunctionListResponseTransform) implFunctionListResponseUnion()      {}
func (FunctionListResponseExtract) implFunctionListResponseUnion()        {}
func (FunctionListResponseAnalyze) implFunctionListResponseUnion()        {}
func (FunctionListResponseClassify) implFunctionListResponseUnion()       {}
func (FunctionListResponseSend) implFunctionListResponseUnion()           {}
func (FunctionListResponseSplit) implFunctionListResponseUnion()          {}
func (FunctionListResponseJoin) implFunctionListResponseUnion()           {}
func (FunctionListResponsePayloadShaping) implFunctionListResponseUnion() {}
func (FunctionListResponseEnrich) implFunctionListResponseUnion()         {}

// Use the following switch statement to find the correct variant
//
//	switch variant := FunctionListResponseUnion.AsAny().(type) {
//	case bem.FunctionListResponseTransform:
//	case bem.FunctionListResponseExtract:
//	case bem.FunctionListResponseAnalyze:
//	case bem.FunctionListResponseClassify:
//	case bem.FunctionListResponseSend:
//	case bem.FunctionListResponseSplit:
//	case bem.FunctionListResponseJoin:
//	case bem.FunctionListResponsePayloadShaping:
//	case bem.FunctionListResponseEnrich:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u FunctionListResponseUnion) AsAny() anyFunctionListResponse {
	switch u.Type {
	case "transform":
		return u.AsTransform()
	case "extract":
		return u.AsExtract()
	case "analyze":
		return u.AsAnalyze()
	case "classify":
		return u.AsClassify()
	case "send":
		return u.AsSend()
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

func (u FunctionListResponseUnion) AsTransform() (v FunctionListResponseTransform) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionListResponseUnion) AsExtract() (v FunctionListResponseExtract) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionListResponseUnion) AsAnalyze() (v FunctionListResponseAnalyze) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionListResponseUnion) AsClassify() (v FunctionListResponseClassify) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionListResponseUnion) AsSend() (v FunctionListResponseSend) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionListResponseUnion) AsSplit() (v FunctionListResponseSplit) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionListResponseUnion) AsJoin() (v FunctionListResponseJoin) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionListResponseUnion) AsPayloadShaping() (v FunctionListResponsePayloadShaping) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionListResponseUnion) AsEnrich() (v FunctionListResponseEnrich) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunctionListResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *FunctionListResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionListResponseTransform struct {
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
func (r FunctionListResponseTransform) RawJSON() string { return r.JSON.raw }
func (r *FunctionListResponseTransform) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A function that extracts structured JSON from documents and images. Accepts a
// wide range of input types including PDFs, images, spreadsheets, emails, and
// more.
type FunctionListResponseExtract struct {
	// Unique identifier of function.
	FunctionID string `json:"functionID" api:"required"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string `json:"functionName" api:"required"`
	// Desired output structure defined in standard JSON Schema convention.
	OutputSchema any `json:"outputSchema" api:"required"`
	// Name of output schema object.
	OutputSchemaName string `json:"outputSchemaName" api:"required"`
	// Whether tabular chunking is enabled. When true, tables in CSV/Excel files are
	// processed in row batches rather than all at once.
	TabularChunkingEnabled bool             `json:"tabularChunkingEnabled" api:"required"`
	Type                   constant.Extract `json:"type" default:"extract"`
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
func (r FunctionListResponseExtract) RawJSON() string { return r.JSON.raw }
func (r *FunctionListResponseExtract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionListResponseAnalyze struct {
	// Whether bounding box extraction is enabled. Only applicable to analyze and
	// extract functions. When true, the function returns the document regions (page,
	// coordinates) from which each field was extracted.
	EnableBoundingBoxes bool `json:"enableBoundingBoxes" api:"required"`
	// Unique identifier of function.
	FunctionID string `json:"functionID" api:"required"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string `json:"functionName" api:"required"`
	// Desired output structure defined in standard JSON Schema convention.
	OutputSchema any `json:"outputSchema" api:"required"`
	// Name of output schema object.
	OutputSchemaName string `json:"outputSchemaName" api:"required"`
	// Reducing the risk of the model stopping early on long documents. Trade-off:
	// Increases total latency.
	PreCount bool             `json:"preCount" api:"required"`
	Type     constant.Analyze `json:"type" default:"analyze"`
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
		EnableBoundingBoxes respjson.Field
		FunctionID          respjson.Field
		FunctionName        respjson.Field
		OutputSchema        respjson.Field
		OutputSchemaName    respjson.Field
		PreCount            respjson.Field
		Type                respjson.Field
		VersionNum          respjson.Field
		Audit               respjson.Field
		DisplayName         respjson.Field
		Tags                respjson.Field
		UsedInWorkflows     respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionListResponseAnalyze) RawJSON() string { return r.JSON.raw }
func (r *FunctionListResponseAnalyze) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V3 read-side shape of a Classify (internally Route) function. Mirrors {
type FunctionListResponseClassify struct {
	// V3 create/update variants of the shared function payloads.
	//
	// The V3 Functions API no longer accepts the legacy `transform` or `analyze`
	// function types when creating new functions or updating existing ones — both have
	// been unified under `extract`. Existing functions of those types remain readable
	// and callable via V3, so the V3 read-side unions still include `transform` and
	// `analyze` variants.
	//
	// The V3 API also renames the internal `route` function type to `classify` on the
	// wire, and the associated `routes` field to `classifications` (type
	// `ClassificationList`). Platform-internal storage and processing still use
	// `route` / `routes`; the rename is applied only at the V3 API boundary.V3-facing
	// name for the list of classifications a classify function can produce.
	Classifications []FunctionListResponseClassifyClassification `json:"classifications" api:"required"`
	// Description of classifier. Can be used to provide additional context on
	// classifier's purpose and expected inputs.
	Description string `json:"description" api:"required"`
	// Email address automatically created by bem. You can forward emails with or
	// without attachments, to be classified.
	EmailAddress string `json:"emailAddress" api:"required"`
	// Unique identifier of function.
	FunctionID string `json:"functionID" api:"required"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string            `json:"functionName" api:"required"`
	Type         constant.Classify `json:"type" default:"classify"`
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
		Classifications respjson.Field
		Description     respjson.Field
		EmailAddress    respjson.Field
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
func (r FunctionListResponseClassify) RawJSON() string { return r.JSON.raw }
func (r *FunctionListResponseClassify) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionListResponseClassifyClassification struct {
	Name            string                                           `json:"name" api:"required"`
	Description     string                                           `json:"description"`
	FunctionID      string                                           `json:"functionID"`
	FunctionName    string                                           `json:"functionName"`
	IsErrorFallback bool                                             `json:"isErrorFallback"`
	Origin          FunctionListResponseClassifyClassificationOrigin `json:"origin"`
	Regex           FunctionListResponseClassifyClassificationRegex  `json:"regex"`
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
func (r FunctionListResponseClassifyClassification) RawJSON() string { return r.JSON.raw }
func (r *FunctionListResponseClassifyClassification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionListResponseClassifyClassificationOrigin struct {
	Email FunctionListResponseClassifyClassificationOriginEmail `json:"email"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionListResponseClassifyClassificationOrigin) RawJSON() string { return r.JSON.raw }
func (r *FunctionListResponseClassifyClassificationOrigin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionListResponseClassifyClassificationOriginEmail struct {
	Patterns []string `json:"patterns"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Patterns    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionListResponseClassifyClassificationOriginEmail) RawJSON() string { return r.JSON.raw }
func (r *FunctionListResponseClassifyClassificationOriginEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionListResponseClassifyClassificationRegex struct {
	Patterns []string `json:"patterns"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Patterns    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionListResponseClassifyClassificationRegex) RawJSON() string { return r.JSON.raw }
func (r *FunctionListResponseClassifyClassificationRegex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A function that delivers workflow outputs to an external destination. Send
// functions receive the output of an upstream workflow node and forward it to a
// webhook, S3 bucket, or Google Drive folder.
type FunctionListResponseSend struct {
	// Destination type for a Send function.
	//
	// Any of "webhook", "s3", "google_drive".
	DestinationType string `json:"destinationType" api:"required"`
	// Unique identifier of function.
	FunctionID string `json:"functionID" api:"required"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string        `json:"functionName" api:"required"`
	Type         constant.Send `json:"type" default:"send"`
	// Version number of function.
	VersionNum int64 `json:"versionNum" api:"required"`
	// Audit trail information for the function.
	Audit FunctionAudit `json:"audit"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName string `json:"displayName"`
	// Google Drive folder ID. Present when destinationType is google_drive. Managed
	// via Paragon OAuth.
	GoogleDriveFolderID string `json:"googleDriveFolderId"`
	// S3 bucket to upload the payload to. Present when destinationType is s3.
	S3Bucket string `json:"s3Bucket"`
	// S3 key prefix (folder path). Optional, present when destinationType is s3.
	S3Prefix string `json:"s3Prefix"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags"`
	// List of workflows that use this function.
	UsedInWorkflows []WorkflowUsageInfo `json:"usedInWorkflows"`
	// Whether webhook payloads are signed with an HMAC-SHA256 `bem-signature` header.
	WebhookSigningEnabled bool `json:"webhookSigningEnabled"`
	// Webhook URL to POST the payload to. Present when destinationType is webhook.
	WebhookURL string `json:"webhookUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DestinationType       respjson.Field
		FunctionID            respjson.Field
		FunctionName          respjson.Field
		Type                  respjson.Field
		VersionNum            respjson.Field
		Audit                 respjson.Field
		DisplayName           respjson.Field
		GoogleDriveFolderID   respjson.Field
		S3Bucket              respjson.Field
		S3Prefix              respjson.Field
		Tags                  respjson.Field
		UsedInWorkflows       respjson.Field
		WebhookSigningEnabled respjson.Field
		WebhookURL            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionListResponseSend) RawJSON() string { return r.JSON.raw }
func (r *FunctionListResponseSend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionListResponseSplit struct {
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
	PrintPageSplitConfig FunctionListResponseSplitPrintPageSplitConfig `json:"printPageSplitConfig"`
	// Configuration for semantic page splitting.
	SemanticPageSplitConfig FunctionListResponseSplitSemanticPageSplitConfig `json:"semanticPageSplitConfig"`
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
func (r FunctionListResponseSplit) RawJSON() string { return r.JSON.raw }
func (r *FunctionListResponseSplit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for print page splitting.
type FunctionListResponseSplitPrintPageSplitConfig struct {
	NextFunctionID string `json:"nextFunctionID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextFunctionID respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionListResponseSplitPrintPageSplitConfig) RawJSON() string { return r.JSON.raw }
func (r *FunctionListResponseSplitPrintPageSplitConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for semantic page splitting.
type FunctionListResponseSplitSemanticPageSplitConfig struct {
	ItemClasses []SplitFunctionSemanticPageItemClass `json:"itemClasses"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemClasses respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionListResponseSplitSemanticPageSplitConfig) RawJSON() string { return r.JSON.raw }
func (r *FunctionListResponseSplitSemanticPageSplitConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionListResponseJoin struct {
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
func (r FunctionListResponseJoin) RawJSON() string { return r.JSON.raw }
func (r *FunctionListResponseJoin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A function that transforms and customizes input payloads using JMESPath
// expressions. Payload shaping allows you to extract specific data, perform
// calculations, and reshape complex input structures into simplified, standardized
// output formats tailored to your downstream systems or business requirements.
type FunctionListResponsePayloadShaping struct {
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
func (r FunctionListResponsePayloadShaping) RawJSON() string { return r.JSON.raw }
func (r *FunctionListResponsePayloadShaping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionListResponseEnrich struct {
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
func (r FunctionListResponseEnrich) RawJSON() string { return r.JSON.raw }
func (r *FunctionListResponseEnrich) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionNewParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfExtract *FunctionNewParamsBodyExtract `json:",inline"`
	// This field is a request body variant, only one variant field can be set. V3 wire
	// form of the Route (classify) function create payload. Mirrors {
	OfClassify *FunctionNewParamsBodyClassify `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfSend *FunctionNewParamsBodySend `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfSplit *FunctionNewParamsBodySplit `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfJoin *FunctionNewParamsBodyJoin `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfPayloadShaping *FunctionNewParamsBodyPayloadShaping `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfEnrich *FunctionNewParamsBodyEnrich `json:",inline"`

	paramObj
}

func (u FunctionNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtract,
		u.OfClassify,
		u.OfSend,
		u.OfSplit,
		u.OfJoin,
		u.OfPayloadShaping,
		u.OfEnrich)
}
func (r *FunctionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FunctionName, Type are required.
type FunctionNewParamsBodyExtract struct {
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string `json:"functionName" api:"required"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Name of output schema object.
	OutputSchemaName param.Opt[string] `json:"outputSchemaName,omitzero"`
	// Whether tabular chunking is enabled. When true, tables in CSV/Excel files are
	// processed in row batches rather than all at once.
	TabularChunkingEnabled param.Opt[bool] `json:"tabularChunkingEnabled,omitzero"`
	// Desired output structure defined in standard JSON Schema convention.
	OutputSchema any `json:"outputSchema,omitzero"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags,omitzero"`
	// This field can be elided, and will marshal its zero value as "extract".
	Type constant.Extract `json:"type" default:"extract"`
	paramObj
}

func (r FunctionNewParamsBodyExtract) MarshalJSON() (data []byte, err error) {
	type shadow FunctionNewParamsBodyExtract
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionNewParamsBodyExtract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V3 wire form of the Route (classify) function create payload. Mirrors {
//
// The properties FunctionName, Type are required.
type FunctionNewParamsBodyClassify struct {
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string `json:"functionName" api:"required"`
	// Description of classifier. Can be used to provide additional context on
	// classifier's purpose and expected inputs.
	Description param.Opt[string] `json:"description,omitzero"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// V3 create/update variants of the shared function payloads.
	//
	// The V3 Functions API no longer accepts the legacy `transform` or `analyze`
	// function types when creating new functions or updating existing ones — both have
	// been unified under `extract`. Existing functions of those types remain readable
	// and callable via V3, so the V3 read-side unions still include `transform` and
	// `analyze` variants.
	//
	// The V3 API also renames the internal `route` function type to `classify` on the
	// wire, and the associated `routes` field to `classifications` (type
	// `ClassificationList`). Platform-internal storage and processing still use
	// `route` / `routes`; the rename is applied only at the V3 API boundary.V3-facing
	// name for the list of classifications a classify function can produce.
	Classifications []FunctionNewParamsBodyClassifyClassification `json:"classifications,omitzero"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags,omitzero"`
	// This field can be elided, and will marshal its zero value as "classify".
	Type constant.Classify `json:"type" default:"classify"`
	paramObj
}

func (r FunctionNewParamsBodyClassify) MarshalJSON() (data []byte, err error) {
	type shadow FunctionNewParamsBodyClassify
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionNewParamsBodyClassify) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type FunctionNewParamsBodyClassifyClassification struct {
	Name            string                                            `json:"name" api:"required"`
	Description     param.Opt[string]                                 `json:"description,omitzero"`
	FunctionID      param.Opt[string]                                 `json:"functionID,omitzero"`
	FunctionName    param.Opt[string]                                 `json:"functionName,omitzero"`
	IsErrorFallback param.Opt[bool]                                   `json:"isErrorFallback,omitzero"`
	Origin          FunctionNewParamsBodyClassifyClassificationOrigin `json:"origin,omitzero"`
	Regex           FunctionNewParamsBodyClassifyClassificationRegex  `json:"regex,omitzero"`
	paramObj
}

func (r FunctionNewParamsBodyClassifyClassification) MarshalJSON() (data []byte, err error) {
	type shadow FunctionNewParamsBodyClassifyClassification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionNewParamsBodyClassifyClassification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionNewParamsBodyClassifyClassificationOrigin struct {
	Email FunctionNewParamsBodyClassifyClassificationOriginEmail `json:"email,omitzero"`
	paramObj
}

func (r FunctionNewParamsBodyClassifyClassificationOrigin) MarshalJSON() (data []byte, err error) {
	type shadow FunctionNewParamsBodyClassifyClassificationOrigin
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionNewParamsBodyClassifyClassificationOrigin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionNewParamsBodyClassifyClassificationOriginEmail struct {
	Patterns []string `json:"patterns,omitzero"`
	paramObj
}

func (r FunctionNewParamsBodyClassifyClassificationOriginEmail) MarshalJSON() (data []byte, err error) {
	type shadow FunctionNewParamsBodyClassifyClassificationOriginEmail
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionNewParamsBodyClassifyClassificationOriginEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionNewParamsBodyClassifyClassificationRegex struct {
	Patterns []string `json:"patterns,omitzero"`
	paramObj
}

func (r FunctionNewParamsBodyClassifyClassificationRegex) MarshalJSON() (data []byte, err error) {
	type shadow FunctionNewParamsBodyClassifyClassificationRegex
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionNewParamsBodyClassifyClassificationRegex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FunctionName, Type are required.
type FunctionNewParamsBodySend struct {
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string `json:"functionName" api:"required"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Google Drive folder ID. Required when destinationType is google_drive. Managed
	// via Paragon OAuth.
	GoogleDriveFolderID param.Opt[string] `json:"googleDriveFolderId,omitzero"`
	// S3 bucket to upload the payload to. Required when destinationType is s3.
	S3Bucket param.Opt[string] `json:"s3Bucket,omitzero"`
	// Optional S3 key prefix (folder path).
	S3Prefix param.Opt[string] `json:"s3Prefix,omitzero"`
	// Whether to sign webhook deliveries with an HMAC-SHA256 `bem-signature` header.
	// Defaults to `true` when omitted — signing is on by default for new send
	// functions. Set explicitly to `false` to disable.
	WebhookSigningEnabled param.Opt[bool] `json:"webhookSigningEnabled,omitzero"`
	// Webhook URL to POST the payload to. Required when destinationType is webhook.
	WebhookURL param.Opt[string] `json:"webhookUrl,omitzero"`
	// Destination type for a Send function.
	//
	// Any of "webhook", "s3", "google_drive".
	DestinationType string `json:"destinationType,omitzero"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags,omitzero"`
	// This field can be elided, and will marshal its zero value as "send".
	Type constant.Send `json:"type" default:"send"`
	paramObj
}

func (r FunctionNewParamsBodySend) MarshalJSON() (data []byte, err error) {
	type shadow FunctionNewParamsBodySend
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionNewParamsBodySend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunctionNewParamsBodySend](
		"destinationType", "webhook", "s3", "google_drive",
	)
}

// The properties FunctionName, Type are required.
type FunctionNewParamsBodySplit struct {
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string `json:"functionName" api:"required"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName             param.Opt[string]                                 `json:"displayName,omitzero"`
	PrintPageSplitConfig    FunctionNewParamsBodySplitPrintPageSplitConfig    `json:"printPageSplitConfig,omitzero"`
	SemanticPageSplitConfig FunctionNewParamsBodySplitSemanticPageSplitConfig `json:"semanticPageSplitConfig,omitzero"`
	// Any of "print_page", "semantic_page".
	SplitType string `json:"splitType,omitzero"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags,omitzero"`
	// This field can be elided, and will marshal its zero value as "split".
	Type constant.Split `json:"type" default:"split"`
	paramObj
}

func (r FunctionNewParamsBodySplit) MarshalJSON() (data []byte, err error) {
	type shadow FunctionNewParamsBodySplit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionNewParamsBodySplit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunctionNewParamsBodySplit](
		"splitType", "print_page", "semantic_page",
	)
}

type FunctionNewParamsBodySplitPrintPageSplitConfig struct {
	NextFunctionID   param.Opt[string] `json:"nextFunctionID,omitzero"`
	NextFunctionName param.Opt[string] `json:"nextFunctionName,omitzero"`
	paramObj
}

func (r FunctionNewParamsBodySplitPrintPageSplitConfig) MarshalJSON() (data []byte, err error) {
	type shadow FunctionNewParamsBodySplitPrintPageSplitConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionNewParamsBodySplitPrintPageSplitConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionNewParamsBodySplitSemanticPageSplitConfig struct {
	ItemClasses []SplitFunctionSemanticPageItemClassParam `json:"itemClasses,omitzero"`
	paramObj
}

func (r FunctionNewParamsBodySplitSemanticPageSplitConfig) MarshalJSON() (data []byte, err error) {
	type shadow FunctionNewParamsBodySplitSemanticPageSplitConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionNewParamsBodySplitSemanticPageSplitConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FunctionName, Type are required.
type FunctionNewParamsBodyJoin struct {
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

func (r FunctionNewParamsBodyJoin) MarshalJSON() (data []byte, err error) {
	type shadow FunctionNewParamsBodyJoin
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionNewParamsBodyJoin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunctionNewParamsBodyJoin](
		"joinType", "standard",
	)
}

// The properties FunctionName, Type are required.
type FunctionNewParamsBodyPayloadShaping struct {
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

func (r FunctionNewParamsBodyPayloadShaping) MarshalJSON() (data []byte, err error) {
	type shadow FunctionNewParamsBodyPayloadShaping
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionNewParamsBodyPayloadShaping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FunctionName, Type are required.
type FunctionNewParamsBodyEnrich struct {
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

func (r FunctionNewParamsBodyEnrich) MarshalJSON() (data []byte, err error) {
	type shadow FunctionNewParamsBodyEnrich
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionNewParamsBodyEnrich) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionUpdateParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfExtract *FunctionUpdateParamsBodyExtract `json:",inline"`
	// This field is a request body variant, only one variant field can be set. V3 wire
	// form of the Route (classify) function upsert payload. Mirrors {
	OfClassify *FunctionUpdateParamsBodyClassify `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfSend *FunctionUpdateParamsBodySend `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfSplit *FunctionUpdateParamsBodySplit `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfJoin *FunctionUpdateParamsBodyJoin `json:",inline"`
	// This field is a request body variant, only one variant field can be set. A
	// function that transforms and customizes input payloads using JMESPath
	// expressions. Payload shaping allows you to extract specific data, perform
	// calculations, and reshape complex input structures into simplified, standardized
	// output formats tailored to your downstream systems or business requirements.
	OfPayloadShaping *FunctionUpdateParamsBodyPayloadShaping `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfEnrich *FunctionUpdateParamsBodyEnrich `json:",inline"`

	paramObj
}

func (u FunctionUpdateParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtract,
		u.OfClassify,
		u.OfSend,
		u.OfSplit,
		u.OfJoin,
		u.OfPayloadShaping,
		u.OfEnrich)
}
func (r *FunctionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type FunctionUpdateParamsBodyExtract struct {
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName param.Opt[string] `json:"functionName,omitzero"`
	// Name of output schema object.
	OutputSchemaName param.Opt[string] `json:"outputSchemaName,omitzero"`
	// Whether tabular chunking is enabled. When true, tables in CSV/Excel files are
	// processed in row batches rather than all at once.
	TabularChunkingEnabled param.Opt[bool] `json:"tabularChunkingEnabled,omitzero"`
	// Desired output structure defined in standard JSON Schema convention.
	OutputSchema any `json:"outputSchema,omitzero"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags,omitzero"`
	// This field can be elided, and will marshal its zero value as "extract".
	Type constant.Extract `json:"type" default:"extract"`
	paramObj
}

func (r FunctionUpdateParamsBodyExtract) MarshalJSON() (data []byte, err error) {
	type shadow FunctionUpdateParamsBodyExtract
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionUpdateParamsBodyExtract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V3 wire form of the Route (classify) function upsert payload. Mirrors {
//
// The property Type is required.
type FunctionUpdateParamsBodyClassify struct {
	// Description of classifier. Can be used to provide additional context on
	// classifier's purpose and expected inputs.
	Description param.Opt[string] `json:"description,omitzero"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName param.Opt[string] `json:"functionName,omitzero"`
	// V3 create/update variants of the shared function payloads.
	//
	// The V3 Functions API no longer accepts the legacy `transform` or `analyze`
	// function types when creating new functions or updating existing ones — both have
	// been unified under `extract`. Existing functions of those types remain readable
	// and callable via V3, so the V3 read-side unions still include `transform` and
	// `analyze` variants.
	//
	// The V3 API also renames the internal `route` function type to `classify` on the
	// wire, and the associated `routes` field to `classifications` (type
	// `ClassificationList`). Platform-internal storage and processing still use
	// `route` / `routes`; the rename is applied only at the V3 API boundary.V3-facing
	// name for the list of classifications a classify function can produce.
	Classifications []FunctionUpdateParamsBodyClassifyClassification `json:"classifications,omitzero"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags,omitzero"`
	// This field can be elided, and will marshal its zero value as "classify".
	Type constant.Classify `json:"type" default:"classify"`
	paramObj
}

func (r FunctionUpdateParamsBodyClassify) MarshalJSON() (data []byte, err error) {
	type shadow FunctionUpdateParamsBodyClassify
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionUpdateParamsBodyClassify) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type FunctionUpdateParamsBodyClassifyClassification struct {
	Name            string                                               `json:"name" api:"required"`
	Description     param.Opt[string]                                    `json:"description,omitzero"`
	FunctionID      param.Opt[string]                                    `json:"functionID,omitzero"`
	FunctionName    param.Opt[string]                                    `json:"functionName,omitzero"`
	IsErrorFallback param.Opt[bool]                                      `json:"isErrorFallback,omitzero"`
	Origin          FunctionUpdateParamsBodyClassifyClassificationOrigin `json:"origin,omitzero"`
	Regex           FunctionUpdateParamsBodyClassifyClassificationRegex  `json:"regex,omitzero"`
	paramObj
}

func (r FunctionUpdateParamsBodyClassifyClassification) MarshalJSON() (data []byte, err error) {
	type shadow FunctionUpdateParamsBodyClassifyClassification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionUpdateParamsBodyClassifyClassification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionUpdateParamsBodyClassifyClassificationOrigin struct {
	Email FunctionUpdateParamsBodyClassifyClassificationOriginEmail `json:"email,omitzero"`
	paramObj
}

func (r FunctionUpdateParamsBodyClassifyClassificationOrigin) MarshalJSON() (data []byte, err error) {
	type shadow FunctionUpdateParamsBodyClassifyClassificationOrigin
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionUpdateParamsBodyClassifyClassificationOrigin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionUpdateParamsBodyClassifyClassificationOriginEmail struct {
	Patterns []string `json:"patterns,omitzero"`
	paramObj
}

func (r FunctionUpdateParamsBodyClassifyClassificationOriginEmail) MarshalJSON() (data []byte, err error) {
	type shadow FunctionUpdateParamsBodyClassifyClassificationOriginEmail
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionUpdateParamsBodyClassifyClassificationOriginEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionUpdateParamsBodyClassifyClassificationRegex struct {
	Patterns []string `json:"patterns,omitzero"`
	paramObj
}

func (r FunctionUpdateParamsBodyClassifyClassificationRegex) MarshalJSON() (data []byte, err error) {
	type shadow FunctionUpdateParamsBodyClassifyClassificationRegex
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionUpdateParamsBodyClassifyClassificationRegex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type FunctionUpdateParamsBodySend struct {
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName param.Opt[string] `json:"functionName,omitzero"`
	// Google Drive folder ID. Required when destinationType is google_drive. Managed
	// via Paragon OAuth.
	GoogleDriveFolderID param.Opt[string] `json:"googleDriveFolderId,omitzero"`
	// S3 bucket to upload the payload to. Required when destinationType is s3.
	S3Bucket param.Opt[string] `json:"s3Bucket,omitzero"`
	// Optional S3 key prefix (folder path).
	S3Prefix param.Opt[string] `json:"s3Prefix,omitzero"`
	// Whether to sign webhook deliveries with an HMAC-SHA256 `bem-signature` header.
	// Defaults to `true` when omitted — signing is on by default for new send
	// functions. Set explicitly to `false` to disable.
	WebhookSigningEnabled param.Opt[bool] `json:"webhookSigningEnabled,omitzero"`
	// Webhook URL to POST the payload to. Required when destinationType is webhook.
	WebhookURL param.Opt[string] `json:"webhookUrl,omitzero"`
	// Destination type for a Send function.
	//
	// Any of "webhook", "s3", "google_drive".
	DestinationType string `json:"destinationType,omitzero"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags,omitzero"`
	// This field can be elided, and will marshal its zero value as "send".
	Type constant.Send `json:"type" default:"send"`
	paramObj
}

func (r FunctionUpdateParamsBodySend) MarshalJSON() (data []byte, err error) {
	type shadow FunctionUpdateParamsBodySend
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionUpdateParamsBodySend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunctionUpdateParamsBodySend](
		"destinationType", "webhook", "s3", "google_drive",
	)
}

// The property Type is required.
type FunctionUpdateParamsBodySplit struct {
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName            param.Opt[string]                                    `json:"functionName,omitzero"`
	PrintPageSplitConfig    FunctionUpdateParamsBodySplitPrintPageSplitConfig    `json:"printPageSplitConfig,omitzero"`
	SemanticPageSplitConfig FunctionUpdateParamsBodySplitSemanticPageSplitConfig `json:"semanticPageSplitConfig,omitzero"`
	// Any of "print_page", "semantic_page".
	SplitType string `json:"splitType,omitzero"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags,omitzero"`
	// This field can be elided, and will marshal its zero value as "split".
	Type constant.Split `json:"type" default:"split"`
	paramObj
}

func (r FunctionUpdateParamsBodySplit) MarshalJSON() (data []byte, err error) {
	type shadow FunctionUpdateParamsBodySplit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionUpdateParamsBodySplit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunctionUpdateParamsBodySplit](
		"splitType", "print_page", "semantic_page",
	)
}

type FunctionUpdateParamsBodySplitPrintPageSplitConfig struct {
	NextFunctionID   param.Opt[string] `json:"nextFunctionID,omitzero"`
	NextFunctionName param.Opt[string] `json:"nextFunctionName,omitzero"`
	paramObj
}

func (r FunctionUpdateParamsBodySplitPrintPageSplitConfig) MarshalJSON() (data []byte, err error) {
	type shadow FunctionUpdateParamsBodySplitPrintPageSplitConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionUpdateParamsBodySplitPrintPageSplitConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionUpdateParamsBodySplitSemanticPageSplitConfig struct {
	ItemClasses []SplitFunctionSemanticPageItemClassParam `json:"itemClasses,omitzero"`
	paramObj
}

func (r FunctionUpdateParamsBodySplitSemanticPageSplitConfig) MarshalJSON() (data []byte, err error) {
	type shadow FunctionUpdateParamsBodySplitSemanticPageSplitConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionUpdateParamsBodySplitSemanticPageSplitConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type FunctionUpdateParamsBodyJoin struct {
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

func (r FunctionUpdateParamsBodyJoin) MarshalJSON() (data []byte, err error) {
	type shadow FunctionUpdateParamsBodyJoin
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionUpdateParamsBodyJoin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunctionUpdateParamsBodyJoin](
		"joinType", "standard",
	)
}

// A function that transforms and customizes input payloads using JMESPath
// expressions. Payload shaping allows you to extract specific data, perform
// calculations, and reshape complex input structures into simplified, standardized
// output formats tailored to your downstream systems or business requirements.
//
// The property Type is required.
type FunctionUpdateParamsBodyPayloadShaping struct {
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

func (r FunctionUpdateParamsBodyPayloadShaping) MarshalJSON() (data []byte, err error) {
	type shadow FunctionUpdateParamsBodyPayloadShaping
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionUpdateParamsBodyPayloadShaping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type FunctionUpdateParamsBodyEnrich struct {
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

func (r FunctionUpdateParamsBodyEnrich) MarshalJSON() (data []byte, err error) {
	type shadow FunctionUpdateParamsBodyEnrich
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionUpdateParamsBodyEnrich) UnmarshalJSON(data []byte) error {
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
