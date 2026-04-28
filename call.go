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
	Outputs []CallOutputUnion `json:"outputs" api:"required"`
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

// CallOutputUnion contains all possible properties and values from
// [CallOutputTransform], [CallOutputExtract], [CallOutputRoute],
// [CallOutputClassify], [CallOutputSplitCollection], [CallOutputSplitItem],
// [ErrorEvent], [CallOutputJoin], [CallOutputEnrich],
// [CallOutputCollectionProcessing], [CallOutputSend].
//
// Use the [CallOutputUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CallOutputUnion struct {
	EventID            string  `json:"eventID"`
	FunctionID         string  `json:"functionID"`
	FunctionName       string  `json:"functionName"`
	ItemCount          int64   `json:"itemCount"`
	ItemOffset         int64   `json:"itemOffset"`
	ReferenceID        string  `json:"referenceID"`
	TransformedContent any     `json:"transformedContent"`
	AvgConfidence      float64 `json:"avgConfidence"`
	CallID             string  `json:"callID"`
	// This field is a union of [CallOutputTransformCorrectedContentUnion],
	// [CallOutputExtractCorrectedContentUnion]
	CorrectedContent CallOutputUnionCorrectedContent `json:"correctedContent"`
	CreatedAt        time.Time                       `json:"createdAt"`
	// Any of "transform", "extract", "route", "classify", "split_collection",
	// "split_item", "error", "join", "enrich", "collection_processing", "send".
	EventType             string `json:"eventType"`
	FieldConfidences      any    `json:"fieldConfidences"`
	FunctionCallID        string `json:"functionCallID"`
	FunctionCallTryNumber int64  `json:"functionCallTryNumber"`
	FunctionVersionNum    int64  `json:"functionVersionNum"`
	// This field is from variant [CallOutputTransform].
	InboundEmail InboundEmailEvent `json:"inboundEmail"`
	// This field is a union of [[]CallOutputTransformInput],
	// [[]CallOutputExtractInput]
	Inputs            CallOutputUnionInputs `json:"inputs"`
	InputType         string                `json:"inputType"`
	InvalidProperties []string              `json:"invalidProperties"`
	// This field is from variant [CallOutputTransform].
	IsRegression bool `json:"isRegression"`
	// This field is from variant [CallOutputTransform].
	LastPublishErrorAt string `json:"lastPublishErrorAt"`
	// This field is a union of [CallOutputTransformMetadata],
	// [CallOutputExtractMetadata], [CallOutputRouteMetadata],
	// [CallOutputClassifyMetadata], [CallOutputSplitCollectionMetadata],
	// [CallOutputSplitItemMetadata], [ErrorEventMetadata], [CallOutputJoinMetadata],
	// [CallOutputEnrichMetadata], [CallOutputCollectionProcessingMetadata],
	// [CallOutputSendMetadata]
	Metadata CallOutputUnionMetadata `json:"metadata"`
	// This field is from variant [CallOutputTransform].
	Metrics CallOutputTransformMetrics `json:"metrics"`
	// This field is from variant [CallOutputTransform].
	OrderMatching bool `json:"orderMatching"`
	// This field is from variant [CallOutputTransform].
	PipelineID string `json:"pipelineID"`
	// This field is from variant [CallOutputTransform].
	PublishedAt        time.Time `json:"publishedAt"`
	S3URL              string    `json:"s3URL"`
	TransformationID   string    `json:"transformationID"`
	WorkflowID         string    `json:"workflowID"`
	WorkflowName       string    `json:"workflowName"`
	WorkflowVersionNum int64     `json:"workflowVersionNum"`
	Choice             string    `json:"choice"`
	OutputType         string    `json:"outputType"`
	// This field is a union of [CallOutputSplitCollectionPrintPageOutput],
	// [CallOutputSplitItemPrintPageOutput]
	PrintPageOutput CallOutputUnionPrintPageOutput `json:"printPageOutput"`
	// This field is a union of [CallOutputSplitCollectionSemanticPageOutput],
	// [CallOutputSplitItemSemanticPageOutput]
	SemanticPageOutput CallOutputUnionSemanticPageOutput `json:"semanticPageOutput"`
	// This field is from variant [ErrorEvent].
	Message string `json:"message"`
	// This field is from variant [CallOutputJoin].
	Items []CallOutputJoinItem `json:"items"`
	// This field is from variant [CallOutputJoin].
	JoinType string `json:"joinType"`
	// This field is from variant [CallOutputEnrich].
	EnrichedContent any `json:"enrichedContent"`
	// This field is from variant [CallOutputCollectionProcessing].
	CollectionID string `json:"collectionID"`
	// This field is from variant [CallOutputCollectionProcessing].
	CollectionName string `json:"collectionName"`
	// This field is from variant [CallOutputCollectionProcessing].
	Operation string `json:"operation"`
	// This field is from variant [CallOutputCollectionProcessing].
	ProcessedCount int64 `json:"processedCount"`
	// This field is from variant [CallOutputCollectionProcessing].
	Status string `json:"status"`
	// This field is from variant [CallOutputCollectionProcessing].
	CollectionItemIDs []string `json:"collectionItemIDs"`
	// This field is from variant [CallOutputCollectionProcessing].
	ErrorMessage string `json:"errorMessage"`
	// This field is from variant [CallOutputSend].
	DeliveryStatus string `json:"deliveryStatus"`
	// This field is from variant [CallOutputSend].
	DestinationType string `json:"destinationType"`
	// This field is from variant [CallOutputSend].
	DeliveredContent any `json:"deliveredContent"`
	// This field is from variant [CallOutputSend].
	GoogleDriveOutput CallOutputSendGoogleDriveOutput `json:"googleDriveOutput"`
	// This field is from variant [CallOutputSend].
	S3Output CallOutputSendS3Output `json:"s3Output"`
	// This field is from variant [CallOutputSend].
	WebhookOutput CallOutputSendWebhookOutput `json:"webhookOutput"`
	JSON          struct {
		EventID               respjson.Field
		FunctionID            respjson.Field
		FunctionName          respjson.Field
		ItemCount             respjson.Field
		ItemOffset            respjson.Field
		ReferenceID           respjson.Field
		TransformedContent    respjson.Field
		AvgConfidence         respjson.Field
		CallID                respjson.Field
		CorrectedContent      respjson.Field
		CreatedAt             respjson.Field
		EventType             respjson.Field
		FieldConfidences      respjson.Field
		FunctionCallID        respjson.Field
		FunctionCallTryNumber respjson.Field
		FunctionVersionNum    respjson.Field
		InboundEmail          respjson.Field
		Inputs                respjson.Field
		InputType             respjson.Field
		InvalidProperties     respjson.Field
		IsRegression          respjson.Field
		LastPublishErrorAt    respjson.Field
		Metadata              respjson.Field
		Metrics               respjson.Field
		OrderMatching         respjson.Field
		PipelineID            respjson.Field
		PublishedAt           respjson.Field
		S3URL                 respjson.Field
		TransformationID      respjson.Field
		WorkflowID            respjson.Field
		WorkflowName          respjson.Field
		WorkflowVersionNum    respjson.Field
		Choice                respjson.Field
		OutputType            respjson.Field
		PrintPageOutput       respjson.Field
		SemanticPageOutput    respjson.Field
		Message               respjson.Field
		Items                 respjson.Field
		JoinType              respjson.Field
		EnrichedContent       respjson.Field
		CollectionID          respjson.Field
		CollectionName        respjson.Field
		Operation             respjson.Field
		ProcessedCount        respjson.Field
		Status                respjson.Field
		CollectionItemIDs     respjson.Field
		ErrorMessage          respjson.Field
		DeliveryStatus        respjson.Field
		DestinationType       respjson.Field
		DeliveredContent      respjson.Field
		GoogleDriveOutput     respjson.Field
		S3Output              respjson.Field
		WebhookOutput         respjson.Field
		raw                   string
	} `json:"-"`
}

// anyCallOutput is implemented by each variant of [CallOutputUnion] to add type
// safety for the return type of [CallOutputUnion.AsAny]
type anyCallOutput interface {
	implCallOutputUnion()
}

func (CallOutputTransform) implCallOutputUnion()            {}
func (CallOutputExtract) implCallOutputUnion()              {}
func (CallOutputRoute) implCallOutputUnion()                {}
func (CallOutputClassify) implCallOutputUnion()             {}
func (CallOutputSplitCollection) implCallOutputUnion()      {}
func (CallOutputSplitItem) implCallOutputUnion()            {}
func (ErrorEvent) implCallOutputUnion()                     {}
func (CallOutputJoin) implCallOutputUnion()                 {}
func (CallOutputEnrich) implCallOutputUnion()               {}
func (CallOutputCollectionProcessing) implCallOutputUnion() {}
func (CallOutputSend) implCallOutputUnion()                 {}

// Use the following switch statement to find the correct variant
//
//	switch variant := CallOutputUnion.AsAny().(type) {
//	case bem.CallOutputTransform:
//	case bem.CallOutputExtract:
//	case bem.CallOutputRoute:
//	case bem.CallOutputClassify:
//	case bem.CallOutputSplitCollection:
//	case bem.CallOutputSplitItem:
//	case bem.ErrorEvent:
//	case bem.CallOutputJoin:
//	case bem.CallOutputEnrich:
//	case bem.CallOutputCollectionProcessing:
//	case bem.CallOutputSend:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u CallOutputUnion) AsAny() anyCallOutput {
	switch u.EventType {
	case "transform":
		return u.AsTransform()
	case "extract":
		return u.AsExtract()
	case "route":
		return u.AsRoute()
	case "classify":
		return u.AsClassify()
	case "split_collection":
		return u.AsSplitCollection()
	case "split_item":
		return u.AsSplitItem()
	case "error":
		return u.AsError()
	case "join":
		return u.AsJoin()
	case "enrich":
		return u.AsEnrich()
	case "collection_processing":
		return u.AsCollectionProcessing()
	case "send":
		return u.AsSend()
	}
	return nil
}

func (u CallOutputUnion) AsTransform() (v CallOutputTransform) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CallOutputUnion) AsExtract() (v CallOutputExtract) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CallOutputUnion) AsRoute() (v CallOutputRoute) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CallOutputUnion) AsClassify() (v CallOutputClassify) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CallOutputUnion) AsSplitCollection() (v CallOutputSplitCollection) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CallOutputUnion) AsSplitItem() (v CallOutputSplitItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CallOutputUnion) AsError() (v ErrorEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CallOutputUnion) AsJoin() (v CallOutputJoin) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CallOutputUnion) AsEnrich() (v CallOutputEnrich) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CallOutputUnion) AsCollectionProcessing() (v CallOutputCollectionProcessing) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CallOutputUnion) AsSend() (v CallOutputSend) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CallOutputUnion) RawJSON() string { return u.JSON.raw }

func (r *CallOutputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CallOutputUnionCorrectedContent is an implicit subunion of [CallOutputUnion].
// CallOutputUnionCorrectedContent provides convenient access to the sub-properties
// of the union.
//
// For type safety it is recommended to directly use a variant of the
// [CallOutputUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAnyArray OfString OfFloat OfBool]
type CallOutputUnionCorrectedContent struct {
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool           `json:",inline"`
	Output []AnyTypeUnion `json:"output"`
	JSON   struct {
		OfAnyArray respjson.Field
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		Output     respjson.Field
		raw        string
	} `json:"-"`
}

func (r *CallOutputUnionCorrectedContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CallOutputUnionInputs is an implicit subunion of [CallOutputUnion].
// CallOutputUnionInputs provides convenient access to the sub-properties of the
// union.
//
// For type safety it is recommended to directly use a variant of the
// [CallOutputUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfCallOutputTransformInputs OfCallOutputExtractInputs]
type CallOutputUnionInputs struct {
	// This field will be present if the value is a [[]CallOutputTransformInput]
	// instead of an object.
	OfCallOutputTransformInputs []CallOutputTransformInput `json:",inline"`
	// This field will be present if the value is a [[]CallOutputExtractInput] instead
	// of an object.
	OfCallOutputExtractInputs []CallOutputExtractInput `json:",inline"`
	JSON                      struct {
		OfCallOutputTransformInputs respjson.Field
		OfCallOutputExtractInputs   respjson.Field
		raw                         string
	} `json:"-"`
}

func (r *CallOutputUnionInputs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CallOutputUnionMetadata is an implicit subunion of [CallOutputUnion].
// CallOutputUnionMetadata provides convenient access to the sub-properties of the
// union.
//
// For type safety it is recommended to directly use a variant of the
// [CallOutputUnion].
type CallOutputUnionMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	JSON                           struct {
		DurationFunctionToEventSeconds respjson.Field
		raw                            string
	} `json:"-"`
}

func (r *CallOutputUnionMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CallOutputUnionPrintPageOutput is an implicit subunion of [CallOutputUnion].
// CallOutputUnionPrintPageOutput provides convenient access to the sub-properties
// of the union.
//
// For type safety it is recommended to directly use a variant of the
// [CallOutputUnion].
type CallOutputUnionPrintPageOutput struct {
	ItemCount int64 `json:"itemCount"`
	// This field is from variant [CallOutputSplitCollectionPrintPageOutput].
	Items []CallOutputSplitCollectionPrintPageOutputItem `json:"items"`
	// This field is from variant [CallOutputSplitItemPrintPageOutput].
	CollectionReferenceID string `json:"collectionReferenceID"`
	// This field is from variant [CallOutputSplitItemPrintPageOutput].
	ItemOffset int64 `json:"itemOffset"`
	// This field is from variant [CallOutputSplitItemPrintPageOutput].
	S3URL string `json:"s3URL"`
	JSON  struct {
		ItemCount             respjson.Field
		Items                 respjson.Field
		CollectionReferenceID respjson.Field
		ItemOffset            respjson.Field
		S3URL                 respjson.Field
		raw                   string
	} `json:"-"`
}

func (r *CallOutputUnionPrintPageOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CallOutputUnionSemanticPageOutput is an implicit subunion of [CallOutputUnion].
// CallOutputUnionSemanticPageOutput provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [CallOutputUnion].
type CallOutputUnionSemanticPageOutput struct {
	ItemCount int64 `json:"itemCount"`
	// This field is from variant [CallOutputSplitCollectionSemanticPageOutput].
	Items     []CallOutputSplitCollectionSemanticPageOutputItem `json:"items"`
	PageCount int64                                             `json:"pageCount"`
	// This field is from variant [CallOutputSplitItemSemanticPageOutput].
	CollectionReferenceID string `json:"collectionReferenceID"`
	// This field is from variant [CallOutputSplitItemSemanticPageOutput].
	ItemClass string `json:"itemClass"`
	// This field is from variant [CallOutputSplitItemSemanticPageOutput].
	ItemClassCount int64 `json:"itemClassCount"`
	// This field is from variant [CallOutputSplitItemSemanticPageOutput].
	ItemClassOffset int64 `json:"itemClassOffset"`
	// This field is from variant [CallOutputSplitItemSemanticPageOutput].
	ItemOffset int64 `json:"itemOffset"`
	// This field is from variant [CallOutputSplitItemSemanticPageOutput].
	PageEnd int64 `json:"pageEnd"`
	// This field is from variant [CallOutputSplitItemSemanticPageOutput].
	PageStart int64 `json:"pageStart"`
	// This field is from variant [CallOutputSplitItemSemanticPageOutput].
	S3URL string `json:"s3URL"`
	JSON  struct {
		ItemCount             respjson.Field
		Items                 respjson.Field
		PageCount             respjson.Field
		CollectionReferenceID respjson.Field
		ItemClass             respjson.Field
		ItemClassCount        respjson.Field
		ItemClassOffset       respjson.Field
		ItemOffset            respjson.Field
		PageEnd               respjson.Field
		PageStart             respjson.Field
		S3URL                 respjson.Field
		raw                   string
	} `json:"-"`
}

func (r *CallOutputUnionSemanticPageOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputTransform struct {
	// Unique ID generated by bem to identify the event.
	EventID string `json:"eventID" api:"required"`
	// Unique identifier of function that this event is associated with.
	FunctionID string `json:"functionID" api:"required"`
	// Unique name of function that this event is associated with.
	FunctionName string `json:"functionName" api:"required"`
	// The number of items that were transformed. Used for batch transformations to
	// indicate how many items were transformed.
	ItemCount int64 `json:"itemCount" api:"required"`
	// The offset of the first item that was transformed. Used for batch
	// transformations to indicate which item in the batch this event corresponds to.
	ItemOffset int64 `json:"itemOffset" api:"required"`
	// The unique ID you use internally to refer to this data point, propagated from
	// the original function input.
	ReferenceID string `json:"referenceID" api:"required"`
	// The transformed content of the input. The structure of this object is defined by
	// the function configuration.
	TransformedContent any `json:"transformedContent" api:"required"`
	// Average confidence score across all extracted fields, in the range [0, 1].
	AvgConfidence float64 `json:"avgConfidence" api:"nullable"`
	// Unique identifier of workflow call that this event is associated with.
	CallID string `json:"callID"`
	// Corrected feedback provided for fine-tuning purposes.
	CorrectedContent CallOutputTransformCorrectedContentUnion `json:"correctedContent" api:"nullable"`
	// Timestamp indicating when the event was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Any of "transform".
	EventType string `json:"eventType"`
	// Per-field confidence scores. A JSON object mapping RFC 6901 JSON Pointer paths
	// (e.g. `"/invoiceNumber"`) to float values in the range [0, 1] indicating the
	// model's confidence in each extracted field value.
	FieldConfidences any `json:"fieldConfidences"`
	// Unique identifier of function call that this event is associated with.
	FunctionCallID string `json:"functionCallID"`
	// The attempt number of the function call that created this event. 1 indexed.
	FunctionCallTryNumber int64 `json:"functionCallTryNumber"`
	// Version number of function that this event is associated with.
	FunctionVersionNum int64 `json:"functionVersionNum"`
	// The inbound email that triggered this event.
	InboundEmail InboundEmailEvent `json:"inboundEmail"`
	// Array of transformation inputs with their types and S3 URLs.
	Inputs []CallOutputTransformInput `json:"inputs" api:"nullable"`
	// The input type of the content you're sending for transformation.
	//
	// Any of "csv", "docx", "email", "heic", "html", "jpeg", "json", "heif", "m4a",
	// "mp3", "pdf", "png", "text", "wav", "webp", "xls", "xlsx", "xml".
	InputType string `json:"inputType"`
	// List of properties that were invalid in the input.
	InvalidProperties []string `json:"invalidProperties"`
	// Indicates whether this transformation was created as part of a regression test.
	IsRegression bool `json:"isRegression"`
	// Last timestamp indicating when the transform was published via webhook and
	// received a non-200 response. Set to `null` on a subsequent retry if the webhook
	// service receives a 200 response.
	LastPublishErrorAt string                      `json:"lastPublishErrorAt" api:"nullable"`
	Metadata           CallOutputTransformMetadata `json:"metadata"`
	// Accuracy, precision, recall, and F1 score when corrected JSON is provided.
	Metrics CallOutputTransformMetrics `json:"metrics" api:"nullable"`
	// Indicates whether array order matters when comparing corrected JSON with
	// extracted JSON.
	OrderMatching bool `json:"orderMatching"`
	// ID of pipeline that transformed the original input data.
	PipelineID string `json:"pipelineID"`
	// Timestamp indicating when the transform was published via webhook and received a
	// successful 200 response. Value is `null` if the transformation hasn't been sent.
	PublishedAt time.Time `json:"publishedAt" format:"date-time"`
	// Presigned S3 URL for the input content uploaded to S3.
	S3URL string `json:"s3URL" api:"nullable"`
	// Unique ID for each transformation output generated by bem following Segment's
	// KSUID conventions.
	TransformationID string `json:"transformationID"`
	// Unique identifier of workflow that this event is associated with.
	WorkflowID string `json:"workflowID"`
	// Name of workflow that this event is associated with.
	WorkflowName string `json:"workflowName"`
	// Version number of workflow that this event is associated with.
	WorkflowVersionNum int64 `json:"workflowVersionNum"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventID               respjson.Field
		FunctionID            respjson.Field
		FunctionName          respjson.Field
		ItemCount             respjson.Field
		ItemOffset            respjson.Field
		ReferenceID           respjson.Field
		TransformedContent    respjson.Field
		AvgConfidence         respjson.Field
		CallID                respjson.Field
		CorrectedContent      respjson.Field
		CreatedAt             respjson.Field
		EventType             respjson.Field
		FieldConfidences      respjson.Field
		FunctionCallID        respjson.Field
		FunctionCallTryNumber respjson.Field
		FunctionVersionNum    respjson.Field
		InboundEmail          respjson.Field
		Inputs                respjson.Field
		InputType             respjson.Field
		InvalidProperties     respjson.Field
		IsRegression          respjson.Field
		LastPublishErrorAt    respjson.Field
		Metadata              respjson.Field
		Metrics               respjson.Field
		OrderMatching         respjson.Field
		PipelineID            respjson.Field
		PublishedAt           respjson.Field
		S3URL                 respjson.Field
		TransformationID      respjson.Field
		WorkflowID            respjson.Field
		WorkflowName          respjson.Field
		WorkflowVersionNum    respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputTransform) RawJSON() string { return r.JSON.raw }
func (r *CallOutputTransform) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CallOutputTransformCorrectedContentUnion contains all possible properties and
// values from [CallOutputTransformCorrectedContentOutput], [[]any], [string],
// [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAnyArray OfString OfFloat OfBool]
type CallOutputTransformCorrectedContentUnion struct {
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field is from variant [CallOutputTransformCorrectedContentOutput].
	Output []AnyTypeUnion `json:"output"`
	JSON   struct {
		OfAnyArray respjson.Field
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		Output     respjson.Field
		raw        string
	} `json:"-"`
}

func (u CallOutputTransformCorrectedContentUnion) AsCallOutputTransformCorrectedContentOutput() (v CallOutputTransformCorrectedContentOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CallOutputTransformCorrectedContentUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CallOutputTransformCorrectedContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CallOutputTransformCorrectedContentUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CallOutputTransformCorrectedContentUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CallOutputTransformCorrectedContentUnion) RawJSON() string { return u.JSON.raw }

func (r *CallOutputTransformCorrectedContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputTransformCorrectedContentOutput struct {
	Output []AnyTypeUnion `json:"output"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputTransformCorrectedContentOutput) RawJSON() string { return r.JSON.raw }
func (r *CallOutputTransformCorrectedContentOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputTransformInput struct {
	InputContent     string `json:"inputContent" api:"nullable"`
	InputType        string `json:"inputType" api:"nullable"`
	JsonInputContent any    `json:"jsonInputContent" api:"nullable"`
	S3URL            string `json:"s3URL" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InputContent     respjson.Field
		InputType        respjson.Field
		JsonInputContent respjson.Field
		S3URL            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputTransformInput) RawJSON() string { return r.JSON.raw }
func (r *CallOutputTransformInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputTransformMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputTransformMetadata) RawJSON() string { return r.JSON.raw }
func (r *CallOutputTransformMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Accuracy, precision, recall, and F1 score when corrected JSON is provided.
type CallOutputTransformMetrics struct {
	Differences []CallOutputTransformMetricsDifference `json:"differences"`
	Metrics     CallOutputTransformMetricsMetrics      `json:"metrics"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Differences respjson.Field
		Metrics     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputTransformMetrics) RawJSON() string { return r.JSON.raw }
func (r *CallOutputTransformMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputTransformMetricsDifference struct {
	Category     string `json:"category"`
	CorrectedVal any    `json:"correctedVal"`
	ExtractedVal any    `json:"extractedVal"`
	JsonPointer  string `json:"jsonPointer"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Category     respjson.Field
		CorrectedVal respjson.Field
		ExtractedVal respjson.Field
		JsonPointer  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputTransformMetricsDifference) RawJSON() string { return r.JSON.raw }
func (r *CallOutputTransformMetricsDifference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputTransformMetricsMetrics struct {
	Accuracy  float64 `json:"accuracy"`
	F1Score   float64 `json:"f1Score"`
	Precision float64 `json:"precision"`
	Recall    float64 `json:"recall"`
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
func (r CallOutputTransformMetricsMetrics) RawJSON() string { return r.JSON.raw }
func (r *CallOutputTransformMetricsMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V3 event variants that do not exist in the shared `Event` union.
//
// `ExtractEvent` and `ClassifyEvent` are emitted only by V3-era function types
// (`extract` and `classify`). The shared `Event` union in
// `specs/events/models.tsp` predates these types and continues to describe V2 /
// V1-alpha responses verbatim; V3 response payloads add the new variants via the
// `EventV3` union below while keeping every shared variant intact for backward
// compatibility.
type CallOutputExtract struct {
	// Unique ID generated by bem to identify the event.
	EventID string `json:"eventID" api:"required"`
	// Unique identifier of function that this event is associated with.
	FunctionID string `json:"functionID" api:"required"`
	// Unique name of function that this event is associated with.
	FunctionName string `json:"functionName" api:"required"`
	// The number of items that were transformed. Used for batch transformations to
	// indicate how many items were transformed.
	ItemCount int64 `json:"itemCount" api:"required"`
	// The offset of the first item that was transformed. Used for batch
	// transformations to indicate which item in the batch this event corresponds to.
	ItemOffset int64 `json:"itemOffset" api:"required"`
	// The unique ID you use internally to refer to this data point, propagated from
	// the original function input.
	ReferenceID string `json:"referenceID" api:"required"`
	// The transformed content of the input. The structure of this object is defined by
	// the function configuration.
	TransformedContent any `json:"transformedContent" api:"required"`
	// Average confidence score across all extracted fields, in the range [0, 1].
	AvgConfidence float64 `json:"avgConfidence" api:"nullable"`
	// Unique identifier of workflow call that this event is associated with.
	CallID string `json:"callID"`
	// Corrected feedback provided for fine-tuning purposes.
	CorrectedContent CallOutputExtractCorrectedContentUnion `json:"correctedContent" api:"nullable"`
	// Timestamp indicating when the event was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Any of "extract".
	EventType string `json:"eventType"`
	// Per-field confidence scores. A JSON object mapping RFC 6901 JSON Pointer paths
	// (e.g. `"/invoiceNumber"`) to float values in the range [0, 1] indicating the
	// model's confidence in each extracted field value.
	FieldConfidences any `json:"fieldConfidences"`
	// Unique identifier of function call that this event is associated with.
	FunctionCallID string `json:"functionCallID"`
	// The attempt number of the function call that created this event. 1 indexed.
	FunctionCallTryNumber int64 `json:"functionCallTryNumber"`
	// Version number of function that this event is associated with.
	FunctionVersionNum int64 `json:"functionVersionNum"`
	// The inbound email that triggered this event.
	InboundEmail InboundEmailEvent `json:"inboundEmail"`
	// Array of transformation inputs with their types and S3 URLs.
	Inputs []CallOutputExtractInput `json:"inputs" api:"nullable"`
	// The input type of the content you're sending for transformation.
	//
	// Any of "csv", "docx", "email", "heic", "html", "jpeg", "json", "heif", "m4a",
	// "mp3", "pdf", "png", "text", "wav", "webp", "xls", "xlsx", "xml".
	InputType string `json:"inputType"`
	// List of properties that were invalid in the input.
	InvalidProperties []string                  `json:"invalidProperties"`
	Metadata          CallOutputExtractMetadata `json:"metadata"`
	// Presigned S3 URL for the input content uploaded to S3.
	S3URL string `json:"s3URL" api:"nullable"`
	// Unique ID for each transformation output generated by bem following Segment's
	// KSUID conventions.
	TransformationID string `json:"transformationID"`
	// Unique identifier of workflow that this event is associated with.
	WorkflowID string `json:"workflowID"`
	// Name of workflow that this event is associated with.
	WorkflowName string `json:"workflowName"`
	// Version number of workflow that this event is associated with.
	WorkflowVersionNum int64 `json:"workflowVersionNum"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventID               respjson.Field
		FunctionID            respjson.Field
		FunctionName          respjson.Field
		ItemCount             respjson.Field
		ItemOffset            respjson.Field
		ReferenceID           respjson.Field
		TransformedContent    respjson.Field
		AvgConfidence         respjson.Field
		CallID                respjson.Field
		CorrectedContent      respjson.Field
		CreatedAt             respjson.Field
		EventType             respjson.Field
		FieldConfidences      respjson.Field
		FunctionCallID        respjson.Field
		FunctionCallTryNumber respjson.Field
		FunctionVersionNum    respjson.Field
		InboundEmail          respjson.Field
		Inputs                respjson.Field
		InputType             respjson.Field
		InvalidProperties     respjson.Field
		Metadata              respjson.Field
		S3URL                 respjson.Field
		TransformationID      respjson.Field
		WorkflowID            respjson.Field
		WorkflowName          respjson.Field
		WorkflowVersionNum    respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputExtract) RawJSON() string { return r.JSON.raw }
func (r *CallOutputExtract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CallOutputExtractCorrectedContentUnion contains all possible properties and
// values from [CallOutputExtractCorrectedContentOutput], [[]any], [string],
// [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAnyArray OfString OfFloat OfBool]
type CallOutputExtractCorrectedContentUnion struct {
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field is from variant [CallOutputExtractCorrectedContentOutput].
	Output []AnyTypeUnion `json:"output"`
	JSON   struct {
		OfAnyArray respjson.Field
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		Output     respjson.Field
		raw        string
	} `json:"-"`
}

func (u CallOutputExtractCorrectedContentUnion) AsCallOutputExtractCorrectedContentOutput() (v CallOutputExtractCorrectedContentOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CallOutputExtractCorrectedContentUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CallOutputExtractCorrectedContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CallOutputExtractCorrectedContentUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CallOutputExtractCorrectedContentUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CallOutputExtractCorrectedContentUnion) RawJSON() string { return u.JSON.raw }

func (r *CallOutputExtractCorrectedContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputExtractCorrectedContentOutput struct {
	Output []AnyTypeUnion `json:"output"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputExtractCorrectedContentOutput) RawJSON() string { return r.JSON.raw }
func (r *CallOutputExtractCorrectedContentOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputExtractInput struct {
	InputContent     string `json:"inputContent" api:"nullable"`
	InputType        string `json:"inputType" api:"nullable"`
	JsonInputContent any    `json:"jsonInputContent" api:"nullable"`
	S3URL            string `json:"s3URL" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InputContent     respjson.Field
		InputType        respjson.Field
		JsonInputContent respjson.Field
		S3URL            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputExtractInput) RawJSON() string { return r.JSON.raw }
func (r *CallOutputExtractInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputExtractMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputExtractMetadata) RawJSON() string { return r.JSON.raw }
func (r *CallOutputExtractMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputRoute struct {
	// The choice made by the router function.
	Choice string `json:"choice" api:"required"`
	// Unique ID generated by bem to identify the event.
	EventID string `json:"eventID" api:"required"`
	// Unique identifier of function that this event is associated with.
	FunctionID string `json:"functionID" api:"required"`
	// Unique name of function that this event is associated with.
	FunctionName string `json:"functionName" api:"required"`
	// The unique ID you use internally to refer to this data point, propagated from
	// the original function input.
	ReferenceID string `json:"referenceID" api:"required"`
	// Unique identifier of workflow call that this event is associated with.
	CallID string `json:"callID"`
	// Timestamp indicating when the event was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Any of "route".
	EventType string `json:"eventType"`
	// Unique identifier of function call that this event is associated with.
	FunctionCallID string `json:"functionCallID"`
	// The attempt number of the function call that created this event. 1 indexed.
	FunctionCallTryNumber int64 `json:"functionCallTryNumber"`
	// Version number of function that this event is associated with.
	FunctionVersionNum int64 `json:"functionVersionNum"`
	// The inbound email that triggered this event.
	InboundEmail InboundEmailEvent       `json:"inboundEmail"`
	Metadata     CallOutputRouteMetadata `json:"metadata"`
	// The presigned S3 URL of the file that was routed.
	S3URL string `json:"s3URL"`
	// Unique identifier of workflow that this event is associated with.
	WorkflowID string `json:"workflowID"`
	// Name of workflow that this event is associated with.
	WorkflowName string `json:"workflowName"`
	// Version number of workflow that this event is associated with.
	WorkflowVersionNum int64 `json:"workflowVersionNum"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Choice                respjson.Field
		EventID               respjson.Field
		FunctionID            respjson.Field
		FunctionName          respjson.Field
		ReferenceID           respjson.Field
		CallID                respjson.Field
		CreatedAt             respjson.Field
		EventType             respjson.Field
		FunctionCallID        respjson.Field
		FunctionCallTryNumber respjson.Field
		FunctionVersionNum    respjson.Field
		InboundEmail          respjson.Field
		Metadata              respjson.Field
		S3URL                 respjson.Field
		WorkflowID            respjson.Field
		WorkflowName          respjson.Field
		WorkflowVersionNum    respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputRoute) RawJSON() string { return r.JSON.raw }
func (r *CallOutputRoute) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputRouteMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputRouteMetadata) RawJSON() string { return r.JSON.raw }
func (r *CallOutputRouteMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputClassify struct {
	// The classification chosen by the classify function.
	Choice string `json:"choice" api:"required"`
	// Unique ID generated by bem to identify the event.
	EventID string `json:"eventID" api:"required"`
	// Unique identifier of function that this event is associated with.
	FunctionID string `json:"functionID" api:"required"`
	// Unique name of function that this event is associated with.
	FunctionName string `json:"functionName" api:"required"`
	// The unique ID you use internally to refer to this data point, propagated from
	// the original function input.
	ReferenceID string `json:"referenceID" api:"required"`
	// Unique identifier of workflow call that this event is associated with.
	CallID string `json:"callID"`
	// Timestamp indicating when the event was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Any of "classify".
	EventType string `json:"eventType"`
	// Unique identifier of function call that this event is associated with.
	FunctionCallID string `json:"functionCallID"`
	// The attempt number of the function call that created this event. 1 indexed.
	FunctionCallTryNumber int64 `json:"functionCallTryNumber"`
	// Version number of function that this event is associated with.
	FunctionVersionNum int64 `json:"functionVersionNum"`
	// The inbound email that triggered this event.
	InboundEmail InboundEmailEvent          `json:"inboundEmail"`
	Metadata     CallOutputClassifyMetadata `json:"metadata"`
	// The presigned S3 URL of the file that was classified.
	S3URL string `json:"s3URL"`
	// Unique identifier of workflow that this event is associated with.
	WorkflowID string `json:"workflowID"`
	// Name of workflow that this event is associated with.
	WorkflowName string `json:"workflowName"`
	// Version number of workflow that this event is associated with.
	WorkflowVersionNum int64 `json:"workflowVersionNum"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Choice                respjson.Field
		EventID               respjson.Field
		FunctionID            respjson.Field
		FunctionName          respjson.Field
		ReferenceID           respjson.Field
		CallID                respjson.Field
		CreatedAt             respjson.Field
		EventType             respjson.Field
		FunctionCallID        respjson.Field
		FunctionCallTryNumber respjson.Field
		FunctionVersionNum    respjson.Field
		InboundEmail          respjson.Field
		Metadata              respjson.Field
		S3URL                 respjson.Field
		WorkflowID            respjson.Field
		WorkflowName          respjson.Field
		WorkflowVersionNum    respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputClassify) RawJSON() string { return r.JSON.raw }
func (r *CallOutputClassify) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputClassifyMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputClassifyMetadata) RawJSON() string { return r.JSON.raw }
func (r *CallOutputClassifyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputSplitCollection struct {
	// Unique ID generated by bem to identify the event.
	EventID string `json:"eventID" api:"required"`
	// Unique identifier of function that this event is associated with.
	FunctionID string `json:"functionID" api:"required"`
	// Unique name of function that this event is associated with.
	FunctionName string `json:"functionName" api:"required"`
	// Any of "print_page", "semantic_page".
	OutputType      string                                   `json:"outputType" api:"required"`
	PrintPageOutput CallOutputSplitCollectionPrintPageOutput `json:"printPageOutput" api:"required"`
	// The unique ID you use internally to refer to this data point, propagated from
	// the original function input.
	ReferenceID        string                                      `json:"referenceID" api:"required"`
	SemanticPageOutput CallOutputSplitCollectionSemanticPageOutput `json:"semanticPageOutput" api:"required"`
	// Unique identifier of workflow call that this event is associated with.
	CallID string `json:"callID"`
	// Timestamp indicating when the event was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Any of "split_collection".
	EventType string `json:"eventType"`
	// Unique identifier of function call that this event is associated with.
	FunctionCallID string `json:"functionCallID"`
	// The attempt number of the function call that created this event. 1 indexed.
	FunctionCallTryNumber int64 `json:"functionCallTryNumber"`
	// Version number of function that this event is associated with.
	FunctionVersionNum int64 `json:"functionVersionNum"`
	// The inbound email that triggered this event.
	InboundEmail InboundEmailEvent                 `json:"inboundEmail"`
	Metadata     CallOutputSplitCollectionMetadata `json:"metadata"`
	// Unique identifier of workflow that this event is associated with.
	WorkflowID string `json:"workflowID"`
	// Name of workflow that this event is associated with.
	WorkflowName string `json:"workflowName"`
	// Version number of workflow that this event is associated with.
	WorkflowVersionNum int64 `json:"workflowVersionNum"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventID               respjson.Field
		FunctionID            respjson.Field
		FunctionName          respjson.Field
		OutputType            respjson.Field
		PrintPageOutput       respjson.Field
		ReferenceID           respjson.Field
		SemanticPageOutput    respjson.Field
		CallID                respjson.Field
		CreatedAt             respjson.Field
		EventType             respjson.Field
		FunctionCallID        respjson.Field
		FunctionCallTryNumber respjson.Field
		FunctionVersionNum    respjson.Field
		InboundEmail          respjson.Field
		Metadata              respjson.Field
		WorkflowID            respjson.Field
		WorkflowName          respjson.Field
		WorkflowVersionNum    respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputSplitCollection) RawJSON() string { return r.JSON.raw }
func (r *CallOutputSplitCollection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputSplitCollectionPrintPageOutput struct {
	ItemCount int64                                          `json:"itemCount"`
	Items     []CallOutputSplitCollectionPrintPageOutputItem `json:"items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemCount   respjson.Field
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputSplitCollectionPrintPageOutput) RawJSON() string { return r.JSON.raw }
func (r *CallOutputSplitCollectionPrintPageOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputSplitCollectionPrintPageOutputItem struct {
	ItemOffset      int64  `json:"itemOffset"`
	ItemReferenceID string `json:"itemReferenceID"`
	S3URL           string `json:"s3URL"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemOffset      respjson.Field
		ItemReferenceID respjson.Field
		S3URL           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputSplitCollectionPrintPageOutputItem) RawJSON() string { return r.JSON.raw }
func (r *CallOutputSplitCollectionPrintPageOutputItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputSplitCollectionSemanticPageOutput struct {
	ItemCount int64                                             `json:"itemCount"`
	Items     []CallOutputSplitCollectionSemanticPageOutputItem `json:"items"`
	PageCount int64                                             `json:"pageCount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemCount   respjson.Field
		Items       respjson.Field
		PageCount   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputSplitCollectionSemanticPageOutput) RawJSON() string { return r.JSON.raw }
func (r *CallOutputSplitCollectionSemanticPageOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputSplitCollectionSemanticPageOutputItem struct {
	ItemClass       string `json:"itemClass"`
	ItemClassCount  int64  `json:"itemClassCount"`
	ItemClassOffset int64  `json:"itemClassOffset"`
	ItemOffset      int64  `json:"itemOffset"`
	ItemReferenceID string `json:"itemReferenceID"`
	PageEnd         int64  `json:"pageEnd"`
	PageStart       int64  `json:"pageStart"`
	S3URL           string `json:"s3URL"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemClass       respjson.Field
		ItemClassCount  respjson.Field
		ItemClassOffset respjson.Field
		ItemOffset      respjson.Field
		ItemReferenceID respjson.Field
		PageEnd         respjson.Field
		PageStart       respjson.Field
		S3URL           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputSplitCollectionSemanticPageOutputItem) RawJSON() string { return r.JSON.raw }
func (r *CallOutputSplitCollectionSemanticPageOutputItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputSplitCollectionMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputSplitCollectionMetadata) RawJSON() string { return r.JSON.raw }
func (r *CallOutputSplitCollectionMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputSplitItem struct {
	// Unique ID generated by bem to identify the event.
	EventID string `json:"eventID" api:"required"`
	// Unique identifier of function that this event is associated with.
	FunctionID string `json:"functionID" api:"required"`
	// Unique name of function that this event is associated with.
	FunctionName string `json:"functionName" api:"required"`
	// Any of "print_page", "semantic_page".
	OutputType string `json:"outputType" api:"required"`
	// The unique ID you use internally to refer to this data point, propagated from
	// the original function input.
	ReferenceID string `json:"referenceID" api:"required"`
	// Unique identifier of workflow call that this event is associated with.
	CallID string `json:"callID"`
	// Timestamp indicating when the event was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Any of "split_item".
	EventType string `json:"eventType"`
	// Unique identifier of function call that this event is associated with.
	FunctionCallID string `json:"functionCallID"`
	// The attempt number of the function call that created this event. 1 indexed.
	FunctionCallTryNumber int64 `json:"functionCallTryNumber"`
	// Version number of function that this event is associated with.
	FunctionVersionNum int64 `json:"functionVersionNum"`
	// The inbound email that triggered this event.
	InboundEmail       InboundEmailEvent                     `json:"inboundEmail"`
	Metadata           CallOutputSplitItemMetadata           `json:"metadata"`
	PrintPageOutput    CallOutputSplitItemPrintPageOutput    `json:"printPageOutput"`
	SemanticPageOutput CallOutputSplitItemSemanticPageOutput `json:"semanticPageOutput"`
	// Unique identifier of workflow that this event is associated with.
	WorkflowID string `json:"workflowID"`
	// Name of workflow that this event is associated with.
	WorkflowName string `json:"workflowName"`
	// Version number of workflow that this event is associated with.
	WorkflowVersionNum int64 `json:"workflowVersionNum"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventID               respjson.Field
		FunctionID            respjson.Field
		FunctionName          respjson.Field
		OutputType            respjson.Field
		ReferenceID           respjson.Field
		CallID                respjson.Field
		CreatedAt             respjson.Field
		EventType             respjson.Field
		FunctionCallID        respjson.Field
		FunctionCallTryNumber respjson.Field
		FunctionVersionNum    respjson.Field
		InboundEmail          respjson.Field
		Metadata              respjson.Field
		PrintPageOutput       respjson.Field
		SemanticPageOutput    respjson.Field
		WorkflowID            respjson.Field
		WorkflowName          respjson.Field
		WorkflowVersionNum    respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputSplitItem) RawJSON() string { return r.JSON.raw }
func (r *CallOutputSplitItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputSplitItemMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputSplitItemMetadata) RawJSON() string { return r.JSON.raw }
func (r *CallOutputSplitItemMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputSplitItemPrintPageOutput struct {
	CollectionReferenceID string `json:"collectionReferenceID"`
	ItemCount             int64  `json:"itemCount"`
	ItemOffset            int64  `json:"itemOffset"`
	S3URL                 string `json:"s3URL"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CollectionReferenceID respjson.Field
		ItemCount             respjson.Field
		ItemOffset            respjson.Field
		S3URL                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputSplitItemPrintPageOutput) RawJSON() string { return r.JSON.raw }
func (r *CallOutputSplitItemPrintPageOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputSplitItemSemanticPageOutput struct {
	CollectionReferenceID string `json:"collectionReferenceID"`
	ItemClass             string `json:"itemClass"`
	ItemClassCount        int64  `json:"itemClassCount"`
	ItemClassOffset       int64  `json:"itemClassOffset"`
	ItemCount             int64  `json:"itemCount"`
	ItemOffset            int64  `json:"itemOffset"`
	PageCount             int64  `json:"pageCount"`
	PageEnd               int64  `json:"pageEnd"`
	PageStart             int64  `json:"pageStart"`
	S3URL                 string `json:"s3URL"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CollectionReferenceID respjson.Field
		ItemClass             respjson.Field
		ItemClassCount        respjson.Field
		ItemClassOffset       respjson.Field
		ItemCount             respjson.Field
		ItemOffset            respjson.Field
		PageCount             respjson.Field
		PageEnd               respjson.Field
		PageStart             respjson.Field
		S3URL                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputSplitItemSemanticPageOutput) RawJSON() string { return r.JSON.raw }
func (r *CallOutputSplitItemSemanticPageOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputJoin struct {
	// Unique ID generated by bem to identify the event.
	EventID string `json:"eventID" api:"required"`
	// Unique identifier of function that this event is associated with.
	FunctionID string `json:"functionID" api:"required"`
	// Unique name of function that this event is associated with.
	FunctionName string `json:"functionName" api:"required"`
	// List of properties that were invalid in the input.
	InvalidProperties []string `json:"invalidProperties" api:"required"`
	// The items that were joined.
	Items []CallOutputJoinItem `json:"items" api:"required"`
	// The type of join that was performed.
	//
	// Any of "standard".
	JoinType string `json:"joinType" api:"required"`
	// The unique ID you use internally to refer to this data point, propagated from
	// the original function input.
	ReferenceID string `json:"referenceID" api:"required"`
	// The transformed content of the input. The structure of this object is defined by
	// the function configuration.
	TransformedContent any `json:"transformedContent" api:"required"`
	// Average confidence score across all extracted fields, in the range [0, 1].
	AvgConfidence float64 `json:"avgConfidence" api:"nullable"`
	// Unique identifier of workflow call that this event is associated with.
	CallID string `json:"callID"`
	// Timestamp indicating when the event was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Any of "join".
	EventType string `json:"eventType"`
	// Per-field confidence scores. A JSON object mapping RFC 6901 JSON Pointer paths
	// (e.g. `"/invoiceNumber"`) to float values in the range [0, 1] indicating the
	// model's confidence in each extracted field value.
	FieldConfidences any `json:"fieldConfidences"`
	// Unique identifier of function call that this event is associated with.
	FunctionCallID string `json:"functionCallID"`
	// The attempt number of the function call that created this event. 1 indexed.
	FunctionCallTryNumber int64 `json:"functionCallTryNumber"`
	// Version number of function that this event is associated with.
	FunctionVersionNum int64 `json:"functionVersionNum"`
	// The inbound email that triggered this event.
	InboundEmail InboundEmailEvent      `json:"inboundEmail"`
	Metadata     CallOutputJoinMetadata `json:"metadata"`
	// Unique ID for each transformation output generated by bem following Segment's
	// KSUID conventions.
	TransformationID string `json:"transformationID"`
	// Unique identifier of workflow that this event is associated with.
	WorkflowID string `json:"workflowID"`
	// Name of workflow that this event is associated with.
	WorkflowName string `json:"workflowName"`
	// Version number of workflow that this event is associated with.
	WorkflowVersionNum int64 `json:"workflowVersionNum"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventID               respjson.Field
		FunctionID            respjson.Field
		FunctionName          respjson.Field
		InvalidProperties     respjson.Field
		Items                 respjson.Field
		JoinType              respjson.Field
		ReferenceID           respjson.Field
		TransformedContent    respjson.Field
		AvgConfidence         respjson.Field
		CallID                respjson.Field
		CreatedAt             respjson.Field
		EventType             respjson.Field
		FieldConfidences      respjson.Field
		FunctionCallID        respjson.Field
		FunctionCallTryNumber respjson.Field
		FunctionVersionNum    respjson.Field
		InboundEmail          respjson.Field
		Metadata              respjson.Field
		TransformationID      respjson.Field
		WorkflowID            respjson.Field
		WorkflowName          respjson.Field
		WorkflowVersionNum    respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputJoin) RawJSON() string { return r.JSON.raw }
func (r *CallOutputJoin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputJoinItem struct {
	// The number of items that were transformed.
	ItemCount int64 `json:"itemCount" api:"required"`
	// The offset of the first item that was transformed. Used for batch
	// transformations to indicate which item in the batch this event corresponds to.
	ItemOffset int64 `json:"itemOffset" api:"required"`
	// The unique ID you use internally to refer to this data point.
	ItemReferenceID string `json:"itemReferenceID" api:"required"`
	// The presigned S3 URL of the file that was joined.
	S3URL string `json:"s3URL"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemCount       respjson.Field
		ItemOffset      respjson.Field
		ItemReferenceID respjson.Field
		S3URL           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputJoinItem) RawJSON() string { return r.JSON.raw }
func (r *CallOutputJoinItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputJoinMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputJoinMetadata) RawJSON() string { return r.JSON.raw }
func (r *CallOutputJoinMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputEnrich struct {
	// The enriched content produced by the enrich function. Contains the input data
	// augmented with results from semantic search against collections.
	EnrichedContent any `json:"enrichedContent" api:"required"`
	// Unique ID generated by bem to identify the event.
	EventID string `json:"eventID" api:"required"`
	// Unique identifier of function that this event is associated with.
	FunctionID string `json:"functionID" api:"required"`
	// Unique name of function that this event is associated with.
	FunctionName string `json:"functionName" api:"required"`
	// The unique ID you use internally to refer to this data point, propagated from
	// the original function input.
	ReferenceID string `json:"referenceID" api:"required"`
	// Unique identifier of workflow call that this event is associated with.
	CallID string `json:"callID"`
	// Timestamp indicating when the event was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Any of "enrich".
	EventType string `json:"eventType"`
	// Unique identifier of function call that this event is associated with.
	FunctionCallID string `json:"functionCallID"`
	// The attempt number of the function call that created this event. 1 indexed.
	FunctionCallTryNumber int64 `json:"functionCallTryNumber"`
	// Version number of function that this event is associated with.
	FunctionVersionNum int64 `json:"functionVersionNum"`
	// The inbound email that triggered this event.
	InboundEmail InboundEmailEvent        `json:"inboundEmail"`
	Metadata     CallOutputEnrichMetadata `json:"metadata"`
	// Unique identifier of workflow that this event is associated with.
	WorkflowID string `json:"workflowID"`
	// Name of workflow that this event is associated with.
	WorkflowName string `json:"workflowName"`
	// Version number of workflow that this event is associated with.
	WorkflowVersionNum int64 `json:"workflowVersionNum"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnrichedContent       respjson.Field
		EventID               respjson.Field
		FunctionID            respjson.Field
		FunctionName          respjson.Field
		ReferenceID           respjson.Field
		CallID                respjson.Field
		CreatedAt             respjson.Field
		EventType             respjson.Field
		FunctionCallID        respjson.Field
		FunctionCallTryNumber respjson.Field
		FunctionVersionNum    respjson.Field
		InboundEmail          respjson.Field
		Metadata              respjson.Field
		WorkflowID            respjson.Field
		WorkflowName          respjson.Field
		WorkflowVersionNum    respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputEnrich) RawJSON() string { return r.JSON.raw }
func (r *CallOutputEnrich) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputEnrichMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputEnrichMetadata) RawJSON() string { return r.JSON.raw }
func (r *CallOutputEnrichMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputCollectionProcessing struct {
	// Unique identifier of the collection.
	CollectionID string `json:"collectionID" api:"required"`
	// Name/path of the collection.
	CollectionName string `json:"collectionName" api:"required"`
	// Unique ID generated by bem to identify the event.
	EventID string `json:"eventID" api:"required"`
	// The operation performed (add or update).
	//
	// Any of "add", "update".
	Operation string `json:"operation" api:"required"`
	// Number of items successfully processed.
	ProcessedCount int64 `json:"processedCount" api:"required"`
	// The unique ID you use internally to refer to this data point, propagated from
	// the original function input.
	ReferenceID string `json:"referenceID" api:"required"`
	// Processing status (success or failed).
	//
	// Any of "success", "failed".
	Status string `json:"status" api:"required"`
	// Array of collection item KSUIDs that were added or updated.
	CollectionItemIDs []string `json:"collectionItemIDs"`
	// Timestamp indicating when the event was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Error message if processing failed.
	ErrorMessage string `json:"errorMessage"`
	// Any of "collection_processing".
	EventType string `json:"eventType"`
	// The attempt number of the function call that created this event. 1 indexed.
	FunctionCallTryNumber int64 `json:"functionCallTryNumber"`
	// The inbound email that triggered this event.
	InboundEmail InboundEmailEvent                      `json:"inboundEmail"`
	Metadata     CallOutputCollectionProcessingMetadata `json:"metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CollectionID          respjson.Field
		CollectionName        respjson.Field
		EventID               respjson.Field
		Operation             respjson.Field
		ProcessedCount        respjson.Field
		ReferenceID           respjson.Field
		Status                respjson.Field
		CollectionItemIDs     respjson.Field
		CreatedAt             respjson.Field
		ErrorMessage          respjson.Field
		EventType             respjson.Field
		FunctionCallTryNumber respjson.Field
		InboundEmail          respjson.Field
		Metadata              respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputCollectionProcessing) RawJSON() string { return r.JSON.raw }
func (r *CallOutputCollectionProcessing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputCollectionProcessingMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputCollectionProcessingMetadata) RawJSON() string { return r.JSON.raw }
func (r *CallOutputCollectionProcessingMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputSend struct {
	// Outcome of a Send function's delivery attempt.
	//
	// Any of "success", "skip".
	DeliveryStatus string `json:"deliveryStatus" api:"required"`
	// Destination type for a Send function.
	//
	// Any of "webhook", "s3", "google_drive".
	DestinationType string `json:"destinationType" api:"required"`
	// Unique ID generated by bem to identify the event.
	EventID string `json:"eventID" api:"required"`
	// Unique identifier of function that this event is associated with.
	FunctionID string `json:"functionID" api:"required"`
	// Unique name of function that this event is associated with.
	FunctionName string `json:"functionName" api:"required"`
	// The unique ID you use internally to refer to this data point, propagated from
	// the original function input.
	ReferenceID string `json:"referenceID" api:"required"`
	// Unique identifier of workflow call that this event is associated with.
	CallID string `json:"callID"`
	// Timestamp indicating when the event was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// The full protocol event JSON that was delivered — identical to what subscription
	// publish would deliver for the same event. For ad-hoc calls with a JSON file
	// input, contains the raw input JSON. For ad-hoc calls with a binary file input,
	// contains {"s3URL": "<presigned-url>"}.
	DeliveredContent any `json:"deliveredContent"`
	// Any of "send".
	EventType string `json:"eventType"`
	// Unique identifier of function call that this event is associated with.
	FunctionCallID string `json:"functionCallID"`
	// The attempt number of the function call that created this event. 1 indexed.
	FunctionCallTryNumber int64 `json:"functionCallTryNumber"`
	// Version number of function that this event is associated with.
	FunctionVersionNum int64 `json:"functionVersionNum"`
	// Metadata returned when a Send function delivers to Google Drive.
	GoogleDriveOutput CallOutputSendGoogleDriveOutput `json:"googleDriveOutput"`
	// The inbound email that triggered this event.
	InboundEmail InboundEmailEvent      `json:"inboundEmail"`
	Metadata     CallOutputSendMetadata `json:"metadata"`
	// Metadata returned when a Send function delivers to an S3 bucket.
	S3Output CallOutputSendS3Output `json:"s3Output"`
	// Metadata returned when a Send function delivers to a webhook.
	WebhookOutput CallOutputSendWebhookOutput `json:"webhookOutput"`
	// Unique identifier of workflow that this event is associated with.
	WorkflowID string `json:"workflowID"`
	// Name of workflow that this event is associated with.
	WorkflowName string `json:"workflowName"`
	// Version number of workflow that this event is associated with.
	WorkflowVersionNum int64 `json:"workflowVersionNum"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeliveryStatus        respjson.Field
		DestinationType       respjson.Field
		EventID               respjson.Field
		FunctionID            respjson.Field
		FunctionName          respjson.Field
		ReferenceID           respjson.Field
		CallID                respjson.Field
		CreatedAt             respjson.Field
		DeliveredContent      respjson.Field
		EventType             respjson.Field
		FunctionCallID        respjson.Field
		FunctionCallTryNumber respjson.Field
		FunctionVersionNum    respjson.Field
		GoogleDriveOutput     respjson.Field
		InboundEmail          respjson.Field
		Metadata              respjson.Field
		S3Output              respjson.Field
		WebhookOutput         respjson.Field
		WorkflowID            respjson.Field
		WorkflowName          respjson.Field
		WorkflowVersionNum    respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputSend) RawJSON() string { return r.JSON.raw }
func (r *CallOutputSend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata returned when a Send function delivers to Google Drive.
type CallOutputSendGoogleDriveOutput struct {
	// Name of the file created in Google Drive.
	FileName string `json:"fileName" api:"required"`
	// ID of the Google Drive folder the file was placed in.
	FolderID string `json:"folderID" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileName    respjson.Field
		FolderID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputSendGoogleDriveOutput) RawJSON() string { return r.JSON.raw }
func (r *CallOutputSendGoogleDriveOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallOutputSendMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputSendMetadata) RawJSON() string { return r.JSON.raw }
func (r *CallOutputSendMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata returned when a Send function delivers to an S3 bucket.
type CallOutputSendS3Output struct {
	// Name of the S3 bucket the payload was written to.
	BucketName string `json:"bucketName" api:"required"`
	// Object key under which the payload was stored.
	Key string `json:"key" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BucketName  respjson.Field
		Key         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputSendS3Output) RawJSON() string { return r.JSON.raw }
func (r *CallOutputSendS3Output) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata returned when a Send function delivers to a webhook.
type CallOutputSendWebhookOutput struct {
	// Raw HTTP response body returned by the webhook endpoint.
	HTTPResponseBody string `json:"httpResponseBody" api:"required"`
	// HTTP status code returned by the webhook endpoint.
	HTTPStatusCode int64 `json:"httpStatusCode" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HTTPResponseBody respjson.Field
		HTTPStatusCode   respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallOutputSendWebhookOutput) RawJSON() string { return r.JSON.raw }
func (r *CallOutputSendWebhookOutput) UnmarshalJSON(data []byte) error {
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
	// Any of "transform", "extract", "route", "classify", "send", "split", "join",
	// "analyze", "payload_shaping", "enrich", "parse".
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
