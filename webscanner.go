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
	"github.com/with-ours/platform-sdk-go/internal/requestconfig"
	"github.com/with-ours/platform-sdk-go/option"
	"github.com/with-ours/platform-sdk-go/packages/param"
	"github.com/with-ours/platform-sdk-go/packages/respjson"
)

// WebScannerService contains methods and other services that help with interacting
// with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebScannerService] method instead.
type WebScannerService struct {
	Options []option.RequestOption
}

// NewWebScannerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebScannerService(opts ...option.RequestOption) (r WebScannerService) {
	r = WebScannerService{}
	r.Options = opts
	return
}

// List every web scanner for this account. Not paginated — accounts have a small
// number of scanners in practice, so the response always fits in a single page.
// Requires scope: webScanner:list
func (r *WebScannerService) List(ctx context.Context, opts ...option.RequestOption) (res *WebScannerListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/web-scanners"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Create a new web scanner for a root domain. A first scan is enqueued
// automatically after creation on a best-effort basis. `rootDomain` is required;
// missing, empty, or malformed values are rejected as HTTP 400. Everything else
// falls back to defaults (`status: Enabled`, `urlLimit: 100`, no excluded
// patterns, no extra seed URLs). The returned entity is the created scanner row
// and may not yet reflect async scan-state changes. Requires scope:
// webScanner:create
func (r *WebScannerService) New(ctx context.Context, body WebScannerNewParams, opts ...option.RequestOption) (res *WebScannerNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/web-scanners"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Find a single web scanner by ID. Requires scope: webScanner:find
func (r *WebScannerService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *WebScannerGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/web-scanners/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Partially update a web scanner. Only the fields you send are changed; omitted
// fields keep their current value. List-valued fields (`excludedPatterns`,
// `includedUrls`) are replaced wholesale when sent. If `rootDomain` is provided
// and malformed, the request is rejected as HTTP 400. Use
// `POST /rest/v1/web-scanners/{id}/trigger` to start a new scan after edits.
// Requires scope: webScanner:update
func (r *WebScannerService) Update(ctx context.Context, id string, body WebScannerUpdateParams, opts ...option.RequestOption) (res *WebScannerUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/web-scanners/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete a web scanner. Associated suppression rules are deleted in the same
// operation. Requires scope: webScanner:delete
func (r *WebScannerService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *WebScannerDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/web-scanners/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Manually kick off a new scan for this web scanner. The request body is empty (or
// `{}`). A successful response means the request was accepted; because the scan
// starts asynchronously, the returned entity may still reflect pre-trigger values
// for fields like `scanStatus` and `lastScanStartedAt`. The trigger is
// rate-limited: a 409 is returned if another scan is already in flight, the
// per-account cooldown has not elapsed, or the request was otherwise rejected; the
// reason is in the response `error` field. Requires scope: webScanner:trigger
func (r *WebScannerService) Trigger(ctx context.Context, id string, opts ...option.RequestOption) (res *WebScannerTriggerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/web-scanners/%s/trigger", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type WebScannerListResponse struct {
	Entities []WebScannerListResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerListResponse) RawJSON() string { return r.JSON.raw }
func (r *WebScannerListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerListResponseEntity struct {
	ID         string `json:"id" api:"required"`
	AccountID  string `json:"accountId" api:"required"`
	RootDomain string `json:"rootDomain" api:"required"`
	// Any of "idle", "scanning".
	ScanStatus string `json:"scanStatus" api:"required"`
	// Any of "Disabled", "Enabled".
	Status                      string   `json:"status" api:"required"`
	CreatedAt                   string   `json:"createdAt" api:"nullable"`
	ExcludedPatterns            []string `json:"excludedPatterns" api:"nullable"`
	IncludedURLs                []string `json:"includedUrls" api:"nullable"`
	LastRunCookieCount          float64  `json:"lastRunCookieCount" api:"nullable"`
	LastRunHighRiskRequestCount float64  `json:"lastRunHighRiskRequestCount" api:"nullable"`
	LastRunRequestCount         float64  `json:"lastRunRequestCount" api:"nullable"`
	LastRunSuccessURLCount      float64  `json:"lastRunSuccessUrlCount" api:"nullable"`
	LastScannedAt               string   `json:"lastScannedAt" api:"nullable"`
	LastScanStartedAt           string   `json:"lastScanStartedAt" api:"nullable"`
	Name                        string   `json:"name" api:"nullable"`
	UpdatedAt                   string   `json:"updatedAt" api:"nullable"`
	URLLimit                    float64  `json:"urlLimit" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                          respjson.Field
		AccountID                   respjson.Field
		RootDomain                  respjson.Field
		ScanStatus                  respjson.Field
		Status                      respjson.Field
		CreatedAt                   respjson.Field
		ExcludedPatterns            respjson.Field
		IncludedURLs                respjson.Field
		LastRunCookieCount          respjson.Field
		LastRunHighRiskRequestCount respjson.Field
		LastRunRequestCount         respjson.Field
		LastRunSuccessURLCount      respjson.Field
		LastScannedAt               respjson.Field
		LastScanStartedAt           respjson.Field
		Name                        respjson.Field
		UpdatedAt                   respjson.Field
		URLLimit                    respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerListResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *WebScannerListResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerNewResponse struct {
	ID         string `json:"id" api:"required"`
	AccountID  string `json:"accountId" api:"required"`
	RootDomain string `json:"rootDomain" api:"required"`
	// Any of "idle", "scanning".
	ScanStatus WebScannerNewResponseScanStatus `json:"scanStatus" api:"required"`
	// Any of "Disabled", "Enabled".
	Status                      WebScannerNewResponseStatus `json:"status" api:"required"`
	CreatedAt                   string                      `json:"createdAt" api:"nullable"`
	ExcludedPatterns            []string                    `json:"excludedPatterns" api:"nullable"`
	IncludedURLs                []string                    `json:"includedUrls" api:"nullable"`
	LastRunCookieCount          float64                     `json:"lastRunCookieCount" api:"nullable"`
	LastRunHighRiskRequestCount float64                     `json:"lastRunHighRiskRequestCount" api:"nullable"`
	LastRunRequestCount         float64                     `json:"lastRunRequestCount" api:"nullable"`
	LastRunSuccessURLCount      float64                     `json:"lastRunSuccessUrlCount" api:"nullable"`
	LastScannedAt               string                      `json:"lastScannedAt" api:"nullable"`
	LastScanStartedAt           string                      `json:"lastScanStartedAt" api:"nullable"`
	Name                        string                      `json:"name" api:"nullable"`
	UpdatedAt                   string                      `json:"updatedAt" api:"nullable"`
	URLLimit                    float64                     `json:"urlLimit" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                          respjson.Field
		AccountID                   respjson.Field
		RootDomain                  respjson.Field
		ScanStatus                  respjson.Field
		Status                      respjson.Field
		CreatedAt                   respjson.Field
		ExcludedPatterns            respjson.Field
		IncludedURLs                respjson.Field
		LastRunCookieCount          respjson.Field
		LastRunHighRiskRequestCount respjson.Field
		LastRunRequestCount         respjson.Field
		LastRunSuccessURLCount      respjson.Field
		LastScannedAt               respjson.Field
		LastScanStartedAt           respjson.Field
		Name                        respjson.Field
		UpdatedAt                   respjson.Field
		URLLimit                    respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WebScannerNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerNewResponseScanStatus string

const (
	WebScannerNewResponseScanStatusIdle     WebScannerNewResponseScanStatus = "idle"
	WebScannerNewResponseScanStatusScanning WebScannerNewResponseScanStatus = "scanning"
)

type WebScannerNewResponseStatus string

const (
	WebScannerNewResponseStatusDisabled WebScannerNewResponseStatus = "Disabled"
	WebScannerNewResponseStatusEnabled  WebScannerNewResponseStatus = "Enabled"
)

type WebScannerGetResponse struct {
	ID         string `json:"id" api:"required"`
	AccountID  string `json:"accountId" api:"required"`
	RootDomain string `json:"rootDomain" api:"required"`
	// Any of "idle", "scanning".
	ScanStatus WebScannerGetResponseScanStatus `json:"scanStatus" api:"required"`
	// Any of "Disabled", "Enabled".
	Status                      WebScannerGetResponseStatus `json:"status" api:"required"`
	CreatedAt                   string                      `json:"createdAt" api:"nullable"`
	ExcludedPatterns            []string                    `json:"excludedPatterns" api:"nullable"`
	IncludedURLs                []string                    `json:"includedUrls" api:"nullable"`
	LastRunCookieCount          float64                     `json:"lastRunCookieCount" api:"nullable"`
	LastRunHighRiskRequestCount float64                     `json:"lastRunHighRiskRequestCount" api:"nullable"`
	LastRunRequestCount         float64                     `json:"lastRunRequestCount" api:"nullable"`
	LastRunSuccessURLCount      float64                     `json:"lastRunSuccessUrlCount" api:"nullable"`
	LastScannedAt               string                      `json:"lastScannedAt" api:"nullable"`
	LastScanStartedAt           string                      `json:"lastScanStartedAt" api:"nullable"`
	Name                        string                      `json:"name" api:"nullable"`
	UpdatedAt                   string                      `json:"updatedAt" api:"nullable"`
	URLLimit                    float64                     `json:"urlLimit" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                          respjson.Field
		AccountID                   respjson.Field
		RootDomain                  respjson.Field
		ScanStatus                  respjson.Field
		Status                      respjson.Field
		CreatedAt                   respjson.Field
		ExcludedPatterns            respjson.Field
		IncludedURLs                respjson.Field
		LastRunCookieCount          respjson.Field
		LastRunHighRiskRequestCount respjson.Field
		LastRunRequestCount         respjson.Field
		LastRunSuccessURLCount      respjson.Field
		LastScannedAt               respjson.Field
		LastScanStartedAt           respjson.Field
		Name                        respjson.Field
		UpdatedAt                   respjson.Field
		URLLimit                    respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WebScannerGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerGetResponseScanStatus string

const (
	WebScannerGetResponseScanStatusIdle     WebScannerGetResponseScanStatus = "idle"
	WebScannerGetResponseScanStatusScanning WebScannerGetResponseScanStatus = "scanning"
)

type WebScannerGetResponseStatus string

const (
	WebScannerGetResponseStatusDisabled WebScannerGetResponseStatus = "Disabled"
	WebScannerGetResponseStatusEnabled  WebScannerGetResponseStatus = "Enabled"
)

type WebScannerUpdateResponse struct {
	ID         string `json:"id" api:"required"`
	AccountID  string `json:"accountId" api:"required"`
	RootDomain string `json:"rootDomain" api:"required"`
	// Any of "idle", "scanning".
	ScanStatus WebScannerUpdateResponseScanStatus `json:"scanStatus" api:"required"`
	// Any of "Disabled", "Enabled".
	Status                      WebScannerUpdateResponseStatus `json:"status" api:"required"`
	CreatedAt                   string                         `json:"createdAt" api:"nullable"`
	ExcludedPatterns            []string                       `json:"excludedPatterns" api:"nullable"`
	IncludedURLs                []string                       `json:"includedUrls" api:"nullable"`
	LastRunCookieCount          float64                        `json:"lastRunCookieCount" api:"nullable"`
	LastRunHighRiskRequestCount float64                        `json:"lastRunHighRiskRequestCount" api:"nullable"`
	LastRunRequestCount         float64                        `json:"lastRunRequestCount" api:"nullable"`
	LastRunSuccessURLCount      float64                        `json:"lastRunSuccessUrlCount" api:"nullable"`
	LastScannedAt               string                         `json:"lastScannedAt" api:"nullable"`
	LastScanStartedAt           string                         `json:"lastScanStartedAt" api:"nullable"`
	Name                        string                         `json:"name" api:"nullable"`
	UpdatedAt                   string                         `json:"updatedAt" api:"nullable"`
	URLLimit                    float64                        `json:"urlLimit" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                          respjson.Field
		AccountID                   respjson.Field
		RootDomain                  respjson.Field
		ScanStatus                  respjson.Field
		Status                      respjson.Field
		CreatedAt                   respjson.Field
		ExcludedPatterns            respjson.Field
		IncludedURLs                respjson.Field
		LastRunCookieCount          respjson.Field
		LastRunHighRiskRequestCount respjson.Field
		LastRunRequestCount         respjson.Field
		LastRunSuccessURLCount      respjson.Field
		LastScannedAt               respjson.Field
		LastScanStartedAt           respjson.Field
		Name                        respjson.Field
		UpdatedAt                   respjson.Field
		URLLimit                    respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *WebScannerUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerUpdateResponseScanStatus string

const (
	WebScannerUpdateResponseScanStatusIdle     WebScannerUpdateResponseScanStatus = "idle"
	WebScannerUpdateResponseScanStatusScanning WebScannerUpdateResponseScanStatus = "scanning"
)

type WebScannerUpdateResponseStatus string

const (
	WebScannerUpdateResponseStatusDisabled WebScannerUpdateResponseStatus = "Disabled"
	WebScannerUpdateResponseStatusEnabled  WebScannerUpdateResponseStatus = "Enabled"
)

type WebScannerDeleteResponse struct {
	// The id of the web scanner that was deleted.
	ID string `json:"id" api:"required"`
	// True when the scanner and its rules were deleted.
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
func (r WebScannerDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *WebScannerDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerTriggerResponse struct {
	ID         string `json:"id" api:"required"`
	AccountID  string `json:"accountId" api:"required"`
	RootDomain string `json:"rootDomain" api:"required"`
	// Any of "idle", "scanning".
	ScanStatus WebScannerTriggerResponseScanStatus `json:"scanStatus" api:"required"`
	// Any of "Disabled", "Enabled".
	Status                      WebScannerTriggerResponseStatus `json:"status" api:"required"`
	CreatedAt                   string                          `json:"createdAt" api:"nullable"`
	ExcludedPatterns            []string                        `json:"excludedPatterns" api:"nullable"`
	IncludedURLs                []string                        `json:"includedUrls" api:"nullable"`
	LastRunCookieCount          float64                         `json:"lastRunCookieCount" api:"nullable"`
	LastRunHighRiskRequestCount float64                         `json:"lastRunHighRiskRequestCount" api:"nullable"`
	LastRunRequestCount         float64                         `json:"lastRunRequestCount" api:"nullable"`
	LastRunSuccessURLCount      float64                         `json:"lastRunSuccessUrlCount" api:"nullable"`
	LastScannedAt               string                          `json:"lastScannedAt" api:"nullable"`
	LastScanStartedAt           string                          `json:"lastScanStartedAt" api:"nullable"`
	Name                        string                          `json:"name" api:"nullable"`
	UpdatedAt                   string                          `json:"updatedAt" api:"nullable"`
	URLLimit                    float64                         `json:"urlLimit" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                          respjson.Field
		AccountID                   respjson.Field
		RootDomain                  respjson.Field
		ScanStatus                  respjson.Field
		Status                      respjson.Field
		CreatedAt                   respjson.Field
		ExcludedPatterns            respjson.Field
		IncludedURLs                respjson.Field
		LastRunCookieCount          respjson.Field
		LastRunHighRiskRequestCount respjson.Field
		LastRunRequestCount         respjson.Field
		LastRunSuccessURLCount      respjson.Field
		LastScannedAt               respjson.Field
		LastScanStartedAt           respjson.Field
		Name                        respjson.Field
		UpdatedAt                   respjson.Field
		URLLimit                    respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerTriggerResponse) RawJSON() string { return r.JSON.raw }
func (r *WebScannerTriggerResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerTriggerResponseScanStatus string

const (
	WebScannerTriggerResponseScanStatusIdle     WebScannerTriggerResponseScanStatus = "idle"
	WebScannerTriggerResponseScanStatusScanning WebScannerTriggerResponseScanStatus = "scanning"
)

type WebScannerTriggerResponseStatus string

const (
	WebScannerTriggerResponseStatusDisabled WebScannerTriggerResponseStatus = "Disabled"
	WebScannerTriggerResponseStatusEnabled  WebScannerTriggerResponseStatus = "Enabled"
)

type WebScannerNewParams struct {
	// Root domain to crawl (e.g. `example.com`). Required on create. Missing or empty
	// values fail request validation as HTTP 400. Present-but-malformed values are
	// rejected as HTTP 400 with the validation reason in `details`.
	RootDomain string            `json:"rootDomain" api:"required"`
	Name       param.Opt[string] `json:"name,omitzero"`
	// Maximum URLs to crawl per scan (1–20,000). Defaults to 100 when omitted.
	URLLimit param.Opt[float64] `json:"urlLimit,omitzero"`
	// URL glob patterns to skip during crawl. Max 100 entries.
	ExcludedPatterns []string `json:"excludedPatterns,omitzero"`
	// Additional seed URLs to include as crawl entry points. Each must be an http(s)
	// URL. Max 100 entries.
	IncludedURLs []string `json:"includedUrls,omitzero"`
	// Any of "Disabled", "Enabled".
	Status WebScannerNewParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r WebScannerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WebScannerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScannerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerNewParamsStatus string

const (
	WebScannerNewParamsStatusDisabled WebScannerNewParamsStatus = "Disabled"
	WebScannerNewParamsStatusEnabled  WebScannerNewParamsStatus = "Enabled"
)

type WebScannerUpdateParams struct {
	Name param.Opt[string] `json:"name,omitzero"`
	// Replace the scanner root domain. When provided, malformed values are rejected as
	// HTTP 400 with the validation reason in `details`.
	RootDomain       param.Opt[string]  `json:"rootDomain,omitzero"`
	URLLimit         param.Opt[float64] `json:"urlLimit,omitzero"`
	ExcludedPatterns []string           `json:"excludedPatterns,omitzero"`
	IncludedURLs     []string           `json:"includedUrls,omitzero"`
	// Any of "Disabled", "Enabled".
	Status WebScannerUpdateParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r WebScannerUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WebScannerUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScannerUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerUpdateParamsStatus string

const (
	WebScannerUpdateParamsStatusDisabled WebScannerUpdateParamsStatus = "Disabled"
	WebScannerUpdateParamsStatusEnabled  WebScannerUpdateParamsStatus = "Enabled"
)
