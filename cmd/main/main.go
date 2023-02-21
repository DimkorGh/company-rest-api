package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"company-rest-api/internal/core/dependencies"
	"company-rest-api/internal/core/server"
)

func main() {
	serverSignalChan := make(chan os.Signal, 1)
	signal.Notify(serverSignalChan, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	depContainer := dependencies.NewContainer(ctx)
	depContainer.Initialize()

	wg := &sync.WaitGroup{}
	wg.Add(1)

	httpServer := &http.Server{
		Addr:              ":" + os.Getenv("API_PORT"),
		Handler:           depContainer.HttpHandler.Router,
		ReadHeaderTimeout: 2 * time.Second,
	}
	srv := server.NewServer(ctx, httpServer)

	go srv.Start(wg)

	<-serverSignalChan

	cancel()
	wg.Wait()
}
