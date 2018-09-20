package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)

	log.Fatal(http.ListenAndServe(":8080", router))
}
func Index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello mux, %q", html.EscapeString(request.URL.Path))
}
