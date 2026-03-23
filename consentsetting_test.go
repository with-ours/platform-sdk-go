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
						ConsentModal: githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsDefaultTranslationValueConsentModal{
							AcceptAllBtn:        "acceptAllBtn",
							Description:         "description",
							RejectAllBtn:        "rejectAllBtn",
							ShowPreferencesBtn:  "showPreferencesBtn",
							Title:               "title",
							CloseIconLabel:      githubcomwithoursplatformsdkgo.String("closeIconLabel"),
							Footer:              githubcomwithoursplatformsdkgo.String("footer"),
							GpcNotification:     githubcomwithoursplatformsdkgo.String("gpcNotification"),
							PrivacyPolicyLabel:  githubcomwithoursplatformsdkgo.String("privacyPolicyLabel"),
							PrivacyPolicyURL:    githubcomwithoursplatformsdkgo.String("privacyPolicyUrl"),
							RevisionMessage:     githubcomwithoursplatformsdkgo.String("revisionMessage"),
							TermsOfServiceLabel: githubcomwithoursplatformsdkgo.String("termsOfServiceLabel"),
							TermsOfServiceURL:   githubcomwithoursplatformsdkgo.String("termsOfServiceUrl"),
						},
						PreferencesModal: githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModal{
							AcceptAllBtn:       "acceptAllBtn",
							CloseIconLabel:     "closeIconLabel",
							RejectAllBtn:       "rejectAllBtn",
							SavePreferencesBtn: "savePreferencesBtn",
							Sections: []githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSection{{
								Description: "description",
								Title:       "title",
								CookieTable: githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSectionCookieTable{
									Body: []githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSectionCookieTableBody{{
										Key:   "key",
										Value: "value",
									}},
									Headers: []githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSectionCookieTableHeader{{
										Key:   "key",
										Value: "value",
									}},
									Caption: githubcomwithoursplatformsdkgo.String("caption"),
								},
								LinkedCategory: githubcomwithoursplatformsdkgo.String("linkedCategory"),
							}},
							Title:               "title",
							ServiceCounterLabel: githubcomwithoursplatformsdkgo.String("serviceCounterLabel"),
						},
					},
				}},
				AutoShow: githubcomwithoursplatformsdkgo.Bool(true),
				AutoShowDismissConfig: githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsDefaultAutoShowDismissConfig{
					PageCount: githubcomwithoursplatformsdkgo.Int(0),
					Seconds:   githubcomwithoursplatformsdkgo.Int(0),
				},
				AutoShowDismissMode:    "after_pages",
				DisablePageInteraction: githubcomwithoursplatformsdkgo.Bool(true),
				GuiOptions: githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsDefaultGuiOptions{
					ConsentModal: githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsDefaultGuiOptionsConsentModal{
						ButtonLayout:       "AcceptAllNecessaryPreferences",
						EqualWeightButtons: githubcomwithoursplatformsdkgo.Bool(true),
						FlipButtons:        githubcomwithoursplatformsdkgo.Bool(true),
						Layout:             "bar",
						Position:           "bottom",
						ShowCloseIcon:      githubcomwithoursplatformsdkgo.Bool(true),
					},
					CssVariables: githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsDefaultGuiOptionsCssVariables{
						ButtonBorderRadius:         githubcomwithoursplatformsdkgo.String("buttonBorderRadius"),
						CookieCategoryBlockBg:      githubcomwithoursplatformsdkgo.String("cookieCategoryBlockBg"),
						CookieCategoryBlockHoverBg: githubcomwithoursplatformsdkgo.String("cookieCategoryBlockHoverBg"),
						FooterBg:                   githubcomwithoursplatformsdkgo.String("footerBg"),
						FooterColor:                githubcomwithoursplatformsdkgo.String("footerColor"),
						FooterLinkColor:            githubcomwithoursplatformsdkgo.String("footerLinkColor"),
						FooterLinkHoverColor:       githubcomwithoursplatformsdkgo.String("footerLinkHoverColor"),
						ModalBg:                    githubcomwithoursplatformsdkgo.String("modalBg"),
						ModalBorderRadius:          githubcomwithoursplatformsdkgo.String("modalBorderRadius"),
						PrimaryButtonBg:            githubcomwithoursplatformsdkgo.String("primaryButtonBg"),
						PrimaryButtonColor:         githubcomwithoursplatformsdkgo.String("primaryButtonColor"),
						PrimaryButtonHoverBg:       githubcomwithoursplatformsdkgo.String("primaryButtonHoverBg"),
						PrimaryButtonHoverColor:    githubcomwithoursplatformsdkgo.String("primaryButtonHoverColor"),
						PrimaryTextColor:           githubcomwithoursplatformsdkgo.String("primaryTextColor"),
						SecondaryButtonBg:          githubcomwithoursplatformsdkgo.String("secondaryButtonBg"),
						SecondaryButtonColor:       githubcomwithoursplatformsdkgo.String("secondaryButtonColor"),
						SecondaryButtonHoverBg:     githubcomwithoursplatformsdkgo.String("secondaryButtonHoverBg"),
						SecondaryButtonHoverColor:  githubcomwithoursplatformsdkgo.String("secondaryButtonHoverColor"),
						SecondaryTextColor:         githubcomwithoursplatformsdkgo.String("secondaryTextColor"),
						ToggleOnBg:                 githubcomwithoursplatformsdkgo.String("toggleOnBg"),
					},
					PreferencesModal: githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsDefaultGuiOptionsPreferencesModal{
						ButtonLayout:       "AcceptAllRejectAllSave",
						EqualWeightButtons: githubcomwithoursplatformsdkgo.Bool(true),
						FlipButtons:        githubcomwithoursplatformsdkgo.Bool(true),
						Layout:             "bar",
						Position:           "left",
					},
				},
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
							ConsentModal: githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsRegionRuleTranslationValueConsentModal{
								AcceptAllBtn:        "acceptAllBtn",
								Description:         "description",
								RejectAllBtn:        "rejectAllBtn",
								ShowPreferencesBtn:  "showPreferencesBtn",
								Title:               "title",
								CloseIconLabel:      githubcomwithoursplatformsdkgo.String("closeIconLabel"),
								Footer:              githubcomwithoursplatformsdkgo.String("footer"),
								GpcNotification:     githubcomwithoursplatformsdkgo.String("gpcNotification"),
								PrivacyPolicyLabel:  githubcomwithoursplatformsdkgo.String("privacyPolicyLabel"),
								PrivacyPolicyURL:    githubcomwithoursplatformsdkgo.String("privacyPolicyUrl"),
								RevisionMessage:     githubcomwithoursplatformsdkgo.String("revisionMessage"),
								TermsOfServiceLabel: githubcomwithoursplatformsdkgo.String("termsOfServiceLabel"),
								TermsOfServiceURL:   githubcomwithoursplatformsdkgo.String("termsOfServiceUrl"),
							},
							PreferencesModal: githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModal{
								AcceptAllBtn:       "acceptAllBtn",
								CloseIconLabel:     "closeIconLabel",
								RejectAllBtn:       "rejectAllBtn",
								SavePreferencesBtn: "savePreferencesBtn",
								Sections: []githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSection{{
									Description: "description",
									Title:       "title",
									CookieTable: githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSectionCookieTable{
										Body: []githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSectionCookieTableBody{{
											Key:   "key",
											Value: "value",
										}},
										Headers: []githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader{{
											Key:   "key",
											Value: "value",
										}},
										Caption: githubcomwithoursplatformsdkgo.String("caption"),
									},
									LinkedCategory: githubcomwithoursplatformsdkgo.String("linkedCategory"),
								}},
								Title:               "title",
								ServiceCounterLabel: githubcomwithoursplatformsdkgo.String("serviceCounterLabel"),
							},
						},
					}},
					AutoShow: githubcomwithoursplatformsdkgo.Bool(true),
					AutoShowDismissConfig: githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsRegionRuleAutoShowDismissConfig{
						PageCount: githubcomwithoursplatformsdkgo.Int(0),
						Seconds:   githubcomwithoursplatformsdkgo.Int(0),
					},
					AutoShowDismissMode:    "after_pages",
					DisablePageInteraction: githubcomwithoursplatformsdkgo.Bool(true),
					GuiOptions: githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsRegionRuleGuiOptions{
						ConsentModal: githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsRegionRuleGuiOptionsConsentModal{
							ButtonLayout:       "AcceptAllNecessaryPreferences",
							EqualWeightButtons: githubcomwithoursplatformsdkgo.Bool(true),
							FlipButtons:        githubcomwithoursplatformsdkgo.Bool(true),
							Layout:             "bar",
							Position:           "bottom",
							ShowCloseIcon:      githubcomwithoursplatformsdkgo.Bool(true),
						},
						CssVariables: githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsRegionRuleGuiOptionsCssVariables{
							ButtonBorderRadius:         githubcomwithoursplatformsdkgo.String("buttonBorderRadius"),
							CookieCategoryBlockBg:      githubcomwithoursplatformsdkgo.String("cookieCategoryBlockBg"),
							CookieCategoryBlockHoverBg: githubcomwithoursplatformsdkgo.String("cookieCategoryBlockHoverBg"),
							FooterBg:                   githubcomwithoursplatformsdkgo.String("footerBg"),
							FooterColor:                githubcomwithoursplatformsdkgo.String("footerColor"),
							FooterLinkColor:            githubcomwithoursplatformsdkgo.String("footerLinkColor"),
							FooterLinkHoverColor:       githubcomwithoursplatformsdkgo.String("footerLinkHoverColor"),
							ModalBg:                    githubcomwithoursplatformsdkgo.String("modalBg"),
							ModalBorderRadius:          githubcomwithoursplatformsdkgo.String("modalBorderRadius"),
							PrimaryButtonBg:            githubcomwithoursplatformsdkgo.String("primaryButtonBg"),
							PrimaryButtonColor:         githubcomwithoursplatformsdkgo.String("primaryButtonColor"),
							PrimaryButtonHoverBg:       githubcomwithoursplatformsdkgo.String("primaryButtonHoverBg"),
							PrimaryButtonHoverColor:    githubcomwithoursplatformsdkgo.String("primaryButtonHoverColor"),
							PrimaryTextColor:           githubcomwithoursplatformsdkgo.String("primaryTextColor"),
							SecondaryButtonBg:          githubcomwithoursplatformsdkgo.String("secondaryButtonBg"),
							SecondaryButtonColor:       githubcomwithoursplatformsdkgo.String("secondaryButtonColor"),
							SecondaryButtonHoverBg:     githubcomwithoursplatformsdkgo.String("secondaryButtonHoverBg"),
							SecondaryButtonHoverColor:  githubcomwithoursplatformsdkgo.String("secondaryButtonHoverColor"),
							SecondaryTextColor:         githubcomwithoursplatformsdkgo.String("secondaryTextColor"),
							ToggleOnBg:                 githubcomwithoursplatformsdkgo.String("toggleOnBg"),
						},
						PreferencesModal: githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsRegionRuleGuiOptionsPreferencesModal{
							ButtonLayout:       "AcceptAllRejectAllSave",
							EqualWeightButtons: githubcomwithoursplatformsdkgo.Bool(true),
							FlipButtons:        githubcomwithoursplatformsdkgo.Bool(true),
							Layout:             "bar",
							Position:           "left",
						},
					},
					HideFromBots:             githubcomwithoursplatformsdkgo.Bool(true),
					ShowVendorsInPreferences: githubcomwithoursplatformsdkgo.Bool(true),
				},
				AdditionalRegions: []string{"string"},
			}},
			Services: []githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsService{{
				InternalNotes:        "internalNotes",
				Label:                "label",
				AdditionalCategories: []string{"string"},
				Category:             githubcomwithoursplatformsdkgo.String("category"),
				DomainPatterns:       []string{"string"},
			}},
			Status:                 githubcomwithoursplatformsdkgo.ConsentSettingUpdateParamsStatusDisabled,
			ConsentCookieName:      githubcomwithoursplatformsdkgo.String("consentCookieName"),
			CustomDomain:           githubcomwithoursplatformsdkgo.String("customDomain"),
			Revision:               githubcomwithoursplatformsdkgo.Float(0),
			SkipBlockingClassNames: []string{"string"},
			WebSDKToken:            githubcomwithoursplatformsdkgo.String("webSDKToken"),
			WhitelistDomains:       []string{"string"},
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
