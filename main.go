package main

import (
	"github.com/keploy/go-sdk/integrations/kmux"
	"github.com/keploy/go-sdk/keploy"
	"log"
)

func main() {
	a := &App{}
	err := a.Initialize(
		"postgres",
		"password",
		"postgres")

	if err != nil {
		log.Fatal("Failed to initialize app", err)
	}

	port := "8000"

	a.Router.Use(kmux.MuxMiddleware(k))
	a.Run(":8000")
}
