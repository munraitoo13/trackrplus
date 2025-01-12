package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var payload UserLogin

	// tries to decode the body data to payload variable
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// gets the jwt token
	token, err := Login(&payload)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}

	// sets the token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var payload UserRegister

	// tries to decode the body data to payload variable
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		fmt.Println(err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// tries to register the user
	if err := Register(&payload); err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to register new user", http.StatusBadRequest)
		return
	}
}
