package config

import (
	"gorm.io/gorm"
)

type Config struct {
	DB            *gorm.DB
	Port          string
	Static        string
	SessionSecret []byte
}

func New(db *gorm.DB) *Config {
	return &Config{
		Port:          ":3000",
		Static:        "assets",
		SessionSecret: []byte("your-secret-key-change-this-in-production"),
		DB:            db,
	}
}
