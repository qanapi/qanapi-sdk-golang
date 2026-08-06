// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package qanapi

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/qanapi/qanapi-sdk-golang/internal/apijson"
	"github.com/qanapi/qanapi-sdk-golang/internal/requestconfig"
	"github.com/qanapi/qanapi-sdk-golang/option"
	"github.com/qanapi/qanapi-sdk-golang/packages/respjson"
)

// V3APIKeyService contains methods and other services that help with interacting
// with the qanapi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV3APIKeyService] method instead.
type V3APIKeyService struct {
	Options []option.RequestOption
}

// NewV3APIKeyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV3APIKeyService(opts ...option.RequestOption) (r V3APIKeyService) {
	r = V3APIKeyService{}
	r.Options = opts
	return
}

// List API Keys
func (r *V3APIKeyService) List(ctx context.Context, opts ...option.RequestOption) (res *[]APIKey, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/api-keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Revoke API Key
func (r *V3APIKeyService) Revoke(ctx context.Context, apiKey int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("v3/api-keys/%v/revoke", apiKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Rotate API Key
func (r *V3APIKeyService) Rotate(ctx context.Context, apiKey int64, opts ...option.RequestOption) (res *V3APIKeyRotateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v3/api-keys/%v/rotate", apiKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Get API Key
func (r *V3APIKeyService) Show(ctx context.Context, apiKey int64, opts ...option.RequestOption) (res *APIKey, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v3/api-keys/%v", apiKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type V3APIKeyRotateResponse struct {
	Key string `json:"key"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3APIKeyRotateResponse) RawJSON() string { return r.JSON.raw }
func (r *V3APIKeyRotateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
