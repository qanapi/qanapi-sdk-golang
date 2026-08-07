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

func TestV2EncryptEncryptDataWithOptionalParams(t *testing.T) {
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
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.V2.Encrypt.EncryptData(context.TODO(), qanapi.V2EncryptEncryptDataParams{
		Data: qanapi.V2EncryptEncryptDataParamsDataUnion{
			OfAnyMap: map[string]any{
				"password": "bar",
			},
		},
		Access: qanapi.V2EncryptEncryptDataParamsAccess{
			ACL: []string{"admin"},
		},
		Attributes: qanapi.V2EncryptEncryptDataParamsAttributes{
			Classification: "confidential",
			Owner:          qanapi.String("alice@example.com"),
			Tags:           []string{"legal"},
		},
		SensitiveFields: []string{"password"},
	})
	if err != nil {
		var apierr *qanapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
