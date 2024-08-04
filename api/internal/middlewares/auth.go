package auth

import (
	"net/http"
)

type Auth struct {
	handler           http.Handler
	Token             string
	AdminToken        string
	NonAdminEndpoints []Endpoint
}

type Endpoint struct {
	Path   string
	Method string
}

func (auth *Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("AuthToken")
	if token == "" {
		http.Error(w, "missing auth token", http.StatusUnauthorized)
		return
	}

	if isAdminEndpoint(auth, r.URL.Path, r.Method) && token != auth.AdminToken {
		http.Error(w, "invalid auth token", http.StatusUnauthorized)
		return
	} else if token != auth.Token && token != auth.AdminToken {
		http.Error(w, "invalid auth token", http.StatusUnauthorized)
		return
	}

	auth.handler.ServeHTTP(w, r)
}

func NewAuth(handlerToWrap http.Handler,
	token string,
	adminToken string,
	notAdminEndpoints []Endpoint) *Auth {
	return &Auth{handlerToWrap, token, adminToken, notAdminEndpoints}
}

func isAdminEndpoint(auth *Auth, path string, method string) bool {
	for i := 0; i < len(auth.NonAdminEndpoints); i++ {
		if auth.NonAdminEndpoints[i].Path == path &&
			auth.NonAdminEndpoints[i].Method == method {
			return false
		}
	}

	return true
}
