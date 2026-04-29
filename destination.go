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

// Update a destination. Requires scope: destination:update
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

// List all destinations. Requires scope: destination:list
func (r *DestinationService) List(ctx context.Context, opts ...option.RequestOption) (res *DestinationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/destinations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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

type DestinationNewResponse struct {
	ID        string `json:"id" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	// Any of "Disabled", "Enabled".
	Status DestinationNewResponseStatus `json:"status" api:"required"`
	// Any of "AWSEventBridge", "AWSKinesis", "AWSLambda", "AWSS3", "AWSSNS",
	// "ActiveCampaignApi", "Admitad", "AmazonDSP", "Amplitude", "AppLovin", "ArtsAI",
	// "Attentive", "Audiohook", "AzureBlob", "BasisPostback", "BingAds", "BingAdsWeb",
	// "Braze", "ConvertABTestingEvent", "Customerio", "DomoWarehouse", "Facebook",
	// "FloodlightSGTM", "FullContact", "G4Analytics", "GA4MeasurementProtocol",
	// "GA4ServerProxy", "Google", "GoogleAds360", "GoogleAdsServerContainer",
	// "GoogleBigQuery", "GoogleBigQueryWarehouse", "GoogleDataManagerEventIngest",
	// "GooglePubSub", "GoogleStorage", "HTTPCustomRequest", "HTTPDestination",
	// "Hubspot", "IHeartMediaMagellan", "Impact", "Iterable", "Klaviyo",
	// "LinkedInAdsCAPI", "LiveIntent", "LiveRampWarehouse", "Mailchimp", "Mixpanel",
	// "NextdoorAds", "OursSyntheticData", "Partnerize", "Pinterest", "Plausible",
	// "Podscribe", "PostHog", "QuantcastCAPI", "QuoraAds", "Reddit", "RokuCAPI",
	// "SnapchatAdsCapi", "Spotify", "StackAdaptAPI", "Taboola", "Tatari",
	// "TheTradeDesk", "TikTok", "VWO", "Viant", "Vibe", "Woopra", "XAds", "Zendesk",
	// "ZoomInfo".
	Type      DestinationNewResponseType `json:"type" api:"required"`
	Name      string                     `json:"name" api:"nullable"`
	UpdatedAt string                     `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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

type DestinationNewResponseType string

const (
	DestinationNewResponseTypeAwsEventBridge               DestinationNewResponseType = "AWSEventBridge"
	DestinationNewResponseTypeAwsKinesis                   DestinationNewResponseType = "AWSKinesis"
	DestinationNewResponseTypeAwsLambda                    DestinationNewResponseType = "AWSLambda"
	DestinationNewResponseTypeAwss3                        DestinationNewResponseType = "AWSS3"
	DestinationNewResponseTypeAwssns                       DestinationNewResponseType = "AWSSNS"
	DestinationNewResponseTypeActiveCampaignAPI            DestinationNewResponseType = "ActiveCampaignApi"
	DestinationNewResponseTypeAdmitad                      DestinationNewResponseType = "Admitad"
	DestinationNewResponseTypeAmazonDsp                    DestinationNewResponseType = "AmazonDSP"
	DestinationNewResponseTypeAmplitude                    DestinationNewResponseType = "Amplitude"
	DestinationNewResponseTypeAppLovin                     DestinationNewResponseType = "AppLovin"
	DestinationNewResponseTypeArtsAI                       DestinationNewResponseType = "ArtsAI"
	DestinationNewResponseTypeAttentive                    DestinationNewResponseType = "Attentive"
	DestinationNewResponseTypeAudiohook                    DestinationNewResponseType = "Audiohook"
	DestinationNewResponseTypeAzureBlob                    DestinationNewResponseType = "AzureBlob"
	DestinationNewResponseTypeBasisPostback                DestinationNewResponseType = "BasisPostback"
	DestinationNewResponseTypeBingAds                      DestinationNewResponseType = "BingAds"
	DestinationNewResponseTypeBingAdsWeb                   DestinationNewResponseType = "BingAdsWeb"
	DestinationNewResponseTypeBraze                        DestinationNewResponseType = "Braze"
	DestinationNewResponseTypeConvertAbTestingEvent        DestinationNewResponseType = "ConvertABTestingEvent"
	DestinationNewResponseTypeCustomerio                   DestinationNewResponseType = "Customerio"
	DestinationNewResponseTypeDomoWarehouse                DestinationNewResponseType = "DomoWarehouse"
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
	DestinationNewResponseTypeMailchimp                    DestinationNewResponseType = "Mailchimp"
	DestinationNewResponseTypeMixpanel                     DestinationNewResponseType = "Mixpanel"
	DestinationNewResponseTypeNextdoorAds                  DestinationNewResponseType = "NextdoorAds"
	DestinationNewResponseTypeOursSyntheticData            DestinationNewResponseType = "OursSyntheticData"
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
	DestinationNewResponseTypeVwo                          DestinationNewResponseType = "VWO"
	DestinationNewResponseTypeViant                        DestinationNewResponseType = "Viant"
	DestinationNewResponseTypeVibe                         DestinationNewResponseType = "Vibe"
	DestinationNewResponseTypeWoopra                       DestinationNewResponseType = "Woopra"
	DestinationNewResponseTypeXAds                         DestinationNewResponseType = "XAds"
	DestinationNewResponseTypeZendesk                      DestinationNewResponseType = "Zendesk"
	DestinationNewResponseTypeZoomInfo                     DestinationNewResponseType = "ZoomInfo"
)

type DestinationGetResponse struct {
	ID        string `json:"id" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	// Any of "Disabled", "Enabled".
	Status DestinationGetResponseStatus `json:"status" api:"required"`
	// Any of "AWSEventBridge", "AWSKinesis", "AWSLambda", "AWSS3", "AWSSNS",
	// "ActiveCampaignApi", "Admitad", "AmazonDSP", "Amplitude", "AppLovin", "ArtsAI",
	// "Attentive", "Audiohook", "AzureBlob", "BasisPostback", "BingAds", "BingAdsWeb",
	// "Braze", "ConvertABTestingEvent", "Customerio", "DomoWarehouse", "Facebook",
	// "FloodlightSGTM", "FullContact", "G4Analytics", "GA4MeasurementProtocol",
	// "GA4ServerProxy", "Google", "GoogleAds360", "GoogleAdsServerContainer",
	// "GoogleBigQuery", "GoogleBigQueryWarehouse", "GoogleDataManagerEventIngest",
	// "GooglePubSub", "GoogleStorage", "HTTPCustomRequest", "HTTPDestination",
	// "Hubspot", "IHeartMediaMagellan", "Impact", "Iterable", "Klaviyo",
	// "LinkedInAdsCAPI", "LiveIntent", "LiveRampWarehouse", "Mailchimp", "Mixpanel",
	// "NextdoorAds", "OursSyntheticData", "Partnerize", "Pinterest", "Plausible",
	// "Podscribe", "PostHog", "QuantcastCAPI", "QuoraAds", "Reddit", "RokuCAPI",
	// "SnapchatAdsCapi", "Spotify", "StackAdaptAPI", "Taboola", "Tatari",
	// "TheTradeDesk", "TikTok", "VWO", "Viant", "Vibe", "Woopra", "XAds", "Zendesk",
	// "ZoomInfo".
	Type                     DestinationGetResponseType `json:"type" api:"required"`
	FacebookConversionAPIKey string                     `json:"facebookConversionAPIKey" api:"nullable"`
	FacebookPixelID          string                     `json:"facebookPixelId" api:"nullable"`
	G4AnalyticsAPIKey        string                     `json:"g4AnalyticsApiKey" api:"nullable"`
	G4AnalyticsMeasurementID string                     `json:"g4AnalyticsMeasurementId" api:"nullable"`
	G4AnalyticsTrackOnPage   bool                       `json:"g4AnalyticsTrackOnPage" api:"nullable"`
	HashingSalt              string                     `json:"hashingSalt" api:"nullable"`
	HTTPDestinationURL       string                     `json:"httpDestinationUrl" api:"nullable"`
	LimitedToSourceIDs       []any                      `json:"limitedToSourceIds" api:"nullable"`
	ManagerGoogleCustomerID  string                     `json:"managerGoogleCustomerId" api:"nullable"`
	Name                     string                     `json:"name" api:"nullable"`
	ProjectAPIKey            string                     `json:"projectAPIKey" api:"nullable"`
	ProjectToken             string                     `json:"projectToken" api:"nullable"`
	SelectedAccountID        string                     `json:"selectedAccountId" api:"nullable"`
	Settings                 any                        `json:"settings" api:"nullable"`
	UpdatedAt                string                     `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		CreatedAt                respjson.Field
		Status                   respjson.Field
		Type                     respjson.Field
		FacebookConversionAPIKey respjson.Field
		FacebookPixelID          respjson.Field
		G4AnalyticsAPIKey        respjson.Field
		G4AnalyticsMeasurementID respjson.Field
		G4AnalyticsTrackOnPage   respjson.Field
		HashingSalt              respjson.Field
		HTTPDestinationURL       respjson.Field
		LimitedToSourceIDs       respjson.Field
		ManagerGoogleCustomerID  respjson.Field
		Name                     respjson.Field
		ProjectAPIKey            respjson.Field
		ProjectToken             respjson.Field
		SelectedAccountID        respjson.Field
		Settings                 respjson.Field
		UpdatedAt                respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
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

type DestinationGetResponseType string

const (
	DestinationGetResponseTypeAwsEventBridge               DestinationGetResponseType = "AWSEventBridge"
	DestinationGetResponseTypeAwsKinesis                   DestinationGetResponseType = "AWSKinesis"
	DestinationGetResponseTypeAwsLambda                    DestinationGetResponseType = "AWSLambda"
	DestinationGetResponseTypeAwss3                        DestinationGetResponseType = "AWSS3"
	DestinationGetResponseTypeAwssns                       DestinationGetResponseType = "AWSSNS"
	DestinationGetResponseTypeActiveCampaignAPI            DestinationGetResponseType = "ActiveCampaignApi"
	DestinationGetResponseTypeAdmitad                      DestinationGetResponseType = "Admitad"
	DestinationGetResponseTypeAmazonDsp                    DestinationGetResponseType = "AmazonDSP"
	DestinationGetResponseTypeAmplitude                    DestinationGetResponseType = "Amplitude"
	DestinationGetResponseTypeAppLovin                     DestinationGetResponseType = "AppLovin"
	DestinationGetResponseTypeArtsAI                       DestinationGetResponseType = "ArtsAI"
	DestinationGetResponseTypeAttentive                    DestinationGetResponseType = "Attentive"
	DestinationGetResponseTypeAudiohook                    DestinationGetResponseType = "Audiohook"
	DestinationGetResponseTypeAzureBlob                    DestinationGetResponseType = "AzureBlob"
	DestinationGetResponseTypeBasisPostback                DestinationGetResponseType = "BasisPostback"
	DestinationGetResponseTypeBingAds                      DestinationGetResponseType = "BingAds"
	DestinationGetResponseTypeBingAdsWeb                   DestinationGetResponseType = "BingAdsWeb"
	DestinationGetResponseTypeBraze                        DestinationGetResponseType = "Braze"
	DestinationGetResponseTypeConvertAbTestingEvent        DestinationGetResponseType = "ConvertABTestingEvent"
	DestinationGetResponseTypeCustomerio                   DestinationGetResponseType = "Customerio"
	DestinationGetResponseTypeDomoWarehouse                DestinationGetResponseType = "DomoWarehouse"
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
	DestinationGetResponseTypeMailchimp                    DestinationGetResponseType = "Mailchimp"
	DestinationGetResponseTypeMixpanel                     DestinationGetResponseType = "Mixpanel"
	DestinationGetResponseTypeNextdoorAds                  DestinationGetResponseType = "NextdoorAds"
	DestinationGetResponseTypeOursSyntheticData            DestinationGetResponseType = "OursSyntheticData"
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
	DestinationGetResponseTypeVwo                          DestinationGetResponseType = "VWO"
	DestinationGetResponseTypeViant                        DestinationGetResponseType = "Viant"
	DestinationGetResponseTypeVibe                         DestinationGetResponseType = "Vibe"
	DestinationGetResponseTypeWoopra                       DestinationGetResponseType = "Woopra"
	DestinationGetResponseTypeXAds                         DestinationGetResponseType = "XAds"
	DestinationGetResponseTypeZendesk                      DestinationGetResponseType = "Zendesk"
	DestinationGetResponseTypeZoomInfo                     DestinationGetResponseType = "ZoomInfo"
)

type DestinationUpdateResponse struct {
	ID        string `json:"id" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	// Any of "Disabled", "Enabled".
	Status DestinationUpdateResponseStatus `json:"status" api:"required"`
	// Any of "AWSEventBridge", "AWSKinesis", "AWSLambda", "AWSS3", "AWSSNS",
	// "ActiveCampaignApi", "Admitad", "AmazonDSP", "Amplitude", "AppLovin", "ArtsAI",
	// "Attentive", "Audiohook", "AzureBlob", "BasisPostback", "BingAds", "BingAdsWeb",
	// "Braze", "ConvertABTestingEvent", "Customerio", "DomoWarehouse", "Facebook",
	// "FloodlightSGTM", "FullContact", "G4Analytics", "GA4MeasurementProtocol",
	// "GA4ServerProxy", "Google", "GoogleAds360", "GoogleAdsServerContainer",
	// "GoogleBigQuery", "GoogleBigQueryWarehouse", "GoogleDataManagerEventIngest",
	// "GooglePubSub", "GoogleStorage", "HTTPCustomRequest", "HTTPDestination",
	// "Hubspot", "IHeartMediaMagellan", "Impact", "Iterable", "Klaviyo",
	// "LinkedInAdsCAPI", "LiveIntent", "LiveRampWarehouse", "Mailchimp", "Mixpanel",
	// "NextdoorAds", "OursSyntheticData", "Partnerize", "Pinterest", "Plausible",
	// "Podscribe", "PostHog", "QuantcastCAPI", "QuoraAds", "Reddit", "RokuCAPI",
	// "SnapchatAdsCapi", "Spotify", "StackAdaptAPI", "Taboola", "Tatari",
	// "TheTradeDesk", "TikTok", "VWO", "Viant", "Vibe", "Woopra", "XAds", "Zendesk",
	// "ZoomInfo".
	Type      DestinationUpdateResponseType `json:"type" api:"required"`
	Name      string                        `json:"name" api:"nullable"`
	UpdatedAt string                        `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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

type DestinationUpdateResponseType string

const (
	DestinationUpdateResponseTypeAwsEventBridge               DestinationUpdateResponseType = "AWSEventBridge"
	DestinationUpdateResponseTypeAwsKinesis                   DestinationUpdateResponseType = "AWSKinesis"
	DestinationUpdateResponseTypeAwsLambda                    DestinationUpdateResponseType = "AWSLambda"
	DestinationUpdateResponseTypeAwss3                        DestinationUpdateResponseType = "AWSS3"
	DestinationUpdateResponseTypeAwssns                       DestinationUpdateResponseType = "AWSSNS"
	DestinationUpdateResponseTypeActiveCampaignAPI            DestinationUpdateResponseType = "ActiveCampaignApi"
	DestinationUpdateResponseTypeAdmitad                      DestinationUpdateResponseType = "Admitad"
	DestinationUpdateResponseTypeAmazonDsp                    DestinationUpdateResponseType = "AmazonDSP"
	DestinationUpdateResponseTypeAmplitude                    DestinationUpdateResponseType = "Amplitude"
	DestinationUpdateResponseTypeAppLovin                     DestinationUpdateResponseType = "AppLovin"
	DestinationUpdateResponseTypeArtsAI                       DestinationUpdateResponseType = "ArtsAI"
	DestinationUpdateResponseTypeAttentive                    DestinationUpdateResponseType = "Attentive"
	DestinationUpdateResponseTypeAudiohook                    DestinationUpdateResponseType = "Audiohook"
	DestinationUpdateResponseTypeAzureBlob                    DestinationUpdateResponseType = "AzureBlob"
	DestinationUpdateResponseTypeBasisPostback                DestinationUpdateResponseType = "BasisPostback"
	DestinationUpdateResponseTypeBingAds                      DestinationUpdateResponseType = "BingAds"
	DestinationUpdateResponseTypeBingAdsWeb                   DestinationUpdateResponseType = "BingAdsWeb"
	DestinationUpdateResponseTypeBraze                        DestinationUpdateResponseType = "Braze"
	DestinationUpdateResponseTypeConvertAbTestingEvent        DestinationUpdateResponseType = "ConvertABTestingEvent"
	DestinationUpdateResponseTypeCustomerio                   DestinationUpdateResponseType = "Customerio"
	DestinationUpdateResponseTypeDomoWarehouse                DestinationUpdateResponseType = "DomoWarehouse"
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
	DestinationUpdateResponseTypeMailchimp                    DestinationUpdateResponseType = "Mailchimp"
	DestinationUpdateResponseTypeMixpanel                     DestinationUpdateResponseType = "Mixpanel"
	DestinationUpdateResponseTypeNextdoorAds                  DestinationUpdateResponseType = "NextdoorAds"
	DestinationUpdateResponseTypeOursSyntheticData            DestinationUpdateResponseType = "OursSyntheticData"
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
	DestinationUpdateResponseTypeVwo                          DestinationUpdateResponseType = "VWO"
	DestinationUpdateResponseTypeViant                        DestinationUpdateResponseType = "Viant"
	DestinationUpdateResponseTypeVibe                         DestinationUpdateResponseType = "Vibe"
	DestinationUpdateResponseTypeWoopra                       DestinationUpdateResponseType = "Woopra"
	DestinationUpdateResponseTypeXAds                         DestinationUpdateResponseType = "XAds"
	DestinationUpdateResponseTypeZendesk                      DestinationUpdateResponseType = "Zendesk"
	DestinationUpdateResponseTypeZoomInfo                     DestinationUpdateResponseType = "ZoomInfo"
)

type DestinationListResponse struct {
	Entities []DestinationListResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DestinationListResponse) RawJSON() string { return r.JSON.raw }
func (r *DestinationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationListResponseEntity struct {
	ID        string `json:"id" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	// Any of "Disabled", "Enabled".
	Status string `json:"status" api:"required"`
	// Any of "AWSEventBridge", "AWSKinesis", "AWSLambda", "AWSS3", "AWSSNS",
	// "ActiveCampaignApi", "Admitad", "AmazonDSP", "Amplitude", "AppLovin", "ArtsAI",
	// "Attentive", "Audiohook", "AzureBlob", "BasisPostback", "BingAds", "BingAdsWeb",
	// "Braze", "ConvertABTestingEvent", "Customerio", "DomoWarehouse", "Facebook",
	// "FloodlightSGTM", "FullContact", "G4Analytics", "GA4MeasurementProtocol",
	// "GA4ServerProxy", "Google", "GoogleAds360", "GoogleAdsServerContainer",
	// "GoogleBigQuery", "GoogleBigQueryWarehouse", "GoogleDataManagerEventIngest",
	// "GooglePubSub", "GoogleStorage", "HTTPCustomRequest", "HTTPDestination",
	// "Hubspot", "IHeartMediaMagellan", "Impact", "Iterable", "Klaviyo",
	// "LinkedInAdsCAPI", "LiveIntent", "LiveRampWarehouse", "Mailchimp", "Mixpanel",
	// "NextdoorAds", "OursSyntheticData", "Partnerize", "Pinterest", "Plausible",
	// "Podscribe", "PostHog", "QuantcastCAPI", "QuoraAds", "Reddit", "RokuCAPI",
	// "SnapchatAdsCapi", "Spotify", "StackAdaptAPI", "Taboola", "Tatari",
	// "TheTradeDesk", "TikTok", "VWO", "Viant", "Vibe", "Woopra", "XAds", "Zendesk",
	// "ZoomInfo".
	Type                     string `json:"type" api:"required"`
	FacebookConversionAPIKey string `json:"facebookConversionAPIKey" api:"nullable"`
	FacebookPixelID          string `json:"facebookPixelId" api:"nullable"`
	G4AnalyticsAPIKey        string `json:"g4AnalyticsApiKey" api:"nullable"`
	G4AnalyticsMeasurementID string `json:"g4AnalyticsMeasurementId" api:"nullable"`
	G4AnalyticsTrackOnPage   bool   `json:"g4AnalyticsTrackOnPage" api:"nullable"`
	HashingSalt              string `json:"hashingSalt" api:"nullable"`
	HTTPDestinationURL       string `json:"httpDestinationUrl" api:"nullable"`
	LimitedToSourceIDs       []any  `json:"limitedToSourceIds" api:"nullable"`
	ManagerGoogleCustomerID  string `json:"managerGoogleCustomerId" api:"nullable"`
	Name                     string `json:"name" api:"nullable"`
	ProjectAPIKey            string `json:"projectAPIKey" api:"nullable"`
	ProjectToken             string `json:"projectToken" api:"nullable"`
	SelectedAccountID        string `json:"selectedAccountId" api:"nullable"`
	Settings                 any    `json:"settings" api:"nullable"`
	UpdatedAt                string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		CreatedAt                respjson.Field
		Status                   respjson.Field
		Type                     respjson.Field
		FacebookConversionAPIKey respjson.Field
		FacebookPixelID          respjson.Field
		G4AnalyticsAPIKey        respjson.Field
		G4AnalyticsMeasurementID respjson.Field
		G4AnalyticsTrackOnPage   respjson.Field
		HashingSalt              respjson.Field
		HTTPDestinationURL       respjson.Field
		LimitedToSourceIDs       respjson.Field
		ManagerGoogleCustomerID  respjson.Field
		Name                     respjson.Field
		ProjectAPIKey            respjson.Field
		ProjectToken             respjson.Field
		SelectedAccountID        respjson.Field
		Settings                 respjson.Field
		UpdatedAt                respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DestinationListResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *DestinationListResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationNewParams struct {
	// Any of "AWSEventBridge", "AWSKinesis", "AWSLambda", "AWSS3", "AWSSNS",
	// "ActiveCampaignApi", "Admitad", "AmazonDSP", "Amplitude", "AppLovin", "ArtsAI",
	// "Attentive", "Audiohook", "AzureBlob", "BasisPostback", "BingAds", "BingAdsWeb",
	// "Braze", "ConvertABTestingEvent", "Customerio", "DomoWarehouse", "Facebook",
	// "FloodlightSGTM", "FullContact", "G4Analytics", "GA4MeasurementProtocol",
	// "GA4ServerProxy", "Google", "GoogleAds360", "GoogleAdsServerContainer",
	// "GoogleBigQuery", "GoogleBigQueryWarehouse", "GoogleDataManagerEventIngest",
	// "GooglePubSub", "GoogleStorage", "HTTPCustomRequest", "HTTPDestination",
	// "Hubspot", "IHeartMediaMagellan", "Impact", "Iterable", "Klaviyo",
	// "LinkedInAdsCAPI", "LiveIntent", "LiveRampWarehouse", "Mailchimp", "Mixpanel",
	// "NextdoorAds", "OursSyntheticData", "Partnerize", "Pinterest", "Plausible",
	// "Podscribe", "PostHog", "QuantcastCAPI", "QuoraAds", "Reddit", "RokuCAPI",
	// "SnapchatAdsCapi", "Spotify", "StackAdaptAPI", "Taboola", "Tatari",
	// "TheTradeDesk", "TikTok", "VWO", "Viant", "Vibe", "Woopra", "XAds", "Zendesk",
	// "ZoomInfo".
	Type DestinationNewParamsType `json:"type,omitzero" api:"required"`
	Name param.Opt[string]        `json:"name,omitzero"`
	paramObj
}

func (r DestinationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DestinationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DestinationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationNewParamsType string

const (
	DestinationNewParamsTypeAwsEventBridge               DestinationNewParamsType = "AWSEventBridge"
	DestinationNewParamsTypeAwsKinesis                   DestinationNewParamsType = "AWSKinesis"
	DestinationNewParamsTypeAwsLambda                    DestinationNewParamsType = "AWSLambda"
	DestinationNewParamsTypeAwss3                        DestinationNewParamsType = "AWSS3"
	DestinationNewParamsTypeAwssns                       DestinationNewParamsType = "AWSSNS"
	DestinationNewParamsTypeActiveCampaignAPI            DestinationNewParamsType = "ActiveCampaignApi"
	DestinationNewParamsTypeAdmitad                      DestinationNewParamsType = "Admitad"
	DestinationNewParamsTypeAmazonDsp                    DestinationNewParamsType = "AmazonDSP"
	DestinationNewParamsTypeAmplitude                    DestinationNewParamsType = "Amplitude"
	DestinationNewParamsTypeAppLovin                     DestinationNewParamsType = "AppLovin"
	DestinationNewParamsTypeArtsAI                       DestinationNewParamsType = "ArtsAI"
	DestinationNewParamsTypeAttentive                    DestinationNewParamsType = "Attentive"
	DestinationNewParamsTypeAudiohook                    DestinationNewParamsType = "Audiohook"
	DestinationNewParamsTypeAzureBlob                    DestinationNewParamsType = "AzureBlob"
	DestinationNewParamsTypeBasisPostback                DestinationNewParamsType = "BasisPostback"
	DestinationNewParamsTypeBingAds                      DestinationNewParamsType = "BingAds"
	DestinationNewParamsTypeBingAdsWeb                   DestinationNewParamsType = "BingAdsWeb"
	DestinationNewParamsTypeBraze                        DestinationNewParamsType = "Braze"
	DestinationNewParamsTypeConvertAbTestingEvent        DestinationNewParamsType = "ConvertABTestingEvent"
	DestinationNewParamsTypeCustomerio                   DestinationNewParamsType = "Customerio"
	DestinationNewParamsTypeDomoWarehouse                DestinationNewParamsType = "DomoWarehouse"
	DestinationNewParamsTypeFacebook                     DestinationNewParamsType = "Facebook"
	DestinationNewParamsTypeFloodlightSgtm               DestinationNewParamsType = "FloodlightSGTM"
	DestinationNewParamsTypeFullContact                  DestinationNewParamsType = "FullContact"
	DestinationNewParamsTypeG4Analytics                  DestinationNewParamsType = "G4Analytics"
	DestinationNewParamsTypeGa4MeasurementProtocol       DestinationNewParamsType = "GA4MeasurementProtocol"
	DestinationNewParamsTypeGa4ServerProxy               DestinationNewParamsType = "GA4ServerProxy"
	DestinationNewParamsTypeGoogle                       DestinationNewParamsType = "Google"
	DestinationNewParamsTypeGoogleAds360                 DestinationNewParamsType = "GoogleAds360"
	DestinationNewParamsTypeGoogleAdsServerContainer     DestinationNewParamsType = "GoogleAdsServerContainer"
	DestinationNewParamsTypeGoogleBigQuery               DestinationNewParamsType = "GoogleBigQuery"
	DestinationNewParamsTypeGoogleBigQueryWarehouse      DestinationNewParamsType = "GoogleBigQueryWarehouse"
	DestinationNewParamsTypeGoogleDataManagerEventIngest DestinationNewParamsType = "GoogleDataManagerEventIngest"
	DestinationNewParamsTypeGooglePubSub                 DestinationNewParamsType = "GooglePubSub"
	DestinationNewParamsTypeGoogleStorage                DestinationNewParamsType = "GoogleStorage"
	DestinationNewParamsTypeHTTPCustomRequest            DestinationNewParamsType = "HTTPCustomRequest"
	DestinationNewParamsTypeHTTPDestination              DestinationNewParamsType = "HTTPDestination"
	DestinationNewParamsTypeHubspot                      DestinationNewParamsType = "Hubspot"
	DestinationNewParamsTypeIHeartMediaMagellan          DestinationNewParamsType = "IHeartMediaMagellan"
	DestinationNewParamsTypeImpact                       DestinationNewParamsType = "Impact"
	DestinationNewParamsTypeIterable                     DestinationNewParamsType = "Iterable"
	DestinationNewParamsTypeKlaviyo                      DestinationNewParamsType = "Klaviyo"
	DestinationNewParamsTypeLinkedInAdsCapi              DestinationNewParamsType = "LinkedInAdsCAPI"
	DestinationNewParamsTypeLiveIntent                   DestinationNewParamsType = "LiveIntent"
	DestinationNewParamsTypeLiveRampWarehouse            DestinationNewParamsType = "LiveRampWarehouse"
	DestinationNewParamsTypeMailchimp                    DestinationNewParamsType = "Mailchimp"
	DestinationNewParamsTypeMixpanel                     DestinationNewParamsType = "Mixpanel"
	DestinationNewParamsTypeNextdoorAds                  DestinationNewParamsType = "NextdoorAds"
	DestinationNewParamsTypeOursSyntheticData            DestinationNewParamsType = "OursSyntheticData"
	DestinationNewParamsTypePartnerize                   DestinationNewParamsType = "Partnerize"
	DestinationNewParamsTypePinterest                    DestinationNewParamsType = "Pinterest"
	DestinationNewParamsTypePlausible                    DestinationNewParamsType = "Plausible"
	DestinationNewParamsTypePodscribe                    DestinationNewParamsType = "Podscribe"
	DestinationNewParamsTypePostHog                      DestinationNewParamsType = "PostHog"
	DestinationNewParamsTypeQuantcastCapi                DestinationNewParamsType = "QuantcastCAPI"
	DestinationNewParamsTypeQuoraAds                     DestinationNewParamsType = "QuoraAds"
	DestinationNewParamsTypeReddit                       DestinationNewParamsType = "Reddit"
	DestinationNewParamsTypeRokuCapi                     DestinationNewParamsType = "RokuCAPI"
	DestinationNewParamsTypeSnapchatAdsCapi              DestinationNewParamsType = "SnapchatAdsCapi"
	DestinationNewParamsTypeSpotify                      DestinationNewParamsType = "Spotify"
	DestinationNewParamsTypeStackAdaptAPI                DestinationNewParamsType = "StackAdaptAPI"
	DestinationNewParamsTypeTaboola                      DestinationNewParamsType = "Taboola"
	DestinationNewParamsTypeTatari                       DestinationNewParamsType = "Tatari"
	DestinationNewParamsTypeTheTradeDesk                 DestinationNewParamsType = "TheTradeDesk"
	DestinationNewParamsTypeTikTok                       DestinationNewParamsType = "TikTok"
	DestinationNewParamsTypeVwo                          DestinationNewParamsType = "VWO"
	DestinationNewParamsTypeViant                        DestinationNewParamsType = "Viant"
	DestinationNewParamsTypeVibe                         DestinationNewParamsType = "Vibe"
	DestinationNewParamsTypeWoopra                       DestinationNewParamsType = "Woopra"
	DestinationNewParamsTypeXAds                         DestinationNewParamsType = "XAds"
	DestinationNewParamsTypeZendesk                      DestinationNewParamsType = "Zendesk"
	DestinationNewParamsTypeZoomInfo                     DestinationNewParamsType = "ZoomInfo"
)

type DestinationUpdateParams struct {
	// Any of "Disabled", "Enabled".
	Status                   DestinationUpdateParamsStatus `json:"status,omitzero" api:"required"`
	FacebookConversionAPIKey param.Opt[string]             `json:"facebookConversionAPIKey,omitzero"`
	FacebookPixelID          param.Opt[string]             `json:"facebookPixelId,omitzero"`
	G4AnalyticsAPIKey        param.Opt[string]             `json:"g4AnalyticsApiKey,omitzero"`
	G4AnalyticsMeasurementID param.Opt[string]             `json:"g4AnalyticsMeasurementId,omitzero"`
	G4AnalyticsTrackOnPage   param.Opt[bool]               `json:"g4AnalyticsTrackOnPage,omitzero"`
	HashingSalt              param.Opt[string]             `json:"hashingSalt,omitzero"`
	HTTPDestinationURL       param.Opt[string]             `json:"httpDestinationUrl,omitzero"`
	ManagerGoogleCustomerID  param.Opt[string]             `json:"managerGoogleCustomerId,omitzero"`
	Name                     param.Opt[string]             `json:"name,omitzero"`
	ProjectAPIKey            param.Opt[string]             `json:"projectAPIKey,omitzero"`
	ProjectToken             param.Opt[string]             `json:"projectToken,omitzero"`
	SelectedAccountID        param.Opt[string]             `json:"selectedAccountId,omitzero"`
	LimitedToSourceIDs       []any                         `json:"limitedToSourceIds,omitzero"`
	Settings                 any                           `json:"settings,omitzero"`
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
