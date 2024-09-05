package utility

import (
	"api/internal/utility/test"
	"net/http"
	"net/http/httptest"
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
	err := test.DummyError{}
	responseWriter := httptest.NewRecorder()

	LogAndErrorResponse(err, responseWriter)

	if responseWriter.Code != http.StatusInternalServerError {
		t.Fatalf("StatusCode was %d instead of %d", responseWriter.Code, http.StatusInternalServerError)
	}
}
