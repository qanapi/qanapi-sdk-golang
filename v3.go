// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package qanapi

import (
	"time"

	"github.com/qanapi/qanapi-sdk-golang/internal/apijson"
	"github.com/qanapi/qanapi-sdk-golang/option"
	"github.com/qanapi/qanapi-sdk-golang/packages/respjson"
)

// V3Service contains methods and other services that help with interacting with
// the qanapi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV3Service] method instead.
type V3Service struct {
	Options        []option.RequestOption
	Roles          V3RoleService
	Configurations V3ConfigurationService
	Users          V3UserService
	APIKeys        V3APIKeyService
	Logs           V3LogService
	Encryption     V3EncryptionService
}

// NewV3Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV3Service(opts ...option.RequestOption) (r V3Service) {
	r = V3Service{}
	r.Options = opts
	r.Roles = NewV3RoleService(opts...)
	r.Configurations = NewV3ConfigurationService(opts...)
	r.Users = NewV3UserService(opts...)
	r.APIKeys = NewV3APIKeyService(opts...)
	r.Logs = NewV3LogService(opts...)
	r.Encryption = NewV3EncryptionService(opts...)
	return
}

type APIKey struct {
	ID     string `json:"id" api:"required"`
	Prefix string `json:"prefix" api:"required"`
	// Any of "active", "revoked".
	Status         APIKeyStatus    `json:"status" api:"required"`
	Configurations []Configuration `json:"configurations"`
	CreatedAt      time.Time       `json:"created_at" format:"date-time"`
	Permissions    []Permission    `json:"permissions"`
	RevokedAt      time.Time       `json:"revoked_at" api:"nullable" format:"date-time"`
	UpdatedAt      time.Time       `json:"updated_at" format:"date-time"`
	User           User            `json:"user"`
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
func (r APIKey) RawJSON() string { return r.JSON.raw }
func (r *APIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyStatus string

const (
	APIKeyStatusActive  APIKeyStatus = "active"
	APIKeyStatusRevoked APIKeyStatus = "revoked"
)

type Configuration struct {
	ID     string  `json:"id" api:"required" format:"uuid"`
	Name   string  `json:"name" api:"required"`
	Type   string  `json:"type" api:"required"`
	Values []Value `json:"values"`
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
func (r Configuration) RawJSON() string { return r.JSON.raw }
func (r *Configuration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Permission struct {
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Permission) RawJSON() string { return r.JSON.raw }
func (r *Permission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Role struct {
	Name        string       `json:"name" api:"required"`
	Description string       `json:"description" api:"nullable"`
	Permissions []Permission `json:"permissions"`
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
func (r Role) RawJSON() string { return r.JSON.raw }
func (r *Role) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type User struct {
	ID               int64     `json:"id" api:"required"`
	Email            string    `json:"email" api:"required" format:"email"`
	Name             string    `json:"name" api:"required"`
	CreatedAt        time.Time `json:"created_at" format:"date-time"`
	Roles            []Role    `json:"roles"`
	TwoFactorEnabled bool      `json:"two_factor_enabled"`
	UpdatedAt        time.Time `json:"updated_at" format:"date-time"`
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
func (r User) RawJSON() string { return r.JSON.raw }
func (r *User) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Value struct {
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
func (r Value) RawJSON() string { return r.JSON.raw }
func (r *Value) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
