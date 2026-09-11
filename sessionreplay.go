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

// SessionReplayService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSessionReplayService] method instead.
type SessionReplayService struct {
	Options []option.RequestOption
}

// NewSessionReplayService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSessionReplayService(opts ...option.RequestOption) (r SessionReplayService) {
	r = SessionReplayService{}
	r.Options = opts
	return
}

// List recorded sessions for a date range. Filter by event, page, visitor, UTM
// fields, or an explicit JSON-encoded session ID list. Use `pagination.nextCursor`
// to retrieve the next page. Requires scope: web-analytics:view
func (r *SessionReplayService) List(ctx context.Context, query SessionReplayListParams, opts ...option.RequestOption) (res *SessionReplayListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/session-replays"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Return the total number of replay-bearing sessions and a daily timeseries for
// the requested date range. Requires scope: web-analytics:view
func (r *SessionReplayService) Overview(ctx context.Context, query SessionReplayOverviewParams, opts ...option.RequestOption) (res *SessionReplayOverviewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/session-replays/overview"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type SessionReplayListResponse struct {
	Items      []SessionReplayListResponseItem     `json:"items" api:"required"`
	Pagination SessionReplayListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionReplayListResponse) RawJSON() string { return r.JSON.raw }
func (r *SessionReplayListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionReplayListResponseItem struct {
	Date       string `json:"date" api:"required"`
	Duration   int64  `json:"duration" api:"required"`
	EventCount int64  `json:"eventCount" api:"required"`
	PageCount  int64  `json:"pageCount" api:"required"`
	SessionID  string `json:"sessionId" api:"required"`
	StartTime  string `json:"startTime" api:"required"`
	VisitorID  string `json:"visitorId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		Duration    respjson.Field
		EventCount  respjson.Field
		PageCount   respjson.Field
		SessionID   respjson.Field
		StartTime   respjson.Field
		VisitorID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionReplayListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *SessionReplayListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionReplayListResponsePagination struct {
	HasMore    bool   `json:"hasMore" api:"required"`
	NextCursor string `json:"nextCursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionReplayListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *SessionReplayListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionReplayOverviewResponse struct {
	Timeseries   []SessionReplayOverviewResponseTimesery `json:"timeseries" api:"required"`
	TotalReplays int64                                   `json:"totalReplays" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Timeseries   respjson.Field
		TotalReplays respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionReplayOverviewResponse) RawJSON() string { return r.JSON.raw }
func (r *SessionReplayOverviewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionReplayOverviewResponseTimesery struct {
	Date  string `json:"date" api:"required"`
	Value int64  `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SessionReplayOverviewResponseTimesery) RawJSON() string { return r.JSON.raw }
func (r *SessionReplayOverviewResponseTimesery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SessionReplayListParams struct {
	// Inclusive lower bound of the replay window as `YYYY-MM-DD`.
	From time.Time `query:"from" api:"required" format:"date" json:"-"`
	// Inclusive upper bound of the replay window as `YYYY-MM-DD`.
	To time.Time `query:"to" api:"required" format:"date" json:"-"`
	// Opaque pagination cursor from pagination.nextCursor in the previous response. Do
	// not decode or modify it. Malformed cursors return 400 Bad Request.
	Cursor    param.Opt[string] `query:"cursor,omitzero" json:"-"`
	EventName param.Opt[string] `query:"eventName,omitzero" json:"-"`
	// Maximum replay sessions to return. Defaults to the report default.
	Limit    param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Pathname param.Opt[string] `query:"pathname,omitzero" json:"-"`
	// Optional JSON-encoded session ID array. Maximum 100 session IDs.
	SessionIDs  param.Opt[string] `query:"sessionIds,omitzero" json:"-"`
	UtmCampaign param.Opt[string] `query:"utmCampaign,omitzero" json:"-"`
	UtmContent  param.Opt[string] `query:"utmContent,omitzero" json:"-"`
	UtmMedium   param.Opt[string] `query:"utmMedium,omitzero" json:"-"`
	UtmName     param.Opt[string] `query:"utmName,omitzero" json:"-"`
	UtmSource   param.Opt[string] `query:"utmSource,omitzero" json:"-"`
	UtmTerm     param.Opt[string] `query:"utmTerm,omitzero" json:"-"`
	VisitorID   param.Opt[string] `query:"visitorId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SessionReplayListParams]'s query parameters as
// `url.Values`.
func (r SessionReplayListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SessionReplayOverviewParams struct {
	// Inclusive lower bound of the replay window as `YYYY-MM-DD`.
	From time.Time `query:"from" api:"required" format:"date" json:"-"`
	// Inclusive upper bound of the replay window as `YYYY-MM-DD`.
	To time.Time `query:"to" api:"required" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [SessionReplayOverviewParams]'s query parameters as
// `url.Values`.
func (r SessionReplayOverviewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
