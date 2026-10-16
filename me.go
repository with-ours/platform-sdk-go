// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/with-ours/platform-sdk-go/v2/internal/apijson"
	"github.com/with-ours/platform-sdk-go/v2/internal/requestconfig"
	"github.com/with-ours/platform-sdk-go/v2/option"
	"github.com/with-ours/platform-sdk-go/v2/packages/respjson"
)

// MeService contains methods and other services that help with interacting with
// the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeService] method instead.
type MeService struct {
	Options []option.RequestOption
}

// NewMeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMeService(opts ...option.RequestOption) (r MeService) {
	r = MeService{}
	r.Options = opts
	return
}

// Identify the authenticated credential. Requires no API-key scope or OAuth user
// permission, so a narrow API key or a new OAuth connection can confirm it is
// valid before calling other operations.
func (r *MeService) Get(ctx context.Context, opts ...option.RequestOption) (res *MeGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type MeGetResponse struct {
	// ISO-8601 expiration of the API key or OAuth access token, or null.
	ExpiresAt time.Time `json:"expiresAt" api:"required" format:"date-time"`
	// Organization the credential acts for.
	OrganizationID string `json:"organizationId" api:"required"`
	// Display name of the organization, or null when unavailable.
	OrganizationName string `json:"organizationName" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt        respjson.Field
		OrganizationID   respjson.Field
		OrganizationName respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MeGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
