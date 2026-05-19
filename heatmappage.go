// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/with-ours/platform-sdk-go/internal/apijson"
	"github.com/with-ours/platform-sdk-go/internal/apiquery"
	"github.com/with-ours/platform-sdk-go/internal/requestconfig"
	"github.com/with-ours/platform-sdk-go/option"
	"github.com/with-ours/platform-sdk-go/packages/pagination"
	"github.com/with-ours/platform-sdk-go/packages/param"
	"github.com/with-ours/platform-sdk-go/packages/respjson"
)

// HeatmapPageService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHeatmapPageService] method instead.
type HeatmapPageService struct {
	options []option.RequestOption
}

// NewHeatmapPageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHeatmapPageService(opts ...option.RequestOption) (r HeatmapPageService) {
	r = HeatmapPageService{}
	r.options = opts
	return
}

// List pages with heatmap coverage in a date window, ranked for triage. Each
// entity is identified by `pageKey` (origin + pathname, query string stripped);
// use that value to drill into `GET /rest/v1/heatmap-pages/summary`. Supports
// cursor pagination — the underlying store uses offset internally, so cursors are
// bounded to roughly 10,000 entries deep; if you need pages beyond that, narrow
// `from`/`to` or add filters rather than paginating further. `from`/`to` are UTC
// calendar days in `YYYY-MM-DD`; the window must be 60 days or fewer. Requires
// scope: web-analytics:view
func (r *HeatmapPageService) List(ctx context.Context, query HeatmapPageListParams, opts ...option.RequestOption) (res *pagination.Cursor[HeatmapPageListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "rest/v1/heatmap-pages"
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

// List pages with heatmap coverage in a date window, ranked for triage. Each
// entity is identified by `pageKey` (origin + pathname, query string stripped);
// use that value to drill into `GET /rest/v1/heatmap-pages/summary`. Supports
// cursor pagination — the underlying store uses offset internally, so cursors are
// bounded to roughly 10,000 entries deep; if you need pages beyond that, narrow
// `from`/`to` or add filters rather than paginating further. `from`/`to` are UTC
// calendar days in `YYYY-MM-DD`; the window must be 60 days or fewer. Requires
// scope: web-analytics:view
func (r *HeatmapPageService) ListAutoPaging(ctx context.Context, query HeatmapPageListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[HeatmapPageListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Bundled per-page heatmap rollup: click bins, dead clicks, rage clicks, scroll
// depth, and up to 5 curated in-app replay URLs in a single payload. Designed as a
// one-call diagnosis surface for an AI assistant or marketer — answers "what is
// happening on this page?" without fanning out across five endpoints. `pageKey`
// comes from `GET /rest/v1/heatmap-pages`. The endpoint is identified by query
// params rather than a path id because heatmap pages are not stored entities; this
// is a documented derived-read exception. `breakpoint` scopes the bin/scroll
// aggregations; replays are returned across all breakpoints regardless of
// `breakpoint` (weighted to cover multiple viewports) so callers can compare
// devices. Requires scope: web-analytics:view
func (r *HeatmapPageService) Summary(ctx context.Context, query HeatmapPageSummaryParams, opts ...option.RequestOption) (res *HeatmapPageSummaryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "rest/v1/heatmap-pages/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type HeatmapPageListResponse struct {
	Breakpoints []HeatmapPageListResponseBreakpoint `json:"breakpoints" api:"required"`
	DeadClicks  int64                               `json:"deadClicks" api:"required"`
	DeadRate    float64                             `json:"deadRate" api:"required"`
	IssueScore  int64                               `json:"issueScore" api:"required"`
	// Stable per-page identifier (origin + pathname, query string stripped). Use this
	// as the `pageKey` argument to `GET /rest/v1/heatmap-pages/summary` and the
	// heatmap MCP tools.
	PageKey     string  `json:"pageKey" api:"required"`
	RageClicks  int64   `json:"rageClicks" api:"required"`
	RageRate    float64 `json:"rageRate" api:"required"`
	TotalClicks int64   `json:"totalClicks" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Breakpoints respjson.Field
		DeadClicks  respjson.Field
		DeadRate    respjson.Field
		IssueScore  respjson.Field
		PageKey     respjson.Field
		RageClicks  respjson.Field
		RageRate    respjson.Field
		TotalClicks respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HeatmapPageListResponse) RawJSON() string { return r.JSON.raw }
func (r *HeatmapPageListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HeatmapPageListResponseBreakpoint struct {
	Breakpoint string `json:"breakpoint" api:"required"`
	Clicks     int64  `json:"clicks" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Breakpoint  respjson.Field
		Clicks      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HeatmapPageListResponseBreakpoint) RawJSON() string { return r.JSON.raw }
func (r *HeatmapPageListResponseBreakpoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HeatmapPageSummaryResponse struct {
	// Aggregated click positions for the requested breakpoint, bucketed into a 64×64
	// grid normalized to the page viewport.
	ClickBins []HeatmapPageSummaryResponseClickBin `json:"clickBins" api:"required"`
	// Click positions where no interactive element was hit. `topElement` is the
	// most-frequently-clicked non-interactive ancestor tag at that bucket, if
	// attributable.
	DeadClicks []HeatmapPageSummaryResponseDeadClick `json:"deadClicks" api:"required"`
	// Click positions where repeated clicks were detected within a short window.
	RageClicks []HeatmapPageSummaryResponseRageClick `json:"rageClicks" api:"required"`
	// Up to 5 curated replay links for this page, ranked by engagement and weighted to
	// cover multiple viewport breakpoints when possible. Each entry is an absolute URL
	// into the OursPrivacy web app — open it to view the recorded session.
	Replays []HeatmapPageSummaryResponseReplay `json:"replays" api:"required"`
	// Distribution of how far visitors scrolled, bucketed in percent of page height.
	// `sessions` is the count of distinct sessions that reached at least that bucket.
	ScrollDepth []HeatmapPageSummaryResponseScrollDepth `json:"scrollDepth" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClickBins   respjson.Field
		DeadClicks  respjson.Field
		RageClicks  respjson.Field
		Replays     respjson.Field
		ScrollDepth respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HeatmapPageSummaryResponse) RawJSON() string { return r.JSON.raw }
func (r *HeatmapPageSummaryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HeatmapPageSummaryResponseClickBin struct {
	BinX   int64 `json:"binX" api:"required"`
	BinY   int64 `json:"binY" api:"required"`
	Clicks int64 `json:"clicks" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BinX        respjson.Field
		BinY        respjson.Field
		Clicks      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HeatmapPageSummaryResponseClickBin) RawJSON() string { return r.JSON.raw }
func (r *HeatmapPageSummaryResponseClickBin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HeatmapPageSummaryResponseDeadClick struct {
	BinX       int64  `json:"binX" api:"required"`
	BinY       int64  `json:"binY" api:"required"`
	DeadClicks int64  `json:"deadClicks" api:"required"`
	TopElement string `json:"topElement" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BinX        respjson.Field
		BinY        respjson.Field
		DeadClicks  respjson.Field
		TopElement  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HeatmapPageSummaryResponseDeadClick) RawJSON() string { return r.JSON.raw }
func (r *HeatmapPageSummaryResponseDeadClick) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HeatmapPageSummaryResponseRageClick struct {
	BinX        int64 `json:"binX" api:"required"`
	BinY        int64 `json:"binY" api:"required"`
	RageEvents  int64 `json:"rageEvents" api:"required"`
	TotalClicks int64 `json:"totalClicks" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BinX        respjson.Field
		BinY        respjson.Field
		RageEvents  respjson.Field
		TotalClicks respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HeatmapPageSummaryResponseRageClick) RawJSON() string { return r.JSON.raw }
func (r *HeatmapPageSummaryResponseRageClick) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HeatmapPageSummaryResponseReplay struct {
	// Viewport bucket the session was recorded under.
	//
	// Any of "desktop", "mobile", "tablet".
	Breakpoint string `json:"breakpoint" api:"required"`
	// Absolute link to view the session replay in the OursPrivacy web app. Opens the
	// replayer pre-scoped to this session, visitor, and date.
	URL string `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Breakpoint  respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HeatmapPageSummaryResponseReplay) RawJSON() string { return r.JSON.raw }
func (r *HeatmapPageSummaryResponseReplay) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HeatmapPageSummaryResponseScrollDepth struct {
	Bucket   int64 `json:"bucket" api:"required"`
	Sessions int64 `json:"sessions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bucket      respjson.Field
		Sessions    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HeatmapPageSummaryResponseScrollDepth) RawJSON() string { return r.JSON.raw }
func (r *HeatmapPageSummaryResponseScrollDepth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HeatmapPageListParams struct {
	// Inclusive lower bound of the heatmap window, as a UTC calendar day in
	// `YYYY-MM-DD` format. The window between `from` and `to` must be 60 days or
	// fewer.
	From string `query:"from" api:"required" json:"-"`
	// Inclusive upper bound of the heatmap window, as a UTC calendar day in
	// `YYYY-MM-DD` format. The window between `from` and `to` must be 60 days or
	// fewer.
	To string `query:"to" api:"required" json:"-"`
	// Maximum number of items to return. Defaults to 25; values below 1 are clamped to
	// 1 and values above 100 are clamped to 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by browser name as captured on events.
	BrowserName param.Opt[string] `query:"browserName,omitzero" json:"-"`
	// Filter by visitor country (ISO country name or code as stored on events).
	Country param.Opt[string] `query:"country,omitzero" json:"-"`
	// Opaque pagination cursor from pagination.nextCursor in the previous response. Do
	// not decode or modify it. Malformed cursors return 400 Bad Request.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Filter by visitor region (state/province as stored on events).
	Region param.Opt[string] `query:"region,omitzero" json:"-"`
	// Case-insensitive substring match against `pageKey`.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Sort key. Defaults to `CLICKS` (descending).
	//
	// Any of "CLICKS", "DEAD_RATE", "ISSUE_SCORE", "RAGE_RATE".
	SortBy HeatmapPageListParamsSortBy `query:"sortBy,omitzero" json:"-"`
	// Sort direction. Defaults to `DESC`.
	//
	// Any of "ASC", "DESC".
	SortDir HeatmapPageListParamsSortDir `query:"sortDir,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [HeatmapPageListParams]'s query parameters as `url.Values`.
func (r HeatmapPageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort key. Defaults to `CLICKS` (descending).
type HeatmapPageListParamsSortBy string

const (
	HeatmapPageListParamsSortByClicks     HeatmapPageListParamsSortBy = "CLICKS"
	HeatmapPageListParamsSortByDeadRate   HeatmapPageListParamsSortBy = "DEAD_RATE"
	HeatmapPageListParamsSortByIssueScore HeatmapPageListParamsSortBy = "ISSUE_SCORE"
	HeatmapPageListParamsSortByRageRate   HeatmapPageListParamsSortBy = "RAGE_RATE"
)

// Sort direction. Defaults to `DESC`.
type HeatmapPageListParamsSortDir string

const (
	HeatmapPageListParamsSortDirAsc  HeatmapPageListParamsSortDir = "ASC"
	HeatmapPageListParamsSortDirDesc HeatmapPageListParamsSortDir = "DESC"
)

type HeatmapPageSummaryParams struct {
	// Viewport bucket the click, dead-click, rage, and scroll-depth aggregations are
	// computed for. Replays are returned for all breakpoints regardless of this value
	// so callers can compare across devices.
	//
	// Any of "desktop", "mobile", "tablet".
	Breakpoint HeatmapPageSummaryParamsBreakpoint `query:"breakpoint,omitzero" api:"required" json:"-"`
	// Inclusive lower bound of the heatmap window, as a UTC calendar day in
	// `YYYY-MM-DD` format. The window between `from` and `to` must be 60 days or
	// fewer.
	From string `query:"from" api:"required" json:"-"`
	// Page identifier returned by `GET /rest/v1/heatmap-pages`. Origin + pathname with
	// the query string stripped (e.g. `https://example.com/pricing`).
	PageKey string `query:"pageKey" api:"required" json:"-"`
	// Inclusive upper bound of the heatmap window, as a UTC calendar day in
	// `YYYY-MM-DD` format. The window between `from` and `to` must be 60 days or
	// fewer.
	To string `query:"to" api:"required" json:"-"`
	// Filter by browser name as captured on events.
	BrowserName param.Opt[string] `query:"browserName,omitzero" json:"-"`
	// Filter by visitor country (ISO country name or code as stored on events).
	Country param.Opt[string] `query:"country,omitzero" json:"-"`
	// Filter by visitor region (state/province as stored on events).
	Region param.Opt[string] `query:"region,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [HeatmapPageSummaryParams]'s query parameters as
// `url.Values`.
func (r HeatmapPageSummaryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Viewport bucket the click, dead-click, rage, and scroll-depth aggregations are
// computed for. Replays are returned for all breakpoints regardless of this value
// so callers can compare across devices.
type HeatmapPageSummaryParamsBreakpoint string

const (
	HeatmapPageSummaryParamsBreakpointDesktop HeatmapPageSummaryParamsBreakpoint = "desktop"
	HeatmapPageSummaryParamsBreakpointMobile  HeatmapPageSummaryParamsBreakpoint = "mobile"
	HeatmapPageSummaryParamsBreakpointTablet  HeatmapPageSummaryParamsBreakpoint = "tablet"
)
