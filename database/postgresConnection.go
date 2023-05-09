package database

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PostgresConnect() {
	if err := godotenv.Load(); err != nil {
		panic("Could not load configuration for database")
	}

	host := os.Getenv("postgresHost")
	port := os.Getenv("postgresPort")
	user := os.Getenv("postgresUser")
	password := os.Getenv("postgresPassword")
	name := os.Getenv("postgresName")

	dsn := "host=" + host + "user=" + user + "password=" + password + "dbname=" + name + "port=" + port + "sslmode=disable"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to PostgresSQL")
	}
}
