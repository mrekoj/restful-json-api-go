package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
)

func TodoShow(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	todoId := vars["todoId"]
	fmt.Fprintln(writer, "Todo show:", todoId)
}

func TodoIndex(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json;charset=UTF-8")
	writer.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(writer).Encode(todos); err != nil {
		panic(err)
	}
}

func Index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome!")
}

func TodoCreate(writer http.ResponseWriter, request *http.Request) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(request.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := request.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &todo); err != nil {
		writer.Header().Set("Content-Type", "application/json;charset=UTF-8")
		writer.WriteHeader(422)
		if err := json.NewEncoder(writer).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateTodo(todo)
	writer.Header().Set("Content-Type", "application/json;charset=UTF-8")
	writer.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(writer).Encode(t); err != nil {
		panic(err)
	}
}
