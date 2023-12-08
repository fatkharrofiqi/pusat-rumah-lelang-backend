package config

import (
	"log"
	"pusat-rumah-lelang-backend/constants"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenDB() *gorm.DB {
	dbURL := constants.DBUser + ":" + constants.DBPassword + "@(" + constants.DBHost + ":" + constants.DBPort + ")/" + constants.DBName + "?charset=utf8&parseTime=True&loc=Local"
	log.Print(dbURL)
	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	return db
}
