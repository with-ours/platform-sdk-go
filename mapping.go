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

type MappingListResponse struct {
	ID               string                       `json:"id" api:"required"`
	IsEnabled        bool                         `json:"isEnabled" api:"required"`
	Mappings         []MappingListResponseMapping `json:"mappings" api:"required"`
	DestinationID    string                       `json:"destinationId" api:"nullable"`
	IsDefaultMapping bool                         `json:"isDefaultMapping" api:"nullable"`
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
func (r MappingListResponse) RawJSON() string { return r.JSON.raw }
func (r *MappingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingListResponseMapping struct {
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
func (r MappingNewResponse) RawJSON() string { return r.JSON.raw }
func (r *MappingNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingNewResponseMapping struct {
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
func (r MappingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MappingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingGetResponseMapping struct {
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
func (r MappingUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *MappingUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingUpdateResponseMapping struct {
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
func (r MappingReorderResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *MappingReorderResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MappingReorderResponseEntityMapping struct {
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
func (r MappingReorderResponseEntityMapping) RawJSON() string { return r.JSON.raw }
func (r *MappingReorderResponseEntityMapping) UnmarshalJSON(data []byte) error {
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
	// themselves `MappingLogic` nodes, so trees nest arbitrarily. Example leaf:
	// `{ "condition": { "property": "$event.event", "operator": "Is", "value": "page_view" } }`.
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
// themselves `MappingLogic` nodes, so trees nest arbitrarily. Example leaf:
// `{ "condition": { "property": "$event.event", "operator": "Is", "value": "page_view" } }`.
// Example combinator: `{ "AND": [{ "condition": ... }, { "OR": [...] }] }`.
type MappingNewParamsLogic struct {
	// All child nodes must match. Each child is a `MappingLogic` node.
	And []any `json:"AND,omitzero"`
	// Any child node must match. Each child is a `MappingLogic` node.
	Or        []any                          `json:"OR,omitzero"`
	Condition MappingNewParamsLogicCondition `json:"condition,omitzero"`
	// Negates a single child `MappingLogic` node.
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
	// Any of "Is", "IsNot", "Contains", "DoesNotContain", "StartsWith", "EndsWith",
	// "IsFalsy", "IsTruthy", "IsNull", "IsNotNull", "IsUndefined", "IsNotUndefined",
	// "IsGreaterThan", "IsGreaterThanOrEqual", "IsLessThan", "IsLessThanOrEqual",
	// "IsIn", "IsNotIn", "IsFoundIn", "IsNotFoundIn", "IsTrue", "IsFalse", "IsBefore",
	// "IsAfter", "IsBetween", "IsOnOrBefore", "IsOnOrAfter", "MatchesRegex",
	// "MatchesRegexIgnoreCase", "DoesNotMatchRegex", "DoesNotMatchRegexIgnoreCase".
	Operator string `json:"operator,omitzero" api:"required"`
	Property string `json:"property" api:"required"`
	Value    string `json:"value" api:"required"`
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
	// themselves `MappingLogic` nodes, so trees nest arbitrarily. Example leaf:
	// `{ "condition": { "property": "$event.event", "operator": "Is", "value": "page_view" } }`.
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
// themselves `MappingLogic` nodes, so trees nest arbitrarily. Example leaf:
// `{ "condition": { "property": "$event.event", "operator": "Is", "value": "page_view" } }`.
// Example combinator: `{ "AND": [{ "condition": ... }, { "OR": [...] }] }`.
type MappingUpdateParamsLogic struct {
	// All child nodes must match. Each child is a `MappingLogic` node.
	And []any `json:"AND,omitzero"`
	// Any child node must match. Each child is a `MappingLogic` node.
	Or        []any                             `json:"OR,omitzero"`
	Condition MappingUpdateParamsLogicCondition `json:"condition,omitzero"`
	// Negates a single child `MappingLogic` node.
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
	// Any of "Is", "IsNot", "Contains", "DoesNotContain", "StartsWith", "EndsWith",
	// "IsFalsy", "IsTruthy", "IsNull", "IsNotNull", "IsUndefined", "IsNotUndefined",
	// "IsGreaterThan", "IsGreaterThanOrEqual", "IsLessThan", "IsLessThanOrEqual",
	// "IsIn", "IsNotIn", "IsFoundIn", "IsNotFoundIn", "IsTrue", "IsFalse", "IsBefore",
	// "IsAfter", "IsBetween", "IsOnOrBefore", "IsOnOrAfter", "MatchesRegex",
	// "MatchesRegexIgnoreCase", "DoesNotMatchRegex", "DoesNotMatchRegexIgnoreCase".
	Operator string `json:"operator,omitzero" api:"required"`
	Property string `json:"property" api:"required"`
	Value    string `json:"value" api:"required"`
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
