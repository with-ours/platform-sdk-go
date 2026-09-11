// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/with-ours/platform-sdk-go/internal/apijson"
	"github.com/with-ours/platform-sdk-go/internal/apiquery"
	"github.com/with-ours/platform-sdk-go/internal/requestconfig"
	"github.com/with-ours/platform-sdk-go/option"
	"github.com/with-ours/platform-sdk-go/packages/param"
	"github.com/with-ours/platform-sdk-go/packages/respjson"
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

// Return privacy-first traffic metrics and a timeseries for the requested date
// range. Filter by source, geography, page, campaign, device, or other supported
// dimensions with the JSON-encoded `filters` query parameter. Requires scope:
// web-analytics:view
func (r *WebAnalyticsService) Overview(ctx context.Context, query WebAnalyticsOverviewParams, opts ...option.RequestOption) (res *WebAnalyticsOverviewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/web-analytics/overview"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Return visitor counts grouped by referrer or UTM source dimension for the
// requested date range. Requires scope: web-analytics:view
func (r *WebAnalyticsService) Sources(ctx context.Context, query WebAnalyticsSourcesParams, opts ...option.RequestOption) (res *WebAnalyticsSourcesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/web-analytics/sources"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Return page-level traffic metrics for top pages, entry pages, or exit pages in
// the requested date range. Requires scope: web-analytics:view
func (r *WebAnalyticsService) Pages(ctx context.Context, query WebAnalyticsPagesParams, opts ...option.RequestOption) (res *WebAnalyticsPagesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/web-analytics/pages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Return visitor counts grouped by country, region, or city for the requested date
// range. Requires scope: web-analytics:view
func (r *WebAnalyticsService) Locations(ctx context.Context, query WebAnalyticsLocationsParams, opts ...option.RequestOption) (res *WebAnalyticsLocationsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/web-analytics/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Return visitor counts grouped by device type, browser, or operating system for
// the requested date range. Requires scope: web-analytics:view
func (r *WebAnalyticsService) Devices(ctx context.Context, query WebAnalyticsDevicesParams, opts ...option.RequestOption) (res *WebAnalyticsDevicesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/web-analytics/devices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Return the distinct visitors active in the most recent 15-minute window,
// optionally scoped to one web source. Requires scope: web-analytics:view
func (r *WebAnalyticsService) CurrentVisitors(ctx context.Context, query WebAnalyticsCurrentVisitorsParams, opts ...option.RequestOption) (res *WebAnalyticsCurrentVisitorsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/web-analytics/current-visitors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Return the next or previous journey steps for a pinned path. The `path` and
// `filters` query parameters are JSON-encoded arrays. Requires scope:
// web-analytics:view
func (r *WebAnalyticsService) Journey(ctx context.Context, query WebAnalyticsJourneyParams, opts ...option.RequestOption) (res *WebAnalyticsJourneyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/web-analytics/journey"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
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

type WebAnalyticsJourneyResponse struct {
	AnchorSessions int64                             `json:"anchorSessions" api:"required"`
	HasMore        bool                              `json:"hasMore" api:"required"`
	Steps          []WebAnalyticsJourneyResponseStep `json:"steps" api:"required"`
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
func (r WebAnalyticsJourneyResponse) RawJSON() string { return r.JSON.raw }
func (r *WebAnalyticsJourneyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebAnalyticsJourneyResponseStep struct {
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
func (r WebAnalyticsJourneyResponseStep) RawJSON() string { return r.JSON.raw }
func (r *WebAnalyticsJourneyResponseStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
	// Optional JSON-encoded array of up to 20 filters. Dimensions: `page`,
	// `entry_page`, `exit_page`, `source`, `medium`, `campaign`, `content`, `term`,
	// `referrer`, `country`, `region`, `city`, `device`, `browser`, `os`. Each filter
	// has a dimension, optional operator (`IS`, `IS_NOT`, `CONTAINS`, `NOT_CONTAINS`;
	// defaults to `IS`), and one or more values. Example:
	// `[{"dimension":"country","values":["United States"]}]`.
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
	// Optional JSON-encoded array of up to 20 filters. Dimensions: `page`,
	// `entry_page`, `exit_page`, `source`, `medium`, `campaign`, `content`, `term`,
	// `referrer`, `country`, `region`, `city`, `device`, `browser`, `os`. Each filter
	// has a dimension, optional operator (`IS`, `IS_NOT`, `CONTAINS`, `NOT_CONTAINS`;
	// defaults to `IS`), and one or more values. Example:
	// `[{"dimension":"country","values":["United States"]}]`.
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
	// Optional JSON-encoded array of up to 20 filters. Dimensions: `page`,
	// `entry_page`, `exit_page`, `source`, `medium`, `campaign`, `content`, `term`,
	// `referrer`, `country`, `region`, `city`, `device`, `browser`, `os`. Each filter
	// has a dimension, optional operator (`IS`, `IS_NOT`, `CONTAINS`, `NOT_CONTAINS`;
	// defaults to `IS`), and one or more values. Example:
	// `[{"dimension":"country","values":["United States"]}]`.
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
	// Optional JSON-encoded array of up to 20 filters. Dimensions: `page`,
	// `entry_page`, `exit_page`, `source`, `medium`, `campaign`, `content`, `term`,
	// `referrer`, `country`, `region`, `city`, `device`, `browser`, `os`. Each filter
	// has a dimension, optional operator (`IS`, `IS_NOT`, `CONTAINS`, `NOT_CONTAINS`;
	// defaults to `IS`), and one or more values. Example:
	// `[{"dimension":"country","values":["United States"]}]`.
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
	// Optional JSON-encoded array of up to 20 filters. Dimensions: `page`,
	// `entry_page`, `exit_page`, `source`, `medium`, `campaign`, `content`, `term`,
	// `referrer`, `country`, `region`, `city`, `device`, `browser`, `os`. Each filter
	// has a dimension, optional operator (`IS`, `IS_NOT`, `CONTAINS`, `NOT_CONTAINS`;
	// defaults to `IS`), and one or more values. Example:
	// `[{"dimension":"country","values":["United States"]}]`.
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

type WebAnalyticsJourneyParams struct {
	// Inclusive lower bound of the analysis window as `YYYY-MM-DD`.
	From time.Time `query:"from" api:"required" format:"date" json:"-"`
	// JSON-encoded ordered path of opaque journey step keys. Use an empty array to
	// request first-column candidates.
	Path string `query:"path" api:"required" json:"-"`
	// Inclusive upper bound of the analysis window as `YYYY-MM-DD`.
	To time.Time `query:"to" api:"required" format:"date" json:"-"`
	// Exclude detected bot sessions. Defaults to true.
	ExcludeBots param.Opt[bool] `query:"excludeBots,omitzero" json:"-"`
	// Optional JSON-encoded array of up to 20 journey filters. Supports web analytics
	// dimensions plus `event_name`, `ep_currency`, `ep_appointment_id`,
	// `ep_appointment_status`, `ep_service_line`, `ep_provider_id`, `ep_location_id`,
	// `ep_booking_channel`, `ep_revenue_type`, `ep_call_outcome`, and `ep_staff_id`.
	// Each filter has a dimension, optional operator (`IS`, `IS_NOT`, `CONTAINS`,
	// `NOT_CONTAINS`; defaults to `IS`), and one or more values.
	Filters param.Opt[string] `query:"filters,omitzero" json:"-"`
	Limit   param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Search  param.Opt[string] `query:"search,omitzero" json:"-"`
	// Optional web source UUID. Omit to aggregate all web sources in the account.
	WebSourceID param.Opt[string] `query:"webSourceId,omitzero" format:"uuid" json:"-"`
	// Any of "forward", "reverse".
	Direction WebAnalyticsJourneyParamsDirection `query:"direction,omitzero" json:"-"`
	// Any of "PAGE", "EVENT".
	StepKind WebAnalyticsJourneyParamsStepKind `query:"stepKind,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebAnalyticsJourneyParams]'s query parameters as
// `url.Values`.
func (r WebAnalyticsJourneyParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebAnalyticsJourneyParamsDirection string

const (
	WebAnalyticsJourneyParamsDirectionForward WebAnalyticsJourneyParamsDirection = "forward"
	WebAnalyticsJourneyParamsDirectionReverse WebAnalyticsJourneyParamsDirection = "reverse"
)

type WebAnalyticsJourneyParamsStepKind string

const (
	WebAnalyticsJourneyParamsStepKindPage  WebAnalyticsJourneyParamsStepKind = "PAGE"
	WebAnalyticsJourneyParamsStepKindEvent WebAnalyticsJourneyParamsStepKind = "EVENT"
)
