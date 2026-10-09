package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg *Config) *gorm.DB {
	db, err := gorm.Open(
		postgres.Open(cfg.Dsn),
		&gorm.Config{
			TranslateError: true,
		},
	)

	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	log.Println("Database connection successful")

	return db
}