package db

import (
	"fmt"
	"gorm.io/gorm"
)

var Database *gorm.DB

type Config struct {
	Host         string
	User         string
	Password     string
	Port         string
	DatabaseName string
}

func InitConfig() *Config {
	dbConfig := Config{
		Host:         "database",
		Port:         "5432",
		User:         "postgres",
		Password:     "123",
		DatabaseName: "microservice",
	}
	return &dbConfig
}

func DbURL(dbConfig *Config) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DatabaseName,
		dbConfig.Port,
	)
}
