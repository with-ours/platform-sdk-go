// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy

import (
	"context"
	"encoding/json"
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

// SourceService contains methods and other services that help with interacting
// with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSourceService] method instead.
type SourceService struct {
	Options []option.RequestOption
}

// NewSourceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSourceService(opts ...option.RequestOption) (r SourceService) {
	r = SourceService{}
	r.Options = opts
	return
}

// List all sources for this account. Supports cursor pagination and optional
// filters for `type`, `status`, and `nameContains`. Results are sorted by creation
// date descending. Requires scope: source:list
func (r *SourceService) List(ctx context.Context, query SourceListParams, opts ...option.RequestOption) (res *pagination.Cursor[SourceListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "rest/v1/sources"
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

// List all sources for this account. Supports cursor pagination and optional
// filters for `type`, `status`, and `nameContains`. Results are sorted by creation
// date descending. Requires scope: source:list
func (r *SourceService) ListAutoPaging(ctx context.Context, query SourceListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[SourceListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Create a new source. Returns the full source entity (same shape as GET
// /sources/{id}) so callers can read all server-assigned fields without a
// follow-up GET. Requires scope: source:create
func (r *SourceService) New(ctx context.Context, body SourceNewParams, opts ...option.RequestOption) (res *SourceNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/sources"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Find a single source by ID. Requires scope: source:view
func (r *SourceService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *SourceGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/sources/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Partially update a source. Only the fields you send are changed; omitted fields
// are unchanged. Send explicit `null` to clear a nullable field. Returns the full
// source entity after the update. Requires scope: source:update
func (r *SourceService) Update(ctx context.Context, id string, body SourceUpdateParams, opts ...option.RequestOption) (res *SourceUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/sources/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete a source. Requires scope: source:delete
func (r *SourceService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *SourceDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/sources/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Returns the install or ingest tokens for a source. Pixel sources (WebSource,
// PixelImage, HTTPApiSource) return
// `{ sourceType: "pixel", token, testToken, installScript, testInstallScript }`.
// Webhook sources (Webhook, CallRail, Formstack, Healthie, etc.) return
// `{ sourceType: "webhook", token, testToken, ingestUrl, testIngestUrl, sampleCurl }`.
// Requires scope: source:view
func (r *SourceService) Tokens(ctx context.Context, id string, opts ...option.RequestOption) (res *SourceTokensResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/sources/%s/tokens", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type SourceListResponse struct {
	ID string `json:"id" api:"required"`
	// Organization id that owns this source.
	AccountID string `json:"accountId" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	// Any of "Disabled", "Enabled".
	Status SourceListResponseStatus `json:"status" api:"required"`
	// Any of "AlchemerWebhook", "AndroidNativeApi", "CSharpApi", "CalComWebhooks",
	// "CalendlyWebhook", "CallRail", "CallTrackingMetrics", "DotNetApi",
	// "FacebookLeadAds", "FormsortWebhooks", "Formstack", "GoLangApi",
	// "HTTPApiSource", "Healthie", "HubspotAppActions", "HubspotFormWebhook",
	// "JotFormWebhooks", "KotlinApi", "NodejsApi", "PHPApi", "PixelImage",
	// "PythonApi", "ReactNativeApi", "RedirectSource", "RubyApi", "SegmentWebPlugin",
	// "TypeformWebhooks", "WebSource", "Webhook", "WhatConverts", "iOSNativeApi".
	Type                  SourceListResponseType `json:"type" api:"required"`
	BotControlMode        string                 `json:"botControlMode" api:"nullable"`
	BotScoreThreshold     float64                `json:"botScoreThreshold" api:"nullable"`
	ExcludeRequestContext bool                   `json:"excludeRequestContext" api:"nullable"`
	// Whether this source exists in the currently published version. A source that is
	// not published will not accept events.
	IsPublished bool `json:"isPublished" api:"nullable"`
	// ISO-8601 timestamp of the most recent event from this source successfully
	// dispatched to a destination.
	LastDispatchedAt string `json:"lastDispatchedAt" api:"nullable"`
	// ISO-8601 timestamp of the most recent inbound request received by this source.
	// Useful for debugging "is my webhook even reaching us?"
	LastTriggeredAt       string `json:"lastTriggeredAt" api:"nullable"`
	Name                  string `json:"name" api:"nullable"`
	ProbabilisticIdentity any    `json:"probabilisticIdentity" api:"nullable"`
	ProjectAPIKey         string `json:"projectAPIKey" api:"nullable"`
	RedirectURL           string `json:"redirectUrl" api:"nullable"`
	SelectedAccountID     string `json:"selectedAccountId" api:"nullable"`
	// Limits which domains can send events to the CDP. When set, only requests from
	// these domains are accepted for this source. Separate from experiment settings
	// `whitelistDomains`, which limits which domains can load your experiments.
	WhitelistDomains []string `json:"whitelistDomains" api:"nullable"`
	WhitelistIPs     []string `json:"whitelistIps" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		AccountID             respjson.Field
		CreatedAt             respjson.Field
		Status                respjson.Field
		Type                  respjson.Field
		BotControlMode        respjson.Field
		BotScoreThreshold     respjson.Field
		ExcludeRequestContext respjson.Field
		IsPublished           respjson.Field
		LastDispatchedAt      respjson.Field
		LastTriggeredAt       respjson.Field
		Name                  respjson.Field
		ProbabilisticIdentity respjson.Field
		ProjectAPIKey         respjson.Field
		RedirectURL           respjson.Field
		SelectedAccountID     respjson.Field
		WhitelistDomains      respjson.Field
		WhitelistIPs          respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceListResponse) RawJSON() string { return r.JSON.raw }
func (r *SourceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SourceListResponseStatus string

const (
	SourceListResponseStatusDisabled SourceListResponseStatus = "Disabled"
	SourceListResponseStatusEnabled  SourceListResponseStatus = "Enabled"
)

type SourceListResponseType string

const (
	SourceListResponseTypeAlchemerWebhook     SourceListResponseType = "AlchemerWebhook"
	SourceListResponseTypeAndroidNativeAPI    SourceListResponseType = "AndroidNativeApi"
	SourceListResponseTypeCSharpAPI           SourceListResponseType = "CSharpApi"
	SourceListResponseTypeCalComWebhooks      SourceListResponseType = "CalComWebhooks"
	SourceListResponseTypeCalendlyWebhook     SourceListResponseType = "CalendlyWebhook"
	SourceListResponseTypeCallRail            SourceListResponseType = "CallRail"
	SourceListResponseTypeCallTrackingMetrics SourceListResponseType = "CallTrackingMetrics"
	SourceListResponseTypeDotNetAPI           SourceListResponseType = "DotNetApi"
	SourceListResponseTypeFacebookLeadAds     SourceListResponseType = "FacebookLeadAds"
	SourceListResponseTypeFormsortWebhooks    SourceListResponseType = "FormsortWebhooks"
	SourceListResponseTypeFormstack           SourceListResponseType = "Formstack"
	SourceListResponseTypeGoLangAPI           SourceListResponseType = "GoLangApi"
	SourceListResponseTypeHTTPAPISource       SourceListResponseType = "HTTPApiSource"
	SourceListResponseTypeHealthie            SourceListResponseType = "Healthie"
	SourceListResponseTypeHubspotAppActions   SourceListResponseType = "HubspotAppActions"
	SourceListResponseTypeHubspotFormWebhook  SourceListResponseType = "HubspotFormWebhook"
	SourceListResponseTypeJotFormWebhooks     SourceListResponseType = "JotFormWebhooks"
	SourceListResponseTypeKotlinAPI           SourceListResponseType = "KotlinApi"
	SourceListResponseTypeNodejsAPI           SourceListResponseType = "NodejsApi"
	SourceListResponseTypePhpAPI              SourceListResponseType = "PHPApi"
	SourceListResponseTypePixelImage          SourceListResponseType = "PixelImage"
	SourceListResponseTypePythonAPI           SourceListResponseType = "PythonApi"
	SourceListResponseTypeReactNativeAPI      SourceListResponseType = "ReactNativeApi"
	SourceListResponseTypeRedirectSource      SourceListResponseType = "RedirectSource"
	SourceListResponseTypeRubyAPI             SourceListResponseType = "RubyApi"
	SourceListResponseTypeSegmentWebPlugin    SourceListResponseType = "SegmentWebPlugin"
	SourceListResponseTypeTypeformWebhooks    SourceListResponseType = "TypeformWebhooks"
	SourceListResponseTypeWebSource           SourceListResponseType = "WebSource"
	SourceListResponseTypeWebhook             SourceListResponseType = "Webhook"
	SourceListResponseTypeWhatConverts        SourceListResponseType = "WhatConverts"
	SourceListResponseTypeIOsNativeAPI        SourceListResponseType = "iOSNativeApi"
)

type SourceNewResponse struct {
	ID string `json:"id" api:"required"`
	// Organization id that owns this source.
	AccountID string `json:"accountId" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	// Any of "Disabled", "Enabled".
	Status SourceNewResponseStatus `json:"status" api:"required"`
	// Any of "AlchemerWebhook", "AndroidNativeApi", "CSharpApi", "CalComWebhooks",
	// "CalendlyWebhook", "CallRail", "CallTrackingMetrics", "DotNetApi",
	// "FacebookLeadAds", "FormsortWebhooks", "Formstack", "GoLangApi",
	// "HTTPApiSource", "Healthie", "HubspotAppActions", "HubspotFormWebhook",
	// "JotFormWebhooks", "KotlinApi", "NodejsApi", "PHPApi", "PixelImage",
	// "PythonApi", "ReactNativeApi", "RedirectSource", "RubyApi", "SegmentWebPlugin",
	// "TypeformWebhooks", "WebSource", "Webhook", "WhatConverts", "iOSNativeApi".
	Type                  SourceNewResponseType `json:"type" api:"required"`
	BotControlMode        string                `json:"botControlMode" api:"nullable"`
	BotScoreThreshold     float64               `json:"botScoreThreshold" api:"nullable"`
	ExcludeRequestContext bool                  `json:"excludeRequestContext" api:"nullable"`
	// Whether this source exists in the currently published version. A source that is
	// not published will not accept events.
	IsPublished bool `json:"isPublished" api:"nullable"`
	// ISO-8601 timestamp of the most recent event from this source successfully
	// dispatched to a destination.
	LastDispatchedAt string `json:"lastDispatchedAt" api:"nullable"`
	// ISO-8601 timestamp of the most recent inbound request received by this source.
	// Useful for debugging "is my webhook even reaching us?"
	LastTriggeredAt       string `json:"lastTriggeredAt" api:"nullable"`
	Name                  string `json:"name" api:"nullable"`
	ProbabilisticIdentity any    `json:"probabilisticIdentity" api:"nullable"`
	ProjectAPIKey         string `json:"projectAPIKey" api:"nullable"`
	RedirectURL           string `json:"redirectUrl" api:"nullable"`
	SelectedAccountID     string `json:"selectedAccountId" api:"nullable"`
	// Limits which domains can send events to the CDP. When set, only requests from
	// these domains are accepted for this source. Separate from experiment settings
	// `whitelistDomains`, which limits which domains can load your experiments.
	WhitelistDomains []string `json:"whitelistDomains" api:"nullable"`
	WhitelistIPs     []string `json:"whitelistIps" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		AccountID             respjson.Field
		CreatedAt             respjson.Field
		Status                respjson.Field
		Type                  respjson.Field
		BotControlMode        respjson.Field
		BotScoreThreshold     respjson.Field
		ExcludeRequestContext respjson.Field
		IsPublished           respjson.Field
		LastDispatchedAt      respjson.Field
		LastTriggeredAt       respjson.Field
		Name                  respjson.Field
		ProbabilisticIdentity respjson.Field
		ProjectAPIKey         respjson.Field
		RedirectURL           respjson.Field
		SelectedAccountID     respjson.Field
		WhitelistDomains      respjson.Field
		WhitelistIPs          respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SourceNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SourceNewResponseStatus string

const (
	SourceNewResponseStatusDisabled SourceNewResponseStatus = "Disabled"
	SourceNewResponseStatusEnabled  SourceNewResponseStatus = "Enabled"
)

type SourceNewResponseType string

const (
	SourceNewResponseTypeAlchemerWebhook     SourceNewResponseType = "AlchemerWebhook"
	SourceNewResponseTypeAndroidNativeAPI    SourceNewResponseType = "AndroidNativeApi"
	SourceNewResponseTypeCSharpAPI           SourceNewResponseType = "CSharpApi"
	SourceNewResponseTypeCalComWebhooks      SourceNewResponseType = "CalComWebhooks"
	SourceNewResponseTypeCalendlyWebhook     SourceNewResponseType = "CalendlyWebhook"
	SourceNewResponseTypeCallRail            SourceNewResponseType = "CallRail"
	SourceNewResponseTypeCallTrackingMetrics SourceNewResponseType = "CallTrackingMetrics"
	SourceNewResponseTypeDotNetAPI           SourceNewResponseType = "DotNetApi"
	SourceNewResponseTypeFacebookLeadAds     SourceNewResponseType = "FacebookLeadAds"
	SourceNewResponseTypeFormsortWebhooks    SourceNewResponseType = "FormsortWebhooks"
	SourceNewResponseTypeFormstack           SourceNewResponseType = "Formstack"
	SourceNewResponseTypeGoLangAPI           SourceNewResponseType = "GoLangApi"
	SourceNewResponseTypeHTTPAPISource       SourceNewResponseType = "HTTPApiSource"
	SourceNewResponseTypeHealthie            SourceNewResponseType = "Healthie"
	SourceNewResponseTypeHubspotAppActions   SourceNewResponseType = "HubspotAppActions"
	SourceNewResponseTypeHubspotFormWebhook  SourceNewResponseType = "HubspotFormWebhook"
	SourceNewResponseTypeJotFormWebhooks     SourceNewResponseType = "JotFormWebhooks"
	SourceNewResponseTypeKotlinAPI           SourceNewResponseType = "KotlinApi"
	SourceNewResponseTypeNodejsAPI           SourceNewResponseType = "NodejsApi"
	SourceNewResponseTypePhpAPI              SourceNewResponseType = "PHPApi"
	SourceNewResponseTypePixelImage          SourceNewResponseType = "PixelImage"
	SourceNewResponseTypePythonAPI           SourceNewResponseType = "PythonApi"
	SourceNewResponseTypeReactNativeAPI      SourceNewResponseType = "ReactNativeApi"
	SourceNewResponseTypeRedirectSource      SourceNewResponseType = "RedirectSource"
	SourceNewResponseTypeRubyAPI             SourceNewResponseType = "RubyApi"
	SourceNewResponseTypeSegmentWebPlugin    SourceNewResponseType = "SegmentWebPlugin"
	SourceNewResponseTypeTypeformWebhooks    SourceNewResponseType = "TypeformWebhooks"
	SourceNewResponseTypeWebSource           SourceNewResponseType = "WebSource"
	SourceNewResponseTypeWebhook             SourceNewResponseType = "Webhook"
	SourceNewResponseTypeWhatConverts        SourceNewResponseType = "WhatConverts"
	SourceNewResponseTypeIOsNativeAPI        SourceNewResponseType = "iOSNativeApi"
)

type SourceGetResponse struct {
	ID string `json:"id" api:"required"`
	// Organization id that owns this source.
	AccountID string `json:"accountId" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	// Any of "Disabled", "Enabled".
	Status SourceGetResponseStatus `json:"status" api:"required"`
	// Any of "AlchemerWebhook", "AndroidNativeApi", "CSharpApi", "CalComWebhooks",
	// "CalendlyWebhook", "CallRail", "CallTrackingMetrics", "DotNetApi",
	// "FacebookLeadAds", "FormsortWebhooks", "Formstack", "GoLangApi",
	// "HTTPApiSource", "Healthie", "HubspotAppActions", "HubspotFormWebhook",
	// "JotFormWebhooks", "KotlinApi", "NodejsApi", "PHPApi", "PixelImage",
	// "PythonApi", "ReactNativeApi", "RedirectSource", "RubyApi", "SegmentWebPlugin",
	// "TypeformWebhooks", "WebSource", "Webhook", "WhatConverts", "iOSNativeApi".
	Type                  SourceGetResponseType `json:"type" api:"required"`
	BotControlMode        string                `json:"botControlMode" api:"nullable"`
	BotScoreThreshold     float64               `json:"botScoreThreshold" api:"nullable"`
	ExcludeRequestContext bool                  `json:"excludeRequestContext" api:"nullable"`
	// Whether this source exists in the currently published version. A source that is
	// not published will not accept events.
	IsPublished bool `json:"isPublished" api:"nullable"`
	// ISO-8601 timestamp of the most recent event from this source successfully
	// dispatched to a destination.
	LastDispatchedAt string `json:"lastDispatchedAt" api:"nullable"`
	// ISO-8601 timestamp of the most recent inbound request received by this source.
	// Useful for debugging "is my webhook even reaching us?"
	LastTriggeredAt       string `json:"lastTriggeredAt" api:"nullable"`
	Name                  string `json:"name" api:"nullable"`
	ProbabilisticIdentity any    `json:"probabilisticIdentity" api:"nullable"`
	ProjectAPIKey         string `json:"projectAPIKey" api:"nullable"`
	RedirectURL           string `json:"redirectUrl" api:"nullable"`
	SelectedAccountID     string `json:"selectedAccountId" api:"nullable"`
	// Limits which domains can send events to the CDP. When set, only requests from
	// these domains are accepted for this source. Separate from experiment settings
	// `whitelistDomains`, which limits which domains can load your experiments.
	WhitelistDomains []string `json:"whitelistDomains" api:"nullable"`
	WhitelistIPs     []string `json:"whitelistIps" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		AccountID             respjson.Field
		CreatedAt             respjson.Field
		Status                respjson.Field
		Type                  respjson.Field
		BotControlMode        respjson.Field
		BotScoreThreshold     respjson.Field
		ExcludeRequestContext respjson.Field
		IsPublished           respjson.Field
		LastDispatchedAt      respjson.Field
		LastTriggeredAt       respjson.Field
		Name                  respjson.Field
		ProbabilisticIdentity respjson.Field
		ProjectAPIKey         respjson.Field
		RedirectURL           respjson.Field
		SelectedAccountID     respjson.Field
		WhitelistDomains      respjson.Field
		WhitelistIPs          respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SourceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SourceGetResponseStatus string

const (
	SourceGetResponseStatusDisabled SourceGetResponseStatus = "Disabled"
	SourceGetResponseStatusEnabled  SourceGetResponseStatus = "Enabled"
)

type SourceGetResponseType string

const (
	SourceGetResponseTypeAlchemerWebhook     SourceGetResponseType = "AlchemerWebhook"
	SourceGetResponseTypeAndroidNativeAPI    SourceGetResponseType = "AndroidNativeApi"
	SourceGetResponseTypeCSharpAPI           SourceGetResponseType = "CSharpApi"
	SourceGetResponseTypeCalComWebhooks      SourceGetResponseType = "CalComWebhooks"
	SourceGetResponseTypeCalendlyWebhook     SourceGetResponseType = "CalendlyWebhook"
	SourceGetResponseTypeCallRail            SourceGetResponseType = "CallRail"
	SourceGetResponseTypeCallTrackingMetrics SourceGetResponseType = "CallTrackingMetrics"
	SourceGetResponseTypeDotNetAPI           SourceGetResponseType = "DotNetApi"
	SourceGetResponseTypeFacebookLeadAds     SourceGetResponseType = "FacebookLeadAds"
	SourceGetResponseTypeFormsortWebhooks    SourceGetResponseType = "FormsortWebhooks"
	SourceGetResponseTypeFormstack           SourceGetResponseType = "Formstack"
	SourceGetResponseTypeGoLangAPI           SourceGetResponseType = "GoLangApi"
	SourceGetResponseTypeHTTPAPISource       SourceGetResponseType = "HTTPApiSource"
	SourceGetResponseTypeHealthie            SourceGetResponseType = "Healthie"
	SourceGetResponseTypeHubspotAppActions   SourceGetResponseType = "HubspotAppActions"
	SourceGetResponseTypeHubspotFormWebhook  SourceGetResponseType = "HubspotFormWebhook"
	SourceGetResponseTypeJotFormWebhooks     SourceGetResponseType = "JotFormWebhooks"
	SourceGetResponseTypeKotlinAPI           SourceGetResponseType = "KotlinApi"
	SourceGetResponseTypeNodejsAPI           SourceGetResponseType = "NodejsApi"
	SourceGetResponseTypePhpAPI              SourceGetResponseType = "PHPApi"
	SourceGetResponseTypePixelImage          SourceGetResponseType = "PixelImage"
	SourceGetResponseTypePythonAPI           SourceGetResponseType = "PythonApi"
	SourceGetResponseTypeReactNativeAPI      SourceGetResponseType = "ReactNativeApi"
	SourceGetResponseTypeRedirectSource      SourceGetResponseType = "RedirectSource"
	SourceGetResponseTypeRubyAPI             SourceGetResponseType = "RubyApi"
	SourceGetResponseTypeSegmentWebPlugin    SourceGetResponseType = "SegmentWebPlugin"
	SourceGetResponseTypeTypeformWebhooks    SourceGetResponseType = "TypeformWebhooks"
	SourceGetResponseTypeWebSource           SourceGetResponseType = "WebSource"
	SourceGetResponseTypeWebhook             SourceGetResponseType = "Webhook"
	SourceGetResponseTypeWhatConverts        SourceGetResponseType = "WhatConverts"
	SourceGetResponseTypeIOsNativeAPI        SourceGetResponseType = "iOSNativeApi"
)

type SourceUpdateResponse struct {
	ID string `json:"id" api:"required"`
	// Organization id that owns this source.
	AccountID string `json:"accountId" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	// Any of "Disabled", "Enabled".
	Status SourceUpdateResponseStatus `json:"status" api:"required"`
	// Any of "AlchemerWebhook", "AndroidNativeApi", "CSharpApi", "CalComWebhooks",
	// "CalendlyWebhook", "CallRail", "CallTrackingMetrics", "DotNetApi",
	// "FacebookLeadAds", "FormsortWebhooks", "Formstack", "GoLangApi",
	// "HTTPApiSource", "Healthie", "HubspotAppActions", "HubspotFormWebhook",
	// "JotFormWebhooks", "KotlinApi", "NodejsApi", "PHPApi", "PixelImage",
	// "PythonApi", "ReactNativeApi", "RedirectSource", "RubyApi", "SegmentWebPlugin",
	// "TypeformWebhooks", "WebSource", "Webhook", "WhatConverts", "iOSNativeApi".
	Type                  SourceUpdateResponseType `json:"type" api:"required"`
	BotControlMode        string                   `json:"botControlMode" api:"nullable"`
	BotScoreThreshold     float64                  `json:"botScoreThreshold" api:"nullable"`
	ExcludeRequestContext bool                     `json:"excludeRequestContext" api:"nullable"`
	// Whether this source exists in the currently published version. A source that is
	// not published will not accept events.
	IsPublished bool `json:"isPublished" api:"nullable"`
	// ISO-8601 timestamp of the most recent event from this source successfully
	// dispatched to a destination.
	LastDispatchedAt string `json:"lastDispatchedAt" api:"nullable"`
	// ISO-8601 timestamp of the most recent inbound request received by this source.
	// Useful for debugging "is my webhook even reaching us?"
	LastTriggeredAt       string `json:"lastTriggeredAt" api:"nullable"`
	Name                  string `json:"name" api:"nullable"`
	ProbabilisticIdentity any    `json:"probabilisticIdentity" api:"nullable"`
	ProjectAPIKey         string `json:"projectAPIKey" api:"nullable"`
	RedirectURL           string `json:"redirectUrl" api:"nullable"`
	SelectedAccountID     string `json:"selectedAccountId" api:"nullable"`
	// Limits which domains can send events to the CDP. When set, only requests from
	// these domains are accepted for this source. Separate from experiment settings
	// `whitelistDomains`, which limits which domains can load your experiments.
	WhitelistDomains []string `json:"whitelistDomains" api:"nullable"`
	WhitelistIPs     []string `json:"whitelistIps" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		AccountID             respjson.Field
		CreatedAt             respjson.Field
		Status                respjson.Field
		Type                  respjson.Field
		BotControlMode        respjson.Field
		BotScoreThreshold     respjson.Field
		ExcludeRequestContext respjson.Field
		IsPublished           respjson.Field
		LastDispatchedAt      respjson.Field
		LastTriggeredAt       respjson.Field
		Name                  respjson.Field
		ProbabilisticIdentity respjson.Field
		ProjectAPIKey         respjson.Field
		RedirectURL           respjson.Field
		SelectedAccountID     respjson.Field
		WhitelistDomains      respjson.Field
		WhitelistIPs          respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *SourceUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SourceUpdateResponseStatus string

const (
	SourceUpdateResponseStatusDisabled SourceUpdateResponseStatus = "Disabled"
	SourceUpdateResponseStatusEnabled  SourceUpdateResponseStatus = "Enabled"
)

type SourceUpdateResponseType string

const (
	SourceUpdateResponseTypeAlchemerWebhook     SourceUpdateResponseType = "AlchemerWebhook"
	SourceUpdateResponseTypeAndroidNativeAPI    SourceUpdateResponseType = "AndroidNativeApi"
	SourceUpdateResponseTypeCSharpAPI           SourceUpdateResponseType = "CSharpApi"
	SourceUpdateResponseTypeCalComWebhooks      SourceUpdateResponseType = "CalComWebhooks"
	SourceUpdateResponseTypeCalendlyWebhook     SourceUpdateResponseType = "CalendlyWebhook"
	SourceUpdateResponseTypeCallRail            SourceUpdateResponseType = "CallRail"
	SourceUpdateResponseTypeCallTrackingMetrics SourceUpdateResponseType = "CallTrackingMetrics"
	SourceUpdateResponseTypeDotNetAPI           SourceUpdateResponseType = "DotNetApi"
	SourceUpdateResponseTypeFacebookLeadAds     SourceUpdateResponseType = "FacebookLeadAds"
	SourceUpdateResponseTypeFormsortWebhooks    SourceUpdateResponseType = "FormsortWebhooks"
	SourceUpdateResponseTypeFormstack           SourceUpdateResponseType = "Formstack"
	SourceUpdateResponseTypeGoLangAPI           SourceUpdateResponseType = "GoLangApi"
	SourceUpdateResponseTypeHTTPAPISource       SourceUpdateResponseType = "HTTPApiSource"
	SourceUpdateResponseTypeHealthie            SourceUpdateResponseType = "Healthie"
	SourceUpdateResponseTypeHubspotAppActions   SourceUpdateResponseType = "HubspotAppActions"
	SourceUpdateResponseTypeHubspotFormWebhook  SourceUpdateResponseType = "HubspotFormWebhook"
	SourceUpdateResponseTypeJotFormWebhooks     SourceUpdateResponseType = "JotFormWebhooks"
	SourceUpdateResponseTypeKotlinAPI           SourceUpdateResponseType = "KotlinApi"
	SourceUpdateResponseTypeNodejsAPI           SourceUpdateResponseType = "NodejsApi"
	SourceUpdateResponseTypePhpAPI              SourceUpdateResponseType = "PHPApi"
	SourceUpdateResponseTypePixelImage          SourceUpdateResponseType = "PixelImage"
	SourceUpdateResponseTypePythonAPI           SourceUpdateResponseType = "PythonApi"
	SourceUpdateResponseTypeReactNativeAPI      SourceUpdateResponseType = "ReactNativeApi"
	SourceUpdateResponseTypeRedirectSource      SourceUpdateResponseType = "RedirectSource"
	SourceUpdateResponseTypeRubyAPI             SourceUpdateResponseType = "RubyApi"
	SourceUpdateResponseTypeSegmentWebPlugin    SourceUpdateResponseType = "SegmentWebPlugin"
	SourceUpdateResponseTypeTypeformWebhooks    SourceUpdateResponseType = "TypeformWebhooks"
	SourceUpdateResponseTypeWebSource           SourceUpdateResponseType = "WebSource"
	SourceUpdateResponseTypeWebhook             SourceUpdateResponseType = "Webhook"
	SourceUpdateResponseTypeWhatConverts        SourceUpdateResponseType = "WhatConverts"
	SourceUpdateResponseTypeIOsNativeAPI        SourceUpdateResponseType = "iOSNativeApi"
)

type SourceDeleteResponse struct {
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
func (r SourceDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *SourceDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SourceTokensResponseUnion contains all possible properties and values from
// [SourceTokensResponseObject], [SourceTokensResponseObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SourceTokensResponseUnion struct {
	Token string `json:"token"`
	// This field is from variant [SourceTokensResponseObject].
	InstallScript string `json:"installScript"`
	SourceType    string `json:"sourceType"`
	// This field is from variant [SourceTokensResponseObject].
	TestInstallScript string `json:"testInstallScript"`
	TestToken         string `json:"testToken"`
	// This field is from variant [SourceTokensResponseObject2].
	IngestURL string `json:"ingestUrl"`
	// This field is from variant [SourceTokensResponseObject2].
	SampleCurl string `json:"sampleCurl"`
	// This field is from variant [SourceTokensResponseObject2].
	TestIngestURL string `json:"testIngestUrl"`
	JSON          struct {
		Token             respjson.Field
		InstallScript     respjson.Field
		SourceType        respjson.Field
		TestInstallScript respjson.Field
		TestToken         respjson.Field
		IngestURL         respjson.Field
		SampleCurl        respjson.Field
		TestIngestURL     respjson.Field
		raw               string
	} `json:"-"`
}

func (u SourceTokensResponseUnion) AsSourceTokensResponseObject() (v SourceTokensResponseObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SourceTokensResponseUnion) AsSourceTokensResponseObject2() (v SourceTokensResponseObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SourceTokensResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *SourceTokensResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SourceTokensResponseObject struct {
	// Install token for the source.
	Token string `json:"token" api:"required"`
	// Ready-to-paste install snippet for the production token, including linked
	// runtime tokens for supported modules.
	InstallScript string `json:"installScript" api:"required"`
	// Discriminator: this is a pixel/web source.
	//
	// Any of "pixel".
	SourceType string `json:"sourceType" api:"required"`
	// Ready-to-paste install snippet for the test token, suitable for validation
	// before a live install.
	TestInstallScript string `json:"testInstallScript" api:"required"`
	// Test-mode token derived from `token`.
	TestToken string `json:"testToken" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token             respjson.Field
		InstallScript     respjson.Field
		SourceType        respjson.Field
		TestInstallScript respjson.Field
		TestToken         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceTokensResponseObject) RawJSON() string { return r.JSON.raw }
func (r *SourceTokensResponseObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SourceTokensResponseObject2 struct {
	// Source token (the source id).
	Token string `json:"token" api:"required"`
	// Production ingest URL for the webhook source.
	IngestURL string `json:"ingestUrl" api:"required"`
	// Example curl command showing how to POST a sample event to the ingest URL. Copy
	// and run to verify connectivity.
	SampleCurl string `json:"sampleCurl" api:"required"`
	// Discriminator: this is a webhook source.
	//
	// Any of "webhook".
	SourceType string `json:"sourceType" api:"required"`
	// Test-mode ingest URL.
	TestIngestURL string `json:"testIngestUrl" api:"required"`
	// Test-mode token derived from `token`.
	TestToken string `json:"testToken" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token         respjson.Field
		IngestURL     respjson.Field
		SampleCurl    respjson.Field
		SourceType    respjson.Field
		TestIngestURL respjson.Field
		TestToken     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceTokensResponseObject2) RawJSON() string { return r.JSON.raw }
func (r *SourceTokensResponseObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SourceListParams struct {
	// Maximum number of items to return. Defaults to 25; values below 1 are clamped to
	// 1 and values above 100 are clamped to 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Opaque pagination cursor from pagination.nextCursor in the previous response. Do
	// not decode or modify it. Malformed cursors return 400 Bad Request.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Case-insensitive substring filter on the source name.
	NameContains param.Opt[string] `query:"nameContains,omitzero" json:"-"`
	// Filter by source status.
	//
	// Any of "Disabled", "Enabled".
	Status SourceListParamsStatus `query:"status,omitzero" json:"-"`
	// Filter by source type.
	//
	// Any of "AlchemerWebhook", "AndroidNativeApi", "CSharpApi", "CalComWebhooks",
	// "CalendlyWebhook", "CallRail", "CallTrackingMetrics", "DotNetApi",
	// "FacebookLeadAds", "FormsortWebhooks", "Formstack", "GoLangApi",
	// "HTTPApiSource", "Healthie", "HubspotAppActions", "HubspotFormWebhook",
	// "JotFormWebhooks", "KotlinApi", "NodejsApi", "PHPApi", "PixelImage",
	// "PythonApi", "ReactNativeApi", "RedirectSource", "RubyApi", "SegmentWebPlugin",
	// "TypeformWebhooks", "WebSource", "Webhook", "WhatConverts", "iOSNativeApi".
	Type SourceListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SourceListParams]'s query parameters as `url.Values`.
func (r SourceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by source status.
type SourceListParamsStatus string

const (
	SourceListParamsStatusDisabled SourceListParamsStatus = "Disabled"
	SourceListParamsStatusEnabled  SourceListParamsStatus = "Enabled"
)

// Filter by source type.
type SourceListParamsType string

const (
	SourceListParamsTypeAlchemerWebhook     SourceListParamsType = "AlchemerWebhook"
	SourceListParamsTypeAndroidNativeAPI    SourceListParamsType = "AndroidNativeApi"
	SourceListParamsTypeCSharpAPI           SourceListParamsType = "CSharpApi"
	SourceListParamsTypeCalComWebhooks      SourceListParamsType = "CalComWebhooks"
	SourceListParamsTypeCalendlyWebhook     SourceListParamsType = "CalendlyWebhook"
	SourceListParamsTypeCallRail            SourceListParamsType = "CallRail"
	SourceListParamsTypeCallTrackingMetrics SourceListParamsType = "CallTrackingMetrics"
	SourceListParamsTypeDotNetAPI           SourceListParamsType = "DotNetApi"
	SourceListParamsTypeFacebookLeadAds     SourceListParamsType = "FacebookLeadAds"
	SourceListParamsTypeFormsortWebhooks    SourceListParamsType = "FormsortWebhooks"
	SourceListParamsTypeFormstack           SourceListParamsType = "Formstack"
	SourceListParamsTypeGoLangAPI           SourceListParamsType = "GoLangApi"
	SourceListParamsTypeHTTPAPISource       SourceListParamsType = "HTTPApiSource"
	SourceListParamsTypeHealthie            SourceListParamsType = "Healthie"
	SourceListParamsTypeHubspotAppActions   SourceListParamsType = "HubspotAppActions"
	SourceListParamsTypeHubspotFormWebhook  SourceListParamsType = "HubspotFormWebhook"
	SourceListParamsTypeJotFormWebhooks     SourceListParamsType = "JotFormWebhooks"
	SourceListParamsTypeKotlinAPI           SourceListParamsType = "KotlinApi"
	SourceListParamsTypeNodejsAPI           SourceListParamsType = "NodejsApi"
	SourceListParamsTypePhpAPI              SourceListParamsType = "PHPApi"
	SourceListParamsTypePixelImage          SourceListParamsType = "PixelImage"
	SourceListParamsTypePythonAPI           SourceListParamsType = "PythonApi"
	SourceListParamsTypeReactNativeAPI      SourceListParamsType = "ReactNativeApi"
	SourceListParamsTypeRedirectSource      SourceListParamsType = "RedirectSource"
	SourceListParamsTypeRubyAPI             SourceListParamsType = "RubyApi"
	SourceListParamsTypeSegmentWebPlugin    SourceListParamsType = "SegmentWebPlugin"
	SourceListParamsTypeTypeformWebhooks    SourceListParamsType = "TypeformWebhooks"
	SourceListParamsTypeWebSource           SourceListParamsType = "WebSource"
	SourceListParamsTypeWebhook             SourceListParamsType = "Webhook"
	SourceListParamsTypeWhatConverts        SourceListParamsType = "WhatConverts"
	SourceListParamsTypeIOsNativeAPI        SourceListParamsType = "iOSNativeApi"
)

type SourceNewParams struct {
	// Any of "AlchemerWebhook", "AndroidNativeApi", "CSharpApi", "CalComWebhooks",
	// "CalendlyWebhook", "CallRail", "CallTrackingMetrics", "DotNetApi",
	// "FacebookLeadAds", "FormsortWebhooks", "Formstack", "GoLangApi",
	// "HTTPApiSource", "Healthie", "HubspotAppActions", "HubspotFormWebhook",
	// "JotFormWebhooks", "KotlinApi", "NodejsApi", "PHPApi", "PixelImage",
	// "PythonApi", "ReactNativeApi", "RedirectSource", "RubyApi", "SegmentWebPlugin",
	// "TypeformWebhooks", "WebSource", "Webhook", "WhatConverts", "iOSNativeApi".
	Type SourceNewParamsType `json:"type,omitzero" api:"required"`
	Name param.Opt[string]   `json:"name,omitzero"`
	paramObj
}

func (r SourceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SourceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SourceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SourceNewParamsType string

const (
	SourceNewParamsTypeAlchemerWebhook     SourceNewParamsType = "AlchemerWebhook"
	SourceNewParamsTypeAndroidNativeAPI    SourceNewParamsType = "AndroidNativeApi"
	SourceNewParamsTypeCSharpAPI           SourceNewParamsType = "CSharpApi"
	SourceNewParamsTypeCalComWebhooks      SourceNewParamsType = "CalComWebhooks"
	SourceNewParamsTypeCalendlyWebhook     SourceNewParamsType = "CalendlyWebhook"
	SourceNewParamsTypeCallRail            SourceNewParamsType = "CallRail"
	SourceNewParamsTypeCallTrackingMetrics SourceNewParamsType = "CallTrackingMetrics"
	SourceNewParamsTypeDotNetAPI           SourceNewParamsType = "DotNetApi"
	SourceNewParamsTypeFacebookLeadAds     SourceNewParamsType = "FacebookLeadAds"
	SourceNewParamsTypeFormsortWebhooks    SourceNewParamsType = "FormsortWebhooks"
	SourceNewParamsTypeFormstack           SourceNewParamsType = "Formstack"
	SourceNewParamsTypeGoLangAPI           SourceNewParamsType = "GoLangApi"
	SourceNewParamsTypeHTTPAPISource       SourceNewParamsType = "HTTPApiSource"
	SourceNewParamsTypeHealthie            SourceNewParamsType = "Healthie"
	SourceNewParamsTypeHubspotAppActions   SourceNewParamsType = "HubspotAppActions"
	SourceNewParamsTypeHubspotFormWebhook  SourceNewParamsType = "HubspotFormWebhook"
	SourceNewParamsTypeJotFormWebhooks     SourceNewParamsType = "JotFormWebhooks"
	SourceNewParamsTypeKotlinAPI           SourceNewParamsType = "KotlinApi"
	SourceNewParamsTypeNodejsAPI           SourceNewParamsType = "NodejsApi"
	SourceNewParamsTypePhpAPI              SourceNewParamsType = "PHPApi"
	SourceNewParamsTypePixelImage          SourceNewParamsType = "PixelImage"
	SourceNewParamsTypePythonAPI           SourceNewParamsType = "PythonApi"
	SourceNewParamsTypeReactNativeAPI      SourceNewParamsType = "ReactNativeApi"
	SourceNewParamsTypeRedirectSource      SourceNewParamsType = "RedirectSource"
	SourceNewParamsTypeRubyAPI             SourceNewParamsType = "RubyApi"
	SourceNewParamsTypeSegmentWebPlugin    SourceNewParamsType = "SegmentWebPlugin"
	SourceNewParamsTypeTypeformWebhooks    SourceNewParamsType = "TypeformWebhooks"
	SourceNewParamsTypeWebSource           SourceNewParamsType = "WebSource"
	SourceNewParamsTypeWebhook             SourceNewParamsType = "Webhook"
	SourceNewParamsTypeWhatConverts        SourceNewParamsType = "WhatConverts"
	SourceNewParamsTypeIOsNativeAPI        SourceNewParamsType = "iOSNativeApi"
)

type SourceUpdateParams struct {
	BotControlMode        param.Opt[string]  `json:"botControlMode,omitzero"`
	BotScoreThreshold     param.Opt[float64] `json:"botScoreThreshold,omitzero"`
	ExcludeRequestContext param.Opt[bool]    `json:"excludeRequestContext,omitzero"`
	Name                  param.Opt[string]  `json:"name,omitzero"`
	ProjectAPIKey         param.Opt[string]  `json:"projectAPIKey,omitzero"`
	RedirectURL           param.Opt[string]  `json:"redirectUrl,omitzero"`
	SelectedAccountID     param.Opt[string]  `json:"selectedAccountId,omitzero"`
	Status                param.Opt[string]  `json:"status,omitzero"`
	ProbabilisticIdentity any                `json:"probabilisticIdentity,omitzero"`
	WhitelistDomains      []string           `json:"whitelistDomains,omitzero"`
	WhitelistIPs          []string           `json:"whitelistIps,omitzero"`
	paramObj
}

func (r SourceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SourceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SourceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
