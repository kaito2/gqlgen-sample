package main

import (
	"github.com/kaito2/gqlgen-sample/internal/adapter/db"
	"github.com/kaito2/gqlgen-sample/internal/adapter/graph/dataloader"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kaito2/gqlgen-sample/internal/adapter/graph"
	"github.com/kaito2/gqlgen-sample/internal/adapter/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// TODO: DI
	db := db.NewDB()
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					DB: db,
				},
			},
		),
	)

	loader := dataloader.NewLoaders(db)
	dataloaderSrv := dataloader.Middleware(loader, srv)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", dataloaderSrv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
