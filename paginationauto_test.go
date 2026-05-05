// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomwithoursplatformsdkgo_test

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
	client := githubcomwithoursplatformsdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	iter := client.Experiments.ListAutoPaging(context.TODO(), githubcomwithoursplatformsdkgo.ExperimentListParams{})
	// The mock server isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		experiment := iter.Current()
		t.Logf("%+v\n", experiment.ID)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
