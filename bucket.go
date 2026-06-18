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
//
// BucketService contains methods and other services that help with interacting
// with the bem API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBucketService] method instead.
type BucketService struct {
	options []option.RequestOption
}

// NewBucketService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBucketService(opts ...option.RequestOption) (r BucketService) {
	r = BucketService{}
	r.options = opts
	return
}

// Create a Bucket
func (r *BucketService) New(ctx context.Context, body BucketNewParams, opts ...option.RequestOption) (res *BucketNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/buckets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a Bucket
func (r *BucketService) Get(ctx context.Context, bucketID string, opts ...option.RequestOption) (res *BucketGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if bucketID == "" {
		err = errors.New("missing required bucketID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/buckets/%s", url.PathEscape(bucketID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a Bucket
func (r *BucketService) Update(ctx context.Context, bucketID string, body BucketUpdateParams, opts ...option.RequestOption) (res *BucketUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if bucketID == "" {
		err = errors.New("missing required bucketID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/buckets/%s", url.PathEscape(bucketID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List Buckets
func (r *BucketService) List(ctx context.Context, query BucketListParams, opts ...option.RequestOption) (res *BucketListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/buckets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a Bucket
func (r *BucketService) Delete(ctx context.Context, bucketID string, body BucketDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if bucketID == "" {
		err = errors.New("missing required bucketID parameter")
		return err
	}
	path := fmt.Sprintf("v3/buckets/%s", url.PathEscape(bucketID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// A Bucket is a named partition of the knowledge graph within an
// account+environment. Entities, mentions, and relations are scoped to a bucket so
// a single account+environment can host multiple isolated graphs.
//
// Every account+environment has exactly one default bucket. The default bucket can
// be renamed but never deleted.
type BucketNewResponse struct {
	// Stable public identifier for the bucket (`bkt_...`).
	BucketID string `json:"bucketID" api:"required"`
	// Creation timestamp (RFC 3339).
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Optional human-facing note about the bucket.
	Description string `json:"description" api:"required"`
	// Whether this is the account+environment's default bucket.
	IsDefault bool `json:"isDefault" api:"required"`
	// Human-facing bucket name. Unique within an account+environment.
	Name string `json:"name" api:"required"`
	// Last-update timestamp (RFC 3339).
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BucketID    respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		IsDefault   respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BucketNewResponse) RawJSON() string { return r.JSON.raw }
func (r *BucketNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Bucket is a named partition of the knowledge graph within an
// account+environment. Entities, mentions, and relations are scoped to a bucket so
// a single account+environment can host multiple isolated graphs.
//
// Every account+environment has exactly one default bucket. The default bucket can
// be renamed but never deleted.
type BucketGetResponse struct {
	// Stable public identifier for the bucket (`bkt_...`).
	BucketID string `json:"bucketID" api:"required"`
	// Creation timestamp (RFC 3339).
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Optional human-facing note about the bucket.
	Description string `json:"description" api:"required"`
	// Whether this is the account+environment's default bucket.
	IsDefault bool `json:"isDefault" api:"required"`
	// Human-facing bucket name. Unique within an account+environment.
	Name string `json:"name" api:"required"`
	// Last-update timestamp (RFC 3339).
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BucketID    respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		IsDefault   respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BucketGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BucketGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Bucket is a named partition of the knowledge graph within an
// account+environment. Entities, mentions, and relations are scoped to a bucket so
// a single account+environment can host multiple isolated graphs.
//
// Every account+environment has exactly one default bucket. The default bucket can
// be renamed but never deleted.
type BucketUpdateResponse struct {
	// Stable public identifier for the bucket (`bkt_...`).
	BucketID string `json:"bucketID" api:"required"`
	// Creation timestamp (RFC 3339).
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Optional human-facing note about the bucket.
	Description string `json:"description" api:"required"`
	// Whether this is the account+environment's default bucket.
	IsDefault bool `json:"isDefault" api:"required"`
	// Human-facing bucket name. Unique within an account+environment.
	Name string `json:"name" api:"required"`
	// Last-update timestamp (RFC 3339).
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BucketID    respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		IsDefault   respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BucketUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *BucketUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response body for listing buckets.
type BucketListResponse struct {
	Buckets []BucketListResponseBucket `json:"buckets" api:"required"`
	// Total number of buckets matching the query, ignoring pagination.
	TotalCount int64 `json:"totalCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Buckets     respjson.Field
		TotalCount  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BucketListResponse) RawJSON() string { return r.JSON.raw }
func (r *BucketListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Bucket is a named partition of the knowledge graph within an
// account+environment. Entities, mentions, and relations are scoped to a bucket so
// a single account+environment can host multiple isolated graphs.
//
// Every account+environment has exactly one default bucket. The default bucket can
// be renamed but never deleted.
type BucketListResponseBucket struct {
	// Stable public identifier for the bucket (`bkt_...`).
	BucketID string `json:"bucketID" api:"required"`
	// Creation timestamp (RFC 3339).
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Optional human-facing note about the bucket.
	Description string `json:"description" api:"required"`
	// Whether this is the account+environment's default bucket.
	IsDefault bool `json:"isDefault" api:"required"`
	// Human-facing bucket name. Unique within an account+environment.
	Name string `json:"name" api:"required"`
	// Last-update timestamp (RFC 3339).
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BucketID    respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		IsDefault   respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BucketListResponseBucket) RawJSON() string { return r.JSON.raw }
func (r *BucketListResponseBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BucketNewParams struct {
	// Bucket name. Required and unique within the account+environment.
	Name string `json:"name" api:"required"`
	// Optional description.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r BucketNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BucketNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BucketNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BucketUpdateParams struct {
	// New description.
	Description param.Opt[string] `json:"description,omitzero"`
	// New name.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r BucketUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BucketUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BucketUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BucketListParams struct {
	// Cursor: return buckets whose `bucketID` sorts before this value.
	EndingBefore param.Opt[string] `query:"endingBefore,omitzero" json:"-"`
	// Maximum number of buckets to return (default 50, max 200).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Case-insensitive substring match on the bucket name.
	NameSubstring param.Opt[string] `query:"nameSubstring,omitzero" json:"-"`
	// Cursor: return buckets whose `bucketID` sorts after this value.
	StartingAfter param.Opt[string] `query:"startingAfter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BucketListParams]'s query parameters as `url.Values`.
func (r BucketListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BucketDeleteParams struct {
	// When `true`, delete the bucket even if it still contains entities (the entities
	// are removed along with it). When omitted or `false`, the request is rejected
	// with `409 Conflict` if the bucket is non-empty.
	//
	// The default bucket can never be deleted regardless of this flag.
	Cascade param.Opt[bool] `query:"cascade,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BucketDeleteParams]'s query parameters as `url.Values`.
func (r BucketDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
