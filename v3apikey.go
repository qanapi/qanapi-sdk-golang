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
	"github.com/qanapi/qanapi-sdk-golang/packages/respjson"
)

// V3APIKeyService contains methods and other services that help with interacting
// with the qanapi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV3APIKeyService] method instead.
type V3APIKeyService struct {
	Options []option.RequestOption
}

// NewV3APIKeyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV3APIKeyService(opts ...option.RequestOption) (r V3APIKeyService) {
	r = V3APIKeyService{}
	r.Options = opts
	return
}

// List API Keys
func (r *V3APIKeyService) List(ctx context.Context, opts ...option.RequestOption) (res *[]V3APIKeyListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/api-keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Revoke API Key
func (r *V3APIKeyService) Revoke(ctx context.Context, apiKey int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("v3/api-keys/%v/revoke", apiKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Rotate API Key
func (r *V3APIKeyService) Rotate(ctx context.Context, apiKey int64, opts ...option.RequestOption) (res *V3APIKeyRotateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v3/api-keys/%v/rotate", apiKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Get API Key
func (r *V3APIKeyService) Show(ctx context.Context, apiKey int64, opts ...option.RequestOption) (res *V3APIKeyShowResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v3/api-keys/%v", apiKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type V3APIKeyListResponse struct {
	ID     string `json:"id" api:"required"`
	Prefix string `json:"prefix" api:"required"`
	// Any of "active", "revoked".
	Status         V3APIKeyListResponseStatus          `json:"status" api:"required"`
	Configurations []V3APIKeyListResponseConfiguration `json:"configurations"`
	CreatedAt      time.Time                           `json:"created_at" format:"date-time"`
	Permissions    []V3APIKeyListResponsePermission    `json:"permissions"`
	RevokedAt      time.Time                           `json:"revoked_at" api:"nullable" format:"date-time"`
	UpdatedAt      time.Time                           `json:"updated_at" format:"date-time"`
	User           V3APIKeyListResponseUser            `json:"user"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Prefix         respjson.Field
		Status         respjson.Field
		Configurations respjson.Field
		CreatedAt      respjson.Field
		Permissions    respjson.Field
		RevokedAt      respjson.Field
		UpdatedAt      respjson.Field
		User           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3APIKeyListResponse) RawJSON() string { return r.JSON.raw }
func (r *V3APIKeyListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3APIKeyListResponseStatus string

const (
	V3APIKeyListResponseStatusActive  V3APIKeyListResponseStatus = "active"
	V3APIKeyListResponseStatusRevoked V3APIKeyListResponseStatus = "revoked"
)

type V3APIKeyListResponseConfiguration struct {
	ID     string                                   `json:"id" api:"required" format:"uuid"`
	Name   string                                   `json:"name" api:"required"`
	Type   string                                   `json:"type" api:"required"`
	Values []V3APIKeyListResponseConfigurationValue `json:"values"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3APIKeyListResponseConfiguration) RawJSON() string { return r.JSON.raw }
func (r *V3APIKeyListResponseConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3APIKeyListResponseConfigurationValue struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3APIKeyListResponseConfigurationValue) RawJSON() string { return r.JSON.raw }
func (r *V3APIKeyListResponseConfigurationValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3APIKeyListResponsePermission struct {
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3APIKeyListResponsePermission) RawJSON() string { return r.JSON.raw }
func (r *V3APIKeyListResponsePermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3APIKeyListResponseUser struct {
	ID               int64                          `json:"id" api:"required"`
	Email            string                         `json:"email" api:"required" format:"email"`
	Name             string                         `json:"name" api:"required"`
	CreatedAt        time.Time                      `json:"created_at" format:"date-time"`
	Roles            []V3APIKeyListResponseUserRole `json:"roles"`
	TwoFactorEnabled bool                           `json:"two_factor_enabled"`
	UpdatedAt        time.Time                      `json:"updated_at" format:"date-time"`
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
func (r V3APIKeyListResponseUser) RawJSON() string { return r.JSON.raw }
func (r *V3APIKeyListResponseUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3APIKeyListResponseUserRole struct {
	Name        string                                   `json:"name" api:"required"`
	Description string                                   `json:"description" api:"nullable"`
	Permissions []V3APIKeyListResponseUserRolePermission `json:"permissions"`
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
func (r V3APIKeyListResponseUserRole) RawJSON() string { return r.JSON.raw }
func (r *V3APIKeyListResponseUserRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3APIKeyListResponseUserRolePermission struct {
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3APIKeyListResponseUserRolePermission) RawJSON() string { return r.JSON.raw }
func (r *V3APIKeyListResponseUserRolePermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3APIKeyRotateResponse struct {
	Key string `json:"key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3APIKeyRotateResponse) RawJSON() string { return r.JSON.raw }
func (r *V3APIKeyRotateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3APIKeyShowResponse struct {
	ID     string `json:"id" api:"required"`
	Prefix string `json:"prefix" api:"required"`
	// Any of "active", "revoked".
	Status         V3APIKeyShowResponseStatus          `json:"status" api:"required"`
	Configurations []V3APIKeyShowResponseConfiguration `json:"configurations"`
	CreatedAt      time.Time                           `json:"created_at" format:"date-time"`
	Permissions    []V3APIKeyShowResponsePermission    `json:"permissions"`
	RevokedAt      time.Time                           `json:"revoked_at" api:"nullable" format:"date-time"`
	UpdatedAt      time.Time                           `json:"updated_at" format:"date-time"`
	User           V3APIKeyShowResponseUser            `json:"user"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Prefix         respjson.Field
		Status         respjson.Field
		Configurations respjson.Field
		CreatedAt      respjson.Field
		Permissions    respjson.Field
		RevokedAt      respjson.Field
		UpdatedAt      respjson.Field
		User           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3APIKeyShowResponse) RawJSON() string { return r.JSON.raw }
func (r *V3APIKeyShowResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3APIKeyShowResponseStatus string

const (
	V3APIKeyShowResponseStatusActive  V3APIKeyShowResponseStatus = "active"
	V3APIKeyShowResponseStatusRevoked V3APIKeyShowResponseStatus = "revoked"
)

type V3APIKeyShowResponseConfiguration struct {
	ID     string                                   `json:"id" api:"required" format:"uuid"`
	Name   string                                   `json:"name" api:"required"`
	Type   string                                   `json:"type" api:"required"`
	Values []V3APIKeyShowResponseConfigurationValue `json:"values"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3APIKeyShowResponseConfiguration) RawJSON() string { return r.JSON.raw }
func (r *V3APIKeyShowResponseConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3APIKeyShowResponseConfigurationValue struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3APIKeyShowResponseConfigurationValue) RawJSON() string { return r.JSON.raw }
func (r *V3APIKeyShowResponseConfigurationValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3APIKeyShowResponsePermission struct {
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3APIKeyShowResponsePermission) RawJSON() string { return r.JSON.raw }
func (r *V3APIKeyShowResponsePermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3APIKeyShowResponseUser struct {
	ID               int64                          `json:"id" api:"required"`
	Email            string                         `json:"email" api:"required" format:"email"`
	Name             string                         `json:"name" api:"required"`
	CreatedAt        time.Time                      `json:"created_at" format:"date-time"`
	Roles            []V3APIKeyShowResponseUserRole `json:"roles"`
	TwoFactorEnabled bool                           `json:"two_factor_enabled"`
	UpdatedAt        time.Time                      `json:"updated_at" format:"date-time"`
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
func (r V3APIKeyShowResponseUser) RawJSON() string { return r.JSON.raw }
func (r *V3APIKeyShowResponseUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3APIKeyShowResponseUserRole struct {
	Name        string                                   `json:"name" api:"required"`
	Description string                                   `json:"description" api:"nullable"`
	Permissions []V3APIKeyShowResponseUserRolePermission `json:"permissions"`
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
func (r V3APIKeyShowResponseUserRole) RawJSON() string { return r.JSON.raw }
func (r *V3APIKeyShowResponseUserRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3APIKeyShowResponseUserRolePermission struct {
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3APIKeyShowResponseUserRolePermission) RawJSON() string { return r.JSON.raw }
func (r *V3APIKeyShowResponseUserRolePermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
