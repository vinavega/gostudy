package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/vinavega/gostudy/internal/database"
)

func (apiCfg *apiCOnfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithErr(w, 400, fmt.Sprint("Error parsing JSON", err))
		return
	}

	if params.Name == "" {
		respondWithErr(w, 400, "Name cannot be empty")
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithErr(w, 400, fmt.Sprint("Error creating user", err))
		return
	}
	respondWithJSON(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiCOnfig) HandlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, databaseUserToUser(user))
}
