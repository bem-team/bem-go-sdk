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

// EntityService contains methods and other services that help with interacting
// with the bem API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntityService] method instead.
type EntityService struct {
	options []option.RequestOption
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
	Synonyms EntitySynonymService
}

// NewEntityService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEntityService(opts ...option.RequestOption) (r EntityService) {
	r = EntityService{}
	r.options = opts
	r.Synonyms = NewEntitySynonymService(opts...)
	return
}

// Update Entity
func (r *EntityService) Update(ctx context.Context, id string, body EntityUpdateParams, opts ...option.RequestOption) (res *EntityUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/entities/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Bulk Seed Entities
func (r *EntityService) BulkNew(ctx context.Context, body EntityBulkNewParams, opts ...option.RequestOption) (res *EntityBulkNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/entities/bulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Bulk Validate Entities
func (r *EntityService) BulkValidate(ctx context.Context, body EntityBulkValidateParams, opts ...option.RequestOption) (res *EntityBulkValidateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/entities/bulk-validate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get an Entity's Relations
func (r *EntityService) GetRelations(ctx context.Context, id string, query EntityGetRelationsParams, opts ...option.RequestOption) (res *EntityGetRelationsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/entities/%s/relations", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get Seed Job Status
func (r *EntityService) GetSeedStatus(ctx context.Context, id string, opts ...option.RequestOption) (res *EntityGetSeedStatusResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/entities/seed/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// An entity record, including its curation status and assigned type.
type EntityUpdateResponse struct {
	// The canonical (longest / most descriptive) surface form.
	Canonical string `json:"canonical" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Public ID (`ent_...`).
	EntityID string `json:"entityID" api:"required"`
	// Total mentions across parsed documents.
	MentionCount int64 `json:"mentionCount" api:"required"`
	// Curation lifecycle state.
	//
	// Any of "extracted", "proposed", "approved", "rejected".
	Status EntityUpdateResponseStatus `json:"status" api:"required"`
	// Distinct surface forms resolved to this entity.
	SurfaceForms []string `json:"surfaceForms" api:"required"`
	// The entity's effective type name (assigned type if set, else inferred).
	Type string `json:"type" api:"required"`
	// Last-update timestamp.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Free-form description.
	Description string `json:"description"`
	// `ety_...` public ID of the assigned type, when one is set.
	TypeID string `json:"typeID"`
	// When the entity was approved/rejected. Present only once validated.
	ValidatedAt time.Time `json:"validatedAt" format:"date-time"`
	// `usr_...` public ID of the validating user (dashboard transitions only).
	ValidatedByUserID string `json:"validatedByUserID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Canonical         respjson.Field
		CreatedAt         respjson.Field
		EntityID          respjson.Field
		MentionCount      respjson.Field
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
func (r EntityUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *EntityUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Curation lifecycle state.
type EntityUpdateResponseStatus string

const (
	EntityUpdateResponseStatusExtracted EntityUpdateResponseStatus = "extracted"
	EntityUpdateResponseStatusProposed  EntityUpdateResponseStatus = "proposed"
	EntityUpdateResponseStatusApproved  EntityUpdateResponseStatus = "approved"
	EntityUpdateResponseStatusRejected  EntityUpdateResponseStatus = "rejected"
)

// `200` response for a synchronously processed (small) batch.
type EntityBulkNewResponse struct {
	// Per-row outcomes, in request order.
	Results []EntityBulkNewResponseResult `json:"results" api:"required"`
	// Per-outcome tally across a batch.
	Summary EntityBulkNewResponseSummary `json:"summary" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityBulkNewResponse) RawJSON() string { return r.JSON.raw }
func (r *EntityBulkNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The outcome of seeding one row.
type EntityBulkNewResponseResult struct {
	// The canonical name from the input row.
	Canonical string `json:"canonical" api:"required"`
	// What happened to this row: `created` (new entity), `merged-with` (matched an
	// existing entity), or `rejected` (see `reason`).
	//
	// Any of "created", "merged-with", "rejected".
	Outcome string `json:"outcome" api:"required"`
	// Public ID (`ent_...`) of the created or merged entity. Absent when rejected.
	EntityID string `json:"entityID"`
	// Human-readable explanation when `outcome` is `rejected`.
	Reason string `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Canonical   respjson.Field
		Outcome     respjson.Field
		EntityID    respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityBulkNewResponseResult) RawJSON() string { return r.JSON.raw }
func (r *EntityBulkNewResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-outcome tally across a batch.
type EntityBulkNewResponseSummary struct {
	// Number of rows that created a new entity.
	Created int64 `json:"created" api:"required"`
	// Number of rows merged into an existing entity.
	Merged int64 `json:"merged" api:"required"`
	// Number of rows rejected.
	Rejected int64 `json:"rejected" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Created     respjson.Field
		Merged      respjson.Field
		Rejected    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityBulkNewResponseSummary) RawJSON() string { return r.JSON.raw }
func (r *EntityBulkNewResponseSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `200` response for `POST /v3/entities/bulk-validate`.
type EntityBulkValidateResponse struct {
	// Per-row outcomes, in request order.
	Results []EntityBulkValidateResponseResult `json:"results" api:"required"`
	// Per-outcome tally across a bulk-validate batch.
	Summary EntityBulkValidateResponseSummary `json:"summary" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityBulkValidateResponse) RawJSON() string { return r.JSON.raw }
func (r *EntityBulkValidateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The outcome of validating one row.
type EntityBulkValidateResponseResult struct {
	// The `ent_...` ID from the request.
	EntityID string `json:"entityID" api:"required"`
	// `validated` (transition applied), `skipped` (not found or not authorized), or
	// `rejected-row` (the transition itself was illegal, e.g. already terminal).
	//
	// Any of "validated", "skipped", "rejected-row".
	Outcome string `json:"outcome" api:"required"`
	// Explanation for a `skipped` or `rejected-row` outcome.
	Reason string `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EntityID    respjson.Field
		Outcome     respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityBulkValidateResponseResult) RawJSON() string { return r.JSON.raw }
func (r *EntityBulkValidateResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-outcome tally across a bulk-validate batch.
type EntityBulkValidateResponseSummary struct {
	// Rows whose transition was illegal.
	RejectedRow int64 `json:"rejectedRow" api:"required"`
	// Rows skipped (not found / not authorized).
	Skipped int64 `json:"skipped" api:"required"`
	// Rows whose transition was applied.
	Validated int64 `json:"validated" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RejectedRow respjson.Field
		Skipped     respjson.Field
		Validated   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityBulkValidateResponseSummary) RawJSON() string { return r.JSON.raw }
func (r *EntityBulkValidateResponseSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response body for `GET /v3/entities/{id}/relations`.
type EntityGetRelationsResponse struct {
	// Edges pointing at the queried entity.
	Inbound []EntityGetRelationsResponseInbound `json:"inbound" api:"required"`
	// Edges pointing away from the queried entity.
	Outbound []EntityGetRelationsResponseOutbound `json:"outbound" api:"required"`
	// Opaque cursor for the next page of edges, or absent on the last page. Pass it
	// back as `cursor`.
	NextCursor string `json:"nextCursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Inbound     respjson.Field
		Outbound    respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityGetRelationsResponse) RawJSON() string { return r.JSON.raw }
func (r *EntityGetRelationsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One edge pointing AT the queried entity (some other entity is the source).
type EntityGetRelationsResponseInbound struct {
	// First-seen timestamp of the edge (RFC 3339).
	FirstSeenAt time.Time `json:"firstSeenAt" api:"required" format:"date-time"`
	// How many times this edge has been observed across parsed documents.
	MentionCount int64 `json:"mentionCount" api:"required"`
	// Free-form relation label (e.g. `author_of`, `affiliated_with`).
	RelationType string `json:"relationType" api:"required"`
	// A compact view of an entity sitting on the far end of a relation edge — the
	// stable public id, the canonical name, and the effective type. The full entity is
	// fetched separately via the entity detail / File System endpoints.
	SourceEntity EntityGetRelationsResponseInboundSourceEntity `json:"sourceEntity" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstSeenAt  respjson.Field
		MentionCount respjson.Field
		RelationType respjson.Field
		SourceEntity respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityGetRelationsResponseInbound) RawJSON() string { return r.JSON.raw }
func (r *EntityGetRelationsResponseInbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A compact view of an entity sitting on the far end of a relation edge — the
// stable public id, the canonical name, and the effective type. The full entity is
// fetched separately via the entity detail / File System endpoints.
type EntityGetRelationsResponseInboundSourceEntity struct {
	// Stable public identifier for the entity (`ent_...`).
	ID string `json:"id" api:"required"`
	// Canonical (most descriptive) surface form of the entity.
	Canonical string `json:"canonical" api:"required"`
	// Effective entity type.
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Canonical   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityGetRelationsResponseInboundSourceEntity) RawJSON() string { return r.JSON.raw }
func (r *EntityGetRelationsResponseInboundSourceEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One edge pointing AWAY from the queried entity (it is the source).
type EntityGetRelationsResponseOutbound struct {
	// First-seen timestamp of the edge (RFC 3339).
	FirstSeenAt time.Time `json:"firstSeenAt" api:"required" format:"date-time"`
	// How many times this edge has been observed across parsed documents.
	MentionCount int64 `json:"mentionCount" api:"required"`
	// Free-form relation label (e.g. `author_of`, `affiliated_with`).
	RelationType string `json:"relationType" api:"required"`
	// A compact view of an entity sitting on the far end of a relation edge — the
	// stable public id, the canonical name, and the effective type. The full entity is
	// fetched separately via the entity detail / File System endpoints.
	TargetEntity EntityGetRelationsResponseOutboundTargetEntity `json:"targetEntity" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstSeenAt  respjson.Field
		MentionCount respjson.Field
		RelationType respjson.Field
		TargetEntity respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityGetRelationsResponseOutbound) RawJSON() string { return r.JSON.raw }
func (r *EntityGetRelationsResponseOutbound) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A compact view of an entity sitting on the far end of a relation edge — the
// stable public id, the canonical name, and the effective type. The full entity is
// fetched separately via the entity detail / File System endpoints.
type EntityGetRelationsResponseOutboundTargetEntity struct {
	// Stable public identifier for the entity (`ent_...`).
	ID string `json:"id" api:"required"`
	// Canonical (most descriptive) surface form of the entity.
	Canonical string `json:"canonical" api:"required"`
	// Effective entity type.
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Canonical   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityGetRelationsResponseOutboundTargetEntity) RawJSON() string { return r.JSON.raw }
func (r *EntityGetRelationsResponseOutboundTargetEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `GET /v3/entities/seed/{id}` response.
type EntityGetSeedStatusResponse struct {
	// Rows that created a new entity.
	CreatedCount int64 `json:"createdCount" api:"required"`
	// Rows merged into an existing entity.
	MergedCount int64 `json:"mergedCount" api:"required"`
	// Rows rejected.
	RejectedCount int64 `json:"rejectedCount" api:"required"`
	// Public ID (`esj_...`) of the seed job.
	SeedJobID string `json:"seedJobID" api:"required"`
	// Lifecycle state.
	//
	// Any of "pending", "processing", "completed", "failed".
	Status EntityGetSeedStatusResponseStatus `json:"status" api:"required"`
	// Total rows in the submitted batch.
	TotalRows int64 `json:"totalRows" api:"required"`
	// Terminal error message when `status` is `failed`.
	Error string `json:"error"`
	// Per-row outcomes. Present only once `status` is `completed`.
	Results []EntityGetSeedStatusResponseResult `json:"results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedCount  respjson.Field
		MergedCount   respjson.Field
		RejectedCount respjson.Field
		SeedJobID     respjson.Field
		Status        respjson.Field
		TotalRows     respjson.Field
		Error         respjson.Field
		Results       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityGetSeedStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *EntityGetSeedStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lifecycle state.
type EntityGetSeedStatusResponseStatus string

const (
	EntityGetSeedStatusResponseStatusPending    EntityGetSeedStatusResponseStatus = "pending"
	EntityGetSeedStatusResponseStatusProcessing EntityGetSeedStatusResponseStatus = "processing"
	EntityGetSeedStatusResponseStatusCompleted  EntityGetSeedStatusResponseStatus = "completed"
	EntityGetSeedStatusResponseStatusFailed     EntityGetSeedStatusResponseStatus = "failed"
)

// The outcome of seeding one row.
type EntityGetSeedStatusResponseResult struct {
	// The canonical name from the input row.
	Canonical string `json:"canonical" api:"required"`
	// What happened to this row: `created` (new entity), `merged-with` (matched an
	// existing entity), or `rejected` (see `reason`).
	//
	// Any of "created", "merged-with", "rejected".
	Outcome string `json:"outcome" api:"required"`
	// Public ID (`ent_...`) of the created or merged entity. Absent when rejected.
	EntityID string `json:"entityID"`
	// Human-readable explanation when `outcome` is `rejected`.
	Reason string `json:"reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Canonical   respjson.Field
		Outcome     respjson.Field
		EntityID    respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityGetSeedStatusResponseResult) RawJSON() string { return r.JSON.raw }
func (r *EntityGetSeedStatusResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EntityUpdateParams struct {
	// The `ety_...` public ID of the type to assign (overriding the bem-inferred
	// type). The empty string clears the assignment. Omit to leave unchanged.
	AssignedTypeID param.Opt[string] `json:"assignedTypeID,omitzero"`
	// Replace the entity's canonical surface form (re-derives its normalized form).
	Canonical param.Opt[string] `json:"canonical,omitzero"`
	// Optional BCP 47 locale tag stamped on any added synonyms.
	Locale param.Opt[string] `json:"locale,omitzero"`
	// Surface forms to attach as `customer_defined` synonyms.
	AddSynonyms []string `json:"addSynonyms,omitzero"`
	// `esn_...` synonym IDs to soft-delete. Only `customer_defined` / `sme_approved`
	// synonyms may be removed; an `extracted` synonym is rejected with `409`.
	RemoveSynonymIDs []string `json:"removeSynonymIDs,omitzero"`
	// Transition the entity's curation status. Only `approved` or `rejected` are
	// accepted, and only from `extracted` or `proposed` (any other transition is
	// rejected with `409`).
	//
	// Any of "approved", "rejected".
	Status EntityUpdateParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r EntityUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow EntityUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EntityUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transition the entity's curation status. Only `approved` or `rejected` are
// accepted, and only from `extracted` or `proposed` (any other transition is
// rejected with `409`).
type EntityUpdateParamsStatus string

const (
	EntityUpdateParamsStatusApproved EntityUpdateParamsStatus = "approved"
	EntityUpdateParamsStatusRejected EntityUpdateParamsStatus = "rejected"
)

type EntityBulkNewParams struct {
	// The entities to seed. Must be non-empty.
	Entities []EntityBulkNewParamsEntity `json:"entities,omitzero" api:"required"`
	// Optional bucket public ID (`bkt_...`) to seed into. Omit to use the
	// account+environment default bucket.
	Bucket param.Opt[string] `json:"bucket,omitzero"`
	// Conflict strategy for an entity that already exists. Only `merge` is supported
	// and it is the default: synonyms are added additively, a longer description
	// replaces the old one, and attributes are merged with new keys winning.
	//
	// Any of "merge".
	OnConflict EntityBulkNewParamsOnConflict `json:"onConflict,omitzero"`
	paramObj
}

func (r EntityBulkNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EntityBulkNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EntityBulkNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One entity to seed in a `POST /v3/entities/bulk` batch.
//
// The properties Canonical, Type are required.
type EntityBulkNewParamsEntity struct {
	// The canonical (longest / most descriptive) surface form for the entity, e.g.
	// `Acme Corporation`. Required. Normalized (lowercased, whitespace-folded) for the
	// uniqueness key.
	Canonical string `json:"canonical" api:"required"`
	// The entity type name, e.g. `instrument` or `organization`. Required. Resolved
	// against your taxonomy and created if it does not yet exist.
	Type string `json:"type" api:"required"`
	// Optional free-form description of the entity.
	Description param.Opt[string] `json:"description,omitzero"`
	// Optional per-entity structured attribute values, e.g.
	// `{ "manufacturer": "Acme", "dosageMg": 50 }`. When the entity's type declares an
	// attribute schema, keys not present in that schema cause the row to be rejected.
	Attributes any `json:"attributes,omitzero"`
	// Optional additional surface forms to attach as `customer_defined` synonyms.
	Synonyms []string `json:"synonyms,omitzero"`
	paramObj
}

func (r EntityBulkNewParamsEntity) MarshalJSON() (data []byte, err error) {
	type shadow EntityBulkNewParamsEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EntityBulkNewParamsEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Conflict strategy for an entity that already exists. Only `merge` is supported
// and it is the default: synonyms are added additively, a longer description
// replaces the old one, and attributes are merged with new keys winning.
type EntityBulkNewParamsOnConflict string

const (
	EntityBulkNewParamsOnConflictMerge EntityBulkNewParamsOnConflict = "merge"
)

type EntityBulkValidateParams struct {
	// The `ent_...` IDs to transition. Must be non-empty.
	EntityIDs []string `json:"entityIDs,omitzero" api:"required"`
	// Terminal status to apply to every entity.
	//
	// Any of "approved", "rejected".
	Status EntityBulkValidateParamsStatus `json:"status,omitzero" api:"required"`
	paramObj
}

func (r EntityBulkValidateParams) MarshalJSON() (data []byte, err error) {
	type shadow EntityBulkValidateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EntityBulkValidateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Terminal status to apply to every entity.
type EntityBulkValidateParamsStatus string

const (
	EntityBulkValidateParamsStatusApproved EntityBulkValidateParamsStatus = "approved"
	EntityBulkValidateParamsStatusRejected EntityBulkValidateParamsStatus = "rejected"
)

type EntityGetRelationsParams struct {
	// Optional bucket public ID (`bkt_...`) to scope the read to one bucket. Omit for
	// the unscoped (all account+environment) view.
	Bucket param.Opt[string] `query:"bucket,omitzero" json:"-"`
	// Cursor: return edges whose KSUID sorts after this value.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of edges to return (default 50, max 200).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Exact-match filter on the relation label.
	RelationType param.Opt[string] `query:"relationType,omitzero" json:"-"`
	// Which edges to return relative to the entity. Defaults to `both`.
	//
	// Any of "inbound", "outbound", "both".
	Direction EntityGetRelationsParamsDirection `query:"direction,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EntityGetRelationsParams]'s query parameters as
// `url.Values`.
func (r EntityGetRelationsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Which edges to return relative to the entity. Defaults to `both`.
type EntityGetRelationsParamsDirection string

const (
	EntityGetRelationsParamsDirectionInbound  EntityGetRelationsParamsDirection = "inbound"
	EntityGetRelationsParamsDirectionOutbound EntityGetRelationsParamsDirection = "outbound"
	EntityGetRelationsParamsDirectionBoth     EntityGetRelationsParamsDirection = "both"
)
