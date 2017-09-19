package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jmoiron/sqlx"
)

//LXGGDB Global DB var
var LXGGDB *sqlx.DB

//JWTSECRET Global secret for jwt signing
var JWTSECRET []byte

func main() {
	settings := loadSettings()
	LXGGDB = loadDB()

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
		r.Post("/containers/new", newContainerHandler)
	})

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "/static")
	staticServer(r, "/", http.Dir(filesDir))

	fmt.Println("Starting server ", settings.serverAddr())
	err := http.ListenAndServe(settings.serverAddr(), r)
	log.Fatal("Error starting server: ", err)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the LXGG api, hit up https://github.com/streatcodes/lxgg for some docs!"))
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
