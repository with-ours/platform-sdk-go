// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/with-ours/platform-sdk-go"
	"github.com/with-ours/platform-sdk-go/internal/testutil"
	"github.com/with-ours/platform-sdk-go/option"
)

func TestWebAnalyticsOverviewWithOptionalParams(t *testing.T) {
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
	_, err := client.WebAnalytics.Overview(context.TODO(), oursprivacy.WebAnalyticsOverviewParams{
		From:         time.Now(),
		Interval:     oursprivacy.WebAnalyticsOverviewParamsIntervalMinute,
		Metric:       oursprivacy.WebAnalyticsOverviewParamsMetricUniqueVisitors,
		To:           time.Now(),
		ExcludeBots:  oursprivacy.Bool(true),
		Filters:      oursprivacy.String("filters"),
		RealtimeFrom: oursprivacy.Time(time.Now()),
		WebSourceID:  oursprivacy.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebAnalyticsSourcesWithOptionalParams(t *testing.T) {
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
	_, err := client.WebAnalytics.Sources(context.TODO(), oursprivacy.WebAnalyticsSourcesParams{
		Dimension:   oursprivacy.WebAnalyticsSourcesParamsDimensionReferrer,
		From:        time.Now(),
		To:          time.Now(),
		ExcludeBots: oursprivacy.Bool(true),
		Filters:     oursprivacy.String("filters"),
		WebSourceID: oursprivacy.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebAnalyticsPagesWithOptionalParams(t *testing.T) {
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
	_, err := client.WebAnalytics.Pages(context.TODO(), oursprivacy.WebAnalyticsPagesParams{
		From:        time.Now(),
		To:          time.Now(),
		View:        oursprivacy.WebAnalyticsPagesParamsViewTop,
		ExcludeBots: oursprivacy.Bool(true),
		Filters:     oursprivacy.String("filters"),
		WebSourceID: oursprivacy.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebAnalyticsLocationsWithOptionalParams(t *testing.T) {
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
	_, err := client.WebAnalytics.Locations(context.TODO(), oursprivacy.WebAnalyticsLocationsParams{
		Dimension:   oursprivacy.WebAnalyticsLocationsParamsDimensionCountry,
		From:        time.Now(),
		To:          time.Now(),
		ExcludeBots: oursprivacy.Bool(true),
		Filters:     oursprivacy.String("filters"),
		WebSourceID: oursprivacy.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebAnalyticsDevicesWithOptionalParams(t *testing.T) {
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
	_, err := client.WebAnalytics.Devices(context.TODO(), oursprivacy.WebAnalyticsDevicesParams{
		Dimension:   oursprivacy.WebAnalyticsDevicesParamsDimensionDevice,
		From:        time.Now(),
		To:          time.Now(),
		ExcludeBots: oursprivacy.Bool(true),
		Filters:     oursprivacy.String("filters"),
		WebSourceID: oursprivacy.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebAnalyticsCurrentVisitorsWithOptionalParams(t *testing.T) {
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
	_, err := client.WebAnalytics.CurrentVisitors(context.TODO(), oursprivacy.WebAnalyticsCurrentVisitorsParams{
		WebSourceID: oursprivacy.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebAnalyticsJourneyWithOptionalParams(t *testing.T) {
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
	_, err := client.WebAnalytics.Journey(context.TODO(), oursprivacy.WebAnalyticsJourneyParams{
		From:        time.Now(),
		Path:        "path",
		To:          time.Now(),
		Direction:   oursprivacy.WebAnalyticsJourneyParamsDirectionForward,
		ExcludeBots: oursprivacy.Bool(true),
		Filters:     oursprivacy.String("filters"),
		Limit:       oursprivacy.Int(1),
		Search:      oursprivacy.String("search"),
		StepKind:    oursprivacy.WebAnalyticsJourneyParamsStepKindPage,
		WebSourceID: oursprivacy.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
