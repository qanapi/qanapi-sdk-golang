// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package qanapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/qanapi/qanapi-sdk-golang/internal/apijson"
	"github.com/qanapi/qanapi-sdk-golang/internal/requestconfig"
	"github.com/qanapi/qanapi-sdk-golang/option"
	"github.com/qanapi/qanapi-sdk-golang/packages/param"
	"github.com/qanapi/qanapi-sdk-golang/packages/respjson"
)

// V3ConfigurationService contains methods and other services that help with
// interacting with the qanapi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV3ConfigurationService] method instead.
type V3ConfigurationService struct {
	Options []option.RequestOption
}

// NewV3ConfigurationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV3ConfigurationService(opts ...option.RequestOption) (r V3ConfigurationService) {
	r = V3ConfigurationService{}
	r.Options = opts
	return
}

// Create configuration
func (r *V3ConfigurationService) New(ctx context.Context, body V3ConfigurationNewParams, opts ...option.RequestOption) (res *V3ConfigurationNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/configurations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update configuration
func (r *V3ConfigurationService) Update(ctx context.Context, configuration string, body V3ConfigurationUpdateParams, opts ...option.RequestOption) (res *V3ConfigurationUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if configuration == "" {
		err = errors.New("missing required configuration parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/configurations/%s", configuration)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// List configurations
func (r *V3ConfigurationService) List(ctx context.Context, opts ...option.RequestOption) (res *[]V3ConfigurationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/configurations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete configuration
func (r *V3ConfigurationService) Delete(ctx context.Context, configuration string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if configuration == "" {
		err = errors.New("missing required configuration parameter")
		return err
	}
	path := fmt.Sprintf("v3/configurations/%s", configuration)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get configuration
func (r *V3ConfigurationService) Show(ctx context.Context, configuration string, opts ...option.RequestOption) (res *V3ConfigurationShowResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if configuration == "" {
		err = errors.New("missing required configuration parameter")
		return nil, err
	}
	path := fmt.Sprintf("v3/configurations/%s", configuration)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type V3ConfigurationNewResponse struct {
	ID     string                            `json:"id" api:"required" format:"uuid"`
	Name   string                            `json:"name" api:"required"`
	Type   string                            `json:"type" api:"required"`
	Values []V3ConfigurationNewResponseValue `json:"values"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3ConfigurationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *V3ConfigurationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3ConfigurationNewResponseValue struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3ConfigurationNewResponseValue) RawJSON() string { return r.JSON.raw }
func (r *V3ConfigurationNewResponseValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3ConfigurationUpdateResponse struct {
	ID     string                               `json:"id" api:"required" format:"uuid"`
	Name   string                               `json:"name" api:"required"`
	Type   string                               `json:"type" api:"required"`
	Values []V3ConfigurationUpdateResponseValue `json:"values"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3ConfigurationUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *V3ConfigurationUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3ConfigurationUpdateResponseValue struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3ConfigurationUpdateResponseValue) RawJSON() string { return r.JSON.raw }
func (r *V3ConfigurationUpdateResponseValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3ConfigurationListResponse struct {
	ID     string                             `json:"id" api:"required" format:"uuid"`
	Name   string                             `json:"name" api:"required"`
	Type   string                             `json:"type" api:"required"`
	Values []V3ConfigurationListResponseValue `json:"values"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3ConfigurationListResponse) RawJSON() string { return r.JSON.raw }
func (r *V3ConfigurationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3ConfigurationListResponseValue struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3ConfigurationListResponseValue) RawJSON() string { return r.JSON.raw }
func (r *V3ConfigurationListResponseValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3ConfigurationShowResponse struct {
	ID     string                             `json:"id" api:"required" format:"uuid"`
	Name   string                             `json:"name" api:"required"`
	Type   string                             `json:"type" api:"required"`
	Values []V3ConfigurationShowResponseValue `json:"values"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3ConfigurationShowResponse) RawJSON() string { return r.JSON.raw }
func (r *V3ConfigurationShowResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3ConfigurationShowResponseValue struct {
	Key   string `json:"key" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3ConfigurationShowResponseValue) RawJSON() string { return r.JSON.raw }
func (r *V3ConfigurationShowResponseValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3ConfigurationNewParams struct {
	Name string `json:"name" api:"required"`
	// Any of "provider", "encryption".
	Type V3ConfigurationNewParamsType `json:"type,omitzero" api:"required"`
	// Required if type is 'provider'
	//
	// Any of "google".
	Provider V3ConfigurationNewParamsProvider `json:"provider,omitzero"`
	paramObj
}

func (r V3ConfigurationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V3ConfigurationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V3ConfigurationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3ConfigurationNewParamsType string

const (
	V3ConfigurationNewParamsTypeProvider   V3ConfigurationNewParamsType = "provider"
	V3ConfigurationNewParamsTypeEncryption V3ConfigurationNewParamsType = "encryption"
)

// Required if type is 'provider'
type V3ConfigurationNewParamsProvider string

const (
	V3ConfigurationNewParamsProviderGoogle V3ConfigurationNewParamsProvider = "google"
)

type V3ConfigurationUpdateParams struct {
	Name string `json:"name" api:"required"`
	paramObj
}

func (r V3ConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V3ConfigurationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V3ConfigurationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
