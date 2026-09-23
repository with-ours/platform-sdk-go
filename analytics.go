// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/with-ours/platform-sdk-go/internal/apijson"
	"github.com/with-ours/platform-sdk-go/internal/apiquery"
	"github.com/with-ours/platform-sdk-go/internal/requestconfig"
	"github.com/with-ours/platform-sdk-go/option"
	"github.com/with-ours/platform-sdk-go/packages/param"
	"github.com/with-ours/platform-sdk-go/packages/respjson"
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
// query definitions. Requires scope: web-analytics:view
func (r *AnalyticsService) QueryCatalog(ctx context.Context, opts ...option.RequestOption) (res *AnalyticsQueryCatalogResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/analytics/query-catalog"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Discover property paths recently observed in event data. Results are suggestions
// from a fixed recent 30-day sample and may be incomplete; callers can still use
// manually entered paths. Use `eventName` and `sourceId` to narrow the
// suggestions. Requires scope: web-analytics:view
func (r *AnalyticsService) PropertySuggestions(ctx context.Context, query AnalyticsPropertySuggestionsParams, opts ...option.RequestOption) (res *AnalyticsPropertySuggestionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/analytics/property-suggestions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AnalyticsQueryCatalogResponse struct {
	Entries     []AnalyticsQueryCatalogResponseEntry `json:"entries" api:"required"`
	Limitations []string                             `json:"limitations" api:"required"`
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
	ID              string   `json:"id" api:"required"`
	Label           string   `json:"label" api:"required"`
	MissingBehavior string   `json:"missingBehavior" api:"required"`
	Operators       []string `json:"operators" api:"required"`
	Provider        string   `json:"provider" api:"required"`
	Scope           string   `json:"scope" api:"required"`
	Type            string   `json:"type" api:"required"`
	AllowedTypes    []string `json:"allowedTypes" api:"nullable"`
	RequiresPath    bool     `json:"requiresPath" api:"nullable"`
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
	Entities []AnalyticsPropertySuggestionsResponseEntity `json:"entities" api:"required"`
	From     time.Time                                    `json:"from" api:"required" format:"date-time"`
	Sampled  bool                                         `json:"sampled" api:"required"`
	To       time.Time                                    `json:"to" api:"required" format:"date-time"`
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
	Path []string `json:"path" api:"required"`
	// Any of "event.properties", "event.visitor_properties".
	Property string `json:"property" api:"required"`
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
