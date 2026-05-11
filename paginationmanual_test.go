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

func TestManualPagination(t *testing.T) {
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
	page, err := client.ExperimentVariants.List(context.TODO(), oursprivacy.ExperimentVariantListParams{
		ExperimentID: "08524dc8-5289-48e8-bf40-b3a7cfa6ca0a",
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, experimentVariant := range page.Entities {
		t.Logf("%+v\n", experimentVariant.ID)
	}
	// The mock server isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, experimentVariant := range page.Entities {
			t.Logf("%+v\n", experimentVariant.ID)
		}
	}
}
