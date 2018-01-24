package pointofsale

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// SellItem is a function that acts as a handler when /API/{sellItem}
func SellItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You sold an item %v \n", strings.Join(r.URL.Query()["sellItem"], " "))
	log.Println("func Sell Item")
	log.Println(r.URL.Query()["sellItem"])
	z := r.URL.Query()["sellItem"]
	log.Printf("id = %v of type %T type %T \n", z, z)
}
