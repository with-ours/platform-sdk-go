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

func TestGlobalDispatchCenterNew(t *testing.T) {
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
	_, err := client.GlobalDispatchCenters.New(context.TODO())
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGlobalDispatchCenterGet(t *testing.T) {
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
	_, err := client.GlobalDispatchCenters.Get(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGlobalDispatchCenterUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.GlobalDispatchCenters.Update(
		context.TODO(),
		"id",
		githubcomwithoursplatformsdkgo.GlobalDispatchCenterUpdateParams{
			Categories: []githubcomwithoursplatformsdkgo.GlobalDispatchCenterUpdateParamsCategory{{
				Description:    githubcomwithoursplatformsdkgo.String("description"),
				DestinationIDs: []string{"string"},
				Logic: githubcomwithoursplatformsdkgo.GlobalDispatchCenterUpdateParamsCategoryLogic{
					And: []githubcomwithoursplatformsdkgo.GlobalDispatchCenterUpdateParamsCategoryLogicAnd{{
						And: []githubcomwithoursplatformsdkgo.GlobalDispatchCenterUpdateParamsCategoryLogicAndAnd{{
							Condition: githubcomwithoursplatformsdkgo.GlobalDispatchCenterUpdateParamsCategoryLogicAndAndCondition{
								Operator: "Contains",
								Property: "property",
								Value:    "value",
							},
						}},
						Condition: githubcomwithoursplatformsdkgo.GlobalDispatchCenterUpdateParamsCategoryLogicAndCondition{
							Operator: "Contains",
							Property: "property",
							Value:    "value",
						},
						Not: githubcomwithoursplatformsdkgo.GlobalDispatchCenterUpdateParamsCategoryLogicAndNot{
							Condition: githubcomwithoursplatformsdkgo.GlobalDispatchCenterUpdateParamsCategoryLogicAndNotCondition{
								Operator: "Contains",
								Property: "property",
								Value:    "value",
							},
						},
						Or: []githubcomwithoursplatformsdkgo.GlobalDispatchCenterUpdateParamsCategoryLogicAndOr{{
							Condition: githubcomwithoursplatformsdkgo.GlobalDispatchCenterUpdateParamsCategoryLogicAndOrCondition{
								Operator: "Contains",
								Property: "property",
								Value:    "value",
							},
						}},
					}},
					Condition: githubcomwithoursplatformsdkgo.GlobalDispatchCenterUpdateParamsCategoryLogicCondition{
						Operator: "Contains",
						Property: "property",
						Value:    "value",
					},
					Not: githubcomwithoursplatformsdkgo.GlobalDispatchCenterUpdateParamsCategoryLogicNot{
						And: []githubcomwithoursplatformsdkgo.GlobalDispatchCenterUpdateParamsCategoryLogicNotAnd{{
							Condition: githubcomwithoursplatformsdkgo.GlobalDispatchCenterUpdateParamsCategoryLogicNotAndCondition{
								Operator: "Contains",
								Property: "property",
								Value:    "value",
							},
						}},
						Condition: githubcomwithoursplatformsdkgo.GlobalDispatchCenterUpdateParamsCategoryLogicNotCondition{
							Operator: "Contains",
							Property: "property",
							Value:    "value",
						},
						Not: githubcomwithoursplatformsdkgo.GlobalDispatchCenterUpdateParamsCategoryLogicNotNot{
							Condition: githubcomwithoursplatformsdkgo.GlobalDispatchCenterUpdateParamsCategoryLogicNotNotCondition{
								Operator: "Contains",
								Property: "property",
								Value:    "value",
							},
						},
						Or: []githubcomwithoursplatformsdkgo.GlobalDispatchCenterUpdateParamsCategoryLogicNotOr{{
							Condition: githubcomwithoursplatformsdkgo.GlobalDispatchCenterUpdateParamsCategoryLogicNotOrCondition{
								Operator: "Contains",
								Property: "property",
								Value:    "value",
							},
						}},
					},
					Or: []githubcomwithoursplatformsdkgo.GlobalDispatchCenterUpdateParamsCategoryLogicOr{{
						And: []githubcomwithoursplatformsdkgo.GlobalDispatchCenterUpdateParamsCategoryLogicOrAnd{{
							Condition: githubcomwithoursplatformsdkgo.GlobalDispatchCenterUpdateParamsCategoryLogicOrAndCondition{
								Operator: "Contains",
								Property: "property",
								Value:    "value",
							},
						}},
						Condition: githubcomwithoursplatformsdkgo.GlobalDispatchCenterUpdateParamsCategoryLogicOrCondition{
							Operator: "Contains",
							Property: "property",
							Value:    "value",
						},
						Not: githubcomwithoursplatformsdkgo.GlobalDispatchCenterUpdateParamsCategoryLogicOrNot{
							Condition: githubcomwithoursplatformsdkgo.GlobalDispatchCenterUpdateParamsCategoryLogicOrNotCondition{
								Operator: "Contains",
								Property: "property",
								Value:    "value",
							},
						},
						Or: []githubcomwithoursplatformsdkgo.GlobalDispatchCenterUpdateParamsCategoryLogicOrOr{{
							Condition: githubcomwithoursplatformsdkgo.GlobalDispatchCenterUpdateParamsCategoryLogicOrOrCondition{
								Operator: "Contains",
								Property: "property",
								Value:    "value",
							},
						}},
					}},
				},
				Name:     githubcomwithoursplatformsdkgo.String("name"),
				Priority: githubcomwithoursplatformsdkgo.Float(0),
			}},
			IsEnabled: githubcomwithoursplatformsdkgo.Bool(true),
			Name:      githubcomwithoursplatformsdkgo.String("name"),
			Notes:     githubcomwithoursplatformsdkgo.String("notes"),
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

func TestGlobalDispatchCenterList(t *testing.T) {
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
	_, err := client.GlobalDispatchCenters.List(context.TODO())
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGlobalDispatchCenterDelete(t *testing.T) {
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
	_, err := client.GlobalDispatchCenters.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
