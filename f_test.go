// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bem_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/bem-team/bem-go-sdk"
	"github.com/bem-team/bem-go-sdk/internal/testutil"
	"github.com/bem-team/bem-go-sdk/option"
)

func TestFNavigateWithOptionalParams(t *testing.T) {
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
	_, err := client.Fs.Navigate(context.TODO(), bem.FNavigateParams{
		Op:        bem.FNavigateParamsOpLs,
		CountOnly: bem.Bool(true),
		Cursor:    bem.String("cursor"),
		Filter: bem.FNavigateParamsFilter{
			FunctionName: bem.String("functionName"),
			Search:       bem.String("search"),
			Since:        bem.Time(time.Now()),
			Type:         bem.String("type"),
		},
		IgnoreCase: bem.Bool(true),
		Limit:      bem.Int(0),
		N:          bem.Int(0),
		Path:       bem.String("path"),
		Pattern:    bem.String("pattern"),
		Range: bem.FNavigateParamsRange{
			Page:         bem.Int(0),
			PageRange:    []int64{0, 0},
			SectionTypes: []string{"string"},
		},
		Regex:  bem.Bool(true),
		Scope:  bem.String("scope"),
		Select: []string{"string"},
	})
	if err != nil {
		var apierr *bem.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
