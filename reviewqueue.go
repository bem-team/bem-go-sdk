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

// The reviewer-facing read surface for entity curation, available on the dashboard
// (JWT) only.
//
//   - **`GET /v3/review-queue`** returns a cursor-paginated set of entities awaiting
//     curation, scoped to your account+environment (and optional `bucket`). Each row
//     is a full entity plus a small preview (up to 2) of its first mentions, so a
//     reviewer can triage without opening every entity.
//
// Filters AND together. `status` (repeatable) defaults to the pre-terminal states
// `extracted` + `proposed` when omitted. `type` (repeatable `ety_…` IDs) matches
// the entity's _effective_ type — its assigned type id, or, for entities with no
// assigned type, its bem-inferred type name. `assignedTo` (`me` or a `usr_…` ID)
// restricts to entities whose effective type the user reviews. `since` (RFC3339)
// filters by creation time. Pagination is cursor-based on `entityID` ascending;
// default limit 50, maximum 200.
//
// ReviewQueueService contains methods and other services that help with
// interacting with the bem API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReviewQueueService] method instead.
type ReviewQueueService struct {
	options []option.RequestOption
}

// NewReviewQueueService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewReviewQueueService(opts ...option.RequestOption) (r ReviewQueueService) {
	r = ReviewQueueService{}
	r.options = opts
	return
}

// **List entities awaiting curation, for a human reviewer's queue.**
//
// Returns a cursor-paginated set of entities scoped to your account+environment
// (and optional `bucket`), each carrying a small preview of its first mentions so
// a reviewer can triage without opening every entity. All filters AND together.
//
//   - **`status`** (repeatable) restricts to the given lifecycle states. Omitting it
//     defaults to the pre-terminal states `extracted` and `proposed`.
//   - **`type`** (repeatable, `ety_...` IDs) matches the entity's _effective_ type:
//     an entity matches when its assigned type is one of these IDs, or it has no
//     assigned type and its bem-inferred type name matches one of them.
//   - **`assignedTo`** (`me` or a `usr_...` ID) restricts to entities whose
//     effective type the given user reviews. `me` resolves to the calling user.
//   - **`since`** (RFC3339) restricts to entities created at or after the time.
//
// Pagination is cursor-based on `entityID` ascending; default limit is 50,
// maximum 200.
func (r *ReviewQueueService) List(ctx context.Context, query ReviewQueueListParams, opts ...option.RequestOption) (res *ReviewQueueListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/review-queue"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// `GET /v3/review-queue` response. Cursor-paginated by `entityID` ascending.
type ReviewQueueListResponse struct {
	// The page of entities awaiting curation.
	Entities []ReviewQueueListResponseEntity `json:"entities" api:"required"`
	// Whether more rows exist beyond this page.
	HasMore bool `json:"hasMore" api:"required"`
	// Opaque cursor to pass as `?cursor=` for the next page. Empty when `hasMore` is
	// false.
	NextCursor string `json:"nextCursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		HasMore     respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReviewQueueListResponse) RawJSON() string { return r.JSON.raw }
func (r *ReviewQueueListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One row of the review queue: an entity plus a small preview of its mentions.
type ReviewQueueListResponseEntity struct {
	// The canonical (longest / most descriptive) surface form.
	Canonical string `json:"canonical" api:"required"`
	// When the entity was created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Public ID (`ent_...`) of the entity.
	EntityID string `json:"entityID" api:"required"`
	// Total mentions across all parsed documents.
	MentionCount int64 `json:"mentionCount" api:"required"`
	// A capped preview (up to 2) of the entity's first mentions, ordered by page then
	// time, so a reviewer can triage without opening each entity.
	PreviewMentions []ReviewQueueListResponseEntityPreviewMention `json:"previewMentions" api:"required"`
	// Curation lifecycle state: `extracted`, `proposed`, `approved`, `rejected`.
	Status string `json:"status" api:"required"`
	// Distinct surface forms that have resolved to this entity.
	SurfaceForms []string `json:"surfaceForms" api:"required"`
	// The effective type name (assigned override if set, else bem-inferred).
	Type string `json:"type" api:"required"`
	// When the entity was last updated.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Free-form description of the entity, when present.
	Description string `json:"description"`
	// Public ID (`ety_...`) of the customer-assigned type, when one is set.
	TypeID string `json:"typeID"`
	// When a human approved/rejected the entity. Omitted while un-validated.
	ValidatedAt time.Time `json:"validatedAt" format:"date-time"`
	// Public ID (`usr_...`) of the user who validated the entity, when known.
	ValidatedByUserID string `json:"validatedByUserID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Canonical         respjson.Field
		CreatedAt         respjson.Field
		EntityID          respjson.Field
		MentionCount      respjson.Field
		PreviewMentions   respjson.Field
		Status            respjson.Field
		SurfaceForms      respjson.Field
		Type              respjson.Field
		UpdatedAt         respjson.Field
		Description       respjson.Field
		TypeID            respjson.Field
		ValidatedAt       respjson.Field
		ValidatedByUserID respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReviewQueueListResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *ReviewQueueListResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single per-document occurrence of an entity, used in review-queue previews.
type ReviewQueueListResponseEntityPreviewMention struct {
	// When this mention was recorded.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Public ID (`ent_...`) of the entity this mention resolves to.
	EntityID string `json:"entityID" api:"required"`
	// Public ID (`emn_...`) of this mention.
	MentionID string `json:"mentionID" api:"required"`
	// 1-indexed page number within the source document.
	Page int64 `json:"page" api:"required"`
	// The user-provided document handle this mention came from.
	ReferenceID string `json:"referenceID" api:"required"`
	// The exact surface string Parse extracted on the page.
	Surface string `json:"surface" api:"required"`
	// The parse-emitted section label this mention sat under, when present.
	SectionLabel string `json:"sectionLabel"`
	// Public ID of the parse transformation that produced this mention, when known.
	TransformationID string `json:"transformationID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt        respjson.Field
		EntityID         respjson.Field
		MentionID        respjson.Field
		Page             respjson.Field
		ReferenceID      respjson.Field
		Surface          respjson.Field
		SectionLabel     respjson.Field
		TransformationID respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReviewQueueListResponseEntityPreviewMention) RawJSON() string { return r.JSON.raw }
func (r *ReviewQueueListResponseEntityPreviewMention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReviewQueueListParams struct {
	// `me` or a `usr_...` ID — restrict to entities whose effective type that user
	// reviews.
	AssignedTo param.Opt[string] `query:"assignedTo,omitzero" json:"-"`
	// Optional bucket public ID (`bkt_...`) to scope to. Omit for all buckets.
	Bucket param.Opt[string] `query:"bucket,omitzero" json:"-"`
	// Cursor — an `entityID` defining your place in the list.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	// RFC3339 timestamp — restrict to entities created at or after this time.
	Since param.Opt[string] `query:"since,omitzero" json:"-"`
	// Restrict to these lifecycle states. Defaults to `extracted` + `proposed`.
	Status []string `query:"status,omitzero" json:"-"`
	// Restrict to entities whose effective type is one of these `ety_...` IDs.
	Type []string `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ReviewQueueListParams]'s query parameters as `url.Values`.
func (r ReviewQueueListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
