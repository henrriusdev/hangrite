package session

import (
	"encoding/gob"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/henrriusdev/hangrite/models"
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

	store = sessions.NewCookieStore(secret)
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 7 days
		HttpOnly: true,
	}
}

func GetStore() *sessions.CookieStore {
	return store
}

// Get session from request
func Get(r *http.Request) (*sessions.Session, error) {
	return store.Get(r, "hangrite_session")
}

// Save session and write to response
func Save(r *http.Request, w http.ResponseWriter, s *sessions.Session) error {
	return s.Save(r, w)
}
