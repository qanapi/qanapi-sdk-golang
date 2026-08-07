// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package qanapi

import (
	"context"
	"net/http"
	"slices"

	"github.com/qanapi/qanapi-sdk-golang/internal/requestconfig"
	"github.com/qanapi/qanapi-sdk-golang/option"
)

// V3RoleService contains methods and other services that help with interacting
// with the qanapi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV3RoleService] method instead.
type V3RoleService struct {
	Options []option.RequestOption
}

// NewV3RoleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV3RoleService(opts ...option.RequestOption) (r V3RoleService) {
	r = V3RoleService{}
	r.Options = opts
	return
}

// List roles
func (r *V3RoleService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Role, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/roles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}
