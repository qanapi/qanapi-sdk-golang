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

// V2DecryptService contains methods and other services that help with interacting
// with the qanapi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2DecryptService] method instead.
type V2DecryptService struct {
	Options []option.RequestOption
}

// NewV2DecryptService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2DecryptService(opts ...option.RequestOption) (r V2DecryptService) {
	r = V2DecryptService{}
	r.Options = opts
	return
}

// Decrypt encrypted payload
func (r *V2DecryptService) DecryptPayload(ctx context.Context, body V2DecryptDecryptPayloadParams, opts ...option.RequestOption) (res *V2DecryptDecryptPayloadResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "decrypt"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// V2DecryptDecryptPayloadResponseUnion contains all possible properties and values
// from [string], [map[string]any], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfV2DecryptDecryptPayloadResponseDecryptDecryptPayloadResponseVariant1Item
// OfAnyArray]
type V2DecryptDecryptPayloadResponseUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfV2DecryptDecryptPayloadResponseDecryptDecryptPayloadResponseVariant1Item any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString                                                                   respjson.Field
		OfV2DecryptDecryptPayloadResponseDecryptDecryptPayloadResponseVariant1Item respjson.Field
		OfAnyArray                                                                 respjson.Field
		raw                                                                        string
	} `json:"-"`
}

func (u V2DecryptDecryptPayloadResponseUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V2DecryptDecryptPayloadResponseUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V2DecryptDecryptPayloadResponseUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V2DecryptDecryptPayloadResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *V2DecryptDecryptPayloadResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2DecryptDecryptPayloadParams struct {
	// The encrypted payload to decrypt.
	//
	// - Can be a string or an object/array with encrypted fields.
	// - Decryption is selective if `sensitiveFields` is provided.
	Data V2DecryptDecryptPayloadParamsDataUnion `json:"data,omitzero" api:"required"`
	// Laravel-style dot-notated paths to fields to decrypt.
	//
	// - Same syntax and behavior as in EncryptRequest.
	// - If omitted, all string values matching encryption prefix are attempted.
	//
	// Examples:
	//
	// - `user.ssn`
	// - `employees.*.payroll.token`
	SensitiveFields []string `json:"sensitiveFields,omitzero"`
	paramObj
}

func (r V2DecryptDecryptPayloadParams) MarshalJSON() (data []byte, err error) {
	type shadow V2DecryptDecryptPayloadParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2DecryptDecryptPayloadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type V2DecryptDecryptPayloadParamsDataUnion struct {
	OfString   param.Opt[string] `json:",omitzero,inline"`
	OfAnyMap   map[string]any    `json:",omitzero,inline"`
	OfAnyArray []any             `json:",omitzero,inline"`
	paramUnion
}

func (u V2DecryptDecryptPayloadParamsDataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfAnyMap, u.OfAnyArray)
}
func (u *V2DecryptDecryptPayloadParamsDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *V2DecryptDecryptPayloadParamsDataUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	}
	return nil
}
