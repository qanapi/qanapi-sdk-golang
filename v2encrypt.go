// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package qanapi

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/qanapi/qanapi-sdk-golang/internal/apijson"
	"github.com/qanapi/qanapi-sdk-golang/internal/requestconfig"
	"github.com/qanapi/qanapi-sdk-golang/option"
	"github.com/qanapi/qanapi-sdk-golang/packages/param"
	"github.com/qanapi/qanapi-sdk-golang/packages/respjson"
)

// V2EncryptService contains methods and other services that help with interacting
// with the qanapi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2EncryptService] method instead.
type V2EncryptService struct {
	Options []option.RequestOption
}

// NewV2EncryptService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2EncryptService(opts ...option.RequestOption) (r V2EncryptService) {
	r = V2EncryptService{}
	r.Options = opts
	return
}

// Encrypt data with optional ACL
func (r *V2EncryptService) EncryptData(ctx context.Context, body V2EncryptEncryptDataParams, opts ...option.RequestOption) (res *V2EncryptEncryptDataResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/encrypt"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// V2EncryptEncryptDataResponseUnion contains all possible properties and values
// from [string], [float64], [map[string]any], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat
// OfV2EncryptEncryptDataResponseEncryptEncryptDataResponseVariant2Item OfAnyArray]
type V2EncryptEncryptDataResponseUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfV2EncryptEncryptDataResponseEncryptEncryptDataResponseVariant2Item any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString                                                             respjson.Field
		OfFloat                                                              respjson.Field
		OfV2EncryptEncryptDataResponseEncryptEncryptDataResponseVariant2Item respjson.Field
		OfAnyArray                                                           respjson.Field
		raw                                                                  string
	} `json:"-"`
}

func (u V2EncryptEncryptDataResponseUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V2EncryptEncryptDataResponseUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V2EncryptEncryptDataResponseUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V2EncryptEncryptDataResponseUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V2EncryptEncryptDataResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *V2EncryptEncryptDataResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2EncryptEncryptDataParams struct {
	// The actual data to encrypt.
	//
	//   - Can be a scalar (string/number), object, or array.
	//   - If the value is an object or array, only the specified `sensitiveFields` are
	//     encrypted.
	Data   V2EncryptEncryptDataParamsDataUnion `json:"data,omitzero" api:"required"`
	Access V2EncryptEncryptDataParamsAccess    `json:"access,omitzero"`
	// Optional metadata describing the data's context.
	Attributes V2EncryptEncryptDataParamsAttributes `json:"attributes,omitzero"`
	// Laravel-style dot-notated paths to fields that should be encrypted.
	//
	// Supports:
	//
	// - Dot notation for nested fields: `user.profile.ssn`
	// - Wildcard `*` for arrays or dynamic keys: `users.*.token`
	//
	// Examples:
	//
	// - `password`
	// - `user.credentials.secret`
	// - `accounts.*.secret`
	// - `teams.*.members.*.email`
	SensitiveFields []string `json:"sensitiveFields,omitzero"`
	paramObj
}

func (r V2EncryptEncryptDataParams) MarshalJSON() (data []byte, err error) {
	type shadow V2EncryptEncryptDataParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2EncryptEncryptDataParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V2EncryptEncryptDataParamsDataUnion struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u V2EncryptEncryptDataParamsDataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfAnyMap, u.OfAnyArray)
}
func (u *V2EncryptEncryptDataParamsDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V2EncryptEncryptDataParamsDataUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	}
	return nil
}

type V2EncryptEncryptDataParamsAccess struct {
	// Access control list — list of user roles authorized to decrypt this data.
	ACL []string `json:"acl,omitzero"`
	paramObj
}

func (r V2EncryptEncryptDataParamsAccess) MarshalJSON() (data []byte, err error) {
	type shadow V2EncryptEncryptDataParamsAccess
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2EncryptEncryptDataParamsAccess) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional metadata describing the data's context.
type V2EncryptEncryptDataParamsAttributes struct {
	Owner param.Opt[string] `json:"owner,omitzero" format:"email"`
	// Any of "public", "internal", "confidential", "restricted".
	Classification string   `json:"classification,omitzero"`
	Tags           []string `json:"tags,omitzero"`
	paramObj
}

func (r V2EncryptEncryptDataParamsAttributes) MarshalJSON() (data []byte, err error) {
	type shadow V2EncryptEncryptDataParamsAttributes
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2EncryptEncryptDataParamsAttributes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2EncryptEncryptDataParamsAttributes](
		"classification", "public", "internal", "confidential", "restricted",
	)
}
