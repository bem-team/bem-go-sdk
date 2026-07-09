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

func TestViewNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Views.New(context.TODO(), bem.ViewNewParams{
		ViewCreate: bem.ViewCreateParam{
			Aggregations: []bem.ViewAggregationParam{{
				Function:            bem.ViewAggregationFunctionCount,
				Name:                "name",
				AggregateColumnName: bem.String("aggregateColumnName"),
				DisplayType:         bem.ViewAggregationDisplayTypeTable,
				GroupByColumnName:   bem.String("groupByColumnName"),
			}},
			Columns: []bem.ViewColumnParam{{
				DisplayOrderIndex: 0,
				Name:              "name",
				ValueSchemaPath:   []string{"string"},
			}},
			Filters: []bem.ViewFilterParam{{
				ColumnName: "columnName",
				FilterType: bem.ViewFilterFilterTypeEqualsString,
				Number:     bem.Float(0),
				String:     bem.String("string"),
			}},
			Functions: []bem.FunctionIdentifierParam{{
				ID:   bem.String("id"),
				Name: bem.String("name"),
			}},
			Name:        "name",
			Description: bem.String("description"),
		},
	})
	if err != nil {
		var apierr *bem.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestViewGet(t *testing.T) {
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
	_, err := client.Views.Get(context.TODO(), "view_id")
	if err != nil {
		var apierr *bem.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestViewUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Views.Update(
		context.TODO(),
		"view_id",
		bem.ViewUpdateParams{
			ViewCreate: bem.ViewCreateParam{
				Aggregations: []bem.ViewAggregationParam{{
					Function:            bem.ViewAggregationFunctionCount,
					Name:                "name",
					AggregateColumnName: bem.String("aggregateColumnName"),
					DisplayType:         bem.ViewAggregationDisplayTypeTable,
					GroupByColumnName:   bem.String("groupByColumnName"),
				}},
				Columns: []bem.ViewColumnParam{{
					DisplayOrderIndex: 0,
					Name:              "name",
					ValueSchemaPath:   []string{"string"},
				}},
				Filters: []bem.ViewFilterParam{{
					ColumnName: "columnName",
					FilterType: bem.ViewFilterFilterTypeEqualsString,
					Number:     bem.Float(0),
					String:     bem.String("string"),
				}},
				Functions: []bem.FunctionIdentifierParam{{
					ID:   bem.String("id"),
					Name: bem.String("name"),
				}},
				Name:        "name",
				Description: bem.String("description"),
			},
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

func TestViewListWithOptionalParams(t *testing.T) {
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
	_, err := client.Views.List(context.TODO(), bem.ViewListParams{
		EndingBefore:      bem.String("endingBefore"),
		FunctionIDs:       []string{"string"},
		FunctionNames:     []string{"string"},
		Limit:             bem.Int(1),
		SortOrder:         bem.ViewListParamsSortOrderAsc,
		StartingAfter:     bem.String("startingAfter"),
		ViewIDs:           []string{"string"},
		ViewNameSubstring: bem.String("viewNameSubstring"),
	})
	if err != nil {
		var apierr *bem.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestViewDelete(t *testing.T) {
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
	err := client.Views.Delete(context.TODO(), "view_id")
	if err != nil {
		var apierr *bem.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestViewGenerateAggregationDataWithOptionalParams(t *testing.T) {
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
	_, err := client.Views.GenerateAggregationData(context.TODO(), bem.ViewGenerateAggregationDataParams{
		Aggregations: []bem.ViewAggregationParam{{
			Function:            bem.ViewAggregationFunctionCount,
			Name:                "name",
			AggregateColumnName: bem.String("aggregateColumnName"),
			DisplayType:         bem.ViewAggregationDisplayTypeTable,
			GroupByColumnName:   bem.String("groupByColumnName"),
		}},
		Columns: []bem.ViewColumnParam{{
			DisplayOrderIndex: 0,
			Name:              "name",
			ValueSchemaPath:   []string{"string"},
		}},
		Filters: []bem.ViewFilterParam{{
			ColumnName: "columnName",
			FilterType: bem.ViewFilterFilterTypeEqualsString,
			Number:     bem.Float(0),
			String:     bem.String("string"),
		}},
		Functions: []bem.FunctionIdentifierParam{{
			ID:   bem.String("id"),
			Name: bem.String("name"),
		}},
		Name: "name",
		TimeWindow: bem.TimeWindowParam{
			End:   time.Now(),
			Start: time.Now(),
		},
		Description: bem.String("description"),
	})
	if err != nil {
		var apierr *bem.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestViewGenerateTableDataWithOptionalParams(t *testing.T) {
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
	_, err := client.Views.GenerateTableData(context.TODO(), bem.ViewGenerateTableDataParams{
		Aggregations: []bem.ViewAggregationParam{{
			Function:            bem.ViewAggregationFunctionCount,
			Name:                "name",
			AggregateColumnName: bem.String("aggregateColumnName"),
			DisplayType:         bem.ViewAggregationDisplayTypeTable,
			GroupByColumnName:   bem.String("groupByColumnName"),
		}},
		Columns: []bem.ViewColumnParam{{
			DisplayOrderIndex: 0,
			Name:              "name",
			ValueSchemaPath:   []string{"string"},
		}},
		Filters: []bem.ViewFilterParam{{
			ColumnName: "columnName",
			FilterType: bem.ViewFilterFilterTypeEqualsString,
			Number:     bem.Float(0),
			String:     bem.String("string"),
		}},
		Functions: []bem.FunctionIdentifierParam{{
			ID:   bem.String("id"),
			Name: bem.String("name"),
		}},
		Name: "name",
		TimeWindow: bem.TimeWindowParam{
			End:   time.Now(),
			Start: time.Now(),
		},
		Description: bem.String("description"),
		Limit:       bem.Int(1),
		Offset:      bem.Int(0),
	})
	if err != nil {
		var apierr *bem.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
