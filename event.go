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
	"github.com/bem-team/bem-go-sdk/internal/requestconfig"
	"github.com/bem-team/bem-go-sdk/option"
	"github.com/bem-team/bem-go-sdk/packages/param"
	"github.com/bem-team/bem-go-sdk/packages/respjson"
)

// Submit training corrections for `extract`, `classify`, and `join` events.
//
// Feedback is event-centric — each correction is attached to an event by its
// `eventID`, and the server resolves the correct underlying storage (extract/join
// transformations or classify route events) from the event's function type.
//
// Split and enrich function types do not support feedback.
//
// EventService contains methods and other services that help with interacting with
// the bem API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventService] method instead.
type EventService struct {
	options []option.RequestOption
}

// NewEventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEventService(opts ...option.RequestOption) (r EventService) {
	r = EventService{}
	r.options = opts
	return
}

// **Submit a correction for an event.**
//
// Accepts training corrections for `extract`, `classify`, and `join` events. For
// extract/join events, `correction` is a JSON object matching the function's
// output schema. For classify events, `correction` is a JSON string matching one
// of the function version's declared classifications.
//
// Submitting feedback again for the same event overwrites the previous correction.
//
// Unsupported function types (split, enrich) return `400`.
func (r *EventService) SubmitFeedback(ctx context.Context, eventID string, body EventSubmitFeedbackParams, opts ...option.RequestOption) (res *EventSubmitFeedbackResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if eventID == "" {
		err = errors.New("missing required eventID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/events/%s/feedback", url.PathEscape(eventID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Echoed response after a correction is recorded.
type EventSubmitFeedbackResponse struct {
	Correction any `json:"correction" api:"required"`
	// Server timestamp when the correction was persisted (RFC 3339).
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	EventID   string    `json:"eventID" api:"required"`
	// Function types that support feedback submission.
	//
	// Any of "extract", "classify", "join".
	FunctionType EventSubmitFeedbackResponseFunctionType `json:"functionType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Correction   respjson.Field
		CreatedAt    respjson.Field
		EventID      respjson.Field
		FunctionType respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventSubmitFeedbackResponse) RawJSON() string { return r.JSON.raw }
func (r *EventSubmitFeedbackResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function types that support feedback submission.
type EventSubmitFeedbackResponseFunctionType string

const (
	EventSubmitFeedbackResponseFunctionTypeExtract  EventSubmitFeedbackResponseFunctionType = "extract"
	EventSubmitFeedbackResponseFunctionTypeClassify EventSubmitFeedbackResponseFunctionType = "classify"
	EventSubmitFeedbackResponseFunctionTypeJoin     EventSubmitFeedbackResponseFunctionType = "join"
)

type EventSubmitFeedbackParams struct {
	Correction    any             `json:"correction,omitzero" api:"required"`
	OrderMatching param.Opt[bool] `json:"orderMatching,omitzero"`
	paramObj
}

func (r EventSubmitFeedbackParams) MarshalJSON() (data []byte, err error) {
	type shadow EventSubmitFeedbackParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EventSubmitFeedbackParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
