package main

import (
	"github.com/alibekabdrakhman1/first-lab/configs"
	"github.com/alibekabdrakhman1/first-lab/internal/service"
	"github.com/alibekabdrakhman1/first-lab/internal/storage"
	"github.com/alibekabdrakhman1/first-lab/internal/transport/http"
	"github.com/alibekabdrakhman1/first-lab/internal/transport/http/handler"
	"log"
)

func main() {
	log.Fatal(run())
}

func run() error {
	cfg, err := configs.New()
	if err != nil {
		return err
	}
	repo, err := storage.New(cfg)
	if err != nil {
		return err
	}
	srv, err := service.NewManager(repo)
	if err != nil {
		return err
	}
	handler := handler.NewHandler(srv)
	router := http.Routes(handler)

	http.Server(cfg, router)
	return nil

}
