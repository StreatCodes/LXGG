package main

import "net/http"
import "encoding/json"
import "fmt"

//Container stores all our container info
type Container struct {
	Name     string
	Template string
	CPUS     int
	RAM      int
	Disk     int
}

func containersAllHandler(w http.ResponseWriter, r *http.Request) {
	// user := r.Context().Value(UserKey("user")).(User)

	//TODO use sqlite
	//TODO use SQLX to get user specif or all containers

	containers := []Container{
		Container{"Container 1", "Web", 1, 1024, 10240},
		Container{"Container 2", "Web", 1, 1024, 10240},
		Container{"Container 3", "IRC", 1, 512, 10240},
		Container{"Container 4", "Mail", 2, 1024, 10240},
		Container{"Container 5", "Mail", 2, 1024, 10240},
		Container{"Container 6", "Mail", 2, 1024, 10240},
		Container{"Container 7", "Monitoring", 1, 1024, 10240},
		Container{"Container 8", "Monitoring", 1, 1024, 10240},
	}

	enc := json.NewEncoder(w)
	err := enc.Encode(containers)
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
