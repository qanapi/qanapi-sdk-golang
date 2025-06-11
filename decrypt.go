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

// DecryptService contains methods and other services that help with interacting
// with the qanapi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDecryptService] method instead.
type DecryptService struct {
	Options []option.RequestOption
}

// NewDecryptService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDecryptService(opts ...option.RequestOption) (r DecryptService) {
	r = DecryptService{}
	r.Options = opts
	return
}

// Decrypt encrypted payload
func (r *DecryptService) DecryptPayload(ctx context.Context, body DecryptDecryptPayloadParams, opts ...option.RequestOption) (res *DecryptDecryptPayloadResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	path := "decrypt"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// DecryptDecryptPayloadResponseUnion contains all possible properties and values
// from [string], [map[string]any], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfDecryptDecryptPayloadResponseMapItem OfAnyArray]
type DecryptDecryptPayloadResponseUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfDecryptDecryptPayloadResponseMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString                               respjson.Field
		OfDecryptDecryptPayloadResponseMapItem respjson.Field
		OfAnyArray                             respjson.Field
		raw                                    string
	} `json:"-"`
}

func (u DecryptDecryptPayloadResponseUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DecryptDecryptPayloadResponseUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DecryptDecryptPayloadResponseUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DecryptDecryptPayloadResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *DecryptDecryptPayloadResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DecryptDecryptPayloadParams struct {
	// The encrypted payload to decrypt.
	//
	// - Can be a string or an object/array with encrypted fields.
	// - Decryption is selective if `sensitiveFields` is provided.
	Data DecryptDecryptPayloadParamsDataUnion `json:"data,omitzero,required"`
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

func (r DecryptDecryptPayloadParams) MarshalJSON() (data []byte, err error) {
	type shadow DecryptDecryptPayloadParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DecryptDecryptPayloadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DecryptDecryptPayloadParamsDataUnion struct {
	OfString   param.Opt[string] `json:",omitzero,inline"`
	OfAnyMap   map[string]any    `json:",omitzero,inline"`
	OfAnyArray []any             `json:",omitzero,inline"`
	paramUnion
}

func (u DecryptDecryptPayloadParamsDataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfAnyMap, u.OfAnyArray)
}
func (u *DecryptDecryptPayloadParamsDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *DecryptDecryptPayloadParamsDataUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	}
	return nil
}
