package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"projects/LDmitryLD/petstore/app/internal/infrastructure/router"
	"projects/LDmitryLD/petstore/app/internal/modules"
	"projects/LDmitryLD/petstore/app/internal/storages"
	"syscall"
	"time"
)

func main() {
	storages := storages.NewStorages()

	services := modules.NewServices(storages)

	controllers := modules.NewControllers(services)

	r := router.NewRouter(controllers)

	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("Starting server")
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server error: ", err.Error())
		}
	}()

	<-sigChan

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Server stopped")
}
