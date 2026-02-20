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

// RestV1AllowedEventService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRestV1AllowedEventService] method instead.
type RestV1AllowedEventService struct {
	Options []option.RequestOption
}

// NewRestV1AllowedEventService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRestV1AllowedEventService(opts ...option.RequestOption) (r RestV1AllowedEventService) {
	r = RestV1AllowedEventService{}
	r.Options = opts
	return
}

// Create a new allowed event. Requires scope: allowedEvent:create
func (r *RestV1AllowedEventService) New(ctx context.Context, body RestV1AllowedEventNewParams, opts ...option.RequestOption) (res *RestV1AllowedEventNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/allowed-events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Find a single allowed event by ID. Requires scope: allowedEvent:find
func (r *RestV1AllowedEventService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *RestV1AllowedEventGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("rest/v1/allowed-events/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all allowed events. Requires scope: allowedEvent:list
func (r *RestV1AllowedEventService) List(ctx context.Context, opts ...option.RequestOption) (res *RestV1AllowedEventListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/allowed-events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a allowed event. Requires scope: allowedEvent:delete
func (r *RestV1AllowedEventService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *RestV1AllowedEventDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("rest/v1/allowed-events/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type RestV1AllowedEventNewResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1AllowedEventNewResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1AllowedEventNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1AllowedEventGetResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1AllowedEventGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1AllowedEventGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1AllowedEventListResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1AllowedEventListResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1AllowedEventListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1AllowedEventDeleteResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RestV1AllowedEventDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *RestV1AllowedEventDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RestV1AllowedEventNewParams struct {
	paramObj
}

func (r RestV1AllowedEventNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RestV1AllowedEventNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RestV1AllowedEventNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
