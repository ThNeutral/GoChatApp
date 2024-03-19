package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ThNeutral/messenger/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load("./.env")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Failed to get PORT env var")
	}

	db_url := os.Getenv("DB_URL")
	if db_url == "" {
		log.Fatal("Failed to get DB_URL env var")
	}

	conn, err := sql.Open("postgres", db_url)
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Post("/register-user", apiCfg.createUser)
	v1Router.Post("/login-user", apiCfg.loginUser)

	v1Router.Get("/user-profile", apiCfg.authMiddleware(apiCfg.getUserProfile))

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	fmt.Printf("Port: %v\n", port)
	log.Fatal(srv.ListenAndServe())
}
