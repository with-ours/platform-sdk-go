// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomwithoursplatformsdkgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/with-ours/platform-sdk-go/internal/apijson"
	"github.com/with-ours/platform-sdk-go/internal/apiquery"
	"github.com/with-ours/platform-sdk-go/internal/requestconfig"
	"github.com/with-ours/platform-sdk-go/option"
	"github.com/with-ours/platform-sdk-go/packages/pagination"
	"github.com/with-ours/platform-sdk-go/packages/param"
	"github.com/with-ours/platform-sdk-go/packages/respjson"
)

// ExperimentService contains methods and other services that help with interacting
// with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExperimentService] method instead.
type ExperimentService struct {
	Options []option.RequestOption
}

// NewExperimentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewExperimentService(opts ...option.RequestOption) (r ExperimentService) {
	r = ExperimentService{}
	r.Options = opts
	return
}

// List experiments for this account. Supports cursor pagination and filtering by
// `status`, `type`, and free-text `search` matched against experiment id, name,
// and description. Combine filters with AND semantics. Requires scope:
// experiment:list
func (r *ExperimentService) List(ctx context.Context, query ExperimentListParams, opts ...option.RequestOption) (res *pagination.Cursor[ExperimentListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "rest/v1/experiments"
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

// List experiments for this account. Supports cursor pagination and filtering by
// `status`, `type`, and free-text `search` matched against experiment id, name,
// and description. Combine filters with AND semantics. Requires scope:
// experiment:list
func (r *ExperimentService) ListAutoPaging(ctx context.Context, query ExperimentListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[ExperimentListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Create a new experiment. Requires scope: experiment:create
func (r *ExperimentService) New(ctx context.Context, body ExperimentNewParams, opts ...option.RequestOption) (res *ExperimentNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/experiments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Find a single experiment by ID. Requires scope: experiment:find
func (r *ExperimentService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ExperimentGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/experiments/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Partially update an experiment. Only the fields you send are changed. Edits are
// only accepted while the experiment is in `draft` status — running, paused, and
// completed experiments return 409 with
// `Experiment can only be edited in draft status`. Use the lifecycle endpoints
// (`/start`, `/pause`, `/resume`, `/stop`) to change status. Requires scope:
// experiment:update
func (r *ExperimentService) Update(ctx context.Context, id string, body ExperimentUpdateParams, opts ...option.RequestOption) (res *ExperimentUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/experiments/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete an experiment. Requires scope: experiment:delete
func (r *ExperimentService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/experiments/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Start an experiment. The request body is optional — send `{}` to use defaults.
// Requires scope: experiment:start
func (r *ExperimentService) Start(ctx context.Context, id string, body ExperimentStartParams, opts ...option.RequestOption) (res *ExperimentStartResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/experiments/%s/start", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Stop an experiment. The request body is optional — send `{}` to stop without
// recording a winner. Requires scope: experiment:stop
func (r *ExperimentService) Stop(ctx context.Context, id string, body ExperimentStopParams, opts ...option.RequestOption) (res *ExperimentStopResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/experiments/%s/stop", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Pause a running experiment. Stops new variant assignments while preserving
// existing ones; the experiment can later be resumed. The request body is
// optional. Requires scope: experiment:stop
func (r *ExperimentService) Pause(ctx context.Context, id string, body ExperimentPauseParams, opts ...option.RequestOption) (res *ExperimentPauseResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/experiments/%s/pause", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Resume a previously-paused experiment so new visitors can be assigned again. The
// request body is optional. Requires scope: experiment:start
func (r *ExperimentService) Resume(ctx context.Context, id string, body ExperimentResumeParams, opts ...option.RequestOption) (res *ExperimentResumeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/experiments/%s/resume", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Aggregate per-variant impressions, conversions, conversion rate, and Bayesian
// probability-to-be-best across the experiment runtime window. Requires scope:
// experiment:find
func (r *ExperimentService) Results(ctx context.Context, id string, query ExperimentResultsParams, opts ...option.RequestOption) (res *ExperimentResultsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/experiments/%s/results", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Per-day per-variant impressions, conversions, and conversion rate, sliced to a
// date range. Use this to chart trends, compare windows, or zoom in on a specific
// period. Pass `startDate` / `endDate` (`YYYY-MM-DD`, UTC, both inclusive) to set
// the window; both default to the full experiment runtime when omitted, so the
// no-arg call returns every day from start to today (or to `stoppedAt` for
// completed experiments). The response orders days oldest-first and omits days
// with no impressions, so an empty `days` array means there was no measured
// traffic in the window. Requires scope: experiment:find
func (r *ExperimentService) ResultsTimeSeries(ctx context.Context, id string, query ExperimentResultsTimeSeriesParams, opts ...option.RequestOption) (res *ExperimentResultsTimeSeriesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/experiments/%s/results-time-series", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ExperimentListResponse struct {
	// Unique identifier for the experiment.
	ID string `json:"id" api:"required"`
	// ISO-8601 timestamp when the experiment was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Stable code-facing key for the experiment. Use this with the headless SDK
	// `getExperimentByKey()` API instead of hard-coding opaque experiment IDs into
	// application code.
	Key string `json:"key" api:"required"`
	// Short, human-readable experiment name.
	Name string `json:"name" api:"required"`
	// Lifecycle state. `draft` is editable, `running` is active, `paused` is
	// temporarily inactive, and `completed` is permanently stopped.
	//
	// Any of "completed", "draft", "paused", "running".
	Status ExperimentListResponseStatus `json:"status" api:"required"`
	// Percent of eligible traffic assigned into the experiment. Use 0 to fully disable
	// enrollment without deleting the experiment.
	TrafficAllocation int64 `json:"trafficAllocation" api:"required"`
	// Optional human-readable hypothesis or summary. In GraphQL this is backed by the
	// experiment hypothesis field.
	Description string `json:"description" api:"nullable"`
	// For redirect variants, whether the original page query string should be
	// forwarded onto the redirect URL.
	IncludeQueryString bool `json:"includeQueryString" api:"nullable"`
	// Configured success metrics. The read shape mirrors the write shape — `metrics`
	// from a GET response can be PATCHed back without modification.
	Metrics ExperimentListResponseMetrics `json:"metrics" api:"nullable"`
	// ISO-8601 timestamp when the experiment most recently entered a running state.
	StartedAt string `json:"startedAt" api:"nullable"`
	// ISO-8601 timestamp when the experiment was completed, if it has been stopped.
	StoppedAt string `json:"stoppedAt" api:"nullable"`
	// Eligibility rules: URL-pattern globs, optional audience, query-param conditions,
	// visitor status, and (server-side) visitor properties. Same shape as the
	// create/patch input.
	TargetingRules ExperimentListResponseTargetingRules `json:"targetingRules" api:"nullable"`
	// Experiment mode. `ab` and `multivariate` use traffic allocation and results;
	// `personalization` is always-on targeting.
	//
	// Any of "ab", "multivariate", "personalization".
	Type ExperimentListResponseType `json:"type" api:"nullable"`
	// ISO-8601 timestamp for the last persisted update, if any.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// Variant ID persisted as the winner when the experiment was stopped. Set via
	// `POST /experiments/{id}/stop` with a `winnerVariantId` body field.
	WinnerVariantID string `json:"winnerVariantId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Key                respjson.Field
		Name               respjson.Field
		Status             respjson.Field
		TrafficAllocation  respjson.Field
		Description        respjson.Field
		IncludeQueryString respjson.Field
		Metrics            respjson.Field
		StartedAt          respjson.Field
		StoppedAt          respjson.Field
		TargetingRules     respjson.Field
		Type               respjson.Field
		UpdatedAt          respjson.Field
		WinnerVariantID    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentListResponse) RawJSON() string { return r.JSON.raw }
func (r *ExperimentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lifecycle state. `draft` is editable, `running` is active, `paused` is
// temporarily inactive, and `completed` is permanently stopped.
type ExperimentListResponseStatus string

const (
	ExperimentListResponseStatusCompleted ExperimentListResponseStatus = "completed"
	ExperimentListResponseStatusDraft     ExperimentListResponseStatus = "draft"
	ExperimentListResponseStatusPaused    ExperimentListResponseStatus = "paused"
	ExperimentListResponseStatusRunning   ExperimentListResponseStatus = "running"
)

// Configured success metrics. The read shape mirrors the write shape — `metrics`
// from a GET response can be PATCHed back without modification.
type ExperimentListResponseMetrics struct {
	// Primary success metric used in the results report.
	Primary any `json:"primary" api:"nullable"`
	// Optional secondary metrics tracked alongside the primary goal.
	Secondary []ExperimentListResponseMetricsSecondary `json:"secondary" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Primary     respjson.Field
		Secondary   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentListResponseMetrics) RawJSON() string { return r.JSON.raw }
func (r *ExperimentListResponseMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentListResponseMetricsSecondary struct {
	// Name of the event used to measure success for this metric.
	EventName string `json:"eventName" api:"nullable"`
	// Optional funnel identifier when the metric is derived from an existing funnel
	// definition.
	FunnelID string `json:"funnelId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventName   respjson.Field
		FunnelID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentListResponseMetricsSecondary) RawJSON() string { return r.JSON.raw }
func (r *ExperimentListResponseMetricsSecondary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Eligibility rules: URL-pattern globs, optional audience, query-param conditions,
// visitor status, and (server-side) visitor properties. Same shape as the
// create/patch input.
type ExperimentListResponseTargetingRules struct {
	// Glob-style URL patterns that must match for the experiment to be eligible. Up to
	// 200 patterns; each pattern up to 2000 characters. An empty array (or omitting
	// the field) matches all URLs — equivalent to `["**"]`.
	URLPatterns []string `json:"urlPatterns" api:"required"`
	// Optional audience identifier used for server-side eligibility filtering.
	AudienceID string `json:"audienceId" api:"nullable"`
	// Additional query-string conditions that must all match for the visitor to
	// qualify.
	QueryParams []ExperimentListResponseTargetingRulesQueryParam `json:"queryParams" api:"nullable"`
	// Optional visitor-property matching rules. These are passed through as JSON for
	// experimentation targeting.
	VisitorProperties any `json:"visitorProperties" api:"nullable"`
	// Whether the experiment should target new visitors, returning visitors, or any
	// visitor.
	VisitorStatus string `json:"visitorStatus" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URLPatterns       respjson.Field
		AudienceID        respjson.Field
		QueryParams       respjson.Field
		VisitorProperties respjson.Field
		VisitorStatus     respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentListResponseTargetingRules) RawJSON() string { return r.JSON.raw }
func (r *ExperimentListResponseTargetingRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentListResponseTargetingRulesQueryParam struct {
	// Query string key to inspect on the current page URL.
	Key string `json:"key" api:"required"`
	// Comparison operator applied to the query string value.
	//
	// Any of "contains", "equals", "exists", "not_equals", "not_exists".
	Operator string `json:"operator" api:"required"`
	// Comparison value used by operators that require one. Omit for `exists` and
	// `not_exists`.
	Value string `json:"value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentListResponseTargetingRulesQueryParam) RawJSON() string { return r.JSON.raw }
func (r *ExperimentListResponseTargetingRulesQueryParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Experiment mode. `ab` and `multivariate` use traffic allocation and results;
// `personalization` is always-on targeting.
type ExperimentListResponseType string

const (
	ExperimentListResponseTypeAb              ExperimentListResponseType = "ab"
	ExperimentListResponseTypeMultivariate    ExperimentListResponseType = "multivariate"
	ExperimentListResponseTypePersonalization ExperimentListResponseType = "personalization"
)

type ExperimentNewResponse struct {
	// Unique identifier for the experiment.
	ID string `json:"id" api:"required"`
	// ISO-8601 timestamp when the experiment was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Stable code-facing key for the experiment. Use this with the headless SDK
	// `getExperimentByKey()` API instead of hard-coding opaque experiment IDs into
	// application code.
	Key string `json:"key" api:"required"`
	// Short, human-readable experiment name.
	Name string `json:"name" api:"required"`
	// Lifecycle state. `draft` is editable, `running` is active, `paused` is
	// temporarily inactive, and `completed` is permanently stopped.
	//
	// Any of "completed", "draft", "paused", "running".
	Status ExperimentNewResponseStatus `json:"status" api:"required"`
	// Percent of eligible traffic assigned into the experiment. Use 0 to fully disable
	// enrollment without deleting the experiment.
	TrafficAllocation int64 `json:"trafficAllocation" api:"required"`
	// All persisted variants for this experiment, including the control variant. A
	// non-personalization experiment needs at least two variants before it can be
	// started.
	Variants []ExperimentNewResponseVariant `json:"variants" api:"required"`
	// Optional human-readable hypothesis or summary. In GraphQL this is backed by the
	// experiment hypothesis field.
	Description string `json:"description" api:"nullable"`
	// For redirect variants, whether the original page query string should be
	// forwarded onto the redirect URL.
	IncludeQueryString bool `json:"includeQueryString" api:"nullable"`
	// Configured success metrics. The read shape mirrors the write shape — `metrics`
	// from a GET response can be PATCHed back without modification.
	Metrics ExperimentNewResponseMetrics `json:"metrics" api:"nullable"`
	// ISO-8601 timestamp when the experiment most recently entered a running state.
	StartedAt string `json:"startedAt" api:"nullable"`
	// ISO-8601 timestamp when the experiment was completed, if it has been stopped.
	StoppedAt string `json:"stoppedAt" api:"nullable"`
	// Eligibility rules: URL-pattern globs, optional audience, query-param conditions,
	// visitor status, and (server-side) visitor properties. Same shape as the
	// create/patch input.
	TargetingRules ExperimentNewResponseTargetingRules `json:"targetingRules" api:"nullable"`
	// Experiment mode. `ab` and `multivariate` use traffic allocation and results;
	// `personalization` is always-on targeting.
	//
	// Any of "ab", "multivariate", "personalization".
	Type ExperimentNewResponseType `json:"type" api:"nullable"`
	// ISO-8601 timestamp for the last persisted update, if any.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// Variant ID persisted as the winner when the experiment was stopped. Set via
	// `POST /experiments/{id}/stop` with a `winnerVariantId` body field.
	WinnerVariantID string `json:"winnerVariantId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Key                respjson.Field
		Name               respjson.Field
		Status             respjson.Field
		TrafficAllocation  respjson.Field
		Variants           respjson.Field
		Description        respjson.Field
		IncludeQueryString respjson.Field
		Metrics            respjson.Field
		StartedAt          respjson.Field
		StoppedAt          respjson.Field
		TargetingRules     respjson.Field
		Type               respjson.Field
		UpdatedAt          respjson.Field
		WinnerVariantID    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ExperimentNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lifecycle state. `draft` is editable, `running` is active, `paused` is
// temporarily inactive, and `completed` is permanently stopped.
type ExperimentNewResponseStatus string

const (
	ExperimentNewResponseStatusCompleted ExperimentNewResponseStatus = "completed"
	ExperimentNewResponseStatusDraft     ExperimentNewResponseStatus = "draft"
	ExperimentNewResponseStatusPaused    ExperimentNewResponseStatus = "paused"
	ExperimentNewResponseStatusRunning   ExperimentNewResponseStatus = "running"
)

type ExperimentNewResponseVariant struct {
	// Unique identifier for this experiment variant.
	ID string `json:"id" api:"required"`
	// Parent experiment ID this variant belongs to.
	ExperimentID string `json:"experimentId" api:"required"`
	// Whether this is the baseline control variant.
	IsControl bool `json:"isControl" api:"required"`
	// Human-readable variant name shown in the dashboard and results.
	Name string `json:"name" api:"required"`
	// Relative traffic weight used when assigning visitors among variants in an active
	// experiment.
	Weight int64 `json:"weight" api:"required"`
	// Ordered list of declarative DOM mutations applied when this variant is assigned.
	DomModifications []ExperimentNewResponseVariantDomModification `json:"domModifications" api:"nullable"`
	// Target URL for redirect variants. Use either a site-relative path such as
	// `/pricing-v2` or an absolute `https://` URL. Cross-origin `http://` URLs are
	// rejected. Omit for DOM modification variants.
	RedirectURL string `json:"redirectUrl" api:"nullable"`
	// How this variant changes the user experience. `dom_modifications` for on-page
	// changes or `redirect` for redirect tests.
	VariantType string `json:"variantType" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		ExperimentID     respjson.Field
		IsControl        respjson.Field
		Name             respjson.Field
		Weight           respjson.Field
		DomModifications respjson.Field
		RedirectURL      respjson.Field
		VariantType      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentNewResponseVariant) RawJSON() string { return r.JSON.raw }
func (r *ExperimentNewResponseVariant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentNewResponseVariantDomModification struct {
	// Mutation to apply when the selector matches. Use `redirectUrl` instead of DOM
	// modifications for redirect variants.
	//
	// Any of "customCss", "customJs", "insertAfter", "insertBefore", "remove",
	// "setAttribute", "setHtml", "setImage", "setStyle", "setText".
	Action string `json:"action" api:"required"`
	// CSS selector used to find the element to modify on the page at runtime.
	Selector string `json:"selector" api:"required"`
	// Canonical action payload. For `setText` / `setHtml` / `customCss` / `customJs` /
	// `setImage` / `insertBefore` / `insertAfter` this is the literal
	// text/HTML/CSS/JS/URL. For `setStyle` and `setAttribute` it is a JSON-stringified
	// `{key: value}` object — prefer the structured `styles` / `attribute` fields
	// below to avoid manual JSON encoding.
	Value string `json:"value" api:"required"`
	// Populated on read for `setAttribute` modifications, parsed from `value`.
	// Customers may also send this field instead of a JSON-stringified `value` on
	// write — see `domModificationInputSchema`.
	Attribute any `json:"attribute" api:"nullable"`
	// Populated on read for `setStyle` modifications, parsed from `value`. Customers
	// may also send this field instead of a JSON-stringified `value` on write — see
	// `domModificationInputSchema`.
	Styles []ExperimentNewResponseVariantDomModificationStyle `json:"styles" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Selector    respjson.Field
		Value       respjson.Field
		Attribute   respjson.Field
		Styles      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentNewResponseVariantDomModification) RawJSON() string { return r.JSON.raw }
func (r *ExperimentNewResponseVariantDomModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentNewResponseVariantDomModificationStyle struct {
	// CSS property name in camelCase or kebab-case.
	Property string `json:"property" api:"required"`
	// CSS value to assign to the property.
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Property    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentNewResponseVariantDomModificationStyle) RawJSON() string { return r.JSON.raw }
func (r *ExperimentNewResponseVariantDomModificationStyle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configured success metrics. The read shape mirrors the write shape — `metrics`
// from a GET response can be PATCHed back without modification.
type ExperimentNewResponseMetrics struct {
	// Primary success metric used in the results report.
	Primary any `json:"primary" api:"nullable"`
	// Optional secondary metrics tracked alongside the primary goal.
	Secondary []ExperimentNewResponseMetricsSecondary `json:"secondary" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Primary     respjson.Field
		Secondary   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentNewResponseMetrics) RawJSON() string { return r.JSON.raw }
func (r *ExperimentNewResponseMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentNewResponseMetricsSecondary struct {
	// Name of the event used to measure success for this metric.
	EventName string `json:"eventName" api:"nullable"`
	// Optional funnel identifier when the metric is derived from an existing funnel
	// definition.
	FunnelID string `json:"funnelId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventName   respjson.Field
		FunnelID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentNewResponseMetricsSecondary) RawJSON() string { return r.JSON.raw }
func (r *ExperimentNewResponseMetricsSecondary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Eligibility rules: URL-pattern globs, optional audience, query-param conditions,
// visitor status, and (server-side) visitor properties. Same shape as the
// create/patch input.
type ExperimentNewResponseTargetingRules struct {
	// Glob-style URL patterns that must match for the experiment to be eligible. Up to
	// 200 patterns; each pattern up to 2000 characters. An empty array (or omitting
	// the field) matches all URLs — equivalent to `["**"]`.
	URLPatterns []string `json:"urlPatterns" api:"required"`
	// Optional audience identifier used for server-side eligibility filtering.
	AudienceID string `json:"audienceId" api:"nullable"`
	// Additional query-string conditions that must all match for the visitor to
	// qualify.
	QueryParams []ExperimentNewResponseTargetingRulesQueryParam `json:"queryParams" api:"nullable"`
	// Optional visitor-property matching rules. These are passed through as JSON for
	// experimentation targeting.
	VisitorProperties any `json:"visitorProperties" api:"nullable"`
	// Whether the experiment should target new visitors, returning visitors, or any
	// visitor.
	VisitorStatus string `json:"visitorStatus" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URLPatterns       respjson.Field
		AudienceID        respjson.Field
		QueryParams       respjson.Field
		VisitorProperties respjson.Field
		VisitorStatus     respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentNewResponseTargetingRules) RawJSON() string { return r.JSON.raw }
func (r *ExperimentNewResponseTargetingRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentNewResponseTargetingRulesQueryParam struct {
	// Query string key to inspect on the current page URL.
	Key string `json:"key" api:"required"`
	// Comparison operator applied to the query string value.
	//
	// Any of "contains", "equals", "exists", "not_equals", "not_exists".
	Operator string `json:"operator" api:"required"`
	// Comparison value used by operators that require one. Omit for `exists` and
	// `not_exists`.
	Value string `json:"value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentNewResponseTargetingRulesQueryParam) RawJSON() string { return r.JSON.raw }
func (r *ExperimentNewResponseTargetingRulesQueryParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Experiment mode. `ab` and `multivariate` use traffic allocation and results;
// `personalization` is always-on targeting.
type ExperimentNewResponseType string

const (
	ExperimentNewResponseTypeAb              ExperimentNewResponseType = "ab"
	ExperimentNewResponseTypeMultivariate    ExperimentNewResponseType = "multivariate"
	ExperimentNewResponseTypePersonalization ExperimentNewResponseType = "personalization"
)

type ExperimentGetResponse struct {
	// Unique identifier for the experiment.
	ID string `json:"id" api:"required"`
	// ISO-8601 timestamp when the experiment was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Stable code-facing key for the experiment. Use this with the headless SDK
	// `getExperimentByKey()` API instead of hard-coding opaque experiment IDs into
	// application code.
	Key string `json:"key" api:"required"`
	// Short, human-readable experiment name.
	Name string `json:"name" api:"required"`
	// Lifecycle state. `draft` is editable, `running` is active, `paused` is
	// temporarily inactive, and `completed` is permanently stopped.
	//
	// Any of "completed", "draft", "paused", "running".
	Status ExperimentGetResponseStatus `json:"status" api:"required"`
	// Percent of eligible traffic assigned into the experiment. Use 0 to fully disable
	// enrollment without deleting the experiment.
	TrafficAllocation int64 `json:"trafficAllocation" api:"required"`
	// All persisted variants for this experiment, including the control variant. A
	// non-personalization experiment needs at least two variants before it can be
	// started.
	Variants []ExperimentGetResponseVariant `json:"variants" api:"required"`
	// Optional human-readable hypothesis or summary. In GraphQL this is backed by the
	// experiment hypothesis field.
	Description string `json:"description" api:"nullable"`
	// For redirect variants, whether the original page query string should be
	// forwarded onto the redirect URL.
	IncludeQueryString bool `json:"includeQueryString" api:"nullable"`
	// Configured success metrics. The read shape mirrors the write shape — `metrics`
	// from a GET response can be PATCHed back without modification.
	Metrics ExperimentGetResponseMetrics `json:"metrics" api:"nullable"`
	// ISO-8601 timestamp when the experiment most recently entered a running state.
	StartedAt string `json:"startedAt" api:"nullable"`
	// ISO-8601 timestamp when the experiment was completed, if it has been stopped.
	StoppedAt string `json:"stoppedAt" api:"nullable"`
	// Eligibility rules: URL-pattern globs, optional audience, query-param conditions,
	// visitor status, and (server-side) visitor properties. Same shape as the
	// create/patch input.
	TargetingRules ExperimentGetResponseTargetingRules `json:"targetingRules" api:"nullable"`
	// Experiment mode. `ab` and `multivariate` use traffic allocation and results;
	// `personalization` is always-on targeting.
	//
	// Any of "ab", "multivariate", "personalization".
	Type ExperimentGetResponseType `json:"type" api:"nullable"`
	// ISO-8601 timestamp for the last persisted update, if any.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// Variant ID persisted as the winner when the experiment was stopped. Set via
	// `POST /experiments/{id}/stop` with a `winnerVariantId` body field.
	WinnerVariantID string `json:"winnerVariantId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Key                respjson.Field
		Name               respjson.Field
		Status             respjson.Field
		TrafficAllocation  respjson.Field
		Variants           respjson.Field
		Description        respjson.Field
		IncludeQueryString respjson.Field
		Metrics            respjson.Field
		StartedAt          respjson.Field
		StoppedAt          respjson.Field
		TargetingRules     respjson.Field
		Type               respjson.Field
		UpdatedAt          respjson.Field
		WinnerVariantID    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ExperimentGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lifecycle state. `draft` is editable, `running` is active, `paused` is
// temporarily inactive, and `completed` is permanently stopped.
type ExperimentGetResponseStatus string

const (
	ExperimentGetResponseStatusCompleted ExperimentGetResponseStatus = "completed"
	ExperimentGetResponseStatusDraft     ExperimentGetResponseStatus = "draft"
	ExperimentGetResponseStatusPaused    ExperimentGetResponseStatus = "paused"
	ExperimentGetResponseStatusRunning   ExperimentGetResponseStatus = "running"
)

type ExperimentGetResponseVariant struct {
	// Unique identifier for this experiment variant.
	ID string `json:"id" api:"required"`
	// Parent experiment ID this variant belongs to.
	ExperimentID string `json:"experimentId" api:"required"`
	// Whether this is the baseline control variant.
	IsControl bool `json:"isControl" api:"required"`
	// Human-readable variant name shown in the dashboard and results.
	Name string `json:"name" api:"required"`
	// Relative traffic weight used when assigning visitors among variants in an active
	// experiment.
	Weight int64 `json:"weight" api:"required"`
	// Ordered list of declarative DOM mutations applied when this variant is assigned.
	DomModifications []ExperimentGetResponseVariantDomModification `json:"domModifications" api:"nullable"`
	// Target URL for redirect variants. Use either a site-relative path such as
	// `/pricing-v2` or an absolute `https://` URL. Cross-origin `http://` URLs are
	// rejected. Omit for DOM modification variants.
	RedirectURL string `json:"redirectUrl" api:"nullable"`
	// How this variant changes the user experience. `dom_modifications` for on-page
	// changes or `redirect` for redirect tests.
	VariantType string `json:"variantType" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		ExperimentID     respjson.Field
		IsControl        respjson.Field
		Name             respjson.Field
		Weight           respjson.Field
		DomModifications respjson.Field
		RedirectURL      respjson.Field
		VariantType      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentGetResponseVariant) RawJSON() string { return r.JSON.raw }
func (r *ExperimentGetResponseVariant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentGetResponseVariantDomModification struct {
	// Mutation to apply when the selector matches. Use `redirectUrl` instead of DOM
	// modifications for redirect variants.
	//
	// Any of "customCss", "customJs", "insertAfter", "insertBefore", "remove",
	// "setAttribute", "setHtml", "setImage", "setStyle", "setText".
	Action string `json:"action" api:"required"`
	// CSS selector used to find the element to modify on the page at runtime.
	Selector string `json:"selector" api:"required"`
	// Canonical action payload. For `setText` / `setHtml` / `customCss` / `customJs` /
	// `setImage` / `insertBefore` / `insertAfter` this is the literal
	// text/HTML/CSS/JS/URL. For `setStyle` and `setAttribute` it is a JSON-stringified
	// `{key: value}` object — prefer the structured `styles` / `attribute` fields
	// below to avoid manual JSON encoding.
	Value string `json:"value" api:"required"`
	// Populated on read for `setAttribute` modifications, parsed from `value`.
	// Customers may also send this field instead of a JSON-stringified `value` on
	// write — see `domModificationInputSchema`.
	Attribute any `json:"attribute" api:"nullable"`
	// Populated on read for `setStyle` modifications, parsed from `value`. Customers
	// may also send this field instead of a JSON-stringified `value` on write — see
	// `domModificationInputSchema`.
	Styles []ExperimentGetResponseVariantDomModificationStyle `json:"styles" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Selector    respjson.Field
		Value       respjson.Field
		Attribute   respjson.Field
		Styles      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentGetResponseVariantDomModification) RawJSON() string { return r.JSON.raw }
func (r *ExperimentGetResponseVariantDomModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentGetResponseVariantDomModificationStyle struct {
	// CSS property name in camelCase or kebab-case.
	Property string `json:"property" api:"required"`
	// CSS value to assign to the property.
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Property    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentGetResponseVariantDomModificationStyle) RawJSON() string { return r.JSON.raw }
func (r *ExperimentGetResponseVariantDomModificationStyle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configured success metrics. The read shape mirrors the write shape — `metrics`
// from a GET response can be PATCHed back without modification.
type ExperimentGetResponseMetrics struct {
	// Primary success metric used in the results report.
	Primary any `json:"primary" api:"nullable"`
	// Optional secondary metrics tracked alongside the primary goal.
	Secondary []ExperimentGetResponseMetricsSecondary `json:"secondary" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Primary     respjson.Field
		Secondary   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentGetResponseMetrics) RawJSON() string { return r.JSON.raw }
func (r *ExperimentGetResponseMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentGetResponseMetricsSecondary struct {
	// Name of the event used to measure success for this metric.
	EventName string `json:"eventName" api:"nullable"`
	// Optional funnel identifier when the metric is derived from an existing funnel
	// definition.
	FunnelID string `json:"funnelId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventName   respjson.Field
		FunnelID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentGetResponseMetricsSecondary) RawJSON() string { return r.JSON.raw }
func (r *ExperimentGetResponseMetricsSecondary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Eligibility rules: URL-pattern globs, optional audience, query-param conditions,
// visitor status, and (server-side) visitor properties. Same shape as the
// create/patch input.
type ExperimentGetResponseTargetingRules struct {
	// Glob-style URL patterns that must match for the experiment to be eligible. Up to
	// 200 patterns; each pattern up to 2000 characters. An empty array (or omitting
	// the field) matches all URLs — equivalent to `["**"]`.
	URLPatterns []string `json:"urlPatterns" api:"required"`
	// Optional audience identifier used for server-side eligibility filtering.
	AudienceID string `json:"audienceId" api:"nullable"`
	// Additional query-string conditions that must all match for the visitor to
	// qualify.
	QueryParams []ExperimentGetResponseTargetingRulesQueryParam `json:"queryParams" api:"nullable"`
	// Optional visitor-property matching rules. These are passed through as JSON for
	// experimentation targeting.
	VisitorProperties any `json:"visitorProperties" api:"nullable"`
	// Whether the experiment should target new visitors, returning visitors, or any
	// visitor.
	VisitorStatus string `json:"visitorStatus" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URLPatterns       respjson.Field
		AudienceID        respjson.Field
		QueryParams       respjson.Field
		VisitorProperties respjson.Field
		VisitorStatus     respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentGetResponseTargetingRules) RawJSON() string { return r.JSON.raw }
func (r *ExperimentGetResponseTargetingRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentGetResponseTargetingRulesQueryParam struct {
	// Query string key to inspect on the current page URL.
	Key string `json:"key" api:"required"`
	// Comparison operator applied to the query string value.
	//
	// Any of "contains", "equals", "exists", "not_equals", "not_exists".
	Operator string `json:"operator" api:"required"`
	// Comparison value used by operators that require one. Omit for `exists` and
	// `not_exists`.
	Value string `json:"value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentGetResponseTargetingRulesQueryParam) RawJSON() string { return r.JSON.raw }
func (r *ExperimentGetResponseTargetingRulesQueryParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Experiment mode. `ab` and `multivariate` use traffic allocation and results;
// `personalization` is always-on targeting.
type ExperimentGetResponseType string

const (
	ExperimentGetResponseTypeAb              ExperimentGetResponseType = "ab"
	ExperimentGetResponseTypeMultivariate    ExperimentGetResponseType = "multivariate"
	ExperimentGetResponseTypePersonalization ExperimentGetResponseType = "personalization"
)

type ExperimentUpdateResponse struct {
	// Unique identifier for the experiment.
	ID string `json:"id" api:"required"`
	// ISO-8601 timestamp when the experiment was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Stable code-facing key for the experiment. Use this with the headless SDK
	// `getExperimentByKey()` API instead of hard-coding opaque experiment IDs into
	// application code.
	Key string `json:"key" api:"required"`
	// Short, human-readable experiment name.
	Name string `json:"name" api:"required"`
	// Lifecycle state. `draft` is editable, `running` is active, `paused` is
	// temporarily inactive, and `completed` is permanently stopped.
	//
	// Any of "completed", "draft", "paused", "running".
	Status ExperimentUpdateResponseStatus `json:"status" api:"required"`
	// Percent of eligible traffic assigned into the experiment. Use 0 to fully disable
	// enrollment without deleting the experiment.
	TrafficAllocation int64 `json:"trafficAllocation" api:"required"`
	// Optional human-readable hypothesis or summary. In GraphQL this is backed by the
	// experiment hypothesis field.
	Description string `json:"description" api:"nullable"`
	// For redirect variants, whether the original page query string should be
	// forwarded onto the redirect URL.
	IncludeQueryString bool `json:"includeQueryString" api:"nullable"`
	// Configured success metrics. The read shape mirrors the write shape — `metrics`
	// from a GET response can be PATCHed back without modification.
	Metrics ExperimentUpdateResponseMetrics `json:"metrics" api:"nullable"`
	// ISO-8601 timestamp when the experiment most recently entered a running state.
	StartedAt string `json:"startedAt" api:"nullable"`
	// ISO-8601 timestamp when the experiment was completed, if it has been stopped.
	StoppedAt string `json:"stoppedAt" api:"nullable"`
	// Eligibility rules: URL-pattern globs, optional audience, query-param conditions,
	// visitor status, and (server-side) visitor properties. Same shape as the
	// create/patch input.
	TargetingRules ExperimentUpdateResponseTargetingRules `json:"targetingRules" api:"nullable"`
	// Experiment mode. `ab` and `multivariate` use traffic allocation and results;
	// `personalization` is always-on targeting.
	//
	// Any of "ab", "multivariate", "personalization".
	Type ExperimentUpdateResponseType `json:"type" api:"nullable"`
	// ISO-8601 timestamp for the last persisted update, if any.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// Variant ID persisted as the winner when the experiment was stopped. Set via
	// `POST /experiments/{id}/stop` with a `winnerVariantId` body field.
	WinnerVariantID string `json:"winnerVariantId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Key                respjson.Field
		Name               respjson.Field
		Status             respjson.Field
		TrafficAllocation  respjson.Field
		Description        respjson.Field
		IncludeQueryString respjson.Field
		Metrics            respjson.Field
		StartedAt          respjson.Field
		StoppedAt          respjson.Field
		TargetingRules     respjson.Field
		Type               respjson.Field
		UpdatedAt          respjson.Field
		WinnerVariantID    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ExperimentUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lifecycle state. `draft` is editable, `running` is active, `paused` is
// temporarily inactive, and `completed` is permanently stopped.
type ExperimentUpdateResponseStatus string

const (
	ExperimentUpdateResponseStatusCompleted ExperimentUpdateResponseStatus = "completed"
	ExperimentUpdateResponseStatusDraft     ExperimentUpdateResponseStatus = "draft"
	ExperimentUpdateResponseStatusPaused    ExperimentUpdateResponseStatus = "paused"
	ExperimentUpdateResponseStatusRunning   ExperimentUpdateResponseStatus = "running"
)

// Configured success metrics. The read shape mirrors the write shape — `metrics`
// from a GET response can be PATCHed back without modification.
type ExperimentUpdateResponseMetrics struct {
	// Primary success metric used in the results report.
	Primary any `json:"primary" api:"nullable"`
	// Optional secondary metrics tracked alongside the primary goal.
	Secondary []ExperimentUpdateResponseMetricsSecondary `json:"secondary" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Primary     respjson.Field
		Secondary   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentUpdateResponseMetrics) RawJSON() string { return r.JSON.raw }
func (r *ExperimentUpdateResponseMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentUpdateResponseMetricsSecondary struct {
	// Name of the event used to measure success for this metric.
	EventName string `json:"eventName" api:"nullable"`
	// Optional funnel identifier when the metric is derived from an existing funnel
	// definition.
	FunnelID string `json:"funnelId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventName   respjson.Field
		FunnelID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentUpdateResponseMetricsSecondary) RawJSON() string { return r.JSON.raw }
func (r *ExperimentUpdateResponseMetricsSecondary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Eligibility rules: URL-pattern globs, optional audience, query-param conditions,
// visitor status, and (server-side) visitor properties. Same shape as the
// create/patch input.
type ExperimentUpdateResponseTargetingRules struct {
	// Glob-style URL patterns that must match for the experiment to be eligible. Up to
	// 200 patterns; each pattern up to 2000 characters. An empty array (or omitting
	// the field) matches all URLs — equivalent to `["**"]`.
	URLPatterns []string `json:"urlPatterns" api:"required"`
	// Optional audience identifier used for server-side eligibility filtering.
	AudienceID string `json:"audienceId" api:"nullable"`
	// Additional query-string conditions that must all match for the visitor to
	// qualify.
	QueryParams []ExperimentUpdateResponseTargetingRulesQueryParam `json:"queryParams" api:"nullable"`
	// Optional visitor-property matching rules. These are passed through as JSON for
	// experimentation targeting.
	VisitorProperties any `json:"visitorProperties" api:"nullable"`
	// Whether the experiment should target new visitors, returning visitors, or any
	// visitor.
	VisitorStatus string `json:"visitorStatus" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URLPatterns       respjson.Field
		AudienceID        respjson.Field
		QueryParams       respjson.Field
		VisitorProperties respjson.Field
		VisitorStatus     respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentUpdateResponseTargetingRules) RawJSON() string { return r.JSON.raw }
func (r *ExperimentUpdateResponseTargetingRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentUpdateResponseTargetingRulesQueryParam struct {
	// Query string key to inspect on the current page URL.
	Key string `json:"key" api:"required"`
	// Comparison operator applied to the query string value.
	//
	// Any of "contains", "equals", "exists", "not_equals", "not_exists".
	Operator string `json:"operator" api:"required"`
	// Comparison value used by operators that require one. Omit for `exists` and
	// `not_exists`.
	Value string `json:"value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentUpdateResponseTargetingRulesQueryParam) RawJSON() string { return r.JSON.raw }
func (r *ExperimentUpdateResponseTargetingRulesQueryParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Experiment mode. `ab` and `multivariate` use traffic allocation and results;
// `personalization` is always-on targeting.
type ExperimentUpdateResponseType string

const (
	ExperimentUpdateResponseTypeAb              ExperimentUpdateResponseType = "ab"
	ExperimentUpdateResponseTypeMultivariate    ExperimentUpdateResponseType = "multivariate"
	ExperimentUpdateResponseTypePersonalization ExperimentUpdateResponseType = "personalization"
)

type ExperimentStartResponse struct {
	// Number of unpublished version changes detected at the time of the lifecycle
	// action. Values greater than 1 mean other unpublished work exists besides this
	// experiment state change.
	ConcurrentVersionChanges int64                             `json:"concurrentVersionChanges" api:"required"`
	Experiment               ExperimentStartResponseExperiment `json:"experiment" api:"required"`
	// Whether the status change was also published to the current live version
	// immediately.
	//
	// Any of "pending_publish", "published".
	PublishStatus ExperimentStartResponsePublishStatus `json:"publishStatus" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConcurrentVersionChanges respjson.Field
		Experiment               respjson.Field
		PublishStatus            respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentStartResponse) RawJSON() string { return r.JSON.raw }
func (r *ExperimentStartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentStartResponseExperiment struct {
	// Unique identifier for the experiment.
	ID string `json:"id" api:"required"`
	// ISO-8601 timestamp when the experiment was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Stable code-facing key for the experiment. Use this with the headless SDK
	// `getExperimentByKey()` API instead of hard-coding opaque experiment IDs into
	// application code.
	Key string `json:"key" api:"required"`
	// Short, human-readable experiment name.
	Name string `json:"name" api:"required"`
	// Lifecycle state. `draft` is editable, `running` is active, `paused` is
	// temporarily inactive, and `completed` is permanently stopped.
	//
	// Any of "completed", "draft", "paused", "running".
	Status string `json:"status" api:"required"`
	// Percent of eligible traffic assigned into the experiment. Use 0 to fully disable
	// enrollment without deleting the experiment.
	TrafficAllocation int64 `json:"trafficAllocation" api:"required"`
	// Optional human-readable hypothesis or summary. In GraphQL this is backed by the
	// experiment hypothesis field.
	Description string `json:"description" api:"nullable"`
	// For redirect variants, whether the original page query string should be
	// forwarded onto the redirect URL.
	IncludeQueryString bool `json:"includeQueryString" api:"nullable"`
	// Configured success metrics. The read shape mirrors the write shape — `metrics`
	// from a GET response can be PATCHed back without modification.
	Metrics ExperimentStartResponseExperimentMetrics `json:"metrics" api:"nullable"`
	// ISO-8601 timestamp when the experiment most recently entered a running state.
	StartedAt string `json:"startedAt" api:"nullable"`
	// ISO-8601 timestamp when the experiment was completed, if it has been stopped.
	StoppedAt string `json:"stoppedAt" api:"nullable"`
	// Eligibility rules: URL-pattern globs, optional audience, query-param conditions,
	// visitor status, and (server-side) visitor properties. Same shape as the
	// create/patch input.
	TargetingRules ExperimentStartResponseExperimentTargetingRules `json:"targetingRules" api:"nullable"`
	// Experiment mode. `ab` and `multivariate` use traffic allocation and results;
	// `personalization` is always-on targeting.
	//
	// Any of "ab", "multivariate", "personalization".
	Type string `json:"type" api:"nullable"`
	// ISO-8601 timestamp for the last persisted update, if any.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// Variant ID persisted as the winner when the experiment was stopped. Set via
	// `POST /experiments/{id}/stop` with a `winnerVariantId` body field.
	WinnerVariantID string `json:"winnerVariantId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Key                respjson.Field
		Name               respjson.Field
		Status             respjson.Field
		TrafficAllocation  respjson.Field
		Description        respjson.Field
		IncludeQueryString respjson.Field
		Metrics            respjson.Field
		StartedAt          respjson.Field
		StoppedAt          respjson.Field
		TargetingRules     respjson.Field
		Type               respjson.Field
		UpdatedAt          respjson.Field
		WinnerVariantID    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentStartResponseExperiment) RawJSON() string { return r.JSON.raw }
func (r *ExperimentStartResponseExperiment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configured success metrics. The read shape mirrors the write shape — `metrics`
// from a GET response can be PATCHed back without modification.
type ExperimentStartResponseExperimentMetrics struct {
	// Primary success metric used in the results report.
	Primary any `json:"primary" api:"nullable"`
	// Optional secondary metrics tracked alongside the primary goal.
	Secondary []ExperimentStartResponseExperimentMetricsSecondary `json:"secondary" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Primary     respjson.Field
		Secondary   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentStartResponseExperimentMetrics) RawJSON() string { return r.JSON.raw }
func (r *ExperimentStartResponseExperimentMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentStartResponseExperimentMetricsSecondary struct {
	// Name of the event used to measure success for this metric.
	EventName string `json:"eventName" api:"nullable"`
	// Optional funnel identifier when the metric is derived from an existing funnel
	// definition.
	FunnelID string `json:"funnelId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventName   respjson.Field
		FunnelID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentStartResponseExperimentMetricsSecondary) RawJSON() string { return r.JSON.raw }
func (r *ExperimentStartResponseExperimentMetricsSecondary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Eligibility rules: URL-pattern globs, optional audience, query-param conditions,
// visitor status, and (server-side) visitor properties. Same shape as the
// create/patch input.
type ExperimentStartResponseExperimentTargetingRules struct {
	// Glob-style URL patterns that must match for the experiment to be eligible. Up to
	// 200 patterns; each pattern up to 2000 characters. An empty array (or omitting
	// the field) matches all URLs — equivalent to `["**"]`.
	URLPatterns []string `json:"urlPatterns" api:"required"`
	// Optional audience identifier used for server-side eligibility filtering.
	AudienceID string `json:"audienceId" api:"nullable"`
	// Additional query-string conditions that must all match for the visitor to
	// qualify.
	QueryParams []ExperimentStartResponseExperimentTargetingRulesQueryParam `json:"queryParams" api:"nullable"`
	// Optional visitor-property matching rules. These are passed through as JSON for
	// experimentation targeting.
	VisitorProperties any `json:"visitorProperties" api:"nullable"`
	// Whether the experiment should target new visitors, returning visitors, or any
	// visitor.
	VisitorStatus string `json:"visitorStatus" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URLPatterns       respjson.Field
		AudienceID        respjson.Field
		QueryParams       respjson.Field
		VisitorProperties respjson.Field
		VisitorStatus     respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentStartResponseExperimentTargetingRules) RawJSON() string { return r.JSON.raw }
func (r *ExperimentStartResponseExperimentTargetingRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentStartResponseExperimentTargetingRulesQueryParam struct {
	// Query string key to inspect on the current page URL.
	Key string `json:"key" api:"required"`
	// Comparison operator applied to the query string value.
	//
	// Any of "contains", "equals", "exists", "not_equals", "not_exists".
	Operator string `json:"operator" api:"required"`
	// Comparison value used by operators that require one. Omit for `exists` and
	// `not_exists`.
	Value string `json:"value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentStartResponseExperimentTargetingRulesQueryParam) RawJSON() string {
	return r.JSON.raw
}
func (r *ExperimentStartResponseExperimentTargetingRulesQueryParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the status change was also published to the current live version
// immediately.
type ExperimentStartResponsePublishStatus string

const (
	ExperimentStartResponsePublishStatusPendingPublish ExperimentStartResponsePublishStatus = "pending_publish"
	ExperimentStartResponsePublishStatusPublished      ExperimentStartResponsePublishStatus = "published"
)

type ExperimentStopResponse struct {
	// Number of unpublished version changes detected at the time of the lifecycle
	// action. Values greater than 1 mean other unpublished work exists besides this
	// experiment state change.
	ConcurrentVersionChanges int64                            `json:"concurrentVersionChanges" api:"required"`
	Experiment               ExperimentStopResponseExperiment `json:"experiment" api:"required"`
	// Whether the status change was also published to the current live version
	// immediately.
	//
	// Any of "pending_publish", "published".
	PublishStatus ExperimentStopResponsePublishStatus `json:"publishStatus" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConcurrentVersionChanges respjson.Field
		Experiment               respjson.Field
		PublishStatus            respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentStopResponse) RawJSON() string { return r.JSON.raw }
func (r *ExperimentStopResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentStopResponseExperiment struct {
	// Unique identifier for the experiment.
	ID string `json:"id" api:"required"`
	// ISO-8601 timestamp when the experiment was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Stable code-facing key for the experiment. Use this with the headless SDK
	// `getExperimentByKey()` API instead of hard-coding opaque experiment IDs into
	// application code.
	Key string `json:"key" api:"required"`
	// Short, human-readable experiment name.
	Name string `json:"name" api:"required"`
	// Lifecycle state. `draft` is editable, `running` is active, `paused` is
	// temporarily inactive, and `completed` is permanently stopped.
	//
	// Any of "completed", "draft", "paused", "running".
	Status string `json:"status" api:"required"`
	// Percent of eligible traffic assigned into the experiment. Use 0 to fully disable
	// enrollment without deleting the experiment.
	TrafficAllocation int64 `json:"trafficAllocation" api:"required"`
	// Optional human-readable hypothesis or summary. In GraphQL this is backed by the
	// experiment hypothesis field.
	Description string `json:"description" api:"nullable"`
	// For redirect variants, whether the original page query string should be
	// forwarded onto the redirect URL.
	IncludeQueryString bool `json:"includeQueryString" api:"nullable"`
	// Configured success metrics. The read shape mirrors the write shape — `metrics`
	// from a GET response can be PATCHed back without modification.
	Metrics ExperimentStopResponseExperimentMetrics `json:"metrics" api:"nullable"`
	// ISO-8601 timestamp when the experiment most recently entered a running state.
	StartedAt string `json:"startedAt" api:"nullable"`
	// ISO-8601 timestamp when the experiment was completed, if it has been stopped.
	StoppedAt string `json:"stoppedAt" api:"nullable"`
	// Eligibility rules: URL-pattern globs, optional audience, query-param conditions,
	// visitor status, and (server-side) visitor properties. Same shape as the
	// create/patch input.
	TargetingRules ExperimentStopResponseExperimentTargetingRules `json:"targetingRules" api:"nullable"`
	// Experiment mode. `ab` and `multivariate` use traffic allocation and results;
	// `personalization` is always-on targeting.
	//
	// Any of "ab", "multivariate", "personalization".
	Type string `json:"type" api:"nullable"`
	// ISO-8601 timestamp for the last persisted update, if any.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// Variant ID persisted as the winner when the experiment was stopped. Set via
	// `POST /experiments/{id}/stop` with a `winnerVariantId` body field.
	WinnerVariantID string `json:"winnerVariantId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Key                respjson.Field
		Name               respjson.Field
		Status             respjson.Field
		TrafficAllocation  respjson.Field
		Description        respjson.Field
		IncludeQueryString respjson.Field
		Metrics            respjson.Field
		StartedAt          respjson.Field
		StoppedAt          respjson.Field
		TargetingRules     respjson.Field
		Type               respjson.Field
		UpdatedAt          respjson.Field
		WinnerVariantID    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentStopResponseExperiment) RawJSON() string { return r.JSON.raw }
func (r *ExperimentStopResponseExperiment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configured success metrics. The read shape mirrors the write shape — `metrics`
// from a GET response can be PATCHed back without modification.
type ExperimentStopResponseExperimentMetrics struct {
	// Primary success metric used in the results report.
	Primary any `json:"primary" api:"nullable"`
	// Optional secondary metrics tracked alongside the primary goal.
	Secondary []ExperimentStopResponseExperimentMetricsSecondary `json:"secondary" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Primary     respjson.Field
		Secondary   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentStopResponseExperimentMetrics) RawJSON() string { return r.JSON.raw }
func (r *ExperimentStopResponseExperimentMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentStopResponseExperimentMetricsSecondary struct {
	// Name of the event used to measure success for this metric.
	EventName string `json:"eventName" api:"nullable"`
	// Optional funnel identifier when the metric is derived from an existing funnel
	// definition.
	FunnelID string `json:"funnelId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventName   respjson.Field
		FunnelID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentStopResponseExperimentMetricsSecondary) RawJSON() string { return r.JSON.raw }
func (r *ExperimentStopResponseExperimentMetricsSecondary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Eligibility rules: URL-pattern globs, optional audience, query-param conditions,
// visitor status, and (server-side) visitor properties. Same shape as the
// create/patch input.
type ExperimentStopResponseExperimentTargetingRules struct {
	// Glob-style URL patterns that must match for the experiment to be eligible. Up to
	// 200 patterns; each pattern up to 2000 characters. An empty array (or omitting
	// the field) matches all URLs — equivalent to `["**"]`.
	URLPatterns []string `json:"urlPatterns" api:"required"`
	// Optional audience identifier used for server-side eligibility filtering.
	AudienceID string `json:"audienceId" api:"nullable"`
	// Additional query-string conditions that must all match for the visitor to
	// qualify.
	QueryParams []ExperimentStopResponseExperimentTargetingRulesQueryParam `json:"queryParams" api:"nullable"`
	// Optional visitor-property matching rules. These are passed through as JSON for
	// experimentation targeting.
	VisitorProperties any `json:"visitorProperties" api:"nullable"`
	// Whether the experiment should target new visitors, returning visitors, or any
	// visitor.
	VisitorStatus string `json:"visitorStatus" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URLPatterns       respjson.Field
		AudienceID        respjson.Field
		QueryParams       respjson.Field
		VisitorProperties respjson.Field
		VisitorStatus     respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentStopResponseExperimentTargetingRules) RawJSON() string { return r.JSON.raw }
func (r *ExperimentStopResponseExperimentTargetingRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentStopResponseExperimentTargetingRulesQueryParam struct {
	// Query string key to inspect on the current page URL.
	Key string `json:"key" api:"required"`
	// Comparison operator applied to the query string value.
	//
	// Any of "contains", "equals", "exists", "not_equals", "not_exists".
	Operator string `json:"operator" api:"required"`
	// Comparison value used by operators that require one. Omit for `exists` and
	// `not_exists`.
	Value string `json:"value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentStopResponseExperimentTargetingRulesQueryParam) RawJSON() string { return r.JSON.raw }
func (r *ExperimentStopResponseExperimentTargetingRulesQueryParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the status change was also published to the current live version
// immediately.
type ExperimentStopResponsePublishStatus string

const (
	ExperimentStopResponsePublishStatusPendingPublish ExperimentStopResponsePublishStatus = "pending_publish"
	ExperimentStopResponsePublishStatusPublished      ExperimentStopResponsePublishStatus = "published"
)

type ExperimentPauseResponse struct {
	// Number of unpublished version changes detected at the time of the lifecycle
	// action. Values greater than 1 mean other unpublished work exists besides this
	// experiment state change.
	ConcurrentVersionChanges int64                             `json:"concurrentVersionChanges" api:"required"`
	Experiment               ExperimentPauseResponseExperiment `json:"experiment" api:"required"`
	// Whether the status change was also published to the current live version
	// immediately.
	//
	// Any of "pending_publish", "published".
	PublishStatus ExperimentPauseResponsePublishStatus `json:"publishStatus" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConcurrentVersionChanges respjson.Field
		Experiment               respjson.Field
		PublishStatus            respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentPauseResponse) RawJSON() string { return r.JSON.raw }
func (r *ExperimentPauseResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentPauseResponseExperiment struct {
	// Unique identifier for the experiment.
	ID string `json:"id" api:"required"`
	// ISO-8601 timestamp when the experiment was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Stable code-facing key for the experiment. Use this with the headless SDK
	// `getExperimentByKey()` API instead of hard-coding opaque experiment IDs into
	// application code.
	Key string `json:"key" api:"required"`
	// Short, human-readable experiment name.
	Name string `json:"name" api:"required"`
	// Lifecycle state. `draft` is editable, `running` is active, `paused` is
	// temporarily inactive, and `completed` is permanently stopped.
	//
	// Any of "completed", "draft", "paused", "running".
	Status string `json:"status" api:"required"`
	// Percent of eligible traffic assigned into the experiment. Use 0 to fully disable
	// enrollment without deleting the experiment.
	TrafficAllocation int64 `json:"trafficAllocation" api:"required"`
	// Optional human-readable hypothesis or summary. In GraphQL this is backed by the
	// experiment hypothesis field.
	Description string `json:"description" api:"nullable"`
	// For redirect variants, whether the original page query string should be
	// forwarded onto the redirect URL.
	IncludeQueryString bool `json:"includeQueryString" api:"nullable"`
	// Configured success metrics. The read shape mirrors the write shape — `metrics`
	// from a GET response can be PATCHed back without modification.
	Metrics ExperimentPauseResponseExperimentMetrics `json:"metrics" api:"nullable"`
	// ISO-8601 timestamp when the experiment most recently entered a running state.
	StartedAt string `json:"startedAt" api:"nullable"`
	// ISO-8601 timestamp when the experiment was completed, if it has been stopped.
	StoppedAt string `json:"stoppedAt" api:"nullable"`
	// Eligibility rules: URL-pattern globs, optional audience, query-param conditions,
	// visitor status, and (server-side) visitor properties. Same shape as the
	// create/patch input.
	TargetingRules ExperimentPauseResponseExperimentTargetingRules `json:"targetingRules" api:"nullable"`
	// Experiment mode. `ab` and `multivariate` use traffic allocation and results;
	// `personalization` is always-on targeting.
	//
	// Any of "ab", "multivariate", "personalization".
	Type string `json:"type" api:"nullable"`
	// ISO-8601 timestamp for the last persisted update, if any.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// Variant ID persisted as the winner when the experiment was stopped. Set via
	// `POST /experiments/{id}/stop` with a `winnerVariantId` body field.
	WinnerVariantID string `json:"winnerVariantId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Key                respjson.Field
		Name               respjson.Field
		Status             respjson.Field
		TrafficAllocation  respjson.Field
		Description        respjson.Field
		IncludeQueryString respjson.Field
		Metrics            respjson.Field
		StartedAt          respjson.Field
		StoppedAt          respjson.Field
		TargetingRules     respjson.Field
		Type               respjson.Field
		UpdatedAt          respjson.Field
		WinnerVariantID    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentPauseResponseExperiment) RawJSON() string { return r.JSON.raw }
func (r *ExperimentPauseResponseExperiment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configured success metrics. The read shape mirrors the write shape — `metrics`
// from a GET response can be PATCHed back without modification.
type ExperimentPauseResponseExperimentMetrics struct {
	// Primary success metric used in the results report.
	Primary any `json:"primary" api:"nullable"`
	// Optional secondary metrics tracked alongside the primary goal.
	Secondary []ExperimentPauseResponseExperimentMetricsSecondary `json:"secondary" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Primary     respjson.Field
		Secondary   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentPauseResponseExperimentMetrics) RawJSON() string { return r.JSON.raw }
func (r *ExperimentPauseResponseExperimentMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentPauseResponseExperimentMetricsSecondary struct {
	// Name of the event used to measure success for this metric.
	EventName string `json:"eventName" api:"nullable"`
	// Optional funnel identifier when the metric is derived from an existing funnel
	// definition.
	FunnelID string `json:"funnelId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventName   respjson.Field
		FunnelID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentPauseResponseExperimentMetricsSecondary) RawJSON() string { return r.JSON.raw }
func (r *ExperimentPauseResponseExperimentMetricsSecondary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Eligibility rules: URL-pattern globs, optional audience, query-param conditions,
// visitor status, and (server-side) visitor properties. Same shape as the
// create/patch input.
type ExperimentPauseResponseExperimentTargetingRules struct {
	// Glob-style URL patterns that must match for the experiment to be eligible. Up to
	// 200 patterns; each pattern up to 2000 characters. An empty array (or omitting
	// the field) matches all URLs — equivalent to `["**"]`.
	URLPatterns []string `json:"urlPatterns" api:"required"`
	// Optional audience identifier used for server-side eligibility filtering.
	AudienceID string `json:"audienceId" api:"nullable"`
	// Additional query-string conditions that must all match for the visitor to
	// qualify.
	QueryParams []ExperimentPauseResponseExperimentTargetingRulesQueryParam `json:"queryParams" api:"nullable"`
	// Optional visitor-property matching rules. These are passed through as JSON for
	// experimentation targeting.
	VisitorProperties any `json:"visitorProperties" api:"nullable"`
	// Whether the experiment should target new visitors, returning visitors, or any
	// visitor.
	VisitorStatus string `json:"visitorStatus" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URLPatterns       respjson.Field
		AudienceID        respjson.Field
		QueryParams       respjson.Field
		VisitorProperties respjson.Field
		VisitorStatus     respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentPauseResponseExperimentTargetingRules) RawJSON() string { return r.JSON.raw }
func (r *ExperimentPauseResponseExperimentTargetingRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentPauseResponseExperimentTargetingRulesQueryParam struct {
	// Query string key to inspect on the current page URL.
	Key string `json:"key" api:"required"`
	// Comparison operator applied to the query string value.
	//
	// Any of "contains", "equals", "exists", "not_equals", "not_exists".
	Operator string `json:"operator" api:"required"`
	// Comparison value used by operators that require one. Omit for `exists` and
	// `not_exists`.
	Value string `json:"value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentPauseResponseExperimentTargetingRulesQueryParam) RawJSON() string {
	return r.JSON.raw
}
func (r *ExperimentPauseResponseExperimentTargetingRulesQueryParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the status change was also published to the current live version
// immediately.
type ExperimentPauseResponsePublishStatus string

const (
	ExperimentPauseResponsePublishStatusPendingPublish ExperimentPauseResponsePublishStatus = "pending_publish"
	ExperimentPauseResponsePublishStatusPublished      ExperimentPauseResponsePublishStatus = "published"
)

type ExperimentResumeResponse struct {
	// Number of unpublished version changes detected at the time of the lifecycle
	// action. Values greater than 1 mean other unpublished work exists besides this
	// experiment state change.
	ConcurrentVersionChanges int64                              `json:"concurrentVersionChanges" api:"required"`
	Experiment               ExperimentResumeResponseExperiment `json:"experiment" api:"required"`
	// Whether the status change was also published to the current live version
	// immediately.
	//
	// Any of "pending_publish", "published".
	PublishStatus ExperimentResumeResponsePublishStatus `json:"publishStatus" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConcurrentVersionChanges respjson.Field
		Experiment               respjson.Field
		PublishStatus            respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentResumeResponse) RawJSON() string { return r.JSON.raw }
func (r *ExperimentResumeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentResumeResponseExperiment struct {
	// Unique identifier for the experiment.
	ID string `json:"id" api:"required"`
	// ISO-8601 timestamp when the experiment was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Stable code-facing key for the experiment. Use this with the headless SDK
	// `getExperimentByKey()` API instead of hard-coding opaque experiment IDs into
	// application code.
	Key string `json:"key" api:"required"`
	// Short, human-readable experiment name.
	Name string `json:"name" api:"required"`
	// Lifecycle state. `draft` is editable, `running` is active, `paused` is
	// temporarily inactive, and `completed` is permanently stopped.
	//
	// Any of "completed", "draft", "paused", "running".
	Status string `json:"status" api:"required"`
	// Percent of eligible traffic assigned into the experiment. Use 0 to fully disable
	// enrollment without deleting the experiment.
	TrafficAllocation int64 `json:"trafficAllocation" api:"required"`
	// Optional human-readable hypothesis or summary. In GraphQL this is backed by the
	// experiment hypothesis field.
	Description string `json:"description" api:"nullable"`
	// For redirect variants, whether the original page query string should be
	// forwarded onto the redirect URL.
	IncludeQueryString bool `json:"includeQueryString" api:"nullable"`
	// Configured success metrics. The read shape mirrors the write shape — `metrics`
	// from a GET response can be PATCHed back without modification.
	Metrics ExperimentResumeResponseExperimentMetrics `json:"metrics" api:"nullable"`
	// ISO-8601 timestamp when the experiment most recently entered a running state.
	StartedAt string `json:"startedAt" api:"nullable"`
	// ISO-8601 timestamp when the experiment was completed, if it has been stopped.
	StoppedAt string `json:"stoppedAt" api:"nullable"`
	// Eligibility rules: URL-pattern globs, optional audience, query-param conditions,
	// visitor status, and (server-side) visitor properties. Same shape as the
	// create/patch input.
	TargetingRules ExperimentResumeResponseExperimentTargetingRules `json:"targetingRules" api:"nullable"`
	// Experiment mode. `ab` and `multivariate` use traffic allocation and results;
	// `personalization` is always-on targeting.
	//
	// Any of "ab", "multivariate", "personalization".
	Type string `json:"type" api:"nullable"`
	// ISO-8601 timestamp for the last persisted update, if any.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// Variant ID persisted as the winner when the experiment was stopped. Set via
	// `POST /experiments/{id}/stop` with a `winnerVariantId` body field.
	WinnerVariantID string `json:"winnerVariantId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Key                respjson.Field
		Name               respjson.Field
		Status             respjson.Field
		TrafficAllocation  respjson.Field
		Description        respjson.Field
		IncludeQueryString respjson.Field
		Metrics            respjson.Field
		StartedAt          respjson.Field
		StoppedAt          respjson.Field
		TargetingRules     respjson.Field
		Type               respjson.Field
		UpdatedAt          respjson.Field
		WinnerVariantID    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentResumeResponseExperiment) RawJSON() string { return r.JSON.raw }
func (r *ExperimentResumeResponseExperiment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configured success metrics. The read shape mirrors the write shape — `metrics`
// from a GET response can be PATCHed back without modification.
type ExperimentResumeResponseExperimentMetrics struct {
	// Primary success metric used in the results report.
	Primary any `json:"primary" api:"nullable"`
	// Optional secondary metrics tracked alongside the primary goal.
	Secondary []ExperimentResumeResponseExperimentMetricsSecondary `json:"secondary" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Primary     respjson.Field
		Secondary   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentResumeResponseExperimentMetrics) RawJSON() string { return r.JSON.raw }
func (r *ExperimentResumeResponseExperimentMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentResumeResponseExperimentMetricsSecondary struct {
	// Name of the event used to measure success for this metric.
	EventName string `json:"eventName" api:"nullable"`
	// Optional funnel identifier when the metric is derived from an existing funnel
	// definition.
	FunnelID string `json:"funnelId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventName   respjson.Field
		FunnelID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentResumeResponseExperimentMetricsSecondary) RawJSON() string { return r.JSON.raw }
func (r *ExperimentResumeResponseExperimentMetricsSecondary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Eligibility rules: URL-pattern globs, optional audience, query-param conditions,
// visitor status, and (server-side) visitor properties. Same shape as the
// create/patch input.
type ExperimentResumeResponseExperimentTargetingRules struct {
	// Glob-style URL patterns that must match for the experiment to be eligible. Up to
	// 200 patterns; each pattern up to 2000 characters. An empty array (or omitting
	// the field) matches all URLs — equivalent to `["**"]`.
	URLPatterns []string `json:"urlPatterns" api:"required"`
	// Optional audience identifier used for server-side eligibility filtering.
	AudienceID string `json:"audienceId" api:"nullable"`
	// Additional query-string conditions that must all match for the visitor to
	// qualify.
	QueryParams []ExperimentResumeResponseExperimentTargetingRulesQueryParam `json:"queryParams" api:"nullable"`
	// Optional visitor-property matching rules. These are passed through as JSON for
	// experimentation targeting.
	VisitorProperties any `json:"visitorProperties" api:"nullable"`
	// Whether the experiment should target new visitors, returning visitors, or any
	// visitor.
	VisitorStatus string `json:"visitorStatus" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URLPatterns       respjson.Field
		AudienceID        respjson.Field
		QueryParams       respjson.Field
		VisitorProperties respjson.Field
		VisitorStatus     respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentResumeResponseExperimentTargetingRules) RawJSON() string { return r.JSON.raw }
func (r *ExperimentResumeResponseExperimentTargetingRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentResumeResponseExperimentTargetingRulesQueryParam struct {
	// Query string key to inspect on the current page URL.
	Key string `json:"key" api:"required"`
	// Comparison operator applied to the query string value.
	//
	// Any of "contains", "equals", "exists", "not_equals", "not_exists".
	Operator string `json:"operator" api:"required"`
	// Comparison value used by operators that require one. Omit for `exists` and
	// `not_exists`.
	Value string `json:"value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentResumeResponseExperimentTargetingRulesQueryParam) RawJSON() string {
	return r.JSON.raw
}
func (r *ExperimentResumeResponseExperimentTargetingRulesQueryParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the status change was also published to the current live version
// immediately.
type ExperimentResumeResponsePublishStatus string

const (
	ExperimentResumeResponsePublishStatusPendingPublish ExperimentResumeResponsePublishStatus = "pending_publish"
	ExperimentResumeResponsePublishStatusPublished      ExperimentResumeResponsePublishStatus = "published"
)

type ExperimentResultsResponse struct {
	// Aggregate performance metrics for each variant in the experiment over the
	// experiment runtime window.
	Variants []ExperimentResultsResponseVariant `json:"variants" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Variants    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentResultsResponse) RawJSON() string { return r.JSON.raw }
func (r *ExperimentResultsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentResultsResponseVariant struct {
	// Variant ID this result row belongs to.
	ID string `json:"id" api:"required"`
	// Conversions divided by impressions for this variant.
	ConversionRate float64 `json:"conversionRate" api:"required"`
	// Number of post-exposure conversions attributed to the variant.
	Conversions int64 `json:"conversions" api:"required"`
	// Number of distinct experiment-impression events counted for the variant.
	Impressions int64 `json:"impressions" api:"required"`
	// Whether this variant is the experiment control. Exactly one variant per
	// experiment is the control.
	IsControl bool `json:"isControl" api:"required"`
	// Human-readable variant name as configured on the experiment (e.g. `Control`,
	// `Treatment`). Lets callers label results without a separate lookup against
	// `/experiment-variants`.
	Name string `json:"name" api:"required"`
	// Bayesian probability that this variant is the best-performing option among all
	// variants in the experiment.
	ProbabilityToBeBest float64 `json:"probabilityToBeBest" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		ConversionRate      respjson.Field
		Conversions         respjson.Field
		Impressions         respjson.Field
		IsControl           respjson.Field
		Name                respjson.Field
		ProbabilityToBeBest respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentResultsResponseVariant) RawJSON() string { return r.JSON.raw }
func (r *ExperimentResultsResponseVariant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentResultsTimeSeriesResponse struct {
	// Per-day metrics for each variant. Days with no impression rows are omitted from
	// the response.
	Days []ExperimentResultsTimeSeriesResponseDay `json:"days" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Days        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentResultsTimeSeriesResponse) RawJSON() string { return r.JSON.raw }
func (r *ExperimentResultsTimeSeriesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentResultsTimeSeriesResponseDay struct {
	// UTC calendar day in `YYYY-MM-DD` format.
	Date     string                                          `json:"date" api:"required"`
	Variants []ExperimentResultsTimeSeriesResponseDayVariant `json:"variants" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		Variants    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentResultsTimeSeriesResponseDay) RawJSON() string { return r.JSON.raw }
func (r *ExperimentResultsTimeSeriesResponseDay) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentResultsTimeSeriesResponseDayVariant struct {
	ID             string  `json:"id" api:"required"`
	ConversionRate float64 `json:"conversionRate" api:"required"`
	Conversions    int64   `json:"conversions" api:"required"`
	Impressions    int64   `json:"impressions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		ConversionRate respjson.Field
		Conversions    respjson.Field
		Impressions    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentResultsTimeSeriesResponseDayVariant) RawJSON() string { return r.JSON.raw }
func (r *ExperimentResultsTimeSeriesResponseDayVariant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentListParams struct {
	// Maximum number of items to return. Defaults to 25; values below 1 are clamped to
	// 1 and values above 100 are clamped to 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Opaque pagination cursor from pagination.nextCursor in the previous response. Do
	// not decode or modify it. Malformed cursors return 400 Bad Request.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Optional case-insensitive text search matched against experiment ID, name, and
	// description.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Optional lifecycle-state filter.
	//
	// Any of "completed", "draft", "paused", "running".
	Status ExperimentListParamsStatus `query:"status,omitzero" json:"-"`
	// Optional experiment-type filter.
	//
	// Any of "ab", "multivariate", "personalization".
	Type ExperimentListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExperimentListParams]'s query parameters as `url.Values`.
func (r ExperimentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Optional lifecycle-state filter.
type ExperimentListParamsStatus string

const (
	ExperimentListParamsStatusCompleted ExperimentListParamsStatus = "completed"
	ExperimentListParamsStatusDraft     ExperimentListParamsStatus = "draft"
	ExperimentListParamsStatusPaused    ExperimentListParamsStatus = "paused"
	ExperimentListParamsStatusRunning   ExperimentListParamsStatus = "running"
)

// Optional experiment-type filter.
type ExperimentListParamsType string

const (
	ExperimentListParamsTypeAb              ExperimentListParamsType = "ab"
	ExperimentListParamsTypeMultivariate    ExperimentListParamsType = "multivariate"
	ExperimentListParamsTypePersonalization ExperimentListParamsType = "personalization"
)

type ExperimentNewParams struct {
	// ID of the experiment settings record that owns this experiment. Call
	// `GET /rest/v1/experiment-settings` first if you need to discover the correct ID
	// for the current account.
	ExperimentSettingsID string `json:"experimentSettingsId" api:"required"`
	// Short experiment name.
	Name string `json:"name" api:"required"`
	// Optional hypothesis or operator note.
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether redirect variants in this experiment should preserve the original
	// request query string.
	IncludeQueryString param.Opt[bool] `json:"includeQueryString,omitzero"`
	// Optional stable code-facing key. When omitted, the API slugifies the name
	// automatically.
	Key param.Opt[string] `json:"key,omitzero"`
	// Initial traffic allocation percentage from 0 to 100.
	TrafficAllocation param.Opt[float64] `json:"trafficAllocation,omitzero"`
	// Goal events. If you send `metrics.primary`, `metrics.primary.eventName` must be
	// a non-blank string. A primary event name is required before the experiment can
	// be started.
	Metrics ExperimentNewParamsMetrics `json:"metrics,omitzero"`
	// Eligibility rules — URL patterns, audience, visitor status, query-param
	// conditions. Omit to inherit defaults.
	TargetingRules ExperimentNewParamsTargetingRules `json:"targetingRules,omitzero"`
	// Experiment mode to create. `ab` and `multivariate` use traffic allocation and
	// results; `personalization` is always-on targeting. Omit to create a standard
	// `ab` experiment.
	//
	// Any of "ab", "multivariate", "personalization".
	Type ExperimentNewParamsType `json:"type,omitzero"`
	paramObj
}

func (r ExperimentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Goal events. If you send `metrics.primary`, `metrics.primary.eventName` must be
// a non-blank string. A primary event name is required before the experiment can
// be started.
type ExperimentNewParamsMetrics struct {
	// Primary success metric. When provided, `eventName` must be a non-blank string.
	Primary ExperimentNewParamsMetricsPrimary `json:"primary,omitzero"`
	// Optional secondary metrics tracked alongside the primary goal.
	Secondary []ExperimentNewParamsMetricsSecondary `json:"secondary,omitzero"`
	paramObj
}

func (r ExperimentNewParamsMetrics) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentNewParamsMetrics
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentNewParamsMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Primary success metric. When provided, `eventName` must be a non-blank string.
type ExperimentNewParamsMetricsPrimary struct {
	// Event name to use as the goal for this metric.
	EventName param.Opt[string] `json:"eventName,omitzero"`
	// Optional funnel identifier when the metric should be derived from an existing
	// funnel definition.
	FunnelID param.Opt[string] `json:"funnelId,omitzero"`
	paramObj
}

func (r ExperimentNewParamsMetricsPrimary) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentNewParamsMetricsPrimary
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentNewParamsMetricsPrimary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentNewParamsMetricsSecondary struct {
	// Event name to use as the goal for this metric.
	EventName param.Opt[string] `json:"eventName,omitzero"`
	// Optional funnel identifier when the metric should be derived from an existing
	// funnel definition.
	FunnelID param.Opt[string] `json:"funnelId,omitzero"`
	paramObj
}

func (r ExperimentNewParamsMetricsSecondary) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentNewParamsMetricsSecondary
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentNewParamsMetricsSecondary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Eligibility rules — URL patterns, audience, visitor status, query-param
// conditions. Omit to inherit defaults.
//
// The property URLPatterns is required.
type ExperimentNewParamsTargetingRules struct {
	// Glob-style URL patterns that must match for the experiment to be eligible. Up to
	// 200 patterns; each pattern up to 2000 characters. An empty array (or omitting
	// the field) matches all URLs — equivalent to `["**"]`.
	URLPatterns []string `json:"urlPatterns,omitzero" api:"required"`
	// Optional audience identifier used for server-side eligibility filtering.
	AudienceID param.Opt[string] `json:"audienceId,omitzero"`
	// Whether the experiment should target new visitors, returning visitors, or any
	// visitor.
	VisitorStatus param.Opt[string] `json:"visitorStatus,omitzero"`
	// Additional query-string conditions that must all match for the visitor to
	// qualify.
	QueryParams []ExperimentNewParamsTargetingRulesQueryParam `json:"queryParams,omitzero"`
	// Optional visitor-property matching rules. These are passed through as JSON for
	// experimentation targeting.
	VisitorProperties any `json:"visitorProperties,omitzero"`
	paramObj
}

func (r ExperimentNewParamsTargetingRules) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentNewParamsTargetingRules
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentNewParamsTargetingRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Key, Operator are required.
type ExperimentNewParamsTargetingRulesQueryParam struct {
	// Query string key to inspect on the current page URL.
	Key string `json:"key" api:"required"`
	// Comparison operator applied to the query string value.
	//
	// Any of "contains", "equals", "exists", "not_equals", "not_exists".
	Operator string `json:"operator,omitzero" api:"required"`
	// Comparison value used by operators that require one. Omit for `exists` and
	// `not_exists`.
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r ExperimentNewParamsTargetingRulesQueryParam) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentNewParamsTargetingRulesQueryParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentNewParamsTargetingRulesQueryParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExperimentNewParamsTargetingRulesQueryParam](
		"operator", "contains", "equals", "exists", "not_equals", "not_exists",
	)
}

// Experiment mode to create. `ab` and `multivariate` use traffic allocation and
// results; `personalization` is always-on targeting. Omit to create a standard
// `ab` experiment.
type ExperimentNewParamsType string

const (
	ExperimentNewParamsTypeAb              ExperimentNewParamsType = "ab"
	ExperimentNewParamsTypeMultivariate    ExperimentNewParamsType = "multivariate"
	ExperimentNewParamsTypePersonalization ExperimentNewParamsType = "personalization"
)

type ExperimentUpdateParams struct {
	// Updated experiment hypothesis or operator note.
	Description param.Opt[string] `json:"description,omitzero"`
	// Updated redirect query-string forwarding behavior for the experiment.
	IncludeQueryString param.Opt[bool] `json:"includeQueryString,omitzero"`
	// Updated stable code-facing key. When blank, the API falls back to a slugified
	// key derived from the current experiment name.
	Key param.Opt[string] `json:"key,omitzero"`
	// Updated experiment name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Updated traffic allocation percentage from 0 to 100.
	TrafficAllocation param.Opt[float64] `json:"trafficAllocation,omitzero"`
	// Updated goal events. Send the full nested object — replaces the previous value,
	// not merged. If you send `metrics.primary`, `metrics.primary.eventName` must be a
	// non-blank string.
	Metrics ExperimentUpdateParamsMetrics `json:"metrics,omitzero"`
	// Updated eligibility rules. Send the full nested object — replaces the previous
	// value, not merged.
	TargetingRules ExperimentUpdateParamsTargetingRules `json:"targetingRules,omitzero"`
	paramObj
}

func (r ExperimentUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Updated goal events. Send the full nested object — replaces the previous value,
// not merged. If you send `metrics.primary`, `metrics.primary.eventName` must be a
// non-blank string.
type ExperimentUpdateParamsMetrics struct {
	// Primary success metric. When provided, `eventName` must be a non-blank string.
	Primary ExperimentUpdateParamsMetricsPrimary `json:"primary,omitzero"`
	// Optional secondary metrics tracked alongside the primary goal.
	Secondary []ExperimentUpdateParamsMetricsSecondary `json:"secondary,omitzero"`
	paramObj
}

func (r ExperimentUpdateParamsMetrics) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentUpdateParamsMetrics
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentUpdateParamsMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Primary success metric. When provided, `eventName` must be a non-blank string.
type ExperimentUpdateParamsMetricsPrimary struct {
	// Event name to use as the goal for this metric.
	EventName param.Opt[string] `json:"eventName,omitzero"`
	// Optional funnel identifier when the metric should be derived from an existing
	// funnel definition.
	FunnelID param.Opt[string] `json:"funnelId,omitzero"`
	paramObj
}

func (r ExperimentUpdateParamsMetricsPrimary) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentUpdateParamsMetricsPrimary
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentUpdateParamsMetricsPrimary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentUpdateParamsMetricsSecondary struct {
	// Event name to use as the goal for this metric.
	EventName param.Opt[string] `json:"eventName,omitzero"`
	// Optional funnel identifier when the metric should be derived from an existing
	// funnel definition.
	FunnelID param.Opt[string] `json:"funnelId,omitzero"`
	paramObj
}

func (r ExperimentUpdateParamsMetricsSecondary) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentUpdateParamsMetricsSecondary
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentUpdateParamsMetricsSecondary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Updated eligibility rules. Send the full nested object — replaces the previous
// value, not merged.
//
// The property URLPatterns is required.
type ExperimentUpdateParamsTargetingRules struct {
	// Glob-style URL patterns that must match for the experiment to be eligible. Up to
	// 200 patterns; each pattern up to 2000 characters. An empty array (or omitting
	// the field) matches all URLs — equivalent to `["**"]`.
	URLPatterns []string `json:"urlPatterns,omitzero" api:"required"`
	// Optional audience identifier used for server-side eligibility filtering.
	AudienceID param.Opt[string] `json:"audienceId,omitzero"`
	// Whether the experiment should target new visitors, returning visitors, or any
	// visitor.
	VisitorStatus param.Opt[string] `json:"visitorStatus,omitzero"`
	// Additional query-string conditions that must all match for the visitor to
	// qualify.
	QueryParams []ExperimentUpdateParamsTargetingRulesQueryParam `json:"queryParams,omitzero"`
	// Optional visitor-property matching rules. These are passed through as JSON for
	// experimentation targeting.
	VisitorProperties any `json:"visitorProperties,omitzero"`
	paramObj
}

func (r ExperimentUpdateParamsTargetingRules) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentUpdateParamsTargetingRules
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentUpdateParamsTargetingRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Key, Operator are required.
type ExperimentUpdateParamsTargetingRulesQueryParam struct {
	// Query string key to inspect on the current page URL.
	Key string `json:"key" api:"required"`
	// Comparison operator applied to the query string value.
	//
	// Any of "contains", "equals", "exists", "not_equals", "not_exists".
	Operator string `json:"operator,omitzero" api:"required"`
	// Comparison value used by operators that require one. Omit for `exists` and
	// `not_exists`.
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r ExperimentUpdateParamsTargetingRulesQueryParam) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentUpdateParamsTargetingRulesQueryParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentUpdateParamsTargetingRulesQueryParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExperimentUpdateParamsTargetingRulesQueryParam](
		"operator", "contains", "equals", "exists", "not_equals", "not_exists",
	)
}

type ExperimentStartParams struct {
	// When true (default on the REST surface), publish the current draft version
	// immediately after starting the experiment. Any other unpublished changes in the
	// same account version are included in that publish. Pass `false` explicitly to
	// stage the change without publishing; the response will report `pending_publish`.
	PublishAfterStart param.Opt[bool] `json:"publishAfterStart,omitzero"`
	paramObj
}

func (r ExperimentStartParams) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentStartParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentStartParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentStopParams struct {
	// Optional winning variant ID to persist when completing the experiment.
	WinnerVariantID param.Opt[string] `json:"winnerVariantId,omitzero"`
	paramObj
}

func (r ExperimentStopParams) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentStopParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentStopParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentPauseParams struct {
	// When true (default on the REST surface), publish the current draft version
	// immediately after pausing the experiment. Any other unpublished changes in the
	// same account version are included in that publish. Pass `false` explicitly to
	// stage the change without publishing; the response will report `pending_publish`.
	PublishAfterPause param.Opt[bool] `json:"publishAfterPause,omitzero"`
	paramObj
}

func (r ExperimentPauseParams) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentPauseParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentPauseParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentResumeParams struct {
	// When true (default on the REST surface), publish the current draft version
	// immediately after resuming the experiment. Any other unpublished changes in the
	// same account version are included in that publish. Pass `false` explicitly to
	// stage the change without publishing; the response will report `pending_publish`.
	PublishAfterResume param.Opt[bool] `json:"publishAfterResume,omitzero"`
	paramObj
}

func (r ExperimentResumeParams) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentResumeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentResumeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentResultsParams struct {
	// Optional override for the conversion event name. When omitted, the experiment
	// primary metric event is used.
	EventName param.Opt[string] `query:"eventName,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExperimentResultsParams]'s query parameters as
// `url.Values`.
func (r ExperimentResultsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExperimentResultsTimeSeriesParams struct {
	// Inclusive upper bound of the response window, as a UTC calendar day in
	// `YYYY-MM-DD` format. Defaults to the experiment stop date for completed
	// experiments, or today for running experiments. Values after that are silently
	// clamped. The window between `startDate` and `endDate` must be 366 days or fewer.
	EndDate param.Opt[string] `query:"endDate,omitzero" json:"-"`
	// Optional override for the conversion event name. When omitted, the experiment
	// primary metric event is used.
	EventName param.Opt[string] `query:"eventName,omitzero" json:"-"`
	// Inclusive lower bound of the response window, as a UTC calendar day in
	// `YYYY-MM-DD` format. Defaults to the experiment start date when omitted. Values
	// before the experiment started are silently clamped to the experiment start.
	StartDate param.Opt[string] `query:"startDate,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExperimentResultsTimeSeriesParams]'s query parameters as
// `url.Values`.
func (r ExperimentResultsTimeSeriesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
