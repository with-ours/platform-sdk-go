// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/with-ours/platform-sdk-go/internal/apijson"
	"github.com/with-ours/platform-sdk-go/internal/apiquery"
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
// rate-limited: a 409 is returned if another scan is already in flight, or if this
// production monitor has `urlLimit >= 5000` and its previous scan completed within
// the last 10 minutes; the reason is in the response `error` field. Requires
// scope: webScanner:trigger
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

// Start a normal full scan with short-lived credentials supplied by your
// scheduler. Every in-scope page in this scan uses the credentials. Set
// `scanSchedule` to `manual` when your scheduler should be the only source of
// scans. Credential values are not returned or included in scan results. Requires
// scope: webScanner:trigger
func (r *WebScannerService) AuthenticatedScan(ctx context.Context, id string, body WebScannerAuthenticatedScanParams, opts ...option.RequestOption) (res *WebScannerAuthenticatedScanResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/web-scanners/%s/authenticated-scan", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Queue an isolated verification capture for one exact in-scope page. This does
// not update monitor history, inventory counts, or the monitor-wide last-scanned
// timestamp. Poll the verification-run endpoint with the returned id for terminal
// evidence. Requires scope: webScanner:trigger
func (r *WebScannerService) TargetedScan(ctx context.Context, id string, body WebScannerTargetedScanParams, opts ...option.RequestOption) (res *WebScannerTargetedScanResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/web-scanners/%s/targeted-scan", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Read a UUID-addressed one-page verification run. The run is isolated from normal
// monitor history and is returned only when it belongs to the requested scanner.
// Requires scope: webScanner:find
func (r *WebScannerService) VerificationRun(ctx context.Context, id string, query WebScannerVerificationRunParams, opts ...option.RequestOption) (res *WebScannerVerificationRunResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/web-scanners/%s/verification-run", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List the one-page verification runs recorded for this scanner, newest first,
// each with its own capture counts. This is how a fix is compared against the
// attempts before it without reading each run individually. Verification runs are
// isolated from monitor history and never affect inventory counts or the
// monitor-wide last-scanned timestamp. Requires scope: webScanner:find
func (r *WebScannerService) VerificationRuns(ctx context.Context, id string, query WebScannerVerificationRunsParams, opts ...option.RequestOption) (res *WebScannerVerificationRunsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/web-scanners/%s/verification-runs", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List the third-party trackers (requests) found on a scan run, with their risk,
// category, the pages they were seen on, and whether each host is already covered
// by a CMP consent service. Defaults to the latest run; pass `date` (an ISO-8601
// timestamp; only the calendar day is used to select the run) to read an earlier
// run. Documented exception to the cursor-pagination standard: paginates with
// `limit` and `offset` because each run is an immutable snapshot. A host that is
// neither covered (`coveredByCmp: false`) nor matched by a suppression rule still
// needs a triage decision — resolve it by adding the host to a CMP consent service
// or by creating a suppression rule with `POST /rest/v1/web-scanner-rules`. Use
// `GET /rest/v1/web-scanners/{id}/summary` for the rolled-up counts. Requires
// scope: webScanner:find
func (r *WebScannerService) Findings(ctx context.Context, id string, query WebScannerFindingsParams, opts ...option.RequestOption) (res *WebScannerFindingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/web-scanners/%s/findings", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List the cookies and local-storage entries observed on a scan run. Defaults to
// the latest run; pass `date` (an ISO-8601 timestamp; only the calendar day is
// used to select the run) to read an earlier run. Cookies paginate with `limit`
// and `offset` (a documented exception to the cursor-pagination standard, since
// each run is an immutable snapshot); local-storage entries are returned in full.
// Requires scope: webScanner:find
func (r *WebScannerService) Cookies(ctx context.Context, id string, query WebScannerCookiesParams, opts ...option.RequestOption) (res *WebScannerCookiesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/web-scanners/%s/cookies", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Compliance summary for a scan run — the rolled-up "what does this site look
// like, and what still needs a decision" view, assembled server-side so you do not
// have to page every finding. Includes total host/vendor/cookie counts, captured
// privacy policies and host coverage, a breakdown by risk and by category,
// coverage (how many hosts are already covered by a CMP consent service or a
// suppression rule vs. how many still need a decision), the new/removed host delta
// versus the previous run, and up to 10 highest-risk hosts that still need a
// decision. A null privacyPolicyUrl means that no policy was captured for the
// hostname. Defaults to the latest run; pass `date` (an ISO-8601 timestamp; only
// the calendar day is used to select the run) to read an earlier run. Clear a host
// that needs a decision by adding it to a CMP consent service or creating a
// suppression rule with `POST /rest/v1/web-scanner-rules`. When the scanner has no
// completed runs, every count is 0 and `runDate` is null. Requires scope:
// webScanner:find
func (r *WebScannerService) Summary(ctx context.Context, id string, query WebScannerSummaryParams, opts ...option.RequestOption) (res *WebScannerSummaryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/web-scanners/%s/summary", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
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
	// Any of "daily", "manual", "monthly", "weekly".
	ScanSchedule string `json:"scanSchedule" api:"required"`
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
	NextScheduledScanAt         string   `json:"nextScheduledScanAt" api:"nullable"`
	UpdatedAt                   string   `json:"updatedAt" api:"nullable"`
	URLLimit                    float64  `json:"urlLimit" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                          respjson.Field
		AccountID                   respjson.Field
		RootDomain                  respjson.Field
		ScanSchedule                respjson.Field
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
		NextScheduledScanAt         respjson.Field
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
	// Any of "daily", "manual", "monthly", "weekly".
	ScanSchedule WebScannerNewResponseScanSchedule `json:"scanSchedule" api:"required"`
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
	NextScheduledScanAt         string                      `json:"nextScheduledScanAt" api:"nullable"`
	UpdatedAt                   string                      `json:"updatedAt" api:"nullable"`
	URLLimit                    float64                     `json:"urlLimit" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                          respjson.Field
		AccountID                   respjson.Field
		RootDomain                  respjson.Field
		ScanSchedule                respjson.Field
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
		NextScheduledScanAt         respjson.Field
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

type WebScannerNewResponseScanSchedule string

const (
	WebScannerNewResponseScanScheduleDaily   WebScannerNewResponseScanSchedule = "daily"
	WebScannerNewResponseScanScheduleManual  WebScannerNewResponseScanSchedule = "manual"
	WebScannerNewResponseScanScheduleMonthly WebScannerNewResponseScanSchedule = "monthly"
	WebScannerNewResponseScanScheduleWeekly  WebScannerNewResponseScanSchedule = "weekly"
)

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
	// Any of "daily", "manual", "monthly", "weekly".
	ScanSchedule WebScannerGetResponseScanSchedule `json:"scanSchedule" api:"required"`
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
	NextScheduledScanAt         string                      `json:"nextScheduledScanAt" api:"nullable"`
	UpdatedAt                   string                      `json:"updatedAt" api:"nullable"`
	URLLimit                    float64                     `json:"urlLimit" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                          respjson.Field
		AccountID                   respjson.Field
		RootDomain                  respjson.Field
		ScanSchedule                respjson.Field
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
		NextScheduledScanAt         respjson.Field
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

type WebScannerGetResponseScanSchedule string

const (
	WebScannerGetResponseScanScheduleDaily   WebScannerGetResponseScanSchedule = "daily"
	WebScannerGetResponseScanScheduleManual  WebScannerGetResponseScanSchedule = "manual"
	WebScannerGetResponseScanScheduleMonthly WebScannerGetResponseScanSchedule = "monthly"
	WebScannerGetResponseScanScheduleWeekly  WebScannerGetResponseScanSchedule = "weekly"
)

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
	// Any of "daily", "manual", "monthly", "weekly".
	ScanSchedule WebScannerUpdateResponseScanSchedule `json:"scanSchedule" api:"required"`
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
	NextScheduledScanAt         string                         `json:"nextScheduledScanAt" api:"nullable"`
	UpdatedAt                   string                         `json:"updatedAt" api:"nullable"`
	URLLimit                    float64                        `json:"urlLimit" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                          respjson.Field
		AccountID                   respjson.Field
		RootDomain                  respjson.Field
		ScanSchedule                respjson.Field
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
		NextScheduledScanAt         respjson.Field
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

type WebScannerUpdateResponseScanSchedule string

const (
	WebScannerUpdateResponseScanScheduleDaily   WebScannerUpdateResponseScanSchedule = "daily"
	WebScannerUpdateResponseScanScheduleManual  WebScannerUpdateResponseScanSchedule = "manual"
	WebScannerUpdateResponseScanScheduleMonthly WebScannerUpdateResponseScanSchedule = "monthly"
	WebScannerUpdateResponseScanScheduleWeekly  WebScannerUpdateResponseScanSchedule = "weekly"
)

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
	// Any of "daily", "manual", "monthly", "weekly".
	ScanSchedule WebScannerTriggerResponseScanSchedule `json:"scanSchedule" api:"required"`
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
	NextScheduledScanAt         string                          `json:"nextScheduledScanAt" api:"nullable"`
	UpdatedAt                   string                          `json:"updatedAt" api:"nullable"`
	URLLimit                    float64                         `json:"urlLimit" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                          respjson.Field
		AccountID                   respjson.Field
		RootDomain                  respjson.Field
		ScanSchedule                respjson.Field
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
		NextScheduledScanAt         respjson.Field
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

type WebScannerTriggerResponseScanSchedule string

const (
	WebScannerTriggerResponseScanScheduleDaily   WebScannerTriggerResponseScanSchedule = "daily"
	WebScannerTriggerResponseScanScheduleManual  WebScannerTriggerResponseScanSchedule = "manual"
	WebScannerTriggerResponseScanScheduleMonthly WebScannerTriggerResponseScanSchedule = "monthly"
	WebScannerTriggerResponseScanScheduleWeekly  WebScannerTriggerResponseScanSchedule = "weekly"
)

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

type WebScannerAuthenticatedScanResponse struct {
	ID         string `json:"id" api:"required"`
	AccountID  string `json:"accountId" api:"required"`
	RootDomain string `json:"rootDomain" api:"required"`
	// Any of "daily", "manual", "monthly", "weekly".
	ScanSchedule WebScannerAuthenticatedScanResponseScanSchedule `json:"scanSchedule" api:"required"`
	// Any of "idle", "scanning".
	ScanStatus WebScannerAuthenticatedScanResponseScanStatus `json:"scanStatus" api:"required"`
	// Any of "Disabled", "Enabled".
	Status                      WebScannerAuthenticatedScanResponseStatus `json:"status" api:"required"`
	CreatedAt                   string                                    `json:"createdAt" api:"nullable"`
	ExcludedPatterns            []string                                  `json:"excludedPatterns" api:"nullable"`
	IncludedURLs                []string                                  `json:"includedUrls" api:"nullable"`
	LastRunCookieCount          float64                                   `json:"lastRunCookieCount" api:"nullable"`
	LastRunHighRiskRequestCount float64                                   `json:"lastRunHighRiskRequestCount" api:"nullable"`
	LastRunRequestCount         float64                                   `json:"lastRunRequestCount" api:"nullable"`
	LastRunSuccessURLCount      float64                                   `json:"lastRunSuccessUrlCount" api:"nullable"`
	LastScannedAt               string                                    `json:"lastScannedAt" api:"nullable"`
	LastScanStartedAt           string                                    `json:"lastScanStartedAt" api:"nullable"`
	Name                        string                                    `json:"name" api:"nullable"`
	NextScheduledScanAt         string                                    `json:"nextScheduledScanAt" api:"nullable"`
	UpdatedAt                   string                                    `json:"updatedAt" api:"nullable"`
	URLLimit                    float64                                   `json:"urlLimit" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                          respjson.Field
		AccountID                   respjson.Field
		RootDomain                  respjson.Field
		ScanSchedule                respjson.Field
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
		NextScheduledScanAt         respjson.Field
		UpdatedAt                   respjson.Field
		URLLimit                    respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerAuthenticatedScanResponse) RawJSON() string { return r.JSON.raw }
func (r *WebScannerAuthenticatedScanResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerAuthenticatedScanResponseScanSchedule string

const (
	WebScannerAuthenticatedScanResponseScanScheduleDaily   WebScannerAuthenticatedScanResponseScanSchedule = "daily"
	WebScannerAuthenticatedScanResponseScanScheduleManual  WebScannerAuthenticatedScanResponseScanSchedule = "manual"
	WebScannerAuthenticatedScanResponseScanScheduleMonthly WebScannerAuthenticatedScanResponseScanSchedule = "monthly"
	WebScannerAuthenticatedScanResponseScanScheduleWeekly  WebScannerAuthenticatedScanResponseScanSchedule = "weekly"
)

type WebScannerAuthenticatedScanResponseScanStatus string

const (
	WebScannerAuthenticatedScanResponseScanStatusIdle     WebScannerAuthenticatedScanResponseScanStatus = "idle"
	WebScannerAuthenticatedScanResponseScanStatusScanning WebScannerAuthenticatedScanResponseScanStatus = "scanning"
)

type WebScannerAuthenticatedScanResponseStatus string

const (
	WebScannerAuthenticatedScanResponseStatusDisabled WebScannerAuthenticatedScanResponseStatus = "Disabled"
	WebScannerAuthenticatedScanResponseStatusEnabled  WebScannerAuthenticatedScanResponseStatus = "Enabled"
)

type WebScannerTargetedScanResponse struct {
	ID           string `json:"id" api:"required" format:"uuid"`
	CreatedAt    string `json:"createdAt" api:"required"`
	RequestedURL string `json:"requestedUrl" api:"required" format:"uri"`
	ScannerID    string `json:"scannerId" api:"required" format:"uuid"`
	// Any of "queued", "in_progress", "completed", "failed".
	Status      WebScannerTargetedScanResponseStatus `json:"status" api:"required"`
	Cause       string                               `json:"cause" api:"nullable"`
	CompletedAt string                               `json:"completedAt" api:"nullable"`
	ResolvedURL string                               `json:"resolvedUrl" api:"nullable"`
	Run         WebScannerTargetedScanResponseRun    `json:"run" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		RequestedURL respjson.Field
		ScannerID    respjson.Field
		Status       respjson.Field
		Cause        respjson.Field
		CompletedAt  respjson.Field
		ResolvedURL  respjson.Field
		Run          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerTargetedScanResponse) RawJSON() string { return r.JSON.raw }
func (r *WebScannerTargetedScanResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerTargetedScanResponseStatus string

const (
	WebScannerTargetedScanResponseStatusQueued     WebScannerTargetedScanResponseStatus = "queued"
	WebScannerTargetedScanResponseStatusInProgress WebScannerTargetedScanResponseStatus = "in_progress"
	WebScannerTargetedScanResponseStatusCompleted  WebScannerTargetedScanResponseStatus = "completed"
	WebScannerTargetedScanResponseStatusFailed     WebScannerTargetedScanResponseStatus = "failed"
)

type WebScannerTargetedScanResponseRun struct {
	CookieCount        int64                                     `json:"cookieCount" api:"required"`
	LocalStorageCount  int64                                     `json:"localStorageCount" api:"required"`
	RequestCount       int64                                     `json:"requestCount" api:"required"`
	VendorCount        int64                                     `json:"vendorCount" api:"required"`
	Metadata           WebScannerTargetedScanResponseRunMetadata `json:"metadata" api:"nullable"`
	SummaryGeneratedAt string                                    `json:"summary_generated_at" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CookieCount        respjson.Field
		LocalStorageCount  respjson.Field
		RequestCount       respjson.Field
		VendorCount        respjson.Field
		Metadata           respjson.Field
		SummaryGeneratedAt respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerTargetedScanResponseRun) RawJSON() string { return r.JSON.raw }
func (r *WebScannerTargetedScanResponseRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerTargetedScanResponseRunMetadata struct {
	FailedURLs       []WebScannerTargetedScanResponseRunMetadataFailedURL `json:"failed_urls" api:"required"`
	SuccessURLs      []string                                             `json:"success_urls" api:"required"`
	ExcludedPatterns []string                                             `json:"excluded_patterns" api:"nullable"`
	ExcludedURLs     []string                                             `json:"excluded_urls" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FailedURLs       respjson.Field
		SuccessURLs      respjson.Field
		ExcludedPatterns respjson.Field
		ExcludedURLs     respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerTargetedScanResponseRunMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebScannerTargetedScanResponseRunMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerTargetedScanResponseRunMetadataFailedURL struct {
	URL        string  `json:"url" api:"required"`
	Details    string  `json:"details" api:"nullable"`
	StatusCode float64 `json:"status_code" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Details     respjson.Field
		StatusCode  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerTargetedScanResponseRunMetadataFailedURL) RawJSON() string { return r.JSON.raw }
func (r *WebScannerTargetedScanResponseRunMetadataFailedURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerVerificationRunResponse struct {
	ID           string `json:"id" api:"required" format:"uuid"`
	CreatedAt    string `json:"createdAt" api:"required"`
	RequestedURL string `json:"requestedUrl" api:"required" format:"uri"`
	ScannerID    string `json:"scannerId" api:"required" format:"uuid"`
	// Any of "queued", "in_progress", "completed", "failed".
	Status      WebScannerVerificationRunResponseStatus `json:"status" api:"required"`
	Cause       string                                  `json:"cause" api:"nullable"`
	CompletedAt string                                  `json:"completedAt" api:"nullable"`
	ResolvedURL string                                  `json:"resolvedUrl" api:"nullable"`
	Run         WebScannerVerificationRunResponseRun    `json:"run" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		RequestedURL respjson.Field
		ScannerID    respjson.Field
		Status       respjson.Field
		Cause        respjson.Field
		CompletedAt  respjson.Field
		ResolvedURL  respjson.Field
		Run          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerVerificationRunResponse) RawJSON() string { return r.JSON.raw }
func (r *WebScannerVerificationRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerVerificationRunResponseStatus string

const (
	WebScannerVerificationRunResponseStatusQueued     WebScannerVerificationRunResponseStatus = "queued"
	WebScannerVerificationRunResponseStatusInProgress WebScannerVerificationRunResponseStatus = "in_progress"
	WebScannerVerificationRunResponseStatusCompleted  WebScannerVerificationRunResponseStatus = "completed"
	WebScannerVerificationRunResponseStatusFailed     WebScannerVerificationRunResponseStatus = "failed"
)

type WebScannerVerificationRunResponseRun struct {
	CookieCount        int64                                        `json:"cookieCount" api:"required"`
	LocalStorageCount  int64                                        `json:"localStorageCount" api:"required"`
	RequestCount       int64                                        `json:"requestCount" api:"required"`
	VendorCount        int64                                        `json:"vendorCount" api:"required"`
	Metadata           WebScannerVerificationRunResponseRunMetadata `json:"metadata" api:"nullable"`
	SummaryGeneratedAt string                                       `json:"summary_generated_at" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CookieCount        respjson.Field
		LocalStorageCount  respjson.Field
		RequestCount       respjson.Field
		VendorCount        respjson.Field
		Metadata           respjson.Field
		SummaryGeneratedAt respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerVerificationRunResponseRun) RawJSON() string { return r.JSON.raw }
func (r *WebScannerVerificationRunResponseRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerVerificationRunResponseRunMetadata struct {
	FailedURLs       []WebScannerVerificationRunResponseRunMetadataFailedURL `json:"failed_urls" api:"required"`
	SuccessURLs      []string                                                `json:"success_urls" api:"required"`
	ExcludedPatterns []string                                                `json:"excluded_patterns" api:"nullable"`
	ExcludedURLs     []string                                                `json:"excluded_urls" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FailedURLs       respjson.Field
		SuccessURLs      respjson.Field
		ExcludedPatterns respjson.Field
		ExcludedURLs     respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerVerificationRunResponseRunMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebScannerVerificationRunResponseRunMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerVerificationRunResponseRunMetadataFailedURL struct {
	URL        string  `json:"url" api:"required"`
	Details    string  `json:"details" api:"nullable"`
	StatusCode float64 `json:"status_code" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Details     respjson.Field
		StatusCode  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerVerificationRunResponseRunMetadataFailedURL) RawJSON() string { return r.JSON.raw }
func (r *WebScannerVerificationRunResponseRunMetadataFailedURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerVerificationRunsResponse struct {
	Items []WebScannerVerificationRunsResponseItem `json:"items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerVerificationRunsResponse) RawJSON() string { return r.JSON.raw }
func (r *WebScannerVerificationRunsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerVerificationRunsResponseItem struct {
	ID           string `json:"id" api:"required" format:"uuid"`
	CreatedAt    string `json:"createdAt" api:"required"`
	RequestedURL string `json:"requestedUrl" api:"required" format:"uri"`
	ScannerID    string `json:"scannerId" api:"required" format:"uuid"`
	// Any of "queued", "in_progress", "completed", "failed".
	Status      string                                    `json:"status" api:"required"`
	Cause       string                                    `json:"cause" api:"nullable"`
	CompletedAt string                                    `json:"completedAt" api:"nullable"`
	ResolvedURL string                                    `json:"resolvedUrl" api:"nullable"`
	Run         WebScannerVerificationRunsResponseItemRun `json:"run" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		RequestedURL respjson.Field
		ScannerID    respjson.Field
		Status       respjson.Field
		Cause        respjson.Field
		CompletedAt  respjson.Field
		ResolvedURL  respjson.Field
		Run          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerVerificationRunsResponseItem) RawJSON() string { return r.JSON.raw }
func (r *WebScannerVerificationRunsResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerVerificationRunsResponseItemRun struct {
	CookieCount        int64                                             `json:"cookieCount" api:"required"`
	LocalStorageCount  int64                                             `json:"localStorageCount" api:"required"`
	RequestCount       int64                                             `json:"requestCount" api:"required"`
	VendorCount        int64                                             `json:"vendorCount" api:"required"`
	Metadata           WebScannerVerificationRunsResponseItemRunMetadata `json:"metadata" api:"nullable"`
	SummaryGeneratedAt string                                            `json:"summary_generated_at" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CookieCount        respjson.Field
		LocalStorageCount  respjson.Field
		RequestCount       respjson.Field
		VendorCount        respjson.Field
		Metadata           respjson.Field
		SummaryGeneratedAt respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerVerificationRunsResponseItemRun) RawJSON() string { return r.JSON.raw }
func (r *WebScannerVerificationRunsResponseItemRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerVerificationRunsResponseItemRunMetadata struct {
	FailedURLs       []WebScannerVerificationRunsResponseItemRunMetadataFailedURL `json:"failed_urls" api:"required"`
	SuccessURLs      []string                                                     `json:"success_urls" api:"required"`
	ExcludedPatterns []string                                                     `json:"excluded_patterns" api:"nullable"`
	ExcludedURLs     []string                                                     `json:"excluded_urls" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FailedURLs       respjson.Field
		SuccessURLs      respjson.Field
		ExcludedPatterns respjson.Field
		ExcludedURLs     respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerVerificationRunsResponseItemRunMetadata) RawJSON() string { return r.JSON.raw }
func (r *WebScannerVerificationRunsResponseItemRunMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerVerificationRunsResponseItemRunMetadataFailedURL struct {
	URL        string  `json:"url" api:"required"`
	Details    string  `json:"details" api:"nullable"`
	StatusCode float64 `json:"status_code" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Details     respjson.Field
		StatusCode  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerVerificationRunsResponseItemRunMetadataFailedURL) RawJSON() string {
	return r.JSON.raw
}
func (r *WebScannerVerificationRunsResponseItemRunMetadataFailedURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerFindingsResponse struct {
	// True when more findings are available beyond the current window.
	HasMore bool `json:"hasMore" api:"required"`
	// Third-party trackers seen on the run. `coveredByCmp` is true when the host is
	// already mapped to a CMP consent service; `risk` is high/medium/low/unknown.
	// Hosts that are neither covered nor matched by a suppression rule are the ones
	// that need a triage decision — clear them by adding the host to a CMP consent
	// service or creating a suppression rule (POST /rest/v1/web-scanner-rules).
	Items []WebScannerFindingsResponseItem `json:"items" api:"required"`
	// Total number of findings in the run.
	Total int64 `json:"total" api:"required"`
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
func (r WebScannerFindingsResponse) RawJSON() string { return r.JSON.raw }
func (r *WebScannerFindingsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerFindingsResponseItem struct {
	CoveredByCmp bool     `json:"coveredByCmp" api:"required"`
	Hostname     string   `json:"hostname" api:"required"`
	SeenOn       []string `json:"seenOn" api:"required"`
	// Any of "audio", "beacon", "document", "eventsource", "fedcm", "fetch", "font",
	// "image", "manifest", "media", "other", "ping", "prefetch", "script",
	// "stylesheet", "texttrack", "video", "websocket", "xhr".
	Types                []string                               `json:"types" api:"required"`
	URLs                 []string                               `json:"urls" api:"required"`
	Category             string                                 `json:"category" api:"nullable"`
	Cookies              []WebScannerFindingsResponseItemCookie `json:"cookies" api:"nullable"`
	CoveredByVendorLabel string                                 `json:"coveredByVendorLabel" api:"nullable"`
	DisplayName          string                                 `json:"displayName" api:"nullable"`
	PrivacyKeywords      []string                               `json:"privacyKeywords" api:"nullable"`
	Risk                 string                                 `json:"risk" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CoveredByCmp         respjson.Field
		Hostname             respjson.Field
		SeenOn               respjson.Field
		Types                respjson.Field
		URLs                 respjson.Field
		Category             respjson.Field
		Cookies              respjson.Field
		CoveredByVendorLabel respjson.Field
		DisplayName          respjson.Field
		PrivacyKeywords      respjson.Field
		Risk                 respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerFindingsResponseItem) RawJSON() string { return r.JSON.raw }
func (r *WebScannerFindingsResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerFindingsResponseItemCookie struct {
	Name   string `json:"name" api:"required"`
	Domain string `json:"domain" api:"nullable"`
	Path   string `json:"path" api:"nullable"`
	Value  string `json:"value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Domain      respjson.Field
		Path        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerFindingsResponseItemCookie) RawJSON() string { return r.JSON.raw }
func (r *WebScannerFindingsResponseItemCookie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerCookiesResponse struct {
	// Cookies observed on the run, paginated by `limit`/`offset`.
	Cookies []WebScannerCookiesResponseCookie `json:"cookies" api:"required"`
	// True when more cookies are available beyond the current window.
	HasMore bool `json:"hasMore" api:"required"`
	// Local-storage entries observed on the run. Returned in full (not paginated).
	LocalStorage      []WebScannerCookiesResponseLocalStorage `json:"localStorage" api:"required"`
	TotalCookies      int64                                   `json:"totalCookies" api:"required"`
	TotalLocalStorage int64                                   `json:"totalLocalStorage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cookies           respjson.Field
		HasMore           respjson.Field
		LocalStorage      respjson.Field
		TotalCookies      respjson.Field
		TotalLocalStorage respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerCookiesResponse) RawJSON() string { return r.JSON.raw }
func (r *WebScannerCookiesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerCookiesResponseCookie struct {
	Name   string `json:"name" api:"required"`
	Domain string `json:"domain" api:"nullable"`
	Path   string `json:"path" api:"nullable"`
	Value  string `json:"value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Domain      respjson.Field
		Path        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerCookiesResponseCookie) RawJSON() string { return r.JSON.raw }
func (r *WebScannerCookiesResponseCookie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerCookiesResponseLocalStorage struct {
	Name  string `json:"name" api:"required"`
	Value string `json:"value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerCookiesResponseLocalStorage) RawJSON() string { return r.JSON.raw }
func (r *WebScannerCookiesResponseLocalStorage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerSummaryResponse struct {
	ByCategory        []WebScannerSummaryResponseByCategory `json:"byCategory" api:"required"`
	CookieCount       int64                                 `json:"cookieCount" api:"required"`
	CountsByRisk      WebScannerSummaryResponseCountsByRisk `json:"countsByRisk" api:"required"`
	Coverage          WebScannerSummaryResponseCoverage     `json:"coverage" api:"required"`
	HostCount         int64                                 `json:"hostCount" api:"required"`
	LocalStorageCount int64                                 `json:"localStorageCount" api:"required"`
	// Distinct privacy policies captured for first-party hostnames in the selected
	// scan run, including the bounded visible policy text.
	PrivacyPolicies    []WebScannerSummaryResponsePrivacyPolicy `json:"privacyPolicies" api:"required"`
	PrivacyPolicyCount int64                                    `json:"privacyPolicyCount" api:"required"`
	// Every successfully crawled first-party hostname and its captured privacy policy
	// URL. A null privacyPolicyUrl means no policy was captured for that hostname.
	PrivacyPolicyHosts []WebScannerSummaryResponsePrivacyPolicyHost `json:"privacyPolicyHosts" api:"required"`
	RootDomain         string                                       `json:"rootDomain" api:"required"`
	ScannerID          string                                       `json:"scannerId" api:"required"`
	// Any of "idle", "scanning".
	ScanStatus WebScannerSummaryResponseScanStatus `json:"scanStatus" api:"required"`
	// Up to 10 hosts that still need a decision (neither CMP-covered nor suppressed),
	// highest risk first. Clear each by adding the host to a CMP consent service or
	// creating a suppression rule (POST /rest/v1/web-scanner-rules) with the reason
	// that explains why it is allowed (baa, internal, approved, compliant, firstParty,
	// ignore).
	TopUncoveredHosts []WebScannerSummaryResponseTopUncoveredHost `json:"topUncoveredHosts" api:"required"`
	VendorCount       int64                                       `json:"vendorCount" api:"required"`
	// Automated accessibility (WCAG 2.1/2.2 A + AA) rollup for the run: `score` is a
	// 0-100 site score (mean of per-page scores; higher is better), with distinct
	// rule-violation `countsByImpact` and the most frequently violated rules in
	// `topViolations` (each with the number of pages it appears on). Covers only the
	// machine-detectable subset of WCAG (~30-40%) — a high score is not a
	// certification of full conformance; manual audit is still required. Null when the
	// run audited no pages.
	Accessibility WebScannerSummaryResponseAccessibility `json:"accessibility" api:"nullable"`
	Delta         WebScannerSummaryResponseDelta         `json:"delta" api:"nullable"`
	RunDate       string                                 `json:"runDate" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByCategory         respjson.Field
		CookieCount        respjson.Field
		CountsByRisk       respjson.Field
		Coverage           respjson.Field
		HostCount          respjson.Field
		LocalStorageCount  respjson.Field
		PrivacyPolicies    respjson.Field
		PrivacyPolicyCount respjson.Field
		PrivacyPolicyHosts respjson.Field
		RootDomain         respjson.Field
		ScannerID          respjson.Field
		ScanStatus         respjson.Field
		TopUncoveredHosts  respjson.Field
		VendorCount        respjson.Field
		Accessibility      respjson.Field
		Delta              respjson.Field
		RunDate            respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerSummaryResponse) RawJSON() string { return r.JSON.raw }
func (r *WebScannerSummaryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerSummaryResponseByCategory struct {
	Category  string `json:"category" api:"required"`
	HostCount int64  `json:"hostCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Category    respjson.Field
		HostCount   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerSummaryResponseByCategory) RawJSON() string { return r.JSON.raw }
func (r *WebScannerSummaryResponseByCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerSummaryResponseCountsByRisk struct {
	High    int64 `json:"high" api:"required"`
	Low     int64 `json:"low" api:"required"`
	Medium  int64 `json:"medium" api:"required"`
	Unknown int64 `json:"unknown" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		High        respjson.Field
		Low         respjson.Field
		Medium      respjson.Field
		Unknown     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerSummaryResponseCountsByRisk) RawJSON() string { return r.JSON.raw }
func (r *WebScannerSummaryResponseCountsByRisk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerSummaryResponseCoverage struct {
	CoveragePercent        int64 `json:"coveragePercent" api:"required"`
	CoveredHostCount       int64 `json:"coveredHostCount" api:"required"`
	NeedsDecisionHostCount int64 `json:"needsDecisionHostCount" api:"required"`
	TotalHostCount         int64 `json:"totalHostCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CoveragePercent        respjson.Field
		CoveredHostCount       respjson.Field
		NeedsDecisionHostCount respjson.Field
		TotalHostCount         respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerSummaryResponseCoverage) RawJSON() string { return r.JSON.raw }
func (r *WebScannerSummaryResponseCoverage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerSummaryResponsePrivacyPolicy struct {
	Hostnames []string `json:"hostnames" api:"required"`
	Text      string   `json:"text" api:"required"`
	URL       string   `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hostnames   respjson.Field
		Text        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerSummaryResponsePrivacyPolicy) RawJSON() string { return r.JSON.raw }
func (r *WebScannerSummaryResponsePrivacyPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerSummaryResponsePrivacyPolicyHost struct {
	Hostname         string `json:"hostname" api:"required"`
	PrivacyPolicyURL string `json:"privacyPolicyUrl" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hostname         respjson.Field
		PrivacyPolicyURL respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerSummaryResponsePrivacyPolicyHost) RawJSON() string { return r.JSON.raw }
func (r *WebScannerSummaryResponsePrivacyPolicyHost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerSummaryResponseScanStatus string

const (
	WebScannerSummaryResponseScanStatusIdle     WebScannerSummaryResponseScanStatus = "idle"
	WebScannerSummaryResponseScanStatusScanning WebScannerSummaryResponseScanStatus = "scanning"
)

type WebScannerSummaryResponseTopUncoveredHost struct {
	CoveredByCmp bool     `json:"coveredByCmp" api:"required"`
	Hostname     string   `json:"hostname" api:"required"`
	SeenOn       []string `json:"seenOn" api:"required"`
	// Any of "audio", "beacon", "document", "eventsource", "fedcm", "fetch", "font",
	// "image", "manifest", "media", "other", "ping", "prefetch", "script",
	// "stylesheet", "texttrack", "video", "websocket", "xhr".
	Types                []string                                          `json:"types" api:"required"`
	URLs                 []string                                          `json:"urls" api:"required"`
	Category             string                                            `json:"category" api:"nullable"`
	Cookies              []WebScannerSummaryResponseTopUncoveredHostCookie `json:"cookies" api:"nullable"`
	CoveredByVendorLabel string                                            `json:"coveredByVendorLabel" api:"nullable"`
	DisplayName          string                                            `json:"displayName" api:"nullable"`
	PrivacyKeywords      []string                                          `json:"privacyKeywords" api:"nullable"`
	Risk                 string                                            `json:"risk" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CoveredByCmp         respjson.Field
		Hostname             respjson.Field
		SeenOn               respjson.Field
		Types                respjson.Field
		URLs                 respjson.Field
		Category             respjson.Field
		Cookies              respjson.Field
		CoveredByVendorLabel respjson.Field
		DisplayName          respjson.Field
		PrivacyKeywords      respjson.Field
		Risk                 respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerSummaryResponseTopUncoveredHost) RawJSON() string { return r.JSON.raw }
func (r *WebScannerSummaryResponseTopUncoveredHost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerSummaryResponseTopUncoveredHostCookie struct {
	Name   string `json:"name" api:"required"`
	Domain string `json:"domain" api:"nullable"`
	Path   string `json:"path" api:"nullable"`
	Value  string `json:"value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Domain      respjson.Field
		Path        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerSummaryResponseTopUncoveredHostCookie) RawJSON() string { return r.JSON.raw }
func (r *WebScannerSummaryResponseTopUncoveredHostCookie) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Automated accessibility (WCAG 2.1/2.2 A + AA) rollup for the run: `score` is a
// 0-100 site score (mean of per-page scores; higher is better), with distinct
// rule-violation `countsByImpact` and the most frequently violated rules in
// `topViolations` (each with the number of pages it appears on). Covers only the
// machine-detectable subset of WCAG (~30-40%) — a high score is not a
// certification of full conformance; manual audit is still required. Null when the
// run audited no pages.
type WebScannerSummaryResponseAccessibility struct {
	CountsByImpact  WebScannerSummaryResponseAccessibilityCountsByImpact `json:"countsByImpact" api:"required"`
	Engine          string                                               `json:"engine" api:"required"`
	PagesEvaluated  int64                                                `json:"pagesEvaluated" api:"required"`
	Score           int64                                                `json:"score" api:"required"`
	TopViolations   []WebScannerSummaryResponseAccessibilityTopViolation `json:"topViolations" api:"required"`
	TotalNodes      int64                                                `json:"totalNodes" api:"required"`
	TotalViolations int64                                                `json:"totalViolations" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CountsByImpact  respjson.Field
		Engine          respjson.Field
		PagesEvaluated  respjson.Field
		Score           respjson.Field
		TopViolations   respjson.Field
		TotalNodes      respjson.Field
		TotalViolations respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerSummaryResponseAccessibility) RawJSON() string { return r.JSON.raw }
func (r *WebScannerSummaryResponseAccessibility) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerSummaryResponseAccessibilityCountsByImpact struct {
	Critical int64 `json:"critical" api:"required"`
	Minor    int64 `json:"minor" api:"required"`
	Moderate int64 `json:"moderate" api:"required"`
	Serious  int64 `json:"serious" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Critical    respjson.Field
		Minor       respjson.Field
		Moderate    respjson.Field
		Serious     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerSummaryResponseAccessibilityCountsByImpact) RawJSON() string { return r.JSON.raw }
func (r *WebScannerSummaryResponseAccessibilityCountsByImpact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerSummaryResponseAccessibilityTopViolation struct {
	ID          string                                                         `json:"id" api:"required"`
	Help        string                                                         `json:"help" api:"required"`
	HelpURL     string                                                         `json:"helpUrl" api:"required"`
	NodeCount   int64                                                          `json:"nodeCount" api:"required"`
	PageCount   int64                                                          `json:"pageCount" api:"required"`
	Pages       []string                                                       `json:"pages" api:"required"`
	SampleNodes []WebScannerSummaryResponseAccessibilityTopViolationSampleNode `json:"sampleNodes" api:"required"`
	WcagTags    []string                                                       `json:"wcagTags" api:"required"`
	Impact      string                                                         `json:"impact" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Help        respjson.Field
		HelpURL     respjson.Field
		NodeCount   respjson.Field
		PageCount   respjson.Field
		Pages       respjson.Field
		SampleNodes respjson.Field
		WcagTags    respjson.Field
		Impact      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerSummaryResponseAccessibilityTopViolation) RawJSON() string { return r.JSON.raw }
func (r *WebScannerSummaryResponseAccessibilityTopViolation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerSummaryResponseAccessibilityTopViolationSampleNode struct {
	HTML           string   `json:"html" api:"required"`
	Target         []string `json:"target" api:"required"`
	FailureSummary string   `json:"failureSummary" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HTML           respjson.Field
		Target         respjson.Field
		FailureSummary respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerSummaryResponseAccessibilityTopViolationSampleNode) RawJSON() string {
	return r.JSON.raw
}
func (r *WebScannerSummaryResponseAccessibilityTopViolationSampleNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerSummaryResponseDelta struct {
	NewHostCount     int64 `json:"newHostCount" api:"required"`
	RemovedHostCount int64 `json:"removedHostCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NewHostCount     respjson.Field
		RemovedHostCount respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerSummaryResponseDelta) RawJSON() string { return r.JSON.raw }
func (r *WebScannerSummaryResponseDelta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
	// How often the scanner crawls this monitor on its own: `daily`, `weekly`,
	// `monthly`, or `manual` to disable scheduled crawls and only run on demand.
	// Defaults to `weekly`. Cadences advance on UTC calendar days from the last
	// completed scan.
	//
	// Any of "daily", "manual", "monthly", "weekly".
	ScanSchedule WebScannerNewParamsScanSchedule `json:"scanSchedule,omitzero"`
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

// How often the scanner crawls this monitor on its own: `daily`, `weekly`,
// `monthly`, or `manual` to disable scheduled crawls and only run on demand.
// Defaults to `weekly`. Cadences advance on UTC calendar days from the last
// completed scan.
type WebScannerNewParamsScanSchedule string

const (
	WebScannerNewParamsScanScheduleDaily   WebScannerNewParamsScanSchedule = "daily"
	WebScannerNewParamsScanScheduleManual  WebScannerNewParamsScanSchedule = "manual"
	WebScannerNewParamsScanScheduleMonthly WebScannerNewParamsScanSchedule = "monthly"
	WebScannerNewParamsScanScheduleWeekly  WebScannerNewParamsScanSchedule = "weekly"
)

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
	// How often the scanner crawls this monitor on its own: `daily`, `weekly`,
	// `monthly`, or `manual` to disable scheduled crawls and only run on demand.
	// Defaults to `weekly`. Cadences advance on UTC calendar days from the last
	// completed scan. Omit to leave the current cadence unchanged.
	//
	// Any of "daily", "manual", "monthly", "weekly".
	ScanSchedule WebScannerUpdateParamsScanSchedule `json:"scanSchedule,omitzero"`
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

// How often the scanner crawls this monitor on its own: `daily`, `weekly`,
// `monthly`, or `manual` to disable scheduled crawls and only run on demand.
// Defaults to `weekly`. Cadences advance on UTC calendar days from the last
// completed scan. Omit to leave the current cadence unchanged.
type WebScannerUpdateParamsScanSchedule string

const (
	WebScannerUpdateParamsScanScheduleDaily   WebScannerUpdateParamsScanSchedule = "daily"
	WebScannerUpdateParamsScanScheduleManual  WebScannerUpdateParamsScanSchedule = "manual"
	WebScannerUpdateParamsScanScheduleMonthly WebScannerUpdateParamsScanSchedule = "monthly"
	WebScannerUpdateParamsScanScheduleWeekly  WebScannerUpdateParamsScanSchedule = "weekly"
)

type WebScannerUpdateParamsStatus string

const (
	WebScannerUpdateParamsStatusDisabled WebScannerUpdateParamsStatus = "Disabled"
	WebScannerUpdateParamsStatusEnabled  WebScannerUpdateParamsStatus = "Enabled"
)

type WebScannerAuthenticatedScanParams struct {
	// Short-lived credentials for this scan. Their values are not returned in API
	// responses or captured in scan results.
	Credentials []WebScannerAuthenticatedScanParamsCredential `json:"credentials,omitzero" api:"required"`
	// Whole-hour validity window for these credentials. Send a fresh credential for
	// every externally scheduled scan.
	ExpiresInHours int64 `json:"expiresInHours" api:"required"`
	paramObj
}

func (r WebScannerAuthenticatedScanParams) MarshalJSON() (data []byte, err error) {
	type shadow WebScannerAuthenticatedScanParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScannerAuthenticatedScanParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Location, Name, Value are required.
type WebScannerAuthenticatedScanParamsCredential struct {
	// Where to place the credential for first-party crawl requests. Header credentials
	// are not sent to third-party requests.
	//
	// Any of "header", "cookie".
	Location string `json:"location,omitzero" api:"required"`
	// HTTP header or cookie name. Use `Authorization` for a bearer-style header.
	Name string `json:"name" api:"required"`
	// Complete credential value, including any required prefix such as `Bearer `.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r WebScannerAuthenticatedScanParamsCredential) MarshalJSON() (data []byte, err error) {
	type shadow WebScannerAuthenticatedScanParamsCredential
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScannerAuthenticatedScanParamsCredential) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebScannerAuthenticatedScanParamsCredential](
		"location", "header", "cookie",
	)
}

type WebScannerTargetedScanParams struct {
	// Exact in-scope HTTP(S) page URL to verify. Fragments are removed; query
	// parameters are retained.
	TargetURL string `json:"targetUrl" api:"required" format:"uri"`
	paramObj
}

func (r WebScannerTargetedScanParams) MarshalJSON() (data []byte, err error) {
	type shadow WebScannerTargetedScanParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScannerTargetedScanParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerVerificationRunParams struct {
	// Verification run UUID returned by the targeted scan command.
	RunID string `query:"runId" api:"required" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [WebScannerVerificationRunParams]'s query parameters as
// `url.Values`.
func (r WebScannerVerificationRunParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebScannerVerificationRunsParams struct {
	// How many runs to read, newest first. Defaults to 20, capped at 50.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebScannerVerificationRunsParams]'s query parameters as
// `url.Values`.
func (r WebScannerVerificationRunsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebScannerFindingsParams struct {
	// Skip this many findings before returning. Use with `limit` for load-more paging.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Which scan run to read, as an ISO-8601 timestamp. Only the UTC calendar day is
	// used to select the run; the time component is ignored. Defaults to the most
	// recent run when omitted.
	Date param.Opt[time.Time] `query:"date,omitzero" format:"date-time" json:"-"`
	// Maximum number of findings to return. Defaults to 25; clamped to 1–100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebScannerFindingsParams]'s query parameters as
// `url.Values`.
func (r WebScannerFindingsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebScannerCookiesParams struct {
	// Skip this many findings before returning. Use with `limit` for load-more paging.
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Which scan run to read, as an ISO-8601 timestamp. Only the UTC calendar day is
	// used to select the run; the time component is ignored. Defaults to the most
	// recent run when omitted.
	Date param.Opt[time.Time] `query:"date,omitzero" format:"date-time" json:"-"`
	// Maximum number of findings to return. Defaults to 25; clamped to 1–100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebScannerCookiesParams]'s query parameters as
// `url.Values`.
func (r WebScannerCookiesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebScannerSummaryParams struct {
	// Which scan run to read, as an ISO-8601 timestamp. Only the UTC calendar day is
	// used to select the run; the time component is ignored. Defaults to the most
	// recent run when omitted.
	Date param.Opt[time.Time] `query:"date,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [WebScannerSummaryParams]'s query parameters as
// `url.Values`.
func (r WebScannerSummaryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
