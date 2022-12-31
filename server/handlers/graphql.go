package handlers

import (
	"gategaurdian/server/graph"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func Graphql(c *fiber.Ctx) error {
	gql := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	gqlHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gql.ServeHTTP(w, r)
	})

	fasthttpadaptor.NewFastHTTPHandler(gqlHandler)(c.Context())
	return nil
}

func PlayGround(c *fiber.Ctx) error {
	gqlPlayground := playground.Handler("GraphQL playground", "/graphql")

	fasthttpadaptor.NewFastHTTPHandler(gqlPlayground)(c.Context())
	return nil
}
