// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/with-ours/platform-sdk-go"
	"github.com/with-ours/platform-sdk-go/internal/testutil"
	"github.com/with-ours/platform-sdk-go/option"
)

func TestSessionReplayListWithOptionalParams(t *testing.T) {
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
	_, err := client.SessionReplays.List(context.TODO(), oursprivacy.SessionReplayListParams{
		From:        time.Now(),
		To:          time.Now(),
		Cursor:      oursprivacy.String("x"),
		EventName:   oursprivacy.String("x"),
		Limit:       oursprivacy.Int(1),
		Pathname:    oursprivacy.String("x"),
		SessionIDs:  oursprivacy.String("sessionIds"),
		UtmCampaign: oursprivacy.String("x"),
		UtmContent:  oursprivacy.String("x"),
		UtmMedium:   oursprivacy.String("x"),
		UtmName:     oursprivacy.String("x"),
		UtmSource:   oursprivacy.String("x"),
		UtmTerm:     oursprivacy.String("x"),
		VisitorID:   oursprivacy.String("x"),
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionReplayOverview(t *testing.T) {
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
	_, err := client.SessionReplays.Overview(context.TODO(), oursprivacy.SessionReplayOverviewParams{
		From: time.Now(),
		To:   time.Now(),
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
