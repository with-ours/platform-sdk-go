// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/with-ours/platform-sdk-go/v2/internal/apijson"
	"github.com/with-ours/platform-sdk-go/v2/internal/apiquery"
	"github.com/with-ours/platform-sdk-go/v2/internal/requestconfig"
	"github.com/with-ours/platform-sdk-go/v2/option"
	"github.com/with-ours/platform-sdk-go/v2/packages/param"
	"github.com/with-ours/platform-sdk-go/v2/packages/respjson"
)

// WebAnalyticsService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebAnalyticsService] method instead.
type WebAnalyticsService struct {
	Options []option.RequestOption
}

// NewWebAnalyticsService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebAnalyticsService(opts ...option.RequestOption) (r WebAnalyticsService) {
	r = WebAnalyticsService{}
	r.Options = opts
	return
}

// Read a fixed Overview report for up to 60 inclusive UTC days. TRAFFIC measures
// visitors with page views and page views; BEHAVIOR measures tracked clicks, dead
// clicks, and rage clicks. Both include web sources only and use the same counting
// and bot/test exclusions as Insights. Requires Analytics access and
// web-analytics:view; no custom-variable permission is required. Requires API-key
// scope or current OAuth user permission: web-analytics:view
func (r *WebAnalyticsService) OverviewInsight(ctx context.Context, query WebAnalyticsOverviewInsightParams, opts ...option.RequestOption) (res *WebAnalyticsOverviewInsightResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/web-analytics/overview-insight"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Return privacy-first traffic metrics and a timeseries for the requested date
// range. Filter by source, geography, page, campaign, device, or other supported
// dimensions with the JSON-encoded `filters` query parameter. Requires API-key
// scope or current OAuth user permission: web-analytics:view
func (r *WebAnalyticsService) Overview(ctx context.Context, query WebAnalyticsOverviewParams, opts ...option.RequestOption) (res *WebAnalyticsOverviewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/web-analytics/overview"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Return visitor counts grouped by referrer or UTM source dimension for the
// requested date range. Requires API-key scope or current OAuth user permission:
// web-analytics:view
func (r *WebAnalyticsService) Sources(ctx context.Context, query WebAnalyticsSourcesParams, opts ...option.RequestOption) (res *WebAnalyticsSourcesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/web-analytics/sources"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Return page-level traffic metrics for top pages, entry pages, or exit pages in
// the requested date range. Requires API-key scope or current OAuth user
// permission: web-analytics:view
func (r *WebAnalyticsService) Pages(ctx context.Context, query WebAnalyticsPagesParams, opts ...option.RequestOption) (res *WebAnalyticsPagesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/web-analytics/pages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Return visitor counts grouped by country, region, or city for the requested date
// range. Requires API-key scope or current OAuth user permission:
// web-analytics:view
func (r *WebAnalyticsService) Locations(ctx context.Context, query WebAnalyticsLocationsParams, opts ...option.RequestOption) (res *WebAnalyticsLocationsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/web-analytics/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Return visitor counts grouped by device type, browser, or operating system for
// the requested date range. Requires API-key scope or current OAuth user
// permission: web-analytics:view
func (r *WebAnalyticsService) Devices(ctx context.Context, query WebAnalyticsDevicesParams, opts ...option.RequestOption) (res *WebAnalyticsDevicesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/web-analytics/devices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Return the distinct visitors active in the most recent 15-minute window,
// optionally scoped to one web source. Requires API-key scope or current OAuth
// user permission: web-analytics:view
func (r *WebAnalyticsService) CurrentVisitors(ctx context.Context, query WebAnalyticsCurrentVisitorsParams, opts ...option.RequestOption) (res *WebAnalyticsCurrentVisitorsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/web-analytics/current-visitors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Return the next or previous journey steps for a pinned path. The `path` and
// `filters` query parameters are JSON-encoded arrays. Requires API-key scope or
// current OAuth user permission: web-analytics:view
func (r *WebAnalyticsService) PathExplorer(ctx context.Context, query WebAnalyticsPathExplorerParams, opts ...option.RequestOption) (res *WebAnalyticsPathExplorerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/web-analytics/path-explorer"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type WebAnalyticsOverviewInsightResponse struct {
	// Any of "clickhouse".
	Engine WebAnalyticsOverviewInsightResponseEngine `json:"engine" api:"required"`
	// Saved report identifier; null for previews.
	ReportID string `json:"reportId" api:"required" format:"uuid"`
	// Saved report revision; null for previews.
	Revision string                                   `json:"revision" api:"required" format:"uuid"`
	Scope    WebAnalyticsOverviewInsightResponseScope `json:"scope" api:"required"`
	// Any of 1.
	SemanticVersion   float64                                     `json:"semanticVersion" api:"required"`
	Series            []WebAnalyticsOverviewInsightResponseSeries `json:"series" api:"required"`
	SourceLimitations []string                                    `json:"sourceLimitations" api:"required"`
	SourcePolicy      string                                      `json:"sourcePolicy" api:"required"`
	SourceRevision    string                                      `json:"sourceRevision" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Engine            respjson.Field
		ReportID          respjson.Field
		Revision          respjson.Field
		Scope             respjson.Field
		SemanticVersion   respjson.Field
		Series            respjson.Field
		SourceLimitations respjson.Field
		SourcePolicy      respjson.Field
		SourceRevision    respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebAnalyticsOverviewInsightResponse) RawJSON() string { return r.JSON.raw }
func (r *WebAnalyticsOverviewInsightResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticsOverviewInsightResponseEngine string

const (
	WebAnalyticsOverviewInsightResponseEngineClickhouse WebAnalyticsOverviewInsightResponseEngine = "clickhouse"
)

type WebAnalyticsOverviewInsightResponseScope struct {
	Cumulative bool      `json:"cumulative" api:"required"`
	ExecutedAt time.Time `json:"executedAt" api:"required" format:"date-time"`
	From       string    `json:"from" api:"required"`
	// Any of "day", "week", "month".
	Interval string `json:"interval" api:"required"`
	// Any of "unknown".
	SourceCompleteness string `json:"sourceCompleteness" api:"required"`
	// Any of "UTC".
	TimeZone   string                                             `json:"timeZone" api:"required"`
	To         string                                             `json:"to" api:"required"`
	Comparison WebAnalyticsOverviewInsightResponseScopeComparison `json:"comparison" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cumulative         respjson.Field
		ExecutedAt         respjson.Field
		From               respjson.Field
		Interval           respjson.Field
		SourceCompleteness respjson.Field
		TimeZone           respjson.Field
		To                 respjson.Field
		Comparison         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebAnalyticsOverviewInsightResponseScope) RawJSON() string { return r.JSON.raw }
func (r *WebAnalyticsOverviewInsightResponseScope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticsOverviewInsightResponseScopeComparison struct {
	From string `json:"from" api:"required"`
	// Any of "previous_period", "previous_year", "custom".
	Kind string `json:"kind" api:"required"`
	To   string `json:"to" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		Kind        respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebAnalyticsOverviewInsightResponseScopeComparison) RawJSON() string { return r.JSON.raw }
func (r *WebAnalyticsOverviewInsightResponseScopeComparison) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticsOverviewInsightResponseSeries struct {
	// Any of "event_count", "unique_visitors", "sum", "average".
	Aggregation     string                                                    `json:"aggregation" api:"required"`
	BreakdownValues []WebAnalyticsOverviewInsightResponseSeriesBreakdownValue `json:"breakdownValues" api:"required"`
	// Any of "previous_period", "previous_year", "custom".
	Comparison string                                           `json:"comparison" api:"required"`
	Key        string                                           `json:"key" api:"required"`
	Label      string                                           `json:"label" api:"required"`
	MetricID   string                                           `json:"metricId" api:"required" format:"uuid"`
	Points     []WebAnalyticsOverviewInsightResponseSeriesPoint `json:"points" api:"required"`
	RangeValue float64                                          `json:"rangeValue" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Aggregation     respjson.Field
		BreakdownValues respjson.Field
		Comparison      respjson.Field
		Key             respjson.Field
		Label           respjson.Field
		MetricID        respjson.Field
		Points          respjson.Field
		RangeValue      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebAnalyticsOverviewInsightResponseSeries) RawJSON() string { return r.JSON.raw }
func (r *WebAnalyticsOverviewInsightResponseSeries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticsOverviewInsightResponseSeriesBreakdownValue struct {
	Property string `json:"property" api:"required"`
	Value    string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Property    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebAnalyticsOverviewInsightResponseSeriesBreakdownValue) RawJSON() string { return r.JSON.raw }
func (r *WebAnalyticsOverviewInsightResponseSeriesBreakdownValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticsOverviewInsightResponseSeriesPoint struct {
	BucketEnd   time.Time `json:"bucketEnd" api:"required" format:"date-time"`
	BucketStart time.Time `json:"bucketStart" api:"required" format:"date-time"`
	Cumulative  float64   `json:"cumulative" api:"required"`
	IsPartial   bool      `json:"isPartial" api:"required"`
	Raw         float64   `json:"raw" api:"required"`
	Smoothed    float64   `json:"smoothed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BucketEnd   respjson.Field
		BucketStart respjson.Field
		Cumulative  respjson.Field
		IsPartial   respjson.Field
		Raw         respjson.Field
		Smoothed    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebAnalyticsOverviewInsightResponseSeriesPoint) RawJSON() string { return r.JSON.raw }
func (r *WebAnalyticsOverviewInsightResponseSeriesPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticsOverviewResponse struct {
	DataUpdatedAt string                                 `json:"dataUpdatedAt" api:"required"`
	Metrics       WebAnalyticsOverviewResponseMetrics    `json:"metrics" api:"required"`
	Timeseries    []WebAnalyticsOverviewResponseTimesery `json:"timeseries" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataUpdatedAt respjson.Field
		Metrics       respjson.Field
		Timeseries    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebAnalyticsOverviewResponse) RawJSON() string { return r.JSON.raw }
func (r *WebAnalyticsOverviewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticsOverviewResponseMetrics struct {
	// Percentage from 0 to 100 of visits with exactly one pageview.
	BounceRate     float64 `json:"bounceRate" api:"required"`
	Pageviews      int64   `json:"pageviews" api:"required"`
	TotalVisits    int64   `json:"totalVisits" api:"required"`
	UniqueVisitors int64   `json:"uniqueVisitors" api:"required"`
	// Average pageviews per visit.
	ViewsPerVisit float64 `json:"viewsPerVisit" api:"required"`
	// Average visit duration in seconds.
	VisitDuration float64 `json:"visitDuration" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BounceRate     respjson.Field
		Pageviews      respjson.Field
		TotalVisits    respjson.Field
		UniqueVisitors respjson.Field
		ViewsPerVisit  respjson.Field
		VisitDuration  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebAnalyticsOverviewResponseMetrics) RawJSON() string { return r.JSON.raw }
func (r *WebAnalyticsOverviewResponseMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticsOverviewResponseTimesery struct {
	Date  string  `json:"date" api:"required"`
	Value float64 `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebAnalyticsOverviewResponseTimesery) RawJSON() string { return r.JSON.raw }
func (r *WebAnalyticsOverviewResponseTimesery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticsSourcesResponse struct {
	DataUpdatedAt string                           `json:"dataUpdatedAt" api:"required"`
	Rows          []WebAnalyticsSourcesResponseRow `json:"rows" api:"required"`
	TotalCount    int64                            `json:"totalCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataUpdatedAt respjson.Field
		Rows          respjson.Field
		TotalCount    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebAnalyticsSourcesResponse) RawJSON() string { return r.JSON.raw }
func (r *WebAnalyticsSourcesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticsSourcesResponseRow struct {
	Name     string `json:"name" api:"required"`
	Visitors int64  `json:"visitors" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Visitors    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebAnalyticsSourcesResponseRow) RawJSON() string { return r.JSON.raw }
func (r *WebAnalyticsSourcesResponseRow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticsPagesResponse struct {
	DataUpdatedAt string                         `json:"dataUpdatedAt" api:"required"`
	Rows          []WebAnalyticsPagesResponseRow `json:"rows" api:"required"`
	TotalCount    int64                          `json:"totalCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataUpdatedAt respjson.Field
		Rows          respjson.Field
		TotalCount    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebAnalyticsPagesResponse) RawJSON() string { return r.JSON.raw }
func (r *WebAnalyticsPagesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticsPagesResponseRow struct {
	// Percentage from 0 to 100 of visits to the page that bounced.
	BounceRate float64 `json:"bounceRate" api:"required"`
	Entries    int64   `json:"entries" api:"required"`
	// Percentage from 0 to 100 of pageviews that ended a visit.
	ExitRate     float64 `json:"exitRate" api:"required"`
	Exits        int64   `json:"exits" api:"required"`
	PageHostname string  `json:"pageHostname" api:"required"`
	PagePath     string  `json:"pagePath" api:"required"`
	Pageviews    int64   `json:"pageviews" api:"required"`
	// Average time on page in seconds when available.
	TimeOnPage float64 `json:"timeOnPage" api:"required"`
	Visitors   int64   `json:"visitors" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BounceRate   respjson.Field
		Entries      respjson.Field
		ExitRate     respjson.Field
		Exits        respjson.Field
		PageHostname respjson.Field
		PagePath     respjson.Field
		Pageviews    respjson.Field
		TimeOnPage   respjson.Field
		Visitors     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebAnalyticsPagesResponseRow) RawJSON() string { return r.JSON.raw }
func (r *WebAnalyticsPagesResponseRow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticsLocationsResponse struct {
	DataUpdatedAt string                             `json:"dataUpdatedAt" api:"required"`
	Rows          []WebAnalyticsLocationsResponseRow `json:"rows" api:"required"`
	TotalCount    int64                              `json:"totalCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataUpdatedAt respjson.Field
		Rows          respjson.Field
		TotalCount    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebAnalyticsLocationsResponse) RawJSON() string { return r.JSON.raw }
func (r *WebAnalyticsLocationsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticsLocationsResponseRow struct {
	Code     string `json:"code" api:"required"`
	Name     string `json:"name" api:"required"`
	Visitors int64  `json:"visitors" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Name        respjson.Field
		Visitors    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebAnalyticsLocationsResponseRow) RawJSON() string { return r.JSON.raw }
func (r *WebAnalyticsLocationsResponseRow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticsDevicesResponse struct {
	DataUpdatedAt string                           `json:"dataUpdatedAt" api:"required"`
	Rows          []WebAnalyticsDevicesResponseRow `json:"rows" api:"required"`
	TotalCount    int64                            `json:"totalCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataUpdatedAt respjson.Field
		Rows          respjson.Field
		TotalCount    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebAnalyticsDevicesResponse) RawJSON() string { return r.JSON.raw }
func (r *WebAnalyticsDevicesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticsDevicesResponseRow struct {
	Name     string `json:"name" api:"required"`
	Visitors int64  `json:"visitors" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Visitors    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebAnalyticsDevicesResponseRow) RawJSON() string { return r.JSON.raw }
func (r *WebAnalyticsDevicesResponseRow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticsCurrentVisitorsResponse struct {
	Count int64 `json:"count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebAnalyticsCurrentVisitorsResponse) RawJSON() string { return r.JSON.raw }
func (r *WebAnalyticsCurrentVisitorsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticsPathExplorerResponse struct {
	AnchorSessions int64                                  `json:"anchorSessions" api:"required"`
	HasMore        bool                                   `json:"hasMore" api:"required"`
	Steps          []WebAnalyticsPathExplorerResponseStep `json:"steps" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AnchorSessions respjson.Field
		HasMore        respjson.Field
		Steps          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebAnalyticsPathExplorerResponse) RawJSON() string { return r.JSON.raw }
func (r *WebAnalyticsPathExplorerResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticsPathExplorerResponseStep struct {
	IsOther    bool   `json:"isOther" api:"required"`
	IsTerminal bool   `json:"isTerminal" api:"required"`
	Key        string `json:"key" api:"required"`
	// Any of "PAGE", "EVENT".
	Kind     string `json:"kind" api:"required"`
	Label    string `json:"label" api:"required"`
	Sessions int64  `json:"sessions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsOther     respjson.Field
		IsTerminal  respjson.Field
		Key         respjson.Field
		Kind        respjson.Field
		Label       respjson.Field
		Sessions    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebAnalyticsPathExplorerResponseStep) RawJSON() string { return r.JSON.raw }
func (r *WebAnalyticsPathExplorerResponseStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticsOverviewInsightParams struct {
	// First included UTC calendar day (YYYY-MM-DD).
	From time.Time `query:"from" api:"required" format:"date" json:"-"`
	// Built-in traffic or click-behavior report.
	//
	// Any of "TRAFFIC", "BEHAVIOR".
	Preset WebAnalyticsOverviewInsightParamsPreset `query:"preset,omitzero" api:"required" json:"-"`
	// Last included UTC calendar day, at most 60 days including from.
	To time.Time `query:"to" api:"required" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [WebAnalyticsOverviewInsightParams]'s query parameters as
// `url.Values`.
func (r WebAnalyticsOverviewInsightParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Built-in traffic or click-behavior report.
type WebAnalyticsOverviewInsightParamsPreset string

const (
	WebAnalyticsOverviewInsightParamsPresetTraffic  WebAnalyticsOverviewInsightParamsPreset = "TRAFFIC"
	WebAnalyticsOverviewInsightParamsPresetBehavior WebAnalyticsOverviewInsightParamsPreset = "BEHAVIOR"
)

type WebAnalyticsOverviewParams struct {
	// Inclusive lower bound of the analysis window as `YYYY-MM-DD`.
	From time.Time `query:"from" api:"required" format:"date" json:"-"`
	// Timeseries bucket interval. Minute queries are capped to the most recent 24
	// hours.
	//
	// Any of "minute", "day", "week", "month".
	Interval WebAnalyticsOverviewParamsInterval `query:"interval,omitzero" api:"required" json:"-"`
	// Any of "unique_visitors", "total_visits", "pageviews", "views_per_visit",
	// "bounce_rate", "visit_duration".
	Metric WebAnalyticsOverviewParamsMetric `query:"metric,omitzero" api:"required" json:"-"`
	// Inclusive upper bound of the analysis window as `YYYY-MM-DD`.
	To time.Time `query:"to" api:"required" format:"date" json:"-"`
	// Exclude detected bot sessions. Defaults to true.
	ExcludeBots param.Opt[bool] `query:"excludeBots,omitzero" json:"-"`
	// JSON-encoded query input.
	Filters param.Opt[string] `query:"filters,omitzero" json:"-"`
	// Optional ISO timestamp used as the lower bound for realtime queries.
	RealtimeFrom param.Opt[time.Time] `query:"realtimeFrom,omitzero" format:"date-time" json:"-"`
	// Optional web source UUID. Omit to aggregate all web sources in the account.
	WebSourceID param.Opt[string] `query:"webSourceId,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [WebAnalyticsOverviewParams]'s query parameters as
// `url.Values`.
func (r WebAnalyticsOverviewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Timeseries bucket interval. Minute queries are capped to the most recent 24
// hours.
type WebAnalyticsOverviewParamsInterval string

const (
	WebAnalyticsOverviewParamsIntervalMinute WebAnalyticsOverviewParamsInterval = "minute"
	WebAnalyticsOverviewParamsIntervalDay    WebAnalyticsOverviewParamsInterval = "day"
	WebAnalyticsOverviewParamsIntervalWeek   WebAnalyticsOverviewParamsInterval = "week"
	WebAnalyticsOverviewParamsIntervalMonth  WebAnalyticsOverviewParamsInterval = "month"
)

type WebAnalyticsOverviewParamsMetric string

const (
	WebAnalyticsOverviewParamsMetricUniqueVisitors WebAnalyticsOverviewParamsMetric = "unique_visitors"
	WebAnalyticsOverviewParamsMetricTotalVisits    WebAnalyticsOverviewParamsMetric = "total_visits"
	WebAnalyticsOverviewParamsMetricPageviews      WebAnalyticsOverviewParamsMetric = "pageviews"
	WebAnalyticsOverviewParamsMetricViewsPerVisit  WebAnalyticsOverviewParamsMetric = "views_per_visit"
	WebAnalyticsOverviewParamsMetricBounceRate     WebAnalyticsOverviewParamsMetric = "bounce_rate"
	WebAnalyticsOverviewParamsMetricVisitDuration  WebAnalyticsOverviewParamsMetric = "visit_duration"
)

type WebAnalyticsSourcesParams struct {
	// Any of "referrer", "campaign", "source", "medium", "content", "term".
	Dimension WebAnalyticsSourcesParamsDimension `query:"dimension,omitzero" api:"required" json:"-"`
	// Inclusive lower bound of the analysis window as `YYYY-MM-DD`.
	From time.Time `query:"from" api:"required" format:"date" json:"-"`
	// Inclusive upper bound of the analysis window as `YYYY-MM-DD`.
	To time.Time `query:"to" api:"required" format:"date" json:"-"`
	// Exclude detected bot sessions. Defaults to true.
	ExcludeBots param.Opt[bool] `query:"excludeBots,omitzero" json:"-"`
	// JSON-encoded query input.
	Filters param.Opt[string] `query:"filters,omitzero" json:"-"`
	// Optional web source UUID. Omit to aggregate all web sources in the account.
	WebSourceID param.Opt[string] `query:"webSourceId,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [WebAnalyticsSourcesParams]'s query parameters as
// `url.Values`.
func (r WebAnalyticsSourcesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebAnalyticsSourcesParamsDimension string

const (
	WebAnalyticsSourcesParamsDimensionReferrer WebAnalyticsSourcesParamsDimension = "referrer"
	WebAnalyticsSourcesParamsDimensionCampaign WebAnalyticsSourcesParamsDimension = "campaign"
	WebAnalyticsSourcesParamsDimensionSource   WebAnalyticsSourcesParamsDimension = "source"
	WebAnalyticsSourcesParamsDimensionMedium   WebAnalyticsSourcesParamsDimension = "medium"
	WebAnalyticsSourcesParamsDimensionContent  WebAnalyticsSourcesParamsDimension = "content"
	WebAnalyticsSourcesParamsDimensionTerm     WebAnalyticsSourcesParamsDimension = "term"
)

type WebAnalyticsPagesParams struct {
	// Inclusive lower bound of the analysis window as `YYYY-MM-DD`.
	From time.Time `query:"from" api:"required" format:"date" json:"-"`
	// Inclusive upper bound of the analysis window as `YYYY-MM-DD`.
	To time.Time `query:"to" api:"required" format:"date" json:"-"`
	// Any of "top", "entry", "exit".
	View WebAnalyticsPagesParamsView `query:"view,omitzero" api:"required" json:"-"`
	// Exclude detected bot sessions. Defaults to true.
	ExcludeBots param.Opt[bool] `query:"excludeBots,omitzero" json:"-"`
	// JSON-encoded query input.
	Filters param.Opt[string] `query:"filters,omitzero" json:"-"`
	// Optional web source UUID. Omit to aggregate all web sources in the account.
	WebSourceID param.Opt[string] `query:"webSourceId,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [WebAnalyticsPagesParams]'s query parameters as
// `url.Values`.
func (r WebAnalyticsPagesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebAnalyticsPagesParamsView string

const (
	WebAnalyticsPagesParamsViewTop   WebAnalyticsPagesParamsView = "top"
	WebAnalyticsPagesParamsViewEntry WebAnalyticsPagesParamsView = "entry"
	WebAnalyticsPagesParamsViewExit  WebAnalyticsPagesParamsView = "exit"
)

type WebAnalyticsLocationsParams struct {
	// Any of "country", "region", "city".
	Dimension WebAnalyticsLocationsParamsDimension `query:"dimension,omitzero" api:"required" json:"-"`
	// Inclusive lower bound of the analysis window as `YYYY-MM-DD`.
	From time.Time `query:"from" api:"required" format:"date" json:"-"`
	// Inclusive upper bound of the analysis window as `YYYY-MM-DD`.
	To time.Time `query:"to" api:"required" format:"date" json:"-"`
	// Exclude detected bot sessions. Defaults to true.
	ExcludeBots param.Opt[bool] `query:"excludeBots,omitzero" json:"-"`
	// JSON-encoded query input.
	Filters param.Opt[string] `query:"filters,omitzero" json:"-"`
	// Optional web source UUID. Omit to aggregate all web sources in the account.
	WebSourceID param.Opt[string] `query:"webSourceId,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [WebAnalyticsLocationsParams]'s query parameters as
// `url.Values`.
func (r WebAnalyticsLocationsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebAnalyticsLocationsParamsDimension string

const (
	WebAnalyticsLocationsParamsDimensionCountry WebAnalyticsLocationsParamsDimension = "country"
	WebAnalyticsLocationsParamsDimensionRegion  WebAnalyticsLocationsParamsDimension = "region"
	WebAnalyticsLocationsParamsDimensionCity    WebAnalyticsLocationsParamsDimension = "city"
)

type WebAnalyticsDevicesParams struct {
	// Any of "device", "browser", "os".
	Dimension WebAnalyticsDevicesParamsDimension `query:"dimension,omitzero" api:"required" json:"-"`
	// Inclusive lower bound of the analysis window as `YYYY-MM-DD`.
	From time.Time `query:"from" api:"required" format:"date" json:"-"`
	// Inclusive upper bound of the analysis window as `YYYY-MM-DD`.
	To time.Time `query:"to" api:"required" format:"date" json:"-"`
	// Exclude detected bot sessions. Defaults to true.
	ExcludeBots param.Opt[bool] `query:"excludeBots,omitzero" json:"-"`
	// JSON-encoded query input.
	Filters param.Opt[string] `query:"filters,omitzero" json:"-"`
	// Optional web source UUID. Omit to aggregate all web sources in the account.
	WebSourceID param.Opt[string] `query:"webSourceId,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [WebAnalyticsDevicesParams]'s query parameters as
// `url.Values`.
func (r WebAnalyticsDevicesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebAnalyticsDevicesParamsDimension string

const (
	WebAnalyticsDevicesParamsDimensionDevice  WebAnalyticsDevicesParamsDimension = "device"
	WebAnalyticsDevicesParamsDimensionBrowser WebAnalyticsDevicesParamsDimension = "browser"
	WebAnalyticsDevicesParamsDimensionOs      WebAnalyticsDevicesParamsDimension = "os"
)

type WebAnalyticsCurrentVisitorsParams struct {
	// Optional web source UUID. Omit to count visitors across all account web sources.
	WebSourceID param.Opt[string] `query:"webSourceId,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [WebAnalyticsCurrentVisitorsParams]'s query parameters as
// `url.Values`.
func (r WebAnalyticsCurrentVisitorsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebAnalyticsPathExplorerParams struct {
	// Inclusive lower bound of the analysis window as `YYYY-MM-DD`.
	From time.Time `query:"from" api:"required" format:"date" json:"-"`
	// JSON-encoded query input.
	Path string `query:"path" api:"required" json:"-"`
	// Inclusive upper bound of the analysis window as `YYYY-MM-DD`.
	To time.Time `query:"to" api:"required" format:"date" json:"-"`
	// Exclude detected bot sessions. Defaults to true.
	ExcludeBots param.Opt[bool] `query:"excludeBots,omitzero" json:"-"`
	// JSON-encoded query input.
	Filters param.Opt[string] `query:"filters,omitzero" json:"-"`
	Limit   param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Search  param.Opt[string] `query:"search,omitzero" json:"-"`
	// Optional web source UUID. Omit to aggregate all web sources in the account.
	WebSourceID param.Opt[string] `query:"webSourceId,omitzero" format:"uuid" json:"-"`
	// Any of "forward", "reverse".
	Direction WebAnalyticsPathExplorerParamsDirection `query:"direction,omitzero" json:"-"`
	// Any of "PAGE", "EVENT".
	StepKind WebAnalyticsPathExplorerParamsStepKind `query:"stepKind,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebAnalyticsPathExplorerParams]'s query parameters as
// `url.Values`.
func (r WebAnalyticsPathExplorerParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebAnalyticsPathExplorerParamsDirection string

const (
	WebAnalyticsPathExplorerParamsDirectionForward WebAnalyticsPathExplorerParamsDirection = "forward"
	WebAnalyticsPathExplorerParamsDirectionReverse WebAnalyticsPathExplorerParamsDirection = "reverse"
)

type WebAnalyticsPathExplorerParamsStepKind string

const (
	WebAnalyticsPathExplorerParamsStepKindPage  WebAnalyticsPathExplorerParamsStepKind = "PAGE"
	WebAnalyticsPathExplorerParamsStepKindEvent WebAnalyticsPathExplorerParamsStepKind = "EVENT"
)
