package handlers

import (
	"gostudy/shared"
	"net/http"
)

func HandlerHealth(w http.ResponseWriter, r *http.Request) {
	shared.RespondWithJSON(w, 200, struct{}{})

}
