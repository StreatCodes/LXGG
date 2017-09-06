package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", rootHandler)
	r.Post("/login", loginHandler)

	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.RequestID)
		r.Use(middleware.RealIP)
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)
		r.Use(validateJWT)

		r.Get("/containers", containersAllHandler)
	})

	http.ListenAndServe(":3001", r)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<html><body>Welcome to the LXGG api, hit up <a href=\"https://github.com/streatcodes/lxgg\">our github page for docs!</a></body></html>"))
}
