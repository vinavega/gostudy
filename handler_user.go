package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/vinavega/gostudy/internal/database"
	"golang.org/x/crypto/bcrypt"
)

func (apiCfg *apiCOnfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
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

	if params.Name == "" || params.Password == "" {
		respondWithErr(w, 400, "Name or password cannot be empty")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), 8)
	if err != nil {
		respondWithErr(w, 400, "failed to hash password")
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Password:  string(hashedPassword),
	})
	if err != nil {
		respondWithErr(w, 400, fmt.Sprint("Error creating user", err))
		return
	}
	respondWithJSON(w, 200, user)
}
