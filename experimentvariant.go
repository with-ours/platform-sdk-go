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

// ExperimentVariantService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExperimentVariantService] method instead.
type ExperimentVariantService struct {
	Options []option.RequestOption
}

// NewExperimentVariantService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewExperimentVariantService(opts ...option.RequestOption) (r ExperimentVariantService) {
	r = ExperimentVariantService{}
	r.Options = opts
	return
}

// List variants for a specific parent experiment. Requires the `experimentId`
// query parameter — variants are always scoped to a single experiment. Supports
// cursor pagination via `limit` and `cursor`; SDK runtimes that need the full set
// in one request can pass `?limit=100`. Requires scope: experiment:find
func (r *ExperimentVariantService) List(ctx context.Context, query ExperimentVariantListParams, opts ...option.RequestOption) (res *pagination.Cursor[ExperimentVariantListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "rest/v1/experiment-variants"
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

// List variants for a specific parent experiment. Requires the `experimentId`
// query parameter — variants are always scoped to a single experiment. Supports
// cursor pagination via `limit` and `cursor`; SDK runtimes that need the full set
// in one request can pass `?limit=100`. Requires scope: experiment:find
func (r *ExperimentVariantService) ListAutoPaging(ctx context.Context, query ExperimentVariantListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[ExperimentVariantListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Create a new experiment variant. Requires scope: experiment:update
func (r *ExperimentVariantService) New(ctx context.Context, body ExperimentVariantNewParams, opts ...option.RequestOption) (res *ExperimentVariantNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/experiment-variants"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Find a single experiment variant by ID. Requires scope: experiment:find
func (r *ExperimentVariantService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ExperimentVariantGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/experiment-variants/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Partially update an experiment variant. Only the fields you send are changed.
// Requires scope: experiment:update
func (r *ExperimentVariantService) Update(ctx context.Context, id string, body ExperimentVariantUpdateParams, opts ...option.RequestOption) (res *ExperimentVariantUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/experiment-variants/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete an experiment variant. Requires scope: experiment:update
func (r *ExperimentVariantService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *ExperimentVariantDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/experiment-variants/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type ExperimentVariantListResponse struct {
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
	DomModifications []ExperimentVariantListResponseDomModification `json:"domModifications" api:"nullable"`
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
func (r ExperimentVariantListResponse) RawJSON() string { return r.JSON.raw }
func (r *ExperimentVariantListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentVariantListResponseDomModification struct {
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
	Styles []ExperimentVariantListResponseDomModificationStyle `json:"styles" api:"nullable"`
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
func (r ExperimentVariantListResponseDomModification) RawJSON() string { return r.JSON.raw }
func (r *ExperimentVariantListResponseDomModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentVariantListResponseDomModificationStyle struct {
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
func (r ExperimentVariantListResponseDomModificationStyle) RawJSON() string { return r.JSON.raw }
func (r *ExperimentVariantListResponseDomModificationStyle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentVariantNewResponse struct {
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
	DomModifications []ExperimentVariantNewResponseDomModification `json:"domModifications" api:"nullable"`
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
func (r ExperimentVariantNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ExperimentVariantNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentVariantNewResponseDomModification struct {
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
	Styles []ExperimentVariantNewResponseDomModificationStyle `json:"styles" api:"nullable"`
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
func (r ExperimentVariantNewResponseDomModification) RawJSON() string { return r.JSON.raw }
func (r *ExperimentVariantNewResponseDomModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentVariantNewResponseDomModificationStyle struct {
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
func (r ExperimentVariantNewResponseDomModificationStyle) RawJSON() string { return r.JSON.raw }
func (r *ExperimentVariantNewResponseDomModificationStyle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentVariantGetResponse struct {
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
	DomModifications []ExperimentVariantGetResponseDomModification `json:"domModifications" api:"nullable"`
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
func (r ExperimentVariantGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ExperimentVariantGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentVariantGetResponseDomModification struct {
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
	Styles []ExperimentVariantGetResponseDomModificationStyle `json:"styles" api:"nullable"`
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
func (r ExperimentVariantGetResponseDomModification) RawJSON() string { return r.JSON.raw }
func (r *ExperimentVariantGetResponseDomModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentVariantGetResponseDomModificationStyle struct {
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
func (r ExperimentVariantGetResponseDomModificationStyle) RawJSON() string { return r.JSON.raw }
func (r *ExperimentVariantGetResponseDomModificationStyle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentVariantUpdateResponse struct {
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
	DomModifications []ExperimentVariantUpdateResponseDomModification `json:"domModifications" api:"nullable"`
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
func (r ExperimentVariantUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ExperimentVariantUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentVariantUpdateResponseDomModification struct {
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
	Styles []ExperimentVariantUpdateResponseDomModificationStyle `json:"styles" api:"nullable"`
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
func (r ExperimentVariantUpdateResponseDomModification) RawJSON() string { return r.JSON.raw }
func (r *ExperimentVariantUpdateResponseDomModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentVariantUpdateResponseDomModificationStyle struct {
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
func (r ExperimentVariantUpdateResponseDomModificationStyle) RawJSON() string { return r.JSON.raw }
func (r *ExperimentVariantUpdateResponseDomModificationStyle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentVariantDeleteResponse struct {
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
	DomModifications []ExperimentVariantDeleteResponseDomModification `json:"domModifications" api:"nullable"`
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
func (r ExperimentVariantDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ExperimentVariantDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentVariantDeleteResponseDomModification struct {
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
	Styles []ExperimentVariantDeleteResponseDomModificationStyle `json:"styles" api:"nullable"`
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
func (r ExperimentVariantDeleteResponseDomModification) RawJSON() string { return r.JSON.raw }
func (r *ExperimentVariantDeleteResponseDomModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentVariantDeleteResponseDomModificationStyle struct {
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
func (r ExperimentVariantDeleteResponseDomModificationStyle) RawJSON() string { return r.JSON.raw }
func (r *ExperimentVariantDeleteResponseDomModificationStyle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentVariantListParams struct {
	// Required. List variants belonging to this parent experiment.
	ExperimentID string `query:"experimentId" api:"required" json:"-"`
	// Maximum number of variants to return. Defaults to 200; values below 1 are
	// clamped to 1 and values above 200 are clamped to 200. Variants per experiment
	// are capped at 200 server-side, so a single request returns the full set.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Opaque pagination cursor from pagination.nextCursor in the previous response. Do
	// not decode or modify it. Malformed cursors return 400 Bad Request.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExperimentVariantListParams]'s query parameters as
// `url.Values`.
func (r ExperimentVariantListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExperimentVariantNewParams struct {
	// Parent experiment ID that will own this new variant.
	ExperimentID string `json:"experimentId" api:"required"`
	// Human-readable name for the new variant.
	Name string `json:"name" api:"required"`
	// Traffic weight to assign to this variant. Weights are relative shares; the
	// runtime normalizes by their sum. Must be a positive integer in the range
	// 1..1_000_000.
	Weight int64 `json:"weight" api:"required"`
	// Mark this variant as the experiment control. Defaults to `false`. The API
	// rejects the request with 409 if the experiment already has a control variant —
	// to swap controls, first PATCH the existing control to clear `isControl`, then
	// create or PATCH the new one with `isControl: true`. The auto-generated control
	// variant created with each new experiment can be replaced this way. DELETE on the
	// control returns 409.
	IsControl param.Opt[bool] `json:"isControl,omitzero"`
	// Required for redirect variants. Use either a site-relative path such as
	// `/pricing-v2` or an absolute `https://` URL. Cross-origin `http://` URLs are
	// rejected. Omit for DOM modification variants.
	RedirectURL param.Opt[string] `json:"redirectUrl,omitzero"`
	// Required for DOM modification variants. Omit for redirect variants. Each entry
	// is `{selector, action, value}`.
	DomModifications []ExperimentVariantNewParamsDomModification `json:"domModifications,omitzero"`
	// Variant delivery mechanism. `dom_modifications` mutates the current page
	// in-place at SDK runtime — use it for copy/style/image/HTML changes that keep
	// visitors on the same URL (headline copy tests, button color, hero image swap).
	// `redirect` routes the visitor to a different URL entirely — use it for
	// landing-page A/B tests, alternate pricing pages, or any test where the _page
	// itself_ is the variable. They are not interchangeable: a redirect variant cannot
	// also tweak DOM, and a dom_modifications variant cannot send the visitor
	// elsewhere.
	//
	// Any of "dom_modifications", "redirect".
	VariantType ExperimentVariantNewParamsVariantType `json:"variantType,omitzero"`
	paramObj
}

func (r ExperimentVariantNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentVariantNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentVariantNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Action, Selector are required.
type ExperimentVariantNewParamsDomModification struct {
	// Mutation to apply when the selector matches.
	//
	// Any of "customCss", "customJs", "insertAfter", "insertBefore", "remove",
	// "setAttribute", "setHtml", "setImage", "setStyle", "setText".
	Action string `json:"action,omitzero" api:"required"`
	// CSS selector for the element to modify at runtime. PREFER specific selectors
	// that match exactly one element: an `id` (`#hero-headline`), a stable `data-*`
	// attribute (`[data-testid="hero-headline"]`), or a unique class/structural chain
	// (`section.hero > h1.headline`). AVOID bare tag selectors like `h1`, `button`, or
	// `img` — modern pages usually contain several, and the runtime applies the
	// mutation to ONLY THE FIRST match, which silently picks the wrong element. If you
	// only have a tag name, scope it with the nearest unique ancestor (e.g. `main h1`,
	// `header nav a:first-of-type`).
	Selector string `json:"selector" api:"required"`
	// Canonical action payload. For `setText` / `setHtml` / `customCss` / `customJs` /
	// `setImage` / `insertBefore` / `insertAfter` this is the literal
	// text/HTML/CSS/JS/URL. For `setStyle` and `setAttribute` it is a JSON-stringified
	// `{key: value}` object — or you can supply the structured `styles` / `attribute`
	// field instead and the server will normalize.
	Value param.Opt[string] `json:"value,omitzero"`
	// Use this for `setAttribute` to avoid JSON-stringifying `{name: value}` yourself.
	// Ignored for other actions.
	Attribute any `json:"attribute,omitzero"`
	// Use this for `setStyle` to avoid JSON-stringifying `{property: value}` yourself.
	// Ignored for other actions.
	Styles []ExperimentVariantNewParamsDomModificationStyle `json:"styles,omitzero"`
	paramObj
}

func (r ExperimentVariantNewParamsDomModification) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentVariantNewParamsDomModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentVariantNewParamsDomModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExperimentVariantNewParamsDomModification](
		"action", "customCss", "customJs", "insertAfter", "insertBefore", "remove", "setAttribute", "setHtml", "setImage", "setStyle", "setText",
	)
}

// The properties Property, Value are required.
type ExperimentVariantNewParamsDomModificationStyle struct {
	// CSS property name in camelCase or kebab-case.
	Property string `json:"property" api:"required"`
	// CSS value to assign to the property.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r ExperimentVariantNewParamsDomModificationStyle) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentVariantNewParamsDomModificationStyle
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentVariantNewParamsDomModificationStyle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Variant delivery mechanism. `dom_modifications` mutates the current page
// in-place at SDK runtime — use it for copy/style/image/HTML changes that keep
// visitors on the same URL (headline copy tests, button color, hero image swap).
// `redirect` routes the visitor to a different URL entirely — use it for
// landing-page A/B tests, alternate pricing pages, or any test where the _page
// itself_ is the variable. They are not interchangeable: a redirect variant cannot
// also tweak DOM, and a dom_modifications variant cannot send the visitor
// elsewhere.
type ExperimentVariantNewParamsVariantType string

const (
	ExperimentVariantNewParamsVariantTypeDomModifications ExperimentVariantNewParamsVariantType = "dom_modifications"
	ExperimentVariantNewParamsVariantTypeRedirect         ExperimentVariantNewParamsVariantType = "redirect"
)

type ExperimentVariantUpdateParams struct {
	// Promote or demote this variant as the control. Promoting a second variant while
	// another already has `isControl: true` is rejected with 409 — clear the existing
	// control first.
	IsControl param.Opt[bool] `json:"isControl,omitzero"`
	// Updated variant name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Updated redirect URL for redirect variants. Use either a site-relative path such
	// as `/pricing-v2` or an absolute `https://` URL. Cross-origin `http://` URLs are
	// rejected.
	RedirectURL param.Opt[string] `json:"redirectUrl,omitzero"`
	// Updated traffic weight relative to other variants. Must be a positive integer in
	// the range 1..1_000_000.
	Weight param.Opt[int64] `json:"weight,omitzero"`
	// Updated declarative DOM mutations. Sending this field replaces the prior list —
	// partial-array merging is not supported.
	DomModifications []ExperimentVariantUpdateParamsDomModification `json:"domModifications,omitzero"`
	// Updated variant delivery mechanism. `dom_modifications` mutates the current page
	// in-place; `redirect` sends the visitor to a different URL — pick based on
	// whether the _page_ or the _content_ is the variable. Changing this also requires
	// updating the matching payload field (`redirectUrl` or `domModifications`).
	//
	// Any of "dom_modifications", "redirect".
	VariantType ExperimentVariantUpdateParamsVariantType `json:"variantType,omitzero"`
	paramObj
}

func (r ExperimentVariantUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentVariantUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentVariantUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Action, Selector are required.
type ExperimentVariantUpdateParamsDomModification struct {
	// Mutation to apply when the selector matches.
	//
	// Any of "customCss", "customJs", "insertAfter", "insertBefore", "remove",
	// "setAttribute", "setHtml", "setImage", "setStyle", "setText".
	Action string `json:"action,omitzero" api:"required"`
	// CSS selector for the element to modify at runtime. PREFER specific selectors
	// that match exactly one element: an `id` (`#hero-headline`), a stable `data-*`
	// attribute (`[data-testid="hero-headline"]`), or a unique class/structural chain
	// (`section.hero > h1.headline`). AVOID bare tag selectors like `h1`, `button`, or
	// `img` — modern pages usually contain several, and the runtime applies the
	// mutation to ONLY THE FIRST match, which silently picks the wrong element. If you
	// only have a tag name, scope it with the nearest unique ancestor (e.g. `main h1`,
	// `header nav a:first-of-type`).
	Selector string `json:"selector" api:"required"`
	// Canonical action payload. For `setText` / `setHtml` / `customCss` / `customJs` /
	// `setImage` / `insertBefore` / `insertAfter` this is the literal
	// text/HTML/CSS/JS/URL. For `setStyle` and `setAttribute` it is a JSON-stringified
	// `{key: value}` object — or you can supply the structured `styles` / `attribute`
	// field instead and the server will normalize.
	Value param.Opt[string] `json:"value,omitzero"`
	// Use this for `setAttribute` to avoid JSON-stringifying `{name: value}` yourself.
	// Ignored for other actions.
	Attribute any `json:"attribute,omitzero"`
	// Use this for `setStyle` to avoid JSON-stringifying `{property: value}` yourself.
	// Ignored for other actions.
	Styles []ExperimentVariantUpdateParamsDomModificationStyle `json:"styles,omitzero"`
	paramObj
}

func (r ExperimentVariantUpdateParamsDomModification) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentVariantUpdateParamsDomModification
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentVariantUpdateParamsDomModification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExperimentVariantUpdateParamsDomModification](
		"action", "customCss", "customJs", "insertAfter", "insertBefore", "remove", "setAttribute", "setHtml", "setImage", "setStyle", "setText",
	)
}

// The properties Property, Value are required.
type ExperimentVariantUpdateParamsDomModificationStyle struct {
	// CSS property name in camelCase or kebab-case.
	Property string `json:"property" api:"required"`
	// CSS value to assign to the property.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r ExperimentVariantUpdateParamsDomModificationStyle) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentVariantUpdateParamsDomModificationStyle
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentVariantUpdateParamsDomModificationStyle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Updated variant delivery mechanism. `dom_modifications` mutates the current page
// in-place; `redirect` sends the visitor to a different URL — pick based on
// whether the _page_ or the _content_ is the variable. Changing this also requires
// updating the matching payload field (`redirectUrl` or `domModifications`).
type ExperimentVariantUpdateParamsVariantType string

const (
	ExperimentVariantUpdateParamsVariantTypeDomModifications ExperimentVariantUpdateParamsVariantType = "dom_modifications"
	ExperimentVariantUpdateParamsVariantTypeRedirect         ExperimentVariantUpdateParamsVariantType = "redirect"
)
