// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package qanapi

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/qanapi-go/internal/apijson"
	"github.com/stainless-sdks/qanapi-go/internal/requestconfig"
	"github.com/stainless-sdks/qanapi-go/option"
	"github.com/stainless-sdks/qanapi-go/packages/param"
	"github.com/stainless-sdks/qanapi-go/packages/respjson"
)

// ScopeService contains methods and other services that help with interacting with
// the qanapi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScopeService] method instead.
type ScopeService struct {
	Options []option.RequestOption
}

// NewScopeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewScopeService(opts ...option.RequestOption) (r ScopeService) {
	r = ScopeService{}
	r.Options = opts
	return
}

// Create a new scope
func (r *ScopeService) New(ctx context.Context, body ScopeNewParams, opts ...option.RequestOption) (res *ScopeNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "scopes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a specific scope
func (r *ScopeService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *ScopeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("scopes/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a scope
func (r *ScopeService) Update(ctx context.Context, id int64, body ScopeUpdateParams, opts ...option.RequestOption) (res *ScopeUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("scopes/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List all scopes
func (r *ScopeService) List(ctx context.Context, opts ...option.RequestOption) (res *[]ScopeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "scopes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a scope
func (r *ScopeService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (res *ScopeDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("scopes/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type ScopeNewResponse struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	Name      string    `json:"name"`
	Route     string    `json:"route"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Route       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScopeNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ScopeNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScopeGetResponse struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	Name      string    `json:"name"`
	Route     string    `json:"route"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Route       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScopeGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ScopeGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScopeUpdateResponse struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	Name      string    `json:"name"`
	Route     string    `json:"route"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Route       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScopeUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ScopeUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScopeListResponse struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	Name      string    `json:"name"`
	Route     string    `json:"route"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Route       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScopeListResponse) RawJSON() string { return r.JSON.raw }
func (r *ScopeListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScopeDeleteResponse struct {
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScopeDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ScopeDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScopeNewParams struct {
	Name  string `json:"name,required"`
	Route string `json:"route,required"`
	paramObj
}

func (r ScopeNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ScopeNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScopeNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScopeUpdateParams struct {
	Name  param.Opt[string] `json:"name,omitzero"`
	Route param.Opt[string] `json:"route,omitzero"`
	paramObj
}

func (r ScopeUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ScopeUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScopeUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
