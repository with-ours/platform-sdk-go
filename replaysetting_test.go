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

func TestReplaySettingNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ReplaySettings.New(context.TODO(), githubcomwithoursplatformsdkgo.ReplaySettingNewParams{
		CustomDomain:     githubcomwithoursplatformsdkgo.String("customDomain"),
		Name:             githubcomwithoursplatformsdkgo.String("name"),
		Status:           githubcomwithoursplatformsdkgo.ReplaySettingNewParamsStatusDisabled,
		WhitelistDomains: []string{"string"},
	})
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestReplaySettingGet(t *testing.T) {
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
	_, err := client.ReplaySettings.Get(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestReplaySettingUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ReplaySettings.Update(
		context.TODO(),
		"id",
		githubcomwithoursplatformsdkgo.ReplaySettingUpdateParams{
			CustomDomain:     githubcomwithoursplatformsdkgo.String("customDomain"),
			Name:             githubcomwithoursplatformsdkgo.String("name"),
			Status:           githubcomwithoursplatformsdkgo.ReplaySettingUpdateParamsStatusDisabled,
			WhitelistDomains: []string{"string"},
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

func TestReplaySettingList(t *testing.T) {
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
	_, err := client.ReplaySettings.List(context.TODO())
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestReplaySettingDelete(t *testing.T) {
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
	_, err := client.ReplaySettings.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
