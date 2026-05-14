// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy

import (
	"github.com/with-ours/platform-sdk-go/option"
)

// WebScannerRuleService contains methods and other services that help with
// interacting with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebScannerRuleService] method instead.
type WebScannerRuleService struct {
	Options []option.RequestOption
}

// NewWebScannerRuleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebScannerRuleService(opts ...option.RequestOption) (r WebScannerRuleService) {
	r = WebScannerRuleService{}
	r.Options = opts
	return
}
