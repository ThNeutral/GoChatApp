package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ThNeutral/messenger/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func findCookieValue(cookies []*http.Cookie, value string) string {
	for _, cookie := range cookies {
		if cookie.Name == value {
			return cookie.Value
		}
	}
	return ""
}

func getAPIKeyFromRequest(header *http.Request) (string, error) {
	rawCookies := header.Cookies()
	if len(rawCookies) == 0 {
		return "", errors.New("no cookie header data")
	}

	authCookie := findCookieValue(rawCookies, "Authorization")
	if authCookie == "" {
		return "", errors.New("no authorization cookie data")
	}

	authHeader := strings.Split(authCookie, " ")
	if len(authHeader) != 2 || authHeader[0] != "Bearer" {
		return "", errors.New("incorrect format of Authorization header. Should be Authorization: Bearer api_key")
	}

	return authHeader[1], nil
}

func (apiCfg *apiConfig) authMiddleware(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		apiKey, err := getAPIKeyFromRequest(r)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Incorrect authorization token. Error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 404, fmt.Sprintf("No user found by this apiKey. Error: %v", err))
			return
		}

		if user.AccessTokenUpdatedAt.Unix()+3600 < time.Now().Unix() {
			respondWithError(w, 400, "Token has expired. To renew token, login again into app using /login-user")
			return
		}

		apiCfg.DB.UpdateAccessTokenExpiryTimeAndGetUser(r.Context(), database.UpdateAccessTokenExpiryTimeAndGetUserParams{
			Email:                user.Email,
			AccessTokenUpdatedAt: time.Now().UTC(),
		})

		handler(w, r, user)
	}
}
