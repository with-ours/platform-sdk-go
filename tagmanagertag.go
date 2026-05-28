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

// TagManagerTagService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTagManagerTagService] method instead.
type TagManagerTagService struct {
	Options []option.RequestOption
}

// NewTagManagerTagService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTagManagerTagService(opts ...option.RequestOption) (r TagManagerTagService) {
	r = TagManagerTagService{}
	r.Options = opts
	return
}

// List tags inside a single tag manager. Requires the `tagManagerId` query
// parameter — tags are always scoped to one parent container. Supports cursor
// pagination via `limit` and `cursor`; the limit clamp is 1000 so a single request
// can return the full set (the web-app workspace renders all tags in one shot).
// Requires scope: tagManagers:find
func (r *TagManagerTagService) List(ctx context.Context, query TagManagerTagListParams, opts ...option.RequestOption) (res *pagination.Cursor[TagManagerTagListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "rest/v1/tag-manager-tags"
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

// List tags inside a single tag manager. Requires the `tagManagerId` query
// parameter — tags are always scoped to one parent container. Supports cursor
// pagination via `limit` and `cursor`; the limit clamp is 1000 so a single request
// can return the full set (the web-app workspace renders all tags in one shot).
// Requires scope: tagManagers:find
func (r *TagManagerTagService) ListAutoPaging(ctx context.Context, query TagManagerTagListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[TagManagerTagListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Create a new tag inside a tag manager. `tagManagerId` is required in the body.
// Newly created tags are not assigned to any folder — assign after creation via
// PATCH with `folderId`. Requires scope: tagManagers:update
func (r *TagManagerTagService) New(ctx context.Context, body TagManagerTagNewParams, opts ...option.RequestOption) (res *TagManagerTagNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/tag-manager-tags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Fetch a single tag by id, including its `folderId` (read-only on this endpoint).
// Requires scope: tagManagers:find
func (r *TagManagerTagService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *TagManagerTagGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/tag-manager-tags/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Partially update a tag. Only the fields you send are changed. Tags cannot be
// moved between tag managers. To assign a tag to a folder, use
// `POST /rest/v1/tag-manager-asset-folders`. Requires scope: tagManagers:update
func (r *TagManagerTagService) Update(ctx context.Context, id string, body TagManagerTagUpdateParams, opts ...option.RequestOption) (res *TagManagerTagUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/tag-manager-tags/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete a tag manager tag. Requires scope: tagManagers:update
func (r *TagManagerTagService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *TagManagerTagDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/tag-manager-tags/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Lists every tag template the platform supports — what `type` discriminator to
// send on create/patch, and the shape of the type-specific `parameters` payload
// (fields, validators, required flags, available values for selects).
// Account-agnostic: the response is the same for every API key. The same registry
// powers server-side validation on `POST` / `PATCH` so what this endpoint
// advertises matches what the server enforces. Requires scope: tagManagers:find
func (r *TagManagerTagService) Types(ctx context.Context, opts ...option.RequestOption) (res *TagManagerTagTypesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/tag-manager-tags/types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type TagManagerTagListResponse struct {
	// Server-assigned UUID for this tag.
	ID        string `json:"id" api:"required"`
	AccountID string `json:"accountId" api:"required"`
	// Triggers that fire this tag. Tag does nothing if this list is empty.
	FireTriggerIDs []string `json:"fireTriggerIds" api:"required"`
	// Human-readable tag name.
	Name string `json:"name" api:"required"`
	// Type-specific JSON configuration. Shape depends on `type` — inspect a known-good
	// tag of the same type for the field set. Empty object is valid for placeholder
	// tags.
	Parameters map[string]any `json:"parameters" api:"required"`
	// Parent tag manager that owns this tag.
	TagManagerID string `json:"tagManagerId" api:"required"`
	// Tag type discriminator. Examples that exist today: `OursTrackTag`,
	// `OursInitTag`, `OursIdentifyTag`, `CustomHtmlTag`. Pick from
	// `GET /tag-manager-tags/types` for the canonical set — names like `GA4Event` are
	// not valid ids.
	Type string `json:"type" api:"required"`
	// Triggers that suppress this tag when they match — evaluated after fire triggers.
	BlockTriggerIDs []string `json:"blockTriggerIds" api:"nullable"`
	CreatedAt       string   `json:"createdAt" api:"nullable"`
	// Defaults to `true` on create.
	Enabled bool `json:"enabled" api:"nullable"`
	// Folder this tag belongs to in the dashboard. Settable via PATCH — send a folder
	// UUID to assign, or `null` to remove from its current folder.
	FolderID string `json:"folderId" api:"nullable"`
	// Lower numbers fire first. Defaults to 0.
	Priority  float64 `json:"priority" api:"nullable"`
	UpdatedAt string  `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AccountID       respjson.Field
		FireTriggerIDs  respjson.Field
		Name            respjson.Field
		Parameters      respjson.Field
		TagManagerID    respjson.Field
		Type            respjson.Field
		BlockTriggerIDs respjson.Field
		CreatedAt       respjson.Field
		Enabled         respjson.Field
		FolderID        respjson.Field
		Priority        respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagManagerTagListResponse) RawJSON() string { return r.JSON.raw }
func (r *TagManagerTagListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerTagNewResponse struct {
	// Server-assigned UUID for this tag.
	ID        string `json:"id" api:"required"`
	AccountID string `json:"accountId" api:"required"`
	// Triggers that fire this tag. Tag does nothing if this list is empty.
	FireTriggerIDs []string `json:"fireTriggerIds" api:"required"`
	// Human-readable tag name.
	Name string `json:"name" api:"required"`
	// Type-specific JSON configuration. Shape depends on `type` — inspect a known-good
	// tag of the same type for the field set. Empty object is valid for placeholder
	// tags.
	Parameters map[string]any `json:"parameters" api:"required"`
	// Parent tag manager that owns this tag.
	TagManagerID string `json:"tagManagerId" api:"required"`
	// Tag type discriminator. Examples that exist today: `OursTrackTag`,
	// `OursInitTag`, `OursIdentifyTag`, `CustomHtmlTag`. Pick from
	// `GET /tag-manager-tags/types` for the canonical set — names like `GA4Event` are
	// not valid ids.
	Type string `json:"type" api:"required"`
	// Triggers that suppress this tag when they match — evaluated after fire triggers.
	BlockTriggerIDs []string `json:"blockTriggerIds" api:"nullable"`
	CreatedAt       string   `json:"createdAt" api:"nullable"`
	// Defaults to `true` on create.
	Enabled bool `json:"enabled" api:"nullable"`
	// Folder this tag belongs to in the dashboard. Settable via PATCH — send a folder
	// UUID to assign, or `null` to remove from its current folder.
	FolderID string `json:"folderId" api:"nullable"`
	// Lower numbers fire first. Defaults to 0.
	Priority  float64 `json:"priority" api:"nullable"`
	UpdatedAt string  `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AccountID       respjson.Field
		FireTriggerIDs  respjson.Field
		Name            respjson.Field
		Parameters      respjson.Field
		TagManagerID    respjson.Field
		Type            respjson.Field
		BlockTriggerIDs respjson.Field
		CreatedAt       respjson.Field
		Enabled         respjson.Field
		FolderID        respjson.Field
		Priority        respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagManagerTagNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TagManagerTagNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerTagGetResponse struct {
	// Server-assigned UUID for this tag.
	ID        string `json:"id" api:"required"`
	AccountID string `json:"accountId" api:"required"`
	// Triggers that fire this tag. Tag does nothing if this list is empty.
	FireTriggerIDs []string `json:"fireTriggerIds" api:"required"`
	// Human-readable tag name.
	Name string `json:"name" api:"required"`
	// Type-specific JSON configuration. Shape depends on `type` — inspect a known-good
	// tag of the same type for the field set. Empty object is valid for placeholder
	// tags.
	Parameters map[string]any `json:"parameters" api:"required"`
	// Parent tag manager that owns this tag.
	TagManagerID string `json:"tagManagerId" api:"required"`
	// Tag type discriminator. Examples that exist today: `OursTrackTag`,
	// `OursInitTag`, `OursIdentifyTag`, `CustomHtmlTag`. Pick from
	// `GET /tag-manager-tags/types` for the canonical set — names like `GA4Event` are
	// not valid ids.
	Type string `json:"type" api:"required"`
	// Triggers that suppress this tag when they match — evaluated after fire triggers.
	BlockTriggerIDs []string `json:"blockTriggerIds" api:"nullable"`
	CreatedAt       string   `json:"createdAt" api:"nullable"`
	// Defaults to `true` on create.
	Enabled bool `json:"enabled" api:"nullable"`
	// Folder this tag belongs to in the dashboard. Settable via PATCH — send a folder
	// UUID to assign, or `null` to remove from its current folder.
	FolderID string `json:"folderId" api:"nullable"`
	// Lower numbers fire first. Defaults to 0.
	Priority  float64 `json:"priority" api:"nullable"`
	UpdatedAt string  `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AccountID       respjson.Field
		FireTriggerIDs  respjson.Field
		Name            respjson.Field
		Parameters      respjson.Field
		TagManagerID    respjson.Field
		Type            respjson.Field
		BlockTriggerIDs respjson.Field
		CreatedAt       respjson.Field
		Enabled         respjson.Field
		FolderID        respjson.Field
		Priority        respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagManagerTagGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TagManagerTagGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerTagUpdateResponse struct {
	// Server-assigned UUID for this tag.
	ID        string `json:"id" api:"required"`
	AccountID string `json:"accountId" api:"required"`
	// Triggers that fire this tag. Tag does nothing if this list is empty.
	FireTriggerIDs []string `json:"fireTriggerIds" api:"required"`
	// Human-readable tag name.
	Name string `json:"name" api:"required"`
	// Type-specific JSON configuration. Shape depends on `type` — inspect a known-good
	// tag of the same type for the field set. Empty object is valid for placeholder
	// tags.
	Parameters map[string]any `json:"parameters" api:"required"`
	// Parent tag manager that owns this tag.
	TagManagerID string `json:"tagManagerId" api:"required"`
	// Tag type discriminator. Examples that exist today: `OursTrackTag`,
	// `OursInitTag`, `OursIdentifyTag`, `CustomHtmlTag`. Pick from
	// `GET /tag-manager-tags/types` for the canonical set — names like `GA4Event` are
	// not valid ids.
	Type string `json:"type" api:"required"`
	// Triggers that suppress this tag when they match — evaluated after fire triggers.
	BlockTriggerIDs []string `json:"blockTriggerIds" api:"nullable"`
	CreatedAt       string   `json:"createdAt" api:"nullable"`
	// Defaults to `true` on create.
	Enabled bool `json:"enabled" api:"nullable"`
	// Folder this tag belongs to in the dashboard. Settable via PATCH — send a folder
	// UUID to assign, or `null` to remove from its current folder.
	FolderID string `json:"folderId" api:"nullable"`
	// Lower numbers fire first. Defaults to 0.
	Priority  float64 `json:"priority" api:"nullable"`
	UpdatedAt string  `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AccountID       respjson.Field
		FireTriggerIDs  respjson.Field
		Name            respjson.Field
		Parameters      respjson.Field
		TagManagerID    respjson.Field
		Type            respjson.Field
		BlockTriggerIDs respjson.Field
		CreatedAt       respjson.Field
		Enabled         respjson.Field
		FolderID        respjson.Field
		Priority        respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagManagerTagUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *TagManagerTagUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerTagDeleteResponse struct {
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
func (r TagManagerTagDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *TagManagerTagDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerTagTypesResponse struct {
	Entities []TagManagerTagTypesResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagManagerTagTypesResponse) RawJSON() string { return r.JSON.raw }
func (r *TagManagerTagTypesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerTagTypesResponseEntity struct {
	// Type discriminator — pass this as `type` on create/patch.
	ID string `json:"id" api:"required"`
	// Grouping label.
	Category string                                  `json:"category" api:"required"`
	Fields   []TagManagerTagTypesResponseEntityField `json:"fields" api:"required"`
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
func (r TagManagerTagTypesResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *TagManagerTagTypesResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerTagTypesResponseEntityField struct {
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
	AvailableValues []TagManagerTagTypesResponseEntityFieldAvailableValue `json:"availableValues" api:"nullable"`
	// Default value when the caller omits the parameter on create.
	Default     map[string]any `json:"default"`
	Description string         `json:"description" api:"nullable"`
	// When `true`, omitting or sending an empty value for this parameter on
	// create/patch returns HTTP 400.
	Required bool `json:"required" api:"nullable"`
	// Server-enforced rules applied to this field at create and patch.
	Validators []TagManagerTagTypesResponseEntityFieldValidator `json:"validators" api:"nullable"`
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
func (r TagManagerTagTypesResponseEntityField) RawJSON() string { return r.JSON.raw }
func (r *TagManagerTagTypesResponseEntityField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerTagTypesResponseEntityFieldAvailableValue struct {
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
func (r TagManagerTagTypesResponseEntityFieldAvailableValue) RawJSON() string { return r.JSON.raw }
func (r *TagManagerTagTypesResponseEntityFieldAvailableValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerTagTypesResponseEntityFieldValidator struct {
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
func (r TagManagerTagTypesResponseEntityFieldValidator) RawJSON() string { return r.JSON.raw }
func (r *TagManagerTagTypesResponseEntityFieldValidator) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerTagListParams struct {
	// Parent tag manager whose tags should be returned.
	TagManagerID string `query:"tagManagerId" api:"required" json:"-"`
	// Maximum number of tags to return. Defaults to 25; values below 1 are clamped to
	// 1 and values above 1000 are clamped to 1000. The web-app passes 1000 to render
	// the full workspace in one request.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Opaque pagination cursor from pagination.nextCursor in the previous response. Do
	// not decode or modify it. Malformed cursors return 400 Bad Request.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TagManagerTagListParams]'s query parameters as
// `url.Values`.
func (r TagManagerTagListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TagManagerTagNewParams struct {
	// Trigger ids that cause this tag to fire. Use `[]` only for placeholder tags.
	FireTriggerIDs []string `json:"fireTriggerIds,omitzero" api:"required"`
	// Human-readable tag name.
	Name string `json:"name" api:"required"`
	// Type-specific JSON configuration. Send `{}` for a placeholder.
	Parameters map[string]any `json:"parameters,omitzero" api:"required"`
	// Parent tag manager that will own the new tag.
	TagManagerID string `json:"tagManagerId" api:"required"`
	// Tag type discriminator. Pick from `GET /tag-manager-tags/types` for the
	// canonical set (e.g. `OursTrackTag`, `OursInitTag`, `CustomHtmlTag`). Names like
	// `GA4Event` are not valid ids.
	Type string `json:"type" api:"required"`
	// Defaults to `true`.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Defaults to 0.
	Priority param.Opt[float64] `json:"priority,omitzero"`
	// Optional trigger ids that block this tag when they match.
	BlockTriggerIDs []string `json:"blockTriggerIds,omitzero"`
	paramObj
}

func (r TagManagerTagNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TagManagerTagNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagManagerTagNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerTagUpdateParams struct {
	// Pause/resume the tag without changing other fields.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Updated priority.
	Priority param.Opt[float64] `json:"priority,omitzero"`
	// Updated tag name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Updated tag type. Pick from `GET /tag-manager-tags/types`.
	Type param.Opt[string] `json:"type,omitzero"`
	// Replaces the block trigger list wholesale. Send `null` to clear.
	BlockTriggerIDs []string `json:"blockTriggerIds,omitzero"`
	// Replaces the fire trigger list wholesale.
	FireTriggerIDs []string `json:"fireTriggerIds,omitzero"`
	// Updated type-specific JSON configuration.
	Parameters map[string]any `json:"parameters,omitzero"`
	paramObj
}

func (r TagManagerTagUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TagManagerTagUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagManagerTagUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
