// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package raindrop

import (
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/option"
)

// GetPiiDataService contains methods and other services that help with interacting
// with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGetPiiDataService] method instead.
type GetPiiDataService struct {
	Options []option.RequestOption
}

// NewGetPiiDataService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGetPiiDataService(opts ...option.RequestOption) (r GetPiiDataService) {
	r = GetPiiDataService{}
	r.Options = opts
	return
}
