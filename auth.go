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

// AuthService contains methods and other services that help with interacting with
// the qanapi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthService] method instead.
type AuthService struct {
	Options []option.RequestOption
}

// NewAuthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAuthService(opts ...option.RequestOption) (r AuthService) {
	r = AuthService{}
	r.Options = opts
	return
}

// Authenticate user and return JWT
func (r *AuthService) Login(ctx context.Context, body AuthLoginParams, opts ...option.RequestOption) (res *AuthLoginResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "auth/login"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Log out the current user
func (r *AuthService) Logout(ctx context.Context, opts ...option.RequestOption) (res *AuthLogoutResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "auth/logout"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Refresh access token using refresh token
func (r *AuthService) RefreshToken(ctx context.Context, opts ...option.RequestOption) (res *AuthRefreshTokenResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "auth/refresh"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Retrieve user profile and roles
func (r *AuthService) GetUserDetails(ctx context.Context, opts ...option.RequestOption) (res *AuthGetUserDetailsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "auth/userdetails"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Revoke the current token
func (r *AuthService) RevokeToken(ctx context.Context, opts ...option.RequestOption) (res *AuthRevokeTokenResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "auth/revoke"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type AuthLoginResponse struct {
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
func (r AuthLoginResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthLoginResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthLogoutResponse struct {
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
func (r AuthLogoutResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthLogoutResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthRefreshTokenResponse struct {
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
func (r AuthRefreshTokenResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthRefreshTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthGetUserDetailsResponse struct {
	ID              int64     `json:"id"`
	Email           string    `json:"email" format:"email"`
	EmailVerifiedAt time.Time `json:"email_verified_at,nullable" format:"date-time"`
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
func (r AuthGetUserDetailsResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthGetUserDetailsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthRevokeTokenResponse struct {
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthRevokeTokenResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthRevokeTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthLoginParams struct {
	Email    string `json:"email,required" format:"email"`
	Password string `json:"password,required"`
	paramObj
}

func (r AuthLoginParams) MarshalJSON() (data []byte, err error) {
	type shadow AuthLoginParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AuthLoginParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
