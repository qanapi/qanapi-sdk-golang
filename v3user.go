// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package qanapi

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/qanapi/qanapi-sdk-golang/internal/apijson"
	"github.com/qanapi/qanapi-sdk-golang/internal/requestconfig"
	"github.com/qanapi/qanapi-sdk-golang/option"
	"github.com/qanapi/qanapi-sdk-golang/packages/param"
	"github.com/qanapi/qanapi-sdk-golang/packages/respjson"
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
func (r *V3UserService) New(ctx context.Context, body V3UserNewParams, opts ...option.RequestOption) (res *V3UserNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List users
func (r *V3UserService) List(ctx context.Context, opts ...option.RequestOption) (res *[]V3UserListResponse, err error) {
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
func (r *V3UserService) Me(ctx context.Context, opts ...option.RequestOption) (res *V3UserMeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/users/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update user
func (r *V3UserService) Patch(ctx context.Context, user int64, body V3UserPatchParams, opts ...option.RequestOption) (res *V3UserPatchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v3/users/%v", user)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Update user
func (r *V3UserService) Restore(ctx context.Context, user int64, body V3UserRestoreParams, opts ...option.RequestOption) (res *V3UserRestoreResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v3/users/%v", user)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Get user
func (r *V3UserService) Show(ctx context.Context, user int64, opts ...option.RequestOption) (res *V3UserShowResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v3/users/%v", user)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type V3UserNewResponse struct {
	ID               int64                   `json:"id" api:"required"`
	Email            string                  `json:"email" api:"required" format:"email"`
	Name             string                  `json:"name" api:"required"`
	CreatedAt        time.Time               `json:"created_at" format:"date-time"`
	Roles            []V3UserNewResponseRole `json:"roles"`
	TwoFactorEnabled bool                    `json:"two_factor_enabled"`
	UpdatedAt        time.Time               `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Email            respjson.Field
		Name             respjson.Field
		CreatedAt        respjson.Field
		Roles            respjson.Field
		TwoFactorEnabled respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3UserNewResponse) RawJSON() string { return r.JSON.raw }
func (r *V3UserNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3UserNewResponseRole struct {
	Name        string                            `json:"name" api:"required"`
	Description string                            `json:"description" api:"nullable"`
	Permissions []V3UserNewResponseRolePermission `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Description respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3UserNewResponseRole) RawJSON() string { return r.JSON.raw }
func (r *V3UserNewResponseRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3UserNewResponseRolePermission struct {
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3UserNewResponseRolePermission) RawJSON() string { return r.JSON.raw }
func (r *V3UserNewResponseRolePermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3UserListResponse struct {
	ID               int64                    `json:"id" api:"required"`
	Email            string                   `json:"email" api:"required" format:"email"`
	Name             string                   `json:"name" api:"required"`
	CreatedAt        time.Time                `json:"created_at" format:"date-time"`
	Roles            []V3UserListResponseRole `json:"roles"`
	TwoFactorEnabled bool                     `json:"two_factor_enabled"`
	UpdatedAt        time.Time                `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Email            respjson.Field
		Name             respjson.Field
		CreatedAt        respjson.Field
		Roles            respjson.Field
		TwoFactorEnabled respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3UserListResponse) RawJSON() string { return r.JSON.raw }
func (r *V3UserListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3UserListResponseRole struct {
	Name        string                             `json:"name" api:"required"`
	Description string                             `json:"description" api:"nullable"`
	Permissions []V3UserListResponseRolePermission `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Description respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3UserListResponseRole) RawJSON() string { return r.JSON.raw }
func (r *V3UserListResponseRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3UserListResponseRolePermission struct {
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3UserListResponseRolePermission) RawJSON() string { return r.JSON.raw }
func (r *V3UserListResponseRolePermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3UserMeResponse struct {
	ID               int64                  `json:"id" api:"required"`
	Email            string                 `json:"email" api:"required" format:"email"`
	Name             string                 `json:"name" api:"required"`
	CreatedAt        time.Time              `json:"created_at" format:"date-time"`
	Roles            []V3UserMeResponseRole `json:"roles"`
	TwoFactorEnabled bool                   `json:"two_factor_enabled"`
	UpdatedAt        time.Time              `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Email            respjson.Field
		Name             respjson.Field
		CreatedAt        respjson.Field
		Roles            respjson.Field
		TwoFactorEnabled respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3UserMeResponse) RawJSON() string { return r.JSON.raw }
func (r *V3UserMeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3UserMeResponseRole struct {
	Name        string                           `json:"name" api:"required"`
	Description string                           `json:"description" api:"nullable"`
	Permissions []V3UserMeResponseRolePermission `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Description respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3UserMeResponseRole) RawJSON() string { return r.JSON.raw }
func (r *V3UserMeResponseRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3UserMeResponseRolePermission struct {
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3UserMeResponseRolePermission) RawJSON() string { return r.JSON.raw }
func (r *V3UserMeResponseRolePermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3UserPatchResponse struct {
	ID               int64                     `json:"id" api:"required"`
	Email            string                    `json:"email" api:"required" format:"email"`
	Name             string                    `json:"name" api:"required"`
	CreatedAt        time.Time                 `json:"created_at" format:"date-time"`
	Roles            []V3UserPatchResponseRole `json:"roles"`
	TwoFactorEnabled bool                      `json:"two_factor_enabled"`
	UpdatedAt        time.Time                 `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Email            respjson.Field
		Name             respjson.Field
		CreatedAt        respjson.Field
		Roles            respjson.Field
		TwoFactorEnabled respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3UserPatchResponse) RawJSON() string { return r.JSON.raw }
func (r *V3UserPatchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3UserPatchResponseRole struct {
	Name        string                              `json:"name" api:"required"`
	Description string                              `json:"description" api:"nullable"`
	Permissions []V3UserPatchResponseRolePermission `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Description respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3UserPatchResponseRole) RawJSON() string { return r.JSON.raw }
func (r *V3UserPatchResponseRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3UserPatchResponseRolePermission struct {
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3UserPatchResponseRolePermission) RawJSON() string { return r.JSON.raw }
func (r *V3UserPatchResponseRolePermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3UserRestoreResponse struct {
	ID               int64                       `json:"id" api:"required"`
	Email            string                      `json:"email" api:"required" format:"email"`
	Name             string                      `json:"name" api:"required"`
	CreatedAt        time.Time                   `json:"created_at" format:"date-time"`
	Roles            []V3UserRestoreResponseRole `json:"roles"`
	TwoFactorEnabled bool                        `json:"two_factor_enabled"`
	UpdatedAt        time.Time                   `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Email            respjson.Field
		Name             respjson.Field
		CreatedAt        respjson.Field
		Roles            respjson.Field
		TwoFactorEnabled respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3UserRestoreResponse) RawJSON() string { return r.JSON.raw }
func (r *V3UserRestoreResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3UserRestoreResponseRole struct {
	Name        string                                `json:"name" api:"required"`
	Description string                                `json:"description" api:"nullable"`
	Permissions []V3UserRestoreResponseRolePermission `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Description respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3UserRestoreResponseRole) RawJSON() string { return r.JSON.raw }
func (r *V3UserRestoreResponseRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3UserRestoreResponseRolePermission struct {
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3UserRestoreResponseRolePermission) RawJSON() string { return r.JSON.raw }
func (r *V3UserRestoreResponseRolePermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3UserShowResponse struct {
	ID               int64                    `json:"id" api:"required"`
	Email            string                   `json:"email" api:"required" format:"email"`
	Name             string                   `json:"name" api:"required"`
	CreatedAt        time.Time                `json:"created_at" format:"date-time"`
	Roles            []V3UserShowResponseRole `json:"roles"`
	TwoFactorEnabled bool                     `json:"two_factor_enabled"`
	UpdatedAt        time.Time                `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Email            respjson.Field
		Name             respjson.Field
		CreatedAt        respjson.Field
		Roles            respjson.Field
		TwoFactorEnabled respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3UserShowResponse) RawJSON() string { return r.JSON.raw }
func (r *V3UserShowResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3UserShowResponseRole struct {
	Name        string                             `json:"name" api:"required"`
	Description string                             `json:"description" api:"nullable"`
	Permissions []V3UserShowResponseRolePermission `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Description respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3UserShowResponseRole) RawJSON() string { return r.JSON.raw }
func (r *V3UserShowResponseRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3UserShowResponseRolePermission struct {
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3UserShowResponseRolePermission) RawJSON() string { return r.JSON.raw }
func (r *V3UserShowResponseRolePermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

type V3UserRestoreParams struct {
	Email            param.Opt[string] `json:"email,omitzero" format:"email"`
	Name             param.Opt[string] `json:"name,omitzero"`
	Role             param.Opt[string] `json:"role,omitzero"`
	TwoFactorEnabled param.Opt[bool]   `json:"two_factor_enabled,omitzero"`
	paramObj
}

func (r V3UserRestoreParams) MarshalJSON() (data []byte, err error) {
	type shadow V3UserRestoreParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V3UserRestoreParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
