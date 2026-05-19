// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/with-ours/platform-sdk-go/internal/apijson"
	"github.com/with-ours/platform-sdk-go/internal/apiquery"
	"github.com/with-ours/platform-sdk-go/internal/requestconfig"
	"github.com/with-ours/platform-sdk-go/option"
	"github.com/with-ours/platform-sdk-go/packages/respjson"
)

// MappingTemplateService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMappingTemplateService] method instead.
type MappingTemplateService struct {
	options []option.RequestOption
}

// NewMappingTemplateService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMappingTemplateService(opts ...option.RequestOption) (r MappingTemplateService) {
	r = MappingTemplateService{}
	r.options = opts
	return
}

// Discover every mapping template available for a destination or source, with full
// property descriptors inlined. Use the returned `id` as `templateId` when calling
// `POST /rest/v1/mappings` (template fat-create variant), and use each entry under
// `mappings[]` to learn the valid `property`, `kind`, `modificationOptions`, and
// any enforced `options`. The `isDefault: true` entry is the destination's
// built-in default template (the one stored at `MAPPER#{destinationId}` when
// configured via `PUT /rest/v1/default-mappings/{destinationId}`). Requires scope:
// mapping:find
func (r *MappingTemplateService) List(ctx context.Context, query MappingTemplateListParams, opts ...option.RequestOption) (res *MappingTemplateListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "rest/v1/mapping-templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type MappingTemplateListResponse struct {
	Entities []MappingTemplateListResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MappingTemplateListResponse) RawJSON() string { return r.JSON.raw }
func (r *MappingTemplateListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingTemplateListResponseEntity struct {
	// Template identifier — pass to `POST /rest/v1/mappings` as `templateId`.
	ID string `json:"id" api:"required"`
	// True for the destination's built-in default template (the one stored at
	// `MAPPER#{destinationId}` when configured). Sources only have one template; it is
	// always default.
	IsDefault   bool                                       `json:"isDefault" api:"required"`
	Mappings    []MappingTemplateListResponseEntityMapping `json:"mappings" api:"required"`
	Name        string                                     `json:"name" api:"required"`
	Description string                                     `json:"description" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		IsDefault   respjson.Field
		Mappings    respjson.Field
		Name        respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MappingTemplateListResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *MappingTemplateListResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingTemplateListResponseEntityMapping struct {
	// Long-form description / tooltip for this property.
	Description string `json:"description" api:"required"`
	IsPii       bool   `json:"isPII" api:"required"`
	// Type information for SDK validation (Text, Integer, Email, Url, IP, Object,
	// KnownObject, Date, DateTime, Array, Boolean, JSON).
	//
	// Any of "Array", "Boolean", "Date", "DateTime", "Email", "IP", "Integer", "JSON",
	// "KnownObject", "Object", "Text", "Url".
	Kind string `json:"kind" api:"required"`
	// Human-readable label (e.g. "Email", "Event Name").
	Label string `json:"label" api:"required"`
	// The template default source expression, e.g. `{{visitor.email}}`.
	Map string `json:"map" api:"required"`
	// The value to send as `mappings[].property` when creating or patching a mapping.
	Property string `json:"property" api:"required"`
	Required bool   `json:"required" api:"required"`
	// The template default modification (hashing / case / URL truncation).
	//
	// Any of "CamelCase", "DmaIP", "DomainOnly", "DomainPathOnly", "DomainPathUTMs",
	// "DomainUTMs", "FakeDomain", "FakeDomainRealPath", "FakeIP", "FullUrl", "Hash",
	// "HashMD5", "HashedCountry", "HashedDateOfBirth", "HashedGender",
	// "HashedNormalized", "HashedNormalizedNoSpecialChars", "HashedPhone",
	// "HashedState", "HashedZip", "KebabCase", "LowerCase", "None", "Null",
	// "Redacted", "RegionalIP", "SnakeCase", "StartCase", "UpperCase".
	Modification string `json:"modification" api:"nullable"`
	// Suggested modification options for this property. Not a whitelist.
	//
	// Any of "CamelCase", "DmaIP", "DomainOnly", "DomainPathOnly", "DomainPathUTMs",
	// "DomainUTMs", "FakeDomain", "FakeDomainRealPath", "FakeIP", "FullUrl", "Hash",
	// "HashMD5", "HashedCountry", "HashedDateOfBirth", "HashedGender",
	// "HashedNormalized", "HashedNormalizedNoSpecialChars", "HashedPhone",
	// "HashedState", "HashedZip", "KebabCase", "LowerCase", "None", "Null",
	// "Redacted", "RegionalIP", "SnakeCase", "StartCase", "UpperCase".
	ModificationOptions []string `json:"modificationOptions" api:"nullable"`
	// When set, the ONLY valid `map` values for this property. Typically used for
	// enum-shaped destinations.
	Options []MappingTemplateListResponseEntityMappingOption `json:"options" api:"nullable"`
	// Non-binding suggestions for the `map` value (e.g. common event names a customer
	// might want to use).
	SuggestedOptions []MappingTemplateListResponseEntityMappingSuggestedOption `json:"suggestedOptions" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description         respjson.Field
		IsPii               respjson.Field
		Kind                respjson.Field
		Label               respjson.Field
		Map                 respjson.Field
		Property            respjson.Field
		Required            respjson.Field
		Modification        respjson.Field
		ModificationOptions respjson.Field
		Options             respjson.Field
		SuggestedOptions    respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MappingTemplateListResponseEntityMapping) RawJSON() string { return r.JSON.raw }
func (r *MappingTemplateListResponseEntityMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingTemplateListResponseEntityMappingOption struct {
	Label string `json:"label" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MappingTemplateListResponseEntityMappingOption) RawJSON() string { return r.JSON.raw }
func (r *MappingTemplateListResponseEntityMappingOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingTemplateListResponseEntityMappingSuggestedOption struct {
	Label string `json:"label" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MappingTemplateListResponseEntityMappingSuggestedOption) RawJSON() string { return r.JSON.raw }
func (r *MappingTemplateListResponseEntityMappingSuggestedOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingTemplateListParams struct {
	// Destination or source id. Required.
	EntityID string `query:"entityId" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [MappingTemplateListParams]'s query parameters as
// `url.Values`.
func (r MappingTemplateListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
