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

func TestExperimentVariantListWithOptionalParams(t *testing.T) {
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
	_, err := client.ExperimentVariants.List(context.TODO(), oursprivacy.ExperimentVariantListParams{
		ExperimentID: "08524dc8-5289-48e8-bf40-b3a7cfa6ca0a",
		Cursor:       oursprivacy.String("cursor"),
		Limit:        oursprivacy.Int(200),
	})
	if err != nil {
		var apierr *oursprivacy.Error
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
	client := oursprivacy.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ExperimentVariants.New(context.TODO(), oursprivacy.ExperimentVariantNewParams{
		ExperimentID: "x",
		Name:         "Variant B",
		Weight:       50,
		DomModifications: []oursprivacy.ExperimentVariantNewParamsDomModification{{
			Action:    "customCss",
			Selector:  "#hero-headline",
			Attribute: map[string]any{},
			Styles: []oursprivacy.ExperimentVariantNewParamsDomModificationStyle{{
				Property: "background-color",
				Value:    "#10B981",
			}},
			Value: oursprivacy.String("Start your free trial"),
		}},
		IsControl:   oursprivacy.Bool(false),
		RedirectURL: oursprivacy.String("https://www.example.com/pricing-v2"),
		VariantType: oursprivacy.ExperimentVariantNewParamsVariantTypeDomModifications,
	})
	if err != nil {
		var apierr *oursprivacy.Error
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
	client := oursprivacy.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ExperimentVariants.Get(context.TODO(), "id")
	if err != nil {
		var apierr *oursprivacy.Error
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
	client := oursprivacy.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ExperimentVariants.Update(
		context.TODO(),
		"id",
		oursprivacy.ExperimentVariantUpdateParams{
			DomModifications: []oursprivacy.ExperimentVariantUpdateParamsDomModification{{
				Action:    "customCss",
				Selector:  "#hero-headline",
				Attribute: map[string]any{},
				Styles: []oursprivacy.ExperimentVariantUpdateParamsDomModificationStyle{{
					Property: "background-color",
					Value:    "#10B981",
				}},
				Value: oursprivacy.String("Start your free trial"),
			}},
			IsControl:   oursprivacy.Bool(true),
			Name:        oursprivacy.String("name"),
			RedirectURL: oursprivacy.String("redirectUrl"),
			VariantType: oursprivacy.ExperimentVariantUpdateParamsVariantTypeDomModifications,
			Weight:      oursprivacy.Int(50),
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

func TestExperimentVariantDelete(t *testing.T) {
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
	_, err := client.ExperimentVariants.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
