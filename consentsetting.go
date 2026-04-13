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
	ID        string `json:"id" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	Kind      string `json:"kind" api:"required"`
	Name      string `json:"name" api:"required"`
	// Any of "Disabled", "Enabled".
	Status    string `json:"status" api:"required"`
	UpdatedAt string `json:"updatedAt" api:"nullable"`
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
func (r ConsentSettingListResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingListResponseEntity) UnmarshalJSON(data []byte) error {
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
