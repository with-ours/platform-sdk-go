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

// ReplaySettingService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReplaySettingService] method instead.
type ReplaySettingService struct {
	Options []option.RequestOption
}

// NewReplaySettingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewReplaySettingService(opts ...option.RequestOption) (r ReplaySettingService) {
	r = ReplaySettingService{}
	r.Options = opts
	return
}

// List the replay configurations on this account. Supports cursor pagination via
// `limit` and `cursor`. Replay settings control which domains may capture session
// replays and where the capture script is hosted. Requires scope:
// replaySettings:list
func (r *ReplaySettingService) List(ctx context.Context, query ReplaySettingListParams, opts ...option.RequestOption) (res *pagination.Cursor[ReplaySettingListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "rest/v1/replay-settings"
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

// List the replay configurations on this account. Supports cursor pagination via
// `limit` and `cursor`. Replay settings control which domains may capture session
// replays and where the capture script is hosted. Requires scope:
// replaySettings:list
func (r *ReplaySettingService) ListAutoPaging(ctx context.Context, query ReplaySettingListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[ReplaySettingListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Create the replay configuration for this account. Each account is limited to one
// replay configuration — calls made when one already exists return HTTP 409 with
// the reason in the response `error` field. Requires scope: replaySettings:create
func (r *ReplaySettingService) New(ctx context.Context, body ReplaySettingNewParams, opts ...option.RequestOption) (res *ReplaySettingNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/replay-settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Fetch a single replay configuration by ID, including its whitelisted domains and
// custom domain. Requires scope: replaySettings:find
func (r *ReplaySettingService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ReplaySettingGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/replay-settings/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update one or more fields on an existing replay configuration. Only the fields
// you send are changed; omitted fields keep their current value. Note that
// `whitelistDomains` is replaced wholesale (not merged with the existing list).
// Requires scope: replaySettings:update
func (r *ReplaySettingService) Update(ctx context.Context, id string, body ReplaySettingUpdateParams, opts ...option.RequestOption) (res *ReplaySettingUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/replay-settings/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete the replay configuration. Capture stops immediately for all whitelisted
// domains. Requires scope: replaySettings:delete
func (r *ReplaySettingService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *ReplaySettingDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/replay-settings/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type ReplaySettingListResponse struct {
	// Stable identifier (UUID) for this replay configuration.
	ID string `json:"id" api:"required"`
	// ISO-8601 timestamp when this configuration was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Human-readable label for this replay configuration. Shown in the dashboard. May
	// be empty.
	Name string `json:"name" api:"required"`
	// Whether session replay capture is currently active. Set to "Enabled" to start
	// capturing replays from whitelisted domains, or "Disabled" to pause capture
	// without losing the configuration.
	//
	// Any of "Disabled", "Enabled".
	Status ReplaySettingListResponseStatus `json:"status" api:"required"`
	// Optional custom domain (CNAME) for hosting the replay capture script. Leave null
	// to use the default Ours Privacy domain.
	CustomDomain string `json:"customDomain" api:"nullable"`
	// ISO-8601 timestamp of the most recent update, or null if never updated.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// Hostnames where session replay capture is permitted. Replays initiated from any
	// host not in this list are dropped. PATCH replaces the list — partial updates are
	// not merged.
	WhitelistDomains []string `json:"whitelistDomains" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		Name             respjson.Field
		Status           respjson.Field
		CustomDomain     respjson.Field
		UpdatedAt        respjson.Field
		WhitelistDomains respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReplaySettingListResponse) RawJSON() string { return r.JSON.raw }
func (r *ReplaySettingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether session replay capture is currently active. Set to "Enabled" to start
// capturing replays from whitelisted domains, or "Disabled" to pause capture
// without losing the configuration.
type ReplaySettingListResponseStatus string

const (
	ReplaySettingListResponseStatusDisabled ReplaySettingListResponseStatus = "Disabled"
	ReplaySettingListResponseStatusEnabled  ReplaySettingListResponseStatus = "Enabled"
)

type ReplaySettingNewResponse struct {
	IsSuccess      bool   `json:"isSuccess" api:"required"`
	Cause          string `json:"cause" api:"nullable"`
	ReplaySettings any    `json:"replaySettings" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsSuccess      respjson.Field
		Cause          respjson.Field
		ReplaySettings respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReplaySettingNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ReplaySettingNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReplaySettingGetResponse struct {
	// Stable identifier (UUID) for this replay configuration.
	ID string `json:"id" api:"required"`
	// ISO-8601 timestamp when this configuration was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Human-readable label for this replay configuration. Shown in the dashboard. May
	// be empty.
	Name string `json:"name" api:"required"`
	// Whether session replay capture is currently active. Set to "Enabled" to start
	// capturing replays from whitelisted domains, or "Disabled" to pause capture
	// without losing the configuration.
	//
	// Any of "Disabled", "Enabled".
	Status ReplaySettingGetResponseStatus `json:"status" api:"required"`
	// Optional custom domain (CNAME) for hosting the replay capture script. Leave null
	// to use the default Ours Privacy domain.
	CustomDomain string `json:"customDomain" api:"nullable"`
	// ISO-8601 timestamp of the most recent update, or null if never updated.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// Hostnames where session replay capture is permitted. Replays initiated from any
	// host not in this list are dropped. PATCH replaces the list — partial updates are
	// not merged.
	WhitelistDomains []string `json:"whitelistDomains" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		Name             respjson.Field
		Status           respjson.Field
		CustomDomain     respjson.Field
		UpdatedAt        respjson.Field
		WhitelistDomains respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReplaySettingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ReplaySettingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether session replay capture is currently active. Set to "Enabled" to start
// capturing replays from whitelisted domains, or "Disabled" to pause capture
// without losing the configuration.
type ReplaySettingGetResponseStatus string

const (
	ReplaySettingGetResponseStatusDisabled ReplaySettingGetResponseStatus = "Disabled"
	ReplaySettingGetResponseStatusEnabled  ReplaySettingGetResponseStatus = "Enabled"
)

type ReplaySettingUpdateResponse struct {
	IsSuccess      bool   `json:"isSuccess" api:"required"`
	Cause          string `json:"cause" api:"nullable"`
	ReplaySettings any    `json:"replaySettings" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsSuccess      respjson.Field
		Cause          respjson.Field
		ReplaySettings respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReplaySettingUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ReplaySettingUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReplaySettingDeleteResponse struct {
	IsSuccess      bool   `json:"isSuccess" api:"required"`
	Cause          string `json:"cause" api:"nullable"`
	ReplaySettings any    `json:"replaySettings" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsSuccess      respjson.Field
		Cause          respjson.Field
		ReplaySettings respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReplaySettingDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ReplaySettingDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReplaySettingListParams struct {
	// Maximum number of items to return. Defaults to 25; values below 1 are clamped to
	// 1 and values above 100 are clamped to 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Opaque pagination cursor from pagination.nextCursor in the previous response. Do
	// not decode or modify it. Malformed cursors return 400 Bad Request.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ReplaySettingListParams]'s query parameters as
// `url.Values`.
func (r ReplaySettingListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ReplaySettingNewParams struct {
	// Optional custom domain (CNAME) for hosting the replay capture script. Leave null
	// to use the default Ours Privacy domain.
	CustomDomain param.Opt[string] `json:"customDomain,omitzero"`
	// Human-readable label for this replay configuration. Shown in the dashboard. May
	// be empty.
	Name param.Opt[string] `json:"name,omitzero"`
	// Whether session replay capture is currently active. Set to "Enabled" to start
	// capturing replays from whitelisted domains, or "Disabled" to pause capture
	// without losing the configuration.
	//
	// Any of "Disabled", "Enabled".
	Status ReplaySettingNewParamsStatus `json:"status,omitzero"`
	// Hostnames where session replay capture is permitted. Replays initiated from any
	// host not in this list are dropped. PATCH replaces the list — partial updates are
	// not merged.
	WhitelistDomains []string `json:"whitelistDomains,omitzero"`
	paramObj
}

func (r ReplaySettingNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ReplaySettingNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReplaySettingNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether session replay capture is currently active. Set to "Enabled" to start
// capturing replays from whitelisted domains, or "Disabled" to pause capture
// without losing the configuration.
type ReplaySettingNewParamsStatus string

const (
	ReplaySettingNewParamsStatusDisabled ReplaySettingNewParamsStatus = "Disabled"
	ReplaySettingNewParamsStatusEnabled  ReplaySettingNewParamsStatus = "Enabled"
)

type ReplaySettingUpdateParams struct {
	// Optional custom domain (CNAME) for hosting the replay capture script. Leave null
	// to use the default Ours Privacy domain.
	CustomDomain param.Opt[string] `json:"customDomain,omitzero"`
	// Human-readable label for this replay configuration. Shown in the dashboard. May
	// be empty.
	Name param.Opt[string] `json:"name,omitzero"`
	// Whether session replay capture is currently active. Set to "Enabled" to start
	// capturing replays from whitelisted domains, or "Disabled" to pause capture
	// without losing the configuration.
	//
	// Any of "Disabled", "Enabled".
	Status ReplaySettingUpdateParamsStatus `json:"status,omitzero"`
	// Hostnames where session replay capture is permitted. Replays initiated from any
	// host not in this list are dropped. PATCH replaces the list — partial updates are
	// not merged.
	WhitelistDomains []string `json:"whitelistDomains,omitzero"`
	paramObj
}

func (r ReplaySettingUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ReplaySettingUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReplaySettingUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether session replay capture is currently active. Set to "Enabled" to start
// capturing replays from whitelisted domains, or "Disabled" to pause capture
// without losing the configuration.
type ReplaySettingUpdateParamsStatus string

const (
	ReplaySettingUpdateParamsStatusDisabled ReplaySettingUpdateParamsStatus = "Disabled"
	ReplaySettingUpdateParamsStatusEnabled  ReplaySettingUpdateParamsStatus = "Enabled"
)
