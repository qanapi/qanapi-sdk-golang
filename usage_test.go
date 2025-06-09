// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package qanapi_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/qanapi-go"
	"github.com/stainless-sdks/qanapi-go/internal/testutil"
	"github.com/stainless-sdks/qanapi-go/option"
)

func TestUsage(t *testing.T) {
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
	response, err := client.Auth.Login(context.TODO(), qanapi.AuthLoginParams{
		Email:    "valid@email.com",
		Password: "secret123",
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response.AccessToken)
}
