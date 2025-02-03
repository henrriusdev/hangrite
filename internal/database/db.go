package database

import (
	"github.com/henrriusdev/hangrite/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("hangrite.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate the schema
	if err := db.AutoMigrate(&models.User{}, &models.Player{}, &models.Certificate{}); err != nil {
		return nil, err
	}

	// Create default admin user if none exists
	var count int64
	if err := db.Model(&models.User{}).Where("role = ?", models.RoleAdmin).Count(&count).Error; err != nil {
		return nil, err
	}

	if count == 0 {
		// Create default admin user
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}

		adminUser := models.User{
			Username: "admin",
			Password: string(hashedPassword),
			Role:     models.RoleAdmin,
		}

		if err := db.Create(&adminUser).Error; err != nil {
			return nil, err
		}
	}

	DB = db
	return db, nil
}
