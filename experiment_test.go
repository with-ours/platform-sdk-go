// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/with-ours/platform-sdk-go"
	"github.com/with-ours/platform-sdk-go/internal/testutil"
	"github.com/with-ours/platform-sdk-go/option"
)

func TestExperimentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Experiments.List(context.TODO(), oursprivacy.ExperimentListParams{
		Cursor: oursprivacy.String("cursor"),
		Limit:  oursprivacy.Int(25),
		Search: oursprivacy.String("pricing hero"),
		Status: oursprivacy.ExperimentListParamsStatusCompleted,
		Type:   oursprivacy.ExperimentListParamsTypeAb,
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExperimentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Experiments.New(context.TODO(), oursprivacy.ExperimentNewParams{
		ExperimentSettingsID: "settings_01HZX9BB73EY2Q37VGK5A0VW7A",
		Name:                 "Homepage Hero Headline Test",
		ControlWeight:        oursprivacy.Int(34),
		Description:          oursprivacy.String("description"),
		IncludeQueryString:   oursprivacy.Bool(true),
		Key:                  oursprivacy.String("homepage-hero-headline-test"),
		Metrics: oursprivacy.ExperimentNewParamsMetrics{
			Primary: oursprivacy.ExperimentNewParamsMetricsPrimary{
				EventName: oursprivacy.String("demo_requested"),
				FunnelID:  oursprivacy.String("funnelId"),
			},
			Secondary: []oursprivacy.ExperimentNewParamsMetricsSecondary{{
				EventName: oursprivacy.String("demo_requested"),
				FunnelID:  oursprivacy.String("funnelId"),
			}},
		},
		TargetingRules: oursprivacy.ExperimentNewParamsTargetingRules{
			URLPatterns: []string{"/pricing*", "get.example.com/learn-more"},
			AudienceID:  oursprivacy.String("audienceId"),
			QueryParams: []oursprivacy.ExperimentNewParamsTargetingRulesQueryParam{{
				Key:      "utm_campaign",
				Operator: "contains",
				Value:    oursprivacy.String("spring-launch"),
			}},
			VisitorProperties: map[string]any{},
			VisitorStatus:     oursprivacy.String("visitorStatus"),
		},
		TrafficAllocation: oursprivacy.Float(100),
		Type:              oursprivacy.ExperimentNewParamsTypeAb,
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExperimentGet(t *testing.T) {
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
	_, err := client.Experiments.Get(context.TODO(), "id")
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExperimentUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Experiments.Update(
		context.TODO(),
		"id",
		oursprivacy.ExperimentUpdateParams{
			Description:        oursprivacy.String("description"),
			IncludeQueryString: oursprivacy.Bool(true),
			Key:                oursprivacy.String("key"),
			Metrics: oursprivacy.ExperimentUpdateParamsMetrics{
				Primary: oursprivacy.ExperimentUpdateParamsMetricsPrimary{
					EventName: oursprivacy.String("demo_requested"),
					FunnelID:  oursprivacy.String("funnelId"),
				},
				Secondary: []oursprivacy.ExperimentUpdateParamsMetricsSecondary{{
					EventName: oursprivacy.String("demo_requested"),
					FunnelID:  oursprivacy.String("funnelId"),
				}},
			},
			Name: oursprivacy.String("name"),
			TargetingRules: oursprivacy.ExperimentUpdateParamsTargetingRules{
				URLPatterns: []string{"/pricing*", "get.example.com/learn-more"},
				AudienceID:  oursprivacy.String("audienceId"),
				QueryParams: []oursprivacy.ExperimentUpdateParamsTargetingRulesQueryParam{{
					Key:      "utm_campaign",
					Operator: "contains",
					Value:    oursprivacy.String("spring-launch"),
				}},
				VisitorProperties: map[string]any{},
				VisitorStatus:     oursprivacy.String("visitorStatus"),
			},
			TrafficAllocation: oursprivacy.Float(0),
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

func TestExperimentDelete(t *testing.T) {
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
	_, err := client.Experiments.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExperimentStartWithOptionalParams(t *testing.T) {
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
	_, err := client.Experiments.Start(
		context.TODO(),
		"id",
		oursprivacy.ExperimentStartParams{
			PublishAfterStart: oursprivacy.Bool(true),
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

func TestExperimentStopWithOptionalParams(t *testing.T) {
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
	_, err := client.Experiments.Stop(
		context.TODO(),
		"id",
		oursprivacy.ExperimentStopParams{
			WinnerVariantID: oursprivacy.String("var_01HZX8YJH3Z3W1R2Q4M5N6P7Q8"),
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

func TestExperimentPauseWithOptionalParams(t *testing.T) {
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
	_, err := client.Experiments.Pause(
		context.TODO(),
		"id",
		oursprivacy.ExperimentPauseParams{
			PublishAfterPause: oursprivacy.Bool(true),
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

func TestExperimentResumeWithOptionalParams(t *testing.T) {
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
	_, err := client.Experiments.Resume(
		context.TODO(),
		"id",
		oursprivacy.ExperimentResumeParams{
			PublishAfterResume: oursprivacy.Bool(true),
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

func TestExperimentResultsWithOptionalParams(t *testing.T) {
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
	_, err := client.Experiments.Results(
		context.TODO(),
		"id",
		oursprivacy.ExperimentResultsParams{
			EventName: oursprivacy.String("demo_requested"),
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

func TestExperimentResultsTimeSeriesWithOptionalParams(t *testing.T) {
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
	_, err := client.Experiments.ResultsTimeSeries(
		context.TODO(),
		"id",
		oursprivacy.ExperimentResultsTimeSeriesParams{
			EndDate:   oursprivacy.String("2026-04-30"),
			EventName: oursprivacy.String("demo_requested"),
			StartDate: oursprivacy.String("2026-04-01"),
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

func TestExperimentSessionReplaysWithOptionalParams(t *testing.T) {
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
	_, err := client.Experiments.SessionReplays(
		context.TODO(),
		"id",
		oursprivacy.ExperimentSessionReplaysParams{
			Cursor:    oursprivacy.String("cursor"),
			Limit:     oursprivacy.Int(25),
			VariantID: oursprivacy.String("var_01HZX8YJH3Z3W1R2Q4M5N6P7Q8"),
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
