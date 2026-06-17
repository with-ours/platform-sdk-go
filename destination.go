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

// DestinationService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDestinationService] method instead.
type DestinationService struct {
	Options []option.RequestOption
}

// NewDestinationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDestinationService(opts ...option.RequestOption) (r DestinationService) {
	r = DestinationService{}
	r.Options = opts
	return
}

// List all destinations. Requires scope: destination:list
func (r *DestinationService) List(ctx context.Context, query DestinationListParams, opts ...option.RequestOption) (res *pagination.Cursor[DestinationListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "rest/v1/destinations"
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

// List all destinations. Requires scope: destination:list
func (r *DestinationService) ListAutoPaging(ctx context.Context, query DestinationListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[DestinationListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Create a new destination. Requires scope: destination:create
func (r *DestinationService) New(ctx context.Context, body DestinationNewParams, opts ...option.RequestOption) (res *DestinationNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/destinations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Find a single destination by ID. Requires scope: destination:find
func (r *DestinationService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *DestinationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/destinations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Partially update a destination. Only the fields you send are changed; omitted
// fields are unchanged. The `settings` object is patch-only: omitted keys keep
// their current value, and send `null` to clear a specific setting. Requires
// scope: destination:update
func (r *DestinationService) Update(ctx context.Context, id string, body DestinationUpdateParams, opts ...option.RequestOption) (res *DestinationUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/destinations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete a destination. Requires scope: destination:delete
func (r *DestinationService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/destinations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Lists every destination type the platform supports, with its human-readable
// label, capability flags (oauth, listsAccounts, supportsRenamedEvents), and the
// settings descriptor used to configure a destination of that type.
// Account-agnostic — the response is the same for every API key. Filter
// client-side to find a specific type (e.g. `Klaviyo`, `Facebook`). Requires
// scope: destination:list
func (r *DestinationService) Types(ctx context.Context, opts ...option.RequestOption) (res *DestinationTypesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/destinations/types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type DestinationListResponse struct {
	ID        string `json:"id" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	// Any of "Disabled", "Enabled".
	Status DestinationListResponseStatus `json:"status" api:"required"`
	// Destination type. Read responses may include warehouse or cloud-storage types
	// that are not creatable through POST /rest/v1/destinations.
	//
	// Any of "AWSEventBridge", "AWSKinesis", "AWSLambda", "AWSS3", "AWSSNS",
	// "ActiveCampaignApi", "Admitad", "AdobeAnalytics", "AmazonDSP", "Amplitude",
	// "AppLovin", "ArtsAI", "Attentive", "Audiohook", "AzureBlob", "BasisPostback",
	// "BeeswaxPostback", "BingAds", "BingAdsWeb", "Braze", "ConvertABTestingEvent",
	// "Customerio", "DatabricksWarehouse", "DomoWarehouse", "Everflow", "Facebook",
	// "FloodlightSGTM", "FullContact", "G4Analytics", "GA4MeasurementProtocol",
	// "GA4ServerProxy", "Google", "GoogleAds360", "GoogleAdsServerContainer",
	// "GoogleBigQuery", "GoogleBigQueryWarehouse", "GoogleDataManagerEventIngest",
	// "GooglePubSub", "GoogleStorage", "HTTPCustomRequest", "HTTPDestination",
	// "Hubspot", "IHeartMediaMagellan", "Impact", "Iterable", "Klaviyo",
	// "LinkedInAdsCAPI", "LiveIntent", "LiveRampWarehouse", "MNTN", "Mailchimp",
	// "Mixpanel", "NextdoorAds", "OpenAIAds", "OursSyntheticData", "Outbrain",
	// "Partnerize", "Pinterest", "Plausible", "Podscribe", "PostHog", "QuantcastCAPI",
	// "QuoraAds", "Reddit", "RokuCAPI", "SnapchatAdsCapi", "Spotify", "StackAdaptAPI",
	// "Taboola", "Tatari", "TheTradeDesk", "TikTok", "UniversalAds", "VWO", "Viant",
	// "ViantCAPI", "Vibe", "Woopra", "XAds", "YelpCAPI", "Zendesk", "ZohoCRM",
	// "ZoomInfo".
	Type               DestinationListResponseType `json:"type" api:"required"`
	HashingSalt        string                      `json:"hashingSalt" api:"nullable"`
	LimitedToSourceIDs []string                    `json:"limitedToSourceIds" api:"nullable"`
	Name               string                      `json:"name" api:"nullable"`
	Settings           any                         `json:"settings" api:"nullable"`
	UpdatedAt          string                      `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Status             respjson.Field
		Type               respjson.Field
		HashingSalt        respjson.Field
		LimitedToSourceIDs respjson.Field
		Name               respjson.Field
		Settings           respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DestinationListResponse) RawJSON() string { return r.JSON.raw }
func (r *DestinationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationListResponseStatus string

const (
	DestinationListResponseStatusDisabled DestinationListResponseStatus = "Disabled"
	DestinationListResponseStatusEnabled  DestinationListResponseStatus = "Enabled"
)

// Destination type. Read responses may include warehouse or cloud-storage types
// that are not creatable through POST /rest/v1/destinations.
type DestinationListResponseType string

const (
	DestinationListResponseTypeAwsEventBridge               DestinationListResponseType = "AWSEventBridge"
	DestinationListResponseTypeAwsKinesis                   DestinationListResponseType = "AWSKinesis"
	DestinationListResponseTypeAwsLambda                    DestinationListResponseType = "AWSLambda"
	DestinationListResponseTypeAwss3                        DestinationListResponseType = "AWSS3"
	DestinationListResponseTypeAwssns                       DestinationListResponseType = "AWSSNS"
	DestinationListResponseTypeActiveCampaignAPI            DestinationListResponseType = "ActiveCampaignApi"
	DestinationListResponseTypeAdmitad                      DestinationListResponseType = "Admitad"
	DestinationListResponseTypeAdobeAnalytics               DestinationListResponseType = "AdobeAnalytics"
	DestinationListResponseTypeAmazonDsp                    DestinationListResponseType = "AmazonDSP"
	DestinationListResponseTypeAmplitude                    DestinationListResponseType = "Amplitude"
	DestinationListResponseTypeAppLovin                     DestinationListResponseType = "AppLovin"
	DestinationListResponseTypeArtsAI                       DestinationListResponseType = "ArtsAI"
	DestinationListResponseTypeAttentive                    DestinationListResponseType = "Attentive"
	DestinationListResponseTypeAudiohook                    DestinationListResponseType = "Audiohook"
	DestinationListResponseTypeAzureBlob                    DestinationListResponseType = "AzureBlob"
	DestinationListResponseTypeBasisPostback                DestinationListResponseType = "BasisPostback"
	DestinationListResponseTypeBeeswaxPostback              DestinationListResponseType = "BeeswaxPostback"
	DestinationListResponseTypeBingAds                      DestinationListResponseType = "BingAds"
	DestinationListResponseTypeBingAdsWeb                   DestinationListResponseType = "BingAdsWeb"
	DestinationListResponseTypeBraze                        DestinationListResponseType = "Braze"
	DestinationListResponseTypeConvertAbTestingEvent        DestinationListResponseType = "ConvertABTestingEvent"
	DestinationListResponseTypeCustomerio                   DestinationListResponseType = "Customerio"
	DestinationListResponseTypeDatabricksWarehouse          DestinationListResponseType = "DatabricksWarehouse"
	DestinationListResponseTypeDomoWarehouse                DestinationListResponseType = "DomoWarehouse"
	DestinationListResponseTypeEverflow                     DestinationListResponseType = "Everflow"
	DestinationListResponseTypeFacebook                     DestinationListResponseType = "Facebook"
	DestinationListResponseTypeFloodlightSgtm               DestinationListResponseType = "FloodlightSGTM"
	DestinationListResponseTypeFullContact                  DestinationListResponseType = "FullContact"
	DestinationListResponseTypeG4Analytics                  DestinationListResponseType = "G4Analytics"
	DestinationListResponseTypeGa4MeasurementProtocol       DestinationListResponseType = "GA4MeasurementProtocol"
	DestinationListResponseTypeGa4ServerProxy               DestinationListResponseType = "GA4ServerProxy"
	DestinationListResponseTypeGoogle                       DestinationListResponseType = "Google"
	DestinationListResponseTypeGoogleAds360                 DestinationListResponseType = "GoogleAds360"
	DestinationListResponseTypeGoogleAdsServerContainer     DestinationListResponseType = "GoogleAdsServerContainer"
	DestinationListResponseTypeGoogleBigQuery               DestinationListResponseType = "GoogleBigQuery"
	DestinationListResponseTypeGoogleBigQueryWarehouse      DestinationListResponseType = "GoogleBigQueryWarehouse"
	DestinationListResponseTypeGoogleDataManagerEventIngest DestinationListResponseType = "GoogleDataManagerEventIngest"
	DestinationListResponseTypeGooglePubSub                 DestinationListResponseType = "GooglePubSub"
	DestinationListResponseTypeGoogleStorage                DestinationListResponseType = "GoogleStorage"
	DestinationListResponseTypeHTTPCustomRequest            DestinationListResponseType = "HTTPCustomRequest"
	DestinationListResponseTypeHTTPDestination              DestinationListResponseType = "HTTPDestination"
	DestinationListResponseTypeHubspot                      DestinationListResponseType = "Hubspot"
	DestinationListResponseTypeIHeartMediaMagellan          DestinationListResponseType = "IHeartMediaMagellan"
	DestinationListResponseTypeImpact                       DestinationListResponseType = "Impact"
	DestinationListResponseTypeIterable                     DestinationListResponseType = "Iterable"
	DestinationListResponseTypeKlaviyo                      DestinationListResponseType = "Klaviyo"
	DestinationListResponseTypeLinkedInAdsCapi              DestinationListResponseType = "LinkedInAdsCAPI"
	DestinationListResponseTypeLiveIntent                   DestinationListResponseType = "LiveIntent"
	DestinationListResponseTypeLiveRampWarehouse            DestinationListResponseType = "LiveRampWarehouse"
	DestinationListResponseTypeMntn                         DestinationListResponseType = "MNTN"
	DestinationListResponseTypeMailchimp                    DestinationListResponseType = "Mailchimp"
	DestinationListResponseTypeMixpanel                     DestinationListResponseType = "Mixpanel"
	DestinationListResponseTypeNextdoorAds                  DestinationListResponseType = "NextdoorAds"
	DestinationListResponseTypeOpenAIAds                    DestinationListResponseType = "OpenAIAds"
	DestinationListResponseTypeOursSyntheticData            DestinationListResponseType = "OursSyntheticData"
	DestinationListResponseTypeOutbrain                     DestinationListResponseType = "Outbrain"
	DestinationListResponseTypePartnerize                   DestinationListResponseType = "Partnerize"
	DestinationListResponseTypePinterest                    DestinationListResponseType = "Pinterest"
	DestinationListResponseTypePlausible                    DestinationListResponseType = "Plausible"
	DestinationListResponseTypePodscribe                    DestinationListResponseType = "Podscribe"
	DestinationListResponseTypePostHog                      DestinationListResponseType = "PostHog"
	DestinationListResponseTypeQuantcastCapi                DestinationListResponseType = "QuantcastCAPI"
	DestinationListResponseTypeQuoraAds                     DestinationListResponseType = "QuoraAds"
	DestinationListResponseTypeReddit                       DestinationListResponseType = "Reddit"
	DestinationListResponseTypeRokuCapi                     DestinationListResponseType = "RokuCAPI"
	DestinationListResponseTypeSnapchatAdsCapi              DestinationListResponseType = "SnapchatAdsCapi"
	DestinationListResponseTypeSpotify                      DestinationListResponseType = "Spotify"
	DestinationListResponseTypeStackAdaptAPI                DestinationListResponseType = "StackAdaptAPI"
	DestinationListResponseTypeTaboola                      DestinationListResponseType = "Taboola"
	DestinationListResponseTypeTatari                       DestinationListResponseType = "Tatari"
	DestinationListResponseTypeTheTradeDesk                 DestinationListResponseType = "TheTradeDesk"
	DestinationListResponseTypeTikTok                       DestinationListResponseType = "TikTok"
	DestinationListResponseTypeUniversalAds                 DestinationListResponseType = "UniversalAds"
	DestinationListResponseTypeVwo                          DestinationListResponseType = "VWO"
	DestinationListResponseTypeViant                        DestinationListResponseType = "Viant"
	DestinationListResponseTypeViantCapi                    DestinationListResponseType = "ViantCAPI"
	DestinationListResponseTypeVibe                         DestinationListResponseType = "Vibe"
	DestinationListResponseTypeWoopra                       DestinationListResponseType = "Woopra"
	DestinationListResponseTypeXAds                         DestinationListResponseType = "XAds"
	DestinationListResponseTypeYelpCapi                     DestinationListResponseType = "YelpCAPI"
	DestinationListResponseTypeZendesk                      DestinationListResponseType = "Zendesk"
	DestinationListResponseTypeZohoCRM                      DestinationListResponseType = "ZohoCRM"
	DestinationListResponseTypeZoomInfo                     DestinationListResponseType = "ZoomInfo"
)

type DestinationNewResponse struct {
	ID        string `json:"id" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	// Any of "Disabled", "Enabled".
	Status DestinationNewResponseStatus `json:"status" api:"required"`
	// Destination type. Read responses may include warehouse or cloud-storage types
	// that are not creatable through POST /rest/v1/destinations.
	//
	// Any of "AWSEventBridge", "AWSKinesis", "AWSLambda", "AWSS3", "AWSSNS",
	// "ActiveCampaignApi", "Admitad", "AdobeAnalytics", "AmazonDSP", "Amplitude",
	// "AppLovin", "ArtsAI", "Attentive", "Audiohook", "AzureBlob", "BasisPostback",
	// "BeeswaxPostback", "BingAds", "BingAdsWeb", "Braze", "ConvertABTestingEvent",
	// "Customerio", "DatabricksWarehouse", "DomoWarehouse", "Everflow", "Facebook",
	// "FloodlightSGTM", "FullContact", "G4Analytics", "GA4MeasurementProtocol",
	// "GA4ServerProxy", "Google", "GoogleAds360", "GoogleAdsServerContainer",
	// "GoogleBigQuery", "GoogleBigQueryWarehouse", "GoogleDataManagerEventIngest",
	// "GooglePubSub", "GoogleStorage", "HTTPCustomRequest", "HTTPDestination",
	// "Hubspot", "IHeartMediaMagellan", "Impact", "Iterable", "Klaviyo",
	// "LinkedInAdsCAPI", "LiveIntent", "LiveRampWarehouse", "MNTN", "Mailchimp",
	// "Mixpanel", "NextdoorAds", "OpenAIAds", "OursSyntheticData", "Outbrain",
	// "Partnerize", "Pinterest", "Plausible", "Podscribe", "PostHog", "QuantcastCAPI",
	// "QuoraAds", "Reddit", "RokuCAPI", "SnapchatAdsCapi", "Spotify", "StackAdaptAPI",
	// "Taboola", "Tatari", "TheTradeDesk", "TikTok", "UniversalAds", "VWO", "Viant",
	// "ViantCAPI", "Vibe", "Woopra", "XAds", "YelpCAPI", "Zendesk", "ZohoCRM",
	// "ZoomInfo".
	Type               DestinationNewResponseType `json:"type" api:"required"`
	HashingSalt        string                     `json:"hashingSalt" api:"nullable"`
	LimitedToSourceIDs []string                   `json:"limitedToSourceIds" api:"nullable"`
	Name               string                     `json:"name" api:"nullable"`
	Settings           any                        `json:"settings" api:"nullable"`
	UpdatedAt          string                     `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Status             respjson.Field
		Type               respjson.Field
		HashingSalt        respjson.Field
		LimitedToSourceIDs respjson.Field
		Name               respjson.Field
		Settings           respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DestinationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *DestinationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationNewResponseStatus string

const (
	DestinationNewResponseStatusDisabled DestinationNewResponseStatus = "Disabled"
	DestinationNewResponseStatusEnabled  DestinationNewResponseStatus = "Enabled"
)

// Destination type. Read responses may include warehouse or cloud-storage types
// that are not creatable through POST /rest/v1/destinations.
type DestinationNewResponseType string

const (
	DestinationNewResponseTypeAwsEventBridge               DestinationNewResponseType = "AWSEventBridge"
	DestinationNewResponseTypeAwsKinesis                   DestinationNewResponseType = "AWSKinesis"
	DestinationNewResponseTypeAwsLambda                    DestinationNewResponseType = "AWSLambda"
	DestinationNewResponseTypeAwss3                        DestinationNewResponseType = "AWSS3"
	DestinationNewResponseTypeAwssns                       DestinationNewResponseType = "AWSSNS"
	DestinationNewResponseTypeActiveCampaignAPI            DestinationNewResponseType = "ActiveCampaignApi"
	DestinationNewResponseTypeAdmitad                      DestinationNewResponseType = "Admitad"
	DestinationNewResponseTypeAdobeAnalytics               DestinationNewResponseType = "AdobeAnalytics"
	DestinationNewResponseTypeAmazonDsp                    DestinationNewResponseType = "AmazonDSP"
	DestinationNewResponseTypeAmplitude                    DestinationNewResponseType = "Amplitude"
	DestinationNewResponseTypeAppLovin                     DestinationNewResponseType = "AppLovin"
	DestinationNewResponseTypeArtsAI                       DestinationNewResponseType = "ArtsAI"
	DestinationNewResponseTypeAttentive                    DestinationNewResponseType = "Attentive"
	DestinationNewResponseTypeAudiohook                    DestinationNewResponseType = "Audiohook"
	DestinationNewResponseTypeAzureBlob                    DestinationNewResponseType = "AzureBlob"
	DestinationNewResponseTypeBasisPostback                DestinationNewResponseType = "BasisPostback"
	DestinationNewResponseTypeBeeswaxPostback              DestinationNewResponseType = "BeeswaxPostback"
	DestinationNewResponseTypeBingAds                      DestinationNewResponseType = "BingAds"
	DestinationNewResponseTypeBingAdsWeb                   DestinationNewResponseType = "BingAdsWeb"
	DestinationNewResponseTypeBraze                        DestinationNewResponseType = "Braze"
	DestinationNewResponseTypeConvertAbTestingEvent        DestinationNewResponseType = "ConvertABTestingEvent"
	DestinationNewResponseTypeCustomerio                   DestinationNewResponseType = "Customerio"
	DestinationNewResponseTypeDatabricksWarehouse          DestinationNewResponseType = "DatabricksWarehouse"
	DestinationNewResponseTypeDomoWarehouse                DestinationNewResponseType = "DomoWarehouse"
	DestinationNewResponseTypeEverflow                     DestinationNewResponseType = "Everflow"
	DestinationNewResponseTypeFacebook                     DestinationNewResponseType = "Facebook"
	DestinationNewResponseTypeFloodlightSgtm               DestinationNewResponseType = "FloodlightSGTM"
	DestinationNewResponseTypeFullContact                  DestinationNewResponseType = "FullContact"
	DestinationNewResponseTypeG4Analytics                  DestinationNewResponseType = "G4Analytics"
	DestinationNewResponseTypeGa4MeasurementProtocol       DestinationNewResponseType = "GA4MeasurementProtocol"
	DestinationNewResponseTypeGa4ServerProxy               DestinationNewResponseType = "GA4ServerProxy"
	DestinationNewResponseTypeGoogle                       DestinationNewResponseType = "Google"
	DestinationNewResponseTypeGoogleAds360                 DestinationNewResponseType = "GoogleAds360"
	DestinationNewResponseTypeGoogleAdsServerContainer     DestinationNewResponseType = "GoogleAdsServerContainer"
	DestinationNewResponseTypeGoogleBigQuery               DestinationNewResponseType = "GoogleBigQuery"
	DestinationNewResponseTypeGoogleBigQueryWarehouse      DestinationNewResponseType = "GoogleBigQueryWarehouse"
	DestinationNewResponseTypeGoogleDataManagerEventIngest DestinationNewResponseType = "GoogleDataManagerEventIngest"
	DestinationNewResponseTypeGooglePubSub                 DestinationNewResponseType = "GooglePubSub"
	DestinationNewResponseTypeGoogleStorage                DestinationNewResponseType = "GoogleStorage"
	DestinationNewResponseTypeHTTPCustomRequest            DestinationNewResponseType = "HTTPCustomRequest"
	DestinationNewResponseTypeHTTPDestination              DestinationNewResponseType = "HTTPDestination"
	DestinationNewResponseTypeHubspot                      DestinationNewResponseType = "Hubspot"
	DestinationNewResponseTypeIHeartMediaMagellan          DestinationNewResponseType = "IHeartMediaMagellan"
	DestinationNewResponseTypeImpact                       DestinationNewResponseType = "Impact"
	DestinationNewResponseTypeIterable                     DestinationNewResponseType = "Iterable"
	DestinationNewResponseTypeKlaviyo                      DestinationNewResponseType = "Klaviyo"
	DestinationNewResponseTypeLinkedInAdsCapi              DestinationNewResponseType = "LinkedInAdsCAPI"
	DestinationNewResponseTypeLiveIntent                   DestinationNewResponseType = "LiveIntent"
	DestinationNewResponseTypeLiveRampWarehouse            DestinationNewResponseType = "LiveRampWarehouse"
	DestinationNewResponseTypeMntn                         DestinationNewResponseType = "MNTN"
	DestinationNewResponseTypeMailchimp                    DestinationNewResponseType = "Mailchimp"
	DestinationNewResponseTypeMixpanel                     DestinationNewResponseType = "Mixpanel"
	DestinationNewResponseTypeNextdoorAds                  DestinationNewResponseType = "NextdoorAds"
	DestinationNewResponseTypeOpenAIAds                    DestinationNewResponseType = "OpenAIAds"
	DestinationNewResponseTypeOursSyntheticData            DestinationNewResponseType = "OursSyntheticData"
	DestinationNewResponseTypeOutbrain                     DestinationNewResponseType = "Outbrain"
	DestinationNewResponseTypePartnerize                   DestinationNewResponseType = "Partnerize"
	DestinationNewResponseTypePinterest                    DestinationNewResponseType = "Pinterest"
	DestinationNewResponseTypePlausible                    DestinationNewResponseType = "Plausible"
	DestinationNewResponseTypePodscribe                    DestinationNewResponseType = "Podscribe"
	DestinationNewResponseTypePostHog                      DestinationNewResponseType = "PostHog"
	DestinationNewResponseTypeQuantcastCapi                DestinationNewResponseType = "QuantcastCAPI"
	DestinationNewResponseTypeQuoraAds                     DestinationNewResponseType = "QuoraAds"
	DestinationNewResponseTypeReddit                       DestinationNewResponseType = "Reddit"
	DestinationNewResponseTypeRokuCapi                     DestinationNewResponseType = "RokuCAPI"
	DestinationNewResponseTypeSnapchatAdsCapi              DestinationNewResponseType = "SnapchatAdsCapi"
	DestinationNewResponseTypeSpotify                      DestinationNewResponseType = "Spotify"
	DestinationNewResponseTypeStackAdaptAPI                DestinationNewResponseType = "StackAdaptAPI"
	DestinationNewResponseTypeTaboola                      DestinationNewResponseType = "Taboola"
	DestinationNewResponseTypeTatari                       DestinationNewResponseType = "Tatari"
	DestinationNewResponseTypeTheTradeDesk                 DestinationNewResponseType = "TheTradeDesk"
	DestinationNewResponseTypeTikTok                       DestinationNewResponseType = "TikTok"
	DestinationNewResponseTypeUniversalAds                 DestinationNewResponseType = "UniversalAds"
	DestinationNewResponseTypeVwo                          DestinationNewResponseType = "VWO"
	DestinationNewResponseTypeViant                        DestinationNewResponseType = "Viant"
	DestinationNewResponseTypeViantCapi                    DestinationNewResponseType = "ViantCAPI"
	DestinationNewResponseTypeVibe                         DestinationNewResponseType = "Vibe"
	DestinationNewResponseTypeWoopra                       DestinationNewResponseType = "Woopra"
	DestinationNewResponseTypeXAds                         DestinationNewResponseType = "XAds"
	DestinationNewResponseTypeYelpCapi                     DestinationNewResponseType = "YelpCAPI"
	DestinationNewResponseTypeZendesk                      DestinationNewResponseType = "Zendesk"
	DestinationNewResponseTypeZohoCRM                      DestinationNewResponseType = "ZohoCRM"
	DestinationNewResponseTypeZoomInfo                     DestinationNewResponseType = "ZoomInfo"
)

type DestinationGetResponse struct {
	ID        string `json:"id" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	// Any of "Disabled", "Enabled".
	Status DestinationGetResponseStatus `json:"status" api:"required"`
	// Destination type. Read responses may include warehouse or cloud-storage types
	// that are not creatable through POST /rest/v1/destinations.
	//
	// Any of "AWSEventBridge", "AWSKinesis", "AWSLambda", "AWSS3", "AWSSNS",
	// "ActiveCampaignApi", "Admitad", "AdobeAnalytics", "AmazonDSP", "Amplitude",
	// "AppLovin", "ArtsAI", "Attentive", "Audiohook", "AzureBlob", "BasisPostback",
	// "BeeswaxPostback", "BingAds", "BingAdsWeb", "Braze", "ConvertABTestingEvent",
	// "Customerio", "DatabricksWarehouse", "DomoWarehouse", "Everflow", "Facebook",
	// "FloodlightSGTM", "FullContact", "G4Analytics", "GA4MeasurementProtocol",
	// "GA4ServerProxy", "Google", "GoogleAds360", "GoogleAdsServerContainer",
	// "GoogleBigQuery", "GoogleBigQueryWarehouse", "GoogleDataManagerEventIngest",
	// "GooglePubSub", "GoogleStorage", "HTTPCustomRequest", "HTTPDestination",
	// "Hubspot", "IHeartMediaMagellan", "Impact", "Iterable", "Klaviyo",
	// "LinkedInAdsCAPI", "LiveIntent", "LiveRampWarehouse", "MNTN", "Mailchimp",
	// "Mixpanel", "NextdoorAds", "OpenAIAds", "OursSyntheticData", "Outbrain",
	// "Partnerize", "Pinterest", "Plausible", "Podscribe", "PostHog", "QuantcastCAPI",
	// "QuoraAds", "Reddit", "RokuCAPI", "SnapchatAdsCapi", "Spotify", "StackAdaptAPI",
	// "Taboola", "Tatari", "TheTradeDesk", "TikTok", "UniversalAds", "VWO", "Viant",
	// "ViantCAPI", "Vibe", "Woopra", "XAds", "YelpCAPI", "Zendesk", "ZohoCRM",
	// "ZoomInfo".
	Type               DestinationGetResponseType `json:"type" api:"required"`
	HashingSalt        string                     `json:"hashingSalt" api:"nullable"`
	LimitedToSourceIDs []string                   `json:"limitedToSourceIds" api:"nullable"`
	Name               string                     `json:"name" api:"nullable"`
	Settings           any                        `json:"settings" api:"nullable"`
	UpdatedAt          string                     `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Status             respjson.Field
		Type               respjson.Field
		HashingSalt        respjson.Field
		LimitedToSourceIDs respjson.Field
		Name               respjson.Field
		Settings           respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DestinationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DestinationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationGetResponseStatus string

const (
	DestinationGetResponseStatusDisabled DestinationGetResponseStatus = "Disabled"
	DestinationGetResponseStatusEnabled  DestinationGetResponseStatus = "Enabled"
)

// Destination type. Read responses may include warehouse or cloud-storage types
// that are not creatable through POST /rest/v1/destinations.
type DestinationGetResponseType string

const (
	DestinationGetResponseTypeAwsEventBridge               DestinationGetResponseType = "AWSEventBridge"
	DestinationGetResponseTypeAwsKinesis                   DestinationGetResponseType = "AWSKinesis"
	DestinationGetResponseTypeAwsLambda                    DestinationGetResponseType = "AWSLambda"
	DestinationGetResponseTypeAwss3                        DestinationGetResponseType = "AWSS3"
	DestinationGetResponseTypeAwssns                       DestinationGetResponseType = "AWSSNS"
	DestinationGetResponseTypeActiveCampaignAPI            DestinationGetResponseType = "ActiveCampaignApi"
	DestinationGetResponseTypeAdmitad                      DestinationGetResponseType = "Admitad"
	DestinationGetResponseTypeAdobeAnalytics               DestinationGetResponseType = "AdobeAnalytics"
	DestinationGetResponseTypeAmazonDsp                    DestinationGetResponseType = "AmazonDSP"
	DestinationGetResponseTypeAmplitude                    DestinationGetResponseType = "Amplitude"
	DestinationGetResponseTypeAppLovin                     DestinationGetResponseType = "AppLovin"
	DestinationGetResponseTypeArtsAI                       DestinationGetResponseType = "ArtsAI"
	DestinationGetResponseTypeAttentive                    DestinationGetResponseType = "Attentive"
	DestinationGetResponseTypeAudiohook                    DestinationGetResponseType = "Audiohook"
	DestinationGetResponseTypeAzureBlob                    DestinationGetResponseType = "AzureBlob"
	DestinationGetResponseTypeBasisPostback                DestinationGetResponseType = "BasisPostback"
	DestinationGetResponseTypeBeeswaxPostback              DestinationGetResponseType = "BeeswaxPostback"
	DestinationGetResponseTypeBingAds                      DestinationGetResponseType = "BingAds"
	DestinationGetResponseTypeBingAdsWeb                   DestinationGetResponseType = "BingAdsWeb"
	DestinationGetResponseTypeBraze                        DestinationGetResponseType = "Braze"
	DestinationGetResponseTypeConvertAbTestingEvent        DestinationGetResponseType = "ConvertABTestingEvent"
	DestinationGetResponseTypeCustomerio                   DestinationGetResponseType = "Customerio"
	DestinationGetResponseTypeDatabricksWarehouse          DestinationGetResponseType = "DatabricksWarehouse"
	DestinationGetResponseTypeDomoWarehouse                DestinationGetResponseType = "DomoWarehouse"
	DestinationGetResponseTypeEverflow                     DestinationGetResponseType = "Everflow"
	DestinationGetResponseTypeFacebook                     DestinationGetResponseType = "Facebook"
	DestinationGetResponseTypeFloodlightSgtm               DestinationGetResponseType = "FloodlightSGTM"
	DestinationGetResponseTypeFullContact                  DestinationGetResponseType = "FullContact"
	DestinationGetResponseTypeG4Analytics                  DestinationGetResponseType = "G4Analytics"
	DestinationGetResponseTypeGa4MeasurementProtocol       DestinationGetResponseType = "GA4MeasurementProtocol"
	DestinationGetResponseTypeGa4ServerProxy               DestinationGetResponseType = "GA4ServerProxy"
	DestinationGetResponseTypeGoogle                       DestinationGetResponseType = "Google"
	DestinationGetResponseTypeGoogleAds360                 DestinationGetResponseType = "GoogleAds360"
	DestinationGetResponseTypeGoogleAdsServerContainer     DestinationGetResponseType = "GoogleAdsServerContainer"
	DestinationGetResponseTypeGoogleBigQuery               DestinationGetResponseType = "GoogleBigQuery"
	DestinationGetResponseTypeGoogleBigQueryWarehouse      DestinationGetResponseType = "GoogleBigQueryWarehouse"
	DestinationGetResponseTypeGoogleDataManagerEventIngest DestinationGetResponseType = "GoogleDataManagerEventIngest"
	DestinationGetResponseTypeGooglePubSub                 DestinationGetResponseType = "GooglePubSub"
	DestinationGetResponseTypeGoogleStorage                DestinationGetResponseType = "GoogleStorage"
	DestinationGetResponseTypeHTTPCustomRequest            DestinationGetResponseType = "HTTPCustomRequest"
	DestinationGetResponseTypeHTTPDestination              DestinationGetResponseType = "HTTPDestination"
	DestinationGetResponseTypeHubspot                      DestinationGetResponseType = "Hubspot"
	DestinationGetResponseTypeIHeartMediaMagellan          DestinationGetResponseType = "IHeartMediaMagellan"
	DestinationGetResponseTypeImpact                       DestinationGetResponseType = "Impact"
	DestinationGetResponseTypeIterable                     DestinationGetResponseType = "Iterable"
	DestinationGetResponseTypeKlaviyo                      DestinationGetResponseType = "Klaviyo"
	DestinationGetResponseTypeLinkedInAdsCapi              DestinationGetResponseType = "LinkedInAdsCAPI"
	DestinationGetResponseTypeLiveIntent                   DestinationGetResponseType = "LiveIntent"
	DestinationGetResponseTypeLiveRampWarehouse            DestinationGetResponseType = "LiveRampWarehouse"
	DestinationGetResponseTypeMntn                         DestinationGetResponseType = "MNTN"
	DestinationGetResponseTypeMailchimp                    DestinationGetResponseType = "Mailchimp"
	DestinationGetResponseTypeMixpanel                     DestinationGetResponseType = "Mixpanel"
	DestinationGetResponseTypeNextdoorAds                  DestinationGetResponseType = "NextdoorAds"
	DestinationGetResponseTypeOpenAIAds                    DestinationGetResponseType = "OpenAIAds"
	DestinationGetResponseTypeOursSyntheticData            DestinationGetResponseType = "OursSyntheticData"
	DestinationGetResponseTypeOutbrain                     DestinationGetResponseType = "Outbrain"
	DestinationGetResponseTypePartnerize                   DestinationGetResponseType = "Partnerize"
	DestinationGetResponseTypePinterest                    DestinationGetResponseType = "Pinterest"
	DestinationGetResponseTypePlausible                    DestinationGetResponseType = "Plausible"
	DestinationGetResponseTypePodscribe                    DestinationGetResponseType = "Podscribe"
	DestinationGetResponseTypePostHog                      DestinationGetResponseType = "PostHog"
	DestinationGetResponseTypeQuantcastCapi                DestinationGetResponseType = "QuantcastCAPI"
	DestinationGetResponseTypeQuoraAds                     DestinationGetResponseType = "QuoraAds"
	DestinationGetResponseTypeReddit                       DestinationGetResponseType = "Reddit"
	DestinationGetResponseTypeRokuCapi                     DestinationGetResponseType = "RokuCAPI"
	DestinationGetResponseTypeSnapchatAdsCapi              DestinationGetResponseType = "SnapchatAdsCapi"
	DestinationGetResponseTypeSpotify                      DestinationGetResponseType = "Spotify"
	DestinationGetResponseTypeStackAdaptAPI                DestinationGetResponseType = "StackAdaptAPI"
	DestinationGetResponseTypeTaboola                      DestinationGetResponseType = "Taboola"
	DestinationGetResponseTypeTatari                       DestinationGetResponseType = "Tatari"
	DestinationGetResponseTypeTheTradeDesk                 DestinationGetResponseType = "TheTradeDesk"
	DestinationGetResponseTypeTikTok                       DestinationGetResponseType = "TikTok"
	DestinationGetResponseTypeUniversalAds                 DestinationGetResponseType = "UniversalAds"
	DestinationGetResponseTypeVwo                          DestinationGetResponseType = "VWO"
	DestinationGetResponseTypeViant                        DestinationGetResponseType = "Viant"
	DestinationGetResponseTypeViantCapi                    DestinationGetResponseType = "ViantCAPI"
	DestinationGetResponseTypeVibe                         DestinationGetResponseType = "Vibe"
	DestinationGetResponseTypeWoopra                       DestinationGetResponseType = "Woopra"
	DestinationGetResponseTypeXAds                         DestinationGetResponseType = "XAds"
	DestinationGetResponseTypeYelpCapi                     DestinationGetResponseType = "YelpCAPI"
	DestinationGetResponseTypeZendesk                      DestinationGetResponseType = "Zendesk"
	DestinationGetResponseTypeZohoCRM                      DestinationGetResponseType = "ZohoCRM"
	DestinationGetResponseTypeZoomInfo                     DestinationGetResponseType = "ZoomInfo"
)

type DestinationUpdateResponse struct {
	ID        string `json:"id" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	// Any of "Disabled", "Enabled".
	Status DestinationUpdateResponseStatus `json:"status" api:"required"`
	// Destination type. Read responses may include warehouse or cloud-storage types
	// that are not creatable through POST /rest/v1/destinations.
	//
	// Any of "AWSEventBridge", "AWSKinesis", "AWSLambda", "AWSS3", "AWSSNS",
	// "ActiveCampaignApi", "Admitad", "AdobeAnalytics", "AmazonDSP", "Amplitude",
	// "AppLovin", "ArtsAI", "Attentive", "Audiohook", "AzureBlob", "BasisPostback",
	// "BeeswaxPostback", "BingAds", "BingAdsWeb", "Braze", "ConvertABTestingEvent",
	// "Customerio", "DatabricksWarehouse", "DomoWarehouse", "Everflow", "Facebook",
	// "FloodlightSGTM", "FullContact", "G4Analytics", "GA4MeasurementProtocol",
	// "GA4ServerProxy", "Google", "GoogleAds360", "GoogleAdsServerContainer",
	// "GoogleBigQuery", "GoogleBigQueryWarehouse", "GoogleDataManagerEventIngest",
	// "GooglePubSub", "GoogleStorage", "HTTPCustomRequest", "HTTPDestination",
	// "Hubspot", "IHeartMediaMagellan", "Impact", "Iterable", "Klaviyo",
	// "LinkedInAdsCAPI", "LiveIntent", "LiveRampWarehouse", "MNTN", "Mailchimp",
	// "Mixpanel", "NextdoorAds", "OpenAIAds", "OursSyntheticData", "Outbrain",
	// "Partnerize", "Pinterest", "Plausible", "Podscribe", "PostHog", "QuantcastCAPI",
	// "QuoraAds", "Reddit", "RokuCAPI", "SnapchatAdsCapi", "Spotify", "StackAdaptAPI",
	// "Taboola", "Tatari", "TheTradeDesk", "TikTok", "UniversalAds", "VWO", "Viant",
	// "ViantCAPI", "Vibe", "Woopra", "XAds", "YelpCAPI", "Zendesk", "ZohoCRM",
	// "ZoomInfo".
	Type               DestinationUpdateResponseType `json:"type" api:"required"`
	HashingSalt        string                        `json:"hashingSalt" api:"nullable"`
	LimitedToSourceIDs []string                      `json:"limitedToSourceIds" api:"nullable"`
	Name               string                        `json:"name" api:"nullable"`
	Settings           any                           `json:"settings" api:"nullable"`
	UpdatedAt          string                        `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Status             respjson.Field
		Type               respjson.Field
		HashingSalt        respjson.Field
		LimitedToSourceIDs respjson.Field
		Name               respjson.Field
		Settings           respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DestinationUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *DestinationUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationUpdateResponseStatus string

const (
	DestinationUpdateResponseStatusDisabled DestinationUpdateResponseStatus = "Disabled"
	DestinationUpdateResponseStatusEnabled  DestinationUpdateResponseStatus = "Enabled"
)

// Destination type. Read responses may include warehouse or cloud-storage types
// that are not creatable through POST /rest/v1/destinations.
type DestinationUpdateResponseType string

const (
	DestinationUpdateResponseTypeAwsEventBridge               DestinationUpdateResponseType = "AWSEventBridge"
	DestinationUpdateResponseTypeAwsKinesis                   DestinationUpdateResponseType = "AWSKinesis"
	DestinationUpdateResponseTypeAwsLambda                    DestinationUpdateResponseType = "AWSLambda"
	DestinationUpdateResponseTypeAwss3                        DestinationUpdateResponseType = "AWSS3"
	DestinationUpdateResponseTypeAwssns                       DestinationUpdateResponseType = "AWSSNS"
	DestinationUpdateResponseTypeActiveCampaignAPI            DestinationUpdateResponseType = "ActiveCampaignApi"
	DestinationUpdateResponseTypeAdmitad                      DestinationUpdateResponseType = "Admitad"
	DestinationUpdateResponseTypeAdobeAnalytics               DestinationUpdateResponseType = "AdobeAnalytics"
	DestinationUpdateResponseTypeAmazonDsp                    DestinationUpdateResponseType = "AmazonDSP"
	DestinationUpdateResponseTypeAmplitude                    DestinationUpdateResponseType = "Amplitude"
	DestinationUpdateResponseTypeAppLovin                     DestinationUpdateResponseType = "AppLovin"
	DestinationUpdateResponseTypeArtsAI                       DestinationUpdateResponseType = "ArtsAI"
	DestinationUpdateResponseTypeAttentive                    DestinationUpdateResponseType = "Attentive"
	DestinationUpdateResponseTypeAudiohook                    DestinationUpdateResponseType = "Audiohook"
	DestinationUpdateResponseTypeAzureBlob                    DestinationUpdateResponseType = "AzureBlob"
	DestinationUpdateResponseTypeBasisPostback                DestinationUpdateResponseType = "BasisPostback"
	DestinationUpdateResponseTypeBeeswaxPostback              DestinationUpdateResponseType = "BeeswaxPostback"
	DestinationUpdateResponseTypeBingAds                      DestinationUpdateResponseType = "BingAds"
	DestinationUpdateResponseTypeBingAdsWeb                   DestinationUpdateResponseType = "BingAdsWeb"
	DestinationUpdateResponseTypeBraze                        DestinationUpdateResponseType = "Braze"
	DestinationUpdateResponseTypeConvertAbTestingEvent        DestinationUpdateResponseType = "ConvertABTestingEvent"
	DestinationUpdateResponseTypeCustomerio                   DestinationUpdateResponseType = "Customerio"
	DestinationUpdateResponseTypeDatabricksWarehouse          DestinationUpdateResponseType = "DatabricksWarehouse"
	DestinationUpdateResponseTypeDomoWarehouse                DestinationUpdateResponseType = "DomoWarehouse"
	DestinationUpdateResponseTypeEverflow                     DestinationUpdateResponseType = "Everflow"
	DestinationUpdateResponseTypeFacebook                     DestinationUpdateResponseType = "Facebook"
	DestinationUpdateResponseTypeFloodlightSgtm               DestinationUpdateResponseType = "FloodlightSGTM"
	DestinationUpdateResponseTypeFullContact                  DestinationUpdateResponseType = "FullContact"
	DestinationUpdateResponseTypeG4Analytics                  DestinationUpdateResponseType = "G4Analytics"
	DestinationUpdateResponseTypeGa4MeasurementProtocol       DestinationUpdateResponseType = "GA4MeasurementProtocol"
	DestinationUpdateResponseTypeGa4ServerProxy               DestinationUpdateResponseType = "GA4ServerProxy"
	DestinationUpdateResponseTypeGoogle                       DestinationUpdateResponseType = "Google"
	DestinationUpdateResponseTypeGoogleAds360                 DestinationUpdateResponseType = "GoogleAds360"
	DestinationUpdateResponseTypeGoogleAdsServerContainer     DestinationUpdateResponseType = "GoogleAdsServerContainer"
	DestinationUpdateResponseTypeGoogleBigQuery               DestinationUpdateResponseType = "GoogleBigQuery"
	DestinationUpdateResponseTypeGoogleBigQueryWarehouse      DestinationUpdateResponseType = "GoogleBigQueryWarehouse"
	DestinationUpdateResponseTypeGoogleDataManagerEventIngest DestinationUpdateResponseType = "GoogleDataManagerEventIngest"
	DestinationUpdateResponseTypeGooglePubSub                 DestinationUpdateResponseType = "GooglePubSub"
	DestinationUpdateResponseTypeGoogleStorage                DestinationUpdateResponseType = "GoogleStorage"
	DestinationUpdateResponseTypeHTTPCustomRequest            DestinationUpdateResponseType = "HTTPCustomRequest"
	DestinationUpdateResponseTypeHTTPDestination              DestinationUpdateResponseType = "HTTPDestination"
	DestinationUpdateResponseTypeHubspot                      DestinationUpdateResponseType = "Hubspot"
	DestinationUpdateResponseTypeIHeartMediaMagellan          DestinationUpdateResponseType = "IHeartMediaMagellan"
	DestinationUpdateResponseTypeImpact                       DestinationUpdateResponseType = "Impact"
	DestinationUpdateResponseTypeIterable                     DestinationUpdateResponseType = "Iterable"
	DestinationUpdateResponseTypeKlaviyo                      DestinationUpdateResponseType = "Klaviyo"
	DestinationUpdateResponseTypeLinkedInAdsCapi              DestinationUpdateResponseType = "LinkedInAdsCAPI"
	DestinationUpdateResponseTypeLiveIntent                   DestinationUpdateResponseType = "LiveIntent"
	DestinationUpdateResponseTypeLiveRampWarehouse            DestinationUpdateResponseType = "LiveRampWarehouse"
	DestinationUpdateResponseTypeMntn                         DestinationUpdateResponseType = "MNTN"
	DestinationUpdateResponseTypeMailchimp                    DestinationUpdateResponseType = "Mailchimp"
	DestinationUpdateResponseTypeMixpanel                     DestinationUpdateResponseType = "Mixpanel"
	DestinationUpdateResponseTypeNextdoorAds                  DestinationUpdateResponseType = "NextdoorAds"
	DestinationUpdateResponseTypeOpenAIAds                    DestinationUpdateResponseType = "OpenAIAds"
	DestinationUpdateResponseTypeOursSyntheticData            DestinationUpdateResponseType = "OursSyntheticData"
	DestinationUpdateResponseTypeOutbrain                     DestinationUpdateResponseType = "Outbrain"
	DestinationUpdateResponseTypePartnerize                   DestinationUpdateResponseType = "Partnerize"
	DestinationUpdateResponseTypePinterest                    DestinationUpdateResponseType = "Pinterest"
	DestinationUpdateResponseTypePlausible                    DestinationUpdateResponseType = "Plausible"
	DestinationUpdateResponseTypePodscribe                    DestinationUpdateResponseType = "Podscribe"
	DestinationUpdateResponseTypePostHog                      DestinationUpdateResponseType = "PostHog"
	DestinationUpdateResponseTypeQuantcastCapi                DestinationUpdateResponseType = "QuantcastCAPI"
	DestinationUpdateResponseTypeQuoraAds                     DestinationUpdateResponseType = "QuoraAds"
	DestinationUpdateResponseTypeReddit                       DestinationUpdateResponseType = "Reddit"
	DestinationUpdateResponseTypeRokuCapi                     DestinationUpdateResponseType = "RokuCAPI"
	DestinationUpdateResponseTypeSnapchatAdsCapi              DestinationUpdateResponseType = "SnapchatAdsCapi"
	DestinationUpdateResponseTypeSpotify                      DestinationUpdateResponseType = "Spotify"
	DestinationUpdateResponseTypeStackAdaptAPI                DestinationUpdateResponseType = "StackAdaptAPI"
	DestinationUpdateResponseTypeTaboola                      DestinationUpdateResponseType = "Taboola"
	DestinationUpdateResponseTypeTatari                       DestinationUpdateResponseType = "Tatari"
	DestinationUpdateResponseTypeTheTradeDesk                 DestinationUpdateResponseType = "TheTradeDesk"
	DestinationUpdateResponseTypeTikTok                       DestinationUpdateResponseType = "TikTok"
	DestinationUpdateResponseTypeUniversalAds                 DestinationUpdateResponseType = "UniversalAds"
	DestinationUpdateResponseTypeVwo                          DestinationUpdateResponseType = "VWO"
	DestinationUpdateResponseTypeViant                        DestinationUpdateResponseType = "Viant"
	DestinationUpdateResponseTypeViantCapi                    DestinationUpdateResponseType = "ViantCAPI"
	DestinationUpdateResponseTypeVibe                         DestinationUpdateResponseType = "Vibe"
	DestinationUpdateResponseTypeWoopra                       DestinationUpdateResponseType = "Woopra"
	DestinationUpdateResponseTypeXAds                         DestinationUpdateResponseType = "XAds"
	DestinationUpdateResponseTypeYelpCapi                     DestinationUpdateResponseType = "YelpCAPI"
	DestinationUpdateResponseTypeZendesk                      DestinationUpdateResponseType = "Zendesk"
	DestinationUpdateResponseTypeZohoCRM                      DestinationUpdateResponseType = "ZohoCRM"
	DestinationUpdateResponseTypeZoomInfo                     DestinationUpdateResponseType = "ZoomInfo"
)

type DestinationTypesResponse struct {
	Entities []DestinationTypesResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DestinationTypesResponse) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypesResponseEntity struct {
	// Any of "Audiohook", "BasisPostback", "Outbrain", "OursSyntheticData",
	// "FullContact", "ZoomInfo", "TheTradeDesk", "Braze", "LiveIntent",
	// "ConvertABTestingEvent", "Customerio", "BingAds", "BingAdsWeb",
	// "HTTPDestination", "Woopra", "HTTPCustomRequest", "Google",
	// "GoogleAdsServerContainer", "G4Analytics", "GA4ServerProxy",
	// "GA4MeasurementProtocol", "GoogleAds360", "Facebook", "Mixpanel", "Amplitude",
	// "TikTok", "Reddit", "Podscribe", "Pinterest", "Mailchimp", "AWSKinesis",
	// "AWSLambda", "AWSSNS", "GooglePubSub", "LinkedInAdsCAPI", "ActiveCampaignApi",
	// "StackAdaptAPI", "Hubspot", "Klaviyo", "XAds", "QuoraAds", "SnapchatAdsCapi",
	// "Partnerize", "NextdoorAds", "Tatari", "Viant", "ViantCAPI", "Impact",
	// "Spotify", "Taboola", "AmazonDSP", "AppLovin", "IHeartMediaMagellan", "Vibe",
	// "GoogleDataManagerEventIngest", "Zendesk", "Iterable", "ArtsAI",
	// "QuantcastCAPI", "FloodlightSGTM", "VWO", "Attentive", "Admitad", "Plausible",
	// "PostHog", "RokuCAPI", "Everflow", "BeeswaxPostback", "AdobeAnalytics",
	// "UniversalAds", "OpenAIAds", "YelpCAPI", "MNTN", "ZohoCRM".
	ID           string                                       `json:"id" api:"required"`
	Capabilities DestinationTypesResponseEntityCapabilities   `json:"capabilities" api:"required"`
	Label        string                                       `json:"label" api:"required"`
	Settings     []DestinationTypesResponseEntitySettingUnion `json:"settings" api:"required"`
	// Any of "deprecated", "ga".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Capabilities respjson.Field
		Label        respjson.Field
		Settings     respjson.Field
		Status       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DestinationTypesResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypesResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypesResponseEntityCapabilities struct {
	ListsAccounts         bool `json:"listsAccounts" api:"required"`
	OAuth                 bool `json:"oauth" api:"required"`
	SupportsRenamedEvents bool `json:"supportsRenamedEvents" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ListsAccounts         respjson.Field
		OAuth                 respjson.Field
		SupportsRenamedEvents respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DestinationTypesResponseEntityCapabilities) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypesResponseEntityCapabilities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DestinationTypesResponseEntitySettingUnion contains all possible properties and
// values from [DestinationTypesResponseEntitySettingObject],
// [DestinationTypesResponseEntitySettingObject2],
// [DestinationTypesResponseEntitySettingObject3],
// [DestinationTypesResponseEntitySettingObject4],
// [DestinationTypesResponseEntitySettingObject5].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type DestinationTypesResponseEntitySettingUnion struct {
	Key   string `json:"key"`
	Label string `json:"label"`
	Type  string `json:"type"`
	// This field is from variant [DestinationTypesResponseEntitySettingObject2].
	Options []DestinationTypesResponseEntitySettingObject2Option `json:"options"`
	// This field is a union of [string], [bool]
	DefaultValue DestinationTypesResponseEntitySettingUnionDefaultValue `json:"defaultValue"`
	Required     bool                                                   `json:"required"`
	Sublabel     string                                                 `json:"sublabel"`
	// This field is from variant [DestinationTypesResponseEntitySettingObject5].
	Placeholder string `json:"placeholder"`
	JSON        struct {
		Key          respjson.Field
		Label        respjson.Field
		Type         respjson.Field
		Options      respjson.Field
		DefaultValue respjson.Field
		Required     respjson.Field
		Sublabel     respjson.Field
		Placeholder  respjson.Field
		raw          string
	} `json:"-"`
}

func (u DestinationTypesResponseEntitySettingUnion) AsDestinationTypesResponseEntitySettingObject() (v DestinationTypesResponseEntitySettingObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DestinationTypesResponseEntitySettingUnion) AsDestinationTypesResponseEntitySettingObject2() (v DestinationTypesResponseEntitySettingObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DestinationTypesResponseEntitySettingUnion) AsDestinationTypesResponseEntitySettingObject3() (v DestinationTypesResponseEntitySettingObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DestinationTypesResponseEntitySettingUnion) AsDestinationTypesResponseEntitySettingObject4() (v DestinationTypesResponseEntitySettingObject4) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DestinationTypesResponseEntitySettingUnion) AsDestinationTypesResponseEntitySettingObject5() (v DestinationTypesResponseEntitySettingObject5) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DestinationTypesResponseEntitySettingUnion) RawJSON() string { return u.JSON.raw }

func (r *DestinationTypesResponseEntitySettingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DestinationTypesResponseEntitySettingUnionDefaultValue is an implicit subunion
// of [DestinationTypesResponseEntitySettingUnion].
// DestinationTypesResponseEntitySettingUnionDefaultValue provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [DestinationTypesResponseEntitySettingUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfBool]
type DestinationTypesResponseEntitySettingUnionDefaultValue struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (r *DestinationTypesResponseEntitySettingUnionDefaultValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypesResponseEntitySettingObject struct {
	Key string `json:"key" api:"required"`
	// Informational display message only. Do not send this key in POST or PATCH
	// settings.
	Label string `json:"label" api:"required"`
	// Any of "Alert".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Label       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DestinationTypesResponseEntitySettingObject) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypesResponseEntitySettingObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypesResponseEntitySettingObject2 struct {
	Key     string                                               `json:"key" api:"required"`
	Label   string                                               `json:"label" api:"required"`
	Options []DestinationTypesResponseEntitySettingObject2Option `json:"options" api:"required"`
	// Any of "Select".
	Type         string `json:"type" api:"required"`
	DefaultValue string `json:"defaultValue" api:"nullable"`
	Required     bool   `json:"required" api:"nullable"`
	Sublabel     string `json:"sublabel" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key          respjson.Field
		Label        respjson.Field
		Options      respjson.Field
		Type         respjson.Field
		DefaultValue respjson.Field
		Required     respjson.Field
		Sublabel     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DestinationTypesResponseEntitySettingObject2) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypesResponseEntitySettingObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypesResponseEntitySettingObject2Option struct {
	Label string `json:"label" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DestinationTypesResponseEntitySettingObject2Option) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypesResponseEntitySettingObject2Option) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypesResponseEntitySettingObject3 struct {
	Key   string `json:"key" api:"required"`
	Label string `json:"label" api:"required"`
	// Any of "Switch".
	Type         string `json:"type" api:"required"`
	DefaultValue bool   `json:"defaultValue" api:"nullable"`
	Required     bool   `json:"required" api:"nullable"`
	Sublabel     string `json:"sublabel" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key          respjson.Field
		Label        respjson.Field
		Type         respjson.Field
		DefaultValue respjson.Field
		Required     respjson.Field
		Sublabel     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DestinationTypesResponseEntitySettingObject3) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypesResponseEntitySettingObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypesResponseEntitySettingObject4 struct {
	Key   string `json:"key" api:"required"`
	Label string `json:"label" api:"required"`
	// Any of "GenericOauth".
	Type     string `json:"type" api:"required"`
	Sublabel string `json:"sublabel" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Label       respjson.Field
		Type        respjson.Field
		Sublabel    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DestinationTypesResponseEntitySettingObject4) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypesResponseEntitySettingObject4) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypesResponseEntitySettingObject5 struct {
	Key         string `json:"key" api:"required"`
	Label       string `json:"label" api:"required"`
	Placeholder string `json:"placeholder" api:"required"`
	// Any of "Text", "Secret".
	Type     DestinationTypesResponseEntitySettingObject5Type `json:"type" api:"required"`
	Required bool                                             `json:"required" api:"nullable"`
	Sublabel string                                           `json:"sublabel" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Label       respjson.Field
		Placeholder respjson.Field
		Type        respjson.Field
		Required    respjson.Field
		Sublabel    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DestinationTypesResponseEntitySettingObject5) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypesResponseEntitySettingObject5) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypesResponseEntitySettingObject5Type string

const (
	DestinationTypesResponseEntitySettingObject5TypeText   DestinationTypesResponseEntitySettingObject5Type = "Text"
	DestinationTypesResponseEntitySettingObject5TypeSecret DestinationTypesResponseEntitySettingObject5Type = "Secret"
)

type DestinationListParams struct {
	// Maximum number of items to return. Defaults to 25; values below 1 are clamped to
	// 1 and values above 100 are clamped to 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Opaque pagination cursor from pagination.nextCursor in the previous response. Do
	// not decode or modify it. Malformed cursors return 400 Bad Request.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Filter destinations by status.
	//
	// Any of "Disabled", "Enabled".
	Status DestinationListParamsStatus `query:"status,omitzero" json:"-"`
	// Filter destinations by destination type.
	//
	// Any of "AWSEventBridge", "AWSKinesis", "AWSLambda", "AWSS3", "AWSSNS",
	// "ActiveCampaignApi", "Admitad", "AdobeAnalytics", "AmazonDSP", "Amplitude",
	// "AppLovin", "ArtsAI", "Attentive", "Audiohook", "AzureBlob", "BasisPostback",
	// "BeeswaxPostback", "BingAds", "BingAdsWeb", "Braze", "ConvertABTestingEvent",
	// "Customerio", "DatabricksWarehouse", "DomoWarehouse", "Everflow", "Facebook",
	// "FloodlightSGTM", "FullContact", "G4Analytics", "GA4MeasurementProtocol",
	// "GA4ServerProxy", "Google", "GoogleAds360", "GoogleAdsServerContainer",
	// "GoogleBigQuery", "GoogleBigQueryWarehouse", "GoogleDataManagerEventIngest",
	// "GooglePubSub", "GoogleStorage", "HTTPCustomRequest", "HTTPDestination",
	// "Hubspot", "IHeartMediaMagellan", "Impact", "Iterable", "Klaviyo",
	// "LinkedInAdsCAPI", "LiveIntent", "LiveRampWarehouse", "MNTN", "Mailchimp",
	// "Mixpanel", "NextdoorAds", "OpenAIAds", "OursSyntheticData", "Outbrain",
	// "Partnerize", "Pinterest", "Plausible", "Podscribe", "PostHog", "QuantcastCAPI",
	// "QuoraAds", "Reddit", "RokuCAPI", "SnapchatAdsCapi", "Spotify", "StackAdaptAPI",
	// "Taboola", "Tatari", "TheTradeDesk", "TikTok", "UniversalAds", "VWO", "Viant",
	// "ViantCAPI", "Vibe", "Woopra", "XAds", "YelpCAPI", "Zendesk", "ZohoCRM",
	// "ZoomInfo".
	Type DestinationListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DestinationListParams]'s query parameters as `url.Values`.
func (r DestinationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter destinations by status.
type DestinationListParamsStatus string

const (
	DestinationListParamsStatusDisabled DestinationListParamsStatus = "Disabled"
	DestinationListParamsStatusEnabled  DestinationListParamsStatus = "Enabled"
)

// Filter destinations by destination type.
type DestinationListParamsType string

const (
	DestinationListParamsTypeAwsEventBridge               DestinationListParamsType = "AWSEventBridge"
	DestinationListParamsTypeAwsKinesis                   DestinationListParamsType = "AWSKinesis"
	DestinationListParamsTypeAwsLambda                    DestinationListParamsType = "AWSLambda"
	DestinationListParamsTypeAwss3                        DestinationListParamsType = "AWSS3"
	DestinationListParamsTypeAwssns                       DestinationListParamsType = "AWSSNS"
	DestinationListParamsTypeActiveCampaignAPI            DestinationListParamsType = "ActiveCampaignApi"
	DestinationListParamsTypeAdmitad                      DestinationListParamsType = "Admitad"
	DestinationListParamsTypeAdobeAnalytics               DestinationListParamsType = "AdobeAnalytics"
	DestinationListParamsTypeAmazonDsp                    DestinationListParamsType = "AmazonDSP"
	DestinationListParamsTypeAmplitude                    DestinationListParamsType = "Amplitude"
	DestinationListParamsTypeAppLovin                     DestinationListParamsType = "AppLovin"
	DestinationListParamsTypeArtsAI                       DestinationListParamsType = "ArtsAI"
	DestinationListParamsTypeAttentive                    DestinationListParamsType = "Attentive"
	DestinationListParamsTypeAudiohook                    DestinationListParamsType = "Audiohook"
	DestinationListParamsTypeAzureBlob                    DestinationListParamsType = "AzureBlob"
	DestinationListParamsTypeBasisPostback                DestinationListParamsType = "BasisPostback"
	DestinationListParamsTypeBeeswaxPostback              DestinationListParamsType = "BeeswaxPostback"
	DestinationListParamsTypeBingAds                      DestinationListParamsType = "BingAds"
	DestinationListParamsTypeBingAdsWeb                   DestinationListParamsType = "BingAdsWeb"
	DestinationListParamsTypeBraze                        DestinationListParamsType = "Braze"
	DestinationListParamsTypeConvertAbTestingEvent        DestinationListParamsType = "ConvertABTestingEvent"
	DestinationListParamsTypeCustomerio                   DestinationListParamsType = "Customerio"
	DestinationListParamsTypeDatabricksWarehouse          DestinationListParamsType = "DatabricksWarehouse"
	DestinationListParamsTypeDomoWarehouse                DestinationListParamsType = "DomoWarehouse"
	DestinationListParamsTypeEverflow                     DestinationListParamsType = "Everflow"
	DestinationListParamsTypeFacebook                     DestinationListParamsType = "Facebook"
	DestinationListParamsTypeFloodlightSgtm               DestinationListParamsType = "FloodlightSGTM"
	DestinationListParamsTypeFullContact                  DestinationListParamsType = "FullContact"
	DestinationListParamsTypeG4Analytics                  DestinationListParamsType = "G4Analytics"
	DestinationListParamsTypeGa4MeasurementProtocol       DestinationListParamsType = "GA4MeasurementProtocol"
	DestinationListParamsTypeGa4ServerProxy               DestinationListParamsType = "GA4ServerProxy"
	DestinationListParamsTypeGoogle                       DestinationListParamsType = "Google"
	DestinationListParamsTypeGoogleAds360                 DestinationListParamsType = "GoogleAds360"
	DestinationListParamsTypeGoogleAdsServerContainer     DestinationListParamsType = "GoogleAdsServerContainer"
	DestinationListParamsTypeGoogleBigQuery               DestinationListParamsType = "GoogleBigQuery"
	DestinationListParamsTypeGoogleBigQueryWarehouse      DestinationListParamsType = "GoogleBigQueryWarehouse"
	DestinationListParamsTypeGoogleDataManagerEventIngest DestinationListParamsType = "GoogleDataManagerEventIngest"
	DestinationListParamsTypeGooglePubSub                 DestinationListParamsType = "GooglePubSub"
	DestinationListParamsTypeGoogleStorage                DestinationListParamsType = "GoogleStorage"
	DestinationListParamsTypeHTTPCustomRequest            DestinationListParamsType = "HTTPCustomRequest"
	DestinationListParamsTypeHTTPDestination              DestinationListParamsType = "HTTPDestination"
	DestinationListParamsTypeHubspot                      DestinationListParamsType = "Hubspot"
	DestinationListParamsTypeIHeartMediaMagellan          DestinationListParamsType = "IHeartMediaMagellan"
	DestinationListParamsTypeImpact                       DestinationListParamsType = "Impact"
	DestinationListParamsTypeIterable                     DestinationListParamsType = "Iterable"
	DestinationListParamsTypeKlaviyo                      DestinationListParamsType = "Klaviyo"
	DestinationListParamsTypeLinkedInAdsCapi              DestinationListParamsType = "LinkedInAdsCAPI"
	DestinationListParamsTypeLiveIntent                   DestinationListParamsType = "LiveIntent"
	DestinationListParamsTypeLiveRampWarehouse            DestinationListParamsType = "LiveRampWarehouse"
	DestinationListParamsTypeMntn                         DestinationListParamsType = "MNTN"
	DestinationListParamsTypeMailchimp                    DestinationListParamsType = "Mailchimp"
	DestinationListParamsTypeMixpanel                     DestinationListParamsType = "Mixpanel"
	DestinationListParamsTypeNextdoorAds                  DestinationListParamsType = "NextdoorAds"
	DestinationListParamsTypeOpenAIAds                    DestinationListParamsType = "OpenAIAds"
	DestinationListParamsTypeOursSyntheticData            DestinationListParamsType = "OursSyntheticData"
	DestinationListParamsTypeOutbrain                     DestinationListParamsType = "Outbrain"
	DestinationListParamsTypePartnerize                   DestinationListParamsType = "Partnerize"
	DestinationListParamsTypePinterest                    DestinationListParamsType = "Pinterest"
	DestinationListParamsTypePlausible                    DestinationListParamsType = "Plausible"
	DestinationListParamsTypePodscribe                    DestinationListParamsType = "Podscribe"
	DestinationListParamsTypePostHog                      DestinationListParamsType = "PostHog"
	DestinationListParamsTypeQuantcastCapi                DestinationListParamsType = "QuantcastCAPI"
	DestinationListParamsTypeQuoraAds                     DestinationListParamsType = "QuoraAds"
	DestinationListParamsTypeReddit                       DestinationListParamsType = "Reddit"
	DestinationListParamsTypeRokuCapi                     DestinationListParamsType = "RokuCAPI"
	DestinationListParamsTypeSnapchatAdsCapi              DestinationListParamsType = "SnapchatAdsCapi"
	DestinationListParamsTypeSpotify                      DestinationListParamsType = "Spotify"
	DestinationListParamsTypeStackAdaptAPI                DestinationListParamsType = "StackAdaptAPI"
	DestinationListParamsTypeTaboola                      DestinationListParamsType = "Taboola"
	DestinationListParamsTypeTatari                       DestinationListParamsType = "Tatari"
	DestinationListParamsTypeTheTradeDesk                 DestinationListParamsType = "TheTradeDesk"
	DestinationListParamsTypeTikTok                       DestinationListParamsType = "TikTok"
	DestinationListParamsTypeUniversalAds                 DestinationListParamsType = "UniversalAds"
	DestinationListParamsTypeVwo                          DestinationListParamsType = "VWO"
	DestinationListParamsTypeViant                        DestinationListParamsType = "Viant"
	DestinationListParamsTypeViantCapi                    DestinationListParamsType = "ViantCAPI"
	DestinationListParamsTypeVibe                         DestinationListParamsType = "Vibe"
	DestinationListParamsTypeWoopra                       DestinationListParamsType = "Woopra"
	DestinationListParamsTypeXAds                         DestinationListParamsType = "XAds"
	DestinationListParamsTypeYelpCapi                     DestinationListParamsType = "YelpCAPI"
	DestinationListParamsTypeZendesk                      DestinationListParamsType = "Zendesk"
	DestinationListParamsTypeZohoCRM                      DestinationListParamsType = "ZohoCRM"
	DestinationListParamsTypeZoomInfo                     DestinationListParamsType = "ZoomInfo"
)

type DestinationNewParams struct {
	// Event-dispatch destination type to create. Warehouse and cloud-storage
	// destination types may appear on read responses but are not creatable through
	// POST.
	//
	// Any of "Audiohook", "BasisPostback", "Outbrain", "OursSyntheticData",
	// "FullContact", "ZoomInfo", "TheTradeDesk", "Braze", "LiveIntent",
	// "ConvertABTestingEvent", "Customerio", "BingAds", "BingAdsWeb",
	// "HTTPDestination", "Woopra", "HTTPCustomRequest", "Google",
	// "GoogleAdsServerContainer", "G4Analytics", "GA4ServerProxy",
	// "GA4MeasurementProtocol", "GoogleAds360", "Facebook", "Mixpanel", "Amplitude",
	// "TikTok", "Reddit", "Podscribe", "Pinterest", "Mailchimp", "AWSKinesis",
	// "AWSLambda", "AWSSNS", "GooglePubSub", "LinkedInAdsCAPI", "ActiveCampaignApi",
	// "StackAdaptAPI", "Hubspot", "Klaviyo", "XAds", "QuoraAds", "SnapchatAdsCapi",
	// "Partnerize", "NextdoorAds", "Tatari", "Viant", "ViantCAPI", "Impact",
	// "Spotify", "Taboola", "AmazonDSP", "AppLovin", "IHeartMediaMagellan", "Vibe",
	// "GoogleDataManagerEventIngest", "Zendesk", "Iterable", "ArtsAI",
	// "QuantcastCAPI", "FloodlightSGTM", "VWO", "Attentive", "Admitad", "Plausible",
	// "PostHog", "RokuCAPI", "Everflow", "BeeswaxPostback", "AdobeAnalytics",
	// "UniversalAds", "OpenAIAds", "YelpCAPI", "MNTN", "ZohoCRM".
	Type DestinationNewParamsType `json:"type,omitzero" api:"required"`
	Name param.Opt[string]        `json:"name,omitzero"`
	// Per-type configuration keys and values. Call GET /rest/v1/destinations/types to
	// get the valid keys for your destination type.
	Settings any `json:"settings,omitzero"`
	paramObj
}

func (r DestinationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DestinationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DestinationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event-dispatch destination type to create. Warehouse and cloud-storage
// destination types may appear on read responses but are not creatable through
// POST.
type DestinationNewParamsType string

const (
	DestinationNewParamsTypeAudiohook                    DestinationNewParamsType = "Audiohook"
	DestinationNewParamsTypeBasisPostback                DestinationNewParamsType = "BasisPostback"
	DestinationNewParamsTypeOutbrain                     DestinationNewParamsType = "Outbrain"
	DestinationNewParamsTypeOursSyntheticData            DestinationNewParamsType = "OursSyntheticData"
	DestinationNewParamsTypeFullContact                  DestinationNewParamsType = "FullContact"
	DestinationNewParamsTypeZoomInfo                     DestinationNewParamsType = "ZoomInfo"
	DestinationNewParamsTypeTheTradeDesk                 DestinationNewParamsType = "TheTradeDesk"
	DestinationNewParamsTypeBraze                        DestinationNewParamsType = "Braze"
	DestinationNewParamsTypeLiveIntent                   DestinationNewParamsType = "LiveIntent"
	DestinationNewParamsTypeConvertAbTestingEvent        DestinationNewParamsType = "ConvertABTestingEvent"
	DestinationNewParamsTypeCustomerio                   DestinationNewParamsType = "Customerio"
	DestinationNewParamsTypeBingAds                      DestinationNewParamsType = "BingAds"
	DestinationNewParamsTypeBingAdsWeb                   DestinationNewParamsType = "BingAdsWeb"
	DestinationNewParamsTypeHTTPDestination              DestinationNewParamsType = "HTTPDestination"
	DestinationNewParamsTypeWoopra                       DestinationNewParamsType = "Woopra"
	DestinationNewParamsTypeHTTPCustomRequest            DestinationNewParamsType = "HTTPCustomRequest"
	DestinationNewParamsTypeGoogle                       DestinationNewParamsType = "Google"
	DestinationNewParamsTypeGoogleAdsServerContainer     DestinationNewParamsType = "GoogleAdsServerContainer"
	DestinationNewParamsTypeG4Analytics                  DestinationNewParamsType = "G4Analytics"
	DestinationNewParamsTypeGa4ServerProxy               DestinationNewParamsType = "GA4ServerProxy"
	DestinationNewParamsTypeGa4MeasurementProtocol       DestinationNewParamsType = "GA4MeasurementProtocol"
	DestinationNewParamsTypeGoogleAds360                 DestinationNewParamsType = "GoogleAds360"
	DestinationNewParamsTypeFacebook                     DestinationNewParamsType = "Facebook"
	DestinationNewParamsTypeMixpanel                     DestinationNewParamsType = "Mixpanel"
	DestinationNewParamsTypeAmplitude                    DestinationNewParamsType = "Amplitude"
	DestinationNewParamsTypeTikTok                       DestinationNewParamsType = "TikTok"
	DestinationNewParamsTypeReddit                       DestinationNewParamsType = "Reddit"
	DestinationNewParamsTypePodscribe                    DestinationNewParamsType = "Podscribe"
	DestinationNewParamsTypePinterest                    DestinationNewParamsType = "Pinterest"
	DestinationNewParamsTypeMailchimp                    DestinationNewParamsType = "Mailchimp"
	DestinationNewParamsTypeAwsKinesis                   DestinationNewParamsType = "AWSKinesis"
	DestinationNewParamsTypeAwsLambda                    DestinationNewParamsType = "AWSLambda"
	DestinationNewParamsTypeAwssns                       DestinationNewParamsType = "AWSSNS"
	DestinationNewParamsTypeGooglePubSub                 DestinationNewParamsType = "GooglePubSub"
	DestinationNewParamsTypeLinkedInAdsCapi              DestinationNewParamsType = "LinkedInAdsCAPI"
	DestinationNewParamsTypeActiveCampaignAPI            DestinationNewParamsType = "ActiveCampaignApi"
	DestinationNewParamsTypeStackAdaptAPI                DestinationNewParamsType = "StackAdaptAPI"
	DestinationNewParamsTypeHubspot                      DestinationNewParamsType = "Hubspot"
	DestinationNewParamsTypeKlaviyo                      DestinationNewParamsType = "Klaviyo"
	DestinationNewParamsTypeXAds                         DestinationNewParamsType = "XAds"
	DestinationNewParamsTypeQuoraAds                     DestinationNewParamsType = "QuoraAds"
	DestinationNewParamsTypeSnapchatAdsCapi              DestinationNewParamsType = "SnapchatAdsCapi"
	DestinationNewParamsTypePartnerize                   DestinationNewParamsType = "Partnerize"
	DestinationNewParamsTypeNextdoorAds                  DestinationNewParamsType = "NextdoorAds"
	DestinationNewParamsTypeTatari                       DestinationNewParamsType = "Tatari"
	DestinationNewParamsTypeViant                        DestinationNewParamsType = "Viant"
	DestinationNewParamsTypeViantCapi                    DestinationNewParamsType = "ViantCAPI"
	DestinationNewParamsTypeImpact                       DestinationNewParamsType = "Impact"
	DestinationNewParamsTypeSpotify                      DestinationNewParamsType = "Spotify"
	DestinationNewParamsTypeTaboola                      DestinationNewParamsType = "Taboola"
	DestinationNewParamsTypeAmazonDsp                    DestinationNewParamsType = "AmazonDSP"
	DestinationNewParamsTypeAppLovin                     DestinationNewParamsType = "AppLovin"
	DestinationNewParamsTypeIHeartMediaMagellan          DestinationNewParamsType = "IHeartMediaMagellan"
	DestinationNewParamsTypeVibe                         DestinationNewParamsType = "Vibe"
	DestinationNewParamsTypeGoogleDataManagerEventIngest DestinationNewParamsType = "GoogleDataManagerEventIngest"
	DestinationNewParamsTypeZendesk                      DestinationNewParamsType = "Zendesk"
	DestinationNewParamsTypeIterable                     DestinationNewParamsType = "Iterable"
	DestinationNewParamsTypeArtsAI                       DestinationNewParamsType = "ArtsAI"
	DestinationNewParamsTypeQuantcastCapi                DestinationNewParamsType = "QuantcastCAPI"
	DestinationNewParamsTypeFloodlightSgtm               DestinationNewParamsType = "FloodlightSGTM"
	DestinationNewParamsTypeVwo                          DestinationNewParamsType = "VWO"
	DestinationNewParamsTypeAttentive                    DestinationNewParamsType = "Attentive"
	DestinationNewParamsTypeAdmitad                      DestinationNewParamsType = "Admitad"
	DestinationNewParamsTypePlausible                    DestinationNewParamsType = "Plausible"
	DestinationNewParamsTypePostHog                      DestinationNewParamsType = "PostHog"
	DestinationNewParamsTypeRokuCapi                     DestinationNewParamsType = "RokuCAPI"
	DestinationNewParamsTypeEverflow                     DestinationNewParamsType = "Everflow"
	DestinationNewParamsTypeBeeswaxPostback              DestinationNewParamsType = "BeeswaxPostback"
	DestinationNewParamsTypeAdobeAnalytics               DestinationNewParamsType = "AdobeAnalytics"
	DestinationNewParamsTypeUniversalAds                 DestinationNewParamsType = "UniversalAds"
	DestinationNewParamsTypeOpenAIAds                    DestinationNewParamsType = "OpenAIAds"
	DestinationNewParamsTypeYelpCapi                     DestinationNewParamsType = "YelpCAPI"
	DestinationNewParamsTypeMntn                         DestinationNewParamsType = "MNTN"
	DestinationNewParamsTypeZohoCRM                      DestinationNewParamsType = "ZohoCRM"
)

type DestinationUpdateParams struct {
	HashingSalt        param.Opt[string] `json:"hashingSalt,omitzero"`
	Name               param.Opt[string] `json:"name,omitzero"`
	LimitedToSourceIDs []string          `json:"limitedToSourceIds,omitzero"`
	// Per-type configuration keys and values. Call GET /rest/v1/destinations/types to
	// get the valid keys for your destination type.
	Settings any `json:"settings,omitzero"`
	// Any of "Disabled", "Enabled".
	Status DestinationUpdateParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r DestinationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DestinationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DestinationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationUpdateParamsStatus string

const (
	DestinationUpdateParamsStatusDisabled DestinationUpdateParamsStatus = "Disabled"
	DestinationUpdateParamsStatusEnabled  DestinationUpdateParamsStatus = "Enabled"
)
