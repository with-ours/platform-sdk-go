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

// ReplaySettingService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReplaySettingService] method instead.
type ReplaySettingService struct {
	Options []option.RequestOption
}

// NewReplaySettingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewReplaySettingService(opts ...option.RequestOption) (r ReplaySettingService) {
	r = ReplaySettingService{}
	r.Options = opts
	return
}

// Create a new replay setting. Requires scope: replaySettings:create
func (r *ReplaySettingService) New(ctx context.Context, body ReplaySettingNewParams, opts ...option.RequestOption) (res *ReplaySettingNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/replay-settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Find a single replay setting by ID. Requires scope: replaySettings:find
func (r *ReplaySettingService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ReplaySettingGetResponse, err error) {
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
func (r *ReplaySettingService) Update(ctx context.Context, id string, body ReplaySettingUpdateParams, opts ...option.RequestOption) (res *ReplaySettingUpdateResponse, err error) {
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
func (r *ReplaySettingService) List(ctx context.Context, opts ...option.RequestOption) (res *ReplaySettingListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/replay-settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a replay setting. Requires scope: replaySettings:delete
func (r *ReplaySettingService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *ReplaySettingDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("rest/v1/replay-settings/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type ReplaySettingNewResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReplaySettingNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ReplaySettingNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReplaySettingGetResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReplaySettingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ReplaySettingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReplaySettingUpdateResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReplaySettingUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ReplaySettingUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReplaySettingListResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReplaySettingListResponse) RawJSON() string { return r.JSON.raw }
func (r *ReplaySettingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReplaySettingDeleteResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReplaySettingDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ReplaySettingDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReplaySettingNewParams struct {
	paramObj
}

func (r ReplaySettingNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ReplaySettingNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReplaySettingNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReplaySettingUpdateParams struct {
	paramObj
}

func (r ReplaySettingUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ReplaySettingUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReplaySettingUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
