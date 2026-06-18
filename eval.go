// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bem

import (
	"context"
	"net/http"
	"slices"

	"github.com/bem-team/bem-go-sdk/internal/apijson"
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
// EvalService contains methods and other services that help with interacting with
// the bem API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEvalService] method instead.
type EvalService struct {
	options []option.RequestOption
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
	Results EvalResultService
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
	Score EvalScoreService
}

// NewEvalService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEvalService(opts ...option.RequestOption) (r EvalService) {
	r = EvalService{}
	r.options = opts
	r.Results = NewEvalResultService(opts...)
	r.Score = NewEvalScoreService(opts...)
	return
}

// **Queue evaluation jobs for a batch of transformations.**
//
// Evaluations run asynchronously and score each transformation's output against
// the function's schema for confidence, hallucination detection, and relevance.
// Transformations must belong to events of a supported type: `extract`,
// `transform`, `analyze`, or `join`.
//
// Returns immediately with a summary of queued vs. skipped transformations and
// per-transformation errors. Poll `GET /v3/eval/results` to retrieve results once
// evaluations complete.
func (r *EvalService) TriggerEvaluation(ctx context.Context, body EvalTriggerEvaluationParams, opts ...option.RequestOption) (res *EvalTriggerEvaluationResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/eval"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Summary of the trigger call. Evaluations run asynchronously; use
// `GET /v3/eval/results` to poll for results.
type EvalTriggerEvaluationResponse struct {
	// Number of evaluation jobs newly queued.
	Queued int64 `json:"queued" api:"required"`
	// Number of transformations skipped because an evaluation job was already pending
	// or already completed for them.
	Skipped int64 `json:"skipped" api:"required"`
	// Map of transformation ID to human-readable error message for any transformations
	// that could not be queued (e.g. not found, unsupported event type).
	Errors any `json:"errors"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Queued      respjson.Field
		Skipped     respjson.Field
		Errors      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvalTriggerEvaluationResponse) RawJSON() string { return r.JSON.raw }
func (r *EvalTriggerEvaluationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalTriggerEvaluationParams struct {
	// Transformation IDs to evaluate. Up to 100 per request.
	TransformationIDs []string `json:"transformationIDs,omitzero" api:"required"`
	// Optional evaluation version (e.g. `0.1.0-gemini`). When omitted the server's
	// default evaluation version is used.
	EvaluationVersion param.Opt[string] `json:"evaluationVersion,omitzero"`
	paramObj
}

func (r EvalTriggerEvaluationParams) MarshalJSON() (data []byte, err error) {
	type shadow EvalTriggerEvaluationParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EvalTriggerEvaluationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
