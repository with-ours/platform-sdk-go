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

func TestDestinationList(t *testing.T) {
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
	_, err := client.Destinations.List(context.TODO())
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDestinationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Destinations.New(context.TODO(), oursprivacy.DestinationNewParams{
		Type: oursprivacy.DestinationNewParamsTypeAwsEventBridge,
		Name: oursprivacy.String("name"),
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDestinationGet(t *testing.T) {
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
	_, err := client.Destinations.Get(context.TODO(), "id")
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDestinationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Destinations.Update(
		context.TODO(),
		"id",
		oursprivacy.DestinationUpdateParams{
			Status:                   oursprivacy.DestinationUpdateParamsStatusDisabled,
			FacebookConversionAPIKey: oursprivacy.String("facebookConversionAPIKey"),
			FacebookPixelID:          oursprivacy.String("facebookPixelId"),
			G4AnalyticsAPIKey:        oursprivacy.String("g4AnalyticsApiKey"),
			G4AnalyticsMeasurementID: oursprivacy.String("g4AnalyticsMeasurementId"),
			G4AnalyticsTrackOnPage:   oursprivacy.Bool(true),
			HashingSalt:              oursprivacy.String("hashingSalt"),
			HTTPDestinationURL:       oursprivacy.String("httpDestinationUrl"),
			LimitedToSourceIDs:       []string{"string"},
			ManagerGoogleCustomerID:  oursprivacy.String("managerGoogleCustomerId"),
			Name:                     oursprivacy.String("name"),
			ProjectAPIKey:            oursprivacy.String("projectAPIKey"),
			ProjectToken:             oursprivacy.String("projectToken"),
			SelectedAccountID:        oursprivacy.String("selectedAccountId"),
			Settings:                 map[string]any{},
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

func TestDestinationDelete(t *testing.T) {
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
	_, err := client.Destinations.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
