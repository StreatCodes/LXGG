package main

import "net/http"
import "encoding/json"
import "fmt"

//Container stores all our container info
type Container struct {
	ID      int64  `db:"id"`
	OwnerID int64  `db:"owner_id"`
	Name    string `db:"name"`
	Tags    string `db:"tags"`
	IP      string `db:"ip"`
}

func containersAllHandler(w http.ResponseWriter, r *http.Request) {
	// user := r.Context().Value(UserKey("user")).(User)

	//TODO get params, check if all_users, then check user is admin, else owned containers
	containers, err := getAllContainers()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching containers from DB: %s", err), http.StatusInternalServerError)
		return
	}

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

	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding body: %s", err), http.StatusInternalServerError)
		return
	}

	//TODO properly validate LXD fields, (though lXD might do that for us)
	id, err := createContainer(container)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating container: %s", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, `{"Status": "Success", "ID":"%d"}`, id)
}
