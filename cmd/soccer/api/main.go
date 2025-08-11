package main

import (
	"context"

	"github.com/Sup3r-Us3r/go-soccer/internal/infra/web/webserver"
)

func main() {
	webServer := webserver.NewWebServer(":8080")

	if err := webServer.Start(); err != nil {
		panic(err)
	}

	// Handle graceful shutdown
	defer func() {
		if err := webServer.Stop(context.Background()); err != nil {
			panic(err)
		}
	}()
}
