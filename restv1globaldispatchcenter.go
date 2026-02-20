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

// RestV1GlobalDispatchCenterService contains methods and other services that help
// with interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRestV1GlobalDispatchCenterService] method instead.
type RestV1GlobalDispatchCenterService struct {
	Options []option.RequestOption
}

// NewRestV1GlobalDispatchCenterService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRestV1GlobalDispatchCenterService(opts ...option.RequestOption) (r RestV1GlobalDispatchCenterService) {
	r = RestV1GlobalDispatchCenterService{}
	r.Options = opts
	return
}

// Create a new global dispatch center. Requires scope: globalDispatch:create
func (r *RestV1GlobalDispatchCenterService) New(ctx context.Context, body RestV1GlobalDispatchCenterNewParams, opts ...option.RequestOption) (res *RestV1GlobalDispatchCenterNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/global-dispatch-centers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Find a single global dispatch center by ID. Requires scope: globalDispatch:find
func (r *RestV1GlobalDispatchCenterService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *RestV1GlobalDispatchCenterGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("rest/v1/global-dispatch-centers/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a global dispatch center. Requires scope: globalDispatch:update
func (r *RestV1GlobalDispatchCenterService) Update(ctx context.Context, id string, body RestV1GlobalDispatchCenterUpdateParams, opts ...option.RequestOption) (res *RestV1GlobalDispatchCenterUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("rest/v1/global-dispatch-centers/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all global dispatch centers. Requires scope: globalDispatch:list
func (r *RestV1GlobalDispatchCenterService) List(ctx context.Context, opts ...option.RequestOption) (res *RestV1GlobalDispatchCenterListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/global-dispatch-centers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a global dispatch center. Requires scope: globalDispatch:delete
func (r *RestV1GlobalDispatchCenterService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *RestV1GlobalDispatchCenterDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("rest/v1/global-dispatch-centers/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type RestV1GlobalDispatchCenterNewResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1GlobalDispatchCenterNewResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1GlobalDispatchCenterNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1GlobalDispatchCenterGetResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1GlobalDispatchCenterGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1GlobalDispatchCenterGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1GlobalDispatchCenterUpdateResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1GlobalDispatchCenterUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1GlobalDispatchCenterUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1GlobalDispatchCenterListResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1GlobalDispatchCenterListResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1GlobalDispatchCenterListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1GlobalDispatchCenterDeleteResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1GlobalDispatchCenterDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1GlobalDispatchCenterDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1GlobalDispatchCenterNewParams struct {
	paramObj
}

func (r RestV1GlobalDispatchCenterNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RestV1GlobalDispatchCenterNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RestV1GlobalDispatchCenterNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1GlobalDispatchCenterUpdateParams struct {
	paramObj
}

func (r RestV1GlobalDispatchCenterUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow RestV1GlobalDispatchCenterUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RestV1GlobalDispatchCenterUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
