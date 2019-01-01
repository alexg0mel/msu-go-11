package handlers

import (
	"net/http"
	"fmt"
	"encoding/json"
	"log"
	"strconv"
)

func (h Handler) HandleTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL.Path)
	defer r.Body.Close()

	switch r.Method {
	case http.MethodGet:
		productsJson, _ := json.Marshal(h.Todos)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(productsJson)
	case http.MethodPost:
		decoder := json.NewDecoder(r.Body)
		todo := Todo{}
		// преобразуем json запрос в структуру
		err := decoder.Decode(&todo)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		*h.Todos = append(*h.Todos, todo)
	case http.MethodPut:
		id := r.URL.Path[len("/todos/"):]
		index, _ := strconv.ParseInt(id, 10, 0)
		(*h.Todos)[index].Done = true
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
