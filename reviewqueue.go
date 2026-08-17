// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bem

import (
	"github.com/bem-team/bem-go-sdk/option"
)

// ReviewQueueService contains methods and other services that help with
// interacting with the bem API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReviewQueueService] method instead.
type ReviewQueueService struct {
	options []option.RequestOption
}

// NewReviewQueueService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewReviewQueueService(opts ...option.RequestOption) (r ReviewQueueService) {
	r = ReviewQueueService{}
	r.options = opts
	return
}
