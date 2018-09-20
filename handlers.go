package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func TodoShow(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	todoId := vars["todoId"]
	fmt.Fprintln(writer, "Todo show:", todoId)
}

func TodoIndex(writer http.ResponseWriter, request *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	if err := json.NewEncoder(writer).Encode(todos); err != nil {
		panic(err)
	}
}

func Index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome!")
}
