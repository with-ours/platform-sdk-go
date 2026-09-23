// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy

import (
	"context"
	"errors"
	"fmt"
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
// scope: web-analytics:view
func (r *FunnelService) List(ctx context.Context, opts ...option.RequestOption) (res *FunnelListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/funnels"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Create a funnel from a versioned definition. Returns the complete saved
// configuration. Requires scope: web-analytics:write
func (r *FunnelService) New(ctx context.Context, body FunnelNewParams, opts ...option.RequestOption) (res *FunnelNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/funnels"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Fetch a single funnel configuration by its id. Returns `404` when the funnel
// does not exist or belongs to a different account. Requires scope:
// web-analytics:view
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
// definition can be replaced by supplying `queryDefinition`. Requires scope:
// web-analytics:write
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
// scope: web-analytics:write
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
// scope: web-analytics:write
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
// span at most 91 days including both endpoints. Requires scope:
// web-analytics:view
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
	QueryDefinition any `json:"queryDefinition" api:"required"`
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
	QueryDefinition any `json:"queryDefinition" api:"required"`
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
	QueryDefinition any `json:"queryDefinition" api:"required"`
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
	QueryDefinition any `json:"queryDefinition" api:"required"`
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
	QueryDefinition any `json:"queryDefinition" api:"required"`
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
	QueryDefinition any               `json:"queryDefinition,omitzero" api:"required"`
	Description     param.Opt[string] `json:"description,omitzero"`
	Watched         param.Opt[bool]   `json:"watched,omitzero"`
	paramObj
}

func (r FunnelNewParams) MarshalJSON() (data []byte, err error) {
	type shadow FunnelNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelUpdateParams struct {
	Description param.Opt[string] `json:"description,omitzero"`
	Watched     param.Opt[bool]   `json:"watched,omitzero"`
	Name        param.Opt[string] `json:"name,omitzero"`
	// Versioned funnel definition with typed steps and filters.
	QueryDefinition any `json:"queryDefinition,omitzero"`
	paramObj
}

func (r FunnelUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow FunnelUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
