package auth

import (
	"net/http"
	"net/http/httptest"
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
	responseWriter := httptest.NewRecorder()

	auth.ServeHTTP(responseWriter, &request)

	if responseWriter.Code != http.StatusUnauthorized {
		t.Fatalf("Response status code should be %d not %d", http.StatusUnauthorized, responseWriter.Code)
	}
}

func TestServeHTTP_WrongAuthToken(t *testing.T) {
	handler := FakeHandler{}
	authToken := "abc"
	path := "pastry"
	method := "GET"

	request, _ := http.NewRequest(method, path, nil)
	request.Header.Add("AuthToken", "nsdoivnws")
	auth := NewAuth(handler, authToken, "adminabc", []Endpoint{
		{
			Path:   path,
			Method: method,
		},
	})
	responseWriter := httptest.NewRecorder()

	auth.ServeHTTP(responseWriter, request)

	if responseWriter.Code != http.StatusUnauthorized {
		t.Fatalf("Response status code should be %d not %d", http.StatusUnauthorized, responseWriter.Code)
	}
}

func TestServeHTTP_CorrectAuthToken(t *testing.T) {
	handler := FakeHandler{}
	authToken := "abc"
	path := "pastry"
	method := "GET"

	request, _ := http.NewRequest(method, path, nil)
	request.Header.Add("AuthToken", authToken)
	auth := NewAuth(handler, authToken, "adminabc", []Endpoint{
		{
			Path:   path,
			Method: method,
		},
	})
	responseWriter := httptest.NewRecorder()

	auth.ServeHTTP(responseWriter, request)

	if responseWriter.Code == http.StatusUnauthorized {
		t.Fatalf("Response status code should not be %d", http.StatusUnauthorized)
	}
}

func TestServeHTTP_WrongAdminAuthToken(t *testing.T) {
	handler := FakeHandler{}
	adminAuthToken := "adminabc"
	path := "pastry"
	method := "GET"

	request, _ := http.NewRequest(method, path, nil)
	request.Header.Add("AuthToken", "nsdoivnws")
	auth := NewAuth(handler, "abc", adminAuthToken, []Endpoint{})
	responseWriter := httptest.NewRecorder()

	auth.ServeHTTP(responseWriter, request)

	if responseWriter.Code != http.StatusUnauthorized {
		t.Fatalf("Response status code should be %d not %d", http.StatusUnauthorized, responseWriter.Code)
	}
}

func TestServeHTTP_CorrectAdminAuthToken(t *testing.T) {
	handler := FakeHandler{}
	adminAuthToken := "adminabc"
	path := "pastry"
	method := "GET"

	request, _ := http.NewRequest(method, path, nil)
	request.Header.Add("AuthToken", adminAuthToken)
	auth := NewAuth(handler, "abc", adminAuthToken, []Endpoint{})
	responseWriter := httptest.NewRecorder()

	auth.ServeHTTP(responseWriter, request)

	if responseWriter.Code == http.StatusUnauthorized {
		t.Fatalf("Response status code should not be %d", http.StatusUnauthorized)
	}
}

type FakeHandler struct {
}

func (FakeHandler) ServeHTTP(http.ResponseWriter, *http.Request) {}
