// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bem_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/bem-go"
	"github.com/stainless-sdks/bem-go/internal/testutil"
	"github.com/stainless-sdks/bem-go/option"
)

func TestErrorGet(t *testing.T) {
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
	_, err := client.Errors.Get(context.TODO(), "eventID")
	if err != nil {
		var apierr *bem.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestErrorListWithOptionalParams(t *testing.T) {
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
	_, err := client.Errors.List(context.TODO(), bem.ErrorListParams{
		CallIDs:              []string{"string"},
		EndingBefore:         bem.String("endingBefore"),
		FunctionIDs:          []string{"string"},
		FunctionNames:        []string{"string"},
		Limit:                bem.Int(1),
		ReferenceIDs:         []string{"string"},
		ReferenceIDSubstring: bem.String("referenceIDSubstring"),
		SortOrder:            bem.ErrorListParamsSortOrderAsc,
		StartingAfter:        bem.String("startingAfter"),
		WorkflowIDs:          []string{"string"},
		WorkflowNames:        []string{"string"},
	})
	if err != nil {
		var apierr *bem.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
