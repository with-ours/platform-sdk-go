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

// GlobalDispatchCenterService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGlobalDispatchCenterService] method instead.
type GlobalDispatchCenterService struct {
	Options []option.RequestOption
}

// NewGlobalDispatchCenterService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGlobalDispatchCenterService(opts ...option.RequestOption) (r GlobalDispatchCenterService) {
	r = GlobalDispatchCenterService{}
	r.Options = opts
	return
}

// Create a new global dispatch center. Requires scope: globalDispatch:create
func (r *GlobalDispatchCenterService) New(ctx context.Context, opts ...option.RequestOption) (res *GlobalDispatchCenterNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/global-dispatch-centers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Find a single global dispatch center by ID. Requires scope: globalDispatch:find
func (r *GlobalDispatchCenterService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *GlobalDispatchCenterGetResponse, err error) {
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
func (r *GlobalDispatchCenterService) Update(ctx context.Context, id string, body GlobalDispatchCenterUpdateParams, opts ...option.RequestOption) (res *GlobalDispatchCenterUpdateResponse, err error) {
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
func (r *GlobalDispatchCenterService) List(ctx context.Context, opts ...option.RequestOption) (res *GlobalDispatchCenterListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/global-dispatch-centers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a global dispatch center. Requires scope: globalDispatch:delete
func (r *GlobalDispatchCenterService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("rest/v1/global-dispatch-centers/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type GlobalDispatchCenterNewResponse struct {
	ID        string `json:"id" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	IsEnabled bool   `json:"isEnabled" api:"required"`
	Kind      string `json:"kind" api:"required"`
	Name      string `json:"name" api:"nullable"`
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		IsEnabled   respjson.Field
		Kind        respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalDispatchCenterNewResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalDispatchCenterNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalDispatchCenterGetResponse struct {
	ID        string `json:"id" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	IsEnabled bool   `json:"isEnabled" api:"required"`
	Kind      string `json:"kind" api:"required"`
	Name      string `json:"name" api:"nullable"`
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		IsEnabled   respjson.Field
		Kind        respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalDispatchCenterGetResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalDispatchCenterGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalDispatchCenterUpdateResponse struct {
	ID        string `json:"id" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	IsEnabled bool   `json:"isEnabled" api:"required"`
	Kind      string `json:"kind" api:"required"`
	Name      string `json:"name" api:"nullable"`
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		IsEnabled   respjson.Field
		Kind        respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalDispatchCenterUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalDispatchCenterUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalDispatchCenterListResponse struct {
	Entities []GlobalDispatchCenterListResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalDispatchCenterListResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalDispatchCenterListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalDispatchCenterListResponseEntity struct {
	ID        string `json:"id" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	IsEnabled bool   `json:"isEnabled" api:"required"`
	Kind      string `json:"kind" api:"required"`
	Name      string `json:"name" api:"nullable"`
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		IsEnabled   respjson.Field
		Kind        respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalDispatchCenterListResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *GlobalDispatchCenterListResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalDispatchCenterUpdateParams struct {
	IsEnabled  param.Opt[bool]                            `json:"isEnabled,omitzero"`
	Name       param.Opt[string]                          `json:"name,omitzero"`
	Notes      param.Opt[string]                          `json:"notes,omitzero"`
	Categories []GlobalDispatchCenterUpdateParamsCategory `json:"categories,omitzero"`
	paramObj
}

func (r GlobalDispatchCenterUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalDispatchCenterUpdateParamsCategory struct {
	Description    param.Opt[string]  `json:"description,omitzero"`
	Name           param.Opt[string]  `json:"name,omitzero"`
	Priority       param.Opt[float64] `json:"priority,omitzero"`
	DestinationIDs []string           `json:"destinationIds,omitzero"`
	Logic          any                `json:"logic,omitzero"`
	paramObj
}

func (r GlobalDispatchCenterUpdateParamsCategory) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterUpdateParamsCategory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterUpdateParamsCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
