// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy_test

import (
	"context"
	"os"
	"testing"

	"github.com/with-ours/platform-sdk-go"
	"github.com/with-ours/platform-sdk-go/internal/testutil"
	"github.com/with-ours/platform-sdk-go/option"
)

func TestAutoPagination(t *testing.T) {
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
	iter := client.Destinations.ListAutoPaging(context.TODO(), oursprivacy.DestinationListParams{})
	// The mock server isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		destination := iter.Current()
		t.Logf("%+v\n", destination.ID)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
