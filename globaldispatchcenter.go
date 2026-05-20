// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy

import (
	"github.com/with-ours/platform-sdk-go/option"
)

// GlobalDispatchCenterService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGlobalDispatchCenterService] method instead.
type GlobalDispatchCenterService struct {
	Options []option.RequestOption
}

// NewGlobalDispatchCenterService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGlobalDispatchCenterService(opts ...option.RequestOption) (r GlobalDispatchCenterService) {
	r = GlobalDispatchCenterService{}
	r.Options = opts
	return
}
