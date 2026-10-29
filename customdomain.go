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
// it is issued. The setup is `null` until created. Requires scope:
// `distributionTenant:find`. Requires API-key scope or current OAuth user
// permission: distributionTenant:find
func (r *CustomDomainService) Status(ctx context.Context, opts ...option.RequestOption) (res *CustomDomainStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "rest/v1/custom-domains/setup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Creates the account custom-domain setup and returns its `setupStatus` lifecycle
// guidance. Each account can have one setup; an existing setup returns HTTP 409.
// Requires scope: `distributionTenant:create`. Requires API-key scope or current
// OAuth user permission: distributionTenant:create
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
// configuring the required CNAME or after the certificate is issued. Requires
// scope: `distributionTenant:update`. Requires API-key scope or current OAuth user
// permission: distributionTenant:update
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
	// Managed certificate issuance status.
	CertificateStatus string                                  `json:"certificateStatus" api:"required"`
	CreatedAt         string                                  `json:"createdAt" api:"required"`
	DomainLimit       int64                                   `json:"domainLimit" api:"required"`
	Domains           []CustomDomainStatusResponseSetupDomain `json:"domains" api:"required"`
	Enabled           bool                                    `json:"enabled" api:"required"`
	// Current custom-domain setup progress and next action.
	SetupStatus CustomDomainStatusResponseSetupSetupStatus `json:"setupStatus" api:"required"`
	Status      string                                     `json:"status" api:"required"`
	UpdatedAt   string                                     `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CertificateActive respjson.Field
		CertificateStatus respjson.Field
		CreatedAt         respjson.Field
		DomainLimit       respjson.Field
		Domains           respjson.Field
		Enabled           respjson.Field
		SetupStatus       respjson.Field
		Status            respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomDomainStatusResponseSetup) RawJSON() string { return r.JSON.raw }
func (r *CustomDomainStatusResponseSetup) UnmarshalJSON(data []byte) error {
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
	BlockingDomains []string `json:"blockingDomains" api:"required"`
	IsReady         bool     `json:"isReady" api:"required"`
	Message         string   `json:"message" api:"required"`
	// Any of "ADD_DOMAINS", "CONFIGURE_DNS", "RECONCILE_CERTIFICATE", "WAIT",
	// "CONTACT_SUPPORT", "NONE".
	NextAction string `json:"nextAction" api:"required"`
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
	// Managed certificate issuance status.
	CertificateStatus string                               `json:"certificateStatus" api:"required"`
	CreatedAt         string                               `json:"createdAt" api:"required"`
	DomainLimit       int64                                `json:"domainLimit" api:"required"`
	Domains           []CustomDomainNewResponseSetupDomain `json:"domains" api:"required"`
	Enabled           bool                                 `json:"enabled" api:"required"`
	// Current custom-domain setup progress and next action.
	SetupStatus CustomDomainNewResponseSetupSetupStatus `json:"setupStatus" api:"required"`
	Status      string                                  `json:"status" api:"required"`
	UpdatedAt   string                                  `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CertificateActive respjson.Field
		CertificateStatus respjson.Field
		CreatedAt         respjson.Field
		DomainLimit       respjson.Field
		Domains           respjson.Field
		Enabled           respjson.Field
		SetupStatus       respjson.Field
		Status            respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomDomainNewResponseSetup) RawJSON() string { return r.JSON.raw }
func (r *CustomDomainNewResponseSetup) UnmarshalJSON(data []byte) error {
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
	BlockingDomains []string `json:"blockingDomains" api:"required"`
	IsReady         bool     `json:"isReady" api:"required"`
	Message         string   `json:"message" api:"required"`
	// Any of "ADD_DOMAINS", "CONFIGURE_DNS", "RECONCILE_CERTIFICATE", "WAIT",
	// "CONTACT_SUPPORT", "NONE".
	NextAction string `json:"nextAction" api:"required"`
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
	// Managed certificate issuance status.
	CertificateStatus string                                   `json:"certificateStatus" api:"required"`
	CreatedAt         string                                   `json:"createdAt" api:"required"`
	DomainLimit       int64                                    `json:"domainLimit" api:"required"`
	Domains           []CustomDomainReplaceResponseSetupDomain `json:"domains" api:"required"`
	Enabled           bool                                     `json:"enabled" api:"required"`
	// Current custom-domain setup progress and next action.
	SetupStatus CustomDomainReplaceResponseSetupSetupStatus `json:"setupStatus" api:"required"`
	Status      string                                      `json:"status" api:"required"`
	UpdatedAt   string                                      `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CertificateActive respjson.Field
		CertificateStatus respjson.Field
		CreatedAt         respjson.Field
		DomainLimit       respjson.Field
		Domains           respjson.Field
		Enabled           respjson.Field
		SetupStatus       respjson.Field
		Status            respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomDomainReplaceResponseSetup) RawJSON() string { return r.JSON.raw }
func (r *CustomDomainReplaceResponseSetup) UnmarshalJSON(data []byte) error {
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
	BlockingDomains []string `json:"blockingDomains" api:"required"`
	IsReady         bool     `json:"isReady" api:"required"`
	Message         string   `json:"message" api:"required"`
	// Any of "ADD_DOMAINS", "CONFIGURE_DNS", "RECONCILE_CERTIFICATE", "WAIT",
	// "CONTACT_SUPPORT", "NONE".
	NextAction string `json:"nextAction" api:"required"`
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
	// Managed certificate issuance status.
	CertificateStatus string                                                `json:"certificateStatus" api:"required"`
	CreatedAt         string                                                `json:"createdAt" api:"required"`
	DomainLimit       int64                                                 `json:"domainLimit" api:"required"`
	Domains           []CustomDomainReconcileCertificateResponseSetupDomain `json:"domains" api:"required"`
	Enabled           bool                                                  `json:"enabled" api:"required"`
	// Current custom-domain setup progress and next action.
	SetupStatus CustomDomainReconcileCertificateResponseSetupSetupStatus `json:"setupStatus" api:"required"`
	Status      string                                                   `json:"status" api:"required"`
	UpdatedAt   string                                                   `json:"updatedAt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CertificateActive respjson.Field
		CertificateStatus respjson.Field
		CreatedAt         respjson.Field
		DomainLimit       respjson.Field
		Domains           respjson.Field
		Enabled           respjson.Field
		SetupStatus       respjson.Field
		Status            respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomDomainReconcileCertificateResponseSetup) RawJSON() string { return r.JSON.raw }
func (r *CustomDomainReconcileCertificateResponseSetup) UnmarshalJSON(data []byte) error {
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
	BlockingDomains []string `json:"blockingDomains" api:"required"`
	IsReady         bool     `json:"isReady" api:"required"`
	Message         string   `json:"message" api:"required"`
	// Any of "ADD_DOMAINS", "CONFIGURE_DNS", "RECONCILE_CERTIFICATE", "WAIT",
	// "CONTACT_SUPPORT", "NONE".
	NextAction string `json:"nextAction" api:"required"`
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
