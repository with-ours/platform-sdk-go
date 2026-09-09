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

func TestConversionJourneySummaryListWithOptionalParams(t *testing.T) {
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
	_, err := client.ConversionJourneySummaries.List(context.TODO(), oursprivacy.ConversionJourneySummaryListParams{
		Cursor: oursprivacy.String("cursor"),
		Limit:  oursprivacy.Int(25),
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConversionJourneySummaryNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ConversionJourneySummaries.New(context.TODO(), oursprivacy.ConversionJourneySummaryNewParams{
		DateFrom:    "2026-06-01",
		DateTo:      "2026-06-30",
		EventName:   "x",
		Name:        "x",
		WindowDays:  1,
		ExcludeBots: oursprivacy.Bool(true),
		Filters: []oursprivacy.ConversionJourneySummaryNewParamsFilter{{
			Dimension: "browser",
			Operator:  "CONTAINS",
			Value:     oursprivacy.String("x"),
			Values:    []string{"x"},
		}},
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

func TestConversionJourneySummaryGet(t *testing.T) {
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
	_, err := client.ConversionJourneySummaries.Get(context.TODO(), "id")
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConversionJourneySummaryUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ConversionJourneySummaries.Update(
		context.TODO(),
		"id",
		oursprivacy.ConversionJourneySummaryUpdateParams{
			DateFrom:    oursprivacy.String("7321-69-10"),
			DateTo:      oursprivacy.String("7321-69-10"),
			EventName:   oursprivacy.String("x"),
			ExcludeBots: oursprivacy.Bool(true),
			Filters: []oursprivacy.ConversionJourneySummaryUpdateParamsFilter{{
				Dimension: "browser",
				Operator:  "CONTAINS",
				Value:     oursprivacy.String("x"),
				Values:    []string{"x"},
			}},
			Name:        oursprivacy.String("x"),
			WebSourceID: oursprivacy.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			WindowDays:  oursprivacy.Int(1),
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

func TestConversionJourneySummaryDelete(t *testing.T) {
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
	_, err := client.ConversionJourneySummaries.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
