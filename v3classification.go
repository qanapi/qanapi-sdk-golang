// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package qanapi

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/qanapi/qanapi-sdk-golang/internal/apijson"
	"github.com/qanapi/qanapi-sdk-golang/internal/apiquery"
	"github.com/qanapi/qanapi-sdk-golang/internal/requestconfig"
	"github.com/qanapi/qanapi-sdk-golang/option"
	"github.com/qanapi/qanapi-sdk-golang/packages/param"
	"github.com/qanapi/qanapi-sdk-golang/packages/respjson"
)

// V3ClassificationService contains methods and other services that help with
// interacting with the qanapi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV3ClassificationService] method instead.
type V3ClassificationService struct {
	Options []option.RequestOption
}

// NewV3ClassificationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV3ClassificationService(opts ...option.RequestOption) (r V3ClassificationService) {
	r = V3ClassificationService{}
	r.Options = opts
	return
}

// Create classification
func (r *V3ClassificationService) New(ctx context.Context, body V3ClassificationNewParams, opts ...option.RequestOption) (res *V3ClassificationNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/classifications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update classification
func (r *V3ClassificationService) Update(ctx context.Context, classification int64, body V3ClassificationUpdateParams, opts ...option.RequestOption) (res *V3ClassificationUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v3/classifications/%v", classification)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List classifications
func (r *V3ClassificationService) List(ctx context.Context, query V3ClassificationListParams, opts ...option.RequestOption) (res *V3ClassificationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/classifications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete classification
func (r *V3ClassificationService) Delete(ctx context.Context, classification int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("v3/classifications/%v", classification)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get classification
func (r *V3ClassificationService) Show(ctx context.Context, classification int64, opts ...option.RequestOption) (res *V3ClassificationShowResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v3/classifications/%v", classification)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type V3ClassificationNewResponse struct {
	ID          int64                                 `json:"id" api:"required"`
	BgColor     string                                `json:"bg_color" api:"required"`
	FgColor     string                                `json:"fg_color" api:"required"`
	Name        string                                `json:"name" api:"required"`
	Slug        string                                `json:"slug" api:"required"`
	Description string                                `json:"description" api:"nullable"`
	Emoji       string                                `json:"emoji" api:"nullable"`
	Providers   []V3ClassificationNewResponseProvider `json:"providers"`
	Roles       []Role                                `json:"roles"`
	Users       []User                                `json:"users"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		BgColor     respjson.Field
		FgColor     respjson.Field
		Name        respjson.Field
		Slug        respjson.Field
		Description respjson.Field
		Emoji       respjson.Field
		Providers   respjson.Field
		Roles       respjson.Field
		Users       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3ClassificationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *V3ClassificationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3ClassificationNewResponseProvider struct {
	GwsGroups []GoogleGroup `json:"gws_groups"`
	Name      string        `json:"name"`
	Uuid      string        `json:"uuid" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GwsGroups   respjson.Field
		Name        respjson.Field
		Uuid        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3ClassificationNewResponseProvider) RawJSON() string { return r.JSON.raw }
func (r *V3ClassificationNewResponseProvider) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3ClassificationUpdateResponse struct {
	ID          int64                                    `json:"id" api:"required"`
	BgColor     string                                   `json:"bg_color" api:"required"`
	FgColor     string                                   `json:"fg_color" api:"required"`
	Name        string                                   `json:"name" api:"required"`
	Slug        string                                   `json:"slug" api:"required"`
	Description string                                   `json:"description" api:"nullable"`
	Emoji       string                                   `json:"emoji" api:"nullable"`
	Providers   []V3ClassificationUpdateResponseProvider `json:"providers"`
	Roles       []Role                                   `json:"roles"`
	Users       []User                                   `json:"users"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		BgColor     respjson.Field
		FgColor     respjson.Field
		Name        respjson.Field
		Slug        respjson.Field
		Description respjson.Field
		Emoji       respjson.Field
		Providers   respjson.Field
		Roles       respjson.Field
		Users       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3ClassificationUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *V3ClassificationUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3ClassificationUpdateResponseProvider struct {
	GwsGroups []GoogleGroup `json:"gws_groups"`
	Name      string        `json:"name"`
	Uuid      string        `json:"uuid" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GwsGroups   respjson.Field
		Name        respjson.Field
		Uuid        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3ClassificationUpdateResponseProvider) RawJSON() string { return r.JSON.raw }
func (r *V3ClassificationUpdateResponseProvider) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3ClassificationListResponse struct {
	CurrentPage  int64                              `json:"current_page"`
	Data         []V3ClassificationListResponseData `json:"data"`
	FirstPageURL string                             `json:"first_page_url" format:"uri"`
	From         int64                              `json:"from" api:"nullable"`
	LastPage     int64                              `json:"last_page"`
	LastPageURL  string                             `json:"last_page_url" format:"uri"`
	Links        []V3ClassificationListResponseLink `json:"links"`
	NextPageURL  string                             `json:"next_page_url" api:"nullable" format:"uri"`
	Path         string                             `json:"path"`
	PerPage      int64                              `json:"per_page"`
	PrevPageURL  string                             `json:"prev_page_url" api:"nullable" format:"uri"`
	To           int64                              `json:"to" api:"nullable"`
	Total        int64                              `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentPage  respjson.Field
		Data         respjson.Field
		FirstPageURL respjson.Field
		From         respjson.Field
		LastPage     respjson.Field
		LastPageURL  respjson.Field
		Links        respjson.Field
		NextPageURL  respjson.Field
		Path         respjson.Field
		PerPage      respjson.Field
		PrevPageURL  respjson.Field
		To           respjson.Field
		Total        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3ClassificationListResponse) RawJSON() string { return r.JSON.raw }
func (r *V3ClassificationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3ClassificationListResponseData struct {
	ID          int64                                      `json:"id" api:"required"`
	BgColor     string                                     `json:"bg_color" api:"required"`
	FgColor     string                                     `json:"fg_color" api:"required"`
	Name        string                                     `json:"name" api:"required"`
	Slug        string                                     `json:"slug" api:"required"`
	Description string                                     `json:"description" api:"nullable"`
	Emoji       string                                     `json:"emoji" api:"nullable"`
	Providers   []V3ClassificationListResponseDataProvider `json:"providers"`
	Roles       []Role                                     `json:"roles"`
	Users       []User                                     `json:"users"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		BgColor     respjson.Field
		FgColor     respjson.Field
		Name        respjson.Field
		Slug        respjson.Field
		Description respjson.Field
		Emoji       respjson.Field
		Providers   respjson.Field
		Roles       respjson.Field
		Users       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3ClassificationListResponseData) RawJSON() string { return r.JSON.raw }
func (r *V3ClassificationListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3ClassificationListResponseDataProvider struct {
	GwsGroups []GoogleGroup `json:"gws_groups"`
	Name      string        `json:"name"`
	Uuid      string        `json:"uuid" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GwsGroups   respjson.Field
		Name        respjson.Field
		Uuid        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3ClassificationListResponseDataProvider) RawJSON() string { return r.JSON.raw }
func (r *V3ClassificationListResponseDataProvider) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3ClassificationListResponseLink struct {
	Active bool   `json:"active"`
	Label  string `json:"label"`
	Page   int64  `json:"page" api:"nullable"`
	URL    string `json:"url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active      respjson.Field
		Label       respjson.Field
		Page        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3ClassificationListResponseLink) RawJSON() string { return r.JSON.raw }
func (r *V3ClassificationListResponseLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3ClassificationShowResponse struct {
	ID          int64                                  `json:"id" api:"required"`
	BgColor     string                                 `json:"bg_color" api:"required"`
	FgColor     string                                 `json:"fg_color" api:"required"`
	Name        string                                 `json:"name" api:"required"`
	Slug        string                                 `json:"slug" api:"required"`
	Description string                                 `json:"description" api:"nullable"`
	Emoji       string                                 `json:"emoji" api:"nullable"`
	Providers   []V3ClassificationShowResponseProvider `json:"providers"`
	Roles       []Role                                 `json:"roles"`
	Users       []User                                 `json:"users"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		BgColor     respjson.Field
		FgColor     respjson.Field
		Name        respjson.Field
		Slug        respjson.Field
		Description respjson.Field
		Emoji       respjson.Field
		Providers   respjson.Field
		Roles       respjson.Field
		Users       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3ClassificationShowResponse) RawJSON() string { return r.JSON.raw }
func (r *V3ClassificationShowResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3ClassificationShowResponseProvider struct {
	GwsGroups []GoogleGroup `json:"gws_groups"`
	Name      string        `json:"name"`
	Uuid      string        `json:"uuid" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GwsGroups   respjson.Field
		Name        respjson.Field
		Uuid        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3ClassificationShowResponseProvider) RawJSON() string { return r.JSON.raw }
func (r *V3ClassificationShowResponseProvider) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3ClassificationNewParams struct {
	BgColor     string            `json:"bg_color" api:"required"`
	FgColor     string            `json:"fg_color" api:"required"`
	Name        string            `json:"name" api:"required"`
	Description param.Opt[string] `json:"description,omitzero"`
	Emoji       param.Opt[string] `json:"emoji,omitzero"`
	// Required if gws_groups is provided.
	ProviderContainerID param.Opt[int64]   `json:"provider_container_id,omitzero"`
	GwsGroups           []GoogleGroupParam `json:"gws_groups,omitzero"`
	// Array of role IDs.
	Roles []int64 `json:"roles,omitzero"`
	// Array of user IDs.
	Users []int64 `json:"users,omitzero"`
	paramObj
}

func (r V3ClassificationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V3ClassificationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V3ClassificationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3ClassificationUpdateParams struct {
	BgColor     string            `json:"bg_color" api:"required"`
	FgColor     string            `json:"fg_color" api:"required"`
	Name        string            `json:"name" api:"required"`
	Description param.Opt[string] `json:"description,omitzero"`
	Emoji       param.Opt[string] `json:"emoji,omitzero"`
	// Required if gws_groups is provided.
	ProviderContainerID param.Opt[int64]   `json:"provider_container_id,omitzero"`
	GwsGroups           []GoogleGroupParam `json:"gws_groups,omitzero"`
	// Array of role IDs.
	Roles []int64 `json:"roles,omitzero"`
	// Array of user IDs.
	Users []int64 `json:"users,omitzero"`
	paramObj
}

func (r V3ClassificationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V3ClassificationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V3ClassificationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3ClassificationListParams struct {
	PerPage param.Opt[int64]  `query:"per_page,omitzero" json:"-"`
	Search  param.Opt[string] `query:"search,omitzero" json:"-"`
	Sort    param.Opt[string] `query:"sort,omitzero" json:"-"`
	User    param.Opt[int64]  `query:"user,omitzero" json:"-"`
	// Any of "asc", "desc".
	Direction V3ClassificationListParamsDirection `query:"direction,omitzero" json:"-"`
	Providers []int64                             `query:"providers,omitzero" json:"-"`
	Roles     []int64                             `query:"roles,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V3ClassificationListParams]'s query parameters as
// `url.Values`.
func (r V3ClassificationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V3ClassificationListParamsDirection string

const (
	V3ClassificationListParamsDirectionAsc  V3ClassificationListParamsDirection = "asc"
	V3ClassificationListParamsDirectionDesc V3ClassificationListParamsDirection = "desc"
)
