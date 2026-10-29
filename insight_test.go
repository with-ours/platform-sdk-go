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

func TestInsightListWithOptionalParams(t *testing.T) {
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
	_, err := client.Insights.List(context.TODO(), oursprivacy.InsightListParams{
		Cursor: oursprivacy.String("cursor"),
		Limit:  oursprivacy.Int(25),
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

func TestInsightNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Insights.New(context.TODO(), oursprivacy.InsightNewParams{
		Definition: oursprivacy.InsightNewParamsDefinition{
			Interval: "day",
			Kind:     "insight",
			Metrics: []oursprivacy.InsightNewParamsDefinitionMetric{{
				ID:          "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				Aggregation: "event_count",
				EventNames:  []string{"x"},
				Name:        "x",
				Filter: oursprivacy.InsightNewParamsDefinitionMetricFilter{
					Filter: oursprivacy.InsightNewParamsDefinitionMetricFilterFilterUnion{
						OfInsightNewsDefinitionMetricFilterFilterObject: &oursprivacy.InsightNewParamsDefinitionMetricFilterFilterObject{
							Children: []any{},
							Kind:     "and",
						},
					},
					Version: 1,
				},
				Property: oursprivacy.InsightNewParamsDefinitionMetricProperty{
					Path:     []string{"x"},
					Property: "event.properties",
					Type:     "number",
				},
			}},
			Version: 1,
			Filter: oursprivacy.InsightNewParamsDefinitionFilter{
				Filter: oursprivacy.InsightNewParamsDefinitionFilterFilterUnion{
					OfInsightNewsDefinitionFilterFilterObject: &oursprivacy.InsightNewParamsDefinitionFilterFilterObject{
						Children: []any{},
						Kind:     "and",
					},
				},
				Version: 1,
			},
		},
		Name: "x",
		ReportContext: oursprivacy.InsightNewParamsReportContext{
			DateRange: oursprivacy.InsightNewParamsReportContextDateRangeUnion{
				OfInsightNewsReportContextDateRangeObject: &oursprivacy.InsightNewParamsReportContextDateRangeObject{
					From:     "7321-69-10",
					To:       "7321-69-10",
					TimeZone: "UTC",
				},
			},
		},
		View: oursprivacy.InsightNewParamsView{
			Cumulative: oursprivacy.Bool(true),
			Series: []oursprivacy.InsightNewParamsViewSeries{{
				MetricID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				Color:    oursprivacy.String("#E1CB97"),
				Hidden:   oursprivacy.Bool(true),
			}},
			Smoothing:       "none",
			SmoothingWindow: oursprivacy.Int(2),
			Stacked:         oursprivacy.Bool(true),
			TableSort: oursprivacy.InsightNewParamsViewTableSort{
				Column:    "label",
				Direction: "asc",
			},
			Visualization: "line",
			YAxisScale:    "linear",
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

func TestInsightGet(t *testing.T) {
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
	_, err := client.Insights.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInsightUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Insights.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		oursprivacy.InsightUpdateParams{
			ExpectedRevision: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Definition: oursprivacy.InsightUpdateParamsDefinition{
				Interval: "day",
				Kind:     "insight",
				Metrics: []oursprivacy.InsightUpdateParamsDefinitionMetric{{
					ID:          "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
					Aggregation: "event_count",
					EventNames:  []string{"x"},
					Name:        "x",
					Filter: oursprivacy.InsightUpdateParamsDefinitionMetricFilter{
						Filter: oursprivacy.InsightUpdateParamsDefinitionMetricFilterFilterUnion{
							OfInsightUpdatesDefinitionMetricFilterFilterObject: &oursprivacy.InsightUpdateParamsDefinitionMetricFilterFilterObject{
								Children: []any{},
								Kind:     "and",
							},
						},
						Version: 1,
					},
					Property: oursprivacy.InsightUpdateParamsDefinitionMetricProperty{
						Path:     []string{"x"},
						Property: "event.properties",
						Type:     "number",
					},
				}},
				Version: 1,
				Filter: oursprivacy.InsightUpdateParamsDefinitionFilter{
					Filter: oursprivacy.InsightUpdateParamsDefinitionFilterFilterUnion{
						OfInsightUpdatesDefinitionFilterFilterObject: &oursprivacy.InsightUpdateParamsDefinitionFilterFilterObject{
							Children: []any{},
							Kind:     "and",
						},
					},
					Version: 1,
				},
			},
			Description: oursprivacy.String("description"),
			Name:        oursprivacy.String("x"),
			ReportContext: oursprivacy.InsightUpdateParamsReportContext{
				DateRange: oursprivacy.InsightUpdateParamsReportContextDateRangeUnion{
					OfInsightUpdatesReportContextDateRangeObject: &oursprivacy.InsightUpdateParamsReportContextDateRangeObject{
						From:     "7321-69-10",
						To:       "7321-69-10",
						TimeZone: "UTC",
					},
				},
			},
			View: oursprivacy.InsightUpdateParamsView{
				Cumulative: oursprivacy.Bool(true),
				Series: []oursprivacy.InsightUpdateParamsViewSeries{{
					MetricID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
					Color:    oursprivacy.String("#E1CB97"),
					Hidden:   oursprivacy.Bool(true),
				}},
				Smoothing:       "none",
				SmoothingWindow: oursprivacy.Int(2),
				Stacked:         oursprivacy.Bool(true),
				TableSort: oursprivacy.InsightUpdateParamsViewTableSort{
					Column:    "label",
					Direction: "asc",
				},
				Visualization: "line",
				YAxisScale:    "linear",
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

func TestInsightDelete(t *testing.T) {
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
	_, err := client.Insights.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		oursprivacy.InsightDeleteParams{
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

func TestInsightCapabilities(t *testing.T) {
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
	_, err := client.Insights.Capabilities(context.TODO())
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInsightResultsWithOptionalParams(t *testing.T) {
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
	_, err := client.Insights.Results(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		oursprivacy.InsightResultsParams{
			ExpectedRevision:        "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Comparison:              oursprivacy.String("comparison"),
			CumulativeOverride:      oursprivacy.Bool(true),
			ReportContextOverride:   oursprivacy.String("reportContextOverride"),
			SmoothingWindowOverride: oursprivacy.Int(2),
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

func TestInsightPreview(t *testing.T) {
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
	_, err := client.Insights.Preview(context.TODO(), oursprivacy.InsightPreviewParams{
		Input: "input",
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
