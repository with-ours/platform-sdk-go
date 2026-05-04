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
func (r *GlobalDispatchCenterService) New(ctx context.Context, body GlobalDispatchCenterNewParams, opts ...option.RequestOption) (res *GlobalDispatchCenterNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/global-dispatch-centers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Find a single global dispatch center by ID. Requires scope: globalDispatch:find
func (r *GlobalDispatchCenterService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *GlobalDispatchCenterGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/global-dispatch-centers/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Partially update a global dispatch center. Only the fields you send are changed.
// Requires scope: globalDispatch:update
func (r *GlobalDispatchCenterService) Update(ctx context.Context, id string, body GlobalDispatchCenterUpdateParams, opts ...option.RequestOption) (res *GlobalDispatchCenterUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/global-dispatch-centers/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List all global dispatch centers. Requires scope: globalDispatch:list
func (r *GlobalDispatchCenterService) List(ctx context.Context, opts ...option.RequestOption) (res *GlobalDispatchCenterListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/global-dispatch-centers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete a global dispatch center. Requires scope: globalDispatch:delete
func (r *GlobalDispatchCenterService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *GlobalDispatchCenterDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/global-dispatch-centers/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type GlobalDispatchCenterNewResponse struct {
	// Server-assigned UUID for this dispatch center.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp when the center was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// When false, the dispatch center is configured but does not route events.
	IsEnabled bool `json:"isEnabled" api:"required"`
	// Discriminator for the entity type. Always "globalDispatchCenter".
	Kind string `json:"kind" api:"required"`
	// Routing categories in priority order (1..N).
	Categories []GlobalDispatchCenterNewResponseCategory `json:"categories" api:"nullable"`
	// Human-readable name shown in the dashboard.
	Name string `json:"name" api:"nullable"`
	// Free-form notes for this center.
	Notes string `json:"notes" api:"nullable"`
	// ISO 8601 timestamp of the last write. Equal to createdAt on a freshly created
	// center; advances on every PATCH.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		IsEnabled   respjson.Field
		Kind        respjson.Field
		Categories  respjson.Field
		Name        respjson.Field
		Notes       respjson.Field
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

type GlobalDispatchCenterNewResponseCategory struct {
	// Display name for the category.
	Name string `json:"name" api:"required"`
	// 1-indexed sort position. Always equals (sorted index + 1) — see PATCH for
	// details.
	Priority int64 `json:"priority" api:"required"`
	// Optional human-readable description.
	Description string `json:"description" api:"nullable"`
	// Destinations that receive events matching this category.
	DestinationIDs []string `json:"destinationIds" api:"nullable"`
	// Optional condition tree gating which events match this category.
	Logic any `json:"logic" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name           respjson.Field
		Priority       respjson.Field
		Description    respjson.Field
		DestinationIDs respjson.Field
		Logic          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalDispatchCenterNewResponseCategory) RawJSON() string { return r.JSON.raw }
func (r *GlobalDispatchCenterNewResponseCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalDispatchCenterGetResponse struct {
	// Server-assigned UUID for this dispatch center.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp when the center was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// When false, the dispatch center is configured but does not route events.
	IsEnabled bool `json:"isEnabled" api:"required"`
	// Discriminator for the entity type. Always "globalDispatchCenter".
	Kind string `json:"kind" api:"required"`
	// Routing categories in priority order (1..N).
	Categories []GlobalDispatchCenterGetResponseCategory `json:"categories" api:"nullable"`
	// Human-readable name shown in the dashboard.
	Name string `json:"name" api:"nullable"`
	// Free-form notes for this center.
	Notes string `json:"notes" api:"nullable"`
	// ISO 8601 timestamp of the last write. Equal to createdAt on a freshly created
	// center; advances on every PATCH.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		IsEnabled   respjson.Field
		Kind        respjson.Field
		Categories  respjson.Field
		Name        respjson.Field
		Notes       respjson.Field
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

type GlobalDispatchCenterGetResponseCategory struct {
	// Display name for the category.
	Name string `json:"name" api:"required"`
	// 1-indexed sort position. Always equals (sorted index + 1) — see PATCH for
	// details.
	Priority int64 `json:"priority" api:"required"`
	// Optional human-readable description.
	Description string `json:"description" api:"nullable"`
	// Destinations that receive events matching this category.
	DestinationIDs []string `json:"destinationIds" api:"nullable"`
	// Optional condition tree gating which events match this category.
	Logic any `json:"logic" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name           respjson.Field
		Priority       respjson.Field
		Description    respjson.Field
		DestinationIDs respjson.Field
		Logic          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalDispatchCenterGetResponseCategory) RawJSON() string { return r.JSON.raw }
func (r *GlobalDispatchCenterGetResponseCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalDispatchCenterUpdateResponse struct {
	// Server-assigned UUID for this dispatch center.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp when the center was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// When false, the dispatch center is configured but does not route events.
	IsEnabled bool `json:"isEnabled" api:"required"`
	// Discriminator for the entity type. Always "globalDispatchCenter".
	Kind string `json:"kind" api:"required"`
	// Routing categories in priority order (1..N).
	Categories []GlobalDispatchCenterUpdateResponseCategory `json:"categories" api:"nullable"`
	// Human-readable name shown in the dashboard.
	Name string `json:"name" api:"nullable"`
	// Free-form notes for this center.
	Notes string `json:"notes" api:"nullable"`
	// ISO 8601 timestamp of the last write. Equal to createdAt on a freshly created
	// center; advances on every PATCH.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		IsEnabled   respjson.Field
		Kind        respjson.Field
		Categories  respjson.Field
		Name        respjson.Field
		Notes       respjson.Field
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

type GlobalDispatchCenterUpdateResponseCategory struct {
	// Display name for the category.
	Name string `json:"name" api:"required"`
	// 1-indexed sort position. Always equals (sorted index + 1) — see PATCH for
	// details.
	Priority int64 `json:"priority" api:"required"`
	// Optional human-readable description.
	Description string `json:"description" api:"nullable"`
	// Destinations that receive events matching this category.
	DestinationIDs []string `json:"destinationIds" api:"nullable"`
	// Optional condition tree gating which events match this category.
	Logic any `json:"logic" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name           respjson.Field
		Priority       respjson.Field
		Description    respjson.Field
		DestinationIDs respjson.Field
		Logic          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalDispatchCenterUpdateResponseCategory) RawJSON() string { return r.JSON.raw }
func (r *GlobalDispatchCenterUpdateResponseCategory) UnmarshalJSON(data []byte) error {
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
	// Server-assigned UUID for this dispatch center.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp when the center was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// When false, the dispatch center is configured but does not route events.
	IsEnabled bool `json:"isEnabled" api:"required"`
	// Discriminator for the entity type. Always "globalDispatchCenter".
	Kind string `json:"kind" api:"required"`
	// Routing categories in priority order (1..N).
	Categories []GlobalDispatchCenterListResponseEntityCategory `json:"categories" api:"nullable"`
	// Human-readable name shown in the dashboard.
	Name string `json:"name" api:"nullable"`
	// Free-form notes for this center.
	Notes string `json:"notes" api:"nullable"`
	// ISO 8601 timestamp of the last write. Equal to createdAt on a freshly created
	// center; advances on every PATCH.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		IsEnabled   respjson.Field
		Kind        respjson.Field
		Categories  respjson.Field
		Name        respjson.Field
		Notes       respjson.Field
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

type GlobalDispatchCenterListResponseEntityCategory struct {
	// Display name for the category.
	Name string `json:"name" api:"required"`
	// 1-indexed sort position. Always equals (sorted index + 1) — see PATCH for
	// details.
	Priority int64 `json:"priority" api:"required"`
	// Optional human-readable description.
	Description string `json:"description" api:"nullable"`
	// Destinations that receive events matching this category.
	DestinationIDs []string `json:"destinationIds" api:"nullable"`
	// Optional condition tree gating which events match this category.
	Logic any `json:"logic" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name           respjson.Field
		Priority       respjson.Field
		Description    respjson.Field
		DestinationIDs respjson.Field
		Logic          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalDispatchCenterListResponseEntityCategory) RawJSON() string { return r.JSON.raw }
func (r *GlobalDispatchCenterListResponseEntityCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalDispatchCenterDeleteResponse struct {
	// The id of the dispatch center that was deleted.
	ID string `json:"id" api:"required"`
	// True when the underlying mutation succeeded; the entity is gone.
	Deleted bool `json:"deleted" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Deleted     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalDispatchCenterDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalDispatchCenterDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalDispatchCenterNewParams struct {
	// Whether the center starts enabled. Defaults to false — opt in by setting true
	// here or via PATCH later.
	IsEnabled param.Opt[bool] `json:"isEnabled,omitzero"`
	// Display name for the new center. Defaults to "Consent Dispatch Center".
	Name param.Opt[string] `json:"name,omitzero"`
	// Free-form notes shown in the dashboard. Not used for routing.
	Notes param.Opt[string] `json:"notes,omitzero"`
	paramObj
}

func (r GlobalDispatchCenterNewParams) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalDispatchCenterUpdateParams struct {
	// Toggle the dispatch center on/off without changing any other config.
	IsEnabled param.Opt[bool] `json:"isEnabled,omitzero"`
	// New display name for the center.
	Name param.Opt[string] `json:"name,omitzero"`
	// Replace the free-form notes.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Full replacement of the categories list. Categories are sorted by priority on
	// write and re-stamped 1..N — see the priority field. Omit to leave existing
	// categories untouched.
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
	// Optional human-readable description shown in the dashboard.
	Description param.Opt[string] `json:"description,omitzero"`
	// Display name for the category. Auto-generated if omitted.
	Name param.Opt[string] `json:"name,omitzero"`
	// Used as a sort key on write. The server sorts categories by this value
	// (ascending), then re-stamps priority as (sorted index + 1) on persist. Send any
	// positive number — gaps are ironed out, duplicate values keep input order via
	// stable sort. Omit to fall to the end.
	Priority param.Opt[float64] `json:"priority,omitzero"`
	// Destinations that receive events matching this category. Stale IDs (deleted
	// destinations or ones from another account) are silently filtered out at write
	// time.
	DestinationIDs []string `json:"destinationIds,omitzero"`
	// Optional condition tree. When set, only matching events route to this category.
	Logic any `json:"logic,omitzero"`
	paramObj
}

func (r GlobalDispatchCenterUpdateParamsCategory) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterUpdateParamsCategory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterUpdateParamsCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
