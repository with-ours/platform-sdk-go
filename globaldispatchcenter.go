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
	"github.com/with-ours/platform-sdk-go/internal/requestconfig"
	"github.com/with-ours/platform-sdk-go/option"
	"github.com/with-ours/platform-sdk-go/packages/param"
	"github.com/with-ours/platform-sdk-go/packages/respjson"
)

// GlobalDispatchCenterService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGlobalDispatchCenterService] method instead.
type GlobalDispatchCenterService struct {
	Options []option.RequestOption
}

// NewGlobalDispatchCenterService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGlobalDispatchCenterService(opts ...option.RequestOption) (r GlobalDispatchCenterService) {
	r = GlobalDispatchCenterService{}
	r.Options = opts
	return
}

// Create a new global dispatch center. Requires scope: globalDispatch:create
func (r *GlobalDispatchCenterService) New(ctx context.Context, opts ...option.RequestOption) (res *GlobalDispatchCenterNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/global-dispatch-centers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Find a single global dispatch center by ID. Requires scope: globalDispatch:find
func (r *GlobalDispatchCenterService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *GlobalDispatchCenterGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/global-dispatch-centers/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update a global dispatch center. Requires scope: globalDispatch:update
func (r *GlobalDispatchCenterService) Update(ctx context.Context, id string, body GlobalDispatchCenterUpdateParams, opts ...option.RequestOption) (res *GlobalDispatchCenterUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/global-dispatch-centers/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List all global dispatch centers. Requires scope: globalDispatch:list
func (r *GlobalDispatchCenterService) List(ctx context.Context, opts ...option.RequestOption) (res *GlobalDispatchCenterListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/global-dispatch-centers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete a global dispatch center. Requires scope: globalDispatch:delete
func (r *GlobalDispatchCenterService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/global-dispatch-centers/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type GlobalDispatchCenterNewResponse struct {
	ID        string `json:"id" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	IsEnabled bool   `json:"isEnabled" api:"required"`
	Kind      string `json:"kind" api:"required"`
	Name      string `json:"name" api:"nullable"`
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		IsEnabled   respjson.Field
		Kind        respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalDispatchCenterNewResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalDispatchCenterNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalDispatchCenterGetResponse struct {
	ID        string `json:"id" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	IsEnabled bool   `json:"isEnabled" api:"required"`
	Kind      string `json:"kind" api:"required"`
	Name      string `json:"name" api:"nullable"`
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		IsEnabled   respjson.Field
		Kind        respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalDispatchCenterGetResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalDispatchCenterGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalDispatchCenterUpdateResponse struct {
	ID        string `json:"id" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	IsEnabled bool   `json:"isEnabled" api:"required"`
	Kind      string `json:"kind" api:"required"`
	Name      string `json:"name" api:"nullable"`
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		IsEnabled   respjson.Field
		Kind        respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalDispatchCenterUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalDispatchCenterUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalDispatchCenterListResponse struct {
	Entities []GlobalDispatchCenterListResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalDispatchCenterListResponse) RawJSON() string { return r.JSON.raw }
func (r *GlobalDispatchCenterListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalDispatchCenterListResponseEntity struct {
	ID        string `json:"id" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	IsEnabled bool   `json:"isEnabled" api:"required"`
	Kind      string `json:"kind" api:"required"`
	Name      string `json:"name" api:"nullable"`
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		IsEnabled   respjson.Field
		Kind        respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlobalDispatchCenterListResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *GlobalDispatchCenterListResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalDispatchCenterUpdateParams struct {
	IsEnabled  param.Opt[bool]                            `json:"isEnabled,omitzero"`
	Name       param.Opt[string]                          `json:"name,omitzero"`
	Notes      param.Opt[string]                          `json:"notes,omitzero"`
	Categories []GlobalDispatchCenterUpdateParamsCategory `json:"categories,omitzero"`
	paramObj
}

func (r GlobalDispatchCenterUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalDispatchCenterUpdateParamsCategory struct {
	Description    param.Opt[string]                             `json:"description,omitzero"`
	Name           param.Opt[string]                             `json:"name,omitzero"`
	Priority       param.Opt[float64]                            `json:"priority,omitzero"`
	DestinationIDs []string                                      `json:"destinationIds,omitzero"`
	Logic          GlobalDispatchCenterUpdateParamsCategoryLogic `json:"logic,omitzero"`
	paramObj
}

func (r GlobalDispatchCenterUpdateParamsCategory) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterUpdateParamsCategory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterUpdateParamsCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalDispatchCenterUpdateParamsCategoryLogic struct {
	And       []GlobalDispatchCenterUpdateParamsCategoryLogicAnd     `json:"AND,omitzero"`
	Condition GlobalDispatchCenterUpdateParamsCategoryLogicCondition `json:"condition,omitzero"`
	Not       GlobalDispatchCenterUpdateParamsCategoryLogicNot       `json:"NOT,omitzero"`
	Or        []GlobalDispatchCenterUpdateParamsCategoryLogicOr      `json:"OR,omitzero"`
	paramObj
}

func (r GlobalDispatchCenterUpdateParamsCategoryLogic) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterUpdateParamsCategoryLogic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterUpdateParamsCategoryLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalDispatchCenterUpdateParamsCategoryLogicAnd struct {
	And       []GlobalDispatchCenterUpdateParamsCategoryLogicAndAnd     `json:"AND,omitzero"`
	Condition GlobalDispatchCenterUpdateParamsCategoryLogicAndCondition `json:"condition,omitzero"`
	Not       GlobalDispatchCenterUpdateParamsCategoryLogicAndNot       `json:"NOT,omitzero"`
	Or        []GlobalDispatchCenterUpdateParamsCategoryLogicAndOr      `json:"OR,omitzero"`
	paramObj
}

func (r GlobalDispatchCenterUpdateParamsCategoryLogicAnd) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterUpdateParamsCategoryLogicAnd
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterUpdateParamsCategoryLogicAnd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalDispatchCenterUpdateParamsCategoryLogicAndAnd struct {
	Condition GlobalDispatchCenterUpdateParamsCategoryLogicAndAndCondition `json:"condition,omitzero"`
	paramObj
}

func (r GlobalDispatchCenterUpdateParamsCategoryLogicAndAnd) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterUpdateParamsCategoryLogicAndAnd
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterUpdateParamsCategoryLogicAndAnd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Operator, Property, Value are required.
type GlobalDispatchCenterUpdateParamsCategoryLogicAndAndCondition struct {
	// Any of "Contains", "DoesNotContain", "DoesNotMatchRegex",
	// "DoesNotMatchRegexIgnoreCase", "EndsWith", "Is", "IsAfter", "IsBefore",
	// "IsBetween", "IsFalse", "IsFalsy", "IsFoundIn", "IsGreaterThan",
	// "IsGreaterThanOrEqual", "IsIn", "IsLessThan", "IsLessThanOrEqual", "IsNot",
	// "IsNotFoundIn", "IsNotIn", "IsNotNull", "IsNotUndefined", "IsNull",
	// "IsOnOrAfter", "IsOnOrBefore", "IsTrue", "IsTruthy", "IsUndefined",
	// "MatchesRegex", "MatchesRegexIgnoreCase", "StartsWith".
	Operator string `json:"operator,omitzero" api:"required"`
	Property string `json:"property" api:"required"`
	Value    string `json:"value" api:"required"`
	paramObj
}

func (r GlobalDispatchCenterUpdateParamsCategoryLogicAndAndCondition) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterUpdateParamsCategoryLogicAndAndCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterUpdateParamsCategoryLogicAndAndCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GlobalDispatchCenterUpdateParamsCategoryLogicAndAndCondition](
		"operator", "Contains", "DoesNotContain", "DoesNotMatchRegex", "DoesNotMatchRegexIgnoreCase", "EndsWith", "Is", "IsAfter", "IsBefore", "IsBetween", "IsFalse", "IsFalsy", "IsFoundIn", "IsGreaterThan", "IsGreaterThanOrEqual", "IsIn", "IsLessThan", "IsLessThanOrEqual", "IsNot", "IsNotFoundIn", "IsNotIn", "IsNotNull", "IsNotUndefined", "IsNull", "IsOnOrAfter", "IsOnOrBefore", "IsTrue", "IsTruthy", "IsUndefined", "MatchesRegex", "MatchesRegexIgnoreCase", "StartsWith",
	)
}

// The properties Operator, Property, Value are required.
type GlobalDispatchCenterUpdateParamsCategoryLogicAndCondition struct {
	// Any of "Contains", "DoesNotContain", "DoesNotMatchRegex",
	// "DoesNotMatchRegexIgnoreCase", "EndsWith", "Is", "IsAfter", "IsBefore",
	// "IsBetween", "IsFalse", "IsFalsy", "IsFoundIn", "IsGreaterThan",
	// "IsGreaterThanOrEqual", "IsIn", "IsLessThan", "IsLessThanOrEqual", "IsNot",
	// "IsNotFoundIn", "IsNotIn", "IsNotNull", "IsNotUndefined", "IsNull",
	// "IsOnOrAfter", "IsOnOrBefore", "IsTrue", "IsTruthy", "IsUndefined",
	// "MatchesRegex", "MatchesRegexIgnoreCase", "StartsWith".
	Operator string `json:"operator,omitzero" api:"required"`
	Property string `json:"property" api:"required"`
	Value    string `json:"value" api:"required"`
	paramObj
}

func (r GlobalDispatchCenterUpdateParamsCategoryLogicAndCondition) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterUpdateParamsCategoryLogicAndCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterUpdateParamsCategoryLogicAndCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GlobalDispatchCenterUpdateParamsCategoryLogicAndCondition](
		"operator", "Contains", "DoesNotContain", "DoesNotMatchRegex", "DoesNotMatchRegexIgnoreCase", "EndsWith", "Is", "IsAfter", "IsBefore", "IsBetween", "IsFalse", "IsFalsy", "IsFoundIn", "IsGreaterThan", "IsGreaterThanOrEqual", "IsIn", "IsLessThan", "IsLessThanOrEqual", "IsNot", "IsNotFoundIn", "IsNotIn", "IsNotNull", "IsNotUndefined", "IsNull", "IsOnOrAfter", "IsOnOrBefore", "IsTrue", "IsTruthy", "IsUndefined", "MatchesRegex", "MatchesRegexIgnoreCase", "StartsWith",
	)
}

type GlobalDispatchCenterUpdateParamsCategoryLogicAndNot struct {
	Condition GlobalDispatchCenterUpdateParamsCategoryLogicAndNotCondition `json:"condition,omitzero"`
	paramObj
}

func (r GlobalDispatchCenterUpdateParamsCategoryLogicAndNot) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterUpdateParamsCategoryLogicAndNot
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterUpdateParamsCategoryLogicAndNot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Operator, Property, Value are required.
type GlobalDispatchCenterUpdateParamsCategoryLogicAndNotCondition struct {
	// Any of "Contains", "DoesNotContain", "DoesNotMatchRegex",
	// "DoesNotMatchRegexIgnoreCase", "EndsWith", "Is", "IsAfter", "IsBefore",
	// "IsBetween", "IsFalse", "IsFalsy", "IsFoundIn", "IsGreaterThan",
	// "IsGreaterThanOrEqual", "IsIn", "IsLessThan", "IsLessThanOrEqual", "IsNot",
	// "IsNotFoundIn", "IsNotIn", "IsNotNull", "IsNotUndefined", "IsNull",
	// "IsOnOrAfter", "IsOnOrBefore", "IsTrue", "IsTruthy", "IsUndefined",
	// "MatchesRegex", "MatchesRegexIgnoreCase", "StartsWith".
	Operator string `json:"operator,omitzero" api:"required"`
	Property string `json:"property" api:"required"`
	Value    string `json:"value" api:"required"`
	paramObj
}

func (r GlobalDispatchCenterUpdateParamsCategoryLogicAndNotCondition) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterUpdateParamsCategoryLogicAndNotCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterUpdateParamsCategoryLogicAndNotCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GlobalDispatchCenterUpdateParamsCategoryLogicAndNotCondition](
		"operator", "Contains", "DoesNotContain", "DoesNotMatchRegex", "DoesNotMatchRegexIgnoreCase", "EndsWith", "Is", "IsAfter", "IsBefore", "IsBetween", "IsFalse", "IsFalsy", "IsFoundIn", "IsGreaterThan", "IsGreaterThanOrEqual", "IsIn", "IsLessThan", "IsLessThanOrEqual", "IsNot", "IsNotFoundIn", "IsNotIn", "IsNotNull", "IsNotUndefined", "IsNull", "IsOnOrAfter", "IsOnOrBefore", "IsTrue", "IsTruthy", "IsUndefined", "MatchesRegex", "MatchesRegexIgnoreCase", "StartsWith",
	)
}

type GlobalDispatchCenterUpdateParamsCategoryLogicAndOr struct {
	Condition GlobalDispatchCenterUpdateParamsCategoryLogicAndOrCondition `json:"condition,omitzero"`
	paramObj
}

func (r GlobalDispatchCenterUpdateParamsCategoryLogicAndOr) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterUpdateParamsCategoryLogicAndOr
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterUpdateParamsCategoryLogicAndOr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Operator, Property, Value are required.
type GlobalDispatchCenterUpdateParamsCategoryLogicAndOrCondition struct {
	// Any of "Contains", "DoesNotContain", "DoesNotMatchRegex",
	// "DoesNotMatchRegexIgnoreCase", "EndsWith", "Is", "IsAfter", "IsBefore",
	// "IsBetween", "IsFalse", "IsFalsy", "IsFoundIn", "IsGreaterThan",
	// "IsGreaterThanOrEqual", "IsIn", "IsLessThan", "IsLessThanOrEqual", "IsNot",
	// "IsNotFoundIn", "IsNotIn", "IsNotNull", "IsNotUndefined", "IsNull",
	// "IsOnOrAfter", "IsOnOrBefore", "IsTrue", "IsTruthy", "IsUndefined",
	// "MatchesRegex", "MatchesRegexIgnoreCase", "StartsWith".
	Operator string `json:"operator,omitzero" api:"required"`
	Property string `json:"property" api:"required"`
	Value    string `json:"value" api:"required"`
	paramObj
}

func (r GlobalDispatchCenterUpdateParamsCategoryLogicAndOrCondition) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterUpdateParamsCategoryLogicAndOrCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterUpdateParamsCategoryLogicAndOrCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GlobalDispatchCenterUpdateParamsCategoryLogicAndOrCondition](
		"operator", "Contains", "DoesNotContain", "DoesNotMatchRegex", "DoesNotMatchRegexIgnoreCase", "EndsWith", "Is", "IsAfter", "IsBefore", "IsBetween", "IsFalse", "IsFalsy", "IsFoundIn", "IsGreaterThan", "IsGreaterThanOrEqual", "IsIn", "IsLessThan", "IsLessThanOrEqual", "IsNot", "IsNotFoundIn", "IsNotIn", "IsNotNull", "IsNotUndefined", "IsNull", "IsOnOrAfter", "IsOnOrBefore", "IsTrue", "IsTruthy", "IsUndefined", "MatchesRegex", "MatchesRegexIgnoreCase", "StartsWith",
	)
}

// The properties Operator, Property, Value are required.
type GlobalDispatchCenterUpdateParamsCategoryLogicCondition struct {
	// Any of "Contains", "DoesNotContain", "DoesNotMatchRegex",
	// "DoesNotMatchRegexIgnoreCase", "EndsWith", "Is", "IsAfter", "IsBefore",
	// "IsBetween", "IsFalse", "IsFalsy", "IsFoundIn", "IsGreaterThan",
	// "IsGreaterThanOrEqual", "IsIn", "IsLessThan", "IsLessThanOrEqual", "IsNot",
	// "IsNotFoundIn", "IsNotIn", "IsNotNull", "IsNotUndefined", "IsNull",
	// "IsOnOrAfter", "IsOnOrBefore", "IsTrue", "IsTruthy", "IsUndefined",
	// "MatchesRegex", "MatchesRegexIgnoreCase", "StartsWith".
	Operator string `json:"operator,omitzero" api:"required"`
	Property string `json:"property" api:"required"`
	Value    string `json:"value" api:"required"`
	paramObj
}

func (r GlobalDispatchCenterUpdateParamsCategoryLogicCondition) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterUpdateParamsCategoryLogicCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterUpdateParamsCategoryLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GlobalDispatchCenterUpdateParamsCategoryLogicCondition](
		"operator", "Contains", "DoesNotContain", "DoesNotMatchRegex", "DoesNotMatchRegexIgnoreCase", "EndsWith", "Is", "IsAfter", "IsBefore", "IsBetween", "IsFalse", "IsFalsy", "IsFoundIn", "IsGreaterThan", "IsGreaterThanOrEqual", "IsIn", "IsLessThan", "IsLessThanOrEqual", "IsNot", "IsNotFoundIn", "IsNotIn", "IsNotNull", "IsNotUndefined", "IsNull", "IsOnOrAfter", "IsOnOrBefore", "IsTrue", "IsTruthy", "IsUndefined", "MatchesRegex", "MatchesRegexIgnoreCase", "StartsWith",
	)
}

type GlobalDispatchCenterUpdateParamsCategoryLogicNot struct {
	And       []GlobalDispatchCenterUpdateParamsCategoryLogicNotAnd     `json:"AND,omitzero"`
	Condition GlobalDispatchCenterUpdateParamsCategoryLogicNotCondition `json:"condition,omitzero"`
	Not       GlobalDispatchCenterUpdateParamsCategoryLogicNotNot       `json:"NOT,omitzero"`
	Or        []GlobalDispatchCenterUpdateParamsCategoryLogicNotOr      `json:"OR,omitzero"`
	paramObj
}

func (r GlobalDispatchCenterUpdateParamsCategoryLogicNot) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterUpdateParamsCategoryLogicNot
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterUpdateParamsCategoryLogicNot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalDispatchCenterUpdateParamsCategoryLogicNotAnd struct {
	Condition GlobalDispatchCenterUpdateParamsCategoryLogicNotAndCondition `json:"condition,omitzero"`
	paramObj
}

func (r GlobalDispatchCenterUpdateParamsCategoryLogicNotAnd) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterUpdateParamsCategoryLogicNotAnd
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterUpdateParamsCategoryLogicNotAnd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Operator, Property, Value are required.
type GlobalDispatchCenterUpdateParamsCategoryLogicNotAndCondition struct {
	// Any of "Contains", "DoesNotContain", "DoesNotMatchRegex",
	// "DoesNotMatchRegexIgnoreCase", "EndsWith", "Is", "IsAfter", "IsBefore",
	// "IsBetween", "IsFalse", "IsFalsy", "IsFoundIn", "IsGreaterThan",
	// "IsGreaterThanOrEqual", "IsIn", "IsLessThan", "IsLessThanOrEqual", "IsNot",
	// "IsNotFoundIn", "IsNotIn", "IsNotNull", "IsNotUndefined", "IsNull",
	// "IsOnOrAfter", "IsOnOrBefore", "IsTrue", "IsTruthy", "IsUndefined",
	// "MatchesRegex", "MatchesRegexIgnoreCase", "StartsWith".
	Operator string `json:"operator,omitzero" api:"required"`
	Property string `json:"property" api:"required"`
	Value    string `json:"value" api:"required"`
	paramObj
}

func (r GlobalDispatchCenterUpdateParamsCategoryLogicNotAndCondition) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterUpdateParamsCategoryLogicNotAndCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterUpdateParamsCategoryLogicNotAndCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GlobalDispatchCenterUpdateParamsCategoryLogicNotAndCondition](
		"operator", "Contains", "DoesNotContain", "DoesNotMatchRegex", "DoesNotMatchRegexIgnoreCase", "EndsWith", "Is", "IsAfter", "IsBefore", "IsBetween", "IsFalse", "IsFalsy", "IsFoundIn", "IsGreaterThan", "IsGreaterThanOrEqual", "IsIn", "IsLessThan", "IsLessThanOrEqual", "IsNot", "IsNotFoundIn", "IsNotIn", "IsNotNull", "IsNotUndefined", "IsNull", "IsOnOrAfter", "IsOnOrBefore", "IsTrue", "IsTruthy", "IsUndefined", "MatchesRegex", "MatchesRegexIgnoreCase", "StartsWith",
	)
}

// The properties Operator, Property, Value are required.
type GlobalDispatchCenterUpdateParamsCategoryLogicNotCondition struct {
	// Any of "Contains", "DoesNotContain", "DoesNotMatchRegex",
	// "DoesNotMatchRegexIgnoreCase", "EndsWith", "Is", "IsAfter", "IsBefore",
	// "IsBetween", "IsFalse", "IsFalsy", "IsFoundIn", "IsGreaterThan",
	// "IsGreaterThanOrEqual", "IsIn", "IsLessThan", "IsLessThanOrEqual", "IsNot",
	// "IsNotFoundIn", "IsNotIn", "IsNotNull", "IsNotUndefined", "IsNull",
	// "IsOnOrAfter", "IsOnOrBefore", "IsTrue", "IsTruthy", "IsUndefined",
	// "MatchesRegex", "MatchesRegexIgnoreCase", "StartsWith".
	Operator string `json:"operator,omitzero" api:"required"`
	Property string `json:"property" api:"required"`
	Value    string `json:"value" api:"required"`
	paramObj
}

func (r GlobalDispatchCenterUpdateParamsCategoryLogicNotCondition) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterUpdateParamsCategoryLogicNotCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterUpdateParamsCategoryLogicNotCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GlobalDispatchCenterUpdateParamsCategoryLogicNotCondition](
		"operator", "Contains", "DoesNotContain", "DoesNotMatchRegex", "DoesNotMatchRegexIgnoreCase", "EndsWith", "Is", "IsAfter", "IsBefore", "IsBetween", "IsFalse", "IsFalsy", "IsFoundIn", "IsGreaterThan", "IsGreaterThanOrEqual", "IsIn", "IsLessThan", "IsLessThanOrEqual", "IsNot", "IsNotFoundIn", "IsNotIn", "IsNotNull", "IsNotUndefined", "IsNull", "IsOnOrAfter", "IsOnOrBefore", "IsTrue", "IsTruthy", "IsUndefined", "MatchesRegex", "MatchesRegexIgnoreCase", "StartsWith",
	)
}

type GlobalDispatchCenterUpdateParamsCategoryLogicNotNot struct {
	Condition GlobalDispatchCenterUpdateParamsCategoryLogicNotNotCondition `json:"condition,omitzero"`
	paramObj
}

func (r GlobalDispatchCenterUpdateParamsCategoryLogicNotNot) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterUpdateParamsCategoryLogicNotNot
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterUpdateParamsCategoryLogicNotNot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Operator, Property, Value are required.
type GlobalDispatchCenterUpdateParamsCategoryLogicNotNotCondition struct {
	// Any of "Contains", "DoesNotContain", "DoesNotMatchRegex",
	// "DoesNotMatchRegexIgnoreCase", "EndsWith", "Is", "IsAfter", "IsBefore",
	// "IsBetween", "IsFalse", "IsFalsy", "IsFoundIn", "IsGreaterThan",
	// "IsGreaterThanOrEqual", "IsIn", "IsLessThan", "IsLessThanOrEqual", "IsNot",
	// "IsNotFoundIn", "IsNotIn", "IsNotNull", "IsNotUndefined", "IsNull",
	// "IsOnOrAfter", "IsOnOrBefore", "IsTrue", "IsTruthy", "IsUndefined",
	// "MatchesRegex", "MatchesRegexIgnoreCase", "StartsWith".
	Operator string `json:"operator,omitzero" api:"required"`
	Property string `json:"property" api:"required"`
	Value    string `json:"value" api:"required"`
	paramObj
}

func (r GlobalDispatchCenterUpdateParamsCategoryLogicNotNotCondition) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterUpdateParamsCategoryLogicNotNotCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterUpdateParamsCategoryLogicNotNotCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GlobalDispatchCenterUpdateParamsCategoryLogicNotNotCondition](
		"operator", "Contains", "DoesNotContain", "DoesNotMatchRegex", "DoesNotMatchRegexIgnoreCase", "EndsWith", "Is", "IsAfter", "IsBefore", "IsBetween", "IsFalse", "IsFalsy", "IsFoundIn", "IsGreaterThan", "IsGreaterThanOrEqual", "IsIn", "IsLessThan", "IsLessThanOrEqual", "IsNot", "IsNotFoundIn", "IsNotIn", "IsNotNull", "IsNotUndefined", "IsNull", "IsOnOrAfter", "IsOnOrBefore", "IsTrue", "IsTruthy", "IsUndefined", "MatchesRegex", "MatchesRegexIgnoreCase", "StartsWith",
	)
}

type GlobalDispatchCenterUpdateParamsCategoryLogicNotOr struct {
	Condition GlobalDispatchCenterUpdateParamsCategoryLogicNotOrCondition `json:"condition,omitzero"`
	paramObj
}

func (r GlobalDispatchCenterUpdateParamsCategoryLogicNotOr) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterUpdateParamsCategoryLogicNotOr
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterUpdateParamsCategoryLogicNotOr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Operator, Property, Value are required.
type GlobalDispatchCenterUpdateParamsCategoryLogicNotOrCondition struct {
	// Any of "Contains", "DoesNotContain", "DoesNotMatchRegex",
	// "DoesNotMatchRegexIgnoreCase", "EndsWith", "Is", "IsAfter", "IsBefore",
	// "IsBetween", "IsFalse", "IsFalsy", "IsFoundIn", "IsGreaterThan",
	// "IsGreaterThanOrEqual", "IsIn", "IsLessThan", "IsLessThanOrEqual", "IsNot",
	// "IsNotFoundIn", "IsNotIn", "IsNotNull", "IsNotUndefined", "IsNull",
	// "IsOnOrAfter", "IsOnOrBefore", "IsTrue", "IsTruthy", "IsUndefined",
	// "MatchesRegex", "MatchesRegexIgnoreCase", "StartsWith".
	Operator string `json:"operator,omitzero" api:"required"`
	Property string `json:"property" api:"required"`
	Value    string `json:"value" api:"required"`
	paramObj
}

func (r GlobalDispatchCenterUpdateParamsCategoryLogicNotOrCondition) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterUpdateParamsCategoryLogicNotOrCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterUpdateParamsCategoryLogicNotOrCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GlobalDispatchCenterUpdateParamsCategoryLogicNotOrCondition](
		"operator", "Contains", "DoesNotContain", "DoesNotMatchRegex", "DoesNotMatchRegexIgnoreCase", "EndsWith", "Is", "IsAfter", "IsBefore", "IsBetween", "IsFalse", "IsFalsy", "IsFoundIn", "IsGreaterThan", "IsGreaterThanOrEqual", "IsIn", "IsLessThan", "IsLessThanOrEqual", "IsNot", "IsNotFoundIn", "IsNotIn", "IsNotNull", "IsNotUndefined", "IsNull", "IsOnOrAfter", "IsOnOrBefore", "IsTrue", "IsTruthy", "IsUndefined", "MatchesRegex", "MatchesRegexIgnoreCase", "StartsWith",
	)
}

type GlobalDispatchCenterUpdateParamsCategoryLogicOr struct {
	And       []GlobalDispatchCenterUpdateParamsCategoryLogicOrAnd     `json:"AND,omitzero"`
	Condition GlobalDispatchCenterUpdateParamsCategoryLogicOrCondition `json:"condition,omitzero"`
	Not       GlobalDispatchCenterUpdateParamsCategoryLogicOrNot       `json:"NOT,omitzero"`
	Or        []GlobalDispatchCenterUpdateParamsCategoryLogicOrOr      `json:"OR,omitzero"`
	paramObj
}

func (r GlobalDispatchCenterUpdateParamsCategoryLogicOr) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterUpdateParamsCategoryLogicOr
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterUpdateParamsCategoryLogicOr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GlobalDispatchCenterUpdateParamsCategoryLogicOrAnd struct {
	Condition GlobalDispatchCenterUpdateParamsCategoryLogicOrAndCondition `json:"condition,omitzero"`
	paramObj
}

func (r GlobalDispatchCenterUpdateParamsCategoryLogicOrAnd) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterUpdateParamsCategoryLogicOrAnd
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterUpdateParamsCategoryLogicOrAnd) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Operator, Property, Value are required.
type GlobalDispatchCenterUpdateParamsCategoryLogicOrAndCondition struct {
	// Any of "Contains", "DoesNotContain", "DoesNotMatchRegex",
	// "DoesNotMatchRegexIgnoreCase", "EndsWith", "Is", "IsAfter", "IsBefore",
	// "IsBetween", "IsFalse", "IsFalsy", "IsFoundIn", "IsGreaterThan",
	// "IsGreaterThanOrEqual", "IsIn", "IsLessThan", "IsLessThanOrEqual", "IsNot",
	// "IsNotFoundIn", "IsNotIn", "IsNotNull", "IsNotUndefined", "IsNull",
	// "IsOnOrAfter", "IsOnOrBefore", "IsTrue", "IsTruthy", "IsUndefined",
	// "MatchesRegex", "MatchesRegexIgnoreCase", "StartsWith".
	Operator string `json:"operator,omitzero" api:"required"`
	Property string `json:"property" api:"required"`
	Value    string `json:"value" api:"required"`
	paramObj
}

func (r GlobalDispatchCenterUpdateParamsCategoryLogicOrAndCondition) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterUpdateParamsCategoryLogicOrAndCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterUpdateParamsCategoryLogicOrAndCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GlobalDispatchCenterUpdateParamsCategoryLogicOrAndCondition](
		"operator", "Contains", "DoesNotContain", "DoesNotMatchRegex", "DoesNotMatchRegexIgnoreCase", "EndsWith", "Is", "IsAfter", "IsBefore", "IsBetween", "IsFalse", "IsFalsy", "IsFoundIn", "IsGreaterThan", "IsGreaterThanOrEqual", "IsIn", "IsLessThan", "IsLessThanOrEqual", "IsNot", "IsNotFoundIn", "IsNotIn", "IsNotNull", "IsNotUndefined", "IsNull", "IsOnOrAfter", "IsOnOrBefore", "IsTrue", "IsTruthy", "IsUndefined", "MatchesRegex", "MatchesRegexIgnoreCase", "StartsWith",
	)
}

// The properties Operator, Property, Value are required.
type GlobalDispatchCenterUpdateParamsCategoryLogicOrCondition struct {
	// Any of "Contains", "DoesNotContain", "DoesNotMatchRegex",
	// "DoesNotMatchRegexIgnoreCase", "EndsWith", "Is", "IsAfter", "IsBefore",
	// "IsBetween", "IsFalse", "IsFalsy", "IsFoundIn", "IsGreaterThan",
	// "IsGreaterThanOrEqual", "IsIn", "IsLessThan", "IsLessThanOrEqual", "IsNot",
	// "IsNotFoundIn", "IsNotIn", "IsNotNull", "IsNotUndefined", "IsNull",
	// "IsOnOrAfter", "IsOnOrBefore", "IsTrue", "IsTruthy", "IsUndefined",
	// "MatchesRegex", "MatchesRegexIgnoreCase", "StartsWith".
	Operator string `json:"operator,omitzero" api:"required"`
	Property string `json:"property" api:"required"`
	Value    string `json:"value" api:"required"`
	paramObj
}

func (r GlobalDispatchCenterUpdateParamsCategoryLogicOrCondition) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterUpdateParamsCategoryLogicOrCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterUpdateParamsCategoryLogicOrCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GlobalDispatchCenterUpdateParamsCategoryLogicOrCondition](
		"operator", "Contains", "DoesNotContain", "DoesNotMatchRegex", "DoesNotMatchRegexIgnoreCase", "EndsWith", "Is", "IsAfter", "IsBefore", "IsBetween", "IsFalse", "IsFalsy", "IsFoundIn", "IsGreaterThan", "IsGreaterThanOrEqual", "IsIn", "IsLessThan", "IsLessThanOrEqual", "IsNot", "IsNotFoundIn", "IsNotIn", "IsNotNull", "IsNotUndefined", "IsNull", "IsOnOrAfter", "IsOnOrBefore", "IsTrue", "IsTruthy", "IsUndefined", "MatchesRegex", "MatchesRegexIgnoreCase", "StartsWith",
	)
}

type GlobalDispatchCenterUpdateParamsCategoryLogicOrNot struct {
	Condition GlobalDispatchCenterUpdateParamsCategoryLogicOrNotCondition `json:"condition,omitzero"`
	paramObj
}

func (r GlobalDispatchCenterUpdateParamsCategoryLogicOrNot) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterUpdateParamsCategoryLogicOrNot
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterUpdateParamsCategoryLogicOrNot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Operator, Property, Value are required.
type GlobalDispatchCenterUpdateParamsCategoryLogicOrNotCondition struct {
	// Any of "Contains", "DoesNotContain", "DoesNotMatchRegex",
	// "DoesNotMatchRegexIgnoreCase", "EndsWith", "Is", "IsAfter", "IsBefore",
	// "IsBetween", "IsFalse", "IsFalsy", "IsFoundIn", "IsGreaterThan",
	// "IsGreaterThanOrEqual", "IsIn", "IsLessThan", "IsLessThanOrEqual", "IsNot",
	// "IsNotFoundIn", "IsNotIn", "IsNotNull", "IsNotUndefined", "IsNull",
	// "IsOnOrAfter", "IsOnOrBefore", "IsTrue", "IsTruthy", "IsUndefined",
	// "MatchesRegex", "MatchesRegexIgnoreCase", "StartsWith".
	Operator string `json:"operator,omitzero" api:"required"`
	Property string `json:"property" api:"required"`
	Value    string `json:"value" api:"required"`
	paramObj
}

func (r GlobalDispatchCenterUpdateParamsCategoryLogicOrNotCondition) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterUpdateParamsCategoryLogicOrNotCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterUpdateParamsCategoryLogicOrNotCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GlobalDispatchCenterUpdateParamsCategoryLogicOrNotCondition](
		"operator", "Contains", "DoesNotContain", "DoesNotMatchRegex", "DoesNotMatchRegexIgnoreCase", "EndsWith", "Is", "IsAfter", "IsBefore", "IsBetween", "IsFalse", "IsFalsy", "IsFoundIn", "IsGreaterThan", "IsGreaterThanOrEqual", "IsIn", "IsLessThan", "IsLessThanOrEqual", "IsNot", "IsNotFoundIn", "IsNotIn", "IsNotNull", "IsNotUndefined", "IsNull", "IsOnOrAfter", "IsOnOrBefore", "IsTrue", "IsTruthy", "IsUndefined", "MatchesRegex", "MatchesRegexIgnoreCase", "StartsWith",
	)
}

type GlobalDispatchCenterUpdateParamsCategoryLogicOrOr struct {
	Condition GlobalDispatchCenterUpdateParamsCategoryLogicOrOrCondition `json:"condition,omitzero"`
	paramObj
}

func (r GlobalDispatchCenterUpdateParamsCategoryLogicOrOr) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterUpdateParamsCategoryLogicOrOr
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterUpdateParamsCategoryLogicOrOr) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Operator, Property, Value are required.
type GlobalDispatchCenterUpdateParamsCategoryLogicOrOrCondition struct {
	// Any of "Contains", "DoesNotContain", "DoesNotMatchRegex",
	// "DoesNotMatchRegexIgnoreCase", "EndsWith", "Is", "IsAfter", "IsBefore",
	// "IsBetween", "IsFalse", "IsFalsy", "IsFoundIn", "IsGreaterThan",
	// "IsGreaterThanOrEqual", "IsIn", "IsLessThan", "IsLessThanOrEqual", "IsNot",
	// "IsNotFoundIn", "IsNotIn", "IsNotNull", "IsNotUndefined", "IsNull",
	// "IsOnOrAfter", "IsOnOrBefore", "IsTrue", "IsTruthy", "IsUndefined",
	// "MatchesRegex", "MatchesRegexIgnoreCase", "StartsWith".
	Operator string `json:"operator,omitzero" api:"required"`
	Property string `json:"property" api:"required"`
	Value    string `json:"value" api:"required"`
	paramObj
}

func (r GlobalDispatchCenterUpdateParamsCategoryLogicOrOrCondition) MarshalJSON() (data []byte, err error) {
	type shadow GlobalDispatchCenterUpdateParamsCategoryLogicOrOrCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GlobalDispatchCenterUpdateParamsCategoryLogicOrOrCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[GlobalDispatchCenterUpdateParamsCategoryLogicOrOrCondition](
		"operator", "Contains", "DoesNotContain", "DoesNotMatchRegex", "DoesNotMatchRegexIgnoreCase", "EndsWith", "Is", "IsAfter", "IsBefore", "IsBetween", "IsFalse", "IsFalsy", "IsFoundIn", "IsGreaterThan", "IsGreaterThanOrEqual", "IsIn", "IsLessThan", "IsLessThanOrEqual", "IsNot", "IsNotFoundIn", "IsNotIn", "IsNotNull", "IsNotUndefined", "IsNull", "IsOnOrAfter", "IsOnOrBefore", "IsTrue", "IsTruthy", "IsUndefined", "MatchesRegex", "MatchesRegexIgnoreCase", "StartsWith",
	)
}
