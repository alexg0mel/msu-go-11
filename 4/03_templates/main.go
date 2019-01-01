package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Todo struct {
	Name string
	Done bool
}

func IsNotDone(todo Todo) bool {
	return !todo.Done
}

func main() {
	tmpl, err := template.New("template.html").Funcs(template.FuncMap{"IsNotDone": IsNotDone}).ParseFiles("template.html")
	if err != nil {
		log.Fatal("Can not expand template", err)
		return
	}

	todos := []Todo{
		{"Выучить Go", false},
		{"Посетить лекцию по вебу", false},
		{"...", false},
		{"Profit", false},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// читаем из urlencoded запроса
			param := r.FormValue("id")
			// преобразуем строку в int
			index, _ := strconv.ParseInt(param, 10, 0)
			todos[index].Done = true
		}

		// исполняем шаблон
		err := tmpl.Execute(w, todos)
		if err != nil {
			// вернем 500 и напишем ошибку
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":8081", nil)
}
