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
// UserService contains methods and other services that help with interacting with
// the bem API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r UserService) {
	r = UserService{}
	r.options = opts
	return
}

// List a User's Reviewer Assignments
func (r *UserService) ListReviewerAssignments(ctx context.Context, userID string, opts ...option.RequestOption) (res *UserListReviewerAssignmentsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required userID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/users/%s/reviewer-assignments", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Response body for the reverse lookup of a user's reviewer assignments.
type UserListReviewerAssignmentsResponse struct {
	Assignments []UserListReviewerAssignmentsResponseAssignment `json:"assignments" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Assignments respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserListReviewerAssignmentsResponse) RawJSON() string { return r.JSON.raw }
func (r *UserListReviewerAssignmentsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One entity type a user reviews, as returned by the reverse-lookup endpoint. The
// type is exposed via its public ID plus its name and description.
type UserListReviewerAssignmentsResponseAssignment struct {
	// When the assignment was created (RFC 3339).
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The entity type's description.
	Description string `json:"description" api:"required"`
	// The entity type's human-facing name.
	Name string `json:"name" api:"required"`
	// Public ID (`ety_...`) of the entity type the user reviews.
	TypeID string `json:"typeID" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Description respjson.Field
		Name        respjson.Field
		TypeID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserListReviewerAssignmentsResponseAssignment) RawJSON() string { return r.JSON.raw }
func (r *UserListReviewerAssignmentsResponseAssignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
