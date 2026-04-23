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

func TestCollectionItemGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Collections.Items.Get(context.TODO(), bem.CollectionItemGetParams{
		CollectionName:        "collectionName",
		IncludeSubcollections: bem.Bool(true),
		Limit:                 bem.Int(1),
		Page:                  bem.Int(1),
	})
	if err != nil {
		var apierr *bem.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCollectionItemUpdate(t *testing.T) {
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
	_, err := client.Collections.Items.Update(context.TODO(), bem.CollectionItemUpdateParams{
		CollectionName: "product_catalog",
		Items: []bem.CollectionItemUpdateParamsItem{{
			CollectionItemID: "clitm_2N6gH8ZKCmvb6BnFcGqhKJ98VzP",
			Data:             "SKU-12345: Updated Industrial Widget - Premium Edition",
		}, {
			CollectionItemID: "clitm_3M7hI9ALDnwc7CoGdHriLK09WaQ",
			Data: map[string]any{
				"sku":      "SKU-67890",
				"name":     "Updated Premium Gear",
				"category": "Hardware",
				"price":    399.99,
			},
		}},
	})
	if err != nil {
		var apierr *bem.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCollectionItemDelete(t *testing.T) {
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
	err := client.Collections.Items.Delete(context.TODO(), bem.CollectionItemDeleteParams{
		CollectionItemID: "collectionItemID",
		CollectionName:   "collectionName",
	})
	if err != nil {
		var apierr *bem.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCollectionItemAdd(t *testing.T) {
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
	_, err := client.Collections.Items.Add(context.TODO(), bem.CollectionItemAddParams{
		CollectionName: "product_catalog",
		Items: []bem.CollectionItemAddParamsItem{{
			Data: map[string]any{
				"sku":      "SKU-11111",
				"name":     "Deluxe Component",
				"category": "Hardware",
				"price":    299.99,
			},
		}, {
			Data: map[string]any{
				"sku":      "SKU-22222",
				"name":     "Standard Part",
				"category": "Tools",
				"price":    49.99,
			},
		}},
	})
	if err != nil {
		var apierr *bem.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
