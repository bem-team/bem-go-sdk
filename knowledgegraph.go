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

// Read the cross-document knowledge graph — the canonical entities and the
// directed relations between them that the Parse pipeline populates when
// `linkAcrossDocuments` is enabled.
//
//   - **`GET /v3/entities/{id}/relations`** returns the inbound and outbound edges
//     incident to one entity, split by direction. Supports `direction`, an exact
//     `relationType` filter, and cursor pagination over edges. A merged-away entity
//     id transparently resolves to its surviving canonical entity.
//   - **`GET /v3/knowledge-graph`** returns the graph as `{ nodes, edges }`,
//     paginating over edges. The `nodes` for a page are the distinct endpoint
//     entities of that page's edges (both endpoints of every edge are included).
//     Filter with `type[]`, `since`, and `search`; an edge is returned only when
//     both of its endpoints survive the entity filters.
//
// Both endpoints take an optional `bucket` (`bkt_...`) to scope the read to a
// single bucket; omit it for the unscoped account+environment view.
//
// KnowledgeGraphService contains methods and other services that help with
// interacting with the bem API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewKnowledgeGraphService] method instead.
type KnowledgeGraphService struct {
	options []option.RequestOption
}

// NewKnowledgeGraphService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewKnowledgeGraphService(opts ...option.RequestOption) (r KnowledgeGraphService) {
	r = KnowledgeGraphService{}
	r.options = opts
	return
}

// Retrieve the Knowledge Graph
func (r *KnowledgeGraphService) Get(ctx context.Context, query KnowledgeGraphGetParams, opts ...option.RequestOption) (res *KnowledgeGraphGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/knowledge-graph"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Response body for `GET /v3/knowledge-graph`. Pagination is over edges; `nodes`
// are the distinct endpoint entities of the returned edge page (both endpoints of
// every edge are included).
type KnowledgeGraphGetResponse struct {
	// The page of edges.
	Edges []KnowledgeGraphGetResponseEdge `json:"edges" api:"required"`
	// Distinct endpoint entities of the returned edge page.
	Nodes []KnowledgeGraphGetResponseNode `json:"nodes" api:"required"`
	// Opaque cursor for the next page of edges, or absent on the last page. Pass it
	// back as `cursor`.
	NextCursor string `json:"nextCursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Edges       respjson.Field
		Nodes       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KnowledgeGraphGetResponse) RawJSON() string { return r.JSON.raw }
func (r *KnowledgeGraphGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One directed edge between two entities, addressed by their public ids.
type KnowledgeGraphGetResponseEdge struct {
	// How many times this edge has been observed.
	MentionCount int64 `json:"mentionCount" api:"required"`
	// Free-form relation label.
	RelationType string `json:"relationType" api:"required"`
	// Source entity public id (`ent_...`).
	SourceID string `json:"sourceId" api:"required"`
	// Target entity public id (`ent_...`).
	TargetID string `json:"targetId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MentionCount respjson.Field
		RelationType respjson.Field
		SourceID     respjson.Field
		TargetID     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KnowledgeGraphGetResponseEdge) RawJSON() string { return r.JSON.raw }
func (r *KnowledgeGraphGetResponseEdge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One entity node in the knowledge graph.
type KnowledgeGraphGetResponseNode struct {
	// Stable public identifier for the entity (`ent_...`).
	ID string `json:"id" api:"required"`
	// Canonical (most descriptive) surface form.
	Canonical string `json:"canonical" api:"required"`
	// Hops from the center node when the request centers the graph on one entity
	// (`nodeID`). The center is depth 0. When the request is uncentered (no `nodeID`),
	// this is 0 for every node.
	Depth int64 `json:"depth" api:"required"`
	// Total mentions of this entity across all parsed documents.
	MentionCount int64 `json:"mentionCount" api:"required"`
	// Effective entity type.
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Canonical    respjson.Field
		Depth        respjson.Field
		MentionCount respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KnowledgeGraphGetResponseNode) RawJSON() string { return r.JSON.raw }
func (r *KnowledgeGraphGetResponseNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KnowledgeGraphGetParams struct {
	// Optional bucket public ID (`bkt_...`) to scope the read to one bucket. Omit for
	// the unscoped (all account+environment) view.
	Bucket param.Opt[string] `query:"bucket,omitzero" json:"-"`
	// Cursor: return edges whose KSUID sorts after this value.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of edges per page (default 50, max 200).
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Maximum hops from the center node. Only meaningful with `nodeID`. Defaults to 2
	// and is clamped down to a system maximum (5).
	MaxDepth param.Opt[int64] `query:"maxDepth,omitzero" json:"-"`
	// Center the graph on this entity (`ent_...`) and only return the subgraph within
	// `maxDepth` hops of it; every node then carries its `depth` (hops from the
	// center, center = 0). Omit for the uncentered whole-graph view. `rootNodeID` and
	// `focusNodeID` are accepted as aliases.
	NodeID param.Opt[string] `query:"nodeID,omitzero" json:"-"`
	// Case-insensitive substring match on canonical names. Both endpoints of an edge
	// must match for the edge (and its nodes) to be returned.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Only edges created at/after this RFC 3339 timestamp.
	Since param.Opt[time.Time] `query:"since,omitzero" format:"date-time" json:"-"`
	// Restrict to entities of these types. An edge is returned only when BOTH of its
	// endpoints survive the type filter.
	Type []string `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [KnowledgeGraphGetParams]'s query parameters as
// `url.Values`.
func (r KnowledgeGraphGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
