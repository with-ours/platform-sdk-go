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

func TestLocationList(t *testing.T) {
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
	_, err := client.Locations.List(context.TODO())
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLocationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Locations.New(context.TODO(), oursprivacy.LocationNewParams{
		AdditionalAddresses: []oursprivacy.LocationNewParamsAdditionalAddress{{
			Latitude:        0,
			Longitude:       0,
			City:            oursprivacy.String("city"),
			Country:         oursprivacy.String("country"),
			Line1:           oursprivacy.String("line1"),
			Line2:           oursprivacy.String("line2"),
			Name:            oursprivacy.String("name"),
			PhoneNumber:     oursprivacy.String("phoneNumber"),
			State:           oursprivacy.String("state"),
			WebsiteLinkText: oursprivacy.String("websiteLinkText"),
			WebsiteURL:      oursprivacy.String("websiteUrl"),
			Zip:             oursprivacy.String("zip"),
		}},
		Center:          map[string]any{},
		City:            oursprivacy.String("city"),
		Country:         oursprivacy.String("country"),
		CustomDomain:    oursprivacy.String("customDomain"),
		Latitude:        oursprivacy.Float(0),
		Line1:           oursprivacy.String("line1"),
		Line2:           oursprivacy.String("line2"),
		Longitude:       oursprivacy.Float(0),
		MapName:         oursprivacy.String("mapName"),
		Name:            oursprivacy.String("name"),
		PhoneNumber:     oursprivacy.String("phoneNumber"),
		State:           oursprivacy.String("state"),
		WebsiteLinkText: oursprivacy.String("websiteLinkText"),
		WebsiteURL:      oursprivacy.String("websiteUrl"),
		Zip:             oursprivacy.String("zip"),
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLocationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Locations.Update(
		context.TODO(),
		"id",
		oursprivacy.LocationUpdateParams{
			AdditionalAddresses: []oursprivacy.LocationUpdateParamsAdditionalAddress{{
				Latitude:        0,
				Longitude:       0,
				City:            oursprivacy.String("city"),
				Country:         oursprivacy.String("country"),
				Line1:           oursprivacy.String("line1"),
				Line2:           oursprivacy.String("line2"),
				Name:            oursprivacy.String("name"),
				PhoneNumber:     oursprivacy.String("phoneNumber"),
				State:           oursprivacy.String("state"),
				WebsiteLinkText: oursprivacy.String("websiteLinkText"),
				WebsiteURL:      oursprivacy.String("websiteUrl"),
				Zip:             oursprivacy.String("zip"),
			}},
			Center:          map[string]any{},
			City:            oursprivacy.String("city"),
			Country:         oursprivacy.String("country"),
			CustomDomain:    oursprivacy.String("customDomain"),
			Latitude:        oursprivacy.Float(0),
			Line1:           oursprivacy.String("line1"),
			Line2:           oursprivacy.String("line2"),
			Longitude:       oursprivacy.Float(0),
			MapName:         oursprivacy.String("mapName"),
			Name:            oursprivacy.String("name"),
			PhoneNumber:     oursprivacy.String("phoneNumber"),
			State:           oursprivacy.String("state"),
			WebsiteLinkText: oursprivacy.String("websiteLinkText"),
			WebsiteURL:      oursprivacy.String("websiteUrl"),
			Zip:             oursprivacy.String("zip"),
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

func TestLocationEmbedCodeWithOptionalParams(t *testing.T) {
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
	_, err := client.Locations.EmbedCode(
		context.TODO(),
		"id",
		oursprivacy.LocationEmbedCodeParams{
			Color:             oursprivacy.String("#007EA8"),
			ColorScheme:       oursprivacy.LocationEmbedCodeParamsColorSchemeLight,
			IncludeAddressBox: oursprivacy.Bool(true),
			IncludeControls:   oursprivacy.LocationEmbedCodeParamsIncludeControlsYes,
			IncludeSeoSchema:  oursprivacy.Bool(true),
			MapStyle:          oursprivacy.LocationEmbedCodeParamsMapStyleStandard,
			Theme:             oursprivacy.LocationEmbedCodeParamsThemeDefault,
			Zoom:              oursprivacy.Int(1),
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
