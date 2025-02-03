package models

import (
	"time"

	"gorm.io/gorm"
)

type Certificate struct {
	gorm.Model
	PlayerID      uint
	Player        Player
	IssueDate     time.Time
	TrainingHours string
}
