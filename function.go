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

// **Create a function.**
//
// The function type (`extract`, `classify`, `split`, `join`, `enrich`, or
// `payload_shaping`) determines which configuration fields are required — see
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
		u.OfParse)
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

func (r CreateFunctionSendParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateFunctionSendParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateFunctionSendParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CreateFunctionSendParam](
		"destinationType", "webhook", "s3", "google_drive",
	)
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

// The properties FunctionName, Type are required.
type CreateFunctionParseParam struct {
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string `json:"functionName" api:"required"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Per-version configuration for a Parse function.
	//
	// Parse renders document pages (PDF, image) via vision LLM and emits structured
	// JSON. The two toggles below independently control entity extraction (a per-call
	// output concern) and cross-document memory linking (an environment-wide concern).
	ParseConfig CreateFunctionParseParseConfigParam `json:"parseConfig,omitzero"`
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

// Per-version configuration for a Parse function.
//
// Parse renders document pages (PDF, image) via vision LLM and emits structured
// JSON. The two toggles below independently control entity extraction (a per-call
// output concern) and cross-document memory linking (an environment-wide concern).
type CreateFunctionParseParseConfigParam struct {
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

func (r CreateFunctionParseParseConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateFunctionParseParseConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateFunctionParseParseConfigParam) UnmarshalJSON(data []byte) error {
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
// [FunctionTransform], [FunctionExtract], [FunctionAnalyze], [FunctionClassify],
// [FunctionSend], [FunctionSplit], [FunctionJoin], [FunctionPayloadShaping],
// [FunctionEnrich], [FunctionParse].
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
	// "payload_shaping", "enrich", "parse".
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
	// This field is from variant [FunctionSend].
	DestinationType string `json:"destinationType"`
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
	ParseConfig FunctionParseParseConfig `json:"parseConfig"`
	JSON        struct {
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
		ParseConfig             respjson.Field
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
func (r FunctionClassify) RawJSON() string { return r.JSON.raw }
func (r *FunctionClassify) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A function that delivers workflow outputs to an external destination. Send
// functions receive the output of an upstream workflow node and forward it to a
// webhook, S3 bucket, or Google Drive folder.
type FunctionSend struct {
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
	// Per-version configuration for a Parse function.
	//
	// Parse renders document pages (PDF, image) via vision LLM and emits structured
	// JSON. The two toggles below independently control entity extraction (a per-call
	// output concern) and cross-document memory linking (an environment-wide concern).
	ParseConfig FunctionParseParseConfig `json:"parseConfig"`
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

// Per-version configuration for a Parse function.
//
// Parse renders document pages (PDF, image) via vision LLM and emits structured
// JSON. The two toggles below independently control entity extraction (a per-call
// output concern) and cross-document memory linking (an environment-wide concern).
type FunctionParseParseConfig struct {
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
		ExtractEntities     respjson.Field
		LinkAcrossDocuments respjson.Field
		Schema              respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionParseParseConfig) RawJSON() string { return r.JSON.raw }
func (r *FunctionParseParseConfig) UnmarshalJSON(data []byte) error {
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
		u.OfParse)
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

func (r UpdateFunctionSendParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateFunctionSendParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateFunctionSendParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UpdateFunctionSendParam](
		"destinationType", "webhook", "s3", "google_drive",
	)
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

// The property Type is required.
type UpdateFunctionParseParam struct {
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName param.Opt[string] `json:"functionName,omitzero"`
	// Per-version configuration for a Parse function.
	//
	// Parse renders document pages (PDF, image) via vision LLM and emits structured
	// JSON. The two toggles below independently control entity extraction (a per-call
	// output concern) and cross-document memory linking (an environment-wide concern).
	ParseConfig UpdateFunctionParseParseConfigParam `json:"parseConfig,omitzero"`
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

// Per-version configuration for a Parse function.
//
// Parse renders document pages (PDF, image) via vision LLM and emits structured
// JSON. The two toggles below independently control entity extraction (a per-call
// output concern) and cross-document memory linking (an environment-wide concern).
type UpdateFunctionParseParseConfigParam struct {
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

func (r UpdateFunctionParseParseConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateFunctionParseParseConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateFunctionParseParseConfigParam) UnmarshalJSON(data []byte) error {
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
