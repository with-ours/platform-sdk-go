// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomwithoursplatformsdkgo

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

// Create a new consent setting. Requires scope: consentSettings:create
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

// Update a consent setting. Requires scope: consentSettings:update
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

// List all consent settings. Requires scope: consentSettings:list
func (r *ConsentSettingService) List(ctx context.Context, opts ...option.RequestOption) (res *ConsentSettingListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/consent-settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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

type ConsentSettingNewResponse struct {
	IsSuccess       bool   `json:"isSuccess" api:"required"`
	Cause           string `json:"cause" api:"nullable"`
	ConsentSettings any    `json:"consentSettings" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsSuccess       respjson.Field
		Cause           respjson.Field
		ConsentSettings respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingGetResponse = any

type ConsentSettingUpdateResponse struct {
	IsSuccess       bool   `json:"isSuccess" api:"required"`
	Cause           string `json:"cause" api:"nullable"`
	ConsentSettings any    `json:"consentSettings" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsSuccess       respjson.Field
		Cause           respjson.Field
		ConsentSettings respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	ID         string                                     `json:"id" api:"required"`
	Categories []ConsentSettingListResponseEntityCategory `json:"categories" api:"required"`
	CreatedAt  string                                     `json:"createdAt" api:"required"`
	Default    ConsentSettingListResponseEntityDefault    `json:"default" api:"required"`
	Kind       string                                     `json:"kind" api:"required"`
	Name       string                                     `json:"name" api:"required"`
	Regions    []ConsentSettingListResponseEntityRegion   `json:"regions" api:"required"`
	Services   []ConsentSettingListResponseEntityService  `json:"services" api:"required"`
	// Any of "Disabled", "Enabled".
	Status                 string  `json:"status" api:"required"`
	ConsentCookieName      string  `json:"consentCookieName" api:"nullable"`
	CustomDomain           string  `json:"customDomain" api:"nullable"`
	Revision               float64 `json:"revision" api:"nullable"`
	SkipBlockingClassNames []any   `json:"skipBlockingClassNames" api:"nullable"`
	UpdatedAt              string  `json:"updatedAt" api:"nullable"`
	WebSDKToken            string  `json:"webSDKToken" api:"nullable"`
	WhitelistDomains       []any   `json:"whitelistDomains" api:"nullable"`
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
	Label    string `json:"label" api:"required"`
	Priority int64  `json:"priority" api:"required"`
	Value    string `json:"value" api:"required"`
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

type ConsentSettingListResponseEntityDefault struct {
	Categories []ConsentSettingListResponseEntityDefaultCategory `json:"categories" api:"required"`
	Language   string                                            `json:"language" api:"required"`
	// Any of "opt_in", "opt_out".
	Mode                     string                                               `json:"mode" api:"required"`
	Translations             []ConsentSettingListResponseEntityDefaultTranslation `json:"translations" api:"required"`
	AutoblockUnknown         bool                                                 `json:"autoblockUnknown" api:"nullable"`
	AutoShow                 bool                                                 `json:"autoShow" api:"nullable"`
	AutoShowDismissConfig    any                                                  `json:"autoShowDismissConfig" api:"nullable"`
	AutoShowDismissMode      string                                               `json:"autoShowDismissMode" api:"nullable"`
	DisablePageInteraction   bool                                                 `json:"disablePageInteraction" api:"nullable"`
	GuiOptions               any                                                  `json:"guiOptions" api:"nullable"`
	HideFromBots             bool                                                 `json:"hideFromBots" api:"nullable"`
	ShowVendorsInPreferences bool                                                 `json:"showVendorsInPreferences" api:"nullable"`
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
	Enabled          bool `json:"enabled" api:"required"`
	AutoDisableOnGpc bool `json:"autoDisableOnGPC" api:"nullable"`
	ReadOnly         bool `json:"readOnly" api:"nullable"`
	ReloadPage       bool `json:"reloadPage" api:"nullable"`
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
	ConsentModal     any `json:"consentModal" api:"nullable"`
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
	RegionCode        string                                     `json:"regionCode" api:"required"`
	Rule              ConsentSettingListResponseEntityRegionRule `json:"rule" api:"required"`
	AdditionalRegions []any                                      `json:"additionalRegions" api:"nullable"`
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
	Categories []ConsentSettingListResponseEntityRegionRuleCategory `json:"categories" api:"required"`
	Language   string                                               `json:"language" api:"required"`
	// Any of "opt_in", "opt_out".
	Mode                     string                                                  `json:"mode" api:"required"`
	Translations             []ConsentSettingListResponseEntityRegionRuleTranslation `json:"translations" api:"required"`
	AutoblockUnknown         bool                                                    `json:"autoblockUnknown" api:"nullable"`
	AutoShow                 bool                                                    `json:"autoShow" api:"nullable"`
	AutoShowDismissConfig    any                                                     `json:"autoShowDismissConfig" api:"nullable"`
	AutoShowDismissMode      string                                                  `json:"autoShowDismissMode" api:"nullable"`
	DisablePageInteraction   bool                                                    `json:"disablePageInteraction" api:"nullable"`
	GuiOptions               any                                                     `json:"guiOptions" api:"nullable"`
	HideFromBots             bool                                                    `json:"hideFromBots" api:"nullable"`
	ShowVendorsInPreferences bool                                                    `json:"showVendorsInPreferences" api:"nullable"`
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
	Enabled          bool `json:"enabled" api:"required"`
	AutoDisableOnGpc bool `json:"autoDisableOnGPC" api:"nullable"`
	ReadOnly         bool `json:"readOnly" api:"nullable"`
	ReloadPage       bool `json:"reloadPage" api:"nullable"`
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
	ConsentModal     any `json:"consentModal" api:"nullable"`
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
	InternalNotes        string `json:"internalNotes" api:"required"`
	Label                string `json:"label" api:"required"`
	AdditionalCategories []any  `json:"additionalCategories" api:"nullable"`
	Category             string `json:"category" api:"nullable"`
	DomainPatterns       []any  `json:"domainPatterns" api:"nullable"`
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

type ConsentSettingDeleteResponse struct {
	ID        string `json:"id" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	Kind      string `json:"kind" api:"required"`
	Name      string `json:"name" api:"required"`
	// Any of "Disabled", "Enabled".
	Status    ConsentSettingDeleteResponseStatus `json:"status" api:"required"`
	UpdatedAt string                             `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Kind        respjson.Field
		Name        respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingDeleteResponseStatus string

const (
	ConsentSettingDeleteResponseStatusDisabled ConsentSettingDeleteResponseStatus = "Disabled"
	ConsentSettingDeleteResponseStatusEnabled  ConsentSettingDeleteResponseStatus = "Enabled"
)

type ConsentSettingUpdateParams struct {
	Categories []ConsentSettingUpdateParamsCategory `json:"categories,omitzero" api:"required"`
	Default    ConsentSettingUpdateParamsDefault    `json:"default,omitzero" api:"required"`
	Name       string                               `json:"name" api:"required"`
	Regions    []ConsentSettingUpdateParamsRegion   `json:"regions,omitzero" api:"required"`
	Services   []ConsentSettingUpdateParamsService  `json:"services,omitzero" api:"required"`
	// Any of "Disabled", "Enabled".
	Status                 ConsentSettingUpdateParamsStatus `json:"status,omitzero" api:"required"`
	ConsentCookieName      param.Opt[string]                `json:"consentCookieName,omitzero"`
	CustomDomain           param.Opt[string]                `json:"customDomain,omitzero"`
	Revision               param.Opt[float64]               `json:"revision,omitzero"`
	WebSDKToken            param.Opt[string]                `json:"webSDKToken,omitzero"`
	SkipBlockingClassNames []any                            `json:"skipBlockingClassNames,omitzero"`
	WhitelistDomains       []any                            `json:"whitelistDomains,omitzero"`
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
	Label    string `json:"label" api:"required"`
	Priority int64  `json:"priority" api:"required"`
	Value    string `json:"value" api:"required"`
	paramObj
}

func (r ConsentSettingUpdateParamsCategory) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsCategory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Categories, Language, Mode, Translations are required.
type ConsentSettingUpdateParamsDefault struct {
	Categories []ConsentSettingUpdateParamsDefaultCategory `json:"categories,omitzero" api:"required"`
	Language   string                                      `json:"language" api:"required"`
	// Any of "opt_in", "opt_out".
	Mode                     string                                         `json:"mode,omitzero" api:"required"`
	Translations             []ConsentSettingUpdateParamsDefaultTranslation `json:"translations,omitzero" api:"required"`
	AutoblockUnknown         param.Opt[bool]                                `json:"autoblockUnknown,omitzero"`
	AutoShow                 param.Opt[bool]                                `json:"autoShow,omitzero"`
	AutoShowDismissMode      param.Opt[string]                              `json:"autoShowDismissMode,omitzero"`
	DisablePageInteraction   param.Opt[bool]                                `json:"disablePageInteraction,omitzero"`
	HideFromBots             param.Opt[bool]                                `json:"hideFromBots,omitzero"`
	ShowVendorsInPreferences param.Opt[bool]                                `json:"showVendorsInPreferences,omitzero"`
	AutoShowDismissConfig    any                                            `json:"autoShowDismissConfig,omitzero"`
	GuiOptions               any                                            `json:"guiOptions,omitzero"`
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
	Enabled          bool            `json:"enabled" api:"required"`
	AutoDisableOnGpc param.Opt[bool] `json:"autoDisableOnGPC,omitzero"`
	ReadOnly         param.Opt[bool] `json:"readOnly,omitzero"`
	ReloadPage       param.Opt[bool] `json:"reloadPage,omitzero"`
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
	ConsentModal     any `json:"consentModal,omitzero"`
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
	RegionCode        string                               `json:"regionCode" api:"required"`
	Rule              ConsentSettingUpdateParamsRegionRule `json:"rule,omitzero" api:"required"`
	AdditionalRegions []any                                `json:"additionalRegions,omitzero"`
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
	Categories []ConsentSettingUpdateParamsRegionRuleCategory `json:"categories,omitzero" api:"required"`
	Language   string                                         `json:"language" api:"required"`
	// Any of "opt_in", "opt_out".
	Mode                     string                                            `json:"mode,omitzero" api:"required"`
	Translations             []ConsentSettingUpdateParamsRegionRuleTranslation `json:"translations,omitzero" api:"required"`
	AutoblockUnknown         param.Opt[bool]                                   `json:"autoblockUnknown,omitzero"`
	AutoShow                 param.Opt[bool]                                   `json:"autoShow,omitzero"`
	AutoShowDismissMode      param.Opt[string]                                 `json:"autoShowDismissMode,omitzero"`
	DisablePageInteraction   param.Opt[bool]                                   `json:"disablePageInteraction,omitzero"`
	HideFromBots             param.Opt[bool]                                   `json:"hideFromBots,omitzero"`
	ShowVendorsInPreferences param.Opt[bool]                                   `json:"showVendorsInPreferences,omitzero"`
	AutoShowDismissConfig    any                                               `json:"autoShowDismissConfig,omitzero"`
	GuiOptions               any                                               `json:"guiOptions,omitzero"`
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
	Enabled          bool            `json:"enabled" api:"required"`
	AutoDisableOnGpc param.Opt[bool] `json:"autoDisableOnGPC,omitzero"`
	ReadOnly         param.Opt[bool] `json:"readOnly,omitzero"`
	ReloadPage       param.Opt[bool] `json:"reloadPage,omitzero"`
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
	ConsentModal     any `json:"consentModal,omitzero"`
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
	InternalNotes        string            `json:"internalNotes" api:"required"`
	Label                string            `json:"label" api:"required"`
	Category             param.Opt[string] `json:"category,omitzero"`
	AdditionalCategories []any             `json:"additionalCategories,omitzero"`
	DomainPatterns       []any             `json:"domainPatterns,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsService) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsService
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingUpdateParamsStatus string

const (
	ConsentSettingUpdateParamsStatusDisabled ConsentSettingUpdateParamsStatus = "Disabled"
	ConsentSettingUpdateParamsStatusEnabled  ConsentSettingUpdateParamsStatus = "Enabled"
)
