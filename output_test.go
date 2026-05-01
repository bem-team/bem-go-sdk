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

func TestOutputGet(t *testing.T) {
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
	_, err := client.Outputs.Get(context.TODO(), "eventID")
	if err != nil {
		var apierr *bem.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOutputListWithOptionalParams(t *testing.T) {
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
	_, err := client.Outputs.List(context.TODO(), bem.OutputListParams{
		CallIDs:              []string{"string"},
		EndingBefore:         bem.String("endingBefore"),
		EventIDs:             []string{"string"},
		EventTypes:           []string{"string"},
		FunctionIDs:          []string{"string"},
		FunctionNames:        []string{"string"},
		FunctionVersionNums:  []int64{0},
		IncludeIntermediate:  bem.Bool(true),
		IsLabelled:           bem.Bool(true),
		IsRegression:         bem.Bool(true),
		Limit:                bem.Int(1),
		ReferenceIDs:         []string{"string"},
		ReferenceIDSubstring: bem.String("referenceIDSubstring"),
		SortOrder:            bem.OutputListParamsSortOrderAsc,
		StartingAfter:        bem.String("startingAfter"),
		TransformationIDs:    []string{"string"},
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
