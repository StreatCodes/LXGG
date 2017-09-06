package main

import "net/http"

func containersAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Fetching container list..."))
}
