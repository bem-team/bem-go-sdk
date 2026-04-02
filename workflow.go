// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bem

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/bem-team/bem-go-sdk/internal/apiform"
	"github.com/bem-team/bem-go-sdk/internal/apijson"
	"github.com/bem-team/bem-go-sdk/internal/apiquery"
	"github.com/bem-team/bem-go-sdk/internal/requestconfig"
	"github.com/bem-team/bem-go-sdk/option"
	"github.com/bem-team/bem-go-sdk/packages/pagination"
	"github.com/bem-team/bem-go-sdk/packages/param"
	"github.com/bem-team/bem-go-sdk/packages/respjson"
)

// Workflow operations
//
// WorkflowService contains methods and other services that help with interacting
// with the bem API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowService] method instead.
type WorkflowService struct {
	options []option.RequestOption
	// Workflow operations
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

// **Invoke a workflow by submitting a multipart form request.**
//
// Workflows can only be called via multipart form in V3. Submit the input file
// along with an optional reference ID for tracking.
//
// ## Synchronous vs Asynchronous
//
// By default the call is created asynchronously and this endpoint returns
// `202 Accepted` immediately with a `pending` call object. Set the `wait` field to
// `true` to block until the call completes (up to 30 seconds):
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
func (r *WorkflowService) Call(ctx context.Context, workflowName string, body WorkflowCallParams, opts ...option.RequestOption) (res *CallGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if workflowName == "" {
		err = errors.New("missing required workflowName parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/workflows/%s/call", url.PathEscape(workflowName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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

type Workflow struct {
	// Unique identifier of workflow.
	ID           string                    `json:"id" api:"required"`
	MainFunction FunctionVersionIdentifier `json:"mainFunction" api:"required"`
	// Unique name of workflow. Must be UNIQUE on a per-environment basis.
	Name string `json:"name" api:"required"`
	// Version number of workflow version.
	VersionNum int64 `json:"versionNum" api:"required"`
	// Audit trail information for the workflow.
	Audit WorkflowAudit `json:"audit"`
	// The date and time the workflow was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Display name of workflow.
	DisplayName string `json:"displayName"`
	// Email address of workflow.
	EmailAddress  string                 `json:"emailAddress"`
	Relationships []WorkflowRelationship `json:"relationships"`
	// Array of tags to categorize and organize workflows.
	Tags []string `json:"tags"`
	// The date and time the workflow was last updated.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		MainFunction  respjson.Field
		Name          respjson.Field
		VersionNum    respjson.Field
		Audit         respjson.Field
		CreatedAt     respjson.Field
		DisplayName   respjson.Field
		EmailAddress  respjson.Field
		Relationships respjson.Field
		Tags          respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Workflow) RawJSON() string { return r.JSON.raw }
func (r *Workflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Audit trail information for the workflow.
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

type WorkflowRelationship struct {
	DestinationFunction FunctionVersionIdentifier `json:"destinationFunction" api:"required"`
	SourceFunction      FunctionVersionIdentifier `json:"sourceFunction" api:"required"`
	// Name of destination.
	DestinationName string `json:"destinationName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DestinationFunction respjson.Field
		SourceFunction      respjson.Field
		DestinationName     respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowRelationship) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRelationship) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DestinationFunction, SourceFunction are required.
type WorkflowRequestRelationshipParam struct {
	DestinationFunction FunctionVersionIdentifierParam `json:"destinationFunction,omitzero" api:"required"`
	SourceFunction      FunctionVersionIdentifierParam `json:"sourceFunction,omitzero" api:"required"`
	// Name of destination.
	DestinationName param.Opt[string] `json:"destinationName,omitzero"`
	paramObj
}

func (r WorkflowRequestRelationshipParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowRequestRelationshipParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowRequestRelationshipParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowNewResponse struct {
	// Error message if the workflow creation failed.
	Error    string   `json:"error"`
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
	Error    string   `json:"error"`
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
	Error    string   `json:"error"`
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
	// Information about functions that were copied when copying to a different
	// environment. Empty when copying within the same environment.
	CopiedFunctions []WorkflowCopyResponseCopiedFunction `json:"copiedFunctions"`
	// The environment where the workflow was copied to.
	Environment string `json:"environment"`
	// Error message if the workflow copy failed.
	Error string `json:"error"`
	// The newly created workflow.
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
	// Display name of workflow.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Name of workflow. Can be updated to rename the workflow. Must be unique within
	// the environment and match the pattern ^[a-zA-Z0-9_-]{1,128}$.
	Name param.Opt[string] `json:"name,omitzero"`
	// Main function for the workflow. The `mainFunction` and `relationships` fields
	// act as a unit and must be provided together, or neither provided.
	//
	//   - If `mainFunction` is provided without `relationships`, relationships will
	//     default to an empty array.
	//   - If `relationships` is provided, `mainFunction` must also be provided
	//     (validation error if missing).
	//   - If neither is provided, both mainFunction and relationships remain unchanged
	//     from the current workflow version.
	MainFunction FunctionVersionIdentifierParam `json:"mainFunction,omitzero"`
	// Relationships between functions in the workflow. The `mainFunction` and
	// `relationships` fields act as a unit and must be provided together, or neither
	// provided.
	//
	//   - If `relationships` is provided, `mainFunction` must also be provided
	//     (validation error if missing).
	//   - If `mainFunction` is provided without `relationships`, relationships will
	//     default to an empty array.
	//   - If neither is provided, both mainFunction and relationships remain unchanged
	//     from the current workflow version.
	Relationships []WorkflowRequestRelationshipParam `json:"relationships,omitzero"`
	// Array of tags to categorize and organize workflows.
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

type WorkflowUpdateParams struct {
	// Display name of workflow.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Name of workflow. Can be updated to rename the workflow. Must be unique within
	// the environment and match the pattern ^[a-zA-Z0-9_-]{1,128}$.
	Name param.Opt[string] `json:"name,omitzero"`
	// Main function for the workflow. The `mainFunction` and `relationships` fields
	// act as a unit and must be provided together, or neither provided.
	//
	//   - If `mainFunction` is provided without `relationships`, relationships will
	//     default to an empty array.
	//   - If `relationships` is provided, `mainFunction` must also be provided
	//     (validation error if missing).
	//   - If neither is provided, both mainFunction and relationships remain unchanged
	//     from the current workflow version.
	MainFunction FunctionVersionIdentifierParam `json:"mainFunction,omitzero"`
	// Relationships between functions in the workflow. The `mainFunction` and
	// `relationships` fields act as a unit and must be provided together, or neither
	// provided.
	//
	//   - If `relationships` is provided, `mainFunction` must also be provided
	//     (validation error if missing).
	//   - If `mainFunction` is provided without `relationships`, relationships will
	//     default to an empty array.
	//   - If neither is provided, both mainFunction and relationships remain unchanged
	//     from the current workflow version.
	Relationships []WorkflowRequestRelationshipParam `json:"relationships,omitzero"`
	// Array of tags to categorize and organize workflows.
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
	// Your reference ID for tracking this call.
	CallReferenceID param.Opt[string] `json:"callReferenceID,omitzero"`
	// When `true`, the endpoint blocks until the call completes (up to 30 seconds) and
	// returns the finished call object. Default: `false`.
	Wait param.Opt[string] `json:"wait,omitzero"`
	// Single input file (for transform, analyze, route, and split functions).
	File any `json:"file,omitzero"`
	// Multiple input files (for join functions).
	Files []any `json:"files,omitzero"`
	paramObj
}

func (r WorkflowCallParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
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
