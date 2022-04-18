package models

import (
	"github.com/jinzhu/gorm"
	"github.com/rest-go-example/pkg/config"
)

var db *gorm.DB

type Person struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       string `json:"age"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Person{})
}

func (p *Person) CreatePerson() *Person {
	db.NewRecord(p)
	db.Create(&p)
	return p
}

func GetAllPerson() []Person {
	var Person []Person
	db.Find(&Person)
	return Person
}

func GetPersonByID(Id int64) (*Person, *gorm.DB) {
	var getPerson Person
	db := db.Where("ID=?", Id).Find(&getPerson)
	return &getPerson, db
}

func DeletePerson(Id int64) Person {
	var person Person
	db.Where("ID=?", Id).Delete(person)
	return person
}
