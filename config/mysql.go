package config

import (
	"pusat-rumah-lelang-backend/constants"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenDB() *gorm.DB {
	dbURL := constants.DBUser + ":" + constants.DBPassword + "@(" + constants.DBHost + ":" + constants.DBPort + ")/" + constants.DBName + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect to database")
	}

	return db
}
