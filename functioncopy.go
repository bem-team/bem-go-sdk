// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bem

import (
	"context"
	"net/http"
	"slices"

	"github.com/bem-team/bem-go-sdk/internal/apijson"
	shimjson "github.com/bem-team/bem-go-sdk/internal/encoding/json"
	"github.com/bem-team/bem-go-sdk/internal/requestconfig"
	"github.com/bem-team/bem-go-sdk/option"
	"github.com/bem-team/bem-go-sdk/packages/param"
)

// Functions are the core building blocks of data transformation in Bem. Each
// function type serves a specific purpose:
//
//   - **Extract**: Extract structured JSON data from unstructured documents (PDFs,
//     emails, images, spreadsheets), with optional layout-aware bounding-box
//     extraction
//   - **Route**: Direct data to different processing paths based on conditions
//   - **Split**: Break multi-page documents into individual pages for parallel
//     processing
//   - **Join**: Combine outputs from multiple function calls into a single result
//   - **Parse**: Render documents into a navigable structure of page-aware sections,
//     named entities, and relationships — designed to be walked by an LLM agent via
//     the [File System API](/api/v3/file-system) (`POST /v3/fs`). Two toggles, both
//     `true` by default: `extractEntities` controls per-document entity and
//     relationship extraction; `linkAcrossDocuments` merges entities into one
//     canonical record per real-world thing across the environment, populating
//     cross-document memory.
//   - **Payload Shaping**: Transform and restructure data using JMESPath expressions
//   - **Enrich**: Enhance data with semantic search against collections
//   - **Send**: Deliver workflow outputs to downstream destinations
//
// Use these endpoints to create, update, list, and manage your functions.
//
// FunctionCopyService contains methods and other services that help with
// interacting with the bem API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFunctionCopyService] method instead.
type FunctionCopyService struct {
	options []option.RequestOption
}

// NewFunctionCopyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFunctionCopyService(opts ...option.RequestOption) (r FunctionCopyService) {
	r = FunctionCopyService{}
	r.options = opts
	return
}

// **Copy a function to a new name within the same environment.**
//
// Forks the source function's current configuration into a brand-new function. The
// copy starts at `versionNum: 1` regardless of how many versions the source has —
// version history is not carried over.
//
// Useful for experimenting with schema or prompt changes against a stable
// production function without disturbing existing callers.
//
// The destination name must be unique in the environment. A copy does not migrate
// workflows: existing workflow nodes continue to reference the original function.
func (r *FunctionCopyService) New(ctx context.Context, body FunctionCopyNewParams, opts ...option.RequestOption) (res *FunctionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/functions/copy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Request to copy an existing function with a new name and optional
// customizations.
//
// The properties SourceFunctionName, TargetFunctionName are required.
type FunctionCopyRequestParam struct {
	// Name of the function to copy from. Must be a valid existing function name.
	SourceFunctionName string `json:"sourceFunctionName" api:"required"`
	// Name for the new copied function. Must be unique within the target environment.
	TargetFunctionName string `json:"targetFunctionName" api:"required"`
	// Optional display name for the copied function. If not provided, defaults to the
	// source function's display name with " (Copy)" appended.
	TargetDisplayName param.Opt[string] `json:"targetDisplayName,omitzero"`
	// Optional environment name to copy the function to. If not provided, the function
	// will be copied within the same environment.
	TargetEnvironment param.Opt[string] `json:"targetEnvironment,omitzero"`
	// Optional array of tags for the copied function. If not provided, defaults to the
	// source function's tags.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r FunctionCopyRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow FunctionCopyRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionCopyRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionCopyNewParams struct {
	// Request to copy an existing function with a new name and optional
	// customizations.
	FunctionCopyRequest FunctionCopyRequestParam
	paramObj
}

func (r FunctionCopyNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.FunctionCopyRequest)
}
func (r *FunctionCopyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
