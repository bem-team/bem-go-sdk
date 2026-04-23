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
// CollectionService contains methods and other services that help with interacting
// with the bem API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCollectionService] method instead.
type CollectionService struct {
	options []option.RequestOption
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
	Items CollectionItemService
}

// NewCollectionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCollectionService(opts ...option.RequestOption) (r CollectionService) {
	r = CollectionService{}
	r.options = opts
	r.Items = NewCollectionItemService(opts...)
	return
}

// Create a Collection
func (r *CollectionService) New(ctx context.Context, body CollectionNewParams, opts ...option.RequestOption) (res *CollectionNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/collections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List Collections
func (r *CollectionService) List(ctx context.Context, query CollectionListParams, opts ...option.RequestOption) (res *CollectionListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/collections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a Collection
func (r *CollectionService) Delete(ctx context.Context, body CollectionDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v3/collections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// Count the number of tokens in the provided texts using the BGE M3 tokenizer.
// This is useful for checking if texts will fit within the embedding model's token
// limit (8,192 tokens per text) before sending them for embedding.
func (r *CollectionService) CountTokens(ctx context.Context, body CollectionCountTokensParams, opts ...option.RequestOption) (res *CollectionCountTokensResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/collections/token-count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Collection details
type CollectionNewResponse struct {
	// Unique identifier for the collection
	CollectionID string `json:"collectionID" api:"required"`
	// The collection name/path. Only letters, digits, underscores, and dots are
	// allowed.
	CollectionName string `json:"collectionName" api:"required"`
	// When the collection was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Number of items in the collection
	ItemCount int64 `json:"itemCount" api:"required"`
	// List of items in the collection (when fetching collection details)
	Items []CollectionNewResponseItem `json:"items"`
	// Number of items per page
	Limit int64 `json:"limit"`
	// Current page number
	Page int64 `json:"page"`
	// Total number of pages
	TotalPages int64 `json:"totalPages"`
	// When the collection was last updated
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CollectionID   respjson.Field
		CollectionName respjson.Field
		CreatedAt      respjson.Field
		ItemCount      respjson.Field
		Items          respjson.Field
		Limit          respjson.Field
		Page           respjson.Field
		TotalPages     respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *CollectionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single item in a collection
type CollectionNewResponseItem struct {
	// Unique identifier for the item
	CollectionItemID string `json:"collectionItemID" api:"required"`
	// When the item was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The data stored in this item
	Data any `json:"data" api:"required"`
	// When the item was last updated
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CollectionItemID respjson.Field
		CreatedAt        respjson.Field
		Data             respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionNewResponseItem) RawJSON() string { return r.JSON.raw }
func (r *CollectionNewResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for listing collections
type CollectionListResponse struct {
	// List of collections
	Collections []CollectionListResponseCollection `json:"collections" api:"required"`
	// Number of collections per page
	Limit int64 `json:"limit" api:"required"`
	// Current page number
	Page int64 `json:"page" api:"required"`
	// Total number of collections
	TotalCount int64 `json:"totalCount" api:"required"`
	// Total number of pages
	TotalPages int64 `json:"totalPages" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Collections respjson.Field
		Limit       respjson.Field
		Page        respjson.Field
		TotalCount  respjson.Field
		TotalPages  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionListResponse) RawJSON() string { return r.JSON.raw }
func (r *CollectionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Collection metadata without items
type CollectionListResponseCollection struct {
	// Unique identifier for the collection
	CollectionID string `json:"collectionID" api:"required"`
	// The collection name/path. Only letters, digits, underscores, and dots are
	// allowed.
	CollectionName string `json:"collectionName" api:"required"`
	// When the collection was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Number of items in the collection
	ItemCount int64 `json:"itemCount" api:"required"`
	// When the collection was last updated
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CollectionID   respjson.Field
		CollectionName respjson.Field
		CreatedAt      respjson.Field
		ItemCount      respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionListResponseCollection) RawJSON() string { return r.JSON.raw }
func (r *CollectionListResponseCollection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for the token count endpoint.
type CollectionCountTokensResponse struct {
	// Maximum tokens allowed per text by the embedding model.
	MaxTokenLimit int64 `json:"max_token_limit"`
	// Number of input texts that exceed `max_token_limit`.
	TextsExceedingLimit int64 `json:"texts_exceeding_limit"`
	// Per-text tokenization results in the same order as the request.
	TokenCounts []CollectionCountTokensResponseTokenCount `json:"token_counts"`
	// Sum of `token_count` across all texts.
	TotalTokens int64 `json:"total_tokens"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxTokenLimit       respjson.Field
		TextsExceedingLimit respjson.Field
		TokenCounts         respjson.Field
		TotalTokens         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionCountTokensResponse) RawJSON() string { return r.JSON.raw }
func (r *CollectionCountTokensResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-text token count result.
type CollectionCountTokensResponseTokenCount struct {
	// Character count of the input text.
	CharCount int64 `json:"char_count"`
	// True if `token_count` exceeds the embedding model's per-text limit.
	ExceedsLimit bool `json:"exceeds_limit"`
	// Zero-based position of this entry in the request `texts` array.
	Index int64 `json:"index"`
	// Number of tokens produced by the tokenizer.
	TokenCount int64 `json:"token_count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CharCount    respjson.Field
		ExceedsLimit respjson.Field
		Index        respjson.Field
		TokenCount   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionCountTokensResponseTokenCount) RawJSON() string { return r.JSON.raw }
func (r *CollectionCountTokensResponseTokenCount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionNewParams struct {
	// Unique name/path for the collection. Supports dot notation for hierarchical
	// paths.
	//
	//   - Only letters (a-z, A-Z), digits (0-9), underscores (\_), and dots (.) are
	//     allowed
	//   - Each segment (between dots) must start with a letter or underscore (not a
	//     digit)
	//   - Segments cannot consist only of digits
	//   - Each segment must be 1-256 characters
	//   - No leading, trailing, or consecutive dots
	//   - Invalid names are rejected with a 400 Bad Request error
	//
	// **Valid Examples:**
	//
	// - 'product_catalog'
	// - 'orders.line_items.sku'
	// - 'customer_data'
	// - 'price_v2'
	//
	// **Invalid Examples:**
	//
	// - 'product-catalog' (contains hyphen)
	// - '123items' (starts with digit)
	// - 'items..data' (consecutive dots)
	// - 'order#123' (contains invalid character #)
	CollectionName string `json:"collectionName" api:"required"`
	paramObj
}

func (r CollectionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CollectionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionListParams struct {
	// Optional substring search filter for collection names (case-insensitive). For
	// example, "premium" will match "customers.premium", "products.premium", etc.
	CollectionNameSearch param.Opt[string] `query:"collectionNameSearch,omitzero" json:"-"`
	// Number of collections per page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number for pagination
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Optional filter to list only collections under a specific parent collection
	// path. For example, "customers" will return "customers", "customers.premium",
	// "customers.premium.vip", etc.
	ParentCollectionName param.Opt[string] `query:"parentCollectionName,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CollectionListParams]'s query parameters as `url.Values`.
func (r CollectionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CollectionDeleteParams struct {
	// The name/path of the collection to delete. Must use only letters, digits,
	// underscores, and dots. Each segment must start with a letter or underscore.
	CollectionName string `query:"collectionName" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [CollectionDeleteParams]'s query parameters as `url.Values`.
func (r CollectionDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CollectionCountTokensParams struct {
	// One or more texts to tokenize.
	Texts []string `json:"texts,omitzero" api:"required"`
	paramObj
}

func (r CollectionCountTokensParams) MarshalJSON() (data []byte, err error) {
	type shadow CollectionCountTokensParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CollectionCountTokensParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
