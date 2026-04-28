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
	"github.com/with-ours/platform-sdk-go/internal/requestconfig"
	"github.com/with-ours/platform-sdk-go/option"
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

// Create a new source. Requires scope: source:create
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

// Update a source. Requires scope: source:update
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

// List all sources. Requires scope: source:list
func (r *SourceService) List(ctx context.Context, opts ...option.RequestOption) (res *SourceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/sources"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete a source. Requires scope: source:delete
func (r *SourceService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/sources/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type SourceNewResponse struct {
	ID        string `json:"id" api:"required"`
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
	Type SourceNewResponseType `json:"type" api:"required"`
	Name string                `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	ID        string `json:"id" api:"required"`
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
	Name                  string                `json:"name" api:"nullable"`
	ProbabilisticIdentity any                   `json:"probabilisticIdentity" api:"nullable"`
	ProjectAPIKey         string                `json:"projectAPIKey" api:"nullable"`
	RedirectURL           string                `json:"redirectUrl" api:"nullable"`
	SelectedAccountID     string                `json:"selectedAccountId" api:"nullable"`
	WhitelistDomains      []any                 `json:"whitelistDomains" api:"nullable"`
	WhitelistIPs          []any                 `json:"whitelistIps" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		CreatedAt             respjson.Field
		Status                respjson.Field
		Type                  respjson.Field
		BotControlMode        respjson.Field
		BotScoreThreshold     respjson.Field
		ExcludeRequestContext respjson.Field
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
	ID        string `json:"id" api:"required"`
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
	Type SourceUpdateResponseType `json:"type" api:"required"`
	Name string                   `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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

type SourceListResponse struct {
	Entities []SourceListResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceListResponse) RawJSON() string { return r.JSON.raw }
func (r *SourceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SourceListResponseEntity struct {
	ID        string `json:"id" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	// Any of "Disabled", "Enabled".
	Status string `json:"status" api:"required"`
	// Any of "AlchemerWebhook", "AndroidNativeApi", "CSharpApi", "CalComWebhooks",
	// "CalendlyWebhook", "CallRail", "CallTrackingMetrics", "DotNetApi",
	// "FacebookLeadAds", "FormsortWebhooks", "Formstack", "GoLangApi",
	// "HTTPApiSource", "Healthie", "HubspotAppActions", "HubspotFormWebhook",
	// "JotFormWebhooks", "KotlinApi", "NodejsApi", "PHPApi", "PixelImage",
	// "PythonApi", "ReactNativeApi", "RedirectSource", "RubyApi", "SegmentWebPlugin",
	// "TypeformWebhooks", "WebSource", "Webhook", "WhatConverts", "iOSNativeApi".
	Type string `json:"type" api:"required"`
	Name string `json:"name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SourceListResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *SourceListResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
	// Any of "Disabled", "Enabled".
	Status                SourceUpdateParamsStatus `json:"status,omitzero" api:"required"`
	BotControlMode        param.Opt[string]        `json:"botControlMode,omitzero"`
	BotScoreThreshold     param.Opt[float64]       `json:"botScoreThreshold,omitzero"`
	ExcludeRequestContext param.Opt[bool]          `json:"excludeRequestContext,omitzero"`
	Name                  param.Opt[string]        `json:"name,omitzero"`
	ProjectAPIKey         param.Opt[string]        `json:"projectAPIKey,omitzero"`
	RedirectURL           param.Opt[string]        `json:"redirectUrl,omitzero"`
	SelectedAccountID     param.Opt[string]        `json:"selectedAccountId,omitzero"`
	ProbabilisticIdentity any                      `json:"probabilisticIdentity,omitzero"`
	WhitelistDomains      []any                    `json:"whitelistDomains,omitzero"`
	WhitelistIPs          []any                    `json:"whitelistIps,omitzero"`
	paramObj
}

func (r SourceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SourceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SourceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SourceUpdateParamsStatus string

const (
	SourceUpdateParamsStatusDisabled SourceUpdateParamsStatus = "Disabled"
	SourceUpdateParamsStatusEnabled  SourceUpdateParamsStatus = "Enabled"
)
