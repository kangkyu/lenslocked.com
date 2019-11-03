package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "contact us if you have any question.")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/contact/", contact)
	r.HandleFunc("/", home)
	http.ListenAndServe(":3000", r)
}
