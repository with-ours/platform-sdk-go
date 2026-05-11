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
	"github.com/with-ours/platform-sdk-go/internal/requestconfig"
	"github.com/with-ours/platform-sdk-go/option"
	"github.com/with-ours/platform-sdk-go/packages/param"
	"github.com/with-ours/platform-sdk-go/packages/respjson"
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

// List all consent settings. Requires scope: consentSettings:list
func (r *ConsentSettingService) List(ctx context.Context, opts ...option.RequestOption) (res *ConsentSettingListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/consent-settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Create a new consent settings record. POST takes no request body — the server
// initializes the record with defaults (Disabled status, opt-out default rule,
// English translations, necessary/analytics/advertising categories, no regions, no
// whitelisted domains). Configure the record afterward with PATCH (partial update)
// or PUT (full replacement). Returns the same shape as GET so you can read the
// server-assigned `id`, default rule, and categories without a follow-up fetch.
// Requires scope: consentSettings:create
func (r *ConsentSettingService) New(ctx context.Context, opts ...option.RequestOption) (res *ConsentSettingNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/consent-settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Find a single consent setting by ID. Requires scope: consentSettings:find
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

// Replace a consent setting. Send the full ConsentSettingsInput body — omitted
// optional fields are reset. Use PATCH for partial updates. Requires scope:
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

// Partially update a consent setting. Send only the fields you want to change —
// every field is optional and unspecified fields are preserved. List-valued fields
// (services, categories, regions) are replaced wholesale when sent. Requires
// scope: consentSettings:update
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

// Delete a consent setting. Requires scope: consentSettings:delete
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
	// Enabled means the CMP serves on whitelisted domains; Disabled means it does not.
	//
	// Any of "Disabled", "Enabled".
	Status string `json:"status" api:"required"`
	// Name of the cookie that stores the user's consent state. Defaults to
	// "op_consent".
	ConsentCookieName string `json:"consentCookieName" api:"nullable"`
	// Optional custom CDN domain for serving the CMP script (e.g.
	// consent.example.com).
	CustomDomain string `json:"customDomain" api:"nullable"`
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
	// Allowlist of domains where this CMP configuration may run. Used at runtime to
	// derive the broadest matching base domain so consent can persist across matching
	// subdomains.
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
	AutoShowDismissConfig any `json:"autoShowDismissConfig" api:"nullable"`
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	AutoShowDismissMode string `json:"autoShowDismissMode" api:"nullable"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction bool `json:"disablePageInteraction" api:"nullable"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions any `json:"guiOptions" api:"nullable"`
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
	ConsentModal any `json:"consentModal" api:"nullable"`
	// Translated copy for the preferences modal.
	PreferencesModal any `json:"preferencesModal" api:"nullable"`
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RegionCode        respjson.Field
		Rule              respjson.Field
		AdditionalRegions respjson.Field
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
	AutoShowDismissConfig any `json:"autoShowDismissConfig" api:"nullable"`
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	AutoShowDismissMode string `json:"autoShowDismissMode" api:"nullable"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction bool `json:"disablePageInteraction" api:"nullable"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions any `json:"guiOptions" api:"nullable"`
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
	ConsentModal any `json:"consentModal" api:"nullable"`
	// Translated copy for the preferences modal.
	PreferencesModal any `json:"preferencesModal" api:"nullable"`
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
	// Enabled means the CMP serves on whitelisted domains; Disabled means it does not.
	//
	// Any of "Disabled", "Enabled".
	Status ConsentSettingNewResponseStatus `json:"status" api:"required"`
	// Name of the cookie that stores the user's consent state. Defaults to
	// "op_consent".
	ConsentCookieName string `json:"consentCookieName" api:"nullable"`
	// Optional custom CDN domain for serving the CMP script (e.g.
	// consent.example.com).
	CustomDomain string `json:"customDomain" api:"nullable"`
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
	// Allowlist of domains where this CMP configuration may run. Used at runtime to
	// derive the broadest matching base domain so consent can persist across matching
	// subdomains.
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
	AutoShowDismissConfig any `json:"autoShowDismissConfig" api:"nullable"`
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	AutoShowDismissMode string `json:"autoShowDismissMode" api:"nullable"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction bool `json:"disablePageInteraction" api:"nullable"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions any `json:"guiOptions" api:"nullable"`
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
	ConsentModal any `json:"consentModal" api:"nullable"`
	// Translated copy for the preferences modal.
	PreferencesModal any `json:"preferencesModal" api:"nullable"`
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RegionCode        respjson.Field
		Rule              respjson.Field
		AdditionalRegions respjson.Field
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
	AutoShowDismissConfig any `json:"autoShowDismissConfig" api:"nullable"`
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	AutoShowDismissMode string `json:"autoShowDismissMode" api:"nullable"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction bool `json:"disablePageInteraction" api:"nullable"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions any `json:"guiOptions" api:"nullable"`
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
	ConsentModal any `json:"consentModal" api:"nullable"`
	// Translated copy for the preferences modal.
	PreferencesModal any `json:"preferencesModal" api:"nullable"`
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

// Enabled means the CMP serves on whitelisted domains; Disabled means it does not.
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
	// Enabled means the CMP serves on whitelisted domains; Disabled means it does not.
	//
	// Any of "Disabled", "Enabled".
	Status ConsentSettingGetResponseStatus `json:"status" api:"required"`
	// Name of the cookie that stores the user's consent state. Defaults to
	// "op_consent".
	ConsentCookieName string `json:"consentCookieName" api:"nullable"`
	// Optional custom CDN domain for serving the CMP script (e.g.
	// consent.example.com).
	CustomDomain string `json:"customDomain" api:"nullable"`
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
	// Allowlist of domains where this CMP configuration may run. Used at runtime to
	// derive the broadest matching base domain so consent can persist across matching
	// subdomains.
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
	AutoShowDismissConfig any `json:"autoShowDismissConfig" api:"nullable"`
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	AutoShowDismissMode string `json:"autoShowDismissMode" api:"nullable"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction bool `json:"disablePageInteraction" api:"nullable"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions any `json:"guiOptions" api:"nullable"`
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
	ConsentModal any `json:"consentModal" api:"nullable"`
	// Translated copy for the preferences modal.
	PreferencesModal any `json:"preferencesModal" api:"nullable"`
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RegionCode        respjson.Field
		Rule              respjson.Field
		AdditionalRegions respjson.Field
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
	AutoShowDismissConfig any `json:"autoShowDismissConfig" api:"nullable"`
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	AutoShowDismissMode string `json:"autoShowDismissMode" api:"nullable"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction bool `json:"disablePageInteraction" api:"nullable"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions any `json:"guiOptions" api:"nullable"`
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
	ConsentModal any `json:"consentModal" api:"nullable"`
	// Translated copy for the preferences modal.
	PreferencesModal any `json:"preferencesModal" api:"nullable"`
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

// Enabled means the CMP serves on whitelisted domains; Disabled means it does not.
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
	// Enabled means the CMP serves on whitelisted domains; Disabled means it does not.
	//
	// Any of "Disabled", "Enabled".
	Status ConsentSettingReplaceResponseStatus `json:"status" api:"required"`
	// Name of the cookie that stores the user's consent state. Defaults to
	// "op_consent".
	ConsentCookieName string `json:"consentCookieName" api:"nullable"`
	// Optional custom CDN domain for serving the CMP script (e.g.
	// consent.example.com).
	CustomDomain string `json:"customDomain" api:"nullable"`
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
	// Allowlist of domains where this CMP configuration may run. Used at runtime to
	// derive the broadest matching base domain so consent can persist across matching
	// subdomains.
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
	AutoShowDismissConfig any `json:"autoShowDismissConfig" api:"nullable"`
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	AutoShowDismissMode string `json:"autoShowDismissMode" api:"nullable"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction bool `json:"disablePageInteraction" api:"nullable"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions any `json:"guiOptions" api:"nullable"`
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
	ConsentModal any `json:"consentModal" api:"nullable"`
	// Translated copy for the preferences modal.
	PreferencesModal any `json:"preferencesModal" api:"nullable"`
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RegionCode        respjson.Field
		Rule              respjson.Field
		AdditionalRegions respjson.Field
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
	AutoShowDismissConfig any `json:"autoShowDismissConfig" api:"nullable"`
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	AutoShowDismissMode string `json:"autoShowDismissMode" api:"nullable"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction bool `json:"disablePageInteraction" api:"nullable"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions any `json:"guiOptions" api:"nullable"`
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
	ConsentModal any `json:"consentModal" api:"nullable"`
	// Translated copy for the preferences modal.
	PreferencesModal any `json:"preferencesModal" api:"nullable"`
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

// Enabled means the CMP serves on whitelisted domains; Disabled means it does not.
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
	// Enabled means the CMP serves on whitelisted domains; Disabled means it does not.
	//
	// Any of "Disabled", "Enabled".
	Status ConsentSettingUpdateResponseStatus `json:"status" api:"required"`
	// Name of the cookie that stores the user's consent state. Defaults to
	// "op_consent".
	ConsentCookieName string `json:"consentCookieName" api:"nullable"`
	// Optional custom CDN domain for serving the CMP script (e.g.
	// consent.example.com).
	CustomDomain string `json:"customDomain" api:"nullable"`
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
	// Allowlist of domains where this CMP configuration may run. Used at runtime to
	// derive the broadest matching base domain so consent can persist across matching
	// subdomains.
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
	AutoShowDismissConfig any `json:"autoShowDismissConfig" api:"nullable"`
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	AutoShowDismissMode string `json:"autoShowDismissMode" api:"nullable"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction bool `json:"disablePageInteraction" api:"nullable"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions any `json:"guiOptions" api:"nullable"`
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
	ConsentModal any `json:"consentModal" api:"nullable"`
	// Translated copy for the preferences modal.
	PreferencesModal any `json:"preferencesModal" api:"nullable"`
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RegionCode        respjson.Field
		Rule              respjson.Field
		AdditionalRegions respjson.Field
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
	AutoShowDismissConfig any `json:"autoShowDismissConfig" api:"nullable"`
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	AutoShowDismissMode string `json:"autoShowDismissMode" api:"nullable"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction bool `json:"disablePageInteraction" api:"nullable"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions any `json:"guiOptions" api:"nullable"`
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
	ConsentModal any `json:"consentModal" api:"nullable"`
	// Translated copy for the preferences modal.
	PreferencesModal any `json:"preferencesModal" api:"nullable"`
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

// Enabled means the CMP serves on whitelisted domains; Disabled means it does not.
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
	// Enabled means the CMP serves on whitelisted domains; Disabled means it does not.
	//
	// Any of "Disabled", "Enabled".
	Status ConsentSettingDeleteResponseStatus `json:"status" api:"required"`
	// Name of the cookie that stores the user's consent state. Defaults to
	// "op_consent".
	ConsentCookieName string `json:"consentCookieName" api:"nullable"`
	// Optional custom CDN domain for serving the CMP script (e.g.
	// consent.example.com).
	CustomDomain string `json:"customDomain" api:"nullable"`
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
	// Allowlist of domains where this CMP configuration may run. Used at runtime to
	// derive the broadest matching base domain so consent can persist across matching
	// subdomains.
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
	AutoShowDismissConfig any `json:"autoShowDismissConfig" api:"nullable"`
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	AutoShowDismissMode string `json:"autoShowDismissMode" api:"nullable"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction bool `json:"disablePageInteraction" api:"nullable"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions any `json:"guiOptions" api:"nullable"`
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
	ConsentModal any `json:"consentModal" api:"nullable"`
	// Translated copy for the preferences modal.
	PreferencesModal any `json:"preferencesModal" api:"nullable"`
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RegionCode        respjson.Field
		Rule              respjson.Field
		AdditionalRegions respjson.Field
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
	AutoShowDismissConfig any `json:"autoShowDismissConfig" api:"nullable"`
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	AutoShowDismissMode string `json:"autoShowDismissMode" api:"nullable"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction bool `json:"disablePageInteraction" api:"nullable"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions any `json:"guiOptions" api:"nullable"`
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
	ConsentModal any `json:"consentModal" api:"nullable"`
	// Translated copy for the preferences modal.
	PreferencesModal any `json:"preferencesModal" api:"nullable"`
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

// Enabled means the CMP serves on whitelisted domains; Disabled means it does not.
type ConsentSettingDeleteResponseStatus string

const (
	ConsentSettingDeleteResponseStatusDisabled ConsentSettingDeleteResponseStatus = "Disabled"
	ConsentSettingDeleteResponseStatusEnabled  ConsentSettingDeleteResponseStatus = "Enabled"
)

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
	// Enabled to serve the CMP, Disabled to take it offline.
	//
	// Any of "Disabled", "Enabled".
	Status ConsentSettingReplaceParamsStatus `json:"status,omitzero" api:"required"`
	// Name of the cookie that stores consent state. Pass null to clear (defaults to
	// "op_consent").
	ConsentCookieName param.Opt[string] `json:"consentCookieName,omitzero"`
	// Custom CDN domain for serving the CMP script. Pass null to clear.
	CustomDomain param.Opt[string] `json:"customDomain,omitzero"`
	// Revision counter. Bump to re-prompt users who already consented.
	Revision param.Opt[float64] `json:"revision,omitzero"`
	// Pixel of the WebSource this CMP is wired into. Pass null to clear the link.
	WebSDKToken param.Opt[string] `json:"webSDKToken,omitzero"`
	// CSS class names that opt scripts out of consent blocking. Each must be a single
	// class token.
	SkipBlockingClassNames []string `json:"skipBlockingClassNames,omitzero"`
	// Allowlist of domains where this CMP runs. Pass null/[] to clear.
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
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	AutoShowDismissMode param.Opt[string] `json:"autoShowDismissMode,omitzero"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction param.Opt[bool] `json:"disablePageInteraction,omitzero"`
	// When true, the modal is suppressed for known bot user agents.
	HideFromBots param.Opt[bool] `json:"hideFromBots,omitzero"`
	// When true, the per-service list (services[]) is rendered inside the preferences
	// modal.
	ShowVendorsInPreferences param.Opt[bool] `json:"showVendorsInPreferences,omitzero"`
	// Threshold config for autoShowDismissMode (page count or seconds).
	AutoShowDismissConfig any `json:"autoShowDismissConfig,omitzero"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions any `json:"guiOptions,omitzero"`
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
	ConsentModal any `json:"consentModal,omitzero"`
	// Translated copy for the preferences modal.
	PreferencesModal any `json:"preferencesModal,omitzero"`
	paramObj
}

func (r ConsentSettingReplaceParamsDefaultTranslationValue) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsDefaultTranslationValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsDefaultTranslationValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties RegionCode, Rule are required.
type ConsentSettingReplaceParamsRegion struct {
	// Region this rule applies to. Use ISO 3166-1 alpha-2 country code ("US", "DE",
	// "BR") or country-subdivision code ("US-CA", "GB-ENG", "CA-ON"). Each region code
	// may appear in only one rule across `regions[]`.
	RegionCode string                                `json:"regionCode" api:"required"`
	Rule       ConsentSettingReplaceParamsRegionRule `json:"rule,omitzero" api:"required"`
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
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	AutoShowDismissMode param.Opt[string] `json:"autoShowDismissMode,omitzero"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction param.Opt[bool] `json:"disablePageInteraction,omitzero"`
	// When true, the modal is suppressed for known bot user agents.
	HideFromBots param.Opt[bool] `json:"hideFromBots,omitzero"`
	// When true, the per-service list (services[]) is rendered inside the preferences
	// modal.
	ShowVendorsInPreferences param.Opt[bool] `json:"showVendorsInPreferences,omitzero"`
	// Threshold config for autoShowDismissMode (page count or seconds).
	AutoShowDismissConfig any `json:"autoShowDismissConfig,omitzero"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions any `json:"guiOptions,omitzero"`
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
	ConsentModal any `json:"consentModal,omitzero"`
	// Translated copy for the preferences modal.
	PreferencesModal any `json:"preferencesModal,omitzero"`
	paramObj
}

func (r ConsentSettingReplaceParamsRegionRuleTranslationValue) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingReplaceParamsRegionRuleTranslationValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingReplaceParamsRegionRuleTranslationValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

// Enabled to serve the CMP, Disabled to take it offline.
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
	// Bump the revision counter to re-prompt users.
	Revision param.Opt[float64] `json:"revision,omitzero"`
	// Set or clear the WebSource pixel link. A non-null token must be a valid
	// WebSource of yours.
	WebSDKToken param.Opt[string] `json:"webSDKToken,omitzero"`
	// Rename the consent settings record.
	Name param.Opt[string] `json:"name,omitzero"`
	// Replace the skipBlockingClassNames list. Pass null/[] to clear.
	SkipBlockingClassNames []string `json:"skipBlockingClassNames,omitzero"`
	// Replace the allowlist. Pass null/[] to clear.
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
	// Toggle Enabled/Disabled without re-sending the rest of the config.
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
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	AutoShowDismissMode param.Opt[string] `json:"autoShowDismissMode,omitzero"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction param.Opt[bool] `json:"disablePageInteraction,omitzero"`
	// When true, the modal is suppressed for known bot user agents.
	HideFromBots param.Opt[bool] `json:"hideFromBots,omitzero"`
	// When true, the per-service list (services[]) is rendered inside the preferences
	// modal.
	ShowVendorsInPreferences param.Opt[bool] `json:"showVendorsInPreferences,omitzero"`
	// Threshold config for autoShowDismissMode (page count or seconds).
	AutoShowDismissConfig any `json:"autoShowDismissConfig,omitzero"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions any `json:"guiOptions,omitzero"`
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
	ConsentModal any `json:"consentModal,omitzero"`
	// Translated copy for the preferences modal.
	PreferencesModal any `json:"preferencesModal,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsDefaultTranslationValue) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsDefaultTranslationValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsDefaultTranslationValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties RegionCode, Rule are required.
type ConsentSettingUpdateParamsRegion struct {
	// Region this rule applies to. Use ISO 3166-1 alpha-2 country code ("US", "DE",
	// "BR") or country-subdivision code ("US-CA", "GB-ENG", "CA-ON"). Each region code
	// may appear in only one rule across `regions[]`.
	RegionCode string                               `json:"regionCode" api:"required"`
	Rule       ConsentSettingUpdateParamsRegionRule `json:"rule,omitzero" api:"required"`
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
	// How the modal is treated as dismissed (never, after_pages, after_seconds).
	AutoShowDismissMode param.Opt[string] `json:"autoShowDismissMode,omitzero"`
	// When true, the rest of the page is locked behind a backdrop until the user
	// chooses.
	DisablePageInteraction param.Opt[bool] `json:"disablePageInteraction,omitzero"`
	// When true, the modal is suppressed for known bot user agents.
	HideFromBots param.Opt[bool] `json:"hideFromBots,omitzero"`
	// When true, the per-service list (services[]) is rendered inside the preferences
	// modal.
	ShowVendorsInPreferences param.Opt[bool] `json:"showVendorsInPreferences,omitzero"`
	// Threshold config for autoShowDismissMode (page count or seconds).
	AutoShowDismissConfig any `json:"autoShowDismissConfig,omitzero"`
	// Visual options for the modals (layout/position/colors).
	GuiOptions any `json:"guiOptions,omitzero"`
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
	ConsentModal any `json:"consentModal,omitzero"`
	// Translated copy for the preferences modal.
	PreferencesModal any `json:"preferencesModal,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsRegionRuleTranslationValue) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsRegionRuleTranslationValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsRegionRuleTranslationValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

// Toggle Enabled/Disabled without re-sending the rest of the config.
type ConsentSettingUpdateParamsStatus string

const (
	ConsentSettingUpdateParamsStatusDisabled ConsentSettingUpdateParamsStatus = "Disabled"
	ConsentSettingUpdateParamsStatusEnabled  ConsentSettingUpdateParamsStatus = "Enabled"
)
