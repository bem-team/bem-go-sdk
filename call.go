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
