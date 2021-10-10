package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"tire/ent"
	"tire/ent/migrate"
	"tire/graph"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"go.uber.org/zap"
)

type serveCmd struct {
	Graph bool `help: "Serve graphql server"`
	Port  int  `help: "Port to serve on" default:8181`
}

func (c *serveCmd) Run() error {
	if c.Graph {
		client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
		if err != nil {
			log.Fatal("opening ent client", err)
		}
		if err := client.Schema.Create(
			context.Background(),
			migrate.WithGlobalUniqueID(true),
		); err != nil {
			log.Fatal("opening ent client", err)
		}
		// TODO configure logger better
		logger, err := zap.NewDevelopment()
		if err != nil {
			log.Fatal("creating logger", err)
		}

		srv := handler.NewDefaultServer(graph.NewSchema(client, logger))

		http.Handle("/", playground.Handler("GraphQL playground", "/query"))
		http.Handle("/query", srv)

		log.Printf("running GQL playground at http://localhost:%d/ ", c.Port)
		log.Fatal(http.ListenAndServe(":"+fmt.Sprintf("%d", c.Port), nil))
		return nil
	}
	return errors.New("nothing to serve")
}
