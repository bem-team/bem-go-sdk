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

// Retrieve terminal non-error output events from workflow calls.
//
// Outputs are events produced by successful terminal function steps — steps that
// completed without errors and did not spawn further downstream function calls. A
// single workflow call may produce multiple outputs (e.g. from a
// split-then-transform pipeline).
//
// Outputs and errors from the same call are not mutually exclusive: a
// partially-completed workflow may have both.
//
// Use `GET /v3/outputs` to list outputs across calls, or
// `GET /v3/outputs/{eventID}` to retrieve a specific output. To get outputs scoped
// to a single call, filter by `callIDs`.
//
// OutputService contains methods and other services that help with interacting
// with the bem API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOutputService] method instead.
type OutputService struct {
	options []option.RequestOption
}

// NewOutputService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOutputService(opts ...option.RequestOption) (r OutputService) {
	r = OutputService{}
	r.options = opts
	return
}

// **Retrieve a single output event by ID.**
//
// Fetches any non-error event by its `eventID`. Returns `404` if the event does
// not exist or if it is an error event (use `GET /v3/errors/{eventID}` for those).
func (r *OutputService) Get(ctx context.Context, eventID string, opts ...option.RequestOption) (res *OutputGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if eventID == "" {
		err = errors.New("missing required eventID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/outputs/%s", url.PathEscape(eventID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// **List terminal non-error output events.**
//
// Returns events that represent successful terminal outputs — primary events
// (non-split-collection) that did not trigger any downstream function calls. Error
// events are excluded; use `GET /v3/errors` to retrieve those.
//
// ## Intermediate Events
//
// By default, intermediate events (those that spawned a downstream function call
// in a multi-step workflow) are excluded. Pass `includeIntermediate=true` to
// include them.
//
// ## Filtering
//
// Filter by call, workflow, function, or reference ID. Multiple filters are ANDed
// together.
func (r *OutputService) List(ctx context.Context, query OutputListParams, opts ...option.RequestOption) (res *pagination.OutputsPage[EventUnion], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v3/outputs"
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

// **List terminal non-error output events.**
//
// Returns events that represent successful terminal outputs — primary events
// (non-split-collection) that did not trigger any downstream function calls. Error
// events are excluded; use `GET /v3/errors` to retrieve those.
//
// ## Intermediate Events
//
// By default, intermediate events (those that spawned a downstream function call
// in a multi-step workflow) are excluded. Pass `includeIntermediate=true` to
// include them.
//
// ## Filtering
//
// Filter by call, workflow, function, or reference ID. Multiple filters are ANDed
// together.
func (r *OutputService) ListAutoPaging(ctx context.Context, query OutputListParams, opts ...option.RequestOption) *pagination.OutputsPageAutoPager[EventUnion] {
	return pagination.NewOutputsPageAutoPager(r.List(ctx, query, opts...))
}

// AnyTypeUnion contains all possible properties and values from [[]any], [string],
// [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAnyArray OfString OfFloat OfBool]
type AnyTypeUnion struct {
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfAnyArray respjson.Field
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		raw        string
	} `json:"-"`
}

func (u AnyTypeUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AnyTypeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AnyTypeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AnyTypeUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AnyTypeUnion) RawJSON() string { return u.JSON.raw }

func (r *AnyTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventUnion contains all possible properties and values from [EventTransform],
// [EventExtract], [EventRoute], [EventClassify], [EventSplitCollection],
// [EventSplitItem], [ErrorEvent], [EventJoin], [EventEnrich],
// [EventCollectionProcessing], [EventSend].
//
// Use the [EventUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type EventUnion struct {
	EventID            string  `json:"eventID"`
	FunctionID         string  `json:"functionID"`
	FunctionName       string  `json:"functionName"`
	ItemCount          int64   `json:"itemCount"`
	ItemOffset         int64   `json:"itemOffset"`
	ReferenceID        string  `json:"referenceID"`
	TransformedContent any     `json:"transformedContent"`
	AvgConfidence      float64 `json:"avgConfidence"`
	CallID             string  `json:"callID"`
	// This field is a union of [EventTransformCorrectedContentUnion],
	// [EventExtractCorrectedContentUnion]
	CorrectedContent EventUnionCorrectedContent `json:"correctedContent"`
	CreatedAt        time.Time                  `json:"createdAt"`
	// Any of "transform", "extract", "route", "classify", "split_collection",
	// "split_item", "error", "join", "enrich", "collection_processing", "send".
	EventType             string `json:"eventType"`
	FieldConfidences      any    `json:"fieldConfidences"`
	FunctionCallID        string `json:"functionCallID"`
	FunctionCallTryNumber int64  `json:"functionCallTryNumber"`
	FunctionVersionNum    int64  `json:"functionVersionNum"`
	// This field is from variant [EventTransform].
	InboundEmail InboundEmailEvent `json:"inboundEmail"`
	// This field is a union of [[]EventTransformInput], [[]EventExtractInput]
	Inputs            EventUnionInputs `json:"inputs"`
	InputType         string           `json:"inputType"`
	InvalidProperties []string         `json:"invalidProperties"`
	// This field is from variant [EventTransform].
	IsRegression bool `json:"isRegression"`
	// This field is from variant [EventTransform].
	LastPublishErrorAt string `json:"lastPublishErrorAt"`
	// This field is a union of [EventTransformMetadata], [EventExtractMetadata],
	// [EventRouteMetadata], [EventClassifyMetadata], [EventSplitCollectionMetadata],
	// [EventSplitItemMetadata], [ErrorEventMetadata], [EventJoinMetadata],
	// [EventEnrichMetadata], [EventCollectionProcessingMetadata], [EventSendMetadata]
	Metadata EventUnionMetadata `json:"metadata"`
	// This field is from variant [EventTransform].
	Metrics EventTransformMetrics `json:"metrics"`
	// This field is from variant [EventTransform].
	OrderMatching bool `json:"orderMatching"`
	// This field is from variant [EventTransform].
	PipelineID string `json:"pipelineID"`
	// This field is from variant [EventTransform].
	PublishedAt        time.Time `json:"publishedAt"`
	S3URL              string    `json:"s3URL"`
	TransformationID   string    `json:"transformationID"`
	WorkflowID         string    `json:"workflowID"`
	WorkflowName       string    `json:"workflowName"`
	WorkflowVersionNum int64     `json:"workflowVersionNum"`
	Choice             string    `json:"choice"`
	OutputType         string    `json:"outputType"`
	// This field is a union of [EventSplitCollectionPrintPageOutput],
	// [EventSplitItemPrintPageOutput]
	PrintPageOutput EventUnionPrintPageOutput `json:"printPageOutput"`
	// This field is a union of [EventSplitCollectionSemanticPageOutput],
	// [EventSplitItemSemanticPageOutput]
	SemanticPageOutput EventUnionSemanticPageOutput `json:"semanticPageOutput"`
	// This field is from variant [ErrorEvent].
	Message string `json:"message"`
	// This field is from variant [EventJoin].
	Items []EventJoinItem `json:"items"`
	// This field is from variant [EventJoin].
	JoinType string `json:"joinType"`
	// This field is from variant [EventEnrich].
	EnrichedContent any `json:"enrichedContent"`
	// This field is from variant [EventCollectionProcessing].
	CollectionID string `json:"collectionID"`
	// This field is from variant [EventCollectionProcessing].
	CollectionName string `json:"collectionName"`
	// This field is from variant [EventCollectionProcessing].
	Operation string `json:"operation"`
	// This field is from variant [EventCollectionProcessing].
	ProcessedCount int64 `json:"processedCount"`
	// This field is from variant [EventCollectionProcessing].
	Status string `json:"status"`
	// This field is from variant [EventCollectionProcessing].
	CollectionItemIDs []string `json:"collectionItemIDs"`
	// This field is from variant [EventCollectionProcessing].
	ErrorMessage string `json:"errorMessage"`
	// This field is from variant [EventSend].
	DeliveryStatus string `json:"deliveryStatus"`
	// This field is from variant [EventSend].
	DestinationType string `json:"destinationType"`
	// This field is from variant [EventSend].
	DeliveredContent any `json:"deliveredContent"`
	// This field is from variant [EventSend].
	GoogleDriveOutput EventSendGoogleDriveOutput `json:"googleDriveOutput"`
	// This field is from variant [EventSend].
	S3Output EventSendS3Output `json:"s3Output"`
	// This field is from variant [EventSend].
	WebhookOutput EventSendWebhookOutput `json:"webhookOutput"`
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

// anyEvent is implemented by each variant of [EventUnion] to add type safety for
// the return type of [EventUnion.AsAny]
type anyEvent interface {
	implEventUnion()
}

func (EventTransform) implEventUnion()            {}
func (EventExtract) implEventUnion()              {}
func (EventRoute) implEventUnion()                {}
func (EventClassify) implEventUnion()             {}
func (EventSplitCollection) implEventUnion()      {}
func (EventSplitItem) implEventUnion()            {}
func (ErrorEvent) implEventUnion()                {}
func (EventJoin) implEventUnion()                 {}
func (EventEnrich) implEventUnion()               {}
func (EventCollectionProcessing) implEventUnion() {}
func (EventSend) implEventUnion()                 {}

// Use the following switch statement to find the correct variant
//
//	switch variant := EventUnion.AsAny().(type) {
//	case bem.EventTransform:
//	case bem.EventExtract:
//	case bem.EventRoute:
//	case bem.EventClassify:
//	case bem.EventSplitCollection:
//	case bem.EventSplitItem:
//	case bem.ErrorEvent:
//	case bem.EventJoin:
//	case bem.EventEnrich:
//	case bem.EventCollectionProcessing:
//	case bem.EventSend:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u EventUnion) AsAny() anyEvent {
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

func (u EventUnion) AsTransform() (v EventTransform) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventUnion) AsExtract() (v EventExtract) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventUnion) AsRoute() (v EventRoute) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventUnion) AsClassify() (v EventClassify) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventUnion) AsSplitCollection() (v EventSplitCollection) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventUnion) AsSplitItem() (v EventSplitItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventUnion) AsError() (v ErrorEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventUnion) AsJoin() (v EventJoin) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventUnion) AsEnrich() (v EventEnrich) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventUnion) AsCollectionProcessing() (v EventCollectionProcessing) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventUnion) AsSend() (v EventSend) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EventUnion) RawJSON() string { return u.JSON.raw }

func (r *EventUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventUnionCorrectedContent is an implicit subunion of [EventUnion].
// EventUnionCorrectedContent provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the [EventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAnyArray OfString OfFloat OfBool]
type EventUnionCorrectedContent struct {
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

func (r *EventUnionCorrectedContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventUnionInputs is an implicit subunion of [EventUnion]. EventUnionInputs
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the [EventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfEventTransformInputs OfEventExtractInputs]
type EventUnionInputs struct {
	// This field will be present if the value is a [[]EventTransformInput] instead of
	// an object.
	OfEventTransformInputs []EventTransformInput `json:",inline"`
	// This field will be present if the value is a [[]EventExtractInput] instead of an
	// object.
	OfEventExtractInputs []EventExtractInput `json:",inline"`
	JSON                 struct {
		OfEventTransformInputs respjson.Field
		OfEventExtractInputs   respjson.Field
		raw                    string
	} `json:"-"`
}

func (r *EventUnionInputs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventUnionMetadata is an implicit subunion of [EventUnion]. EventUnionMetadata
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the [EventUnion].
type EventUnionMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	JSON                           struct {
		DurationFunctionToEventSeconds respjson.Field
		raw                            string
	} `json:"-"`
}

func (r *EventUnionMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventUnionPrintPageOutput is an implicit subunion of [EventUnion].
// EventUnionPrintPageOutput provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the [EventUnion].
type EventUnionPrintPageOutput struct {
	ItemCount int64 `json:"itemCount"`
	// This field is from variant [EventSplitCollectionPrintPageOutput].
	Items []EventSplitCollectionPrintPageOutputItem `json:"items"`
	// This field is from variant [EventSplitItemPrintPageOutput].
	CollectionReferenceID string `json:"collectionReferenceID"`
	// This field is from variant [EventSplitItemPrintPageOutput].
	ItemOffset int64 `json:"itemOffset"`
	// This field is from variant [EventSplitItemPrintPageOutput].
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

func (r *EventUnionPrintPageOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventUnionSemanticPageOutput is an implicit subunion of [EventUnion].
// EventUnionSemanticPageOutput provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the [EventUnion].
type EventUnionSemanticPageOutput struct {
	ItemCount int64 `json:"itemCount"`
	// This field is from variant [EventSplitCollectionSemanticPageOutput].
	Items     []EventSplitCollectionSemanticPageOutputItem `json:"items"`
	PageCount int64                                        `json:"pageCount"`
	// This field is from variant [EventSplitItemSemanticPageOutput].
	CollectionReferenceID string `json:"collectionReferenceID"`
	// This field is from variant [EventSplitItemSemanticPageOutput].
	ItemClass string `json:"itemClass"`
	// This field is from variant [EventSplitItemSemanticPageOutput].
	ItemClassCount int64 `json:"itemClassCount"`
	// This field is from variant [EventSplitItemSemanticPageOutput].
	ItemClassOffset int64 `json:"itemClassOffset"`
	// This field is from variant [EventSplitItemSemanticPageOutput].
	ItemOffset int64 `json:"itemOffset"`
	// This field is from variant [EventSplitItemSemanticPageOutput].
	PageEnd int64 `json:"pageEnd"`
	// This field is from variant [EventSplitItemSemanticPageOutput].
	PageStart int64 `json:"pageStart"`
	// This field is from variant [EventSplitItemSemanticPageOutput].
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

func (r *EventUnionSemanticPageOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventTransform struct {
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
	CorrectedContent EventTransformCorrectedContentUnion `json:"correctedContent" api:"nullable"`
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
	Inputs []EventTransformInput `json:"inputs" api:"nullable"`
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
	LastPublishErrorAt string                 `json:"lastPublishErrorAt" api:"nullable"`
	Metadata           EventTransformMetadata `json:"metadata"`
	// Accuracy, precision, recall, and F1 score when corrected JSON is provided.
	Metrics EventTransformMetrics `json:"metrics" api:"nullable"`
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
func (r EventTransform) RawJSON() string { return r.JSON.raw }
func (r *EventTransform) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventTransformCorrectedContentUnion contains all possible properties and values
// from [EventTransformCorrectedContentOutput], [[]any], [string], [float64],
// [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAnyArray OfString OfFloat OfBool]
type EventTransformCorrectedContentUnion struct {
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field is from variant [EventTransformCorrectedContentOutput].
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

func (u EventTransformCorrectedContentUnion) AsEventTransformCorrectedContentOutput() (v EventTransformCorrectedContentOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventTransformCorrectedContentUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventTransformCorrectedContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventTransformCorrectedContentUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventTransformCorrectedContentUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EventTransformCorrectedContentUnion) RawJSON() string { return u.JSON.raw }

func (r *EventTransformCorrectedContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventTransformCorrectedContentOutput struct {
	Output []AnyTypeUnion `json:"output"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventTransformCorrectedContentOutput) RawJSON() string { return r.JSON.raw }
func (r *EventTransformCorrectedContentOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventTransformInput struct {
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
func (r EventTransformInput) RawJSON() string { return r.JSON.raw }
func (r *EventTransformInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventTransformMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventTransformMetadata) RawJSON() string { return r.JSON.raw }
func (r *EventTransformMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Accuracy, precision, recall, and F1 score when corrected JSON is provided.
type EventTransformMetrics struct {
	Differences []EventTransformMetricsDifference `json:"differences"`
	Metrics     EventTransformMetricsMetrics      `json:"metrics"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Differences respjson.Field
		Metrics     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventTransformMetrics) RawJSON() string { return r.JSON.raw }
func (r *EventTransformMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventTransformMetricsDifference struct {
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
func (r EventTransformMetricsDifference) RawJSON() string { return r.JSON.raw }
func (r *EventTransformMetricsDifference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventTransformMetricsMetrics struct {
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
func (r EventTransformMetricsMetrics) RawJSON() string { return r.JSON.raw }
func (r *EventTransformMetricsMetrics) UnmarshalJSON(data []byte) error {
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
type EventExtract struct {
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
	CorrectedContent EventExtractCorrectedContentUnion `json:"correctedContent" api:"nullable"`
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
	Inputs []EventExtractInput `json:"inputs" api:"nullable"`
	// The input type of the content you're sending for transformation.
	//
	// Any of "csv", "docx", "email", "heic", "html", "jpeg", "json", "heif", "m4a",
	// "mp3", "pdf", "png", "text", "wav", "webp", "xls", "xlsx", "xml".
	InputType string `json:"inputType"`
	// List of properties that were invalid in the input.
	InvalidProperties []string             `json:"invalidProperties"`
	Metadata          EventExtractMetadata `json:"metadata"`
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
func (r EventExtract) RawJSON() string { return r.JSON.raw }
func (r *EventExtract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventExtractCorrectedContentUnion contains all possible properties and values
// from [EventExtractCorrectedContentOutput], [[]any], [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAnyArray OfString OfFloat OfBool]
type EventExtractCorrectedContentUnion struct {
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field is from variant [EventExtractCorrectedContentOutput].
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

func (u EventExtractCorrectedContentUnion) AsEventExtractCorrectedContentOutput() (v EventExtractCorrectedContentOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventExtractCorrectedContentUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventExtractCorrectedContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventExtractCorrectedContentUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventExtractCorrectedContentUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EventExtractCorrectedContentUnion) RawJSON() string { return u.JSON.raw }

func (r *EventExtractCorrectedContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventExtractCorrectedContentOutput struct {
	Output []AnyTypeUnion `json:"output"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventExtractCorrectedContentOutput) RawJSON() string { return r.JSON.raw }
func (r *EventExtractCorrectedContentOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventExtractInput struct {
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
func (r EventExtractInput) RawJSON() string { return r.JSON.raw }
func (r *EventExtractInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventExtractMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventExtractMetadata) RawJSON() string { return r.JSON.raw }
func (r *EventExtractMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventRoute struct {
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
	InboundEmail InboundEmailEvent  `json:"inboundEmail"`
	Metadata     EventRouteMetadata `json:"metadata"`
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
func (r EventRoute) RawJSON() string { return r.JSON.raw }
func (r *EventRoute) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventRouteMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventRouteMetadata) RawJSON() string { return r.JSON.raw }
func (r *EventRouteMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventClassify struct {
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
	InboundEmail InboundEmailEvent     `json:"inboundEmail"`
	Metadata     EventClassifyMetadata `json:"metadata"`
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
func (r EventClassify) RawJSON() string { return r.JSON.raw }
func (r *EventClassify) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventClassifyMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventClassifyMetadata) RawJSON() string { return r.JSON.raw }
func (r *EventClassifyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventSplitCollection struct {
	// Unique ID generated by bem to identify the event.
	EventID string `json:"eventID" api:"required"`
	// Unique identifier of function that this event is associated with.
	FunctionID string `json:"functionID" api:"required"`
	// Unique name of function that this event is associated with.
	FunctionName string `json:"functionName" api:"required"`
	// Any of "print_page", "semantic_page".
	OutputType      string                              `json:"outputType" api:"required"`
	PrintPageOutput EventSplitCollectionPrintPageOutput `json:"printPageOutput" api:"required"`
	// The unique ID you use internally to refer to this data point, propagated from
	// the original function input.
	ReferenceID        string                                 `json:"referenceID" api:"required"`
	SemanticPageOutput EventSplitCollectionSemanticPageOutput `json:"semanticPageOutput" api:"required"`
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
	InboundEmail InboundEmailEvent            `json:"inboundEmail"`
	Metadata     EventSplitCollectionMetadata `json:"metadata"`
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
func (r EventSplitCollection) RawJSON() string { return r.JSON.raw }
func (r *EventSplitCollection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventSplitCollectionPrintPageOutput struct {
	ItemCount int64                                     `json:"itemCount"`
	Items     []EventSplitCollectionPrintPageOutputItem `json:"items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemCount   respjson.Field
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventSplitCollectionPrintPageOutput) RawJSON() string { return r.JSON.raw }
func (r *EventSplitCollectionPrintPageOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventSplitCollectionPrintPageOutputItem struct {
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
func (r EventSplitCollectionPrintPageOutputItem) RawJSON() string { return r.JSON.raw }
func (r *EventSplitCollectionPrintPageOutputItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventSplitCollectionSemanticPageOutput struct {
	ItemCount int64                                        `json:"itemCount"`
	Items     []EventSplitCollectionSemanticPageOutputItem `json:"items"`
	PageCount int64                                        `json:"pageCount"`
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
func (r EventSplitCollectionSemanticPageOutput) RawJSON() string { return r.JSON.raw }
func (r *EventSplitCollectionSemanticPageOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventSplitCollectionSemanticPageOutputItem struct {
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
func (r EventSplitCollectionSemanticPageOutputItem) RawJSON() string { return r.JSON.raw }
func (r *EventSplitCollectionSemanticPageOutputItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventSplitCollectionMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventSplitCollectionMetadata) RawJSON() string { return r.JSON.raw }
func (r *EventSplitCollectionMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventSplitItem struct {
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
	InboundEmail       InboundEmailEvent                `json:"inboundEmail"`
	Metadata           EventSplitItemMetadata           `json:"metadata"`
	PrintPageOutput    EventSplitItemPrintPageOutput    `json:"printPageOutput"`
	SemanticPageOutput EventSplitItemSemanticPageOutput `json:"semanticPageOutput"`
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
func (r EventSplitItem) RawJSON() string { return r.JSON.raw }
func (r *EventSplitItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventSplitItemMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventSplitItemMetadata) RawJSON() string { return r.JSON.raw }
func (r *EventSplitItemMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventSplitItemPrintPageOutput struct {
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
func (r EventSplitItemPrintPageOutput) RawJSON() string { return r.JSON.raw }
func (r *EventSplitItemPrintPageOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventSplitItemSemanticPageOutput struct {
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
func (r EventSplitItemSemanticPageOutput) RawJSON() string { return r.JSON.raw }
func (r *EventSplitItemSemanticPageOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventJoin struct {
	// Unique ID generated by bem to identify the event.
	EventID string `json:"eventID" api:"required"`
	// Unique identifier of function that this event is associated with.
	FunctionID string `json:"functionID" api:"required"`
	// Unique name of function that this event is associated with.
	FunctionName string `json:"functionName" api:"required"`
	// List of properties that were invalid in the input.
	InvalidProperties []string `json:"invalidProperties" api:"required"`
	// The items that were joined.
	Items []EventJoinItem `json:"items" api:"required"`
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
	InboundEmail InboundEmailEvent `json:"inboundEmail"`
	Metadata     EventJoinMetadata `json:"metadata"`
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
func (r EventJoin) RawJSON() string { return r.JSON.raw }
func (r *EventJoin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventJoinItem struct {
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
func (r EventJoinItem) RawJSON() string { return r.JSON.raw }
func (r *EventJoinItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventJoinMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventJoinMetadata) RawJSON() string { return r.JSON.raw }
func (r *EventJoinMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventEnrich struct {
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
	InboundEmail InboundEmailEvent   `json:"inboundEmail"`
	Metadata     EventEnrichMetadata `json:"metadata"`
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
func (r EventEnrich) RawJSON() string { return r.JSON.raw }
func (r *EventEnrich) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventEnrichMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventEnrichMetadata) RawJSON() string { return r.JSON.raw }
func (r *EventEnrichMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventCollectionProcessing struct {
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
	InboundEmail InboundEmailEvent                 `json:"inboundEmail"`
	Metadata     EventCollectionProcessingMetadata `json:"metadata"`
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
func (r EventCollectionProcessing) RawJSON() string { return r.JSON.raw }
func (r *EventCollectionProcessing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventCollectionProcessingMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventCollectionProcessingMetadata) RawJSON() string { return r.JSON.raw }
func (r *EventCollectionProcessingMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventSend struct {
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
	GoogleDriveOutput EventSendGoogleDriveOutput `json:"googleDriveOutput"`
	// The inbound email that triggered this event.
	InboundEmail InboundEmailEvent `json:"inboundEmail"`
	Metadata     EventSendMetadata `json:"metadata"`
	// Metadata returned when a Send function delivers to an S3 bucket.
	S3Output EventSendS3Output `json:"s3Output"`
	// Metadata returned when a Send function delivers to a webhook.
	WebhookOutput EventSendWebhookOutput `json:"webhookOutput"`
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
func (r EventSend) RawJSON() string { return r.JSON.raw }
func (r *EventSend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata returned when a Send function delivers to Google Drive.
type EventSendGoogleDriveOutput struct {
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
func (r EventSendGoogleDriveOutput) RawJSON() string { return r.JSON.raw }
func (r *EventSendGoogleDriveOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventSendMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventSendMetadata) RawJSON() string { return r.JSON.raw }
func (r *EventSendMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata returned when a Send function delivers to an S3 bucket.
type EventSendS3Output struct {
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
func (r EventSendS3Output) RawJSON() string { return r.JSON.raw }
func (r *EventSendS3Output) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata returned when a Send function delivers to a webhook.
type EventSendWebhookOutput struct {
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
func (r EventSendWebhookOutput) RawJSON() string { return r.JSON.raw }
func (r *EventSendWebhookOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponse struct {
	// V3 read-side event union. Superset of the shared `Event` union: it contains
	// every shared variant verbatim (backward compatible) and adds the V3-only
	// `extract` and `classify` variants.
	Output EventUnion `json:"output" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputGetResponse) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListParams struct {
	EndingBefore param.Opt[string] `query:"endingBefore,omitzero" json:"-"`
	// When `true`, includes intermediate events (those that spawned a downstream
	// function call). Default: `false`.
	IncludeIntermediate param.Opt[bool] `query:"includeIntermediate,omitzero" json:"-"`
	// If `true`, only outputs with a corrected (labelled) payload. If `false`, only
	// outputs that are not labelled. If omitted, no filter is applied.
	IsLabelled param.Opt[bool] `query:"isLabelled,omitzero" json:"-"`
	// If `true`, only regression-marked outputs. If `false`, only non-regression
	// outputs. If omitted, no filter is applied.
	//
	// Note: clients migrating from `/v1-beta/transformations` should pass
	// `isRegression=false` explicitly to preserve the legacy default (regressions
	// hidden unless explicitly requested).
	IsRegression param.Opt[bool]  `query:"isRegression,omitzero" json:"-"`
	Limit        param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Case-insensitive substring match against `referenceID`.
	ReferenceIDSubstring param.Opt[string] `query:"referenceIDSubstring,omitzero" json:"-"`
	StartingAfter        param.Opt[string] `query:"startingAfter,omitzero" json:"-"`
	// Filter to outputs from specific calls.
	CallIDs []string `query:"callIDs,omitzero" json:"-"`
	// Filter to specific output events by their event IDs (KSUIDs).
	EventIDs      []string `query:"eventIDs,omitzero" json:"-"`
	FunctionIDs   []string `query:"functionIDs,omitzero" json:"-"`
	FunctionNames []string `query:"functionNames,omitzero" json:"-"`
	// Filter to specific function version numbers.
	FunctionVersionNums []int64  `query:"functionVersionNums,omitzero" json:"-"`
	ReferenceIDs        []string `query:"referenceIDs,omitzero" json:"-"`
	// Any of "asc", "desc".
	SortOrder OutputListParamsSortOrder `query:"sortOrder,omitzero" json:"-"`
	// Filter by legacy transformation IDs. Provided for backwards compatibility with
	// clients migrating from `/v1-beta/transformations`.
	TransformationIDs []string `query:"transformationIDs,omitzero" json:"-"`
	WorkflowIDs       []string `query:"workflowIDs,omitzero" json:"-"`
	WorkflowNames     []string `query:"workflowNames,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OutputListParams]'s query parameters as `url.Values`.
func (r OutputListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OutputListParamsSortOrder string

const (
	OutputListParamsSortOrderAsc  OutputListParamsSortOrder = "asc"
	OutputListParamsSortOrderDesc OutputListParamsSortOrder = "desc"
)
