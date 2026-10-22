// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/with-ours/platform-sdk-go/v2/internal/apijson"
	"github.com/with-ours/platform-sdk-go/v2/internal/apiquery"
	"github.com/with-ours/platform-sdk-go/v2/internal/requestconfig"
	"github.com/with-ours/platform-sdk-go/v2/option"
	"github.com/with-ours/platform-sdk-go/v2/packages/pagination"
	"github.com/with-ours/platform-sdk-go/v2/packages/param"
	"github.com/with-ours/platform-sdk-go/v2/packages/respjson"
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
// with AND semantics. Requires API-key scope or current OAuth user permission:
// version:list
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
// with AND semantics. Requires API-key scope or current OAuth user permission:
// version:list
func (r *VersionService) ListAutoPaging(ctx context.Context, query VersionListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[VersionListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, query, opts...))
}

// Publish the current draft of non-experiment entities (destinations, mappings,
// experiment settings, governance rules, etc.) as a new version. Newly created or
// modified DRAFT experiments and experiment variants are NOT shipped here — call
// `POST /rest/v1/experiments/{id}/start` instead, which atomically publishes the
// experiment and its variants by default (`publishAfterStart: true`). Returns the
// full Version on success. Returns HTTP 409 with the reason in the response
// `error` field when there are no draft changes to publish, when another publish
// is already in flight, or when the action otherwise conflicts with current state.
// To re-publish an existing version, use POST /rest/v1/versions/{id}/publish
// instead.
//
// Supports cherry-picking via the `include*` body fields, each an entity-id
// allowlist for one collection. Omit every `include*` field (or send each as `[]`)
// to publish all draft changes. Once any `include*` field lists ids, only the
// listed changes publish, and every collection you omit or send as `[]` keeps its
// published state. Requires API-key scope or current OAuth user permission:
// version:publish
func (r *VersionService) New(ctx context.Context, body VersionNewParams, opts ...option.RequestOption) (res *VersionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/versions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Find a single version by ID. Requires API-key scope or current OAuth user
// permission: version:find
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
// API-key scope or current OAuth user permission: version:update
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
// Requires API-key scope or current OAuth user permission: version:publish
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
// redacted. Useful for IaC export, audit, and backup workflows. Requires API-key
// scope or current OAuth user permission: version:find
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
// Each modified entity includes `changes` with dot paths, redacted old/new values,
// whole arrays, and null for absent values.
//
// Draft experiments and experiment variants do NOT appear in any diff — only
// experiments that are already running (or paused/completed) snapshot into
// versions. Use the experiment lifecycle endpoints (`/start`, `/stop`, etc.) to
// ship experiment changes; use these diff endpoints only for non-experiment
// entities (destinations, mappings, experiment settings, governance rules, etc.).
//
//   - `GET /rest/v1/versions/draft/diff` — compare the current draft of
//     non-experiment entities against the latest published version. Use this to
//     preview what would be included in a `POST /rest/v1/versions` call. (`draft` is
//     a literal path segment — there is no version with that ID; it identifies the
//     comparison target.)
//   - `GET /rest/v1/versions/{id}/diff` — compare that specific version against the
//     latest published version.
//   - `GET /rest/v1/versions/{id}/diff?against={otherId}` — compare two specific
//     versions. `otherId` must be a published version UUID. Use `draft` as the path
//     id to compare the live draft against that version. Requires API-key scope or
//     current OAuth user permission: version:find
func (r *VersionService) Diff(ctx context.Context, id VersionDiffParamsID, query VersionDiffParams, opts ...option.RequestOption) (res *VersionDiffResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("rest/v1/versions/%v/diff", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Check whether one entity has unpublished changes, without fetching the org-wide
// draft diff.
// `GET /rest/v1/versions/draft/status?collection={collection}&entityId={id}`
// returns `status: "added" | "modified" | "removed" | null` for that entity, with
// the same semantics as `GET /rest/v1/versions/draft/diff` — it shares the diff's
// field blacklist, so passive fields (timestamps, sync bookkeeping) do not count
// as a change.
//
// `status` is `null` for every outcome other than a real pending change: the
// entity already matches the published version, the id is unknown, the id belongs
// to another account, or the account has never published anything. Use
// `GET /rest/v1/versions/draft/diff` when you need the full change set instead of
// one entity. (`draft` is a literal path segment identifying the live draft as the
// target, which is why the entity is passed as `entityId` rather than `id`.)
// Requires API-key scope or current OAuth user permission: version:find
func (r *VersionService) Status(ctx context.Context, id VersionStatusParamsID, query VersionStatusParams, opts ...option.RequestOption) (res *VersionStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("rest/v1/versions/%v/status", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Revert one or more pending draft changes back to the latest published version
// (the "abandon a change" action). Send a body listing the entities to revert;
// each is handled independently — a modified entity is restored to its published
// values, a deleted entity is recreated, and a newly-added entity is discarded. An
// entity that already matches the published version is left untouched and is not
// counted (a no-op). An entity present in neither the draft nor the published
// version is reported in `notFoundCount`. `remainingChanges` is the total number
// of pending draft changes left after this call, which may stay above zero when
// other unreverted changes remain.
//
// The `experiments` and `experimentVariants` collections are not supported here —
// use the experiment lifecycle endpoints (`/start`, `/stop`) for those — but
// `experimentSettings` is supported. To discard every pending change at once, use
// `POST /rest/v1/versions/draft/abandon`. (`draft` is a literal path segment
// identifying the live draft as the target.) Requires API-key scope or current
// OAuth user permission: version:publish
func (r *VersionService) Revert(ctx context.Context, id VersionRevertParamsID, body VersionRevertParams, opts ...option.RequestOption) (res *VersionRevertResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("rest/v1/versions/%v/revert", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Abandon ALL pending draft changes, resetting every supported entity collection
// back to the latest published version. Running experiments and experiment
// variants are not affected (ship those via the experiment lifecycle endpoints),
// but experiment settings are reset. Returns HTTP 409 when there is no published
// version to revert to, or when another publish or abandon is already in flight.
// To revert only specific changes, use `POST /rest/v1/versions/draft/revert`.
// (`draft` is a literal path segment identifying the live draft as the target.)
// Requires API-key scope or current OAuth user permission: version:publish
func (r *VersionService) Abandon(ctx context.Context, id VersionAbandonParamsID, opts ...option.RequestOption) (res *VersionAbandonResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("rest/v1/versions/%v/abandon", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
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
	// Identifier of the user who most recently published this version.
	PublishedByUserID string `json:"publishedByUserId" api:"nullable"`
	// Display name of the user who most recently published this version.
	PublishedByUserName string `json:"publishedByUserName" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		CreatedAt           respjson.Field
		IsPublished         respjson.Field
		VersionNumber       respjson.Field
		Name                respjson.Field
		Notes               respjson.Field
		PublishedAt         respjson.Field
		PublishedByUserID   respjson.Field
		PublishedByUserName respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
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
	// Identifier of the user who most recently published this version.
	PublishedByUserID string `json:"publishedByUserId" api:"nullable"`
	// Display name of the user who most recently published this version.
	PublishedByUserName string `json:"publishedByUserName" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		CreatedAt           respjson.Field
		IsPublished         respjson.Field
		VersionNumber       respjson.Field
		Name                respjson.Field
		Notes               respjson.Field
		PublishedAt         respjson.Field
		PublishedByUserID   respjson.Field
		PublishedByUserName respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
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
	// Identifier of the user who most recently published this version.
	PublishedByUserID string `json:"publishedByUserId" api:"nullable"`
	// Display name of the user who most recently published this version.
	PublishedByUserName string `json:"publishedByUserName" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		CreatedAt           respjson.Field
		IsPublished         respjson.Field
		VersionNumber       respjson.Field
		Name                respjson.Field
		Notes               respjson.Field
		PublishedAt         respjson.Field
		PublishedByUserID   respjson.Field
		PublishedByUserName respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
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
	// Identifier of the user who most recently published this version.
	PublishedByUserID string `json:"publishedByUserId" api:"nullable"`
	// Display name of the user who most recently published this version.
	PublishedByUserName string `json:"publishedByUserName" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		CreatedAt           respjson.Field
		IsPublished         respjson.Field
		VersionNumber       respjson.Field
		Name                respjson.Field
		Notes               respjson.Field
		PublishedAt         respjson.Field
		PublishedByUserID   respjson.Field
		PublishedByUserName respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
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
	// Identifier of the user who most recently published this version.
	PublishedByUserID string `json:"publishedByUserId" api:"nullable"`
	// Display name of the user who most recently published this version.
	PublishedByUserName string `json:"publishedByUserName" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		CreatedAt           respjson.Field
		IsPublished         respjson.Field
		VersionNumber       respjson.Field
		Name                respjson.Field
		Notes               respjson.Field
		PublishedAt         respjson.Field
		PublishedByUserID   respjson.Field
		PublishedByUserName respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
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
	ShortLinks               VersionDiffResponseDifferencesShortLinks               `json:"shortLinks" api:"required"`
	Sources                  VersionDiffResponseDifferencesSources                  `json:"sources" api:"required"`
	TagManagers              VersionDiffResponseDifferencesTagManagers              `json:"tagManagers" api:"required"`
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
		ShortLinks               respjson.Field
		Sources                  respjson.Field
		TagManagers              respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// Changed fields with credentials redacted. Nested objects use dot paths; arrays
	// are whole values. Missing values are null.
	Changes []VersionDiffResponseDifferencesAllowedEventsModifiedChange `json:"changes" api:"required"`
	New     VersionDiffResponseDifferencesAllowedEventsModifiedNew      `json:"new" api:"required"`
	Old     VersionDiffResponseDifferencesAllowedEventsModifiedOld      `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Changes     respjson.Field
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

type VersionDiffResponseDifferencesAllowedEventsModifiedChange struct {
	New  VersionDiffResponseDifferencesAllowedEventsModifiedChangeNewUnion `json:"new" api:"required"`
	Old  VersionDiffResponseDifferencesAllowedEventsModifiedChangeOldUnion `json:"old" api:"required"`
	Path string                                                            `json:"path" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesAllowedEventsModifiedChange) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesAllowedEventsModifiedChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesAllowedEventsModifiedChangeNewUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesAllowedEventsModifiedChangeNewArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesAllowedEventsModifiedChangeNewArray
// OfVersionDiffResponseDifferencesAllowedEventsModifiedChangeNewMapItem]
type VersionDiffResponseDifferencesAllowedEventsModifiedChangeNewUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesAllowedEventsModifiedChangeNewArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesAllowedEventsModifiedChangeNewArray []VersionDiffResponseDifferencesAllowedEventsModifiedChangeNewArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesAllowedEventsModifiedChangeNewMapItem any `json:",inline"`
	JSON                                                                  struct {
		OfString                                                              respjson.Field
		OfFloat                                                               respjson.Field
		OfBool                                                                respjson.Field
		OfVersionDiffResponseDifferencesAllowedEventsModifiedChangeNewArray   respjson.Field
		OfVersionDiffResponseDifferencesAllowedEventsModifiedChangeNewMapItem respjson.Field
		raw                                                                   string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesAllowedEventsModifiedChangeNewUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesAllowedEventsModifiedChangeNewUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesAllowedEventsModifiedChangeNewUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesAllowedEventsModifiedChangeNewUnion) AsVersionDiffResponseDifferencesAllowedEventsModifiedChangeNewArray() (v []VersionDiffResponseDifferencesAllowedEventsModifiedChangeNewArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesAllowedEventsModifiedChangeNewUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesAllowedEventsModifiedChangeNewUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesAllowedEventsModifiedChangeNewUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesAllowedEventsModifiedChangeNewArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesAllowedEventsModifiedChangeNewArrayItemMapItem]
type VersionDiffResponseDifferencesAllowedEventsModifiedChangeNewArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesAllowedEventsModifiedChangeNewArrayItemMapItem any `json:",inline"`
	JSON                                                                           struct {
		OfString                                                                       respjson.Field
		OfFloat                                                                        respjson.Field
		OfBool                                                                         respjson.Field
		OfAnyArray                                                                     respjson.Field
		OfVersionDiffResponseDifferencesAllowedEventsModifiedChangeNewArrayItemMapItem respjson.Field
		raw                                                                            string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesAllowedEventsModifiedChangeNewArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesAllowedEventsModifiedChangeNewArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesAllowedEventsModifiedChangeNewArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesAllowedEventsModifiedChangeNewArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesAllowedEventsModifiedChangeNewArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesAllowedEventsModifiedChangeNewArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesAllowedEventsModifiedChangeNewArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesAllowedEventsModifiedChangeOldUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesAllowedEventsModifiedChangeOldArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesAllowedEventsModifiedChangeOldArray
// OfVersionDiffResponseDifferencesAllowedEventsModifiedChangeOldMapItem]
type VersionDiffResponseDifferencesAllowedEventsModifiedChangeOldUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesAllowedEventsModifiedChangeOldArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesAllowedEventsModifiedChangeOldArray []VersionDiffResponseDifferencesAllowedEventsModifiedChangeOldArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesAllowedEventsModifiedChangeOldMapItem any `json:",inline"`
	JSON                                                                  struct {
		OfString                                                              respjson.Field
		OfFloat                                                               respjson.Field
		OfBool                                                                respjson.Field
		OfVersionDiffResponseDifferencesAllowedEventsModifiedChangeOldArray   respjson.Field
		OfVersionDiffResponseDifferencesAllowedEventsModifiedChangeOldMapItem respjson.Field
		raw                                                                   string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesAllowedEventsModifiedChangeOldUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesAllowedEventsModifiedChangeOldUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesAllowedEventsModifiedChangeOldUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesAllowedEventsModifiedChangeOldUnion) AsVersionDiffResponseDifferencesAllowedEventsModifiedChangeOldArray() (v []VersionDiffResponseDifferencesAllowedEventsModifiedChangeOldArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesAllowedEventsModifiedChangeOldUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesAllowedEventsModifiedChangeOldUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesAllowedEventsModifiedChangeOldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesAllowedEventsModifiedChangeOldArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesAllowedEventsModifiedChangeOldArrayItemMapItem]
type VersionDiffResponseDifferencesAllowedEventsModifiedChangeOldArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesAllowedEventsModifiedChangeOldArrayItemMapItem any `json:",inline"`
	JSON                                                                           struct {
		OfString                                                                       respjson.Field
		OfFloat                                                                        respjson.Field
		OfBool                                                                         respjson.Field
		OfAnyArray                                                                     respjson.Field
		OfVersionDiffResponseDifferencesAllowedEventsModifiedChangeOldArrayItemMapItem respjson.Field
		raw                                                                            string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesAllowedEventsModifiedChangeOldArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesAllowedEventsModifiedChangeOldArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesAllowedEventsModifiedChangeOldArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesAllowedEventsModifiedChangeOldArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesAllowedEventsModifiedChangeOldArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesAllowedEventsModifiedChangeOldArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesAllowedEventsModifiedChangeOldArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesAllowedEventsModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// Changed fields with credentials redacted. Nested objects use dot paths; arrays
	// are whole values. Missing values are null.
	Changes []VersionDiffResponseDifferencesConsentSettingsModifiedChange `json:"changes" api:"required"`
	New     VersionDiffResponseDifferencesConsentSettingsModifiedNew      `json:"new" api:"required"`
	Old     VersionDiffResponseDifferencesConsentSettingsModifiedOld      `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Changes     respjson.Field
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

type VersionDiffResponseDifferencesConsentSettingsModifiedChange struct {
	New  VersionDiffResponseDifferencesConsentSettingsModifiedChangeNewUnion `json:"new" api:"required"`
	Old  VersionDiffResponseDifferencesConsentSettingsModifiedChangeOldUnion `json:"old" api:"required"`
	Path string                                                              `json:"path" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesConsentSettingsModifiedChange) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesConsentSettingsModifiedChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesConsentSettingsModifiedChangeNewUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesConsentSettingsModifiedChangeNewArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesConsentSettingsModifiedChangeNewArray
// OfVersionDiffResponseDifferencesConsentSettingsModifiedChangeNewMapItem]
type VersionDiffResponseDifferencesConsentSettingsModifiedChangeNewUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesConsentSettingsModifiedChangeNewArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesConsentSettingsModifiedChangeNewArray []VersionDiffResponseDifferencesConsentSettingsModifiedChangeNewArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesConsentSettingsModifiedChangeNewMapItem any `json:",inline"`
	JSON                                                                    struct {
		OfString                                                                respjson.Field
		OfFloat                                                                 respjson.Field
		OfBool                                                                  respjson.Field
		OfVersionDiffResponseDifferencesConsentSettingsModifiedChangeNewArray   respjson.Field
		OfVersionDiffResponseDifferencesConsentSettingsModifiedChangeNewMapItem respjson.Field
		raw                                                                     string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesConsentSettingsModifiedChangeNewUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesConsentSettingsModifiedChangeNewUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesConsentSettingsModifiedChangeNewUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesConsentSettingsModifiedChangeNewUnion) AsVersionDiffResponseDifferencesConsentSettingsModifiedChangeNewArray() (v []VersionDiffResponseDifferencesConsentSettingsModifiedChangeNewArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesConsentSettingsModifiedChangeNewUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesConsentSettingsModifiedChangeNewUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesConsentSettingsModifiedChangeNewUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesConsentSettingsModifiedChangeNewArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesConsentSettingsModifiedChangeNewArrayItemMapItem]
type VersionDiffResponseDifferencesConsentSettingsModifiedChangeNewArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesConsentSettingsModifiedChangeNewArrayItemMapItem any `json:",inline"`
	JSON                                                                             struct {
		OfString                                                                         respjson.Field
		OfFloat                                                                          respjson.Field
		OfBool                                                                           respjson.Field
		OfAnyArray                                                                       respjson.Field
		OfVersionDiffResponseDifferencesConsentSettingsModifiedChangeNewArrayItemMapItem respjson.Field
		raw                                                                              string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesConsentSettingsModifiedChangeNewArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesConsentSettingsModifiedChangeNewArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesConsentSettingsModifiedChangeNewArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesConsentSettingsModifiedChangeNewArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesConsentSettingsModifiedChangeNewArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesConsentSettingsModifiedChangeNewArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesConsentSettingsModifiedChangeNewArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesConsentSettingsModifiedChangeOldUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesConsentSettingsModifiedChangeOldArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesConsentSettingsModifiedChangeOldArray
// OfVersionDiffResponseDifferencesConsentSettingsModifiedChangeOldMapItem]
type VersionDiffResponseDifferencesConsentSettingsModifiedChangeOldUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesConsentSettingsModifiedChangeOldArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesConsentSettingsModifiedChangeOldArray []VersionDiffResponseDifferencesConsentSettingsModifiedChangeOldArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesConsentSettingsModifiedChangeOldMapItem any `json:",inline"`
	JSON                                                                    struct {
		OfString                                                                respjson.Field
		OfFloat                                                                 respjson.Field
		OfBool                                                                  respjson.Field
		OfVersionDiffResponseDifferencesConsentSettingsModifiedChangeOldArray   respjson.Field
		OfVersionDiffResponseDifferencesConsentSettingsModifiedChangeOldMapItem respjson.Field
		raw                                                                     string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesConsentSettingsModifiedChangeOldUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesConsentSettingsModifiedChangeOldUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesConsentSettingsModifiedChangeOldUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesConsentSettingsModifiedChangeOldUnion) AsVersionDiffResponseDifferencesConsentSettingsModifiedChangeOldArray() (v []VersionDiffResponseDifferencesConsentSettingsModifiedChangeOldArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesConsentSettingsModifiedChangeOldUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesConsentSettingsModifiedChangeOldUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesConsentSettingsModifiedChangeOldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesConsentSettingsModifiedChangeOldArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesConsentSettingsModifiedChangeOldArrayItemMapItem]
type VersionDiffResponseDifferencesConsentSettingsModifiedChangeOldArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesConsentSettingsModifiedChangeOldArrayItemMapItem any `json:",inline"`
	JSON                                                                             struct {
		OfString                                                                         respjson.Field
		OfFloat                                                                          respjson.Field
		OfBool                                                                           respjson.Field
		OfAnyArray                                                                       respjson.Field
		OfVersionDiffResponseDifferencesConsentSettingsModifiedChangeOldArrayItemMapItem respjson.Field
		raw                                                                              string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesConsentSettingsModifiedChangeOldArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesConsentSettingsModifiedChangeOldArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesConsentSettingsModifiedChangeOldArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesConsentSettingsModifiedChangeOldArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesConsentSettingsModifiedChangeOldArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesConsentSettingsModifiedChangeOldArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesConsentSettingsModifiedChangeOldArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesConsentSettingsModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// Changed fields with credentials redacted. Nested objects use dot paths; arrays
	// are whole values. Missing values are null.
	Changes []VersionDiffResponseDifferencesDataGovernanceEventsModifiedChange `json:"changes" api:"required"`
	New     VersionDiffResponseDifferencesDataGovernanceEventsModifiedNew      `json:"new" api:"required"`
	Old     VersionDiffResponseDifferencesDataGovernanceEventsModifiedOld      `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Changes     respjson.Field
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

type VersionDiffResponseDifferencesDataGovernanceEventsModifiedChange struct {
	New  VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewUnion `json:"new" api:"required"`
	Old  VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldUnion `json:"old" api:"required"`
	Path string                                                                   `json:"path" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesDataGovernanceEventsModifiedChange) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesDataGovernanceEventsModifiedChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewArray
// OfVersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewMapItem]
type VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewArray []VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewMapItem any `json:",inline"`
	JSON                                                                         struct {
		OfString                                                                     respjson.Field
		OfFloat                                                                      respjson.Field
		OfBool                                                                       respjson.Field
		OfVersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewArray   respjson.Field
		OfVersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewMapItem respjson.Field
		raw                                                                          string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewUnion) AsVersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewArray() (v []VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewArrayItemMapItem]
type VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewArrayItemMapItem any `json:",inline"`
	JSON                                                                                  struct {
		OfString                                                                              respjson.Field
		OfFloat                                                                               respjson.Field
		OfBool                                                                                respjson.Field
		OfAnyArray                                                                            respjson.Field
		OfVersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewArrayItemMapItem respjson.Field
		raw                                                                                   string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeNewArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldArray
// OfVersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldMapItem]
type VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldArray []VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldMapItem any `json:",inline"`
	JSON                                                                         struct {
		OfString                                                                     respjson.Field
		OfFloat                                                                      respjson.Field
		OfBool                                                                       respjson.Field
		OfVersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldArray   respjson.Field
		OfVersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldMapItem respjson.Field
		raw                                                                          string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldUnion) AsVersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldArray() (v []VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldArrayItemMapItem]
type VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldArrayItemMapItem any `json:",inline"`
	JSON                                                                                  struct {
		OfString                                                                              respjson.Field
		OfFloat                                                                               respjson.Field
		OfBool                                                                                respjson.Field
		OfAnyArray                                                                            respjson.Field
		OfVersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldArrayItemMapItem respjson.Field
		raw                                                                                   string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesDataGovernanceEventsModifiedChangeOldArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesDataGovernanceEventsModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// Changed fields with credentials redacted. Nested objects use dot paths; arrays
	// are whole values. Missing values are null.
	Changes []VersionDiffResponseDifferencesDataGovernanceRulesModifiedChange `json:"changes" api:"required"`
	New     VersionDiffResponseDifferencesDataGovernanceRulesModifiedNew      `json:"new" api:"required"`
	Old     VersionDiffResponseDifferencesDataGovernanceRulesModifiedOld      `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Changes     respjson.Field
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

type VersionDiffResponseDifferencesDataGovernanceRulesModifiedChange struct {
	New  VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewUnion `json:"new" api:"required"`
	Old  VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldUnion `json:"old" api:"required"`
	Path string                                                                  `json:"path" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesDataGovernanceRulesModifiedChange) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesDataGovernanceRulesModifiedChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewUnion contains
// all possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewArray
// OfVersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewMapItem]
type VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewArray []VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewMapItem any `json:",inline"`
	JSON                                                                        struct {
		OfString                                                                    respjson.Field
		OfFloat                                                                     respjson.Field
		OfBool                                                                      respjson.Field
		OfVersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewArray   respjson.Field
		OfVersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewMapItem respjson.Field
		raw                                                                         string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewUnion) AsVersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewArray() (v []VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewArrayItemMapItem]
type VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewArrayItemMapItem any `json:",inline"`
	JSON                                                                                 struct {
		OfString                                                                             respjson.Field
		OfFloat                                                                              respjson.Field
		OfBool                                                                               respjson.Field
		OfAnyArray                                                                           respjson.Field
		OfVersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewArrayItemMapItem respjson.Field
		raw                                                                                  string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeNewArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldUnion contains
// all possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldArray
// OfVersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldMapItem]
type VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldArray []VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldMapItem any `json:",inline"`
	JSON                                                                        struct {
		OfString                                                                    respjson.Field
		OfFloat                                                                     respjson.Field
		OfBool                                                                      respjson.Field
		OfVersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldArray   respjson.Field
		OfVersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldMapItem respjson.Field
		raw                                                                         string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldUnion) AsVersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldArray() (v []VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldArrayItemMapItem]
type VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldArrayItemMapItem any `json:",inline"`
	JSON                                                                                 struct {
		OfString                                                                             respjson.Field
		OfFloat                                                                              respjson.Field
		OfBool                                                                               respjson.Field
		OfAnyArray                                                                           respjson.Field
		OfVersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldArrayItemMapItem respjson.Field
		raw                                                                                  string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesDataGovernanceRulesModifiedChangeOldArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesDataGovernanceRulesModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// Changed fields with credentials redacted. Nested objects use dot paths; arrays
	// are whole values. Missing values are null.
	Changes []VersionDiffResponseDifferencesDestinationsModifiedChange `json:"changes" api:"required"`
	New     VersionDiffResponseDifferencesDestinationsModifiedNew      `json:"new" api:"required"`
	Old     VersionDiffResponseDifferencesDestinationsModifiedOld      `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Changes     respjson.Field
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

type VersionDiffResponseDifferencesDestinationsModifiedChange struct {
	New  VersionDiffResponseDifferencesDestinationsModifiedChangeNewUnion `json:"new" api:"required"`
	Old  VersionDiffResponseDifferencesDestinationsModifiedChangeOldUnion `json:"old" api:"required"`
	Path string                                                           `json:"path" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesDestinationsModifiedChange) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesDestinationsModifiedChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesDestinationsModifiedChangeNewUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesDestinationsModifiedChangeNewArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesDestinationsModifiedChangeNewArray
// OfVersionDiffResponseDifferencesDestinationsModifiedChangeNewMapItem]
type VersionDiffResponseDifferencesDestinationsModifiedChangeNewUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesDestinationsModifiedChangeNewArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesDestinationsModifiedChangeNewArray []VersionDiffResponseDifferencesDestinationsModifiedChangeNewArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesDestinationsModifiedChangeNewMapItem any `json:",inline"`
	JSON                                                                 struct {
		OfString                                                             respjson.Field
		OfFloat                                                              respjson.Field
		OfBool                                                               respjson.Field
		OfVersionDiffResponseDifferencesDestinationsModifiedChangeNewArray   respjson.Field
		OfVersionDiffResponseDifferencesDestinationsModifiedChangeNewMapItem respjson.Field
		raw                                                                  string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesDestinationsModifiedChangeNewUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDestinationsModifiedChangeNewUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDestinationsModifiedChangeNewUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDestinationsModifiedChangeNewUnion) AsVersionDiffResponseDifferencesDestinationsModifiedChangeNewArray() (v []VersionDiffResponseDifferencesDestinationsModifiedChangeNewArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDestinationsModifiedChangeNewUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesDestinationsModifiedChangeNewUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesDestinationsModifiedChangeNewUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesDestinationsModifiedChangeNewArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesDestinationsModifiedChangeNewArrayItemMapItem]
type VersionDiffResponseDifferencesDestinationsModifiedChangeNewArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesDestinationsModifiedChangeNewArrayItemMapItem any `json:",inline"`
	JSON                                                                          struct {
		OfString                                                                      respjson.Field
		OfFloat                                                                       respjson.Field
		OfBool                                                                        respjson.Field
		OfAnyArray                                                                    respjson.Field
		OfVersionDiffResponseDifferencesDestinationsModifiedChangeNewArrayItemMapItem respjson.Field
		raw                                                                           string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesDestinationsModifiedChangeNewArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDestinationsModifiedChangeNewArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDestinationsModifiedChangeNewArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDestinationsModifiedChangeNewArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDestinationsModifiedChangeNewArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesDestinationsModifiedChangeNewArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesDestinationsModifiedChangeNewArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesDestinationsModifiedChangeOldUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesDestinationsModifiedChangeOldArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesDestinationsModifiedChangeOldArray
// OfVersionDiffResponseDifferencesDestinationsModifiedChangeOldMapItem]
type VersionDiffResponseDifferencesDestinationsModifiedChangeOldUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesDestinationsModifiedChangeOldArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesDestinationsModifiedChangeOldArray []VersionDiffResponseDifferencesDestinationsModifiedChangeOldArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesDestinationsModifiedChangeOldMapItem any `json:",inline"`
	JSON                                                                 struct {
		OfString                                                             respjson.Field
		OfFloat                                                              respjson.Field
		OfBool                                                               respjson.Field
		OfVersionDiffResponseDifferencesDestinationsModifiedChangeOldArray   respjson.Field
		OfVersionDiffResponseDifferencesDestinationsModifiedChangeOldMapItem respjson.Field
		raw                                                                  string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesDestinationsModifiedChangeOldUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDestinationsModifiedChangeOldUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDestinationsModifiedChangeOldUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDestinationsModifiedChangeOldUnion) AsVersionDiffResponseDifferencesDestinationsModifiedChangeOldArray() (v []VersionDiffResponseDifferencesDestinationsModifiedChangeOldArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDestinationsModifiedChangeOldUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesDestinationsModifiedChangeOldUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesDestinationsModifiedChangeOldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesDestinationsModifiedChangeOldArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesDestinationsModifiedChangeOldArrayItemMapItem]
type VersionDiffResponseDifferencesDestinationsModifiedChangeOldArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesDestinationsModifiedChangeOldArrayItemMapItem any `json:",inline"`
	JSON                                                                          struct {
		OfString                                                                      respjson.Field
		OfFloat                                                                       respjson.Field
		OfBool                                                                        respjson.Field
		OfAnyArray                                                                    respjson.Field
		OfVersionDiffResponseDifferencesDestinationsModifiedChangeOldArrayItemMapItem respjson.Field
		raw                                                                           string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesDestinationsModifiedChangeOldArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDestinationsModifiedChangeOldArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDestinationsModifiedChangeOldArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDestinationsModifiedChangeOldArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesDestinationsModifiedChangeOldArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesDestinationsModifiedChangeOldArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesDestinationsModifiedChangeOldArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesDestinationsModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// Changed fields with credentials redacted. Nested objects use dot paths; arrays
	// are whole values. Missing values are null.
	Changes []VersionDiffResponseDifferencesExperimentsModifiedChange `json:"changes" api:"required"`
	New     VersionDiffResponseDifferencesExperimentsModifiedNew      `json:"new" api:"required"`
	Old     VersionDiffResponseDifferencesExperimentsModifiedOld      `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Changes     respjson.Field
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

type VersionDiffResponseDifferencesExperimentsModifiedChange struct {
	New  VersionDiffResponseDifferencesExperimentsModifiedChangeNewUnion `json:"new" api:"required"`
	Old  VersionDiffResponseDifferencesExperimentsModifiedChangeOldUnion `json:"old" api:"required"`
	Path string                                                          `json:"path" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesExperimentsModifiedChange) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesExperimentsModifiedChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesExperimentsModifiedChangeNewUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesExperimentsModifiedChangeNewArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesExperimentsModifiedChangeNewArray
// OfVersionDiffResponseDifferencesExperimentsModifiedChangeNewMapItem]
type VersionDiffResponseDifferencesExperimentsModifiedChangeNewUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesExperimentsModifiedChangeNewArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesExperimentsModifiedChangeNewArray []VersionDiffResponseDifferencesExperimentsModifiedChangeNewArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesExperimentsModifiedChangeNewMapItem any `json:",inline"`
	JSON                                                                struct {
		OfString                                                            respjson.Field
		OfFloat                                                             respjson.Field
		OfBool                                                              respjson.Field
		OfVersionDiffResponseDifferencesExperimentsModifiedChangeNewArray   respjson.Field
		OfVersionDiffResponseDifferencesExperimentsModifiedChangeNewMapItem respjson.Field
		raw                                                                 string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesExperimentsModifiedChangeNewUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentsModifiedChangeNewUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentsModifiedChangeNewUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentsModifiedChangeNewUnion) AsVersionDiffResponseDifferencesExperimentsModifiedChangeNewArray() (v []VersionDiffResponseDifferencesExperimentsModifiedChangeNewArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentsModifiedChangeNewUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesExperimentsModifiedChangeNewUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesExperimentsModifiedChangeNewUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesExperimentsModifiedChangeNewArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesExperimentsModifiedChangeNewArrayItemMapItem]
type VersionDiffResponseDifferencesExperimentsModifiedChangeNewArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesExperimentsModifiedChangeNewArrayItemMapItem any `json:",inline"`
	JSON                                                                         struct {
		OfString                                                                     respjson.Field
		OfFloat                                                                      respjson.Field
		OfBool                                                                       respjson.Field
		OfAnyArray                                                                   respjson.Field
		OfVersionDiffResponseDifferencesExperimentsModifiedChangeNewArrayItemMapItem respjson.Field
		raw                                                                          string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesExperimentsModifiedChangeNewArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentsModifiedChangeNewArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentsModifiedChangeNewArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentsModifiedChangeNewArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentsModifiedChangeNewArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesExperimentsModifiedChangeNewArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesExperimentsModifiedChangeNewArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesExperimentsModifiedChangeOldUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesExperimentsModifiedChangeOldArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesExperimentsModifiedChangeOldArray
// OfVersionDiffResponseDifferencesExperimentsModifiedChangeOldMapItem]
type VersionDiffResponseDifferencesExperimentsModifiedChangeOldUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesExperimentsModifiedChangeOldArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesExperimentsModifiedChangeOldArray []VersionDiffResponseDifferencesExperimentsModifiedChangeOldArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesExperimentsModifiedChangeOldMapItem any `json:",inline"`
	JSON                                                                struct {
		OfString                                                            respjson.Field
		OfFloat                                                             respjson.Field
		OfBool                                                              respjson.Field
		OfVersionDiffResponseDifferencesExperimentsModifiedChangeOldArray   respjson.Field
		OfVersionDiffResponseDifferencesExperimentsModifiedChangeOldMapItem respjson.Field
		raw                                                                 string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesExperimentsModifiedChangeOldUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentsModifiedChangeOldUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentsModifiedChangeOldUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentsModifiedChangeOldUnion) AsVersionDiffResponseDifferencesExperimentsModifiedChangeOldArray() (v []VersionDiffResponseDifferencesExperimentsModifiedChangeOldArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentsModifiedChangeOldUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesExperimentsModifiedChangeOldUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesExperimentsModifiedChangeOldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesExperimentsModifiedChangeOldArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesExperimentsModifiedChangeOldArrayItemMapItem]
type VersionDiffResponseDifferencesExperimentsModifiedChangeOldArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesExperimentsModifiedChangeOldArrayItemMapItem any `json:",inline"`
	JSON                                                                         struct {
		OfString                                                                     respjson.Field
		OfFloat                                                                      respjson.Field
		OfBool                                                                       respjson.Field
		OfAnyArray                                                                   respjson.Field
		OfVersionDiffResponseDifferencesExperimentsModifiedChangeOldArrayItemMapItem respjson.Field
		raw                                                                          string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesExperimentsModifiedChangeOldArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentsModifiedChangeOldArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentsModifiedChangeOldArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentsModifiedChangeOldArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentsModifiedChangeOldArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesExperimentsModifiedChangeOldArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesExperimentsModifiedChangeOldArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesExperimentsModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// Changed fields with credentials redacted. Nested objects use dot paths; arrays
	// are whole values. Missing values are null.
	Changes []VersionDiffResponseDifferencesExperimentSettingsModifiedChange `json:"changes" api:"required"`
	New     VersionDiffResponseDifferencesExperimentSettingsModifiedNew      `json:"new" api:"required"`
	Old     VersionDiffResponseDifferencesExperimentSettingsModifiedOld      `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Changes     respjson.Field
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

type VersionDiffResponseDifferencesExperimentSettingsModifiedChange struct {
	New  VersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewUnion `json:"new" api:"required"`
	Old  VersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldUnion `json:"old" api:"required"`
	Path string                                                                 `json:"path" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesExperimentSettingsModifiedChange) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesExperimentSettingsModifiedChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewUnion contains
// all possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewArray
// OfVersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewMapItem]
type VersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewArray []VersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewMapItem any `json:",inline"`
	JSON                                                                       struct {
		OfString                                                                   respjson.Field
		OfFloat                                                                    respjson.Field
		OfBool                                                                     respjson.Field
		OfVersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewArray   respjson.Field
		OfVersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewMapItem respjson.Field
		raw                                                                        string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewUnion) AsVersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewArray() (v []VersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewArrayItemMapItem]
type VersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewArrayItemMapItem any `json:",inline"`
	JSON                                                                                struct {
		OfString                                                                            respjson.Field
		OfFloat                                                                             respjson.Field
		OfBool                                                                              respjson.Field
		OfAnyArray                                                                          respjson.Field
		OfVersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewArrayItemMapItem respjson.Field
		raw                                                                                 string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesExperimentSettingsModifiedChangeNewArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldUnion contains
// all possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldArray
// OfVersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldMapItem]
type VersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldArray []VersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldMapItem any `json:",inline"`
	JSON                                                                       struct {
		OfString                                                                   respjson.Field
		OfFloat                                                                    respjson.Field
		OfBool                                                                     respjson.Field
		OfVersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldArray   respjson.Field
		OfVersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldMapItem respjson.Field
		raw                                                                        string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldUnion) AsVersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldArray() (v []VersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldArrayItemMapItem]
type VersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldArrayItemMapItem any `json:",inline"`
	JSON                                                                                struct {
		OfString                                                                            respjson.Field
		OfFloat                                                                             respjson.Field
		OfBool                                                                              respjson.Field
		OfAnyArray                                                                          respjson.Field
		OfVersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldArrayItemMapItem respjson.Field
		raw                                                                                 string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesExperimentSettingsModifiedChangeOldArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesExperimentSettingsModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// Changed fields with credentials redacted. Nested objects use dot paths; arrays
	// are whole values. Missing values are null.
	Changes []VersionDiffResponseDifferencesExperimentVariantsModifiedChange `json:"changes" api:"required"`
	New     VersionDiffResponseDifferencesExperimentVariantsModifiedNew      `json:"new" api:"required"`
	Old     VersionDiffResponseDifferencesExperimentVariantsModifiedOld      `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Changes     respjson.Field
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

type VersionDiffResponseDifferencesExperimentVariantsModifiedChange struct {
	New  VersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewUnion `json:"new" api:"required"`
	Old  VersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldUnion `json:"old" api:"required"`
	Path string                                                                 `json:"path" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesExperimentVariantsModifiedChange) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesExperimentVariantsModifiedChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewUnion contains
// all possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewArray
// OfVersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewMapItem]
type VersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewArray []VersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewMapItem any `json:",inline"`
	JSON                                                                       struct {
		OfString                                                                   respjson.Field
		OfFloat                                                                    respjson.Field
		OfBool                                                                     respjson.Field
		OfVersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewArray   respjson.Field
		OfVersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewMapItem respjson.Field
		raw                                                                        string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewUnion) AsVersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewArray() (v []VersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewArrayItemMapItem]
type VersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewArrayItemMapItem any `json:",inline"`
	JSON                                                                                struct {
		OfString                                                                            respjson.Field
		OfFloat                                                                             respjson.Field
		OfBool                                                                              respjson.Field
		OfAnyArray                                                                          respjson.Field
		OfVersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewArrayItemMapItem respjson.Field
		raw                                                                                 string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesExperimentVariantsModifiedChangeNewArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldUnion contains
// all possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldArray
// OfVersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldMapItem]
type VersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldArray []VersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldMapItem any `json:",inline"`
	JSON                                                                       struct {
		OfString                                                                   respjson.Field
		OfFloat                                                                    respjson.Field
		OfBool                                                                     respjson.Field
		OfVersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldArray   respjson.Field
		OfVersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldMapItem respjson.Field
		raw                                                                        string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldUnion) AsVersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldArray() (v []VersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldArrayItemMapItem]
type VersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldArrayItemMapItem any `json:",inline"`
	JSON                                                                                struct {
		OfString                                                                            respjson.Field
		OfFloat                                                                             respjson.Field
		OfBool                                                                              respjson.Field
		OfAnyArray                                                                          respjson.Field
		OfVersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldArrayItemMapItem respjson.Field
		raw                                                                                 string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesExperimentVariantsModifiedChangeOldArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesExperimentVariantsModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// Changed fields with credentials redacted. Nested objects use dot paths; arrays
	// are whole values. Missing values are null.
	Changes []VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChange `json:"changes" api:"required"`
	New     VersionDiffResponseDifferencesExternalAllowedEventDataModifiedNew      `json:"new" api:"required"`
	Old     VersionDiffResponseDifferencesExternalAllowedEventDataModifiedOld      `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Changes     respjson.Field
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

type VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChange struct {
	New  VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewUnion `json:"new" api:"required"`
	Old  VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldUnion `json:"old" api:"required"`
	Path string                                                                       `json:"path" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChange) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewArray
// OfVersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewMapItem]
type VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewArray []VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewMapItem any `json:",inline"`
	JSON                                                                             struct {
		OfString                                                                         respjson.Field
		OfFloat                                                                          respjson.Field
		OfBool                                                                           respjson.Field
		OfVersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewArray   respjson.Field
		OfVersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewMapItem respjson.Field
		raw                                                                              string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewUnion) AsVersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewArray() (v []VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewArrayItemMapItem]
type VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewArrayItemMapItem any `json:",inline"`
	JSON                                                                                      struct {
		OfString                                                                                  respjson.Field
		OfFloat                                                                                   respjson.Field
		OfBool                                                                                    respjson.Field
		OfAnyArray                                                                                respjson.Field
		OfVersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewArrayItemMapItem respjson.Field
		raw                                                                                       string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeNewArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldArray
// OfVersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldMapItem]
type VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldArray []VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldMapItem any `json:",inline"`
	JSON                                                                             struct {
		OfString                                                                         respjson.Field
		OfFloat                                                                          respjson.Field
		OfBool                                                                           respjson.Field
		OfVersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldArray   respjson.Field
		OfVersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldMapItem respjson.Field
		raw                                                                              string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldUnion) AsVersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldArray() (v []VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldArrayItemMapItem]
type VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldArrayItemMapItem any `json:",inline"`
	JSON                                                                                      struct {
		OfString                                                                                  respjson.Field
		OfFloat                                                                                   respjson.Field
		OfBool                                                                                    respjson.Field
		OfAnyArray                                                                                respjson.Field
		OfVersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldArrayItemMapItem respjson.Field
		raw                                                                                       string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesExternalAllowedEventDataModifiedChangeOldArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesExternalAllowedEventDataModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// Changed fields with credentials redacted. Nested objects use dot paths; arrays
	// are whole values. Missing values are null.
	Changes []VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChange `json:"changes" api:"required"`
	New     VersionDiffResponseDifferencesGlobalDispatchCentersModifiedNew      `json:"new" api:"required"`
	Old     VersionDiffResponseDifferencesGlobalDispatchCentersModifiedOld      `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Changes     respjson.Field
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

type VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChange struct {
	New  VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewUnion `json:"new" api:"required"`
	Old  VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldUnion `json:"old" api:"required"`
	Path string                                                                    `json:"path" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChange) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewArray
// OfVersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewMapItem]
type VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewArray []VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewMapItem any `json:",inline"`
	JSON                                                                          struct {
		OfString                                                                      respjson.Field
		OfFloat                                                                       respjson.Field
		OfBool                                                                        respjson.Field
		OfVersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewArray   respjson.Field
		OfVersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewMapItem respjson.Field
		raw                                                                           string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewUnion) AsVersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewArray() (v []VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewArrayItemMapItem]
type VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewArrayItemMapItem any `json:",inline"`
	JSON                                                                                   struct {
		OfString                                                                               respjson.Field
		OfFloat                                                                                respjson.Field
		OfBool                                                                                 respjson.Field
		OfAnyArray                                                                             respjson.Field
		OfVersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewArrayItemMapItem respjson.Field
		raw                                                                                    string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeNewArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldArray
// OfVersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldMapItem]
type VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldArray []VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldMapItem any `json:",inline"`
	JSON                                                                          struct {
		OfString                                                                      respjson.Field
		OfFloat                                                                       respjson.Field
		OfBool                                                                        respjson.Field
		OfVersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldArray   respjson.Field
		OfVersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldMapItem respjson.Field
		raw                                                                           string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldUnion) AsVersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldArray() (v []VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldArrayItemMapItem]
type VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldArrayItemMapItem any `json:",inline"`
	JSON                                                                                   struct {
		OfString                                                                               respjson.Field
		OfFloat                                                                                respjson.Field
		OfBool                                                                                 respjson.Field
		OfAnyArray                                                                             respjson.Field
		OfVersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldArrayItemMapItem respjson.Field
		raw                                                                                    string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesGlobalDispatchCentersModifiedChangeOldArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesGlobalDispatchCentersModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// Changed fields with credentials redacted. Nested objects use dot paths; arrays
	// are whole values. Missing values are null.
	Changes []VersionDiffResponseDifferencesMappingsModifiedChange `json:"changes" api:"required"`
	New     VersionDiffResponseDifferencesMappingsModifiedNew      `json:"new" api:"required"`
	Old     VersionDiffResponseDifferencesMappingsModifiedOld      `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Changes     respjson.Field
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

type VersionDiffResponseDifferencesMappingsModifiedChange struct {
	New  VersionDiffResponseDifferencesMappingsModifiedChangeNewUnion `json:"new" api:"required"`
	Old  VersionDiffResponseDifferencesMappingsModifiedChangeOldUnion `json:"old" api:"required"`
	Path string                                                       `json:"path" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesMappingsModifiedChange) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesMappingsModifiedChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesMappingsModifiedChangeNewUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesMappingsModifiedChangeNewArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesMappingsModifiedChangeNewArray
// OfVersionDiffResponseDifferencesMappingsModifiedChangeNewMapItem]
type VersionDiffResponseDifferencesMappingsModifiedChangeNewUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesMappingsModifiedChangeNewArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesMappingsModifiedChangeNewArray []VersionDiffResponseDifferencesMappingsModifiedChangeNewArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesMappingsModifiedChangeNewMapItem any `json:",inline"`
	JSON                                                             struct {
		OfString                                                         respjson.Field
		OfFloat                                                          respjson.Field
		OfBool                                                           respjson.Field
		OfVersionDiffResponseDifferencesMappingsModifiedChangeNewArray   respjson.Field
		OfVersionDiffResponseDifferencesMappingsModifiedChangeNewMapItem respjson.Field
		raw                                                              string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesMappingsModifiedChangeNewUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesMappingsModifiedChangeNewUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesMappingsModifiedChangeNewUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesMappingsModifiedChangeNewUnion) AsVersionDiffResponseDifferencesMappingsModifiedChangeNewArray() (v []VersionDiffResponseDifferencesMappingsModifiedChangeNewArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesMappingsModifiedChangeNewUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesMappingsModifiedChangeNewUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesMappingsModifiedChangeNewUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesMappingsModifiedChangeNewArrayItemUnion contains
// all possible properties and values from [string], [float64], [bool], [[]any],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesMappingsModifiedChangeNewArrayItemMapItem]
type VersionDiffResponseDifferencesMappingsModifiedChangeNewArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesMappingsModifiedChangeNewArrayItemMapItem any `json:",inline"`
	JSON                                                                      struct {
		OfString                                                                  respjson.Field
		OfFloat                                                                   respjson.Field
		OfBool                                                                    respjson.Field
		OfAnyArray                                                                respjson.Field
		OfVersionDiffResponseDifferencesMappingsModifiedChangeNewArrayItemMapItem respjson.Field
		raw                                                                       string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesMappingsModifiedChangeNewArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesMappingsModifiedChangeNewArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesMappingsModifiedChangeNewArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesMappingsModifiedChangeNewArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesMappingsModifiedChangeNewArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesMappingsModifiedChangeNewArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesMappingsModifiedChangeNewArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesMappingsModifiedChangeOldUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesMappingsModifiedChangeOldArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesMappingsModifiedChangeOldArray
// OfVersionDiffResponseDifferencesMappingsModifiedChangeOldMapItem]
type VersionDiffResponseDifferencesMappingsModifiedChangeOldUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesMappingsModifiedChangeOldArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesMappingsModifiedChangeOldArray []VersionDiffResponseDifferencesMappingsModifiedChangeOldArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesMappingsModifiedChangeOldMapItem any `json:",inline"`
	JSON                                                             struct {
		OfString                                                         respjson.Field
		OfFloat                                                          respjson.Field
		OfBool                                                           respjson.Field
		OfVersionDiffResponseDifferencesMappingsModifiedChangeOldArray   respjson.Field
		OfVersionDiffResponseDifferencesMappingsModifiedChangeOldMapItem respjson.Field
		raw                                                              string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesMappingsModifiedChangeOldUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesMappingsModifiedChangeOldUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesMappingsModifiedChangeOldUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesMappingsModifiedChangeOldUnion) AsVersionDiffResponseDifferencesMappingsModifiedChangeOldArray() (v []VersionDiffResponseDifferencesMappingsModifiedChangeOldArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesMappingsModifiedChangeOldUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesMappingsModifiedChangeOldUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesMappingsModifiedChangeOldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesMappingsModifiedChangeOldArrayItemUnion contains
// all possible properties and values from [string], [float64], [bool], [[]any],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesMappingsModifiedChangeOldArrayItemMapItem]
type VersionDiffResponseDifferencesMappingsModifiedChangeOldArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesMappingsModifiedChangeOldArrayItemMapItem any `json:",inline"`
	JSON                                                                      struct {
		OfString                                                                  respjson.Field
		OfFloat                                                                   respjson.Field
		OfBool                                                                    respjson.Field
		OfAnyArray                                                                respjson.Field
		OfVersionDiffResponseDifferencesMappingsModifiedChangeOldArrayItemMapItem respjson.Field
		raw                                                                       string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesMappingsModifiedChangeOldArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesMappingsModifiedChangeOldArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesMappingsModifiedChangeOldArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesMappingsModifiedChangeOldArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesMappingsModifiedChangeOldArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesMappingsModifiedChangeOldArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesMappingsModifiedChangeOldArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesMappingsModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// Changed fields with credentials redacted. Nested objects use dot paths; arrays
	// are whole values. Missing values are null.
	Changes []VersionDiffResponseDifferencesReplaySettingsModifiedChange `json:"changes" api:"required"`
	New     VersionDiffResponseDifferencesReplaySettingsModifiedNew      `json:"new" api:"required"`
	Old     VersionDiffResponseDifferencesReplaySettingsModifiedOld      `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Changes     respjson.Field
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

type VersionDiffResponseDifferencesReplaySettingsModifiedChange struct {
	New  VersionDiffResponseDifferencesReplaySettingsModifiedChangeNewUnion `json:"new" api:"required"`
	Old  VersionDiffResponseDifferencesReplaySettingsModifiedChangeOldUnion `json:"old" api:"required"`
	Path string                                                             `json:"path" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesReplaySettingsModifiedChange) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesReplaySettingsModifiedChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesReplaySettingsModifiedChangeNewUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesReplaySettingsModifiedChangeNewArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesReplaySettingsModifiedChangeNewArray
// OfVersionDiffResponseDifferencesReplaySettingsModifiedChangeNewMapItem]
type VersionDiffResponseDifferencesReplaySettingsModifiedChangeNewUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesReplaySettingsModifiedChangeNewArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesReplaySettingsModifiedChangeNewArray []VersionDiffResponseDifferencesReplaySettingsModifiedChangeNewArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesReplaySettingsModifiedChangeNewMapItem any `json:",inline"`
	JSON                                                                   struct {
		OfString                                                               respjson.Field
		OfFloat                                                                respjson.Field
		OfBool                                                                 respjson.Field
		OfVersionDiffResponseDifferencesReplaySettingsModifiedChangeNewArray   respjson.Field
		OfVersionDiffResponseDifferencesReplaySettingsModifiedChangeNewMapItem respjson.Field
		raw                                                                    string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesReplaySettingsModifiedChangeNewUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesReplaySettingsModifiedChangeNewUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesReplaySettingsModifiedChangeNewUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesReplaySettingsModifiedChangeNewUnion) AsVersionDiffResponseDifferencesReplaySettingsModifiedChangeNewArray() (v []VersionDiffResponseDifferencesReplaySettingsModifiedChangeNewArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesReplaySettingsModifiedChangeNewUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesReplaySettingsModifiedChangeNewUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesReplaySettingsModifiedChangeNewUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesReplaySettingsModifiedChangeNewArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesReplaySettingsModifiedChangeNewArrayItemMapItem]
type VersionDiffResponseDifferencesReplaySettingsModifiedChangeNewArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesReplaySettingsModifiedChangeNewArrayItemMapItem any `json:",inline"`
	JSON                                                                            struct {
		OfString                                                                        respjson.Field
		OfFloat                                                                         respjson.Field
		OfBool                                                                          respjson.Field
		OfAnyArray                                                                      respjson.Field
		OfVersionDiffResponseDifferencesReplaySettingsModifiedChangeNewArrayItemMapItem respjson.Field
		raw                                                                             string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesReplaySettingsModifiedChangeNewArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesReplaySettingsModifiedChangeNewArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesReplaySettingsModifiedChangeNewArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesReplaySettingsModifiedChangeNewArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesReplaySettingsModifiedChangeNewArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesReplaySettingsModifiedChangeNewArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesReplaySettingsModifiedChangeNewArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesReplaySettingsModifiedChangeOldUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesReplaySettingsModifiedChangeOldArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesReplaySettingsModifiedChangeOldArray
// OfVersionDiffResponseDifferencesReplaySettingsModifiedChangeOldMapItem]
type VersionDiffResponseDifferencesReplaySettingsModifiedChangeOldUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesReplaySettingsModifiedChangeOldArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesReplaySettingsModifiedChangeOldArray []VersionDiffResponseDifferencesReplaySettingsModifiedChangeOldArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesReplaySettingsModifiedChangeOldMapItem any `json:",inline"`
	JSON                                                                   struct {
		OfString                                                               respjson.Field
		OfFloat                                                                respjson.Field
		OfBool                                                                 respjson.Field
		OfVersionDiffResponseDifferencesReplaySettingsModifiedChangeOldArray   respjson.Field
		OfVersionDiffResponseDifferencesReplaySettingsModifiedChangeOldMapItem respjson.Field
		raw                                                                    string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesReplaySettingsModifiedChangeOldUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesReplaySettingsModifiedChangeOldUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesReplaySettingsModifiedChangeOldUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesReplaySettingsModifiedChangeOldUnion) AsVersionDiffResponseDifferencesReplaySettingsModifiedChangeOldArray() (v []VersionDiffResponseDifferencesReplaySettingsModifiedChangeOldArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesReplaySettingsModifiedChangeOldUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesReplaySettingsModifiedChangeOldUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesReplaySettingsModifiedChangeOldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesReplaySettingsModifiedChangeOldArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesReplaySettingsModifiedChangeOldArrayItemMapItem]
type VersionDiffResponseDifferencesReplaySettingsModifiedChangeOldArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesReplaySettingsModifiedChangeOldArrayItemMapItem any `json:",inline"`
	JSON                                                                            struct {
		OfString                                                                        respjson.Field
		OfFloat                                                                         respjson.Field
		OfBool                                                                          respjson.Field
		OfAnyArray                                                                      respjson.Field
		OfVersionDiffResponseDifferencesReplaySettingsModifiedChangeOldArrayItemMapItem respjson.Field
		raw                                                                             string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesReplaySettingsModifiedChangeOldArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesReplaySettingsModifiedChangeOldArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesReplaySettingsModifiedChangeOldArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesReplaySettingsModifiedChangeOldArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesReplaySettingsModifiedChangeOldArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesReplaySettingsModifiedChangeOldArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesReplaySettingsModifiedChangeOldArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesReplaySettingsModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesReplaySettingsRemoved) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesReplaySettingsRemoved) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesShortLinks struct {
	Added []VersionDiffResponseDifferencesShortLinksAdded `json:"added" api:"required"`
	// Entities present in both snapshots but with at least one field changed. `old` is
	// the snapshot of the entity in the baseline (latest published); `new` is the
	// snapshot in the comparison target (the draft).
	Modified []VersionDiffResponseDifferencesShortLinksModified `json:"modified" api:"required"`
	Removed  []VersionDiffResponseDifferencesShortLinksRemoved  `json:"removed" api:"required"`
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
func (r VersionDiffResponseDifferencesShortLinks) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesShortLinks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesShortLinksAdded struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesShortLinksAdded) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesShortLinksAdded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesShortLinksModified struct {
	// Changed fields with credentials redacted. Nested objects use dot paths; arrays
	// are whole values. Missing values are null.
	Changes []VersionDiffResponseDifferencesShortLinksModifiedChange `json:"changes" api:"required"`
	New     VersionDiffResponseDifferencesShortLinksModifiedNew      `json:"new" api:"required"`
	Old     VersionDiffResponseDifferencesShortLinksModifiedOld      `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Changes     respjson.Field
		New         respjson.Field
		Old         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesShortLinksModified) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesShortLinksModified) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesShortLinksModifiedChange struct {
	New  VersionDiffResponseDifferencesShortLinksModifiedChangeNewUnion `json:"new" api:"required"`
	Old  VersionDiffResponseDifferencesShortLinksModifiedChangeOldUnion `json:"old" api:"required"`
	Path string                                                         `json:"path" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesShortLinksModifiedChange) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesShortLinksModifiedChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesShortLinksModifiedChangeNewUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesShortLinksModifiedChangeNewArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesShortLinksModifiedChangeNewArray
// OfVersionDiffResponseDifferencesShortLinksModifiedChangeNewMapItem]
type VersionDiffResponseDifferencesShortLinksModifiedChangeNewUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesShortLinksModifiedChangeNewArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesShortLinksModifiedChangeNewArray []VersionDiffResponseDifferencesShortLinksModifiedChangeNewArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesShortLinksModifiedChangeNewMapItem any `json:",inline"`
	JSON                                                               struct {
		OfString                                                           respjson.Field
		OfFloat                                                            respjson.Field
		OfBool                                                             respjson.Field
		OfVersionDiffResponseDifferencesShortLinksModifiedChangeNewArray   respjson.Field
		OfVersionDiffResponseDifferencesShortLinksModifiedChangeNewMapItem respjson.Field
		raw                                                                string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesShortLinksModifiedChangeNewUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesShortLinksModifiedChangeNewUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesShortLinksModifiedChangeNewUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesShortLinksModifiedChangeNewUnion) AsVersionDiffResponseDifferencesShortLinksModifiedChangeNewArray() (v []VersionDiffResponseDifferencesShortLinksModifiedChangeNewArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesShortLinksModifiedChangeNewUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesShortLinksModifiedChangeNewUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesShortLinksModifiedChangeNewUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesShortLinksModifiedChangeNewArrayItemUnion contains
// all possible properties and values from [string], [float64], [bool], [[]any],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesShortLinksModifiedChangeNewArrayItemMapItem]
type VersionDiffResponseDifferencesShortLinksModifiedChangeNewArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesShortLinksModifiedChangeNewArrayItemMapItem any `json:",inline"`
	JSON                                                                        struct {
		OfString                                                                    respjson.Field
		OfFloat                                                                     respjson.Field
		OfBool                                                                      respjson.Field
		OfAnyArray                                                                  respjson.Field
		OfVersionDiffResponseDifferencesShortLinksModifiedChangeNewArrayItemMapItem respjson.Field
		raw                                                                         string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesShortLinksModifiedChangeNewArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesShortLinksModifiedChangeNewArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesShortLinksModifiedChangeNewArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesShortLinksModifiedChangeNewArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesShortLinksModifiedChangeNewArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesShortLinksModifiedChangeNewArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesShortLinksModifiedChangeNewArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesShortLinksModifiedChangeOldUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesShortLinksModifiedChangeOldArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesShortLinksModifiedChangeOldArray
// OfVersionDiffResponseDifferencesShortLinksModifiedChangeOldMapItem]
type VersionDiffResponseDifferencesShortLinksModifiedChangeOldUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesShortLinksModifiedChangeOldArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesShortLinksModifiedChangeOldArray []VersionDiffResponseDifferencesShortLinksModifiedChangeOldArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesShortLinksModifiedChangeOldMapItem any `json:",inline"`
	JSON                                                               struct {
		OfString                                                           respjson.Field
		OfFloat                                                            respjson.Field
		OfBool                                                             respjson.Field
		OfVersionDiffResponseDifferencesShortLinksModifiedChangeOldArray   respjson.Field
		OfVersionDiffResponseDifferencesShortLinksModifiedChangeOldMapItem respjson.Field
		raw                                                                string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesShortLinksModifiedChangeOldUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesShortLinksModifiedChangeOldUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesShortLinksModifiedChangeOldUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesShortLinksModifiedChangeOldUnion) AsVersionDiffResponseDifferencesShortLinksModifiedChangeOldArray() (v []VersionDiffResponseDifferencesShortLinksModifiedChangeOldArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesShortLinksModifiedChangeOldUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesShortLinksModifiedChangeOldUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesShortLinksModifiedChangeOldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesShortLinksModifiedChangeOldArrayItemUnion contains
// all possible properties and values from [string], [float64], [bool], [[]any],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesShortLinksModifiedChangeOldArrayItemMapItem]
type VersionDiffResponseDifferencesShortLinksModifiedChangeOldArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesShortLinksModifiedChangeOldArrayItemMapItem any `json:",inline"`
	JSON                                                                        struct {
		OfString                                                                    respjson.Field
		OfFloat                                                                     respjson.Field
		OfBool                                                                      respjson.Field
		OfAnyArray                                                                  respjson.Field
		OfVersionDiffResponseDifferencesShortLinksModifiedChangeOldArrayItemMapItem respjson.Field
		raw                                                                         string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesShortLinksModifiedChangeOldArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesShortLinksModifiedChangeOldArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesShortLinksModifiedChangeOldArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesShortLinksModifiedChangeOldArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesShortLinksModifiedChangeOldArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesShortLinksModifiedChangeOldArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesShortLinksModifiedChangeOldArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesShortLinksModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesShortLinksModifiedNew) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesShortLinksModifiedNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesShortLinksModifiedOld struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesShortLinksModifiedOld) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesShortLinksModifiedOld) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesShortLinksRemoved struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesShortLinksRemoved) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesShortLinksRemoved) UnmarshalJSON(data []byte) error {
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// Changed fields with credentials redacted. Nested objects use dot paths; arrays
	// are whole values. Missing values are null.
	Changes []VersionDiffResponseDifferencesSourcesModifiedChange `json:"changes" api:"required"`
	New     VersionDiffResponseDifferencesSourcesModifiedNew      `json:"new" api:"required"`
	Old     VersionDiffResponseDifferencesSourcesModifiedOld      `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Changes     respjson.Field
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

type VersionDiffResponseDifferencesSourcesModifiedChange struct {
	New  VersionDiffResponseDifferencesSourcesModifiedChangeNewUnion `json:"new" api:"required"`
	Old  VersionDiffResponseDifferencesSourcesModifiedChangeOldUnion `json:"old" api:"required"`
	Path string                                                      `json:"path" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesSourcesModifiedChange) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesSourcesModifiedChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesSourcesModifiedChangeNewUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesSourcesModifiedChangeNewArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesSourcesModifiedChangeNewArray
// OfVersionDiffResponseDifferencesSourcesModifiedChangeNewMapItem]
type VersionDiffResponseDifferencesSourcesModifiedChangeNewUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesSourcesModifiedChangeNewArrayItemUnion] instead
	// of an object.
	OfVersionDiffResponseDifferencesSourcesModifiedChangeNewArray []VersionDiffResponseDifferencesSourcesModifiedChangeNewArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesSourcesModifiedChangeNewMapItem any `json:",inline"`
	JSON                                                            struct {
		OfString                                                        respjson.Field
		OfFloat                                                         respjson.Field
		OfBool                                                          respjson.Field
		OfVersionDiffResponseDifferencesSourcesModifiedChangeNewArray   respjson.Field
		OfVersionDiffResponseDifferencesSourcesModifiedChangeNewMapItem respjson.Field
		raw                                                             string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesSourcesModifiedChangeNewUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesSourcesModifiedChangeNewUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesSourcesModifiedChangeNewUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesSourcesModifiedChangeNewUnion) AsVersionDiffResponseDifferencesSourcesModifiedChangeNewArray() (v []VersionDiffResponseDifferencesSourcesModifiedChangeNewArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesSourcesModifiedChangeNewUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesSourcesModifiedChangeNewUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesSourcesModifiedChangeNewUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesSourcesModifiedChangeNewArrayItemUnion contains
// all possible properties and values from [string], [float64], [bool], [[]any],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesSourcesModifiedChangeNewArrayItemMapItem]
type VersionDiffResponseDifferencesSourcesModifiedChangeNewArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesSourcesModifiedChangeNewArrayItemMapItem any `json:",inline"`
	JSON                                                                     struct {
		OfString                                                                 respjson.Field
		OfFloat                                                                  respjson.Field
		OfBool                                                                   respjson.Field
		OfAnyArray                                                               respjson.Field
		OfVersionDiffResponseDifferencesSourcesModifiedChangeNewArrayItemMapItem respjson.Field
		raw                                                                      string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesSourcesModifiedChangeNewArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesSourcesModifiedChangeNewArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesSourcesModifiedChangeNewArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesSourcesModifiedChangeNewArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesSourcesModifiedChangeNewArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesSourcesModifiedChangeNewArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesSourcesModifiedChangeNewArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesSourcesModifiedChangeOldUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesSourcesModifiedChangeOldArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesSourcesModifiedChangeOldArray
// OfVersionDiffResponseDifferencesSourcesModifiedChangeOldMapItem]
type VersionDiffResponseDifferencesSourcesModifiedChangeOldUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesSourcesModifiedChangeOldArrayItemUnion] instead
	// of an object.
	OfVersionDiffResponseDifferencesSourcesModifiedChangeOldArray []VersionDiffResponseDifferencesSourcesModifiedChangeOldArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesSourcesModifiedChangeOldMapItem any `json:",inline"`
	JSON                                                            struct {
		OfString                                                        respjson.Field
		OfFloat                                                         respjson.Field
		OfBool                                                          respjson.Field
		OfVersionDiffResponseDifferencesSourcesModifiedChangeOldArray   respjson.Field
		OfVersionDiffResponseDifferencesSourcesModifiedChangeOldMapItem respjson.Field
		raw                                                             string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesSourcesModifiedChangeOldUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesSourcesModifiedChangeOldUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesSourcesModifiedChangeOldUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesSourcesModifiedChangeOldUnion) AsVersionDiffResponseDifferencesSourcesModifiedChangeOldArray() (v []VersionDiffResponseDifferencesSourcesModifiedChangeOldArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesSourcesModifiedChangeOldUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesSourcesModifiedChangeOldUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesSourcesModifiedChangeOldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesSourcesModifiedChangeOldArrayItemUnion contains
// all possible properties and values from [string], [float64], [bool], [[]any],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesSourcesModifiedChangeOldArrayItemMapItem]
type VersionDiffResponseDifferencesSourcesModifiedChangeOldArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesSourcesModifiedChangeOldArrayItemMapItem any `json:",inline"`
	JSON                                                                     struct {
		OfString                                                                 respjson.Field
		OfFloat                                                                  respjson.Field
		OfBool                                                                   respjson.Field
		OfAnyArray                                                               respjson.Field
		OfVersionDiffResponseDifferencesSourcesModifiedChangeOldArrayItemMapItem respjson.Field
		raw                                                                      string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesSourcesModifiedChangeOldArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesSourcesModifiedChangeOldArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesSourcesModifiedChangeOldArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesSourcesModifiedChangeOldArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesSourcesModifiedChangeOldArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesSourcesModifiedChangeOldArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesSourcesModifiedChangeOldArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesSourcesModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesSourcesRemoved) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesSourcesRemoved) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesTagManagers struct {
	Added []VersionDiffResponseDifferencesTagManagersAdded `json:"added" api:"required"`
	// Entities present in both snapshots but with at least one field changed. `old` is
	// the snapshot of the entity in the baseline (latest published); `new` is the
	// snapshot in the comparison target (the draft).
	Modified []VersionDiffResponseDifferencesTagManagersModified `json:"modified" api:"required"`
	Removed  []VersionDiffResponseDifferencesTagManagersRemoved  `json:"removed" api:"required"`
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
func (r VersionDiffResponseDifferencesTagManagers) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesTagManagers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesTagManagersAdded struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesTagManagersAdded) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesTagManagersAdded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesTagManagersModified struct {
	// Changed fields with credentials redacted. Nested objects use dot paths; arrays
	// are whole values. Missing values are null.
	Changes []VersionDiffResponseDifferencesTagManagersModifiedChange `json:"changes" api:"required"`
	New     VersionDiffResponseDifferencesTagManagersModifiedNew      `json:"new" api:"required"`
	Old     VersionDiffResponseDifferencesTagManagersModifiedOld      `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Changes     respjson.Field
		New         respjson.Field
		Old         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesTagManagersModified) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesTagManagersModified) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesTagManagersModifiedChange struct {
	New  VersionDiffResponseDifferencesTagManagersModifiedChangeNewUnion `json:"new" api:"required"`
	Old  VersionDiffResponseDifferencesTagManagersModifiedChangeOldUnion `json:"old" api:"required"`
	Path string                                                          `json:"path" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesTagManagersModifiedChange) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesTagManagersModifiedChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesTagManagersModifiedChangeNewUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesTagManagersModifiedChangeNewArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesTagManagersModifiedChangeNewArray
// OfVersionDiffResponseDifferencesTagManagersModifiedChangeNewMapItem]
type VersionDiffResponseDifferencesTagManagersModifiedChangeNewUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesTagManagersModifiedChangeNewArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesTagManagersModifiedChangeNewArray []VersionDiffResponseDifferencesTagManagersModifiedChangeNewArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesTagManagersModifiedChangeNewMapItem any `json:",inline"`
	JSON                                                                struct {
		OfString                                                            respjson.Field
		OfFloat                                                             respjson.Field
		OfBool                                                              respjson.Field
		OfVersionDiffResponseDifferencesTagManagersModifiedChangeNewArray   respjson.Field
		OfVersionDiffResponseDifferencesTagManagersModifiedChangeNewMapItem respjson.Field
		raw                                                                 string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesTagManagersModifiedChangeNewUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagersModifiedChangeNewUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagersModifiedChangeNewUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagersModifiedChangeNewUnion) AsVersionDiffResponseDifferencesTagManagersModifiedChangeNewArray() (v []VersionDiffResponseDifferencesTagManagersModifiedChangeNewArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagersModifiedChangeNewUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesTagManagersModifiedChangeNewUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesTagManagersModifiedChangeNewUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesTagManagersModifiedChangeNewArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesTagManagersModifiedChangeNewArrayItemMapItem]
type VersionDiffResponseDifferencesTagManagersModifiedChangeNewArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesTagManagersModifiedChangeNewArrayItemMapItem any `json:",inline"`
	JSON                                                                         struct {
		OfString                                                                     respjson.Field
		OfFloat                                                                      respjson.Field
		OfBool                                                                       respjson.Field
		OfAnyArray                                                                   respjson.Field
		OfVersionDiffResponseDifferencesTagManagersModifiedChangeNewArrayItemMapItem respjson.Field
		raw                                                                          string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesTagManagersModifiedChangeNewArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagersModifiedChangeNewArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagersModifiedChangeNewArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagersModifiedChangeNewArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagersModifiedChangeNewArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesTagManagersModifiedChangeNewArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesTagManagersModifiedChangeNewArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesTagManagersModifiedChangeOldUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesTagManagersModifiedChangeOldArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesTagManagersModifiedChangeOldArray
// OfVersionDiffResponseDifferencesTagManagersModifiedChangeOldMapItem]
type VersionDiffResponseDifferencesTagManagersModifiedChangeOldUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesTagManagersModifiedChangeOldArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesTagManagersModifiedChangeOldArray []VersionDiffResponseDifferencesTagManagersModifiedChangeOldArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesTagManagersModifiedChangeOldMapItem any `json:",inline"`
	JSON                                                                struct {
		OfString                                                            respjson.Field
		OfFloat                                                             respjson.Field
		OfBool                                                              respjson.Field
		OfVersionDiffResponseDifferencesTagManagersModifiedChangeOldArray   respjson.Field
		OfVersionDiffResponseDifferencesTagManagersModifiedChangeOldMapItem respjson.Field
		raw                                                                 string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesTagManagersModifiedChangeOldUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagersModifiedChangeOldUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagersModifiedChangeOldUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagersModifiedChangeOldUnion) AsVersionDiffResponseDifferencesTagManagersModifiedChangeOldArray() (v []VersionDiffResponseDifferencesTagManagersModifiedChangeOldArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagersModifiedChangeOldUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesTagManagersModifiedChangeOldUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesTagManagersModifiedChangeOldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesTagManagersModifiedChangeOldArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesTagManagersModifiedChangeOldArrayItemMapItem]
type VersionDiffResponseDifferencesTagManagersModifiedChangeOldArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesTagManagersModifiedChangeOldArrayItemMapItem any `json:",inline"`
	JSON                                                                         struct {
		OfString                                                                     respjson.Field
		OfFloat                                                                      respjson.Field
		OfBool                                                                       respjson.Field
		OfAnyArray                                                                   respjson.Field
		OfVersionDiffResponseDifferencesTagManagersModifiedChangeOldArrayItemMapItem respjson.Field
		raw                                                                          string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesTagManagersModifiedChangeOldArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagersModifiedChangeOldArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagersModifiedChangeOldArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagersModifiedChangeOldArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagersModifiedChangeOldArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesTagManagersModifiedChangeOldArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesTagManagersModifiedChangeOldArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesTagManagersModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesTagManagersModifiedNew) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesTagManagersModifiedNew) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesTagManagersModifiedOld struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesTagManagersModifiedOld) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesTagManagersModifiedOld) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesTagManagersRemoved struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesTagManagersRemoved) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesTagManagersRemoved) UnmarshalJSON(data []byte) error {
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// Changed fields with credentials redacted. Nested objects use dot paths; arrays
	// are whole values. Missing values are null.
	Changes []VersionDiffResponseDifferencesTagManagerTagsModifiedChange `json:"changes" api:"required"`
	New     VersionDiffResponseDifferencesTagManagerTagsModifiedNew      `json:"new" api:"required"`
	Old     VersionDiffResponseDifferencesTagManagerTagsModifiedOld      `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Changes     respjson.Field
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

type VersionDiffResponseDifferencesTagManagerTagsModifiedChange struct {
	New  VersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewUnion `json:"new" api:"required"`
	Old  VersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldUnion `json:"old" api:"required"`
	Path string                                                             `json:"path" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesTagManagerTagsModifiedChange) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesTagManagerTagsModifiedChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewArray
// OfVersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewMapItem]
type VersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewArray []VersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewMapItem any `json:",inline"`
	JSON                                                                   struct {
		OfString                                                               respjson.Field
		OfFloat                                                                respjson.Field
		OfBool                                                                 respjson.Field
		OfVersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewArray   respjson.Field
		OfVersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewMapItem respjson.Field
		raw                                                                    string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewUnion) AsVersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewArray() (v []VersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewArrayItemMapItem]
type VersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewArrayItemMapItem any `json:",inline"`
	JSON                                                                            struct {
		OfString                                                                        respjson.Field
		OfFloat                                                                         respjson.Field
		OfBool                                                                          respjson.Field
		OfAnyArray                                                                      respjson.Field
		OfVersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewArrayItemMapItem respjson.Field
		raw                                                                             string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesTagManagerTagsModifiedChangeNewArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldArray
// OfVersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldMapItem]
type VersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldArray []VersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldMapItem any `json:",inline"`
	JSON                                                                   struct {
		OfString                                                               respjson.Field
		OfFloat                                                                respjson.Field
		OfBool                                                                 respjson.Field
		OfVersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldArray   respjson.Field
		OfVersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldMapItem respjson.Field
		raw                                                                    string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldUnion) AsVersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldArray() (v []VersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldArrayItemMapItem]
type VersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldArrayItemMapItem any `json:",inline"`
	JSON                                                                            struct {
		OfString                                                                        respjson.Field
		OfFloat                                                                         respjson.Field
		OfBool                                                                          respjson.Field
		OfAnyArray                                                                      respjson.Field
		OfVersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldArrayItemMapItem respjson.Field
		raw                                                                             string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesTagManagerTagsModifiedChangeOldArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesTagManagerTagsModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// Changed fields with credentials redacted. Nested objects use dot paths; arrays
	// are whole values. Missing values are null.
	Changes []VersionDiffResponseDifferencesTagManagerTriggersModifiedChange `json:"changes" api:"required"`
	New     VersionDiffResponseDifferencesTagManagerTriggersModifiedNew      `json:"new" api:"required"`
	Old     VersionDiffResponseDifferencesTagManagerTriggersModifiedOld      `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Changes     respjson.Field
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

type VersionDiffResponseDifferencesTagManagerTriggersModifiedChange struct {
	New  VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewUnion `json:"new" api:"required"`
	Old  VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldUnion `json:"old" api:"required"`
	Path string                                                                 `json:"path" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesTagManagerTriggersModifiedChange) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesTagManagerTriggersModifiedChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewUnion contains
// all possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewArray
// OfVersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewMapItem]
type VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewArray []VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewMapItem any `json:",inline"`
	JSON                                                                       struct {
		OfString                                                                   respjson.Field
		OfFloat                                                                    respjson.Field
		OfBool                                                                     respjson.Field
		OfVersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewArray   respjson.Field
		OfVersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewMapItem respjson.Field
		raw                                                                        string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewUnion) AsVersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewArray() (v []VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewArrayItemMapItem]
type VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewArrayItemMapItem any `json:",inline"`
	JSON                                                                                struct {
		OfString                                                                            respjson.Field
		OfFloat                                                                             respjson.Field
		OfBool                                                                              respjson.Field
		OfAnyArray                                                                          respjson.Field
		OfVersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewArrayItemMapItem respjson.Field
		raw                                                                                 string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeNewArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldUnion contains
// all possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldArray
// OfVersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldMapItem]
type VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldArray []VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldMapItem any `json:",inline"`
	JSON                                                                       struct {
		OfString                                                                   respjson.Field
		OfFloat                                                                    respjson.Field
		OfBool                                                                     respjson.Field
		OfVersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldArray   respjson.Field
		OfVersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldMapItem respjson.Field
		raw                                                                        string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldUnion) AsVersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldArray() (v []VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldArrayItemMapItem]
type VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldArrayItemMapItem any `json:",inline"`
	JSON                                                                                struct {
		OfString                                                                            respjson.Field
		OfFloat                                                                             respjson.Field
		OfBool                                                                              respjson.Field
		OfAnyArray                                                                          respjson.Field
		OfVersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldArrayItemMapItem respjson.Field
		raw                                                                                 string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesTagManagerTriggersModifiedChangeOldArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesTagManagerTriggersModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// Changed fields with credentials redacted. Nested objects use dot paths; arrays
	// are whole values. Missing values are null.
	Changes []VersionDiffResponseDifferencesTagManagerVariablesModifiedChange `json:"changes" api:"required"`
	New     VersionDiffResponseDifferencesTagManagerVariablesModifiedNew      `json:"new" api:"required"`
	Old     VersionDiffResponseDifferencesTagManagerVariablesModifiedOld      `json:"old" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Changes     respjson.Field
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

type VersionDiffResponseDifferencesTagManagerVariablesModifiedChange struct {
	New  VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewUnion `json:"new" api:"required"`
	Old  VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldUnion `json:"old" api:"required"`
	Path string                                                                  `json:"path" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		New         respjson.Field
		Old         respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesTagManagerVariablesModifiedChange) RawJSON() string {
	return r.JSON.raw
}
func (r *VersionDiffResponseDifferencesTagManagerVariablesModifiedChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewUnion contains
// all possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewArray
// OfVersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewMapItem]
type VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewArray []VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewMapItem any `json:",inline"`
	JSON                                                                        struct {
		OfString                                                                    respjson.Field
		OfFloat                                                                     respjson.Field
		OfBool                                                                      respjson.Field
		OfVersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewArray   respjson.Field
		OfVersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewMapItem respjson.Field
		raw                                                                         string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewUnion) AsVersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewArray() (v []VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewArrayItemMapItem]
type VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewArrayItemMapItem any `json:",inline"`
	JSON                                                                                 struct {
		OfString                                                                             respjson.Field
		OfFloat                                                                              respjson.Field
		OfBool                                                                               respjson.Field
		OfAnyArray                                                                           respjson.Field
		OfVersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewArrayItemMapItem respjson.Field
		raw                                                                                  string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeNewArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldUnion contains
// all possible properties and values from [string], [float64], [bool],
// [[]VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldArrayItemUnion],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfVersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldArray
// OfVersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldMapItem]
type VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldArrayItemUnion]
	// instead of an object.
	OfVersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldArray []VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldArrayItemUnion `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldMapItem any `json:",inline"`
	JSON                                                                        struct {
		OfString                                                                    respjson.Field
		OfFloat                                                                     respjson.Field
		OfBool                                                                      respjson.Field
		OfVersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldArray   respjson.Field
		OfVersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldMapItem respjson.Field
		raw                                                                         string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldUnion) AsVersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldArray() (v []VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool],
// [[]any], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray
// OfVersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldArrayItemMapItem]
type VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfVersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldArrayItemMapItem any `json:",inline"`
	JSON                                                                                 struct {
		OfString                                                                             respjson.Field
		OfFloat                                                                              respjson.Field
		OfBool                                                                               respjson.Field
		OfAnyArray                                                                           respjson.Field
		OfVersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldArrayItemMapItem respjson.Field
		raw                                                                                  string
	} `json:"-"`
}

func (u VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldArrayItemUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldArrayItemUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *VersionDiffResponseDifferencesTagManagerVariablesModifiedChangeOldArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionDiffResponseDifferencesTagManagerVariablesModifiedNew struct {
	ID string `json:"id" api:"required"`
	// Human-readable label for the entity at the snapshot it was captured in. For most
	// collections this is the entity's `name` field; for `allowedEvents` it is a
	// computed summary of the event's key fields. Two `modified` items can therefore
	// have identical `old.name` and `new.name` even though their underlying records
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
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
	// differ — the change is in fields not surfaced by the summary. Read `changes` for
	// the changed fields.
	Name string `json:"name" api:"nullable"`
	// Optional change-classifier. Currently set to `'Reordered'` on `mappings`
	// modifications when the only change is priority reordering — this lets clients
	// de-emphasize them in change-review UI. `null` for every other diff item.
	Summary string `json:"summary" api:"nullable"`
	// Parent tag-manager id for `tagManagerTags`, `tagManagerTriggers`, and
	// `tagManagerVariables`. `null` for every other collection.
	TagManagerID string `json:"tagManagerId" api:"nullable"`
	// Entity subtype when the collection has one — for `sources` it matches the
	// source's `type` field, letting clients distinguish kinds of source. `null` for
	// collections without a subtype.
	Type string `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Name         respjson.Field
		Summary      respjson.Field
		TagManagerID respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionDiffResponseDifferencesTagManagerVariablesRemoved) RawJSON() string { return r.JSON.raw }
func (r *VersionDiffResponseDifferencesTagManagerVariablesRemoved) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionStatusResponse struct {
	// Echo of the requested entity id.
	ID string `json:"id" api:"required"`
	// Echo of the requested collection.
	//
	// Any of "allowedEvents", "consentSettings", "dataGovernanceEvents",
	// "dataGovernanceRules", "destinations", "experimentSettings",
	// "externalAllowedEventData", "globalDispatchCenters", "mappings",
	// "replaySettings", "sources", "tagManagerTags", "tagManagerTriggers",
	// "tagManagerVariables", "tagManagers".
	Collection VersionStatusResponseCollection `json:"collection" api:"required"`
	// How the entity differs from the latest published version: `added` exists only in
	// the draft, `modified` exists in both with unpublished edits, `removed` exists
	// only in the published version. `null` covers every other outcome — the entity
	// matches the published version, is unknown, belongs to another account, or the
	// account has never published — so callers cannot use it to probe for entities
	// outside their account.
	//
	// Any of "added", "modified", "removed".
	Status VersionStatusResponseStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Collection  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *VersionStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Echo of the requested collection.
type VersionStatusResponseCollection string

const (
	VersionStatusResponseCollectionAllowedEvents            VersionStatusResponseCollection = "allowedEvents"
	VersionStatusResponseCollectionConsentSettings          VersionStatusResponseCollection = "consentSettings"
	VersionStatusResponseCollectionDataGovernanceEvents     VersionStatusResponseCollection = "dataGovernanceEvents"
	VersionStatusResponseCollectionDataGovernanceRules      VersionStatusResponseCollection = "dataGovernanceRules"
	VersionStatusResponseCollectionDestinations             VersionStatusResponseCollection = "destinations"
	VersionStatusResponseCollectionExperimentSettings       VersionStatusResponseCollection = "experimentSettings"
	VersionStatusResponseCollectionExternalAllowedEventData VersionStatusResponseCollection = "externalAllowedEventData"
	VersionStatusResponseCollectionGlobalDispatchCenters    VersionStatusResponseCollection = "globalDispatchCenters"
	VersionStatusResponseCollectionMappings                 VersionStatusResponseCollection = "mappings"
	VersionStatusResponseCollectionReplaySettings           VersionStatusResponseCollection = "replaySettings"
	VersionStatusResponseCollectionSources                  VersionStatusResponseCollection = "sources"
	VersionStatusResponseCollectionTagManagerTags           VersionStatusResponseCollection = "tagManagerTags"
	VersionStatusResponseCollectionTagManagerTriggers       VersionStatusResponseCollection = "tagManagerTriggers"
	VersionStatusResponseCollectionTagManagerVariables      VersionStatusResponseCollection = "tagManagerVariables"
	VersionStatusResponseCollectionTagManagers              VersionStatusResponseCollection = "tagManagers"
)

// How the entity differs from the latest published version: `added` exists only in
// the draft, `modified` exists in both with unpublished edits, `removed` exists
// only in the published version. `null` covers every other outcome — the entity
// matches the published version, is unknown, belongs to another account, or the
// account has never published — so callers cannot use it to probe for entities
// outside their account.
type VersionStatusResponseStatus string

const (
	VersionStatusResponseStatusAdded    VersionStatusResponseStatus = "added"
	VersionStatusResponseStatusModified VersionStatusResponseStatus = "modified"
	VersionStatusResponseStatusRemoved  VersionStatusResponseStatus = "removed"
)

type VersionRevertResponse struct {
	// Newly-added draft entities that were deleted.
	DiscardedCount float64 `json:"discardedCount" api:"required"`
	// Requested entities present in neither the draft nor the latest published
	// version.
	NotFoundCount float64 `json:"notFoundCount" api:"required"`
	// Total pending draft changes remaining after this operation.
	RemainingChanges float64 `json:"remainingChanges" api:"required"`
	// Entities recreated because they had been deleted from the draft.
	RestoredCount float64 `json:"restoredCount" api:"required"`
	// Entities overwritten back to their published values.
	RevertedCount float64 `json:"revertedCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DiscardedCount   respjson.Field
		NotFoundCount    respjson.Field
		RemainingChanges respjson.Field
		RestoredCount    respjson.Field
		RevertedCount    respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionRevertResponse) RawJSON() string { return r.JSON.raw }
func (r *VersionRevertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionAbandonResponse struct {
	// Newly-added draft entities that were deleted.
	DiscardedCount float64 `json:"discardedCount" api:"required"`
	// Requested entities present in neither the draft nor the latest published
	// version.
	NotFoundCount float64 `json:"notFoundCount" api:"required"`
	// Total pending draft changes remaining after this operation.
	RemainingChanges float64 `json:"remainingChanges" api:"required"`
	// Entities recreated because they had been deleted from the draft.
	RestoredCount float64 `json:"restoredCount" api:"required"`
	// Entities overwritten back to their published values.
	RevertedCount float64 `json:"revertedCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DiscardedCount   respjson.Field
		NotFoundCount    respjson.Field
		RemainingChanges respjson.Field
		RestoredCount    respjson.Field
		RevertedCount    respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionAbandonResponse) RawJSON() string { return r.JSON.raw }
func (r *VersionAbandonResponse) UnmarshalJSON(data []byte) error {
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
	// Allow publishing web sources without allowed domains for this version. Defaults
	// to false. Sources without allowed domains accept events from any domain.
	BypassWebSourceDomainValidation param.Opt[bool] `json:"bypassWebSourceDomainValidation,omitzero"`
	// Optional human-readable label for the new version.
	Name param.Opt[string] `json:"name,omitzero"`
	// Optional release notes describing what changed in this version.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Cherry-pick: allowed event ids (slug-like strings) to include from the draft.
	// Omit every `include*` field (or send each as `[]`) to publish all draft changes.
	// Once any `include*` field lists ids, only the listed changes publish, and every
	// collection you omit or send as `[]` keeps its published state.
	IncludeAllowedEvents []string `json:"includeAllowedEvents,omitzero"`
	// Cherry-pick: consent settings ids to include from the draft. Omit every
	// `include*` field (or send each as `[]`) to publish all draft changes. Once any
	// `include*` field lists ids, only the listed changes publish, and every
	// collection you omit or send as `[]` keeps its published state.
	IncludeConsentSettings []string `json:"includeConsentSettings,omitzero"`
	// Cherry-pick: data governance event UUIDs to include from the draft. Omit every
	// `include*` field (or send each as `[]`) to publish all draft changes. Once any
	// `include*` field lists ids, only the listed changes publish, and every
	// collection you omit or send as `[]` keeps its published state.
	IncludeDataGovernanceEvents []string `json:"includeDataGovernanceEvents,omitzero" format:"uuid"`
	// Cherry-pick: data governance rule UUIDs to include from the draft. Omit every
	// `include*` field (or send each as `[]`) to publish all draft changes. Once any
	// `include*` field lists ids, only the listed changes publish, and every
	// collection you omit or send as `[]` keeps its published state.
	IncludeDataGovernanceRules []string `json:"includeDataGovernanceRules,omitzero" format:"uuid"`
	// Cherry-pick: destination UUIDs to include from the draft. Omit every `include*`
	// field (or send each as `[]`) to publish all draft changes. Once any `include*`
	// field lists ids, only the listed changes publish, and every collection you omit
	// or send as `[]` keeps its published state.
	IncludeDestinations []string `json:"includeDestinations,omitzero" format:"uuid"`
	// Cherry-pick: experiment UUIDs to include from the draft. Omit every `include*`
	// field (or send each as `[]`) to publish all draft changes. Once any `include*`
	// field lists ids, only the listed changes publish, and every collection you omit
	// or send as `[]` keeps its published state.
	IncludeExperiments []string `json:"includeExperiments,omitzero" format:"uuid"`
	// Cherry-pick: experiment settings UUIDs to include from the draft. Omit every
	// `include*` field (or send each as `[]`) to publish all draft changes. Once any
	// `include*` field lists ids, only the listed changes publish, and every
	// collection you omit or send as `[]` keeps its published state.
	IncludeExperimentSettings []string `json:"includeExperimentSettings,omitzero" format:"uuid"`
	// Cherry-pick: experiment variant UUIDs to include from the draft. Omit every
	// `include*` field (or send each as `[]`) to publish all draft changes. Once any
	// `include*` field lists ids, only the listed changes publish, and every
	// collection you omit or send as `[]` keeps its published state.
	IncludeExperimentVariants []string `json:"includeExperimentVariants,omitzero" format:"uuid"`
	// Cherry-pick: external allowed event data ids to include from the draft. Omit
	// every `include*` field (or send each as `[]`) to publish all draft changes. Once
	// any `include*` field lists ids, only the listed changes publish, and every
	// collection you omit or send as `[]` keeps its published state.
	IncludeExternalAllowedEventData []string `json:"includeExternalAllowedEventData,omitzero"`
	// Cherry-pick: global dispatch center ids to include from the draft. Omit every
	// `include*` field (or send each as `[]`) to publish all draft changes. Once any
	// `include*` field lists ids, only the listed changes publish, and every
	// collection you omit or send as `[]` keeps its published state.
	IncludeGlobalDispatchCenters []string `json:"includeGlobalDispatchCenters,omitzero"`
	// Cherry-pick: mapping ids to include from the draft. Omit every `include*` field
	// (or send each as `[]`) to publish all draft changes. Once any `include*` field
	// lists ids, only the listed changes publish, and every collection you omit or
	// send as `[]` keeps its published state.
	IncludeMappings []string `json:"includeMappings,omitzero"`
	// Cherry-pick: replay settings ids to include from the draft. Omit every
	// `include*` field (or send each as `[]`) to publish all draft changes. Once any
	// `include*` field lists ids, only the listed changes publish, and every
	// collection you omit or send as `[]` keeps its published state.
	IncludeReplaySettings []string `json:"includeReplaySettings,omitzero"`
	// Cherry-pick: source UUIDs to include from the draft. Omit every `include*` field
	// (or send each as `[]`) to publish all draft changes. Once any `include*` field
	// lists ids, only the listed changes publish, and every collection you omit or
	// send as `[]` keeps its published state.
	IncludeSources []string `json:"includeSources,omitzero" format:"uuid"`
	// Cherry-pick: tag manager container UUIDs to include from the draft. Selecting a
	// container publishes its own settings only; list its tags, triggers, and
	// variables to publish their changes. Selecting a container deleted from the draft
	// also removes its tags, triggers, and variables. Omit every `include*` field (or
	// send each as `[]`) to publish all draft changes. Once any `include*` field lists
	// ids, only the listed changes publish, and every collection you omit or send as
	// `[]` keeps its published state.
	IncludeTagManagers []string `json:"includeTagManagers,omitzero" format:"uuid"`
	// Cherry-pick: tag manager tag UUIDs to include from the draft. If the item's
	// container has never been published, the container publishes with it. Omit every
	// `include*` field (or send each as `[]`) to publish all draft changes. Once any
	// `include*` field lists ids, only the listed changes publish, and every
	// collection you omit or send as `[]` keeps its published state.
	IncludeTagManagerTags []string `json:"includeTagManagerTags,omitzero" format:"uuid"`
	// Cherry-pick: tag manager trigger UUIDs to include from the draft. If the item's
	// container has never been published, the container publishes with it. Omit every
	// `include*` field (or send each as `[]`) to publish all draft changes. Once any
	// `include*` field lists ids, only the listed changes publish, and every
	// collection you omit or send as `[]` keeps its published state.
	IncludeTagManagerTriggers []string `json:"includeTagManagerTriggers,omitzero" format:"uuid"`
	// Cherry-pick: tag manager variable UUIDs to include from the draft. If the item's
	// container has never been published, the container publishes with it. Omit every
	// `include*` field (or send each as `[]`) to publish all draft changes. Once any
	// `include*` field lists ids, only the listed changes publish, and every
	// collection you omit or send as `[]` keeps its published state.
	IncludeTagManagerVariables []string `json:"includeTagManagerVariables,omitzero" format:"uuid"`
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
	// Baseline version id to compare the path version against. Omit or pass null for
	// the latest published version. Pass a version UUID to compute a
	// version-vs-version diff.
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

type VersionStatusParams struct {
	// Entity collection the id belongs to. `experiments` and `experimentVariants` are
	// not supported — draft experiments never snapshot into a version — but
	// `experimentSettings` is.
	//
	// Any of "allowedEvents", "consentSettings", "dataGovernanceEvents",
	// "dataGovernanceRules", "destinations", "experimentSettings",
	// "externalAllowedEventData", "globalDispatchCenters", "mappings",
	// "replaySettings", "sources", "tagManagerTags", "tagManagerTriggers",
	// "tagManagerVariables", "tagManagers".
	Collection VersionStatusParamsCollection `query:"collection,omitzero" api:"required" json:"-"`
	// Id of the entity to check. Most collections use UUIDs; `allowedEvents`,
	// `externalAllowedEventData`, `mappings`, `consentSettings`,
	// `globalDispatchCenters`, and `replaySettings` use slug-like strings.
	EntityID string `query:"entityId" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [VersionStatusParams]'s query parameters as `url.Values`.
func (r VersionStatusParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VersionStatusParamsID string

const (
	VersionStatusParamsIDDraft VersionStatusParamsID = "draft"
)

// Entity collection the id belongs to. `experiments` and `experimentVariants` are
// not supported — draft experiments never snapshot into a version — but
// `experimentSettings` is.
type VersionStatusParamsCollection string

const (
	VersionStatusParamsCollectionAllowedEvents            VersionStatusParamsCollection = "allowedEvents"
	VersionStatusParamsCollectionConsentSettings          VersionStatusParamsCollection = "consentSettings"
	VersionStatusParamsCollectionDataGovernanceEvents     VersionStatusParamsCollection = "dataGovernanceEvents"
	VersionStatusParamsCollectionDataGovernanceRules      VersionStatusParamsCollection = "dataGovernanceRules"
	VersionStatusParamsCollectionDestinations             VersionStatusParamsCollection = "destinations"
	VersionStatusParamsCollectionExperimentSettings       VersionStatusParamsCollection = "experimentSettings"
	VersionStatusParamsCollectionExternalAllowedEventData VersionStatusParamsCollection = "externalAllowedEventData"
	VersionStatusParamsCollectionGlobalDispatchCenters    VersionStatusParamsCollection = "globalDispatchCenters"
	VersionStatusParamsCollectionMappings                 VersionStatusParamsCollection = "mappings"
	VersionStatusParamsCollectionReplaySettings           VersionStatusParamsCollection = "replaySettings"
	VersionStatusParamsCollectionSources                  VersionStatusParamsCollection = "sources"
	VersionStatusParamsCollectionTagManagerTags           VersionStatusParamsCollection = "tagManagerTags"
	VersionStatusParamsCollectionTagManagerTriggers       VersionStatusParamsCollection = "tagManagerTriggers"
	VersionStatusParamsCollectionTagManagerVariables      VersionStatusParamsCollection = "tagManagerVariables"
	VersionStatusParamsCollectionTagManagers              VersionStatusParamsCollection = "tagManagers"
)

type VersionRevertParams struct {
	// Draft entities to revert back to the latest published version. Each entry is
	// processed independently: a modified entity is restored to its published values,
	// a deleted entity is recreated, and a newly-added entity is discarded. Pass a
	// single entry to revert one change.
	Entities []VersionRevertParamsEntity `json:"entities,omitzero" api:"required"`
	paramObj
}

func (r VersionRevertParams) MarshalJSON() (data []byte, err error) {
	type shadow VersionRevertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VersionRevertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionRevertParamsID string

const (
	VersionRevertParamsIDDraft VersionRevertParamsID = "draft"
)

// The properties ID, Collection are required.
type VersionRevertParamsEntity struct {
	// Id of the draft entity to revert.
	ID string `json:"id" api:"required"`
	// Entity collection the id belongs to. `experiments` and `experimentVariants` are
	// not supported — use the experiment lifecycle endpoints for those — but
	// `experimentSettings` is.
	//
	// Any of "allowedEvents", "consentSettings", "dataGovernanceEvents",
	// "dataGovernanceRules", "destinations", "experimentSettings",
	// "externalAllowedEventData", "globalDispatchCenters", "mappings",
	// "replaySettings", "sources", "tagManagerTags", "tagManagerTriggers",
	// "tagManagerVariables", "tagManagers".
	Collection string `json:"collection,omitzero" api:"required"`
	paramObj
}

func (r VersionRevertParamsEntity) MarshalJSON() (data []byte, err error) {
	type shadow VersionRevertParamsEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VersionRevertParamsEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VersionRevertParamsEntity](
		"collection", "allowedEvents", "consentSettings", "dataGovernanceEvents", "dataGovernanceRules", "destinations", "experimentSettings", "externalAllowedEventData", "globalDispatchCenters", "mappings", "replaySettings", "sources", "tagManagerTags", "tagManagerTriggers", "tagManagerVariables", "tagManagers",
	)
}

type VersionAbandonParamsID string

const (
	VersionAbandonParamsIDDraft VersionAbandonParamsID = "draft"
)
