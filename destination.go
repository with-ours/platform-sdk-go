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

// DestinationService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDestinationService] method instead.
type DestinationService struct {
	Options []option.RequestOption
}

// NewDestinationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDestinationService(opts ...option.RequestOption) (r DestinationService) {
	r = DestinationService{}
	r.Options = opts
	return
}

// Create a new destination. Requires scope: destination:create
func (r *DestinationService) New(ctx context.Context, body DestinationNewParams, opts ...option.RequestOption) (res *DestinationNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/destinations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Find a single destination by ID. Requires scope: destination:find
func (r *DestinationService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *DestinationGetResponse, err error) {
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
func (r *DestinationService) Update(ctx context.Context, id string, body DestinationUpdateParams, opts ...option.RequestOption) (res *DestinationUpdateResponse, err error) {
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
func (r *DestinationService) List(ctx context.Context, opts ...option.RequestOption) (res *DestinationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/destinations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a destination. Requires scope: destination:delete
func (r *DestinationService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *DestinationDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("rest/v1/destinations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type DestinationNewResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DestinationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *DestinationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationGetResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DestinationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DestinationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationUpdateResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DestinationUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *DestinationUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationListResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DestinationListResponse) RawJSON() string { return r.JSON.raw }
func (r *DestinationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationDeleteResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DestinationDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *DestinationDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationNewParams struct {
	paramObj
}

func (r DestinationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DestinationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DestinationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationUpdateParams struct {
	paramObj
}

func (r DestinationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DestinationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DestinationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
