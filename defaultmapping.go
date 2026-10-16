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
	"github.com/with-ours/platform-sdk-go/v2/internal/requestconfig"
	"github.com/with-ours/platform-sdk-go/v2/option"
	"github.com/with-ours/platform-sdk-go/v2/packages/param"
	"github.com/with-ours/platform-sdk-go/v2/packages/respjson"
)

// DefaultMappingService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDefaultMappingService] method instead.
type DefaultMappingService struct {
	Options []option.RequestOption
}

// NewDefaultMappingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDefaultMappingService(opts ...option.RequestOption) (r DefaultMappingService) {
	r = DefaultMappingService{}
	r.Options = opts
	return
}

// List every stored default mapping for the account, one per destination that has
// ever written a default. Destinations that have not yet written a default mapping
// do not appear here. Use `GET /rest/v1/default-mappings/{destinationId}` to fetch
// the hydrated would-be row for a specific destination. Requires API-key scope or
// current OAuth user permission: mapping:list
func (r *DefaultMappingService) List(ctx context.Context, opts ...option.RequestOption) (res *DefaultMappingListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/default-mappings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Fetch the destination's default mapping by destination id. Returns a hydrated
// row with empty `mappings[]` when no default mapping has been written yet (so
// callers do not need to handle a 404-vs-200 branch). Each destination has at most
// one default mapping. Requires API-key scope or current OAuth user permission:
// mapping:find
func (r *DefaultMappingService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *DefaultMappingGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/default-mappings/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Upsert the destination default mapping. Always replaces `mappings[]` wholesale
// (default mappings have no merge-partial semantic). Default mappings cannot have
// custom `logic`; the field is not accepted on this endpoint. Requires API-key
// scope or current OAuth user permission: mapping:update
func (r *DefaultMappingService) Replace(ctx context.Context, id string, body DefaultMappingReplaceParams, opts ...option.RequestOption) (res *DefaultMappingReplaceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/default-mappings/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

type DefaultMappingListResponse struct {
	Entities []DefaultMappingListResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DefaultMappingListResponse) RawJSON() string { return r.JSON.raw }
func (r *DefaultMappingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultMappingListResponseEntity struct {
	ID               string                                    `json:"id" api:"required"`
	IsEnabled        bool                                      `json:"isEnabled" api:"required"`
	Mappings         []DefaultMappingListResponseEntityMapping `json:"mappings" api:"required"`
	DestinationID    string                                    `json:"destinationId" api:"nullable"`
	IsDefaultMapping bool                                      `json:"isDefaultMapping" api:"nullable"`
	// Condition tree gating when this mapping fires. A node is either a leaf
	// `condition` or a combinator (`AND`, `OR`, `NOT`). Combinator children are
	// themselves logic nodes, so trees nest arbitrarily.
	//
	// Example leaf:
	// `{ "condition": { "property": "$event.event", "operator": "Is", "value": "page_view" } }`.
	//
	// Example combinator: `{ "AND": [{ "condition": ... }, { "OR": [...] }] }`.
	Logic        DefaultMappingListResponseEntityLogic `json:"logic" api:"nullable"`
	Name         string                                `json:"name" api:"nullable"`
	Priority     float64                               `json:"priority" api:"nullable"`
	SourceID     string                                `json:"sourceId" api:"nullable"`
	TemplateID   string                                `json:"templateId" api:"nullable"`
	TemplateName string                                `json:"templateName" api:"nullable"`
	UpdatedAt    string                                `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		IsEnabled        respjson.Field
		Mappings         respjson.Field
		DestinationID    respjson.Field
		IsDefaultMapping respjson.Field
		Logic            respjson.Field
		Name             respjson.Field
		Priority         respjson.Field
		SourceID         respjson.Field
		TemplateID       respjson.Field
		TemplateName     respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DefaultMappingListResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *DefaultMappingListResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultMappingListResponseEntityMapping struct {
	// Source expression sent to the destination for this `property`. Use `{{...}}`
	// template syntax to substitute values from the event/visitor record:
	// `{{event.event}}`, `{{event.event_properties.value}}`, `{{visitor.email}}`. Bare
	// strings (no `{{}}`) are sent verbatim. Note: `{{...}}` template syntax belongs
	// HERE, NOT in `logic.condition.property` — logic conditions use bare dotted paths
	// like `$event.event_properties.value`.
	Map string `json:"map" api:"required"`
	// Destination-side field name. Comes from the destination template — discover the
	// valid set via `GET /rest/v1/mapping-templates?entityId=...`.
	Property string `json:"property" api:"required"`
	// Any of "CamelCase", "DmaIP", "DomainOnly", "DomainPathOnly", "DomainPathUTMs",
	// "DomainUTMs", "FakeDomain", "FakeDomainRealPath", "FakeIP", "FullUrl", "Hash",
	// "HashMD5", "HashedCountry", "HashedDateOfBirth", "HashedGender",
	// "HashedNormalized", "HashedNormalizedNoSpecialChars", "HashedPhone",
	// "HashedState", "HashedZip", "KebabCase", "LowerCase", "None", "Null",
	// "Redacted", "RegionalIP", "SnakeCase", "StartCase", "UpperCase".
	Modification string `json:"modification" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Map          respjson.Field
		Property     respjson.Field
		Modification respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DefaultMappingListResponseEntityMapping) RawJSON() string { return r.JSON.raw }
func (r *DefaultMappingListResponseEntityMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Condition tree gating when this mapping fires. A node is either a leaf
// `condition` or a combinator (`AND`, `OR`, `NOT`). Combinator children are
// themselves logic nodes, so trees nest arbitrarily.
//
// Example leaf:
// `{ "condition": { "property": "$event.event", "operator": "Is", "value": "page_view" } }`.
//
// Example combinator: `{ "AND": [{ "condition": ... }, { "OR": [...] }] }`.
type DefaultMappingListResponseEntityLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And       []map[string]DefaultMappingListResponseEntityLogicAndUnion `json:"AND" api:"nullable"`
	Condition DefaultMappingListResponseEntityLogicCondition             `json:"condition" api:"nullable"`
	// Negates a single child logic node.
	Not map[string]DefaultMappingListResponseEntityLogicNotUnion `json:"NOT" api:"nullable"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or []map[string]DefaultMappingListResponseEntityLogicOrUnion `json:"OR" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		And         respjson.Field
		Condition   respjson.Field
		Not         respjson.Field
		Or          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DefaultMappingListResponseEntityLogic) RawJSON() string { return r.JSON.raw }
func (r *DefaultMappingListResponseEntityLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DefaultMappingListResponseEntityLogicAndUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfDefaultMappingListResponseEntityLogicAndMapItem]
type DefaultMappingListResponseEntityLogicAndUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfDefaultMappingListResponseEntityLogicAndMapItem any `json:",inline"`
	JSON                                              struct {
		OfString                                          respjson.Field
		OfFloat                                           respjson.Field
		OfBool                                            respjson.Field
		OfAnyArray                                        respjson.Field
		OfDefaultMappingListResponseEntityLogicAndMapItem respjson.Field
		raw                                               string
	} `json:"-"`
}

func (u DefaultMappingListResponseEntityLogicAndUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingListResponseEntityLogicAndUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingListResponseEntityLogicAndUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingListResponseEntityLogicAndUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingListResponseEntityLogicAndUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DefaultMappingListResponseEntityLogicAndUnion) RawJSON() string { return u.JSON.raw }

func (r *DefaultMappingListResponseEntityLogicAndUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultMappingListResponseEntityLogicCondition struct {
	// Comparison verb in PascalCase. Equality/text: `Is`, `IsNot`, `Contains`,
	// `DoesNotContain`, `StartsWith`, `EndsWith`. Truthiness/nullability: `IsFalsy`,
	// `IsTruthy`, `IsNull`, `IsNotNull`, `IsUndefined`, `IsNotUndefined`, `IsTrue`,
	// `IsFalse`. Numeric: `IsGreaterThan`, `IsGreaterThanOrEqual`, `IsLessThan`,
	// `IsLessThanOrEqual`. Set membership: `IsIn`, `IsNotIn`, `IsFoundIn`,
	// `IsNotFoundIn`. Date: `IsBefore`, `IsAfter`, `IsBetween`, `IsOnOrBefore`,
	// `IsOnOrAfter`. Regex: `MatchesRegex`, `MatchesRegexIgnoreCase`,
	// `DoesNotMatchRegex`, `DoesNotMatchRegexIgnoreCase`.
	//
	// Any of "Is", "IsNot", "Contains", "DoesNotContain", "StartsWith", "EndsWith",
	// "IsFalsy", "IsTruthy", "IsNull", "IsNotNull", "IsUndefined", "IsNotUndefined",
	// "IsGreaterThan", "IsGreaterThanOrEqual", "IsLessThan", "IsLessThanOrEqual",
	// "IsIn", "IsNotIn", "IsFoundIn", "IsNotFoundIn", "IsTrue", "IsFalse", "IsBefore",
	// "IsAfter", "IsBetween", "IsOnOrBefore", "IsOnOrAfter", "MatchesRegex",
	// "MatchesRegexIgnoreCase", "DoesNotMatchRegex", "DoesNotMatchRegexIgnoreCase".
	Operator string `json:"operator" api:"required"`
	// Bare dotted path into the event/visitor record. Examples: `$event.event`,
	// `$event.event_properties.value`, `visitor.consent.marketing`. The leading `$` is
	// optional and stripped before lookup. Do **not** use `{{...}}` here — that
	// template syntax is for mapping values (`mappings[].map`), not logic conditions,
	// and would be compared as a literal string.
	Property string `json:"property" api:"required"`
	// String compared against the resolved property. Operators that take no value
	// (`IsFalsy`, `IsTruthy`, `IsNull`, `IsNotNull`, `IsUndefined`, `IsNotUndefined`,
	// `IsTrue`, `IsFalse`) ignore this field — send `""`.
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator    respjson.Field
		Property    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DefaultMappingListResponseEntityLogicCondition) RawJSON() string { return r.JSON.raw }
func (r *DefaultMappingListResponseEntityLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DefaultMappingListResponseEntityLogicNotUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfDefaultMappingListResponseEntityLogicNotMapItem]
type DefaultMappingListResponseEntityLogicNotUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfDefaultMappingListResponseEntityLogicNotMapItem any `json:",inline"`
	JSON                                              struct {
		OfString                                          respjson.Field
		OfFloat                                           respjson.Field
		OfBool                                            respjson.Field
		OfAnyArray                                        respjson.Field
		OfDefaultMappingListResponseEntityLogicNotMapItem respjson.Field
		raw                                               string
	} `json:"-"`
}

func (u DefaultMappingListResponseEntityLogicNotUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingListResponseEntityLogicNotUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingListResponseEntityLogicNotUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingListResponseEntityLogicNotUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingListResponseEntityLogicNotUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DefaultMappingListResponseEntityLogicNotUnion) RawJSON() string { return u.JSON.raw }

func (r *DefaultMappingListResponseEntityLogicNotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DefaultMappingListResponseEntityLogicOrUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfDefaultMappingListResponseEntityLogicOrMapItem]
type DefaultMappingListResponseEntityLogicOrUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfDefaultMappingListResponseEntityLogicOrMapItem any `json:",inline"`
	JSON                                             struct {
		OfString                                         respjson.Field
		OfFloat                                          respjson.Field
		OfBool                                           respjson.Field
		OfAnyArray                                       respjson.Field
		OfDefaultMappingListResponseEntityLogicOrMapItem respjson.Field
		raw                                              string
	} `json:"-"`
}

func (u DefaultMappingListResponseEntityLogicOrUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingListResponseEntityLogicOrUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingListResponseEntityLogicOrUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingListResponseEntityLogicOrUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingListResponseEntityLogicOrUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DefaultMappingListResponseEntityLogicOrUnion) RawJSON() string { return u.JSON.raw }

func (r *DefaultMappingListResponseEntityLogicOrUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultMappingGetResponse struct {
	ID               string                             `json:"id" api:"required"`
	IsEnabled        bool                               `json:"isEnabled" api:"required"`
	Mappings         []DefaultMappingGetResponseMapping `json:"mappings" api:"required"`
	DestinationID    string                             `json:"destinationId" api:"nullable"`
	IsDefaultMapping bool                               `json:"isDefaultMapping" api:"nullable"`
	// Condition tree gating when this mapping fires. A node is either a leaf
	// `condition` or a combinator (`AND`, `OR`, `NOT`). Combinator children are
	// themselves logic nodes, so trees nest arbitrarily.
	//
	// Example leaf:
	// `{ "condition": { "property": "$event.event", "operator": "Is", "value": "page_view" } }`.
	//
	// Example combinator: `{ "AND": [{ "condition": ... }, { "OR": [...] }] }`.
	Logic        DefaultMappingGetResponseLogic `json:"logic" api:"nullable"`
	Name         string                         `json:"name" api:"nullable"`
	Priority     float64                        `json:"priority" api:"nullable"`
	SourceID     string                         `json:"sourceId" api:"nullable"`
	TemplateID   string                         `json:"templateId" api:"nullable"`
	TemplateName string                         `json:"templateName" api:"nullable"`
	UpdatedAt    string                         `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		IsEnabled        respjson.Field
		Mappings         respjson.Field
		DestinationID    respjson.Field
		IsDefaultMapping respjson.Field
		Logic            respjson.Field
		Name             respjson.Field
		Priority         respjson.Field
		SourceID         respjson.Field
		TemplateID       respjson.Field
		TemplateName     respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DefaultMappingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DefaultMappingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultMappingGetResponseMapping struct {
	// Source expression sent to the destination for this `property`. Use `{{...}}`
	// template syntax to substitute values from the event/visitor record:
	// `{{event.event}}`, `{{event.event_properties.value}}`, `{{visitor.email}}`. Bare
	// strings (no `{{}}`) are sent verbatim. Note: `{{...}}` template syntax belongs
	// HERE, NOT in `logic.condition.property` — logic conditions use bare dotted paths
	// like `$event.event_properties.value`.
	Map string `json:"map" api:"required"`
	// Destination-side field name. Comes from the destination template — discover the
	// valid set via `GET /rest/v1/mapping-templates?entityId=...`.
	Property string `json:"property" api:"required"`
	// Any of "CamelCase", "DmaIP", "DomainOnly", "DomainPathOnly", "DomainPathUTMs",
	// "DomainUTMs", "FakeDomain", "FakeDomainRealPath", "FakeIP", "FullUrl", "Hash",
	// "HashMD5", "HashedCountry", "HashedDateOfBirth", "HashedGender",
	// "HashedNormalized", "HashedNormalizedNoSpecialChars", "HashedPhone",
	// "HashedState", "HashedZip", "KebabCase", "LowerCase", "None", "Null",
	// "Redacted", "RegionalIP", "SnakeCase", "StartCase", "UpperCase".
	Modification string `json:"modification" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Map          respjson.Field
		Property     respjson.Field
		Modification respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DefaultMappingGetResponseMapping) RawJSON() string { return r.JSON.raw }
func (r *DefaultMappingGetResponseMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Condition tree gating when this mapping fires. A node is either a leaf
// `condition` or a combinator (`AND`, `OR`, `NOT`). Combinator children are
// themselves logic nodes, so trees nest arbitrarily.
//
// Example leaf:
// `{ "condition": { "property": "$event.event", "operator": "Is", "value": "page_view" } }`.
//
// Example combinator: `{ "AND": [{ "condition": ... }, { "OR": [...] }] }`.
type DefaultMappingGetResponseLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And       []map[string]DefaultMappingGetResponseLogicAndUnion `json:"AND" api:"nullable"`
	Condition DefaultMappingGetResponseLogicCondition             `json:"condition" api:"nullable"`
	// Negates a single child logic node.
	Not map[string]DefaultMappingGetResponseLogicNotUnion `json:"NOT" api:"nullable"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or []map[string]DefaultMappingGetResponseLogicOrUnion `json:"OR" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		And         respjson.Field
		Condition   respjson.Field
		Not         respjson.Field
		Or          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DefaultMappingGetResponseLogic) RawJSON() string { return r.JSON.raw }
func (r *DefaultMappingGetResponseLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DefaultMappingGetResponseLogicAndUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfDefaultMappingGetResponseLogicAndMapItem]
type DefaultMappingGetResponseLogicAndUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfDefaultMappingGetResponseLogicAndMapItem any `json:",inline"`
	JSON                                       struct {
		OfString                                   respjson.Field
		OfFloat                                    respjson.Field
		OfBool                                     respjson.Field
		OfAnyArray                                 respjson.Field
		OfDefaultMappingGetResponseLogicAndMapItem respjson.Field
		raw                                        string
	} `json:"-"`
}

func (u DefaultMappingGetResponseLogicAndUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingGetResponseLogicAndUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingGetResponseLogicAndUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingGetResponseLogicAndUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingGetResponseLogicAndUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DefaultMappingGetResponseLogicAndUnion) RawJSON() string { return u.JSON.raw }

func (r *DefaultMappingGetResponseLogicAndUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultMappingGetResponseLogicCondition struct {
	// Comparison verb in PascalCase. Equality/text: `Is`, `IsNot`, `Contains`,
	// `DoesNotContain`, `StartsWith`, `EndsWith`. Truthiness/nullability: `IsFalsy`,
	// `IsTruthy`, `IsNull`, `IsNotNull`, `IsUndefined`, `IsNotUndefined`, `IsTrue`,
	// `IsFalse`. Numeric: `IsGreaterThan`, `IsGreaterThanOrEqual`, `IsLessThan`,
	// `IsLessThanOrEqual`. Set membership: `IsIn`, `IsNotIn`, `IsFoundIn`,
	// `IsNotFoundIn`. Date: `IsBefore`, `IsAfter`, `IsBetween`, `IsOnOrBefore`,
	// `IsOnOrAfter`. Regex: `MatchesRegex`, `MatchesRegexIgnoreCase`,
	// `DoesNotMatchRegex`, `DoesNotMatchRegexIgnoreCase`.
	//
	// Any of "Is", "IsNot", "Contains", "DoesNotContain", "StartsWith", "EndsWith",
	// "IsFalsy", "IsTruthy", "IsNull", "IsNotNull", "IsUndefined", "IsNotUndefined",
	// "IsGreaterThan", "IsGreaterThanOrEqual", "IsLessThan", "IsLessThanOrEqual",
	// "IsIn", "IsNotIn", "IsFoundIn", "IsNotFoundIn", "IsTrue", "IsFalse", "IsBefore",
	// "IsAfter", "IsBetween", "IsOnOrBefore", "IsOnOrAfter", "MatchesRegex",
	// "MatchesRegexIgnoreCase", "DoesNotMatchRegex", "DoesNotMatchRegexIgnoreCase".
	Operator string `json:"operator" api:"required"`
	// Bare dotted path into the event/visitor record. Examples: `$event.event`,
	// `$event.event_properties.value`, `visitor.consent.marketing`. The leading `$` is
	// optional and stripped before lookup. Do **not** use `{{...}}` here — that
	// template syntax is for mapping values (`mappings[].map`), not logic conditions,
	// and would be compared as a literal string.
	Property string `json:"property" api:"required"`
	// String compared against the resolved property. Operators that take no value
	// (`IsFalsy`, `IsTruthy`, `IsNull`, `IsNotNull`, `IsUndefined`, `IsNotUndefined`,
	// `IsTrue`, `IsFalse`) ignore this field — send `""`.
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator    respjson.Field
		Property    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DefaultMappingGetResponseLogicCondition) RawJSON() string { return r.JSON.raw }
func (r *DefaultMappingGetResponseLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DefaultMappingGetResponseLogicNotUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfDefaultMappingGetResponseLogicNotMapItem]
type DefaultMappingGetResponseLogicNotUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfDefaultMappingGetResponseLogicNotMapItem any `json:",inline"`
	JSON                                       struct {
		OfString                                   respjson.Field
		OfFloat                                    respjson.Field
		OfBool                                     respjson.Field
		OfAnyArray                                 respjson.Field
		OfDefaultMappingGetResponseLogicNotMapItem respjson.Field
		raw                                        string
	} `json:"-"`
}

func (u DefaultMappingGetResponseLogicNotUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingGetResponseLogicNotUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingGetResponseLogicNotUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingGetResponseLogicNotUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingGetResponseLogicNotUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DefaultMappingGetResponseLogicNotUnion) RawJSON() string { return u.JSON.raw }

func (r *DefaultMappingGetResponseLogicNotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DefaultMappingGetResponseLogicOrUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfDefaultMappingGetResponseLogicOrMapItem]
type DefaultMappingGetResponseLogicOrUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfDefaultMappingGetResponseLogicOrMapItem any `json:",inline"`
	JSON                                      struct {
		OfString                                  respjson.Field
		OfFloat                                   respjson.Field
		OfBool                                    respjson.Field
		OfAnyArray                                respjson.Field
		OfDefaultMappingGetResponseLogicOrMapItem respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u DefaultMappingGetResponseLogicOrUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingGetResponseLogicOrUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingGetResponseLogicOrUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingGetResponseLogicOrUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingGetResponseLogicOrUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DefaultMappingGetResponseLogicOrUnion) RawJSON() string { return u.JSON.raw }

func (r *DefaultMappingGetResponseLogicOrUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultMappingReplaceResponse struct {
	ID               string                                 `json:"id" api:"required"`
	IsEnabled        bool                                   `json:"isEnabled" api:"required"`
	Mappings         []DefaultMappingReplaceResponseMapping `json:"mappings" api:"required"`
	DestinationID    string                                 `json:"destinationId" api:"nullable"`
	IsDefaultMapping bool                                   `json:"isDefaultMapping" api:"nullable"`
	// Condition tree gating when this mapping fires. A node is either a leaf
	// `condition` or a combinator (`AND`, `OR`, `NOT`). Combinator children are
	// themselves logic nodes, so trees nest arbitrarily.
	//
	// Example leaf:
	// `{ "condition": { "property": "$event.event", "operator": "Is", "value": "page_view" } }`.
	//
	// Example combinator: `{ "AND": [{ "condition": ... }, { "OR": [...] }] }`.
	Logic        DefaultMappingReplaceResponseLogic `json:"logic" api:"nullable"`
	Name         string                             `json:"name" api:"nullable"`
	Priority     float64                            `json:"priority" api:"nullable"`
	SourceID     string                             `json:"sourceId" api:"nullable"`
	TemplateID   string                             `json:"templateId" api:"nullable"`
	TemplateName string                             `json:"templateName" api:"nullable"`
	UpdatedAt    string                             `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		IsEnabled        respjson.Field
		Mappings         respjson.Field
		DestinationID    respjson.Field
		IsDefaultMapping respjson.Field
		Logic            respjson.Field
		Name             respjson.Field
		Priority         respjson.Field
		SourceID         respjson.Field
		TemplateID       respjson.Field
		TemplateName     respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DefaultMappingReplaceResponse) RawJSON() string { return r.JSON.raw }
func (r *DefaultMappingReplaceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultMappingReplaceResponseMapping struct {
	// Source expression sent to the destination for this `property`. Use `{{...}}`
	// template syntax to substitute values from the event/visitor record:
	// `{{event.event}}`, `{{event.event_properties.value}}`, `{{visitor.email}}`. Bare
	// strings (no `{{}}`) are sent verbatim. Note: `{{...}}` template syntax belongs
	// HERE, NOT in `logic.condition.property` — logic conditions use bare dotted paths
	// like `$event.event_properties.value`.
	Map string `json:"map" api:"required"`
	// Destination-side field name. Comes from the destination template — discover the
	// valid set via `GET /rest/v1/mapping-templates?entityId=...`.
	Property string `json:"property" api:"required"`
	// Any of "CamelCase", "DmaIP", "DomainOnly", "DomainPathOnly", "DomainPathUTMs",
	// "DomainUTMs", "FakeDomain", "FakeDomainRealPath", "FakeIP", "FullUrl", "Hash",
	// "HashMD5", "HashedCountry", "HashedDateOfBirth", "HashedGender",
	// "HashedNormalized", "HashedNormalizedNoSpecialChars", "HashedPhone",
	// "HashedState", "HashedZip", "KebabCase", "LowerCase", "None", "Null",
	// "Redacted", "RegionalIP", "SnakeCase", "StartCase", "UpperCase".
	Modification string `json:"modification" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Map          respjson.Field
		Property     respjson.Field
		Modification respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DefaultMappingReplaceResponseMapping) RawJSON() string { return r.JSON.raw }
func (r *DefaultMappingReplaceResponseMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Condition tree gating when this mapping fires. A node is either a leaf
// `condition` or a combinator (`AND`, `OR`, `NOT`). Combinator children are
// themselves logic nodes, so trees nest arbitrarily.
//
// Example leaf:
// `{ "condition": { "property": "$event.event", "operator": "Is", "value": "page_view" } }`.
//
// Example combinator: `{ "AND": [{ "condition": ... }, { "OR": [...] }] }`.
type DefaultMappingReplaceResponseLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And       []map[string]DefaultMappingReplaceResponseLogicAndUnion `json:"AND" api:"nullable"`
	Condition DefaultMappingReplaceResponseLogicCondition             `json:"condition" api:"nullable"`
	// Negates a single child logic node.
	Not map[string]DefaultMappingReplaceResponseLogicNotUnion `json:"NOT" api:"nullable"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or []map[string]DefaultMappingReplaceResponseLogicOrUnion `json:"OR" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		And         respjson.Field
		Condition   respjson.Field
		Not         respjson.Field
		Or          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DefaultMappingReplaceResponseLogic) RawJSON() string { return r.JSON.raw }
func (r *DefaultMappingReplaceResponseLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DefaultMappingReplaceResponseLogicAndUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfDefaultMappingReplaceResponseLogicAndMapItem]
type DefaultMappingReplaceResponseLogicAndUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfDefaultMappingReplaceResponseLogicAndMapItem any `json:",inline"`
	JSON                                           struct {
		OfString                                       respjson.Field
		OfFloat                                        respjson.Field
		OfBool                                         respjson.Field
		OfAnyArray                                     respjson.Field
		OfDefaultMappingReplaceResponseLogicAndMapItem respjson.Field
		raw                                            string
	} `json:"-"`
}

func (u DefaultMappingReplaceResponseLogicAndUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingReplaceResponseLogicAndUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingReplaceResponseLogicAndUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingReplaceResponseLogicAndUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingReplaceResponseLogicAndUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DefaultMappingReplaceResponseLogicAndUnion) RawJSON() string { return u.JSON.raw }

func (r *DefaultMappingReplaceResponseLogicAndUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultMappingReplaceResponseLogicCondition struct {
	// Comparison verb in PascalCase. Equality/text: `Is`, `IsNot`, `Contains`,
	// `DoesNotContain`, `StartsWith`, `EndsWith`. Truthiness/nullability: `IsFalsy`,
	// `IsTruthy`, `IsNull`, `IsNotNull`, `IsUndefined`, `IsNotUndefined`, `IsTrue`,
	// `IsFalse`. Numeric: `IsGreaterThan`, `IsGreaterThanOrEqual`, `IsLessThan`,
	// `IsLessThanOrEqual`. Set membership: `IsIn`, `IsNotIn`, `IsFoundIn`,
	// `IsNotFoundIn`. Date: `IsBefore`, `IsAfter`, `IsBetween`, `IsOnOrBefore`,
	// `IsOnOrAfter`. Regex: `MatchesRegex`, `MatchesRegexIgnoreCase`,
	// `DoesNotMatchRegex`, `DoesNotMatchRegexIgnoreCase`.
	//
	// Any of "Is", "IsNot", "Contains", "DoesNotContain", "StartsWith", "EndsWith",
	// "IsFalsy", "IsTruthy", "IsNull", "IsNotNull", "IsUndefined", "IsNotUndefined",
	// "IsGreaterThan", "IsGreaterThanOrEqual", "IsLessThan", "IsLessThanOrEqual",
	// "IsIn", "IsNotIn", "IsFoundIn", "IsNotFoundIn", "IsTrue", "IsFalse", "IsBefore",
	// "IsAfter", "IsBetween", "IsOnOrBefore", "IsOnOrAfter", "MatchesRegex",
	// "MatchesRegexIgnoreCase", "DoesNotMatchRegex", "DoesNotMatchRegexIgnoreCase".
	Operator string `json:"operator" api:"required"`
	// Bare dotted path into the event/visitor record. Examples: `$event.event`,
	// `$event.event_properties.value`, `visitor.consent.marketing`. The leading `$` is
	// optional and stripped before lookup. Do **not** use `{{...}}` here — that
	// template syntax is for mapping values (`mappings[].map`), not logic conditions,
	// and would be compared as a literal string.
	Property string `json:"property" api:"required"`
	// String compared against the resolved property. Operators that take no value
	// (`IsFalsy`, `IsTruthy`, `IsNull`, `IsNotNull`, `IsUndefined`, `IsNotUndefined`,
	// `IsTrue`, `IsFalse`) ignore this field — send `""`.
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator    respjson.Field
		Property    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DefaultMappingReplaceResponseLogicCondition) RawJSON() string { return r.JSON.raw }
func (r *DefaultMappingReplaceResponseLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DefaultMappingReplaceResponseLogicNotUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfDefaultMappingReplaceResponseLogicNotMapItem]
type DefaultMappingReplaceResponseLogicNotUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfDefaultMappingReplaceResponseLogicNotMapItem any `json:",inline"`
	JSON                                           struct {
		OfString                                       respjson.Field
		OfFloat                                        respjson.Field
		OfBool                                         respjson.Field
		OfAnyArray                                     respjson.Field
		OfDefaultMappingReplaceResponseLogicNotMapItem respjson.Field
		raw                                            string
	} `json:"-"`
}

func (u DefaultMappingReplaceResponseLogicNotUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingReplaceResponseLogicNotUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingReplaceResponseLogicNotUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingReplaceResponseLogicNotUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingReplaceResponseLogicNotUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DefaultMappingReplaceResponseLogicNotUnion) RawJSON() string { return u.JSON.raw }

func (r *DefaultMappingReplaceResponseLogicNotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DefaultMappingReplaceResponseLogicOrUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfDefaultMappingReplaceResponseLogicOrMapItem]
type DefaultMappingReplaceResponseLogicOrUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfDefaultMappingReplaceResponseLogicOrMapItem any `json:",inline"`
	JSON                                          struct {
		OfString                                      respjson.Field
		OfFloat                                       respjson.Field
		OfBool                                        respjson.Field
		OfAnyArray                                    respjson.Field
		OfDefaultMappingReplaceResponseLogicOrMapItem respjson.Field
		raw                                           string
	} `json:"-"`
}

func (u DefaultMappingReplaceResponseLogicOrUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingReplaceResponseLogicOrUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingReplaceResponseLogicOrUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingReplaceResponseLogicOrUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DefaultMappingReplaceResponseLogicOrUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DefaultMappingReplaceResponseLogicOrUnion) RawJSON() string { return u.JSON.raw }

func (r *DefaultMappingReplaceResponseLogicOrUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultMappingReplaceParams struct {
	// Property mappings to persist as the destination default. Use
	// `GET /rest/v1/mapping-templates?entityId={destinationId}` to discover the valid
	// `property` values.
	Mappings []DefaultMappingReplaceParamsMapping `json:"mappings,omitzero" api:"required"`
	// Toggle the default mapping on/off. Defaults to `true` when omitted. `null` is
	// treated as omitted.
	IsEnabled param.Opt[bool] `json:"isEnabled,omitzero"`
	paramObj
}

func (r DefaultMappingReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow DefaultMappingReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DefaultMappingReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Map, Property are required.
type DefaultMappingReplaceParamsMapping struct {
	// Source expression sent to the destination for this `property`. Use `{{...}}`
	// template syntax to substitute values from the event/visitor record:
	// `{{event.event}}`, `{{event.event_properties.value}}`, `{{visitor.email}}`. Bare
	// strings (no `{{}}`) are sent verbatim. Note: `{{...}}` template syntax belongs
	// HERE, NOT in `logic.condition.property` — logic conditions use bare dotted paths
	// like `$event.event_properties.value`.
	Map string `json:"map" api:"required"`
	// Destination-side field name. Comes from the destination template — discover the
	// valid set via `GET /rest/v1/mapping-templates?entityId=...`.
	Property string `json:"property" api:"required"`
	// Any of "CamelCase", "DmaIP", "DomainOnly", "DomainPathOnly", "DomainPathUTMs",
	// "DomainUTMs", "FakeDomain", "FakeDomainRealPath", "FakeIP", "FullUrl", "Hash",
	// "HashMD5", "HashedCountry", "HashedDateOfBirth", "HashedGender",
	// "HashedNormalized", "HashedNormalizedNoSpecialChars", "HashedPhone",
	// "HashedState", "HashedZip", "KebabCase", "LowerCase", "None", "Null",
	// "Redacted", "RegionalIP", "SnakeCase", "StartCase", "UpperCase".
	Modification string `json:"modification,omitzero"`
	paramObj
}

func (r DefaultMappingReplaceParamsMapping) MarshalJSON() (data []byte, err error) {
	type shadow DefaultMappingReplaceParamsMapping
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DefaultMappingReplaceParamsMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DefaultMappingReplaceParamsMapping](
		"modification", "CamelCase", "DmaIP", "DomainOnly", "DomainPathOnly", "DomainPathUTMs", "DomainUTMs", "FakeDomain", "FakeDomainRealPath", "FakeIP", "FullUrl", "Hash", "HashMD5", "HashedCountry", "HashedDateOfBirth", "HashedGender", "HashedNormalized", "HashedNormalizedNoSpecialChars", "HashedPhone", "HashedState", "HashedZip", "KebabCase", "LowerCase", "None", "Null", "Redacted", "RegionalIP", "SnakeCase", "StartCase", "UpperCase",
	)
}
