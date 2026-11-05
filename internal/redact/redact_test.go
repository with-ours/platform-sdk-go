// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package redact

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"testing"
)

const secret = "SENTINEL_SECRET_DO_NOT_LOG"

func TestBodyRedactsAtAnyDepthAndKeepsTheRest(t *testing.T) {
	body := `{"event":"still_visible","token":"` + secret + `","nested":[{"token": ["` + secret + `", {"a":"` + secret + `"}]}],"text":"say \"token\""}`
	got := String(body)
	if strings.Contains(got, secret) {
		t.Fatalf("secret in %s", got)
	}
	want := `{"event":"still_visible","token":"***","nested":[{"token": "***"}],"text":"say \"token\""}`
	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestBodyHandlesLargeBodies(t *testing.T) {
	padding := strings.Repeat("x", 1<<20)
	escapes := strings.Repeat(`\"`, 1<<16)
	got := String(`{"a":"` + padding + `","b":"` + escapes + `","token":"` + secret + `","c":"` + padding + `"}`)
	if strings.Contains(got, secret) || !strings.Contains(got, `"token":"***"`) || !strings.Contains(got, padding) {
		t.Fatalf("large body was not redacted correctly (%d bytes)", len(got))
	}
}

func TestBodyKeepsNonJSONAndOtherFields(t *testing.T) {
	for _, body := range []string{"", "not json", `{"event":"x"}`, `["a","b"]`} {
		if got := String(body); got != body {
			t.Fatalf("got %s, want %s", got, body)
		}
	}
}

func TestDumpsRedactAndKeepTheOriginalBody(t *testing.T) {
	body := `{"event":"still_visible","token":"` + secret + `"}`
	req, err := http.NewRequest("POST", "https://example.test/x", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+secret)

	for name, dump := range map[string]func(*http.Request, bool) ([]byte, error){
		"DumpRequest":    DumpRequest,
		"DumpRequestOut": DumpRequestOut,
	} {
		out, err := dump(req, true)
		if err != nil {
			t.Fatalf("%s: %v", name, err)
		}
		if bytes.Contains(out, []byte(secret)) || !bytes.Contains(out, []byte("still_visible")) {
			t.Fatalf("%s: %s", name, out)
		}
	}

	sent, err := io.ReadAll(req.Body)
	if err != nil || string(sent) != body {
		t.Fatalf("request body changed to %q: %v", sent, err)
	}
	if req.Header.Get("Authorization") != "Bearer "+secret {
		t.Fatalf("request header changed to %q", req.Header.Get("Authorization"))
	}
}

func TestDumpResponseRedactsAndKeepsTheOriginalBody(t *testing.T) {
	body := `{"event":"still_visible","token":"` + secret + `"}`
	resp := &http.Response{
		StatusCode:    400,
		ProtoMajor:    1,
		ProtoMinor:    1,
		Header:        http.Header{"Set-Cookie": {"session=" + secret}},
		Body:          io.NopCloser(strings.NewReader(body)),
		ContentLength: int64(len(body)),
	}
	out, err := DumpResponse(resp, true)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(out, []byte(secret)) || !bytes.Contains(out, []byte("still_visible")) {
		t.Fatalf("%s", out)
	}
	kept, err := io.ReadAll(resp.Body)
	if err != nil || string(kept) != body {
		t.Fatalf("response body changed to %q: %v", kept, err)
	}
}
