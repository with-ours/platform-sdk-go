// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy

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

// List every funnel configured on this account. Each funnel includes its step
// configuration, funnel type, and conversion window. Funnel results are computed
// on demand, so `status` is always `READY` and `reportDateRange` is always `null`;
// both fields are retained for backward compatibility and should not be used to
// decide whether results are available. Requires scope: web-analytics:view
func (r *FunnelService) List(ctx context.Context, opts ...option.RequestOption) (res *FunnelListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/funnels"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Create a session-based funnel with 2 to 10 ordered event steps. Returns the
// complete saved configuration. Requires scope: web-analytics:write
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

// Update one or more Funnel fields. Omitted fields remain unchanged. Send `null`
// to clear an optional field. `globalLogic` and legacy `utmFilters` cannot be set
// together. Requires scope: web-analytics:write
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

// Duplicate a funnel configuration in the same account. The copy keeps the funnel
// steps and settings, receives a new ID, and is named `Copy of …`. Requires scope:
// web-analytics:write
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

// Compute funnel step analytics for a funnel over a date window. Returns per-step
// visitor counts, conversion rates, drop-off rates, average time to next step, and
// sample session IDs for replay. Results are computed on demand from event data at
// request time, so any date window within the supported range returns current
// results. `to` must be on or after `from`, and the window may span at most 31
// days including both endpoints. Requires scope: web-analytics:view
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
	// All funnels configured on this account.
	Entities []FunnelListResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	// Any of "SESSION_BASED", "VISITOR_BASED".
	FunnelType string `json:"funnelType" api:"required"`
	Name       string `json:"name" api:"required"`
	// Any of "READY", "PROCESSING".
	Status           string                                   `json:"status" api:"required"`
	Steps            []FunnelListResponseEntityStep           `json:"steps" api:"required"`
	UpdatedAt        string                                   `json:"updatedAt" api:"required"`
	ConversionWindow FunnelListResponseEntityConversionWindow `json:"conversionWindow" api:"nullable"`
	// Any of "UNIQUES", "TOTALS", "SESSIONS".
	CountingMethod string `json:"countingMethod" api:"nullable"`
	Description    string `json:"description" api:"nullable"`
	// Nested visitor logic for the entire funnel. When supplied, this replaces legacy
	// UTM filters.
	GlobalLogic     FunnelListResponseEntityGlobalLogic     `json:"globalLogic" api:"nullable"`
	ReportDateRange FunnelListResponseEntityReportDateRange `json:"reportDateRange" api:"nullable"`
	// Any of "EXACT", "ANY".
	StepOrder string `json:"stepOrder" api:"nullable"`
	// Legacy exact-match UTM filters. Do not combine with globalLogic; globalLogic
	// takes precedence.
	UtmFilters any  `json:"utmFilters" api:"nullable"`
	Watched    bool `json:"watched" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt        respjson.Field
		FunnelID         respjson.Field
		FunnelType       respjson.Field
		Name             respjson.Field
		Status           respjson.Field
		Steps            respjson.Field
		UpdatedAt        respjson.Field
		ConversionWindow respjson.Field
		CountingMethod   respjson.Field
		Description      respjson.Field
		GlobalLogic      respjson.Field
		ReportDateRange  respjson.Field
		StepOrder        respjson.Field
		UtmFilters       respjson.Field
		Watched          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelListResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *FunnelListResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelListResponseEntityStep struct {
	EventName string `json:"eventName" api:"required"`
	Name      string `json:"name" api:"required"`
	Order     int64  `json:"order" api:"required"`
	StepID    string `json:"stepId" api:"required"`
	// Step-level event filters (JSON object).
	Filters any `json:"filters"`
	// Step-level event logic.
	Logic FunnelListResponseEntityStepLogic `json:"logic" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventName   respjson.Field
		Name        respjson.Field
		Order       respjson.Field
		StepID      respjson.Field
		Filters     respjson.Field
		Logic       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelListResponseEntityStep) RawJSON() string { return r.JSON.raw }
func (r *FunnelListResponseEntityStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Step-level event logic.
type FunnelListResponseEntityStepLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And       []any                                      `json:"AND" api:"nullable"`
	Condition FunnelListResponseEntityStepLogicCondition `json:"condition" api:"nullable"`
	// Negates a single child logic node.
	Not any `json:"NOT"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or []any `json:"OR" api:"nullable"`
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
func (r FunnelListResponseEntityStepLogic) RawJSON() string { return r.JSON.raw }
func (r *FunnelListResponseEntityStepLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelListResponseEntityStepLogicCondition struct {
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
func (r FunnelListResponseEntityStepLogicCondition) RawJSON() string { return r.JSON.raw }
func (r *FunnelListResponseEntityStepLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelListResponseEntityConversionWindow struct {
	// Any of "MINUTES", "HOURS", "DAYS".
	Unit  string `json:"unit" api:"required"`
	Value int64  `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Unit        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelListResponseEntityConversionWindow) RawJSON() string { return r.JSON.raw }
func (r *FunnelListResponseEntityConversionWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Nested visitor logic for the entire funnel. When supplied, this replaces legacy
// UTM filters.
type FunnelListResponseEntityGlobalLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And       []any                                        `json:"AND" api:"nullable"`
	Condition FunnelListResponseEntityGlobalLogicCondition `json:"condition" api:"nullable"`
	// Negates a single child logic node.
	Not any `json:"NOT"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or []any `json:"OR" api:"nullable"`
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
func (r FunnelListResponseEntityGlobalLogic) RawJSON() string { return r.JSON.raw }
func (r *FunnelListResponseEntityGlobalLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelListResponseEntityGlobalLogicCondition struct {
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
func (r FunnelListResponseEntityGlobalLogicCondition) RawJSON() string { return r.JSON.raw }
func (r *FunnelListResponseEntityGlobalLogicCondition) UnmarshalJSON(data []byte) error {
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
	// Any of "SESSION_BASED", "VISITOR_BASED".
	FunnelType FunnelNewResponseFunnelType `json:"funnelType" api:"required"`
	Name       string                      `json:"name" api:"required"`
	// Any of "READY", "PROCESSING".
	Status           FunnelNewResponseStatus           `json:"status" api:"required"`
	Steps            []FunnelNewResponseStep           `json:"steps" api:"required"`
	UpdatedAt        string                            `json:"updatedAt" api:"required"`
	ConversionWindow FunnelNewResponseConversionWindow `json:"conversionWindow" api:"nullable"`
	// Any of "UNIQUES", "TOTALS", "SESSIONS".
	CountingMethod FunnelNewResponseCountingMethod `json:"countingMethod" api:"nullable"`
	Description    string                          `json:"description" api:"nullable"`
	// Nested visitor logic for the entire funnel. When supplied, this replaces legacy
	// UTM filters.
	GlobalLogic     FunnelNewResponseGlobalLogic     `json:"globalLogic" api:"nullable"`
	ReportDateRange FunnelNewResponseReportDateRange `json:"reportDateRange" api:"nullable"`
	// Any of "EXACT", "ANY".
	StepOrder FunnelNewResponseStepOrder `json:"stepOrder" api:"nullable"`
	// Legacy exact-match UTM filters. Do not combine with globalLogic; globalLogic
	// takes precedence.
	UtmFilters any  `json:"utmFilters" api:"nullable"`
	Watched    bool `json:"watched" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt        respjson.Field
		FunnelID         respjson.Field
		FunnelType       respjson.Field
		Name             respjson.Field
		Status           respjson.Field
		Steps            respjson.Field
		UpdatedAt        respjson.Field
		ConversionWindow respjson.Field
		CountingMethod   respjson.Field
		Description      respjson.Field
		GlobalLogic      respjson.Field
		ReportDateRange  respjson.Field
		StepOrder        respjson.Field
		UtmFilters       respjson.Field
		Watched          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelNewResponse) RawJSON() string { return r.JSON.raw }
func (r *FunnelNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelNewResponseFunnelType string

const (
	FunnelNewResponseFunnelTypeSessionBased FunnelNewResponseFunnelType = "SESSION_BASED"
	FunnelNewResponseFunnelTypeVisitorBased FunnelNewResponseFunnelType = "VISITOR_BASED"
)

type FunnelNewResponseStatus string

const (
	FunnelNewResponseStatusReady      FunnelNewResponseStatus = "READY"
	FunnelNewResponseStatusProcessing FunnelNewResponseStatus = "PROCESSING"
)

type FunnelNewResponseStep struct {
	EventName string `json:"eventName" api:"required"`
	Name      string `json:"name" api:"required"`
	Order     int64  `json:"order" api:"required"`
	StepID    string `json:"stepId" api:"required"`
	// Step-level event filters (JSON object).
	Filters any `json:"filters"`
	// Step-level event logic.
	Logic FunnelNewResponseStepLogic `json:"logic" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventName   respjson.Field
		Name        respjson.Field
		Order       respjson.Field
		StepID      respjson.Field
		Filters     respjson.Field
		Logic       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelNewResponseStep) RawJSON() string { return r.JSON.raw }
func (r *FunnelNewResponseStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Step-level event logic.
type FunnelNewResponseStepLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And       []any                               `json:"AND" api:"nullable"`
	Condition FunnelNewResponseStepLogicCondition `json:"condition" api:"nullable"`
	// Negates a single child logic node.
	Not any `json:"NOT"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or []any `json:"OR" api:"nullable"`
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
func (r FunnelNewResponseStepLogic) RawJSON() string { return r.JSON.raw }
func (r *FunnelNewResponseStepLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelNewResponseStepLogicCondition struct {
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
func (r FunnelNewResponseStepLogicCondition) RawJSON() string { return r.JSON.raw }
func (r *FunnelNewResponseStepLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelNewResponseConversionWindow struct {
	// Any of "MINUTES", "HOURS", "DAYS".
	Unit  string `json:"unit" api:"required"`
	Value int64  `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Unit        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelNewResponseConversionWindow) RawJSON() string { return r.JSON.raw }
func (r *FunnelNewResponseConversionWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelNewResponseCountingMethod string

const (
	FunnelNewResponseCountingMethodUniques  FunnelNewResponseCountingMethod = "UNIQUES"
	FunnelNewResponseCountingMethodTotals   FunnelNewResponseCountingMethod = "TOTALS"
	FunnelNewResponseCountingMethodSessions FunnelNewResponseCountingMethod = "SESSIONS"
)

// Nested visitor logic for the entire funnel. When supplied, this replaces legacy
// UTM filters.
type FunnelNewResponseGlobalLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And       []any                                 `json:"AND" api:"nullable"`
	Condition FunnelNewResponseGlobalLogicCondition `json:"condition" api:"nullable"`
	// Negates a single child logic node.
	Not any `json:"NOT"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or []any `json:"OR" api:"nullable"`
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
func (r FunnelNewResponseGlobalLogic) RawJSON() string { return r.JSON.raw }
func (r *FunnelNewResponseGlobalLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelNewResponseGlobalLogicCondition struct {
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
func (r FunnelNewResponseGlobalLogicCondition) RawJSON() string { return r.JSON.raw }
func (r *FunnelNewResponseGlobalLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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

type FunnelNewResponseStepOrder string

const (
	FunnelNewResponseStepOrderExact FunnelNewResponseStepOrder = "EXACT"
	FunnelNewResponseStepOrderAny   FunnelNewResponseStepOrder = "ANY"
)

// Funnel configuration details.
type FunnelGetResponse struct {
	CreatedAt string `json:"createdAt" api:"required"`
	FunnelID  string `json:"funnelId" api:"required" format:"uuid"`
	// Any of "SESSION_BASED", "VISITOR_BASED".
	FunnelType FunnelGetResponseFunnelType `json:"funnelType" api:"required"`
	Name       string                      `json:"name" api:"required"`
	// Any of "READY", "PROCESSING".
	Status           FunnelGetResponseStatus           `json:"status" api:"required"`
	Steps            []FunnelGetResponseStep           `json:"steps" api:"required"`
	UpdatedAt        string                            `json:"updatedAt" api:"required"`
	ConversionWindow FunnelGetResponseConversionWindow `json:"conversionWindow" api:"nullable"`
	// Any of "UNIQUES", "TOTALS", "SESSIONS".
	CountingMethod FunnelGetResponseCountingMethod `json:"countingMethod" api:"nullable"`
	Description    string                          `json:"description" api:"nullable"`
	// Nested visitor logic for the entire funnel. When supplied, this replaces legacy
	// UTM filters.
	GlobalLogic     FunnelGetResponseGlobalLogic     `json:"globalLogic" api:"nullable"`
	ReportDateRange FunnelGetResponseReportDateRange `json:"reportDateRange" api:"nullable"`
	// Any of "EXACT", "ANY".
	StepOrder FunnelGetResponseStepOrder `json:"stepOrder" api:"nullable"`
	// Legacy exact-match UTM filters. Do not combine with globalLogic; globalLogic
	// takes precedence.
	UtmFilters any  `json:"utmFilters" api:"nullable"`
	Watched    bool `json:"watched" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt        respjson.Field
		FunnelID         respjson.Field
		FunnelType       respjson.Field
		Name             respjson.Field
		Status           respjson.Field
		Steps            respjson.Field
		UpdatedAt        respjson.Field
		ConversionWindow respjson.Field
		CountingMethod   respjson.Field
		Description      respjson.Field
		GlobalLogic      respjson.Field
		ReportDateRange  respjson.Field
		StepOrder        respjson.Field
		UtmFilters       respjson.Field
		Watched          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelGetResponse) RawJSON() string { return r.JSON.raw }
func (r *FunnelGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelGetResponseFunnelType string

const (
	FunnelGetResponseFunnelTypeSessionBased FunnelGetResponseFunnelType = "SESSION_BASED"
	FunnelGetResponseFunnelTypeVisitorBased FunnelGetResponseFunnelType = "VISITOR_BASED"
)

type FunnelGetResponseStatus string

const (
	FunnelGetResponseStatusReady      FunnelGetResponseStatus = "READY"
	FunnelGetResponseStatusProcessing FunnelGetResponseStatus = "PROCESSING"
)

type FunnelGetResponseStep struct {
	EventName string `json:"eventName" api:"required"`
	Name      string `json:"name" api:"required"`
	Order     int64  `json:"order" api:"required"`
	StepID    string `json:"stepId" api:"required"`
	// Step-level event filters (JSON object).
	Filters any `json:"filters"`
	// Step-level event logic.
	Logic FunnelGetResponseStepLogic `json:"logic" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventName   respjson.Field
		Name        respjson.Field
		Order       respjson.Field
		StepID      respjson.Field
		Filters     respjson.Field
		Logic       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelGetResponseStep) RawJSON() string { return r.JSON.raw }
func (r *FunnelGetResponseStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Step-level event logic.
type FunnelGetResponseStepLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And       []any                               `json:"AND" api:"nullable"`
	Condition FunnelGetResponseStepLogicCondition `json:"condition" api:"nullable"`
	// Negates a single child logic node.
	Not any `json:"NOT"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or []any `json:"OR" api:"nullable"`
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
func (r FunnelGetResponseStepLogic) RawJSON() string { return r.JSON.raw }
func (r *FunnelGetResponseStepLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelGetResponseStepLogicCondition struct {
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
func (r FunnelGetResponseStepLogicCondition) RawJSON() string { return r.JSON.raw }
func (r *FunnelGetResponseStepLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelGetResponseConversionWindow struct {
	// Any of "MINUTES", "HOURS", "DAYS".
	Unit  string `json:"unit" api:"required"`
	Value int64  `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Unit        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelGetResponseConversionWindow) RawJSON() string { return r.JSON.raw }
func (r *FunnelGetResponseConversionWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelGetResponseCountingMethod string

const (
	FunnelGetResponseCountingMethodUniques  FunnelGetResponseCountingMethod = "UNIQUES"
	FunnelGetResponseCountingMethodTotals   FunnelGetResponseCountingMethod = "TOTALS"
	FunnelGetResponseCountingMethodSessions FunnelGetResponseCountingMethod = "SESSIONS"
)

// Nested visitor logic for the entire funnel. When supplied, this replaces legacy
// UTM filters.
type FunnelGetResponseGlobalLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And       []any                                 `json:"AND" api:"nullable"`
	Condition FunnelGetResponseGlobalLogicCondition `json:"condition" api:"nullable"`
	// Negates a single child logic node.
	Not any `json:"NOT"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or []any `json:"OR" api:"nullable"`
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
func (r FunnelGetResponseGlobalLogic) RawJSON() string { return r.JSON.raw }
func (r *FunnelGetResponseGlobalLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelGetResponseGlobalLogicCondition struct {
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
func (r FunnelGetResponseGlobalLogicCondition) RawJSON() string { return r.JSON.raw }
func (r *FunnelGetResponseGlobalLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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

type FunnelGetResponseStepOrder string

const (
	FunnelGetResponseStepOrderExact FunnelGetResponseStepOrder = "EXACT"
	FunnelGetResponseStepOrderAny   FunnelGetResponseStepOrder = "ANY"
)

type FunnelUpdateResponse struct {
	CreatedAt string `json:"createdAt" api:"required"`
	FunnelID  string `json:"funnelId" api:"required" format:"uuid"`
	// Any of "SESSION_BASED", "VISITOR_BASED".
	FunnelType FunnelUpdateResponseFunnelType `json:"funnelType" api:"required"`
	Name       string                         `json:"name" api:"required"`
	// Any of "READY", "PROCESSING".
	Status           FunnelUpdateResponseStatus           `json:"status" api:"required"`
	Steps            []FunnelUpdateResponseStep           `json:"steps" api:"required"`
	UpdatedAt        string                               `json:"updatedAt" api:"required"`
	ConversionWindow FunnelUpdateResponseConversionWindow `json:"conversionWindow" api:"nullable"`
	// Any of "UNIQUES", "TOTALS", "SESSIONS".
	CountingMethod FunnelUpdateResponseCountingMethod `json:"countingMethod" api:"nullable"`
	Description    string                             `json:"description" api:"nullable"`
	// Nested visitor logic for the entire funnel. When supplied, this replaces legacy
	// UTM filters.
	GlobalLogic     FunnelUpdateResponseGlobalLogic     `json:"globalLogic" api:"nullable"`
	ReportDateRange FunnelUpdateResponseReportDateRange `json:"reportDateRange" api:"nullable"`
	// Any of "EXACT", "ANY".
	StepOrder FunnelUpdateResponseStepOrder `json:"stepOrder" api:"nullable"`
	// Legacy exact-match UTM filters. Do not combine with globalLogic; globalLogic
	// takes precedence.
	UtmFilters any  `json:"utmFilters" api:"nullable"`
	Watched    bool `json:"watched" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt        respjson.Field
		FunnelID         respjson.Field
		FunnelType       respjson.Field
		Name             respjson.Field
		Status           respjson.Field
		Steps            respjson.Field
		UpdatedAt        respjson.Field
		ConversionWindow respjson.Field
		CountingMethod   respjson.Field
		Description      respjson.Field
		GlobalLogic      respjson.Field
		ReportDateRange  respjson.Field
		StepOrder        respjson.Field
		UtmFilters       respjson.Field
		Watched          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *FunnelUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelUpdateResponseFunnelType string

const (
	FunnelUpdateResponseFunnelTypeSessionBased FunnelUpdateResponseFunnelType = "SESSION_BASED"
	FunnelUpdateResponseFunnelTypeVisitorBased FunnelUpdateResponseFunnelType = "VISITOR_BASED"
)

type FunnelUpdateResponseStatus string

const (
	FunnelUpdateResponseStatusReady      FunnelUpdateResponseStatus = "READY"
	FunnelUpdateResponseStatusProcessing FunnelUpdateResponseStatus = "PROCESSING"
)

type FunnelUpdateResponseStep struct {
	EventName string `json:"eventName" api:"required"`
	Name      string `json:"name" api:"required"`
	Order     int64  `json:"order" api:"required"`
	StepID    string `json:"stepId" api:"required"`
	// Step-level event filters (JSON object).
	Filters any `json:"filters"`
	// Step-level event logic.
	Logic FunnelUpdateResponseStepLogic `json:"logic" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventName   respjson.Field
		Name        respjson.Field
		Order       respjson.Field
		StepID      respjson.Field
		Filters     respjson.Field
		Logic       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelUpdateResponseStep) RawJSON() string { return r.JSON.raw }
func (r *FunnelUpdateResponseStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Step-level event logic.
type FunnelUpdateResponseStepLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And       []any                                  `json:"AND" api:"nullable"`
	Condition FunnelUpdateResponseStepLogicCondition `json:"condition" api:"nullable"`
	// Negates a single child logic node.
	Not any `json:"NOT"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or []any `json:"OR" api:"nullable"`
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
func (r FunnelUpdateResponseStepLogic) RawJSON() string { return r.JSON.raw }
func (r *FunnelUpdateResponseStepLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelUpdateResponseStepLogicCondition struct {
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
func (r FunnelUpdateResponseStepLogicCondition) RawJSON() string { return r.JSON.raw }
func (r *FunnelUpdateResponseStepLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelUpdateResponseConversionWindow struct {
	// Any of "MINUTES", "HOURS", "DAYS".
	Unit  string `json:"unit" api:"required"`
	Value int64  `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Unit        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelUpdateResponseConversionWindow) RawJSON() string { return r.JSON.raw }
func (r *FunnelUpdateResponseConversionWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelUpdateResponseCountingMethod string

const (
	FunnelUpdateResponseCountingMethodUniques  FunnelUpdateResponseCountingMethod = "UNIQUES"
	FunnelUpdateResponseCountingMethodTotals   FunnelUpdateResponseCountingMethod = "TOTALS"
	FunnelUpdateResponseCountingMethodSessions FunnelUpdateResponseCountingMethod = "SESSIONS"
)

// Nested visitor logic for the entire funnel. When supplied, this replaces legacy
// UTM filters.
type FunnelUpdateResponseGlobalLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And       []any                                    `json:"AND" api:"nullable"`
	Condition FunnelUpdateResponseGlobalLogicCondition `json:"condition" api:"nullable"`
	// Negates a single child logic node.
	Not any `json:"NOT"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or []any `json:"OR" api:"nullable"`
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
func (r FunnelUpdateResponseGlobalLogic) RawJSON() string { return r.JSON.raw }
func (r *FunnelUpdateResponseGlobalLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelUpdateResponseGlobalLogicCondition struct {
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
func (r FunnelUpdateResponseGlobalLogicCondition) RawJSON() string { return r.JSON.raw }
func (r *FunnelUpdateResponseGlobalLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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

type FunnelUpdateResponseStepOrder string

const (
	FunnelUpdateResponseStepOrderExact FunnelUpdateResponseStepOrder = "EXACT"
	FunnelUpdateResponseStepOrderAny   FunnelUpdateResponseStepOrder = "ANY"
)

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
	// Any of "SESSION_BASED", "VISITOR_BASED".
	FunnelType FunnelDuplicateResponseFunnelType `json:"funnelType" api:"required"`
	Name       string                            `json:"name" api:"required"`
	// Any of "READY", "PROCESSING".
	Status           FunnelDuplicateResponseStatus           `json:"status" api:"required"`
	Steps            []FunnelDuplicateResponseStep           `json:"steps" api:"required"`
	UpdatedAt        string                                  `json:"updatedAt" api:"required"`
	ConversionWindow FunnelDuplicateResponseConversionWindow `json:"conversionWindow" api:"nullable"`
	// Any of "UNIQUES", "TOTALS", "SESSIONS".
	CountingMethod FunnelDuplicateResponseCountingMethod `json:"countingMethod" api:"nullable"`
	Description    string                                `json:"description" api:"nullable"`
	// Nested visitor logic for the entire funnel. When supplied, this replaces legacy
	// UTM filters.
	GlobalLogic     FunnelDuplicateResponseGlobalLogic     `json:"globalLogic" api:"nullable"`
	ReportDateRange FunnelDuplicateResponseReportDateRange `json:"reportDateRange" api:"nullable"`
	// Any of "EXACT", "ANY".
	StepOrder FunnelDuplicateResponseStepOrder `json:"stepOrder" api:"nullable"`
	// Legacy exact-match UTM filters. Do not combine with globalLogic; globalLogic
	// takes precedence.
	UtmFilters any  `json:"utmFilters" api:"nullable"`
	Watched    bool `json:"watched" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt        respjson.Field
		FunnelID         respjson.Field
		FunnelType       respjson.Field
		Name             respjson.Field
		Status           respjson.Field
		Steps            respjson.Field
		UpdatedAt        respjson.Field
		ConversionWindow respjson.Field
		CountingMethod   respjson.Field
		Description      respjson.Field
		GlobalLogic      respjson.Field
		ReportDateRange  respjson.Field
		StepOrder        respjson.Field
		UtmFilters       respjson.Field
		Watched          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelDuplicateResponse) RawJSON() string { return r.JSON.raw }
func (r *FunnelDuplicateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelDuplicateResponseFunnelType string

const (
	FunnelDuplicateResponseFunnelTypeSessionBased FunnelDuplicateResponseFunnelType = "SESSION_BASED"
	FunnelDuplicateResponseFunnelTypeVisitorBased FunnelDuplicateResponseFunnelType = "VISITOR_BASED"
)

type FunnelDuplicateResponseStatus string

const (
	FunnelDuplicateResponseStatusReady      FunnelDuplicateResponseStatus = "READY"
	FunnelDuplicateResponseStatusProcessing FunnelDuplicateResponseStatus = "PROCESSING"
)

type FunnelDuplicateResponseStep struct {
	EventName string `json:"eventName" api:"required"`
	Name      string `json:"name" api:"required"`
	Order     int64  `json:"order" api:"required"`
	StepID    string `json:"stepId" api:"required"`
	// Step-level event filters (JSON object).
	Filters any `json:"filters"`
	// Step-level event logic.
	Logic FunnelDuplicateResponseStepLogic `json:"logic" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventName   respjson.Field
		Name        respjson.Field
		Order       respjson.Field
		StepID      respjson.Field
		Filters     respjson.Field
		Logic       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelDuplicateResponseStep) RawJSON() string { return r.JSON.raw }
func (r *FunnelDuplicateResponseStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Step-level event logic.
type FunnelDuplicateResponseStepLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And       []any                                     `json:"AND" api:"nullable"`
	Condition FunnelDuplicateResponseStepLogicCondition `json:"condition" api:"nullable"`
	// Negates a single child logic node.
	Not any `json:"NOT"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or []any `json:"OR" api:"nullable"`
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
func (r FunnelDuplicateResponseStepLogic) RawJSON() string { return r.JSON.raw }
func (r *FunnelDuplicateResponseStepLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelDuplicateResponseStepLogicCondition struct {
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
func (r FunnelDuplicateResponseStepLogicCondition) RawJSON() string { return r.JSON.raw }
func (r *FunnelDuplicateResponseStepLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelDuplicateResponseConversionWindow struct {
	// Any of "MINUTES", "HOURS", "DAYS".
	Unit  string `json:"unit" api:"required"`
	Value int64  `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Unit        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelDuplicateResponseConversionWindow) RawJSON() string { return r.JSON.raw }
func (r *FunnelDuplicateResponseConversionWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelDuplicateResponseCountingMethod string

const (
	FunnelDuplicateResponseCountingMethodUniques  FunnelDuplicateResponseCountingMethod = "UNIQUES"
	FunnelDuplicateResponseCountingMethodTotals   FunnelDuplicateResponseCountingMethod = "TOTALS"
	FunnelDuplicateResponseCountingMethodSessions FunnelDuplicateResponseCountingMethod = "SESSIONS"
)

// Nested visitor logic for the entire funnel. When supplied, this replaces legacy
// UTM filters.
type FunnelDuplicateResponseGlobalLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And       []any                                       `json:"AND" api:"nullable"`
	Condition FunnelDuplicateResponseGlobalLogicCondition `json:"condition" api:"nullable"`
	// Negates a single child logic node.
	Not any `json:"NOT"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or []any `json:"OR" api:"nullable"`
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
func (r FunnelDuplicateResponseGlobalLogic) RawJSON() string { return r.JSON.raw }
func (r *FunnelDuplicateResponseGlobalLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelDuplicateResponseGlobalLogicCondition struct {
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
func (r FunnelDuplicateResponseGlobalLogicCondition) RawJSON() string { return r.JSON.raw }
func (r *FunnelDuplicateResponseGlobalLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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

type FunnelDuplicateResponseStepOrder string

const (
	FunnelDuplicateResponseStepOrderExact FunnelDuplicateResponseStepOrder = "EXACT"
	FunnelDuplicateResponseStepOrderAny   FunnelDuplicateResponseStepOrder = "ANY"
)

type FunnelResultsResponse struct {
	// Conversion rate from first step to last step as a percentage.
	OverallConversionRate float64 `json:"overallConversionRate" api:"required"`
	// Per-step funnel analytics, ordered by step number.
	Steps []FunnelResultsResponseStep `json:"steps" api:"required"`
	// Total number of visitors who entered the funnel (entered step 1).
	TotalVisitors int64 `json:"totalVisitors" api:"required"`
	// Average time from first step to last step in seconds. Null when no completions.
	OverallAvgTimeToConversion float64 `json:"overallAvgTimeToConversion" api:"nullable"`
	// Present when the results are wider than the funnel as configured. Some step
	// conditions could not be expressed as a query and were ignored, so the counts
	// above include visitors those conditions would have excluded. Absent when the
	// whole definition was applied. The dashboard surfaces the same caveat.
	Warning string `json:"warning"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OverallConversionRate      respjson.Field
		Steps                      respjson.Field
		TotalVisitors              respjson.Field
		OverallAvgTimeToConversion respjson.Field
		Warning                    respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunnelResultsResponse) RawJSON() string { return r.JSON.raw }
func (r *FunnelResultsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelResultsResponseStep struct {
	ConversionCount       int64    `json:"conversionCount" api:"required"`
	ConversionRate        float64  `json:"conversionRate" api:"required"`
	DropOffRate           float64  `json:"dropOffRate" api:"required"`
	DropOffSessionIDs     []string `json:"dropOffSessionIds" api:"required"`
	OverallConversionRate float64  `json:"overallConversionRate" api:"required"`
	SessionIDs            []string `json:"sessionIds" api:"required"`
	StepNumber            int64    `json:"stepNumber" api:"required"`
	VisitorCount          int64    `json:"visitorCount" api:"required"`
	AvgTimeToNextStep     float64  `json:"avgTimeToNextStep" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConversionCount       respjson.Field
		ConversionRate        respjson.Field
		DropOffRate           respjson.Field
		DropOffSessionIDs     respjson.Field
		OverallConversionRate respjson.Field
		SessionIDs            respjson.Field
		StepNumber            respjson.Field
		VisitorCount          respjson.Field
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
	Name             string                `json:"name" api:"required"`
	Steps            []FunnelNewParamsStep `json:"steps,omitzero" api:"required"`
	CountingMethod   param.Opt[string]     `json:"countingMethod,omitzero"`
	Description      param.Opt[string]     `json:"description,omitzero"`
	StepOrder        param.Opt[string]     `json:"stepOrder,omitzero"`
	Watched          param.Opt[bool]       `json:"watched,omitzero"`
	ConversionWindow any                   `json:"conversionWindow,omitzero"`
	// Nested visitor logic for the entire funnel. When supplied, this replaces legacy
	// UTM filters.
	GlobalLogic FunnelNewParamsGlobalLogic `json:"globalLogic,omitzero"`
	// Legacy exact-match UTM filters. Do not combine with globalLogic; globalLogic
	// takes precedence.
	UtmFilters any `json:"utmFilters,omitzero"`
	// Funnels are session-based. `SESSION_BASED` is the only supported value and is
	// applied when omitted.
	//
	// Any of "SESSION_BASED".
	FunnelType FunnelNewParamsFunnelType `json:"funnelType,omitzero"`
	paramObj
}

func (r FunnelNewParams) MarshalJSON() (data []byte, err error) {
	type shadow FunnelNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties EventName, Name, Order are required.
type FunnelNewParamsStep struct {
	EventName string                   `json:"eventName" api:"required"`
	Name      string                   `json:"name" api:"required"`
	Order     int64                    `json:"order" api:"required"`
	Filters   any                      `json:"filters,omitzero"`
	Logic     FunnelNewParamsStepLogic `json:"logic,omitzero"`
	paramObj
}

func (r FunnelNewParamsStep) MarshalJSON() (data []byte, err error) {
	type shadow FunnelNewParamsStep
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelNewParamsStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelNewParamsStepLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And       []any                             `json:"AND,omitzero"`
	Condition FunnelNewParamsStepLogicCondition `json:"condition,omitzero"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or []any `json:"OR,omitzero"`
	// Negates a single child logic node.
	Not any `json:"NOT,omitzero"`
	paramObj
}

func (r FunnelNewParamsStepLogic) MarshalJSON() (data []byte, err error) {
	type shadow FunnelNewParamsStepLogic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelNewParamsStepLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Operator, Property, Value are required.
type FunnelNewParamsStepLogicCondition struct {
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

func (r FunnelNewParamsStepLogicCondition) MarshalJSON() (data []byte, err error) {
	type shadow FunnelNewParamsStepLogicCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelNewParamsStepLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelNewParamsStepLogicCondition](
		"operator", "Is", "IsNot", "Contains", "DoesNotContain", "StartsWith", "EndsWith", "IsFalsy", "IsTruthy", "IsNull", "IsNotNull", "IsUndefined", "IsNotUndefined", "IsGreaterThan", "IsGreaterThanOrEqual", "IsLessThan", "IsLessThanOrEqual", "IsIn", "IsNotIn", "IsFoundIn", "IsNotFoundIn", "IsTrue", "IsFalse", "IsBefore", "IsAfter", "IsBetween", "IsOnOrBefore", "IsOnOrAfter", "MatchesRegex", "MatchesRegexIgnoreCase", "DoesNotMatchRegex", "DoesNotMatchRegexIgnoreCase",
	)
}

// Funnels are session-based. `SESSION_BASED` is the only supported value and is
// applied when omitted.
type FunnelNewParamsFunnelType string

const (
	FunnelNewParamsFunnelTypeSessionBased FunnelNewParamsFunnelType = "SESSION_BASED"
)

// Nested visitor logic for the entire funnel. When supplied, this replaces legacy
// UTM filters.
type FunnelNewParamsGlobalLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And       []any                               `json:"AND,omitzero"`
	Condition FunnelNewParamsGlobalLogicCondition `json:"condition,omitzero"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or []any `json:"OR,omitzero"`
	// Negates a single child logic node.
	Not any `json:"NOT,omitzero"`
	paramObj
}

func (r FunnelNewParamsGlobalLogic) MarshalJSON() (data []byte, err error) {
	type shadow FunnelNewParamsGlobalLogic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelNewParamsGlobalLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Operator, Property, Value are required.
type FunnelNewParamsGlobalLogicCondition struct {
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

func (r FunnelNewParamsGlobalLogicCondition) MarshalJSON() (data []byte, err error) {
	type shadow FunnelNewParamsGlobalLogicCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelNewParamsGlobalLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelNewParamsGlobalLogicCondition](
		"operator", "Is", "IsNot", "Contains", "DoesNotContain", "StartsWith", "EndsWith", "IsFalsy", "IsTruthy", "IsNull", "IsNotNull", "IsUndefined", "IsNotUndefined", "IsGreaterThan", "IsGreaterThanOrEqual", "IsLessThan", "IsLessThanOrEqual", "IsIn", "IsNotIn", "IsFoundIn", "IsNotFoundIn", "IsTrue", "IsFalse", "IsBefore", "IsAfter", "IsBetween", "IsOnOrBefore", "IsOnOrAfter", "MatchesRegex", "MatchesRegexIgnoreCase", "DoesNotMatchRegex", "DoesNotMatchRegexIgnoreCase",
	)
}

type FunnelUpdateParams struct {
	CountingMethod   param.Opt[string] `json:"countingMethod,omitzero"`
	Description      param.Opt[string] `json:"description,omitzero"`
	StepOrder        param.Opt[string] `json:"stepOrder,omitzero"`
	Watched          param.Opt[bool]   `json:"watched,omitzero"`
	Name             param.Opt[string] `json:"name,omitzero"`
	ConversionWindow any               `json:"conversionWindow,omitzero"`
	// Nested visitor logic for the entire funnel. When supplied, this replaces legacy
	// UTM filters.
	GlobalLogic FunnelUpdateParamsGlobalLogic `json:"globalLogic,omitzero"`
	// Legacy exact-match UTM filters. Do not combine with globalLogic; globalLogic
	// takes precedence.
	UtmFilters any `json:"utmFilters,omitzero"`
	// Any of "SESSION_BASED".
	FunnelType FunnelUpdateParamsFunnelType `json:"funnelType,omitzero"`
	Steps      []FunnelUpdateParamsStep     `json:"steps,omitzero"`
	paramObj
}

func (r FunnelUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow FunnelUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelUpdateParamsFunnelType string

const (
	FunnelUpdateParamsFunnelTypeSessionBased FunnelUpdateParamsFunnelType = "SESSION_BASED"
)

// Nested visitor logic for the entire funnel. When supplied, this replaces legacy
// UTM filters.
type FunnelUpdateParamsGlobalLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And       []any                                  `json:"AND,omitzero"`
	Condition FunnelUpdateParamsGlobalLogicCondition `json:"condition,omitzero"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or []any `json:"OR,omitzero"`
	// Negates a single child logic node.
	Not any `json:"NOT,omitzero"`
	paramObj
}

func (r FunnelUpdateParamsGlobalLogic) MarshalJSON() (data []byte, err error) {
	type shadow FunnelUpdateParamsGlobalLogic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelUpdateParamsGlobalLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Operator, Property, Value are required.
type FunnelUpdateParamsGlobalLogicCondition struct {
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

func (r FunnelUpdateParamsGlobalLogicCondition) MarshalJSON() (data []byte, err error) {
	type shadow FunnelUpdateParamsGlobalLogicCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelUpdateParamsGlobalLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelUpdateParamsGlobalLogicCondition](
		"operator", "Is", "IsNot", "Contains", "DoesNotContain", "StartsWith", "EndsWith", "IsFalsy", "IsTruthy", "IsNull", "IsNotNull", "IsUndefined", "IsNotUndefined", "IsGreaterThan", "IsGreaterThanOrEqual", "IsLessThan", "IsLessThanOrEqual", "IsIn", "IsNotIn", "IsFoundIn", "IsNotFoundIn", "IsTrue", "IsFalse", "IsBefore", "IsAfter", "IsBetween", "IsOnOrBefore", "IsOnOrAfter", "MatchesRegex", "MatchesRegexIgnoreCase", "DoesNotMatchRegex", "DoesNotMatchRegexIgnoreCase",
	)
}

// The properties EventName, Name, Order are required.
type FunnelUpdateParamsStep struct {
	EventName string                      `json:"eventName" api:"required"`
	Name      string                      `json:"name" api:"required"`
	Order     int64                       `json:"order" api:"required"`
	Filters   any                         `json:"filters,omitzero"`
	Logic     FunnelUpdateParamsStepLogic `json:"logic,omitzero"`
	paramObj
}

func (r FunnelUpdateParamsStep) MarshalJSON() (data []byte, err error) {
	type shadow FunnelUpdateParamsStep
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelUpdateParamsStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunnelUpdateParamsStepLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And       []any                                `json:"AND,omitzero"`
	Condition FunnelUpdateParamsStepLogicCondition `json:"condition,omitzero"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or []any `json:"OR,omitzero"`
	// Negates a single child logic node.
	Not any `json:"NOT,omitzero"`
	paramObj
}

func (r FunnelUpdateParamsStepLogic) MarshalJSON() (data []byte, err error) {
	type shadow FunnelUpdateParamsStepLogic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelUpdateParamsStepLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Operator, Property, Value are required.
type FunnelUpdateParamsStepLogicCondition struct {
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

func (r FunnelUpdateParamsStepLogicCondition) MarshalJSON() (data []byte, err error) {
	type shadow FunnelUpdateParamsStepLogicCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunnelUpdateParamsStepLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FunnelUpdateParamsStepLogicCondition](
		"operator", "Is", "IsNot", "Contains", "DoesNotContain", "StartsWith", "EndsWith", "IsFalsy", "IsTruthy", "IsNull", "IsNotNull", "IsUndefined", "IsNotUndefined", "IsGreaterThan", "IsGreaterThanOrEqual", "IsLessThan", "IsLessThanOrEqual", "IsIn", "IsNotIn", "IsFoundIn", "IsNotFoundIn", "IsTrue", "IsFalse", "IsBefore", "IsAfter", "IsBetween", "IsOnOrBefore", "IsOnOrAfter", "MatchesRegex", "MatchesRegexIgnoreCase", "DoesNotMatchRegex", "DoesNotMatchRegexIgnoreCase",
	)
}

type FunnelResultsParams struct {
	// Inclusive lower bound of the analysis window, as a UTC calendar day in
	// `YYYY-MM-DD` format. The window may span at most 31 days including both
	// endpoints.
	From string `query:"from" api:"required" json:"-"`
	// Inclusive upper bound of the analysis window, as a UTC calendar day in
	// `YYYY-MM-DD` format. Must be on or after `from`, and the window may span at most
	// 31 days including both endpoints.
	To string `query:"to" api:"required" json:"-"`
	// Restrict the funnel to sessions whose `utm_campaign` exactly matches this value.
	UtmCampaign param.Opt[string] `query:"utmCampaign,omitzero" json:"-"`
	// Restrict the funnel to sessions whose `utm_content` exactly matches this value.
	UtmContent param.Opt[string] `query:"utmContent,omitzero" json:"-"`
	// Restrict the funnel to sessions whose `utm_medium` exactly matches this value.
	UtmMedium param.Opt[string] `query:"utmMedium,omitzero" json:"-"`
	// Accepted for backward compatibility but NOT applied — there is no campaign-name
	// dimension on funnel sessions. Use `utmCampaign` instead.
	UtmName param.Opt[string] `query:"utmName,omitzero" json:"-"`
	// Restrict the funnel to sessions whose `utm_source` exactly matches this value.
	UtmSource param.Opt[string] `query:"utmSource,omitzero" json:"-"`
	// Restrict the funnel to sessions whose `utm_term` exactly matches this value.
	UtmTerm param.Opt[string] `query:"utmTerm,omitzero" json:"-"`
	// Accepted for backward compatibility but NOT applied. Funnel sessions carry a
	// single attribution set, so there is no initial vs. last-touch distinction to
	// select between.
	//
	// Any of "INITIAL", "LAST_TOUCH".
	AttributionType FunnelResultsParamsAttributionType `query:"attributionType,omitzero" json:"-"`
	// Restrict the funnel to sessions on a device class. `MOBILE` matches phone
	// sessions; `DESKTOP` matches every session that is not a phone, tablet, TV,
	// console, wearable, XR, or embedded device. `ALL` (the default) applies no device
	// filter.
	//
	// Any of "DESKTOP", "MOBILE", "ALL".
	DeviceType FunnelResultsParamsDeviceType `query:"deviceType,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FunnelResultsParams]'s query parameters as `url.Values`.
func (r FunnelResultsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Accepted for backward compatibility but NOT applied. Funnel sessions carry a
// single attribution set, so there is no initial vs. last-touch distinction to
// select between.
type FunnelResultsParamsAttributionType string

const (
	FunnelResultsParamsAttributionTypeInitial   FunnelResultsParamsAttributionType = "INITIAL"
	FunnelResultsParamsAttributionTypeLastTouch FunnelResultsParamsAttributionType = "LAST_TOUCH"
)

// Restrict the funnel to sessions on a device class. `MOBILE` matches phone
// sessions; `DESKTOP` matches every session that is not a phone, tablet, TV,
// console, wearable, XR, or embedded device. `ALL` (the default) applies no device
// filter.
type FunnelResultsParamsDeviceType string

const (
	FunnelResultsParamsDeviceTypeDesktop FunnelResultsParamsDeviceType = "DESKTOP"
	FunnelResultsParamsDeviceTypeMobile  FunnelResultsParamsDeviceType = "MOBILE"
	FunnelResultsParamsDeviceTypeAll     FunnelResultsParamsDeviceType = "ALL"
)
