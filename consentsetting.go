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
	Categories []any  `json:"categories,omitzero" api:"required"`
	Name       string `json:"name" api:"required"`
	Regions    []any  `json:"regions,omitzero" api:"required"`
	Services   []any  `json:"services,omitzero" api:"required"`
	// Any of "Disabled", "Enabled".
	Status                 ConsentSettingUpdateParamsStatus `json:"status,omitzero" api:"required"`
	ConsentCookieName      param.Opt[string]                `json:"consentCookieName,omitzero"`
	Revision               param.Opt[float64]               `json:"revision,omitzero"`
	WebSDKToken            param.Opt[string]                `json:"webSDKToken,omitzero"`
	SkipBlockingClassNames []string                         `json:"skipBlockingClassNames,omitzero"`
	WhitelistDomains       []any                            `json:"whitelistDomains,omitzero"`
	CustomDomain           any                              `json:"customDomain,omitzero"`
	Default                any                              `json:"default,omitzero"`
	paramObj
}

func (r ConsentSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingUpdateParamsStatus string

const (
	ConsentSettingUpdateParamsStatusDisabled ConsentSettingUpdateParamsStatus = "Disabled"
	ConsentSettingUpdateParamsStatusEnabled  ConsentSettingUpdateParamsStatus = "Enabled"
)
