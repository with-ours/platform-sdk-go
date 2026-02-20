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
func (r *ConsentSettingService) New(ctx context.Context, body ConsentSettingNewParams, opts ...option.RequestOption) (res *ConsentSettingNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/consent-settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Find a single consent setting by ID. Requires scope: consentSettings:find
func (r *ConsentSettingService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ConsentSettingGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("rest/v1/consent-settings/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a consent setting. Requires scope: consentSettings:update
func (r *ConsentSettingService) Update(ctx context.Context, id string, body ConsentSettingUpdateParams, opts ...option.RequestOption) (res *ConsentSettingUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("rest/v1/consent-settings/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all consent settings. Requires scope: consentSettings:list
func (r *ConsentSettingService) List(ctx context.Context, opts ...option.RequestOption) (res *ConsentSettingListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/consent-settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a consent setting. Requires scope: consentSettings:delete
func (r *ConsentSettingService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *ConsentSettingDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("rest/v1/consent-settings/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type ConsentSettingNewResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingGetResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingUpdateResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingListResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingListResponse) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingDeleteResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConsentSettingDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ConsentSettingDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingNewParams struct {
	paramObj
}

func (r ConsentSettingNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConsentSettingUpdateParams struct {
	paramObj
}

func (r ConsentSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ConsentSettingUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConsentSettingUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
