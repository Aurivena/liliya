package server

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
	}
	logrus.Info("server started successfully")
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) {
	logrus.Info("server shutdown process started")

	if err := s.httpServer.Shutdown(ctx); err != nil {
		logrus.Error(err.Error())
	} else {
		logrus.Info("http listener shutdown successfully")
	}

	logrus.Info("server shutdown process completed successfully")
}
