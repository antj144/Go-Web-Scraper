package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github./go-chi/chi"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	"github.com/lib/pg"
)

type apiConfig strut {
	DB *database.Queries
}

func main() {
	fmt.Print("hello World")

	godoenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal(("PORT is not found in the environment"))
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal(("DB_URL is not found in the environment"))
	}

	conn, sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn)

	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string, {"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))


	v1Router := chi.NewRouter
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/user", apiCfg.handlerCreateUser)
	v1Router.GET("/user", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Printf("server starting on port %v", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port:", portString)

}
