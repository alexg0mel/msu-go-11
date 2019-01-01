package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

// PrintByID print student by id
func PrintByID(id int64) {
	var fio string
	var info sql.NullString
	// var info string
	var score int
	row := db.QueryRow("SELECT fio, info, score FROM students WHERE id = $1", id)
	// fmt.Println(row)
	err := row.Scan(&fio, &info, &score)
	PanicOnErr(err)
	fmt.Println("PrintByID:", id, "fio:", fio, "info:", info, "score:", score)
}

func main() {
	var err error
	// создаём структуру базы
	// но соединение происходит только при мервом запросе
	db, err = sql.Open("postgres", "user=postgres dbname=msu-go-11 sslmode=disable")
	PanicOnErr(err)

	// проверяем что подключение реально произошло ( делаем запрос )
	err = db.Ping()
	PanicOnErr(err)

	// итерируемся по многим записям
	// Exec исполняет запрос и возвращает записи
	rows, err := db.Query("SELECT fio FROM students")
	PanicOnErr(err)
	for rows.Next() {
		var fio string
		err = rows.Scan(&fio)
		PanicOnErr(err)
		fmt.Println("rows.Next fio: ", fio)
	}
	rows.Close()

	var fio string
	row := db.QueryRow("SELECT fio FROM students WHERE id = 1")
	err = row.Scan(&fio)
	PanicOnErr(err)
	fmt.Println("db.QueryRow fio: ", fio)

	// Exec исполняет запрос и возвращает сколько строк было затронуто и последнйи ИД вставленной записи
	// символ ? является placeholder-ом. все последующие значения авто-экранируются и подставляются с правильным кавычками
	var lastID int64
	err = db.QueryRow(
		"INSERT INTO students (fio, score) VALUES ($1, 0) RETURNING id",
		"Ivan Ivanov",
	).Scan(&lastID)
	PanicOnErr(err)

	fmt.Println("Insert - LastInsertId: ", lastID)

	PrintByID(lastID)

	// Exec исполняет запрос и возвращает сколько строк было затронуто и последнйи ИД вставленной записи ( 0 в данном случае )
	// символ ? является placeholder-ом. все последующие значения авто-экранируются и подставляются с правильным кавычками
	result, err := db.Exec(
		"UPDATE students SET info = $1 WHERE id = $2",
		"test user",
		lastID,
	)
	PanicOnErr(err)

	affected, err := result.RowsAffected()
	PanicOnErr(err)
	fmt.Println("Update - RowsAffected", affected)

	PrintByID(lastID)

	// использование prepared statements
	// Prepare подготовливает запись ( шлёт запрос на сервере, там он парсится и готов принимать данные)
	stmt, err := db.Prepare("UPDATE students SET info = $1, score = $2 WHERE id = $2")
	PanicOnErr(err)
	// Exec для prepares statement отправляет даныне на сервер - там запрос уже распашрен, только исполняется с новыми данными
	result, err = stmt.Exec("prapared statements update", lastID)
	PanicOnErr(err)

	affected, err = result.RowsAffected()
	PanicOnErr(err)
	fmt.Println("Update - RowsAffected", affected)

	PrintByID(lastID)

	fmt.Println("OpenConnections", db.Stats().OpenConnections)

}

//PanicOnErr panics on error
func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
