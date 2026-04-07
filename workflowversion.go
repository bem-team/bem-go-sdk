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
	"github.com/bem-team/bem-go-sdk/packages/pagination"
	"github.com/bem-team/bem-go-sdk/packages/param"
	"github.com/bem-team/bem-go-sdk/packages/respjson"
)

// Workflow operations
//
// WorkflowVersionService contains methods and other services that help with
// interacting with the bem API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowVersionService] method instead.
type WorkflowVersionService struct {
	options []option.RequestOption
}

// NewWorkflowVersionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkflowVersionService(opts ...option.RequestOption) (r WorkflowVersionService) {
	r = WorkflowVersionService{}
	r.options = opts
	return
}

// Get a Workflow Version
func (r *WorkflowVersionService) Get(ctx context.Context, versionNum int64, query WorkflowVersionGetParams, opts ...option.RequestOption) (res *WorkflowVersionGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.WorkflowName == "" {
		err = errors.New("missing required workflowName parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/workflows/%s/versions/%v", url.PathEscape(query.WorkflowName), versionNum)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List Workflow Versions
func (r *WorkflowVersionService) List(ctx context.Context, workflowName string, query WorkflowVersionListParams, opts ...option.RequestOption) (res *pagination.WorkflowVersionsPage[WorkflowVersionListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if workflowName == "" {
		err = errors.New("missing required workflowName parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/workflows/%s/versions", url.PathEscape(workflowName))
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Workflow Versions
func (r *WorkflowVersionService) ListAutoPaging(ctx context.Context, workflowName string, query WorkflowVersionListParams, opts ...option.RequestOption) *pagination.WorkflowVersionsPageAutoPager[WorkflowVersionListResponse] {
	return pagination.NewWorkflowVersionsPageAutoPager(r.List(ctx, workflowName, query, opts...))
}

type WorkflowVersionGetResponse struct {
	// Error message if the workflow version retrieval failed.
	Error string `json:"error"`
	// V3 read representation of a workflow version.
	Workflow WorkflowVersionGetResponseWorkflow `json:"workflow"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		Workflow    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowVersionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowVersionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V3 read representation of a workflow version.
type WorkflowVersionGetResponseWorkflow struct {
	// Unique identifier of the workflow.
	ID string `json:"id" api:"required"`
	// The date and time the workflow was created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// All directed edges in this workflow version's DAG.
	Edges []WorkflowVersionGetResponseWorkflowEdge `json:"edges" api:"required"`
	// Name of the entry-point call-site node.
	MainNodeName string `json:"mainNodeName" api:"required"`
	// Unique name of the workflow within the environment.
	Name string `json:"name" api:"required"`
	// All call-site nodes in this workflow version's DAG.
	Nodes []WorkflowVersionGetResponseWorkflowNode `json:"nodes" api:"required"`
	// The date and time the workflow was last updated.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Version number of this workflow version.
	VersionNum int64 `json:"versionNum" api:"required"`
	// Audit trail information.
	Audit WorkflowVersionGetResponseWorkflowAudit `json:"audit"`
	// Human-readable display name.
	DisplayName string `json:"displayName"`
	// Inbound email address associated with the workflow, if any.
	EmailAddress string `json:"emailAddress"`
	// Tags associated with the workflow.
	Tags []string `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Edges        respjson.Field
		MainNodeName respjson.Field
		Name         respjson.Field
		Nodes        respjson.Field
		UpdatedAt    respjson.Field
		VersionNum   respjson.Field
		Audit        respjson.Field
		DisplayName  respjson.Field
		EmailAddress respjson.Field
		Tags         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowVersionGetResponseWorkflow) RawJSON() string { return r.JSON.raw }
func (r *WorkflowVersionGetResponseWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Read representation of a directed edge between call-site nodes.
type WorkflowVersionGetResponseWorkflowEdge struct {
	// Name of the destination node.
	DestinationNodeName string `json:"destinationNodeName" api:"required"`
	// Name of the source node.
	SourceNodeName string `json:"sourceNodeName" api:"required"`
	// Labelled outlet on the source node, if any.
	DestinationName string `json:"destinationName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DestinationNodeName respjson.Field
		SourceNodeName      respjson.Field
		DestinationName     respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowVersionGetResponseWorkflowEdge) RawJSON() string { return r.JSON.raw }
func (r *WorkflowVersionGetResponseWorkflowEdge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Read representation of a call-site node.
type WorkflowVersionGetResponseWorkflowNode struct {
	// Function (and version) executing at this call site.
	Function FunctionVersionIdentifier `json:"function" api:"required"`
	// Name of this call site, unique within the workflow version.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Function    respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowVersionGetResponseWorkflowNode) RawJSON() string { return r.JSON.raw }
func (r *WorkflowVersionGetResponseWorkflowNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Audit trail information.
type WorkflowVersionGetResponseWorkflowAudit struct {
	// Information about who created the current version.
	VersionCreatedBy UserActionSummary `json:"versionCreatedBy"`
	// Information about who created the workflow.
	WorkflowCreatedBy UserActionSummary `json:"workflowCreatedBy"`
	// Information about who last updated the workflow.
	WorkflowLastUpdatedBy UserActionSummary `json:"workflowLastUpdatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		VersionCreatedBy      respjson.Field
		WorkflowCreatedBy     respjson.Field
		WorkflowLastUpdatedBy respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowVersionGetResponseWorkflowAudit) RawJSON() string { return r.JSON.raw }
func (r *WorkflowVersionGetResponseWorkflowAudit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V3 read representation of a workflow version.
type WorkflowVersionListResponse struct {
	// Unique identifier of the workflow.
	ID string `json:"id" api:"required"`
	// The date and time the workflow was created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// All directed edges in this workflow version's DAG.
	Edges []WorkflowVersionListResponseEdge `json:"edges" api:"required"`
	// Name of the entry-point call-site node.
	MainNodeName string `json:"mainNodeName" api:"required"`
	// Unique name of the workflow within the environment.
	Name string `json:"name" api:"required"`
	// All call-site nodes in this workflow version's DAG.
	Nodes []WorkflowVersionListResponseNode `json:"nodes" api:"required"`
	// The date and time the workflow was last updated.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// Version number of this workflow version.
	VersionNum int64 `json:"versionNum" api:"required"`
	// Audit trail information.
	Audit WorkflowVersionListResponseAudit `json:"audit"`
	// Human-readable display name.
	DisplayName string `json:"displayName"`
	// Inbound email address associated with the workflow, if any.
	EmailAddress string `json:"emailAddress"`
	// Tags associated with the workflow.
	Tags []string `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Edges        respjson.Field
		MainNodeName respjson.Field
		Name         respjson.Field
		Nodes        respjson.Field
		UpdatedAt    respjson.Field
		VersionNum   respjson.Field
		Audit        respjson.Field
		DisplayName  respjson.Field
		EmailAddress respjson.Field
		Tags         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowVersionListResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowVersionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Read representation of a directed edge between call-site nodes.
type WorkflowVersionListResponseEdge struct {
	// Name of the destination node.
	DestinationNodeName string `json:"destinationNodeName" api:"required"`
	// Name of the source node.
	SourceNodeName string `json:"sourceNodeName" api:"required"`
	// Labelled outlet on the source node, if any.
	DestinationName string `json:"destinationName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DestinationNodeName respjson.Field
		SourceNodeName      respjson.Field
		DestinationName     respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowVersionListResponseEdge) RawJSON() string { return r.JSON.raw }
func (r *WorkflowVersionListResponseEdge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Read representation of a call-site node.
type WorkflowVersionListResponseNode struct {
	// Function (and version) executing at this call site.
	Function FunctionVersionIdentifier `json:"function" api:"required"`
	// Name of this call site, unique within the workflow version.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Function    respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowVersionListResponseNode) RawJSON() string { return r.JSON.raw }
func (r *WorkflowVersionListResponseNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Audit trail information.
type WorkflowVersionListResponseAudit struct {
	// Information about who created the current version.
	VersionCreatedBy UserActionSummary `json:"versionCreatedBy"`
	// Information about who created the workflow.
	WorkflowCreatedBy UserActionSummary `json:"workflowCreatedBy"`
	// Information about who last updated the workflow.
	WorkflowLastUpdatedBy UserActionSummary `json:"workflowLastUpdatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		VersionCreatedBy      respjson.Field
		WorkflowCreatedBy     respjson.Field
		WorkflowLastUpdatedBy respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowVersionListResponseAudit) RawJSON() string { return r.JSON.raw }
func (r *WorkflowVersionListResponseAudit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowVersionGetParams struct {
	WorkflowName string `path:"workflowName" api:"required" json:"-"`
	paramObj
}

type WorkflowVersionListParams struct {
	EndingBefore  param.Opt[int64] `query:"endingBefore,omitzero" json:"-"`
	Limit         param.Opt[int64] `query:"limit,omitzero" json:"-"`
	StartingAfter param.Opt[int64] `query:"startingAfter,omitzero" json:"-"`
	// Any of "asc", "desc".
	SortOrder WorkflowVersionListParamsSortOrder `query:"sortOrder,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkflowVersionListParams]'s query parameters as
// `url.Values`.
func (r WorkflowVersionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkflowVersionListParamsSortOrder string

const (
	WorkflowVersionListParamsSortOrderAsc  WorkflowVersionListParamsSortOrder = "asc"
	WorkflowVersionListParamsSortOrderDesc WorkflowVersionListParamsSortOrder = "desc"
)
