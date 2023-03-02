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

	dp := dependencies.NewContainer(ctx)
	dp.Initialize()

	wg := &sync.WaitGroup{}
	wg.Add(1)

	httpServer := &http.Server{
		Addr:              dp.Conf.Server.Port,
		Handler:           dp.HttpHandler.Router,
		ReadHeaderTimeout: dp.Conf.Server.ReadHeaderTimeout,
	}
	srv := server.NewServer(ctx, httpServer)

	go srv.Start(wg)

	<-serverSignalChan

	cancel()
	wg.Wait()
}
