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
	"github.com/with-ours/platform-sdk-go/internal/requestconfig"
	"github.com/with-ours/platform-sdk-go/option"
	"github.com/with-ours/platform-sdk-go/packages/param"
	"github.com/with-ours/platform-sdk-go/packages/respjson"
)

// DefaultMappingService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDefaultMappingService] method instead.
type DefaultMappingService struct {
	Options []option.RequestOption
}

// NewDefaultMappingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDefaultMappingService(opts ...option.RequestOption) (r DefaultMappingService) {
	r = DefaultMappingService{}
	r.Options = opts
	return
}

// List every stored default mapping for the account, one per destination that has
// ever written a default. Destinations that have not yet written a default mapping
// do not appear here. Use `GET /rest/v1/default-mappings/{destinationId}` to fetch
// the hydrated would-be row for a specific destination. Requires scope:
// mapping:list
func (r *DefaultMappingService) List(ctx context.Context, opts ...option.RequestOption) (res *DefaultMappingListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/default-mappings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Fetch the destination's default mapping by destination id. Returns a hydrated
// row with empty `mappings[]` when no default mapping has been written yet (so
// callers do not need to handle a 404-vs-200 branch). Each destination has at most
// one default mapping. Requires scope: mapping:find
func (r *DefaultMappingService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *DefaultMappingGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/default-mappings/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Upsert the destination default mapping. Always replaces `mappings[]` wholesale
// (default mappings have no merge-partial semantic). Default mappings cannot have
// custom `logic`; the field is not accepted on this endpoint. Requires scope:
// mapping:update
func (r *DefaultMappingService) Replace(ctx context.Context, id string, body DefaultMappingReplaceParams, opts ...option.RequestOption) (res *DefaultMappingReplaceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/default-mappings/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

type DefaultMappingListResponse struct {
	Entities []DefaultMappingListResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DefaultMappingListResponse) RawJSON() string { return r.JSON.raw }
func (r *DefaultMappingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultMappingListResponseEntity struct {
	ID               string                                    `json:"id" api:"required"`
	IsEnabled        bool                                      `json:"isEnabled" api:"required"`
	Mappings         []DefaultMappingListResponseEntityMapping `json:"mappings" api:"required"`
	DestinationID    string                                    `json:"destinationId" api:"nullable"`
	IsDefaultMapping bool                                      `json:"isDefaultMapping" api:"nullable"`
	// Condition tree gating when this mapping fires. A node is either a leaf
	// `condition` or a combinator (`AND`, `OR`, `NOT`). Combinator children are
	// themselves `MappingLogic` nodes, so trees nest arbitrarily. Example leaf:
	// `{ "condition": { "property": "$event.event", "operator": "Is", "value": "page_view" } }`.
	// Example combinator: `{ "AND": [{ "condition": ... }, { "OR": [...] }] }`.
	Logic        any     `json:"logic"`
	Name         string  `json:"name" api:"nullable"`
	Priority     float64 `json:"priority" api:"nullable"`
	SourceID     string  `json:"sourceId" api:"nullable"`
	TemplateID   string  `json:"templateId" api:"nullable"`
	TemplateName string  `json:"templateName" api:"nullable"`
	UpdatedAt    string  `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		IsEnabled        respjson.Field
		Mappings         respjson.Field
		DestinationID    respjson.Field
		IsDefaultMapping respjson.Field
		Logic            respjson.Field
		Name             respjson.Field
		Priority         respjson.Field
		SourceID         respjson.Field
		TemplateID       respjson.Field
		TemplateName     respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DefaultMappingListResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *DefaultMappingListResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultMappingListResponseEntityMapping struct {
	Map      string `json:"map" api:"required"`
	Property string `json:"property" api:"required"`
	// Any of "CamelCase", "DmaIP", "DomainOnly", "DomainPathOnly", "DomainPathUTMs",
	// "DomainUTMs", "FakeDomain", "FakeDomainRealPath", "FakeIP", "FullUrl", "Hash",
	// "HashMD5", "HashedCountry", "HashedDateOfBirth", "HashedGender",
	// "HashedNormalized", "HashedNormalizedNoSpecialChars", "HashedPhone",
	// "HashedState", "HashedZip", "KebabCase", "LowerCase", "None", "Null",
	// "Redacted", "RegionalIP", "SnakeCase", "StartCase", "UpperCase".
	Modification string `json:"modification" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Map          respjson.Field
		Property     respjson.Field
		Modification respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DefaultMappingListResponseEntityMapping) RawJSON() string { return r.JSON.raw }
func (r *DefaultMappingListResponseEntityMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultMappingGetResponse struct {
	ID               string                             `json:"id" api:"required"`
	IsEnabled        bool                               `json:"isEnabled" api:"required"`
	Mappings         []DefaultMappingGetResponseMapping `json:"mappings" api:"required"`
	DestinationID    string                             `json:"destinationId" api:"nullable"`
	IsDefaultMapping bool                               `json:"isDefaultMapping" api:"nullable"`
	// Condition tree gating when this mapping fires. A node is either a leaf
	// `condition` or a combinator (`AND`, `OR`, `NOT`). Combinator children are
	// themselves `MappingLogic` nodes, so trees nest arbitrarily. Example leaf:
	// `{ "condition": { "property": "$event.event", "operator": "Is", "value": "page_view" } }`.
	// Example combinator: `{ "AND": [{ "condition": ... }, { "OR": [...] }] }`.
	Logic        any     `json:"logic"`
	Name         string  `json:"name" api:"nullable"`
	Priority     float64 `json:"priority" api:"nullable"`
	SourceID     string  `json:"sourceId" api:"nullable"`
	TemplateID   string  `json:"templateId" api:"nullable"`
	TemplateName string  `json:"templateName" api:"nullable"`
	UpdatedAt    string  `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		IsEnabled        respjson.Field
		Mappings         respjson.Field
		DestinationID    respjson.Field
		IsDefaultMapping respjson.Field
		Logic            respjson.Field
		Name             respjson.Field
		Priority         respjson.Field
		SourceID         respjson.Field
		TemplateID       respjson.Field
		TemplateName     respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DefaultMappingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DefaultMappingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultMappingGetResponseMapping struct {
	Map      string `json:"map" api:"required"`
	Property string `json:"property" api:"required"`
	// Any of "CamelCase", "DmaIP", "DomainOnly", "DomainPathOnly", "DomainPathUTMs",
	// "DomainUTMs", "FakeDomain", "FakeDomainRealPath", "FakeIP", "FullUrl", "Hash",
	// "HashMD5", "HashedCountry", "HashedDateOfBirth", "HashedGender",
	// "HashedNormalized", "HashedNormalizedNoSpecialChars", "HashedPhone",
	// "HashedState", "HashedZip", "KebabCase", "LowerCase", "None", "Null",
	// "Redacted", "RegionalIP", "SnakeCase", "StartCase", "UpperCase".
	Modification string `json:"modification" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Map          respjson.Field
		Property     respjson.Field
		Modification respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DefaultMappingGetResponseMapping) RawJSON() string { return r.JSON.raw }
func (r *DefaultMappingGetResponseMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultMappingReplaceResponse struct {
	ID               string                                 `json:"id" api:"required"`
	IsEnabled        bool                                   `json:"isEnabled" api:"required"`
	Mappings         []DefaultMappingReplaceResponseMapping `json:"mappings" api:"required"`
	DestinationID    string                                 `json:"destinationId" api:"nullable"`
	IsDefaultMapping bool                                   `json:"isDefaultMapping" api:"nullable"`
	// Condition tree gating when this mapping fires. A node is either a leaf
	// `condition` or a combinator (`AND`, `OR`, `NOT`). Combinator children are
	// themselves `MappingLogic` nodes, so trees nest arbitrarily. Example leaf:
	// `{ "condition": { "property": "$event.event", "operator": "Is", "value": "page_view" } }`.
	// Example combinator: `{ "AND": [{ "condition": ... }, { "OR": [...] }] }`.
	Logic        any     `json:"logic"`
	Name         string  `json:"name" api:"nullable"`
	Priority     float64 `json:"priority" api:"nullable"`
	SourceID     string  `json:"sourceId" api:"nullable"`
	TemplateID   string  `json:"templateId" api:"nullable"`
	TemplateName string  `json:"templateName" api:"nullable"`
	UpdatedAt    string  `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		IsEnabled        respjson.Field
		Mappings         respjson.Field
		DestinationID    respjson.Field
		IsDefaultMapping respjson.Field
		Logic            respjson.Field
		Name             respjson.Field
		Priority         respjson.Field
		SourceID         respjson.Field
		TemplateID       respjson.Field
		TemplateName     respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DefaultMappingReplaceResponse) RawJSON() string { return r.JSON.raw }
func (r *DefaultMappingReplaceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultMappingReplaceResponseMapping struct {
	Map      string `json:"map" api:"required"`
	Property string `json:"property" api:"required"`
	// Any of "CamelCase", "DmaIP", "DomainOnly", "DomainPathOnly", "DomainPathUTMs",
	// "DomainUTMs", "FakeDomain", "FakeDomainRealPath", "FakeIP", "FullUrl", "Hash",
	// "HashMD5", "HashedCountry", "HashedDateOfBirth", "HashedGender",
	// "HashedNormalized", "HashedNormalizedNoSpecialChars", "HashedPhone",
	// "HashedState", "HashedZip", "KebabCase", "LowerCase", "None", "Null",
	// "Redacted", "RegionalIP", "SnakeCase", "StartCase", "UpperCase".
	Modification string `json:"modification" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Map          respjson.Field
		Property     respjson.Field
		Modification respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DefaultMappingReplaceResponseMapping) RawJSON() string { return r.JSON.raw }
func (r *DefaultMappingReplaceResponseMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DefaultMappingReplaceParams struct {
	// Property mappings to persist as the destination default. Use
	// `GET /rest/v1/mapping-templates?entityId={destinationId}` to discover the valid
	// `property` values.
	Mappings []DefaultMappingReplaceParamsMapping `json:"mappings,omitzero" api:"required"`
	// Toggle the default mapping on/off. Defaults to `true` when omitted. `null` is
	// treated as omitted.
	IsEnabled param.Opt[bool] `json:"isEnabled,omitzero"`
	paramObj
}

func (r DefaultMappingReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow DefaultMappingReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DefaultMappingReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Map, Property are required.
type DefaultMappingReplaceParamsMapping struct {
	Map      string `json:"map" api:"required"`
	Property string `json:"property" api:"required"`
	// Any of "CamelCase", "DmaIP", "DomainOnly", "DomainPathOnly", "DomainPathUTMs",
	// "DomainUTMs", "FakeDomain", "FakeDomainRealPath", "FakeIP", "FullUrl", "Hash",
	// "HashMD5", "HashedCountry", "HashedDateOfBirth", "HashedGender",
	// "HashedNormalized", "HashedNormalizedNoSpecialChars", "HashedPhone",
	// "HashedState", "HashedZip", "KebabCase", "LowerCase", "None", "Null",
	// "Redacted", "RegionalIP", "SnakeCase", "StartCase", "UpperCase".
	Modification string `json:"modification,omitzero"`
	paramObj
}

func (r DefaultMappingReplaceParamsMapping) MarshalJSON() (data []byte, err error) {
	type shadow DefaultMappingReplaceParamsMapping
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DefaultMappingReplaceParamsMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DefaultMappingReplaceParamsMapping](
		"modification", "CamelCase", "DmaIP", "DomainOnly", "DomainPathOnly", "DomainPathUTMs", "DomainUTMs", "FakeDomain", "FakeDomainRealPath", "FakeIP", "FullUrl", "Hash", "HashMD5", "HashedCountry", "HashedDateOfBirth", "HashedGender", "HashedNormalized", "HashedNormalizedNoSpecialChars", "HashedPhone", "HashedState", "HashedZip", "KebabCase", "LowerCase", "None", "Null", "Redacted", "RegionalIP", "SnakeCase", "StartCase", "UpperCase",
	)
}
