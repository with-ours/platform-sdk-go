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

func TestConsentSettingList(t *testing.T) {
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
	_, err := client.ConsentSettings.List(context.TODO())
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConsentSettingNew(t *testing.T) {
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
	_, err := client.ConsentSettings.New(context.TODO())
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConsentSettingGet(t *testing.T) {
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
	_, err := client.ConsentSettings.Get(context.TODO(), "id")
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConsentSettingReplaceWithOptionalParams(t *testing.T) {
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
	_, err := client.ConsentSettings.Replace(
		context.TODO(),
		"id",
		oursprivacy.ConsentSettingReplaceParams{
			Categories: []oursprivacy.ConsentSettingReplaceParamsCategory{{
				Label:    "label",
				Priority: 0,
				Value:    "value",
			}},
			Default: oursprivacy.ConsentSettingReplaceParamsDefault{
				Categories: []oursprivacy.ConsentSettingReplaceParamsDefaultCategory{{
					Key: "key",
					Value: oursprivacy.ConsentSettingReplaceParamsDefaultCategoryValue{
						Enabled:          true,
						AutoDisableOnGpc: oursprivacy.Bool(true),
						ReadOnly:         oursprivacy.Bool(true),
						ReloadPage:       oursprivacy.Bool(true),
					},
				}},
				Language: "en",
				Mode:     "opt_in",
				Translations: []oursprivacy.ConsentSettingReplaceParamsDefaultTranslation{{
					Language: "en",
					Value: oursprivacy.ConsentSettingReplaceParamsDefaultTranslationValue{
						ConsentModal:     map[string]any{},
						PreferencesModal: map[string]any{},
					},
				}},
				AutoblockUnknown:         oursprivacy.Bool(true),
				AutoShow:                 oursprivacy.Bool(true),
				AutoShowDismissConfig:    map[string]any{},
				AutoShowDismissMode:      oursprivacy.String("autoShowDismissMode"),
				DisablePageInteraction:   oursprivacy.Bool(true),
				GuiOptions:               map[string]any{},
				HideFromBots:             oursprivacy.Bool(true),
				ShowVendorsInPreferences: oursprivacy.Bool(true),
			},
			Name: "name",
			Regions: []oursprivacy.ConsentSettingReplaceParamsRegion{{
				RegionCode: "US-CA",
				Rule: oursprivacy.ConsentSettingReplaceParamsRegionRule{
					Categories: []oursprivacy.ConsentSettingReplaceParamsRegionRuleCategory{{
						Key: "key",
						Value: oursprivacy.ConsentSettingReplaceParamsRegionRuleCategoryValue{
							Enabled:          true,
							AutoDisableOnGpc: oursprivacy.Bool(true),
							ReadOnly:         oursprivacy.Bool(true),
							ReloadPage:       oursprivacy.Bool(true),
						},
					}},
					Language: "en",
					Mode:     "opt_in",
					Translations: []oursprivacy.ConsentSettingReplaceParamsRegionRuleTranslation{{
						Language: "en",
						Value: oursprivacy.ConsentSettingReplaceParamsRegionRuleTranslationValue{
							ConsentModal:     map[string]any{},
							PreferencesModal: map[string]any{},
						},
					}},
					AutoblockUnknown:         oursprivacy.Bool(true),
					AutoShow:                 oursprivacy.Bool(true),
					AutoShowDismissConfig:    map[string]any{},
					AutoShowDismissMode:      oursprivacy.String("autoShowDismissMode"),
					DisablePageInteraction:   oursprivacy.Bool(true),
					GuiOptions:               map[string]any{},
					HideFromBots:             oursprivacy.Bool(true),
					ShowVendorsInPreferences: oursprivacy.Bool(true),
				},
				AdditionalRegions: []string{"string"},
			}},
			Services: []oursprivacy.ConsentSettingReplaceParamsService{{
				InternalNotes:        "internalNotes",
				Label:                "label",
				AdditionalCategories: []string{"string"},
				Category:             oursprivacy.String("category"),
				DomainPatterns:       []string{"string"},
			}},
			Status:                 oursprivacy.ConsentSettingReplaceParamsStatusDisabled,
			ConsentCookieName:      oursprivacy.String("consentCookieName"),
			CustomDomain:           oursprivacy.String("customDomain"),
			Revision:               oursprivacy.Float(0),
			SkipBlockingClassNames: []string{"string"},
			WebSDKToken:            oursprivacy.String("webSDKToken"),
			WhitelistDomains:       []string{"string"},
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

func TestConsentSettingUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ConsentSettings.Update(
		context.TODO(),
		"id",
		oursprivacy.ConsentSettingUpdateParams{
			Categories: []oursprivacy.ConsentSettingUpdateParamsCategory{{
				Label:    "label",
				Priority: 0,
				Value:    "value",
			}},
			ConsentCookieName: oursprivacy.String("consentCookieName"),
			CustomDomain:      oursprivacy.String("customDomain"),
			Default: oursprivacy.ConsentSettingUpdateParamsDefault{
				Categories: []oursprivacy.ConsentSettingUpdateParamsDefaultCategory{{
					Key: "key",
					Value: oursprivacy.ConsentSettingUpdateParamsDefaultCategoryValue{
						Enabled:          true,
						AutoDisableOnGpc: oursprivacy.Bool(true),
						ReadOnly:         oursprivacy.Bool(true),
						ReloadPage:       oursprivacy.Bool(true),
					},
				}},
				Language: "en",
				Mode:     "opt_in",
				Translations: []oursprivacy.ConsentSettingUpdateParamsDefaultTranslation{{
					Language: "en",
					Value: oursprivacy.ConsentSettingUpdateParamsDefaultTranslationValue{
						ConsentModal:     map[string]any{},
						PreferencesModal: map[string]any{},
					},
				}},
				AutoblockUnknown:         oursprivacy.Bool(true),
				AutoShow:                 oursprivacy.Bool(true),
				AutoShowDismissConfig:    map[string]any{},
				AutoShowDismissMode:      oursprivacy.String("autoShowDismissMode"),
				DisablePageInteraction:   oursprivacy.Bool(true),
				GuiOptions:               map[string]any{},
				HideFromBots:             oursprivacy.Bool(true),
				ShowVendorsInPreferences: oursprivacy.Bool(true),
			},
			Name: oursprivacy.String("name"),
			Regions: []oursprivacy.ConsentSettingUpdateParamsRegion{{
				RegionCode: "US-CA",
				Rule: oursprivacy.ConsentSettingUpdateParamsRegionRule{
					Categories: []oursprivacy.ConsentSettingUpdateParamsRegionRuleCategory{{
						Key: "key",
						Value: oursprivacy.ConsentSettingUpdateParamsRegionRuleCategoryValue{
							Enabled:          true,
							AutoDisableOnGpc: oursprivacy.Bool(true),
							ReadOnly:         oursprivacy.Bool(true),
							ReloadPage:       oursprivacy.Bool(true),
						},
					}},
					Language: "en",
					Mode:     "opt_in",
					Translations: []oursprivacy.ConsentSettingUpdateParamsRegionRuleTranslation{{
						Language: "en",
						Value: oursprivacy.ConsentSettingUpdateParamsRegionRuleTranslationValue{
							ConsentModal:     map[string]any{},
							PreferencesModal: map[string]any{},
						},
					}},
					AutoblockUnknown:         oursprivacy.Bool(true),
					AutoShow:                 oursprivacy.Bool(true),
					AutoShowDismissConfig:    map[string]any{},
					AutoShowDismissMode:      oursprivacy.String("autoShowDismissMode"),
					DisablePageInteraction:   oursprivacy.Bool(true),
					GuiOptions:               map[string]any{},
					HideFromBots:             oursprivacy.Bool(true),
					ShowVendorsInPreferences: oursprivacy.Bool(true),
				},
				AdditionalRegions: []string{"string"},
			}},
			Revision: oursprivacy.Float(0),
			Services: []oursprivacy.ConsentSettingUpdateParamsService{{
				InternalNotes:        "internalNotes",
				Label:                "label",
				AdditionalCategories: []string{"string"},
				Category:             oursprivacy.String("category"),
				DomainPatterns:       []string{"string"},
			}},
			SkipBlockingClassNames: []string{"string"},
			Status:                 oursprivacy.ConsentSettingUpdateParamsStatusDisabled,
			WebSDKToken:            oursprivacy.String("webSDKToken"),
			WhitelistDomains:       []string{"string"},
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

func TestConsentSettingDelete(t *testing.T) {
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
	_, err := client.ConsentSettings.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConsentSettingAnalyticsWithOptionalParams(t *testing.T) {
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
	_, err := client.ConsentSettings.Analytics(
		context.TODO(),
		"id",
		oursprivacy.ConsentSettingAnalyticsParams{
			From:                      "2026-04-01",
			To:                        "2026-04-30",
			CompareWithPreviousPeriod: oursprivacy.Bool(true),
			Granularity:               oursprivacy.ConsentSettingAnalyticsParamsGranularityDaily,
			PagePath:                  oursprivacy.String("/pricing"),
			Regions:                   oursprivacy.String("California"),
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

func TestConsentSettingPageAnalysisWithOptionalParams(t *testing.T) {
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
	_, err := client.ConsentSettings.PageAnalysis(
		context.TODO(),
		"id",
		oursprivacy.ConsentSettingPageAnalysisParams{
			From:    "2026-04-01",
			To:      "2026-04-30",
			Limit:   oursprivacy.Int(1),
			Offset:  oursprivacy.Int(0),
			Regions: oursprivacy.String("California"),
			Search:  oursprivacy.String("/checkout"),
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

func TestConsentSettingAnalyticsByRegion(t *testing.T) {
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
	_, err := client.ConsentSettings.AnalyticsByRegion(
		context.TODO(),
		"id",
		oursprivacy.ConsentSettingAnalyticsByRegionParams{
			From: "2026-04-01",
			To:   "2026-04-30",
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
