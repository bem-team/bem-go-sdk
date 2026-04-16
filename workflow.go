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
)

// Workflows orchestrate one or more functions into a directed acyclic graph (DAG)
// for document processing.
//
// Use these endpoints to create, update, list, and manage workflows, and to invoke
// them with file input via `POST /v3/workflows/{workflowName}/call`.
//
// The call endpoint accepts files as either multipart form data or JSON with
// base64-encoded content. In the Bem CLI, use `@path/to/file` inside JSON values
// to automatically read and encode files:
//
// ```
//
//	bem workflows call --workflow-name my-workflow \
//	  --input.single-file '{"inputContent": "@file.pdf", "inputType": "pdf"}' \
//	  --wait
//
// ```
//
// WorkflowService contains methods and other services that help with interacting
// with the bem API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowService] method instead.
type WorkflowService struct {
	options []option.RequestOption
	// Workflows orchestrate one or more functions into a directed acyclic graph (DAG)
	// for document processing.
	//
	// Use these endpoints to create, update, list, and manage workflows, and to invoke
	// them with file input via `POST /v3/workflows/{workflowName}/call`.
	//
	// The call endpoint accepts files as either multipart form data or JSON with
	// base64-encoded content. In the Bem CLI, use `@path/to/file` inside JSON values
	// to automatically read and encode files:
	//
	// ```
	//
	//	bem workflows call --workflow-name my-workflow \
	//	  --input.single-file '{"inputContent": "@file.pdf", "inputType": "pdf"}' \
	//	  --wait
	//
	// ```
	Versions WorkflowVersionService
}

// NewWorkflowService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkflowService(opts ...option.RequestOption) (r WorkflowService) {
	r = WorkflowService{}
	r.options = opts
	r.Versions = NewWorkflowVersionService(opts...)
	return
}

// Create a Workflow
func (r *WorkflowService) New(ctx context.Context, body WorkflowNewParams, opts ...option.RequestOption) (res *WorkflowNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/workflows"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a Workflow
func (r *WorkflowService) Get(ctx context.Context, workflowName string, opts ...option.RequestOption) (res *WorkflowGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if workflowName == "" {
		err = errors.New("missing required workflowName parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/workflows/%s", url.PathEscape(workflowName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a Workflow
func (r *WorkflowService) Update(ctx context.Context, workflowName string, body WorkflowUpdateParams, opts ...option.RequestOption) (res *WorkflowUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if workflowName == "" {
		err = errors.New("missing required workflowName parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/workflows/%s", url.PathEscape(workflowName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List Workflows
func (r *WorkflowService) List(ctx context.Context, query WorkflowListParams, opts ...option.RequestOption) (res *pagination.WorkflowsPage[Workflow], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v3/workflows"
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

// List Workflows
func (r *WorkflowService) ListAutoPaging(ctx context.Context, query WorkflowListParams, opts ...option.RequestOption) *pagination.WorkflowsPageAutoPager[Workflow] {
	return pagination.NewWorkflowsPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a Workflow
func (r *WorkflowService) Delete(ctx context.Context, workflowName string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if workflowName == "" {
		err = errors.New("missing required workflowName parameter")
		return err
	}
	path := fmt.Sprintf("v3/workflows/%s", url.PathEscape(workflowName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// **Invoke a workflow.**
//
// Submit the input file as either a multipart form request or a JSON request with
// base64-encoded file content. The workflow name is derived from the URL path.
//
// ## Input Formats
//
//   - **Multipart form** (`multipart/form-data`): attach the file directly via the
//     `file` or `files` fields. Set `wait` in the form body to control synchronous
//     behaviour.
//   - **JSON** (`application/json`): base64-encode the file content and set it in
//     `input.singleFile.inputContent` or `input.batchFiles.inputs[*].inputContent`.
//     Pass `wait=true` as a query parameter to control synchronous behaviour.
//
// ## Synchronous vs Asynchronous
//
// By default the call is created asynchronously and this endpoint returns
// `202 Accepted` immediately with a `pending` call object. Set `wait` to `true` to
// block until the call completes (up to 30 seconds):
//
//   - On success: returns `200 OK` with the completed call, `outputs` populated
//   - On failure: returns `500 Internal Server Error` with the call and an `error`
//     message
//   - On timeout: returns `202 Accepted` with the still-running call
//
// ## Tracking
//
// Poll `GET /v3/calls/{callID}` to check status, or configure a webhook
// subscription to receive events when the call finishes.
//
// ## CLI Usage
//
// Use `@path/to/file` inside JSON string values to embed file contents
// automatically. Binary files (PDF, images, audio) are base64-encoded; text files
// are embedded as strings.
//
// Single file (synchronous):
//
// ```bash
//
//	bem workflows call \
//	  --workflow-name my-workflow \
//	  --input.single-file '{"inputContent": "@invoice.pdf", "inputType": "pdf"}' \
//	  --wait
//
// ```
//
// Single file (asynchronous, returns callID immediately):
//
// ```bash
//
//	bem workflows call \
//	  --workflow-name my-workflow \
//	  --input.single-file '{"inputContent": "@invoice.pdf", "inputType": "pdf"}'
//
// ```
//
// Batch files:
//
// ```bash
//
//	bem workflows call \
//	  --workflow-name my-workflow \
//	  --input.batch-files '{"inputs": [{"inputContent": "@a.pdf", "inputType": "pdf"}, {"inputContent": "@b.png", "inputType": "png"}]}'
//
// ```
//
// Alternative: pass the full `--input` flag as JSON:
//
// ```bash
//
//	bem workflows call \
//	  --workflow-name my-workflow \
//	  --input '{"singleFile": {"inputContent": "@invoice.pdf", "inputType": "pdf"}}' \
//	  --wait
//
// ```
//
// **Important:** `--wait` is a boolean flag. Use `--wait` or `--wait=true`. Do
// **not** use `--wait true` (with a space) — the `true` will be parsed as an
// unexpected positional argument.
//
// Supported `inputType` values: csv, docx, email, heic, heif, html, jpeg, json,
// m4a, mp3, pdf, png, text, wav, webp, xls, xlsx, xml.
func (r *WorkflowService) Call(ctx context.Context, workflowName string, params WorkflowCallParams, opts ...option.RequestOption) (res *CallGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if workflowName == "" {
		err = errors.New("missing required workflowName parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/workflows/%s/call", url.PathEscape(workflowName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Copy a Workflow
func (r *WorkflowService) Copy(ctx context.Context, body WorkflowCopyParams, opts ...option.RequestOption) (res *WorkflowCopyResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/workflows/copy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type FunctionVersionIdentifier struct {
	// Unique identifier of function. Provide either id or name, not both.
	ID string `json:"id"`
	// Name of function. Must be UNIQUE on a per-environment basis. Provide either id
	// or name, not both.
	Name string `json:"name"`
	// Version number of function.
	VersionNum int64 `json:"versionNum"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		VersionNum  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionVersionIdentifier) RawJSON() string { return r.JSON.raw }
func (r *FunctionVersionIdentifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FunctionVersionIdentifier to a
// FunctionVersionIdentifierParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FunctionVersionIdentifierParam.Overrides()
func (r FunctionVersionIdentifier) ToParam() FunctionVersionIdentifierParam {
	return param.Override[FunctionVersionIdentifierParam](json.RawMessage(r.RawJSON()))
}

type FunctionVersionIdentifierParam struct {
	// Unique identifier of function. Provide either id or name, not both.
	ID param.Opt[string] `json:"id,omitzero"`
	// Name of function. Must be UNIQUE on a per-environment basis. Provide either id
	// or name, not both.
	Name param.Opt[string] `json:"name,omitzero"`
	// Version number of function.
	VersionNum param.Opt[int64] `json:"versionNum,omitzero"`
	paramObj
}

func (r FunctionVersionIdentifierParam) MarshalJSON() (data []byte, err error) {
	type shadow FunctionVersionIdentifierParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionVersionIdentifierParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V3 read representation of a workflow version.
type Workflow struct {
	// Unique identifier of the workflow.
	ID string `json:"id" api:"required"`
	// The date and time the workflow was created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// All directed edges in this workflow version's DAG.
	Edges []WorkflowEdgeResponse `json:"edges" api:"required"`
	// Name of the entry-point call-site node.
	MainNodeName string `json:"mainNodeName" api:"required"`
	// Unique name of the workflow within the environment.
	Name string `json:"name" api:"required"`
	// All call-site nodes in this workflow version's DAG.
	Nodes []WorkflowNodeResponse `json:"nodes" api:"required"`
	// The date and time the workflow was last updated.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Version number of this workflow version.
	VersionNum int64 `json:"versionNum" api:"required"`
	// Audit trail information.
	Audit WorkflowAudit `json:"audit"`
	// Human-readable display name.
	DisplayName string `json:"displayName"`
	// Inbound email address associated with the workflow, if any.
	EmailAddress string `json:"emailAddress"`
	// Tags associated with the workflow.
	Tags []string `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Edges        respjson.Field
		MainNodeName respjson.Field
		Name         respjson.Field
		Nodes        respjson.Field
		UpdatedAt    respjson.Field
		VersionNum   respjson.Field
		Audit        respjson.Field
		DisplayName  respjson.Field
		EmailAddress respjson.Field
		Tags         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Workflow) RawJSON() string { return r.JSON.raw }
func (r *Workflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowAudit struct {
	// Information about who created the current version.
	VersionCreatedBy UserActionSummary `json:"versionCreatedBy"`
	// Information about who created the workflow.
	WorkflowCreatedBy UserActionSummary `json:"workflowCreatedBy"`
	// Information about who last updated the workflow.
	WorkflowLastUpdatedBy UserActionSummary `json:"workflowLastUpdatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		VersionCreatedBy      respjson.Field
		WorkflowCreatedBy     respjson.Field
		WorkflowLastUpdatedBy respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowAudit) RawJSON() string { return r.JSON.raw }
func (r *WorkflowAudit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Read representation of a directed edge between call-site nodes.
type WorkflowEdgeResponse struct {
	// Name of the destination node.
	DestinationNodeName string `json:"destinationNodeName" api:"required"`
	// Name of the source node.
	SourceNodeName string `json:"sourceNodeName" api:"required"`
	// Labelled outlet on the source node, if any.
	DestinationName string `json:"destinationName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DestinationNodeName respjson.Field
		SourceNodeName      respjson.Field
		DestinationName     respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowEdgeResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowEdgeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Read representation of a call-site node.
type WorkflowNodeResponse struct {
	// Function (and version) executing at this call site.
	Function FunctionVersionIdentifier `json:"function" api:"required"`
	// Name of this call site, unique within the workflow version.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Function    respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowNodeResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowNodeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowNewResponse struct {
	// Error message if the workflow creation failed.
	Error string `json:"error"`
	// V3 read representation of a workflow version.
	Workflow Workflow `json:"workflow"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		Workflow    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowGetResponse struct {
	// Error message if the workflow retrieval failed.
	Error string `json:"error"`
	// V3 read representation of a workflow version.
	Workflow Workflow `json:"workflow"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		Workflow    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowUpdateResponse struct {
	// Error message if the workflow update failed.
	Error string `json:"error"`
	// V3 read representation of a workflow version.
	Workflow Workflow `json:"workflow"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		Workflow    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowCopyResponse struct {
	// Functions that were copied when copying to a different environment. Empty when
	// copying within the same environment.
	CopiedFunctions []WorkflowCopyResponseCopiedFunction `json:"copiedFunctions"`
	// The environment the workflow was copied to.
	Environment string `json:"environment"`
	// Error message if the workflow copy failed.
	Error string `json:"error"`
	// V3 read representation of a workflow version.
	Workflow Workflow `json:"workflow"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CopiedFunctions respjson.Field
		Environment     respjson.Field
		Error           respjson.Field
		Workflow        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowCopyResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowCopyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowCopyResponseCopiedFunction struct {
	// ID of the source function that was copied.
	SourceFunctionID string `json:"sourceFunctionID" api:"required"`
	// Name of the source function that was copied.
	SourceFunctionName string `json:"sourceFunctionName" api:"required"`
	// Version number of the source function that was copied.
	SourceVersionNum int64 `json:"sourceVersionNum" api:"required"`
	// ID of the newly created function in the target environment.
	TargetFunctionID string `json:"targetFunctionID" api:"required"`
	// Name of the newly created function in the target environment.
	TargetFunctionName string `json:"targetFunctionName" api:"required"`
	// Version number of the newly created function in the target environment.
	TargetVersionNum int64 `json:"targetVersionNum" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SourceFunctionID   respjson.Field
		SourceFunctionName respjson.Field
		SourceVersionNum   respjson.Field
		TargetFunctionID   respjson.Field
		TargetFunctionName respjson.Field
		TargetVersionNum   respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowCopyResponseCopiedFunction) RawJSON() string { return r.JSON.raw }
func (r *WorkflowCopyResponseCopiedFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowNewParams struct {
	// Name of the entry-point node. Must not be a destination of any edge.
	MainNodeName string `json:"mainNodeName" api:"required"`
	// Unique name for the workflow. Must match `^[a-zA-Z0-9_-]{1,128}$`.
	Name string `json:"name" api:"required"`
	// Call-site nodes in the DAG. At least one is required.
	Nodes []WorkflowNewParamsNode `json:"nodes,omitzero" api:"required"`
	// Human-readable display name.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Directed edges between nodes. Omit or leave empty for single-node workflows.
	Edges []WorkflowNewParamsEdge `json:"edges,omitzero"`
	// Tags to categorize and organize the workflow.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r WorkflowNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single function call-site node in a workflow DAG.
//
// The property Function is required.
type WorkflowNewParamsNode struct {
	// The function (and version) to execute at this call site.
	Function FunctionVersionIdentifierParam `json:"function,omitzero" api:"required"`
	// Name for this call site. Must be unique within the workflow version. Defaults to
	// the function's own name when omitted.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r WorkflowNewParamsNode) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsNode
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A directed edge between two named call-site nodes.
//
// The properties DestinationNodeName, SourceNodeName are required.
type WorkflowNewParamsEdge struct {
	// Name of the destination node.
	DestinationNodeName string `json:"destinationNodeName" api:"required"`
	// Name of the source node.
	SourceNodeName string `json:"sourceNodeName" api:"required"`
	// Labelled outlet on the source node that activates this edge. Omit for the
	// default (unlabelled) outlet.
	DestinationName param.Opt[string] `json:"destinationName,omitzero"`
	paramObj
}

func (r WorkflowNewParamsEdge) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParamsEdge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParamsEdge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowUpdateParams struct {
	// Human-readable display name.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// `mainNodeName`, `nodes`, and `edges` must be provided together to update the DAG
	// topology. If none are provided the topology is copied unchanged from the current
	// version.
	MainNodeName param.Opt[string] `json:"mainNodeName,omitzero"`
	// New name for the workflow (renames it). Must match `^[a-zA-Z0-9_-]{1,128}$`.
	Name  param.Opt[string]          `json:"name,omitzero"`
	Edges []WorkflowUpdateParamsEdge `json:"edges,omitzero"`
	Nodes []WorkflowUpdateParamsNode `json:"nodes,omitzero"`
	// Tags to categorize and organize the workflow.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r WorkflowUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A directed edge between two named call-site nodes.
//
// The properties DestinationNodeName, SourceNodeName are required.
type WorkflowUpdateParamsEdge struct {
	// Name of the destination node.
	DestinationNodeName string `json:"destinationNodeName" api:"required"`
	// Name of the source node.
	SourceNodeName string `json:"sourceNodeName" api:"required"`
	// Labelled outlet on the source node that activates this edge. Omit for the
	// default (unlabelled) outlet.
	DestinationName param.Opt[string] `json:"destinationName,omitzero"`
	paramObj
}

func (r WorkflowUpdateParamsEdge) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpdateParamsEdge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowUpdateParamsEdge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single function call-site node in a workflow DAG.
//
// The property Function is required.
type WorkflowUpdateParamsNode struct {
	// The function (and version) to execute at this call site.
	Function FunctionVersionIdentifierParam `json:"function,omitzero" api:"required"`
	// Name for this call site. Must be unique within the workflow version. Defaults to
	// the function's own name when omitted.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r WorkflowUpdateParamsNode) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpdateParamsNode
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowUpdateParamsNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowListParams struct {
	DisplayName   param.Opt[string] `query:"displayName,omitzero" json:"-"`
	EndingBefore  param.Opt[string] `query:"endingBefore,omitzero" json:"-"`
	Limit         param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	StartingAfter param.Opt[string] `query:"startingAfter,omitzero" json:"-"`
	FunctionIDs   []string          `query:"functionIDs,omitzero" json:"-"`
	FunctionNames []string          `query:"functionNames,omitzero" json:"-"`
	// Any of "asc", "desc".
	SortOrder     WorkflowListParamsSortOrder `query:"sortOrder,omitzero" json:"-"`
	Tags          []string                    `query:"tags,omitzero" json:"-"`
	WorkflowIDs   []string                    `query:"workflowIDs,omitzero" json:"-"`
	WorkflowNames []string                    `query:"workflowNames,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkflowListParams]'s query parameters as `url.Values`.
func (r WorkflowListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkflowListParamsSortOrder string

const (
	WorkflowListParamsSortOrderAsc  WorkflowListParamsSortOrder = "asc"
	WorkflowListParamsSortOrderDesc WorkflowListParamsSortOrder = "desc"
)

type WorkflowCallParams struct {
	// Input file(s) for a call. Provide exactly one of `singleFile` or `batchFiles`.
	//
	// In the CLI, use the nested flags `--input.single-file` or `--input.batch-files`
	// with `@path/to/file` for automatic file embedding:
	// `--input.single-file '{"inputContent": "@invoice.pdf", "inputType": "pdf"}' --wait`
	Input WorkflowCallParamsInput `json:"input,omitzero" api:"required"`
	// Block until the call completes (up to 30 seconds) and return the finished call
	// object. Default: `false`. This is a boolean flag — use `--wait` or
	// `--wait=true`, not `--wait true`.
	Wait param.Opt[bool] `query:"wait,omitzero" json:"-"`
	// Your reference ID for tracking this call.
	CallReferenceID param.Opt[string] `json:"callReferenceID,omitzero"`
	// Arbitrary JSON object attached to this call. Stored on the call record and
	// injected into `transformedContent` under the reserved `_metadata` key (alongside
	// `referenceID`). Must be a JSON object. Maximum size: 4 KB.
	Metadata any `json:"metadata,omitzero"`
	paramObj
}

func (r WorkflowCallParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowCallParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowCallParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [WorkflowCallParams]'s query parameters as `url.Values`.
func (r WorkflowCallParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Input file(s) for a call. Provide exactly one of `singleFile` or `batchFiles`.
//
// In the CLI, use the nested flags `--input.single-file` or `--input.batch-files`
// with `@path/to/file` for automatic file embedding:
// `--input.single-file '{"inputContent": "@invoice.pdf", "inputType": "pdf"}' --wait`
type WorkflowCallParamsInput struct {
	// Multiple files to process in one call. Each item in the `inputs` array has its
	// own `inputContent` and `inputType`.
	BatchFiles WorkflowCallParamsInputBatchFiles `json:"batchFiles,omitzero"`
	// A single file input with base64-encoded content.
	//
	// When using the Bem CLI, use `@path/to/file` in the `inputContent` field to
	// automatically read and base64-encode the file:
	// `--input.single-file '{"inputContent": "@file.pdf", "inputType": "pdf"}' --wait`
	SingleFile WorkflowCallParamsInputSingleFile `json:"singleFile,omitzero"`
	paramObj
}

func (r WorkflowCallParamsInput) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowCallParamsInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowCallParamsInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Multiple files to process in one call. Each item in the `inputs` array has its
// own `inputContent` and `inputType`.
type WorkflowCallParamsInputBatchFiles struct {
	Inputs []WorkflowCallParamsInputBatchFilesInput `json:"inputs,omitzero"`
	paramObj
}

func (r WorkflowCallParamsInputBatchFiles) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowCallParamsInputBatchFiles
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowCallParamsInputBatchFiles) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties InputContent, InputType are required.
type WorkflowCallParamsInputBatchFilesInput struct {
	// Base64-encoded file content. In the Bem CLI, use `@path/to/file` to embed file
	// contents automatically.
	InputContent string `json:"inputContent" api:"required"`
	// The input type of the content you're sending for transformation.
	//
	// Any of "csv", "docx", "email", "heic", "html", "jpeg", "json", "heif", "m4a",
	// "mp3", "pdf", "png", "text", "wav", "webp", "xls", "xlsx", "xml".
	InputType       string            `json:"inputType,omitzero" api:"required"`
	ItemReferenceID param.Opt[string] `json:"itemReferenceID,omitzero"`
	paramObj
}

func (r WorkflowCallParamsInputBatchFilesInput) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowCallParamsInputBatchFilesInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowCallParamsInputBatchFilesInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowCallParamsInputBatchFilesInput](
		"inputType", "csv", "docx", "email", "heic", "html", "jpeg", "json", "heif", "m4a", "mp3", "pdf", "png", "text", "wav", "webp", "xls", "xlsx", "xml",
	)
}

// A single file input with base64-encoded content.
//
// When using the Bem CLI, use `@path/to/file` in the `inputContent` field to
// automatically read and base64-encode the file:
// `--input.single-file '{"inputContent": "@file.pdf", "inputType": "pdf"}' --wait`
//
// The properties InputContent, InputType are required.
type WorkflowCallParamsInputSingleFile struct {
	// Base64-encoded file content. In the Bem CLI, use `@path/to/file` to embed file
	// contents automatically.
	InputContent string `json:"inputContent" api:"required"`
	// The input type of the content you're sending for transformation.
	//
	// Any of "csv", "docx", "email", "heic", "html", "jpeg", "json", "heif", "m4a",
	// "mp3", "pdf", "png", "text", "wav", "webp", "xls", "xlsx", "xml".
	InputType string `json:"inputType,omitzero" api:"required"`
	paramObj
}

func (r WorkflowCallParamsInputSingleFile) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowCallParamsInputSingleFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowCallParamsInputSingleFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowCallParamsInputSingleFile](
		"inputType", "csv", "docx", "email", "heic", "html", "jpeg", "json", "heif", "m4a", "mp3", "pdf", "png", "text", "wav", "webp", "xls", "xlsx", "xml",
	)
}

type WorkflowCopyParams struct {
	// Name of the source workflow to copy from.
	SourceWorkflowName string `json:"sourceWorkflowName" api:"required"`
	// Name for the new copied workflow. Must be unique within the target environment.
	TargetWorkflowName string `json:"targetWorkflowName" api:"required"`
	// Optional version number of the source workflow to copy. If not provided, copies
	// the current version.
	SourceWorkflowVersionNum param.Opt[int64] `json:"sourceWorkflowVersionNum,omitzero"`
	// Optional display name for the copied workflow. If not provided, uses the source
	// workflow's display name with " (Copy)" appended.
	TargetDisplayName param.Opt[string] `json:"targetDisplayName,omitzero"`
	// Optional target environment name. If provided, copies the workflow to a
	// different environment. When copying to a different environment, all functions
	// used in the workflow will also be copied.
	TargetEnvironment param.Opt[string] `json:"targetEnvironment,omitzero"`
	// Optional tags for the copied workflow. If not provided, uses the source
	// workflow's tags.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r WorkflowCopyParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowCopyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowCopyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
