// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bem

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/bem-go/internal/apijson"
	"github.com/stainless-sdks/bem-go/internal/apiquery"
	"github.com/stainless-sdks/bem-go/internal/requestconfig"
	"github.com/stainless-sdks/bem-go/option"
	"github.com/stainless-sdks/bem-go/packages/pagination"
	"github.com/stainless-sdks/bem-go/packages/param"
	"github.com/stainless-sdks/bem-go/packages/respjson"
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
func (r *WorkflowVersionService) List(ctx context.Context, workflowName string, query WorkflowVersionListParams, opts ...option.RequestOption) (res *pagination.WorkflowVersionsPage[Workflow], err error) {
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
func (r *WorkflowVersionService) ListAutoPaging(ctx context.Context, workflowName string, query WorkflowVersionListParams, opts ...option.RequestOption) *pagination.WorkflowVersionsPageAutoPager[Workflow] {
	return pagination.NewWorkflowVersionsPageAutoPager(r.List(ctx, workflowName, query, opts...))
}

type WorkflowVersionGetResponse struct {
	// Error message if the workflow version retrieval failed.
	Error    string   `json:"error"`
	Workflow Workflow `json:"workflow"`
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
