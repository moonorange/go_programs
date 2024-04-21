package main

import (
	"bff/client"
	"bff/graph"
	"context"
	"log"
	"net/http"
	"os"

	"connectrpc.com/connect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/protogo/gen"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func ListTasksByTag(ctx context.Context) (*gen.ListTasksByTagResponse, error) {
	qs := client.NewTaskServiceClient("localhost:8081")
	res, err := qs.ListTasksByTag(ctx, connect.NewRequest(&gen.ListTasksByTagRequest{
		TagName: "tag1",
	}))
	if err != nil {
		return nil, err
	}
	return res.Msg, nil
}
