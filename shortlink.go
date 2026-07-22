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

// ShortLinkService contains methods and other services that help with interacting
// with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewShortLinkService] method instead.
type ShortLinkService struct {
	Options []option.RequestOption
}

// NewShortLinkService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewShortLinkService(opts ...option.RequestOption) (r ShortLinkService) {
	r = ShortLinkService{}
	r.Options = opts
	return
}

// List all short links (QR codes / redirects) for this account, newest first.
// Supports cursor pagination and optional `status` and `nameContains` filters.
// Each entity bundles the destination URL, the composed public `shortUrl`, and the
// QR/campaign design. Requires scope: source:list
func (r *ShortLinkService) List(ctx context.Context, query ShortLinkListParams, opts ...option.RequestOption) (res *pagination.Cursor[ShortLinkListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "rest/v1/short-links"
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

// List all short links (QR codes / redirects) for this account, newest first.
// Supports cursor pagination and optional `status` and `nameContains` filters.
// Each entity bundles the destination URL, the composed public `shortUrl`, and the
// QR/campaign design. Requires scope: source:list
func (r *ShortLinkService) ListAutoPaging(ctx context.Context, query ShortLinkListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[ShortLinkListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Create a short link (QR code / redirect) with its destination, campaign tags,
// and QR styling in a single call. The short code is generated automatically; the
// response `shortUrl` is the public URL the QR encodes. All body fields are
// optional — send `{}` to create an unconfigured link and fill it in later with
// PATCH. A newly created short link only resolves at the edge once a version is
// published. Requires scope: source:create
func (r *ShortLinkService) New(ctx context.Context, body ShortLinkNewParams, opts ...option.RequestOption) (res *ShortLinkNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/short-links"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Fetch a single short link by id, including its destination, composed `shortUrl`,
// and QR/campaign design. Returns 404 when no short link matches the id or it
// belongs to a different account. Requires scope: source:view
func (r *ShortLinkService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ShortLinkGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/short-links/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Partially update a short link. Only the fields you send are changed; omitted
// fields are unchanged. Send explicit `null` to clear `redirectUrl`. The `utm` and
// `qr` objects are replaced wholesale when sent. Returns the full short link
// entity after the update. Requires scope: source:update
func (r *ShortLinkService) Update(ctx context.Context, id string, body ShortLinkUpdateParams, opts ...option.RequestOption) (res *ShortLinkUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/short-links/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete a short link and its QR/campaign design. After deletion the short URL
// stops resolving on the next publish. Requires scope: source:delete
func (r *ShortLinkService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *ShortLinkDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/short-links/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Aggregate click analytics for a short link over a date window: total and unique
// clicks, a time series (daily or hourly), and breakdowns by country, city, and
// device. QR scans are counted as clicks. Pass `from`/`to` as UTC calendar days
// (`YYYY-MM-DD`); set `granularity=HOURLY` for hourly buckets and
// `excludeBots=false` to include bot traffic. Requires the `shortlink:reporting`
// scope, which is gated separately because analytics data is PHI-bearing. Requires
// scope: shortlink:reporting
func (r *ShortLinkService) Results(ctx context.Context, id string, query ShortLinkResultsParams, opts ...option.RequestOption) (res *ShortLinkResultsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/short-links/%s/results", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ShortLinkListResponse struct {
	ID string `json:"id" api:"required"`
	// Organization id that owns this short link.
	AccountID string `json:"accountId" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	// Any of "Disabled", "Enabled".
	Status ShortLinkListResponseStatus `json:"status" api:"required"`
	// Whether this short link exists in the currently published version. An
	// unpublished short link does not resolve at the edge.
	IsPublished bool   `json:"isPublished" api:"nullable"`
	Name        string `json:"name" api:"nullable"`
	// The short code embedded in the public URL (`/redirect/{pixel}`).
	// Server-assigned.
	Pixel string `json:"pixel" api:"nullable"`
	// The destination URL this short link redirects to.
	RedirectURL string `json:"redirectUrl" api:"nullable"`
	// QR styling + campaign tags. Null until the link is styled.
	ShortLinkDesign any `json:"shortLinkDesign" api:"nullable"`
	// The public short-link URL that the QR encodes and callers share. Composed from
	// the short code, the link name (sent as the tracked event), and the campaign
	// tags. Also resolves on any branded custom domains configured for the account.
	ShortURL string `json:"shortUrl" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AccountID       respjson.Field
		CreatedAt       respjson.Field
		Status          respjson.Field
		IsPublished     respjson.Field
		Name            respjson.Field
		Pixel           respjson.Field
		RedirectURL     respjson.Field
		ShortLinkDesign respjson.Field
		ShortURL        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortLinkListResponse) RawJSON() string { return r.JSON.raw }
func (r *ShortLinkListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortLinkListResponseStatus string

const (
	ShortLinkListResponseStatusDisabled ShortLinkListResponseStatus = "Disabled"
	ShortLinkListResponseStatusEnabled  ShortLinkListResponseStatus = "Enabled"
)

type ShortLinkNewResponse struct {
	ID string `json:"id" api:"required"`
	// Organization id that owns this short link.
	AccountID string `json:"accountId" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	// Any of "Disabled", "Enabled".
	Status ShortLinkNewResponseStatus `json:"status" api:"required"`
	// Whether this short link exists in the currently published version. An
	// unpublished short link does not resolve at the edge.
	IsPublished bool   `json:"isPublished" api:"nullable"`
	Name        string `json:"name" api:"nullable"`
	// The short code embedded in the public URL (`/redirect/{pixel}`).
	// Server-assigned.
	Pixel string `json:"pixel" api:"nullable"`
	// The destination URL this short link redirects to.
	RedirectURL string `json:"redirectUrl" api:"nullable"`
	// QR styling + campaign tags. Null until the link is styled.
	ShortLinkDesign any `json:"shortLinkDesign" api:"nullable"`
	// The public short-link URL that the QR encodes and callers share. Composed from
	// the short code, the link name (sent as the tracked event), and the campaign
	// tags. Also resolves on any branded custom domains configured for the account.
	ShortURL string `json:"shortUrl" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AccountID       respjson.Field
		CreatedAt       respjson.Field
		Status          respjson.Field
		IsPublished     respjson.Field
		Name            respjson.Field
		Pixel           respjson.Field
		RedirectURL     respjson.Field
		ShortLinkDesign respjson.Field
		ShortURL        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortLinkNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ShortLinkNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortLinkNewResponseStatus string

const (
	ShortLinkNewResponseStatusDisabled ShortLinkNewResponseStatus = "Disabled"
	ShortLinkNewResponseStatusEnabled  ShortLinkNewResponseStatus = "Enabled"
)

type ShortLinkGetResponse struct {
	ID string `json:"id" api:"required"`
	// Organization id that owns this short link.
	AccountID string `json:"accountId" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	// Any of "Disabled", "Enabled".
	Status ShortLinkGetResponseStatus `json:"status" api:"required"`
	// Whether this short link exists in the currently published version. An
	// unpublished short link does not resolve at the edge.
	IsPublished bool   `json:"isPublished" api:"nullable"`
	Name        string `json:"name" api:"nullable"`
	// The short code embedded in the public URL (`/redirect/{pixel}`).
	// Server-assigned.
	Pixel string `json:"pixel" api:"nullable"`
	// The destination URL this short link redirects to.
	RedirectURL string `json:"redirectUrl" api:"nullable"`
	// QR styling + campaign tags. Null until the link is styled.
	ShortLinkDesign any `json:"shortLinkDesign" api:"nullable"`
	// The public short-link URL that the QR encodes and callers share. Composed from
	// the short code, the link name (sent as the tracked event), and the campaign
	// tags. Also resolves on any branded custom domains configured for the account.
	ShortURL string `json:"shortUrl" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AccountID       respjson.Field
		CreatedAt       respjson.Field
		Status          respjson.Field
		IsPublished     respjson.Field
		Name            respjson.Field
		Pixel           respjson.Field
		RedirectURL     respjson.Field
		ShortLinkDesign respjson.Field
		ShortURL        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortLinkGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ShortLinkGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortLinkGetResponseStatus string

const (
	ShortLinkGetResponseStatusDisabled ShortLinkGetResponseStatus = "Disabled"
	ShortLinkGetResponseStatusEnabled  ShortLinkGetResponseStatus = "Enabled"
)

type ShortLinkUpdateResponse struct {
	ID string `json:"id" api:"required"`
	// Organization id that owns this short link.
	AccountID string `json:"accountId" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	// Any of "Disabled", "Enabled".
	Status ShortLinkUpdateResponseStatus `json:"status" api:"required"`
	// Whether this short link exists in the currently published version. An
	// unpublished short link does not resolve at the edge.
	IsPublished bool   `json:"isPublished" api:"nullable"`
	Name        string `json:"name" api:"nullable"`
	// The short code embedded in the public URL (`/redirect/{pixel}`).
	// Server-assigned.
	Pixel string `json:"pixel" api:"nullable"`
	// The destination URL this short link redirects to.
	RedirectURL string `json:"redirectUrl" api:"nullable"`
	// QR styling + campaign tags. Null until the link is styled.
	ShortLinkDesign any `json:"shortLinkDesign" api:"nullable"`
	// The public short-link URL that the QR encodes and callers share. Composed from
	// the short code, the link name (sent as the tracked event), and the campaign
	// tags. Also resolves on any branded custom domains configured for the account.
	ShortURL string `json:"shortUrl" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AccountID       respjson.Field
		CreatedAt       respjson.Field
		Status          respjson.Field
		IsPublished     respjson.Field
		Name            respjson.Field
		Pixel           respjson.Field
		RedirectURL     respjson.Field
		ShortLinkDesign respjson.Field
		ShortURL        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortLinkUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ShortLinkUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortLinkUpdateResponseStatus string

const (
	ShortLinkUpdateResponseStatusDisabled ShortLinkUpdateResponseStatus = "Disabled"
	ShortLinkUpdateResponseStatusEnabled  ShortLinkUpdateResponseStatus = "Enabled"
)

type ShortLinkDeleteResponse struct {
	// Any of true.
	Deleted bool `json:"deleted" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deleted     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortLinkDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ShortLinkDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortLinkResultsResponse struct {
	Devices      []ShortLinkResultsResponseDevice       `json:"devices" api:"required"`
	GeoByCity    []ShortLinkResultsResponseGeoByCity    `json:"geoByCity" api:"required"`
	GeoByCountry []ShortLinkResultsResponseGeoByCountry `json:"geoByCountry" api:"required"`
	TimeSeries   []ShortLinkResultsResponseTimeSeries   `json:"timeSeries" api:"required"`
	TotalClicks  float64                                `json:"totalClicks" api:"required"`
	UniqueClicks float64                                `json:"uniqueClicks" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Devices      respjson.Field
		GeoByCity    respjson.Field
		GeoByCountry respjson.Field
		TimeSeries   respjson.Field
		TotalClicks  respjson.Field
		UniqueClicks respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortLinkResultsResponse) RawJSON() string { return r.JSON.raw }
func (r *ShortLinkResultsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortLinkResultsResponseDevice struct {
	Clicks       float64 `json:"clicks" api:"required"`
	Name         string  `json:"name" api:"required"`
	UniqueClicks float64 `json:"uniqueClicks" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Clicks       respjson.Field
		Name         respjson.Field
		UniqueClicks respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortLinkResultsResponseDevice) RawJSON() string { return r.JSON.raw }
func (r *ShortLinkResultsResponseDevice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortLinkResultsResponseGeoByCity struct {
	Clicks       float64 `json:"clicks" api:"required"`
	Name         string  `json:"name" api:"required"`
	UniqueClicks float64 `json:"uniqueClicks" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Clicks       respjson.Field
		Name         respjson.Field
		UniqueClicks respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortLinkResultsResponseGeoByCity) RawJSON() string { return r.JSON.raw }
func (r *ShortLinkResultsResponseGeoByCity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortLinkResultsResponseGeoByCountry struct {
	Clicks       float64 `json:"clicks" api:"required"`
	Name         string  `json:"name" api:"required"`
	UniqueClicks float64 `json:"uniqueClicks" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Clicks       respjson.Field
		Name         respjson.Field
		UniqueClicks respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortLinkResultsResponseGeoByCountry) RawJSON() string { return r.JSON.raw }
func (r *ShortLinkResultsResponseGeoByCountry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortLinkResultsResponseTimeSeries struct {
	Clicks       float64 `json:"clicks" api:"required"`
	Period       string  `json:"period" api:"required"`
	UniqueClicks float64 `json:"uniqueClicks" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Clicks       respjson.Field
		Period       respjson.Field
		UniqueClicks respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortLinkResultsResponseTimeSeries) RawJSON() string { return r.JSON.raw }
func (r *ShortLinkResultsResponseTimeSeries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortLinkListParams struct {
	// Maximum number of items to return. Defaults to 25; values below 1 are clamped to
	// 1 and values above 100 are clamped to 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Opaque pagination cursor from pagination.nextCursor in the previous response. Do
	// not decode or modify it. Malformed cursors return 400 Bad Request.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Case-insensitive substring filter on the short link name.
	NameContains param.Opt[string] `query:"nameContains,omitzero" json:"-"`
	// Filter by short link status.
	//
	// Any of "Disabled", "Enabled".
	Status ShortLinkListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ShortLinkListParams]'s query parameters as `url.Values`.
func (r ShortLinkListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by short link status.
type ShortLinkListParamsStatus string

const (
	ShortLinkListParamsStatusDisabled ShortLinkListParamsStatus = "Disabled"
	ShortLinkListParamsStatusEnabled  ShortLinkListParamsStatus = "Enabled"
)

type ShortLinkNewParams struct {
	// Human-readable name. Also sent as the tracked event name on every click/scan.
	Name param.Opt[string] `json:"name,omitzero"`
	// Destination URL the short link redirects to. Must be a valid URL.
	RedirectURL param.Opt[string] `json:"redirectUrl,omitzero"`
	// QR code visual styling.
	Qr any `json:"qr,omitzero"`
	// Campaign / UTM tags appended to the tracked short-link URL.
	Utm any `json:"utm,omitzero"`
	paramObj
}

func (r ShortLinkNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ShortLinkNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ShortLinkNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortLinkUpdateParams struct {
	Name param.Opt[string] `json:"name,omitzero"`
	// Destination URL the short link redirects to. Must be a valid URL. Send `null` to
	// clear it.
	RedirectURL param.Opt[string] `json:"redirectUrl,omitzero"`
	// Whether the short link resolves at the edge. Send `Enabled` or `Disabled`;
	// `null` is rejected since storage cannot represent it.
	Status param.Opt[string] `json:"status,omitzero"`
	Qr     any               `json:"qr,omitzero"`
	Utm    any               `json:"utm,omitzero"`
	paramObj
}

func (r ShortLinkUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ShortLinkUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ShortLinkUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortLinkResultsParams struct {
	// Inclusive lower bound of the report window, a UTC calendar day in `YYYY-MM-DD`.
	From string `query:"from" api:"required" json:"-"`
	// Inclusive upper bound of the report window, a UTC calendar day in `YYYY-MM-DD`.
	To string `query:"to" api:"required" json:"-"`
	// Exclude bot traffic from the counts. Defaults to `true`.
	ExcludeBots param.Opt[bool] `query:"excludeBots,omitzero" json:"-"`
	// Time-series bucket size. Defaults to `DAILY`.
	//
	// Any of "DAILY", "HOURLY".
	Granularity ShortLinkResultsParamsGranularity `query:"granularity,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ShortLinkResultsParams]'s query parameters as `url.Values`.
func (r ShortLinkResultsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Time-series bucket size. Defaults to `DAILY`.
type ShortLinkResultsParamsGranularity string

const (
	ShortLinkResultsParamsGranularityDaily  ShortLinkResultsParamsGranularity = "DAILY"
	ShortLinkResultsParamsGranularityHourly ShortLinkResultsParamsGranularity = "HOURLY"
)
