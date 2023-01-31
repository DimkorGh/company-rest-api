package server

import (
	"context"
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"
)

type Server struct {
	ctx        context.Context
	HttpServer *http.Server
}

func NewServer(ctx context.Context, httpServer *http.Server) *Server {
	return &Server{
		ctx:        ctx,
		HttpServer: httpServer,
	}
}

func (srv *Server) Start(wg *sync.WaitGroup) {
	defer wg.Done()

	go func() {
		if err := srv.HttpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Error from http server: %s", err.Error())
		}
	}()

	<-srv.ctx.Done()
	if err := srv.HttpServer.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Error while closing http server: %s", err.Error())
	}
}
