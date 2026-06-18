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
	"github.com/bem-team/bem-go-sdk/internal/requestconfig"
	"github.com/bem-team/bem-go-sdk/option"
	"github.com/bem-team/bem-go-sdk/packages/param"
	"github.com/bem-team/bem-go-sdk/packages/respjson"
)

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
//
// EntityTypeReviewerService contains methods and other services that help with
// interacting with the bem API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntityTypeReviewerService] method instead.
type EntityTypeReviewerService struct {
	options []option.RequestOption
}

// NewEntityTypeReviewerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEntityTypeReviewerService(opts ...option.RequestOption) (r EntityTypeReviewerService) {
	r = EntityTypeReviewerService{}
	r.options = opts
	return
}

// List Reviewers
func (r *EntityTypeReviewerService) List(ctx context.Context, typeID string, opts ...option.RequestOption) (res *EntityTypeReviewerListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if typeID == "" {
		err = errors.New("missing required typeID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/entity-types/%s/reviewers", url.PathEscape(typeID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Assign a Reviewer
func (r *EntityTypeReviewerService) Assign(ctx context.Context, typeID string, body EntityTypeReviewerAssignParams, opts ...option.RequestOption) (res *EntityTypeReviewerAssignResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if typeID == "" {
		err = errors.New("missing required typeID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/entity-types/%s/reviewers", url.PathEscape(typeID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Remove a Reviewer
func (r *EntityTypeReviewerService) Remove(ctx context.Context, userID string, body EntityTypeReviewerRemoveParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.TypeID == "" {
		err = errors.New("missing required typeID parameter")
		return err
	}
	if userID == "" {
		err = errors.New("missing required userID parameter")
		return err
	}
	path := fmt.Sprintf("v3/entity-types/%s/reviewers/%s", url.PathEscape(body.TypeID), url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Response body for listing the reviewers of an entity type.
type EntityTypeReviewerListResponse struct {
	Reviewers []EntityTypeReviewerListResponseReviewer `json:"reviewers" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reviewers   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityTypeReviewerListResponse) RawJSON() string { return r.JSON.raw }
func (r *EntityTypeReviewerListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A reviewer assignment links a user to an entity type they are responsible for
// reviewing. The assignment is scoped to an account+environment and is unique per
// (entity type, user).
type EntityTypeReviewerListResponseReviewer struct {
	// When the assignment was created (RFC 3339).
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The assigned user's email.
	Email string `json:"email" api:"required"`
	// Stable public identifier for the assignment (`etr_...`).
	ReviewerID string `json:"reviewerID" api:"required"`
	// The assigned user's account role (for example `operator`, `admin`).
	Role string `json:"role" api:"required"`
	// Public identifier of the assigned user (`usr_...`).
	UserID string `json:"userID" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Email       respjson.Field
		ReviewerID  respjson.Field
		Role        respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityTypeReviewerListResponseReviewer) RawJSON() string { return r.JSON.raw }
func (r *EntityTypeReviewerListResponseReviewer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A reviewer assignment links a user to an entity type they are responsible for
// reviewing. The assignment is scoped to an account+environment and is unique per
// (entity type, user).
type EntityTypeReviewerAssignResponse struct {
	// When the assignment was created (RFC 3339).
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The assigned user's email.
	Email string `json:"email" api:"required"`
	// Stable public identifier for the assignment (`etr_...`).
	ReviewerID string `json:"reviewerID" api:"required"`
	// The assigned user's account role (for example `operator`, `admin`).
	Role string `json:"role" api:"required"`
	// Public identifier of the assigned user (`usr_...`).
	UserID string `json:"userID" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Email       respjson.Field
		ReviewerID  respjson.Field
		Role        respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityTypeReviewerAssignResponse) RawJSON() string { return r.JSON.raw }
func (r *EntityTypeReviewerAssignResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EntityTypeReviewerAssignParams struct {
	// Public ID (`usr_...`) of the user to assign. Must belong to the account.
	UserID string `json:"userID" api:"required"`
	paramObj
}

func (r EntityTypeReviewerAssignParams) MarshalJSON() (data []byte, err error) {
	type shadow EntityTypeReviewerAssignParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EntityTypeReviewerAssignParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EntityTypeReviewerRemoveParams struct {
	TypeID string `path:"typeID" api:"required" json:"-"`
	paramObj
}
