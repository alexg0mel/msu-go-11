package main

import (
	"net/http"
	"../handlers"
	"fmt"
)

func main() {
	handler := handlers.Handler{
		Todos: &[]handlers.Todo{
			{"Выучить Go", false},
		},
	}

	http.HandleFunc("/", handler.HandleIndex)
	http.HandleFunc("/todos/", handler.HandleTodos)

	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println("Error: ", err.Error())
	}
}
