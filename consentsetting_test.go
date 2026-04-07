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

func TestConsentSettingNew(t *testing.T) {
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
	_, err := client.ConsentSettings.New(context.TODO())
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
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
	client := githubcomwithoursplatformsdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ConsentSettings.Get(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
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
	client := githubcomwithoursplatformsdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ConsentSettings.Update(
		context.TODO(),
		"id",
		githubcomwithoursplatformsdkgo.ConsentSettingUpdateParams{
			Categories: []githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsCategory{{
				Label:    "label",
				Priority: 0,
				Value:    "value",
			}},
			Default: githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsDefault{
				Categories: []githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsDefaultCategory{{
					Key: "key",
					Value: githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsDefaultCategoryValue{
						Enabled:          true,
						AutoDisableOnGpc: githubcomwithoursplatformsdkgo.Bool(true),
						ReadOnly:         githubcomwithoursplatformsdkgo.Bool(true),
						ReloadPage:       githubcomwithoursplatformsdkgo.Bool(true),
					},
				}},
				Language: "language",
				Mode:     "opt_in",
				Translations: []githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsDefaultTranslation{{
					Language: "language",
					Value: githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsDefaultTranslationValue{
						ConsentModal:     map[string]any{},
						PreferencesModal: map[string]any{},
					},
				}},
				AutoShow:                 githubcomwithoursplatformsdkgo.Bool(true),
				AutoShowDismissConfig:    map[string]any{},
				AutoShowDismissMode:      githubcomwithoursplatformsdkgo.String("autoShowDismissMode"),
				DisablePageInteraction:   githubcomwithoursplatformsdkgo.Bool(true),
				GuiOptions:               map[string]any{},
				HideFromBots:             githubcomwithoursplatformsdkgo.Bool(true),
				ShowVendorsInPreferences: githubcomwithoursplatformsdkgo.Bool(true),
			},
			Name: "name",
			Regions: []githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsRegion{{
				RegionCode: "regionCode",
				Rule: githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsRegionRule{
					Categories: []githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsRegionRuleCategory{{
						Key: "key",
						Value: githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsRegionRuleCategoryValue{
							Enabled:          true,
							AutoDisableOnGpc: githubcomwithoursplatformsdkgo.Bool(true),
							ReadOnly:         githubcomwithoursplatformsdkgo.Bool(true),
							ReloadPage:       githubcomwithoursplatformsdkgo.Bool(true),
						},
					}},
					Language: "language",
					Mode:     "opt_in",
					Translations: []githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsRegionRuleTranslation{{
						Language: "language",
						Value: githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsRegionRuleTranslationValue{
							ConsentModal:     map[string]any{},
							PreferencesModal: map[string]any{},
						},
					}},
					AutoShow:                 githubcomwithoursplatformsdkgo.Bool(true),
					AutoShowDismissConfig:    map[string]any{},
					AutoShowDismissMode:      githubcomwithoursplatformsdkgo.String("autoShowDismissMode"),
					DisablePageInteraction:   githubcomwithoursplatformsdkgo.Bool(true),
					GuiOptions:               map[string]any{},
					HideFromBots:             githubcomwithoursplatformsdkgo.Bool(true),
					ShowVendorsInPreferences: githubcomwithoursplatformsdkgo.Bool(true),
				},
				AdditionalRegions: []any{map[string]any{}},
			}},
			Services: []githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsService{{
				InternalNotes:        "internalNotes",
				Label:                "label",
				AdditionalCategories: []any{map[string]any{}},
				Category:             githubcomwithoursplatformsdkgo.String("category"),
				DomainPatterns:       []any{map[string]any{}},
			}},
			Status:                 githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsStatusDisabled,
			ConsentCookieName:      githubcomwithoursplatformsdkgo.String("consentCookieName"),
			CustomDomain:           githubcomwithoursplatformsdkgo.String("customDomain"),
			Revision:               githubcomwithoursplatformsdkgo.Float(0),
			SkipBlockingClassNames: []any{map[string]any{}},
			WebSDKToken:            githubcomwithoursplatformsdkgo.String("webSDKToken"),
			WhitelistDomains:       []any{map[string]any{}},
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

func TestConsentSettingList(t *testing.T) {
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
	_, err := client.ConsentSettings.List(context.TODO())
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
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
	client := githubcomwithoursplatformsdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ConsentSettings.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
