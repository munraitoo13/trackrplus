package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginPayload *LoginPayload

	// tries to decode the body data to payload variable
	if err := json.NewDecoder(r.Body).Decode(&loginPayload); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// gets the jwt token
	token, err := LoginService(loginPayload)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}

	// sets the token as a cookie
	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: token,
	})
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var registerPayload *RegisterPayload

	// tries to decode the body data to payload variable
	if err := json.NewDecoder(r.Body).Decode(&registerPayload); err != nil {
		fmt.Println(err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// tries to register the user
	if err := RegisterService(registerPayload); err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to register new user", http.StatusBadRequest)
		return
	}
}
