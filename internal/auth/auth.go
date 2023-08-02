package auth

import (
	"errors"
	"net/http"
)

func GetAPIKey(header http.Header) (string, error) {
	authorization := header.Get("ApiKey")
	if authorization == "" {
		return "", errors.New("Unauthorized")
	}

	return authorization, nil
}
