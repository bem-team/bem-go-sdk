// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bem

import (
	"github.com/bem-team/bem-go-sdk/option"
)

// EntityTypeReviewerService contains methods and other services that help with
// interacting with the bem API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntityTypeReviewerService] method instead.
type EntityTypeReviewerService struct {
	options []option.RequestOption
}

// NewEntityTypeReviewerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEntityTypeReviewerService(opts ...option.RequestOption) (r EntityTypeReviewerService) {
	r = EntityTypeReviewerService{}
	r.options = opts
	return
}
