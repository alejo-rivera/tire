package graph

import (
	"tire/ent"

	"github.com/99designs/gqlgen/graphql"
	"go.uber.org/zap"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	dbClient *ent.Client
	*zap.Logger
}

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client, logger *zap.Logger) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: &Resolver{
			client,
			logger,
		},
	})
}
