// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bem

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
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
// EvalScoreService contains methods and other services that help with interacting
// with the bem API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEvalScoreService] method instead.
type EvalScoreService struct {
	options []option.RequestOption
}

// NewEvalScoreService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEvalScoreService(opts ...option.RequestOption) (r EvalScoreService) {
	r = EvalScoreService{}
	r.options = opts
	return
}

// **Score a function against a list of (input, expected) pairs.**
//
// Submits a batch of `(input, expected)` pairs, runs the named function over each
// input, and returns per-pair + aggregate accuracy metrics comparing the
// function's actual output to the provided expected JSON.
//
// Scoring runs asynchronously. The response carries a `scoreRunID`; poll
// `GET /v3/eval/score/{scoreRunID}` until `status` is one of `completed`, `error`,
// or `cancelled`.
//
// This request says only _what to extract_. How the output is compared against the
// expected value happens on the GET, recomputed from stored JSON each time.
func (r *EvalScoreService) New(ctx context.Context, body EvalScoreNewParams, opts ...option.RequestOption) (res *EvalScoreNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/eval/score"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// **Get the status and per-pair results of a score run.**
//
// The comparison happens here, not in the run: the function's output is compared
// against the expected value on every read, under the configuration supplied
// below. Re-reading the same run with different settings returns different metrics
// and costs nothing — no model calls are repeated.
//
// Comparison is exact and takes no configuration: a value matches the expected one
// or it is a miss. It is still redone on every read, so the numbers reflect the
// stored data as it is now.
//
// Returns `aggregate` once `status` reaches `completed` or `error`. `perPair` is
// populated incrementally — each pair's `fieldResults` appears as its underlying
// function call terminates.
func (r *EvalScoreService) Get(ctx context.Context, scoreRunID string, opts ...option.RequestOption) (res *EvalScoreRun, err error) {
	opts = slices.Concat(r.options, opts)
	if scoreRunID == "" {
		err = errors.New("missing required scoreRunID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/eval/score/%s", url.PathEscape(scoreRunID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// **Cancel an in-flight score run.**
//
// Transitions the run to `cancelled`. Function calls already in flight are allowed
// to finish (best-effort cancellation via the job queue); results from completed
// pairs may still appear in subsequent GETs.
func (r *EvalScoreService) Cancel(ctx context.Context, scoreRunID string, opts ...option.RequestOption) (res *EvalScoreRun, err error) {
	opts = slices.Concat(r.options, opts)
	if scoreRunID == "" {
		err = errors.New("missing required scoreRunID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/eval/score/%s/cancel", url.PathEscape(scoreRunID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Full status payload returned by `GET /v3/eval/score/{scoreRunID}`.
//
// Scoring takes no configuration: a value matches the expected one or it is a
// miss. The comparison is still recomputed on every read from the stored JSON, so
// the numbers reflect the data as it is now rather than as it was when the run
// executed.
type EvalScoreRun struct {
	FunctionName       string `json:"functionName" api:"required"`
	FunctionVersionNum int64  `json:"functionVersionNum" api:"required"`
	// Per-pair results. `fieldResults` appears once a pair has an output to compare.
	PerPair []EvalScoreRunPerPair `json:"perPair" api:"required"`
	// Counts across all pairs.
	Progress   EvalScoreRunProgress `json:"progress" api:"required"`
	ScoreRunID string               `json:"scoreRunID" api:"required"`
	// Status values for an eval-score run.
	//
	// Any of "pending", "initializing", "running", "completed", "error", "cancelled".
	Status EvalScoreRunStatus `json:"status" api:"required"`
	// Populated once `status` is `completed` or `error`.
	Aggregate EvalScoreRunAggregate `json:"aggregate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FunctionName       respjson.Field
		FunctionVersionNum respjson.Field
		PerPair            respjson.Field
		Progress           respjson.Field
		ScoreRunID         respjson.Field
		Status             respjson.Field
		Aggregate          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvalScoreRun) RawJSON() string { return r.JSON.raw }
func (r *EvalScoreRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-pair result.
type EvalScoreRunPerPair struct {
	PairIndex int64 `json:"pairIndex" api:"required"`
	// Per-pair status.
	//
	// Any of "pending", "running", "completed", "failed".
	Status string `json:"status" api:"required"`
	// The function call that produced the actual output, if any.
	CallID string `json:"callID"`
	// Error message if the underlying function call failed.
	ErrorMessage string `json:"errorMessage"`
	// Per-leaf comparator output. Present only after the pair has been compared.
	FieldResults []EvalScoreRunPerPairFieldResult `json:"fieldResults"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PairIndex    respjson.Field
		Status       respjson.Field
		CallID       respjson.Field
		ErrorMessage respjson.Field
		FieldResults respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvalScoreRunPerPair) RawJSON() string { return r.JSON.raw }
func (r *EvalScoreRunPerPair) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One leaf in `expected ∪ actual`.
type EvalScoreRunPerPairFieldResult struct {
	// Classification, in the same vocabulary the model-comparison endpoint reports.
	// Comparison is exact — a value matches or it does not:
	//
	// - `match`: both present and deep-equal
	// - `mismatch`: both present, different
	// - `missing`: expected present, actual absent
	// - `extra`: actual present, expected absent
	//
	// Any of "match", "mismatch", "missing", "extra".
	Match string `json:"match" api:"required"`
	// JSON Pointer to the leaf.
	Path   string `json:"path" api:"required"`
	Actual any    `json:"actual"`
	// Populated for every non-identical numeric pair; `actual - expected`. Reported as
	// evidence only — numbers have no threshold, so a delta tells you how far off a
	// value was without ever excusing it.
	Delta    float64 `json:"delta"`
	Expected any     `json:"expected"`
	// Populated for every non-identical string pair; the Levenshtein ratio in
	// `[0, 1]`. Reported as evidence: it says how close a wrong value was, which never
	// makes it right.
	Similarity float64 `json:"similarity"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Match       respjson.Field
		Path        respjson.Field
		Actual      respjson.Field
		Delta       respjson.Field
		Expected    respjson.Field
		Similarity  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvalScoreRunPerPairFieldResult) RawJSON() string { return r.JSON.raw }
func (r *EvalScoreRunPerPairFieldResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Counts across all pairs.
type EvalScoreRunProgress struct {
	Completed int64 `json:"completed" api:"required"`
	Failed    int64 `json:"failed" api:"required"`
	Total     int64 `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Completed   respjson.Field
		Failed      respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvalScoreRunProgress) RawJSON() string { return r.JSON.raw }
func (r *EvalScoreRunProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Populated once `status` is `completed` or `error`.
type EvalScoreRunAggregate struct {
	Extras              int64   `json:"extras" api:"required"`
	F1                  float64 `json:"f1" api:"required"`
	Matches             int64   `json:"matches" api:"required"`
	Mismatches          int64   `json:"mismatches" api:"required"`
	Missing             int64   `json:"missing" api:"required"`
	Precision           float64 `json:"precision" api:"required"`
	Recall              float64 `json:"recall" api:"required"`
	TotalFieldsActual   int64   `json:"totalFieldsActual" api:"required"`
	TotalFieldsExpected int64   `json:"totalFieldsExpected" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Extras              respjson.Field
		F1                  respjson.Field
		Matches             respjson.Field
		Mismatches          respjson.Field
		Missing             respjson.Field
		Precision           respjson.Field
		Recall              respjson.Field
		TotalFieldsActual   respjson.Field
		TotalFieldsExpected respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvalScoreRunAggregate) RawJSON() string { return r.JSON.raw }
func (r *EvalScoreRunAggregate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status values for an eval-score run.
type EvalScoreRunStatus string

const (
	EvalScoreRunStatusPending      EvalScoreRunStatus = "pending"
	EvalScoreRunStatusInitializing EvalScoreRunStatus = "initializing"
	EvalScoreRunStatusRunning      EvalScoreRunStatus = "running"
	EvalScoreRunStatusCompleted    EvalScoreRunStatus = "completed"
	EvalScoreRunStatusError        EvalScoreRunStatus = "error"
	EvalScoreRunStatusCancelled    EvalScoreRunStatus = "cancelled"
)

// A single file input with base64-encoded content.
//
// When using the Bem CLI, use `@path/to/file` in the `inputContent` field to
// automatically read and base64-encode the file:
// `--input.single-file '{"inputContent": "@file.pdf", "inputType": "pdf"}' --wait`
//
// The properties InputContent, InputType are required.
type FileInputParam struct {
	// Base64-encoded file content. In the Bem CLI, use `@path/to/file` to embed file
	// contents automatically.
	InputContent string `json:"inputContent" api:"required"`
	// The input type of the content you're sending for transformation.
	//
	// Must match the actual file format. See `InputType` for allowed values.
	//
	// Any of "csv", "docx", "email", "heic", "html", "jfif", "jpeg", "json", "heif",
	// "m4a", "mov", "mp3", "mp4", "pdf", "png", "pptx", "text", "wav", "webp", "xls",
	// "xlsx", "xml".
	InputType InputType `json:"inputType,omitzero" api:"required"`
	paramObj
}

func (r FileInputParam) MarshalJSON() (data []byte, err error) {
	type shadow FileInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Returned by `POST /v3/eval/score`.
type EvalScoreNewResponse struct {
	// Run identifier. Use with `GET /v3/eval/score/{scoreRunID}`.
	ScoreRunID string `json:"scoreRunID" api:"required"`
	// Status values for an eval-score run.
	//
	// Any of "pending", "initializing", "running", "completed", "error", "cancelled".
	Status EvalScoreRunStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ScoreRunID  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvalScoreNewResponse) RawJSON() string { return r.JSON.raw }
func (r *EvalScoreNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvalScoreNewParams struct {
	// Name of the function to score. Must be of type extract, transform, or analyze.
	FunctionName string `json:"functionName" api:"required"`
	// A saved Golden Data Set (`gds_…`) to score against. Mutually exclusive with
	// `pairs`; provide exactly one. Its input / corrected / schema columns are
	// resolved by column role. When it carries a `schema`-role column, scoring types
	// each row against that ground-truth schema instead of the function's own schema —
	// so results hold up as functions/schemas evolve.
	DatasetID param.Opt[string] `json:"datasetID,omitzero"`
	// Optional version number to score against. P0: only the function's current
	// version is accepted; passing a different version returns 422.
	FunctionVersionNum param.Opt[int64] `json:"functionVersionNum,omitzero"`
	// Inline `(input, expected)` pairs to score, up to 1000 per request. Mutually
	// exclusive with `datasetID`; provide exactly one.
	Pairs []EvalScoreNewParamsPair `json:"pairs,omitzero"`
	paramObj
}

func (r EvalScoreNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EvalScoreNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EvalScoreNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One `(input, expected)` pair.
//
// The properties Expected, Input are required.
type EvalScoreNewParamsPair struct {
	// Expected output for this input, as a JSON value. The comparator walks
	// `expected ∪ actual` and produces a per-leaf classification.
	Expected any `json:"expected,omitzero" api:"required"`
	// The file input to feed into the function.
	Input FileInputParam `json:"input,omitzero" api:"required"`
	paramObj
}

func (r EvalScoreNewParamsPair) MarshalJSON() (data []byte, err error) {
	type shadow EvalScoreNewParamsPair
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EvalScoreNewParamsPair) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
