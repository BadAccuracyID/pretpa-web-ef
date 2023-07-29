package main

import (
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/badaccuracyid/tpa-web-ef/internal/graph"
	"github.com/badaccuracyid/tpa-web-ef/internal/graph/resolver"
	"github.com/badaccuracyid/tpa-web-ef/internal/middleware"
	"github.com/badaccuracyid/tpa-web-ef/internal/service"
	"github.com/badaccuracyid/tpa-web-ef/pkg/database"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
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
	graphConfig := graph.Config{Resolvers: &resolver.Resolver{DB: getDatabase}}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graphConfig))
	srv.AddTransport(transport.Websocket{})

	jwtSecret := os.Getenv("JWT_SECRET")
	jwtService := service.NewJWTService(jwtSecret)
	authMiddleware := middleware.NewAuthMiddleware(jwtService)

	router := mux.NewRouter()
	router.Use(authMiddleware.Middleware)

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	handle := cors.AllowAll().Handler(router)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, handle))
}
