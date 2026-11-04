// Review pending config changes, and publish them only when asked to. Config edits
// write to a draft and go live on publish. Experiments ship separately.
//
// By default this reads the draft and prints what it would do:
//
//	OURS_PRIVACY_API_KEY=<api key> go run ./examples/review-and-publish-draft
//
// To publish every pending draft change as a new live version:
//
//	OURS_PRIVACY_API_KEY=<api key> OURS_PRIVACY_EXAMPLE_PUBLISH=1 go run ./examples/review-and-publish-draft
//
// Set OURS_PRIVACY_PLATFORM_BASE_URL to use a host other than production.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	oursprivacy "github.com/with-ours/platform-sdk-go/v2"
)

const versionName = "Example release"

func main() {
	if os.Getenv("OURS_PRIVACY_API_KEY") == "" {
		log.Fatal("OURS_PRIVACY_API_KEY is required: set it to your Platform API key")
	}
	ctx := context.Background()
	client := oursprivacy.NewClient()

	diff, err := client.Versions.Diff(ctx, oursprivacy.VersionDiffParamsIDDraft, oursprivacy.VersionDiffParams{})
	if err != nil {
		log.Fatal(err)
	}
	if diff.Count == 0 {
		fmt.Println("The draft has no pending changes, so there is nothing to publish.")
		return
	}
	fmt.Printf("%.0f pending draft changes:\n%s\n", diff.Count, diff.Differences.RawJSON())

	if os.Getenv("OURS_PRIVACY_EXAMPLE_PUBLISH") != "1" {
		fmt.Printf("Dry run: would publish these changes as version %q. Set OURS_PRIVACY_EXAMPLE_PUBLISH=1 to publish.\n", versionName)
		return
	}

	// With no include* fields, this publishes everything in the draft at this moment.
	version, err := client.Versions.New(ctx, oursprivacy.VersionNewParams{Name: oursprivacy.String(versionName)})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("published version:", version.ID)
}
