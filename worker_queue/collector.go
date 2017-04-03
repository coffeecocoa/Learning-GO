/*
receive client requests for work, builds work request that the
workers can understand, and pushes the work onto the the end of
the work queue
*/

package main

import (
	"fmt"
	"net/http"
	"time"
)

// a bufferd channel that we can send work request on
var WorkQueue = make(chan WorkRequest, 100)

func Collector(w http.ResponseWriter, r *http.Request) {
	//make sure we can only be called with an http post request
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	//parse the delay
	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Bad delay value: "+err.Error(), http.StatusBadRequest)
		return
	}

	//now, we retrieve the person's name from the request
	name := r.FormValue("name")

	//just do quick bit of sanity checking to make the client actually
	//provided us with a name
	if name == "" {
		http.Error(w, "You must specify a name.", http.StatusBadRequest)
		return
	}

	//now we take the delay and the person's name, and make workrequest out of them
	work := WorkRequest{Name: name, Delay: delay}

	//push the work onto the queue
	WorkQueue <- work
	fmt.Println("Work request queued")
	return
}
