package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PostgresConnect() {
	_, err := gorm.Open(postgres.Open(""), &gorm.Config{})
	if err != nil {
		panic("Could not connect to PostgresSQL")
	}
}
