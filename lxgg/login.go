package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

//UserKey is our type for passing the user context between middlewares
type UserKey string

//User contains user information
type User struct {
	ID    int64  `db:"id"`
	User  string `db:"username"`
	Pass  string `db:"password" json:"-"`
	Admin bool   `db:"admin"`
	jwt.StandardClaims
}

type loginRequest struct {
	Username string
	Password string
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var loginReq loginRequest
	err := dec.Decode(&loginReq)

	if err != nil {
		http.Error(w, "Error decoding body", http.StatusInternalServerError)
		log.Println("Error decoding login requst body: ", err)
		return
	}

	user, err := verifyLogin(w, loginReq.Username, loginReq.Password)
	if err != nil {
		http.Error(w, "Error verifying user", http.StatusInternalServerError)
		log.Println("Error verifying user: ", err)
		return
	}

	//Create our JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":    user.ID,
		"User":  user.User,
		"Admin": user.Admin,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(JWTSECRET)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error signing JWT: %s", err), http.StatusInternalServerError)
		return
	}

	//Create our cookie and send back a success login
	cookie := http.Cookie{Path: "/", Name: "lxgg_session", Value: tokenString}
	http.SetCookie(w, &cookie)

	_, err = w.Write([]byte(`"Status":"Success"`))
	if err != nil {
		log.Println("Error writing response: ", err)
		return
	}
}

func verifyLogin(w http.ResponseWriter, username, password string) (User, error) {
	var user []User

	//TODO make sure this isn't nil when empty
	err := LXGGDB.Select(&user, `SELECT * FROM users WHERE username=$1`, username)
	if err != nil {
		return User{}, err
	}

	if len(user) < 1 {
		return User{}, errors.New("No users associated with that username")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user[0].Pass), []byte(password))
	if err != nil {
		return User{}, err
	}

	return user[0], nil
}
