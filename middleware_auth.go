package main

import (
	"fmt"
	"gostudy/internal/auth"
	"gostudy/internal/database"
	"gostudy/shared"
	"net/http"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiCOnfig) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			shared.RespondWithErr(w, 400, fmt.Sprint("erro ao extrais apikey do header ", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			shared.RespondWithErr(w, 400, fmt.Sprint("error getting api_key from database", err))
			return
		}
		handler(w, r, user)
	}
}
