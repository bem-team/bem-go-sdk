// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bem

import (
	"context"
	"errors"
	"fmt"
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

// Manage the human-readable surface forms (synonyms) attached to a canonical
// entity. Synonyms feed the matcher's exact-match path, so adding the right
// synonyms improves cross-document entity resolution.
//
//   - **`POST /v3/entities/{id}/synonyms`** attaches a `customer_defined` synonym.
//     If the same normalized form already exists as an `extracted` synonym, it is
//     upgraded to `customer_defined` (so the matcher weights it higher); an existing
//     customer/SME synonym is returned unchanged.
//   - **`DELETE /v3/entities/{id}/synonyms/{synonymID}`** soft-deletes a synonym.
//     Only `customer_defined` and `sme_approved` synonyms are deletable; `extracted`
//     synonyms are resolver-owned and the request is rejected with `409 Conflict`.
//
// A merged-away entity id transparently resolves to its surviving canonical
// entity, so a synonym added to a stale id lands on the entity that persists.
//
// EntitySynonymService contains methods and other services that help with
// interacting with the bem API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntitySynonymService] method instead.
type EntitySynonymService struct {
	options []option.RequestOption
}

// NewEntitySynonymService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEntitySynonymService(opts ...option.RequestOption) (r EntitySynonymService) {
	r = EntitySynonymService{}
	r.options = opts
	return
}

// Add a Synonym to an Entity
func (r *EntitySynonymService) Add(ctx context.Context, id string, params EntitySynonymAddParams, opts ...option.RequestOption) (res *EntitySynonymAddResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/entities/%s/synonyms", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Remove a Synonym from an Entity
func (r *EntitySynonymService) Remove(ctx context.Context, synonymID string, params EntitySynonymRemoveParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if synonymID == "" {
		err = errors.New("missing required synonymID parameter")
		return err
	}
	path := fmt.Sprintf("v3/entities/%s/synonyms/%s", url.PathEscape(params.ID), url.PathEscape(synonymID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return err
}

// One synonym attached to an entity.
type EntitySynonymAddResponse struct {
	// Creation timestamp of the synonym (RFC 3339).
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Lowercased, whitespace-folded form of `text`.
	NormalizedText string `json:"normalizedText" api:"required"`
	// Provenance of the synonym. `customer_defined` and `sme_approved` synonyms are
	// deletable; `extracted` synonyms are resolver-owned and cannot be deleted.
	//
	// Any of "extracted", "customer_defined", "sme_approved".
	Source EntitySynonymAddResponseSource `json:"source" api:"required"`
	// Stable public identifier for the synonym (`esn_...`).
	SynonymID string `json:"synonymID" api:"required"`
	// The human-readable synonym as authored.
	Text string `json:"text" api:"required"`
	// Optional BCP 47 locale tag, when one was supplied.
	Locale string `json:"locale"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt      respjson.Field
		NormalizedText respjson.Field
		Source         respjson.Field
		SynonymID      respjson.Field
		Text           respjson.Field
		Locale         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntitySynonymAddResponse) RawJSON() string { return r.JSON.raw }
func (r *EntitySynonymAddResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provenance of the synonym. `customer_defined` and `sme_approved` synonyms are
// deletable; `extracted` synonyms are resolver-owned and cannot be deleted.
type EntitySynonymAddResponseSource string

const (
	EntitySynonymAddResponseSourceExtracted       EntitySynonymAddResponseSource = "extracted"
	EntitySynonymAddResponseSourceCustomerDefined EntitySynonymAddResponseSource = "customer_defined"
	EntitySynonymAddResponseSourceSmeApproved     EntitySynonymAddResponseSource = "sme_approved"
)

type EntitySynonymAddParams struct {
	// The human-readable synonym surface form to attach (e.g. `Acme Corp`, `ACME`). It
	// is normalized (lowercased, whitespace-folded) for the uniqueness key and the
	// matcher's exact-match path.
	Text string `json:"text" api:"required"`
	// Optional bucket public ID (`bkt_...`) to scope the entity lookup to one bucket.
	// Omit for the unscoped (all account+environment) view.
	Bucket param.Opt[string] `query:"bucket,omitzero" json:"-"`
	// Optional BCP 47 locale tag (e.g. `en-US`) for language-specific synonyms.
	Locale param.Opt[string] `json:"locale,omitzero"`
	paramObj
}

func (r EntitySynonymAddParams) MarshalJSON() (data []byte, err error) {
	type shadow EntitySynonymAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EntitySynonymAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [EntitySynonymAddParams]'s query parameters as `url.Values`.
func (r EntitySynonymAddParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EntitySynonymRemoveParams struct {
	ID string `path:"id" api:"required" json:"-"`
	// Optional bucket public ID (`bkt_...`) to scope the entity lookup to one bucket.
	// Omit for the unscoped (all account+environment) view.
	Bucket param.Opt[string] `query:"bucket,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EntitySynonymRemoveParams]'s query parameters as
// `url.Values`.
func (r EntitySynonymRemoveParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
