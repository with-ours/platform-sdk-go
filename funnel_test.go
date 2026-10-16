// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/with-ours/platform-sdk-go/v2"
	"github.com/with-ours/platform-sdk-go/v2/internal/testutil"
	"github.com/with-ours/platform-sdk-go/v2/option"
)

func TestFunnelList(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := oursprivacy.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Funnels.List(context.TODO())
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFunnelNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := oursprivacy.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Funnels.New(context.TODO(), oursprivacy.FunnelNewParams{
		Name: "x",
		QueryDefinition: oursprivacy.FunnelNewParamsQueryDefinition{
			Counting: "unique",
			Identity: "session",
			Kind:     "funnel",
			Order:    "ordered",
			Reentry:  "first",
			Steps: []oursprivacy.FunnelNewParamsQueryDefinitionStep{{
				ID:        "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				EventName: "x",
				Name:      "x",
				Filter: oursprivacy.FunnelNewParamsQueryDefinitionStepFilter{
					Filter: oursprivacy.FunnelNewParamsQueryDefinitionStepFilterFilterUnion{
						OfFunnelNewsQueryDefinitionStepFilterFilterObject: &oursprivacy.FunnelNewParamsQueryDefinitionStepFilterFilterObject{
							Children: []any{},
							Kind:     "and",
						},
					},
					Version: 1,
				},
			}, {
				ID:        "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				EventName: "x",
				Name:      "x",
				Filter: oursprivacy.FunnelNewParamsQueryDefinitionStepFilter{
					Filter: oursprivacy.FunnelNewParamsQueryDefinitionStepFilterFilterUnion{
						OfFunnelNewsQueryDefinitionStepFilterFilterObject: &oursprivacy.FunnelNewParamsQueryDefinitionStepFilterFilterObject{
							Children: []any{},
							Kind:     "and",
						},
					},
					Version: 1,
				},
			}},
			Version:  1,
			WindowMs: 1,
			EntryFilter: oursprivacy.FunnelNewParamsQueryDefinitionEntryFilter{
				Filter: oursprivacy.FunnelNewParamsQueryDefinitionEntryFilterFilterUnion{
					OfFunnelNewsQueryDefinitionEntryFilterFilterObject: &oursprivacy.FunnelNewParamsQueryDefinitionEntryFilterFilterObject{
						Children: []any{},
						Kind:     "and",
					},
				},
				Version: 1,
			},
			Filter: oursprivacy.FunnelNewParamsQueryDefinitionFilter{
				Filter: oursprivacy.FunnelNewParamsQueryDefinitionFilterFilterUnion{
					OfFunnelNewsQueryDefinitionFilterFilterObject: &oursprivacy.FunnelNewParamsQueryDefinitionFilterFilterObject{
						Children: []any{},
						Kind:     "and",
					},
				},
				Version: 1,
			},
		},
		Description: oursprivacy.String("description"),
		Watched:     oursprivacy.Bool(true),
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFunnelGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := oursprivacy.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Funnels.Get(context.TODO(), "id")
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFunnelUpdateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := oursprivacy.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Funnels.Update(
		context.TODO(),
		"id",
		oursprivacy.FunnelUpdateParams{
			Description: oursprivacy.String("description"),
			Name:        oursprivacy.String("x"),
			QueryDefinition: oursprivacy.FunnelUpdateParamsQueryDefinition{
				Counting: "unique",
				Identity: "session",
				Kind:     "funnel",
				Order:    "ordered",
				Reentry:  "first",
				Steps: []oursprivacy.FunnelUpdateParamsQueryDefinitionStep{{
					ID:        "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
					EventName: "x",
					Name:      "x",
					Filter: oursprivacy.FunnelUpdateParamsQueryDefinitionStepFilter{
						Filter: oursprivacy.FunnelUpdateParamsQueryDefinitionStepFilterFilterUnion{
							OfFunnelUpdatesQueryDefinitionStepFilterFilterObject: &oursprivacy.FunnelUpdateParamsQueryDefinitionStepFilterFilterObject{
								Children: []any{},
								Kind:     "and",
							},
						},
						Version: 1,
					},
				}, {
					ID:        "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
					EventName: "x",
					Name:      "x",
					Filter: oursprivacy.FunnelUpdateParamsQueryDefinitionStepFilter{
						Filter: oursprivacy.FunnelUpdateParamsQueryDefinitionStepFilterFilterUnion{
							OfFunnelUpdatesQueryDefinitionStepFilterFilterObject: &oursprivacy.FunnelUpdateParamsQueryDefinitionStepFilterFilterObject{
								Children: []any{},
								Kind:     "and",
							},
						},
						Version: 1,
					},
				}},
				Version:  1,
				WindowMs: 1,
				EntryFilter: oursprivacy.FunnelUpdateParamsQueryDefinitionEntryFilter{
					Filter: oursprivacy.FunnelUpdateParamsQueryDefinitionEntryFilterFilterUnion{
						OfFunnelUpdatesQueryDefinitionEntryFilterFilterObject: &oursprivacy.FunnelUpdateParamsQueryDefinitionEntryFilterFilterObject{
							Children: []any{},
							Kind:     "and",
						},
					},
					Version: 1,
				},
				Filter: oursprivacy.FunnelUpdateParamsQueryDefinitionFilter{
					Filter: oursprivacy.FunnelUpdateParamsQueryDefinitionFilterFilterUnion{
						OfFunnelUpdatesQueryDefinitionFilterFilterObject: &oursprivacy.FunnelUpdateParamsQueryDefinitionFilterFilterObject{
							Children: []any{},
							Kind:     "and",
						},
					},
					Version: 1,
				},
			},
			Watched: oursprivacy.Bool(true),
		},
	)
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFunnelDelete(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := oursprivacy.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Funnels.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFunnelDuplicate(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := oursprivacy.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Funnels.Duplicate(context.TODO(), "id")
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFunnelResultsWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := oursprivacy.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Funnels.Results(
		context.TODO(),
		"id",
		oursprivacy.FunnelResultsParams{
			From:                        "2026-06-01",
			To:                          "2026-06-30",
			ExpectedDefinitionUpdatedAt: oursprivacy.Time(time.Now()),
		},
	)
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
