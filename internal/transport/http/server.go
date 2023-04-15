package http

import (
	"context"
	"github.com/alibekabdrakhman1/first-lab/configs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Server(cfg *configs.Config, router http.Handler) {
	srv := &http.Server{
		Addr:    cfg.HTTP.PORT,
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
