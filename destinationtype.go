// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	"github.com/with-ours/platform-sdk-go/internal/apijson"
	"github.com/with-ours/platform-sdk-go/internal/requestconfig"
	"github.com/with-ours/platform-sdk-go/option"
	"github.com/with-ours/platform-sdk-go/packages/respjson"
)

// DestinationTypeService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDestinationTypeService] method instead.
type DestinationTypeService struct {
	Options []option.RequestOption
}

// NewDestinationTypeService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDestinationTypeService(opts ...option.RequestOption) (r DestinationTypeService) {
	r = DestinationTypeService{}
	r.Options = opts
	return
}

// Lists every destination type the platform supports, with its human-readable
// label, capability flags (oauth, listsAccounts, supportsRenamedEvents), and the
// settings descriptor used to configure a destination of that type.
// Account-agnostic — the response is the same for every API key. Requires scope:
// destination:list
func (r *DestinationTypeService) List(ctx context.Context, opts ...option.RequestOption) (res *DestinationTypeListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/destination-types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Fetch the descriptor for a single destination type by its identifier (e.g.
// `Klaviyo`, `Facebook`, `GoogleDataManagerEventIngest`). Returns 404 if the
// identifier is unknown. Requires scope: destination:list
func (r *DestinationTypeService) Get(ctx context.Context, id DestinationTypeGetParamsID, opts ...option.RequestOption) (res *DestinationTypeGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("rest/v1/destination-types/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type DestinationTypeListResponse struct {
	Entities []DestinationTypeListResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DestinationTypeListResponse) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypeListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypeListResponseEntity struct {
	// Any of "AWSEventBridge", "AWSKinesis", "AWSLambda", "AWSS3", "AWSSNS",
	// "ActiveCampaignApi", "Admitad", "AmazonDSP", "Amplitude", "AppLovin", "ArtsAI",
	// "Attentive", "Audiohook", "AzureBlob", "BasisPostback", "BeeswaxPostback",
	// "BingAds", "BingAdsWeb", "Braze", "ConvertABTestingEvent", "Customerio",
	// "DomoWarehouse", "Everflow", "Facebook", "FloodlightSGTM", "FullContact",
	// "G4Analytics", "GA4MeasurementProtocol", "GA4ServerProxy", "Google",
	// "GoogleAds360", "GoogleAdsServerContainer", "GoogleBigQuery",
	// "GoogleBigQueryWarehouse", "GoogleDataManagerEventIngest", "GooglePubSub",
	// "GoogleStorage", "HTTPCustomRequest", "HTTPDestination", "Hubspot",
	// "IHeartMediaMagellan", "Impact", "Iterable", "Klaviyo", "LinkedInAdsCAPI",
	// "LiveIntent", "LiveRampWarehouse", "Mailchimp", "Mixpanel", "NextdoorAds",
	// "OursSyntheticData", "Partnerize", "Pinterest", "Plausible", "Podscribe",
	// "PostHog", "QuantcastCAPI", "QuoraAds", "Reddit", "RokuCAPI", "SnapchatAdsCapi",
	// "Spotify", "StackAdaptAPI", "Taboola", "Tatari", "TheTradeDesk", "TikTok",
	// "VWO", "Viant", "Vibe", "Woopra", "XAds", "Zendesk", "ZoomInfo".
	ID           string                                          `json:"id" api:"required"`
	Capabilities DestinationTypeListResponseEntityCapabilities   `json:"capabilities" api:"required"`
	Label        string                                          `json:"label" api:"required"`
	Settings     []DestinationTypeListResponseEntitySettingUnion `json:"settings" api:"required"`
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
func (r DestinationTypeListResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypeListResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypeListResponseEntityCapabilities struct {
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
func (r DestinationTypeListResponseEntityCapabilities) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypeListResponseEntityCapabilities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DestinationTypeListResponseEntitySettingUnion contains all possible properties
// and values from [DestinationTypeListResponseEntitySettingObject],
// [DestinationTypeListResponseEntitySettingObject2],
// [DestinationTypeListResponseEntitySettingObject3],
// [DestinationTypeListResponseEntitySettingObject4],
// [DestinationTypeListResponseEntitySettingObject5].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type DestinationTypeListResponseEntitySettingUnion struct {
	Key   string `json:"key"`
	Label string `json:"label"`
	Type  string `json:"type"`
	// This field is from variant [DestinationTypeListResponseEntitySettingObject2].
	Options  []DestinationTypeListResponseEntitySettingObject2Option `json:"options"`
	Required bool                                                    `json:"required"`
	Sublabel string                                                  `json:"sublabel"`
	// This field is from variant [DestinationTypeListResponseEntitySettingObject3].
	DefaultValue bool `json:"defaultValue"`
	// This field is from variant [DestinationTypeListResponseEntitySettingObject5].
	Placeholder string `json:"placeholder"`
	JSON        struct {
		Key          respjson.Field
		Label        respjson.Field
		Type         respjson.Field
		Options      respjson.Field
		Required     respjson.Field
		Sublabel     respjson.Field
		DefaultValue respjson.Field
		Placeholder  respjson.Field
		raw          string
	} `json:"-"`
}

func (u DestinationTypeListResponseEntitySettingUnion) AsDestinationTypeListResponseEntitySettingObject() (v DestinationTypeListResponseEntitySettingObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DestinationTypeListResponseEntitySettingUnion) AsDestinationTypeListResponseEntitySettingObject2() (v DestinationTypeListResponseEntitySettingObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DestinationTypeListResponseEntitySettingUnion) AsDestinationTypeListResponseEntitySettingObject3() (v DestinationTypeListResponseEntitySettingObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DestinationTypeListResponseEntitySettingUnion) AsDestinationTypeListResponseEntitySettingObject4() (v DestinationTypeListResponseEntitySettingObject4) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DestinationTypeListResponseEntitySettingUnion) AsDestinationTypeListResponseEntitySettingObject5() (v DestinationTypeListResponseEntitySettingObject5) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DestinationTypeListResponseEntitySettingUnion) RawJSON() string { return u.JSON.raw }

func (r *DestinationTypeListResponseEntitySettingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypeListResponseEntitySettingObject struct {
	Key   string `json:"key" api:"required"`
	Label string `json:"label" api:"required"`
	// Any of "Alert", "GenericOauth", "Secret", "Select", "Switch", "Text".
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
func (r DestinationTypeListResponseEntitySettingObject) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypeListResponseEntitySettingObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypeListResponseEntitySettingObject2 struct {
	Key     string                                                  `json:"key" api:"required"`
	Label   string                                                  `json:"label" api:"required"`
	Options []DestinationTypeListResponseEntitySettingObject2Option `json:"options" api:"required"`
	// Any of "Alert", "GenericOauth", "Secret", "Select", "Switch", "Text".
	Type     string `json:"type" api:"required"`
	Required bool   `json:"required" api:"nullable"`
	Sublabel string `json:"sublabel" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Label       respjson.Field
		Options     respjson.Field
		Type        respjson.Field
		Required    respjson.Field
		Sublabel    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DestinationTypeListResponseEntitySettingObject2) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypeListResponseEntitySettingObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypeListResponseEntitySettingObject2Option struct {
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
func (r DestinationTypeListResponseEntitySettingObject2Option) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypeListResponseEntitySettingObject2Option) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypeListResponseEntitySettingObject3 struct {
	Key   string `json:"key" api:"required"`
	Label string `json:"label" api:"required"`
	// Any of "Alert", "GenericOauth", "Secret", "Select", "Switch", "Text".
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
func (r DestinationTypeListResponseEntitySettingObject3) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypeListResponseEntitySettingObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypeListResponseEntitySettingObject4 struct {
	Key   string `json:"key" api:"required"`
	Label string `json:"label" api:"required"`
	// Any of "Alert", "GenericOauth", "Secret", "Select", "Switch", "Text".
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
func (r DestinationTypeListResponseEntitySettingObject4) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypeListResponseEntitySettingObject4) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypeListResponseEntitySettingObject5 struct {
	Key         string `json:"key" api:"required"`
	Label       string `json:"label" api:"required"`
	Placeholder string `json:"placeholder" api:"required"`
	// Any of "Alert", "GenericOauth", "Secret", "Select", "Switch", "Text".
	Type     string `json:"type" api:"required"`
	Required bool   `json:"required" api:"nullable"`
	Sublabel string `json:"sublabel" api:"nullable"`
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
func (r DestinationTypeListResponseEntitySettingObject5) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypeListResponseEntitySettingObject5) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypeGetResponse struct {
	// Any of "AWSEventBridge", "AWSKinesis", "AWSLambda", "AWSS3", "AWSSNS",
	// "ActiveCampaignApi", "Admitad", "AmazonDSP", "Amplitude", "AppLovin", "ArtsAI",
	// "Attentive", "Audiohook", "AzureBlob", "BasisPostback", "BeeswaxPostback",
	// "BingAds", "BingAdsWeb", "Braze", "ConvertABTestingEvent", "Customerio",
	// "DomoWarehouse", "Everflow", "Facebook", "FloodlightSGTM", "FullContact",
	// "G4Analytics", "GA4MeasurementProtocol", "GA4ServerProxy", "Google",
	// "GoogleAds360", "GoogleAdsServerContainer", "GoogleBigQuery",
	// "GoogleBigQueryWarehouse", "GoogleDataManagerEventIngest", "GooglePubSub",
	// "GoogleStorage", "HTTPCustomRequest", "HTTPDestination", "Hubspot",
	// "IHeartMediaMagellan", "Impact", "Iterable", "Klaviyo", "LinkedInAdsCAPI",
	// "LiveIntent", "LiveRampWarehouse", "Mailchimp", "Mixpanel", "NextdoorAds",
	// "OursSyntheticData", "Partnerize", "Pinterest", "Plausible", "Podscribe",
	// "PostHog", "QuantcastCAPI", "QuoraAds", "Reddit", "RokuCAPI", "SnapchatAdsCapi",
	// "Spotify", "StackAdaptAPI", "Taboola", "Tatari", "TheTradeDesk", "TikTok",
	// "VWO", "Viant", "Vibe", "Woopra", "XAds", "Zendesk", "ZoomInfo".
	ID           DestinationTypeGetResponseID             `json:"id" api:"required"`
	Capabilities DestinationTypeGetResponseCapabilities   `json:"capabilities" api:"required"`
	Label        string                                   `json:"label" api:"required"`
	Settings     []DestinationTypeGetResponseSettingUnion `json:"settings" api:"required"`
	// Any of "deprecated", "ga".
	Status DestinationTypeGetResponseStatus `json:"status" api:"required"`
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
func (r DestinationTypeGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypeGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypeGetResponseID string

const (
	DestinationTypeGetResponseIDAwsEventBridge               DestinationTypeGetResponseID = "AWSEventBridge"
	DestinationTypeGetResponseIDAwsKinesis                   DestinationTypeGetResponseID = "AWSKinesis"
	DestinationTypeGetResponseIDAwsLambda                    DestinationTypeGetResponseID = "AWSLambda"
	DestinationTypeGetResponseIDAwss3                        DestinationTypeGetResponseID = "AWSS3"
	DestinationTypeGetResponseIDAwssns                       DestinationTypeGetResponseID = "AWSSNS"
	DestinationTypeGetResponseIDActiveCampaignAPI            DestinationTypeGetResponseID = "ActiveCampaignApi"
	DestinationTypeGetResponseIDAdmitad                      DestinationTypeGetResponseID = "Admitad"
	DestinationTypeGetResponseIDAmazonDsp                    DestinationTypeGetResponseID = "AmazonDSP"
	DestinationTypeGetResponseIDAmplitude                    DestinationTypeGetResponseID = "Amplitude"
	DestinationTypeGetResponseIDAppLovin                     DestinationTypeGetResponseID = "AppLovin"
	DestinationTypeGetResponseIDArtsAI                       DestinationTypeGetResponseID = "ArtsAI"
	DestinationTypeGetResponseIDAttentive                    DestinationTypeGetResponseID = "Attentive"
	DestinationTypeGetResponseIDAudiohook                    DestinationTypeGetResponseID = "Audiohook"
	DestinationTypeGetResponseIDAzureBlob                    DestinationTypeGetResponseID = "AzureBlob"
	DestinationTypeGetResponseIDBasisPostback                DestinationTypeGetResponseID = "BasisPostback"
	DestinationTypeGetResponseIDBeeswaxPostback              DestinationTypeGetResponseID = "BeeswaxPostback"
	DestinationTypeGetResponseIDBingAds                      DestinationTypeGetResponseID = "BingAds"
	DestinationTypeGetResponseIDBingAdsWeb                   DestinationTypeGetResponseID = "BingAdsWeb"
	DestinationTypeGetResponseIDBraze                        DestinationTypeGetResponseID = "Braze"
	DestinationTypeGetResponseIDConvertAbTestingEvent        DestinationTypeGetResponseID = "ConvertABTestingEvent"
	DestinationTypeGetResponseIDCustomerio                   DestinationTypeGetResponseID = "Customerio"
	DestinationTypeGetResponseIDDomoWarehouse                DestinationTypeGetResponseID = "DomoWarehouse"
	DestinationTypeGetResponseIDEverflow                     DestinationTypeGetResponseID = "Everflow"
	DestinationTypeGetResponseIDFacebook                     DestinationTypeGetResponseID = "Facebook"
	DestinationTypeGetResponseIDFloodlightSgtm               DestinationTypeGetResponseID = "FloodlightSGTM"
	DestinationTypeGetResponseIDFullContact                  DestinationTypeGetResponseID = "FullContact"
	DestinationTypeGetResponseIDG4Analytics                  DestinationTypeGetResponseID = "G4Analytics"
	DestinationTypeGetResponseIDGa4MeasurementProtocol       DestinationTypeGetResponseID = "GA4MeasurementProtocol"
	DestinationTypeGetResponseIDGa4ServerProxy               DestinationTypeGetResponseID = "GA4ServerProxy"
	DestinationTypeGetResponseIDGoogle                       DestinationTypeGetResponseID = "Google"
	DestinationTypeGetResponseIDGoogleAds360                 DestinationTypeGetResponseID = "GoogleAds360"
	DestinationTypeGetResponseIDGoogleAdsServerContainer     DestinationTypeGetResponseID = "GoogleAdsServerContainer"
	DestinationTypeGetResponseIDGoogleBigQuery               DestinationTypeGetResponseID = "GoogleBigQuery"
	DestinationTypeGetResponseIDGoogleBigQueryWarehouse      DestinationTypeGetResponseID = "GoogleBigQueryWarehouse"
	DestinationTypeGetResponseIDGoogleDataManagerEventIngest DestinationTypeGetResponseID = "GoogleDataManagerEventIngest"
	DestinationTypeGetResponseIDGooglePubSub                 DestinationTypeGetResponseID = "GooglePubSub"
	DestinationTypeGetResponseIDGoogleStorage                DestinationTypeGetResponseID = "GoogleStorage"
	DestinationTypeGetResponseIDHTTPCustomRequest            DestinationTypeGetResponseID = "HTTPCustomRequest"
	DestinationTypeGetResponseIDHTTPDestination              DestinationTypeGetResponseID = "HTTPDestination"
	DestinationTypeGetResponseIDHubspot                      DestinationTypeGetResponseID = "Hubspot"
	DestinationTypeGetResponseIDIHeartMediaMagellan          DestinationTypeGetResponseID = "IHeartMediaMagellan"
	DestinationTypeGetResponseIDImpact                       DestinationTypeGetResponseID = "Impact"
	DestinationTypeGetResponseIDIterable                     DestinationTypeGetResponseID = "Iterable"
	DestinationTypeGetResponseIDKlaviyo                      DestinationTypeGetResponseID = "Klaviyo"
	DestinationTypeGetResponseIDLinkedInAdsCapi              DestinationTypeGetResponseID = "LinkedInAdsCAPI"
	DestinationTypeGetResponseIDLiveIntent                   DestinationTypeGetResponseID = "LiveIntent"
	DestinationTypeGetResponseIDLiveRampWarehouse            DestinationTypeGetResponseID = "LiveRampWarehouse"
	DestinationTypeGetResponseIDMailchimp                    DestinationTypeGetResponseID = "Mailchimp"
	DestinationTypeGetResponseIDMixpanel                     DestinationTypeGetResponseID = "Mixpanel"
	DestinationTypeGetResponseIDNextdoorAds                  DestinationTypeGetResponseID = "NextdoorAds"
	DestinationTypeGetResponseIDOursSyntheticData            DestinationTypeGetResponseID = "OursSyntheticData"
	DestinationTypeGetResponseIDPartnerize                   DestinationTypeGetResponseID = "Partnerize"
	DestinationTypeGetResponseIDPinterest                    DestinationTypeGetResponseID = "Pinterest"
	DestinationTypeGetResponseIDPlausible                    DestinationTypeGetResponseID = "Plausible"
	DestinationTypeGetResponseIDPodscribe                    DestinationTypeGetResponseID = "Podscribe"
	DestinationTypeGetResponseIDPostHog                      DestinationTypeGetResponseID = "PostHog"
	DestinationTypeGetResponseIDQuantcastCapi                DestinationTypeGetResponseID = "QuantcastCAPI"
	DestinationTypeGetResponseIDQuoraAds                     DestinationTypeGetResponseID = "QuoraAds"
	DestinationTypeGetResponseIDReddit                       DestinationTypeGetResponseID = "Reddit"
	DestinationTypeGetResponseIDRokuCapi                     DestinationTypeGetResponseID = "RokuCAPI"
	DestinationTypeGetResponseIDSnapchatAdsCapi              DestinationTypeGetResponseID = "SnapchatAdsCapi"
	DestinationTypeGetResponseIDSpotify                      DestinationTypeGetResponseID = "Spotify"
	DestinationTypeGetResponseIDStackAdaptAPI                DestinationTypeGetResponseID = "StackAdaptAPI"
	DestinationTypeGetResponseIDTaboola                      DestinationTypeGetResponseID = "Taboola"
	DestinationTypeGetResponseIDTatari                       DestinationTypeGetResponseID = "Tatari"
	DestinationTypeGetResponseIDTheTradeDesk                 DestinationTypeGetResponseID = "TheTradeDesk"
	DestinationTypeGetResponseIDTikTok                       DestinationTypeGetResponseID = "TikTok"
	DestinationTypeGetResponseIDVwo                          DestinationTypeGetResponseID = "VWO"
	DestinationTypeGetResponseIDViant                        DestinationTypeGetResponseID = "Viant"
	DestinationTypeGetResponseIDVibe                         DestinationTypeGetResponseID = "Vibe"
	DestinationTypeGetResponseIDWoopra                       DestinationTypeGetResponseID = "Woopra"
	DestinationTypeGetResponseIDXAds                         DestinationTypeGetResponseID = "XAds"
	DestinationTypeGetResponseIDZendesk                      DestinationTypeGetResponseID = "Zendesk"
	DestinationTypeGetResponseIDZoomInfo                     DestinationTypeGetResponseID = "ZoomInfo"
)

type DestinationTypeGetResponseCapabilities struct {
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
func (r DestinationTypeGetResponseCapabilities) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypeGetResponseCapabilities) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DestinationTypeGetResponseSettingUnion contains all possible properties and
// values from [DestinationTypeGetResponseSettingObject],
// [DestinationTypeGetResponseSettingObject2],
// [DestinationTypeGetResponseSettingObject3],
// [DestinationTypeGetResponseSettingObject4],
// [DestinationTypeGetResponseSettingObject5].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type DestinationTypeGetResponseSettingUnion struct {
	Key   string `json:"key"`
	Label string `json:"label"`
	Type  string `json:"type"`
	// This field is from variant [DestinationTypeGetResponseSettingObject2].
	Options  []DestinationTypeGetResponseSettingObject2Option `json:"options"`
	Required bool                                             `json:"required"`
	Sublabel string                                           `json:"sublabel"`
	// This field is from variant [DestinationTypeGetResponseSettingObject3].
	DefaultValue bool `json:"defaultValue"`
	// This field is from variant [DestinationTypeGetResponseSettingObject5].
	Placeholder string `json:"placeholder"`
	JSON        struct {
		Key          respjson.Field
		Label        respjson.Field
		Type         respjson.Field
		Options      respjson.Field
		Required     respjson.Field
		Sublabel     respjson.Field
		DefaultValue respjson.Field
		Placeholder  respjson.Field
		raw          string
	} `json:"-"`
}

func (u DestinationTypeGetResponseSettingUnion) AsDestinationTypeGetResponseSettingObject() (v DestinationTypeGetResponseSettingObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DestinationTypeGetResponseSettingUnion) AsDestinationTypeGetResponseSettingObject2() (v DestinationTypeGetResponseSettingObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DestinationTypeGetResponseSettingUnion) AsDestinationTypeGetResponseSettingObject3() (v DestinationTypeGetResponseSettingObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DestinationTypeGetResponseSettingUnion) AsDestinationTypeGetResponseSettingObject4() (v DestinationTypeGetResponseSettingObject4) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DestinationTypeGetResponseSettingUnion) AsDestinationTypeGetResponseSettingObject5() (v DestinationTypeGetResponseSettingObject5) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DestinationTypeGetResponseSettingUnion) RawJSON() string { return u.JSON.raw }

func (r *DestinationTypeGetResponseSettingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypeGetResponseSettingObject struct {
	Key   string `json:"key" api:"required"`
	Label string `json:"label" api:"required"`
	// Any of "Alert", "GenericOauth", "Secret", "Select", "Switch", "Text".
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
func (r DestinationTypeGetResponseSettingObject) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypeGetResponseSettingObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypeGetResponseSettingObject2 struct {
	Key     string                                           `json:"key" api:"required"`
	Label   string                                           `json:"label" api:"required"`
	Options []DestinationTypeGetResponseSettingObject2Option `json:"options" api:"required"`
	// Any of "Alert", "GenericOauth", "Secret", "Select", "Switch", "Text".
	Type     string `json:"type" api:"required"`
	Required bool   `json:"required" api:"nullable"`
	Sublabel string `json:"sublabel" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Label       respjson.Field
		Options     respjson.Field
		Type        respjson.Field
		Required    respjson.Field
		Sublabel    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DestinationTypeGetResponseSettingObject2) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypeGetResponseSettingObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypeGetResponseSettingObject2Option struct {
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
func (r DestinationTypeGetResponseSettingObject2Option) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypeGetResponseSettingObject2Option) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypeGetResponseSettingObject3 struct {
	Key   string `json:"key" api:"required"`
	Label string `json:"label" api:"required"`
	// Any of "Alert", "GenericOauth", "Secret", "Select", "Switch", "Text".
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
func (r DestinationTypeGetResponseSettingObject3) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypeGetResponseSettingObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypeGetResponseSettingObject4 struct {
	Key   string `json:"key" api:"required"`
	Label string `json:"label" api:"required"`
	// Any of "Alert", "GenericOauth", "Secret", "Select", "Switch", "Text".
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
func (r DestinationTypeGetResponseSettingObject4) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypeGetResponseSettingObject4) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypeGetResponseSettingObject5 struct {
	Key         string `json:"key" api:"required"`
	Label       string `json:"label" api:"required"`
	Placeholder string `json:"placeholder" api:"required"`
	// Any of "Alert", "GenericOauth", "Secret", "Select", "Switch", "Text".
	Type     string `json:"type" api:"required"`
	Required bool   `json:"required" api:"nullable"`
	Sublabel string `json:"sublabel" api:"nullable"`
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
func (r DestinationTypeGetResponseSettingObject5) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypeGetResponseSettingObject5) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypeGetResponseStatus string

const (
	DestinationTypeGetResponseStatusDeprecated DestinationTypeGetResponseStatus = "deprecated"
	DestinationTypeGetResponseStatusGa         DestinationTypeGetResponseStatus = "ga"
)

type DestinationTypeGetParamsID string

const (
	DestinationTypeGetParamsIDAwsEventBridge               DestinationTypeGetParamsID = "AWSEventBridge"
	DestinationTypeGetParamsIDAwsKinesis                   DestinationTypeGetParamsID = "AWSKinesis"
	DestinationTypeGetParamsIDAwsLambda                    DestinationTypeGetParamsID = "AWSLambda"
	DestinationTypeGetParamsIDAwss3                        DestinationTypeGetParamsID = "AWSS3"
	DestinationTypeGetParamsIDAwssns                       DestinationTypeGetParamsID = "AWSSNS"
	DestinationTypeGetParamsIDActiveCampaignAPI            DestinationTypeGetParamsID = "ActiveCampaignApi"
	DestinationTypeGetParamsIDAdmitad                      DestinationTypeGetParamsID = "Admitad"
	DestinationTypeGetParamsIDAmazonDsp                    DestinationTypeGetParamsID = "AmazonDSP"
	DestinationTypeGetParamsIDAmplitude                    DestinationTypeGetParamsID = "Amplitude"
	DestinationTypeGetParamsIDAppLovin                     DestinationTypeGetParamsID = "AppLovin"
	DestinationTypeGetParamsIDArtsAI                       DestinationTypeGetParamsID = "ArtsAI"
	DestinationTypeGetParamsIDAttentive                    DestinationTypeGetParamsID = "Attentive"
	DestinationTypeGetParamsIDAudiohook                    DestinationTypeGetParamsID = "Audiohook"
	DestinationTypeGetParamsIDAzureBlob                    DestinationTypeGetParamsID = "AzureBlob"
	DestinationTypeGetParamsIDBasisPostback                DestinationTypeGetParamsID = "BasisPostback"
	DestinationTypeGetParamsIDBeeswaxPostback              DestinationTypeGetParamsID = "BeeswaxPostback"
	DestinationTypeGetParamsIDBingAds                      DestinationTypeGetParamsID = "BingAds"
	DestinationTypeGetParamsIDBingAdsWeb                   DestinationTypeGetParamsID = "BingAdsWeb"
	DestinationTypeGetParamsIDBraze                        DestinationTypeGetParamsID = "Braze"
	DestinationTypeGetParamsIDConvertAbTestingEvent        DestinationTypeGetParamsID = "ConvertABTestingEvent"
	DestinationTypeGetParamsIDCustomerio                   DestinationTypeGetParamsID = "Customerio"
	DestinationTypeGetParamsIDDomoWarehouse                DestinationTypeGetParamsID = "DomoWarehouse"
	DestinationTypeGetParamsIDEverflow                     DestinationTypeGetParamsID = "Everflow"
	DestinationTypeGetParamsIDFacebook                     DestinationTypeGetParamsID = "Facebook"
	DestinationTypeGetParamsIDFloodlightSgtm               DestinationTypeGetParamsID = "FloodlightSGTM"
	DestinationTypeGetParamsIDFullContact                  DestinationTypeGetParamsID = "FullContact"
	DestinationTypeGetParamsIDG4Analytics                  DestinationTypeGetParamsID = "G4Analytics"
	DestinationTypeGetParamsIDGa4MeasurementProtocol       DestinationTypeGetParamsID = "GA4MeasurementProtocol"
	DestinationTypeGetParamsIDGa4ServerProxy               DestinationTypeGetParamsID = "GA4ServerProxy"
	DestinationTypeGetParamsIDGoogle                       DestinationTypeGetParamsID = "Google"
	DestinationTypeGetParamsIDGoogleAds360                 DestinationTypeGetParamsID = "GoogleAds360"
	DestinationTypeGetParamsIDGoogleAdsServerContainer     DestinationTypeGetParamsID = "GoogleAdsServerContainer"
	DestinationTypeGetParamsIDGoogleBigQuery               DestinationTypeGetParamsID = "GoogleBigQuery"
	DestinationTypeGetParamsIDGoogleBigQueryWarehouse      DestinationTypeGetParamsID = "GoogleBigQueryWarehouse"
	DestinationTypeGetParamsIDGoogleDataManagerEventIngest DestinationTypeGetParamsID = "GoogleDataManagerEventIngest"
	DestinationTypeGetParamsIDGooglePubSub                 DestinationTypeGetParamsID = "GooglePubSub"
	DestinationTypeGetParamsIDGoogleStorage                DestinationTypeGetParamsID = "GoogleStorage"
	DestinationTypeGetParamsIDHTTPCustomRequest            DestinationTypeGetParamsID = "HTTPCustomRequest"
	DestinationTypeGetParamsIDHTTPDestination              DestinationTypeGetParamsID = "HTTPDestination"
	DestinationTypeGetParamsIDHubspot                      DestinationTypeGetParamsID = "Hubspot"
	DestinationTypeGetParamsIDIHeartMediaMagellan          DestinationTypeGetParamsID = "IHeartMediaMagellan"
	DestinationTypeGetParamsIDImpact                       DestinationTypeGetParamsID = "Impact"
	DestinationTypeGetParamsIDIterable                     DestinationTypeGetParamsID = "Iterable"
	DestinationTypeGetParamsIDKlaviyo                      DestinationTypeGetParamsID = "Klaviyo"
	DestinationTypeGetParamsIDLinkedInAdsCapi              DestinationTypeGetParamsID = "LinkedInAdsCAPI"
	DestinationTypeGetParamsIDLiveIntent                   DestinationTypeGetParamsID = "LiveIntent"
	DestinationTypeGetParamsIDLiveRampWarehouse            DestinationTypeGetParamsID = "LiveRampWarehouse"
	DestinationTypeGetParamsIDMailchimp                    DestinationTypeGetParamsID = "Mailchimp"
	DestinationTypeGetParamsIDMixpanel                     DestinationTypeGetParamsID = "Mixpanel"
	DestinationTypeGetParamsIDNextdoorAds                  DestinationTypeGetParamsID = "NextdoorAds"
	DestinationTypeGetParamsIDOursSyntheticData            DestinationTypeGetParamsID = "OursSyntheticData"
	DestinationTypeGetParamsIDPartnerize                   DestinationTypeGetParamsID = "Partnerize"
	DestinationTypeGetParamsIDPinterest                    DestinationTypeGetParamsID = "Pinterest"
	DestinationTypeGetParamsIDPlausible                    DestinationTypeGetParamsID = "Plausible"
	DestinationTypeGetParamsIDPodscribe                    DestinationTypeGetParamsID = "Podscribe"
	DestinationTypeGetParamsIDPostHog                      DestinationTypeGetParamsID = "PostHog"
	DestinationTypeGetParamsIDQuantcastCapi                DestinationTypeGetParamsID = "QuantcastCAPI"
	DestinationTypeGetParamsIDQuoraAds                     DestinationTypeGetParamsID = "QuoraAds"
	DestinationTypeGetParamsIDReddit                       DestinationTypeGetParamsID = "Reddit"
	DestinationTypeGetParamsIDRokuCapi                     DestinationTypeGetParamsID = "RokuCAPI"
	DestinationTypeGetParamsIDSnapchatAdsCapi              DestinationTypeGetParamsID = "SnapchatAdsCapi"
	DestinationTypeGetParamsIDSpotify                      DestinationTypeGetParamsID = "Spotify"
	DestinationTypeGetParamsIDStackAdaptAPI                DestinationTypeGetParamsID = "StackAdaptAPI"
	DestinationTypeGetParamsIDTaboola                      DestinationTypeGetParamsID = "Taboola"
	DestinationTypeGetParamsIDTatari                       DestinationTypeGetParamsID = "Tatari"
	DestinationTypeGetParamsIDTheTradeDesk                 DestinationTypeGetParamsID = "TheTradeDesk"
	DestinationTypeGetParamsIDTikTok                       DestinationTypeGetParamsID = "TikTok"
	DestinationTypeGetParamsIDVwo                          DestinationTypeGetParamsID = "VWO"
	DestinationTypeGetParamsIDViant                        DestinationTypeGetParamsID = "Viant"
	DestinationTypeGetParamsIDVibe                         DestinationTypeGetParamsID = "Vibe"
	DestinationTypeGetParamsIDWoopra                       DestinationTypeGetParamsID = "Woopra"
	DestinationTypeGetParamsIDXAds                         DestinationTypeGetParamsID = "XAds"
	DestinationTypeGetParamsIDZendesk                      DestinationTypeGetParamsID = "Zendesk"
	DestinationTypeGetParamsIDZoomInfo                     DestinationTypeGetParamsID = "ZoomInfo"
)
