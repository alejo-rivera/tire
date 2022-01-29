package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
	"tire/config"
	"tire/ent"
	"tire/ent/migrate"
	"tire/graph"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type serveCmd struct{}

func (c *serveCmd) Run(conf config.Data) error {
	if conf.ServerConfig.Enable {
		// Prep DB
		client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
		if err != nil {
			log.Fatal("opening ent client", err)
		}
		if err := client.Schema.Create(
			context.Background(),
			migrate.WithGlobalUniqueID(true),
		); err != nil {
			log.Fatal("creating ent schema", err)
		}

		logger, err := createLogger(conf.LoggingLevel)
		if err != nil {
			log.Fatal("creating logger", err)
		}

		router := gin.New()

		// sMiddleware
		// TODO update CORS before going to prod
		router.Use(cors.New(cors.Config{
			AllowAllOrigins: true,
			MaxAge:          12 * time.Hour,
		}))
		// Add a ginzap middleware, which:
		//   - Logs all requests, like a combined access and error log.
		//   - RFC3339 with UTC time format.
		// Logs all panic to error log
		//   - stack means whether output the stack info.
		router.Use(ginzap.Ginzap(logger, time.RFC3339, true), ginzap.RecoveryWithZap(logger, true))

		// Routes
		router.POST("/query", graphqlHandler(client, logger))
		router.GET("/", playgroundHandler())

		router.Run(fmt.Sprintf(":%d", conf.ServerConfig.Port))

		return nil
	}
	return errors.New("nothing to serve")
}

func createLogger(loggingType config.LoggingType) (*zap.Logger, error) {
	var loggerConfig zap.Config
	switch loggingType {
	case config.Debug:
		loggerConfig = zap.NewDevelopmentConfig()
		loggerConfig.Encoding = "console"
	case config.Release:
		loggerConfig = zap.NewProductionConfig()
		loggerConfig.Encoding = "json"
	}
	return loggerConfig.Build()
}

// Defining the Graphql handler
func graphqlHandler(client *ent.Client, logger *zap.Logger) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewSchema(client, logger))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
