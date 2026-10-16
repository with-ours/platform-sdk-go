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
// API-key scope or current OAuth user permission: report:list-events
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
// API-key scope or current OAuth user permission: report:list-events
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
// they are generated and returned. Requires API-key scope or current OAuth user
// permission: test-event:create
func (r *TestEventService) New(ctx context.Context, body TestEventNewParams, opts ...option.RequestOption) (res *TestEventNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/test-events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve one browser, server, or synthetic event captured in debug mode within
// the last 48 hours. Use an id returned by the create or list endpoint. Events
// that have not arrived, expired events, and events outside your account
// return 404. Requires API-key scope or current OAuth user permission:
// report:view-event
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
// Requires API-key scope or current OAuth user permission: report:view-dispatch
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
	IsTestEvent       bool                                                 `json:"isTestEvent" api:"required"`
	Time              string                                               `json:"time" api:"required"`
	VisitorID         string                                               `json:"visitorId" api:"required"`
	DefaultProperties map[string]TestEventListResponseDefaultPropertyUnion `json:"defaultProperties" api:"nullable"`
	EventProperties   map[string]TestEventListResponseEventPropertyUnion   `json:"eventProperties" api:"nullable"`
	RawData           string                                               `json:"rawData" api:"nullable"`
	RequestContext    map[string]TestEventListResponseRequestContextUnion  `json:"requestContext" api:"nullable"`
	SourceID          string                                               `json:"sourceId" api:"nullable"`
	SourceName        string                                               `json:"sourceName" api:"nullable"`
	SourceType        string                                               `json:"sourceType" api:"nullable"`
	UserProperties    map[string]TestEventListResponseUserPropertyUnion    `json:"userProperties" api:"nullable"`
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

// TestEventListResponseDefaultPropertyUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfTestEventListResponseDefaultPropertyMapItem]
type TestEventListResponseDefaultPropertyUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfTestEventListResponseDefaultPropertyMapItem any `json:",inline"`
	JSON                                          struct {
		OfString                                      respjson.Field
		OfFloat                                       respjson.Field
		OfBool                                        respjson.Field
		OfAnyArray                                    respjson.Field
		OfTestEventListResponseDefaultPropertyMapItem respjson.Field
		raw                                           string
	} `json:"-"`
}

func (u TestEventListResponseDefaultPropertyUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventListResponseDefaultPropertyUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventListResponseDefaultPropertyUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventListResponseDefaultPropertyUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventListResponseDefaultPropertyUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TestEventListResponseDefaultPropertyUnion) RawJSON() string { return u.JSON.raw }

func (r *TestEventListResponseDefaultPropertyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TestEventListResponseEventPropertyUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfTestEventListResponseEventPropertyMapItem]
type TestEventListResponseEventPropertyUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfTestEventListResponseEventPropertyMapItem any `json:",inline"`
	JSON                                        struct {
		OfString                                    respjson.Field
		OfFloat                                     respjson.Field
		OfBool                                      respjson.Field
		OfAnyArray                                  respjson.Field
		OfTestEventListResponseEventPropertyMapItem respjson.Field
		raw                                         string
	} `json:"-"`
}

func (u TestEventListResponseEventPropertyUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventListResponseEventPropertyUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventListResponseEventPropertyUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventListResponseEventPropertyUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventListResponseEventPropertyUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TestEventListResponseEventPropertyUnion) RawJSON() string { return u.JSON.raw }

func (r *TestEventListResponseEventPropertyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TestEventListResponseRequestContextUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfTestEventListResponseRequestContextMapItem]
type TestEventListResponseRequestContextUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfTestEventListResponseRequestContextMapItem any `json:",inline"`
	JSON                                         struct {
		OfString                                     respjson.Field
		OfFloat                                      respjson.Field
		OfBool                                       respjson.Field
		OfAnyArray                                   respjson.Field
		OfTestEventListResponseRequestContextMapItem respjson.Field
		raw                                          string
	} `json:"-"`
}

func (u TestEventListResponseRequestContextUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventListResponseRequestContextUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventListResponseRequestContextUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventListResponseRequestContextUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventListResponseRequestContextUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TestEventListResponseRequestContextUnion) RawJSON() string { return u.JSON.raw }

func (r *TestEventListResponseRequestContextUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TestEventListResponseUserPropertyUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfTestEventListResponseUserPropertyMapItem]
type TestEventListResponseUserPropertyUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfTestEventListResponseUserPropertyMapItem any `json:",inline"`
	JSON                                       struct {
		OfString                                   respjson.Field
		OfFloat                                    respjson.Field
		OfBool                                     respjson.Field
		OfAnyArray                                 respjson.Field
		OfTestEventListResponseUserPropertyMapItem respjson.Field
		raw                                        string
	} `json:"-"`
}

func (u TestEventListResponseUserPropertyUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventListResponseUserPropertyUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventListResponseUserPropertyUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventListResponseUserPropertyUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventListResponseUserPropertyUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TestEventListResponseUserPropertyUnion) RawJSON() string { return u.JSON.raw }

func (r *TestEventListResponseUserPropertyUnion) UnmarshalJSON(data []byte) error {
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
	IsTestEvent       bool                                                `json:"isTestEvent" api:"required"`
	Time              string                                              `json:"time" api:"required"`
	VisitorID         string                                              `json:"visitorId" api:"required"`
	DefaultProperties map[string]TestEventGetResponseDefaultPropertyUnion `json:"defaultProperties" api:"nullable"`
	EventProperties   map[string]TestEventGetResponseEventPropertyUnion   `json:"eventProperties" api:"nullable"`
	RawData           string                                              `json:"rawData" api:"nullable"`
	RequestContext    map[string]TestEventGetResponseRequestContextUnion  `json:"requestContext" api:"nullable"`
	SourceID          string                                              `json:"sourceId" api:"nullable"`
	SourceName        string                                              `json:"sourceName" api:"nullable"`
	SourceType        string                                              `json:"sourceType" api:"nullable"`
	UserProperties    map[string]TestEventGetResponseUserPropertyUnion    `json:"userProperties" api:"nullable"`
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

// TestEventGetResponseDefaultPropertyUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfTestEventGetResponseDefaultPropertyMapItem]
type TestEventGetResponseDefaultPropertyUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfTestEventGetResponseDefaultPropertyMapItem any `json:",inline"`
	JSON                                         struct {
		OfString                                     respjson.Field
		OfFloat                                      respjson.Field
		OfBool                                       respjson.Field
		OfAnyArray                                   respjson.Field
		OfTestEventGetResponseDefaultPropertyMapItem respjson.Field
		raw                                          string
	} `json:"-"`
}

func (u TestEventGetResponseDefaultPropertyUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventGetResponseDefaultPropertyUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventGetResponseDefaultPropertyUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventGetResponseDefaultPropertyUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventGetResponseDefaultPropertyUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TestEventGetResponseDefaultPropertyUnion) RawJSON() string { return u.JSON.raw }

func (r *TestEventGetResponseDefaultPropertyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TestEventGetResponseEventPropertyUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfTestEventGetResponseEventPropertyMapItem]
type TestEventGetResponseEventPropertyUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfTestEventGetResponseEventPropertyMapItem any `json:",inline"`
	JSON                                       struct {
		OfString                                   respjson.Field
		OfFloat                                    respjson.Field
		OfBool                                     respjson.Field
		OfAnyArray                                 respjson.Field
		OfTestEventGetResponseEventPropertyMapItem respjson.Field
		raw                                        string
	} `json:"-"`
}

func (u TestEventGetResponseEventPropertyUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventGetResponseEventPropertyUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventGetResponseEventPropertyUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventGetResponseEventPropertyUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventGetResponseEventPropertyUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TestEventGetResponseEventPropertyUnion) RawJSON() string { return u.JSON.raw }

func (r *TestEventGetResponseEventPropertyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TestEventGetResponseRequestContextUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfTestEventGetResponseRequestContextMapItem]
type TestEventGetResponseRequestContextUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfTestEventGetResponseRequestContextMapItem any `json:",inline"`
	JSON                                        struct {
		OfString                                    respjson.Field
		OfFloat                                     respjson.Field
		OfBool                                      respjson.Field
		OfAnyArray                                  respjson.Field
		OfTestEventGetResponseRequestContextMapItem respjson.Field
		raw                                         string
	} `json:"-"`
}

func (u TestEventGetResponseRequestContextUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventGetResponseRequestContextUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventGetResponseRequestContextUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventGetResponseRequestContextUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventGetResponseRequestContextUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TestEventGetResponseRequestContextUnion) RawJSON() string { return u.JSON.raw }

func (r *TestEventGetResponseRequestContextUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TestEventGetResponseUserPropertyUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfTestEventGetResponseUserPropertyMapItem]
type TestEventGetResponseUserPropertyUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfTestEventGetResponseUserPropertyMapItem any `json:",inline"`
	JSON                                      struct {
		OfString                                  respjson.Field
		OfFloat                                   respjson.Field
		OfBool                                    respjson.Field
		OfAnyArray                                respjson.Field
		OfTestEventGetResponseUserPropertyMapItem respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u TestEventGetResponseUserPropertyUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventGetResponseUserPropertyUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventGetResponseUserPropertyUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventGetResponseUserPropertyUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventGetResponseUserPropertyUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TestEventGetResponseUserPropertyUnion) RawJSON() string { return u.JSON.raw }

func (r *TestEventGetResponseUserPropertyUnion) UnmarshalJSON(data []byte) error {
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
	Payload map[string]TestEventDispatchesResponseEntityPayloadUnion `json:"payload" api:"nullable"`
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

// TestEventDispatchesResponseEntityPayloadUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfTestEventDispatchesResponseEntityPayloadMapItem]
type TestEventDispatchesResponseEntityPayloadUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfTestEventDispatchesResponseEntityPayloadMapItem any `json:",inline"`
	JSON                                              struct {
		OfString                                          respjson.Field
		OfFloat                                           respjson.Field
		OfBool                                            respjson.Field
		OfAnyArray                                        respjson.Field
		OfTestEventDispatchesResponseEntityPayloadMapItem respjson.Field
		raw                                               string
	} `json:"-"`
}

func (u TestEventDispatchesResponseEntityPayloadUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventDispatchesResponseEntityPayloadUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventDispatchesResponseEntityPayloadUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventDispatchesResponseEntityPayloadUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TestEventDispatchesResponseEntityPayloadUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TestEventDispatchesResponseEntityPayloadUnion) RawJSON() string { return u.JSON.raw }

func (r *TestEventDispatchesResponseEntityPayloadUnion) UnmarshalJSON(data []byte) error {
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
	SourceID param.Opt[string] `json:"sourceId,omitzero" format:"uuid"`
	// Visitor identity to send the event as. Generated when omitted. Letters, numbers,
	// hyphens, and underscores only.
	VisitorID param.Opt[string] `json:"visitorId,omitzero"`
	// Known context properties used by ingest and destination mappings. Use
	// current_url for the page URL. Unknown properties are ignored.
	DefaultProperties any `json:"defaultProperties,omitzero"`
	// Event-level properties available to destination mappings.
	EventProperties map[string]TestEventNewParamsEventPropertyUnion `json:"eventProperties,omitzero"`
	// Person-level properties available to destination mappings.
	UserProperties map[string]TestEventNewParamsUserPropertyUnion `json:"userProperties,omitzero"`
	paramObj
}

func (r TestEventNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TestEventNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TestEventNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TestEventNewParamsEventPropertyUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u TestEventNewParamsEventPropertyUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *TestEventNewParamsEventPropertyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TestEventNewParamsUserPropertyUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u TestEventNewParamsUserPropertyUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *TestEventNewParamsUserPropertyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}
