package main

import "log"

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
