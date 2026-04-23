// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bem

import (
	"context"
	"net/http"
	"os"
	"slices"

	"github.com/bem-team/bem-go-sdk/internal/requestconfig"
	"github.com/bem-team/bem-go-sdk/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the bem API. You should not instantiate this client directly,
// and instead use the [NewClient] method instead.
type Client struct {
	options []option.RequestOption
	// Functions are the core building blocks of data transformation in Bem. Each
	// function type serves a specific purpose:
	//
	//   - **Extract**: Extract structured JSON data from unstructured documents (PDFs,
	//     emails, images, spreadsheets), with optional layout-aware bounding-box
	//     extraction
	//   - **Route**: Direct data to different processing paths based on conditions
	//   - **Split**: Break multi-page documents into individual pages for parallel
	//     processing
	//   - **Join**: Combine outputs from multiple function calls into a single result
	//   - **Payload Shaping**: Transform and restructure data using JMESPath expressions
	//   - **Enrich**: Enhance data with semantic search against collections
	//   - **Send**: Deliver workflow outputs to downstream destinations
	//
	// Use these endpoints to create, update, list, and manage your functions.
	Functions FunctionService
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
	Calls CallService
	// Retrieve terminal error events from workflow calls.
	//
	// Errors are events produced by function steps that failed during processing. A
	// single workflow call may produce multiple error events if several steps fail
	// independently.
	//
	// Errors and outputs from the same call are not mutually exclusive: a
	// partially-completed workflow may have both.
	//
	// Use `GET /v3/errors` to list errors across calls, or `GET /v3/errors/{eventID}`
	// to retrieve a specific error. To get errors scoped to a single call, filter by
	// `callIDs`.
	Errors ErrorService
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
	Outputs OutputService
	// Workflows orchestrate one or more functions into a directed acyclic graph (DAG)
	// for document processing.
	//
	// Use these endpoints to create, update, list, and manage workflows, and to invoke
	// them with file input via `POST /v3/workflows/{workflowName}/call`.
	//
	// The call endpoint accepts files as either multipart form data or JSON with
	// base64-encoded content. In the Bem CLI, use `@path/to/file` inside JSON values
	// to automatically read and encode files:
	//
	// ```
	//
	//	bem workflows call --workflow-name my-workflow \
	//	  --input.single-file '{"inputContent": "@file.pdf", "inputType": "pdf"}' \
	//	  --wait
	//
	// ```
	Workflows WorkflowService
	// Infer JSON Schemas from uploaded documents using AI.
	//
	// Upload a file (PDF, image, spreadsheet, email, etc.) and receive a
	// general-purpose JSON Schema that captures the document's structure. The inferred
	// schema can be used directly as the `outputSchema` when creating Extract
	// functions.
	//
	// The schema is designed to be broadly applicable to documents of the same type,
	// not just the specific file uploaded.
	InferSchema InferSchemaService
	// Collections are named groups of embedded items used by Enrich functions for
	// semantic search.
	//
	// Each collection is referenced by a `collectionName`, which supports dot notation
	// for hierarchical paths (e.g. `customers.premium.vip`). Names must contain only
	// letters, digits, underscores, and dots, and each segment must start with a
	// letter or underscore.
	//
	// ## Items
	//
	// Items carry either a string or a JSON object in their `data` field. When items
	// are added or updated, their `data` is embedded asynchronously —
	// `POST /v3/collections/items` and `PUT /v3/collections/items` return immediately
	// with a `pending` status and an `eventID` that can be correlated with webhook
	// notifications once processing completes.
	//
	// ## Listing and hierarchy
	//
	// Use `GET /v3/collections` with `parentCollectionName` to list collections under
	// a path, or `collectionNameSearch` for a case-insensitive substring match.
	// `GET /v3/collections/items` retrieves a specific collection's items; pass
	// `includeSubcollections=true` to fold in items from all descendant collections.
	//
	// ## Token counting
	//
	// Use `POST /v3/collections/token-count` to check whether texts fit within the
	// embedding model's 8,192-token-per-text limit before submitting them for
	// embedding.
	Collections CollectionService
	// Submit training corrections for `extract`, `classify`, and `join` events.
	//
	// Feedback is event-centric — each correction is attached to an event by its
	// `eventID`, and the server resolves the correct underlying storage (extract/join
	// transformations or classify route events) from the event's function type.
	//
	// Split and enrich function types do not support feedback.
	Events EventService
	// Manage the webhook signing secret used to authenticate outbound webhook
	// deliveries.
	//
	// When a signing secret is active, every webhook delivery includes a
	// `bem-signature` header in the format `t={unix_timestamp},v1={hex_hmac_sha256}`.
	// The signature covers `{timestamp}.{raw_request_body}` and can be verified using
	// HMAC-SHA256 with your secret.
	//
	// Rotate the secret at any time with `POST /v3/webhook-secret`. To avoid downtime
	// during rotation, update your verification logic to accept both the old and new
	// secret briefly before revoking the old one.
	WebhookSecret WebhookSecretService
}

// DefaultClientOptions read from the environment (BEM_API_KEY, BEM_BASE_URL). This
// should be used to initialize new clients.
func DefaultClientOptions() []option.RequestOption {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("BEM_BASE_URL"); ok {
		defaults = append(defaults, option.WithBaseURL(o))
	}
	if o, ok := os.LookupEnv("BEM_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	return defaults
}

// NewClient generates a new client with the default option read from the
// environment (BEM_API_KEY, BEM_BASE_URL). The option passed in as arguments are
// applied after these default arguments, and all option will be passed down to the
// services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r Client) {
	opts = append(DefaultClientOptions(), opts...)

	r = Client{options: opts}

	r.Functions = NewFunctionService(opts...)
	r.Calls = NewCallService(opts...)
	r.Errors = NewErrorService(opts...)
	r.Outputs = NewOutputService(opts...)
	r.Workflows = NewWorkflowService(opts...)
	r.InferSchema = NewInferSchemaService(opts...)
	r.Collections = NewCollectionService(opts...)
	r.Events = NewEventService(opts...)
	r.WebhookSecret = NewWebhookSecretService(opts...)

	return
}

// Execute makes a request with the given context, method, URL, request params,
// response, and request options. This is useful for hitting undocumented endpoints
// while retaining the base URL, auth, retries, and other options from the client.
//
// If a byte slice or an [io.Reader] is supplied to params, it will be used as-is
// for the request body.
//
// The params is by default serialized into the body using [encoding/json]. If your
// type implements a MarshalJSON function, it will be used instead to serialize the
// request. If a URLQuery method is implemented, the returned [url.Values] will be
// used as query strings to the url.
//
// If your params struct uses [param.Field], you must provide either [MarshalJSON],
// [URLQuery], and/or [MarshalForm] functions. It is undefined behavior to use a
// struct uses [param.Field] without specifying how it is serialized.
//
// Any "…Params" object defined in this library can be used as the request
// argument. Note that 'path' arguments will not be forwarded into the url.
//
// The response body will be deserialized into the res variable, depending on its
// type:
//
//   - A pointer to a [*http.Response] is populated by the raw response.
//   - A pointer to a byte array will be populated with the contents of the request
//     body.
//   - A pointer to any other type uses this library's default JSON decoding, which
//     respects UnmarshalJSON if it is defined on the type.
//   - A nil value will not read the response body.
//
// For even greater flexibility, see [option.WithResponseInto] and
// [option.WithResponseBodyInto].
func (r *Client) Execute(ctx context.Context, method string, path string, params any, res any, opts ...option.RequestOption) error {
	opts = slices.Concat(r.options, opts)
	return requestconfig.ExecuteNewRequest(ctx, method, path, params, res, opts...)
}

// Get makes a GET request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Get(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodGet, path, params, res, opts...)
}

// Post makes a POST request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Post(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPost, path, params, res, opts...)
}

// Put makes a PUT request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Put(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPut, path, params, res, opts...)
}

// Patch makes a PATCH request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Patch(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPatch, path, params, res, opts...)
}

// Delete makes a DELETE request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Delete(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodDelete, path, params, res, opts...)
}
