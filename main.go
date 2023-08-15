package main

import (
	"database/sql"
	"gostudy/api/handlers"
	"gostudy/internal/database"
	"gostudy/middleware"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not set in this environment")
	}

	dbURL := os.Getenv("POSTGRES_URL")
	if dbURL == "" {
		log.Fatal("POSTGRES_URL is not set in this environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Could not connect to Postgres Database")
	}

	userCfg := handlers.UserCfg{
		DB: database.New(conn),
	}

	feedCfg := handlers.FeedCfg{
		DB: database.New(conn),
	}

	authCfg := middleware.AuthCfg{
		DB: database.New(conn),
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
	v1Router.Get("/health", handlers.HandlerHealth)
	v1Router.Post("/user", userCfg.HandlerCreateUser)
	v1Router.Get("/me", authCfg.MiddlewareAuth(userCfg.HandlerGetUser))
	v1Router.Post("/feed", authCfg.MiddlewareAuth(feedCfg.HandlerCreateFeed))

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)
	erro := srv.ListenAndServe()
	if erro != nil {
		log.Fatal(erro)
	}

}
