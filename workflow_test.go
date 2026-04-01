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

func TestWorkflowNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.New(context.TODO(), bem.WorkflowNewParams{
		DisplayName: bem.String("displayName"),
		MainFunction: bem.FunctionVersionIdentifierParam{
			ID:         bem.String("id"),
			Name:       bem.String("name"),
			VersionNum: bem.Int(0),
		},
		Name: bem.String("name"),
		Relationships: []bem.WorkflowRequestRelationshipParam{{
			DestinationFunction: bem.FunctionVersionIdentifierParam{
				ID:         bem.String("id"),
				Name:       bem.String("name"),
				VersionNum: bem.Int(0),
			},
			SourceFunction: bem.FunctionVersionIdentifierParam{
				ID:         bem.String("id"),
				Name:       bem.String("name"),
				VersionNum: bem.Int(0),
			},
			DestinationName: bem.String("destinationName"),
		}},
		Tags: []string{"string"},
	})
	if err != nil {
		var apierr *bem.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowGet(t *testing.T) {
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
	_, err := client.Workflows.Get(context.TODO(), "workflowName")
	if err != nil {
		var apierr *bem.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Update(
		context.TODO(),
		"workflowName",
		bem.WorkflowUpdateParams{
			DisplayName: bem.String("displayName"),
			MainFunction: bem.FunctionVersionIdentifierParam{
				ID:         bem.String("id"),
				Name:       bem.String("name"),
				VersionNum: bem.Int(0),
			},
			Name: bem.String("name"),
			Relationships: []bem.WorkflowRequestRelationshipParam{{
				DestinationFunction: bem.FunctionVersionIdentifierParam{
					ID:         bem.String("id"),
					Name:       bem.String("name"),
					VersionNum: bem.Int(0),
				},
				SourceFunction: bem.FunctionVersionIdentifierParam{
					ID:         bem.String("id"),
					Name:       bem.String("name"),
					VersionNum: bem.Int(0),
				},
				DestinationName: bem.String("destinationName"),
			}},
			Tags: []string{"string"},
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

func TestWorkflowListWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.List(context.TODO(), bem.WorkflowListParams{
		DisplayName:   bem.String("displayName"),
		EndingBefore:  bem.String("endingBefore"),
		FunctionIDs:   []string{"string"},
		FunctionNames: []string{"string"},
		Limit:         bem.Int(1),
		SortOrder:     bem.WorkflowListParamsSortOrderAsc,
		StartingAfter: bem.String("startingAfter"),
		Tags:          []string{"string"},
		WorkflowIDs:   []string{"string"},
		WorkflowNames: []string{"string"},
	})
	if err != nil {
		var apierr *bem.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowDelete(t *testing.T) {
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
	err := client.Workflows.Delete(context.TODO(), "workflowName")
	if err != nil {
		var apierr *bem.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowCallWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Call(
		context.TODO(),
		"workflowName",
		bem.WorkflowCallParams{
			CallReferenceID: bem.String("callReferenceID"),
			File:            map[string]any{},
			Files:           []string{"string"},
			Wait:            bem.String("wait"),
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

func TestWorkflowCopyWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Copy(context.TODO(), bem.WorkflowCopyParams{
		SourceWorkflowName:       "sourceWorkflowName",
		TargetWorkflowName:       "targetWorkflowName",
		SourceWorkflowVersionNum: bem.Int(1),
		Tags:                     []string{"string"},
		TargetDisplayName:        bem.String("targetDisplayName"),
		TargetEnvironment:        bem.String("targetEnvironment"),
	})
	if err != nil {
		var apierr *bem.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
