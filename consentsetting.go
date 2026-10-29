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

	"github.com/with-ours/platform-sdk-go/v2/internal/apijson"
	"github.com/with-ours/platform-sdk-go/v2/internal/apiquery"
	"github.com/with-ours/platform-sdk-go/v2/internal/requestconfig"
	"github.com/with-ours/platform-sdk-go/v2/option"
	"github.com/with-ours/platform-sdk-go/v2/packages/param"
	"github.com/with-ours/platform-sdk-go/v2/packages/respjson"
)

// ConsentSettingService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConsentSettingService] method instead.
type ConsentSettingService struct {
	Options []option.RequestOption
}

// NewConsentSettingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConsentSettingService(opts ...option.RequestOption) (r ConsentSettingService) {
	r = ConsentSettingService{}
	r.Options = opts
	return
}

// List all consent settings. Requires API-key scope or current OAuth user
// permission: consentSettings:list
func (r *ConsentSettingService) List(ctx context.Context, opts ...option.RequestOption) (res *ConsentSettingListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/consent-settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Create a new consent setting. POST takes no request body. The new draft is
// Disabled with no allowed domains, an opt-out default rule, English translations,
// necessary/analytics/advertising categories, and no regions. Configure it with
// PATCH or PUT before enabling it. Returns the same fields as GET, including its
// `id`, default rule, and categories. Requires API-key scope or current OAuth user
// permission: consentSettings:create
func (r *ConsentSettingService) New(ctx context.Context, opts ...option.RequestOption) (res *ConsentSettingNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/consent-settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Find a single consent setting by ID. Requires API-key scope or current OAuth
// user permission: consentSettings:find
func (r *ConsentSettingService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ConsentSettingGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/consent-settings/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a consent setting with all required fields. Omitted optional top-level
// fields retain their saved values; the required default rule and lists are
// replaced by supplied values. Send null or [] to clear an optional field when
// allowed. Enabled requires at least one allowed domain in this request. Omitting
// or clearing `whitelistDomains` with null or [] while Enabled returns an HTTP 400
// validation error. Disabled settings may have no allowed domains. Use PATCH for
// partial updates. Requires API-key scope or current OAuth user permission:
// consentSettings:update
func (r *ConsentSettingService) Replace(ctx context.Context, id string, body ConsentSettingReplaceParams, opts ...option.RequestOption) (res *ConsentSettingReplaceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/consent-settings/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Partially update a consent setting. Omitted fields keep their saved values.
// List-valued fields (services, categories, regions) are replaced in full when
// sent. Omitting `whitelistDomains` preserves saved domains; sending it replaces
// the list. Clear allowed domains with null or [] only when the resulting status
// is Disabled. Enabled requires at least one allowed domain, either already saved
// or sent with this request. Invalid combinations return HTTP 400 validation
// errors. Requires API-key scope or current OAuth user permission:
// consentSettings:update
func (r *ConsentSettingService) Update(ctx context.Context, id string, body ConsentSettingUpdateParams, opts ...option.RequestOption) (res *ConsentSettingUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/consent-settings/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete a consent setting. Requires API-key scope or current OAuth user
// permission: consentSettings:delete
func (r *ConsentSettingService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *ConsentSettingDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/consent-settings/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Time-series consent analytics for a single consent settings record: banner
// views, opt-ins, opt-outs, close-icon clicks, and derived opt-in/out rates per
// UTC hour, day, Monday-based week, or calendar month. The window is zero-filled
// so callers get a contiguous series, and rates are person-level
// (`COUNT(DISTINCT visitor_id)`). Use the JSON-encoded shared analytics `filter`
// to scope events by source, page, visitor, or session properties; use
// `compareWithPreviousPeriod=true` to also receive the matching prior window.
// Included days: at most 180 days for `DAILY`, 365 for `WEEKLY`, 365 for
// `MONTHLY`, or 14 for `HOURLY`. Partial buckets identify their actual selected
// dates. Previous-period comparisons are supported for daily and hourly output.
// Requires the API-key scope `report:consent-analytics` (this endpoint returns
// consent analytics report data, which is PHI-bearing and gated separately from
// consent-settings management). Requires API-key scope or current OAuth user
// permission: report:consent-analytics
func (r *ConsentSettingService) Analytics(ctx context.Context, id string, query ConsentSettingAnalyticsParams, opts ...option.RequestOption) (res *ConsentSettingAnalyticsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/consent-settings/%s/analytics", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns daily banner reach, explicit opt-ins and opt-outs, dismissal events, and
// consent rates across the account. Includes historical events without a
// consent-setting ID. Visitor totals are deduplicated across settings and dates;
// dismissals count events. Rates use the larger of banner reach or opt-ins plus
// opt-outs plus dismissals as their denominator, matching Consent Analytics.
// Requires Analytics access and report:consent-analytics. Limited to 60 included
// UTC calendar days. Requires API-key scope or current OAuth user permission:
// report:consent-analytics
func (r *ConsentSettingService) Overview(ctx context.Context, query ConsentSettingOverviewParams, opts ...option.RequestOption) (res *ConsentSettingOverviewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/consent-settings/analytics-overview"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Monthly consent trends over up to 365 included days. Returns unique visitor
// counts within each UTC calendar month, event counts for dismissals, rates, and
// partial-month bounds. Uses the same shared analytics filter as the other consent
// reports. This series-only output omits whole-window totals and breakdowns.
// Requires API-key scope or current OAuth user permission:
// report:consent-analytics
func (r *ConsentSettingService) AnalyticsMonthly(ctx context.Context, id string, query ConsentSettingAnalyticsMonthlyParams, opts ...option.RequestOption) (res *ConsentSettingAnalyticsMonthlyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/consent-settings/%s/analytics-monthly", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns interval limits for a Consent report kind: SUMMARY, TRENDS, or PAGES.
// The organization is determined by the API key. Default included days: at most
// 180 days for `DAILY`, 365 for `WEEKLY`, 365 for `MONTHLY`, or 14 for `HOURLY`;
// organization limits may narrow these allowances. TRENDS requires
// report:consent-analytics, PAGES requires report:consent-page-analysis, and
// SUMMARY requires both. Dates, filters, search and pagination are not part of
// this request. Requires API-key scope or current OAuth user permission:
// report:consent-analytics or report:consent-page-analysis
func (r *ConsentSettingService) AnalyticsCapabilities(ctx context.Context, query ConsentSettingAnalyticsCapabilitiesParams, opts ...option.RequestOption) (res *ConsentSettingAnalyticsCapabilitiesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/consent-settings/analytics-capabilities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Per-page consent breakdown for one consent settings record, ranked by opt-outs
// (descending). Each row bundles banner views, opt-outs, close-icon clicks, and
// the derived opt-out rate. Returns `{ entities, pagination }`. Continue with
// `pagination.nextCursor`; results sort by opt-outs descending, then page path
// ascending. `search` is a substring match against `pathname`; the JSON-encoded
// shared analytics `filter` scopes the underlying events. Requires the API-key
// scope `report:consent-page-analysis` (PHI-bearing report data). Requires API-key
// scope or current OAuth user permission: report:consent-page-analysis
func (r *ConsentSettingService) PageAnalysis(ctx context.Context, id string, query ConsentSettingPageAnalysisParams, opts ...option.RequestOption) (res *ConsentSettingPageAnalysisResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/consent-settings/%s/page-analysis", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Region-grouped consent totals for the window: banner views, opt-ins, opt-outs,
// close-icon clicks, and derived rates per visitor `country_region_name`, ranked
// by banner views (descending). Visitors whose region cannot be resolved (e.g. bot
// traffic, IP geo failure) are bucketed under the literal `Unknown` so per-region
// counts always sum to the global totals. Use the JSON-encoded shared analytics
// `filter` to scope the underlying events. Requires the API-key scope
// `report:consent-analytics` (PHI-bearing report data). Requires API-key scope or
// current OAuth user permission: report:consent-analytics
func (r *ConsentSettingService) AnalyticsByRegion(ctx context.Context, id string, query ConsentSettingAnalyticsByRegionParams, opts ...option.RequestOption) (res *ConsentSettingAnalyticsByRegionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/consent-settings/%s/analytics-by-region", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ConsentSettingListResponse struct {
	Entities []ConsentSettingListResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponse) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingListResponseEntity struct {
	// Server-assigned UUID for this consent settings record.
	ID string `json:"id" api:"required"`
	// Top-level consent categories (e.g. necessary / analytics / advertising). Server
	// re-stamps `priority` to 0..N on write.
	Categories []ConsentSettingListResponseEntityCategory `json:"categories" api:"required"`
	// ISO 8601 timestamp when the record was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Default rule used when the user is not in any region listed in `regions[]`.
	Default ConsentSettingListResponseEntityDefault `json:"default" api:"required"`
	// Discriminator for the entity type. Always "consentSettings".
	Kind string `json:"kind" api:"required"`
	// Human-readable name shown in the dashboard.
	Name string `json:"name" api:"required"`
	// Per-region rule overrides. The first rule whose `regionCode`/`additionalRegions`
	// includes the user's region wins; otherwise `default` applies.
	Regions []ConsentSettingListResponseEntityRegion `json:"regions" api:"required"`
	// Per-service entries powering "show vendors" and category-aware blocking.
	Services []ConsentSettingListResponseEntityService `json:"services" api:"required"`
	// Enabled serves the CMP on its allowed domains and requires at least one.
	// Disabled does not serve the CMP and may have no allowed domains.
	//
	// Any of "Disabled", "Enabled".
	Status string `json:"status" api:"required"`
	// Name of the cookie that stores the user's consent state. Defaults to
	// "op_consent".
	ConsentCookieName string `json:"consentCookieName" api:"nullable"`
	// Optional custom CDN domain for serving the CMP script (e.g.
	// consent.example.com).
	CustomDomain string `json:"customDomain" api:"nullable"`
	// Cookie/localStorage key for the visitor (device) ID. Must exactly match the web
	// SDK's cookie_names.device_id (default "ours_device_id") so consent and SDK
	// events resolve to the same visitor.
	DeviceIDCookieName string `json:"deviceIdCookieName" api:"nullable"`
	// Revision counter. Bump this to force users who already consented to see the
	// modal again (the SDK compares the persisted revision against this value).
	Revision float64 `json:"revision" api:"nullable"`
	// CSS class names that opt scripts out of consent blocking. Each entry must be a
	// single class token (no whitespace).
	SkipBlockingClassNames []string `json:"skipBlockingClassNames" api:"nullable"`
	// ISO 8601 timestamp of the last write. Null on a freshly created record.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// Pixel of the WebSource that this CMP is wired into. Setting this to a token that
	// is not a valid WebSource of yours is rejected; use null to clear the link.
	WebSDKToken string `json:"webSDKToken" api:"nullable"`
	// Domains where this CMP may run. An Enabled CMP requires at least one allowed
	// domain; with none configured, the CMP does not run. Matching domains may share
	// consent across subdomains.
	WhitelistDomains []string `json:"whitelistDomains" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		Categories             respjson.Field
		CreatedAt              respjson.Field
		Default                respjson.Field
		Kind                   respjson.Field
		Name                   respjson.Field
		Regions                respjson.Field
		Services               respjson.Field
		Status                 respjson.Field
		ConsentCookieName      respjson.Field
		CustomDomain           respjson.Field
		DeviceIDCookieName     respjson.Field
		Revision               respjson.Field
		SkipBlockingClassNames respjson.Field
		UpdatedAt              respjson.Field
		WebSDKToken            respjson.Field
		WhitelistDomains       respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingListResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingListResponseEntityCategory struct {
	// Human-readable label shown next to the toggle in the preferences modal.
	Label string `json:"label" api:"required"`
	// Sort key. Lower numbers render first. Server re-stamps to 0..N on write — send
	// any integer, gaps and duplicates are ironed out.
	Priority int64 `json:"priority" api:"required"`
	// Stable identifier referenced by services and translation sections.
	// Conventionally lowercase (e.g. "necessary", "analytics", "advertising").
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		Priority    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityCategory) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingListResponseEntityCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default rule used when the user is not in any region listed in `regions[]`.
type ConsentSettingListResponseEntityDefault struct {
	// Per-category default config for this rule. Every category defined in the
	// top-level `categories[].value` should have an entry here.
	Categories []ConsentSettingListResponseEntityDefaultCategory `json:"categories" api:"required"`
	// BCP 47 default language for this rule. Must have a matching entry in
	// `translations`. Examples: "en", "en-US", "es", "de".
	Language string `json:"language" api:"required"`
	// opt_in: scripts blocked until user accepts (GDPR style). opt_out: scripts run by
	// default until user rejects (CCPA style).
	//
	// Any of "opt_in", "opt_out".
	Mode string `json:"mode" api:"required"`
	// All UI copy, keyed by language. Must include an entry whose `language` matches
	// the rule's `language` field.
	Translations []ConsentSettingListResponseEntityDefaultTranslation `json:"translations" api:"required"`
	// When true, scripts not classified by services[] are blocked until the user opts
	// in.
	AutoblockUnknown bool `json:"autoblockUnknown" api:"nullable"`
	// When true, the consent modal auto-opens on page load.
	AutoShow bool `json:"autoShow" api:"nullable"`
	// Threshold config for autoShowDismissMode (page count or seconds).
	AutoShowDismissConfig ConsentSettingListResponseEntityDefaultAutoShowDismissConfig `json:"autoShowDismissConfig" api:"nullable"`
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	//
	// Any of "after_pages", "after_seconds", "never".
	AutoShowDismissMode string `json:"autoShowDismissMode" api:"nullable"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction bool `json:"disablePageInteraction" api:"nullable"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions ConsentSettingListResponseEntityDefaultGuiOptions `json:"guiOptions" api:"nullable"`
	// When true, the modal is suppressed for known bot user agents.
	HideFromBots bool `json:"hideFromBots" api:"nullable"`
	// When true, the per-service list (services[]) is rendered inside the preferences
	// modal.
	ShowVendorsInPreferences bool `json:"showVendorsInPreferences" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories               respjson.Field
		Language                 respjson.Field
		Mode                     respjson.Field
		Translations             respjson.Field
		AutoblockUnknown         respjson.Field
		AutoShow                 respjson.Field
		AutoShowDismissConfig    respjson.Field
		AutoShowDismissMode      respjson.Field
		DisablePageInteraction   respjson.Field
		GuiOptions               respjson.Field
		HideFromBots             respjson.Field
		ShowVendorsInPreferences respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityDefault) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingListResponseEntityDefault) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingListResponseEntityDefaultCategory struct {
	// Category value (matches `categories[].value`) this entry configures.
	Key   string                                               `json:"key" api:"required"`
	Value ConsentSettingListResponseEntityDefaultCategoryValue `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityDefaultCategory) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingListResponseEntityDefaultCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingListResponseEntityDefaultCategoryValue struct {
	// Whether this category is on by default before the user interacts.
	Enabled bool `json:"enabled" api:"required"`
	// When true, this category defaults off if the browser sends Sec-GPC: 1.
	AutoDisableOnGpc bool `json:"autoDisableOnGPC" api:"nullable"`
	// When true, the user cannot toggle this category in the preferences modal.
	ReadOnly bool `json:"readOnly" api:"nullable"`
	// When true, the page reloads after this category is toggled so newly-allowed
	// scripts can run.
	ReloadPage bool `json:"reloadPage" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled          respjson.Field
		AutoDisableOnGpc respjson.Field
		ReadOnly         respjson.Field
		ReloadPage       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityDefaultCategoryValue) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingListResponseEntityDefaultCategoryValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingListResponseEntityDefaultTranslation struct {
	// BCP 47 language tag identifying which translation this entry provides. Examples:
	// "en", "en-US", "es", "fr-CA". The default rule's `language` must appear here.
	Language string                                                  `json:"language" api:"required"`
	Value    ConsentSettingListResponseEntityDefaultTranslationValue `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Language    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityDefaultTranslation) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingListResponseEntityDefaultTranslation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingListResponseEntityDefaultTranslationValue struct {
	// Translated copy for the initial consent modal.
	ConsentModal ConsentSettingListResponseEntityDefaultTranslationValueConsentModal `json:"consentModal" api:"nullable"`
	// Translated copy for the preferences modal.
	PreferencesModal ConsentSettingListResponseEntityDefaultTranslationValuePreferencesModal `json:"preferencesModal" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConsentModal     respjson.Field
		PreferencesModal respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityDefaultTranslationValue) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingListResponseEntityDefaultTranslationValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the initial consent modal.
type ConsentSettingListResponseEntityDefaultTranslationValueConsentModal struct {
	// Label for the "accept all" button.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Body text shown in the consent modal.
	Description string `json:"description" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the button that opens the preferences modal.
	ShowPreferencesBtn string `json:"showPreferencesBtn" api:"required"`
	// Heading shown at the top of the consent modal.
	Title string `json:"title" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel string `json:"closeIconLabel" api:"nullable"`
	// Optional footer HTML/text shown below the body.
	Footer string `json:"footer" api:"nullable"`
	// Optional text shown when the browser is sending Sec-GPC: 1.
	GpcNotification string `json:"gpcNotification" api:"nullable"`
	// Link label for the privacy policy (defaults to "Privacy Policy").
	PrivacyPolicyLabel string `json:"privacyPolicyLabel" api:"nullable"`
	// Privacy policy link. Accepts an absolute URL (https://...) or a relative path
	// (/legal/privacy) so one config can run across multiple domains. javascript:,
	// data:, and vbscript: schemes are rejected.
	PrivacyPolicyURL string `json:"privacyPolicyUrl" api:"nullable"`
	// Optional message shown when the consent revision has bumped (re-prompt copy).
	RevisionMessage string `json:"revisionMessage" api:"nullable"`
	// Link label for terms of service (defaults to "Terms of Service").
	TermsOfServiceLabel string `json:"termsOfServiceLabel" api:"nullable"`
	// Terms of service link. Same scheme rules as privacyPolicyUrl: absolute or
	// relative path only.
	TermsOfServiceURL string `json:"termsOfServiceUrl" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptAllBtn        respjson.Field
		Description         respjson.Field
		RejectAllBtn        respjson.Field
		ShowPreferencesBtn  respjson.Field
		Title               respjson.Field
		CloseIconLabel      respjson.Field
		Footer              respjson.Field
		GpcNotification     respjson.Field
		PrivacyPolicyLabel  respjson.Field
		PrivacyPolicyURL    respjson.Field
		RevisionMessage     respjson.Field
		TermsOfServiceLabel respjson.Field
		TermsOfServiceURL   respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityDefaultTranslationValueConsentModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingListResponseEntityDefaultTranslationValueConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the preferences modal.
type ConsentSettingListResponseEntityDefaultTranslationValuePreferencesModal struct {
	// Label for the "accept all" button in the preferences modal.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel string `json:"closeIconLabel" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the "save my choices" button.
	SavePreferencesBtn string `json:"savePreferencesBtn" api:"required"`
	// Per-category sections plus a leading general section. The server backfills
	// sections for each defined category if absent.
	Sections []ConsentSettingListResponseEntityDefaultTranslationValuePreferencesModalSection `json:"sections" api:"required"`
	// Heading for the preferences modal.
	Title string `json:"title" api:"required"`
	// Optional collapsible info block rendered after the category sections. Renders
	// only when both title and body are non-empty. Use it for long-form text unrelated
	// to any specific category (e.g. a third-party advertisers disclosure).
	AdditionalInfo ConsentSettingListResponseEntityDefaultTranslationValuePreferencesModalAdditionalInfo `json:"additionalInfo" api:"nullable"`
	// Pluralizable label for the "N services" counter (e.g. "Service|Services").
	ServiceCounterLabel string `json:"serviceCounterLabel" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptAllBtn        respjson.Field
		CloseIconLabel      respjson.Field
		RejectAllBtn        respjson.Field
		SavePreferencesBtn  respjson.Field
		Sections            respjson.Field
		Title               respjson.Field
		AdditionalInfo      respjson.Field
		ServiceCounterLabel respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityDefaultTranslationValuePreferencesModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingListResponseEntityDefaultTranslationValuePreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingListResponseEntityDefaultTranslationValuePreferencesModalSection struct {
	// Body text for this preferences-modal section.
	Description string `json:"description" api:"required"`
	// Heading for this preferences-modal section.
	Title string `json:"title" api:"required"`
	// Optional disclosure table listing cookies/services in this section.
	CookieTable ConsentSettingListResponseEntityDefaultTranslationValuePreferencesModalSectionCookieTable `json:"cookieTable" api:"nullable"`
	// Category value this section configures. Use null/omit for the introductory
	// section that applies to no specific category.
	LinkedCategory string `json:"linkedCategory" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description    respjson.Field
		Title          respjson.Field
		CookieTable    respjson.Field
		LinkedCategory respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityDefaultTranslationValuePreferencesModalSection) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingListResponseEntityDefaultTranslationValuePreferencesModalSection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional disclosure table listing cookies/services in this section.
type ConsentSettingListResponseEntityDefaultTranslationValuePreferencesModalSectionCookieTable struct {
	// Rows of the cookie table — each row is a {key,value} pair (e.g. cookie name →
	// purpose).
	Body []ConsentSettingListResponseEntityDefaultTranslationValuePreferencesModalSectionCookieTableBody `json:"body" api:"required"`
	// Header cells. Same {key,value} shape as rows.
	Headers []ConsentSettingListResponseEntityDefaultTranslationValuePreferencesModalSectionCookieTableHeader `json:"headers" api:"required"`
	// Optional table caption.
	Caption string `json:"caption" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Headers     respjson.Field
		Caption     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityDefaultTranslationValuePreferencesModalSectionCookieTable) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingListResponseEntityDefaultTranslationValuePreferencesModalSectionCookieTable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingListResponseEntityDefaultTranslationValuePreferencesModalSectionCookieTableBody struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityDefaultTranslationValuePreferencesModalSectionCookieTableBody) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingListResponseEntityDefaultTranslationValuePreferencesModalSectionCookieTableBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingListResponseEntityDefaultTranslationValuePreferencesModalSectionCookieTableHeader struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityDefaultTranslationValuePreferencesModalSectionCookieTableHeader) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingListResponseEntityDefaultTranslationValuePreferencesModalSectionCookieTableHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional collapsible info block rendered after the category sections. Renders
// only when both title and body are non-empty. Use it for long-form text unrelated
// to any specific category (e.g. a third-party advertisers disclosure).
type ConsentSettingListResponseEntityDefaultTranslationValuePreferencesModalAdditionalInfo struct {
	// Body text for the additional-info block. Supports the rich-text markdown subset
	// used by section descriptions: line breaks, [text](url) links (https:// or
	// relative paths), and **bold**. HTML is stripped.
	Body string `json:"body" api:"required"`
	// Heading for the additional-info block.
	Title string `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityDefaultTranslationValuePreferencesModalAdditionalInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingListResponseEntityDefaultTranslationValuePreferencesModalAdditionalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Threshold config for autoShowDismissMode (page count or seconds).
type ConsentSettingListResponseEntityDefaultAutoShowDismissConfig struct {
	// Pages-viewed threshold when autoShowDismissMode = after_pages.
	PageCount int64 `json:"pageCount" api:"nullable"`
	// Seconds-on-page threshold when autoShowDismissMode = after_seconds.
	Seconds int64 `json:"seconds" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageCount   respjson.Field
		Seconds     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityDefaultAutoShowDismissConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingListResponseEntityDefaultAutoShowDismissConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the modals (layout/position/colors).
type ConsentSettingListResponseEntityDefaultGuiOptions struct {
	// Visual options for the initial consent modal.
	ConsentModal ConsentSettingListResponseEntityDefaultGuiOptionsConsentModal `json:"consentModal" api:"nullable"`
	// Global CSS variable overrides applied to both modals.
	CssVariables ConsentSettingListResponseEntityDefaultGuiOptionsCssVariables `json:"cssVariables" api:"nullable"`
	// Visual options for the per-category preferences modal.
	PreferencesModal ConsentSettingListResponseEntityDefaultGuiOptionsPreferencesModal `json:"preferencesModal" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConsentModal     respjson.Field
		CssVariables     respjson.Field
		PreferencesModal respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityDefaultGuiOptions) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingListResponseEntityDefaultGuiOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the initial consent modal.
type ConsentSettingListResponseEntityDefaultGuiOptionsConsentModal struct {
	// Which buttons to render in the consent modal and in what arrangement.
	//
	// Any of "AcceptAllNecessaryPreferences", "AcceptAllPreferences", "AcceptOnly",
	// "InformationOnly", "PreferencesOnly".
	ButtonLayout string `json:"buttonLayout" api:"nullable"`
	// When true, accept and reject buttons share the same visual weight (no
	// primary/secondary).
	EqualWeightButtons bool `json:"equalWeightButtons" api:"nullable"`
	// When true, swap accept/reject button order.
	FlipButtons bool `json:"flipButtons" api:"nullable"`
	// Modal layout style (box, bar, cloud, etc.).
	//
	// Any of "bar", "bar_inline", "box", "box_inline", "box_wide", "cloud",
	// "cloud_inline".
	Layout string `json:"layout" api:"nullable"`
	// Where the modal anchors on the viewport (e.g. bottom_left, top_right,
	// middle_center).
	//
	// Any of "bottom", "bottom_center", "bottom_left", "bottom_right", "middle",
	// "middle_center", "middle_left", "middle_right", "top", "top_center", "top_left",
	// "top_right".
	Position string `json:"position" api:"nullable"`
	// Show an X to dismiss the modal without choosing.
	ShowCloseIcon bool `json:"showCloseIcon" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonLayout       respjson.Field
		EqualWeightButtons respjson.Field
		FlipButtons        respjson.Field
		Layout             respjson.Field
		Position           respjson.Field
		ShowCloseIcon      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityDefaultGuiOptionsConsentModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingListResponseEntityDefaultGuiOptionsConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Global CSS variable overrides applied to both modals.
type ConsentSettingListResponseEntityDefaultGuiOptionsCssVariables struct {
	// CSS length, e.g. "4px" or "0.25rem".
	ButtonBorderRadius string `json:"buttonBorderRadius" api:"nullable"`
	// CSS color for the per-category block background in the preferences modal.
	CookieCategoryBlockBg string `json:"cookieCategoryBlockBg" api:"nullable"`
	// CSS color for the per-category block background on hover.
	CookieCategoryBlockHoverBg string `json:"cookieCategoryBlockHoverBg" api:"nullable"`
	// CSS color for the modal footer background.
	FooterBg string `json:"footerBg" api:"nullable"`
	// CSS color for footer text.
	FooterColor string `json:"footerColor" api:"nullable"`
	// CSS color for footer links.
	FooterLinkColor string `json:"footerLinkColor" api:"nullable"`
	// CSS color for footer links on hover.
	FooterLinkHoverColor string `json:"footerLinkHoverColor" api:"nullable"`
	// CSS color for the modal background.
	ModalBg string `json:"modalBg" api:"nullable"`
	// CSS length for modal corner radius.
	ModalBorderRadius string `json:"modalBorderRadius" api:"nullable"`
	// CSS color for the primary button background.
	PrimaryButtonBg string `json:"primaryButtonBg" api:"nullable"`
	// CSS color for primary button text.
	PrimaryButtonColor string `json:"primaryButtonColor" api:"nullable"`
	// CSS color for the primary button background on hover.
	PrimaryButtonHoverBg string `json:"primaryButtonHoverBg" api:"nullable"`
	// CSS color for primary button text on hover.
	PrimaryButtonHoverColor string `json:"primaryButtonHoverColor" api:"nullable"`
	// CSS color for primary text in the modal body.
	PrimaryTextColor string `json:"primaryTextColor" api:"nullable"`
	// CSS color for the secondary button background.
	SecondaryButtonBg string `json:"secondaryButtonBg" api:"nullable"`
	// CSS color for secondary button text.
	SecondaryButtonColor string `json:"secondaryButtonColor" api:"nullable"`
	// CSS color for secondary button background on hover.
	SecondaryButtonHoverBg string `json:"secondaryButtonHoverBg" api:"nullable"`
	// CSS color for secondary button text on hover.
	SecondaryButtonHoverColor string `json:"secondaryButtonHoverColor" api:"nullable"`
	// CSS color for secondary text in the modal body.
	SecondaryTextColor string `json:"secondaryTextColor" api:"nullable"`
	// CSS color for the "on" state of category toggles.
	ToggleOnBg string `json:"toggleOnBg" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonBorderRadius         respjson.Field
		CookieCategoryBlockBg      respjson.Field
		CookieCategoryBlockHoverBg respjson.Field
		FooterBg                   respjson.Field
		FooterColor                respjson.Field
		FooterLinkColor            respjson.Field
		FooterLinkHoverColor       respjson.Field
		ModalBg                    respjson.Field
		ModalBorderRadius          respjson.Field
		PrimaryButtonBg            respjson.Field
		PrimaryButtonColor         respjson.Field
		PrimaryButtonHoverBg       respjson.Field
		PrimaryButtonHoverColor    respjson.Field
		PrimaryTextColor           respjson.Field
		SecondaryButtonBg          respjson.Field
		SecondaryButtonColor       respjson.Field
		SecondaryButtonHoverBg     respjson.Field
		SecondaryButtonHoverColor  respjson.Field
		SecondaryTextColor         respjson.Field
		ToggleOnBg                 respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityDefaultGuiOptionsCssVariables) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingListResponseEntityDefaultGuiOptionsCssVariables) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the per-category preferences modal.
type ConsentSettingListResponseEntityDefaultGuiOptionsPreferencesModal struct {
	// Button layout for the preferences modal.
	//
	// Any of "AcceptAllRejectAllSave", "AcceptAllSave".
	ButtonLayout string `json:"buttonLayout" api:"nullable"`
	// When true, accept/reject share the same visual weight.
	EqualWeightButtons bool `json:"equalWeightButtons" api:"nullable"`
	// When true, swap accept/reject button order.
	FlipButtons bool `json:"flipButtons" api:"nullable"`
	// Preferences modal layout (box, bar, bar_wide).
	//
	// Any of "bar", "bar_wide", "box".
	Layout string `json:"layout" api:"nullable"`
	// Preferences modal anchor (left, right).
	//
	// Any of "left", "right".
	Position string `json:"position" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonLayout       respjson.Field
		EqualWeightButtons respjson.Field
		FlipButtons        respjson.Field
		Layout             respjson.Field
		Position           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityDefaultGuiOptionsPreferencesModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingListResponseEntityDefaultGuiOptionsPreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingListResponseEntityRegion struct {
	// Region this rule applies to. Use ISO 3166-1 alpha-2 country code ("US", "DE",
	// "BR") or country-subdivision code ("US-CA", "GB-ENG", "CA-ON"). Each region code
	// may appear in only one rule across `regions[]`.
	RegionCode string                                     `json:"regionCode" api:"required"`
	Rule       ConsentSettingListResponseEntityRegionRule `json:"rule" api:"required"`
	// Other region codes that should reuse this rule. Same code-format rules as
	// `regionCode`. Cannot include `regionCode` itself, cannot duplicate, cannot
	// overlap with another rule's regions.
	AdditionalRegions []string `json:"additionalRegions" api:"nullable"`
	// Optional label for this override group, e.g. "CCPA States". Shown in the
	// dashboard instead of the region codes. Pass null to clear it and fall back to
	// the codes.
	Name string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RegionCode        respjson.Field
		Rule              respjson.Field
		AdditionalRegions respjson.Field
		Name              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityRegion) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingListResponseEntityRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingListResponseEntityRegionRule struct {
	// Per-category default config for this rule. Every category defined in the
	// top-level `categories[].value` should have an entry here.
	Categories []ConsentSettingListResponseEntityRegionRuleCategory `json:"categories" api:"required"`
	// BCP 47 default language for this rule. Must have a matching entry in
	// `translations`. Examples: "en", "en-US", "es", "de".
	Language string `json:"language" api:"required"`
	// opt_in: scripts blocked until user accepts (GDPR style). opt_out: scripts run by
	// default until user rejects (CCPA style).
	//
	// Any of "opt_in", "opt_out".
	Mode string `json:"mode" api:"required"`
	// All UI copy, keyed by language. Must include an entry whose `language` matches
	// the rule's `language` field.
	Translations []ConsentSettingListResponseEntityRegionRuleTranslation `json:"translations" api:"required"`
	// When true, scripts not classified by services[] are blocked until the user opts
	// in.
	AutoblockUnknown bool `json:"autoblockUnknown" api:"nullable"`
	// When true, the consent modal auto-opens on page load.
	AutoShow bool `json:"autoShow" api:"nullable"`
	// Threshold config for autoShowDismissMode (page count or seconds).
	AutoShowDismissConfig ConsentSettingListResponseEntityRegionRuleAutoShowDismissConfig `json:"autoShowDismissConfig" api:"nullable"`
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	//
	// Any of "after_pages", "after_seconds", "never".
	AutoShowDismissMode string `json:"autoShowDismissMode" api:"nullable"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction bool `json:"disablePageInteraction" api:"nullable"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions ConsentSettingListResponseEntityRegionRuleGuiOptions `json:"guiOptions" api:"nullable"`
	// When true, the modal is suppressed for known bot user agents.
	HideFromBots bool `json:"hideFromBots" api:"nullable"`
	// When true, the per-service list (services[]) is rendered inside the preferences
	// modal.
	ShowVendorsInPreferences bool `json:"showVendorsInPreferences" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories               respjson.Field
		Language                 respjson.Field
		Mode                     respjson.Field
		Translations             respjson.Field
		AutoblockUnknown         respjson.Field
		AutoShow                 respjson.Field
		AutoShowDismissConfig    respjson.Field
		AutoShowDismissMode      respjson.Field
		DisablePageInteraction   respjson.Field
		GuiOptions               respjson.Field
		HideFromBots             respjson.Field
		ShowVendorsInPreferences respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityRegionRule) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingListResponseEntityRegionRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingListResponseEntityRegionRuleCategory struct {
	// Category value (matches `categories[].value`) this entry configures.
	Key   string                                                  `json:"key" api:"required"`
	Value ConsentSettingListResponseEntityRegionRuleCategoryValue `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityRegionRuleCategory) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingListResponseEntityRegionRuleCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingListResponseEntityRegionRuleCategoryValue struct {
	// Whether this category is on by default before the user interacts.
	Enabled bool `json:"enabled" api:"required"`
	// When true, this category defaults off if the browser sends Sec-GPC: 1.
	AutoDisableOnGpc bool `json:"autoDisableOnGPC" api:"nullable"`
	// When true, the user cannot toggle this category in the preferences modal.
	ReadOnly bool `json:"readOnly" api:"nullable"`
	// When true, the page reloads after this category is toggled so newly-allowed
	// scripts can run.
	ReloadPage bool `json:"reloadPage" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled          respjson.Field
		AutoDisableOnGpc respjson.Field
		ReadOnly         respjson.Field
		ReloadPage       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityRegionRuleCategoryValue) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingListResponseEntityRegionRuleCategoryValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingListResponseEntityRegionRuleTranslation struct {
	// BCP 47 language tag identifying which translation this entry provides. Examples:
	// "en", "en-US", "es", "fr-CA". The default rule's `language` must appear here.
	Language string                                                     `json:"language" api:"required"`
	Value    ConsentSettingListResponseEntityRegionRuleTranslationValue `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Language    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityRegionRuleTranslation) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingListResponseEntityRegionRuleTranslation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingListResponseEntityRegionRuleTranslationValue struct {
	// Translated copy for the initial consent modal.
	ConsentModal ConsentSettingListResponseEntityRegionRuleTranslationValueConsentModal `json:"consentModal" api:"nullable"`
	// Translated copy for the preferences modal.
	PreferencesModal ConsentSettingListResponseEntityRegionRuleTranslationValuePreferencesModal `json:"preferencesModal" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConsentModal     respjson.Field
		PreferencesModal respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityRegionRuleTranslationValue) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingListResponseEntityRegionRuleTranslationValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the initial consent modal.
type ConsentSettingListResponseEntityRegionRuleTranslationValueConsentModal struct {
	// Label for the "accept all" button.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Body text shown in the consent modal.
	Description string `json:"description" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the button that opens the preferences modal.
	ShowPreferencesBtn string `json:"showPreferencesBtn" api:"required"`
	// Heading shown at the top of the consent modal.
	Title string `json:"title" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel string `json:"closeIconLabel" api:"nullable"`
	// Optional footer HTML/text shown below the body.
	Footer string `json:"footer" api:"nullable"`
	// Optional text shown when the browser is sending Sec-GPC: 1.
	GpcNotification string `json:"gpcNotification" api:"nullable"`
	// Link label for the privacy policy (defaults to "Privacy Policy").
	PrivacyPolicyLabel string `json:"privacyPolicyLabel" api:"nullable"`
	// Privacy policy link. Accepts an absolute URL (https://...) or a relative path
	// (/legal/privacy) so one config can run across multiple domains. javascript:,
	// data:, and vbscript: schemes are rejected.
	PrivacyPolicyURL string `json:"privacyPolicyUrl" api:"nullable"`
	// Optional message shown when the consent revision has bumped (re-prompt copy).
	RevisionMessage string `json:"revisionMessage" api:"nullable"`
	// Link label for terms of service (defaults to "Terms of Service").
	TermsOfServiceLabel string `json:"termsOfServiceLabel" api:"nullable"`
	// Terms of service link. Same scheme rules as privacyPolicyUrl: absolute or
	// relative path only.
	TermsOfServiceURL string `json:"termsOfServiceUrl" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptAllBtn        respjson.Field
		Description         respjson.Field
		RejectAllBtn        respjson.Field
		ShowPreferencesBtn  respjson.Field
		Title               respjson.Field
		CloseIconLabel      respjson.Field
		Footer              respjson.Field
		GpcNotification     respjson.Field
		PrivacyPolicyLabel  respjson.Field
		PrivacyPolicyURL    respjson.Field
		RevisionMessage     respjson.Field
		TermsOfServiceLabel respjson.Field
		TermsOfServiceURL   respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityRegionRuleTranslationValueConsentModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingListResponseEntityRegionRuleTranslationValueConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the preferences modal.
type ConsentSettingListResponseEntityRegionRuleTranslationValuePreferencesModal struct {
	// Label for the "accept all" button in the preferences modal.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel string `json:"closeIconLabel" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the "save my choices" button.
	SavePreferencesBtn string `json:"savePreferencesBtn" api:"required"`
	// Per-category sections plus a leading general section. The server backfills
	// sections for each defined category if absent.
	Sections []ConsentSettingListResponseEntityRegionRuleTranslationValuePreferencesModalSection `json:"sections" api:"required"`
	// Heading for the preferences modal.
	Title string `json:"title" api:"required"`
	// Optional collapsible info block rendered after the category sections. Renders
	// only when both title and body are non-empty. Use it for long-form text unrelated
	// to any specific category (e.g. a third-party advertisers disclosure).
	AdditionalInfo ConsentSettingListResponseEntityRegionRuleTranslationValuePreferencesModalAdditionalInfo `json:"additionalInfo" api:"nullable"`
	// Pluralizable label for the "N services" counter (e.g. "Service|Services").
	ServiceCounterLabel string `json:"serviceCounterLabel" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptAllBtn        respjson.Field
		CloseIconLabel      respjson.Field
		RejectAllBtn        respjson.Field
		SavePreferencesBtn  respjson.Field
		Sections            respjson.Field
		Title               respjson.Field
		AdditionalInfo      respjson.Field
		ServiceCounterLabel respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityRegionRuleTranslationValuePreferencesModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingListResponseEntityRegionRuleTranslationValuePreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingListResponseEntityRegionRuleTranslationValuePreferencesModalSection struct {
	// Body text for this preferences-modal section.
	Description string `json:"description" api:"required"`
	// Heading for this preferences-modal section.
	Title string `json:"title" api:"required"`
	// Optional disclosure table listing cookies/services in this section.
	CookieTable ConsentSettingListResponseEntityRegionRuleTranslationValuePreferencesModalSectionCookieTable `json:"cookieTable" api:"nullable"`
	// Category value this section configures. Use null/omit for the introductory
	// section that applies to no specific category.
	LinkedCategory string `json:"linkedCategory" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description    respjson.Field
		Title          respjson.Field
		CookieTable    respjson.Field
		LinkedCategory respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityRegionRuleTranslationValuePreferencesModalSection) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingListResponseEntityRegionRuleTranslationValuePreferencesModalSection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional disclosure table listing cookies/services in this section.
type ConsentSettingListResponseEntityRegionRuleTranslationValuePreferencesModalSectionCookieTable struct {
	// Rows of the cookie table — each row is a {key,value} pair (e.g. cookie name →
	// purpose).
	Body []ConsentSettingListResponseEntityRegionRuleTranslationValuePreferencesModalSectionCookieTableBody `json:"body" api:"required"`
	// Header cells. Same {key,value} shape as rows.
	Headers []ConsentSettingListResponseEntityRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader `json:"headers" api:"required"`
	// Optional table caption.
	Caption string `json:"caption" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Headers     respjson.Field
		Caption     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityRegionRuleTranslationValuePreferencesModalSectionCookieTable) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingListResponseEntityRegionRuleTranslationValuePreferencesModalSectionCookieTable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingListResponseEntityRegionRuleTranslationValuePreferencesModalSectionCookieTableBody struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityRegionRuleTranslationValuePreferencesModalSectionCookieTableBody) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingListResponseEntityRegionRuleTranslationValuePreferencesModalSectionCookieTableBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingListResponseEntityRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingListResponseEntityRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional collapsible info block rendered after the category sections. Renders
// only when both title and body are non-empty. Use it for long-form text unrelated
// to any specific category (e.g. a third-party advertisers disclosure).
type ConsentSettingListResponseEntityRegionRuleTranslationValuePreferencesModalAdditionalInfo struct {
	// Body text for the additional-info block. Supports the rich-text markdown subset
	// used by section descriptions: line breaks, [text](url) links (https:// or
	// relative paths), and **bold**. HTML is stripped.
	Body string `json:"body" api:"required"`
	// Heading for the additional-info block.
	Title string `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityRegionRuleTranslationValuePreferencesModalAdditionalInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingListResponseEntityRegionRuleTranslationValuePreferencesModalAdditionalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Threshold config for autoShowDismissMode (page count or seconds).
type ConsentSettingListResponseEntityRegionRuleAutoShowDismissConfig struct {
	// Pages-viewed threshold when autoShowDismissMode = after_pages.
	PageCount int64 `json:"pageCount" api:"nullable"`
	// Seconds-on-page threshold when autoShowDismissMode = after_seconds.
	Seconds int64 `json:"seconds" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageCount   respjson.Field
		Seconds     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityRegionRuleAutoShowDismissConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingListResponseEntityRegionRuleAutoShowDismissConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the modals (layout/position/colors).
type ConsentSettingListResponseEntityRegionRuleGuiOptions struct {
	// Visual options for the initial consent modal.
	ConsentModal ConsentSettingListResponseEntityRegionRuleGuiOptionsConsentModal `json:"consentModal" api:"nullable"`
	// Global CSS variable overrides applied to both modals.
	CssVariables ConsentSettingListResponseEntityRegionRuleGuiOptionsCssVariables `json:"cssVariables" api:"nullable"`
	// Visual options for the per-category preferences modal.
	PreferencesModal ConsentSettingListResponseEntityRegionRuleGuiOptionsPreferencesModal `json:"preferencesModal" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConsentModal     respjson.Field
		CssVariables     respjson.Field
		PreferencesModal respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityRegionRuleGuiOptions) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingListResponseEntityRegionRuleGuiOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the initial consent modal.
type ConsentSettingListResponseEntityRegionRuleGuiOptionsConsentModal struct {
	// Which buttons to render in the consent modal and in what arrangement.
	//
	// Any of "AcceptAllNecessaryPreferences", "AcceptAllPreferences", "AcceptOnly",
	// "InformationOnly", "PreferencesOnly".
	ButtonLayout string `json:"buttonLayout" api:"nullable"`
	// When true, accept and reject buttons share the same visual weight (no
	// primary/secondary).
	EqualWeightButtons bool `json:"equalWeightButtons" api:"nullable"`
	// When true, swap accept/reject button order.
	FlipButtons bool `json:"flipButtons" api:"nullable"`
	// Modal layout style (box, bar, cloud, etc.).
	//
	// Any of "bar", "bar_inline", "box", "box_inline", "box_wide", "cloud",
	// "cloud_inline".
	Layout string `json:"layout" api:"nullable"`
	// Where the modal anchors on the viewport (e.g. bottom_left, top_right,
	// middle_center).
	//
	// Any of "bottom", "bottom_center", "bottom_left", "bottom_right", "middle",
	// "middle_center", "middle_left", "middle_right", "top", "top_center", "top_left",
	// "top_right".
	Position string `json:"position" api:"nullable"`
	// Show an X to dismiss the modal without choosing.
	ShowCloseIcon bool `json:"showCloseIcon" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonLayout       respjson.Field
		EqualWeightButtons respjson.Field
		FlipButtons        respjson.Field
		Layout             respjson.Field
		Position           respjson.Field
		ShowCloseIcon      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityRegionRuleGuiOptionsConsentModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingListResponseEntityRegionRuleGuiOptionsConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Global CSS variable overrides applied to both modals.
type ConsentSettingListResponseEntityRegionRuleGuiOptionsCssVariables struct {
	// CSS length, e.g. "4px" or "0.25rem".
	ButtonBorderRadius string `json:"buttonBorderRadius" api:"nullable"`
	// CSS color for the per-category block background in the preferences modal.
	CookieCategoryBlockBg string `json:"cookieCategoryBlockBg" api:"nullable"`
	// CSS color for the per-category block background on hover.
	CookieCategoryBlockHoverBg string `json:"cookieCategoryBlockHoverBg" api:"nullable"`
	// CSS color for the modal footer background.
	FooterBg string `json:"footerBg" api:"nullable"`
	// CSS color for footer text.
	FooterColor string `json:"footerColor" api:"nullable"`
	// CSS color for footer links.
	FooterLinkColor string `json:"footerLinkColor" api:"nullable"`
	// CSS color for footer links on hover.
	FooterLinkHoverColor string `json:"footerLinkHoverColor" api:"nullable"`
	// CSS color for the modal background.
	ModalBg string `json:"modalBg" api:"nullable"`
	// CSS length for modal corner radius.
	ModalBorderRadius string `json:"modalBorderRadius" api:"nullable"`
	// CSS color for the primary button background.
	PrimaryButtonBg string `json:"primaryButtonBg" api:"nullable"`
	// CSS color for primary button text.
	PrimaryButtonColor string `json:"primaryButtonColor" api:"nullable"`
	// CSS color for the primary button background on hover.
	PrimaryButtonHoverBg string `json:"primaryButtonHoverBg" api:"nullable"`
	// CSS color for primary button text on hover.
	PrimaryButtonHoverColor string `json:"primaryButtonHoverColor" api:"nullable"`
	// CSS color for primary text in the modal body.
	PrimaryTextColor string `json:"primaryTextColor" api:"nullable"`
	// CSS color for the secondary button background.
	SecondaryButtonBg string `json:"secondaryButtonBg" api:"nullable"`
	// CSS color for secondary button text.
	SecondaryButtonColor string `json:"secondaryButtonColor" api:"nullable"`
	// CSS color for secondary button background on hover.
	SecondaryButtonHoverBg string `json:"secondaryButtonHoverBg" api:"nullable"`
	// CSS color for secondary button text on hover.
	SecondaryButtonHoverColor string `json:"secondaryButtonHoverColor" api:"nullable"`
	// CSS color for secondary text in the modal body.
	SecondaryTextColor string `json:"secondaryTextColor" api:"nullable"`
	// CSS color for the "on" state of category toggles.
	ToggleOnBg string `json:"toggleOnBg" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonBorderRadius         respjson.Field
		CookieCategoryBlockBg      respjson.Field
		CookieCategoryBlockHoverBg respjson.Field
		FooterBg                   respjson.Field
		FooterColor                respjson.Field
		FooterLinkColor            respjson.Field
		FooterLinkHoverColor       respjson.Field
		ModalBg                    respjson.Field
		ModalBorderRadius          respjson.Field
		PrimaryButtonBg            respjson.Field
		PrimaryButtonColor         respjson.Field
		PrimaryButtonHoverBg       respjson.Field
		PrimaryButtonHoverColor    respjson.Field
		PrimaryTextColor           respjson.Field
		SecondaryButtonBg          respjson.Field
		SecondaryButtonColor       respjson.Field
		SecondaryButtonHoverBg     respjson.Field
		SecondaryButtonHoverColor  respjson.Field
		SecondaryTextColor         respjson.Field
		ToggleOnBg                 respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityRegionRuleGuiOptionsCssVariables) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingListResponseEntityRegionRuleGuiOptionsCssVariables) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the per-category preferences modal.
type ConsentSettingListResponseEntityRegionRuleGuiOptionsPreferencesModal struct {
	// Button layout for the preferences modal.
	//
	// Any of "AcceptAllRejectAllSave", "AcceptAllSave".
	ButtonLayout string `json:"buttonLayout" api:"nullable"`
	// When true, accept/reject share the same visual weight.
	EqualWeightButtons bool `json:"equalWeightButtons" api:"nullable"`
	// When true, swap accept/reject button order.
	FlipButtons bool `json:"flipButtons" api:"nullable"`
	// Preferences modal layout (box, bar, bar_wide).
	//
	// Any of "bar", "bar_wide", "box".
	Layout string `json:"layout" api:"nullable"`
	// Preferences modal anchor (left, right).
	//
	// Any of "left", "right".
	Position string `json:"position" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonLayout       respjson.Field
		EqualWeightButtons respjson.Field
		FlipButtons        respjson.Field
		Layout             respjson.Field
		Position           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityRegionRuleGuiOptionsPreferencesModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingListResponseEntityRegionRuleGuiOptionsPreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingListResponseEntityService struct {
	// Internal notes shown to admins in the dashboard. Not user-facing.
	InternalNotes string `json:"internalNotes" api:"required"`
	// Display name for this service in the preferences modal.
	Label string `json:"label" api:"required"`
	// Extra category values this service belongs to. Each must match a
	// `categories[].value`.
	AdditionalCategories []string `json:"additionalCategories" api:"nullable"`
	// Primary category value this service belongs to. Must match one of the top-level
	// `categories[].value` entries.
	Category string `json:"category" api:"nullable"`
	// Domains/paths this service matches. Patterns matching the CMP's own scripts
	// (e.g. cdn.oursprivacy.com/cmp-init) are rejected to prevent the CMP blocking
	// itself — use a more specific path like cdn.oursprivacy.com/main.js to block a
	// specific script.
	DomainPatterns []string `json:"domainPatterns" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InternalNotes        respjson.Field
		Label                respjson.Field
		AdditionalCategories respjson.Field
		Category             respjson.Field
		DomainPatterns       respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponseEntityService) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingListResponseEntityService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingNewResponse struct {
	// Server-assigned UUID for this consent settings record.
	ID string `json:"id" api:"required"`
	// Top-level consent categories (e.g. necessary / analytics / advertising). Server
	// re-stamps `priority` to 0..N on write.
	Categories []ConsentSettingNewResponseCategory `json:"categories" api:"required"`
	// ISO 8601 timestamp when the record was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Default rule used when the user is not in any region listed in `regions[]`.
	Default ConsentSettingNewResponseDefault `json:"default" api:"required"`
	// Discriminator for the entity type. Always "consentSettings".
	Kind string `json:"kind" api:"required"`
	// Human-readable name shown in the dashboard.
	Name string `json:"name" api:"required"`
	// Per-region rule overrides. The first rule whose `regionCode`/`additionalRegions`
	// includes the user's region wins; otherwise `default` applies.
	Regions []ConsentSettingNewResponseRegion `json:"regions" api:"required"`
	// Per-service entries powering "show vendors" and category-aware blocking.
	Services []ConsentSettingNewResponseService `json:"services" api:"required"`
	// Enabled serves the CMP on its allowed domains and requires at least one.
	// Disabled does not serve the CMP and may have no allowed domains.
	//
	// Any of "Disabled", "Enabled".
	Status ConsentSettingNewResponseStatus `json:"status" api:"required"`
	// Name of the cookie that stores the user's consent state. Defaults to
	// "op_consent".
	ConsentCookieName string `json:"consentCookieName" api:"nullable"`
	// Optional custom CDN domain for serving the CMP script (e.g.
	// consent.example.com).
	CustomDomain string `json:"customDomain" api:"nullable"`
	// Cookie/localStorage key for the visitor (device) ID. Must exactly match the web
	// SDK's cookie_names.device_id (default "ours_device_id") so consent and SDK
	// events resolve to the same visitor.
	DeviceIDCookieName string `json:"deviceIdCookieName" api:"nullable"`
	// Revision counter. Bump this to force users who already consented to see the
	// modal again (the SDK compares the persisted revision against this value).
	Revision float64 `json:"revision" api:"nullable"`
	// CSS class names that opt scripts out of consent blocking. Each entry must be a
	// single class token (no whitespace).
	SkipBlockingClassNames []string `json:"skipBlockingClassNames" api:"nullable"`
	// ISO 8601 timestamp of the last write. Null on a freshly created record.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// Pixel of the WebSource that this CMP is wired into. Setting this to a token that
	// is not a valid WebSource of yours is rejected; use null to clear the link.
	WebSDKToken string `json:"webSDKToken" api:"nullable"`
	// Domains where this CMP may run. An Enabled CMP requires at least one allowed
	// domain; with none configured, the CMP does not run. Matching domains may share
	// consent across subdomains.
	WhitelistDomains []string `json:"whitelistDomains" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		Categories             respjson.Field
		CreatedAt              respjson.Field
		Default                respjson.Field
		Kind                   respjson.Field
		Name                   respjson.Field
		Regions                respjson.Field
		Services               respjson.Field
		Status                 respjson.Field
		ConsentCookieName      respjson.Field
		CustomDomain           respjson.Field
		DeviceIDCookieName     respjson.Field
		Revision               respjson.Field
		SkipBlockingClassNames respjson.Field
		UpdatedAt              respjson.Field
		WebSDKToken            respjson.Field
		WhitelistDomains       respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingNewResponseCategory struct {
	// Human-readable label shown next to the toggle in the preferences modal.
	Label string `json:"label" api:"required"`
	// Sort key. Lower numbers render first. Server re-stamps to 0..N on write — send
	// any integer, gaps and duplicates are ironed out.
	Priority int64 `json:"priority" api:"required"`
	// Stable identifier referenced by services and translation sections.
	// Conventionally lowercase (e.g. "necessary", "analytics", "advertising").
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		Priority    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseCategory) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingNewResponseCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default rule used when the user is not in any region listed in `regions[]`.
type ConsentSettingNewResponseDefault struct {
	// Per-category default config for this rule. Every category defined in the
	// top-level `categories[].value` should have an entry here.
	Categories []ConsentSettingNewResponseDefaultCategory `json:"categories" api:"required"`
	// BCP 47 default language for this rule. Must have a matching entry in
	// `translations`. Examples: "en", "en-US", "es", "de".
	Language string `json:"language" api:"required"`
	// opt_in: scripts blocked until user accepts (GDPR style). opt_out: scripts run by
	// default until user rejects (CCPA style).
	//
	// Any of "opt_in", "opt_out".
	Mode string `json:"mode" api:"required"`
	// All UI copy, keyed by language. Must include an entry whose `language` matches
	// the rule's `language` field.
	Translations []ConsentSettingNewResponseDefaultTranslation `json:"translations" api:"required"`
	// When true, scripts not classified by services[] are blocked until the user opts
	// in.
	AutoblockUnknown bool `json:"autoblockUnknown" api:"nullable"`
	// When true, the consent modal auto-opens on page load.
	AutoShow bool `json:"autoShow" api:"nullable"`
	// Threshold config for autoShowDismissMode (page count or seconds).
	AutoShowDismissConfig ConsentSettingNewResponseDefaultAutoShowDismissConfig `json:"autoShowDismissConfig" api:"nullable"`
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	//
	// Any of "after_pages", "after_seconds", "never".
	AutoShowDismissMode string `json:"autoShowDismissMode" api:"nullable"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction bool `json:"disablePageInteraction" api:"nullable"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions ConsentSettingNewResponseDefaultGuiOptions `json:"guiOptions" api:"nullable"`
	// When true, the modal is suppressed for known bot user agents.
	HideFromBots bool `json:"hideFromBots" api:"nullable"`
	// When true, the per-service list (services[]) is rendered inside the preferences
	// modal.
	ShowVendorsInPreferences bool `json:"showVendorsInPreferences" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories               respjson.Field
		Language                 respjson.Field
		Mode                     respjson.Field
		Translations             respjson.Field
		AutoblockUnknown         respjson.Field
		AutoShow                 respjson.Field
		AutoShowDismissConfig    respjson.Field
		AutoShowDismissMode      respjson.Field
		DisablePageInteraction   respjson.Field
		GuiOptions               respjson.Field
		HideFromBots             respjson.Field
		ShowVendorsInPreferences respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseDefault) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingNewResponseDefault) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingNewResponseDefaultCategory struct {
	// Category value (matches `categories[].value`) this entry configures.
	Key   string                                        `json:"key" api:"required"`
	Value ConsentSettingNewResponseDefaultCategoryValue `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseDefaultCategory) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingNewResponseDefaultCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingNewResponseDefaultCategoryValue struct {
	// Whether this category is on by default before the user interacts.
	Enabled bool `json:"enabled" api:"required"`
	// When true, this category defaults off if the browser sends Sec-GPC: 1.
	AutoDisableOnGpc bool `json:"autoDisableOnGPC" api:"nullable"`
	// When true, the user cannot toggle this category in the preferences modal.
	ReadOnly bool `json:"readOnly" api:"nullable"`
	// When true, the page reloads after this category is toggled so newly-allowed
	// scripts can run.
	ReloadPage bool `json:"reloadPage" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled          respjson.Field
		AutoDisableOnGpc respjson.Field
		ReadOnly         respjson.Field
		ReloadPage       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseDefaultCategoryValue) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingNewResponseDefaultCategoryValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingNewResponseDefaultTranslation struct {
	// BCP 47 language tag identifying which translation this entry provides. Examples:
	// "en", "en-US", "es", "fr-CA". The default rule's `language` must appear here.
	Language string                                           `json:"language" api:"required"`
	Value    ConsentSettingNewResponseDefaultTranslationValue `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Language    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseDefaultTranslation) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingNewResponseDefaultTranslation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingNewResponseDefaultTranslationValue struct {
	// Translated copy for the initial consent modal.
	ConsentModal ConsentSettingNewResponseDefaultTranslationValueConsentModal `json:"consentModal" api:"nullable"`
	// Translated copy for the preferences modal.
	PreferencesModal ConsentSettingNewResponseDefaultTranslationValuePreferencesModal `json:"preferencesModal" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConsentModal     respjson.Field
		PreferencesModal respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseDefaultTranslationValue) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingNewResponseDefaultTranslationValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the initial consent modal.
type ConsentSettingNewResponseDefaultTranslationValueConsentModal struct {
	// Label for the "accept all" button.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Body text shown in the consent modal.
	Description string `json:"description" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the button that opens the preferences modal.
	ShowPreferencesBtn string `json:"showPreferencesBtn" api:"required"`
	// Heading shown at the top of the consent modal.
	Title string `json:"title" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel string `json:"closeIconLabel" api:"nullable"`
	// Optional footer HTML/text shown below the body.
	Footer string `json:"footer" api:"nullable"`
	// Optional text shown when the browser is sending Sec-GPC: 1.
	GpcNotification string `json:"gpcNotification" api:"nullable"`
	// Link label for the privacy policy (defaults to "Privacy Policy").
	PrivacyPolicyLabel string `json:"privacyPolicyLabel" api:"nullable"`
	// Privacy policy link. Accepts an absolute URL (https://...) or a relative path
	// (/legal/privacy) so one config can run across multiple domains. javascript:,
	// data:, and vbscript: schemes are rejected.
	PrivacyPolicyURL string `json:"privacyPolicyUrl" api:"nullable"`
	// Optional message shown when the consent revision has bumped (re-prompt copy).
	RevisionMessage string `json:"revisionMessage" api:"nullable"`
	// Link label for terms of service (defaults to "Terms of Service").
	TermsOfServiceLabel string `json:"termsOfServiceLabel" api:"nullable"`
	// Terms of service link. Same scheme rules as privacyPolicyUrl: absolute or
	// relative path only.
	TermsOfServiceURL string `json:"termsOfServiceUrl" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptAllBtn        respjson.Field
		Description         respjson.Field
		RejectAllBtn        respjson.Field
		ShowPreferencesBtn  respjson.Field
		Title               respjson.Field
		CloseIconLabel      respjson.Field
		Footer              respjson.Field
		GpcNotification     respjson.Field
		PrivacyPolicyLabel  respjson.Field
		PrivacyPolicyURL    respjson.Field
		RevisionMessage     respjson.Field
		TermsOfServiceLabel respjson.Field
		TermsOfServiceURL   respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseDefaultTranslationValueConsentModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingNewResponseDefaultTranslationValueConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the preferences modal.
type ConsentSettingNewResponseDefaultTranslationValuePreferencesModal struct {
	// Label for the "accept all" button in the preferences modal.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel string `json:"closeIconLabel" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the "save my choices" button.
	SavePreferencesBtn string `json:"savePreferencesBtn" api:"required"`
	// Per-category sections plus a leading general section. The server backfills
	// sections for each defined category if absent.
	Sections []ConsentSettingNewResponseDefaultTranslationValuePreferencesModalSection `json:"sections" api:"required"`
	// Heading for the preferences modal.
	Title string `json:"title" api:"required"`
	// Optional collapsible info block rendered after the category sections. Renders
	// only when both title and body are non-empty. Use it for long-form text unrelated
	// to any specific category (e.g. a third-party advertisers disclosure).
	AdditionalInfo ConsentSettingNewResponseDefaultTranslationValuePreferencesModalAdditionalInfo `json:"additionalInfo" api:"nullable"`
	// Pluralizable label for the "N services" counter (e.g. "Service|Services").
	ServiceCounterLabel string `json:"serviceCounterLabel" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptAllBtn        respjson.Field
		CloseIconLabel      respjson.Field
		RejectAllBtn        respjson.Field
		SavePreferencesBtn  respjson.Field
		Sections            respjson.Field
		Title               respjson.Field
		AdditionalInfo      respjson.Field
		ServiceCounterLabel respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseDefaultTranslationValuePreferencesModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingNewResponseDefaultTranslationValuePreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingNewResponseDefaultTranslationValuePreferencesModalSection struct {
	// Body text for this preferences-modal section.
	Description string `json:"description" api:"required"`
	// Heading for this preferences-modal section.
	Title string `json:"title" api:"required"`
	// Optional disclosure table listing cookies/services in this section.
	CookieTable ConsentSettingNewResponseDefaultTranslationValuePreferencesModalSectionCookieTable `json:"cookieTable" api:"nullable"`
	// Category value this section configures. Use null/omit for the introductory
	// section that applies to no specific category.
	LinkedCategory string `json:"linkedCategory" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description    respjson.Field
		Title          respjson.Field
		CookieTable    respjson.Field
		LinkedCategory respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseDefaultTranslationValuePreferencesModalSection) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingNewResponseDefaultTranslationValuePreferencesModalSection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional disclosure table listing cookies/services in this section.
type ConsentSettingNewResponseDefaultTranslationValuePreferencesModalSectionCookieTable struct {
	// Rows of the cookie table — each row is a {key,value} pair (e.g. cookie name →
	// purpose).
	Body []ConsentSettingNewResponseDefaultTranslationValuePreferencesModalSectionCookieTableBody `json:"body" api:"required"`
	// Header cells. Same {key,value} shape as rows.
	Headers []ConsentSettingNewResponseDefaultTranslationValuePreferencesModalSectionCookieTableHeader `json:"headers" api:"required"`
	// Optional table caption.
	Caption string `json:"caption" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Headers     respjson.Field
		Caption     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseDefaultTranslationValuePreferencesModalSectionCookieTable) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingNewResponseDefaultTranslationValuePreferencesModalSectionCookieTable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingNewResponseDefaultTranslationValuePreferencesModalSectionCookieTableBody struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseDefaultTranslationValuePreferencesModalSectionCookieTableBody) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingNewResponseDefaultTranslationValuePreferencesModalSectionCookieTableBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingNewResponseDefaultTranslationValuePreferencesModalSectionCookieTableHeader struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseDefaultTranslationValuePreferencesModalSectionCookieTableHeader) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingNewResponseDefaultTranslationValuePreferencesModalSectionCookieTableHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional collapsible info block rendered after the category sections. Renders
// only when both title and body are non-empty. Use it for long-form text unrelated
// to any specific category (e.g. a third-party advertisers disclosure).
type ConsentSettingNewResponseDefaultTranslationValuePreferencesModalAdditionalInfo struct {
	// Body text for the additional-info block. Supports the rich-text markdown subset
	// used by section descriptions: line breaks, [text](url) links (https:// or
	// relative paths), and **bold**. HTML is stripped.
	Body string `json:"body" api:"required"`
	// Heading for the additional-info block.
	Title string `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseDefaultTranslationValuePreferencesModalAdditionalInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingNewResponseDefaultTranslationValuePreferencesModalAdditionalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Threshold config for autoShowDismissMode (page count or seconds).
type ConsentSettingNewResponseDefaultAutoShowDismissConfig struct {
	// Pages-viewed threshold when autoShowDismissMode = after_pages.
	PageCount int64 `json:"pageCount" api:"nullable"`
	// Seconds-on-page threshold when autoShowDismissMode = after_seconds.
	Seconds int64 `json:"seconds" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageCount   respjson.Field
		Seconds     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseDefaultAutoShowDismissConfig) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingNewResponseDefaultAutoShowDismissConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the modals (layout/position/colors).
type ConsentSettingNewResponseDefaultGuiOptions struct {
	// Visual options for the initial consent modal.
	ConsentModal ConsentSettingNewResponseDefaultGuiOptionsConsentModal `json:"consentModal" api:"nullable"`
	// Global CSS variable overrides applied to both modals.
	CssVariables ConsentSettingNewResponseDefaultGuiOptionsCssVariables `json:"cssVariables" api:"nullable"`
	// Visual options for the per-category preferences modal.
	PreferencesModal ConsentSettingNewResponseDefaultGuiOptionsPreferencesModal `json:"preferencesModal" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConsentModal     respjson.Field
		CssVariables     respjson.Field
		PreferencesModal respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseDefaultGuiOptions) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingNewResponseDefaultGuiOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the initial consent modal.
type ConsentSettingNewResponseDefaultGuiOptionsConsentModal struct {
	// Which buttons to render in the consent modal and in what arrangement.
	//
	// Any of "AcceptAllNecessaryPreferences", "AcceptAllPreferences", "AcceptOnly",
	// "InformationOnly", "PreferencesOnly".
	ButtonLayout string `json:"buttonLayout" api:"nullable"`
	// When true, accept and reject buttons share the same visual weight (no
	// primary/secondary).
	EqualWeightButtons bool `json:"equalWeightButtons" api:"nullable"`
	// When true, swap accept/reject button order.
	FlipButtons bool `json:"flipButtons" api:"nullable"`
	// Modal layout style (box, bar, cloud, etc.).
	//
	// Any of "bar", "bar_inline", "box", "box_inline", "box_wide", "cloud",
	// "cloud_inline".
	Layout string `json:"layout" api:"nullable"`
	// Where the modal anchors on the viewport (e.g. bottom_left, top_right,
	// middle_center).
	//
	// Any of "bottom", "bottom_center", "bottom_left", "bottom_right", "middle",
	// "middle_center", "middle_left", "middle_right", "top", "top_center", "top_left",
	// "top_right".
	Position string `json:"position" api:"nullable"`
	// Show an X to dismiss the modal without choosing.
	ShowCloseIcon bool `json:"showCloseIcon" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonLayout       respjson.Field
		EqualWeightButtons respjson.Field
		FlipButtons        respjson.Field
		Layout             respjson.Field
		Position           respjson.Field
		ShowCloseIcon      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseDefaultGuiOptionsConsentModal) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingNewResponseDefaultGuiOptionsConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Global CSS variable overrides applied to both modals.
type ConsentSettingNewResponseDefaultGuiOptionsCssVariables struct {
	// CSS length, e.g. "4px" or "0.25rem".
	ButtonBorderRadius string `json:"buttonBorderRadius" api:"nullable"`
	// CSS color for the per-category block background in the preferences modal.
	CookieCategoryBlockBg string `json:"cookieCategoryBlockBg" api:"nullable"`
	// CSS color for the per-category block background on hover.
	CookieCategoryBlockHoverBg string `json:"cookieCategoryBlockHoverBg" api:"nullable"`
	// CSS color for the modal footer background.
	FooterBg string `json:"footerBg" api:"nullable"`
	// CSS color for footer text.
	FooterColor string `json:"footerColor" api:"nullable"`
	// CSS color for footer links.
	FooterLinkColor string `json:"footerLinkColor" api:"nullable"`
	// CSS color for footer links on hover.
	FooterLinkHoverColor string `json:"footerLinkHoverColor" api:"nullable"`
	// CSS color for the modal background.
	ModalBg string `json:"modalBg" api:"nullable"`
	// CSS length for modal corner radius.
	ModalBorderRadius string `json:"modalBorderRadius" api:"nullable"`
	// CSS color for the primary button background.
	PrimaryButtonBg string `json:"primaryButtonBg" api:"nullable"`
	// CSS color for primary button text.
	PrimaryButtonColor string `json:"primaryButtonColor" api:"nullable"`
	// CSS color for the primary button background on hover.
	PrimaryButtonHoverBg string `json:"primaryButtonHoverBg" api:"nullable"`
	// CSS color for primary button text on hover.
	PrimaryButtonHoverColor string `json:"primaryButtonHoverColor" api:"nullable"`
	// CSS color for primary text in the modal body.
	PrimaryTextColor string `json:"primaryTextColor" api:"nullable"`
	// CSS color for the secondary button background.
	SecondaryButtonBg string `json:"secondaryButtonBg" api:"nullable"`
	// CSS color for secondary button text.
	SecondaryButtonColor string `json:"secondaryButtonColor" api:"nullable"`
	// CSS color for secondary button background on hover.
	SecondaryButtonHoverBg string `json:"secondaryButtonHoverBg" api:"nullable"`
	// CSS color for secondary button text on hover.
	SecondaryButtonHoverColor string `json:"secondaryButtonHoverColor" api:"nullable"`
	// CSS color for secondary text in the modal body.
	SecondaryTextColor string `json:"secondaryTextColor" api:"nullable"`
	// CSS color for the "on" state of category toggles.
	ToggleOnBg string `json:"toggleOnBg" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonBorderRadius         respjson.Field
		CookieCategoryBlockBg      respjson.Field
		CookieCategoryBlockHoverBg respjson.Field
		FooterBg                   respjson.Field
		FooterColor                respjson.Field
		FooterLinkColor            respjson.Field
		FooterLinkHoverColor       respjson.Field
		ModalBg                    respjson.Field
		ModalBorderRadius          respjson.Field
		PrimaryButtonBg            respjson.Field
		PrimaryButtonColor         respjson.Field
		PrimaryButtonHoverBg       respjson.Field
		PrimaryButtonHoverColor    respjson.Field
		PrimaryTextColor           respjson.Field
		SecondaryButtonBg          respjson.Field
		SecondaryButtonColor       respjson.Field
		SecondaryButtonHoverBg     respjson.Field
		SecondaryButtonHoverColor  respjson.Field
		SecondaryTextColor         respjson.Field
		ToggleOnBg                 respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseDefaultGuiOptionsCssVariables) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingNewResponseDefaultGuiOptionsCssVariables) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the per-category preferences modal.
type ConsentSettingNewResponseDefaultGuiOptionsPreferencesModal struct {
	// Button layout for the preferences modal.
	//
	// Any of "AcceptAllRejectAllSave", "AcceptAllSave".
	ButtonLayout string `json:"buttonLayout" api:"nullable"`
	// When true, accept/reject share the same visual weight.
	EqualWeightButtons bool `json:"equalWeightButtons" api:"nullable"`
	// When true, swap accept/reject button order.
	FlipButtons bool `json:"flipButtons" api:"nullable"`
	// Preferences modal layout (box, bar, bar_wide).
	//
	// Any of "bar", "bar_wide", "box".
	Layout string `json:"layout" api:"nullable"`
	// Preferences modal anchor (left, right).
	//
	// Any of "left", "right".
	Position string `json:"position" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonLayout       respjson.Field
		EqualWeightButtons respjson.Field
		FlipButtons        respjson.Field
		Layout             respjson.Field
		Position           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseDefaultGuiOptionsPreferencesModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingNewResponseDefaultGuiOptionsPreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingNewResponseRegion struct {
	// Region this rule applies to. Use ISO 3166-1 alpha-2 country code ("US", "DE",
	// "BR") or country-subdivision code ("US-CA", "GB-ENG", "CA-ON"). Each region code
	// may appear in only one rule across `regions[]`.
	RegionCode string                              `json:"regionCode" api:"required"`
	Rule       ConsentSettingNewResponseRegionRule `json:"rule" api:"required"`
	// Other region codes that should reuse this rule. Same code-format rules as
	// `regionCode`. Cannot include `regionCode` itself, cannot duplicate, cannot
	// overlap with another rule's regions.
	AdditionalRegions []string `json:"additionalRegions" api:"nullable"`
	// Optional label for this override group, e.g. "CCPA States". Shown in the
	// dashboard instead of the region codes. Pass null to clear it and fall back to
	// the codes.
	Name string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RegionCode        respjson.Field
		Rule              respjson.Field
		AdditionalRegions respjson.Field
		Name              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseRegion) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingNewResponseRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingNewResponseRegionRule struct {
	// Per-category default config for this rule. Every category defined in the
	// top-level `categories[].value` should have an entry here.
	Categories []ConsentSettingNewResponseRegionRuleCategory `json:"categories" api:"required"`
	// BCP 47 default language for this rule. Must have a matching entry in
	// `translations`. Examples: "en", "en-US", "es", "de".
	Language string `json:"language" api:"required"`
	// opt_in: scripts blocked until user accepts (GDPR style). opt_out: scripts run by
	// default until user rejects (CCPA style).
	//
	// Any of "opt_in", "opt_out".
	Mode string `json:"mode" api:"required"`
	// All UI copy, keyed by language. Must include an entry whose `language` matches
	// the rule's `language` field.
	Translations []ConsentSettingNewResponseRegionRuleTranslation `json:"translations" api:"required"`
	// When true, scripts not classified by services[] are blocked until the user opts
	// in.
	AutoblockUnknown bool `json:"autoblockUnknown" api:"nullable"`
	// When true, the consent modal auto-opens on page load.
	AutoShow bool `json:"autoShow" api:"nullable"`
	// Threshold config for autoShowDismissMode (page count or seconds).
	AutoShowDismissConfig ConsentSettingNewResponseRegionRuleAutoShowDismissConfig `json:"autoShowDismissConfig" api:"nullable"`
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	//
	// Any of "after_pages", "after_seconds", "never".
	AutoShowDismissMode string `json:"autoShowDismissMode" api:"nullable"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction bool `json:"disablePageInteraction" api:"nullable"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions ConsentSettingNewResponseRegionRuleGuiOptions `json:"guiOptions" api:"nullable"`
	// When true, the modal is suppressed for known bot user agents.
	HideFromBots bool `json:"hideFromBots" api:"nullable"`
	// When true, the per-service list (services[]) is rendered inside the preferences
	// modal.
	ShowVendorsInPreferences bool `json:"showVendorsInPreferences" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories               respjson.Field
		Language                 respjson.Field
		Mode                     respjson.Field
		Translations             respjson.Field
		AutoblockUnknown         respjson.Field
		AutoShow                 respjson.Field
		AutoShowDismissConfig    respjson.Field
		AutoShowDismissMode      respjson.Field
		DisablePageInteraction   respjson.Field
		GuiOptions               respjson.Field
		HideFromBots             respjson.Field
		ShowVendorsInPreferences respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseRegionRule) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingNewResponseRegionRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingNewResponseRegionRuleCategory struct {
	// Category value (matches `categories[].value`) this entry configures.
	Key   string                                           `json:"key" api:"required"`
	Value ConsentSettingNewResponseRegionRuleCategoryValue `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseRegionRuleCategory) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingNewResponseRegionRuleCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingNewResponseRegionRuleCategoryValue struct {
	// Whether this category is on by default before the user interacts.
	Enabled bool `json:"enabled" api:"required"`
	// When true, this category defaults off if the browser sends Sec-GPC: 1.
	AutoDisableOnGpc bool `json:"autoDisableOnGPC" api:"nullable"`
	// When true, the user cannot toggle this category in the preferences modal.
	ReadOnly bool `json:"readOnly" api:"nullable"`
	// When true, the page reloads after this category is toggled so newly-allowed
	// scripts can run.
	ReloadPage bool `json:"reloadPage" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled          respjson.Field
		AutoDisableOnGpc respjson.Field
		ReadOnly         respjson.Field
		ReloadPage       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseRegionRuleCategoryValue) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingNewResponseRegionRuleCategoryValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingNewResponseRegionRuleTranslation struct {
	// BCP 47 language tag identifying which translation this entry provides. Examples:
	// "en", "en-US", "es", "fr-CA". The default rule's `language` must appear here.
	Language string                                              `json:"language" api:"required"`
	Value    ConsentSettingNewResponseRegionRuleTranslationValue `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Language    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseRegionRuleTranslation) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingNewResponseRegionRuleTranslation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingNewResponseRegionRuleTranslationValue struct {
	// Translated copy for the initial consent modal.
	ConsentModal ConsentSettingNewResponseRegionRuleTranslationValueConsentModal `json:"consentModal" api:"nullable"`
	// Translated copy for the preferences modal.
	PreferencesModal ConsentSettingNewResponseRegionRuleTranslationValuePreferencesModal `json:"preferencesModal" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConsentModal     respjson.Field
		PreferencesModal respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseRegionRuleTranslationValue) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingNewResponseRegionRuleTranslationValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the initial consent modal.
type ConsentSettingNewResponseRegionRuleTranslationValueConsentModal struct {
	// Label for the "accept all" button.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Body text shown in the consent modal.
	Description string `json:"description" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the button that opens the preferences modal.
	ShowPreferencesBtn string `json:"showPreferencesBtn" api:"required"`
	// Heading shown at the top of the consent modal.
	Title string `json:"title" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel string `json:"closeIconLabel" api:"nullable"`
	// Optional footer HTML/text shown below the body.
	Footer string `json:"footer" api:"nullable"`
	// Optional text shown when the browser is sending Sec-GPC: 1.
	GpcNotification string `json:"gpcNotification" api:"nullable"`
	// Link label for the privacy policy (defaults to "Privacy Policy").
	PrivacyPolicyLabel string `json:"privacyPolicyLabel" api:"nullable"`
	// Privacy policy link. Accepts an absolute URL (https://...) or a relative path
	// (/legal/privacy) so one config can run across multiple domains. javascript:,
	// data:, and vbscript: schemes are rejected.
	PrivacyPolicyURL string `json:"privacyPolicyUrl" api:"nullable"`
	// Optional message shown when the consent revision has bumped (re-prompt copy).
	RevisionMessage string `json:"revisionMessage" api:"nullable"`
	// Link label for terms of service (defaults to "Terms of Service").
	TermsOfServiceLabel string `json:"termsOfServiceLabel" api:"nullable"`
	// Terms of service link. Same scheme rules as privacyPolicyUrl: absolute or
	// relative path only.
	TermsOfServiceURL string `json:"termsOfServiceUrl" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptAllBtn        respjson.Field
		Description         respjson.Field
		RejectAllBtn        respjson.Field
		ShowPreferencesBtn  respjson.Field
		Title               respjson.Field
		CloseIconLabel      respjson.Field
		Footer              respjson.Field
		GpcNotification     respjson.Field
		PrivacyPolicyLabel  respjson.Field
		PrivacyPolicyURL    respjson.Field
		RevisionMessage     respjson.Field
		TermsOfServiceLabel respjson.Field
		TermsOfServiceURL   respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseRegionRuleTranslationValueConsentModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingNewResponseRegionRuleTranslationValueConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the preferences modal.
type ConsentSettingNewResponseRegionRuleTranslationValuePreferencesModal struct {
	// Label for the "accept all" button in the preferences modal.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel string `json:"closeIconLabel" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the "save my choices" button.
	SavePreferencesBtn string `json:"savePreferencesBtn" api:"required"`
	// Per-category sections plus a leading general section. The server backfills
	// sections for each defined category if absent.
	Sections []ConsentSettingNewResponseRegionRuleTranslationValuePreferencesModalSection `json:"sections" api:"required"`
	// Heading for the preferences modal.
	Title string `json:"title" api:"required"`
	// Optional collapsible info block rendered after the category sections. Renders
	// only when both title and body are non-empty. Use it for long-form text unrelated
	// to any specific category (e.g. a third-party advertisers disclosure).
	AdditionalInfo ConsentSettingNewResponseRegionRuleTranslationValuePreferencesModalAdditionalInfo `json:"additionalInfo" api:"nullable"`
	// Pluralizable label for the "N services" counter (e.g. "Service|Services").
	ServiceCounterLabel string `json:"serviceCounterLabel" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptAllBtn        respjson.Field
		CloseIconLabel      respjson.Field
		RejectAllBtn        respjson.Field
		SavePreferencesBtn  respjson.Field
		Sections            respjson.Field
		Title               respjson.Field
		AdditionalInfo      respjson.Field
		ServiceCounterLabel respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseRegionRuleTranslationValuePreferencesModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingNewResponseRegionRuleTranslationValuePreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingNewResponseRegionRuleTranslationValuePreferencesModalSection struct {
	// Body text for this preferences-modal section.
	Description string `json:"description" api:"required"`
	// Heading for this preferences-modal section.
	Title string `json:"title" api:"required"`
	// Optional disclosure table listing cookies/services in this section.
	CookieTable ConsentSettingNewResponseRegionRuleTranslationValuePreferencesModalSectionCookieTable `json:"cookieTable" api:"nullable"`
	// Category value this section configures. Use null/omit for the introductory
	// section that applies to no specific category.
	LinkedCategory string `json:"linkedCategory" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description    respjson.Field
		Title          respjson.Field
		CookieTable    respjson.Field
		LinkedCategory respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseRegionRuleTranslationValuePreferencesModalSection) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingNewResponseRegionRuleTranslationValuePreferencesModalSection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional disclosure table listing cookies/services in this section.
type ConsentSettingNewResponseRegionRuleTranslationValuePreferencesModalSectionCookieTable struct {
	// Rows of the cookie table — each row is a {key,value} pair (e.g. cookie name →
	// purpose).
	Body []ConsentSettingNewResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableBody `json:"body" api:"required"`
	// Header cells. Same {key,value} shape as rows.
	Headers []ConsentSettingNewResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader `json:"headers" api:"required"`
	// Optional table caption.
	Caption string `json:"caption" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Headers     respjson.Field
		Caption     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseRegionRuleTranslationValuePreferencesModalSectionCookieTable) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingNewResponseRegionRuleTranslationValuePreferencesModalSectionCookieTable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingNewResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableBody struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableBody) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingNewResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingNewResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingNewResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional collapsible info block rendered after the category sections. Renders
// only when both title and body are non-empty. Use it for long-form text unrelated
// to any specific category (e.g. a third-party advertisers disclosure).
type ConsentSettingNewResponseRegionRuleTranslationValuePreferencesModalAdditionalInfo struct {
	// Body text for the additional-info block. Supports the rich-text markdown subset
	// used by section descriptions: line breaks, [text](url) links (https:// or
	// relative paths), and **bold**. HTML is stripped.
	Body string `json:"body" api:"required"`
	// Heading for the additional-info block.
	Title string `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseRegionRuleTranslationValuePreferencesModalAdditionalInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingNewResponseRegionRuleTranslationValuePreferencesModalAdditionalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Threshold config for autoShowDismissMode (page count or seconds).
type ConsentSettingNewResponseRegionRuleAutoShowDismissConfig struct {
	// Pages-viewed threshold when autoShowDismissMode = after_pages.
	PageCount int64 `json:"pageCount" api:"nullable"`
	// Seconds-on-page threshold when autoShowDismissMode = after_seconds.
	Seconds int64 `json:"seconds" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageCount   respjson.Field
		Seconds     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseRegionRuleAutoShowDismissConfig) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingNewResponseRegionRuleAutoShowDismissConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the modals (layout/position/colors).
type ConsentSettingNewResponseRegionRuleGuiOptions struct {
	// Visual options for the initial consent modal.
	ConsentModal ConsentSettingNewResponseRegionRuleGuiOptionsConsentModal `json:"consentModal" api:"nullable"`
	// Global CSS variable overrides applied to both modals.
	CssVariables ConsentSettingNewResponseRegionRuleGuiOptionsCssVariables `json:"cssVariables" api:"nullable"`
	// Visual options for the per-category preferences modal.
	PreferencesModal ConsentSettingNewResponseRegionRuleGuiOptionsPreferencesModal `json:"preferencesModal" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConsentModal     respjson.Field
		CssVariables     respjson.Field
		PreferencesModal respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseRegionRuleGuiOptions) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingNewResponseRegionRuleGuiOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the initial consent modal.
type ConsentSettingNewResponseRegionRuleGuiOptionsConsentModal struct {
	// Which buttons to render in the consent modal and in what arrangement.
	//
	// Any of "AcceptAllNecessaryPreferences", "AcceptAllPreferences", "AcceptOnly",
	// "InformationOnly", "PreferencesOnly".
	ButtonLayout string `json:"buttonLayout" api:"nullable"`
	// When true, accept and reject buttons share the same visual weight (no
	// primary/secondary).
	EqualWeightButtons bool `json:"equalWeightButtons" api:"nullable"`
	// When true, swap accept/reject button order.
	FlipButtons bool `json:"flipButtons" api:"nullable"`
	// Modal layout style (box, bar, cloud, etc.).
	//
	// Any of "bar", "bar_inline", "box", "box_inline", "box_wide", "cloud",
	// "cloud_inline".
	Layout string `json:"layout" api:"nullable"`
	// Where the modal anchors on the viewport (e.g. bottom_left, top_right,
	// middle_center).
	//
	// Any of "bottom", "bottom_center", "bottom_left", "bottom_right", "middle",
	// "middle_center", "middle_left", "middle_right", "top", "top_center", "top_left",
	// "top_right".
	Position string `json:"position" api:"nullable"`
	// Show an X to dismiss the modal without choosing.
	ShowCloseIcon bool `json:"showCloseIcon" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonLayout       respjson.Field
		EqualWeightButtons respjson.Field
		FlipButtons        respjson.Field
		Layout             respjson.Field
		Position           respjson.Field
		ShowCloseIcon      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseRegionRuleGuiOptionsConsentModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingNewResponseRegionRuleGuiOptionsConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Global CSS variable overrides applied to both modals.
type ConsentSettingNewResponseRegionRuleGuiOptionsCssVariables struct {
	// CSS length, e.g. "4px" or "0.25rem".
	ButtonBorderRadius string `json:"buttonBorderRadius" api:"nullable"`
	// CSS color for the per-category block background in the preferences modal.
	CookieCategoryBlockBg string `json:"cookieCategoryBlockBg" api:"nullable"`
	// CSS color for the per-category block background on hover.
	CookieCategoryBlockHoverBg string `json:"cookieCategoryBlockHoverBg" api:"nullable"`
	// CSS color for the modal footer background.
	FooterBg string `json:"footerBg" api:"nullable"`
	// CSS color for footer text.
	FooterColor string `json:"footerColor" api:"nullable"`
	// CSS color for footer links.
	FooterLinkColor string `json:"footerLinkColor" api:"nullable"`
	// CSS color for footer links on hover.
	FooterLinkHoverColor string `json:"footerLinkHoverColor" api:"nullable"`
	// CSS color for the modal background.
	ModalBg string `json:"modalBg" api:"nullable"`
	// CSS length for modal corner radius.
	ModalBorderRadius string `json:"modalBorderRadius" api:"nullable"`
	// CSS color for the primary button background.
	PrimaryButtonBg string `json:"primaryButtonBg" api:"nullable"`
	// CSS color for primary button text.
	PrimaryButtonColor string `json:"primaryButtonColor" api:"nullable"`
	// CSS color for the primary button background on hover.
	PrimaryButtonHoverBg string `json:"primaryButtonHoverBg" api:"nullable"`
	// CSS color for primary button text on hover.
	PrimaryButtonHoverColor string `json:"primaryButtonHoverColor" api:"nullable"`
	// CSS color for primary text in the modal body.
	PrimaryTextColor string `json:"primaryTextColor" api:"nullable"`
	// CSS color for the secondary button background.
	SecondaryButtonBg string `json:"secondaryButtonBg" api:"nullable"`
	// CSS color for secondary button text.
	SecondaryButtonColor string `json:"secondaryButtonColor" api:"nullable"`
	// CSS color for secondary button background on hover.
	SecondaryButtonHoverBg string `json:"secondaryButtonHoverBg" api:"nullable"`
	// CSS color for secondary button text on hover.
	SecondaryButtonHoverColor string `json:"secondaryButtonHoverColor" api:"nullable"`
	// CSS color for secondary text in the modal body.
	SecondaryTextColor string `json:"secondaryTextColor" api:"nullable"`
	// CSS color for the "on" state of category toggles.
	ToggleOnBg string `json:"toggleOnBg" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonBorderRadius         respjson.Field
		CookieCategoryBlockBg      respjson.Field
		CookieCategoryBlockHoverBg respjson.Field
		FooterBg                   respjson.Field
		FooterColor                respjson.Field
		FooterLinkColor            respjson.Field
		FooterLinkHoverColor       respjson.Field
		ModalBg                    respjson.Field
		ModalBorderRadius          respjson.Field
		PrimaryButtonBg            respjson.Field
		PrimaryButtonColor         respjson.Field
		PrimaryButtonHoverBg       respjson.Field
		PrimaryButtonHoverColor    respjson.Field
		PrimaryTextColor           respjson.Field
		SecondaryButtonBg          respjson.Field
		SecondaryButtonColor       respjson.Field
		SecondaryButtonHoverBg     respjson.Field
		SecondaryButtonHoverColor  respjson.Field
		SecondaryTextColor         respjson.Field
		ToggleOnBg                 respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseRegionRuleGuiOptionsCssVariables) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingNewResponseRegionRuleGuiOptionsCssVariables) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the per-category preferences modal.
type ConsentSettingNewResponseRegionRuleGuiOptionsPreferencesModal struct {
	// Button layout for the preferences modal.
	//
	// Any of "AcceptAllRejectAllSave", "AcceptAllSave".
	ButtonLayout string `json:"buttonLayout" api:"nullable"`
	// When true, accept/reject share the same visual weight.
	EqualWeightButtons bool `json:"equalWeightButtons" api:"nullable"`
	// When true, swap accept/reject button order.
	FlipButtons bool `json:"flipButtons" api:"nullable"`
	// Preferences modal layout (box, bar, bar_wide).
	//
	// Any of "bar", "bar_wide", "box".
	Layout string `json:"layout" api:"nullable"`
	// Preferences modal anchor (left, right).
	//
	// Any of "left", "right".
	Position string `json:"position" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonLayout       respjson.Field
		EqualWeightButtons respjson.Field
		FlipButtons        respjson.Field
		Layout             respjson.Field
		Position           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseRegionRuleGuiOptionsPreferencesModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingNewResponseRegionRuleGuiOptionsPreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingNewResponseService struct {
	// Internal notes shown to admins in the dashboard. Not user-facing.
	InternalNotes string `json:"internalNotes" api:"required"`
	// Display name for this service in the preferences modal.
	Label string `json:"label" api:"required"`
	// Extra category values this service belongs to. Each must match a
	// `categories[].value`.
	AdditionalCategories []string `json:"additionalCategories" api:"nullable"`
	// Primary category value this service belongs to. Must match one of the top-level
	// `categories[].value` entries.
	Category string `json:"category" api:"nullable"`
	// Domains/paths this service matches. Patterns matching the CMP's own scripts
	// (e.g. cdn.oursprivacy.com/cmp-init) are rejected to prevent the CMP blocking
	// itself — use a more specific path like cdn.oursprivacy.com/main.js to block a
	// specific script.
	DomainPatterns []string `json:"domainPatterns" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InternalNotes        respjson.Field
		Label                respjson.Field
		AdditionalCategories respjson.Field
		Category             respjson.Field
		DomainPatterns       respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponseService) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingNewResponseService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enabled serves the CMP on its allowed domains and requires at least one.
// Disabled does not serve the CMP and may have no allowed domains.
type ConsentSettingNewResponseStatus string

const (
	ConsentSettingNewResponseStatusDisabled ConsentSettingNewResponseStatus = "Disabled"
	ConsentSettingNewResponseStatusEnabled  ConsentSettingNewResponseStatus = "Enabled"
)

type ConsentSettingGetResponse struct {
	// Server-assigned UUID for this consent settings record.
	ID string `json:"id" api:"required"`
	// Top-level consent categories (e.g. necessary / analytics / advertising). Server
	// re-stamps `priority` to 0..N on write.
	Categories []ConsentSettingGetResponseCategory `json:"categories" api:"required"`
	// ISO 8601 timestamp when the record was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Default rule used when the user is not in any region listed in `regions[]`.
	Default ConsentSettingGetResponseDefault `json:"default" api:"required"`
	// Discriminator for the entity type. Always "consentSettings".
	Kind string `json:"kind" api:"required"`
	// Human-readable name shown in the dashboard.
	Name string `json:"name" api:"required"`
	// Per-region rule overrides. The first rule whose `regionCode`/`additionalRegions`
	// includes the user's region wins; otherwise `default` applies.
	Regions []ConsentSettingGetResponseRegion `json:"regions" api:"required"`
	// Per-service entries powering "show vendors" and category-aware blocking.
	Services []ConsentSettingGetResponseService `json:"services" api:"required"`
	// Enabled serves the CMP on its allowed domains and requires at least one.
	// Disabled does not serve the CMP and may have no allowed domains.
	//
	// Any of "Disabled", "Enabled".
	Status ConsentSettingGetResponseStatus `json:"status" api:"required"`
	// Name of the cookie that stores the user's consent state. Defaults to
	// "op_consent".
	ConsentCookieName string `json:"consentCookieName" api:"nullable"`
	// Optional custom CDN domain for serving the CMP script (e.g.
	// consent.example.com).
	CustomDomain string `json:"customDomain" api:"nullable"`
	// Cookie/localStorage key for the visitor (device) ID. Must exactly match the web
	// SDK's cookie_names.device_id (default "ours_device_id") so consent and SDK
	// events resolve to the same visitor.
	DeviceIDCookieName string `json:"deviceIdCookieName" api:"nullable"`
	// Revision counter. Bump this to force users who already consented to see the
	// modal again (the SDK compares the persisted revision against this value).
	Revision float64 `json:"revision" api:"nullable"`
	// CSS class names that opt scripts out of consent blocking. Each entry must be a
	// single class token (no whitespace).
	SkipBlockingClassNames []string `json:"skipBlockingClassNames" api:"nullable"`
	// ISO 8601 timestamp of the last write. Null on a freshly created record.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// Pixel of the WebSource that this CMP is wired into. Setting this to a token that
	// is not a valid WebSource of yours is rejected; use null to clear the link.
	WebSDKToken string `json:"webSDKToken" api:"nullable"`
	// Domains where this CMP may run. An Enabled CMP requires at least one allowed
	// domain; with none configured, the CMP does not run. Matching domains may share
	// consent across subdomains.
	WhitelistDomains []string `json:"whitelistDomains" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		Categories             respjson.Field
		CreatedAt              respjson.Field
		Default                respjson.Field
		Kind                   respjson.Field
		Name                   respjson.Field
		Regions                respjson.Field
		Services               respjson.Field
		Status                 respjson.Field
		ConsentCookieName      respjson.Field
		CustomDomain           respjson.Field
		DeviceIDCookieName     respjson.Field
		Revision               respjson.Field
		SkipBlockingClassNames respjson.Field
		UpdatedAt              respjson.Field
		WebSDKToken            respjson.Field
		WhitelistDomains       respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingGetResponseCategory struct {
	// Human-readable label shown next to the toggle in the preferences modal.
	Label string `json:"label" api:"required"`
	// Sort key. Lower numbers render first. Server re-stamps to 0..N on write — send
	// any integer, gaps and duplicates are ironed out.
	Priority int64 `json:"priority" api:"required"`
	// Stable identifier referenced by services and translation sections.
	// Conventionally lowercase (e.g. "necessary", "analytics", "advertising").
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		Priority    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseCategory) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingGetResponseCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default rule used when the user is not in any region listed in `regions[]`.
type ConsentSettingGetResponseDefault struct {
	// Per-category default config for this rule. Every category defined in the
	// top-level `categories[].value` should have an entry here.
	Categories []ConsentSettingGetResponseDefaultCategory `json:"categories" api:"required"`
	// BCP 47 default language for this rule. Must have a matching entry in
	// `translations`. Examples: "en", "en-US", "es", "de".
	Language string `json:"language" api:"required"`
	// opt_in: scripts blocked until user accepts (GDPR style). opt_out: scripts run by
	// default until user rejects (CCPA style).
	//
	// Any of "opt_in", "opt_out".
	Mode string `json:"mode" api:"required"`
	// All UI copy, keyed by language. Must include an entry whose `language` matches
	// the rule's `language` field.
	Translations []ConsentSettingGetResponseDefaultTranslation `json:"translations" api:"required"`
	// When true, scripts not classified by services[] are blocked until the user opts
	// in.
	AutoblockUnknown bool `json:"autoblockUnknown" api:"nullable"`
	// When true, the consent modal auto-opens on page load.
	AutoShow bool `json:"autoShow" api:"nullable"`
	// Threshold config for autoShowDismissMode (page count or seconds).
	AutoShowDismissConfig ConsentSettingGetResponseDefaultAutoShowDismissConfig `json:"autoShowDismissConfig" api:"nullable"`
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	//
	// Any of "after_pages", "after_seconds", "never".
	AutoShowDismissMode string `json:"autoShowDismissMode" api:"nullable"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction bool `json:"disablePageInteraction" api:"nullable"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions ConsentSettingGetResponseDefaultGuiOptions `json:"guiOptions" api:"nullable"`
	// When true, the modal is suppressed for known bot user agents.
	HideFromBots bool `json:"hideFromBots" api:"nullable"`
	// When true, the per-service list (services[]) is rendered inside the preferences
	// modal.
	ShowVendorsInPreferences bool `json:"showVendorsInPreferences" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories               respjson.Field
		Language                 respjson.Field
		Mode                     respjson.Field
		Translations             respjson.Field
		AutoblockUnknown         respjson.Field
		AutoShow                 respjson.Field
		AutoShowDismissConfig    respjson.Field
		AutoShowDismissMode      respjson.Field
		DisablePageInteraction   respjson.Field
		GuiOptions               respjson.Field
		HideFromBots             respjson.Field
		ShowVendorsInPreferences respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseDefault) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingGetResponseDefault) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingGetResponseDefaultCategory struct {
	// Category value (matches `categories[].value`) this entry configures.
	Key   string                                        `json:"key" api:"required"`
	Value ConsentSettingGetResponseDefaultCategoryValue `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseDefaultCategory) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingGetResponseDefaultCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingGetResponseDefaultCategoryValue struct {
	// Whether this category is on by default before the user interacts.
	Enabled bool `json:"enabled" api:"required"`
	// When true, this category defaults off if the browser sends Sec-GPC: 1.
	AutoDisableOnGpc bool `json:"autoDisableOnGPC" api:"nullable"`
	// When true, the user cannot toggle this category in the preferences modal.
	ReadOnly bool `json:"readOnly" api:"nullable"`
	// When true, the page reloads after this category is toggled so newly-allowed
	// scripts can run.
	ReloadPage bool `json:"reloadPage" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled          respjson.Field
		AutoDisableOnGpc respjson.Field
		ReadOnly         respjson.Field
		ReloadPage       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseDefaultCategoryValue) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingGetResponseDefaultCategoryValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingGetResponseDefaultTranslation struct {
	// BCP 47 language tag identifying which translation this entry provides. Examples:
	// "en", "en-US", "es", "fr-CA". The default rule's `language` must appear here.
	Language string                                           `json:"language" api:"required"`
	Value    ConsentSettingGetResponseDefaultTranslationValue `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Language    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseDefaultTranslation) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingGetResponseDefaultTranslation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingGetResponseDefaultTranslationValue struct {
	// Translated copy for the initial consent modal.
	ConsentModal ConsentSettingGetResponseDefaultTranslationValueConsentModal `json:"consentModal" api:"nullable"`
	// Translated copy for the preferences modal.
	PreferencesModal ConsentSettingGetResponseDefaultTranslationValuePreferencesModal `json:"preferencesModal" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConsentModal     respjson.Field
		PreferencesModal respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseDefaultTranslationValue) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingGetResponseDefaultTranslationValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the initial consent modal.
type ConsentSettingGetResponseDefaultTranslationValueConsentModal struct {
	// Label for the "accept all" button.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Body text shown in the consent modal.
	Description string `json:"description" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the button that opens the preferences modal.
	ShowPreferencesBtn string `json:"showPreferencesBtn" api:"required"`
	// Heading shown at the top of the consent modal.
	Title string `json:"title" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel string `json:"closeIconLabel" api:"nullable"`
	// Optional footer HTML/text shown below the body.
	Footer string `json:"footer" api:"nullable"`
	// Optional text shown when the browser is sending Sec-GPC: 1.
	GpcNotification string `json:"gpcNotification" api:"nullable"`
	// Link label for the privacy policy (defaults to "Privacy Policy").
	PrivacyPolicyLabel string `json:"privacyPolicyLabel" api:"nullable"`
	// Privacy policy link. Accepts an absolute URL (https://...) or a relative path
	// (/legal/privacy) so one config can run across multiple domains. javascript:,
	// data:, and vbscript: schemes are rejected.
	PrivacyPolicyURL string `json:"privacyPolicyUrl" api:"nullable"`
	// Optional message shown when the consent revision has bumped (re-prompt copy).
	RevisionMessage string `json:"revisionMessage" api:"nullable"`
	// Link label for terms of service (defaults to "Terms of Service").
	TermsOfServiceLabel string `json:"termsOfServiceLabel" api:"nullable"`
	// Terms of service link. Same scheme rules as privacyPolicyUrl: absolute or
	// relative path only.
	TermsOfServiceURL string `json:"termsOfServiceUrl" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptAllBtn        respjson.Field
		Description         respjson.Field
		RejectAllBtn        respjson.Field
		ShowPreferencesBtn  respjson.Field
		Title               respjson.Field
		CloseIconLabel      respjson.Field
		Footer              respjson.Field
		GpcNotification     respjson.Field
		PrivacyPolicyLabel  respjson.Field
		PrivacyPolicyURL    respjson.Field
		RevisionMessage     respjson.Field
		TermsOfServiceLabel respjson.Field
		TermsOfServiceURL   respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseDefaultTranslationValueConsentModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingGetResponseDefaultTranslationValueConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the preferences modal.
type ConsentSettingGetResponseDefaultTranslationValuePreferencesModal struct {
	// Label for the "accept all" button in the preferences modal.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel string `json:"closeIconLabel" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the "save my choices" button.
	SavePreferencesBtn string `json:"savePreferencesBtn" api:"required"`
	// Per-category sections plus a leading general section. The server backfills
	// sections for each defined category if absent.
	Sections []ConsentSettingGetResponseDefaultTranslationValuePreferencesModalSection `json:"sections" api:"required"`
	// Heading for the preferences modal.
	Title string `json:"title" api:"required"`
	// Optional collapsible info block rendered after the category sections. Renders
	// only when both title and body are non-empty. Use it for long-form text unrelated
	// to any specific category (e.g. a third-party advertisers disclosure).
	AdditionalInfo ConsentSettingGetResponseDefaultTranslationValuePreferencesModalAdditionalInfo `json:"additionalInfo" api:"nullable"`
	// Pluralizable label for the "N services" counter (e.g. "Service|Services").
	ServiceCounterLabel string `json:"serviceCounterLabel" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptAllBtn        respjson.Field
		CloseIconLabel      respjson.Field
		RejectAllBtn        respjson.Field
		SavePreferencesBtn  respjson.Field
		Sections            respjson.Field
		Title               respjson.Field
		AdditionalInfo      respjson.Field
		ServiceCounterLabel respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseDefaultTranslationValuePreferencesModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingGetResponseDefaultTranslationValuePreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingGetResponseDefaultTranslationValuePreferencesModalSection struct {
	// Body text for this preferences-modal section.
	Description string `json:"description" api:"required"`
	// Heading for this preferences-modal section.
	Title string `json:"title" api:"required"`
	// Optional disclosure table listing cookies/services in this section.
	CookieTable ConsentSettingGetResponseDefaultTranslationValuePreferencesModalSectionCookieTable `json:"cookieTable" api:"nullable"`
	// Category value this section configures. Use null/omit for the introductory
	// section that applies to no specific category.
	LinkedCategory string `json:"linkedCategory" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description    respjson.Field
		Title          respjson.Field
		CookieTable    respjson.Field
		LinkedCategory respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseDefaultTranslationValuePreferencesModalSection) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingGetResponseDefaultTranslationValuePreferencesModalSection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional disclosure table listing cookies/services in this section.
type ConsentSettingGetResponseDefaultTranslationValuePreferencesModalSectionCookieTable struct {
	// Rows of the cookie table — each row is a {key,value} pair (e.g. cookie name →
	// purpose).
	Body []ConsentSettingGetResponseDefaultTranslationValuePreferencesModalSectionCookieTableBody `json:"body" api:"required"`
	// Header cells. Same {key,value} shape as rows.
	Headers []ConsentSettingGetResponseDefaultTranslationValuePreferencesModalSectionCookieTableHeader `json:"headers" api:"required"`
	// Optional table caption.
	Caption string `json:"caption" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Headers     respjson.Field
		Caption     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseDefaultTranslationValuePreferencesModalSectionCookieTable) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingGetResponseDefaultTranslationValuePreferencesModalSectionCookieTable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingGetResponseDefaultTranslationValuePreferencesModalSectionCookieTableBody struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseDefaultTranslationValuePreferencesModalSectionCookieTableBody) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingGetResponseDefaultTranslationValuePreferencesModalSectionCookieTableBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingGetResponseDefaultTranslationValuePreferencesModalSectionCookieTableHeader struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseDefaultTranslationValuePreferencesModalSectionCookieTableHeader) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingGetResponseDefaultTranslationValuePreferencesModalSectionCookieTableHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional collapsible info block rendered after the category sections. Renders
// only when both title and body are non-empty. Use it for long-form text unrelated
// to any specific category (e.g. a third-party advertisers disclosure).
type ConsentSettingGetResponseDefaultTranslationValuePreferencesModalAdditionalInfo struct {
	// Body text for the additional-info block. Supports the rich-text markdown subset
	// used by section descriptions: line breaks, [text](url) links (https:// or
	// relative paths), and **bold**. HTML is stripped.
	Body string `json:"body" api:"required"`
	// Heading for the additional-info block.
	Title string `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseDefaultTranslationValuePreferencesModalAdditionalInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingGetResponseDefaultTranslationValuePreferencesModalAdditionalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Threshold config for autoShowDismissMode (page count or seconds).
type ConsentSettingGetResponseDefaultAutoShowDismissConfig struct {
	// Pages-viewed threshold when autoShowDismissMode = after_pages.
	PageCount int64 `json:"pageCount" api:"nullable"`
	// Seconds-on-page threshold when autoShowDismissMode = after_seconds.
	Seconds int64 `json:"seconds" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageCount   respjson.Field
		Seconds     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseDefaultAutoShowDismissConfig) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingGetResponseDefaultAutoShowDismissConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the modals (layout/position/colors).
type ConsentSettingGetResponseDefaultGuiOptions struct {
	// Visual options for the initial consent modal.
	ConsentModal ConsentSettingGetResponseDefaultGuiOptionsConsentModal `json:"consentModal" api:"nullable"`
	// Global CSS variable overrides applied to both modals.
	CssVariables ConsentSettingGetResponseDefaultGuiOptionsCssVariables `json:"cssVariables" api:"nullable"`
	// Visual options for the per-category preferences modal.
	PreferencesModal ConsentSettingGetResponseDefaultGuiOptionsPreferencesModal `json:"preferencesModal" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConsentModal     respjson.Field
		CssVariables     respjson.Field
		PreferencesModal respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseDefaultGuiOptions) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingGetResponseDefaultGuiOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the initial consent modal.
type ConsentSettingGetResponseDefaultGuiOptionsConsentModal struct {
	// Which buttons to render in the consent modal and in what arrangement.
	//
	// Any of "AcceptAllNecessaryPreferences", "AcceptAllPreferences", "AcceptOnly",
	// "InformationOnly", "PreferencesOnly".
	ButtonLayout string `json:"buttonLayout" api:"nullable"`
	// When true, accept and reject buttons share the same visual weight (no
	// primary/secondary).
	EqualWeightButtons bool `json:"equalWeightButtons" api:"nullable"`
	// When true, swap accept/reject button order.
	FlipButtons bool `json:"flipButtons" api:"nullable"`
	// Modal layout style (box, bar, cloud, etc.).
	//
	// Any of "bar", "bar_inline", "box", "box_inline", "box_wide", "cloud",
	// "cloud_inline".
	Layout string `json:"layout" api:"nullable"`
	// Where the modal anchors on the viewport (e.g. bottom_left, top_right,
	// middle_center).
	//
	// Any of "bottom", "bottom_center", "bottom_left", "bottom_right", "middle",
	// "middle_center", "middle_left", "middle_right", "top", "top_center", "top_left",
	// "top_right".
	Position string `json:"position" api:"nullable"`
	// Show an X to dismiss the modal without choosing.
	ShowCloseIcon bool `json:"showCloseIcon" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonLayout       respjson.Field
		EqualWeightButtons respjson.Field
		FlipButtons        respjson.Field
		Layout             respjson.Field
		Position           respjson.Field
		ShowCloseIcon      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseDefaultGuiOptionsConsentModal) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingGetResponseDefaultGuiOptionsConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Global CSS variable overrides applied to both modals.
type ConsentSettingGetResponseDefaultGuiOptionsCssVariables struct {
	// CSS length, e.g. "4px" or "0.25rem".
	ButtonBorderRadius string `json:"buttonBorderRadius" api:"nullable"`
	// CSS color for the per-category block background in the preferences modal.
	CookieCategoryBlockBg string `json:"cookieCategoryBlockBg" api:"nullable"`
	// CSS color for the per-category block background on hover.
	CookieCategoryBlockHoverBg string `json:"cookieCategoryBlockHoverBg" api:"nullable"`
	// CSS color for the modal footer background.
	FooterBg string `json:"footerBg" api:"nullable"`
	// CSS color for footer text.
	FooterColor string `json:"footerColor" api:"nullable"`
	// CSS color for footer links.
	FooterLinkColor string `json:"footerLinkColor" api:"nullable"`
	// CSS color for footer links on hover.
	FooterLinkHoverColor string `json:"footerLinkHoverColor" api:"nullable"`
	// CSS color for the modal background.
	ModalBg string `json:"modalBg" api:"nullable"`
	// CSS length for modal corner radius.
	ModalBorderRadius string `json:"modalBorderRadius" api:"nullable"`
	// CSS color for the primary button background.
	PrimaryButtonBg string `json:"primaryButtonBg" api:"nullable"`
	// CSS color for primary button text.
	PrimaryButtonColor string `json:"primaryButtonColor" api:"nullable"`
	// CSS color for the primary button background on hover.
	PrimaryButtonHoverBg string `json:"primaryButtonHoverBg" api:"nullable"`
	// CSS color for primary button text on hover.
	PrimaryButtonHoverColor string `json:"primaryButtonHoverColor" api:"nullable"`
	// CSS color for primary text in the modal body.
	PrimaryTextColor string `json:"primaryTextColor" api:"nullable"`
	// CSS color for the secondary button background.
	SecondaryButtonBg string `json:"secondaryButtonBg" api:"nullable"`
	// CSS color for secondary button text.
	SecondaryButtonColor string `json:"secondaryButtonColor" api:"nullable"`
	// CSS color for secondary button background on hover.
	SecondaryButtonHoverBg string `json:"secondaryButtonHoverBg" api:"nullable"`
	// CSS color for secondary button text on hover.
	SecondaryButtonHoverColor string `json:"secondaryButtonHoverColor" api:"nullable"`
	// CSS color for secondary text in the modal body.
	SecondaryTextColor string `json:"secondaryTextColor" api:"nullable"`
	// CSS color for the "on" state of category toggles.
	ToggleOnBg string `json:"toggleOnBg" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonBorderRadius         respjson.Field
		CookieCategoryBlockBg      respjson.Field
		CookieCategoryBlockHoverBg respjson.Field
		FooterBg                   respjson.Field
		FooterColor                respjson.Field
		FooterLinkColor            respjson.Field
		FooterLinkHoverColor       respjson.Field
		ModalBg                    respjson.Field
		ModalBorderRadius          respjson.Field
		PrimaryButtonBg            respjson.Field
		PrimaryButtonColor         respjson.Field
		PrimaryButtonHoverBg       respjson.Field
		PrimaryButtonHoverColor    respjson.Field
		PrimaryTextColor           respjson.Field
		SecondaryButtonBg          respjson.Field
		SecondaryButtonColor       respjson.Field
		SecondaryButtonHoverBg     respjson.Field
		SecondaryButtonHoverColor  respjson.Field
		SecondaryTextColor         respjson.Field
		ToggleOnBg                 respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseDefaultGuiOptionsCssVariables) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingGetResponseDefaultGuiOptionsCssVariables) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the per-category preferences modal.
type ConsentSettingGetResponseDefaultGuiOptionsPreferencesModal struct {
	// Button layout for the preferences modal.
	//
	// Any of "AcceptAllRejectAllSave", "AcceptAllSave".
	ButtonLayout string `json:"buttonLayout" api:"nullable"`
	// When true, accept/reject share the same visual weight.
	EqualWeightButtons bool `json:"equalWeightButtons" api:"nullable"`
	// When true, swap accept/reject button order.
	FlipButtons bool `json:"flipButtons" api:"nullable"`
	// Preferences modal layout (box, bar, bar_wide).
	//
	// Any of "bar", "bar_wide", "box".
	Layout string `json:"layout" api:"nullable"`
	// Preferences modal anchor (left, right).
	//
	// Any of "left", "right".
	Position string `json:"position" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonLayout       respjson.Field
		EqualWeightButtons respjson.Field
		FlipButtons        respjson.Field
		Layout             respjson.Field
		Position           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseDefaultGuiOptionsPreferencesModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingGetResponseDefaultGuiOptionsPreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingGetResponseRegion struct {
	// Region this rule applies to. Use ISO 3166-1 alpha-2 country code ("US", "DE",
	// "BR") or country-subdivision code ("US-CA", "GB-ENG", "CA-ON"). Each region code
	// may appear in only one rule across `regions[]`.
	RegionCode string                              `json:"regionCode" api:"required"`
	Rule       ConsentSettingGetResponseRegionRule `json:"rule" api:"required"`
	// Other region codes that should reuse this rule. Same code-format rules as
	// `regionCode`. Cannot include `regionCode` itself, cannot duplicate, cannot
	// overlap with another rule's regions.
	AdditionalRegions []string `json:"additionalRegions" api:"nullable"`
	// Optional label for this override group, e.g. "CCPA States". Shown in the
	// dashboard instead of the region codes. Pass null to clear it and fall back to
	// the codes.
	Name string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RegionCode        respjson.Field
		Rule              respjson.Field
		AdditionalRegions respjson.Field
		Name              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseRegion) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingGetResponseRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingGetResponseRegionRule struct {
	// Per-category default config for this rule. Every category defined in the
	// top-level `categories[].value` should have an entry here.
	Categories []ConsentSettingGetResponseRegionRuleCategory `json:"categories" api:"required"`
	// BCP 47 default language for this rule. Must have a matching entry in
	// `translations`. Examples: "en", "en-US", "es", "de".
	Language string `json:"language" api:"required"`
	// opt_in: scripts blocked until user accepts (GDPR style). opt_out: scripts run by
	// default until user rejects (CCPA style).
	//
	// Any of "opt_in", "opt_out".
	Mode string `json:"mode" api:"required"`
	// All UI copy, keyed by language. Must include an entry whose `language` matches
	// the rule's `language` field.
	Translations []ConsentSettingGetResponseRegionRuleTranslation `json:"translations" api:"required"`
	// When true, scripts not classified by services[] are blocked until the user opts
	// in.
	AutoblockUnknown bool `json:"autoblockUnknown" api:"nullable"`
	// When true, the consent modal auto-opens on page load.
	AutoShow bool `json:"autoShow" api:"nullable"`
	// Threshold config for autoShowDismissMode (page count or seconds).
	AutoShowDismissConfig ConsentSettingGetResponseRegionRuleAutoShowDismissConfig `json:"autoShowDismissConfig" api:"nullable"`
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	//
	// Any of "after_pages", "after_seconds", "never".
	AutoShowDismissMode string `json:"autoShowDismissMode" api:"nullable"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction bool `json:"disablePageInteraction" api:"nullable"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions ConsentSettingGetResponseRegionRuleGuiOptions `json:"guiOptions" api:"nullable"`
	// When true, the modal is suppressed for known bot user agents.
	HideFromBots bool `json:"hideFromBots" api:"nullable"`
	// When true, the per-service list (services[]) is rendered inside the preferences
	// modal.
	ShowVendorsInPreferences bool `json:"showVendorsInPreferences" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories               respjson.Field
		Language                 respjson.Field
		Mode                     respjson.Field
		Translations             respjson.Field
		AutoblockUnknown         respjson.Field
		AutoShow                 respjson.Field
		AutoShowDismissConfig    respjson.Field
		AutoShowDismissMode      respjson.Field
		DisablePageInteraction   respjson.Field
		GuiOptions               respjson.Field
		HideFromBots             respjson.Field
		ShowVendorsInPreferences respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseRegionRule) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingGetResponseRegionRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingGetResponseRegionRuleCategory struct {
	// Category value (matches `categories[].value`) this entry configures.
	Key   string                                           `json:"key" api:"required"`
	Value ConsentSettingGetResponseRegionRuleCategoryValue `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseRegionRuleCategory) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingGetResponseRegionRuleCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingGetResponseRegionRuleCategoryValue struct {
	// Whether this category is on by default before the user interacts.
	Enabled bool `json:"enabled" api:"required"`
	// When true, this category defaults off if the browser sends Sec-GPC: 1.
	AutoDisableOnGpc bool `json:"autoDisableOnGPC" api:"nullable"`
	// When true, the user cannot toggle this category in the preferences modal.
	ReadOnly bool `json:"readOnly" api:"nullable"`
	// When true, the page reloads after this category is toggled so newly-allowed
	// scripts can run.
	ReloadPage bool `json:"reloadPage" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled          respjson.Field
		AutoDisableOnGpc respjson.Field
		ReadOnly         respjson.Field
		ReloadPage       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseRegionRuleCategoryValue) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingGetResponseRegionRuleCategoryValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingGetResponseRegionRuleTranslation struct {
	// BCP 47 language tag identifying which translation this entry provides. Examples:
	// "en", "en-US", "es", "fr-CA". The default rule's `language` must appear here.
	Language string                                              `json:"language" api:"required"`
	Value    ConsentSettingGetResponseRegionRuleTranslationValue `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Language    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseRegionRuleTranslation) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingGetResponseRegionRuleTranslation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingGetResponseRegionRuleTranslationValue struct {
	// Translated copy for the initial consent modal.
	ConsentModal ConsentSettingGetResponseRegionRuleTranslationValueConsentModal `json:"consentModal" api:"nullable"`
	// Translated copy for the preferences modal.
	PreferencesModal ConsentSettingGetResponseRegionRuleTranslationValuePreferencesModal `json:"preferencesModal" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConsentModal     respjson.Field
		PreferencesModal respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseRegionRuleTranslationValue) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingGetResponseRegionRuleTranslationValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the initial consent modal.
type ConsentSettingGetResponseRegionRuleTranslationValueConsentModal struct {
	// Label for the "accept all" button.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Body text shown in the consent modal.
	Description string `json:"description" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the button that opens the preferences modal.
	ShowPreferencesBtn string `json:"showPreferencesBtn" api:"required"`
	// Heading shown at the top of the consent modal.
	Title string `json:"title" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel string `json:"closeIconLabel" api:"nullable"`
	// Optional footer HTML/text shown below the body.
	Footer string `json:"footer" api:"nullable"`
	// Optional text shown when the browser is sending Sec-GPC: 1.
	GpcNotification string `json:"gpcNotification" api:"nullable"`
	// Link label for the privacy policy (defaults to "Privacy Policy").
	PrivacyPolicyLabel string `json:"privacyPolicyLabel" api:"nullable"`
	// Privacy policy link. Accepts an absolute URL (https://...) or a relative path
	// (/legal/privacy) so one config can run across multiple domains. javascript:,
	// data:, and vbscript: schemes are rejected.
	PrivacyPolicyURL string `json:"privacyPolicyUrl" api:"nullable"`
	// Optional message shown when the consent revision has bumped (re-prompt copy).
	RevisionMessage string `json:"revisionMessage" api:"nullable"`
	// Link label for terms of service (defaults to "Terms of Service").
	TermsOfServiceLabel string `json:"termsOfServiceLabel" api:"nullable"`
	// Terms of service link. Same scheme rules as privacyPolicyUrl: absolute or
	// relative path only.
	TermsOfServiceURL string `json:"termsOfServiceUrl" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptAllBtn        respjson.Field
		Description         respjson.Field
		RejectAllBtn        respjson.Field
		ShowPreferencesBtn  respjson.Field
		Title               respjson.Field
		CloseIconLabel      respjson.Field
		Footer              respjson.Field
		GpcNotification     respjson.Field
		PrivacyPolicyLabel  respjson.Field
		PrivacyPolicyURL    respjson.Field
		RevisionMessage     respjson.Field
		TermsOfServiceLabel respjson.Field
		TermsOfServiceURL   respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseRegionRuleTranslationValueConsentModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingGetResponseRegionRuleTranslationValueConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the preferences modal.
type ConsentSettingGetResponseRegionRuleTranslationValuePreferencesModal struct {
	// Label for the "accept all" button in the preferences modal.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel string `json:"closeIconLabel" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the "save my choices" button.
	SavePreferencesBtn string `json:"savePreferencesBtn" api:"required"`
	// Per-category sections plus a leading general section. The server backfills
	// sections for each defined category if absent.
	Sections []ConsentSettingGetResponseRegionRuleTranslationValuePreferencesModalSection `json:"sections" api:"required"`
	// Heading for the preferences modal.
	Title string `json:"title" api:"required"`
	// Optional collapsible info block rendered after the category sections. Renders
	// only when both title and body are non-empty. Use it for long-form text unrelated
	// to any specific category (e.g. a third-party advertisers disclosure).
	AdditionalInfo ConsentSettingGetResponseRegionRuleTranslationValuePreferencesModalAdditionalInfo `json:"additionalInfo" api:"nullable"`
	// Pluralizable label for the "N services" counter (e.g. "Service|Services").
	ServiceCounterLabel string `json:"serviceCounterLabel" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptAllBtn        respjson.Field
		CloseIconLabel      respjson.Field
		RejectAllBtn        respjson.Field
		SavePreferencesBtn  respjson.Field
		Sections            respjson.Field
		Title               respjson.Field
		AdditionalInfo      respjson.Field
		ServiceCounterLabel respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseRegionRuleTranslationValuePreferencesModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingGetResponseRegionRuleTranslationValuePreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingGetResponseRegionRuleTranslationValuePreferencesModalSection struct {
	// Body text for this preferences-modal section.
	Description string `json:"description" api:"required"`
	// Heading for this preferences-modal section.
	Title string `json:"title" api:"required"`
	// Optional disclosure table listing cookies/services in this section.
	CookieTable ConsentSettingGetResponseRegionRuleTranslationValuePreferencesModalSectionCookieTable `json:"cookieTable" api:"nullable"`
	// Category value this section configures. Use null/omit for the introductory
	// section that applies to no specific category.
	LinkedCategory string `json:"linkedCategory" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description    respjson.Field
		Title          respjson.Field
		CookieTable    respjson.Field
		LinkedCategory respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseRegionRuleTranslationValuePreferencesModalSection) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingGetResponseRegionRuleTranslationValuePreferencesModalSection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional disclosure table listing cookies/services in this section.
type ConsentSettingGetResponseRegionRuleTranslationValuePreferencesModalSectionCookieTable struct {
	// Rows of the cookie table — each row is a {key,value} pair (e.g. cookie name →
	// purpose).
	Body []ConsentSettingGetResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableBody `json:"body" api:"required"`
	// Header cells. Same {key,value} shape as rows.
	Headers []ConsentSettingGetResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader `json:"headers" api:"required"`
	// Optional table caption.
	Caption string `json:"caption" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Headers     respjson.Field
		Caption     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseRegionRuleTranslationValuePreferencesModalSectionCookieTable) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingGetResponseRegionRuleTranslationValuePreferencesModalSectionCookieTable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingGetResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableBody struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableBody) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingGetResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingGetResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingGetResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional collapsible info block rendered after the category sections. Renders
// only when both title and body are non-empty. Use it for long-form text unrelated
// to any specific category (e.g. a third-party advertisers disclosure).
type ConsentSettingGetResponseRegionRuleTranslationValuePreferencesModalAdditionalInfo struct {
	// Body text for the additional-info block. Supports the rich-text markdown subset
	// used by section descriptions: line breaks, [text](url) links (https:// or
	// relative paths), and **bold**. HTML is stripped.
	Body string `json:"body" api:"required"`
	// Heading for the additional-info block.
	Title string `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseRegionRuleTranslationValuePreferencesModalAdditionalInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingGetResponseRegionRuleTranslationValuePreferencesModalAdditionalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Threshold config for autoShowDismissMode (page count or seconds).
type ConsentSettingGetResponseRegionRuleAutoShowDismissConfig struct {
	// Pages-viewed threshold when autoShowDismissMode = after_pages.
	PageCount int64 `json:"pageCount" api:"nullable"`
	// Seconds-on-page threshold when autoShowDismissMode = after_seconds.
	Seconds int64 `json:"seconds" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageCount   respjson.Field
		Seconds     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseRegionRuleAutoShowDismissConfig) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingGetResponseRegionRuleAutoShowDismissConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the modals (layout/position/colors).
type ConsentSettingGetResponseRegionRuleGuiOptions struct {
	// Visual options for the initial consent modal.
	ConsentModal ConsentSettingGetResponseRegionRuleGuiOptionsConsentModal `json:"consentModal" api:"nullable"`
	// Global CSS variable overrides applied to both modals.
	CssVariables ConsentSettingGetResponseRegionRuleGuiOptionsCssVariables `json:"cssVariables" api:"nullable"`
	// Visual options for the per-category preferences modal.
	PreferencesModal ConsentSettingGetResponseRegionRuleGuiOptionsPreferencesModal `json:"preferencesModal" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConsentModal     respjson.Field
		CssVariables     respjson.Field
		PreferencesModal respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseRegionRuleGuiOptions) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingGetResponseRegionRuleGuiOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the initial consent modal.
type ConsentSettingGetResponseRegionRuleGuiOptionsConsentModal struct {
	// Which buttons to render in the consent modal and in what arrangement.
	//
	// Any of "AcceptAllNecessaryPreferences", "AcceptAllPreferences", "AcceptOnly",
	// "InformationOnly", "PreferencesOnly".
	ButtonLayout string `json:"buttonLayout" api:"nullable"`
	// When true, accept and reject buttons share the same visual weight (no
	// primary/secondary).
	EqualWeightButtons bool `json:"equalWeightButtons" api:"nullable"`
	// When true, swap accept/reject button order.
	FlipButtons bool `json:"flipButtons" api:"nullable"`
	// Modal layout style (box, bar, cloud, etc.).
	//
	// Any of "bar", "bar_inline", "box", "box_inline", "box_wide", "cloud",
	// "cloud_inline".
	Layout string `json:"layout" api:"nullable"`
	// Where the modal anchors on the viewport (e.g. bottom_left, top_right,
	// middle_center).
	//
	// Any of "bottom", "bottom_center", "bottom_left", "bottom_right", "middle",
	// "middle_center", "middle_left", "middle_right", "top", "top_center", "top_left",
	// "top_right".
	Position string `json:"position" api:"nullable"`
	// Show an X to dismiss the modal without choosing.
	ShowCloseIcon bool `json:"showCloseIcon" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonLayout       respjson.Field
		EqualWeightButtons respjson.Field
		FlipButtons        respjson.Field
		Layout             respjson.Field
		Position           respjson.Field
		ShowCloseIcon      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseRegionRuleGuiOptionsConsentModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingGetResponseRegionRuleGuiOptionsConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Global CSS variable overrides applied to both modals.
type ConsentSettingGetResponseRegionRuleGuiOptionsCssVariables struct {
	// CSS length, e.g. "4px" or "0.25rem".
	ButtonBorderRadius string `json:"buttonBorderRadius" api:"nullable"`
	// CSS color for the per-category block background in the preferences modal.
	CookieCategoryBlockBg string `json:"cookieCategoryBlockBg" api:"nullable"`
	// CSS color for the per-category block background on hover.
	CookieCategoryBlockHoverBg string `json:"cookieCategoryBlockHoverBg" api:"nullable"`
	// CSS color for the modal footer background.
	FooterBg string `json:"footerBg" api:"nullable"`
	// CSS color for footer text.
	FooterColor string `json:"footerColor" api:"nullable"`
	// CSS color for footer links.
	FooterLinkColor string `json:"footerLinkColor" api:"nullable"`
	// CSS color for footer links on hover.
	FooterLinkHoverColor string `json:"footerLinkHoverColor" api:"nullable"`
	// CSS color for the modal background.
	ModalBg string `json:"modalBg" api:"nullable"`
	// CSS length for modal corner radius.
	ModalBorderRadius string `json:"modalBorderRadius" api:"nullable"`
	// CSS color for the primary button background.
	PrimaryButtonBg string `json:"primaryButtonBg" api:"nullable"`
	// CSS color for primary button text.
	PrimaryButtonColor string `json:"primaryButtonColor" api:"nullable"`
	// CSS color for the primary button background on hover.
	PrimaryButtonHoverBg string `json:"primaryButtonHoverBg" api:"nullable"`
	// CSS color for primary button text on hover.
	PrimaryButtonHoverColor string `json:"primaryButtonHoverColor" api:"nullable"`
	// CSS color for primary text in the modal body.
	PrimaryTextColor string `json:"primaryTextColor" api:"nullable"`
	// CSS color for the secondary button background.
	SecondaryButtonBg string `json:"secondaryButtonBg" api:"nullable"`
	// CSS color for secondary button text.
	SecondaryButtonColor string `json:"secondaryButtonColor" api:"nullable"`
	// CSS color for secondary button background on hover.
	SecondaryButtonHoverBg string `json:"secondaryButtonHoverBg" api:"nullable"`
	// CSS color for secondary button text on hover.
	SecondaryButtonHoverColor string `json:"secondaryButtonHoverColor" api:"nullable"`
	// CSS color for secondary text in the modal body.
	SecondaryTextColor string `json:"secondaryTextColor" api:"nullable"`
	// CSS color for the "on" state of category toggles.
	ToggleOnBg string `json:"toggleOnBg" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonBorderRadius         respjson.Field
		CookieCategoryBlockBg      respjson.Field
		CookieCategoryBlockHoverBg respjson.Field
		FooterBg                   respjson.Field
		FooterColor                respjson.Field
		FooterLinkColor            respjson.Field
		FooterLinkHoverColor       respjson.Field
		ModalBg                    respjson.Field
		ModalBorderRadius          respjson.Field
		PrimaryButtonBg            respjson.Field
		PrimaryButtonColor         respjson.Field
		PrimaryButtonHoverBg       respjson.Field
		PrimaryButtonHoverColor    respjson.Field
		PrimaryTextColor           respjson.Field
		SecondaryButtonBg          respjson.Field
		SecondaryButtonColor       respjson.Field
		SecondaryButtonHoverBg     respjson.Field
		SecondaryButtonHoverColor  respjson.Field
		SecondaryTextColor         respjson.Field
		ToggleOnBg                 respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseRegionRuleGuiOptionsCssVariables) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingGetResponseRegionRuleGuiOptionsCssVariables) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the per-category preferences modal.
type ConsentSettingGetResponseRegionRuleGuiOptionsPreferencesModal struct {
	// Button layout for the preferences modal.
	//
	// Any of "AcceptAllRejectAllSave", "AcceptAllSave".
	ButtonLayout string `json:"buttonLayout" api:"nullable"`
	// When true, accept/reject share the same visual weight.
	EqualWeightButtons bool `json:"equalWeightButtons" api:"nullable"`
	// When true, swap accept/reject button order.
	FlipButtons bool `json:"flipButtons" api:"nullable"`
	// Preferences modal layout (box, bar, bar_wide).
	//
	// Any of "bar", "bar_wide", "box".
	Layout string `json:"layout" api:"nullable"`
	// Preferences modal anchor (left, right).
	//
	// Any of "left", "right".
	Position string `json:"position" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonLayout       respjson.Field
		EqualWeightButtons respjson.Field
		FlipButtons        respjson.Field
		Layout             respjson.Field
		Position           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseRegionRuleGuiOptionsPreferencesModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingGetResponseRegionRuleGuiOptionsPreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingGetResponseService struct {
	// Internal notes shown to admins in the dashboard. Not user-facing.
	InternalNotes string `json:"internalNotes" api:"required"`
	// Display name for this service in the preferences modal.
	Label string `json:"label" api:"required"`
	// Extra category values this service belongs to. Each must match a
	// `categories[].value`.
	AdditionalCategories []string `json:"additionalCategories" api:"nullable"`
	// Primary category value this service belongs to. Must match one of the top-level
	// `categories[].value` entries.
	Category string `json:"category" api:"nullable"`
	// Domains/paths this service matches. Patterns matching the CMP's own scripts
	// (e.g. cdn.oursprivacy.com/cmp-init) are rejected to prevent the CMP blocking
	// itself — use a more specific path like cdn.oursprivacy.com/main.js to block a
	// specific script.
	DomainPatterns []string `json:"domainPatterns" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InternalNotes        respjson.Field
		Label                respjson.Field
		AdditionalCategories respjson.Field
		Category             respjson.Field
		DomainPatterns       respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponseService) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingGetResponseService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enabled serves the CMP on its allowed domains and requires at least one.
// Disabled does not serve the CMP and may have no allowed domains.
type ConsentSettingGetResponseStatus string

const (
	ConsentSettingGetResponseStatusDisabled ConsentSettingGetResponseStatus = "Disabled"
	ConsentSettingGetResponseStatusEnabled  ConsentSettingGetResponseStatus = "Enabled"
)

type ConsentSettingReplaceResponse struct {
	// Server-assigned UUID for this consent settings record.
	ID string `json:"id" api:"required"`
	// Top-level consent categories (e.g. necessary / analytics / advertising). Server
	// re-stamps `priority` to 0..N on write.
	Categories []ConsentSettingReplaceResponseCategory `json:"categories" api:"required"`
	// ISO 8601 timestamp when the record was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Default rule used when the user is not in any region listed in `regions[]`.
	Default ConsentSettingReplaceResponseDefault `json:"default" api:"required"`
	// Discriminator for the entity type. Always "consentSettings".
	Kind string `json:"kind" api:"required"`
	// Human-readable name shown in the dashboard.
	Name string `json:"name" api:"required"`
	// Per-region rule overrides. The first rule whose `regionCode`/`additionalRegions`
	// includes the user's region wins; otherwise `default` applies.
	Regions []ConsentSettingReplaceResponseRegion `json:"regions" api:"required"`
	// Per-service entries powering "show vendors" and category-aware blocking.
	Services []ConsentSettingReplaceResponseService `json:"services" api:"required"`
	// Enabled serves the CMP on its allowed domains and requires at least one.
	// Disabled does not serve the CMP and may have no allowed domains.
	//
	// Any of "Disabled", "Enabled".
	Status ConsentSettingReplaceResponseStatus `json:"status" api:"required"`
	// Name of the cookie that stores the user's consent state. Defaults to
	// "op_consent".
	ConsentCookieName string `json:"consentCookieName" api:"nullable"`
	// Optional custom CDN domain for serving the CMP script (e.g.
	// consent.example.com).
	CustomDomain string `json:"customDomain" api:"nullable"`
	// Cookie/localStorage key for the visitor (device) ID. Must exactly match the web
	// SDK's cookie_names.device_id (default "ours_device_id") so consent and SDK
	// events resolve to the same visitor.
	DeviceIDCookieName string `json:"deviceIdCookieName" api:"nullable"`
	// Revision counter. Bump this to force users who already consented to see the
	// modal again (the SDK compares the persisted revision against this value).
	Revision float64 `json:"revision" api:"nullable"`
	// CSS class names that opt scripts out of consent blocking. Each entry must be a
	// single class token (no whitespace).
	SkipBlockingClassNames []string `json:"skipBlockingClassNames" api:"nullable"`
	// ISO 8601 timestamp of the last write. Null on a freshly created record.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// Pixel of the WebSource that this CMP is wired into. Setting this to a token that
	// is not a valid WebSource of yours is rejected; use null to clear the link.
	WebSDKToken string `json:"webSDKToken" api:"nullable"`
	// Domains where this CMP may run. An Enabled CMP requires at least one allowed
	// domain; with none configured, the CMP does not run. Matching domains may share
	// consent across subdomains.
	WhitelistDomains []string `json:"whitelistDomains" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		Categories             respjson.Field
		CreatedAt              respjson.Field
		Default                respjson.Field
		Kind                   respjson.Field
		Name                   respjson.Field
		Regions                respjson.Field
		Services               respjson.Field
		Status                 respjson.Field
		ConsentCookieName      respjson.Field
		CustomDomain           respjson.Field
		DeviceIDCookieName     respjson.Field
		Revision               respjson.Field
		SkipBlockingClassNames respjson.Field
		UpdatedAt              respjson.Field
		WebSDKToken            respjson.Field
		WhitelistDomains       respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponse) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingReplaceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingReplaceResponseCategory struct {
	// Human-readable label shown next to the toggle in the preferences modal.
	Label string `json:"label" api:"required"`
	// Sort key. Lower numbers render first. Server re-stamps to 0..N on write — send
	// any integer, gaps and duplicates are ironed out.
	Priority int64 `json:"priority" api:"required"`
	// Stable identifier referenced by services and translation sections.
	// Conventionally lowercase (e.g. "necessary", "analytics", "advertising").
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		Priority    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseCategory) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingReplaceResponseCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default rule used when the user is not in any region listed in `regions[]`.
type ConsentSettingReplaceResponseDefault struct {
	// Per-category default config for this rule. Every category defined in the
	// top-level `categories[].value` should have an entry here.
	Categories []ConsentSettingReplaceResponseDefaultCategory `json:"categories" api:"required"`
	// BCP 47 default language for this rule. Must have a matching entry in
	// `translations`. Examples: "en", "en-US", "es", "de".
	Language string `json:"language" api:"required"`
	// opt_in: scripts blocked until user accepts (GDPR style). opt_out: scripts run by
	// default until user rejects (CCPA style).
	//
	// Any of "opt_in", "opt_out".
	Mode string `json:"mode" api:"required"`
	// All UI copy, keyed by language. Must include an entry whose `language` matches
	// the rule's `language` field.
	Translations []ConsentSettingReplaceResponseDefaultTranslation `json:"translations" api:"required"`
	// When true, scripts not classified by services[] are blocked until the user opts
	// in.
	AutoblockUnknown bool `json:"autoblockUnknown" api:"nullable"`
	// When true, the consent modal auto-opens on page load.
	AutoShow bool `json:"autoShow" api:"nullable"`
	// Threshold config for autoShowDismissMode (page count or seconds).
	AutoShowDismissConfig ConsentSettingReplaceResponseDefaultAutoShowDismissConfig `json:"autoShowDismissConfig" api:"nullable"`
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	//
	// Any of "after_pages", "after_seconds", "never".
	AutoShowDismissMode string `json:"autoShowDismissMode" api:"nullable"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction bool `json:"disablePageInteraction" api:"nullable"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions ConsentSettingReplaceResponseDefaultGuiOptions `json:"guiOptions" api:"nullable"`
	// When true, the modal is suppressed for known bot user agents.
	HideFromBots bool `json:"hideFromBots" api:"nullable"`
	// When true, the per-service list (services[]) is rendered inside the preferences
	// modal.
	ShowVendorsInPreferences bool `json:"showVendorsInPreferences" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories               respjson.Field
		Language                 respjson.Field
		Mode                     respjson.Field
		Translations             respjson.Field
		AutoblockUnknown         respjson.Field
		AutoShow                 respjson.Field
		AutoShowDismissConfig    respjson.Field
		AutoShowDismissMode      respjson.Field
		DisablePageInteraction   respjson.Field
		GuiOptions               respjson.Field
		HideFromBots             respjson.Field
		ShowVendorsInPreferences respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseDefault) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingReplaceResponseDefault) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingReplaceResponseDefaultCategory struct {
	// Category value (matches `categories[].value`) this entry configures.
	Key   string                                            `json:"key" api:"required"`
	Value ConsentSettingReplaceResponseDefaultCategoryValue `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseDefaultCategory) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingReplaceResponseDefaultCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingReplaceResponseDefaultCategoryValue struct {
	// Whether this category is on by default before the user interacts.
	Enabled bool `json:"enabled" api:"required"`
	// When true, this category defaults off if the browser sends Sec-GPC: 1.
	AutoDisableOnGpc bool `json:"autoDisableOnGPC" api:"nullable"`
	// When true, the user cannot toggle this category in the preferences modal.
	ReadOnly bool `json:"readOnly" api:"nullable"`
	// When true, the page reloads after this category is toggled so newly-allowed
	// scripts can run.
	ReloadPage bool `json:"reloadPage" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled          respjson.Field
		AutoDisableOnGpc respjson.Field
		ReadOnly         respjson.Field
		ReloadPage       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseDefaultCategoryValue) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingReplaceResponseDefaultCategoryValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingReplaceResponseDefaultTranslation struct {
	// BCP 47 language tag identifying which translation this entry provides. Examples:
	// "en", "en-US", "es", "fr-CA". The default rule's `language` must appear here.
	Language string                                               `json:"language" api:"required"`
	Value    ConsentSettingReplaceResponseDefaultTranslationValue `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Language    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseDefaultTranslation) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingReplaceResponseDefaultTranslation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingReplaceResponseDefaultTranslationValue struct {
	// Translated copy for the initial consent modal.
	ConsentModal ConsentSettingReplaceResponseDefaultTranslationValueConsentModal `json:"consentModal" api:"nullable"`
	// Translated copy for the preferences modal.
	PreferencesModal ConsentSettingReplaceResponseDefaultTranslationValuePreferencesModal `json:"preferencesModal" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConsentModal     respjson.Field
		PreferencesModal respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseDefaultTranslationValue) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingReplaceResponseDefaultTranslationValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the initial consent modal.
type ConsentSettingReplaceResponseDefaultTranslationValueConsentModal struct {
	// Label for the "accept all" button.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Body text shown in the consent modal.
	Description string `json:"description" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the button that opens the preferences modal.
	ShowPreferencesBtn string `json:"showPreferencesBtn" api:"required"`
	// Heading shown at the top of the consent modal.
	Title string `json:"title" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel string `json:"closeIconLabel" api:"nullable"`
	// Optional footer HTML/text shown below the body.
	Footer string `json:"footer" api:"nullable"`
	// Optional text shown when the browser is sending Sec-GPC: 1.
	GpcNotification string `json:"gpcNotification" api:"nullable"`
	// Link label for the privacy policy (defaults to "Privacy Policy").
	PrivacyPolicyLabel string `json:"privacyPolicyLabel" api:"nullable"`
	// Privacy policy link. Accepts an absolute URL (https://...) or a relative path
	// (/legal/privacy) so one config can run across multiple domains. javascript:,
	// data:, and vbscript: schemes are rejected.
	PrivacyPolicyURL string `json:"privacyPolicyUrl" api:"nullable"`
	// Optional message shown when the consent revision has bumped (re-prompt copy).
	RevisionMessage string `json:"revisionMessage" api:"nullable"`
	// Link label for terms of service (defaults to "Terms of Service").
	TermsOfServiceLabel string `json:"termsOfServiceLabel" api:"nullable"`
	// Terms of service link. Same scheme rules as privacyPolicyUrl: absolute or
	// relative path only.
	TermsOfServiceURL string `json:"termsOfServiceUrl" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptAllBtn        respjson.Field
		Description         respjson.Field
		RejectAllBtn        respjson.Field
		ShowPreferencesBtn  respjson.Field
		Title               respjson.Field
		CloseIconLabel      respjson.Field
		Footer              respjson.Field
		GpcNotification     respjson.Field
		PrivacyPolicyLabel  respjson.Field
		PrivacyPolicyURL    respjson.Field
		RevisionMessage     respjson.Field
		TermsOfServiceLabel respjson.Field
		TermsOfServiceURL   respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseDefaultTranslationValueConsentModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingReplaceResponseDefaultTranslationValueConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the preferences modal.
type ConsentSettingReplaceResponseDefaultTranslationValuePreferencesModal struct {
	// Label for the "accept all" button in the preferences modal.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel string `json:"closeIconLabel" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the "save my choices" button.
	SavePreferencesBtn string `json:"savePreferencesBtn" api:"required"`
	// Per-category sections plus a leading general section. The server backfills
	// sections for each defined category if absent.
	Sections []ConsentSettingReplaceResponseDefaultTranslationValuePreferencesModalSection `json:"sections" api:"required"`
	// Heading for the preferences modal.
	Title string `json:"title" api:"required"`
	// Optional collapsible info block rendered after the category sections. Renders
	// only when both title and body are non-empty. Use it for long-form text unrelated
	// to any specific category (e.g. a third-party advertisers disclosure).
	AdditionalInfo ConsentSettingReplaceResponseDefaultTranslationValuePreferencesModalAdditionalInfo `json:"additionalInfo" api:"nullable"`
	// Pluralizable label for the "N services" counter (e.g. "Service|Services").
	ServiceCounterLabel string `json:"serviceCounterLabel" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptAllBtn        respjson.Field
		CloseIconLabel      respjson.Field
		RejectAllBtn        respjson.Field
		SavePreferencesBtn  respjson.Field
		Sections            respjson.Field
		Title               respjson.Field
		AdditionalInfo      respjson.Field
		ServiceCounterLabel respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseDefaultTranslationValuePreferencesModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingReplaceResponseDefaultTranslationValuePreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingReplaceResponseDefaultTranslationValuePreferencesModalSection struct {
	// Body text for this preferences-modal section.
	Description string `json:"description" api:"required"`
	// Heading for this preferences-modal section.
	Title string `json:"title" api:"required"`
	// Optional disclosure table listing cookies/services in this section.
	CookieTable ConsentSettingReplaceResponseDefaultTranslationValuePreferencesModalSectionCookieTable `json:"cookieTable" api:"nullable"`
	// Category value this section configures. Use null/omit for the introductory
	// section that applies to no specific category.
	LinkedCategory string `json:"linkedCategory" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description    respjson.Field
		Title          respjson.Field
		CookieTable    respjson.Field
		LinkedCategory respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseDefaultTranslationValuePreferencesModalSection) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingReplaceResponseDefaultTranslationValuePreferencesModalSection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional disclosure table listing cookies/services in this section.
type ConsentSettingReplaceResponseDefaultTranslationValuePreferencesModalSectionCookieTable struct {
	// Rows of the cookie table — each row is a {key,value} pair (e.g. cookie name →
	// purpose).
	Body []ConsentSettingReplaceResponseDefaultTranslationValuePreferencesModalSectionCookieTableBody `json:"body" api:"required"`
	// Header cells. Same {key,value} shape as rows.
	Headers []ConsentSettingReplaceResponseDefaultTranslationValuePreferencesModalSectionCookieTableHeader `json:"headers" api:"required"`
	// Optional table caption.
	Caption string `json:"caption" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Headers     respjson.Field
		Caption     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseDefaultTranslationValuePreferencesModalSectionCookieTable) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingReplaceResponseDefaultTranslationValuePreferencesModalSectionCookieTable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingReplaceResponseDefaultTranslationValuePreferencesModalSectionCookieTableBody struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseDefaultTranslationValuePreferencesModalSectionCookieTableBody) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingReplaceResponseDefaultTranslationValuePreferencesModalSectionCookieTableBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingReplaceResponseDefaultTranslationValuePreferencesModalSectionCookieTableHeader struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseDefaultTranslationValuePreferencesModalSectionCookieTableHeader) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingReplaceResponseDefaultTranslationValuePreferencesModalSectionCookieTableHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional collapsible info block rendered after the category sections. Renders
// only when both title and body are non-empty. Use it for long-form text unrelated
// to any specific category (e.g. a third-party advertisers disclosure).
type ConsentSettingReplaceResponseDefaultTranslationValuePreferencesModalAdditionalInfo struct {
	// Body text for the additional-info block. Supports the rich-text markdown subset
	// used by section descriptions: line breaks, [text](url) links (https:// or
	// relative paths), and **bold**. HTML is stripped.
	Body string `json:"body" api:"required"`
	// Heading for the additional-info block.
	Title string `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseDefaultTranslationValuePreferencesModalAdditionalInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingReplaceResponseDefaultTranslationValuePreferencesModalAdditionalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Threshold config for autoShowDismissMode (page count or seconds).
type ConsentSettingReplaceResponseDefaultAutoShowDismissConfig struct {
	// Pages-viewed threshold when autoShowDismissMode = after_pages.
	PageCount int64 `json:"pageCount" api:"nullable"`
	// Seconds-on-page threshold when autoShowDismissMode = after_seconds.
	Seconds int64 `json:"seconds" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageCount   respjson.Field
		Seconds     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseDefaultAutoShowDismissConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingReplaceResponseDefaultAutoShowDismissConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the modals (layout/position/colors).
type ConsentSettingReplaceResponseDefaultGuiOptions struct {
	// Visual options for the initial consent modal.
	ConsentModal ConsentSettingReplaceResponseDefaultGuiOptionsConsentModal `json:"consentModal" api:"nullable"`
	// Global CSS variable overrides applied to both modals.
	CssVariables ConsentSettingReplaceResponseDefaultGuiOptionsCssVariables `json:"cssVariables" api:"nullable"`
	// Visual options for the per-category preferences modal.
	PreferencesModal ConsentSettingReplaceResponseDefaultGuiOptionsPreferencesModal `json:"preferencesModal" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConsentModal     respjson.Field
		CssVariables     respjson.Field
		PreferencesModal respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseDefaultGuiOptions) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingReplaceResponseDefaultGuiOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the initial consent modal.
type ConsentSettingReplaceResponseDefaultGuiOptionsConsentModal struct {
	// Which buttons to render in the consent modal and in what arrangement.
	//
	// Any of "AcceptAllNecessaryPreferences", "AcceptAllPreferences", "AcceptOnly",
	// "InformationOnly", "PreferencesOnly".
	ButtonLayout string `json:"buttonLayout" api:"nullable"`
	// When true, accept and reject buttons share the same visual weight (no
	// primary/secondary).
	EqualWeightButtons bool `json:"equalWeightButtons" api:"nullable"`
	// When true, swap accept/reject button order.
	FlipButtons bool `json:"flipButtons" api:"nullable"`
	// Modal layout style (box, bar, cloud, etc.).
	//
	// Any of "bar", "bar_inline", "box", "box_inline", "box_wide", "cloud",
	// "cloud_inline".
	Layout string `json:"layout" api:"nullable"`
	// Where the modal anchors on the viewport (e.g. bottom_left, top_right,
	// middle_center).
	//
	// Any of "bottom", "bottom_center", "bottom_left", "bottom_right", "middle",
	// "middle_center", "middle_left", "middle_right", "top", "top_center", "top_left",
	// "top_right".
	Position string `json:"position" api:"nullable"`
	// Show an X to dismiss the modal without choosing.
	ShowCloseIcon bool `json:"showCloseIcon" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonLayout       respjson.Field
		EqualWeightButtons respjson.Field
		FlipButtons        respjson.Field
		Layout             respjson.Field
		Position           respjson.Field
		ShowCloseIcon      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseDefaultGuiOptionsConsentModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingReplaceResponseDefaultGuiOptionsConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Global CSS variable overrides applied to both modals.
type ConsentSettingReplaceResponseDefaultGuiOptionsCssVariables struct {
	// CSS length, e.g. "4px" or "0.25rem".
	ButtonBorderRadius string `json:"buttonBorderRadius" api:"nullable"`
	// CSS color for the per-category block background in the preferences modal.
	CookieCategoryBlockBg string `json:"cookieCategoryBlockBg" api:"nullable"`
	// CSS color for the per-category block background on hover.
	CookieCategoryBlockHoverBg string `json:"cookieCategoryBlockHoverBg" api:"nullable"`
	// CSS color for the modal footer background.
	FooterBg string `json:"footerBg" api:"nullable"`
	// CSS color for footer text.
	FooterColor string `json:"footerColor" api:"nullable"`
	// CSS color for footer links.
	FooterLinkColor string `json:"footerLinkColor" api:"nullable"`
	// CSS color for footer links on hover.
	FooterLinkHoverColor string `json:"footerLinkHoverColor" api:"nullable"`
	// CSS color for the modal background.
	ModalBg string `json:"modalBg" api:"nullable"`
	// CSS length for modal corner radius.
	ModalBorderRadius string `json:"modalBorderRadius" api:"nullable"`
	// CSS color for the primary button background.
	PrimaryButtonBg string `json:"primaryButtonBg" api:"nullable"`
	// CSS color for primary button text.
	PrimaryButtonColor string `json:"primaryButtonColor" api:"nullable"`
	// CSS color for the primary button background on hover.
	PrimaryButtonHoverBg string `json:"primaryButtonHoverBg" api:"nullable"`
	// CSS color for primary button text on hover.
	PrimaryButtonHoverColor string `json:"primaryButtonHoverColor" api:"nullable"`
	// CSS color for primary text in the modal body.
	PrimaryTextColor string `json:"primaryTextColor" api:"nullable"`
	// CSS color for the secondary button background.
	SecondaryButtonBg string `json:"secondaryButtonBg" api:"nullable"`
	// CSS color for secondary button text.
	SecondaryButtonColor string `json:"secondaryButtonColor" api:"nullable"`
	// CSS color for secondary button background on hover.
	SecondaryButtonHoverBg string `json:"secondaryButtonHoverBg" api:"nullable"`
	// CSS color for secondary button text on hover.
	SecondaryButtonHoverColor string `json:"secondaryButtonHoverColor" api:"nullable"`
	// CSS color for secondary text in the modal body.
	SecondaryTextColor string `json:"secondaryTextColor" api:"nullable"`
	// CSS color for the "on" state of category toggles.
	ToggleOnBg string `json:"toggleOnBg" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonBorderRadius         respjson.Field
		CookieCategoryBlockBg      respjson.Field
		CookieCategoryBlockHoverBg respjson.Field
		FooterBg                   respjson.Field
		FooterColor                respjson.Field
		FooterLinkColor            respjson.Field
		FooterLinkHoverColor       respjson.Field
		ModalBg                    respjson.Field
		ModalBorderRadius          respjson.Field
		PrimaryButtonBg            respjson.Field
		PrimaryButtonColor         respjson.Field
		PrimaryButtonHoverBg       respjson.Field
		PrimaryButtonHoverColor    respjson.Field
		PrimaryTextColor           respjson.Field
		SecondaryButtonBg          respjson.Field
		SecondaryButtonColor       respjson.Field
		SecondaryButtonHoverBg     respjson.Field
		SecondaryButtonHoverColor  respjson.Field
		SecondaryTextColor         respjson.Field
		ToggleOnBg                 respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseDefaultGuiOptionsCssVariables) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingReplaceResponseDefaultGuiOptionsCssVariables) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the per-category preferences modal.
type ConsentSettingReplaceResponseDefaultGuiOptionsPreferencesModal struct {
	// Button layout for the preferences modal.
	//
	// Any of "AcceptAllRejectAllSave", "AcceptAllSave".
	ButtonLayout string `json:"buttonLayout" api:"nullable"`
	// When true, accept/reject share the same visual weight.
	EqualWeightButtons bool `json:"equalWeightButtons" api:"nullable"`
	// When true, swap accept/reject button order.
	FlipButtons bool `json:"flipButtons" api:"nullable"`
	// Preferences modal layout (box, bar, bar_wide).
	//
	// Any of "bar", "bar_wide", "box".
	Layout string `json:"layout" api:"nullable"`
	// Preferences modal anchor (left, right).
	//
	// Any of "left", "right".
	Position string `json:"position" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonLayout       respjson.Field
		EqualWeightButtons respjson.Field
		FlipButtons        respjson.Field
		Layout             respjson.Field
		Position           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseDefaultGuiOptionsPreferencesModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingReplaceResponseDefaultGuiOptionsPreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingReplaceResponseRegion struct {
	// Region this rule applies to. Use ISO 3166-1 alpha-2 country code ("US", "DE",
	// "BR") or country-subdivision code ("US-CA", "GB-ENG", "CA-ON"). Each region code
	// may appear in only one rule across `regions[]`.
	RegionCode string                                  `json:"regionCode" api:"required"`
	Rule       ConsentSettingReplaceResponseRegionRule `json:"rule" api:"required"`
	// Other region codes that should reuse this rule. Same code-format rules as
	// `regionCode`. Cannot include `regionCode` itself, cannot duplicate, cannot
	// overlap with another rule's regions.
	AdditionalRegions []string `json:"additionalRegions" api:"nullable"`
	// Optional label for this override group, e.g. "CCPA States". Shown in the
	// dashboard instead of the region codes. Pass null to clear it and fall back to
	// the codes.
	Name string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RegionCode        respjson.Field
		Rule              respjson.Field
		AdditionalRegions respjson.Field
		Name              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseRegion) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingReplaceResponseRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingReplaceResponseRegionRule struct {
	// Per-category default config for this rule. Every category defined in the
	// top-level `categories[].value` should have an entry here.
	Categories []ConsentSettingReplaceResponseRegionRuleCategory `json:"categories" api:"required"`
	// BCP 47 default language for this rule. Must have a matching entry in
	// `translations`. Examples: "en", "en-US", "es", "de".
	Language string `json:"language" api:"required"`
	// opt_in: scripts blocked until user accepts (GDPR style). opt_out: scripts run by
	// default until user rejects (CCPA style).
	//
	// Any of "opt_in", "opt_out".
	Mode string `json:"mode" api:"required"`
	// All UI copy, keyed by language. Must include an entry whose `language` matches
	// the rule's `language` field.
	Translations []ConsentSettingReplaceResponseRegionRuleTranslation `json:"translations" api:"required"`
	// When true, scripts not classified by services[] are blocked until the user opts
	// in.
	AutoblockUnknown bool `json:"autoblockUnknown" api:"nullable"`
	// When true, the consent modal auto-opens on page load.
	AutoShow bool `json:"autoShow" api:"nullable"`
	// Threshold config for autoShowDismissMode (page count or seconds).
	AutoShowDismissConfig ConsentSettingReplaceResponseRegionRuleAutoShowDismissConfig `json:"autoShowDismissConfig" api:"nullable"`
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	//
	// Any of "after_pages", "after_seconds", "never".
	AutoShowDismissMode string `json:"autoShowDismissMode" api:"nullable"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction bool `json:"disablePageInteraction" api:"nullable"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions ConsentSettingReplaceResponseRegionRuleGuiOptions `json:"guiOptions" api:"nullable"`
	// When true, the modal is suppressed for known bot user agents.
	HideFromBots bool `json:"hideFromBots" api:"nullable"`
	// When true, the per-service list (services[]) is rendered inside the preferences
	// modal.
	ShowVendorsInPreferences bool `json:"showVendorsInPreferences" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories               respjson.Field
		Language                 respjson.Field
		Mode                     respjson.Field
		Translations             respjson.Field
		AutoblockUnknown         respjson.Field
		AutoShow                 respjson.Field
		AutoShowDismissConfig    respjson.Field
		AutoShowDismissMode      respjson.Field
		DisablePageInteraction   respjson.Field
		GuiOptions               respjson.Field
		HideFromBots             respjson.Field
		ShowVendorsInPreferences respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseRegionRule) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingReplaceResponseRegionRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingReplaceResponseRegionRuleCategory struct {
	// Category value (matches `categories[].value`) this entry configures.
	Key   string                                               `json:"key" api:"required"`
	Value ConsentSettingReplaceResponseRegionRuleCategoryValue `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseRegionRuleCategory) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingReplaceResponseRegionRuleCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingReplaceResponseRegionRuleCategoryValue struct {
	// Whether this category is on by default before the user interacts.
	Enabled bool `json:"enabled" api:"required"`
	// When true, this category defaults off if the browser sends Sec-GPC: 1.
	AutoDisableOnGpc bool `json:"autoDisableOnGPC" api:"nullable"`
	// When true, the user cannot toggle this category in the preferences modal.
	ReadOnly bool `json:"readOnly" api:"nullable"`
	// When true, the page reloads after this category is toggled so newly-allowed
	// scripts can run.
	ReloadPage bool `json:"reloadPage" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled          respjson.Field
		AutoDisableOnGpc respjson.Field
		ReadOnly         respjson.Field
		ReloadPage       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseRegionRuleCategoryValue) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingReplaceResponseRegionRuleCategoryValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingReplaceResponseRegionRuleTranslation struct {
	// BCP 47 language tag identifying which translation this entry provides. Examples:
	// "en", "en-US", "es", "fr-CA". The default rule's `language` must appear here.
	Language string                                                  `json:"language" api:"required"`
	Value    ConsentSettingReplaceResponseRegionRuleTranslationValue `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Language    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseRegionRuleTranslation) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingReplaceResponseRegionRuleTranslation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingReplaceResponseRegionRuleTranslationValue struct {
	// Translated copy for the initial consent modal.
	ConsentModal ConsentSettingReplaceResponseRegionRuleTranslationValueConsentModal `json:"consentModal" api:"nullable"`
	// Translated copy for the preferences modal.
	PreferencesModal ConsentSettingReplaceResponseRegionRuleTranslationValuePreferencesModal `json:"preferencesModal" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConsentModal     respjson.Field
		PreferencesModal respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseRegionRuleTranslationValue) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingReplaceResponseRegionRuleTranslationValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the initial consent modal.
type ConsentSettingReplaceResponseRegionRuleTranslationValueConsentModal struct {
	// Label for the "accept all" button.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Body text shown in the consent modal.
	Description string `json:"description" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the button that opens the preferences modal.
	ShowPreferencesBtn string `json:"showPreferencesBtn" api:"required"`
	// Heading shown at the top of the consent modal.
	Title string `json:"title" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel string `json:"closeIconLabel" api:"nullable"`
	// Optional footer HTML/text shown below the body.
	Footer string `json:"footer" api:"nullable"`
	// Optional text shown when the browser is sending Sec-GPC: 1.
	GpcNotification string `json:"gpcNotification" api:"nullable"`
	// Link label for the privacy policy (defaults to "Privacy Policy").
	PrivacyPolicyLabel string `json:"privacyPolicyLabel" api:"nullable"`
	// Privacy policy link. Accepts an absolute URL (https://...) or a relative path
	// (/legal/privacy) so one config can run across multiple domains. javascript:,
	// data:, and vbscript: schemes are rejected.
	PrivacyPolicyURL string `json:"privacyPolicyUrl" api:"nullable"`
	// Optional message shown when the consent revision has bumped (re-prompt copy).
	RevisionMessage string `json:"revisionMessage" api:"nullable"`
	// Link label for terms of service (defaults to "Terms of Service").
	TermsOfServiceLabel string `json:"termsOfServiceLabel" api:"nullable"`
	// Terms of service link. Same scheme rules as privacyPolicyUrl: absolute or
	// relative path only.
	TermsOfServiceURL string `json:"termsOfServiceUrl" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptAllBtn        respjson.Field
		Description         respjson.Field
		RejectAllBtn        respjson.Field
		ShowPreferencesBtn  respjson.Field
		Title               respjson.Field
		CloseIconLabel      respjson.Field
		Footer              respjson.Field
		GpcNotification     respjson.Field
		PrivacyPolicyLabel  respjson.Field
		PrivacyPolicyURL    respjson.Field
		RevisionMessage     respjson.Field
		TermsOfServiceLabel respjson.Field
		TermsOfServiceURL   respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseRegionRuleTranslationValueConsentModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingReplaceResponseRegionRuleTranslationValueConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the preferences modal.
type ConsentSettingReplaceResponseRegionRuleTranslationValuePreferencesModal struct {
	// Label for the "accept all" button in the preferences modal.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel string `json:"closeIconLabel" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the "save my choices" button.
	SavePreferencesBtn string `json:"savePreferencesBtn" api:"required"`
	// Per-category sections plus a leading general section. The server backfills
	// sections for each defined category if absent.
	Sections []ConsentSettingReplaceResponseRegionRuleTranslationValuePreferencesModalSection `json:"sections" api:"required"`
	// Heading for the preferences modal.
	Title string `json:"title" api:"required"`
	// Optional collapsible info block rendered after the category sections. Renders
	// only when both title and body are non-empty. Use it for long-form text unrelated
	// to any specific category (e.g. a third-party advertisers disclosure).
	AdditionalInfo ConsentSettingReplaceResponseRegionRuleTranslationValuePreferencesModalAdditionalInfo `json:"additionalInfo" api:"nullable"`
	// Pluralizable label for the "N services" counter (e.g. "Service|Services").
	ServiceCounterLabel string `json:"serviceCounterLabel" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptAllBtn        respjson.Field
		CloseIconLabel      respjson.Field
		RejectAllBtn        respjson.Field
		SavePreferencesBtn  respjson.Field
		Sections            respjson.Field
		Title               respjson.Field
		AdditionalInfo      respjson.Field
		ServiceCounterLabel respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseRegionRuleTranslationValuePreferencesModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingReplaceResponseRegionRuleTranslationValuePreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingReplaceResponseRegionRuleTranslationValuePreferencesModalSection struct {
	// Body text for this preferences-modal section.
	Description string `json:"description" api:"required"`
	// Heading for this preferences-modal section.
	Title string `json:"title" api:"required"`
	// Optional disclosure table listing cookies/services in this section.
	CookieTable ConsentSettingReplaceResponseRegionRuleTranslationValuePreferencesModalSectionCookieTable `json:"cookieTable" api:"nullable"`
	// Category value this section configures. Use null/omit for the introductory
	// section that applies to no specific category.
	LinkedCategory string `json:"linkedCategory" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description    respjson.Field
		Title          respjson.Field
		CookieTable    respjson.Field
		LinkedCategory respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseRegionRuleTranslationValuePreferencesModalSection) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingReplaceResponseRegionRuleTranslationValuePreferencesModalSection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional disclosure table listing cookies/services in this section.
type ConsentSettingReplaceResponseRegionRuleTranslationValuePreferencesModalSectionCookieTable struct {
	// Rows of the cookie table — each row is a {key,value} pair (e.g. cookie name →
	// purpose).
	Body []ConsentSettingReplaceResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableBody `json:"body" api:"required"`
	// Header cells. Same {key,value} shape as rows.
	Headers []ConsentSettingReplaceResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader `json:"headers" api:"required"`
	// Optional table caption.
	Caption string `json:"caption" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Headers     respjson.Field
		Caption     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseRegionRuleTranslationValuePreferencesModalSectionCookieTable) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingReplaceResponseRegionRuleTranslationValuePreferencesModalSectionCookieTable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingReplaceResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableBody struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableBody) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingReplaceResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingReplaceResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingReplaceResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional collapsible info block rendered after the category sections. Renders
// only when both title and body are non-empty. Use it for long-form text unrelated
// to any specific category (e.g. a third-party advertisers disclosure).
type ConsentSettingReplaceResponseRegionRuleTranslationValuePreferencesModalAdditionalInfo struct {
	// Body text for the additional-info block. Supports the rich-text markdown subset
	// used by section descriptions: line breaks, [text](url) links (https:// or
	// relative paths), and **bold**. HTML is stripped.
	Body string `json:"body" api:"required"`
	// Heading for the additional-info block.
	Title string `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseRegionRuleTranslationValuePreferencesModalAdditionalInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingReplaceResponseRegionRuleTranslationValuePreferencesModalAdditionalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Threshold config for autoShowDismissMode (page count or seconds).
type ConsentSettingReplaceResponseRegionRuleAutoShowDismissConfig struct {
	// Pages-viewed threshold when autoShowDismissMode = after_pages.
	PageCount int64 `json:"pageCount" api:"nullable"`
	// Seconds-on-page threshold when autoShowDismissMode = after_seconds.
	Seconds int64 `json:"seconds" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageCount   respjson.Field
		Seconds     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseRegionRuleAutoShowDismissConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingReplaceResponseRegionRuleAutoShowDismissConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the modals (layout/position/colors).
type ConsentSettingReplaceResponseRegionRuleGuiOptions struct {
	// Visual options for the initial consent modal.
	ConsentModal ConsentSettingReplaceResponseRegionRuleGuiOptionsConsentModal `json:"consentModal" api:"nullable"`
	// Global CSS variable overrides applied to both modals.
	CssVariables ConsentSettingReplaceResponseRegionRuleGuiOptionsCssVariables `json:"cssVariables" api:"nullable"`
	// Visual options for the per-category preferences modal.
	PreferencesModal ConsentSettingReplaceResponseRegionRuleGuiOptionsPreferencesModal `json:"preferencesModal" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConsentModal     respjson.Field
		CssVariables     respjson.Field
		PreferencesModal respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseRegionRuleGuiOptions) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingReplaceResponseRegionRuleGuiOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the initial consent modal.
type ConsentSettingReplaceResponseRegionRuleGuiOptionsConsentModal struct {
	// Which buttons to render in the consent modal and in what arrangement.
	//
	// Any of "AcceptAllNecessaryPreferences", "AcceptAllPreferences", "AcceptOnly",
	// "InformationOnly", "PreferencesOnly".
	ButtonLayout string `json:"buttonLayout" api:"nullable"`
	// When true, accept and reject buttons share the same visual weight (no
	// primary/secondary).
	EqualWeightButtons bool `json:"equalWeightButtons" api:"nullable"`
	// When true, swap accept/reject button order.
	FlipButtons bool `json:"flipButtons" api:"nullable"`
	// Modal layout style (box, bar, cloud, etc.).
	//
	// Any of "bar", "bar_inline", "box", "box_inline", "box_wide", "cloud",
	// "cloud_inline".
	Layout string `json:"layout" api:"nullable"`
	// Where the modal anchors on the viewport (e.g. bottom_left, top_right,
	// middle_center).
	//
	// Any of "bottom", "bottom_center", "bottom_left", "bottom_right", "middle",
	// "middle_center", "middle_left", "middle_right", "top", "top_center", "top_left",
	// "top_right".
	Position string `json:"position" api:"nullable"`
	// Show an X to dismiss the modal without choosing.
	ShowCloseIcon bool `json:"showCloseIcon" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonLayout       respjson.Field
		EqualWeightButtons respjson.Field
		FlipButtons        respjson.Field
		Layout             respjson.Field
		Position           respjson.Field
		ShowCloseIcon      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseRegionRuleGuiOptionsConsentModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingReplaceResponseRegionRuleGuiOptionsConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Global CSS variable overrides applied to both modals.
type ConsentSettingReplaceResponseRegionRuleGuiOptionsCssVariables struct {
	// CSS length, e.g. "4px" or "0.25rem".
	ButtonBorderRadius string `json:"buttonBorderRadius" api:"nullable"`
	// CSS color for the per-category block background in the preferences modal.
	CookieCategoryBlockBg string `json:"cookieCategoryBlockBg" api:"nullable"`
	// CSS color for the per-category block background on hover.
	CookieCategoryBlockHoverBg string `json:"cookieCategoryBlockHoverBg" api:"nullable"`
	// CSS color for the modal footer background.
	FooterBg string `json:"footerBg" api:"nullable"`
	// CSS color for footer text.
	FooterColor string `json:"footerColor" api:"nullable"`
	// CSS color for footer links.
	FooterLinkColor string `json:"footerLinkColor" api:"nullable"`
	// CSS color for footer links on hover.
	FooterLinkHoverColor string `json:"footerLinkHoverColor" api:"nullable"`
	// CSS color for the modal background.
	ModalBg string `json:"modalBg" api:"nullable"`
	// CSS length for modal corner radius.
	ModalBorderRadius string `json:"modalBorderRadius" api:"nullable"`
	// CSS color for the primary button background.
	PrimaryButtonBg string `json:"primaryButtonBg" api:"nullable"`
	// CSS color for primary button text.
	PrimaryButtonColor string `json:"primaryButtonColor" api:"nullable"`
	// CSS color for the primary button background on hover.
	PrimaryButtonHoverBg string `json:"primaryButtonHoverBg" api:"nullable"`
	// CSS color for primary button text on hover.
	PrimaryButtonHoverColor string `json:"primaryButtonHoverColor" api:"nullable"`
	// CSS color for primary text in the modal body.
	PrimaryTextColor string `json:"primaryTextColor" api:"nullable"`
	// CSS color for the secondary button background.
	SecondaryButtonBg string `json:"secondaryButtonBg" api:"nullable"`
	// CSS color for secondary button text.
	SecondaryButtonColor string `json:"secondaryButtonColor" api:"nullable"`
	// CSS color for secondary button background on hover.
	SecondaryButtonHoverBg string `json:"secondaryButtonHoverBg" api:"nullable"`
	// CSS color for secondary button text on hover.
	SecondaryButtonHoverColor string `json:"secondaryButtonHoverColor" api:"nullable"`
	// CSS color for secondary text in the modal body.
	SecondaryTextColor string `json:"secondaryTextColor" api:"nullable"`
	// CSS color for the "on" state of category toggles.
	ToggleOnBg string `json:"toggleOnBg" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonBorderRadius         respjson.Field
		CookieCategoryBlockBg      respjson.Field
		CookieCategoryBlockHoverBg respjson.Field
		FooterBg                   respjson.Field
		FooterColor                respjson.Field
		FooterLinkColor            respjson.Field
		FooterLinkHoverColor       respjson.Field
		ModalBg                    respjson.Field
		ModalBorderRadius          respjson.Field
		PrimaryButtonBg            respjson.Field
		PrimaryButtonColor         respjson.Field
		PrimaryButtonHoverBg       respjson.Field
		PrimaryButtonHoverColor    respjson.Field
		PrimaryTextColor           respjson.Field
		SecondaryButtonBg          respjson.Field
		SecondaryButtonColor       respjson.Field
		SecondaryButtonHoverBg     respjson.Field
		SecondaryButtonHoverColor  respjson.Field
		SecondaryTextColor         respjson.Field
		ToggleOnBg                 respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseRegionRuleGuiOptionsCssVariables) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingReplaceResponseRegionRuleGuiOptionsCssVariables) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the per-category preferences modal.
type ConsentSettingReplaceResponseRegionRuleGuiOptionsPreferencesModal struct {
	// Button layout for the preferences modal.
	//
	// Any of "AcceptAllRejectAllSave", "AcceptAllSave".
	ButtonLayout string `json:"buttonLayout" api:"nullable"`
	// When true, accept/reject share the same visual weight.
	EqualWeightButtons bool `json:"equalWeightButtons" api:"nullable"`
	// When true, swap accept/reject button order.
	FlipButtons bool `json:"flipButtons" api:"nullable"`
	// Preferences modal layout (box, bar, bar_wide).
	//
	// Any of "bar", "bar_wide", "box".
	Layout string `json:"layout" api:"nullable"`
	// Preferences modal anchor (left, right).
	//
	// Any of "left", "right".
	Position string `json:"position" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonLayout       respjson.Field
		EqualWeightButtons respjson.Field
		FlipButtons        respjson.Field
		Layout             respjson.Field
		Position           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseRegionRuleGuiOptionsPreferencesModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingReplaceResponseRegionRuleGuiOptionsPreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingReplaceResponseService struct {
	// Internal notes shown to admins in the dashboard. Not user-facing.
	InternalNotes string `json:"internalNotes" api:"required"`
	// Display name for this service in the preferences modal.
	Label string `json:"label" api:"required"`
	// Extra category values this service belongs to. Each must match a
	// `categories[].value`.
	AdditionalCategories []string `json:"additionalCategories" api:"nullable"`
	// Primary category value this service belongs to. Must match one of the top-level
	// `categories[].value` entries.
	Category string `json:"category" api:"nullable"`
	// Domains/paths this service matches. Patterns matching the CMP's own scripts
	// (e.g. cdn.oursprivacy.com/cmp-init) are rejected to prevent the CMP blocking
	// itself — use a more specific path like cdn.oursprivacy.com/main.js to block a
	// specific script.
	DomainPatterns []string `json:"domainPatterns" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InternalNotes        respjson.Field
		Label                respjson.Field
		AdditionalCategories respjson.Field
		Category             respjson.Field
		DomainPatterns       respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingReplaceResponseService) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingReplaceResponseService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enabled serves the CMP on its allowed domains and requires at least one.
// Disabled does not serve the CMP and may have no allowed domains.
type ConsentSettingReplaceResponseStatus string

const (
	ConsentSettingReplaceResponseStatusDisabled ConsentSettingReplaceResponseStatus = "Disabled"
	ConsentSettingReplaceResponseStatusEnabled  ConsentSettingReplaceResponseStatus = "Enabled"
)

type ConsentSettingUpdateResponse struct {
	// Server-assigned UUID for this consent settings record.
	ID string `json:"id" api:"required"`
	// Top-level consent categories (e.g. necessary / analytics / advertising). Server
	// re-stamps `priority` to 0..N on write.
	Categories []ConsentSettingUpdateResponseCategory `json:"categories" api:"required"`
	// ISO 8601 timestamp when the record was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Default rule used when the user is not in any region listed in `regions[]`.
	Default ConsentSettingUpdateResponseDefault `json:"default" api:"required"`
	// Discriminator for the entity type. Always "consentSettings".
	Kind string `json:"kind" api:"required"`
	// Human-readable name shown in the dashboard.
	Name string `json:"name" api:"required"`
	// Per-region rule overrides. The first rule whose `regionCode`/`additionalRegions`
	// includes the user's region wins; otherwise `default` applies.
	Regions []ConsentSettingUpdateResponseRegion `json:"regions" api:"required"`
	// Per-service entries powering "show vendors" and category-aware blocking.
	Services []ConsentSettingUpdateResponseService `json:"services" api:"required"`
	// Enabled serves the CMP on its allowed domains and requires at least one.
	// Disabled does not serve the CMP and may have no allowed domains.
	//
	// Any of "Disabled", "Enabled".
	Status ConsentSettingUpdateResponseStatus `json:"status" api:"required"`
	// Name of the cookie that stores the user's consent state. Defaults to
	// "op_consent".
	ConsentCookieName string `json:"consentCookieName" api:"nullable"`
	// Optional custom CDN domain for serving the CMP script (e.g.
	// consent.example.com).
	CustomDomain string `json:"customDomain" api:"nullable"`
	// Cookie/localStorage key for the visitor (device) ID. Must exactly match the web
	// SDK's cookie_names.device_id (default "ours_device_id") so consent and SDK
	// events resolve to the same visitor.
	DeviceIDCookieName string `json:"deviceIdCookieName" api:"nullable"`
	// Revision counter. Bump this to force users who already consented to see the
	// modal again (the SDK compares the persisted revision against this value).
	Revision float64 `json:"revision" api:"nullable"`
	// CSS class names that opt scripts out of consent blocking. Each entry must be a
	// single class token (no whitespace).
	SkipBlockingClassNames []string `json:"skipBlockingClassNames" api:"nullable"`
	// ISO 8601 timestamp of the last write. Null on a freshly created record.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// Pixel of the WebSource that this CMP is wired into. Setting this to a token that
	// is not a valid WebSource of yours is rejected; use null to clear the link.
	WebSDKToken string `json:"webSDKToken" api:"nullable"`
	// Domains where this CMP may run. An Enabled CMP requires at least one allowed
	// domain; with none configured, the CMP does not run. Matching domains may share
	// consent across subdomains.
	WhitelistDomains []string `json:"whitelistDomains" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		Categories             respjson.Field
		CreatedAt              respjson.Field
		Default                respjson.Field
		Kind                   respjson.Field
		Name                   respjson.Field
		Regions                respjson.Field
		Services               respjson.Field
		Status                 respjson.Field
		ConsentCookieName      respjson.Field
		CustomDomain           respjson.Field
		DeviceIDCookieName     respjson.Field
		Revision               respjson.Field
		SkipBlockingClassNames respjson.Field
		UpdatedAt              respjson.Field
		WebSDKToken            respjson.Field
		WhitelistDomains       respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingUpdateResponseCategory struct {
	// Human-readable label shown next to the toggle in the preferences modal.
	Label string `json:"label" api:"required"`
	// Sort key. Lower numbers render first. Server re-stamps to 0..N on write — send
	// any integer, gaps and duplicates are ironed out.
	Priority int64 `json:"priority" api:"required"`
	// Stable identifier referenced by services and translation sections.
	// Conventionally lowercase (e.g. "necessary", "analytics", "advertising").
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		Priority    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseCategory) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingUpdateResponseCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default rule used when the user is not in any region listed in `regions[]`.
type ConsentSettingUpdateResponseDefault struct {
	// Per-category default config for this rule. Every category defined in the
	// top-level `categories[].value` should have an entry here.
	Categories []ConsentSettingUpdateResponseDefaultCategory `json:"categories" api:"required"`
	// BCP 47 default language for this rule. Must have a matching entry in
	// `translations`. Examples: "en", "en-US", "es", "de".
	Language string `json:"language" api:"required"`
	// opt_in: scripts blocked until user accepts (GDPR style). opt_out: scripts run by
	// default until user rejects (CCPA style).
	//
	// Any of "opt_in", "opt_out".
	Mode string `json:"mode" api:"required"`
	// All UI copy, keyed by language. Must include an entry whose `language` matches
	// the rule's `language` field.
	Translations []ConsentSettingUpdateResponseDefaultTranslation `json:"translations" api:"required"`
	// When true, scripts not classified by services[] are blocked until the user opts
	// in.
	AutoblockUnknown bool `json:"autoblockUnknown" api:"nullable"`
	// When true, the consent modal auto-opens on page load.
	AutoShow bool `json:"autoShow" api:"nullable"`
	// Threshold config for autoShowDismissMode (page count or seconds).
	AutoShowDismissConfig ConsentSettingUpdateResponseDefaultAutoShowDismissConfig `json:"autoShowDismissConfig" api:"nullable"`
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	//
	// Any of "after_pages", "after_seconds", "never".
	AutoShowDismissMode string `json:"autoShowDismissMode" api:"nullable"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction bool `json:"disablePageInteraction" api:"nullable"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions ConsentSettingUpdateResponseDefaultGuiOptions `json:"guiOptions" api:"nullable"`
	// When true, the modal is suppressed for known bot user agents.
	HideFromBots bool `json:"hideFromBots" api:"nullable"`
	// When true, the per-service list (services[]) is rendered inside the preferences
	// modal.
	ShowVendorsInPreferences bool `json:"showVendorsInPreferences" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories               respjson.Field
		Language                 respjson.Field
		Mode                     respjson.Field
		Translations             respjson.Field
		AutoblockUnknown         respjson.Field
		AutoShow                 respjson.Field
		AutoShowDismissConfig    respjson.Field
		AutoShowDismissMode      respjson.Field
		DisablePageInteraction   respjson.Field
		GuiOptions               respjson.Field
		HideFromBots             respjson.Field
		ShowVendorsInPreferences respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseDefault) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingUpdateResponseDefault) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingUpdateResponseDefaultCategory struct {
	// Category value (matches `categories[].value`) this entry configures.
	Key   string                                           `json:"key" api:"required"`
	Value ConsentSettingUpdateResponseDefaultCategoryValue `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseDefaultCategory) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingUpdateResponseDefaultCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingUpdateResponseDefaultCategoryValue struct {
	// Whether this category is on by default before the user interacts.
	Enabled bool `json:"enabled" api:"required"`
	// When true, this category defaults off if the browser sends Sec-GPC: 1.
	AutoDisableOnGpc bool `json:"autoDisableOnGPC" api:"nullable"`
	// When true, the user cannot toggle this category in the preferences modal.
	ReadOnly bool `json:"readOnly" api:"nullable"`
	// When true, the page reloads after this category is toggled so newly-allowed
	// scripts can run.
	ReloadPage bool `json:"reloadPage" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled          respjson.Field
		AutoDisableOnGpc respjson.Field
		ReadOnly         respjson.Field
		ReloadPage       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseDefaultCategoryValue) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingUpdateResponseDefaultCategoryValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingUpdateResponseDefaultTranslation struct {
	// BCP 47 language tag identifying which translation this entry provides. Examples:
	// "en", "en-US", "es", "fr-CA". The default rule's `language` must appear here.
	Language string                                              `json:"language" api:"required"`
	Value    ConsentSettingUpdateResponseDefaultTranslationValue `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Language    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseDefaultTranslation) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingUpdateResponseDefaultTranslation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingUpdateResponseDefaultTranslationValue struct {
	// Translated copy for the initial consent modal.
	ConsentModal ConsentSettingUpdateResponseDefaultTranslationValueConsentModal `json:"consentModal" api:"nullable"`
	// Translated copy for the preferences modal.
	PreferencesModal ConsentSettingUpdateResponseDefaultTranslationValuePreferencesModal `json:"preferencesModal" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConsentModal     respjson.Field
		PreferencesModal respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseDefaultTranslationValue) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingUpdateResponseDefaultTranslationValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the initial consent modal.
type ConsentSettingUpdateResponseDefaultTranslationValueConsentModal struct {
	// Label for the "accept all" button.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Body text shown in the consent modal.
	Description string `json:"description" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the button that opens the preferences modal.
	ShowPreferencesBtn string `json:"showPreferencesBtn" api:"required"`
	// Heading shown at the top of the consent modal.
	Title string `json:"title" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel string `json:"closeIconLabel" api:"nullable"`
	// Optional footer HTML/text shown below the body.
	Footer string `json:"footer" api:"nullable"`
	// Optional text shown when the browser is sending Sec-GPC: 1.
	GpcNotification string `json:"gpcNotification" api:"nullable"`
	// Link label for the privacy policy (defaults to "Privacy Policy").
	PrivacyPolicyLabel string `json:"privacyPolicyLabel" api:"nullable"`
	// Privacy policy link. Accepts an absolute URL (https://...) or a relative path
	// (/legal/privacy) so one config can run across multiple domains. javascript:,
	// data:, and vbscript: schemes are rejected.
	PrivacyPolicyURL string `json:"privacyPolicyUrl" api:"nullable"`
	// Optional message shown when the consent revision has bumped (re-prompt copy).
	RevisionMessage string `json:"revisionMessage" api:"nullable"`
	// Link label for terms of service (defaults to "Terms of Service").
	TermsOfServiceLabel string `json:"termsOfServiceLabel" api:"nullable"`
	// Terms of service link. Same scheme rules as privacyPolicyUrl: absolute or
	// relative path only.
	TermsOfServiceURL string `json:"termsOfServiceUrl" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptAllBtn        respjson.Field
		Description         respjson.Field
		RejectAllBtn        respjson.Field
		ShowPreferencesBtn  respjson.Field
		Title               respjson.Field
		CloseIconLabel      respjson.Field
		Footer              respjson.Field
		GpcNotification     respjson.Field
		PrivacyPolicyLabel  respjson.Field
		PrivacyPolicyURL    respjson.Field
		RevisionMessage     respjson.Field
		TermsOfServiceLabel respjson.Field
		TermsOfServiceURL   respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseDefaultTranslationValueConsentModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingUpdateResponseDefaultTranslationValueConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the preferences modal.
type ConsentSettingUpdateResponseDefaultTranslationValuePreferencesModal struct {
	// Label for the "accept all" button in the preferences modal.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel string `json:"closeIconLabel" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the "save my choices" button.
	SavePreferencesBtn string `json:"savePreferencesBtn" api:"required"`
	// Per-category sections plus a leading general section. The server backfills
	// sections for each defined category if absent.
	Sections []ConsentSettingUpdateResponseDefaultTranslationValuePreferencesModalSection `json:"sections" api:"required"`
	// Heading for the preferences modal.
	Title string `json:"title" api:"required"`
	// Optional collapsible info block rendered after the category sections. Renders
	// only when both title and body are non-empty. Use it for long-form text unrelated
	// to any specific category (e.g. a third-party advertisers disclosure).
	AdditionalInfo ConsentSettingUpdateResponseDefaultTranslationValuePreferencesModalAdditionalInfo `json:"additionalInfo" api:"nullable"`
	// Pluralizable label for the "N services" counter (e.g. "Service|Services").
	ServiceCounterLabel string `json:"serviceCounterLabel" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptAllBtn        respjson.Field
		CloseIconLabel      respjson.Field
		RejectAllBtn        respjson.Field
		SavePreferencesBtn  respjson.Field
		Sections            respjson.Field
		Title               respjson.Field
		AdditionalInfo      respjson.Field
		ServiceCounterLabel respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseDefaultTranslationValuePreferencesModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingUpdateResponseDefaultTranslationValuePreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingUpdateResponseDefaultTranslationValuePreferencesModalSection struct {
	// Body text for this preferences-modal section.
	Description string `json:"description" api:"required"`
	// Heading for this preferences-modal section.
	Title string `json:"title" api:"required"`
	// Optional disclosure table listing cookies/services in this section.
	CookieTable ConsentSettingUpdateResponseDefaultTranslationValuePreferencesModalSectionCookieTable `json:"cookieTable" api:"nullable"`
	// Category value this section configures. Use null/omit for the introductory
	// section that applies to no specific category.
	LinkedCategory string `json:"linkedCategory" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description    respjson.Field
		Title          respjson.Field
		CookieTable    respjson.Field
		LinkedCategory respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseDefaultTranslationValuePreferencesModalSection) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingUpdateResponseDefaultTranslationValuePreferencesModalSection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional disclosure table listing cookies/services in this section.
type ConsentSettingUpdateResponseDefaultTranslationValuePreferencesModalSectionCookieTable struct {
	// Rows of the cookie table — each row is a {key,value} pair (e.g. cookie name →
	// purpose).
	Body []ConsentSettingUpdateResponseDefaultTranslationValuePreferencesModalSectionCookieTableBody `json:"body" api:"required"`
	// Header cells. Same {key,value} shape as rows.
	Headers []ConsentSettingUpdateResponseDefaultTranslationValuePreferencesModalSectionCookieTableHeader `json:"headers" api:"required"`
	// Optional table caption.
	Caption string `json:"caption" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Headers     respjson.Field
		Caption     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseDefaultTranslationValuePreferencesModalSectionCookieTable) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingUpdateResponseDefaultTranslationValuePreferencesModalSectionCookieTable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingUpdateResponseDefaultTranslationValuePreferencesModalSectionCookieTableBody struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseDefaultTranslationValuePreferencesModalSectionCookieTableBody) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingUpdateResponseDefaultTranslationValuePreferencesModalSectionCookieTableBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingUpdateResponseDefaultTranslationValuePreferencesModalSectionCookieTableHeader struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseDefaultTranslationValuePreferencesModalSectionCookieTableHeader) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingUpdateResponseDefaultTranslationValuePreferencesModalSectionCookieTableHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional collapsible info block rendered after the category sections. Renders
// only when both title and body are non-empty. Use it for long-form text unrelated
// to any specific category (e.g. a third-party advertisers disclosure).
type ConsentSettingUpdateResponseDefaultTranslationValuePreferencesModalAdditionalInfo struct {
	// Body text for the additional-info block. Supports the rich-text markdown subset
	// used by section descriptions: line breaks, [text](url) links (https:// or
	// relative paths), and **bold**. HTML is stripped.
	Body string `json:"body" api:"required"`
	// Heading for the additional-info block.
	Title string `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseDefaultTranslationValuePreferencesModalAdditionalInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingUpdateResponseDefaultTranslationValuePreferencesModalAdditionalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Threshold config for autoShowDismissMode (page count or seconds).
type ConsentSettingUpdateResponseDefaultAutoShowDismissConfig struct {
	// Pages-viewed threshold when autoShowDismissMode = after_pages.
	PageCount int64 `json:"pageCount" api:"nullable"`
	// Seconds-on-page threshold when autoShowDismissMode = after_seconds.
	Seconds int64 `json:"seconds" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageCount   respjson.Field
		Seconds     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseDefaultAutoShowDismissConfig) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingUpdateResponseDefaultAutoShowDismissConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the modals (layout/position/colors).
type ConsentSettingUpdateResponseDefaultGuiOptions struct {
	// Visual options for the initial consent modal.
	ConsentModal ConsentSettingUpdateResponseDefaultGuiOptionsConsentModal `json:"consentModal" api:"nullable"`
	// Global CSS variable overrides applied to both modals.
	CssVariables ConsentSettingUpdateResponseDefaultGuiOptionsCssVariables `json:"cssVariables" api:"nullable"`
	// Visual options for the per-category preferences modal.
	PreferencesModal ConsentSettingUpdateResponseDefaultGuiOptionsPreferencesModal `json:"preferencesModal" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConsentModal     respjson.Field
		CssVariables     respjson.Field
		PreferencesModal respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseDefaultGuiOptions) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingUpdateResponseDefaultGuiOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the initial consent modal.
type ConsentSettingUpdateResponseDefaultGuiOptionsConsentModal struct {
	// Which buttons to render in the consent modal and in what arrangement.
	//
	// Any of "AcceptAllNecessaryPreferences", "AcceptAllPreferences", "AcceptOnly",
	// "InformationOnly", "PreferencesOnly".
	ButtonLayout string `json:"buttonLayout" api:"nullable"`
	// When true, accept and reject buttons share the same visual weight (no
	// primary/secondary).
	EqualWeightButtons bool `json:"equalWeightButtons" api:"nullable"`
	// When true, swap accept/reject button order.
	FlipButtons bool `json:"flipButtons" api:"nullable"`
	// Modal layout style (box, bar, cloud, etc.).
	//
	// Any of "bar", "bar_inline", "box", "box_inline", "box_wide", "cloud",
	// "cloud_inline".
	Layout string `json:"layout" api:"nullable"`
	// Where the modal anchors on the viewport (e.g. bottom_left, top_right,
	// middle_center).
	//
	// Any of "bottom", "bottom_center", "bottom_left", "bottom_right", "middle",
	// "middle_center", "middle_left", "middle_right", "top", "top_center", "top_left",
	// "top_right".
	Position string `json:"position" api:"nullable"`
	// Show an X to dismiss the modal without choosing.
	ShowCloseIcon bool `json:"showCloseIcon" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonLayout       respjson.Field
		EqualWeightButtons respjson.Field
		FlipButtons        respjson.Field
		Layout             respjson.Field
		Position           respjson.Field
		ShowCloseIcon      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseDefaultGuiOptionsConsentModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingUpdateResponseDefaultGuiOptionsConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Global CSS variable overrides applied to both modals.
type ConsentSettingUpdateResponseDefaultGuiOptionsCssVariables struct {
	// CSS length, e.g. "4px" or "0.25rem".
	ButtonBorderRadius string `json:"buttonBorderRadius" api:"nullable"`
	// CSS color for the per-category block background in the preferences modal.
	CookieCategoryBlockBg string `json:"cookieCategoryBlockBg" api:"nullable"`
	// CSS color for the per-category block background on hover.
	CookieCategoryBlockHoverBg string `json:"cookieCategoryBlockHoverBg" api:"nullable"`
	// CSS color for the modal footer background.
	FooterBg string `json:"footerBg" api:"nullable"`
	// CSS color for footer text.
	FooterColor string `json:"footerColor" api:"nullable"`
	// CSS color for footer links.
	FooterLinkColor string `json:"footerLinkColor" api:"nullable"`
	// CSS color for footer links on hover.
	FooterLinkHoverColor string `json:"footerLinkHoverColor" api:"nullable"`
	// CSS color for the modal background.
	ModalBg string `json:"modalBg" api:"nullable"`
	// CSS length for modal corner radius.
	ModalBorderRadius string `json:"modalBorderRadius" api:"nullable"`
	// CSS color for the primary button background.
	PrimaryButtonBg string `json:"primaryButtonBg" api:"nullable"`
	// CSS color for primary button text.
	PrimaryButtonColor string `json:"primaryButtonColor" api:"nullable"`
	// CSS color for the primary button background on hover.
	PrimaryButtonHoverBg string `json:"primaryButtonHoverBg" api:"nullable"`
	// CSS color for primary button text on hover.
	PrimaryButtonHoverColor string `json:"primaryButtonHoverColor" api:"nullable"`
	// CSS color for primary text in the modal body.
	PrimaryTextColor string `json:"primaryTextColor" api:"nullable"`
	// CSS color for the secondary button background.
	SecondaryButtonBg string `json:"secondaryButtonBg" api:"nullable"`
	// CSS color for secondary button text.
	SecondaryButtonColor string `json:"secondaryButtonColor" api:"nullable"`
	// CSS color for secondary button background on hover.
	SecondaryButtonHoverBg string `json:"secondaryButtonHoverBg" api:"nullable"`
	// CSS color for secondary button text on hover.
	SecondaryButtonHoverColor string `json:"secondaryButtonHoverColor" api:"nullable"`
	// CSS color for secondary text in the modal body.
	SecondaryTextColor string `json:"secondaryTextColor" api:"nullable"`
	// CSS color for the "on" state of category toggles.
	ToggleOnBg string `json:"toggleOnBg" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonBorderRadius         respjson.Field
		CookieCategoryBlockBg      respjson.Field
		CookieCategoryBlockHoverBg respjson.Field
		FooterBg                   respjson.Field
		FooterColor                respjson.Field
		FooterLinkColor            respjson.Field
		FooterLinkHoverColor       respjson.Field
		ModalBg                    respjson.Field
		ModalBorderRadius          respjson.Field
		PrimaryButtonBg            respjson.Field
		PrimaryButtonColor         respjson.Field
		PrimaryButtonHoverBg       respjson.Field
		PrimaryButtonHoverColor    respjson.Field
		PrimaryTextColor           respjson.Field
		SecondaryButtonBg          respjson.Field
		SecondaryButtonColor       respjson.Field
		SecondaryButtonHoverBg     respjson.Field
		SecondaryButtonHoverColor  respjson.Field
		SecondaryTextColor         respjson.Field
		ToggleOnBg                 respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseDefaultGuiOptionsCssVariables) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingUpdateResponseDefaultGuiOptionsCssVariables) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the per-category preferences modal.
type ConsentSettingUpdateResponseDefaultGuiOptionsPreferencesModal struct {
	// Button layout for the preferences modal.
	//
	// Any of "AcceptAllRejectAllSave", "AcceptAllSave".
	ButtonLayout string `json:"buttonLayout" api:"nullable"`
	// When true, accept/reject share the same visual weight.
	EqualWeightButtons bool `json:"equalWeightButtons" api:"nullable"`
	// When true, swap accept/reject button order.
	FlipButtons bool `json:"flipButtons" api:"nullable"`
	// Preferences modal layout (box, bar, bar_wide).
	//
	// Any of "bar", "bar_wide", "box".
	Layout string `json:"layout" api:"nullable"`
	// Preferences modal anchor (left, right).
	//
	// Any of "left", "right".
	Position string `json:"position" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonLayout       respjson.Field
		EqualWeightButtons respjson.Field
		FlipButtons        respjson.Field
		Layout             respjson.Field
		Position           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseDefaultGuiOptionsPreferencesModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingUpdateResponseDefaultGuiOptionsPreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingUpdateResponseRegion struct {
	// Region this rule applies to. Use ISO 3166-1 alpha-2 country code ("US", "DE",
	// "BR") or country-subdivision code ("US-CA", "GB-ENG", "CA-ON"). Each region code
	// may appear in only one rule across `regions[]`.
	RegionCode string                                 `json:"regionCode" api:"required"`
	Rule       ConsentSettingUpdateResponseRegionRule `json:"rule" api:"required"`
	// Other region codes that should reuse this rule. Same code-format rules as
	// `regionCode`. Cannot include `regionCode` itself, cannot duplicate, cannot
	// overlap with another rule's regions.
	AdditionalRegions []string `json:"additionalRegions" api:"nullable"`
	// Optional label for this override group, e.g. "CCPA States". Shown in the
	// dashboard instead of the region codes. Pass null to clear it and fall back to
	// the codes.
	Name string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RegionCode        respjson.Field
		Rule              respjson.Field
		AdditionalRegions respjson.Field
		Name              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseRegion) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingUpdateResponseRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingUpdateResponseRegionRule struct {
	// Per-category default config for this rule. Every category defined in the
	// top-level `categories[].value` should have an entry here.
	Categories []ConsentSettingUpdateResponseRegionRuleCategory `json:"categories" api:"required"`
	// BCP 47 default language for this rule. Must have a matching entry in
	// `translations`. Examples: "en", "en-US", "es", "de".
	Language string `json:"language" api:"required"`
	// opt_in: scripts blocked until user accepts (GDPR style). opt_out: scripts run by
	// default until user rejects (CCPA style).
	//
	// Any of "opt_in", "opt_out".
	Mode string `json:"mode" api:"required"`
	// All UI copy, keyed by language. Must include an entry whose `language` matches
	// the rule's `language` field.
	Translations []ConsentSettingUpdateResponseRegionRuleTranslation `json:"translations" api:"required"`
	// When true, scripts not classified by services[] are blocked until the user opts
	// in.
	AutoblockUnknown bool `json:"autoblockUnknown" api:"nullable"`
	// When true, the consent modal auto-opens on page load.
	AutoShow bool `json:"autoShow" api:"nullable"`
	// Threshold config for autoShowDismissMode (page count or seconds).
	AutoShowDismissConfig ConsentSettingUpdateResponseRegionRuleAutoShowDismissConfig `json:"autoShowDismissConfig" api:"nullable"`
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	//
	// Any of "after_pages", "after_seconds", "never".
	AutoShowDismissMode string `json:"autoShowDismissMode" api:"nullable"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction bool `json:"disablePageInteraction" api:"nullable"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions ConsentSettingUpdateResponseRegionRuleGuiOptions `json:"guiOptions" api:"nullable"`
	// When true, the modal is suppressed for known bot user agents.
	HideFromBots bool `json:"hideFromBots" api:"nullable"`
	// When true, the per-service list (services[]) is rendered inside the preferences
	// modal.
	ShowVendorsInPreferences bool `json:"showVendorsInPreferences" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories               respjson.Field
		Language                 respjson.Field
		Mode                     respjson.Field
		Translations             respjson.Field
		AutoblockUnknown         respjson.Field
		AutoShow                 respjson.Field
		AutoShowDismissConfig    respjson.Field
		AutoShowDismissMode      respjson.Field
		DisablePageInteraction   respjson.Field
		GuiOptions               respjson.Field
		HideFromBots             respjson.Field
		ShowVendorsInPreferences respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseRegionRule) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingUpdateResponseRegionRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingUpdateResponseRegionRuleCategory struct {
	// Category value (matches `categories[].value`) this entry configures.
	Key   string                                              `json:"key" api:"required"`
	Value ConsentSettingUpdateResponseRegionRuleCategoryValue `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseRegionRuleCategory) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingUpdateResponseRegionRuleCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingUpdateResponseRegionRuleCategoryValue struct {
	// Whether this category is on by default before the user interacts.
	Enabled bool `json:"enabled" api:"required"`
	// When true, this category defaults off if the browser sends Sec-GPC: 1.
	AutoDisableOnGpc bool `json:"autoDisableOnGPC" api:"nullable"`
	// When true, the user cannot toggle this category in the preferences modal.
	ReadOnly bool `json:"readOnly" api:"nullable"`
	// When true, the page reloads after this category is toggled so newly-allowed
	// scripts can run.
	ReloadPage bool `json:"reloadPage" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled          respjson.Field
		AutoDisableOnGpc respjson.Field
		ReadOnly         respjson.Field
		ReloadPage       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseRegionRuleCategoryValue) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingUpdateResponseRegionRuleCategoryValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingUpdateResponseRegionRuleTranslation struct {
	// BCP 47 language tag identifying which translation this entry provides. Examples:
	// "en", "en-US", "es", "fr-CA". The default rule's `language` must appear here.
	Language string                                                 `json:"language" api:"required"`
	Value    ConsentSettingUpdateResponseRegionRuleTranslationValue `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Language    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseRegionRuleTranslation) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingUpdateResponseRegionRuleTranslation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingUpdateResponseRegionRuleTranslationValue struct {
	// Translated copy for the initial consent modal.
	ConsentModal ConsentSettingUpdateResponseRegionRuleTranslationValueConsentModal `json:"consentModal" api:"nullable"`
	// Translated copy for the preferences modal.
	PreferencesModal ConsentSettingUpdateResponseRegionRuleTranslationValuePreferencesModal `json:"preferencesModal" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConsentModal     respjson.Field
		PreferencesModal respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseRegionRuleTranslationValue) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingUpdateResponseRegionRuleTranslationValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the initial consent modal.
type ConsentSettingUpdateResponseRegionRuleTranslationValueConsentModal struct {
	// Label for the "accept all" button.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Body text shown in the consent modal.
	Description string `json:"description" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the button that opens the preferences modal.
	ShowPreferencesBtn string `json:"showPreferencesBtn" api:"required"`
	// Heading shown at the top of the consent modal.
	Title string `json:"title" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel string `json:"closeIconLabel" api:"nullable"`
	// Optional footer HTML/text shown below the body.
	Footer string `json:"footer" api:"nullable"`
	// Optional text shown when the browser is sending Sec-GPC: 1.
	GpcNotification string `json:"gpcNotification" api:"nullable"`
	// Link label for the privacy policy (defaults to "Privacy Policy").
	PrivacyPolicyLabel string `json:"privacyPolicyLabel" api:"nullable"`
	// Privacy policy link. Accepts an absolute URL (https://...) or a relative path
	// (/legal/privacy) so one config can run across multiple domains. javascript:,
	// data:, and vbscript: schemes are rejected.
	PrivacyPolicyURL string `json:"privacyPolicyUrl" api:"nullable"`
	// Optional message shown when the consent revision has bumped (re-prompt copy).
	RevisionMessage string `json:"revisionMessage" api:"nullable"`
	// Link label for terms of service (defaults to "Terms of Service").
	TermsOfServiceLabel string `json:"termsOfServiceLabel" api:"nullable"`
	// Terms of service link. Same scheme rules as privacyPolicyUrl: absolute or
	// relative path only.
	TermsOfServiceURL string `json:"termsOfServiceUrl" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptAllBtn        respjson.Field
		Description         respjson.Field
		RejectAllBtn        respjson.Field
		ShowPreferencesBtn  respjson.Field
		Title               respjson.Field
		CloseIconLabel      respjson.Field
		Footer              respjson.Field
		GpcNotification     respjson.Field
		PrivacyPolicyLabel  respjson.Field
		PrivacyPolicyURL    respjson.Field
		RevisionMessage     respjson.Field
		TermsOfServiceLabel respjson.Field
		TermsOfServiceURL   respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseRegionRuleTranslationValueConsentModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingUpdateResponseRegionRuleTranslationValueConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the preferences modal.
type ConsentSettingUpdateResponseRegionRuleTranslationValuePreferencesModal struct {
	// Label for the "accept all" button in the preferences modal.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel string `json:"closeIconLabel" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the "save my choices" button.
	SavePreferencesBtn string `json:"savePreferencesBtn" api:"required"`
	// Per-category sections plus a leading general section. The server backfills
	// sections for each defined category if absent.
	Sections []ConsentSettingUpdateResponseRegionRuleTranslationValuePreferencesModalSection `json:"sections" api:"required"`
	// Heading for the preferences modal.
	Title string `json:"title" api:"required"`
	// Optional collapsible info block rendered after the category sections. Renders
	// only when both title and body are non-empty. Use it for long-form text unrelated
	// to any specific category (e.g. a third-party advertisers disclosure).
	AdditionalInfo ConsentSettingUpdateResponseRegionRuleTranslationValuePreferencesModalAdditionalInfo `json:"additionalInfo" api:"nullable"`
	// Pluralizable label for the "N services" counter (e.g. "Service|Services").
	ServiceCounterLabel string `json:"serviceCounterLabel" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptAllBtn        respjson.Field
		CloseIconLabel      respjson.Field
		RejectAllBtn        respjson.Field
		SavePreferencesBtn  respjson.Field
		Sections            respjson.Field
		Title               respjson.Field
		AdditionalInfo      respjson.Field
		ServiceCounterLabel respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseRegionRuleTranslationValuePreferencesModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingUpdateResponseRegionRuleTranslationValuePreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingUpdateResponseRegionRuleTranslationValuePreferencesModalSection struct {
	// Body text for this preferences-modal section.
	Description string `json:"description" api:"required"`
	// Heading for this preferences-modal section.
	Title string `json:"title" api:"required"`
	// Optional disclosure table listing cookies/services in this section.
	CookieTable ConsentSettingUpdateResponseRegionRuleTranslationValuePreferencesModalSectionCookieTable `json:"cookieTable" api:"nullable"`
	// Category value this section configures. Use null/omit for the introductory
	// section that applies to no specific category.
	LinkedCategory string `json:"linkedCategory" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description    respjson.Field
		Title          respjson.Field
		CookieTable    respjson.Field
		LinkedCategory respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseRegionRuleTranslationValuePreferencesModalSection) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingUpdateResponseRegionRuleTranslationValuePreferencesModalSection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional disclosure table listing cookies/services in this section.
type ConsentSettingUpdateResponseRegionRuleTranslationValuePreferencesModalSectionCookieTable struct {
	// Rows of the cookie table — each row is a {key,value} pair (e.g. cookie name →
	// purpose).
	Body []ConsentSettingUpdateResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableBody `json:"body" api:"required"`
	// Header cells. Same {key,value} shape as rows.
	Headers []ConsentSettingUpdateResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader `json:"headers" api:"required"`
	// Optional table caption.
	Caption string `json:"caption" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Headers     respjson.Field
		Caption     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseRegionRuleTranslationValuePreferencesModalSectionCookieTable) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingUpdateResponseRegionRuleTranslationValuePreferencesModalSectionCookieTable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingUpdateResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableBody struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableBody) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingUpdateResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingUpdateResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingUpdateResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional collapsible info block rendered after the category sections. Renders
// only when both title and body are non-empty. Use it for long-form text unrelated
// to any specific category (e.g. a third-party advertisers disclosure).
type ConsentSettingUpdateResponseRegionRuleTranslationValuePreferencesModalAdditionalInfo struct {
	// Body text for the additional-info block. Supports the rich-text markdown subset
	// used by section descriptions: line breaks, [text](url) links (https:// or
	// relative paths), and **bold**. HTML is stripped.
	Body string `json:"body" api:"required"`
	// Heading for the additional-info block.
	Title string `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseRegionRuleTranslationValuePreferencesModalAdditionalInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingUpdateResponseRegionRuleTranslationValuePreferencesModalAdditionalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Threshold config for autoShowDismissMode (page count or seconds).
type ConsentSettingUpdateResponseRegionRuleAutoShowDismissConfig struct {
	// Pages-viewed threshold when autoShowDismissMode = after_pages.
	PageCount int64 `json:"pageCount" api:"nullable"`
	// Seconds-on-page threshold when autoShowDismissMode = after_seconds.
	Seconds int64 `json:"seconds" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageCount   respjson.Field
		Seconds     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseRegionRuleAutoShowDismissConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingUpdateResponseRegionRuleAutoShowDismissConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the modals (layout/position/colors).
type ConsentSettingUpdateResponseRegionRuleGuiOptions struct {
	// Visual options for the initial consent modal.
	ConsentModal ConsentSettingUpdateResponseRegionRuleGuiOptionsConsentModal `json:"consentModal" api:"nullable"`
	// Global CSS variable overrides applied to both modals.
	CssVariables ConsentSettingUpdateResponseRegionRuleGuiOptionsCssVariables `json:"cssVariables" api:"nullable"`
	// Visual options for the per-category preferences modal.
	PreferencesModal ConsentSettingUpdateResponseRegionRuleGuiOptionsPreferencesModal `json:"preferencesModal" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConsentModal     respjson.Field
		CssVariables     respjson.Field
		PreferencesModal respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseRegionRuleGuiOptions) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingUpdateResponseRegionRuleGuiOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the initial consent modal.
type ConsentSettingUpdateResponseRegionRuleGuiOptionsConsentModal struct {
	// Which buttons to render in the consent modal and in what arrangement.
	//
	// Any of "AcceptAllNecessaryPreferences", "AcceptAllPreferences", "AcceptOnly",
	// "InformationOnly", "PreferencesOnly".
	ButtonLayout string `json:"buttonLayout" api:"nullable"`
	// When true, accept and reject buttons share the same visual weight (no
	// primary/secondary).
	EqualWeightButtons bool `json:"equalWeightButtons" api:"nullable"`
	// When true, swap accept/reject button order.
	FlipButtons bool `json:"flipButtons" api:"nullable"`
	// Modal layout style (box, bar, cloud, etc.).
	//
	// Any of "bar", "bar_inline", "box", "box_inline", "box_wide", "cloud",
	// "cloud_inline".
	Layout string `json:"layout" api:"nullable"`
	// Where the modal anchors on the viewport (e.g. bottom_left, top_right,
	// middle_center).
	//
	// Any of "bottom", "bottom_center", "bottom_left", "bottom_right", "middle",
	// "middle_center", "middle_left", "middle_right", "top", "top_center", "top_left",
	// "top_right".
	Position string `json:"position" api:"nullable"`
	// Show an X to dismiss the modal without choosing.
	ShowCloseIcon bool `json:"showCloseIcon" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonLayout       respjson.Field
		EqualWeightButtons respjson.Field
		FlipButtons        respjson.Field
		Layout             respjson.Field
		Position           respjson.Field
		ShowCloseIcon      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseRegionRuleGuiOptionsConsentModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingUpdateResponseRegionRuleGuiOptionsConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Global CSS variable overrides applied to both modals.
type ConsentSettingUpdateResponseRegionRuleGuiOptionsCssVariables struct {
	// CSS length, e.g. "4px" or "0.25rem".
	ButtonBorderRadius string `json:"buttonBorderRadius" api:"nullable"`
	// CSS color for the per-category block background in the preferences modal.
	CookieCategoryBlockBg string `json:"cookieCategoryBlockBg" api:"nullable"`
	// CSS color for the per-category block background on hover.
	CookieCategoryBlockHoverBg string `json:"cookieCategoryBlockHoverBg" api:"nullable"`
	// CSS color for the modal footer background.
	FooterBg string `json:"footerBg" api:"nullable"`
	// CSS color for footer text.
	FooterColor string `json:"footerColor" api:"nullable"`
	// CSS color for footer links.
	FooterLinkColor string `json:"footerLinkColor" api:"nullable"`
	// CSS color for footer links on hover.
	FooterLinkHoverColor string `json:"footerLinkHoverColor" api:"nullable"`
	// CSS color for the modal background.
	ModalBg string `json:"modalBg" api:"nullable"`
	// CSS length for modal corner radius.
	ModalBorderRadius string `json:"modalBorderRadius" api:"nullable"`
	// CSS color for the primary button background.
	PrimaryButtonBg string `json:"primaryButtonBg" api:"nullable"`
	// CSS color for primary button text.
	PrimaryButtonColor string `json:"primaryButtonColor" api:"nullable"`
	// CSS color for the primary button background on hover.
	PrimaryButtonHoverBg string `json:"primaryButtonHoverBg" api:"nullable"`
	// CSS color for primary button text on hover.
	PrimaryButtonHoverColor string `json:"primaryButtonHoverColor" api:"nullable"`
	// CSS color for primary text in the modal body.
	PrimaryTextColor string `json:"primaryTextColor" api:"nullable"`
	// CSS color for the secondary button background.
	SecondaryButtonBg string `json:"secondaryButtonBg" api:"nullable"`
	// CSS color for secondary button text.
	SecondaryButtonColor string `json:"secondaryButtonColor" api:"nullable"`
	// CSS color for secondary button background on hover.
	SecondaryButtonHoverBg string `json:"secondaryButtonHoverBg" api:"nullable"`
	// CSS color for secondary button text on hover.
	SecondaryButtonHoverColor string `json:"secondaryButtonHoverColor" api:"nullable"`
	// CSS color for secondary text in the modal body.
	SecondaryTextColor string `json:"secondaryTextColor" api:"nullable"`
	// CSS color for the "on" state of category toggles.
	ToggleOnBg string `json:"toggleOnBg" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonBorderRadius         respjson.Field
		CookieCategoryBlockBg      respjson.Field
		CookieCategoryBlockHoverBg respjson.Field
		FooterBg                   respjson.Field
		FooterColor                respjson.Field
		FooterLinkColor            respjson.Field
		FooterLinkHoverColor       respjson.Field
		ModalBg                    respjson.Field
		ModalBorderRadius          respjson.Field
		PrimaryButtonBg            respjson.Field
		PrimaryButtonColor         respjson.Field
		PrimaryButtonHoverBg       respjson.Field
		PrimaryButtonHoverColor    respjson.Field
		PrimaryTextColor           respjson.Field
		SecondaryButtonBg          respjson.Field
		SecondaryButtonColor       respjson.Field
		SecondaryButtonHoverBg     respjson.Field
		SecondaryButtonHoverColor  respjson.Field
		SecondaryTextColor         respjson.Field
		ToggleOnBg                 respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseRegionRuleGuiOptionsCssVariables) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingUpdateResponseRegionRuleGuiOptionsCssVariables) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the per-category preferences modal.
type ConsentSettingUpdateResponseRegionRuleGuiOptionsPreferencesModal struct {
	// Button layout for the preferences modal.
	//
	// Any of "AcceptAllRejectAllSave", "AcceptAllSave".
	ButtonLayout string `json:"buttonLayout" api:"nullable"`
	// When true, accept/reject share the same visual weight.
	EqualWeightButtons bool `json:"equalWeightButtons" api:"nullable"`
	// When true, swap accept/reject button order.
	FlipButtons bool `json:"flipButtons" api:"nullable"`
	// Preferences modal layout (box, bar, bar_wide).
	//
	// Any of "bar", "bar_wide", "box".
	Layout string `json:"layout" api:"nullable"`
	// Preferences modal anchor (left, right).
	//
	// Any of "left", "right".
	Position string `json:"position" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonLayout       respjson.Field
		EqualWeightButtons respjson.Field
		FlipButtons        respjson.Field
		Layout             respjson.Field
		Position           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseRegionRuleGuiOptionsPreferencesModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingUpdateResponseRegionRuleGuiOptionsPreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingUpdateResponseService struct {
	// Internal notes shown to admins in the dashboard. Not user-facing.
	InternalNotes string `json:"internalNotes" api:"required"`
	// Display name for this service in the preferences modal.
	Label string `json:"label" api:"required"`
	// Extra category values this service belongs to. Each must match a
	// `categories[].value`.
	AdditionalCategories []string `json:"additionalCategories" api:"nullable"`
	// Primary category value this service belongs to. Must match one of the top-level
	// `categories[].value` entries.
	Category string `json:"category" api:"nullable"`
	// Domains/paths this service matches. Patterns matching the CMP's own scripts
	// (e.g. cdn.oursprivacy.com/cmp-init) are rejected to prevent the CMP blocking
	// itself — use a more specific path like cdn.oursprivacy.com/main.js to block a
	// specific script.
	DomainPatterns []string `json:"domainPatterns" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InternalNotes        respjson.Field
		Label                respjson.Field
		AdditionalCategories respjson.Field
		Category             respjson.Field
		DomainPatterns       respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponseService) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingUpdateResponseService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enabled serves the CMP on its allowed domains and requires at least one.
// Disabled does not serve the CMP and may have no allowed domains.
type ConsentSettingUpdateResponseStatus string

const (
	ConsentSettingUpdateResponseStatusDisabled ConsentSettingUpdateResponseStatus = "Disabled"
	ConsentSettingUpdateResponseStatusEnabled  ConsentSettingUpdateResponseStatus = "Enabled"
)

type ConsentSettingDeleteResponse struct {
	// Server-assigned UUID for this consent settings record.
	ID string `json:"id" api:"required"`
	// Top-level consent categories (e.g. necessary / analytics / advertising). Server
	// re-stamps `priority` to 0..N on write.
	Categories []ConsentSettingDeleteResponseCategory `json:"categories" api:"required"`
	// ISO 8601 timestamp when the record was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Default rule used when the user is not in any region listed in `regions[]`.
	Default ConsentSettingDeleteResponseDefault `json:"default" api:"required"`
	// Discriminator for the entity type. Always "consentSettings".
	Kind string `json:"kind" api:"required"`
	// Human-readable name shown in the dashboard.
	Name string `json:"name" api:"required"`
	// Per-region rule overrides. The first rule whose `regionCode`/`additionalRegions`
	// includes the user's region wins; otherwise `default` applies.
	Regions []ConsentSettingDeleteResponseRegion `json:"regions" api:"required"`
	// Per-service entries powering "show vendors" and category-aware blocking.
	Services []ConsentSettingDeleteResponseService `json:"services" api:"required"`
	// Enabled serves the CMP on its allowed domains and requires at least one.
	// Disabled does not serve the CMP and may have no allowed domains.
	//
	// Any of "Disabled", "Enabled".
	Status ConsentSettingDeleteResponseStatus `json:"status" api:"required"`
	// Name of the cookie that stores the user's consent state. Defaults to
	// "op_consent".
	ConsentCookieName string `json:"consentCookieName" api:"nullable"`
	// Optional custom CDN domain for serving the CMP script (e.g.
	// consent.example.com).
	CustomDomain string `json:"customDomain" api:"nullable"`
	// Cookie/localStorage key for the visitor (device) ID. Must exactly match the web
	// SDK's cookie_names.device_id (default "ours_device_id") so consent and SDK
	// events resolve to the same visitor.
	DeviceIDCookieName string `json:"deviceIdCookieName" api:"nullable"`
	// Revision counter. Bump this to force users who already consented to see the
	// modal again (the SDK compares the persisted revision against this value).
	Revision float64 `json:"revision" api:"nullable"`
	// CSS class names that opt scripts out of consent blocking. Each entry must be a
	// single class token (no whitespace).
	SkipBlockingClassNames []string `json:"skipBlockingClassNames" api:"nullable"`
	// ISO 8601 timestamp of the last write. Null on a freshly created record.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// Pixel of the WebSource that this CMP is wired into. Setting this to a token that
	// is not a valid WebSource of yours is rejected; use null to clear the link.
	WebSDKToken string `json:"webSDKToken" api:"nullable"`
	// Domains where this CMP may run. An Enabled CMP requires at least one allowed
	// domain; with none configured, the CMP does not run. Matching domains may share
	// consent across subdomains.
	WhitelistDomains []string `json:"whitelistDomains" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		Categories             respjson.Field
		CreatedAt              respjson.Field
		Default                respjson.Field
		Kind                   respjson.Field
		Name                   respjson.Field
		Regions                respjson.Field
		Services               respjson.Field
		Status                 respjson.Field
		ConsentCookieName      respjson.Field
		CustomDomain           respjson.Field
		DeviceIDCookieName     respjson.Field
		Revision               respjson.Field
		SkipBlockingClassNames respjson.Field
		UpdatedAt              respjson.Field
		WebSDKToken            respjson.Field
		WhitelistDomains       respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingDeleteResponseCategory struct {
	// Human-readable label shown next to the toggle in the preferences modal.
	Label string `json:"label" api:"required"`
	// Sort key. Lower numbers render first. Server re-stamps to 0..N on write — send
	// any integer, gaps and duplicates are ironed out.
	Priority int64 `json:"priority" api:"required"`
	// Stable identifier referenced by services and translation sections.
	// Conventionally lowercase (e.g. "necessary", "analytics", "advertising").
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		Priority    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseCategory) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingDeleteResponseCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default rule used when the user is not in any region listed in `regions[]`.
type ConsentSettingDeleteResponseDefault struct {
	// Per-category default config for this rule. Every category defined in the
	// top-level `categories[].value` should have an entry here.
	Categories []ConsentSettingDeleteResponseDefaultCategory `json:"categories" api:"required"`
	// BCP 47 default language for this rule. Must have a matching entry in
	// `translations`. Examples: "en", "en-US", "es", "de".
	Language string `json:"language" api:"required"`
	// opt_in: scripts blocked until user accepts (GDPR style). opt_out: scripts run by
	// default until user rejects (CCPA style).
	//
	// Any of "opt_in", "opt_out".
	Mode string `json:"mode" api:"required"`
	// All UI copy, keyed by language. Must include an entry whose `language` matches
	// the rule's `language` field.
	Translations []ConsentSettingDeleteResponseDefaultTranslation `json:"translations" api:"required"`
	// When true, scripts not classified by services[] are blocked until the user opts
	// in.
	AutoblockUnknown bool `json:"autoblockUnknown" api:"nullable"`
	// When true, the consent modal auto-opens on page load.
	AutoShow bool `json:"autoShow" api:"nullable"`
	// Threshold config for autoShowDismissMode (page count or seconds).
	AutoShowDismissConfig ConsentSettingDeleteResponseDefaultAutoShowDismissConfig `json:"autoShowDismissConfig" api:"nullable"`
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	//
	// Any of "after_pages", "after_seconds", "never".
	AutoShowDismissMode string `json:"autoShowDismissMode" api:"nullable"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction bool `json:"disablePageInteraction" api:"nullable"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions ConsentSettingDeleteResponseDefaultGuiOptions `json:"guiOptions" api:"nullable"`
	// When true, the modal is suppressed for known bot user agents.
	HideFromBots bool `json:"hideFromBots" api:"nullable"`
	// When true, the per-service list (services[]) is rendered inside the preferences
	// modal.
	ShowVendorsInPreferences bool `json:"showVendorsInPreferences" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories               respjson.Field
		Language                 respjson.Field
		Mode                     respjson.Field
		Translations             respjson.Field
		AutoblockUnknown         respjson.Field
		AutoShow                 respjson.Field
		AutoShowDismissConfig    respjson.Field
		AutoShowDismissMode      respjson.Field
		DisablePageInteraction   respjson.Field
		GuiOptions               respjson.Field
		HideFromBots             respjson.Field
		ShowVendorsInPreferences respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseDefault) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingDeleteResponseDefault) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingDeleteResponseDefaultCategory struct {
	// Category value (matches `categories[].value`) this entry configures.
	Key   string                                           `json:"key" api:"required"`
	Value ConsentSettingDeleteResponseDefaultCategoryValue `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseDefaultCategory) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingDeleteResponseDefaultCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingDeleteResponseDefaultCategoryValue struct {
	// Whether this category is on by default before the user interacts.
	Enabled bool `json:"enabled" api:"required"`
	// When true, this category defaults off if the browser sends Sec-GPC: 1.
	AutoDisableOnGpc bool `json:"autoDisableOnGPC" api:"nullable"`
	// When true, the user cannot toggle this category in the preferences modal.
	ReadOnly bool `json:"readOnly" api:"nullable"`
	// When true, the page reloads after this category is toggled so newly-allowed
	// scripts can run.
	ReloadPage bool `json:"reloadPage" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled          respjson.Field
		AutoDisableOnGpc respjson.Field
		ReadOnly         respjson.Field
		ReloadPage       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseDefaultCategoryValue) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingDeleteResponseDefaultCategoryValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingDeleteResponseDefaultTranslation struct {
	// BCP 47 language tag identifying which translation this entry provides. Examples:
	// "en", "en-US", "es", "fr-CA". The default rule's `language` must appear here.
	Language string                                              `json:"language" api:"required"`
	Value    ConsentSettingDeleteResponseDefaultTranslationValue `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Language    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseDefaultTranslation) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingDeleteResponseDefaultTranslation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingDeleteResponseDefaultTranslationValue struct {
	// Translated copy for the initial consent modal.
	ConsentModal ConsentSettingDeleteResponseDefaultTranslationValueConsentModal `json:"consentModal" api:"nullable"`
	// Translated copy for the preferences modal.
	PreferencesModal ConsentSettingDeleteResponseDefaultTranslationValuePreferencesModal `json:"preferencesModal" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConsentModal     respjson.Field
		PreferencesModal respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseDefaultTranslationValue) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingDeleteResponseDefaultTranslationValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the initial consent modal.
type ConsentSettingDeleteResponseDefaultTranslationValueConsentModal struct {
	// Label for the "accept all" button.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Body text shown in the consent modal.
	Description string `json:"description" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the button that opens the preferences modal.
	ShowPreferencesBtn string `json:"showPreferencesBtn" api:"required"`
	// Heading shown at the top of the consent modal.
	Title string `json:"title" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel string `json:"closeIconLabel" api:"nullable"`
	// Optional footer HTML/text shown below the body.
	Footer string `json:"footer" api:"nullable"`
	// Optional text shown when the browser is sending Sec-GPC: 1.
	GpcNotification string `json:"gpcNotification" api:"nullable"`
	// Link label for the privacy policy (defaults to "Privacy Policy").
	PrivacyPolicyLabel string `json:"privacyPolicyLabel" api:"nullable"`
	// Privacy policy link. Accepts an absolute URL (https://...) or a relative path
	// (/legal/privacy) so one config can run across multiple domains. javascript:,
	// data:, and vbscript: schemes are rejected.
	PrivacyPolicyURL string `json:"privacyPolicyUrl" api:"nullable"`
	// Optional message shown when the consent revision has bumped (re-prompt copy).
	RevisionMessage string `json:"revisionMessage" api:"nullable"`
	// Link label for terms of service (defaults to "Terms of Service").
	TermsOfServiceLabel string `json:"termsOfServiceLabel" api:"nullable"`
	// Terms of service link. Same scheme rules as privacyPolicyUrl: absolute or
	// relative path only.
	TermsOfServiceURL string `json:"termsOfServiceUrl" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptAllBtn        respjson.Field
		Description         respjson.Field
		RejectAllBtn        respjson.Field
		ShowPreferencesBtn  respjson.Field
		Title               respjson.Field
		CloseIconLabel      respjson.Field
		Footer              respjson.Field
		GpcNotification     respjson.Field
		PrivacyPolicyLabel  respjson.Field
		PrivacyPolicyURL    respjson.Field
		RevisionMessage     respjson.Field
		TermsOfServiceLabel respjson.Field
		TermsOfServiceURL   respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseDefaultTranslationValueConsentModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingDeleteResponseDefaultTranslationValueConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the preferences modal.
type ConsentSettingDeleteResponseDefaultTranslationValuePreferencesModal struct {
	// Label for the "accept all" button in the preferences modal.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel string `json:"closeIconLabel" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the "save my choices" button.
	SavePreferencesBtn string `json:"savePreferencesBtn" api:"required"`
	// Per-category sections plus a leading general section. The server backfills
	// sections for each defined category if absent.
	Sections []ConsentSettingDeleteResponseDefaultTranslationValuePreferencesModalSection `json:"sections" api:"required"`
	// Heading for the preferences modal.
	Title string `json:"title" api:"required"`
	// Optional collapsible info block rendered after the category sections. Renders
	// only when both title and body are non-empty. Use it for long-form text unrelated
	// to any specific category (e.g. a third-party advertisers disclosure).
	AdditionalInfo ConsentSettingDeleteResponseDefaultTranslationValuePreferencesModalAdditionalInfo `json:"additionalInfo" api:"nullable"`
	// Pluralizable label for the "N services" counter (e.g. "Service|Services").
	ServiceCounterLabel string `json:"serviceCounterLabel" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptAllBtn        respjson.Field
		CloseIconLabel      respjson.Field
		RejectAllBtn        respjson.Field
		SavePreferencesBtn  respjson.Field
		Sections            respjson.Field
		Title               respjson.Field
		AdditionalInfo      respjson.Field
		ServiceCounterLabel respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseDefaultTranslationValuePreferencesModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingDeleteResponseDefaultTranslationValuePreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingDeleteResponseDefaultTranslationValuePreferencesModalSection struct {
	// Body text for this preferences-modal section.
	Description string `json:"description" api:"required"`
	// Heading for this preferences-modal section.
	Title string `json:"title" api:"required"`
	// Optional disclosure table listing cookies/services in this section.
	CookieTable ConsentSettingDeleteResponseDefaultTranslationValuePreferencesModalSectionCookieTable `json:"cookieTable" api:"nullable"`
	// Category value this section configures. Use null/omit for the introductory
	// section that applies to no specific category.
	LinkedCategory string `json:"linkedCategory" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description    respjson.Field
		Title          respjson.Field
		CookieTable    respjson.Field
		LinkedCategory respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseDefaultTranslationValuePreferencesModalSection) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingDeleteResponseDefaultTranslationValuePreferencesModalSection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional disclosure table listing cookies/services in this section.
type ConsentSettingDeleteResponseDefaultTranslationValuePreferencesModalSectionCookieTable struct {
	// Rows of the cookie table — each row is a {key,value} pair (e.g. cookie name →
	// purpose).
	Body []ConsentSettingDeleteResponseDefaultTranslationValuePreferencesModalSectionCookieTableBody `json:"body" api:"required"`
	// Header cells. Same {key,value} shape as rows.
	Headers []ConsentSettingDeleteResponseDefaultTranslationValuePreferencesModalSectionCookieTableHeader `json:"headers" api:"required"`
	// Optional table caption.
	Caption string `json:"caption" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Headers     respjson.Field
		Caption     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseDefaultTranslationValuePreferencesModalSectionCookieTable) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingDeleteResponseDefaultTranslationValuePreferencesModalSectionCookieTable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingDeleteResponseDefaultTranslationValuePreferencesModalSectionCookieTableBody struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseDefaultTranslationValuePreferencesModalSectionCookieTableBody) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingDeleteResponseDefaultTranslationValuePreferencesModalSectionCookieTableBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingDeleteResponseDefaultTranslationValuePreferencesModalSectionCookieTableHeader struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseDefaultTranslationValuePreferencesModalSectionCookieTableHeader) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingDeleteResponseDefaultTranslationValuePreferencesModalSectionCookieTableHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional collapsible info block rendered after the category sections. Renders
// only when both title and body are non-empty. Use it for long-form text unrelated
// to any specific category (e.g. a third-party advertisers disclosure).
type ConsentSettingDeleteResponseDefaultTranslationValuePreferencesModalAdditionalInfo struct {
	// Body text for the additional-info block. Supports the rich-text markdown subset
	// used by section descriptions: line breaks, [text](url) links (https:// or
	// relative paths), and **bold**. HTML is stripped.
	Body string `json:"body" api:"required"`
	// Heading for the additional-info block.
	Title string `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseDefaultTranslationValuePreferencesModalAdditionalInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingDeleteResponseDefaultTranslationValuePreferencesModalAdditionalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Threshold config for autoShowDismissMode (page count or seconds).
type ConsentSettingDeleteResponseDefaultAutoShowDismissConfig struct {
	// Pages-viewed threshold when autoShowDismissMode = after_pages.
	PageCount int64 `json:"pageCount" api:"nullable"`
	// Seconds-on-page threshold when autoShowDismissMode = after_seconds.
	Seconds int64 `json:"seconds" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageCount   respjson.Field
		Seconds     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseDefaultAutoShowDismissConfig) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingDeleteResponseDefaultAutoShowDismissConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the modals (layout/position/colors).
type ConsentSettingDeleteResponseDefaultGuiOptions struct {
	// Visual options for the initial consent modal.
	ConsentModal ConsentSettingDeleteResponseDefaultGuiOptionsConsentModal `json:"consentModal" api:"nullable"`
	// Global CSS variable overrides applied to both modals.
	CssVariables ConsentSettingDeleteResponseDefaultGuiOptionsCssVariables `json:"cssVariables" api:"nullable"`
	// Visual options for the per-category preferences modal.
	PreferencesModal ConsentSettingDeleteResponseDefaultGuiOptionsPreferencesModal `json:"preferencesModal" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConsentModal     respjson.Field
		CssVariables     respjson.Field
		PreferencesModal respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseDefaultGuiOptions) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingDeleteResponseDefaultGuiOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the initial consent modal.
type ConsentSettingDeleteResponseDefaultGuiOptionsConsentModal struct {
	// Which buttons to render in the consent modal and in what arrangement.
	//
	// Any of "AcceptAllNecessaryPreferences", "AcceptAllPreferences", "AcceptOnly",
	// "InformationOnly", "PreferencesOnly".
	ButtonLayout string `json:"buttonLayout" api:"nullable"`
	// When true, accept and reject buttons share the same visual weight (no
	// primary/secondary).
	EqualWeightButtons bool `json:"equalWeightButtons" api:"nullable"`
	// When true, swap accept/reject button order.
	FlipButtons bool `json:"flipButtons" api:"nullable"`
	// Modal layout style (box, bar, cloud, etc.).
	//
	// Any of "bar", "bar_inline", "box", "box_inline", "box_wide", "cloud",
	// "cloud_inline".
	Layout string `json:"layout" api:"nullable"`
	// Where the modal anchors on the viewport (e.g. bottom_left, top_right,
	// middle_center).
	//
	// Any of "bottom", "bottom_center", "bottom_left", "bottom_right", "middle",
	// "middle_center", "middle_left", "middle_right", "top", "top_center", "top_left",
	// "top_right".
	Position string `json:"position" api:"nullable"`
	// Show an X to dismiss the modal without choosing.
	ShowCloseIcon bool `json:"showCloseIcon" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonLayout       respjson.Field
		EqualWeightButtons respjson.Field
		FlipButtons        respjson.Field
		Layout             respjson.Field
		Position           respjson.Field
		ShowCloseIcon      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseDefaultGuiOptionsConsentModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingDeleteResponseDefaultGuiOptionsConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Global CSS variable overrides applied to both modals.
type ConsentSettingDeleteResponseDefaultGuiOptionsCssVariables struct {
	// CSS length, e.g. "4px" or "0.25rem".
	ButtonBorderRadius string `json:"buttonBorderRadius" api:"nullable"`
	// CSS color for the per-category block background in the preferences modal.
	CookieCategoryBlockBg string `json:"cookieCategoryBlockBg" api:"nullable"`
	// CSS color for the per-category block background on hover.
	CookieCategoryBlockHoverBg string `json:"cookieCategoryBlockHoverBg" api:"nullable"`
	// CSS color for the modal footer background.
	FooterBg string `json:"footerBg" api:"nullable"`
	// CSS color for footer text.
	FooterColor string `json:"footerColor" api:"nullable"`
	// CSS color for footer links.
	FooterLinkColor string `json:"footerLinkColor" api:"nullable"`
	// CSS color for footer links on hover.
	FooterLinkHoverColor string `json:"footerLinkHoverColor" api:"nullable"`
	// CSS color for the modal background.
	ModalBg string `json:"modalBg" api:"nullable"`
	// CSS length for modal corner radius.
	ModalBorderRadius string `json:"modalBorderRadius" api:"nullable"`
	// CSS color for the primary button background.
	PrimaryButtonBg string `json:"primaryButtonBg" api:"nullable"`
	// CSS color for primary button text.
	PrimaryButtonColor string `json:"primaryButtonColor" api:"nullable"`
	// CSS color for the primary button background on hover.
	PrimaryButtonHoverBg string `json:"primaryButtonHoverBg" api:"nullable"`
	// CSS color for primary button text on hover.
	PrimaryButtonHoverColor string `json:"primaryButtonHoverColor" api:"nullable"`
	// CSS color for primary text in the modal body.
	PrimaryTextColor string `json:"primaryTextColor" api:"nullable"`
	// CSS color for the secondary button background.
	SecondaryButtonBg string `json:"secondaryButtonBg" api:"nullable"`
	// CSS color for secondary button text.
	SecondaryButtonColor string `json:"secondaryButtonColor" api:"nullable"`
	// CSS color for secondary button background on hover.
	SecondaryButtonHoverBg string `json:"secondaryButtonHoverBg" api:"nullable"`
	// CSS color for secondary button text on hover.
	SecondaryButtonHoverColor string `json:"secondaryButtonHoverColor" api:"nullable"`
	// CSS color for secondary text in the modal body.
	SecondaryTextColor string `json:"secondaryTextColor" api:"nullable"`
	// CSS color for the "on" state of category toggles.
	ToggleOnBg string `json:"toggleOnBg" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonBorderRadius         respjson.Field
		CookieCategoryBlockBg      respjson.Field
		CookieCategoryBlockHoverBg respjson.Field
		FooterBg                   respjson.Field
		FooterColor                respjson.Field
		FooterLinkColor            respjson.Field
		FooterLinkHoverColor       respjson.Field
		ModalBg                    respjson.Field
		ModalBorderRadius          respjson.Field
		PrimaryButtonBg            respjson.Field
		PrimaryButtonColor         respjson.Field
		PrimaryButtonHoverBg       respjson.Field
		PrimaryButtonHoverColor    respjson.Field
		PrimaryTextColor           respjson.Field
		SecondaryButtonBg          respjson.Field
		SecondaryButtonColor       respjson.Field
		SecondaryButtonHoverBg     respjson.Field
		SecondaryButtonHoverColor  respjson.Field
		SecondaryTextColor         respjson.Field
		ToggleOnBg                 respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseDefaultGuiOptionsCssVariables) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingDeleteResponseDefaultGuiOptionsCssVariables) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the per-category preferences modal.
type ConsentSettingDeleteResponseDefaultGuiOptionsPreferencesModal struct {
	// Button layout for the preferences modal.
	//
	// Any of "AcceptAllRejectAllSave", "AcceptAllSave".
	ButtonLayout string `json:"buttonLayout" api:"nullable"`
	// When true, accept/reject share the same visual weight.
	EqualWeightButtons bool `json:"equalWeightButtons" api:"nullable"`
	// When true, swap accept/reject button order.
	FlipButtons bool `json:"flipButtons" api:"nullable"`
	// Preferences modal layout (box, bar, bar_wide).
	//
	// Any of "bar", "bar_wide", "box".
	Layout string `json:"layout" api:"nullable"`
	// Preferences modal anchor (left, right).
	//
	// Any of "left", "right".
	Position string `json:"position" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonLayout       respjson.Field
		EqualWeightButtons respjson.Field
		FlipButtons        respjson.Field
		Layout             respjson.Field
		Position           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseDefaultGuiOptionsPreferencesModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingDeleteResponseDefaultGuiOptionsPreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingDeleteResponseRegion struct {
	// Region this rule applies to. Use ISO 3166-1 alpha-2 country code ("US", "DE",
	// "BR") or country-subdivision code ("US-CA", "GB-ENG", "CA-ON"). Each region code
	// may appear in only one rule across `regions[]`.
	RegionCode string                                 `json:"regionCode" api:"required"`
	Rule       ConsentSettingDeleteResponseRegionRule `json:"rule" api:"required"`
	// Other region codes that should reuse this rule. Same code-format rules as
	// `regionCode`. Cannot include `regionCode` itself, cannot duplicate, cannot
	// overlap with another rule's regions.
	AdditionalRegions []string `json:"additionalRegions" api:"nullable"`
	// Optional label for this override group, e.g. "CCPA States". Shown in the
	// dashboard instead of the region codes. Pass null to clear it and fall back to
	// the codes.
	Name string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RegionCode        respjson.Field
		Rule              respjson.Field
		AdditionalRegions respjson.Field
		Name              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseRegion) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingDeleteResponseRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingDeleteResponseRegionRule struct {
	// Per-category default config for this rule. Every category defined in the
	// top-level `categories[].value` should have an entry here.
	Categories []ConsentSettingDeleteResponseRegionRuleCategory `json:"categories" api:"required"`
	// BCP 47 default language for this rule. Must have a matching entry in
	// `translations`. Examples: "en", "en-US", "es", "de".
	Language string `json:"language" api:"required"`
	// opt_in: scripts blocked until user accepts (GDPR style). opt_out: scripts run by
	// default until user rejects (CCPA style).
	//
	// Any of "opt_in", "opt_out".
	Mode string `json:"mode" api:"required"`
	// All UI copy, keyed by language. Must include an entry whose `language` matches
	// the rule's `language` field.
	Translations []ConsentSettingDeleteResponseRegionRuleTranslation `json:"translations" api:"required"`
	// When true, scripts not classified by services[] are blocked until the user opts
	// in.
	AutoblockUnknown bool `json:"autoblockUnknown" api:"nullable"`
	// When true, the consent modal auto-opens on page load.
	AutoShow bool `json:"autoShow" api:"nullable"`
	// Threshold config for autoShowDismissMode (page count or seconds).
	AutoShowDismissConfig ConsentSettingDeleteResponseRegionRuleAutoShowDismissConfig `json:"autoShowDismissConfig" api:"nullable"`
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	//
	// Any of "after_pages", "after_seconds", "never".
	AutoShowDismissMode string `json:"autoShowDismissMode" api:"nullable"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction bool `json:"disablePageInteraction" api:"nullable"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions ConsentSettingDeleteResponseRegionRuleGuiOptions `json:"guiOptions" api:"nullable"`
	// When true, the modal is suppressed for known bot user agents.
	HideFromBots bool `json:"hideFromBots" api:"nullable"`
	// When true, the per-service list (services[]) is rendered inside the preferences
	// modal.
	ShowVendorsInPreferences bool `json:"showVendorsInPreferences" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories               respjson.Field
		Language                 respjson.Field
		Mode                     respjson.Field
		Translations             respjson.Field
		AutoblockUnknown         respjson.Field
		AutoShow                 respjson.Field
		AutoShowDismissConfig    respjson.Field
		AutoShowDismissMode      respjson.Field
		DisablePageInteraction   respjson.Field
		GuiOptions               respjson.Field
		HideFromBots             respjson.Field
		ShowVendorsInPreferences respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseRegionRule) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingDeleteResponseRegionRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingDeleteResponseRegionRuleCategory struct {
	// Category value (matches `categories[].value`) this entry configures.
	Key   string                                              `json:"key" api:"required"`
	Value ConsentSettingDeleteResponseRegionRuleCategoryValue `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseRegionRuleCategory) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingDeleteResponseRegionRuleCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingDeleteResponseRegionRuleCategoryValue struct {
	// Whether this category is on by default before the user interacts.
	Enabled bool `json:"enabled" api:"required"`
	// When true, this category defaults off if the browser sends Sec-GPC: 1.
	AutoDisableOnGpc bool `json:"autoDisableOnGPC" api:"nullable"`
	// When true, the user cannot toggle this category in the preferences modal.
	ReadOnly bool `json:"readOnly" api:"nullable"`
	// When true, the page reloads after this category is toggled so newly-allowed
	// scripts can run.
	ReloadPage bool `json:"reloadPage" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled          respjson.Field
		AutoDisableOnGpc respjson.Field
		ReadOnly         respjson.Field
		ReloadPage       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseRegionRuleCategoryValue) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingDeleteResponseRegionRuleCategoryValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingDeleteResponseRegionRuleTranslation struct {
	// BCP 47 language tag identifying which translation this entry provides. Examples:
	// "en", "en-US", "es", "fr-CA". The default rule's `language` must appear here.
	Language string                                                 `json:"language" api:"required"`
	Value    ConsentSettingDeleteResponseRegionRuleTranslationValue `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Language    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseRegionRuleTranslation) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingDeleteResponseRegionRuleTranslation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingDeleteResponseRegionRuleTranslationValue struct {
	// Translated copy for the initial consent modal.
	ConsentModal ConsentSettingDeleteResponseRegionRuleTranslationValueConsentModal `json:"consentModal" api:"nullable"`
	// Translated copy for the preferences modal.
	PreferencesModal ConsentSettingDeleteResponseRegionRuleTranslationValuePreferencesModal `json:"preferencesModal" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConsentModal     respjson.Field
		PreferencesModal respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseRegionRuleTranslationValue) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingDeleteResponseRegionRuleTranslationValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the initial consent modal.
type ConsentSettingDeleteResponseRegionRuleTranslationValueConsentModal struct {
	// Label for the "accept all" button.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Body text shown in the consent modal.
	Description string `json:"description" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the button that opens the preferences modal.
	ShowPreferencesBtn string `json:"showPreferencesBtn" api:"required"`
	// Heading shown at the top of the consent modal.
	Title string `json:"title" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel string `json:"closeIconLabel" api:"nullable"`
	// Optional footer HTML/text shown below the body.
	Footer string `json:"footer" api:"nullable"`
	// Optional text shown when the browser is sending Sec-GPC: 1.
	GpcNotification string `json:"gpcNotification" api:"nullable"`
	// Link label for the privacy policy (defaults to "Privacy Policy").
	PrivacyPolicyLabel string `json:"privacyPolicyLabel" api:"nullable"`
	// Privacy policy link. Accepts an absolute URL (https://...) or a relative path
	// (/legal/privacy) so one config can run across multiple domains. javascript:,
	// data:, and vbscript: schemes are rejected.
	PrivacyPolicyURL string `json:"privacyPolicyUrl" api:"nullable"`
	// Optional message shown when the consent revision has bumped (re-prompt copy).
	RevisionMessage string `json:"revisionMessage" api:"nullable"`
	// Link label for terms of service (defaults to "Terms of Service").
	TermsOfServiceLabel string `json:"termsOfServiceLabel" api:"nullable"`
	// Terms of service link. Same scheme rules as privacyPolicyUrl: absolute or
	// relative path only.
	TermsOfServiceURL string `json:"termsOfServiceUrl" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptAllBtn        respjson.Field
		Description         respjson.Field
		RejectAllBtn        respjson.Field
		ShowPreferencesBtn  respjson.Field
		Title               respjson.Field
		CloseIconLabel      respjson.Field
		Footer              respjson.Field
		GpcNotification     respjson.Field
		PrivacyPolicyLabel  respjson.Field
		PrivacyPolicyURL    respjson.Field
		RevisionMessage     respjson.Field
		TermsOfServiceLabel respjson.Field
		TermsOfServiceURL   respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseRegionRuleTranslationValueConsentModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingDeleteResponseRegionRuleTranslationValueConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the preferences modal.
type ConsentSettingDeleteResponseRegionRuleTranslationValuePreferencesModal struct {
	// Label for the "accept all" button in the preferences modal.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel string `json:"closeIconLabel" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the "save my choices" button.
	SavePreferencesBtn string `json:"savePreferencesBtn" api:"required"`
	// Per-category sections plus a leading general section. The server backfills
	// sections for each defined category if absent.
	Sections []ConsentSettingDeleteResponseRegionRuleTranslationValuePreferencesModalSection `json:"sections" api:"required"`
	// Heading for the preferences modal.
	Title string `json:"title" api:"required"`
	// Optional collapsible info block rendered after the category sections. Renders
	// only when both title and body are non-empty. Use it for long-form text unrelated
	// to any specific category (e.g. a third-party advertisers disclosure).
	AdditionalInfo ConsentSettingDeleteResponseRegionRuleTranslationValuePreferencesModalAdditionalInfo `json:"additionalInfo" api:"nullable"`
	// Pluralizable label for the "N services" counter (e.g. "Service|Services").
	ServiceCounterLabel string `json:"serviceCounterLabel" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptAllBtn        respjson.Field
		CloseIconLabel      respjson.Field
		RejectAllBtn        respjson.Field
		SavePreferencesBtn  respjson.Field
		Sections            respjson.Field
		Title               respjson.Field
		AdditionalInfo      respjson.Field
		ServiceCounterLabel respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseRegionRuleTranslationValuePreferencesModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingDeleteResponseRegionRuleTranslationValuePreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingDeleteResponseRegionRuleTranslationValuePreferencesModalSection struct {
	// Body text for this preferences-modal section.
	Description string `json:"description" api:"required"`
	// Heading for this preferences-modal section.
	Title string `json:"title" api:"required"`
	// Optional disclosure table listing cookies/services in this section.
	CookieTable ConsentSettingDeleteResponseRegionRuleTranslationValuePreferencesModalSectionCookieTable `json:"cookieTable" api:"nullable"`
	// Category value this section configures. Use null/omit for the introductory
	// section that applies to no specific category.
	LinkedCategory string `json:"linkedCategory" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description    respjson.Field
		Title          respjson.Field
		CookieTable    respjson.Field
		LinkedCategory respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseRegionRuleTranslationValuePreferencesModalSection) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingDeleteResponseRegionRuleTranslationValuePreferencesModalSection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional disclosure table listing cookies/services in this section.
type ConsentSettingDeleteResponseRegionRuleTranslationValuePreferencesModalSectionCookieTable struct {
	// Rows of the cookie table — each row is a {key,value} pair (e.g. cookie name →
	// purpose).
	Body []ConsentSettingDeleteResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableBody `json:"body" api:"required"`
	// Header cells. Same {key,value} shape as rows.
	Headers []ConsentSettingDeleteResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader `json:"headers" api:"required"`
	// Optional table caption.
	Caption string `json:"caption" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Headers     respjson.Field
		Caption     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseRegionRuleTranslationValuePreferencesModalSectionCookieTable) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingDeleteResponseRegionRuleTranslationValuePreferencesModalSectionCookieTable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingDeleteResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableBody struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableBody) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingDeleteResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingDeleteResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingDeleteResponseRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional collapsible info block rendered after the category sections. Renders
// only when both title and body are non-empty. Use it for long-form text unrelated
// to any specific category (e.g. a third-party advertisers disclosure).
type ConsentSettingDeleteResponseRegionRuleTranslationValuePreferencesModalAdditionalInfo struct {
	// Body text for the additional-info block. Supports the rich-text markdown subset
	// used by section descriptions: line breaks, [text](url) links (https:// or
	// relative paths), and **bold**. HTML is stripped.
	Body string `json:"body" api:"required"`
	// Heading for the additional-info block.
	Title string `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body        respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseRegionRuleTranslationValuePreferencesModalAdditionalInfo) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingDeleteResponseRegionRuleTranslationValuePreferencesModalAdditionalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Threshold config for autoShowDismissMode (page count or seconds).
type ConsentSettingDeleteResponseRegionRuleAutoShowDismissConfig struct {
	// Pages-viewed threshold when autoShowDismissMode = after_pages.
	PageCount int64 `json:"pageCount" api:"nullable"`
	// Seconds-on-page threshold when autoShowDismissMode = after_seconds.
	Seconds int64 `json:"seconds" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageCount   respjson.Field
		Seconds     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseRegionRuleAutoShowDismissConfig) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingDeleteResponseRegionRuleAutoShowDismissConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the modals (layout/position/colors).
type ConsentSettingDeleteResponseRegionRuleGuiOptions struct {
	// Visual options for the initial consent modal.
	ConsentModal ConsentSettingDeleteResponseRegionRuleGuiOptionsConsentModal `json:"consentModal" api:"nullable"`
	// Global CSS variable overrides applied to both modals.
	CssVariables ConsentSettingDeleteResponseRegionRuleGuiOptionsCssVariables `json:"cssVariables" api:"nullable"`
	// Visual options for the per-category preferences modal.
	PreferencesModal ConsentSettingDeleteResponseRegionRuleGuiOptionsPreferencesModal `json:"preferencesModal" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConsentModal     respjson.Field
		CssVariables     respjson.Field
		PreferencesModal respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseRegionRuleGuiOptions) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingDeleteResponseRegionRuleGuiOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the initial consent modal.
type ConsentSettingDeleteResponseRegionRuleGuiOptionsConsentModal struct {
	// Which buttons to render in the consent modal and in what arrangement.
	//
	// Any of "AcceptAllNecessaryPreferences", "AcceptAllPreferences", "AcceptOnly",
	// "InformationOnly", "PreferencesOnly".
	ButtonLayout string `json:"buttonLayout" api:"nullable"`
	// When true, accept and reject buttons share the same visual weight (no
	// primary/secondary).
	EqualWeightButtons bool `json:"equalWeightButtons" api:"nullable"`
	// When true, swap accept/reject button order.
	FlipButtons bool `json:"flipButtons" api:"nullable"`
	// Modal layout style (box, bar, cloud, etc.).
	//
	// Any of "bar", "bar_inline", "box", "box_inline", "box_wide", "cloud",
	// "cloud_inline".
	Layout string `json:"layout" api:"nullable"`
	// Where the modal anchors on the viewport (e.g. bottom_left, top_right,
	// middle_center).
	//
	// Any of "bottom", "bottom_center", "bottom_left", "bottom_right", "middle",
	// "middle_center", "middle_left", "middle_right", "top", "top_center", "top_left",
	// "top_right".
	Position string `json:"position" api:"nullable"`
	// Show an X to dismiss the modal without choosing.
	ShowCloseIcon bool `json:"showCloseIcon" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonLayout       respjson.Field
		EqualWeightButtons respjson.Field
		FlipButtons        respjson.Field
		Layout             respjson.Field
		Position           respjson.Field
		ShowCloseIcon      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseRegionRuleGuiOptionsConsentModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingDeleteResponseRegionRuleGuiOptionsConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Global CSS variable overrides applied to both modals.
type ConsentSettingDeleteResponseRegionRuleGuiOptionsCssVariables struct {
	// CSS length, e.g. "4px" or "0.25rem".
	ButtonBorderRadius string `json:"buttonBorderRadius" api:"nullable"`
	// CSS color for the per-category block background in the preferences modal.
	CookieCategoryBlockBg string `json:"cookieCategoryBlockBg" api:"nullable"`
	// CSS color for the per-category block background on hover.
	CookieCategoryBlockHoverBg string `json:"cookieCategoryBlockHoverBg" api:"nullable"`
	// CSS color for the modal footer background.
	FooterBg string `json:"footerBg" api:"nullable"`
	// CSS color for footer text.
	FooterColor string `json:"footerColor" api:"nullable"`
	// CSS color for footer links.
	FooterLinkColor string `json:"footerLinkColor" api:"nullable"`
	// CSS color for footer links on hover.
	FooterLinkHoverColor string `json:"footerLinkHoverColor" api:"nullable"`
	// CSS color for the modal background.
	ModalBg string `json:"modalBg" api:"nullable"`
	// CSS length for modal corner radius.
	ModalBorderRadius string `json:"modalBorderRadius" api:"nullable"`
	// CSS color for the primary button background.
	PrimaryButtonBg string `json:"primaryButtonBg" api:"nullable"`
	// CSS color for primary button text.
	PrimaryButtonColor string `json:"primaryButtonColor" api:"nullable"`
	// CSS color for the primary button background on hover.
	PrimaryButtonHoverBg string `json:"primaryButtonHoverBg" api:"nullable"`
	// CSS color for primary button text on hover.
	PrimaryButtonHoverColor string `json:"primaryButtonHoverColor" api:"nullable"`
	// CSS color for primary text in the modal body.
	PrimaryTextColor string `json:"primaryTextColor" api:"nullable"`
	// CSS color for the secondary button background.
	SecondaryButtonBg string `json:"secondaryButtonBg" api:"nullable"`
	// CSS color for secondary button text.
	SecondaryButtonColor string `json:"secondaryButtonColor" api:"nullable"`
	// CSS color for secondary button background on hover.
	SecondaryButtonHoverBg string `json:"secondaryButtonHoverBg" api:"nullable"`
	// CSS color for secondary button text on hover.
	SecondaryButtonHoverColor string `json:"secondaryButtonHoverColor" api:"nullable"`
	// CSS color for secondary text in the modal body.
	SecondaryTextColor string `json:"secondaryTextColor" api:"nullable"`
	// CSS color for the "on" state of category toggles.
	ToggleOnBg string `json:"toggleOnBg" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonBorderRadius         respjson.Field
		CookieCategoryBlockBg      respjson.Field
		CookieCategoryBlockHoverBg respjson.Field
		FooterBg                   respjson.Field
		FooterColor                respjson.Field
		FooterLinkColor            respjson.Field
		FooterLinkHoverColor       respjson.Field
		ModalBg                    respjson.Field
		ModalBorderRadius          respjson.Field
		PrimaryButtonBg            respjson.Field
		PrimaryButtonColor         respjson.Field
		PrimaryButtonHoverBg       respjson.Field
		PrimaryButtonHoverColor    respjson.Field
		PrimaryTextColor           respjson.Field
		SecondaryButtonBg          respjson.Field
		SecondaryButtonColor       respjson.Field
		SecondaryButtonHoverBg     respjson.Field
		SecondaryButtonHoverColor  respjson.Field
		SecondaryTextColor         respjson.Field
		ToggleOnBg                 respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseRegionRuleGuiOptionsCssVariables) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingDeleteResponseRegionRuleGuiOptionsCssVariables) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the per-category preferences modal.
type ConsentSettingDeleteResponseRegionRuleGuiOptionsPreferencesModal struct {
	// Button layout for the preferences modal.
	//
	// Any of "AcceptAllRejectAllSave", "AcceptAllSave".
	ButtonLayout string `json:"buttonLayout" api:"nullable"`
	// When true, accept/reject share the same visual weight.
	EqualWeightButtons bool `json:"equalWeightButtons" api:"nullable"`
	// When true, swap accept/reject button order.
	FlipButtons bool `json:"flipButtons" api:"nullable"`
	// Preferences modal layout (box, bar, bar_wide).
	//
	// Any of "bar", "bar_wide", "box".
	Layout string `json:"layout" api:"nullable"`
	// Preferences modal anchor (left, right).
	//
	// Any of "left", "right".
	Position string `json:"position" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ButtonLayout       respjson.Field
		EqualWeightButtons respjson.Field
		FlipButtons        respjson.Field
		Layout             respjson.Field
		Position           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseRegionRuleGuiOptionsPreferencesModal) RawJSON() string {
	return r.JSON.raw
}
func (r *ConsentSettingDeleteResponseRegionRuleGuiOptionsPreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingDeleteResponseService struct {
	// Internal notes shown to admins in the dashboard. Not user-facing.
	InternalNotes string `json:"internalNotes" api:"required"`
	// Display name for this service in the preferences modal.
	Label string `json:"label" api:"required"`
	// Extra category values this service belongs to. Each must match a
	// `categories[].value`.
	AdditionalCategories []string `json:"additionalCategories" api:"nullable"`
	// Primary category value this service belongs to. Must match one of the top-level
	// `categories[].value` entries.
	Category string `json:"category" api:"nullable"`
	// Domains/paths this service matches. Patterns matching the CMP's own scripts
	// (e.g. cdn.oursprivacy.com/cmp-init) are rejected to prevent the CMP blocking
	// itself — use a more specific path like cdn.oursprivacy.com/main.js to block a
	// specific script.
	DomainPatterns []string `json:"domainPatterns" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InternalNotes        respjson.Field
		Label                respjson.Field
		AdditionalCategories respjson.Field
		Category             respjson.Field
		DomainPatterns       respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponseService) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingDeleteResponseService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enabled serves the CMP on its allowed domains and requires at least one.
// Disabled does not serve the CMP and may have no allowed domains.
type ConsentSettingDeleteResponseStatus string

const (
	ConsentSettingDeleteResponseStatusDisabled ConsentSettingDeleteResponseStatus = "Disabled"
	ConsentSettingDeleteResponseStatusEnabled  ConsentSettingDeleteResponseStatus = "Enabled"
)

type ConsentSettingAnalyticsResponse struct {
	// One entry per time bucket (hour, day, week, or month, depending on
	// `granularity`) covering the full window — empty windows are zero-filled so
	// callers can render contiguous time series without gap-handling logic.
	Items  []ConsentSettingAnalyticsResponseItem `json:"items" api:"required"`
	Totals ConsentSettingAnalyticsResponseTotals `json:"totals" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Totals      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingAnalyticsResponse) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingAnalyticsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingAnalyticsResponseItem struct {
	BannerViews                     int64   `json:"bannerViews" api:"required"`
	CloseIconClicks                 int64   `json:"closeIconClicks" api:"required"`
	DateTime                        string  `json:"dateTime" api:"required"`
	ExplicitOptIns                  int64   `json:"explicitOptIns" api:"required"`
	ExplicitOptOuts                 int64   `json:"explicitOptOuts" api:"required"`
	OptInRate                       float64 `json:"optInRate" api:"required"`
	OptOutRate                      float64 `json:"optOutRate" api:"required"`
	PartialPeriod                   bool    `json:"partialPeriod" api:"required"`
	PeriodFrom                      string  `json:"periodFrom" api:"required"`
	PeriodTo                        string  `json:"periodTo" api:"required"`
	PercentageChangeBannerViews     float64 `json:"percentageChangeBannerViews" api:"nullable"`
	PercentageChangeCloseIconClicks float64 `json:"percentageChangeCloseIconClicks" api:"nullable"`
	PercentageChangeExplicitOptIns  float64 `json:"percentageChangeExplicitOptIns" api:"nullable"`
	PercentageChangeExplicitOptOuts float64 `json:"percentageChangeExplicitOptOuts" api:"nullable"`
	PreviousBannerViews             float64 `json:"previousBannerViews" api:"nullable"`
	PreviousCloseIconClicks         float64 `json:"previousCloseIconClicks" api:"nullable"`
	PreviousExplicitOptIns          float64 `json:"previousExplicitOptIns" api:"nullable"`
	PreviousExplicitOptOuts         float64 `json:"previousExplicitOptOuts" api:"nullable"`
	PreviousOptInRate               float64 `json:"previousOptInRate" api:"nullable"`
	PreviousOptOutRate              float64 `json:"previousOptOutRate" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BannerViews                     respjson.Field
		CloseIconClicks                 respjson.Field
		DateTime                        respjson.Field
		ExplicitOptIns                  respjson.Field
		ExplicitOptOuts                 respjson.Field
		OptInRate                       respjson.Field
		OptOutRate                      respjson.Field
		PartialPeriod                   respjson.Field
		PeriodFrom                      respjson.Field
		PeriodTo                        respjson.Field
		PercentageChangeBannerViews     respjson.Field
		PercentageChangeCloseIconClicks respjson.Field
		PercentageChangeExplicitOptIns  respjson.Field
		PercentageChangeExplicitOptOuts respjson.Field
		PreviousBannerViews             respjson.Field
		PreviousCloseIconClicks         respjson.Field
		PreviousExplicitOptIns          respjson.Field
		PreviousExplicitOptOuts         respjson.Field
		PreviousOptInRate               respjson.Field
		PreviousOptOutRate              respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingAnalyticsResponseItem) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingAnalyticsResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingAnalyticsResponseTotals struct {
	BannerViews     int64   `json:"bannerViews" api:"required"`
	CloseIconClicks int64   `json:"closeIconClicks" api:"required"`
	ExplicitOptIns  int64   `json:"explicitOptIns" api:"required"`
	ExplicitOptOuts int64   `json:"explicitOptOuts" api:"required"`
	OptInRate       float64 `json:"optInRate" api:"required"`
	OptOutRate      float64 `json:"optOutRate" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BannerViews     respjson.Field
		CloseIconClicks respjson.Field
		ExplicitOptIns  respjson.Field
		ExplicitOptOuts respjson.Field
		OptInRate       respjson.Field
		OptOutRate      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingAnalyticsResponseTotals) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingAnalyticsResponseTotals) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingOverviewResponse struct {
	Items []ConsentSettingOverviewResponseItem `json:"items" api:"required"`
	// Whole-window counts; do not sum daily unique visitors.
	Totals ConsentSettingOverviewResponseTotals `json:"totals" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Totals      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingOverviewResponse) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingOverviewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingOverviewResponseItem struct {
	// Unique visitors shown the banner.
	BannerViews int64 `json:"bannerViews" api:"required"`
	// Banner dismissal events.
	CloseIconClicks int64     `json:"closeIconClicks" api:"required"`
	DateTime        time.Time `json:"dateTime" api:"required" format:"date-time"`
	// Unique visitors who accepted all.
	ExplicitOptIns int64 `json:"explicitOptIns" api:"required"`
	// Unique visitors who accepted necessary categories only.
	ExplicitOptOuts int64 `json:"explicitOptOuts" api:"required"`
	// Opt-in percentage using the Consent Analytics denominator.
	OptInRate float64 `json:"optInRate" api:"required"`
	// Opt-out percentage using the Consent Analytics denominator.
	OptOutRate    float64   `json:"optOutRate" api:"required"`
	PartialPeriod bool      `json:"partialPeriod" api:"required"`
	PeriodFrom    time.Time `json:"periodFrom" api:"required" format:"date"`
	PeriodTo      time.Time `json:"periodTo" api:"required" format:"date"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BannerViews     respjson.Field
		CloseIconClicks respjson.Field
		DateTime        respjson.Field
		ExplicitOptIns  respjson.Field
		ExplicitOptOuts respjson.Field
		OptInRate       respjson.Field
		OptOutRate      respjson.Field
		PartialPeriod   respjson.Field
		PeriodFrom      respjson.Field
		PeriodTo        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingOverviewResponseItem) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingOverviewResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whole-window counts; do not sum daily unique visitors.
type ConsentSettingOverviewResponseTotals struct {
	// Unique visitors shown the banner.
	BannerViews int64 `json:"bannerViews" api:"required"`
	// Banner dismissal events.
	CloseIconClicks int64 `json:"closeIconClicks" api:"required"`
	// Unique visitors who accepted all.
	ExplicitOptIns int64 `json:"explicitOptIns" api:"required"`
	// Unique visitors who accepted necessary categories only.
	ExplicitOptOuts int64 `json:"explicitOptOuts" api:"required"`
	// Opt-in percentage using the Consent Analytics denominator.
	OptInRate float64 `json:"optInRate" api:"required"`
	// Opt-out percentage using the Consent Analytics denominator.
	OptOutRate float64 `json:"optOutRate" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BannerViews     respjson.Field
		CloseIconClicks respjson.Field
		ExplicitOptIns  respjson.Field
		ExplicitOptOuts respjson.Field
		OptInRate       respjson.Field
		OptOutRate      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingOverviewResponseTotals) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingOverviewResponseTotals) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingAnalyticsMonthlyResponse struct {
	Items []ConsentSettingAnalyticsMonthlyResponseItem `json:"items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingAnalyticsMonthlyResponse) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingAnalyticsMonthlyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingAnalyticsMonthlyResponseItem struct {
	BannerViews     int64   `json:"bannerViews" api:"required"`
	CloseIconClicks int64   `json:"closeIconClicks" api:"required"`
	DateTime        string  `json:"dateTime" api:"required"`
	ExplicitOptIns  int64   `json:"explicitOptIns" api:"required"`
	ExplicitOptOuts int64   `json:"explicitOptOuts" api:"required"`
	OptInRate       float64 `json:"optInRate" api:"required"`
	OptOutRate      float64 `json:"optOutRate" api:"required"`
	PartialPeriod   bool    `json:"partialPeriod" api:"required"`
	PeriodFrom      string  `json:"periodFrom" api:"required"`
	PeriodTo        string  `json:"periodTo" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BannerViews     respjson.Field
		CloseIconClicks respjson.Field
		DateTime        respjson.Field
		ExplicitOptIns  respjson.Field
		ExplicitOptOuts respjson.Field
		OptInRate       respjson.Field
		OptOutRate      respjson.Field
		PartialPeriod   respjson.Field
		PeriodFrom      respjson.Field
		PeriodTo        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingAnalyticsMonthlyResponseItem) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingAnalyticsMonthlyResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingAnalyticsCapabilitiesResponse struct {
	ExpiresAt     string                                                `json:"expiresAt" api:"required"`
	Intervals     []ConsentSettingAnalyticsCapabilitiesResponseInterval `json:"intervals" api:"required"`
	PolicyVersion string                                                `json:"policyVersion" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt     respjson.Field
		Intervals     respjson.Field
		PolicyVersion respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingAnalyticsCapabilitiesResponse) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingAnalyticsCapabilitiesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingAnalyticsCapabilitiesResponseInterval struct {
	ComparisonSupported bool `json:"comparisonSupported" api:"required"`
	// Any of "DAILY", "HOURLY", "MONTHLY", "WEEKLY".
	Granularity     string `json:"granularity" api:"required"`
	MaxIncludedDays int64  `json:"maxIncludedDays" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ComparisonSupported respjson.Field
		Granularity         respjson.Field
		MaxIncludedDays     respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingAnalyticsCapabilitiesResponseInterval) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingAnalyticsCapabilitiesResponseInterval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingPageAnalysisResponse struct {
	// Pages ranked by opt-outs descending, then page path ascending.
	Entities   []ConsentSettingPageAnalysisResponseEntity   `json:"entities" api:"required"`
	Pagination ConsentSettingPageAnalysisResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingPageAnalysisResponse) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingPageAnalysisResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingPageAnalysisResponseEntity struct {
	BannerViews     int64   `json:"bannerViews" api:"required"`
	CloseIconClicks int64   `json:"closeIconClicks" api:"required"`
	OptOutRate      float64 `json:"optOutRate" api:"required"`
	OptOuts         int64   `json:"optOuts" api:"required"`
	Page            string  `json:"page" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BannerViews     respjson.Field
		CloseIconClicks respjson.Field
		OptOutRate      respjson.Field
		OptOuts         respjson.Field
		Page            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingPageAnalysisResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingPageAnalysisResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingPageAnalysisResponsePagination struct {
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
func (r ConsentSettingPageAnalysisResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingPageAnalysisResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingAnalyticsByRegionResponse struct {
	// Per-region totals for the window, ranked by banner views (descending). Visitors
	// whose region cannot be resolved (e.g. bot traffic, IP geo failure) are bucketed
	// under the literal `Unknown`, so the per-region counts always sum to the global
	// totals.
	Regions []ConsentSettingAnalyticsByRegionResponseRegion `json:"regions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Regions     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingAnalyticsByRegionResponse) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingAnalyticsByRegionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingAnalyticsByRegionResponseRegion struct {
	BannerViews     int64   `json:"bannerViews" api:"required"`
	CloseIconClicks int64   `json:"closeIconClicks" api:"required"`
	ExplicitOptIns  int64   `json:"explicitOptIns" api:"required"`
	ExplicitOptOuts int64   `json:"explicitOptOuts" api:"required"`
	OptInRate       float64 `json:"optInRate" api:"required"`
	OptOutRate      float64 `json:"optOutRate" api:"required"`
	RegionName      string  `json:"regionName" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BannerViews     respjson.Field
		CloseIconClicks respjson.Field
		ExplicitOptIns  respjson.Field
		ExplicitOptOuts respjson.Field
		OptInRate       respjson.Field
		OptOutRate      respjson.Field
		RegionName      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingAnalyticsByRegionResponseRegion) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingAnalyticsByRegionResponseRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingReplaceParams struct {
	// Top-level consent categories. Server re-stamps `priority` to 0..N.
	Categories []ConsentSettingReplaceParamsCategory `json:"categories,omitzero" api:"required"`
	// Default rule used when the user is not in any region listed in `regions[]`.
	Default ConsentSettingReplaceParamsDefault `json:"default,omitzero" api:"required"`
	// Human-readable name shown in the dashboard.
	Name string `json:"name" api:"required"`
	// Per-region rule overrides. Each `regionCode` must be unique across rules and
	// must not appear in any other rule's `additionalRegions`.
	Regions []ConsentSettingReplaceParamsRegion `json:"regions,omitzero" api:"required"`
	// Per-service entries powering "show vendors" and category-aware blocking. Empty
	// array clears the list.
	Services []ConsentSettingReplaceParamsService `json:"services,omitzero" api:"required"`
	// Enabled requires at least one allowed domain. Disabled may have none.
	//
	// Any of "Disabled", "Enabled".
	Status ConsentSettingReplaceParamsStatus `json:"status,omitzero" api:"required"`
	// Name of the cookie that stores consent state. Pass null to clear (defaults to
	// "op_consent").
	ConsentCookieName param.Opt[string] `json:"consentCookieName,omitzero"`
	// Custom CDN domain for serving the CMP script. Pass null to clear.
	CustomDomain param.Opt[string] `json:"customDomain,omitzero"`
	// Cookie/localStorage key for the visitor (device) ID. Must exactly match the web
	// SDK's cookie_names.device_id (default "ours_device_id"). Pass null to clear.
	DeviceIDCookieName param.Opt[string] `json:"deviceIdCookieName,omitzero"`
	// Revision counter. Bump to re-prompt users who already consented.
	Revision param.Opt[float64] `json:"revision,omitzero"`
	// Pixel of the WebSource this CMP is wired into. Pass null to clear the link.
	WebSDKToken param.Opt[string] `json:"webSDKToken,omitzero"`
	// CSS class names that opt scripts out of consent blocking. Each must be a single
	// class token.
	SkipBlockingClassNames []string `json:"skipBlockingClassNames,omitzero"`
	// Domains where this CMP may run. Enabled requires at least one allowed domain;
	// omitting, null, or [] fails validation when status is Enabled. Disabled may have
	// none.
	WhitelistDomains []string `json:"whitelistDomains,omitzero"`
	paramObj
}

func (r ConsentSettingReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Label, Priority, Value are required.
type ConsentSettingReplaceParamsCategory struct {
	// Human-readable label shown next to the toggle in the preferences modal.
	Label string `json:"label" api:"required"`
	// Sort key. Lower numbers render first. Server re-stamps to 0..N on write — send
	// any integer, gaps and duplicates are ironed out.
	Priority int64 `json:"priority" api:"required"`
	// Stable identifier referenced by services and translation sections.
	// Conventionally lowercase (e.g. "necessary", "analytics", "advertising").
	Value string `json:"value" api:"required"`
	paramObj
}

func (r ConsentSettingReplaceParamsCategory) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsCategory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default rule used when the user is not in any region listed in `regions[]`.
//
// The properties Categories, Language, Mode, Translations are required.
type ConsentSettingReplaceParamsDefault struct {
	// Per-category default config for this rule. Every category defined in the
	// top-level `categories[].value` should have an entry here.
	Categories []ConsentSettingReplaceParamsDefaultCategory `json:"categories,omitzero" api:"required"`
	// BCP 47 default language for this rule. Must have a matching entry in
	// `translations`. Examples: "en", "en-US", "es", "de".
	Language string `json:"language" api:"required"`
	// opt_in: scripts blocked until user accepts (GDPR style). opt_out: scripts run by
	// default until user rejects (CCPA style).
	//
	// Any of "opt_in", "opt_out".
	Mode string `json:"mode,omitzero" api:"required"`
	// All UI copy, keyed by language. Must include an entry whose `language` matches
	// the rule's `language` field.
	Translations []ConsentSettingReplaceParamsDefaultTranslation `json:"translations,omitzero" api:"required"`
	// When true, scripts not classified by services[] are blocked until the user opts
	// in.
	AutoblockUnknown param.Opt[bool] `json:"autoblockUnknown,omitzero"`
	// When true, the consent modal auto-opens on page load.
	AutoShow param.Opt[bool] `json:"autoShow,omitzero"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction param.Opt[bool] `json:"disablePageInteraction,omitzero"`
	// When true, the modal is suppressed for known bot user agents.
	HideFromBots param.Opt[bool] `json:"hideFromBots,omitzero"`
	// When true, the per-service list (services[]) is rendered inside the preferences
	// modal.
	ShowVendorsInPreferences param.Opt[bool] `json:"showVendorsInPreferences,omitzero"`
	// Threshold config for autoShowDismissMode (page count or seconds).
	AutoShowDismissConfig ConsentSettingReplaceParamsDefaultAutoShowDismissConfig `json:"autoShowDismissConfig,omitzero"`
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	//
	// Any of "after_pages", "after_seconds", "never".
	AutoShowDismissMode string `json:"autoShowDismissMode,omitzero"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions ConsentSettingReplaceParamsDefaultGuiOptions `json:"guiOptions,omitzero"`
	paramObj
}

func (r ConsentSettingReplaceParamsDefault) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsDefault
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsDefault) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConsentSettingReplaceParamsDefault](
		"mode", "opt_in", "opt_out",
	)
	apijson.RegisterFieldValidator[ConsentSettingReplaceParamsDefault](
		"autoShowDismissMode", "after_pages", "after_seconds", "never",
	)
}

// The properties Key, Value are required.
type ConsentSettingReplaceParamsDefaultCategory struct {
	// Category value (matches `categories[].value`) this entry configures.
	Key   string                                          `json:"key" api:"required"`
	Value ConsentSettingReplaceParamsDefaultCategoryValue `json:"value,omitzero" api:"required"`
	paramObj
}

func (r ConsentSettingReplaceParamsDefaultCategory) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsDefaultCategory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsDefaultCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Enabled is required.
type ConsentSettingReplaceParamsDefaultCategoryValue struct {
	// Whether this category is on by default before the user interacts.
	Enabled bool `json:"enabled" api:"required"`
	// When true, this category defaults off if the browser sends Sec-GPC: 1.
	AutoDisableOnGpc param.Opt[bool] `json:"autoDisableOnGPC,omitzero"`
	// When true, the user cannot toggle this category in the preferences modal.
	ReadOnly param.Opt[bool] `json:"readOnly,omitzero"`
	// When true, the page reloads after this category is toggled so newly-allowed
	// scripts can run.
	ReloadPage param.Opt[bool] `json:"reloadPage,omitzero"`
	paramObj
}

func (r ConsentSettingReplaceParamsDefaultCategoryValue) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsDefaultCategoryValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsDefaultCategoryValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Language, Value are required.
type ConsentSettingReplaceParamsDefaultTranslation struct {
	// BCP 47 language tag identifying which translation this entry provides. Examples:
	// "en", "en-US", "es", "fr-CA". The default rule's `language` must appear here.
	Language string                                             `json:"language" api:"required"`
	Value    ConsentSettingReplaceParamsDefaultTranslationValue `json:"value,omitzero" api:"required"`
	paramObj
}

func (r ConsentSettingReplaceParamsDefaultTranslation) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsDefaultTranslation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsDefaultTranslation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingReplaceParamsDefaultTranslationValue struct {
	// Translated copy for the initial consent modal.
	ConsentModal ConsentSettingReplaceParamsDefaultTranslationValueConsentModal `json:"consentModal,omitzero"`
	// Translated copy for the preferences modal.
	PreferencesModal ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModal `json:"preferencesModal,omitzero"`
	paramObj
}

func (r ConsentSettingReplaceParamsDefaultTranslationValue) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsDefaultTranslationValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsDefaultTranslationValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the initial consent modal.
//
// The properties AcceptAllBtn, Description, RejectAllBtn, ShowPreferencesBtn,
// Title are required.
type ConsentSettingReplaceParamsDefaultTranslationValueConsentModal struct {
	// Label for the "accept all" button.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Body text shown in the consent modal.
	Description string `json:"description" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the button that opens the preferences modal.
	ShowPreferencesBtn string `json:"showPreferencesBtn" api:"required"`
	// Heading shown at the top of the consent modal.
	Title string `json:"title" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel param.Opt[string] `json:"closeIconLabel,omitzero"`
	// Optional footer HTML/text shown below the body.
	Footer param.Opt[string] `json:"footer,omitzero"`
	// Optional text shown when the browser is sending Sec-GPC: 1.
	GpcNotification param.Opt[string] `json:"gpcNotification,omitzero"`
	// Link label for the privacy policy (defaults to "Privacy Policy").
	PrivacyPolicyLabel param.Opt[string] `json:"privacyPolicyLabel,omitzero"`
	// Privacy policy link. Accepts an absolute URL (https://...) or a relative path
	// (/legal/privacy) so one config can run across multiple domains. javascript:,
	// data:, and vbscript: schemes are rejected.
	PrivacyPolicyURL param.Opt[string] `json:"privacyPolicyUrl,omitzero"`
	// Optional message shown when the consent revision has bumped (re-prompt copy).
	RevisionMessage param.Opt[string] `json:"revisionMessage,omitzero"`
	// Link label for terms of service (defaults to "Terms of Service").
	TermsOfServiceLabel param.Opt[string] `json:"termsOfServiceLabel,omitzero"`
	// Terms of service link. Same scheme rules as privacyPolicyUrl: absolute or
	// relative path only.
	TermsOfServiceURL param.Opt[string] `json:"termsOfServiceUrl,omitzero"`
	paramObj
}

func (r ConsentSettingReplaceParamsDefaultTranslationValueConsentModal) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsDefaultTranslationValueConsentModal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsDefaultTranslationValueConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the preferences modal.
//
// The properties AcceptAllBtn, CloseIconLabel, RejectAllBtn, SavePreferencesBtn,
// Sections, Title are required.
type ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModal struct {
	// Label for the "accept all" button in the preferences modal.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel string `json:"closeIconLabel" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the "save my choices" button.
	SavePreferencesBtn string `json:"savePreferencesBtn" api:"required"`
	// Per-category sections plus a leading general section. The server backfills
	// sections for each defined category if absent.
	Sections []ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModalSection `json:"sections,omitzero" api:"required"`
	// Heading for the preferences modal.
	Title string `json:"title" api:"required"`
	// Pluralizable label for the "N services" counter (e.g. "Service|Services").
	ServiceCounterLabel param.Opt[string] `json:"serviceCounterLabel,omitzero"`
	// Optional collapsible info block rendered after the category sections. Renders
	// only when both title and body are non-empty. Use it for long-form text unrelated
	// to any specific category (e.g. a third-party advertisers disclosure).
	AdditionalInfo ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModalAdditionalInfo `json:"additionalInfo,omitzero"`
	paramObj
}

func (r ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModal) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Description, Title are required.
type ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModalSection struct {
	// Body text for this preferences-modal section.
	Description string `json:"description" api:"required"`
	// Heading for this preferences-modal section.
	Title string `json:"title" api:"required"`
	// Category value this section configures. Use null/omit for the introductory
	// section that applies to no specific category.
	LinkedCategory param.Opt[string] `json:"linkedCategory,omitzero"`
	// Optional disclosure table listing cookies/services in this section.
	CookieTable ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModalSectionCookieTable `json:"cookieTable,omitzero"`
	paramObj
}

func (r ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModalSection) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModalSection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModalSection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional disclosure table listing cookies/services in this section.
//
// The properties Body, Headers are required.
type ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModalSectionCookieTable struct {
	// Rows of the cookie table — each row is a {key,value} pair (e.g. cookie name →
	// purpose).
	Body []ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModalSectionCookieTableBody `json:"body,omitzero" api:"required"`
	// Header cells. Same {key,value} shape as rows.
	Headers []ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModalSectionCookieTableHeader `json:"headers,omitzero" api:"required"`
	// Optional table caption.
	Caption param.Opt[string] `json:"caption,omitzero"`
	paramObj
}

func (r ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModalSectionCookieTable) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModalSectionCookieTable
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModalSectionCookieTable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Key, Value are required.
type ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModalSectionCookieTableBody struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	paramObj
}

func (r ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModalSectionCookieTableBody) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModalSectionCookieTableBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModalSectionCookieTableBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Key, Value are required.
type ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModalSectionCookieTableHeader struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	paramObj
}

func (r ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModalSectionCookieTableHeader) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModalSectionCookieTableHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModalSectionCookieTableHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional collapsible info block rendered after the category sections. Renders
// only when both title and body are non-empty. Use it for long-form text unrelated
// to any specific category (e.g. a third-party advertisers disclosure).
//
// The properties Body, Title are required.
type ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModalAdditionalInfo struct {
	// Body text for the additional-info block. Supports the rich-text markdown subset
	// used by section descriptions: line breaks, [text](url) links (https:// or
	// relative paths), and **bold**. HTML is stripped.
	Body string `json:"body" api:"required"`
	// Heading for the additional-info block.
	Title string `json:"title" api:"required"`
	paramObj
}

func (r ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModalAdditionalInfo) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModalAdditionalInfo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModalAdditionalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Threshold config for autoShowDismissMode (page count or seconds).
type ConsentSettingReplaceParamsDefaultAutoShowDismissConfig struct {
	// Pages-viewed threshold when autoShowDismissMode = after_pages.
	PageCount param.Opt[int64] `json:"pageCount,omitzero"`
	// Seconds-on-page threshold when autoShowDismissMode = after_seconds.
	Seconds param.Opt[int64] `json:"seconds,omitzero"`
	paramObj
}

func (r ConsentSettingReplaceParamsDefaultAutoShowDismissConfig) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsDefaultAutoShowDismissConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsDefaultAutoShowDismissConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the modals (layout/position/colors).
type ConsentSettingReplaceParamsDefaultGuiOptions struct {
	// Visual options for the initial consent modal.
	ConsentModal ConsentSettingReplaceParamsDefaultGuiOptionsConsentModal `json:"consentModal,omitzero"`
	// Global CSS variable overrides applied to both modals.
	CssVariables ConsentSettingReplaceParamsDefaultGuiOptionsCssVariables `json:"cssVariables,omitzero"`
	// Visual options for the per-category preferences modal.
	PreferencesModal ConsentSettingReplaceParamsDefaultGuiOptionsPreferencesModal `json:"preferencesModal,omitzero"`
	paramObj
}

func (r ConsentSettingReplaceParamsDefaultGuiOptions) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsDefaultGuiOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsDefaultGuiOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the initial consent modal.
type ConsentSettingReplaceParamsDefaultGuiOptionsConsentModal struct {
	// When true, accept and reject buttons share the same visual weight (no
	// primary/secondary).
	EqualWeightButtons param.Opt[bool] `json:"equalWeightButtons,omitzero"`
	// When true, swap accept/reject button order.
	FlipButtons param.Opt[bool] `json:"flipButtons,omitzero"`
	// Show an X to dismiss the modal without choosing.
	ShowCloseIcon param.Opt[bool] `json:"showCloseIcon,omitzero"`
	// Which buttons to render in the consent modal and in what arrangement.
	//
	// Any of "AcceptAllNecessaryPreferences", "AcceptAllPreferences", "AcceptOnly",
	// "InformationOnly", "PreferencesOnly".
	ButtonLayout string `json:"buttonLayout,omitzero"`
	// Modal layout style (box, bar, cloud, etc.).
	//
	// Any of "bar", "bar_inline", "box", "box_inline", "box_wide", "cloud",
	// "cloud_inline".
	Layout string `json:"layout,omitzero"`
	// Where the modal anchors on the viewport (e.g. bottom_left, top_right,
	// middle_center).
	//
	// Any of "bottom", "bottom_center", "bottom_left", "bottom_right", "middle",
	// "middle_center", "middle_left", "middle_right", "top", "top_center", "top_left",
	// "top_right".
	Position string `json:"position,omitzero"`
	paramObj
}

func (r ConsentSettingReplaceParamsDefaultGuiOptionsConsentModal) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsDefaultGuiOptionsConsentModal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsDefaultGuiOptionsConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConsentSettingReplaceParamsDefaultGuiOptionsConsentModal](
		"buttonLayout", "AcceptAllNecessaryPreferences", "AcceptAllPreferences", "AcceptOnly", "InformationOnly", "PreferencesOnly",
	)
	apijson.RegisterFieldValidator[ConsentSettingReplaceParamsDefaultGuiOptionsConsentModal](
		"layout", "bar", "bar_inline", "box", "box_inline", "box_wide", "cloud", "cloud_inline",
	)
	apijson.RegisterFieldValidator[ConsentSettingReplaceParamsDefaultGuiOptionsConsentModal](
		"position", "bottom", "bottom_center", "bottom_left", "bottom_right", "middle", "middle_center", "middle_left", "middle_right", "top", "top_center", "top_left", "top_right",
	)
}

// Global CSS variable overrides applied to both modals.
type ConsentSettingReplaceParamsDefaultGuiOptionsCssVariables struct {
	// CSS length, e.g. "4px" or "0.25rem".
	ButtonBorderRadius param.Opt[string] `json:"buttonBorderRadius,omitzero"`
	// CSS color for the per-category block background in the preferences modal.
	CookieCategoryBlockBg param.Opt[string] `json:"cookieCategoryBlockBg,omitzero"`
	// CSS color for the per-category block background on hover.
	CookieCategoryBlockHoverBg param.Opt[string] `json:"cookieCategoryBlockHoverBg,omitzero"`
	// CSS color for the modal footer background.
	FooterBg param.Opt[string] `json:"footerBg,omitzero"`
	// CSS color for footer text.
	FooterColor param.Opt[string] `json:"footerColor,omitzero"`
	// CSS color for footer links.
	FooterLinkColor param.Opt[string] `json:"footerLinkColor,omitzero"`
	// CSS color for footer links on hover.
	FooterLinkHoverColor param.Opt[string] `json:"footerLinkHoverColor,omitzero"`
	// CSS color for the modal background.
	ModalBg param.Opt[string] `json:"modalBg,omitzero"`
	// CSS length for modal corner radius.
	ModalBorderRadius param.Opt[string] `json:"modalBorderRadius,omitzero"`
	// CSS color for the primary button background.
	PrimaryButtonBg param.Opt[string] `json:"primaryButtonBg,omitzero"`
	// CSS color for primary button text.
	PrimaryButtonColor param.Opt[string] `json:"primaryButtonColor,omitzero"`
	// CSS color for the primary button background on hover.
	PrimaryButtonHoverBg param.Opt[string] `json:"primaryButtonHoverBg,omitzero"`
	// CSS color for primary button text on hover.
	PrimaryButtonHoverColor param.Opt[string] `json:"primaryButtonHoverColor,omitzero"`
	// CSS color for primary text in the modal body.
	PrimaryTextColor param.Opt[string] `json:"primaryTextColor,omitzero"`
	// CSS color for the secondary button background.
	SecondaryButtonBg param.Opt[string] `json:"secondaryButtonBg,omitzero"`
	// CSS color for secondary button text.
	SecondaryButtonColor param.Opt[string] `json:"secondaryButtonColor,omitzero"`
	// CSS color for secondary button background on hover.
	SecondaryButtonHoverBg param.Opt[string] `json:"secondaryButtonHoverBg,omitzero"`
	// CSS color for secondary button text on hover.
	SecondaryButtonHoverColor param.Opt[string] `json:"secondaryButtonHoverColor,omitzero"`
	// CSS color for secondary text in the modal body.
	SecondaryTextColor param.Opt[string] `json:"secondaryTextColor,omitzero"`
	// CSS color for the "on" state of category toggles.
	ToggleOnBg param.Opt[string] `json:"toggleOnBg,omitzero"`
	paramObj
}

func (r ConsentSettingReplaceParamsDefaultGuiOptionsCssVariables) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsDefaultGuiOptionsCssVariables
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsDefaultGuiOptionsCssVariables) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the per-category preferences modal.
type ConsentSettingReplaceParamsDefaultGuiOptionsPreferencesModal struct {
	// When true, accept/reject share the same visual weight.
	EqualWeightButtons param.Opt[bool] `json:"equalWeightButtons,omitzero"`
	// When true, swap accept/reject button order.
	FlipButtons param.Opt[bool] `json:"flipButtons,omitzero"`
	// Button layout for the preferences modal.
	//
	// Any of "AcceptAllRejectAllSave", "AcceptAllSave".
	ButtonLayout string `json:"buttonLayout,omitzero"`
	// Preferences modal layout (box, bar, bar_wide).
	//
	// Any of "bar", "bar_wide", "box".
	Layout string `json:"layout,omitzero"`
	// Preferences modal anchor (left, right).
	//
	// Any of "left", "right".
	Position string `json:"position,omitzero"`
	paramObj
}

func (r ConsentSettingReplaceParamsDefaultGuiOptionsPreferencesModal) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsDefaultGuiOptionsPreferencesModal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsDefaultGuiOptionsPreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConsentSettingReplaceParamsDefaultGuiOptionsPreferencesModal](
		"buttonLayout", "AcceptAllRejectAllSave", "AcceptAllSave",
	)
	apijson.RegisterFieldValidator[ConsentSettingReplaceParamsDefaultGuiOptionsPreferencesModal](
		"layout", "bar", "bar_wide", "box",
	)
	apijson.RegisterFieldValidator[ConsentSettingReplaceParamsDefaultGuiOptionsPreferencesModal](
		"position", "left", "right",
	)
}

// The properties RegionCode, Rule are required.
type ConsentSettingReplaceParamsRegion struct {
	// Region this rule applies to. Use ISO 3166-1 alpha-2 country code ("US", "DE",
	// "BR") or country-subdivision code ("US-CA", "GB-ENG", "CA-ON"). Each region code
	// may appear in only one rule across `regions[]`.
	RegionCode string                                `json:"regionCode" api:"required"`
	Rule       ConsentSettingReplaceParamsRegionRule `json:"rule,omitzero" api:"required"`
	// Optional label for this override group, e.g. "CCPA States". Shown in the
	// dashboard instead of the region codes. Pass null to clear it and fall back to
	// the codes.
	Name param.Opt[string] `json:"name,omitzero"`
	// Other region codes that should reuse this rule. Same code-format rules as
	// `regionCode`. Cannot include `regionCode` itself, cannot duplicate, cannot
	// overlap with another rule's regions.
	AdditionalRegions []string `json:"additionalRegions,omitzero"`
	paramObj
}

func (r ConsentSettingReplaceParamsRegion) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsRegion
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Categories, Language, Mode, Translations are required.
type ConsentSettingReplaceParamsRegionRule struct {
	// Per-category default config for this rule. Every category defined in the
	// top-level `categories[].value` should have an entry here.
	Categories []ConsentSettingReplaceParamsRegionRuleCategory `json:"categories,omitzero" api:"required"`
	// BCP 47 default language for this rule. Must have a matching entry in
	// `translations`. Examples: "en", "en-US", "es", "de".
	Language string `json:"language" api:"required"`
	// opt_in: scripts blocked until user accepts (GDPR style). opt_out: scripts run by
	// default until user rejects (CCPA style).
	//
	// Any of "opt_in", "opt_out".
	Mode string `json:"mode,omitzero" api:"required"`
	// All UI copy, keyed by language. Must include an entry whose `language` matches
	// the rule's `language` field.
	Translations []ConsentSettingReplaceParamsRegionRuleTranslation `json:"translations,omitzero" api:"required"`
	// When true, scripts not classified by services[] are blocked until the user opts
	// in.
	AutoblockUnknown param.Opt[bool] `json:"autoblockUnknown,omitzero"`
	// When true, the consent modal auto-opens on page load.
	AutoShow param.Opt[bool] `json:"autoShow,omitzero"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction param.Opt[bool] `json:"disablePageInteraction,omitzero"`
	// When true, the modal is suppressed for known bot user agents.
	HideFromBots param.Opt[bool] `json:"hideFromBots,omitzero"`
	// When true, the per-service list (services[]) is rendered inside the preferences
	// modal.
	ShowVendorsInPreferences param.Opt[bool] `json:"showVendorsInPreferences,omitzero"`
	// Threshold config for autoShowDismissMode (page count or seconds).
	AutoShowDismissConfig ConsentSettingReplaceParamsRegionRuleAutoShowDismissConfig `json:"autoShowDismissConfig,omitzero"`
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	//
	// Any of "after_pages", "after_seconds", "never".
	AutoShowDismissMode string `json:"autoShowDismissMode,omitzero"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions ConsentSettingReplaceParamsRegionRuleGuiOptions `json:"guiOptions,omitzero"`
	paramObj
}

func (r ConsentSettingReplaceParamsRegionRule) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsRegionRule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsRegionRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConsentSettingReplaceParamsRegionRule](
		"mode", "opt_in", "opt_out",
	)
	apijson.RegisterFieldValidator[ConsentSettingReplaceParamsRegionRule](
		"autoShowDismissMode", "after_pages", "after_seconds", "never",
	)
}

// The properties Key, Value are required.
type ConsentSettingReplaceParamsRegionRuleCategory struct {
	// Category value (matches `categories[].value`) this entry configures.
	Key   string                                             `json:"key" api:"required"`
	Value ConsentSettingReplaceParamsRegionRuleCategoryValue `json:"value,omitzero" api:"required"`
	paramObj
}

func (r ConsentSettingReplaceParamsRegionRuleCategory) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsRegionRuleCategory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsRegionRuleCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Enabled is required.
type ConsentSettingReplaceParamsRegionRuleCategoryValue struct {
	// Whether this category is on by default before the user interacts.
	Enabled bool `json:"enabled" api:"required"`
	// When true, this category defaults off if the browser sends Sec-GPC: 1.
	AutoDisableOnGpc param.Opt[bool] `json:"autoDisableOnGPC,omitzero"`
	// When true, the user cannot toggle this category in the preferences modal.
	ReadOnly param.Opt[bool] `json:"readOnly,omitzero"`
	// When true, the page reloads after this category is toggled so newly-allowed
	// scripts can run.
	ReloadPage param.Opt[bool] `json:"reloadPage,omitzero"`
	paramObj
}

func (r ConsentSettingReplaceParamsRegionRuleCategoryValue) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsRegionRuleCategoryValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsRegionRuleCategoryValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Language, Value are required.
type ConsentSettingReplaceParamsRegionRuleTranslation struct {
	// BCP 47 language tag identifying which translation this entry provides. Examples:
	// "en", "en-US", "es", "fr-CA". The default rule's `language` must appear here.
	Language string                                                `json:"language" api:"required"`
	Value    ConsentSettingReplaceParamsRegionRuleTranslationValue `json:"value,omitzero" api:"required"`
	paramObj
}

func (r ConsentSettingReplaceParamsRegionRuleTranslation) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsRegionRuleTranslation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsRegionRuleTranslation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingReplaceParamsRegionRuleTranslationValue struct {
	// Translated copy for the initial consent modal.
	ConsentModal ConsentSettingReplaceParamsRegionRuleTranslationValueConsentModal `json:"consentModal,omitzero"`
	// Translated copy for the preferences modal.
	PreferencesModal ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModal `json:"preferencesModal,omitzero"`
	paramObj
}

func (r ConsentSettingReplaceParamsRegionRuleTranslationValue) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsRegionRuleTranslationValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsRegionRuleTranslationValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the initial consent modal.
//
// The properties AcceptAllBtn, Description, RejectAllBtn, ShowPreferencesBtn,
// Title are required.
type ConsentSettingReplaceParamsRegionRuleTranslationValueConsentModal struct {
	// Label for the "accept all" button.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Body text shown in the consent modal.
	Description string `json:"description" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the button that opens the preferences modal.
	ShowPreferencesBtn string `json:"showPreferencesBtn" api:"required"`
	// Heading shown at the top of the consent modal.
	Title string `json:"title" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel param.Opt[string] `json:"closeIconLabel,omitzero"`
	// Optional footer HTML/text shown below the body.
	Footer param.Opt[string] `json:"footer,omitzero"`
	// Optional text shown when the browser is sending Sec-GPC: 1.
	GpcNotification param.Opt[string] `json:"gpcNotification,omitzero"`
	// Link label for the privacy policy (defaults to "Privacy Policy").
	PrivacyPolicyLabel param.Opt[string] `json:"privacyPolicyLabel,omitzero"`
	// Privacy policy link. Accepts an absolute URL (https://...) or a relative path
	// (/legal/privacy) so one config can run across multiple domains. javascript:,
	// data:, and vbscript: schemes are rejected.
	PrivacyPolicyURL param.Opt[string] `json:"privacyPolicyUrl,omitzero"`
	// Optional message shown when the consent revision has bumped (re-prompt copy).
	RevisionMessage param.Opt[string] `json:"revisionMessage,omitzero"`
	// Link label for terms of service (defaults to "Terms of Service").
	TermsOfServiceLabel param.Opt[string] `json:"termsOfServiceLabel,omitzero"`
	// Terms of service link. Same scheme rules as privacyPolicyUrl: absolute or
	// relative path only.
	TermsOfServiceURL param.Opt[string] `json:"termsOfServiceUrl,omitzero"`
	paramObj
}

func (r ConsentSettingReplaceParamsRegionRuleTranslationValueConsentModal) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsRegionRuleTranslationValueConsentModal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsRegionRuleTranslationValueConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the preferences modal.
//
// The properties AcceptAllBtn, CloseIconLabel, RejectAllBtn, SavePreferencesBtn,
// Sections, Title are required.
type ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModal struct {
	// Label for the "accept all" button in the preferences modal.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel string `json:"closeIconLabel" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the "save my choices" button.
	SavePreferencesBtn string `json:"savePreferencesBtn" api:"required"`
	// Per-category sections plus a leading general section. The server backfills
	// sections for each defined category if absent.
	Sections []ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModalSection `json:"sections,omitzero" api:"required"`
	// Heading for the preferences modal.
	Title string `json:"title" api:"required"`
	// Pluralizable label for the "N services" counter (e.g. "Service|Services").
	ServiceCounterLabel param.Opt[string] `json:"serviceCounterLabel,omitzero"`
	// Optional collapsible info block rendered after the category sections. Renders
	// only when both title and body are non-empty. Use it for long-form text unrelated
	// to any specific category (e.g. a third-party advertisers disclosure).
	AdditionalInfo ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModalAdditionalInfo `json:"additionalInfo,omitzero"`
	paramObj
}

func (r ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModal) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Description, Title are required.
type ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModalSection struct {
	// Body text for this preferences-modal section.
	Description string `json:"description" api:"required"`
	// Heading for this preferences-modal section.
	Title string `json:"title" api:"required"`
	// Category value this section configures. Use null/omit for the introductory
	// section that applies to no specific category.
	LinkedCategory param.Opt[string] `json:"linkedCategory,omitzero"`
	// Optional disclosure table listing cookies/services in this section.
	CookieTable ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModalSectionCookieTable `json:"cookieTable,omitzero"`
	paramObj
}

func (r ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModalSection) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModalSection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModalSection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional disclosure table listing cookies/services in this section.
//
// The properties Body, Headers are required.
type ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModalSectionCookieTable struct {
	// Rows of the cookie table — each row is a {key,value} pair (e.g. cookie name →
	// purpose).
	Body []ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModalSectionCookieTableBody `json:"body,omitzero" api:"required"`
	// Header cells. Same {key,value} shape as rows.
	Headers []ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader `json:"headers,omitzero" api:"required"`
	// Optional table caption.
	Caption param.Opt[string] `json:"caption,omitzero"`
	paramObj
}

func (r ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModalSectionCookieTable) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModalSectionCookieTable
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModalSectionCookieTable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Key, Value are required.
type ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModalSectionCookieTableBody struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	paramObj
}

func (r ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModalSectionCookieTableBody) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModalSectionCookieTableBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModalSectionCookieTableBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Key, Value are required.
type ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	paramObj
}

func (r ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional collapsible info block rendered after the category sections. Renders
// only when both title and body are non-empty. Use it for long-form text unrelated
// to any specific category (e.g. a third-party advertisers disclosure).
//
// The properties Body, Title are required.
type ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModalAdditionalInfo struct {
	// Body text for the additional-info block. Supports the rich-text markdown subset
	// used by section descriptions: line breaks, [text](url) links (https:// or
	// relative paths), and **bold**. HTML is stripped.
	Body string `json:"body" api:"required"`
	// Heading for the additional-info block.
	Title string `json:"title" api:"required"`
	paramObj
}

func (r ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModalAdditionalInfo) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModalAdditionalInfo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModalAdditionalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Threshold config for autoShowDismissMode (page count or seconds).
type ConsentSettingReplaceParamsRegionRuleAutoShowDismissConfig struct {
	// Pages-viewed threshold when autoShowDismissMode = after_pages.
	PageCount param.Opt[int64] `json:"pageCount,omitzero"`
	// Seconds-on-page threshold when autoShowDismissMode = after_seconds.
	Seconds param.Opt[int64] `json:"seconds,omitzero"`
	paramObj
}

func (r ConsentSettingReplaceParamsRegionRuleAutoShowDismissConfig) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsRegionRuleAutoShowDismissConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsRegionRuleAutoShowDismissConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the modals (layout/position/colors).
type ConsentSettingReplaceParamsRegionRuleGuiOptions struct {
	// Visual options for the initial consent modal.
	ConsentModal ConsentSettingReplaceParamsRegionRuleGuiOptionsConsentModal `json:"consentModal,omitzero"`
	// Global CSS variable overrides applied to both modals.
	CssVariables ConsentSettingReplaceParamsRegionRuleGuiOptionsCssVariables `json:"cssVariables,omitzero"`
	// Visual options for the per-category preferences modal.
	PreferencesModal ConsentSettingReplaceParamsRegionRuleGuiOptionsPreferencesModal `json:"preferencesModal,omitzero"`
	paramObj
}

func (r ConsentSettingReplaceParamsRegionRuleGuiOptions) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsRegionRuleGuiOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsRegionRuleGuiOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the initial consent modal.
type ConsentSettingReplaceParamsRegionRuleGuiOptionsConsentModal struct {
	// When true, accept and reject buttons share the same visual weight (no
	// primary/secondary).
	EqualWeightButtons param.Opt[bool] `json:"equalWeightButtons,omitzero"`
	// When true, swap accept/reject button order.
	FlipButtons param.Opt[bool] `json:"flipButtons,omitzero"`
	// Show an X to dismiss the modal without choosing.
	ShowCloseIcon param.Opt[bool] `json:"showCloseIcon,omitzero"`
	// Which buttons to render in the consent modal and in what arrangement.
	//
	// Any of "AcceptAllNecessaryPreferences", "AcceptAllPreferences", "AcceptOnly",
	// "InformationOnly", "PreferencesOnly".
	ButtonLayout string `json:"buttonLayout,omitzero"`
	// Modal layout style (box, bar, cloud, etc.).
	//
	// Any of "bar", "bar_inline", "box", "box_inline", "box_wide", "cloud",
	// "cloud_inline".
	Layout string `json:"layout,omitzero"`
	// Where the modal anchors on the viewport (e.g. bottom_left, top_right,
	// middle_center).
	//
	// Any of "bottom", "bottom_center", "bottom_left", "bottom_right", "middle",
	// "middle_center", "middle_left", "middle_right", "top", "top_center", "top_left",
	// "top_right".
	Position string `json:"position,omitzero"`
	paramObj
}

func (r ConsentSettingReplaceParamsRegionRuleGuiOptionsConsentModal) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsRegionRuleGuiOptionsConsentModal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsRegionRuleGuiOptionsConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConsentSettingReplaceParamsRegionRuleGuiOptionsConsentModal](
		"buttonLayout", "AcceptAllNecessaryPreferences", "AcceptAllPreferences", "AcceptOnly", "InformationOnly", "PreferencesOnly",
	)
	apijson.RegisterFieldValidator[ConsentSettingReplaceParamsRegionRuleGuiOptionsConsentModal](
		"layout", "bar", "bar_inline", "box", "box_inline", "box_wide", "cloud", "cloud_inline",
	)
	apijson.RegisterFieldValidator[ConsentSettingReplaceParamsRegionRuleGuiOptionsConsentModal](
		"position", "bottom", "bottom_center", "bottom_left", "bottom_right", "middle", "middle_center", "middle_left", "middle_right", "top", "top_center", "top_left", "top_right",
	)
}

// Global CSS variable overrides applied to both modals.
type ConsentSettingReplaceParamsRegionRuleGuiOptionsCssVariables struct {
	// CSS length, e.g. "4px" or "0.25rem".
	ButtonBorderRadius param.Opt[string] `json:"buttonBorderRadius,omitzero"`
	// CSS color for the per-category block background in the preferences modal.
	CookieCategoryBlockBg param.Opt[string] `json:"cookieCategoryBlockBg,omitzero"`
	// CSS color for the per-category block background on hover.
	CookieCategoryBlockHoverBg param.Opt[string] `json:"cookieCategoryBlockHoverBg,omitzero"`
	// CSS color for the modal footer background.
	FooterBg param.Opt[string] `json:"footerBg,omitzero"`
	// CSS color for footer text.
	FooterColor param.Opt[string] `json:"footerColor,omitzero"`
	// CSS color for footer links.
	FooterLinkColor param.Opt[string] `json:"footerLinkColor,omitzero"`
	// CSS color for footer links on hover.
	FooterLinkHoverColor param.Opt[string] `json:"footerLinkHoverColor,omitzero"`
	// CSS color for the modal background.
	ModalBg param.Opt[string] `json:"modalBg,omitzero"`
	// CSS length for modal corner radius.
	ModalBorderRadius param.Opt[string] `json:"modalBorderRadius,omitzero"`
	// CSS color for the primary button background.
	PrimaryButtonBg param.Opt[string] `json:"primaryButtonBg,omitzero"`
	// CSS color for primary button text.
	PrimaryButtonColor param.Opt[string] `json:"primaryButtonColor,omitzero"`
	// CSS color for the primary button background on hover.
	PrimaryButtonHoverBg param.Opt[string] `json:"primaryButtonHoverBg,omitzero"`
	// CSS color for primary button text on hover.
	PrimaryButtonHoverColor param.Opt[string] `json:"primaryButtonHoverColor,omitzero"`
	// CSS color for primary text in the modal body.
	PrimaryTextColor param.Opt[string] `json:"primaryTextColor,omitzero"`
	// CSS color for the secondary button background.
	SecondaryButtonBg param.Opt[string] `json:"secondaryButtonBg,omitzero"`
	// CSS color for secondary button text.
	SecondaryButtonColor param.Opt[string] `json:"secondaryButtonColor,omitzero"`
	// CSS color for secondary button background on hover.
	SecondaryButtonHoverBg param.Opt[string] `json:"secondaryButtonHoverBg,omitzero"`
	// CSS color for secondary button text on hover.
	SecondaryButtonHoverColor param.Opt[string] `json:"secondaryButtonHoverColor,omitzero"`
	// CSS color for secondary text in the modal body.
	SecondaryTextColor param.Opt[string] `json:"secondaryTextColor,omitzero"`
	// CSS color for the "on" state of category toggles.
	ToggleOnBg param.Opt[string] `json:"toggleOnBg,omitzero"`
	paramObj
}

func (r ConsentSettingReplaceParamsRegionRuleGuiOptionsCssVariables) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsRegionRuleGuiOptionsCssVariables
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsRegionRuleGuiOptionsCssVariables) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the per-category preferences modal.
type ConsentSettingReplaceParamsRegionRuleGuiOptionsPreferencesModal struct {
	// When true, accept/reject share the same visual weight.
	EqualWeightButtons param.Opt[bool] `json:"equalWeightButtons,omitzero"`
	// When true, swap accept/reject button order.
	FlipButtons param.Opt[bool] `json:"flipButtons,omitzero"`
	// Button layout for the preferences modal.
	//
	// Any of "AcceptAllRejectAllSave", "AcceptAllSave".
	ButtonLayout string `json:"buttonLayout,omitzero"`
	// Preferences modal layout (box, bar, bar_wide).
	//
	// Any of "bar", "bar_wide", "box".
	Layout string `json:"layout,omitzero"`
	// Preferences modal anchor (left, right).
	//
	// Any of "left", "right".
	Position string `json:"position,omitzero"`
	paramObj
}

func (r ConsentSettingReplaceParamsRegionRuleGuiOptionsPreferencesModal) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsRegionRuleGuiOptionsPreferencesModal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsRegionRuleGuiOptionsPreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConsentSettingReplaceParamsRegionRuleGuiOptionsPreferencesModal](
		"buttonLayout", "AcceptAllRejectAllSave", "AcceptAllSave",
	)
	apijson.RegisterFieldValidator[ConsentSettingReplaceParamsRegionRuleGuiOptionsPreferencesModal](
		"layout", "bar", "bar_wide", "box",
	)
	apijson.RegisterFieldValidator[ConsentSettingReplaceParamsRegionRuleGuiOptionsPreferencesModal](
		"position", "left", "right",
	)
}

// The properties InternalNotes, Label are required.
type ConsentSettingReplaceParamsService struct {
	// Internal notes shown to admins in the dashboard. Not user-facing.
	InternalNotes string `json:"internalNotes" api:"required"`
	// Display name for this service in the preferences modal.
	Label string `json:"label" api:"required"`
	// Primary category value this service belongs to. Must match one of the top-level
	// `categories[].value` entries.
	Category param.Opt[string] `json:"category,omitzero"`
	// Extra category values this service belongs to. Each must match a
	// `categories[].value`.
	AdditionalCategories []string `json:"additionalCategories,omitzero"`
	// Domains/paths this service matches. Patterns matching the CMP's own scripts
	// (e.g. cdn.oursprivacy.com/cmp-init) are rejected to prevent the CMP blocking
	// itself — use a more specific path like cdn.oursprivacy.com/main.js to block a
	// specific script.
	DomainPatterns []string `json:"domainPatterns,omitzero"`
	paramObj
}

func (r ConsentSettingReplaceParamsService) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsService
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enabled requires at least one allowed domain. Disabled may have none.
type ConsentSettingReplaceParamsStatus string

const (
	ConsentSettingReplaceParamsStatusDisabled ConsentSettingReplaceParamsStatus = "Disabled"
	ConsentSettingReplaceParamsStatusEnabled  ConsentSettingReplaceParamsStatus = "Enabled"
)

type ConsentSettingUpdateParams struct {
	// Set or clear the consent cookie name.
	ConsentCookieName param.Opt[string] `json:"consentCookieName,omitzero"`
	// Set or clear the custom CDN domain.
	CustomDomain param.Opt[string] `json:"customDomain,omitzero"`
	// Cookie/localStorage key for the visitor (device) ID. Must exactly match the web
	// SDK's cookie_names.device_id (default "ours_device_id"). Pass null to clear.
	DeviceIDCookieName param.Opt[string] `json:"deviceIdCookieName,omitzero"`
	// Bump the revision counter to re-prompt users.
	Revision param.Opt[float64] `json:"revision,omitzero"`
	// Set or clear the WebSource pixel link. A non-null token must be a valid
	// WebSource of yours.
	WebSDKToken param.Opt[string] `json:"webSDKToken,omitzero"`
	// Rename the consent settings record.
	Name param.Opt[string] `json:"name,omitzero"`
	// Replace the skipBlockingClassNames list. Pass null/[] to clear.
	SkipBlockingClassNames []string `json:"skipBlockingClassNames,omitzero"`
	// Replace the allowed domains when sent. Omit to preserve saved domains. Pass null
	// or [] to clear only when the resulting status is Disabled.
	WhitelistDomains []string `json:"whitelistDomains,omitzero"`
	// Replace the entire categories list. Omit to leave existing categories untouched.
	Categories []ConsentSettingUpdateParamsCategory `json:"categories,omitzero"`
	// Replace the default rule wholesale. Omit to leave it untouched.
	Default ConsentSettingUpdateParamsDefault `json:"default,omitzero"`
	// Replace the entire regions list. Omit to leave it untouched. To change one
	// region, send the full regions array with that region's rule modified.
	Regions []ConsentSettingUpdateParamsRegion `json:"regions,omitzero"`
	// Replace the entire services list. Omit to leave existing services untouched.
	Services []ConsentSettingUpdateParamsService `json:"services,omitzero"`
	// Set Enabled or Disabled. Enabling requires at least one allowed domain, either
	// already saved or sent with this request.
	//
	// Any of "Disabled", "Enabled".
	Status ConsentSettingUpdateParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Label, Priority, Value are required.
type ConsentSettingUpdateParamsCategory struct {
	// Human-readable label shown next to the toggle in the preferences modal.
	Label string `json:"label" api:"required"`
	// Sort key. Lower numbers render first. Server re-stamps to 0..N on write — send
	// any integer, gaps and duplicates are ironed out.
	Priority int64 `json:"priority" api:"required"`
	// Stable identifier referenced by services and translation sections.
	// Conventionally lowercase (e.g. "necessary", "analytics", "advertising").
	Value string `json:"value" api:"required"`
	paramObj
}

func (r ConsentSettingUpdateParamsCategory) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsCategory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Replace the default rule wholesale. Omit to leave it untouched.
//
// The properties Categories, Language, Mode, Translations are required.
type ConsentSettingUpdateParamsDefault struct {
	// Per-category default config for this rule. Every category defined in the
	// top-level `categories[].value` should have an entry here.
	Categories []ConsentSettingUpdateParamsDefaultCategory `json:"categories,omitzero" api:"required"`
	// BCP 47 default language for this rule. Must have a matching entry in
	// `translations`. Examples: "en", "en-US", "es", "de".
	Language string `json:"language" api:"required"`
	// opt_in: scripts blocked until user accepts (GDPR style). opt_out: scripts run by
	// default until user rejects (CCPA style).
	//
	// Any of "opt_in", "opt_out".
	Mode string `json:"mode,omitzero" api:"required"`
	// All UI copy, keyed by language. Must include an entry whose `language` matches
	// the rule's `language` field.
	Translations []ConsentSettingUpdateParamsDefaultTranslation `json:"translations,omitzero" api:"required"`
	// When true, scripts not classified by services[] are blocked until the user opts
	// in.
	AutoblockUnknown param.Opt[bool] `json:"autoblockUnknown,omitzero"`
	// When true, the consent modal auto-opens on page load.
	AutoShow param.Opt[bool] `json:"autoShow,omitzero"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction param.Opt[bool] `json:"disablePageInteraction,omitzero"`
	// When true, the modal is suppressed for known bot user agents.
	HideFromBots param.Opt[bool] `json:"hideFromBots,omitzero"`
	// When true, the per-service list (services[]) is rendered inside the preferences
	// modal.
	ShowVendorsInPreferences param.Opt[bool] `json:"showVendorsInPreferences,omitzero"`
	// Threshold config for autoShowDismissMode (page count or seconds).
	AutoShowDismissConfig ConsentSettingUpdateParamsDefaultAutoShowDismissConfig `json:"autoShowDismissConfig,omitzero"`
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	//
	// Any of "after_pages", "after_seconds", "never".
	AutoShowDismissMode string `json:"autoShowDismissMode,omitzero"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions ConsentSettingUpdateParamsDefaultGuiOptions `json:"guiOptions,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsDefault) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsDefault
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsDefault) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConsentSettingUpdateParamsDefault](
		"mode", "opt_in", "opt_out",
	)
	apijson.RegisterFieldValidator[ConsentSettingUpdateParamsDefault](
		"autoShowDismissMode", "after_pages", "after_seconds", "never",
	)
}

// The properties Key, Value are required.
type ConsentSettingUpdateParamsDefaultCategory struct {
	// Category value (matches `categories[].value`) this entry configures.
	Key   string                                         `json:"key" api:"required"`
	Value ConsentSettingUpdateParamsDefaultCategoryValue `json:"value,omitzero" api:"required"`
	paramObj
}

func (r ConsentSettingUpdateParamsDefaultCategory) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsDefaultCategory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsDefaultCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Enabled is required.
type ConsentSettingUpdateParamsDefaultCategoryValue struct {
	// Whether this category is on by default before the user interacts.
	Enabled bool `json:"enabled" api:"required"`
	// When true, this category defaults off if the browser sends Sec-GPC: 1.
	AutoDisableOnGpc param.Opt[bool] `json:"autoDisableOnGPC,omitzero"`
	// When true, the user cannot toggle this category in the preferences modal.
	ReadOnly param.Opt[bool] `json:"readOnly,omitzero"`
	// When true, the page reloads after this category is toggled so newly-allowed
	// scripts can run.
	ReloadPage param.Opt[bool] `json:"reloadPage,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsDefaultCategoryValue) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsDefaultCategoryValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsDefaultCategoryValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Language, Value are required.
type ConsentSettingUpdateParamsDefaultTranslation struct {
	// BCP 47 language tag identifying which translation this entry provides. Examples:
	// "en", "en-US", "es", "fr-CA". The default rule's `language` must appear here.
	Language string                                            `json:"language" api:"required"`
	Value    ConsentSettingUpdateParamsDefaultTranslationValue `json:"value,omitzero" api:"required"`
	paramObj
}

func (r ConsentSettingUpdateParamsDefaultTranslation) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsDefaultTranslation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsDefaultTranslation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingUpdateParamsDefaultTranslationValue struct {
	// Translated copy for the initial consent modal.
	ConsentModal ConsentSettingUpdateParamsDefaultTranslationValueConsentModal `json:"consentModal,omitzero"`
	// Translated copy for the preferences modal.
	PreferencesModal ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModal `json:"preferencesModal,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsDefaultTranslationValue) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsDefaultTranslationValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsDefaultTranslationValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the initial consent modal.
//
// The properties AcceptAllBtn, Description, RejectAllBtn, ShowPreferencesBtn,
// Title are required.
type ConsentSettingUpdateParamsDefaultTranslationValueConsentModal struct {
	// Label for the "accept all" button.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Body text shown in the consent modal.
	Description string `json:"description" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the button that opens the preferences modal.
	ShowPreferencesBtn string `json:"showPreferencesBtn" api:"required"`
	// Heading shown at the top of the consent modal.
	Title string `json:"title" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel param.Opt[string] `json:"closeIconLabel,omitzero"`
	// Optional footer HTML/text shown below the body.
	Footer param.Opt[string] `json:"footer,omitzero"`
	// Optional text shown when the browser is sending Sec-GPC: 1.
	GpcNotification param.Opt[string] `json:"gpcNotification,omitzero"`
	// Link label for the privacy policy (defaults to "Privacy Policy").
	PrivacyPolicyLabel param.Opt[string] `json:"privacyPolicyLabel,omitzero"`
	// Privacy policy link. Accepts an absolute URL (https://...) or a relative path
	// (/legal/privacy) so one config can run across multiple domains. javascript:,
	// data:, and vbscript: schemes are rejected.
	PrivacyPolicyURL param.Opt[string] `json:"privacyPolicyUrl,omitzero"`
	// Optional message shown when the consent revision has bumped (re-prompt copy).
	RevisionMessage param.Opt[string] `json:"revisionMessage,omitzero"`
	// Link label for terms of service (defaults to "Terms of Service").
	TermsOfServiceLabel param.Opt[string] `json:"termsOfServiceLabel,omitzero"`
	// Terms of service link. Same scheme rules as privacyPolicyUrl: absolute or
	// relative path only.
	TermsOfServiceURL param.Opt[string] `json:"termsOfServiceUrl,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsDefaultTranslationValueConsentModal) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsDefaultTranslationValueConsentModal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsDefaultTranslationValueConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the preferences modal.
//
// The properties AcceptAllBtn, CloseIconLabel, RejectAllBtn, SavePreferencesBtn,
// Sections, Title are required.
type ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModal struct {
	// Label for the "accept all" button in the preferences modal.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel string `json:"closeIconLabel" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the "save my choices" button.
	SavePreferencesBtn string `json:"savePreferencesBtn" api:"required"`
	// Per-category sections plus a leading general section. The server backfills
	// sections for each defined category if absent.
	Sections []ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSection `json:"sections,omitzero" api:"required"`
	// Heading for the preferences modal.
	Title string `json:"title" api:"required"`
	// Pluralizable label for the "N services" counter (e.g. "Service|Services").
	ServiceCounterLabel param.Opt[string] `json:"serviceCounterLabel,omitzero"`
	// Optional collapsible info block rendered after the category sections. Renders
	// only when both title and body are non-empty. Use it for long-form text unrelated
	// to any specific category (e.g. a third-party advertisers disclosure).
	AdditionalInfo ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalAdditionalInfo `json:"additionalInfo,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModal) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Description, Title are required.
type ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSection struct {
	// Body text for this preferences-modal section.
	Description string `json:"description" api:"required"`
	// Heading for this preferences-modal section.
	Title string `json:"title" api:"required"`
	// Category value this section configures. Use null/omit for the introductory
	// section that applies to no specific category.
	LinkedCategory param.Opt[string] `json:"linkedCategory,omitzero"`
	// Optional disclosure table listing cookies/services in this section.
	CookieTable ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSectionCookieTable `json:"cookieTable,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSection) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional disclosure table listing cookies/services in this section.
//
// The properties Body, Headers are required.
type ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSectionCookieTable struct {
	// Rows of the cookie table — each row is a {key,value} pair (e.g. cookie name →
	// purpose).
	Body []ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSectionCookieTableBody `json:"body,omitzero" api:"required"`
	// Header cells. Same {key,value} shape as rows.
	Headers []ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSectionCookieTableHeader `json:"headers,omitzero" api:"required"`
	// Optional table caption.
	Caption param.Opt[string] `json:"caption,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSectionCookieTable) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSectionCookieTable
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSectionCookieTable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Key, Value are required.
type ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSectionCookieTableBody struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	paramObj
}

func (r ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSectionCookieTableBody) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSectionCookieTableBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSectionCookieTableBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Key, Value are required.
type ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSectionCookieTableHeader struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	paramObj
}

func (r ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSectionCookieTableHeader) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSectionCookieTableHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSectionCookieTableHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional collapsible info block rendered after the category sections. Renders
// only when both title and body are non-empty. Use it for long-form text unrelated
// to any specific category (e.g. a third-party advertisers disclosure).
//
// The properties Body, Title are required.
type ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalAdditionalInfo struct {
	// Body text for the additional-info block. Supports the rich-text markdown subset
	// used by section descriptions: line breaks, [text](url) links (https:// or
	// relative paths), and **bold**. HTML is stripped.
	Body string `json:"body" api:"required"`
	// Heading for the additional-info block.
	Title string `json:"title" api:"required"`
	paramObj
}

func (r ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalAdditionalInfo) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalAdditionalInfo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalAdditionalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Threshold config for autoShowDismissMode (page count or seconds).
type ConsentSettingUpdateParamsDefaultAutoShowDismissConfig struct {
	// Pages-viewed threshold when autoShowDismissMode = after_pages.
	PageCount param.Opt[int64] `json:"pageCount,omitzero"`
	// Seconds-on-page threshold when autoShowDismissMode = after_seconds.
	Seconds param.Opt[int64] `json:"seconds,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsDefaultAutoShowDismissConfig) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsDefaultAutoShowDismissConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsDefaultAutoShowDismissConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the modals (layout/position/colors).
type ConsentSettingUpdateParamsDefaultGuiOptions struct {
	// Visual options for the initial consent modal.
	ConsentModal ConsentSettingUpdateParamsDefaultGuiOptionsConsentModal `json:"consentModal,omitzero"`
	// Global CSS variable overrides applied to both modals.
	CssVariables ConsentSettingUpdateParamsDefaultGuiOptionsCssVariables `json:"cssVariables,omitzero"`
	// Visual options for the per-category preferences modal.
	PreferencesModal ConsentSettingUpdateParamsDefaultGuiOptionsPreferencesModal `json:"preferencesModal,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsDefaultGuiOptions) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsDefaultGuiOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsDefaultGuiOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the initial consent modal.
type ConsentSettingUpdateParamsDefaultGuiOptionsConsentModal struct {
	// When true, accept and reject buttons share the same visual weight (no
	// primary/secondary).
	EqualWeightButtons param.Opt[bool] `json:"equalWeightButtons,omitzero"`
	// When true, swap accept/reject button order.
	FlipButtons param.Opt[bool] `json:"flipButtons,omitzero"`
	// Show an X to dismiss the modal without choosing.
	ShowCloseIcon param.Opt[bool] `json:"showCloseIcon,omitzero"`
	// Which buttons to render in the consent modal and in what arrangement.
	//
	// Any of "AcceptAllNecessaryPreferences", "AcceptAllPreferences", "AcceptOnly",
	// "InformationOnly", "PreferencesOnly".
	ButtonLayout string `json:"buttonLayout,omitzero"`
	// Modal layout style (box, bar, cloud, etc.).
	//
	// Any of "bar", "bar_inline", "box", "box_inline", "box_wide", "cloud",
	// "cloud_inline".
	Layout string `json:"layout,omitzero"`
	// Where the modal anchors on the viewport (e.g. bottom_left, top_right,
	// middle_center).
	//
	// Any of "bottom", "bottom_center", "bottom_left", "bottom_right", "middle",
	// "middle_center", "middle_left", "middle_right", "top", "top_center", "top_left",
	// "top_right".
	Position string `json:"position,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsDefaultGuiOptionsConsentModal) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsDefaultGuiOptionsConsentModal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsDefaultGuiOptionsConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConsentSettingUpdateParamsDefaultGuiOptionsConsentModal](
		"buttonLayout", "AcceptAllNecessaryPreferences", "AcceptAllPreferences", "AcceptOnly", "InformationOnly", "PreferencesOnly",
	)
	apijson.RegisterFieldValidator[ConsentSettingUpdateParamsDefaultGuiOptionsConsentModal](
		"layout", "bar", "bar_inline", "box", "box_inline", "box_wide", "cloud", "cloud_inline",
	)
	apijson.RegisterFieldValidator[ConsentSettingUpdateParamsDefaultGuiOptionsConsentModal](
		"position", "bottom", "bottom_center", "bottom_left", "bottom_right", "middle", "middle_center", "middle_left", "middle_right", "top", "top_center", "top_left", "top_right",
	)
}

// Global CSS variable overrides applied to both modals.
type ConsentSettingUpdateParamsDefaultGuiOptionsCssVariables struct {
	// CSS length, e.g. "4px" or "0.25rem".
	ButtonBorderRadius param.Opt[string] `json:"buttonBorderRadius,omitzero"`
	// CSS color for the per-category block background in the preferences modal.
	CookieCategoryBlockBg param.Opt[string] `json:"cookieCategoryBlockBg,omitzero"`
	// CSS color for the per-category block background on hover.
	CookieCategoryBlockHoverBg param.Opt[string] `json:"cookieCategoryBlockHoverBg,omitzero"`
	// CSS color for the modal footer background.
	FooterBg param.Opt[string] `json:"footerBg,omitzero"`
	// CSS color for footer text.
	FooterColor param.Opt[string] `json:"footerColor,omitzero"`
	// CSS color for footer links.
	FooterLinkColor param.Opt[string] `json:"footerLinkColor,omitzero"`
	// CSS color for footer links on hover.
	FooterLinkHoverColor param.Opt[string] `json:"footerLinkHoverColor,omitzero"`
	// CSS color for the modal background.
	ModalBg param.Opt[string] `json:"modalBg,omitzero"`
	// CSS length for modal corner radius.
	ModalBorderRadius param.Opt[string] `json:"modalBorderRadius,omitzero"`
	// CSS color for the primary button background.
	PrimaryButtonBg param.Opt[string] `json:"primaryButtonBg,omitzero"`
	// CSS color for primary button text.
	PrimaryButtonColor param.Opt[string] `json:"primaryButtonColor,omitzero"`
	// CSS color for the primary button background on hover.
	PrimaryButtonHoverBg param.Opt[string] `json:"primaryButtonHoverBg,omitzero"`
	// CSS color for primary button text on hover.
	PrimaryButtonHoverColor param.Opt[string] `json:"primaryButtonHoverColor,omitzero"`
	// CSS color for primary text in the modal body.
	PrimaryTextColor param.Opt[string] `json:"primaryTextColor,omitzero"`
	// CSS color for the secondary button background.
	SecondaryButtonBg param.Opt[string] `json:"secondaryButtonBg,omitzero"`
	// CSS color for secondary button text.
	SecondaryButtonColor param.Opt[string] `json:"secondaryButtonColor,omitzero"`
	// CSS color for secondary button background on hover.
	SecondaryButtonHoverBg param.Opt[string] `json:"secondaryButtonHoverBg,omitzero"`
	// CSS color for secondary button text on hover.
	SecondaryButtonHoverColor param.Opt[string] `json:"secondaryButtonHoverColor,omitzero"`
	// CSS color for secondary text in the modal body.
	SecondaryTextColor param.Opt[string] `json:"secondaryTextColor,omitzero"`
	// CSS color for the "on" state of category toggles.
	ToggleOnBg param.Opt[string] `json:"toggleOnBg,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsDefaultGuiOptionsCssVariables) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsDefaultGuiOptionsCssVariables
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsDefaultGuiOptionsCssVariables) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the per-category preferences modal.
type ConsentSettingUpdateParamsDefaultGuiOptionsPreferencesModal struct {
	// When true, accept/reject share the same visual weight.
	EqualWeightButtons param.Opt[bool] `json:"equalWeightButtons,omitzero"`
	// When true, swap accept/reject button order.
	FlipButtons param.Opt[bool] `json:"flipButtons,omitzero"`
	// Button layout for the preferences modal.
	//
	// Any of "AcceptAllRejectAllSave", "AcceptAllSave".
	ButtonLayout string `json:"buttonLayout,omitzero"`
	// Preferences modal layout (box, bar, bar_wide).
	//
	// Any of "bar", "bar_wide", "box".
	Layout string `json:"layout,omitzero"`
	// Preferences modal anchor (left, right).
	//
	// Any of "left", "right".
	Position string `json:"position,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsDefaultGuiOptionsPreferencesModal) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsDefaultGuiOptionsPreferencesModal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsDefaultGuiOptionsPreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConsentSettingUpdateParamsDefaultGuiOptionsPreferencesModal](
		"buttonLayout", "AcceptAllRejectAllSave", "AcceptAllSave",
	)
	apijson.RegisterFieldValidator[ConsentSettingUpdateParamsDefaultGuiOptionsPreferencesModal](
		"layout", "bar", "bar_wide", "box",
	)
	apijson.RegisterFieldValidator[ConsentSettingUpdateParamsDefaultGuiOptionsPreferencesModal](
		"position", "left", "right",
	)
}

// The properties RegionCode, Rule are required.
type ConsentSettingUpdateParamsRegion struct {
	// Region this rule applies to. Use ISO 3166-1 alpha-2 country code ("US", "DE",
	// "BR") or country-subdivision code ("US-CA", "GB-ENG", "CA-ON"). Each region code
	// may appear in only one rule across `regions[]`.
	RegionCode string                               `json:"regionCode" api:"required"`
	Rule       ConsentSettingUpdateParamsRegionRule `json:"rule,omitzero" api:"required"`
	// Optional label for this override group, e.g. "CCPA States". Shown in the
	// dashboard instead of the region codes. Pass null to clear it and fall back to
	// the codes.
	Name param.Opt[string] `json:"name,omitzero"`
	// Other region codes that should reuse this rule. Same code-format rules as
	// `regionCode`. Cannot include `regionCode` itself, cannot duplicate, cannot
	// overlap with another rule's regions.
	AdditionalRegions []string `json:"additionalRegions,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsRegion) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsRegion
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Categories, Language, Mode, Translations are required.
type ConsentSettingUpdateParamsRegionRule struct {
	// Per-category default config for this rule. Every category defined in the
	// top-level `categories[].value` should have an entry here.
	Categories []ConsentSettingUpdateParamsRegionRuleCategory `json:"categories,omitzero" api:"required"`
	// BCP 47 default language for this rule. Must have a matching entry in
	// `translations`. Examples: "en", "en-US", "es", "de".
	Language string `json:"language" api:"required"`
	// opt_in: scripts blocked until user accepts (GDPR style). opt_out: scripts run by
	// default until user rejects (CCPA style).
	//
	// Any of "opt_in", "opt_out".
	Mode string `json:"mode,omitzero" api:"required"`
	// All UI copy, keyed by language. Must include an entry whose `language` matches
	// the rule's `language` field.
	Translations []ConsentSettingUpdateParamsRegionRuleTranslation `json:"translations,omitzero" api:"required"`
	// When true, scripts not classified by services[] are blocked until the user opts
	// in.
	AutoblockUnknown param.Opt[bool] `json:"autoblockUnknown,omitzero"`
	// When true, the consent modal auto-opens on page load.
	AutoShow param.Opt[bool] `json:"autoShow,omitzero"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction param.Opt[bool] `json:"disablePageInteraction,omitzero"`
	// When true, the modal is suppressed for known bot user agents.
	HideFromBots param.Opt[bool] `json:"hideFromBots,omitzero"`
	// When true, the per-service list (services[]) is rendered inside the preferences
	// modal.
	ShowVendorsInPreferences param.Opt[bool] `json:"showVendorsInPreferences,omitzero"`
	// Threshold config for autoShowDismissMode (page count or seconds).
	AutoShowDismissConfig ConsentSettingUpdateParamsRegionRuleAutoShowDismissConfig `json:"autoShowDismissConfig,omitzero"`
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	//
	// Any of "after_pages", "after_seconds", "never".
	AutoShowDismissMode string `json:"autoShowDismissMode,omitzero"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions ConsentSettingUpdateParamsRegionRuleGuiOptions `json:"guiOptions,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsRegionRule) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsRegionRule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsRegionRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConsentSettingUpdateParamsRegionRule](
		"mode", "opt_in", "opt_out",
	)
	apijson.RegisterFieldValidator[ConsentSettingUpdateParamsRegionRule](
		"autoShowDismissMode", "after_pages", "after_seconds", "never",
	)
}

// The properties Key, Value are required.
type ConsentSettingUpdateParamsRegionRuleCategory struct {
	// Category value (matches `categories[].value`) this entry configures.
	Key   string                                            `json:"key" api:"required"`
	Value ConsentSettingUpdateParamsRegionRuleCategoryValue `json:"value,omitzero" api:"required"`
	paramObj
}

func (r ConsentSettingUpdateParamsRegionRuleCategory) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsRegionRuleCategory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsRegionRuleCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Enabled is required.
type ConsentSettingUpdateParamsRegionRuleCategoryValue struct {
	// Whether this category is on by default before the user interacts.
	Enabled bool `json:"enabled" api:"required"`
	// When true, this category defaults off if the browser sends Sec-GPC: 1.
	AutoDisableOnGpc param.Opt[bool] `json:"autoDisableOnGPC,omitzero"`
	// When true, the user cannot toggle this category in the preferences modal.
	ReadOnly param.Opt[bool] `json:"readOnly,omitzero"`
	// When true, the page reloads after this category is toggled so newly-allowed
	// scripts can run.
	ReloadPage param.Opt[bool] `json:"reloadPage,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsRegionRuleCategoryValue) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsRegionRuleCategoryValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsRegionRuleCategoryValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Language, Value are required.
type ConsentSettingUpdateParamsRegionRuleTranslation struct {
	// BCP 47 language tag identifying which translation this entry provides. Examples:
	// "en", "en-US", "es", "fr-CA". The default rule's `language` must appear here.
	Language string                                               `json:"language" api:"required"`
	Value    ConsentSettingUpdateParamsRegionRuleTranslationValue `json:"value,omitzero" api:"required"`
	paramObj
}

func (r ConsentSettingUpdateParamsRegionRuleTranslation) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsRegionRuleTranslation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsRegionRuleTranslation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingUpdateParamsRegionRuleTranslationValue struct {
	// Translated copy for the initial consent modal.
	ConsentModal ConsentSettingUpdateParamsRegionRuleTranslationValueConsentModal `json:"consentModal,omitzero"`
	// Translated copy for the preferences modal.
	PreferencesModal ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModal `json:"preferencesModal,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsRegionRuleTranslationValue) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsRegionRuleTranslationValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsRegionRuleTranslationValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the initial consent modal.
//
// The properties AcceptAllBtn, Description, RejectAllBtn, ShowPreferencesBtn,
// Title are required.
type ConsentSettingUpdateParamsRegionRuleTranslationValueConsentModal struct {
	// Label for the "accept all" button.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Body text shown in the consent modal.
	Description string `json:"description" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the button that opens the preferences modal.
	ShowPreferencesBtn string `json:"showPreferencesBtn" api:"required"`
	// Heading shown at the top of the consent modal.
	Title string `json:"title" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel param.Opt[string] `json:"closeIconLabel,omitzero"`
	// Optional footer HTML/text shown below the body.
	Footer param.Opt[string] `json:"footer,omitzero"`
	// Optional text shown when the browser is sending Sec-GPC: 1.
	GpcNotification param.Opt[string] `json:"gpcNotification,omitzero"`
	// Link label for the privacy policy (defaults to "Privacy Policy").
	PrivacyPolicyLabel param.Opt[string] `json:"privacyPolicyLabel,omitzero"`
	// Privacy policy link. Accepts an absolute URL (https://...) or a relative path
	// (/legal/privacy) so one config can run across multiple domains. javascript:,
	// data:, and vbscript: schemes are rejected.
	PrivacyPolicyURL param.Opt[string] `json:"privacyPolicyUrl,omitzero"`
	// Optional message shown when the consent revision has bumped (re-prompt copy).
	RevisionMessage param.Opt[string] `json:"revisionMessage,omitzero"`
	// Link label for terms of service (defaults to "Terms of Service").
	TermsOfServiceLabel param.Opt[string] `json:"termsOfServiceLabel,omitzero"`
	// Terms of service link. Same scheme rules as privacyPolicyUrl: absolute or
	// relative path only.
	TermsOfServiceURL param.Opt[string] `json:"termsOfServiceUrl,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsRegionRuleTranslationValueConsentModal) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsRegionRuleTranslationValueConsentModal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsRegionRuleTranslationValueConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Translated copy for the preferences modal.
//
// The properties AcceptAllBtn, CloseIconLabel, RejectAllBtn, SavePreferencesBtn,
// Sections, Title are required.
type ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModal struct {
	// Label for the "accept all" button in the preferences modal.
	AcceptAllBtn string `json:"acceptAllBtn" api:"required"`
	// Accessibility label for the close (X) icon.
	CloseIconLabel string `json:"closeIconLabel" api:"required"`
	// Label for the "reject all" / "necessary only" button.
	RejectAllBtn string `json:"rejectAllBtn" api:"required"`
	// Label for the "save my choices" button.
	SavePreferencesBtn string `json:"savePreferencesBtn" api:"required"`
	// Per-category sections plus a leading general section. The server backfills
	// sections for each defined category if absent.
	Sections []ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSection `json:"sections,omitzero" api:"required"`
	// Heading for the preferences modal.
	Title string `json:"title" api:"required"`
	// Pluralizable label for the "N services" counter (e.g. "Service|Services").
	ServiceCounterLabel param.Opt[string] `json:"serviceCounterLabel,omitzero"`
	// Optional collapsible info block rendered after the category sections. Renders
	// only when both title and body are non-empty. Use it for long-form text unrelated
	// to any specific category (e.g. a third-party advertisers disclosure).
	AdditionalInfo ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalAdditionalInfo `json:"additionalInfo,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModal) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Description, Title are required.
type ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSection struct {
	// Body text for this preferences-modal section.
	Description string `json:"description" api:"required"`
	// Heading for this preferences-modal section.
	Title string `json:"title" api:"required"`
	// Category value this section configures. Use null/omit for the introductory
	// section that applies to no specific category.
	LinkedCategory param.Opt[string] `json:"linkedCategory,omitzero"`
	// Optional disclosure table listing cookies/services in this section.
	CookieTable ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSectionCookieTable `json:"cookieTable,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSection) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional disclosure table listing cookies/services in this section.
//
// The properties Body, Headers are required.
type ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSectionCookieTable struct {
	// Rows of the cookie table — each row is a {key,value} pair (e.g. cookie name →
	// purpose).
	Body []ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSectionCookieTableBody `json:"body,omitzero" api:"required"`
	// Header cells. Same {key,value} shape as rows.
	Headers []ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader `json:"headers,omitzero" api:"required"`
	// Optional table caption.
	Caption param.Opt[string] `json:"caption,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSectionCookieTable) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSectionCookieTable
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSectionCookieTable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Key, Value are required.
type ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSectionCookieTableBody struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	paramObj
}

func (r ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSectionCookieTableBody) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSectionCookieTableBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSectionCookieTableBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Key, Value are required.
type ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	paramObj
}

func (r ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional collapsible info block rendered after the category sections. Renders
// only when both title and body are non-empty. Use it for long-form text unrelated
// to any specific category (e.g. a third-party advertisers disclosure).
//
// The properties Body, Title are required.
type ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalAdditionalInfo struct {
	// Body text for the additional-info block. Supports the rich-text markdown subset
	// used by section descriptions: line breaks, [text](url) links (https:// or
	// relative paths), and **bold**. HTML is stripped.
	Body string `json:"body" api:"required"`
	// Heading for the additional-info block.
	Title string `json:"title" api:"required"`
	paramObj
}

func (r ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalAdditionalInfo) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalAdditionalInfo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalAdditionalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Threshold config for autoShowDismissMode (page count or seconds).
type ConsentSettingUpdateParamsRegionRuleAutoShowDismissConfig struct {
	// Pages-viewed threshold when autoShowDismissMode = after_pages.
	PageCount param.Opt[int64] `json:"pageCount,omitzero"`
	// Seconds-on-page threshold when autoShowDismissMode = after_seconds.
	Seconds param.Opt[int64] `json:"seconds,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsRegionRuleAutoShowDismissConfig) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsRegionRuleAutoShowDismissConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsRegionRuleAutoShowDismissConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the modals (layout/position/colors).
type ConsentSettingUpdateParamsRegionRuleGuiOptions struct {
	// Visual options for the initial consent modal.
	ConsentModal ConsentSettingUpdateParamsRegionRuleGuiOptionsConsentModal `json:"consentModal,omitzero"`
	// Global CSS variable overrides applied to both modals.
	CssVariables ConsentSettingUpdateParamsRegionRuleGuiOptionsCssVariables `json:"cssVariables,omitzero"`
	// Visual options for the per-category preferences modal.
	PreferencesModal ConsentSettingUpdateParamsRegionRuleGuiOptionsPreferencesModal `json:"preferencesModal,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsRegionRuleGuiOptions) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsRegionRuleGuiOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsRegionRuleGuiOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the initial consent modal.
type ConsentSettingUpdateParamsRegionRuleGuiOptionsConsentModal struct {
	// When true, accept and reject buttons share the same visual weight (no
	// primary/secondary).
	EqualWeightButtons param.Opt[bool] `json:"equalWeightButtons,omitzero"`
	// When true, swap accept/reject button order.
	FlipButtons param.Opt[bool] `json:"flipButtons,omitzero"`
	// Show an X to dismiss the modal without choosing.
	ShowCloseIcon param.Opt[bool] `json:"showCloseIcon,omitzero"`
	// Which buttons to render in the consent modal and in what arrangement.
	//
	// Any of "AcceptAllNecessaryPreferences", "AcceptAllPreferences", "AcceptOnly",
	// "InformationOnly", "PreferencesOnly".
	ButtonLayout string `json:"buttonLayout,omitzero"`
	// Modal layout style (box, bar, cloud, etc.).
	//
	// Any of "bar", "bar_inline", "box", "box_inline", "box_wide", "cloud",
	// "cloud_inline".
	Layout string `json:"layout,omitzero"`
	// Where the modal anchors on the viewport (e.g. bottom_left, top_right,
	// middle_center).
	//
	// Any of "bottom", "bottom_center", "bottom_left", "bottom_right", "middle",
	// "middle_center", "middle_left", "middle_right", "top", "top_center", "top_left",
	// "top_right".
	Position string `json:"position,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsRegionRuleGuiOptionsConsentModal) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsRegionRuleGuiOptionsConsentModal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsRegionRuleGuiOptionsConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConsentSettingUpdateParamsRegionRuleGuiOptionsConsentModal](
		"buttonLayout", "AcceptAllNecessaryPreferences", "AcceptAllPreferences", "AcceptOnly", "InformationOnly", "PreferencesOnly",
	)
	apijson.RegisterFieldValidator[ConsentSettingUpdateParamsRegionRuleGuiOptionsConsentModal](
		"layout", "bar", "bar_inline", "box", "box_inline", "box_wide", "cloud", "cloud_inline",
	)
	apijson.RegisterFieldValidator[ConsentSettingUpdateParamsRegionRuleGuiOptionsConsentModal](
		"position", "bottom", "bottom_center", "bottom_left", "bottom_right", "middle", "middle_center", "middle_left", "middle_right", "top", "top_center", "top_left", "top_right",
	)
}

// Global CSS variable overrides applied to both modals.
type ConsentSettingUpdateParamsRegionRuleGuiOptionsCssVariables struct {
	// CSS length, e.g. "4px" or "0.25rem".
	ButtonBorderRadius param.Opt[string] `json:"buttonBorderRadius,omitzero"`
	// CSS color for the per-category block background in the preferences modal.
	CookieCategoryBlockBg param.Opt[string] `json:"cookieCategoryBlockBg,omitzero"`
	// CSS color for the per-category block background on hover.
	CookieCategoryBlockHoverBg param.Opt[string] `json:"cookieCategoryBlockHoverBg,omitzero"`
	// CSS color for the modal footer background.
	FooterBg param.Opt[string] `json:"footerBg,omitzero"`
	// CSS color for footer text.
	FooterColor param.Opt[string] `json:"footerColor,omitzero"`
	// CSS color for footer links.
	FooterLinkColor param.Opt[string] `json:"footerLinkColor,omitzero"`
	// CSS color for footer links on hover.
	FooterLinkHoverColor param.Opt[string] `json:"footerLinkHoverColor,omitzero"`
	// CSS color for the modal background.
	ModalBg param.Opt[string] `json:"modalBg,omitzero"`
	// CSS length for modal corner radius.
	ModalBorderRadius param.Opt[string] `json:"modalBorderRadius,omitzero"`
	// CSS color for the primary button background.
	PrimaryButtonBg param.Opt[string] `json:"primaryButtonBg,omitzero"`
	// CSS color for primary button text.
	PrimaryButtonColor param.Opt[string] `json:"primaryButtonColor,omitzero"`
	// CSS color for the primary button background on hover.
	PrimaryButtonHoverBg param.Opt[string] `json:"primaryButtonHoverBg,omitzero"`
	// CSS color for primary button text on hover.
	PrimaryButtonHoverColor param.Opt[string] `json:"primaryButtonHoverColor,omitzero"`
	// CSS color for primary text in the modal body.
	PrimaryTextColor param.Opt[string] `json:"primaryTextColor,omitzero"`
	// CSS color for the secondary button background.
	SecondaryButtonBg param.Opt[string] `json:"secondaryButtonBg,omitzero"`
	// CSS color for secondary button text.
	SecondaryButtonColor param.Opt[string] `json:"secondaryButtonColor,omitzero"`
	// CSS color for secondary button background on hover.
	SecondaryButtonHoverBg param.Opt[string] `json:"secondaryButtonHoverBg,omitzero"`
	// CSS color for secondary button text on hover.
	SecondaryButtonHoverColor param.Opt[string] `json:"secondaryButtonHoverColor,omitzero"`
	// CSS color for secondary text in the modal body.
	SecondaryTextColor param.Opt[string] `json:"secondaryTextColor,omitzero"`
	// CSS color for the "on" state of category toggles.
	ToggleOnBg param.Opt[string] `json:"toggleOnBg,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsRegionRuleGuiOptionsCssVariables) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsRegionRuleGuiOptionsCssVariables
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsRegionRuleGuiOptionsCssVariables) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Visual options for the per-category preferences modal.
type ConsentSettingUpdateParamsRegionRuleGuiOptionsPreferencesModal struct {
	// When true, accept/reject share the same visual weight.
	EqualWeightButtons param.Opt[bool] `json:"equalWeightButtons,omitzero"`
	// When true, swap accept/reject button order.
	FlipButtons param.Opt[bool] `json:"flipButtons,omitzero"`
	// Button layout for the preferences modal.
	//
	// Any of "AcceptAllRejectAllSave", "AcceptAllSave".
	ButtonLayout string `json:"buttonLayout,omitzero"`
	// Preferences modal layout (box, bar, bar_wide).
	//
	// Any of "bar", "bar_wide", "box".
	Layout string `json:"layout,omitzero"`
	// Preferences modal anchor (left, right).
	//
	// Any of "left", "right".
	Position string `json:"position,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsRegionRuleGuiOptionsPreferencesModal) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsRegionRuleGuiOptionsPreferencesModal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsRegionRuleGuiOptionsPreferencesModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConsentSettingUpdateParamsRegionRuleGuiOptionsPreferencesModal](
		"buttonLayout", "AcceptAllRejectAllSave", "AcceptAllSave",
	)
	apijson.RegisterFieldValidator[ConsentSettingUpdateParamsRegionRuleGuiOptionsPreferencesModal](
		"layout", "bar", "bar_wide", "box",
	)
	apijson.RegisterFieldValidator[ConsentSettingUpdateParamsRegionRuleGuiOptionsPreferencesModal](
		"position", "left", "right",
	)
}

// The properties InternalNotes, Label are required.
type ConsentSettingUpdateParamsService struct {
	// Internal notes shown to admins in the dashboard. Not user-facing.
	InternalNotes string `json:"internalNotes" api:"required"`
	// Display name for this service in the preferences modal.
	Label string `json:"label" api:"required"`
	// Primary category value this service belongs to. Must match one of the top-level
	// `categories[].value` entries.
	Category param.Opt[string] `json:"category,omitzero"`
	// Extra category values this service belongs to. Each must match a
	// `categories[].value`.
	AdditionalCategories []string `json:"additionalCategories,omitzero"`
	// Domains/paths this service matches. Patterns matching the CMP's own scripts
	// (e.g. cdn.oursprivacy.com/cmp-init) are rejected to prevent the CMP blocking
	// itself — use a more specific path like cdn.oursprivacy.com/main.js to block a
	// specific script.
	DomainPatterns []string `json:"domainPatterns,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsService) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsService
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Set Enabled or Disabled. Enabling requires at least one allowed domain, either
// already saved or sent with this request.
type ConsentSettingUpdateParamsStatus string

const (
	ConsentSettingUpdateParamsStatusDisabled ConsentSettingUpdateParamsStatus = "Disabled"
	ConsentSettingUpdateParamsStatusEnabled  ConsentSettingUpdateParamsStatus = "Enabled"
)

type ConsentSettingAnalyticsParams struct {
	// Inclusive lower bound of the analytics window, as a UTC calendar day in
	// `YYYY-MM-DD` format. Count both selected dates: at most 180 days for `DAILY`,
	// 365 for `WEEKLY`, 365 for `MONTHLY`, or 14 for `HOURLY`.
	From string `query:"from" api:"required" json:"-"`
	// Inclusive upper bound of the analytics window, as a UTC calendar day in
	// `YYYY-MM-DD` format. Count both selected dates: at most 180 days for `DAILY`,
	// 365 for `WEEKLY`, 365 for `MONTHLY`, or 14 for `HOURLY`.
	To string `query:"to" api:"required" json:"-"`
	// Supported for DAILY and HOURLY. When `true`, each bucket also returns the
	// matching bucket from the immediately preceding window of equal length (in
	// `previous*` and `percentageChange*` fields). Defaults to `false`.
	CompareWithPreviousPeriod param.Opt[bool] `query:"compareWithPreviousPeriod,omitzero" json:"-"`
	// Optional custom comparison-window lower bound, as a UTC calendar day in
	// `YYYY-MM-DD` format. Provide together with `comparisonTo`; the window between
	// them must be the same length as `from`–`to` (a mismatch returns `400`). When
	// omitted (with `compareWithPreviousPeriod=true`), the immediately preceding
	// equal-length period is used.
	ComparisonFrom param.Opt[string] `query:"comparisonFrom,omitzero" json:"-"`
	// Optional custom comparison-window upper bound, as a UTC calendar day in
	// `YYYY-MM-DD` format. Provide together with `comparisonFrom`; the window must
	// match the `from`–`to` length. Ignored unless both comparison bounds are
	// supplied.
	ComparisonTo param.Opt[string] `query:"comparisonTo,omitzero" json:"-"`
	// JSON-encoded versioned analytics filter. Uses the shared analytics catalog
	// property IDs and operators, including event, session, and custom properties.
	Filter param.Opt[string] `query:"filter,omitzero" json:"-"`
	// Reporting interval: DAILY (default), WEEKLY (UTC weeks starting Monday), MONTHLY
	// (UTC calendar months), or HOURLY. Included-day limits: at most 180 days for
	// `DAILY`, 365 for `WEEKLY`, 365 for `MONTHLY`, or 14 for `HOURLY`. Breakdown
	// reports aggregate the entire selected range.
	//
	// Any of "DAILY", "HOURLY", "MONTHLY", "WEEKLY".
	Granularity ConsentSettingAnalyticsParamsGranularity `query:"granularity,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConsentSettingAnalyticsParams]'s query parameters as
// `url.Values`.
func (r ConsentSettingAnalyticsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Reporting interval: DAILY (default), WEEKLY (UTC weeks starting Monday), MONTHLY
// (UTC calendar months), or HOURLY. Included-day limits: at most 180 days for
// `DAILY`, 365 for `WEEKLY`, 365 for `MONTHLY`, or 14 for `HOURLY`. Breakdown
// reports aggregate the entire selected range.
type ConsentSettingAnalyticsParamsGranularity string

const (
	ConsentSettingAnalyticsParamsGranularityDaily   ConsentSettingAnalyticsParamsGranularity = "DAILY"
	ConsentSettingAnalyticsParamsGranularityHourly  ConsentSettingAnalyticsParamsGranularity = "HOURLY"
	ConsentSettingAnalyticsParamsGranularityMonthly ConsentSettingAnalyticsParamsGranularity = "MONTHLY"
	ConsentSettingAnalyticsParamsGranularityWeekly  ConsentSettingAnalyticsParamsGranularity = "WEEKLY"
)

type ConsentSettingOverviewParams struct {
	// First included UTC calendar day (YYYY-MM-DD).
	From time.Time `query:"from" api:"required" format:"date" json:"-"`
	// Last included UTC calendar day, at most 60 days including from.
	To time.Time `query:"to" api:"required" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [ConsentSettingOverviewParams]'s query parameters as
// `url.Values`.
func (r ConsentSettingOverviewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConsentSettingAnalyticsMonthlyParams struct {
	// Inclusive lower bound of the analytics window, as a UTC calendar day in
	// `YYYY-MM-DD` format. Count both selected dates: at most 180 days for `DAILY`,
	// 365 for `WEEKLY`, 365 for `MONTHLY`, or 14 for `HOURLY`.
	From string `query:"from" api:"required" json:"-"`
	// Inclusive upper bound of the analytics window, as a UTC calendar day in
	// `YYYY-MM-DD` format. Count both selected dates: at most 180 days for `DAILY`,
	// 365 for `WEEKLY`, 365 for `MONTHLY`, or 14 for `HOURLY`.
	To string `query:"to" api:"required" json:"-"`
	// JSON-encoded versioned analytics filter. Uses the shared analytics catalog
	// property IDs and operators, including event, session, and custom properties.
	Filter param.Opt[string] `query:"filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConsentSettingAnalyticsMonthlyParams]'s query parameters as
// `url.Values`.
func (r ConsentSettingAnalyticsMonthlyParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConsentSettingAnalyticsCapabilitiesParams struct {
	// Any of "PAGES", "SUMMARY", "TRENDS".
	Kind ConsentSettingAnalyticsCapabilitiesParamsKind `query:"kind,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConsentSettingAnalyticsCapabilitiesParams]'s query
// parameters as `url.Values`.
func (r ConsentSettingAnalyticsCapabilitiesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConsentSettingAnalyticsCapabilitiesParamsKind string

const (
	ConsentSettingAnalyticsCapabilitiesParamsKindPages   ConsentSettingAnalyticsCapabilitiesParamsKind = "PAGES"
	ConsentSettingAnalyticsCapabilitiesParamsKindSummary ConsentSettingAnalyticsCapabilitiesParamsKind = "SUMMARY"
	ConsentSettingAnalyticsCapabilitiesParamsKindTrends  ConsentSettingAnalyticsCapabilitiesParamsKind = "TRENDS"
)

type ConsentSettingPageAnalysisParams struct {
	// Inclusive lower bound of the analytics window, as a UTC calendar day in
	// `YYYY-MM-DD` format. Count both selected dates: at most 180 days for `DAILY`,
	// 365 for `WEEKLY`, 365 for `MONTHLY`, or 14 for `HOURLY`.
	From string `query:"from" api:"required" json:"-"`
	// Inclusive upper bound of the analytics window, as a UTC calendar day in
	// `YYYY-MM-DD` format. Count both selected dates: at most 180 days for `DAILY`,
	// 365 for `WEEKLY`, 365 for `MONTHLY`, or 14 for `HOURLY`.
	To string `query:"to" api:"required" json:"-"`
	// Maximum number of pages to return. Defaults to 50; clamped to 1–200.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Opaque cursor from pagination.nextCursor. Omit for the first page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// JSON-encoded versioned analytics filter. Uses the shared analytics catalog
	// property IDs and operators, including event, session, and custom properties.
	Filter param.Opt[string] `query:"filter,omitzero" json:"-"`
	// Case-sensitive substring match against the page path.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Reporting interval: DAILY (default), WEEKLY (UTC weeks starting Monday), MONTHLY
	// (UTC calendar months), or HOURLY. Included-day limits: at most 180 days for
	// `DAILY`, 365 for `WEEKLY`, 365 for `MONTHLY`, or 14 for `HOURLY`. Breakdown
	// reports aggregate the entire selected range.
	//
	// Any of "DAILY", "HOURLY", "MONTHLY", "WEEKLY".
	Granularity ConsentSettingPageAnalysisParamsGranularity `query:"granularity,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConsentSettingPageAnalysisParams]'s query parameters as
// `url.Values`.
func (r ConsentSettingPageAnalysisParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Reporting interval: DAILY (default), WEEKLY (UTC weeks starting Monday), MONTHLY
// (UTC calendar months), or HOURLY. Included-day limits: at most 180 days for
// `DAILY`, 365 for `WEEKLY`, 365 for `MONTHLY`, or 14 for `HOURLY`. Breakdown
// reports aggregate the entire selected range.
type ConsentSettingPageAnalysisParamsGranularity string

const (
	ConsentSettingPageAnalysisParamsGranularityDaily   ConsentSettingPageAnalysisParamsGranularity = "DAILY"
	ConsentSettingPageAnalysisParamsGranularityHourly  ConsentSettingPageAnalysisParamsGranularity = "HOURLY"
	ConsentSettingPageAnalysisParamsGranularityMonthly ConsentSettingPageAnalysisParamsGranularity = "MONTHLY"
	ConsentSettingPageAnalysisParamsGranularityWeekly  ConsentSettingPageAnalysisParamsGranularity = "WEEKLY"
)

type ConsentSettingAnalyticsByRegionParams struct {
	// Inclusive lower bound of the analytics window, as a UTC calendar day in
	// `YYYY-MM-DD` format. Count both selected dates: at most 180 days for `DAILY`,
	// 365 for `WEEKLY`, 365 for `MONTHLY`, or 14 for `HOURLY`.
	From string `query:"from" api:"required" json:"-"`
	// Inclusive upper bound of the analytics window, as a UTC calendar day in
	// `YYYY-MM-DD` format. Count both selected dates: at most 180 days for `DAILY`,
	// 365 for `WEEKLY`, 365 for `MONTHLY`, or 14 for `HOURLY`.
	To string `query:"to" api:"required" json:"-"`
	// JSON-encoded versioned analytics filter. Uses the shared analytics catalog
	// property IDs and operators, including event, session, and custom properties.
	Filter param.Opt[string] `query:"filter,omitzero" json:"-"`
	// Reporting interval: DAILY (default), WEEKLY (UTC weeks starting Monday), MONTHLY
	// (UTC calendar months), or HOURLY. Included-day limits: at most 180 days for
	// `DAILY`, 365 for `WEEKLY`, 365 for `MONTHLY`, or 14 for `HOURLY`. Breakdown
	// reports aggregate the entire selected range.
	//
	// Any of "DAILY", "HOURLY", "MONTHLY", "WEEKLY".
	Granularity ConsentSettingAnalyticsByRegionParamsGranularity `query:"granularity,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConsentSettingAnalyticsByRegionParams]'s query parameters
// as `url.Values`.
func (r ConsentSettingAnalyticsByRegionParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Reporting interval: DAILY (default), WEEKLY (UTC weeks starting Monday), MONTHLY
// (UTC calendar months), or HOURLY. Included-day limits: at most 180 days for
// `DAILY`, 365 for `WEEKLY`, 365 for `MONTHLY`, or 14 for `HOURLY`. Breakdown
// reports aggregate the entire selected range.
type ConsentSettingAnalyticsByRegionParamsGranularity string

const (
	ConsentSettingAnalyticsByRegionParamsGranularityDaily   ConsentSettingAnalyticsByRegionParamsGranularity = "DAILY"
	ConsentSettingAnalyticsByRegionParamsGranularityHourly  ConsentSettingAnalyticsByRegionParamsGranularity = "HOURLY"
	ConsentSettingAnalyticsByRegionParamsGranularityMonthly ConsentSettingAnalyticsByRegionParamsGranularity = "MONTHLY"
	ConsentSettingAnalyticsByRegionParamsGranularityWeekly  ConsentSettingAnalyticsByRegionParamsGranularity = "WEEKLY"
)
