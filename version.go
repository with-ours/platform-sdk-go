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

// VersionService contains methods and other services that help with interacting
// with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVersionService] method instead.
type VersionService struct {
	Options []option.RequestOption
}

// NewVersionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVersionService(opts ...option.RequestOption) (r VersionService) {
	r = VersionService{}
	r.Options = opts
	return
}

// Create a new version. Requires scope: version:publish
func (r *VersionService) New(ctx context.Context, body VersionNewParams, opts ...option.RequestOption) (res *VersionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/versions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Find a single version by ID. Requires scope: version:find
func (r *VersionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *VersionGetResponse, err error) {
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
func (r *VersionService) Update(ctx context.Context, id string, body VersionUpdateParams, opts ...option.RequestOption) (res *VersionUpdateResponse, err error) {
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
func (r *VersionService) List(ctx context.Context, opts ...option.RequestOption) (res *VersionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/versions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type VersionNewResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *VersionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionGetResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *VersionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionUpdateResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *VersionUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionListResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionListResponse) RawJSON() string { return r.JSON.raw }
func (r *VersionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionNewParams struct {
	paramObj
}

func (r VersionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VersionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VersionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionUpdateParams struct {
	paramObj
}

func (r VersionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow VersionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VersionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
