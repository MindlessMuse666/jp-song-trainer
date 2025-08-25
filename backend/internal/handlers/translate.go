package handlers

import (
	"encoding/json"
	"net/http"
)

type TranslateRequest struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Lyrics string `json:"lyrics"`
}

type TranslateResponce struct {
	SongID string  `json:"songId"`
	Blocks []Block `json:"blocks"`
}

type Block struct {
	ID    string `json:"id"`
	Orig  string `json:"orig"`
	Trans string `json:"trans"`
}

func Translate(w http.ResponseWriter, r *http.Request) {
	var req TranslateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	// TODO: Deepseek API, но пока что - мок
	resp := TranslateResponce{
		SongID: "demo-1",
		Blocks: []Block{
			{ID: "b1", Orig: "君の名は", Trans: "Твоё имя"},
			{ID: "b2", Orig: "世界の終わり", Trans: "Конец света"},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
