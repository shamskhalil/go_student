package models

type Student struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

var count rune = -1
var studentStore []Student = []Student{}

func AddOne(newStudent Student) rune {
	count++
	studentStore = append(studentStore, newStudent)
	return count
}

func GetOne(id int) Student {
	return studentStore[id]
}

func GetAll() []Student {
	return studentStore
}

func DeleteOne(id int) {
	studentStore = append(studentStore[:id], studentStore[id+1:]...)
}

func UpdateOne(id int, newStudent Student) {
	studentStore[id] = newStudent
}
