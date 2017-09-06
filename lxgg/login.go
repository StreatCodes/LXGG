package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

//User contains user information
type User struct {
	User  string
	Admin bool
	jwt.StandardClaims
}

type loginRequest struct {
	Username string
	Password string
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var loginReq loginRequest
	dec.Decode(&loginReq)

	if loginReq.Username == "lxgg" && loginReq.Password == "password" {
		//Create our JWT
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"User":  loginReq.Username,
			"Admin": true,
		})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString([]byte("8asdnSJ871SKJN8*6asdj 12n3k12j3n9as87cha89&"))
		if err != nil {
			http.Error(w, fmt.Sprintf("Error signing JWT: %s", err), http.StatusInternalServerError)
			return
		}

		//Create our cookie and send back a success login
		cookie := http.Cookie{Name: "session", Value: tokenString}
		http.SetCookie(w, &cookie)

		w.Write([]byte(`"Status":"Success"`))
	} else {
		w.Write([]byte(`"Status":"Invalid creds"`))
		http.Error(w, "Invalid creds", http.StatusUnauthorized)
	}
}
