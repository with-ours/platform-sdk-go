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
	"github.com/with-ours/platform-sdk-go/packages/pagination"
	"github.com/with-ours/platform-sdk-go/packages/param"
	"github.com/with-ours/platform-sdk-go/packages/respjson"
)

// ConversionJourneySummaryService contains methods and other services that help
// with interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConversionJourneySummaryService] method instead.
type ConversionJourneySummaryService struct {
	Options []option.RequestOption
}

// NewConversionJourneySummaryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewConversionJourneySummaryService(opts ...option.RequestOption) (r ConversionJourneySummaryService) {
	r = ConversionJourneySummaryService{}
	r.Options = opts
	return
}

// List saved Conversion Journey Summary configurations, most recently updated
// first. Supports cursor pagination. Each result contains the conversion event,
// analysis window, attribution window, filters, and bot/source settings needed to
// reopen the saved analysis. Requires scope: web-analytics:view
func (r *ConversionJourneySummaryService) List(ctx context.Context, query ConversionJourneySummaryListParams, opts ...option.RequestOption) (res *pagination.Cursor[ConversionJourneySummaryListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "rest/v1/conversion-journey-summaries"
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

// List saved Conversion Journey Summary configurations, most recently updated
// first. Supports cursor pagination. Each result contains the conversion event,
// analysis window, attribution window, filters, and bot/source settings needed to
// reopen the saved analysis. Requires scope: web-analytics:view
func (r *ConversionJourneySummaryService) ListAutoPaging(ctx context.Context, query ConversionJourneySummaryListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[ConversionJourneySummaryListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Save a named Conversion Journey Summary configuration. Returns the full saved
// summary so callers can reopen the same analysis without a follow-up request.
// Each account can save up to 100 summaries. Requires scope: web-analytics:write
func (r *ConversionJourneySummaryService) New(ctx context.Context, body ConversionJourneySummaryNewParams, opts ...option.RequestOption) (res *ConversionJourneySummaryNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/conversion-journey-summaries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Fetch a saved Conversion Journey Summary by its id. Returns 404 when it does not
// exist. Requires scope: web-analytics:view
func (r *ConversionJourneySummaryService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ConversionJourneySummaryGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/conversion-journey-summaries/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update one or more fields on a saved Conversion Journey Summary. Omitted fields
// remain unchanged. When provided, `filters` replaces the complete saved filter
// list. Send `null` for `webSourceId` or `excludeBots` to clear that optional
// setting. Requires scope: web-analytics:write
func (r *ConversionJourneySummaryService) Update(ctx context.Context, id string, body ConversionJourneySummaryUpdateParams, opts ...option.RequestOption) (res *ConversionJourneySummaryUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/conversion-journey-summaries/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete a saved Conversion Journey Summary. The underlying analytics data is
// unaffected. Requires scope: web-analytics:write
func (r *ConversionJourneySummaryService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *ConversionJourneySummaryDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/conversion-journey-summaries/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type ConversionJourneySummaryListResponse struct {
	CreatedAt   string                                       `json:"createdAt" api:"required"`
	DateFrom    string                                       `json:"dateFrom" api:"required"`
	DateTo      string                                       `json:"dateTo" api:"required"`
	EventName   string                                       `json:"eventName" api:"required"`
	Filters     []ConversionJourneySummaryListResponseFilter `json:"filters" api:"required"`
	Name        string                                       `json:"name" api:"required"`
	SummaryID   string                                       `json:"summaryId" api:"required" format:"uuid"`
	UpdatedAt   string                                       `json:"updatedAt" api:"required"`
	WindowDays  int64                                        `json:"windowDays" api:"required"`
	ExcludeBots bool                                         `json:"excludeBots" api:"nullable"`
	WebSourceID string                                       `json:"webSourceId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		DateFrom    respjson.Field
		DateTo      respjson.Field
		EventName   respjson.Field
		Filters     respjson.Field
		Name        respjson.Field
		SummaryID   respjson.Field
		UpdatedAt   respjson.Field
		WindowDays  respjson.Field
		ExcludeBots respjson.Field
		WebSourceID respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversionJourneySummaryListResponse) RawJSON() string { return r.JSON.raw }
func (r *ConversionJourneySummaryListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversionJourneySummaryListResponseFilter struct {
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
func (r ConversionJourneySummaryListResponseFilter) RawJSON() string { return r.JSON.raw }
func (r *ConversionJourneySummaryListResponseFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversionJourneySummaryNewResponse struct {
	CreatedAt   string                                      `json:"createdAt" api:"required"`
	DateFrom    string                                      `json:"dateFrom" api:"required"`
	DateTo      string                                      `json:"dateTo" api:"required"`
	EventName   string                                      `json:"eventName" api:"required"`
	Filters     []ConversionJourneySummaryNewResponseFilter `json:"filters" api:"required"`
	Name        string                                      `json:"name" api:"required"`
	SummaryID   string                                      `json:"summaryId" api:"required" format:"uuid"`
	UpdatedAt   string                                      `json:"updatedAt" api:"required"`
	WindowDays  int64                                       `json:"windowDays" api:"required"`
	ExcludeBots bool                                        `json:"excludeBots" api:"nullable"`
	WebSourceID string                                      `json:"webSourceId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		DateFrom    respjson.Field
		DateTo      respjson.Field
		EventName   respjson.Field
		Filters     respjson.Field
		Name        respjson.Field
		SummaryID   respjson.Field
		UpdatedAt   respjson.Field
		WindowDays  respjson.Field
		ExcludeBots respjson.Field
		WebSourceID respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversionJourneySummaryNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ConversionJourneySummaryNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversionJourneySummaryNewResponseFilter struct {
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
func (r ConversionJourneySummaryNewResponseFilter) RawJSON() string { return r.JSON.raw }
func (r *ConversionJourneySummaryNewResponseFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversionJourneySummaryGetResponse struct {
	CreatedAt   string                                      `json:"createdAt" api:"required"`
	DateFrom    string                                      `json:"dateFrom" api:"required"`
	DateTo      string                                      `json:"dateTo" api:"required"`
	EventName   string                                      `json:"eventName" api:"required"`
	Filters     []ConversionJourneySummaryGetResponseFilter `json:"filters" api:"required"`
	Name        string                                      `json:"name" api:"required"`
	SummaryID   string                                      `json:"summaryId" api:"required" format:"uuid"`
	UpdatedAt   string                                      `json:"updatedAt" api:"required"`
	WindowDays  int64                                       `json:"windowDays" api:"required"`
	ExcludeBots bool                                        `json:"excludeBots" api:"nullable"`
	WebSourceID string                                      `json:"webSourceId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		DateFrom    respjson.Field
		DateTo      respjson.Field
		EventName   respjson.Field
		Filters     respjson.Field
		Name        respjson.Field
		SummaryID   respjson.Field
		UpdatedAt   respjson.Field
		WindowDays  respjson.Field
		ExcludeBots respjson.Field
		WebSourceID respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversionJourneySummaryGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ConversionJourneySummaryGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversionJourneySummaryGetResponseFilter struct {
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
func (r ConversionJourneySummaryGetResponseFilter) RawJSON() string { return r.JSON.raw }
func (r *ConversionJourneySummaryGetResponseFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversionJourneySummaryUpdateResponse struct {
	CreatedAt   string                                         `json:"createdAt" api:"required"`
	DateFrom    string                                         `json:"dateFrom" api:"required"`
	DateTo      string                                         `json:"dateTo" api:"required"`
	EventName   string                                         `json:"eventName" api:"required"`
	Filters     []ConversionJourneySummaryUpdateResponseFilter `json:"filters" api:"required"`
	Name        string                                         `json:"name" api:"required"`
	SummaryID   string                                         `json:"summaryId" api:"required" format:"uuid"`
	UpdatedAt   string                                         `json:"updatedAt" api:"required"`
	WindowDays  int64                                          `json:"windowDays" api:"required"`
	ExcludeBots bool                                           `json:"excludeBots" api:"nullable"`
	WebSourceID string                                         `json:"webSourceId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		DateFrom    respjson.Field
		DateTo      respjson.Field
		EventName   respjson.Field
		Filters     respjson.Field
		Name        respjson.Field
		SummaryID   respjson.Field
		UpdatedAt   respjson.Field
		WindowDays  respjson.Field
		ExcludeBots respjson.Field
		WebSourceID respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversionJourneySummaryUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ConversionJourneySummaryUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversionJourneySummaryUpdateResponseFilter struct {
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
func (r ConversionJourneySummaryUpdateResponseFilter) RawJSON() string { return r.JSON.raw }
func (r *ConversionJourneySummaryUpdateResponseFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversionJourneySummaryDeleteResponse struct {
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
func (r ConversionJourneySummaryDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ConversionJourneySummaryDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversionJourneySummaryListParams struct {
	// Maximum number of items to return. Defaults to 25; values below 1 are clamped to
	// 1 and values above 100 are clamped to 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Opaque pagination cursor from pagination.nextCursor in the previous response. Do
	// not decode or modify it. Malformed cursors return 400 Bad Request.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConversionJourneySummaryListParams]'s query parameters as
// `url.Values`.
func (r ConversionJourneySummaryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConversionJourneySummaryNewParams struct {
	// Inclusive start of the saved analysis window in `YYYY-MM-DD` format.
	DateFrom string `json:"dateFrom" api:"required"`
	// Inclusive end of the saved analysis window in `YYYY-MM-DD` format.
	DateTo      string                                    `json:"dateTo" api:"required"`
	EventName   string                                    `json:"eventName" api:"required"`
	Name        string                                    `json:"name" api:"required"`
	WindowDays  int64                                     `json:"windowDays" api:"required"`
	ExcludeBots param.Opt[bool]                           `json:"excludeBots,omitzero"`
	WebSourceID param.Opt[string]                         `json:"webSourceId,omitzero" format:"uuid"`
	Filters     []ConversionJourneySummaryNewParamsFilter `json:"filters,omitzero"`
	paramObj
}

func (r ConversionJourneySummaryNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ConversionJourneySummaryNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversionJourneySummaryNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Dimension is required.
type ConversionJourneySummaryNewParamsFilter struct {
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

func (r ConversionJourneySummaryNewParamsFilter) MarshalJSON() (data []byte, err error) {
	type shadow ConversionJourneySummaryNewParamsFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversionJourneySummaryNewParamsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversionJourneySummaryNewParamsFilter](
		"dimension", "browser", "campaign", "city", "content", "country", "device", "entry_page", "exit_page", "medium", "os", "page", "referrer", "region", "source", "term",
	)
	apijson.RegisterFieldValidator[ConversionJourneySummaryNewParamsFilter](
		"operator", "CONTAINS", "IS", "IS_NOT", "NOT_CONTAINS",
	)
}

type ConversionJourneySummaryUpdateParams struct {
	ExcludeBots param.Opt[bool]                              `json:"excludeBots,omitzero"`
	WebSourceID param.Opt[string]                            `json:"webSourceId,omitzero" format:"uuid"`
	DateFrom    param.Opt[string]                            `json:"dateFrom,omitzero"`
	DateTo      param.Opt[string]                            `json:"dateTo,omitzero"`
	EventName   param.Opt[string]                            `json:"eventName,omitzero"`
	Name        param.Opt[string]                            `json:"name,omitzero"`
	WindowDays  param.Opt[int64]                             `json:"windowDays,omitzero"`
	Filters     []ConversionJourneySummaryUpdateParamsFilter `json:"filters,omitzero"`
	paramObj
}

func (r ConversionJourneySummaryUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ConversionJourneySummaryUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversionJourneySummaryUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Dimension is required.
type ConversionJourneySummaryUpdateParamsFilter struct {
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

func (r ConversionJourneySummaryUpdateParamsFilter) MarshalJSON() (data []byte, err error) {
	type shadow ConversionJourneySummaryUpdateParamsFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversionJourneySummaryUpdateParamsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversionJourneySummaryUpdateParamsFilter](
		"dimension", "browser", "campaign", "city", "content", "country", "device", "entry_page", "exit_page", "medium", "os", "page", "referrer", "region", "source", "term",
	)
	apijson.RegisterFieldValidator[ConversionJourneySummaryUpdateParamsFilter](
		"operator", "CONTAINS", "IS", "IS_NOT", "NOT_CONTAINS",
	)
}
