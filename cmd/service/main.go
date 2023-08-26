package main

import (
	"log"
	"net/http"

	"github.com/lyuvolkova/avitoTest-dynSeg/internal/app/service"
	"github.com/lyuvolkova/avitoTest-dynSeg/internal/pkg/server"
)

func main() {
	app := service.New()

	mux := server.NewMux(app)
	serviceWebServer := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Server started", serviceWebServer.Addr)

	if err := serviceWebServer.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}

}
