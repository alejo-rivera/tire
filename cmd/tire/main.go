package main

import (
	"log"

	"github.com/alecthomas/kong"
	_ "github.com/mattn/go-sqlite3"
)

var CLI struct {
	Serve serveCmd `cmd:"" help:"Run as app  server"`
}

func main() {
	ctx := kong.Parse(&CLI)
	err := ctx.Run()
	if err != nil {
		log.Fatalf("parsing_cli: %s", err)
	}
}
