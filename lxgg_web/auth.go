package main

import "net/http"

func validateJWT(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		//TODO
		// if rip := realIP(r); rip != "" {
		// 	r.RemoteAddr = rip
		// }
		// h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
