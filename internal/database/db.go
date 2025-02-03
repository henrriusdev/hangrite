package database

import (
	"log"
	"os"
	"path/filepath"

	"github.com/henrriusdev/hangrite/config"
	"github.com/henrriusdev/hangrite/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() (*gorm.DB, error) {
	dbDir := filepath.Dir(config.Env.DBPath)

	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		log.Printf("Failed to create directory: %v", err)
		return nil, err
	}

	db, err := gorm.Open(sqlite.Open(config.Env.DBPath), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to open database: %v", err)
		return nil, err
	}

	// Auto migrate the schema
	if err := db.AutoMigrate(&models.User{}, &models.Player{}, &models.Certificate{}); err != nil {
		log.Printf("Failed to auto migrate: %v", err)
		return nil, err
	}

	// Create default admin user if none exists
	var count int64
	if err := db.Model(&models.User{}).Where("role = ?", models.RoleAdmin).Count(&count).Error; err != nil {
		log.Printf("Failed to count admin users: %v", err)
		return nil, err
	}

	if count == 0 {
		// Create default admin user
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("Failed to generate admin password: %v", err)
			return nil, err
		}

		adminUser := models.User{
			Username: "admin",
			Password: string(hashedPassword),
			Role:     models.RoleAdmin,
		}

		if err := db.Create(&adminUser).Error; err != nil {
			log.Printf("Failed to create admin user: %v", err)
			return nil, err
		}
		log.Println("Successfully created default admin user")
	}

	DB = db
	return db, nil
}
