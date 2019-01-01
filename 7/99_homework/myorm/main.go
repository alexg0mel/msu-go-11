package main

import (
	"./user"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// создаём базу
	var err error
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/msu-go-11?charset=utf8&interpolateParams=true")
	PanicOnErr(err)
	db.SetMaxOpenConns(10)
	err = db.Ping()
	PanicOnErr(err)

	user.SetDB(db)

	u := user.User{}
	err = u.FindByPK(1)
	PanicOnErr(err)
	fmt.Printf("найден по первичному ключу: %+v\n", u)

	// создаем новую запись
	newUser := user.User{
		Login: "Ivan Ivanov",
	}
	err = newUser.Create()
	PanicOnErr(err)
	fmt.Println(newUser.ID)

	// обновляем запись
	newUser.Info = "occupation: programmer"
	newUser.Balance = 10
	newUser.Status = 1
	newUser.SomeInnerFlag = true
	err = newUser.Update()
	PanicOnErr(err)

	u2 := user.User{}
	u2.FindByPK(newUser.ID)
	fmt.Printf("найден по первичному ключу после сохранения: %+v\n", u2)
}

//PanicOnErr panics on error
func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
