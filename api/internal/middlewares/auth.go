package auth

import (
	"net/http"
)

type Auth struct {
	handler http.Handler
	Token   string
}

func (auth *Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("AuthToken")
	if token == "" {
		http.Error(w, "missing auth token", http.StatusUnauthorized)
		return
	}

	if token != auth.Token {
		http.Error(w, "invalid auth token", http.StatusUnauthorized)
		return
	}

	auth.handler.ServeHTTP(w, r)
}

func NewAuth(handlerToWrap http.Handler, token string) *Auth {
	return &Auth{handlerToWrap, token}
}
