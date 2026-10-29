// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/with-ours/platform-sdk-go/v2/internal/apijson"
	"github.com/with-ours/platform-sdk-go/v2/internal/apiquery"
	"github.com/with-ours/platform-sdk-go/v2/internal/requestconfig"
	"github.com/with-ours/platform-sdk-go/v2/option"
	"github.com/with-ours/platform-sdk-go/v2/packages/param"
	"github.com/with-ours/platform-sdk-go/v2/packages/respjson"
)

// ReportBrandingService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReportBrandingService] method instead.
type ReportBrandingService struct {
	Options []option.RequestOption
}

// NewReportBrandingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewReportBrandingService(opts ...option.RequestOption) (r ReportBrandingService) {
	r = ReportBrandingService{}
	r.Options = opts
	return
}

// Get the organization logo shown on Dashboards and Dashboard PDFs. Returns {
// logo: null } when no logo is set. Pass includeDataUrl=false to omit the image
// data and read only its format, dimensions, size, and update time. Requires
// Analytics Suite access. Requires API-key scope or current OAuth user permission:
// web-analytics:view
func (r *ReportBrandingService) Get(ctx context.Context, query ReportBrandingGetParams, opts ...option.RequestOption) (res *ReportBrandingGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/report-branding/logo"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Replace the organization logo used on Dashboards and Dashboard PDFs. Send
// contentType and base64 data for a PNG, JPEG, or WebP image of at most 524288
// bytes and 4096 pixels in each dimension. The bytes must match contentType; SVG
// is not accepted, so rasterize vector logos first. Requires Analytics Suite
// access and organization:manage-report-branding. Requires API-key scope or
// current OAuth user permission: organization:manage-report-branding
func (r *ReportBrandingService) Replace(ctx context.Context, body ReportBrandingReplaceParams, opts ...option.RequestOption) (res *ReportBrandingReplaceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/report-branding/logo"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Remove the organization logo from Dashboards and Dashboard PDFs. Succeeds when
// no logo is set. Requires Analytics Suite access and
// organization:manage-report-branding. Requires API-key scope or current OAuth
// user permission: organization:manage-report-branding
func (r *ReportBrandingService) Delete(ctx context.Context, opts ...option.RequestOption) (res *ReportBrandingDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/report-branding/logo"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type ReportBrandingGetResponse struct {
	// Organization logo shown on Dashboards and Dashboard PDFs, or null when none is
	// set.
	Logo ReportBrandingGetResponseLogo `json:"logo" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Logo        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportBrandingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ReportBrandingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Organization logo shown on Dashboards and Dashboard PDFs, or null when none is
// set.
type ReportBrandingGetResponseLogo struct {
	// Decoded image size in bytes.
	ByteSize int64 `json:"byteSize" api:"required"`
	// Any of "image/png", "image/jpeg", "image/webp".
	ContentType string `json:"contentType" api:"required"`
	// Image height in pixels.
	Height          int64     `json:"height" api:"required"`
	UpdatedAt       time.Time `json:"updatedAt" api:"required" format:"date-time"`
	UpdatedByUserID string    `json:"updatedByUserId" api:"required"`
	// Image width in pixels.
	Width int64 `json:"width" api:"required"`
	// data: URL containing the base64 image. Returned by GET unless includeDataUrl is
	// false; omitted from write responses.
	DataURL string `json:"dataUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByteSize        respjson.Field
		ContentType     respjson.Field
		Height          respjson.Field
		UpdatedAt       respjson.Field
		UpdatedByUserID respjson.Field
		Width           respjson.Field
		DataURL         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportBrandingGetResponseLogo) RawJSON() string { return r.JSON.raw }
func (r *ReportBrandingGetResponseLogo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportBrandingReplaceResponse struct {
	// Organization logo shown on Dashboards and Dashboard PDFs, or null when none is
	// set.
	Logo ReportBrandingReplaceResponseLogo `json:"logo" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Logo        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportBrandingReplaceResponse) RawJSON() string { return r.JSON.raw }
func (r *ReportBrandingReplaceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Organization logo shown on Dashboards and Dashboard PDFs, or null when none is
// set.
type ReportBrandingReplaceResponseLogo struct {
	// Decoded image size in bytes.
	ByteSize int64 `json:"byteSize" api:"required"`
	// Any of "image/png", "image/jpeg", "image/webp".
	ContentType string `json:"contentType" api:"required"`
	// Image height in pixels.
	Height          int64     `json:"height" api:"required"`
	UpdatedAt       time.Time `json:"updatedAt" api:"required" format:"date-time"`
	UpdatedByUserID string    `json:"updatedByUserId" api:"required"`
	// Image width in pixels.
	Width int64 `json:"width" api:"required"`
	// data: URL containing the base64 image. Returned by GET unless includeDataUrl is
	// false; omitted from write responses.
	DataURL string `json:"dataUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByteSize        respjson.Field
		ContentType     respjson.Field
		Height          respjson.Field
		UpdatedAt       respjson.Field
		UpdatedByUserID respjson.Field
		Width           respjson.Field
		DataURL         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportBrandingReplaceResponseLogo) RawJSON() string { return r.JSON.raw }
func (r *ReportBrandingReplaceResponseLogo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportBrandingDeleteResponse struct {
	// Organization logo shown on Dashboards and Dashboard PDFs, or null when none is
	// set.
	Logo ReportBrandingDeleteResponseLogo `json:"logo" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Logo        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportBrandingDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ReportBrandingDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Organization logo shown on Dashboards and Dashboard PDFs, or null when none is
// set.
type ReportBrandingDeleteResponseLogo struct {
	// Decoded image size in bytes.
	ByteSize int64 `json:"byteSize" api:"required"`
	// Any of "image/png", "image/jpeg", "image/webp".
	ContentType string `json:"contentType" api:"required"`
	// Image height in pixels.
	Height          int64     `json:"height" api:"required"`
	UpdatedAt       time.Time `json:"updatedAt" api:"required" format:"date-time"`
	UpdatedByUserID string    `json:"updatedByUserId" api:"required"`
	// Image width in pixels.
	Width int64 `json:"width" api:"required"`
	// data: URL containing the base64 image. Returned by GET unless includeDataUrl is
	// false; omitted from write responses.
	DataURL string `json:"dataUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByteSize        respjson.Field
		ContentType     respjson.Field
		Height          respjson.Field
		UpdatedAt       respjson.Field
		UpdatedByUserID respjson.Field
		Width           respjson.Field
		DataURL         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportBrandingDeleteResponseLogo) RawJSON() string { return r.JSON.raw }
func (r *ReportBrandingDeleteResponseLogo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportBrandingGetParams struct {
	// Whether to return the image as logo.dataUrl. Defaults to true; pass false to
	// read only the logo metadata.
	IncludeDataURL param.Opt[bool] `query:"includeDataUrl,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ReportBrandingGetParams]'s query parameters as
// `url.Values`.
func (r ReportBrandingGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ReportBrandingReplaceParams struct {
	// Image format: image/png, image/jpeg, or image/webp. SVG is not accepted;
	// rasterize vector logos first. Must match the uploaded bytes.
	//
	// Any of "image/png", "image/jpeg", "image/webp".
	ContentType ReportBrandingReplaceParamsContentType `json:"contentType,omitzero" api:"required"`
	// Standard base64 image bytes without a data: prefix. At most 524288 decoded bytes
	// and 4096 pixels in each dimension.
	Data string `json:"data" api:"required"`
	paramObj
}

func (r ReportBrandingReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow ReportBrandingReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReportBrandingReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image format: image/png, image/jpeg, or image/webp. SVG is not accepted;
// rasterize vector logos first. Must match the uploaded bytes.
type ReportBrandingReplaceParamsContentType string

const (
	ReportBrandingReplaceParamsContentTypeImagePng  ReportBrandingReplaceParamsContentType = "image/png"
	ReportBrandingReplaceParamsContentTypeImageJpeg ReportBrandingReplaceParamsContentType = "image/jpeg"
	ReportBrandingReplaceParamsContentTypeImageWebp ReportBrandingReplaceParamsContentType = "image/webp"
)
