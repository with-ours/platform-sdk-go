// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/with-ours/platform-sdk-go/v2"
	"github.com/with-ours/platform-sdk-go/v2/internal/testutil"
	"github.com/with-ours/platform-sdk-go/v2/option"
)

func TestAttributionReportListWithOptionalParams(t *testing.T) {
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
	_, err := client.AttributionReports.List(context.TODO(), oursprivacy.AttributionReportListParams{
		Cursor: oursprivacy.String("cursor"),
		Limit:  oursprivacy.Int(0),
		Search: oursprivacy.String("search"),
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttributionReportNewWithOptionalParams(t *testing.T) {
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
	_, err := client.AttributionReports.New(context.TODO(), oursprivacy.AttributionReportNewParams{
		Definition: oursprivacy.AttributionReportNewParamsDefinition{
			Kind: "attribution",
			Outcomes: []oursprivacy.AttributionReportNewParamsDefinitionOutcome{{
				ID:         "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				EventNames: []string{"x"},
				Label:      "x",
				Filter: oursprivacy.AttributionReportNewParamsDefinitionOutcomeFilter{
					Filter: oursprivacy.AttributionReportNewParamsDefinitionOutcomeFilterFilterUnion{
						OfAttributionReportNewsDefinitionOutcomeFilterFilterObject: &oursprivacy.AttributionReportNewParamsDefinitionOutcomeFilterFilterObject{
							Children: []any{},
							Kind:     "and",
						},
					},
					Version: 1,
				},
				Value: oursprivacy.AttributionReportNewParamsDefinitionOutcomeValue{
					Path:     []string{"x"},
					Property: "event.properties",
					Type:     "number",
				},
			}},
			Version:  1,
			Counting: "unique",
			Currency: "USD",
			Filter: oursprivacy.AttributionReportNewParamsDefinitionFilter{
				Filter: oursprivacy.AttributionReportNewParamsDefinitionFilterFilterUnion{
					OfAttributionReportNewsDefinitionFilterFilterObject: &oursprivacy.AttributionReportNewParamsDefinitionFilterFilterObject{
						Children: []any{},
						Kind:     "and",
					},
				},
				Version: 1,
			},
			LookbackDays: 7,
			Model:        "first_touch",
		},
		Name: "x",
		ReportContext: oursprivacy.AttributionReportNewParamsReportContext{
			DateRange: oursprivacy.AttributionReportNewParamsReportContextDateRangeUnion{
				OfAttributionReportNewsReportContextDateRangeObject: &oursprivacy.AttributionReportNewParamsReportContextDateRangeObject{
					From:     "7321-69-10",
					To:       "7321-69-10",
					TimeZone: "UTC",
				},
			},
		},
		Description: oursprivacy.String("description"),
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttributionReportGet(t *testing.T) {
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
	_, err := client.AttributionReports.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttributionReportUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.AttributionReports.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		oursprivacy.AttributionReportUpdateParams{
			ExpectedRevision: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Definition: oursprivacy.AttributionReportUpdateParamsDefinition{
				Kind: "attribution",
				Outcomes: []oursprivacy.AttributionReportUpdateParamsDefinitionOutcome{{
					ID:         "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
					EventNames: []string{"x"},
					Label:      "x",
					Filter: oursprivacy.AttributionReportUpdateParamsDefinitionOutcomeFilter{
						Filter: oursprivacy.AttributionReportUpdateParamsDefinitionOutcomeFilterFilterUnion{
							OfAttributionReportUpdatesDefinitionOutcomeFilterFilterObject: &oursprivacy.AttributionReportUpdateParamsDefinitionOutcomeFilterFilterObject{
								Children: []any{},
								Kind:     "and",
							},
						},
						Version: 1,
					},
					Value: oursprivacy.AttributionReportUpdateParamsDefinitionOutcomeValue{
						Path:     []string{"x"},
						Property: "event.properties",
						Type:     "number",
					},
				}},
				Version:  1,
				Counting: "unique",
				Currency: "USD",
				Filter: oursprivacy.AttributionReportUpdateParamsDefinitionFilter{
					Filter: oursprivacy.AttributionReportUpdateParamsDefinitionFilterFilterUnion{
						OfAttributionReportUpdatesDefinitionFilterFilterObject: &oursprivacy.AttributionReportUpdateParamsDefinitionFilterFilterObject{
							Children: []any{},
							Kind:     "and",
						},
					},
					Version: 1,
				},
				LookbackDays: 7,
				Model:        "first_touch",
			},
			Description: oursprivacy.String("description"),
			Name:        oursprivacy.String("x"),
			ReportContext: oursprivacy.AttributionReportUpdateParamsReportContext{
				DateRange: oursprivacy.AttributionReportUpdateParamsReportContextDateRangeUnion{
					OfAttributionReportUpdatesReportContextDateRangeObject: &oursprivacy.AttributionReportUpdateParamsReportContextDateRangeObject{
						From:     "7321-69-10",
						To:       "7321-69-10",
						TimeZone: "UTC",
					},
				},
			},
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

func TestAttributionReportDelete(t *testing.T) {
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
	_, err := client.AttributionReports.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		oursprivacy.AttributionReportDeleteParams{
			ExpectedRevision: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
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

func TestAttributionReportDuplicate(t *testing.T) {
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
	_, err := client.AttributionReports.Duplicate(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		oursprivacy.AttributionReportDuplicateParams{
			ExpectedRevision: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
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

func TestAttributionReportPreview(t *testing.T) {
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
	_, err := client.AttributionReports.Preview(context.TODO(), oursprivacy.AttributionReportPreviewParams{
		Input: `{"definition":{"version":1,"kind":"attribution","outcomes":[{"id":"11111111-1111-4111-8111-111111111111","label":"Bookings","eventNames":["appointment_booked"]}],"model":"last_touch","counting":"unique","lookbackDays":30,"currency":"USD"},"from":"2026-09-01","to":"2026-09-30","comparison":{"kind":"previous_period"}}`,
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttributionReportVisitors(t *testing.T) {
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
	_, err := client.AttributionReports.Visitors(context.TODO(), oursprivacy.AttributionReportVisitorsParams{
		Input: `{"definition":{"version":1,"kind":"attribution","outcomes":[{"id":"11111111-1111-4111-8111-111111111111","label":"Bookings","eventNames":["appointment_booked"]}],"model":"last_touch","counting":"unique","lookbackDays":30,"currency":"USD"},"from":"2026-09-01","to":"2026-09-30","comparison":{"kind":"previous_period"},"selection":{"outcomeId":"11111111-1111-4111-8111-111111111111","period":"current","grain":"source_medium","channel":"Paid search","sourceMedium":"google / cpc","touchRole":"last"},"limit":25}`,
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttributionReportInsights(t *testing.T) {
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
	_, err := client.AttributionReports.Insights(context.TODO(), oursprivacy.AttributionReportInsightsParams{
		Input: `{"definition":{"version":1,"kind":"attribution","outcomes":[{"id":"11111111-1111-4111-8111-111111111111","label":"Bookings","eventNames":["appointment_booked"]}],"model":"last_touch","counting":"unique","lookbackDays":30,"currency":"USD"},"from":"2026-09-01","to":"2026-09-30","comparison":{"kind":"previous_period"}}`,
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
