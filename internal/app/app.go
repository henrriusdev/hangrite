package app

import (
	"log"
	"net/http"

	"github.com/henrriusdev/hangrite/config"
	"github.com/henrriusdev/hangrite/internal/handlers"
	"github.com/henrriusdev/hangrite/internal/session"
)

type App struct {
	conf     *config.Config
	handlers *handlers.Handlers
}

func New(conf *config.Config) *App {
	println(config.Env.SessionSecret)
	session.Init([]byte(config.Env.SessionSecret))

	return &App{
		conf:     conf,
		handlers: handlers.New(conf),
	}
}

func (a *App) Run() error {
	fs := http.FileServer(http.Dir(config.Env.Static))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	a.setupRoutes()

	log.Printf("Server starting on http://0.0.0.0%s\n", config.Env.Port)
	return http.ListenAndServe("0.0.0.0:8080", nil)
}

func (a *App) setupRoutes() {
	http.HandleFunc("/", a.handlers.Home)
	http.HandleFunc("/login", a.handlers.LoginHandler)
	http.HandleFunc("/auth/login", a.handlers.LoginHandler)
	http.HandleFunc("/auth/logout", a.handlers.LogoutHandler)

	// Admin-only routes (protected by middleware)
	http.HandleFunc("/register", a.handlers.RequireAdmin(a.handlers.RegisterHandler))
	http.HandleFunc("/auth/register", a.handlers.RequireAdmin(a.handlers.RegisterHandler))

	// Player routes
	http.HandleFunc("/players", a.handlers.RequireAdmin(a.handlers.ListPlayers))
	http.HandleFunc("/players/new", a.handlers.RequireAdmin(a.handlers.NewPlayer))
	http.HandleFunc("/players/{id}/edit", a.handlers.RequireAdmin(a.handlers.EditPlayer))
	http.HandleFunc("/certificates/generate", a.handlers.RequireAdmin(a.handlers.GenerateCertificate))
}
