// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package option

import (
	"log"
	"net/http"

	"github.com/with-ours/platform-sdk-go/v2/internal/redact"
)

// WithDebugLog logs the HTTP request and response content, with secrets replaced.
// If the logger parameter is nil, it uses the default logger.
//
// WithDebugLog is for debugging and development purposes only.
// It should not be used in production code. The behavior and interface
// of WithDebugLog is not guaranteed to be stable.
func WithDebugLog(logger *log.Logger) RequestOption {
	return WithMiddleware(func(req *http.Request, nxt MiddlewareNext) (*http.Response, error) {
		if logger == nil {
			logger = log.Default()
		}

		if reqBytes, err := redact.DumpRequest(req, true); err == nil {
			logger.Printf("Request Content:\n%s\n", reqBytes)
		}

		resp, err := nxt(req)
		if err != nil {
			return resp, err
		}

		if respBytes, err := redact.DumpResponse(resp, true); err == nil {
			logger.Printf("Response Content:\n%s\n", respBytes)
		}

		return resp, err
	})
}
