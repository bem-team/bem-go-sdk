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
		MainNodeName: "mainNodeName",
		Name:         "name",
		Nodes: []bem.WorkflowNewParamsNode{{
			Function: bem.FunctionVersionIdentifierParam{
				ID:         bem.String("id"),
				Name:       bem.String("name"),
				VersionNum: bem.Int(0),
			},
			Metadata: map[string]any{},
			Name:     bem.String("name"),
		}},
		Connectors: []bem.WorkflowNewParamsConnector{{
			Name:        "name",
			Type:        "paragon",
			ConnectorID: bem.String("connectorID"),
			Paragon: bem.WorkflowNewParamsConnectorParagon{
				Configuration: map[string]any{},
				Integration:   bem.String("integration"),
			},
		}},
		DisplayName: bem.String("displayName"),
		Edges: []bem.WorkflowNewParamsEdge{{
			DestinationNodeName: "destinationNodeName",
			SourceNodeName:      "sourceNodeName",
			DestinationName:     bem.String("destinationName"),
			Metadata:            map[string]any{},
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
			Connectors: []bem.WorkflowUpdateParamsConnector{{
				Name:        "name",
				Type:        "paragon",
				ConnectorID: bem.String("connectorID"),
				Paragon: bem.WorkflowUpdateParamsConnectorParagon{
					Configuration: map[string]any{},
					Integration:   bem.String("integration"),
				},
			}},
			DisplayName: bem.String("displayName"),
			Edges: []bem.WorkflowUpdateParamsEdge{{
				DestinationNodeName: "destinationNodeName",
				SourceNodeName:      "sourceNodeName",
				DestinationName:     bem.String("destinationName"),
				Metadata:            map[string]any{},
			}},
			MainNodeName: bem.String("mainNodeName"),
			Name:         bem.String("name"),
			Nodes: []bem.WorkflowUpdateParamsNode{{
				Function: bem.FunctionVersionIdentifierParam{
					ID:         bem.String("id"),
					Name:       bem.String("name"),
					VersionNum: bem.Int(0),
				},
				Metadata: map[string]any{},
				Name:     bem.String("name"),
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
			Input: bem.WorkflowCallParamsInput{
				BatchFiles: bem.WorkflowCallParamsInputBatchFiles{
					Inputs: []bem.WorkflowCallParamsInputBatchFilesInput{{
						InputContent:    "inputContent",
						InputType:       "csv",
						ItemReferenceID: bem.String("itemReferenceID"),
					}},
				},
				SingleFile: bem.WorkflowCallParamsInputSingleFile{
					InputContent: "inputContent",
					InputType:    "csv",
				},
			},
			Wait:            bem.Bool(true),
			CallReferenceID: bem.String("callReferenceID"),
			Metadata:        map[string]any{},
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
