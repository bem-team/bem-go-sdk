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

// **Fetch evaluation results for a batch of transformations (POST).**
//
// For each requested transformation ID the response reports one of three states: a
// completed `result`, still-`pending`, or `failed`. The POST variant accepts the
// ID list in the request body; use the `GET` variant with query parameters for
// simpler clients.
func (r *EvalResultService) FetchResults(ctx context.Context, body EvalResultFetchResultsParams, opts ...option.RequestOption) (res *EvalResultFetchResultsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/eval/results"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// **Fetch evaluation results for a batch of transformations.**
//
// Identical behavior to the POST variant; accepts transformation IDs as a
// comma-separated `transformationIDs` query parameter. Limited to 100 IDs per
// request.
func (r *EvalResultService) GetResults(ctx context.Context, query EvalResultGetResultsParams, opts ...option.RequestOption) (res *EvalResultGetResultsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/eval/results"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Batched response containing the evaluation state for every requested
// transformation ID, partitioned into completed `results`, still-running
// `pending`, and terminal `failed` groups.
type EvalResultFetchResultsResponse struct {
	// Completed evaluation results, keyed by transformation ID.
	//
	// A transformation appears here only if its evaluation completed successfully.
	// Still-running evaluations appear in `pending`; failed evaluations appear in
	// `failed`.
	Results any `json:"results" api:"required"`
	// Reserved map of transformation ID to error message for validation failures on
	// the request itself. Populated only in edge cases.
	Errors any `json:"errors"`
	// Transformations whose evaluation failed or was not found.
	Failed []EvalResultFetchResultsResponseFailed `json:"failed"`
	// Transformations whose evaluation is still running.
	Pending []EvalResultFetchResultsResponsePending `json:"pending"`
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
func (r EvalResultFetchResultsResponse) RawJSON() string { return r.JSON.raw }
func (r *EvalResultFetchResultsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A transformation whose evaluation failed or was not found.
type EvalResultFetchResultsResponseFailed struct {
	// Server timestamp associated with the failure.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Human-readable failure reason.
	ErrorMessage     string `json:"errorMessage" api:"required"`
	TransformationID string `json:"transformationId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt        respjson.Field
		ErrorMessage     respjson.Field
		TransformationID respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvalResultFetchResultsResponseFailed) RawJSON() string { return r.JSON.raw }
func (r *EvalResultFetchResultsResponseFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A transformation whose evaluation is still running.
type EvalResultFetchResultsResponsePending struct {
	// Server timestamp when the evaluation was queued.
	CreatedAt        time.Time `json:"createdAt" api:"required" format:"date-time"`
	TransformationID string    `json:"transformationId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt        respjson.Field
		TransformationID respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvalResultFetchResultsResponsePending) RawJSON() string { return r.JSON.raw }
func (r *EvalResultFetchResultsResponsePending) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Batched response containing the evaluation state for every requested
// transformation ID, partitioned into completed `results`, still-running
// `pending`, and terminal `failed` groups.
type EvalResultGetResultsResponse struct {
	// Completed evaluation results, keyed by transformation ID.
	//
	// A transformation appears here only if its evaluation completed successfully.
	// Still-running evaluations appear in `pending`; failed evaluations appear in
	// `failed`.
	Results any `json:"results" api:"required"`
	// Reserved map of transformation ID to error message for validation failures on
	// the request itself. Populated only in edge cases.
	Errors any `json:"errors"`
	// Transformations whose evaluation failed or was not found.
	Failed []EvalResultGetResultsResponseFailed `json:"failed"`
	// Transformations whose evaluation is still running.
	Pending []EvalResultGetResultsResponsePending `json:"pending"`
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
func (r EvalResultGetResultsResponse) RawJSON() string { return r.JSON.raw }
func (r *EvalResultGetResultsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A transformation whose evaluation failed or was not found.
type EvalResultGetResultsResponseFailed struct {
	// Server timestamp associated with the failure.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Human-readable failure reason.
	ErrorMessage     string `json:"errorMessage" api:"required"`
	TransformationID string `json:"transformationId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt        respjson.Field
		ErrorMessage     respjson.Field
		TransformationID respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvalResultGetResultsResponseFailed) RawJSON() string { return r.JSON.raw }
func (r *EvalResultGetResultsResponseFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A transformation whose evaluation is still running.
type EvalResultGetResultsResponsePending struct {
	// Server timestamp when the evaluation was queued.
	CreatedAt        time.Time `json:"createdAt" api:"required" format:"date-time"`
	TransformationID string    `json:"transformationId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt        respjson.Field
		TransformationID respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvalResultGetResultsResponsePending) RawJSON() string { return r.JSON.raw }
func (r *EvalResultGetResultsResponsePending) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalResultFetchResultsParams struct {
	// Transformation IDs to fetch results for. Up to 100 per request.
	TransformationIDs []string `json:"transformationIDs,omitzero" api:"required"`
	// Optional evaluation version filter.
	EvaluationVersion param.Opt[string] `json:"evaluationVersion,omitzero"`
	paramObj
}

func (r EvalResultFetchResultsParams) MarshalJSON() (data []byte, err error) {
	type shadow EvalResultFetchResultsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EvalResultFetchResultsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalResultGetResultsParams struct {
	// Comma-separated list of transformation IDs to fetch results for. Between 1 and
	// 100 IDs per request.
	TransformationIDs string `query:"transformationIDs" api:"required" json:"-"`
	// Optional evaluation version filter.
	EvaluationVersion param.Opt[string] `query:"evaluationVersion,omitzero" json:"-"`
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
