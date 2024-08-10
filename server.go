package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/mtariq99/graphqlexample/config"
	"github.com/mtariq99/graphqlexample/graph"
)

func init() {
	if err := config.LoadConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	Resolver := graph.NewResolver()
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: Resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.Cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+config.Cfg.ServerPort, nil))
}
