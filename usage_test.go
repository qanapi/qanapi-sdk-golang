// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package qanapiqanapisdkgolang_test

import (
	"context"
	"os"
	"testing"

	"github.com/qanapi/qanapi-sdk-golang"
	"github.com/qanapi/qanapi-sdk-golang/internal/testutil"
	"github.com/qanapi/qanapi-sdk-golang/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := qanapiqanapisdkgolang.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithSubdomain("My-Subdomain"),
	)
	response, err := client.Auth.Login(context.TODO(), qanapiqanapisdkgolang.AuthLoginParams{
		Email:    "valid@email.com",
		Password: "secret123",
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response.AccessToken)
}
