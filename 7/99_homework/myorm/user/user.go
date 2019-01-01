package user

//myorm:users
type User struct {
	ID            uint   `myorm:"primary_key"`     // первичный ключ, в него мы пишем LastInsertId
	Login         string `myorm:"column:username"` // поле называется username в таблице
	Info          string `myorm:"null"`            // поле может иметь тип null
	Balance       int
	Status        int
	SomeInnerFlag bool `myorm:"-"` //поля нет в таблице, игнорируем его
}
