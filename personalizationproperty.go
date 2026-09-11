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
	"github.com/with-ours/platform-sdk-go/packages/pagination"
	"github.com/with-ours/platform-sdk-go/packages/param"
	"github.com/with-ours/platform-sdk-go/packages/respjson"
)

// PersonalizationPropertyService contains methods and other services that help
// with interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPersonalizationPropertyService] method instead.
type PersonalizationPropertyService struct {
	Options []option.RequestOption
}

// NewPersonalizationPropertyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPersonalizationPropertyService(opts ...option.RequestOption) (r PersonalizationPropertyService) {
	r = PersonalizationPropertyService{}
	r.Options = opts
	return
}

// List personalization properties for one experiment settings record. Requires the
// `experimentSettingsId` query parameter — properties are always scoped to a
// single record; list the records with `GET /rest/v1/experiment-settings`.
// Supports cursor pagination via `limit` and `cursor`; the limit clamp is 1000 so
// a single request can return the full set. Requires scope:
// experimentSettings:list
func (r *PersonalizationPropertyService) List(ctx context.Context, query PersonalizationPropertyListParams, opts ...option.RequestOption) (res *pagination.Cursor[PersonalizationPropertyListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "rest/v1/personalization-properties"
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

// List personalization properties for one experiment settings record. Requires the
// `experimentSettingsId` query parameter — properties are always scoped to a
// single record; list the records with `GET /rest/v1/experiment-settings`.
// Supports cursor pagination via `limit` and `cursor`; the limit clamp is 1000 so
// a single request can return the full set. Requires scope:
// experimentSettings:list
func (r *PersonalizationPropertyService) ListAutoPaging(ctx context.Context, query PersonalizationPropertyListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[PersonalizationPropertyListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Create a personalization property on an experiment settings record. The new rule
// is published automatically and starts accumulating from the next matching event
// — no separate publish call is needed. `propertyKey` must be unique within the
// parent record. Requires scope: experimentSettings:update
func (r *PersonalizationPropertyService) New(ctx context.Context, body PersonalizationPropertyNewParams, opts ...option.RequestOption) (res *PersonalizationPropertyNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/personalization-properties"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Find a single personalization property by ID. Returns 404 when no property
// matches the supplied id. Requires scope: experimentSettings:find
func (r *PersonalizationPropertyService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PersonalizationPropertyGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/personalization-properties/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Partially update a personalization property. Only the fields you send are
// changed, and the update is published automatically. Sending `triggerConditions`
// replaces the prior list — partial-array merging is not supported. Values already
// accumulated for visitors are kept; the new rule applies to events from here on.
// Returns 404 when no property matches the supplied id. Requires scope:
// experimentSettings:update
func (r *PersonalizationPropertyService) Update(ctx context.Context, id string, body PersonalizationPropertyUpdateParams, opts ...option.RequestOption) (res *PersonalizationPropertyUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/personalization-properties/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete a personalization property. The rule stops accumulating immediately;
// values already recorded for visitors are no longer maintained. Returns 404 when
// no property matches the supplied id. Requires scope: experimentSettings:update
func (r *PersonalizationPropertyService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/personalization-properties/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type PersonalizationPropertyListResponse struct {
	// Unique identifier for this personalization property.
	ID string `json:"id" api:"required"`
	// How repeated trigger events fold into the stored value. `set_true` records
	// `true` the first time the event matches — use it for one-shot interest flags.
	// `increment` counts matching events as a number. `latest_value` stores the most
	// recent value read from `valueField`, replacing whatever was there before.
	//
	// Any of "increment", "latest_value", "set_true".
	Accumulator PersonalizationPropertyListResponseAccumulator `json:"accumulator" api:"required"`
	// ISO 8601 timestamp of when the property was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Experiment settings record that owns this property.
	ExperimentSettingsID string `json:"experimentSettingsId" api:"required"`
	// Name this value is stored and read under. Must be unique within the parent
	// experiment settings record — a duplicate key is rejected with 409. Because
	// accumulated values are delivered to the visitor's browser and are readable
	// there, never accumulate secrets, credentials, PHI, or confidential data into a
	// property.
	PropertyKey string `json:"propertyKey" api:"required"`
	// Whether the property is accumulating. `Disabled` properties stop accumulating
	// but keep values already recorded.
	//
	// Any of "Disabled", "Enabled".
	Status PersonalizationPropertyListResponseStatus `json:"status" api:"required"`
	// Name of the tracked event that advances this property. Every event with this
	// name is evaluated against `triggerConditions`.
	TriggerEventName string `json:"triggerEventName" api:"required"`
	// Optional filters the triggering event must satisfy. All conditions must match
	// (AND). Omit or send an empty list to accumulate on every event with the matching
	// name.
	TriggerConditions []PersonalizationPropertyListResponseTriggerCondition `json:"triggerConditions" api:"nullable"`
	// ISO 8601 timestamp of the last update, or null when never updated.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// Event field to read the stored value from. Required for `latest_value`; ignored
	// by `set_true` and `increment`. Only scalar values (string, number, boolean) are
	// stored. Must be one of `event.name`, `event.properties.<name>`, or
	// `event.context.<name>` where `<name>` is one of `current_url`, `referrer`,
	// `utm_source`, `utm_medium`, `utm_campaign`, `utm_content`, or `utm_term`. Any
	// other path is rejected with 400.
	ValueField string `json:"valueField" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Accumulator          respjson.Field
		CreatedAt            respjson.Field
		ExperimentSettingsID respjson.Field
		PropertyKey          respjson.Field
		Status               respjson.Field
		TriggerEventName     respjson.Field
		TriggerConditions    respjson.Field
		UpdatedAt            respjson.Field
		ValueField           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonalizationPropertyListResponse) RawJSON() string { return r.JSON.raw }
func (r *PersonalizationPropertyListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How repeated trigger events fold into the stored value. `set_true` records
// `true` the first time the event matches — use it for one-shot interest flags.
// `increment` counts matching events as a number. `latest_value` stores the most
// recent value read from `valueField`, replacing whatever was there before.
type PersonalizationPropertyListResponseAccumulator string

const (
	PersonalizationPropertyListResponseAccumulatorIncrement   PersonalizationPropertyListResponseAccumulator = "increment"
	PersonalizationPropertyListResponseAccumulatorLatestValue PersonalizationPropertyListResponseAccumulator = "latest_value"
	PersonalizationPropertyListResponseAccumulatorSetTrue     PersonalizationPropertyListResponseAccumulator = "set_true"
)

// Whether the property is accumulating. `Disabled` properties stop accumulating
// but keep values already recorded.
type PersonalizationPropertyListResponseStatus string

const (
	PersonalizationPropertyListResponseStatusDisabled PersonalizationPropertyListResponseStatus = "Disabled"
	PersonalizationPropertyListResponseStatusEnabled  PersonalizationPropertyListResponseStatus = "Enabled"
)

type PersonalizationPropertyListResponseTriggerCondition struct {
	// Event field the condition reads. Must be one of `event.name`,
	// `event.properties.<name>`, or `event.context.<name>` where `<name>` is one of
	// `current_url`, `referrer`, `utm_source`, `utm_medium`, `utm_campaign`,
	// `utm_content`, or `utm_term`. Any other path is rejected with 400.
	Field string `json:"field" api:"required"`
	// Comparison applied to the field. `exists` and `not_exists` ignore `value`;
	// `regex` matches the field against `value` as a regular expression.
	//
	// Any of "contains", "equals", "exists", "not_equals", "not_exists", "regex".
	Operator string `json:"operator" api:"required"`
	// Scalar value the operator compares against. Omit for `exists` and `not_exists`.
	// Nested objects and arrays are rejected.
	Value string `json:"value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonalizationPropertyListResponseTriggerCondition) RawJSON() string { return r.JSON.raw }
func (r *PersonalizationPropertyListResponseTriggerCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonalizationPropertyNewResponse struct {
	// Unique identifier for this personalization property.
	ID string `json:"id" api:"required"`
	// How repeated trigger events fold into the stored value. `set_true` records
	// `true` the first time the event matches — use it for one-shot interest flags.
	// `increment` counts matching events as a number. `latest_value` stores the most
	// recent value read from `valueField`, replacing whatever was there before.
	//
	// Any of "increment", "latest_value", "set_true".
	Accumulator PersonalizationPropertyNewResponseAccumulator `json:"accumulator" api:"required"`
	// ISO 8601 timestamp of when the property was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Experiment settings record that owns this property.
	ExperimentSettingsID string `json:"experimentSettingsId" api:"required"`
	// Name this value is stored and read under. Must be unique within the parent
	// experiment settings record — a duplicate key is rejected with 409. Because
	// accumulated values are delivered to the visitor's browser and are readable
	// there, never accumulate secrets, credentials, PHI, or confidential data into a
	// property.
	PropertyKey string `json:"propertyKey" api:"required"`
	// Whether the property is accumulating. `Disabled` properties stop accumulating
	// but keep values already recorded.
	//
	// Any of "Disabled", "Enabled".
	Status PersonalizationPropertyNewResponseStatus `json:"status" api:"required"`
	// Name of the tracked event that advances this property. Every event with this
	// name is evaluated against `triggerConditions`.
	TriggerEventName string `json:"triggerEventName" api:"required"`
	// Optional filters the triggering event must satisfy. All conditions must match
	// (AND). Omit or send an empty list to accumulate on every event with the matching
	// name.
	TriggerConditions []PersonalizationPropertyNewResponseTriggerCondition `json:"triggerConditions" api:"nullable"`
	// ISO 8601 timestamp of the last update, or null when never updated.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// Event field to read the stored value from. Required for `latest_value`; ignored
	// by `set_true` and `increment`. Only scalar values (string, number, boolean) are
	// stored. Must be one of `event.name`, `event.properties.<name>`, or
	// `event.context.<name>` where `<name>` is one of `current_url`, `referrer`,
	// `utm_source`, `utm_medium`, `utm_campaign`, `utm_content`, or `utm_term`. Any
	// other path is rejected with 400.
	ValueField string `json:"valueField" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Accumulator          respjson.Field
		CreatedAt            respjson.Field
		ExperimentSettingsID respjson.Field
		PropertyKey          respjson.Field
		Status               respjson.Field
		TriggerEventName     respjson.Field
		TriggerConditions    respjson.Field
		UpdatedAt            respjson.Field
		ValueField           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonalizationPropertyNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PersonalizationPropertyNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How repeated trigger events fold into the stored value. `set_true` records
// `true` the first time the event matches — use it for one-shot interest flags.
// `increment` counts matching events as a number. `latest_value` stores the most
// recent value read from `valueField`, replacing whatever was there before.
type PersonalizationPropertyNewResponseAccumulator string

const (
	PersonalizationPropertyNewResponseAccumulatorIncrement   PersonalizationPropertyNewResponseAccumulator = "increment"
	PersonalizationPropertyNewResponseAccumulatorLatestValue PersonalizationPropertyNewResponseAccumulator = "latest_value"
	PersonalizationPropertyNewResponseAccumulatorSetTrue     PersonalizationPropertyNewResponseAccumulator = "set_true"
)

// Whether the property is accumulating. `Disabled` properties stop accumulating
// but keep values already recorded.
type PersonalizationPropertyNewResponseStatus string

const (
	PersonalizationPropertyNewResponseStatusDisabled PersonalizationPropertyNewResponseStatus = "Disabled"
	PersonalizationPropertyNewResponseStatusEnabled  PersonalizationPropertyNewResponseStatus = "Enabled"
)

type PersonalizationPropertyNewResponseTriggerCondition struct {
	// Event field the condition reads. Must be one of `event.name`,
	// `event.properties.<name>`, or `event.context.<name>` where `<name>` is one of
	// `current_url`, `referrer`, `utm_source`, `utm_medium`, `utm_campaign`,
	// `utm_content`, or `utm_term`. Any other path is rejected with 400.
	Field string `json:"field" api:"required"`
	// Comparison applied to the field. `exists` and `not_exists` ignore `value`;
	// `regex` matches the field against `value` as a regular expression.
	//
	// Any of "contains", "equals", "exists", "not_equals", "not_exists", "regex".
	Operator string `json:"operator" api:"required"`
	// Scalar value the operator compares against. Omit for `exists` and `not_exists`.
	// Nested objects and arrays are rejected.
	Value string `json:"value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonalizationPropertyNewResponseTriggerCondition) RawJSON() string { return r.JSON.raw }
func (r *PersonalizationPropertyNewResponseTriggerCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonalizationPropertyGetResponse struct {
	// Unique identifier for this personalization property.
	ID string `json:"id" api:"required"`
	// How repeated trigger events fold into the stored value. `set_true` records
	// `true` the first time the event matches — use it for one-shot interest flags.
	// `increment` counts matching events as a number. `latest_value` stores the most
	// recent value read from `valueField`, replacing whatever was there before.
	//
	// Any of "increment", "latest_value", "set_true".
	Accumulator PersonalizationPropertyGetResponseAccumulator `json:"accumulator" api:"required"`
	// ISO 8601 timestamp of when the property was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Experiment settings record that owns this property.
	ExperimentSettingsID string `json:"experimentSettingsId" api:"required"`
	// Name this value is stored and read under. Must be unique within the parent
	// experiment settings record — a duplicate key is rejected with 409. Because
	// accumulated values are delivered to the visitor's browser and are readable
	// there, never accumulate secrets, credentials, PHI, or confidential data into a
	// property.
	PropertyKey string `json:"propertyKey" api:"required"`
	// Whether the property is accumulating. `Disabled` properties stop accumulating
	// but keep values already recorded.
	//
	// Any of "Disabled", "Enabled".
	Status PersonalizationPropertyGetResponseStatus `json:"status" api:"required"`
	// Name of the tracked event that advances this property. Every event with this
	// name is evaluated against `triggerConditions`.
	TriggerEventName string `json:"triggerEventName" api:"required"`
	// Optional filters the triggering event must satisfy. All conditions must match
	// (AND). Omit or send an empty list to accumulate on every event with the matching
	// name.
	TriggerConditions []PersonalizationPropertyGetResponseTriggerCondition `json:"triggerConditions" api:"nullable"`
	// ISO 8601 timestamp of the last update, or null when never updated.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// Event field to read the stored value from. Required for `latest_value`; ignored
	// by `set_true` and `increment`. Only scalar values (string, number, boolean) are
	// stored. Must be one of `event.name`, `event.properties.<name>`, or
	// `event.context.<name>` where `<name>` is one of `current_url`, `referrer`,
	// `utm_source`, `utm_medium`, `utm_campaign`, `utm_content`, or `utm_term`. Any
	// other path is rejected with 400.
	ValueField string `json:"valueField" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Accumulator          respjson.Field
		CreatedAt            respjson.Field
		ExperimentSettingsID respjson.Field
		PropertyKey          respjson.Field
		Status               respjson.Field
		TriggerEventName     respjson.Field
		TriggerConditions    respjson.Field
		UpdatedAt            respjson.Field
		ValueField           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonalizationPropertyGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PersonalizationPropertyGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How repeated trigger events fold into the stored value. `set_true` records
// `true` the first time the event matches — use it for one-shot interest flags.
// `increment` counts matching events as a number. `latest_value` stores the most
// recent value read from `valueField`, replacing whatever was there before.
type PersonalizationPropertyGetResponseAccumulator string

const (
	PersonalizationPropertyGetResponseAccumulatorIncrement   PersonalizationPropertyGetResponseAccumulator = "increment"
	PersonalizationPropertyGetResponseAccumulatorLatestValue PersonalizationPropertyGetResponseAccumulator = "latest_value"
	PersonalizationPropertyGetResponseAccumulatorSetTrue     PersonalizationPropertyGetResponseAccumulator = "set_true"
)

// Whether the property is accumulating. `Disabled` properties stop accumulating
// but keep values already recorded.
type PersonalizationPropertyGetResponseStatus string

const (
	PersonalizationPropertyGetResponseStatusDisabled PersonalizationPropertyGetResponseStatus = "Disabled"
	PersonalizationPropertyGetResponseStatusEnabled  PersonalizationPropertyGetResponseStatus = "Enabled"
)

type PersonalizationPropertyGetResponseTriggerCondition struct {
	// Event field the condition reads. Must be one of `event.name`,
	// `event.properties.<name>`, or `event.context.<name>` where `<name>` is one of
	// `current_url`, `referrer`, `utm_source`, `utm_medium`, `utm_campaign`,
	// `utm_content`, or `utm_term`. Any other path is rejected with 400.
	Field string `json:"field" api:"required"`
	// Comparison applied to the field. `exists` and `not_exists` ignore `value`;
	// `regex` matches the field against `value` as a regular expression.
	//
	// Any of "contains", "equals", "exists", "not_equals", "not_exists", "regex".
	Operator string `json:"operator" api:"required"`
	// Scalar value the operator compares against. Omit for `exists` and `not_exists`.
	// Nested objects and arrays are rejected.
	Value string `json:"value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonalizationPropertyGetResponseTriggerCondition) RawJSON() string { return r.JSON.raw }
func (r *PersonalizationPropertyGetResponseTriggerCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonalizationPropertyUpdateResponse struct {
	// Unique identifier for this personalization property.
	ID string `json:"id" api:"required"`
	// How repeated trigger events fold into the stored value. `set_true` records
	// `true` the first time the event matches — use it for one-shot interest flags.
	// `increment` counts matching events as a number. `latest_value` stores the most
	// recent value read from `valueField`, replacing whatever was there before.
	//
	// Any of "increment", "latest_value", "set_true".
	Accumulator PersonalizationPropertyUpdateResponseAccumulator `json:"accumulator" api:"required"`
	// ISO 8601 timestamp of when the property was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Experiment settings record that owns this property.
	ExperimentSettingsID string `json:"experimentSettingsId" api:"required"`
	// Name this value is stored and read under. Must be unique within the parent
	// experiment settings record — a duplicate key is rejected with 409. Because
	// accumulated values are delivered to the visitor's browser and are readable
	// there, never accumulate secrets, credentials, PHI, or confidential data into a
	// property.
	PropertyKey string `json:"propertyKey" api:"required"`
	// Whether the property is accumulating. `Disabled` properties stop accumulating
	// but keep values already recorded.
	//
	// Any of "Disabled", "Enabled".
	Status PersonalizationPropertyUpdateResponseStatus `json:"status" api:"required"`
	// Name of the tracked event that advances this property. Every event with this
	// name is evaluated against `triggerConditions`.
	TriggerEventName string `json:"triggerEventName" api:"required"`
	// Optional filters the triggering event must satisfy. All conditions must match
	// (AND). Omit or send an empty list to accumulate on every event with the matching
	// name.
	TriggerConditions []PersonalizationPropertyUpdateResponseTriggerCondition `json:"triggerConditions" api:"nullable"`
	// ISO 8601 timestamp of the last update, or null when never updated.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// Event field to read the stored value from. Required for `latest_value`; ignored
	// by `set_true` and `increment`. Only scalar values (string, number, boolean) are
	// stored. Must be one of `event.name`, `event.properties.<name>`, or
	// `event.context.<name>` where `<name>` is one of `current_url`, `referrer`,
	// `utm_source`, `utm_medium`, `utm_campaign`, `utm_content`, or `utm_term`. Any
	// other path is rejected with 400.
	ValueField string `json:"valueField" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Accumulator          respjson.Field
		CreatedAt            respjson.Field
		ExperimentSettingsID respjson.Field
		PropertyKey          respjson.Field
		Status               respjson.Field
		TriggerEventName     respjson.Field
		TriggerConditions    respjson.Field
		UpdatedAt            respjson.Field
		ValueField           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonalizationPropertyUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *PersonalizationPropertyUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How repeated trigger events fold into the stored value. `set_true` records
// `true` the first time the event matches — use it for one-shot interest flags.
// `increment` counts matching events as a number. `latest_value` stores the most
// recent value read from `valueField`, replacing whatever was there before.
type PersonalizationPropertyUpdateResponseAccumulator string

const (
	PersonalizationPropertyUpdateResponseAccumulatorIncrement   PersonalizationPropertyUpdateResponseAccumulator = "increment"
	PersonalizationPropertyUpdateResponseAccumulatorLatestValue PersonalizationPropertyUpdateResponseAccumulator = "latest_value"
	PersonalizationPropertyUpdateResponseAccumulatorSetTrue     PersonalizationPropertyUpdateResponseAccumulator = "set_true"
)

// Whether the property is accumulating. `Disabled` properties stop accumulating
// but keep values already recorded.
type PersonalizationPropertyUpdateResponseStatus string

const (
	PersonalizationPropertyUpdateResponseStatusDisabled PersonalizationPropertyUpdateResponseStatus = "Disabled"
	PersonalizationPropertyUpdateResponseStatusEnabled  PersonalizationPropertyUpdateResponseStatus = "Enabled"
)

type PersonalizationPropertyUpdateResponseTriggerCondition struct {
	// Event field the condition reads. Must be one of `event.name`,
	// `event.properties.<name>`, or `event.context.<name>` where `<name>` is one of
	// `current_url`, `referrer`, `utm_source`, `utm_medium`, `utm_campaign`,
	// `utm_content`, or `utm_term`. Any other path is rejected with 400.
	Field string `json:"field" api:"required"`
	// Comparison applied to the field. `exists` and `not_exists` ignore `value`;
	// `regex` matches the field against `value` as a regular expression.
	//
	// Any of "contains", "equals", "exists", "not_equals", "not_exists", "regex".
	Operator string `json:"operator" api:"required"`
	// Scalar value the operator compares against. Omit for `exists` and `not_exists`.
	// Nested objects and arrays are rejected.
	Value string `json:"value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		Operator    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonalizationPropertyUpdateResponseTriggerCondition) RawJSON() string { return r.JSON.raw }
func (r *PersonalizationPropertyUpdateResponseTriggerCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonalizationPropertyListParams struct {
	// Required. List properties belonging to this experiment settings record. Get the
	// id from `GET /rest/v1/experiment-settings`.
	ExperimentSettingsID string `query:"experimentSettingsId" api:"required" json:"-"`
	// Maximum number of items to return. Defaults to 25; values below 1 are clamped to
	// 1 and values above 100 are clamped to 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Opaque pagination cursor from pagination.nextCursor in the previous response. Do
	// not decode or modify it. Malformed cursors return 400 Bad Request.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PersonalizationPropertyListParams]'s query parameters as
// `url.Values`.
func (r PersonalizationPropertyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PersonalizationPropertyNewParams struct {
	// How repeated trigger events fold into the stored value. `set_true` records
	// `true` the first time the event matches — use it for one-shot interest flags.
	// `increment` counts matching events as a number. `latest_value` stores the most
	// recent value read from `valueField`, replacing whatever was there before.
	//
	// Any of "increment", "latest_value", "set_true".
	Accumulator PersonalizationPropertyNewParamsAccumulator `json:"accumulator,omitzero" api:"required"`
	// Experiment settings record that will own this property. Get the id from
	// `GET /rest/v1/experiment-settings`.
	ExperimentSettingsID string `json:"experimentSettingsId" api:"required"`
	// Name this value is stored and read under. Must be unique within the parent
	// experiment settings record — a duplicate key is rejected with 409. Because
	// accumulated values are delivered to the visitor's browser and are readable
	// there, never accumulate secrets, credentials, PHI, or confidential data into a
	// property.
	PropertyKey string `json:"propertyKey" api:"required"`
	// Name of the tracked event that advances this property. Every event with this
	// name is evaluated against `triggerConditions`.
	TriggerEventName string `json:"triggerEventName" api:"required"`
	// Event field to read the stored value from. Required for `latest_value`; ignored
	// by `set_true` and `increment`. Only scalar values (string, number, boolean) are
	// stored. Must be one of `event.name`, `event.properties.<name>`, or
	// `event.context.<name>` where `<name>` is one of `current_url`, `referrer`,
	// `utm_source`, `utm_medium`, `utm_campaign`, `utm_content`, or `utm_term`. Any
	// other path is rejected with 400.
	ValueField param.Opt[string] `json:"valueField,omitzero"`
	// Optional filters the triggering event must satisfy. All conditions must match
	// (AND). Omit or send an empty list to accumulate on every event with the matching
	// name.
	TriggerConditions []PersonalizationPropertyNewParamsTriggerCondition `json:"triggerConditions,omitzero"`
	paramObj
}

func (r PersonalizationPropertyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PersonalizationPropertyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonalizationPropertyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How repeated trigger events fold into the stored value. `set_true` records
// `true` the first time the event matches — use it for one-shot interest flags.
// `increment` counts matching events as a number. `latest_value` stores the most
// recent value read from `valueField`, replacing whatever was there before.
type PersonalizationPropertyNewParamsAccumulator string

const (
	PersonalizationPropertyNewParamsAccumulatorIncrement   PersonalizationPropertyNewParamsAccumulator = "increment"
	PersonalizationPropertyNewParamsAccumulatorLatestValue PersonalizationPropertyNewParamsAccumulator = "latest_value"
	PersonalizationPropertyNewParamsAccumulatorSetTrue     PersonalizationPropertyNewParamsAccumulator = "set_true"
)

// The properties Field, Operator are required.
type PersonalizationPropertyNewParamsTriggerCondition struct {
	// Event field the condition reads. Must be one of `event.name`,
	// `event.properties.<name>`, or `event.context.<name>` where `<name>` is one of
	// `current_url`, `referrer`, `utm_source`, `utm_medium`, `utm_campaign`,
	// `utm_content`, or `utm_term`. Any other path is rejected with 400.
	Field string `json:"field" api:"required"`
	// Comparison applied to the field. `exists` and `not_exists` ignore `value`;
	// `regex` matches the field against `value` as a regular expression.
	//
	// Any of "contains", "equals", "exists", "not_equals", "not_exists", "regex".
	Operator string `json:"operator,omitzero" api:"required"`
	// Scalar value the operator compares against. Omit for `exists` and `not_exists`.
	// Nested objects and arrays are rejected.
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r PersonalizationPropertyNewParamsTriggerCondition) MarshalJSON() (data []byte, err error) {
	type shadow PersonalizationPropertyNewParamsTriggerCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonalizationPropertyNewParamsTriggerCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PersonalizationPropertyNewParamsTriggerCondition](
		"operator", "contains", "equals", "exists", "not_equals", "not_exists", "regex",
	)
}

type PersonalizationPropertyUpdateParams struct {
	// Renamed key. Name this value is stored and read under. Must be unique within the
	// parent experiment settings record — a duplicate key is rejected with 409.
	// Because accumulated values are delivered to the visitor's browser and are
	// readable there, never accumulate secrets, credentials, PHI, or confidential data
	// into a property. Omit this field or send `null` to leave it unchanged.
	PropertyKey param.Opt[string] `json:"propertyKey,omitzero"`
	// Name of the tracked event that advances this property. Every event with this
	// name is evaluated against `triggerConditions`. Omit this field or send `null` to
	// leave it unchanged.
	TriggerEventName param.Opt[string] `json:"triggerEventName,omitzero"`
	// Event field to read the stored value from. Required for `latest_value`; ignored
	// by `set_true` and `increment`. Only scalar values (string, number, boolean) are
	// stored. Must be one of `event.name`, `event.properties.<name>`, or
	// `event.context.<name>` where `<name>` is one of `current_url`, `referrer`,
	// `utm_source`, `utm_medium`, `utm_campaign`, `utm_content`, or `utm_term`. Any
	// other path is rejected with 400.
	ValueField param.Opt[string] `json:"valueField,omitzero"`
	// How repeated trigger events fold into the stored value. `set_true` records
	// `true` the first time the event matches — use it for one-shot interest flags.
	// `increment` counts matching events as a number. `latest_value` stores the most
	// recent value read from `valueField`, replacing whatever was there before. Omit
	// this field or send `null` to leave it unchanged.
	//
	// Any of "increment", "latest_value", "set_true".
	Accumulator PersonalizationPropertyUpdateParamsAccumulator `json:"accumulator,omitzero"`
	// Pause or resume accumulation. `Disabled` stops accumulating without discarding
	// values already recorded. Omit this field or send `null` to leave it unchanged.
	//
	// Any of "Disabled", "Enabled".
	Status PersonalizationPropertyUpdateParamsStatus `json:"status,omitzero"`
	// Optional filters the triggering event must satisfy. All conditions must match
	// (AND). Omit or send an empty list to accumulate on every event with the matching
	// name. Sending this field replaces the prior list — partial-array merging is not
	// supported. Send `null` or an empty array to clear every condition, which makes
	// the rule fire on each occurrence of the trigger event.
	TriggerConditions []PersonalizationPropertyUpdateParamsTriggerCondition `json:"triggerConditions,omitzero"`
	paramObj
}

func (r PersonalizationPropertyUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PersonalizationPropertyUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonalizationPropertyUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How repeated trigger events fold into the stored value. `set_true` records
// `true` the first time the event matches — use it for one-shot interest flags.
// `increment` counts matching events as a number. `latest_value` stores the most
// recent value read from `valueField`, replacing whatever was there before. Omit
// this field or send `null` to leave it unchanged.
type PersonalizationPropertyUpdateParamsAccumulator string

const (
	PersonalizationPropertyUpdateParamsAccumulatorIncrement   PersonalizationPropertyUpdateParamsAccumulator = "increment"
	PersonalizationPropertyUpdateParamsAccumulatorLatestValue PersonalizationPropertyUpdateParamsAccumulator = "latest_value"
	PersonalizationPropertyUpdateParamsAccumulatorSetTrue     PersonalizationPropertyUpdateParamsAccumulator = "set_true"
)

// Pause or resume accumulation. `Disabled` stops accumulating without discarding
// values already recorded. Omit this field or send `null` to leave it unchanged.
type PersonalizationPropertyUpdateParamsStatus string

const (
	PersonalizationPropertyUpdateParamsStatusDisabled PersonalizationPropertyUpdateParamsStatus = "Disabled"
	PersonalizationPropertyUpdateParamsStatusEnabled  PersonalizationPropertyUpdateParamsStatus = "Enabled"
)

// The properties Field, Operator are required.
type PersonalizationPropertyUpdateParamsTriggerCondition struct {
	// Event field the condition reads. Must be one of `event.name`,
	// `event.properties.<name>`, or `event.context.<name>` where `<name>` is one of
	// `current_url`, `referrer`, `utm_source`, `utm_medium`, `utm_campaign`,
	// `utm_content`, or `utm_term`. Any other path is rejected with 400.
	Field string `json:"field" api:"required"`
	// Comparison applied to the field. `exists` and `not_exists` ignore `value`;
	// `regex` matches the field against `value` as a regular expression.
	//
	// Any of "contains", "equals", "exists", "not_equals", "not_exists", "regex".
	Operator string `json:"operator,omitzero" api:"required"`
	// Scalar value the operator compares against. Omit for `exists` and `not_exists`.
	// Nested objects and arrays are rejected.
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r PersonalizationPropertyUpdateParamsTriggerCondition) MarshalJSON() (data []byte, err error) {
	type shadow PersonalizationPropertyUpdateParamsTriggerCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PersonalizationPropertyUpdateParamsTriggerCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PersonalizationPropertyUpdateParamsTriggerCondition](
		"operator", "contains", "equals", "exists", "not_equals", "not_exists", "regex",
	)
}
