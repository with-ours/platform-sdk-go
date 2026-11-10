// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package redact

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httputil"
	"regexp"
	"strings"
)

// Package redact keeps secrets out of debug output. It never changes what is
// sent to the API.
//
// This package is for internal use only.

const placeholder = "***"

var sensitiveHeaders = []string{"authorization", "api-key", "x-api-key", "cookie", "set-cookie"}

// sensitiveFields holds field names in lower case, without underscores or hyphens.
var sensitiveFields = map[string]struct{}{
	"accesstoken":       {},
	"apikey":            {},
	"clientsecret":      {},
	"credentials":       {},
	"ingesturl":         {},
	"installscript":     {},
	"password":          {},
	"projectapikey":     {},
	"refreshtoken":      {},
	"samplecurl":        {},
	"testingesturl":     {},
	"testinstallscript": {},
	"testtoken":         {},
	"token":             {},
	"websdktoken":       {},
}

var jsonKey = regexp.MustCompile(`"((?:[^"\\]|\\.)*)"\s*:\s*`)

var fieldSeparators = strings.NewReplacer("_", "", "-", "")

func isSensitiveField(name string) bool {
	_, ok := sensitiveFields[fieldSeparators.Replace(strings.ToLower(name))]
	return ok
}

// Headers returns headers with sensitive values replaced, leaving headers unchanged.
func Headers(headers http.Header) http.Header {
	redacted := headers.Clone()
	for _, name := range sensitiveHeaders {
		values := redacted.Values(name)
		if len(values) == 0 {
			continue
		}
		redacted.Del(name)
		for range values {
			redacted.Add(name, placeholder)
		}
	}
	return redacted
}

// Body returns a JSON body with the values of sensitive fields replaced at any
// depth, leaving body unchanged. Anything else in body is kept as it is.
func Body(body []byte) []byte {
	var out []byte
	next := 0
	for _, match := range jsonKey.FindAllSubmatchIndex(body, -1) {
		if match[0] < next || !isSensitiveField(string(body[match[2]:match[3]])) {
			continue
		}
		out = append(out, body[next:match[1]]...)
		out = append(out, '"')
		out = append(out, placeholder...)
		out = append(out, '"')
		next = valueEnd(body, match[1])
	}
	if next == 0 {
		return body
	}
	return append(out, body[next:]...)
}

// String is [Body] for text.
func String(body string) string {
	return string(Body([]byte(body)))
}

// valueEnd returns the index just past the JSON value starting at start.
func valueEnd(body []byte, start int) int {
	depth := 0
	inString := false
	for i := start; i < len(body); i++ {
		c := body[i]
		switch {
		case inString:
			if c == '\\' {
				i++
			} else if c == '"' {
				inString = false
				if depth == 0 {
					return i + 1
				}
			}
		case c == '"':
			inString = true
		case c == '{' || c == '[':
			depth++
		case c == '}' || c == ']':
			if depth == 0 {
				return i
			}
			depth--
			if depth == 0 {
				return i + 1
			}
		case depth == 0 && (c == ',' || c == ' ' || c == '\t' || c == '\r' || c == '\n'):
			return i
		}
	}
	return len(body)
}

func readBody(body io.ReadCloser) ([]byte, io.ReadCloser, error) {
	data, err := io.ReadAll(body)
	body.Close()
	return data, io.NopCloser(bytes.NewReader(data)), err
}

// redactedRequest returns a copy of req to dump. req keeps a readable body with
// its original bytes.
func redactedRequest(req *http.Request, body bool) (*http.Request, error) {
	redacted := req.Clone(req.Context())
	redacted.Header = Headers(req.Header)
	if !body || req.Body == nil || req.Body == http.NoBody {
		return redacted, nil
	}
	data, restored, err := readBody(req.Body)
	req.Body = restored
	if err != nil {
		return nil, err
	}
	data = Body(data)
	redacted.Body = io.NopCloser(bytes.NewReader(data))
	redacted.GetBody = nil
	redacted.ContentLength = int64(len(data))
	return redacted, nil
}

// DumpRequest is [httputil.DumpRequest] with secrets replaced.
func DumpRequest(req *http.Request, body bool) ([]byte, error) {
	redacted, err := redactedRequest(req, body)
	if err != nil {
		return nil, err
	}
	return httputil.DumpRequest(redacted, body)
}

// DumpRequestOut is [httputil.DumpRequestOut] with secrets replaced.
func DumpRequestOut(req *http.Request, body bool) ([]byte, error) {
	redacted, err := redactedRequest(req, body)
	if err != nil {
		return nil, err
	}
	return httputil.DumpRequestOut(redacted, body)
}

// DumpResponse is [httputil.DumpResponse] with secrets replaced. resp keeps a
// readable body with its original bytes.
func DumpResponse(resp *http.Response, body bool) ([]byte, error) {
	redacted := *resp
	redacted.Header = Headers(resp.Header)
	if body && resp.Body != nil && resp.Body != http.NoBody {
		data, restored, err := readBody(resp.Body)
		resp.Body = restored
		if err != nil {
			return nil, err
		}
		data = Body(data)
		redacted.Body = io.NopCloser(bytes.NewReader(data))
		if redacted.ContentLength > 0 {
			redacted.ContentLength = int64(len(data))
		}
	}
	return httputil.DumpResponse(&redacted, body)
}

// WriteParams writes params as JSON with secrets replaced, for the Format methods
// of params that hold a secret. Nothing is written from a value that can't be marshaled.
func WriteParams(w io.Writer, params any) {
	data, err := json.Marshal(params)
	if err != nil {
		io.WriteString(w, "(params omitted: "+placeholder+")")
		return
	}
	w.Write(Body(data))
}
