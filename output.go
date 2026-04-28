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
func (r *OutputService) List(ctx context.Context, query OutputListParams, opts ...option.RequestOption) (res *pagination.OutputsPage[OutputListResponseUnion], err error) {
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
func (r *OutputService) ListAutoPaging(ctx context.Context, query OutputListParams, opts ...option.RequestOption) *pagination.OutputsPageAutoPager[OutputListResponseUnion] {
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

type OutputGetResponse struct {
	// V3 read-side event union. Superset of the shared `Event` union: it contains
	// every shared variant verbatim (backward compatible) and adds the V3-only
	// `extract` and `classify` variants.
	Output OutputGetResponseOutputUnion `json:"output" api:"required"`
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

// OutputGetResponseOutputUnion contains all possible properties and values from
// [OutputGetResponseOutputTransform], [OutputGetResponseOutputExtract],
// [OutputGetResponseOutputRoute], [OutputGetResponseOutputClassify],
// [OutputGetResponseOutputSplitCollection], [OutputGetResponseOutputSplitItem],
// [ErrorEvent], [OutputGetResponseOutputJoin], [OutputGetResponseOutputEnrich],
// [OutputGetResponseOutputCollectionProcessing], [OutputGetResponseOutputSend].
//
// Use the [OutputGetResponseOutputUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type OutputGetResponseOutputUnion struct {
	EventID            string  `json:"eventID"`
	FunctionID         string  `json:"functionID"`
	FunctionName       string  `json:"functionName"`
	ItemCount          int64   `json:"itemCount"`
	ItemOffset         int64   `json:"itemOffset"`
	ReferenceID        string  `json:"referenceID"`
	TransformedContent any     `json:"transformedContent"`
	AvgConfidence      float64 `json:"avgConfidence"`
	CallID             string  `json:"callID"`
	// This field is a union of
	// [OutputGetResponseOutputTransformCorrectedContentUnion],
	// [OutputGetResponseOutputExtractCorrectedContentUnion]
	CorrectedContent OutputGetResponseOutputUnionCorrectedContent `json:"correctedContent"`
	CreatedAt        time.Time                                    `json:"createdAt"`
	// Any of "transform", "extract", "route", "classify", "split_collection",
	// "split_item", "error", "join", "enrich", "collection_processing", "send".
	EventType             string `json:"eventType"`
	FieldConfidences      any    `json:"fieldConfidences"`
	FunctionCallID        string `json:"functionCallID"`
	FunctionCallTryNumber int64  `json:"functionCallTryNumber"`
	FunctionVersionNum    int64  `json:"functionVersionNum"`
	// This field is from variant [OutputGetResponseOutputTransform].
	InboundEmail InboundEmailEvent `json:"inboundEmail"`
	// This field is a union of [[]OutputGetResponseOutputTransformInput],
	// [[]OutputGetResponseOutputExtractInput]
	Inputs            OutputGetResponseOutputUnionInputs `json:"inputs"`
	InputType         string                             `json:"inputType"`
	InvalidProperties []string                           `json:"invalidProperties"`
	// This field is from variant [OutputGetResponseOutputTransform].
	IsRegression bool `json:"isRegression"`
	// This field is from variant [OutputGetResponseOutputTransform].
	LastPublishErrorAt string `json:"lastPublishErrorAt"`
	// This field is a union of [OutputGetResponseOutputTransformMetadata],
	// [OutputGetResponseOutputExtractMetadata],
	// [OutputGetResponseOutputRouteMetadata],
	// [OutputGetResponseOutputClassifyMetadata],
	// [OutputGetResponseOutputSplitCollectionMetadata],
	// [OutputGetResponseOutputSplitItemMetadata], [ErrorEventMetadata],
	// [OutputGetResponseOutputJoinMetadata], [OutputGetResponseOutputEnrichMetadata],
	// [OutputGetResponseOutputCollectionProcessingMetadata],
	// [OutputGetResponseOutputSendMetadata]
	Metadata OutputGetResponseOutputUnionMetadata `json:"metadata"`
	// This field is from variant [OutputGetResponseOutputTransform].
	Metrics OutputGetResponseOutputTransformMetrics `json:"metrics"`
	// This field is from variant [OutputGetResponseOutputTransform].
	OrderMatching bool `json:"orderMatching"`
	// This field is from variant [OutputGetResponseOutputTransform].
	PipelineID string `json:"pipelineID"`
	// This field is from variant [OutputGetResponseOutputTransform].
	PublishedAt        time.Time `json:"publishedAt"`
	S3URL              string    `json:"s3URL"`
	TransformationID   string    `json:"transformationID"`
	WorkflowID         string    `json:"workflowID"`
	WorkflowName       string    `json:"workflowName"`
	WorkflowVersionNum int64     `json:"workflowVersionNum"`
	Choice             string    `json:"choice"`
	OutputType         string    `json:"outputType"`
	// This field is a union of
	// [OutputGetResponseOutputSplitCollectionPrintPageOutput],
	// [OutputGetResponseOutputSplitItemPrintPageOutput]
	PrintPageOutput OutputGetResponseOutputUnionPrintPageOutput `json:"printPageOutput"`
	// This field is a union of
	// [OutputGetResponseOutputSplitCollectionSemanticPageOutput],
	// [OutputGetResponseOutputSplitItemSemanticPageOutput]
	SemanticPageOutput OutputGetResponseOutputUnionSemanticPageOutput `json:"semanticPageOutput"`
	// This field is from variant [ErrorEvent].
	Message string `json:"message"`
	// This field is from variant [OutputGetResponseOutputJoin].
	Items []OutputGetResponseOutputJoinItem `json:"items"`
	// This field is from variant [OutputGetResponseOutputJoin].
	JoinType string `json:"joinType"`
	// This field is from variant [OutputGetResponseOutputEnrich].
	EnrichedContent any `json:"enrichedContent"`
	// This field is from variant [OutputGetResponseOutputCollectionProcessing].
	CollectionID string `json:"collectionID"`
	// This field is from variant [OutputGetResponseOutputCollectionProcessing].
	CollectionName string `json:"collectionName"`
	// This field is from variant [OutputGetResponseOutputCollectionProcessing].
	Operation string `json:"operation"`
	// This field is from variant [OutputGetResponseOutputCollectionProcessing].
	ProcessedCount int64 `json:"processedCount"`
	// This field is from variant [OutputGetResponseOutputCollectionProcessing].
	Status string `json:"status"`
	// This field is from variant [OutputGetResponseOutputCollectionProcessing].
	CollectionItemIDs []string `json:"collectionItemIDs"`
	// This field is from variant [OutputGetResponseOutputCollectionProcessing].
	ErrorMessage string `json:"errorMessage"`
	// This field is from variant [OutputGetResponseOutputSend].
	DeliveryStatus string `json:"deliveryStatus"`
	// This field is from variant [OutputGetResponseOutputSend].
	DestinationType string `json:"destinationType"`
	// This field is from variant [OutputGetResponseOutputSend].
	DeliveredContent any `json:"deliveredContent"`
	// This field is from variant [OutputGetResponseOutputSend].
	GoogleDriveOutput OutputGetResponseOutputSendGoogleDriveOutput `json:"googleDriveOutput"`
	// This field is from variant [OutputGetResponseOutputSend].
	S3Output OutputGetResponseOutputSendS3Output `json:"s3Output"`
	// This field is from variant [OutputGetResponseOutputSend].
	WebhookOutput OutputGetResponseOutputSendWebhookOutput `json:"webhookOutput"`
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

// anyOutputGetResponseOutput is implemented by each variant of
// [OutputGetResponseOutputUnion] to add type safety for the return type of
// [OutputGetResponseOutputUnion.AsAny]
type anyOutputGetResponseOutput interface {
	implOutputGetResponseOutputUnion()
}

func (OutputGetResponseOutputTransform) implOutputGetResponseOutputUnion()            {}
func (OutputGetResponseOutputExtract) implOutputGetResponseOutputUnion()              {}
func (OutputGetResponseOutputRoute) implOutputGetResponseOutputUnion()                {}
func (OutputGetResponseOutputClassify) implOutputGetResponseOutputUnion()             {}
func (OutputGetResponseOutputSplitCollection) implOutputGetResponseOutputUnion()      {}
func (OutputGetResponseOutputSplitItem) implOutputGetResponseOutputUnion()            {}
func (ErrorEvent) implOutputGetResponseOutputUnion()                                  {}
func (OutputGetResponseOutputJoin) implOutputGetResponseOutputUnion()                 {}
func (OutputGetResponseOutputEnrich) implOutputGetResponseOutputUnion()               {}
func (OutputGetResponseOutputCollectionProcessing) implOutputGetResponseOutputUnion() {}
func (OutputGetResponseOutputSend) implOutputGetResponseOutputUnion()                 {}

// Use the following switch statement to find the correct variant
//
//	switch variant := OutputGetResponseOutputUnion.AsAny().(type) {
//	case bem.OutputGetResponseOutputTransform:
//	case bem.OutputGetResponseOutputExtract:
//	case bem.OutputGetResponseOutputRoute:
//	case bem.OutputGetResponseOutputClassify:
//	case bem.OutputGetResponseOutputSplitCollection:
//	case bem.OutputGetResponseOutputSplitItem:
//	case bem.ErrorEvent:
//	case bem.OutputGetResponseOutputJoin:
//	case bem.OutputGetResponseOutputEnrich:
//	case bem.OutputGetResponseOutputCollectionProcessing:
//	case bem.OutputGetResponseOutputSend:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u OutputGetResponseOutputUnion) AsAny() anyOutputGetResponseOutput {
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

func (u OutputGetResponseOutputUnion) AsTransform() (v OutputGetResponseOutputTransform) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputGetResponseOutputUnion) AsExtract() (v OutputGetResponseOutputExtract) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputGetResponseOutputUnion) AsRoute() (v OutputGetResponseOutputRoute) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputGetResponseOutputUnion) AsClassify() (v OutputGetResponseOutputClassify) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputGetResponseOutputUnion) AsSplitCollection() (v OutputGetResponseOutputSplitCollection) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputGetResponseOutputUnion) AsSplitItem() (v OutputGetResponseOutputSplitItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputGetResponseOutputUnion) AsError() (v ErrorEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputGetResponseOutputUnion) AsJoin() (v OutputGetResponseOutputJoin) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputGetResponseOutputUnion) AsEnrich() (v OutputGetResponseOutputEnrich) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputGetResponseOutputUnion) AsCollectionProcessing() (v OutputGetResponseOutputCollectionProcessing) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputGetResponseOutputUnion) AsSend() (v OutputGetResponseOutputSend) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OutputGetResponseOutputUnion) RawJSON() string { return u.JSON.raw }

func (r *OutputGetResponseOutputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OutputGetResponseOutputUnionCorrectedContent is an implicit subunion of
// [OutputGetResponseOutputUnion]. OutputGetResponseOutputUnionCorrectedContent
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [OutputGetResponseOutputUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAnyArray OfString OfFloat OfBool]
type OutputGetResponseOutputUnionCorrectedContent struct {
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

func (r *OutputGetResponseOutputUnionCorrectedContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OutputGetResponseOutputUnionInputs is an implicit subunion of
// [OutputGetResponseOutputUnion]. OutputGetResponseOutputUnionInputs provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [OutputGetResponseOutputUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfOutputGetResponseOutputTransformInputs
// OfOutputGetResponseOutputExtractInputs]
type OutputGetResponseOutputUnionInputs struct {
	// This field will be present if the value is a
	// [[]OutputGetResponseOutputTransformInput] instead of an object.
	OfOutputGetResponseOutputTransformInputs []OutputGetResponseOutputTransformInput `json:",inline"`
	// This field will be present if the value is a
	// [[]OutputGetResponseOutputExtractInput] instead of an object.
	OfOutputGetResponseOutputExtractInputs []OutputGetResponseOutputExtractInput `json:",inline"`
	JSON                                   struct {
		OfOutputGetResponseOutputTransformInputs respjson.Field
		OfOutputGetResponseOutputExtractInputs   respjson.Field
		raw                                      string
	} `json:"-"`
}

func (r *OutputGetResponseOutputUnionInputs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OutputGetResponseOutputUnionMetadata is an implicit subunion of
// [OutputGetResponseOutputUnion]. OutputGetResponseOutputUnionMetadata provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [OutputGetResponseOutputUnion].
type OutputGetResponseOutputUnionMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	JSON                           struct {
		DurationFunctionToEventSeconds respjson.Field
		raw                            string
	} `json:"-"`
}

func (r *OutputGetResponseOutputUnionMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OutputGetResponseOutputUnionPrintPageOutput is an implicit subunion of
// [OutputGetResponseOutputUnion]. OutputGetResponseOutputUnionPrintPageOutput
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [OutputGetResponseOutputUnion].
type OutputGetResponseOutputUnionPrintPageOutput struct {
	ItemCount int64 `json:"itemCount"`
	// This field is from variant
	// [OutputGetResponseOutputSplitCollectionPrintPageOutput].
	Items []OutputGetResponseOutputSplitCollectionPrintPageOutputItem `json:"items"`
	// This field is from variant [OutputGetResponseOutputSplitItemPrintPageOutput].
	CollectionReferenceID string `json:"collectionReferenceID"`
	// This field is from variant [OutputGetResponseOutputSplitItemPrintPageOutput].
	ItemOffset int64 `json:"itemOffset"`
	// This field is from variant [OutputGetResponseOutputSplitItemPrintPageOutput].
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

func (r *OutputGetResponseOutputUnionPrintPageOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OutputGetResponseOutputUnionSemanticPageOutput is an implicit subunion of
// [OutputGetResponseOutputUnion]. OutputGetResponseOutputUnionSemanticPageOutput
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [OutputGetResponseOutputUnion].
type OutputGetResponseOutputUnionSemanticPageOutput struct {
	ItemCount int64 `json:"itemCount"`
	// This field is from variant
	// [OutputGetResponseOutputSplitCollectionSemanticPageOutput].
	Items     []OutputGetResponseOutputSplitCollectionSemanticPageOutputItem `json:"items"`
	PageCount int64                                                          `json:"pageCount"`
	// This field is from variant [OutputGetResponseOutputSplitItemSemanticPageOutput].
	CollectionReferenceID string `json:"collectionReferenceID"`
	// This field is from variant [OutputGetResponseOutputSplitItemSemanticPageOutput].
	ItemClass string `json:"itemClass"`
	// This field is from variant [OutputGetResponseOutputSplitItemSemanticPageOutput].
	ItemClassCount int64 `json:"itemClassCount"`
	// This field is from variant [OutputGetResponseOutputSplitItemSemanticPageOutput].
	ItemClassOffset int64 `json:"itemClassOffset"`
	// This field is from variant [OutputGetResponseOutputSplitItemSemanticPageOutput].
	ItemOffset int64 `json:"itemOffset"`
	// This field is from variant [OutputGetResponseOutputSplitItemSemanticPageOutput].
	PageEnd int64 `json:"pageEnd"`
	// This field is from variant [OutputGetResponseOutputSplitItemSemanticPageOutput].
	PageStart int64 `json:"pageStart"`
	// This field is from variant [OutputGetResponseOutputSplitItemSemanticPageOutput].
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

func (r *OutputGetResponseOutputUnionSemanticPageOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputTransform struct {
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
	CorrectedContent OutputGetResponseOutputTransformCorrectedContentUnion `json:"correctedContent" api:"nullable"`
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
	Inputs []OutputGetResponseOutputTransformInput `json:"inputs" api:"nullable"`
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
	LastPublishErrorAt string                                   `json:"lastPublishErrorAt" api:"nullable"`
	Metadata           OutputGetResponseOutputTransformMetadata `json:"metadata"`
	// Accuracy, precision, recall, and F1 score when corrected JSON is provided.
	Metrics OutputGetResponseOutputTransformMetrics `json:"metrics" api:"nullable"`
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
func (r OutputGetResponseOutputTransform) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputTransform) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OutputGetResponseOutputTransformCorrectedContentUnion contains all possible
// properties and values from
// [OutputGetResponseOutputTransformCorrectedContentOutput], [[]any], [string],
// [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAnyArray OfString OfFloat OfBool]
type OutputGetResponseOutputTransformCorrectedContentUnion struct {
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field is from variant
	// [OutputGetResponseOutputTransformCorrectedContentOutput].
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

func (u OutputGetResponseOutputTransformCorrectedContentUnion) AsOutputGetResponseOutputTransformCorrectedContentOutput() (v OutputGetResponseOutputTransformCorrectedContentOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputGetResponseOutputTransformCorrectedContentUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputGetResponseOutputTransformCorrectedContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputGetResponseOutputTransformCorrectedContentUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputGetResponseOutputTransformCorrectedContentUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OutputGetResponseOutputTransformCorrectedContentUnion) RawJSON() string { return u.JSON.raw }

func (r *OutputGetResponseOutputTransformCorrectedContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputTransformCorrectedContentOutput struct {
	Output []AnyTypeUnion `json:"output"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputGetResponseOutputTransformCorrectedContentOutput) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputTransformCorrectedContentOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputTransformInput struct {
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
func (r OutputGetResponseOutputTransformInput) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputTransformInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputTransformMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputGetResponseOutputTransformMetadata) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputTransformMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Accuracy, precision, recall, and F1 score when corrected JSON is provided.
type OutputGetResponseOutputTransformMetrics struct {
	Differences []OutputGetResponseOutputTransformMetricsDifference `json:"differences"`
	Metrics     OutputGetResponseOutputTransformMetricsMetrics      `json:"metrics"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Differences respjson.Field
		Metrics     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputGetResponseOutputTransformMetrics) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputTransformMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputTransformMetricsDifference struct {
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
func (r OutputGetResponseOutputTransformMetricsDifference) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputTransformMetricsDifference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputTransformMetricsMetrics struct {
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
func (r OutputGetResponseOutputTransformMetricsMetrics) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputTransformMetricsMetrics) UnmarshalJSON(data []byte) error {
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
type OutputGetResponseOutputExtract struct {
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
	CorrectedContent OutputGetResponseOutputExtractCorrectedContentUnion `json:"correctedContent" api:"nullable"`
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
	Inputs []OutputGetResponseOutputExtractInput `json:"inputs" api:"nullable"`
	// The input type of the content you're sending for transformation.
	//
	// Any of "csv", "docx", "email", "heic", "html", "jpeg", "json", "heif", "m4a",
	// "mp3", "pdf", "png", "text", "wav", "webp", "xls", "xlsx", "xml".
	InputType string `json:"inputType"`
	// List of properties that were invalid in the input.
	InvalidProperties []string                               `json:"invalidProperties"`
	Metadata          OutputGetResponseOutputExtractMetadata `json:"metadata"`
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
func (r OutputGetResponseOutputExtract) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputExtract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OutputGetResponseOutputExtractCorrectedContentUnion contains all possible
// properties and values from
// [OutputGetResponseOutputExtractCorrectedContentOutput], [[]any], [string],
// [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAnyArray OfString OfFloat OfBool]
type OutputGetResponseOutputExtractCorrectedContentUnion struct {
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field is from variant
	// [OutputGetResponseOutputExtractCorrectedContentOutput].
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

func (u OutputGetResponseOutputExtractCorrectedContentUnion) AsOutputGetResponseOutputExtractCorrectedContentOutput() (v OutputGetResponseOutputExtractCorrectedContentOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputGetResponseOutputExtractCorrectedContentUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputGetResponseOutputExtractCorrectedContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputGetResponseOutputExtractCorrectedContentUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputGetResponseOutputExtractCorrectedContentUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OutputGetResponseOutputExtractCorrectedContentUnion) RawJSON() string { return u.JSON.raw }

func (r *OutputGetResponseOutputExtractCorrectedContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputExtractCorrectedContentOutput struct {
	Output []AnyTypeUnion `json:"output"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputGetResponseOutputExtractCorrectedContentOutput) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputExtractCorrectedContentOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputExtractInput struct {
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
func (r OutputGetResponseOutputExtractInput) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputExtractInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputExtractMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputGetResponseOutputExtractMetadata) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputExtractMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputRoute struct {
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
	InboundEmail InboundEmailEvent                    `json:"inboundEmail"`
	Metadata     OutputGetResponseOutputRouteMetadata `json:"metadata"`
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
func (r OutputGetResponseOutputRoute) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputRoute) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputRouteMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputGetResponseOutputRouteMetadata) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputRouteMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputClassify struct {
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
	InboundEmail InboundEmailEvent                       `json:"inboundEmail"`
	Metadata     OutputGetResponseOutputClassifyMetadata `json:"metadata"`
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
func (r OutputGetResponseOutputClassify) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputClassify) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputClassifyMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputGetResponseOutputClassifyMetadata) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputClassifyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputSplitCollection struct {
	// Unique ID generated by bem to identify the event.
	EventID string `json:"eventID" api:"required"`
	// Unique identifier of function that this event is associated with.
	FunctionID string `json:"functionID" api:"required"`
	// Unique name of function that this event is associated with.
	FunctionName string `json:"functionName" api:"required"`
	// Any of "print_page", "semantic_page".
	OutputType      string                                                `json:"outputType" api:"required"`
	PrintPageOutput OutputGetResponseOutputSplitCollectionPrintPageOutput `json:"printPageOutput" api:"required"`
	// The unique ID you use internally to refer to this data point, propagated from
	// the original function input.
	ReferenceID        string                                                   `json:"referenceID" api:"required"`
	SemanticPageOutput OutputGetResponseOutputSplitCollectionSemanticPageOutput `json:"semanticPageOutput" api:"required"`
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
	InboundEmail InboundEmailEvent                              `json:"inboundEmail"`
	Metadata     OutputGetResponseOutputSplitCollectionMetadata `json:"metadata"`
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
func (r OutputGetResponseOutputSplitCollection) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputSplitCollection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputSplitCollectionPrintPageOutput struct {
	ItemCount int64                                                       `json:"itemCount"`
	Items     []OutputGetResponseOutputSplitCollectionPrintPageOutputItem `json:"items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemCount   respjson.Field
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputGetResponseOutputSplitCollectionPrintPageOutput) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputSplitCollectionPrintPageOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputSplitCollectionPrintPageOutputItem struct {
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
func (r OutputGetResponseOutputSplitCollectionPrintPageOutputItem) RawJSON() string {
	return r.JSON.raw
}
func (r *OutputGetResponseOutputSplitCollectionPrintPageOutputItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputSplitCollectionSemanticPageOutput struct {
	ItemCount int64                                                          `json:"itemCount"`
	Items     []OutputGetResponseOutputSplitCollectionSemanticPageOutputItem `json:"items"`
	PageCount int64                                                          `json:"pageCount"`
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
func (r OutputGetResponseOutputSplitCollectionSemanticPageOutput) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputSplitCollectionSemanticPageOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputSplitCollectionSemanticPageOutputItem struct {
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
func (r OutputGetResponseOutputSplitCollectionSemanticPageOutputItem) RawJSON() string {
	return r.JSON.raw
}
func (r *OutputGetResponseOutputSplitCollectionSemanticPageOutputItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputSplitCollectionMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputGetResponseOutputSplitCollectionMetadata) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputSplitCollectionMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputSplitItem struct {
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
	InboundEmail       InboundEmailEvent                                  `json:"inboundEmail"`
	Metadata           OutputGetResponseOutputSplitItemMetadata           `json:"metadata"`
	PrintPageOutput    OutputGetResponseOutputSplitItemPrintPageOutput    `json:"printPageOutput"`
	SemanticPageOutput OutputGetResponseOutputSplitItemSemanticPageOutput `json:"semanticPageOutput"`
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
func (r OutputGetResponseOutputSplitItem) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputSplitItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputSplitItemMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputGetResponseOutputSplitItemMetadata) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputSplitItemMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputSplitItemPrintPageOutput struct {
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
func (r OutputGetResponseOutputSplitItemPrintPageOutput) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputSplitItemPrintPageOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputSplitItemSemanticPageOutput struct {
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
func (r OutputGetResponseOutputSplitItemSemanticPageOutput) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputSplitItemSemanticPageOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputJoin struct {
	// Unique ID generated by bem to identify the event.
	EventID string `json:"eventID" api:"required"`
	// Unique identifier of function that this event is associated with.
	FunctionID string `json:"functionID" api:"required"`
	// Unique name of function that this event is associated with.
	FunctionName string `json:"functionName" api:"required"`
	// List of properties that were invalid in the input.
	InvalidProperties []string `json:"invalidProperties" api:"required"`
	// The items that were joined.
	Items []OutputGetResponseOutputJoinItem `json:"items" api:"required"`
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
	InboundEmail InboundEmailEvent                   `json:"inboundEmail"`
	Metadata     OutputGetResponseOutputJoinMetadata `json:"metadata"`
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
func (r OutputGetResponseOutputJoin) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputJoin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputJoinItem struct {
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
func (r OutputGetResponseOutputJoinItem) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputJoinItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputJoinMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputGetResponseOutputJoinMetadata) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputJoinMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputEnrich struct {
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
	InboundEmail InboundEmailEvent                     `json:"inboundEmail"`
	Metadata     OutputGetResponseOutputEnrichMetadata `json:"metadata"`
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
func (r OutputGetResponseOutputEnrich) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputEnrich) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputEnrichMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputGetResponseOutputEnrichMetadata) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputEnrichMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputCollectionProcessing struct {
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
	InboundEmail InboundEmailEvent                                   `json:"inboundEmail"`
	Metadata     OutputGetResponseOutputCollectionProcessingMetadata `json:"metadata"`
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
func (r OutputGetResponseOutputCollectionProcessing) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputCollectionProcessing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputCollectionProcessingMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputGetResponseOutputCollectionProcessingMetadata) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputCollectionProcessingMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputSend struct {
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
	GoogleDriveOutput OutputGetResponseOutputSendGoogleDriveOutput `json:"googleDriveOutput"`
	// The inbound email that triggered this event.
	InboundEmail InboundEmailEvent                   `json:"inboundEmail"`
	Metadata     OutputGetResponseOutputSendMetadata `json:"metadata"`
	// Metadata returned when a Send function delivers to an S3 bucket.
	S3Output OutputGetResponseOutputSendS3Output `json:"s3Output"`
	// Metadata returned when a Send function delivers to a webhook.
	WebhookOutput OutputGetResponseOutputSendWebhookOutput `json:"webhookOutput"`
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
func (r OutputGetResponseOutputSend) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputSend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata returned when a Send function delivers to Google Drive.
type OutputGetResponseOutputSendGoogleDriveOutput struct {
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
func (r OutputGetResponseOutputSendGoogleDriveOutput) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputSendGoogleDriveOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputGetResponseOutputSendMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputGetResponseOutputSendMetadata) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputSendMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata returned when a Send function delivers to an S3 bucket.
type OutputGetResponseOutputSendS3Output struct {
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
func (r OutputGetResponseOutputSendS3Output) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputSendS3Output) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata returned when a Send function delivers to a webhook.
type OutputGetResponseOutputSendWebhookOutput struct {
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
func (r OutputGetResponseOutputSendWebhookOutput) RawJSON() string { return r.JSON.raw }
func (r *OutputGetResponseOutputSendWebhookOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OutputListResponseUnion contains all possible properties and values from
// [OutputListResponseTransform], [OutputListResponseExtract],
// [OutputListResponseRoute], [OutputListResponseClassify],
// [OutputListResponseSplitCollection], [OutputListResponseSplitItem],
// [ErrorEvent], [OutputListResponseJoin], [OutputListResponseEnrich],
// [OutputListResponseCollectionProcessing], [OutputListResponseSend].
//
// Use the [OutputListResponseUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type OutputListResponseUnion struct {
	EventID            string  `json:"eventID"`
	FunctionID         string  `json:"functionID"`
	FunctionName       string  `json:"functionName"`
	ItemCount          int64   `json:"itemCount"`
	ItemOffset         int64   `json:"itemOffset"`
	ReferenceID        string  `json:"referenceID"`
	TransformedContent any     `json:"transformedContent"`
	AvgConfidence      float64 `json:"avgConfidence"`
	CallID             string  `json:"callID"`
	// This field is a union of [OutputListResponseTransformCorrectedContentUnion],
	// [OutputListResponseExtractCorrectedContentUnion]
	CorrectedContent OutputListResponseUnionCorrectedContent `json:"correctedContent"`
	CreatedAt        time.Time                               `json:"createdAt"`
	// Any of "transform", "extract", "route", "classify", "split_collection",
	// "split_item", "error", "join", "enrich", "collection_processing", "send".
	EventType             string `json:"eventType"`
	FieldConfidences      any    `json:"fieldConfidences"`
	FunctionCallID        string `json:"functionCallID"`
	FunctionCallTryNumber int64  `json:"functionCallTryNumber"`
	FunctionVersionNum    int64  `json:"functionVersionNum"`
	// This field is from variant [OutputListResponseTransform].
	InboundEmail InboundEmailEvent `json:"inboundEmail"`
	// This field is a union of [[]OutputListResponseTransformInput],
	// [[]OutputListResponseExtractInput]
	Inputs            OutputListResponseUnionInputs `json:"inputs"`
	InputType         string                        `json:"inputType"`
	InvalidProperties []string                      `json:"invalidProperties"`
	// This field is from variant [OutputListResponseTransform].
	IsRegression bool `json:"isRegression"`
	// This field is from variant [OutputListResponseTransform].
	LastPublishErrorAt string `json:"lastPublishErrorAt"`
	// This field is a union of [OutputListResponseTransformMetadata],
	// [OutputListResponseExtractMetadata], [OutputListResponseRouteMetadata],
	// [OutputListResponseClassifyMetadata],
	// [OutputListResponseSplitCollectionMetadata],
	// [OutputListResponseSplitItemMetadata], [ErrorEventMetadata],
	// [OutputListResponseJoinMetadata], [OutputListResponseEnrichMetadata],
	// [OutputListResponseCollectionProcessingMetadata],
	// [OutputListResponseSendMetadata]
	Metadata OutputListResponseUnionMetadata `json:"metadata"`
	// This field is from variant [OutputListResponseTransform].
	Metrics OutputListResponseTransformMetrics `json:"metrics"`
	// This field is from variant [OutputListResponseTransform].
	OrderMatching bool `json:"orderMatching"`
	// This field is from variant [OutputListResponseTransform].
	PipelineID string `json:"pipelineID"`
	// This field is from variant [OutputListResponseTransform].
	PublishedAt        time.Time `json:"publishedAt"`
	S3URL              string    `json:"s3URL"`
	TransformationID   string    `json:"transformationID"`
	WorkflowID         string    `json:"workflowID"`
	WorkflowName       string    `json:"workflowName"`
	WorkflowVersionNum int64     `json:"workflowVersionNum"`
	Choice             string    `json:"choice"`
	OutputType         string    `json:"outputType"`
	// This field is a union of [OutputListResponseSplitCollectionPrintPageOutput],
	// [OutputListResponseSplitItemPrintPageOutput]
	PrintPageOutput OutputListResponseUnionPrintPageOutput `json:"printPageOutput"`
	// This field is a union of [OutputListResponseSplitCollectionSemanticPageOutput],
	// [OutputListResponseSplitItemSemanticPageOutput]
	SemanticPageOutput OutputListResponseUnionSemanticPageOutput `json:"semanticPageOutput"`
	// This field is from variant [ErrorEvent].
	Message string `json:"message"`
	// This field is from variant [OutputListResponseJoin].
	Items []OutputListResponseJoinItem `json:"items"`
	// This field is from variant [OutputListResponseJoin].
	JoinType string `json:"joinType"`
	// This field is from variant [OutputListResponseEnrich].
	EnrichedContent any `json:"enrichedContent"`
	// This field is from variant [OutputListResponseCollectionProcessing].
	CollectionID string `json:"collectionID"`
	// This field is from variant [OutputListResponseCollectionProcessing].
	CollectionName string `json:"collectionName"`
	// This field is from variant [OutputListResponseCollectionProcessing].
	Operation string `json:"operation"`
	// This field is from variant [OutputListResponseCollectionProcessing].
	ProcessedCount int64 `json:"processedCount"`
	// This field is from variant [OutputListResponseCollectionProcessing].
	Status string `json:"status"`
	// This field is from variant [OutputListResponseCollectionProcessing].
	CollectionItemIDs []string `json:"collectionItemIDs"`
	// This field is from variant [OutputListResponseCollectionProcessing].
	ErrorMessage string `json:"errorMessage"`
	// This field is from variant [OutputListResponseSend].
	DeliveryStatus string `json:"deliveryStatus"`
	// This field is from variant [OutputListResponseSend].
	DestinationType string `json:"destinationType"`
	// This field is from variant [OutputListResponseSend].
	DeliveredContent any `json:"deliveredContent"`
	// This field is from variant [OutputListResponseSend].
	GoogleDriveOutput OutputListResponseSendGoogleDriveOutput `json:"googleDriveOutput"`
	// This field is from variant [OutputListResponseSend].
	S3Output OutputListResponseSendS3Output `json:"s3Output"`
	// This field is from variant [OutputListResponseSend].
	WebhookOutput OutputListResponseSendWebhookOutput `json:"webhookOutput"`
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

// anyOutputListResponse is implemented by each variant of
// [OutputListResponseUnion] to add type safety for the return type of
// [OutputListResponseUnion.AsAny]
type anyOutputListResponse interface {
	implOutputListResponseUnion()
}

func (OutputListResponseTransform) implOutputListResponseUnion()            {}
func (OutputListResponseExtract) implOutputListResponseUnion()              {}
func (OutputListResponseRoute) implOutputListResponseUnion()                {}
func (OutputListResponseClassify) implOutputListResponseUnion()             {}
func (OutputListResponseSplitCollection) implOutputListResponseUnion()      {}
func (OutputListResponseSplitItem) implOutputListResponseUnion()            {}
func (ErrorEvent) implOutputListResponseUnion()                             {}
func (OutputListResponseJoin) implOutputListResponseUnion()                 {}
func (OutputListResponseEnrich) implOutputListResponseUnion()               {}
func (OutputListResponseCollectionProcessing) implOutputListResponseUnion() {}
func (OutputListResponseSend) implOutputListResponseUnion()                 {}

// Use the following switch statement to find the correct variant
//
//	switch variant := OutputListResponseUnion.AsAny().(type) {
//	case bem.OutputListResponseTransform:
//	case bem.OutputListResponseExtract:
//	case bem.OutputListResponseRoute:
//	case bem.OutputListResponseClassify:
//	case bem.OutputListResponseSplitCollection:
//	case bem.OutputListResponseSplitItem:
//	case bem.ErrorEvent:
//	case bem.OutputListResponseJoin:
//	case bem.OutputListResponseEnrich:
//	case bem.OutputListResponseCollectionProcessing:
//	case bem.OutputListResponseSend:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u OutputListResponseUnion) AsAny() anyOutputListResponse {
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

func (u OutputListResponseUnion) AsTransform() (v OutputListResponseTransform) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputListResponseUnion) AsExtract() (v OutputListResponseExtract) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputListResponseUnion) AsRoute() (v OutputListResponseRoute) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputListResponseUnion) AsClassify() (v OutputListResponseClassify) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputListResponseUnion) AsSplitCollection() (v OutputListResponseSplitCollection) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputListResponseUnion) AsSplitItem() (v OutputListResponseSplitItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputListResponseUnion) AsError() (v ErrorEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputListResponseUnion) AsJoin() (v OutputListResponseJoin) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputListResponseUnion) AsEnrich() (v OutputListResponseEnrich) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputListResponseUnion) AsCollectionProcessing() (v OutputListResponseCollectionProcessing) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputListResponseUnion) AsSend() (v OutputListResponseSend) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OutputListResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *OutputListResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OutputListResponseUnionCorrectedContent is an implicit subunion of
// [OutputListResponseUnion]. OutputListResponseUnionCorrectedContent provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [OutputListResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAnyArray OfString OfFloat OfBool]
type OutputListResponseUnionCorrectedContent struct {
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

func (r *OutputListResponseUnionCorrectedContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OutputListResponseUnionInputs is an implicit subunion of
// [OutputListResponseUnion]. OutputListResponseUnionInputs provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [OutputListResponseUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfOutputListResponseTransformInputs
// OfOutputListResponseExtractInputs]
type OutputListResponseUnionInputs struct {
	// This field will be present if the value is a
	// [[]OutputListResponseTransformInput] instead of an object.
	OfOutputListResponseTransformInputs []OutputListResponseTransformInput `json:",inline"`
	// This field will be present if the value is a [[]OutputListResponseExtractInput]
	// instead of an object.
	OfOutputListResponseExtractInputs []OutputListResponseExtractInput `json:",inline"`
	JSON                              struct {
		OfOutputListResponseTransformInputs respjson.Field
		OfOutputListResponseExtractInputs   respjson.Field
		raw                                 string
	} `json:"-"`
}

func (r *OutputListResponseUnionInputs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OutputListResponseUnionMetadata is an implicit subunion of
// [OutputListResponseUnion]. OutputListResponseUnionMetadata provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [OutputListResponseUnion].
type OutputListResponseUnionMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	JSON                           struct {
		DurationFunctionToEventSeconds respjson.Field
		raw                            string
	} `json:"-"`
}

func (r *OutputListResponseUnionMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OutputListResponseUnionPrintPageOutput is an implicit subunion of
// [OutputListResponseUnion]. OutputListResponseUnionPrintPageOutput provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [OutputListResponseUnion].
type OutputListResponseUnionPrintPageOutput struct {
	ItemCount int64 `json:"itemCount"`
	// This field is from variant [OutputListResponseSplitCollectionPrintPageOutput].
	Items []OutputListResponseSplitCollectionPrintPageOutputItem `json:"items"`
	// This field is from variant [OutputListResponseSplitItemPrintPageOutput].
	CollectionReferenceID string `json:"collectionReferenceID"`
	// This field is from variant [OutputListResponseSplitItemPrintPageOutput].
	ItemOffset int64 `json:"itemOffset"`
	// This field is from variant [OutputListResponseSplitItemPrintPageOutput].
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

func (r *OutputListResponseUnionPrintPageOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OutputListResponseUnionSemanticPageOutput is an implicit subunion of
// [OutputListResponseUnion]. OutputListResponseUnionSemanticPageOutput provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [OutputListResponseUnion].
type OutputListResponseUnionSemanticPageOutput struct {
	ItemCount int64 `json:"itemCount"`
	// This field is from variant
	// [OutputListResponseSplitCollectionSemanticPageOutput].
	Items     []OutputListResponseSplitCollectionSemanticPageOutputItem `json:"items"`
	PageCount int64                                                     `json:"pageCount"`
	// This field is from variant [OutputListResponseSplitItemSemanticPageOutput].
	CollectionReferenceID string `json:"collectionReferenceID"`
	// This field is from variant [OutputListResponseSplitItemSemanticPageOutput].
	ItemClass string `json:"itemClass"`
	// This field is from variant [OutputListResponseSplitItemSemanticPageOutput].
	ItemClassCount int64 `json:"itemClassCount"`
	// This field is from variant [OutputListResponseSplitItemSemanticPageOutput].
	ItemClassOffset int64 `json:"itemClassOffset"`
	// This field is from variant [OutputListResponseSplitItemSemanticPageOutput].
	ItemOffset int64 `json:"itemOffset"`
	// This field is from variant [OutputListResponseSplitItemSemanticPageOutput].
	PageEnd int64 `json:"pageEnd"`
	// This field is from variant [OutputListResponseSplitItemSemanticPageOutput].
	PageStart int64 `json:"pageStart"`
	// This field is from variant [OutputListResponseSplitItemSemanticPageOutput].
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

func (r *OutputListResponseUnionSemanticPageOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseTransform struct {
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
	CorrectedContent OutputListResponseTransformCorrectedContentUnion `json:"correctedContent" api:"nullable"`
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
	Inputs []OutputListResponseTransformInput `json:"inputs" api:"nullable"`
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
	LastPublishErrorAt string                              `json:"lastPublishErrorAt" api:"nullable"`
	Metadata           OutputListResponseTransformMetadata `json:"metadata"`
	// Accuracy, precision, recall, and F1 score when corrected JSON is provided.
	Metrics OutputListResponseTransformMetrics `json:"metrics" api:"nullable"`
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
func (r OutputListResponseTransform) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseTransform) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OutputListResponseTransformCorrectedContentUnion contains all possible
// properties and values from [OutputListResponseTransformCorrectedContentOutput],
// [[]any], [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAnyArray OfString OfFloat OfBool]
type OutputListResponseTransformCorrectedContentUnion struct {
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field is from variant [OutputListResponseTransformCorrectedContentOutput].
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

func (u OutputListResponseTransformCorrectedContentUnion) AsOutputListResponseTransformCorrectedContentOutput() (v OutputListResponseTransformCorrectedContentOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputListResponseTransformCorrectedContentUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputListResponseTransformCorrectedContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputListResponseTransformCorrectedContentUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputListResponseTransformCorrectedContentUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OutputListResponseTransformCorrectedContentUnion) RawJSON() string { return u.JSON.raw }

func (r *OutputListResponseTransformCorrectedContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseTransformCorrectedContentOutput struct {
	Output []AnyTypeUnion `json:"output"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputListResponseTransformCorrectedContentOutput) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseTransformCorrectedContentOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseTransformInput struct {
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
func (r OutputListResponseTransformInput) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseTransformInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseTransformMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputListResponseTransformMetadata) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseTransformMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Accuracy, precision, recall, and F1 score when corrected JSON is provided.
type OutputListResponseTransformMetrics struct {
	Differences []OutputListResponseTransformMetricsDifference `json:"differences"`
	Metrics     OutputListResponseTransformMetricsMetrics      `json:"metrics"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Differences respjson.Field
		Metrics     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputListResponseTransformMetrics) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseTransformMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseTransformMetricsDifference struct {
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
func (r OutputListResponseTransformMetricsDifference) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseTransformMetricsDifference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseTransformMetricsMetrics struct {
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
func (r OutputListResponseTransformMetricsMetrics) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseTransformMetricsMetrics) UnmarshalJSON(data []byte) error {
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
type OutputListResponseExtract struct {
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
	CorrectedContent OutputListResponseExtractCorrectedContentUnion `json:"correctedContent" api:"nullable"`
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
	Inputs []OutputListResponseExtractInput `json:"inputs" api:"nullable"`
	// The input type of the content you're sending for transformation.
	//
	// Any of "csv", "docx", "email", "heic", "html", "jpeg", "json", "heif", "m4a",
	// "mp3", "pdf", "png", "text", "wav", "webp", "xls", "xlsx", "xml".
	InputType string `json:"inputType"`
	// List of properties that were invalid in the input.
	InvalidProperties []string                          `json:"invalidProperties"`
	Metadata          OutputListResponseExtractMetadata `json:"metadata"`
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
func (r OutputListResponseExtract) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseExtract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OutputListResponseExtractCorrectedContentUnion contains all possible properties
// and values from [OutputListResponseExtractCorrectedContentOutput], [[]any],
// [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAnyArray OfString OfFloat OfBool]
type OutputListResponseExtractCorrectedContentUnion struct {
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field is from variant [OutputListResponseExtractCorrectedContentOutput].
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

func (u OutputListResponseExtractCorrectedContentUnion) AsOutputListResponseExtractCorrectedContentOutput() (v OutputListResponseExtractCorrectedContentOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputListResponseExtractCorrectedContentUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputListResponseExtractCorrectedContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputListResponseExtractCorrectedContentUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OutputListResponseExtractCorrectedContentUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OutputListResponseExtractCorrectedContentUnion) RawJSON() string { return u.JSON.raw }

func (r *OutputListResponseExtractCorrectedContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseExtractCorrectedContentOutput struct {
	Output []AnyTypeUnion `json:"output"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputListResponseExtractCorrectedContentOutput) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseExtractCorrectedContentOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseExtractInput struct {
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
func (r OutputListResponseExtractInput) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseExtractInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseExtractMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputListResponseExtractMetadata) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseExtractMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseRoute struct {
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
	InboundEmail InboundEmailEvent               `json:"inboundEmail"`
	Metadata     OutputListResponseRouteMetadata `json:"metadata"`
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
func (r OutputListResponseRoute) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseRoute) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseRouteMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputListResponseRouteMetadata) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseRouteMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseClassify struct {
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
	InboundEmail InboundEmailEvent                  `json:"inboundEmail"`
	Metadata     OutputListResponseClassifyMetadata `json:"metadata"`
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
func (r OutputListResponseClassify) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseClassify) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseClassifyMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputListResponseClassifyMetadata) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseClassifyMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseSplitCollection struct {
	// Unique ID generated by bem to identify the event.
	EventID string `json:"eventID" api:"required"`
	// Unique identifier of function that this event is associated with.
	FunctionID string `json:"functionID" api:"required"`
	// Unique name of function that this event is associated with.
	FunctionName string `json:"functionName" api:"required"`
	// Any of "print_page", "semantic_page".
	OutputType      string                                           `json:"outputType" api:"required"`
	PrintPageOutput OutputListResponseSplitCollectionPrintPageOutput `json:"printPageOutput" api:"required"`
	// The unique ID you use internally to refer to this data point, propagated from
	// the original function input.
	ReferenceID        string                                              `json:"referenceID" api:"required"`
	SemanticPageOutput OutputListResponseSplitCollectionSemanticPageOutput `json:"semanticPageOutput" api:"required"`
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
	InboundEmail InboundEmailEvent                         `json:"inboundEmail"`
	Metadata     OutputListResponseSplitCollectionMetadata `json:"metadata"`
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
func (r OutputListResponseSplitCollection) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseSplitCollection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseSplitCollectionPrintPageOutput struct {
	ItemCount int64                                                  `json:"itemCount"`
	Items     []OutputListResponseSplitCollectionPrintPageOutputItem `json:"items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemCount   respjson.Field
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputListResponseSplitCollectionPrintPageOutput) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseSplitCollectionPrintPageOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseSplitCollectionPrintPageOutputItem struct {
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
func (r OutputListResponseSplitCollectionPrintPageOutputItem) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseSplitCollectionPrintPageOutputItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseSplitCollectionSemanticPageOutput struct {
	ItemCount int64                                                     `json:"itemCount"`
	Items     []OutputListResponseSplitCollectionSemanticPageOutputItem `json:"items"`
	PageCount int64                                                     `json:"pageCount"`
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
func (r OutputListResponseSplitCollectionSemanticPageOutput) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseSplitCollectionSemanticPageOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseSplitCollectionSemanticPageOutputItem struct {
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
func (r OutputListResponseSplitCollectionSemanticPageOutputItem) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseSplitCollectionSemanticPageOutputItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseSplitCollectionMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputListResponseSplitCollectionMetadata) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseSplitCollectionMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseSplitItem struct {
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
	InboundEmail       InboundEmailEvent                             `json:"inboundEmail"`
	Metadata           OutputListResponseSplitItemMetadata           `json:"metadata"`
	PrintPageOutput    OutputListResponseSplitItemPrintPageOutput    `json:"printPageOutput"`
	SemanticPageOutput OutputListResponseSplitItemSemanticPageOutput `json:"semanticPageOutput"`
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
func (r OutputListResponseSplitItem) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseSplitItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseSplitItemMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputListResponseSplitItemMetadata) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseSplitItemMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseSplitItemPrintPageOutput struct {
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
func (r OutputListResponseSplitItemPrintPageOutput) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseSplitItemPrintPageOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseSplitItemSemanticPageOutput struct {
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
func (r OutputListResponseSplitItemSemanticPageOutput) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseSplitItemSemanticPageOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseJoin struct {
	// Unique ID generated by bem to identify the event.
	EventID string `json:"eventID" api:"required"`
	// Unique identifier of function that this event is associated with.
	FunctionID string `json:"functionID" api:"required"`
	// Unique name of function that this event is associated with.
	FunctionName string `json:"functionName" api:"required"`
	// List of properties that were invalid in the input.
	InvalidProperties []string `json:"invalidProperties" api:"required"`
	// The items that were joined.
	Items []OutputListResponseJoinItem `json:"items" api:"required"`
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
	InboundEmail InboundEmailEvent              `json:"inboundEmail"`
	Metadata     OutputListResponseJoinMetadata `json:"metadata"`
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
func (r OutputListResponseJoin) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseJoin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseJoinItem struct {
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
func (r OutputListResponseJoinItem) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseJoinItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseJoinMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputListResponseJoinMetadata) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseJoinMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseEnrich struct {
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
	InboundEmail InboundEmailEvent                `json:"inboundEmail"`
	Metadata     OutputListResponseEnrichMetadata `json:"metadata"`
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
func (r OutputListResponseEnrich) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseEnrich) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseEnrichMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputListResponseEnrichMetadata) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseEnrichMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseCollectionProcessing struct {
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
	InboundEmail InboundEmailEvent                              `json:"inboundEmail"`
	Metadata     OutputListResponseCollectionProcessingMetadata `json:"metadata"`
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
func (r OutputListResponseCollectionProcessing) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseCollectionProcessing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseCollectionProcessingMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputListResponseCollectionProcessingMetadata) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseCollectionProcessingMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseSend struct {
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
	GoogleDriveOutput OutputListResponseSendGoogleDriveOutput `json:"googleDriveOutput"`
	// The inbound email that triggered this event.
	InboundEmail InboundEmailEvent              `json:"inboundEmail"`
	Metadata     OutputListResponseSendMetadata `json:"metadata"`
	// Metadata returned when a Send function delivers to an S3 bucket.
	S3Output OutputListResponseSendS3Output `json:"s3Output"`
	// Metadata returned when a Send function delivers to a webhook.
	WebhookOutput OutputListResponseSendWebhookOutput `json:"webhookOutput"`
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
func (r OutputListResponseSend) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseSend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata returned when a Send function delivers to Google Drive.
type OutputListResponseSendGoogleDriveOutput struct {
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
func (r OutputListResponseSendGoogleDriveOutput) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseSendGoogleDriveOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OutputListResponseSendMetadata struct {
	DurationFunctionToEventSeconds float64 `json:"durationFunctionToEventSeconds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationFunctionToEventSeconds respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OutputListResponseSendMetadata) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseSendMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata returned when a Send function delivers to an S3 bucket.
type OutputListResponseSendS3Output struct {
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
func (r OutputListResponseSendS3Output) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseSendS3Output) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata returned when a Send function delivers to a webhook.
type OutputListResponseSendWebhookOutput struct {
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
func (r OutputListResponseSendWebhookOutput) RawJSON() string { return r.JSON.raw }
func (r *OutputListResponseSendWebhookOutput) UnmarshalJSON(data []byte) error {
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
