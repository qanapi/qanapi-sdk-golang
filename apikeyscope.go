// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package qanapi

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/qanapi/qanapi-sdk-golang/internal/apijson"
	"github.com/qanapi/qanapi-sdk-golang/internal/requestconfig"
	"github.com/qanapi/qanapi-sdk-golang/option"
	"github.com/qanapi/qanapi-sdk-golang/packages/param"
	"github.com/qanapi/qanapi-sdk-golang/packages/respjson"
)

// APIKeyScopeService contains methods and other services that help with
// interacting with the qanapi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIKeyScopeService] method instead.
type APIKeyScopeService struct {
	Options []option.RequestOption
}

// NewAPIKeyScopeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAPIKeyScopeService(opts ...option.RequestOption) (r APIKeyScopeService) {
	r = APIKeyScopeService{}
	r.Options = opts
	return
}

// Retrieve scopes associated with an API Key
func (r *APIKeyScopeService) Get(ctx context.Context, apiKey int64, opts ...option.RequestOption) (res *[]APIKeyScopeGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("api-keys/%v/scopes", apiKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Attach scopes to an API Key
func (r *APIKeyScopeService) Attach(ctx context.Context, apiKey int64, body APIKeyScopeAttachParams, opts ...option.RequestOption) (res *APIKeyScopeAttachResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("api-keys/%v/scopes/attach", apiKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Detach scopes from an API Key
func (r *APIKeyScopeService) Detach(ctx context.Context, apiKey int64, body APIKeyScopeDetachParams, opts ...option.RequestOption) (res *APIKeyScopeDetachResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("api-keys/%v/scopes/detach", apiKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Sync scopes of an API Key
func (r *APIKeyScopeService) Sync(ctx context.Context, apiKey int64, body APIKeyScopeSyncParams, opts ...option.RequestOption) (res *APIKeyScopeSyncResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("api-keys/%v/scopes/sync", apiKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type APIKeyScopeGetResponse struct {
	ID        int64                       `json:"id"`
	CreatedAt time.Time                   `json:"created_at" format:"date-time"`
	Name      string                      `json:"name"`
	Pivot     APIKeyScopeGetResponsePivot `json:"pivot"`
	Route     string                      `json:"route"`
	UpdatedAt time.Time                   `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Pivot       respjson.Field
		Route       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKeyScopeGetResponse) RawJSON() string { return r.JSON.raw }
func (r *APIKeyScopeGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyScopeGetResponsePivot struct {
	APIKeyID int64 `json:"api_key_id"`
	ScopeID  int64 `json:"scope_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKeyID    respjson.Field
		ScopeID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKeyScopeGetResponsePivot) RawJSON() string { return r.JSON.raw }
func (r *APIKeyScopeGetResponsePivot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyScopeAttachResponse struct {
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKeyScopeAttachResponse) RawJSON() string { return r.JSON.raw }
func (r *APIKeyScopeAttachResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyScopeDetachResponse struct {
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKeyScopeDetachResponse) RawJSON() string { return r.JSON.raw }
func (r *APIKeyScopeDetachResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyScopeSyncResponse struct {
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKeyScopeSyncResponse) RawJSON() string { return r.JSON.raw }
func (r *APIKeyScopeSyncResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyScopeAttachParams struct {
	// List of scope IDs to attach
	ScopeIDs []int64 `json:"scope_ids,omitzero,required"`
	paramObj
}

func (r APIKeyScopeAttachParams) MarshalJSON() (data []byte, err error) {
	type shadow APIKeyScopeAttachParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIKeyScopeAttachParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyScopeDetachParams struct {
	// List of scope IDs to detach
	ScopeIDs []int64 `json:"scope_ids,omitzero,required"`
	paramObj
}

func (r APIKeyScopeDetachParams) MarshalJSON() (data []byte, err error) {
	type shadow APIKeyScopeDetachParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIKeyScopeDetachParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyScopeSyncParams struct {
	// List of scope IDs to sync
	ScopeIDs []int64 `json:"scope_ids,omitzero,required"`
	paramObj
}

func (r APIKeyScopeSyncParams) MarshalJSON() (data []byte, err error) {
	type shadow APIKeyScopeSyncParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIKeyScopeSyncParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
