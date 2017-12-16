package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>RKTPOS, startin up....</h1><p>yep, here we go</p>")
}

func main() {
	fmt.Println("Starting RKT......")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
