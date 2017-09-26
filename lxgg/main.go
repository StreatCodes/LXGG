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
	"github.com/lxc/lxd/client"
)

//LXGGDB Global DB var
var LXGGDB *sqlx.DB

//LXDCONN Global LXD connection var
var LXDCONN lxd.ContainerServer

//JWTSECRET Global secret for jwt signing
var JWTSECRET []byte

func main() {
	settings := loadSettings()
	LXGGDB = loadDB()

	var err error
	LXDCONN, err = lxd.ConnectLXDUnix("", nil)
	if err != nil {
		log.Fatal("Could not connect to LXD: ", err)
	}

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
		r.Get("/images", imagesAllHandler)
		r.Post("/containers/new", newContainerHandler)
	})

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "/static")
	staticServer(r, "/", http.Dir(filesDir))

	fmt.Println("Starting server ", settings.serverAddr())
	err = http.ListenAndServe(settings.serverAddr(), r)
	log.Fatal("Error starting server: ", err)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Welcome to the LXGG api, hit up https://github.com/streatcodes/lxgg for some docs!"))
	if err != nil {
		log.Println("Error writing response: ", err)
		return
	}
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
