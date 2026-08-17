// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bem_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/bem-team/bem-go-sdk"
	"github.com/bem-team/bem-go-sdk/internal/testutil"
	"github.com/bem-team/bem-go-sdk/option"
)

func TestEntityTypeNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := bem.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.EntityTypes.New(context.TODO(), bem.EntityTypeNewParams{
		Name:            "Drug",
		AttributeSchema: map[string]any{},
		Description:     bem.String("A pharmaceutical compound"),
		ParentTypeID:    bem.String("parentTypeID"),
	})
	if err != nil {
		var apierr *bem.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityTypeGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := bem.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.EntityTypes.Get(context.TODO(), "typeID")
	if err != nil {
		var apierr *bem.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityTypeUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := bem.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.EntityTypes.Update(
		context.TODO(),
		"typeID",
		bem.EntityTypeUpdateParams{
			AttributeSchema: map[string]any{},
			Description:     bem.String("description"),
			ParentTypeID:    bem.String("parentTypeID"),
		},
	)
	if err != nil {
		var apierr *bem.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityTypeDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := bem.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.EntityTypes.Delete(context.TODO(), "typeID")
	if err != nil {
		var apierr *bem.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
