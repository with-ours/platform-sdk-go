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
	"github.com/with-ours/platform-sdk-go/packages/param"
	"github.com/with-ours/platform-sdk-go/packages/respjson"
)

// AttributionService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAttributionService] method instead.
type AttributionService struct {
	Options []option.RequestOption
}

// NewAttributionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAttributionService(opts ...option.RequestOption) (r AttributionService) {
	r = AttributionService{}
	r.Options = opts
	return
}

// Returns the top-15 values for each UTM dimension (source, medium, campaign,
// content, term, name) and referring domain attributed to the conversion event on
// a first-touch basis for the given date window. Use `from`/`to` to set the
// analysis window (max 60 days). Optionally filter to a specific UTM combo with
// `utmSource`, `utmMedium`, etc. The counts represent unique visitors who
// performed the specified `eventName` and were attributed to each UTM value.
// Requires scope: web-analytics:view
func (r *AttributionService) Initial(ctx context.Context, query AttributionInitialParams, opts ...option.RequestOption) (res *AttributionInitialResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/attribution/initial"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the top-15 values for each UTM dimension (source, medium, campaign,
// content, term, name) and referring domain attributed to the conversion event on
// a last-touch basis for the given date window. Use `from`/`to` to set the
// analysis window (max 60 days). The counts represent unique visitors who
// performed the specified `eventName` and were attributed to each UTM value on
// their most recent session. Requires scope: web-analytics:view
func (r *AttributionService) LastTouch(ctx context.Context, query AttributionLastTouchParams, opts ...option.RequestOption) (res *AttributionLastTouchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/attribution/last-touch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Multi-touch conversion attribution: returns a source → medium → campaign
// hierarchy with attributed converter credits distributed according to the
// selected `attributionModel`. Scoped to all web sources by default; optionally
// narrow to a single web source via `webSourceId`. Date range is capped at 31
// days; lookback window is capped at 60 days. Requires scope: web-analytics:view
func (r *AttributionService) Conversion(ctx context.Context, query AttributionConversionParams, opts ...option.RequestOption) (res *AttributionConversionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/attribution/conversion"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Audience performance conversion report: returns a summary of converters and
// conversion rate for the selected event and date window, a per-day timeseries,
// and a UTM source/medium/campaign breakdown. Optionally compare against the
// preceding period of equal length when `attributionWindow` is `IN_RANGE`. Date
// range is capped at 60 days. Requires scope: web-analytics:view
func (r *AttributionService) AudienceConversion(ctx context.Context, query AttributionAudienceConversionParams, opts ...option.RequestOption) (res *AttributionAudienceConversionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/attribution/audience-conversion"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Compare up to 5 UTM dimension combinations side-by-side for a single conversion
// event. Each combo returns the unique visitors, sessions, total events, and
// derived conversion rate for that UTM filter within the window. Requires both
// `web-analytics:view` and `report:event-count-by-day` API-key scopes. Date range
// is capped at 31 days. Pass `combos` as a single JSON-encoded array:
// `combos=[{"utmSource":"google","utmMedium":"cpc"},{"utmSource":"meta"}]`.
// Requires scope: web-analytics:view
func (r *AttributionService) UtmComparison(ctx context.Context, query AttributionUtmComparisonParams, opts ...option.RequestOption) (res *AttributionUtmComparisonResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/attribution/utm-comparison"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AttributionInitialResponse struct {
	InitialReferringDomain []AttributionInitialResponseInitialReferringDomain `json:"initial_referring_domain" api:"required"`
	InitialUtmCampaign     []AttributionInitialResponseInitialUtmCampaign     `json:"initial_utm_campaign" api:"required"`
	InitialUtmContent      []AttributionInitialResponseInitialUtmContent      `json:"initial_utm_content" api:"required"`
	InitialUtmMedium       []AttributionInitialResponseInitialUtmMedium       `json:"initial_utm_medium" api:"required"`
	InitialUtmName         []AttributionInitialResponseInitialUtmName         `json:"initial_utm_name" api:"required"`
	InitialUtmSource       []AttributionInitialResponseInitialUtmSource       `json:"initial_utm_source" api:"required"`
	InitialUtmTerm         []AttributionInitialResponseInitialUtmTerm         `json:"initial_utm_term" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InitialReferringDomain respjson.Field
		InitialUtmCampaign     respjson.Field
		InitialUtmContent      respjson.Field
		InitialUtmMedium       respjson.Field
		InitialUtmName         respjson.Field
		InitialUtmSource       respjson.Field
		InitialUtmTerm         respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionInitialResponse) RawJSON() string { return r.JSON.raw }
func (r *AttributionInitialResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionInitialResponseInitialReferringDomain struct {
	Count int64  `json:"count" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionInitialResponseInitialReferringDomain) RawJSON() string { return r.JSON.raw }
func (r *AttributionInitialResponseInitialReferringDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionInitialResponseInitialUtmCampaign struct {
	Count int64  `json:"count" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionInitialResponseInitialUtmCampaign) RawJSON() string { return r.JSON.raw }
func (r *AttributionInitialResponseInitialUtmCampaign) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionInitialResponseInitialUtmContent struct {
	Count int64  `json:"count" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionInitialResponseInitialUtmContent) RawJSON() string { return r.JSON.raw }
func (r *AttributionInitialResponseInitialUtmContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionInitialResponseInitialUtmMedium struct {
	Count int64  `json:"count" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionInitialResponseInitialUtmMedium) RawJSON() string { return r.JSON.raw }
func (r *AttributionInitialResponseInitialUtmMedium) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionInitialResponseInitialUtmName struct {
	Count int64  `json:"count" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionInitialResponseInitialUtmName) RawJSON() string { return r.JSON.raw }
func (r *AttributionInitialResponseInitialUtmName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionInitialResponseInitialUtmSource struct {
	Count int64  `json:"count" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionInitialResponseInitialUtmSource) RawJSON() string { return r.JSON.raw }
func (r *AttributionInitialResponseInitialUtmSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionInitialResponseInitialUtmTerm struct {
	Count int64  `json:"count" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionInitialResponseInitialUtmTerm) RawJSON() string { return r.JSON.raw }
func (r *AttributionInitialResponseInitialUtmTerm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionLastTouchResponse struct {
	ReferringDomain []AttributionLastTouchResponseReferringDomain `json:"referring_domain" api:"required"`
	UtmCampaign     []AttributionLastTouchResponseUtmCampaign     `json:"utm_campaign" api:"required"`
	UtmContent      []AttributionLastTouchResponseUtmContent      `json:"utm_content" api:"required"`
	UtmMedium       []AttributionLastTouchResponseUtmMedium       `json:"utm_medium" api:"required"`
	UtmName         []AttributionLastTouchResponseUtmName         `json:"utm_name" api:"required"`
	UtmSource       []AttributionLastTouchResponseUtmSource       `json:"utm_source" api:"required"`
	UtmTerm         []AttributionLastTouchResponseUtmTerm         `json:"utm_term" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReferringDomain respjson.Field
		UtmCampaign     respjson.Field
		UtmContent      respjson.Field
		UtmMedium       respjson.Field
		UtmName         respjson.Field
		UtmSource       respjson.Field
		UtmTerm         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionLastTouchResponse) RawJSON() string { return r.JSON.raw }
func (r *AttributionLastTouchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionLastTouchResponseReferringDomain struct {
	Count int64  `json:"count" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionLastTouchResponseReferringDomain) RawJSON() string { return r.JSON.raw }
func (r *AttributionLastTouchResponseReferringDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionLastTouchResponseUtmCampaign struct {
	Count int64  `json:"count" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionLastTouchResponseUtmCampaign) RawJSON() string { return r.JSON.raw }
func (r *AttributionLastTouchResponseUtmCampaign) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionLastTouchResponseUtmContent struct {
	Count int64  `json:"count" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionLastTouchResponseUtmContent) RawJSON() string { return r.JSON.raw }
func (r *AttributionLastTouchResponseUtmContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionLastTouchResponseUtmMedium struct {
	Count int64  `json:"count" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionLastTouchResponseUtmMedium) RawJSON() string { return r.JSON.raw }
func (r *AttributionLastTouchResponseUtmMedium) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionLastTouchResponseUtmName struct {
	Count int64  `json:"count" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionLastTouchResponseUtmName) RawJSON() string { return r.JSON.raw }
func (r *AttributionLastTouchResponseUtmName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionLastTouchResponseUtmSource struct {
	Count int64  `json:"count" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionLastTouchResponseUtmSource) RawJSON() string { return r.JSON.raw }
func (r *AttributionLastTouchResponseUtmSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionLastTouchResponseUtmTerm struct {
	Count int64  `json:"count" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionLastTouchResponseUtmTerm) RawJSON() string { return r.JSON.raw }
func (r *AttributionLastTouchResponseUtmTerm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionConversionResponse struct {
	IsTruncated   bool                                 `json:"isTruncated" api:"required"`
	MaxLeafRows   int64                                `json:"maxLeafRows" api:"required"`
	Nodes         []AttributionConversionResponseNode  `json:"nodes" api:"required"`
	Summary       AttributionConversionResponseSummary `json:"summary" api:"required"`
	TotalLeafRows int64                                `json:"totalLeafRows" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsTruncated   respjson.Field
		MaxLeafRows   respjson.Field
		Nodes         respjson.Field
		Summary       respjson.Field
		TotalLeafRows respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionConversionResponse) RawJSON() string { return r.JSON.raw }
func (r *AttributionConversionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionConversionResponseNode struct {
	AttributedConverterCredit float64 `json:"attributedConverterCredit" api:"required"`
	Converters                int64   `json:"converters" api:"required"`
	// Any of "SOURCE", "MEDIUM", "CAMPAIGN".
	Level    string `json:"level" api:"required"`
	Sessions int64  `json:"sessions" api:"required"`
	Source   string `json:"source" api:"required"`
	Campaign string `json:"campaign" api:"nullable"`
	Medium   string `json:"medium" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttributedConverterCredit respjson.Field
		Converters                respjson.Field
		Level                     respjson.Field
		Sessions                  respjson.Field
		Source                    respjson.Field
		Campaign                  respjson.Field
		Medium                    respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionConversionResponseNode) RawJSON() string { return r.JSON.raw }
func (r *AttributionConversionResponseNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionConversionResponseSummary struct {
	AttributedConverters int64 `json:"attributedConverters" api:"required"`
	ScopeConverters      int64 `json:"scopeConverters" api:"required"`
	TotalEventConverters int64 `json:"totalEventConverters" api:"required"`
	TotalSessions        int64 `json:"totalSessions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttributedConverters respjson.Field
		ScopeConverters      respjson.Field
		TotalEventConverters respjson.Field
		TotalSessions        respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionConversionResponseSummary) RawJSON() string { return r.JSON.raw }
func (r *AttributionConversionResponseSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionAudienceConversionResponse struct {
	Breakdown       []AttributionAudienceConversionResponseBreakdown     `json:"breakdown" api:"required"`
	Summary         AttributionAudienceConversionResponseSummary         `json:"summary" api:"required"`
	Timeseries      []AttributionAudienceConversionResponseTimesery      `json:"timeseries" api:"required"`
	PreviousSummary AttributionAudienceConversionResponsePreviousSummary `json:"previousSummary" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Breakdown       respjson.Field
		Summary         respjson.Field
		Timeseries      respjson.Field
		PreviousSummary respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionAudienceConversionResponse) RawJSON() string { return r.JSON.raw }
func (r *AttributionAudienceConversionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionAudienceConversionResponseBreakdown struct {
	Campaign    string  `json:"campaign" api:"required"`
	Conversions int64   `json:"conversions" api:"required"`
	Converters  int64   `json:"converters" api:"required"`
	Medium      string  `json:"medium" api:"required"`
	Source      string  `json:"source" api:"required"`
	TotalValue  float64 `json:"totalValue" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Campaign    respjson.Field
		Conversions respjson.Field
		Converters  respjson.Field
		Medium      respjson.Field
		Source      respjson.Field
		TotalValue  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionAudienceConversionResponseBreakdown) RawJSON() string { return r.JSON.raw }
func (r *AttributionAudienceConversionResponseBreakdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionAudienceConversionResponseSummary struct {
	AudienceSize                 int64   `json:"audienceSize" api:"required"`
	AvgValuePerConversion        float64 `json:"avgValuePerConversion" api:"required"`
	AvgValuePerConvertingVisitor float64 `json:"avgValuePerConvertingVisitor" api:"required"`
	ConversionRate               float64 `json:"conversionRate" api:"required"`
	Conversions                  int64   `json:"conversions" api:"required"`
	Converters                   int64   `json:"converters" api:"required"`
	TotalValue                   float64 `json:"totalValue" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AudienceSize                 respjson.Field
		AvgValuePerConversion        respjson.Field
		AvgValuePerConvertingVisitor respjson.Field
		ConversionRate               respjson.Field
		Conversions                  respjson.Field
		Converters                   respjson.Field
		TotalValue                   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionAudienceConversionResponseSummary) RawJSON() string { return r.JSON.raw }
func (r *AttributionAudienceConversionResponseSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionAudienceConversionResponseTimesery struct {
	Conversions int64   `json:"conversions" api:"required"`
	Date        string  `json:"date" api:"required"`
	TotalValue  float64 `json:"totalValue" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Conversions respjson.Field
		Date        respjson.Field
		TotalValue  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionAudienceConversionResponseTimesery) RawJSON() string { return r.JSON.raw }
func (r *AttributionAudienceConversionResponseTimesery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionAudienceConversionResponsePreviousSummary struct {
	AudienceSize                 int64   `json:"audienceSize" api:"required"`
	AvgValuePerConversion        float64 `json:"avgValuePerConversion" api:"required"`
	AvgValuePerConvertingVisitor float64 `json:"avgValuePerConvertingVisitor" api:"required"`
	ConversionRate               float64 `json:"conversionRate" api:"required"`
	Conversions                  int64   `json:"conversions" api:"required"`
	Converters                   int64   `json:"converters" api:"required"`
	TotalValue                   float64 `json:"totalValue" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AudienceSize                 respjson.Field
		AvgValuePerConversion        respjson.Field
		AvgValuePerConvertingVisitor respjson.Field
		ConversionRate               respjson.Field
		Conversions                  respjson.Field
		Converters                   respjson.Field
		TotalValue                   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionAudienceConversionResponsePreviousSummary) RawJSON() string { return r.JSON.raw }
func (r *AttributionAudienceConversionResponsePreviousSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionUtmComparisonResponse struct {
	// Per-combo metrics in the same order as the input `combos` array.
	Combos []AttributionUtmComparisonResponseCombo `json:"combos" api:"required"`
	// The conversion event that was analyzed.
	EventName string `json:"eventName" api:"required"`
	// The requested date range as returned by the server.
	Range AttributionUtmComparisonResponseRange `json:"range" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Combos      respjson.Field
		EventName   respjson.Field
		Range       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionUtmComparisonResponse) RawJSON() string { return r.JSON.raw }
func (r *AttributionUtmComparisonResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionUtmComparisonResponseCombo struct {
	Combo          AttributionUtmComparisonResponseComboCombo `json:"combo" api:"required"`
	ConversionRate float64                                    `json:"conversionRate" api:"required"`
	Events         int64                                      `json:"events" api:"required"`
	Sessions       int64                                      `json:"sessions" api:"required"`
	Visitors       int64                                      `json:"visitors" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Combo          respjson.Field
		ConversionRate respjson.Field
		Events         respjson.Field
		Sessions       respjson.Field
		Visitors       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionUtmComparisonResponseCombo) RawJSON() string { return r.JSON.raw }
func (r *AttributionUtmComparisonResponseCombo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionUtmComparisonResponseComboCombo struct {
	UtmCampaign string `json:"utmCampaign" api:"nullable"`
	UtmContent  string `json:"utmContent" api:"nullable"`
	UtmMedium   string `json:"utmMedium" api:"nullable"`
	UtmSource   string `json:"utmSource" api:"nullable"`
	UtmTerm     string `json:"utmTerm" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		UtmCampaign respjson.Field
		UtmContent  respjson.Field
		UtmMedium   respjson.Field
		UtmSource   respjson.Field
		UtmTerm     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionUtmComparisonResponseComboCombo) RawJSON() string { return r.JSON.raw }
func (r *AttributionUtmComparisonResponseComboCombo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The requested date range as returned by the server.
type AttributionUtmComparisonResponseRange struct {
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
func (r AttributionUtmComparisonResponseRange) RawJSON() string { return r.JSON.raw }
func (r *AttributionUtmComparisonResponseRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionInitialParams struct {
	// Conversion event to count. Must be a selectable conversion event.
	EventName string `query:"eventName" api:"required" json:"-"`
	// Inclusive lower bound of the analytics window, as a UTC calendar day in
	// `YYYY-MM-DD` format. The window between `from` and `to` must be 60 days or
	// fewer.
	From string `query:"from" api:"required" json:"-"`
	// Inclusive upper bound of the analytics window, as a UTC calendar day in
	// `YYYY-MM-DD` format. The window between `from` and `to` must be 60 days or
	// fewer.
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
	// Attribution type for UTM filter matching. Defaults to `INITIAL`.
	//
	// Any of "INITIAL", "LAST_TOUCH".
	AttributionType AttributionInitialParamsAttributionType `query:"attributionType,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AttributionInitialParams]'s query parameters as
// `url.Values`.
func (r AttributionInitialParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Attribution type for UTM filter matching. Defaults to `INITIAL`.
type AttributionInitialParamsAttributionType string

const (
	AttributionInitialParamsAttributionTypeInitial   AttributionInitialParamsAttributionType = "INITIAL"
	AttributionInitialParamsAttributionTypeLastTouch AttributionInitialParamsAttributionType = "LAST_TOUCH"
)

type AttributionLastTouchParams struct {
	// Conversion event to count. Must be a selectable conversion event.
	EventName string `query:"eventName" api:"required" json:"-"`
	// Inclusive lower bound of the analytics window, as a UTC calendar day in
	// `YYYY-MM-DD` format. The window between `from` and `to` must be 60 days or
	// fewer.
	From string `query:"from" api:"required" json:"-"`
	// Inclusive upper bound of the analytics window, as a UTC calendar day in
	// `YYYY-MM-DD` format. The window between `from` and `to` must be 60 days or
	// fewer.
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
	// Attribution type for UTM filter matching. Defaults to `LAST_TOUCH`.
	//
	// Any of "INITIAL", "LAST_TOUCH".
	AttributionType AttributionLastTouchParamsAttributionType `query:"attributionType,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AttributionLastTouchParams]'s query parameters as
// `url.Values`.
func (r AttributionLastTouchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Attribution type for UTM filter matching. Defaults to `LAST_TOUCH`.
type AttributionLastTouchParamsAttributionType string

const (
	AttributionLastTouchParamsAttributionTypeInitial   AttributionLastTouchParamsAttributionType = "INITIAL"
	AttributionLastTouchParamsAttributionTypeLastTouch AttributionLastTouchParamsAttributionType = "LAST_TOUCH"
)

type AttributionConversionParams struct {
	// Attribution model to apply to multi-touch conversion paths.
	//
	// Any of "FIRST_TOUCH", "LAST_TOUCH", "LINEAR", "POSITION_BASED".
	AttributionModel AttributionConversionParamsAttributionModel `query:"attributionModel,omitzero" api:"required" json:"-"`
	// Conversion event to attribute. Must be a selectable conversion event.
	EventName string `query:"eventName" api:"required" json:"-"`
	// Inclusive lower bound of the analytics window, as a UTC calendar day in
	// `YYYY-MM-DD` format. The window between `from` and `to` must be 31 days or
	// fewer.
	From string `query:"from" api:"required" json:"-"`
	// Inclusive upper bound of the analytics window, as a UTC calendar day in
	// `YYYY-MM-DD` format. The window between `from` and `to` must be 31 days or
	// fewer.
	To string `query:"to" api:"required" json:"-"`
	// Maximum number of leaf-level attribution rows to return. Defaults to 1000.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Scope to a single web source by id, or omit for all sources (account-wide).
	WebSourceID param.Opt[string] `query:"webSourceId,omitzero" json:"-"`
	// How far back before each conversion to consider touchpoints. Capped at 60 days
	// for this report. Defaults to `THIRTY_DAYS`.
	//
	// Any of "SEVEN_DAYS", "FOURTEEN_DAYS", "THIRTY_DAYS", "SIXTY_DAYS",
	// "NINETY_DAYS", "ONE_HUNDRED_EIGHTY_DAYS", "UNLIMITED".
	LookbackWindow AttributionConversionParamsLookbackWindow `query:"lookbackWindow,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AttributionConversionParams]'s query parameters as
// `url.Values`.
func (r AttributionConversionParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Attribution model to apply to multi-touch conversion paths.
type AttributionConversionParamsAttributionModel string

const (
	AttributionConversionParamsAttributionModelFirstTouch    AttributionConversionParamsAttributionModel = "FIRST_TOUCH"
	AttributionConversionParamsAttributionModelLastTouch     AttributionConversionParamsAttributionModel = "LAST_TOUCH"
	AttributionConversionParamsAttributionModelLinear        AttributionConversionParamsAttributionModel = "LINEAR"
	AttributionConversionParamsAttributionModelPositionBased AttributionConversionParamsAttributionModel = "POSITION_BASED"
)

// How far back before each conversion to consider touchpoints. Capped at 60 days
// for this report. Defaults to `THIRTY_DAYS`.
type AttributionConversionParamsLookbackWindow string

const (
	AttributionConversionParamsLookbackWindowSevenDays            AttributionConversionParamsLookbackWindow = "SEVEN_DAYS"
	AttributionConversionParamsLookbackWindowFourteenDays         AttributionConversionParamsLookbackWindow = "FOURTEEN_DAYS"
	AttributionConversionParamsLookbackWindowThirtyDays           AttributionConversionParamsLookbackWindow = "THIRTY_DAYS"
	AttributionConversionParamsLookbackWindowSixtyDays            AttributionConversionParamsLookbackWindow = "SIXTY_DAYS"
	AttributionConversionParamsLookbackWindowNinetyDays           AttributionConversionParamsLookbackWindow = "NINETY_DAYS"
	AttributionConversionParamsLookbackWindowOneHundredEightyDays AttributionConversionParamsLookbackWindow = "ONE_HUNDRED_EIGHTY_DAYS"
	AttributionConversionParamsLookbackWindowUnlimited            AttributionConversionParamsLookbackWindow = "UNLIMITED"
)

type AttributionAudienceConversionParams struct {
	// Conversion event to analyze.
	EventName string `query:"eventName" api:"required" json:"-"`
	// Inclusive lower bound of the analytics window, as a UTC calendar day in
	// `YYYY-MM-DD` format. The window between `from` and `to` must be 60 days or
	// fewer.
	From string `query:"from" api:"required" json:"-"`
	// Inclusive upper bound of the analytics window, as a UTC calendar day in
	// `YYYY-MM-DD` format. The window between `from` and `to` must be 60 days or
	// fewer.
	To string `query:"to" api:"required" json:"-"`
	// Attribution window: `IN_RANGE` or a number of lookback days (e.g. `7`, `30`).
	// Defaults to `IN_RANGE`.
	AttributionWindow param.Opt[string] `query:"attributionWindow,omitzero" json:"-"`
	// Event property to sum as conversion value.
	ValueProperty param.Opt[string] `query:"valueProperty,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AttributionAudienceConversionParams]'s query parameters as
// `url.Values`.
func (r AttributionAudienceConversionParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AttributionUtmComparisonParams struct {
	// JSON-encoded array of UTM dimension combos to compare side-by-side (min 1, max
	// 5). Each combo is an object with optional `utmSource`, `utmMedium`,
	// `utmCampaign`, `utmContent`, `utmTerm` fields.
	Combos string `query:"combos" api:"required" json:"-"`
	// Conversion event to compare across UTM combos.
	EventName string `query:"eventName" api:"required" json:"-"`
	// Inclusive lower bound of the analytics window, as a UTC calendar day in
	// `YYYY-MM-DD` format. The window between `from` and `to` must be 31 days or
	// fewer.
	From string `query:"from" api:"required" json:"-"`
	// Inclusive upper bound of the analytics window, as a UTC calendar day in
	// `YYYY-MM-DD` format. The window between `from` and `to` must be 31 days or
	// fewer.
	To string `query:"to" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [AttributionUtmComparisonParams]'s query parameters as
// `url.Values`.
func (r AttributionUtmComparisonParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
