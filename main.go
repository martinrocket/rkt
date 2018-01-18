package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//root directory html page. Users will follow link to start using the application in the correct folder.
var myPage = `

<p>RocketPOS - To start POS, please follow this link <a href="http://localhost:8080/pointofsale/">http://localhost:8080/pointofsale/</a></p>

`

/* using this web site to build app
 * https://www.sohamkamani.com/blog/2017/09/13/how-to-build-a-web-application-in-golang/
 *
 */

// posRouter builds a new Router that will be started by the webserver
func posRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/callme/", handler2).Methods("GET")
	r.HandleFunc("/callme2/", handler4).Methods("GET")
	r.HandleFunc("/pointofsale/", handler3).Methods("GET") // alows GET from webserver root directory
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	return r
}

func main() {
	// create the router and start Listen and Server
	r := posRouter()
	http.ListenAndServe(":8080", r)
	fmt.Println("RKTPOS running...")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "You called me!")
}

func handler4(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Now you called me!")
}

func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "clalled POS")
}
