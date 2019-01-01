package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	sess *mgo.Session
)

type student struct {
	ID    bson.ObjectId `json:"id" bson:"_id"`
	Fio   string        `json:"fio" bson:"fio"`
	Info  string        `json:"info" bson:"info"`
	Score int           `json:"score" bson:"score"`
}

func main() {
	var err error
	sess, err = mgo.Dial("mongodb://localhost")
	PanicOnErr(err)

	// если коллекции не будет, то она создасться автоматически
	collection := sess.DB("msu-go-11").C("students")

	index := mgo.Index{
		Key: []string{"fio"},
	}
	err = collection.EnsureIndex(index)
	PanicOnErr(err)

	// для монги нет такого красивого дампа SQL, так что я вставляю демо-запись если коллекция пуста
	if n, _ := collection.Count(); n == 0 {
		firstStudent := &student{bson.NewObjectId(), "Vasily Romanov", "work: mail.ru group", 10}
		err = collection.Insert(firstStudent)
		PanicOnErr(err)
	}

	var allStudents []student
	// bson.M{} - это типа условия для поиска
	err = collection.Find(bson.M{}).All(&allStudents)
	PanicOnErr(err)
	for i, v := range allStudents {
		fmt.Printf("student[%d]: %+v\n", i, v)
	}

	//генерим какой-то ИДшник
	id := bson.NewObjectId()
	// bson.M{"_id": id} - задание условия для поиска
	var nonExistentSturent student
	err = collection.Find(bson.M{"_id": id}).One(&nonExistentSturent)
	PanicOnErr(err)

	secondStudent := &student{id, "Ivan Ivanov", "", 0}
	err = collection.Insert(secondStudent)
	PanicOnErr(err)

	err = collection.Find(bson.M{"_id": id}).One(&nonExistentSturent)
	PanicOnErr(err)
	fmt.Printf("Second Student: %+v\n", nonExistentSturent)

	collection.UpdateAll(
		bson.M{"fio": "Ivan Ivanov"},
		bson.M{
			"$set": bson.M{"Info": "all Ivans info"},
		},
	)

	secondStudent.Info = "single record"
	collection.Update(bson.M{"_id": secondStudent.ID}, &secondStudent)

	err = collection.Find(bson.M{"_id": id}).One(&nonExistentSturent)
	PanicOnErr(err)
	fmt.Printf("Second Student after update: %+v\n", nonExistentSturent)

}

//PanicOnErr panics on error
func PanicOnErr(err error) {
	if err == mgo.ErrNotFound {
		fmt.Println("Record not found")
	} else if err != nil {
		PanicOnErr(err)
	}

}
