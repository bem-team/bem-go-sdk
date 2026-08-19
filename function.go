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
	"github.com/bem-team/bem-go-sdk/shared/constant"
)

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
	//   - **Parse**: Render documents into a navigable structure of page-aware sections,
	//     named entities, and relationships — designed to be walked by an LLM agent via
	//     the [File System API](/api/v3/file-system) (`POST /v3/fs`). Two toggles, both
	//     `true` by default: `extractEntities` controls per-document entity and
	//     relationship extraction; `linkAcrossDocuments` merges entities into one
	//     canonical record per real-world thing across the environment, populating
	//     cross-document memory.
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
	//   - **Parse**: Render documents into a navigable structure of page-aware sections,
	//     named entities, and relationships — designed to be walked by an LLM agent via
	//     the [File System API](/api/v3/file-system) (`POST /v3/fs`). Two toggles, both
	//     `true` by default: `extractEntities` controls per-document entity and
	//     relationship extraction; `linkAcrossDocuments` merges entities into one
	//     canonical record per real-world thing across the environment, populating
	//     cross-document memory.
	//   - **Payload Shaping**: Transform and restructure data using JMESPath expressions
	//   - **Enrich**: Enhance data with semantic search against collections
	//   - **Send**: Deliver workflow outputs to downstream destinations
	//
	// Use these endpoints to create, update, list, and manage your functions.
	Versions FunctionVersionService
	// Monitor, evaluate, and iterate on the quality of every function in your
	// environment. Function Accuracy bundles two complementary loops:
	//
	// ## Evaluations (`/v3/eval`)
	//
	// Trigger and retrieve per-transformation evaluations. Evaluations run
	// asynchronously and score each transformation's output against the function's
	// schema for confidence, per-field hallucination detection, and relevance.
	// Supported for `extract`, `transform`, `analyze`, and `join` events.
	//
	//  1. **Trigger** — `POST /v3/eval` queues jobs for a batch of transformation IDs.
	//  2. **Poll** — `GET /v3/eval/results` returns the current state of each requested
	//     ID, partitioned into `results`, `pending`, and `failed`. Accepts either
	//     `eventIDs` (preferred) or `transformationIDs` as a comma-separated query
	//     parameter, and always keys the response by event KSUID.
	//
	// Up to 100 IDs may be submitted per request.
	//
	// ## Metrics, review, regression (`/v3/functions/{metrics,review,regression,compare}`)
	//
	// Roll evaluation results and user corrections up into actionable function-level
	// signal:
	//
	//   - **`GET /v3/functions/metrics`** — aggregate accuracy, precision, recall, F1,
	//     and confusion-matrix counts per function.
	//   - **`POST /v3/functions/review`** — sample-size estimation, confidence-bucketed
	//     distribution, PR-AUC, and per-threshold confidence intervals (Wald or Wilson)
	//     for picking review cutoffs.
	//   - **`POST /v3/functions/regression`** — replay corrected historical inputs
	//     against a new function version, producing a labeled regression dataset.
	//   - **`POST /v3/functions/regression/corrections`** — propagate baseline
	//     corrections onto the regression dataset so it can be scored.
	//   - **`POST /v3/functions/compare`** — compute aggregate and field-level lift
	//     between any two versions, optionally scoped to the regression dataset.
	//
	// All five endpoints support `extract` end-to-end on both the vision and OCR
	// paths, alongside the legacy `transform` / `analyze` / `join` types.
	Regression FunctionRegressionService
}

// NewFunctionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFunctionService(opts ...option.RequestOption) (r FunctionService) {
	r = FunctionService{}
	r.options = opts
	r.Copy = NewFunctionCopyService(opts...)
	r.Versions = NewFunctionVersionService(opts...)
	r.Regression = NewFunctionRegressionService(opts...)
	return
}

// **Create a function.**
//
// The function `type` determines which configuration fields are required — see the
// `CreateFunctionV3` discriminated union and
// [Function types overview](/guide/function-types/overview) for the per-type
// contract.
//
// The response contains both `functionID` and `functionName`. Either is a stable
// handle you can use elsewhere; most workflows reference functions by
// `functionName` because it's human-readable.
//
// ## Naming rules
//
//   - `functionName` must be unique per environment.
//   - Allowed characters: letters, digits, hyphens, and underscores.
//   - Names cannot be reused after deletion within the same environment for at least
//     the retention window of the previous record.
//
// The new function is created at `versionNum: 1`. Subsequent
// `PATCH /v3/functions/{functionName}` calls produce new versions — the version-1
// configuration remains immutable and addressable.
func (r *FunctionService) New(ctx context.Context, body FunctionNewParams, opts ...option.RequestOption) (res *FunctionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/functions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// **Retrieve a function's current version by name.**
//
// Returns the function record with its `currentVersionNum` and the configuration
// of that version. To inspect a historical version, use
// `GET /v3/functions/{functionName}/versions/{versionNum}`.
func (r *FunctionService) Get(ctx context.Context, functionName string, query FunctionGetParams, opts ...option.RequestOption) (res *FunctionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if functionName == "" {
		err = errors.New("missing required functionName parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/functions/%s", url.PathEscape(functionName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// **Update a function. Updates create a new version.**
//
// The previous version remains addressable and immutable. Workflow nodes that
// pinned the function with a `versionNum` continue to use the pinned version;
// nodes that reference the function by name with no version automatically pick up
// the new version on their next call.
//
// ## What you can change
//
// Any field allowed by the function's type. Most commonly: `outputSchema` (for
// `extract`/`join`), `classifications` (for `classify`), `displayName`, and
// `tags`.
//
// ## Versioning behaviour
//
//   - Each successful update increments `currentVersionNum` by 1.
//   - `displayName`, `tags`, and `functionName` updates also create a new version,
//     so the version history is a complete record of every change.
//   - To revert, fetch the previous version and re-submit its configuration as a new
//     update — versions themselves are immutable.
func (r *FunctionService) Update(ctx context.Context, functionName string, body FunctionUpdateParams, opts ...option.RequestOption) (res *FunctionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if functionName == "" {
		err = errors.New("missing required functionName parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/functions/%s", url.PathEscape(functionName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// **List functions in the current environment.**
//
// Returns each function's current version. Combine filters freely — they AND
// together.
//
// ## Filtering
//
//   - `functionIDs` / `functionNames`: exact-match identity filters.
//   - `displayName`: case-insensitive substring match.
//   - `types`: one or more of `extract`, `classify`, `split`, `join`, `enrich`,
//     `payload_shaping`. Legacy `transform`, `analyze`, `route`, and `send` types
//     remain readable via this filter.
//   - `tags`: returns functions tagged with any of the supplied tags.
//   - `workflowIDs` / `workflowNames`: returns only functions referenced by the
//     named workflows. Useful for "what functions does this workflow depend on?"
//     lookups.
//   - `workflowIDVersionNums` / `workflowNameVersionNums`: the same lookup pinned to
//     a specific workflow version.
//
// ## Pagination
//
// Cursor-based with `startingAfter` and `endingBefore` (functionIDs). Default
// limit 50, maximum 100.
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

// **List functions in the current environment.**
//
// Returns each function's current version. Combine filters freely — they AND
// together.
//
// ## Filtering
//
//   - `functionIDs` / `functionNames`: exact-match identity filters.
//   - `displayName`: case-insensitive substring match.
//   - `types`: one or more of `extract`, `classify`, `split`, `join`, `enrich`,
//     `payload_shaping`. Legacy `transform`, `analyze`, `route`, and `send` types
//     remain readable via this filter.
//   - `tags`: returns functions tagged with any of the supplied tags.
//   - `workflowIDs` / `workflowNames`: returns only functions referenced by the
//     named workflows. Useful for "what functions does this workflow depend on?"
//     lookups.
//   - `workflowIDVersionNums` / `workflowNameVersionNums`: the same lookup pinned to
//     a specific workflow version.
//
// ## Pagination
//
// Cursor-based with `startingAfter` and `endingBefore` (functionIDs). Default
// limit 50, maximum 100.
func (r *FunctionService) ListAutoPaging(ctx context.Context, query FunctionListParams, opts ...option.RequestOption) *pagination.FunctionsPageAutoPager[FunctionUnion] {
	return pagination.NewFunctionsPageAutoPager(r.List(ctx, query, opts...))
}

// **Delete a function and every one of its versions.**
//
// Permanent. Running and queued calls that reference this function continue to
// completion against the version they captured at call time, but no new calls can
// target it.
//
// ## Before deleting
//
// Workflow nodes that reference this function will fail at call time after
// deletion. List workflows that reference it first:
//
// ```
// GET /v3/workflows?functionNames=my-function
// ```
//
// Update or remove those workflows, or create a replacement function and re-point
// the workflow nodes, before deleting.
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

// **Compare metrics between two function versions.**
//
// Computes aggregate and field-level lift/regression between any two versions of a
// function: accuracy, precision, recall, F1, and PR-AUC. Field-level changes are
// returned only for fields whose lift exceeds 1% in either direction.
//
// Supported for every function type that produces labeled transformations:
// `extract`, `transform`, `analyze`, `join`. Pass `isRegression: true` to compare
// only the regression dataset (rows produced by `POST /v3/functions/regression`) —
// the canonical way to judge a candidate version before promoting it.
//
// Defaults: `baselineVersionNum = currentVersionNum - 1`,
// `comparisonVersionNum = currentVersionNum`.
func (r *FunctionService) CompareMetrics(ctx context.Context, body FunctionCompareMetricsParams, opts ...option.RequestOption) (res *FunctionCompareMetricsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/functions/compare"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// **Estimate human review requirements for a function.**
//
// Combines confusion-matrix metrics with the per-transformation evaluation scores
// (confidence / hallucination / relevance produced by the eval service) to
// compute:
//
//   - A confidence-bucketed distribution of the function's outputs.
//   - Sample-size estimates at configurable margin-of-error and confidence levels
//     (Wald or Wilson intervals).
//   - A precision-recall AUC and a per-threshold matrix you can use to pick a review
//     cutoff.
//
// Supported for every function type that produces transformations and feeds the
// auto-evaluation pipeline: `extract`, `transform`, `analyze`, `join`. Extract
// works on both vision (PDF/PNG/JPEG/HEIC/HEIF/WebP) and OCR-routed inputs.
//
// Pass `isRegression: true` to scope the review to transformations created by a
// previous regression run (see `POST /v3/functions/regression`).
func (r *FunctionService) EstimateReviewRequirements(ctx context.Context, body FunctionEstimateReviewRequirementsParams, opts ...option.RequestOption) (res *FunctionEstimateReviewRequirementsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/functions/review"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// **Retrieve performance metrics for functions based on labeled transformation
// data.**
//
// Calculates accuracy, precision, recall, F1, and the underlying confusion-matrix
// counts for each matching function by comparing model outputs against user
// corrections. Metrics are aggregated across every transformation the function has
// produced, regardless of function type — `extract`, `transform`, `analyze`, and
// `join` all populate the same `metrics` column on the transformation row, so v3
// surfaces all of them uniformly.
//
// ## Filtering
//
// Combine `functionIDs` / `functionNames` / `types` to narrow the result set.
// `types` accepts `extract` alongside the legacy `transform` / `analyze` types
// (which remain readable). Pagination is cursor-based.
//
// ## Requirements
//
// A function only shows non-zero metrics once at least one of its transformations
// has been labeled — submit corrections via `POST /v3/events/{eventID}/feedback`.
func (r *FunctionService) GetMetrics(ctx context.Context, query FunctionGetMetricsParams, opts ...option.RequestOption) (res *FunctionGetMetricsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/functions/metrics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ClassificationListItem struct {
	Name            string                       `json:"name" api:"required"`
	Description     string                       `json:"description"`
	FunctionID      string                       `json:"functionID"`
	FunctionName    string                       `json:"functionName"`
	IsErrorFallback bool                         `json:"isErrorFallback"`
	Origin          ClassificationListItemOrigin `json:"origin"`
	Regex           ClassificationListItemRegex  `json:"regex"`
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
func (r ClassificationListItem) RawJSON() string { return r.JSON.raw }
func (r *ClassificationListItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ClassificationListItem to a ClassificationListItemParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ClassificationListItemParam.Overrides()
func (r ClassificationListItem) ToParam() ClassificationListItemParam {
	return param.Override[ClassificationListItemParam](json.RawMessage(r.RawJSON()))
}

type ClassificationListItemOrigin struct {
	Email ClassificationListItemOriginEmail `json:"email"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClassificationListItemOrigin) RawJSON() string { return r.JSON.raw }
func (r *ClassificationListItemOrigin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClassificationListItemOriginEmail struct {
	Patterns []string `json:"patterns"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Patterns    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClassificationListItemOriginEmail) RawJSON() string { return r.JSON.raw }
func (r *ClassificationListItemOriginEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClassificationListItemRegex struct {
	Patterns []string `json:"patterns"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Patterns    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClassificationListItemRegex) RawJSON() string { return r.JSON.raw }
func (r *ClassificationListItemRegex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type ClassificationListItemParam struct {
	Name            string                            `json:"name" api:"required"`
	Description     param.Opt[string]                 `json:"description,omitzero"`
	FunctionID      param.Opt[string]                 `json:"functionID,omitzero"`
	FunctionName    param.Opt[string]                 `json:"functionName,omitzero"`
	IsErrorFallback param.Opt[bool]                   `json:"isErrorFallback,omitzero"`
	Origin          ClassificationListItemOriginParam `json:"origin,omitzero"`
	Regex           ClassificationListItemRegexParam  `json:"regex,omitzero"`
	paramObj
}

func (r ClassificationListItemParam) MarshalJSON() (data []byte, err error) {
	type shadow ClassificationListItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClassificationListItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClassificationListItemOriginParam struct {
	Email ClassificationListItemOriginEmailParam `json:"email,omitzero"`
	paramObj
}

func (r ClassificationListItemOriginParam) MarshalJSON() (data []byte, err error) {
	type shadow ClassificationListItemOriginParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClassificationListItemOriginParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClassificationListItemOriginEmailParam struct {
	Patterns []string `json:"patterns,omitzero"`
	paramObj
}

func (r ClassificationListItemOriginEmailParam) MarshalJSON() (data []byte, err error) {
	type shadow ClassificationListItemOriginEmailParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClassificationListItemOriginEmailParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClassificationListItemRegexParam struct {
	Patterns []string `json:"patterns,omitzero"`
	paramObj
}

func (r ClassificationListItemRegexParam) MarshalJSON() (data []byte, err error) {
	type shadow ClassificationListItemRegexParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClassificationListItemRegexParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func CreateFunctionParamOfExtract(functionName string) CreateFunctionUnionParam {
	var extract CreateFunctionExtractParam
	extract.FunctionName = functionName
	return CreateFunctionUnionParam{OfExtract: &extract}
}

func CreateFunctionParamOfClassify(functionName string) CreateFunctionUnionParam {
	var classify CreateFunctionClassifyParam
	classify.FunctionName = functionName
	return CreateFunctionUnionParam{OfClassify: &classify}
}

func CreateFunctionParamOfSend(functionName string) CreateFunctionUnionParam {
	var send CreateFunctionSendParam
	send.FunctionName = functionName
	return CreateFunctionUnionParam{OfSend: &send}
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

func CreateFunctionParamOfParse(functionName string) CreateFunctionUnionParam {
	var parse CreateFunctionParseParam
	parse.FunctionName = functionName
	return CreateFunctionUnionParam{OfParse: &parse}
}

func CreateFunctionParamOfRender(functionName string, renderConfig RenderConfigInputParam) CreateFunctionUnionParam {
	var render CreateFunctionRenderParam
	render.FunctionName = functionName
	render.RenderConfig = renderConfig
	return CreateFunctionUnionParam{OfRender: &render}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CreateFunctionUnionParam struct {
	OfExtract        *CreateFunctionExtractParam        `json:",omitzero,inline"`
	OfClassify       *CreateFunctionClassifyParam       `json:",omitzero,inline"`
	OfSend           *CreateFunctionSendParam           `json:",omitzero,inline"`
	OfSplit          *CreateFunctionSplitParam          `json:",omitzero,inline"`
	OfJoin           *CreateFunctionJoinParam           `json:",omitzero,inline"`
	OfPayloadShaping *CreateFunctionPayloadShapingParam `json:",omitzero,inline"`
	OfEnrich         *CreateFunctionEnrichParam         `json:",omitzero,inline"`
	OfParse          *CreateFunctionParseParam          `json:",omitzero,inline"`
	OfRender         *CreateFunctionRenderParam         `json:",omitzero,inline"`
	paramUnion
}

func (u CreateFunctionUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtract,
		u.OfClassify,
		u.OfSend,
		u.OfSplit,
		u.OfJoin,
		u.OfPayloadShaping,
		u.OfEnrich,
		u.OfParse,
		u.OfRender)
}
func (u *CreateFunctionUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[CreateFunctionUnionParam](
		"type",
		apijson.Discriminator[CreateFunctionExtractParam]("extract"),
		apijson.Discriminator[CreateFunctionClassifyParam]("classify"),
		apijson.Discriminator[CreateFunctionSendParam]("send"),
		apijson.Discriminator[CreateFunctionSplitParam]("split"),
		apijson.Discriminator[CreateFunctionJoinParam]("join"),
		apijson.Discriminator[CreateFunctionPayloadShapingParam]("payload_shaping"),
		apijson.Discriminator[CreateFunctionEnrichParam]("enrich"),
		apijson.Discriminator[CreateFunctionParseParam]("parse"),
		apijson.Discriminator[CreateFunctionRenderParam]("render"),
	)
}

// The properties FunctionName, Type are required.
type CreateFunctionExtractParam struct {
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string `json:"functionName" api:"required"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Whether bounding box extraction is enabled. Applies to vision input types (pdf,
	// png, jpeg, heic, heif, webp) that dispatch through the analyze path. When true,
	// the function returns the document regions (page, coordinates) from which each
	// field was extracted. Enabling this automatically configures the function to use
	// the bounding box model. Disabling resets to the default.
	EnableBoundingBoxes param.Opt[bool] `json:"enableBoundingBoxes,omitzero"`
	// Name of output schema object.
	OutputSchemaName param.Opt[string] `json:"outputSchemaName,omitzero"`
	// Reducing the risk of the model stopping early on long documents. Trade-off:
	// Increases total latency. Compatible with `enableBoundingBoxes`.
	PreCount param.Opt[bool] `json:"preCount,omitzero"`
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

func (r CreateFunctionExtractParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateFunctionExtractParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateFunctionExtractParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V3 wire form of the classify function create payload.
//
// The properties FunctionName, Type are required.
type CreateFunctionClassifyParam struct {
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string `json:"functionName" api:"required"`
	// Description of classifier. Can be used to provide additional context on
	// classifier's purpose and expected inputs.
	Description param.Opt[string] `json:"description,omitzero"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// When true, image and PDF inputs are sent directly to the model for routing
	// instead of being OCR'd to text first. Defaults to true for new classify
	// functions and false for the legacy route type.
	NativeVisualInput param.Opt[bool] `json:"nativeVisualInput,omitzero"`
	// List of classifications a classify function can produce. Shares the underlying
	// route list shape.
	Classifications []ClassificationListItemParam `json:"classifications,omitzero"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags,omitzero"`
	// This field can be elided, and will marshal its zero value as "classify".
	Type constant.Classify `json:"type" default:"classify"`
	paramObj
}

func (r CreateFunctionClassifyParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateFunctionClassifyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateFunctionClassifyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FunctionName, Type are required.
type CreateFunctionSendParam struct {
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
	// Where the payload is delivered.
	//
	// Any of "webhook", "s3", "google_drive".
	DestinationType SendDestinationType `json:"destinationType,omitzero"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags,omitzero"`
	// This field can be elided, and will marshal its zero value as "send".
	Type constant.Send `json:"type" default:"send"`
	paramObj
}

func (r CreateFunctionSendParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateFunctionSendParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateFunctionSendParam) UnmarshalJSON(data []byte) error {
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
	// Configuration for an enrich function.
	//
	// **How Enrich Functions Work:**
	//
	// Enrich functions augment JSON input with data from external sources. They take
	// JSON input (typically from a previous function), extract specified fields, fetch
	// or search for matching data, and inject the results back into the JSON.
	//
	// **Data Sources:**
	//
	//   - **Collections** (`source: "collection"`): Vector/keyword search against a BEM
	//     collection. Best for semantic matching against pre-indexed documents.
	//   - **Endpoints** (`source: "endpoint"`): HTTP call to any user-provided REST API.
	//     Best for looking up live data from CRMs, ERPs, or other external systems.
	//     Optionally uses LLM agent reasoning to rank candidates returned by the
	//     endpoint.
	//
	// **Input Requirements:**
	//
	// - Must receive JSON input (typically from a previous function's output)
	//
	// **Example Use Cases:**
	//
	//   - Match product descriptions to SKU codes from a product catalog collection
	//   - Enrich customer data with account details from a CRM endpoint
	//   - Use LLM agent reasoning to fuzzy-match line item descriptions to catalog
	//     products
	//
	// **Configuration:**
	//
	// - Define named endpoints (for endpoint-source steps)
	// - Define one or more enrichment steps; steps are executed sequentially
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

// The properties FunctionName, Type are required.
type CreateFunctionParseParam struct {
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string `json:"functionName" api:"required"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Cross-cutting toggles for Parse functions. Mirrors the `extraConfig` surface on
	// Extract / Join — separated from `parseConfig` so the per-call Parse output shape
	// stays distinct from operator-level execution flags.
	ExtraConfig ParseExtraFunctionConfigParam `json:"extraConfig,omitzero"`
	// Per-version configuration for a Parse function.
	//
	// Parse renders document pages (PDF, image) via vision LLM and emits structured
	// JSON. The two toggles below independently control entity extraction (a per-call
	// output concern) and cross-document memory linking (an environment-wide concern).
	ParseConfig ParseConfigParam `json:"parseConfig,omitzero"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags,omitzero"`
	// This field can be elided, and will marshal its zero value as "parse".
	Type constant.Parse `json:"type" default:"parse"`
	paramObj
}

func (r CreateFunctionParseParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateFunctionParseParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateFunctionParseParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FunctionName, RenderConfig, Type are required.
type CreateFunctionRenderParam struct {
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string `json:"functionName" api:"required"`
	// Render configuration. Required at create time — a Render function without a
	// template has nothing to bind data to. Update bodies may omit this for partial
	// edits.
	RenderConfig RenderConfigInputParam `json:"renderConfig,omitzero" api:"required"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags,omitzero"`
	// This field can be elided, and will marshal its zero value as "render".
	Type constant.Render `json:"type" default:"render"`
	paramObj
}

func (r CreateFunctionRenderParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateFunctionRenderParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateFunctionRenderParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for an enrich function.
//
// **How Enrich Functions Work:**
//
// Enrich functions augment JSON input with data from external sources. They take
// JSON input (typically from a previous function), extract specified fields, fetch
// or search for matching data, and inject the results back into the JSON.
//
// **Data Sources:**
//
//   - **Collections** (`source: "collection"`): Vector/keyword search against a BEM
//     collection. Best for semantic matching against pre-indexed documents.
//   - **Endpoints** (`source: "endpoint"`): HTTP call to any user-provided REST API.
//     Best for looking up live data from CRMs, ERPs, or other external systems.
//     Optionally uses LLM agent reasoning to rank candidates returned by the
//     endpoint.
//
// **Input Requirements:**
//
// - Must receive JSON input (typically from a previous function's output)
//
// **Example Use Cases:**
//
//   - Match product descriptions to SKU codes from a product catalog collection
//   - Enrich customer data with account details from a CRM endpoint
//   - Use LLM agent reasoning to fuzzy-match line item descriptions to catalog
//     products
//
// **Configuration:**
//
// - Define named endpoints (for endpoint-source steps)
// - Define one or more enrichment steps; steps are executed sequentially
type EnrichConfig struct {
	// Array of enrichment steps to execute sequentially.
	Steps []EnrichStep `json:"steps" api:"required"`
	// Named HTTP endpoints available to endpoint-source steps. Each endpoint must have
	// a unique `name` referenced by the step's `endpointName`. Required when any step
	// uses `source: "endpoint"`.
	Endpoints []EnrichConfigEndpoint `json:"endpoints"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Steps       respjson.Field
		Endpoints   respjson.Field
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

// A named HTTP endpoint that an enrich step can call to fetch enrichment data.
//
// The platform makes one request per extracted source value, substituting the
// value as a query parameter or body template placeholder. The raw response (or
// the sub-value selected by `responsePath`) is injected into the output, or passed
// to LLM agent reasoning when `matchInstructions` is set.
//
// **Request formats:**
//
//   - `GET`: Appends `?{queryParam}={value}` to the URL.
//   - `POST`: Sends `bodyTemplate` as the request body, replacing `{value}` with the
//     extracted value.
type EnrichConfigEndpoint struct {
	// HTTP method to use.
	//
	// Any of "GET", "POST".
	Method string `json:"method" api:"required"`
	// Unique name for this endpoint, referenced by enrichStep.endpointName.
	Name string `json:"name" api:"required"`
	// Full URL of the endpoint (must be http:// or https://).
	URL string `json:"url" api:"required"`
	// JSON body template for POST requests. **Required for POST endpoints.** Must
	// contain the `{value}` placeholder, which is replaced with the extracted source
	// value at runtime.
	//
	// Example: `bodyTemplate: "{\"query\": \"{value}\", \"limit\": 10}"`
	BodyTemplate string `json:"bodyTemplate"`
	// Additional HTTP headers to include in every request (e.g.
	// `Authorization: Bearer <token>`).
	Headers any `json:"headers"`
	// Natural-language instructions for LLM agent reasoning.
	//
	// When set, the candidates fetched from the endpoint are passed to an LLM with
	// these instructions, which selects the best match(es) and returns them ranked
	// best-first. Each injected result has the shape
	// `{ data, rank, confidence, reasoning? }` (rank is 1-based, 1 = best).
	//
	// When omitted, the raw fetched value is injected without any LLM involvement.
	MatchInstructions string `json:"matchInstructions"`
	// Maximum number of ranked matches to return per source value when
	// `matchInstructions` is set (default: 1). Ignored when `matchInstructions` is
	// empty.
	MatchTopK int64 `json:"matchTopK"`
	// LLM batch size during agent reasoning (default: 50). All candidates — across all
	// fetched pages — are scored in batches of this size. Smaller values reduce
	// per-call token usage; larger values mean fewer LLM calls. Ignored when
	// `matchInstructions` is empty.
	MaxCandidates int64 `json:"maxCandidates"`
	// Maximum number of pages to fetch (default: 10). Acts as a safety cap against
	// infinite pagination loops when the server never returns an empty cursor.
	MaxPages int64 `json:"maxPages"`
	// Query parameter name used to pass the cursor on subsequent GET requests, or the
	// `{placeholder}` name used in the POST `bodyTemplate` (e.g. `"cursor"`,
	// `"pageToken"`, `"offset"`).
	//
	// Must be set together with `nextPagePath`.
	NextPageParam string `json:"nextPageParam"`
	// JMESPath expression applied to each raw response to extract the cursor or token
	// for the next page (e.g. `"nextCursor"`, `"pagination.nextToken"`). An absent,
	// null, or empty-string result stops pagination. Both string and numeric values
	// are supported — numbers are converted to their decimal string representation
	// before being forwarded as a query parameter.
	//
	// Must be set together with `nextPageParam`.
	//
	// **Supported pagination styles:**
	//
	//   - **Cursor/token-based** — server returns an opaque token in the response body
	//     (e.g. `{"nextCursor": "abc123"}`). Set `nextPagePath: "nextCursor"` and the
	//     platform forwards it verbatim on the next request.
	//   - **Server-computed offset/page** — server echoes back the next offset or page
	//     number in the response body (e.g. `{"nextOffset": 50}` or `{"nextPage": 2}`).
	//     Set `nextPagePath: "nextOffset"` and the platform forwards the value as-is.
	//
	// **Not supported:**
	//
	//   - **Client-computed offset** — APIs where the client must compute
	//     `offset += limit` itself (e.g. `?offset=0&limit=50` with no next-offset in the
	//     response). Workaround: ask the API provider to return the next offset in the
	//     response body, or bake a fixed page size into the URL and use a server-side
	//     cursor instead.
	//   - **Client-computed page number** — APIs where the client increments `?page=N`
	//     itself with no next-page value in the response. Same workaround applies.
	//   - **Link header** — `Link: <url>; rel="next"` in HTTP response headers. The
	//     platform only inspects the response body.
	NextPagePath string `json:"nextPagePath"`
	// Query parameter name used to pass the extracted source value. **Required for GET
	// endpoints.** The value is URL-encoded and appended as
	// `?{queryParam}={sourceValue}`.
	//
	// Example: `queryParam: "q"` → `GET /products?q=blue+widget`
	QueryParam string `json:"queryParam"`
	// JMESPath expression applied to the response body to extract the enrichment
	// value. Omit to use the entire response body as the result.
	//
	// **For agent reasoning:** use a wildcard projection (e.g. `items[*]` or
	// `results[*].data`) so the endpoint's list of candidates is flattened into an
	// array before being passed to the LLM. A non-wildcard path (e.g. `data.product`)
	// extracts a single value treated as one candidate.
	//
	// **Response size:** the platform reads at most 50 MB of the response body before
	// decoding, regardless of the Content-Length header.
	ResponsePath string `json:"responsePath"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method            respjson.Field
		Name              respjson.Field
		URL               respjson.Field
		BodyTemplate      respjson.Field
		Headers           respjson.Field
		MatchInstructions respjson.Field
		MatchTopK         respjson.Field
		MaxCandidates     respjson.Field
		MaxPages          respjson.Field
		NextPageParam     respjson.Field
		NextPagePath      respjson.Field
		QueryParam        respjson.Field
		ResponsePath      respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnrichConfigEndpoint) RawJSON() string { return r.JSON.raw }
func (r *EnrichConfigEndpoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for an enrich function.
//
// **How Enrich Functions Work:**
//
// Enrich functions augment JSON input with data from external sources. They take
// JSON input (typically from a previous function), extract specified fields, fetch
// or search for matching data, and inject the results back into the JSON.
//
// **Data Sources:**
//
//   - **Collections** (`source: "collection"`): Vector/keyword search against a BEM
//     collection. Best for semantic matching against pre-indexed documents.
//   - **Endpoints** (`source: "endpoint"`): HTTP call to any user-provided REST API.
//     Best for looking up live data from CRMs, ERPs, or other external systems.
//     Optionally uses LLM agent reasoning to rank candidates returned by the
//     endpoint.
//
// **Input Requirements:**
//
// - Must receive JSON input (typically from a previous function's output)
//
// **Example Use Cases:**
//
//   - Match product descriptions to SKU codes from a product catalog collection
//   - Enrich customer data with account details from a CRM endpoint
//   - Use LLM agent reasoning to fuzzy-match line item descriptions to catalog
//     products
//
// **Configuration:**
//
// - Define named endpoints (for endpoint-source steps)
// - Define one or more enrichment steps; steps are executed sequentially
//
// The property Steps is required.
type EnrichConfigParam struct {
	// Array of enrichment steps to execute sequentially.
	Steps []EnrichStepParam `json:"steps,omitzero" api:"required"`
	// Named HTTP endpoints available to endpoint-source steps. Each endpoint must have
	// a unique `name` referenced by the step's `endpointName`. Required when any step
	// uses `source: "endpoint"`.
	Endpoints []EnrichConfigEndpointParam `json:"endpoints,omitzero"`
	paramObj
}

func (r EnrichConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow EnrichConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnrichConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A named HTTP endpoint that an enrich step can call to fetch enrichment data.
//
// The platform makes one request per extracted source value, substituting the
// value as a query parameter or body template placeholder. The raw response (or
// the sub-value selected by `responsePath`) is injected into the output, or passed
// to LLM agent reasoning when `matchInstructions` is set.
//
// **Request formats:**
//
//   - `GET`: Appends `?{queryParam}={value}` to the URL.
//   - `POST`: Sends `bodyTemplate` as the request body, replacing `{value}` with the
//     extracted value.
//
// The properties Method, Name, URL are required.
type EnrichConfigEndpointParam struct {
	// HTTP method to use.
	//
	// Any of "GET", "POST".
	Method string `json:"method,omitzero" api:"required"`
	// Unique name for this endpoint, referenced by enrichStep.endpointName.
	Name string `json:"name" api:"required"`
	// Full URL of the endpoint (must be http:// or https://).
	URL string `json:"url" api:"required"`
	// JSON body template for POST requests. **Required for POST endpoints.** Must
	// contain the `{value}` placeholder, which is replaced with the extracted source
	// value at runtime.
	//
	// Example: `bodyTemplate: "{\"query\": \"{value}\", \"limit\": 10}"`
	BodyTemplate param.Opt[string] `json:"bodyTemplate,omitzero"`
	// Natural-language instructions for LLM agent reasoning.
	//
	// When set, the candidates fetched from the endpoint are passed to an LLM with
	// these instructions, which selects the best match(es) and returns them ranked
	// best-first. Each injected result has the shape
	// `{ data, rank, confidence, reasoning? }` (rank is 1-based, 1 = best).
	//
	// When omitted, the raw fetched value is injected without any LLM involvement.
	MatchInstructions param.Opt[string] `json:"matchInstructions,omitzero"`
	// Maximum number of ranked matches to return per source value when
	// `matchInstructions` is set (default: 1). Ignored when `matchInstructions` is
	// empty.
	MatchTopK param.Opt[int64] `json:"matchTopK,omitzero"`
	// LLM batch size during agent reasoning (default: 50). All candidates — across all
	// fetched pages — are scored in batches of this size. Smaller values reduce
	// per-call token usage; larger values mean fewer LLM calls. Ignored when
	// `matchInstructions` is empty.
	MaxCandidates param.Opt[int64] `json:"maxCandidates,omitzero"`
	// Maximum number of pages to fetch (default: 10). Acts as a safety cap against
	// infinite pagination loops when the server never returns an empty cursor.
	MaxPages param.Opt[int64] `json:"maxPages,omitzero"`
	// Query parameter name used to pass the cursor on subsequent GET requests, or the
	// `{placeholder}` name used in the POST `bodyTemplate` (e.g. `"cursor"`,
	// `"pageToken"`, `"offset"`).
	//
	// Must be set together with `nextPagePath`.
	NextPageParam param.Opt[string] `json:"nextPageParam,omitzero"`
	// JMESPath expression applied to each raw response to extract the cursor or token
	// for the next page (e.g. `"nextCursor"`, `"pagination.nextToken"`). An absent,
	// null, or empty-string result stops pagination. Both string and numeric values
	// are supported — numbers are converted to their decimal string representation
	// before being forwarded as a query parameter.
	//
	// Must be set together with `nextPageParam`.
	//
	// **Supported pagination styles:**
	//
	//   - **Cursor/token-based** — server returns an opaque token in the response body
	//     (e.g. `{"nextCursor": "abc123"}`). Set `nextPagePath: "nextCursor"` and the
	//     platform forwards it verbatim on the next request.
	//   - **Server-computed offset/page** — server echoes back the next offset or page
	//     number in the response body (e.g. `{"nextOffset": 50}` or `{"nextPage": 2}`).
	//     Set `nextPagePath: "nextOffset"` and the platform forwards the value as-is.
	//
	// **Not supported:**
	//
	//   - **Client-computed offset** — APIs where the client must compute
	//     `offset += limit` itself (e.g. `?offset=0&limit=50` with no next-offset in the
	//     response). Workaround: ask the API provider to return the next offset in the
	//     response body, or bake a fixed page size into the URL and use a server-side
	//     cursor instead.
	//   - **Client-computed page number** — APIs where the client increments `?page=N`
	//     itself with no next-page value in the response. Same workaround applies.
	//   - **Link header** — `Link: <url>; rel="next"` in HTTP response headers. The
	//     platform only inspects the response body.
	NextPagePath param.Opt[string] `json:"nextPagePath,omitzero"`
	// Query parameter name used to pass the extracted source value. **Required for GET
	// endpoints.** The value is URL-encoded and appended as
	// `?{queryParam}={sourceValue}`.
	//
	// Example: `queryParam: "q"` → `GET /products?q=blue+widget`
	QueryParam param.Opt[string] `json:"queryParam,omitzero"`
	// JMESPath expression applied to the response body to extract the enrichment
	// value. Omit to use the entire response body as the result.
	//
	// **For agent reasoning:** use a wildcard projection (e.g. `items[*]` or
	// `results[*].data`) so the endpoint's list of candidates is flattened into an
	// array before being passed to the LLM. A non-wildcard path (e.g. `data.product`)
	// extracts a single value treated as one candidate.
	//
	// **Response size:** the platform reads at most 50 MB of the response body before
	// decoding, regardless of the Content-Length header.
	ResponsePath param.Opt[string] `json:"responsePath,omitzero"`
	// Additional HTTP headers to include in every request (e.g.
	// `Authorization: Bearer <token>`).
	Headers any `json:"headers,omitzero"`
	paramObj
}

func (r EnrichConfigEndpointParam) MarshalJSON() (data []byte, err error) {
	type shadow EnrichConfigEndpointParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnrichConfigEndpointParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EnrichConfigEndpointParam](
		"method", "GET", "POST",
	)
}

// Single enrichment step configuration.
//
// **Process Flow (collection source):**
//
//  1. Extract values from `sourceField` using JMESPath
//  2. Perform search against the specified collection (semantic, exact, or hybrid
//     based on `searchMode`)
//  3. Return top K matches sorted by relevance (best match first)
//  4. Inject results into `targetField`
//
// **Process Flow (endpoint source):**
//
//  1. Extract values from `sourceField` using JMESPath
//  2. Call the named endpoint once per extracted value, following pagination if
//     `nextPagePath`/`nextPageParam` are configured on the endpoint
//  3. Optionally apply LLM agent reasoning to rank candidates
//     (`matchInstructions`), batching across all fetched pages in groups of
//     `maxCandidates`
//  4. Inject results into `targetField`
//
// **Collection Search Modes** (`source: "collection"` only):
//
//   - `semantic` (default): Vector similarity search — best for natural language and
//     conceptual matching
//   - `exact`: Exact keyword matching — best for SKU numbers, IDs, routing numbers
//   - `hybrid`: Combined semantic + keyword search — best for tags and categories
//
// **Result Format (collection source, exact mode — no re-ranking):**
//
// - Always an array sorted by relevance (best match first)
// - Each element: `{ id, data }`
//
// **Result Format (collection source, semantic/hybrid — re-ranking always on):**
//
//   - Re-ranking uses a fixed, built-in instruction to the LLM (rank the candidates
//     by how well each matches the source value); it is not configurable per step
//   - Array of matches, best first:
//     `[{ id, data, rank, confidence?, reasoning?, score?, scoreType? }, ...]`
//   - `id` is the collection item the match came from (e.g. `"clitm_…"`) — a durable
//     handle that survives edits to the item's data, and joins directly against the
//     collection. Where the same payload spans several rows (results are
//     de-duplicated by payload, and the uniqueness constraint is per collection +
//     embedding model), the oldest is the representative. It is how a candidate is
//     referenced when submitting ground-truth re-rankings via
//     `POST /v3/events/{eventID}/enrich-feedback`
//   - `rank` is 1-based (1 = best)
//   - `confidence` is the LLM's 0–1 score. It is present only for entries the LLM
//     ranked and **omitted** for backfilled entries (see below) — a missing
//     `confidence` means "not ranked by the LLM", not a score of 0
//   - `score` is the retrieval score and `scoreType` says which metric it is:
//     `"cosineDistance"` for semantic or `"hybridScore"` for hybrid. Both are 0–2
//     dissimilarities where **lower = better** — hybrid's Reciprocal Rank Fusion
//     score is mapped onto the same scale as cosine distance (0 = top of both
//     rankings). Included only when `includeScore` is set
//   - Results are de-duplicated by item payload, so they are distinct. Length is
//     `min(distinct candidates retrieved, topK)`; semantic additionally drops
//     candidates beyond `scoreThreshold`. The LLM re-orders the survivors; if it
//     ranks fewer than that length, the remaining survivors are backfilled in
//     retrieval (score) order with `confidence` omitted
//
// **Result Format (endpoint source, no matchInstructions):**
//
//   - Always an array; the raw fetched value is the single element
//   - These elements are the raw fetched values, so they carry no `id`. Ground-truth
//     re-ranking references candidates by `id`, so a field enriched this way cannot
//     be re-ranked
//
// **Result Format (endpoint source, with matchInstructions):**
//
//   - Array of LLM-ranked matches, best first:
//     `[{ id, data, rank, confidence, reasoning? }, ...]`
//   - `rank` is 1-based (1 = best); `confidence` is the LLM's 0–1 score
//   - `id` is a content hash of `data` (e.g. `"h_a5fef997ef9f8992"`) — identical
//     data always yields the same id. Endpoint candidates have no collection item to
//     name, so unlike collection matches they are identified by content; the `h_`
//     prefix tells the two apart
//   - Length capped by `enrichEndpoint.matchTopK` (default 1)
type EnrichStep struct {
	// JMESPath expression to extract source data. Can extract a single value or an
	// array. Each extracted value is looked up independently.
	SourceField string `json:"sourceField" api:"required"`
	// Field path where enriched results should be placed. Use simple field names
	// (e.g., "enriched_products"). Results are always injected as an array (list),
	// regardless of topK value.
	TargetField string `json:"targetField" api:"required"`
	// Name of the collection to search against. Required when `source` is
	// `"collection"`. The collection must exist and contain items. Supports
	// hierarchical paths when used with `includeSubcollections`.
	CollectionName string `json:"collectionName"`
	// Name of an endpoint defined in `enrichConfig.endpoints`. Required when `source`
	// is `"endpoint"`.
	EndpointName string `json:"endpointName"`
	// Whether to include retrieval scores in results.
	//
	// When enabled, each result includes a `score` field and a `scoreType` identifying
	// the metric:
	//
	//   - `"cosineDistance"` (semantic): 0.0 (perfect match) to 2.0 (completely
	//     dissimilar) — lower is better.
	//   - `"hybridScore"` (hybrid): an RRF score mapped onto cosine distance's 0–2 scale
	//     — lower is better (0.0 = top of both rankings).
	IncludeScore bool `json:"includeScore"`
	// When true, searches all collections under the hierarchical path. For example,
	// "customers" will match "customers", "customers.premium", etc.
	IncludeSubcollections bool `json:"includeSubcollections"`
	// Maximum cosine distance threshold for filtering results (default: 0.6). Results
	// with cosine distance above this threshold are excluded.
	//
	// **Applies to `semantic` and `hybrid` search modes.** For `hybrid`, the
	// Reciprocal Rank Fusion score is mapped onto the same 0–2 dissimilarity scale as
	// cosine distance, so a single threshold works for both. `exact` uses keyword
	// matching and ignores this setting. Note the default `0.6` is calibrated for
	// cosine distance and is relatively strict for hybrid.
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
	// **hybrid**: Fuses the dense (semantic) and sparse (keyword) rankings with
	// weighted Reciprocal Rank Fusion (k=60, 0.5 dense / 0.5 sparse). Because RRF
	// combines rank positions rather than raw scores, semantic meaning and exact-token
	// overlap contribute on the same scale.
	//
	// - Use for: Tags, categories, partial identifiers
	// - Example: Balances semantic meaning with exact keyword matching
	//
	// Any of "semantic", "exact", "hybrid".
	SearchMode EnrichStepSearchMode `json:"searchMode"`
	// Where to fetch enrichment data from (default: `"collection"`).
	//
	//   - `"collection"`: Vector/keyword search against a BEM collection. Requires
	//     `collectionName`.
	//   - `"endpoint"`: HTTP call to a named endpoint defined in
	//     `enrichConfig.endpoints`. Requires `endpointName`.
	//
	// Any of "collection", "endpoint".
	Source EnrichStepSource `json:"source"`
	// Number of top matching results to return per query (default: 1). Results are
	// always returned as an array (list), sorted best match first (by cosine distance
	// for `semantic`/`exact`, or by fused relevance score for `hybrid`). Duplicate
	// items are collapsed, so results are distinct: you get `topK` distinct matches
	// unless the collection contains fewer.
	//
	// - 1: Returns array with single best match: `[{...}]`
	// - > 1: Returns array with multiple matches: `[{...}, {...}, ...]`
	//
	// When re-ranking is on (the default for `semantic`/`hybrid`), `topK` is still the
	// number of results returned — re-ranking changes their order, not the count. The
	// candidate pool the LLM chooses from is widened internally to at least 5, so even
	// `topK: 1` re-ranks a real pool and returns the single best match.
	TopK int64 `json:"topK"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SourceField           respjson.Field
		TargetField           respjson.Field
		CollectionName        respjson.Field
		EndpointName          respjson.Field
		IncludeScore          respjson.Field
		IncludeSubcollections respjson.Field
		ScoreThreshold        respjson.Field
		SearchMode            respjson.Field
		Source                respjson.Field
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
// **hybrid**: Fuses the dense (semantic) and sparse (keyword) rankings with
// weighted Reciprocal Rank Fusion (k=60, 0.5 dense / 0.5 sparse). Because RRF
// combines rank positions rather than raw scores, semantic meaning and exact-token
// overlap contribute on the same scale.
//
// - Use for: Tags, categories, partial identifiers
// - Example: Balances semantic meaning with exact keyword matching
type EnrichStepSearchMode string

const (
	EnrichStepSearchModeSemantic EnrichStepSearchMode = "semantic"
	EnrichStepSearchModeExact    EnrichStepSearchMode = "exact"
	EnrichStepSearchModeHybrid   EnrichStepSearchMode = "hybrid"
)

// Where to fetch enrichment data from (default: `"collection"`).
//
//   - `"collection"`: Vector/keyword search against a BEM collection. Requires
//     `collectionName`.
//   - `"endpoint"`: HTTP call to a named endpoint defined in
//     `enrichConfig.endpoints`. Requires `endpointName`.
type EnrichStepSource string

const (
	EnrichStepSourceCollection EnrichStepSource = "collection"
	EnrichStepSourceEndpoint   EnrichStepSource = "endpoint"
)

// Single enrichment step configuration.
//
// **Process Flow (collection source):**
//
//  1. Extract values from `sourceField` using JMESPath
//  2. Perform search against the specified collection (semantic, exact, or hybrid
//     based on `searchMode`)
//  3. Return top K matches sorted by relevance (best match first)
//  4. Inject results into `targetField`
//
// **Process Flow (endpoint source):**
//
//  1. Extract values from `sourceField` using JMESPath
//  2. Call the named endpoint once per extracted value, following pagination if
//     `nextPagePath`/`nextPageParam` are configured on the endpoint
//  3. Optionally apply LLM agent reasoning to rank candidates
//     (`matchInstructions`), batching across all fetched pages in groups of
//     `maxCandidates`
//  4. Inject results into `targetField`
//
// **Collection Search Modes** (`source: "collection"` only):
//
//   - `semantic` (default): Vector similarity search — best for natural language and
//     conceptual matching
//   - `exact`: Exact keyword matching — best for SKU numbers, IDs, routing numbers
//   - `hybrid`: Combined semantic + keyword search — best for tags and categories
//
// **Result Format (collection source, exact mode — no re-ranking):**
//
// - Always an array sorted by relevance (best match first)
// - Each element: `{ id, data }`
//
// **Result Format (collection source, semantic/hybrid — re-ranking always on):**
//
//   - Re-ranking uses a fixed, built-in instruction to the LLM (rank the candidates
//     by how well each matches the source value); it is not configurable per step
//   - Array of matches, best first:
//     `[{ id, data, rank, confidence?, reasoning?, score?, scoreType? }, ...]`
//   - `id` is the collection item the match came from (e.g. `"clitm_…"`) — a durable
//     handle that survives edits to the item's data, and joins directly against the
//     collection. Where the same payload spans several rows (results are
//     de-duplicated by payload, and the uniqueness constraint is per collection +
//     embedding model), the oldest is the representative. It is how a candidate is
//     referenced when submitting ground-truth re-rankings via
//     `POST /v3/events/{eventID}/enrich-feedback`
//   - `rank` is 1-based (1 = best)
//   - `confidence` is the LLM's 0–1 score. It is present only for entries the LLM
//     ranked and **omitted** for backfilled entries (see below) — a missing
//     `confidence` means "not ranked by the LLM", not a score of 0
//   - `score` is the retrieval score and `scoreType` says which metric it is:
//     `"cosineDistance"` for semantic or `"hybridScore"` for hybrid. Both are 0–2
//     dissimilarities where **lower = better** — hybrid's Reciprocal Rank Fusion
//     score is mapped onto the same scale as cosine distance (0 = top of both
//     rankings). Included only when `includeScore` is set
//   - Results are de-duplicated by item payload, so they are distinct. Length is
//     `min(distinct candidates retrieved, topK)`; semantic additionally drops
//     candidates beyond `scoreThreshold`. The LLM re-orders the survivors; if it
//     ranks fewer than that length, the remaining survivors are backfilled in
//     retrieval (score) order with `confidence` omitted
//
// **Result Format (endpoint source, no matchInstructions):**
//
//   - Always an array; the raw fetched value is the single element
//   - These elements are the raw fetched values, so they carry no `id`. Ground-truth
//     re-ranking references candidates by `id`, so a field enriched this way cannot
//     be re-ranked
//
// **Result Format (endpoint source, with matchInstructions):**
//
//   - Array of LLM-ranked matches, best first:
//     `[{ id, data, rank, confidence, reasoning? }, ...]`
//   - `rank` is 1-based (1 = best); `confidence` is the LLM's 0–1 score
//   - `id` is a content hash of `data` (e.g. `"h_a5fef997ef9f8992"`) — identical
//     data always yields the same id. Endpoint candidates have no collection item to
//     name, so unlike collection matches they are identified by content; the `h_`
//     prefix tells the two apart
//   - Length capped by `enrichEndpoint.matchTopK` (default 1)
//
// The properties SourceField, TargetField are required.
type EnrichStepParam struct {
	// JMESPath expression to extract source data. Can extract a single value or an
	// array. Each extracted value is looked up independently.
	SourceField string `json:"sourceField" api:"required"`
	// Field path where enriched results should be placed. Use simple field names
	// (e.g., "enriched_products"). Results are always injected as an array (list),
	// regardless of topK value.
	TargetField string `json:"targetField" api:"required"`
	// Name of the collection to search against. Required when `source` is
	// `"collection"`. The collection must exist and contain items. Supports
	// hierarchical paths when used with `includeSubcollections`.
	CollectionName param.Opt[string] `json:"collectionName,omitzero"`
	// Name of an endpoint defined in `enrichConfig.endpoints`. Required when `source`
	// is `"endpoint"`.
	EndpointName param.Opt[string] `json:"endpointName,omitzero"`
	// Whether to include retrieval scores in results.
	//
	// When enabled, each result includes a `score` field and a `scoreType` identifying
	// the metric:
	//
	//   - `"cosineDistance"` (semantic): 0.0 (perfect match) to 2.0 (completely
	//     dissimilar) — lower is better.
	//   - `"hybridScore"` (hybrid): an RRF score mapped onto cosine distance's 0–2 scale
	//     — lower is better (0.0 = top of both rankings).
	IncludeScore param.Opt[bool] `json:"includeScore,omitzero"`
	// When true, searches all collections under the hierarchical path. For example,
	// "customers" will match "customers", "customers.premium", etc.
	IncludeSubcollections param.Opt[bool] `json:"includeSubcollections,omitzero"`
	// Maximum cosine distance threshold for filtering results (default: 0.6). Results
	// with cosine distance above this threshold are excluded.
	//
	// **Applies to `semantic` and `hybrid` search modes.** For `hybrid`, the
	// Reciprocal Rank Fusion score is mapped onto the same 0–2 dissimilarity scale as
	// cosine distance, so a single threshold works for both. `exact` uses keyword
	// matching and ignores this setting. Note the default `0.6` is calibrated for
	// cosine distance and is relatively strict for hybrid.
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
	// always returned as an array (list), sorted best match first (by cosine distance
	// for `semantic`/`exact`, or by fused relevance score for `hybrid`). Duplicate
	// items are collapsed, so results are distinct: you get `topK` distinct matches
	// unless the collection contains fewer.
	//
	// - 1: Returns array with single best match: `[{...}]`
	// - > 1: Returns array with multiple matches: `[{...}, {...}, ...]`
	//
	// When re-ranking is on (the default for `semantic`/`hybrid`), `topK` is still the
	// number of results returned — re-ranking changes their order, not the count. The
	// candidate pool the LLM chooses from is widened internally to at least 5, so even
	// `topK: 1` re-ranks a real pool and returns the single best match.
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
	// **hybrid**: Fuses the dense (semantic) and sparse (keyword) rankings with
	// weighted Reciprocal Rank Fusion (k=60, 0.5 dense / 0.5 sparse). Because RRF
	// combines rank positions rather than raw scores, semantic meaning and exact-token
	// overlap contribute on the same scale.
	//
	// - Use for: Tags, categories, partial identifiers
	// - Example: Balances semantic meaning with exact keyword matching
	//
	// Any of "semantic", "exact", "hybrid".
	SearchMode EnrichStepSearchMode `json:"searchMode,omitzero"`
	// Where to fetch enrichment data from (default: `"collection"`).
	//
	//   - `"collection"`: Vector/keyword search against a BEM collection. Requires
	//     `collectionName`.
	//   - `"endpoint"`: HTTP call to a named endpoint defined in
	//     `enrichConfig.endpoints`. Requires `endpointName`.
	//
	// Any of "collection", "endpoint".
	Source EnrichStepSource `json:"source,omitzero"`
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
// [FunctionTransform], [FunctionExtract], [FunctionAnalyze], [FunctionClassify],
// [FunctionSend], [FunctionSplit], [FunctionJoin], [FunctionPayloadShaping],
// [FunctionEnrich], [FunctionParse], [FunctionRender].
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
	// Any of "transform", "extract", "analyze", "classify", "send", "split", "join",
	// "payload_shaping", "enrich", "parse", "render".
	Type       string `json:"type"`
	VersionNum int64  `json:"versionNum"`
	// This field is from variant [FunctionTransform].
	Audit               FunctionAudit       `json:"audit"`
	DisplayName         string              `json:"displayName"`
	Tags                []string            `json:"tags"`
	UsedInWorkflows     []WorkflowUsageInfo `json:"usedInWorkflows"`
	EnableBoundingBoxes bool                `json:"enableBoundingBoxes"`
	PreCount            bool                `json:"preCount"`
	// This field is from variant [FunctionClassify].
	Classifications []ClassificationListItem `json:"classifications"`
	Description     string                   `json:"description"`
	// This field is from variant [FunctionClassify].
	NativeVisualInput bool `json:"nativeVisualInput"`
	// This field is from variant [FunctionSend].
	DestinationType SendDestinationType `json:"destinationType"`
	// This field is from variant [FunctionSend].
	GoogleDriveFolderID string `json:"googleDriveFolderId"`
	// This field is from variant [FunctionSend].
	S3Bucket string `json:"s3Bucket"`
	// This field is from variant [FunctionSend].
	S3Prefix string `json:"s3Prefix"`
	// This field is from variant [FunctionSend].
	WebhookSigningEnabled bool `json:"webhookSigningEnabled"`
	// This field is from variant [FunctionSend].
	WebhookURL string `json:"webhookUrl"`
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
	// This field is from variant [FunctionParse].
	ExtraConfig ParseExtraFunctionConfig `json:"extraConfig"`
	// This field is from variant [FunctionParse].
	ParseConfig ParseConfig `json:"parseConfig"`
	// This field is from variant [FunctionRender].
	RenderConfig RenderConfig `json:"renderConfig"`
	JSON         struct {
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
		NativeVisualInput       respjson.Field
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
		ExtraConfig             respjson.Field
		ParseConfig             respjson.Field
		RenderConfig            respjson.Field
		raw                     string
	} `json:"-"`
}

// anyFunction is implemented by each variant of [FunctionUnion] to add type safety
// for the return type of [FunctionUnion.AsAny]
type anyFunction interface {
	implFunctionUnion()
}

func (FunctionTransform) implFunctionUnion()      {}
func (FunctionExtract) implFunctionUnion()        {}
func (FunctionAnalyze) implFunctionUnion()        {}
func (FunctionClassify) implFunctionUnion()       {}
func (FunctionSend) implFunctionUnion()           {}
func (FunctionSplit) implFunctionUnion()          {}
func (FunctionJoin) implFunctionUnion()           {}
func (FunctionPayloadShaping) implFunctionUnion() {}
func (FunctionEnrich) implFunctionUnion()         {}
func (FunctionParse) implFunctionUnion()          {}
func (FunctionRender) implFunctionUnion()         {}

// Use the following switch statement to find the correct variant
//
//	switch variant := FunctionUnion.AsAny().(type) {
//	case bem.FunctionTransform:
//	case bem.FunctionExtract:
//	case bem.FunctionAnalyze:
//	case bem.FunctionClassify:
//	case bem.FunctionSend:
//	case bem.FunctionSplit:
//	case bem.FunctionJoin:
//	case bem.FunctionPayloadShaping:
//	case bem.FunctionEnrich:
//	case bem.FunctionParse:
//	case bem.FunctionRender:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u FunctionUnion) AsAny() anyFunction {
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
	case "parse":
		return u.AsParse()
	case "render":
		return u.AsRender()
	}
	return nil
}

func (u FunctionUnion) AsTransform() (v FunctionTransform) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionUnion) AsExtract() (v FunctionExtract) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionUnion) AsAnalyze() (v FunctionAnalyze) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionUnion) AsClassify() (v FunctionClassify) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionUnion) AsSend() (v FunctionSend) {
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

func (u FunctionUnion) AsParse() (v FunctionParse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionUnion) AsRender() (v FunctionRender) {
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

// A function that extracts structured JSON from documents and images. Accepts a
// wide range of input types including PDFs, images, spreadsheets, emails, and
// more.
type FunctionExtract struct {
	// Whether bounding box extraction is enabled. Applies to vision input types (pdf,
	// png, jpeg, heic, heif, webp) that dispatch through the analyze path. When true,
	// the function returns the document regions (page, coordinates) from which each
	// field was extracted.
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
	Type     constant.Extract `json:"type" default:"extract"`
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
func (r FunctionExtract) RawJSON() string { return r.JSON.raw }
func (r *FunctionExtract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionAnalyze struct {
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
func (r FunctionAnalyze) RawJSON() string { return r.JSON.raw }
func (r *FunctionAnalyze) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionClassify struct {
	// List of classifications a classify function can produce. Shares the underlying
	// route list shape.
	Classifications []ClassificationListItem `json:"classifications" api:"required"`
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
	// When true, image and PDF inputs are sent directly to the model for routing
	// instead of being OCR'd to text first. Defaults to true for new classify
	// functions and false for the legacy route type.
	NativeVisualInput bool `json:"nativeVisualInput"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags"`
	// List of workflows that use this function.
	UsedInWorkflows []WorkflowUsageInfo `json:"usedInWorkflows"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Classifications   respjson.Field
		Description       respjson.Field
		EmailAddress      respjson.Field
		FunctionID        respjson.Field
		FunctionName      respjson.Field
		Type              respjson.Field
		VersionNum        respjson.Field
		Audit             respjson.Field
		DisplayName       respjson.Field
		NativeVisualInput respjson.Field
		Tags              respjson.Field
		UsedInWorkflows   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionClassify) RawJSON() string { return r.JSON.raw }
func (r *FunctionClassify) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A function that delivers workflow outputs to an external destination. Send
// functions receive the output of an upstream workflow node and forward it to a
// webhook, S3 bucket, or Google Drive folder.
type FunctionSend struct {
	// Where the payload is delivered.
	//
	// Any of "webhook", "s3", "google_drive".
	DestinationType SendDestinationType `json:"destinationType" api:"required"`
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
func (r FunctionSend) RawJSON() string { return r.JSON.raw }
func (r *FunctionSend) UnmarshalJSON(data []byte) error {
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
	// Configuration for an enrich function.
	//
	// **How Enrich Functions Work:**
	//
	// Enrich functions augment JSON input with data from external sources. They take
	// JSON input (typically from a previous function), extract specified fields, fetch
	// or search for matching data, and inject the results back into the JSON.
	//
	// **Data Sources:**
	//
	//   - **Collections** (`source: "collection"`): Vector/keyword search against a BEM
	//     collection. Best for semantic matching against pre-indexed documents.
	//   - **Endpoints** (`source: "endpoint"`): HTTP call to any user-provided REST API.
	//     Best for looking up live data from CRMs, ERPs, or other external systems.
	//     Optionally uses LLM agent reasoning to rank candidates returned by the
	//     endpoint.
	//
	// **Input Requirements:**
	//
	// - Must receive JSON input (typically from a previous function's output)
	//
	// **Example Use Cases:**
	//
	//   - Match product descriptions to SKU codes from a product catalog collection
	//   - Enrich customer data with account details from a CRM endpoint
	//   - Use LLM agent reasoning to fuzzy-match line item descriptions to catalog
	//     products
	//
	// **Configuration:**
	//
	// - Define named endpoints (for endpoint-source steps)
	// - Define one or more enrichment steps; steps are executed sequentially
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

type FunctionParse struct {
	// Unique identifier of function.
	FunctionID string `json:"functionID" api:"required"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string         `json:"functionName" api:"required"`
	Type         constant.Parse `json:"type" default:"parse"`
	// Version number of function.
	VersionNum int64 `json:"versionNum" api:"required"`
	// Audit trail information for the function.
	Audit FunctionAudit `json:"audit"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName string `json:"displayName"`
	// Cross-cutting toggles for Parse functions. Mirrors the `extraConfig` surface on
	// Extract / Join — separated from `parseConfig` so the per-call Parse output shape
	// stays distinct from operator-level execution flags.
	ExtraConfig ParseExtraFunctionConfig `json:"extraConfig"`
	// Per-version configuration for a Parse function.
	//
	// Parse renders document pages (PDF, image) via vision LLM and emits structured
	// JSON. The two toggles below independently control entity extraction (a per-call
	// output concern) and cross-document memory linking (an environment-wide concern).
	ParseConfig ParseConfig `json:"parseConfig"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags"`
	// List of workflows that use this function.
	UsedInWorkflows []WorkflowUsageInfo `json:"usedInWorkflows"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FunctionID      respjson.Field
		FunctionName    respjson.Field
		Type            respjson.Field
		VersionNum      respjson.Field
		Audit           respjson.Field
		DisplayName     respjson.Field
		ExtraConfig     respjson.Field
		ParseConfig     respjson.Field
		Tags            respjson.Field
		UsedInWorkflows respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionParse) RawJSON() string { return r.JSON.raw }
func (r *FunctionParse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionRender struct {
	// Unique identifier of function.
	FunctionID string `json:"functionID" api:"required"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string          `json:"functionName" api:"required"`
	Type         constant.Render `json:"type" default:"render"`
	// Version number of function.
	VersionNum int64 `json:"versionNum" api:"required"`
	// Audit trail information for the function.
	Audit FunctionAudit `json:"audit"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName string `json:"displayName"`
	// Per-version configuration for a Render function.
	//
	// Render emits a `.docx` from schema-typed JSON by composing the JSON into a
	// `.docx` template. The template document is stored server-side; this response
	// exposes only the contract derived from it. Schema validation runs internally in
	// the ML service against the bundled core schema; no customer-supplied schema
	// rides this surface.
	RenderConfig RenderConfig `json:"renderConfig"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags"`
	// List of workflows that use this function.
	UsedInWorkflows []WorkflowUsageInfo `json:"usedInWorkflows"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FunctionID      respjson.Field
		FunctionName    respjson.Field
		Type            respjson.Field
		VersionNum      respjson.Field
		Audit           respjson.Field
		DisplayName     respjson.Field
		RenderConfig    respjson.Field
		Tags            respjson.Field
		UsedInWorkflows respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionRender) RawJSON() string { return r.JSON.raw }
func (r *FunctionRender) UnmarshalJSON(data []byte) error {
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
	FunctionTypeExtract        FunctionType = "extract"
	FunctionTypeRoute          FunctionType = "route"
	FunctionTypeClassify       FunctionType = "classify"
	FunctionTypeSend           FunctionType = "send"
	FunctionTypeSplit          FunctionType = "split"
	FunctionTypeJoin           FunctionType = "join"
	FunctionTypeAnalyze        FunctionType = "analyze"
	FunctionTypePayloadShaping FunctionType = "payload_shaping"
	FunctionTypeEnrich         FunctionType = "enrich"
	FunctionTypeParse          FunctionType = "parse"
	FunctionTypeRender         FunctionType = "render"
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

// Comparison of a single metric between two versions
type MetricComparison struct {
	// Value in baseline version (null if not available)
	BaselineValue float64 `json:"baselineValue" api:"nullable"`
	// Value in comparison version (null if not available)
	ComparisonValue float64 `json:"comparisonValue" api:"nullable"`
	// Absolute difference (comparisonValue - baselineValue)
	Difference float64 `json:"difference" api:"nullable"`
	// **Percentage change from baseline to comparison**
	//
	// Formula: ((comparisonValue - baselineValue) / baselineValue) \* 100
	//
	// - Positive values indicate improvement
	// - Negative values indicate regression
	LiftPercent float64 `json:"liftPercent" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaselineValue   respjson.Field
		ComparisonValue respjson.Field
		Difference      respjson.Field
		LiftPercent     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricComparison) RawJSON() string { return r.JSON.raw }
func (r *MetricComparison) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Comprehensive performance metrics
type Metrics struct {
	// Overall accuracy
	Accuracy float64 `json:"accuracy" api:"nullable"`
	// F1 Score (harmonic mean of precision and recall)
	F1Score float64 `json:"f1Score" api:"nullable"`
	// False Negatives
	Fn int64 `json:"fn"`
	// False Positives
	Fp int64 `json:"fp"`
	// Precision (TP / (TP + FP))
	Precision float64 `json:"precision" api:"nullable"`
	// Recall (TP / (TP + FN))
	Recall float64 `json:"recall" api:"nullable"`
	// True Negatives
	Tn int64 `json:"tn"`
	// True Positives
	Tp int64 `json:"tp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accuracy    respjson.Field
		F1Score     respjson.Field
		Fn          respjson.Field
		Fp          respjson.Field
		Precision   respjson.Field
		Recall      respjson.Field
		Tn          respjson.Field
		Tp          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Metrics) RawJSON() string { return r.JSON.raw }
func (r *Metrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Comparison of metrics between two versions
type MetricsComparison struct {
	// Comparison of a single metric between two versions
	Accuracy MetricComparison `json:"accuracy"`
	// Comparison of a single metric between two versions
	F1Score MetricComparison `json:"f1Score"`
	// Comparison of a single metric between two versions
	Precision MetricComparison `json:"precision"`
	// Comparison of a single metric between two versions
	Recall MetricComparison `json:"recall"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accuracy    respjson.Field
		F1Score     respjson.Field
		Precision   respjson.Field
		Recall      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsComparison) RawJSON() string { return r.JSON.raw }
func (r *MetricsComparison) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detailed performance metrics and analysis
type MetricsDetails struct {
	// Aggregate confusion matrix metrics across all fields
	AggregateMetrics Metrics `json:"aggregateMetrics"`
	// Enhanced field metrics with comprehensive analytics
	FieldMetrics []MetricsDetailsFieldMetric `json:"fieldMetrics"`
	// Area Under the Precision-Recall Curve
	PrecisionRecallAuc float64 `json:"precisionRecallAuc"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AggregateMetrics   respjson.Field
		FieldMetrics       respjson.Field
		PrecisionRecallAuc respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsDetails) RawJSON() string { return r.JSON.raw }
func (r *MetricsDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enhanced field metrics with comprehensive analytics
type MetricsDetailsFieldMetric struct {
	// JSON path to the field
	FieldPath string `json:"fieldPath" api:"required"`
	// Comprehensive performance metrics
	Metrics Metrics `json:"metrics"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FieldPath   respjson.Field
		Metrics     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetricsDetailsFieldMetric) RawJSON() string { return r.JSON.raw }
func (r *MetricsDetailsFieldMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-version configuration for a Parse function.
//
// Parse renders document pages (PDF, image) via vision LLM and emits structured
// JSON. The two toggles below independently control entity extraction (a per-call
// output concern) and cross-document memory linking (an environment-wide concern).
type ParseConfig struct {
	// Optional bucket NAME that parse-extracted entities land in when no call-level
	// bucket is supplied. Lower precedence than a call-level bucket, higher than the
	// account+environment default.
	DefaultBucket string `json:"defaultBucket"`
	// When true, extract named entities (people, organizations, products, studies,
	// identifiers, etc.) and the relationships between them, and dedupe by canonical
	// name within the document. When false, only `sections[]` is extracted;
	// `entities[]` and `relationships[]` come back empty in the parse output. Defaults
	// to true.
	ExtractEntities bool `json:"extractEntities"`
	// When true, link this document's entities to entities seen in earlier documents
	// in this environment, building one canonical record per real-world thing across
	// the corpus. Visible in the Memory tab and queryable via `POST /v3/fs` (op=find /
	// open / xref). Doesn't change this call's parse output. Requires
	// `extractEntities=true`. Defaults to true.
	LinkAcrossDocuments bool `json:"linkAcrossDocuments"`
	// Optional JSONSchema. When provided, each chunk performs schema-guided
	// extraction. When absent, chunks perform open-ended discovery and return
	// sections, entities, and relationships per the discovery schema.
	Schema any `json:"schema"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DefaultBucket       respjson.Field
		ExtractEntities     respjson.Field
		LinkAcrossDocuments respjson.Field
		Schema              respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseConfig) RawJSON() string { return r.JSON.raw }
func (r *ParseConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ParseConfig to a ParseConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ParseConfigParam.Overrides()
func (r ParseConfig) ToParam() ParseConfigParam {
	return param.Override[ParseConfigParam](json.RawMessage(r.RawJSON()))
}

// Per-version configuration for a Parse function.
//
// Parse renders document pages (PDF, image) via vision LLM and emits structured
// JSON. The two toggles below independently control entity extraction (a per-call
// output concern) and cross-document memory linking (an environment-wide concern).
type ParseConfigParam struct {
	// Optional bucket NAME that parse-extracted entities land in when no call-level
	// bucket is supplied. Lower precedence than a call-level bucket, higher than the
	// account+environment default.
	DefaultBucket param.Opt[string] `json:"defaultBucket,omitzero"`
	// When true, extract named entities (people, organizations, products, studies,
	// identifiers, etc.) and the relationships between them, and dedupe by canonical
	// name within the document. When false, only `sections[]` is extracted;
	// `entities[]` and `relationships[]` come back empty in the parse output. Defaults
	// to true.
	ExtractEntities param.Opt[bool] `json:"extractEntities,omitzero"`
	// When true, link this document's entities to entities seen in earlier documents
	// in this environment, building one canonical record per real-world thing across
	// the corpus. Visible in the Memory tab and queryable via `POST /v3/fs` (op=find /
	// open / xref). Doesn't change this call's parse output. Requires
	// `extractEntities=true`. Defaults to true.
	LinkAcrossDocuments param.Opt[bool] `json:"linkAcrossDocuments,omitzero"`
	// Optional JSONSchema. When provided, each chunk performs schema-guided
	// extraction. When absent, chunks perform open-ended discovery and return
	// sections, entities, and relationships per the discovery schema.
	Schema any `json:"schema,omitzero"`
	paramObj
}

func (r ParseConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow ParseConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParseConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cross-cutting toggles for Parse functions. Mirrors the `extraConfig` surface on
// Extract / Join — separated from `parseConfig` so the per-call Parse output shape
// stays distinct from operator-level execution flags.
type ParseExtraFunctionConfig struct {
	// When true, return per-section and per-entity-mention coordinates in the parse
	// event's `fieldBoundingBoxes` map (same shape as Extract: JSON Pointer key →
	// array of `{page, left, top, width, height}` with coordinates normalized to [0,
	// 1]). Keys are `/sections/{N}` and `/entities/{N}/occurrences/{M}` into the parse
	// output. Only applies to the open-ended discovery path (no `schema`) and to
	// vision input types. Bedrock-backed parse functions silently return an empty map
	// (no native bbox support). Defaults to false.
	EnableBoundingBoxes bool `json:"enableBoundingBoxes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnableBoundingBoxes respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseExtraFunctionConfig) RawJSON() string { return r.JSON.raw }
func (r *ParseExtraFunctionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ParseExtraFunctionConfig to a
// ParseExtraFunctionConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ParseExtraFunctionConfigParam.Overrides()
func (r ParseExtraFunctionConfig) ToParam() ParseExtraFunctionConfigParam {
	return param.Override[ParseExtraFunctionConfigParam](json.RawMessage(r.RawJSON()))
}

// Cross-cutting toggles for Parse functions. Mirrors the `extraConfig` surface on
// Extract / Join — separated from `parseConfig` so the per-call Parse output shape
// stays distinct from operator-level execution flags.
type ParseExtraFunctionConfigParam struct {
	// When true, return per-section and per-entity-mention coordinates in the parse
	// event's `fieldBoundingBoxes` map (same shape as Extract: JSON Pointer key →
	// array of `{page, left, top, width, height}` with coordinates normalized to [0,
	// 1]). Keys are `/sections/{N}` and `/entities/{N}/occurrences/{M}` into the parse
	// output. Only applies to the open-ended discovery path (no `schema`) and to
	// vision input types. Bedrock-backed parse functions silently return an empty map
	// (no native bbox support). Defaults to false.
	EnableBoundingBoxes param.Opt[bool] `json:"enableBoundingBoxes,omitzero"`
	paramObj
}

func (r ParseExtraFunctionConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow ParseExtraFunctionConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParseExtraFunctionConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Confidence interval for a rate/proportion using Wald (normal approximation)
// method by default.
//
// Wald confidence intervals use the normal approximation to the binomial
// distribution. For extreme rates or small sample sizes, Wilson confidence
// intervals may be more appropriate.
type RateConfidenceInterval struct {
	// Current number of samples/observations available
	CurrentSample int64 `json:"currentSample" api:"required"`
	// Minimum number of samples needed for reliable confidence interval calculation
	SampleNeeded int64 `json:"sampleNeeded" api:"required"`
	// Lower bound of the confidence interval (null if insufficient sample size)
	CiLower float64 `json:"ciLower" api:"nullable"`
	// Upper bound of the confidence interval (null if insufficient sample size)
	CiUpper float64 `json:"ciUpper" api:"nullable"`
	// Point estimate (observed rate) at the center of the interval (null if
	// insufficient sample size)
	Mid float64 `json:"mid" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentSample respjson.Field
		SampleNeeded  respjson.Field
		CiLower       respjson.Field
		CiUpper       respjson.Field
		Mid           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RateConfidenceInterval) RawJSON() string { return r.JSON.raw }
func (r *RateConfidenceInterval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-version configuration for a Render function.
//
// Render emits a `.docx` from schema-typed JSON by composing the JSON into a
// `.docx` template. The template document is stored server-side; this response
// exposes only the contract derived from it. Schema validation runs internally in
// the ML service against the bundled core schema; no customer-supplied schema
// rides this surface.
type RenderConfig struct {
	// The uploaded template: its filename, a short-lived presigned download URL, and
	// the placeholder/style contract derived from it. Absent on configs created before
	// template capture existed.
	Template RenderConfigTemplate `json:"template"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Template    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RenderConfig) RawJSON() string { return r.JSON.raw }
func (r *RenderConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The uploaded template: its filename, a short-lived presigned download URL, and
// the placeholder/style contract derived from it. Absent on configs created before
// template capture existed.
type RenderConfigTemplate struct {
	// Short-lived presigned URL to download the stored `.docx`. The private storage
	// location is never exposed.
	DownloadURL string `json:"downloadURL" format:"uri"`
	// Supported list kinds (`decimal`, `bullet`) the template's `numbering.xml`
	// defines an `abstractNum` for. Empty means the template can hold no list, so any
	// list primitive will fail at render.
	//
	// Any of "decimal", "bullet".
	ListKinds []string `json:"listKinds"`
	// Original filename of the uploaded template (e.g. `contract.docx`), echoed back
	// for display. Absent on templates uploaded before the filename was captured.
	Name string `json:"name"`
	// The placeholder contract derived from the template at create/update time. Absent
	// on configs created before create/update-time validation existed.
	Placeholders RenderConfigTemplatePlaceholders `json:"placeholders"`
	// Paragraph/character style IDs the uploaded template defines and the rendered
	// output can reference. Derived from the template's `styles.xml` at create/update
	// time.
	StyleIDs []string `json:"styleIds"`
	// Style IDs whose type is table — the styles a `table` primitive's required
	// `styleId` can name. Empty means the template defines no table style, so any
	// table primitive will fail at render.
	TableStyleIDs []string `json:"tableStyleIds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DownloadURL   respjson.Field
		ListKinds     respjson.Field
		Name          respjson.Field
		Placeholders  respjson.Field
		StyleIDs      respjson.Field
		TableStyleIDs respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RenderConfigTemplate) RawJSON() string { return r.JSON.raw }
func (r *RenderConfigTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The placeholder contract derived from the template at create/update time. Absent
// on configs created before create/update-time validation existed.
type RenderConfigTemplatePlaceholders struct {
	BlockKeys  []string `json:"blockKeys" api:"required"`
	StringKeys []string `json:"stringKeys" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BlockKeys   respjson.Field
		StringKeys  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RenderConfigTemplatePlaceholders) RawJSON() string { return r.JSON.raw }
func (r *RenderConfigTemplatePlaceholders) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request-side render configuration. Carries the template document as
// base64-encoded `.docx` bytes: the server validates them, stores the template,
// and derives the placeholder/style-id contract at create/update time, so clients
// never submit `placeholders` or `styleIds`. The response shape (`RenderConfig`)
// returns the derived contract.
//
// The property Template is required.
type RenderConfigInputParam struct {
	Template RenderConfigInputTemplateParam `json:"template,omitzero" api:"required"`
	paramObj
}

func (r RenderConfigInputParam) MarshalJSON() (data []byte, err error) {
	type shadow RenderConfigInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RenderConfigInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Base64 is required.
type RenderConfigInputTemplateParam struct {
	// Base64-encoded `.docx` bytes. In the Bem CLI, use `@path/to/file` to embed it
	// automatically.
	Base64 string `json:"base64" api:"required"`
	// Original upload filename (e.g. `contract.docx`), stored for display only. Does
	// not affect where the template is stored.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r RenderConfigInputTemplateParam) MarshalJSON() (data []byte, err error) {
	type shadow RenderConfigInputTemplateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RenderConfigInputTemplateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Destination type for a Send function.
type SendDestinationType string

const (
	SendDestinationTypeWebhook     SendDestinationType = "webhook"
	SendDestinationTypeS3          SendDestinationType = "s3"
	SendDestinationTypeGoogleDrive SendDestinationType = "google_drive"
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type UpdateFunctionUnionParam struct {
	OfExtract        *UpdateFunctionExtractParam        `json:",omitzero,inline"`
	OfClassify       *UpdateFunctionClassifyParam       `json:",omitzero,inline"`
	OfSend           *UpdateFunctionSendParam           `json:",omitzero,inline"`
	OfSplit          *UpdateFunctionSplitParam          `json:",omitzero,inline"`
	OfJoin           *UpdateFunctionJoinParam           `json:",omitzero,inline"`
	OfPayloadShaping *UpdateFunctionPayloadShapingParam `json:",omitzero,inline"`
	OfEnrich         *UpdateFunctionEnrichParam         `json:",omitzero,inline"`
	OfParse          *UpdateFunctionParseParam          `json:",omitzero,inline"`
	OfRender         *UpdateFunctionRenderParam         `json:",omitzero,inline"`
	paramUnion
}

func (u UpdateFunctionUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtract,
		u.OfClassify,
		u.OfSend,
		u.OfSplit,
		u.OfJoin,
		u.OfPayloadShaping,
		u.OfEnrich,
		u.OfParse,
		u.OfRender)
}
func (u *UpdateFunctionUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[UpdateFunctionUnionParam](
		"type",
		apijson.Discriminator[UpdateFunctionExtractParam]("extract"),
		apijson.Discriminator[UpdateFunctionClassifyParam]("classify"),
		apijson.Discriminator[UpdateFunctionSendParam]("send"),
		apijson.Discriminator[UpdateFunctionSplitParam]("split"),
		apijson.Discriminator[UpdateFunctionJoinParam]("join"),
		apijson.Discriminator[UpdateFunctionPayloadShapingParam]("payload_shaping"),
		apijson.Discriminator[UpdateFunctionEnrichParam]("enrich"),
		apijson.Discriminator[UpdateFunctionParseParam]("parse"),
		apijson.Discriminator[UpdateFunctionRenderParam]("render"),
	)
}

// The property Type is required.
type UpdateFunctionExtractParam struct {
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Whether bounding box extraction is enabled. Applies to vision input types (pdf,
	// png, jpeg, heic, heif, webp) that dispatch through the analyze path. When true,
	// the function returns the document regions (page, coordinates) from which each
	// field was extracted. Enabling this automatically configures the function to use
	// the bounding box model. Disabling resets to the default.
	EnableBoundingBoxes param.Opt[bool] `json:"enableBoundingBoxes,omitzero"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName param.Opt[string] `json:"functionName,omitzero"`
	// Name of output schema object.
	OutputSchemaName param.Opt[string] `json:"outputSchemaName,omitzero"`
	// Reducing the risk of the model stopping early on long documents. Trade-off:
	// Increases total latency. Compatible with `enableBoundingBoxes`.
	PreCount param.Opt[bool] `json:"preCount,omitzero"`
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

func (r UpdateFunctionExtractParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateFunctionExtractParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateFunctionExtractParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V3 create/update variants of the shared function payloads.
//
// The V3 Functions API no longer accepts the legacy `transform` or `analyze`
// function types when creating new functions or updating existing ones — both have
// been unified under `extract`. Existing functions of those types remain readable
// and callable via V3, so the V3 read-side unions still include `transform` and
// `analyze` variants.
//
// The V3 API also exposes `classify` in place of the legacy `route` type on
// create/update, with `classifications` in place of `routes`. Read-side
// `ClassifyFunction` / `ClassifyFunctionVersion` / `ClassificationList` are
// defined in the shared functions models and used by both the V2 and V3 response
// unions (existing classify functions are returned from V2 GET endpoints
// verbatim).V3 wire form of the classify function upsert payload.
//
// The property Type is required.
type UpdateFunctionClassifyParam struct {
	// Description of classifier. Can be used to provide additional context on
	// classifier's purpose and expected inputs.
	Description param.Opt[string] `json:"description,omitzero"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName param.Opt[string] `json:"functionName,omitzero"`
	// When true, image and PDF inputs are sent directly to the model for routing
	// instead of being OCR'd to text first. Defaults to true for new classify
	// functions and false for the legacy route type.
	NativeVisualInput param.Opt[bool] `json:"nativeVisualInput,omitzero"`
	// List of classifications a classify function can produce. Shares the underlying
	// route list shape.
	Classifications []ClassificationListItemParam `json:"classifications,omitzero"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags,omitzero"`
	// This field can be elided, and will marshal its zero value as "classify".
	Type constant.Classify `json:"type" default:"classify"`
	paramObj
}

func (r UpdateFunctionClassifyParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateFunctionClassifyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateFunctionClassifyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type UpdateFunctionSendParam struct {
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
	// Where the payload is delivered.
	//
	// Any of "webhook", "s3", "google_drive".
	DestinationType SendDestinationType `json:"destinationType,omitzero"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags,omitzero"`
	// This field can be elided, and will marshal its zero value as "send".
	Type constant.Send `json:"type" default:"send"`
	paramObj
}

func (r UpdateFunctionSendParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateFunctionSendParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateFunctionSendParam) UnmarshalJSON(data []byte) error {
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
	// Configuration for an enrich function.
	//
	// **How Enrich Functions Work:**
	//
	// Enrich functions augment JSON input with data from external sources. They take
	// JSON input (typically from a previous function), extract specified fields, fetch
	// or search for matching data, and inject the results back into the JSON.
	//
	// **Data Sources:**
	//
	//   - **Collections** (`source: "collection"`): Vector/keyword search against a BEM
	//     collection. Best for semantic matching against pre-indexed documents.
	//   - **Endpoints** (`source: "endpoint"`): HTTP call to any user-provided REST API.
	//     Best for looking up live data from CRMs, ERPs, or other external systems.
	//     Optionally uses LLM agent reasoning to rank candidates returned by the
	//     endpoint.
	//
	// **Input Requirements:**
	//
	// - Must receive JSON input (typically from a previous function's output)
	//
	// **Example Use Cases:**
	//
	//   - Match product descriptions to SKU codes from a product catalog collection
	//   - Enrich customer data with account details from a CRM endpoint
	//   - Use LLM agent reasoning to fuzzy-match line item descriptions to catalog
	//     products
	//
	// **Configuration:**
	//
	// - Define named endpoints (for endpoint-source steps)
	// - Define one or more enrichment steps; steps are executed sequentially
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

// The property Type is required.
type UpdateFunctionParseParam struct {
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName param.Opt[string] `json:"functionName,omitzero"`
	// Cross-cutting toggles for Parse functions. Mirrors the `extraConfig` surface on
	// Extract / Join — separated from `parseConfig` so the per-call Parse output shape
	// stays distinct from operator-level execution flags.
	ExtraConfig ParseExtraFunctionConfigParam `json:"extraConfig,omitzero"`
	// Per-version configuration for a Parse function.
	//
	// Parse renders document pages (PDF, image) via vision LLM and emits structured
	// JSON. The two toggles below independently control entity extraction (a per-call
	// output concern) and cross-document memory linking (an environment-wide concern).
	ParseConfig ParseConfigParam `json:"parseConfig,omitzero"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags,omitzero"`
	// This field can be elided, and will marshal its zero value as "parse".
	Type constant.Parse `json:"type" default:"parse"`
	paramObj
}

func (r UpdateFunctionParseParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateFunctionParseParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateFunctionParseParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type UpdateFunctionRenderParam struct {
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName param.Opt[string] `json:"functionName,omitzero"`
	// Request-side render configuration. Carries the template document as
	// base64-encoded `.docx` bytes: the server validates them, stores the template,
	// and derives the placeholder/style-id contract at create/update time, so clients
	// never submit `placeholders` or `styleIds`. The response shape (`RenderConfig`)
	// returns the derived contract.
	RenderConfig RenderConfigInputParam `json:"renderConfig,omitzero"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags,omitzero"`
	// This field can be elided, and will marshal its zero value as "render".
	Type constant.Render `json:"type" default:"render"`
	paramObj
}

func (r UpdateFunctionRenderParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateFunctionRenderParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateFunctionRenderParam) UnmarshalJSON(data []byte) error {
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

// **Response containing metrics comparison between two function versions**
//
// Shows absolute differences, lift percentages, and field-level changes.
type FunctionCompareMetricsResponse struct {
	// Baseline version number used for comparison
	BaselineVersionNum int64 `json:"baselineVersionNum" api:"required"`
	// Comparison version number
	ComparisonVersionNum int64 `json:"comparisonVersionNum" api:"required"`
	// Name of the compared function
	FunctionName string `json:"functionName" api:"required"`
	// Comparison of metrics between two versions
	AggregateComparison MetricsComparison `json:"aggregateComparison"`
	// Detailed performance metrics and analysis
	BaselineMetrics MetricsDetails `json:"baselineMetrics"`
	// Number of transformations used to calculate baseline metrics
	BaselineTransformationCount int64 `json:"baselineTransformationCount"`
	// Detailed performance metrics and analysis
	ComparisonMetrics MetricsDetails `json:"comparisonMetrics"`
	// Number of transformations used to calculate comparison metrics
	ComparisonTransformationCount int64 `json:"comparisonTransformationCount"`
	// **Field-level metrics that changed significantly**
	//
	// Only includes fields where metrics changed by more than 1%.
	FieldMetricsChanges []FunctionCompareMetricsResponseFieldMetricsChange `json:"fieldMetricsChanges"`
	// Optional message with additional details
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaselineVersionNum            respjson.Field
		ComparisonVersionNum          respjson.Field
		FunctionName                  respjson.Field
		AggregateComparison           respjson.Field
		BaselineMetrics               respjson.Field
		BaselineTransformationCount   respjson.Field
		ComparisonMetrics             respjson.Field
		ComparisonTransformationCount respjson.Field
		FieldMetricsChanges           respjson.Field
		Message                       respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionCompareMetricsResponse) RawJSON() string { return r.JSON.raw }
func (r *FunctionCompareMetricsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Comparison of field-level metrics
type FunctionCompareMetricsResponseFieldMetricsChange struct {
	// Comparison of metrics between two versions
	Comparison MetricsComparison `json:"comparison" api:"required"`
	// JSON pointer path to the field
	FieldPath string `json:"fieldPath" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Comparison  respjson.Field
		FieldPath   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionCompareMetricsResponseFieldMetricsChange) RawJSON() string { return r.JSON.raw }
func (r *FunctionCompareMetricsResponseFieldMetricsChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing review requirements estimate
type FunctionEstimateReviewRequirementsResponse struct {
	// Detailed review requirements estimate
	Estimate FunctionEstimateReviewRequirementsResponseEstimate `json:"estimate" api:"required"`
	// Name of the analyzed function
	FunctionName string `json:"functionName" api:"required"`
	// Version number of the function that was analyzed
	FunctionVersionNum int64 `json:"functionVersionNum" api:"required"`
	// Detailed performance metrics and analysis
	Metrics MetricsDetails `json:"metrics"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Estimate           respjson.Field
		FunctionName       respjson.Field
		FunctionVersionNum respjson.Field
		Metrics            respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionEstimateReviewRequirementsResponse) RawJSON() string { return r.JSON.raw }
func (r *FunctionEstimateReviewRequirementsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Detailed review requirements estimate
type FunctionEstimateReviewRequirementsResponseEstimate struct {
	// Distribution of confidence levels
	ConfidenceDistribution FunctionEstimateReviewRequirementsResponseEstimateConfidenceDistribution `json:"confidenceDistribution" api:"required"`
	// Number of transformations already labeled
	LabeledTransformations int64 `json:"labeledTransformations" api:"required"`
	// Number of transformations without evaluation data
	MissingEvaluations int64 `json:"missingEvaluations" api:"required"`
	// Statistical analysis across confidence thresholds
	ThresholdMatrix []FunctionEstimateReviewRequirementsResponseEstimateThresholdMatrix `json:"thresholdMatrix" api:"required"`
	// Total number of transformations analyzed
	TotalTransformations int64 `json:"totalTransformations" api:"required"`
	// Number of transformations not yet labeled
	UnlabeledTransformations int64 `json:"unlabeledTransformations" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConfidenceDistribution   respjson.Field
		LabeledTransformations   respjson.Field
		MissingEvaluations       respjson.Field
		ThresholdMatrix          respjson.Field
		TotalTransformations     respjson.Field
		UnlabeledTransformations respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionEstimateReviewRequirementsResponseEstimate) RawJSON() string { return r.JSON.raw }
func (r *FunctionEstimateReviewRequirementsResponseEstimate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Distribution of confidence levels
type FunctionEstimateReviewRequirementsResponseEstimateConfidenceDistribution struct {
	High   int64 `json:"high"`
	Low    int64 `json:"low"`
	Medium int64 `json:"medium"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		High        respjson.Field
		Low         respjson.Field
		Medium      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionEstimateReviewRequirementsResponseEstimateConfidenceDistribution) RawJSON() string {
	return r.JSON.raw
}
func (r *FunctionEstimateReviewRequirementsResponseEstimateConfidenceDistribution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Results for a specific confidence threshold analysis
type FunctionEstimateReviewRequirementsResponseEstimateThresholdMatrix struct {
	// False Negatives
	Fn int64 `json:"fn" api:"required"`
	// False Positives
	Fp int64 `json:"fp" api:"required"`
	// Confidence threshold value
	Threshold float64 `json:"threshold" api:"required"`
	// True Negatives
	Tn int64 `json:"tn" api:"required"`
	// True Positives
	Tp int64 `json:"tp" api:"required"`
	// Accuracy confidence intervals for samples above threshold, by confidence level.
	// Keys are confidence levels as strings ("90", "95", "99"). Values contain
	// statistical confidence intervals.
	AccuracyAboveThreshold FunctionEstimateReviewRequirementsResponseEstimateThresholdMatrixAccuracyAboveThreshold `json:"accuracyAboveThreshold"`
	// False Discovery Rate confidence intervals by confidence level. Keys are
	// confidence levels as strings ("90", "95", "99"). Values contain statistical
	// confidence intervals.
	FalseDiscoveryRate FunctionEstimateReviewRequirementsResponseEstimateThresholdMatrixFalseDiscoveryRate `json:"falseDiscoveryRate"`
	// False Positive Rate confidence intervals by confidence level. Keys are
	// confidence levels as strings ("90", "95", "99"). Values contain statistical
	// confidence intervals.
	FalsePositiveRate FunctionEstimateReviewRequirementsResponseEstimateThresholdMatrixFalsePositiveRate `json:"falsePositiveRate"`
	// Precision confidence intervals by confidence level. Keys are confidence levels
	// as strings ("90", "95", "99"). Values contain statistical confidence intervals.
	Precision FunctionEstimateReviewRequirementsResponseEstimateThresholdMatrixPrecision `json:"precision"`
	// Recall confidence intervals by confidence level. Keys are confidence levels as
	// strings ("90", "95", "99"). Values contain statistical confidence intervals.
	Recall FunctionEstimateReviewRequirementsResponseEstimateThresholdMatrixRecall `json:"recall"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Fn                     respjson.Field
		Fp                     respjson.Field
		Threshold              respjson.Field
		Tn                     respjson.Field
		Tp                     respjson.Field
		AccuracyAboveThreshold respjson.Field
		FalseDiscoveryRate     respjson.Field
		FalsePositiveRate      respjson.Field
		Precision              respjson.Field
		Recall                 respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionEstimateReviewRequirementsResponseEstimateThresholdMatrix) RawJSON() string {
	return r.JSON.raw
}
func (r *FunctionEstimateReviewRequirementsResponseEstimateThresholdMatrix) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Accuracy confidence intervals for samples above threshold, by confidence level.
// Keys are confidence levels as strings ("90", "95", "99"). Values contain
// statistical confidence intervals.
type FunctionEstimateReviewRequirementsResponseEstimateThresholdMatrixAccuracyAboveThreshold struct {
	// Confidence interval for a rate/proportion using Wald (normal approximation)
	// method by default.
	//
	// Wald confidence intervals use the normal approximation to the binomial
	// distribution. For extreme rates or small sample sizes, Wilson confidence
	// intervals may be more appropriate.
	Number95 RateConfidenceInterval `json:"95"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Number95    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionEstimateReviewRequirementsResponseEstimateThresholdMatrixAccuracyAboveThreshold) RawJSON() string {
	return r.JSON.raw
}
func (r *FunctionEstimateReviewRequirementsResponseEstimateThresholdMatrixAccuracyAboveThreshold) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// False Discovery Rate confidence intervals by confidence level. Keys are
// confidence levels as strings ("90", "95", "99"). Values contain statistical
// confidence intervals.
type FunctionEstimateReviewRequirementsResponseEstimateThresholdMatrixFalseDiscoveryRate struct {
	// Confidence interval for a rate/proportion using Wald (normal approximation)
	// method by default.
	//
	// Wald confidence intervals use the normal approximation to the binomial
	// distribution. For extreme rates or small sample sizes, Wilson confidence
	// intervals may be more appropriate.
	Number95 RateConfidenceInterval `json:"95"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Number95    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionEstimateReviewRequirementsResponseEstimateThresholdMatrixFalseDiscoveryRate) RawJSON() string {
	return r.JSON.raw
}
func (r *FunctionEstimateReviewRequirementsResponseEstimateThresholdMatrixFalseDiscoveryRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// False Positive Rate confidence intervals by confidence level. Keys are
// confidence levels as strings ("90", "95", "99"). Values contain statistical
// confidence intervals.
type FunctionEstimateReviewRequirementsResponseEstimateThresholdMatrixFalsePositiveRate struct {
	// Confidence interval for a rate/proportion using Wald (normal approximation)
	// method by default.
	//
	// Wald confidence intervals use the normal approximation to the binomial
	// distribution. For extreme rates or small sample sizes, Wilson confidence
	// intervals may be more appropriate.
	Number95 RateConfidenceInterval `json:"95"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Number95    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionEstimateReviewRequirementsResponseEstimateThresholdMatrixFalsePositiveRate) RawJSON() string {
	return r.JSON.raw
}
func (r *FunctionEstimateReviewRequirementsResponseEstimateThresholdMatrixFalsePositiveRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Precision confidence intervals by confidence level. Keys are confidence levels
// as strings ("90", "95", "99"). Values contain statistical confidence intervals.
type FunctionEstimateReviewRequirementsResponseEstimateThresholdMatrixPrecision struct {
	// Confidence interval for a rate/proportion using Wald (normal approximation)
	// method by default.
	//
	// Wald confidence intervals use the normal approximation to the binomial
	// distribution. For extreme rates or small sample sizes, Wilson confidence
	// intervals may be more appropriate.
	Number95 RateConfidenceInterval `json:"95"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Number95    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionEstimateReviewRequirementsResponseEstimateThresholdMatrixPrecision) RawJSON() string {
	return r.JSON.raw
}
func (r *FunctionEstimateReviewRequirementsResponseEstimateThresholdMatrixPrecision) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Recall confidence intervals by confidence level. Keys are confidence levels as
// strings ("90", "95", "99"). Values contain statistical confidence intervals.
type FunctionEstimateReviewRequirementsResponseEstimateThresholdMatrixRecall struct {
	// Confidence interval for a rate/proportion using Wald (normal approximation)
	// method by default.
	//
	// Wald confidence intervals use the normal approximation to the binomial
	// distribution. For extreme rates or small sample sizes, Wilson confidence
	// intervals may be more appropriate.
	Number95 RateConfidenceInterval `json:"95"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Number95    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionEstimateReviewRequirementsResponseEstimateThresholdMatrixRecall) RawJSON() string {
	return r.JSON.raw
}
func (r *FunctionEstimateReviewRequirementsResponseEstimateThresholdMatrixRecall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionGetMetricsResponse struct {
	Functions []FunctionGetMetricsResponseFunction `json:"functions" api:"required"`
	// Total number of functions
	TotalCount int64 `json:"totalCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Functions   respjson.Field
		TotalCount  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionGetMetricsResponse) RawJSON() string { return r.JSON.raw }
func (r *FunctionGetMetricsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionGetMetricsResponseFunction struct {
	// The function name
	FunctionName string                                    `json:"functionName" api:"required"`
	Metrics      FunctionGetMetricsResponseFunctionMetrics `json:"metrics" api:"required"`
	// Number of transformations that have been labeled/evaluated for metrics
	// calculation
	TotalLabeledResults int64 `json:"totalLabeledResults" api:"required"`
	// Total number of results processed by the function
	TotalResults int64 `json:"totalResults" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FunctionName        respjson.Field
		Metrics             respjson.Field
		TotalLabeledResults respjson.Field
		TotalResults        respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionGetMetricsResponseFunction) RawJSON() string { return r.JSON.raw }
func (r *FunctionGetMetricsResponseFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionGetMetricsResponseFunctionMetrics struct {
	Accuracy  float64 `json:"accuracy" api:"required"`
	F1Score   float64 `json:"f1Score" api:"required"`
	Fn        int64   `json:"fn" api:"required"`
	Fp        int64   `json:"fp" api:"required"`
	Precision float64 `json:"precision" api:"required"`
	Recall    float64 `json:"recall" api:"required"`
	Tn        int64   `json:"tn" api:"required"`
	Tp        int64   `json:"tp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accuracy    respjson.Field
		F1Score     respjson.Field
		Fn          respjson.Field
		Fp          respjson.Field
		Precision   respjson.Field
		Recall      respjson.Field
		Tn          respjson.Field
		Tp          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionGetMetricsResponseFunctionMetrics) RawJSON() string { return r.JSON.raw }
func (r *FunctionGetMetricsResponseFunctionMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionNewParams struct {
	// V3 wire form of the classify function create payload.
	CreateFunction CreateFunctionUnionParam
	paramObj
}

func (r FunctionNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateFunction)
}
func (r *FunctionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionGetParams struct {
	// Populate the function's `extraConfig` block. Omitted or `false` by default, in
	// which case `extraConfig` is absent from the response.
	IncludeExtraSettings param.Opt[bool] `query:"includeExtraSettings,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FunctionGetParams]'s query parameters as `url.Values`.
func (r FunctionGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FunctionUpdateParams struct {
	// V3 create/update variants of the shared function payloads.
	//
	// The V3 Functions API no longer accepts the legacy `transform` or `analyze`
	// function types when creating new functions or updating existing ones — both have
	// been unified under `extract`. Existing functions of those types remain readable
	// and callable via V3, so the V3 read-side unions still include `transform` and
	// `analyze` variants.
	//
	// The V3 API also exposes `classify` in place of the legacy `route` type on
	// create/update, with `classifications` in place of `routes`. Read-side
	// `ClassifyFunction` / `ClassifyFunctionVersion` / `ClassificationList` are
	// defined in the shared functions models and used by both the V2 and V3 response
	// unions (existing classify functions are returned from V2 GET endpoints
	// verbatim).V3 wire form of the classify function upsert payload.
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
	DisplayName  param.Opt[string] `query:"displayName,omitzero" json:"-"`
	EndingBefore param.Opt[string] `query:"endingBefore,omitzero" json:"-"`
	// Populate each function's `extraConfig` block. Omitted or `false` by default, in
	// which case `extraConfig` is absent from the response.
	IncludeExtraSettings param.Opt[bool]   `query:"includeExtraSettings,omitzero" json:"-"`
	Limit                param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	StartingAfter        param.Opt[string] `query:"startingAfter,omitzero" json:"-"`
	FunctionIDs          []string          `query:"functionIDs,omitzero" json:"-"`
	FunctionNames        []string          `query:"functionNames,omitzero" json:"-"`
	// Any of "asc", "desc".
	SortOrder   FunctionListParamsSortOrder `query:"sortOrder,omitzero" json:"-"`
	Tags        []string                    `query:"tags,omitzero" json:"-"`
	Types       []FunctionType              `query:"types,omitzero" json:"-"`
	WorkflowIDs []string                    `query:"workflowIDs,omitzero" json:"-"`
	// Return only functions referenced by a specific workflow version. Each entry is
	// `<workflowID>.<versionNum>` — for example `wf_2c9AXIj48cUYJtCuv1gsQtHGDzK.3`.
	WorkflowIDVersionNums []string `query:"workflowIDVersionNums,omitzero" json:"-"`
	WorkflowNames         []string `query:"workflowNames,omitzero" json:"-"`
	// Return only functions referenced by a specific workflow version, keyed by
	// workflow name. Each entry is `<workflowName>.<versionNum>` — for example
	// `invoice-pipeline.3`.
	WorkflowNameVersionNums []string `query:"workflowNameVersionNums,omitzero" json:"-"`
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

type FunctionCompareMetricsParams struct {
	// Name of the function to compare versions for
	FunctionName string `json:"functionName" api:"required"`
	// **Baseline version number for comparison**
	//
	// If not provided, defaults to the previous version (current - 1).
	BaselineVersionNum param.Opt[int64] `json:"baselineVersionNum,omitzero"`
	// **Comparison version number**
	//
	// If not provided, defaults to the current version.
	ComparisonVersionNum param.Opt[int64] `json:"comparisonVersionNum,omitzero"`
	// **Whether to compare regression test data only**
	//
	// If true, only compares transformations marked as regression tests.
	IsRegression param.Opt[bool] `json:"isRegression,omitzero"`
	paramObj
}

func (r FunctionCompareMetricsParams) MarshalJSON() (data []byte, err error) {
	type shadow FunctionCompareMetricsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionCompareMetricsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionEstimateReviewRequirementsParams struct {
	// Name of the function to analyze
	FunctionName string `json:"functionName" api:"required"`
	// Optional function version number to analyze. If not provided, uses the
	// latest/current version of the function.
	FunctionVersionNum param.Opt[int64] `json:"functionVersionNum,omitzero"`
	// Internal flag indicating if the request is from a regression test
	IsRegression param.Opt[bool] `json:"isRegression,omitzero"`
	// Margin of error for statistical calculations
	MarginOfError param.Opt[float64] `json:"marginOfError,omitzero"`
	// Maximum confidence threshold to analyze
	ThresholdMax param.Opt[float64] `json:"thresholdMax,omitzero"`
	// Minimum confidence threshold to analyze
	ThresholdMin param.Opt[float64] `json:"thresholdMin,omitzero"`
	// Step size for threshold analysis (smaller = more granular)
	ThresholdStep param.Opt[float64] `json:"thresholdStep,omitzero"`
	// Confidence levels for statistical analysis as integers representing percentages
	// (e.g., [90, 95, 99] for 90%, 95%, 99%). IMPORTANT: Only integers are accepted,
	// floats like 0.95 will be rejected.
	ConfidenceLevels []int64 `json:"confidenceLevels,omitzero"`
	// Confidence interval calculation method (default "wald").
	//
	// - "wald": Normal approximation method (faster, standard)
	// - "wilson": Wilson score interval (more robust for extreme rates)
	//
	// Any of "wald", "wilson".
	ConfidenceMethod FunctionEstimateReviewRequirementsParamsConfidenceMethod `json:"confidenceMethod,omitzero"`
	// Optional evaluation version to filter evaluations by. Must be one of the
	// supported versions. If not provided, defaults to "0.1.0-gemini".
	//
	// Any of "0.1.0-gemini".
	EvaluationVersion FunctionEstimateReviewRequirementsParamsEvaluationVersion `json:"evaluationVersion,omitzero"`
	paramObj
}

func (r FunctionEstimateReviewRequirementsParams) MarshalJSON() (data []byte, err error) {
	type shadow FunctionEstimateReviewRequirementsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionEstimateReviewRequirementsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Confidence interval calculation method (default "wald").
//
// - "wald": Normal approximation method (faster, standard)
// - "wilson": Wilson score interval (more robust for extreme rates)
type FunctionEstimateReviewRequirementsParamsConfidenceMethod string

const (
	FunctionEstimateReviewRequirementsParamsConfidenceMethodWald   FunctionEstimateReviewRequirementsParamsConfidenceMethod = "wald"
	FunctionEstimateReviewRequirementsParamsConfidenceMethodWilson FunctionEstimateReviewRequirementsParamsConfidenceMethod = "wilson"
)

// Optional evaluation version to filter evaluations by. Must be one of the
// supported versions. If not provided, defaults to "0.1.0-gemini".
type FunctionEstimateReviewRequirementsParamsEvaluationVersion string

const (
	FunctionEstimateReviewRequirementsParamsEvaluationVersion0_1_0Gemini FunctionEstimateReviewRequirementsParamsEvaluationVersion = "0.1.0-gemini"
)

type FunctionGetMetricsParams struct {
	// Case-insensitive substring match on the function display name.
	DisplayName param.Opt[string] `query:"displayName,omitzero" json:"-"`
	// Cursor — a `functionID` defining your place in the list.
	EndingBefore param.Opt[string] `query:"endingBefore,omitzero" json:"-"`
	Limit        param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	// Cursor — a `functionID` defining your place in the list.
	StartingAfter param.Opt[string] `query:"startingAfter,omitzero" json:"-"`
	FunctionIDs   []string          `query:"functionIDs,omitzero" json:"-"`
	FunctionNames []string          `query:"functionNames,omitzero" json:"-"`
	// Sort direction over the result set (default `asc`). Pagination works
	// symmetrically in both directions via `startingAfter` / `endingBefore`.
	//
	// Any of "asc", "desc".
	SortOrder FunctionGetMetricsParamsSortOrder `query:"sortOrder,omitzero" json:"-"`
	// Returns metrics for functions tagged with any of the supplied tags.
	Tags  []string       `query:"tags,omitzero" json:"-"`
	Types []FunctionType `query:"types,omitzero" json:"-"`
	// Returns metrics only for functions referenced by the named workflows.
	WorkflowIDs []string `query:"workflowIDs,omitzero" json:"-"`
	// Narrow the workflow filter to a specific workflow version. Each entry is
	// `<workflowID>.<versionNum>`.
	WorkflowIDVersionNums []string `query:"workflowIDVersionNums,omitzero" json:"-"`
	// Returns metrics only for functions referenced by the named workflows.
	WorkflowNames []string `query:"workflowNames,omitzero" json:"-"`
	// Narrow the workflow filter to a specific workflow version, keyed by workflow
	// name. Each entry is `<workflowName>.<versionNum>`.
	WorkflowNameVersionNums []string `query:"workflowNameVersionNums,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FunctionGetMetricsParams]'s query parameters as
// `url.Values`.
func (r FunctionGetMetricsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort direction over the result set (default `asc`). Pagination works
// symmetrically in both directions via `startingAfter` / `endingBefore`.
type FunctionGetMetricsParamsSortOrder string

const (
	FunctionGetMetricsParamsSortOrderAsc  FunctionGetMetricsParamsSortOrder = "asc"
	FunctionGetMetricsParamsSortOrderDesc FunctionGetMetricsParamsSortOrder = "desc"
)
