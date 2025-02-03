package handlers

import (
	"net/http"

	"github.com/henrriusdev/hangrite/config"
	"github.com/henrriusdev/hangrite/internal/session"
	"github.com/henrriusdev/hangrite/models"
	"gorm.io/gorm"
)

type Handlers struct {
	db *gorm.DB
}

func New(config *config.Config) *Handlers {
	return &Handlers{
		db: config.DB,
	}
}

// Middleware to require admin role
func (h *Handlers) RequireAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := h.GetCurrentUser(r)
		if user == nil || user.Role != models.RoleAdmin {
			// If it's an AJAX request, return 401 instead of redirecting
			if r.Header.Get("X-Requested-With") == "XMLHttpRequest" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}

func (h *Handlers) GetCurrentUser(r *http.Request) *models.CurrentUser {
	sess, err := session.Get(r)
	if err != nil {
		return nil
	}

	userID, ok := sess.Values[session.UserIDKey].(uint)
	if !ok {
		return nil
	}

	var user models.User
	if err := h.db.First(&user, userID).Error; err != nil {
		return nil
	}

	return &models.CurrentUser{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	}
}
