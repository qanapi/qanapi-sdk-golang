// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package qanapi

import (
	"github.com/qanapi/qanapi-sdk-golang/option"
)

// V2Service contains methods and other services that help with interacting with
// the qanapi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2Service] method instead.
type V2Service struct {
	Options []option.RequestOption
	Auth    V2AuthService
	Encrypt V2EncryptService
	Decrypt V2DecryptService
	APIKeys V2APIKeyService
}

// NewV2Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV2Service(opts ...option.RequestOption) (r V2Service) {
	r = V2Service{}
	r.Options = opts
	r.Auth = NewV2AuthService(opts...)
	r.Encrypt = NewV2EncryptService(opts...)
	r.Decrypt = NewV2DecryptService(opts...)
	r.APIKeys = NewV2APIKeyService(opts...)
	return
}
