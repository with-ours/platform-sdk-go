// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy

import (
	"context"
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

// AnalyticsService contains methods and other services that help with interacting
// with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAnalyticsService] method instead.
type AnalyticsService struct {
	Options []option.RequestOption
}

// NewAnalyticsService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAnalyticsService(opts ...option.RequestOption) (r AnalyticsService) {
	r = AnalyticsService{}
	r.Options = opts
	return
}

// Discover the filter properties, operators, and providers available to analytics
// query definitions. Requires API-key scope or current OAuth user permission:
// variables:find-default
func (r *AnalyticsService) QueryCatalog(ctx context.Context, opts ...option.RequestOption) (res *AnalyticsQueryCatalogResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/analytics/query-catalog"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Discover property paths recently observed in event data. Results are suggestions
// from a fixed recent 30-day sample and may be incomplete; callers can still use
// manually entered paths. Use `eventName` and `sourceId` to narrow the
// suggestions. Requires API-key scope or current OAuth user permission:
// variables:find-custom, web-analytics:view
func (r *AnalyticsService) PropertySuggestions(ctx context.Context, query AnalyticsPropertySuggestionsParams, opts ...option.RequestOption) (res *AnalyticsPropertySuggestionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/analytics/property-suggestions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Discover the event names this account has collected in the last 30 days, for
// Insight metric `eventNames`, the `eventName` of funnel steps and journey flows,
// or the `eventName` filter of `GET /rest/v1/analytics/property-suggestions`.
// Returns `{ entities }` in alphabetical order. The list is not paginated: it
// holds up to the 500 most frequent names, and newly seen events may take a few
// minutes to appear. Names missing from the list still work when entered exactly.
// Requires API-key scope or current OAuth user permission: web-analytics:view
func (r *AnalyticsService) EventNames(ctx context.Context, opts ...option.RequestOption) (res *AnalyticsEventNamesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/analytics/event-names"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type AnalyticsQueryCatalogResponse struct {
	// Fields available to analytics report filters.
	Entries []AnalyticsQueryCatalogResponseEntry `json:"entries" api:"required"`
	// Known limits of analytics filtering.
	Limitations []string `json:"limitations" api:"required"`
	// Catalog format version.
	//
	// Any of 1.
	Version float64 `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entries     respjson.Field
		Limitations respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AnalyticsQueryCatalogResponse) RawJSON() string { return r.JSON.raw }
func (r *AnalyticsQueryCatalogResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsQueryCatalogResponseEntry struct {
	// Catalog identifier for this filterable field.
	ID string `json:"id" api:"required"`
	// Human-readable field name.
	Label string `json:"label" api:"required"`
	// How a filter treats events where the field is absent.
	MissingBehavior string `json:"missingBehavior" api:"required"`
	// Operators this field accepts in a filter condition.
	Operators []string `json:"operators" api:"required"`
	// Data source that supplies the field value.
	Provider string `json:"provider" api:"required"`
	// Whether the field belongs to the event or to its session.
	Scope string `json:"scope" api:"required"`
	// Value type the filter compares against.
	Type string `json:"type" api:"required"`
	// Value types allowed for a path-based field.
	AllowedTypes []string `json:"allowedTypes" api:"nullable"`
	// True when a property path must be supplied, as for custom event properties.
	RequiresPath bool `json:"requiresPath" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Label           respjson.Field
		MissingBehavior respjson.Field
		Operators       respjson.Field
		Provider        respjson.Field
		Scope           respjson.Field
		Type            respjson.Field
		AllowedTypes    respjson.Field
		RequiresPath    respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AnalyticsQueryCatalogResponseEntry) RawJSON() string { return r.JSON.raw }
func (r *AnalyticsQueryCatalogResponseEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsPropertySuggestionsResponse struct {
	// Property paths and value types observed in a sample of recent events.
	Entities []AnalyticsPropertySuggestionsResponseEntity `json:"entities" api:"required"`
	// Start of the observation window the suggestions were drawn from.
	From time.Time `json:"from" api:"required" format:"date-time"`
	// True when the suggestions come from a sample, so some paths may be missing.
	Sampled bool `json:"sampled" api:"required"`
	// End of the observation window the suggestions were drawn from.
	To time.Time `json:"to" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		From        respjson.Field
		Sampled     respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AnalyticsPropertySuggestionsResponse) RawJSON() string { return r.JSON.raw }
func (r *AnalyticsPropertySuggestionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsPropertySuggestionsResponseEntity struct {
	// Property path segments, outermost first.
	Path []string `json:"path" api:"required"`
	// Property bag the path belongs to.
	//
	// Any of "event.properties", "event.visitor_properties".
	Property string `json:"property" api:"required"`
	// Value type observed at this path.
	//
	// Any of "string", "number", "boolean".
	Type string `json:"type" api:"required"`
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
func (r AnalyticsPropertySuggestionsResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *AnalyticsPropertySuggestionsResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsEventNamesResponse struct {
	// Distinct event names seen in the last 30 days, alphabetical. Not paginated: the
	// list holds up to the 500 most frequent names.
	Entities []AnalyticsEventNamesResponseEntity `json:"entities" api:"required"`
	// Start of the observation window the suggestions were drawn from.
	From time.Time `json:"from" api:"required" format:"date-time"`
	// True when rare names may be missing from the list. Missing names still work when
	// entered exactly.
	Sampled bool `json:"sampled" api:"required"`
	// End of the observation window the suggestions were drawn from.
	To time.Time `json:"to" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		From        respjson.Field
		Sampled     respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AnalyticsEventNamesResponse) RawJSON() string { return r.JSON.raw }
func (r *AnalyticsEventNamesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsEventNamesResponseEntity struct {
	// Event name exactly as collected, including `$`-prefixed system events.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AnalyticsEventNamesResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *AnalyticsEventNamesResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsPropertySuggestionsParams struct {
	// Optional exact event name to scope the suggestions.
	EventName param.Opt[string] `query:"eventName,omitzero" json:"-"`
	// Optional exact source identifier to scope the suggestions.
	SourceID param.Opt[string] `query:"sourceId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AnalyticsPropertySuggestionsParams]'s query parameters as
// `url.Values`.
func (r AnalyticsPropertySuggestionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
