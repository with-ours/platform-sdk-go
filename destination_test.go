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

func TestDestinationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Destinations.New(context.TODO(), githubcomwithoursplatformsdkgo.DestinationNewParams{
		Type: githubcomwithoursplatformsdkgo.DestinationNewParamsTypeAwsEventBridge,
		Name: githubcomwithoursplatformsdkgo.String("name"),
	})
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
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
	client := githubcomwithoursplatformsdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Destinations.Get(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
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
	client := githubcomwithoursplatformsdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Destinations.Update(
		context.TODO(),
		"id",
		githubcomwithoursplatformsdkgo.DestinationUpdateParams{
			Status:                   githubcomwithoursplatformsdkgo.DestinationUpdateParamsStatusDisabled,
			FacebookConversionAPIKey: githubcomwithoursplatformsdkgo.String("facebookConversionAPIKey"),
			FacebookPixelID:          githubcomwithoursplatformsdkgo.String("facebookPixelId"),
			G4AnalyticsAPIKey:        githubcomwithoursplatformsdkgo.String("g4AnalyticsApiKey"),
			G4AnalyticsMeasurementID: githubcomwithoursplatformsdkgo.String("g4AnalyticsMeasurementId"),
			G4AnalyticsTrackOnPage:   githubcomwithoursplatformsdkgo.Bool(true),
			HashingSalt:              githubcomwithoursplatformsdkgo.String("hashingSalt"),
			HTTPDestinationURL:       githubcomwithoursplatformsdkgo.String("httpDestinationUrl"),
			LimitedToSourceIDs:       []string{"string"},
			ManagerGoogleCustomerID:  githubcomwithoursplatformsdkgo.String("managerGoogleCustomerId"),
			Name:                     githubcomwithoursplatformsdkgo.String("name"),
			ProjectAPIKey:            githubcomwithoursplatformsdkgo.String("projectAPIKey"),
			ProjectToken:             githubcomwithoursplatformsdkgo.String("projectToken"),
			SelectedAccountID:        githubcomwithoursplatformsdkgo.String("selectedAccountId"),
			Settings: map[string]githubcomwithoursplatformsdkgo.DestinationUpdateParamsSettingUnion{
				"foo": {
					OfString: githubcomwithoursplatformsdkgo.String("string"),
				},
			},
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

func TestDestinationList(t *testing.T) {
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
	_, err := client.Destinations.List(context.TODO())
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
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
	client := githubcomwithoursplatformsdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Destinations.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
