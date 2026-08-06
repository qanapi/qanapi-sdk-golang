// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package qanapi

import (
	"github.com/qanapi/qanapi-sdk-golang/option"
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
