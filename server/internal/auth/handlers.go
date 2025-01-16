package auth

import (
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	serv *AuthService
}

func NewAuthHandler(serv *AuthService) *AuthHandler {
	return &AuthHandler{serv}
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	loginPayload := LoginPayload{}

	// tries to decode the body data to payload variable
	if err := json.NewDecoder(r.Body).Decode(&loginPayload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// gets the jwt token
	token, err := h.serv.LoginService(loginPayload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// sets the token as a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
	})

	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	registerPayload := RegisterPayload{}

	// tries to decode the body data to payload variable
	if err := json.NewDecoder(r.Body).Decode(&registerPayload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// tries to register the user
	if err := h.serv.RegisterService(registerPayload); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
