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

func TestJourneyFlowListWithOptionalParams(t *testing.T) {
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
	_, err := client.JourneyFlows.List(context.TODO(), oursprivacy.JourneyFlowListParams{
		Cursor: oursprivacy.String("cursor"),
		Limit:  oursprivacy.Int(25),
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

func TestJourneyFlowNewWithOptionalParams(t *testing.T) {
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
	_, err := client.JourneyFlows.New(context.TODO(), oursprivacy.JourneyFlowNewParams{
		Definition: oursprivacy.JourneyFlowNewParamsDefinition{
			Anchors: []oursprivacy.JourneyFlowNewParamsDefinitionAnchor{{
				ID: "x",
				Alternatives: []oursprivacy.JourneyFlowNewParamsDefinitionAnchorAlternativeUnion{{
					OfJourneyFlowNewsDefinitionAnchorAlternativeObject: &oursprivacy.JourneyFlowNewParamsDefinitionAnchorAlternativeObject{
						EventName: "x",
						Kind:      "event",
						Filter: oursprivacy.JourneyFlowNewParamsDefinitionAnchorAlternativeObjectFilter{
							Filter: oursprivacy.JourneyFlowNewParamsDefinitionAnchorAlternativeObjectFilterFilterUnion{
								OfJourneyFlowNewsDefinitionAnchorAlternativeObjectFilterFilterObject: &oursprivacy.JourneyFlowNewParamsDefinitionAnchorAlternativeObjectFilterFilterObject{
									Children: []any{},
									Kind:     "and",
								},
							},
							Version: 1,
						},
					},
				}},
				Filter: oursprivacy.JourneyFlowNewParamsDefinitionAnchorFilter{
					Filter: oursprivacy.JourneyFlowNewParamsDefinitionAnchorFilterFilterUnion{
						OfJourneyFlowNewsDefinitionAnchorFilterFilterObject: &oursprivacy.JourneyFlowNewParamsDefinitionAnchorFilterFilterObject{
							Children: []any{},
							Kind:     "and",
						},
					},
					Version: 1,
				},
				Label: oursprivacy.String("label"),
			}},
			Exploration: oursprivacy.JourneyFlowNewParamsDefinitionExploration{
				After:  oursprivacy.Int(0),
				Before: oursprivacy.Int(0),
				Between: []oursprivacy.JourneyFlowNewParamsDefinitionExplorationBetween{{
					AfterFrom:    0,
					BeforeTo:     0,
					FromAnchorID: "x",
					ToAnchorID:   "x",
				}},
				CollapseRepeats: oursprivacy.Bool(true),
				Focus: []oursprivacy.JourneyFlowNewParamsDefinitionExplorationFocusUnion{{
					OfJourneyFlowNewsDefinitionExplorationFocusObject: &oursprivacy.JourneyFlowNewParamsDefinitionExplorationFocusObject{
						Position: oursprivacy.JourneyFlowNewParamsDefinitionExplorationFocusObjectPosition{
							AnchorID: "x",
							Offset:   0,
							Side:     "anchor",
						},
						Matcher: oursprivacy.JourneyFlowNewParamsDefinitionExplorationFocusObjectMatcherUnion{
							OfJourneyFlowNewsDefinitionExplorationFocusObjectMatcherObject: &oursprivacy.JourneyFlowNewParamsDefinitionExplorationFocusObjectMatcherObject{
								EventName: "x",
								Kind:      "event",
								Filter: oursprivacy.JourneyFlowNewParamsDefinitionExplorationFocusObjectMatcherObjectFilter{
									Filter: oursprivacy.JourneyFlowNewParamsDefinitionExplorationFocusObjectMatcherObjectFilterFilterUnion{
										OfJourneyFlowNewsDefinitionExplorationFocusObjectMatcherObjectFilterFilterObject: &oursprivacy.JourneyFlowNewParamsDefinitionExplorationFocusObjectMatcherObjectFilterFilterObject{
											Children: []any{},
											Kind:     "and",
										},
									},
									Version: 1,
								},
							},
						},
					},
				}},
				HiddenEvents: []oursprivacy.JourneyFlowNewParamsDefinitionExplorationHiddenEventUnion{{
					OfJourneyFlowNewsDefinitionExplorationHiddenEventObject: &oursprivacy.JourneyFlowNewParamsDefinitionExplorationHiddenEventObject{
						EventName: "x",
						Kind:      "event",
						Filter: oursprivacy.JourneyFlowNewParamsDefinitionExplorationHiddenEventObjectFilter{
							Filter: oursprivacy.JourneyFlowNewParamsDefinitionExplorationHiddenEventObjectFilterFilterUnion{
								OfJourneyFlowNewsDefinitionExplorationHiddenEventObjectFilterFilterObject: &oursprivacy.JourneyFlowNewParamsDefinitionExplorationHiddenEventObjectFilterFilterObject{
									Children: []any{},
									Kind:     "and",
								},
							},
							Version: 1,
						},
					},
				}},
				PageKey: "path",
				PathFilter: oursprivacy.JourneyFlowNewParamsDefinitionExplorationPathFilter{
					Filter: oursprivacy.JourneyFlowNewParamsDefinitionExplorationPathFilterFilterUnion{
						OfJourneyFlowNewsDefinitionExplorationPathFilterFilterObject: &oursprivacy.JourneyFlowNewParamsDefinitionExplorationPathFilterFilterObject{
							Children: []any{},
							Kind:     "and",
						},
					},
					Version: 1,
				},
				StepKinds: []string{"event"},
			},
			Kind: "journey-flow",
			Scope: oursprivacy.JourneyFlowNewParamsDefinitionScope{
				ExcludeBots: true,
				SourceIDs:   []string{"x"},
				SessionFilter: oursprivacy.JourneyFlowNewParamsDefinitionScopeSessionFilter{
					Filter: oursprivacy.JourneyFlowNewParamsDefinitionScopeSessionFilterFilterUnion{
						OfJourneyFlowNewsDefinitionScopeSessionFilterFilterObject: &oursprivacy.JourneyFlowNewParamsDefinitionScopeSessionFilterFilterObject{
							Children: []any{},
							Kind:     "and",
						},
					},
					Version: 1,
				},
			},
			Version: 1,
			AnchorFilter: oursprivacy.JourneyFlowNewParamsDefinitionAnchorFilter{
				Filter: oursprivacy.JourneyFlowNewParamsDefinitionAnchorFilterFilterUnion{
					OfJourneyFlowNewsDefinitionAnchorFilterFilterObject: &oursprivacy.JourneyFlowNewParamsDefinitionAnchorFilterFilterObject{
						Children: []any{},
						Kind:     "and",
					},
				},
				Version: 1,
			},
			Breakdown: oursprivacy.JourneyFlowNewParamsDefinitionBreakdown{
				AtAnchorID: "x",
				Property: oursprivacy.JourneyFlowNewParamsDefinitionBreakdownProperty{
					Property: "x",
					Type:     "string",
					Path:     []string{"x"},
				},
				Limit: oursprivacy.Int(1),
			},
			ContextWindowMs:    oursprivacy.Int(1),
			ConversionWindowMs: oursprivacy.Int(1),
			Counting:           "unique",
			EntryFilter: oursprivacy.JourneyFlowNewParamsDefinitionEntryFilter{
				Filter: oursprivacy.JourneyFlowNewParamsDefinitionEntryFilterFilterUnion{
					OfJourneyFlowNewsDefinitionEntryFilterFilterObject: &oursprivacy.JourneyFlowNewParamsDefinitionEntryFilterFilterObject{
						Children: []any{},
						Kind:     "and",
					},
				},
				Version: 1,
			},
			Exclusions: []oursprivacy.JourneyFlowNewParamsDefinitionExclusion{{
				ID:           "x",
				FromAnchorID: "x",
				Matcher: oursprivacy.JourneyFlowNewParamsDefinitionExclusionMatcherUnion{
					OfJourneyFlowNewsDefinitionExclusionMatcherObject: &oursprivacy.JourneyFlowNewParamsDefinitionExclusionMatcherObject{
						EventName: "x",
						Kind:      "event",
						Filter: oursprivacy.JourneyFlowNewParamsDefinitionExclusionMatcherObjectFilter{
							Filter: oursprivacy.JourneyFlowNewParamsDefinitionExclusionMatcherObjectFilterFilterUnion{
								OfJourneyFlowNewsDefinitionExclusionMatcherObjectFilterFilterObject: &oursprivacy.JourneyFlowNewParamsDefinitionExclusionMatcherObjectFilterFilterObject{
									Children: []any{},
									Kind:     "and",
								},
							},
							Version: 1,
						},
					},
				},
				ToAnchorID: "x",
			}},
			HoldConstant: []oursprivacy.JourneyFlowNewParamsDefinitionHoldConstant{{
				Property: "x",
				Type:     "string",
				Path:     []string{"x"},
			}},
			NodeExpansions: []oursprivacy.JourneyFlowNewParamsDefinitionNodeExpansion{{
				Matcher: oursprivacy.JourneyFlowNewParamsDefinitionNodeExpansionMatcherUnion{
					OfJourneyFlowNewsDefinitionNodeExpansionMatcherObject: &oursprivacy.JourneyFlowNewParamsDefinitionNodeExpansionMatcherObject{
						EventName: "x",
						Kind:      "event",
						Filter: oursprivacy.JourneyFlowNewParamsDefinitionNodeExpansionMatcherObjectFilter{
							Filter: oursprivacy.JourneyFlowNewParamsDefinitionNodeExpansionMatcherObjectFilterFilterUnion{
								OfJourneyFlowNewsDefinitionNodeExpansionMatcherObjectFilterFilterObject: &oursprivacy.JourneyFlowNewParamsDefinitionNodeExpansionMatcherObjectFilterFilterObject{
									Children: []any{},
									Kind:     "and",
								},
							},
							Version: 1,
						},
					},
				},
				Position: oursprivacy.JourneyFlowNewParamsDefinitionNodeExpansionPosition{
					AnchorID: "x",
					Offset:   0,
					Side:     "anchor",
				},
				Property: oursprivacy.JourneyFlowNewParamsDefinitionNodeExpansionProperty{
					Property: "x",
					Type:     "string",
					Path:     []string{"x"},
				},
				Limit: oursprivacy.Int(1),
			}},
			Reentry: "first",
		},
		Name: "x",
		ReportContext: oursprivacy.JourneyFlowNewParamsReportContext{
			DateRange: oursprivacy.JourneyFlowNewParamsReportContextDateRange{
				From:     "7321-69-10",
				TimeZone: "x",
				To:       "7321-69-10",
			},
			Comparisons: []oursprivacy.JourneyFlowNewParamsReportContextComparison{{
				ID:    "x",
				Label: "x",
				DateRange: oursprivacy.JourneyFlowNewParamsReportContextComparisonDateRange{
					From:     "7321-69-10",
					TimeZone: "x",
					To:       "7321-69-10",
				},
				EntryFilter: oursprivacy.JourneyFlowNewParamsReportContextComparisonEntryFilter{
					Filter: oursprivacy.JourneyFlowNewParamsReportContextComparisonEntryFilterFilterUnion{
						OfJourneyFlowNewsReportContextComparisonEntryFilterFilterObject: &oursprivacy.JourneyFlowNewParamsReportContextComparisonEntryFilterFilterObject{
							Children: []any{},
							Kind:     "and",
						},
					},
					Version: 1,
				},
				SessionFilter: oursprivacy.JourneyFlowNewParamsReportContextComparisonSessionFilter{
					Filter: oursprivacy.JourneyFlowNewParamsReportContextComparisonSessionFilterFilterUnion{
						OfJourneyFlowNewsReportContextComparisonSessionFilterFilterObject: &oursprivacy.JourneyFlowNewParamsReportContextComparisonSessionFilterFilterObject{
							Children: []any{},
							Kind:     "and",
						},
					},
					Version: 1,
				},
			}},
		},
		View: oursprivacy.JourneyFlowNewParamsView{
			ShowDropOff:          oursprivacy.Bool(true),
			ShowElapsedTime:      oursprivacy.Bool(true),
			TopEventsPerPosition: oursprivacy.Int(1),
			TopPaths:             oursprivacy.Int(1),
			Visualization:        "sankey",
		},
		Description: oursprivacy.String("description"),
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestJourneyFlowGet(t *testing.T) {
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
	_, err := client.JourneyFlows.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestJourneyFlowUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.JourneyFlows.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		oursprivacy.JourneyFlowUpdateParams{
			ExpectedRevision: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Definition: oursprivacy.JourneyFlowUpdateParamsDefinition{
				Anchors: []oursprivacy.JourneyFlowUpdateParamsDefinitionAnchor{{
					ID: "x",
					Alternatives: []oursprivacy.JourneyFlowUpdateParamsDefinitionAnchorAlternativeUnion{{
						OfJourneyFlowUpdatesDefinitionAnchorAlternativeObject: &oursprivacy.JourneyFlowUpdateParamsDefinitionAnchorAlternativeObject{
							EventName: "x",
							Kind:      "event",
							Filter: oursprivacy.JourneyFlowUpdateParamsDefinitionAnchorAlternativeObjectFilter{
								Filter: oursprivacy.JourneyFlowUpdateParamsDefinitionAnchorAlternativeObjectFilterFilterUnion{
									OfJourneyFlowUpdatesDefinitionAnchorAlternativeObjectFilterFilterObject: &oursprivacy.JourneyFlowUpdateParamsDefinitionAnchorAlternativeObjectFilterFilterObject{
										Children: []any{},
										Kind:     "and",
									},
								},
								Version: 1,
							},
						},
					}},
					Filter: oursprivacy.JourneyFlowUpdateParamsDefinitionAnchorFilter{
						Filter: oursprivacy.JourneyFlowUpdateParamsDefinitionAnchorFilterFilterUnion{
							OfJourneyFlowUpdatesDefinitionAnchorFilterFilterObject: &oursprivacy.JourneyFlowUpdateParamsDefinitionAnchorFilterFilterObject{
								Children: []any{},
								Kind:     "and",
							},
						},
						Version: 1,
					},
					Label: oursprivacy.String("label"),
				}},
				Exploration: oursprivacy.JourneyFlowUpdateParamsDefinitionExploration{
					After:  oursprivacy.Int(0),
					Before: oursprivacy.Int(0),
					Between: []oursprivacy.JourneyFlowUpdateParamsDefinitionExplorationBetween{{
						AfterFrom:    0,
						BeforeTo:     0,
						FromAnchorID: "x",
						ToAnchorID:   "x",
					}},
					CollapseRepeats: oursprivacy.Bool(true),
					Focus: []oursprivacy.JourneyFlowUpdateParamsDefinitionExplorationFocusUnion{{
						OfJourneyFlowUpdatesDefinitionExplorationFocusObject: &oursprivacy.JourneyFlowUpdateParamsDefinitionExplorationFocusObject{
							Position: oursprivacy.JourneyFlowUpdateParamsDefinitionExplorationFocusObjectPosition{
								AnchorID: "x",
								Offset:   0,
								Side:     "anchor",
							},
							Matcher: oursprivacy.JourneyFlowUpdateParamsDefinitionExplorationFocusObjectMatcherUnion{
								OfJourneyFlowUpdatesDefinitionExplorationFocusObjectMatcherObject: &oursprivacy.JourneyFlowUpdateParamsDefinitionExplorationFocusObjectMatcherObject{
									EventName: "x",
									Kind:      "event",
									Filter: oursprivacy.JourneyFlowUpdateParamsDefinitionExplorationFocusObjectMatcherObjectFilter{
										Filter: oursprivacy.JourneyFlowUpdateParamsDefinitionExplorationFocusObjectMatcherObjectFilterFilterUnion{
											OfJourneyFlowUpdatesDefinitionExplorationFocusObjectMatcherObjectFilterFilterObject: &oursprivacy.JourneyFlowUpdateParamsDefinitionExplorationFocusObjectMatcherObjectFilterFilterObject{
												Children: []any{},
												Kind:     "and",
											},
										},
										Version: 1,
									},
								},
							},
						},
					}},
					HiddenEvents: []oursprivacy.JourneyFlowUpdateParamsDefinitionExplorationHiddenEventUnion{{
						OfJourneyFlowUpdatesDefinitionExplorationHiddenEventObject: &oursprivacy.JourneyFlowUpdateParamsDefinitionExplorationHiddenEventObject{
							EventName: "x",
							Kind:      "event",
							Filter: oursprivacy.JourneyFlowUpdateParamsDefinitionExplorationHiddenEventObjectFilter{
								Filter: oursprivacy.JourneyFlowUpdateParamsDefinitionExplorationHiddenEventObjectFilterFilterUnion{
									OfJourneyFlowUpdatesDefinitionExplorationHiddenEventObjectFilterFilterObject: &oursprivacy.JourneyFlowUpdateParamsDefinitionExplorationHiddenEventObjectFilterFilterObject{
										Children: []any{},
										Kind:     "and",
									},
								},
								Version: 1,
							},
						},
					}},
					PageKey: "path",
					PathFilter: oursprivacy.JourneyFlowUpdateParamsDefinitionExplorationPathFilter{
						Filter: oursprivacy.JourneyFlowUpdateParamsDefinitionExplorationPathFilterFilterUnion{
							OfJourneyFlowUpdatesDefinitionExplorationPathFilterFilterObject: &oursprivacy.JourneyFlowUpdateParamsDefinitionExplorationPathFilterFilterObject{
								Children: []any{},
								Kind:     "and",
							},
						},
						Version: 1,
					},
					StepKinds: []string{"event"},
				},
				Kind: "journey-flow",
				Scope: oursprivacy.JourneyFlowUpdateParamsDefinitionScope{
					ExcludeBots: true,
					SourceIDs:   []string{"x"},
					SessionFilter: oursprivacy.JourneyFlowUpdateParamsDefinitionScopeSessionFilter{
						Filter: oursprivacy.JourneyFlowUpdateParamsDefinitionScopeSessionFilterFilterUnion{
							OfJourneyFlowUpdatesDefinitionScopeSessionFilterFilterObject: &oursprivacy.JourneyFlowUpdateParamsDefinitionScopeSessionFilterFilterObject{
								Children: []any{},
								Kind:     "and",
							},
						},
						Version: 1,
					},
				},
				Version: 1,
				AnchorFilter: oursprivacy.JourneyFlowUpdateParamsDefinitionAnchorFilter{
					Filter: oursprivacy.JourneyFlowUpdateParamsDefinitionAnchorFilterFilterUnion{
						OfJourneyFlowUpdatesDefinitionAnchorFilterFilterObject: &oursprivacy.JourneyFlowUpdateParamsDefinitionAnchorFilterFilterObject{
							Children: []any{},
							Kind:     "and",
						},
					},
					Version: 1,
				},
				Breakdown: oursprivacy.JourneyFlowUpdateParamsDefinitionBreakdown{
					AtAnchorID: "x",
					Property: oursprivacy.JourneyFlowUpdateParamsDefinitionBreakdownProperty{
						Property: "x",
						Type:     "string",
						Path:     []string{"x"},
					},
					Limit: oursprivacy.Int(1),
				},
				ContextWindowMs:    oursprivacy.Int(1),
				ConversionWindowMs: oursprivacy.Int(1),
				Counting:           "unique",
				EntryFilter: oursprivacy.JourneyFlowUpdateParamsDefinitionEntryFilter{
					Filter: oursprivacy.JourneyFlowUpdateParamsDefinitionEntryFilterFilterUnion{
						OfJourneyFlowUpdatesDefinitionEntryFilterFilterObject: &oursprivacy.JourneyFlowUpdateParamsDefinitionEntryFilterFilterObject{
							Children: []any{},
							Kind:     "and",
						},
					},
					Version: 1,
				},
				Exclusions: []oursprivacy.JourneyFlowUpdateParamsDefinitionExclusion{{
					ID:           "x",
					FromAnchorID: "x",
					Matcher: oursprivacy.JourneyFlowUpdateParamsDefinitionExclusionMatcherUnion{
						OfJourneyFlowUpdatesDefinitionExclusionMatcherObject: &oursprivacy.JourneyFlowUpdateParamsDefinitionExclusionMatcherObject{
							EventName: "x",
							Kind:      "event",
							Filter: oursprivacy.JourneyFlowUpdateParamsDefinitionExclusionMatcherObjectFilter{
								Filter: oursprivacy.JourneyFlowUpdateParamsDefinitionExclusionMatcherObjectFilterFilterUnion{
									OfJourneyFlowUpdatesDefinitionExclusionMatcherObjectFilterFilterObject: &oursprivacy.JourneyFlowUpdateParamsDefinitionExclusionMatcherObjectFilterFilterObject{
										Children: []any{},
										Kind:     "and",
									},
								},
								Version: 1,
							},
						},
					},
					ToAnchorID: "x",
				}},
				HoldConstant: []oursprivacy.JourneyFlowUpdateParamsDefinitionHoldConstant{{
					Property: "x",
					Type:     "string",
					Path:     []string{"x"},
				}},
				NodeExpansions: []oursprivacy.JourneyFlowUpdateParamsDefinitionNodeExpansion{{
					Matcher: oursprivacy.JourneyFlowUpdateParamsDefinitionNodeExpansionMatcherUnion{
						OfJourneyFlowUpdatesDefinitionNodeExpansionMatcherObject: &oursprivacy.JourneyFlowUpdateParamsDefinitionNodeExpansionMatcherObject{
							EventName: "x",
							Kind:      "event",
							Filter: oursprivacy.JourneyFlowUpdateParamsDefinitionNodeExpansionMatcherObjectFilter{
								Filter: oursprivacy.JourneyFlowUpdateParamsDefinitionNodeExpansionMatcherObjectFilterFilterUnion{
									OfJourneyFlowUpdatesDefinitionNodeExpansionMatcherObjectFilterFilterObject: &oursprivacy.JourneyFlowUpdateParamsDefinitionNodeExpansionMatcherObjectFilterFilterObject{
										Children: []any{},
										Kind:     "and",
									},
								},
								Version: 1,
							},
						},
					},
					Position: oursprivacy.JourneyFlowUpdateParamsDefinitionNodeExpansionPosition{
						AnchorID: "x",
						Offset:   0,
						Side:     "anchor",
					},
					Property: oursprivacy.JourneyFlowUpdateParamsDefinitionNodeExpansionProperty{
						Property: "x",
						Type:     "string",
						Path:     []string{"x"},
					},
					Limit: oursprivacy.Int(1),
				}},
				Reentry: "first",
			},
			Description: oursprivacy.String("description"),
			Name:        oursprivacy.String("x"),
			ReportContext: oursprivacy.JourneyFlowUpdateParamsReportContext{
				DateRange: oursprivacy.JourneyFlowUpdateParamsReportContextDateRange{
					From:     "7321-69-10",
					TimeZone: "x",
					To:       "7321-69-10",
				},
				Comparisons: []oursprivacy.JourneyFlowUpdateParamsReportContextComparison{{
					ID:    "x",
					Label: "x",
					DateRange: oursprivacy.JourneyFlowUpdateParamsReportContextComparisonDateRange{
						From:     "7321-69-10",
						TimeZone: "x",
						To:       "7321-69-10",
					},
					EntryFilter: oursprivacy.JourneyFlowUpdateParamsReportContextComparisonEntryFilter{
						Filter: oursprivacy.JourneyFlowUpdateParamsReportContextComparisonEntryFilterFilterUnion{
							OfJourneyFlowUpdatesReportContextComparisonEntryFilterFilterObject: &oursprivacy.JourneyFlowUpdateParamsReportContextComparisonEntryFilterFilterObject{
								Children: []any{},
								Kind:     "and",
							},
						},
						Version: 1,
					},
					SessionFilter: oursprivacy.JourneyFlowUpdateParamsReportContextComparisonSessionFilter{
						Filter: oursprivacy.JourneyFlowUpdateParamsReportContextComparisonSessionFilterFilterUnion{
							OfJourneyFlowUpdatesReportContextComparisonSessionFilterFilterObject: &oursprivacy.JourneyFlowUpdateParamsReportContextComparisonSessionFilterFilterObject{
								Children: []any{},
								Kind:     "and",
							},
						},
						Version: 1,
					},
				}},
			},
			View: oursprivacy.JourneyFlowUpdateParamsView{
				ShowDropOff:          oursprivacy.Bool(true),
				ShowElapsedTime:      oursprivacy.Bool(true),
				TopEventsPerPosition: oursprivacy.Int(1),
				TopPaths:             oursprivacy.Int(1),
				Visualization:        "sankey",
			},
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

func TestJourneyFlowDelete(t *testing.T) {
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
	_, err := client.JourneyFlows.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		oursprivacy.JourneyFlowDeleteParams{
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

func TestJourneyFlowCapabilities(t *testing.T) {
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
	_, err := client.JourneyFlows.Capabilities(context.TODO())
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestJourneyFlowResultsWithOptionalParams(t *testing.T) {
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
	_, err := client.JourneyFlows.Results(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		oursprivacy.JourneyFlowResultsParams{
			ExpectedRevision:      "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Refresh:               oursprivacy.Bool(true),
			ReportContextOverride: oursprivacy.String("reportContextOverride"),
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

func TestJourneyFlowExportWithOptionalParams(t *testing.T) {
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
	_, err := client.JourneyFlows.Export(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		oursprivacy.JourneyFlowExportParams{
			ExpectedRevision:      "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Refresh:               oursprivacy.Bool(true),
			ReportContextOverride: oursprivacy.String("reportContextOverride"),
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

func TestJourneyFlowPreview(t *testing.T) {
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
	_, err := client.JourneyFlows.Preview(context.TODO(), oursprivacy.JourneyFlowPreviewParams{
		Input: "input",
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
