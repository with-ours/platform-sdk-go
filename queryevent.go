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

// QueryEventService contains methods and other services that help with interacting
// with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQueryEventService] method instead.
type QueryEventService struct {
	Options []option.RequestOption
}

// NewQueryEventService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewQueryEventService(opts ...option.RequestOption) (r QueryEventService) {
	r = QueryEventService{}
	r.Options = opts
	return
}

// List matching event occurrences. JSON-encode input in the query string.
// includeCount defaults to false. Requests that include raw event or visitor
// properties also require event-detail permission, checked by the resolver; a
// missing permission is masked as 404 Not Found. Requires API-key scope or current
// OAuth user permission: report:list-events
func (r *QueryEventService) List(ctx context.Context, query QueryEventListParams, opts ...option.RequestOption) (res *QueryEventListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/query-events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Return Query Events date, page, selection, projection, and export limits.
// Requires API-key scope or current OAuth user permission: report:list-events
func (r *QueryEventService) Capabilities(ctx context.Context, opts ...option.RequestOption) (res *QueryEventCapabilitiesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/query-events/capabilities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List the filter fields and operators supported by Query Events. Requires API-key
// scope or current OAuth user permission: report:list-events
func (r *QueryEventService) Catalog(ctx context.Context, opts ...option.RequestOption) (res *QueryEventCatalogResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/query-events/query-catalog"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Suggest observed event and visitor property paths. Optionally scope by exact
// event name or source id. Requires API-key scope or current OAuth user
// permission: report:list-events
func (r *QueryEventService) PropertySuggestions(ctx context.Context, query QueryEventPropertySuggestionsParams, opts ...option.RequestOption) (res *QueryEventPropertySuggestionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/query-events/property-suggestions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List observed event names available to Query Events. Requires API-key scope or
// current OAuth user permission: report:list-events
func (r *QueryEventService) EventNames(ctx context.Context, opts ...option.RequestOption) (res *QueryEventEventNamesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/query-events/event-names"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Find one event occurrence using the same JSON-encoded input and cursor returned
// by a Query Events list. Requires API-key scope or current OAuth user permission:
// report:list-events, report:view-event
func (r *QueryEventService) Occurrence(ctx context.Context, query QueryEventOccurrenceParams, opts ...option.RequestOption) (res *QueryEventOccurrenceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/query-events/occurrence"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Export matching events as CSV or JSON, up to the supported row limit.
// JSON-encode input in the query string. Requires API-key scope or current OAuth
// user permission: report:list-events, report:view-event
func (r *QueryEventService) Export(ctx context.Context, query QueryEventExportParams, opts ...option.RequestOption) (res *QueryEventExportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/query-events/export"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type QueryEventListResponse struct {
	Freshness    QueryEventListResponseFreshness  `json:"freshness" api:"required"`
	MatchedCount float64                          `json:"matchedCount" api:"required"`
	Pagination   QueryEventListResponsePagination `json:"pagination" api:"required"`
	Rows         []QueryEventListResponseRow      `json:"rows" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Freshness    respjson.Field
		MatchedCount respjson.Field
		Pagination   respjson.Field
		Rows         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryEventListResponse) RawJSON() string { return r.JSON.raw }
func (r *QueryEventListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryEventListResponseFreshness struct {
	// Any of "UNKNOWN".
	CoverageStatus    string    `json:"coverageStatus" api:"required"`
	CoverageWatermark time.Time `json:"coverageWatermark" api:"required" format:"date-time"`
	SyncCompletedAt   time.Time `json:"syncCompletedAt" api:"nullable" format:"date-time"`
	SyncCursor        time.Time `json:"syncCursor" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CoverageStatus    respjson.Field
		CoverageWatermark respjson.Field
		SyncCompletedAt   respjson.Field
		SyncCursor        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryEventListResponseFreshness) RawJSON() string { return r.JSON.raw }
func (r *QueryEventListResponseFreshness) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryEventListResponsePagination struct {
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
func (r QueryEventListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *QueryEventListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryEventListResponseRow struct {
	Columns           []QueryEventListResponseRowColumn `json:"columns" api:"required"`
	Cursor            string                            `json:"cursor" api:"required"`
	DistinctID        string                            `json:"distinctId" api:"required"`
	Email             string                            `json:"email" api:"required"`
	EventName         string                            `json:"eventName" api:"required"`
	EventProperties   map[string]any                    `json:"eventProperties" api:"required"`
	IsBot             bool                              `json:"isBot" api:"required"`
	IsTestEvent       bool                              `json:"isTestEvent" api:"required"`
	PagePath          string                            `json:"pagePath" api:"required"`
	SessionID         string                            `json:"sessionId" api:"required"`
	SourceID          string                            `json:"sourceId" api:"required"`
	SourceName        string                            `json:"sourceName" api:"required"`
	SourceType        string                            `json:"sourceType" api:"required"`
	Timestamp         time.Time                         `json:"timestamp" api:"required" format:"date-time"`
	VisitorID         string                            `json:"visitorId" api:"required"`
	VisitorProperties map[string]any                    `json:"visitorProperties" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Columns           respjson.Field
		Cursor            respjson.Field
		DistinctID        respjson.Field
		Email             respjson.Field
		EventName         respjson.Field
		EventProperties   respjson.Field
		IsBot             respjson.Field
		IsTestEvent       respjson.Field
		PagePath          respjson.Field
		SessionID         respjson.Field
		SourceID          respjson.Field
		SourceName        respjson.Field
		SourceType        respjson.Field
		Timestamp         respjson.Field
		VisitorID         respjson.Field
		VisitorProperties respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryEventListResponseRow) RawJSON() string { return r.JSON.raw }
func (r *QueryEventListResponseRow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryEventListResponseRowColumn struct {
	Path     []string `json:"path" api:"required"`
	Property string   `json:"property" api:"required"`
	// Any of "STRING", "NUMBER", "BOOLEAN".
	Type  string `json:"type" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Path        respjson.Field
		Property    respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryEventListResponseRowColumn) RawJSON() string { return r.JSON.raw }
func (r *QueryEventListResponseRowColumn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryEventCapabilitiesResponse struct {
	DefaultPageSize int64 `json:"defaultPageSize" api:"required"`
	// Any of 20.
	MaxColumns float64 `json:"maxColumns" api:"required"`
	// Any of 366.
	MaxDateRangeDays float64 `json:"maxDateRangeDays" api:"required"`
	// Any of 100.
	MaxExportRows float64 `json:"maxExportRows" api:"required"`
	// Any of 250.
	MaxPageSize float64 `json:"maxPageSize" api:"required"`
	// Any of 10.
	MaxSelections                     float64 `json:"maxSelections" api:"required"`
	SupportsVerifiedCoverageWatermark bool    `json:"supportsVerifiedCoverageWatermark" api:"required"`
	// Any of 1.
	Version float64 `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DefaultPageSize                   respjson.Field
		MaxColumns                        respjson.Field
		MaxDateRangeDays                  respjson.Field
		MaxExportRows                     respjson.Field
		MaxPageSize                       respjson.Field
		MaxSelections                     respjson.Field
		SupportsVerifiedCoverageWatermark respjson.Field
		Version                           respjson.Field
		ExtraFields                       map[string]respjson.Field
		raw                               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryEventCapabilitiesResponse) RawJSON() string { return r.JSON.raw }
func (r *QueryEventCapabilitiesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryEventCatalogResponse struct {
	// Fields available to analytics report filters.
	Entries []QueryEventCatalogResponseEntry `json:"entries" api:"required"`
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
func (r QueryEventCatalogResponse) RawJSON() string { return r.JSON.raw }
func (r *QueryEventCatalogResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryEventCatalogResponseEntry struct {
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
func (r QueryEventCatalogResponseEntry) RawJSON() string { return r.JSON.raw }
func (r *QueryEventCatalogResponseEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryEventPropertySuggestionsResponse struct {
	// Property paths and value types observed in a sample of recent events.
	Entities []QueryEventPropertySuggestionsResponseEntity `json:"entities" api:"required"`
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
func (r QueryEventPropertySuggestionsResponse) RawJSON() string { return r.JSON.raw }
func (r *QueryEventPropertySuggestionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryEventPropertySuggestionsResponseEntity struct {
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
func (r QueryEventPropertySuggestionsResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *QueryEventPropertySuggestionsResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryEventEventNamesResponse struct {
	// Distinct event names seen in the last 30 days, alphabetical. Not paginated: the
	// list holds up to the 500 most frequent names.
	Entities []QueryEventEventNamesResponseEntity `json:"entities" api:"required"`
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
func (r QueryEventEventNamesResponse) RawJSON() string { return r.JSON.raw }
func (r *QueryEventEventNamesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryEventEventNamesResponseEntity struct {
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
func (r QueryEventEventNamesResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *QueryEventEventNamesResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryEventOccurrenceResponse struct {
	Columns           []QueryEventOccurrenceResponseColumn `json:"columns" api:"required"`
	Cursor            string                               `json:"cursor" api:"required"`
	DistinctID        string                               `json:"distinctId" api:"required"`
	Email             string                               `json:"email" api:"required"`
	EventName         string                               `json:"eventName" api:"required"`
	EventProperties   map[string]any                       `json:"eventProperties" api:"required"`
	IsBot             bool                                 `json:"isBot" api:"required"`
	IsTestEvent       bool                                 `json:"isTestEvent" api:"required"`
	PagePath          string                               `json:"pagePath" api:"required"`
	SessionID         string                               `json:"sessionId" api:"required"`
	SourceID          string                               `json:"sourceId" api:"required"`
	SourceName        string                               `json:"sourceName" api:"required"`
	SourceType        string                               `json:"sourceType" api:"required"`
	Timestamp         time.Time                            `json:"timestamp" api:"required" format:"date-time"`
	VisitorID         string                               `json:"visitorId" api:"required"`
	VisitorProperties map[string]any                       `json:"visitorProperties" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Columns           respjson.Field
		Cursor            respjson.Field
		DistinctID        respjson.Field
		Email             respjson.Field
		EventName         respjson.Field
		EventProperties   respjson.Field
		IsBot             respjson.Field
		IsTestEvent       respjson.Field
		PagePath          respjson.Field
		SessionID         respjson.Field
		SourceID          respjson.Field
		SourceName        respjson.Field
		SourceType        respjson.Field
		Timestamp         respjson.Field
		VisitorID         respjson.Field
		VisitorProperties respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryEventOccurrenceResponse) RawJSON() string { return r.JSON.raw }
func (r *QueryEventOccurrenceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryEventOccurrenceResponseColumn struct {
	Path     []string `json:"path" api:"required"`
	Property string   `json:"property" api:"required"`
	// Any of "STRING", "NUMBER", "BOOLEAN".
	Type  string `json:"type" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Path        respjson.Field
		Property    respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryEventOccurrenceResponseColumn) RawJSON() string { return r.JSON.raw }
func (r *QueryEventOccurrenceResponseColumn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryEventExportResponse struct {
	Content string `json:"content" api:"required"`
	// Any of "CSV", "JSON".
	Format     QueryEventExportResponseFormat     `json:"format" api:"required"`
	Pagination QueryEventExportResponsePagination `json:"pagination" api:"required"`
	RowCount   int64                              `json:"rowCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Format      respjson.Field
		Pagination  respjson.Field
		RowCount    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryEventExportResponse) RawJSON() string { return r.JSON.raw }
func (r *QueryEventExportResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryEventExportResponseFormat string

const (
	QueryEventExportResponseFormatCsv  QueryEventExportResponseFormat = "CSV"
	QueryEventExportResponseFormatJson QueryEventExportResponseFormat = "JSON"
)

type QueryEventExportResponsePagination struct {
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
func (r QueryEventExportResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *QueryEventExportResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryEventListParams struct {
	// Query Events filters and projection. JSON-encode input in the REST query string.
	Input string `query:"input" api:"required" json:"-"`
	// Cursor from the previous page.
	Cursor       param.Opt[string] `query:"cursor,omitzero" json:"-"`
	IncludeCount param.Opt[bool]   `query:"includeCount,omitzero" json:"-"`
	Limit        param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [QueryEventListParams]'s query parameters as `url.Values`.
func (r QueryEventListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type QueryEventPropertySuggestionsParams struct {
	// Optional exact event name to scope the suggestions.
	EventName param.Opt[string] `query:"eventName,omitzero" json:"-"`
	// Optional exact source identifier to scope the suggestions.
	SourceID param.Opt[string] `query:"sourceId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [QueryEventPropertySuggestionsParams]'s query parameters as
// `url.Values`.
func (r QueryEventPropertySuggestionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type QueryEventOccurrenceParams struct {
	// Cursor from the page containing this event.
	Cursor string `query:"cursor" api:"required" json:"-"`
	// The same filters used to obtain the event cursor. JSON-encode input in the REST
	// query string.
	Input string `query:"input" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [QueryEventOccurrenceParams]'s query parameters as
// `url.Values`.
func (r QueryEventOccurrenceParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type QueryEventExportParams struct {
	// Export encoding.
	//
	// Any of "CSV", "JSON".
	Format QueryEventExportParamsFormat `query:"format,omitzero" api:"required" json:"-"`
	// Query Events filters and projection. JSON-encode input in the REST query string.
	Input string `query:"input" api:"required" json:"-"`
	// Start after this event cursor.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [QueryEventExportParams]'s query parameters as `url.Values`.
func (r QueryEventExportParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Export encoding.
type QueryEventExportParamsFormat string

const (
	QueryEventExportParamsFormatCsv  QueryEventExportParamsFormat = "CSV"
	QueryEventExportParamsFormatJson QueryEventExportParamsFormat = "JSON"
)
