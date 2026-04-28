// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bem

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/bem-team/bem-go-sdk/internal/apijson"
	"github.com/bem-team/bem-go-sdk/internal/requestconfig"
	"github.com/bem-team/bem-go-sdk/option"
	"github.com/bem-team/bem-go-sdk/packages/param"
	"github.com/bem-team/bem-go-sdk/packages/respjson"
)

// Unix-shell-style nav over parsed documents and the cross-doc memory store.
//
// `POST /v3/fs` is a single op-driven endpoint designed for LLM agents and
// programmatic consumers that want to walk a corpus the way they'd walk a
// filesystem.
//
// ## Doc-level ops (every parsed document)
//
// - `ls` — list parsed documents with rich per-doc metadata.
// - `cat` — read one doc's parse JSON, sliced (`range`) or projected (`select`).
// - `head` — first N sections of one doc.
// - `grep` — substring or regex search; `scope`, `path`, `countOnly` available.
// - `stat` — metadata only (page/section/entity counts, timestamps).
//
// ## Memory-level ops (require `linkAcrossDocuments: true` on the parse function)
//
// - `find` — list canonical entities across the corpus.
// - `open` — entity + mentions.
// - `xref` — for one entity, sections across docs that mention it (with content).
//
// Memory ops return an empty list with a `hint` when no docs in this environment
// have been memory-linked.
//
// ## Pagination
//
// List ops paginate by cursor — pass the previous response's `nextCursor` back as
// `cursor`; `hasMore: false` signals the last page. Same idiom as `/v3/calls` and
// `/v3/outputs`.
//
// FService contains methods and other services that help with interacting with the
// bem API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFService] method instead.
type FService struct {
	options []option.RequestOption
}

// NewFService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFService(opts ...option.RequestOption) (r FService) {
	r = FService{}
	r.options = opts
	return
}

// **Navigate parsed documents and the cross-doc memory store via Unix-shell
// verbs.**
//
// `POST /v3/fs` is a single op-driven endpoint that lets an LLM agent (or any
// programmatic client) walk a corpus the way it would walk a filesystem — `ls` to
// list, `cat` to read, `grep` to search, `head` for a quick peek, `stat` for
// metadata, and `find` / `open` / `xref` for the cross-doc entity memory layer.
//
// The body always carries an `op` field; other fields apply per op. The response
// envelope is uniform: `{op, data, hasMore?, nextCursor?, count?, hint?}`.
//
// ## Doc-level ops (work on every parsed document)
//
//   - `ls`: list parsed documents with `pageCount`, `sectionCount`, `entityCount`,
//     and a short `previewEntities` array.
//   - `cat`: read one doc's full parse JSON, optionally sliced by `range` (page /
//     pageRange / sectionTypes) or projected by `select` (dotted paths like
//     `["sections.label", "sections.page"]`).
//   - `head`: first N sections of one doc.
//   - `grep`: substring or regex search across recent parse outputs. `scope`
//     restricts to `sections` / `entities` / `relationships`. `path` scopes to one
//     doc. `countOnly: true` returns only the hit count.
//   - `stat`: metadata only — page/section/entity counts, timestamps.
//
// ## Memory-level ops (require `linkAcrossDocuments: true` on the parse function)
//
//   - `find`: list canonical entities, filterable by `type`, `search`, `since`.
//     Returns an empty list with a `hint` when no docs have been memory-linked.
//   - `open`: fetch one entity plus its mentions across docs.
//   - `xref`: for one entity, return the actual sections (with content) across docs
//     that mention it. The "where exactly is X discussed" loop in one round trip.
//
// ## Pagination
//
// List ops (`ls`, `find`) paginate by cursor: pass the last item's `nextCursor`
// from a previous response to fetch the next page; `hasMore: false` signals the
// last page. Same idiom as `/v3/calls` and `/v3/outputs`.
func (r *FService) Navigate(ctx context.Context, body FNavigateParams, opts ...option.RequestOption) (res *FNavigateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/fs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Uniform response shape returned for every `op`. `data` is op-specific JSON (a
// list, an object, or a string), but the wrapper is constant so a client only
// learns one parse path.
type FNavigateResponse struct {
	// Op-specific payload. See per-op shapes below.
	Data any `json:"data" api:"required"`
	// Operations exposed by `POST /v3/fs`.
	//
	// The verbs and their flag names mirror Unix tools so an LLM agent's existing
	// vocabulary maps directly:
	//
	// - `ls` — list parsed documents
	// - `cat` — read one parsed doc (optionally sliced by range / projected by select)
	// - `grep` — substring or regex search across parse outputs
	// - `head` — first N sections of one doc
	// - `stat` — metadata only (page count, section count, parsed at, ...)
	// - `find` — list canonical entities (cross-doc memory)
	// - `open` — entity + mentions
	// - `xref` — entity → sections across docs that mention it
	//
	// Doc-level ops (ls, cat, grep, head, stat) work on every parsed document,
	// regardless of how the parse function was configured.
	//
	// Memory-level ops (find, open, xref) operate on the global entities table which
	// is only populated when the parse function had `linkAcrossDocuments: true`. On
	// environments with no memory-linked docs they return empty data with a hint
	// pointing at the toggle.
	//
	// Any of "ls", "find", "open", "cat", "grep", "xref", "stat", "head".
	Op FNavigateResponseOp `json:"op" api:"required"`
	// Set for ops that return a count rather than a list (`grep` with
	// `countOnly=true`) or as a sanity check on lists.
	Count int64 `json:"count"`
	// True when more pages exist for cursor-paginated ops.
	HasMore bool `json:"hasMore"`
	// Optional human-readable note. Surfaced on memory-level ops (`find` / `open` /
	// `xref`) when the corpus has no memory-linked docs, pointing users at the
	// `linkAcrossDocuments` toggle on the parse function.
	Hint string `json:"hint"`
	// Cursor to pass as `cursor` in the next request to fetch the next page. Empty
	// when `hasMore=false`.
	NextCursor string `json:"nextCursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Op          respjson.Field
		Count       respjson.Field
		HasMore     respjson.Field
		Hint        respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FNavigateResponse) RawJSON() string { return r.JSON.raw }
func (r *FNavigateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Operations exposed by `POST /v3/fs`.
//
// The verbs and their flag names mirror Unix tools so an LLM agent's existing
// vocabulary maps directly:
//
// - `ls` — list parsed documents
// - `cat` — read one parsed doc (optionally sliced by range / projected by select)
// - `grep` — substring or regex search across parse outputs
// - `head` — first N sections of one doc
// - `stat` — metadata only (page count, section count, parsed at, ...)
// - `find` — list canonical entities (cross-doc memory)
// - `open` — entity + mentions
// - `xref` — entity → sections across docs that mention it
//
// Doc-level ops (ls, cat, grep, head, stat) work on every parsed document,
// regardless of how the parse function was configured.
//
// Memory-level ops (find, open, xref) operate on the global entities table which
// is only populated when the parse function had `linkAcrossDocuments: true`. On
// environments with no memory-linked docs they return empty data with a hint
// pointing at the toggle.
type FNavigateResponseOp string

const (
	FNavigateResponseOpLs   FNavigateResponseOp = "ls"
	FNavigateResponseOpFind FNavigateResponseOp = "find"
	FNavigateResponseOpOpen FNavigateResponseOp = "open"
	FNavigateResponseOpCat  FNavigateResponseOp = "cat"
	FNavigateResponseOpGrep FNavigateResponseOp = "grep"
	FNavigateResponseOpXref FNavigateResponseOp = "xref"
	FNavigateResponseOpStat FNavigateResponseOp = "stat"
	FNavigateResponseOpHead FNavigateResponseOp = "head"
)

type FNavigateParams struct {
	// Operations exposed by `POST /v3/fs`.
	//
	// The verbs and their flag names mirror Unix tools so an LLM agent's existing
	// vocabulary maps directly:
	//
	// - `ls` — list parsed documents
	// - `cat` — read one parsed doc (optionally sliced by range / projected by select)
	// - `grep` — substring or regex search across parse outputs
	// - `head` — first N sections of one doc
	// - `stat` — metadata only (page count, section count, parsed at, ...)
	// - `find` — list canonical entities (cross-doc memory)
	// - `open` — entity + mentions
	// - `xref` — entity → sections across docs that mention it
	//
	// Doc-level ops (ls, cat, grep, head, stat) work on every parsed document,
	// regardless of how the parse function was configured.
	//
	// Memory-level ops (find, open, xref) operate on the global entities table which
	// is only populated when the parse function had `linkAcrossDocuments: true`. On
	// environments with no memory-linked docs they return empty data with a hint
	// pointing at the toggle.
	//
	// Any of "ls", "find", "open", "cat", "grep", "xref", "stat", "head".
	Op FNavigateParamsOp `json:"op,omitzero" api:"required"`
	// When true, return only the hit count without snippet payload. Cheaper than
	// fetching matches when the agent only wants a yes/no.
	CountOnly param.Opt[bool] `json:"countOnly,omitzero"`
	// Pagination cursor. Pass the last item's ID from a previous response
	// (`nextCursor`) to fetch the next page.
	Cursor param.Opt[string] `json:"cursor,omitzero"`
	// When true (default), substring/regex matching is case-insensitive.
	IgnoreCase param.Opt[bool] `json:"ignoreCase,omitzero"`
	// Maximum results to return. Defaults vary per op (25–50).
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// First-N count for `op=head`. Defaults to 10.
	N param.Opt[int64] `json:"n,omitzero"`
	// Identifier for ops that operate on a single resource:
	//
	// - cat / head / stat: a parsed document, by `referenceID` or `transformationID`.
	// - open / xref / stat: an entity, by `entityID`.
	Path param.Opt[string] `json:"path,omitzero"`
	// Substring or regex pattern for `op=grep`.
	Pattern param.Opt[string] `json:"pattern,omitzero"`
	// When true, `pattern` is interpreted as a Go regex. Default false.
	Regex param.Opt[bool] `json:"regex,omitzero"`
	// Restricts grep to one part of the parse output. One of `"sections"`,
	// `"entities"`, `"relationships"`, `"all"` (default).
	Scope param.Opt[string] `json:"scope,omitzero"`
	// Filter options for `op=ls` and `op=find`.
	Filter FNavigateParamsFilter `json:"filter,omitzero"`
	// Slice the parse output along page or section dimensions. Used with `op=cat`.
	Range FNavigateParamsRange `json:"range,omitzero"`
	// Project the parse output to specific dotted paths (e.g.
	// `["sections.label", "sections.page"]`), letting an agent map a doc's structure
	// cheaply before reading content. Used with `op=cat`.
	Select []string `json:"select,omitzero"`
	paramObj
}

func (r FNavigateParams) MarshalJSON() (data []byte, err error) {
	type shadow FNavigateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FNavigateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Operations exposed by `POST /v3/fs`.
//
// The verbs and their flag names mirror Unix tools so an LLM agent's existing
// vocabulary maps directly:
//
// - `ls` — list parsed documents
// - `cat` — read one parsed doc (optionally sliced by range / projected by select)
// - `grep` — substring or regex search across parse outputs
// - `head` — first N sections of one doc
// - `stat` — metadata only (page count, section count, parsed at, ...)
// - `find` — list canonical entities (cross-doc memory)
// - `open` — entity + mentions
// - `xref` — entity → sections across docs that mention it
//
// Doc-level ops (ls, cat, grep, head, stat) work on every parsed document,
// regardless of how the parse function was configured.
//
// Memory-level ops (find, open, xref) operate on the global entities table which
// is only populated when the parse function had `linkAcrossDocuments: true`. On
// environments with no memory-linked docs they return empty data with a hint
// pointing at the toggle.
type FNavigateParamsOp string

const (
	FNavigateParamsOpLs   FNavigateParamsOp = "ls"
	FNavigateParamsOpFind FNavigateParamsOp = "find"
	FNavigateParamsOpOpen FNavigateParamsOp = "open"
	FNavigateParamsOpCat  FNavigateParamsOp = "cat"
	FNavigateParamsOpGrep FNavigateParamsOp = "grep"
	FNavigateParamsOpXref FNavigateParamsOp = "xref"
	FNavigateParamsOpStat FNavigateParamsOp = "stat"
	FNavigateParamsOpHead FNavigateParamsOp = "head"
)

// Filter options for `op=ls` and `op=find`.
type FNavigateParamsFilter struct {
	// Match a parsed doc's source function name exactly.
	FunctionName param.Opt[string] `json:"functionName,omitzero"`
	// Substring match on canonical name (entities) or `referenceID` (parsed docs).
	// Case-insensitive.
	Search param.Opt[string] `json:"search,omitzero"`
	// Restrict to resources created at or after this timestamp.
	Since param.Opt[time.Time] `json:"since,omitzero" format:"date-time"`
	// Match an entity's `type` field exactly (e.g. `"drug"`, `"study"`).
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r FNavigateParamsFilter) MarshalJSON() (data []byte, err error) {
	type shadow FNavigateParamsFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FNavigateParamsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Slice the parse output along page or section dimensions. Used with `op=cat`.
type FNavigateParamsRange struct {
	// Restrict sections to one page (1-indexed).
	Page param.Opt[int64] `json:"page,omitzero"`
	// Restrict sections to an inclusive page range. Two-element array of `[from, to]`
	// (both 1-indexed).
	PageRange []int64 `json:"pageRange,omitzero"`
	// Keep only sections whose `type` matches one of these (e.g. `["table", "list"]`).
	SectionTypes []string `json:"sectionTypes,omitzero"`
	paramObj
}

func (r FNavigateParamsRange) MarshalJSON() (data []byte, err error) {
	type shadow FNavigateParamsRange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FNavigateParamsRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
