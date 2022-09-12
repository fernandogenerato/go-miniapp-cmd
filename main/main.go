package main

import (
	"log"
	"miniapp/cmd/app"
	"os"
)

// usage : go run main.go ip --host google.com.br
func main() {
	app := app.Build()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
