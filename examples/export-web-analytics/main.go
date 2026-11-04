// Export the last 7 days of top pages for every web source, as CSV on stdout.
// This only reads data.
//
//	OURS_PRIVACY_API_KEY=<api key> go run ./examples/export-web-analytics
//
// Set OURS_PRIVACY_PLATFORM_BASE_URL to read from a host other than production.
//
// The pages report has no cursor of its own, so the pagination here walks the
// list of web sources and requests one report per source.
package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	oursprivacy "github.com/with-ours/platform-sdk-go/v2"
)

func main() {
	if os.Getenv("OURS_PRIVACY_API_KEY") == "" {
		log.Fatal("OURS_PRIVACY_API_KEY is required: set it to your Platform API key")
	}
	ctx := context.Background()
	client := oursprivacy.NewClient()

	to := time.Now().UTC()
	from := to.AddDate(0, 0, -7)

	out := csv.NewWriter(os.Stdout)
	out.Write([]string{"webSourceId", "hostname", "path", "pageviews", "visitors"})

	// ListAutoPaging fetches further pages as the loop reaches them.
	sources := client.Sources.ListAutoPaging(ctx, oursprivacy.SourceListParams{
		Type:  oursprivacy.SourceListParamsTypeWebSource,
		Limit: oursprivacy.Int(100),
	})
	for sources.Next() {
		source := sources.Current()
		report, err := client.WebAnalytics.Pages(ctx, oursprivacy.WebAnalyticsPagesParams{
			From:        from,
			To:          to,
			View:        oursprivacy.WebAnalyticsPagesParamsViewTop,
			WebSourceID: oursprivacy.String(source.ID),
		})
		if err != nil {
			log.Fatal(err)
		}
		for _, row := range report.Rows {
			out.Write([]string{
				source.ID, row.PageHostname, row.PagePath,
				fmt.Sprint(row.Pageviews), fmt.Sprint(row.Visitors),
			})
		}
		out.Flush() // keep finished sources if a later request fails
	}
	if err := sources.Err(); err != nil {
		log.Fatal(err)
	}
	if err := out.Error(); err != nil {
		log.Fatal(err)
	}
}
