// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package raindrop

import (
	"github.com/LiquidMetal-AI/lm-raindrop-go-sdk/option"
)

// UpdateMetadataService contains methods and other services that help with
// interacting with the raindrop API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUpdateMetadataService] method instead.
type UpdateMetadataService struct {
	Options []option.RequestOption
}

// NewUpdateMetadataService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUpdateMetadataService(opts ...option.RequestOption) (r UpdateMetadataService) {
	r = UpdateMetadataService{}
	r.Options = opts
	return
}
