package main

import (
	"context"
	"log"
	"tire/config"

	"github.com/alecthomas/kong"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sethvargo/go-envconfig"
)

var CLI struct {
	Serve serveCmd `cmd:"" help:"Run as app server"`
}

func main() {
	ctx := context.Background()
	// Parse Env vars
	var config config.Data
	if err := envconfig.Process(ctx, &config); err != nil {
		log.Fatal(err)
	}
	// Parse CLI
	kongCTX := kong.Parse(&CLI)
	err := kongCTX.Run(config)
	if err != nil {
		log.Fatalf("parsing_cli: %s", err)
	}
}
