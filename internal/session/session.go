package session

import (
	"encoding/gob"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/henrriusdev/hangrite/models"
)

const (
	// SessionName is the name of the cookie
	SessionName = "hangrite_session"
)

var (
	store *sessions.CookieStore
	// Key for getting the user ID from session
	UserIDKey = "user_id"
	// Key for getting the user role from session
	UserRoleKey = "user_role"
)

func Init(secret []byte) {
	// Register types with gob encoder
	gob.Register(models.Role(""))
	gob.Register(uint(0))

	// Create a simple cookie store with just the secret
	store = sessions.NewCookieStore(secret)

	// Configure session cookie
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 7 days
		HttpOnly: true,
		Secure:   false, // Set to true in production with HTTPS
		SameSite: http.SameSiteLaxMode,
	}
}

func GetStore() *sessions.CookieStore {
	if store == nil {
		panic("Session store not initialized")
	}
	return store
}

// Get session from request
func Get(r *http.Request) (*sessions.Session, error) {
	if store == nil {
		panic("Session store not initialized")
	}
	return store.Get(r, SessionName)
}

// Save session and write to response
func Save(r *http.Request, w http.ResponseWriter, s *sessions.Session) error {
	if store == nil {
		panic("Session store not initialized")
	}
	return s.Save(r, w)
}
