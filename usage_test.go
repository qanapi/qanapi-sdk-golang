// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package qanapi_test

import (
	"context"
	"os"
	"testing"

	"github.com/qanapi/qanapi-sdk-golang"
	"github.com/qanapi/qanapi-sdk-golang/internal/testutil"
	"github.com/qanapi/qanapi-sdk-golang/option"
)

func TestUsage(t *testing.T) {
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
	response, err := client.V3.Encryption.Encrypt(
		context.TODO(),
		"proxy",
		qanapi.V3EncryptionEncryptParams{
			Body: map[string]any{
				"name":    "bar",
				"email":   "bar",
				"ssn":     "bar",
				"dob":     "bar",
				"address": "bar",
			},
			XQanapiFields: "x-qanapi-fields",
		},
	)
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response)
}
