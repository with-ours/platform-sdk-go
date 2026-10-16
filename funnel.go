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
	"time"

	"github.com/with-ours/platform-sdk-go/v2/internal/apijson"
	"github.com/with-ours/platform-sdk-go/v2/internal/apiquery"
	"github.com/with-ours/platform-sdk-go/v2/internal/requestconfig"
	"github.com/with-ours/platform-sdk-go/v2/option"
	"github.com/with-ours/platform-sdk-go/v2/packages/param"
	"github.com/with-ours/platform-sdk-go/v2/packages/respjson"
)

// FunnelService contains methods and other services that help with interacting
// with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFunnelService] method instead.
type FunnelService struct {
	Options []option.RequestOption
}

// NewFunnelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFunnelService(opts ...option.RequestOption) (r FunnelService) {
	r = FunnelService{}
	r.Options = opts
	return
}

// List readable funnels configured on this account. Each funnel includes its
// canonical versioned definition. If some funnels cannot be loaded,
// `unavailableCount` and `warnings` identify the incomplete result; omitted
// funnels must not be treated as deleted. A complete list has
// `unavailableCount: 0` and no warnings. Funnel results are computed on demand, so
// `status` is always `READY` and `reportDateRange` is always `null`. Requires
// API-key scope or current OAuth user permission: web-analytics:view
func (r *FunnelService) List(ctx context.Context, opts ...option.RequestOption) (res *FunnelListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/funnels"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Create a funnel from a versioned definition. Returns the complete saved
// configuration. Requires API-key scope or current OAuth user permission:
// web-analytics:write
func (r *FunnelService) New(ctx context.Context, body FunnelNewParams, opts ...option.RequestOption) (res *FunnelNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/funnels"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Fetch a single funnel configuration by its id. Returns `404` when the funnel
// does not exist or belongs to a different account. Requires API-key scope or
// current OAuth user permission: web-analytics:view
func (r *FunnelService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *FunnelGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/funnels/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update one or more funnel fields. Omitted fields remain unchanged. The canonical
// definition can be replaced by supplying `queryDefinition`. Requires API-key
// scope or current OAuth user permission: web-analytics:write
func (r *FunnelService) Update(ctx context.Context, id string, body FunnelUpdateParams, opts ...option.RequestOption) (res *FunnelUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/funnels/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete a Funnel configuration. Existing analytics data is unaffected. Requires
// API-key scope or current OAuth user permission: web-analytics:write
func (r *FunnelService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *FunnelDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/funnels/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Duplicate a funnel configuration in the same account. The copy keeps the
// canonical definition, receives a new ID, and is named `Copy of …`. Requires
// API-key scope or current OAuth user permission: web-analytics:write
func (r *FunnelService) Duplicate(ctx context.Context, id string, opts ...option.RequestOption) (res *FunnelDuplicateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/funnels/%s/duplicate", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Compute funnel step analytics from the funnel’s saved `queryDefinition` over a
// requested date window. Returns per-step visitor counts, conversion rates,
// drop-off rates, average time to next step, sample session IDs for replay, and
// the resolved entry/observation scope. Results are computed on demand from event
// data at request time; ad hoc filter, saved-scope, and web-source overrides are
// not accepted. Observation completeness is null when the source coverage
// watermark is unavailable. `to` must be on or after `from`, and the window may
// span at most 91 days including both endpoints. Requires API-key scope or current
// OAuth user permission: web-analytics:view
func (r *FunnelService) Results(ctx context.Context, id string, query FunnelResultsParams, opts ...option.RequestOption) (res *FunnelResultsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/funnels/%s/results", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type FunnelListResponse struct {
	// The readable funnels configured on this account.
	Entities []FunnelListResponseEntity `json:"entities" api:"required"`
	// Number of funnels that could not be loaded. A positive count means the list is
	// incomplete.
	UnavailableCount int64 `json:"unavailableCount" api:"required"`
	// Warnings about incomplete results. Empty when every funnel was loaded.
	Warnings []string `json:"warnings" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities         respjson.Field
		UnavailableCount respjson.Field
		Warnings         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelListResponse) RawJSON() string { return r.JSON.raw }
func (r *FunnelListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelListResponseEntity struct {
	CreatedAt string `json:"createdAt" api:"required"`
	FunnelID  string `json:"funnelId" api:"required" format:"uuid"`
	Name      string `json:"name" api:"required"`
	// Versioned funnel definition with typed steps and filters.
	QueryDefinition FunnelListResponseEntityQueryDefinition `json:"queryDefinition" api:"required"`
	// Any of "READY", "PROCESSING".
	Status          string                                  `json:"status" api:"required"`
	UpdatedAt       string                                  `json:"updatedAt" api:"required"`
	Description     string                                  `json:"description" api:"nullable"`
	ReportDateRange FunnelListResponseEntityReportDateRange `json:"reportDateRange" api:"nullable"`
	Watched         bool                                    `json:"watched" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt       respjson.Field
		FunnelID        respjson.Field
		Name            respjson.Field
		QueryDefinition respjson.Field
		Status          respjson.Field
		UpdatedAt       respjson.Field
		Description     respjson.Field
		ReportDateRange respjson.Field
		Watched         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelListResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *FunnelListResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Versioned funnel definition with typed steps and filters.
type FunnelListResponseEntityQueryDefinition struct {
	// Any of "unique", "events", "sessions".
	Counting string `json:"counting" api:"required"`
	// Any of "session", "visitor".
	Identity string `json:"identity" api:"required"`
	// Any of "funnel".
	Kind string `json:"kind" api:"required"`
	// Any of "ordered", "any".
	Order string `json:"order" api:"required"`
	// Any of "first", "best".
	Reentry string                                        `json:"reentry" api:"required"`
	Steps   []FunnelListResponseEntityQueryDefinitionStep `json:"steps" api:"required"`
	// Any of 1.
	Version     float64                                            `json:"version" api:"required"`
	WindowMs    int64                                              `json:"windowMs" api:"required"`
	EntryFilter FunnelListResponseEntityQueryDefinitionEntryFilter `json:"entryFilter"`
	Filter      FunnelListResponseEntityQueryDefinitionFilter      `json:"filter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Counting    respjson.Field
		Identity    respjson.Field
		Kind        respjson.Field
		Order       respjson.Field
		Reentry     respjson.Field
		Steps       respjson.Field
		Version     respjson.Field
		WindowMs    respjson.Field
		EntryFilter respjson.Field
		Filter      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelListResponseEntityQueryDefinition) RawJSON() string { return r.JSON.raw }
func (r *FunnelListResponseEntityQueryDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelListResponseEntityQueryDefinitionStep struct {
	ID        string                                            `json:"id" api:"required" format:"uuid"`
	EventName string                                            `json:"eventName" api:"required"`
	Name      string                                            `json:"name" api:"required"`
	Filter    FunnelListResponseEntityQueryDefinitionStepFilter `json:"filter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventName   respjson.Field
		Name        respjson.Field
		Filter      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelListResponseEntityQueryDefinitionStep) RawJSON() string { return r.JSON.raw }
func (r *FunnelListResponseEntityQueryDefinitionStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelListResponseEntityQueryDefinitionStepFilter struct {
	Filter FunnelListResponseEntityQueryDefinitionStepFilterFilterUnion `json:"filter" api:"required"`
	// Any of 1.
	Version float64 `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filter      respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelListResponseEntityQueryDefinitionStepFilter) RawJSON() string { return r.JSON.raw }
func (r *FunnelListResponseEntityQueryDefinitionStepFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelListResponseEntityQueryDefinitionStepFilterFilterUnion contains all
// possible properties and values from
// [FunnelListResponseEntityQueryDefinitionStepFilterFilterObject],
// [FunnelListResponseEntityQueryDefinitionStepFilterFilterObject2],
// [FunnelListResponseEntityQueryDefinitionStepFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FunnelListResponseEntityQueryDefinitionStepFilterFilterUnion struct {
	// This field is from variant
	// [FunnelListResponseEntityQueryDefinitionStepFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant
	// [FunnelListResponseEntityQueryDefinitionStepFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant
	// [FunnelListResponseEntityQueryDefinitionStepFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant
	// [FunnelListResponseEntityQueryDefinitionStepFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant
	// [FunnelListResponseEntityQueryDefinitionStepFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant
	// [FunnelListResponseEntityQueryDefinitionStepFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant
	// [FunnelListResponseEntityQueryDefinitionStepFilterFilterObject3].
	Value FunnelListResponseEntityQueryDefinitionStepFilterFilterObject3ValueUnion `json:"value"`
	JSON  struct {
		Children respjson.Field
		Kind     respjson.Field
		Child    respjson.Field
		Operator respjson.Field
		Property respjson.Field
		Path     respjson.Field
		Type     respjson.Field
		Value    respjson.Field
		raw      string
	} `json:"-"`
}

func (u FunnelListResponseEntityQueryDefinitionStepFilterFilterUnion) AsFunnelListResponseEntityQueryDefinitionStepFilterFilterObject() (v FunnelListResponseEntityQueryDefinitionStepFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelListResponseEntityQueryDefinitionStepFilterFilterUnion) AsFunnelListResponseEntityQueryDefinitionStepFilterFilterObject2() (v FunnelListResponseEntityQueryDefinitionStepFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelListResponseEntityQueryDefinitionStepFilterFilterUnion) AsFunnelListResponseEntityQueryDefinitionStepFilterFilterObject3() (v FunnelListResponseEntityQueryDefinitionStepFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelListResponseEntityQueryDefinitionStepFilterFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelListResponseEntityQueryDefinitionStepFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelListResponseEntityQueryDefinitionStepFilterFilterObject struct {
	Children []any `json:"children" api:"required"`
	// Any of "and", "or".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Children    respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelListResponseEntityQueryDefinitionStepFilterFilterObject) RawJSON() string {
	return r.JSON.raw
}
func (r *FunnelListResponseEntityQueryDefinitionStepFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelListResponseEntityQueryDefinitionStepFilterFilterObject2 struct {
	Child any `json:"child" api:"required"`
	// Any of "not".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Child       respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelListResponseEntityQueryDefinitionStepFilterFilterObject2) RawJSON() string {
	return r.JSON.raw
}
func (r *FunnelListResponseEntityQueryDefinitionStepFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelListResponseEntityQueryDefinitionStepFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path"`
	// Any of "string", "number", "boolean".
	Type  string                                                                   `json:"type"`
	Value FunnelListResponseEntityQueryDefinitionStepFilterFilterObject3ValueUnion `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		Operator    respjson.Field
		Property    respjson.Field
		Path        respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelListResponseEntityQueryDefinitionStepFilterFilterObject3) RawJSON() string {
	return r.JSON.raw
}
func (r *FunnelListResponseEntityQueryDefinitionStepFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelListResponseEntityQueryDefinitionStepFilterFilterObject3ValueUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]FunnelListResponseEntityQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfFunnelListResponseEntityQueryDefinitionStepFilterFilterObject3ValueArray]
type FunnelListResponseEntityQueryDefinitionStepFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]FunnelListResponseEntityQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion]
	// instead of an object.
	OfFunnelListResponseEntityQueryDefinitionStepFilterFilterObject3ValueArray []FunnelListResponseEntityQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                                       struct {
		OfString                                                                   respjson.Field
		OfFloat                                                                    respjson.Field
		OfBool                                                                     respjson.Field
		OfFunnelListResponseEntityQueryDefinitionStepFilterFilterObject3ValueArray respjson.Field
		raw                                                                        string
	} `json:"-"`
}

func (u FunnelListResponseEntityQueryDefinitionStepFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelListResponseEntityQueryDefinitionStepFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelListResponseEntityQueryDefinitionStepFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelListResponseEntityQueryDefinitionStepFilterFilterObject3ValueUnion) AsFunnelListResponseEntityQueryDefinitionStepFilterFilterObject3ValueArray() (v []FunnelListResponseEntityQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelListResponseEntityQueryDefinitionStepFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelListResponseEntityQueryDefinitionStepFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelListResponseEntityQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type FunnelListResponseEntityQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u FunnelListResponseEntityQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelListResponseEntityQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelListResponseEntityQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelListResponseEntityQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelListResponseEntityQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelListResponseEntityQueryDefinitionEntryFilter struct {
	Filter FunnelListResponseEntityQueryDefinitionEntryFilterFilterUnion `json:"filter" api:"required"`
	// Any of 1.
	Version float64 `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filter      respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelListResponseEntityQueryDefinitionEntryFilter) RawJSON() string { return r.JSON.raw }
func (r *FunnelListResponseEntityQueryDefinitionEntryFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelListResponseEntityQueryDefinitionEntryFilterFilterUnion contains all
// possible properties and values from
// [FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject],
// [FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject2],
// [FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FunnelListResponseEntityQueryDefinitionEntryFilterFilterUnion struct {
	// This field is from variant
	// [FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant
	// [FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant
	// [FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant
	// [FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant
	// [FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant
	// [FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant
	// [FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3].
	Value FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3ValueUnion `json:"value"`
	JSON  struct {
		Children respjson.Field
		Kind     respjson.Field
		Child    respjson.Field
		Operator respjson.Field
		Property respjson.Field
		Path     respjson.Field
		Type     respjson.Field
		Value    respjson.Field
		raw      string
	} `json:"-"`
}

func (u FunnelListResponseEntityQueryDefinitionEntryFilterFilterUnion) AsFunnelListResponseEntityQueryDefinitionEntryFilterFilterObject() (v FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelListResponseEntityQueryDefinitionEntryFilterFilterUnion) AsFunnelListResponseEntityQueryDefinitionEntryFilterFilterObject2() (v FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelListResponseEntityQueryDefinitionEntryFilterFilterUnion) AsFunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3() (v FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelListResponseEntityQueryDefinitionEntryFilterFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelListResponseEntityQueryDefinitionEntryFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject struct {
	Children []any `json:"children" api:"required"`
	// Any of "and", "or".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Children    respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject) RawJSON() string {
	return r.JSON.raw
}
func (r *FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject2 struct {
	Child any `json:"child" api:"required"`
	// Any of "not".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Child       respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject2) RawJSON() string {
	return r.JSON.raw
}
func (r *FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path"`
	// Any of "string", "number", "boolean".
	Type  string                                                                    `json:"type"`
	Value FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3ValueUnion `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		Operator    respjson.Field
		Property    respjson.Field
		Path        respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3) RawJSON() string {
	return r.JSON.raw
}
func (r *FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3ValueUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfFunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3ValueArray]
type FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion]
	// instead of an object.
	OfFunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3ValueArray []FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                                        struct {
		OfString                                                                    respjson.Field
		OfFloat                                                                     respjson.Field
		OfBool                                                                      respjson.Field
		OfFunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3ValueArray respjson.Field
		raw                                                                         string
	} `json:"-"`
}

func (u FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3ValueUnion) AsFunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3ValueArray() (v []FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelListResponseEntityQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelListResponseEntityQueryDefinitionFilter struct {
	Filter FunnelListResponseEntityQueryDefinitionFilterFilterUnion `json:"filter" api:"required"`
	// Any of 1.
	Version float64 `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filter      respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelListResponseEntityQueryDefinitionFilter) RawJSON() string { return r.JSON.raw }
func (r *FunnelListResponseEntityQueryDefinitionFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelListResponseEntityQueryDefinitionFilterFilterUnion contains all possible
// properties and values from
// [FunnelListResponseEntityQueryDefinitionFilterFilterObject],
// [FunnelListResponseEntityQueryDefinitionFilterFilterObject2],
// [FunnelListResponseEntityQueryDefinitionFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FunnelListResponseEntityQueryDefinitionFilterFilterUnion struct {
	// This field is from variant
	// [FunnelListResponseEntityQueryDefinitionFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant
	// [FunnelListResponseEntityQueryDefinitionFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant
	// [FunnelListResponseEntityQueryDefinitionFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant
	// [FunnelListResponseEntityQueryDefinitionFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant
	// [FunnelListResponseEntityQueryDefinitionFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant
	// [FunnelListResponseEntityQueryDefinitionFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant
	// [FunnelListResponseEntityQueryDefinitionFilterFilterObject3].
	Value FunnelListResponseEntityQueryDefinitionFilterFilterObject3ValueUnion `json:"value"`
	JSON  struct {
		Children respjson.Field
		Kind     respjson.Field
		Child    respjson.Field
		Operator respjson.Field
		Property respjson.Field
		Path     respjson.Field
		Type     respjson.Field
		Value    respjson.Field
		raw      string
	} `json:"-"`
}

func (u FunnelListResponseEntityQueryDefinitionFilterFilterUnion) AsFunnelListResponseEntityQueryDefinitionFilterFilterObject() (v FunnelListResponseEntityQueryDefinitionFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelListResponseEntityQueryDefinitionFilterFilterUnion) AsFunnelListResponseEntityQueryDefinitionFilterFilterObject2() (v FunnelListResponseEntityQueryDefinitionFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelListResponseEntityQueryDefinitionFilterFilterUnion) AsFunnelListResponseEntityQueryDefinitionFilterFilterObject3() (v FunnelListResponseEntityQueryDefinitionFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelListResponseEntityQueryDefinitionFilterFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *FunnelListResponseEntityQueryDefinitionFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelListResponseEntityQueryDefinitionFilterFilterObject struct {
	Children []any `json:"children" api:"required"`
	// Any of "and", "or".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Children    respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelListResponseEntityQueryDefinitionFilterFilterObject) RawJSON() string {
	return r.JSON.raw
}
func (r *FunnelListResponseEntityQueryDefinitionFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelListResponseEntityQueryDefinitionFilterFilterObject2 struct {
	Child any `json:"child" api:"required"`
	// Any of "not".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Child       respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelListResponseEntityQueryDefinitionFilterFilterObject2) RawJSON() string {
	return r.JSON.raw
}
func (r *FunnelListResponseEntityQueryDefinitionFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelListResponseEntityQueryDefinitionFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path"`
	// Any of "string", "number", "boolean".
	Type  string                                                               `json:"type"`
	Value FunnelListResponseEntityQueryDefinitionFilterFilterObject3ValueUnion `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		Operator    respjson.Field
		Property    respjson.Field
		Path        respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelListResponseEntityQueryDefinitionFilterFilterObject3) RawJSON() string {
	return r.JSON.raw
}
func (r *FunnelListResponseEntityQueryDefinitionFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelListResponseEntityQueryDefinitionFilterFilterObject3ValueUnion contains
// all possible properties and values from [string], [float64], [bool],
// [[]FunnelListResponseEntityQueryDefinitionFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfFunnelListResponseEntityQueryDefinitionFilterFilterObject3ValueArray]
type FunnelListResponseEntityQueryDefinitionFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]FunnelListResponseEntityQueryDefinitionFilterFilterObject3ValueArrayItemUnion]
	// instead of an object.
	OfFunnelListResponseEntityQueryDefinitionFilterFilterObject3ValueArray []FunnelListResponseEntityQueryDefinitionFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                                   struct {
		OfString                                                               respjson.Field
		OfFloat                                                                respjson.Field
		OfBool                                                                 respjson.Field
		OfFunnelListResponseEntityQueryDefinitionFilterFilterObject3ValueArray respjson.Field
		raw                                                                    string
	} `json:"-"`
}

func (u FunnelListResponseEntityQueryDefinitionFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelListResponseEntityQueryDefinitionFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelListResponseEntityQueryDefinitionFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelListResponseEntityQueryDefinitionFilterFilterObject3ValueUnion) AsFunnelListResponseEntityQueryDefinitionFilterFilterObject3ValueArray() (v []FunnelListResponseEntityQueryDefinitionFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelListResponseEntityQueryDefinitionFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelListResponseEntityQueryDefinitionFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelListResponseEntityQueryDefinitionFilterFilterObject3ValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type FunnelListResponseEntityQueryDefinitionFilterFilterObject3ValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u FunnelListResponseEntityQueryDefinitionFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelListResponseEntityQueryDefinitionFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelListResponseEntityQueryDefinitionFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelListResponseEntityQueryDefinitionFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelListResponseEntityQueryDefinitionFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelListResponseEntityReportDateRange struct {
	From string `json:"from" api:"required"`
	To   string `json:"to" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelListResponseEntityReportDateRange) RawJSON() string { return r.JSON.raw }
func (r *FunnelListResponseEntityReportDateRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelNewResponse struct {
	CreatedAt string `json:"createdAt" api:"required"`
	FunnelID  string `json:"funnelId" api:"required" format:"uuid"`
	Name      string `json:"name" api:"required"`
	// Versioned funnel definition with typed steps and filters.
	QueryDefinition FunnelNewResponseQueryDefinition `json:"queryDefinition" api:"required"`
	// Any of "READY", "PROCESSING".
	Status          FunnelNewResponseStatus          `json:"status" api:"required"`
	UpdatedAt       string                           `json:"updatedAt" api:"required"`
	Description     string                           `json:"description" api:"nullable"`
	ReportDateRange FunnelNewResponseReportDateRange `json:"reportDateRange" api:"nullable"`
	Watched         bool                             `json:"watched" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt       respjson.Field
		FunnelID        respjson.Field
		Name            respjson.Field
		QueryDefinition respjson.Field
		Status          respjson.Field
		UpdatedAt       respjson.Field
		Description     respjson.Field
		ReportDateRange respjson.Field
		Watched         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelNewResponse) RawJSON() string { return r.JSON.raw }
func (r *FunnelNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Versioned funnel definition with typed steps and filters.
type FunnelNewResponseQueryDefinition struct {
	// Any of "unique", "events", "sessions".
	Counting string `json:"counting" api:"required"`
	// Any of "session", "visitor".
	Identity string `json:"identity" api:"required"`
	// Any of "funnel".
	Kind string `json:"kind" api:"required"`
	// Any of "ordered", "any".
	Order string `json:"order" api:"required"`
	// Any of "first", "best".
	Reentry string                                 `json:"reentry" api:"required"`
	Steps   []FunnelNewResponseQueryDefinitionStep `json:"steps" api:"required"`
	// Any of 1.
	Version     float64                                     `json:"version" api:"required"`
	WindowMs    int64                                       `json:"windowMs" api:"required"`
	EntryFilter FunnelNewResponseQueryDefinitionEntryFilter `json:"entryFilter"`
	Filter      FunnelNewResponseQueryDefinitionFilter      `json:"filter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Counting    respjson.Field
		Identity    respjson.Field
		Kind        respjson.Field
		Order       respjson.Field
		Reentry     respjson.Field
		Steps       respjson.Field
		Version     respjson.Field
		WindowMs    respjson.Field
		EntryFilter respjson.Field
		Filter      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelNewResponseQueryDefinition) RawJSON() string { return r.JSON.raw }
func (r *FunnelNewResponseQueryDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelNewResponseQueryDefinitionStep struct {
	ID        string                                     `json:"id" api:"required" format:"uuid"`
	EventName string                                     `json:"eventName" api:"required"`
	Name      string                                     `json:"name" api:"required"`
	Filter    FunnelNewResponseQueryDefinitionStepFilter `json:"filter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventName   respjson.Field
		Name        respjson.Field
		Filter      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelNewResponseQueryDefinitionStep) RawJSON() string { return r.JSON.raw }
func (r *FunnelNewResponseQueryDefinitionStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelNewResponseQueryDefinitionStepFilter struct {
	Filter FunnelNewResponseQueryDefinitionStepFilterFilterUnion `json:"filter" api:"required"`
	// Any of 1.
	Version float64 `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filter      respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelNewResponseQueryDefinitionStepFilter) RawJSON() string { return r.JSON.raw }
func (r *FunnelNewResponseQueryDefinitionStepFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelNewResponseQueryDefinitionStepFilterFilterUnion contains all possible
// properties and values from
// [FunnelNewResponseQueryDefinitionStepFilterFilterObject],
// [FunnelNewResponseQueryDefinitionStepFilterFilterObject2],
// [FunnelNewResponseQueryDefinitionStepFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FunnelNewResponseQueryDefinitionStepFilterFilterUnion struct {
	// This field is from variant
	// [FunnelNewResponseQueryDefinitionStepFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant
	// [FunnelNewResponseQueryDefinitionStepFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant
	// [FunnelNewResponseQueryDefinitionStepFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant
	// [FunnelNewResponseQueryDefinitionStepFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant
	// [FunnelNewResponseQueryDefinitionStepFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant
	// [FunnelNewResponseQueryDefinitionStepFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant
	// [FunnelNewResponseQueryDefinitionStepFilterFilterObject3].
	Value FunnelNewResponseQueryDefinitionStepFilterFilterObject3ValueUnion `json:"value"`
	JSON  struct {
		Children respjson.Field
		Kind     respjson.Field
		Child    respjson.Field
		Operator respjson.Field
		Property respjson.Field
		Path     respjson.Field
		Type     respjson.Field
		Value    respjson.Field
		raw      string
	} `json:"-"`
}

func (u FunnelNewResponseQueryDefinitionStepFilterFilterUnion) AsFunnelNewResponseQueryDefinitionStepFilterFilterObject() (v FunnelNewResponseQueryDefinitionStepFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelNewResponseQueryDefinitionStepFilterFilterUnion) AsFunnelNewResponseQueryDefinitionStepFilterFilterObject2() (v FunnelNewResponseQueryDefinitionStepFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelNewResponseQueryDefinitionStepFilterFilterUnion) AsFunnelNewResponseQueryDefinitionStepFilterFilterObject3() (v FunnelNewResponseQueryDefinitionStepFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelNewResponseQueryDefinitionStepFilterFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *FunnelNewResponseQueryDefinitionStepFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelNewResponseQueryDefinitionStepFilterFilterObject struct {
	Children []any `json:"children" api:"required"`
	// Any of "and", "or".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Children    respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelNewResponseQueryDefinitionStepFilterFilterObject) RawJSON() string { return r.JSON.raw }
func (r *FunnelNewResponseQueryDefinitionStepFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelNewResponseQueryDefinitionStepFilterFilterObject2 struct {
	Child any `json:"child" api:"required"`
	// Any of "not".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Child       respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelNewResponseQueryDefinitionStepFilterFilterObject2) RawJSON() string { return r.JSON.raw }
func (r *FunnelNewResponseQueryDefinitionStepFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelNewResponseQueryDefinitionStepFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path"`
	// Any of "string", "number", "boolean".
	Type  string                                                            `json:"type"`
	Value FunnelNewResponseQueryDefinitionStepFilterFilterObject3ValueUnion `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		Operator    respjson.Field
		Property    respjson.Field
		Path        respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelNewResponseQueryDefinitionStepFilterFilterObject3) RawJSON() string { return r.JSON.raw }
func (r *FunnelNewResponseQueryDefinitionStepFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelNewResponseQueryDefinitionStepFilterFilterObject3ValueUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]FunnelNewResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfFunnelNewResponseQueryDefinitionStepFilterFilterObject3ValueArray]
type FunnelNewResponseQueryDefinitionStepFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]FunnelNewResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion]
	// instead of an object.
	OfFunnelNewResponseQueryDefinitionStepFilterFilterObject3ValueArray []FunnelNewResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                                struct {
		OfString                                                            respjson.Field
		OfFloat                                                             respjson.Field
		OfBool                                                              respjson.Field
		OfFunnelNewResponseQueryDefinitionStepFilterFilterObject3ValueArray respjson.Field
		raw                                                                 string
	} `json:"-"`
}

func (u FunnelNewResponseQueryDefinitionStepFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelNewResponseQueryDefinitionStepFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelNewResponseQueryDefinitionStepFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelNewResponseQueryDefinitionStepFilterFilterObject3ValueUnion) AsFunnelNewResponseQueryDefinitionStepFilterFilterObject3ValueArray() (v []FunnelNewResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelNewResponseQueryDefinitionStepFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelNewResponseQueryDefinitionStepFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelNewResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type FunnelNewResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u FunnelNewResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelNewResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelNewResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelNewResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelNewResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelNewResponseQueryDefinitionEntryFilter struct {
	Filter FunnelNewResponseQueryDefinitionEntryFilterFilterUnion `json:"filter" api:"required"`
	// Any of 1.
	Version float64 `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filter      respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelNewResponseQueryDefinitionEntryFilter) RawJSON() string { return r.JSON.raw }
func (r *FunnelNewResponseQueryDefinitionEntryFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelNewResponseQueryDefinitionEntryFilterFilterUnion contains all possible
// properties and values from
// [FunnelNewResponseQueryDefinitionEntryFilterFilterObject],
// [FunnelNewResponseQueryDefinitionEntryFilterFilterObject2],
// [FunnelNewResponseQueryDefinitionEntryFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FunnelNewResponseQueryDefinitionEntryFilterFilterUnion struct {
	// This field is from variant
	// [FunnelNewResponseQueryDefinitionEntryFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant
	// [FunnelNewResponseQueryDefinitionEntryFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant
	// [FunnelNewResponseQueryDefinitionEntryFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant
	// [FunnelNewResponseQueryDefinitionEntryFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant
	// [FunnelNewResponseQueryDefinitionEntryFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant
	// [FunnelNewResponseQueryDefinitionEntryFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant
	// [FunnelNewResponseQueryDefinitionEntryFilterFilterObject3].
	Value FunnelNewResponseQueryDefinitionEntryFilterFilterObject3ValueUnion `json:"value"`
	JSON  struct {
		Children respjson.Field
		Kind     respjson.Field
		Child    respjson.Field
		Operator respjson.Field
		Property respjson.Field
		Path     respjson.Field
		Type     respjson.Field
		Value    respjson.Field
		raw      string
	} `json:"-"`
}

func (u FunnelNewResponseQueryDefinitionEntryFilterFilterUnion) AsFunnelNewResponseQueryDefinitionEntryFilterFilterObject() (v FunnelNewResponseQueryDefinitionEntryFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelNewResponseQueryDefinitionEntryFilterFilterUnion) AsFunnelNewResponseQueryDefinitionEntryFilterFilterObject2() (v FunnelNewResponseQueryDefinitionEntryFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelNewResponseQueryDefinitionEntryFilterFilterUnion) AsFunnelNewResponseQueryDefinitionEntryFilterFilterObject3() (v FunnelNewResponseQueryDefinitionEntryFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelNewResponseQueryDefinitionEntryFilterFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *FunnelNewResponseQueryDefinitionEntryFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelNewResponseQueryDefinitionEntryFilterFilterObject struct {
	Children []any `json:"children" api:"required"`
	// Any of "and", "or".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Children    respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelNewResponseQueryDefinitionEntryFilterFilterObject) RawJSON() string { return r.JSON.raw }
func (r *FunnelNewResponseQueryDefinitionEntryFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelNewResponseQueryDefinitionEntryFilterFilterObject2 struct {
	Child any `json:"child" api:"required"`
	// Any of "not".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Child       respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelNewResponseQueryDefinitionEntryFilterFilterObject2) RawJSON() string { return r.JSON.raw }
func (r *FunnelNewResponseQueryDefinitionEntryFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelNewResponseQueryDefinitionEntryFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path"`
	// Any of "string", "number", "boolean".
	Type  string                                                             `json:"type"`
	Value FunnelNewResponseQueryDefinitionEntryFilterFilterObject3ValueUnion `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		Operator    respjson.Field
		Property    respjson.Field
		Path        respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelNewResponseQueryDefinitionEntryFilterFilterObject3) RawJSON() string { return r.JSON.raw }
func (r *FunnelNewResponseQueryDefinitionEntryFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelNewResponseQueryDefinitionEntryFilterFilterObject3ValueUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]FunnelNewResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfFunnelNewResponseQueryDefinitionEntryFilterFilterObject3ValueArray]
type FunnelNewResponseQueryDefinitionEntryFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]FunnelNewResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion]
	// instead of an object.
	OfFunnelNewResponseQueryDefinitionEntryFilterFilterObject3ValueArray []FunnelNewResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                                 struct {
		OfString                                                             respjson.Field
		OfFloat                                                              respjson.Field
		OfBool                                                               respjson.Field
		OfFunnelNewResponseQueryDefinitionEntryFilterFilterObject3ValueArray respjson.Field
		raw                                                                  string
	} `json:"-"`
}

func (u FunnelNewResponseQueryDefinitionEntryFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelNewResponseQueryDefinitionEntryFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelNewResponseQueryDefinitionEntryFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelNewResponseQueryDefinitionEntryFilterFilterObject3ValueUnion) AsFunnelNewResponseQueryDefinitionEntryFilterFilterObject3ValueArray() (v []FunnelNewResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelNewResponseQueryDefinitionEntryFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelNewResponseQueryDefinitionEntryFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelNewResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type FunnelNewResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u FunnelNewResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelNewResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelNewResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelNewResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelNewResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelNewResponseQueryDefinitionFilter struct {
	Filter FunnelNewResponseQueryDefinitionFilterFilterUnion `json:"filter" api:"required"`
	// Any of 1.
	Version float64 `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filter      respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelNewResponseQueryDefinitionFilter) RawJSON() string { return r.JSON.raw }
func (r *FunnelNewResponseQueryDefinitionFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelNewResponseQueryDefinitionFilterFilterUnion contains all possible
// properties and values from [FunnelNewResponseQueryDefinitionFilterFilterObject],
// [FunnelNewResponseQueryDefinitionFilterFilterObject2],
// [FunnelNewResponseQueryDefinitionFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FunnelNewResponseQueryDefinitionFilterFilterUnion struct {
	// This field is from variant [FunnelNewResponseQueryDefinitionFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant
	// [FunnelNewResponseQueryDefinitionFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant
	// [FunnelNewResponseQueryDefinitionFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant
	// [FunnelNewResponseQueryDefinitionFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant
	// [FunnelNewResponseQueryDefinitionFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant
	// [FunnelNewResponseQueryDefinitionFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant
	// [FunnelNewResponseQueryDefinitionFilterFilterObject3].
	Value FunnelNewResponseQueryDefinitionFilterFilterObject3ValueUnion `json:"value"`
	JSON  struct {
		Children respjson.Field
		Kind     respjson.Field
		Child    respjson.Field
		Operator respjson.Field
		Property respjson.Field
		Path     respjson.Field
		Type     respjson.Field
		Value    respjson.Field
		raw      string
	} `json:"-"`
}

func (u FunnelNewResponseQueryDefinitionFilterFilterUnion) AsFunnelNewResponseQueryDefinitionFilterFilterObject() (v FunnelNewResponseQueryDefinitionFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelNewResponseQueryDefinitionFilterFilterUnion) AsFunnelNewResponseQueryDefinitionFilterFilterObject2() (v FunnelNewResponseQueryDefinitionFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelNewResponseQueryDefinitionFilterFilterUnion) AsFunnelNewResponseQueryDefinitionFilterFilterObject3() (v FunnelNewResponseQueryDefinitionFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelNewResponseQueryDefinitionFilterFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *FunnelNewResponseQueryDefinitionFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelNewResponseQueryDefinitionFilterFilterObject struct {
	Children []any `json:"children" api:"required"`
	// Any of "and", "or".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Children    respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelNewResponseQueryDefinitionFilterFilterObject) RawJSON() string { return r.JSON.raw }
func (r *FunnelNewResponseQueryDefinitionFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelNewResponseQueryDefinitionFilterFilterObject2 struct {
	Child any `json:"child" api:"required"`
	// Any of "not".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Child       respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelNewResponseQueryDefinitionFilterFilterObject2) RawJSON() string { return r.JSON.raw }
func (r *FunnelNewResponseQueryDefinitionFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelNewResponseQueryDefinitionFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path"`
	// Any of "string", "number", "boolean".
	Type  string                                                        `json:"type"`
	Value FunnelNewResponseQueryDefinitionFilterFilterObject3ValueUnion `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		Operator    respjson.Field
		Property    respjson.Field
		Path        respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelNewResponseQueryDefinitionFilterFilterObject3) RawJSON() string { return r.JSON.raw }
func (r *FunnelNewResponseQueryDefinitionFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelNewResponseQueryDefinitionFilterFilterObject3ValueUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]FunnelNewResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfFunnelNewResponseQueryDefinitionFilterFilterObject3ValueArray]
type FunnelNewResponseQueryDefinitionFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]FunnelNewResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion]
	// instead of an object.
	OfFunnelNewResponseQueryDefinitionFilterFilterObject3ValueArray []FunnelNewResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                            struct {
		OfString                                                        respjson.Field
		OfFloat                                                         respjson.Field
		OfBool                                                          respjson.Field
		OfFunnelNewResponseQueryDefinitionFilterFilterObject3ValueArray respjson.Field
		raw                                                             string
	} `json:"-"`
}

func (u FunnelNewResponseQueryDefinitionFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelNewResponseQueryDefinitionFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelNewResponseQueryDefinitionFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelNewResponseQueryDefinitionFilterFilterObject3ValueUnion) AsFunnelNewResponseQueryDefinitionFilterFilterObject3ValueArray() (v []FunnelNewResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelNewResponseQueryDefinitionFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelNewResponseQueryDefinitionFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelNewResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion contains
// all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type FunnelNewResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u FunnelNewResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelNewResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelNewResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelNewResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelNewResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelNewResponseStatus string

const (
	FunnelNewResponseStatusReady      FunnelNewResponseStatus = "READY"
	FunnelNewResponseStatusProcessing FunnelNewResponseStatus = "PROCESSING"
)

type FunnelNewResponseReportDateRange struct {
	From string `json:"from" api:"required"`
	To   string `json:"to" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelNewResponseReportDateRange) RawJSON() string { return r.JSON.raw }
func (r *FunnelNewResponseReportDateRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Funnel configuration details.
type FunnelGetResponse struct {
	CreatedAt string `json:"createdAt" api:"required"`
	FunnelID  string `json:"funnelId" api:"required" format:"uuid"`
	Name      string `json:"name" api:"required"`
	// Versioned funnel definition with typed steps and filters.
	QueryDefinition FunnelGetResponseQueryDefinition `json:"queryDefinition" api:"required"`
	// Any of "READY", "PROCESSING".
	Status          FunnelGetResponseStatus          `json:"status" api:"required"`
	UpdatedAt       string                           `json:"updatedAt" api:"required"`
	Description     string                           `json:"description" api:"nullable"`
	ReportDateRange FunnelGetResponseReportDateRange `json:"reportDateRange" api:"nullable"`
	Watched         bool                             `json:"watched" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt       respjson.Field
		FunnelID        respjson.Field
		Name            respjson.Field
		QueryDefinition respjson.Field
		Status          respjson.Field
		UpdatedAt       respjson.Field
		Description     respjson.Field
		ReportDateRange respjson.Field
		Watched         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelGetResponse) RawJSON() string { return r.JSON.raw }
func (r *FunnelGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Versioned funnel definition with typed steps and filters.
type FunnelGetResponseQueryDefinition struct {
	// Any of "unique", "events", "sessions".
	Counting string `json:"counting" api:"required"`
	// Any of "session", "visitor".
	Identity string `json:"identity" api:"required"`
	// Any of "funnel".
	Kind string `json:"kind" api:"required"`
	// Any of "ordered", "any".
	Order string `json:"order" api:"required"`
	// Any of "first", "best".
	Reentry string                                 `json:"reentry" api:"required"`
	Steps   []FunnelGetResponseQueryDefinitionStep `json:"steps" api:"required"`
	// Any of 1.
	Version     float64                                     `json:"version" api:"required"`
	WindowMs    int64                                       `json:"windowMs" api:"required"`
	EntryFilter FunnelGetResponseQueryDefinitionEntryFilter `json:"entryFilter"`
	Filter      FunnelGetResponseQueryDefinitionFilter      `json:"filter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Counting    respjson.Field
		Identity    respjson.Field
		Kind        respjson.Field
		Order       respjson.Field
		Reentry     respjson.Field
		Steps       respjson.Field
		Version     respjson.Field
		WindowMs    respjson.Field
		EntryFilter respjson.Field
		Filter      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelGetResponseQueryDefinition) RawJSON() string { return r.JSON.raw }
func (r *FunnelGetResponseQueryDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelGetResponseQueryDefinitionStep struct {
	ID        string                                     `json:"id" api:"required" format:"uuid"`
	EventName string                                     `json:"eventName" api:"required"`
	Name      string                                     `json:"name" api:"required"`
	Filter    FunnelGetResponseQueryDefinitionStepFilter `json:"filter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventName   respjson.Field
		Name        respjson.Field
		Filter      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelGetResponseQueryDefinitionStep) RawJSON() string { return r.JSON.raw }
func (r *FunnelGetResponseQueryDefinitionStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelGetResponseQueryDefinitionStepFilter struct {
	Filter FunnelGetResponseQueryDefinitionStepFilterFilterUnion `json:"filter" api:"required"`
	// Any of 1.
	Version float64 `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filter      respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelGetResponseQueryDefinitionStepFilter) RawJSON() string { return r.JSON.raw }
func (r *FunnelGetResponseQueryDefinitionStepFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelGetResponseQueryDefinitionStepFilterFilterUnion contains all possible
// properties and values from
// [FunnelGetResponseQueryDefinitionStepFilterFilterObject],
// [FunnelGetResponseQueryDefinitionStepFilterFilterObject2],
// [FunnelGetResponseQueryDefinitionStepFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FunnelGetResponseQueryDefinitionStepFilterFilterUnion struct {
	// This field is from variant
	// [FunnelGetResponseQueryDefinitionStepFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant
	// [FunnelGetResponseQueryDefinitionStepFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant
	// [FunnelGetResponseQueryDefinitionStepFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant
	// [FunnelGetResponseQueryDefinitionStepFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant
	// [FunnelGetResponseQueryDefinitionStepFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant
	// [FunnelGetResponseQueryDefinitionStepFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant
	// [FunnelGetResponseQueryDefinitionStepFilterFilterObject3].
	Value FunnelGetResponseQueryDefinitionStepFilterFilterObject3ValueUnion `json:"value"`
	JSON  struct {
		Children respjson.Field
		Kind     respjson.Field
		Child    respjson.Field
		Operator respjson.Field
		Property respjson.Field
		Path     respjson.Field
		Type     respjson.Field
		Value    respjson.Field
		raw      string
	} `json:"-"`
}

func (u FunnelGetResponseQueryDefinitionStepFilterFilterUnion) AsFunnelGetResponseQueryDefinitionStepFilterFilterObject() (v FunnelGetResponseQueryDefinitionStepFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelGetResponseQueryDefinitionStepFilterFilterUnion) AsFunnelGetResponseQueryDefinitionStepFilterFilterObject2() (v FunnelGetResponseQueryDefinitionStepFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelGetResponseQueryDefinitionStepFilterFilterUnion) AsFunnelGetResponseQueryDefinitionStepFilterFilterObject3() (v FunnelGetResponseQueryDefinitionStepFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelGetResponseQueryDefinitionStepFilterFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *FunnelGetResponseQueryDefinitionStepFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelGetResponseQueryDefinitionStepFilterFilterObject struct {
	Children []any `json:"children" api:"required"`
	// Any of "and", "or".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Children    respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelGetResponseQueryDefinitionStepFilterFilterObject) RawJSON() string { return r.JSON.raw }
func (r *FunnelGetResponseQueryDefinitionStepFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelGetResponseQueryDefinitionStepFilterFilterObject2 struct {
	Child any `json:"child" api:"required"`
	// Any of "not".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Child       respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelGetResponseQueryDefinitionStepFilterFilterObject2) RawJSON() string { return r.JSON.raw }
func (r *FunnelGetResponseQueryDefinitionStepFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelGetResponseQueryDefinitionStepFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path"`
	// Any of "string", "number", "boolean".
	Type  string                                                            `json:"type"`
	Value FunnelGetResponseQueryDefinitionStepFilterFilterObject3ValueUnion `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		Operator    respjson.Field
		Property    respjson.Field
		Path        respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelGetResponseQueryDefinitionStepFilterFilterObject3) RawJSON() string { return r.JSON.raw }
func (r *FunnelGetResponseQueryDefinitionStepFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelGetResponseQueryDefinitionStepFilterFilterObject3ValueUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]FunnelGetResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfFunnelGetResponseQueryDefinitionStepFilterFilterObject3ValueArray]
type FunnelGetResponseQueryDefinitionStepFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]FunnelGetResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion]
	// instead of an object.
	OfFunnelGetResponseQueryDefinitionStepFilterFilterObject3ValueArray []FunnelGetResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                                struct {
		OfString                                                            respjson.Field
		OfFloat                                                             respjson.Field
		OfBool                                                              respjson.Field
		OfFunnelGetResponseQueryDefinitionStepFilterFilterObject3ValueArray respjson.Field
		raw                                                                 string
	} `json:"-"`
}

func (u FunnelGetResponseQueryDefinitionStepFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelGetResponseQueryDefinitionStepFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelGetResponseQueryDefinitionStepFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelGetResponseQueryDefinitionStepFilterFilterObject3ValueUnion) AsFunnelGetResponseQueryDefinitionStepFilterFilterObject3ValueArray() (v []FunnelGetResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelGetResponseQueryDefinitionStepFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelGetResponseQueryDefinitionStepFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelGetResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type FunnelGetResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u FunnelGetResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelGetResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelGetResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelGetResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelGetResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelGetResponseQueryDefinitionEntryFilter struct {
	Filter FunnelGetResponseQueryDefinitionEntryFilterFilterUnion `json:"filter" api:"required"`
	// Any of 1.
	Version float64 `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filter      respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelGetResponseQueryDefinitionEntryFilter) RawJSON() string { return r.JSON.raw }
func (r *FunnelGetResponseQueryDefinitionEntryFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelGetResponseQueryDefinitionEntryFilterFilterUnion contains all possible
// properties and values from
// [FunnelGetResponseQueryDefinitionEntryFilterFilterObject],
// [FunnelGetResponseQueryDefinitionEntryFilterFilterObject2],
// [FunnelGetResponseQueryDefinitionEntryFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FunnelGetResponseQueryDefinitionEntryFilterFilterUnion struct {
	// This field is from variant
	// [FunnelGetResponseQueryDefinitionEntryFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant
	// [FunnelGetResponseQueryDefinitionEntryFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant
	// [FunnelGetResponseQueryDefinitionEntryFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant
	// [FunnelGetResponseQueryDefinitionEntryFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant
	// [FunnelGetResponseQueryDefinitionEntryFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant
	// [FunnelGetResponseQueryDefinitionEntryFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant
	// [FunnelGetResponseQueryDefinitionEntryFilterFilterObject3].
	Value FunnelGetResponseQueryDefinitionEntryFilterFilterObject3ValueUnion `json:"value"`
	JSON  struct {
		Children respjson.Field
		Kind     respjson.Field
		Child    respjson.Field
		Operator respjson.Field
		Property respjson.Field
		Path     respjson.Field
		Type     respjson.Field
		Value    respjson.Field
		raw      string
	} `json:"-"`
}

func (u FunnelGetResponseQueryDefinitionEntryFilterFilterUnion) AsFunnelGetResponseQueryDefinitionEntryFilterFilterObject() (v FunnelGetResponseQueryDefinitionEntryFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelGetResponseQueryDefinitionEntryFilterFilterUnion) AsFunnelGetResponseQueryDefinitionEntryFilterFilterObject2() (v FunnelGetResponseQueryDefinitionEntryFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelGetResponseQueryDefinitionEntryFilterFilterUnion) AsFunnelGetResponseQueryDefinitionEntryFilterFilterObject3() (v FunnelGetResponseQueryDefinitionEntryFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelGetResponseQueryDefinitionEntryFilterFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *FunnelGetResponseQueryDefinitionEntryFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelGetResponseQueryDefinitionEntryFilterFilterObject struct {
	Children []any `json:"children" api:"required"`
	// Any of "and", "or".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Children    respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelGetResponseQueryDefinitionEntryFilterFilterObject) RawJSON() string { return r.JSON.raw }
func (r *FunnelGetResponseQueryDefinitionEntryFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelGetResponseQueryDefinitionEntryFilterFilterObject2 struct {
	Child any `json:"child" api:"required"`
	// Any of "not".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Child       respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelGetResponseQueryDefinitionEntryFilterFilterObject2) RawJSON() string { return r.JSON.raw }
func (r *FunnelGetResponseQueryDefinitionEntryFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelGetResponseQueryDefinitionEntryFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path"`
	// Any of "string", "number", "boolean".
	Type  string                                                             `json:"type"`
	Value FunnelGetResponseQueryDefinitionEntryFilterFilterObject3ValueUnion `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		Operator    respjson.Field
		Property    respjson.Field
		Path        respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelGetResponseQueryDefinitionEntryFilterFilterObject3) RawJSON() string { return r.JSON.raw }
func (r *FunnelGetResponseQueryDefinitionEntryFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelGetResponseQueryDefinitionEntryFilterFilterObject3ValueUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]FunnelGetResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfFunnelGetResponseQueryDefinitionEntryFilterFilterObject3ValueArray]
type FunnelGetResponseQueryDefinitionEntryFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]FunnelGetResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion]
	// instead of an object.
	OfFunnelGetResponseQueryDefinitionEntryFilterFilterObject3ValueArray []FunnelGetResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                                 struct {
		OfString                                                             respjson.Field
		OfFloat                                                              respjson.Field
		OfBool                                                               respjson.Field
		OfFunnelGetResponseQueryDefinitionEntryFilterFilterObject3ValueArray respjson.Field
		raw                                                                  string
	} `json:"-"`
}

func (u FunnelGetResponseQueryDefinitionEntryFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelGetResponseQueryDefinitionEntryFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelGetResponseQueryDefinitionEntryFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelGetResponseQueryDefinitionEntryFilterFilterObject3ValueUnion) AsFunnelGetResponseQueryDefinitionEntryFilterFilterObject3ValueArray() (v []FunnelGetResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelGetResponseQueryDefinitionEntryFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelGetResponseQueryDefinitionEntryFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelGetResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type FunnelGetResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u FunnelGetResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelGetResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelGetResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelGetResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelGetResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelGetResponseQueryDefinitionFilter struct {
	Filter FunnelGetResponseQueryDefinitionFilterFilterUnion `json:"filter" api:"required"`
	// Any of 1.
	Version float64 `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filter      respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelGetResponseQueryDefinitionFilter) RawJSON() string { return r.JSON.raw }
func (r *FunnelGetResponseQueryDefinitionFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelGetResponseQueryDefinitionFilterFilterUnion contains all possible
// properties and values from [FunnelGetResponseQueryDefinitionFilterFilterObject],
// [FunnelGetResponseQueryDefinitionFilterFilterObject2],
// [FunnelGetResponseQueryDefinitionFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FunnelGetResponseQueryDefinitionFilterFilterUnion struct {
	// This field is from variant [FunnelGetResponseQueryDefinitionFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant
	// [FunnelGetResponseQueryDefinitionFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant
	// [FunnelGetResponseQueryDefinitionFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant
	// [FunnelGetResponseQueryDefinitionFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant
	// [FunnelGetResponseQueryDefinitionFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant
	// [FunnelGetResponseQueryDefinitionFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant
	// [FunnelGetResponseQueryDefinitionFilterFilterObject3].
	Value FunnelGetResponseQueryDefinitionFilterFilterObject3ValueUnion `json:"value"`
	JSON  struct {
		Children respjson.Field
		Kind     respjson.Field
		Child    respjson.Field
		Operator respjson.Field
		Property respjson.Field
		Path     respjson.Field
		Type     respjson.Field
		Value    respjson.Field
		raw      string
	} `json:"-"`
}

func (u FunnelGetResponseQueryDefinitionFilterFilterUnion) AsFunnelGetResponseQueryDefinitionFilterFilterObject() (v FunnelGetResponseQueryDefinitionFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelGetResponseQueryDefinitionFilterFilterUnion) AsFunnelGetResponseQueryDefinitionFilterFilterObject2() (v FunnelGetResponseQueryDefinitionFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelGetResponseQueryDefinitionFilterFilterUnion) AsFunnelGetResponseQueryDefinitionFilterFilterObject3() (v FunnelGetResponseQueryDefinitionFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelGetResponseQueryDefinitionFilterFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *FunnelGetResponseQueryDefinitionFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelGetResponseQueryDefinitionFilterFilterObject struct {
	Children []any `json:"children" api:"required"`
	// Any of "and", "or".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Children    respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelGetResponseQueryDefinitionFilterFilterObject) RawJSON() string { return r.JSON.raw }
func (r *FunnelGetResponseQueryDefinitionFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelGetResponseQueryDefinitionFilterFilterObject2 struct {
	Child any `json:"child" api:"required"`
	// Any of "not".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Child       respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelGetResponseQueryDefinitionFilterFilterObject2) RawJSON() string { return r.JSON.raw }
func (r *FunnelGetResponseQueryDefinitionFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelGetResponseQueryDefinitionFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path"`
	// Any of "string", "number", "boolean".
	Type  string                                                        `json:"type"`
	Value FunnelGetResponseQueryDefinitionFilterFilterObject3ValueUnion `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		Operator    respjson.Field
		Property    respjson.Field
		Path        respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelGetResponseQueryDefinitionFilterFilterObject3) RawJSON() string { return r.JSON.raw }
func (r *FunnelGetResponseQueryDefinitionFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelGetResponseQueryDefinitionFilterFilterObject3ValueUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]FunnelGetResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfFunnelGetResponseQueryDefinitionFilterFilterObject3ValueArray]
type FunnelGetResponseQueryDefinitionFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]FunnelGetResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion]
	// instead of an object.
	OfFunnelGetResponseQueryDefinitionFilterFilterObject3ValueArray []FunnelGetResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                            struct {
		OfString                                                        respjson.Field
		OfFloat                                                         respjson.Field
		OfBool                                                          respjson.Field
		OfFunnelGetResponseQueryDefinitionFilterFilterObject3ValueArray respjson.Field
		raw                                                             string
	} `json:"-"`
}

func (u FunnelGetResponseQueryDefinitionFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelGetResponseQueryDefinitionFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelGetResponseQueryDefinitionFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelGetResponseQueryDefinitionFilterFilterObject3ValueUnion) AsFunnelGetResponseQueryDefinitionFilterFilterObject3ValueArray() (v []FunnelGetResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelGetResponseQueryDefinitionFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelGetResponseQueryDefinitionFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelGetResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion contains
// all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type FunnelGetResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u FunnelGetResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelGetResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelGetResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelGetResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelGetResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelGetResponseStatus string

const (
	FunnelGetResponseStatusReady      FunnelGetResponseStatus = "READY"
	FunnelGetResponseStatusProcessing FunnelGetResponseStatus = "PROCESSING"
)

type FunnelGetResponseReportDateRange struct {
	From string `json:"from" api:"required"`
	To   string `json:"to" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelGetResponseReportDateRange) RawJSON() string { return r.JSON.raw }
func (r *FunnelGetResponseReportDateRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelUpdateResponse struct {
	CreatedAt string `json:"createdAt" api:"required"`
	FunnelID  string `json:"funnelId" api:"required" format:"uuid"`
	Name      string `json:"name" api:"required"`
	// Versioned funnel definition with typed steps and filters.
	QueryDefinition FunnelUpdateResponseQueryDefinition `json:"queryDefinition" api:"required"`
	// Any of "READY", "PROCESSING".
	Status          FunnelUpdateResponseStatus          `json:"status" api:"required"`
	UpdatedAt       string                              `json:"updatedAt" api:"required"`
	Description     string                              `json:"description" api:"nullable"`
	ReportDateRange FunnelUpdateResponseReportDateRange `json:"reportDateRange" api:"nullable"`
	Watched         bool                                `json:"watched" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt       respjson.Field
		FunnelID        respjson.Field
		Name            respjson.Field
		QueryDefinition respjson.Field
		Status          respjson.Field
		UpdatedAt       respjson.Field
		Description     respjson.Field
		ReportDateRange respjson.Field
		Watched         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *FunnelUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Versioned funnel definition with typed steps and filters.
type FunnelUpdateResponseQueryDefinition struct {
	// Any of "unique", "events", "sessions".
	Counting string `json:"counting" api:"required"`
	// Any of "session", "visitor".
	Identity string `json:"identity" api:"required"`
	// Any of "funnel".
	Kind string `json:"kind" api:"required"`
	// Any of "ordered", "any".
	Order string `json:"order" api:"required"`
	// Any of "first", "best".
	Reentry string                                    `json:"reentry" api:"required"`
	Steps   []FunnelUpdateResponseQueryDefinitionStep `json:"steps" api:"required"`
	// Any of 1.
	Version     float64                                        `json:"version" api:"required"`
	WindowMs    int64                                          `json:"windowMs" api:"required"`
	EntryFilter FunnelUpdateResponseQueryDefinitionEntryFilter `json:"entryFilter"`
	Filter      FunnelUpdateResponseQueryDefinitionFilter      `json:"filter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Counting    respjson.Field
		Identity    respjson.Field
		Kind        respjson.Field
		Order       respjson.Field
		Reentry     respjson.Field
		Steps       respjson.Field
		Version     respjson.Field
		WindowMs    respjson.Field
		EntryFilter respjson.Field
		Filter      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelUpdateResponseQueryDefinition) RawJSON() string { return r.JSON.raw }
func (r *FunnelUpdateResponseQueryDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelUpdateResponseQueryDefinitionStep struct {
	ID        string                                        `json:"id" api:"required" format:"uuid"`
	EventName string                                        `json:"eventName" api:"required"`
	Name      string                                        `json:"name" api:"required"`
	Filter    FunnelUpdateResponseQueryDefinitionStepFilter `json:"filter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventName   respjson.Field
		Name        respjson.Field
		Filter      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelUpdateResponseQueryDefinitionStep) RawJSON() string { return r.JSON.raw }
func (r *FunnelUpdateResponseQueryDefinitionStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelUpdateResponseQueryDefinitionStepFilter struct {
	Filter FunnelUpdateResponseQueryDefinitionStepFilterFilterUnion `json:"filter" api:"required"`
	// Any of 1.
	Version float64 `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filter      respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelUpdateResponseQueryDefinitionStepFilter) RawJSON() string { return r.JSON.raw }
func (r *FunnelUpdateResponseQueryDefinitionStepFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelUpdateResponseQueryDefinitionStepFilterFilterUnion contains all possible
// properties and values from
// [FunnelUpdateResponseQueryDefinitionStepFilterFilterObject],
// [FunnelUpdateResponseQueryDefinitionStepFilterFilterObject2],
// [FunnelUpdateResponseQueryDefinitionStepFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FunnelUpdateResponseQueryDefinitionStepFilterFilterUnion struct {
	// This field is from variant
	// [FunnelUpdateResponseQueryDefinitionStepFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant
	// [FunnelUpdateResponseQueryDefinitionStepFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant
	// [FunnelUpdateResponseQueryDefinitionStepFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant
	// [FunnelUpdateResponseQueryDefinitionStepFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant
	// [FunnelUpdateResponseQueryDefinitionStepFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant
	// [FunnelUpdateResponseQueryDefinitionStepFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant
	// [FunnelUpdateResponseQueryDefinitionStepFilterFilterObject3].
	Value FunnelUpdateResponseQueryDefinitionStepFilterFilterObject3ValueUnion `json:"value"`
	JSON  struct {
		Children respjson.Field
		Kind     respjson.Field
		Child    respjson.Field
		Operator respjson.Field
		Property respjson.Field
		Path     respjson.Field
		Type     respjson.Field
		Value    respjson.Field
		raw      string
	} `json:"-"`
}

func (u FunnelUpdateResponseQueryDefinitionStepFilterFilterUnion) AsFunnelUpdateResponseQueryDefinitionStepFilterFilterObject() (v FunnelUpdateResponseQueryDefinitionStepFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelUpdateResponseQueryDefinitionStepFilterFilterUnion) AsFunnelUpdateResponseQueryDefinitionStepFilterFilterObject2() (v FunnelUpdateResponseQueryDefinitionStepFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelUpdateResponseQueryDefinitionStepFilterFilterUnion) AsFunnelUpdateResponseQueryDefinitionStepFilterFilterObject3() (v FunnelUpdateResponseQueryDefinitionStepFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelUpdateResponseQueryDefinitionStepFilterFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *FunnelUpdateResponseQueryDefinitionStepFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelUpdateResponseQueryDefinitionStepFilterFilterObject struct {
	Children []any `json:"children" api:"required"`
	// Any of "and", "or".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Children    respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelUpdateResponseQueryDefinitionStepFilterFilterObject) RawJSON() string {
	return r.JSON.raw
}
func (r *FunnelUpdateResponseQueryDefinitionStepFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelUpdateResponseQueryDefinitionStepFilterFilterObject2 struct {
	Child any `json:"child" api:"required"`
	// Any of "not".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Child       respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelUpdateResponseQueryDefinitionStepFilterFilterObject2) RawJSON() string {
	return r.JSON.raw
}
func (r *FunnelUpdateResponseQueryDefinitionStepFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelUpdateResponseQueryDefinitionStepFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path"`
	// Any of "string", "number", "boolean".
	Type  string                                                               `json:"type"`
	Value FunnelUpdateResponseQueryDefinitionStepFilterFilterObject3ValueUnion `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		Operator    respjson.Field
		Property    respjson.Field
		Path        respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelUpdateResponseQueryDefinitionStepFilterFilterObject3) RawJSON() string {
	return r.JSON.raw
}
func (r *FunnelUpdateResponseQueryDefinitionStepFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelUpdateResponseQueryDefinitionStepFilterFilterObject3ValueUnion contains
// all possible properties and values from [string], [float64], [bool],
// [[]FunnelUpdateResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfFunnelUpdateResponseQueryDefinitionStepFilterFilterObject3ValueArray]
type FunnelUpdateResponseQueryDefinitionStepFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]FunnelUpdateResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion]
	// instead of an object.
	OfFunnelUpdateResponseQueryDefinitionStepFilterFilterObject3ValueArray []FunnelUpdateResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                                   struct {
		OfString                                                               respjson.Field
		OfFloat                                                                respjson.Field
		OfBool                                                                 respjson.Field
		OfFunnelUpdateResponseQueryDefinitionStepFilterFilterObject3ValueArray respjson.Field
		raw                                                                    string
	} `json:"-"`
}

func (u FunnelUpdateResponseQueryDefinitionStepFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelUpdateResponseQueryDefinitionStepFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelUpdateResponseQueryDefinitionStepFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelUpdateResponseQueryDefinitionStepFilterFilterObject3ValueUnion) AsFunnelUpdateResponseQueryDefinitionStepFilterFilterObject3ValueArray() (v []FunnelUpdateResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelUpdateResponseQueryDefinitionStepFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelUpdateResponseQueryDefinitionStepFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelUpdateResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type FunnelUpdateResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u FunnelUpdateResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelUpdateResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelUpdateResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelUpdateResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelUpdateResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelUpdateResponseQueryDefinitionEntryFilter struct {
	Filter FunnelUpdateResponseQueryDefinitionEntryFilterFilterUnion `json:"filter" api:"required"`
	// Any of 1.
	Version float64 `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filter      respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelUpdateResponseQueryDefinitionEntryFilter) RawJSON() string { return r.JSON.raw }
func (r *FunnelUpdateResponseQueryDefinitionEntryFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelUpdateResponseQueryDefinitionEntryFilterFilterUnion contains all possible
// properties and values from
// [FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject],
// [FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject2],
// [FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FunnelUpdateResponseQueryDefinitionEntryFilterFilterUnion struct {
	// This field is from variant
	// [FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant
	// [FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant
	// [FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant
	// [FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant
	// [FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant
	// [FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant
	// [FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3].
	Value FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3ValueUnion `json:"value"`
	JSON  struct {
		Children respjson.Field
		Kind     respjson.Field
		Child    respjson.Field
		Operator respjson.Field
		Property respjson.Field
		Path     respjson.Field
		Type     respjson.Field
		Value    respjson.Field
		raw      string
	} `json:"-"`
}

func (u FunnelUpdateResponseQueryDefinitionEntryFilterFilterUnion) AsFunnelUpdateResponseQueryDefinitionEntryFilterFilterObject() (v FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelUpdateResponseQueryDefinitionEntryFilterFilterUnion) AsFunnelUpdateResponseQueryDefinitionEntryFilterFilterObject2() (v FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelUpdateResponseQueryDefinitionEntryFilterFilterUnion) AsFunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3() (v FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelUpdateResponseQueryDefinitionEntryFilterFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelUpdateResponseQueryDefinitionEntryFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject struct {
	Children []any `json:"children" api:"required"`
	// Any of "and", "or".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Children    respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject) RawJSON() string {
	return r.JSON.raw
}
func (r *FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject2 struct {
	Child any `json:"child" api:"required"`
	// Any of "not".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Child       respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject2) RawJSON() string {
	return r.JSON.raw
}
func (r *FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path"`
	// Any of "string", "number", "boolean".
	Type  string                                                                `json:"type"`
	Value FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3ValueUnion `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		Operator    respjson.Field
		Property    respjson.Field
		Path        respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3) RawJSON() string {
	return r.JSON.raw
}
func (r *FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3ValueUnion contains
// all possible properties and values from [string], [float64], [bool],
// [[]FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfFunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3ValueArray]
type FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion]
	// instead of an object.
	OfFunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3ValueArray []FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                                    struct {
		OfString                                                                respjson.Field
		OfFloat                                                                 respjson.Field
		OfBool                                                                  respjson.Field
		OfFunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3ValueArray respjson.Field
		raw                                                                     string
	} `json:"-"`
}

func (u FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3ValueUnion) AsFunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3ValueArray() (v []FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelUpdateResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelUpdateResponseQueryDefinitionFilter struct {
	Filter FunnelUpdateResponseQueryDefinitionFilterFilterUnion `json:"filter" api:"required"`
	// Any of 1.
	Version float64 `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filter      respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelUpdateResponseQueryDefinitionFilter) RawJSON() string { return r.JSON.raw }
func (r *FunnelUpdateResponseQueryDefinitionFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelUpdateResponseQueryDefinitionFilterFilterUnion contains all possible
// properties and values from
// [FunnelUpdateResponseQueryDefinitionFilterFilterObject],
// [FunnelUpdateResponseQueryDefinitionFilterFilterObject2],
// [FunnelUpdateResponseQueryDefinitionFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FunnelUpdateResponseQueryDefinitionFilterFilterUnion struct {
	// This field is from variant
	// [FunnelUpdateResponseQueryDefinitionFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant
	// [FunnelUpdateResponseQueryDefinitionFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant
	// [FunnelUpdateResponseQueryDefinitionFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant
	// [FunnelUpdateResponseQueryDefinitionFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant
	// [FunnelUpdateResponseQueryDefinitionFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant
	// [FunnelUpdateResponseQueryDefinitionFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant
	// [FunnelUpdateResponseQueryDefinitionFilterFilterObject3].
	Value FunnelUpdateResponseQueryDefinitionFilterFilterObject3ValueUnion `json:"value"`
	JSON  struct {
		Children respjson.Field
		Kind     respjson.Field
		Child    respjson.Field
		Operator respjson.Field
		Property respjson.Field
		Path     respjson.Field
		Type     respjson.Field
		Value    respjson.Field
		raw      string
	} `json:"-"`
}

func (u FunnelUpdateResponseQueryDefinitionFilterFilterUnion) AsFunnelUpdateResponseQueryDefinitionFilterFilterObject() (v FunnelUpdateResponseQueryDefinitionFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelUpdateResponseQueryDefinitionFilterFilterUnion) AsFunnelUpdateResponseQueryDefinitionFilterFilterObject2() (v FunnelUpdateResponseQueryDefinitionFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelUpdateResponseQueryDefinitionFilterFilterUnion) AsFunnelUpdateResponseQueryDefinitionFilterFilterObject3() (v FunnelUpdateResponseQueryDefinitionFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelUpdateResponseQueryDefinitionFilterFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *FunnelUpdateResponseQueryDefinitionFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelUpdateResponseQueryDefinitionFilterFilterObject struct {
	Children []any `json:"children" api:"required"`
	// Any of "and", "or".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Children    respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelUpdateResponseQueryDefinitionFilterFilterObject) RawJSON() string { return r.JSON.raw }
func (r *FunnelUpdateResponseQueryDefinitionFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelUpdateResponseQueryDefinitionFilterFilterObject2 struct {
	Child any `json:"child" api:"required"`
	// Any of "not".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Child       respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelUpdateResponseQueryDefinitionFilterFilterObject2) RawJSON() string { return r.JSON.raw }
func (r *FunnelUpdateResponseQueryDefinitionFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelUpdateResponseQueryDefinitionFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path"`
	// Any of "string", "number", "boolean".
	Type  string                                                           `json:"type"`
	Value FunnelUpdateResponseQueryDefinitionFilterFilterObject3ValueUnion `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		Operator    respjson.Field
		Property    respjson.Field
		Path        respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelUpdateResponseQueryDefinitionFilterFilterObject3) RawJSON() string { return r.JSON.raw }
func (r *FunnelUpdateResponseQueryDefinitionFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelUpdateResponseQueryDefinitionFilterFilterObject3ValueUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]FunnelUpdateResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfFunnelUpdateResponseQueryDefinitionFilterFilterObject3ValueArray]
type FunnelUpdateResponseQueryDefinitionFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]FunnelUpdateResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion]
	// instead of an object.
	OfFunnelUpdateResponseQueryDefinitionFilterFilterObject3ValueArray []FunnelUpdateResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                               struct {
		OfString                                                           respjson.Field
		OfFloat                                                            respjson.Field
		OfBool                                                             respjson.Field
		OfFunnelUpdateResponseQueryDefinitionFilterFilterObject3ValueArray respjson.Field
		raw                                                                string
	} `json:"-"`
}

func (u FunnelUpdateResponseQueryDefinitionFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelUpdateResponseQueryDefinitionFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelUpdateResponseQueryDefinitionFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelUpdateResponseQueryDefinitionFilterFilterObject3ValueUnion) AsFunnelUpdateResponseQueryDefinitionFilterFilterObject3ValueArray() (v []FunnelUpdateResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelUpdateResponseQueryDefinitionFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelUpdateResponseQueryDefinitionFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelUpdateResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type FunnelUpdateResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u FunnelUpdateResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelUpdateResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelUpdateResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelUpdateResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelUpdateResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelUpdateResponseStatus string

const (
	FunnelUpdateResponseStatusReady      FunnelUpdateResponseStatus = "READY"
	FunnelUpdateResponseStatusProcessing FunnelUpdateResponseStatus = "PROCESSING"
)

type FunnelUpdateResponseReportDateRange struct {
	From string `json:"from" api:"required"`
	To   string `json:"to" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelUpdateResponseReportDateRange) RawJSON() string { return r.JSON.raw }
func (r *FunnelUpdateResponseReportDateRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelDeleteResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Any of true.
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
func (r FunnelDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *FunnelDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelDuplicateResponse struct {
	CreatedAt string `json:"createdAt" api:"required"`
	FunnelID  string `json:"funnelId" api:"required" format:"uuid"`
	Name      string `json:"name" api:"required"`
	// Versioned funnel definition with typed steps and filters.
	QueryDefinition FunnelDuplicateResponseQueryDefinition `json:"queryDefinition" api:"required"`
	// Any of "READY", "PROCESSING".
	Status          FunnelDuplicateResponseStatus          `json:"status" api:"required"`
	UpdatedAt       string                                 `json:"updatedAt" api:"required"`
	Description     string                                 `json:"description" api:"nullable"`
	ReportDateRange FunnelDuplicateResponseReportDateRange `json:"reportDateRange" api:"nullable"`
	Watched         bool                                   `json:"watched" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt       respjson.Field
		FunnelID        respjson.Field
		Name            respjson.Field
		QueryDefinition respjson.Field
		Status          respjson.Field
		UpdatedAt       respjson.Field
		Description     respjson.Field
		ReportDateRange respjson.Field
		Watched         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelDuplicateResponse) RawJSON() string { return r.JSON.raw }
func (r *FunnelDuplicateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Versioned funnel definition with typed steps and filters.
type FunnelDuplicateResponseQueryDefinition struct {
	// Any of "unique", "events", "sessions".
	Counting string `json:"counting" api:"required"`
	// Any of "session", "visitor".
	Identity string `json:"identity" api:"required"`
	// Any of "funnel".
	Kind string `json:"kind" api:"required"`
	// Any of "ordered", "any".
	Order string `json:"order" api:"required"`
	// Any of "first", "best".
	Reentry string                                       `json:"reentry" api:"required"`
	Steps   []FunnelDuplicateResponseQueryDefinitionStep `json:"steps" api:"required"`
	// Any of 1.
	Version     float64                                           `json:"version" api:"required"`
	WindowMs    int64                                             `json:"windowMs" api:"required"`
	EntryFilter FunnelDuplicateResponseQueryDefinitionEntryFilter `json:"entryFilter"`
	Filter      FunnelDuplicateResponseQueryDefinitionFilter      `json:"filter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Counting    respjson.Field
		Identity    respjson.Field
		Kind        respjson.Field
		Order       respjson.Field
		Reentry     respjson.Field
		Steps       respjson.Field
		Version     respjson.Field
		WindowMs    respjson.Field
		EntryFilter respjson.Field
		Filter      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelDuplicateResponseQueryDefinition) RawJSON() string { return r.JSON.raw }
func (r *FunnelDuplicateResponseQueryDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelDuplicateResponseQueryDefinitionStep struct {
	ID        string                                           `json:"id" api:"required" format:"uuid"`
	EventName string                                           `json:"eventName" api:"required"`
	Name      string                                           `json:"name" api:"required"`
	Filter    FunnelDuplicateResponseQueryDefinitionStepFilter `json:"filter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventName   respjson.Field
		Name        respjson.Field
		Filter      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelDuplicateResponseQueryDefinitionStep) RawJSON() string { return r.JSON.raw }
func (r *FunnelDuplicateResponseQueryDefinitionStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelDuplicateResponseQueryDefinitionStepFilter struct {
	Filter FunnelDuplicateResponseQueryDefinitionStepFilterFilterUnion `json:"filter" api:"required"`
	// Any of 1.
	Version float64 `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filter      respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelDuplicateResponseQueryDefinitionStepFilter) RawJSON() string { return r.JSON.raw }
func (r *FunnelDuplicateResponseQueryDefinitionStepFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelDuplicateResponseQueryDefinitionStepFilterFilterUnion contains all
// possible properties and values from
// [FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject],
// [FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject2],
// [FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FunnelDuplicateResponseQueryDefinitionStepFilterFilterUnion struct {
	// This field is from variant
	// [FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant
	// [FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant
	// [FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant
	// [FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant
	// [FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant
	// [FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant
	// [FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3].
	Value FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3ValueUnion `json:"value"`
	JSON  struct {
		Children respjson.Field
		Kind     respjson.Field
		Child    respjson.Field
		Operator respjson.Field
		Property respjson.Field
		Path     respjson.Field
		Type     respjson.Field
		Value    respjson.Field
		raw      string
	} `json:"-"`
}

func (u FunnelDuplicateResponseQueryDefinitionStepFilterFilterUnion) AsFunnelDuplicateResponseQueryDefinitionStepFilterFilterObject() (v FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelDuplicateResponseQueryDefinitionStepFilterFilterUnion) AsFunnelDuplicateResponseQueryDefinitionStepFilterFilterObject2() (v FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelDuplicateResponseQueryDefinitionStepFilterFilterUnion) AsFunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3() (v FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelDuplicateResponseQueryDefinitionStepFilterFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelDuplicateResponseQueryDefinitionStepFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject struct {
	Children []any `json:"children" api:"required"`
	// Any of "and", "or".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Children    respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject) RawJSON() string {
	return r.JSON.raw
}
func (r *FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject2 struct {
	Child any `json:"child" api:"required"`
	// Any of "not".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Child       respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject2) RawJSON() string {
	return r.JSON.raw
}
func (r *FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path"`
	// Any of "string", "number", "boolean".
	Type  string                                                                  `json:"type"`
	Value FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3ValueUnion `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		Operator    respjson.Field
		Property    respjson.Field
		Path        respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3) RawJSON() string {
	return r.JSON.raw
}
func (r *FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3ValueUnion contains
// all possible properties and values from [string], [float64], [bool],
// [[]FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfFunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3ValueArray]
type FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion]
	// instead of an object.
	OfFunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3ValueArray []FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                                      struct {
		OfString                                                                  respjson.Field
		OfFloat                                                                   respjson.Field
		OfBool                                                                    respjson.Field
		OfFunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3ValueArray respjson.Field
		raw                                                                       string
	} `json:"-"`
}

func (u FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3ValueUnion) AsFunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3ValueArray() (v []FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelDuplicateResponseQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelDuplicateResponseQueryDefinitionEntryFilter struct {
	Filter FunnelDuplicateResponseQueryDefinitionEntryFilterFilterUnion `json:"filter" api:"required"`
	// Any of 1.
	Version float64 `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filter      respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelDuplicateResponseQueryDefinitionEntryFilter) RawJSON() string { return r.JSON.raw }
func (r *FunnelDuplicateResponseQueryDefinitionEntryFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelDuplicateResponseQueryDefinitionEntryFilterFilterUnion contains all
// possible properties and values from
// [FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject],
// [FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject2],
// [FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FunnelDuplicateResponseQueryDefinitionEntryFilterFilterUnion struct {
	// This field is from variant
	// [FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant
	// [FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant
	// [FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant
	// [FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant
	// [FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant
	// [FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant
	// [FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3].
	Value FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3ValueUnion `json:"value"`
	JSON  struct {
		Children respjson.Field
		Kind     respjson.Field
		Child    respjson.Field
		Operator respjson.Field
		Property respjson.Field
		Path     respjson.Field
		Type     respjson.Field
		Value    respjson.Field
		raw      string
	} `json:"-"`
}

func (u FunnelDuplicateResponseQueryDefinitionEntryFilterFilterUnion) AsFunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject() (v FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelDuplicateResponseQueryDefinitionEntryFilterFilterUnion) AsFunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject2() (v FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelDuplicateResponseQueryDefinitionEntryFilterFilterUnion) AsFunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3() (v FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelDuplicateResponseQueryDefinitionEntryFilterFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelDuplicateResponseQueryDefinitionEntryFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject struct {
	Children []any `json:"children" api:"required"`
	// Any of "and", "or".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Children    respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject) RawJSON() string {
	return r.JSON.raw
}
func (r *FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject2 struct {
	Child any `json:"child" api:"required"`
	// Any of "not".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Child       respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject2) RawJSON() string {
	return r.JSON.raw
}
func (r *FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path"`
	// Any of "string", "number", "boolean".
	Type  string                                                                   `json:"type"`
	Value FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3ValueUnion `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		Operator    respjson.Field
		Property    respjson.Field
		Path        respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3) RawJSON() string {
	return r.JSON.raw
}
func (r *FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3ValueUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfFunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3ValueArray]
type FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion]
	// instead of an object.
	OfFunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3ValueArray []FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                                       struct {
		OfString                                                                   respjson.Field
		OfFloat                                                                    respjson.Field
		OfBool                                                                     respjson.Field
		OfFunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3ValueArray respjson.Field
		raw                                                                        string
	} `json:"-"`
}

func (u FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3ValueUnion) AsFunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3ValueArray() (v []FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelDuplicateResponseQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelDuplicateResponseQueryDefinitionFilter struct {
	Filter FunnelDuplicateResponseQueryDefinitionFilterFilterUnion `json:"filter" api:"required"`
	// Any of 1.
	Version float64 `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filter      respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelDuplicateResponseQueryDefinitionFilter) RawJSON() string { return r.JSON.raw }
func (r *FunnelDuplicateResponseQueryDefinitionFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelDuplicateResponseQueryDefinitionFilterFilterUnion contains all possible
// properties and values from
// [FunnelDuplicateResponseQueryDefinitionFilterFilterObject],
// [FunnelDuplicateResponseQueryDefinitionFilterFilterObject2],
// [FunnelDuplicateResponseQueryDefinitionFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FunnelDuplicateResponseQueryDefinitionFilterFilterUnion struct {
	// This field is from variant
	// [FunnelDuplicateResponseQueryDefinitionFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant
	// [FunnelDuplicateResponseQueryDefinitionFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant
	// [FunnelDuplicateResponseQueryDefinitionFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant
	// [FunnelDuplicateResponseQueryDefinitionFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant
	// [FunnelDuplicateResponseQueryDefinitionFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant
	// [FunnelDuplicateResponseQueryDefinitionFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant
	// [FunnelDuplicateResponseQueryDefinitionFilterFilterObject3].
	Value FunnelDuplicateResponseQueryDefinitionFilterFilterObject3ValueUnion `json:"value"`
	JSON  struct {
		Children respjson.Field
		Kind     respjson.Field
		Child    respjson.Field
		Operator respjson.Field
		Property respjson.Field
		Path     respjson.Field
		Type     respjson.Field
		Value    respjson.Field
		raw      string
	} `json:"-"`
}

func (u FunnelDuplicateResponseQueryDefinitionFilterFilterUnion) AsFunnelDuplicateResponseQueryDefinitionFilterFilterObject() (v FunnelDuplicateResponseQueryDefinitionFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelDuplicateResponseQueryDefinitionFilterFilterUnion) AsFunnelDuplicateResponseQueryDefinitionFilterFilterObject2() (v FunnelDuplicateResponseQueryDefinitionFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelDuplicateResponseQueryDefinitionFilterFilterUnion) AsFunnelDuplicateResponseQueryDefinitionFilterFilterObject3() (v FunnelDuplicateResponseQueryDefinitionFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelDuplicateResponseQueryDefinitionFilterFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *FunnelDuplicateResponseQueryDefinitionFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelDuplicateResponseQueryDefinitionFilterFilterObject struct {
	Children []any `json:"children" api:"required"`
	// Any of "and", "or".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Children    respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelDuplicateResponseQueryDefinitionFilterFilterObject) RawJSON() string { return r.JSON.raw }
func (r *FunnelDuplicateResponseQueryDefinitionFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelDuplicateResponseQueryDefinitionFilterFilterObject2 struct {
	Child any `json:"child" api:"required"`
	// Any of "not".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Child       respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelDuplicateResponseQueryDefinitionFilterFilterObject2) RawJSON() string {
	return r.JSON.raw
}
func (r *FunnelDuplicateResponseQueryDefinitionFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelDuplicateResponseQueryDefinitionFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path"`
	// Any of "string", "number", "boolean".
	Type  string                                                              `json:"type"`
	Value FunnelDuplicateResponseQueryDefinitionFilterFilterObject3ValueUnion `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		Operator    respjson.Field
		Property    respjson.Field
		Path        respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelDuplicateResponseQueryDefinitionFilterFilterObject3) RawJSON() string {
	return r.JSON.raw
}
func (r *FunnelDuplicateResponseQueryDefinitionFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelDuplicateResponseQueryDefinitionFilterFilterObject3ValueUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]FunnelDuplicateResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfFunnelDuplicateResponseQueryDefinitionFilterFilterObject3ValueArray]
type FunnelDuplicateResponseQueryDefinitionFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]FunnelDuplicateResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion]
	// instead of an object.
	OfFunnelDuplicateResponseQueryDefinitionFilterFilterObject3ValueArray []FunnelDuplicateResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                                  struct {
		OfString                                                              respjson.Field
		OfFloat                                                               respjson.Field
		OfBool                                                                respjson.Field
		OfFunnelDuplicateResponseQueryDefinitionFilterFilterObject3ValueArray respjson.Field
		raw                                                                   string
	} `json:"-"`
}

func (u FunnelDuplicateResponseQueryDefinitionFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelDuplicateResponseQueryDefinitionFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelDuplicateResponseQueryDefinitionFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelDuplicateResponseQueryDefinitionFilterFilterObject3ValueUnion) AsFunnelDuplicateResponseQueryDefinitionFilterFilterObject3ValueArray() (v []FunnelDuplicateResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelDuplicateResponseQueryDefinitionFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelDuplicateResponseQueryDefinitionFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FunnelDuplicateResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type FunnelDuplicateResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u FunnelDuplicateResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelDuplicateResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FunnelDuplicateResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FunnelDuplicateResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FunnelDuplicateResponseQueryDefinitionFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelDuplicateResponseStatus string

const (
	FunnelDuplicateResponseStatusReady      FunnelDuplicateResponseStatus = "READY"
	FunnelDuplicateResponseStatusProcessing FunnelDuplicateResponseStatus = "PROCESSING"
)

type FunnelDuplicateResponseReportDateRange struct {
	From string `json:"from" api:"required"`
	To   string `json:"to" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelDuplicateResponseReportDateRange) RawJSON() string { return r.JSON.raw }
func (r *FunnelDuplicateResponseReportDateRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelResultsResponse struct {
	DefinitionUpdatedAt time.Time `json:"definitionUpdatedAt" api:"required" format:"date-time"`
	Engine              string    `json:"engine" api:"required"`
	// Conversion rate from first step to last step as a percentage.
	OverallConversionRate float64                    `json:"overallConversionRate" api:"required"`
	Scope                 FunnelResultsResponseScope `json:"scope" api:"required"`
	// Any of 1.
	SemanticVersion float64 `json:"semanticVersion" api:"required"`
	// Known Events v3 source limitations: equal-load conflicts, corrections that move
	// across physical keys, and source completeness are unverified.
	//
	// Any of "equal-load-conflicts-unverified", "cross-key-corrections-unverified",
	// "source-completeness-unverified".
	SourceLimitations []string `json:"sourceLimitations" api:"required"`
	// Any of "events-v3-physical-key".
	SourcePolicy FunnelResultsResponseSourcePolicy `json:"sourcePolicy" api:"required"`
	// Per-step funnel analytics, ordered by step number.
	Steps []FunnelResultsResponseStep `json:"steps" api:"required"`
	// Total number of visitors who entered the funnel (entered step 1).
	TotalVisitors int64 `json:"totalVisitors" api:"required"`
	// Average time from first step to last step in seconds. Null when no completions.
	OverallAvgTimeToConversion float64 `json:"overallAvgTimeToConversion" api:"nullable"`
	SourceRevision             string  `json:"sourceRevision" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DefinitionUpdatedAt        respjson.Field
		Engine                     respjson.Field
		OverallConversionRate      respjson.Field
		Scope                      respjson.Field
		SemanticVersion            respjson.Field
		SourceLimitations          respjson.Field
		SourcePolicy               respjson.Field
		Steps                      respjson.Field
		TotalVisitors              respjson.Field
		OverallAvgTimeToConversion respjson.Field
		SourceRevision             respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelResultsResponse) RawJSON() string { return r.JSON.raw }
func (r *FunnelResultsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelResultsResponseScope struct {
	EntryFrom       string `json:"entryFrom" api:"required"`
	EntryTo         string `json:"entryTo" api:"required"`
	ObservationFrom string `json:"observationFrom" api:"required"`
	ObservationTo   string `json:"observationTo" api:"required"`
	// Any of "UTC".
	Timezone            string `json:"timezone" api:"required"`
	WindowMs            int64  `json:"windowMs" api:"required"`
	ObservationComplete bool   `json:"observationComplete" api:"nullable"`
	SettledEntryTo      string `json:"settledEntryTo" api:"nullable"`
	Watermark           string `json:"watermark" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EntryFrom           respjson.Field
		EntryTo             respjson.Field
		ObservationFrom     respjson.Field
		ObservationTo       respjson.Field
		Timezone            respjson.Field
		WindowMs            respjson.Field
		ObservationComplete respjson.Field
		SettledEntryTo      respjson.Field
		Watermark           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelResultsResponseScope) RawJSON() string { return r.JSON.raw }
func (r *FunnelResultsResponseScope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelResultsResponseSourcePolicy string

const (
	FunnelResultsResponseSourcePolicyEventsV3PhysicalKey FunnelResultsResponseSourcePolicy = "events-v3-physical-key"
)

type FunnelResultsResponseStep struct {
	ConversionCount       int64    `json:"conversionCount" api:"required"`
	ConversionRate        float64  `json:"conversionRate" api:"required"`
	DropOffRate           float64  `json:"dropOffRate" api:"required"`
	DropOffSessionIDs     []string `json:"dropOffSessionIds" api:"required"`
	DropOffVisitorIDs     []string `json:"dropOffVisitorIds" api:"required"`
	OverallConversionRate float64  `json:"overallConversionRate" api:"required"`
	SessionIDs            []string `json:"sessionIds" api:"required"`
	StepNumber            int64    `json:"stepNumber" api:"required"`
	VisitorCount          int64    `json:"visitorCount" api:"required"`
	VisitorIDs            []string `json:"visitorIds" api:"required"`
	AvgTimeToNextStep     float64  `json:"avgTimeToNextStep" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConversionCount       respjson.Field
		ConversionRate        respjson.Field
		DropOffRate           respjson.Field
		DropOffSessionIDs     respjson.Field
		DropOffVisitorIDs     respjson.Field
		OverallConversionRate respjson.Field
		SessionIDs            respjson.Field
		StepNumber            respjson.Field
		VisitorCount          respjson.Field
		VisitorIDs            respjson.Field
		AvgTimeToNextStep     respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelResultsResponseStep) RawJSON() string { return r.JSON.raw }
func (r *FunnelResultsResponseStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelNewParams struct {
	Name string `json:"name" api:"required"`
	// Versioned funnel definition with typed steps and filters.
	QueryDefinition FunnelNewParamsQueryDefinition `json:"queryDefinition,omitzero" api:"required"`
	Description     param.Opt[string]              `json:"description,omitzero"`
	Watched         param.Opt[bool]                `json:"watched,omitzero"`
	paramObj
}

func (r FunnelNewParams) MarshalJSON() (data []byte, err error) {
	type shadow FunnelNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Versioned funnel definition with typed steps and filters.
//
// The properties Counting, Identity, Kind, Order, Reentry, Steps, Version,
// WindowMs are required.
type FunnelNewParamsQueryDefinition struct {
	// Any of "unique", "events", "sessions".
	Counting string `json:"counting,omitzero" api:"required"`
	// Any of "session", "visitor".
	Identity string `json:"identity,omitzero" api:"required"`
	// Any of "funnel".
	Kind string `json:"kind,omitzero" api:"required"`
	// Any of "ordered", "any".
	Order string `json:"order,omitzero" api:"required"`
	// Any of "first", "best".
	Reentry string                               `json:"reentry,omitzero" api:"required"`
	Steps   []FunnelNewParamsQueryDefinitionStep `json:"steps,omitzero" api:"required"`
	// Any of 1.
	Version     float64                                   `json:"version,omitzero" api:"required"`
	WindowMs    int64                                     `json:"windowMs" api:"required"`
	EntryFilter FunnelNewParamsQueryDefinitionEntryFilter `json:"entryFilter,omitzero"`
	Filter      FunnelNewParamsQueryDefinitionFilter      `json:"filter,omitzero"`
	paramObj
}

func (r FunnelNewParamsQueryDefinition) MarshalJSON() (data []byte, err error) {
	type shadow FunnelNewParamsQueryDefinition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelNewParamsQueryDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelNewParamsQueryDefinition](
		"counting", "unique", "events", "sessions",
	)
	apijson.RegisterFieldValidator[FunnelNewParamsQueryDefinition](
		"identity", "session", "visitor",
	)
	apijson.RegisterFieldValidator[FunnelNewParamsQueryDefinition](
		"kind", "funnel",
	)
	apijson.RegisterFieldValidator[FunnelNewParamsQueryDefinition](
		"order", "ordered", "any",
	)
	apijson.RegisterFieldValidator[FunnelNewParamsQueryDefinition](
		"reentry", "first", "best",
	)
	apijson.RegisterFieldValidator[FunnelNewParamsQueryDefinition](
		"version", 1,
	)
}

// The properties ID, EventName, Name are required.
type FunnelNewParamsQueryDefinitionStep struct {
	ID        string                                   `json:"id" api:"required" format:"uuid"`
	EventName string                                   `json:"eventName" api:"required"`
	Name      string                                   `json:"name" api:"required"`
	Filter    FunnelNewParamsQueryDefinitionStepFilter `json:"filter,omitzero"`
	paramObj
}

func (r FunnelNewParamsQueryDefinitionStep) MarshalJSON() (data []byte, err error) {
	type shadow FunnelNewParamsQueryDefinitionStep
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelNewParamsQueryDefinitionStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Filter, Version are required.
type FunnelNewParamsQueryDefinitionStepFilter struct {
	Filter FunnelNewParamsQueryDefinitionStepFilterFilterUnion `json:"filter,omitzero" api:"required"`
	// Any of 1.
	Version float64 `json:"version,omitzero" api:"required"`
	paramObj
}

func (r FunnelNewParamsQueryDefinitionStepFilter) MarshalJSON() (data []byte, err error) {
	type shadow FunnelNewParamsQueryDefinitionStepFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelNewParamsQueryDefinitionStepFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelNewParamsQueryDefinitionStepFilter](
		"version", 1,
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FunnelNewParamsQueryDefinitionStepFilterFilterUnion struct {
	OfFunnelNewsQueryDefinitionStepFilterFilterObject  *FunnelNewParamsQueryDefinitionStepFilterFilterObject  `json:",omitzero,inline"`
	OfFunnelNewsQueryDefinitionStepFilterFilterObject2 *FunnelNewParamsQueryDefinitionStepFilterFilterObject2 `json:",omitzero,inline"`
	OfFunnelNewsQueryDefinitionStepFilterFilterObject3 *FunnelNewParamsQueryDefinitionStepFilterFilterObject3 `json:",omitzero,inline"`
	paramUnion
}

func (u FunnelNewParamsQueryDefinitionStepFilterFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFunnelNewsQueryDefinitionStepFilterFilterObject, u.OfFunnelNewsQueryDefinitionStepFilterFilterObject2, u.OfFunnelNewsQueryDefinitionStepFilterFilterObject3)
}
func (u *FunnelNewParamsQueryDefinitionStepFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Children, Kind are required.
type FunnelNewParamsQueryDefinitionStepFilterFilterObject struct {
	Children []any `json:"children,omitzero" api:"required"`
	// Any of "and", "or".
	Kind string `json:"kind,omitzero" api:"required"`
	paramObj
}

func (r FunnelNewParamsQueryDefinitionStepFilterFilterObject) MarshalJSON() (data []byte, err error) {
	type shadow FunnelNewParamsQueryDefinitionStepFilterFilterObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelNewParamsQueryDefinitionStepFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelNewParamsQueryDefinitionStepFilterFilterObject](
		"kind", "and", "or",
	)
}

// The properties Child, Kind are required.
type FunnelNewParamsQueryDefinitionStepFilterFilterObject2 struct {
	Child any `json:"child,omitzero" api:"required"`
	// Any of "not".
	Kind string `json:"kind,omitzero" api:"required"`
	paramObj
}

func (r FunnelNewParamsQueryDefinitionStepFilterFilterObject2) MarshalJSON() (data []byte, err error) {
	type shadow FunnelNewParamsQueryDefinitionStepFilterFilterObject2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelNewParamsQueryDefinitionStepFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelNewParamsQueryDefinitionStepFilterFilterObject2](
		"kind", "not",
	)
}

// The properties Kind, Operator, Property are required.
type FunnelNewParamsQueryDefinitionStepFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind,omitzero" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator,omitzero" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path,omitzero"`
	// Any of "string", "number", "boolean".
	Type  string                                                          `json:"type,omitzero"`
	Value FunnelNewParamsQueryDefinitionStepFilterFilterObject3ValueUnion `json:"value,omitzero"`
	paramObj
}

func (r FunnelNewParamsQueryDefinitionStepFilterFilterObject3) MarshalJSON() (data []byte, err error) {
	type shadow FunnelNewParamsQueryDefinitionStepFilterFilterObject3
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelNewParamsQueryDefinitionStepFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelNewParamsQueryDefinitionStepFilterFilterObject3](
		"kind", "predicate",
	)
	apijson.RegisterFieldValidator[FunnelNewParamsQueryDefinitionStepFilterFilterObject3](
		"operator", "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains", "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in", "exists", "missing", "is_null",
	)
	apijson.RegisterFieldValidator[FunnelNewParamsQueryDefinitionStepFilterFilterObject3](
		"type", "string", "number", "boolean",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FunnelNewParamsQueryDefinitionStepFilterFilterObject3ValueUnion struct {
	OfString                                                     param.Opt[string]                                                          `json:",omitzero,inline"`
	OfFloat                                                      param.Opt[float64]                                                         `json:",omitzero,inline"`
	OfBool                                                       param.Opt[bool]                                                            `json:",omitzero,inline"`
	OfFunnelNewsQueryDefinitionStepFilterFilterObject3ValueArray []FunnelNewParamsQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u FunnelNewParamsQueryDefinitionStepFilterFilterObject3ValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool, u.OfFunnelNewsQueryDefinitionStepFilterFilterObject3ValueArray)
}
func (u *FunnelNewParamsQueryDefinitionStepFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FunnelNewParamsQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u FunnelNewParamsQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *FunnelNewParamsQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Filter, Version are required.
type FunnelNewParamsQueryDefinitionEntryFilter struct {
	Filter FunnelNewParamsQueryDefinitionEntryFilterFilterUnion `json:"filter,omitzero" api:"required"`
	// Any of 1.
	Version float64 `json:"version,omitzero" api:"required"`
	paramObj
}

func (r FunnelNewParamsQueryDefinitionEntryFilter) MarshalJSON() (data []byte, err error) {
	type shadow FunnelNewParamsQueryDefinitionEntryFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelNewParamsQueryDefinitionEntryFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelNewParamsQueryDefinitionEntryFilter](
		"version", 1,
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FunnelNewParamsQueryDefinitionEntryFilterFilterUnion struct {
	OfFunnelNewsQueryDefinitionEntryFilterFilterObject  *FunnelNewParamsQueryDefinitionEntryFilterFilterObject  `json:",omitzero,inline"`
	OfFunnelNewsQueryDefinitionEntryFilterFilterObject2 *FunnelNewParamsQueryDefinitionEntryFilterFilterObject2 `json:",omitzero,inline"`
	OfFunnelNewsQueryDefinitionEntryFilterFilterObject3 *FunnelNewParamsQueryDefinitionEntryFilterFilterObject3 `json:",omitzero,inline"`
	paramUnion
}

func (u FunnelNewParamsQueryDefinitionEntryFilterFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFunnelNewsQueryDefinitionEntryFilterFilterObject, u.OfFunnelNewsQueryDefinitionEntryFilterFilterObject2, u.OfFunnelNewsQueryDefinitionEntryFilterFilterObject3)
}
func (u *FunnelNewParamsQueryDefinitionEntryFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Children, Kind are required.
type FunnelNewParamsQueryDefinitionEntryFilterFilterObject struct {
	Children []any `json:"children,omitzero" api:"required"`
	// Any of "and", "or".
	Kind string `json:"kind,omitzero" api:"required"`
	paramObj
}

func (r FunnelNewParamsQueryDefinitionEntryFilterFilterObject) MarshalJSON() (data []byte, err error) {
	type shadow FunnelNewParamsQueryDefinitionEntryFilterFilterObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelNewParamsQueryDefinitionEntryFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelNewParamsQueryDefinitionEntryFilterFilterObject](
		"kind", "and", "or",
	)
}

// The properties Child, Kind are required.
type FunnelNewParamsQueryDefinitionEntryFilterFilterObject2 struct {
	Child any `json:"child,omitzero" api:"required"`
	// Any of "not".
	Kind string `json:"kind,omitzero" api:"required"`
	paramObj
}

func (r FunnelNewParamsQueryDefinitionEntryFilterFilterObject2) MarshalJSON() (data []byte, err error) {
	type shadow FunnelNewParamsQueryDefinitionEntryFilterFilterObject2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelNewParamsQueryDefinitionEntryFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelNewParamsQueryDefinitionEntryFilterFilterObject2](
		"kind", "not",
	)
}

// The properties Kind, Operator, Property are required.
type FunnelNewParamsQueryDefinitionEntryFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind,omitzero" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator,omitzero" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path,omitzero"`
	// Any of "string", "number", "boolean".
	Type  string                                                           `json:"type,omitzero"`
	Value FunnelNewParamsQueryDefinitionEntryFilterFilterObject3ValueUnion `json:"value,omitzero"`
	paramObj
}

func (r FunnelNewParamsQueryDefinitionEntryFilterFilterObject3) MarshalJSON() (data []byte, err error) {
	type shadow FunnelNewParamsQueryDefinitionEntryFilterFilterObject3
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelNewParamsQueryDefinitionEntryFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelNewParamsQueryDefinitionEntryFilterFilterObject3](
		"kind", "predicate",
	)
	apijson.RegisterFieldValidator[FunnelNewParamsQueryDefinitionEntryFilterFilterObject3](
		"operator", "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains", "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in", "exists", "missing", "is_null",
	)
	apijson.RegisterFieldValidator[FunnelNewParamsQueryDefinitionEntryFilterFilterObject3](
		"type", "string", "number", "boolean",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FunnelNewParamsQueryDefinitionEntryFilterFilterObject3ValueUnion struct {
	OfString                                                      param.Opt[string]                                                           `json:",omitzero,inline"`
	OfFloat                                                       param.Opt[float64]                                                          `json:",omitzero,inline"`
	OfBool                                                        param.Opt[bool]                                                             `json:",omitzero,inline"`
	OfFunnelNewsQueryDefinitionEntryFilterFilterObject3ValueArray []FunnelNewParamsQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u FunnelNewParamsQueryDefinitionEntryFilterFilterObject3ValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool, u.OfFunnelNewsQueryDefinitionEntryFilterFilterObject3ValueArray)
}
func (u *FunnelNewParamsQueryDefinitionEntryFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FunnelNewParamsQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u FunnelNewParamsQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *FunnelNewParamsQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Filter, Version are required.
type FunnelNewParamsQueryDefinitionFilter struct {
	Filter FunnelNewParamsQueryDefinitionFilterFilterUnion `json:"filter,omitzero" api:"required"`
	// Any of 1.
	Version float64 `json:"version,omitzero" api:"required"`
	paramObj
}

func (r FunnelNewParamsQueryDefinitionFilter) MarshalJSON() (data []byte, err error) {
	type shadow FunnelNewParamsQueryDefinitionFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelNewParamsQueryDefinitionFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelNewParamsQueryDefinitionFilter](
		"version", 1,
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FunnelNewParamsQueryDefinitionFilterFilterUnion struct {
	OfFunnelNewsQueryDefinitionFilterFilterObject  *FunnelNewParamsQueryDefinitionFilterFilterObject  `json:",omitzero,inline"`
	OfFunnelNewsQueryDefinitionFilterFilterObject2 *FunnelNewParamsQueryDefinitionFilterFilterObject2 `json:",omitzero,inline"`
	OfFunnelNewsQueryDefinitionFilterFilterObject3 *FunnelNewParamsQueryDefinitionFilterFilterObject3 `json:",omitzero,inline"`
	paramUnion
}

func (u FunnelNewParamsQueryDefinitionFilterFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFunnelNewsQueryDefinitionFilterFilterObject, u.OfFunnelNewsQueryDefinitionFilterFilterObject2, u.OfFunnelNewsQueryDefinitionFilterFilterObject3)
}
func (u *FunnelNewParamsQueryDefinitionFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Children, Kind are required.
type FunnelNewParamsQueryDefinitionFilterFilterObject struct {
	Children []any `json:"children,omitzero" api:"required"`
	// Any of "and", "or".
	Kind string `json:"kind,omitzero" api:"required"`
	paramObj
}

func (r FunnelNewParamsQueryDefinitionFilterFilterObject) MarshalJSON() (data []byte, err error) {
	type shadow FunnelNewParamsQueryDefinitionFilterFilterObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelNewParamsQueryDefinitionFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelNewParamsQueryDefinitionFilterFilterObject](
		"kind", "and", "or",
	)
}

// The properties Child, Kind are required.
type FunnelNewParamsQueryDefinitionFilterFilterObject2 struct {
	Child any `json:"child,omitzero" api:"required"`
	// Any of "not".
	Kind string `json:"kind,omitzero" api:"required"`
	paramObj
}

func (r FunnelNewParamsQueryDefinitionFilterFilterObject2) MarshalJSON() (data []byte, err error) {
	type shadow FunnelNewParamsQueryDefinitionFilterFilterObject2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelNewParamsQueryDefinitionFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelNewParamsQueryDefinitionFilterFilterObject2](
		"kind", "not",
	)
}

// The properties Kind, Operator, Property are required.
type FunnelNewParamsQueryDefinitionFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind,omitzero" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator,omitzero" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path,omitzero"`
	// Any of "string", "number", "boolean".
	Type  string                                                      `json:"type,omitzero"`
	Value FunnelNewParamsQueryDefinitionFilterFilterObject3ValueUnion `json:"value,omitzero"`
	paramObj
}

func (r FunnelNewParamsQueryDefinitionFilterFilterObject3) MarshalJSON() (data []byte, err error) {
	type shadow FunnelNewParamsQueryDefinitionFilterFilterObject3
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelNewParamsQueryDefinitionFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelNewParamsQueryDefinitionFilterFilterObject3](
		"kind", "predicate",
	)
	apijson.RegisterFieldValidator[FunnelNewParamsQueryDefinitionFilterFilterObject3](
		"operator", "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains", "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in", "exists", "missing", "is_null",
	)
	apijson.RegisterFieldValidator[FunnelNewParamsQueryDefinitionFilterFilterObject3](
		"type", "string", "number", "boolean",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FunnelNewParamsQueryDefinitionFilterFilterObject3ValueUnion struct {
	OfString                                                 param.Opt[string]                                                      `json:",omitzero,inline"`
	OfFloat                                                  param.Opt[float64]                                                     `json:",omitzero,inline"`
	OfBool                                                   param.Opt[bool]                                                        `json:",omitzero,inline"`
	OfFunnelNewsQueryDefinitionFilterFilterObject3ValueArray []FunnelNewParamsQueryDefinitionFilterFilterObject3ValueArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u FunnelNewParamsQueryDefinitionFilterFilterObject3ValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool, u.OfFunnelNewsQueryDefinitionFilterFilterObject3ValueArray)
}
func (u *FunnelNewParamsQueryDefinitionFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FunnelNewParamsQueryDefinitionFilterFilterObject3ValueArrayItemUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u FunnelNewParamsQueryDefinitionFilterFilterObject3ValueArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *FunnelNewParamsQueryDefinitionFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type FunnelUpdateParams struct {
	Description param.Opt[string] `json:"description,omitzero"`
	Watched     param.Opt[bool]   `json:"watched,omitzero"`
	Name        param.Opt[string] `json:"name,omitzero"`
	// Versioned funnel definition with typed steps and filters.
	QueryDefinition FunnelUpdateParamsQueryDefinition `json:"queryDefinition,omitzero"`
	paramObj
}

func (r FunnelUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow FunnelUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Versioned funnel definition with typed steps and filters.
//
// The properties Counting, Identity, Kind, Order, Reentry, Steps, Version,
// WindowMs are required.
type FunnelUpdateParamsQueryDefinition struct {
	// Any of "unique", "events", "sessions".
	Counting string `json:"counting,omitzero" api:"required"`
	// Any of "session", "visitor".
	Identity string `json:"identity,omitzero" api:"required"`
	// Any of "funnel".
	Kind string `json:"kind,omitzero" api:"required"`
	// Any of "ordered", "any".
	Order string `json:"order,omitzero" api:"required"`
	// Any of "first", "best".
	Reentry string                                  `json:"reentry,omitzero" api:"required"`
	Steps   []FunnelUpdateParamsQueryDefinitionStep `json:"steps,omitzero" api:"required"`
	// Any of 1.
	Version     float64                                      `json:"version,omitzero" api:"required"`
	WindowMs    int64                                        `json:"windowMs" api:"required"`
	EntryFilter FunnelUpdateParamsQueryDefinitionEntryFilter `json:"entryFilter,omitzero"`
	Filter      FunnelUpdateParamsQueryDefinitionFilter      `json:"filter,omitzero"`
	paramObj
}

func (r FunnelUpdateParamsQueryDefinition) MarshalJSON() (data []byte, err error) {
	type shadow FunnelUpdateParamsQueryDefinition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelUpdateParamsQueryDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelUpdateParamsQueryDefinition](
		"counting", "unique", "events", "sessions",
	)
	apijson.RegisterFieldValidator[FunnelUpdateParamsQueryDefinition](
		"identity", "session", "visitor",
	)
	apijson.RegisterFieldValidator[FunnelUpdateParamsQueryDefinition](
		"kind", "funnel",
	)
	apijson.RegisterFieldValidator[FunnelUpdateParamsQueryDefinition](
		"order", "ordered", "any",
	)
	apijson.RegisterFieldValidator[FunnelUpdateParamsQueryDefinition](
		"reentry", "first", "best",
	)
	apijson.RegisterFieldValidator[FunnelUpdateParamsQueryDefinition](
		"version", 1,
	)
}

// The properties ID, EventName, Name are required.
type FunnelUpdateParamsQueryDefinitionStep struct {
	ID        string                                      `json:"id" api:"required" format:"uuid"`
	EventName string                                      `json:"eventName" api:"required"`
	Name      string                                      `json:"name" api:"required"`
	Filter    FunnelUpdateParamsQueryDefinitionStepFilter `json:"filter,omitzero"`
	paramObj
}

func (r FunnelUpdateParamsQueryDefinitionStep) MarshalJSON() (data []byte, err error) {
	type shadow FunnelUpdateParamsQueryDefinitionStep
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelUpdateParamsQueryDefinitionStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Filter, Version are required.
type FunnelUpdateParamsQueryDefinitionStepFilter struct {
	Filter FunnelUpdateParamsQueryDefinitionStepFilterFilterUnion `json:"filter,omitzero" api:"required"`
	// Any of 1.
	Version float64 `json:"version,omitzero" api:"required"`
	paramObj
}

func (r FunnelUpdateParamsQueryDefinitionStepFilter) MarshalJSON() (data []byte, err error) {
	type shadow FunnelUpdateParamsQueryDefinitionStepFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelUpdateParamsQueryDefinitionStepFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelUpdateParamsQueryDefinitionStepFilter](
		"version", 1,
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FunnelUpdateParamsQueryDefinitionStepFilterFilterUnion struct {
	OfFunnelUpdatesQueryDefinitionStepFilterFilterObject  *FunnelUpdateParamsQueryDefinitionStepFilterFilterObject  `json:",omitzero,inline"`
	OfFunnelUpdatesQueryDefinitionStepFilterFilterObject2 *FunnelUpdateParamsQueryDefinitionStepFilterFilterObject2 `json:",omitzero,inline"`
	OfFunnelUpdatesQueryDefinitionStepFilterFilterObject3 *FunnelUpdateParamsQueryDefinitionStepFilterFilterObject3 `json:",omitzero,inline"`
	paramUnion
}

func (u FunnelUpdateParamsQueryDefinitionStepFilterFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFunnelUpdatesQueryDefinitionStepFilterFilterObject, u.OfFunnelUpdatesQueryDefinitionStepFilterFilterObject2, u.OfFunnelUpdatesQueryDefinitionStepFilterFilterObject3)
}
func (u *FunnelUpdateParamsQueryDefinitionStepFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Children, Kind are required.
type FunnelUpdateParamsQueryDefinitionStepFilterFilterObject struct {
	Children []any `json:"children,omitzero" api:"required"`
	// Any of "and", "or".
	Kind string `json:"kind,omitzero" api:"required"`
	paramObj
}

func (r FunnelUpdateParamsQueryDefinitionStepFilterFilterObject) MarshalJSON() (data []byte, err error) {
	type shadow FunnelUpdateParamsQueryDefinitionStepFilterFilterObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelUpdateParamsQueryDefinitionStepFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelUpdateParamsQueryDefinitionStepFilterFilterObject](
		"kind", "and", "or",
	)
}

// The properties Child, Kind are required.
type FunnelUpdateParamsQueryDefinitionStepFilterFilterObject2 struct {
	Child any `json:"child,omitzero" api:"required"`
	// Any of "not".
	Kind string `json:"kind,omitzero" api:"required"`
	paramObj
}

func (r FunnelUpdateParamsQueryDefinitionStepFilterFilterObject2) MarshalJSON() (data []byte, err error) {
	type shadow FunnelUpdateParamsQueryDefinitionStepFilterFilterObject2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelUpdateParamsQueryDefinitionStepFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelUpdateParamsQueryDefinitionStepFilterFilterObject2](
		"kind", "not",
	)
}

// The properties Kind, Operator, Property are required.
type FunnelUpdateParamsQueryDefinitionStepFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind,omitzero" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator,omitzero" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path,omitzero"`
	// Any of "string", "number", "boolean".
	Type  string                                                             `json:"type,omitzero"`
	Value FunnelUpdateParamsQueryDefinitionStepFilterFilterObject3ValueUnion `json:"value,omitzero"`
	paramObj
}

func (r FunnelUpdateParamsQueryDefinitionStepFilterFilterObject3) MarshalJSON() (data []byte, err error) {
	type shadow FunnelUpdateParamsQueryDefinitionStepFilterFilterObject3
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelUpdateParamsQueryDefinitionStepFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelUpdateParamsQueryDefinitionStepFilterFilterObject3](
		"kind", "predicate",
	)
	apijson.RegisterFieldValidator[FunnelUpdateParamsQueryDefinitionStepFilterFilterObject3](
		"operator", "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains", "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in", "exists", "missing", "is_null",
	)
	apijson.RegisterFieldValidator[FunnelUpdateParamsQueryDefinitionStepFilterFilterObject3](
		"type", "string", "number", "boolean",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FunnelUpdateParamsQueryDefinitionStepFilterFilterObject3ValueUnion struct {
	OfString                                                        param.Opt[string]                                                             `json:",omitzero,inline"`
	OfFloat                                                         param.Opt[float64]                                                            `json:",omitzero,inline"`
	OfBool                                                          param.Opt[bool]                                                               `json:",omitzero,inline"`
	OfFunnelUpdatesQueryDefinitionStepFilterFilterObject3ValueArray []FunnelUpdateParamsQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u FunnelUpdateParamsQueryDefinitionStepFilterFilterObject3ValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool, u.OfFunnelUpdatesQueryDefinitionStepFilterFilterObject3ValueArray)
}
func (u *FunnelUpdateParamsQueryDefinitionStepFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FunnelUpdateParamsQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u FunnelUpdateParamsQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *FunnelUpdateParamsQueryDefinitionStepFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Filter, Version are required.
type FunnelUpdateParamsQueryDefinitionEntryFilter struct {
	Filter FunnelUpdateParamsQueryDefinitionEntryFilterFilterUnion `json:"filter,omitzero" api:"required"`
	// Any of 1.
	Version float64 `json:"version,omitzero" api:"required"`
	paramObj
}

func (r FunnelUpdateParamsQueryDefinitionEntryFilter) MarshalJSON() (data []byte, err error) {
	type shadow FunnelUpdateParamsQueryDefinitionEntryFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelUpdateParamsQueryDefinitionEntryFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelUpdateParamsQueryDefinitionEntryFilter](
		"version", 1,
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FunnelUpdateParamsQueryDefinitionEntryFilterFilterUnion struct {
	OfFunnelUpdatesQueryDefinitionEntryFilterFilterObject  *FunnelUpdateParamsQueryDefinitionEntryFilterFilterObject  `json:",omitzero,inline"`
	OfFunnelUpdatesQueryDefinitionEntryFilterFilterObject2 *FunnelUpdateParamsQueryDefinitionEntryFilterFilterObject2 `json:",omitzero,inline"`
	OfFunnelUpdatesQueryDefinitionEntryFilterFilterObject3 *FunnelUpdateParamsQueryDefinitionEntryFilterFilterObject3 `json:",omitzero,inline"`
	paramUnion
}

func (u FunnelUpdateParamsQueryDefinitionEntryFilterFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFunnelUpdatesQueryDefinitionEntryFilterFilterObject, u.OfFunnelUpdatesQueryDefinitionEntryFilterFilterObject2, u.OfFunnelUpdatesQueryDefinitionEntryFilterFilterObject3)
}
func (u *FunnelUpdateParamsQueryDefinitionEntryFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Children, Kind are required.
type FunnelUpdateParamsQueryDefinitionEntryFilterFilterObject struct {
	Children []any `json:"children,omitzero" api:"required"`
	// Any of "and", "or".
	Kind string `json:"kind,omitzero" api:"required"`
	paramObj
}

func (r FunnelUpdateParamsQueryDefinitionEntryFilterFilterObject) MarshalJSON() (data []byte, err error) {
	type shadow FunnelUpdateParamsQueryDefinitionEntryFilterFilterObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelUpdateParamsQueryDefinitionEntryFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelUpdateParamsQueryDefinitionEntryFilterFilterObject](
		"kind", "and", "or",
	)
}

// The properties Child, Kind are required.
type FunnelUpdateParamsQueryDefinitionEntryFilterFilterObject2 struct {
	Child any `json:"child,omitzero" api:"required"`
	// Any of "not".
	Kind string `json:"kind,omitzero" api:"required"`
	paramObj
}

func (r FunnelUpdateParamsQueryDefinitionEntryFilterFilterObject2) MarshalJSON() (data []byte, err error) {
	type shadow FunnelUpdateParamsQueryDefinitionEntryFilterFilterObject2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelUpdateParamsQueryDefinitionEntryFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelUpdateParamsQueryDefinitionEntryFilterFilterObject2](
		"kind", "not",
	)
}

// The properties Kind, Operator, Property are required.
type FunnelUpdateParamsQueryDefinitionEntryFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind,omitzero" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator,omitzero" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path,omitzero"`
	// Any of "string", "number", "boolean".
	Type  string                                                              `json:"type,omitzero"`
	Value FunnelUpdateParamsQueryDefinitionEntryFilterFilterObject3ValueUnion `json:"value,omitzero"`
	paramObj
}

func (r FunnelUpdateParamsQueryDefinitionEntryFilterFilterObject3) MarshalJSON() (data []byte, err error) {
	type shadow FunnelUpdateParamsQueryDefinitionEntryFilterFilterObject3
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelUpdateParamsQueryDefinitionEntryFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelUpdateParamsQueryDefinitionEntryFilterFilterObject3](
		"kind", "predicate",
	)
	apijson.RegisterFieldValidator[FunnelUpdateParamsQueryDefinitionEntryFilterFilterObject3](
		"operator", "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains", "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in", "exists", "missing", "is_null",
	)
	apijson.RegisterFieldValidator[FunnelUpdateParamsQueryDefinitionEntryFilterFilterObject3](
		"type", "string", "number", "boolean",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FunnelUpdateParamsQueryDefinitionEntryFilterFilterObject3ValueUnion struct {
	OfString                                                         param.Opt[string]                                                              `json:",omitzero,inline"`
	OfFloat                                                          param.Opt[float64]                                                             `json:",omitzero,inline"`
	OfBool                                                           param.Opt[bool]                                                                `json:",omitzero,inline"`
	OfFunnelUpdatesQueryDefinitionEntryFilterFilterObject3ValueArray []FunnelUpdateParamsQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u FunnelUpdateParamsQueryDefinitionEntryFilterFilterObject3ValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool, u.OfFunnelUpdatesQueryDefinitionEntryFilterFilterObject3ValueArray)
}
func (u *FunnelUpdateParamsQueryDefinitionEntryFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FunnelUpdateParamsQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u FunnelUpdateParamsQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *FunnelUpdateParamsQueryDefinitionEntryFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Filter, Version are required.
type FunnelUpdateParamsQueryDefinitionFilter struct {
	Filter FunnelUpdateParamsQueryDefinitionFilterFilterUnion `json:"filter,omitzero" api:"required"`
	// Any of 1.
	Version float64 `json:"version,omitzero" api:"required"`
	paramObj
}

func (r FunnelUpdateParamsQueryDefinitionFilter) MarshalJSON() (data []byte, err error) {
	type shadow FunnelUpdateParamsQueryDefinitionFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelUpdateParamsQueryDefinitionFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelUpdateParamsQueryDefinitionFilter](
		"version", 1,
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FunnelUpdateParamsQueryDefinitionFilterFilterUnion struct {
	OfFunnelUpdatesQueryDefinitionFilterFilterObject  *FunnelUpdateParamsQueryDefinitionFilterFilterObject  `json:",omitzero,inline"`
	OfFunnelUpdatesQueryDefinitionFilterFilterObject2 *FunnelUpdateParamsQueryDefinitionFilterFilterObject2 `json:",omitzero,inline"`
	OfFunnelUpdatesQueryDefinitionFilterFilterObject3 *FunnelUpdateParamsQueryDefinitionFilterFilterObject3 `json:",omitzero,inline"`
	paramUnion
}

func (u FunnelUpdateParamsQueryDefinitionFilterFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFunnelUpdatesQueryDefinitionFilterFilterObject, u.OfFunnelUpdatesQueryDefinitionFilterFilterObject2, u.OfFunnelUpdatesQueryDefinitionFilterFilterObject3)
}
func (u *FunnelUpdateParamsQueryDefinitionFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Children, Kind are required.
type FunnelUpdateParamsQueryDefinitionFilterFilterObject struct {
	Children []any `json:"children,omitzero" api:"required"`
	// Any of "and", "or".
	Kind string `json:"kind,omitzero" api:"required"`
	paramObj
}

func (r FunnelUpdateParamsQueryDefinitionFilterFilterObject) MarshalJSON() (data []byte, err error) {
	type shadow FunnelUpdateParamsQueryDefinitionFilterFilterObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelUpdateParamsQueryDefinitionFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelUpdateParamsQueryDefinitionFilterFilterObject](
		"kind", "and", "or",
	)
}

// The properties Child, Kind are required.
type FunnelUpdateParamsQueryDefinitionFilterFilterObject2 struct {
	Child any `json:"child,omitzero" api:"required"`
	// Any of "not".
	Kind string `json:"kind,omitzero" api:"required"`
	paramObj
}

func (r FunnelUpdateParamsQueryDefinitionFilterFilterObject2) MarshalJSON() (data []byte, err error) {
	type shadow FunnelUpdateParamsQueryDefinitionFilterFilterObject2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelUpdateParamsQueryDefinitionFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelUpdateParamsQueryDefinitionFilterFilterObject2](
		"kind", "not",
	)
}

// The properties Kind, Operator, Property are required.
type FunnelUpdateParamsQueryDefinitionFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind,omitzero" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator,omitzero" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path,omitzero"`
	// Any of "string", "number", "boolean".
	Type  string                                                         `json:"type,omitzero"`
	Value FunnelUpdateParamsQueryDefinitionFilterFilterObject3ValueUnion `json:"value,omitzero"`
	paramObj
}

func (r FunnelUpdateParamsQueryDefinitionFilterFilterObject3) MarshalJSON() (data []byte, err error) {
	type shadow FunnelUpdateParamsQueryDefinitionFilterFilterObject3
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelUpdateParamsQueryDefinitionFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelUpdateParamsQueryDefinitionFilterFilterObject3](
		"kind", "predicate",
	)
	apijson.RegisterFieldValidator[FunnelUpdateParamsQueryDefinitionFilterFilterObject3](
		"operator", "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains", "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in", "exists", "missing", "is_null",
	)
	apijson.RegisterFieldValidator[FunnelUpdateParamsQueryDefinitionFilterFilterObject3](
		"type", "string", "number", "boolean",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FunnelUpdateParamsQueryDefinitionFilterFilterObject3ValueUnion struct {
	OfString                                                    param.Opt[string]                                                         `json:",omitzero,inline"`
	OfFloat                                                     param.Opt[float64]                                                        `json:",omitzero,inline"`
	OfBool                                                      param.Opt[bool]                                                           `json:",omitzero,inline"`
	OfFunnelUpdatesQueryDefinitionFilterFilterObject3ValueArray []FunnelUpdateParamsQueryDefinitionFilterFilterObject3ValueArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u FunnelUpdateParamsQueryDefinitionFilterFilterObject3ValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool, u.OfFunnelUpdatesQueryDefinitionFilterFilterObject3ValueArray)
}
func (u *FunnelUpdateParamsQueryDefinitionFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FunnelUpdateParamsQueryDefinitionFilterFilterObject3ValueArrayItemUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u FunnelUpdateParamsQueryDefinitionFilterFilterObject3ValueArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *FunnelUpdateParamsQueryDefinitionFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type FunnelResultsParams struct {
	// Inclusive lower bound of the analysis window, as a UTC calendar day in
	// `YYYY-MM-DD` format. The window may span at most 91 days including both
	// endpoints.
	From string `query:"from" api:"required" json:"-"`
	// Inclusive upper bound of the analysis window, as a UTC calendar day in
	// `YYYY-MM-DD` format. Must be on or after `from`, and the window may span at most
	// 91 days including both endpoints.
	To string `query:"to" api:"required" json:"-"`
	// Require this saved definition revision. Returns a conflict if the funnel has
	// changed.
	ExpectedDefinitionUpdatedAt param.Opt[time.Time] `query:"expectedDefinitionUpdatedAt,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [FunnelResultsParams]'s query parameters as `url.Values`.
func (r FunnelResultsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
