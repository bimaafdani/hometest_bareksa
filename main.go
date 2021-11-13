package main

import "HoteTestBareksa/interfaces"

// // open comment to deploy gcp app engine
// "google.golang.org/appengine"

func main() {
	// // open comment to deploy gcp app engine
	// router := interfaces.Routes()
	// http.Handle("/", router)
	// appengine.Main() // Start the gcp server

	interfaces.Run(8000)
}
