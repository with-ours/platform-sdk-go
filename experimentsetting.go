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

// ExperimentSettingService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExperimentSettingService] method instead.
type ExperimentSettingService struct {
	Options []option.RequestOption
}

// NewExperimentSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewExperimentSettingService(opts ...option.RequestOption) (r ExperimentSettingService) {
	r = ExperimentSettingService{}
	r.Options = opts
	return
}

// List experiment settings records for the account. Use the returned `id` as
// `experimentSettingsId` when creating an experiment. Requires scope:
// experimentSettings:list
func (r *ExperimentSettingService) List(ctx context.Context, opts ...option.RequestOption) (res *ExperimentSettingListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/experiment-settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Create the account-level experimentation bootstrap record. Most accounts should
// only ever have one. Requires scope: experimentSettings:create
func (r *ExperimentSettingService) New(ctx context.Context, body ExperimentSettingNewParams, opts ...option.RequestOption) (res *ExperimentSettingNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/experiment-settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Find a single experiment settings record by ID. Returns 404 when no record
// matches the supplied id. Requires scope: experimentSettings:find
func (r *ExperimentSettingService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ExperimentSettingGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/experiment-settings/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Partially update an experiment settings. Only the fields you send are changed.
// Requires scope: experimentSettings:update
func (r *ExperimentSettingService) Update(ctx context.Context, id string, body ExperimentSettingUpdateParams, opts ...option.RequestOption) (res *ExperimentSettingUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/experiment-settings/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Delete the experimentation bootstrap record. This also deletes child
// experiments, variants, and personalization properties owned by it. Requires
// scope: experimentSettings:delete
func (r *ExperimentSettingService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *ExperimentSettingDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/experiment-settings/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type ExperimentSettingListResponse struct {
	// Experiment settings records available to the current account. Use the `id` from
	// this response as `experimentSettingsId` when creating an experiment. Most
	// accounts have a single record; this list is not paginated.
	Entities []ExperimentSettingListResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentSettingListResponse) RawJSON() string { return r.JSON.raw }
func (r *ExperimentSettingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentSettingListResponseEntity struct {
	// Unique identifier for the account-level experiment settings record.
	ID string `json:"id" api:"required"`
	// Account that owns this experiment settings record.
	AccountID string `json:"accountId" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	// Human-readable name for this experimentation configuration.
	Name string `json:"name" api:"required"`
	// Pixel token used by the experiments runtime and CDN configuration. This is
	// informative for REST clients; use the settings `id` for createExperiment.
	Pixel string `json:"pixel" api:"required"`
	// Cookie name used to persist sticky variant assignments in the browser.
	CookieName      string `json:"cookieName" api:"nullable"`
	CreatedByUserID string `json:"createdByUserId" api:"nullable"`
	UpdatedAt       string `json:"updatedAt" api:"nullable"`
	UpdatedByUserID string `json:"updatedByUserId" api:"nullable"`
	// Limits which domains can load your experiments. When set, experiments using this
	// settings record are only served on these domains; the SDK refuses to load
	// anywhere else and your experiments never run on those hosts. Separate from
	// source `whitelistDomains`, which limits which domains can send events to the
	// CDP.
	WhitelistDomains []string `json:"whitelistDomains" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AccountID        respjson.Field
		CreatedAt        respjson.Field
		Name             respjson.Field
		Pixel            respjson.Field
		CookieName       respjson.Field
		CreatedByUserID  respjson.Field
		UpdatedAt        respjson.Field
		UpdatedByUserID  respjson.Field
		WhitelistDomains respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentSettingListResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *ExperimentSettingListResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentSettingNewResponse struct {
	// Unique identifier for the account-level experiment settings record.
	ID string `json:"id" api:"required"`
	// Account that owns this experiment settings record.
	AccountID string `json:"accountId" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	// Human-readable name for this experimentation configuration.
	Name string `json:"name" api:"required"`
	// Pixel token used by the experiments runtime and CDN configuration. This is
	// informative for REST clients; use the settings `id` for createExperiment.
	Pixel string `json:"pixel" api:"required"`
	// Cookie name used to persist sticky variant assignments in the browser.
	CookieName      string `json:"cookieName" api:"nullable"`
	CreatedByUserID string `json:"createdByUserId" api:"nullable"`
	UpdatedAt       string `json:"updatedAt" api:"nullable"`
	UpdatedByUserID string `json:"updatedByUserId" api:"nullable"`
	// Limits which domains can load your experiments. When set, experiments using this
	// settings record are only served on these domains; the SDK refuses to load
	// anywhere else and your experiments never run on those hosts. Separate from
	// source `whitelistDomains`, which limits which domains can send events to the
	// CDP.
	WhitelistDomains []string `json:"whitelistDomains" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AccountID        respjson.Field
		CreatedAt        respjson.Field
		Name             respjson.Field
		Pixel            respjson.Field
		CookieName       respjson.Field
		CreatedByUserID  respjson.Field
		UpdatedAt        respjson.Field
		UpdatedByUserID  respjson.Field
		WhitelistDomains respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentSettingNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ExperimentSettingNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentSettingGetResponse struct {
	// Unique identifier for the account-level experiment settings record.
	ID string `json:"id" api:"required"`
	// Account that owns this experiment settings record.
	AccountID string `json:"accountId" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	// Human-readable name for this experimentation configuration.
	Name string `json:"name" api:"required"`
	// Pixel token used by the experiments runtime and CDN configuration. This is
	// informative for REST clients; use the settings `id` for createExperiment.
	Pixel string `json:"pixel" api:"required"`
	// Cookie name used to persist sticky variant assignments in the browser.
	CookieName      string `json:"cookieName" api:"nullable"`
	CreatedByUserID string `json:"createdByUserId" api:"nullable"`
	UpdatedAt       string `json:"updatedAt" api:"nullable"`
	UpdatedByUserID string `json:"updatedByUserId" api:"nullable"`
	// Limits which domains can load your experiments. When set, experiments using this
	// settings record are only served on these domains; the SDK refuses to load
	// anywhere else and your experiments never run on those hosts. Separate from
	// source `whitelistDomains`, which limits which domains can send events to the
	// CDP.
	WhitelistDomains []string `json:"whitelistDomains" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AccountID        respjson.Field
		CreatedAt        respjson.Field
		Name             respjson.Field
		Pixel            respjson.Field
		CookieName       respjson.Field
		CreatedByUserID  respjson.Field
		UpdatedAt        respjson.Field
		UpdatedByUserID  respjson.Field
		WhitelistDomains respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentSettingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ExperimentSettingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentSettingUpdateResponse struct {
	// Unique identifier for the account-level experiment settings record.
	ID string `json:"id" api:"required"`
	// Account that owns this experiment settings record.
	AccountID string `json:"accountId" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	// Human-readable name for this experimentation configuration.
	Name string `json:"name" api:"required"`
	// Pixel token used by the experiments runtime and CDN configuration. This is
	// informative for REST clients; use the settings `id` for createExperiment.
	Pixel string `json:"pixel" api:"required"`
	// Cookie name used to persist sticky variant assignments in the browser.
	CookieName      string `json:"cookieName" api:"nullable"`
	CreatedByUserID string `json:"createdByUserId" api:"nullable"`
	UpdatedAt       string `json:"updatedAt" api:"nullable"`
	UpdatedByUserID string `json:"updatedByUserId" api:"nullable"`
	// Limits which domains can load your experiments. When set, experiments using this
	// settings record are only served on these domains; the SDK refuses to load
	// anywhere else and your experiments never run on those hosts. Separate from
	// source `whitelistDomains`, which limits which domains can send events to the
	// CDP.
	WhitelistDomains []string `json:"whitelistDomains" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AccountID        respjson.Field
		CreatedAt        respjson.Field
		Name             respjson.Field
		Pixel            respjson.Field
		CookieName       respjson.Field
		CreatedByUserID  respjson.Field
		UpdatedAt        respjson.Field
		UpdatedByUserID  respjson.Field
		WhitelistDomains respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentSettingUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ExperimentSettingUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentSettingDeleteResponse struct {
	// Unique identifier for the account-level experiment settings record.
	ID string `json:"id" api:"required"`
	// Account that owns this experiment settings record.
	AccountID string `json:"accountId" api:"required"`
	CreatedAt string `json:"createdAt" api:"required"`
	// Human-readable name for this experimentation configuration.
	Name string `json:"name" api:"required"`
	// Pixel token used by the experiments runtime and CDN configuration. This is
	// informative for REST clients; use the settings `id` for createExperiment.
	Pixel string `json:"pixel" api:"required"`
	// Cookie name used to persist sticky variant assignments in the browser.
	CookieName      string `json:"cookieName" api:"nullable"`
	CreatedByUserID string `json:"createdByUserId" api:"nullable"`
	UpdatedAt       string `json:"updatedAt" api:"nullable"`
	UpdatedByUserID string `json:"updatedByUserId" api:"nullable"`
	// Limits which domains can load your experiments. When set, experiments using this
	// settings record are only served on these domains; the SDK refuses to load
	// anywhere else and your experiments never run on those hosts. Separate from
	// source `whitelistDomains`, which limits which domains can send events to the
	// CDP.
	WhitelistDomains []string `json:"whitelistDomains" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AccountID        respjson.Field
		CreatedAt        respjson.Field
		Name             respjson.Field
		Pixel            respjson.Field
		CookieName       respjson.Field
		CreatedByUserID  respjson.Field
		UpdatedAt        respjson.Field
		UpdatedByUserID  respjson.Field
		WhitelistDomains respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExperimentSettingDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ExperimentSettingDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentSettingNewParams struct {
	// Cookie name used to persist sticky variant assignments in the browser. Defaults
	// to `_cord_exp` when omitted on create.
	CookieName param.Opt[string] `json:"cookieName,omitzero"`
	// Human-readable name for this experimentation configuration. Defaults to
	// `Experiment Settings` when omitted on create.
	Name param.Opt[string] `json:"name,omitzero"`
	// Limits which domains can load your experiments. When set, experiments using this
	// settings record are only served on these domains; the SDK refuses to load
	// anywhere else and your experiments never run on those hosts. Separate from
	// source `whitelistDomains`, which limits which domains can send events to the
	// CDP.
	WhitelistDomains []string `json:"whitelistDomains,omitzero"`
	paramObj
}

func (r ExperimentSettingNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentSettingNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentSettingNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExperimentSettingUpdateParams struct {
	// Cookie name used to persist sticky variant assignments in the browser. Defaults
	// to `_cord_exp` when omitted on create.
	CookieName param.Opt[string] `json:"cookieName,omitzero"`
	// Human-readable name for this experimentation configuration. Defaults to
	// `Experiment Settings` when omitted on create.
	Name param.Opt[string] `json:"name,omitzero"`
	// Limits which domains can load your experiments. When set, experiments using this
	// settings record are only served on these domains; the SDK refuses to load
	// anywhere else and your experiments never run on those hosts. Separate from
	// source `whitelistDomains`, which limits which domains can send events to the
	// CDP.
	WhitelistDomains []string `json:"whitelistDomains,omitzero"`
	paramObj
}

func (r ExperimentSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ExperimentSettingUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExperimentSettingUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
