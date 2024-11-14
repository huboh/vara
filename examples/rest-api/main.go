package main

import (
	"log"

	"github.com/huboh/vara"

	"github.com/huboh/vara/examples/rest-api/modules/app"
)

const (
	host = "localhost"
	port = "5009"
)

func main() {
	app, err := vara.New(&app.Module{})
	if err != nil {
		log.Fatal("failed to create vara app: ", err)
	}

	err = app.Listen(host, port)
	if err != nil {
		log.Fatal("failed to start app server: ", err)
	}
}
