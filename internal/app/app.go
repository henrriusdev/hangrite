package app

import (
	"log"
	"net/http"

	"github.com/henrriusdev/hangrite/config"
	"github.com/henrriusdev/hangrite/internal/handlers"
	"github.com/henrriusdev/hangrite/internal/session"
)

type App struct {
	config   *config.Config
	handlers *handlers.Handlers
}

func New(config *config.Config) *App {
	session.Init(config.SessionSecret)

	return &App{
		config:   config,
		handlers: handlers.New(config),
	}
}

func (a *App) Run() error {
	fs := http.FileServer(http.Dir(a.config.Static))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	a.setupRoutes()

	log.Printf("Server starting on http://localhost%s\n", a.config.Port)
	return http.ListenAndServe(a.config.Port, nil)
}

func (a *App) setupRoutes() {
	http.HandleFunc("/", a.handlers.Home)
	http.HandleFunc("/login", a.handlers.LoginHandler)
	http.HandleFunc("/auth/login", a.handlers.LoginHandler)
	http.HandleFunc("/auth/logout", a.handlers.LogoutHandler)

	// Admin-only routes (protected by middleware)
	http.HandleFunc("/register", a.handlers.RequireAdmin(a.handlers.RegisterHandler))
	http.HandleFunc("/auth/register", a.handlers.RequireAdmin(a.handlers.RegisterHandler))

	// Player routes (protected by middleware)
	http.HandleFunc("/players", a.handlers.RequireAdmin(a.handlers.ListPlayers))
	http.HandleFunc("/players/new", a.handlers.RequireAdmin(a.handlers.NewPlayer))
	http.HandleFunc("/players/{id}/edit", a.handlers.RequireAdmin(a.handlers.EditPlayer))
	http.HandleFunc("/certificates/generate", a.handlers.RequireAdmin(a.handlers.GenerateCertificate))
}
