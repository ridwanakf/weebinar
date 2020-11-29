package main

import (
	"log"

	"github.com/ridwanakf/weebinar/internal/app"
	"github.com/ridwanakf/weebinar/internal/delivery/rest"
)

func main() {
	// init app
	WeebinarApp, err := app.NewWeebinarApp()
	if err != nil {
		log.Fatalf("error initiating app %+v", err)
	}
	defer func() {
		errs := WeebinarApp.Close()
		for e := range errs {
			log.Println(e)
		}
	}()

	rest.Start(WeebinarApp)
}
