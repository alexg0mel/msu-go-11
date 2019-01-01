package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

type Student struct {
	ID    uint `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Fio   string
	Info  string
	Score int
}

func (u *Student) TableName() string {
	return "students"
}

func (u *Student) BeforeSave() (err error) {
	fmt.Println("trigger on before save")
	return
}

// PrintByID print student by id
func PrintByID(id uint) {
	st := Student{}
	err := db.Find(&st, id).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("Record not found", id)
	} else {
		PanicOnErr(err)
	}
	fmt.Printf("PrintByID: %+v, data: %+v\n", id, st)
}

func main() {
	var err error

	db, err = gorm.Open("mysql", "root@tcp(localhost:3306)/msu-go-11?charset=utf8")
	PanicOnErr(err)
	defer db.Close()

	//непосредственно подключаемся к базе
	db.DB()
	db.DB().Ping()

	//выбираем одиночную запись
	PrintByID(1)
	PrintByID(100500)

	// выборка по всем студентам
	allStudents := []Student{}
	db.Find(&allStudents)
	for i, v := range allStudents {
		fmt.Printf("students[%d] %+v\n", i, v)
	}

	// создаем новую запись
	newStudent := Student{
		Fio: "Ivan Ivanov",
	}
	db.Create(&newStudent)
	fmt.Println(newStudent.ID)
	PrintByID(newStudent.ID)

	// обновляем запись
	newStudent.Info = "occupation: programmer"
	newStudent.Score = 10
	db.Save(newStudent)
	PrintByID(newStudent.ID)

	return

}

//PanicOnErr panics on error
func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
