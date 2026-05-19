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
	"github.com/with-ours/platform-sdk-go/packages/param"
	"github.com/with-ours/platform-sdk-go/packages/respjson"
)

// LocationService contains methods and other services that help with interacting
// with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLocationService] method instead.
type LocationService struct {
	Options []option.RequestOption
}

// NewLocationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLocationService(opts ...option.RequestOption) (r LocationService) {
	r = LocationService{}
	r.Options = opts
	return
}

// List every location for this account. Not paginated — each account has a small
// map-count limit (single digits in practice) so the response always fits in a
// single page. Requires scope: maps:list
func (r *LocationService) List(ctx context.Context, opts ...option.RequestOption) (res *LocationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Create a new location (map embed). All address fields are optional and can be
// filled in later via PATCH. Returns the slim entity with the server-assigned `id`
// so callers can immediately request `GET /rest/v1/locations/{id}/embed-code`.
// Requires scope: maps:create
func (r *LocationService) New(ctx context.Context, body LocationNewParams, opts ...option.RequestOption) (res *LocationNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Partially update a location. Only the fields you send are changed.
// `additionalAddresses` is replaced wholesale when sent — partial item updates are
// not merged. The map's computed center is recalculated on every PATCH from the
// latest coordinates. Requires scope: maps:update
func (r *LocationService) Update(ctx context.Context, id string, body LocationUpdateParams, opts ...option.RequestOption) (res *LocationUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/locations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Generate the paste-ready HTML embed snippet for a location. The response is a
// single self-contained HTML string (a `<style>` block + `<div>` wrapping an
// `<iframe>` pointed at the maps CDN for the current stage, plus an optional
// JSON-LD `<script>`). Customize the render with the optional query params
// (`color`, `theme`, `colorScheme`, `mapStyle`, `includeAddressBox`, `zoom`,
// `includeControls`, `includeSEOSchema`); all have sane defaults. Requires scope:
// maps:find
func (r *LocationService) EmbedCode(ctx context.Context, id string, query LocationEmbedCodeParams, opts ...option.RequestOption) (res *LocationEmbedCodeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/locations/%s/embed-code", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type LocationListResponse struct {
	Entities []LocationListResponseEntity `json:"entities" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LocationListResponse) RawJSON() string { return r.JSON.raw }
func (r *LocationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LocationListResponseEntity struct {
	ID                  string                                        `json:"id" api:"required"`
	AccountID           string                                        `json:"accountId" api:"required"`
	AdditionalAddresses []LocationListResponseEntityAdditionalAddress `json:"additionalAddresses" api:"nullable"`
	Center              any                                           `json:"center" api:"nullable"`
	City                string                                        `json:"city" api:"nullable"`
	Country             string                                        `json:"country" api:"nullable"`
	CreatedAt           string                                        `json:"createdAt" api:"nullable"`
	CustomDomain        string                                        `json:"customDomain" api:"nullable"`
	Latitude            float64                                       `json:"latitude" api:"nullable"`
	Line1               string                                        `json:"line1" api:"nullable"`
	Line2               string                                        `json:"line2" api:"nullable"`
	Longitude           float64                                       `json:"longitude" api:"nullable"`
	MapName             string                                        `json:"mapName" api:"nullable"`
	Name                string                                        `json:"name" api:"nullable"`
	PhoneNumber         string                                        `json:"phoneNumber" api:"nullable"`
	State               string                                        `json:"state" api:"nullable"`
	UpdatedAt           string                                        `json:"updatedAt" api:"nullable"`
	WebsiteLinkText     string                                        `json:"websiteLinkText" api:"nullable"`
	WebsiteURL          string                                        `json:"websiteUrl" api:"nullable"`
	Zip                 string                                        `json:"zip" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AccountID           respjson.Field
		AdditionalAddresses respjson.Field
		Center              respjson.Field
		City                respjson.Field
		Country             respjson.Field
		CreatedAt           respjson.Field
		CustomDomain        respjson.Field
		Latitude            respjson.Field
		Line1               respjson.Field
		Line2               respjson.Field
		Longitude           respjson.Field
		MapName             respjson.Field
		Name                respjson.Field
		PhoneNumber         respjson.Field
		State               respjson.Field
		UpdatedAt           respjson.Field
		WebsiteLinkText     respjson.Field
		WebsiteURL          respjson.Field
		Zip                 respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LocationListResponseEntity) RawJSON() string { return r.JSON.raw }
func (r *LocationListResponseEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LocationListResponseEntityAdditionalAddress struct {
	City            string  `json:"city" api:"nullable"`
	Country         string  `json:"country" api:"nullable"`
	Latitude        float64 `json:"latitude" api:"nullable"`
	Line1           string  `json:"line1" api:"nullable"`
	Line2           string  `json:"line2" api:"nullable"`
	Longitude       float64 `json:"longitude" api:"nullable"`
	Name            string  `json:"name" api:"nullable"`
	PhoneNumber     string  `json:"phoneNumber" api:"nullable"`
	State           string  `json:"state" api:"nullable"`
	WebsiteLinkText string  `json:"websiteLinkText" api:"nullable"`
	WebsiteURL      string  `json:"websiteUrl" api:"nullable"`
	Zip             string  `json:"zip" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City            respjson.Field
		Country         respjson.Field
		Latitude        respjson.Field
		Line1           respjson.Field
		Line2           respjson.Field
		Longitude       respjson.Field
		Name            respjson.Field
		PhoneNumber     respjson.Field
		State           respjson.Field
		WebsiteLinkText respjson.Field
		WebsiteURL      respjson.Field
		Zip             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LocationListResponseEntityAdditionalAddress) RawJSON() string { return r.JSON.raw }
func (r *LocationListResponseEntityAdditionalAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LocationNewResponse struct {
	ID                  string                                 `json:"id" api:"required"`
	AccountID           string                                 `json:"accountId" api:"required"`
	AdditionalAddresses []LocationNewResponseAdditionalAddress `json:"additionalAddresses" api:"nullable"`
	Center              any                                    `json:"center" api:"nullable"`
	City                string                                 `json:"city" api:"nullable"`
	Country             string                                 `json:"country" api:"nullable"`
	CreatedAt           string                                 `json:"createdAt" api:"nullable"`
	CustomDomain        string                                 `json:"customDomain" api:"nullable"`
	Latitude            float64                                `json:"latitude" api:"nullable"`
	Line1               string                                 `json:"line1" api:"nullable"`
	Line2               string                                 `json:"line2" api:"nullable"`
	Longitude           float64                                `json:"longitude" api:"nullable"`
	MapName             string                                 `json:"mapName" api:"nullable"`
	Name                string                                 `json:"name" api:"nullable"`
	PhoneNumber         string                                 `json:"phoneNumber" api:"nullable"`
	State               string                                 `json:"state" api:"nullable"`
	UpdatedAt           string                                 `json:"updatedAt" api:"nullable"`
	WebsiteLinkText     string                                 `json:"websiteLinkText" api:"nullable"`
	WebsiteURL          string                                 `json:"websiteUrl" api:"nullable"`
	Zip                 string                                 `json:"zip" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AccountID           respjson.Field
		AdditionalAddresses respjson.Field
		Center              respjson.Field
		City                respjson.Field
		Country             respjson.Field
		CreatedAt           respjson.Field
		CustomDomain        respjson.Field
		Latitude            respjson.Field
		Line1               respjson.Field
		Line2               respjson.Field
		Longitude           respjson.Field
		MapName             respjson.Field
		Name                respjson.Field
		PhoneNumber         respjson.Field
		State               respjson.Field
		UpdatedAt           respjson.Field
		WebsiteLinkText     respjson.Field
		WebsiteURL          respjson.Field
		Zip                 respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LocationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *LocationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LocationNewResponseAdditionalAddress struct {
	City            string  `json:"city" api:"nullable"`
	Country         string  `json:"country" api:"nullable"`
	Latitude        float64 `json:"latitude" api:"nullable"`
	Line1           string  `json:"line1" api:"nullable"`
	Line2           string  `json:"line2" api:"nullable"`
	Longitude       float64 `json:"longitude" api:"nullable"`
	Name            string  `json:"name" api:"nullable"`
	PhoneNumber     string  `json:"phoneNumber" api:"nullable"`
	State           string  `json:"state" api:"nullable"`
	WebsiteLinkText string  `json:"websiteLinkText" api:"nullable"`
	WebsiteURL      string  `json:"websiteUrl" api:"nullable"`
	Zip             string  `json:"zip" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City            respjson.Field
		Country         respjson.Field
		Latitude        respjson.Field
		Line1           respjson.Field
		Line2           respjson.Field
		Longitude       respjson.Field
		Name            respjson.Field
		PhoneNumber     respjson.Field
		State           respjson.Field
		WebsiteLinkText respjson.Field
		WebsiteURL      respjson.Field
		Zip             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LocationNewResponseAdditionalAddress) RawJSON() string { return r.JSON.raw }
func (r *LocationNewResponseAdditionalAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LocationUpdateResponse struct {
	ID                  string                                    `json:"id" api:"required"`
	AccountID           string                                    `json:"accountId" api:"required"`
	AdditionalAddresses []LocationUpdateResponseAdditionalAddress `json:"additionalAddresses" api:"nullable"`
	Center              any                                       `json:"center" api:"nullable"`
	City                string                                    `json:"city" api:"nullable"`
	Country             string                                    `json:"country" api:"nullable"`
	CreatedAt           string                                    `json:"createdAt" api:"nullable"`
	CustomDomain        string                                    `json:"customDomain" api:"nullable"`
	Latitude            float64                                   `json:"latitude" api:"nullable"`
	Line1               string                                    `json:"line1" api:"nullable"`
	Line2               string                                    `json:"line2" api:"nullable"`
	Longitude           float64                                   `json:"longitude" api:"nullable"`
	MapName             string                                    `json:"mapName" api:"nullable"`
	Name                string                                    `json:"name" api:"nullable"`
	PhoneNumber         string                                    `json:"phoneNumber" api:"nullable"`
	State               string                                    `json:"state" api:"nullable"`
	UpdatedAt           string                                    `json:"updatedAt" api:"nullable"`
	WebsiteLinkText     string                                    `json:"websiteLinkText" api:"nullable"`
	WebsiteURL          string                                    `json:"websiteUrl" api:"nullable"`
	Zip                 string                                    `json:"zip" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AccountID           respjson.Field
		AdditionalAddresses respjson.Field
		Center              respjson.Field
		City                respjson.Field
		Country             respjson.Field
		CreatedAt           respjson.Field
		CustomDomain        respjson.Field
		Latitude            respjson.Field
		Line1               respjson.Field
		Line2               respjson.Field
		Longitude           respjson.Field
		MapName             respjson.Field
		Name                respjson.Field
		PhoneNumber         respjson.Field
		State               respjson.Field
		UpdatedAt           respjson.Field
		WebsiteLinkText     respjson.Field
		WebsiteURL          respjson.Field
		Zip                 respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LocationUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *LocationUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LocationUpdateResponseAdditionalAddress struct {
	City            string  `json:"city" api:"nullable"`
	Country         string  `json:"country" api:"nullable"`
	Latitude        float64 `json:"latitude" api:"nullable"`
	Line1           string  `json:"line1" api:"nullable"`
	Line2           string  `json:"line2" api:"nullable"`
	Longitude       float64 `json:"longitude" api:"nullable"`
	Name            string  `json:"name" api:"nullable"`
	PhoneNumber     string  `json:"phoneNumber" api:"nullable"`
	State           string  `json:"state" api:"nullable"`
	WebsiteLinkText string  `json:"websiteLinkText" api:"nullable"`
	WebsiteURL      string  `json:"websiteUrl" api:"nullable"`
	Zip             string  `json:"zip" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City            respjson.Field
		Country         respjson.Field
		Latitude        respjson.Field
		Line1           respjson.Field
		Line2           respjson.Field
		Longitude       respjson.Field
		Name            respjson.Field
		PhoneNumber     respjson.Field
		State           respjson.Field
		WebsiteLinkText respjson.Field
		WebsiteURL      respjson.Field
		Zip             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LocationUpdateResponseAdditionalAddress) RawJSON() string { return r.JSON.raw }
func (r *LocationUpdateResponseAdditionalAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LocationEmbedCodeResponse struct {
	// Self-contained HTML snippet (a `<style>` + `<div>` wrapping an `<iframe>`, plus
	// an optional JSON-LD `<script>`) ready to paste into any page. The iframe `src`
	// points to the maps CDN for the current stage.
	EmbedCode string `json:"embedCode" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EmbedCode   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LocationEmbedCodeResponse) RawJSON() string { return r.JSON.raw }
func (r *LocationEmbedCodeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LocationNewParams struct {
	City                param.Opt[string]                    `json:"city,omitzero"`
	Country             param.Opt[string]                    `json:"country,omitzero"`
	CustomDomain        param.Opt[string]                    `json:"customDomain,omitzero"`
	Latitude            param.Opt[float64]                   `json:"latitude,omitzero"`
	Line1               param.Opt[string]                    `json:"line1,omitzero"`
	Line2               param.Opt[string]                    `json:"line2,omitzero"`
	Longitude           param.Opt[float64]                   `json:"longitude,omitzero"`
	MapName             param.Opt[string]                    `json:"mapName,omitzero"`
	Name                param.Opt[string]                    `json:"name,omitzero"`
	PhoneNumber         param.Opt[string]                    `json:"phoneNumber,omitzero"`
	State               param.Opt[string]                    `json:"state,omitzero"`
	WebsiteLinkText     param.Opt[string]                    `json:"websiteLinkText,omitzero"`
	WebsiteURL          param.Opt[string]                    `json:"websiteUrl,omitzero"`
	Zip                 param.Opt[string]                    `json:"zip,omitzero"`
	AdditionalAddresses []LocationNewParamsAdditionalAddress `json:"additionalAddresses,omitzero"`
	Center              any                                  `json:"center,omitzero"`
	paramObj
}

func (r LocationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LocationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LocationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Latitude, Longitude are required.
type LocationNewParamsAdditionalAddress struct {
	Latitude        float64           `json:"latitude" api:"required"`
	Longitude       float64           `json:"longitude" api:"required"`
	City            param.Opt[string] `json:"city,omitzero"`
	Country         param.Opt[string] `json:"country,omitzero"`
	Line1           param.Opt[string] `json:"line1,omitzero"`
	Line2           param.Opt[string] `json:"line2,omitzero"`
	Name            param.Opt[string] `json:"name,omitzero"`
	PhoneNumber     param.Opt[string] `json:"phoneNumber,omitzero"`
	State           param.Opt[string] `json:"state,omitzero"`
	WebsiteLinkText param.Opt[string] `json:"websiteLinkText,omitzero"`
	WebsiteURL      param.Opt[string] `json:"websiteUrl,omitzero"`
	Zip             param.Opt[string] `json:"zip,omitzero"`
	paramObj
}

func (r LocationNewParamsAdditionalAddress) MarshalJSON() (data []byte, err error) {
	type shadow LocationNewParamsAdditionalAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LocationNewParamsAdditionalAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LocationUpdateParams struct {
	City                param.Opt[string]                       `json:"city,omitzero"`
	Country             param.Opt[string]                       `json:"country,omitzero"`
	CustomDomain        param.Opt[string]                       `json:"customDomain,omitzero"`
	Latitude            param.Opt[float64]                      `json:"latitude,omitzero"`
	Line1               param.Opt[string]                       `json:"line1,omitzero"`
	Line2               param.Opt[string]                       `json:"line2,omitzero"`
	Longitude           param.Opt[float64]                      `json:"longitude,omitzero"`
	MapName             param.Opt[string]                       `json:"mapName,omitzero"`
	Name                param.Opt[string]                       `json:"name,omitzero"`
	PhoneNumber         param.Opt[string]                       `json:"phoneNumber,omitzero"`
	State               param.Opt[string]                       `json:"state,omitzero"`
	WebsiteLinkText     param.Opt[string]                       `json:"websiteLinkText,omitzero"`
	WebsiteURL          param.Opt[string]                       `json:"websiteUrl,omitzero"`
	Zip                 param.Opt[string]                       `json:"zip,omitzero"`
	AdditionalAddresses []LocationUpdateParamsAdditionalAddress `json:"additionalAddresses,omitzero"`
	Center              any                                     `json:"center,omitzero"`
	paramObj
}

func (r LocationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow LocationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LocationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Latitude, Longitude are required.
type LocationUpdateParamsAdditionalAddress struct {
	Latitude        float64           `json:"latitude" api:"required"`
	Longitude       float64           `json:"longitude" api:"required"`
	City            param.Opt[string] `json:"city,omitzero"`
	Country         param.Opt[string] `json:"country,omitzero"`
	Line1           param.Opt[string] `json:"line1,omitzero"`
	Line2           param.Opt[string] `json:"line2,omitzero"`
	Name            param.Opt[string] `json:"name,omitzero"`
	PhoneNumber     param.Opt[string] `json:"phoneNumber,omitzero"`
	State           param.Opt[string] `json:"state,omitzero"`
	WebsiteLinkText param.Opt[string] `json:"websiteLinkText,omitzero"`
	WebsiteURL      param.Opt[string] `json:"websiteUrl,omitzero"`
	Zip             param.Opt[string] `json:"zip,omitzero"`
	paramObj
}

func (r LocationUpdateParamsAdditionalAddress) MarshalJSON() (data []byte, err error) {
	type shadow LocationUpdateParamsAdditionalAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LocationUpdateParamsAdditionalAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LocationEmbedCodeParams struct {
	// Brand color used in the embedded map UI. Any CSS color string.
	Color param.Opt[string] `query:"color,omitzero" json:"-"`
	// Render the address sidebar overlay next to the map. Send `false` to hide.
	IncludeAddressBox param.Opt[bool] `query:"includeAddressBox,omitzero" json:"-"`
	// Emit a `schema.org` Place JSON-LD block alongside the iframe so search engines
	// can index the location.
	IncludeSeoSchema param.Opt[bool] `query:"includeSEOSchema,omitzero" json:"-"`
	// Initial map zoom level (Google-style 1–20).
	Zoom param.Opt[int64] `query:"zoom,omitzero" json:"-"`
	// Light or dark color scheme.
	//
	// Any of "light", "dark".
	ColorScheme LocationEmbedCodeParamsColorScheme `query:"colorScheme,omitzero" json:"-"`
	// Whether the embed renders map controls. `accessible` enables keyboard-navigable
	// controls.
	//
	// Any of "yes", "no", "accessible".
	IncludeControls LocationEmbedCodeParamsIncludeControls `query:"includeControls,omitzero" json:"-"`
	// Base map style.
	//
	// Any of "Standard", "Monochrome".
	MapStyle LocationEmbedCodeParamsMapStyle `query:"mapStyle,omitzero" json:"-"`
	// Visual theme variant.
	//
	// Any of "default", "modern".
	Theme LocationEmbedCodeParamsTheme `query:"theme,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LocationEmbedCodeParams]'s query parameters as
// `url.Values`.
func (r LocationEmbedCodeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Light or dark color scheme.
type LocationEmbedCodeParamsColorScheme string

const (
	LocationEmbedCodeParamsColorSchemeLight LocationEmbedCodeParamsColorScheme = "light"
	LocationEmbedCodeParamsColorSchemeDark  LocationEmbedCodeParamsColorScheme = "dark"
)

// Whether the embed renders map controls. `accessible` enables keyboard-navigable
// controls.
type LocationEmbedCodeParamsIncludeControls string

const (
	LocationEmbedCodeParamsIncludeControlsYes        LocationEmbedCodeParamsIncludeControls = "yes"
	LocationEmbedCodeParamsIncludeControlsNo         LocationEmbedCodeParamsIncludeControls = "no"
	LocationEmbedCodeParamsIncludeControlsAccessible LocationEmbedCodeParamsIncludeControls = "accessible"
)

// Base map style.
type LocationEmbedCodeParamsMapStyle string

const (
	LocationEmbedCodeParamsMapStyleStandard   LocationEmbedCodeParamsMapStyle = "Standard"
	LocationEmbedCodeParamsMapStyleMonochrome LocationEmbedCodeParamsMapStyle = "Monochrome"
)

// Visual theme variant.
type LocationEmbedCodeParamsTheme string

const (
	LocationEmbedCodeParamsThemeDefault LocationEmbedCodeParamsTheme = "default"
	LocationEmbedCodeParamsThemeModern  LocationEmbedCodeParamsTheme = "modern"
)
