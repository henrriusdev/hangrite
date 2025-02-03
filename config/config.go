package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type EnvVar struct {
	Port          string `env:"PORT" envDefault:":3000"`
	Static        string `env:"STATIC" envDefault:"assets"`
	SessionSecret string `env:"SESSION_SECRET" envDefault:"your-secret-key-change-this-in-production"`
	DBPath        string `env:"DB_PATH" envDefault:"./data/hangrite.db"`
}

var Env EnvVar

type Config struct {
	DB *gorm.DB
}

func getEnvPath() string {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	return filepath.Join(basepath, "..", ".env")
}

func LoadEnv() {
	if err := godotenv.Load(getEnvPath()); err != nil {
		log.Println("No .env file found")
	}

	if err := env.Parse(&Env); err != nil {
		log.Fatalf("Failed to parse env: %v", err)
	}
}

func New(db *gorm.DB) *Config {
	return &Config{
		DB: db,
	}
}
