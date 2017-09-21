package main

import (
	"context"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func validateJWT(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("lxgg_session")
		if err != nil {
			http.Error(w, "Can't read lxgg_session cookie, have you logged in?", http.StatusForbidden)
			return
		}

		//TODO make a new type for this so it doesn't have all the JWT junk in it down the line
		var user User
		token, err := jwt.ParseWithClaims(cookie.Value, &user, func(token *jwt.Token) (interface{}, error) {
			return JWTSECRET, nil
		})

		//check token is valid
		if err != nil || !token.Valid {
			//Delete cookie
			//TODO make sure this is working
			delCookie := http.Cookie{Path: "/", Name: "lxgg_session", Value: "", Expires: time.Now().Add(-100 * time.Hour)}
			http.SetCookie(w, &delCookie)

			http.Error(w, "Couldn't verify token integrity.", http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), UserKey("user"), user)
		h.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
