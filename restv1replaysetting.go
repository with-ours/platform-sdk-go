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

// RestV1ReplaySettingService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRestV1ReplaySettingService] method instead.
type RestV1ReplaySettingService struct {
	Options []option.RequestOption
}

// NewRestV1ReplaySettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRestV1ReplaySettingService(opts ...option.RequestOption) (r RestV1ReplaySettingService) {
	r = RestV1ReplaySettingService{}
	r.Options = opts
	return
}

// Create a new replay setting. Requires scope: replaySettings:create
func (r *RestV1ReplaySettingService) New(ctx context.Context, body RestV1ReplaySettingNewParams, opts ...option.RequestOption) (res *RestV1ReplaySettingNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/replay-settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Find a single replay setting by ID. Requires scope: replaySettings:find
func (r *RestV1ReplaySettingService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *RestV1ReplaySettingGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("rest/v1/replay-settings/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a replay setting. Requires scope: replaySettings:update
func (r *RestV1ReplaySettingService) Update(ctx context.Context, id string, body RestV1ReplaySettingUpdateParams, opts ...option.RequestOption) (res *RestV1ReplaySettingUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("rest/v1/replay-settings/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all replay settings. Requires scope: replaySettings:list
func (r *RestV1ReplaySettingService) List(ctx context.Context, opts ...option.RequestOption) (res *RestV1ReplaySettingListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/replay-settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a replay setting. Requires scope: replaySettings:delete
func (r *RestV1ReplaySettingService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *RestV1ReplaySettingDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("rest/v1/replay-settings/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type RestV1ReplaySettingNewResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1ReplaySettingNewResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1ReplaySettingNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1ReplaySettingGetResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1ReplaySettingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1ReplaySettingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1ReplaySettingUpdateResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1ReplaySettingUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1ReplaySettingUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1ReplaySettingListResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1ReplaySettingListResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1ReplaySettingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1ReplaySettingDeleteResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1ReplaySettingDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1ReplaySettingDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1ReplaySettingNewParams struct {
	paramObj
}

func (r RestV1ReplaySettingNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RestV1ReplaySettingNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RestV1ReplaySettingNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1ReplaySettingUpdateParams struct {
	paramObj
}

func (r RestV1ReplaySettingUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow RestV1ReplaySettingUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RestV1ReplaySettingUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
