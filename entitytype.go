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
//
// EntityTypeService contains methods and other services that help with interacting
// with the bem API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntityTypeService] method instead.
type EntityTypeService struct {
	options []option.RequestOption
	// Reviewer assignments link users to the entity types they are responsible for
	// reviewing, scoped to an account+environment. These are dashboard-only endpoints:
	// an assignment needs a user identity, which only the dashboard (JWT) surface
	// carries.
	//
	//   - **`POST /v3/entity-types/{typeID}/reviewers`** assigns a user as a reviewer of
	//     the type. The assignment is idempotent: re-assigning an existing reviewer
	//     returns the existing assignment. Requires the `admin` role.
	//   - **`GET /v3/entity-types/{typeID}/reviewers`** lists the users assigned to
	//     review the type, with each user's email and role. Requires the `operator`
	//     role.
	//   - **`DELETE /v3/entity-types/{typeID}/reviewers/{userID}`** removes an
	//     assignment. Requires the `admin` role.
	//   - **`GET /v3/users/{userID}/reviewer-assignments`** is the reverse lookup: the
	//     entity types a user reviews. A user may read their own assignments; reading
	//     another user's assignments requires the `admin` role.
	Reviewers EntityTypeReviewerService
}

// NewEntityTypeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEntityTypeService(opts ...option.RequestOption) (r EntityTypeService) {
	r = EntityTypeService{}
	r.options = opts
	r.Reviewers = NewEntityTypeReviewerService(opts...)
	return
}

// Create an Entity Type
func (r *EntityTypeService) New(ctx context.Context, body EntityTypeNewParams, opts ...option.RequestOption) (res *EntityType, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/entity-types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get an Entity Type
func (r *EntityTypeService) Get(ctx context.Context, typeID string, opts ...option.RequestOption) (res *EntityType, err error) {
	opts = slices.Concat(r.options, opts)
	if typeID == "" {
		err = errors.New("missing required typeID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/entity-types/%s", url.PathEscape(typeID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update an Entity Type
func (r *EntityTypeService) Update(ctx context.Context, typeID string, body EntityTypeUpdateParams, opts ...option.RequestOption) (res *EntityType, err error) {
	opts = slices.Concat(r.options, opts)
	if typeID == "" {
		err = errors.New("missing required typeID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/entity-types/%s", url.PathEscape(typeID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List Entity Types
func (r *EntityTypeService) List(ctx context.Context, query EntityTypeListParams, opts ...option.RequestOption) (res *EntityTypeListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/entity-types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete an Entity Type
func (r *EntityTypeService) Delete(ctx context.Context, typeID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if typeID == "" {
		err = errors.New("missing required typeID parameter")
		return err
	}
	path := fmt.Sprintf("v3/entity-types/%s", url.PathEscape(typeID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// An EntityType is a customer-defined type in the knowledge-graph taxonomy, scoped
// to an account+environment. Types may be organised into hierarchies via
// `parentTypeID`, and may carry per-type structured attribute metadata in
// `attributeSchema` (for example `{"unit": "mg", "range": [0, 100]}`).
type EntityType struct {
	// Creation timestamp (RFC 3339).
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Optional human-facing note about the type.
	Description string `json:"description" api:"required"`
	// Human-facing type name. Unique within an account+environment, and immutable once
	// set.
	Name string `json:"name" api:"required"`
	// Public ID (`ety_...`) of the parent type, or an empty string when the type is
	// top-level.
	ParentTypeID string `json:"parentTypeID" api:"required"`
	// Stable public identifier for the entity type (`ety_...`).
	TypeID string `json:"typeID" api:"required"`
	// Last-update timestamp (RFC 3339).
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Optional per-type structured attribute metadata.
	AttributeSchema any `json:"attributeSchema"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt       respjson.Field
		Description     respjson.Field
		Name            respjson.Field
		ParentTypeID    respjson.Field
		TypeID          respjson.Field
		UpdatedAt       respjson.Field
		AttributeSchema respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityType) RawJSON() string { return r.JSON.raw }
func (r *EntityType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response body for listing entity types.
type EntityTypeListResponse struct {
	EntityTypes []EntityType `json:"entityTypes" api:"required"`
	// Total number of entity types matching the query, ignoring pagination.
	TotalCount int64 `json:"totalCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EntityTypes respjson.Field
		TotalCount  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityTypeListResponse) RawJSON() string { return r.JSON.raw }
func (r *EntityTypeListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EntityTypeNewParams struct {
	// Type name. Required and unique within the account+environment.
	Name string `json:"name" api:"required"`
	// Optional description.
	Description param.Opt[string] `json:"description,omitzero"`
	// Optional public ID (`ety_...`) of the parent type. Must belong to the same
	// account+environment.
	ParentTypeID param.Opt[string] `json:"parentTypeID,omitzero"`
	// Optional per-type structured attribute metadata.
	AttributeSchema any `json:"attributeSchema,omitzero"`
	paramObj
}

func (r EntityTypeNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EntityTypeNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EntityTypeNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EntityTypeUpdateParams struct {
	// New description.
	Description param.Opt[string] `json:"description,omitzero"`
	// New parent type public ID (`ety_...`), or an empty string to clear the parent
	// (promote to top-level). Must belong to the same account+environment and may not
	// be the type itself.
	ParentTypeID param.Opt[string] `json:"parentTypeID,omitzero"`
	// New per-type structured attribute metadata.
	AttributeSchema any `json:"attributeSchema,omitzero"`
	paramObj
}

func (r EntityTypeUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow EntityTypeUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EntityTypeUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EntityTypeListParams struct {
	// Cursor: return types whose `typeID` sorts before this value.
	EndingBefore param.Opt[string] `query:"endingBefore,omitzero" json:"-"`
	// Maximum number of entity types to return (default 50, max 200).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter to the direct children of this parent type (`ety_...`).
	ParentTypeID param.Opt[string] `query:"parentTypeId,omitzero" json:"-"`
	// Cursor: return types whose `typeID` sorts after this value.
	StartingAfter param.Opt[string] `query:"startingAfter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EntityTypeListParams]'s query parameters as `url.Values`.
func (r EntityTypeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
