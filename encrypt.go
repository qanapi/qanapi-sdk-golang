// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package qanapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/qanapi/qanapi-sdk-golang/internal/apijson"
	"github.com/qanapi/qanapi-sdk-golang/internal/requestconfig"
	"github.com/qanapi/qanapi-sdk-golang/option"
	"github.com/qanapi/qanapi-sdk-golang/packages/param"
	"github.com/qanapi/qanapi-sdk-golang/packages/respjson"
)

// EncryptService contains methods and other services that help with interacting
// with the qanapi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEncryptService] method instead.
type EncryptService struct {
	Options []option.RequestOption
}

// NewEncryptService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEncryptService(opts ...option.RequestOption) (r EncryptService) {
	r = EncryptService{}
	r.Options = opts
	return
}

// Encrypt data with optional ACL
func (r *EncryptService) EncryptData(ctx context.Context, body EncryptEncryptDataParams, opts ...option.RequestOption) (res *EncryptEncryptDataResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	path := "encrypt"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The property Data is required.
type EncryptedDataParam struct {
	// The actual data to encrypt.
	//
	//   - Can be a scalar (string/number), object, or array.
	//   - If the value is an object or array, only the specified `sensitiveFields` are
	//     encrypted.
	Data   EncryptedDataDataUnionParam `json:"data,omitzero,required"`
	Access EncryptedDataAccessParam    `json:"access,omitzero"`
	// Optional metadata describing the data's context.
	Attributes EncryptedDataAttributesParam `json:"attributes,omitzero"`
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

func (r EncryptedDataParam) MarshalJSON() (data []byte, err error) {
	type shadow EncryptedDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EncryptedDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EncryptedDataDataUnionParam struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u EncryptedDataDataUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfAnyMap, u.OfAnyArray)
}
func (u *EncryptedDataDataUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EncryptedDataDataUnionParam) asAny() any {
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

type EncryptedDataAccessParam struct {
	// Access control list — list of user roles authorized to decrypt this data.
	ACL []string `json:"acl,omitzero"`
	paramObj
}

func (r EncryptedDataAccessParam) MarshalJSON() (data []byte, err error) {
	type shadow EncryptedDataAccessParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EncryptedDataAccessParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional metadata describing the data's context.
type EncryptedDataAttributesParam struct {
	Owner param.Opt[string] `json:"owner,omitzero" format:"email"`
	// Any of "public", "internal", "confidential", "restricted".
	Classification string   `json:"classification,omitzero"`
	Tags           []string `json:"tags,omitzero"`
	paramObj
}

func (r EncryptedDataAttributesParam) MarshalJSON() (data []byte, err error) {
	type shadow EncryptedDataAttributesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EncryptedDataAttributesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EncryptedDataAttributesParam](
		"classification", "public", "internal", "confidential", "restricted",
	)
}

// EncryptEncryptDataResponseUnion contains all possible properties and values from
// [string], [float64], [map[string]any], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfEncryptEncryptDataResponseMapItem OfAnyArray]
type EncryptEncryptDataResponseUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfEncryptEncryptDataResponseMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString                            respjson.Field
		OfFloat                             respjson.Field
		OfEncryptEncryptDataResponseMapItem respjson.Field
		OfAnyArray                          respjson.Field
		raw                                 string
	} `json:"-"`
}

func (u EncryptEncryptDataResponseUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EncryptEncryptDataResponseUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EncryptEncryptDataResponseUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EncryptEncryptDataResponseUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EncryptEncryptDataResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *EncryptEncryptDataResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EncryptEncryptDataParams struct {
	EncryptedData EncryptedDataParam
	paramObj
}

func (r EncryptEncryptDataParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.EncryptedData)
}
func (r *EncryptEncryptDataParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.EncryptedData)
}
