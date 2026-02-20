// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacyplatform

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/ours-privacy-platform-go/internal/apijson"
	"github.com/stainless-sdks/ours-privacy-platform-go/internal/requestconfig"
	"github.com/stainless-sdks/ours-privacy-platform-go/option"
	"github.com/stainless-sdks/ours-privacy-platform-go/packages/param"
	"github.com/stainless-sdks/ours-privacy-platform-go/packages/respjson"
)

// RestV1ConsentSettingService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRestV1ConsentSettingService] method instead.
type RestV1ConsentSettingService struct {
	Options []option.RequestOption
}

// NewRestV1ConsentSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRestV1ConsentSettingService(opts ...option.RequestOption) (r RestV1ConsentSettingService) {
	r = RestV1ConsentSettingService{}
	r.Options = opts
	return
}

// Create a new consent setting. Requires scope: consentSettings:create
func (r *RestV1ConsentSettingService) New(ctx context.Context, body RestV1ConsentSettingNewParams, opts ...option.RequestOption) (res *RestV1ConsentSettingNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/consent-settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Find a single consent setting by ID. Requires scope: consentSettings:find
func (r *RestV1ConsentSettingService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *RestV1ConsentSettingGetResponse, err error) {
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
func (r *RestV1ConsentSettingService) Update(ctx context.Context, id string, body RestV1ConsentSettingUpdateParams, opts ...option.RequestOption) (res *RestV1ConsentSettingUpdateResponse, err error) {
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
func (r *RestV1ConsentSettingService) List(ctx context.Context, opts ...option.RequestOption) (res *RestV1ConsentSettingListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/consent-settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a consent setting. Requires scope: consentSettings:delete
func (r *RestV1ConsentSettingService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *RestV1ConsentSettingDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("rest/v1/consent-settings/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type RestV1ConsentSettingNewResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1ConsentSettingNewResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1ConsentSettingNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1ConsentSettingGetResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1ConsentSettingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1ConsentSettingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1ConsentSettingUpdateResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1ConsentSettingUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1ConsentSettingUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1ConsentSettingListResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1ConsentSettingListResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1ConsentSettingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1ConsentSettingDeleteResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1ConsentSettingDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1ConsentSettingDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1ConsentSettingNewParams struct {
	paramObj
}

func (r RestV1ConsentSettingNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RestV1ConsentSettingNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RestV1ConsentSettingNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1ConsentSettingUpdateParams struct {
	paramObj
}

func (r RestV1ConsentSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow RestV1ConsentSettingUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RestV1ConsentSettingUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
