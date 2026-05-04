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
	"github.com/with-ours/platform-sdk-go/internal/apiquery"
	"github.com/with-ours/platform-sdk-go/internal/requestconfig"
	"github.com/with-ours/platform-sdk-go/option"
	"github.com/with-ours/platform-sdk-go/packages/param"
	"github.com/with-ours/platform-sdk-go/packages/respjson"
)

// VersionService contains methods and other services that help with interacting
// with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVersionService] method instead.
type VersionService struct {
	Options []option.RequestOption
}

// NewVersionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVersionService(opts ...option.RequestOption) (r VersionService) {
	r = VersionService{}
	r.Options = opts
	return
}

// Publish the current draft (i.e. all unpublished entity changes) as a new
// version. Returns the full Version on success. Returns HTTP 409 with the reason
// in the response `error` field when there are no draft changes to publish, when
// another publish is already in flight, or when the action otherwise conflicts
// with current state. To re-publish an existing version, use POST
// /rest/v1/versions/{id}/publish instead. Requires scope: version:publish
func (r *VersionService) New(ctx context.Context, body VersionNewParams, opts ...option.RequestOption) (res *VersionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/versions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Find a single version by ID. Requires scope: version:find
func (r *VersionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *VersionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/versions/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Partially update a version. Only the fields you send are changed. Requires
// scope: version:update
func (r *VersionService) Update(ctx context.Context, id string, body VersionUpdateParams, opts ...option.RequestOption) (res *VersionUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/versions/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List versions for this account, newest first. Supports cursor pagination and
// filtering by `isPublished`, `nameContains`, and `notesContains`. Combine filters
// with AND semantics. Requires scope: version:list
func (r *VersionService) List(ctx context.Context, query VersionListParams, opts ...option.RequestOption) (res *VersionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/versions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type VersionNewResponse struct {
	ID            string  `json:"id" api:"required"`
	CreatedAt     string  `json:"createdAt" api:"required"`
	IsPublished   bool    `json:"isPublished" api:"required"`
	VersionNumber float64 `json:"versionNumber" api:"required"`
	Name          string  `json:"name" api:"nullable"`
	Notes         string  `json:"notes" api:"nullable"`
	// When this version was most recently published. NOT cleared when a newer version
	// is published — `publishedAt` reflects the most recent successful publish of this
	// row, regardless of whether `isPublished` is currently true. Use `isPublished` to
	// determine the current live version.
	PublishedAt string `json:"publishedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		IsPublished   respjson.Field
		VersionNumber respjson.Field
		Name          respjson.Field
		Notes         respjson.Field
		PublishedAt   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *VersionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionGetResponse struct {
	ID            string  `json:"id" api:"required"`
	CreatedAt     string  `json:"createdAt" api:"required"`
	IsPublished   bool    `json:"isPublished" api:"required"`
	VersionNumber float64 `json:"versionNumber" api:"required"`
	Name          string  `json:"name" api:"nullable"`
	Notes         string  `json:"notes" api:"nullable"`
	// When this version was most recently published. NOT cleared when a newer version
	// is published — `publishedAt` reflects the most recent successful publish of this
	// row, regardless of whether `isPublished` is currently true. Use `isPublished` to
	// determine the current live version.
	PublishedAt string `json:"publishedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		IsPublished   respjson.Field
		VersionNumber respjson.Field
		Name          respjson.Field
		Notes         respjson.Field
		PublishedAt   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *VersionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionUpdateResponse struct {
	ID            string  `json:"id" api:"required"`
	CreatedAt     string  `json:"createdAt" api:"required"`
	IsPublished   bool    `json:"isPublished" api:"required"`
	VersionNumber float64 `json:"versionNumber" api:"required"`
	Name          string  `json:"name" api:"nullable"`
	Notes         string  `json:"notes" api:"nullable"`
	// When this version was most recently published. NOT cleared when a newer version
	// is published — `publishedAt` reflects the most recent successful publish of this
	// row, regardless of whether `isPublished` is currently true. Use `isPublished` to
	// determine the current live version.
	PublishedAt string `json:"publishedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		IsPublished   respjson.Field
		VersionNumber respjson.Field
		Name          respjson.Field
		Notes         respjson.Field
		PublishedAt   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *VersionUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionListResponse struct {
	Entities   []VersionListResponseEntity   `json:"entities" api:"required"`
	Pagination VersionListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionListResponse) RawJSON() string { return r.JSON.raw }
func (r *VersionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionListResponseEntity struct {
	ID            string  `json:"id" api:"required"`
	CreatedAt     string  `json:"createdAt" api:"required"`
	IsPublished   bool    `json:"isPublished" api:"required"`
	VersionNumber float64 `json:"versionNumber" api:"required"`
	Name          string  `json:"name" api:"nullable"`
	// When this version was most recently published. NOT cleared when a newer version
	// is published — `publishedAt` reflects the most recent successful publish of this
	// row, regardless of whether `isPublished` is currently true. Use `isPublished` to
	// determine the current live version.
	PublishedAt string `json:"publishedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		IsPublished   respjson.Field
		VersionNumber respjson.Field
		Name          respjson.Field
		PublishedAt   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionListResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *VersionListResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionListResponsePagination struct {
	HasMore    bool   `json:"hasMore" api:"required"`
	NextCursor string `json:"nextCursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *VersionListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionNewParams struct {
	Name                            param.Opt[string] `json:"name,omitzero"`
	Notes                           param.Opt[string] `json:"notes,omitzero"`
	IncludeAllowedEvents            []string          `json:"includeAllowedEvents,omitzero"`
	IncludeConsentSettings          []string          `json:"includeConsentSettings,omitzero"`
	IncludeDestinations             []string          `json:"includeDestinations,omitzero"`
	IncludeExternalAllowedEventData []string          `json:"includeExternalAllowedEventData,omitzero"`
	IncludeGlobalDispatchCenters    []string          `json:"includeGlobalDispatchCenters,omitzero"`
	IncludeMappings                 []string          `json:"includeMappings,omitzero"`
	IncludeReplaySettings           []string          `json:"includeReplaySettings,omitzero"`
	IncludeSources                  []string          `json:"includeSources,omitzero"`
	IncludeTagManagerTags           []string          `json:"includeTagManagerTags,omitzero"`
	IncludeTagManagerTriggers       []string          `json:"includeTagManagerTriggers,omitzero"`
	IncludeTagManagerVariables      []string          `json:"includeTagManagerVariables,omitzero"`
	paramObj
}

func (r VersionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VersionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VersionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionUpdateParams struct {
	Name  param.Opt[string] `json:"name,omitzero"`
	Notes param.Opt[string] `json:"notes,omitzero"`
	paramObj
}

func (r VersionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow VersionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VersionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionListParams struct {
	// Maximum number of versions to return. Defaults to 25; values below 1 are clamped
	// to 1 and values above 100 are clamped to 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Opaque pagination cursor from pagination.nextCursor in the previous response. Do
	// not decode or modify it. Malformed cursors return 400 Bad Request.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Case-insensitive substring match on the version name.
	NameContains param.Opt[string] `query:"nameContains,omitzero" json:"-"`
	// Case-insensitive substring match on the version notes.
	NotesContains param.Opt[string] `query:"notesContains,omitzero" json:"-"`
	// Filter to only published or unpublished versions.
	//
	// Any of "true", "false".
	IsPublished VersionListParamsIsPublished `query:"isPublished,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VersionListParams]'s query parameters as `url.Values`.
func (r VersionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter to only published or unpublished versions.
type VersionListParamsIsPublished string

const (
	VersionListParamsIsPublishedTrue  VersionListParamsIsPublished = "true"
	VersionListParamsIsPublishedFalse VersionListParamsIsPublished = "false"
)
