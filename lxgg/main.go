package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	// r.Get("/", rootHandler)
	r.Post("/login", loginHandler)

	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.RequestID)
		r.Use(middleware.RealIP)
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)
		r.Use(validateJWT)

		r.Get("/containers", containersAllHandler)
	})

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "/static")
	staticServer(r, "/", http.Dir(filesDir))

	http.ListenAndServe(":3001", r)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<html><body>Welcome to the LXGG api, hit up <a href=\"https://github.com/streatcodes/lxgg\">our github page for docs!</a></body></html>"))
}

func staticServer(r chi.Router, path string, root http.FileSystem) {
	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		fs.ServeHTTP(w, r)
	}))
}
