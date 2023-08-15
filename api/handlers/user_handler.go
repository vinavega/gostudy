package handlers

import (
	"encoding/json"
	"fmt"
	"gostudy/db"
	"gostudy/internal/database"
	"gostudy/shared"
	"gostudy/utils"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type UserCfg db.ApiCOnfig

func (apiCfg *UserCfg) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		shared.RespondWithErr(w, 400, fmt.Sprint("Error parsing JSON", err))
		return
	}

	if params.Name == "" {
		shared.RespondWithErr(w, 400, "Name cannot be empty")
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		shared.RespondWithErr(w, 400, fmt.Sprint("Error creating user", err))
		return
	}
	shared.RespondWithJSON(w, 200, utils.DatabaseUserToUser(user))
}

func (apiCfg *UserCfg) HandlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	shared.RespondWithJSON(w, 200, utils.DatabaseUserToUser(user))
}
