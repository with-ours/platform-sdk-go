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

// TestEventService contains methods and other services that help with interacting
// with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTestEventService] method instead.
type TestEventService struct {
	Options []option.RequestOption
}

// NewTestEventService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTestEventService(opts ...option.RequestOption) (r TestEventService) {
	r = TestEventService{}
	r.Options = opts
	return
}

// List browser, server, and synthetic events captured with debug mode during the
// last 48 hours, newest received first. Includes live-token and test-token events;
// `isTestEvent` indicates whether a test token was used. Returns full event
// properties. Pass an entity id unchanged to the detail or dispatch endpoint. Use
// `limit` (default 25, maximum 100) and `cursor` to page through results. Requires
// scope: report:list-events
func (r *TestEventService) List(ctx context.Context, query TestEventListParams, opts ...option.RequestOption) (res *pagination.Cursor[TestEventListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "rest/v1/test-events"
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

// List browser, server, and synthetic events captured with debug mode during the
// last 48 hours, newest received first. Includes live-token and test-token events;
// `isTestEvent` indicates whether a test token was used. Returns full event
// properties. Pass an entity id unchanged to the detail or dispatch endpoint. Use
// `limit` (default 25, maximum 100) and `cursor` to page through results. Requires
// scope: report:list-events
func (r *TestEventService) ListAutoPaging(ctx context.Context, query TestEventListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[TestEventListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Send a test event through the same processing your production events go through,
// using the source’s live token with debug capture enabled. This can deliver real
// data to configured destinations. The response acknowledges acceptance for
// processing; it does not contain processing results. The event is captured in
// Recent Events. Use the returned `id` with
// `GET /rest/v1/test-events/{id}/dispatches` to retrieve recorded dispatches.
// Supply `visitorId` and `distinctId` to choose the identity yourself; otherwise
// they are generated and returned. Requires scope: test-event:create
func (r *TestEventService) New(ctx context.Context, body TestEventNewParams, opts ...option.RequestOption) (res *TestEventNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/test-events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve one browser, server, or synthetic event captured in debug mode within
// the last 48 hours. Use an id returned by the create or list endpoint. Events
// that have not arrived, expired events, and events outside your account
// return 404. Requires scope: report:view-event
func (r *TestEventService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *TestEventGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/test-events/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve dispatches recorded for a debug-captured browser, server, or synthetic
// event. Use an id returned by the create or list endpoint. The entities array
// contains mapped payloads and recorded success or failure details. It is empty
// until dispatch records arrive, and may remain empty when no dispatch occurs.
// Results can grow as processing continues; this endpoint does not predict
// destinations or signal completion. Requires the report:view-dispatch scope.
// Requires scope: report:view-dispatch
func (r *TestEventService) Dispatches(ctx context.Context, id string, opts ...option.RequestOption) (res *TestEventDispatchesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/test-events/%s/dispatches", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type TestEventListResponse struct {
	ID         string `json:"id" api:"required"`
	DistinctID string `json:"distinctId" api:"required"`
	EventName  string `json:"eventName" api:"required"`
	// Whether the event used a test token. Debug-captured live-token events have false
	// here.
	IsTestEvent       bool   `json:"isTestEvent" api:"required"`
	Time              string `json:"time" api:"required"`
	VisitorID         string `json:"visitorId" api:"required"`
	DefaultProperties any    `json:"defaultProperties" api:"nullable"`
	EventProperties   any    `json:"eventProperties" api:"nullable"`
	RawData           string `json:"rawData" api:"nullable"`
	RequestContext    any    `json:"requestContext" api:"nullable"`
	SourceID          string `json:"sourceId" api:"nullable"`
	SourceName        string `json:"sourceName" api:"nullable"`
	SourceType        string `json:"sourceType" api:"nullable"`
	UserProperties    any    `json:"userProperties" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		DistinctID        respjson.Field
		EventName         respjson.Field
		IsTestEvent       respjson.Field
		Time              respjson.Field
		VisitorID         respjson.Field
		DefaultProperties respjson.Field
		EventProperties   respjson.Field
		RawData           respjson.Field
		RequestContext    respjson.Field
		SourceID          respjson.Field
		SourceName        respjson.Field
		SourceType        respjson.Field
		UserProperties    respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TestEventListResponse) RawJSON() string { return r.JSON.raw }
func (r *TestEventListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TestEventNewResponse struct {
	// Identifier to pass to the event detail or dispatch endpoint.
	ID string `json:"id" api:"required"`
	// When the event was accepted for processing, as an ISO 8601 timestamp.
	AcceptedAt string `json:"acceptedAt" api:"required"`
	DistinctID string `json:"distinctId" api:"required"`
	EventName  string `json:"eventName" api:"required"`
	SourceID   string `json:"sourceId" api:"required"`
	VisitorID  string `json:"visitorId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AcceptedAt  respjson.Field
		DistinctID  respjson.Field
		EventName   respjson.Field
		SourceID    respjson.Field
		VisitorID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TestEventNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TestEventNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TestEventGetResponse struct {
	ID         string `json:"id" api:"required"`
	DistinctID string `json:"distinctId" api:"required"`
	EventName  string `json:"eventName" api:"required"`
	// Whether the event used a test token. Debug-captured live-token events have false
	// here.
	IsTestEvent       bool   `json:"isTestEvent" api:"required"`
	Time              string `json:"time" api:"required"`
	VisitorID         string `json:"visitorId" api:"required"`
	DefaultProperties any    `json:"defaultProperties" api:"nullable"`
	EventProperties   any    `json:"eventProperties" api:"nullable"`
	RawData           string `json:"rawData" api:"nullable"`
	RequestContext    any    `json:"requestContext" api:"nullable"`
	SourceID          string `json:"sourceId" api:"nullable"`
	SourceName        string `json:"sourceName" api:"nullable"`
	SourceType        string `json:"sourceType" api:"nullable"`
	UserProperties    any    `json:"userProperties" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		DistinctID        respjson.Field
		EventName         respjson.Field
		IsTestEvent       respjson.Field
		Time              respjson.Field
		VisitorID         respjson.Field
		DefaultProperties respjson.Field
		EventProperties   respjson.Field
		RawData           respjson.Field
		RequestContext    respjson.Field
		SourceID          respjson.Field
		SourceName        respjson.Field
		SourceType        respjson.Field
		UserProperties    respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TestEventGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TestEventGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TestEventDispatchesResponse struct {
	ID string `json:"id" api:"required"`
	// Dispatches recorded so far. Empty until dispatches arrive; no predicted
	// destinations or completion signal. Events created through this API can be
	// delivered to live destinations.
	Entities []TestEventDispatchesResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TestEventDispatchesResponse) RawJSON() string { return r.JSON.raw }
func (r *TestEventDispatchesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TestEventDispatchesResponseEntity struct {
	DestinationID    string `json:"destinationId" api:"required"`
	AllowedEventID   string `json:"allowedEventId" api:"nullable"`
	AllowedEventName string `json:"allowedEventName" api:"nullable"`
	DestinationName  string `json:"destinationName" api:"nullable"`
	DestinationType  string `json:"destinationType" api:"nullable"`
	// When the send was attempted, as an ISO 8601 timestamp.
	DispatchedAt string `json:"dispatchedAt" api:"nullable"`
	// When the destination accepted the send, as an ISO 8601 timestamp.
	DispatchedSuccessAt string `json:"dispatchedSuccessAt" api:"nullable"`
	// Failure detail when `success` is false.
	Message string `json:"message" api:"nullable"`
	// The mapped payload recorded for this dispatch.
	Payload any `json:"payload" api:"nullable"`
	// Whether this send succeeded. Null while the attempt is still in flight.
	Success bool `json:"success" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DestinationID       respjson.Field
		AllowedEventID      respjson.Field
		AllowedEventName    respjson.Field
		DestinationName     respjson.Field
		DestinationType     respjson.Field
		DispatchedAt        respjson.Field
		DispatchedSuccessAt respjson.Field
		Message             respjson.Field
		Payload             respjson.Field
		Success             respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TestEventDispatchesResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *TestEventDispatchesResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TestEventListParams struct {
	// Maximum number of items to return. Defaults to 25; values below 1 are clamped to
	// 1 and values above 100 are clamped to 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Opaque pagination cursor from pagination.nextCursor in the previous response. Do
	// not decode or modify it. Malformed cursors return 400 Bad Request.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TestEventListParams]'s query parameters as `url.Values`.
func (r TestEventListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TestEventNewParams struct {
	// Name of the event to send, matching an event name configured on this account.
	EventName string `json:"eventName" api:"required"`
	// Event identity for this specific event. Generated when omitted. Letters,
	// numbers, hyphens, and underscores only.
	DistinctID param.Opt[string] `json:"distinctId,omitzero"`
	// Source to attribute the event to. Defaults to the account's first enabled source
	// that accepts events over the API.
	SourceID param.Opt[string] `json:"sourceId,omitzero"`
	// Visitor identity to send the event as. Generated when omitted. Letters, numbers,
	// hyphens, and underscores only.
	VisitorID param.Opt[string] `json:"visitorId,omitzero"`
	// Known context properties used by ingest and destination mappings. Use
	// current_url for the page URL. Unknown properties are ignored.
	DefaultProperties TestEventNewParamsDefaultProperties `json:"defaultProperties,omitzero"`
	// Event-level properties available to destination mappings.
	EventProperties any `json:"eventProperties,omitzero"`
	// Person-level properties available to destination mappings.
	UserProperties any `json:"userProperties,omitzero"`
	paramObj
}

func (r TestEventNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TestEventNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TestEventNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Known context properties used by ingest and destination mappings. Use
// current_url for the page URL. Unknown properties are ignored.
type TestEventNewParamsDefaultProperties struct {
	// The Everflow affiliate Click (Transaction) ID, captured from the
	// `_ef_transaction_id` URL parameter. Ex: ef_click_abc123
	EfTransactionID param.Opt[string] `json:"_ef_transaction_id,omitzero"`
	// The active time in milliseconds that the user had this tab active
	ActiveDuration param.Opt[float64] `json:"activeDuration,omitzero"`
	// The ad id for detected in the session. This is set by the web sdk automatically.
	AdID param.Opt[string] `json:"ad_id,omitzero"`
	// The Admitad (Mitgo) affiliate Click ID. Ex: admitad_uid_abc123
	AdmitadUid param.Opt[string] `json:"admitad_uid,omitzero"`
	// The adset id for detected in the session. This is set by the web sdk
	// automatically.
	AdsetID param.Opt[string] `json:"adset_id,omitzero"`
	// The AppLovin alart query parameter. Ex: alart123
	Alart param.Opt[string] `json:"alart,omitzero"`
	// The AppLovin aleid query parameter. Ex: aleid123
	Aleid param.Opt[string] `json:"aleid,omitzero"`
	// The AppLovin pixel cookie value (\_axwrt). Web-only.
	Axwrt param.Opt[string] `json:"axwrt,omitzero"`
	// The Basis DSP Click ID. Ex: basis_cid123
	BasisCid param.Opt[string] `json:"basis_cid,omitzero"`
	// The Beeswax (FreeWheel Buyer Cloud) auction ID, captured from the
	// `{{AUCTION_ID}}` macro on creative click URLs. Ex: bx-auc-abc123
	BeeswaxAuctionID param.Opt[string] `json:"beeswax_auction_id,omitzero"`
	// The language of the browser. Ex: en-US
	BrowserLanguage param.Opt[string] `json:"browser_language,omitzero"`
	// The name of the browser. Ex: Chrome
	BrowserName param.Opt[string] `json:"browser_name,omitzero"`
	// The version of the browser. Ex: 114.0
	BrowserVersion param.Opt[string] `json:"browser_version,omitzero"`
	// The campaign id for detected in the session. This is set by the web sdk
	// automatically.
	CampaignID param.Opt[string] `json:"campaign_id,omitzero"`
	// The Click ID. Ex: clickid123
	Clickid param.Opt[string] `json:"clickid,omitzero"`
	// The Generic Click ID. Ex: clid123
	Clid param.Opt[string] `json:"clid,omitzero"`
	// The architecture of the CPU. Ex: x64
	CPUArchitecture param.Opt[string] `json:"cpu_architecture,omitzero"`
	// The full url (including query params) of the current page
	CurrentURL param.Opt[string] `json:"current_url,omitzero"`
	// The DoubleClick Click ID. Ex: dclid123
	Dclid param.Opt[string] `json:"dclid,omitzero"`
	// The model of the device. Ex: iPhone 13
	DeviceModel param.Opt[string] `json:"device_model,omitzero"`
	// The type of device the user is using. Ex: mobile
	DeviceType param.Opt[string] `json:"device_type,omitzero"`
	// The vendor of the device. Ex: Apple
	DeviceVendor param.Opt[string] `json:"device_vendor,omitzero"`
	// The time in milliseconds since the page was loaded // script was loaded
	Duration param.Opt[float64] `json:"duration,omitzero"`
	// The browsers encoding. Ex: UTF-8
	Encoding param.Opt[string] `json:"encoding,omitzero"`
	// The name of the browser engine. Ex: Blink
	EngineName param.Opt[string] `json:"engine_name,omitzero"`
	// The version of the browser engine. Ex: 114.0
	EngineVersion param.Opt[string] `json:"engine_version,omitzero"`
	// The Pinterest Click ID. Ex: epik456
	Epik param.Opt[string] `json:"epik,omitzero"`
	// Facebook Click ID with prefix format for Conversions API tracking. Ex:
	// fb.1.1554763741205.AbCdEfGhIjKlMnOpQrStUvWxYz1234567890
	Fbc param.Opt[string] `json:"fbc,omitzero"`
	// Raw Facebook Click ID query parameter without prefix from ad clicks. Ex:
	// AbCdEfGhIjKlMnOpQrStUvWxYz1234567890
	Fbclid param.Opt[string] `json:"fbclid,omitzero"`
	// Facebook Browser ID parameter for identifying browsers and attributing events.
	// Ex: fb.1.1554763741205.1098115397
	Fbp param.Opt[string] `json:"fbp,omitzero"`
	// Deprecated
	Fv param.Opt[bool] `json:"fv,omitzero"`
	// The Google Ad Source. Ex: google
	GadSource param.Opt[string] `json:"gad_source,omitzero"`
	// The Google Braid ID. Ex: gbraid123
	Gbraid param.Opt[string] `json:"gbraid,omitzero"`
	// The Google Click ID. Ex: gclid123
	Gclid param.Opt[string] `json:"gclid,omitzero"`
	// The host of the current page. Ex: example.com
	Host param.Opt[string] `json:"host,omitzero"`
	// Whether the user is in an iframe. Ex: true
	Iframe param.Opt[bool] `json:"iframe,omitzero"`
	// The Impact Click ID reference. Ex: im_ref123
	ImRef param.Opt[string] `json:"im_ref,omitzero"`
	// The IP address of the user. Ex: 127.0.0.1
	IP param.Opt[string] `json:"ip,omitzero"`
	// The Impact Click ID. Ex: irclickid123
	Irclickid param.Opt[string] `json:"irclickid,omitzero"`
	// Whether we have detected that the user is a bot. This is set automatically by
	// the Ours server primarily for events tracked through the web SDK.
	IsBot param.Opt[bool] `json:"is_bot,omitzero"`
	// The LinkedIn Click ID. Ex: li_fat_id123
	LiFatID param.Opt[string] `json:"li_fat_id,omitzero"`
	// The Microsoft Click ID. Ex: msclkid123
	Msclkid param.Opt[string] `json:"msclkid,omitzero"`
	// The NextDoor Click ID. Ex: ndclid123
	Ndclid param.Opt[string] `json:"ndclid,omitzero"`
	// Deprecated
	NewS param.Opt[bool] `json:"new_s,omitzero"`
	// The Outbrain click ID, captured from the `ob_click_id` URL parameter (Outbrain
	// `{{ob_click_id}}` macro) on the landing page. Ex: ob_click_abc123
	ObClickID param.Opt[string] `json:"ob_click_id,omitzero"`
	// The OpenAI Ads privacy-preserving reference, captured from the `oppref` URL
	// parameter on landing pages (the OpenAI Pixel also stores it in a `__oppref`
	// cookie). Sent to OpenAI Ads on Conversions API events for attribution. Ex:
	// oppref_abc
	Oppref param.Opt[string] `json:"oppref,omitzero"`
	// The name of the operating system. Ex: Windows
	OsName param.Opt[string] `json:"os_name,omitzero"`
	// The version of the operating system. Ex: 10.0
	OsVersion param.Opt[string] `json:"os_version,omitzero"`
	// A random set of numbers for the page load
	PageHash param.Opt[float64] `json:"page_hash,omitzero"`
	// The pathname of the current page. Ex: /home
	Pathname param.Opt[string] `json:"pathname,omitzero"`
	// The Quora Click ID. Ex: qclid123
	Qclid param.Opt[string] `json:"qclid,omitzero"`
	// The Reddit Click ID. Ex: rdt_cid123
	RdtCid param.Opt[string] `json:"rdt_cid,omitzero"`
	// The time the event was received by an Ours server in ISO format
	ReceivedAt param.Opt[string] `json:"received_at,omitzero"`
	// The referrer URL of the current page
	Referrer param.Opt[string] `json:"referrer,omitzero"`
	// The referring domain of the current page
	ReferringDomain param.Opt[string] `json:"referring_domain,omitzero"`
	// The StackAdapt Tracking ID. Ex: sacid123
	Sacid param.Opt[string] `json:"sacid,omitzero"`
	// The SnapChat Click ID. Ex: sccid123
	Sccid param.Opt[string] `json:"sccid,omitzero"`
	// The height of the screen. Ex: 1080
	ScreenHeight param.Opt[float64] `json:"screen_height,omitzero"`
	// The width of the screen. Ex: 1920
	ScreenWidth param.Opt[float64] `json:"screen_width,omitzero"`
	// The number of sessions the user has had. Ex: 3
	SessionCount param.Opt[float64] `json:"sessionCount,omitzero"`
	// The session ID as assigned automatically by the web SDK. This is required for
	// session replay
	Sid param.Opt[string] `json:"sid,omitzero"`
	Sr  param.Opt[string] `json:"sr,omitzero"`
	// The title of the current page
	Title param.Opt[string] `json:"title,omitzero"`
	// The TikTok Click ID. Ex: ttclid123
	Ttclid param.Opt[string] `json:"ttclid,omitzero"`
	// The Twitter Click ID. Ex: twclid123
	Twclid param.Opt[string] `json:"twclid,omitzero"`
	// User agent as a full list of strings.
	Uafvl param.Opt[string] `json:"uafvl,omitzero"`
	// The user agent of the browser
	UserAgent param.Opt[string] `json:"user_agent,omitzero"`
	// The UTM Campaign. The web SDK automatically captures this from the query params.
	UtmCampaign param.Opt[string] `json:"utm_campaign,omitzero"`
	// The UTM Content. The web SDK automatically captures this from the query params.
	UtmContent param.Opt[string] `json:"utm_content,omitzero"`
	// The UTM Medium. The web SDK automatically captures this from the query params.
	UtmMedium param.Opt[string] `json:"utm_medium,omitzero"`
	// The UTM Name. The web SDK automatically captures this from the query params.
	UtmName param.Opt[string] `json:"utm_name,omitzero"`
	// The UTM Source. The web SDK automatically captures this from the query params.
	UtmSource param.Opt[string] `json:"utm_source,omitzero"`
	// The UTM Term. The web SDK automatically captures this from the query params.
	UtmTerm param.Opt[string] `json:"utm_term,omitzero"`
	// The SDK version (e.g., web SDK or ingest-sdk-\* via generated SDK headers)
	Version param.Opt[string] `json:"version,omitzero"`
	// The Viant (Adelphic) Click ID, captured from the `viant_click_id` URL parameter
	// (Viant `${ADELPHIC_CLICKID}` macro). Sent as `xid` on Viant postbacks. Ex:
	// viant_click_abc123
	ViantClickID param.Opt[string] `json:"viant_click_id,omitzero"`
	// The Viant (Adelphic) Impression ID, captured from the `viant_impression_id` URL
	// parameter (Viant `${ADELPHIC_IMPRESSIONID}` macro). Sent as `imp_id` on Viant
	// postbacks for post-view attribution. Ex: viant_imp_abc123
	ViantImpressionID param.Opt[string] `json:"viant_impression_id,omitzero"`
	// The WBRAID Identifier. The web SDK automatically captures this from the query
	// params.
	Wbraid param.Opt[string] `json:"wbraid,omitzero"`
	// Whether the user is in a webview. Ex: true
	Webview param.Opt[bool] `json:"webview,omitzero"`
	paramObj
}

func (r TestEventNewParamsDefaultProperties) MarshalJSON() (data []byte, err error) {
	type shadow TestEventNewParamsDefaultProperties
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TestEventNewParamsDefaultProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
