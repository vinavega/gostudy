package main

import (
	"fmt"
	"net/http"

	"github.com/vinavega/gostudy/internal/auth"
	"github.com/vinavega/gostudy/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiCOnfig) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithErr(w, 400, fmt.Sprint("erro ao extrais apikey do header ", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithErr(w, 400, fmt.Sprint("error getting api_key from database", err))
			return
		}
		handler(w, r, user)
	}
}
