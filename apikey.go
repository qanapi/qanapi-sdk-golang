// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package qanapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/qanapi/qanapi-sdk-golang/internal/apijson"
	"github.com/qanapi/qanapi-sdk-golang/internal/requestconfig"
	"github.com/qanapi/qanapi-sdk-golang/option"
	"github.com/qanapi/qanapi-sdk-golang/packages/respjson"
)

// APIKeyService contains methods and other services that help with interacting
// with the qanapi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIKeyService] method instead.
type APIKeyService struct {
	Options []option.RequestOption
	Scopes  APIKeyScopeService
}

// NewAPIKeyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAPIKeyService(opts ...option.RequestOption) (r APIKeyService) {
	r = APIKeyService{}
	r.Options = opts
	r.Scopes = NewAPIKeyScopeService(opts...)
	return
}

// Revoke an API Key
func (r *APIKeyService) Revoke(ctx context.Context, apiKey string, opts ...option.RequestOption) (res *APIKeyRevokeResponse, err error) {
	opts = append(r.Options[:], opts...)
	if apiKey == "" {
		err = errors.New("missing required apiKey parameter")
		return
	}
	path := fmt.Sprintf("api-keys/%s/revoke", apiKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

// Rotate an API Key
func (r *APIKeyService) Rotate(ctx context.Context, apiKey string, opts ...option.RequestOption) (res *APIKeyRotateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if apiKey == "" {
		err = errors.New("missing required apiKey parameter")
		return
	}
	path := fmt.Sprintf("api-keys/%s/rotate", apiKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

type APIKeyRevokeResponse struct {
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKeyRevokeResponse) RawJSON() string { return r.JSON.raw }
func (r *APIKeyRevokeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyRotateResponse struct {
	APIKey  string `json:"api_key"`
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKey      respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKeyRotateResponse) RawJSON() string { return r.JSON.raw }
func (r *APIKeyRotateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
