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
	"github.com/with-ours/platform-sdk-go/v2/packages/respjson"
)

// QueryVisitorService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQueryVisitorService] method instead.
type QueryVisitorService struct {
	Options []option.RequestOption
}

// NewQueryVisitorService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewQueryVisitorService(opts ...option.RequestOption) (r QueryVisitorService) {
	r = QueryVisitorService{}
	r.Options = opts
	return
}

// List visitors with matching events and their permitted current-profile fields.
// JSON-encode input in the query string; cursors stay bound to the account and
// filters, while page size may change. Requires API-key scope or current OAuth
// user permission: web-analytics:view, report:list-events
func (r *QueryVisitorService) List(ctx context.Context, query QueryVisitorListParams, opts ...option.RequestOption) (res *QueryVisitorListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/query-visitors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Return Query Visitors date, page, selection, and touch-filter capabilities.
// Requires API-key scope or current OAuth user permission: web-analytics:view,
// report:list-events
func (r *QueryVisitorService) Capabilities(ctx context.Context, opts ...option.RequestOption) (res *QueryVisitorCapabilitiesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/query-visitors/capabilities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Search current visitor profiles for the Query Visitors typeahead. Pass returned
// visitorIds to the list input to filter the result. Requires API-key scope or
// current OAuth user permission: web-analytics:view, report:list-events,
// report:view-visitor
func (r *QueryVisitorService) ProfileSearch(ctx context.Context, query QueryVisitorProfileSearchParams, opts ...option.RequestOption) (res *QueryVisitorProfileSearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/query-visitors/profile-search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type QueryVisitorListResponse struct {
	Freshness    QueryVisitorListResponseFreshness  `json:"freshness" api:"required"`
	MatchedCount float64                            `json:"matchedCount" api:"required"`
	Pagination   QueryVisitorListResponsePagination `json:"pagination" api:"required"`
	Rows         []QueryVisitorListResponseRow      `json:"rows" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Freshness    respjson.Field
		MatchedCount respjson.Field
		Pagination   respjson.Field
		Rows         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryVisitorListResponse) RawJSON() string { return r.JSON.raw }
func (r *QueryVisitorListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryVisitorListResponseFreshness struct {
	// Any of "UNKNOWN".
	CoverageStatus    string    `json:"coverageStatus" api:"required"`
	CoverageWatermark time.Time `json:"coverageWatermark" api:"required" format:"date-time"`
	SyncCompletedAt   time.Time `json:"syncCompletedAt" api:"nullable" format:"date-time"`
	SyncCursor        time.Time `json:"syncCursor" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CoverageStatus    respjson.Field
		CoverageWatermark respjson.Field
		SyncCompletedAt   respjson.Field
		SyncCursor        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryVisitorListResponseFreshness) RawJSON() string { return r.JSON.raw }
func (r *QueryVisitorListResponseFreshness) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryVisitorListResponsePagination struct {
	HasMore    bool   `json:"hasMore" api:"required"`
	NextCursor string `json:"nextCursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryVisitorListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *QueryVisitorListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryVisitorListResponseRow struct {
	Match     QueryVisitorListResponseRowMatch   `json:"match" api:"required"`
	Profile   QueryVisitorListResponseRowProfile `json:"profile" api:"required"`
	VisitorID string                             `json:"visitorId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Match       respjson.Field
		Profile     respjson.Field
		VisitorID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryVisitorListResponseRow) RawJSON() string { return r.JSON.raw }
func (r *QueryVisitorListResponseRow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryVisitorListResponseRowMatch struct {
	EventCount   float64   `json:"eventCount" api:"required"`
	FirstEventAt time.Time `json:"firstEventAt" api:"required" format:"date-time"`
	LastEventAt  time.Time `json:"lastEventAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventCount   respjson.Field
		FirstEventAt respjson.Field
		LastEventAt  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryVisitorListResponseRowMatch) RawJSON() string { return r.JSON.raw }
func (r *QueryVisitorListResponseRowMatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryVisitorListResponseRowProfile struct {
	// Current profile email, when available.
	Email string `json:"email" api:"required"`
	// Current external ID, when available.
	ExternalID string `json:"externalId" api:"required"`
	// Current first name, when permitted and available.
	FirstName string `json:"firstName" api:"required"`
	// First activity recorded on the current profile.
	FirstSeenAt time.Time `json:"firstSeenAt" api:"required" format:"date-time"`
	// Current last name, when permitted and available.
	LastName string `json:"lastName" api:"required"`
	// Most recent activity recorded on the current profile.
	LastSeenAt time.Time `json:"lastSeenAt" api:"required" format:"date-time"`
	// Current profile lookup status.
	//
	// Any of "FOUND", "NOT_FOUND", "UNAVAILABLE".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		ExternalID  respjson.Field
		FirstName   respjson.Field
		FirstSeenAt respjson.Field
		LastName    respjson.Field
		LastSeenAt  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryVisitorListResponseRowProfile) RawJSON() string { return r.JSON.raw }
func (r *QueryVisitorListResponseRowProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryVisitorCapabilitiesResponse struct {
	// Any of 25.
	DefaultPageSize          float64 `json:"defaultPageSize" api:"required"`
	DefaultTouchLookbackDays int64   `json:"defaultTouchLookbackDays" api:"required"`
	// Any of 31.
	MaxDateRangeDays float64 `json:"maxDateRangeDays" api:"required"`
	// Any of 100.
	MaxPageSize float64 `json:"maxPageSize" api:"required"`
	// Any of 10.
	MaxSelections                     float64 `json:"maxSelections" api:"required"`
	SupportsVerifiedCoverageWatermark bool    `json:"supportsVerifiedCoverageWatermark" api:"required"`
	TouchEnabled                      bool    `json:"touchEnabled" api:"required"`
	TouchLookbackDays                 []int64 `json:"touchLookbackDays" api:"required"`
	// Any of 1.
	Version float64 `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DefaultPageSize                   respjson.Field
		DefaultTouchLookbackDays          respjson.Field
		MaxDateRangeDays                  respjson.Field
		MaxPageSize                       respjson.Field
		MaxSelections                     respjson.Field
		SupportsVerifiedCoverageWatermark respjson.Field
		TouchEnabled                      respjson.Field
		TouchLookbackDays                 respjson.Field
		Version                           respjson.Field
		ExtraFields                       map[string]respjson.Field
		raw                               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryVisitorCapabilitiesResponse) RawJSON() string { return r.JSON.raw }
func (r *QueryVisitorCapabilitiesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryVisitorProfileSearchResponse struct {
	Matches    []QueryVisitorProfileSearchResponseMatch `json:"matches" api:"required"`
	Truncated  bool                                     `json:"truncated" api:"required"`
	VisitorIDs []string                                 `json:"visitorIds" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Matches     respjson.Field
		Truncated   respjson.Field
		VisitorIDs  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryVisitorProfileSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *QueryVisitorProfileSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryVisitorProfileSearchResponseMatch struct {
	Email      string    `json:"email" api:"required"`
	ExternalID string    `json:"externalId" api:"required"`
	FirstName  string    `json:"firstName" api:"required"`
	LastName   string    `json:"lastName" api:"required"`
	LastSeenAt time.Time `json:"lastSeenAt" api:"required" format:"date-time"`
	VisitorID  string    `json:"visitorId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		ExternalID  respjson.Field
		FirstName   respjson.Field
		LastName    respjson.Field
		LastSeenAt  respjson.Field
		VisitorID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryVisitorProfileSearchResponseMatch) RawJSON() string { return r.JSON.raw }
func (r *QueryVisitorProfileSearchResponseMatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryVisitorListParams struct {
	// Query Visitors filters and page. JSON-encode input in the REST query string.
	Input string `query:"input" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [QueryVisitorListParams]'s query parameters as `url.Values`.
func (r QueryVisitorListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type QueryVisitorProfileSearchParams struct {
	// Current profile search filters. JSON-encode input in the REST query string.
	Input string `query:"input" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [QueryVisitorProfileSearchParams]'s query parameters as
// `url.Values`.
func (r QueryVisitorProfileSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
