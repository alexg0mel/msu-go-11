package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type TestUser struct {
	ID       int
	Username string
}

func main() {
	db, err := gorm.Open("mysql", "stuff:stuff@/data?charset=utf8&parseTime=True&loc=Local")

	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	// After db connection is created.
	db.CreateTable(&TestUser{})

	// Also some useful functions
	fmt.Println("Table created ", db.HasTable(&TestUser{})) // => true
	db.DropTable(&TestUser{})

	defer db.Close()

	complex(db)
}

func complex(db *gorm.DB) {
	if !db.HasTable(&User{}) {
		db.CreateTable(&User{})
	}

	//fill(db)
	query(db)
}

func fill(db *gorm.DB) {
	user := User{
		FirstName: "Arthur",
		LastName:  "Dent",
		Username:  "adent",
		Salary:    5000,
	}

	db.Create(&user)

	// Seeding tables:
	var users []User = []User{
		User{Username: "foobar", FirstName: "Foo", LastName: "Bar", Salary: 200},
		User{Username: "helloworld", FirstName: "Hello", LastName: "World", Salary: 200},
		User{Username: "john", FirstName: "John", Salary: 200},
	}

	for _, user := range users {
		db.Create(&user)
	}
}

func query(db *gorm.DB) {
	// Bind data to variable u with conditions
	u := User{FirstName: "Arthur"}
	//db.Where(&u).First(&u)
	// OR
	//db.Find(&u)
	fmt.Println(u)

	// Query multiple records
	users := []User{}
	db.Where(&User{Salary: 200}).Find(&users)
	fmt.Println(users)
}

type User struct {
	// auto-populate columns: id, created_at, updated_at, deleted_at
	gorm.Model

	// Set column type manually
	Username string `sql:"type:VARCHAR(255);not null;unique"`

	// Set default value
	LastName string `sql:"DEFAULT:'Smith'"`

	// Custom column name instead of default snake_case format
	FirstName string `gorm:"column:FirstName"`

	Role string

	Salary int64
}

func (u *User) TableName() string {
	// custom table name, this is default
	return "users"
}

// func (u *User) BeforeSave() (err error) {
// 	if u.Role != "admin" {
// 		err = errors.New("Permission denied.")
// 	}
// 	return
// }
