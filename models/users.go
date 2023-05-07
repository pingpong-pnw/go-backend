package models

type Users struct {
	id        uint
	firstName string
	lastName  string
	email     string `gorm:"unique"`
	password  string
}
