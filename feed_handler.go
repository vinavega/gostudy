package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/vinavega/gostudy/internal/database"
)

func (apiCfg *apiCOnfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithErr(w, 400, fmt.Sprint("Error parsing JSON", err))
		return
	}

	if params.Name == "" || params.Url == "" {
		respondWithErr(w, 400, "Name  or url cannot be empty")
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithErr(w, 400, fmt.Sprint("Error creating feed", err))
		return
	}
	respondWithJSON(w, 200, databaseFeedToFeed(feed))
}
