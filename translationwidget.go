// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy

import (
	"context"
	"errors"
	"fmt"
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

// TranslationWidgetService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTranslationWidgetService] method instead.
type TranslationWidgetService struct {
	Options []option.RequestOption
}

// NewTranslationWidgetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTranslationWidgetService(opts ...option.RequestOption) (r TranslationWidgetService) {
	r = TranslationWidgetService{}
	r.Options = opts
	return
}

// Return usage totals and language, host, and page breakdowns for one translation
// widget over the requested date range. Requires scope:
// report:translation-analytics
func (r *TranslationWidgetService) Analytics(ctx context.Context, id string, query TranslationWidgetAnalyticsParams, opts ...option.RequestOption) (res *TranslationWidgetAnalyticsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/translation-widgets/%s/analytics", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type TranslationWidgetAnalyticsResponse struct {
	ByHost            []TranslationWidgetAnalyticsResponseByHost     `json:"byHost" api:"required"`
	ByLanguage        []TranslationWidgetAnalyticsResponseByLanguage `json:"byLanguage" api:"required"`
	LanguagesUsed     int64                                          `json:"languagesUsed" api:"required"`
	TopPages          []TranslationWidgetAnalyticsResponseTopPage    `json:"topPages" api:"required"`
	TotalTranslations int64                                          `json:"totalTranslations" api:"required"`
	UniqueUsers       int64                                          `json:"uniqueUsers" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByHost            respjson.Field
		ByLanguage        respjson.Field
		LanguagesUsed     respjson.Field
		TopPages          respjson.Field
		TotalTranslations respjson.Field
		UniqueUsers       respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranslationWidgetAnalyticsResponse) RawJSON() string { return r.JSON.raw }
func (r *TranslationWidgetAnalyticsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TranslationWidgetAnalyticsResponseByHost struct {
	Host         string `json:"host" api:"required"`
	Translations int64  `json:"translations" api:"required"`
	UniqueUsers  int64  `json:"uniqueUsers" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Host         respjson.Field
		Translations respjson.Field
		UniqueUsers  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranslationWidgetAnalyticsResponseByHost) RawJSON() string { return r.JSON.raw }
func (r *TranslationWidgetAnalyticsResponseByHost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TranslationWidgetAnalyticsResponseByLanguage struct {
	LanguageCode string `json:"languageCode" api:"required"`
	Translations int64  `json:"translations" api:"required"`
	UniqueUsers  int64  `json:"uniqueUsers" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LanguageCode respjson.Field
		Translations respjson.Field
		UniqueUsers  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranslationWidgetAnalyticsResponseByLanguage) RawJSON() string { return r.JSON.raw }
func (r *TranslationWidgetAnalyticsResponseByLanguage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TranslationWidgetAnalyticsResponseTopPage struct {
	ByLanguage   []TranslationWidgetAnalyticsResponseTopPageByLanguage `json:"byLanguage" api:"required"`
	Translations int64                                                 `json:"translations" api:"required"`
	UniqueUsers  int64                                                 `json:"uniqueUsers" api:"required"`
	URL          string                                                `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByLanguage   respjson.Field
		Translations respjson.Field
		UniqueUsers  respjson.Field
		URL          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranslationWidgetAnalyticsResponseTopPage) RawJSON() string { return r.JSON.raw }
func (r *TranslationWidgetAnalyticsResponseTopPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TranslationWidgetAnalyticsResponseTopPageByLanguage struct {
	LanguageCode string `json:"languageCode" api:"required"`
	Translations int64  `json:"translations" api:"required"`
	UniqueUsers  int64  `json:"uniqueUsers" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LanguageCode respjson.Field
		Translations respjson.Field
		UniqueUsers  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranslationWidgetAnalyticsResponseTopPageByLanguage) RawJSON() string { return r.JSON.raw }
func (r *TranslationWidgetAnalyticsResponseTopPageByLanguage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TranslationWidgetAnalyticsParams struct {
	// Inclusive lower bound of the analytics window as `YYYY-MM-DD`.
	From time.Time `query:"from" api:"required" format:"date" json:"-"`
	// Inclusive upper bound of the analytics window as `YYYY-MM-DD`.
	To time.Time `query:"to" api:"required" format:"date" json:"-"`
	// Maximum rows to return for each breakdown. Defaults to 50.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TranslationWidgetAnalyticsParams]'s query parameters as
// `url.Values`.
func (r TranslationWidgetAnalyticsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
