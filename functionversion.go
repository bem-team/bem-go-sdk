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

// **Retrieve a specific historical version of a function.**
//
// Versions are immutable. Use this endpoint to inspect what a function looked like
// at the moment a particular call was made — every event and transformation
// records the function version it ran against.
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

// **List every version of a function.**
//
// Returns the full version history, newest-first. Each row captures the
// configuration the function had between updates. Useful for audits ("when did
// this schema change?") and for diffing two versions before promoting an update to
// production.
func (r *FunctionVersionService) List(ctx context.Context, functionName string, opts ...option.RequestOption) (res *ListFunctionVersionsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if functionName == "" {
		err = errors.New("missing required functionName parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/functions/%s/versions", url.PathEscape(functionName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// FunctionVersionUnion contains all possible properties and values from
// [FunctionVersionTransform], [FunctionVersionExtract], [FunctionVersionAnalyze],
// [FunctionVersionClassify], [FunctionVersionSend], [FunctionVersionSplit],
// [FunctionVersionJoin], [FunctionVersionEnrich], [FunctionVersionPayloadShaping],
// [FunctionVersionParse].
//
// Use the [FunctionVersionUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FunctionVersionUnion struct {
	EmailAddress           string `json:"emailAddress"`
	FunctionID             string `json:"functionID"`
	FunctionName           string `json:"functionName"`
	OutputSchema           any    `json:"outputSchema"`
	OutputSchemaName       string `json:"outputSchemaName"`
	TabularChunkingEnabled bool   `json:"tabularChunkingEnabled"`
	// Any of "transform", "extract", "analyze", "classify", "send", "split", "join",
	// "enrich", "payload_shaping", "parse".
	Type       string `json:"type"`
	VersionNum int64  `json:"versionNum"`
	// This field is from variant [FunctionVersionTransform].
	Audit               FunctionAudit       `json:"audit"`
	CreatedAt           time.Time           `json:"createdAt"`
	DisplayName         string              `json:"displayName"`
	Tags                []string            `json:"tags"`
	UsedInWorkflows     []WorkflowUsageInfo `json:"usedInWorkflows"`
	EnableBoundingBoxes bool                `json:"enableBoundingBoxes"`
	PreCount            bool                `json:"preCount"`
	// This field is from variant [FunctionVersionClassify].
	Classifications []ClassificationListItem `json:"classifications"`
	Description     string                   `json:"description"`
	// This field is from variant [FunctionVersionSend].
	DestinationType string `json:"destinationType"`
	// This field is from variant [FunctionVersionSend].
	GoogleDriveFolderID string `json:"googleDriveFolderId"`
	// This field is from variant [FunctionVersionSend].
	S3Bucket string `json:"s3Bucket"`
	// This field is from variant [FunctionVersionSend].
	S3Prefix string `json:"s3Prefix"`
	// This field is from variant [FunctionVersionSend].
	WebhookSigningEnabled bool `json:"webhookSigningEnabled"`
	// This field is from variant [FunctionVersionSend].
	WebhookURL string `json:"webhookUrl"`
	// This field is from variant [FunctionVersionSplit].
	SplitType string `json:"splitType"`
	// This field is from variant [FunctionVersionSplit].
	PrintPageSplitConfig FunctionVersionSplitPrintPageSplitConfig `json:"printPageSplitConfig"`
	// This field is from variant [FunctionVersionSplit].
	SemanticPageSplitConfig FunctionVersionSplitSemanticPageSplitConfig `json:"semanticPageSplitConfig"`
	// This field is from variant [FunctionVersionJoin].
	JoinType string `json:"joinType"`
	// This field is from variant [FunctionVersionEnrich].
	Config EnrichConfig `json:"config"`
	// This field is from variant [FunctionVersionPayloadShaping].
	ShapingSchema string `json:"shapingSchema"`
	// This field is from variant [FunctionVersionParse].
	ParseConfig FunctionVersionParseParseConfig `json:"parseConfig"`
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
		ParseConfig             respjson.Field
		raw                     string
	} `json:"-"`
}

// anyFunctionVersion is implemented by each variant of [FunctionVersionUnion] to
// add type safety for the return type of [FunctionVersionUnion.AsAny]
type anyFunctionVersion interface {
	implFunctionVersionUnion()
}

func (FunctionVersionTransform) implFunctionVersionUnion()      {}
func (FunctionVersionExtract) implFunctionVersionUnion()        {}
func (FunctionVersionAnalyze) implFunctionVersionUnion()        {}
func (FunctionVersionClassify) implFunctionVersionUnion()       {}
func (FunctionVersionSend) implFunctionVersionUnion()           {}
func (FunctionVersionSplit) implFunctionVersionUnion()          {}
func (FunctionVersionJoin) implFunctionVersionUnion()           {}
func (FunctionVersionEnrich) implFunctionVersionUnion()         {}
func (FunctionVersionPayloadShaping) implFunctionVersionUnion() {}
func (FunctionVersionParse) implFunctionVersionUnion()          {}

// Use the following switch statement to find the correct variant
//
//	switch variant := FunctionVersionUnion.AsAny().(type) {
//	case bem.FunctionVersionTransform:
//	case bem.FunctionVersionExtract:
//	case bem.FunctionVersionAnalyze:
//	case bem.FunctionVersionClassify:
//	case bem.FunctionVersionSend:
//	case bem.FunctionVersionSplit:
//	case bem.FunctionVersionJoin:
//	case bem.FunctionVersionEnrich:
//	case bem.FunctionVersionPayloadShaping:
//	case bem.FunctionVersionParse:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u FunctionVersionUnion) AsAny() anyFunctionVersion {
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
	case "parse":
		return u.AsParse()
	}
	return nil
}

func (u FunctionVersionUnion) AsTransform() (v FunctionVersionTransform) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionVersionUnion) AsExtract() (v FunctionVersionExtract) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionVersionUnion) AsAnalyze() (v FunctionVersionAnalyze) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionVersionUnion) AsClassify() (v FunctionVersionClassify) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionVersionUnion) AsSend() (v FunctionVersionSend) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionVersionUnion) AsSplit() (v FunctionVersionSplit) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionVersionUnion) AsJoin() (v FunctionVersionJoin) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionVersionUnion) AsEnrich() (v FunctionVersionEnrich) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionVersionUnion) AsPayloadShaping() (v FunctionVersionPayloadShaping) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunctionVersionUnion) AsParse() (v FunctionVersionParse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunctionVersionUnion) RawJSON() string { return u.JSON.raw }

func (r *FunctionVersionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionTransform struct {
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
func (r FunctionVersionTransform) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionTransform) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionExtract struct {
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
	PreCount bool `json:"preCount" api:"required"`
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
		EnableBoundingBoxes    respjson.Field
		FunctionID             respjson.Field
		FunctionName           respjson.Field
		OutputSchema           respjson.Field
		OutputSchemaName       respjson.Field
		PreCount               respjson.Field
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
func (r FunctionVersionExtract) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionExtract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionAnalyze struct {
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
func (r FunctionVersionAnalyze) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionAnalyze) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionClassify struct {
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
func (r FunctionVersionClassify) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionClassify) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionSend struct {
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
func (r FunctionVersionSend) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionSend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionSplit struct {
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
	DisplayName             string                                      `json:"displayName"`
	PrintPageSplitConfig    FunctionVersionSplitPrintPageSplitConfig    `json:"printPageSplitConfig"`
	SemanticPageSplitConfig FunctionVersionSplitSemanticPageSplitConfig `json:"semanticPageSplitConfig"`
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
func (r FunctionVersionSplit) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionSplit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionSplitPrintPageSplitConfig struct {
	NextFunctionID string `json:"nextFunctionID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextFunctionID respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionVersionSplitPrintPageSplitConfig) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionSplitPrintPageSplitConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionSplitSemanticPageSplitConfig struct {
	ItemClasses []SplitFunctionSemanticPageItemClass `json:"itemClasses"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemClasses respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionVersionSplitSemanticPageSplitConfig) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionSplitSemanticPageSplitConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionJoin struct {
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
func (r FunctionVersionJoin) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionJoin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionEnrich struct {
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
func (r FunctionVersionEnrich) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionEnrich) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A version of a payload shaping function that transforms and customizes input
// payloads using JMESPath expressions. Payload shaping allows you to extract
// specific data, perform calculations, and reshape complex input structures into
// simplified, standardized output formats tailored to your downstream systems or
// business requirements.
type FunctionVersionPayloadShaping struct {
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
func (r FunctionVersionPayloadShaping) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionPayloadShaping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionVersionParse struct {
	// Unique identifier of function.
	FunctionID string `json:"functionID" api:"required"`
	// Name of function. Must be UNIQUE on a per-environment basis.
	FunctionName string         `json:"functionName" api:"required"`
	Type         constant.Parse `json:"type" default:"parse"`
	// Version number of function.
	VersionNum int64 `json:"versionNum" api:"required"`
	// Audit trail information for the function version.
	Audit FunctionAudit `json:"audit"`
	// The date and time the function version was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Display name of function. Human-readable name to help you identify the function.
	DisplayName string `json:"displayName"`
	// Per-version configuration for a Parse function.
	//
	// Parse renders document pages (PDF, image) via vision LLM and emits structured
	// JSON. The two toggles below independently control entity extraction (a per-call
	// output concern) and cross-document memory linking (an environment-wide concern).
	ParseConfig FunctionVersionParseParseConfig `json:"parseConfig"`
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
		CreatedAt       respjson.Field
		DisplayName     respjson.Field
		ParseConfig     respjson.Field
		Tags            respjson.Field
		UsedInWorkflows respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionVersionParse) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionParse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-version configuration for a Parse function.
//
// Parse renders document pages (PDF, image) via vision LLM and emits structured
// JSON. The two toggles below independently control entity extraction (a per-call
// output concern) and cross-document memory linking (an environment-wide concern).
type FunctionVersionParseParseConfig struct {
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
func (r FunctionVersionParseParseConfig) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionParseParseConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListFunctionVersionsResponse struct {
	// The total number of results available.
	TotalCount int64                  `json:"totalCount"`
	Versions   []FunctionVersionUnion `json:"versions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TotalCount  respjson.Field
		Versions    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListFunctionVersionsResponse) RawJSON() string { return r.JSON.raw }
func (r *ListFunctionVersionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single-function-version response wrapper used by V3 endpoints.
type FunctionVersionGetResponse struct {
	// V3 read-side union for function versions. Same shape as the shared
	// `FunctionVersion` union but with `classify` in place of `route`.
	Function FunctionVersionUnion `json:"function" api:"required"`
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

type FunctionVersionGetParams struct {
	FunctionName string `path:"functionName" api:"required" json:"-"`
	paramObj
}
