package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	lxd "github.com/lxc/lxd/client"
)

func imagesAllHandler(w http.ResponseWriter, r *http.Request) {
	imgconn, err := lxd.ConnectPublicLXD("https://us.images.linuxcontainers.org/", nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error connecting to public LXD server: %s", err), http.StatusInternalServerError)
		return
	}

	images, err := imgconn.GetImages()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching image list from LXD: %s", err), http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)
	err = enc.Encode(images)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding image list: %s", err), http.StatusInternalServerError)
		return
	}
}
