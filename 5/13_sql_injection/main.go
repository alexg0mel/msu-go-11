package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

var (
	db *sql.DB
)

var loginFormTmpl = `
<html>
	<body>
	<form action="/login" method="post">
		Login: <input type="text" name="login">
		Password: <input type="password" name="password">
		<input type="submit" value="Login">
	</form>
	</body>
</html>
`

func main() {

	var err error
	// создаём структуру базы
	// но соединение происходит только при мервом запросе
	// db, err = sql.Open("mysql", "root@tcp(localhost:3360)/msu-go-11?charset=utf8")
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/msu-go-11?charset=utf8")
	PanicOnErr(err)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(loginFormTmpl))
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		var id int
		var login string
		var body string

		// парсим POST
		r.ParseForm()

		// fmt.Printf("Form data: %+v\n", r.Form)

		inputLogin := r.Form["login"][0]
		body += fmt.Sprintln("inputLogin:", inputLogin)

		// ПЛОХО! НЕ ДЕЛАЙТЕ ТАК!
		// параметры не экранированы должным образом
		// мы подставляем в запрос параметр как есть
		query := fmt.Sprintf("SELECT id, login FROM users WHERE login = '%s' LIMIT 1", inputLogin)
		body += fmt.Sprintln("Sprint query:", query)
		row := db.QueryRow(query)
		err := row.Scan(&id, &login)

		if err == sql.ErrNoRows {
			body += fmt.Sprintln("Sprint case: NOT FOUND")
		} else {
			PanicOnErr(err)
			body += fmt.Sprintln("Sprint case: id:", id, "login:", login)
		}

		// ПРАВИЛЬНО
		// Мы используем плейсхолдеры, там параметры будет экранирован должным образом
		row = db.QueryRow("SELECT id, login FROM users WHERE login = ? LIMIT 1", inputLogin)
		err = row.Scan(&id, &login)
		if err == sql.ErrNoRows {
			body += fmt.Sprintln("Placeholders case: NOT FOUND")
		} else {
			PanicOnErr(err)
			body += fmt.Sprintln("Placeholders id:", id, "login:", login)
		}

		w.Write([]byte(body))
	})

	http.ListenAndServe(":8081", nil)
}

//PanicOnErr panics on error
func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
