package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/lyuvolkova/avitoTest-dynSeg/internal/app/service"
	"github.com/lyuvolkova/avitoTest-dynSeg/internal/pkg/server"
	"github.com/lyuvolkova/avitoTest-dynSeg/internal/pkg/storage/provider"
	"github.com/lyuvolkova/avitoTest-dynSeg/internal/pkg/storage/segmentsrepo"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pg, err := provider.Postgres(ctx, os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatalln(err)
	}
	defer pg.Close()

	sp := segmentsrepo.New(pg)
	app := service.New(sp)

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
