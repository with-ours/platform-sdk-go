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

// MappingService contains methods and other services that help with interacting
// with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMappingService] method instead.
type MappingService struct {
	Options []option.RequestOption
}

// NewMappingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMappingService(opts ...option.RequestOption) (r MappingService) {
	r = MappingService{}
	r.Options = opts
	return
}

// List mappings for an entity (a source or destination). Requires the `entityId`
// query parameter. Supports cursor pagination via `limit` and `cursor`. Sorted by
// `priority` ascending, then by `id` for deterministic pagination. Requires
// API-key scope or current OAuth user permission: mapping:list
func (r *MappingService) List(ctx context.Context, query MappingListParams, opts ...option.RequestOption) (res *pagination.Cursor[MappingListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "rest/v1/mappings"
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

// List mappings for an entity (a source or destination). Requires the `entityId`
// query parameter. Supports cursor pagination via `limit` and `cursor`. Sorted by
// `priority` ascending, then by `id` for deterministic pagination. Requires
// API-key scope or current OAuth user permission: mapping:list
func (r *MappingService) ListAutoPaging(ctx context.Context, query MappingListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[MappingListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Create a mapping. Two body shapes are accepted:
//
//  1. Quick-create — `{ allowedEventId, destinationId }`. Binds an allowed event to
//     a destination. Returns a slim entity with empty `mappings[]`; follow up with
//     PATCH to populate fields.
//  2. Template fat-create —
//     `{ entityId, templateId, mappings?, logic?, isEnabled?, name?, insertAfterIdx? }`.
//     Lands a fully-shaped mapping in one round-trip. Use
//     `GET /rest/v1/mapping-templates?entityId=...` to discover the valid
//     `templateId` and `mappings[].property` values.
//
// Sending both `allowedEventId` and `templateId` returns 400. Requires API-key
// scope or current OAuth user permission: mapping:create
func (r *MappingService) New(ctx context.Context, body MappingNewParams, opts ...option.RequestOption) (res *MappingNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/mappings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Find a single mapping by ID. Requires API-key scope or current OAuth user
// permission: mapping:find
func (r *MappingService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *MappingGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/mappings/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Partially update a mapping. Only the fields you send are changed. Send
// `isEnabled: false` to pause the mapping without changing other fields (mirrors
// `status` on destinations). `mappings[]` is replaced wholesale when sent.
// Requires API-key scope or current OAuth user permission: mapping:update
func (r *MappingService) Update(ctx context.Context, id string, body MappingUpdateParams, opts ...option.RequestOption) (res *MappingUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/mappings/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete a mapping. Requires API-key scope or current OAuth user permission:
// mapping:delete
func (r *MappingService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/mappings/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Preview a destination-specific event rename before setting or clearing it. Check
// `supportsRenamedEvents` in `GET /rest/v1/destinations/types` first. If
// `willCascade` is true, review all `affectedEventNames` before calling
// `PUT /rest/v1/mappings/rename` with `confirmCascade: true`. Requires API-key
// scope or current OAuth user permission: mapping:find
func (r *MappingService) RenamePreview(ctx context.Context, query MappingRenamePreviewParams, opts ...option.RequestOption) (res *MappingRenamePreviewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/mappings/rename-preview"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Set or clear a destination-specific allowed-event rename. Send
// `renamedEventName: null` to clear it; clearing an existing mapper also requires
// the `mapping:delete` scope. Preview first with
// `GET /rest/v1/mappings/rename-preview`; a shared mapper requires
// `confirmCascade: true`, and all affected events will change. Publish a new
// version to activate the change. Requires API-key scope or current OAuth user
// permission: mapping:update
func (r *MappingService) Rename(ctx context.Context, body MappingRenameParams, opts ...option.RequestOption) (res *MappingRenameResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/mappings/rename"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Reassign `priority` for a set of mappings. Pass `{ uuids: [...] }` with the
// mapping ids in their new order — index 0 becomes the highest-priority mapping.
// All ids must belong to the same parent entity (source or destination); mixing
// entities returns 400. Requires API-key scope or current OAuth user permission:
// mapping:update
func (r *MappingService) Reorder(ctx context.Context, body MappingReorderParams, opts ...option.RequestOption) (res *MappingReorderResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/mappings/reorder"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Discover every mapping template available for a destination or source, with full
// property descriptors inlined. Use the returned `id` as `templateId` when calling
// `POST /rest/v1/mappings` (template fat-create variant), and use each entry under
// `mappings[]` to learn the valid `property`, `kind`, `modificationOptions`, and
// any enforced `options`. The `isDefault: true` entry is the destination's
// built-in default template, configured via
// `PUT /rest/v1/default-mappings/{destinationId}`. Requires API-key scope or
// current OAuth user permission: mapping:find
func (r *MappingService) Templates(ctx context.Context, query MappingTemplatesParams, opts ...option.RequestOption) (res *MappingTemplatesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/mappings/templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Lists the platform-provided variables that any mapping `value` can reference
// (e.g. `event.email`, `event.request_context.ip`, `visitor.id`). Account-agnostic
// discovery — use these paths as the right-hand side of a mapping field. Requires
// API-key scope or current OAuth user permission: variables:find-default
func (r *MappingService) DefaultVariables(ctx context.Context, opts ...option.RequestOption) (res *MappingDefaultVariablesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/mappings/default-variables"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists the custom variables observed in this account’s recent event stream (last
// 14 days). These are dot-paths under `event.event_properties.*` that callers can
// target in mapping `value` fields. Newly seen properties may take a few minutes
// to appear; an empty list means no custom properties have been observed yet for
// this account. Requires API-key scope or current OAuth user permission:
// variables:find-custom
func (r *MappingService) CustomVariables(ctx context.Context, opts ...option.RequestOption) (res *MappingCustomVariablesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/mappings/custom-variables"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists every value accepted on a mapping field’s `modification` property, with a
// human-readable label and one-sentence description. Account-agnostic. Use this
// alongside `GET /rest/v1/mapping-templates` to render a labelled modification
// picker without hardcoding the enum. Requires API-key scope or current OAuth user
// permission: variables:find-default
func (r *MappingService) Modifications(ctx context.Context, opts ...option.RequestOption) (res *MappingModificationsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/mappings/modifications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type MappingListResponse struct {
	ID               string                       `json:"id" api:"required"`
	IsEnabled        bool                         `json:"isEnabled" api:"required"`
	Mappings         []MappingListResponseMapping `json:"mappings" api:"required"`
	DestinationID    string                       `json:"destinationId" api:"nullable"`
	IsDefaultMapping bool                         `json:"isDefaultMapping" api:"nullable"`
	// Condition tree gating when this mapping fires. A node is either a leaf
	// `condition` or a combinator (`AND`, `OR`, `NOT`). Combinator children are
	// themselves logic nodes, so trees nest arbitrarily.
	//
	// Example leaf:
	// `{ "condition": { "property": "$event.event", "operator": "Is", "value": "page_view" } }`.
	//
	// Example combinator: `{ "AND": [{ "condition": ... }, { "OR": [...] }] }`.
	Logic        MappingListResponseLogic `json:"logic" api:"nullable"`
	Name         string                   `json:"name" api:"nullable"`
	Priority     float64                  `json:"priority" api:"nullable"`
	SourceID     string                   `json:"sourceId" api:"nullable"`
	TemplateID   string                   `json:"templateId" api:"nullable"`
	TemplateName string                   `json:"templateName" api:"nullable"`
	UpdatedAt    string                   `json:"updatedAt" api:"nullable"`
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
func (r MappingListResponse) RawJSON() string { return r.JSON.raw }
func (r *MappingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingListResponseMapping struct {
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
func (r MappingListResponseMapping) RawJSON() string { return r.JSON.raw }
func (r *MappingListResponseMapping) UnmarshalJSON(data []byte) error {
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
type MappingListResponseLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And       []map[string]MappingListResponseLogicAndUnion `json:"AND" api:"nullable"`
	Condition MappingListResponseLogicCondition             `json:"condition" api:"nullable"`
	// Negates a single child logic node.
	Not map[string]MappingListResponseLogicNotUnion `json:"NOT" api:"nullable"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or []map[string]MappingListResponseLogicOrUnion `json:"OR" api:"nullable"`
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
func (r MappingListResponseLogic) RawJSON() string { return r.JSON.raw }
func (r *MappingListResponseLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MappingListResponseLogicAndUnion contains all possible properties and values
// from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfMappingListResponseLogicAndMapItem]
type MappingListResponseLogicAndUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfMappingListResponseLogicAndMapItem any `json:",inline"`
	JSON                                 struct {
		OfString                             respjson.Field
		OfFloat                              respjson.Field
		OfBool                               respjson.Field
		OfAnyArray                           respjson.Field
		OfMappingListResponseLogicAndMapItem respjson.Field
		raw                                  string
	} `json:"-"`
}

func (u MappingListResponseLogicAndUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingListResponseLogicAndUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingListResponseLogicAndUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingListResponseLogicAndUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingListResponseLogicAndUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MappingListResponseLogicAndUnion) RawJSON() string { return u.JSON.raw }

func (r *MappingListResponseLogicAndUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingListResponseLogicCondition struct {
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
func (r MappingListResponseLogicCondition) RawJSON() string { return r.JSON.raw }
func (r *MappingListResponseLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MappingListResponseLogicNotUnion contains all possible properties and values
// from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfMappingListResponseLogicNotMapItem]
type MappingListResponseLogicNotUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfMappingListResponseLogicNotMapItem any `json:",inline"`
	JSON                                 struct {
		OfString                             respjson.Field
		OfFloat                              respjson.Field
		OfBool                               respjson.Field
		OfAnyArray                           respjson.Field
		OfMappingListResponseLogicNotMapItem respjson.Field
		raw                                  string
	} `json:"-"`
}

func (u MappingListResponseLogicNotUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingListResponseLogicNotUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingListResponseLogicNotUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingListResponseLogicNotUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingListResponseLogicNotUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MappingListResponseLogicNotUnion) RawJSON() string { return u.JSON.raw }

func (r *MappingListResponseLogicNotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MappingListResponseLogicOrUnion contains all possible properties and values from
// [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfMappingListResponseLogicOrMapItem]
type MappingListResponseLogicOrUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfMappingListResponseLogicOrMapItem any `json:",inline"`
	JSON                                struct {
		OfString                            respjson.Field
		OfFloat                             respjson.Field
		OfBool                              respjson.Field
		OfAnyArray                          respjson.Field
		OfMappingListResponseLogicOrMapItem respjson.Field
		raw                                 string
	} `json:"-"`
}

func (u MappingListResponseLogicOrUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingListResponseLogicOrUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingListResponseLogicOrUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingListResponseLogicOrUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingListResponseLogicOrUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MappingListResponseLogicOrUnion) RawJSON() string { return u.JSON.raw }

func (r *MappingListResponseLogicOrUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingNewResponse struct {
	ID               string                      `json:"id" api:"required"`
	IsEnabled        bool                        `json:"isEnabled" api:"required"`
	Mappings         []MappingNewResponseMapping `json:"mappings" api:"required"`
	DestinationID    string                      `json:"destinationId" api:"nullable"`
	IsDefaultMapping bool                        `json:"isDefaultMapping" api:"nullable"`
	// Condition tree gating when this mapping fires. A node is either a leaf
	// `condition` or a combinator (`AND`, `OR`, `NOT`). Combinator children are
	// themselves logic nodes, so trees nest arbitrarily.
	//
	// Example leaf:
	// `{ "condition": { "property": "$event.event", "operator": "Is", "value": "page_view" } }`.
	//
	// Example combinator: `{ "AND": [{ "condition": ... }, { "OR": [...] }] }`.
	Logic        MappingNewResponseLogic `json:"logic" api:"nullable"`
	Name         string                  `json:"name" api:"nullable"`
	Priority     float64                 `json:"priority" api:"nullable"`
	SourceID     string                  `json:"sourceId" api:"nullable"`
	TemplateID   string                  `json:"templateId" api:"nullable"`
	TemplateName string                  `json:"templateName" api:"nullable"`
	UpdatedAt    string                  `json:"updatedAt" api:"nullable"`
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
func (r MappingNewResponse) RawJSON() string { return r.JSON.raw }
func (r *MappingNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingNewResponseMapping struct {
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
func (r MappingNewResponseMapping) RawJSON() string { return r.JSON.raw }
func (r *MappingNewResponseMapping) UnmarshalJSON(data []byte) error {
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
type MappingNewResponseLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And       []map[string]MappingNewResponseLogicAndUnion `json:"AND" api:"nullable"`
	Condition MappingNewResponseLogicCondition             `json:"condition" api:"nullable"`
	// Negates a single child logic node.
	Not map[string]MappingNewResponseLogicNotUnion `json:"NOT" api:"nullable"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or []map[string]MappingNewResponseLogicOrUnion `json:"OR" api:"nullable"`
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
func (r MappingNewResponseLogic) RawJSON() string { return r.JSON.raw }
func (r *MappingNewResponseLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MappingNewResponseLogicAndUnion contains all possible properties and values from
// [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfMappingNewResponseLogicAndMapItem]
type MappingNewResponseLogicAndUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfMappingNewResponseLogicAndMapItem any `json:",inline"`
	JSON                                struct {
		OfString                            respjson.Field
		OfFloat                             respjson.Field
		OfBool                              respjson.Field
		OfAnyArray                          respjson.Field
		OfMappingNewResponseLogicAndMapItem respjson.Field
		raw                                 string
	} `json:"-"`
}

func (u MappingNewResponseLogicAndUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingNewResponseLogicAndUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingNewResponseLogicAndUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingNewResponseLogicAndUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingNewResponseLogicAndUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MappingNewResponseLogicAndUnion) RawJSON() string { return u.JSON.raw }

func (r *MappingNewResponseLogicAndUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingNewResponseLogicCondition struct {
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
func (r MappingNewResponseLogicCondition) RawJSON() string { return r.JSON.raw }
func (r *MappingNewResponseLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MappingNewResponseLogicNotUnion contains all possible properties and values from
// [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfMappingNewResponseLogicNotMapItem]
type MappingNewResponseLogicNotUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfMappingNewResponseLogicNotMapItem any `json:",inline"`
	JSON                                struct {
		OfString                            respjson.Field
		OfFloat                             respjson.Field
		OfBool                              respjson.Field
		OfAnyArray                          respjson.Field
		OfMappingNewResponseLogicNotMapItem respjson.Field
		raw                                 string
	} `json:"-"`
}

func (u MappingNewResponseLogicNotUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingNewResponseLogicNotUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingNewResponseLogicNotUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingNewResponseLogicNotUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingNewResponseLogicNotUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MappingNewResponseLogicNotUnion) RawJSON() string { return u.JSON.raw }

func (r *MappingNewResponseLogicNotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MappingNewResponseLogicOrUnion contains all possible properties and values from
// [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfMappingNewResponseLogicOrMapItem]
type MappingNewResponseLogicOrUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfMappingNewResponseLogicOrMapItem any `json:",inline"`
	JSON                               struct {
		OfString                           respjson.Field
		OfFloat                            respjson.Field
		OfBool                             respjson.Field
		OfAnyArray                         respjson.Field
		OfMappingNewResponseLogicOrMapItem respjson.Field
		raw                                string
	} `json:"-"`
}

func (u MappingNewResponseLogicOrUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingNewResponseLogicOrUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingNewResponseLogicOrUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingNewResponseLogicOrUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingNewResponseLogicOrUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MappingNewResponseLogicOrUnion) RawJSON() string { return u.JSON.raw }

func (r *MappingNewResponseLogicOrUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingGetResponse struct {
	ID               string                      `json:"id" api:"required"`
	IsEnabled        bool                        `json:"isEnabled" api:"required"`
	Mappings         []MappingGetResponseMapping `json:"mappings" api:"required"`
	DestinationID    string                      `json:"destinationId" api:"nullable"`
	IsDefaultMapping bool                        `json:"isDefaultMapping" api:"nullable"`
	// Condition tree gating when this mapping fires. A node is either a leaf
	// `condition` or a combinator (`AND`, `OR`, `NOT`). Combinator children are
	// themselves logic nodes, so trees nest arbitrarily.
	//
	// Example leaf:
	// `{ "condition": { "property": "$event.event", "operator": "Is", "value": "page_view" } }`.
	//
	// Example combinator: `{ "AND": [{ "condition": ... }, { "OR": [...] }] }`.
	Logic        MappingGetResponseLogic `json:"logic" api:"nullable"`
	Name         string                  `json:"name" api:"nullable"`
	Priority     float64                 `json:"priority" api:"nullable"`
	SourceID     string                  `json:"sourceId" api:"nullable"`
	TemplateID   string                  `json:"templateId" api:"nullable"`
	TemplateName string                  `json:"templateName" api:"nullable"`
	UpdatedAt    string                  `json:"updatedAt" api:"nullable"`
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
func (r MappingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MappingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingGetResponseMapping struct {
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
func (r MappingGetResponseMapping) RawJSON() string { return r.JSON.raw }
func (r *MappingGetResponseMapping) UnmarshalJSON(data []byte) error {
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
type MappingGetResponseLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And       []map[string]MappingGetResponseLogicAndUnion `json:"AND" api:"nullable"`
	Condition MappingGetResponseLogicCondition             `json:"condition" api:"nullable"`
	// Negates a single child logic node.
	Not map[string]MappingGetResponseLogicNotUnion `json:"NOT" api:"nullable"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or []map[string]MappingGetResponseLogicOrUnion `json:"OR" api:"nullable"`
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
func (r MappingGetResponseLogic) RawJSON() string { return r.JSON.raw }
func (r *MappingGetResponseLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MappingGetResponseLogicAndUnion contains all possible properties and values from
// [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfMappingGetResponseLogicAndMapItem]
type MappingGetResponseLogicAndUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfMappingGetResponseLogicAndMapItem any `json:",inline"`
	JSON                                struct {
		OfString                            respjson.Field
		OfFloat                             respjson.Field
		OfBool                              respjson.Field
		OfAnyArray                          respjson.Field
		OfMappingGetResponseLogicAndMapItem respjson.Field
		raw                                 string
	} `json:"-"`
}

func (u MappingGetResponseLogicAndUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingGetResponseLogicAndUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingGetResponseLogicAndUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingGetResponseLogicAndUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingGetResponseLogicAndUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MappingGetResponseLogicAndUnion) RawJSON() string { return u.JSON.raw }

func (r *MappingGetResponseLogicAndUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingGetResponseLogicCondition struct {
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
func (r MappingGetResponseLogicCondition) RawJSON() string { return r.JSON.raw }
func (r *MappingGetResponseLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MappingGetResponseLogicNotUnion contains all possible properties and values from
// [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfMappingGetResponseLogicNotMapItem]
type MappingGetResponseLogicNotUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfMappingGetResponseLogicNotMapItem any `json:",inline"`
	JSON                                struct {
		OfString                            respjson.Field
		OfFloat                             respjson.Field
		OfBool                              respjson.Field
		OfAnyArray                          respjson.Field
		OfMappingGetResponseLogicNotMapItem respjson.Field
		raw                                 string
	} `json:"-"`
}

func (u MappingGetResponseLogicNotUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingGetResponseLogicNotUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingGetResponseLogicNotUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingGetResponseLogicNotUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingGetResponseLogicNotUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MappingGetResponseLogicNotUnion) RawJSON() string { return u.JSON.raw }

func (r *MappingGetResponseLogicNotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MappingGetResponseLogicOrUnion contains all possible properties and values from
// [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfMappingGetResponseLogicOrMapItem]
type MappingGetResponseLogicOrUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfMappingGetResponseLogicOrMapItem any `json:",inline"`
	JSON                               struct {
		OfString                           respjson.Field
		OfFloat                            respjson.Field
		OfBool                             respjson.Field
		OfAnyArray                         respjson.Field
		OfMappingGetResponseLogicOrMapItem respjson.Field
		raw                                string
	} `json:"-"`
}

func (u MappingGetResponseLogicOrUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingGetResponseLogicOrUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingGetResponseLogicOrUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingGetResponseLogicOrUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingGetResponseLogicOrUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MappingGetResponseLogicOrUnion) RawJSON() string { return u.JSON.raw }

func (r *MappingGetResponseLogicOrUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingUpdateResponse struct {
	ID               string                         `json:"id" api:"required"`
	IsEnabled        bool                           `json:"isEnabled" api:"required"`
	Mappings         []MappingUpdateResponseMapping `json:"mappings" api:"required"`
	DestinationID    string                         `json:"destinationId" api:"nullable"`
	IsDefaultMapping bool                           `json:"isDefaultMapping" api:"nullable"`
	// Condition tree gating when this mapping fires. A node is either a leaf
	// `condition` or a combinator (`AND`, `OR`, `NOT`). Combinator children are
	// themselves logic nodes, so trees nest arbitrarily.
	//
	// Example leaf:
	// `{ "condition": { "property": "$event.event", "operator": "Is", "value": "page_view" } }`.
	//
	// Example combinator: `{ "AND": [{ "condition": ... }, { "OR": [...] }] }`.
	Logic        MappingUpdateResponseLogic `json:"logic" api:"nullable"`
	Name         string                     `json:"name" api:"nullable"`
	Priority     float64                    `json:"priority" api:"nullable"`
	SourceID     string                     `json:"sourceId" api:"nullable"`
	TemplateID   string                     `json:"templateId" api:"nullable"`
	TemplateName string                     `json:"templateName" api:"nullable"`
	UpdatedAt    string                     `json:"updatedAt" api:"nullable"`
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
func (r MappingUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *MappingUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingUpdateResponseMapping struct {
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
func (r MappingUpdateResponseMapping) RawJSON() string { return r.JSON.raw }
func (r *MappingUpdateResponseMapping) UnmarshalJSON(data []byte) error {
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
type MappingUpdateResponseLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And       []map[string]MappingUpdateResponseLogicAndUnion `json:"AND" api:"nullable"`
	Condition MappingUpdateResponseLogicCondition             `json:"condition" api:"nullable"`
	// Negates a single child logic node.
	Not map[string]MappingUpdateResponseLogicNotUnion `json:"NOT" api:"nullable"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or []map[string]MappingUpdateResponseLogicOrUnion `json:"OR" api:"nullable"`
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
func (r MappingUpdateResponseLogic) RawJSON() string { return r.JSON.raw }
func (r *MappingUpdateResponseLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MappingUpdateResponseLogicAndUnion contains all possible properties and values
// from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfMappingUpdateResponseLogicAndMapItem]
type MappingUpdateResponseLogicAndUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfMappingUpdateResponseLogicAndMapItem any `json:",inline"`
	JSON                                   struct {
		OfString                               respjson.Field
		OfFloat                                respjson.Field
		OfBool                                 respjson.Field
		OfAnyArray                             respjson.Field
		OfMappingUpdateResponseLogicAndMapItem respjson.Field
		raw                                    string
	} `json:"-"`
}

func (u MappingUpdateResponseLogicAndUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingUpdateResponseLogicAndUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingUpdateResponseLogicAndUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingUpdateResponseLogicAndUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingUpdateResponseLogicAndUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MappingUpdateResponseLogicAndUnion) RawJSON() string { return u.JSON.raw }

func (r *MappingUpdateResponseLogicAndUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingUpdateResponseLogicCondition struct {
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
func (r MappingUpdateResponseLogicCondition) RawJSON() string { return r.JSON.raw }
func (r *MappingUpdateResponseLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MappingUpdateResponseLogicNotUnion contains all possible properties and values
// from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfMappingUpdateResponseLogicNotMapItem]
type MappingUpdateResponseLogicNotUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfMappingUpdateResponseLogicNotMapItem any `json:",inline"`
	JSON                                   struct {
		OfString                               respjson.Field
		OfFloat                                respjson.Field
		OfBool                                 respjson.Field
		OfAnyArray                             respjson.Field
		OfMappingUpdateResponseLogicNotMapItem respjson.Field
		raw                                    string
	} `json:"-"`
}

func (u MappingUpdateResponseLogicNotUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingUpdateResponseLogicNotUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingUpdateResponseLogicNotUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingUpdateResponseLogicNotUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingUpdateResponseLogicNotUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MappingUpdateResponseLogicNotUnion) RawJSON() string { return u.JSON.raw }

func (r *MappingUpdateResponseLogicNotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MappingUpdateResponseLogicOrUnion contains all possible properties and values
// from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfMappingUpdateResponseLogicOrMapItem]
type MappingUpdateResponseLogicOrUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfMappingUpdateResponseLogicOrMapItem any `json:",inline"`
	JSON                                  struct {
		OfString                              respjson.Field
		OfFloat                               respjson.Field
		OfBool                                respjson.Field
		OfAnyArray                            respjson.Field
		OfMappingUpdateResponseLogicOrMapItem respjson.Field
		raw                                   string
	} `json:"-"`
}

func (u MappingUpdateResponseLogicOrUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingUpdateResponseLogicOrUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingUpdateResponseLogicOrUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingUpdateResponseLogicOrUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingUpdateResponseLogicOrUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MappingUpdateResponseLogicOrUnion) RawJSON() string { return u.JSON.raw }

func (r *MappingUpdateResponseLogicOrUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingRenamePreviewResponse struct {
	AffectedEventNames []string `json:"affectedEventNames" api:"required"`
	WillCascade        bool     `json:"willCascade" api:"required"`
	MapperName         string   `json:"mapperName" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AffectedEventNames respjson.Field
		WillCascade        respjson.Field
		MapperName         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MappingRenamePreviewResponse) RawJSON() string { return r.JSON.raw }
func (r *MappingRenamePreviewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingRenameResponse struct {
	ID               string                         `json:"id" api:"required"`
	IsEnabled        bool                           `json:"isEnabled" api:"required"`
	Mappings         []MappingRenameResponseMapping `json:"mappings" api:"required"`
	DestinationID    string                         `json:"destinationId" api:"nullable"`
	IsDefaultMapping bool                           `json:"isDefaultMapping" api:"nullable"`
	// Condition tree gating when this mapping fires. A node is either a leaf
	// `condition` or a combinator (`AND`, `OR`, `NOT`). Combinator children are
	// themselves logic nodes, so trees nest arbitrarily.
	//
	// Example leaf:
	// `{ "condition": { "property": "$event.event", "operator": "Is", "value": "page_view" } }`.
	//
	// Example combinator: `{ "AND": [{ "condition": ... }, { "OR": [...] }] }`.
	Logic        MappingRenameResponseLogic `json:"logic" api:"nullable"`
	Name         string                     `json:"name" api:"nullable"`
	Priority     float64                    `json:"priority" api:"nullable"`
	SourceID     string                     `json:"sourceId" api:"nullable"`
	TemplateID   string                     `json:"templateId" api:"nullable"`
	TemplateName string                     `json:"templateName" api:"nullable"`
	UpdatedAt    string                     `json:"updatedAt" api:"nullable"`
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
func (r MappingRenameResponse) RawJSON() string { return r.JSON.raw }
func (r *MappingRenameResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingRenameResponseMapping struct {
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
func (r MappingRenameResponseMapping) RawJSON() string { return r.JSON.raw }
func (r *MappingRenameResponseMapping) UnmarshalJSON(data []byte) error {
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
type MappingRenameResponseLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And       []map[string]MappingRenameResponseLogicAndUnion `json:"AND" api:"nullable"`
	Condition MappingRenameResponseLogicCondition             `json:"condition" api:"nullable"`
	// Negates a single child logic node.
	Not map[string]MappingRenameResponseLogicNotUnion `json:"NOT" api:"nullable"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or []map[string]MappingRenameResponseLogicOrUnion `json:"OR" api:"nullable"`
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
func (r MappingRenameResponseLogic) RawJSON() string { return r.JSON.raw }
func (r *MappingRenameResponseLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MappingRenameResponseLogicAndUnion contains all possible properties and values
// from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfMappingRenameResponseLogicAndMapItem]
type MappingRenameResponseLogicAndUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfMappingRenameResponseLogicAndMapItem any `json:",inline"`
	JSON                                   struct {
		OfString                               respjson.Field
		OfFloat                                respjson.Field
		OfBool                                 respjson.Field
		OfAnyArray                             respjson.Field
		OfMappingRenameResponseLogicAndMapItem respjson.Field
		raw                                    string
	} `json:"-"`
}

func (u MappingRenameResponseLogicAndUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingRenameResponseLogicAndUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingRenameResponseLogicAndUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingRenameResponseLogicAndUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingRenameResponseLogicAndUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MappingRenameResponseLogicAndUnion) RawJSON() string { return u.JSON.raw }

func (r *MappingRenameResponseLogicAndUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingRenameResponseLogicCondition struct {
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
func (r MappingRenameResponseLogicCondition) RawJSON() string { return r.JSON.raw }
func (r *MappingRenameResponseLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MappingRenameResponseLogicNotUnion contains all possible properties and values
// from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfMappingRenameResponseLogicNotMapItem]
type MappingRenameResponseLogicNotUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfMappingRenameResponseLogicNotMapItem any `json:",inline"`
	JSON                                   struct {
		OfString                               respjson.Field
		OfFloat                                respjson.Field
		OfBool                                 respjson.Field
		OfAnyArray                             respjson.Field
		OfMappingRenameResponseLogicNotMapItem respjson.Field
		raw                                    string
	} `json:"-"`
}

func (u MappingRenameResponseLogicNotUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingRenameResponseLogicNotUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingRenameResponseLogicNotUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingRenameResponseLogicNotUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingRenameResponseLogicNotUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MappingRenameResponseLogicNotUnion) RawJSON() string { return u.JSON.raw }

func (r *MappingRenameResponseLogicNotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MappingRenameResponseLogicOrUnion contains all possible properties and values
// from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfMappingRenameResponseLogicOrMapItem]
type MappingRenameResponseLogicOrUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfMappingRenameResponseLogicOrMapItem any `json:",inline"`
	JSON                                  struct {
		OfString                              respjson.Field
		OfFloat                               respjson.Field
		OfBool                                respjson.Field
		OfAnyArray                            respjson.Field
		OfMappingRenameResponseLogicOrMapItem respjson.Field
		raw                                   string
	} `json:"-"`
}

func (u MappingRenameResponseLogicOrUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingRenameResponseLogicOrUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingRenameResponseLogicOrUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingRenameResponseLogicOrUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingRenameResponseLogicOrUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MappingRenameResponseLogicOrUnion) RawJSON() string { return u.JSON.raw }

func (r *MappingRenameResponseLogicOrUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingReorderResponse struct {
	Entities []MappingReorderResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MappingReorderResponse) RawJSON() string { return r.JSON.raw }
func (r *MappingReorderResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingReorderResponseEntity struct {
	ID               string                                `json:"id" api:"required"`
	IsEnabled        bool                                  `json:"isEnabled" api:"required"`
	Mappings         []MappingReorderResponseEntityMapping `json:"mappings" api:"required"`
	DestinationID    string                                `json:"destinationId" api:"nullable"`
	IsDefaultMapping bool                                  `json:"isDefaultMapping" api:"nullable"`
	// Condition tree gating when this mapping fires. A node is either a leaf
	// `condition` or a combinator (`AND`, `OR`, `NOT`). Combinator children are
	// themselves logic nodes, so trees nest arbitrarily.
	//
	// Example leaf:
	// `{ "condition": { "property": "$event.event", "operator": "Is", "value": "page_view" } }`.
	//
	// Example combinator: `{ "AND": [{ "condition": ... }, { "OR": [...] }] }`.
	Logic        MappingReorderResponseEntityLogic `json:"logic" api:"nullable"`
	Name         string                            `json:"name" api:"nullable"`
	Priority     float64                           `json:"priority" api:"nullable"`
	SourceID     string                            `json:"sourceId" api:"nullable"`
	TemplateID   string                            `json:"templateId" api:"nullable"`
	TemplateName string                            `json:"templateName" api:"nullable"`
	UpdatedAt    string                            `json:"updatedAt" api:"nullable"`
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
func (r MappingReorderResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *MappingReorderResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingReorderResponseEntityMapping struct {
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
func (r MappingReorderResponseEntityMapping) RawJSON() string { return r.JSON.raw }
func (r *MappingReorderResponseEntityMapping) UnmarshalJSON(data []byte) error {
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
type MappingReorderResponseEntityLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And       []map[string]MappingReorderResponseEntityLogicAndUnion `json:"AND" api:"nullable"`
	Condition MappingReorderResponseEntityLogicCondition             `json:"condition" api:"nullable"`
	// Negates a single child logic node.
	Not map[string]MappingReorderResponseEntityLogicNotUnion `json:"NOT" api:"nullable"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or []map[string]MappingReorderResponseEntityLogicOrUnion `json:"OR" api:"nullable"`
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
func (r MappingReorderResponseEntityLogic) RawJSON() string { return r.JSON.raw }
func (r *MappingReorderResponseEntityLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MappingReorderResponseEntityLogicAndUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfMappingReorderResponseEntityLogicAndMapItem]
type MappingReorderResponseEntityLogicAndUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfMappingReorderResponseEntityLogicAndMapItem any `json:",inline"`
	JSON                                          struct {
		OfString                                      respjson.Field
		OfFloat                                       respjson.Field
		OfBool                                        respjson.Field
		OfAnyArray                                    respjson.Field
		OfMappingReorderResponseEntityLogicAndMapItem respjson.Field
		raw                                           string
	} `json:"-"`
}

func (u MappingReorderResponseEntityLogicAndUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingReorderResponseEntityLogicAndUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingReorderResponseEntityLogicAndUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingReorderResponseEntityLogicAndUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingReorderResponseEntityLogicAndUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MappingReorderResponseEntityLogicAndUnion) RawJSON() string { return u.JSON.raw }

func (r *MappingReorderResponseEntityLogicAndUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingReorderResponseEntityLogicCondition struct {
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
func (r MappingReorderResponseEntityLogicCondition) RawJSON() string { return r.JSON.raw }
func (r *MappingReorderResponseEntityLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MappingReorderResponseEntityLogicNotUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfMappingReorderResponseEntityLogicNotMapItem]
type MappingReorderResponseEntityLogicNotUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfMappingReorderResponseEntityLogicNotMapItem any `json:",inline"`
	JSON                                          struct {
		OfString                                      respjson.Field
		OfFloat                                       respjson.Field
		OfBool                                        respjson.Field
		OfAnyArray                                    respjson.Field
		OfMappingReorderResponseEntityLogicNotMapItem respjson.Field
		raw                                           string
	} `json:"-"`
}

func (u MappingReorderResponseEntityLogicNotUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingReorderResponseEntityLogicNotUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingReorderResponseEntityLogicNotUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingReorderResponseEntityLogicNotUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingReorderResponseEntityLogicNotUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MappingReorderResponseEntityLogicNotUnion) RawJSON() string { return u.JSON.raw }

func (r *MappingReorderResponseEntityLogicNotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MappingReorderResponseEntityLogicOrUnion contains all possible properties and
// values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfMappingReorderResponseEntityLogicOrMapItem]
type MappingReorderResponseEntityLogicOrUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfMappingReorderResponseEntityLogicOrMapItem any `json:",inline"`
	JSON                                         struct {
		OfString                                     respjson.Field
		OfFloat                                      respjson.Field
		OfBool                                       respjson.Field
		OfAnyArray                                   respjson.Field
		OfMappingReorderResponseEntityLogicOrMapItem respjson.Field
		raw                                          string
	} `json:"-"`
}

func (u MappingReorderResponseEntityLogicOrUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingReorderResponseEntityLogicOrUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingReorderResponseEntityLogicOrUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingReorderResponseEntityLogicOrUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MappingReorderResponseEntityLogicOrUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MappingReorderResponseEntityLogicOrUnion) RawJSON() string { return u.JSON.raw }

func (r *MappingReorderResponseEntityLogicOrUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingTemplatesResponse struct {
	Entities []MappingTemplatesResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MappingTemplatesResponse) RawJSON() string { return r.JSON.raw }
func (r *MappingTemplatesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingTemplatesResponseEntity struct {
	// Template identifier — pass to `POST /rest/v1/mappings` as `templateId`.
	ID string `json:"id" api:"required"`
	// True for the destination's built-in default template (the one stored at
	// `MAPPER#{destinationId}` when configured). Sources only have one template; it is
	// always default.
	IsDefault   bool                                    `json:"isDefault" api:"required"`
	Mappings    []MappingTemplatesResponseEntityMapping `json:"mappings" api:"required"`
	Name        string                                  `json:"name" api:"required"`
	Description string                                  `json:"description" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		IsDefault   respjson.Field
		Mappings    respjson.Field
		Name        respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MappingTemplatesResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *MappingTemplatesResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingTemplatesResponseEntityMapping struct {
	// Long-form description / tooltip for this property.
	Description string `json:"description" api:"required"`
	IsPii       bool   `json:"isPII" api:"required"`
	// Type information for SDK validation (Text, Integer, Email, Url, IP, Object,
	// KnownObject, Date, DateTime, Array, Boolean, JSON).
	//
	// Any of "Array", "Boolean", "Date", "DateTime", "Email", "IP", "Integer", "JSON",
	// "KnownObject", "Object", "Text", "Url".
	Kind string `json:"kind" api:"required"`
	// Human-readable label (e.g. "Email", "Event Name").
	Label string `json:"label" api:"required"`
	// The template default source expression, e.g. `{{visitor.email}}`.
	Map string `json:"map" api:"required"`
	// The value to send as `mappings[].property` when creating or patching a mapping.
	Property string `json:"property" api:"required"`
	Required bool   `json:"required" api:"required"`
	// The template default modification (hashing / case / URL truncation).
	//
	// Any of "CamelCase", "DmaIP", "DomainOnly", "DomainPathOnly", "DomainPathUTMs",
	// "DomainUTMs", "FakeDomain", "FakeDomainRealPath", "FakeIP", "FullUrl", "Hash",
	// "HashMD5", "HashedCountry", "HashedDateOfBirth", "HashedGender",
	// "HashedNormalized", "HashedNormalizedNoSpecialChars", "HashedPhone",
	// "HashedState", "HashedZip", "KebabCase", "LowerCase", "None", "Null",
	// "Redacted", "RegionalIP", "SnakeCase", "StartCase", "UpperCase".
	Modification string `json:"modification" api:"nullable"`
	// Suggested modification options for this property. Not a whitelist.
	//
	// Any of "CamelCase", "DmaIP", "DomainOnly", "DomainPathOnly", "DomainPathUTMs",
	// "DomainUTMs", "FakeDomain", "FakeDomainRealPath", "FakeIP", "FullUrl", "Hash",
	// "HashMD5", "HashedCountry", "HashedDateOfBirth", "HashedGender",
	// "HashedNormalized", "HashedNormalizedNoSpecialChars", "HashedPhone",
	// "HashedState", "HashedZip", "KebabCase", "LowerCase", "None", "Null",
	// "Redacted", "RegionalIP", "SnakeCase", "StartCase", "UpperCase".
	ModificationOptions []string `json:"modificationOptions" api:"nullable"`
	// When set, the ONLY valid `map` values for this property. Typically used for
	// enum-shaped destinations.
	Options []MappingTemplatesResponseEntityMappingOption `json:"options" api:"nullable"`
	// Non-binding suggestions for the `map` value (e.g. common event names a customer
	// might want to use).
	SuggestedOptions []MappingTemplatesResponseEntityMappingSuggestedOption `json:"suggestedOptions" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description         respjson.Field
		IsPii               respjson.Field
		Kind                respjson.Field
		Label               respjson.Field
		Map                 respjson.Field
		Property            respjson.Field
		Required            respjson.Field
		Modification        respjson.Field
		ModificationOptions respjson.Field
		Options             respjson.Field
		SuggestedOptions    respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MappingTemplatesResponseEntityMapping) RawJSON() string { return r.JSON.raw }
func (r *MappingTemplatesResponseEntityMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingTemplatesResponseEntityMappingOption struct {
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
func (r MappingTemplatesResponseEntityMappingOption) RawJSON() string { return r.JSON.raw }
func (r *MappingTemplatesResponseEntityMappingOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingTemplatesResponseEntityMappingSuggestedOption struct {
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
func (r MappingTemplatesResponseEntityMappingSuggestedOption) RawJSON() string { return r.JSON.raw }
func (r *MappingTemplatesResponseEntityMappingSuggestedOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingDefaultVariablesResponse struct {
	Entities []MappingDefaultVariablesResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MappingDefaultVariablesResponse) RawJSON() string { return r.JSON.raw }
func (r *MappingDefaultVariablesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingDefaultVariablesResponseEntity struct {
	// Sample values observed for this path (empty for unsampled defaults).
	Examples []string `json:"examples" api:"required"`
	// Human-readable display name.
	Name string `json:"name" api:"required"`
	// Dot-path used in mapping `value` fields (e.g. `event.email`).
	Path string `json:"path" api:"required"`
	// Relative popularity rank. Higher means more frequently set across events.
	Popularity float64 `json:"popularity" api:"required"`
	// Optional long-form context shown in the variable dictionary drawer.
	AdvancedInfo string `json:"advancedInfo" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Examples     respjson.Field
		Name         respjson.Field
		Path         respjson.Field
		Popularity   respjson.Field
		AdvancedInfo respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MappingDefaultVariablesResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *MappingDefaultVariablesResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingCustomVariablesResponse struct {
	Entities []MappingCustomVariablesResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MappingCustomVariablesResponse) RawJSON() string { return r.JSON.raw }
func (r *MappingCustomVariablesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingCustomVariablesResponseEntity struct {
	// Sample values observed for this path (empty for unsampled defaults).
	Examples []string `json:"examples" api:"required"`
	// Human-readable display name.
	Name string `json:"name" api:"required"`
	// Dot-path used in mapping `value` fields (e.g. `event.email`).
	Path string `json:"path" api:"required"`
	// Relative popularity rank. Higher means more frequently set across events.
	Popularity float64 `json:"popularity" api:"required"`
	// Optional long-form context shown in the variable dictionary drawer.
	AdvancedInfo string `json:"advancedInfo" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Examples     respjson.Field
		Name         respjson.Field
		Path         respjson.Field
		Popularity   respjson.Field
		AdvancedInfo respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MappingCustomVariablesResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *MappingCustomVariablesResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingModificationsResponse struct {
	Entities []MappingModificationsResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MappingModificationsResponse) RawJSON() string { return r.JSON.raw }
func (r *MappingModificationsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingModificationsResponseEntity struct {
	// One-sentence explanation of what the modification does to the mapped value.
	Description string `json:"description" api:"required"`
	// Short human-readable name (suitable for picker labels).
	Label string `json:"label" api:"required"`
	// Enum value to send on `modification` fields when authoring a mapping.
	//
	// Any of "CamelCase", "DmaIP", "DomainOnly", "DomainPathOnly", "DomainPathUTMs",
	// "DomainUTMs", "FakeDomain", "FakeDomainRealPath", "FakeIP", "FullUrl", "Hash",
	// "HashMD5", "HashedCountry", "HashedDateOfBirth", "HashedGender",
	// "HashedNormalized", "HashedNormalizedNoSpecialChars", "HashedPhone",
	// "HashedState", "HashedZip", "KebabCase", "LowerCase", "None", "Null",
	// "Redacted", "RegionalIP", "SnakeCase", "StartCase", "UpperCase".
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Label       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MappingModificationsResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *MappingModificationsResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingListParams struct {
	// Filter mappings by their parent entity id. Must be a destination id or source
	// id.
	EntityID string `query:"entityId" api:"required" json:"-"`
	// Maximum number of mappings to return. Defaults to 1000; values below 1 are
	// clamped to 1 and values above 1000 are clamped to 1000. Most accounts can fetch
	// the full list in one request.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Opaque pagination cursor from pagination.nextCursor in the previous response. Do
	// not decode or modify it. Malformed cursors return 400 Bad Request.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MappingListParams]'s query parameters as `url.Values`.
func (r MappingListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MappingNewParams struct {
	// Template fat-create only. Override the auto-generated mapping name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Quick-create variant: allowed event to bind the new mapping to. Required
	// together with `destinationId`. Mutually exclusive with `templateId`/`entityId`.
	AllowedEventID param.Opt[string] `json:"allowedEventId,omitzero"`
	// Quick-create variant: destination that should receive events matched by this
	// mapping. Required together with `allowedEventId`.
	DestinationID param.Opt[string] `json:"destinationId,omitzero"`
	// Template fat-create variant: destination or source id this mapping belongs to.
	// Required together with `templateId`.
	EntityID param.Opt[string] `json:"entityId,omitzero"`
	// Template fat-create only. Zero-based position in the destination/source priority
	// order to insert this mapping after. Omit to append at the end.
	InsertAfterIdx param.Opt[int64] `json:"insertAfterIdx,omitzero"`
	// Template fat-create only. Initial enabled state. Defaults to `true`.
	IsEnabled param.Opt[bool] `json:"isEnabled,omitzero"`
	// Template fat-create variant: template id from `GET /rest/v1/mapping-templates`.
	// Picks the property descriptor set used to validate `mappings[].property`.
	// Required together with `entityId`.
	TemplateID param.Opt[string] `json:"templateId,omitzero"`
	// Condition tree gating when this mapping fires. A node is either a leaf
	// `condition` or a combinator (`AND`, `OR`, `NOT`). Combinator children are
	// themselves logic nodes, so trees nest arbitrarily.
	//
	// Example leaf:
	// `{ "condition": { "property": "$event.event", "operator": "Is", "value": "page_view" } }`.
	//
	// Example combinator: `{ "AND": [{ "condition": ... }, { "OR": [...] }] }`.
	Logic MappingNewParamsLogic `json:"logic,omitzero"`
	// Template fat-create only. Optional initial property mappings. When omitted the
	// mapping is seeded with template defaults for sources and non-default destination
	// templates, and with `[]` for default destination templates.
	Mappings []MappingNewParamsMapping `json:"mappings,omitzero"`
	paramObj
}

func (r MappingNewParams) MarshalJSON() (data []byte, err error) {
	type shadow MappingNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MappingNewParams) UnmarshalJSON(data []byte) error {
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
type MappingNewParamsLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And       []map[string]MappingNewParamsLogicAndUnion `json:"AND,omitzero"`
	Condition MappingNewParamsLogicCondition             `json:"condition,omitzero"`
	// Negates a single child logic node.
	Not map[string]MappingNewParamsLogicNotUnion `json:"NOT,omitzero"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or []map[string]MappingNewParamsLogicOrUnion `json:"OR,omitzero"`
	paramObj
}

func (r MappingNewParamsLogic) MarshalJSON() (data []byte, err error) {
	type shadow MappingNewParamsLogic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MappingNewParamsLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MappingNewParamsLogicAndUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u MappingNewParamsLogicAndUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *MappingNewParamsLogicAndUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Operator, Property, Value are required.
type MappingNewParamsLogicCondition struct {
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
	Operator string `json:"operator,omitzero" api:"required"`
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
	paramObj
}

func (r MappingNewParamsLogicCondition) MarshalJSON() (data []byte, err error) {
	type shadow MappingNewParamsLogicCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MappingNewParamsLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MappingNewParamsLogicCondition](
		"operator", "Is", "IsNot", "Contains", "DoesNotContain", "StartsWith", "EndsWith", "IsFalsy", "IsTruthy", "IsNull", "IsNotNull", "IsUndefined", "IsNotUndefined", "IsGreaterThan", "IsGreaterThanOrEqual", "IsLessThan", "IsLessThanOrEqual", "IsIn", "IsNotIn", "IsFoundIn", "IsNotFoundIn", "IsTrue", "IsFalse", "IsBefore", "IsAfter", "IsBetween", "IsOnOrBefore", "IsOnOrAfter", "MatchesRegex", "MatchesRegexIgnoreCase", "DoesNotMatchRegex", "DoesNotMatchRegexIgnoreCase",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MappingNewParamsLogicNotUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u MappingNewParamsLogicNotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *MappingNewParamsLogicNotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MappingNewParamsLogicOrUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u MappingNewParamsLogicOrUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *MappingNewParamsLogicOrUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Map, Property are required.
type MappingNewParamsMapping struct {
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

func (r MappingNewParamsMapping) MarshalJSON() (data []byte, err error) {
	type shadow MappingNewParamsMapping
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MappingNewParamsMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MappingNewParamsMapping](
		"modification", "CamelCase", "DmaIP", "DomainOnly", "DomainPathOnly", "DomainPathUTMs", "DomainUTMs", "FakeDomain", "FakeDomainRealPath", "FakeIP", "FullUrl", "Hash", "HashMD5", "HashedCountry", "HashedDateOfBirth", "HashedGender", "HashedNormalized", "HashedNormalizedNoSpecialChars", "HashedPhone", "HashedState", "HashedZip", "KebabCase", "LowerCase", "None", "Null", "Redacted", "RegionalIP", "SnakeCase", "StartCase", "UpperCase",
	)
}

type MappingUpdateParams struct {
	// Flip the mapping on/off without changing other fields. `null` is treated as
	// omitted.
	IsEnabled param.Opt[bool]   `json:"isEnabled,omitzero"`
	Name      param.Opt[string] `json:"name,omitzero"`
	// Condition tree gating when this mapping fires. A node is either a leaf
	// `condition` or a combinator (`AND`, `OR`, `NOT`). Combinator children are
	// themselves logic nodes, so trees nest arbitrarily.
	//
	// Example leaf:
	// `{ "condition": { "property": "$event.event", "operator": "Is", "value": "page_view" } }`.
	//
	// Example combinator: `{ "AND": [{ "condition": ... }, { "OR": [...] }] }`.
	Logic    MappingUpdateParamsLogic     `json:"logic,omitzero"`
	Mappings []MappingUpdateParamsMapping `json:"mappings,omitzero"`
	paramObj
}

func (r MappingUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow MappingUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MappingUpdateParams) UnmarshalJSON(data []byte) error {
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
type MappingUpdateParamsLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And       []map[string]MappingUpdateParamsLogicAndUnion `json:"AND,omitzero"`
	Condition MappingUpdateParamsLogicCondition             `json:"condition,omitzero"`
	// Negates a single child logic node.
	Not map[string]MappingUpdateParamsLogicNotUnion `json:"NOT,omitzero"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or []map[string]MappingUpdateParamsLogicOrUnion `json:"OR,omitzero"`
	paramObj
}

func (r MappingUpdateParamsLogic) MarshalJSON() (data []byte, err error) {
	type shadow MappingUpdateParamsLogic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MappingUpdateParamsLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MappingUpdateParamsLogicAndUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u MappingUpdateParamsLogicAndUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *MappingUpdateParamsLogicAndUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Operator, Property, Value are required.
type MappingUpdateParamsLogicCondition struct {
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
	Operator string `json:"operator,omitzero" api:"required"`
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
	paramObj
}

func (r MappingUpdateParamsLogicCondition) MarshalJSON() (data []byte, err error) {
	type shadow MappingUpdateParamsLogicCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MappingUpdateParamsLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MappingUpdateParamsLogicCondition](
		"operator", "Is", "IsNot", "Contains", "DoesNotContain", "StartsWith", "EndsWith", "IsFalsy", "IsTruthy", "IsNull", "IsNotNull", "IsUndefined", "IsNotUndefined", "IsGreaterThan", "IsGreaterThanOrEqual", "IsLessThan", "IsLessThanOrEqual", "IsIn", "IsNotIn", "IsFoundIn", "IsNotFoundIn", "IsTrue", "IsFalse", "IsBefore", "IsAfter", "IsBetween", "IsOnOrBefore", "IsOnOrAfter", "MatchesRegex", "MatchesRegexIgnoreCase", "DoesNotMatchRegex", "DoesNotMatchRegexIgnoreCase",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MappingUpdateParamsLogicNotUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u MappingUpdateParamsLogicNotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *MappingUpdateParamsLogicNotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MappingUpdateParamsLogicOrUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u MappingUpdateParamsLogicOrUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *MappingUpdateParamsLogicOrUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Map, Property are required.
type MappingUpdateParamsMapping struct {
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

func (r MappingUpdateParamsMapping) MarshalJSON() (data []byte, err error) {
	type shadow MappingUpdateParamsMapping
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MappingUpdateParamsMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MappingUpdateParamsMapping](
		"modification", "CamelCase", "DmaIP", "DomainOnly", "DomainPathOnly", "DomainPathUTMs", "DomainUTMs", "FakeDomain", "FakeDomainRealPath", "FakeIP", "FullUrl", "Hash", "HashMD5", "HashedCountry", "HashedDateOfBirth", "HashedGender", "HashedNormalized", "HashedNormalizedNoSpecialChars", "HashedPhone", "HashedState", "HashedZip", "KebabCase", "LowerCase", "None", "Null", "Redacted", "RegionalIP", "SnakeCase", "StartCase", "UpperCase",
	)
}

type MappingRenamePreviewParams struct {
	// Allowed event id to preview for this destination.
	AllowedEventID string `query:"allowedEventId" api:"required" format:"uuid" json:"-"`
	// Destination id. Discover renaming support through
	// `GET /rest/v1/destinations/types`.
	DestinationID string `query:"destinationId" api:"required" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [MappingRenamePreviewParams]'s query parameters as
// `url.Values`.
func (r MappingRenamePreviewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MappingRenameParams struct {
	// Name sent to the destination. Send `null` to clear the rename.
	RenamedEventName param.Opt[string] `json:"renamedEventName,omitzero" api:"required"`
	// Allowed event id to rename for this destination.
	AllowedEventID string `json:"allowedEventId" api:"required" format:"uuid"`
	// Destination id for the rename.
	DestinationID string `json:"destinationId" api:"required" format:"uuid"`
	// Set to `true` only after preview confirms which allowed events a shared mapper
	// will change. Defaults to `false`.
	ConfirmCascade param.Opt[bool] `json:"confirmCascade,omitzero"`
	paramObj
}

func (r MappingRenameParams) MarshalJSON() (data []byte, err error) {
	type shadow MappingRenameParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MappingRenameParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingReorderParams struct {
	// Mapping ids in their new priority order, low priority index first. All ids must
	// belong to the same parent entity (source or destination).
	Uuids []string `json:"uuids,omitzero" api:"required"`
	paramObj
}

func (r MappingReorderParams) MarshalJSON() (data []byte, err error) {
	type shadow MappingReorderParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MappingReorderParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingTemplatesParams struct {
	// Destination or source id. Required.
	EntityID string `query:"entityId" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [MappingTemplatesParams]'s query parameters as `url.Values`.
func (r MappingTemplatesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
