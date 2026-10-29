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

// DataGovernanceService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDataGovernanceService] method instead.
type DataGovernanceService struct {
	Options []option.RequestOption
}

// NewDataGovernanceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDataGovernanceService(opts ...option.RequestOption) (r DataGovernanceService) {
	r = DataGovernanceService{}
	r.Options = opts
	return
}

// List data-governance records as { entities, pagination }. Supports limit and
// cursor; pass pagination.nextCursor to retrieve the next page. Records are
// ordered newest first, with stable id ordering for equal creation times. Each
// account currently has at most one record, so entities is empty or contains one
// record. Categories stop dispatch to their destinationIds when their logic
// matches. Requires API-key scope or current OAuth user permission:
// globalDispatch:list
func (r *DataGovernanceService) List(ctx context.Context, query DataGovernanceListParams, opts ...option.RequestOption) (res *pagination.Cursor[DataGovernanceListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "rest/v1/data-governance"
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

// List data-governance records as { entities, pagination }. Supports limit and
// cursor; pass pagination.nextCursor to retrieve the next page. Records are
// ordered newest first, with stable id ordering for equal creation times. Each
// account currently has at most one record, so entities is empty or contains one
// record. Categories stop dispatch to their destinationIds when their logic
// matches. Requires API-key scope or current OAuth user permission:
// globalDispatch:list
func (r *DataGovernanceService) ListAutoPaging(ctx context.Context, query DataGovernanceListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[DataGovernanceListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Create the data-governance record for this account. Each account may have at
// most one — a second POST returns 409. Body is optional; defaults are
// `isEnabled: false` and no categories. Categories are added later via PATCH.
// Requires API-key scope or current OAuth user permission: globalDispatch:create
func (r *DataGovernanceService) New(ctx context.Context, body DataGovernanceNewParams, opts ...option.RequestOption) (res *DataGovernanceNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/data-governance"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Fetch the data-governance record by id, including its categories (logic,
// destinations, priority). Returns the same 404 when the record is missing or
// inaccessible. Requires API-key scope or current OAuth user permission:
// globalDispatch:find
func (r *DataGovernanceService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *DataGovernanceGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/data-governance/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Partially update the data-governance record. Top-level fields (`name`, `notes`,
// `isEnabled`) follow the standard PATCH semantic — only the fields you send are
// changed.
//
// `categories` is the documented exception: when sent, it is **replaced
// wholesale**. There is no partial-merge for individual categories. To change one
// category, fetch the current record, modify the array, and PATCH it back.
//
// Categories are sorted ascending by the `priority` you supplied and then
// renumbered `1..N` in the response, so the returned `priority` values are always
// sequential with no gaps. Stale `destinationIds` (deleted destinations or
// destinations on another account) are silently filtered out — the response echoes
// the filtered list, so a follow-up GET is not required to see what was saved.
// Requires API-key scope or current OAuth user permission: globalDispatch:update
func (r *DataGovernanceService) Update(ctx context.Context, id string, body DataGovernanceUpdateParams, opts ...option.RequestOption) (res *DataGovernanceUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/data-governance/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete the data-governance record from the draft. Publish a version for inbound
// events to flow through without category-level gating. Create a new record with
// POST to reinstate governance. Returns the same 404 when the record is missing or
// inaccessible. Requires API-key scope or current OAuth user permission:
// globalDispatch:delete
func (r *DataGovernanceService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *DataGovernanceDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/data-governance/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type DataGovernanceListResponse struct {
	// Server-assigned UUID for this data-governance record.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp when the record was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// When false, data governance is configured but does not gate dispatch — inbound
	// events flow through to destinations regardless of category logic.
	IsEnabled bool `json:"isEnabled" api:"required"`
	// Discriminator for this resource. Always "dataGovernance".
	Kind string `json:"kind" api:"required"`
	// Governance categories in priority order (1..N).
	Categories []DataGovernanceListResponseCategory `json:"categories" api:"nullable"`
	// Human-readable name shown in the dashboard.
	Name string `json:"name" api:"nullable"`
	// Free-form notes for this record.
	Notes string `json:"notes" api:"nullable"`
	// ISO 8601 timestamp of the last write. Equal to createdAt on a freshly created
	// record; advances on every PATCH.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		IsEnabled   respjson.Field
		Kind        respjson.Field
		Categories  respjson.Field
		Name        respjson.Field
		Notes       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DataGovernanceListResponse) RawJSON() string { return r.JSON.raw }
func (r *DataGovernanceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DataGovernanceListResponseCategory struct {
	// Display name for the category.
	Name string `json:"name" api:"required"`
	// 1-indexed sort position. Always equals (sorted index + 1) — see PATCH for
	// details.
	Priority int64 `json:"priority" api:"required"`
	// Optional human-readable description.
	Description string `json:"description" api:"nullable"`
	// Destinations gated by this category when its logic evaluates to TRUE.
	DestinationIDs []string `json:"destinationIds" api:"nullable"`
	// Condition tree evaluated against each inbound event. Write conditions that
	// evaluate **TRUE for events you want to STOP**. A node is either a leaf
	// `condition` or a combinator (`AND`, `OR`, `NOT`); combinator children are
	// themselves logic nodes, so trees nest arbitrarily.
	//
	// Discovery: `GET /rest/v1/mappings/default-variables` lists the canonical
	// platform-provided `property` paths (visitor consent arrays, event fields,
	// request context, identity fields). Custom `event.event_properties.*` paths are
	// caller-defined.
	//
	// Example leaf (stop dispatch when the visitor rejected the `advertising` consent
	// category):
	// `{ "condition": { "property": "visitor.consent.rejected_categories", "operator": "Contains", "value": "advertising" } }`.
	//
	// Example combinator:
	// `{ "AND": [{ "condition": { "property": "visitor.consent.rejected_categories", "operator": "Contains", "value": "advertising" } }, { "OR": [/* nested logic nodes */] }] }`.
	Logic DataGovernanceListResponseCategoryLogic `json:"logic" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name           respjson.Field
		Priority       respjson.Field
		Description    respjson.Field
		DestinationIDs respjson.Field
		Logic          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DataGovernanceListResponseCategory) RawJSON() string { return r.JSON.raw }
func (r *DataGovernanceListResponseCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Condition tree evaluated against each inbound event. Write conditions that
// evaluate **TRUE for events you want to STOP**. A node is either a leaf
// `condition` or a combinator (`AND`, `OR`, `NOT`); combinator children are
// themselves logic nodes, so trees nest arbitrarily.
//
// Discovery: `GET /rest/v1/mappings/default-variables` lists the canonical
// platform-provided `property` paths (visitor consent arrays, event fields,
// request context, identity fields). Custom `event.event_properties.*` paths are
// caller-defined.
//
// Example leaf (stop dispatch when the visitor rejected the `advertising` consent
// category):
// `{ "condition": { "property": "visitor.consent.rejected_categories", "operator": "Contains", "value": "advertising" } }`.
//
// Example combinator:
// `{ "AND": [{ "condition": { "property": "visitor.consent.rejected_categories", "operator": "Contains", "value": "advertising" } }, { "OR": [/* nested logic nodes */] }] }`.
type DataGovernanceListResponseCategoryLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And       []map[string]DataGovernanceListResponseCategoryLogicAndUnion `json:"AND" api:"nullable"`
	Condition DataGovernanceListResponseCategoryLogicCondition             `json:"condition" api:"nullable"`
	// Negates a single child logic node.
	Not map[string]DataGovernanceListResponseCategoryLogicNotUnion `json:"NOT" api:"nullable"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or []map[string]DataGovernanceListResponseCategoryLogicOrUnion `json:"OR" api:"nullable"`
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
func (r DataGovernanceListResponseCategoryLogic) RawJSON() string { return r.JSON.raw }
func (r *DataGovernanceListResponseCategoryLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DataGovernanceListResponseCategoryLogicAndUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfDataGovernanceListResponseCategoryLogicAndMapItem]
type DataGovernanceListResponseCategoryLogicAndUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfDataGovernanceListResponseCategoryLogicAndMapItem any `json:",inline"`
	JSON                                                struct {
		OfString                                            respjson.Field
		OfFloat                                             respjson.Field
		OfBool                                              respjson.Field
		OfAnyArray                                          respjson.Field
		OfDataGovernanceListResponseCategoryLogicAndMapItem respjson.Field
		raw                                                 string
	} `json:"-"`
}

func (u DataGovernanceListResponseCategoryLogicAndUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceListResponseCategoryLogicAndUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceListResponseCategoryLogicAndUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceListResponseCategoryLogicAndUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceListResponseCategoryLogicAndUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DataGovernanceListResponseCategoryLogicAndUnion) RawJSON() string { return u.JSON.raw }

func (r *DataGovernanceListResponseCategoryLogicAndUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DataGovernanceListResponseCategoryLogicCondition struct {
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
func (r DataGovernanceListResponseCategoryLogicCondition) RawJSON() string { return r.JSON.raw }
func (r *DataGovernanceListResponseCategoryLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DataGovernanceListResponseCategoryLogicNotUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfDataGovernanceListResponseCategoryLogicNotMapItem]
type DataGovernanceListResponseCategoryLogicNotUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfDataGovernanceListResponseCategoryLogicNotMapItem any `json:",inline"`
	JSON                                                struct {
		OfString                                            respjson.Field
		OfFloat                                             respjson.Field
		OfBool                                              respjson.Field
		OfAnyArray                                          respjson.Field
		OfDataGovernanceListResponseCategoryLogicNotMapItem respjson.Field
		raw                                                 string
	} `json:"-"`
}

func (u DataGovernanceListResponseCategoryLogicNotUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceListResponseCategoryLogicNotUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceListResponseCategoryLogicNotUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceListResponseCategoryLogicNotUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceListResponseCategoryLogicNotUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DataGovernanceListResponseCategoryLogicNotUnion) RawJSON() string { return u.JSON.raw }

func (r *DataGovernanceListResponseCategoryLogicNotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DataGovernanceListResponseCategoryLogicOrUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfDataGovernanceListResponseCategoryLogicOrMapItem]
type DataGovernanceListResponseCategoryLogicOrUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfDataGovernanceListResponseCategoryLogicOrMapItem any `json:",inline"`
	JSON                                               struct {
		OfString                                           respjson.Field
		OfFloat                                            respjson.Field
		OfBool                                             respjson.Field
		OfAnyArray                                         respjson.Field
		OfDataGovernanceListResponseCategoryLogicOrMapItem respjson.Field
		raw                                                string
	} `json:"-"`
}

func (u DataGovernanceListResponseCategoryLogicOrUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceListResponseCategoryLogicOrUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceListResponseCategoryLogicOrUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceListResponseCategoryLogicOrUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceListResponseCategoryLogicOrUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DataGovernanceListResponseCategoryLogicOrUnion) RawJSON() string { return u.JSON.raw }

func (r *DataGovernanceListResponseCategoryLogicOrUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DataGovernanceNewResponse struct {
	// Server-assigned UUID for this data-governance record.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp when the record was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// When false, data governance is configured but does not gate dispatch — inbound
	// events flow through to destinations regardless of category logic.
	IsEnabled bool `json:"isEnabled" api:"required"`
	// Discriminator for this resource. Always "dataGovernance".
	Kind string `json:"kind" api:"required"`
	// Governance categories in priority order (1..N).
	Categories []DataGovernanceNewResponseCategory `json:"categories" api:"nullable"`
	// Human-readable name shown in the dashboard.
	Name string `json:"name" api:"nullable"`
	// Free-form notes for this record.
	Notes string `json:"notes" api:"nullable"`
	// ISO 8601 timestamp of the last write. Equal to createdAt on a freshly created
	// record; advances on every PATCH.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		IsEnabled   respjson.Field
		Kind        respjson.Field
		Categories  respjson.Field
		Name        respjson.Field
		Notes       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DataGovernanceNewResponse) RawJSON() string { return r.JSON.raw }
func (r *DataGovernanceNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DataGovernanceNewResponseCategory struct {
	// Display name for the category.
	Name string `json:"name" api:"required"`
	// 1-indexed sort position. Always equals (sorted index + 1) — see PATCH for
	// details.
	Priority int64 `json:"priority" api:"required"`
	// Optional human-readable description.
	Description string `json:"description" api:"nullable"`
	// Destinations gated by this category when its logic evaluates to TRUE.
	DestinationIDs []string `json:"destinationIds" api:"nullable"`
	// Condition tree evaluated against each inbound event. Write conditions that
	// evaluate **TRUE for events you want to STOP**. A node is either a leaf
	// `condition` or a combinator (`AND`, `OR`, `NOT`); combinator children are
	// themselves logic nodes, so trees nest arbitrarily.
	//
	// Discovery: `GET /rest/v1/mappings/default-variables` lists the canonical
	// platform-provided `property` paths (visitor consent arrays, event fields,
	// request context, identity fields). Custom `event.event_properties.*` paths are
	// caller-defined.
	//
	// Example leaf (stop dispatch when the visitor rejected the `advertising` consent
	// category):
	// `{ "condition": { "property": "visitor.consent.rejected_categories", "operator": "Contains", "value": "advertising" } }`.
	//
	// Example combinator:
	// `{ "AND": [{ "condition": { "property": "visitor.consent.rejected_categories", "operator": "Contains", "value": "advertising" } }, { "OR": [/* nested logic nodes */] }] }`.
	Logic DataGovernanceNewResponseCategoryLogic `json:"logic" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name           respjson.Field
		Priority       respjson.Field
		Description    respjson.Field
		DestinationIDs respjson.Field
		Logic          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DataGovernanceNewResponseCategory) RawJSON() string { return r.JSON.raw }
func (r *DataGovernanceNewResponseCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Condition tree evaluated against each inbound event. Write conditions that
// evaluate **TRUE for events you want to STOP**. A node is either a leaf
// `condition` or a combinator (`AND`, `OR`, `NOT`); combinator children are
// themselves logic nodes, so trees nest arbitrarily.
//
// Discovery: `GET /rest/v1/mappings/default-variables` lists the canonical
// platform-provided `property` paths (visitor consent arrays, event fields,
// request context, identity fields). Custom `event.event_properties.*` paths are
// caller-defined.
//
// Example leaf (stop dispatch when the visitor rejected the `advertising` consent
// category):
// `{ "condition": { "property": "visitor.consent.rejected_categories", "operator": "Contains", "value": "advertising" } }`.
//
// Example combinator:
// `{ "AND": [{ "condition": { "property": "visitor.consent.rejected_categories", "operator": "Contains", "value": "advertising" } }, { "OR": [/* nested logic nodes */] }] }`.
type DataGovernanceNewResponseCategoryLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And       []map[string]DataGovernanceNewResponseCategoryLogicAndUnion `json:"AND" api:"nullable"`
	Condition DataGovernanceNewResponseCategoryLogicCondition             `json:"condition" api:"nullable"`
	// Negates a single child logic node.
	Not map[string]DataGovernanceNewResponseCategoryLogicNotUnion `json:"NOT" api:"nullable"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or []map[string]DataGovernanceNewResponseCategoryLogicOrUnion `json:"OR" api:"nullable"`
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
func (r DataGovernanceNewResponseCategoryLogic) RawJSON() string { return r.JSON.raw }
func (r *DataGovernanceNewResponseCategoryLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DataGovernanceNewResponseCategoryLogicAndUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfDataGovernanceNewResponseCategoryLogicAndMapItem]
type DataGovernanceNewResponseCategoryLogicAndUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfDataGovernanceNewResponseCategoryLogicAndMapItem any `json:",inline"`
	JSON                                               struct {
		OfString                                           respjson.Field
		OfFloat                                            respjson.Field
		OfBool                                             respjson.Field
		OfAnyArray                                         respjson.Field
		OfDataGovernanceNewResponseCategoryLogicAndMapItem respjson.Field
		raw                                                string
	} `json:"-"`
}

func (u DataGovernanceNewResponseCategoryLogicAndUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceNewResponseCategoryLogicAndUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceNewResponseCategoryLogicAndUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceNewResponseCategoryLogicAndUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceNewResponseCategoryLogicAndUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DataGovernanceNewResponseCategoryLogicAndUnion) RawJSON() string { return u.JSON.raw }

func (r *DataGovernanceNewResponseCategoryLogicAndUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DataGovernanceNewResponseCategoryLogicCondition struct {
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
func (r DataGovernanceNewResponseCategoryLogicCondition) RawJSON() string { return r.JSON.raw }
func (r *DataGovernanceNewResponseCategoryLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DataGovernanceNewResponseCategoryLogicNotUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfDataGovernanceNewResponseCategoryLogicNotMapItem]
type DataGovernanceNewResponseCategoryLogicNotUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfDataGovernanceNewResponseCategoryLogicNotMapItem any `json:",inline"`
	JSON                                               struct {
		OfString                                           respjson.Field
		OfFloat                                            respjson.Field
		OfBool                                             respjson.Field
		OfAnyArray                                         respjson.Field
		OfDataGovernanceNewResponseCategoryLogicNotMapItem respjson.Field
		raw                                                string
	} `json:"-"`
}

func (u DataGovernanceNewResponseCategoryLogicNotUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceNewResponseCategoryLogicNotUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceNewResponseCategoryLogicNotUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceNewResponseCategoryLogicNotUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceNewResponseCategoryLogicNotUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DataGovernanceNewResponseCategoryLogicNotUnion) RawJSON() string { return u.JSON.raw }

func (r *DataGovernanceNewResponseCategoryLogicNotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DataGovernanceNewResponseCategoryLogicOrUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfDataGovernanceNewResponseCategoryLogicOrMapItem]
type DataGovernanceNewResponseCategoryLogicOrUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfDataGovernanceNewResponseCategoryLogicOrMapItem any `json:",inline"`
	JSON                                              struct {
		OfString                                          respjson.Field
		OfFloat                                           respjson.Field
		OfBool                                            respjson.Field
		OfAnyArray                                        respjson.Field
		OfDataGovernanceNewResponseCategoryLogicOrMapItem respjson.Field
		raw                                               string
	} `json:"-"`
}

func (u DataGovernanceNewResponseCategoryLogicOrUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceNewResponseCategoryLogicOrUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceNewResponseCategoryLogicOrUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceNewResponseCategoryLogicOrUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceNewResponseCategoryLogicOrUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DataGovernanceNewResponseCategoryLogicOrUnion) RawJSON() string { return u.JSON.raw }

func (r *DataGovernanceNewResponseCategoryLogicOrUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DataGovernanceGetResponse struct {
	// Server-assigned UUID for this data-governance record.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp when the record was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// When false, data governance is configured but does not gate dispatch — inbound
	// events flow through to destinations regardless of category logic.
	IsEnabled bool `json:"isEnabled" api:"required"`
	// Discriminator for this resource. Always "dataGovernance".
	Kind string `json:"kind" api:"required"`
	// Governance categories in priority order (1..N).
	Categories []DataGovernanceGetResponseCategory `json:"categories" api:"nullable"`
	// Human-readable name shown in the dashboard.
	Name string `json:"name" api:"nullable"`
	// Free-form notes for this record.
	Notes string `json:"notes" api:"nullable"`
	// ISO 8601 timestamp of the last write. Equal to createdAt on a freshly created
	// record; advances on every PATCH.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		IsEnabled   respjson.Field
		Kind        respjson.Field
		Categories  respjson.Field
		Name        respjson.Field
		Notes       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DataGovernanceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DataGovernanceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DataGovernanceGetResponseCategory struct {
	// Display name for the category.
	Name string `json:"name" api:"required"`
	// 1-indexed sort position. Always equals (sorted index + 1) — see PATCH for
	// details.
	Priority int64 `json:"priority" api:"required"`
	// Optional human-readable description.
	Description string `json:"description" api:"nullable"`
	// Destinations gated by this category when its logic evaluates to TRUE.
	DestinationIDs []string `json:"destinationIds" api:"nullable"`
	// Condition tree evaluated against each inbound event. Write conditions that
	// evaluate **TRUE for events you want to STOP**. A node is either a leaf
	// `condition` or a combinator (`AND`, `OR`, `NOT`); combinator children are
	// themselves logic nodes, so trees nest arbitrarily.
	//
	// Discovery: `GET /rest/v1/mappings/default-variables` lists the canonical
	// platform-provided `property` paths (visitor consent arrays, event fields,
	// request context, identity fields). Custom `event.event_properties.*` paths are
	// caller-defined.
	//
	// Example leaf (stop dispatch when the visitor rejected the `advertising` consent
	// category):
	// `{ "condition": { "property": "visitor.consent.rejected_categories", "operator": "Contains", "value": "advertising" } }`.
	//
	// Example combinator:
	// `{ "AND": [{ "condition": { "property": "visitor.consent.rejected_categories", "operator": "Contains", "value": "advertising" } }, { "OR": [/* nested logic nodes */] }] }`.
	Logic DataGovernanceGetResponseCategoryLogic `json:"logic" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name           respjson.Field
		Priority       respjson.Field
		Description    respjson.Field
		DestinationIDs respjson.Field
		Logic          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DataGovernanceGetResponseCategory) RawJSON() string { return r.JSON.raw }
func (r *DataGovernanceGetResponseCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Condition tree evaluated against each inbound event. Write conditions that
// evaluate **TRUE for events you want to STOP**. A node is either a leaf
// `condition` or a combinator (`AND`, `OR`, `NOT`); combinator children are
// themselves logic nodes, so trees nest arbitrarily.
//
// Discovery: `GET /rest/v1/mappings/default-variables` lists the canonical
// platform-provided `property` paths (visitor consent arrays, event fields,
// request context, identity fields). Custom `event.event_properties.*` paths are
// caller-defined.
//
// Example leaf (stop dispatch when the visitor rejected the `advertising` consent
// category):
// `{ "condition": { "property": "visitor.consent.rejected_categories", "operator": "Contains", "value": "advertising" } }`.
//
// Example combinator:
// `{ "AND": [{ "condition": { "property": "visitor.consent.rejected_categories", "operator": "Contains", "value": "advertising" } }, { "OR": [/* nested logic nodes */] }] }`.
type DataGovernanceGetResponseCategoryLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And       []map[string]DataGovernanceGetResponseCategoryLogicAndUnion `json:"AND" api:"nullable"`
	Condition DataGovernanceGetResponseCategoryLogicCondition             `json:"condition" api:"nullable"`
	// Negates a single child logic node.
	Not map[string]DataGovernanceGetResponseCategoryLogicNotUnion `json:"NOT" api:"nullable"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or []map[string]DataGovernanceGetResponseCategoryLogicOrUnion `json:"OR" api:"nullable"`
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
func (r DataGovernanceGetResponseCategoryLogic) RawJSON() string { return r.JSON.raw }
func (r *DataGovernanceGetResponseCategoryLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DataGovernanceGetResponseCategoryLogicAndUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfDataGovernanceGetResponseCategoryLogicAndMapItem]
type DataGovernanceGetResponseCategoryLogicAndUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfDataGovernanceGetResponseCategoryLogicAndMapItem any `json:",inline"`
	JSON                                               struct {
		OfString                                           respjson.Field
		OfFloat                                            respjson.Field
		OfBool                                             respjson.Field
		OfAnyArray                                         respjson.Field
		OfDataGovernanceGetResponseCategoryLogicAndMapItem respjson.Field
		raw                                                string
	} `json:"-"`
}

func (u DataGovernanceGetResponseCategoryLogicAndUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceGetResponseCategoryLogicAndUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceGetResponseCategoryLogicAndUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceGetResponseCategoryLogicAndUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceGetResponseCategoryLogicAndUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DataGovernanceGetResponseCategoryLogicAndUnion) RawJSON() string { return u.JSON.raw }

func (r *DataGovernanceGetResponseCategoryLogicAndUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DataGovernanceGetResponseCategoryLogicCondition struct {
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
func (r DataGovernanceGetResponseCategoryLogicCondition) RawJSON() string { return r.JSON.raw }
func (r *DataGovernanceGetResponseCategoryLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DataGovernanceGetResponseCategoryLogicNotUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfDataGovernanceGetResponseCategoryLogicNotMapItem]
type DataGovernanceGetResponseCategoryLogicNotUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfDataGovernanceGetResponseCategoryLogicNotMapItem any `json:",inline"`
	JSON                                               struct {
		OfString                                           respjson.Field
		OfFloat                                            respjson.Field
		OfBool                                             respjson.Field
		OfAnyArray                                         respjson.Field
		OfDataGovernanceGetResponseCategoryLogicNotMapItem respjson.Field
		raw                                                string
	} `json:"-"`
}

func (u DataGovernanceGetResponseCategoryLogicNotUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceGetResponseCategoryLogicNotUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceGetResponseCategoryLogicNotUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceGetResponseCategoryLogicNotUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceGetResponseCategoryLogicNotUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DataGovernanceGetResponseCategoryLogicNotUnion) RawJSON() string { return u.JSON.raw }

func (r *DataGovernanceGetResponseCategoryLogicNotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DataGovernanceGetResponseCategoryLogicOrUnion contains all possible properties
// and values from [string], [float64], [bool], [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfDataGovernanceGetResponseCategoryLogicOrMapItem]
type DataGovernanceGetResponseCategoryLogicOrUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfDataGovernanceGetResponseCategoryLogicOrMapItem any `json:",inline"`
	JSON                                              struct {
		OfString                                          respjson.Field
		OfFloat                                           respjson.Field
		OfBool                                            respjson.Field
		OfAnyArray                                        respjson.Field
		OfDataGovernanceGetResponseCategoryLogicOrMapItem respjson.Field
		raw                                               string
	} `json:"-"`
}

func (u DataGovernanceGetResponseCategoryLogicOrUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceGetResponseCategoryLogicOrUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceGetResponseCategoryLogicOrUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceGetResponseCategoryLogicOrUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceGetResponseCategoryLogicOrUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DataGovernanceGetResponseCategoryLogicOrUnion) RawJSON() string { return u.JSON.raw }

func (r *DataGovernanceGetResponseCategoryLogicOrUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DataGovernanceUpdateResponse struct {
	// Server-assigned UUID for this data-governance record.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp when the record was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// When false, data governance is configured but does not gate dispatch — inbound
	// events flow through to destinations regardless of category logic.
	IsEnabled bool `json:"isEnabled" api:"required"`
	// Discriminator for this resource. Always "dataGovernance".
	Kind string `json:"kind" api:"required"`
	// Governance categories in priority order (1..N).
	Categories []DataGovernanceUpdateResponseCategory `json:"categories" api:"nullable"`
	// Human-readable name shown in the dashboard.
	Name string `json:"name" api:"nullable"`
	// Free-form notes for this record.
	Notes string `json:"notes" api:"nullable"`
	// ISO 8601 timestamp of the last write. Equal to createdAt on a freshly created
	// record; advances on every PATCH.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		IsEnabled   respjson.Field
		Kind        respjson.Field
		Categories  respjson.Field
		Name        respjson.Field
		Notes       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DataGovernanceUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *DataGovernanceUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DataGovernanceUpdateResponseCategory struct {
	// Display name for the category.
	Name string `json:"name" api:"required"`
	// 1-indexed sort position. Always equals (sorted index + 1) — see PATCH for
	// details.
	Priority int64 `json:"priority" api:"required"`
	// Optional human-readable description.
	Description string `json:"description" api:"nullable"`
	// Destinations gated by this category when its logic evaluates to TRUE.
	DestinationIDs []string `json:"destinationIds" api:"nullable"`
	// Condition tree evaluated against each inbound event. Write conditions that
	// evaluate **TRUE for events you want to STOP**. A node is either a leaf
	// `condition` or a combinator (`AND`, `OR`, `NOT`); combinator children are
	// themselves logic nodes, so trees nest arbitrarily.
	//
	// Discovery: `GET /rest/v1/mappings/default-variables` lists the canonical
	// platform-provided `property` paths (visitor consent arrays, event fields,
	// request context, identity fields). Custom `event.event_properties.*` paths are
	// caller-defined.
	//
	// Example leaf (stop dispatch when the visitor rejected the `advertising` consent
	// category):
	// `{ "condition": { "property": "visitor.consent.rejected_categories", "operator": "Contains", "value": "advertising" } }`.
	//
	// Example combinator:
	// `{ "AND": [{ "condition": { "property": "visitor.consent.rejected_categories", "operator": "Contains", "value": "advertising" } }, { "OR": [/* nested logic nodes */] }] }`.
	Logic DataGovernanceUpdateResponseCategoryLogic `json:"logic" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name           respjson.Field
		Priority       respjson.Field
		Description    respjson.Field
		DestinationIDs respjson.Field
		Logic          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DataGovernanceUpdateResponseCategory) RawJSON() string { return r.JSON.raw }
func (r *DataGovernanceUpdateResponseCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Condition tree evaluated against each inbound event. Write conditions that
// evaluate **TRUE for events you want to STOP**. A node is either a leaf
// `condition` or a combinator (`AND`, `OR`, `NOT`); combinator children are
// themselves logic nodes, so trees nest arbitrarily.
//
// Discovery: `GET /rest/v1/mappings/default-variables` lists the canonical
// platform-provided `property` paths (visitor consent arrays, event fields,
// request context, identity fields). Custom `event.event_properties.*` paths are
// caller-defined.
//
// Example leaf (stop dispatch when the visitor rejected the `advertising` consent
// category):
// `{ "condition": { "property": "visitor.consent.rejected_categories", "operator": "Contains", "value": "advertising" } }`.
//
// Example combinator:
// `{ "AND": [{ "condition": { "property": "visitor.consent.rejected_categories", "operator": "Contains", "value": "advertising" } }, { "OR": [/* nested logic nodes */] }] }`.
type DataGovernanceUpdateResponseCategoryLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And       []map[string]DataGovernanceUpdateResponseCategoryLogicAndUnion `json:"AND" api:"nullable"`
	Condition DataGovernanceUpdateResponseCategoryLogicCondition             `json:"condition" api:"nullable"`
	// Negates a single child logic node.
	Not map[string]DataGovernanceUpdateResponseCategoryLogicNotUnion `json:"NOT" api:"nullable"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or []map[string]DataGovernanceUpdateResponseCategoryLogicOrUnion `json:"OR" api:"nullable"`
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
func (r DataGovernanceUpdateResponseCategoryLogic) RawJSON() string { return r.JSON.raw }
func (r *DataGovernanceUpdateResponseCategoryLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DataGovernanceUpdateResponseCategoryLogicAndUnion contains all possible
// properties and values from [string], [float64], [bool], [[]any],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfDataGovernanceUpdateResponseCategoryLogicAndMapItem]
type DataGovernanceUpdateResponseCategoryLogicAndUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfDataGovernanceUpdateResponseCategoryLogicAndMapItem any `json:",inline"`
	JSON                                                  struct {
		OfString                                              respjson.Field
		OfFloat                                               respjson.Field
		OfBool                                                respjson.Field
		OfAnyArray                                            respjson.Field
		OfDataGovernanceUpdateResponseCategoryLogicAndMapItem respjson.Field
		raw                                                   string
	} `json:"-"`
}

func (u DataGovernanceUpdateResponseCategoryLogicAndUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceUpdateResponseCategoryLogicAndUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceUpdateResponseCategoryLogicAndUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceUpdateResponseCategoryLogicAndUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceUpdateResponseCategoryLogicAndUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DataGovernanceUpdateResponseCategoryLogicAndUnion) RawJSON() string { return u.JSON.raw }

func (r *DataGovernanceUpdateResponseCategoryLogicAndUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DataGovernanceUpdateResponseCategoryLogicCondition struct {
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
func (r DataGovernanceUpdateResponseCategoryLogicCondition) RawJSON() string { return r.JSON.raw }
func (r *DataGovernanceUpdateResponseCategoryLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DataGovernanceUpdateResponseCategoryLogicNotUnion contains all possible
// properties and values from [string], [float64], [bool], [[]any],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfDataGovernanceUpdateResponseCategoryLogicNotMapItem]
type DataGovernanceUpdateResponseCategoryLogicNotUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfDataGovernanceUpdateResponseCategoryLogicNotMapItem any `json:",inline"`
	JSON                                                  struct {
		OfString                                              respjson.Field
		OfFloat                                               respjson.Field
		OfBool                                                respjson.Field
		OfAnyArray                                            respjson.Field
		OfDataGovernanceUpdateResponseCategoryLogicNotMapItem respjson.Field
		raw                                                   string
	} `json:"-"`
}

func (u DataGovernanceUpdateResponseCategoryLogicNotUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceUpdateResponseCategoryLogicNotUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceUpdateResponseCategoryLogicNotUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceUpdateResponseCategoryLogicNotUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceUpdateResponseCategoryLogicNotUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DataGovernanceUpdateResponseCategoryLogicNotUnion) RawJSON() string { return u.JSON.raw }

func (r *DataGovernanceUpdateResponseCategoryLogicNotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DataGovernanceUpdateResponseCategoryLogicOrUnion contains all possible
// properties and values from [string], [float64], [bool], [[]any],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfDataGovernanceUpdateResponseCategoryLogicOrMapItem]
type DataGovernanceUpdateResponseCategoryLogicOrUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfDataGovernanceUpdateResponseCategoryLogicOrMapItem any `json:",inline"`
	JSON                                                 struct {
		OfString                                             respjson.Field
		OfFloat                                              respjson.Field
		OfBool                                               respjson.Field
		OfAnyArray                                           respjson.Field
		OfDataGovernanceUpdateResponseCategoryLogicOrMapItem respjson.Field
		raw                                                  string
	} `json:"-"`
}

func (u DataGovernanceUpdateResponseCategoryLogicOrUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceUpdateResponseCategoryLogicOrUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceUpdateResponseCategoryLogicOrUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceUpdateResponseCategoryLogicOrUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataGovernanceUpdateResponseCategoryLogicOrUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DataGovernanceUpdateResponseCategoryLogicOrUnion) RawJSON() string { return u.JSON.raw }

func (r *DataGovernanceUpdateResponseCategoryLogicOrUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DataGovernanceDeleteResponse struct {
	// The id of the data-governance record that was deleted.
	ID string `json:"id" api:"required"`
	// True when the record was deleted.
	Deleted bool `json:"deleted" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Deleted     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DataGovernanceDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *DataGovernanceDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DataGovernanceListParams struct {
	// Maximum number of items to return. Defaults to 25; values below 1 are clamped to
	// 1 and values above 100 are clamped to 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Opaque pagination cursor from pagination.nextCursor in the previous response. Do
	// not decode or modify it. Malformed cursors return 400 Bad Request.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DataGovernanceListParams]'s query parameters as
// `url.Values`.
func (r DataGovernanceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DataGovernanceNewParams struct {
	// Whether the record starts enabled. Defaults to false — opt in by setting true
	// here or via PATCH later. When disabled, every category is bypassed and inbound
	// events flow through to destinations regardless of consent state.
	IsEnabled param.Opt[bool] `json:"isEnabled,omitzero"`
	// Display name for the new record. Defaults to "Data Governance".
	Name param.Opt[string] `json:"name,omitzero"`
	// Free-form notes shown in the dashboard. Not used for routing.
	Notes param.Opt[string] `json:"notes,omitzero"`
	paramObj
}

func (r DataGovernanceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DataGovernanceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DataGovernanceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DataGovernanceUpdateParams struct {
	// Toggle data governance on/off without changing any other config.
	IsEnabled param.Opt[bool] `json:"isEnabled,omitzero"`
	// New display name for the record.
	Name param.Opt[string] `json:"name,omitzero"`
	// Replace the free-form notes.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Full replacement of the categories list. The sent array becomes the new state —
	// there is no partial-merge semantic for categories. Categories are sorted by
	// priority on write and re-stamped 1..N. Omit to leave existing categories
	// untouched.
	Categories []DataGovernanceUpdateParamsCategory `json:"categories,omitzero"`
	paramObj
}

func (r DataGovernanceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DataGovernanceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DataGovernanceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DataGovernanceUpdateParamsCategory struct {
	// Optional human-readable description shown in the dashboard.
	Description param.Opt[string] `json:"description,omitzero"`
	// Display name for the category. Auto-generated if omitted.
	Name param.Opt[string] `json:"name,omitzero"`
	// Used as a sort key on write. The server sorts categories by this value
	// (ascending), then re-stamps priority as (sorted index + 1) on persist. Send any
	// positive number — gaps are ironed out, duplicate values keep input order via
	// stable sort. Omit to fall to the end.
	Priority param.Opt[float64] `json:"priority,omitzero"`
	// Destinations gated by this category. When the category logic evaluates to TRUE
	// for an inbound event, dispatch to every destination on this list is stopped.
	// Stale IDs (deleted destinations or destinations on another account) are silently
	// filtered out at write time.
	DestinationIDs []string `json:"destinationIds,omitzero"`
	// Condition tree evaluated against each inbound event. Write conditions that
	// evaluate **TRUE for events you want to STOP**. A node is either a leaf
	// `condition` or a combinator (`AND`, `OR`, `NOT`); combinator children are
	// themselves logic nodes, so trees nest arbitrarily.
	//
	// Discovery: `GET /rest/v1/mappings/default-variables` lists the canonical
	// platform-provided `property` paths (visitor consent arrays, event fields,
	// request context, identity fields). Custom `event.event_properties.*` paths are
	// caller-defined.
	//
	// Example leaf (stop dispatch when the visitor rejected the `advertising` consent
	// category):
	// `{ "condition": { "property": "visitor.consent.rejected_categories", "operator": "Contains", "value": "advertising" } }`.
	//
	// Example combinator:
	// `{ "AND": [{ "condition": { "property": "visitor.consent.rejected_categories", "operator": "Contains", "value": "advertising" } }, { "OR": [/* nested logic nodes */] }] }`.
	Logic DataGovernanceUpdateParamsCategoryLogic `json:"logic,omitzero"`
	paramObj
}

func (r DataGovernanceUpdateParamsCategory) MarshalJSON() (data []byte, err error) {
	type shadow DataGovernanceUpdateParamsCategory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DataGovernanceUpdateParamsCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Condition tree evaluated against each inbound event. Write conditions that
// evaluate **TRUE for events you want to STOP**. A node is either a leaf
// `condition` or a combinator (`AND`, `OR`, `NOT`); combinator children are
// themselves logic nodes, so trees nest arbitrarily.
//
// Discovery: `GET /rest/v1/mappings/default-variables` lists the canonical
// platform-provided `property` paths (visitor consent arrays, event fields,
// request context, identity fields). Custom `event.event_properties.*` paths are
// caller-defined.
//
// Example leaf (stop dispatch when the visitor rejected the `advertising` consent
// category):
// `{ "condition": { "property": "visitor.consent.rejected_categories", "operator": "Contains", "value": "advertising" } }`.
//
// Example combinator:
// `{ "AND": [{ "condition": { "property": "visitor.consent.rejected_categories", "operator": "Contains", "value": "advertising" } }, { "OR": [/* nested logic nodes */] }] }`.
type DataGovernanceUpdateParamsCategoryLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And       []map[string]DataGovernanceUpdateParamsCategoryLogicAndUnion `json:"AND,omitzero"`
	Condition DataGovernanceUpdateParamsCategoryLogicCondition             `json:"condition,omitzero"`
	// Negates a single child logic node.
	Not map[string]DataGovernanceUpdateParamsCategoryLogicNotUnion `json:"NOT,omitzero"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or []map[string]DataGovernanceUpdateParamsCategoryLogicOrUnion `json:"OR,omitzero"`
	paramObj
}

func (r DataGovernanceUpdateParamsCategoryLogic) MarshalJSON() (data []byte, err error) {
	type shadow DataGovernanceUpdateParamsCategoryLogic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DataGovernanceUpdateParamsCategoryLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DataGovernanceUpdateParamsCategoryLogicAndUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u DataGovernanceUpdateParamsCategoryLogicAndUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *DataGovernanceUpdateParamsCategoryLogicAndUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Operator, Property, Value are required.
type DataGovernanceUpdateParamsCategoryLogicCondition struct {
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

func (r DataGovernanceUpdateParamsCategoryLogicCondition) MarshalJSON() (data []byte, err error) {
	type shadow DataGovernanceUpdateParamsCategoryLogicCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DataGovernanceUpdateParamsCategoryLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DataGovernanceUpdateParamsCategoryLogicCondition](
		"operator", "Is", "IsNot", "Contains", "DoesNotContain", "StartsWith", "EndsWith", "IsFalsy", "IsTruthy", "IsNull", "IsNotNull", "IsUndefined", "IsNotUndefined", "IsGreaterThan", "IsGreaterThanOrEqual", "IsLessThan", "IsLessThanOrEqual", "IsIn", "IsNotIn", "IsFoundIn", "IsNotFoundIn", "IsTrue", "IsFalse", "IsBefore", "IsAfter", "IsBetween", "IsOnOrBefore", "IsOnOrAfter", "MatchesRegex", "MatchesRegexIgnoreCase", "DoesNotMatchRegex", "DoesNotMatchRegexIgnoreCase",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DataGovernanceUpdateParamsCategoryLogicNotUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u DataGovernanceUpdateParamsCategoryLogicNotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *DataGovernanceUpdateParamsCategoryLogicNotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DataGovernanceUpdateParamsCategoryLogicOrUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u DataGovernanceUpdateParamsCategoryLogicOrUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfAnyArray,
		u.OfAnyMap)
}
func (u *DataGovernanceUpdateParamsCategoryLogicOrUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}
