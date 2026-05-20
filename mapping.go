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

// MappingService contains methods and other services that help with interacting
// with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMappingService] method instead.
type MappingService struct {
	Options []option.RequestOption
}

// NewMappingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMappingService(opts ...option.RequestOption) (r MappingService) {
	r = MappingService{}
	r.Options = opts
	return
}

// List mappings for an entity (a source or destination). Requires the `entityId`
// query parameter. Supports cursor pagination via `limit` and `cursor`. Sorted by
// `priority` ascending, then by `id` for deterministic pagination. Requires scope:
// mapping:list
func (r *MappingService) List(ctx context.Context, query MappingListParams, opts ...option.RequestOption) (res *pagination.Cursor[MappingListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "rest/v1/mappings"
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

// List mappings for an entity (a source or destination). Requires the `entityId`
// query parameter. Supports cursor pagination via `limit` and `cursor`. Sorted by
// `priority` ascending, then by `id` for deterministic pagination. Requires scope:
// mapping:list
func (r *MappingService) ListAutoPaging(ctx context.Context, query MappingListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[MappingListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Create a mapping. Two body shapes are accepted:
//
//  1. Quick-create — `{ allowedEventId, destinationId }`. Binds an allowed event to
//     a destination. Returns a slim entity with empty `mappings[]`; follow up with
//     PATCH to populate fields.
//  2. Template fat-create —
//     `{ entityId, templateId, mappings?, logic?, isEnabled?, name?, insertAfterIdx? }`.
//     Lands a fully-shaped mapping in one round-trip. Use
//     `GET /rest/v1/mapping-templates?entityId=...` to discover the valid
//     `templateId` and `mappings[].property` values.
//
// Sending both `allowedEventId` and `templateId` returns 400. Requires scope:
// mapping:create
func (r *MappingService) New(ctx context.Context, body MappingNewParams, opts ...option.RequestOption) (res *MappingNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/mappings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Find a single mapping by ID. Requires scope: mapping:find
func (r *MappingService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *MappingGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/mappings/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Partially update a mapping. Only the fields you send are changed. Send
// `isEnabled: false` to pause the mapping without changing other fields (mirrors
// `status` on destinations). `mappings[]` is replaced wholesale when sent.
// Requires scope: mapping:update
func (r *MappingService) Update(ctx context.Context, id string, body MappingUpdateParams, opts ...option.RequestOption) (res *MappingUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/mappings/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete a mapping. Requires scope: mapping:delete
func (r *MappingService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/mappings/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Reassign `priority` for a set of mappings. Pass `{ uuids: [...] }` with the
// mapping ids in their new order — index 0 becomes the highest-priority mapping.
// All ids must belong to the same parent entity (source or destination); mixing
// entities returns 400. Requires scope: mapping:update
func (r *MappingService) Reorder(ctx context.Context, body MappingReorderParams, opts ...option.RequestOption) (res *MappingReorderResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/mappings/reorder"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Discover every mapping template available for a destination or source, with full
// property descriptors inlined. Use the returned `id` as `templateId` when calling
// `POST /rest/v1/mappings` (template fat-create variant), and use each entry under
// `mappings[]` to learn the valid `property`, `kind`, `modificationOptions`, and
// any enforced `options`. The `isDefault: true` entry is the destination's
// built-in default template (the one stored at `MAPPER#{destinationId}` when
// configured via `PUT /rest/v1/default-mappings/{destinationId}`). Requires scope:
// mapping:find
func (r *MappingService) Templates(ctx context.Context, query MappingTemplatesParams, opts ...option.RequestOption) (res *MappingTemplatesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/mappings/templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Lists the platform-provided variables that any mapping `value` can reference
// (e.g. `event.email`, `event.request_context.ip`, `visitor.id`). Account-agnostic
// discovery — use these paths as the right-hand side of a mapping field. Requires
// scope: variables:find-default
func (r *MappingService) DefaultVariables(ctx context.Context, opts ...option.RequestOption) (res *MappingDefaultVariablesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/mappings/default-variables"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists the custom variables observed in this account’s recent event stream (last
// 14 days). These are dot-paths under `event.event_properties.*` that callers can
// target in mapping `value` fields. The result is cached for 10 minutes; an empty
// list means no custom properties have been seen yet for this account. Requires
// scope: variables:find-custom
func (r *MappingService) CustomVariables(ctx context.Context, opts ...option.RequestOption) (res *MappingCustomVariablesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/mappings/custom-variables"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists every value accepted on a mapping field’s `modification` property, with a
// human-readable label and one-sentence description. Account-agnostic. Use this
// alongside `GET /rest/v1/mapping-templates` to render a labelled modification
// picker without hardcoding the enum. Requires scope: variables:find-default
func (r *MappingService) Modifications(ctx context.Context, opts ...option.RequestOption) (res *MappingModificationsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/mappings/modifications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type MappingListResponse struct {
	ID               string                       `json:"id" api:"required"`
	IsEnabled        bool                         `json:"isEnabled" api:"required"`
	Mappings         []MappingListResponseMapping `json:"mappings" api:"required"`
	DestinationID    string                       `json:"destinationId" api:"nullable"`
	IsDefaultMapping bool                         `json:"isDefaultMapping" api:"nullable"`
	// Condition tree gating when this mapping fires. A node is either a leaf
	// `condition` or a combinator (`AND`, `OR`, `NOT`). Combinator children are
	// themselves logic nodes, so trees nest arbitrarily.
	//
	// Example leaf:
	// `{ "condition": { "property": "$event.event", "operator": "Is", "value": "page_view" } }`.
	//
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
func (r MappingListResponse) RawJSON() string { return r.JSON.raw }
func (r *MappingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingListResponseMapping struct {
	// Source expression sent to the destination for this `property`. Use `{{...}}`
	// template syntax to substitute values from the event/visitor record:
	// `{{event.event}}`, `{{event.event_properties.value}}`, `{{visitor.email}}`. Bare
	// strings (no `{{}}`) are sent verbatim. Note: `{{...}}` template syntax belongs
	// HERE, NOT in `logic.condition.property` — logic conditions use bare dotted paths
	// like `$event.event_properties.value`.
	Map string `json:"map" api:"required"`
	// Destination-side field name. Comes from the destination template — discover the
	// valid set via `GET /rest/v1/mapping-templates?entityId=...`.
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
func (r MappingListResponseMapping) RawJSON() string { return r.JSON.raw }
func (r *MappingListResponseMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingNewResponse struct {
	ID               string                      `json:"id" api:"required"`
	IsEnabled        bool                        `json:"isEnabled" api:"required"`
	Mappings         []MappingNewResponseMapping `json:"mappings" api:"required"`
	DestinationID    string                      `json:"destinationId" api:"nullable"`
	IsDefaultMapping bool                        `json:"isDefaultMapping" api:"nullable"`
	// Condition tree gating when this mapping fires. A node is either a leaf
	// `condition` or a combinator (`AND`, `OR`, `NOT`). Combinator children are
	// themselves logic nodes, so trees nest arbitrarily.
	//
	// Example leaf:
	// `{ "condition": { "property": "$event.event", "operator": "Is", "value": "page_view" } }`.
	//
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
func (r MappingNewResponse) RawJSON() string { return r.JSON.raw }
func (r *MappingNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingNewResponseMapping struct {
	// Source expression sent to the destination for this `property`. Use `{{...}}`
	// template syntax to substitute values from the event/visitor record:
	// `{{event.event}}`, `{{event.event_properties.value}}`, `{{visitor.email}}`. Bare
	// strings (no `{{}}`) are sent verbatim. Note: `{{...}}` template syntax belongs
	// HERE, NOT in `logic.condition.property` — logic conditions use bare dotted paths
	// like `$event.event_properties.value`.
	Map string `json:"map" api:"required"`
	// Destination-side field name. Comes from the destination template — discover the
	// valid set via `GET /rest/v1/mapping-templates?entityId=...`.
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
func (r MappingNewResponseMapping) RawJSON() string { return r.JSON.raw }
func (r *MappingNewResponseMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingGetResponse struct {
	ID               string                      `json:"id" api:"required"`
	IsEnabled        bool                        `json:"isEnabled" api:"required"`
	Mappings         []MappingGetResponseMapping `json:"mappings" api:"required"`
	DestinationID    string                      `json:"destinationId" api:"nullable"`
	IsDefaultMapping bool                        `json:"isDefaultMapping" api:"nullable"`
	// Condition tree gating when this mapping fires. A node is either a leaf
	// `condition` or a combinator (`AND`, `OR`, `NOT`). Combinator children are
	// themselves logic nodes, so trees nest arbitrarily.
	//
	// Example leaf:
	// `{ "condition": { "property": "$event.event", "operator": "Is", "value": "page_view" } }`.
	//
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
func (r MappingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MappingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingGetResponseMapping struct {
	// Source expression sent to the destination for this `property`. Use `{{...}}`
	// template syntax to substitute values from the event/visitor record:
	// `{{event.event}}`, `{{event.event_properties.value}}`, `{{visitor.email}}`. Bare
	// strings (no `{{}}`) are sent verbatim. Note: `{{...}}` template syntax belongs
	// HERE, NOT in `logic.condition.property` — logic conditions use bare dotted paths
	// like `$event.event_properties.value`.
	Map string `json:"map" api:"required"`
	// Destination-side field name. Comes from the destination template — discover the
	// valid set via `GET /rest/v1/mapping-templates?entityId=...`.
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
func (r MappingGetResponseMapping) RawJSON() string { return r.JSON.raw }
func (r *MappingGetResponseMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingUpdateResponse struct {
	ID               string                         `json:"id" api:"required"`
	IsEnabled        bool                           `json:"isEnabled" api:"required"`
	Mappings         []MappingUpdateResponseMapping `json:"mappings" api:"required"`
	DestinationID    string                         `json:"destinationId" api:"nullable"`
	IsDefaultMapping bool                           `json:"isDefaultMapping" api:"nullable"`
	// Condition tree gating when this mapping fires. A node is either a leaf
	// `condition` or a combinator (`AND`, `OR`, `NOT`). Combinator children are
	// themselves logic nodes, so trees nest arbitrarily.
	//
	// Example leaf:
	// `{ "condition": { "property": "$event.event", "operator": "Is", "value": "page_view" } }`.
	//
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
func (r MappingUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *MappingUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingUpdateResponseMapping struct {
	// Source expression sent to the destination for this `property`. Use `{{...}}`
	// template syntax to substitute values from the event/visitor record:
	// `{{event.event}}`, `{{event.event_properties.value}}`, `{{visitor.email}}`. Bare
	// strings (no `{{}}`) are sent verbatim. Note: `{{...}}` template syntax belongs
	// HERE, NOT in `logic.condition.property` — logic conditions use bare dotted paths
	// like `$event.event_properties.value`.
	Map string `json:"map" api:"required"`
	// Destination-side field name. Comes from the destination template — discover the
	// valid set via `GET /rest/v1/mapping-templates?entityId=...`.
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
func (r MappingUpdateResponseMapping) RawJSON() string { return r.JSON.raw }
func (r *MappingUpdateResponseMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingReorderResponse struct {
	Entities []MappingReorderResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MappingReorderResponse) RawJSON() string { return r.JSON.raw }
func (r *MappingReorderResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingReorderResponseEntity struct {
	ID               string                                `json:"id" api:"required"`
	IsEnabled        bool                                  `json:"isEnabled" api:"required"`
	Mappings         []MappingReorderResponseEntityMapping `json:"mappings" api:"required"`
	DestinationID    string                                `json:"destinationId" api:"nullable"`
	IsDefaultMapping bool                                  `json:"isDefaultMapping" api:"nullable"`
	// Condition tree gating when this mapping fires. A node is either a leaf
	// `condition` or a combinator (`AND`, `OR`, `NOT`). Combinator children are
	// themselves logic nodes, so trees nest arbitrarily.
	//
	// Example leaf:
	// `{ "condition": { "property": "$event.event", "operator": "Is", "value": "page_view" } }`.
	//
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
func (r MappingReorderResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *MappingReorderResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingReorderResponseEntityMapping struct {
	// Source expression sent to the destination for this `property`. Use `{{...}}`
	// template syntax to substitute values from the event/visitor record:
	// `{{event.event}}`, `{{event.event_properties.value}}`, `{{visitor.email}}`. Bare
	// strings (no `{{}}`) are sent verbatim. Note: `{{...}}` template syntax belongs
	// HERE, NOT in `logic.condition.property` — logic conditions use bare dotted paths
	// like `$event.event_properties.value`.
	Map string `json:"map" api:"required"`
	// Destination-side field name. Comes from the destination template — discover the
	// valid set via `GET /rest/v1/mapping-templates?entityId=...`.
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
func (r MappingReorderResponseEntityMapping) RawJSON() string { return r.JSON.raw }
func (r *MappingReorderResponseEntityMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingTemplatesResponse struct {
	Entities []MappingTemplatesResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MappingTemplatesResponse) RawJSON() string { return r.JSON.raw }
func (r *MappingTemplatesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingTemplatesResponseEntity struct {
	// Template identifier — pass to `POST /rest/v1/mappings` as `templateId`.
	ID string `json:"id" api:"required"`
	// True for the destination's built-in default template (the one stored at
	// `MAPPER#{destinationId}` when configured). Sources only have one template; it is
	// always default.
	IsDefault   bool                                    `json:"isDefault" api:"required"`
	Mappings    []MappingTemplatesResponseEntityMapping `json:"mappings" api:"required"`
	Name        string                                  `json:"name" api:"required"`
	Description string                                  `json:"description" api:"nullable"`
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
func (r MappingTemplatesResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *MappingTemplatesResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingTemplatesResponseEntityMapping struct {
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
	Options []MappingTemplatesResponseEntityMappingOption `json:"options" api:"nullable"`
	// Non-binding suggestions for the `map` value (e.g. common event names a customer
	// might want to use).
	SuggestedOptions []MappingTemplatesResponseEntityMappingSuggestedOption `json:"suggestedOptions" api:"nullable"`
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
func (r MappingTemplatesResponseEntityMapping) RawJSON() string { return r.JSON.raw }
func (r *MappingTemplatesResponseEntityMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingTemplatesResponseEntityMappingOption struct {
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
func (r MappingTemplatesResponseEntityMappingOption) RawJSON() string { return r.JSON.raw }
func (r *MappingTemplatesResponseEntityMappingOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingTemplatesResponseEntityMappingSuggestedOption struct {
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
func (r MappingTemplatesResponseEntityMappingSuggestedOption) RawJSON() string { return r.JSON.raw }
func (r *MappingTemplatesResponseEntityMappingSuggestedOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingDefaultVariablesResponse struct {
	Entities []MappingDefaultVariablesResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MappingDefaultVariablesResponse) RawJSON() string { return r.JSON.raw }
func (r *MappingDefaultVariablesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingDefaultVariablesResponseEntity struct {
	// Sample values observed for this path (empty for unsampled defaults).
	Examples []string `json:"examples" api:"required"`
	// Human-readable display name.
	Name string `json:"name" api:"required"`
	// Dot-path used in mapping `value` fields (e.g. `event.email`).
	Path string `json:"path" api:"required"`
	// Relative popularity rank. Higher means more frequently set across events.
	Popularity float64 `json:"popularity" api:"required"`
	// Optional long-form context shown in the variable dictionary drawer.
	AdvancedInfo string `json:"advancedInfo" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Examples     respjson.Field
		Name         respjson.Field
		Path         respjson.Field
		Popularity   respjson.Field
		AdvancedInfo respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MappingDefaultVariablesResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *MappingDefaultVariablesResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingCustomVariablesResponse struct {
	Entities []MappingCustomVariablesResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MappingCustomVariablesResponse) RawJSON() string { return r.JSON.raw }
func (r *MappingCustomVariablesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingCustomVariablesResponseEntity struct {
	// Sample values observed for this path (empty for unsampled defaults).
	Examples []string `json:"examples" api:"required"`
	// Human-readable display name.
	Name string `json:"name" api:"required"`
	// Dot-path used in mapping `value` fields (e.g. `event.email`).
	Path string `json:"path" api:"required"`
	// Relative popularity rank. Higher means more frequently set across events.
	Popularity float64 `json:"popularity" api:"required"`
	// Optional long-form context shown in the variable dictionary drawer.
	AdvancedInfo string `json:"advancedInfo" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Examples     respjson.Field
		Name         respjson.Field
		Path         respjson.Field
		Popularity   respjson.Field
		AdvancedInfo respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MappingCustomVariablesResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *MappingCustomVariablesResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingModificationsResponse struct {
	Entities []MappingModificationsResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MappingModificationsResponse) RawJSON() string { return r.JSON.raw }
func (r *MappingModificationsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingModificationsResponseEntity struct {
	// One-sentence explanation of what the modification does to the mapped value.
	Description string `json:"description" api:"required"`
	// Short human-readable name (suitable for picker labels).
	Label string `json:"label" api:"required"`
	// Enum value to send on `modification` fields when authoring a mapping.
	//
	// Any of "CamelCase", "DmaIP", "DomainOnly", "DomainPathOnly", "DomainPathUTMs",
	// "DomainUTMs", "FakeDomain", "FakeDomainRealPath", "FakeIP", "FullUrl", "Hash",
	// "HashMD5", "HashedCountry", "HashedDateOfBirth", "HashedGender",
	// "HashedNormalized", "HashedNormalizedNoSpecialChars", "HashedPhone",
	// "HashedState", "HashedZip", "KebabCase", "LowerCase", "None", "Null",
	// "Redacted", "RegionalIP", "SnakeCase", "StartCase", "UpperCase".
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Label       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MappingModificationsResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *MappingModificationsResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingListParams struct {
	// Filter mappings by their parent entity id. Must be a destination id or source
	// id.
	EntityID string `query:"entityId" api:"required" json:"-"`
	// Maximum number of mappings to return. Defaults to 1000; values below 1 are
	// clamped to 1 and values above 1000 are clamped to 1000. Most accounts can fetch
	// the full list in one request.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Opaque pagination cursor from pagination.nextCursor in the previous response. Do
	// not decode or modify it. Malformed cursors return 400 Bad Request.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MappingListParams]'s query parameters as `url.Values`.
func (r MappingListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MappingNewParams struct {
	// Template fat-create only. Override the auto-generated mapping name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Quick-create variant: allowed event to bind the new mapping to. Required
	// together with `destinationId`. Mutually exclusive with `templateId`/`entityId`.
	AllowedEventID param.Opt[string] `json:"allowedEventId,omitzero"`
	// Quick-create variant: destination that should receive events matched by this
	// mapping. Required together with `allowedEventId`.
	DestinationID param.Opt[string] `json:"destinationId,omitzero"`
	// Template fat-create variant: destination or source id this mapping belongs to.
	// Required together with `templateId`.
	EntityID param.Opt[string] `json:"entityId,omitzero"`
	// Template fat-create only. Zero-based position in the destination/source priority
	// order to insert this mapping after. Omit to append at the end.
	InsertAfterIdx param.Opt[int64] `json:"insertAfterIdx,omitzero"`
	// Template fat-create only. Initial enabled state. Defaults to `true`.
	IsEnabled param.Opt[bool] `json:"isEnabled,omitzero"`
	// Template fat-create variant: template id from `GET /rest/v1/mapping-templates`.
	// Picks the property descriptor set used to validate `mappings[].property`.
	// Required together with `entityId`.
	TemplateID param.Opt[string] `json:"templateId,omitzero"`
	// Condition tree gating when this mapping fires. A node is either a leaf
	// `condition` or a combinator (`AND`, `OR`, `NOT`). Combinator children are
	// themselves logic nodes, so trees nest arbitrarily.
	//
	// Example leaf:
	// `{ "condition": { "property": "$event.event", "operator": "Is", "value": "page_view" } }`.
	//
	// Example combinator: `{ "AND": [{ "condition": ... }, { "OR": [...] }] }`.
	Logic MappingNewParamsLogic `json:"logic,omitzero"`
	// Template fat-create only. Optional initial property mappings. When omitted the
	// mapping is seeded with template defaults for sources and non-default destination
	// templates, and with `[]` for default destination templates.
	Mappings []MappingNewParamsMapping `json:"mappings,omitzero"`
	paramObj
}

func (r MappingNewParams) MarshalJSON() (data []byte, err error) {
	type shadow MappingNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MappingNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Condition tree gating when this mapping fires. A node is either a leaf
// `condition` or a combinator (`AND`, `OR`, `NOT`). Combinator children are
// themselves logic nodes, so trees nest arbitrarily.
//
// Example leaf:
// `{ "condition": { "property": "$event.event", "operator": "Is", "value": "page_view" } }`.
//
// Example combinator: `{ "AND": [{ "condition": ... }, { "OR": [...] }] }`.
type MappingNewParamsLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And []any `json:"AND,omitzero"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or        []any                          `json:"OR,omitzero"`
	Condition MappingNewParamsLogicCondition `json:"condition,omitzero"`
	// Negates a single child logic node.
	Not any `json:"NOT,omitzero"`
	paramObj
}

func (r MappingNewParamsLogic) MarshalJSON() (data []byte, err error) {
	type shadow MappingNewParamsLogic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MappingNewParamsLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Operator, Property, Value are required.
type MappingNewParamsLogicCondition struct {
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

func (r MappingNewParamsLogicCondition) MarshalJSON() (data []byte, err error) {
	type shadow MappingNewParamsLogicCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MappingNewParamsLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MappingNewParamsLogicCondition](
		"operator", "Is", "IsNot", "Contains", "DoesNotContain", "StartsWith", "EndsWith", "IsFalsy", "IsTruthy", "IsNull", "IsNotNull", "IsUndefined", "IsNotUndefined", "IsGreaterThan", "IsGreaterThanOrEqual", "IsLessThan", "IsLessThanOrEqual", "IsIn", "IsNotIn", "IsFoundIn", "IsNotFoundIn", "IsTrue", "IsFalse", "IsBefore", "IsAfter", "IsBetween", "IsOnOrBefore", "IsOnOrAfter", "MatchesRegex", "MatchesRegexIgnoreCase", "DoesNotMatchRegex", "DoesNotMatchRegexIgnoreCase",
	)
}

// The properties Map, Property are required.
type MappingNewParamsMapping struct {
	// Source expression sent to the destination for this `property`. Use `{{...}}`
	// template syntax to substitute values from the event/visitor record:
	// `{{event.event}}`, `{{event.event_properties.value}}`, `{{visitor.email}}`. Bare
	// strings (no `{{}}`) are sent verbatim. Note: `{{...}}` template syntax belongs
	// HERE, NOT in `logic.condition.property` — logic conditions use bare dotted paths
	// like `$event.event_properties.value`.
	Map string `json:"map" api:"required"`
	// Destination-side field name. Comes from the destination template — discover the
	// valid set via `GET /rest/v1/mapping-templates?entityId=...`.
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

func (r MappingNewParamsMapping) MarshalJSON() (data []byte, err error) {
	type shadow MappingNewParamsMapping
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MappingNewParamsMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MappingNewParamsMapping](
		"modification", "CamelCase", "DmaIP", "DomainOnly", "DomainPathOnly", "DomainPathUTMs", "DomainUTMs", "FakeDomain", "FakeDomainRealPath", "FakeIP", "FullUrl", "Hash", "HashMD5", "HashedCountry", "HashedDateOfBirth", "HashedGender", "HashedNormalized", "HashedNormalizedNoSpecialChars", "HashedPhone", "HashedState", "HashedZip", "KebabCase", "LowerCase", "None", "Null", "Redacted", "RegionalIP", "SnakeCase", "StartCase", "UpperCase",
	)
}

type MappingUpdateParams struct {
	// Flip the mapping on/off without changing other fields. `null` is treated as
	// omitted.
	IsEnabled param.Opt[bool]   `json:"isEnabled,omitzero"`
	Name      param.Opt[string] `json:"name,omitzero"`
	// Condition tree gating when this mapping fires. A node is either a leaf
	// `condition` or a combinator (`AND`, `OR`, `NOT`). Combinator children are
	// themselves logic nodes, so trees nest arbitrarily.
	//
	// Example leaf:
	// `{ "condition": { "property": "$event.event", "operator": "Is", "value": "page_view" } }`.
	//
	// Example combinator: `{ "AND": [{ "condition": ... }, { "OR": [...] }] }`.
	Logic    MappingUpdateParamsLogic     `json:"logic,omitzero"`
	Mappings []MappingUpdateParamsMapping `json:"mappings,omitzero"`
	paramObj
}

func (r MappingUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow MappingUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MappingUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Condition tree gating when this mapping fires. A node is either a leaf
// `condition` or a combinator (`AND`, `OR`, `NOT`). Combinator children are
// themselves logic nodes, so trees nest arbitrarily.
//
// Example leaf:
// `{ "condition": { "property": "$event.event", "operator": "Is", "value": "page_view" } }`.
//
// Example combinator: `{ "AND": [{ "condition": ... }, { "OR": [...] }] }`.
type MappingUpdateParamsLogic struct {
	// All child nodes must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	And []any `json:"AND,omitzero"`
	// Any child node must match. Each child is itself a logic node (leaf `condition`
	// or combinator).
	Or        []any                             `json:"OR,omitzero"`
	Condition MappingUpdateParamsLogicCondition `json:"condition,omitzero"`
	// Negates a single child logic node.
	Not any `json:"NOT,omitzero"`
	paramObj
}

func (r MappingUpdateParamsLogic) MarshalJSON() (data []byte, err error) {
	type shadow MappingUpdateParamsLogic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MappingUpdateParamsLogic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Operator, Property, Value are required.
type MappingUpdateParamsLogicCondition struct {
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

func (r MappingUpdateParamsLogicCondition) MarshalJSON() (data []byte, err error) {
	type shadow MappingUpdateParamsLogicCondition
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MappingUpdateParamsLogicCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MappingUpdateParamsLogicCondition](
		"operator", "Is", "IsNot", "Contains", "DoesNotContain", "StartsWith", "EndsWith", "IsFalsy", "IsTruthy", "IsNull", "IsNotNull", "IsUndefined", "IsNotUndefined", "IsGreaterThan", "IsGreaterThanOrEqual", "IsLessThan", "IsLessThanOrEqual", "IsIn", "IsNotIn", "IsFoundIn", "IsNotFoundIn", "IsTrue", "IsFalse", "IsBefore", "IsAfter", "IsBetween", "IsOnOrBefore", "IsOnOrAfter", "MatchesRegex", "MatchesRegexIgnoreCase", "DoesNotMatchRegex", "DoesNotMatchRegexIgnoreCase",
	)
}

// The properties Map, Property are required.
type MappingUpdateParamsMapping struct {
	// Source expression sent to the destination for this `property`. Use `{{...}}`
	// template syntax to substitute values from the event/visitor record:
	// `{{event.event}}`, `{{event.event_properties.value}}`, `{{visitor.email}}`. Bare
	// strings (no `{{}}`) are sent verbatim. Note: `{{...}}` template syntax belongs
	// HERE, NOT in `logic.condition.property` — logic conditions use bare dotted paths
	// like `$event.event_properties.value`.
	Map string `json:"map" api:"required"`
	// Destination-side field name. Comes from the destination template — discover the
	// valid set via `GET /rest/v1/mapping-templates?entityId=...`.
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

func (r MappingUpdateParamsMapping) MarshalJSON() (data []byte, err error) {
	type shadow MappingUpdateParamsMapping
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MappingUpdateParamsMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MappingUpdateParamsMapping](
		"modification", "CamelCase", "DmaIP", "DomainOnly", "DomainPathOnly", "DomainPathUTMs", "DomainUTMs", "FakeDomain", "FakeDomainRealPath", "FakeIP", "FullUrl", "Hash", "HashMD5", "HashedCountry", "HashedDateOfBirth", "HashedGender", "HashedNormalized", "HashedNormalizedNoSpecialChars", "HashedPhone", "HashedState", "HashedZip", "KebabCase", "LowerCase", "None", "Null", "Redacted", "RegionalIP", "SnakeCase", "StartCase", "UpperCase",
	)
}

type MappingReorderParams struct {
	// Mapping ids in their new priority order, low priority index first. All ids must
	// belong to the same parent entity (source or destination).
	Uuids []string `json:"uuids,omitzero" api:"required"`
	paramObj
}

func (r MappingReorderParams) MarshalJSON() (data []byte, err error) {
	type shadow MappingReorderParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MappingReorderParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingTemplatesParams struct {
	// Destination or source id. Required.
	EntityID string `query:"entityId" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [MappingTemplatesParams]'s query parameters as `url.Values`.
func (r MappingTemplatesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
