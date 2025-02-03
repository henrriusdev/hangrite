package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/henrriusdev/hangrite/models"
	"github.com/henrriusdev/hangrite/pages/certificates"
	"github.com/henrriusdev/hangrite/pages/players"
)

func (h *Handlers) ListPlayers(w http.ResponseWriter, r *http.Request) {
	user := h.GetCurrentUser(r)

	var playerList []models.Player
	if err := h.config.DB.Find(&playerList).Error; err != nil {
		http.Error(w, "Error al cargar peloteros", http.StatusInternalServerError)
		return
	}

	players.List(playerList, user).Render(r.Context(), w)
}

func (h *Handlers) NewPlayer(w http.ResponseWriter, r *http.Request) {
	user := h.GetCurrentUser(r)

	if r.Method == http.MethodGet {
		players.Form(models.PlayerForm{
			EntryDate: time.Now(),
		}, true, user).Render(r.Context(), w)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
		return
	}

	entryDate, err := time.Parse("2006-01-02", r.FormValue("entry_date"))
	if err != nil {
		http.Error(w, "Fecha inválida", http.StatusBadRequest)
		return
	}

	player := models.Player{
		Name:      r.FormValue("name"),
		IDNumber:  r.FormValue("id_number"),
		Position:  r.FormValue("position"),
		EntryDate: entryDate,
	}

	if err := h.config.DB.Create(&player).Error; err != nil {
		players.Form(models.PlayerForm{
			ID:        player.ID,
			Name:      player.Name,
			IDNumber:  player.IDNumber,
			Position:  player.Position,
			EntryDate: player.EntryDate,
			HasError:  true,
			ErrorMsg:  "Error al crear el pelotero",
		}, true, user).Render(r.Context(), w)
		return
	}

	http.Redirect(w, r, "/players", http.StatusSeeOther)
}

func (h *Handlers) EditPlayer(w http.ResponseWriter, r *http.Request) {
	user := h.GetCurrentUser(r)
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var player models.Player
	if err := h.config.DB.First(&player, id).Error; err != nil {
		http.Error(w, "Pelotero no encontrado", http.StatusNotFound)
		return
	}

	if r.Method == http.MethodGet {
		players.Form(models.PlayerForm{
			ID:        player.ID,
			Name:      player.Name,
			IDNumber:  player.IDNumber,
			Position:  player.Position,
			EntryDate: player.EntryDate,
		}, false, user).Render(r.Context(), w)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
		return
	}

	entryDate, err := time.Parse("2006-01-02", r.FormValue("entry_date"))
	if err != nil {
		http.Error(w, "Fecha inválida", http.StatusBadRequest)
		return
	}

	updates := models.Player{
		Name:      r.FormValue("name"),
		IDNumber:  r.FormValue("id_number"),
		Position:  r.FormValue("position"),
		EntryDate: entryDate,
	}

	if err := h.config.DB.Model(&player).Updates(updates).Error; err != nil {
		players.Form(models.PlayerForm{
			ID:        player.ID,
			Name:      updates.Name,
			IDNumber:  updates.IDNumber,
			Position:  updates.Position,
			EntryDate: updates.EntryDate,
			HasError:  true,
			ErrorMsg:  "Error al actualizar el pelotero",
		}, false, user).Render(r.Context(), w)
		return
	}

	http.Redirect(w, r, "/players", http.StatusSeeOther)
}

func (h *Handlers) GenerateCertificate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
		return
	}

	playerID := r.PostForm.Get("player_id")
	if playerID == "" {
		http.Error(w, "ID de pelotero es requerido", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(playerID, 10, 64)
	if err != nil {
		http.Error(w, "ID de pelotero inválido", http.StatusBadRequest)
		return
	}

	var player models.Player
	if err := h.config.DB.First(&player, id).Error; err != nil {
		http.Error(w, "Pelotero no encontrado", http.StatusNotFound)
		return
	}

	certificate := models.Certificate{
		PlayerID:      uint(id),
		IssueDate:     time.Now(),
		TrainingHours: r.PostForm.Get("training_hours"),
	}

	if err := h.config.DB.Create(&certificate).Error; err != nil {
		http.Error(w, fmt.Sprintf("Error al generar la constancia: %v", err), http.StatusInternalServerError)
		return
	}

	certificates.Certificate(player, certificate).Render(r.Context(), w)
}
