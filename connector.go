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

// Connectors are integrations that trigger a Bem workflow from an external system.
//
// A connector binds an inbound source — currently Box or a Paragon-managed
// integration such as Google Drive — to a specific workflow (by `workflowName` or
// `workflowID`). When the source observes a new file, Bem invokes the bound
// workflow against that file.
//
// Use these endpoints to create, list, and remove connectors. The fields used at
// create time depend on the connector `type`: Box connectors require Box
// credentials and a folder to watch, while Paragon connectors carry a
// `paragonIntegration` identifier and an integration-specific
// `paragonConfiguration` object (for example, `{ "folderId": "..." }` for Google
// Drive).
//
// ConnectorService contains methods and other services that help with interacting
// with the bem API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectorService] method instead.
type ConnectorService struct {
	options []option.RequestOption
}

// NewConnectorService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConnectorService(opts ...option.RequestOption) (r ConnectorService) {
	r = ConnectorService{}
	r.options = opts
	return
}

// Create a Connector
func (r *ConnectorService) New(ctx context.Context, body ConnectorNewParams, opts ...option.RequestOption) (res *Connector, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/connectors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List Connectors
func (r *ConnectorService) List(ctx context.Context, query ConnectorListParams, opts ...option.RequestOption) (res *ConnectorListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/connectors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a Connector
func (r *ConnectorService) Delete(ctx context.Context, connectorID string, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	if connectorID == "" {
		err = errors.New("missing required connectorID parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/connectors/%s", url.PathEscape(connectorID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// A Connector represents an integration that triggers a Bem workflow from an
// external system.
type Connector struct {
	// Box client ID (from your Box application).
	BoxClientID string `json:"boxClientID" api:"required"`
	// Box client secret (from your Box application).
	//
	// Note: This value is sensitive and should be stored securely.
	BoxClientSecret string `json:"boxClientSecret" api:"required"`
	// Box enterprise ID.
	BoxEnterpriseID string `json:"boxEnterpriseID" api:"required"`
	// Box folder ID to watch for new uploads.
	BoxFolderID string `json:"boxFolderID" api:"required"`
	// Unique identifier for the connector.
	ConnectorID string `json:"connectorID" api:"required"`
	// Human-friendly name for this connector.
	Name string `json:"name" api:"required"`
	// Configuration specific to the type of integration.
	ParagonConfiguration any `json:"paragonConfiguration" api:"required"`
	// Paragon integration, eg. "googledrive".
	ParagonIntegration string `json:"paragonIntegration" api:"required"`
	// Paragon sync ID.
	ParagonSyncID string `json:"paragonSyncID" api:"required"`
	// The connector type.
	//
	// Any of "box", "paragon".
	Type ConnectorType `json:"type" api:"required"`
	// Workflow API ID that will be triggered by this connector.
	WorkflowID string `json:"workflowID" api:"required"`
	// Workflow name that will be triggered by this connector.
	WorkflowName string `json:"workflowName" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BoxClientID          respjson.Field
		BoxClientSecret      respjson.Field
		BoxEnterpriseID      respjson.Field
		BoxFolderID          respjson.Field
		ConnectorID          respjson.Field
		Name                 respjson.Field
		ParagonConfiguration respjson.Field
		ParagonIntegration   respjson.Field
		ParagonSyncID        respjson.Field
		Type                 respjson.Field
		WorkflowID           respjson.Field
		WorkflowName         respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Connector) RawJSON() string { return r.JSON.raw }
func (r *Connector) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Connector type.
type ConnectorType string

const (
	ConnectorTypeBox     ConnectorType = "box"
	ConnectorTypeParagon ConnectorType = "paragon"
)

// Response body for listing connectors.
type ConnectorListResponse struct {
	Connectors []Connector `json:"connectors" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Connectors  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectorListResponse) RawJSON() string { return r.JSON.raw }
func (r *ConnectorListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectorNewParams struct {
	// Human-friendly name for this connector.
	Name string `json:"name" api:"required"`
	// The connector type.
	//
	// Any of "box", "paragon".
	Type ConnectorType `json:"type,omitzero" api:"required"`
	// Box client ID (from your Box application).
	BoxClientID param.Opt[string] `json:"boxClientID,omitzero"`
	// Box client secret (from your Box application).
	BoxClientSecret param.Opt[string] `json:"boxClientSecret,omitzero"`
	// Box enterprise ID.
	BoxEnterpriseID param.Opt[string] `json:"boxEnterpriseID,omitzero"`
	// Box folder ID to watch for new uploads.
	BoxFolderID param.Opt[string] `json:"boxFolderID,omitzero"`
	// Paragon integration, eg. "googledrive".
	ParagonIntegration param.Opt[string] `json:"paragonIntegration,omitzero"`
	// One of `workflowID` or `workflowName` must be provided.
	//
	// If both are provided, they must refer to the same workflow.
	WorkflowID param.Opt[string] `json:"workflowID,omitzero"`
	// One of `workflowID` or `workflowName` must be provided.
	//
	// If both are provided, they must refer to the same workflow.
	WorkflowName param.Opt[string] `json:"workflowName,omitzero"`
	// Configuration specific to the type of integration.
	ParagonConfiguration any `json:"paragonConfiguration,omitzero"`
	paramObj
}

func (r ConnectorNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ConnectorNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConnectorNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectorListParams struct {
	// Filter connectors by workflow API ID (e.g. `wf_...`).
	//
	// If both `workflowID` and `workflowName` are provided, results must match both.
	WorkflowID param.Opt[string] `query:"workflowID,omitzero" json:"-"`
	// Filter connectors by workflow name (exact match).
	//
	// If both `workflowID` and `workflowName` are provided, results must match both.
	WorkflowName param.Opt[string] `query:"workflowName,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConnectorListParams]'s query parameters as `url.Values`.
func (r ConnectorListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
