// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacyplatform

import (
	"github.com/stainless-sdks/ours-privacy-platform-go/option"
)

// RestV1Service contains methods and other services that help with interacting
// with the ours-privacy-platform API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRestV1Service] method instead.
type RestV1Service struct {
	Options               []option.RequestOption
	Destinations          RestV1DestinationService
	Sources               RestV1SourceService
	AllowedEvents         RestV1AllowedEventService
	ConsentSettings       RestV1ConsentSettingService
	GlobalDispatchCenters RestV1GlobalDispatchCenterService
	ReplaySettings        RestV1ReplaySettingService
	Versions              RestV1VersionService
}

// NewRestV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRestV1Service(opts ...option.RequestOption) (r RestV1Service) {
	r = RestV1Service{}
	r.Options = opts
	r.Destinations = NewRestV1DestinationService(opts...)
	r.Sources = NewRestV1SourceService(opts...)
	r.AllowedEvents = NewRestV1AllowedEventService(opts...)
	r.ConsentSettings = NewRestV1ConsentSettingService(opts...)
	r.GlobalDispatchCenters = NewRestV1GlobalDispatchCenterService(opts...)
	r.ReplaySettings = NewRestV1ReplaySettingService(opts...)
	r.Versions = NewRestV1VersionService(opts...)
	return
}
