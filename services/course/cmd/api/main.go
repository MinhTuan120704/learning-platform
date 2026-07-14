package main

import (
	"fmt"
	"log"

	"github.com/MinhTuan120704/learning-platform/services/course/internal/bootstrap"
)

func main() {
	app, err := bootstrap.New()
	if err != nil {
		log.Fatal(err)
	}

	addr := fmt.Sprintf(
		"%s:%s",
		app.Config.HTTP.Host,
		app.Config.HTTP.Port,
	)

	log.Printf("Course Service listening on %s", addr)

	if err := app.Router.Run(addr); err != nil {
		log.Fatal(err)
	}
}
