package main

import (
	"fmt"
	"net/http"
)

var myPage = `

<h1>RKTPOS, startin up....</h1>
<p>yep, here we go</p>
<p>yep yep</p>

`

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, myPage)
}

func main() {
	fmt.Println("Starting RKT......")
	fmt.Print("Select an option: ") // Prompt user to select from menu.
	var selection []byte            // Variable to hold selection from user input
	for {
		http.HandleFunc("/", handler)
		http.ListenAndServe(":8080", nil)
		fmt.Scanln(&selection) // Scan user input and assign it to memory location for var selection
		fmt.Println()
		if string(selection) == "x" {
			http.Shutdown()
		}
	}
}
