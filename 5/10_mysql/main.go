package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	println("on start")
}

var (
	db *sql.DB
)

// PrintByID print student by id
func PrintByID(id int64) {
	var fio string
	var info sql.NullString
	// var info string
	var score int
	row := db.QueryRow("SELECT fio, info, score FROM students WHERE id = ?", id)
	// fmt.Println(row)
	err := row.Scan(&fio, &info, &score)
	PanicOnErr(err)
	fmt.Println("PrintByID:", id, "fio:", fio, "info:", info, "score:", score)
}

func main() {
	var err error
	// создаём структуру базы
	// но соединение происходит только при мервом запросе
	db, err = sql.Open("mysql", "root@tcp(localhost:3306)/msu-go-11?charset=utf8&interpolateParams=true")
	PanicOnErr(err)

	db.SetMaxOpenConns(10)

	// проверяем что подключение реально произошло ( делаем запрос )
	err = db.Ping()
	PanicOnErr(err)

	// итерируемся по многим записям
	// Exec исполняет запрос и возвращает записи
	rows, err := db.Query("SELECT fio, score FROM students")
	PanicOnErr(err)
	for rows.Next() {
		var fio string
		var score string
		err = rows.Scan(&fio, &score)
		PanicOnErr(err)
		fmt.Println("rows.Next fio: ", fio, score)
	}
	// надо закрывать соединения, иначе будем течь
	rows.Close()

	var fio string
	row := db.QueryRow("SELECT fio FROM students WHERE id = 1")
	err = row.Scan(&fio)
	PanicOnErr(err)
	fmt.Println("db.QueryRow fio: ", fio)

	// Exec исполняет запрос и возвращает сколько строк было затронуто и последнйи ИД вставленной записи
	// символ ? является placeholder-ом. все последующие значения авто-экранируются и подставляются с правильным кавычками
	result, err := db.Exec(
		"INSERT INTO students (`fio`, `score`) VALUES (?, 0)",
		"Ivan Ivanov",
	)
	PanicOnErr(err)

	affected, err := result.RowsAffected()
	PanicOnErr(err)
	lastID, err := result.LastInsertId()
	PanicOnErr(err)

	fmt.Println("Insert - RowsAffected", affected, "LastInsertId: ", lastID)

	PrintByID(lastID)

	// Exec исполняет запрос и возвращает сколько строк было затронуто и последнйи ИД вставленной записи ( 0 в данном случае )
	// символ ? является placeholder-ом. все последующие значения авто-экранируются и подставляются с правильным кавычками
	result, err = db.Exec(
		"UPDATE students SET info = ? WHERE id = ?",
		"test user",
		lastID,
	)
	PanicOnErr(err)

	affected, err = result.RowsAffected()
	PanicOnErr(err)
	fmt.Println("Update - RowsAffected", affected)

	PrintByID(lastID)

	// использование prepared statements
	// Prepare подготовливает запись ( шлёт запрос на сервере, там он парсится и готов принимать данные)
	stmt, err := db.Prepare("UPDATE students SET info = ?, score = ? WHERE id = ?")
	PanicOnErr(err)
	// Exec для prepares statement отправляет даныне на сервер - там запрос уже распашрен, только исполняется с новыми данными
	result, err = stmt.Exec("prapared statements update", lastID, lastID)
	PanicOnErr(err)
	result, err = stmt.Exec("8 update", lastID, 8)
	PanicOnErr(err)

	affected, err = result.RowsAffected()
	PanicOnErr(err)
	fmt.Println("Update - RowsAffected", affected)

	PrintByID(lastID)

	return

	fmt.Println("OpenConnections", db.Stats().OpenConnections)

}

//PanicOnErr panics on error
func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
