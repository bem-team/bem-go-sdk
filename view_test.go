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

func TestViewNew(t *testing.T) {
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
		Aggregations: []bem.ViewNewParamsAggregation{{
			Function:            "count",
			Name:                "name",
			AggregateColumnName: bem.String("aggregateColumnName"),
			GroupByColumnName:   bem.String("groupByColumnName"),
		}},
		Columns: []bem.ViewNewParamsColumn{{
			DisplayOrderIndex: 0,
			Name:              "name",
			ValueSchemaPath:   []string{"string"},
		}},
		Filters: []bem.ViewNewParamsFilter{{
			ColumnName: "columnName",
			FilterType: "equals_string",
			Number:     bem.Float(0),
			String:     bem.String("string"),
		}},
		Functions: []bem.ViewNewParamsFunction{{
			ID:   bem.String("id"),
			Name: bem.String("name"),
		}},
		Name: "name",
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

func TestViewUpdate(t *testing.T) {
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
			Aggregations: []bem.ViewUpdateParamsAggregation{{
				Function:            "count",
				Name:                "name",
				AggregateColumnName: bem.String("aggregateColumnName"),
				GroupByColumnName:   bem.String("groupByColumnName"),
			}},
			Columns: []bem.ViewUpdateParamsColumn{{
				DisplayOrderIndex: 0,
				Name:              "name",
				ValueSchemaPath:   []string{"string"},
			}},
			Filters: []bem.ViewUpdateParamsFilter{{
				ColumnName: "columnName",
				FilterType: "equals_string",
				Number:     bem.Float(0),
				String:     bem.String("string"),
			}},
			Functions: []bem.ViewUpdateParamsFunction{{
				ID:   bem.String("id"),
				Name: bem.String("name"),
			}},
			Name: "name",
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

func TestViewGenerateAggregationData(t *testing.T) {
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
		Aggregations: []bem.ViewGenerateAggregationDataParamsAggregation{{
			Function:            "count",
			Name:                "name",
			AggregateColumnName: bem.String("aggregateColumnName"),
			GroupByColumnName:   bem.String("groupByColumnName"),
		}},
		Columns: []bem.ViewGenerateAggregationDataParamsColumn{{
			DisplayOrderIndex: 0,
			Name:              "name",
			ValueSchemaPath:   []string{"string"},
		}},
		Filters: []bem.ViewGenerateAggregationDataParamsFilter{{
			ColumnName: "columnName",
			FilterType: "equals_string",
			Number:     bem.Float(0),
			String:     bem.String("string"),
		}},
		Functions: []bem.ViewGenerateAggregationDataParamsFunction{{
			ID:   bem.String("id"),
			Name: bem.String("name"),
		}},
		Name: "name",
		TimeWindow: bem.ViewGenerateAggregationDataParamsTimeWindow{
			End:   time.Now(),
			Start: time.Now(),
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
		Aggregations: []bem.ViewGenerateTableDataParamsAggregation{{
			Function:            "count",
			Name:                "name",
			AggregateColumnName: bem.String("aggregateColumnName"),
			GroupByColumnName:   bem.String("groupByColumnName"),
		}},
		Columns: []bem.ViewGenerateTableDataParamsColumn{{
			DisplayOrderIndex: 0,
			Name:              "name",
			ValueSchemaPath:   []string{"string"},
		}},
		Filters: []bem.ViewGenerateTableDataParamsFilter{{
			ColumnName: "columnName",
			FilterType: "equals_string",
			Number:     bem.Float(0),
			String:     bem.String("string"),
		}},
		Functions: []bem.ViewGenerateTableDataParamsFunction{{
			ID:   bem.String("id"),
			Name: bem.String("name"),
		}},
		Name: "name",
		TimeWindow: bem.ViewGenerateTableDataParamsTimeWindow{
			End:   time.Now(),
			Start: time.Now(),
		},
		Limit:  bem.Int(1),
		Offset: bem.Int(0),
	})
	if err != nil {
		var apierr *bem.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
