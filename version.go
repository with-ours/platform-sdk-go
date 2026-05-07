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

// VersionService contains methods and other services that help with interacting
// with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVersionService] method instead.
type VersionService struct {
	Options []option.RequestOption
}

// NewVersionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVersionService(opts ...option.RequestOption) (r VersionService) {
	r = VersionService{}
	r.Options = opts
	return
}

// List versions for this account, newest first. Supports cursor pagination and
// filtering by `isPublished`, `nameContains`, and `notesContains`. Combine filters
// with AND semantics. Requires scope: version:list
func (r *VersionService) List(ctx context.Context, query VersionListParams, opts ...option.RequestOption) (res *pagination.Cursor[VersionListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "rest/v1/versions"
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

// List versions for this account, newest first. Supports cursor pagination and
// filtering by `isPublished`, `nameContains`, and `notesContains`. Combine filters
// with AND semantics. Requires scope: version:list
func (r *VersionService) ListAutoPaging(ctx context.Context, query VersionListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[VersionListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Publish the current draft (i.e. all unpublished entity changes) as a new
// version. Returns the full Version on success. Returns HTTP 409 with the reason
// in the response `error` field when there are no draft changes to publish, when
// another publish is already in flight, or when the action otherwise conflicts
// with current state. To re-publish an existing version, use POST
// /rest/v1/versions/{id}/publish instead. Requires scope: version:publish
func (r *VersionService) New(ctx context.Context, body VersionNewParams, opts ...option.RequestOption) (res *VersionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/versions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Find a single version by ID. Requires scope: version:find
func (r *VersionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *VersionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/versions/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Partially update a version. Only the fields you send are changed. Requires
// scope: version:update
func (r *VersionService) Update(ctx context.Context, id string, body VersionUpdateParams, opts ...option.RequestOption) (res *VersionUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/versions/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Re-publish an existing (previously created) version by ID. Use this to roll back
// to an older snapshot. Returns 409 if the version is already published, was
// created more than 45 days ago, or another publish is already in flight. To
// create-and-publish from current draft state, use POST /rest/v1/versions instead.
// Requires scope: version:publish
func (r *VersionService) Publish(ctx context.Context, id string, opts ...option.RequestOption) (res *VersionPublishResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/versions/%s/publish", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Retrieve the full JSON snapshot captured by a version — every entity
// (destinations, sources, mappings, consent settings, etc.) as it existed when
// this version was published. Sensitive fields (API keys, tokens, secrets) are
// redacted. Useful for IaC export, audit, and backup workflows. Requires scope:
// version:find
func (r *VersionService) Snapshot(ctx context.Context, id string, opts ...option.RequestOption) (res *VersionSnapshotResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/versions/%s/snapshot", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Compare two versions of the account configuration. Returns
// added/removed/modified entities grouped by collection, plus a total `count`.
//
//   - `GET /rest/v1/versions/draft/diff` — compare the current draft (all
//     unpublished entity changes) against the latest published version. Use this to
//     preview what would be included in a `POST /rest/v1/versions` call. (`draft` is
//     a literal path segment — there is no version with that ID; it identifies the
//     comparison target.)
//   - `GET /rest/v1/versions/{id}/diff` — compare that specific version against the
//     latest published version.
//   - `GET /rest/v1/versions/{id}/diff?against={otherId}` — compare two specific
//     versions. `otherId` may also be `draft` to diff a published snapshot against
//     the live draft state. Requires scope: version:find
func (r *VersionService) Diff(ctx context.Context, id VersionDiffParamsID, query VersionDiffParams, opts ...option.RequestOption) (res *VersionDiffResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("rest/v1/versions/%v/diff", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type VersionListResponse struct {
	ID            string  `json:"id" api:"required"`
	CreatedAt     string  `json:"createdAt" api:"required"`
	IsPublished   bool    `json:"isPublished" api:"required"`
	VersionNumber float64 `json:"versionNumber" api:"required"`
	Name          string  `json:"name" api:"nullable"`
	Notes         string  `json:"notes" api:"nullable"`
	// When this version was most recently published. NOT cleared when a newer version
	// is published — `publishedAt` reflects the most recent successful publish of this
	// row, regardless of whether `isPublished` is currently true. Use `isPublished` to
	// determine the current live version.
	PublishedAt string `json:"publishedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		IsPublished   respjson.Field
		VersionNumber respjson.Field
		Name          respjson.Field
		Notes         respjson.Field
		PublishedAt   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionListResponse) RawJSON() string { return r.JSON.raw }
func (r *VersionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionNewResponse struct {
	ID            string  `json:"id" api:"required"`
	CreatedAt     string  `json:"createdAt" api:"required"`
	IsPublished   bool    `json:"isPublished" api:"required"`
	VersionNumber float64 `json:"versionNumber" api:"required"`
	Name          string  `json:"name" api:"nullable"`
	Notes         string  `json:"notes" api:"nullable"`
	// When this version was most recently published. NOT cleared when a newer version
	// is published — `publishedAt` reflects the most recent successful publish of this
	// row, regardless of whether `isPublished` is currently true. Use `isPublished` to
	// determine the current live version.
	PublishedAt string `json:"publishedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		IsPublished   respjson.Field
		VersionNumber respjson.Field
		Name          respjson.Field
		Notes         respjson.Field
		PublishedAt   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *VersionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionGetResponse struct {
	ID            string  `json:"id" api:"required"`
	CreatedAt     string  `json:"createdAt" api:"required"`
	IsPublished   bool    `json:"isPublished" api:"required"`
	VersionNumber float64 `json:"versionNumber" api:"required"`
	Name          string  `json:"name" api:"nullable"`
	Notes         string  `json:"notes" api:"nullable"`
	// When this version was most recently published. NOT cleared when a newer version
	// is published — `publishedAt` reflects the most recent successful publish of this
	// row, regardless of whether `isPublished` is currently true. Use `isPublished` to
	// determine the current live version.
	PublishedAt string `json:"publishedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		IsPublished   respjson.Field
		VersionNumber respjson.Field
		Name          respjson.Field
		Notes         respjson.Field
		PublishedAt   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *VersionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionUpdateResponse struct {
	ID            string  `json:"id" api:"required"`
	CreatedAt     string  `json:"createdAt" api:"required"`
	IsPublished   bool    `json:"isPublished" api:"required"`
	VersionNumber float64 `json:"versionNumber" api:"required"`
	Name          string  `json:"name" api:"nullable"`
	Notes         string  `json:"notes" api:"nullable"`
	// When this version was most recently published. NOT cleared when a newer version
	// is published — `publishedAt` reflects the most recent successful publish of this
	// row, regardless of whether `isPublished` is currently true. Use `isPublished` to
	// determine the current live version.
	PublishedAt string `json:"publishedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		IsPublished   respjson.Field
		VersionNumber respjson.Field
		Name          respjson.Field
		Notes         respjson.Field
		PublishedAt   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *VersionUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionPublishResponse struct {
	ID            string  `json:"id" api:"required"`
	CreatedAt     string  `json:"createdAt" api:"required"`
	IsPublished   bool    `json:"isPublished" api:"required"`
	VersionNumber float64 `json:"versionNumber" api:"required"`
	Name          string  `json:"name" api:"nullable"`
	Notes         string  `json:"notes" api:"nullable"`
	// When this version was most recently published. NOT cleared when a newer version
	// is published — `publishedAt` reflects the most recent successful publish of this
	// row, regardless of whether `isPublished` is currently true. Use `isPublished` to
	// determine the current live version.
	PublishedAt string `json:"publishedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		IsPublished   respjson.Field
		VersionNumber respjson.Field
		Name          respjson.Field
		Notes         respjson.Field
		PublishedAt   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionPublishResponse) RawJSON() string { return r.JSON.raw }
func (r *VersionPublishResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionSnapshotResponse struct {
	ID string `json:"id" api:"required"`
	// The full entity snapshot captured by this version. Keys are entity collection
	// names (destinations, sources, allowedEvents, mappings, consentSettings,
	// replaySettings, globalDispatchCenters, etc.) and values are arrays of the
	// entities as they existed at publish time. Sensitive fields (API keys, tokens,
	// secrets) are redacted. Returns null when the version snapshot has been pruned.
	JsonContent   map[string]any `json:"jsonContent" api:"required"`
	VersionNumber float64        `json:"versionNumber" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		JsonContent   respjson.Field
		VersionNumber respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionSnapshotResponse) RawJSON() string { return r.JSON.raw }
func (r *VersionSnapshotResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponse struct {
	// Total number of differences across all entity collections. Equals the sum of
	// `added + removed + modified` across every collection in `differences`.
	Count       float64                        `json:"count" api:"required"`
	Differences VersionDiffResponseDifferences `json:"differences" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Differences respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponse) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferences struct {
	AllowedEvents            VersionDiffResponseDifferencesAllowedEvents            `json:"allowedEvents" api:"required"`
	ConsentSettings          VersionDiffResponseDifferencesConsentSettings          `json:"consentSettings" api:"required"`
	DataGovernanceEvents     VersionDiffResponseDifferencesDataGovernanceEvents     `json:"dataGovernanceEvents" api:"required"`
	DataGovernanceRules      VersionDiffResponseDifferencesDataGovernanceRules      `json:"dataGovernanceRules" api:"required"`
	Destinations             VersionDiffResponseDifferencesDestinations             `json:"destinations" api:"required"`
	Experiments              VersionDiffResponseDifferencesExperiments              `json:"experiments" api:"required"`
	ExperimentSettings       VersionDiffResponseDifferencesExperimentSettings       `json:"experimentSettings" api:"required"`
	ExperimentVariants       VersionDiffResponseDifferencesExperimentVariants       `json:"experimentVariants" api:"required"`
	ExternalAllowedEventData VersionDiffResponseDifferencesExternalAllowedEventData `json:"externalAllowedEventData" api:"required"`
	GlobalDispatchCenters    VersionDiffResponseDifferencesGlobalDispatchCenters    `json:"globalDispatchCenters" api:"required"`
	Mappings                 VersionDiffResponseDifferencesMappings                 `json:"mappings" api:"required"`
	ReplaySettings           VersionDiffResponseDifferencesReplaySettings           `json:"replaySettings" api:"required"`
	Sources                  VersionDiffResponseDifferencesSources                  `json:"sources" api:"required"`
	TagManagerTags           VersionDiffResponseDifferencesTagManagerTags           `json:"tagManagerTags" api:"required"`
	TagManagerTriggers       VersionDiffResponseDifferencesTagManagerTriggers       `json:"tagManagerTriggers" api:"required"`
	TagManagerVariables      VersionDiffResponseDifferencesTagManagerVariables      `json:"tagManagerVariables" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowedEvents            respjson.Field
		ConsentSettings          respjson.Field
		DataGovernanceEvents     respjson.Field
		DataGovernanceRules      respjson.Field
		Destinations             respjson.Field
		Experiments              respjson.Field
		ExperimentSettings       respjson.Field
		ExperimentVariants       respjson.Field
		ExternalAllowedEventData respjson.Field
		GlobalDispatchCenters    respjson.Field
		Mappings                 respjson.Field
		ReplaySettings           respjson.Field
		Sources                  respjson.Field
		TagManagerTags           respjson.Field
		TagManagerTriggers       respjson.Field
		TagManagerVariables      respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferences) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferences) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesAllowedEvents struct {
	Added []VersionDiffResponseDifferencesAllowedEventsAdded `json:"added" api:"required"`
	// Entities present in both snapshots but with at least one field changed. `old` is
	// the snapshot of the entity in the baseline (latest published); `new` is the
	// snapshot in the comparison target (the draft).
	Modified []VersionDiffResponseDifferencesAllowedEventsModified `json:"modified" api:"required"`
	Removed  []VersionDiffResponseDifferencesAllowedEventsRemoved  `json:"removed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Added       respjson.Field
		Modified    respjson.Field
		Removed     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesAllowedEvents) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesAllowedEvents) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesAllowedEventsAdded struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesAllowedEventsAdded) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesAllowedEventsAdded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesAllowedEventsModified struct {
	New VersionDiffResponseDifferencesAllowedEventsModifiedNew `json:"new" api:"required"`
	Old VersionDiffResponseDifferencesAllowedEventsModifiedOld `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesAllowedEventsModified) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesAllowedEventsModified) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesAllowedEventsModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesAllowedEventsModifiedNew) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesAllowedEventsModifiedNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesAllowedEventsModifiedOld struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesAllowedEventsModifiedOld) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesAllowedEventsModifiedOld) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesAllowedEventsRemoved struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesAllowedEventsRemoved) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesAllowedEventsRemoved) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesConsentSettings struct {
	Added []VersionDiffResponseDifferencesConsentSettingsAdded `json:"added" api:"required"`
	// Entities present in both snapshots but with at least one field changed. `old` is
	// the snapshot of the entity in the baseline (latest published); `new` is the
	// snapshot in the comparison target (the draft).
	Modified []VersionDiffResponseDifferencesConsentSettingsModified `json:"modified" api:"required"`
	Removed  []VersionDiffResponseDifferencesConsentSettingsRemoved  `json:"removed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Added       respjson.Field
		Modified    respjson.Field
		Removed     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesConsentSettings) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesConsentSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesConsentSettingsAdded struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesConsentSettingsAdded) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesConsentSettingsAdded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesConsentSettingsModified struct {
	New VersionDiffResponseDifferencesConsentSettingsModifiedNew `json:"new" api:"required"`
	Old VersionDiffResponseDifferencesConsentSettingsModifiedOld `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesConsentSettingsModified) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesConsentSettingsModified) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesConsentSettingsModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesConsentSettingsModifiedNew) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesConsentSettingsModifiedNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesConsentSettingsModifiedOld struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesConsentSettingsModifiedOld) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesConsentSettingsModifiedOld) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesConsentSettingsRemoved struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesConsentSettingsRemoved) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesConsentSettingsRemoved) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesDataGovernanceEvents struct {
	Added []VersionDiffResponseDifferencesDataGovernanceEventsAdded `json:"added" api:"required"`
	// Entities present in both snapshots but with at least one field changed. `old` is
	// the snapshot of the entity in the baseline (latest published); `new` is the
	// snapshot in the comparison target (the draft).
	Modified []VersionDiffResponseDifferencesDataGovernanceEventsModified `json:"modified" api:"required"`
	Removed  []VersionDiffResponseDifferencesDataGovernanceEventsRemoved  `json:"removed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Added       respjson.Field
		Modified    respjson.Field
		Removed     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesDataGovernanceEvents) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesDataGovernanceEvents) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesDataGovernanceEventsAdded struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesDataGovernanceEventsAdded) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesDataGovernanceEventsAdded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesDataGovernanceEventsModified struct {
	New VersionDiffResponseDifferencesDataGovernanceEventsModifiedNew `json:"new" api:"required"`
	Old VersionDiffResponseDifferencesDataGovernanceEventsModifiedOld `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesDataGovernanceEventsModified) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesDataGovernanceEventsModified) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesDataGovernanceEventsModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesDataGovernanceEventsModifiedNew) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesDataGovernanceEventsModifiedNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesDataGovernanceEventsModifiedOld struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesDataGovernanceEventsModifiedOld) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesDataGovernanceEventsModifiedOld) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesDataGovernanceEventsRemoved struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesDataGovernanceEventsRemoved) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesDataGovernanceEventsRemoved) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesDataGovernanceRules struct {
	Added []VersionDiffResponseDifferencesDataGovernanceRulesAdded `json:"added" api:"required"`
	// Entities present in both snapshots but with at least one field changed. `old` is
	// the snapshot of the entity in the baseline (latest published); `new` is the
	// snapshot in the comparison target (the draft).
	Modified []VersionDiffResponseDifferencesDataGovernanceRulesModified `json:"modified" api:"required"`
	Removed  []VersionDiffResponseDifferencesDataGovernanceRulesRemoved  `json:"removed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Added       respjson.Field
		Modified    respjson.Field
		Removed     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesDataGovernanceRules) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesDataGovernanceRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesDataGovernanceRulesAdded struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesDataGovernanceRulesAdded) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesDataGovernanceRulesAdded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesDataGovernanceRulesModified struct {
	New VersionDiffResponseDifferencesDataGovernanceRulesModifiedNew `json:"new" api:"required"`
	Old VersionDiffResponseDifferencesDataGovernanceRulesModifiedOld `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesDataGovernanceRulesModified) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesDataGovernanceRulesModified) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesDataGovernanceRulesModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesDataGovernanceRulesModifiedNew) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesDataGovernanceRulesModifiedNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesDataGovernanceRulesModifiedOld struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesDataGovernanceRulesModifiedOld) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesDataGovernanceRulesModifiedOld) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesDataGovernanceRulesRemoved struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesDataGovernanceRulesRemoved) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesDataGovernanceRulesRemoved) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesDestinations struct {
	Added []VersionDiffResponseDifferencesDestinationsAdded `json:"added" api:"required"`
	// Entities present in both snapshots but with at least one field changed. `old` is
	// the snapshot of the entity in the baseline (latest published); `new` is the
	// snapshot in the comparison target (the draft).
	Modified []VersionDiffResponseDifferencesDestinationsModified `json:"modified" api:"required"`
	Removed  []VersionDiffResponseDifferencesDestinationsRemoved  `json:"removed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Added       respjson.Field
		Modified    respjson.Field
		Removed     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesDestinations) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesDestinations) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesDestinationsAdded struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesDestinationsAdded) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesDestinationsAdded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesDestinationsModified struct {
	New VersionDiffResponseDifferencesDestinationsModifiedNew `json:"new" api:"required"`
	Old VersionDiffResponseDifferencesDestinationsModifiedOld `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesDestinationsModified) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesDestinationsModified) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesDestinationsModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesDestinationsModifiedNew) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesDestinationsModifiedNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesDestinationsModifiedOld struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesDestinationsModifiedOld) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesDestinationsModifiedOld) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesDestinationsRemoved struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesDestinationsRemoved) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesDestinationsRemoved) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesExperiments struct {
	Added []VersionDiffResponseDifferencesExperimentsAdded `json:"added" api:"required"`
	// Entities present in both snapshots but with at least one field changed. `old` is
	// the snapshot of the entity in the baseline (latest published); `new` is the
	// snapshot in the comparison target (the draft).
	Modified []VersionDiffResponseDifferencesExperimentsModified `json:"modified" api:"required"`
	Removed  []VersionDiffResponseDifferencesExperimentsRemoved  `json:"removed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Added       respjson.Field
		Modified    respjson.Field
		Removed     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesExperiments) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesExperiments) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesExperimentsAdded struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesExperimentsAdded) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesExperimentsAdded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesExperimentsModified struct {
	New VersionDiffResponseDifferencesExperimentsModifiedNew `json:"new" api:"required"`
	Old VersionDiffResponseDifferencesExperimentsModifiedOld `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesExperimentsModified) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesExperimentsModified) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesExperimentsModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesExperimentsModifiedNew) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesExperimentsModifiedNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesExperimentsModifiedOld struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesExperimentsModifiedOld) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesExperimentsModifiedOld) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesExperimentsRemoved struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesExperimentsRemoved) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesExperimentsRemoved) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesExperimentSettings struct {
	Added []VersionDiffResponseDifferencesExperimentSettingsAdded `json:"added" api:"required"`
	// Entities present in both snapshots but with at least one field changed. `old` is
	// the snapshot of the entity in the baseline (latest published); `new` is the
	// snapshot in the comparison target (the draft).
	Modified []VersionDiffResponseDifferencesExperimentSettingsModified `json:"modified" api:"required"`
	Removed  []VersionDiffResponseDifferencesExperimentSettingsRemoved  `json:"removed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Added       respjson.Field
		Modified    respjson.Field
		Removed     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesExperimentSettings) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesExperimentSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesExperimentSettingsAdded struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesExperimentSettingsAdded) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesExperimentSettingsAdded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesExperimentSettingsModified struct {
	New VersionDiffResponseDifferencesExperimentSettingsModifiedNew `json:"new" api:"required"`
	Old VersionDiffResponseDifferencesExperimentSettingsModifiedOld `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesExperimentSettingsModified) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesExperimentSettingsModified) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesExperimentSettingsModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesExperimentSettingsModifiedNew) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesExperimentSettingsModifiedNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesExperimentSettingsModifiedOld struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesExperimentSettingsModifiedOld) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesExperimentSettingsModifiedOld) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesExperimentSettingsRemoved struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesExperimentSettingsRemoved) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesExperimentSettingsRemoved) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesExperimentVariants struct {
	Added []VersionDiffResponseDifferencesExperimentVariantsAdded `json:"added" api:"required"`
	// Entities present in both snapshots but with at least one field changed. `old` is
	// the snapshot of the entity in the baseline (latest published); `new` is the
	// snapshot in the comparison target (the draft).
	Modified []VersionDiffResponseDifferencesExperimentVariantsModified `json:"modified" api:"required"`
	Removed  []VersionDiffResponseDifferencesExperimentVariantsRemoved  `json:"removed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Added       respjson.Field
		Modified    respjson.Field
		Removed     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesExperimentVariants) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesExperimentVariants) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesExperimentVariantsAdded struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesExperimentVariantsAdded) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesExperimentVariantsAdded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesExperimentVariantsModified struct {
	New VersionDiffResponseDifferencesExperimentVariantsModifiedNew `json:"new" api:"required"`
	Old VersionDiffResponseDifferencesExperimentVariantsModifiedOld `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesExperimentVariantsModified) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesExperimentVariantsModified) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesExperimentVariantsModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesExperimentVariantsModifiedNew) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesExperimentVariantsModifiedNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesExperimentVariantsModifiedOld struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesExperimentVariantsModifiedOld) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesExperimentVariantsModifiedOld) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesExperimentVariantsRemoved struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesExperimentVariantsRemoved) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesExperimentVariantsRemoved) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesExternalAllowedEventData struct {
	Added []VersionDiffResponseDifferencesExternalAllowedEventDataAdded `json:"added" api:"required"`
	// Entities present in both snapshots but with at least one field changed. `old` is
	// the snapshot of the entity in the baseline (latest published); `new` is the
	// snapshot in the comparison target (the draft).
	Modified []VersionDiffResponseDifferencesExternalAllowedEventDataModified `json:"modified" api:"required"`
	Removed  []VersionDiffResponseDifferencesExternalAllowedEventDataRemoved  `json:"removed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Added       respjson.Field
		Modified    respjson.Field
		Removed     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesExternalAllowedEventData) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesExternalAllowedEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesExternalAllowedEventDataAdded struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesExternalAllowedEventDataAdded) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesExternalAllowedEventDataAdded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesExternalAllowedEventDataModified struct {
	New VersionDiffResponseDifferencesExternalAllowedEventDataModifiedNew `json:"new" api:"required"`
	Old VersionDiffResponseDifferencesExternalAllowedEventDataModifiedOld `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesExternalAllowedEventDataModified) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesExternalAllowedEventDataModified) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesExternalAllowedEventDataModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesExternalAllowedEventDataModifiedNew) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesExternalAllowedEventDataModifiedNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesExternalAllowedEventDataModifiedOld struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesExternalAllowedEventDataModifiedOld) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesExternalAllowedEventDataModifiedOld) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesExternalAllowedEventDataRemoved struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesExternalAllowedEventDataRemoved) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesExternalAllowedEventDataRemoved) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesGlobalDispatchCenters struct {
	Added []VersionDiffResponseDifferencesGlobalDispatchCentersAdded `json:"added" api:"required"`
	// Entities present in both snapshots but with at least one field changed. `old` is
	// the snapshot of the entity in the baseline (latest published); `new` is the
	// snapshot in the comparison target (the draft).
	Modified []VersionDiffResponseDifferencesGlobalDispatchCentersModified `json:"modified" api:"required"`
	Removed  []VersionDiffResponseDifferencesGlobalDispatchCentersRemoved  `json:"removed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Added       respjson.Field
		Modified    respjson.Field
		Removed     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesGlobalDispatchCenters) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesGlobalDispatchCenters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesGlobalDispatchCentersAdded struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesGlobalDispatchCentersAdded) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesGlobalDispatchCentersAdded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesGlobalDispatchCentersModified struct {
	New VersionDiffResponseDifferencesGlobalDispatchCentersModifiedNew `json:"new" api:"required"`
	Old VersionDiffResponseDifferencesGlobalDispatchCentersModifiedOld `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesGlobalDispatchCentersModified) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesGlobalDispatchCentersModified) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesGlobalDispatchCentersModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesGlobalDispatchCentersModifiedNew) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesGlobalDispatchCentersModifiedNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesGlobalDispatchCentersModifiedOld struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesGlobalDispatchCentersModifiedOld) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesGlobalDispatchCentersModifiedOld) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesGlobalDispatchCentersRemoved struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesGlobalDispatchCentersRemoved) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesGlobalDispatchCentersRemoved) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesMappings struct {
	Added []VersionDiffResponseDifferencesMappingsAdded `json:"added" api:"required"`
	// Entities present in both snapshots but with at least one field changed. `old` is
	// the snapshot of the entity in the baseline (latest published); `new` is the
	// snapshot in the comparison target (the draft).
	Modified []VersionDiffResponseDifferencesMappingsModified `json:"modified" api:"required"`
	Removed  []VersionDiffResponseDifferencesMappingsRemoved  `json:"removed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Added       respjson.Field
		Modified    respjson.Field
		Removed     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesMappings) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesMappings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesMappingsAdded struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesMappingsAdded) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesMappingsAdded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesMappingsModified struct {
	New VersionDiffResponseDifferencesMappingsModifiedNew `json:"new" api:"required"`
	Old VersionDiffResponseDifferencesMappingsModifiedOld `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesMappingsModified) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesMappingsModified) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesMappingsModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesMappingsModifiedNew) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesMappingsModifiedNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesMappingsModifiedOld struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesMappingsModifiedOld) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesMappingsModifiedOld) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesMappingsRemoved struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesMappingsRemoved) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesMappingsRemoved) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesReplaySettings struct {
	Added []VersionDiffResponseDifferencesReplaySettingsAdded `json:"added" api:"required"`
	// Entities present in both snapshots but with at least one field changed. `old` is
	// the snapshot of the entity in the baseline (latest published); `new` is the
	// snapshot in the comparison target (the draft).
	Modified []VersionDiffResponseDifferencesReplaySettingsModified `json:"modified" api:"required"`
	Removed  []VersionDiffResponseDifferencesReplaySettingsRemoved  `json:"removed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Added       respjson.Field
		Modified    respjson.Field
		Removed     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesReplaySettings) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesReplaySettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesReplaySettingsAdded struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesReplaySettingsAdded) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesReplaySettingsAdded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesReplaySettingsModified struct {
	New VersionDiffResponseDifferencesReplaySettingsModifiedNew `json:"new" api:"required"`
	Old VersionDiffResponseDifferencesReplaySettingsModifiedOld `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesReplaySettingsModified) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesReplaySettingsModified) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesReplaySettingsModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesReplaySettingsModifiedNew) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesReplaySettingsModifiedNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesReplaySettingsModifiedOld struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesReplaySettingsModifiedOld) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesReplaySettingsModifiedOld) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesReplaySettingsRemoved struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesReplaySettingsRemoved) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesReplaySettingsRemoved) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesSources struct {
	Added []VersionDiffResponseDifferencesSourcesAdded `json:"added" api:"required"`
	// Entities present in both snapshots but with at least one field changed. `old` is
	// the snapshot of the entity in the baseline (latest published); `new` is the
	// snapshot in the comparison target (the draft).
	Modified []VersionDiffResponseDifferencesSourcesModified `json:"modified" api:"required"`
	Removed  []VersionDiffResponseDifferencesSourcesRemoved  `json:"removed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Added       respjson.Field
		Modified    respjson.Field
		Removed     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesSources) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesSources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesSourcesAdded struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesSourcesAdded) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesSourcesAdded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesSourcesModified struct {
	New VersionDiffResponseDifferencesSourcesModifiedNew `json:"new" api:"required"`
	Old VersionDiffResponseDifferencesSourcesModifiedOld `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesSourcesModified) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesSourcesModified) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesSourcesModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesSourcesModifiedNew) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesSourcesModifiedNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesSourcesModifiedOld struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesSourcesModifiedOld) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesSourcesModifiedOld) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesSourcesRemoved struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesSourcesRemoved) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesSourcesRemoved) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesTagManagerTags struct {
	Added []VersionDiffResponseDifferencesTagManagerTagsAdded `json:"added" api:"required"`
	// Entities present in both snapshots but with at least one field changed. `old` is
	// the snapshot of the entity in the baseline (latest published); `new` is the
	// snapshot in the comparison target (the draft).
	Modified []VersionDiffResponseDifferencesTagManagerTagsModified `json:"modified" api:"required"`
	Removed  []VersionDiffResponseDifferencesTagManagerTagsRemoved  `json:"removed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Added       respjson.Field
		Modified    respjson.Field
		Removed     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesTagManagerTags) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesTagManagerTags) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesTagManagerTagsAdded struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesTagManagerTagsAdded) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesTagManagerTagsAdded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesTagManagerTagsModified struct {
	New VersionDiffResponseDifferencesTagManagerTagsModifiedNew `json:"new" api:"required"`
	Old VersionDiffResponseDifferencesTagManagerTagsModifiedOld `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesTagManagerTagsModified) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesTagManagerTagsModified) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesTagManagerTagsModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesTagManagerTagsModifiedNew) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesTagManagerTagsModifiedNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesTagManagerTagsModifiedOld struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesTagManagerTagsModifiedOld) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesTagManagerTagsModifiedOld) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesTagManagerTagsRemoved struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesTagManagerTagsRemoved) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesTagManagerTagsRemoved) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesTagManagerTriggers struct {
	Added []VersionDiffResponseDifferencesTagManagerTriggersAdded `json:"added" api:"required"`
	// Entities present in both snapshots but with at least one field changed. `old` is
	// the snapshot of the entity in the baseline (latest published); `new` is the
	// snapshot in the comparison target (the draft).
	Modified []VersionDiffResponseDifferencesTagManagerTriggersModified `json:"modified" api:"required"`
	Removed  []VersionDiffResponseDifferencesTagManagerTriggersRemoved  `json:"removed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Added       respjson.Field
		Modified    respjson.Field
		Removed     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesTagManagerTriggers) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesTagManagerTriggers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesTagManagerTriggersAdded struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesTagManagerTriggersAdded) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesTagManagerTriggersAdded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesTagManagerTriggersModified struct {
	New VersionDiffResponseDifferencesTagManagerTriggersModifiedNew `json:"new" api:"required"`
	Old VersionDiffResponseDifferencesTagManagerTriggersModifiedOld `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesTagManagerTriggersModified) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesTagManagerTriggersModified) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesTagManagerTriggersModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesTagManagerTriggersModifiedNew) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesTagManagerTriggersModifiedNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesTagManagerTriggersModifiedOld struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesTagManagerTriggersModifiedOld) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesTagManagerTriggersModifiedOld) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesTagManagerTriggersRemoved struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesTagManagerTriggersRemoved) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesTagManagerTriggersRemoved) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesTagManagerVariables struct {
	Added []VersionDiffResponseDifferencesTagManagerVariablesAdded `json:"added" api:"required"`
	// Entities present in both snapshots but with at least one field changed. `old` is
	// the snapshot of the entity in the baseline (latest published); `new` is the
	// snapshot in the comparison target (the draft).
	Modified []VersionDiffResponseDifferencesTagManagerVariablesModified `json:"modified" api:"required"`
	Removed  []VersionDiffResponseDifferencesTagManagerVariablesRemoved  `json:"removed" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Added       respjson.Field
		Modified    respjson.Field
		Removed     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesTagManagerVariables) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesTagManagerVariables) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesTagManagerVariablesAdded struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesTagManagerVariablesAdded) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesTagManagerVariablesAdded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesTagManagerVariablesModified struct {
	New VersionDiffResponseDifferencesTagManagerVariablesModifiedNew `json:"new" api:"required"`
	Old VersionDiffResponseDifferencesTagManagerVariablesModifiedOld `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesTagManagerVariablesModified) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesTagManagerVariablesModified) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesTagManagerVariablesModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesTagManagerVariablesModifiedNew) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesTagManagerVariablesModifiedNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesTagManagerVariablesModifiedOld struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesTagManagerVariablesModifiedOld) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesTagManagerVariablesModifiedOld) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesTagManagerVariablesRemoved struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Use the `id` to
	// fetch full detail.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesTagManagerVariablesRemoved) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesTagManagerVariablesRemoved) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionListParams struct {
	// Maximum number of items to return. Defaults to 25; values below 1 are clamped to
	// 1 and values above 100 are clamped to 100.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Opaque pagination cursor from pagination.nextCursor in the previous response. Do
	// not decode or modify it. Malformed cursors return 400 Bad Request.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Case-insensitive substring match on the version name.
	NameContains param.Opt[string] `query:"nameContains,omitzero" json:"-"`
	// Case-insensitive substring match on the version notes.
	NotesContains param.Opt[string] `query:"notesContains,omitzero" json:"-"`
	// Filter to only published or unpublished versions.
	//
	// Any of "true", "false".
	IsPublished VersionListParamsIsPublished `query:"isPublished,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VersionListParams]'s query parameters as `url.Values`.
func (r VersionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter to only published or unpublished versions.
type VersionListParamsIsPublished string

const (
	VersionListParamsIsPublishedTrue  VersionListParamsIsPublished = "true"
	VersionListParamsIsPublishedFalse VersionListParamsIsPublished = "false"
)

type VersionNewParams struct {
	Name  param.Opt[string] `json:"name,omitzero"`
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Cherry-pick: allowed event ids (slug-like strings) to include from the draft.
	// Omit or send `[]` to include all draft changes in this collection.
	IncludeAllowedEvents []string `json:"includeAllowedEvents,omitzero"`
	// Cherry-pick: consent settings ids to include from the draft. Omit or send `[]`
	// to include all draft changes in this collection.
	IncludeConsentSettings []string `json:"includeConsentSettings,omitzero"`
	// Cherry-pick: data governance event UUIDs to include from the draft. Omit or send
	// `[]` to include all draft changes in this collection.
	IncludeDataGovernanceEvents []string `json:"includeDataGovernanceEvents,omitzero"`
	// Cherry-pick: data governance rule UUIDs to include from the draft. Omit or send
	// `[]` to include all draft changes in this collection.
	IncludeDataGovernanceRules []string `json:"includeDataGovernanceRules,omitzero"`
	// Cherry-pick: destination UUIDs to include from the draft. Omit or send `[]` to
	// include all draft changes in this collection.
	IncludeDestinations []string `json:"includeDestinations,omitzero"`
	// Cherry-pick: experiment UUIDs to include from the draft. Omit or send `[]` to
	// include all draft changes in this collection.
	IncludeExperiments []string `json:"includeExperiments,omitzero"`
	// Cherry-pick: experiment settings UUIDs to include from the draft. Omit or send
	// `[]` to include all draft changes in this collection.
	IncludeExperimentSettings []string `json:"includeExperimentSettings,omitzero"`
	// Cherry-pick: experiment variant UUIDs to include from the draft. Omit or send
	// `[]` to include all draft changes in this collection.
	IncludeExperimentVariants []string `json:"includeExperimentVariants,omitzero"`
	// Cherry-pick: external allowed event data ids to include from the draft. Omit or
	// send `[]` to include all draft changes in this collection.
	IncludeExternalAllowedEventData []string `json:"includeExternalAllowedEventData,omitzero"`
	// Cherry-pick: global dispatch center ids to include from the draft. Omit or send
	// `[]` to include all draft changes in this collection.
	IncludeGlobalDispatchCenters []string `json:"includeGlobalDispatchCenters,omitzero"`
	// Cherry-pick: mapping ids to include from the draft. Omit or send `[]` to include
	// all draft changes in this collection.
	IncludeMappings []string `json:"includeMappings,omitzero"`
	// Cherry-pick: replay settings ids to include from the draft. Omit or send `[]` to
	// include all draft changes in this collection.
	IncludeReplaySettings []string `json:"includeReplaySettings,omitzero"`
	// Cherry-pick: source UUIDs to include from the draft. Omit or send `[]` to
	// include all draft changes in this collection.
	IncludeSources []string `json:"includeSources,omitzero"`
	// Cherry-pick: tag manager tag UUIDs to include from the draft. Omit or send `[]`
	// to include all draft changes in this collection.
	IncludeTagManagerTags []string `json:"includeTagManagerTags,omitzero"`
	// Cherry-pick: tag manager trigger UUIDs to include from the draft. Omit or send
	// `[]` to include all draft changes in this collection.
	IncludeTagManagerTriggers []string `json:"includeTagManagerTriggers,omitzero"`
	// Cherry-pick: tag manager variable UUIDs to include from the draft. Omit or send
	// `[]` to include all draft changes in this collection.
	IncludeTagManagerVariables []string `json:"includeTagManagerVariables,omitzero"`
	paramObj
}

func (r VersionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VersionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VersionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionUpdateParams struct {
	Name  param.Opt[string] `json:"name,omitzero"`
	Notes param.Opt[string] `json:"notes,omitzero"`
	paramObj
}

func (r VersionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow VersionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VersionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffParams struct {
	// Baseline version id to compare the path version against. Omit for the latest
	// published version. Pass a version UUID to compute a version-vs-version diff.
	Against param.Opt[string] `query:"against,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [VersionDiffParams]'s query parameters as `url.Values`.
func (r VersionDiffParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VersionDiffParamsID string

const (
	VersionDiffParamsIDDraft VersionDiffParamsID = "draft"
)
