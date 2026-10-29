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
	"github.com/with-ours/platform-sdk-go/v2/packages/pagination"
	"github.com/with-ours/platform-sdk-go/v2/packages/param"
	"github.com/with-ours/platform-sdk-go/v2/packages/respjson"
)

// InsightService contains methods and other services that help with interacting
// with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInsightService] method instead.
type InsightService struct {
	Options []option.RequestOption
}

// NewInsightService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInsightService(opts ...option.RequestOption) (r InsightService) {
	r = InsightService{}
	r.Options = opts
	return
}

// List saved Insights, most recently updated first. Supports cursor pagination and
// case-insensitive name/description search. Preserve unavailableCount and warnings
// when presenting partial lists. Requires API-key scope or current OAuth user
// permission: web-analytics:view, variables:find-custom
func (r *InsightService) List(ctx context.Context, query InsightListParams, opts ...option.RequestOption) (res *pagination.Cursor[InsightListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "rest/v1/insights"
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

// List saved Insights, most recently updated first. Supports cursor pagination and
// case-insensitive name/description search. Preserve unavailableCount and warnings
// when presenting partial lists. Requires API-key scope or current OAuth user
// permission: web-analytics:view, variables:find-custom
func (r *InsightService) ListAutoPaging(ctx context.Context, query InsightListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[InsightListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Save an Insight report. Use GET /rest/v1/insights/capabilities to discover
// supported features and GET /rest/v1/insights/{id}/result to compute results.
// Requires API-key scope or current OAuth user permission: web-analytics:write,
// variables:find-custom
func (r *InsightService) New(ctx context.Context, body InsightNewParams, opts ...option.RequestOption) (res *InsightNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/insights"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a saved Insight by id, including its revision, metric definition, date
// range, and chart settings. Requires API-key scope or current OAuth user
// permission: web-analytics:view, variables:find-custom
func (r *InsightService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *InsightGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/insights/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a saved Insight using expectedRevision. Omitted top-level fields remain
// unchanged; supplied definition, reportContext, and view replace the entire
// nested value. Null clears the description. Requires API-key scope or current
// OAuth user permission: web-analytics:write, variables:find-custom
func (r *InsightService) Update(ctx context.Context, id string, body InsightUpdateParams, opts ...option.RequestOption) (res *InsightUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/insights/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete a saved Insight using expectedRevision. Reload the report before retrying
// a revision conflict. Requires API-key scope or current OAuth user permission:
// web-analytics:write, variables:find-custom
func (r *InsightService) Delete(ctx context.Context, id string, body InsightDeleteParams, opts ...option.RequestOption) (res *InsightDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/insights/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// Discover supported Insight metrics, intervals, date presets, visualizations,
// comparisons, and limits before creating a report. Requires API-key scope or
// current OAuth user permission: web-analytics:view, variables:find-custom
func (r *InsightService) Capabilities(ctx context.Context, opts ...option.RequestOption) (res *InsightCapabilitiesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/insights/capabilities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Compute a saved Insight using expectedRevision with optional date, smoothing,
// cumulative, and comparison overrides. Overrides do not change the saved report.
// Requires API-key scope or current OAuth user permission: web-analytics:view,
// variables:find-custom
func (r *InsightService) Results(ctx context.Context, id string, query InsightResultsParams, opts ...option.RequestOption) (res *InsightResultsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/insights/%s/result", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Compute an unsaved Insight preview from its metric definition and date range.
// JSON-encode input in the query string; maximum 16384 bytes before URL encoding.
// Requires API-key scope or current OAuth user permission: web-analytics:view,
// variables:find-custom
func (r *InsightService) Preview(ctx context.Context, query InsightPreviewParams, opts ...option.RequestOption) (res *InsightPreviewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/insights/preview"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type InsightListResponse struct {
	ID              string    `json:"id" api:"required" format:"uuid"`
	CreatedAt       time.Time `json:"createdAt" api:"required" format:"date-time"`
	CreatedByUserID string    `json:"createdByUserId" api:"required"`
	// Version 1 Insight definition: one to ten metrics and a day, week, or month
	// interval. Sum and average require a numeric event property. Filters use the
	// analytics property catalog.
	Definition InsightListResponseDefinition `json:"definition" api:"required"`
	// Display name for the saved Insight.
	Name string `json:"name" api:"required"`
	// UTC date range using either a preset or inclusive from/to calendar dates, up to
	// 366 days.
	ReportContext   InsightListResponseReportContext `json:"reportContext" api:"required"`
	Revision        string                           `json:"revision" api:"required" format:"uuid"`
	UpdatedAt       time.Time                        `json:"updatedAt" api:"required" format:"date-time"`
	UpdatedByUserID string                           `json:"updatedByUserId" api:"required"`
	// Complete chart settings. Replaces the saved view when supplied on PATCH; omitted
	// view fields reset to their defaults.
	View InsightListResponseView `json:"view" api:"required"`
	// Optional description; null clears it.
	Description string `json:"description" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		CreatedByUserID respjson.Field
		Definition      respjson.Field
		Name            respjson.Field
		ReportContext   respjson.Field
		Revision        respjson.Field
		UpdatedAt       respjson.Field
		UpdatedByUserID respjson.Field
		View            respjson.Field
		Description     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightListResponse) RawJSON() string { return r.JSON.raw }
func (r *InsightListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Version 1 Insight definition: one to ten metrics and a day, week, or month
// interval. Sum and average require a numeric event property. Filters use the
// analytics property catalog.
type InsightListResponseDefinition struct {
	// Any of "day", "week", "month".
	Interval string `json:"interval" api:"required"`
	// Any of "insight".
	Kind    string                                `json:"kind" api:"required"`
	Metrics []InsightListResponseDefinitionMetric `json:"metrics" api:"required"`
	// Any of 1.
	Version float64                             `json:"version" api:"required"`
	Filter  InsightListResponseDefinitionFilter `json:"filter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Interval    respjson.Field
		Kind        respjson.Field
		Metrics     respjson.Field
		Version     respjson.Field
		Filter      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightListResponseDefinition) RawJSON() string { return r.JSON.raw }
func (r *InsightListResponseDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightListResponseDefinitionMetric struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Any of "event_count", "unique_visitors", "sum", "average".
	Aggregation string                                      `json:"aggregation" api:"required"`
	EventNames  []string                                    `json:"eventNames" api:"required"`
	Name        string                                      `json:"name" api:"required"`
	Filter      InsightListResponseDefinitionMetricFilter   `json:"filter"`
	Property    InsightListResponseDefinitionMetricProperty `json:"property"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Aggregation respjson.Field
		EventNames  respjson.Field
		Name        respjson.Field
		Filter      respjson.Field
		Property    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightListResponseDefinitionMetric) RawJSON() string { return r.JSON.raw }
func (r *InsightListResponseDefinitionMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightListResponseDefinitionMetricFilter struct {
	Filter InsightListResponseDefinitionMetricFilterFilterUnion `json:"filter" api:"required"`
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
func (r InsightListResponseDefinitionMetricFilter) RawJSON() string { return r.JSON.raw }
func (r *InsightListResponseDefinitionMetricFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InsightListResponseDefinitionMetricFilterFilterUnion contains all possible
// properties and values from
// [InsightListResponseDefinitionMetricFilterFilterObject],
// [InsightListResponseDefinitionMetricFilterFilterObject2],
// [InsightListResponseDefinitionMetricFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InsightListResponseDefinitionMetricFilterFilterUnion struct {
	// This field is from variant
	// [InsightListResponseDefinitionMetricFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant
	// [InsightListResponseDefinitionMetricFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant
	// [InsightListResponseDefinitionMetricFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant
	// [InsightListResponseDefinitionMetricFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant
	// [InsightListResponseDefinitionMetricFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant
	// [InsightListResponseDefinitionMetricFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant
	// [InsightListResponseDefinitionMetricFilterFilterObject3].
	Value InsightListResponseDefinitionMetricFilterFilterObject3ValueUnion `json:"value"`
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

func (u InsightListResponseDefinitionMetricFilterFilterUnion) AsInsightListResponseDefinitionMetricFilterFilterObject() (v InsightListResponseDefinitionMetricFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightListResponseDefinitionMetricFilterFilterUnion) AsInsightListResponseDefinitionMetricFilterFilterObject2() (v InsightListResponseDefinitionMetricFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightListResponseDefinitionMetricFilterFilterUnion) AsInsightListResponseDefinitionMetricFilterFilterObject3() (v InsightListResponseDefinitionMetricFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InsightListResponseDefinitionMetricFilterFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *InsightListResponseDefinitionMetricFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightListResponseDefinitionMetricFilterFilterObject struct {
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
func (r InsightListResponseDefinitionMetricFilterFilterObject) RawJSON() string { return r.JSON.raw }
func (r *InsightListResponseDefinitionMetricFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightListResponseDefinitionMetricFilterFilterObject2 struct {
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
func (r InsightListResponseDefinitionMetricFilterFilterObject2) RawJSON() string { return r.JSON.raw }
func (r *InsightListResponseDefinitionMetricFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightListResponseDefinitionMetricFilterFilterObject3 struct {
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
	Value InsightListResponseDefinitionMetricFilterFilterObject3ValueUnion `json:"value"`
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
func (r InsightListResponseDefinitionMetricFilterFilterObject3) RawJSON() string { return r.JSON.raw }
func (r *InsightListResponseDefinitionMetricFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InsightListResponseDefinitionMetricFilterFilterObject3ValueUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]InsightListResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfInsightListResponseDefinitionMetricFilterFilterObject3ValueArray]
type InsightListResponseDefinitionMetricFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]InsightListResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion]
	// instead of an object.
	OfInsightListResponseDefinitionMetricFilterFilterObject3ValueArray []InsightListResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                               struct {
		OfString                                                           respjson.Field
		OfFloat                                                            respjson.Field
		OfBool                                                             respjson.Field
		OfInsightListResponseDefinitionMetricFilterFilterObject3ValueArray respjson.Field
		raw                                                                string
	} `json:"-"`
}

func (u InsightListResponseDefinitionMetricFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightListResponseDefinitionMetricFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightListResponseDefinitionMetricFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightListResponseDefinitionMetricFilterFilterObject3ValueUnion) AsInsightListResponseDefinitionMetricFilterFilterObject3ValueArray() (v []InsightListResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InsightListResponseDefinitionMetricFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *InsightListResponseDefinitionMetricFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InsightListResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type InsightListResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion struct {
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

func (u InsightListResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightListResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightListResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InsightListResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *InsightListResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightListResponseDefinitionMetricProperty struct {
	Path []string `json:"path" api:"required"`
	// Any of "event.properties".
	Property string `json:"property" api:"required"`
	// Any of "number".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Path        respjson.Field
		Property    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightListResponseDefinitionMetricProperty) RawJSON() string { return r.JSON.raw }
func (r *InsightListResponseDefinitionMetricProperty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightListResponseDefinitionFilter struct {
	Filter InsightListResponseDefinitionFilterFilterUnion `json:"filter" api:"required"`
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
func (r InsightListResponseDefinitionFilter) RawJSON() string { return r.JSON.raw }
func (r *InsightListResponseDefinitionFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InsightListResponseDefinitionFilterFilterUnion contains all possible properties
// and values from [InsightListResponseDefinitionFilterFilterObject],
// [InsightListResponseDefinitionFilterFilterObject2],
// [InsightListResponseDefinitionFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InsightListResponseDefinitionFilterFilterUnion struct {
	// This field is from variant [InsightListResponseDefinitionFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant [InsightListResponseDefinitionFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant [InsightListResponseDefinitionFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant [InsightListResponseDefinitionFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant [InsightListResponseDefinitionFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant [InsightListResponseDefinitionFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant [InsightListResponseDefinitionFilterFilterObject3].
	Value InsightListResponseDefinitionFilterFilterObject3ValueUnion `json:"value"`
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

func (u InsightListResponseDefinitionFilterFilterUnion) AsInsightListResponseDefinitionFilterFilterObject() (v InsightListResponseDefinitionFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightListResponseDefinitionFilterFilterUnion) AsInsightListResponseDefinitionFilterFilterObject2() (v InsightListResponseDefinitionFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightListResponseDefinitionFilterFilterUnion) AsInsightListResponseDefinitionFilterFilterObject3() (v InsightListResponseDefinitionFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InsightListResponseDefinitionFilterFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *InsightListResponseDefinitionFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightListResponseDefinitionFilterFilterObject struct {
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
func (r InsightListResponseDefinitionFilterFilterObject) RawJSON() string { return r.JSON.raw }
func (r *InsightListResponseDefinitionFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightListResponseDefinitionFilterFilterObject2 struct {
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
func (r InsightListResponseDefinitionFilterFilterObject2) RawJSON() string { return r.JSON.raw }
func (r *InsightListResponseDefinitionFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightListResponseDefinitionFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path"`
	// Any of "string", "number", "boolean".
	Type  string                                                     `json:"type"`
	Value InsightListResponseDefinitionFilterFilterObject3ValueUnion `json:"value"`
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
func (r InsightListResponseDefinitionFilterFilterObject3) RawJSON() string { return r.JSON.raw }
func (r *InsightListResponseDefinitionFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InsightListResponseDefinitionFilterFilterObject3ValueUnion contains all possible
// properties and values from [string], [float64], [bool],
// [[]InsightListResponseDefinitionFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfInsightListResponseDefinitionFilterFilterObject3ValueArray]
type InsightListResponseDefinitionFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]InsightListResponseDefinitionFilterFilterObject3ValueArrayItemUnion] instead
	// of an object.
	OfInsightListResponseDefinitionFilterFilterObject3ValueArray []InsightListResponseDefinitionFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                         struct {
		OfString                                                     respjson.Field
		OfFloat                                                      respjson.Field
		OfBool                                                       respjson.Field
		OfInsightListResponseDefinitionFilterFilterObject3ValueArray respjson.Field
		raw                                                          string
	} `json:"-"`
}

func (u InsightListResponseDefinitionFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightListResponseDefinitionFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightListResponseDefinitionFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightListResponseDefinitionFilterFilterObject3ValueUnion) AsInsightListResponseDefinitionFilterFilterObject3ValueArray() (v []InsightListResponseDefinitionFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InsightListResponseDefinitionFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *InsightListResponseDefinitionFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InsightListResponseDefinitionFilterFilterObject3ValueArrayItemUnion contains all
// possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type InsightListResponseDefinitionFilterFilterObject3ValueArrayItemUnion struct {
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

func (u InsightListResponseDefinitionFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightListResponseDefinitionFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightListResponseDefinitionFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InsightListResponseDefinitionFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *InsightListResponseDefinitionFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UTC date range using either a preset or inclusive from/to calendar dates, up to
// 366 days.
type InsightListResponseReportContext struct {
	DateRange InsightListResponseReportContextDateRangeUnion `json:"dateRange" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DateRange   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightListResponseReportContext) RawJSON() string { return r.JSON.raw }
func (r *InsightListResponseReportContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InsightListResponseReportContextDateRangeUnion contains all possible properties
// and values from [InsightListResponseReportContextDateRangeObject],
// [InsightListResponseReportContextDateRangeObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InsightListResponseReportContextDateRangeUnion struct {
	// This field is from variant [InsightListResponseReportContextDateRangeObject].
	From string `json:"from"`
	// This field is from variant [InsightListResponseReportContextDateRangeObject].
	To       string `json:"to"`
	TimeZone string `json:"timeZone"`
	// This field is from variant [InsightListResponseReportContextDateRangeObject2].
	Preset string `json:"preset"`
	JSON   struct {
		From     respjson.Field
		To       respjson.Field
		TimeZone respjson.Field
		Preset   respjson.Field
		raw      string
	} `json:"-"`
}

func (u InsightListResponseReportContextDateRangeUnion) AsInsightListResponseReportContextDateRangeObject() (v InsightListResponseReportContextDateRangeObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightListResponseReportContextDateRangeUnion) AsInsightListResponseReportContextDateRangeObject2() (v InsightListResponseReportContextDateRangeObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InsightListResponseReportContextDateRangeUnion) RawJSON() string { return u.JSON.raw }

func (r *InsightListResponseReportContextDateRangeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightListResponseReportContextDateRangeObject struct {
	From string `json:"from" api:"required"`
	To   string `json:"to" api:"required"`
	// Any of "UTC".
	TimeZone string `json:"timeZone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		To          respjson.Field
		TimeZone    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightListResponseReportContextDateRangeObject) RawJSON() string { return r.JSON.raw }
func (r *InsightListResponseReportContextDateRangeObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightListResponseReportContextDateRangeObject2 struct {
	// Any of "today", "yesterday", "last_7_days", "last_30_days", "last_90_days",
	// "last_12_months", "week_to_date", "month_to_date", "year_to_date".
	Preset string `json:"preset" api:"required"`
	// Any of "UTC".
	TimeZone string `json:"timeZone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Preset      respjson.Field
		TimeZone    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightListResponseReportContextDateRangeObject2) RawJSON() string { return r.JSON.raw }
func (r *InsightListResponseReportContextDateRangeObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete chart settings. Replaces the saved view when supplied on PATCH; omitted
// view fields reset to their defaults.
type InsightListResponseView struct {
	Cumulative bool                            `json:"cumulative"`
	Series     []InsightListResponseViewSeries `json:"series"`
	// Any of "none", "rolling_average".
	Smoothing       string                           `json:"smoothing"`
	SmoothingWindow int64                            `json:"smoothingWindow"`
	Stacked         bool                             `json:"stacked"`
	TableSort       InsightListResponseViewTableSort `json:"tableSort"`
	// Any of "line", "column", "bar", "pie", "table", "metric".
	Visualization string `json:"visualization"`
	// Any of "linear", "log".
	YAxisScale string `json:"yAxisScale"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cumulative      respjson.Field
		Series          respjson.Field
		Smoothing       respjson.Field
		SmoothingWindow respjson.Field
		Stacked         respjson.Field
		TableSort       respjson.Field
		Visualization   respjson.Field
		YAxisScale      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightListResponseView) RawJSON() string { return r.JSON.raw }
func (r *InsightListResponseView) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightListResponseViewSeries struct {
	MetricID string `json:"metricId" api:"required" format:"uuid"`
	Color    string `json:"color"`
	Hidden   bool   `json:"hidden"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MetricID    respjson.Field
		Color       respjson.Field
		Hidden      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightListResponseViewSeries) RawJSON() string { return r.JSON.raw }
func (r *InsightListResponseViewSeries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightListResponseViewTableSort struct {
	// Any of "label", "range_value", "average".
	Column string `json:"column" api:"required"`
	// Any of "asc", "desc".
	Direction string `json:"direction" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Column      respjson.Field
		Direction   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightListResponseViewTableSort) RawJSON() string { return r.JSON.raw }
func (r *InsightListResponseViewTableSort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightNewResponse struct {
	ID              string    `json:"id" api:"required" format:"uuid"`
	CreatedAt       time.Time `json:"createdAt" api:"required" format:"date-time"`
	CreatedByUserID string    `json:"createdByUserId" api:"required"`
	// Version 1 Insight definition: one to ten metrics and a day, week, or month
	// interval. Sum and average require a numeric event property. Filters use the
	// analytics property catalog.
	Definition InsightNewResponseDefinition `json:"definition" api:"required"`
	// Display name for the saved Insight.
	Name string `json:"name" api:"required"`
	// UTC date range using either a preset or inclusive from/to calendar dates, up to
	// 366 days.
	ReportContext   InsightNewResponseReportContext `json:"reportContext" api:"required"`
	Revision        string                          `json:"revision" api:"required" format:"uuid"`
	UpdatedAt       time.Time                       `json:"updatedAt" api:"required" format:"date-time"`
	UpdatedByUserID string                          `json:"updatedByUserId" api:"required"`
	// Complete chart settings. Replaces the saved view when supplied on PATCH; omitted
	// view fields reset to their defaults.
	View InsightNewResponseView `json:"view" api:"required"`
	// Optional description; null clears it.
	Description string `json:"description" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		CreatedByUserID respjson.Field
		Definition      respjson.Field
		Name            respjson.Field
		ReportContext   respjson.Field
		Revision        respjson.Field
		UpdatedAt       respjson.Field
		UpdatedByUserID respjson.Field
		View            respjson.Field
		Description     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightNewResponse) RawJSON() string { return r.JSON.raw }
func (r *InsightNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Version 1 Insight definition: one to ten metrics and a day, week, or month
// interval. Sum and average require a numeric event property. Filters use the
// analytics property catalog.
type InsightNewResponseDefinition struct {
	// Any of "day", "week", "month".
	Interval string `json:"interval" api:"required"`
	// Any of "insight".
	Kind    string                               `json:"kind" api:"required"`
	Metrics []InsightNewResponseDefinitionMetric `json:"metrics" api:"required"`
	// Any of 1.
	Version float64                            `json:"version" api:"required"`
	Filter  InsightNewResponseDefinitionFilter `json:"filter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Interval    respjson.Field
		Kind        respjson.Field
		Metrics     respjson.Field
		Version     respjson.Field
		Filter      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightNewResponseDefinition) RawJSON() string { return r.JSON.raw }
func (r *InsightNewResponseDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightNewResponseDefinitionMetric struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Any of "event_count", "unique_visitors", "sum", "average".
	Aggregation string                                     `json:"aggregation" api:"required"`
	EventNames  []string                                   `json:"eventNames" api:"required"`
	Name        string                                     `json:"name" api:"required"`
	Filter      InsightNewResponseDefinitionMetricFilter   `json:"filter"`
	Property    InsightNewResponseDefinitionMetricProperty `json:"property"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Aggregation respjson.Field
		EventNames  respjson.Field
		Name        respjson.Field
		Filter      respjson.Field
		Property    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightNewResponseDefinitionMetric) RawJSON() string { return r.JSON.raw }
func (r *InsightNewResponseDefinitionMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightNewResponseDefinitionMetricFilter struct {
	Filter InsightNewResponseDefinitionMetricFilterFilterUnion `json:"filter" api:"required"`
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
func (r InsightNewResponseDefinitionMetricFilter) RawJSON() string { return r.JSON.raw }
func (r *InsightNewResponseDefinitionMetricFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InsightNewResponseDefinitionMetricFilterFilterUnion contains all possible
// properties and values from
// [InsightNewResponseDefinitionMetricFilterFilterObject],
// [InsightNewResponseDefinitionMetricFilterFilterObject2],
// [InsightNewResponseDefinitionMetricFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InsightNewResponseDefinitionMetricFilterFilterUnion struct {
	// This field is from variant
	// [InsightNewResponseDefinitionMetricFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant
	// [InsightNewResponseDefinitionMetricFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant
	// [InsightNewResponseDefinitionMetricFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant
	// [InsightNewResponseDefinitionMetricFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant
	// [InsightNewResponseDefinitionMetricFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant
	// [InsightNewResponseDefinitionMetricFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant
	// [InsightNewResponseDefinitionMetricFilterFilterObject3].
	Value InsightNewResponseDefinitionMetricFilterFilterObject3ValueUnion `json:"value"`
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

func (u InsightNewResponseDefinitionMetricFilterFilterUnion) AsInsightNewResponseDefinitionMetricFilterFilterObject() (v InsightNewResponseDefinitionMetricFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightNewResponseDefinitionMetricFilterFilterUnion) AsInsightNewResponseDefinitionMetricFilterFilterObject2() (v InsightNewResponseDefinitionMetricFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightNewResponseDefinitionMetricFilterFilterUnion) AsInsightNewResponseDefinitionMetricFilterFilterObject3() (v InsightNewResponseDefinitionMetricFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InsightNewResponseDefinitionMetricFilterFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *InsightNewResponseDefinitionMetricFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightNewResponseDefinitionMetricFilterFilterObject struct {
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
func (r InsightNewResponseDefinitionMetricFilterFilterObject) RawJSON() string { return r.JSON.raw }
func (r *InsightNewResponseDefinitionMetricFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightNewResponseDefinitionMetricFilterFilterObject2 struct {
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
func (r InsightNewResponseDefinitionMetricFilterFilterObject2) RawJSON() string { return r.JSON.raw }
func (r *InsightNewResponseDefinitionMetricFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightNewResponseDefinitionMetricFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path"`
	// Any of "string", "number", "boolean".
	Type  string                                                          `json:"type"`
	Value InsightNewResponseDefinitionMetricFilterFilterObject3ValueUnion `json:"value"`
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
func (r InsightNewResponseDefinitionMetricFilterFilterObject3) RawJSON() string { return r.JSON.raw }
func (r *InsightNewResponseDefinitionMetricFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InsightNewResponseDefinitionMetricFilterFilterObject3ValueUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]InsightNewResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfInsightNewResponseDefinitionMetricFilterFilterObject3ValueArray]
type InsightNewResponseDefinitionMetricFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]InsightNewResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion]
	// instead of an object.
	OfInsightNewResponseDefinitionMetricFilterFilterObject3ValueArray []InsightNewResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                              struct {
		OfString                                                          respjson.Field
		OfFloat                                                           respjson.Field
		OfBool                                                            respjson.Field
		OfInsightNewResponseDefinitionMetricFilterFilterObject3ValueArray respjson.Field
		raw                                                               string
	} `json:"-"`
}

func (u InsightNewResponseDefinitionMetricFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightNewResponseDefinitionMetricFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightNewResponseDefinitionMetricFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightNewResponseDefinitionMetricFilterFilterObject3ValueUnion) AsInsightNewResponseDefinitionMetricFilterFilterObject3ValueArray() (v []InsightNewResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InsightNewResponseDefinitionMetricFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *InsightNewResponseDefinitionMetricFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InsightNewResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type InsightNewResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion struct {
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

func (u InsightNewResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightNewResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightNewResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InsightNewResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *InsightNewResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightNewResponseDefinitionMetricProperty struct {
	Path []string `json:"path" api:"required"`
	// Any of "event.properties".
	Property string `json:"property" api:"required"`
	// Any of "number".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Path        respjson.Field
		Property    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightNewResponseDefinitionMetricProperty) RawJSON() string { return r.JSON.raw }
func (r *InsightNewResponseDefinitionMetricProperty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightNewResponseDefinitionFilter struct {
	Filter InsightNewResponseDefinitionFilterFilterUnion `json:"filter" api:"required"`
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
func (r InsightNewResponseDefinitionFilter) RawJSON() string { return r.JSON.raw }
func (r *InsightNewResponseDefinitionFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InsightNewResponseDefinitionFilterFilterUnion contains all possible properties
// and values from [InsightNewResponseDefinitionFilterFilterObject],
// [InsightNewResponseDefinitionFilterFilterObject2],
// [InsightNewResponseDefinitionFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InsightNewResponseDefinitionFilterFilterUnion struct {
	// This field is from variant [InsightNewResponseDefinitionFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant [InsightNewResponseDefinitionFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant [InsightNewResponseDefinitionFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant [InsightNewResponseDefinitionFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant [InsightNewResponseDefinitionFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant [InsightNewResponseDefinitionFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant [InsightNewResponseDefinitionFilterFilterObject3].
	Value InsightNewResponseDefinitionFilterFilterObject3ValueUnion `json:"value"`
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

func (u InsightNewResponseDefinitionFilterFilterUnion) AsInsightNewResponseDefinitionFilterFilterObject() (v InsightNewResponseDefinitionFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightNewResponseDefinitionFilterFilterUnion) AsInsightNewResponseDefinitionFilterFilterObject2() (v InsightNewResponseDefinitionFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightNewResponseDefinitionFilterFilterUnion) AsInsightNewResponseDefinitionFilterFilterObject3() (v InsightNewResponseDefinitionFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InsightNewResponseDefinitionFilterFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *InsightNewResponseDefinitionFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightNewResponseDefinitionFilterFilterObject struct {
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
func (r InsightNewResponseDefinitionFilterFilterObject) RawJSON() string { return r.JSON.raw }
func (r *InsightNewResponseDefinitionFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightNewResponseDefinitionFilterFilterObject2 struct {
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
func (r InsightNewResponseDefinitionFilterFilterObject2) RawJSON() string { return r.JSON.raw }
func (r *InsightNewResponseDefinitionFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightNewResponseDefinitionFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path"`
	// Any of "string", "number", "boolean".
	Type  string                                                    `json:"type"`
	Value InsightNewResponseDefinitionFilterFilterObject3ValueUnion `json:"value"`
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
func (r InsightNewResponseDefinitionFilterFilterObject3) RawJSON() string { return r.JSON.raw }
func (r *InsightNewResponseDefinitionFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InsightNewResponseDefinitionFilterFilterObject3ValueUnion contains all possible
// properties and values from [string], [float64], [bool],
// [[]InsightNewResponseDefinitionFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfInsightNewResponseDefinitionFilterFilterObject3ValueArray]
type InsightNewResponseDefinitionFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]InsightNewResponseDefinitionFilterFilterObject3ValueArrayItemUnion] instead
	// of an object.
	OfInsightNewResponseDefinitionFilterFilterObject3ValueArray []InsightNewResponseDefinitionFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                        struct {
		OfString                                                    respjson.Field
		OfFloat                                                     respjson.Field
		OfBool                                                      respjson.Field
		OfInsightNewResponseDefinitionFilterFilterObject3ValueArray respjson.Field
		raw                                                         string
	} `json:"-"`
}

func (u InsightNewResponseDefinitionFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightNewResponseDefinitionFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightNewResponseDefinitionFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightNewResponseDefinitionFilterFilterObject3ValueUnion) AsInsightNewResponseDefinitionFilterFilterObject3ValueArray() (v []InsightNewResponseDefinitionFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InsightNewResponseDefinitionFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *InsightNewResponseDefinitionFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InsightNewResponseDefinitionFilterFilterObject3ValueArrayItemUnion contains all
// possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type InsightNewResponseDefinitionFilterFilterObject3ValueArrayItemUnion struct {
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

func (u InsightNewResponseDefinitionFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightNewResponseDefinitionFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightNewResponseDefinitionFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InsightNewResponseDefinitionFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *InsightNewResponseDefinitionFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UTC date range using either a preset or inclusive from/to calendar dates, up to
// 366 days.
type InsightNewResponseReportContext struct {
	DateRange InsightNewResponseReportContextDateRangeUnion `json:"dateRange" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DateRange   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightNewResponseReportContext) RawJSON() string { return r.JSON.raw }
func (r *InsightNewResponseReportContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InsightNewResponseReportContextDateRangeUnion contains all possible properties
// and values from [InsightNewResponseReportContextDateRangeObject],
// [InsightNewResponseReportContextDateRangeObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InsightNewResponseReportContextDateRangeUnion struct {
	// This field is from variant [InsightNewResponseReportContextDateRangeObject].
	From string `json:"from"`
	// This field is from variant [InsightNewResponseReportContextDateRangeObject].
	To       string `json:"to"`
	TimeZone string `json:"timeZone"`
	// This field is from variant [InsightNewResponseReportContextDateRangeObject2].
	Preset string `json:"preset"`
	JSON   struct {
		From     respjson.Field
		To       respjson.Field
		TimeZone respjson.Field
		Preset   respjson.Field
		raw      string
	} `json:"-"`
}

func (u InsightNewResponseReportContextDateRangeUnion) AsInsightNewResponseReportContextDateRangeObject() (v InsightNewResponseReportContextDateRangeObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightNewResponseReportContextDateRangeUnion) AsInsightNewResponseReportContextDateRangeObject2() (v InsightNewResponseReportContextDateRangeObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InsightNewResponseReportContextDateRangeUnion) RawJSON() string { return u.JSON.raw }

func (r *InsightNewResponseReportContextDateRangeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightNewResponseReportContextDateRangeObject struct {
	From string `json:"from" api:"required"`
	To   string `json:"to" api:"required"`
	// Any of "UTC".
	TimeZone string `json:"timeZone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		To          respjson.Field
		TimeZone    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightNewResponseReportContextDateRangeObject) RawJSON() string { return r.JSON.raw }
func (r *InsightNewResponseReportContextDateRangeObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightNewResponseReportContextDateRangeObject2 struct {
	// Any of "today", "yesterday", "last_7_days", "last_30_days", "last_90_days",
	// "last_12_months", "week_to_date", "month_to_date", "year_to_date".
	Preset string `json:"preset" api:"required"`
	// Any of "UTC".
	TimeZone string `json:"timeZone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Preset      respjson.Field
		TimeZone    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightNewResponseReportContextDateRangeObject2) RawJSON() string { return r.JSON.raw }
func (r *InsightNewResponseReportContextDateRangeObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete chart settings. Replaces the saved view when supplied on PATCH; omitted
// view fields reset to their defaults.
type InsightNewResponseView struct {
	Cumulative bool                           `json:"cumulative"`
	Series     []InsightNewResponseViewSeries `json:"series"`
	// Any of "none", "rolling_average".
	Smoothing       string                          `json:"smoothing"`
	SmoothingWindow int64                           `json:"smoothingWindow"`
	Stacked         bool                            `json:"stacked"`
	TableSort       InsightNewResponseViewTableSort `json:"tableSort"`
	// Any of "line", "column", "bar", "pie", "table", "metric".
	Visualization string `json:"visualization"`
	// Any of "linear", "log".
	YAxisScale string `json:"yAxisScale"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cumulative      respjson.Field
		Series          respjson.Field
		Smoothing       respjson.Field
		SmoothingWindow respjson.Field
		Stacked         respjson.Field
		TableSort       respjson.Field
		Visualization   respjson.Field
		YAxisScale      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightNewResponseView) RawJSON() string { return r.JSON.raw }
func (r *InsightNewResponseView) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightNewResponseViewSeries struct {
	MetricID string `json:"metricId" api:"required" format:"uuid"`
	Color    string `json:"color"`
	Hidden   bool   `json:"hidden"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MetricID    respjson.Field
		Color       respjson.Field
		Hidden      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightNewResponseViewSeries) RawJSON() string { return r.JSON.raw }
func (r *InsightNewResponseViewSeries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightNewResponseViewTableSort struct {
	// Any of "label", "range_value", "average".
	Column string `json:"column" api:"required"`
	// Any of "asc", "desc".
	Direction string `json:"direction" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Column      respjson.Field
		Direction   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightNewResponseViewTableSort) RawJSON() string { return r.JSON.raw }
func (r *InsightNewResponseViewTableSort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightGetResponse struct {
	ID              string    `json:"id" api:"required" format:"uuid"`
	CreatedAt       time.Time `json:"createdAt" api:"required" format:"date-time"`
	CreatedByUserID string    `json:"createdByUserId" api:"required"`
	// Version 1 Insight definition: one to ten metrics and a day, week, or month
	// interval. Sum and average require a numeric event property. Filters use the
	// analytics property catalog.
	Definition InsightGetResponseDefinition `json:"definition" api:"required"`
	// Display name for the saved Insight.
	Name string `json:"name" api:"required"`
	// UTC date range using either a preset or inclusive from/to calendar dates, up to
	// 366 days.
	ReportContext   InsightGetResponseReportContext `json:"reportContext" api:"required"`
	Revision        string                          `json:"revision" api:"required" format:"uuid"`
	UpdatedAt       time.Time                       `json:"updatedAt" api:"required" format:"date-time"`
	UpdatedByUserID string                          `json:"updatedByUserId" api:"required"`
	// Complete chart settings. Replaces the saved view when supplied on PATCH; omitted
	// view fields reset to their defaults.
	View InsightGetResponseView `json:"view" api:"required"`
	// Optional description; null clears it.
	Description string `json:"description" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		CreatedByUserID respjson.Field
		Definition      respjson.Field
		Name            respjson.Field
		ReportContext   respjson.Field
		Revision        respjson.Field
		UpdatedAt       respjson.Field
		UpdatedByUserID respjson.Field
		View            respjson.Field
		Description     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightGetResponse) RawJSON() string { return r.JSON.raw }
func (r *InsightGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Version 1 Insight definition: one to ten metrics and a day, week, or month
// interval. Sum and average require a numeric event property. Filters use the
// analytics property catalog.
type InsightGetResponseDefinition struct {
	// Any of "day", "week", "month".
	Interval string `json:"interval" api:"required"`
	// Any of "insight".
	Kind    string                               `json:"kind" api:"required"`
	Metrics []InsightGetResponseDefinitionMetric `json:"metrics" api:"required"`
	// Any of 1.
	Version float64                            `json:"version" api:"required"`
	Filter  InsightGetResponseDefinitionFilter `json:"filter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Interval    respjson.Field
		Kind        respjson.Field
		Metrics     respjson.Field
		Version     respjson.Field
		Filter      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightGetResponseDefinition) RawJSON() string { return r.JSON.raw }
func (r *InsightGetResponseDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightGetResponseDefinitionMetric struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Any of "event_count", "unique_visitors", "sum", "average".
	Aggregation string                                     `json:"aggregation" api:"required"`
	EventNames  []string                                   `json:"eventNames" api:"required"`
	Name        string                                     `json:"name" api:"required"`
	Filter      InsightGetResponseDefinitionMetricFilter   `json:"filter"`
	Property    InsightGetResponseDefinitionMetricProperty `json:"property"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Aggregation respjson.Field
		EventNames  respjson.Field
		Name        respjson.Field
		Filter      respjson.Field
		Property    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightGetResponseDefinitionMetric) RawJSON() string { return r.JSON.raw }
func (r *InsightGetResponseDefinitionMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightGetResponseDefinitionMetricFilter struct {
	Filter InsightGetResponseDefinitionMetricFilterFilterUnion `json:"filter" api:"required"`
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
func (r InsightGetResponseDefinitionMetricFilter) RawJSON() string { return r.JSON.raw }
func (r *InsightGetResponseDefinitionMetricFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InsightGetResponseDefinitionMetricFilterFilterUnion contains all possible
// properties and values from
// [InsightGetResponseDefinitionMetricFilterFilterObject],
// [InsightGetResponseDefinitionMetricFilterFilterObject2],
// [InsightGetResponseDefinitionMetricFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InsightGetResponseDefinitionMetricFilterFilterUnion struct {
	// This field is from variant
	// [InsightGetResponseDefinitionMetricFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant
	// [InsightGetResponseDefinitionMetricFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant
	// [InsightGetResponseDefinitionMetricFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant
	// [InsightGetResponseDefinitionMetricFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant
	// [InsightGetResponseDefinitionMetricFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant
	// [InsightGetResponseDefinitionMetricFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant
	// [InsightGetResponseDefinitionMetricFilterFilterObject3].
	Value InsightGetResponseDefinitionMetricFilterFilterObject3ValueUnion `json:"value"`
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

func (u InsightGetResponseDefinitionMetricFilterFilterUnion) AsInsightGetResponseDefinitionMetricFilterFilterObject() (v InsightGetResponseDefinitionMetricFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightGetResponseDefinitionMetricFilterFilterUnion) AsInsightGetResponseDefinitionMetricFilterFilterObject2() (v InsightGetResponseDefinitionMetricFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightGetResponseDefinitionMetricFilterFilterUnion) AsInsightGetResponseDefinitionMetricFilterFilterObject3() (v InsightGetResponseDefinitionMetricFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InsightGetResponseDefinitionMetricFilterFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *InsightGetResponseDefinitionMetricFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightGetResponseDefinitionMetricFilterFilterObject struct {
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
func (r InsightGetResponseDefinitionMetricFilterFilterObject) RawJSON() string { return r.JSON.raw }
func (r *InsightGetResponseDefinitionMetricFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightGetResponseDefinitionMetricFilterFilterObject2 struct {
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
func (r InsightGetResponseDefinitionMetricFilterFilterObject2) RawJSON() string { return r.JSON.raw }
func (r *InsightGetResponseDefinitionMetricFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightGetResponseDefinitionMetricFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path"`
	// Any of "string", "number", "boolean".
	Type  string                                                          `json:"type"`
	Value InsightGetResponseDefinitionMetricFilterFilterObject3ValueUnion `json:"value"`
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
func (r InsightGetResponseDefinitionMetricFilterFilterObject3) RawJSON() string { return r.JSON.raw }
func (r *InsightGetResponseDefinitionMetricFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InsightGetResponseDefinitionMetricFilterFilterObject3ValueUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]InsightGetResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfInsightGetResponseDefinitionMetricFilterFilterObject3ValueArray]
type InsightGetResponseDefinitionMetricFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]InsightGetResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion]
	// instead of an object.
	OfInsightGetResponseDefinitionMetricFilterFilterObject3ValueArray []InsightGetResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                              struct {
		OfString                                                          respjson.Field
		OfFloat                                                           respjson.Field
		OfBool                                                            respjson.Field
		OfInsightGetResponseDefinitionMetricFilterFilterObject3ValueArray respjson.Field
		raw                                                               string
	} `json:"-"`
}

func (u InsightGetResponseDefinitionMetricFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightGetResponseDefinitionMetricFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightGetResponseDefinitionMetricFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightGetResponseDefinitionMetricFilterFilterObject3ValueUnion) AsInsightGetResponseDefinitionMetricFilterFilterObject3ValueArray() (v []InsightGetResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InsightGetResponseDefinitionMetricFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *InsightGetResponseDefinitionMetricFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InsightGetResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type InsightGetResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion struct {
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

func (u InsightGetResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightGetResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightGetResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InsightGetResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *InsightGetResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightGetResponseDefinitionMetricProperty struct {
	Path []string `json:"path" api:"required"`
	// Any of "event.properties".
	Property string `json:"property" api:"required"`
	// Any of "number".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Path        respjson.Field
		Property    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightGetResponseDefinitionMetricProperty) RawJSON() string { return r.JSON.raw }
func (r *InsightGetResponseDefinitionMetricProperty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightGetResponseDefinitionFilter struct {
	Filter InsightGetResponseDefinitionFilterFilterUnion `json:"filter" api:"required"`
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
func (r InsightGetResponseDefinitionFilter) RawJSON() string { return r.JSON.raw }
func (r *InsightGetResponseDefinitionFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InsightGetResponseDefinitionFilterFilterUnion contains all possible properties
// and values from [InsightGetResponseDefinitionFilterFilterObject],
// [InsightGetResponseDefinitionFilterFilterObject2],
// [InsightGetResponseDefinitionFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InsightGetResponseDefinitionFilterFilterUnion struct {
	// This field is from variant [InsightGetResponseDefinitionFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant [InsightGetResponseDefinitionFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant [InsightGetResponseDefinitionFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant [InsightGetResponseDefinitionFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant [InsightGetResponseDefinitionFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant [InsightGetResponseDefinitionFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant [InsightGetResponseDefinitionFilterFilterObject3].
	Value InsightGetResponseDefinitionFilterFilterObject3ValueUnion `json:"value"`
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

func (u InsightGetResponseDefinitionFilterFilterUnion) AsInsightGetResponseDefinitionFilterFilterObject() (v InsightGetResponseDefinitionFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightGetResponseDefinitionFilterFilterUnion) AsInsightGetResponseDefinitionFilterFilterObject2() (v InsightGetResponseDefinitionFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightGetResponseDefinitionFilterFilterUnion) AsInsightGetResponseDefinitionFilterFilterObject3() (v InsightGetResponseDefinitionFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InsightGetResponseDefinitionFilterFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *InsightGetResponseDefinitionFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightGetResponseDefinitionFilterFilterObject struct {
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
func (r InsightGetResponseDefinitionFilterFilterObject) RawJSON() string { return r.JSON.raw }
func (r *InsightGetResponseDefinitionFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightGetResponseDefinitionFilterFilterObject2 struct {
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
func (r InsightGetResponseDefinitionFilterFilterObject2) RawJSON() string { return r.JSON.raw }
func (r *InsightGetResponseDefinitionFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightGetResponseDefinitionFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path"`
	// Any of "string", "number", "boolean".
	Type  string                                                    `json:"type"`
	Value InsightGetResponseDefinitionFilterFilterObject3ValueUnion `json:"value"`
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
func (r InsightGetResponseDefinitionFilterFilterObject3) RawJSON() string { return r.JSON.raw }
func (r *InsightGetResponseDefinitionFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InsightGetResponseDefinitionFilterFilterObject3ValueUnion contains all possible
// properties and values from [string], [float64], [bool],
// [[]InsightGetResponseDefinitionFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfInsightGetResponseDefinitionFilterFilterObject3ValueArray]
type InsightGetResponseDefinitionFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]InsightGetResponseDefinitionFilterFilterObject3ValueArrayItemUnion] instead
	// of an object.
	OfInsightGetResponseDefinitionFilterFilterObject3ValueArray []InsightGetResponseDefinitionFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                        struct {
		OfString                                                    respjson.Field
		OfFloat                                                     respjson.Field
		OfBool                                                      respjson.Field
		OfInsightGetResponseDefinitionFilterFilterObject3ValueArray respjson.Field
		raw                                                         string
	} `json:"-"`
}

func (u InsightGetResponseDefinitionFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightGetResponseDefinitionFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightGetResponseDefinitionFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightGetResponseDefinitionFilterFilterObject3ValueUnion) AsInsightGetResponseDefinitionFilterFilterObject3ValueArray() (v []InsightGetResponseDefinitionFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InsightGetResponseDefinitionFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *InsightGetResponseDefinitionFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InsightGetResponseDefinitionFilterFilterObject3ValueArrayItemUnion contains all
// possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type InsightGetResponseDefinitionFilterFilterObject3ValueArrayItemUnion struct {
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

func (u InsightGetResponseDefinitionFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightGetResponseDefinitionFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightGetResponseDefinitionFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InsightGetResponseDefinitionFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *InsightGetResponseDefinitionFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UTC date range using either a preset or inclusive from/to calendar dates, up to
// 366 days.
type InsightGetResponseReportContext struct {
	DateRange InsightGetResponseReportContextDateRangeUnion `json:"dateRange" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DateRange   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightGetResponseReportContext) RawJSON() string { return r.JSON.raw }
func (r *InsightGetResponseReportContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InsightGetResponseReportContextDateRangeUnion contains all possible properties
// and values from [InsightGetResponseReportContextDateRangeObject],
// [InsightGetResponseReportContextDateRangeObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InsightGetResponseReportContextDateRangeUnion struct {
	// This field is from variant [InsightGetResponseReportContextDateRangeObject].
	From string `json:"from"`
	// This field is from variant [InsightGetResponseReportContextDateRangeObject].
	To       string `json:"to"`
	TimeZone string `json:"timeZone"`
	// This field is from variant [InsightGetResponseReportContextDateRangeObject2].
	Preset string `json:"preset"`
	JSON   struct {
		From     respjson.Field
		To       respjson.Field
		TimeZone respjson.Field
		Preset   respjson.Field
		raw      string
	} `json:"-"`
}

func (u InsightGetResponseReportContextDateRangeUnion) AsInsightGetResponseReportContextDateRangeObject() (v InsightGetResponseReportContextDateRangeObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightGetResponseReportContextDateRangeUnion) AsInsightGetResponseReportContextDateRangeObject2() (v InsightGetResponseReportContextDateRangeObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InsightGetResponseReportContextDateRangeUnion) RawJSON() string { return u.JSON.raw }

func (r *InsightGetResponseReportContextDateRangeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightGetResponseReportContextDateRangeObject struct {
	From string `json:"from" api:"required"`
	To   string `json:"to" api:"required"`
	// Any of "UTC".
	TimeZone string `json:"timeZone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		To          respjson.Field
		TimeZone    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightGetResponseReportContextDateRangeObject) RawJSON() string { return r.JSON.raw }
func (r *InsightGetResponseReportContextDateRangeObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightGetResponseReportContextDateRangeObject2 struct {
	// Any of "today", "yesterday", "last_7_days", "last_30_days", "last_90_days",
	// "last_12_months", "week_to_date", "month_to_date", "year_to_date".
	Preset string `json:"preset" api:"required"`
	// Any of "UTC".
	TimeZone string `json:"timeZone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Preset      respjson.Field
		TimeZone    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightGetResponseReportContextDateRangeObject2) RawJSON() string { return r.JSON.raw }
func (r *InsightGetResponseReportContextDateRangeObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete chart settings. Replaces the saved view when supplied on PATCH; omitted
// view fields reset to their defaults.
type InsightGetResponseView struct {
	Cumulative bool                           `json:"cumulative"`
	Series     []InsightGetResponseViewSeries `json:"series"`
	// Any of "none", "rolling_average".
	Smoothing       string                          `json:"smoothing"`
	SmoothingWindow int64                           `json:"smoothingWindow"`
	Stacked         bool                            `json:"stacked"`
	TableSort       InsightGetResponseViewTableSort `json:"tableSort"`
	// Any of "line", "column", "bar", "pie", "table", "metric".
	Visualization string `json:"visualization"`
	// Any of "linear", "log".
	YAxisScale string `json:"yAxisScale"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cumulative      respjson.Field
		Series          respjson.Field
		Smoothing       respjson.Field
		SmoothingWindow respjson.Field
		Stacked         respjson.Field
		TableSort       respjson.Field
		Visualization   respjson.Field
		YAxisScale      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightGetResponseView) RawJSON() string { return r.JSON.raw }
func (r *InsightGetResponseView) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightGetResponseViewSeries struct {
	MetricID string `json:"metricId" api:"required" format:"uuid"`
	Color    string `json:"color"`
	Hidden   bool   `json:"hidden"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MetricID    respjson.Field
		Color       respjson.Field
		Hidden      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightGetResponseViewSeries) RawJSON() string { return r.JSON.raw }
func (r *InsightGetResponseViewSeries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightGetResponseViewTableSort struct {
	// Any of "label", "range_value", "average".
	Column string `json:"column" api:"required"`
	// Any of "asc", "desc".
	Direction string `json:"direction" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Column      respjson.Field
		Direction   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightGetResponseViewTableSort) RawJSON() string { return r.JSON.raw }
func (r *InsightGetResponseViewTableSort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightUpdateResponse struct {
	ID              string    `json:"id" api:"required" format:"uuid"`
	CreatedAt       time.Time `json:"createdAt" api:"required" format:"date-time"`
	CreatedByUserID string    `json:"createdByUserId" api:"required"`
	// Version 1 Insight definition: one to ten metrics and a day, week, or month
	// interval. Sum and average require a numeric event property. Filters use the
	// analytics property catalog.
	Definition InsightUpdateResponseDefinition `json:"definition" api:"required"`
	// Display name for the saved Insight.
	Name string `json:"name" api:"required"`
	// UTC date range using either a preset or inclusive from/to calendar dates, up to
	// 366 days.
	ReportContext   InsightUpdateResponseReportContext `json:"reportContext" api:"required"`
	Revision        string                             `json:"revision" api:"required" format:"uuid"`
	UpdatedAt       time.Time                          `json:"updatedAt" api:"required" format:"date-time"`
	UpdatedByUserID string                             `json:"updatedByUserId" api:"required"`
	// Complete chart settings. Replaces the saved view when supplied on PATCH; omitted
	// view fields reset to their defaults.
	View InsightUpdateResponseView `json:"view" api:"required"`
	// Optional description; null clears it.
	Description string `json:"description" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		CreatedByUserID respjson.Field
		Definition      respjson.Field
		Name            respjson.Field
		ReportContext   respjson.Field
		Revision        respjson.Field
		UpdatedAt       respjson.Field
		UpdatedByUserID respjson.Field
		View            respjson.Field
		Description     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *InsightUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Version 1 Insight definition: one to ten metrics and a day, week, or month
// interval. Sum and average require a numeric event property. Filters use the
// analytics property catalog.
type InsightUpdateResponseDefinition struct {
	// Any of "day", "week", "month".
	Interval string `json:"interval" api:"required"`
	// Any of "insight".
	Kind    string                                  `json:"kind" api:"required"`
	Metrics []InsightUpdateResponseDefinitionMetric `json:"metrics" api:"required"`
	// Any of 1.
	Version float64                               `json:"version" api:"required"`
	Filter  InsightUpdateResponseDefinitionFilter `json:"filter"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Interval    respjson.Field
		Kind        respjson.Field
		Metrics     respjson.Field
		Version     respjson.Field
		Filter      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightUpdateResponseDefinition) RawJSON() string { return r.JSON.raw }
func (r *InsightUpdateResponseDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightUpdateResponseDefinitionMetric struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Any of "event_count", "unique_visitors", "sum", "average".
	Aggregation string                                        `json:"aggregation" api:"required"`
	EventNames  []string                                      `json:"eventNames" api:"required"`
	Name        string                                        `json:"name" api:"required"`
	Filter      InsightUpdateResponseDefinitionMetricFilter   `json:"filter"`
	Property    InsightUpdateResponseDefinitionMetricProperty `json:"property"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Aggregation respjson.Field
		EventNames  respjson.Field
		Name        respjson.Field
		Filter      respjson.Field
		Property    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightUpdateResponseDefinitionMetric) RawJSON() string { return r.JSON.raw }
func (r *InsightUpdateResponseDefinitionMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightUpdateResponseDefinitionMetricFilter struct {
	Filter InsightUpdateResponseDefinitionMetricFilterFilterUnion `json:"filter" api:"required"`
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
func (r InsightUpdateResponseDefinitionMetricFilter) RawJSON() string { return r.JSON.raw }
func (r *InsightUpdateResponseDefinitionMetricFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InsightUpdateResponseDefinitionMetricFilterFilterUnion contains all possible
// properties and values from
// [InsightUpdateResponseDefinitionMetricFilterFilterObject],
// [InsightUpdateResponseDefinitionMetricFilterFilterObject2],
// [InsightUpdateResponseDefinitionMetricFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InsightUpdateResponseDefinitionMetricFilterFilterUnion struct {
	// This field is from variant
	// [InsightUpdateResponseDefinitionMetricFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant
	// [InsightUpdateResponseDefinitionMetricFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant
	// [InsightUpdateResponseDefinitionMetricFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant
	// [InsightUpdateResponseDefinitionMetricFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant
	// [InsightUpdateResponseDefinitionMetricFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant
	// [InsightUpdateResponseDefinitionMetricFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant
	// [InsightUpdateResponseDefinitionMetricFilterFilterObject3].
	Value InsightUpdateResponseDefinitionMetricFilterFilterObject3ValueUnion `json:"value"`
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

func (u InsightUpdateResponseDefinitionMetricFilterFilterUnion) AsInsightUpdateResponseDefinitionMetricFilterFilterObject() (v InsightUpdateResponseDefinitionMetricFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightUpdateResponseDefinitionMetricFilterFilterUnion) AsInsightUpdateResponseDefinitionMetricFilterFilterObject2() (v InsightUpdateResponseDefinitionMetricFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightUpdateResponseDefinitionMetricFilterFilterUnion) AsInsightUpdateResponseDefinitionMetricFilterFilterObject3() (v InsightUpdateResponseDefinitionMetricFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InsightUpdateResponseDefinitionMetricFilterFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *InsightUpdateResponseDefinitionMetricFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightUpdateResponseDefinitionMetricFilterFilterObject struct {
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
func (r InsightUpdateResponseDefinitionMetricFilterFilterObject) RawJSON() string { return r.JSON.raw }
func (r *InsightUpdateResponseDefinitionMetricFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightUpdateResponseDefinitionMetricFilterFilterObject2 struct {
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
func (r InsightUpdateResponseDefinitionMetricFilterFilterObject2) RawJSON() string { return r.JSON.raw }
func (r *InsightUpdateResponseDefinitionMetricFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightUpdateResponseDefinitionMetricFilterFilterObject3 struct {
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
	Value InsightUpdateResponseDefinitionMetricFilterFilterObject3ValueUnion `json:"value"`
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
func (r InsightUpdateResponseDefinitionMetricFilterFilterObject3) RawJSON() string { return r.JSON.raw }
func (r *InsightUpdateResponseDefinitionMetricFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InsightUpdateResponseDefinitionMetricFilterFilterObject3ValueUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]InsightUpdateResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfInsightUpdateResponseDefinitionMetricFilterFilterObject3ValueArray]
type InsightUpdateResponseDefinitionMetricFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]InsightUpdateResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion]
	// instead of an object.
	OfInsightUpdateResponseDefinitionMetricFilterFilterObject3ValueArray []InsightUpdateResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                                 struct {
		OfString                                                             respjson.Field
		OfFloat                                                              respjson.Field
		OfBool                                                               respjson.Field
		OfInsightUpdateResponseDefinitionMetricFilterFilterObject3ValueArray respjson.Field
		raw                                                                  string
	} `json:"-"`
}

func (u InsightUpdateResponseDefinitionMetricFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightUpdateResponseDefinitionMetricFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightUpdateResponseDefinitionMetricFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightUpdateResponseDefinitionMetricFilterFilterObject3ValueUnion) AsInsightUpdateResponseDefinitionMetricFilterFilterObject3ValueArray() (v []InsightUpdateResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InsightUpdateResponseDefinitionMetricFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *InsightUpdateResponseDefinitionMetricFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InsightUpdateResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type InsightUpdateResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion struct {
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

func (u InsightUpdateResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightUpdateResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightUpdateResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InsightUpdateResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *InsightUpdateResponseDefinitionMetricFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightUpdateResponseDefinitionMetricProperty struct {
	Path []string `json:"path" api:"required"`
	// Any of "event.properties".
	Property string `json:"property" api:"required"`
	// Any of "number".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Path        respjson.Field
		Property    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightUpdateResponseDefinitionMetricProperty) RawJSON() string { return r.JSON.raw }
func (r *InsightUpdateResponseDefinitionMetricProperty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightUpdateResponseDefinitionFilter struct {
	Filter InsightUpdateResponseDefinitionFilterFilterUnion `json:"filter" api:"required"`
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
func (r InsightUpdateResponseDefinitionFilter) RawJSON() string { return r.JSON.raw }
func (r *InsightUpdateResponseDefinitionFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InsightUpdateResponseDefinitionFilterFilterUnion contains all possible
// properties and values from [InsightUpdateResponseDefinitionFilterFilterObject],
// [InsightUpdateResponseDefinitionFilterFilterObject2],
// [InsightUpdateResponseDefinitionFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InsightUpdateResponseDefinitionFilterFilterUnion struct {
	// This field is from variant [InsightUpdateResponseDefinitionFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant [InsightUpdateResponseDefinitionFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant [InsightUpdateResponseDefinitionFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant [InsightUpdateResponseDefinitionFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant [InsightUpdateResponseDefinitionFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant [InsightUpdateResponseDefinitionFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant [InsightUpdateResponseDefinitionFilterFilterObject3].
	Value InsightUpdateResponseDefinitionFilterFilterObject3ValueUnion `json:"value"`
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

func (u InsightUpdateResponseDefinitionFilterFilterUnion) AsInsightUpdateResponseDefinitionFilterFilterObject() (v InsightUpdateResponseDefinitionFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightUpdateResponseDefinitionFilterFilterUnion) AsInsightUpdateResponseDefinitionFilterFilterObject2() (v InsightUpdateResponseDefinitionFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightUpdateResponseDefinitionFilterFilterUnion) AsInsightUpdateResponseDefinitionFilterFilterObject3() (v InsightUpdateResponseDefinitionFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InsightUpdateResponseDefinitionFilterFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *InsightUpdateResponseDefinitionFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightUpdateResponseDefinitionFilterFilterObject struct {
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
func (r InsightUpdateResponseDefinitionFilterFilterObject) RawJSON() string { return r.JSON.raw }
func (r *InsightUpdateResponseDefinitionFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightUpdateResponseDefinitionFilterFilterObject2 struct {
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
func (r InsightUpdateResponseDefinitionFilterFilterObject2) RawJSON() string { return r.JSON.raw }
func (r *InsightUpdateResponseDefinitionFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightUpdateResponseDefinitionFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path"`
	// Any of "string", "number", "boolean".
	Type  string                                                       `json:"type"`
	Value InsightUpdateResponseDefinitionFilterFilterObject3ValueUnion `json:"value"`
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
func (r InsightUpdateResponseDefinitionFilterFilterObject3) RawJSON() string { return r.JSON.raw }
func (r *InsightUpdateResponseDefinitionFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InsightUpdateResponseDefinitionFilterFilterObject3ValueUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]InsightUpdateResponseDefinitionFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfInsightUpdateResponseDefinitionFilterFilterObject3ValueArray]
type InsightUpdateResponseDefinitionFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]InsightUpdateResponseDefinitionFilterFilterObject3ValueArrayItemUnion]
	// instead of an object.
	OfInsightUpdateResponseDefinitionFilterFilterObject3ValueArray []InsightUpdateResponseDefinitionFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                           struct {
		OfString                                                       respjson.Field
		OfFloat                                                        respjson.Field
		OfBool                                                         respjson.Field
		OfInsightUpdateResponseDefinitionFilterFilterObject3ValueArray respjson.Field
		raw                                                            string
	} `json:"-"`
}

func (u InsightUpdateResponseDefinitionFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightUpdateResponseDefinitionFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightUpdateResponseDefinitionFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightUpdateResponseDefinitionFilterFilterObject3ValueUnion) AsInsightUpdateResponseDefinitionFilterFilterObject3ValueArray() (v []InsightUpdateResponseDefinitionFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InsightUpdateResponseDefinitionFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *InsightUpdateResponseDefinitionFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InsightUpdateResponseDefinitionFilterFilterObject3ValueArrayItemUnion contains
// all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type InsightUpdateResponseDefinitionFilterFilterObject3ValueArrayItemUnion struct {
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

func (u InsightUpdateResponseDefinitionFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightUpdateResponseDefinitionFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightUpdateResponseDefinitionFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InsightUpdateResponseDefinitionFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *InsightUpdateResponseDefinitionFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UTC date range using either a preset or inclusive from/to calendar dates, up to
// 366 days.
type InsightUpdateResponseReportContext struct {
	DateRange InsightUpdateResponseReportContextDateRangeUnion `json:"dateRange" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DateRange   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightUpdateResponseReportContext) RawJSON() string { return r.JSON.raw }
func (r *InsightUpdateResponseReportContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InsightUpdateResponseReportContextDateRangeUnion contains all possible
// properties and values from [InsightUpdateResponseReportContextDateRangeObject],
// [InsightUpdateResponseReportContextDateRangeObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InsightUpdateResponseReportContextDateRangeUnion struct {
	// This field is from variant [InsightUpdateResponseReportContextDateRangeObject].
	From string `json:"from"`
	// This field is from variant [InsightUpdateResponseReportContextDateRangeObject].
	To       string `json:"to"`
	TimeZone string `json:"timeZone"`
	// This field is from variant [InsightUpdateResponseReportContextDateRangeObject2].
	Preset string `json:"preset"`
	JSON   struct {
		From     respjson.Field
		To       respjson.Field
		TimeZone respjson.Field
		Preset   respjson.Field
		raw      string
	} `json:"-"`
}

func (u InsightUpdateResponseReportContextDateRangeUnion) AsInsightUpdateResponseReportContextDateRangeObject() (v InsightUpdateResponseReportContextDateRangeObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InsightUpdateResponseReportContextDateRangeUnion) AsInsightUpdateResponseReportContextDateRangeObject2() (v InsightUpdateResponseReportContextDateRangeObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InsightUpdateResponseReportContextDateRangeUnion) RawJSON() string { return u.JSON.raw }

func (r *InsightUpdateResponseReportContextDateRangeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightUpdateResponseReportContextDateRangeObject struct {
	From string `json:"from" api:"required"`
	To   string `json:"to" api:"required"`
	// Any of "UTC".
	TimeZone string `json:"timeZone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		To          respjson.Field
		TimeZone    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightUpdateResponseReportContextDateRangeObject) RawJSON() string { return r.JSON.raw }
func (r *InsightUpdateResponseReportContextDateRangeObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightUpdateResponseReportContextDateRangeObject2 struct {
	// Any of "today", "yesterday", "last_7_days", "last_30_days", "last_90_days",
	// "last_12_months", "week_to_date", "month_to_date", "year_to_date".
	Preset string `json:"preset" api:"required"`
	// Any of "UTC".
	TimeZone string `json:"timeZone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Preset      respjson.Field
		TimeZone    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightUpdateResponseReportContextDateRangeObject2) RawJSON() string { return r.JSON.raw }
func (r *InsightUpdateResponseReportContextDateRangeObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Complete chart settings. Replaces the saved view when supplied on PATCH; omitted
// view fields reset to their defaults.
type InsightUpdateResponseView struct {
	Cumulative bool                              `json:"cumulative"`
	Series     []InsightUpdateResponseViewSeries `json:"series"`
	// Any of "none", "rolling_average".
	Smoothing       string                             `json:"smoothing"`
	SmoothingWindow int64                              `json:"smoothingWindow"`
	Stacked         bool                               `json:"stacked"`
	TableSort       InsightUpdateResponseViewTableSort `json:"tableSort"`
	// Any of "line", "column", "bar", "pie", "table", "metric".
	Visualization string `json:"visualization"`
	// Any of "linear", "log".
	YAxisScale string `json:"yAxisScale"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cumulative      respjson.Field
		Series          respjson.Field
		Smoothing       respjson.Field
		SmoothingWindow respjson.Field
		Stacked         respjson.Field
		TableSort       respjson.Field
		Visualization   respjson.Field
		YAxisScale      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightUpdateResponseView) RawJSON() string { return r.JSON.raw }
func (r *InsightUpdateResponseView) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightUpdateResponseViewSeries struct {
	MetricID string `json:"metricId" api:"required" format:"uuid"`
	Color    string `json:"color"`
	Hidden   bool   `json:"hidden"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MetricID    respjson.Field
		Color       respjson.Field
		Hidden      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightUpdateResponseViewSeries) RawJSON() string { return r.JSON.raw }
func (r *InsightUpdateResponseViewSeries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightUpdateResponseViewTableSort struct {
	// Any of "label", "range_value", "average".
	Column string `json:"column" api:"required"`
	// Any of "asc", "desc".
	Direction string `json:"direction" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Column      respjson.Field
		Direction   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightUpdateResponseViewTableSort) RawJSON() string { return r.JSON.raw }
func (r *InsightUpdateResponseViewTableSort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightDeleteResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Any of true.
	Deleted         bool   `json:"deleted" api:"required"`
	DeletedRevision string `json:"deletedRevision" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Deleted         respjson.Field
		DeletedRevision respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *InsightDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightCapabilitiesResponse struct {
	// Any of 1.
	DefinitionVersion float64 `json:"definitionVersion" api:"required"`
	// Any of 366.
	MaxDateRangeDays float64 `json:"maxDateRangeDays" api:"required"`
	// Any of 10.
	MaxEventNamesPerMetric float64 `json:"maxEventNamesPerMetric" api:"required"`
	// Any of 10.
	MaxMetrics float64 `json:"maxMetrics" api:"required"`
	// Any of 30.
	MaxSmoothingWindow float64 `json:"maxSmoothingWindow" api:"required"`
	// Any of 2.
	MinSmoothingWindow float64 `json:"minSmoothingWindow" api:"required"`
	// Any of "event_count", "unique_visitors", "sum", "average".
	SupportedAggregations []string `json:"supportedAggregations" api:"required"`
	// Any of "linear", "log".
	SupportedAxisScales []string `json:"supportedAxisScales" api:"required"`
	// Any of "previous_period", "previous_year", "custom".
	SupportedComparisons []string `json:"supportedComparisons" api:"required"`
	// Any of "event_count", "unique_visitors", "sum", "average".
	SupportedCumulativeAggregations []string `json:"supportedCumulativeAggregations" api:"required"`
	// Any of "today", "yesterday", "last_7_days", "last_30_days", "last_90_days",
	// "last_12_months", "week_to_date", "month_to_date", "year_to_date".
	SupportedDatePresets []string `json:"supportedDatePresets" api:"required"`
	// Any of "day", "week", "month".
	SupportedIntervals []string `json:"supportedIntervals" api:"required"`
	// Any of "none", "rolling_average".
	SupportedSmoothing []string `json:"supportedSmoothing" api:"required"`
	// Any of "line", "column", "bar", "pie", "table", "metric".
	SupportedVisualizations []string `json:"supportedVisualizations" api:"required"`
	// Any of false.
	SupportsBreakdowns bool `json:"supportsBreakdowns" api:"required"`
	// Any of true.
	SupportsComparisons bool `json:"supportsComparisons" api:"required"`
	// Any of true.
	SupportsMetricFilter bool `json:"supportsMetricFilter" api:"required"`
	// Any of false.
	SupportsNonUtcTimeZones bool `json:"supportsNonUtcTimeZones" api:"required"`
	// Any of true.
	SupportsNumericProperty bool `json:"supportsNumericProperty" api:"required"`
	// Any of true.
	SupportsReportFilter bool `json:"supportsReportFilter" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DefinitionVersion               respjson.Field
		MaxDateRangeDays                respjson.Field
		MaxEventNamesPerMetric          respjson.Field
		MaxMetrics                      respjson.Field
		MaxSmoothingWindow              respjson.Field
		MinSmoothingWindow              respjson.Field
		SupportedAggregations           respjson.Field
		SupportedAxisScales             respjson.Field
		SupportedComparisons            respjson.Field
		SupportedCumulativeAggregations respjson.Field
		SupportedDatePresets            respjson.Field
		SupportedIntervals              respjson.Field
		SupportedSmoothing              respjson.Field
		SupportedVisualizations         respjson.Field
		SupportsBreakdowns              respjson.Field
		SupportsComparisons             respjson.Field
		SupportsMetricFilter            respjson.Field
		SupportsNonUtcTimeZones         respjson.Field
		SupportsNumericProperty         respjson.Field
		SupportsReportFilter            respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightCapabilitiesResponse) RawJSON() string { return r.JSON.raw }
func (r *InsightCapabilitiesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightResultsResponse struct {
	// Any of "clickhouse".
	Engine InsightResultsResponseEngine `json:"engine" api:"required"`
	// Saved report identifier; null for previews.
	ReportID string `json:"reportId" api:"required" format:"uuid"`
	// Saved report revision; null for previews.
	Revision string                      `json:"revision" api:"required" format:"uuid"`
	Scope    InsightResultsResponseScope `json:"scope" api:"required"`
	// Any of 1.
	SemanticVersion   float64                        `json:"semanticVersion" api:"required"`
	Series            []InsightResultsResponseSeries `json:"series" api:"required"`
	SourceLimitations []string                       `json:"sourceLimitations" api:"required"`
	SourcePolicy      string                         `json:"sourcePolicy" api:"required"`
	SourceRevision    string                         `json:"sourceRevision" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Engine            respjson.Field
		ReportID          respjson.Field
		Revision          respjson.Field
		Scope             respjson.Field
		SemanticVersion   respjson.Field
		Series            respjson.Field
		SourceLimitations respjson.Field
		SourcePolicy      respjson.Field
		SourceRevision    respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightResultsResponse) RawJSON() string { return r.JSON.raw }
func (r *InsightResultsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightResultsResponseEngine string

const (
	InsightResultsResponseEngineClickhouse InsightResultsResponseEngine = "clickhouse"
)

type InsightResultsResponseScope struct {
	Cumulative bool      `json:"cumulative" api:"required"`
	ExecutedAt time.Time `json:"executedAt" api:"required" format:"date-time"`
	From       string    `json:"from" api:"required"`
	// Any of "day", "week", "month".
	Interval string `json:"interval" api:"required"`
	// Any of "unknown".
	SourceCompleteness string `json:"sourceCompleteness" api:"required"`
	// Any of "UTC".
	TimeZone   string                                `json:"timeZone" api:"required"`
	To         string                                `json:"to" api:"required"`
	Comparison InsightResultsResponseScopeComparison `json:"comparison" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cumulative         respjson.Field
		ExecutedAt         respjson.Field
		From               respjson.Field
		Interval           respjson.Field
		SourceCompleteness respjson.Field
		TimeZone           respjson.Field
		To                 respjson.Field
		Comparison         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightResultsResponseScope) RawJSON() string { return r.JSON.raw }
func (r *InsightResultsResponseScope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightResultsResponseScopeComparison struct {
	From string `json:"from" api:"required"`
	// Any of "previous_period", "previous_year", "custom".
	Kind string `json:"kind" api:"required"`
	To   string `json:"to" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		Kind        respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightResultsResponseScopeComparison) RawJSON() string { return r.JSON.raw }
func (r *InsightResultsResponseScopeComparison) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightResultsResponseSeries struct {
	// Any of "event_count", "unique_visitors", "sum", "average".
	Aggregation     string                                       `json:"aggregation" api:"required"`
	BreakdownValues []InsightResultsResponseSeriesBreakdownValue `json:"breakdownValues" api:"required"`
	// Any of "previous_period", "previous_year", "custom".
	Comparison string                              `json:"comparison" api:"required"`
	Key        string                              `json:"key" api:"required"`
	Label      string                              `json:"label" api:"required"`
	MetricID   string                              `json:"metricId" api:"required" format:"uuid"`
	Points     []InsightResultsResponseSeriesPoint `json:"points" api:"required"`
	RangeValue float64                             `json:"rangeValue" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Aggregation     respjson.Field
		BreakdownValues respjson.Field
		Comparison      respjson.Field
		Key             respjson.Field
		Label           respjson.Field
		MetricID        respjson.Field
		Points          respjson.Field
		RangeValue      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightResultsResponseSeries) RawJSON() string { return r.JSON.raw }
func (r *InsightResultsResponseSeries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightResultsResponseSeriesBreakdownValue struct {
	Property string `json:"property" api:"required"`
	Value    string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Property    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightResultsResponseSeriesBreakdownValue) RawJSON() string { return r.JSON.raw }
func (r *InsightResultsResponseSeriesBreakdownValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightResultsResponseSeriesPoint struct {
	BucketEnd   time.Time `json:"bucketEnd" api:"required" format:"date-time"`
	BucketStart time.Time `json:"bucketStart" api:"required" format:"date-time"`
	Cumulative  float64   `json:"cumulative" api:"required"`
	IsPartial   bool      `json:"isPartial" api:"required"`
	Raw         float64   `json:"raw" api:"required"`
	Smoothed    float64   `json:"smoothed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BucketEnd   respjson.Field
		BucketStart respjson.Field
		Cumulative  respjson.Field
		IsPartial   respjson.Field
		Raw         respjson.Field
		Smoothed    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightResultsResponseSeriesPoint) RawJSON() string { return r.JSON.raw }
func (r *InsightResultsResponseSeriesPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightPreviewResponse struct {
	// Any of "clickhouse".
	Engine InsightPreviewResponseEngine `json:"engine" api:"required"`
	// Saved report identifier; null for previews.
	ReportID string `json:"reportId" api:"required" format:"uuid"`
	// Saved report revision; null for previews.
	Revision string                      `json:"revision" api:"required" format:"uuid"`
	Scope    InsightPreviewResponseScope `json:"scope" api:"required"`
	// Any of 1.
	SemanticVersion   float64                        `json:"semanticVersion" api:"required"`
	Series            []InsightPreviewResponseSeries `json:"series" api:"required"`
	SourceLimitations []string                       `json:"sourceLimitations" api:"required"`
	SourcePolicy      string                         `json:"sourcePolicy" api:"required"`
	SourceRevision    string                         `json:"sourceRevision" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Engine            respjson.Field
		ReportID          respjson.Field
		Revision          respjson.Field
		Scope             respjson.Field
		SemanticVersion   respjson.Field
		Series            respjson.Field
		SourceLimitations respjson.Field
		SourcePolicy      respjson.Field
		SourceRevision    respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightPreviewResponse) RawJSON() string { return r.JSON.raw }
func (r *InsightPreviewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightPreviewResponseEngine string

const (
	InsightPreviewResponseEngineClickhouse InsightPreviewResponseEngine = "clickhouse"
)

type InsightPreviewResponseScope struct {
	Cumulative bool      `json:"cumulative" api:"required"`
	ExecutedAt time.Time `json:"executedAt" api:"required" format:"date-time"`
	From       string    `json:"from" api:"required"`
	// Any of "day", "week", "month".
	Interval string `json:"interval" api:"required"`
	// Any of "unknown".
	SourceCompleteness string `json:"sourceCompleteness" api:"required"`
	// Any of "UTC".
	TimeZone   string                                `json:"timeZone" api:"required"`
	To         string                                `json:"to" api:"required"`
	Comparison InsightPreviewResponseScopeComparison `json:"comparison" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cumulative         respjson.Field
		ExecutedAt         respjson.Field
		From               respjson.Field
		Interval           respjson.Field
		SourceCompleteness respjson.Field
		TimeZone           respjson.Field
		To                 respjson.Field
		Comparison         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightPreviewResponseScope) RawJSON() string { return r.JSON.raw }
func (r *InsightPreviewResponseScope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightPreviewResponseScopeComparison struct {
	From string `json:"from" api:"required"`
	// Any of "previous_period", "previous_year", "custom".
	Kind string `json:"kind" api:"required"`
	To   string `json:"to" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		Kind        respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightPreviewResponseScopeComparison) RawJSON() string { return r.JSON.raw }
func (r *InsightPreviewResponseScopeComparison) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightPreviewResponseSeries struct {
	// Any of "event_count", "unique_visitors", "sum", "average".
	Aggregation     string                                       `json:"aggregation" api:"required"`
	BreakdownValues []InsightPreviewResponseSeriesBreakdownValue `json:"breakdownValues" api:"required"`
	// Any of "previous_period", "previous_year", "custom".
	Comparison string                              `json:"comparison" api:"required"`
	Key        string                              `json:"key" api:"required"`
	Label      string                              `json:"label" api:"required"`
	MetricID   string                              `json:"metricId" api:"required" format:"uuid"`
	Points     []InsightPreviewResponseSeriesPoint `json:"points" api:"required"`
	RangeValue float64                             `json:"rangeValue" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Aggregation     respjson.Field
		BreakdownValues respjson.Field
		Comparison      respjson.Field
		Key             respjson.Field
		Label           respjson.Field
		MetricID        respjson.Field
		Points          respjson.Field
		RangeValue      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightPreviewResponseSeries) RawJSON() string { return r.JSON.raw }
func (r *InsightPreviewResponseSeries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightPreviewResponseSeriesBreakdownValue struct {
	Property string `json:"property" api:"required"`
	Value    string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Property    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightPreviewResponseSeriesBreakdownValue) RawJSON() string { return r.JSON.raw }
func (r *InsightPreviewResponseSeriesBreakdownValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightPreviewResponseSeriesPoint struct {
	BucketEnd   time.Time `json:"bucketEnd" api:"required" format:"date-time"`
	BucketStart time.Time `json:"bucketStart" api:"required" format:"date-time"`
	Cumulative  float64   `json:"cumulative" api:"required"`
	IsPartial   bool      `json:"isPartial" api:"required"`
	Raw         float64   `json:"raw" api:"required"`
	Smoothed    float64   `json:"smoothed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BucketEnd   respjson.Field
		BucketStart respjson.Field
		Cumulative  respjson.Field
		IsPartial   respjson.Field
		Raw         respjson.Field
		Smoothed    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsightPreviewResponseSeriesPoint) RawJSON() string { return r.JSON.raw }
func (r *InsightPreviewResponseSeriesPoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InsightListParams struct {
	// Maximum number of items to return. Defaults to 25; values below 1 are clamped to
	// 1 and values above 100 are clamped to 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Opaque pagination cursor from pagination.nextCursor in the previous response. Do
	// not decode or modify it. Malformed cursors return 400 Bad Request.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Case-insensitive search in the name and description.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InsightListParams]'s query parameters as `url.Values`.
func (r InsightListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InsightNewParams struct {
	// Version 1 Insight definition: one to ten metrics and a day, week, or month
	// interval. Sum and average require a numeric event property. Filters use the
	// analytics property catalog.
	Definition InsightNewParamsDefinition `json:"definition,omitzero" api:"required"`
	// Display name for the saved Insight.
	Name string `json:"name" api:"required"`
	// UTC date range using either a preset or inclusive from/to calendar dates, up to
	// 366 days.
	ReportContext InsightNewParamsReportContext `json:"reportContext,omitzero" api:"required"`
	// Complete chart settings. Replaces the saved view when supplied on PATCH; omitted
	// view fields reset to their defaults.
	View InsightNewParamsView `json:"view,omitzero" api:"required"`
	// Optional description; null clears it.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r InsightNewParams) MarshalJSON() (data []byte, err error) {
	type shadow InsightNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Version 1 Insight definition: one to ten metrics and a day, week, or month
// interval. Sum and average require a numeric event property. Filters use the
// analytics property catalog.
//
// The properties Interval, Kind, Metrics, Version are required.
type InsightNewParamsDefinition struct {
	// Any of "day", "week", "month".
	Interval string `json:"interval,omitzero" api:"required"`
	// Any of "insight".
	Kind    string                             `json:"kind,omitzero" api:"required"`
	Metrics []InsightNewParamsDefinitionMetric `json:"metrics,omitzero" api:"required"`
	// Any of 1.
	Version float64                          `json:"version,omitzero" api:"required"`
	Filter  InsightNewParamsDefinitionFilter `json:"filter,omitzero"`
	paramObj
}

func (r InsightNewParamsDefinition) MarshalJSON() (data []byte, err error) {
	type shadow InsightNewParamsDefinition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightNewParamsDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InsightNewParamsDefinition](
		"interval", "day", "week", "month",
	)
	apijson.RegisterFieldValidator[InsightNewParamsDefinition](
		"kind", "insight",
	)
	apijson.RegisterFieldValidator[InsightNewParamsDefinition](
		"version", 1,
	)
}

// The properties ID, Aggregation, EventNames, Name are required.
type InsightNewParamsDefinitionMetric struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Any of "event_count", "unique_visitors", "sum", "average".
	Aggregation string                                   `json:"aggregation,omitzero" api:"required"`
	EventNames  []string                                 `json:"eventNames,omitzero" api:"required"`
	Name        string                                   `json:"name" api:"required"`
	Filter      InsightNewParamsDefinitionMetricFilter   `json:"filter,omitzero"`
	Property    InsightNewParamsDefinitionMetricProperty `json:"property,omitzero"`
	paramObj
}

func (r InsightNewParamsDefinitionMetric) MarshalJSON() (data []byte, err error) {
	type shadow InsightNewParamsDefinitionMetric
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightNewParamsDefinitionMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InsightNewParamsDefinitionMetric](
		"aggregation", "event_count", "unique_visitors", "sum", "average",
	)
}

// The properties Filter, Version are required.
type InsightNewParamsDefinitionMetricFilter struct {
	Filter InsightNewParamsDefinitionMetricFilterFilterUnion `json:"filter,omitzero" api:"required"`
	// Any of 1.
	Version float64 `json:"version,omitzero" api:"required"`
	paramObj
}

func (r InsightNewParamsDefinitionMetricFilter) MarshalJSON() (data []byte, err error) {
	type shadow InsightNewParamsDefinitionMetricFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightNewParamsDefinitionMetricFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InsightNewParamsDefinitionMetricFilter](
		"version", 1,
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InsightNewParamsDefinitionMetricFilterFilterUnion struct {
	OfInsightNewsDefinitionMetricFilterFilterObject  *InsightNewParamsDefinitionMetricFilterFilterObject  `json:",omitzero,inline"`
	OfInsightNewsDefinitionMetricFilterFilterObject2 *InsightNewParamsDefinitionMetricFilterFilterObject2 `json:",omitzero,inline"`
	OfInsightNewsDefinitionMetricFilterFilterObject3 *InsightNewParamsDefinitionMetricFilterFilterObject3 `json:",omitzero,inline"`
	paramUnion
}

func (u InsightNewParamsDefinitionMetricFilterFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInsightNewsDefinitionMetricFilterFilterObject, u.OfInsightNewsDefinitionMetricFilterFilterObject2, u.OfInsightNewsDefinitionMetricFilterFilterObject3)
}
func (u *InsightNewParamsDefinitionMetricFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Children, Kind are required.
type InsightNewParamsDefinitionMetricFilterFilterObject struct {
	Children []any `json:"children,omitzero" api:"required"`
	// Any of "and", "or".
	Kind string `json:"kind,omitzero" api:"required"`
	paramObj
}

func (r InsightNewParamsDefinitionMetricFilterFilterObject) MarshalJSON() (data []byte, err error) {
	type shadow InsightNewParamsDefinitionMetricFilterFilterObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightNewParamsDefinitionMetricFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InsightNewParamsDefinitionMetricFilterFilterObject](
		"kind", "and", "or",
	)
}

// The properties Child, Kind are required.
type InsightNewParamsDefinitionMetricFilterFilterObject2 struct {
	Child any `json:"child,omitzero" api:"required"`
	// Any of "not".
	Kind string `json:"kind,omitzero" api:"required"`
	paramObj
}

func (r InsightNewParamsDefinitionMetricFilterFilterObject2) MarshalJSON() (data []byte, err error) {
	type shadow InsightNewParamsDefinitionMetricFilterFilterObject2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightNewParamsDefinitionMetricFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InsightNewParamsDefinitionMetricFilterFilterObject2](
		"kind", "not",
	)
}

// The properties Kind, Operator, Property are required.
type InsightNewParamsDefinitionMetricFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind,omitzero" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator,omitzero" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path,omitzero"`
	// Any of "string", "number", "boolean".
	Type  string                                                        `json:"type,omitzero"`
	Value InsightNewParamsDefinitionMetricFilterFilterObject3ValueUnion `json:"value,omitzero"`
	paramObj
}

func (r InsightNewParamsDefinitionMetricFilterFilterObject3) MarshalJSON() (data []byte, err error) {
	type shadow InsightNewParamsDefinitionMetricFilterFilterObject3
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightNewParamsDefinitionMetricFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InsightNewParamsDefinitionMetricFilterFilterObject3](
		"kind", "predicate",
	)
	apijson.RegisterFieldValidator[InsightNewParamsDefinitionMetricFilterFilterObject3](
		"operator", "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains", "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in", "exists", "missing", "is_null",
	)
	apijson.RegisterFieldValidator[InsightNewParamsDefinitionMetricFilterFilterObject3](
		"type", "string", "number", "boolean",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InsightNewParamsDefinitionMetricFilterFilterObject3ValueUnion struct {
	OfString                                                   param.Opt[string]                                                        `json:",omitzero,inline"`
	OfFloat                                                    param.Opt[float64]                                                       `json:",omitzero,inline"`
	OfBool                                                     param.Opt[bool]                                                          `json:",omitzero,inline"`
	OfInsightNewsDefinitionMetricFilterFilterObject3ValueArray []InsightNewParamsDefinitionMetricFilterFilterObject3ValueArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u InsightNewParamsDefinitionMetricFilterFilterObject3ValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool, u.OfInsightNewsDefinitionMetricFilterFilterObject3ValueArray)
}
func (u *InsightNewParamsDefinitionMetricFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InsightNewParamsDefinitionMetricFilterFilterObject3ValueArrayItemUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u InsightNewParamsDefinitionMetricFilterFilterObject3ValueArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *InsightNewParamsDefinitionMetricFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Path, Property are required.
type InsightNewParamsDefinitionMetricProperty struct {
	Path []string `json:"path,omitzero" api:"required"`
	// Any of "event.properties".
	Property string `json:"property,omitzero" api:"required"`
	// Any of "number".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r InsightNewParamsDefinitionMetricProperty) MarshalJSON() (data []byte, err error) {
	type shadow InsightNewParamsDefinitionMetricProperty
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightNewParamsDefinitionMetricProperty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InsightNewParamsDefinitionMetricProperty](
		"property", "event.properties",
	)
	apijson.RegisterFieldValidator[InsightNewParamsDefinitionMetricProperty](
		"type", "number",
	)
}

// The properties Filter, Version are required.
type InsightNewParamsDefinitionFilter struct {
	Filter InsightNewParamsDefinitionFilterFilterUnion `json:"filter,omitzero" api:"required"`
	// Any of 1.
	Version float64 `json:"version,omitzero" api:"required"`
	paramObj
}

func (r InsightNewParamsDefinitionFilter) MarshalJSON() (data []byte, err error) {
	type shadow InsightNewParamsDefinitionFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightNewParamsDefinitionFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InsightNewParamsDefinitionFilter](
		"version", 1,
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InsightNewParamsDefinitionFilterFilterUnion struct {
	OfInsightNewsDefinitionFilterFilterObject  *InsightNewParamsDefinitionFilterFilterObject  `json:",omitzero,inline"`
	OfInsightNewsDefinitionFilterFilterObject2 *InsightNewParamsDefinitionFilterFilterObject2 `json:",omitzero,inline"`
	OfInsightNewsDefinitionFilterFilterObject3 *InsightNewParamsDefinitionFilterFilterObject3 `json:",omitzero,inline"`
	paramUnion
}

func (u InsightNewParamsDefinitionFilterFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInsightNewsDefinitionFilterFilterObject, u.OfInsightNewsDefinitionFilterFilterObject2, u.OfInsightNewsDefinitionFilterFilterObject3)
}
func (u *InsightNewParamsDefinitionFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Children, Kind are required.
type InsightNewParamsDefinitionFilterFilterObject struct {
	Children []any `json:"children,omitzero" api:"required"`
	// Any of "and", "or".
	Kind string `json:"kind,omitzero" api:"required"`
	paramObj
}

func (r InsightNewParamsDefinitionFilterFilterObject) MarshalJSON() (data []byte, err error) {
	type shadow InsightNewParamsDefinitionFilterFilterObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightNewParamsDefinitionFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InsightNewParamsDefinitionFilterFilterObject](
		"kind", "and", "or",
	)
}

// The properties Child, Kind are required.
type InsightNewParamsDefinitionFilterFilterObject2 struct {
	Child any `json:"child,omitzero" api:"required"`
	// Any of "not".
	Kind string `json:"kind,omitzero" api:"required"`
	paramObj
}

func (r InsightNewParamsDefinitionFilterFilterObject2) MarshalJSON() (data []byte, err error) {
	type shadow InsightNewParamsDefinitionFilterFilterObject2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightNewParamsDefinitionFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InsightNewParamsDefinitionFilterFilterObject2](
		"kind", "not",
	)
}

// The properties Kind, Operator, Property are required.
type InsightNewParamsDefinitionFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind,omitzero" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator,omitzero" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path,omitzero"`
	// Any of "string", "number", "boolean".
	Type  string                                                  `json:"type,omitzero"`
	Value InsightNewParamsDefinitionFilterFilterObject3ValueUnion `json:"value,omitzero"`
	paramObj
}

func (r InsightNewParamsDefinitionFilterFilterObject3) MarshalJSON() (data []byte, err error) {
	type shadow InsightNewParamsDefinitionFilterFilterObject3
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightNewParamsDefinitionFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InsightNewParamsDefinitionFilterFilterObject3](
		"kind", "predicate",
	)
	apijson.RegisterFieldValidator[InsightNewParamsDefinitionFilterFilterObject3](
		"operator", "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains", "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in", "exists", "missing", "is_null",
	)
	apijson.RegisterFieldValidator[InsightNewParamsDefinitionFilterFilterObject3](
		"type", "string", "number", "boolean",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InsightNewParamsDefinitionFilterFilterObject3ValueUnion struct {
	OfString                                             param.Opt[string]                                                  `json:",omitzero,inline"`
	OfFloat                                              param.Opt[float64]                                                 `json:",omitzero,inline"`
	OfBool                                               param.Opt[bool]                                                    `json:",omitzero,inline"`
	OfInsightNewsDefinitionFilterFilterObject3ValueArray []InsightNewParamsDefinitionFilterFilterObject3ValueArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u InsightNewParamsDefinitionFilterFilterObject3ValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool, u.OfInsightNewsDefinitionFilterFilterObject3ValueArray)
}
func (u *InsightNewParamsDefinitionFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InsightNewParamsDefinitionFilterFilterObject3ValueArrayItemUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u InsightNewParamsDefinitionFilterFilterObject3ValueArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *InsightNewParamsDefinitionFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// UTC date range using either a preset or inclusive from/to calendar dates, up to
// 366 days.
//
// The property DateRange is required.
type InsightNewParamsReportContext struct {
	DateRange InsightNewParamsReportContextDateRangeUnion `json:"dateRange,omitzero" api:"required"`
	paramObj
}

func (r InsightNewParamsReportContext) MarshalJSON() (data []byte, err error) {
	type shadow InsightNewParamsReportContext
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightNewParamsReportContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InsightNewParamsReportContextDateRangeUnion struct {
	OfInsightNewsReportContextDateRangeObject  *InsightNewParamsReportContextDateRangeObject  `json:",omitzero,inline"`
	OfInsightNewsReportContextDateRangeObject2 *InsightNewParamsReportContextDateRangeObject2 `json:",omitzero,inline"`
	paramUnion
}

func (u InsightNewParamsReportContextDateRangeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInsightNewsReportContextDateRangeObject, u.OfInsightNewsReportContextDateRangeObject2)
}
func (u *InsightNewParamsReportContextDateRangeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties From, To are required.
type InsightNewParamsReportContextDateRangeObject struct {
	From string `json:"from" api:"required"`
	To   string `json:"to" api:"required"`
	// Any of "UTC".
	TimeZone string `json:"timeZone,omitzero"`
	paramObj
}

func (r InsightNewParamsReportContextDateRangeObject) MarshalJSON() (data []byte, err error) {
	type shadow InsightNewParamsReportContextDateRangeObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightNewParamsReportContextDateRangeObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InsightNewParamsReportContextDateRangeObject](
		"timeZone", "UTC",
	)
}

// The property Preset is required.
type InsightNewParamsReportContextDateRangeObject2 struct {
	// Any of "today", "yesterday", "last_7_days", "last_30_days", "last_90_days",
	// "last_12_months", "week_to_date", "month_to_date", "year_to_date".
	Preset string `json:"preset,omitzero" api:"required"`
	// Any of "UTC".
	TimeZone string `json:"timeZone,omitzero"`
	paramObj
}

func (r InsightNewParamsReportContextDateRangeObject2) MarshalJSON() (data []byte, err error) {
	type shadow InsightNewParamsReportContextDateRangeObject2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightNewParamsReportContextDateRangeObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InsightNewParamsReportContextDateRangeObject2](
		"preset", "today", "yesterday", "last_7_days", "last_30_days", "last_90_days", "last_12_months", "week_to_date", "month_to_date", "year_to_date",
	)
	apijson.RegisterFieldValidator[InsightNewParamsReportContextDateRangeObject2](
		"timeZone", "UTC",
	)
}

// Complete chart settings. Replaces the saved view when supplied on PATCH; omitted
// view fields reset to their defaults.
type InsightNewParamsView struct {
	Cumulative      param.Opt[bool]              `json:"cumulative,omitzero"`
	SmoothingWindow param.Opt[int64]             `json:"smoothingWindow,omitzero"`
	Stacked         param.Opt[bool]              `json:"stacked,omitzero"`
	Series          []InsightNewParamsViewSeries `json:"series,omitzero"`
	// Any of "none", "rolling_average".
	Smoothing string                        `json:"smoothing,omitzero"`
	TableSort InsightNewParamsViewTableSort `json:"tableSort,omitzero"`
	// Any of "line", "column", "bar", "pie", "table", "metric".
	Visualization string `json:"visualization,omitzero"`
	// Any of "linear", "log".
	YAxisScale string `json:"yAxisScale,omitzero"`
	paramObj
}

func (r InsightNewParamsView) MarshalJSON() (data []byte, err error) {
	type shadow InsightNewParamsView
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightNewParamsView) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InsightNewParamsView](
		"smoothing", "none", "rolling_average",
	)
	apijson.RegisterFieldValidator[InsightNewParamsView](
		"visualization", "line", "column", "bar", "pie", "table", "metric",
	)
	apijson.RegisterFieldValidator[InsightNewParamsView](
		"yAxisScale", "linear", "log",
	)
}

// The property MetricID is required.
type InsightNewParamsViewSeries struct {
	MetricID string            `json:"metricId" api:"required" format:"uuid"`
	Color    param.Opt[string] `json:"color,omitzero"`
	Hidden   param.Opt[bool]   `json:"hidden,omitzero"`
	paramObj
}

func (r InsightNewParamsViewSeries) MarshalJSON() (data []byte, err error) {
	type shadow InsightNewParamsViewSeries
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightNewParamsViewSeries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Column, Direction are required.
type InsightNewParamsViewTableSort struct {
	// Any of "label", "range_value", "average".
	Column string `json:"column,omitzero" api:"required"`
	// Any of "asc", "desc".
	Direction string `json:"direction,omitzero" api:"required"`
	paramObj
}

func (r InsightNewParamsViewTableSort) MarshalJSON() (data []byte, err error) {
	type shadow InsightNewParamsViewTableSort
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightNewParamsViewTableSort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InsightNewParamsViewTableSort](
		"column", "label", "range_value", "average",
	)
	apijson.RegisterFieldValidator[InsightNewParamsViewTableSort](
		"direction", "asc", "desc",
	)
}

type InsightUpdateParams struct {
	// Revision returned by get/list or the last write. A stale revision returns 409;
	// reload before retrying.
	ExpectedRevision string `json:"expectedRevision" api:"required" format:"uuid"`
	// Optional description; null clears it.
	Description param.Opt[string] `json:"description,omitzero"`
	// Display name for the saved Insight.
	Name param.Opt[string] `json:"name,omitzero"`
	// Version 1 Insight definition: one to ten metrics and a day, week, or month
	// interval. Sum and average require a numeric event property. Filters use the
	// analytics property catalog.
	Definition InsightUpdateParamsDefinition `json:"definition,omitzero"`
	// UTC date range using either a preset or inclusive from/to calendar dates, up to
	// 366 days.
	ReportContext InsightUpdateParamsReportContext `json:"reportContext,omitzero"`
	// Complete chart settings. Replaces the saved view when supplied on PATCH; omitted
	// view fields reset to their defaults.
	View InsightUpdateParamsView `json:"view,omitzero"`
	paramObj
}

func (r InsightUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow InsightUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Version 1 Insight definition: one to ten metrics and a day, week, or month
// interval. Sum and average require a numeric event property. Filters use the
// analytics property catalog.
//
// The properties Interval, Kind, Metrics, Version are required.
type InsightUpdateParamsDefinition struct {
	// Any of "day", "week", "month".
	Interval string `json:"interval,omitzero" api:"required"`
	// Any of "insight".
	Kind    string                                `json:"kind,omitzero" api:"required"`
	Metrics []InsightUpdateParamsDefinitionMetric `json:"metrics,omitzero" api:"required"`
	// Any of 1.
	Version float64                             `json:"version,omitzero" api:"required"`
	Filter  InsightUpdateParamsDefinitionFilter `json:"filter,omitzero"`
	paramObj
}

func (r InsightUpdateParamsDefinition) MarshalJSON() (data []byte, err error) {
	type shadow InsightUpdateParamsDefinition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightUpdateParamsDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InsightUpdateParamsDefinition](
		"interval", "day", "week", "month",
	)
	apijson.RegisterFieldValidator[InsightUpdateParamsDefinition](
		"kind", "insight",
	)
	apijson.RegisterFieldValidator[InsightUpdateParamsDefinition](
		"version", 1,
	)
}

// The properties ID, Aggregation, EventNames, Name are required.
type InsightUpdateParamsDefinitionMetric struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Any of "event_count", "unique_visitors", "sum", "average".
	Aggregation string                                      `json:"aggregation,omitzero" api:"required"`
	EventNames  []string                                    `json:"eventNames,omitzero" api:"required"`
	Name        string                                      `json:"name" api:"required"`
	Filter      InsightUpdateParamsDefinitionMetricFilter   `json:"filter,omitzero"`
	Property    InsightUpdateParamsDefinitionMetricProperty `json:"property,omitzero"`
	paramObj
}

func (r InsightUpdateParamsDefinitionMetric) MarshalJSON() (data []byte, err error) {
	type shadow InsightUpdateParamsDefinitionMetric
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightUpdateParamsDefinitionMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InsightUpdateParamsDefinitionMetric](
		"aggregation", "event_count", "unique_visitors", "sum", "average",
	)
}

// The properties Filter, Version are required.
type InsightUpdateParamsDefinitionMetricFilter struct {
	Filter InsightUpdateParamsDefinitionMetricFilterFilterUnion `json:"filter,omitzero" api:"required"`
	// Any of 1.
	Version float64 `json:"version,omitzero" api:"required"`
	paramObj
}

func (r InsightUpdateParamsDefinitionMetricFilter) MarshalJSON() (data []byte, err error) {
	type shadow InsightUpdateParamsDefinitionMetricFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightUpdateParamsDefinitionMetricFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InsightUpdateParamsDefinitionMetricFilter](
		"version", 1,
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InsightUpdateParamsDefinitionMetricFilterFilterUnion struct {
	OfInsightUpdatesDefinitionMetricFilterFilterObject  *InsightUpdateParamsDefinitionMetricFilterFilterObject  `json:",omitzero,inline"`
	OfInsightUpdatesDefinitionMetricFilterFilterObject2 *InsightUpdateParamsDefinitionMetricFilterFilterObject2 `json:",omitzero,inline"`
	OfInsightUpdatesDefinitionMetricFilterFilterObject3 *InsightUpdateParamsDefinitionMetricFilterFilterObject3 `json:",omitzero,inline"`
	paramUnion
}

func (u InsightUpdateParamsDefinitionMetricFilterFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInsightUpdatesDefinitionMetricFilterFilterObject, u.OfInsightUpdatesDefinitionMetricFilterFilterObject2, u.OfInsightUpdatesDefinitionMetricFilterFilterObject3)
}
func (u *InsightUpdateParamsDefinitionMetricFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Children, Kind are required.
type InsightUpdateParamsDefinitionMetricFilterFilterObject struct {
	Children []any `json:"children,omitzero" api:"required"`
	// Any of "and", "or".
	Kind string `json:"kind,omitzero" api:"required"`
	paramObj
}

func (r InsightUpdateParamsDefinitionMetricFilterFilterObject) MarshalJSON() (data []byte, err error) {
	type shadow InsightUpdateParamsDefinitionMetricFilterFilterObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightUpdateParamsDefinitionMetricFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InsightUpdateParamsDefinitionMetricFilterFilterObject](
		"kind", "and", "or",
	)
}

// The properties Child, Kind are required.
type InsightUpdateParamsDefinitionMetricFilterFilterObject2 struct {
	Child any `json:"child,omitzero" api:"required"`
	// Any of "not".
	Kind string `json:"kind,omitzero" api:"required"`
	paramObj
}

func (r InsightUpdateParamsDefinitionMetricFilterFilterObject2) MarshalJSON() (data []byte, err error) {
	type shadow InsightUpdateParamsDefinitionMetricFilterFilterObject2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightUpdateParamsDefinitionMetricFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InsightUpdateParamsDefinitionMetricFilterFilterObject2](
		"kind", "not",
	)
}

// The properties Kind, Operator, Property are required.
type InsightUpdateParamsDefinitionMetricFilterFilterObject3 struct {
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
	Value InsightUpdateParamsDefinitionMetricFilterFilterObject3ValueUnion `json:"value,omitzero"`
	paramObj
}

func (r InsightUpdateParamsDefinitionMetricFilterFilterObject3) MarshalJSON() (data []byte, err error) {
	type shadow InsightUpdateParamsDefinitionMetricFilterFilterObject3
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightUpdateParamsDefinitionMetricFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InsightUpdateParamsDefinitionMetricFilterFilterObject3](
		"kind", "predicate",
	)
	apijson.RegisterFieldValidator[InsightUpdateParamsDefinitionMetricFilterFilterObject3](
		"operator", "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains", "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in", "exists", "missing", "is_null",
	)
	apijson.RegisterFieldValidator[InsightUpdateParamsDefinitionMetricFilterFilterObject3](
		"type", "string", "number", "boolean",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InsightUpdateParamsDefinitionMetricFilterFilterObject3ValueUnion struct {
	OfString                                                      param.Opt[string]                                                           `json:",omitzero,inline"`
	OfFloat                                                       param.Opt[float64]                                                          `json:",omitzero,inline"`
	OfBool                                                        param.Opt[bool]                                                             `json:",omitzero,inline"`
	OfInsightUpdatesDefinitionMetricFilterFilterObject3ValueArray []InsightUpdateParamsDefinitionMetricFilterFilterObject3ValueArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u InsightUpdateParamsDefinitionMetricFilterFilterObject3ValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool, u.OfInsightUpdatesDefinitionMetricFilterFilterObject3ValueArray)
}
func (u *InsightUpdateParamsDefinitionMetricFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InsightUpdateParamsDefinitionMetricFilterFilterObject3ValueArrayItemUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u InsightUpdateParamsDefinitionMetricFilterFilterObject3ValueArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *InsightUpdateParamsDefinitionMetricFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Path, Property are required.
type InsightUpdateParamsDefinitionMetricProperty struct {
	Path []string `json:"path,omitzero" api:"required"`
	// Any of "event.properties".
	Property string `json:"property,omitzero" api:"required"`
	// Any of "number".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r InsightUpdateParamsDefinitionMetricProperty) MarshalJSON() (data []byte, err error) {
	type shadow InsightUpdateParamsDefinitionMetricProperty
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightUpdateParamsDefinitionMetricProperty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InsightUpdateParamsDefinitionMetricProperty](
		"property", "event.properties",
	)
	apijson.RegisterFieldValidator[InsightUpdateParamsDefinitionMetricProperty](
		"type", "number",
	)
}

// The properties Filter, Version are required.
type InsightUpdateParamsDefinitionFilter struct {
	Filter InsightUpdateParamsDefinitionFilterFilterUnion `json:"filter,omitzero" api:"required"`
	// Any of 1.
	Version float64 `json:"version,omitzero" api:"required"`
	paramObj
}

func (r InsightUpdateParamsDefinitionFilter) MarshalJSON() (data []byte, err error) {
	type shadow InsightUpdateParamsDefinitionFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightUpdateParamsDefinitionFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InsightUpdateParamsDefinitionFilter](
		"version", 1,
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InsightUpdateParamsDefinitionFilterFilterUnion struct {
	OfInsightUpdatesDefinitionFilterFilterObject  *InsightUpdateParamsDefinitionFilterFilterObject  `json:",omitzero,inline"`
	OfInsightUpdatesDefinitionFilterFilterObject2 *InsightUpdateParamsDefinitionFilterFilterObject2 `json:",omitzero,inline"`
	OfInsightUpdatesDefinitionFilterFilterObject3 *InsightUpdateParamsDefinitionFilterFilterObject3 `json:",omitzero,inline"`
	paramUnion
}

func (u InsightUpdateParamsDefinitionFilterFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInsightUpdatesDefinitionFilterFilterObject, u.OfInsightUpdatesDefinitionFilterFilterObject2, u.OfInsightUpdatesDefinitionFilterFilterObject3)
}
func (u *InsightUpdateParamsDefinitionFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Children, Kind are required.
type InsightUpdateParamsDefinitionFilterFilterObject struct {
	Children []any `json:"children,omitzero" api:"required"`
	// Any of "and", "or".
	Kind string `json:"kind,omitzero" api:"required"`
	paramObj
}

func (r InsightUpdateParamsDefinitionFilterFilterObject) MarshalJSON() (data []byte, err error) {
	type shadow InsightUpdateParamsDefinitionFilterFilterObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightUpdateParamsDefinitionFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InsightUpdateParamsDefinitionFilterFilterObject](
		"kind", "and", "or",
	)
}

// The properties Child, Kind are required.
type InsightUpdateParamsDefinitionFilterFilterObject2 struct {
	Child any `json:"child,omitzero" api:"required"`
	// Any of "not".
	Kind string `json:"kind,omitzero" api:"required"`
	paramObj
}

func (r InsightUpdateParamsDefinitionFilterFilterObject2) MarshalJSON() (data []byte, err error) {
	type shadow InsightUpdateParamsDefinitionFilterFilterObject2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightUpdateParamsDefinitionFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InsightUpdateParamsDefinitionFilterFilterObject2](
		"kind", "not",
	)
}

// The properties Kind, Operator, Property are required.
type InsightUpdateParamsDefinitionFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind,omitzero" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator,omitzero" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path,omitzero"`
	// Any of "string", "number", "boolean".
	Type  string                                                     `json:"type,omitzero"`
	Value InsightUpdateParamsDefinitionFilterFilterObject3ValueUnion `json:"value,omitzero"`
	paramObj
}

func (r InsightUpdateParamsDefinitionFilterFilterObject3) MarshalJSON() (data []byte, err error) {
	type shadow InsightUpdateParamsDefinitionFilterFilterObject3
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightUpdateParamsDefinitionFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InsightUpdateParamsDefinitionFilterFilterObject3](
		"kind", "predicate",
	)
	apijson.RegisterFieldValidator[InsightUpdateParamsDefinitionFilterFilterObject3](
		"operator", "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains", "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in", "exists", "missing", "is_null",
	)
	apijson.RegisterFieldValidator[InsightUpdateParamsDefinitionFilterFilterObject3](
		"type", "string", "number", "boolean",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InsightUpdateParamsDefinitionFilterFilterObject3ValueUnion struct {
	OfString                                                param.Opt[string]                                                     `json:",omitzero,inline"`
	OfFloat                                                 param.Opt[float64]                                                    `json:",omitzero,inline"`
	OfBool                                                  param.Opt[bool]                                                       `json:",omitzero,inline"`
	OfInsightUpdatesDefinitionFilterFilterObject3ValueArray []InsightUpdateParamsDefinitionFilterFilterObject3ValueArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u InsightUpdateParamsDefinitionFilterFilterObject3ValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool, u.OfInsightUpdatesDefinitionFilterFilterObject3ValueArray)
}
func (u *InsightUpdateParamsDefinitionFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InsightUpdateParamsDefinitionFilterFilterObject3ValueArrayItemUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u InsightUpdateParamsDefinitionFilterFilterObject3ValueArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *InsightUpdateParamsDefinitionFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// UTC date range using either a preset or inclusive from/to calendar dates, up to
// 366 days.
//
// The property DateRange is required.
type InsightUpdateParamsReportContext struct {
	DateRange InsightUpdateParamsReportContextDateRangeUnion `json:"dateRange,omitzero" api:"required"`
	paramObj
}

func (r InsightUpdateParamsReportContext) MarshalJSON() (data []byte, err error) {
	type shadow InsightUpdateParamsReportContext
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightUpdateParamsReportContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InsightUpdateParamsReportContextDateRangeUnion struct {
	OfInsightUpdatesReportContextDateRangeObject  *InsightUpdateParamsReportContextDateRangeObject  `json:",omitzero,inline"`
	OfInsightUpdatesReportContextDateRangeObject2 *InsightUpdateParamsReportContextDateRangeObject2 `json:",omitzero,inline"`
	paramUnion
}

func (u InsightUpdateParamsReportContextDateRangeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInsightUpdatesReportContextDateRangeObject, u.OfInsightUpdatesReportContextDateRangeObject2)
}
func (u *InsightUpdateParamsReportContextDateRangeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties From, To are required.
type InsightUpdateParamsReportContextDateRangeObject struct {
	From string `json:"from" api:"required"`
	To   string `json:"to" api:"required"`
	// Any of "UTC".
	TimeZone string `json:"timeZone,omitzero"`
	paramObj
}

func (r InsightUpdateParamsReportContextDateRangeObject) MarshalJSON() (data []byte, err error) {
	type shadow InsightUpdateParamsReportContextDateRangeObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightUpdateParamsReportContextDateRangeObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InsightUpdateParamsReportContextDateRangeObject](
		"timeZone", "UTC",
	)
}

// The property Preset is required.
type InsightUpdateParamsReportContextDateRangeObject2 struct {
	// Any of "today", "yesterday", "last_7_days", "last_30_days", "last_90_days",
	// "last_12_months", "week_to_date", "month_to_date", "year_to_date".
	Preset string `json:"preset,omitzero" api:"required"`
	// Any of "UTC".
	TimeZone string `json:"timeZone,omitzero"`
	paramObj
}

func (r InsightUpdateParamsReportContextDateRangeObject2) MarshalJSON() (data []byte, err error) {
	type shadow InsightUpdateParamsReportContextDateRangeObject2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightUpdateParamsReportContextDateRangeObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InsightUpdateParamsReportContextDateRangeObject2](
		"preset", "today", "yesterday", "last_7_days", "last_30_days", "last_90_days", "last_12_months", "week_to_date", "month_to_date", "year_to_date",
	)
	apijson.RegisterFieldValidator[InsightUpdateParamsReportContextDateRangeObject2](
		"timeZone", "UTC",
	)
}

// Complete chart settings. Replaces the saved view when supplied on PATCH; omitted
// view fields reset to their defaults.
type InsightUpdateParamsView struct {
	Cumulative      param.Opt[bool]                 `json:"cumulative,omitzero"`
	SmoothingWindow param.Opt[int64]                `json:"smoothingWindow,omitzero"`
	Stacked         param.Opt[bool]                 `json:"stacked,omitzero"`
	Series          []InsightUpdateParamsViewSeries `json:"series,omitzero"`
	// Any of "none", "rolling_average".
	Smoothing string                           `json:"smoothing,omitzero"`
	TableSort InsightUpdateParamsViewTableSort `json:"tableSort,omitzero"`
	// Any of "line", "column", "bar", "pie", "table", "metric".
	Visualization string `json:"visualization,omitzero"`
	// Any of "linear", "log".
	YAxisScale string `json:"yAxisScale,omitzero"`
	paramObj
}

func (r InsightUpdateParamsView) MarshalJSON() (data []byte, err error) {
	type shadow InsightUpdateParamsView
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightUpdateParamsView) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InsightUpdateParamsView](
		"smoothing", "none", "rolling_average",
	)
	apijson.RegisterFieldValidator[InsightUpdateParamsView](
		"visualization", "line", "column", "bar", "pie", "table", "metric",
	)
	apijson.RegisterFieldValidator[InsightUpdateParamsView](
		"yAxisScale", "linear", "log",
	)
}

// The property MetricID is required.
type InsightUpdateParamsViewSeries struct {
	MetricID string            `json:"metricId" api:"required" format:"uuid"`
	Color    param.Opt[string] `json:"color,omitzero"`
	Hidden   param.Opt[bool]   `json:"hidden,omitzero"`
	paramObj
}

func (r InsightUpdateParamsViewSeries) MarshalJSON() (data []byte, err error) {
	type shadow InsightUpdateParamsViewSeries
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightUpdateParamsViewSeries) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Column, Direction are required.
type InsightUpdateParamsViewTableSort struct {
	// Any of "label", "range_value", "average".
	Column string `json:"column,omitzero" api:"required"`
	// Any of "asc", "desc".
	Direction string `json:"direction,omitzero" api:"required"`
	paramObj
}

func (r InsightUpdateParamsViewTableSort) MarshalJSON() (data []byte, err error) {
	type shadow InsightUpdateParamsViewTableSort
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InsightUpdateParamsViewTableSort) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InsightUpdateParamsViewTableSort](
		"column", "label", "range_value", "average",
	)
	apijson.RegisterFieldValidator[InsightUpdateParamsViewTableSort](
		"direction", "asc", "desc",
	)
}

type InsightDeleteParams struct {
	// Revision returned by get/list or the last write. A stale revision returns 409;
	// reload before retrying.
	ExpectedRevision string `query:"expectedRevision" api:"required" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [InsightDeleteParams]'s query parameters as `url.Values`.
func (r InsightDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InsightResultsParams struct {
	// Revision returned by get/list or the last write. A stale revision returns 409;
	// reload before retrying.
	ExpectedRevision string `query:"expectedRevision" api:"required" format:"uuid" json:"-"`
	// Omit or send null to use the saved cumulative setting; true or false overrides
	// it for this query.
	CumulativeOverride param.Opt[bool] `query:"cumulativeOverride,omitzero" json:"-"`
	// Omit to use saved smoothing, send null to disable it, or supply a window of 2–30
	// buckets.
	SmoothingWindowOverride param.Opt[int64] `query:"smoothingWindowOverride,omitzero" json:"-"`
	// Compare with the previous period, previous year, or a custom starting date.
	// JSON-encoded in the REST query string; never saved.
	Comparison param.Opt[string] `query:"comparison,omitzero" json:"-"`
	// Complete date range override for this query. JSON-encoded in the REST query
	// string; does not change the saved report.
	ReportContextOverride param.Opt[string] `query:"reportContextOverride,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InsightResultsParams]'s query parameters as `url.Values`.
func (r InsightResultsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InsightPreviewParams struct {
	// Object with required definition and reportContext, plus optional
	// smoothingWindow, cumulative, and comparison. Chart view settings belong to saved
	// reports. JSON-encode input in the REST query string; maximum 16384 bytes.
	Input string `query:"input" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [InsightPreviewParams]'s query parameters as `url.Values`.
func (r InsightPreviewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
