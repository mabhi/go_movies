package database

import (
	"fmt"
	"go-movies-be/src/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, username, password, database, port)
	// fmt.Printf("connection strng %s", dsn)
	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("Could not connect with the database! %s", err))
	}

}

func AutoMigrate() {
	if DB != nil {
		DB.AutoMigrate(models.Movie{}, models.Genre{})
	} else {
		fmt.Println("💀 Db is nil")
	}
}
