// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bem

import (
	"context"
	"net/http"
	"slices"

	"github.com/bem-team/bem-go-sdk/internal/apijson"
	"github.com/bem-team/bem-go-sdk/internal/requestconfig"
	"github.com/bem-team/bem-go-sdk/option"
	"github.com/bem-team/bem-go-sdk/packages/respjson"
)

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
//
// WebhookSecretService contains methods and other services that help with
// interacting with the bem API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookSecretService] method instead.
type WebhookSecretService struct {
	options []option.RequestOption
}

// NewWebhookSecretService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebhookSecretService(opts ...option.RequestOption) (r WebhookSecretService) {
	r = WebhookSecretService{}
	r.options = opts
	return
}

// **Generate a new webhook signing secret.**
//
// Creates a new signing secret for this environment (or replaces the existing
// one). The new secret is returned in full exactly once — store it securely.
//
// After rotation all newly delivered webhooks will be signed with the new secret.
// Update your verification logic before calling this endpoint if you need
// zero-downtime rotation.
func (r *WebhookSecretService) New(ctx context.Context, opts ...option.RequestOption) (res *WebhookSecret, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/webhook-secret"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// **Get the current webhook signing secret.**
//
// Returns the active secret used to sign outbound webhook deliveries via the
// `bem-signature` header. Returns 404 if no secret has been generated for this
// environment yet.
//
// Use the secret to verify incoming webhook payloads:
//
// 1. Parse `bem-signature: t={timestamp},v1={signature}`.
// 2. Construct the signed string: `{timestamp}.{raw request body}`.
// 3. Compute HMAC-SHA256 of that string using the secret.
// 4. Compare the hex digest against `v1`.
// 5. Reject requests where the timestamp is more than a few minutes old.
func (r *WebhookSecretService) Get(ctx context.Context, opts ...option.RequestOption) (res *WebhookSecret, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/webhook-secret"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// **Revoke the current webhook signing secret.**
//
// Deletes the active signing secret. Webhook deliveries will continue but will no
// longer include a `bem-signature` header until a new secret is generated.
func (r *WebhookSecretService) Revoke(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v3/webhook-secret"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Webhook signing secret used to verify `bem-signature` headers on delivered
// webhooks.
type WebhookSecret struct {
	// The signing secret value. Store this securely — it is shown in full only on
	// generation.
	Secret string `json:"secret" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookSecret) RawJSON() string { return r.JSON.raw }
func (r *WebhookSecret) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
