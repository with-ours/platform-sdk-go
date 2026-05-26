// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy

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

// TagManagerService contains methods and other services that help with interacting
// with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTagManagerService] method instead.
type TagManagerService struct {
	Options []option.RequestOption
}

// NewTagManagerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTagManagerService(opts ...option.RequestOption) (r TagManagerService) {
	r = TagManagerService{}
	r.Options = opts
	return
}

// List every tag manager on this account. Each tag manager is a pixel-scoped
// container of tags, triggers, variables, and folders. Not paginated — accounts
// are capped at a small number of tag managers in practice, so the response fits
// in a single page. Requires scope: tagManagers:list
func (r *TagManagerService) List(ctx context.Context, opts ...option.RequestOption) (res *TagManagerListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/tag-managers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Create a new tag manager. The server seeds three default triggers
// (`Initialization`, `PageView`, `DomReady`) and one `OursInitTag` so the
// container is immediately usable — call
// `GET /tag-manager-triggers?tagManagerId={id}` right after create to grab their
// server-assigned ids so you can reuse them in `fireTriggerIds` instead of
// redundantly creating a second `PageView`/`DomReady`/`Initialization`. Returns
// the bare entity. Accounts have a per-account tag manager limit — exceeding it
// returns 409 with the reason in the response `error` field. Requires scope:
// tagManagers:create
func (r *TagManagerService) New(ctx context.Context, body TagManagerNewParams, opts ...option.RequestOption) (res *TagManagerNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/tag-managers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Fetch a single tag manager by id, including its server-assigned `pixel` token
// used by the install snippet. Requires scope: tagManagers:find
func (r *TagManagerService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *TagManagerGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/tag-managers/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Partially update a tag manager. Only the fields you send are changed; omitted
// fields keep their current value. Send `dataLayerName: null` to clear the
// override and fall back to the SDK default. Requires scope: tagManagers:update
func (r *TagManagerService) Update(ctx context.Context, id string, body TagManagerUpdateParams, opts ...option.RequestOption) (res *TagManagerUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/tag-managers/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete a tag manager. Child tags, triggers, variables, and folders are
// cascade-deleted with the container. Requires scope: tagManagers:delete
func (r *TagManagerService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *TagManagerDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/tag-managers/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type TagManagerListResponse struct {
	Entities []TagManagerListResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagManagerListResponse) RawJSON() string { return r.JSON.raw }
func (r *TagManagerListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerListResponseEntity struct {
	// Server-assigned UUID for this tag manager container.
	ID string `json:"id" api:"required"`
	// Account that owns this tag manager.
	AccountID string `json:"accountId" api:"required"`
	// Human-readable name for the tag manager.
	Name string `json:"name" api:"required"`
	// Server-assigned pixel/container token. Used in the tag-manager install snippet
	// served on customer sites — do not regenerate via the API.
	Pixel string `json:"pixel" api:"required"`
	// ISO 8601 timestamp when the tag manager was created.
	CreatedAt string `json:"createdAt" api:"nullable"`
	// Window-global name of the customer data layer that triggers and variables read
	// from (e.g. `dataLayer`). Defaults to `null`, which means the SDK falls back to
	// its built-in name.
	DataLayerName string `json:"dataLayerName" api:"nullable"`
	// ISO 8601 timestamp of the last update.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		AccountID     respjson.Field
		Name          respjson.Field
		Pixel         respjson.Field
		CreatedAt     respjson.Field
		DataLayerName respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagManagerListResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *TagManagerListResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerNewResponse struct {
	// Server-assigned UUID for this tag manager container.
	ID string `json:"id" api:"required"`
	// Account that owns this tag manager.
	AccountID string `json:"accountId" api:"required"`
	// Human-readable name for the tag manager.
	Name string `json:"name" api:"required"`
	// Server-assigned pixel/container token. Used in the tag-manager install snippet
	// served on customer sites — do not regenerate via the API.
	Pixel string `json:"pixel" api:"required"`
	// ISO 8601 timestamp when the tag manager was created.
	CreatedAt string `json:"createdAt" api:"nullable"`
	// Window-global name of the customer data layer that triggers and variables read
	// from (e.g. `dataLayer`). Defaults to `null`, which means the SDK falls back to
	// its built-in name.
	DataLayerName string `json:"dataLayerName" api:"nullable"`
	// ISO 8601 timestamp of the last update.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		AccountID     respjson.Field
		Name          respjson.Field
		Pixel         respjson.Field
		CreatedAt     respjson.Field
		DataLayerName respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagManagerNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TagManagerNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerGetResponse struct {
	// Server-assigned UUID for this tag manager container.
	ID string `json:"id" api:"required"`
	// Account that owns this tag manager.
	AccountID string `json:"accountId" api:"required"`
	// Human-readable name for the tag manager.
	Name string `json:"name" api:"required"`
	// Server-assigned pixel/container token. Used in the tag-manager install snippet
	// served on customer sites — do not regenerate via the API.
	Pixel string `json:"pixel" api:"required"`
	// ISO 8601 timestamp when the tag manager was created.
	CreatedAt string `json:"createdAt" api:"nullable"`
	// Window-global name of the customer data layer that triggers and variables read
	// from (e.g. `dataLayer`). Defaults to `null`, which means the SDK falls back to
	// its built-in name.
	DataLayerName string `json:"dataLayerName" api:"nullable"`
	// ISO 8601 timestamp of the last update.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		AccountID     respjson.Field
		Name          respjson.Field
		Pixel         respjson.Field
		CreatedAt     respjson.Field
		DataLayerName respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagManagerGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TagManagerGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerUpdateResponse struct {
	// Server-assigned UUID for this tag manager container.
	ID string `json:"id" api:"required"`
	// Account that owns this tag manager.
	AccountID string `json:"accountId" api:"required"`
	// Human-readable name for the tag manager.
	Name string `json:"name" api:"required"`
	// Server-assigned pixel/container token. Used in the tag-manager install snippet
	// served on customer sites — do not regenerate via the API.
	Pixel string `json:"pixel" api:"required"`
	// ISO 8601 timestamp when the tag manager was created.
	CreatedAt string `json:"createdAt" api:"nullable"`
	// Window-global name of the customer data layer that triggers and variables read
	// from (e.g. `dataLayer`). Defaults to `null`, which means the SDK falls back to
	// its built-in name.
	DataLayerName string `json:"dataLayerName" api:"nullable"`
	// ISO 8601 timestamp of the last update.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		AccountID     respjson.Field
		Name          respjson.Field
		Pixel         respjson.Field
		CreatedAt     respjson.Field
		DataLayerName respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagManagerUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *TagManagerUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerDeleteResponse struct {
	// The id of the tag manager that was deleted.
	ID string `json:"id" api:"required"`
	// True when the underlying mutation succeeded.
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
func (r TagManagerDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *TagManagerDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerNewParams struct {
	// Human-readable name for the new tag manager.
	Name string `json:"name" api:"required"`
	// Optional global data-layer name (e.g. `dataLayer`). Omit to use the SDK default.
	DataLayerName param.Opt[string] `json:"dataLayerName,omitzero"`
	paramObj
}

func (r TagManagerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TagManagerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagManagerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerUpdateParams struct {
	// New data-layer name. Send `null` to clear and fall back to the SDK default.
	DataLayerName param.Opt[string] `json:"dataLayerName,omitzero"`
	// New display name.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r TagManagerUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TagManagerUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagManagerUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
