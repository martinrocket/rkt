package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	pos "github.com/martinrocket/rkt/pointofsale"
)

//root directory html page. Users will follow link to start using the application in the correct folder.
var myPage = `

<p>RocketPOS - To start POS, please follow this link <a href="http://localhost:8080/pointofsale/">http://localhost:8080/pointofsale/</a></p>

`

// posRouter builds a new Router that will be started by the webserver
func posRouter() *mux.Router {
	r := mux.NewRouter()
	log.Println("...")
	r.HandleFunc("/API/{id:[10]}", handler1)
	r.HandleFunc("/API/2/", handler2).Methods("GET")
	r.HandleFunc("/API/3/", handler3).Methods("GET")
	r.HandleFunc("/API/{sellItem}", pos.SellItem).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	return r
}

func main() {
	// create the router and start Listen and Server
	r := posRouter()

	http.ListenAndServe(":8080", r)

	log.Println("RKTPOS running...")

}

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "You called me! \nh1")
	log.Println("handler1")
	log.Println(r.URL.Path)
	log.Println(r.URL.Query()["id"])
	z := r.URL.Query()["id"]
	log.Printf("id = %v of type %T type %T \n", z, z)
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "You called me! h2")
	log.Println("handler2")

}

func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "You called me! h3")
	log.Println("handler3")
}
