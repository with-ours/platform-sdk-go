// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy

import (
	"context"
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

// List links managed through Short Links for this account, newest first. Ordinary
// Redirect Tracking sources and legacy links without a saved short-link design are
// excluded. Supports cursor pagination and optional `status`, `nameContains`, and
// `search` filters. Search matches names, destinations, and codes. Each entity
// bundles the destination URL, immutable code, path format, and QR/campaign
// design. Requires API-key scope or current OAuth user permission: source:list
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

// List links managed through Short Links for this account, newest first. Ordinary
// Redirect Tracking sources and legacy links without a saved short-link design are
// excluded. Supports cursor pagination and optional `status`, `nameContains`, and
// `search` filters. Search matches names, destinations, and codes. Each entity
// bundles the destination URL, immutable code, path format, and QR/campaign
// design. Requires API-key scope or current OAuth user permission: source:list
func (r *ShortLinkService) ListAutoPaging(ctx context.Context, query ShortLinkListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[ShortLinkListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Create a short link (QR code / redirect) with its destination, campaign tags,
// and QR styling in a single call. The server atomically reserves an immutable
// compact code; `shortUrl` resolves as `/r/{pixel}` without tracking query
// parameters. All body fields are optional: send `{}` to create an unconfigured
// link and fill it in later with PATCH. A newly created short link only resolves
// at the edge once a version is published. Requires API-key scope or current OAuth
// user permission: source:create
func (r *ShortLinkService) New(ctx context.Context, body ShortLinkNewParams, opts ...option.RequestOption) (res *ShortLinkNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/short-links"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Fetch a single short link by id, including its destination, immutable code,
// composed `shortUrl`, and QR/campaign design. Returns 404 when no short link
// matches the id or it belongs to a different account. Requires API-key scope or
// current OAuth user permission: source:view
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
// entity after the update. Requires API-key scope or current OAuth user
// permission: source:update
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
// stops resolving on the next publish. Requires API-key scope or current OAuth
// user permission: source:delete
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

// Duplicate a short link with its campaign tags and QR styling. The copy starts
// disabled with a new code. Requires permission to view the original and create a
// source. Requires API-key scope or current OAuth user permission: source:create
func (r *ShortLinkService) Clone(ctx context.Context, id string, opts ...option.RequestOption) (res *ShortLinkCloneResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/short-links/%s/clone", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Aggregate click analytics for a short link over a date window: total and unique
// clicks, a time series (daily or hourly), and breakdowns by country, city, and
// device. QR scans are counted as clicks. Pass `from`/`to` as UTC calendar days
// (`YYYY-MM-DD`); set `granularity=HOURLY` for hourly buckets and
// `excludeBots=false` to include bot traffic. Requires the `shortlink:reporting`
// scope, which is gated separately because analytics data is PHI-bearing. Requires
// API-key scope or current OAuth user permission: shortlink:reporting
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
	// Whether saved published design campaign tags are added to the destination when
	// enabled. Existing destination UTMs stay unchanged. Defaults to true for new
	// short links and false for existing links without a stored setting.
	ForwardUtmToDestination bool `json:"forwardUtmToDestination" api:"required"`
	// Any of "Disabled", "Enabled".
	Status ShortLinkListResponseStatus `json:"status" api:"required"`
	// Whether a user selected this code instead of using a generated code.
	HasCustomShortLinkCode bool `json:"hasCustomShortLinkCode" api:"nullable"`
	// Whether this short link exists in the currently published version. An
	// unpublished short link does not resolve at the edge.
	IsPublished bool   `json:"isPublished" api:"nullable"`
	Name        string `json:"name" api:"nullable"`
	// The immutable short code embedded in the public URL (`/r/{pixel}` for new
	// links). Server-assigned.
	Pixel string `json:"pixel" api:"nullable"`
	// The destination URL this short link redirects to.
	RedirectURL string `json:"redirectUrl" api:"nullable"`
	// The public code embedded in the short-link URL.
	ShortLinkCode string `json:"shortLinkCode" api:"nullable"`
	// QR styling + campaign tags. New links include an empty design; legacy redirects
	// may return null.
	ShortLinkDesign ShortLinkListResponseShortLinkDesign `json:"shortLinkDesign" api:"nullable"`
	// The public short-link URL that the QR encodes and callers share. New links use
	// `/r/{pixel}`; tracked event and campaign defaults are stored server-side. Also
	// resolves on branded custom domains configured for the account.
	ShortURL string `json:"shortUrl" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		AccountID               respjson.Field
		CreatedAt               respjson.Field
		ForwardUtmToDestination respjson.Field
		Status                  respjson.Field
		HasCustomShortLinkCode  respjson.Field
		IsPublished             respjson.Field
		Name                    respjson.Field
		Pixel                   respjson.Field
		RedirectURL             respjson.Field
		ShortLinkCode           respjson.Field
		ShortLinkDesign         respjson.Field
		ShortURL                respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
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

// QR styling + campaign tags. New links include an empty design; legacy redirects
// may return null.
type ShortLinkListResponseShortLinkDesign struct {
	Qr  ShortLinkListResponseShortLinkDesignQr  `json:"qr" api:"nullable"`
	Utm ShortLinkListResponseShortLinkDesignUtm `json:"utm" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Qr          respjson.Field
		Utm         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortLinkListResponseShortLinkDesign) RawJSON() string { return r.JSON.raw }
func (r *ShortLinkListResponseShortLinkDesign) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortLinkListResponseShortLinkDesignQr struct {
	// Background color.
	BgColor string `json:"bgColor" api:"nullable"`
	// Finder (corner) dot shape: `square` or `dot`.
	CornerDotStyle string `json:"cornerDotStyle" api:"nullable"`
	// Finder (corner) square shape: `square`, `dot`, or `extra-rounded`.
	CornerSquareStyle string `json:"cornerSquareStyle" api:"nullable"`
	// Module shape: `square`, `rounded`, `dots`, `classy`, `classy-rounded`, or
	// `extra-rounded`.
	DotStyle string `json:"dotStyle" api:"nullable"`
	// Foreground (module) color.
	FgColor string `json:"fgColor" api:"nullable"`
	// Download frame: `square`, `rounded`, or `circle`.
	FrameShape string `json:"frameShape" api:"nullable"`
	// A two-stop gradient fill for the QR modules.
	Gradient ShortLinkListResponseShortLinkDesignQrGradient `json:"gradient" api:"nullable"`
	// Base64 data URL of a centered logo overlay.
	Logo string `json:"logo" api:"nullable"`
	// Logo size as a fraction of the QR (0–0.4).
	LogoSizeRatio float64 `json:"logoSizeRatio" api:"nullable"`
	// The selected preset theme id, for re-highlighting it in a UI.
	Theme string `json:"theme" api:"nullable"`
	// When true, the background is exported transparent.
	TransparentBackground bool `json:"transparentBackground" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BgColor               respjson.Field
		CornerDotStyle        respjson.Field
		CornerSquareStyle     respjson.Field
		DotStyle              respjson.Field
		FgColor               respjson.Field
		FrameShape            respjson.Field
		Gradient              respjson.Field
		Logo                  respjson.Field
		LogoSizeRatio         respjson.Field
		Theme                 respjson.Field
		TransparentBackground respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortLinkListResponseShortLinkDesignQr) RawJSON() string { return r.JSON.raw }
func (r *ShortLinkListResponseShortLinkDesignQr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A two-stop gradient fill for the QR modules.
type ShortLinkListResponseShortLinkDesignQrGradient struct {
	// Gradient end color.
	EndColor string `json:"endColor" api:"required"`
	// Gradient start color.
	StartColor string `json:"startColor" api:"required"`
	// Gradient type: `linear` or `radial`.
	Type string `json:"type" api:"required"`
	// Rotation in radians (linear gradients only).
	Rotation float64 `json:"rotation" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndColor    respjson.Field
		StartColor  respjson.Field
		Type        respjson.Field
		Rotation    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortLinkListResponseShortLinkDesignQrGradient) RawJSON() string { return r.JSON.raw }
func (r *ShortLinkListResponseShortLinkDesignQrGradient) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortLinkListResponseShortLinkDesignUtm struct {
	// The `utm_campaign` value appended to the tracked short-link URL.
	Campaign string `json:"campaign" api:"nullable"`
	// The `utm_content` value.
	Content string `json:"content" api:"nullable"`
	// The `utm_medium` value.
	Medium string `json:"medium" api:"nullable"`
	// The `utm_source` value.
	Source string `json:"source" api:"nullable"`
	// The `utm_term` value.
	Term string `json:"term" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Campaign    respjson.Field
		Content     respjson.Field
		Medium      respjson.Field
		Source      respjson.Field
		Term        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortLinkListResponseShortLinkDesignUtm) RawJSON() string { return r.JSON.raw }
func (r *ShortLinkListResponseShortLinkDesignUtm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortLinkNewResponse struct {
	ID string `json:"id" api:"required"`
	// Organization id that owns this short link.
	AccountID string `json:"accountId" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	// Whether saved published design campaign tags are added to the destination when
	// enabled. Existing destination UTMs stay unchanged. Defaults to true for new
	// short links and false for existing links without a stored setting.
	ForwardUtmToDestination bool `json:"forwardUtmToDestination" api:"required"`
	// Any of "Disabled", "Enabled".
	Status ShortLinkNewResponseStatus `json:"status" api:"required"`
	// Whether a user selected this code instead of using a generated code.
	HasCustomShortLinkCode bool `json:"hasCustomShortLinkCode" api:"nullable"`
	// Whether this short link exists in the currently published version. An
	// unpublished short link does not resolve at the edge.
	IsPublished bool   `json:"isPublished" api:"nullable"`
	Name        string `json:"name" api:"nullable"`
	// The immutable short code embedded in the public URL (`/r/{pixel}` for new
	// links). Server-assigned.
	Pixel string `json:"pixel" api:"nullable"`
	// The destination URL this short link redirects to.
	RedirectURL string `json:"redirectUrl" api:"nullable"`
	// The public code embedded in the short-link URL.
	ShortLinkCode string `json:"shortLinkCode" api:"nullable"`
	// QR styling + campaign tags. New links include an empty design; legacy redirects
	// may return null.
	ShortLinkDesign ShortLinkNewResponseShortLinkDesign `json:"shortLinkDesign" api:"nullable"`
	// The public short-link URL that the QR encodes and callers share. New links use
	// `/r/{pixel}`; tracked event and campaign defaults are stored server-side. Also
	// resolves on branded custom domains configured for the account.
	ShortURL string `json:"shortUrl" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		AccountID               respjson.Field
		CreatedAt               respjson.Field
		ForwardUtmToDestination respjson.Field
		Status                  respjson.Field
		HasCustomShortLinkCode  respjson.Field
		IsPublished             respjson.Field
		Name                    respjson.Field
		Pixel                   respjson.Field
		RedirectURL             respjson.Field
		ShortLinkCode           respjson.Field
		ShortLinkDesign         respjson.Field
		ShortURL                respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
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

// QR styling + campaign tags. New links include an empty design; legacy redirects
// may return null.
type ShortLinkNewResponseShortLinkDesign struct {
	Qr  ShortLinkNewResponseShortLinkDesignQr  `json:"qr" api:"nullable"`
	Utm ShortLinkNewResponseShortLinkDesignUtm `json:"utm" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Qr          respjson.Field
		Utm         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortLinkNewResponseShortLinkDesign) RawJSON() string { return r.JSON.raw }
func (r *ShortLinkNewResponseShortLinkDesign) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortLinkNewResponseShortLinkDesignQr struct {
	// Background color.
	BgColor string `json:"bgColor" api:"nullable"`
	// Finder (corner) dot shape: `square` or `dot`.
	CornerDotStyle string `json:"cornerDotStyle" api:"nullable"`
	// Finder (corner) square shape: `square`, `dot`, or `extra-rounded`.
	CornerSquareStyle string `json:"cornerSquareStyle" api:"nullable"`
	// Module shape: `square`, `rounded`, `dots`, `classy`, `classy-rounded`, or
	// `extra-rounded`.
	DotStyle string `json:"dotStyle" api:"nullable"`
	// Foreground (module) color.
	FgColor string `json:"fgColor" api:"nullable"`
	// Download frame: `square`, `rounded`, or `circle`.
	FrameShape string `json:"frameShape" api:"nullable"`
	// A two-stop gradient fill for the QR modules.
	Gradient ShortLinkNewResponseShortLinkDesignQrGradient `json:"gradient" api:"nullable"`
	// Base64 data URL of a centered logo overlay.
	Logo string `json:"logo" api:"nullable"`
	// Logo size as a fraction of the QR (0–0.4).
	LogoSizeRatio float64 `json:"logoSizeRatio" api:"nullable"`
	// The selected preset theme id, for re-highlighting it in a UI.
	Theme string `json:"theme" api:"nullable"`
	// When true, the background is exported transparent.
	TransparentBackground bool `json:"transparentBackground" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BgColor               respjson.Field
		CornerDotStyle        respjson.Field
		CornerSquareStyle     respjson.Field
		DotStyle              respjson.Field
		FgColor               respjson.Field
		FrameShape            respjson.Field
		Gradient              respjson.Field
		Logo                  respjson.Field
		LogoSizeRatio         respjson.Field
		Theme                 respjson.Field
		TransparentBackground respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortLinkNewResponseShortLinkDesignQr) RawJSON() string { return r.JSON.raw }
func (r *ShortLinkNewResponseShortLinkDesignQr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A two-stop gradient fill for the QR modules.
type ShortLinkNewResponseShortLinkDesignQrGradient struct {
	// Gradient end color.
	EndColor string `json:"endColor" api:"required"`
	// Gradient start color.
	StartColor string `json:"startColor" api:"required"`
	// Gradient type: `linear` or `radial`.
	Type string `json:"type" api:"required"`
	// Rotation in radians (linear gradients only).
	Rotation float64 `json:"rotation" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndColor    respjson.Field
		StartColor  respjson.Field
		Type        respjson.Field
		Rotation    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortLinkNewResponseShortLinkDesignQrGradient) RawJSON() string { return r.JSON.raw }
func (r *ShortLinkNewResponseShortLinkDesignQrGradient) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortLinkNewResponseShortLinkDesignUtm struct {
	// The `utm_campaign` value appended to the tracked short-link URL.
	Campaign string `json:"campaign" api:"nullable"`
	// The `utm_content` value.
	Content string `json:"content" api:"nullable"`
	// The `utm_medium` value.
	Medium string `json:"medium" api:"nullable"`
	// The `utm_source` value.
	Source string `json:"source" api:"nullable"`
	// The `utm_term` value.
	Term string `json:"term" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Campaign    respjson.Field
		Content     respjson.Field
		Medium      respjson.Field
		Source      respjson.Field
		Term        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortLinkNewResponseShortLinkDesignUtm) RawJSON() string { return r.JSON.raw }
func (r *ShortLinkNewResponseShortLinkDesignUtm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortLinkGetResponse struct {
	ID string `json:"id" api:"required"`
	// Organization id that owns this short link.
	AccountID string `json:"accountId" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	// Whether saved published design campaign tags are added to the destination when
	// enabled. Existing destination UTMs stay unchanged. Defaults to true for new
	// short links and false for existing links without a stored setting.
	ForwardUtmToDestination bool `json:"forwardUtmToDestination" api:"required"`
	// Any of "Disabled", "Enabled".
	Status ShortLinkGetResponseStatus `json:"status" api:"required"`
	// Whether a user selected this code instead of using a generated code.
	HasCustomShortLinkCode bool `json:"hasCustomShortLinkCode" api:"nullable"`
	// Whether this short link exists in the currently published version. An
	// unpublished short link does not resolve at the edge.
	IsPublished bool   `json:"isPublished" api:"nullable"`
	Name        string `json:"name" api:"nullable"`
	// The immutable short code embedded in the public URL (`/r/{pixel}` for new
	// links). Server-assigned.
	Pixel string `json:"pixel" api:"nullable"`
	// The destination URL this short link redirects to.
	RedirectURL string `json:"redirectUrl" api:"nullable"`
	// The public code embedded in the short-link URL.
	ShortLinkCode string `json:"shortLinkCode" api:"nullable"`
	// QR styling + campaign tags. New links include an empty design; legacy redirects
	// may return null.
	ShortLinkDesign ShortLinkGetResponseShortLinkDesign `json:"shortLinkDesign" api:"nullable"`
	// The public short-link URL that the QR encodes and callers share. New links use
	// `/r/{pixel}`; tracked event and campaign defaults are stored server-side. Also
	// resolves on branded custom domains configured for the account.
	ShortURL string `json:"shortUrl" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		AccountID               respjson.Field
		CreatedAt               respjson.Field
		ForwardUtmToDestination respjson.Field
		Status                  respjson.Field
		HasCustomShortLinkCode  respjson.Field
		IsPublished             respjson.Field
		Name                    respjson.Field
		Pixel                   respjson.Field
		RedirectURL             respjson.Field
		ShortLinkCode           respjson.Field
		ShortLinkDesign         respjson.Field
		ShortURL                respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
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

// QR styling + campaign tags. New links include an empty design; legacy redirects
// may return null.
type ShortLinkGetResponseShortLinkDesign struct {
	Qr  ShortLinkGetResponseShortLinkDesignQr  `json:"qr" api:"nullable"`
	Utm ShortLinkGetResponseShortLinkDesignUtm `json:"utm" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Qr          respjson.Field
		Utm         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortLinkGetResponseShortLinkDesign) RawJSON() string { return r.JSON.raw }
func (r *ShortLinkGetResponseShortLinkDesign) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortLinkGetResponseShortLinkDesignQr struct {
	// Background color.
	BgColor string `json:"bgColor" api:"nullable"`
	// Finder (corner) dot shape: `square` or `dot`.
	CornerDotStyle string `json:"cornerDotStyle" api:"nullable"`
	// Finder (corner) square shape: `square`, `dot`, or `extra-rounded`.
	CornerSquareStyle string `json:"cornerSquareStyle" api:"nullable"`
	// Module shape: `square`, `rounded`, `dots`, `classy`, `classy-rounded`, or
	// `extra-rounded`.
	DotStyle string `json:"dotStyle" api:"nullable"`
	// Foreground (module) color.
	FgColor string `json:"fgColor" api:"nullable"`
	// Download frame: `square`, `rounded`, or `circle`.
	FrameShape string `json:"frameShape" api:"nullable"`
	// A two-stop gradient fill for the QR modules.
	Gradient ShortLinkGetResponseShortLinkDesignQrGradient `json:"gradient" api:"nullable"`
	// Base64 data URL of a centered logo overlay.
	Logo string `json:"logo" api:"nullable"`
	// Logo size as a fraction of the QR (0–0.4).
	LogoSizeRatio float64 `json:"logoSizeRatio" api:"nullable"`
	// The selected preset theme id, for re-highlighting it in a UI.
	Theme string `json:"theme" api:"nullable"`
	// When true, the background is exported transparent.
	TransparentBackground bool `json:"transparentBackground" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BgColor               respjson.Field
		CornerDotStyle        respjson.Field
		CornerSquareStyle     respjson.Field
		DotStyle              respjson.Field
		FgColor               respjson.Field
		FrameShape            respjson.Field
		Gradient              respjson.Field
		Logo                  respjson.Field
		LogoSizeRatio         respjson.Field
		Theme                 respjson.Field
		TransparentBackground respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortLinkGetResponseShortLinkDesignQr) RawJSON() string { return r.JSON.raw }
func (r *ShortLinkGetResponseShortLinkDesignQr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A two-stop gradient fill for the QR modules.
type ShortLinkGetResponseShortLinkDesignQrGradient struct {
	// Gradient end color.
	EndColor string `json:"endColor" api:"required"`
	// Gradient start color.
	StartColor string `json:"startColor" api:"required"`
	// Gradient type: `linear` or `radial`.
	Type string `json:"type" api:"required"`
	// Rotation in radians (linear gradients only).
	Rotation float64 `json:"rotation" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndColor    respjson.Field
		StartColor  respjson.Field
		Type        respjson.Field
		Rotation    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortLinkGetResponseShortLinkDesignQrGradient) RawJSON() string { return r.JSON.raw }
func (r *ShortLinkGetResponseShortLinkDesignQrGradient) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortLinkGetResponseShortLinkDesignUtm struct {
	// The `utm_campaign` value appended to the tracked short-link URL.
	Campaign string `json:"campaign" api:"nullable"`
	// The `utm_content` value.
	Content string `json:"content" api:"nullable"`
	// The `utm_medium` value.
	Medium string `json:"medium" api:"nullable"`
	// The `utm_source` value.
	Source string `json:"source" api:"nullable"`
	// The `utm_term` value.
	Term string `json:"term" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Campaign    respjson.Field
		Content     respjson.Field
		Medium      respjson.Field
		Source      respjson.Field
		Term        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortLinkGetResponseShortLinkDesignUtm) RawJSON() string { return r.JSON.raw }
func (r *ShortLinkGetResponseShortLinkDesignUtm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortLinkUpdateResponse struct {
	ID string `json:"id" api:"required"`
	// Organization id that owns this short link.
	AccountID string `json:"accountId" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	// Whether saved published design campaign tags are added to the destination when
	// enabled. Existing destination UTMs stay unchanged. Defaults to true for new
	// short links and false for existing links without a stored setting.
	ForwardUtmToDestination bool `json:"forwardUtmToDestination" api:"required"`
	// Any of "Disabled", "Enabled".
	Status ShortLinkUpdateResponseStatus `json:"status" api:"required"`
	// Whether a user selected this code instead of using a generated code.
	HasCustomShortLinkCode bool `json:"hasCustomShortLinkCode" api:"nullable"`
	// Whether this short link exists in the currently published version. An
	// unpublished short link does not resolve at the edge.
	IsPublished bool   `json:"isPublished" api:"nullable"`
	Name        string `json:"name" api:"nullable"`
	// The immutable short code embedded in the public URL (`/r/{pixel}` for new
	// links). Server-assigned.
	Pixel string `json:"pixel" api:"nullable"`
	// The destination URL this short link redirects to.
	RedirectURL string `json:"redirectUrl" api:"nullable"`
	// The public code embedded in the short-link URL.
	ShortLinkCode string `json:"shortLinkCode" api:"nullable"`
	// QR styling + campaign tags. New links include an empty design; legacy redirects
	// may return null.
	ShortLinkDesign ShortLinkUpdateResponseShortLinkDesign `json:"shortLinkDesign" api:"nullable"`
	// The public short-link URL that the QR encodes and callers share. New links use
	// `/r/{pixel}`; tracked event and campaign defaults are stored server-side. Also
	// resolves on branded custom domains configured for the account.
	ShortURL string `json:"shortUrl" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		AccountID               respjson.Field
		CreatedAt               respjson.Field
		ForwardUtmToDestination respjson.Field
		Status                  respjson.Field
		HasCustomShortLinkCode  respjson.Field
		IsPublished             respjson.Field
		Name                    respjson.Field
		Pixel                   respjson.Field
		RedirectURL             respjson.Field
		ShortLinkCode           respjson.Field
		ShortLinkDesign         respjson.Field
		ShortURL                respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
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

// QR styling + campaign tags. New links include an empty design; legacy redirects
// may return null.
type ShortLinkUpdateResponseShortLinkDesign struct {
	Qr  ShortLinkUpdateResponseShortLinkDesignQr  `json:"qr" api:"nullable"`
	Utm ShortLinkUpdateResponseShortLinkDesignUtm `json:"utm" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Qr          respjson.Field
		Utm         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortLinkUpdateResponseShortLinkDesign) RawJSON() string { return r.JSON.raw }
func (r *ShortLinkUpdateResponseShortLinkDesign) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortLinkUpdateResponseShortLinkDesignQr struct {
	// Background color.
	BgColor string `json:"bgColor" api:"nullable"`
	// Finder (corner) dot shape: `square` or `dot`.
	CornerDotStyle string `json:"cornerDotStyle" api:"nullable"`
	// Finder (corner) square shape: `square`, `dot`, or `extra-rounded`.
	CornerSquareStyle string `json:"cornerSquareStyle" api:"nullable"`
	// Module shape: `square`, `rounded`, `dots`, `classy`, `classy-rounded`, or
	// `extra-rounded`.
	DotStyle string `json:"dotStyle" api:"nullable"`
	// Foreground (module) color.
	FgColor string `json:"fgColor" api:"nullable"`
	// Download frame: `square`, `rounded`, or `circle`.
	FrameShape string `json:"frameShape" api:"nullable"`
	// A two-stop gradient fill for the QR modules.
	Gradient ShortLinkUpdateResponseShortLinkDesignQrGradient `json:"gradient" api:"nullable"`
	// Base64 data URL of a centered logo overlay.
	Logo string `json:"logo" api:"nullable"`
	// Logo size as a fraction of the QR (0–0.4).
	LogoSizeRatio float64 `json:"logoSizeRatio" api:"nullable"`
	// The selected preset theme id, for re-highlighting it in a UI.
	Theme string `json:"theme" api:"nullable"`
	// When true, the background is exported transparent.
	TransparentBackground bool `json:"transparentBackground" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BgColor               respjson.Field
		CornerDotStyle        respjson.Field
		CornerSquareStyle     respjson.Field
		DotStyle              respjson.Field
		FgColor               respjson.Field
		FrameShape            respjson.Field
		Gradient              respjson.Field
		Logo                  respjson.Field
		LogoSizeRatio         respjson.Field
		Theme                 respjson.Field
		TransparentBackground respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortLinkUpdateResponseShortLinkDesignQr) RawJSON() string { return r.JSON.raw }
func (r *ShortLinkUpdateResponseShortLinkDesignQr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A two-stop gradient fill for the QR modules.
type ShortLinkUpdateResponseShortLinkDesignQrGradient struct {
	// Gradient end color.
	EndColor string `json:"endColor" api:"required"`
	// Gradient start color.
	StartColor string `json:"startColor" api:"required"`
	// Gradient type: `linear` or `radial`.
	Type string `json:"type" api:"required"`
	// Rotation in radians (linear gradients only).
	Rotation float64 `json:"rotation" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndColor    respjson.Field
		StartColor  respjson.Field
		Type        respjson.Field
		Rotation    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortLinkUpdateResponseShortLinkDesignQrGradient) RawJSON() string { return r.JSON.raw }
func (r *ShortLinkUpdateResponseShortLinkDesignQrGradient) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortLinkUpdateResponseShortLinkDesignUtm struct {
	// The `utm_campaign` value appended to the tracked short-link URL.
	Campaign string `json:"campaign" api:"nullable"`
	// The `utm_content` value.
	Content string `json:"content" api:"nullable"`
	// The `utm_medium` value.
	Medium string `json:"medium" api:"nullable"`
	// The `utm_source` value.
	Source string `json:"source" api:"nullable"`
	// The `utm_term` value.
	Term string `json:"term" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Campaign    respjson.Field
		Content     respjson.Field
		Medium      respjson.Field
		Source      respjson.Field
		Term        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortLinkUpdateResponseShortLinkDesignUtm) RawJSON() string { return r.JSON.raw }
func (r *ShortLinkUpdateResponseShortLinkDesignUtm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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

type ShortLinkCloneResponse struct {
	ID string `json:"id" api:"required"`
	// Organization id that owns this short link.
	AccountID string `json:"accountId" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	// Whether saved published design campaign tags are added to the destination when
	// enabled. Existing destination UTMs stay unchanged. Defaults to true for new
	// short links and false for existing links without a stored setting.
	ForwardUtmToDestination bool `json:"forwardUtmToDestination" api:"required"`
	// Any of "Disabled", "Enabled".
	Status ShortLinkCloneResponseStatus `json:"status" api:"required"`
	// Whether a user selected this code instead of using a generated code.
	HasCustomShortLinkCode bool `json:"hasCustomShortLinkCode" api:"nullable"`
	// Whether this short link exists in the currently published version. An
	// unpublished short link does not resolve at the edge.
	IsPublished bool   `json:"isPublished" api:"nullable"`
	Name        string `json:"name" api:"nullable"`
	// The immutable short code embedded in the public URL (`/r/{pixel}` for new
	// links). Server-assigned.
	Pixel string `json:"pixel" api:"nullable"`
	// The destination URL this short link redirects to.
	RedirectURL string `json:"redirectUrl" api:"nullable"`
	// The public code embedded in the short-link URL.
	ShortLinkCode string `json:"shortLinkCode" api:"nullable"`
	// QR styling + campaign tags. New links include an empty design; legacy redirects
	// may return null.
	ShortLinkDesign ShortLinkCloneResponseShortLinkDesign `json:"shortLinkDesign" api:"nullable"`
	// The public short-link URL that the QR encodes and callers share. New links use
	// `/r/{pixel}`; tracked event and campaign defaults are stored server-side. Also
	// resolves on branded custom domains configured for the account.
	ShortURL string `json:"shortUrl" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		AccountID               respjson.Field
		CreatedAt               respjson.Field
		ForwardUtmToDestination respjson.Field
		Status                  respjson.Field
		HasCustomShortLinkCode  respjson.Field
		IsPublished             respjson.Field
		Name                    respjson.Field
		Pixel                   respjson.Field
		RedirectURL             respjson.Field
		ShortLinkCode           respjson.Field
		ShortLinkDesign         respjson.Field
		ShortURL                respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortLinkCloneResponse) RawJSON() string { return r.JSON.raw }
func (r *ShortLinkCloneResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortLinkCloneResponseStatus string

const (
	ShortLinkCloneResponseStatusDisabled ShortLinkCloneResponseStatus = "Disabled"
	ShortLinkCloneResponseStatusEnabled  ShortLinkCloneResponseStatus = "Enabled"
)

// QR styling + campaign tags. New links include an empty design; legacy redirects
// may return null.
type ShortLinkCloneResponseShortLinkDesign struct {
	Qr  ShortLinkCloneResponseShortLinkDesignQr  `json:"qr" api:"nullable"`
	Utm ShortLinkCloneResponseShortLinkDesignUtm `json:"utm" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Qr          respjson.Field
		Utm         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortLinkCloneResponseShortLinkDesign) RawJSON() string { return r.JSON.raw }
func (r *ShortLinkCloneResponseShortLinkDesign) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortLinkCloneResponseShortLinkDesignQr struct {
	// Background color.
	BgColor string `json:"bgColor" api:"nullable"`
	// Finder (corner) dot shape: `square` or `dot`.
	CornerDotStyle string `json:"cornerDotStyle" api:"nullable"`
	// Finder (corner) square shape: `square`, `dot`, or `extra-rounded`.
	CornerSquareStyle string `json:"cornerSquareStyle" api:"nullable"`
	// Module shape: `square`, `rounded`, `dots`, `classy`, `classy-rounded`, or
	// `extra-rounded`.
	DotStyle string `json:"dotStyle" api:"nullable"`
	// Foreground (module) color.
	FgColor string `json:"fgColor" api:"nullable"`
	// Download frame: `square`, `rounded`, or `circle`.
	FrameShape string `json:"frameShape" api:"nullable"`
	// A two-stop gradient fill for the QR modules.
	Gradient ShortLinkCloneResponseShortLinkDesignQrGradient `json:"gradient" api:"nullable"`
	// Base64 data URL of a centered logo overlay.
	Logo string `json:"logo" api:"nullable"`
	// Logo size as a fraction of the QR (0–0.4).
	LogoSizeRatio float64 `json:"logoSizeRatio" api:"nullable"`
	// The selected preset theme id, for re-highlighting it in a UI.
	Theme string `json:"theme" api:"nullable"`
	// When true, the background is exported transparent.
	TransparentBackground bool `json:"transparentBackground" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BgColor               respjson.Field
		CornerDotStyle        respjson.Field
		CornerSquareStyle     respjson.Field
		DotStyle              respjson.Field
		FgColor               respjson.Field
		FrameShape            respjson.Field
		Gradient              respjson.Field
		Logo                  respjson.Field
		LogoSizeRatio         respjson.Field
		Theme                 respjson.Field
		TransparentBackground respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortLinkCloneResponseShortLinkDesignQr) RawJSON() string { return r.JSON.raw }
func (r *ShortLinkCloneResponseShortLinkDesignQr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A two-stop gradient fill for the QR modules.
type ShortLinkCloneResponseShortLinkDesignQrGradient struct {
	// Gradient end color.
	EndColor string `json:"endColor" api:"required"`
	// Gradient start color.
	StartColor string `json:"startColor" api:"required"`
	// Gradient type: `linear` or `radial`.
	Type string `json:"type" api:"required"`
	// Rotation in radians (linear gradients only).
	Rotation float64 `json:"rotation" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndColor    respjson.Field
		StartColor  respjson.Field
		Type        respjson.Field
		Rotation    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortLinkCloneResponseShortLinkDesignQrGradient) RawJSON() string { return r.JSON.raw }
func (r *ShortLinkCloneResponseShortLinkDesignQrGradient) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortLinkCloneResponseShortLinkDesignUtm struct {
	// The `utm_campaign` value appended to the tracked short-link URL.
	Campaign string `json:"campaign" api:"nullable"`
	// The `utm_content` value.
	Content string `json:"content" api:"nullable"`
	// The `utm_medium` value.
	Medium string `json:"medium" api:"nullable"`
	// The `utm_source` value.
	Source string `json:"source" api:"nullable"`
	// The `utm_term` value.
	Term string `json:"term" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Campaign    respjson.Field
		Content     respjson.Field
		Medium      respjson.Field
		Source      respjson.Field
		Term        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShortLinkCloneResponseShortLinkDesignUtm) RawJSON() string { return r.JSON.raw }
func (r *ShortLinkCloneResponseShortLinkDesignUtm) UnmarshalJSON(data []byte) error {
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
	// Case-insensitive search across name, destination URL, and public code.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
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
	// Optional custom code in the shared public short-link namespace.
	Code param.Opt[string] `json:"code,omitzero"`
	// Human-readable name. Also sent as the tracked event name on every click/scan.
	Name param.Opt[string] `json:"name,omitzero"`
	// Destination URL the short link redirects to. Must be a valid URL.
	RedirectURL param.Opt[string] `json:"redirectUrl,omitzero" format:"uri"`
	// When enabled, add saved published design campaign tags to the destination after
	// publishing. Defaults to true when omitted; null is rejected.
	ForwardUtmToDestination param.Opt[bool] `json:"forwardUtmToDestination,omitzero"`
	// QR code visual styling.
	Qr ShortLinkNewParamsQr `json:"qr,omitzero"`
	// Campaign / UTM tags appended to the tracked short-link URL.
	Utm ShortLinkNewParamsUtm `json:"utm,omitzero"`
	paramObj
}

func (r ShortLinkNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ShortLinkNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ShortLinkNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// QR code visual styling.
type ShortLinkNewParamsQr struct {
	// Background color.
	BgColor param.Opt[string] `json:"bgColor,omitzero"`
	// Foreground (module) color.
	FgColor param.Opt[string] `json:"fgColor,omitzero"`
	// Base64 data URL of a centered logo overlay.
	Logo param.Opt[string] `json:"logo,omitzero"`
	// Logo size as a fraction of the QR.
	LogoSizeRatio param.Opt[float64] `json:"logoSizeRatio,omitzero"`
	// The selected preset theme id, for re-highlighting it in a UI.
	Theme param.Opt[string] `json:"theme,omitzero"`
	// When true, the background is exported transparent.
	TransparentBackground param.Opt[bool] `json:"transparentBackground,omitzero"`
	// Finder (corner) dot shape.
	//
	// Any of "square", "dot".
	CornerDotStyle string `json:"cornerDotStyle,omitzero"`
	// Finder (corner) square shape.
	//
	// Any of "square", "dot", "extra-rounded".
	CornerSquareStyle string `json:"cornerSquareStyle,omitzero"`
	// Module shape.
	//
	// Any of "square", "rounded", "dots", "classy", "classy-rounded", "extra-rounded".
	DotStyle string `json:"dotStyle,omitzero"`
	// Download frame.
	//
	// Any of "square", "rounded", "circle".
	FrameShape string `json:"frameShape,omitzero"`
	// A two-stop gradient fill for the QR modules.
	Gradient ShortLinkNewParamsQrGradient `json:"gradient,omitzero"`
	paramObj
}

func (r ShortLinkNewParamsQr) MarshalJSON() (data []byte, err error) {
	type shadow ShortLinkNewParamsQr
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ShortLinkNewParamsQr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ShortLinkNewParamsQr](
		"cornerDotStyle", "square", "dot",
	)
	apijson.RegisterFieldValidator[ShortLinkNewParamsQr](
		"cornerSquareStyle", "square", "dot", "extra-rounded",
	)
	apijson.RegisterFieldValidator[ShortLinkNewParamsQr](
		"dotStyle", "square", "rounded", "dots", "classy", "classy-rounded", "extra-rounded",
	)
	apijson.RegisterFieldValidator[ShortLinkNewParamsQr](
		"frameShape", "square", "rounded", "circle",
	)
}

// A two-stop gradient fill for the QR modules.
//
// The properties EndColor, StartColor, Type are required.
type ShortLinkNewParamsQrGradient struct {
	// Gradient end color.
	EndColor string `json:"endColor" api:"required"`
	// Gradient start color.
	StartColor string `json:"startColor" api:"required"`
	// Gradient type.
	//
	// Any of "linear", "radial".
	Type string `json:"type,omitzero" api:"required"`
	// Rotation in radians (linear gradients only).
	Rotation param.Opt[float64] `json:"rotation,omitzero"`
	paramObj
}

func (r ShortLinkNewParamsQrGradient) MarshalJSON() (data []byte, err error) {
	type shadow ShortLinkNewParamsQrGradient
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ShortLinkNewParamsQrGradient) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ShortLinkNewParamsQrGradient](
		"type", "linear", "radial",
	)
}

// Campaign / UTM tags appended to the tracked short-link URL.
type ShortLinkNewParamsUtm struct {
	// The `utm_campaign` value appended to the tracked short-link URL.
	Campaign param.Opt[string] `json:"campaign,omitzero"`
	// The `utm_content` value.
	Content param.Opt[string] `json:"content,omitzero"`
	// The `utm_medium` value.
	Medium param.Opt[string] `json:"medium,omitzero"`
	// The `utm_source` value.
	Source param.Opt[string] `json:"source,omitzero"`
	// The `utm_term` value.
	Term param.Opt[string] `json:"term,omitzero"`
	paramObj
}

func (r ShortLinkNewParamsUtm) MarshalJSON() (data []byte, err error) {
	type shadow ShortLinkNewParamsUtm
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ShortLinkNewParamsUtm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortLinkUpdateParams struct {
	Code param.Opt[string] `json:"code,omitzero"`
	Name param.Opt[string] `json:"name,omitzero"`
	// Destination URL the short link redirects to. Must be a valid URL. Send `null` to
	// clear it.
	RedirectURL param.Opt[string] `json:"redirectUrl,omitzero" format:"uri"`
	// When enabled, add saved published design campaign tags to the destination after
	// publishing. Omit to leave unchanged; null is rejected.
	ForwardUtmToDestination param.Opt[bool]         `json:"forwardUtmToDestination,omitzero"`
	Qr                      ShortLinkUpdateParamsQr `json:"qr,omitzero"`
	// Whether the short link resolves at the edge. Send `Enabled` or `Disabled`;
	// `null` is rejected since storage cannot represent it.
	//
	// Any of "Disabled", "Enabled".
	Status ShortLinkUpdateParamsStatus `json:"status,omitzero"`
	Utm    ShortLinkUpdateParamsUtm    `json:"utm,omitzero"`
	paramObj
}

func (r ShortLinkUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ShortLinkUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ShortLinkUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShortLinkUpdateParamsQr struct {
	// Background color.
	BgColor param.Opt[string] `json:"bgColor,omitzero"`
	// Foreground (module) color.
	FgColor param.Opt[string] `json:"fgColor,omitzero"`
	// Base64 data URL of a centered logo overlay.
	Logo param.Opt[string] `json:"logo,omitzero"`
	// Logo size as a fraction of the QR.
	LogoSizeRatio param.Opt[float64] `json:"logoSizeRatio,omitzero"`
	// The selected preset theme id, for re-highlighting it in a UI.
	Theme param.Opt[string] `json:"theme,omitzero"`
	// When true, the background is exported transparent.
	TransparentBackground param.Opt[bool] `json:"transparentBackground,omitzero"`
	// Finder (corner) dot shape.
	//
	// Any of "square", "dot".
	CornerDotStyle string `json:"cornerDotStyle,omitzero"`
	// Finder (corner) square shape.
	//
	// Any of "square", "dot", "extra-rounded".
	CornerSquareStyle string `json:"cornerSquareStyle,omitzero"`
	// Module shape.
	//
	// Any of "square", "rounded", "dots", "classy", "classy-rounded", "extra-rounded".
	DotStyle string `json:"dotStyle,omitzero"`
	// Download frame.
	//
	// Any of "square", "rounded", "circle".
	FrameShape string `json:"frameShape,omitzero"`
	// A two-stop gradient fill for the QR modules.
	Gradient ShortLinkUpdateParamsQrGradient `json:"gradient,omitzero"`
	paramObj
}

func (r ShortLinkUpdateParamsQr) MarshalJSON() (data []byte, err error) {
	type shadow ShortLinkUpdateParamsQr
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ShortLinkUpdateParamsQr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ShortLinkUpdateParamsQr](
		"cornerDotStyle", "square", "dot",
	)
	apijson.RegisterFieldValidator[ShortLinkUpdateParamsQr](
		"cornerSquareStyle", "square", "dot", "extra-rounded",
	)
	apijson.RegisterFieldValidator[ShortLinkUpdateParamsQr](
		"dotStyle", "square", "rounded", "dots", "classy", "classy-rounded", "extra-rounded",
	)
	apijson.RegisterFieldValidator[ShortLinkUpdateParamsQr](
		"frameShape", "square", "rounded", "circle",
	)
}

// A two-stop gradient fill for the QR modules.
//
// The properties EndColor, StartColor, Type are required.
type ShortLinkUpdateParamsQrGradient struct {
	// Gradient end color.
	EndColor string `json:"endColor" api:"required"`
	// Gradient start color.
	StartColor string `json:"startColor" api:"required"`
	// Gradient type.
	//
	// Any of "linear", "radial".
	Type string `json:"type,omitzero" api:"required"`
	// Rotation in radians (linear gradients only).
	Rotation param.Opt[float64] `json:"rotation,omitzero"`
	paramObj
}

func (r ShortLinkUpdateParamsQrGradient) MarshalJSON() (data []byte, err error) {
	type shadow ShortLinkUpdateParamsQrGradient
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ShortLinkUpdateParamsQrGradient) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ShortLinkUpdateParamsQrGradient](
		"type", "linear", "radial",
	)
}

// Whether the short link resolves at the edge. Send `Enabled` or `Disabled`;
// `null` is rejected since storage cannot represent it.
type ShortLinkUpdateParamsStatus string

const (
	ShortLinkUpdateParamsStatusDisabled ShortLinkUpdateParamsStatus = "Disabled"
	ShortLinkUpdateParamsStatusEnabled  ShortLinkUpdateParamsStatus = "Enabled"
)

type ShortLinkUpdateParamsUtm struct {
	// The `utm_campaign` value appended to the tracked short-link URL.
	Campaign param.Opt[string] `json:"campaign,omitzero"`
	// The `utm_content` value.
	Content param.Opt[string] `json:"content,omitzero"`
	// The `utm_medium` value.
	Medium param.Opt[string] `json:"medium,omitzero"`
	// The `utm_source` value.
	Source param.Opt[string] `json:"source,omitzero"`
	// The `utm_term` value.
	Term param.Opt[string] `json:"term,omitzero"`
	paramObj
}

func (r ShortLinkUpdateParamsUtm) MarshalJSON() (data []byte, err error) {
	type shadow ShortLinkUpdateParamsUtm
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ShortLinkUpdateParamsUtm) UnmarshalJSON(data []byte) error {
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
