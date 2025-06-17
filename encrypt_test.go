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

func TestEncryptEncryptDataWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Encrypt.EncryptData(context.TODO(), qanapi.EncryptEncryptDataParams{
		Encrypt: qanapi.EncryptParam{
			Data: qanapi.EncryptDataUnionParam{
				OfAnyMap: map[string]any{
					"password": "bar",
				},
			},
			Access: qanapi.EncryptAccessParam{
				ACL: []string{"admin"},
			},
			Attributes: qanapi.EncryptAttributesParam{
				Classification: "confidential",
				Owner:          qanapi.String("alice@example.com"),
				Tags:           []string{"legal"},
			},
			SensitiveFields: []string{"password"},
		},
	})
	if err != nil {
		var apierr *qanapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
