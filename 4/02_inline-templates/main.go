package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"log"
)

func main() {
	// создаем новый шаблон из строки
	tmpl := template.New("main")
	tmpl, _ = tmpl.Parse(
		`<div style="display: inline-block; border: 1px solid #aaa; border-radius: 3px; padding: 30px; margin: 20px;">
			<pre>{{.}}</pre>
		</div>`)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		c := http.Client{}
		resp, err := c.Get("http://artii.herokuapp.com/make?text=" + path)
		if err != nil {
			// при ошибке кидаем 500 и уходим
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error"))
			return
		}

		defer resp.Body.Close()

		// читаем весь Body
		body, _ := ioutil.ReadAll(resp.Body)

		// отдаем шаблон с данными
		tmpl.Execute(w, string(body))
	})

	http.ListenAndServe(":8081", nil)
}
