// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package qanapi

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/qanapi/qanapi-sdk-golang/internal/apijson"
	"github.com/qanapi/qanapi-sdk-golang/internal/requestconfig"
	"github.com/qanapi/qanapi-sdk-golang/option"
	"github.com/qanapi/qanapi-sdk-golang/packages/param"
)

// V3UserService contains methods and other services that help with interacting
// with the qanapi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV3UserService] method instead.
type V3UserService struct {
	Options []option.RequestOption
}

// NewV3UserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV3UserService(opts ...option.RequestOption) (r V3UserService) {
	r = V3UserService{}
	r.Options = opts
	return
}

// Create user
func (r *V3UserService) New(ctx context.Context, body V3UserNewParams, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List users
func (r *V3UserService) List(ctx context.Context, opts ...option.RequestOption) (res *[]User, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete user
func (r *V3UserService) Delete(ctx context.Context, user int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("v3/users/%v", user)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get current user
func (r *V3UserService) Me(ctx context.Context, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/users/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update user
func (r *V3UserService) Patch(ctx context.Context, user int64, body V3UserPatchParams, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v3/users/%v", user)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Restore user
func (r *V3UserService) Restore(ctx context.Context, user int64, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v3/users/%v/restore", user)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return res, err
}

// Get user
func (r *V3UserService) Show(ctx context.Context, user int64, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v3/users/%v", user)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type V3UserNewParams struct {
	Email string `json:"email" api:"required" format:"email"`
	// Name of the role to assign
	Role string `json:"role" api:"required"`
	paramObj
}

func (r V3UserNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V3UserNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V3UserNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3UserPatchParams struct {
	Email            param.Opt[string] `json:"email,omitzero" format:"email"`
	Name             param.Opt[string] `json:"name,omitzero"`
	Role             param.Opt[string] `json:"role,omitzero"`
	TwoFactorEnabled param.Opt[bool]   `json:"two_factor_enabled,omitzero"`
	paramObj
}

func (r V3UserPatchParams) MarshalJSON() (data []byte, err error) {
	type shadow V3UserPatchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V3UserPatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
