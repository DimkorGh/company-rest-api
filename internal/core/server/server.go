package server

import (
	"context"
	"net/http"
	"sync"

	"company-rest-api/internal/core/log"
)

type Server struct {
	ctx        context.Context
	logger     log.LoggerInt
	HttpServer *http.Server
}

func NewServer(ctx context.Context, logger log.LoggerInt, httpServer *http.Server) *Server {
	return &Server{
		ctx:        ctx,
		logger:     logger,
		HttpServer: httpServer,
	}
}

func (srv *Server) Start(wg *sync.WaitGroup) {
	defer wg.Done()

	go func() {
		if err := srv.HttpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			srv.logger.Fatalf("Error from http server: %s", err.Error())
		}
	}()

	<-srv.ctx.Done()
	if err := srv.HttpServer.Shutdown(context.Background()); err != nil {
		srv.logger.Errorf("Error while closing http server: %s", err.Error())
	}
}
