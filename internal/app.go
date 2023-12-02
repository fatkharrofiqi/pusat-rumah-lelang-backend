package internal

import (
	"pusat-rumah-lelang-backend/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDatabase() (*gorm.DB, error) {
	config, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	dbURL := config.DBUser + ":" + config.DBPassword + "@(" + config.DBHost + ":" + config.DBPort + ")/" + config.DBName + "?charset=utf8&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return DB, nil
}
