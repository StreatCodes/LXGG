package main

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

func validateJWT(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil {
			http.Error(w, "Can't read session cookie, have you logged in?", http.StatusForbidden)
			return
		}

		var user User
		token, err := jwt.ParseWithClaims(cookie.Value, &user, func(token *jwt.Token) (interface{}, error) {
			return []byte("8asdnSJ871SKJN8*6asdj 12n3k12j3n9as87cha89&"), nil
		})

		//check token is valid
		if err != nil || !token.Valid {
			//Delete cookie
			//TODO make sure this is working
			cookie.MaxAge = -1
			http.SetCookie(w, cookie)

			http.Error(w, "Couldn't verify token integrity.", http.StatusForbidden)
			return
		}

		fmt.Printf("%v %v %v", user.User, user.Admin, user.StandardClaims.ExpiresAt)

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
