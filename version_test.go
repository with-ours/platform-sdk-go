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

func TestVersionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Versions.List(context.TODO(), githubcomwithoursplatformsdkgo.VersionListParams{
		Cursor:        githubcomwithoursplatformsdkgo.String("cursor"),
		IsPublished:   githubcomwithoursplatformsdkgo.VersionListParamsIsPublishedTrue,
		Limit:         githubcomwithoursplatformsdkgo.Int(25),
		NameContains:  githubcomwithoursplatformsdkgo.String("nameContains"),
		NotesContains: githubcomwithoursplatformsdkgo.String("notesContains"),
	})
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVersionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Versions.New(context.TODO(), githubcomwithoursplatformsdkgo.VersionNewParams{
		IncludeAllowedEvents:            []string{"string"},
		IncludeConsentSettings:          []string{"string"},
		IncludeDataGovernanceEvents:     []string{"string"},
		IncludeDataGovernanceRules:      []string{"string"},
		IncludeDestinations:             []string{"string"},
		IncludeExperiments:              []string{"string"},
		IncludeExperimentSettings:       []string{"string"},
		IncludeExperimentVariants:       []string{"string"},
		IncludeExternalAllowedEventData: []string{"string"},
		IncludeGlobalDispatchCenters:    []string{"string"},
		IncludeMappings:                 []string{"string"},
		IncludeReplaySettings:           []string{"string"},
		IncludeSources:                  []string{"string"},
		IncludeTagManagerTags:           []string{"string"},
		IncludeTagManagerTriggers:       []string{"string"},
		IncludeTagManagerVariables:      []string{"string"},
		Name:                            githubcomwithoursplatformsdkgo.String("name"),
		Notes:                           githubcomwithoursplatformsdkgo.String("notes"),
	})
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVersionGet(t *testing.T) {
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
	_, err := client.Versions.Get(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVersionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Versions.Update(
		context.TODO(),
		"id",
		githubcomwithoursplatformsdkgo.VersionUpdateParams{
			Name:  githubcomwithoursplatformsdkgo.String("name"),
			Notes: githubcomwithoursplatformsdkgo.String("notes"),
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

func TestVersionPublish(t *testing.T) {
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
	_, err := client.Versions.Publish(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVersionSnapshot(t *testing.T) {
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
	_, err := client.Versions.Snapshot(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomwithoursplatformsdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVersionDiffWithOptionalParams(t *testing.T) {
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
	_, err := client.Versions.Diff(
		context.TODO(),
		githubcomwithoursplatformsdkgo.VersionDiffParamsIDDraft,
		githubcomwithoursplatformsdkgo.VersionDiffParams{
			Against: githubcomwithoursplatformsdkgo.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
