package main

import "net/http"

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You may login."))
}
