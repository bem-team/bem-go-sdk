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
	//   - **Transform**: Extract structured JSON data from unstructured documents (PDFs,
	//     emails, images)
	//   - **Analyze**: Perform visual analysis on documents to extract layout-aware
	//     information
	//   - **Route**: Direct data to different processing paths based on conditions
	//   - **Split**: Break multi-page documents into individual pages for parallel
	//     processing
	//   - **Join**: Combine outputs from multiple function calls into a single result
	//   - **Payload Shaping**: Transform and restructure data using JMESPath expressions
	//   - **Enrich**: Enhance data with semantic search against collections
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
	// Workflow operations
	Workflows WorkflowService
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
