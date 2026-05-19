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
	"github.com/with-ours/platform-sdk-go/packages/param"
	"github.com/with-ours/platform-sdk-go/packages/respjson"
)

// WebScannerRuleService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebScannerRuleService] method instead.
type WebScannerRuleService struct {
	options []option.RequestOption
}

// NewWebScannerRuleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebScannerRuleService(opts ...option.RequestOption) (r WebScannerRuleService) {
	r = WebScannerRuleService{}
	r.options = opts
	return
}

// List suppression rules for a single web scanner. Requires the `scannerId` query
// parameter — rules are always scoped to a parent scanner. Not paginated; the
// per-scanner rule count is bounded. Requires scope: webScanner:find
func (r *WebScannerRuleService) List(ctx context.Context, query WebScannerRuleListParams, opts ...option.RequestOption) (res *WebScannerRuleListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "rest/v1/web-scanner-rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Create a suppression rule on a web scanner. Auth is enforced against the parent
// scanner via `webScanner:update`. At least one of `cookiePatterns`,
// `domainPatterns`, or `scriptPatterns` should be set for the rule to match
// anything; omitted pattern arrays default to `[]`. Requires scope:
// webScanner:update
func (r *WebScannerRuleService) New(ctx context.Context, body WebScannerRuleNewParams, opts ...option.RequestOption) (res *WebScannerRuleNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "rest/v1/web-scanner-rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Find a single web scanner rule by ID. Requires scope: webScanner:find
func (r *WebScannerRuleService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *WebScannerRuleGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/web-scanner-rules/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Partially update a suppression rule. Only the fields you send are changed.
// List-valued fields (`cookiePatterns`, `domainPatterns`, `scriptPatterns`) are
// replaced wholesale when sent. Requires scope: webScanner:update
func (r *WebScannerRuleService) Update(ctx context.Context, id string, body WebScannerRuleUpdateParams, opts ...option.RequestOption) (res *WebScannerRuleUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/web-scanner-rules/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete a web scanner rule. Requires scope: webScanner:update
func (r *WebScannerRuleService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *WebScannerRuleDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/web-scanner-rules/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type WebScannerRuleListResponse struct {
	Entities []WebScannerRuleListResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerRuleListResponse) RawJSON() string { return r.JSON.raw }
func (r *WebScannerRuleListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerRuleListResponseEntity struct {
	ID              string   `json:"id" api:"required"`
	AccountID       string   `json:"accountId" api:"required"`
	CookiePatterns  []string `json:"cookiePatterns" api:"required"`
	CreatedAt       string   `json:"createdAt" api:"required"`
	DomainPatterns  []string `json:"domainPatterns" api:"required"`
	Name            string   `json:"name" api:"required"`
	Priority        int64    `json:"priority" api:"required"`
	ScannerID       string   `json:"scannerId" api:"required"`
	ScriptPatterns  []string `json:"scriptPatterns" api:"required"`
	CreatedByUserID string   `json:"createdByUserId" api:"nullable"`
	Notes           string   `json:"notes" api:"nullable"`
	// Any of "approved", "baa", "compliant", "firstParty", "ignore", "internal",
	// "other".
	Reason          string `json:"reason" api:"nullable"`
	UpdatedAt       string `json:"updatedAt" api:"nullable"`
	UpdatedByUserID string `json:"updatedByUserId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AccountID       respjson.Field
		CookiePatterns  respjson.Field
		CreatedAt       respjson.Field
		DomainPatterns  respjson.Field
		Name            respjson.Field
		Priority        respjson.Field
		ScannerID       respjson.Field
		ScriptPatterns  respjson.Field
		CreatedByUserID respjson.Field
		Notes           respjson.Field
		Reason          respjson.Field
		UpdatedAt       respjson.Field
		UpdatedByUserID respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerRuleListResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *WebScannerRuleListResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerRuleNewResponse struct {
	ID              string   `json:"id" api:"required"`
	AccountID       string   `json:"accountId" api:"required"`
	CookiePatterns  []string `json:"cookiePatterns" api:"required"`
	CreatedAt       string   `json:"createdAt" api:"required"`
	DomainPatterns  []string `json:"domainPatterns" api:"required"`
	Name            string   `json:"name" api:"required"`
	Priority        int64    `json:"priority" api:"required"`
	ScannerID       string   `json:"scannerId" api:"required"`
	ScriptPatterns  []string `json:"scriptPatterns" api:"required"`
	CreatedByUserID string   `json:"createdByUserId" api:"nullable"`
	Notes           string   `json:"notes" api:"nullable"`
	// Any of "approved", "baa", "compliant", "firstParty", "ignore", "internal",
	// "other".
	Reason          WebScannerRuleNewResponseReason `json:"reason" api:"nullable"`
	UpdatedAt       string                          `json:"updatedAt" api:"nullable"`
	UpdatedByUserID string                          `json:"updatedByUserId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AccountID       respjson.Field
		CookiePatterns  respjson.Field
		CreatedAt       respjson.Field
		DomainPatterns  respjson.Field
		Name            respjson.Field
		Priority        respjson.Field
		ScannerID       respjson.Field
		ScriptPatterns  respjson.Field
		CreatedByUserID respjson.Field
		Notes           respjson.Field
		Reason          respjson.Field
		UpdatedAt       respjson.Field
		UpdatedByUserID respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerRuleNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WebScannerRuleNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerRuleNewResponseReason string

const (
	WebScannerRuleNewResponseReasonApproved   WebScannerRuleNewResponseReason = "approved"
	WebScannerRuleNewResponseReasonBaa        WebScannerRuleNewResponseReason = "baa"
	WebScannerRuleNewResponseReasonCompliant  WebScannerRuleNewResponseReason = "compliant"
	WebScannerRuleNewResponseReasonFirstParty WebScannerRuleNewResponseReason = "firstParty"
	WebScannerRuleNewResponseReasonIgnore     WebScannerRuleNewResponseReason = "ignore"
	WebScannerRuleNewResponseReasonInternal   WebScannerRuleNewResponseReason = "internal"
	WebScannerRuleNewResponseReasonOther      WebScannerRuleNewResponseReason = "other"
)

type WebScannerRuleGetResponse struct {
	ID              string   `json:"id" api:"required"`
	AccountID       string   `json:"accountId" api:"required"`
	CookiePatterns  []string `json:"cookiePatterns" api:"required"`
	CreatedAt       string   `json:"createdAt" api:"required"`
	DomainPatterns  []string `json:"domainPatterns" api:"required"`
	Name            string   `json:"name" api:"required"`
	Priority        int64    `json:"priority" api:"required"`
	ScannerID       string   `json:"scannerId" api:"required"`
	ScriptPatterns  []string `json:"scriptPatterns" api:"required"`
	CreatedByUserID string   `json:"createdByUserId" api:"nullable"`
	Notes           string   `json:"notes" api:"nullable"`
	// Any of "approved", "baa", "compliant", "firstParty", "ignore", "internal",
	// "other".
	Reason          WebScannerRuleGetResponseReason `json:"reason" api:"nullable"`
	UpdatedAt       string                          `json:"updatedAt" api:"nullable"`
	UpdatedByUserID string                          `json:"updatedByUserId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AccountID       respjson.Field
		CookiePatterns  respjson.Field
		CreatedAt       respjson.Field
		DomainPatterns  respjson.Field
		Name            respjson.Field
		Priority        respjson.Field
		ScannerID       respjson.Field
		ScriptPatterns  respjson.Field
		CreatedByUserID respjson.Field
		Notes           respjson.Field
		Reason          respjson.Field
		UpdatedAt       respjson.Field
		UpdatedByUserID respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerRuleGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WebScannerRuleGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerRuleGetResponseReason string

const (
	WebScannerRuleGetResponseReasonApproved   WebScannerRuleGetResponseReason = "approved"
	WebScannerRuleGetResponseReasonBaa        WebScannerRuleGetResponseReason = "baa"
	WebScannerRuleGetResponseReasonCompliant  WebScannerRuleGetResponseReason = "compliant"
	WebScannerRuleGetResponseReasonFirstParty WebScannerRuleGetResponseReason = "firstParty"
	WebScannerRuleGetResponseReasonIgnore     WebScannerRuleGetResponseReason = "ignore"
	WebScannerRuleGetResponseReasonInternal   WebScannerRuleGetResponseReason = "internal"
	WebScannerRuleGetResponseReasonOther      WebScannerRuleGetResponseReason = "other"
)

type WebScannerRuleUpdateResponse struct {
	ID              string   `json:"id" api:"required"`
	AccountID       string   `json:"accountId" api:"required"`
	CookiePatterns  []string `json:"cookiePatterns" api:"required"`
	CreatedAt       string   `json:"createdAt" api:"required"`
	DomainPatterns  []string `json:"domainPatterns" api:"required"`
	Name            string   `json:"name" api:"required"`
	Priority        int64    `json:"priority" api:"required"`
	ScannerID       string   `json:"scannerId" api:"required"`
	ScriptPatterns  []string `json:"scriptPatterns" api:"required"`
	CreatedByUserID string   `json:"createdByUserId" api:"nullable"`
	Notes           string   `json:"notes" api:"nullable"`
	// Any of "approved", "baa", "compliant", "firstParty", "ignore", "internal",
	// "other".
	Reason          WebScannerRuleUpdateResponseReason `json:"reason" api:"nullable"`
	UpdatedAt       string                             `json:"updatedAt" api:"nullable"`
	UpdatedByUserID string                             `json:"updatedByUserId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AccountID       respjson.Field
		CookiePatterns  respjson.Field
		CreatedAt       respjson.Field
		DomainPatterns  respjson.Field
		Name            respjson.Field
		Priority        respjson.Field
		ScannerID       respjson.Field
		ScriptPatterns  respjson.Field
		CreatedByUserID respjson.Field
		Notes           respjson.Field
		Reason          respjson.Field
		UpdatedAt       respjson.Field
		UpdatedByUserID respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebScannerRuleUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *WebScannerRuleUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerRuleUpdateResponseReason string

const (
	WebScannerRuleUpdateResponseReasonApproved   WebScannerRuleUpdateResponseReason = "approved"
	WebScannerRuleUpdateResponseReasonBaa        WebScannerRuleUpdateResponseReason = "baa"
	WebScannerRuleUpdateResponseReasonCompliant  WebScannerRuleUpdateResponseReason = "compliant"
	WebScannerRuleUpdateResponseReasonFirstParty WebScannerRuleUpdateResponseReason = "firstParty"
	WebScannerRuleUpdateResponseReasonIgnore     WebScannerRuleUpdateResponseReason = "ignore"
	WebScannerRuleUpdateResponseReasonInternal   WebScannerRuleUpdateResponseReason = "internal"
	WebScannerRuleUpdateResponseReasonOther      WebScannerRuleUpdateResponseReason = "other"
)

type WebScannerRuleDeleteResponse struct {
	// The id of the suppression rule that was deleted.
	ID string `json:"id" api:"required"`
	// True when the underlying mutation succeeded.
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
func (r WebScannerRuleDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *WebScannerRuleDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebScannerRuleListParams struct {
	// The web scanner whose suppression rules should be returned.
	ScannerID string `query:"scannerId" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [WebScannerRuleListParams]'s query parameters as
// `url.Values`.
func (r WebScannerRuleListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebScannerRuleNewParams struct {
	// User-friendly name for the suppression rule.
	Name string `json:"name" api:"required"`
	// Rule priority (1–10,000). Lower numbers are evaluated first when multiple rules
	// match.
	Priority int64 `json:"priority" api:"required"`
	// The web scanner this rule belongs to.
	ScannerID string `json:"scannerId" api:"required"`
	// Free-form notes about why this rule exists or what it covers. Trimmed
	// server-side; empty strings become `null`.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Why this rule was added. Surfaced in audit views. Send `null` to clear an
	// existing reason on patch.
	//
	// Any of "approved", "baa", "compliant", "firstParty", "ignore", "internal",
	// "other".
	Reason WebScannerRuleNewParamsReason `json:"reason,omitzero"`
	// Glob patterns matched against cookie names (e.g. `_ga*`). Max 100 entries. When
	// sent on PATCH, replaces the existing list wholesale.
	CookiePatterns []string `json:"cookiePatterns,omitzero"`
	// Glob patterns matched against cookie domain / script hostname (e.g.
	// `*.google-analytics.com`). Max 100 entries. When sent on PATCH, replaces the
	// existing list wholesale.
	DomainPatterns []string `json:"domainPatterns,omitzero"`
	// Glob patterns matched against full script URLs (e.g.
	// `https://www.googletagmanager.com/gtm.js?id=*`). Max 100 entries. When sent on
	// PATCH, replaces the existing list wholesale.
	ScriptPatterns []string `json:"scriptPatterns,omitzero"`
	paramObj
}

func (r WebScannerRuleNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WebScannerRuleNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScannerRuleNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Why this rule was added. Surfaced in audit views. Send `null` to clear an
// existing reason on patch.
type WebScannerRuleNewParamsReason string

const (
	WebScannerRuleNewParamsReasonApproved   WebScannerRuleNewParamsReason = "approved"
	WebScannerRuleNewParamsReasonBaa        WebScannerRuleNewParamsReason = "baa"
	WebScannerRuleNewParamsReasonCompliant  WebScannerRuleNewParamsReason = "compliant"
	WebScannerRuleNewParamsReasonFirstParty WebScannerRuleNewParamsReason = "firstParty"
	WebScannerRuleNewParamsReasonIgnore     WebScannerRuleNewParamsReason = "ignore"
	WebScannerRuleNewParamsReasonInternal   WebScannerRuleNewParamsReason = "internal"
	WebScannerRuleNewParamsReasonOther      WebScannerRuleNewParamsReason = "other"
)

type WebScannerRuleUpdateParams struct {
	// Free-form notes about why this rule exists or what it covers. Trimmed
	// server-side; empty strings become `null`.
	Notes    param.Opt[string] `json:"notes,omitzero"`
	Name     param.Opt[string] `json:"name,omitzero"`
	Priority param.Opt[int64]  `json:"priority,omitzero"`
	// Why this rule was added. Surfaced in audit views. Send `null` to clear an
	// existing reason on patch.
	//
	// Any of "approved", "baa", "compliant", "firstParty", "ignore", "internal",
	// "other".
	Reason WebScannerRuleUpdateParamsReason `json:"reason,omitzero"`
	// Glob patterns matched against cookie names (e.g. `_ga*`). Max 100 entries. When
	// sent on PATCH, replaces the existing list wholesale.
	CookiePatterns []string `json:"cookiePatterns,omitzero"`
	// Glob patterns matched against cookie domain / script hostname (e.g.
	// `*.google-analytics.com`). Max 100 entries. When sent on PATCH, replaces the
	// existing list wholesale.
	DomainPatterns []string `json:"domainPatterns,omitzero"`
	// Glob patterns matched against full script URLs (e.g.
	// `https://www.googletagmanager.com/gtm.js?id=*`). Max 100 entries. When sent on
	// PATCH, replaces the existing list wholesale.
	ScriptPatterns []string `json:"scriptPatterns,omitzero"`
	paramObj
}

func (r WebScannerRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WebScannerRuleUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebScannerRuleUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Why this rule was added. Surfaced in audit views. Send `null` to clear an
// existing reason on patch.
type WebScannerRuleUpdateParamsReason string

const (
	WebScannerRuleUpdateParamsReasonApproved   WebScannerRuleUpdateParamsReason = "approved"
	WebScannerRuleUpdateParamsReasonBaa        WebScannerRuleUpdateParamsReason = "baa"
	WebScannerRuleUpdateParamsReasonCompliant  WebScannerRuleUpdateParamsReason = "compliant"
	WebScannerRuleUpdateParamsReasonFirstParty WebScannerRuleUpdateParamsReason = "firstParty"
	WebScannerRuleUpdateParamsReasonIgnore     WebScannerRuleUpdateParamsReason = "ignore"
	WebScannerRuleUpdateParamsReasonInternal   WebScannerRuleUpdateParamsReason = "internal"
	WebScannerRuleUpdateParamsReasonOther      WebScannerRuleUpdateParamsReason = "other"
)
