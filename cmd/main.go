package main

import (
	"context"
	service2 "github.com/alibekabdrakhman1/first-lab/internal/service"
	storage2 "github.com/alibekabdrakhman1/first-lab/internal/storage"
	http2 "github.com/alibekabdrakhman1/first-lab/internal/transport/http"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alibekabdrakhman1/first-lab/configs"
)

func main() {
	run()
}

func run() {
	cfg, _ := configs.New()
	storage, err := storage2.New()
	if err != nil {
		log.Fatal(err)
	}
	service, err := service2.NewManager(storage)
	if err != nil {
		log.Fatal(err)
	}
	handler := http2.NewHandler(service)
	router := http2.Routes(handler)
	srv := &http.Server{
		Addr:    cfg.PORT,
		Handler: router,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("Server Started")

	<-done
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")

}
