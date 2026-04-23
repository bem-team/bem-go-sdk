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
func (r *WebhookSecretService) New(ctx context.Context, opts ...option.RequestOption) (res *WebhookSecretNewResponse, err error) {
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
func (r *WebhookSecretService) Get(ctx context.Context, opts ...option.RequestOption) (res *WebhookSecretGetResponse, err error) {
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
type WebhookSecretNewResponse struct {
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
func (r WebhookSecretNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookSecretNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Webhook signing secret used to verify `bem-signature` headers on delivered
// webhooks.
type WebhookSecretGetResponse struct {
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
func (r WebhookSecretGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookSecretGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
