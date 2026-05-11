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

func TestHeatmapPageListWithOptionalParams(t *testing.T) {
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
	_, err := client.HeatmapPages.List(context.TODO(), oursprivacy.HeatmapPageListParams{
		From:        "2026-04-01",
		To:          "2026-04-30",
		BrowserName: oursprivacy.String("Chrome"),
		Country:     oursprivacy.String("x"),
		Cursor:      oursprivacy.String("cursor"),
		Limit:       oursprivacy.Int(25),
		Region:      oursprivacy.String("x"),
		Search:      oursprivacy.String("/pricing"),
		SortBy:      oursprivacy.HeatmapPageListParamsSortByClicks,
		SortDir:     oursprivacy.HeatmapPageListParamsSortDirAsc,
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHeatmapPageSummaryWithOptionalParams(t *testing.T) {
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
	_, err := client.HeatmapPages.Summary(context.TODO(), oursprivacy.HeatmapPageSummaryParams{
		Breakpoint:  oursprivacy.HeatmapPageSummaryParamsBreakpointDesktop,
		From:        "2026-04-01",
		PageKey:     "https://example.com/pricing",
		To:          "2026-04-30",
		BrowserName: oursprivacy.String("Chrome"),
		Country:     oursprivacy.String("x"),
		Region:      oursprivacy.String("x"),
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
