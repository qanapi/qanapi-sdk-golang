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

func TestV3APIKeyList(t *testing.T) {
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
	_, err := client.V3.APIKeys.List(context.TODO())
	if err != nil {
		var apierr *qanapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV3APIKeyRevoke(t *testing.T) {
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
	err := client.V3.APIKeys.Revoke(context.TODO(), 0)
	if err != nil {
		var apierr *qanapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV3APIKeyRotate(t *testing.T) {
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
	_, err := client.V3.APIKeys.Rotate(context.TODO(), 0)
	if err != nil {
		var apierr *qanapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV3APIKeyShow(t *testing.T) {
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
	_, err := client.V3.APIKeys.Show(context.TODO(), 0)
	if err != nil {
		var apierr *qanapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
