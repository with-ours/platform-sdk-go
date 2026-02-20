// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacyplatform

import (
	"github.com/stainless-sdks/ours-privacy-platform-go/option"
)

// RestService contains methods and other services that help with interacting with
// the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRestService] method instead.
type RestService struct {
	Options []option.RequestOption
	V1      RestV1Service
}

// NewRestService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRestService(opts ...option.RequestOption) (r RestService) {
	r = RestService{}
	r.Options = opts
	r.V1 = NewRestV1Service(opts...)
	return
}
