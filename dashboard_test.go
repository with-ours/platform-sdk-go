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

func TestDashboardListWithOptionalParams(t *testing.T) {
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
	_, err := client.Dashboards.List(context.TODO(), oursprivacy.DashboardListParams{
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

func TestDashboardNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Dashboards.New(context.TODO(), oursprivacy.DashboardNewParams{
		Name: "x",
		ReportContext: oursprivacy.DashboardNewParamsReportContext{
			DateRange: oursprivacy.DashboardNewParamsReportContextDateRangeUnion{
				OfDashboardNewsReportContextDateRangeObject: &oursprivacy.DashboardNewParamsReportContextDateRangeObject{
					From:     "7321-69-10",
					To:       "7321-69-10",
					TimeZone: "UTC",
				},
			},
		},
		Tiles: []oursprivacy.DashboardNewParamsTileUnion{{
			OfDashboardNewsTileObject: &oursprivacy.DashboardNewParamsTileObject{
				ID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				Layout: oursprivacy.DashboardNewParamsTileObjectLayout{
					Height: 240,
					Width:  3,
				},
				Report: oursprivacy.DashboardNewParamsTileObjectReport{
					ID:   "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
					Kind: "insight",
				},
			},
		}},
		Description: oursprivacy.String("description"),
		Filter: oursprivacy.DashboardNewParamsFilter{
			Filter: oursprivacy.DashboardNewParamsFilterFilterUnion{
				OfDashboardNewsFilterFilterObject: &oursprivacy.DashboardNewParamsFilterFilterObject{
					Children: []oursprivacy.DashboardNewParamsFilterFilterObjectChildUnion{{
						OfDashboardNewsFilterFilterObjectChildObject: &oursprivacy.DashboardNewParamsFilterFilterObjectChildObject{
							Children: []any{},
							Kind:     "and",
						},
					}},
					Kind: "and",
				},
			},
			Version: 1,
		},
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDashboardGet(t *testing.T) {
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
	_, err := client.Dashboards.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDashboardUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Dashboards.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		oursprivacy.DashboardUpdateParams{
			ExpectedRevision: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Description:      oursprivacy.String("description"),
			Filter: oursprivacy.DashboardUpdateParamsFilter{
				Filter: oursprivacy.DashboardUpdateParamsFilterFilterUnion{
					OfDashboardUpdatesFilterFilterObject: &oursprivacy.DashboardUpdateParamsFilterFilterObject{
						Children: []oursprivacy.DashboardUpdateParamsFilterFilterObjectChildUnion{{
							OfDashboardUpdatesFilterFilterObjectChildObject: &oursprivacy.DashboardUpdateParamsFilterFilterObjectChildObject{
								Children: []any{},
								Kind:     "and",
							},
						}},
						Kind: "and",
					},
				},
				Version: 1,
			},
			Name: oursprivacy.String("x"),
			ReportContext: oursprivacy.DashboardUpdateParamsReportContext{
				DateRange: oursprivacy.DashboardUpdateParamsReportContextDateRangeUnion{
					OfDashboardUpdatesReportContextDateRangeObject: &oursprivacy.DashboardUpdateParamsReportContextDateRangeObject{
						From:     "7321-69-10",
						To:       "7321-69-10",
						TimeZone: "UTC",
					},
				},
			},
			Tiles: []oursprivacy.DashboardUpdateParamsTileUnion{{
				OfDashboardUpdatesTileObject: &oursprivacy.DashboardUpdateParamsTileObject{
					ID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
					Layout: oursprivacy.DashboardUpdateParamsTileObjectLayout{
						Height: 240,
						Width:  3,
					},
					Report: oursprivacy.DashboardUpdateParamsTileObjectReport{
						ID:   "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
						Kind: "insight",
					},
				},
			}},
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

func TestDashboardDelete(t *testing.T) {
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
	_, err := client.Dashboards.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		oursprivacy.DashboardDeleteParams{
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

func TestDashboardDuplicate(t *testing.T) {
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
	_, err := client.Dashboards.Duplicate(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		oursprivacy.DashboardDuplicateParams{
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

func TestDashboardPreview(t *testing.T) {
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
	_, err := client.Dashboards.Preview(context.TODO(), oursprivacy.DashboardPreviewParams{
		Input: `{"tileId":"11111111-1111-4111-8111-111111111111","report":{"kind":"insight","id":"22222222-2222-4222-8222-222222222222"},"context":{"from":"2026-09-01","to":"2026-09-07"},"comparison":{"kind":"previous_period"}}`,
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDashboardResultsWithOptionalParams(t *testing.T) {
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
	_, err := client.Dashboards.Results(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		oursprivacy.DashboardResultsParams{
			Context:          "context",
			ExpectedRevision: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			TileID:           "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Comparison:       oursprivacy.String("comparison"),
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

func TestDashboardReplayTileResult(t *testing.T) {
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
	_, err := client.Dashboards.ReplayTileResult(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		oursprivacy.DashboardReplayTileResultParams{
			Context:          "context",
			ExpectedRevision: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			TileID:           "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
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

func TestDashboardReplayPreview(t *testing.T) {
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
	_, err := client.Dashboards.ReplayPreview(context.TODO(), oursprivacy.DashboardReplayPreviewParams{
		Input: `{"tileId":"11111111-1111-4111-8111-111111111111","replays":{"title":"Recent checkouts","pathname":"/checkout","limit":10},"context":{"from":"2026-09-01","to":"2026-09-07"}}`,
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
