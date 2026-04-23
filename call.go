// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bem

import (
	"context"
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

// The Calls API provides a unified interface for invoking both **Workflows** and
// **Functions**.
//
// Use this API when you want to:
//
// - Execute a complete workflow that chains multiple functions together
// - Call a single function directly without defining a workflow
// - Submit batch requests with multiple inputs in a single API call
// - Track execution status using call reference IDs
//
// **Key Difference**: Calls vs Function Calls
//
//   - **Calls API** (`/v3/calls`): High-level API for invoking workflows or
//     functions by name/ID. Supports batch processing and workflow orchestration.
//   - **Function Calls API** (`/v3/functions/{functionName}/call`): Direct function
//     invocation with function-type-specific arguments. Better for granular control
//     over individual function calls.
//
// CallService contains methods and other services that help with interacting with
// the bem API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCallService] method instead.
type CallService struct {
	options []option.RequestOption
}

// NewCallService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCallService(opts ...option.RequestOption) (r CallService) {
	r = CallService{}
	r.options = opts
	return
}

// **Retrieve a workflow call by ID.**
//
// Returns the full call object including status, workflow details, terminal
// outputs, and terminal errors. `outputs` and `errors` are both populated once the
// call finishes — they are not mutually exclusive (a partially-completed workflow
// may have both).
//
// ## Status
//
// | Status      | Description                                                 |
// | ----------- | ----------------------------------------------------------- |
// | `pending`   | Queued, not yet started                                     |
// | `running`   | Currently executing                                         |
// | `completed` | All enclosed function calls finished without errors         |
// | `failed`    | One or more enclosed function calls produced an error event |
//
// Poll this endpoint or configure a webhook subscription to detect completion.
func (r *CallService) Get(ctx context.Context, callID string, opts ...option.RequestOption) (res *CallGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if callID == "" {
		err = errors.New("missing required callID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/calls/%s", url.PathEscape(callID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// **List workflow calls with filtering and pagination.**
//
// Returns calls created via `POST /v3/workflows/{workflowName}/call`.
//
// ## Filtering
//
// - `callIDs`: Specific call identifiers
// - `referenceIDs`: Your custom reference IDs
// - `workflowIDs` / `workflowNames`: Filter by workflow
//
// ## Pagination
//
// Use `startingAfter` and `endingBefore` cursors with a default limit of 50.
func (r *CallService) List(ctx context.Context, query CallListParams, opts ...option.RequestOption) (res *pagination.CallsPage[Call], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v3/calls"
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

// **List workflow calls with filtering and pagination.**
//
// Returns calls created via `POST /v3/workflows/{workflowName}/call`.
//
// ## Filtering
//
// - `callIDs`: Specific call identifiers
// - `referenceIDs`: Your custom reference IDs
// - `workflowIDs` / `workflowNames`: Filter by workflow
//
// ## Pagination
//
// Use `startingAfter` and `endingBefore` cursors with a default limit of 50.
func (r *CallService) ListAutoPaging(ctx context.Context, query CallListParams, opts ...option.RequestOption) *pagination.CallsPageAutoPager[Call] {
	return pagination.NewCallsPageAutoPager(r.List(ctx, query, opts...))
}

// **Retrieve the full execution trace of a workflow call.**
//
// Returns all function calls and events emitted during the call as flat arrays.
// The DAG can be reconstructed using `FunctionCallResponseBase.sourceEventID` (the
// event that spawned each function call) and each event's `functionCallID` (the
// function call that emitted it).
//
// ## Graph structure
//
//   - A function call with no `sourceEventID` is the root.
//   - An event's `functionCallID` points to the function call that emitted it.
//   - A function call's `sourceEventID` points to the event that triggered it.
//   - `workflowNodeName` identifies the DAG node; `incomingDestinationName`
//     identifies the labelled outlet used to reach this call (absent for unlabelled
//     edges and root calls).
//
// The trace is available as soon as the call exists and grows as execution
// proceeds.
func (r *CallService) GetTrace(ctx context.Context, callID string, opts ...option.RequestOption) (res *CallGetTraceResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if callID == "" {
		err = errors.New("missing required callID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/calls/%s/trace", url.PathEscape(callID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// A workflow call returned by the V3 API.
//
// Compared to the V2 `Call` model:
//
//   - Terminal outputs are split into `outputs` (non-error events) and `errors`
//     (error events)
//   - `callType` and function-scoped fields are removed — V3 calls are always
//     workflow calls
//   - The deprecated `functionCalls` field is removed (use
//     `GET /v3/calls/{callID}/trace`)
//   - `url` and `traceUrl` hint fields are included for resource discovery
type Call struct {
	// Unique identifier of the call.
	CallID string `json:"callID" api:"required"`
	// The date and time the call was created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Terminal error events of this call. Workflow calls are not atomic — `errors` and
	// `outputs` may both be non-empty if some enclosed function calls succeeded and
	// others failed.
	//
	// Retrieve individual errors via `GET /v3/errors/{eventID}`.
	Errors []ErrorEvent `json:"errors" api:"required"`
	// Terminal non-error outputs of this call: primary events (non-split-collection)
	// that did not trigger any downstream function calls. Workflow calls are not
	// atomic — `outputs` and `errors` may both be non-empty if some enclosed function
	// calls succeeded and others failed.
	//
	// Each element is a polymorphic event object; inspect `eventType` to determine the
	// type. Retrieve individual outputs via `GET /v3/outputs/{eventID}`.
	Outputs []EventUnion `json:"outputs" api:"required"`
	// Hint URL for the full execution trace: `GET /v3/calls/{callID}/trace`.
	TraceURL string `json:"traceUrl" api:"required"`
	// Hint URL for retrieving this call: `GET /v3/calls/{callID}`.
	URL string `json:"url" api:"required"`
	// Your reference ID for this call, propagated from the original request.
	CallReferenceID string `json:"callReferenceID"`
	// The date and time the call finished. Only set once status is `completed` or
	// `failed`.
	FinishedAt time.Time `json:"finishedAt" format:"date-time"`
	// Input to the main function call.
	Input CallInput `json:"input"`
	// Status of call.
	//
	// Any of "pending", "running", "completed", "failed".
	Status CallStatus `json:"status"`
	// Unique identifier of the workflow.
	WorkflowID string `json:"workflowID"`
	// Name of the workflow.
	WorkflowName string `json:"workflowName"`
	// Version number of the workflow.
	WorkflowVersionNum int64 `json:"workflowVersionNum"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallID             respjson.Field
		CreatedAt          respjson.Field
		Errors             respjson.Field
		Outputs            respjson.Field
		TraceURL           respjson.Field
		URL                respjson.Field
		CallReferenceID    respjson.Field
		FinishedAt         respjson.Field
		Input              respjson.Field
		Status             respjson.Field
		WorkflowID         respjson.Field
		WorkflowName       respjson.Field
		WorkflowVersionNum respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Call) RawJSON() string { return r.JSON.raw }
func (r *Call) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Input to the main function call.
type CallInput struct {
	BatchFiles CallInputBatchFiles `json:"batchFiles"`
	SingleFile CallInputSingleFile `json:"singleFile"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BatchFiles  respjson.Field
		SingleFile  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallInput) RawJSON() string { return r.JSON.raw }
func (r *CallInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallInputBatchFiles struct {
	Inputs []CallInputBatchFilesInput `json:"inputs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Inputs      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallInputBatchFiles) RawJSON() string { return r.JSON.raw }
func (r *CallInputBatchFiles) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallInputBatchFilesInput struct {
	// Input type of the file
	InputType string `json:"inputType"`
	// Item reference ID
	ItemReferenceID string `json:"itemReferenceID"`
	// Presigned S3 URL for the file
	S3URL string `json:"s3URL"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InputType       respjson.Field
		ItemReferenceID respjson.Field
		S3URL           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallInputBatchFilesInput) RawJSON() string { return r.JSON.raw }
func (r *CallInputBatchFilesInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallInputSingleFile struct {
	// Input type of the file
	InputType string `json:"inputType"`
	// Presigned S3 URL for the file
	S3URL string `json:"s3URL"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InputType   respjson.Field
		S3URL       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallInputSingleFile) RawJSON() string { return r.JSON.raw }
func (r *CallInputSingleFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of call.
type CallStatus string

const (
	CallStatusPending   CallStatus = "pending"
	CallStatusRunning   CallStatus = "running"
	CallStatusCompleted CallStatus = "completed"
	CallStatusFailed    CallStatus = "failed"
)

type CallGetResponse struct {
	// A workflow call returned by the V3 API.
	//
	// Compared to the V2 `Call` model:
	//
	//   - Terminal outputs are split into `outputs` (non-error events) and `errors`
	//     (error events)
	//   - `callType` and function-scoped fields are removed — V3 calls are always
	//     workflow calls
	//   - The deprecated `functionCalls` field is removed (use
	//     `GET /v3/calls/{callID}/trace`)
	//   - `url` and `traceUrl` hint fields are included for resource discovery
	Call Call `json:"call"`
	// Error message if the call retrieval failed, or if the call itself failed when
	// using `wait=true`.
	Error string `json:"error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Call        respjson.Field
		Error       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallGetResponse) RawJSON() string { return r.JSON.raw }
func (r *CallGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from `GET /v3/calls/{callID}/trace`.
//
// Contains the full execution DAG as flat arrays of function calls and events.
// Reconstruct the graph using `FunctionCallResponseBase.sourceEventID` (the event
// that spawned each function call) and each event's `functionCallID` (the function
// call that emitted it).
type CallGetTraceResponse struct {
	// Error message if trace retrieval failed.
	Error string `json:"error"`
	// Full execution DAG of a call as flat arrays. Reconstruct the graph using
	// FunctionCallResponseBase.sourceEventID and each event's functionCallID.
	Trace CallGetTraceResponseTrace `json:"trace"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		Trace       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallGetTraceResponse) RawJSON() string { return r.JSON.raw }
func (r *CallGetTraceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full execution DAG of a call as flat arrays. Reconstruct the graph using
// FunctionCallResponseBase.sourceEventID and each event's functionCallID.
type CallGetTraceResponseTrace struct {
	// All events emitted within this call, polymorphic by eventType.
	Events []any `json:"events" api:"required"`
	// All function calls executed within this call.
	FunctionCalls []CallGetTraceResponseTraceFunctionCall `json:"functionCalls" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Events        respjson.Field
		FunctionCalls respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallGetTraceResponseTrace) RawJSON() string { return r.JSON.raw }
func (r *CallGetTraceResponseTrace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallGetTraceResponseTraceFunctionCall struct {
	// Unique identifier for this function call
	FunctionCallID string `json:"functionCallID" api:"required"`
	// ID of the function that was called
	FunctionID string `json:"functionID" api:"required"`
	// Name of the function that was called
	FunctionName string `json:"functionName" api:"required"`
	// User-provided reference ID for tracking
	ReferenceID string `json:"referenceID" api:"required"`
	// The date and time this function call started.
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// The status of the action.
	//
	// Any of "pending", "running", "completed", "failed".
	Status string `json:"status" api:"required"`
	// The type of the function.
	//
	// Any of "transform", "extract", "route", "send", "split", "join", "analyze",
	// "payload_shaping", "enrich".
	Type FunctionType `json:"type" api:"required"`
	// Array of activity steps for this function call
	Activity []CallGetTraceResponseTraceFunctionCallActivity `json:"activity"`
	// The date and time this function call finished. Absent while still running.
	FinishedAt time.Time `json:"finishedAt" format:"date-time"`
	// Version number of the function
	FunctionVersionNum int64 `json:"functionVersionNum"`
	// The labelled outlet on the upstream node that routed execution to this call.
	// Absent for root calls, unlabelled edges, and pre-migration rows.
	IncomingDestinationName string `json:"incomingDestinationName"`
	// Array of all file inputs with their S3 URLs
	Inputs []CallGetTraceResponseTraceFunctionCallInput `json:"inputs"`
	// Input type for single file input (set when there's exactly one file input)
	InputType string `json:"inputType"`
	// Presigned S3 URL for single file input (set when there's exactly one file input)
	S3URL string `json:"s3URL"`
	// ID of the event that spawned this function call (for DAG reconstruction). Nil
	// for the root function call.
	SourceEventID string `json:"sourceEventID"`
	// ID of the function call that spawned this function call (for DAG reconstruction)
	SourceFunctionCallID string `json:"sourceFunctionCallID"`
	// ID of the workflow call this function call belongs to (top-level execution
	// context)
	WorkflowCallID string `json:"workflowCallID"`
	// Name of the workflow DAG call-site node this function call is executing. Absent
	// for non-workflow calls and pre-migration rows.
	WorkflowNodeName string `json:"workflowNodeName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FunctionCallID          respjson.Field
		FunctionID              respjson.Field
		FunctionName            respjson.Field
		ReferenceID             respjson.Field
		StartedAt               respjson.Field
		Status                  respjson.Field
		Type                    respjson.Field
		Activity                respjson.Field
		FinishedAt              respjson.Field
		FunctionVersionNum      respjson.Field
		IncomingDestinationName respjson.Field
		Inputs                  respjson.Field
		InputType               respjson.Field
		S3URL                   respjson.Field
		SourceEventID           respjson.Field
		SourceFunctionCallID    respjson.Field
		WorkflowCallID          respjson.Field
		WorkflowNodeName        respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallGetTraceResponseTraceFunctionCall) RawJSON() string { return r.JSON.raw }
func (r *CallGetTraceResponseTraceFunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallGetTraceResponseTraceFunctionCallActivity struct {
	DisplayName string `json:"displayName"`
	// Any of "pending", "running", "completed", "failed".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayName respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallGetTraceResponseTraceFunctionCallActivity) RawJSON() string { return r.JSON.raw }
func (r *CallGetTraceResponseTraceFunctionCallActivity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallGetTraceResponseTraceFunctionCallInput struct {
	// Input type of the file
	InputType string `json:"inputType"`
	// Item reference ID for batch inputs
	ItemReferenceID string `json:"itemReferenceID"`
	// Presigned S3 URL for the file input
	S3URL string `json:"s3URL"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InputType       respjson.Field
		ItemReferenceID respjson.Field
		S3URL           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallGetTraceResponseTraceFunctionCallInput) RawJSON() string { return r.JSON.raw }
func (r *CallGetTraceResponseTraceFunctionCallInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallListParams struct {
	EndingBefore param.Opt[string] `query:"endingBefore,omitzero" json:"-"`
	Limit        param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	// Case-insensitive substring match against `callReferenceID`.
	ReferenceIDSubstring param.Opt[string] `query:"referenceIDSubstring,omitzero" json:"-"`
	StartingAfter        param.Opt[string] `query:"startingAfter,omitzero" json:"-"`
	CallIDs              []string          `query:"callIDs,omitzero" json:"-"`
	ReferenceIDs         []string          `query:"referenceIDs,omitzero" json:"-"`
	// Any of "asc", "desc".
	SortOrder CallListParamsSortOrder `query:"sortOrder,omitzero" json:"-"`
	// Filter by one or more statuses.
	//
	// Any of "pending", "running", "completed", "failed".
	Statuses      []string `query:"statuses,omitzero" json:"-"`
	WorkflowIDs   []string `query:"workflowIDs,omitzero" json:"-"`
	WorkflowNames []string `query:"workflowNames,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CallListParams]'s query parameters as `url.Values`.
func (r CallListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CallListParamsSortOrder string

const (
	CallListParamsSortOrderAsc  CallListParamsSortOrder = "asc"
	CallListParamsSortOrderDesc CallListParamsSortOrder = "desc"
)
