// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/with-ours/platform-sdk-go/v2/internal/apijson"
	"github.com/with-ours/platform-sdk-go/v2/internal/apiquery"
	"github.com/with-ours/platform-sdk-go/v2/internal/requestconfig"
	"github.com/with-ours/platform-sdk-go/v2/option"
	"github.com/with-ours/platform-sdk-go/v2/packages/pagination"
	"github.com/with-ours/platform-sdk-go/v2/packages/param"
	"github.com/with-ours/platform-sdk-go/v2/packages/respjson"
)

// TagManagerVariableService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTagManagerVariableService] method instead.
type TagManagerVariableService struct {
	Options []option.RequestOption
}

// NewTagManagerVariableService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTagManagerVariableService(opts ...option.RequestOption) (r TagManagerVariableService) {
	r = TagManagerVariableService{}
	r.Options = opts
	return
}

// List variables inside a single tag manager. Requires the `tagManagerId` query
// parameter — variables are always scoped to one parent container. Supports cursor
// pagination via `limit` and `cursor`; the limit clamp is 1000 so a single request
// can return the full set (the web-app workspace renders all variables in one
// shot). Requires API-key scope or current OAuth user permission: tagManagers:find
func (r *TagManagerVariableService) List(ctx context.Context, query TagManagerVariableListParams, opts ...option.RequestOption) (res *pagination.Cursor[TagManagerVariableListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "rest/v1/tag-manager-variables"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List variables inside a single tag manager. Requires the `tagManagerId` query
// parameter — variables are always scoped to one parent container. Supports cursor
// pagination via `limit` and `cursor`; the limit clamp is 1000 so a single request
// can return the full set (the web-app workspace renders all variables in one
// shot). Requires API-key scope or current OAuth user permission: tagManagers:find
func (r *TagManagerVariableService) ListAutoPaging(ctx context.Context, query TagManagerVariableListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[TagManagerVariableListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Create a new variable inside a tag manager. `tagManagerId` is required in the
// body. Known input failures (e.g. duplicate variable name within the tag manager)
// are returned as HTTP 409 with the reason in the response `error` field. Requires
// API-key scope or current OAuth user permission: tagManagers:update
func (r *TagManagerVariableService) New(ctx context.Context, body TagManagerVariableNewParams, opts ...option.RequestOption) (res *TagManagerVariableNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/tag-manager-variables"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Find a single tag manager variable by ID. Requires API-key scope or current
// OAuth user permission: tagManagers:find
func (r *TagManagerVariableService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *TagManagerVariableGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/tag-manager-variables/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Partially update a variable. Only the fields you send are changed. Name
// collisions with other variables in the same tag manager return 409 with the
// reason in the response `error` field. To assign a variable to a folder, use
// `POST /rest/v1/tag-manager-asset-folders`. Requires API-key scope or current
// OAuth user permission: tagManagers:update
func (r *TagManagerVariableService) Update(ctx context.Context, id string, body TagManagerVariableUpdateParams, opts ...option.RequestOption) (res *TagManagerVariableUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/tag-manager-variables/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete a tag manager variable. Requires API-key scope or current OAuth user
// permission: tagManagers:update
func (r *TagManagerVariableService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *TagManagerVariableDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/tag-manager-variables/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Lists every variable template the platform supports — what `type` discriminator
// to send on create/patch, the shape of the type-specific `parameters` payload,
// and `supportsVariables` (whether the variable's own parameter fields may
// reference `{{OtherVariable}}` at runtime). Account-agnostic: the response is the
// same for every API key. Requires API-key scope or current OAuth user permission:
// tagManagers:find
func (r *TagManagerVariableService) Types(ctx context.Context, opts ...option.RequestOption) (res *TagManagerVariableTypesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/tag-manager-variables/types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type TagManagerVariableListResponse struct {
	ID        string `json:"id" api:"required"`
	AccountID string `json:"accountId" api:"required"`
	Name      string `json:"name" api:"required"`
	// Type-specific configuration.
	Parameters   map[string]any `json:"parameters" api:"required"`
	TagManagerID string         `json:"tagManagerId" api:"required"`
	// Variable type discriminator. Examples that exist today: `DataLayer`, `Constant`,
	// `Cookie`, `Url`, `UrlParameter`, `Weekday`, `RandomNumber`. Pick from
	// `GET /tag-manager-variables/types` for the canonical set.
	Type      string `json:"type" api:"required"`
	CreatedAt string `json:"createdAt" api:"nullable"`
	// Default value returned when no rule matches. JSON value — type depends on
	// `type`.
	DefaultValue TagManagerVariableListResponseDefaultValueUnion `json:"defaultValue" api:"nullable"`
	Enabled      bool                                            `json:"enabled" api:"nullable"`
	// Folder this variable belongs to. Settable via PATCH — send a folder UUID to
	// assign, or `null` to remove from its current folder.
	FolderID string `json:"folderId" api:"nullable"`
	// Optional lookup table for `LookUpTable`-style variables. JSON value.
	LookUpTable TagManagerVariableListResponseLookUpTableUnion `json:"lookUpTable" api:"nullable"`
	UpdatedAt   string                                         `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AccountID    respjson.Field
		Name         respjson.Field
		Parameters   respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
		CreatedAt    respjson.Field
		DefaultValue respjson.Field
		Enabled      respjson.Field
		FolderID     respjson.Field
		LookUpTable  respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagManagerVariableListResponse) RawJSON() string { return r.JSON.raw }
func (r *TagManagerVariableListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TagManagerVariableListResponseDefaultValueUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfTagManagerVariableListResponseDefaultValueMapItem]
type TagManagerVariableListResponseDefaultValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfTagManagerVariableListResponseDefaultValueMapItem any `json:",inline"`
	JSON                                                struct {
		OfString                                            respjson.Field
		OfFloat                                             respjson.Field
		OfBool                                              respjson.Field
		OfAnyArray                                          respjson.Field
		OfTagManagerVariableListResponseDefaultValueMapItem respjson.Field
		raw                                                 string
	} `json:"-"`
}

func (u TagManagerVariableListResponseDefaultValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableListResponseDefaultValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableListResponseDefaultValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableListResponseDefaultValueUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableListResponseDefaultValueUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TagManagerVariableListResponseDefaultValueUnion) RawJSON() string { return u.JSON.raw }

func (r *TagManagerVariableListResponseDefaultValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TagManagerVariableListResponseLookUpTableUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfTagManagerVariableListResponseLookUpTableMapItem]
type TagManagerVariableListResponseLookUpTableUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfTagManagerVariableListResponseLookUpTableMapItem any `json:",inline"`
	JSON                                               struct {
		OfString                                           respjson.Field
		OfFloat                                            respjson.Field
		OfBool                                             respjson.Field
		OfAnyArray                                         respjson.Field
		OfTagManagerVariableListResponseLookUpTableMapItem respjson.Field
		raw                                                string
	} `json:"-"`
}

func (u TagManagerVariableListResponseLookUpTableUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableListResponseLookUpTableUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableListResponseLookUpTableUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableListResponseLookUpTableUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableListResponseLookUpTableUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TagManagerVariableListResponseLookUpTableUnion) RawJSON() string { return u.JSON.raw }

func (r *TagManagerVariableListResponseLookUpTableUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerVariableNewResponse struct {
	ID        string `json:"id" api:"required"`
	AccountID string `json:"accountId" api:"required"`
	Name      string `json:"name" api:"required"`
	// Type-specific configuration.
	Parameters   map[string]any `json:"parameters" api:"required"`
	TagManagerID string         `json:"tagManagerId" api:"required"`
	// Variable type discriminator. Examples that exist today: `DataLayer`, `Constant`,
	// `Cookie`, `Url`, `UrlParameter`, `Weekday`, `RandomNumber`. Pick from
	// `GET /tag-manager-variables/types` for the canonical set.
	Type      string `json:"type" api:"required"`
	CreatedAt string `json:"createdAt" api:"nullable"`
	// Default value returned when no rule matches. JSON value — type depends on
	// `type`.
	DefaultValue TagManagerVariableNewResponseDefaultValueUnion `json:"defaultValue" api:"nullable"`
	Enabled      bool                                           `json:"enabled" api:"nullable"`
	// Folder this variable belongs to. Settable via PATCH — send a folder UUID to
	// assign, or `null` to remove from its current folder.
	FolderID string `json:"folderId" api:"nullable"`
	// Optional lookup table for `LookUpTable`-style variables. JSON value.
	LookUpTable TagManagerVariableNewResponseLookUpTableUnion `json:"lookUpTable" api:"nullable"`
	UpdatedAt   string                                        `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AccountID    respjson.Field
		Name         respjson.Field
		Parameters   respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
		CreatedAt    respjson.Field
		DefaultValue respjson.Field
		Enabled      respjson.Field
		FolderID     respjson.Field
		LookUpTable  respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagManagerVariableNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TagManagerVariableNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TagManagerVariableNewResponseDefaultValueUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfTagManagerVariableNewResponseDefaultValueMapItem]
type TagManagerVariableNewResponseDefaultValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfTagManagerVariableNewResponseDefaultValueMapItem any `json:",inline"`
	JSON                                               struct {
		OfString                                           respjson.Field
		OfFloat                                            respjson.Field
		OfBool                                             respjson.Field
		OfAnyArray                                         respjson.Field
		OfTagManagerVariableNewResponseDefaultValueMapItem respjson.Field
		raw                                                string
	} `json:"-"`
}

func (u TagManagerVariableNewResponseDefaultValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableNewResponseDefaultValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableNewResponseDefaultValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableNewResponseDefaultValueUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableNewResponseDefaultValueUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TagManagerVariableNewResponseDefaultValueUnion) RawJSON() string { return u.JSON.raw }

func (r *TagManagerVariableNewResponseDefaultValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TagManagerVariableNewResponseLookUpTableUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfTagManagerVariableNewResponseLookUpTableMapItem]
type TagManagerVariableNewResponseLookUpTableUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfTagManagerVariableNewResponseLookUpTableMapItem any `json:",inline"`
	JSON                                              struct {
		OfString                                          respjson.Field
		OfFloat                                           respjson.Field
		OfBool                                            respjson.Field
		OfAnyArray                                        respjson.Field
		OfTagManagerVariableNewResponseLookUpTableMapItem respjson.Field
		raw                                               string
	} `json:"-"`
}

func (u TagManagerVariableNewResponseLookUpTableUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableNewResponseLookUpTableUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableNewResponseLookUpTableUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableNewResponseLookUpTableUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableNewResponseLookUpTableUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TagManagerVariableNewResponseLookUpTableUnion) RawJSON() string { return u.JSON.raw }

func (r *TagManagerVariableNewResponseLookUpTableUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerVariableGetResponse struct {
	ID        string `json:"id" api:"required"`
	AccountID string `json:"accountId" api:"required"`
	Name      string `json:"name" api:"required"`
	// Type-specific configuration.
	Parameters   map[string]any `json:"parameters" api:"required"`
	TagManagerID string         `json:"tagManagerId" api:"required"`
	// Variable type discriminator. Examples that exist today: `DataLayer`, `Constant`,
	// `Cookie`, `Url`, `UrlParameter`, `Weekday`, `RandomNumber`. Pick from
	// `GET /tag-manager-variables/types` for the canonical set.
	Type      string `json:"type" api:"required"`
	CreatedAt string `json:"createdAt" api:"nullable"`
	// Default value returned when no rule matches. JSON value — type depends on
	// `type`.
	DefaultValue TagManagerVariableGetResponseDefaultValueUnion `json:"defaultValue" api:"nullable"`
	Enabled      bool                                           `json:"enabled" api:"nullable"`
	// Folder this variable belongs to. Settable via PATCH — send a folder UUID to
	// assign, or `null` to remove from its current folder.
	FolderID string `json:"folderId" api:"nullable"`
	// Optional lookup table for `LookUpTable`-style variables. JSON value.
	LookUpTable TagManagerVariableGetResponseLookUpTableUnion `json:"lookUpTable" api:"nullable"`
	UpdatedAt   string                                        `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AccountID    respjson.Field
		Name         respjson.Field
		Parameters   respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
		CreatedAt    respjson.Field
		DefaultValue respjson.Field
		Enabled      respjson.Field
		FolderID     respjson.Field
		LookUpTable  respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagManagerVariableGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TagManagerVariableGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TagManagerVariableGetResponseDefaultValueUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfTagManagerVariableGetResponseDefaultValueMapItem]
type TagManagerVariableGetResponseDefaultValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfTagManagerVariableGetResponseDefaultValueMapItem any `json:",inline"`
	JSON                                               struct {
		OfString                                           respjson.Field
		OfFloat                                            respjson.Field
		OfBool                                             respjson.Field
		OfAnyArray                                         respjson.Field
		OfTagManagerVariableGetResponseDefaultValueMapItem respjson.Field
		raw                                                string
	} `json:"-"`
}

func (u TagManagerVariableGetResponseDefaultValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableGetResponseDefaultValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableGetResponseDefaultValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableGetResponseDefaultValueUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableGetResponseDefaultValueUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TagManagerVariableGetResponseDefaultValueUnion) RawJSON() string { return u.JSON.raw }

func (r *TagManagerVariableGetResponseDefaultValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TagManagerVariableGetResponseLookUpTableUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfTagManagerVariableGetResponseLookUpTableMapItem]
type TagManagerVariableGetResponseLookUpTableUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfTagManagerVariableGetResponseLookUpTableMapItem any `json:",inline"`
	JSON                                              struct {
		OfString                                          respjson.Field
		OfFloat                                           respjson.Field
		OfBool                                            respjson.Field
		OfAnyArray                                        respjson.Field
		OfTagManagerVariableGetResponseLookUpTableMapItem respjson.Field
		raw                                               string
	} `json:"-"`
}

func (u TagManagerVariableGetResponseLookUpTableUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableGetResponseLookUpTableUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableGetResponseLookUpTableUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableGetResponseLookUpTableUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableGetResponseLookUpTableUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TagManagerVariableGetResponseLookUpTableUnion) RawJSON() string { return u.JSON.raw }

func (r *TagManagerVariableGetResponseLookUpTableUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerVariableUpdateResponse struct {
	ID        string `json:"id" api:"required"`
	AccountID string `json:"accountId" api:"required"`
	Name      string `json:"name" api:"required"`
	// Type-specific configuration.
	Parameters   map[string]any `json:"parameters" api:"required"`
	TagManagerID string         `json:"tagManagerId" api:"required"`
	// Variable type discriminator. Examples that exist today: `DataLayer`, `Constant`,
	// `Cookie`, `Url`, `UrlParameter`, `Weekday`, `RandomNumber`. Pick from
	// `GET /tag-manager-variables/types` for the canonical set.
	Type      string `json:"type" api:"required"`
	CreatedAt string `json:"createdAt" api:"nullable"`
	// Default value returned when no rule matches. JSON value — type depends on
	// `type`.
	DefaultValue TagManagerVariableUpdateResponseDefaultValueUnion `json:"defaultValue" api:"nullable"`
	Enabled      bool                                              `json:"enabled" api:"nullable"`
	// Folder this variable belongs to. Settable via PATCH — send a folder UUID to
	// assign, or `null` to remove from its current folder.
	FolderID string `json:"folderId" api:"nullable"`
	// Optional lookup table for `LookUpTable`-style variables. JSON value.
	LookUpTable TagManagerVariableUpdateResponseLookUpTableUnion `json:"lookUpTable" api:"nullable"`
	UpdatedAt   string                                           `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AccountID    respjson.Field
		Name         respjson.Field
		Parameters   respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
		CreatedAt    respjson.Field
		DefaultValue respjson.Field
		Enabled      respjson.Field
		FolderID     respjson.Field
		LookUpTable  respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagManagerVariableUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *TagManagerVariableUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TagManagerVariableUpdateResponseDefaultValueUnion contains all possible
// properties and values from [string], [float64], [bool], [[]any],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfTagManagerVariableUpdateResponseDefaultValueMapItem]
type TagManagerVariableUpdateResponseDefaultValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfTagManagerVariableUpdateResponseDefaultValueMapItem any `json:",inline"`
	JSON                                                  struct {
		OfString                                              respjson.Field
		OfFloat                                               respjson.Field
		OfBool                                                respjson.Field
		OfAnyArray                                            respjson.Field
		OfTagManagerVariableUpdateResponseDefaultValueMapItem respjson.Field
		raw                                                   string
	} `json:"-"`
}

func (u TagManagerVariableUpdateResponseDefaultValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableUpdateResponseDefaultValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableUpdateResponseDefaultValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableUpdateResponseDefaultValueUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableUpdateResponseDefaultValueUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TagManagerVariableUpdateResponseDefaultValueUnion) RawJSON() string { return u.JSON.raw }

func (r *TagManagerVariableUpdateResponseDefaultValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TagManagerVariableUpdateResponseLookUpTableUnion contains all possible
// properties and values from [string], [float64], [bool], [[]any],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfTagManagerVariableUpdateResponseLookUpTableMapItem]
type TagManagerVariableUpdateResponseLookUpTableUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfTagManagerVariableUpdateResponseLookUpTableMapItem any `json:",inline"`
	JSON                                                 struct {
		OfString                                             respjson.Field
		OfFloat                                              respjson.Field
		OfBool                                               respjson.Field
		OfAnyArray                                           respjson.Field
		OfTagManagerVariableUpdateResponseLookUpTableMapItem respjson.Field
		raw                                                  string
	} `json:"-"`
}

func (u TagManagerVariableUpdateResponseLookUpTableUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableUpdateResponseLookUpTableUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableUpdateResponseLookUpTableUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableUpdateResponseLookUpTableUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableUpdateResponseLookUpTableUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TagManagerVariableUpdateResponseLookUpTableUnion) RawJSON() string { return u.JSON.raw }

func (r *TagManagerVariableUpdateResponseLookUpTableUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerVariableDeleteResponse struct {
	ID      string `json:"id" api:"required"`
	Deleted bool   `json:"deleted" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Deleted     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagManagerVariableDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *TagManagerVariableDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerVariableTypesResponse struct {
	Entities []TagManagerVariableTypesResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagManagerVariableTypesResponse) RawJSON() string { return r.JSON.raw }
func (r *TagManagerVariableTypesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerVariableTypesResponseEntity struct {
	// Type discriminator — pass this as `type` on create/patch.
	ID string `json:"id" api:"required"`
	// Grouping label.
	Category string                                       `json:"category" api:"required"`
	Fields   []TagManagerVariableTypesResponseEntityField `json:"fields" api:"required"`
	// Human-readable display name.
	Name        string `json:"name" api:"required"`
	Description string `json:"description" api:"nullable"`
	// When `true`, this variable type's parameter fields can themselves contain
	// `{{OtherVariable}}` references that the SDK resolves at runtime.
	SupportsVariables bool `json:"supportsVariables" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Category          respjson.Field
		Fields            respjson.Field
		Name              respjson.Field
		Description       respjson.Field
		SupportsVariables respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagManagerVariableTypesResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *TagManagerVariableTypesResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerVariableTypesResponseEntityField struct {
	// Parameter key that goes in the `parameters` payload at create/patch.
	ID string `json:"id" api:"required"`
	// Human-readable title for the field.
	Title string `json:"title" api:"required"`
	// Underlying data type of the parameter value.
	//
	// Any of "STRING", "BOOLEAN", "INTEGER", "FLOAT", "TABLE".
	Type string `json:"type" api:"required"`
	// For TABLE-typed fields, the predefined keys each row may contain.
	AllowedKeys []string `json:"allowedKeys" api:"nullable"`
	// When present, the field accepts only one of these values. Send the `value`
	// string in `parameters`; the `label` is for display.
	AvailableValues []TagManagerVariableTypesResponseEntityFieldAvailableValue `json:"availableValues" api:"nullable"`
	// Default value when the caller omits the parameter on create.
	Default     TagManagerVariableTypesResponseEntityFieldDefaultUnion `json:"default" api:"nullable"`
	Description string                                                 `json:"description" api:"nullable"`
	// When `true`, omitting or sending an empty value for this parameter on
	// create/patch returns HTTP 400.
	Required bool `json:"required" api:"nullable"`
	// Server-enforced rules applied to this field at create and patch.
	Validators []TagManagerVariableTypesResponseEntityFieldValidator `json:"validators" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Title           respjson.Field
		Type            respjson.Field
		AllowedKeys     respjson.Field
		AvailableValues respjson.Field
		Default         respjson.Field
		Description     respjson.Field
		Required        respjson.Field
		Validators      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagManagerVariableTypesResponseEntityField) RawJSON() string { return r.JSON.raw }
func (r *TagManagerVariableTypesResponseEntityField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerVariableTypesResponseEntityFieldAvailableValue struct {
	Label string `json:"label" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagManagerVariableTypesResponseEntityFieldAvailableValue) RawJSON() string { return r.JSON.raw }
func (r *TagManagerVariableTypesResponseEntityFieldAvailableValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TagManagerVariableTypesResponseEntityFieldDefaultUnion contains all possible
// properties and values from [string], [float64], [bool], [[]any],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfTagManagerVariableTypesResponseEntityFieldDefaultMapItem]
type TagManagerVariableTypesResponseEntityFieldDefaultUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfTagManagerVariableTypesResponseEntityFieldDefaultMapItem any `json:",inline"`
	JSON                                                       struct {
		OfString                                                   respjson.Field
		OfFloat                                                    respjson.Field
		OfBool                                                     respjson.Field
		OfAnyArray                                                 respjson.Field
		OfTagManagerVariableTypesResponseEntityFieldDefaultMapItem respjson.Field
		raw                                                        string
	} `json:"-"`
}

func (u TagManagerVariableTypesResponseEntityFieldDefaultUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableTypesResponseEntityFieldDefaultUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableTypesResponseEntityFieldDefaultUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableTypesResponseEntityFieldDefaultUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TagManagerVariableTypesResponseEntityFieldDefaultUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TagManagerVariableTypesResponseEntityFieldDefaultUnion) RawJSON() string { return u.JSON.raw }

func (r *TagManagerVariableTypesResponseEntityFieldDefaultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerVariableTypesResponseEntityFieldValidator struct {
	// Any of "NotEmpty", "CharacterLength", "Url", "Email", "Number", "Range".
	Type string  `json:"type" api:"required"`
	Max  float64 `json:"max" api:"nullable"`
	Min  float64 `json:"min" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Max         respjson.Field
		Min         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagManagerVariableTypesResponseEntityFieldValidator) RawJSON() string { return r.JSON.raw }
func (r *TagManagerVariableTypesResponseEntityFieldValidator) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerVariableListParams struct {
	// Parent tag manager whose variables should be returned.
	TagManagerID string `query:"tagManagerId" api:"required" json:"-"`
	// Maximum number of variables to return. Defaults to 25; values below 1 are
	// clamped to 1 and values above 1000 are clamped to 1000. The web-app passes 1000
	// to render the full workspace in one request.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Opaque pagination cursor from pagination.nextCursor in the previous response. Do
	// not decode or modify it. Malformed cursors return 400 Bad Request.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TagManagerVariableListParams]'s query parameters as
// `url.Values`.
func (r TagManagerVariableListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TagManagerVariableNewParams struct {
	Name string `json:"name" api:"required"`
	// Type-specific JSON configuration.
	Parameters map[string]any `json:"parameters,omitzero" api:"required"`
	// Parent tag manager that will own the new variable.
	TagManagerID string `json:"tagManagerId" api:"required"`
	// Variable type discriminator. Pick from `GET /tag-manager-variables/types` for
	// the canonical set (e.g. `DataLayer`, `Constant`, `Cookie`, `Url`).
	Type    string          `json:"type" api:"required"`
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Optional default value. JSON value of any type.
	DefaultValue TagManagerVariableNewParamsDefaultValueUnion `json:"defaultValue,omitzero"`
	// Optional lookup table for `LookUpTable` variables.
	LookUpTable TagManagerVariableNewParamsLookUpTableUnion `json:"lookUpTable,omitzero"`
	paramObj
}

func (r TagManagerVariableNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TagManagerVariableNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagManagerVariableNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TagManagerVariableNewParamsDefaultValueUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u TagManagerVariableNewParamsDefaultValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *TagManagerVariableNewParamsDefaultValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TagManagerVariableNewParamsLookUpTableUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u TagManagerVariableNewParamsLookUpTableUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *TagManagerVariableNewParamsLookUpTableUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type TagManagerVariableUpdateParams struct {
	// Pause/resume the variable without changing other fields.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Updated variable name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Updated variable type. Pick from `GET /tag-manager-variables/types`.
	Type param.Opt[string] `json:"type,omitzero"`
	// Updated default value. JSON value of any type.
	DefaultValue TagManagerVariableUpdateParamsDefaultValueUnion `json:"defaultValue,omitzero"`
	// Updated lookup table payload.
	LookUpTable TagManagerVariableUpdateParamsLookUpTableUnion `json:"lookUpTable,omitzero"`
	// Updated type-specific JSON configuration.
	Parameters map[string]any `json:"parameters,omitzero"`
	paramObj
}

func (r TagManagerVariableUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TagManagerVariableUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagManagerVariableUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TagManagerVariableUpdateParamsDefaultValueUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u TagManagerVariableUpdateParamsDefaultValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *TagManagerVariableUpdateParamsDefaultValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TagManagerVariableUpdateParamsLookUpTableUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u TagManagerVariableUpdateParamsLookUpTableUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *TagManagerVariableUpdateParamsLookUpTableUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}
