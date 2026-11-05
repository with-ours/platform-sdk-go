// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oursprivacy_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/with-ours/platform-sdk-go/v2"
	"github.com/with-ours/platform-sdk-go/v2/option"
)

const redactionSecret = "SENTINEL_SECRET_DO_NOT_LOG"

type redactionTransport struct {
	fn func(req *http.Request) (*http.Response, error)
}

func (t *redactionTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.fn(req)
}

func TestDebugLogAndDumpsKeepSecretsOut(t *testing.T) {
	var sent []byte
	var logged bytes.Buffer
	client := oursprivacy.NewClient(
		option.WithAPIKey("My API Key"),
		option.WithMaxRetries(0),
		option.WithDebugLog(log.New(&logged, "", 0)),
		option.WithHTTPClient(&http.Client{
			Transport: &redactionTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					sent, _ = io.ReadAll(req.Body)
					return &http.Response{
						StatusCode: http.StatusBadRequest,
						Header: http.Header{
							"Content-Type": []string{"application/json"},
							"Set-Cookie":   []string{"session=" + redactionSecret},
						},
						Body:    io.NopCloser(strings.NewReader(`{"token":"` + redactionSecret + `","code":"still_visible"}`)),
						Request: req,
					}, nil
				},
			},
		}),
	)
	body := map[string]any{
		"event":  "still_visible",
		"token":  redactionSecret,
		"nested": []map[string]any{{"token": redactionSecret}},
	}
	err := client.Post(
		context.Background(),
		"foo",
		body,
		nil,
		option.WithHeader("Authorization", "Bearer "+redactionSecret),
	)

	var apierr *oursprivacy.Error
	if !errors.As(err, &apierr) {
		t.Fatalf("expected an API error, got %v", err)
	}
	outputs := map[string]string{
		"debug log":    logged.String(),
		"DumpRequest":  string(apierr.DumpRequest(true)),
		"DumpResponse": string(apierr.DumpResponse(true)),
		"Error":        apierr.Error(),
	}
	for name, output := range outputs {
		if strings.Contains(output, redactionSecret) {
			t.Errorf("%s contains the secret:\n%s", name, output)
		}
		if !strings.Contains(output, "still_visible") {
			t.Errorf("%s lost the rest of the body:\n%s", name, output)
		}
	}
	if !strings.Contains(string(sent), redactionSecret) {
		t.Errorf("the secret must still be sent, got %s", sent)
	}
}
