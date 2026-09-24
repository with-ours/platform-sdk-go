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

func TestAttributionInitialWithOptionalParams(t *testing.T) {
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
	_, err := client.Attribution.Initial(context.TODO(), oursprivacy.AttributionInitialParams{
		EventName:       "purchase",
		From:            "2026-05-01",
		To:              "2026-06-30",
		AttributionType: oursprivacy.AttributionInitialParamsAttributionTypeInitial,
		UtmCampaign:     oursprivacy.String("x"),
		UtmContent:      oursprivacy.String("x"),
		UtmMedium:       oursprivacy.String("x"),
		UtmName:         oursprivacy.String("x"),
		UtmSource:       oursprivacy.String("x"),
		UtmTerm:         oursprivacy.String("x"),
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttributionLastTouchWithOptionalParams(t *testing.T) {
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
	_, err := client.Attribution.LastTouch(context.TODO(), oursprivacy.AttributionLastTouchParams{
		EventName:       "purchase",
		From:            "2026-05-01",
		To:              "2026-06-30",
		AttributionType: oursprivacy.AttributionLastTouchParamsAttributionTypeInitial,
		UtmCampaign:     oursprivacy.String("x"),
		UtmContent:      oursprivacy.String("x"),
		UtmMedium:       oursprivacy.String("x"),
		UtmName:         oursprivacy.String("x"),
		UtmSource:       oursprivacy.String("x"),
		UtmTerm:         oursprivacy.String("x"),
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttributionConversionWithOptionalParams(t *testing.T) {
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
	_, err := client.Attribution.Conversion(context.TODO(), oursprivacy.AttributionConversionParams{
		AttributionModel: oursprivacy.AttributionConversionParamsAttributionModelFirstTouch,
		EventName:        "purchase",
		From:             "2026-06-01",
		To:               "2026-06-30",
		Limit:            oursprivacy.Int(1),
		LookbackWindow:   oursprivacy.AttributionConversionParamsLookbackWindowSevenDays,
		WebSourceID:      oursprivacy.String("x"),
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttributionAudienceConversionWithOptionalParams(t *testing.T) {
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
	_, err := client.Attribution.AudienceConversion(context.TODO(), oursprivacy.AttributionAudienceConversionParams{
		EventName:         "purchase",
		From:              "2026-05-01",
		To:                "2026-06-30",
		AttributionWindow: oursprivacy.String("IN_RANGE"),
		ExcludeBots:       oursprivacy.AttributionAudienceConversionParamsExcludeBotsTrue,
		ValueProperty:     oursprivacy.String("revenue"),
		WebSourceID:       oursprivacy.String("550e8400-e29b-41d4-a716-446655440000"),
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAttributionUtmComparison(t *testing.T) {
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
	_, err := client.Attribution.UtmComparison(context.TODO(), oursprivacy.AttributionUtmComparisonParams{
		Combos:    `[{"utmSource":"google","utmMedium":"cpc"},{"utmSource":"meta"}]`,
		EventName: "purchase",
		From:      "2026-06-01",
		To:        "2026-06-30",
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
