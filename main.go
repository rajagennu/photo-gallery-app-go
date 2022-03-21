package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")

}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "To get in touch with us, please send us email to "+
		" mailto:\"rajag@example.com\"")

}

func faq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<b>1. What is this website about ?</b><br>"+
		" this is a photoshoot gallery app for great collages")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1 style=\"align:right;\">Are you sure, you are on right Planet ?")
}

func main() {
	// http.HandleFunc("/", handlerFunc)
	// http.ListenAndServe(":3000", nil)
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.HandleFunc("/FAQ", faq)

	var h http.Handler = http.HandlerFunc(notFound)
	r.NotFoundHandler = h

	http.ListenAndServe(":3000", r)
}
