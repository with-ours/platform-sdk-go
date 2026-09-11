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

// AudienceConversionReportService contains methods and other services that help
// with interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAudienceConversionReportService] method instead.
type AudienceConversionReportService struct {
	Options []option.RequestOption
}

// NewAudienceConversionReportService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAudienceConversionReportService(opts ...option.RequestOption) (r AudienceConversionReportService) {
	r = AudienceConversionReportService{}
	r.Options = opts
	return
}

// List saved Audience Performance report configurations, most recently updated
// first. Requires scope: web-analytics:view
func (r *AudienceConversionReportService) List(ctx context.Context, opts ...option.RequestOption) (res *AudienceConversionReportListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/audience-conversion-reports"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Save an Audience Performance report configuration. Returns the full report so
// callers can run or update it without another request. Requires scope:
// web-analytics:write
func (r *AudienceConversionReportService) New(ctx context.Context, body AudienceConversionReportNewParams, opts ...option.RequestOption) (res *AudienceConversionReportNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/audience-conversion-reports"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Fetch a saved Audience Performance report by id. Requires scope:
// web-analytics:view
func (r *AudienceConversionReportService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AudienceConversionReportGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/audience-conversion-reports/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a saved Audience Performance report. Omitted fields remain unchanged,
// `filters: []` clears all filters, and dates must be sent or cleared as a pair.
// Requires scope: web-analytics:write
func (r *AudienceConversionReportService) Update(ctx context.Context, id string, body AudienceConversionReportUpdateParams, opts ...option.RequestOption) (res *AudienceConversionReportUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/audience-conversion-reports/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete a saved Audience Performance report configuration. Requires scope:
// web-analytics:write
func (r *AudienceConversionReportService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *AudienceConversionReportDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/audience-conversion-reports/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Run a saved Audience Performance report. `from` and `to` override the saved date
// range when provided together. Returns 400 when neither a saved range nor an
// override is available. Requires scope: web-analytics:view
func (r *AudienceConversionReportService) Results(ctx context.Context, id string, query AudienceConversionReportResultsParams, opts ...option.RequestOption) (res *AudienceConversionReportResultsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/audience-conversion-reports/%s/results", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AudienceConversionReportListResponse struct {
	Entities []AudienceConversionReportListResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudienceConversionReportListResponse) RawJSON() string { return r.JSON.raw }
func (r *AudienceConversionReportListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudienceConversionReportListResponseEntity struct {
	AttributionWindow string                                             `json:"attributionWindow" api:"required"`
	CreatedAt         string                                             `json:"createdAt" api:"required"`
	EventName         string                                             `json:"eventName" api:"required"`
	Filters           []AudienceConversionReportListResponseEntityFilter `json:"filters" api:"required"`
	Name              string                                             `json:"name" api:"required"`
	ReportID          string                                             `json:"reportId" api:"required" format:"uuid"`
	UpdatedAt         string                                             `json:"updatedAt" api:"required"`
	ValueProperty     string                                             `json:"valueProperty" api:"required"`
	DateFrom          string                                             `json:"dateFrom" api:"nullable"`
	DateTo            string                                             `json:"dateTo" api:"nullable"`
	ExcludeBots       bool                                               `json:"excludeBots" api:"nullable"`
	WebSourceID       string                                             `json:"webSourceId" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttributionWindow respjson.Field
		CreatedAt         respjson.Field
		EventName         respjson.Field
		Filters           respjson.Field
		Name              respjson.Field
		ReportID          respjson.Field
		UpdatedAt         respjson.Field
		ValueProperty     respjson.Field
		DateFrom          respjson.Field
		DateTo            respjson.Field
		ExcludeBots       respjson.Field
		WebSourceID       respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudienceConversionReportListResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *AudienceConversionReportListResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudienceConversionReportListResponseEntityFilter struct {
	// Any of "browser", "campaign", "city", "content", "country", "device",
	// "entry_page", "exit_page", "medium", "os", "page", "referrer", "region",
	// "source", "term".
	Dimension string `json:"dimension" api:"required"`
	// Any of "CONTAINS", "IS", "IS_NOT", "NOT_CONTAINS".
	Operator string   `json:"operator" api:"required"`
	Values   []string `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dimension   respjson.Field
		Operator    respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudienceConversionReportListResponseEntityFilter) RawJSON() string { return r.JSON.raw }
func (r *AudienceConversionReportListResponseEntityFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudienceConversionReportNewResponse struct {
	AttributionWindow string                                      `json:"attributionWindow" api:"required"`
	CreatedAt         string                                      `json:"createdAt" api:"required"`
	EventName         string                                      `json:"eventName" api:"required"`
	Filters           []AudienceConversionReportNewResponseFilter `json:"filters" api:"required"`
	Name              string                                      `json:"name" api:"required"`
	ReportID          string                                      `json:"reportId" api:"required" format:"uuid"`
	UpdatedAt         string                                      `json:"updatedAt" api:"required"`
	ValueProperty     string                                      `json:"valueProperty" api:"required"`
	DateFrom          string                                      `json:"dateFrom" api:"nullable"`
	DateTo            string                                      `json:"dateTo" api:"nullable"`
	ExcludeBots       bool                                        `json:"excludeBots" api:"nullable"`
	WebSourceID       string                                      `json:"webSourceId" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttributionWindow respjson.Field
		CreatedAt         respjson.Field
		EventName         respjson.Field
		Filters           respjson.Field
		Name              respjson.Field
		ReportID          respjson.Field
		UpdatedAt         respjson.Field
		ValueProperty     respjson.Field
		DateFrom          respjson.Field
		DateTo            respjson.Field
		ExcludeBots       respjson.Field
		WebSourceID       respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudienceConversionReportNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AudienceConversionReportNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudienceConversionReportNewResponseFilter struct {
	// Any of "browser", "campaign", "city", "content", "country", "device",
	// "entry_page", "exit_page", "medium", "os", "page", "referrer", "region",
	// "source", "term".
	Dimension string `json:"dimension" api:"required"`
	// Any of "CONTAINS", "IS", "IS_NOT", "NOT_CONTAINS".
	Operator string   `json:"operator" api:"required"`
	Values   []string `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dimension   respjson.Field
		Operator    respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudienceConversionReportNewResponseFilter) RawJSON() string { return r.JSON.raw }
func (r *AudienceConversionReportNewResponseFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudienceConversionReportGetResponse struct {
	AttributionWindow string                                      `json:"attributionWindow" api:"required"`
	CreatedAt         string                                      `json:"createdAt" api:"required"`
	EventName         string                                      `json:"eventName" api:"required"`
	Filters           []AudienceConversionReportGetResponseFilter `json:"filters" api:"required"`
	Name              string                                      `json:"name" api:"required"`
	ReportID          string                                      `json:"reportId" api:"required" format:"uuid"`
	UpdatedAt         string                                      `json:"updatedAt" api:"required"`
	ValueProperty     string                                      `json:"valueProperty" api:"required"`
	DateFrom          string                                      `json:"dateFrom" api:"nullable"`
	DateTo            string                                      `json:"dateTo" api:"nullable"`
	ExcludeBots       bool                                        `json:"excludeBots" api:"nullable"`
	WebSourceID       string                                      `json:"webSourceId" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttributionWindow respjson.Field
		CreatedAt         respjson.Field
		EventName         respjson.Field
		Filters           respjson.Field
		Name              respjson.Field
		ReportID          respjson.Field
		UpdatedAt         respjson.Field
		ValueProperty     respjson.Field
		DateFrom          respjson.Field
		DateTo            respjson.Field
		ExcludeBots       respjson.Field
		WebSourceID       respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudienceConversionReportGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AudienceConversionReportGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudienceConversionReportGetResponseFilter struct {
	// Any of "browser", "campaign", "city", "content", "country", "device",
	// "entry_page", "exit_page", "medium", "os", "page", "referrer", "region",
	// "source", "term".
	Dimension string `json:"dimension" api:"required"`
	// Any of "CONTAINS", "IS", "IS_NOT", "NOT_CONTAINS".
	Operator string   `json:"operator" api:"required"`
	Values   []string `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dimension   respjson.Field
		Operator    respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudienceConversionReportGetResponseFilter) RawJSON() string { return r.JSON.raw }
func (r *AudienceConversionReportGetResponseFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudienceConversionReportUpdateResponse struct {
	AttributionWindow string                                         `json:"attributionWindow" api:"required"`
	CreatedAt         string                                         `json:"createdAt" api:"required"`
	EventName         string                                         `json:"eventName" api:"required"`
	Filters           []AudienceConversionReportUpdateResponseFilter `json:"filters" api:"required"`
	Name              string                                         `json:"name" api:"required"`
	ReportID          string                                         `json:"reportId" api:"required" format:"uuid"`
	UpdatedAt         string                                         `json:"updatedAt" api:"required"`
	ValueProperty     string                                         `json:"valueProperty" api:"required"`
	DateFrom          string                                         `json:"dateFrom" api:"nullable"`
	DateTo            string                                         `json:"dateTo" api:"nullable"`
	ExcludeBots       bool                                           `json:"excludeBots" api:"nullable"`
	WebSourceID       string                                         `json:"webSourceId" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttributionWindow respjson.Field
		CreatedAt         respjson.Field
		EventName         respjson.Field
		Filters           respjson.Field
		Name              respjson.Field
		ReportID          respjson.Field
		UpdatedAt         respjson.Field
		ValueProperty     respjson.Field
		DateFrom          respjson.Field
		DateTo            respjson.Field
		ExcludeBots       respjson.Field
		WebSourceID       respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudienceConversionReportUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *AudienceConversionReportUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudienceConversionReportUpdateResponseFilter struct {
	// Any of "browser", "campaign", "city", "content", "country", "device",
	// "entry_page", "exit_page", "medium", "os", "page", "referrer", "region",
	// "source", "term".
	Dimension string `json:"dimension" api:"required"`
	// Any of "CONTAINS", "IS", "IS_NOT", "NOT_CONTAINS".
	Operator string   `json:"operator" api:"required"`
	Values   []string `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dimension   respjson.Field
		Operator    respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudienceConversionReportUpdateResponseFilter) RawJSON() string { return r.JSON.raw }
func (r *AudienceConversionReportUpdateResponseFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudienceConversionReportDeleteResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Any of true.
	Deleted bool `json:"deleted" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Deleted     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudienceConversionReportDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *AudienceConversionReportDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudienceConversionReportResultsResponse struct {
	Breakdown       []AudienceConversionReportResultsResponseBreakdown     `json:"breakdown" api:"required"`
	Summary         AudienceConversionReportResultsResponseSummary         `json:"summary" api:"required"`
	Timeseries      []AudienceConversionReportResultsResponseTimesery      `json:"timeseries" api:"required"`
	PreviousSummary AudienceConversionReportResultsResponsePreviousSummary `json:"previousSummary" api:"nullable"`
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
func (r AudienceConversionReportResultsResponse) RawJSON() string { return r.JSON.raw }
func (r *AudienceConversionReportResultsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudienceConversionReportResultsResponseBreakdown struct {
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
func (r AudienceConversionReportResultsResponseBreakdown) RawJSON() string { return r.JSON.raw }
func (r *AudienceConversionReportResultsResponseBreakdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudienceConversionReportResultsResponseSummary struct {
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
func (r AudienceConversionReportResultsResponseSummary) RawJSON() string { return r.JSON.raw }
func (r *AudienceConversionReportResultsResponseSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudienceConversionReportResultsResponseTimesery struct {
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
func (r AudienceConversionReportResultsResponseTimesery) RawJSON() string { return r.JSON.raw }
func (r *AudienceConversionReportResultsResponseTimesery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudienceConversionReportResultsResponsePreviousSummary struct {
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
func (r AudienceConversionReportResultsResponsePreviousSummary) RawJSON() string { return r.JSON.raw }
func (r *AudienceConversionReportResultsResponsePreviousSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudienceConversionReportNewParams struct {
	AttributionWindow string                                    `json:"attributionWindow" api:"required"`
	EventName         string                                    `json:"eventName" api:"required"`
	Name              string                                    `json:"name" api:"required"`
	ValueProperty     string                                    `json:"valueProperty" api:"required"`
	DateFrom          param.Opt[string]                         `json:"dateFrom,omitzero"`
	DateTo            param.Opt[string]                         `json:"dateTo,omitzero"`
	ExcludeBots       param.Opt[bool]                           `json:"excludeBots,omitzero"`
	WebSourceID       param.Opt[string]                         `json:"webSourceId,omitzero" format:"uuid"`
	Filters           []AudienceConversionReportNewParamsFilter `json:"filters,omitzero"`
	paramObj
}

func (r AudienceConversionReportNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AudienceConversionReportNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AudienceConversionReportNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Dimension is required.
type AudienceConversionReportNewParamsFilter struct {
	// Any of "browser", "campaign", "city", "content", "country", "device",
	// "entry_page", "exit_page", "medium", "os", "page", "referrer", "region",
	// "source", "term".
	Dimension string            `json:"dimension,omitzero" api:"required"`
	Value     param.Opt[string] `json:"value,omitzero"`
	// Any of "CONTAINS", "IS", "IS_NOT", "NOT_CONTAINS".
	Operator string   `json:"operator,omitzero"`
	Values   []string `json:"values,omitzero"`
	paramObj
}

func (r AudienceConversionReportNewParamsFilter) MarshalJSON() (data []byte, err error) {
	type shadow AudienceConversionReportNewParamsFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AudienceConversionReportNewParamsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AudienceConversionReportNewParamsFilter](
		"dimension", "browser", "campaign", "city", "content", "country", "device", "entry_page", "exit_page", "medium", "os", "page", "referrer", "region", "source", "term",
	)
	apijson.RegisterFieldValidator[AudienceConversionReportNewParamsFilter](
		"operator", "CONTAINS", "IS", "IS_NOT", "NOT_CONTAINS",
	)
}

type AudienceConversionReportUpdateParams struct {
	DateFrom          param.Opt[string]                            `json:"dateFrom,omitzero"`
	DateTo            param.Opt[string]                            `json:"dateTo,omitzero"`
	ExcludeBots       param.Opt[bool]                              `json:"excludeBots,omitzero"`
	WebSourceID       param.Opt[string]                            `json:"webSourceId,omitzero" format:"uuid"`
	AttributionWindow param.Opt[string]                            `json:"attributionWindow,omitzero"`
	EventName         param.Opt[string]                            `json:"eventName,omitzero"`
	Name              param.Opt[string]                            `json:"name,omitzero"`
	ValueProperty     param.Opt[string]                            `json:"valueProperty,omitzero"`
	Filters           []AudienceConversionReportUpdateParamsFilter `json:"filters,omitzero"`
	paramObj
}

func (r AudienceConversionReportUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AudienceConversionReportUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AudienceConversionReportUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Dimension is required.
type AudienceConversionReportUpdateParamsFilter struct {
	// Any of "browser", "campaign", "city", "content", "country", "device",
	// "entry_page", "exit_page", "medium", "os", "page", "referrer", "region",
	// "source", "term".
	Dimension string            `json:"dimension,omitzero" api:"required"`
	Value     param.Opt[string] `json:"value,omitzero"`
	// Any of "CONTAINS", "IS", "IS_NOT", "NOT_CONTAINS".
	Operator string   `json:"operator,omitzero"`
	Values   []string `json:"values,omitzero"`
	paramObj
}

func (r AudienceConversionReportUpdateParamsFilter) MarshalJSON() (data []byte, err error) {
	type shadow AudienceConversionReportUpdateParamsFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AudienceConversionReportUpdateParamsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AudienceConversionReportUpdateParamsFilter](
		"dimension", "browser", "campaign", "city", "content", "country", "device", "entry_page", "exit_page", "medium", "os", "page", "referrer", "region", "source", "term",
	)
	apijson.RegisterFieldValidator[AudienceConversionReportUpdateParamsFilter](
		"operator", "CONTAINS", "IS", "IS_NOT", "NOT_CONTAINS",
	)
}

type AudienceConversionReportResultsParams struct {
	From param.Opt[string] `query:"from,omitzero" json:"-"`
	To   param.Opt[string] `query:"to,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AudienceConversionReportResultsParams]'s query parameters
// as `url.Values`.
func (r AudienceConversionReportResultsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
