package handlers

import (
	"net/http"

	"github.com/henrriusdev/hangrite/internal/session"
	"github.com/henrriusdev/hangrite/models"
	"github.com/henrriusdev/hangrite/pages/auth"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handlers) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		auth.Login(models.LoginForm{}).Render(r.Context(), w)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	form := models.LoginForm{
		Username: username,
		Password: password,
	}

	var user models.User
	result := h.config.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		form.HasError = true
		form.ErrorMsg = "Usuario o contraseña incorrectos"
		auth.Login(form).Render(r.Context(), w)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		form.HasError = true
		form.ErrorMsg = "Usuario o contraseña incorrectos"
		auth.Login(form).Render(r.Context(), w)
		return
	}

	// Create session
	sess, err := session.Get(r)
	if err != nil {
		http.Error(w, "Error interno del servidor: "+err.Error(), http.StatusInternalServerError)
		return
	}

	sess.Values[session.UserIDKey] = user.ID
	sess.Values[session.UserRoleKey] = user.Role
	err = session.Save(r, w, sess)
	if err != nil {
		http.Error(w, "Error interno del servidor: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handlers) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	user := h.GetCurrentUser(r)

	if r.Method == http.MethodGet {
		auth.Register(models.LoginForm{}, user).Render(r.Context(), w)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	form := models.LoginForm{
		Username: username,
		Password: password,
	}

	// Check if username already exists
	var existingUser models.User
	result := h.config.DB.Where("username = ?", username).First(&existingUser)
	if result.Error == nil {
		form.HasError = true
		form.ErrorMsg = "El usuario ya existe"
		auth.Register(form, user).Render(r.Context(), w)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		form.HasError = true
		form.ErrorMsg = "Error al crear usuario"
		auth.Register(form, user).Render(r.Context(), w)
		return
	}

	newUser := models.User{
		Username: username,
		Password: string(hashedPassword),
		Role:     models.RoleUser,
	}

	result = h.config.DB.Create(&newUser)
	if result.Error != nil {
		form.HasError = true
		form.ErrorMsg = "Error al crear usuario"
		auth.Register(form, user).Render(r.Context(), w)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (h *Handlers) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sess, err := session.Get(r)
	if err != nil {
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
		return
	}

	sess.Options.MaxAge = -1
	err = session.Save(r, w, sess)
	if err != nil {
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
