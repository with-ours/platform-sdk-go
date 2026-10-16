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
						ConsentModal: oursprivacy.ConsentSettingReplaceParamsDefaultTranslationValueConsentModal{
							AcceptAllBtn:        "acceptAllBtn",
							Description:         "description",
							RejectAllBtn:        "rejectAllBtn",
							ShowPreferencesBtn:  "showPreferencesBtn",
							Title:               "title",
							CloseIconLabel:      oursprivacy.String("closeIconLabel"),
							Footer:              oursprivacy.String("footer"),
							GpcNotification:     oursprivacy.String("gpcNotification"),
							PrivacyPolicyLabel:  oursprivacy.String("privacyPolicyLabel"),
							PrivacyPolicyURL:    oursprivacy.String("privacyPolicyUrl"),
							RevisionMessage:     oursprivacy.String("revisionMessage"),
							TermsOfServiceLabel: oursprivacy.String("termsOfServiceLabel"),
							TermsOfServiceURL:   oursprivacy.String("termsOfServiceUrl"),
						},
						PreferencesModal: oursprivacy.ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModal{
							AcceptAllBtn:       "acceptAllBtn",
							CloseIconLabel:     "closeIconLabel",
							RejectAllBtn:       "rejectAllBtn",
							SavePreferencesBtn: "savePreferencesBtn",
							Sections: []oursprivacy.ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModalSection{{
								Description: "description",
								Title:       "title",
								CookieTable: oursprivacy.ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModalSectionCookieTable{
									Body: []oursprivacy.ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModalSectionCookieTableBody{{
										Key:   "key",
										Value: "value",
									}},
									Headers: []oursprivacy.ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModalSectionCookieTableHeader{{
										Key:   "key",
										Value: "value",
									}},
									Caption: oursprivacy.String("caption"),
								},
								LinkedCategory: oursprivacy.String("linkedCategory"),
							}},
							Title: "title",
							AdditionalInfo: oursprivacy.ConsentSettingReplaceParamsDefaultTranslationValuePreferencesModalAdditionalInfo{
								Body:  "body",
								Title: "title",
							},
							ServiceCounterLabel: oursprivacy.String("serviceCounterLabel"),
						},
					},
				}},
				AutoblockUnknown: oursprivacy.Bool(true),
				AutoShow:         oursprivacy.Bool(true),
				AutoShowDismissConfig: oursprivacy.ConsentSettingReplaceParamsDefaultAutoShowDismissConfig{
					PageCount: oursprivacy.Int(0),
					Seconds:   oursprivacy.Int(0),
				},
				AutoShowDismissMode:    "after_pages",
				DisablePageInteraction: oursprivacy.Bool(true),
				GuiOptions: oursprivacy.ConsentSettingReplaceParamsDefaultGuiOptions{
					ConsentModal: oursprivacy.ConsentSettingReplaceParamsDefaultGuiOptionsConsentModal{
						ButtonLayout:       "AcceptAllNecessaryPreferences",
						EqualWeightButtons: oursprivacy.Bool(true),
						FlipButtons:        oursprivacy.Bool(true),
						Layout:             "bar",
						Position:           "bottom",
						ShowCloseIcon:      oursprivacy.Bool(true),
					},
					CssVariables: oursprivacy.ConsentSettingReplaceParamsDefaultGuiOptionsCssVariables{
						ButtonBorderRadius:         oursprivacy.String("buttonBorderRadius"),
						CookieCategoryBlockBg:      oursprivacy.String("cookieCategoryBlockBg"),
						CookieCategoryBlockHoverBg: oursprivacy.String("cookieCategoryBlockHoverBg"),
						FooterBg:                   oursprivacy.String("footerBg"),
						FooterColor:                oursprivacy.String("footerColor"),
						FooterLinkColor:            oursprivacy.String("footerLinkColor"),
						FooterLinkHoverColor:       oursprivacy.String("footerLinkHoverColor"),
						ModalBg:                    oursprivacy.String("modalBg"),
						ModalBorderRadius:          oursprivacy.String("modalBorderRadius"),
						PrimaryButtonBg:            oursprivacy.String("primaryButtonBg"),
						PrimaryButtonColor:         oursprivacy.String("primaryButtonColor"),
						PrimaryButtonHoverBg:       oursprivacy.String("primaryButtonHoverBg"),
						PrimaryButtonHoverColor:    oursprivacy.String("primaryButtonHoverColor"),
						PrimaryTextColor:           oursprivacy.String("primaryTextColor"),
						SecondaryButtonBg:          oursprivacy.String("secondaryButtonBg"),
						SecondaryButtonColor:       oursprivacy.String("secondaryButtonColor"),
						SecondaryButtonHoverBg:     oursprivacy.String("secondaryButtonHoverBg"),
						SecondaryButtonHoverColor:  oursprivacy.String("secondaryButtonHoverColor"),
						SecondaryTextColor:         oursprivacy.String("secondaryTextColor"),
						ToggleOnBg:                 oursprivacy.String("toggleOnBg"),
					},
					PreferencesModal: oursprivacy.ConsentSettingReplaceParamsDefaultGuiOptionsPreferencesModal{
						ButtonLayout:       "AcceptAllRejectAllSave",
						EqualWeightButtons: oursprivacy.Bool(true),
						FlipButtons:        oursprivacy.Bool(true),
						Layout:             "bar",
						Position:           "left",
					},
				},
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
							ConsentModal: oursprivacy.ConsentSettingReplaceParamsRegionRuleTranslationValueConsentModal{
								AcceptAllBtn:        "acceptAllBtn",
								Description:         "description",
								RejectAllBtn:        "rejectAllBtn",
								ShowPreferencesBtn:  "showPreferencesBtn",
								Title:               "title",
								CloseIconLabel:      oursprivacy.String("closeIconLabel"),
								Footer:              oursprivacy.String("footer"),
								GpcNotification:     oursprivacy.String("gpcNotification"),
								PrivacyPolicyLabel:  oursprivacy.String("privacyPolicyLabel"),
								PrivacyPolicyURL:    oursprivacy.String("privacyPolicyUrl"),
								RevisionMessage:     oursprivacy.String("revisionMessage"),
								TermsOfServiceLabel: oursprivacy.String("termsOfServiceLabel"),
								TermsOfServiceURL:   oursprivacy.String("termsOfServiceUrl"),
							},
							PreferencesModal: oursprivacy.ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModal{
								AcceptAllBtn:       "acceptAllBtn",
								CloseIconLabel:     "closeIconLabel",
								RejectAllBtn:       "rejectAllBtn",
								SavePreferencesBtn: "savePreferencesBtn",
								Sections: []oursprivacy.ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModalSection{{
									Description: "description",
									Title:       "title",
									CookieTable: oursprivacy.ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModalSectionCookieTable{
										Body: []oursprivacy.ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModalSectionCookieTableBody{{
											Key:   "key",
											Value: "value",
										}},
										Headers: []oursprivacy.ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader{{
											Key:   "key",
											Value: "value",
										}},
										Caption: oursprivacy.String("caption"),
									},
									LinkedCategory: oursprivacy.String("linkedCategory"),
								}},
								Title: "title",
								AdditionalInfo: oursprivacy.ConsentSettingReplaceParamsRegionRuleTranslationValuePreferencesModalAdditionalInfo{
									Body:  "body",
									Title: "title",
								},
								ServiceCounterLabel: oursprivacy.String("serviceCounterLabel"),
							},
						},
					}},
					AutoblockUnknown: oursprivacy.Bool(true),
					AutoShow:         oursprivacy.Bool(true),
					AutoShowDismissConfig: oursprivacy.ConsentSettingReplaceParamsRegionRuleAutoShowDismissConfig{
						PageCount: oursprivacy.Int(0),
						Seconds:   oursprivacy.Int(0),
					},
					AutoShowDismissMode:    "after_pages",
					DisablePageInteraction: oursprivacy.Bool(true),
					GuiOptions: oursprivacy.ConsentSettingReplaceParamsRegionRuleGuiOptions{
						ConsentModal: oursprivacy.ConsentSettingReplaceParamsRegionRuleGuiOptionsConsentModal{
							ButtonLayout:       "AcceptAllNecessaryPreferences",
							EqualWeightButtons: oursprivacy.Bool(true),
							FlipButtons:        oursprivacy.Bool(true),
							Layout:             "bar",
							Position:           "bottom",
							ShowCloseIcon:      oursprivacy.Bool(true),
						},
						CssVariables: oursprivacy.ConsentSettingReplaceParamsRegionRuleGuiOptionsCssVariables{
							ButtonBorderRadius:         oursprivacy.String("buttonBorderRadius"),
							CookieCategoryBlockBg:      oursprivacy.String("cookieCategoryBlockBg"),
							CookieCategoryBlockHoverBg: oursprivacy.String("cookieCategoryBlockHoverBg"),
							FooterBg:                   oursprivacy.String("footerBg"),
							FooterColor:                oursprivacy.String("footerColor"),
							FooterLinkColor:            oursprivacy.String("footerLinkColor"),
							FooterLinkHoverColor:       oursprivacy.String("footerLinkHoverColor"),
							ModalBg:                    oursprivacy.String("modalBg"),
							ModalBorderRadius:          oursprivacy.String("modalBorderRadius"),
							PrimaryButtonBg:            oursprivacy.String("primaryButtonBg"),
							PrimaryButtonColor:         oursprivacy.String("primaryButtonColor"),
							PrimaryButtonHoverBg:       oursprivacy.String("primaryButtonHoverBg"),
							PrimaryButtonHoverColor:    oursprivacy.String("primaryButtonHoverColor"),
							PrimaryTextColor:           oursprivacy.String("primaryTextColor"),
							SecondaryButtonBg:          oursprivacy.String("secondaryButtonBg"),
							SecondaryButtonColor:       oursprivacy.String("secondaryButtonColor"),
							SecondaryButtonHoverBg:     oursprivacy.String("secondaryButtonHoverBg"),
							SecondaryButtonHoverColor:  oursprivacy.String("secondaryButtonHoverColor"),
							SecondaryTextColor:         oursprivacy.String("secondaryTextColor"),
							ToggleOnBg:                 oursprivacy.String("toggleOnBg"),
						},
						PreferencesModal: oursprivacy.ConsentSettingReplaceParamsRegionRuleGuiOptionsPreferencesModal{
							ButtonLayout:       "AcceptAllRejectAllSave",
							EqualWeightButtons: oursprivacy.Bool(true),
							FlipButtons:        oursprivacy.Bool(true),
							Layout:             "bar",
							Position:           "left",
						},
					},
					HideFromBots:             oursprivacy.Bool(true),
					ShowVendorsInPreferences: oursprivacy.Bool(true),
				},
				AdditionalRegions: []string{"string"},
				Name:              oursprivacy.String("name"),
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
			DeviceIDCookieName:     oursprivacy.String("deviceIdCookieName"),
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
						ConsentModal: oursprivacy.ConsentSettingUpdateParamsDefaultTranslationValueConsentModal{
							AcceptAllBtn:        "acceptAllBtn",
							Description:         "description",
							RejectAllBtn:        "rejectAllBtn",
							ShowPreferencesBtn:  "showPreferencesBtn",
							Title:               "title",
							CloseIconLabel:      oursprivacy.String("closeIconLabel"),
							Footer:              oursprivacy.String("footer"),
							GpcNotification:     oursprivacy.String("gpcNotification"),
							PrivacyPolicyLabel:  oursprivacy.String("privacyPolicyLabel"),
							PrivacyPolicyURL:    oursprivacy.String("privacyPolicyUrl"),
							RevisionMessage:     oursprivacy.String("revisionMessage"),
							TermsOfServiceLabel: oursprivacy.String("termsOfServiceLabel"),
							TermsOfServiceURL:   oursprivacy.String("termsOfServiceUrl"),
						},
						PreferencesModal: oursprivacy.ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModal{
							AcceptAllBtn:       "acceptAllBtn",
							CloseIconLabel:     "closeIconLabel",
							RejectAllBtn:       "rejectAllBtn",
							SavePreferencesBtn: "savePreferencesBtn",
							Sections: []oursprivacy.ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSection{{
								Description: "description",
								Title:       "title",
								CookieTable: oursprivacy.ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSectionCookieTable{
									Body: []oursprivacy.ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSectionCookieTableBody{{
										Key:   "key",
										Value: "value",
									}},
									Headers: []oursprivacy.ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalSectionCookieTableHeader{{
										Key:   "key",
										Value: "value",
									}},
									Caption: oursprivacy.String("caption"),
								},
								LinkedCategory: oursprivacy.String("linkedCategory"),
							}},
							Title: "title",
							AdditionalInfo: oursprivacy.ConsentSettingUpdateParamsDefaultTranslationValuePreferencesModalAdditionalInfo{
								Body:  "body",
								Title: "title",
							},
							ServiceCounterLabel: oursprivacy.String("serviceCounterLabel"),
						},
					},
				}},
				AutoblockUnknown: oursprivacy.Bool(true),
				AutoShow:         oursprivacy.Bool(true),
				AutoShowDismissConfig: oursprivacy.ConsentSettingUpdateParamsDefaultAutoShowDismissConfig{
					PageCount: oursprivacy.Int(0),
					Seconds:   oursprivacy.Int(0),
				},
				AutoShowDismissMode:    "after_pages",
				DisablePageInteraction: oursprivacy.Bool(true),
				GuiOptions: oursprivacy.ConsentSettingUpdateParamsDefaultGuiOptions{
					ConsentModal: oursprivacy.ConsentSettingUpdateParamsDefaultGuiOptionsConsentModal{
						ButtonLayout:       "AcceptAllNecessaryPreferences",
						EqualWeightButtons: oursprivacy.Bool(true),
						FlipButtons:        oursprivacy.Bool(true),
						Layout:             "bar",
						Position:           "bottom",
						ShowCloseIcon:      oursprivacy.Bool(true),
					},
					CssVariables: oursprivacy.ConsentSettingUpdateParamsDefaultGuiOptionsCssVariables{
						ButtonBorderRadius:         oursprivacy.String("buttonBorderRadius"),
						CookieCategoryBlockBg:      oursprivacy.String("cookieCategoryBlockBg"),
						CookieCategoryBlockHoverBg: oursprivacy.String("cookieCategoryBlockHoverBg"),
						FooterBg:                   oursprivacy.String("footerBg"),
						FooterColor:                oursprivacy.String("footerColor"),
						FooterLinkColor:            oursprivacy.String("footerLinkColor"),
						FooterLinkHoverColor:       oursprivacy.String("footerLinkHoverColor"),
						ModalBg:                    oursprivacy.String("modalBg"),
						ModalBorderRadius:          oursprivacy.String("modalBorderRadius"),
						PrimaryButtonBg:            oursprivacy.String("primaryButtonBg"),
						PrimaryButtonColor:         oursprivacy.String("primaryButtonColor"),
						PrimaryButtonHoverBg:       oursprivacy.String("primaryButtonHoverBg"),
						PrimaryButtonHoverColor:    oursprivacy.String("primaryButtonHoverColor"),
						PrimaryTextColor:           oursprivacy.String("primaryTextColor"),
						SecondaryButtonBg:          oursprivacy.String("secondaryButtonBg"),
						SecondaryButtonColor:       oursprivacy.String("secondaryButtonColor"),
						SecondaryButtonHoverBg:     oursprivacy.String("secondaryButtonHoverBg"),
						SecondaryButtonHoverColor:  oursprivacy.String("secondaryButtonHoverColor"),
						SecondaryTextColor:         oursprivacy.String("secondaryTextColor"),
						ToggleOnBg:                 oursprivacy.String("toggleOnBg"),
					},
					PreferencesModal: oursprivacy.ConsentSettingUpdateParamsDefaultGuiOptionsPreferencesModal{
						ButtonLayout:       "AcceptAllRejectAllSave",
						EqualWeightButtons: oursprivacy.Bool(true),
						FlipButtons:        oursprivacy.Bool(true),
						Layout:             "bar",
						Position:           "left",
					},
				},
				HideFromBots:             oursprivacy.Bool(true),
				ShowVendorsInPreferences: oursprivacy.Bool(true),
			},
			DeviceIDCookieName: oursprivacy.String("deviceIdCookieName"),
			Name:               oursprivacy.String("name"),
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
							ConsentModal: oursprivacy.ConsentSettingUpdateParamsRegionRuleTranslationValueConsentModal{
								AcceptAllBtn:        "acceptAllBtn",
								Description:         "description",
								RejectAllBtn:        "rejectAllBtn",
								ShowPreferencesBtn:  "showPreferencesBtn",
								Title:               "title",
								CloseIconLabel:      oursprivacy.String("closeIconLabel"),
								Footer:              oursprivacy.String("footer"),
								GpcNotification:     oursprivacy.String("gpcNotification"),
								PrivacyPolicyLabel:  oursprivacy.String("privacyPolicyLabel"),
								PrivacyPolicyURL:    oursprivacy.String("privacyPolicyUrl"),
								RevisionMessage:     oursprivacy.String("revisionMessage"),
								TermsOfServiceLabel: oursprivacy.String("termsOfServiceLabel"),
								TermsOfServiceURL:   oursprivacy.String("termsOfServiceUrl"),
							},
							PreferencesModal: oursprivacy.ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModal{
								AcceptAllBtn:       "acceptAllBtn",
								CloseIconLabel:     "closeIconLabel",
								RejectAllBtn:       "rejectAllBtn",
								SavePreferencesBtn: "savePreferencesBtn",
								Sections: []oursprivacy.ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSection{{
									Description: "description",
									Title:       "title",
									CookieTable: oursprivacy.ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSectionCookieTable{
										Body: []oursprivacy.ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSectionCookieTableBody{{
											Key:   "key",
											Value: "value",
										}},
										Headers: []oursprivacy.ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalSectionCookieTableHeader{{
											Key:   "key",
											Value: "value",
										}},
										Caption: oursprivacy.String("caption"),
									},
									LinkedCategory: oursprivacy.String("linkedCategory"),
								}},
								Title: "title",
								AdditionalInfo: oursprivacy.ConsentSettingUpdateParamsRegionRuleTranslationValuePreferencesModalAdditionalInfo{
									Body:  "body",
									Title: "title",
								},
								ServiceCounterLabel: oursprivacy.String("serviceCounterLabel"),
							},
						},
					}},
					AutoblockUnknown: oursprivacy.Bool(true),
					AutoShow:         oursprivacy.Bool(true),
					AutoShowDismissConfig: oursprivacy.ConsentSettingUpdateParamsRegionRuleAutoShowDismissConfig{
						PageCount: oursprivacy.Int(0),
						Seconds:   oursprivacy.Int(0),
					},
					AutoShowDismissMode:    "after_pages",
					DisablePageInteraction: oursprivacy.Bool(true),
					GuiOptions: oursprivacy.ConsentSettingUpdateParamsRegionRuleGuiOptions{
						ConsentModal: oursprivacy.ConsentSettingUpdateParamsRegionRuleGuiOptionsConsentModal{
							ButtonLayout:       "AcceptAllNecessaryPreferences",
							EqualWeightButtons: oursprivacy.Bool(true),
							FlipButtons:        oursprivacy.Bool(true),
							Layout:             "bar",
							Position:           "bottom",
							ShowCloseIcon:      oursprivacy.Bool(true),
						},
						CssVariables: oursprivacy.ConsentSettingUpdateParamsRegionRuleGuiOptionsCssVariables{
							ButtonBorderRadius:         oursprivacy.String("buttonBorderRadius"),
							CookieCategoryBlockBg:      oursprivacy.String("cookieCategoryBlockBg"),
							CookieCategoryBlockHoverBg: oursprivacy.String("cookieCategoryBlockHoverBg"),
							FooterBg:                   oursprivacy.String("footerBg"),
							FooterColor:                oursprivacy.String("footerColor"),
							FooterLinkColor:            oursprivacy.String("footerLinkColor"),
							FooterLinkHoverColor:       oursprivacy.String("footerLinkHoverColor"),
							ModalBg:                    oursprivacy.String("modalBg"),
							ModalBorderRadius:          oursprivacy.String("modalBorderRadius"),
							PrimaryButtonBg:            oursprivacy.String("primaryButtonBg"),
							PrimaryButtonColor:         oursprivacy.String("primaryButtonColor"),
							PrimaryButtonHoverBg:       oursprivacy.String("primaryButtonHoverBg"),
							PrimaryButtonHoverColor:    oursprivacy.String("primaryButtonHoverColor"),
							PrimaryTextColor:           oursprivacy.String("primaryTextColor"),
							SecondaryButtonBg:          oursprivacy.String("secondaryButtonBg"),
							SecondaryButtonColor:       oursprivacy.String("secondaryButtonColor"),
							SecondaryButtonHoverBg:     oursprivacy.String("secondaryButtonHoverBg"),
							SecondaryButtonHoverColor:  oursprivacy.String("secondaryButtonHoverColor"),
							SecondaryTextColor:         oursprivacy.String("secondaryTextColor"),
							ToggleOnBg:                 oursprivacy.String("toggleOnBg"),
						},
						PreferencesModal: oursprivacy.ConsentSettingUpdateParamsRegionRuleGuiOptionsPreferencesModal{
							ButtonLayout:       "AcceptAllRejectAllSave",
							EqualWeightButtons: oursprivacy.Bool(true),
							FlipButtons:        oursprivacy.Bool(true),
							Layout:             "bar",
							Position:           "left",
						},
					},
					HideFromBots:             oursprivacy.Bool(true),
					ShowVendorsInPreferences: oursprivacy.Bool(true),
				},
				AdditionalRegions: []string{"string"},
				Name:              oursprivacy.String("name"),
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
			ComparisonFrom:            oursprivacy.String("2026-03-01"),
			ComparisonTo:              oursprivacy.String("2026-03-14"),
			Filter:                    oursprivacy.String(`{"version":1,"filter":{"kind":"predicate","property":"event.source_id","operator":"eq","value":"source-1"}}`),
			Granularity:               oursprivacy.ConsentSettingAnalyticsParamsGranularityDaily,
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

func TestConsentSettingAnalyticsMonthlyWithOptionalParams(t *testing.T) {
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
	_, err := client.ConsentSettings.AnalyticsMonthly(
		context.TODO(),
		"id",
		oursprivacy.ConsentSettingAnalyticsMonthlyParams{
			From:   "2026-04-01",
			To:     "2026-04-30",
			Filter: oursprivacy.String(`{"version":1,"filter":{"kind":"predicate","property":"event.source_id","operator":"eq","value":"source-1"}}`),
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

func TestConsentSettingAnalyticsCapabilitiesWithOptionalParams(t *testing.T) {
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
	_, err := client.ConsentSettings.AnalyticsCapabilities(context.TODO(), oursprivacy.ConsentSettingAnalyticsCapabilitiesParams{
		Kind: oursprivacy.ConsentSettingAnalyticsCapabilitiesParamsKindPages,
	})
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
			From:        "2026-04-01",
			To:          "2026-04-30",
			Cursor:      oursprivacy.String("cursor"),
			Filter:      oursprivacy.String(`{"version":1,"filter":{"kind":"predicate","property":"event.source_id","operator":"eq","value":"source-1"}}`),
			Granularity: oursprivacy.ConsentSettingPageAnalysisParamsGranularityDaily,
			Limit:       oursprivacy.Int(0),
			Search:      oursprivacy.String("/checkout"),
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

func TestConsentSettingAnalyticsByRegionWithOptionalParams(t *testing.T) {
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
			From:        "2026-04-01",
			To:          "2026-04-30",
			Filter:      oursprivacy.String(`{"version":1,"filter":{"kind":"predicate","property":"event.source_id","operator":"eq","value":"source-1"}}`),
			Granularity: oursprivacy.ConsentSettingAnalyticsByRegionParamsGranularityDaily,
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
