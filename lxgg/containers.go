package main

import "net/http"

func containersAllHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(UserKey("user")).(User)

	if user.Admin {
		w.Write([]byte("Fetching ALL container list..."))
	} else {
		w.Write([]byte("Fetching USER container list..."))
	}
}
