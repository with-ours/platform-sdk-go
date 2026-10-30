// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/with-ours/platform-sdk-go/v2/internal/apijson"
	"github.com/with-ours/platform-sdk-go/v2/internal/requestconfig"
	"github.com/with-ours/platform-sdk-go/v2/option"
	"github.com/with-ours/platform-sdk-go/v2/packages/param"
	"github.com/with-ours/platform-sdk-go/v2/packages/respjson"
)

// CustomDomainService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomDomainService] method instead.
type CustomDomainService struct {
	Options []option.RequestOption
}

// NewCustomDomainService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomDomainService(opts ...option.RequestOption) (r CustomDomainService) {
	r = CustomDomainService{}
	r.Options = opts
	return
}

// Returns the account custom-domain setup, observed DNS values, and `setupStatus`
// lifecycle guidance, including readiness, blocking domains, and the next
// recommended action. Reading status also advances certificate setup: it starts
// certificate validation once DNS is configured and activates the certificate once
// it is issued. Terminal certificate failures include `certificateFailure` with a
// provider code, explanation, and failed validation domains; a successful status
// read still returns HTTP 200. DNS blockers can coexist with a certificate
// failure. Refresh does not restart a terminal failed request; contact support
// after remediation. The setup is `null` until created. Requires scope:
// `distributionTenant:find`. Requires API-key scope or current OAuth user
// permission: distributionTenant:find
func (r *CustomDomainService) Status(ctx context.Context, opts ...option.RequestOption) (res *CustomDomainStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/custom-domains/setup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Creates the account custom-domain setup and returns its `setupStatus` lifecycle
// guidance with `certificateFailure: null`. Each account can have one setup; an
// existing setup returns HTTP 409. Requires scope: `distributionTenant:create`.
// Requires API-key scope or current OAuth user permission:
// distributionTenant:create
func (r *CustomDomainService) New(ctx context.Context, opts ...option.RequestOption) (res *CustomDomainNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/custom-domains"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Replaces the full custom-domain list for this account and returns `setupStatus`
// lifecycle guidance. The setup `id` is the path identifier. The account allowance
// is returned as `domainLimit` in the setup response (10 by default, up to 99 with
// an increased allowance). An empty `domains` list removes all custom domains.
// Duplicate domains are rejected. Domains under `oursprivacy.com` are rejected.
// Domain delivery and certificate settings are managed by Ours Privacy; DNS
// changes are customer-managed. Requires scope: `distributionTenant:update`.
// Requires API-key scope or current OAuth user permission:
// distributionTenant:update
func (r *CustomDomainService) Replace(ctx context.Context, id string, body CustomDomainReplaceParams, opts ...option.RequestOption) (res *CustomDomainReplaceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/custom-domains/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Checks DNS and certificate status, applies any required certificate changes, and
// returns updated `setupStatus` lifecycle guidance. Optional: reading the setup
// status performs the same steps automatically. Use it to retry immediately after
// configuring the required CNAME or after the certificate is issued. It does not
// restart a terminal failed certificate request; contact support after correcting
// the reported problem. Requires scope: `distributionTenant:update`. Requires
// API-key scope or current OAuth user permission: distributionTenant:update
func (r *CustomDomainService) ReconcileCertificate(ctx context.Context, id string, opts ...option.RequestOption) (res *CustomDomainReconcileCertificateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rest/v1/custom-domains/%s/certificate-validation", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type CustomDomainStatusResponse struct {
	Setup CustomDomainStatusResponseSetup `json:"setup" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Setup       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomDomainStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomDomainStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomDomainStatusResponseSetup struct {
	// Custom-domain setup ID.
	ID string `json:"id" api:"required"`
	// Whether the managed certificate is attached to serve the custom domains.
	CertificateActive bool `json:"certificateActive" api:"required"`
	// Current terminal certificate failure. Null for new, pending, or issued requests;
	// independent of DNS lifecycle blockers.
	CertificateFailure CustomDomainStatusResponseSetupCertificateFailure `json:"certificateFailure" api:"required"`
	// Managed certificate issuance status.
	CertificateStatus string `json:"certificateStatus" api:"required"`
	// ISO 8601 timestamp when the setup was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Maximum number of custom domains this account can attach.
	DomainLimit int64 `json:"domainLimit" api:"required"`
	// Custom domains attached to the setup.
	Domains []CustomDomainStatusResponseSetupDomain `json:"domains" api:"required"`
	// Whether the setup is enabled to serve traffic.
	Enabled bool `json:"enabled" api:"required"`
	// Current custom-domain setup progress and next action.
	SetupStatus CustomDomainStatusResponseSetupSetupStatus `json:"setupStatus" api:"required"`
	// Delivery deployment status, or null when unavailable.
	Status string `json:"status" api:"required"`
	// ISO 8601 timestamp when the setup was last updated.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CertificateActive  respjson.Field
		CertificateFailure respjson.Field
		CertificateStatus  respjson.Field
		CreatedAt          respjson.Field
		DomainLimit        respjson.Field
		Domains            respjson.Field
		Enabled            respjson.Field
		SetupStatus        respjson.Field
		Status             respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomDomainStatusResponseSetup) RawJSON() string { return r.JSON.raw }
func (r *CustomDomainStatusResponseSetup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current terminal certificate failure. Null for new, pending, or issued requests;
// independent of DNS lifecycle blockers.
type CustomDomainStatusResponseSetupCertificateFailure struct {
	// Certificate failure reason, such as CAA_ERROR when DNS CAA records do not
	// authorize the issuer. Other values may appear. Null when unavailable.
	Code string `json:"code" api:"required"`
	// Customer domains with failed certificate validation, not individually proven CAA
	// policy defects. An empty list means none were identified.
	Domains []string `json:"domains" api:"required"`
	// Certificate failure explanation and next steps.
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Domains     respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomDomainStatusResponseSetupCertificateFailure) RawJSON() string { return r.JSON.raw }
func (r *CustomDomainStatusResponseSetupCertificateFailure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomDomainStatusResponseSetupDomain struct {
	// Required CNAME target for this domain.
	CnameValue string `json:"cnameValue" api:"required"`
	// Whether DNS points to the required CNAME target.
	//
	// Any of "configured", "notConfigured", "wrongEndpoint", "error".
	DNSStatus string `json:"dnsStatus" api:"required"`
	// DNS value currently observed for this domain.
	DNSValue string `json:"dnsValue" api:"required"`
	// Custom domain attached to the account.
	Domain string `json:"domain" api:"required"`
	// Domain attachment status.
	//
	// Any of "active", "inactive".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CnameValue  respjson.Field
		DNSStatus   respjson.Field
		DNSValue    respjson.Field
		Domain      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomDomainStatusResponseSetupDomain) RawJSON() string { return r.JSON.raw }
func (r *CustomDomainStatusResponseSetupDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current custom-domain setup progress and next action.
type CustomDomainStatusResponseSetupSetupStatus struct {
	// Custom domains whose DNS currently blocks setup. Empty when DNS is not the
	// blocker.
	BlockingDomains []string `json:"blockingDomains" api:"required"`
	// Whether every custom domain is ready to serve traffic.
	IsReady bool `json:"isReady" api:"required"`
	// Customer-facing explanation of the current stage.
	Message string `json:"message" api:"required"`
	// Recommended next step to move setup forward.
	//
	// Any of "ADD_DOMAINS", "CONFIGURE_DNS", "RECONCILE_CERTIFICATE", "WAIT",
	// "CONTACT_SUPPORT", "NONE".
	NextAction string `json:"nextAction" api:"required"`
	// Current setup stage. READY means every custom domain can serve traffic.
	//
	// Any of "NO_DOMAINS", "DNS_REQUIRED", "CERTIFICATE_VALIDATION_REQUIRED",
	// "CERTIFICATE_PENDING", "CERTIFICATE_ACTIVATION_REQUIRED", "DEPLOYING", "READY",
	// "CERTIFICATE_FAILED", "DISABLED".
	State string `json:"state" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BlockingDomains respjson.Field
		IsReady         respjson.Field
		Message         respjson.Field
		NextAction      respjson.Field
		State           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomDomainStatusResponseSetupSetupStatus) RawJSON() string { return r.JSON.raw }
func (r *CustomDomainStatusResponseSetupSetupStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomDomainNewResponse struct {
	Setup CustomDomainNewResponseSetup `json:"setup" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Setup       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomDomainNewResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomDomainNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomDomainNewResponseSetup struct {
	// Custom-domain setup ID.
	ID string `json:"id" api:"required"`
	// Whether the managed certificate is attached to serve the custom domains.
	CertificateActive bool `json:"certificateActive" api:"required"`
	// Current terminal certificate failure. Null for new, pending, or issued requests;
	// independent of DNS lifecycle blockers.
	CertificateFailure CustomDomainNewResponseSetupCertificateFailure `json:"certificateFailure" api:"required"`
	// Managed certificate issuance status.
	CertificateStatus string `json:"certificateStatus" api:"required"`
	// ISO 8601 timestamp when the setup was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Maximum number of custom domains this account can attach.
	DomainLimit int64 `json:"domainLimit" api:"required"`
	// Custom domains attached to the setup.
	Domains []CustomDomainNewResponseSetupDomain `json:"domains" api:"required"`
	// Whether the setup is enabled to serve traffic.
	Enabled bool `json:"enabled" api:"required"`
	// Current custom-domain setup progress and next action.
	SetupStatus CustomDomainNewResponseSetupSetupStatus `json:"setupStatus" api:"required"`
	// Delivery deployment status, or null when unavailable.
	Status string `json:"status" api:"required"`
	// ISO 8601 timestamp when the setup was last updated.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CertificateActive  respjson.Field
		CertificateFailure respjson.Field
		CertificateStatus  respjson.Field
		CreatedAt          respjson.Field
		DomainLimit        respjson.Field
		Domains            respjson.Field
		Enabled            respjson.Field
		SetupStatus        respjson.Field
		Status             respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomDomainNewResponseSetup) RawJSON() string { return r.JSON.raw }
func (r *CustomDomainNewResponseSetup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current terminal certificate failure. Null for new, pending, or issued requests;
// independent of DNS lifecycle blockers.
type CustomDomainNewResponseSetupCertificateFailure struct {
	// Certificate failure reason, such as CAA_ERROR when DNS CAA records do not
	// authorize the issuer. Other values may appear. Null when unavailable.
	Code string `json:"code" api:"required"`
	// Customer domains with failed certificate validation, not individually proven CAA
	// policy defects. An empty list means none were identified.
	Domains []string `json:"domains" api:"required"`
	// Certificate failure explanation and next steps.
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Domains     respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomDomainNewResponseSetupCertificateFailure) RawJSON() string { return r.JSON.raw }
func (r *CustomDomainNewResponseSetupCertificateFailure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomDomainNewResponseSetupDomain struct {
	// Required CNAME target for this domain.
	CnameValue string `json:"cnameValue" api:"required"`
	// Whether DNS points to the required CNAME target.
	//
	// Any of "configured", "notConfigured", "wrongEndpoint", "error".
	DNSStatus string `json:"dnsStatus" api:"required"`
	// DNS value currently observed for this domain.
	DNSValue string `json:"dnsValue" api:"required"`
	// Custom domain attached to the account.
	Domain string `json:"domain" api:"required"`
	// Domain attachment status.
	//
	// Any of "active", "inactive".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CnameValue  respjson.Field
		DNSStatus   respjson.Field
		DNSValue    respjson.Field
		Domain      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomDomainNewResponseSetupDomain) RawJSON() string { return r.JSON.raw }
func (r *CustomDomainNewResponseSetupDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current custom-domain setup progress and next action.
type CustomDomainNewResponseSetupSetupStatus struct {
	// Custom domains whose DNS currently blocks setup. Empty when DNS is not the
	// blocker.
	BlockingDomains []string `json:"blockingDomains" api:"required"`
	// Whether every custom domain is ready to serve traffic.
	IsReady bool `json:"isReady" api:"required"`
	// Customer-facing explanation of the current stage.
	Message string `json:"message" api:"required"`
	// Recommended next step to move setup forward.
	//
	// Any of "ADD_DOMAINS", "CONFIGURE_DNS", "RECONCILE_CERTIFICATE", "WAIT",
	// "CONTACT_SUPPORT", "NONE".
	NextAction string `json:"nextAction" api:"required"`
	// Current setup stage. READY means every custom domain can serve traffic.
	//
	// Any of "NO_DOMAINS", "DNS_REQUIRED", "CERTIFICATE_VALIDATION_REQUIRED",
	// "CERTIFICATE_PENDING", "CERTIFICATE_ACTIVATION_REQUIRED", "DEPLOYING", "READY",
	// "CERTIFICATE_FAILED", "DISABLED".
	State string `json:"state" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BlockingDomains respjson.Field
		IsReady         respjson.Field
		Message         respjson.Field
		NextAction      respjson.Field
		State           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomDomainNewResponseSetupSetupStatus) RawJSON() string { return r.JSON.raw }
func (r *CustomDomainNewResponseSetupSetupStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomDomainReplaceResponse struct {
	Setup CustomDomainReplaceResponseSetup `json:"setup" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Setup       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomDomainReplaceResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomDomainReplaceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomDomainReplaceResponseSetup struct {
	// Custom-domain setup ID.
	ID string `json:"id" api:"required"`
	// Whether the managed certificate is attached to serve the custom domains.
	CertificateActive bool `json:"certificateActive" api:"required"`
	// Current terminal certificate failure. Null for new, pending, or issued requests;
	// independent of DNS lifecycle blockers.
	CertificateFailure CustomDomainReplaceResponseSetupCertificateFailure `json:"certificateFailure" api:"required"`
	// Managed certificate issuance status.
	CertificateStatus string `json:"certificateStatus" api:"required"`
	// ISO 8601 timestamp when the setup was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Maximum number of custom domains this account can attach.
	DomainLimit int64 `json:"domainLimit" api:"required"`
	// Custom domains attached to the setup.
	Domains []CustomDomainReplaceResponseSetupDomain `json:"domains" api:"required"`
	// Whether the setup is enabled to serve traffic.
	Enabled bool `json:"enabled" api:"required"`
	// Current custom-domain setup progress and next action.
	SetupStatus CustomDomainReplaceResponseSetupSetupStatus `json:"setupStatus" api:"required"`
	// Delivery deployment status, or null when unavailable.
	Status string `json:"status" api:"required"`
	// ISO 8601 timestamp when the setup was last updated.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CertificateActive  respjson.Field
		CertificateFailure respjson.Field
		CertificateStatus  respjson.Field
		CreatedAt          respjson.Field
		DomainLimit        respjson.Field
		Domains            respjson.Field
		Enabled            respjson.Field
		SetupStatus        respjson.Field
		Status             respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomDomainReplaceResponseSetup) RawJSON() string { return r.JSON.raw }
func (r *CustomDomainReplaceResponseSetup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current terminal certificate failure. Null for new, pending, or issued requests;
// independent of DNS lifecycle blockers.
type CustomDomainReplaceResponseSetupCertificateFailure struct {
	// Certificate failure reason, such as CAA_ERROR when DNS CAA records do not
	// authorize the issuer. Other values may appear. Null when unavailable.
	Code string `json:"code" api:"required"`
	// Customer domains with failed certificate validation, not individually proven CAA
	// policy defects. An empty list means none were identified.
	Domains []string `json:"domains" api:"required"`
	// Certificate failure explanation and next steps.
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Domains     respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomDomainReplaceResponseSetupCertificateFailure) RawJSON() string { return r.JSON.raw }
func (r *CustomDomainReplaceResponseSetupCertificateFailure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomDomainReplaceResponseSetupDomain struct {
	// Required CNAME target for this domain.
	CnameValue string `json:"cnameValue" api:"required"`
	// Whether DNS points to the required CNAME target.
	//
	// Any of "configured", "notConfigured", "wrongEndpoint", "error".
	DNSStatus string `json:"dnsStatus" api:"required"`
	// DNS value currently observed for this domain.
	DNSValue string `json:"dnsValue" api:"required"`
	// Custom domain attached to the account.
	Domain string `json:"domain" api:"required"`
	// Domain attachment status.
	//
	// Any of "active", "inactive".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CnameValue  respjson.Field
		DNSStatus   respjson.Field
		DNSValue    respjson.Field
		Domain      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomDomainReplaceResponseSetupDomain) RawJSON() string { return r.JSON.raw }
func (r *CustomDomainReplaceResponseSetupDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current custom-domain setup progress and next action.
type CustomDomainReplaceResponseSetupSetupStatus struct {
	// Custom domains whose DNS currently blocks setup. Empty when DNS is not the
	// blocker.
	BlockingDomains []string `json:"blockingDomains" api:"required"`
	// Whether every custom domain is ready to serve traffic.
	IsReady bool `json:"isReady" api:"required"`
	// Customer-facing explanation of the current stage.
	Message string `json:"message" api:"required"`
	// Recommended next step to move setup forward.
	//
	// Any of "ADD_DOMAINS", "CONFIGURE_DNS", "RECONCILE_CERTIFICATE", "WAIT",
	// "CONTACT_SUPPORT", "NONE".
	NextAction string `json:"nextAction" api:"required"`
	// Current setup stage. READY means every custom domain can serve traffic.
	//
	// Any of "NO_DOMAINS", "DNS_REQUIRED", "CERTIFICATE_VALIDATION_REQUIRED",
	// "CERTIFICATE_PENDING", "CERTIFICATE_ACTIVATION_REQUIRED", "DEPLOYING", "READY",
	// "CERTIFICATE_FAILED", "DISABLED".
	State string `json:"state" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BlockingDomains respjson.Field
		IsReady         respjson.Field
		Message         respjson.Field
		NextAction      respjson.Field
		State           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomDomainReplaceResponseSetupSetupStatus) RawJSON() string { return r.JSON.raw }
func (r *CustomDomainReplaceResponseSetupSetupStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomDomainReconcileCertificateResponse struct {
	Setup CustomDomainReconcileCertificateResponseSetup `json:"setup" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Setup       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomDomainReconcileCertificateResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomDomainReconcileCertificateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomDomainReconcileCertificateResponseSetup struct {
	// Custom-domain setup ID.
	ID string `json:"id" api:"required"`
	// Whether the managed certificate is attached to serve the custom domains.
	CertificateActive bool `json:"certificateActive" api:"required"`
	// Current terminal certificate failure. Null for new, pending, or issued requests;
	// independent of DNS lifecycle blockers.
	CertificateFailure CustomDomainReconcileCertificateResponseSetupCertificateFailure `json:"certificateFailure" api:"required"`
	// Managed certificate issuance status.
	CertificateStatus string `json:"certificateStatus" api:"required"`
	// ISO 8601 timestamp when the setup was created.
	CreatedAt string `json:"createdAt" api:"required"`
	// Maximum number of custom domains this account can attach.
	DomainLimit int64 `json:"domainLimit" api:"required"`
	// Custom domains attached to the setup.
	Domains []CustomDomainReconcileCertificateResponseSetupDomain `json:"domains" api:"required"`
	// Whether the setup is enabled to serve traffic.
	Enabled bool `json:"enabled" api:"required"`
	// Current custom-domain setup progress and next action.
	SetupStatus CustomDomainReconcileCertificateResponseSetupSetupStatus `json:"setupStatus" api:"required"`
	// Delivery deployment status, or null when unavailable.
	Status string `json:"status" api:"required"`
	// ISO 8601 timestamp when the setup was last updated.
	UpdatedAt string `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CertificateActive  respjson.Field
		CertificateFailure respjson.Field
		CertificateStatus  respjson.Field
		CreatedAt          respjson.Field
		DomainLimit        respjson.Field
		Domains            respjson.Field
		Enabled            respjson.Field
		SetupStatus        respjson.Field
		Status             respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomDomainReconcileCertificateResponseSetup) RawJSON() string { return r.JSON.raw }
func (r *CustomDomainReconcileCertificateResponseSetup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current terminal certificate failure. Null for new, pending, or issued requests;
// independent of DNS lifecycle blockers.
type CustomDomainReconcileCertificateResponseSetupCertificateFailure struct {
	// Certificate failure reason, such as CAA_ERROR when DNS CAA records do not
	// authorize the issuer. Other values may appear. Null when unavailable.
	Code string `json:"code" api:"required"`
	// Customer domains with failed certificate validation, not individually proven CAA
	// policy defects. An empty list means none were identified.
	Domains []string `json:"domains" api:"required"`
	// Certificate failure explanation and next steps.
	Message string `json:"message" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Domains     respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomDomainReconcileCertificateResponseSetupCertificateFailure) RawJSON() string {
	return r.JSON.raw
}
func (r *CustomDomainReconcileCertificateResponseSetupCertificateFailure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomDomainReconcileCertificateResponseSetupDomain struct {
	// Required CNAME target for this domain.
	CnameValue string `json:"cnameValue" api:"required"`
	// Whether DNS points to the required CNAME target.
	//
	// Any of "configured", "notConfigured", "wrongEndpoint", "error".
	DNSStatus string `json:"dnsStatus" api:"required"`
	// DNS value currently observed for this domain.
	DNSValue string `json:"dnsValue" api:"required"`
	// Custom domain attached to the account.
	Domain string `json:"domain" api:"required"`
	// Domain attachment status.
	//
	// Any of "active", "inactive".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CnameValue  respjson.Field
		DNSStatus   respjson.Field
		DNSValue    respjson.Field
		Domain      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomDomainReconcileCertificateResponseSetupDomain) RawJSON() string { return r.JSON.raw }
func (r *CustomDomainReconcileCertificateResponseSetupDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current custom-domain setup progress and next action.
type CustomDomainReconcileCertificateResponseSetupSetupStatus struct {
	// Custom domains whose DNS currently blocks setup. Empty when DNS is not the
	// blocker.
	BlockingDomains []string `json:"blockingDomains" api:"required"`
	// Whether every custom domain is ready to serve traffic.
	IsReady bool `json:"isReady" api:"required"`
	// Customer-facing explanation of the current stage.
	Message string `json:"message" api:"required"`
	// Recommended next step to move setup forward.
	//
	// Any of "ADD_DOMAINS", "CONFIGURE_DNS", "RECONCILE_CERTIFICATE", "WAIT",
	// "CONTACT_SUPPORT", "NONE".
	NextAction string `json:"nextAction" api:"required"`
	// Current setup stage. READY means every custom domain can serve traffic.
	//
	// Any of "NO_DOMAINS", "DNS_REQUIRED", "CERTIFICATE_VALIDATION_REQUIRED",
	// "CERTIFICATE_PENDING", "CERTIFICATE_ACTIVATION_REQUIRED", "DEPLOYING", "READY",
	// "CERTIFICATE_FAILED", "DISABLED".
	State string `json:"state" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BlockingDomains respjson.Field
		IsReady         respjson.Field
		Message         respjson.Field
		NextAction      respjson.Field
		State           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomDomainReconcileCertificateResponseSetupSetupStatus) RawJSON() string { return r.JSON.raw }
func (r *CustomDomainReconcileCertificateResponseSetupSetupStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomDomainReplaceParams struct {
	// Replacement list of up to 99 custom domains, subject to the account's
	// domainLimit (10 by default). An empty list removes all custom domains. Domains
	// under oursprivacy.com are rejected.
	Domains []string `json:"domains,omitzero" api:"required"`
	paramObj
}

func (r CustomDomainReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomDomainReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomDomainReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
