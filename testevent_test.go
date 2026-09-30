// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/with-ours/platform-sdk-go/v2"
	"github.com/with-ours/platform-sdk-go/v2/internal/testutil"
	"github.com/with-ours/platform-sdk-go/v2/option"
)

func TestTestEventListWithOptionalParams(t *testing.T) {
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
	_, err := client.TestEvents.List(context.TODO(), oursprivacy.TestEventListParams{
		Cursor: oursprivacy.String("cursor"),
		Limit:  oursprivacy.Int(25),
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTestEventNewWithOptionalParams(t *testing.T) {
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
	_, err := client.TestEvents.New(context.TODO(), oursprivacy.TestEventNewParams{
		EventName: "Purchase",
		DefaultProperties: oursprivacy.TestEventNewParamsDefaultProperties{
			EfTransactionID:   oursprivacy.String("_ef_transaction_id"),
			ActiveDuration:    oursprivacy.Float(0),
			AdID:              oursprivacy.String("ad_id"),
			AdmitadUid:        oursprivacy.String("admitad_uid"),
			AdsetID:           oursprivacy.String("adset_id"),
			Alart:             oursprivacy.String("alart"),
			Aleid:             oursprivacy.String("aleid"),
			Axwrt:             oursprivacy.String("axwrt"),
			BasisCid:          oursprivacy.String("basis_cid"),
			BeeswaxAuctionID:  oursprivacy.String("beeswax_auction_id"),
			BrowserLanguage:   oursprivacy.String("browser_language"),
			BrowserName:       oursprivacy.String("browser_name"),
			BrowserVersion:    oursprivacy.String("browser_version"),
			CampaignID:        oursprivacy.String("campaign_id"),
			Clickid:           oursprivacy.String("clickid"),
			Clid:              oursprivacy.String("clid"),
			CPUArchitecture:   oursprivacy.String("cpu_architecture"),
			CurrentURL:        oursprivacy.String("https://example.com/checkout"),
			Dclid:             oursprivacy.String("dclid"),
			DeviceModel:       oursprivacy.String("device_model"),
			DeviceType:        oursprivacy.String("device_type"),
			DeviceVendor:      oursprivacy.String("device_vendor"),
			Duration:          oursprivacy.Float(0),
			Encoding:          oursprivacy.String("encoding"),
			EngineName:        oursprivacy.String("engine_name"),
			EngineVersion:     oursprivacy.String("engine_version"),
			Epik:              oursprivacy.String("epik"),
			Fbc:               oursprivacy.String("fbc"),
			Fbclid:            oursprivacy.String("fbclid"),
			Fbp:               oursprivacy.String("fbp"),
			Fv:                oursprivacy.Bool(true),
			GadSource:         oursprivacy.String("gad_source"),
			Gbraid:            oursprivacy.String("gbraid"),
			Gclid:             oursprivacy.String("gclid"),
			Host:              oursprivacy.String("host"),
			Iframe:            oursprivacy.Bool(true),
			ImRef:             oursprivacy.String("im_ref"),
			IP:                oursprivacy.String("ip"),
			Irclickid:         oursprivacy.String("irclickid"),
			IsBot:             oursprivacy.Bool(true),
			LiFatID:           oursprivacy.String("li_fat_id"),
			Msclkid:           oursprivacy.String("msclkid"),
			Ndclid:            oursprivacy.String("ndclid"),
			NewS:              oursprivacy.Bool(true),
			ObClickID:         oursprivacy.String("ob_click_id"),
			Oppref:            oursprivacy.String("oppref"),
			OsName:            oursprivacy.String("os_name"),
			OsVersion:         oursprivacy.String("os_version"),
			PageHash:          oursprivacy.Float(0),
			Pathname:          oursprivacy.String("pathname"),
			Qclid:             oursprivacy.String("qclid"),
			RdtCid:            oursprivacy.String("rdt_cid"),
			ReceivedAt:        oursprivacy.String("received_at"),
			Referrer:          oursprivacy.String("https://example.com"),
			ReferringDomain:   oursprivacy.String("referring_domain"),
			Sacid:             oursprivacy.String("sacid"),
			Sccid:             oursprivacy.String("sccid"),
			ScreenHeight:      oursprivacy.Float(0),
			ScreenWidth:       oursprivacy.Float(0),
			SessionCount:      oursprivacy.Float(0),
			Sid:               oursprivacy.String("sid"),
			Sr:                oursprivacy.String("sr"),
			Title:             oursprivacy.String("title"),
			Ttclid:            oursprivacy.String("ttclid"),
			Twclid:            oursprivacy.String("twclid"),
			Uafvl:             oursprivacy.String("uafvl"),
			UserAgent:         oursprivacy.String("user_agent"),
			UtmCampaign:       oursprivacy.String("utm_campaign"),
			UtmContent:        oursprivacy.String("utm_content"),
			UtmMedium:         oursprivacy.String("utm_medium"),
			UtmName:           oursprivacy.String("utm_name"),
			UtmSource:         oursprivacy.String("utm_source"),
			UtmTerm:           oursprivacy.String("utm_term"),
			Version:           oursprivacy.String("version"),
			ViantClickID:      oursprivacy.String("viant_click_id"),
			ViantImpressionID: oursprivacy.String("viant_impression_id"),
			Wbraid:            oursprivacy.String("wbraid"),
			Webview:           oursprivacy.Bool(true),
		},
		DistinctID: oursprivacy.String("distinctId"),
		EventProperties: map[string]any{
			"revenue":  42.5,
			"currency": "USD",
		},
		SourceID: oursprivacy.String("sourceId"),
		UserProperties: map[string]any{
			"email": "test@example.com",
		},
		VisitorID: oursprivacy.String("visitorId"),
	})
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTestEventGet(t *testing.T) {
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
	_, err := client.TestEvents.Get(context.TODO(), "a1b2c3d4:ck9x8y7z6w5v4u3t")
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTestEventDispatches(t *testing.T) {
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
	_, err := client.TestEvents.Dispatches(context.TODO(), "a1b2c3d4:ck9x8y7z6w5v4u3t")
	if err != nil {
		var apierr *oursprivacy.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
