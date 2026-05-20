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

// AllowedEventService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAllowedEventService] method instead.
type AllowedEventService struct {
	Options []option.RequestOption
}

// NewAllowedEventService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAllowedEventService(opts ...option.RequestOption) (r AllowedEventService) {
	r = AllowedEventService{}
	r.Options = opts
	return
}

// List every allowed event for this account. Allowed events sit between sources
// and destinations in the dispatch flow — only inbound events whose `event` field
// matches the `name` of an allowed event (case-insensitive) can be routed to that
// event's `destinationIds`. Events without a matching allowed event are dropped.
// The list is not paginated; the per-account count is bounded. System events
// (names beginning with `$`, e.g. `$heatmap_click`) are hidden from the response —
// only `$identify` is creatable as an allowed event. Requires scope:
// allowedEvent:list
func (r *AllowedEventService) List(ctx context.Context, opts ...option.RequestOption) (res *AllowedEventListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/allowed-events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Create a new allowed event for this account.
//
//   - `name` is required, trimmed, case-insensitively unique within the account, and
//     rejected if it exceeds the platform event-name length limit.
//   - Names starting with `$` are reserved for system events. Only `$identify` is
//     accepted.
//   - `destinationIds` is optional. Unknown ids and ids belonging to other accounts
//     are silently filtered out at write time (the destination must exist on this
//     account to be saved).
//
// Returns the full entity so callers can read the server-assigned `id`,
// `createdAt`, and the filtered `destinationIds` without a follow-up GET. Known
// input failures (duplicate name, name length, `$`-prefix reservation, empty name)
// are returned as HTTP 409 with the reason in the response `error` field. Requires
// scope: allowedEvent:create
func (r *AllowedEventService) New(ctx context.Context, body AllowedEventNewParams, opts ...option.RequestOption) (res *AllowedEventNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/allowed-events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Fetch a single allowed event by id. Returns 404 when no record matches the
// supplied id or it belongs to a different account. Requires scope:
// allowedEvent:find
func (r *AllowedEventService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AllowedEventGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/allowed-events/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Partially update an allowed event. Only the fields you send are changed; omitted
// fields keep their current value.
//
//   - `destinationIds` is replaced wholesale when sent — the canonical way to add or
//     remove a destination is to fetch, modify the array, and PATCH it back. Stale
//     ids (deleted destinations or destinations on another account) are silently
//     filtered out at write time.
//   - `name` is subject to the same rules as create: case-insensitive uniqueness,
//     length cap, `$`-prefix reservation.
//   - `trigger` accepts `null` to clear the existing value.
//
// Returns the full entity. Known input failures (duplicate name, length,
// `$`-prefix, empty name) are returned as HTTP 409 with the reason in the response
// `error` field. Requires scope: allowedEvent:update
func (r *AllowedEventService) Update(ctx context.Context, id string, body AllowedEventUpdateParams, opts ...option.RequestOption) (res *AllowedEventUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/allowed-events/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete an allowed event. After deletion, inbound events whose `event` field
// matches the deleted name are no longer routed and are dropped at the allow-list
// stage of the dispatch flow. Requires scope: allowedEvent:delete
func (r *AllowedEventService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/allowed-events/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type AllowedEventListResponse struct {
	Entities []AllowedEventListResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AllowedEventListResponse) RawJSON() string { return r.JSON.raw }
func (r *AllowedEventListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AllowedEventListResponseEntity struct {
	// Server-assigned UUID for this allowed event.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp when the allowed event was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Destinations that receive this event. Empty array means the event is allowed but
	// routed nowhere (effectively dropped). PATCH replaces this list wholesale.
	DestinationIDs []string `json:"destinationIds" api:"required"`
	// Case-insensitive event name. Inbound events whose `event` field matches this
	// value are gated through this allowed event. Reserved: names starting with `$`
	// are system events and cannot be created (except `$identify`).
	Name string `json:"name" api:"required"`
	// Optional free-form trigger description. Not used by the dispatch pipeline —
	// surfaced in the dashboard so teams can record where each event fires.
	Trigger string `json:"trigger" api:"nullable"`
	// ISO 8601 timestamp of the last PATCH. Equal to createdAt on a freshly created
	// event.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		DestinationIDs respjson.Field
		Name           respjson.Field
		Trigger        respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AllowedEventListResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *AllowedEventListResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AllowedEventNewResponse struct {
	// Server-assigned UUID for this allowed event.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp when the allowed event was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Destinations that receive this event. Empty array means the event is allowed but
	// routed nowhere (effectively dropped). PATCH replaces this list wholesale.
	DestinationIDs []string `json:"destinationIds" api:"required"`
	// Case-insensitive event name. Inbound events whose `event` field matches this
	// value are gated through this allowed event. Reserved: names starting with `$`
	// are system events and cannot be created (except `$identify`).
	Name string `json:"name" api:"required"`
	// ISO 8601 timestamp of the most recent successful dispatch to any destination on
	// `destinationIds`. Lags `lastTriggeredAt` when consent or governance rules drop
	// the event before dispatch.
	LastDispatchedAt string `json:"lastDispatchedAt" api:"nullable"`
	// ISO 8601 timestamp of the most recent inbound event observed for this name.
	// Useful for spotting events that were configured but never fired. Null when never
	// observed.
	LastTriggeredAt string `json:"lastTriggeredAt" api:"nullable"`
	// Optional free-form trigger description. Not used by the dispatch pipeline —
	// surfaced in the dashboard so teams can record where each event fires.
	Trigger string `json:"trigger" api:"nullable"`
	// ISO 8601 timestamp of the last PATCH. Equal to createdAt on a freshly created
	// event.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		DestinationIDs   respjson.Field
		Name             respjson.Field
		LastDispatchedAt respjson.Field
		LastTriggeredAt  respjson.Field
		Trigger          respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AllowedEventNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AllowedEventNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AllowedEventGetResponse struct {
	// Server-assigned UUID for this allowed event.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp when the allowed event was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Destinations that receive this event. Empty array means the event is allowed but
	// routed nowhere (effectively dropped). PATCH replaces this list wholesale.
	DestinationIDs []string `json:"destinationIds" api:"required"`
	// Case-insensitive event name. Inbound events whose `event` field matches this
	// value are gated through this allowed event. Reserved: names starting with `$`
	// are system events and cannot be created (except `$identify`).
	Name string `json:"name" api:"required"`
	// ISO 8601 timestamp of the most recent successful dispatch to any destination on
	// `destinationIds`. Lags `lastTriggeredAt` when consent or governance rules drop
	// the event before dispatch.
	LastDispatchedAt string `json:"lastDispatchedAt" api:"nullable"`
	// ISO 8601 timestamp of the most recent inbound event observed for this name.
	// Useful for spotting events that were configured but never fired. Null when never
	// observed.
	LastTriggeredAt string `json:"lastTriggeredAt" api:"nullable"`
	// Optional free-form trigger description. Not used by the dispatch pipeline —
	// surfaced in the dashboard so teams can record where each event fires.
	Trigger string `json:"trigger" api:"nullable"`
	// ISO 8601 timestamp of the last PATCH. Equal to createdAt on a freshly created
	// event.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		DestinationIDs   respjson.Field
		Name             respjson.Field
		LastDispatchedAt respjson.Field
		LastTriggeredAt  respjson.Field
		Trigger          respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AllowedEventGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AllowedEventGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AllowedEventUpdateResponse struct {
	// Server-assigned UUID for this allowed event.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp when the allowed event was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Destinations that receive this event. Empty array means the event is allowed but
	// routed nowhere (effectively dropped). PATCH replaces this list wholesale.
	DestinationIDs []string `json:"destinationIds" api:"required"`
	// Case-insensitive event name. Inbound events whose `event` field matches this
	// value are gated through this allowed event. Reserved: names starting with `$`
	// are system events and cannot be created (except `$identify`).
	Name string `json:"name" api:"required"`
	// ISO 8601 timestamp of the most recent successful dispatch to any destination on
	// `destinationIds`. Lags `lastTriggeredAt` when consent or governance rules drop
	// the event before dispatch.
	LastDispatchedAt string `json:"lastDispatchedAt" api:"nullable"`
	// ISO 8601 timestamp of the most recent inbound event observed for this name.
	// Useful for spotting events that were configured but never fired. Null when never
	// observed.
	LastTriggeredAt string `json:"lastTriggeredAt" api:"nullable"`
	// Optional free-form trigger description. Not used by the dispatch pipeline —
	// surfaced in the dashboard so teams can record where each event fires.
	Trigger string `json:"trigger" api:"nullable"`
	// ISO 8601 timestamp of the last PATCH. Equal to createdAt on a freshly created
	// event.
	UpdatedAt string `json:"updatedAt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		DestinationIDs   respjson.Field
		Name             respjson.Field
		LastDispatchedAt respjson.Field
		LastTriggeredAt  respjson.Field
		Trigger          respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AllowedEventUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *AllowedEventUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AllowedEventNewParams struct {
	Name           string   `json:"name" api:"required"`
	DestinationIDs []string `json:"destinationIds,omitzero"`
	paramObj
}

func (r AllowedEventNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AllowedEventNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AllowedEventNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AllowedEventUpdateParams struct {
	// New event name. Subject to the same rules as create: case-insensitive uniqueness
	// within the account, max length enforced by the platform, and the `$`-prefix
	// reservation (only `$identify` is allowed). Omit to leave the name unchanged.
	Name param.Opt[string] `json:"name,omitzero"`
	// Free-form trigger description shown in the dashboard. Send `null` to clear. Not
	// used by the dispatch pipeline.
	Trigger param.Opt[string] `json:"trigger,omitzero"`
	// Destinations that should receive this event. Wholesale replacement — the sent
	// list becomes the new value. Stale IDs (destinations from another account or
	// deleted destinations) are silently filtered out at write time. Send `[]` to gate
	// the event from every destination.
	DestinationIDs []string `json:"destinationIds,omitzero"`
	paramObj
}

func (r AllowedEventUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AllowedEventUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AllowedEventUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
