package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HistoryItem struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
}

// Мок-объекты истории
var history = []HistoryItem{
	{ID: "1", Title: "君の名は", Artist: "RADWIMPS"},
	{ID: "2", Title: "世界の終わり", Artist: "SEKAI NO OWARI"},
}

func GetHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(history)
}

func DeleteHistory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	for i, item := range history {
		if item.ID == id {
			history = append(history[:i], history[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "not found", http.StatusNotFound)
}
