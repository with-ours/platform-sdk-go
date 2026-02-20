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

// AllowedEventService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAllowedEventService] method instead.
type AllowedEventService struct {
	Options []option.RequestOption
}

// NewAllowedEventService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAllowedEventService(opts ...option.RequestOption) (r AllowedEventService) {
	r = AllowedEventService{}
	r.Options = opts
	return
}

// Create a new allowed event. Requires scope: allowedEvent:create
func (r *AllowedEventService) New(ctx context.Context, body AllowedEventNewParams, opts ...option.RequestOption) (res *AllowedEventNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/allowed-events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Find a single allowed event by ID. Requires scope: allowedEvent:find
func (r *AllowedEventService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AllowedEventGetResponse, err error) {
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
func (r *AllowedEventService) List(ctx context.Context, opts ...option.RequestOption) (res *AllowedEventListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/allowed-events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a allowed event. Requires scope: allowedEvent:delete
func (r *AllowedEventService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *AllowedEventDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("rest/v1/allowed-events/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AllowedEventNewResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AllowedEventNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AllowedEventNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AllowedEventGetResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AllowedEventGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AllowedEventGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AllowedEventListResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AllowedEventListResponse) RawJSON() string { return r.JSON.raw }
func (r *AllowedEventListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AllowedEventDeleteResponse struct {
	Data any `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AllowedEventDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *AllowedEventDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AllowedEventNewParams struct {
	paramObj
}

func (r AllowedEventNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AllowedEventNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AllowedEventNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
