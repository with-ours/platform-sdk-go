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

// AttributionReportService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAttributionReportService] method instead.
type AttributionReportService struct {
	Options []option.RequestOption
}

// NewAttributionReportService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAttributionReportService(opts ...option.RequestOption) (r AttributionReportService) {
	r = AttributionReportService{}
	r.Options = opts
	return
}

// List saved attribution reports in stable ID order with case-insensitive name
// search. Defaults to 20 results; each request scans at most 100 reports. Search
// and unavailable or deleted entries can yield short pages. Continue with
// pagination.nextCursor while hasMore is true and preserve unavailableCount.
// Requires Analytics Suite access. Requires API-key scope or current OAuth user
// permission: web-analytics:view
func (r *AttributionReportService) List(ctx context.Context, query AttributionReportListParams, opts ...option.RequestOption) (res *pagination.Cursor[AttributionReportListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "rest/v1/attribution-reports"
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

// List saved attribution reports in stable ID order with case-insensitive name
// search. Defaults to 20 results; each request scans at most 100 reports. Search
// and unavailable or deleted entries can yield short pages. Continue with
// pagination.nextCursor while hasMore is true and preserve unavailableCount.
// Requires Analytics Suite access. Requires API-key scope or current OAuth user
// permission: web-analytics:view
func (r *AttributionReportService) ListAutoPaging(ctx context.Context, query AttributionReportListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[AttributionReportListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Save an attribution report with a default date context and a definition of up to
// five conversion outcomes, the credit model, counting unit, lookback window, and
// currency. Use GET /rest/v1/attribution-reports/preview to compute results.
// Requires Analytics Suite access. Requires API-key scope or current OAuth user
// permission: web-analytics:write
func (r *AttributionReportService) New(ctx context.Context, body AttributionReportNewParams, opts ...option.RequestOption) (res *AttributionReportNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/attribution-reports"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a saved attribution report including its revision, default date context, and
// attribution definition. Requires Analytics Suite access; unavailable reports
// return 404. Requires API-key scope or current OAuth user permission:
// web-analytics:view
func (r *AttributionReportService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AttributionReportGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/attribution-reports/%s", url.PathEscape(id))
	err = requestconfig.CheckPathSegments(path)
	if err != nil {
		return nil, err
	}
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update an attribution report using expectedRevision. Omitted fields remain
// unchanged; null clears the description. Supplied reportContext and definition
// replace the entire saved value. Requires Analytics Suite access. Requires
// API-key scope or current OAuth user permission: web-analytics:write
func (r *AttributionReportService) Update(ctx context.Context, id string, body AttributionReportUpdateParams, opts ...option.RequestOption) (res *AttributionReportUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/attribution-reports/%s", url.PathEscape(id))
	err = requestconfig.CheckPathSegments(path)
	if err != nil {
		return nil, err
	}
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete an attribution report using expectedRevision. Requires Analytics Suite
// access. Requires API-key scope or current OAuth user permission:
// web-analytics:write
func (r *AttributionReportService) Delete(ctx context.Context, id string, body AttributionReportDeleteParams, opts ...option.RequestOption) (res *AttributionReportDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/attribution-reports/%s", url.PathEscape(id))
	err = requestconfig.CheckPathSegments(path)
	if err != nil {
		return nil, err
	}
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// Duplicate an attribution report using the expectedRevision in the request body.
// Copies the description, date context, and definition under a new id with
// "(copy)" appended to the name. Requires Analytics Suite access. Requires API-key
// scope or current OAuth user permission: web-analytics:write, web-analytics:view
func (r *AttributionReportService) Duplicate(ctx context.Context, id string, body AttributionReportDuplicateParams, opts ...option.RequestOption) (res *AttributionReportDuplicateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/attribution-reports/%s/duplicate", url.PathEscape(id))
	err = requestconfig.CheckPathSegments(path)
	if err != nil {
		return nil, err
	}
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Compute attribution credit for a saved or draft definition without saving it.
// JSON-encode input as {definition,from,to,comparison?,interval?}; dates are
// inclusive UTC calendar strings spanning at most 92 days. Results include a total
// row, channel, source/medium, campaign, and term breakdowns, and a time series,
// with credit under every model. Maximum 16384 bytes before URL encoding. Requires
// Analytics Suite access. Requires API-key scope or current OAuth user permission:
// web-analytics:view
func (r *AttributionReportService) Preview(ctx context.Context, query AttributionReportPreviewParams, opts ...option.RequestOption) (res *AttributionReportPreviewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/attribution-reports/preview"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List the visitors behind one number in an attribution result: an outcome,
// period, row, touch role (first, last, any, or assist), and optional interval
// bucket. Send the same definition, dates, comparison, and interval that produced
// the result. Each visitor includes the credit and value they give that row, their
// credited first and last touch, conversion time, sessions in the lookback window,
// and up to three recent session replays. totalUnits equals the first- or
// last-touch credit for those roles and the row creditedUnits for any. Ordered
// newest conversion first; continue with pagination.nextCursor while hasMore is
// true. Requires Analytics Suite access and permission to view visitor data.
// Requires API-key scope or current OAuth user permission: web-analytics:view,
// report:view-visitor
func (r *AttributionReportService) Visitors(ctx context.Context, query AttributionReportVisitorsParams, opts ...option.RequestOption) (res *AttributionReportVisitorsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/attribution-reports/visitors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Explain what moved in an attribution result: new sources, campaigns, and terms
// with credited conversions and no sessions in the prior 180 days; the largest
// share-adjusted credit changes against the comparison period; channels credited
// very differently by first and last touch; and rows that went quiet. Takes the
// same JSON input as the preview. Without a comparison only new values are
// returned. Returns at most 8 aggregate findings with their baselines and sample
// sizes, ranked by effect size and not tested for significance. Requires Analytics
// Suite access. Requires API-key scope or current OAuth user permission:
// web-analytics:view
func (r *AttributionReportService) Insights(ctx context.Context, query AttributionReportInsightsParams, opts ...option.RequestOption) (res *AttributionReportInsightsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/attribution-reports/insights"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AttributionReportListResponse struct {
	ID              string    `json:"id" api:"required" format:"uuid"`
	CreatedAt       time.Time `json:"createdAt" api:"required" format:"date-time"`
	CreatedByUserID string    `json:"createdByUserId" api:"required"`
	// Version 1 attribution definition: one to five conversion outcomes, each with one
	// to ten unique event names, an optional filter, and an optional numeric event
	// property summed as value. Set the credit model (default last_touch), counting
	// unit (default unique), lookback window of 7, 14, 30, 60, or 90 days (default
	// 30), an optional filter applied to every outcome, and the value currency
	// (default USD). Outcome ids identify outcomes in results. Replaces the saved
	// definition on PATCH.
	Definition AttributionReportListResponseDefinition `json:"definition" api:"required"`
	// Display name for the saved attribution report.
	Name string `json:"name" api:"required"`
	// Default UTC date range: a preset or inclusive from/to calendar dates. Replaces
	// the saved context on PATCH. Results cover at most 92 inclusive days.
	ReportContext   AttributionReportListResponseReportContext `json:"reportContext" api:"required"`
	Revision        string                                     `json:"revision" api:"required" format:"uuid"`
	UpdatedAt       time.Time                                  `json:"updatedAt" api:"required" format:"date-time"`
	UpdatedByUserID string                                     `json:"updatedByUserId" api:"required"`
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
		Description     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportListResponse) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Version 1 attribution definition: one to five conversion outcomes, each with one
// to ten unique event names, an optional filter, and an optional numeric event
// property summed as value. Set the credit model (default last_touch), counting
// unit (default unique), lookback window of 7, 14, 30, 60, or 90 days (default
// 30), an optional filter applied to every outcome, and the value currency
// (default USD). Outcome ids identify outcomes in results. Replaces the saved
// definition on PATCH.
type AttributionReportListResponseDefinition struct {
	// Any of "attribution".
	Kind     string                                           `json:"kind" api:"required"`
	Outcomes []AttributionReportListResponseDefinitionOutcome `json:"outcomes" api:"required"`
	// Any of 1.
	Version float64 `json:"version" api:"required"`
	// Any of "unique", "totals", "sessions".
	Counting string `json:"counting"`
	// Any of "USD", "CAD", "GBP", "EUR", "AUD".
	Currency string                                        `json:"currency"`
	Filter   AttributionReportListResponseDefinitionFilter `json:"filter"`
	// Any of 7, 14, 30, 60, 90.
	LookbackDays float64 `json:"lookbackDays"`
	// Any of "first_touch", "last_touch", "linear", "u_shaped", "j_shaped",
	// "time_decay".
	Model string `json:"model"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind         respjson.Field
		Outcomes     respjson.Field
		Version      respjson.Field
		Counting     respjson.Field
		Currency     respjson.Field
		Filter       respjson.Field
		LookbackDays respjson.Field
		Model        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportListResponseDefinition) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportListResponseDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportListResponseDefinitionOutcome struct {
	ID         string                                               `json:"id" api:"required" format:"uuid"`
	EventNames []string                                             `json:"eventNames" api:"required"`
	Label      string                                               `json:"label" api:"required"`
	Filter     AttributionReportListResponseDefinitionOutcomeFilter `json:"filter"`
	Value      AttributionReportListResponseDefinitionOutcomeValue  `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventNames  respjson.Field
		Label       respjson.Field
		Filter      respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportListResponseDefinitionOutcome) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportListResponseDefinitionOutcome) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportListResponseDefinitionOutcomeFilter struct {
	Filter AttributionReportListResponseDefinitionOutcomeFilterFilterUnion `json:"filter" api:"required"`
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
func (r AttributionReportListResponseDefinitionOutcomeFilter) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportListResponseDefinitionOutcomeFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportListResponseDefinitionOutcomeFilterFilterUnion contains all
// possible properties and values from
// [AttributionReportListResponseDefinitionOutcomeFilterFilterObject],
// [AttributionReportListResponseDefinitionOutcomeFilterFilterObject2],
// [AttributionReportListResponseDefinitionOutcomeFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AttributionReportListResponseDefinitionOutcomeFilterFilterUnion struct {
	// This field is from variant
	// [AttributionReportListResponseDefinitionOutcomeFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant
	// [AttributionReportListResponseDefinitionOutcomeFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant
	// [AttributionReportListResponseDefinitionOutcomeFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant
	// [AttributionReportListResponseDefinitionOutcomeFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant
	// [AttributionReportListResponseDefinitionOutcomeFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant
	// [AttributionReportListResponseDefinitionOutcomeFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant
	// [AttributionReportListResponseDefinitionOutcomeFilterFilterObject3].
	Value AttributionReportListResponseDefinitionOutcomeFilterFilterObject3ValueUnion `json:"value"`
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

func (u AttributionReportListResponseDefinitionOutcomeFilterFilterUnion) AsAttributionReportListResponseDefinitionOutcomeFilterFilterObject() (v AttributionReportListResponseDefinitionOutcomeFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportListResponseDefinitionOutcomeFilterFilterUnion) AsAttributionReportListResponseDefinitionOutcomeFilterFilterObject2() (v AttributionReportListResponseDefinitionOutcomeFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportListResponseDefinitionOutcomeFilterFilterUnion) AsAttributionReportListResponseDefinitionOutcomeFilterFilterObject3() (v AttributionReportListResponseDefinitionOutcomeFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportListResponseDefinitionOutcomeFilterFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AttributionReportListResponseDefinitionOutcomeFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportListResponseDefinitionOutcomeFilterFilterObject struct {
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
func (r AttributionReportListResponseDefinitionOutcomeFilterFilterObject) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportListResponseDefinitionOutcomeFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportListResponseDefinitionOutcomeFilterFilterObject2 struct {
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
func (r AttributionReportListResponseDefinitionOutcomeFilterFilterObject2) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportListResponseDefinitionOutcomeFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportListResponseDefinitionOutcomeFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path"`
	// Any of "string", "number", "boolean".
	Type  string                                                                      `json:"type"`
	Value AttributionReportListResponseDefinitionOutcomeFilterFilterObject3ValueUnion `json:"value"`
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
func (r AttributionReportListResponseDefinitionOutcomeFilterFilterObject3) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportListResponseDefinitionOutcomeFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportListResponseDefinitionOutcomeFilterFilterObject3ValueUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]AttributionReportListResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAttributionReportListResponseDefinitionOutcomeFilterFilterObject3ValueArray]
type AttributionReportListResponseDefinitionOutcomeFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AttributionReportListResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion]
	// instead of an object.
	OfAttributionReportListResponseDefinitionOutcomeFilterFilterObject3ValueArray []AttributionReportListResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                                          struct {
		OfString                                                                      respjson.Field
		OfFloat                                                                       respjson.Field
		OfBool                                                                        respjson.Field
		OfAttributionReportListResponseDefinitionOutcomeFilterFilterObject3ValueArray respjson.Field
		raw                                                                           string
	} `json:"-"`
}

func (u AttributionReportListResponseDefinitionOutcomeFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportListResponseDefinitionOutcomeFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportListResponseDefinitionOutcomeFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportListResponseDefinitionOutcomeFilterFilterObject3ValueUnion) AsAttributionReportListResponseDefinitionOutcomeFilterFilterObject3ValueArray() (v []AttributionReportListResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportListResponseDefinitionOutcomeFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AttributionReportListResponseDefinitionOutcomeFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportListResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type AttributionReportListResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion struct {
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

func (u AttributionReportListResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportListResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportListResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportListResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AttributionReportListResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportListResponseDefinitionOutcomeValue struct {
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
func (r AttributionReportListResponseDefinitionOutcomeValue) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportListResponseDefinitionOutcomeValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportListResponseDefinitionFilter struct {
	Filter AttributionReportListResponseDefinitionFilterFilterUnion `json:"filter" api:"required"`
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
func (r AttributionReportListResponseDefinitionFilter) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportListResponseDefinitionFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportListResponseDefinitionFilterFilterUnion contains all possible
// properties and values from
// [AttributionReportListResponseDefinitionFilterFilterObject],
// [AttributionReportListResponseDefinitionFilterFilterObject2],
// [AttributionReportListResponseDefinitionFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AttributionReportListResponseDefinitionFilterFilterUnion struct {
	// This field is from variant
	// [AttributionReportListResponseDefinitionFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant
	// [AttributionReportListResponseDefinitionFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant
	// [AttributionReportListResponseDefinitionFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant
	// [AttributionReportListResponseDefinitionFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant
	// [AttributionReportListResponseDefinitionFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant
	// [AttributionReportListResponseDefinitionFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant
	// [AttributionReportListResponseDefinitionFilterFilterObject3].
	Value AttributionReportListResponseDefinitionFilterFilterObject3ValueUnion `json:"value"`
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

func (u AttributionReportListResponseDefinitionFilterFilterUnion) AsAttributionReportListResponseDefinitionFilterFilterObject() (v AttributionReportListResponseDefinitionFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportListResponseDefinitionFilterFilterUnion) AsAttributionReportListResponseDefinitionFilterFilterObject2() (v AttributionReportListResponseDefinitionFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportListResponseDefinitionFilterFilterUnion) AsAttributionReportListResponseDefinitionFilterFilterObject3() (v AttributionReportListResponseDefinitionFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportListResponseDefinitionFilterFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *AttributionReportListResponseDefinitionFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportListResponseDefinitionFilterFilterObject struct {
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
func (r AttributionReportListResponseDefinitionFilterFilterObject) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportListResponseDefinitionFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportListResponseDefinitionFilterFilterObject2 struct {
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
func (r AttributionReportListResponseDefinitionFilterFilterObject2) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportListResponseDefinitionFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportListResponseDefinitionFilterFilterObject3 struct {
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
	Value AttributionReportListResponseDefinitionFilterFilterObject3ValueUnion `json:"value"`
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
func (r AttributionReportListResponseDefinitionFilterFilterObject3) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportListResponseDefinitionFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportListResponseDefinitionFilterFilterObject3ValueUnion contains
// all possible properties and values from [string], [float64], [bool],
// [[]AttributionReportListResponseDefinitionFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAttributionReportListResponseDefinitionFilterFilterObject3ValueArray]
type AttributionReportListResponseDefinitionFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AttributionReportListResponseDefinitionFilterFilterObject3ValueArrayItemUnion]
	// instead of an object.
	OfAttributionReportListResponseDefinitionFilterFilterObject3ValueArray []AttributionReportListResponseDefinitionFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                                   struct {
		OfString                                                               respjson.Field
		OfFloat                                                                respjson.Field
		OfBool                                                                 respjson.Field
		OfAttributionReportListResponseDefinitionFilterFilterObject3ValueArray respjson.Field
		raw                                                                    string
	} `json:"-"`
}

func (u AttributionReportListResponseDefinitionFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportListResponseDefinitionFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportListResponseDefinitionFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportListResponseDefinitionFilterFilterObject3ValueUnion) AsAttributionReportListResponseDefinitionFilterFilterObject3ValueArray() (v []AttributionReportListResponseDefinitionFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportListResponseDefinitionFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AttributionReportListResponseDefinitionFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportListResponseDefinitionFilterFilterObject3ValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type AttributionReportListResponseDefinitionFilterFilterObject3ValueArrayItemUnion struct {
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

func (u AttributionReportListResponseDefinitionFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportListResponseDefinitionFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportListResponseDefinitionFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportListResponseDefinitionFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AttributionReportListResponseDefinitionFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default UTC date range: a preset or inclusive from/to calendar dates. Replaces
// the saved context on PATCH. Results cover at most 92 inclusive days.
type AttributionReportListResponseReportContext struct {
	DateRange AttributionReportListResponseReportContextDateRangeUnion `json:"dateRange" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DateRange   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportListResponseReportContext) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportListResponseReportContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportListResponseReportContextDateRangeUnion contains all possible
// properties and values from
// [AttributionReportListResponseReportContextDateRangeObject],
// [AttributionReportListResponseReportContextDateRangeObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AttributionReportListResponseReportContextDateRangeUnion struct {
	// This field is from variant
	// [AttributionReportListResponseReportContextDateRangeObject].
	From string `json:"from"`
	// This field is from variant
	// [AttributionReportListResponseReportContextDateRangeObject].
	To       string `json:"to"`
	TimeZone string `json:"timeZone"`
	// This field is from variant
	// [AttributionReportListResponseReportContextDateRangeObject2].
	Preset string `json:"preset"`
	JSON   struct {
		From     respjson.Field
		To       respjson.Field
		TimeZone respjson.Field
		Preset   respjson.Field
		raw      string
	} `json:"-"`
}

func (u AttributionReportListResponseReportContextDateRangeUnion) AsAttributionReportListResponseReportContextDateRangeObject() (v AttributionReportListResponseReportContextDateRangeObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportListResponseReportContextDateRangeUnion) AsAttributionReportListResponseReportContextDateRangeObject2() (v AttributionReportListResponseReportContextDateRangeObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportListResponseReportContextDateRangeUnion) RawJSON() string { return u.JSON.raw }

func (r *AttributionReportListResponseReportContextDateRangeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportListResponseReportContextDateRangeObject struct {
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
func (r AttributionReportListResponseReportContextDateRangeObject) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportListResponseReportContextDateRangeObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportListResponseReportContextDateRangeObject2 struct {
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
func (r AttributionReportListResponseReportContextDateRangeObject2) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportListResponseReportContextDateRangeObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportNewResponse struct {
	ID              string    `json:"id" api:"required" format:"uuid"`
	CreatedAt       time.Time `json:"createdAt" api:"required" format:"date-time"`
	CreatedByUserID string    `json:"createdByUserId" api:"required"`
	// Version 1 attribution definition: one to five conversion outcomes, each with one
	// to ten unique event names, an optional filter, and an optional numeric event
	// property summed as value. Set the credit model (default last_touch), counting
	// unit (default unique), lookback window of 7, 14, 30, 60, or 90 days (default
	// 30), an optional filter applied to every outcome, and the value currency
	// (default USD). Outcome ids identify outcomes in results. Replaces the saved
	// definition on PATCH.
	Definition AttributionReportNewResponseDefinition `json:"definition" api:"required"`
	// Display name for the saved attribution report.
	Name string `json:"name" api:"required"`
	// Default UTC date range: a preset or inclusive from/to calendar dates. Replaces
	// the saved context on PATCH. Results cover at most 92 inclusive days.
	ReportContext   AttributionReportNewResponseReportContext `json:"reportContext" api:"required"`
	Revision        string                                    `json:"revision" api:"required" format:"uuid"`
	UpdatedAt       time.Time                                 `json:"updatedAt" api:"required" format:"date-time"`
	UpdatedByUserID string                                    `json:"updatedByUserId" api:"required"`
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
		Description     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Version 1 attribution definition: one to five conversion outcomes, each with one
// to ten unique event names, an optional filter, and an optional numeric event
// property summed as value. Set the credit model (default last_touch), counting
// unit (default unique), lookback window of 7, 14, 30, 60, or 90 days (default
// 30), an optional filter applied to every outcome, and the value currency
// (default USD). Outcome ids identify outcomes in results. Replaces the saved
// definition on PATCH.
type AttributionReportNewResponseDefinition struct {
	// Any of "attribution".
	Kind     string                                          `json:"kind" api:"required"`
	Outcomes []AttributionReportNewResponseDefinitionOutcome `json:"outcomes" api:"required"`
	// Any of 1.
	Version float64 `json:"version" api:"required"`
	// Any of "unique", "totals", "sessions".
	Counting string `json:"counting"`
	// Any of "USD", "CAD", "GBP", "EUR", "AUD".
	Currency string                                       `json:"currency"`
	Filter   AttributionReportNewResponseDefinitionFilter `json:"filter"`
	// Any of 7, 14, 30, 60, 90.
	LookbackDays float64 `json:"lookbackDays"`
	// Any of "first_touch", "last_touch", "linear", "u_shaped", "j_shaped",
	// "time_decay".
	Model string `json:"model"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind         respjson.Field
		Outcomes     respjson.Field
		Version      respjson.Field
		Counting     respjson.Field
		Currency     respjson.Field
		Filter       respjson.Field
		LookbackDays respjson.Field
		Model        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportNewResponseDefinition) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportNewResponseDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportNewResponseDefinitionOutcome struct {
	ID         string                                              `json:"id" api:"required" format:"uuid"`
	EventNames []string                                            `json:"eventNames" api:"required"`
	Label      string                                              `json:"label" api:"required"`
	Filter     AttributionReportNewResponseDefinitionOutcomeFilter `json:"filter"`
	Value      AttributionReportNewResponseDefinitionOutcomeValue  `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventNames  respjson.Field
		Label       respjson.Field
		Filter      respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportNewResponseDefinitionOutcome) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportNewResponseDefinitionOutcome) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportNewResponseDefinitionOutcomeFilter struct {
	Filter AttributionReportNewResponseDefinitionOutcomeFilterFilterUnion `json:"filter" api:"required"`
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
func (r AttributionReportNewResponseDefinitionOutcomeFilter) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportNewResponseDefinitionOutcomeFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportNewResponseDefinitionOutcomeFilterFilterUnion contains all
// possible properties and values from
// [AttributionReportNewResponseDefinitionOutcomeFilterFilterObject],
// [AttributionReportNewResponseDefinitionOutcomeFilterFilterObject2],
// [AttributionReportNewResponseDefinitionOutcomeFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AttributionReportNewResponseDefinitionOutcomeFilterFilterUnion struct {
	// This field is from variant
	// [AttributionReportNewResponseDefinitionOutcomeFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant
	// [AttributionReportNewResponseDefinitionOutcomeFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant
	// [AttributionReportNewResponseDefinitionOutcomeFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant
	// [AttributionReportNewResponseDefinitionOutcomeFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant
	// [AttributionReportNewResponseDefinitionOutcomeFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant
	// [AttributionReportNewResponseDefinitionOutcomeFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant
	// [AttributionReportNewResponseDefinitionOutcomeFilterFilterObject3].
	Value AttributionReportNewResponseDefinitionOutcomeFilterFilterObject3ValueUnion `json:"value"`
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

func (u AttributionReportNewResponseDefinitionOutcomeFilterFilterUnion) AsAttributionReportNewResponseDefinitionOutcomeFilterFilterObject() (v AttributionReportNewResponseDefinitionOutcomeFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportNewResponseDefinitionOutcomeFilterFilterUnion) AsAttributionReportNewResponseDefinitionOutcomeFilterFilterObject2() (v AttributionReportNewResponseDefinitionOutcomeFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportNewResponseDefinitionOutcomeFilterFilterUnion) AsAttributionReportNewResponseDefinitionOutcomeFilterFilterObject3() (v AttributionReportNewResponseDefinitionOutcomeFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportNewResponseDefinitionOutcomeFilterFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AttributionReportNewResponseDefinitionOutcomeFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportNewResponseDefinitionOutcomeFilterFilterObject struct {
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
func (r AttributionReportNewResponseDefinitionOutcomeFilterFilterObject) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportNewResponseDefinitionOutcomeFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportNewResponseDefinitionOutcomeFilterFilterObject2 struct {
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
func (r AttributionReportNewResponseDefinitionOutcomeFilterFilterObject2) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportNewResponseDefinitionOutcomeFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportNewResponseDefinitionOutcomeFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path"`
	// Any of "string", "number", "boolean".
	Type  string                                                                     `json:"type"`
	Value AttributionReportNewResponseDefinitionOutcomeFilterFilterObject3ValueUnion `json:"value"`
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
func (r AttributionReportNewResponseDefinitionOutcomeFilterFilterObject3) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportNewResponseDefinitionOutcomeFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportNewResponseDefinitionOutcomeFilterFilterObject3ValueUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]AttributionReportNewResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAttributionReportNewResponseDefinitionOutcomeFilterFilterObject3ValueArray]
type AttributionReportNewResponseDefinitionOutcomeFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AttributionReportNewResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion]
	// instead of an object.
	OfAttributionReportNewResponseDefinitionOutcomeFilterFilterObject3ValueArray []AttributionReportNewResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                                         struct {
		OfString                                                                     respjson.Field
		OfFloat                                                                      respjson.Field
		OfBool                                                                       respjson.Field
		OfAttributionReportNewResponseDefinitionOutcomeFilterFilterObject3ValueArray respjson.Field
		raw                                                                          string
	} `json:"-"`
}

func (u AttributionReportNewResponseDefinitionOutcomeFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportNewResponseDefinitionOutcomeFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportNewResponseDefinitionOutcomeFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportNewResponseDefinitionOutcomeFilterFilterObject3ValueUnion) AsAttributionReportNewResponseDefinitionOutcomeFilterFilterObject3ValueArray() (v []AttributionReportNewResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportNewResponseDefinitionOutcomeFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AttributionReportNewResponseDefinitionOutcomeFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportNewResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type AttributionReportNewResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion struct {
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

func (u AttributionReportNewResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportNewResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportNewResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportNewResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AttributionReportNewResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportNewResponseDefinitionOutcomeValue struct {
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
func (r AttributionReportNewResponseDefinitionOutcomeValue) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportNewResponseDefinitionOutcomeValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportNewResponseDefinitionFilter struct {
	Filter AttributionReportNewResponseDefinitionFilterFilterUnion `json:"filter" api:"required"`
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
func (r AttributionReportNewResponseDefinitionFilter) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportNewResponseDefinitionFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportNewResponseDefinitionFilterFilterUnion contains all possible
// properties and values from
// [AttributionReportNewResponseDefinitionFilterFilterObject],
// [AttributionReportNewResponseDefinitionFilterFilterObject2],
// [AttributionReportNewResponseDefinitionFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AttributionReportNewResponseDefinitionFilterFilterUnion struct {
	// This field is from variant
	// [AttributionReportNewResponseDefinitionFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant
	// [AttributionReportNewResponseDefinitionFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant
	// [AttributionReportNewResponseDefinitionFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant
	// [AttributionReportNewResponseDefinitionFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant
	// [AttributionReportNewResponseDefinitionFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant
	// [AttributionReportNewResponseDefinitionFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant
	// [AttributionReportNewResponseDefinitionFilterFilterObject3].
	Value AttributionReportNewResponseDefinitionFilterFilterObject3ValueUnion `json:"value"`
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

func (u AttributionReportNewResponseDefinitionFilterFilterUnion) AsAttributionReportNewResponseDefinitionFilterFilterObject() (v AttributionReportNewResponseDefinitionFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportNewResponseDefinitionFilterFilterUnion) AsAttributionReportNewResponseDefinitionFilterFilterObject2() (v AttributionReportNewResponseDefinitionFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportNewResponseDefinitionFilterFilterUnion) AsAttributionReportNewResponseDefinitionFilterFilterObject3() (v AttributionReportNewResponseDefinitionFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportNewResponseDefinitionFilterFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *AttributionReportNewResponseDefinitionFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportNewResponseDefinitionFilterFilterObject struct {
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
func (r AttributionReportNewResponseDefinitionFilterFilterObject) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportNewResponseDefinitionFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportNewResponseDefinitionFilterFilterObject2 struct {
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
func (r AttributionReportNewResponseDefinitionFilterFilterObject2) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportNewResponseDefinitionFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportNewResponseDefinitionFilterFilterObject3 struct {
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
	Value AttributionReportNewResponseDefinitionFilterFilterObject3ValueUnion `json:"value"`
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
func (r AttributionReportNewResponseDefinitionFilterFilterObject3) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportNewResponseDefinitionFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportNewResponseDefinitionFilterFilterObject3ValueUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]AttributionReportNewResponseDefinitionFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAttributionReportNewResponseDefinitionFilterFilterObject3ValueArray]
type AttributionReportNewResponseDefinitionFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AttributionReportNewResponseDefinitionFilterFilterObject3ValueArrayItemUnion]
	// instead of an object.
	OfAttributionReportNewResponseDefinitionFilterFilterObject3ValueArray []AttributionReportNewResponseDefinitionFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                                  struct {
		OfString                                                              respjson.Field
		OfFloat                                                               respjson.Field
		OfBool                                                                respjson.Field
		OfAttributionReportNewResponseDefinitionFilterFilterObject3ValueArray respjson.Field
		raw                                                                   string
	} `json:"-"`
}

func (u AttributionReportNewResponseDefinitionFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportNewResponseDefinitionFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportNewResponseDefinitionFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportNewResponseDefinitionFilterFilterObject3ValueUnion) AsAttributionReportNewResponseDefinitionFilterFilterObject3ValueArray() (v []AttributionReportNewResponseDefinitionFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportNewResponseDefinitionFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AttributionReportNewResponseDefinitionFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportNewResponseDefinitionFilterFilterObject3ValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type AttributionReportNewResponseDefinitionFilterFilterObject3ValueArrayItemUnion struct {
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

func (u AttributionReportNewResponseDefinitionFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportNewResponseDefinitionFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportNewResponseDefinitionFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportNewResponseDefinitionFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AttributionReportNewResponseDefinitionFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default UTC date range: a preset or inclusive from/to calendar dates. Replaces
// the saved context on PATCH. Results cover at most 92 inclusive days.
type AttributionReportNewResponseReportContext struct {
	DateRange AttributionReportNewResponseReportContextDateRangeUnion `json:"dateRange" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DateRange   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportNewResponseReportContext) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportNewResponseReportContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportNewResponseReportContextDateRangeUnion contains all possible
// properties and values from
// [AttributionReportNewResponseReportContextDateRangeObject],
// [AttributionReportNewResponseReportContextDateRangeObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AttributionReportNewResponseReportContextDateRangeUnion struct {
	// This field is from variant
	// [AttributionReportNewResponseReportContextDateRangeObject].
	From string `json:"from"`
	// This field is from variant
	// [AttributionReportNewResponseReportContextDateRangeObject].
	To       string `json:"to"`
	TimeZone string `json:"timeZone"`
	// This field is from variant
	// [AttributionReportNewResponseReportContextDateRangeObject2].
	Preset string `json:"preset"`
	JSON   struct {
		From     respjson.Field
		To       respjson.Field
		TimeZone respjson.Field
		Preset   respjson.Field
		raw      string
	} `json:"-"`
}

func (u AttributionReportNewResponseReportContextDateRangeUnion) AsAttributionReportNewResponseReportContextDateRangeObject() (v AttributionReportNewResponseReportContextDateRangeObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportNewResponseReportContextDateRangeUnion) AsAttributionReportNewResponseReportContextDateRangeObject2() (v AttributionReportNewResponseReportContextDateRangeObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportNewResponseReportContextDateRangeUnion) RawJSON() string { return u.JSON.raw }

func (r *AttributionReportNewResponseReportContextDateRangeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportNewResponseReportContextDateRangeObject struct {
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
func (r AttributionReportNewResponseReportContextDateRangeObject) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportNewResponseReportContextDateRangeObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportNewResponseReportContextDateRangeObject2 struct {
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
func (r AttributionReportNewResponseReportContextDateRangeObject2) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportNewResponseReportContextDateRangeObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportGetResponse struct {
	ID              string    `json:"id" api:"required" format:"uuid"`
	CreatedAt       time.Time `json:"createdAt" api:"required" format:"date-time"`
	CreatedByUserID string    `json:"createdByUserId" api:"required"`
	// Version 1 attribution definition: one to five conversion outcomes, each with one
	// to ten unique event names, an optional filter, and an optional numeric event
	// property summed as value. Set the credit model (default last_touch), counting
	// unit (default unique), lookback window of 7, 14, 30, 60, or 90 days (default
	// 30), an optional filter applied to every outcome, and the value currency
	// (default USD). Outcome ids identify outcomes in results. Replaces the saved
	// definition on PATCH.
	Definition AttributionReportGetResponseDefinition `json:"definition" api:"required"`
	// Display name for the saved attribution report.
	Name string `json:"name" api:"required"`
	// Default UTC date range: a preset or inclusive from/to calendar dates. Replaces
	// the saved context on PATCH. Results cover at most 92 inclusive days.
	ReportContext   AttributionReportGetResponseReportContext `json:"reportContext" api:"required"`
	Revision        string                                    `json:"revision" api:"required" format:"uuid"`
	UpdatedAt       time.Time                                 `json:"updatedAt" api:"required" format:"date-time"`
	UpdatedByUserID string                                    `json:"updatedByUserId" api:"required"`
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
		Description     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Version 1 attribution definition: one to five conversion outcomes, each with one
// to ten unique event names, an optional filter, and an optional numeric event
// property summed as value. Set the credit model (default last_touch), counting
// unit (default unique), lookback window of 7, 14, 30, 60, or 90 days (default
// 30), an optional filter applied to every outcome, and the value currency
// (default USD). Outcome ids identify outcomes in results. Replaces the saved
// definition on PATCH.
type AttributionReportGetResponseDefinition struct {
	// Any of "attribution".
	Kind     string                                          `json:"kind" api:"required"`
	Outcomes []AttributionReportGetResponseDefinitionOutcome `json:"outcomes" api:"required"`
	// Any of 1.
	Version float64 `json:"version" api:"required"`
	// Any of "unique", "totals", "sessions".
	Counting string `json:"counting"`
	// Any of "USD", "CAD", "GBP", "EUR", "AUD".
	Currency string                                       `json:"currency"`
	Filter   AttributionReportGetResponseDefinitionFilter `json:"filter"`
	// Any of 7, 14, 30, 60, 90.
	LookbackDays float64 `json:"lookbackDays"`
	// Any of "first_touch", "last_touch", "linear", "u_shaped", "j_shaped",
	// "time_decay".
	Model string `json:"model"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind         respjson.Field
		Outcomes     respjson.Field
		Version      respjson.Field
		Counting     respjson.Field
		Currency     respjson.Field
		Filter       respjson.Field
		LookbackDays respjson.Field
		Model        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportGetResponseDefinition) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportGetResponseDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportGetResponseDefinitionOutcome struct {
	ID         string                                              `json:"id" api:"required" format:"uuid"`
	EventNames []string                                            `json:"eventNames" api:"required"`
	Label      string                                              `json:"label" api:"required"`
	Filter     AttributionReportGetResponseDefinitionOutcomeFilter `json:"filter"`
	Value      AttributionReportGetResponseDefinitionOutcomeValue  `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventNames  respjson.Field
		Label       respjson.Field
		Filter      respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportGetResponseDefinitionOutcome) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportGetResponseDefinitionOutcome) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportGetResponseDefinitionOutcomeFilter struct {
	Filter AttributionReportGetResponseDefinitionOutcomeFilterFilterUnion `json:"filter" api:"required"`
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
func (r AttributionReportGetResponseDefinitionOutcomeFilter) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportGetResponseDefinitionOutcomeFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportGetResponseDefinitionOutcomeFilterFilterUnion contains all
// possible properties and values from
// [AttributionReportGetResponseDefinitionOutcomeFilterFilterObject],
// [AttributionReportGetResponseDefinitionOutcomeFilterFilterObject2],
// [AttributionReportGetResponseDefinitionOutcomeFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AttributionReportGetResponseDefinitionOutcomeFilterFilterUnion struct {
	// This field is from variant
	// [AttributionReportGetResponseDefinitionOutcomeFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant
	// [AttributionReportGetResponseDefinitionOutcomeFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant
	// [AttributionReportGetResponseDefinitionOutcomeFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant
	// [AttributionReportGetResponseDefinitionOutcomeFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant
	// [AttributionReportGetResponseDefinitionOutcomeFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant
	// [AttributionReportGetResponseDefinitionOutcomeFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant
	// [AttributionReportGetResponseDefinitionOutcomeFilterFilterObject3].
	Value AttributionReportGetResponseDefinitionOutcomeFilterFilterObject3ValueUnion `json:"value"`
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

func (u AttributionReportGetResponseDefinitionOutcomeFilterFilterUnion) AsAttributionReportGetResponseDefinitionOutcomeFilterFilterObject() (v AttributionReportGetResponseDefinitionOutcomeFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportGetResponseDefinitionOutcomeFilterFilterUnion) AsAttributionReportGetResponseDefinitionOutcomeFilterFilterObject2() (v AttributionReportGetResponseDefinitionOutcomeFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportGetResponseDefinitionOutcomeFilterFilterUnion) AsAttributionReportGetResponseDefinitionOutcomeFilterFilterObject3() (v AttributionReportGetResponseDefinitionOutcomeFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportGetResponseDefinitionOutcomeFilterFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AttributionReportGetResponseDefinitionOutcomeFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportGetResponseDefinitionOutcomeFilterFilterObject struct {
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
func (r AttributionReportGetResponseDefinitionOutcomeFilterFilterObject) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportGetResponseDefinitionOutcomeFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportGetResponseDefinitionOutcomeFilterFilterObject2 struct {
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
func (r AttributionReportGetResponseDefinitionOutcomeFilterFilterObject2) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportGetResponseDefinitionOutcomeFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportGetResponseDefinitionOutcomeFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path"`
	// Any of "string", "number", "boolean".
	Type  string                                                                     `json:"type"`
	Value AttributionReportGetResponseDefinitionOutcomeFilterFilterObject3ValueUnion `json:"value"`
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
func (r AttributionReportGetResponseDefinitionOutcomeFilterFilterObject3) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportGetResponseDefinitionOutcomeFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportGetResponseDefinitionOutcomeFilterFilterObject3ValueUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]AttributionReportGetResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAttributionReportGetResponseDefinitionOutcomeFilterFilterObject3ValueArray]
type AttributionReportGetResponseDefinitionOutcomeFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AttributionReportGetResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion]
	// instead of an object.
	OfAttributionReportGetResponseDefinitionOutcomeFilterFilterObject3ValueArray []AttributionReportGetResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                                         struct {
		OfString                                                                     respjson.Field
		OfFloat                                                                      respjson.Field
		OfBool                                                                       respjson.Field
		OfAttributionReportGetResponseDefinitionOutcomeFilterFilterObject3ValueArray respjson.Field
		raw                                                                          string
	} `json:"-"`
}

func (u AttributionReportGetResponseDefinitionOutcomeFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportGetResponseDefinitionOutcomeFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportGetResponseDefinitionOutcomeFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportGetResponseDefinitionOutcomeFilterFilterObject3ValueUnion) AsAttributionReportGetResponseDefinitionOutcomeFilterFilterObject3ValueArray() (v []AttributionReportGetResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportGetResponseDefinitionOutcomeFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AttributionReportGetResponseDefinitionOutcomeFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportGetResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type AttributionReportGetResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion struct {
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

func (u AttributionReportGetResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportGetResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportGetResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportGetResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AttributionReportGetResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportGetResponseDefinitionOutcomeValue struct {
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
func (r AttributionReportGetResponseDefinitionOutcomeValue) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportGetResponseDefinitionOutcomeValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportGetResponseDefinitionFilter struct {
	Filter AttributionReportGetResponseDefinitionFilterFilterUnion `json:"filter" api:"required"`
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
func (r AttributionReportGetResponseDefinitionFilter) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportGetResponseDefinitionFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportGetResponseDefinitionFilterFilterUnion contains all possible
// properties and values from
// [AttributionReportGetResponseDefinitionFilterFilterObject],
// [AttributionReportGetResponseDefinitionFilterFilterObject2],
// [AttributionReportGetResponseDefinitionFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AttributionReportGetResponseDefinitionFilterFilterUnion struct {
	// This field is from variant
	// [AttributionReportGetResponseDefinitionFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant
	// [AttributionReportGetResponseDefinitionFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant
	// [AttributionReportGetResponseDefinitionFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant
	// [AttributionReportGetResponseDefinitionFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant
	// [AttributionReportGetResponseDefinitionFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant
	// [AttributionReportGetResponseDefinitionFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant
	// [AttributionReportGetResponseDefinitionFilterFilterObject3].
	Value AttributionReportGetResponseDefinitionFilterFilterObject3ValueUnion `json:"value"`
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

func (u AttributionReportGetResponseDefinitionFilterFilterUnion) AsAttributionReportGetResponseDefinitionFilterFilterObject() (v AttributionReportGetResponseDefinitionFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportGetResponseDefinitionFilterFilterUnion) AsAttributionReportGetResponseDefinitionFilterFilterObject2() (v AttributionReportGetResponseDefinitionFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportGetResponseDefinitionFilterFilterUnion) AsAttributionReportGetResponseDefinitionFilterFilterObject3() (v AttributionReportGetResponseDefinitionFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportGetResponseDefinitionFilterFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *AttributionReportGetResponseDefinitionFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportGetResponseDefinitionFilterFilterObject struct {
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
func (r AttributionReportGetResponseDefinitionFilterFilterObject) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportGetResponseDefinitionFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportGetResponseDefinitionFilterFilterObject2 struct {
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
func (r AttributionReportGetResponseDefinitionFilterFilterObject2) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportGetResponseDefinitionFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportGetResponseDefinitionFilterFilterObject3 struct {
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
	Value AttributionReportGetResponseDefinitionFilterFilterObject3ValueUnion `json:"value"`
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
func (r AttributionReportGetResponseDefinitionFilterFilterObject3) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportGetResponseDefinitionFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportGetResponseDefinitionFilterFilterObject3ValueUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]AttributionReportGetResponseDefinitionFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAttributionReportGetResponseDefinitionFilterFilterObject3ValueArray]
type AttributionReportGetResponseDefinitionFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AttributionReportGetResponseDefinitionFilterFilterObject3ValueArrayItemUnion]
	// instead of an object.
	OfAttributionReportGetResponseDefinitionFilterFilterObject3ValueArray []AttributionReportGetResponseDefinitionFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                                  struct {
		OfString                                                              respjson.Field
		OfFloat                                                               respjson.Field
		OfBool                                                                respjson.Field
		OfAttributionReportGetResponseDefinitionFilterFilterObject3ValueArray respjson.Field
		raw                                                                   string
	} `json:"-"`
}

func (u AttributionReportGetResponseDefinitionFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportGetResponseDefinitionFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportGetResponseDefinitionFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportGetResponseDefinitionFilterFilterObject3ValueUnion) AsAttributionReportGetResponseDefinitionFilterFilterObject3ValueArray() (v []AttributionReportGetResponseDefinitionFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportGetResponseDefinitionFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AttributionReportGetResponseDefinitionFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportGetResponseDefinitionFilterFilterObject3ValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type AttributionReportGetResponseDefinitionFilterFilterObject3ValueArrayItemUnion struct {
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

func (u AttributionReportGetResponseDefinitionFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportGetResponseDefinitionFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportGetResponseDefinitionFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportGetResponseDefinitionFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AttributionReportGetResponseDefinitionFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default UTC date range: a preset or inclusive from/to calendar dates. Replaces
// the saved context on PATCH. Results cover at most 92 inclusive days.
type AttributionReportGetResponseReportContext struct {
	DateRange AttributionReportGetResponseReportContextDateRangeUnion `json:"dateRange" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DateRange   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportGetResponseReportContext) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportGetResponseReportContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportGetResponseReportContextDateRangeUnion contains all possible
// properties and values from
// [AttributionReportGetResponseReportContextDateRangeObject],
// [AttributionReportGetResponseReportContextDateRangeObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AttributionReportGetResponseReportContextDateRangeUnion struct {
	// This field is from variant
	// [AttributionReportGetResponseReportContextDateRangeObject].
	From string `json:"from"`
	// This field is from variant
	// [AttributionReportGetResponseReportContextDateRangeObject].
	To       string `json:"to"`
	TimeZone string `json:"timeZone"`
	// This field is from variant
	// [AttributionReportGetResponseReportContextDateRangeObject2].
	Preset string `json:"preset"`
	JSON   struct {
		From     respjson.Field
		To       respjson.Field
		TimeZone respjson.Field
		Preset   respjson.Field
		raw      string
	} `json:"-"`
}

func (u AttributionReportGetResponseReportContextDateRangeUnion) AsAttributionReportGetResponseReportContextDateRangeObject() (v AttributionReportGetResponseReportContextDateRangeObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportGetResponseReportContextDateRangeUnion) AsAttributionReportGetResponseReportContextDateRangeObject2() (v AttributionReportGetResponseReportContextDateRangeObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportGetResponseReportContextDateRangeUnion) RawJSON() string { return u.JSON.raw }

func (r *AttributionReportGetResponseReportContextDateRangeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportGetResponseReportContextDateRangeObject struct {
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
func (r AttributionReportGetResponseReportContextDateRangeObject) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportGetResponseReportContextDateRangeObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportGetResponseReportContextDateRangeObject2 struct {
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
func (r AttributionReportGetResponseReportContextDateRangeObject2) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportGetResponseReportContextDateRangeObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportUpdateResponse struct {
	ID              string    `json:"id" api:"required" format:"uuid"`
	CreatedAt       time.Time `json:"createdAt" api:"required" format:"date-time"`
	CreatedByUserID string    `json:"createdByUserId" api:"required"`
	// Version 1 attribution definition: one to five conversion outcomes, each with one
	// to ten unique event names, an optional filter, and an optional numeric event
	// property summed as value. Set the credit model (default last_touch), counting
	// unit (default unique), lookback window of 7, 14, 30, 60, or 90 days (default
	// 30), an optional filter applied to every outcome, and the value currency
	// (default USD). Outcome ids identify outcomes in results. Replaces the saved
	// definition on PATCH.
	Definition AttributionReportUpdateResponseDefinition `json:"definition" api:"required"`
	// Display name for the saved attribution report.
	Name string `json:"name" api:"required"`
	// Default UTC date range: a preset or inclusive from/to calendar dates. Replaces
	// the saved context on PATCH. Results cover at most 92 inclusive days.
	ReportContext   AttributionReportUpdateResponseReportContext `json:"reportContext" api:"required"`
	Revision        string                                       `json:"revision" api:"required" format:"uuid"`
	UpdatedAt       time.Time                                    `json:"updatedAt" api:"required" format:"date-time"`
	UpdatedByUserID string                                       `json:"updatedByUserId" api:"required"`
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
		Description     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Version 1 attribution definition: one to five conversion outcomes, each with one
// to ten unique event names, an optional filter, and an optional numeric event
// property summed as value. Set the credit model (default last_touch), counting
// unit (default unique), lookback window of 7, 14, 30, 60, or 90 days (default
// 30), an optional filter applied to every outcome, and the value currency
// (default USD). Outcome ids identify outcomes in results. Replaces the saved
// definition on PATCH.
type AttributionReportUpdateResponseDefinition struct {
	// Any of "attribution".
	Kind     string                                             `json:"kind" api:"required"`
	Outcomes []AttributionReportUpdateResponseDefinitionOutcome `json:"outcomes" api:"required"`
	// Any of 1.
	Version float64 `json:"version" api:"required"`
	// Any of "unique", "totals", "sessions".
	Counting string `json:"counting"`
	// Any of "USD", "CAD", "GBP", "EUR", "AUD".
	Currency string                                          `json:"currency"`
	Filter   AttributionReportUpdateResponseDefinitionFilter `json:"filter"`
	// Any of 7, 14, 30, 60, 90.
	LookbackDays float64 `json:"lookbackDays"`
	// Any of "first_touch", "last_touch", "linear", "u_shaped", "j_shaped",
	// "time_decay".
	Model string `json:"model"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind         respjson.Field
		Outcomes     respjson.Field
		Version      respjson.Field
		Counting     respjson.Field
		Currency     respjson.Field
		Filter       respjson.Field
		LookbackDays respjson.Field
		Model        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportUpdateResponseDefinition) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportUpdateResponseDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportUpdateResponseDefinitionOutcome struct {
	ID         string                                                 `json:"id" api:"required" format:"uuid"`
	EventNames []string                                               `json:"eventNames" api:"required"`
	Label      string                                                 `json:"label" api:"required"`
	Filter     AttributionReportUpdateResponseDefinitionOutcomeFilter `json:"filter"`
	Value      AttributionReportUpdateResponseDefinitionOutcomeValue  `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventNames  respjson.Field
		Label       respjson.Field
		Filter      respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportUpdateResponseDefinitionOutcome) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportUpdateResponseDefinitionOutcome) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportUpdateResponseDefinitionOutcomeFilter struct {
	Filter AttributionReportUpdateResponseDefinitionOutcomeFilterFilterUnion `json:"filter" api:"required"`
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
func (r AttributionReportUpdateResponseDefinitionOutcomeFilter) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportUpdateResponseDefinitionOutcomeFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportUpdateResponseDefinitionOutcomeFilterFilterUnion contains all
// possible properties and values from
// [AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject],
// [AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject2],
// [AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AttributionReportUpdateResponseDefinitionOutcomeFilterFilterUnion struct {
	// This field is from variant
	// [AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant
	// [AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant
	// [AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant
	// [AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant
	// [AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant
	// [AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant
	// [AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3].
	Value AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3ValueUnion `json:"value"`
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

func (u AttributionReportUpdateResponseDefinitionOutcomeFilterFilterUnion) AsAttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject() (v AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportUpdateResponseDefinitionOutcomeFilterFilterUnion) AsAttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject2() (v AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportUpdateResponseDefinitionOutcomeFilterFilterUnion) AsAttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3() (v AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportUpdateResponseDefinitionOutcomeFilterFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AttributionReportUpdateResponseDefinitionOutcomeFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject struct {
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
func (r AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject2 struct {
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
func (r AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject2) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path"`
	// Any of "string", "number", "boolean".
	Type  string                                                                        `json:"type"`
	Value AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3ValueUnion `json:"value"`
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
func (r AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3ValueUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3ValueArray]
type AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion]
	// instead of an object.
	OfAttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3ValueArray []AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                                            struct {
		OfString                                                                        respjson.Field
		OfFloat                                                                         respjson.Field
		OfBool                                                                          respjson.Field
		OfAttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3ValueArray respjson.Field
		raw                                                                             string
	} `json:"-"`
}

func (u AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3ValueUnion) AsAttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3ValueArray() (v []AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion struct {
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

func (u AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AttributionReportUpdateResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportUpdateResponseDefinitionOutcomeValue struct {
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
func (r AttributionReportUpdateResponseDefinitionOutcomeValue) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportUpdateResponseDefinitionOutcomeValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportUpdateResponseDefinitionFilter struct {
	Filter AttributionReportUpdateResponseDefinitionFilterFilterUnion `json:"filter" api:"required"`
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
func (r AttributionReportUpdateResponseDefinitionFilter) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportUpdateResponseDefinitionFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportUpdateResponseDefinitionFilterFilterUnion contains all possible
// properties and values from
// [AttributionReportUpdateResponseDefinitionFilterFilterObject],
// [AttributionReportUpdateResponseDefinitionFilterFilterObject2],
// [AttributionReportUpdateResponseDefinitionFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AttributionReportUpdateResponseDefinitionFilterFilterUnion struct {
	// This field is from variant
	// [AttributionReportUpdateResponseDefinitionFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant
	// [AttributionReportUpdateResponseDefinitionFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant
	// [AttributionReportUpdateResponseDefinitionFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant
	// [AttributionReportUpdateResponseDefinitionFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant
	// [AttributionReportUpdateResponseDefinitionFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant
	// [AttributionReportUpdateResponseDefinitionFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant
	// [AttributionReportUpdateResponseDefinitionFilterFilterObject3].
	Value AttributionReportUpdateResponseDefinitionFilterFilterObject3ValueUnion `json:"value"`
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

func (u AttributionReportUpdateResponseDefinitionFilterFilterUnion) AsAttributionReportUpdateResponseDefinitionFilterFilterObject() (v AttributionReportUpdateResponseDefinitionFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportUpdateResponseDefinitionFilterFilterUnion) AsAttributionReportUpdateResponseDefinitionFilterFilterObject2() (v AttributionReportUpdateResponseDefinitionFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportUpdateResponseDefinitionFilterFilterUnion) AsAttributionReportUpdateResponseDefinitionFilterFilterObject3() (v AttributionReportUpdateResponseDefinitionFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportUpdateResponseDefinitionFilterFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AttributionReportUpdateResponseDefinitionFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportUpdateResponseDefinitionFilterFilterObject struct {
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
func (r AttributionReportUpdateResponseDefinitionFilterFilterObject) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportUpdateResponseDefinitionFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportUpdateResponseDefinitionFilterFilterObject2 struct {
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
func (r AttributionReportUpdateResponseDefinitionFilterFilterObject2) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportUpdateResponseDefinitionFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportUpdateResponseDefinitionFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path"`
	// Any of "string", "number", "boolean".
	Type  string                                                                 `json:"type"`
	Value AttributionReportUpdateResponseDefinitionFilterFilterObject3ValueUnion `json:"value"`
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
func (r AttributionReportUpdateResponseDefinitionFilterFilterObject3) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportUpdateResponseDefinitionFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportUpdateResponseDefinitionFilterFilterObject3ValueUnion contains
// all possible properties and values from [string], [float64], [bool],
// [[]AttributionReportUpdateResponseDefinitionFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAttributionReportUpdateResponseDefinitionFilterFilterObject3ValueArray]
type AttributionReportUpdateResponseDefinitionFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AttributionReportUpdateResponseDefinitionFilterFilterObject3ValueArrayItemUnion]
	// instead of an object.
	OfAttributionReportUpdateResponseDefinitionFilterFilterObject3ValueArray []AttributionReportUpdateResponseDefinitionFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                                     struct {
		OfString                                                                 respjson.Field
		OfFloat                                                                  respjson.Field
		OfBool                                                                   respjson.Field
		OfAttributionReportUpdateResponseDefinitionFilterFilterObject3ValueArray respjson.Field
		raw                                                                      string
	} `json:"-"`
}

func (u AttributionReportUpdateResponseDefinitionFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportUpdateResponseDefinitionFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportUpdateResponseDefinitionFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportUpdateResponseDefinitionFilterFilterObject3ValueUnion) AsAttributionReportUpdateResponseDefinitionFilterFilterObject3ValueArray() (v []AttributionReportUpdateResponseDefinitionFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportUpdateResponseDefinitionFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AttributionReportUpdateResponseDefinitionFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportUpdateResponseDefinitionFilterFilterObject3ValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type AttributionReportUpdateResponseDefinitionFilterFilterObject3ValueArrayItemUnion struct {
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

func (u AttributionReportUpdateResponseDefinitionFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportUpdateResponseDefinitionFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportUpdateResponseDefinitionFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportUpdateResponseDefinitionFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AttributionReportUpdateResponseDefinitionFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default UTC date range: a preset or inclusive from/to calendar dates. Replaces
// the saved context on PATCH. Results cover at most 92 inclusive days.
type AttributionReportUpdateResponseReportContext struct {
	DateRange AttributionReportUpdateResponseReportContextDateRangeUnion `json:"dateRange" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DateRange   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportUpdateResponseReportContext) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportUpdateResponseReportContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportUpdateResponseReportContextDateRangeUnion contains all possible
// properties and values from
// [AttributionReportUpdateResponseReportContextDateRangeObject],
// [AttributionReportUpdateResponseReportContextDateRangeObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AttributionReportUpdateResponseReportContextDateRangeUnion struct {
	// This field is from variant
	// [AttributionReportUpdateResponseReportContextDateRangeObject].
	From string `json:"from"`
	// This field is from variant
	// [AttributionReportUpdateResponseReportContextDateRangeObject].
	To       string `json:"to"`
	TimeZone string `json:"timeZone"`
	// This field is from variant
	// [AttributionReportUpdateResponseReportContextDateRangeObject2].
	Preset string `json:"preset"`
	JSON   struct {
		From     respjson.Field
		To       respjson.Field
		TimeZone respjson.Field
		Preset   respjson.Field
		raw      string
	} `json:"-"`
}

func (u AttributionReportUpdateResponseReportContextDateRangeUnion) AsAttributionReportUpdateResponseReportContextDateRangeObject() (v AttributionReportUpdateResponseReportContextDateRangeObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportUpdateResponseReportContextDateRangeUnion) AsAttributionReportUpdateResponseReportContextDateRangeObject2() (v AttributionReportUpdateResponseReportContextDateRangeObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportUpdateResponseReportContextDateRangeUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AttributionReportUpdateResponseReportContextDateRangeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportUpdateResponseReportContextDateRangeObject struct {
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
func (r AttributionReportUpdateResponseReportContextDateRangeObject) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportUpdateResponseReportContextDateRangeObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportUpdateResponseReportContextDateRangeObject2 struct {
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
func (r AttributionReportUpdateResponseReportContextDateRangeObject2) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportUpdateResponseReportContextDateRangeObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportDeleteResponse struct {
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
func (r AttributionReportDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportDuplicateResponse struct {
	ID              string    `json:"id" api:"required" format:"uuid"`
	CreatedAt       time.Time `json:"createdAt" api:"required" format:"date-time"`
	CreatedByUserID string    `json:"createdByUserId" api:"required"`
	// Version 1 attribution definition: one to five conversion outcomes, each with one
	// to ten unique event names, an optional filter, and an optional numeric event
	// property summed as value. Set the credit model (default last_touch), counting
	// unit (default unique), lookback window of 7, 14, 30, 60, or 90 days (default
	// 30), an optional filter applied to every outcome, and the value currency
	// (default USD). Outcome ids identify outcomes in results. Replaces the saved
	// definition on PATCH.
	Definition AttributionReportDuplicateResponseDefinition `json:"definition" api:"required"`
	// Display name for the saved attribution report.
	Name string `json:"name" api:"required"`
	// Default UTC date range: a preset or inclusive from/to calendar dates. Replaces
	// the saved context on PATCH. Results cover at most 92 inclusive days.
	ReportContext   AttributionReportDuplicateResponseReportContext `json:"reportContext" api:"required"`
	Revision        string                                          `json:"revision" api:"required" format:"uuid"`
	UpdatedAt       time.Time                                       `json:"updatedAt" api:"required" format:"date-time"`
	UpdatedByUserID string                                          `json:"updatedByUserId" api:"required"`
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
		Description     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportDuplicateResponse) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportDuplicateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Version 1 attribution definition: one to five conversion outcomes, each with one
// to ten unique event names, an optional filter, and an optional numeric event
// property summed as value. Set the credit model (default last_touch), counting
// unit (default unique), lookback window of 7, 14, 30, 60, or 90 days (default
// 30), an optional filter applied to every outcome, and the value currency
// (default USD). Outcome ids identify outcomes in results. Replaces the saved
// definition on PATCH.
type AttributionReportDuplicateResponseDefinition struct {
	// Any of "attribution".
	Kind     string                                                `json:"kind" api:"required"`
	Outcomes []AttributionReportDuplicateResponseDefinitionOutcome `json:"outcomes" api:"required"`
	// Any of 1.
	Version float64 `json:"version" api:"required"`
	// Any of "unique", "totals", "sessions".
	Counting string `json:"counting"`
	// Any of "USD", "CAD", "GBP", "EUR", "AUD".
	Currency string                                             `json:"currency"`
	Filter   AttributionReportDuplicateResponseDefinitionFilter `json:"filter"`
	// Any of 7, 14, 30, 60, 90.
	LookbackDays float64 `json:"lookbackDays"`
	// Any of "first_touch", "last_touch", "linear", "u_shaped", "j_shaped",
	// "time_decay".
	Model string `json:"model"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind         respjson.Field
		Outcomes     respjson.Field
		Version      respjson.Field
		Counting     respjson.Field
		Currency     respjson.Field
		Filter       respjson.Field
		LookbackDays respjson.Field
		Model        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportDuplicateResponseDefinition) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportDuplicateResponseDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportDuplicateResponseDefinitionOutcome struct {
	ID         string                                                    `json:"id" api:"required" format:"uuid"`
	EventNames []string                                                  `json:"eventNames" api:"required"`
	Label      string                                                    `json:"label" api:"required"`
	Filter     AttributionReportDuplicateResponseDefinitionOutcomeFilter `json:"filter"`
	Value      AttributionReportDuplicateResponseDefinitionOutcomeValue  `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventNames  respjson.Field
		Label       respjson.Field
		Filter      respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportDuplicateResponseDefinitionOutcome) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportDuplicateResponseDefinitionOutcome) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportDuplicateResponseDefinitionOutcomeFilter struct {
	Filter AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterUnion `json:"filter" api:"required"`
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
func (r AttributionReportDuplicateResponseDefinitionOutcomeFilter) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportDuplicateResponseDefinitionOutcomeFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterUnion contains
// all possible properties and values from
// [AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject],
// [AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject2],
// [AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterUnion struct {
	// This field is from variant
	// [AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant
	// [AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant
	// [AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant
	// [AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant
	// [AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant
	// [AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant
	// [AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3].
	Value AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3ValueUnion `json:"value"`
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

func (u AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterUnion) AsAttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject() (v AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterUnion) AsAttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject2() (v AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterUnion) AsAttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3() (v AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject struct {
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
func (r AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject2 struct {
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
func (r AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject2) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path"`
	// Any of "string", "number", "boolean".
	Type  string                                                                           `json:"type"`
	Value AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3ValueUnion `json:"value"`
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
func (r AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3ValueUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3ValueArray]
type AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion]
	// instead of an object.
	OfAttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3ValueArray []AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                                               struct {
		OfString                                                                           respjson.Field
		OfFloat                                                                            respjson.Field
		OfBool                                                                             respjson.Field
		OfAttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3ValueArray respjson.Field
		raw                                                                                string
	} `json:"-"`
}

func (u AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3ValueUnion) AsAttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3ValueArray() (v []AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion struct {
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

func (u AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AttributionReportDuplicateResponseDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportDuplicateResponseDefinitionOutcomeValue struct {
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
func (r AttributionReportDuplicateResponseDefinitionOutcomeValue) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportDuplicateResponseDefinitionOutcomeValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportDuplicateResponseDefinitionFilter struct {
	Filter AttributionReportDuplicateResponseDefinitionFilterFilterUnion `json:"filter" api:"required"`
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
func (r AttributionReportDuplicateResponseDefinitionFilter) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportDuplicateResponseDefinitionFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportDuplicateResponseDefinitionFilterFilterUnion contains all
// possible properties and values from
// [AttributionReportDuplicateResponseDefinitionFilterFilterObject],
// [AttributionReportDuplicateResponseDefinitionFilterFilterObject2],
// [AttributionReportDuplicateResponseDefinitionFilterFilterObject3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AttributionReportDuplicateResponseDefinitionFilterFilterUnion struct {
	// This field is from variant
	// [AttributionReportDuplicateResponseDefinitionFilterFilterObject].
	Children []any  `json:"children"`
	Kind     string `json:"kind"`
	// This field is from variant
	// [AttributionReportDuplicateResponseDefinitionFilterFilterObject2].
	Child any `json:"child"`
	// This field is from variant
	// [AttributionReportDuplicateResponseDefinitionFilterFilterObject3].
	Operator string `json:"operator"`
	// This field is from variant
	// [AttributionReportDuplicateResponseDefinitionFilterFilterObject3].
	Property string `json:"property"`
	// This field is from variant
	// [AttributionReportDuplicateResponseDefinitionFilterFilterObject3].
	Path []string `json:"path"`
	// This field is from variant
	// [AttributionReportDuplicateResponseDefinitionFilterFilterObject3].
	Type string `json:"type"`
	// This field is from variant
	// [AttributionReportDuplicateResponseDefinitionFilterFilterObject3].
	Value AttributionReportDuplicateResponseDefinitionFilterFilterObject3ValueUnion `json:"value"`
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

func (u AttributionReportDuplicateResponseDefinitionFilterFilterUnion) AsAttributionReportDuplicateResponseDefinitionFilterFilterObject() (v AttributionReportDuplicateResponseDefinitionFilterFilterObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportDuplicateResponseDefinitionFilterFilterUnion) AsAttributionReportDuplicateResponseDefinitionFilterFilterObject2() (v AttributionReportDuplicateResponseDefinitionFilterFilterObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportDuplicateResponseDefinitionFilterFilterUnion) AsAttributionReportDuplicateResponseDefinitionFilterFilterObject3() (v AttributionReportDuplicateResponseDefinitionFilterFilterObject3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportDuplicateResponseDefinitionFilterFilterUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AttributionReportDuplicateResponseDefinitionFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportDuplicateResponseDefinitionFilterFilterObject struct {
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
func (r AttributionReportDuplicateResponseDefinitionFilterFilterObject) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportDuplicateResponseDefinitionFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportDuplicateResponseDefinitionFilterFilterObject2 struct {
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
func (r AttributionReportDuplicateResponseDefinitionFilterFilterObject2) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportDuplicateResponseDefinitionFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportDuplicateResponseDefinitionFilterFilterObject3 struct {
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
	Value AttributionReportDuplicateResponseDefinitionFilterFilterObject3ValueUnion `json:"value"`
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
func (r AttributionReportDuplicateResponseDefinitionFilterFilterObject3) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportDuplicateResponseDefinitionFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportDuplicateResponseDefinitionFilterFilterObject3ValueUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]AttributionReportDuplicateResponseDefinitionFilterFilterObject3ValueArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAttributionReportDuplicateResponseDefinitionFilterFilterObject3ValueArray]
type AttributionReportDuplicateResponseDefinitionFilterFilterObject3ValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]AttributionReportDuplicateResponseDefinitionFilterFilterObject3ValueArrayItemUnion]
	// instead of an object.
	OfAttributionReportDuplicateResponseDefinitionFilterFilterObject3ValueArray []AttributionReportDuplicateResponseDefinitionFilterFilterObject3ValueArrayItemUnion `json:",inline"`
	JSON                                                                        struct {
		OfString                                                                    respjson.Field
		OfFloat                                                                     respjson.Field
		OfBool                                                                      respjson.Field
		OfAttributionReportDuplicateResponseDefinitionFilterFilterObject3ValueArray respjson.Field
		raw                                                                         string
	} `json:"-"`
}

func (u AttributionReportDuplicateResponseDefinitionFilterFilterObject3ValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportDuplicateResponseDefinitionFilterFilterObject3ValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportDuplicateResponseDefinitionFilterFilterObject3ValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportDuplicateResponseDefinitionFilterFilterObject3ValueUnion) AsAttributionReportDuplicateResponseDefinitionFilterFilterObject3ValueArray() (v []AttributionReportDuplicateResponseDefinitionFilterFilterObject3ValueArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportDuplicateResponseDefinitionFilterFilterObject3ValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AttributionReportDuplicateResponseDefinitionFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportDuplicateResponseDefinitionFilterFilterObject3ValueArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type AttributionReportDuplicateResponseDefinitionFilterFilterObject3ValueArrayItemUnion struct {
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

func (u AttributionReportDuplicateResponseDefinitionFilterFilterObject3ValueArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportDuplicateResponseDefinitionFilterFilterObject3ValueArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportDuplicateResponseDefinitionFilterFilterObject3ValueArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportDuplicateResponseDefinitionFilterFilterObject3ValueArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AttributionReportDuplicateResponseDefinitionFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default UTC date range: a preset or inclusive from/to calendar dates. Replaces
// the saved context on PATCH. Results cover at most 92 inclusive days.
type AttributionReportDuplicateResponseReportContext struct {
	DateRange AttributionReportDuplicateResponseReportContextDateRangeUnion `json:"dateRange" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DateRange   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportDuplicateResponseReportContext) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportDuplicateResponseReportContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AttributionReportDuplicateResponseReportContextDateRangeUnion contains all
// possible properties and values from
// [AttributionReportDuplicateResponseReportContextDateRangeObject],
// [AttributionReportDuplicateResponseReportContextDateRangeObject2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AttributionReportDuplicateResponseReportContextDateRangeUnion struct {
	// This field is from variant
	// [AttributionReportDuplicateResponseReportContextDateRangeObject].
	From string `json:"from"`
	// This field is from variant
	// [AttributionReportDuplicateResponseReportContextDateRangeObject].
	To       string `json:"to"`
	TimeZone string `json:"timeZone"`
	// This field is from variant
	// [AttributionReportDuplicateResponseReportContextDateRangeObject2].
	Preset string `json:"preset"`
	JSON   struct {
		From     respjson.Field
		To       respjson.Field
		TimeZone respjson.Field
		Preset   respjson.Field
		raw      string
	} `json:"-"`
}

func (u AttributionReportDuplicateResponseReportContextDateRangeUnion) AsAttributionReportDuplicateResponseReportContextDateRangeObject() (v AttributionReportDuplicateResponseReportContextDateRangeObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AttributionReportDuplicateResponseReportContextDateRangeUnion) AsAttributionReportDuplicateResponseReportContextDateRangeObject2() (v AttributionReportDuplicateResponseReportContextDateRangeObject2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AttributionReportDuplicateResponseReportContextDateRangeUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AttributionReportDuplicateResponseReportContextDateRangeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportDuplicateResponseReportContextDateRangeObject struct {
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
func (r AttributionReportDuplicateResponseReportContextDateRangeObject) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportDuplicateResponseReportContextDateRangeObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportDuplicateResponseReportContextDateRangeObject2 struct {
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
func (r AttributionReportDuplicateResponseReportContextDateRangeObject2) RawJSON() string {
	return r.JSON.raw
}
func (r *AttributionReportDuplicateResponseReportContextDateRangeObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportPreviewResponse struct {
	ChannelRulesVersion int64     `json:"channelRulesVersion" api:"required"`
	ComputedAt          time.Time `json:"computedAt" api:"required" format:"date-time"`
	// Any of "unique", "totals", "sessions".
	Counting AttributionReportPreviewResponseCounting `json:"counting" api:"required"`
	From     string                                   `json:"from" api:"required"`
	// Any of "day", "week".
	Interval AttributionReportPreviewResponseInterval `json:"interval" api:"required"`
	// True when more breakdown rows matched than the response returns; the total row
	// is always complete.
	IsTruncated bool `json:"isTruncated" api:"required"`
	// Time-series credit per bucket, period, and outcome; channel is null for the
	// all-channel series.
	Points []AttributionReportPreviewResponsePoint `json:"points" api:"required"`
	// Inclusive comparison start date; null without a comparison.
	PreviousFrom string `json:"previousFrom" api:"required"`
	// Inclusive comparison end date; null without a comparison.
	PreviousTo string `json:"previousTo" api:"required"`
	// The total row followed by channel, source/medium, campaign, and term breakdown
	// rows, ordered by current-period credit under the selected model.
	Rows []AttributionReportPreviewResponseRow `json:"rows" api:"required"`
	To   string                                `json:"to" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChannelRulesVersion respjson.Field
		ComputedAt          respjson.Field
		Counting            respjson.Field
		From                respjson.Field
		Interval            respjson.Field
		IsTruncated         respjson.Field
		Points              respjson.Field
		PreviousFrom        respjson.Field
		PreviousTo          respjson.Field
		Rows                respjson.Field
		To                  respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportPreviewResponse) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportPreviewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportPreviewResponseCounting string

const (
	AttributionReportPreviewResponseCountingUnique   AttributionReportPreviewResponseCounting = "unique"
	AttributionReportPreviewResponseCountingTotals   AttributionReportPreviewResponseCounting = "totals"
	AttributionReportPreviewResponseCountingSessions AttributionReportPreviewResponseCounting = "sessions"
)

type AttributionReportPreviewResponseInterval string

const (
	AttributionReportPreviewResponseIntervalDay  AttributionReportPreviewResponseInterval = "day"
	AttributionReportPreviewResponseIntervalWeek AttributionReportPreviewResponseInterval = "week"
)

type AttributionReportPreviewResponsePoint struct {
	Bucket  string `json:"bucket" api:"required"`
	Channel string `json:"channel" api:"required"`
	// Credit under every model, keyed by model name, for this bucket.
	Credit    AttributionReportPreviewResponsePointCredit `json:"credit" api:"required"`
	OutcomeID string                                      `json:"outcomeId" api:"required" format:"uuid"`
	// Any of "current", "previous".
	Period string `json:"period" api:"required"`
	// Attributed value under every model for this bucket; null when the outcome has no
	// value property.
	Value AttributionReportPreviewResponsePointValue `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bucket      respjson.Field
		Channel     respjson.Field
		Credit      respjson.Field
		OutcomeID   respjson.Field
		Period      respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportPreviewResponsePoint) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportPreviewResponsePoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit under every model, keyed by model name, for this bucket.
type AttributionReportPreviewResponsePointCredit struct {
	FirstTouch float64 `json:"first_touch" api:"required"`
	JShaped    float64 `json:"j_shaped" api:"required"`
	LastTouch  float64 `json:"last_touch" api:"required"`
	Linear     float64 `json:"linear" api:"required"`
	TimeDecay  float64 `json:"time_decay" api:"required"`
	UShaped    float64 `json:"u_shaped" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstTouch  respjson.Field
		JShaped     respjson.Field
		LastTouch   respjson.Field
		Linear      respjson.Field
		TimeDecay   respjson.Field
		UShaped     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportPreviewResponsePointCredit) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportPreviewResponsePointCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Attributed value under every model for this bucket; null when the outcome has no
// value property.
type AttributionReportPreviewResponsePointValue struct {
	FirstTouch float64 `json:"first_touch" api:"required"`
	JShaped    float64 `json:"j_shaped" api:"required"`
	LastTouch  float64 `json:"last_touch" api:"required"`
	Linear     float64 `json:"linear" api:"required"`
	TimeDecay  float64 `json:"time_decay" api:"required"`
	UShaped    float64 `json:"u_shaped" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstTouch  respjson.Field
		JShaped     respjson.Field
		LastTouch   respjson.Field
		Linear      respjson.Field
		TimeDecay   respjson.Field
		UShaped     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportPreviewResponsePointValue) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportPreviewResponsePointValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportPreviewResponseRow struct {
	Campaign string `json:"campaign" api:"required"`
	// One cell per outcome and period, current period first, in outcome order.
	Cells   []AttributionReportPreviewResponseRowCell `json:"cells" api:"required"`
	Channel string                                    `json:"channel" api:"required"`
	// Any of "total", "channel", "source_medium", "campaign", "term".
	Grain        string `json:"grain" api:"required"`
	SourceMedium string `json:"sourceMedium" api:"required"`
	Term         string `json:"term" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Campaign     respjson.Field
		Cells        respjson.Field
		Channel      respjson.Field
		Grain        respjson.Field
		SourceMedium respjson.Field
		Term         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportPreviewResponseRow) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportPreviewResponseRow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportPreviewResponseRowCell struct {
	AssistedVisitors int64 `json:"assistedVisitors" api:"required"`
	// Credit under every model, keyed by model name, in the definition counting unit.
	// Compare models without recomputing.
	Credit              AttributionReportPreviewResponseRowCellCredit `json:"credit" api:"required"`
	CreditedUnits       int64                                         `json:"creditedUnits" api:"required"`
	MedianDaysToConvert float64                                       `json:"medianDaysToConvert" api:"required"`
	OutcomeID           string                                        `json:"outcomeId" api:"required" format:"uuid"`
	// Any of "current", "previous".
	Period string `json:"period" api:"required"`
	// Attributed value under every model, in the definition currency; null when the
	// outcome has no value property.
	Value AttributionReportPreviewResponseRowCellValue `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssistedVisitors    respjson.Field
		Credit              respjson.Field
		CreditedUnits       respjson.Field
		MedianDaysToConvert respjson.Field
		OutcomeID           respjson.Field
		Period              respjson.Field
		Value               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportPreviewResponseRowCell) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportPreviewResponseRowCell) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Credit under every model, keyed by model name, in the definition counting unit.
// Compare models without recomputing.
type AttributionReportPreviewResponseRowCellCredit struct {
	FirstTouch float64 `json:"first_touch" api:"required"`
	JShaped    float64 `json:"j_shaped" api:"required"`
	LastTouch  float64 `json:"last_touch" api:"required"`
	Linear     float64 `json:"linear" api:"required"`
	TimeDecay  float64 `json:"time_decay" api:"required"`
	UShaped    float64 `json:"u_shaped" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstTouch  respjson.Field
		JShaped     respjson.Field
		LastTouch   respjson.Field
		Linear      respjson.Field
		TimeDecay   respjson.Field
		UShaped     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportPreviewResponseRowCellCredit) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportPreviewResponseRowCellCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Attributed value under every model, in the definition currency; null when the
// outcome has no value property.
type AttributionReportPreviewResponseRowCellValue struct {
	FirstTouch float64 `json:"first_touch" api:"required"`
	JShaped    float64 `json:"j_shaped" api:"required"`
	LastTouch  float64 `json:"last_touch" api:"required"`
	Linear     float64 `json:"linear" api:"required"`
	TimeDecay  float64 `json:"time_decay" api:"required"`
	UShaped    float64 `json:"u_shaped" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstTouch  respjson.Field
		JShaped     respjson.Field
		LastTouch   respjson.Field
		Linear      respjson.Field
		TimeDecay   respjson.Field
		UShaped     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportPreviewResponseRowCellValue) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportPreviewResponseRowCellValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportVisitorsResponse struct {
	// Visitors ordered by latest matching conversion, newest first, then visitor id.
	Entities   []AttributionReportVisitorsResponseEntity   `json:"entities" api:"required"`
	Pagination AttributionReportVisitorsResponsePagination `json:"pagination" api:"required"`
	// Visitors behind the selection.
	TotalCount int64 `json:"totalCount" api:"required"`
	// Counted units behind the selection. Equals the first- or last-touch credit for
	// those roles and the row's creditedUnits for any.
	TotalUnits int64 `json:"totalUnits" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		Pagination  respjson.Field
		TotalCount  respjson.Field
		TotalUnits  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportVisitorsResponse) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportVisitorsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportVisitorsResponseEntity struct {
	// When the latest matching conversion happened (UTC).
	ConvertedAt time.Time `json:"convertedAt" api:"required" format:"date-time"`
	// Credit this visitor gives the selected row under the definition's model.
	Credit float64 `json:"credit" api:"required"`
	// The credited first touch of the latest matching conversion; null when it had no
	// eligible touch.
	FirstTouch AttributionReportVisitorsResponseEntityFirstTouch `json:"firstTouch" api:"required"`
	// The credited last touch of the latest matching conversion; null when it had no
	// eligible touch.
	LastTouch AttributionReportVisitorsResponseEntityLastTouch `json:"lastTouch" api:"required"`
	// Up to three recorded sessions from the matching lookback windows, newest first,
	// within session replay retention.
	LatestReplays []AttributionReportVisitorsResponseEntityLatestReplay `json:"latestReplays" api:"required"`
	// Sessions in the latest matching conversion's lookback window.
	SessionCount int64 `json:"sessionCount" api:"required"`
	// Matching conversions (totals counting) or converting sessions (sessions
	// counting); always 1 under unique counting.
	Units int64 `json:"units" api:"required"`
	// Value this visitor credits to the selected row under the definition's model;
	// null when the outcome has no value property.
	Value     float64 `json:"value" api:"required"`
	VisitorID string  `json:"visitorId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConvertedAt   respjson.Field
		Credit        respjson.Field
		FirstTouch    respjson.Field
		LastTouch     respjson.Field
		LatestReplays respjson.Field
		SessionCount  respjson.Field
		Units         respjson.Field
		Value         respjson.Field
		VisitorID     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportVisitorsResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportVisitorsResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The credited first touch of the latest matching conversion; null when it had no
// eligible touch.
type AttributionReportVisitorsResponseEntityFirstTouch struct {
	At           time.Time `json:"at" api:"required" format:"date-time"`
	Campaign     string    `json:"campaign" api:"required"`
	Channel      string    `json:"channel" api:"required"`
	SessionID    string    `json:"sessionId" api:"required"`
	SourceMedium string    `json:"sourceMedium" api:"required"`
	Term         string    `json:"term" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		At           respjson.Field
		Campaign     respjson.Field
		Channel      respjson.Field
		SessionID    respjson.Field
		SourceMedium respjson.Field
		Term         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportVisitorsResponseEntityFirstTouch) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportVisitorsResponseEntityFirstTouch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The credited last touch of the latest matching conversion; null when it had no
// eligible touch.
type AttributionReportVisitorsResponseEntityLastTouch struct {
	At           time.Time `json:"at" api:"required" format:"date-time"`
	Campaign     string    `json:"campaign" api:"required"`
	Channel      string    `json:"channel" api:"required"`
	SessionID    string    `json:"sessionId" api:"required"`
	SourceMedium string    `json:"sourceMedium" api:"required"`
	Term         string    `json:"term" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		At           respjson.Field
		Campaign     respjson.Field
		Channel      respjson.Field
		SessionID    respjson.Field
		SourceMedium respjson.Field
		Term         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportVisitorsResponseEntityLastTouch) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportVisitorsResponseEntityLastTouch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportVisitorsResponseEntityLatestReplay struct {
	Date      string    `json:"date" api:"required"`
	SessionID string    `json:"sessionId" api:"required"`
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		SessionID   respjson.Field
		StartedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportVisitorsResponseEntityLatestReplay) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportVisitorsResponseEntityLatestReplay) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportVisitorsResponsePagination struct {
	HasMore    bool   `json:"hasMore" api:"required"`
	NextCursor string `json:"nextCursor" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportVisitorsResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportVisitorsResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportInsightsResponse struct {
	ComputedAt time.Time `json:"computedAt" api:"required" format:"date-time"`
	// Any of "unique", "totals", "sessions".
	Counting AttributionReportInsightsResponseCounting `json:"counting" api:"required"`
	From     string                                    `json:"from" api:"required"`
	// At most 8 findings, most notable first, one per row. Aggregates only; no visitor
	// identifiers.
	Insights []AttributionReportInsightsResponseInsight `json:"insights" api:"required"`
	// Any of "first_touch", "last_touch", "linear", "u_shaped", "j_shaped",
	// "time_decay".
	Model AttributionReportInsightsResponseModel `json:"model" api:"required"`
	// A value is new when it had no sessions in this many days before from.
	NewValueLookbackDays int64 `json:"newValueLookbackDays" api:"required"`
	// Inclusive comparison start date; null without a comparison.
	PreviousFrom string `json:"previousFrom" api:"required"`
	// Inclusive comparison end date; null without a comparison.
	PreviousTo string `json:"previousTo" api:"required"`
	To         string `json:"to" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ComputedAt           respjson.Field
		Counting             respjson.Field
		From                 respjson.Field
		Insights             respjson.Field
		Model                respjson.Field
		NewValueLookbackDays respjson.Field
		PreviousFrom         respjson.Field
		PreviousTo           respjson.Field
		To                   respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportInsightsResponse) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportInsightsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportInsightsResponseCounting string

const (
	AttributionReportInsightsResponseCountingUnique   AttributionReportInsightsResponseCounting = "unique"
	AttributionReportInsightsResponseCountingTotals   AttributionReportInsightsResponseCounting = "totals"
	AttributionReportInsightsResponseCountingSessions AttributionReportInsightsResponseCounting = "sessions"
)

type AttributionReportInsightsResponseInsight struct {
	// Stable identifier for the same input, made of the kind, outcome, and row.
	ID string `json:"id" api:"required"`
	// The number current is compared with; see baselineKind.
	Baseline float64 `json:"baseline" api:"required"`
	// previous_period: baseline is the same model in the comparison period.
	// other_model: baseline is baselineModel in the current period. lookback: baseline
	// is the number of sessions with the value before the range, always 0.
	//
	// Any of "previous_period", "other_model", "lookback".
	BaselineKind string `json:"baselineKind" api:"required"`
	// Any of "first_touch", "last_touch", "linear", "u_shaped", "j_shaped",
	// "time_decay".
	BaselineModel string `json:"baselineModel" api:"required"`
	// The row's share of the outcome's baseline credit; null for new values.
	BaselineShare float64 `json:"baselineShare" api:"required"`
	// Relative change from baseline; null when the baseline is 0.
	Change float64 `json:"change" api:"required"`
	// Credit in the current period under model, in the definition counting unit.
	Current float64 `json:"current" api:"required"`
	// The row's share of the outcome's current credit; null when the outcome has none.
	CurrentShare float64 `json:"currentShare" api:"required"`
	// current minus baseline.
	Delta float64 `json:"delta" api:"required"`
	// Child rows that together explain most of a move, largest first.
	Drivers []string `json:"drivers" api:"required"`
	// Always descriptive: findings are ranked by effect size above minimum volumes and
	// are not tested for significance.
	//
	// Any of "descriptive".
	Evidence string `json:"evidence" api:"required"`
	// First session date with the new value; new_value only.
	FirstSeen string `json:"firstSeen" api:"required"`
	// new_value: a source / medium, campaign, or term with credited conversions and no
	// sessions in the lookback before the range. mover_up / mover_down: the largest
	// share-adjusted credit changes against the comparison period. model_disagreement:
	// a channel or source / medium credited at least twice as much by last touch as by
	// first touch, or the reverse. went_quiet: a row that had credit in the comparison
	// period and almost none now.
	//
	// Any of "new_value", "mover_up", "mover_down", "model_disagreement",
	// "went_quiet".
	Kind string `json:"kind" api:"required"`
	// Credit model of current.
	//
	// Any of "first_touch", "last_touch", "linear", "u_shaped", "j_shaped",
	// "time_decay".
	Model string `json:"model" api:"required"`
	// Relative change of the outcome's total credit, so a row's move can be read
	// against the account's.
	OutcomeChange float64 `json:"outcomeChange" api:"required"`
	// Outcome id from the definition the numbers belong to.
	OutcomeID string `json:"outcomeId" api:"required" format:"uuid"`
	// Counted units with a touch on this row in the comparison period; null without
	// one.
	PreviousSampleUnits int64 `json:"previousSampleUnits" api:"required"`
	// The breakdown row the finding is about; matches a preview row by grain, channel,
	// sourceMedium, campaign, and term.
	Row AttributionReportInsightsResponseInsightRow `json:"row" api:"required"`
	// Sessions with the new value in the range; new_value only.
	SampleSessions int64 `json:"sampleSessions" api:"required"`
	// Counted units with a touch on this row in the current period.
	SampleUnits int64 `json:"sampleUnits" api:"required"`
	// Ranking score from 0 to 1, weighted by kind.
	Score float64 `json:"score" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Baseline            respjson.Field
		BaselineKind        respjson.Field
		BaselineModel       respjson.Field
		BaselineShare       respjson.Field
		Change              respjson.Field
		Current             respjson.Field
		CurrentShare        respjson.Field
		Delta               respjson.Field
		Drivers             respjson.Field
		Evidence            respjson.Field
		FirstSeen           respjson.Field
		Kind                respjson.Field
		Model               respjson.Field
		OutcomeChange       respjson.Field
		OutcomeID           respjson.Field
		PreviousSampleUnits respjson.Field
		Row                 respjson.Field
		SampleSessions      respjson.Field
		SampleUnits         respjson.Field
		Score               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportInsightsResponseInsight) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportInsightsResponseInsight) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The breakdown row the finding is about; matches a preview row by grain, channel,
// sourceMedium, campaign, and term.
type AttributionReportInsightsResponseInsightRow struct {
	Campaign string `json:"campaign" api:"required"`
	Channel  string `json:"channel" api:"required"`
	// Any of "channel", "source_medium", "campaign", "term".
	Grain        string `json:"grain" api:"required"`
	SourceMedium string `json:"sourceMedium" api:"required"`
	Term         string `json:"term" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Campaign     respjson.Field
		Channel      respjson.Field
		Grain        respjson.Field
		SourceMedium respjson.Field
		Term         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributionReportInsightsResponseInsightRow) RawJSON() string { return r.JSON.raw }
func (r *AttributionReportInsightsResponseInsightRow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportInsightsResponseModel string

const (
	AttributionReportInsightsResponseModelFirstTouch AttributionReportInsightsResponseModel = "first_touch"
	AttributionReportInsightsResponseModelLastTouch  AttributionReportInsightsResponseModel = "last_touch"
	AttributionReportInsightsResponseModelLinear     AttributionReportInsightsResponseModel = "linear"
	AttributionReportInsightsResponseModelUShaped    AttributionReportInsightsResponseModel = "u_shaped"
	AttributionReportInsightsResponseModelJShaped    AttributionReportInsightsResponseModel = "j_shaped"
	AttributionReportInsightsResponseModelTimeDecay  AttributionReportInsightsResponseModel = "time_decay"
)

type AttributionReportListParams struct {
	// Opaque pagination.nextCursor from the previous page, scoped to the account and
	// search text. Continue while pagination.hasMore is true; changing the search
	// requires starting without a cursor.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum attribution reports to return. Defaults to 20; values above 100 are
	// capped at 100. Each page scans at most 100 reports in stable ID order, so pages
	// may be short even when more results exist.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Case-insensitive substring search on attribution report names.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AttributionReportListParams]'s query parameters as
// `url.Values`.
func (r AttributionReportListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AttributionReportNewParams struct {
	// Version 1 attribution definition: one to five conversion outcomes, each with one
	// to ten unique event names, an optional filter, and an optional numeric event
	// property summed as value. Set the credit model (default last_touch), counting
	// unit (default unique), lookback window of 7, 14, 30, 60, or 90 days (default
	// 30), an optional filter applied to every outcome, and the value currency
	// (default USD). Outcome ids identify outcomes in results. Replaces the saved
	// definition on PATCH.
	Definition AttributionReportNewParamsDefinition `json:"definition,omitzero" api:"required"`
	// Display name for the saved attribution report.
	Name string `json:"name" api:"required"`
	// Default UTC date range: a preset or inclusive from/to calendar dates. Replaces
	// the saved context on PATCH. Results cover at most 92 inclusive days.
	ReportContext AttributionReportNewParamsReportContext `json:"reportContext,omitzero" api:"required"`
	// Optional description; null clears it.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r AttributionReportNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AttributionReportNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributionReportNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Version 1 attribution definition: one to five conversion outcomes, each with one
// to ten unique event names, an optional filter, and an optional numeric event
// property summed as value. Set the credit model (default last_touch), counting
// unit (default unique), lookback window of 7, 14, 30, 60, or 90 days (default
// 30), an optional filter applied to every outcome, and the value currency
// (default USD). Outcome ids identify outcomes in results. Replaces the saved
// definition on PATCH.
//
// The properties Kind, Outcomes, Version are required.
type AttributionReportNewParamsDefinition struct {
	// Any of "attribution".
	Kind     string                                        `json:"kind,omitzero" api:"required"`
	Outcomes []AttributionReportNewParamsDefinitionOutcome `json:"outcomes,omitzero" api:"required"`
	// Any of 1.
	Version float64 `json:"version,omitzero" api:"required"`
	// Any of "unique", "totals", "sessions".
	Counting string `json:"counting,omitzero"`
	// Any of "USD", "CAD", "GBP", "EUR", "AUD".
	Currency string                                     `json:"currency,omitzero"`
	Filter   AttributionReportNewParamsDefinitionFilter `json:"filter,omitzero"`
	// Any of 7, 14, 30, 60, 90.
	LookbackDays float64 `json:"lookbackDays,omitzero"`
	// Any of "first_touch", "last_touch", "linear", "u_shaped", "j_shaped",
	// "time_decay".
	Model string `json:"model,omitzero"`
	paramObj
}

func (r AttributionReportNewParamsDefinition) MarshalJSON() (data []byte, err error) {
	type shadow AttributionReportNewParamsDefinition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributionReportNewParamsDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AttributionReportNewParamsDefinition](
		"kind", "attribution",
	)
	apijson.RegisterFieldValidator[AttributionReportNewParamsDefinition](
		"version", 1,
	)
	apijson.RegisterFieldValidator[AttributionReportNewParamsDefinition](
		"counting", "unique", "totals", "sessions",
	)
	apijson.RegisterFieldValidator[AttributionReportNewParamsDefinition](
		"currency", "USD", "CAD", "GBP", "EUR", "AUD",
	)
	apijson.RegisterFieldValidator[AttributionReportNewParamsDefinition](
		"lookbackDays", 7, 14, 30, 60, 90,
	)
	apijson.RegisterFieldValidator[AttributionReportNewParamsDefinition](
		"model", "first_touch", "last_touch", "linear", "u_shaped", "j_shaped", "time_decay",
	)
}

// The properties ID, EventNames, Label are required.
type AttributionReportNewParamsDefinitionOutcome struct {
	ID         string                                            `json:"id" api:"required" format:"uuid"`
	EventNames []string                                          `json:"eventNames,omitzero" api:"required"`
	Label      string                                            `json:"label" api:"required"`
	Filter     AttributionReportNewParamsDefinitionOutcomeFilter `json:"filter,omitzero"`
	Value      AttributionReportNewParamsDefinitionOutcomeValue  `json:"value,omitzero"`
	paramObj
}

func (r AttributionReportNewParamsDefinitionOutcome) MarshalJSON() (data []byte, err error) {
	type shadow AttributionReportNewParamsDefinitionOutcome
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributionReportNewParamsDefinitionOutcome) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Filter, Version are required.
type AttributionReportNewParamsDefinitionOutcomeFilter struct {
	Filter AttributionReportNewParamsDefinitionOutcomeFilterFilterUnion `json:"filter,omitzero" api:"required"`
	// Any of 1.
	Version float64 `json:"version,omitzero" api:"required"`
	paramObj
}

func (r AttributionReportNewParamsDefinitionOutcomeFilter) MarshalJSON() (data []byte, err error) {
	type shadow AttributionReportNewParamsDefinitionOutcomeFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributionReportNewParamsDefinitionOutcomeFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AttributionReportNewParamsDefinitionOutcomeFilter](
		"version", 1,
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AttributionReportNewParamsDefinitionOutcomeFilterFilterUnion struct {
	OfAttributionReportNewsDefinitionOutcomeFilterFilterObject  *AttributionReportNewParamsDefinitionOutcomeFilterFilterObject  `json:",omitzero,inline"`
	OfAttributionReportNewsDefinitionOutcomeFilterFilterObject2 *AttributionReportNewParamsDefinitionOutcomeFilterFilterObject2 `json:",omitzero,inline"`
	OfAttributionReportNewsDefinitionOutcomeFilterFilterObject3 *AttributionReportNewParamsDefinitionOutcomeFilterFilterObject3 `json:",omitzero,inline"`
	paramUnion
}

func (u AttributionReportNewParamsDefinitionOutcomeFilterFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAttributionReportNewsDefinitionOutcomeFilterFilterObject, u.OfAttributionReportNewsDefinitionOutcomeFilterFilterObject2, u.OfAttributionReportNewsDefinitionOutcomeFilterFilterObject3)
}
func (u *AttributionReportNewParamsDefinitionOutcomeFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Children, Kind are required.
type AttributionReportNewParamsDefinitionOutcomeFilterFilterObject struct {
	Children []any `json:"children,omitzero" api:"required"`
	// Any of "and", "or".
	Kind string `json:"kind,omitzero" api:"required"`
	paramObj
}

func (r AttributionReportNewParamsDefinitionOutcomeFilterFilterObject) MarshalJSON() (data []byte, err error) {
	type shadow AttributionReportNewParamsDefinitionOutcomeFilterFilterObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributionReportNewParamsDefinitionOutcomeFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AttributionReportNewParamsDefinitionOutcomeFilterFilterObject](
		"kind", "and", "or",
	)
}

// The properties Child, Kind are required.
type AttributionReportNewParamsDefinitionOutcomeFilterFilterObject2 struct {
	Child any `json:"child,omitzero" api:"required"`
	// Any of "not".
	Kind string `json:"kind,omitzero" api:"required"`
	paramObj
}

func (r AttributionReportNewParamsDefinitionOutcomeFilterFilterObject2) MarshalJSON() (data []byte, err error) {
	type shadow AttributionReportNewParamsDefinitionOutcomeFilterFilterObject2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributionReportNewParamsDefinitionOutcomeFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AttributionReportNewParamsDefinitionOutcomeFilterFilterObject2](
		"kind", "not",
	)
}

// The properties Kind, Operator, Property are required.
type AttributionReportNewParamsDefinitionOutcomeFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind,omitzero" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator,omitzero" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path,omitzero"`
	// Any of "string", "number", "boolean".
	Type  string                                                                   `json:"type,omitzero"`
	Value AttributionReportNewParamsDefinitionOutcomeFilterFilterObject3ValueUnion `json:"value,omitzero"`
	paramObj
}

func (r AttributionReportNewParamsDefinitionOutcomeFilterFilterObject3) MarshalJSON() (data []byte, err error) {
	type shadow AttributionReportNewParamsDefinitionOutcomeFilterFilterObject3
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributionReportNewParamsDefinitionOutcomeFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AttributionReportNewParamsDefinitionOutcomeFilterFilterObject3](
		"kind", "predicate",
	)
	apijson.RegisterFieldValidator[AttributionReportNewParamsDefinitionOutcomeFilterFilterObject3](
		"operator", "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains", "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in", "exists", "missing", "is_null",
	)
	apijson.RegisterFieldValidator[AttributionReportNewParamsDefinitionOutcomeFilterFilterObject3](
		"type", "string", "number", "boolean",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AttributionReportNewParamsDefinitionOutcomeFilterFilterObject3ValueUnion struct {
	OfString                                                              param.Opt[string]                                                                   `json:",omitzero,inline"`
	OfFloat                                                               param.Opt[float64]                                                                  `json:",omitzero,inline"`
	OfBool                                                                param.Opt[bool]                                                                     `json:",omitzero,inline"`
	OfAttributionReportNewsDefinitionOutcomeFilterFilterObject3ValueArray []AttributionReportNewParamsDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u AttributionReportNewParamsDefinitionOutcomeFilterFilterObject3ValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool, u.OfAttributionReportNewsDefinitionOutcomeFilterFilterObject3ValueArray)
}
func (u *AttributionReportNewParamsDefinitionOutcomeFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AttributionReportNewParamsDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u AttributionReportNewParamsDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *AttributionReportNewParamsDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Path, Property are required.
type AttributionReportNewParamsDefinitionOutcomeValue struct {
	Path []string `json:"path,omitzero" api:"required"`
	// Any of "event.properties".
	Property string `json:"property,omitzero" api:"required"`
	// Any of "number".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r AttributionReportNewParamsDefinitionOutcomeValue) MarshalJSON() (data []byte, err error) {
	type shadow AttributionReportNewParamsDefinitionOutcomeValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributionReportNewParamsDefinitionOutcomeValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AttributionReportNewParamsDefinitionOutcomeValue](
		"property", "event.properties",
	)
	apijson.RegisterFieldValidator[AttributionReportNewParamsDefinitionOutcomeValue](
		"type", "number",
	)
}

// The properties Filter, Version are required.
type AttributionReportNewParamsDefinitionFilter struct {
	Filter AttributionReportNewParamsDefinitionFilterFilterUnion `json:"filter,omitzero" api:"required"`
	// Any of 1.
	Version float64 `json:"version,omitzero" api:"required"`
	paramObj
}

func (r AttributionReportNewParamsDefinitionFilter) MarshalJSON() (data []byte, err error) {
	type shadow AttributionReportNewParamsDefinitionFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributionReportNewParamsDefinitionFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AttributionReportNewParamsDefinitionFilter](
		"version", 1,
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AttributionReportNewParamsDefinitionFilterFilterUnion struct {
	OfAttributionReportNewsDefinitionFilterFilterObject  *AttributionReportNewParamsDefinitionFilterFilterObject  `json:",omitzero,inline"`
	OfAttributionReportNewsDefinitionFilterFilterObject2 *AttributionReportNewParamsDefinitionFilterFilterObject2 `json:",omitzero,inline"`
	OfAttributionReportNewsDefinitionFilterFilterObject3 *AttributionReportNewParamsDefinitionFilterFilterObject3 `json:",omitzero,inline"`
	paramUnion
}

func (u AttributionReportNewParamsDefinitionFilterFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAttributionReportNewsDefinitionFilterFilterObject, u.OfAttributionReportNewsDefinitionFilterFilterObject2, u.OfAttributionReportNewsDefinitionFilterFilterObject3)
}
func (u *AttributionReportNewParamsDefinitionFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Children, Kind are required.
type AttributionReportNewParamsDefinitionFilterFilterObject struct {
	Children []any `json:"children,omitzero" api:"required"`
	// Any of "and", "or".
	Kind string `json:"kind,omitzero" api:"required"`
	paramObj
}

func (r AttributionReportNewParamsDefinitionFilterFilterObject) MarshalJSON() (data []byte, err error) {
	type shadow AttributionReportNewParamsDefinitionFilterFilterObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributionReportNewParamsDefinitionFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AttributionReportNewParamsDefinitionFilterFilterObject](
		"kind", "and", "or",
	)
}

// The properties Child, Kind are required.
type AttributionReportNewParamsDefinitionFilterFilterObject2 struct {
	Child any `json:"child,omitzero" api:"required"`
	// Any of "not".
	Kind string `json:"kind,omitzero" api:"required"`
	paramObj
}

func (r AttributionReportNewParamsDefinitionFilterFilterObject2) MarshalJSON() (data []byte, err error) {
	type shadow AttributionReportNewParamsDefinitionFilterFilterObject2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributionReportNewParamsDefinitionFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AttributionReportNewParamsDefinitionFilterFilterObject2](
		"kind", "not",
	)
}

// The properties Kind, Operator, Property are required.
type AttributionReportNewParamsDefinitionFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind,omitzero" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator,omitzero" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path,omitzero"`
	// Any of "string", "number", "boolean".
	Type  string                                                            `json:"type,omitzero"`
	Value AttributionReportNewParamsDefinitionFilterFilterObject3ValueUnion `json:"value,omitzero"`
	paramObj
}

func (r AttributionReportNewParamsDefinitionFilterFilterObject3) MarshalJSON() (data []byte, err error) {
	type shadow AttributionReportNewParamsDefinitionFilterFilterObject3
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributionReportNewParamsDefinitionFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AttributionReportNewParamsDefinitionFilterFilterObject3](
		"kind", "predicate",
	)
	apijson.RegisterFieldValidator[AttributionReportNewParamsDefinitionFilterFilterObject3](
		"operator", "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains", "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in", "exists", "missing", "is_null",
	)
	apijson.RegisterFieldValidator[AttributionReportNewParamsDefinitionFilterFilterObject3](
		"type", "string", "number", "boolean",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AttributionReportNewParamsDefinitionFilterFilterObject3ValueUnion struct {
	OfString                                                       param.Opt[string]                                                            `json:",omitzero,inline"`
	OfFloat                                                        param.Opt[float64]                                                           `json:",omitzero,inline"`
	OfBool                                                         param.Opt[bool]                                                              `json:",omitzero,inline"`
	OfAttributionReportNewsDefinitionFilterFilterObject3ValueArray []AttributionReportNewParamsDefinitionFilterFilterObject3ValueArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u AttributionReportNewParamsDefinitionFilterFilterObject3ValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool, u.OfAttributionReportNewsDefinitionFilterFilterObject3ValueArray)
}
func (u *AttributionReportNewParamsDefinitionFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AttributionReportNewParamsDefinitionFilterFilterObject3ValueArrayItemUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u AttributionReportNewParamsDefinitionFilterFilterObject3ValueArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *AttributionReportNewParamsDefinitionFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Default UTC date range: a preset or inclusive from/to calendar dates. Replaces
// the saved context on PATCH. Results cover at most 92 inclusive days.
//
// The property DateRange is required.
type AttributionReportNewParamsReportContext struct {
	DateRange AttributionReportNewParamsReportContextDateRangeUnion `json:"dateRange,omitzero" api:"required"`
	paramObj
}

func (r AttributionReportNewParamsReportContext) MarshalJSON() (data []byte, err error) {
	type shadow AttributionReportNewParamsReportContext
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributionReportNewParamsReportContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AttributionReportNewParamsReportContextDateRangeUnion struct {
	OfAttributionReportNewsReportContextDateRangeObject  *AttributionReportNewParamsReportContextDateRangeObject  `json:",omitzero,inline"`
	OfAttributionReportNewsReportContextDateRangeObject2 *AttributionReportNewParamsReportContextDateRangeObject2 `json:",omitzero,inline"`
	paramUnion
}

func (u AttributionReportNewParamsReportContextDateRangeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAttributionReportNewsReportContextDateRangeObject, u.OfAttributionReportNewsReportContextDateRangeObject2)
}
func (u *AttributionReportNewParamsReportContextDateRangeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties From, To are required.
type AttributionReportNewParamsReportContextDateRangeObject struct {
	From string `json:"from" api:"required"`
	To   string `json:"to" api:"required"`
	// Any of "UTC".
	TimeZone string `json:"timeZone,omitzero"`
	paramObj
}

func (r AttributionReportNewParamsReportContextDateRangeObject) MarshalJSON() (data []byte, err error) {
	type shadow AttributionReportNewParamsReportContextDateRangeObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributionReportNewParamsReportContextDateRangeObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AttributionReportNewParamsReportContextDateRangeObject](
		"timeZone", "UTC",
	)
}

// The property Preset is required.
type AttributionReportNewParamsReportContextDateRangeObject2 struct {
	// Any of "today", "yesterday", "last_7_days", "last_30_days", "last_90_days",
	// "last_12_months", "week_to_date", "month_to_date", "year_to_date".
	Preset string `json:"preset,omitzero" api:"required"`
	// Any of "UTC".
	TimeZone string `json:"timeZone,omitzero"`
	paramObj
}

func (r AttributionReportNewParamsReportContextDateRangeObject2) MarshalJSON() (data []byte, err error) {
	type shadow AttributionReportNewParamsReportContextDateRangeObject2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributionReportNewParamsReportContextDateRangeObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AttributionReportNewParamsReportContextDateRangeObject2](
		"preset", "today", "yesterday", "last_7_days", "last_30_days", "last_90_days", "last_12_months", "week_to_date", "month_to_date", "year_to_date",
	)
	apijson.RegisterFieldValidator[AttributionReportNewParamsReportContextDateRangeObject2](
		"timeZone", "UTC",
	)
}

type AttributionReportUpdateParams struct {
	// Current revision from get_attribution_report, list_attribution_reports, or the
	// last write. A stale revision returns 409; reload before retrying.
	ExpectedRevision string `json:"expectedRevision" api:"required" format:"uuid"`
	// Optional description; null clears it.
	Description param.Opt[string] `json:"description,omitzero"`
	// Display name for the saved attribution report.
	Name param.Opt[string] `json:"name,omitzero"`
	// Version 1 attribution definition: one to five conversion outcomes, each with one
	// to ten unique event names, an optional filter, and an optional numeric event
	// property summed as value. Set the credit model (default last_touch), counting
	// unit (default unique), lookback window of 7, 14, 30, 60, or 90 days (default
	// 30), an optional filter applied to every outcome, and the value currency
	// (default USD). Outcome ids identify outcomes in results. Replaces the saved
	// definition on PATCH.
	Definition AttributionReportUpdateParamsDefinition `json:"definition,omitzero"`
	// Default UTC date range: a preset or inclusive from/to calendar dates. Replaces
	// the saved context on PATCH. Results cover at most 92 inclusive days.
	ReportContext AttributionReportUpdateParamsReportContext `json:"reportContext,omitzero"`
	paramObj
}

func (r AttributionReportUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AttributionReportUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributionReportUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Version 1 attribution definition: one to five conversion outcomes, each with one
// to ten unique event names, an optional filter, and an optional numeric event
// property summed as value. Set the credit model (default last_touch), counting
// unit (default unique), lookback window of 7, 14, 30, 60, or 90 days (default
// 30), an optional filter applied to every outcome, and the value currency
// (default USD). Outcome ids identify outcomes in results. Replaces the saved
// definition on PATCH.
//
// The properties Kind, Outcomes, Version are required.
type AttributionReportUpdateParamsDefinition struct {
	// Any of "attribution".
	Kind     string                                           `json:"kind,omitzero" api:"required"`
	Outcomes []AttributionReportUpdateParamsDefinitionOutcome `json:"outcomes,omitzero" api:"required"`
	// Any of 1.
	Version float64 `json:"version,omitzero" api:"required"`
	// Any of "unique", "totals", "sessions".
	Counting string `json:"counting,omitzero"`
	// Any of "USD", "CAD", "GBP", "EUR", "AUD".
	Currency string                                        `json:"currency,omitzero"`
	Filter   AttributionReportUpdateParamsDefinitionFilter `json:"filter,omitzero"`
	// Any of 7, 14, 30, 60, 90.
	LookbackDays float64 `json:"lookbackDays,omitzero"`
	// Any of "first_touch", "last_touch", "linear", "u_shaped", "j_shaped",
	// "time_decay".
	Model string `json:"model,omitzero"`
	paramObj
}

func (r AttributionReportUpdateParamsDefinition) MarshalJSON() (data []byte, err error) {
	type shadow AttributionReportUpdateParamsDefinition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributionReportUpdateParamsDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AttributionReportUpdateParamsDefinition](
		"kind", "attribution",
	)
	apijson.RegisterFieldValidator[AttributionReportUpdateParamsDefinition](
		"version", 1,
	)
	apijson.RegisterFieldValidator[AttributionReportUpdateParamsDefinition](
		"counting", "unique", "totals", "sessions",
	)
	apijson.RegisterFieldValidator[AttributionReportUpdateParamsDefinition](
		"currency", "USD", "CAD", "GBP", "EUR", "AUD",
	)
	apijson.RegisterFieldValidator[AttributionReportUpdateParamsDefinition](
		"lookbackDays", 7, 14, 30, 60, 90,
	)
	apijson.RegisterFieldValidator[AttributionReportUpdateParamsDefinition](
		"model", "first_touch", "last_touch", "linear", "u_shaped", "j_shaped", "time_decay",
	)
}

// The properties ID, EventNames, Label are required.
type AttributionReportUpdateParamsDefinitionOutcome struct {
	ID         string                                               `json:"id" api:"required" format:"uuid"`
	EventNames []string                                             `json:"eventNames,omitzero" api:"required"`
	Label      string                                               `json:"label" api:"required"`
	Filter     AttributionReportUpdateParamsDefinitionOutcomeFilter `json:"filter,omitzero"`
	Value      AttributionReportUpdateParamsDefinitionOutcomeValue  `json:"value,omitzero"`
	paramObj
}

func (r AttributionReportUpdateParamsDefinitionOutcome) MarshalJSON() (data []byte, err error) {
	type shadow AttributionReportUpdateParamsDefinitionOutcome
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributionReportUpdateParamsDefinitionOutcome) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Filter, Version are required.
type AttributionReportUpdateParamsDefinitionOutcomeFilter struct {
	Filter AttributionReportUpdateParamsDefinitionOutcomeFilterFilterUnion `json:"filter,omitzero" api:"required"`
	// Any of 1.
	Version float64 `json:"version,omitzero" api:"required"`
	paramObj
}

func (r AttributionReportUpdateParamsDefinitionOutcomeFilter) MarshalJSON() (data []byte, err error) {
	type shadow AttributionReportUpdateParamsDefinitionOutcomeFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributionReportUpdateParamsDefinitionOutcomeFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AttributionReportUpdateParamsDefinitionOutcomeFilter](
		"version", 1,
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AttributionReportUpdateParamsDefinitionOutcomeFilterFilterUnion struct {
	OfAttributionReportUpdatesDefinitionOutcomeFilterFilterObject  *AttributionReportUpdateParamsDefinitionOutcomeFilterFilterObject  `json:",omitzero,inline"`
	OfAttributionReportUpdatesDefinitionOutcomeFilterFilterObject2 *AttributionReportUpdateParamsDefinitionOutcomeFilterFilterObject2 `json:",omitzero,inline"`
	OfAttributionReportUpdatesDefinitionOutcomeFilterFilterObject3 *AttributionReportUpdateParamsDefinitionOutcomeFilterFilterObject3 `json:",omitzero,inline"`
	paramUnion
}

func (u AttributionReportUpdateParamsDefinitionOutcomeFilterFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAttributionReportUpdatesDefinitionOutcomeFilterFilterObject, u.OfAttributionReportUpdatesDefinitionOutcomeFilterFilterObject2, u.OfAttributionReportUpdatesDefinitionOutcomeFilterFilterObject3)
}
func (u *AttributionReportUpdateParamsDefinitionOutcomeFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Children, Kind are required.
type AttributionReportUpdateParamsDefinitionOutcomeFilterFilterObject struct {
	Children []any `json:"children,omitzero" api:"required"`
	// Any of "and", "or".
	Kind string `json:"kind,omitzero" api:"required"`
	paramObj
}

func (r AttributionReportUpdateParamsDefinitionOutcomeFilterFilterObject) MarshalJSON() (data []byte, err error) {
	type shadow AttributionReportUpdateParamsDefinitionOutcomeFilterFilterObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributionReportUpdateParamsDefinitionOutcomeFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AttributionReportUpdateParamsDefinitionOutcomeFilterFilterObject](
		"kind", "and", "or",
	)
}

// The properties Child, Kind are required.
type AttributionReportUpdateParamsDefinitionOutcomeFilterFilterObject2 struct {
	Child any `json:"child,omitzero" api:"required"`
	// Any of "not".
	Kind string `json:"kind,omitzero" api:"required"`
	paramObj
}

func (r AttributionReportUpdateParamsDefinitionOutcomeFilterFilterObject2) MarshalJSON() (data []byte, err error) {
	type shadow AttributionReportUpdateParamsDefinitionOutcomeFilterFilterObject2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributionReportUpdateParamsDefinitionOutcomeFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AttributionReportUpdateParamsDefinitionOutcomeFilterFilterObject2](
		"kind", "not",
	)
}

// The properties Kind, Operator, Property are required.
type AttributionReportUpdateParamsDefinitionOutcomeFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind,omitzero" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator,omitzero" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path,omitzero"`
	// Any of "string", "number", "boolean".
	Type  string                                                                      `json:"type,omitzero"`
	Value AttributionReportUpdateParamsDefinitionOutcomeFilterFilterObject3ValueUnion `json:"value,omitzero"`
	paramObj
}

func (r AttributionReportUpdateParamsDefinitionOutcomeFilterFilterObject3) MarshalJSON() (data []byte, err error) {
	type shadow AttributionReportUpdateParamsDefinitionOutcomeFilterFilterObject3
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributionReportUpdateParamsDefinitionOutcomeFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AttributionReportUpdateParamsDefinitionOutcomeFilterFilterObject3](
		"kind", "predicate",
	)
	apijson.RegisterFieldValidator[AttributionReportUpdateParamsDefinitionOutcomeFilterFilterObject3](
		"operator", "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains", "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in", "exists", "missing", "is_null",
	)
	apijson.RegisterFieldValidator[AttributionReportUpdateParamsDefinitionOutcomeFilterFilterObject3](
		"type", "string", "number", "boolean",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AttributionReportUpdateParamsDefinitionOutcomeFilterFilterObject3ValueUnion struct {
	OfString                                                                 param.Opt[string]                                                                      `json:",omitzero,inline"`
	OfFloat                                                                  param.Opt[float64]                                                                     `json:",omitzero,inline"`
	OfBool                                                                   param.Opt[bool]                                                                        `json:",omitzero,inline"`
	OfAttributionReportUpdatesDefinitionOutcomeFilterFilterObject3ValueArray []AttributionReportUpdateParamsDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u AttributionReportUpdateParamsDefinitionOutcomeFilterFilterObject3ValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool, u.OfAttributionReportUpdatesDefinitionOutcomeFilterFilterObject3ValueArray)
}
func (u *AttributionReportUpdateParamsDefinitionOutcomeFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AttributionReportUpdateParamsDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u AttributionReportUpdateParamsDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *AttributionReportUpdateParamsDefinitionOutcomeFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Path, Property are required.
type AttributionReportUpdateParamsDefinitionOutcomeValue struct {
	Path []string `json:"path,omitzero" api:"required"`
	// Any of "event.properties".
	Property string `json:"property,omitzero" api:"required"`
	// Any of "number".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r AttributionReportUpdateParamsDefinitionOutcomeValue) MarshalJSON() (data []byte, err error) {
	type shadow AttributionReportUpdateParamsDefinitionOutcomeValue
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributionReportUpdateParamsDefinitionOutcomeValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AttributionReportUpdateParamsDefinitionOutcomeValue](
		"property", "event.properties",
	)
	apijson.RegisterFieldValidator[AttributionReportUpdateParamsDefinitionOutcomeValue](
		"type", "number",
	)
}

// The properties Filter, Version are required.
type AttributionReportUpdateParamsDefinitionFilter struct {
	Filter AttributionReportUpdateParamsDefinitionFilterFilterUnion `json:"filter,omitzero" api:"required"`
	// Any of 1.
	Version float64 `json:"version,omitzero" api:"required"`
	paramObj
}

func (r AttributionReportUpdateParamsDefinitionFilter) MarshalJSON() (data []byte, err error) {
	type shadow AttributionReportUpdateParamsDefinitionFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributionReportUpdateParamsDefinitionFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AttributionReportUpdateParamsDefinitionFilter](
		"version", 1,
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AttributionReportUpdateParamsDefinitionFilterFilterUnion struct {
	OfAttributionReportUpdatesDefinitionFilterFilterObject  *AttributionReportUpdateParamsDefinitionFilterFilterObject  `json:",omitzero,inline"`
	OfAttributionReportUpdatesDefinitionFilterFilterObject2 *AttributionReportUpdateParamsDefinitionFilterFilterObject2 `json:",omitzero,inline"`
	OfAttributionReportUpdatesDefinitionFilterFilterObject3 *AttributionReportUpdateParamsDefinitionFilterFilterObject3 `json:",omitzero,inline"`
	paramUnion
}

func (u AttributionReportUpdateParamsDefinitionFilterFilterUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAttributionReportUpdatesDefinitionFilterFilterObject, u.OfAttributionReportUpdatesDefinitionFilterFilterObject2, u.OfAttributionReportUpdatesDefinitionFilterFilterObject3)
}
func (u *AttributionReportUpdateParamsDefinitionFilterFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties Children, Kind are required.
type AttributionReportUpdateParamsDefinitionFilterFilterObject struct {
	Children []any `json:"children,omitzero" api:"required"`
	// Any of "and", "or".
	Kind string `json:"kind,omitzero" api:"required"`
	paramObj
}

func (r AttributionReportUpdateParamsDefinitionFilterFilterObject) MarshalJSON() (data []byte, err error) {
	type shadow AttributionReportUpdateParamsDefinitionFilterFilterObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributionReportUpdateParamsDefinitionFilterFilterObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AttributionReportUpdateParamsDefinitionFilterFilterObject](
		"kind", "and", "or",
	)
}

// The properties Child, Kind are required.
type AttributionReportUpdateParamsDefinitionFilterFilterObject2 struct {
	Child any `json:"child,omitzero" api:"required"`
	// Any of "not".
	Kind string `json:"kind,omitzero" api:"required"`
	paramObj
}

func (r AttributionReportUpdateParamsDefinitionFilterFilterObject2) MarshalJSON() (data []byte, err error) {
	type shadow AttributionReportUpdateParamsDefinitionFilterFilterObject2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributionReportUpdateParamsDefinitionFilterFilterObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AttributionReportUpdateParamsDefinitionFilterFilterObject2](
		"kind", "not",
	)
}

// The properties Kind, Operator, Property are required.
type AttributionReportUpdateParamsDefinitionFilterFilterObject3 struct {
	// Any of "predicate".
	Kind string `json:"kind,omitzero" api:"required"`
	// Any of "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains",
	// "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in",
	// "exists", "missing", "is_null".
	Operator string   `json:"operator,omitzero" api:"required"`
	Property string   `json:"property" api:"required"`
	Path     []string `json:"path,omitzero"`
	// Any of "string", "number", "boolean".
	Type  string                                                               `json:"type,omitzero"`
	Value AttributionReportUpdateParamsDefinitionFilterFilterObject3ValueUnion `json:"value,omitzero"`
	paramObj
}

func (r AttributionReportUpdateParamsDefinitionFilterFilterObject3) MarshalJSON() (data []byte, err error) {
	type shadow AttributionReportUpdateParamsDefinitionFilterFilterObject3
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributionReportUpdateParamsDefinitionFilterFilterObject3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AttributionReportUpdateParamsDefinitionFilterFilterObject3](
		"kind", "predicate",
	)
	apijson.RegisterFieldValidator[AttributionReportUpdateParamsDefinitionFilterFilterObject3](
		"operator", "eq", "neq", "gt", "gte", "lt", "lte", "contains", "not_contains", "contains_ci", "not_contains_ci", "starts_with", "ends_with", "in", "not_in", "exists", "missing", "is_null",
	)
	apijson.RegisterFieldValidator[AttributionReportUpdateParamsDefinitionFilterFilterObject3](
		"type", "string", "number", "boolean",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AttributionReportUpdateParamsDefinitionFilterFilterObject3ValueUnion struct {
	OfString                                                          param.Opt[string]                                                               `json:",omitzero,inline"`
	OfFloat                                                           param.Opt[float64]                                                              `json:",omitzero,inline"`
	OfBool                                                            param.Opt[bool]                                                                 `json:",omitzero,inline"`
	OfAttributionReportUpdatesDefinitionFilterFilterObject3ValueArray []AttributionReportUpdateParamsDefinitionFilterFilterObject3ValueArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u AttributionReportUpdateParamsDefinitionFilterFilterObject3ValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool, u.OfAttributionReportUpdatesDefinitionFilterFilterObject3ValueArray)
}
func (u *AttributionReportUpdateParamsDefinitionFilterFilterObject3ValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AttributionReportUpdateParamsDefinitionFilterFilterObject3ValueArrayItemUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u AttributionReportUpdateParamsDefinitionFilterFilterObject3ValueArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *AttributionReportUpdateParamsDefinitionFilterFilterObject3ValueArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Default UTC date range: a preset or inclusive from/to calendar dates. Replaces
// the saved context on PATCH. Results cover at most 92 inclusive days.
//
// The property DateRange is required.
type AttributionReportUpdateParamsReportContext struct {
	DateRange AttributionReportUpdateParamsReportContextDateRangeUnion `json:"dateRange,omitzero" api:"required"`
	paramObj
}

func (r AttributionReportUpdateParamsReportContext) MarshalJSON() (data []byte, err error) {
	type shadow AttributionReportUpdateParamsReportContext
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributionReportUpdateParamsReportContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AttributionReportUpdateParamsReportContextDateRangeUnion struct {
	OfAttributionReportUpdatesReportContextDateRangeObject  *AttributionReportUpdateParamsReportContextDateRangeObject  `json:",omitzero,inline"`
	OfAttributionReportUpdatesReportContextDateRangeObject2 *AttributionReportUpdateParamsReportContextDateRangeObject2 `json:",omitzero,inline"`
	paramUnion
}

func (u AttributionReportUpdateParamsReportContextDateRangeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAttributionReportUpdatesReportContextDateRangeObject, u.OfAttributionReportUpdatesReportContextDateRangeObject2)
}
func (u *AttributionReportUpdateParamsReportContextDateRangeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// The properties From, To are required.
type AttributionReportUpdateParamsReportContextDateRangeObject struct {
	From string `json:"from" api:"required"`
	To   string `json:"to" api:"required"`
	// Any of "UTC".
	TimeZone string `json:"timeZone,omitzero"`
	paramObj
}

func (r AttributionReportUpdateParamsReportContextDateRangeObject) MarshalJSON() (data []byte, err error) {
	type shadow AttributionReportUpdateParamsReportContextDateRangeObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributionReportUpdateParamsReportContextDateRangeObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AttributionReportUpdateParamsReportContextDateRangeObject](
		"timeZone", "UTC",
	)
}

// The property Preset is required.
type AttributionReportUpdateParamsReportContextDateRangeObject2 struct {
	// Any of "today", "yesterday", "last_7_days", "last_30_days", "last_90_days",
	// "last_12_months", "week_to_date", "month_to_date", "year_to_date".
	Preset string `json:"preset,omitzero" api:"required"`
	// Any of "UTC".
	TimeZone string `json:"timeZone,omitzero"`
	paramObj
}

func (r AttributionReportUpdateParamsReportContextDateRangeObject2) MarshalJSON() (data []byte, err error) {
	type shadow AttributionReportUpdateParamsReportContextDateRangeObject2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributionReportUpdateParamsReportContextDateRangeObject2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AttributionReportUpdateParamsReportContextDateRangeObject2](
		"preset", "today", "yesterday", "last_7_days", "last_30_days", "last_90_days", "last_12_months", "week_to_date", "month_to_date", "year_to_date",
	)
	apijson.RegisterFieldValidator[AttributionReportUpdateParamsReportContextDateRangeObject2](
		"timeZone", "UTC",
	)
}

type AttributionReportDeleteParams struct {
	// Current revision from get_attribution_report, list_attribution_reports, or the
	// last write. A stale revision returns 409; reload before retrying.
	ExpectedRevision string `query:"expectedRevision" api:"required" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [AttributionReportDeleteParams]'s query parameters as
// `url.Values`.
func (r AttributionReportDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AttributionReportDuplicateParams struct {
	// Current revision from get_attribution_report, list_attribution_reports, or the
	// last write. A stale revision returns 409; reload before retrying.
	ExpectedRevision string `json:"expectedRevision" api:"required" format:"uuid"`
	paramObj
}

func (r AttributionReportDuplicateParams) MarshalJSON() (data []byte, err error) {
	type shadow AttributionReportDuplicateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributionReportDuplicateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AttributionReportPreviewParams struct {
	// Required input: definition, from, and to; comparison and interval are optional.
	// JSON-encode this object in the REST query string; maximum 16384 bytes before URL
	// encoding. Computing a result never saves or changes a report.
	Input string `query:"input" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [AttributionReportPreviewParams]'s query parameters as
// `url.Values`.
func (r AttributionReportPreviewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AttributionReportVisitorsParams struct {
	// Required input: the same definition, from, to, and optional comparison and
	// interval that produced the result, plus the selection to open; limit and cursor
	// are optional. JSON-encode this object in the REST query string; maximum 16384
	// bytes before URL encoding.
	Input string `query:"input" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [AttributionReportVisitorsParams]'s query parameters as
// `url.Values`.
func (r AttributionReportVisitorsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AttributionReportInsightsParams struct {
	// Required input: definition, from, and to; comparison and interval are optional.
	// Use the same input as the preview to explain what moved in that result.
	// JSON-encode this object in the REST query string; maximum 16384 bytes before URL
	// encoding. Nothing is saved.
	Input string `query:"input" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [AttributionReportInsightsParams]'s query parameters as
// `url.Values`.
func (r AttributionReportInsightsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
