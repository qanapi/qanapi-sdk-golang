// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package qanapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/qanapi/qanapi-sdk-golang/internal/apijson"
	shimjson "github.com/qanapi/qanapi-sdk-golang/internal/encoding/json"
	"github.com/qanapi/qanapi-sdk-golang/internal/requestconfig"
	"github.com/qanapi/qanapi-sdk-golang/option"
	"github.com/qanapi/qanapi-sdk-golang/packages/param"
)

// V3EncryptionService contains methods and other services that help with
// interacting with the qanapi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV3EncryptionService] method instead.
type V3EncryptionService struct {
	Options []option.RequestOption
}

// NewV3EncryptionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV3EncryptionService(opts ...option.RequestOption) (r V3EncryptionService) {
	r = V3EncryptionService{}
	r.Options = opts
	return
}

// Decrypt data
func (r *V3EncryptionService) Decrypt(ctx context.Context, proxy string, params V3EncryptionDecryptParams, opts ...option.RequestOption) (res *V3EncryptionDecryptResponse, err error) {
	if !param.IsOmitted(params.XQanapiFields) {
		opts = append(opts, option.WithHeader("x-qanapi-fields", fmt.Sprintf("%v", params.XQanapiFields.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if proxy == "" {
		err = errors.New("missing required proxy parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/encryption/%s/decrypt", proxy)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Encrypt data
func (r *V3EncryptionService) Encrypt(ctx context.Context, proxy string, params V3EncryptionEncryptParams, opts ...option.RequestOption) (res *V3EncryptionEncryptResponse, err error) {
	if !param.IsOmitted(params.XQanapiFields) {
		opts = append(opts, option.WithHeader("x-qanapi-fields", fmt.Sprintf("%v", params.XQanapiFields)))
	}
	if !param.IsOmitted(params.XQanapiDestination) {
		opts = append(opts, option.WithHeader("x-qanapi-destination", fmt.Sprintf("%v", params.XQanapiDestination.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if proxy == "" {
		err = errors.New("missing required proxy parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/encryption/%s/encrypt", proxy)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type V3EncryptionDecryptResponse map[string]any

type V3EncryptionEncryptResponse map[string]any

type V3EncryptionDecryptParams struct {
	// A JSON object to decrypt fields on. A maximum depth of 32 is allowed.
	Data          map[string]any
	XQanapiFields param.Opt[string] `header:"x-qanapi-fields,omitzero" json:"-"`
	paramObj
}

func (r V3EncryptionDecryptParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Data)
}
func (r *V3EncryptionDecryptParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3EncryptionEncryptParams struct {
	// A JSON object to encrypt fields on. A maximum depth of 32 is allowed.
	Data               map[string]any
	XQanapiFields      string            `header:"x-qanapi-fields" api:"required" json:"-"`
	XQanapiDestination param.Opt[string] `header:"x-qanapi-destination,omitzero" json:"-"`
	paramObj
}

func (r V3EncryptionEncryptParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Data)
}
func (r *V3EncryptionEncryptParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
