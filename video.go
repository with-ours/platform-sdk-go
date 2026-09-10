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

// VideoService contains methods and other services that help with interacting with
// the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVideoService] method instead.
type VideoService struct {
	Options []option.RequestOption
}

// NewVideoService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there is
// one), and before any request-specific options.
func NewVideoService(opts ...option.RequestOption) (r VideoService) {
	r = VideoService{}
	r.Options = opts
	return
}

// List videos for the account, newest first. Supports cursor pagination and an
// optional case-insensitive title filter. Requires scope: media:list
func (r *VideoService) List(ctx context.Context, query VideoListParams, opts ...option.RequestOption) (res *pagination.Cursor[VideoListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "rest/v1/videos"
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

// List videos for the account, newest first. Supports cursor pagination and an
// optional case-insensitive title filter. Requires scope: media:list
func (r *VideoService) ListAutoPaging(ctx context.Context, query VideoListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[VideoListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Create a video record and return a temporary upload target for the original MP4
// or WebM file. Upload the file directly using the returned URL and matching
// content type, then poll the video to observe processing progress. Requires scope:
// media:create
func (r *VideoService) New(ctx context.Context, body VideoNewParams, opts ...option.RequestOption) (res *VideoNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/videos"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Fetch a video and its current playback asset availability. The processed video,
// poster, and transcript are prepared asynchronously after upload. Requires scope:
// media:find
func (r *VideoService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *VideoGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/videos/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Partially update video metadata. Only fields included in the body change; send
// `null` to clear a nullable field. Requires scope: media:update
func (r *VideoService) Update(ctx context.Context, id string, body VideoUpdateParams, opts ...option.RequestOption) (res *VideoUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/videos/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete a video and its related assets. Requires scope: media:delete
func (r *VideoService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *VideoDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/videos/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Read the current WebVTT transcript. Transcript text is available wherever the
// video is embedded, so do not include PHI or other confidential information.
// Requires scope: media:find
func (r *VideoService) Transcript(ctx context.Context, id string, opts ...option.RequestOption) (res *VideoTranscriptResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/videos/%s/transcript", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Replace the transcript with VTT or SRT text. SRT is normalized to WebVTT.
// Transcript text is available wherever the video is embedded, so do not include
// PHI or other confidential information. Requires scope: media:update
func (r *VideoService) UpdateTranscript(ctx context.Context, id string, body VideoUpdateTranscriptParams, opts ...option.RequestOption) (res *VideoUpdateTranscriptResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/videos/%s/transcript", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Return per-video starts, unique viewers, completion rate, and average watch time
// for a date window. This derived report uses `limit` and `offset` pagination;
// `total` is the number of rows returned through the current offset, not a total
// match count. Requires scope: report:video-analytics
func (r *VideoService) Analytics(ctx context.Context, query VideoAnalyticsParams, opts ...option.RequestOption) (res *VideoAnalyticsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/videos/analytics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Return daily or hourly starts, unique viewers, completions, and completion rate
// for one video. Daily windows support up to 90 days; hourly windows support up to
// 14 days. Requires scope: report:video-analytics
func (r *VideoService) AnalyticsTimeseries(ctx context.Context, id string, query VideoAnalyticsTimeseriesParams, opts ...option.RequestOption) (res *VideoAnalyticsTimeseriesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/videos/%s/analytics", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type VideoListResponse struct {
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
func (r VideoListResponse) RawJSON() string { return r.JSON.raw }
func (r *VideoListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoNewResponse struct {
	ID        string `json:"id" api:"required"`
	AccountID string `json:"accountId" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	// Any of "Video".
	Type                  string                 `json:"type" api:"required"`
	Upload                VideoNewResponseUpload `json:"upload" api:"required"`
	CaptionsUpdatedAt     string                 `json:"captionsUpdatedAt" api:"nullable"`
	CaptionsUpdatedByName string                 `json:"captionsUpdatedByName" api:"nullable"`
	Description           string                 `json:"description" api:"nullable"`
	Duration              float64                `json:"duration" api:"nullable"`
	HasVideoUpload        bool                   `json:"hasVideoUpload" api:"nullable"`
	Height                float64                `json:"height" api:"nullable"`
	Name                  string                 `json:"name" api:"nullable"`
	UpdatedAt             string                 `json:"updatedAt" api:"nullable"`
	Width                 float64                `json:"width" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		AccountID             respjson.Field
		CreatedAt             respjson.Field
		Type                  respjson.Field
		Upload                respjson.Field
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
func (r VideoNewResponse) RawJSON() string { return r.JSON.raw }
func (r *VideoNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoNewResponseUpload struct {
	// Any of "MP4", "WEBM".
	MimeType VideoNewResponseUploadMimeType `json:"mimeType" api:"required"`
	URL      string                         `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MimeType    respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoNewResponseUpload) RawJSON() string { return r.JSON.raw }
func (r *VideoNewResponseUpload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoNewResponseUploadMimeType string

const (
	VideoNewResponseUploadMimeTypeMP4  VideoNewResponseUploadMimeType = "MP4"
	VideoNewResponseUploadMimeTypeWebm VideoNewResponseUploadMimeType = "WEBM"
)

type VideoGetResponse struct {
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
	ResolvedValues        any     `json:"resolvedValues" api:"nullable"`
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
		ResolvedValues        respjson.Field
		UpdatedAt             respjson.Field
		Width                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoGetResponse) RawJSON() string { return r.JSON.raw }
func (r *VideoGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoUpdateResponse struct {
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
func (r VideoUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *VideoUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoDeleteResponse struct {
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
func (r VideoDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *VideoDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTranscriptResponse struct {
	Content string `json:"content" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoTranscriptResponse) RawJSON() string { return r.JSON.raw }
func (r *VideoTranscriptResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoUpdateTranscriptResponse struct {
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
func (r VideoUpdateTranscriptResponse) RawJSON() string { return r.JSON.raw }
func (r *VideoUpdateTranscriptResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoAnalyticsResponse struct {
	HasMore bool                         `json:"hasMore" api:"required"`
	Items   []VideoAnalyticsResponseItem `json:"items" api:"required"`
	Total   int64                        `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		Items       respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoAnalyticsResponse) RawJSON() string { return r.JSON.raw }
func (r *VideoAnalyticsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoAnalyticsResponseItem struct {
	AvgVideoDurationSeconds float64 `json:"avgVideoDurationSeconds" api:"required"`
	AvgWatchTimeSeconds     float64 `json:"avgWatchTimeSeconds" api:"required"`
	CompletionRate          float64 `json:"completionRate" api:"required"`
	UniqueViewers           int64   `json:"uniqueViewers" api:"required"`
	VideoStarts             int64   `json:"videoStarts" api:"required"`
	VideoTitle              string  `json:"videoTitle" api:"required"`
	VideoID                 string  `json:"videoId" api:"nullable"`
	VideoURL                string  `json:"videoUrl" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvgVideoDurationSeconds respjson.Field
		AvgWatchTimeSeconds     respjson.Field
		CompletionRate          respjson.Field
		UniqueViewers           respjson.Field
		VideoStarts             respjson.Field
		VideoTitle              respjson.Field
		VideoID                 respjson.Field
		VideoURL                respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoAnalyticsResponseItem) RawJSON() string { return r.JSON.raw }
func (r *VideoAnalyticsResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoAnalyticsTimeseriesResponse struct {
	Items []VideoAnalyticsTimeseriesResponseItem `json:"items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoAnalyticsTimeseriesResponse) RawJSON() string { return r.JSON.raw }
func (r *VideoAnalyticsTimeseriesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoAnalyticsTimeseriesResponseItem struct {
	CompletionRate float64 `json:"completionRate" api:"required"`
	Completions    int64   `json:"completions" api:"required"`
	DateTime       string  `json:"dateTime" api:"required"`
	UniqueViewers  int64   `json:"uniqueViewers" api:"required"`
	VideoStarts    int64   `json:"videoStarts" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletionRate respjson.Field
		Completions    respjson.Field
		DateTime       respjson.Field
		UniqueViewers  respjson.Field
		VideoStarts    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoAnalyticsTimeseriesResponseItem) RawJSON() string { return r.JSON.raw }
func (r *VideoAnalyticsTimeseriesResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoListParams struct {
	// Maximum number of items to return. Defaults to 25; values below 1 are clamped to
	// 1 and values above 100 are clamped to 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Opaque pagination cursor from pagination.nextCursor in the previous response. Do
	// not decode or modify it. Malformed cursors return 400 Bad Request.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Case-insensitive substring match on the video title.
	NameContains param.Opt[string] `query:"nameContains,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VideoListParams]'s query parameters as `url.Values`.
func (r VideoListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VideoNewParams struct {
	// Content type for the original video upload: `MP4` or `WEBM`.
	//
	// Any of "MP4", "WEBM".
	MimeType    VideoNewParamsMimeType `json:"mimeType,omitzero" api:"required"`
	Description param.Opt[string]      `json:"description,omitzero"`
	// Video title. Defaults to `New Video` when omitted.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r VideoNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VideoNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VideoNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Content type for the original video upload: `MP4` or `WEBM`.
type VideoNewParamsMimeType string

const (
	VideoNewParamsMimeTypeMP4  VideoNewParamsMimeType = "MP4"
	VideoNewParamsMimeTypeWebm VideoNewParamsMimeType = "WEBM"
)

type VideoUpdateParams struct {
	Description    param.Opt[string]  `json:"description,omitzero"`
	Duration       param.Opt[float64] `json:"duration,omitzero"`
	HasVideoUpload param.Opt[bool]    `json:"hasVideoUpload,omitzero"`
	Height         param.Opt[float64] `json:"height,omitzero"`
	Name           param.Opt[string]  `json:"name,omitzero"`
	Width          param.Opt[float64] `json:"width,omitzero"`
	paramObj
}

func (r VideoUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow VideoUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VideoUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoUpdateTranscriptParams struct {
	// Transcript text, limited to 1 MB.
	Content string `json:"content" api:"required"`
	// Transcript source format. SRT content is normalized to WebVTT before it is
	// saved.
	//
	// Any of "SRT", "VTT".
	Format VideoUpdateTranscriptParamsFormat `json:"format,omitzero" api:"required"`
	paramObj
}

func (r VideoUpdateTranscriptParams) MarshalJSON() (data []byte, err error) {
	type shadow VideoUpdateTranscriptParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VideoUpdateTranscriptParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transcript source format. SRT content is normalized to WebVTT before it is saved.
type VideoUpdateTranscriptParamsFormat string

const (
	VideoUpdateTranscriptParamsFormatSrt VideoUpdateTranscriptParamsFormat = "SRT"
	VideoUpdateTranscriptParamsFormatVtt VideoUpdateTranscriptParamsFormat = "VTT"
)

type VideoAnalyticsParams struct {
	// Inclusive UTC start day in `YYYY-MM-DD` format.
	From string `query:"from" api:"required" json:"-"`
	// Inclusive UTC end day in `YYYY-MM-DD` format.
	To string `query:"to" api:"required" json:"-"`
	// Maximum number of video rows to return. Defaults to 50.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Zero-based row offset. This report is an intentional offset-pagination exception.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VideoAnalyticsParams]'s query parameters as `url.Values`.
func (r VideoAnalyticsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VideoAnalyticsTimeseriesParams struct {
	// Inclusive UTC start day in `YYYY-MM-DD` format.
	From string `query:"from" api:"required" json:"-"`
	// Inclusive UTC end day in `YYYY-MM-DD` format.
	To string `query:"to" api:"required" json:"-"`
	// Bucket size. Defaults to `DAILY`; `HOURLY` supports windows of up to 14 days.
	//
	// Any of "DAILY", "HOURLY".
	Granularity VideoAnalyticsTimeseriesParamsGranularity `query:"granularity,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VideoAnalyticsTimeseriesParams]'s query parameters as
// `url.Values`.
func (r VideoAnalyticsTimeseriesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Bucket size. Defaults to `DAILY`; `HOURLY` supports windows of up to 14 days.
type VideoAnalyticsTimeseriesParamsGranularity string

const (
	VideoAnalyticsTimeseriesParamsGranularityDaily  VideoAnalyticsTimeseriesParamsGranularity = "DAILY"
	VideoAnalyticsTimeseriesParamsGranularityHourly VideoAnalyticsTimeseriesParamsGranularity = "HOURLY"
)
