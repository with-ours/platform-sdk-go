// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/with-ours/platform-sdk-go/internal/apijson"
	"github.com/with-ours/platform-sdk-go/internal/apiquery"
	"github.com/with-ours/platform-sdk-go/internal/requestconfig"
	"github.com/with-ours/platform-sdk-go/option"
	"github.com/with-ours/platform-sdk-go/packages/param"
	"github.com/with-ours/platform-sdk-go/packages/respjson"
)

// FunnelService contains methods and other services that help with interacting
// with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFunnelService] method instead.
type FunnelService struct {
	Options []option.RequestOption
}

// NewFunnelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFunnelService(opts ...option.RequestOption) (r FunnelService) {
	r = FunnelService{}
	r.Options = opts
	return
}

// List every funnel configured on this account. Each funnel includes its step
// configuration, funnel type, conversion window, and current processing status.
// The available report date range (if any pre-computed reports exist) is returned
// in `reportDateRange`. Requires scope: web-analytics:view
func (r *FunnelService) List(ctx context.Context, opts ...option.RequestOption) (res *FunnelListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/funnels"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Fetch a single funnel configuration by its id. Returns `404` when the funnel
// does not exist or belongs to a different account. Requires scope:
// web-analytics:view
func (r *FunnelService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *FunnelGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/funnels/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Compute funnel step analytics for a funnel over a date window. Returns per-step
// visitor counts, conversion rates, drop-off rates, average time to next step, and
// sample session IDs for replay. Funnel results are pre-computed daily from S3;
// reports outside the `reportDateRange` shown on the funnel config will return
// empty steps. Requires scope: web-analytics:view
func (r *FunnelService) Results(ctx context.Context, id string, query FunnelResultsParams, opts ...option.RequestOption) (res *FunnelResultsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/funnels/%s/results", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type FunnelListResponse struct {
	// All funnels configured on this account.
	Entities []FunnelListResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelListResponse) RawJSON() string { return r.JSON.raw }
func (r *FunnelListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelListResponseEntity struct {
	CreatedAt string `json:"createdAt" api:"required"`
	FunnelID  string `json:"funnelId" api:"required" format:"uuid"`
	// Any of "SESSION_BASED", "VISITOR_BASED".
	FunnelType string `json:"funnelType" api:"required"`
	Name       string `json:"name" api:"required"`
	// Any of "READY", "PROCESSING".
	Status           string                                   `json:"status" api:"required"`
	Steps            []FunnelListResponseEntityStep           `json:"steps" api:"required"`
	UpdatedAt        string                                   `json:"updatedAt" api:"required"`
	ConversionWindow FunnelListResponseEntityConversionWindow `json:"conversionWindow" api:"nullable"`
	// Any of "UNIQUES", "TOTALS", "SESSIONS".
	CountingMethod  string                                  `json:"countingMethod" api:"nullable"`
	Description     string                                  `json:"description" api:"nullable"`
	ReportDateRange FunnelListResponseEntityReportDateRange `json:"reportDateRange" api:"nullable"`
	// Any of "EXACT", "ANY".
	StepOrder string `json:"stepOrder" api:"nullable"`
	// UTM filter object (JSON).
	UtmFilters any  `json:"utmFilters"`
	Watched    bool `json:"watched" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt        respjson.Field
		FunnelID         respjson.Field
		FunnelType       respjson.Field
		Name             respjson.Field
		Status           respjson.Field
		Steps            respjson.Field
		UpdatedAt        respjson.Field
		ConversionWindow respjson.Field
		CountingMethod   respjson.Field
		Description      respjson.Field
		ReportDateRange  respjson.Field
		StepOrder        respjson.Field
		UtmFilters       respjson.Field
		Watched          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelListResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *FunnelListResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelListResponseEntityStep struct {
	EventName string `json:"eventName" api:"required"`
	Name      string `json:"name" api:"required"`
	Order     int64  `json:"order" api:"required"`
	StepID    string `json:"stepId" api:"required"`
	// Step-level event filters (JSON object).
	Filters any `json:"filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventName   respjson.Field
		Name        respjson.Field
		Order       respjson.Field
		StepID      respjson.Field
		Filters     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelListResponseEntityStep) RawJSON() string { return r.JSON.raw }
func (r *FunnelListResponseEntityStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelListResponseEntityConversionWindow struct {
	// Any of "MINUTES", "HOURS", "DAYS".
	Unit  string `json:"unit" api:"required"`
	Value int64  `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Unit        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelListResponseEntityConversionWindow) RawJSON() string { return r.JSON.raw }
func (r *FunnelListResponseEntityConversionWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelListResponseEntityReportDateRange struct {
	From string `json:"from" api:"required"`
	To   string `json:"to" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelListResponseEntityReportDateRange) RawJSON() string { return r.JSON.raw }
func (r *FunnelListResponseEntityReportDateRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Funnel configuration details.
type FunnelGetResponse struct {
	CreatedAt string `json:"createdAt" api:"required"`
	FunnelID  string `json:"funnelId" api:"required" format:"uuid"`
	// Any of "SESSION_BASED", "VISITOR_BASED".
	FunnelType FunnelGetResponseFunnelType `json:"funnelType" api:"required"`
	Name       string                      `json:"name" api:"required"`
	// Any of "READY", "PROCESSING".
	Status           FunnelGetResponseStatus           `json:"status" api:"required"`
	Steps            []FunnelGetResponseStep           `json:"steps" api:"required"`
	UpdatedAt        string                            `json:"updatedAt" api:"required"`
	ConversionWindow FunnelGetResponseConversionWindow `json:"conversionWindow" api:"nullable"`
	// Any of "UNIQUES", "TOTALS", "SESSIONS".
	CountingMethod  FunnelGetResponseCountingMethod  `json:"countingMethod" api:"nullable"`
	Description     string                           `json:"description" api:"nullable"`
	ReportDateRange FunnelGetResponseReportDateRange `json:"reportDateRange" api:"nullable"`
	// Any of "EXACT", "ANY".
	StepOrder FunnelGetResponseStepOrder `json:"stepOrder" api:"nullable"`
	// UTM filter object (JSON).
	UtmFilters any  `json:"utmFilters"`
	Watched    bool `json:"watched" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt        respjson.Field
		FunnelID         respjson.Field
		FunnelType       respjson.Field
		Name             respjson.Field
		Status           respjson.Field
		Steps            respjson.Field
		UpdatedAt        respjson.Field
		ConversionWindow respjson.Field
		CountingMethod   respjson.Field
		Description      respjson.Field
		ReportDateRange  respjson.Field
		StepOrder        respjson.Field
		UtmFilters       respjson.Field
		Watched          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelGetResponse) RawJSON() string { return r.JSON.raw }
func (r *FunnelGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelGetResponseFunnelType string

const (
	FunnelGetResponseFunnelTypeSessionBased FunnelGetResponseFunnelType = "SESSION_BASED"
	FunnelGetResponseFunnelTypeVisitorBased FunnelGetResponseFunnelType = "VISITOR_BASED"
)

type FunnelGetResponseStatus string

const (
	FunnelGetResponseStatusReady      FunnelGetResponseStatus = "READY"
	FunnelGetResponseStatusProcessing FunnelGetResponseStatus = "PROCESSING"
)

type FunnelGetResponseStep struct {
	EventName string `json:"eventName" api:"required"`
	Name      string `json:"name" api:"required"`
	Order     int64  `json:"order" api:"required"`
	StepID    string `json:"stepId" api:"required"`
	// Step-level event filters (JSON object).
	Filters any `json:"filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventName   respjson.Field
		Name        respjson.Field
		Order       respjson.Field
		StepID      respjson.Field
		Filters     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelGetResponseStep) RawJSON() string { return r.JSON.raw }
func (r *FunnelGetResponseStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelGetResponseConversionWindow struct {
	// Any of "MINUTES", "HOURS", "DAYS".
	Unit  string `json:"unit" api:"required"`
	Value int64  `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Unit        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelGetResponseConversionWindow) RawJSON() string { return r.JSON.raw }
func (r *FunnelGetResponseConversionWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelGetResponseCountingMethod string

const (
	FunnelGetResponseCountingMethodUniques  FunnelGetResponseCountingMethod = "UNIQUES"
	FunnelGetResponseCountingMethodTotals   FunnelGetResponseCountingMethod = "TOTALS"
	FunnelGetResponseCountingMethodSessions FunnelGetResponseCountingMethod = "SESSIONS"
)

type FunnelGetResponseReportDateRange struct {
	From string `json:"from" api:"required"`
	To   string `json:"to" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelGetResponseReportDateRange) RawJSON() string { return r.JSON.raw }
func (r *FunnelGetResponseReportDateRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelGetResponseStepOrder string

const (
	FunnelGetResponseStepOrderExact FunnelGetResponseStepOrder = "EXACT"
	FunnelGetResponseStepOrderAny   FunnelGetResponseStepOrder = "ANY"
)

type FunnelResultsResponse struct {
	// Conversion rate from first step to last step as a percentage.
	OverallConversionRate float64 `json:"overallConversionRate" api:"required"`
	// Per-step funnel analytics, ordered by step number.
	Steps []FunnelResultsResponseStep `json:"steps" api:"required"`
	// Total number of visitors who entered the funnel (entered step 1).
	TotalVisitors int64 `json:"totalVisitors" api:"required"`
	// Average time from first step to last step in seconds. Null when no completions.
	OverallAvgTimeToConversion float64 `json:"overallAvgTimeToConversion" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OverallConversionRate      respjson.Field
		Steps                      respjson.Field
		TotalVisitors              respjson.Field
		OverallAvgTimeToConversion respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelResultsResponse) RawJSON() string { return r.JSON.raw }
func (r *FunnelResultsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelResultsResponseStep struct {
	ConversionCount       int64    `json:"conversionCount" api:"required"`
	ConversionRate        float64  `json:"conversionRate" api:"required"`
	DropOffRate           float64  `json:"dropOffRate" api:"required"`
	DropOffSessionIDs     []string `json:"dropOffSessionIds" api:"required"`
	OverallConversionRate float64  `json:"overallConversionRate" api:"required"`
	SessionIDs            []string `json:"sessionIds" api:"required"`
	StepNumber            int64    `json:"stepNumber" api:"required"`
	VisitorCount          int64    `json:"visitorCount" api:"required"`
	AvgTimeToNextStep     float64  `json:"avgTimeToNextStep" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConversionCount       respjson.Field
		ConversionRate        respjson.Field
		DropOffRate           respjson.Field
		DropOffSessionIDs     respjson.Field
		OverallConversionRate respjson.Field
		SessionIDs            respjson.Field
		StepNumber            respjson.Field
		VisitorCount          respjson.Field
		AvgTimeToNextStep     respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelResultsResponseStep) RawJSON() string { return r.JSON.raw }
func (r *FunnelResultsResponseStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelResultsParams struct {
	// Inclusive lower bound of the analysis window, as a UTC calendar day in
	// `YYYY-MM-DD` format.
	From string `query:"from" api:"required" json:"-"`
	// Inclusive upper bound of the analysis window, as a UTC calendar day in
	// `YYYY-MM-DD` format.
	To string `query:"to" api:"required" json:"-"`
	// Filter by UTM campaign.
	UtmCampaign param.Opt[string] `query:"utmCampaign,omitzero" json:"-"`
	// Filter by UTM content.
	UtmContent param.Opt[string] `query:"utmContent,omitzero" json:"-"`
	// Filter by UTM medium.
	UtmMedium param.Opt[string] `query:"utmMedium,omitzero" json:"-"`
	// Filter by UTM name.
	UtmName param.Opt[string] `query:"utmName,omitzero" json:"-"`
	// Filter by UTM source.
	UtmSource param.Opt[string] `query:"utmSource,omitzero" json:"-"`
	// Filter by UTM term.
	UtmTerm param.Opt[string] `query:"utmTerm,omitzero" json:"-"`
	// Attribution type for UTM filter matching in funnel steps.
	//
	// Any of "INITIAL", "LAST_TOUCH".
	AttributionType FunnelResultsParamsAttributionType `query:"attributionType,omitzero" json:"-"`
	// Filter funnel analytics to a specific device type. Defaults to `ALL`.
	//
	// Any of "DESKTOP", "MOBILE", "ALL".
	DeviceType FunnelResultsParamsDeviceType `query:"deviceType,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FunnelResultsParams]'s query parameters as `url.Values`.
func (r FunnelResultsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Attribution type for UTM filter matching in funnel steps.
type FunnelResultsParamsAttributionType string

const (
	FunnelResultsParamsAttributionTypeInitial   FunnelResultsParamsAttributionType = "INITIAL"
	FunnelResultsParamsAttributionTypeLastTouch FunnelResultsParamsAttributionType = "LAST_TOUCH"
)

// Filter funnel analytics to a specific device type. Defaults to `ALL`.
type FunnelResultsParamsDeviceType string

const (
	FunnelResultsParamsDeviceTypeDesktop FunnelResultsParamsDeviceType = "DESKTOP"
	FunnelResultsParamsDeviceTypeMobile  FunnelResultsParamsDeviceType = "MOBILE"
	FunnelResultsParamsDeviceTypeAll     FunnelResultsParamsDeviceType = "ALL"
)
