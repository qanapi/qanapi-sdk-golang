// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package qanapi_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/qanapi/qanapi-sdk-golang"
	"github.com/qanapi/qanapi-sdk-golang/internal/testutil"
	"github.com/qanapi/qanapi-sdk-golang/option"
)

func TestV3ClassificationNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := qanapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithSubdomain("My-Subdomain"),
	)
	_, err := client.V3.Classifications.New(context.TODO(), qanapi.V3ClassificationNewParams{
		BgColor:     "#e1cb97",
		FgColor:     "#e1cb97",
		Name:        "name",
		Description: qanapi.String("description"),
		Emoji:       qanapi.String("emoji"),
		GwsGroups: []qanapi.GoogleGroupParam{{
			ID:    "id",
			Email: qanapi.String("email"),
			Name:  qanapi.String("name"),
		}},
		ProviderContainerID: qanapi.Int(0),
		Roles:               []int64{0},
		Users:               []int64{0},
	})
	if err != nil {
		var apierr *qanapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV3ClassificationUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := qanapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithSubdomain("My-Subdomain"),
	)
	_, err := client.V3.Classifications.Update(
		context.TODO(),
		0,
		qanapi.V3ClassificationUpdateParams{
			BgColor:     "#e1cb97",
			FgColor:     "#e1cb97",
			Name:        "name",
			Description: qanapi.String("description"),
			Emoji:       qanapi.String("emoji"),
			GwsGroups: []qanapi.GoogleGroupParam{{
				ID:    "id",
				Email: qanapi.String("email"),
				Name:  qanapi.String("name"),
			}},
			ProviderContainerID: qanapi.Int(0),
			Roles:               []int64{0},
			Users:               []int64{0},
		},
	)
	if err != nil {
		var apierr *qanapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV3ClassificationListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := qanapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithSubdomain("My-Subdomain"),
	)
	_, err := client.V3.Classifications.List(context.TODO(), qanapi.V3ClassificationListParams{
		Direction: qanapi.V3ClassificationListParamsDirectionAsc,
		PerPage:   qanapi.Int(0),
		Providers: []int64{0},
		Roles:     []int64{0},
		Search:    qanapi.String("search"),
		Sort:      qanapi.String("sort"),
		User:      qanapi.Int(0),
	})
	if err != nil {
		var apierr *qanapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV3ClassificationDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := qanapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithSubdomain("My-Subdomain"),
	)
	err := client.V3.Classifications.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *qanapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV3ClassificationShow(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := qanapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithSubdomain("My-Subdomain"),
	)
	_, err := client.V3.Classifications.Show(context.TODO(), 0)
	if err != nil {
		var apierr *qanapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
