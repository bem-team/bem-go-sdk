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
	"github.com/bem-team/bem-go-sdk/internal/requestconfig"
	"github.com/bem-team/bem-go-sdk/option"
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
// FunctionVersionService contains methods and other services that help with
// interacting with the bem API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFunctionVersionService] method instead.
type FunctionVersionService struct {
	options []option.RequestOption
}

// NewFunctionVersionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFunctionVersionService(opts ...option.RequestOption) (r FunctionVersionService) {
	r = FunctionVersionService{}
	r.options = opts
	return
}

// Get a Function Version
func (r *FunctionVersionService) Get(ctx context.Context, versionNum int64, query FunctionVersionGetParams, opts ...option.RequestOption) (res *FunctionVersionGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.FunctionName == "" {
		err = errors.New("missing required functionName parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/functions/%s/versions/%v", url.PathEscape(query.FunctionName), versionNum)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List Function Versions
func (r *FunctionVersionService) List(ctx context.Context, functionName string, opts ...option.RequestOption) (res *FunctionVersionListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if functionName == "" {
		err = errors.New("missing required functionName parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/functions/%s/versions", url.PathEscape(functionName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Single-function-version response wrapper used by V3 endpoints.
type FunctionVersionGetResponse struct {
	// V3 read-side union for function versions. Same shape as the shared
	// `FunctionVersion` union but with `classify` in place of `route`.
	Function FunctionVersionGetResponseFunctionUnion `json:"function" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Function    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionVersionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunctionVersionGetResponseFunctionUnion contains all possible properties and
// values from [FunctionVersionGetResponseFunctionTransform],
// [FunctionVersionGetResponseFunctionExtract],
// [FunctionVersionGetResponseFunctionAnalyze],
// [FunctionVersionGetResponseFunctionClassify],
// [FunctionVersionGetResponseFunctionSend],
// [FunctionVersionGetResponseFunctionSplit],
// [FunctionVersionGetResponseFunctionJoin],
// [FunctionVersionGetResponseFunctionEnrich],
// [FunctionVersionGetResponseFunctionPayloadShaping].
//
// Use the [FunctionVersionGetResponseFunctionUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FunctionVersionGetResponseFunctionUnion struct {
	EmailAddress           string `json:"emailAddress"`
	FunctionID             string `json:"functionID"`
	FunctionName           string `json:"functionName"`
	OutputSchema           any    `json:"outputSchema"`
	OutputSchemaName       string `json:"outputSchemaName"`
	TabularChunkingEnabled bool   `json:"tabularChunkingEnabled"`
	// Any of "transform", "extract", "analyze", "classify", "send", "split", "join",
	// "enrich", "payload_shaping".
	Type       string `json:"type"`
	VersionNum int64  `json:"versionNum"`
	// This field is from variant [FunctionVersionGetResponseFunctionTransform].
	Audit           FunctionAudit       `json:"audit"`
	CreatedAt       time.Time           `json:"createdAt"`
	DisplayName     string              `json:"displayName"`
	Tags            []string            `json:"tags"`
	UsedInWorkflows []WorkflowUsageInfo `json:"usedInWorkflows"`
	// This field is from variant [FunctionVersionGetResponseFunctionAnalyze].
	EnableBoundingBoxes bool `json:"enableBoundingBoxes"`
	// This field is from variant [FunctionVersionGetResponseFunctionAnalyze].
	PreCount bool `json:"preCount"`
	// This field is from variant [FunctionVersionGetResponseFunctionClassify].
	Classifications []FunctionVersionGetResponseFunctionClassifyClassification `json:"classifications"`
	Description     string                                                     `json:"description"`
	// This field is from variant [FunctionVersionGetResponseFunctionSend].
	DestinationType string `json:"destinationType"`
	// This field is from variant [FunctionVersionGetResponseFunctionSend].
	GoogleDriveFolderID string `json:"googleDriveFolderId"`
	// This field is from variant [FunctionVersionGetResponseFunctionSend].
	S3Bucket string `json:"s3Bucket"`
	// This field is from variant [FunctionVersionGetResponseFunctionSend].
	S3Prefix string `json:"s3Prefix"`
	// This field is from variant [FunctionVersionGetResponseFunctionSend].
	WebhookSigningEnabled bool `json:"webhookSigningEnabled"`
	// This field is from variant [FunctionVersionGetResponseFunctionSend].
	WebhookURL string `json:"webhookUrl"`
	// This field is from variant [FunctionVersionGetResponseFunctionSplit].
	SplitType string `json:"splitType"`
	// This field is from variant [FunctionVersionGetResponseFunctionSplit].
	PrintPageSplitConfig FunctionVersionGetResponseFunctionSplitPrintPageSplitConfig `json:"printPageSplitConfig"`
	// This field is from variant [FunctionVersionGetResponseFunctionSplit].
	SemanticPageSplitConfig FunctionVersionGetResponseFunctionSplitSemanticPageSplitConfig `json:"semanticPageSplitConfig"`
	// This field is from variant [FunctionVersionGetResponseFunctionJoin].
	JoinType string `json:"joinType"`
	// This field is from variant [FunctionVersionGetResponseFunctionEnrich].
	Config EnrichConfig `json:"config"`
	// This field is from variant [FunctionVersionGetResponseFunctionPayloadShaping].
	ShapingSchema string `json:"shapingSchema"`
	JSON          struct {
		EmailAddress            respjson.Field
		FunctionID              respjson.Field
		FunctionName            respjson.Field
		OutputSchema            respjson.Field
		OutputSchemaName        respjson.Field
		TabularChunkingEnabled  respjson.Field
		Type                    respjson.Field
		VersionNum              respjson.Field
		Audit                   respjson.Field
		CreatedAt               respjson.Field
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
		Config                  respjson.Field
		ShapingSchema           respjson.Field
		raw                     string
	} `json:"-"`
}

// anyFunctionVersionGetResponseFunction is implemented by each variant of
// [FunctionVersionGetResponseFunctionUnion] to add type safety for the return type
// of [FunctionVersionGetResponseFunctionUnion.AsAny]
type anyFunctionVersionGetResponseFunction interface {
	implFunctionVersionGetResponseFunctionUnion()
}

func (FunctionVersionGetResponseFunctionTransform) implFunctionVersionGetResponseFunctionUnion() {}
func (FunctionVersionGetResponseFunctionExtract) implFunctionVersionGetResponseFunctionUnion()   {}
func (FunctionVersionGetResponseFunctionAnalyze) implFunctionVersionGetResponseFunctionUnion()   {}
func (FunctionVersionGetResponseFunctionClassify) implFunctionVersionGetResponseFunctionUnion()  {}
func (FunctionVersionGetResponseFunctionSend) implFunctionVersionGetResponseFunctionUnion()      {}
func (FunctionVersionGetResponseFunctionSplit) implFunctionVersionGetResponseFunctionUnion()     {}
func (FunctionVersionGetResponseFunctionJoin) implFunctionVersionGetResponseFunctionUnion()      {}
func (FunctionVersionGetResponseFunctionEnrich) implFunctionVersionGetResponseFunctionUnion()    {}
func (FunctionVersionGetResponseFunctionPayloadShaping) implFunctionVersionGetResponseFunctionUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := FunctionVersionGetResponseFunctionUnion.AsAny().(type) {
//	case bem.FunctionVersionGetResponseFunctionTransform:
//	case bem.FunctionVersionGetResponseFunctionExtract:
//	case bem.FunctionVersionGetResponseFunctionAnalyze:
//	case bem.FunctionVersionGetResponseFunctionClassify:
//	case bem.FunctionVersionGetResponseFunctionSend:
//	case bem.FunctionVersionGetResponseFunctionSplit:
//	case bem.FunctionVersionGetResponseFunctionJoin:
//	case bem.FunctionVersionGetResponseFunctionEnrich:
//	case bem.FunctionVersionGetResponseFunctionPayloadShaping:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u FunctionVersionGetResponseFunctionUnion) AsAny() anyFunctionVersionGetResponseFunction {
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
	case "enrich":
		return u.AsEnrich()
	case "payload_shaping":
		return u.AsPayloadShaping()
	}
	return nil
}

func (u FunctionVersionGetResponseFunctionUnion) AsTransform() (v FunctionVersionGetResponseFunctionTransform) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionVersionGetResponseFunctionUnion) AsExtract() (v FunctionVersionGetResponseFunctionExtract) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionVersionGetResponseFunctionUnion) AsAnalyze() (v FunctionVersionGetResponseFunctionAnalyze) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionVersionGetResponseFunctionUnion) AsClassify() (v FunctionVersionGetResponseFunctionClassify) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionVersionGetResponseFunctionUnion) AsSend() (v FunctionVersionGetResponseFunctionSend) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionVersionGetResponseFunctionUnion) AsSplit() (v FunctionVersionGetResponseFunctionSplit) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionVersionGetResponseFunctionUnion) AsJoin() (v FunctionVersionGetResponseFunctionJoin) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionVersionGetResponseFunctionUnion) AsEnrich() (v FunctionVersionGetResponseFunctionEnrich) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionVersionGetResponseFunctionUnion) AsPayloadShaping() (v FunctionVersionGetResponseFunctionPayloadShaping) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunctionVersionGetResponseFunctionUnion) RawJSON() string { return u.JSON.raw }

func (r *FunctionVersionGetResponseFunctionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionGetResponseFunctionTransform struct {
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
	// Audit trail information for the function version.
	Audit FunctionAudit `json:"audit"`
	// The date and time the function version was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
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
		CreatedAt              respjson.Field
		DisplayName            respjson.Field
		Tags                   respjson.Field
		UsedInWorkflows        respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionVersionGetResponseFunctionTransform) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionGetResponseFunctionTransform) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionGetResponseFunctionExtract struct {
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
	// Audit trail information for the function version.
	Audit FunctionAudit `json:"audit"`
	// The date and time the function version was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
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
		CreatedAt              respjson.Field
		DisplayName            respjson.Field
		Tags                   respjson.Field
		UsedInWorkflows        respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionVersionGetResponseFunctionExtract) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionGetResponseFunctionExtract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionGetResponseFunctionAnalyze struct {
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
	// Audit trail information for the function version.
	Audit FunctionAudit `json:"audit"`
	// The date and time the function version was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
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
		CreatedAt           respjson.Field
		DisplayName         respjson.Field
		Tags                respjson.Field
		UsedInWorkflows     respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionVersionGetResponseFunctionAnalyze) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionGetResponseFunctionAnalyze) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V3 read-side shape of a Classify (internally Route) function version. Mirrors {
type FunctionVersionGetResponseFunctionClassify struct {
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
	Classifications []FunctionVersionGetResponseFunctionClassifyClassification `json:"classifications" api:"required"`
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
	// Audit trail information for the function version.
	Audit FunctionAudit `json:"audit"`
	// The date and time the function version was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
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
		CreatedAt       respjson.Field
		DisplayName     respjson.Field
		Tags            respjson.Field
		UsedInWorkflows respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionVersionGetResponseFunctionClassify) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionGetResponseFunctionClassify) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionGetResponseFunctionClassifyClassification struct {
	Name            string                                                         `json:"name" api:"required"`
	Description     string                                                         `json:"description"`
	FunctionID      string                                                         `json:"functionID"`
	FunctionName    string                                                         `json:"functionName"`
	IsErrorFallback bool                                                           `json:"isErrorFallback"`
	Origin          FunctionVersionGetResponseFunctionClassifyClassificationOrigin `json:"origin"`
	Regex           FunctionVersionGetResponseFunctionClassifyClassificationRegex  `json:"regex"`
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
func (r FunctionVersionGetResponseFunctionClassifyClassification) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionGetResponseFunctionClassifyClassification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionGetResponseFunctionClassifyClassificationOrigin struct {
	Email FunctionVersionGetResponseFunctionClassifyClassificationOriginEmail `json:"email"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionVersionGetResponseFunctionClassifyClassificationOrigin) RawJSON() string {
	return r.JSON.raw
}
func (r *FunctionVersionGetResponseFunctionClassifyClassificationOrigin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionGetResponseFunctionClassifyClassificationOriginEmail struct {
	Patterns []string `json:"patterns"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Patterns    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionVersionGetResponseFunctionClassifyClassificationOriginEmail) RawJSON() string {
	return r.JSON.raw
}
func (r *FunctionVersionGetResponseFunctionClassifyClassificationOriginEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionGetResponseFunctionClassifyClassificationRegex struct {
	Patterns []string `json:"patterns"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Patterns    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionVersionGetResponseFunctionClassifyClassificationRegex) RawJSON() string {
	return r.JSON.raw
}
func (r *FunctionVersionGetResponseFunctionClassifyClassificationRegex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionGetResponseFunctionSend struct {
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
	// Audit trail information for the function version.
	Audit FunctionAudit `json:"audit"`
	// The date and time the function version was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName         string `json:"displayName"`
	GoogleDriveFolderID string `json:"googleDriveFolderId"`
	S3Bucket            string `json:"s3Bucket"`
	S3Prefix            string `json:"s3Prefix"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags"`
	// List of workflows that use this function.
	UsedInWorkflows []WorkflowUsageInfo `json:"usedInWorkflows"`
	// Whether webhook deliveries are signed with an HMAC-SHA256 `bem-signature`
	// header.
	WebhookSigningEnabled bool   `json:"webhookSigningEnabled"`
	WebhookURL            string `json:"webhookUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DestinationType       respjson.Field
		FunctionID            respjson.Field
		FunctionName          respjson.Field
		Type                  respjson.Field
		VersionNum            respjson.Field
		Audit                 respjson.Field
		CreatedAt             respjson.Field
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
func (r FunctionVersionGetResponseFunctionSend) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionGetResponseFunctionSend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionGetResponseFunctionSplit struct {
	// Unique identifier of function.
	FunctionID string `json:"functionID" api:"required"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string `json:"functionName" api:"required"`
	// Any of "print_page", "semantic_page".
	SplitType string         `json:"splitType" api:"required"`
	Type      constant.Split `json:"type" default:"split"`
	// Version number of function.
	VersionNum int64 `json:"versionNum" api:"required"`
	// Audit trail information for the function version.
	Audit FunctionAudit `json:"audit"`
	// The date and time the function version was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName             string                                                         `json:"displayName"`
	PrintPageSplitConfig    FunctionVersionGetResponseFunctionSplitPrintPageSplitConfig    `json:"printPageSplitConfig"`
	SemanticPageSplitConfig FunctionVersionGetResponseFunctionSplitSemanticPageSplitConfig `json:"semanticPageSplitConfig"`
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
		CreatedAt               respjson.Field
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
func (r FunctionVersionGetResponseFunctionSplit) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionGetResponseFunctionSplit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionGetResponseFunctionSplitPrintPageSplitConfig struct {
	NextFunctionID string `json:"nextFunctionID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextFunctionID respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionVersionGetResponseFunctionSplitPrintPageSplitConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *FunctionVersionGetResponseFunctionSplitPrintPageSplitConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionGetResponseFunctionSplitSemanticPageSplitConfig struct {
	ItemClasses []SplitFunctionSemanticPageItemClass `json:"itemClasses"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemClasses respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionVersionGetResponseFunctionSplitSemanticPageSplitConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *FunctionVersionGetResponseFunctionSplitSemanticPageSplitConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionGetResponseFunctionJoin struct {
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
	// Audit trail information for the function version.
	Audit FunctionAudit `json:"audit"`
	// The date and time the function version was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
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
		CreatedAt        respjson.Field
		DisplayName      respjson.Field
		Tags             respjson.Field
		UsedInWorkflows  respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionVersionGetResponseFunctionJoin) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionGetResponseFunctionJoin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionGetResponseFunctionEnrich struct {
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
	// Audit trail information for the function version.
	Audit FunctionAudit `json:"audit"`
	// The date and time the function version was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
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
		CreatedAt       respjson.Field
		DisplayName     respjson.Field
		Tags            respjson.Field
		UsedInWorkflows respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionVersionGetResponseFunctionEnrich) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionGetResponseFunctionEnrich) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A version of a payload shaping function that transforms and customizes input
// payloads using JMESPath expressions. Payload shaping allows you to extract
// specific data, perform calculations, and reshape complex input structures into
// simplified, standardized output formats tailored to your downstream systems or
// business requirements.
type FunctionVersionGetResponseFunctionPayloadShaping struct {
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
	// Audit trail information for the function version.
	Audit FunctionAudit `json:"audit"`
	// The date and time the function version was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
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
		CreatedAt       respjson.Field
		DisplayName     respjson.Field
		Tags            respjson.Field
		UsedInWorkflows respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionVersionGetResponseFunctionPayloadShaping) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionGetResponseFunctionPayloadShaping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionListResponse struct {
	// The total number of results available.
	TotalCount int64                                     `json:"totalCount"`
	Versions   []FunctionVersionListResponseVersionUnion `json:"versions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TotalCount  respjson.Field
		Versions    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionVersionListResponse) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunctionVersionListResponseVersionUnion contains all possible properties and
// values from [FunctionVersionListResponseVersionTransform],
// [FunctionVersionListResponseVersionExtract],
// [FunctionVersionListResponseVersionAnalyze],
// [FunctionVersionListResponseVersionClassify],
// [FunctionVersionListResponseVersionSend],
// [FunctionVersionListResponseVersionSplit],
// [FunctionVersionListResponseVersionJoin],
// [FunctionVersionListResponseVersionEnrich],
// [FunctionVersionListResponseVersionPayloadShaping].
//
// Use the [FunctionVersionListResponseVersionUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FunctionVersionListResponseVersionUnion struct {
	EmailAddress           string `json:"emailAddress"`
	FunctionID             string `json:"functionID"`
	FunctionName           string `json:"functionName"`
	OutputSchema           any    `json:"outputSchema"`
	OutputSchemaName       string `json:"outputSchemaName"`
	TabularChunkingEnabled bool   `json:"tabularChunkingEnabled"`
	// Any of "transform", "extract", "analyze", "classify", "send", "split", "join",
	// "enrich", "payload_shaping".
	Type       string `json:"type"`
	VersionNum int64  `json:"versionNum"`
	// This field is from variant [FunctionVersionListResponseVersionTransform].
	Audit           FunctionAudit       `json:"audit"`
	CreatedAt       time.Time           `json:"createdAt"`
	DisplayName     string              `json:"displayName"`
	Tags            []string            `json:"tags"`
	UsedInWorkflows []WorkflowUsageInfo `json:"usedInWorkflows"`
	// This field is from variant [FunctionVersionListResponseVersionAnalyze].
	EnableBoundingBoxes bool `json:"enableBoundingBoxes"`
	// This field is from variant [FunctionVersionListResponseVersionAnalyze].
	PreCount bool `json:"preCount"`
	// This field is from variant [FunctionVersionListResponseVersionClassify].
	Classifications []FunctionVersionListResponseVersionClassifyClassification `json:"classifications"`
	Description     string                                                     `json:"description"`
	// This field is from variant [FunctionVersionListResponseVersionSend].
	DestinationType string `json:"destinationType"`
	// This field is from variant [FunctionVersionListResponseVersionSend].
	GoogleDriveFolderID string `json:"googleDriveFolderId"`
	// This field is from variant [FunctionVersionListResponseVersionSend].
	S3Bucket string `json:"s3Bucket"`
	// This field is from variant [FunctionVersionListResponseVersionSend].
	S3Prefix string `json:"s3Prefix"`
	// This field is from variant [FunctionVersionListResponseVersionSend].
	WebhookSigningEnabled bool `json:"webhookSigningEnabled"`
	// This field is from variant [FunctionVersionListResponseVersionSend].
	WebhookURL string `json:"webhookUrl"`
	// This field is from variant [FunctionVersionListResponseVersionSplit].
	SplitType string `json:"splitType"`
	// This field is from variant [FunctionVersionListResponseVersionSplit].
	PrintPageSplitConfig FunctionVersionListResponseVersionSplitPrintPageSplitConfig `json:"printPageSplitConfig"`
	// This field is from variant [FunctionVersionListResponseVersionSplit].
	SemanticPageSplitConfig FunctionVersionListResponseVersionSplitSemanticPageSplitConfig `json:"semanticPageSplitConfig"`
	// This field is from variant [FunctionVersionListResponseVersionJoin].
	JoinType string `json:"joinType"`
	// This field is from variant [FunctionVersionListResponseVersionEnrich].
	Config EnrichConfig `json:"config"`
	// This field is from variant [FunctionVersionListResponseVersionPayloadShaping].
	ShapingSchema string `json:"shapingSchema"`
	JSON          struct {
		EmailAddress            respjson.Field
		FunctionID              respjson.Field
		FunctionName            respjson.Field
		OutputSchema            respjson.Field
		OutputSchemaName        respjson.Field
		TabularChunkingEnabled  respjson.Field
		Type                    respjson.Field
		VersionNum              respjson.Field
		Audit                   respjson.Field
		CreatedAt               respjson.Field
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
		Config                  respjson.Field
		ShapingSchema           respjson.Field
		raw                     string
	} `json:"-"`
}

// anyFunctionVersionListResponseVersion is implemented by each variant of
// [FunctionVersionListResponseVersionUnion] to add type safety for the return type
// of [FunctionVersionListResponseVersionUnion.AsAny]
type anyFunctionVersionListResponseVersion interface {
	implFunctionVersionListResponseVersionUnion()
}

func (FunctionVersionListResponseVersionTransform) implFunctionVersionListResponseVersionUnion() {}
func (FunctionVersionListResponseVersionExtract) implFunctionVersionListResponseVersionUnion()   {}
func (FunctionVersionListResponseVersionAnalyze) implFunctionVersionListResponseVersionUnion()   {}
func (FunctionVersionListResponseVersionClassify) implFunctionVersionListResponseVersionUnion()  {}
func (FunctionVersionListResponseVersionSend) implFunctionVersionListResponseVersionUnion()      {}
func (FunctionVersionListResponseVersionSplit) implFunctionVersionListResponseVersionUnion()     {}
func (FunctionVersionListResponseVersionJoin) implFunctionVersionListResponseVersionUnion()      {}
func (FunctionVersionListResponseVersionEnrich) implFunctionVersionListResponseVersionUnion()    {}
func (FunctionVersionListResponseVersionPayloadShaping) implFunctionVersionListResponseVersionUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := FunctionVersionListResponseVersionUnion.AsAny().(type) {
//	case bem.FunctionVersionListResponseVersionTransform:
//	case bem.FunctionVersionListResponseVersionExtract:
//	case bem.FunctionVersionListResponseVersionAnalyze:
//	case bem.FunctionVersionListResponseVersionClassify:
//	case bem.FunctionVersionListResponseVersionSend:
//	case bem.FunctionVersionListResponseVersionSplit:
//	case bem.FunctionVersionListResponseVersionJoin:
//	case bem.FunctionVersionListResponseVersionEnrich:
//	case bem.FunctionVersionListResponseVersionPayloadShaping:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u FunctionVersionListResponseVersionUnion) AsAny() anyFunctionVersionListResponseVersion {
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
	case "enrich":
		return u.AsEnrich()
	case "payload_shaping":
		return u.AsPayloadShaping()
	}
	return nil
}

func (u FunctionVersionListResponseVersionUnion) AsTransform() (v FunctionVersionListResponseVersionTransform) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionVersionListResponseVersionUnion) AsExtract() (v FunctionVersionListResponseVersionExtract) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionVersionListResponseVersionUnion) AsAnalyze() (v FunctionVersionListResponseVersionAnalyze) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionVersionListResponseVersionUnion) AsClassify() (v FunctionVersionListResponseVersionClassify) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionVersionListResponseVersionUnion) AsSend() (v FunctionVersionListResponseVersionSend) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionVersionListResponseVersionUnion) AsSplit() (v FunctionVersionListResponseVersionSplit) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionVersionListResponseVersionUnion) AsJoin() (v FunctionVersionListResponseVersionJoin) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionVersionListResponseVersionUnion) AsEnrich() (v FunctionVersionListResponseVersionEnrich) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionVersionListResponseVersionUnion) AsPayloadShaping() (v FunctionVersionListResponseVersionPayloadShaping) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunctionVersionListResponseVersionUnion) RawJSON() string { return u.JSON.raw }

func (r *FunctionVersionListResponseVersionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionListResponseVersionTransform struct {
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
	// Audit trail information for the function version.
	Audit FunctionAudit `json:"audit"`
	// The date and time the function version was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
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
		CreatedAt              respjson.Field
		DisplayName            respjson.Field
		Tags                   respjson.Field
		UsedInWorkflows        respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionVersionListResponseVersionTransform) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionListResponseVersionTransform) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionListResponseVersionExtract struct {
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
	// Audit trail information for the function version.
	Audit FunctionAudit `json:"audit"`
	// The date and time the function version was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
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
		CreatedAt              respjson.Field
		DisplayName            respjson.Field
		Tags                   respjson.Field
		UsedInWorkflows        respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionVersionListResponseVersionExtract) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionListResponseVersionExtract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionListResponseVersionAnalyze struct {
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
	// Audit trail information for the function version.
	Audit FunctionAudit `json:"audit"`
	// The date and time the function version was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
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
		CreatedAt           respjson.Field
		DisplayName         respjson.Field
		Tags                respjson.Field
		UsedInWorkflows     respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionVersionListResponseVersionAnalyze) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionListResponseVersionAnalyze) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V3 read-side shape of a Classify (internally Route) function version. Mirrors {
type FunctionVersionListResponseVersionClassify struct {
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
	Classifications []FunctionVersionListResponseVersionClassifyClassification `json:"classifications" api:"required"`
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
	// Audit trail information for the function version.
	Audit FunctionAudit `json:"audit"`
	// The date and time the function version was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
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
		CreatedAt       respjson.Field
		DisplayName     respjson.Field
		Tags            respjson.Field
		UsedInWorkflows respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionVersionListResponseVersionClassify) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionListResponseVersionClassify) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionListResponseVersionClassifyClassification struct {
	Name            string                                                         `json:"name" api:"required"`
	Description     string                                                         `json:"description"`
	FunctionID      string                                                         `json:"functionID"`
	FunctionName    string                                                         `json:"functionName"`
	IsErrorFallback bool                                                           `json:"isErrorFallback"`
	Origin          FunctionVersionListResponseVersionClassifyClassificationOrigin `json:"origin"`
	Regex           FunctionVersionListResponseVersionClassifyClassificationRegex  `json:"regex"`
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
func (r FunctionVersionListResponseVersionClassifyClassification) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionListResponseVersionClassifyClassification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionListResponseVersionClassifyClassificationOrigin struct {
	Email FunctionVersionListResponseVersionClassifyClassificationOriginEmail `json:"email"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionVersionListResponseVersionClassifyClassificationOrigin) RawJSON() string {
	return r.JSON.raw
}
func (r *FunctionVersionListResponseVersionClassifyClassificationOrigin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionListResponseVersionClassifyClassificationOriginEmail struct {
	Patterns []string `json:"patterns"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Patterns    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionVersionListResponseVersionClassifyClassificationOriginEmail) RawJSON() string {
	return r.JSON.raw
}
func (r *FunctionVersionListResponseVersionClassifyClassificationOriginEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionListResponseVersionClassifyClassificationRegex struct {
	Patterns []string `json:"patterns"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Patterns    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionVersionListResponseVersionClassifyClassificationRegex) RawJSON() string {
	return r.JSON.raw
}
func (r *FunctionVersionListResponseVersionClassifyClassificationRegex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionListResponseVersionSend struct {
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
	// Audit trail information for the function version.
	Audit FunctionAudit `json:"audit"`
	// The date and time the function version was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName         string `json:"displayName"`
	GoogleDriveFolderID string `json:"googleDriveFolderId"`
	S3Bucket            string `json:"s3Bucket"`
	S3Prefix            string `json:"s3Prefix"`
	// Array of tags to categorize and organize functions.
	Tags []string `json:"tags"`
	// List of workflows that use this function.
	UsedInWorkflows []WorkflowUsageInfo `json:"usedInWorkflows"`
	// Whether webhook deliveries are signed with an HMAC-SHA256 `bem-signature`
	// header.
	WebhookSigningEnabled bool   `json:"webhookSigningEnabled"`
	WebhookURL            string `json:"webhookUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DestinationType       respjson.Field
		FunctionID            respjson.Field
		FunctionName          respjson.Field
		Type                  respjson.Field
		VersionNum            respjson.Field
		Audit                 respjson.Field
		CreatedAt             respjson.Field
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
func (r FunctionVersionListResponseVersionSend) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionListResponseVersionSend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionListResponseVersionSplit struct {
	// Unique identifier of function.
	FunctionID string `json:"functionID" api:"required"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string `json:"functionName" api:"required"`
	// Any of "print_page", "semantic_page".
	SplitType string         `json:"splitType" api:"required"`
	Type      constant.Split `json:"type" default:"split"`
	// Version number of function.
	VersionNum int64 `json:"versionNum" api:"required"`
	// Audit trail information for the function version.
	Audit FunctionAudit `json:"audit"`
	// The date and time the function version was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName             string                                                         `json:"displayName"`
	PrintPageSplitConfig    FunctionVersionListResponseVersionSplitPrintPageSplitConfig    `json:"printPageSplitConfig"`
	SemanticPageSplitConfig FunctionVersionListResponseVersionSplitSemanticPageSplitConfig `json:"semanticPageSplitConfig"`
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
		CreatedAt               respjson.Field
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
func (r FunctionVersionListResponseVersionSplit) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionListResponseVersionSplit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionListResponseVersionSplitPrintPageSplitConfig struct {
	NextFunctionID string `json:"nextFunctionID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextFunctionID respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionVersionListResponseVersionSplitPrintPageSplitConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *FunctionVersionListResponseVersionSplitPrintPageSplitConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionListResponseVersionSplitSemanticPageSplitConfig struct {
	ItemClasses []SplitFunctionSemanticPageItemClass `json:"itemClasses"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemClasses respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionVersionListResponseVersionSplitSemanticPageSplitConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *FunctionVersionListResponseVersionSplitSemanticPageSplitConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionListResponseVersionJoin struct {
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
	// Audit trail information for the function version.
	Audit FunctionAudit `json:"audit"`
	// The date and time the function version was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
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
		CreatedAt        respjson.Field
		DisplayName      respjson.Field
		Tags             respjson.Field
		UsedInWorkflows  respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionVersionListResponseVersionJoin) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionListResponseVersionJoin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionListResponseVersionEnrich struct {
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
	// Audit trail information for the function version.
	Audit FunctionAudit `json:"audit"`
	// The date and time the function version was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
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
		CreatedAt       respjson.Field
		DisplayName     respjson.Field
		Tags            respjson.Field
		UsedInWorkflows respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionVersionListResponseVersionEnrich) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionListResponseVersionEnrich) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A version of a payload shaping function that transforms and customizes input
// payloads using JMESPath expressions. Payload shaping allows you to extract
// specific data, perform calculations, and reshape complex input structures into
// simplified, standardized output formats tailored to your downstream systems or
// business requirements.
type FunctionVersionListResponseVersionPayloadShaping struct {
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
	// Audit trail information for the function version.
	Audit FunctionAudit `json:"audit"`
	// The date and time the function version was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
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
		CreatedAt       respjson.Field
		DisplayName     respjson.Field
		Tags            respjson.Field
		UsedInWorkflows respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionVersionListResponseVersionPayloadShaping) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionListResponseVersionPayloadShaping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionGetParams struct {
	FunctionName string `path:"functionName" api:"required" json:"-"`
	paramObj
}
