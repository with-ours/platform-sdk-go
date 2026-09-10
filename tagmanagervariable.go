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
	"github.com/with-ours/platform-sdk-go/internal/apiquery"
	"github.com/with-ours/platform-sdk-go/internal/requestconfig"
	"github.com/with-ours/platform-sdk-go/option"
	"github.com/with-ours/platform-sdk-go/packages/pagination"
	"github.com/with-ours/platform-sdk-go/packages/param"
	"github.com/with-ours/platform-sdk-go/packages/respjson"
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
// shot). Requires scope: tagManagers:find
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
// shot). Requires scope: tagManagers:find
func (r *TagManagerVariableService) ListAutoPaging(ctx context.Context, query TagManagerVariableListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[TagManagerVariableListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Create a new variable inside a tag manager. `tagManagerId` is required in the
// body. Known input failures (e.g. duplicate variable name within the tag manager)
// are returned as HTTP 409 with the reason in the response `error` field. Requires
// scope: tagManagers:update
func (r *TagManagerVariableService) New(ctx context.Context, body TagManagerVariableNewParams, opts ...option.RequestOption) (res *TagManagerVariableNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/tag-manager-variables"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Find a single tag manager variable by ID. Requires scope: tagManagers:find
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
// `POST /rest/v1/tag-manager-asset-folders`. Requires scope: tagManagers:update
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

// Delete a tag manager variable. Requires scope: tagManagers:update
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
// same for every API key. Requires scope: tagManagers:find
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
	DefaultValue map[string]any `json:"defaultValue"`
	Enabled      bool           `json:"enabled" api:"nullable"`
	// Folder this variable belongs to. Settable via PATCH — send a folder UUID to
	// assign, or `null` to remove from its current folder.
	FolderID string `json:"folderId" api:"nullable"`
	// Optional lookup table for `LookUpTable`-style variables. JSON value.
	LookUpTable map[string]any `json:"lookUpTable"`
	UpdatedAt   string         `json:"updatedAt" api:"nullable"`
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
	DefaultValue map[string]any `json:"defaultValue"`
	Enabled      bool           `json:"enabled" api:"nullable"`
	// Folder this variable belongs to. Settable via PATCH — send a folder UUID to
	// assign, or `null` to remove from its current folder.
	FolderID string `json:"folderId" api:"nullable"`
	// Optional lookup table for `LookUpTable`-style variables. JSON value.
	LookUpTable map[string]any `json:"lookUpTable"`
	UpdatedAt   string         `json:"updatedAt" api:"nullable"`
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
	DefaultValue map[string]any `json:"defaultValue"`
	Enabled      bool           `json:"enabled" api:"nullable"`
	// Folder this variable belongs to. Settable via PATCH — send a folder UUID to
	// assign, or `null` to remove from its current folder.
	FolderID string `json:"folderId" api:"nullable"`
	// Optional lookup table for `LookUpTable`-style variables. JSON value.
	LookUpTable map[string]any `json:"lookUpTable"`
	UpdatedAt   string         `json:"updatedAt" api:"nullable"`
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
	DefaultValue map[string]any `json:"defaultValue"`
	Enabled      bool           `json:"enabled" api:"nullable"`
	// Folder this variable belongs to. Settable via PATCH — send a folder UUID to
	// assign, or `null` to remove from its current folder.
	FolderID string `json:"folderId" api:"nullable"`
	// Optional lookup table for `LookUpTable`-style variables. JSON value.
	LookUpTable map[string]any `json:"lookUpTable"`
	UpdatedAt   string         `json:"updatedAt" api:"nullable"`
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
	Default     map[string]any `json:"default"`
	Description string         `json:"description" api:"nullable"`
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
	DefaultValue map[string]any `json:"defaultValue,omitzero"`
	// Optional lookup table for `LookUpTable` variables.
	LookUpTable map[string]any `json:"lookUpTable,omitzero"`
	paramObj
}

func (r TagManagerVariableNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TagManagerVariableNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagManagerVariableNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerVariableUpdateParams struct {
	// Pause/resume the variable without changing other fields.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Updated variable name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Updated variable type. Pick from `GET /tag-manager-variables/types`.
	Type param.Opt[string] `json:"type,omitzero"`
	// Updated default value. JSON value of any type.
	DefaultValue map[string]any `json:"defaultValue,omitzero"`
	// Updated lookup table payload.
	LookUpTable map[string]any `json:"lookUpTable,omitzero"`
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
