package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"company-rest-api/internal/core/dependencies"
	"company-rest-api/internal/core/server"
)

func main() {
	serverSignalChan := make(chan os.Signal, 1)
	signal.Notify(serverSignalChan, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	dc := dependencies.NewContainer(ctx)
	dc.Initialize()

	wg := &sync.WaitGroup{}
	wg.Add(1)

	httpServer := &http.Server{
		Addr:              dc.Conf.Server.Port,
		Handler:           dc.HttpHandler.Router,
		ReadHeaderTimeout: dc.Conf.Server.ReadHeaderTimeout,
	}
	srv := server.NewServer(ctx, dc.Logger, httpServer)

	go srv.Start(wg)

	<-serverSignalChan

	cancel()
	wg.Wait()
}
