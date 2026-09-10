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

// TagManagerFolderService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTagManagerFolderService] method instead.
type TagManagerFolderService struct {
	Options []option.RequestOption
}

// NewTagManagerFolderService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTagManagerFolderService(opts ...option.RequestOption) (r TagManagerFolderService) {
	r = TagManagerFolderService{}
	r.Options = opts
	return
}

// List folders inside a single tag manager. Folders are dashboard-only
// organizational containers — they do not affect tag evaluation. Requires the
// `tagManagerId` query parameter. Supports cursor pagination via `limit` and
// `cursor`; the limit clamp is 1000 so a single request can return the full set
// (the web-app workspace renders all folders in one shot). Requires scope:
// tagManagers:find
func (r *TagManagerFolderService) List(ctx context.Context, query TagManagerFolderListParams, opts ...option.RequestOption) (res *pagination.Cursor[TagManagerFolderListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "rest/v1/tag-manager-folders"
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

// List folders inside a single tag manager. Folders are dashboard-only
// organizational containers — they do not affect tag evaluation. Requires the
// `tagManagerId` query parameter. Supports cursor pagination via `limit` and
// `cursor`; the limit clamp is 1000 so a single request can return the full set
// (the web-app workspace renders all folders in one shot). Requires scope:
// tagManagers:find
func (r *TagManagerFolderService) ListAutoPaging(ctx context.Context, query TagManagerFolderListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[TagManagerFolderListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Create a folder inside a tag manager. `tagManagerId` is required in the body.
// Names are case-insensitively unique within the tag manager — collisions return
// 409 with the reason in the response `error` field. Requires scope:
// tagManagers:update
func (r *TagManagerFolderService) New(ctx context.Context, body TagManagerFolderNewParams, opts ...option.RequestOption) (res *TagManagerFolderNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/tag-manager-folders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Find a single tag manager folder by ID. Requires scope: tagManagers:find
func (r *TagManagerFolderService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *TagManagerFolderGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/tag-manager-folders/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Rename a folder. The new name must be case-insensitively unique within the tag
// manager; collisions return 409. Requires scope: tagManagers:update
func (r *TagManagerFolderService) Update(ctx context.Context, id string, body TagManagerFolderUpdateParams, opts ...option.RequestOption) (res *TagManagerFolderUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/tag-manager-folders/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete a folder. Tags, triggers, and variables previously assigned to the folder
// are no longer grouped under it; the assets themselves are not deleted. Requires
// scope: tagManagers:update
func (r *TagManagerFolderService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *TagManagerFolderDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/tag-manager-folders/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type TagManagerFolderListResponse struct {
	ID           string `json:"id" api:"required"`
	AccountID    string `json:"accountId" api:"required"`
	Name         string `json:"name" api:"required"`
	TagManagerID string `json:"tagManagerId" api:"required"`
	CreatedAt    string `json:"createdAt" api:"nullable"`
	UpdatedAt    string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AccountID    respjson.Field
		Name         respjson.Field
		TagManagerID respjson.Field
		CreatedAt    respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagManagerFolderListResponse) RawJSON() string { return r.JSON.raw }
func (r *TagManagerFolderListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerFolderNewResponse struct {
	ID           string `json:"id" api:"required"`
	AccountID    string `json:"accountId" api:"required"`
	Name         string `json:"name" api:"required"`
	TagManagerID string `json:"tagManagerId" api:"required"`
	CreatedAt    string `json:"createdAt" api:"nullable"`
	UpdatedAt    string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AccountID    respjson.Field
		Name         respjson.Field
		TagManagerID respjson.Field
		CreatedAt    respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagManagerFolderNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TagManagerFolderNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerFolderGetResponse struct {
	ID           string `json:"id" api:"required"`
	AccountID    string `json:"accountId" api:"required"`
	Name         string `json:"name" api:"required"`
	TagManagerID string `json:"tagManagerId" api:"required"`
	CreatedAt    string `json:"createdAt" api:"nullable"`
	UpdatedAt    string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AccountID    respjson.Field
		Name         respjson.Field
		TagManagerID respjson.Field
		CreatedAt    respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagManagerFolderGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TagManagerFolderGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerFolderUpdateResponse struct {
	ID           string `json:"id" api:"required"`
	AccountID    string `json:"accountId" api:"required"`
	Name         string `json:"name" api:"required"`
	TagManagerID string `json:"tagManagerId" api:"required"`
	CreatedAt    string `json:"createdAt" api:"nullable"`
	UpdatedAt    string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AccountID    respjson.Field
		Name         respjson.Field
		TagManagerID respjson.Field
		CreatedAt    respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagManagerFolderUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *TagManagerFolderUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerFolderDeleteResponse struct {
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
func (r TagManagerFolderDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *TagManagerFolderDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerFolderListParams struct {
	// Parent tag manager whose folders should be returned.
	TagManagerID string `query:"tagManagerId" api:"required" json:"-"`
	// Maximum number of folders to return. Defaults to 25; values below 1 are clamped
	// to 1 and values above 1000 are clamped to 1000. The web-app passes 1000 to
	// render the full workspace in one request.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Opaque pagination cursor from pagination.nextCursor in the previous response. Do
	// not decode or modify it. Malformed cursors return 400 Bad Request.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TagManagerFolderListParams]'s query parameters as
// `url.Values`.
func (r TagManagerFolderListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TagManagerFolderNewParams struct {
	// Folder name. Case-insensitively unique within the tag manager.
	Name string `json:"name" api:"required"`
	// Parent tag manager that will own the new folder.
	TagManagerID string `json:"tagManagerId" api:"required"`
	paramObj
}

func (r TagManagerFolderNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TagManagerFolderNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagManagerFolderNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagManagerFolderUpdateParams struct {
	// New folder name. Case-insensitively unique within the tag manager.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r TagManagerFolderUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TagManagerFolderUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagManagerFolderUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
