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

// RestV1VersionService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRestV1VersionService] method instead.
type RestV1VersionService struct {
	Options []option.RequestOption
}

// NewRestV1VersionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRestV1VersionService(opts ...option.RequestOption) (r RestV1VersionService) {
	r = RestV1VersionService{}
	r.Options = opts
	return
}

// Create a new version. Requires scope: version:publish
func (r *RestV1VersionService) New(ctx context.Context, body RestV1VersionNewParams, opts ...option.RequestOption) (res *RestV1VersionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/versions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Find a single version by ID. Requires scope: version:find
func (r *RestV1VersionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *RestV1VersionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("rest/v1/versions/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a version. Requires scope: version:update
func (r *RestV1VersionService) Update(ctx context.Context, id string, body RestV1VersionUpdateParams, opts ...option.RequestOption) (res *RestV1VersionUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("rest/v1/versions/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all versions. Requires scope: version:list
func (r *RestV1VersionService) List(ctx context.Context, opts ...option.RequestOption) (res *RestV1VersionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/versions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type RestV1VersionNewResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1VersionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1VersionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1VersionGetResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1VersionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1VersionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1VersionUpdateResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1VersionUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1VersionUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1VersionListResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1VersionListResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1VersionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1VersionNewParams struct {
	paramObj
}

func (r RestV1VersionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RestV1VersionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RestV1VersionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1VersionUpdateParams struct {
	paramObj
}

func (r RestV1VersionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow RestV1VersionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RestV1VersionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
