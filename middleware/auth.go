package middleware

import (
	"fmt"
	"gostudy/internal/auth"
	"gostudy/internal/database"
	"gostudy/shared"
	"gostudy/utils"
	"net/http"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

type AuthCfg utils.ApiCOnfig

func (apiCfg *AuthCfg) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			shared.RespondWithErr(w, 400, fmt.Sprint("erro ao extrair apikey do header ", err))
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
