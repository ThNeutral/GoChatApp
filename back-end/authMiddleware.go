package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/ThNeutral/messenger/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func getAPIKeyFromHeader(header http.Header) (string, error) {
	authHeaderString := header.Get("Authorization")
	if authHeaderString == "" {
		return "", errors.New("no authorization header data")
	}

	authHeader := strings.Split(authHeaderString, " ")
	if len(authHeader) != 2 || authHeader[0] != "Bearer" {
		return "", errors.New("incorrect format of Authorization header. Should be Authorization: Bearer api_key")
	}

	return authHeader[1], nil
}

func (apiCfg *apiConfig) authMiddleware(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := getAPIKeyFromHeader(r.Header)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Incorrect authorization token. Error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 404, fmt.Sprintf("No user found by this apiKey. Error: %v", err))
			return
		}

		handler(w, r, user)
	}
}
