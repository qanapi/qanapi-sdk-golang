// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package qanapi

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/qanapi/qanapi-sdk-golang/internal/apijson"
	"github.com/qanapi/qanapi-sdk-golang/internal/requestconfig"
	"github.com/qanapi/qanapi-sdk-golang/option"
	"github.com/qanapi/qanapi-sdk-golang/packages/param"
	"github.com/qanapi/qanapi-sdk-golang/packages/respjson"
)

// V2AuthService contains methods and other services that help with interacting
// with the qanapi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2AuthService] method instead.
type V2AuthService struct {
	Options []option.RequestOption
}

// NewV2AuthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV2AuthService(opts ...option.RequestOption) (r V2AuthService) {
	r = V2AuthService{}
	r.Options = opts
	return
}

// Authenticate user and return JWT
func (r *V2AuthService) Login(ctx context.Context, body V2AuthLoginParams, opts ...option.RequestOption) (res *V2AuthLoginResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "auth/login"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Log out the current user
func (r *V2AuthService) Logout(ctx context.Context, opts ...option.RequestOption) (res *V2AuthLogoutResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "auth/logout"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Refresh access token using refresh token
func (r *V2AuthService) RefreshToken(ctx context.Context, opts ...option.RequestOption) (res *V2AuthRefreshTokenResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "auth/refresh"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Retrieve user profile and roles
func (r *V2AuthService) GetUserDetails(ctx context.Context, opts ...option.RequestOption) (res *V2AuthGetUserDetailsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "auth/userdetails"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Revoke the current token
func (r *V2AuthService) RevokeToken(ctx context.Context, opts ...option.RequestOption) (res *V2AuthRevokeTokenResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "auth/revoke"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type V2AuthLoginResponse struct {
	// JWT Bearer token
	AccessToken string `json:"access_token"`
	// Token expiration in seconds
	ExpiresIn int64 `json:"expires_in"`
	// Token Type
	TokenType string `json:"token_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessToken respjson.Field
		ExpiresIn   respjson.Field
		TokenType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2AuthLoginResponse) RawJSON() string { return r.JSON.raw }
func (r *V2AuthLoginResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2AuthLogoutResponse struct {
	Message string `json:"message"`
	User    string `json:"user" format:"email"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		User        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2AuthLogoutResponse) RawJSON() string { return r.JSON.raw }
func (r *V2AuthLogoutResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2AuthRefreshTokenResponse struct {
	// JWT access token
	AccessToken string `json:"access_token"`
	// Token expiration time in seconds
	ExpiresIn int64  `json:"expires_in"`
	TokenType string `json:"token_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessToken respjson.Field
		ExpiresIn   respjson.Field
		TokenType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2AuthRefreshTokenResponse) RawJSON() string { return r.JSON.raw }
func (r *V2AuthRefreshTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2AuthGetUserDetailsResponse struct {
	ID              int64     `json:"id"`
	Email           string    `json:"email" format:"email"`
	EmailVerifiedAt time.Time `json:"email_verified_at" api:"nullable" format:"date-time"`
	FirstLogin      int64     `json:"first_login"`
	GravatarURL     string    `json:"gravatar_url" format:"uri"`
	Name            string    `json:"name"`
	Roles           []string  `json:"roles"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Email           respjson.Field
		EmailVerifiedAt respjson.Field
		FirstLogin      respjson.Field
		GravatarURL     respjson.Field
		Name            respjson.Field
		Roles           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2AuthGetUserDetailsResponse) RawJSON() string { return r.JSON.raw }
func (r *V2AuthGetUserDetailsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2AuthRevokeTokenResponse struct {
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2AuthRevokeTokenResponse) RawJSON() string { return r.JSON.raw }
func (r *V2AuthRevokeTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2AuthLoginParams struct {
	Email    string `json:"email" api:"required" format:"email"`
	Password string `json:"password" api:"required"`
	paramObj
}

func (r V2AuthLoginParams) MarshalJSON() (data []byte, err error) {
	type shadow V2AuthLoginParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2AuthLoginParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
