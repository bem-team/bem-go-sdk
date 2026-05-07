// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bem

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/bem-team/bem-go-sdk/internal/apijson"
	"github.com/bem-team/bem-go-sdk/internal/apiquery"
	"github.com/bem-team/bem-go-sdk/internal/requestconfig"
	"github.com/bem-team/bem-go-sdk/option"
	"github.com/bem-team/bem-go-sdk/packages/param"
	"github.com/bem-team/bem-go-sdk/packages/respjson"
)

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
//
// CollectionItemService contains methods and other services that help with
// interacting with the bem API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCollectionItemService] method instead.
type CollectionItemService struct {
	options []option.RequestOption
}

// NewCollectionItemService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCollectionItemService(opts ...option.RequestOption) (r CollectionItemService) {
	r = CollectionItemService{}
	r.options = opts
	return
}

// Get a Collection
func (r *CollectionItemService) Get(ctx context.Context, query CollectionItemGetParams, opts ...option.RequestOption) (res *Collection, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/collections/items"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Update existing items in a Collection
func (r *CollectionItemService) Update(ctx context.Context, body CollectionItemUpdateParams, opts ...option.RequestOption) (res *CollectionItemUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/collections/items"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Delete an item from a Collection
func (r *CollectionItemService) Delete(ctx context.Context, body CollectionItemDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v3/collections/items"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// Add new items to a Collection
func (r *CollectionItemService) Add(ctx context.Context, body CollectionItemAddParams, opts ...option.RequestOption) (res *CollectionItemAddResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/collections/items"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Response after queuing items for async update
type CollectionItemUpdateResponse struct {
	// Event ID for tracking this operation. Use this to correlate with webhook
	// notifications.
	EventID string `json:"eventID" api:"required"`
	// Status message
	Message string `json:"message" api:"required"`
	// Processing status
	//
	// Any of "pending".
	Status CollectionItemUpdateResponseStatus `json:"status" api:"required"`
	// Array of items that were updated (only present in synchronous mode, deprecated)
	Items []CollectionItem `json:"items"`
	// Number of items updated (only present in synchronous mode, deprecated)
	UpdatedCount int64 `json:"updatedCount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventID      respjson.Field
		Message      respjson.Field
		Status       respjson.Field
		Items        respjson.Field
		UpdatedCount respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionItemUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *CollectionItemUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Processing status
type CollectionItemUpdateResponseStatus string

const (
	CollectionItemUpdateResponseStatusPending CollectionItemUpdateResponseStatus = "pending"
)

// Response after queuing items for async processing
type CollectionItemAddResponse struct {
	// Event ID for tracking this operation. Use this to correlate with webhook
	// notifications.
	EventID string `json:"eventID" api:"required"`
	// Status message
	Message string `json:"message" api:"required"`
	// Processing status
	//
	// Any of "pending".
	Status CollectionItemAddResponseStatus `json:"status" api:"required"`
	// Number of new items added (only present in synchronous mode, deprecated)
	AddedCount int64 `json:"addedCount"`
	// Array of items that were added (only present in synchronous mode, deprecated)
	Items []CollectionItem `json:"items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventID     respjson.Field
		Message     respjson.Field
		Status      respjson.Field
		AddedCount  respjson.Field
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionItemAddResponse) RawJSON() string { return r.JSON.raw }
func (r *CollectionItemAddResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Processing status
type CollectionItemAddResponseStatus string

const (
	CollectionItemAddResponseStatusPending CollectionItemAddResponseStatus = "pending"
)

type CollectionItemGetParams struct {
	// The name/path of the collection. Must use only letters, digits, underscores, and
	// dots. Each segment must start with a letter or underscore.
	CollectionName string `query:"collectionName" api:"required" json:"-"`
	// When true, includes items from all subcollections under the specified collection
	// path. For example, querying "customers" with this flag will return items from
	// "customers", "customers.premium", "customers.premium.vip", etc.
	IncludeSubcollections param.Opt[bool] `query:"includeSubcollections,omitzero" json:"-"`
	// Number of items per page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number for pagination
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CollectionItemGetParams]'s query parameters as
// `url.Values`.
func (r CollectionItemGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CollectionItemUpdateParams struct {
	// The name/path of the collection. Must use only letters, digits, underscores, and
	// dots. Each segment must start with a letter or underscore.
	CollectionName string `json:"collectionName" api:"required"`
	// Array of items to update (maximum 100 items per request)
	Items []CollectionItemUpdateParamsItem `json:"items,omitzero" api:"required"`
	paramObj
}

func (r CollectionItemUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CollectionItemUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionItemUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Data for updating an existing item in a collection
//
// The properties CollectionItemID, Data are required.
type CollectionItemUpdateParamsItem struct {
	// Unique identifier for the item to update
	CollectionItemID string `json:"collectionItemID" api:"required"`
	// The updated data to be embedded and stored (string or JSON object)
	Data any `json:"data,omitzero" api:"required"`
	paramObj
}

func (r CollectionItemUpdateParamsItem) MarshalJSON() (data []byte, err error) {
	type shadow CollectionItemUpdateParamsItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionItemUpdateParamsItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionItemDeleteParams struct {
	// The unique identifier of the item to delete
	CollectionItemID string `query:"collectionItemID" api:"required" json:"-"`
	// The name/path of the collection. Must use only letters, digits, underscores, and
	// dots. Each segment must start with a letter or underscore.
	CollectionName string `query:"collectionName" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [CollectionItemDeleteParams]'s query parameters as
// `url.Values`.
func (r CollectionItemDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CollectionItemAddParams struct {
	// The name/path of the collection. Must use only letters, digits, underscores, and
	// dots. Each segment must start with a letter or underscore.
	CollectionName string `json:"collectionName" api:"required"`
	// Array of items to add (maximum 100 items per request)
	Items []CollectionItemAddParamsItem `json:"items,omitzero" api:"required"`
	paramObj
}

func (r CollectionItemAddParams) MarshalJSON() (data []byte, err error) {
	type shadow CollectionItemAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionItemAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Data for creating a new item in a collection
//
// The property Data is required.
type CollectionItemAddParamsItem struct {
	// The data to be embedded and stored (string or JSON object)
	Data any `json:"data,omitzero" api:"required"`
	paramObj
}

func (r CollectionItemAddParamsItem) MarshalJSON() (data []byte, err error) {
	type shadow CollectionItemAddParamsItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionItemAddParamsItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
