package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type CommentUpdate struct {
	Grammar string `json:"grammar"`
	Kanji   string `json:"kanji"`
	Culture string `json:"culture"`
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var req CommentUpdate
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	// TODO: обновление в базе (сейчас мок)
	resp := map[string]any{
		"blockId": id,
		"updated": req,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
