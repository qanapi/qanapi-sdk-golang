// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package qanapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/qanapi/qanapi-sdk-golang/internal/apijson"
	"github.com/qanapi/qanapi-sdk-golang/internal/apiquery"
	"github.com/qanapi/qanapi-sdk-golang/internal/requestconfig"
	"github.com/qanapi/qanapi-sdk-golang/option"
	"github.com/qanapi/qanapi-sdk-golang/packages/param"
	"github.com/qanapi/qanapi-sdk-golang/packages/respjson"
)

// V3LogService contains methods and other services that help with interacting with
// the qanapi API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV3LogService] method instead.
type V3LogService struct {
	Options []option.RequestOption
}

// NewV3LogService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV3LogService(opts ...option.RequestOption) (r V3LogService) {
	r = V3LogService{}
	r.Options = opts
	return
}

// Get activity logs
func (r *V3LogService) Activity(ctx context.Context, query V3LogActivityParams, opts ...option.RequestOption) (res *V3LogActivityResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/logs/activity"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get API logs
func (r *V3LogService) API(ctx context.Context, query V3LogAPIParams, opts ...option.RequestOption) (res *V3LogAPIResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/logs/api"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get Qanapi Flow logs
func (r *V3LogService) QanapiFlow(ctx context.Context, query V3LogQanapiFlowParams, opts ...option.RequestOption) (res *V3LogQanapiFlowResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/logs/qanapi-flow"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get unified logs
func (r *V3LogService) Unified(ctx context.Context, query V3LogUnifiedParams, opts ...option.RequestOption) (res *V3LogUnifiedResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v3/logs/unified"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type V3LogActivityResponse struct {
	CurrentPage  int64                       `json:"current_page"`
	Data         []V3LogActivityResponseData `json:"data"`
	FirstPageURL string                      `json:"first_page_url" format:"uri"`
	From         int64                       `json:"from" api:"nullable"`
	LastPage     int64                       `json:"last_page"`
	LastPageURL  string                      `json:"last_page_url" format:"uri"`
	Links        []V3LogActivityResponseLink `json:"links"`
	NextPageURL  string                      `json:"next_page_url" api:"nullable" format:"uri"`
	Path         string                      `json:"path"`
	PerPage      int64                       `json:"per_page"`
	PrevPageURL  string                      `json:"prev_page_url" api:"nullable" format:"uri"`
	To           int64                       `json:"to" api:"nullable"`
	Total        int64                       `json:"total"`
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
func (r V3LogActivityResponse) RawJSON() string { return r.JSON.raw }
func (r *V3LogActivityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogActivityResponseData struct {
	Action      string                        `json:"action"`
	Description string                        `json:"description"`
	IP          string                        `json:"ip" api:"nullable"`
	Timestamp   time.Time                     `json:"timestamp" format:"date-time"`
	User        V3LogActivityResponseDataUser `json:"user"`
	When        string                        `json:"when"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Description respjson.Field
		IP          respjson.Field
		Timestamp   respjson.Field
		User        respjson.Field
		When        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3LogActivityResponseData) RawJSON() string { return r.JSON.raw }
func (r *V3LogActivityResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogActivityResponseDataUser struct {
	ID               int64                               `json:"id" api:"required"`
	Email            string                              `json:"email" api:"required" format:"email"`
	Name             string                              `json:"name" api:"required"`
	CreatedAt        time.Time                           `json:"created_at" format:"date-time"`
	Roles            []V3LogActivityResponseDataUserRole `json:"roles"`
	TwoFactorEnabled bool                                `json:"two_factor_enabled"`
	UpdatedAt        time.Time                           `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Email            respjson.Field
		Name             respjson.Field
		CreatedAt        respjson.Field
		Roles            respjson.Field
		TwoFactorEnabled respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3LogActivityResponseDataUser) RawJSON() string { return r.JSON.raw }
func (r *V3LogActivityResponseDataUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogActivityResponseDataUserRole struct {
	Name        string                                        `json:"name" api:"required"`
	Description string                                        `json:"description" api:"nullable"`
	Permissions []V3LogActivityResponseDataUserRolePermission `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Description respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3LogActivityResponseDataUserRole) RawJSON() string { return r.JSON.raw }
func (r *V3LogActivityResponseDataUserRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogActivityResponseDataUserRolePermission struct {
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3LogActivityResponseDataUserRolePermission) RawJSON() string { return r.JSON.raw }
func (r *V3LogActivityResponseDataUserRolePermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogActivityResponseLink struct {
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
func (r V3LogActivityResponseLink) RawJSON() string { return r.JSON.raw }
func (r *V3LogActivityResponseLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogAPIResponse struct {
	CurrentPage  int64                  `json:"current_page"`
	Data         []V3LogAPIResponseData `json:"data"`
	FirstPageURL string                 `json:"first_page_url" format:"uri"`
	From         int64                  `json:"from" api:"nullable"`
	LastPage     int64                  `json:"last_page"`
	LastPageURL  string                 `json:"last_page_url" format:"uri"`
	Links        []V3LogAPIResponseLink `json:"links"`
	NextPageURL  string                 `json:"next_page_url" api:"nullable" format:"uri"`
	Path         string                 `json:"path"`
	PerPage      int64                  `json:"per_page"`
	PrevPageURL  string                 `json:"prev_page_url" api:"nullable" format:"uri"`
	To           int64                  `json:"to" api:"nullable"`
	Total        int64                  `json:"total"`
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
func (r V3LogAPIResponse) RawJSON() string { return r.JSON.raw }
func (r *V3LogAPIResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogAPIResponseData struct {
	APIKey          V3LogAPIResponseDataAPIKey `json:"api_key"`
	APIKeyID        int64                      `json:"api_key_id"`
	ConfigurationID int64                      `json:"configuration_id" api:"nullable"`
	CreatedAt       time.Time                  `json:"created_at" format:"date-time"`
	Domain          string                     `json:"domain"`
	Endpoint        string                     `json:"endpoint"`
	Method          string                     `json:"method"`
	Proxied         bool                       `json:"proxied"`
	ProxiedTo       string                     `json:"proxied_to" api:"nullable"`
	RequestID       string                     `json:"request_id" api:"nullable"`
	StatusCode      int64                      `json:"status_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKey          respjson.Field
		APIKeyID        respjson.Field
		ConfigurationID respjson.Field
		CreatedAt       respjson.Field
		Domain          respjson.Field
		Endpoint        respjson.Field
		Method          respjson.Field
		Proxied         respjson.Field
		ProxiedTo       respjson.Field
		RequestID       respjson.Field
		StatusCode      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3LogAPIResponseData) RawJSON() string { return r.JSON.raw }
func (r *V3LogAPIResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogAPIResponseDataAPIKey struct {
	ID     string `json:"id" api:"required"`
	Prefix string `json:"prefix" api:"required"`
	// Any of "active", "revoked".
	Status         string                                    `json:"status" api:"required"`
	Configurations []V3LogAPIResponseDataAPIKeyConfiguration `json:"configurations"`
	CreatedAt      time.Time                                 `json:"created_at" format:"date-time"`
	Permissions    []V3LogAPIResponseDataAPIKeyPermission    `json:"permissions"`
	RevokedAt      time.Time                                 `json:"revoked_at" api:"nullable" format:"date-time"`
	UpdatedAt      time.Time                                 `json:"updated_at" format:"date-time"`
	User           V3LogAPIResponseDataAPIKeyUser            `json:"user"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Prefix         respjson.Field
		Status         respjson.Field
		Configurations respjson.Field
		CreatedAt      respjson.Field
		Permissions    respjson.Field
		RevokedAt      respjson.Field
		UpdatedAt      respjson.Field
		User           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3LogAPIResponseDataAPIKey) RawJSON() string { return r.JSON.raw }
func (r *V3LogAPIResponseDataAPIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogAPIResponseDataAPIKeyConfiguration struct {
	ID     string                                         `json:"id" api:"required" format:"uuid"`
	Name   string                                         `json:"name" api:"required"`
	Type   string                                         `json:"type" api:"required"`
	Values []V3LogAPIResponseDataAPIKeyConfigurationValue `json:"values"`
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
func (r V3LogAPIResponseDataAPIKeyConfiguration) RawJSON() string { return r.JSON.raw }
func (r *V3LogAPIResponseDataAPIKeyConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogAPIResponseDataAPIKeyConfigurationValue struct {
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
func (r V3LogAPIResponseDataAPIKeyConfigurationValue) RawJSON() string { return r.JSON.raw }
func (r *V3LogAPIResponseDataAPIKeyConfigurationValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogAPIResponseDataAPIKeyPermission struct {
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3LogAPIResponseDataAPIKeyPermission) RawJSON() string { return r.JSON.raw }
func (r *V3LogAPIResponseDataAPIKeyPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogAPIResponseDataAPIKeyUser struct {
	ID               int64                                `json:"id" api:"required"`
	Email            string                               `json:"email" api:"required" format:"email"`
	Name             string                               `json:"name" api:"required"`
	CreatedAt        time.Time                            `json:"created_at" format:"date-time"`
	Roles            []V3LogAPIResponseDataAPIKeyUserRole `json:"roles"`
	TwoFactorEnabled bool                                 `json:"two_factor_enabled"`
	UpdatedAt        time.Time                            `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Email            respjson.Field
		Name             respjson.Field
		CreatedAt        respjson.Field
		Roles            respjson.Field
		TwoFactorEnabled respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3LogAPIResponseDataAPIKeyUser) RawJSON() string { return r.JSON.raw }
func (r *V3LogAPIResponseDataAPIKeyUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogAPIResponseDataAPIKeyUserRole struct {
	Name        string                                         `json:"name" api:"required"`
	Description string                                         `json:"description" api:"nullable"`
	Permissions []V3LogAPIResponseDataAPIKeyUserRolePermission `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Description respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3LogAPIResponseDataAPIKeyUserRole) RawJSON() string { return r.JSON.raw }
func (r *V3LogAPIResponseDataAPIKeyUserRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogAPIResponseDataAPIKeyUserRolePermission struct {
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3LogAPIResponseDataAPIKeyUserRolePermission) RawJSON() string { return r.JSON.raw }
func (r *V3LogAPIResponseDataAPIKeyUserRolePermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogAPIResponseLink struct {
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
func (r V3LogAPIResponseLink) RawJSON() string { return r.JSON.raw }
func (r *V3LogAPIResponseLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogQanapiFlowResponse struct {
	CurrentPage  int64                         `json:"current_page"`
	Data         []V3LogQanapiFlowResponseData `json:"data"`
	FirstPageURL string                        `json:"first_page_url" format:"uri"`
	From         int64                         `json:"from" api:"nullable"`
	LastPage     int64                         `json:"last_page"`
	LastPageURL  string                        `json:"last_page_url" format:"uri"`
	Links        []V3LogQanapiFlowResponseLink `json:"links"`
	NextPageURL  string                        `json:"next_page_url" api:"nullable" format:"uri"`
	Path         string                        `json:"path"`
	PerPage      int64                         `json:"per_page"`
	PrevPageURL  string                        `json:"prev_page_url" api:"nullable" format:"uri"`
	To           int64                         `json:"to" api:"nullable"`
	Total        int64                         `json:"total"`
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
func (r V3LogQanapiFlowResponse) RawJSON() string { return r.JSON.raw }
func (r *V3LogQanapiFlowResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogQanapiFlowResponseData struct {
	Action          string    `json:"action"`
	ConfigurationID int64     `json:"configuration_id" api:"nullable"`
	CreatedAt       time.Time `json:"created_at" format:"date-time"`
	Email           string    `json:"email" api:"nullable" format:"email"`
	RequestID       string    `json:"request_id" api:"nullable"`
	Type            string    `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action          respjson.Field
		ConfigurationID respjson.Field
		CreatedAt       respjson.Field
		Email           respjson.Field
		RequestID       respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3LogQanapiFlowResponseData) RawJSON() string { return r.JSON.raw }
func (r *V3LogQanapiFlowResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogQanapiFlowResponseLink struct {
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
func (r V3LogQanapiFlowResponseLink) RawJSON() string { return r.JSON.raw }
func (r *V3LogQanapiFlowResponseLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogUnifiedResponse struct {
	CurrentPage  int64                      `json:"current_page"`
	Data         []V3LogUnifiedResponseData `json:"data"`
	FirstPageURL string                     `json:"first_page_url" format:"uri"`
	From         int64                      `json:"from" api:"nullable"`
	LastPage     int64                      `json:"last_page"`
	LastPageURL  string                     `json:"last_page_url" format:"uri"`
	Links        []V3LogUnifiedResponseLink `json:"links"`
	NextPageURL  string                     `json:"next_page_url" api:"nullable" format:"uri"`
	Path         string                     `json:"path"`
	PerPage      int64                      `json:"per_page"`
	PrevPageURL  string                     `json:"prev_page_url" api:"nullable" format:"uri"`
	To           int64                      `json:"to" api:"nullable"`
	Total        int64                      `json:"total"`
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
func (r V3LogUnifiedResponse) RawJSON() string { return r.JSON.raw }
func (r *V3LogUnifiedResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogUnifiedResponseData struct {
	Action        string                                `json:"action"`
	CauserEmail   string                                `json:"causer_email" api:"nullable" format:"email"`
	Configuration V3LogUnifiedResponseDataConfiguration `json:"configuration"`
	Description   string                                `json:"description"`
	Details       any                                   `json:"details" api:"nullable"`
	FullLog       V3LogUnifiedResponseDataFullLogUnion  `json:"full_log"`
	// Any of "activity", "api", "usage".
	LogType    string                       `json:"log_type"`
	RequestID  string                       `json:"request_id" api:"nullable"`
	StatusCode int64                        `json:"status_code" api:"nullable"`
	Timestamp  time.Time                    `json:"timestamp" format:"date-time"`
	User       V3LogUnifiedResponseDataUser `json:"user"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action        respjson.Field
		CauserEmail   respjson.Field
		Configuration respjson.Field
		Description   respjson.Field
		Details       respjson.Field
		FullLog       respjson.Field
		LogType       respjson.Field
		RequestID     respjson.Field
		StatusCode    respjson.Field
		Timestamp     respjson.Field
		User          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3LogUnifiedResponseData) RawJSON() string { return r.JSON.raw }
func (r *V3LogUnifiedResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogUnifiedResponseDataConfiguration struct {
	ID     string                                       `json:"id" api:"required" format:"uuid"`
	Name   string                                       `json:"name" api:"required"`
	Type   string                                       `json:"type" api:"required"`
	Values []V3LogUnifiedResponseDataConfigurationValue `json:"values"`
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
func (r V3LogUnifiedResponseDataConfiguration) RawJSON() string { return r.JSON.raw }
func (r *V3LogUnifiedResponseDataConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogUnifiedResponseDataConfigurationValue struct {
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
func (r V3LogUnifiedResponseDataConfigurationValue) RawJSON() string { return r.JSON.raw }
func (r *V3LogUnifiedResponseDataConfigurationValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V3LogUnifiedResponseDataFullLogUnion contains all possible properties and values
// from [V3LogUnifiedResponseDataFullLogActivityLog],
// [V3LogUnifiedResponseDataFullLogAPILog],
// [V3LogUnifiedResponseDataFullLogQanapiFlowLog].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type V3LogUnifiedResponseDataFullLogUnion struct {
	Action string `json:"action"`
	// This field is from variant [V3LogUnifiedResponseDataFullLogActivityLog].
	Description string `json:"description"`
	// This field is from variant [V3LogUnifiedResponseDataFullLogActivityLog].
	IP string `json:"ip"`
	// This field is from variant [V3LogUnifiedResponseDataFullLogActivityLog].
	Timestamp time.Time `json:"timestamp"`
	// This field is from variant [V3LogUnifiedResponseDataFullLogActivityLog].
	User V3LogUnifiedResponseDataFullLogActivityLogUser `json:"user"`
	// This field is from variant [V3LogUnifiedResponseDataFullLogActivityLog].
	When string `json:"when"`
	// This field is from variant [V3LogUnifiedResponseDataFullLogAPILog].
	APIKey V3LogUnifiedResponseDataFullLogAPILogAPIKey `json:"api_key"`
	// This field is from variant [V3LogUnifiedResponseDataFullLogAPILog].
	APIKeyID        int64     `json:"api_key_id"`
	ConfigurationID int64     `json:"configuration_id"`
	CreatedAt       time.Time `json:"created_at"`
	// This field is from variant [V3LogUnifiedResponseDataFullLogAPILog].
	Domain string `json:"domain"`
	// This field is from variant [V3LogUnifiedResponseDataFullLogAPILog].
	Endpoint string `json:"endpoint"`
	// This field is from variant [V3LogUnifiedResponseDataFullLogAPILog].
	Method string `json:"method"`
	// This field is from variant [V3LogUnifiedResponseDataFullLogAPILog].
	Proxied bool `json:"proxied"`
	// This field is from variant [V3LogUnifiedResponseDataFullLogAPILog].
	ProxiedTo string `json:"proxied_to"`
	RequestID string `json:"request_id"`
	// This field is from variant [V3LogUnifiedResponseDataFullLogAPILog].
	StatusCode int64 `json:"status_code"`
	// This field is from variant [V3LogUnifiedResponseDataFullLogQanapiFlowLog].
	Email string `json:"email"`
	// This field is from variant [V3LogUnifiedResponseDataFullLogQanapiFlowLog].
	Type string `json:"type"`
	JSON struct {
		Action          respjson.Field
		Description     respjson.Field
		IP              respjson.Field
		Timestamp       respjson.Field
		User            respjson.Field
		When            respjson.Field
		APIKey          respjson.Field
		APIKeyID        respjson.Field
		ConfigurationID respjson.Field
		CreatedAt       respjson.Field
		Domain          respjson.Field
		Endpoint        respjson.Field
		Method          respjson.Field
		Proxied         respjson.Field
		ProxiedTo       respjson.Field
		RequestID       respjson.Field
		StatusCode      respjson.Field
		Email           respjson.Field
		Type            respjson.Field
		raw             string
	} `json:"-"`
}

func (u V3LogUnifiedResponseDataFullLogUnion) AsV3LogUnifiedResponseDataFullLogActivityLog() (v V3LogUnifiedResponseDataFullLogActivityLog) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V3LogUnifiedResponseDataFullLogUnion) AsV3LogUnifiedResponseDataFullLogAPILog() (v V3LogUnifiedResponseDataFullLogAPILog) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V3LogUnifiedResponseDataFullLogUnion) AsV3LogUnifiedResponseDataFullLogQanapiFlowLog() (v V3LogUnifiedResponseDataFullLogQanapiFlowLog) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V3LogUnifiedResponseDataFullLogUnion) RawJSON() string { return u.JSON.raw }

func (r *V3LogUnifiedResponseDataFullLogUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogUnifiedResponseDataFullLogActivityLog struct {
	Action      string                                         `json:"action"`
	Description string                                         `json:"description"`
	IP          string                                         `json:"ip" api:"nullable"`
	Timestamp   time.Time                                      `json:"timestamp" format:"date-time"`
	User        V3LogUnifiedResponseDataFullLogActivityLogUser `json:"user"`
	When        string                                         `json:"when"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Description respjson.Field
		IP          respjson.Field
		Timestamp   respjson.Field
		User        respjson.Field
		When        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3LogUnifiedResponseDataFullLogActivityLog) RawJSON() string { return r.JSON.raw }
func (r *V3LogUnifiedResponseDataFullLogActivityLog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogUnifiedResponseDataFullLogActivityLogUser struct {
	ID               int64                                                `json:"id" api:"required"`
	Email            string                                               `json:"email" api:"required" format:"email"`
	Name             string                                               `json:"name" api:"required"`
	CreatedAt        time.Time                                            `json:"created_at" format:"date-time"`
	Roles            []V3LogUnifiedResponseDataFullLogActivityLogUserRole `json:"roles"`
	TwoFactorEnabled bool                                                 `json:"two_factor_enabled"`
	UpdatedAt        time.Time                                            `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Email            respjson.Field
		Name             respjson.Field
		CreatedAt        respjson.Field
		Roles            respjson.Field
		TwoFactorEnabled respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3LogUnifiedResponseDataFullLogActivityLogUser) RawJSON() string { return r.JSON.raw }
func (r *V3LogUnifiedResponseDataFullLogActivityLogUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogUnifiedResponseDataFullLogActivityLogUserRole struct {
	Name        string                                                         `json:"name" api:"required"`
	Description string                                                         `json:"description" api:"nullable"`
	Permissions []V3LogUnifiedResponseDataFullLogActivityLogUserRolePermission `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Description respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3LogUnifiedResponseDataFullLogActivityLogUserRole) RawJSON() string { return r.JSON.raw }
func (r *V3LogUnifiedResponseDataFullLogActivityLogUserRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogUnifiedResponseDataFullLogActivityLogUserRolePermission struct {
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3LogUnifiedResponseDataFullLogActivityLogUserRolePermission) RawJSON() string {
	return r.JSON.raw
}
func (r *V3LogUnifiedResponseDataFullLogActivityLogUserRolePermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogUnifiedResponseDataFullLogAPILog struct {
	APIKey          V3LogUnifiedResponseDataFullLogAPILogAPIKey `json:"api_key"`
	APIKeyID        int64                                       `json:"api_key_id"`
	ConfigurationID int64                                       `json:"configuration_id" api:"nullable"`
	CreatedAt       time.Time                                   `json:"created_at" format:"date-time"`
	Domain          string                                      `json:"domain"`
	Endpoint        string                                      `json:"endpoint"`
	Method          string                                      `json:"method"`
	Proxied         bool                                        `json:"proxied"`
	ProxiedTo       string                                      `json:"proxied_to" api:"nullable"`
	RequestID       string                                      `json:"request_id" api:"nullable"`
	StatusCode      int64                                       `json:"status_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKey          respjson.Field
		APIKeyID        respjson.Field
		ConfigurationID respjson.Field
		CreatedAt       respjson.Field
		Domain          respjson.Field
		Endpoint        respjson.Field
		Method          respjson.Field
		Proxied         respjson.Field
		ProxiedTo       respjson.Field
		RequestID       respjson.Field
		StatusCode      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3LogUnifiedResponseDataFullLogAPILog) RawJSON() string { return r.JSON.raw }
func (r *V3LogUnifiedResponseDataFullLogAPILog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogUnifiedResponseDataFullLogAPILogAPIKey struct {
	ID     string `json:"id" api:"required"`
	Prefix string `json:"prefix" api:"required"`
	// Any of "active", "revoked".
	Status         string                                                     `json:"status" api:"required"`
	Configurations []V3LogUnifiedResponseDataFullLogAPILogAPIKeyConfiguration `json:"configurations"`
	CreatedAt      time.Time                                                  `json:"created_at" format:"date-time"`
	Permissions    []V3LogUnifiedResponseDataFullLogAPILogAPIKeyPermission    `json:"permissions"`
	RevokedAt      time.Time                                                  `json:"revoked_at" api:"nullable" format:"date-time"`
	UpdatedAt      time.Time                                                  `json:"updated_at" format:"date-time"`
	User           V3LogUnifiedResponseDataFullLogAPILogAPIKeyUser            `json:"user"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Prefix         respjson.Field
		Status         respjson.Field
		Configurations respjson.Field
		CreatedAt      respjson.Field
		Permissions    respjson.Field
		RevokedAt      respjson.Field
		UpdatedAt      respjson.Field
		User           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3LogUnifiedResponseDataFullLogAPILogAPIKey) RawJSON() string { return r.JSON.raw }
func (r *V3LogUnifiedResponseDataFullLogAPILogAPIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogUnifiedResponseDataFullLogAPILogAPIKeyConfiguration struct {
	ID     string                                                          `json:"id" api:"required" format:"uuid"`
	Name   string                                                          `json:"name" api:"required"`
	Type   string                                                          `json:"type" api:"required"`
	Values []V3LogUnifiedResponseDataFullLogAPILogAPIKeyConfigurationValue `json:"values"`
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
func (r V3LogUnifiedResponseDataFullLogAPILogAPIKeyConfiguration) RawJSON() string { return r.JSON.raw }
func (r *V3LogUnifiedResponseDataFullLogAPILogAPIKeyConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogUnifiedResponseDataFullLogAPILogAPIKeyConfigurationValue struct {
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
func (r V3LogUnifiedResponseDataFullLogAPILogAPIKeyConfigurationValue) RawJSON() string {
	return r.JSON.raw
}
func (r *V3LogUnifiedResponseDataFullLogAPILogAPIKeyConfigurationValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogUnifiedResponseDataFullLogAPILogAPIKeyPermission struct {
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3LogUnifiedResponseDataFullLogAPILogAPIKeyPermission) RawJSON() string { return r.JSON.raw }
func (r *V3LogUnifiedResponseDataFullLogAPILogAPIKeyPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogUnifiedResponseDataFullLogAPILogAPIKeyUser struct {
	ID               int64                                                 `json:"id" api:"required"`
	Email            string                                                `json:"email" api:"required" format:"email"`
	Name             string                                                `json:"name" api:"required"`
	CreatedAt        time.Time                                             `json:"created_at" format:"date-time"`
	Roles            []V3LogUnifiedResponseDataFullLogAPILogAPIKeyUserRole `json:"roles"`
	TwoFactorEnabled bool                                                  `json:"two_factor_enabled"`
	UpdatedAt        time.Time                                             `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Email            respjson.Field
		Name             respjson.Field
		CreatedAt        respjson.Field
		Roles            respjson.Field
		TwoFactorEnabled respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3LogUnifiedResponseDataFullLogAPILogAPIKeyUser) RawJSON() string { return r.JSON.raw }
func (r *V3LogUnifiedResponseDataFullLogAPILogAPIKeyUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogUnifiedResponseDataFullLogAPILogAPIKeyUserRole struct {
	Name        string                                                          `json:"name" api:"required"`
	Description string                                                          `json:"description" api:"nullable"`
	Permissions []V3LogUnifiedResponseDataFullLogAPILogAPIKeyUserRolePermission `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Description respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3LogUnifiedResponseDataFullLogAPILogAPIKeyUserRole) RawJSON() string { return r.JSON.raw }
func (r *V3LogUnifiedResponseDataFullLogAPILogAPIKeyUserRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogUnifiedResponseDataFullLogAPILogAPIKeyUserRolePermission struct {
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3LogUnifiedResponseDataFullLogAPILogAPIKeyUserRolePermission) RawJSON() string {
	return r.JSON.raw
}
func (r *V3LogUnifiedResponseDataFullLogAPILogAPIKeyUserRolePermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogUnifiedResponseDataFullLogQanapiFlowLog struct {
	Action          string    `json:"action"`
	ConfigurationID int64     `json:"configuration_id" api:"nullable"`
	CreatedAt       time.Time `json:"created_at" format:"date-time"`
	Email           string    `json:"email" api:"nullable" format:"email"`
	RequestID       string    `json:"request_id" api:"nullable"`
	Type            string    `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action          respjson.Field
		ConfigurationID respjson.Field
		CreatedAt       respjson.Field
		Email           respjson.Field
		RequestID       respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3LogUnifiedResponseDataFullLogQanapiFlowLog) RawJSON() string { return r.JSON.raw }
func (r *V3LogUnifiedResponseDataFullLogQanapiFlowLog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogUnifiedResponseDataUser struct {
	ID               int64                              `json:"id" api:"required"`
	Email            string                             `json:"email" api:"required" format:"email"`
	Name             string                             `json:"name" api:"required"`
	CreatedAt        time.Time                          `json:"created_at" format:"date-time"`
	Roles            []V3LogUnifiedResponseDataUserRole `json:"roles"`
	TwoFactorEnabled bool                               `json:"two_factor_enabled"`
	UpdatedAt        time.Time                          `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Email            respjson.Field
		Name             respjson.Field
		CreatedAt        respjson.Field
		Roles            respjson.Field
		TwoFactorEnabled respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3LogUnifiedResponseDataUser) RawJSON() string { return r.JSON.raw }
func (r *V3LogUnifiedResponseDataUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogUnifiedResponseDataUserRole struct {
	Name        string                                       `json:"name" api:"required"`
	Description string                                       `json:"description" api:"nullable"`
	Permissions []V3LogUnifiedResponseDataUserRolePermission `json:"permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Description respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3LogUnifiedResponseDataUserRole) RawJSON() string { return r.JSON.raw }
func (r *V3LogUnifiedResponseDataUserRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogUnifiedResponseDataUserRolePermission struct {
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V3LogUnifiedResponseDataUserRolePermission) RawJSON() string { return r.JSON.raw }
func (r *V3LogUnifiedResponseDataUserRolePermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogUnifiedResponseLink struct {
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
func (r V3LogUnifiedResponseLink) RawJSON() string { return r.JSON.raw }
func (r *V3LogUnifiedResponseLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V3LogActivityParams struct {
	LogName param.Opt[string] `query:"logName,omitzero" json:"-"`
	Page    param.Opt[int64]  `query:"page,omitzero" json:"-"`
	PerPage param.Opt[int64]  `query:"per_page,omitzero" json:"-"`
	// User ID filter
	User   param.Opt[int64] `query:"user,omitzero" json:"-"`
	UserID param.Opt[int64] `query:"user_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V3LogActivityParams]'s query parameters as `url.Values`.
func (r V3LogActivityParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V3LogAPIParams struct {
	// API Key ID filter
	APIKey  param.Opt[int64] `query:"apiKey,omitzero" json:"-"`
	Page    param.Opt[int64] `query:"page,omitzero" json:"-"`
	PerPage param.Opt[int64] `query:"per_page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V3LogAPIParams]'s query parameters as `url.Values`.
func (r V3LogAPIParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V3LogQanapiFlowParams struct {
	Page    param.Opt[int64] `query:"page,omitzero" json:"-"`
	PerPage param.Opt[int64] `query:"per_page,omitzero" json:"-"`
	// Integration type filter
	Type param.Opt[string] `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V3LogQanapiFlowParams]'s query parameters as `url.Values`.
func (r V3LogQanapiFlowParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V3LogUnifiedParams struct {
	Action      param.Opt[string] `query:"action,omitzero" json:"-"`
	CauserEmail param.Opt[string] `query:"causer_email,omitzero" json:"-"`
	Description param.Opt[string] `query:"description,omitzero" json:"-"`
	Details     param.Opt[string] `query:"details,omitzero" json:"-"`
	Page        param.Opt[int64]  `query:"page,omitzero" json:"-"`
	PerPage     param.Opt[int64]  `query:"per_page,omitzero" json:"-"`
	RequestID   param.Opt[string] `query:"request_id,omitzero" json:"-"`
	StatusCode  param.Opt[int64]  `query:"status_code,omitzero" json:"-"`
	UserID      param.Opt[int64]  `query:"user_id,omitzero" json:"-"`
	// Any of "activity", "api", "usage".
	LogType V3LogUnifiedParamsLogType `query:"log_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V3LogUnifiedParams]'s query parameters as `url.Values`.
func (r V3LogUnifiedParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V3LogUnifiedParamsLogType string

const (
	V3LogUnifiedParamsLogTypeActivity V3LogUnifiedParamsLogType = "activity"
	V3LogUnifiedParamsLogTypeAPI      V3LogUnifiedParamsLogType = "api"
	V3LogUnifiedParamsLogTypeUsage    V3LogUnifiedParamsLogType = "usage"
)
