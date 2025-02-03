package models

import (
	"time"

	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Name         string        `json:"name"`
	IDNumber     string        `json:"id_number" gorm:"uniqueIndex"`
	Position     string        `json:"position"`
	EntryDate    time.Time     `json:"entry_date"`
	Certificates []Certificate `json:"certificates"`
}

type PlayerForm struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	IDNumber  string    `json:"id_number"`
	Position  string    `json:"position"`
	EntryDate time.Time `json:"entry_date"`
	HasError  bool      `json:"-"`
	ErrorMsg  string    `json:"-"`
}
