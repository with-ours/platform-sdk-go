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
	options []option.RequestOption
}

// NewDestinationTypeService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDestinationTypeService(opts ...option.RequestOption) (r DestinationTypeService) {
	r = DestinationTypeService{}
	r.options = opts
	return
}

// Lists every destination type the platform supports, with its human-readable
// label, capability flags (oauth, listsAccounts, supportsRenamedEvents), and the
// settings descriptor used to configure a destination of that type.
// Account-agnostic — the response is the same for every API key. Requires scope:
// destination:list
func (r *DestinationTypeService) List(ctx context.Context, opts ...option.RequestOption) (res *DestinationTypeListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "rest/v1/destination-types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Fetch the descriptor for a single destination type by its identifier (e.g.
// `Klaviyo`, `Facebook`, `GoogleDataManagerEventIngest`). Returns 404 if the
// identifier is unknown. Requires scope: destination:list
func (r *DestinationTypeService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *DestinationTypeGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/destination-types/%s", url.PathEscape(id))
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
	// Any of "Audiohook", "BasisPostback", "OursSyntheticData", "FullContact",
	// "ZoomInfo", "TheTradeDesk", "Braze", "LiveIntent", "ConvertABTestingEvent",
	// "Customerio", "BingAds", "BingAdsWeb", "HTTPDestination", "Woopra",
	// "HTTPCustomRequest", "Google", "GoogleAdsServerContainer", "G4Analytics",
	// "GA4ServerProxy", "GA4MeasurementProtocol", "GoogleAds360", "Facebook",
	// "Mixpanel", "Amplitude", "TikTok", "Reddit", "Podscribe", "Pinterest",
	// "Mailchimp", "AWSKinesis", "AWSLambda", "GooglePubSub", "LinkedInAdsCAPI",
	// "ActiveCampaignApi", "StackAdaptAPI", "Hubspot", "Klaviyo", "XAds", "QuoraAds",
	// "SnapchatAdsCapi", "Partnerize", "NextdoorAds", "Tatari", "Viant", "Impact",
	// "Spotify", "Taboola", "AmazonDSP", "AppLovin", "IHeartMediaMagellan", "Vibe",
	// "GoogleDataManagerEventIngest", "Zendesk", "Iterable", "ArtsAI",
	// "QuantcastCAPI", "FloodlightSGTM", "VWO", "Attentive", "Admitad", "Plausible",
	// "PostHog", "RokuCAPI", "Everflow", "BeeswaxPostback", "AdobeAnalytics".
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
func (r DestinationTypeListResponseEntitySettingObject) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypeListResponseEntitySettingObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypeListResponseEntitySettingObject2 struct {
	Key     string                                                  `json:"key" api:"required"`
	Label   string                                                  `json:"label" api:"required"`
	Options []DestinationTypeListResponseEntitySettingObject2Option `json:"options" api:"required"`
	// Any of "Select".
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
func (r DestinationTypeListResponseEntitySettingObject3) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypeListResponseEntitySettingObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypeListResponseEntitySettingObject4 struct {
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
func (r DestinationTypeListResponseEntitySettingObject4) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypeListResponseEntitySettingObject4) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypeListResponseEntitySettingObject5 struct {
	Key         string `json:"key" api:"required"`
	Label       string `json:"label" api:"required"`
	Placeholder string `json:"placeholder" api:"required"`
	// Any of "Text", "Secret".
	Type     DestinationTypeListResponseEntitySettingObject5Type `json:"type" api:"required"`
	Required bool                                                `json:"required" api:"nullable"`
	Sublabel string                                              `json:"sublabel" api:"nullable"`
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

type DestinationTypeListResponseEntitySettingObject5Type string

const (
	DestinationTypeListResponseEntitySettingObject5TypeText   DestinationTypeListResponseEntitySettingObject5Type = "Text"
	DestinationTypeListResponseEntitySettingObject5TypeSecret DestinationTypeListResponseEntitySettingObject5Type = "Secret"
)

type DestinationTypeGetResponse struct {
	// Any of "Audiohook", "BasisPostback", "OursSyntheticData", "FullContact",
	// "ZoomInfo", "TheTradeDesk", "Braze", "LiveIntent", "ConvertABTestingEvent",
	// "Customerio", "BingAds", "BingAdsWeb", "HTTPDestination", "Woopra",
	// "HTTPCustomRequest", "Google", "GoogleAdsServerContainer", "G4Analytics",
	// "GA4ServerProxy", "GA4MeasurementProtocol", "GoogleAds360", "Facebook",
	// "Mixpanel", "Amplitude", "TikTok", "Reddit", "Podscribe", "Pinterest",
	// "Mailchimp", "AWSKinesis", "AWSLambda", "GooglePubSub", "LinkedInAdsCAPI",
	// "ActiveCampaignApi", "StackAdaptAPI", "Hubspot", "Klaviyo", "XAds", "QuoraAds",
	// "SnapchatAdsCapi", "Partnerize", "NextdoorAds", "Tatari", "Viant", "Impact",
	// "Spotify", "Taboola", "AmazonDSP", "AppLovin", "IHeartMediaMagellan", "Vibe",
	// "GoogleDataManagerEventIngest", "Zendesk", "Iterable", "ArtsAI",
	// "QuantcastCAPI", "FloodlightSGTM", "VWO", "Attentive", "Admitad", "Plausible",
	// "PostHog", "RokuCAPI", "Everflow", "BeeswaxPostback", "AdobeAnalytics".
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
	DestinationTypeGetResponseIDAudiohook                    DestinationTypeGetResponseID = "Audiohook"
	DestinationTypeGetResponseIDBasisPostback                DestinationTypeGetResponseID = "BasisPostback"
	DestinationTypeGetResponseIDOursSyntheticData            DestinationTypeGetResponseID = "OursSyntheticData"
	DestinationTypeGetResponseIDFullContact                  DestinationTypeGetResponseID = "FullContact"
	DestinationTypeGetResponseIDZoomInfo                     DestinationTypeGetResponseID = "ZoomInfo"
	DestinationTypeGetResponseIDTheTradeDesk                 DestinationTypeGetResponseID = "TheTradeDesk"
	DestinationTypeGetResponseIDBraze                        DestinationTypeGetResponseID = "Braze"
	DestinationTypeGetResponseIDLiveIntent                   DestinationTypeGetResponseID = "LiveIntent"
	DestinationTypeGetResponseIDConvertAbTestingEvent        DestinationTypeGetResponseID = "ConvertABTestingEvent"
	DestinationTypeGetResponseIDCustomerio                   DestinationTypeGetResponseID = "Customerio"
	DestinationTypeGetResponseIDBingAds                      DestinationTypeGetResponseID = "BingAds"
	DestinationTypeGetResponseIDBingAdsWeb                   DestinationTypeGetResponseID = "BingAdsWeb"
	DestinationTypeGetResponseIDHTTPDestination              DestinationTypeGetResponseID = "HTTPDestination"
	DestinationTypeGetResponseIDWoopra                       DestinationTypeGetResponseID = "Woopra"
	DestinationTypeGetResponseIDHTTPCustomRequest            DestinationTypeGetResponseID = "HTTPCustomRequest"
	DestinationTypeGetResponseIDGoogle                       DestinationTypeGetResponseID = "Google"
	DestinationTypeGetResponseIDGoogleAdsServerContainer     DestinationTypeGetResponseID = "GoogleAdsServerContainer"
	DestinationTypeGetResponseIDG4Analytics                  DestinationTypeGetResponseID = "G4Analytics"
	DestinationTypeGetResponseIDGa4ServerProxy               DestinationTypeGetResponseID = "GA4ServerProxy"
	DestinationTypeGetResponseIDGa4MeasurementProtocol       DestinationTypeGetResponseID = "GA4MeasurementProtocol"
	DestinationTypeGetResponseIDGoogleAds360                 DestinationTypeGetResponseID = "GoogleAds360"
	DestinationTypeGetResponseIDFacebook                     DestinationTypeGetResponseID = "Facebook"
	DestinationTypeGetResponseIDMixpanel                     DestinationTypeGetResponseID = "Mixpanel"
	DestinationTypeGetResponseIDAmplitude                    DestinationTypeGetResponseID = "Amplitude"
	DestinationTypeGetResponseIDTikTok                       DestinationTypeGetResponseID = "TikTok"
	DestinationTypeGetResponseIDReddit                       DestinationTypeGetResponseID = "Reddit"
	DestinationTypeGetResponseIDPodscribe                    DestinationTypeGetResponseID = "Podscribe"
	DestinationTypeGetResponseIDPinterest                    DestinationTypeGetResponseID = "Pinterest"
	DestinationTypeGetResponseIDMailchimp                    DestinationTypeGetResponseID = "Mailchimp"
	DestinationTypeGetResponseIDAwsKinesis                   DestinationTypeGetResponseID = "AWSKinesis"
	DestinationTypeGetResponseIDAwsLambda                    DestinationTypeGetResponseID = "AWSLambda"
	DestinationTypeGetResponseIDGooglePubSub                 DestinationTypeGetResponseID = "GooglePubSub"
	DestinationTypeGetResponseIDLinkedInAdsCapi              DestinationTypeGetResponseID = "LinkedInAdsCAPI"
	DestinationTypeGetResponseIDActiveCampaignAPI            DestinationTypeGetResponseID = "ActiveCampaignApi"
	DestinationTypeGetResponseIDStackAdaptAPI                DestinationTypeGetResponseID = "StackAdaptAPI"
	DestinationTypeGetResponseIDHubspot                      DestinationTypeGetResponseID = "Hubspot"
	DestinationTypeGetResponseIDKlaviyo                      DestinationTypeGetResponseID = "Klaviyo"
	DestinationTypeGetResponseIDXAds                         DestinationTypeGetResponseID = "XAds"
	DestinationTypeGetResponseIDQuoraAds                     DestinationTypeGetResponseID = "QuoraAds"
	DestinationTypeGetResponseIDSnapchatAdsCapi              DestinationTypeGetResponseID = "SnapchatAdsCapi"
	DestinationTypeGetResponseIDPartnerize                   DestinationTypeGetResponseID = "Partnerize"
	DestinationTypeGetResponseIDNextdoorAds                  DestinationTypeGetResponseID = "NextdoorAds"
	DestinationTypeGetResponseIDTatari                       DestinationTypeGetResponseID = "Tatari"
	DestinationTypeGetResponseIDViant                        DestinationTypeGetResponseID = "Viant"
	DestinationTypeGetResponseIDImpact                       DestinationTypeGetResponseID = "Impact"
	DestinationTypeGetResponseIDSpotify                      DestinationTypeGetResponseID = "Spotify"
	DestinationTypeGetResponseIDTaboola                      DestinationTypeGetResponseID = "Taboola"
	DestinationTypeGetResponseIDAmazonDsp                    DestinationTypeGetResponseID = "AmazonDSP"
	DestinationTypeGetResponseIDAppLovin                     DestinationTypeGetResponseID = "AppLovin"
	DestinationTypeGetResponseIDIHeartMediaMagellan          DestinationTypeGetResponseID = "IHeartMediaMagellan"
	DestinationTypeGetResponseIDVibe                         DestinationTypeGetResponseID = "Vibe"
	DestinationTypeGetResponseIDGoogleDataManagerEventIngest DestinationTypeGetResponseID = "GoogleDataManagerEventIngest"
	DestinationTypeGetResponseIDZendesk                      DestinationTypeGetResponseID = "Zendesk"
	DestinationTypeGetResponseIDIterable                     DestinationTypeGetResponseID = "Iterable"
	DestinationTypeGetResponseIDArtsAI                       DestinationTypeGetResponseID = "ArtsAI"
	DestinationTypeGetResponseIDQuantcastCapi                DestinationTypeGetResponseID = "QuantcastCAPI"
	DestinationTypeGetResponseIDFloodlightSgtm               DestinationTypeGetResponseID = "FloodlightSGTM"
	DestinationTypeGetResponseIDVwo                          DestinationTypeGetResponseID = "VWO"
	DestinationTypeGetResponseIDAttentive                    DestinationTypeGetResponseID = "Attentive"
	DestinationTypeGetResponseIDAdmitad                      DestinationTypeGetResponseID = "Admitad"
	DestinationTypeGetResponseIDPlausible                    DestinationTypeGetResponseID = "Plausible"
	DestinationTypeGetResponseIDPostHog                      DestinationTypeGetResponseID = "PostHog"
	DestinationTypeGetResponseIDRokuCapi                     DestinationTypeGetResponseID = "RokuCAPI"
	DestinationTypeGetResponseIDEverflow                     DestinationTypeGetResponseID = "Everflow"
	DestinationTypeGetResponseIDBeeswaxPostback              DestinationTypeGetResponseID = "BeeswaxPostback"
	DestinationTypeGetResponseIDAdobeAnalytics               DestinationTypeGetResponseID = "AdobeAnalytics"
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
func (r DestinationTypeGetResponseSettingObject) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypeGetResponseSettingObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypeGetResponseSettingObject2 struct {
	Key     string                                           `json:"key" api:"required"`
	Label   string                                           `json:"label" api:"required"`
	Options []DestinationTypeGetResponseSettingObject2Option `json:"options" api:"required"`
	// Any of "Select".
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
func (r DestinationTypeGetResponseSettingObject3) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypeGetResponseSettingObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypeGetResponseSettingObject4 struct {
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
func (r DestinationTypeGetResponseSettingObject4) RawJSON() string { return r.JSON.raw }
func (r *DestinationTypeGetResponseSettingObject4) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationTypeGetResponseSettingObject5 struct {
	Key         string `json:"key" api:"required"`
	Label       string `json:"label" api:"required"`
	Placeholder string `json:"placeholder" api:"required"`
	// Any of "Text", "Secret".
	Type     DestinationTypeGetResponseSettingObject5Type `json:"type" api:"required"`
	Required bool                                         `json:"required" api:"nullable"`
	Sublabel string                                       `json:"sublabel" api:"nullable"`
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

type DestinationTypeGetResponseSettingObject5Type string

const (
	DestinationTypeGetResponseSettingObject5TypeText   DestinationTypeGetResponseSettingObject5Type = "Text"
	DestinationTypeGetResponseSettingObject5TypeSecret DestinationTypeGetResponseSettingObject5Type = "Secret"
)

type DestinationTypeGetResponseStatus string

const (
	DestinationTypeGetResponseStatusDeprecated DestinationTypeGetResponseStatus = "deprecated"
	DestinationTypeGetResponseStatusGa         DestinationTypeGetResponseStatus = "ga"
)
