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

// RestV1DestinationService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRestV1DestinationService] method instead.
type RestV1DestinationService struct {
	Options []option.RequestOption
}

// NewRestV1DestinationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRestV1DestinationService(opts ...option.RequestOption) (r RestV1DestinationService) {
	r = RestV1DestinationService{}
	r.Options = opts
	return
}

// Create a new destination. Requires scope: destination:create
func (r *RestV1DestinationService) New(ctx context.Context, body RestV1DestinationNewParams, opts ...option.RequestOption) (res *RestV1DestinationNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/destinations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Find a single destination by ID. Requires scope: destination:find
func (r *RestV1DestinationService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *RestV1DestinationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("rest/v1/destinations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a destination. Requires scope: destination:update
func (r *RestV1DestinationService) Update(ctx context.Context, id string, body RestV1DestinationUpdateParams, opts ...option.RequestOption) (res *RestV1DestinationUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("rest/v1/destinations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all destinations. Requires scope: destination:list
func (r *RestV1DestinationService) List(ctx context.Context, opts ...option.RequestOption) (res *RestV1DestinationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/destinations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a destination. Requires scope: destination:delete
func (r *RestV1DestinationService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *RestV1DestinationDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("rest/v1/destinations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type RestV1DestinationNewResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1DestinationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1DestinationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1DestinationGetResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1DestinationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1DestinationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1DestinationUpdateResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1DestinationUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1DestinationUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1DestinationListResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1DestinationListResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1DestinationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1DestinationDeleteResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1DestinationDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1DestinationDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1DestinationNewParams struct {
	paramObj
}

func (r RestV1DestinationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RestV1DestinationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RestV1DestinationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1DestinationUpdateParams struct {
	paramObj
}

func (r RestV1DestinationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow RestV1DestinationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RestV1DestinationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
