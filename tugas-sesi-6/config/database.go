package config

import (
	"fmt"
	"os"
	"prakerja-tugas6/utils"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))

	var (
		host     = os.Getenv("DATABASE_HOST")
		user     = os.Getenv("DATABASE_USERNAME")
		password = os.Getenv("DATABASE_PASSWORD")
		dbName   = os.Getenv("DATABASE_NAME")
		timeZone = os.Getenv("DATABASE_TIMEZONE")
	)

	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=%s", host, port, user, password, dbName, timeZone)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	utils.ErrorPanic(err)

	return db
}
