package config

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(viper *viper.Viper, log *logrus.Logger) *gorm.DB {
	username := viper.GetString("DB_USERNAME")
	password := viper.GetString("DB_PASSWORD")
	host := viper.GetString("DB_HOST")
	port := viper.GetInt("DB_PORT")
	database := viper.GetString("DB_NAME")
	idleConnection := viper.GetInt("DB_IDLE_CONNECTION")
	maxConnection := viper.GetInt("DB_MAX_CONNECTION")
	maxLifeTimeConnection := viper.GetInt("DB_MAX_LIFE_TIME_CONNECTION")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(&logrusWriter{
			Logger: log,
		}, logger.Config{
			SlowThreshold:             time.Second * 5,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			LogLevel:                  logger.Info,
		}),
	})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	connection, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	connection.SetMaxIdleConns(idleConnection)
	connection.SetMaxOpenConns(maxConnection)
	connection.SetConnMaxLifetime(time.Second * time.Duration(maxLifeTimeConnection))

	return db
}

type logrusWriter struct {
	Logger *logrus.Logger
}

func (l *logrusWriter) Printf(message string, args ...interface{}) {
	l.Logger.Tracef(message, args...)
}
