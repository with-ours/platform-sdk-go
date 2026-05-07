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

func TestExperimentVariantListWithOptionalParams(t *testing.T) {
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
	_, err := client.ExperimentVariants.List(context.TODO(), githubcomwithoursplatformsdkgo.ExperimentVariantListParams{
		ExperimentID: "08524dc8-5289-48e8-bf40-b3a7cfa6ca0a",
		Cursor:       githubcomwithoursplatformsdkgo.String("cursor"),
		Limit:        githubcomwithoursplatformsdkgo.Int(200),
	})
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExperimentVariantNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ExperimentVariants.New(context.TODO(), githubcomwithoursplatformsdkgo.ExperimentVariantNewParams{
		ExperimentID: "x",
		Name:         "Variant B",
		Weight:       50,
		DomModifications: []githubcomwithoursplatformsdkgo.ExperimentVariantNewParamsDomModification{{
			Action:    "customCss",
			Selector:  "h1.hero-title",
			Attribute: map[string]any{},
			Styles: []githubcomwithoursplatformsdkgo.ExperimentVariantNewParamsDomModificationStyle{{
				Property: "background-color",
				Value:    "#10B981",
			}},
			Value: githubcomwithoursplatformsdkgo.String("Start your free trial"),
		}},
		IsControl:   githubcomwithoursplatformsdkgo.Bool(false),
		RedirectURL: githubcomwithoursplatformsdkgo.String("https://www.example.com/pricing-v2"),
		VariantType: githubcomwithoursplatformsdkgo.ExperimentVariantNewParamsVariantTypeDomModifications,
	})
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExperimentVariantGet(t *testing.T) {
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
	_, err := client.ExperimentVariants.Get(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExperimentVariantUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ExperimentVariants.Update(
		context.TODO(),
		"id",
		githubcomwithoursplatformsdkgo.ExperimentVariantUpdateParams{
			DomModifications: []githubcomwithoursplatformsdkgo.ExperimentVariantUpdateParamsDomModification{{
				Action:    "customCss",
				Selector:  "h1.hero-title",
				Attribute: map[string]any{},
				Styles: []githubcomwithoursplatformsdkgo.ExperimentVariantUpdateParamsDomModificationStyle{{
					Property: "background-color",
					Value:    "#10B981",
				}},
				Value: githubcomwithoursplatformsdkgo.String("Start your free trial"),
			}},
			IsControl:   githubcomwithoursplatformsdkgo.Bool(true),
			Name:        githubcomwithoursplatformsdkgo.String("name"),
			RedirectURL: githubcomwithoursplatformsdkgo.String("redirectUrl"),
			VariantType: githubcomwithoursplatformsdkgo.ExperimentVariantUpdateParamsVariantTypeDomModifications,
			Weight:      githubcomwithoursplatformsdkgo.Int(50),
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

func TestExperimentVariantDelete(t *testing.T) {
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
	_, err := client.ExperimentVariants.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
