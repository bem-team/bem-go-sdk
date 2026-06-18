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

func TestEntityUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Entities.Update(
		context.TODO(),
		"id",
		bem.EntityUpdateParams{
			AddSynonyms:      []string{"string"},
			AssignedTypeID:   bem.String("assignedTypeID"),
			Canonical:        bem.String("canonical"),
			Locale:           bem.String("locale"),
			RemoveSynonymIDs: []string{"string"},
			Status:           bem.EntityUpdateParamsStatusApproved,
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

func TestEntityBulkNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Entities.BulkNew(context.TODO(), bem.EntityBulkNewParams{
		Entities: []bem.EntityBulkNewParamsEntity{{
			Canonical: "Acme Corporation",
			Type:      "organization",
			Attributes: map[string]any{
				"headquarters": "Springfield",
			},
			Description: bem.String("Industrial conglomerate"),
			Synonyms:    []string{"ACME", "Acme Corp"},
		}},
		Bucket:     bem.String("bucket"),
		OnConflict: bem.EntityBulkNewParamsOnConflictMerge,
	})
	if err != nil {
		var apierr *bem.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityBulkValidate(t *testing.T) {
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
	_, err := client.Entities.BulkValidate(context.TODO(), bem.EntityBulkValidateParams{
		EntityIDs: []string{"ent_2abc", "ent_2def"},
		Status:    bem.EntityBulkValidateParamsStatusApproved,
	})
	if err != nil {
		var apierr *bem.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityGetRelationsWithOptionalParams(t *testing.T) {
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
	_, err := client.Entities.GetRelations(
		context.TODO(),
		"id",
		bem.EntityGetRelationsParams{
			Bucket:       bem.String("bucket"),
			Cursor:       bem.String("cursor"),
			Direction:    bem.EntityGetRelationsParamsDirectionInbound,
			Limit:        bem.Int(0),
			RelationType: bem.String("relationType"),
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

func TestEntityGetSeedStatus(t *testing.T) {
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
	_, err := client.Entities.GetSeedStatus(context.TODO(), "id")
	if err != nil {
		var apierr *bem.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
