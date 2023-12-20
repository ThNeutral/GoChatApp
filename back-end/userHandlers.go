package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ThNeutral/messenger/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) createUser(w http.ResponseWriter, r *http.Request) {
	type params struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	parameters := params{}
	err := decoder.Decode(&parameters)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Recieved malformed JSON. Error: %v", err))
		return
	}
	if parameters.Email == "" || parameters.Password == "" || parameters.Username == "" {
		respondWithError(w, 400, "Recieved incorrect JSON format. Expected {username: string, email: string, password: string}")
		return
	}

	_, err = apiCfg.DB.GetUserByEmail(r.Context(), parameters.Email)

	if err == nil {
		respondWithError(w, 409, "User with such email already exists")
		return
	}

	newUser, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Username:  parameters.Username,
		Email:     parameters.Email,
		Password:  parameters.Password,
	})

	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Failed to create user. Error: %v", err))
		return
	}

	type response struct {
		Message     string `json:"message"`
		AccessToken string `json:"access_token"`
	}

	respondWithJSON(w, 201, response{
		Message:     "Successfully created new user",
		AccessToken: newUser.AccessToken,
	})
}

func (apiCfg *apiConfig) loginUser(w http.ResponseWriter, r *http.Request) {
	type params struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)

	parameters := params{}
	err := decoder.Decode(&parameters)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Recieved malformed JSON. Error: %v", err))
		return
	}
	if parameters.Email == "" || parameters.Password == "" {
		respondWithError(w, 400, "Recieved incorrect JSON format. Expected {email: string, password: string}")
		return
	}

	user, err := apiCfg.DB.GetUserByEmailAndPassword(r.Context(), database.GetUserByEmailAndPasswordParams{
		Email:    parameters.Email,
		Password: parameters.Password,
	})
	if err != nil {
		respondWithError(w, 404, "User with given email and password does not exist")
		return
	}

	type response struct {
		Message     string `json:"message"`
		AccessToken string `json:"access_token"`
	}

	respondWithJSON(w, 200, response{
		Message:     "User with given information has been found",
		AccessToken: user.AccessToken,
	})
}

func (apiCfg *apiConfig) getUserProfile(w http.ResponseWriter, r *http.Request, user database.User) {
	type response struct {
		CreatedAt time.Time
		UpdatedAt time.Time
		Username  string
		Email     string
	}

	respondWithJSON(w, 200, response{
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Username:  user.Username,
		Email:     user.Password,
	})
}
