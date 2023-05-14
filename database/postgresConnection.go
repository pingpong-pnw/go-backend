package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/pingpong-pnw/go-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func PostgresConnect() {
	if err := godotenv.Load(); err != nil {
		panic("Could not load configuration for database")
	}

	host := os.Getenv("postgresHost")
	port := os.Getenv("postgresPort")
	user := os.Getenv("postgresUser")
	password := os.Getenv("postgresPassword")
	name := os.Getenv("postgresName")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, name, port)
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		panic("Could not connect to PostgresSQL")
	}
	DB = connection
	connection.AutoMigrate(models.Users{})
}
