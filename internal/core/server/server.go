package server

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(httpServer *http.Server) *Server {
	return &Server{
		httpServer: httpServer,
	}
}

func (server *Server) Start(wg *sync.WaitGroup, closeServerChan <-chan bool) {
	go func() {
		if err := server.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Error from http server: %s", err.Error())
		}
	}()

	select {
	case <-closeServerChan:
	}

	server.httpServer.Shutdown(context.Background())
	wg.Done()
}
