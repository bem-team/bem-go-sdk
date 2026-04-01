// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bem_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/bem-go"
	"github.com/stainless-sdks/bem-go/internal/testutil"
	"github.com/stainless-sdks/bem-go/option"
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
	client := bem.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	functionResponse, err := client.Functions.New(context.TODO(), bem.FunctionNewParamsCreateTransformFunction{
		FunctionName: "functionName",
		Type:         bem.FunctionNewParamsCreateTransformFunctionTypeTransform,
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", functionResponse.Function)
}
