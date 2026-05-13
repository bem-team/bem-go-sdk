// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bem

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/bem-team/bem-go-sdk/internal/apijson"
	"github.com/bem-team/bem-go-sdk/internal/apiquery"
	"github.com/bem-team/bem-go-sdk/internal/requestconfig"
	"github.com/bem-team/bem-go-sdk/option"
	"github.com/bem-team/bem-go-sdk/packages/param"
	"github.com/bem-team/bem-go-sdk/packages/respjson"
)

// Monitor, evaluate, and iterate on the quality of every function in your
// environment. Function Accuracy bundles two complementary loops:
//
// ## Evaluations (`/v3/eval`)
//
// Trigger and retrieve per-transformation evaluations. Evaluations run
// asynchronously and score each transformation's output against the function's
// schema for confidence, per-field hallucination detection, and relevance.
// Supported for `extract`, `transform`, `analyze`, and `join` events.
//
//  1. **Trigger** — `POST /v3/eval` queues jobs for a batch of transformation IDs.
//  2. **Poll** — `GET /v3/eval/results` returns the current state of each requested
//     ID, partitioned into `results`, `pending`, and `failed`. Accepts either
//     `eventIDs` (preferred) or `transformationIDs` as a comma-separated query
//     parameter, and always keys the response by event KSUID.
//
// Up to 100 IDs may be submitted per request.
//
// ## Metrics, review, regression (`/v3/functions/{metrics,review,regression,compare}`)
//
// Roll evaluation results and user corrections up into actionable function-level
// signal:
//
//   - **`GET /v3/functions/metrics`** — aggregate accuracy, precision, recall, F1,
//     and confusion-matrix counts per function.
//   - **`POST /v3/functions/review`** — sample-size estimation, confidence-bucketed
//     distribution, PR-AUC, and per-threshold confidence intervals (Wald or Wilson)
//     for picking review cutoffs.
//   - **`POST /v3/functions/regression`** — replay corrected historical inputs
//     against a new function version, producing a labeled regression dataset.
//   - **`POST /v3/functions/regression/corrections`** — propagate baseline
//     corrections onto the regression dataset so it can be scored.
//   - **`POST /v3/functions/compare`** — compute aggregate and field-level lift
//     between any two versions, optionally scoped to the regression dataset.
//
// All five endpoints support `extract` end-to-end on both the vision and OCR
// paths, alongside the legacy `transform` / `analyze` / `join` types.
//
// EvalResultService contains methods and other services that help with interacting
// with the bem API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEvalResultService] method instead.
type EvalResultService struct {
	options []option.RequestOption
}

// NewEvalResultService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEvalResultService(opts ...option.RequestOption) (r EvalResultService) {
	r = EvalResultService{}
	r.options = opts
	return
}

// **Fetch evaluation results for a batch of events.**
//
// Pass either `eventIDs` (preferred — the externally-stable V3 identifier) or
// `transformationIDs` as a comma-separated query parameter. Exactly one of the two
// must be provided. Up to 100 IDs per request.
//
// For each requested ID the response reports one of three states: a completed
// `result`, still-`pending`, or `failed`. Results, pending, and failed entries are
// all keyed by event KSUID regardless of which input form was used.
func (r *EvalResultService) GetResults(ctx context.Context, query EvalResultGetResultsParams, opts ...option.RequestOption) (res *EvaluationResults, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/eval/results"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Batched response containing the evaluation state for every requested ID,
// partitioned into completed `results`, still-running `pending`, and terminal
// `failed` groups. All identifiers in the response are event KSUIDs regardless of
// whether the request used `eventIDs` or `transformationIDs`.
type EvaluationResults struct {
	// Completed evaluation results, keyed by event KSUID.
	//
	// An event appears here only if its evaluation completed successfully.
	// Still-running evaluations appear in `pending`; failed evaluations appear in
	// `failed`.
	Results any `json:"results" api:"required"`
	// Reserved map of event KSUID to error message for validation failures on the
	// request itself. Populated only in edge cases.
	Errors any `json:"errors"`
	// Events whose evaluation failed or was not found.
	Failed []EvaluationResultsFailed `json:"failed"`
	// Events whose evaluation is still running.
	Pending []EvaluationResultsPending `json:"pending"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Errors      respjson.Field
		Failed      respjson.Field
		Pending     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvaluationResults) RawJSON() string { return r.JSON.raw }
func (r *EvaluationResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An event whose evaluation failed or was not found.
type EvaluationResultsFailed struct {
	// Server timestamp associated with the failure.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Human-readable failure reason.
	ErrorMessage string `json:"errorMessage" api:"required"`
	// Event KSUID.
	EventID string `json:"eventID" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt    respjson.Field
		ErrorMessage respjson.Field
		EventID      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvaluationResultsFailed) RawJSON() string { return r.JSON.raw }
func (r *EvaluationResultsFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An event whose evaluation is still running.
type EvaluationResultsPending struct {
	// Server timestamp when the evaluation was queued.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Event KSUID.
	EventID string `json:"eventID" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		EventID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvaluationResultsPending) RawJSON() string { return r.JSON.raw }
func (r *EvaluationResultsPending) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalResultGetResultsParams struct {
	// Optional evaluation version filter.
	EvaluationVersion param.Opt[string] `query:"evaluationVersion,omitzero" json:"-"`
	// Comma-separated list of event KSUIDs to fetch results for. Between 1 and 100 IDs
	// per request. Mutually exclusive with `transformationIDs`.
	EventIDs param.Opt[string] `query:"eventIDs,omitzero" json:"-"`
	// Comma-separated list of transformation IDs to fetch results for. Between 1 and
	// 100 IDs per request. Mutually exclusive with `eventIDs`. Prefer `eventIDs` for
	// new integrations.
	TransformationIDs param.Opt[string] `query:"transformationIDs,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EvalResultGetResultsParams]'s query parameters as
// `url.Values`.
func (r EvalResultGetResultsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
