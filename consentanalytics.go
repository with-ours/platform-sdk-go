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
	"github.com/with-ours/platform-sdk-go/packages/respjson"
)

// ConsentAnalyticsService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConsentAnalyticsService] method instead.
type ConsentAnalyticsService struct {
	Options []option.RequestOption
}

// NewConsentAnalyticsService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConsentAnalyticsService(opts ...option.RequestOption) (r ConsentAnalyticsService) {
	r = ConsentAnalyticsService{}
	r.Options = opts
	return
}

// Account-wide blocking stats from the Global Consent Center for the window: how
// many dispatches were attempted, how many were blocked, and a per-category
// breakdown of the blocks (with `percentageBlocked` = share of `totalDispatches`).
// Not scoped to a single consent settings record — this aggregates across every
// destination in the account. The endpoint is identified by query params rather
// than a path id because the report is account-scoped; this is a documented
// derived-read exception. Requires the API-key scope
// `report:global-consent-center-analytics` (this is the account-wide consent
// analytics report and is gated separately from consent-settings list). Requires
// scope: report:global-consent-center-analytics
func (r *ConsentAnalyticsService) List(ctx context.Context, query ConsentAnalyticsListParams, opts ...option.RequestOption) (res *ConsentAnalyticsListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/consent-analytics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ConsentAnalyticsListResponse struct {
	// Per-category breakdown of blocked dispatches. `categoryName` is `Unknown` when
	// the upstream block message could not be parsed; `percentageBlocked` is the share
	// of `totalDispatches` blocked under that category.
	Items []ConsentAnalyticsListResponseItem `json:"items" api:"required"`
	// Total dispatches blocked by the Global Consent Center across all categories.
	TotalBlocked int64 `json:"totalBlocked" api:"required"`
	// Total dispatches attempted across all destinations for the window.
	TotalDispatches int64 `json:"totalDispatches" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items           respjson.Field
		TotalBlocked    respjson.Field
		TotalDispatches respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentAnalyticsListResponse) RawJSON() string { return r.JSON.raw }
func (r *ConsentAnalyticsListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentAnalyticsListResponseItem struct {
	BlockedCount      int64   `json:"blockedCount" api:"required"`
	CategoryName      string  `json:"categoryName" api:"required"`
	PercentageBlocked float64 `json:"percentageBlocked" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BlockedCount      respjson.Field
		CategoryName      respjson.Field
		PercentageBlocked respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentAnalyticsListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *ConsentAnalyticsListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentAnalyticsListParams struct {
	// Inclusive lower bound of the analytics window, as a UTC calendar day in
	// `YYYY-MM-DD` format. The window between `from` and `to` must be 60 days or
	// fewer.
	From string `query:"from" api:"required" json:"-"`
	// Inclusive upper bound of the analytics window, as a UTC calendar day in
	// `YYYY-MM-DD` format. The window between `from` and `to` must be 60 days or
	// fewer.
	To string `query:"to" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [ConsentAnalyticsListParams]'s query parameters as
// `url.Values`.
func (r ConsentAnalyticsListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
