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
	IsSuccess       bool                                     `json:"isSuccess" api:"required"`
	Cause           string                                   `json:"cause" api:"nullable"`
	ConsentSettings ConsentSettingNewResponseConsentSettings `json:"consentSettings" api:"nullable"`
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

type ConsentSettingNewResponseConsentSettings struct {
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
func (r ConsentSettingNewResponseConsentSettings) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingNewResponseConsentSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingGetResponse struct {
	ID        string `json:"id" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	Kind      string `json:"kind" api:"required"`
	Name      string `json:"name" api:"required"`
	// Any of "Disabled", "Enabled".
	Status    ConsentSettingGetResponseStatus `json:"status" api:"required"`
	UpdatedAt string                          `json:"updatedAt" api:"nullable"`
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
func (r ConsentSettingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingGetResponseStatus string

const (
	ConsentSettingGetResponseStatusDisabled ConsentSettingGetResponseStatus = "Disabled"
	ConsentSettingGetResponseStatusEnabled  ConsentSettingGetResponseStatus = "Enabled"
)

type ConsentSettingUpdateResponse struct {
	IsSuccess       bool                                        `json:"isSuccess" api:"required"`
	Cause           string                                      `json:"cause" api:"nullable"`
	ConsentSettings ConsentSettingUpdateResponseConsentSettings `json:"consentSettings" api:"nullable"`
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

type ConsentSettingUpdateResponseConsentSettings struct {
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
func (r ConsentSettingUpdateResponseConsentSettings) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingUpdateResponseConsentSettings) UnmarshalJSON(data []byte) error {
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
	SkipBlockingClassNames []string                         `json:"skipBlockingClassNames,omitzero"`
	WhitelistDomains       []string                         `json:"whitelistDomains,omitzero"`
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
	Mode                     string                                                 `json:"mode,omitzero" api:"required"`
	Translations             []ConsentSettingUpdateParamsDefaultTranslation         `json:"translations,omitzero" api:"required"`
	AutoShow                 param.Opt[bool]                                        `json:"autoShow,omitzero"`
	DisablePageInteraction   param.Opt[bool]                                        `json:"disablePageInteraction,omitzero"`
	HideFromBots             param.Opt[bool]                                        `json:"hideFromBots,omitzero"`
	ShowVendorsInPreferences param.Opt[bool]                                        `json:"showVendorsInPreferences,omitzero"`
	AutoShowDismissConfig    ConsentSettingUpdateParamsDefaultAutoShowDismissConfig `json:"autoShowDismissConfig,omitzero"`
	// Any of "after_pages", "after_seconds", "never".
	AutoShowDismissMode string                                      `json:"autoShowDismissMode,omitzero"`
	GuiOptions          ConsentSettingUpdateParamsDefaultGuiOptions `json:"guiOptions,omitzero"`
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
	ConsentModal     ConsentSettingUpdateParamsDefaultTranslationValueConsentModal     `json:"consentModal,omitzero"`
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

// The properties AcceptAllBtn, Description, RejectAllBtn, ShowPreferencesBtn,
// Title are required.
type ConsentSettingUpdateParamsDefaultTranslationValueConsentModal struct {
	AcceptAllBtn        string            `json:"acceptAllBtn" api:"required"`
	Description         string            `json:"description" api:"required"`
	RejectAllBtn        string            `json:"rejectAllBtn" api:"required"`
	ShowPreferencesBtn  string            `json:"showPreferencesBtn" api:"required"`
	Title               string            `json:"title" api:"required"`
	CloseIconLabel      param.Opt[string] `json:"closeIconLabel,omitzero"`
	Footer              param.Opt[string] `json:"footer,omitzero"`
	GpcNotification     param.Opt[string] `json:"gpcNotification,omitzero"`
	PrivacyPolicyLabel  param.Opt[string] `json:"privacyPolicyLabel,omitzero"`
	PrivacyPolicyURL    param.Opt[string] `json:"privacyPolicyUrl,omitzero"`
	RevisionMessage     param.Opt[string] `json:"revisionMessage,omitzero"`
	TermsOfServiceLabel param.Opt[string] `json:"termsOfServiceLabel,omitzero"`
	TermsOfServiceURL   param.Opt[string] `json:"termsOfServiceUrl,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsDefaultTranslationValueConsentModal) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsDefaultTranslationValueConsentModal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsDefaultTranslationValueConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AcceptAllBtn, CloseIconLabel, RejectAllBtn, SavePreferencesBtn,
// Sections, Title are required.
type ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModal struct {
	AcceptAllBtn        string                                                                     `json:"acceptAllBtn" api:"required"`
	CloseIconLabel      string                                                                     `json:"closeIconLabel" api:"required"`
	RejectAllBtn        string                                                                     `json:"rejectAllBtn" api:"required"`
	SavePreferencesBtn  string                                                                     `json:"savePreferencesBtn" api:"required"`
	Sections            []ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSection `json:"sections,omitzero" api:"required"`
	Title               string                                                                     `json:"title" api:"required"`
	ServiceCounterLabel param.Opt[string]                                                          `json:"serviceCounterLabel,omitzero"`
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
	Description    string                                                                              `json:"description" api:"required"`
	Title          string                                                                              `json:"title" api:"required"`
	LinkedCategory param.Opt[string]                                                                   `json:"linkedCategory,omitzero"`
	CookieTable    ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSectionCookieTable `json:"cookieTable,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSection) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Body, Headers are required.
type ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSectionCookieTable struct {
	Body    []ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSectionCookieTableBody   `json:"body,omitzero" api:"required"`
	Headers []ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSectionCookieTableHeader `json:"headers,omitzero" api:"required"`
	Caption param.Opt[string]                                                                           `json:"caption,omitzero"`
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

type ConsentSettingUpdateParamsDefaultAutoShowDismissConfig struct {
	PageCount param.Opt[int64] `json:"pageCount,omitzero"`
	Seconds   param.Opt[int64] `json:"seconds,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsDefaultAutoShowDismissConfig) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsDefaultAutoShowDismissConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsDefaultAutoShowDismissConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingUpdateParamsDefaultGuiOptions struct {
	ConsentModal     ConsentSettingUpdateParamsDefaultGuiOptionsConsentModal     `json:"consentModal,omitzero"`
	CssVariables     ConsentSettingUpdateParamsDefaultGuiOptionsCssVariables     `json:"cssVariables,omitzero"`
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

type ConsentSettingUpdateParamsDefaultGuiOptionsConsentModal struct {
	EqualWeightButtons param.Opt[bool] `json:"equalWeightButtons,omitzero"`
	FlipButtons        param.Opt[bool] `json:"flipButtons,omitzero"`
	ShowCloseIcon      param.Opt[bool] `json:"showCloseIcon,omitzero"`
	// Any of "AcceptAllNecessaryPreferences", "AcceptAllPreferences", "AcceptOnly",
	// "InformationOnly", "PreferencesOnly".
	ButtonLayout string `json:"buttonLayout,omitzero"`
	// Any of "bar", "bar_inline", "box", "box_inline", "box_wide", "cloud",
	// "cloud_inline".
	Layout string `json:"layout,omitzero"`
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

type ConsentSettingUpdateParamsDefaultGuiOptionsCssVariables struct {
	ButtonBorderRadius        param.Opt[string] `json:"buttonBorderRadius,omitzero"`
	FooterBg                  param.Opt[string] `json:"footerBg,omitzero"`
	FooterColor               param.Opt[string] `json:"footerColor,omitzero"`
	FooterLinkColor           param.Opt[string] `json:"footerLinkColor,omitzero"`
	FooterLinkHoverColor      param.Opt[string] `json:"footerLinkHoverColor,omitzero"`
	ModalBg                   param.Opt[string] `json:"modalBg,omitzero"`
	ModalBorderRadius         param.Opt[string] `json:"modalBorderRadius,omitzero"`
	PrimaryButtonBg           param.Opt[string] `json:"primaryButtonBg,omitzero"`
	PrimaryButtonColor        param.Opt[string] `json:"primaryButtonColor,omitzero"`
	PrimaryButtonHoverBg      param.Opt[string] `json:"primaryButtonHoverBg,omitzero"`
	PrimaryButtonHoverColor   param.Opt[string] `json:"primaryButtonHoverColor,omitzero"`
	PrimaryTextColor          param.Opt[string] `json:"primaryTextColor,omitzero"`
	SecondaryButtonBg         param.Opt[string] `json:"secondaryButtonBg,omitzero"`
	SecondaryButtonColor      param.Opt[string] `json:"secondaryButtonColor,omitzero"`
	SecondaryButtonHoverBg    param.Opt[string] `json:"secondaryButtonHoverBg,omitzero"`
	SecondaryButtonHoverColor param.Opt[string] `json:"secondaryButtonHoverColor,omitzero"`
	SecondaryTextColor        param.Opt[string] `json:"secondaryTextColor,omitzero"`
	ToggleOnBg                param.Opt[string] `json:"toggleOnBg,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsDefaultGuiOptionsCssVariables) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsDefaultGuiOptionsCssVariables
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsDefaultGuiOptionsCssVariables) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingUpdateParamsDefaultGuiOptionsPreferencesModal struct {
	EqualWeightButtons param.Opt[bool] `json:"equalWeightButtons,omitzero"`
	FlipButtons        param.Opt[bool] `json:"flipButtons,omitzero"`
	// Any of "AcceptAllRejectAllSave", "AcceptAllSave".
	ButtonLayout string `json:"buttonLayout,omitzero"`
	// Any of "bar", "bar_wide", "box".
	Layout string `json:"layout,omitzero"`
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
	RegionCode        string                               `json:"regionCode" api:"required"`
	Rule              ConsentSettingUpdateParamsRegionRule `json:"rule,omitzero" api:"required"`
	AdditionalRegions []string                             `json:"additionalRegions,omitzero"`
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
	Mode                     string                                                    `json:"mode,omitzero" api:"required"`
	Translations             []ConsentSettingUpdateParamsRegionRuleTranslation         `json:"translations,omitzero" api:"required"`
	AutoShow                 param.Opt[bool]                                           `json:"autoShow,omitzero"`
	DisablePageInteraction   param.Opt[bool]                                           `json:"disablePageInteraction,omitzero"`
	HideFromBots             param.Opt[bool]                                           `json:"hideFromBots,omitzero"`
	ShowVendorsInPreferences param.Opt[bool]                                           `json:"showVendorsInPreferences,omitzero"`
	AutoShowDismissConfig    ConsentSettingUpdateParamsRegionRuleAutoShowDismissConfig `json:"autoShowDismissConfig,omitzero"`
	// Any of "after_pages", "after_seconds", "never".
	AutoShowDismissMode string                                         `json:"autoShowDismissMode,omitzero"`
	GuiOptions          ConsentSettingUpdateParamsRegionRuleGuiOptions `json:"guiOptions,omitzero"`
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
	ConsentModal     ConsentSettingUpdateParamsRegionRuleTranslationValueConsentModal     `json:"consentModal,omitzero"`
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

// The properties AcceptAllBtn, Description, RejectAllBtn, ShowPreferencesBtn,
// Title are required.
type ConsentSettingUpdateParamsRegionRuleTranslationValueConsentModal struct {
	AcceptAllBtn        string            `json:"acceptAllBtn" api:"required"`
	Description         string            `json:"description" api:"required"`
	RejectAllBtn        string            `json:"rejectAllBtn" api:"required"`
	ShowPreferencesBtn  string            `json:"showPreferencesBtn" api:"required"`
	Title               string            `json:"title" api:"required"`
	CloseIconLabel      param.Opt[string] `json:"closeIconLabel,omitzero"`
	Footer              param.Opt[string] `json:"footer,omitzero"`
	GpcNotification     param.Opt[string] `json:"gpcNotification,omitzero"`
	PrivacyPolicyLabel  param.Opt[string] `json:"privacyPolicyLabel,omitzero"`
	PrivacyPolicyURL    param.Opt[string] `json:"privacyPolicyUrl,omitzero"`
	RevisionMessage     param.Opt[string] `json:"revisionMessage,omitzero"`
	TermsOfServiceLabel param.Opt[string] `json:"termsOfServiceLabel,omitzero"`
	TermsOfServiceURL   param.Opt[string] `json:"termsOfServiceUrl,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsRegionRuleTranslationValueConsentModal) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsRegionRuleTranslationValueConsentModal
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsRegionRuleTranslationValueConsentModal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AcceptAllBtn, CloseIconLabel, RejectAllBtn, SavePreferencesBtn,
// Sections, Title are required.
type ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModal struct {
	AcceptAllBtn        string                                                                        `json:"acceptAllBtn" api:"required"`
	CloseIconLabel      string                                                                        `json:"closeIconLabel" api:"required"`
	RejectAllBtn        string                                                                        `json:"rejectAllBtn" api:"required"`
	SavePreferencesBtn  string                                                                        `json:"savePreferencesBtn" api:"required"`
	Sections            []ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSection `json:"sections,omitzero" api:"required"`
	Title               string                                                                        `json:"title" api:"required"`
	ServiceCounterLabel param.Opt[string]                                                             `json:"serviceCounterLabel,omitzero"`
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
	Description    string                                                                                 `json:"description" api:"required"`
	Title          string                                                                                 `json:"title" api:"required"`
	LinkedCategory param.Opt[string]                                                                      `json:"linkedCategory,omitzero"`
	CookieTable    ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSectionCookieTable `json:"cookieTable,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSection) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSection
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Body, Headers are required.
type ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSectionCookieTable struct {
	Body    []ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSectionCookieTableBody   `json:"body,omitzero" api:"required"`
	Headers []ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader `json:"headers,omitzero" api:"required"`
	Caption param.Opt[string]                                                                              `json:"caption,omitzero"`
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

type ConsentSettingUpdateParamsRegionRuleAutoShowDismissConfig struct {
	PageCount param.Opt[int64] `json:"pageCount,omitzero"`
	Seconds   param.Opt[int64] `json:"seconds,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsRegionRuleAutoShowDismissConfig) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsRegionRuleAutoShowDismissConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsRegionRuleAutoShowDismissConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingUpdateParamsRegionRuleGuiOptions struct {
	ConsentModal     ConsentSettingUpdateParamsRegionRuleGuiOptionsConsentModal     `json:"consentModal,omitzero"`
	CssVariables     ConsentSettingUpdateParamsRegionRuleGuiOptionsCssVariables     `json:"cssVariables,omitzero"`
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

type ConsentSettingUpdateParamsRegionRuleGuiOptionsConsentModal struct {
	EqualWeightButtons param.Opt[bool] `json:"equalWeightButtons,omitzero"`
	FlipButtons        param.Opt[bool] `json:"flipButtons,omitzero"`
	ShowCloseIcon      param.Opt[bool] `json:"showCloseIcon,omitzero"`
	// Any of "AcceptAllNecessaryPreferences", "AcceptAllPreferences", "AcceptOnly",
	// "InformationOnly", "PreferencesOnly".
	ButtonLayout string `json:"buttonLayout,omitzero"`
	// Any of "bar", "bar_inline", "box", "box_inline", "box_wide", "cloud",
	// "cloud_inline".
	Layout string `json:"layout,omitzero"`
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

type ConsentSettingUpdateParamsRegionRuleGuiOptionsCssVariables struct {
	ButtonBorderRadius        param.Opt[string] `json:"buttonBorderRadius,omitzero"`
	FooterBg                  param.Opt[string] `json:"footerBg,omitzero"`
	FooterColor               param.Opt[string] `json:"footerColor,omitzero"`
	FooterLinkColor           param.Opt[string] `json:"footerLinkColor,omitzero"`
	FooterLinkHoverColor      param.Opt[string] `json:"footerLinkHoverColor,omitzero"`
	ModalBg                   param.Opt[string] `json:"modalBg,omitzero"`
	ModalBorderRadius         param.Opt[string] `json:"modalBorderRadius,omitzero"`
	PrimaryButtonBg           param.Opt[string] `json:"primaryButtonBg,omitzero"`
	PrimaryButtonColor        param.Opt[string] `json:"primaryButtonColor,omitzero"`
	PrimaryButtonHoverBg      param.Opt[string] `json:"primaryButtonHoverBg,omitzero"`
	PrimaryButtonHoverColor   param.Opt[string] `json:"primaryButtonHoverColor,omitzero"`
	PrimaryTextColor          param.Opt[string] `json:"primaryTextColor,omitzero"`
	SecondaryButtonBg         param.Opt[string] `json:"secondaryButtonBg,omitzero"`
	SecondaryButtonColor      param.Opt[string] `json:"secondaryButtonColor,omitzero"`
	SecondaryButtonHoverBg    param.Opt[string] `json:"secondaryButtonHoverBg,omitzero"`
	SecondaryButtonHoverColor param.Opt[string] `json:"secondaryButtonHoverColor,omitzero"`
	SecondaryTextColor        param.Opt[string] `json:"secondaryTextColor,omitzero"`
	ToggleOnBg                param.Opt[string] `json:"toggleOnBg,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParamsRegionRuleGuiOptionsCssVariables) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParamsRegionRuleGuiOptionsCssVariables
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParamsRegionRuleGuiOptionsCssVariables) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingUpdateParamsRegionRuleGuiOptionsPreferencesModal struct {
	EqualWeightButtons param.Opt[bool] `json:"equalWeightButtons,omitzero"`
	FlipButtons        param.Opt[bool] `json:"flipButtons,omitzero"`
	// Any of "AcceptAllRejectAllSave", "AcceptAllSave".
	ButtonLayout string `json:"buttonLayout,omitzero"`
	// Any of "bar", "bar_wide", "box".
	Layout string `json:"layout,omitzero"`
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
	InternalNotes        string            `json:"internalNotes" api:"required"`
	Label                string            `json:"label" api:"required"`
	Category             param.Opt[string] `json:"category,omitzero"`
	AdditionalCategories []string          `json:"additionalCategories,omitzero"`
	DomainPatterns       []string          `json:"domainPatterns,omitzero"`
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
