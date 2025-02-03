package handlers

import (
	"net/http"

	"github.com/henrriusdev/hangrite/pages"
)

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	user := h.GetCurrentUser(r)
	pages.Home(user).Render(r.Context(), w)
}
