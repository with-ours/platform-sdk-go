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

// TagManagerTriggerService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTagManagerTriggerService] method instead.
type TagManagerTriggerService struct {
	Options []option.RequestOption
}

// NewTagManagerTriggerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTagManagerTriggerService(opts ...option.RequestOption) (r TagManagerTriggerService) {
	r = TagManagerTriggerService{}
	r.Options = opts
	return
}

// List triggers inside a single tag manager. Requires the `tagManagerId` query
// parameter — triggers are always scoped to one parent container. Supports cursor
// pagination via `limit` and `cursor`; the limit clamp is 1000 so a single request
// can return the full set (the web-app workspace renders all triggers in one
// shot). Requires scope: tagManagers:find
func (r *TagManagerTriggerService) List(ctx context.Context, query TagManagerTriggerListParams, opts ...option.RequestOption) (res *pagination.Cursor[TagManagerTriggerListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "rest/v1/tag-manager-triggers"
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

// List triggers inside a single tag manager. Requires the `tagManagerId` query
// parameter — triggers are always scoped to one parent container. Supports cursor
// pagination via `limit` and `cursor`; the limit clamp is 1000 so a single request
// can return the full set (the web-app workspace renders all triggers in one
// shot). Requires scope: tagManagers:find
func (r *TagManagerTriggerService) ListAutoPaging(ctx context.Context, query TagManagerTriggerListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[TagManagerTriggerListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Create a new trigger inside a tag manager. `tagManagerId` is required in the
// body. Send `conditions: []` for an unconditional trigger; otherwise supply
// type-specific condition objects. Requires scope: tagManagers:update
func (r *TagManagerTriggerService) New(ctx context.Context, body TagManagerTriggerNewParams, opts ...option.RequestOption) (res *TagManagerTriggerNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/tag-manager-triggers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Find a single tag manager trigger by ID. Requires scope: tagManagers:find
func (r *TagManagerTriggerService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *TagManagerTriggerGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/tag-manager-triggers/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Partially update a trigger. Only the fields you send are changed. `conditions`
// is replaced wholesale when sent. Requires scope: tagManagers:update
func (r *TagManagerTriggerService) Update(ctx context.Context, id string, body TagManagerTriggerUpdateParams, opts ...option.RequestOption) (res *TagManagerTriggerUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/tag-manager-triggers/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete a tag manager trigger. Requires scope: tagManagers:update
func (r *TagManagerTriggerService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *TagManagerTriggerDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/tag-manager-triggers/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Lists every trigger template the platform supports — what `type` discriminator
// to send on create/patch, and the shape of the type-specific `parameters`
// payload. Trigger `conditions` are evaluated at runtime (per-trigger, see the
// resource docs) and are not part of this descriptor. Account-agnostic: the
// response is the same for every API key. Requires scope: tagManagers:find
func (r *TagManagerTriggerService) Types(ctx context.Context, opts ...option.RequestOption) (res *TagManagerTriggerTypesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/tag-manager-triggers/types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type TagManagerTriggerListResponse struct {
	ID        string `json:"id" api:"required"`
	AccountID string `json:"accountId" api:"required"`
	// Conditions that must hold for the trigger to match. Use `[]` for unconditional
	// triggers.
	Conditions []map[string]any `json:"conditions" api:"required"`
	Name       string           `json:"name" api:"required"`
	// Type-specific configuration. Send `{}` for a no-op trigger.
	Parameters   map[string]any `json:"parameters" api:"required"`
	TagManagerID string         `json:"tagManagerId" api:"required"`
	// Must equal `type` — send the same string in both fields. The server rejects any
	// divergent value.
	Trigger string `json:"Trigger" api:"required"`
	// Trigger type discriminator. Examples that exist today: `PageView`, `DomReady`,
	// `Initialization`, `AllElementsClick`, `AllLinksClick`, `FormSubmit`,
	// `CustomEvent`, `ScrollReach`, `Timer`. Pick from
	// `GET /tag-manager-triggers/types` for the canonical set. Note there is no plain
	// `Click` id; use one of the `All*Click` variants.
	Type      string `json:"type" api:"required"`
	CreatedAt string `json:"createdAt" api:"nullable"`
	Enabled   bool   `json:"enabled" api:"nullable"`
	// Folder this trigger belongs to. Read-only on this endpoint — use the GraphQL
	// `assignTagManagerAssetToFolder` mutation to change.
	FolderID  string `json:"folderId" api:"nullable"`
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AccountID    respjson.Field
		Conditions   respjson.Field
		Name         respjson.Field
		Parameters   respjson.Field
		TagManagerID respjson.Field
		Trigger      respjson.Field
		Type         respjson.Field
		CreatedAt    respjson.Field
		Enabled      respjson.Field
		FolderID     respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagManagerTriggerListResponse) RawJSON() string { return r.JSON.raw }
func (r *TagManagerTriggerListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerTriggerNewResponse struct {
	ID        string `json:"id" api:"required"`
	AccountID string `json:"accountId" api:"required"`
	// Conditions that must hold for the trigger to match. Use `[]` for unconditional
	// triggers.
	Conditions []map[string]any `json:"conditions" api:"required"`
	Name       string           `json:"name" api:"required"`
	// Type-specific configuration. Send `{}` for a no-op trigger.
	Parameters   map[string]any `json:"parameters" api:"required"`
	TagManagerID string         `json:"tagManagerId" api:"required"`
	// Must equal `type` — send the same string in both fields. The server rejects any
	// divergent value.
	Trigger string `json:"Trigger" api:"required"`
	// Trigger type discriminator. Examples that exist today: `PageView`, `DomReady`,
	// `Initialization`, `AllElementsClick`, `AllLinksClick`, `FormSubmit`,
	// `CustomEvent`, `ScrollReach`, `Timer`. Pick from
	// `GET /tag-manager-triggers/types` for the canonical set. Note there is no plain
	// `Click` id; use one of the `All*Click` variants.
	Type      string `json:"type" api:"required"`
	CreatedAt string `json:"createdAt" api:"nullable"`
	Enabled   bool   `json:"enabled" api:"nullable"`
	// Folder this trigger belongs to. Read-only on this endpoint — use the GraphQL
	// `assignTagManagerAssetToFolder` mutation to change.
	FolderID  string `json:"folderId" api:"nullable"`
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AccountID    respjson.Field
		Conditions   respjson.Field
		Name         respjson.Field
		Parameters   respjson.Field
		TagManagerID respjson.Field
		Trigger      respjson.Field
		Type         respjson.Field
		CreatedAt    respjson.Field
		Enabled      respjson.Field
		FolderID     respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagManagerTriggerNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TagManagerTriggerNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerTriggerGetResponse struct {
	ID        string `json:"id" api:"required"`
	AccountID string `json:"accountId" api:"required"`
	// Conditions that must hold for the trigger to match. Use `[]` for unconditional
	// triggers.
	Conditions []map[string]any `json:"conditions" api:"required"`
	Name       string           `json:"name" api:"required"`
	// Type-specific configuration. Send `{}` for a no-op trigger.
	Parameters   map[string]any `json:"parameters" api:"required"`
	TagManagerID string         `json:"tagManagerId" api:"required"`
	// Must equal `type` — send the same string in both fields. The server rejects any
	// divergent value.
	Trigger string `json:"Trigger" api:"required"`
	// Trigger type discriminator. Examples that exist today: `PageView`, `DomReady`,
	// `Initialization`, `AllElementsClick`, `AllLinksClick`, `FormSubmit`,
	// `CustomEvent`, `ScrollReach`, `Timer`. Pick from
	// `GET /tag-manager-triggers/types` for the canonical set. Note there is no plain
	// `Click` id; use one of the `All*Click` variants.
	Type      string `json:"type" api:"required"`
	CreatedAt string `json:"createdAt" api:"nullable"`
	Enabled   bool   `json:"enabled" api:"nullable"`
	// Folder this trigger belongs to. Read-only on this endpoint — use the GraphQL
	// `assignTagManagerAssetToFolder` mutation to change.
	FolderID  string `json:"folderId" api:"nullable"`
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AccountID    respjson.Field
		Conditions   respjson.Field
		Name         respjson.Field
		Parameters   respjson.Field
		TagManagerID respjson.Field
		Trigger      respjson.Field
		Type         respjson.Field
		CreatedAt    respjson.Field
		Enabled      respjson.Field
		FolderID     respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagManagerTriggerGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TagManagerTriggerGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerTriggerUpdateResponse struct {
	ID        string `json:"id" api:"required"`
	AccountID string `json:"accountId" api:"required"`
	// Conditions that must hold for the trigger to match. Use `[]` for unconditional
	// triggers.
	Conditions []map[string]any `json:"conditions" api:"required"`
	Name       string           `json:"name" api:"required"`
	// Type-specific configuration. Send `{}` for a no-op trigger.
	Parameters   map[string]any `json:"parameters" api:"required"`
	TagManagerID string         `json:"tagManagerId" api:"required"`
	// Must equal `type` — send the same string in both fields. The server rejects any
	// divergent value.
	Trigger string `json:"Trigger" api:"required"`
	// Trigger type discriminator. Examples that exist today: `PageView`, `DomReady`,
	// `Initialization`, `AllElementsClick`, `AllLinksClick`, `FormSubmit`,
	// `CustomEvent`, `ScrollReach`, `Timer`. Pick from
	// `GET /tag-manager-triggers/types` for the canonical set. Note there is no plain
	// `Click` id; use one of the `All*Click` variants.
	Type      string `json:"type" api:"required"`
	CreatedAt string `json:"createdAt" api:"nullable"`
	Enabled   bool   `json:"enabled" api:"nullable"`
	// Folder this trigger belongs to. Read-only on this endpoint — use the GraphQL
	// `assignTagManagerAssetToFolder` mutation to change.
	FolderID  string `json:"folderId" api:"nullable"`
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AccountID    respjson.Field
		Conditions   respjson.Field
		Name         respjson.Field
		Parameters   respjson.Field
		TagManagerID respjson.Field
		Trigger      respjson.Field
		Type         respjson.Field
		CreatedAt    respjson.Field
		Enabled      respjson.Field
		FolderID     respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagManagerTriggerUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *TagManagerTriggerUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerTriggerDeleteResponse struct {
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
func (r TagManagerTriggerDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *TagManagerTriggerDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerTriggerTypesResponse struct {
	Entities []TagManagerTriggerTypesResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagManagerTriggerTypesResponse) RawJSON() string { return r.JSON.raw }
func (r *TagManagerTriggerTypesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerTriggerTypesResponseEntity struct {
	// Type discriminator — pass this as `type` on create/patch.
	ID string `json:"id" api:"required"`
	// Grouping label.
	Category string                                      `json:"category" api:"required"`
	Fields   []TagManagerTriggerTypesResponseEntityField `json:"fields" api:"required"`
	// Human-readable display name.
	Name        string `json:"name" api:"required"`
	Description string `json:"description" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Category    respjson.Field
		Fields      respjson.Field
		Name        respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagManagerTriggerTypesResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *TagManagerTriggerTypesResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerTriggerTypesResponseEntityField struct {
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
	AvailableValues []TagManagerTriggerTypesResponseEntityFieldAvailableValue `json:"availableValues" api:"nullable"`
	// Default value when the caller omits the parameter on create.
	Default     map[string]any `json:"default"`
	Description string         `json:"description" api:"nullable"`
	// When `true`, omitting or sending an empty value for this parameter on
	// create/patch returns HTTP 400.
	Required bool `json:"required" api:"nullable"`
	// Server-enforced rules applied to this field at create and patch.
	Validators []TagManagerTriggerTypesResponseEntityFieldValidator `json:"validators" api:"nullable"`
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
func (r TagManagerTriggerTypesResponseEntityField) RawJSON() string { return r.JSON.raw }
func (r *TagManagerTriggerTypesResponseEntityField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerTriggerTypesResponseEntityFieldAvailableValue struct {
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
func (r TagManagerTriggerTypesResponseEntityFieldAvailableValue) RawJSON() string { return r.JSON.raw }
func (r *TagManagerTriggerTypesResponseEntityFieldAvailableValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerTriggerTypesResponseEntityFieldValidator struct {
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
func (r TagManagerTriggerTypesResponseEntityFieldValidator) RawJSON() string { return r.JSON.raw }
func (r *TagManagerTriggerTypesResponseEntityFieldValidator) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerTriggerListParams struct {
	// Parent tag manager whose triggers should be returned.
	TagManagerID string `query:"tagManagerId" api:"required" json:"-"`
	// Maximum number of triggers to return. Defaults to 25; values below 1 are clamped
	// to 1 and values above 1000 are clamped to 1000. The web-app passes 1000 to
	// render the full workspace in one request.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Opaque pagination cursor from pagination.nextCursor in the previous response. Do
	// not decode or modify it. Malformed cursors return 400 Bad Request.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TagManagerTriggerListParams]'s query parameters as
// `url.Values`.
func (r TagManagerTriggerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TagManagerTriggerNewParams struct {
	// Match conditions; use `[]` for an unconditional trigger.
	Conditions []map[string]any `json:"conditions,omitzero" api:"required"`
	Name       string           `json:"name" api:"required"`
	// Type-specific JSON configuration.
	Parameters map[string]any `json:"parameters,omitzero" api:"required"`
	// Parent tag manager that will own the new trigger.
	TagManagerID string `json:"tagManagerId" api:"required"`
	// Must equal `type` — send the same string in both fields. The server rejects any
	// divergent value.
	Trigger string `json:"Trigger" api:"required"`
	// Trigger type discriminator. Pick from `GET /tag-manager-triggers/types` for the
	// canonical set (e.g. `PageView`, `CustomEvent`, `AllElementsClick`).
	Type    string          `json:"type" api:"required"`
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r TagManagerTriggerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TagManagerTriggerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagManagerTriggerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerTriggerUpdateParams struct {
	// Pause/resume the trigger without changing other fields.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Updated trigger name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Must equal `type`. Omit both fields, or send both with the same value — the
	// server rejects any divergence.
	Trigger param.Opt[string] `json:"Trigger,omitzero"`
	// Updated trigger type. Pick from `GET /tag-manager-triggers/types`. When changing
	// `type`, send the new value in `Trigger` as well (they must match).
	Type param.Opt[string] `json:"type,omitzero"`
	// Replaces conditions wholesale when sent. Use `[]` for an unconditional trigger.
	Conditions []map[string]any `json:"conditions,omitzero"`
	// Updated type-specific JSON configuration.
	Parameters map[string]any `json:"parameters,omitzero"`
	paramObj
}

func (r TagManagerTriggerUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TagManagerTriggerUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagManagerTriggerUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
