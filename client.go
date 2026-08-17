// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bem

import (
	"context"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/bem-team/bem-go-sdk/internal/requestconfig"
	"github.com/bem-team/bem-go-sdk/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the bem API. You should not instantiate this client directly,
// and instead use the [NewClient] method instead.
type Client struct {
	options   []option.RequestOption
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
	Events   EventService
	Webhooks WebhookService
	// bem POSTs a JSON event to your configured webhook URL each time a subscribed
	// function call, workflow output, or collection-processing job fires. This section
	// is the reference for those deliveries: the payload shape per event type, plus
	// the endpoints you use to manage the signing secret.
	//
	// Every variant shares the same envelope — function/workflow IDs, timestamps, the
	// inbound email that triggered the call, and so on — and adds a payload field that
	// depends on the function type. The `eventType` field on the body is the
	// discriminator: dispatch on it to select which payload shape to expect. SDKs
	// generated from this spec expose a `webhooks.unwrap()` helper that performs the
	// dispatch and returns a typed event.
	//
	// ## Payloads
	//
	// | `eventType`             | Payload                                                                      | Schema                      |
	// | ----------------------- | ---------------------------------------------------------------------------- | --------------------------- |
	// | `extract`               | [Extract event](/api/v3/webhooks/events/extract)                             | `ExtractEvent`              |
	// | `classify`              | [Classify event](/api/v3/webhooks/events/classify)                           | `ClassifyEvent`             |
	// | `parse`                 | [Parse event](/api/v3/webhooks/events/parse)                                 | `ParseEvent`                |
	// | `split_collection`      | [Split collection event](/api/v3/webhooks/events/split-collection)           | `SplitCollectionEvent`      |
	// | `split_item`            | [Split item event](/api/v3/webhooks/events/split-item)                       | `SplitItemEvent`            |
	// | `join`                  | [Join event](/api/v3/webhooks/events/join)                                   | `JoinEvent`                 |
	// | `enrich`                | [Enrich event](/api/v3/webhooks/events/enrich)                               | `EnrichEvent`               |
	// | `payload_shaping`       | [Payload shaping event](/api/v3/webhooks/events/payload-shaping)             | `PayloadShapingEvent`       |
	// | `send`                  | [Send event](/api/v3/webhooks/events/send)                                   | `SendEvent`                 |
	// | `evaluation`            | [Evaluation event](/api/v3/webhooks/events/evaluation)                       | `EvaluationEvent`           |
	// | `collection_processing` | [Collection processing event](/api/v3/webhooks/events/collection-processing) | `collectionProcessingEvent` |
	// | `error`                 | [Error event](/api/v3/webhooks/events/error)                                 | `ErrorEvent`                |
	//
	// ## Signing secret
	//
	// Every delivery includes a `bem-signature` header in the format
	// `t={unix_timestamp},v1={hex_hmac_sha256}`. The signature covers
	// `{timestamp}.{raw_request_body}` and is computed with HMAC-SHA256 using the
	// active signing secret for your environment.
	//
	// To verify a payload:
	//
	//  1. Parse `bem-signature: t={timestamp},v1={signature}`.
	//  2. Construct the signed string: `{timestamp}.{raw_request_body}`.
	//  3. Compute HMAC-SHA256 of that string using your secret.
	//  4. Reject the request if the hex digest doesn't match `v1`, or if the timestamp
	//     is more than a few minutes old.
	//
	// Manage the secret with these endpoints:
	//
	//   - [**Generate a signing secret**](/api/v3/webhooks/secret/generate-secret) —
	//     `POST /v3/webhook-secret`. Returns the new secret in full exactly once.
	//   - [**Get the signing secret**](/api/v3/webhooks/secret/get-secret) —
	//     `GET /v3/webhook-secret`. Returns the active secret.
	//   - [**Revoke the signing secret**](/api/v3/webhooks/secret/revoke-secret) —
	//     `DELETE /v3/webhook-secret`. Webhook deliveries continue but are unsigned
	//     until a new secret is generated.
	//
	// For zero-downtime rotation, briefly accept both the old and new secret in your
	// verification logic before revoking the old one.
	//
	// ## Retries
	//
	// bem treats any non-2XX response (or a transport failure) as a delivery error and
	// retries with exponential backoff. Return a 2XX as soon as you have durably
	// queued the payload — do not block on downstream work.
	WebhookSecret WebhookSecretService
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
	Eval EvalService
	// Unix-shell-style nav over parsed documents and the cross-doc memory store.
	//
	// `POST /v3/fs` is a single op-driven endpoint designed for LLM agents and
	// programmatic consumers that want to walk a corpus the way they'd walk a
	// filesystem.
	//
	// ## Doc-level ops (every parsed document)
	//
	// - `ls` — list parsed documents with rich per-doc metadata.
	// - `cat` — read one doc's parse JSON, sliced (`range`) or projected (`select`).
	// - `head` — first N sections of one doc.
	// - `grep` — substring or regex search; `scope`, `path`, `countOnly` available.
	// - `stat` — metadata only (page/section/entity counts, timestamps).
	//
	// ## Memory-level ops (require `linkAcrossDocuments: true` on the parse function)
	//
	// - `find` — list canonical entities across the corpus.
	// - `open` — entity + mentions.
	// - `xref` — for one entity, sections across docs that mention it (with content).
	//
	// Memory ops return an empty list with a `hint` when no docs in this environment
	// have been memory-linked.
	//
	// ## Pagination
	//
	// List ops paginate by cursor — pass the previous response's `nextCursor` back as
	// `cursor`; `hasMore: false` signals the last page. Same idiom as `/v3/calls` and
	// `/v3/outputs`.
	Fs FService
	// Connectors are integrations that trigger a Bem workflow from an external system.
	//
	// A connector binds an inbound source — currently Box or a Paragon-managed
	// integration such as Google Drive — to a specific workflow (by `workflowName` or
	// `workflowID`). When the source observes a new file, Bem invokes the bound
	// workflow against that file.
	//
	// Use these endpoints to create, list, and remove connectors. The fields used at
	// create time depend on the connector `type`: Box connectors require Box
	// credentials and a folder to watch, while Paragon connectors carry a
	// `paragonIntegration` identifier and an integration-specific
	// `paragonConfiguration` object (for example, `{ "folderId": "..." }` for Google
	// Drive).
	Connectors ConnectorService
	// Subscriptions wire up notifications for the events your functions and
	// collections produce.
	//
	// Most subscriptions target a single function (by `functionName` or `functionID`)
	// or a single collection (by `collectionName` or `collectionID`) and select a
	// `type` corresponding to the event you want to receive — for example `transform`,
	// `route`, `join`, `evaluation`, `error`, `enrich`, or `collection_processing`.
	//
	// Entity-lifecycle events are account-wide and target no function or collection.
	// Set `type` to one of the following and provide a `webhookURL` (these event types
	// support webhook delivery only):
	//
	//   - `entity_proposed` — an entity entered the `proposed` curation status (queued
	//     for review).
	//   - `entity_validated` — an entity was approved/validated by a reviewer.
	//   - `entity_rejected` — an entity was rejected by a reviewer.
	//
	// Each entity-lifecycle delivery is a JSON POST describing the transition
	// (`entityID`, `typeName`, `priorStatus`, `newStatus`, optional `actorUserID` and
	// `reason`, and a `timestamp`).
	//
	// Deliveries can be sent to any combination of:
	//
	// - `webhookURL` — HTTPS endpoint that receives a JSON POST per event.
	// - `s3Bucket` + `s3FilePath` — sync output JSON into an AWS S3 prefix you own.
	// - `googleDriveFolderID` — drop output JSON into a Google Drive folder.
	//
	// Use `disabled: true` to pause delivery without deleting the subscription.
	// Updates follow conventional PATCH semantics — only the fields you include are
	// changed.
	Subscriptions SubscriptionService
	// Views are tabular projections over the `transformations` your functions produce
	// — a saved query that turns raw extracted JSON into a filterable, paginatable,
	// aggregatable table.
	//
	// ## Anatomy
	//
	// A view declares:
	//
	//   - One or more **functions** to read from (by `functionID` or `functionName`).
	//   - A list of **columns**, each pinned to a `valueSchemaPath` (a JSON Pointer into
	//     the function's output schema).
	//   - Optional **filters** (string equality, numeric comparators, null-checks) and
	//     **aggregations** (`count`, `count_distinct`, `sum`, `average`, `min`, `max`).
	//
	// Views are versioned: every update produces a new version, and the previous
	// version remains immutable and addressable. Function types that produce
	// transformations with an output schema — `extract`, `transform`, `analyze`,
	// `join` — are all queryable through views; `extract` works uniformly across
	// vision and OCR inputs.
	//
	// ## Reading data
	//
	//   - **`POST /v3/views/table-data`** — paginated rows of column values. Each row
	//     reports the underlying event's `eventID` (the externally-stable KSUID used
	//     everywhere else in V3) plus the projected column values.
	//   - **`POST /v3/views/aggregation-data`** — group-by-able aggregate values across
	//     the same query surface.
	//
	// Both endpoints take a `timeWindow` to bound the transformation set and require
	// at least one `function` to read from.
	Views ViewService
	// Buckets are named partitions of the knowledge graph within an
	// account+environment. Entities, mentions, and relations are scoped to a bucket so
	// a single account+environment can host multiple isolated graphs — for example one
	// per data source or workspace.
	//
	// Every account+environment has exactly one **default** bucket, used by unscoped
	// flows. The default bucket can be renamed but never deleted.
	//
	// Use these endpoints to create, list, fetch, rename, and delete buckets:
	//
	//   - **`POST /v3/buckets`** creates a non-default bucket.
	//   - **`GET /v3/buckets`** lists buckets with cursor pagination (`startingAfter` /
	//     `endingBefore` over `bucketID`).
	//   - **`PATCH /v3/buckets/{bucketID}`** updates `name` and/or `description`.
	//   - **`DELETE /v3/buckets/{bucketID}`** soft-deletes a bucket. A non-empty bucket
	//     is rejected with `409 Conflict` unless `?cascade=true` is passed; the default
	//     bucket can never be deleted.
	Buckets  BucketService
	Entities EntityService
	// Entity Types are the customer-defined taxonomy for the knowledge graph, scoped
	// to an account+environment. Each type has a unique, immutable name and can be
	// organised into hierarchies via `parentTypeID`. A type may carry per-type
	// structured attribute metadata in `attributeSchema` (for example
	// `{"unit": "mg", "range": [0, 100]}`).
	//
	// Use these endpoints to create, list, fetch, update, and delete entity types:
	//
	//   - **`POST /v3/entity-types`** creates a type, optionally under a parent.
	//   - **`GET /v3/entity-types`** lists types with cursor pagination (`startingAfter`
	//     / `endingBefore` over `typeID`) and an optional `parentTypeId` filter for
	//     direct children.
	//   - **`PATCH /v3/entity-types/{typeID}`** updates `description`, `parentTypeID`,
	//     and/or `attributeSchema`. The `name` is immutable.
	//   - **`DELETE /v3/entity-types/{typeID}`** soft-deletes a type. The request is
	//     rejected with `409 Conflict` while any live entity is assigned to the type or
	//     any live child type points at it.
	EntityTypes EntityTypeService
	// Read the cross-document knowledge graph — the canonical entities and the
	// directed relations between them that the Parse pipeline populates when
	// `linkAcrossDocuments` is enabled.
	//
	//   - **`GET /v3/entities/{id}/relations`** returns the inbound and outbound edges
	//     incident to one entity, split by direction. Supports `direction`, an exact
	//     `relationType` filter, and cursor pagination over edges. A merged-away entity
	//     id transparently resolves to its surviving canonical entity.
	//   - **`GET /v3/knowledge-graph`** returns the graph as `{ nodes, edges }`,
	//     paginating over edges. The `nodes` for a page are the distinct endpoint
	//     entities of that page's edges (both endpoints of every edge are included).
	//     Filter with `type[]`, `since`, and `search`; an edge is returned only when
	//     both of its endpoints survive the entity filters.
	//
	// Both endpoints take an optional `bucket` (`bkt_...`) to scope the read to a
	// single bucket; omit it for the unscoped account+environment view.
	KnowledgeGraph KnowledgeGraphService
}

// DefaultClientOptions read from the environment (BEM_API_KEY, BEM_BASE_URL). This
// should be used to initialize new clients.
func DefaultClientOptions() []option.RequestOption {
	defaults := []option.RequestOption{option.WithHTTPClient(defaultHTTPClient()), option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("BEM_BASE_URL"); ok {
		defaults = append(defaults, option.WithBaseURL(o))
	}
	if o, ok := os.LookupEnv("BEM_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	if o, ok := os.LookupEnv("BEM_CUSTOM_HEADERS"); ok {
		for _, line := range strings.Split(o, "\n") {
			colon := strings.Index(line, ":")
			if colon >= 0 {
				defaults = append(defaults, option.WithHeader(strings.TrimSpace(line[:colon]), strings.TrimSpace(line[colon+1:])))
			}
		}
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
	r.Webhooks = NewWebhookService(opts...)
	r.WebhookSecret = NewWebhookSecretService(opts...)
	r.Eval = NewEvalService(opts...)
	r.Fs = NewFService(opts...)
	r.Connectors = NewConnectorService(opts...)
	r.Subscriptions = NewSubscriptionService(opts...)
	r.Views = NewViewService(opts...)
	r.Buckets = NewBucketService(opts...)
	r.Entities = NewEntityService(opts...)
	r.EntityTypes = NewEntityTypeService(opts...)
	r.KnowledgeGraph = NewKnowledgeGraphService(opts...)

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
