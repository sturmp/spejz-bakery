package auth

import (
	"api/internal/utility"
	"net/http"
	"net/url"
	"testing"
)

func TestIsAdminEndpoint_WithAdminEndpoint(t *testing.T) {
	handler := FakeHandler{}
	path := "admin"
	method := "GET"
	endpoints := []Endpoint{}
	auth := NewAuth(handler, "abcd", "adminabcd", endpoints)

	result := isAdminEndpoint(auth, path, method)

	if !result {
		t.Fatalf("%s - %s should be identified as admin endpoint", path, method)
	}
}

func TestIsAdminEndpoint_WithNonAdminEndpoint(t *testing.T) {
	handler := FakeHandler{}
	path := "pastry"
	method := "GET"
	endpoints := []Endpoint{
		{
			Path:   path,
			Method: method,
		},
	}
	auth := NewAuth(handler, "abcd", "adminabcd", endpoints)

	result := isAdminEndpoint(auth, path, method)

	if result {
		t.Fatalf("%s - %s should be identified as non admin endpoint", path, method)
	}
}

func TestServeHTTP_MissingAuthToken(t *testing.T) {
	request := http.Request{}
	request.Header = http.Header{}
	handler := FakeHandler{}
	auth := NewAuth(handler, "abc", "adminabc", []Endpoint{})
	responseWriter := utility.FakeResponseWriter{}

	auth.ServeHTTP(&responseWriter, &request)

	if responseWriter.StatusCode != http.StatusUnauthorized {
		t.Fatalf("Response status code should be %d not %d", http.StatusUnauthorized, responseWriter.StatusCode)
	}
}

func TestServeHTTP_WrongAuthToken(t *testing.T) {
	request := http.Request{}
	request.Header = http.Header{}
	handler := FakeHandler{}
	authToken := "abc"
	path := "pastry"
	method := "GET"

	request.Header.Add("AuthToken", "nsdoivnws")
	request.URL = &url.URL{
		Path: path,
	}
	request.Method = method
	auth := NewAuth(handler, authToken, "adminabc", []Endpoint{
		{
			Path:   path,
			Method: method,
		},
	})
	responseWriter := utility.FakeResponseWriter{}

	auth.ServeHTTP(&responseWriter, &request)

	if responseWriter.StatusCode != http.StatusUnauthorized {
		t.Fatalf("Response status code should be %d not %d", http.StatusUnauthorized, responseWriter.StatusCode)
	}
}

func TestServeHTTP_CorrectAuthToken(t *testing.T) {
	request := http.Request{}
	request.Header = http.Header{}
	handler := FakeHandler{}
	authToken := "abc"
	path := "pastry"
	method := "GET"

	request.Header.Add("AuthToken", authToken)
	request.URL = &url.URL{
		Path: path,
	}
	request.Method = method
	auth := NewAuth(handler, authToken, "adminabc", []Endpoint{
		{
			Path:   path,
			Method: method,
		},
	})
	responseWriter := utility.FakeResponseWriter{}

	auth.ServeHTTP(&responseWriter, &request)

	if responseWriter.StatusCode == http.StatusUnauthorized {
		t.Fatalf("Response status code should not be %d", http.StatusUnauthorized)
	}
}

func TestServeHTTP_WrongAdminAuthToken(t *testing.T) {
	request := http.Request{}
	request.Header = http.Header{}
	handler := FakeHandler{}
	adminAuthToken := "adminabc"
	path := "pastry"
	method := "GET"

	request.Header.Add("AuthToken", "nsdoivnws")
	request.URL = &url.URL{
		Path: path,
	}
	request.Method = method
	auth := NewAuth(handler, "abc", adminAuthToken, []Endpoint{})
	responseWriter := utility.FakeResponseWriter{}

	auth.ServeHTTP(&responseWriter, &request)

	if responseWriter.StatusCode != http.StatusUnauthorized {
		t.Fatalf("Response status code should be %d not %d", http.StatusUnauthorized, responseWriter.StatusCode)
	}
}

func TestServeHTTP_CorrectAdminAuthToken(t *testing.T) {
	request := http.Request{}
	request.Header = http.Header{}
	handler := FakeHandler{}
	adminAuthToken := "adminabc"
	path := "pastry"
	method := "GET"

	request.Header.Add("AuthToken", adminAuthToken)
	request.URL = &url.URL{
		Path: path,
	}
	request.Method = method
	auth := NewAuth(handler, "abc", adminAuthToken, []Endpoint{})
	responseWriter := utility.FakeResponseWriter{}

	auth.ServeHTTP(&responseWriter, &request)

	if responseWriter.StatusCode == http.StatusUnauthorized {
		t.Fatalf("Response status code should not be %d", http.StatusUnauthorized)
	}
}

type FakeHandler struct {
}

func (FakeHandler) ServeHTTP(http.ResponseWriter, *http.Request) {}
