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

// RestV1SourceService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRestV1SourceService] method instead.
type RestV1SourceService struct {
	Options []option.RequestOption
}

// NewRestV1SourceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRestV1SourceService(opts ...option.RequestOption) (r RestV1SourceService) {
	r = RestV1SourceService{}
	r.Options = opts
	return
}

// Create a new source. Requires scope: source:create
func (r *RestV1SourceService) New(ctx context.Context, body RestV1SourceNewParams, opts ...option.RequestOption) (res *RestV1SourceNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/sources"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Find a single source by ID. Requires scope: source:view
func (r *RestV1SourceService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *RestV1SourceGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("rest/v1/sources/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a source. Requires scope: source:update
func (r *RestV1SourceService) Update(ctx context.Context, id string, body RestV1SourceUpdateParams, opts ...option.RequestOption) (res *RestV1SourceUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("rest/v1/sources/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all sources. Requires scope: source:list
func (r *RestV1SourceService) List(ctx context.Context, opts ...option.RequestOption) (res *RestV1SourceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/sources"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a source. Requires scope: source:delete
func (r *RestV1SourceService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *RestV1SourceDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("rest/v1/sources/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type RestV1SourceNewResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1SourceNewResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1SourceNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1SourceGetResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1SourceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1SourceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1SourceUpdateResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1SourceUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1SourceUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1SourceListResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1SourceListResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1SourceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1SourceDeleteResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1SourceDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1SourceDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1SourceNewParams struct {
	paramObj
}

func (r RestV1SourceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RestV1SourceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RestV1SourceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1SourceUpdateParams struct {
	paramObj
}

func (r RestV1SourceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow RestV1SourceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RestV1SourceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
