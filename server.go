package main

import (
	"github.com/badaccuracyid/tpa-web-ef/database"
	"github.com/badaccuracyid/tpa-web-ef/middleware"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/badaccuracyid/tpa-web-ef/graph"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {
	var err = godotenv.Load()
	if err != nil {
		panic("failed to load env file: " + err.Error())
	}

	_, err = database.MigrateTables()
	if err != nil {
		panic("failed to migrate tables: " + err.Error())
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	getDatabase, err := database.GetDatabase()
	if err != nil {
		panic("failed to get database: " + err.Error())
	}
	graphConfig := graph.Config{Resolvers: &graph.Resolver{DB: getDatabase}}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graphConfig))

	router := mux.NewRouter()
	router.Use(middleware.AuthMiddleware)

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
