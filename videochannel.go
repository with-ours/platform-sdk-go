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

// VideoChannelService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVideoChannelService] method instead.
type VideoChannelService struct {
	Options []option.RequestOption
}

// NewVideoChannelService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVideoChannelService(opts ...option.RequestOption) (r VideoChannelService) {
	r = VideoChannelService{}
	r.Options = opts
	return
}

// List video channels for the account, sorted by name. Supports cursor pagination
// via `limit` and `cursor`; the limit clamp is 1000 so a single request can return
// the full set. Entries omit `resolvedValues` — fetch a channel by id for its video
// count and embed output. Requires scope: videoChannel:list
func (r *VideoChannelService) List(ctx context.Context, query VideoChannelListParams, opts ...option.RequestOption) (res *pagination.Cursor[VideoChannelListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "rest/v1/video-channels"
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

// List video channels for the account, sorted by name. Supports cursor pagination
// via `limit` and `cursor`; the limit clamp is 1000 so a single request can return
// the full set. Entries omit `resolvedValues` — fetch a channel by id for its video
// count and embed output. Requires scope: videoChannel:list
func (r *VideoChannelService) ListAutoPaging(ctx context.Context, query VideoChannelListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[VideoChannelListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Create a video channel. Only `name` is accepted here; set branding and publish it
// with PATCH, and add videos with `POST /rest/v1/video-channels/{id}/media`. New
// channels start unpublished, so the page is not reachable until you send
// `isPublished: true`. Requires scope: videoChannel:create
func (r *VideoChannelService) New(ctx context.Context, body VideoChannelNewParams, opts ...option.RequestOption) (res *VideoChannelNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/video-channels"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Fetch a channel with its branding, publish state, video count, shareable page
// URL, and paste-ready embed code. Requires scope: videoChannel:find
func (r *VideoChannelService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *VideoChannelGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/video-channels/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Partially update a channel. Only fields included in the body change; send `null`
// to clear a nullable field. Sending `isPublished: true` makes the channel page
// reachable and renders it from the current videos and branding; `false` takes it
// offline. `logoMediaId` must reference an image in your media library. Requires
// scope: videoChannel:update
func (r *VideoChannelService) Update(ctx context.Context, id string, body VideoChannelUpdateParams, opts ...option.RequestOption) (res *VideoChannelUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/video-channels/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete a channel and take its page offline. The videos it listed are not deleted
// — only their membership in this channel. Requires scope: videoChannel:delete
func (r *VideoChannelService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *VideoChannelDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/video-channels/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// List the videos in a channel, ordered by their position on the page. Not
// paginated: a channel holds a bounded set of videos, so the full ordered list is
// always returned. Videos whose media no longer resolves are omitted rather than
// returned as broken entries. Requires scope: videoChannel:find
func (r *VideoChannelService) Media(ctx context.Context, id string, opts ...option.RequestOption) (res *VideoChannelMediaResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/video-channels/%s/media", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Add a video to a channel. Omit `position` to append it to the end. A video can
// belong to several channels, so adding it here does not remove it from any other.
// Calling this again for a video already in the channel updates its position instead
// of adding a duplicate, and keeps its current slot when `position` is omitted. The
// returned `id` is a composite membership key, not a UUID. Requires scope:
// videoChannel:update
func (r *VideoChannelService) AssignMedia(ctx context.Context, id string, body VideoChannelAssignMediaParams, opts ...option.RequestOption) (res *VideoChannelAssignMediaResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/video-channels/%s/media", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Remove one video from a channel, identified by the `mediaId` query parameter. The
// video itself is not deleted and stays in any other channel it belongs to.
// Idempotent — removing a video that is not in the channel succeeds and returns the
// channel unchanged. Requires scope: videoChannel:update
func (r *VideoChannelService) RemoveMedia(ctx context.Context, id string, query VideoChannelRemoveMediaParams, opts ...option.RequestOption) (res *VideoChannelRemoveMediaResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/video-channels/%s/media", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, query, &res, opts...)
	return res, err
}

// Set the display order of a channel’s videos. Send every video id currently in the
// channel in the order you want them shown — index 0 appears first. A partial list,
// or an id that is not in the channel, returns 400 so a caller working from a stale
// view learns it is out of date instead of getting a partial write. Requires scope:
// videoChannel:update
func (r *VideoChannelService) Reorder(ctx context.Context, id string, body VideoChannelReorderParams, opts ...option.RequestOption) (res *VideoChannelReorderResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/video-channels/%s/reorder", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type VideoChannelListResponse struct {
	ID          string `json:"id" api:"required"`
	AccountID   string `json:"accountId" api:"required"`
	Name        string `json:"name" api:"required"`
	BrandColor  string `json:"brandColor" api:"nullable"`
	CreatedAt   string `json:"createdAt" api:"nullable"`
	Description string `json:"description" api:"nullable"`
	FooterText  string `json:"footerText" api:"nullable"`
	IsPublished bool   `json:"isPublished" api:"nullable"`
	LogoMediaID string `json:"logoMediaId" api:"nullable"`
	UpdatedAt   string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AccountID   respjson.Field
		Name        respjson.Field
		BrandColor  respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		FooterText  respjson.Field
		IsPublished respjson.Field
		LogoMediaID respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoChannelListResponse) RawJSON() string { return r.JSON.raw }
func (r *VideoChannelListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoChannelNewResponse struct {
	ID          string `json:"id" api:"required"`
	AccountID   string `json:"accountId" api:"required"`
	Name        string `json:"name" api:"required"`
	BrandColor  string `json:"brandColor" api:"nullable"`
	CreatedAt   string `json:"createdAt" api:"nullable"`
	Description string `json:"description" api:"nullable"`
	FooterText  string `json:"footerText" api:"nullable"`
	IsPublished bool   `json:"isPublished" api:"nullable"`
	LogoMediaID string `json:"logoMediaId" api:"nullable"`
	UpdatedAt   string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AccountID   respjson.Field
		Name        respjson.Field
		BrandColor  respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		FooterText  respjson.Field
		IsPublished respjson.Field
		LogoMediaID respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoChannelNewResponse) RawJSON() string { return r.JSON.raw }
func (r *VideoChannelNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoChannelGetResponse struct {
	ID             string `json:"id" api:"required"`
	AccountID      string `json:"accountId" api:"required"`
	Name           string `json:"name" api:"required"`
	BrandColor     string `json:"brandColor" api:"nullable"`
	CreatedAt      string `json:"createdAt" api:"nullable"`
	Description    string `json:"description" api:"nullable"`
	FooterText     string `json:"footerText" api:"nullable"`
	IsPublished    bool   `json:"isPublished" api:"nullable"`
	LogoMediaID    string `json:"logoMediaId" api:"nullable"`
	ResolvedValues any    `json:"resolvedValues" api:"nullable"`
	UpdatedAt      string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		AccountID      respjson.Field
		Name           respjson.Field
		BrandColor     respjson.Field
		CreatedAt      respjson.Field
		Description    respjson.Field
		FooterText     respjson.Field
		IsPublished    respjson.Field
		LogoMediaID    respjson.Field
		ResolvedValues respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoChannelGetResponse) RawJSON() string { return r.JSON.raw }
func (r *VideoChannelGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoChannelUpdateResponse struct {
	ID             string `json:"id" api:"required"`
	AccountID      string `json:"accountId" api:"required"`
	Name           string `json:"name" api:"required"`
	BrandColor     string `json:"brandColor" api:"nullable"`
	CreatedAt      string `json:"createdAt" api:"nullable"`
	Description    string `json:"description" api:"nullable"`
	FooterText     string `json:"footerText" api:"nullable"`
	IsPublished    bool   `json:"isPublished" api:"nullable"`
	LogoMediaID    string `json:"logoMediaId" api:"nullable"`
	ResolvedValues any    `json:"resolvedValues" api:"nullable"`
	UpdatedAt      string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		AccountID      respjson.Field
		Name           respjson.Field
		BrandColor     respjson.Field
		CreatedAt      respjson.Field
		Description    respjson.Field
		FooterText     respjson.Field
		IsPublished    respjson.Field
		LogoMediaID    respjson.Field
		ResolvedValues respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoChannelUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *VideoChannelUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoChannelDeleteResponse struct {
	ID string `json:"id" api:"required"`
	// Any of true.
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
func (r VideoChannelDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *VideoChannelDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoChannelMediaResponse struct {
	Entities []VideoChannelMediaResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoChannelMediaResponse) RawJSON() string { return r.JSON.raw }
func (r *VideoChannelMediaResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoChannelMediaResponseEntity struct {
	ID        string `json:"id" api:"required"`
	AccountID string `json:"accountId" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	// Any of "Video".
	Type                  string  `json:"type" api:"required"`
	CaptionsUpdatedAt     string  `json:"captionsUpdatedAt" api:"nullable"`
	CaptionsUpdatedByName string  `json:"captionsUpdatedByName" api:"nullable"`
	Description           string  `json:"description" api:"nullable"`
	Duration              float64 `json:"duration" api:"nullable"`
	HasVideoUpload        bool    `json:"hasVideoUpload" api:"nullable"`
	Height                float64 `json:"height" api:"nullable"`
	Name                  string  `json:"name" api:"nullable"`
	UpdatedAt             string  `json:"updatedAt" api:"nullable"`
	Width                 float64 `json:"width" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		AccountID             respjson.Field
		CreatedAt             respjson.Field
		Type                  respjson.Field
		CaptionsUpdatedAt     respjson.Field
		CaptionsUpdatedByName respjson.Field
		Description           respjson.Field
		Duration              respjson.Field
		HasVideoUpload        respjson.Field
		Height                respjson.Field
		Name                  respjson.Field
		UpdatedAt             respjson.Field
		Width                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoChannelMediaResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *VideoChannelMediaResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoChannelAssignMediaResponse struct {
	ID        string  `json:"id" api:"required"`
	AccountID string  `json:"accountId" api:"required"`
	MediaID   string  `json:"mediaId" api:"required"`
	ChannelID string  `json:"channelId" api:"nullable"`
	CreatedAt string  `json:"createdAt" api:"nullable"`
	Position  float64 `json:"position" api:"nullable"`
	UpdatedAt string  `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AccountID   respjson.Field
		MediaID     respjson.Field
		ChannelID   respjson.Field
		CreatedAt   respjson.Field
		Position    respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoChannelAssignMediaResponse) RawJSON() string { return r.JSON.raw }
func (r *VideoChannelAssignMediaResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoChannelRemoveMediaResponse struct {
	ID             string `json:"id" api:"required"`
	AccountID      string `json:"accountId" api:"required"`
	Name           string `json:"name" api:"required"`
	BrandColor     string `json:"brandColor" api:"nullable"`
	CreatedAt      string `json:"createdAt" api:"nullable"`
	Description    string `json:"description" api:"nullable"`
	FooterText     string `json:"footerText" api:"nullable"`
	IsPublished    bool   `json:"isPublished" api:"nullable"`
	LogoMediaID    string `json:"logoMediaId" api:"nullable"`
	ResolvedValues any    `json:"resolvedValues" api:"nullable"`
	UpdatedAt      string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		AccountID      respjson.Field
		Name           respjson.Field
		BrandColor     respjson.Field
		CreatedAt      respjson.Field
		Description    respjson.Field
		FooterText     respjson.Field
		IsPublished    respjson.Field
		LogoMediaID    respjson.Field
		ResolvedValues respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoChannelRemoveMediaResponse) RawJSON() string { return r.JSON.raw }
func (r *VideoChannelRemoveMediaResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoChannelReorderResponse struct {
	ID             string `json:"id" api:"required"`
	AccountID      string `json:"accountId" api:"required"`
	Name           string `json:"name" api:"required"`
	BrandColor     string `json:"brandColor" api:"nullable"`
	CreatedAt      string `json:"createdAt" api:"nullable"`
	Description    string `json:"description" api:"nullable"`
	FooterText     string `json:"footerText" api:"nullable"`
	IsPublished    bool   `json:"isPublished" api:"nullable"`
	LogoMediaID    string `json:"logoMediaId" api:"nullable"`
	ResolvedValues any    `json:"resolvedValues" api:"nullable"`
	UpdatedAt      string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		AccountID      respjson.Field
		Name           respjson.Field
		BrandColor     respjson.Field
		CreatedAt      respjson.Field
		Description    respjson.Field
		FooterText     respjson.Field
		IsPublished    respjson.Field
		LogoMediaID    respjson.Field
		ResolvedValues respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoChannelReorderResponse) RawJSON() string { return r.JSON.raw }
func (r *VideoChannelReorderResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoChannelListParams struct {
	// Maximum number of channels to return. Defaults to 25; values below 1 are clamped
	// to 1 and values above 1000 are clamped to 1000.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Opaque pagination cursor from pagination.nextCursor in the previous response. Do
	// not decode or modify it. Malformed cursors return 400 Bad Request.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VideoChannelListParams]'s query parameters as `url.Values`.
func (r VideoChannelListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VideoChannelNewParams struct {
	// Channel name. Case-insensitively unique within the account.
	Name string `json:"name,omitzero" api:"required"`
	paramObj
}

func (r VideoChannelNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VideoChannelNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VideoChannelNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoChannelUpdateParams struct {
	// Accent color used on the channel page. Any CSS color string.
	BrandColor  param.Opt[string] `json:"brandColor,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	FooterText  param.Opt[string] `json:"footerText,omitzero"`
	// Whether the channel page is publicly reachable. Publishing renders the page from
	// the current videos and branding; unpublishing takes it offline.
	IsPublished param.Opt[bool] `json:"isPublished,omitzero"`
	// Id of an image in your media library to show as the channel logo. Must be an
	// image; send `null` to clear it.
	LogoMediaID param.Opt[string] `json:"logoMediaId,omitzero"`
	// New channel name. Must stay case-insensitively unique within the account.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r VideoChannelUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow VideoChannelUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VideoChannelUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoChannelAssignMediaParams struct {
	// Id of the video to add to the channel. Must be a video, not an image.
	MediaID string `json:"mediaId,omitzero" api:"required" format:"uuid"`
	// Zero-based slot in the channel order. Omit to append to the end; omitting it on a
	// video that is already in the channel keeps its current slot.
	Position param.Opt[int64] `json:"position,omitzero"`
	paramObj
}

func (r VideoChannelAssignMediaParams) MarshalJSON() (data []byte, err error) {
	type shadow VideoChannelAssignMediaParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VideoChannelAssignMediaParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoChannelRemoveMediaParams struct {
	// Id of the video to remove from this channel.
	MediaID string `query:"mediaId" api:"required" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [VideoChannelRemoveMediaParams]'s query parameters as
// `url.Values`.
func (r VideoChannelRemoveMediaParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VideoChannelReorderParams struct {
	// Every video id currently in the channel, in the order you want them shown.
	// Partial lists and ids that are not in the channel are rejected.
	MediaIDs []string `json:"mediaIds,omitzero" api:"required"`
	paramObj
}

func (r VideoChannelReorderParams) MarshalJSON() (data []byte, err error) {
	type shadow VideoChannelReorderParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VideoChannelReorderParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
