// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"
	"reflect"

	"github.com/bem-team/bem-go-sdk/internal/apijson"
	"github.com/bem-team/bem-go-sdk/internal/requestconfig"
	"github.com/bem-team/bem-go-sdk/option"
	"github.com/bem-team/bem-go-sdk/packages/param"
	"github.com/bem-team/bem-go-sdk/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type FunctionsPage[T any] struct {
	Functions []T `json:"functions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Functions   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r FunctionsPage[T]) RawJSON() string { return r.JSON.raw }
func (r *FunctionsPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *FunctionsPage[T]) GetNextPage() (res *FunctionsPage[T], err error) {
	if len(r.Functions) == 0 {
		return nil, nil
	}
	items := r.Functions
	if items == nil || len(items) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	value := reflect.ValueOf(items[len(items)-1])
	field := value.FieldByName("FunctionID")
	err = cfg.Apply(option.WithQuery("startingAfter", field.Interface().(string)))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *FunctionsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &FunctionsPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type FunctionsPageAutoPager[T any] struct {
	page *FunctionsPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewFunctionsPageAutoPager[T any](page *FunctionsPage[T], err error) *FunctionsPageAutoPager[T] {
	return &FunctionsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *FunctionsPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Functions) == 0 {
		return false
	}
	if r.idx >= len(r.page.Functions) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Functions) == 0 {
			return false
		}
	}
	r.cur = r.page.Functions[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *FunctionsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *FunctionsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *FunctionsPageAutoPager[T]) Index() int {
	return r.run
}

type WorkflowsPage[T any] struct {
	Workflows []T `json:"workflows"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Workflows   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r WorkflowsPage[T]) RawJSON() string { return r.JSON.raw }
func (r *WorkflowsPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *WorkflowsPage[T]) GetNextPage() (res *WorkflowsPage[T], err error) {
	if len(r.Workflows) == 0 {
		return nil, nil
	}
	items := r.Workflows
	if items == nil || len(items) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	value := reflect.ValueOf(items[len(items)-1])
	field := value.FieldByName("ID")
	err = cfg.Apply(option.WithQuery("startingAfter", field.Interface().(string)))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *WorkflowsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &WorkflowsPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type WorkflowsPageAutoPager[T any] struct {
	page *WorkflowsPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewWorkflowsPageAutoPager[T any](page *WorkflowsPage[T], err error) *WorkflowsPageAutoPager[T] {
	return &WorkflowsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *WorkflowsPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Workflows) == 0 {
		return false
	}
	if r.idx >= len(r.page.Workflows) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Workflows) == 0 {
			return false
		}
	}
	r.cur = r.page.Workflows[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *WorkflowsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *WorkflowsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *WorkflowsPageAutoPager[T]) Index() int {
	return r.run
}

type CallsPage[T any] struct {
	Calls []T `json:"calls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Calls       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r CallsPage[T]) RawJSON() string { return r.JSON.raw }
func (r *CallsPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CallsPage[T]) GetNextPage() (res *CallsPage[T], err error) {
	if len(r.Calls) == 0 {
		return nil, nil
	}
	items := r.Calls
	if items == nil || len(items) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	value := reflect.ValueOf(items[len(items)-1])
	field := value.FieldByName("CallID")
	err = cfg.Apply(option.WithQuery("startingAfter", field.Interface().(string)))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *CallsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CallsPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CallsPageAutoPager[T any] struct {
	page *CallsPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCallsPageAutoPager[T any](page *CallsPage[T], err error) *CallsPageAutoPager[T] {
	return &CallsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CallsPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Calls) == 0 {
		return false
	}
	if r.idx >= len(r.page.Calls) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Calls) == 0 {
			return false
		}
	}
	r.cur = r.page.Calls[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CallsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *CallsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *CallsPageAutoPager[T]) Index() int {
	return r.run
}

type OutputsPage[T any] struct {
	Outputs []T `json:"outputs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Outputs     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r OutputsPage[T]) RawJSON() string { return r.JSON.raw }
func (r *OutputsPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *OutputsPage[T]) GetNextPage() (res *OutputsPage[T], err error) {
	if len(r.Outputs) == 0 {
		return nil, nil
	}
	items := r.Outputs
	if items == nil || len(items) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	value := reflect.ValueOf(items[len(items)-1])
	field := value.FieldByName("EventID")
	err = cfg.Apply(option.WithQuery("startingAfter", field.Interface().(string)))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *OutputsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &OutputsPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type OutputsPageAutoPager[T any] struct {
	page *OutputsPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewOutputsPageAutoPager[T any](page *OutputsPage[T], err error) *OutputsPageAutoPager[T] {
	return &OutputsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *OutputsPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Outputs) == 0 {
		return false
	}
	if r.idx >= len(r.page.Outputs) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Outputs) == 0 {
			return false
		}
	}
	r.cur = r.page.Outputs[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *OutputsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *OutputsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *OutputsPageAutoPager[T]) Index() int {
	return r.run
}

type ErrorsPage[T any] struct {
	Errors []T `json:"errors"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Errors      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r ErrorsPage[T]) RawJSON() string { return r.JSON.raw }
func (r *ErrorsPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *ErrorsPage[T]) GetNextPage() (res *ErrorsPage[T], err error) {
	if len(r.Errors) == 0 {
		return nil, nil
	}
	items := r.Errors
	if items == nil || len(items) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	value := reflect.ValueOf(items[len(items)-1])
	field := value.FieldByName("EventID")
	err = cfg.Apply(option.WithQuery("startingAfter", field.Interface().(string)))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *ErrorsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &ErrorsPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type ErrorsPageAutoPager[T any] struct {
	page *ErrorsPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewErrorsPageAutoPager[T any](page *ErrorsPage[T], err error) *ErrorsPageAutoPager[T] {
	return &ErrorsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *ErrorsPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Errors) == 0 {
		return false
	}
	if r.idx >= len(r.page.Errors) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Errors) == 0 {
			return false
		}
	}
	r.cur = r.page.Errors[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *ErrorsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *ErrorsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *ErrorsPageAutoPager[T]) Index() int {
	return r.run
}

type WorkflowVersionsPage[T any] struct {
	Versions []T `json:"versions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Versions    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r WorkflowVersionsPage[T]) RawJSON() string { return r.JSON.raw }
func (r *WorkflowVersionsPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *WorkflowVersionsPage[T]) GetNextPage() (res *WorkflowVersionsPage[T], err error) {
	if len(r.Versions) == 0 {
		return nil, nil
	}
	items := r.Versions
	if items == nil || len(items) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	value := reflect.ValueOf(items[len(items)-1])
	field := value.FieldByName("VersionNum")
	err = cfg.Apply(option.WithQuery("startingAfter", field.Interface().(string)))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *WorkflowVersionsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &WorkflowVersionsPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type WorkflowVersionsPageAutoPager[T any] struct {
	page *WorkflowVersionsPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewWorkflowVersionsPageAutoPager[T any](page *WorkflowVersionsPage[T], err error) *WorkflowVersionsPageAutoPager[T] {
	return &WorkflowVersionsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *WorkflowVersionsPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Versions) == 0 {
		return false
	}
	if r.idx >= len(r.page.Versions) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Versions) == 0 {
			return false
		}
	}
	r.cur = r.page.Versions[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *WorkflowVersionsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *WorkflowVersionsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *WorkflowVersionsPageAutoPager[T]) Index() int {
	return r.run
}
