//go:build testing

package utility

import "net/http"

type FakeResponseWriter struct {
	StatusCode int
}

func (response *FakeResponseWriter) Header() http.Header {
	return http.Header{}
}

func (response *FakeResponseWriter) Write([]byte) (int, error) {
	return 0, nil
}

func (response *FakeResponseWriter) WriteHeader(statusCode int) {
	response.StatusCode = statusCode
}
