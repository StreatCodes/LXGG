package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/lxc/lxd/client"

	"github.com/lxc/lxd/shared/api"
)

//Container stores all our container info
type Container struct {
	ID      int64  `db:"id"`
	OwnerID int64  `db:"owner_id"`
	Name    string `db:"name"`
	Tags    string `db:"tags"`
	IP      string `db:"ip"`
	Image   string
}

func containersAllHandler(w http.ResponseWriter, r *http.Request) {
	// user := r.Context().Value(UserKey("user")).(User)

	//Get list of containers from LXD
	//TODO get params, check if all_users, then check user is admin, else owned containers
	containers, err := LXDCONN.GetContainers()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching containers from LXD: %s", err), http.StatusInternalServerError)
		return
	}

	//Response
	enc := json.NewEncoder(w)
	err = enc.Encode(containers)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding container response: %s", err), http.StatusInternalServerError)
		return
	}
	// if user.Admin {
	// 	w.Write([]byte("Fetching ALL container list..."))
	// } else {
	// 	w.Write([]byte("Fetching USER container list..."))
	// }

}

func newContainerHandler(w http.ResponseWriter, r *http.Request) {
	var container Container

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&container)

	fmt.Println("A")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding body: %s", err), http.StatusInternalServerError)
		return
	}

	//getimage from finger print
	//createcontainerfromimage
	//TODO properly validate LXD fields, (though lXD might do that for us)

	fmt.Println("B")
	imgconn, err := lxd.ConnectPublicLXD("https://us.images.linuxcontainers.org/", nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error connecting to public LXD server: %s", err), http.StatusInternalServerError)
		return
	}

	fmt.Println("C")
	img, _, err := imgconn.GetImage(container.Image)
	if err != nil {
		log.Println(container.Image)
		http.Error(w, fmt.Sprintf("Error fetching image corresponding to fingerprint: %s", err), http.StatusInternalServerError)
		return
	}

	op, err := LXDCONN.CreateContainerFromImage(imgconn, *img, api.ContainersPost{Name: container.Name})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating container: %s", err), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	fmt.Println("E")
	//Use an alternative method
	err = op.Wait()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error waiting for LXD response: %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Println("F")
	// id, err := createContainer(container)
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("Error creating container: %s", err), http.StatusInternalServerError)
	// 	return
	// }

	_, err = fmt.Fprint(w, `{"Status": "Success"}`)
	if err != nil {
		log.Println("Error writing success response: ", err)
		return
	}
	fmt.Println("G")
}
