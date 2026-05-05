// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomwithoursplatformsdkgo_test

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
	client := githubcomwithoursplatformsdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Experiments.List(context.TODO(), githubcomwithoursplatformsdkgo.ExperimentListParams{
		Cursor: githubcomwithoursplatformsdkgo.String("cursor"),
		Limit:  githubcomwithoursplatformsdkgo.Int(25),
		Search: githubcomwithoursplatformsdkgo.String("pricing hero"),
		Status: githubcomwithoursplatformsdkgo.ExperimentListParamsStatusCompleted,
		Type:   githubcomwithoursplatformsdkgo.ExperimentListParamsTypeAb,
	})
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
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
	client := githubcomwithoursplatformsdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Experiments.New(context.TODO(), githubcomwithoursplatformsdkgo.ExperimentNewParams{
		ExperimentSettingsID: "settings_01HZX9BB73EY2Q37VGK5A0VW7A",
		Name:                 "Homepage Hero Headline Test",
		Description:          githubcomwithoursplatformsdkgo.String("description"),
		IncludeQueryString:   githubcomwithoursplatformsdkgo.Bool(true),
		Key:                  githubcomwithoursplatformsdkgo.String("homepage-hero-headline-test"),
		Metrics: githubcomwithoursplatformsdkgo.ExperimentNewParamsMetrics{
			Primary: githubcomwithoursplatformsdkgo.ExperimentNewParamsMetricsPrimary{
				EventName: githubcomwithoursplatformsdkgo.String("demo_requested"),
				FunnelID:  githubcomwithoursplatformsdkgo.String("funnelId"),
			},
			Secondary: []githubcomwithoursplatformsdkgo.ExperimentNewParamsMetricsSecondary{{
				EventName: githubcomwithoursplatformsdkgo.String("demo_requested"),
				FunnelID:  githubcomwithoursplatformsdkgo.String("funnelId"),
			}},
		},
		TargetingRules: githubcomwithoursplatformsdkgo.ExperimentNewParamsTargetingRules{
			URLPatterns: []string{"/pricing*", "/enterprise"},
			AudienceID:  githubcomwithoursplatformsdkgo.String("audienceId"),
			QueryParams: []githubcomwithoursplatformsdkgo.ExperimentNewParamsTargetingRulesQueryParam{{
				Key:      "utm_campaign",
				Operator: "contains",
				Value:    githubcomwithoursplatformsdkgo.String("spring-launch"),
			}},
			VisitorProperties: map[string]any{},
			VisitorStatus:     githubcomwithoursplatformsdkgo.String("visitorStatus"),
		},
		TrafficAllocation: githubcomwithoursplatformsdkgo.Float(100),
		Type:              githubcomwithoursplatformsdkgo.ExperimentNewParamsTypeAb,
	})
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
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
	client := githubcomwithoursplatformsdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Experiments.Get(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
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
	client := githubcomwithoursplatformsdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Experiments.Update(
		context.TODO(),
		"id",
		githubcomwithoursplatformsdkgo.ExperimentUpdateParams{
			Description:        githubcomwithoursplatformsdkgo.String("description"),
			IncludeQueryString: githubcomwithoursplatformsdkgo.Bool(true),
			Key:                githubcomwithoursplatformsdkgo.String("key"),
			Metrics: githubcomwithoursplatformsdkgo.ExperimentUpdateParamsMetrics{
				Primary: githubcomwithoursplatformsdkgo.ExperimentUpdateParamsMetricsPrimary{
					EventName: githubcomwithoursplatformsdkgo.String("demo_requested"),
					FunnelID:  githubcomwithoursplatformsdkgo.String("funnelId"),
				},
				Secondary: []githubcomwithoursplatformsdkgo.ExperimentUpdateParamsMetricsSecondary{{
					EventName: githubcomwithoursplatformsdkgo.String("demo_requested"),
					FunnelID:  githubcomwithoursplatformsdkgo.String("funnelId"),
				}},
			},
			Name: githubcomwithoursplatformsdkgo.String("name"),
			TargetingRules: githubcomwithoursplatformsdkgo.ExperimentUpdateParamsTargetingRules{
				URLPatterns: []string{"/pricing*", "/enterprise"},
				AudienceID:  githubcomwithoursplatformsdkgo.String("audienceId"),
				QueryParams: []githubcomwithoursplatformsdkgo.ExperimentUpdateParamsTargetingRulesQueryParam{{
					Key:      "utm_campaign",
					Operator: "contains",
					Value:    githubcomwithoursplatformsdkgo.String("spring-launch"),
				}},
				VisitorProperties: map[string]any{},
				VisitorStatus:     githubcomwithoursplatformsdkgo.String("visitorStatus"),
			},
			TrafficAllocation: githubcomwithoursplatformsdkgo.Float(0),
		},
	)
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
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
	client := githubcomwithoursplatformsdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Experiments.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
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
	client := githubcomwithoursplatformsdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Experiments.Start(
		context.TODO(),
		"id",
		githubcomwithoursplatformsdkgo.ExperimentStartParams{
			PublishAfterStart: githubcomwithoursplatformsdkgo.Bool(true),
		},
	)
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
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
	client := githubcomwithoursplatformsdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Experiments.Stop(
		context.TODO(),
		"id",
		githubcomwithoursplatformsdkgo.ExperimentStopParams{
			WinnerVariantID: githubcomwithoursplatformsdkgo.String("var_01HZX8YJH3Z3W1R2Q4M5N6P7Q8"),
		},
	)
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
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
	client := githubcomwithoursplatformsdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Experiments.Pause(
		context.TODO(),
		"id",
		githubcomwithoursplatformsdkgo.ExperimentPauseParams{
			PublishAfterPause: githubcomwithoursplatformsdkgo.Bool(true),
		},
	)
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
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
	client := githubcomwithoursplatformsdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Experiments.Resume(
		context.TODO(),
		"id",
		githubcomwithoursplatformsdkgo.ExperimentResumeParams{
			PublishAfterResume: githubcomwithoursplatformsdkgo.Bool(true),
		},
	)
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
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
	client := githubcomwithoursplatformsdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Experiments.Results(
		context.TODO(),
		"id",
		githubcomwithoursplatformsdkgo.ExperimentResultsParams{
			EventName: githubcomwithoursplatformsdkgo.String("demo_requested"),
		},
	)
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
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
	client := githubcomwithoursplatformsdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Experiments.ResultsTimeSeries(
		context.TODO(),
		"id",
		githubcomwithoursplatformsdkgo.ExperimentResultsTimeSeriesParams{
			EndDate:   githubcomwithoursplatformsdkgo.String("2026-04-30"),
			EventName: githubcomwithoursplatformsdkgo.String("demo_requested"),
			StartDate: githubcomwithoursplatformsdkgo.String("2026-04-01"),
		},
	)
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
