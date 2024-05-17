package service

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB = nil

func createDatabase() *gorm.DB {
	//! Important! Append "parseTime=true" as a param in the DSN, otherwise there are bugs related to time conversion
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func GetDatabase() *gorm.DB {
	if database == nil {
		database = createDatabase()
		Migrate(database)
	}

	return database
}
