// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bem

import (
	"context"
	"errors"
	"fmt"
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

// Subscriptions wire up notifications for the events your functions and
// collections produce.
//
// Most subscriptions target a single function (by `functionName` or `functionID`)
// or a single collection (by `collectionName` or `collectionID`) and select a
// `type` corresponding to the event you want to receive — for example `transform`,
// `route`, `join`, `evaluation`, `error`, `enrich`, or `collection_processing`.
//
// Entity-lifecycle events are account-wide and target no function or collection.
// Set `type` to one of the following and provide a `webhookURL` (these event types
// support webhook delivery only):
//
//   - `entity_proposed` — an entity entered the `proposed` curation status (queued
//     for review).
//   - `entity_validated` — an entity was approved/validated by a reviewer.
//   - `entity_rejected` — an entity was rejected by a reviewer.
//
// Each entity-lifecycle delivery is a JSON POST describing the transition
// (`entityID`, `typeName`, `priorStatus`, `newStatus`, optional `actorUserID` and
// `reason`, and a `timestamp`).
//
// Deliveries can be sent to any combination of:
//
// - `webhookURL` — HTTPS endpoint that receives a JSON POST per event.
// - `s3Bucket` + `s3FilePath` — sync output JSON into an AWS S3 prefix you own.
// - `googleDriveFolderID` — drop output JSON into a Google Drive folder.
//
// Use `disabled: true` to pause delivery without deleting the subscription.
// Updates follow conventional PATCH semantics — only the fields you include are
// changed.
//
// SubscriptionService contains methods and other services that help with
// interacting with the bem API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubscriptionService] method instead.
type SubscriptionService struct {
	options []option.RequestOption
}

// NewSubscriptionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSubscriptionService(opts ...option.RequestOption) (r SubscriptionService) {
	r = SubscriptionService{}
	r.options = opts
	return
}

// Creates a new subscription to listen to transform or error events.
func (r *SubscriptionService) New(ctx context.Context, body SubscriptionNewParams, opts ...option.RequestOption) (res *SubscriptionV3, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a Subscription
func (r *SubscriptionService) Get(ctx context.Context, subscriptionID string, opts ...option.RequestOption) (res *SubscriptionV3, err error) {
	opts = slices.Concat(r.options, opts)
	if subscriptionID == "" {
		err = errors.New("missing required subscriptionID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/subscriptions/%s", url.PathEscape(subscriptionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an existing subscription. Follow conventional PATCH behavior, so only
// included fields will be updated.
func (r *SubscriptionService) Update(ctx context.Context, subscriptionID string, body SubscriptionUpdateParams, opts ...option.RequestOption) (res *SubscriptionV3, err error) {
	opts = slices.Concat(r.options, opts)
	if subscriptionID == "" {
		err = errors.New("missing required subscriptionID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/subscriptions/%s", url.PathEscape(subscriptionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List Subscriptions
func (r *SubscriptionService) List(ctx context.Context, query SubscriptionListParams, opts ...option.RequestOption) (res *[]SubscriptionV3, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes an existing subscription.
func (r *SubscriptionService) Delete(ctx context.Context, subscriptionID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscriptionID parameter")
		return err
	}
	path := fmt.Sprintf("v3/subscriptions/%s", url.PathEscape(subscriptionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type SubscriptionV3 struct {
	// Name of subscription.
	Name string `json:"name" api:"required"`
	// The unique identifier of the subscription.
	SubscriptionID string `json:"subscriptionID" api:"required"`
	// Type of subscription.
	//
	// Any of "transform", "analyze", "route", "join", "split_collection",
	// "split_item", "evaluation", "error", "payload_shaping", "enrich",
	// "collection_processing".
	Type SubscriptionV3Type `json:"type" api:"required"`
	// Unique identifier of collection this subscription listens to.
	CollectionID string `json:"collectionID"`
	// Name of collection this subscription listens to.
	CollectionName string `json:"collectionName"`
	// Toggles whether subscription is active or not.
	Disabled bool `json:"disabled"`
	// Unique identifier of function this subscription listens to.
	FunctionID string `json:"functionID"`
	// Unique name of function this subscription listens to.
	FunctionName string `json:"functionName"`
	// Google Drive folder ID for syncing output data to Google Drive.
	GoogleDriveFolderID string `json:"googleDriveFolderID"`
	// S3 bucket name for syncing output data to AWS S3.
	S3Bucket string `json:"s3Bucket"`
	// S3 file path for syncing output data to AWS S3.
	S3FilePath string `json:"s3FilePath"`
	// URL bem will send webhook requests to.
	WebhookURL string `json:"webhookURL"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name                respjson.Field
		SubscriptionID      respjson.Field
		Type                respjson.Field
		CollectionID        respjson.Field
		CollectionName      respjson.Field
		Disabled            respjson.Field
		FunctionID          respjson.Field
		FunctionName        respjson.Field
		GoogleDriveFolderID respjson.Field
		S3Bucket            respjson.Field
		S3FilePath          respjson.Field
		WebhookURL          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubscriptionV3) RawJSON() string { return r.JSON.raw }
func (r *SubscriptionV3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of subscription.
type SubscriptionV3Type string

const (
	SubscriptionV3TypeTransform            SubscriptionV3Type = "transform"
	SubscriptionV3TypeAnalyze              SubscriptionV3Type = "analyze"
	SubscriptionV3TypeRoute                SubscriptionV3Type = "route"
	SubscriptionV3TypeJoin                 SubscriptionV3Type = "join"
	SubscriptionV3TypeSplitCollection      SubscriptionV3Type = "split_collection"
	SubscriptionV3TypeSplitItem            SubscriptionV3Type = "split_item"
	SubscriptionV3TypeEvaluation           SubscriptionV3Type = "evaluation"
	SubscriptionV3TypeError                SubscriptionV3Type = "error"
	SubscriptionV3TypePayloadShaping       SubscriptionV3Type = "payload_shaping"
	SubscriptionV3TypeEnrich               SubscriptionV3Type = "enrich"
	SubscriptionV3TypeCollectionProcessing SubscriptionV3Type = "collection_processing"
)

type SubscriptionNewParams struct {
	// Name of subscription.
	Name string `json:"name" api:"required"`
	// Type of subscription.
	//
	// Any of "transform", "analyze", "route", "join", "split_collection",
	// "split_item", "evaluation", "error", "payload_shaping", "enrich",
	// "collection_processing".
	Type SubscriptionNewParamsType `json:"type,omitzero" api:"required"`
	// Unique identifier of collection this subscription listens to (alternative to
	// collectionName).
	CollectionID param.Opt[string] `json:"collectionID,omitzero"`
	// Name of collection this subscription listens to (required for collection-based
	// subscriptions).
	CollectionName param.Opt[string] `json:"collectionName,omitzero"`
	// Toggles whether subscription is active or not.
	Disabled param.Opt[bool] `json:"disabled,omitzero"`
	// Unique identifier of function this subscription listens to (alternative to
	// functionName).
	FunctionID param.Opt[string] `json:"functionID,omitzero"`
	// Unique name of function this subscription listens to (required for
	// function-based subscriptions).
	FunctionName param.Opt[string] `json:"functionName,omitzero"`
	// Google Drive folder ID for syncing output data to Google Drive.
	GoogleDriveFolderID param.Opt[string] `json:"googleDriveFolderID,omitzero"`
	// S3 bucket name for syncing output data to AWS S3.
	S3Bucket param.Opt[string] `json:"s3Bucket,omitzero"`
	// S3 file path for syncing output data to AWS S3.
	S3FilePath param.Opt[string] `json:"s3FilePath,omitzero"`
	// URL bem will send webhook requests to.
	WebhookURL param.Opt[string] `json:"webhookURL,omitzero"`
	paramObj
}

func (r SubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SubscriptionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubscriptionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of subscription.
type SubscriptionNewParamsType string

const (
	SubscriptionNewParamsTypeTransform            SubscriptionNewParamsType = "transform"
	SubscriptionNewParamsTypeAnalyze              SubscriptionNewParamsType = "analyze"
	SubscriptionNewParamsTypeRoute                SubscriptionNewParamsType = "route"
	SubscriptionNewParamsTypeJoin                 SubscriptionNewParamsType = "join"
	SubscriptionNewParamsTypeSplitCollection      SubscriptionNewParamsType = "split_collection"
	SubscriptionNewParamsTypeSplitItem            SubscriptionNewParamsType = "split_item"
	SubscriptionNewParamsTypeEvaluation           SubscriptionNewParamsType = "evaluation"
	SubscriptionNewParamsTypeError                SubscriptionNewParamsType = "error"
	SubscriptionNewParamsTypePayloadShaping       SubscriptionNewParamsType = "payload_shaping"
	SubscriptionNewParamsTypeEnrich               SubscriptionNewParamsType = "enrich"
	SubscriptionNewParamsTypeCollectionProcessing SubscriptionNewParamsType = "collection_processing"
)

type SubscriptionUpdateParams struct {
	// Toggles whether subscription is active or not.
	Disabled param.Opt[bool] `json:"disabled,omitzero"`
	// Unique name of function this subscription listens to.
	FunctionName param.Opt[string] `json:"functionName,omitzero"`
	// Google Drive folder ID for syncing output data to Google Drive.
	GoogleDriveFolderID param.Opt[string] `json:"googleDriveFolderID,omitzero"`
	// Name of subscription.
	Name param.Opt[string] `json:"name,omitzero"`
	// S3 bucket name for syncing output data to AWS S3.
	S3Bucket param.Opt[string] `json:"s3Bucket,omitzero"`
	// S3 file path for syncing output data to AWS S3.
	S3FilePath param.Opt[string] `json:"s3FilePath,omitzero"`
	// URL bem will send webhook requests to.
	WebhookURL param.Opt[string] `json:"webhookURL,omitzero"`
	// Type of subscription.
	//
	// Any of "transform", "analyze", "route", "join", "split_collection",
	// "split_item", "evaluation", "error", "payload_shaping", "enrich",
	// "collection_processing".
	Type SubscriptionUpdateParamsType `json:"type,omitzero"`
	paramObj
}

func (r SubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SubscriptionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubscriptionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of subscription.
type SubscriptionUpdateParamsType string

const (
	SubscriptionUpdateParamsTypeTransform            SubscriptionUpdateParamsType = "transform"
	SubscriptionUpdateParamsTypeAnalyze              SubscriptionUpdateParamsType = "analyze"
	SubscriptionUpdateParamsTypeRoute                SubscriptionUpdateParamsType = "route"
	SubscriptionUpdateParamsTypeJoin                 SubscriptionUpdateParamsType = "join"
	SubscriptionUpdateParamsTypeSplitCollection      SubscriptionUpdateParamsType = "split_collection"
	SubscriptionUpdateParamsTypeSplitItem            SubscriptionUpdateParamsType = "split_item"
	SubscriptionUpdateParamsTypeEvaluation           SubscriptionUpdateParamsType = "evaluation"
	SubscriptionUpdateParamsTypeError                SubscriptionUpdateParamsType = "error"
	SubscriptionUpdateParamsTypePayloadShaping       SubscriptionUpdateParamsType = "payload_shaping"
	SubscriptionUpdateParamsTypeEnrich               SubscriptionUpdateParamsType = "enrich"
	SubscriptionUpdateParamsTypeCollectionProcessing SubscriptionUpdateParamsType = "collection_processing"
)

type SubscriptionListParams struct {
	// A cursor to use in pagination. `endingBefore` is a task ID that defines your
	// place in the list. For example, if you make a list request and receive 50
	// objects, starting with `sub_2c9AXIj48cUYJtCuv1gsQtHGDzK`, your subsequent call
	// can include `endingBefore=sub_2c9AXIj48cUYJtCuv1gsQtHGDzK` to fetch the previous
	// page of the list.
	EndingBefore param.Opt[string] `query:"endingBefore,omitzero" json:"-"`
	// This specifies a limit on the number of objects to return, ranging between 1
	// and 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// A cursor to use in pagination. `startingAfter` is a task ID that defines your
	// place in the list. For example, if you make a list request and receive 50
	// objects, ending with `sub_2c9AXIj48cUYJtCuv1gsQtHGDzK`, your subsequent call can
	// include `startingAfter=sub_2c9AXIj48cUYJtCuv1gsQtHGDzK` to fetch the next page
	// of the list.
	StartingAfter param.Opt[string] `query:"startingAfter,omitzero" json:"-"`
	// Filters to subscriptions linked to included array of function names.
	FunctionNames []string `query:"functionNames,omitzero" json:"-"`
	// Specifies sorting behavior. The two options are `asc` and `desc` to sort
	// ascending and descending respectively, with default sort being ascending. Paging
	// works in both directions.
	//
	// Any of "asc", "desc".
	SortOrder SubscriptionListParamsSortOrder `query:"sortOrder,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SubscriptionListParams]'s query parameters as `url.Values`.
func (r SubscriptionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies sorting behavior. The two options are `asc` and `desc` to sort
// ascending and descending respectively, with default sort being ascending. Paging
// works in both directions.
type SubscriptionListParamsSortOrder string

const (
	SubscriptionListParamsSortOrderAsc  SubscriptionListParamsSortOrder = "asc"
	SubscriptionListParamsSortOrderDesc SubscriptionListParamsSortOrder = "desc"
)
