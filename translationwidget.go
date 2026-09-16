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

// List every translation widget configured on the account, including the domains
// it runs on, its appearance settings, and the languages it offers. Not paginated
// — widgets are capped by the account's translation widget limit. Requires scope:
// translationWidget:list
func (r *TranslationWidgetService) List(ctx context.Context, opts ...option.RequestOption) (res *TranslationWidgetListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/translation-widgets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Create a translation widget. Every field is optional — anything omitted comes
// back as `null` and the widget falls back to its built-in appearance, and
// omitting `enabledLanguages` offers every supported language. Returns the full
// widget, including its id, so it can be installed without a follow-up request.
// Requires scope: translationWidget:create
func (r *TranslationWidgetService) New(ctx context.Context, body TranslationWidgetNewParams, opts ...option.RequestOption) (res *TranslationWidgetNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/translation-widgets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Fetch one translation widget by its id. Returns 404 when it does not exist.
// Requires scope: translationWidget:find
func (r *TranslationWidgetService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *TranslationWidgetGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/translation-widgets/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
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

type TranslationWidgetListResponse struct {
	// Every translation widget on the account. Not paginated — widgets are capped by
	// the account's translation widget limit, so the full set always fits in one
	// response.
	Entities []TranslationWidgetListResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranslationWidgetListResponse) RawJSON() string { return r.JSON.raw }
func (r *TranslationWidgetListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TranslationWidgetListResponseEntity struct {
	// Unique identifier for the translation widget.
	ID string `json:"id" api:"required"`
	// ISO-8601 timestamp when the widget was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Accent colour for the widget button and modal, as a hex value — `#RGB`,
	// `#RRGGBB`, or `#RRGGBBAA`.
	BrandColor string `json:"brandColor" api:"nullable"`
	// Custom domain used to serve the widget script (e.g. `translate.example.com`).
	// Leave null to use the default Ours Privacy domain. Send the bare host — no
	// scheme, port, or path.
	CustomDomain string `json:"customDomain" api:"nullable"`
	// Languages offered in the widget, as ISO codes such as `en`, `es`, `fr`, `zh-TW`.
	// Omit or send an empty array to offer every supported language. An unrecognised
	// code is rejected with 400 rather than silently ignored.
	EnabledLanguages []string `json:"enabledLanguages" api:"nullable"`
	// Layout of the language-selection modal: `standard` list, `grid` of languages, or
	// a `dropdown`.
	//
	// Any of "standard", "grid", "dropdown".
	ModalVariant string `json:"modalVariant" api:"nullable"`
	// Terms the widget keeps in their original form in every language — brand names,
	// product names, and the like. Up to 500 terms of 200 characters each; duplicates
	// and purely numeric terms are dropped.
	NoTranslateTerms []string `json:"noTranslateTerms" api:"nullable"`
	// Corner of the viewport the widget button is anchored to.
	//
	// Any of "bottom-right", "bottom-left", "top-right", "top-left".
	Position string `json:"position" api:"nullable"`
	// Colour scheme the widget renders in.
	//
	// Any of "light", "dark".
	Theme string `json:"theme" api:"nullable"`
	// ISO-8601 timestamp of the most recent change. Matches `createdAt` on a widget
	// that has never been edited.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// Domains where this widget is allowed to run. The widget refuses to load anywhere
	// else, so a missing host means the widget never appears. Send bare hosts — no
	// scheme, port, or path.
	WhitelistedDomains []string `json:"whitelistedDomains" api:"nullable"`
	// Visual style of the widget button: `compact` for an icon-sized button,
	// `extended` for a labelled button, `minimal` for the least intrusive treatment.
	//
	// Any of "compact", "extended", "minimal".
	WidgetVariant string `json:"widgetVariant" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		BrandColor         respjson.Field
		CustomDomain       respjson.Field
		EnabledLanguages   respjson.Field
		ModalVariant       respjson.Field
		NoTranslateTerms   respjson.Field
		Position           respjson.Field
		Theme              respjson.Field
		UpdatedAt          respjson.Field
		WhitelistedDomains respjson.Field
		WidgetVariant      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranslationWidgetListResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *TranslationWidgetListResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TranslationWidgetNewResponse struct {
	// Unique identifier for the translation widget.
	ID string `json:"id" api:"required"`
	// ISO-8601 timestamp when the widget was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Accent colour for the widget button and modal, as a hex value — `#RGB`,
	// `#RRGGBB`, or `#RRGGBBAA`.
	BrandColor string `json:"brandColor" api:"nullable"`
	// Custom domain used to serve the widget script (e.g. `translate.example.com`).
	// Leave null to use the default Ours Privacy domain. Send the bare host — no
	// scheme, port, or path.
	CustomDomain string `json:"customDomain" api:"nullable"`
	// Languages offered in the widget, as ISO codes such as `en`, `es`, `fr`, `zh-TW`.
	// Omit or send an empty array to offer every supported language. An unrecognised
	// code is rejected with 400 rather than silently ignored.
	EnabledLanguages []string `json:"enabledLanguages" api:"nullable"`
	// Layout of the language-selection modal: `standard` list, `grid` of languages, or
	// a `dropdown`.
	//
	// Any of "standard", "grid", "dropdown".
	ModalVariant TranslationWidgetNewResponseModalVariant `json:"modalVariant" api:"nullable"`
	// Terms the widget keeps in their original form in every language — brand names,
	// product names, and the like. Up to 500 terms of 200 characters each; duplicates
	// and purely numeric terms are dropped.
	NoTranslateTerms []string `json:"noTranslateTerms" api:"nullable"`
	// Corner of the viewport the widget button is anchored to.
	//
	// Any of "bottom-right", "bottom-left", "top-right", "top-left".
	Position TranslationWidgetNewResponsePosition `json:"position" api:"nullable"`
	// Colour scheme the widget renders in.
	//
	// Any of "light", "dark".
	Theme TranslationWidgetNewResponseTheme `json:"theme" api:"nullable"`
	// ISO-8601 timestamp of the most recent change. Matches `createdAt` on a widget
	// that has never been edited.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// Domains where this widget is allowed to run. The widget refuses to load anywhere
	// else, so a missing host means the widget never appears. Send bare hosts — no
	// scheme, port, or path.
	WhitelistedDomains []string `json:"whitelistedDomains" api:"nullable"`
	// Visual style of the widget button: `compact` for an icon-sized button,
	// `extended` for a labelled button, `minimal` for the least intrusive treatment.
	//
	// Any of "compact", "extended", "minimal".
	WidgetVariant TranslationWidgetNewResponseWidgetVariant `json:"widgetVariant" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		BrandColor         respjson.Field
		CustomDomain       respjson.Field
		EnabledLanguages   respjson.Field
		ModalVariant       respjson.Field
		NoTranslateTerms   respjson.Field
		Position           respjson.Field
		Theme              respjson.Field
		UpdatedAt          respjson.Field
		WhitelistedDomains respjson.Field
		WidgetVariant      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranslationWidgetNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TranslationWidgetNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Layout of the language-selection modal: `standard` list, `grid` of languages, or
// a `dropdown`.
type TranslationWidgetNewResponseModalVariant string

const (
	TranslationWidgetNewResponseModalVariantStandard TranslationWidgetNewResponseModalVariant = "standard"
	TranslationWidgetNewResponseModalVariantGrid     TranslationWidgetNewResponseModalVariant = "grid"
	TranslationWidgetNewResponseModalVariantDropdown TranslationWidgetNewResponseModalVariant = "dropdown"
)

// Corner of the viewport the widget button is anchored to.
type TranslationWidgetNewResponsePosition string

const (
	TranslationWidgetNewResponsePositionBottomRight TranslationWidgetNewResponsePosition = "bottom-right"
	TranslationWidgetNewResponsePositionBottomLeft  TranslationWidgetNewResponsePosition = "bottom-left"
	TranslationWidgetNewResponsePositionTopRight    TranslationWidgetNewResponsePosition = "top-right"
	TranslationWidgetNewResponsePositionTopLeft     TranslationWidgetNewResponsePosition = "top-left"
)

// Colour scheme the widget renders in.
type TranslationWidgetNewResponseTheme string

const (
	TranslationWidgetNewResponseThemeLight TranslationWidgetNewResponseTheme = "light"
	TranslationWidgetNewResponseThemeDark  TranslationWidgetNewResponseTheme = "dark"
)

// Visual style of the widget button: `compact` for an icon-sized button,
// `extended` for a labelled button, `minimal` for the least intrusive treatment.
type TranslationWidgetNewResponseWidgetVariant string

const (
	TranslationWidgetNewResponseWidgetVariantCompact  TranslationWidgetNewResponseWidgetVariant = "compact"
	TranslationWidgetNewResponseWidgetVariantExtended TranslationWidgetNewResponseWidgetVariant = "extended"
	TranslationWidgetNewResponseWidgetVariantMinimal  TranslationWidgetNewResponseWidgetVariant = "minimal"
)

type TranslationWidgetGetResponse struct {
	// Unique identifier for the translation widget.
	ID string `json:"id" api:"required"`
	// ISO-8601 timestamp when the widget was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Accent colour for the widget button and modal, as a hex value — `#RGB`,
	// `#RRGGBB`, or `#RRGGBBAA`.
	BrandColor string `json:"brandColor" api:"nullable"`
	// Custom domain used to serve the widget script (e.g. `translate.example.com`).
	// Leave null to use the default Ours Privacy domain. Send the bare host — no
	// scheme, port, or path.
	CustomDomain string `json:"customDomain" api:"nullable"`
	// Languages offered in the widget, as ISO codes such as `en`, `es`, `fr`, `zh-TW`.
	// Omit or send an empty array to offer every supported language. An unrecognised
	// code is rejected with 400 rather than silently ignored.
	EnabledLanguages []string `json:"enabledLanguages" api:"nullable"`
	// Layout of the language-selection modal: `standard` list, `grid` of languages, or
	// a `dropdown`.
	//
	// Any of "standard", "grid", "dropdown".
	ModalVariant TranslationWidgetGetResponseModalVariant `json:"modalVariant" api:"nullable"`
	// Terms the widget keeps in their original form in every language — brand names,
	// product names, and the like. Up to 500 terms of 200 characters each; duplicates
	// and purely numeric terms are dropped.
	NoTranslateTerms []string `json:"noTranslateTerms" api:"nullable"`
	// Corner of the viewport the widget button is anchored to.
	//
	// Any of "bottom-right", "bottom-left", "top-right", "top-left".
	Position TranslationWidgetGetResponsePosition `json:"position" api:"nullable"`
	// Colour scheme the widget renders in.
	//
	// Any of "light", "dark".
	Theme TranslationWidgetGetResponseTheme `json:"theme" api:"nullable"`
	// ISO-8601 timestamp of the most recent change. Matches `createdAt` on a widget
	// that has never been edited.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// Domains where this widget is allowed to run. The widget refuses to load anywhere
	// else, so a missing host means the widget never appears. Send bare hosts — no
	// scheme, port, or path.
	WhitelistedDomains []string `json:"whitelistedDomains" api:"nullable"`
	// Visual style of the widget button: `compact` for an icon-sized button,
	// `extended` for a labelled button, `minimal` for the least intrusive treatment.
	//
	// Any of "compact", "extended", "minimal".
	WidgetVariant TranslationWidgetGetResponseWidgetVariant `json:"widgetVariant" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		BrandColor         respjson.Field
		CustomDomain       respjson.Field
		EnabledLanguages   respjson.Field
		ModalVariant       respjson.Field
		NoTranslateTerms   respjson.Field
		Position           respjson.Field
		Theme              respjson.Field
		UpdatedAt          respjson.Field
		WhitelistedDomains respjson.Field
		WidgetVariant      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TranslationWidgetGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TranslationWidgetGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Layout of the language-selection modal: `standard` list, `grid` of languages, or
// a `dropdown`.
type TranslationWidgetGetResponseModalVariant string

const (
	TranslationWidgetGetResponseModalVariantStandard TranslationWidgetGetResponseModalVariant = "standard"
	TranslationWidgetGetResponseModalVariantGrid     TranslationWidgetGetResponseModalVariant = "grid"
	TranslationWidgetGetResponseModalVariantDropdown TranslationWidgetGetResponseModalVariant = "dropdown"
)

// Corner of the viewport the widget button is anchored to.
type TranslationWidgetGetResponsePosition string

const (
	TranslationWidgetGetResponsePositionBottomRight TranslationWidgetGetResponsePosition = "bottom-right"
	TranslationWidgetGetResponsePositionBottomLeft  TranslationWidgetGetResponsePosition = "bottom-left"
	TranslationWidgetGetResponsePositionTopRight    TranslationWidgetGetResponsePosition = "top-right"
	TranslationWidgetGetResponsePositionTopLeft     TranslationWidgetGetResponsePosition = "top-left"
)

// Colour scheme the widget renders in.
type TranslationWidgetGetResponseTheme string

const (
	TranslationWidgetGetResponseThemeLight TranslationWidgetGetResponseTheme = "light"
	TranslationWidgetGetResponseThemeDark  TranslationWidgetGetResponseTheme = "dark"
)

// Visual style of the widget button: `compact` for an icon-sized button,
// `extended` for a labelled button, `minimal` for the least intrusive treatment.
type TranslationWidgetGetResponseWidgetVariant string

const (
	TranslationWidgetGetResponseWidgetVariantCompact  TranslationWidgetGetResponseWidgetVariant = "compact"
	TranslationWidgetGetResponseWidgetVariantExtended TranslationWidgetGetResponseWidgetVariant = "extended"
	TranslationWidgetGetResponseWidgetVariantMinimal  TranslationWidgetGetResponseWidgetVariant = "minimal"
)

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

type TranslationWidgetNewParams struct {
	// Accent colour for the widget button and modal, as a hex value — `#RGB`,
	// `#RRGGBB`, or `#RRGGBBAA`.
	BrandColor param.Opt[string] `json:"brandColor,omitzero"`
	// Custom domain used to serve the widget script (e.g. `translate.example.com`).
	// Leave null to use the default Ours Privacy domain. Send the bare host — no
	// scheme, port, or path.
	CustomDomain param.Opt[string] `json:"customDomain,omitzero"`
	// Languages offered in the widget, as ISO codes such as `en`, `es`, `fr`, `zh-TW`.
	// Omit or send an empty array to offer every supported language. An unrecognised
	// code is rejected with 400 rather than silently ignored.
	//
	// Any of "af", "sq", "am", "ar", "hy", "az", "bn", "bs", "bg", "ca", "zh",
	// "zh-TW", "hr", "cs", "da", "fa-AF", "nl", "en", "et", "fa", "tl", "fi", "fr",
	// "fr-CA", "ka", "de", "el", "gu", "ht", "ha", "he", "hi", "hu", "is", "id", "ga",
	// "it", "ja", "kn", "kk", "ko", "lv", "lt", "mk", "ms", "ml", "mt", "mr", "mn",
	// "no", "ps", "pl", "pt", "pt-PT", "pa", "ro", "ru", "sr", "si", "sk", "sl", "so",
	// "es", "es-MX", "sw", "sv", "ta", "te", "th", "tr", "uk", "ur", "uz", "vi", "cy".
	EnabledLanguages []string `json:"enabledLanguages,omitzero"`
	// Terms the widget keeps in their original form in every language — brand names,
	// product names, and the like. Up to 500 terms of 200 characters each; duplicates
	// and purely numeric terms are dropped.
	NoTranslateTerms []string `json:"noTranslateTerms,omitzero"`
	// Domains where this widget is allowed to run. The widget refuses to load anywhere
	// else, so a missing host means the widget never appears. Send bare hosts — no
	// scheme, port, or path.
	WhitelistedDomains []string `json:"whitelistedDomains,omitzero"`
	// Layout of the language-selection modal: `standard` list, `grid` of languages, or
	// a `dropdown`.
	//
	// Any of "standard", "grid", "dropdown".
	ModalVariant TranslationWidgetNewParamsModalVariant `json:"modalVariant,omitzero"`
	// Corner of the viewport the widget button is anchored to.
	//
	// Any of "bottom-right", "bottom-left", "top-right", "top-left".
	Position TranslationWidgetNewParamsPosition `json:"position,omitzero"`
	// Colour scheme the widget renders in.
	//
	// Any of "light", "dark".
	Theme TranslationWidgetNewParamsTheme `json:"theme,omitzero"`
	// Visual style of the widget button: `compact` for an icon-sized button,
	// `extended` for a labelled button, `minimal` for the least intrusive treatment.
	//
	// Any of "compact", "extended", "minimal".
	WidgetVariant TranslationWidgetNewParamsWidgetVariant `json:"widgetVariant,omitzero"`
	paramObj
}

func (r TranslationWidgetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TranslationWidgetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TranslationWidgetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Layout of the language-selection modal: `standard` list, `grid` of languages, or
// a `dropdown`.
type TranslationWidgetNewParamsModalVariant string

const (
	TranslationWidgetNewParamsModalVariantStandard TranslationWidgetNewParamsModalVariant = "standard"
	TranslationWidgetNewParamsModalVariantGrid     TranslationWidgetNewParamsModalVariant = "grid"
	TranslationWidgetNewParamsModalVariantDropdown TranslationWidgetNewParamsModalVariant = "dropdown"
)

// Corner of the viewport the widget button is anchored to.
type TranslationWidgetNewParamsPosition string

const (
	TranslationWidgetNewParamsPositionBottomRight TranslationWidgetNewParamsPosition = "bottom-right"
	TranslationWidgetNewParamsPositionBottomLeft  TranslationWidgetNewParamsPosition = "bottom-left"
	TranslationWidgetNewParamsPositionTopRight    TranslationWidgetNewParamsPosition = "top-right"
	TranslationWidgetNewParamsPositionTopLeft     TranslationWidgetNewParamsPosition = "top-left"
)

// Colour scheme the widget renders in.
type TranslationWidgetNewParamsTheme string

const (
	TranslationWidgetNewParamsThemeLight TranslationWidgetNewParamsTheme = "light"
	TranslationWidgetNewParamsThemeDark  TranslationWidgetNewParamsTheme = "dark"
)

// Visual style of the widget button: `compact` for an icon-sized button,
// `extended` for a labelled button, `minimal` for the least intrusive treatment.
type TranslationWidgetNewParamsWidgetVariant string

const (
	TranslationWidgetNewParamsWidgetVariantCompact  TranslationWidgetNewParamsWidgetVariant = "compact"
	TranslationWidgetNewParamsWidgetVariantExtended TranslationWidgetNewParamsWidgetVariant = "extended"
	TranslationWidgetNewParamsWidgetVariantMinimal  TranslationWidgetNewParamsWidgetVariant = "minimal"
)

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
