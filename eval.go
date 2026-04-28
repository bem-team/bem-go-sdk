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

// Trigger and retrieve evaluations for completed transformations.
//
// Evaluations run asynchronously and score each transformation's output against
// the function's schema for confidence, per-field hallucination detection, and
// relevance. Evaluations are supported for `extract`, `transform`, `analyze`, and
// `join` events.
//
// ## Lifecycle
//
//  1. **Trigger** — `POST /v3/eval` queues jobs for a batch of transformation IDs
//     and returns immediately with `queued` / `skipped` counts plus per-ID errors.
//  2. **Poll** — `POST /v3/eval/results` (body) or `GET /v3/eval/results` (query)
//     returns the current state of each requested transformation, partitioned into
//     `results` (completed), `pending` (still running), and `failed` (terminal
//     failures or unknown transformation IDs).
//
// Up to 100 transformation IDs may be submitted per request.
//
// EvalService contains methods and other services that help with interacting with
// the bem API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEvalService] method instead.
type EvalService struct {
	options []option.RequestOption
	// Trigger and retrieve evaluations for completed transformations.
	//
	// Evaluations run asynchronously and score each transformation's output against
	// the function's schema for confidence, per-field hallucination detection, and
	// relevance. Evaluations are supported for `extract`, `transform`, `analyze`, and
	// `join` events.
	//
	// ## Lifecycle
	//
	//  1. **Trigger** — `POST /v3/eval` queues jobs for a batch of transformation IDs
	//     and returns immediately with `queued` / `skipped` counts plus per-ID errors.
	//  2. **Poll** — `POST /v3/eval/results` (body) or `GET /v3/eval/results` (query)
	//     returns the current state of each requested transformation, partitioned into
	//     `results` (completed), `pending` (still running), and `failed` (terminal
	//     failures or unknown transformation IDs).
	//
	// Up to 100 transformation IDs may be submitted per request.
	Results EvalResultService
}

// NewEvalService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEvalService(opts ...option.RequestOption) (r EvalService) {
	r = EvalService{}
	r.options = opts
	r.Results = NewEvalResultService(opts...)
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
// per-transformation errors. Poll `POST /v3/eval/results` or
// `GET /v3/eval/results` to retrieve results once evaluations complete.
func (r *EvalService) TriggerEvaluation(ctx context.Context, body EvalTriggerEvaluationParams, opts ...option.RequestOption) (res *EvalTriggerEvaluationResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/eval"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Summary of the trigger call. Evaluations run asynchronously; use
// `POST /v3/eval/results` or `GET /v3/eval/results` to poll for results.
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
