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
		Steps: []oursprivacy.FunnelNewParamsStep{{
			EventName: "x",
			Name:      "x",
			Order:     0,
			Filters:   map[string]any{},
			Logic: oursprivacy.FunnelNewParamsStepLogic{
				And: []any{map[string]any{}},
				Condition: oursprivacy.FunnelNewParamsStepLogicCondition{
					Operator: "Is",
					Property: "property",
					Value:    "value",
				},
				Not: map[string]any{},
				Or:  []any{map[string]any{}},
			},
		}, {
			EventName: "x",
			Name:      "x",
			Order:     0,
			Filters:   map[string]any{},
			Logic: oursprivacy.FunnelNewParamsStepLogic{
				And: []any{map[string]any{}},
				Condition: oursprivacy.FunnelNewParamsStepLogicCondition{
					Operator: "Is",
					Property: "property",
					Value:    "value",
				},
				Not: map[string]any{},
				Or:  []any{map[string]any{}},
			},
		}},
		ConversionWindow: map[string]any{},
		CountingMethod:   oursprivacy.String("countingMethod"),
		Description:      oursprivacy.String("description"),
		FunnelType:       oursprivacy.FunnelNewParamsFunnelTypeSessionBased,
		GlobalLogic: oursprivacy.FunnelNewParamsGlobalLogic{
			And: []any{map[string]any{}},
			Condition: oursprivacy.FunnelNewParamsGlobalLogicCondition{
				Operator: "Is",
				Property: "property",
				Value:    "value",
			},
			Not: map[string]any{},
			Or:  []any{map[string]any{}},
		},
		StepOrder:  oursprivacy.String("stepOrder"),
		UtmFilters: map[string]any{},
		Watched:    oursprivacy.Bool(true),
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
			ConversionWindow: map[string]any{},
			CountingMethod:   oursprivacy.String("countingMethod"),
			Description:      oursprivacy.String("description"),
			FunnelType:       oursprivacy.FunnelUpdateParamsFunnelTypeSessionBased,
			GlobalLogic: oursprivacy.FunnelUpdateParamsGlobalLogic{
				And: []any{map[string]any{}},
				Condition: oursprivacy.FunnelUpdateParamsGlobalLogicCondition{
					Operator: "Is",
					Property: "property",
					Value:    "value",
				},
				Not: map[string]any{},
				Or:  []any{map[string]any{}},
			},
			Name:      oursprivacy.String("x"),
			StepOrder: oursprivacy.String("stepOrder"),
			Steps: []oursprivacy.FunnelUpdateParamsStep{{
				EventName: "x",
				Name:      "x",
				Order:     0,
				Filters:   map[string]any{},
				Logic: oursprivacy.FunnelUpdateParamsStepLogic{
					And: []any{map[string]any{}},
					Condition: oursprivacy.FunnelUpdateParamsStepLogicCondition{
						Operator: "Is",
						Property: "property",
						Value:    "value",
					},
					Not: map[string]any{},
					Or:  []any{map[string]any{}},
				},
			}, {
				EventName: "x",
				Name:      "x",
				Order:     0,
				Filters:   map[string]any{},
				Logic: oursprivacy.FunnelUpdateParamsStepLogic{
					And: []any{map[string]any{}},
					Condition: oursprivacy.FunnelUpdateParamsStepLogicCondition{
						Operator: "Is",
						Property: "property",
						Value:    "value",
					},
					Not: map[string]any{},
					Or:  []any{map[string]any{}},
				},
			}},
			UtmFilters: map[string]any{},
			Watched:    oursprivacy.Bool(true),
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
			From:            "2026-06-01",
			To:              "2026-06-30",
			AttributionType: oursprivacy.FunnelResultsParamsAttributionTypeInitial,
			DeviceType:      oursprivacy.FunnelResultsParamsDeviceTypeDesktop,
			UtmCampaign:     oursprivacy.String("spring-promo"),
			UtmContent:      oursprivacy.String("x"),
			UtmMedium:       oursprivacy.String("cpc"),
			UtmName:         oursprivacy.String("x"),
			UtmSource:       oursprivacy.String("google"),
			UtmTerm:         oursprivacy.String("x"),
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
