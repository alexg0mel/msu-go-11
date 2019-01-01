package main

import (
	"io/ioutil"
	"net/http"
	"log"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		// Создаем http клиент. В стуктуру можно передать таймаут, куки и прочую инфу о запросе
		c := http.Client{}
		resp, err := c.Get("http://artii.herokuapp.com/make?text=" + path)
		if err != nil {
			log.Println(err)
		}
		// нужно закрыть тело, когда прочитаем что нужно
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		// статус - ОК
		w.WriteHeader(200)
		w.Write(body)
	})

	http.ListenAndServe(":8081", nil)
}
