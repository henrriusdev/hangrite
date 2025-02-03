package main

import (
	"log"

	"github.com/henrriusdev/hangrite/config"
	"github.com/henrriusdev/hangrite/internal/app"
	"github.com/henrriusdev/hangrite/internal/database"
)

func main() {
	db, err := database.Init()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	config := config.New(db)

	app := app.New(config)
	if err := app.Run(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
