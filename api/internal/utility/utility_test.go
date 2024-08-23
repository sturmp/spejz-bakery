package utility

import (
	"net/http"
	"testing"
)

func TestGetLanguageOrDefault_Default(t *testing.T) {
	request := http.Request{}
	languageCode := GetLanguageOrDefault(&request)
	expected := "en"

	if languageCode != expected {
		t.Fatalf("languageCode was %s instead of %s", languageCode, expected)
	}
}

func TestGetLanguageOrDefault_FromHeader(t *testing.T) {
	request := http.Request{}
	request.Header = http.Header{}
	request.Header.Add("Accept-Language", "hu")

	languageCode := GetLanguageOrDefault(&request)
	expected := "hu"

	if languageCode != expected {
		t.Fatalf("languageCode was %s instead of %s", languageCode, expected)
	}
}

func TestLogAndErrorResponse_ResponseCode(t *testing.T) {
	err := DummyError{}
	responseWriter := FakeResponseWriter{}

	LogAndErrorResponse(err, &responseWriter)

	if responseWriter.StatusCode != http.StatusInternalServerError {
		t.Fatalf("StatusCode was %d instead of %d", responseWriter.StatusCode, http.StatusInternalServerError)
	}
}

type DummyError struct{}

func (err DummyError) Error() string {
	return "Error"
}

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
